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
	"github.com/FranciscoCastillo41/test-go/backend/internal/repository"
	"github.com/FranciscoCastillo41/test-go/backend/internal/services"
	"github.com/FranciscoCastillo41/test-go/backend/internal/store"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// config
	cfg := config.Load()
	if cfg.DBURL == "" {
		log.Fatal("DB_URL is required (Supabase Postgres URI, e.g. postgres://user:pass@host:5432/postgres?sslmode=require)")
	}

	// database (pgxpool)
	pool, err := store.OpenPostgres(cfg.DBURL)
	if err != nil {
		log.Fatalf("db open: %v", err)
	}
	defer pool.Close()

	// migrations (MVP: run a single SQL file at startup)
	if mig, err := os.ReadFile("migrations/001_init.sql"); err == nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if _, err := pool.Exec(ctx, string(mig)); err != nil {
			log.Fatalf("migration: %v", err)
		}
	} else {
		log.Printf("warning: migrations/001_init.sql not found: %v", err)
	}

	// dependency injection (repos → services → router)
	widgetRepo := repository.NewPGWidgetRepo(pool)
	widgetSvc := services.NewWidgetService(widgetRepo)

	router := httpserver.BuildRouter(httpserver.Deps{
		Widgets: widgetSvc,
	})

	// HTTP server with timeouts
	srv := &http.Server{
		Addr:              cfg.HTTPAddress, // e.g. ":8080"
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	// start
	go func() {
		log.Printf("server starting on %s (env=%s)", cfg.HTTPAddress, cfg.AppEnvironment)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
	log.Println("server stopped")
}
