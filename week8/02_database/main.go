package main

// Week 9 Exercise — Replace the in-memory Store from week8/01_http_chi
// with a real PostgreSQL backend using pgxpool.
//
// Before running:
//   1. Start Postgres and create the DB:
//        createdb go_learning
//   2. Apply the schema:
//        psql -d go_learning -f schema.sql
//   3. Set the connection string env var:
//        export DATABASE_URL="postgres://localhost:5432/go_learning?sslmode=disable"
//   4. Run:
//        go run main.go
//
// Node.js bridge:
//   pgxpool.New  ≈  new pg.Pool({ ... })
//   pool.QueryRow ≈  pool.query (single row)
//   row.Scan     ≈  destructuring the result row
//   pgx.ErrNoRows ≈  result.rows.length === 0

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ---------------------------------------------------------------------------
// Domain
// ---------------------------------------------------------------------------

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

// ---------------------------------------------------------------------------
// Store — now backed by Postgres instead of an in-memory map.
//
// The handler layer (below) is IDENTICAL to week8/01_http_chi — only the
// store implementation changes. This is the Go interface/struct pattern doing
// its job: swappable internals, stable API surface.
// ---------------------------------------------------------------------------

type Store struct {
	pool *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{pool: pool}
}

// Create inserts a new user and returns the full record (including DB-generated id).
// $1, $2 are positional parameters — Postgres style (not ? like MySQL/SQLite).
// RETURNING lets us get the generated id and created_at in a single round-trip.
func (s *Store) Create(ctx context.Context, name string, age int) (User, error) {
	var u User
	err := s.pool.QueryRow(ctx,
		`INSERT INTO users (name, age) VALUES ($1, $2)
		 RETURNING id, name, age, created_at`,
		name, age,
	).Scan(&u.ID, &u.Name, &u.Age, &u.CreatedAt)
	return u, err
}

// Get fetches a single user by id.
// pgx.ErrNoRows is returned when no row matches — handle it like you'd check
// result.rows.length === 0 in Node, but as an explicit error value.
func (s *Store) Get(ctx context.Context, id int) (User, error) {
	var u User
	err := s.pool.QueryRow(ctx,
		`SELECT id, name, age, created_at FROM users WHERE id = $1`,
		id,
	).Scan(&u.ID, &u.Name, &u.Age, &u.CreatedAt)
	return u, err // caller checks errors.Is(err, pgx.ErrNoRows)
}

// Update modifies name and age for the given id.
// If no row matched, pgx.ErrNoRows is returned via RETURNING's empty result set.
func (s *Store) Update(ctx context.Context, id int, name string, age int) (User, error) {
	var u User
	err := s.pool.QueryRow(ctx,
		`UPDATE users SET name=$2, age=$3 WHERE id=$1
		 RETURNING id, name, age, created_at`,
		id, name, age,
	).Scan(&u.ID, &u.Name, &u.Age, &u.CreatedAt)
	return u, err
}

// Delete removes a user. CommandTag.RowsAffected() tells us if anything was deleted.
func (s *Store) Delete(ctx context.Context, id int) (bool, error) {
	tag, err := s.pool.Exec(ctx,
		`DELETE FROM users WHERE id = $1`, id,
	)
	if err != nil {
		return false, err
	}
	return tag.RowsAffected() > 0, nil
}

// ---------------------------------------------------------------------------
// Handlers — identical structure to week8/01_http_chi.
// Only change: methods now accept ctx and return errors instead of bool.
// ---------------------------------------------------------------------------

type Handler struct {
	store *Store
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if body.Name == "" {
		writeError(w, http.StatusUnprocessableEntity, "name is required")
		return
	}

	user, err := h.store.Create(r.Context(), body.Name, body.Age)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not create user")
		return
	}

	writeJSON(w, http.StatusCreated, user)
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := urlParamInt(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	user, err := h.store.Get(r.Context(), id)
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, http.StatusNotFound, "user not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not fetch user")
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := urlParamInt(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var body struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := h.store.Update(r.Context(), id, body.Name, body.Age)
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, http.StatusNotFound, "user not found")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not update user")
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := urlParamInt(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	deleted, err := h.store.Delete(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "could not delete user")
		return
	}
	if !deleted {
		writeError(w, http.StatusNotFound, "user not found")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

func urlParamInt(r *http.Request, key string) (int, error) {
	raw := chi.URLParam(r, key)
	if raw == "" {
		return 0, errors.New("missing param")
	}
	return strconv.Atoi(raw)
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

func main() {
	// DATABASE_URL example: "postgres://localhost:5432/go_learning?sslmode=disable"
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL env var is required")
	}

	// pgxpool.New creates a managed connection pool.
	// The context here controls how long we wait for the initial pool setup.
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("unable to create connection pool: %v", err)
	}
	defer pool.Close()

	// Ping verifies the connection before we start serving traffic.
	// In Node you'd do pool.connect() then release; this is the Go equivalent.
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("unable to reach database: %v", err)
	}
	log.Println("connected to database")

	store := NewStore(pool)
	h := &Handler{store: store}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/users", h.createUser)
	r.Get("/users/{id}", h.getUser)
	r.Put("/users/{id}", h.updateUser)
	r.Delete("/users/{id}", h.deleteUser)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("listening on :8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
