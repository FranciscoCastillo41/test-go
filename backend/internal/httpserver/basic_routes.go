package httpserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

// BasicRoutes exposes simple demo endpoints under /v1.
func BasicRoutes() chi.Router {
	r := chi.NewRouter()

	// GET /v1/hello
	r.Get("/hello", func(w http.ResponseWriter, _ *http.Request) {
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

	// POST /v1/echo
	r.Post("/echo", func(w http.ResponseWriter, r *http.Request) {
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

	return r
}
