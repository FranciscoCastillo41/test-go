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
	if cfg.SupabaseJWTSecret == "" {
		log.Fatal("SUPABASE_JWT_SECRET is required")
	}

	// database (pgxpool)
	pool, err := store.OpenPostgres(cfg.DBURL)
	if err != nil {
		log.Fatalf("db open: %v", err)
	}
	defer pool.Close()

	// migrations - run all migration files in order
	migrationFiles := []string{"migrations/001_init.sql", "migrations/002_init.sql"}
	for _, migFile := range migrationFiles {
		if mig, err := os.ReadFile(migFile); err == nil {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			if _, err := pool.Exec(ctx, string(mig)); err != nil {
				cancel()
				log.Fatalf("migration %s failed: %v", migFile, err)
			}
			cancel()
			log.Printf("migration %s applied successfully", migFile)
		} else {
			log.Printf("warning: %s not found: %v", migFile, err)
		}
	}

	// dependency injection (repos → services → router)
	widgetRepo := repository.NewPGWidgetRepo(pool)
	widgetSvc := services.NewWidgetService(widgetRepo)
	
	userRepo := repository.NewPGUserRepo(pool)
	userSvc := services.NewUserService(userRepo)

	router := httpserver.BuildRouter(httpserver.Deps{
		Widgets: widgetSvc,
		Users:   userSvc,
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
