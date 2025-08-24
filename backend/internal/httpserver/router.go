package httpserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/FranciscoCastillo41/test-go/backend/internal/config"
)

// BuildRouter returns an http.Handler with basic routes and safe defaults.
func BuildRouter() http.Handler {

	cfg := config.Load()

	// Middlewares
	r := chi.NewRouter()
	r.Use(SimpleCORS(cfg.CORSOrigins))
	r.Use(RequestLogger)

	// Health check
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	// Versioned API
	r.Route("/v1", func(v1 chi.Router) {
		// GET demo
		v1.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
			type hello struct {
				Message   string `json:"message"`
				Timestamp string `json:"timestamp"`
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(hello{
				Message:   "Hello from Go 👋",
				Timestamp: time.Now().UTC().Format(time.RFC3339),
			})
		})

		// POST echo demo
		v1.Post("/echo", func(w http.ResponseWriter, r *http.Request) {
			var body map[string]any
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				http.Error(w, "invalid JSON", http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(map[string]any{
				"received": body,
				"note":     "It works! You posted JSON and I echoed it back.",
			})
		})
	})

	return r
}
