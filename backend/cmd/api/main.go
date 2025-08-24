package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/FranciscoCastillo41/test-go/backend/internal/config"
	"github.com/FranciscoCastillo41/test-go/backend/internal/httpserver"
)

func main() {
	// Load config (env vars with defaults)
	cfg := config.Load()

	// Build router with a couple of simple endpoints
	router := httpserver.BuildRouter()

	// Create HTTP server with safe timeouts
	server := &http.Server{
		Addr:              cfg.HTTPAddress,
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	// Start server in the background
	go func() {
		log.Printf("server starting on %s (env=%s)", cfg.HTTPAddress, cfg.AppEnvironment)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server listen error: %v", err)
		}
	}()

	// Wait for Ctrl+C / SIGTERM, then shutdown gracefully
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}
	log.Println("server stopped")
}
