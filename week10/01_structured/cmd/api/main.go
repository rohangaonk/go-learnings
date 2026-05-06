package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-learning/week10/01_structured/internal/config"
	"go-learning/week10/01_structured/internal/server"
	"go-learning/week10/01_structured/internal/store"
)

// main is a wiring file — load config, build deps, start server.
// All business logic lives in internal/. This file should never grow large.
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config: %v", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	slog.SetDefault(logger)

	s := store.New()
	srv := server.New(cfg, s, logger)

	// Start server in a goroutine so we can also listen for shutdown signals
	go func() {
		if err := srv.Run(); err != nil {
			logger.Error("server error", "err", err)
		}
	}()

	// Wait for SIGINT (Ctrl+C) or SIGTERM (Kubernetes, systemd)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	logger.Info("shutting down")

	// 10-second timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("shutdown error", "err", err)
	}

	logger.Info("shutdown complete")
}
