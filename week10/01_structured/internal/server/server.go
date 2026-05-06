package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"

	"go-learning/week10/01_structured/internal/config"
	"go-learning/week10/01_structured/internal/store"
)

// Server holds all dependencies.
//
// Node.js bridge: like an Express app class that you'd instantiate with
//   new App({ store, logger, config })
// except Go is explicit — no DI framework, no magic. You see exactly what
// each component needs.
type Server struct {
	router *chi.Mux
	store  *store.Store
	log    *slog.Logger
	cfg    config.Config
	httpSrv    *http.Server
}

func New(cfg config.Config, s *store.Store, log *slog.Logger) *Server {
	srv := &Server{
		store: s,
		log:   log,
		cfg:   cfg,
	}
	srv.router = srv.buildRouter()
	srv.httpSrv = &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      srv.router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	return srv
}

// Run starts the HTTP server. Blocks until the server stops.
func (s *Server) Run() error {
	s.log.Info("server starting", "port", s.cfg.Port)
	if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Shutdown gracefully stops the server with the given context deadline.
// Used for clean shutdown on SIGTERM/SIGINT.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpSrv.Shutdown(ctx)
}

func (s *Server) buildRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(s.recoverer)

	r.Post("/users", s.createUser)
	r.Get("/users/{id}", s.getUser)
	r.Put("/users/{id}", s.updateUser)
	r.Delete("/users/{id}", s.deleteUser)

	return r
}

// recoverer is a middleware that catches panics, logs them at Error level
// with the panic value and stack trace, then returns a 500.
//
// Node.js bridge: like an Express error-handling middleware
//   app.use((err, req, res, next) => { logger.error(err); res.status(500)... })
// except in Go, panics aren't errors — they're a separate mechanism, and you
// must use recover() inside a deferred function to catch them.
func (s *Server) recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if val := recover(); val != nil {
				s.log.Error("panic recovered",
					"panic", fmt.Sprintf("%v", val),
					"stack", string(debug.Stack()),
					"method", r.Method,
					"path", r.URL.Path,
				)
				writeError(w, http.StatusInternalServerError, "internal server error")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// ---------------------------------------------------------------------------
// Handlers
// ---------------------------------------------------------------------------

func (s *Server) createUser(w http.ResponseWriter, r *http.Request) {
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

	user := s.store.Create(body.Name, body.Age)
	s.log.Info("user created", "user_id", user.ID, "name", user.Name)
	writeJSON(w, http.StatusCreated, user)
}

func (s *Server) getUser(w http.ResponseWriter, r *http.Request) {
	id, err := urlParamInt(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	user, ok := s.store.Get(id)
	if !ok {
		s.log.Info("user not found", "user_id", id)
		writeError(w, http.StatusNotFound, "user not found")
		return
	}

	s.log.Info("user fetched", "user_id", id)
	writeJSON(w, http.StatusOK, user)
}

func (s *Server) updateUser(w http.ResponseWriter, r *http.Request) {
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

	user, ok := s.store.Update(id, body.Name, body.Age)
	if !ok {
		s.log.Info("user not found for update", "user_id", id)
		writeError(w, http.StatusNotFound, "user not found")
		return
	}

	s.log.Info("user updated", "user_id", id)
	writeJSON(w, http.StatusOK, user)
}

func (s *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := urlParamInt(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if !s.store.Delete(id) {
		s.log.Info("user not found for delete", "user_id", id)
		writeError(w, http.StatusNotFound, "user not found")
		return
	}

	s.log.Info("user deleted", "user_id", id)
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
