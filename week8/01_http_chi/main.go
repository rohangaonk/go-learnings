package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// ---------------------------------------------------------------------------
// Domain
// ---------------------------------------------------------------------------

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// ---------------------------------------------------------------------------
// In-memory store
//
// RWMutex: multiple readers can hold the lock simultaneously, but a writer
// gets exclusive access. Right tool when reads >> writes.
// In Node you'd never think about this — the event loop serialises everything.
// ---------------------------------------------------------------------------

type Store struct {
	mu     sync.RWMutex
	users  map[int]User
	nextID int
}

func NewStore() *Store {
	return &Store{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (s *Store) Create(name string, age int) User {
	s.mu.Lock()
	defer s.mu.Unlock()
	u := User{ID: s.nextID, Name: name, Age: age}
	s.users[s.nextID] = u
	s.nextID++
	return u
}

func (s *Store) Get(id int) (User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	u, ok := s.users[id]
	return u, ok
}

func (s *Store) Update(id int, name string, age int) (User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	u, ok := s.users[id]
	if !ok {
		return User{}, false
	}
	u.Name = name
	u.Age = age
	s.users[id] = u
	return u, true
}

func (s *Store) Delete(id int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.users[id]
	if !ok {
		return false
	}
	delete(s.users, id)
	return true
}

// ---------------------------------------------------------------------------
// Handlers
//
// Handlers are methods on a struct that holds dependencies (the store).
// This is the Go alternative to Express's closure-over-dependencies pattern:
//   const handler = (store) => (req, res) => { ... }
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

	user := h.store.Create(body.Name, body.Age)
	writeJSON(w, http.StatusCreated, user)
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := urlParamInt(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	user, ok := h.store.Get(id)
	if !ok {
		writeError(w, http.StatusNotFound, "user not found")
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

	user, ok := h.store.Update(id, body.Name, body.Age)
	if !ok {
		writeError(w, http.StatusNotFound, "user not found")
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

	if !h.store.Delete(id) {
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
	store := NewStore()
	h := &Handler{store: store}

	r := chi.NewRouter()

	// Middleware — equivalent to app.use() in Express
	r.Use(middleware.Logger)    // logs method, path, status, latency
	r.Use(middleware.Recoverer) // catches panics, returns 500 instead of crashing

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
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
