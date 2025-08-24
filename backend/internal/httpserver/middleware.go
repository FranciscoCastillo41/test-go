package httpserver

import (
	"log"
	"net/http"
	"strings"
	"time"
)

// RequestLogger logs method, path, and duration.
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s - %dms", r.Method, r.URL.Path, time.Since(start).Milliseconds())
	})
}

// SimpleCORS applies a basic CORS policy for the given comma-separated origins.
// Example: "http://localhost:3000,https://your-frontend.vercel.app"
// Use "*" to allow all (dev only).
func SimpleCORS(allowed string) func(http.Handler) http.Handler {
	originSet := map[string]struct{}{}
	for _, o := range strings.Split(allowed, ",") {
		if o = strings.TrimSpace(o); o != "" {
			originSet[o] = struct{}{}
		}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")

			if allowed == "*" {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			} else if _, ok := originSet[origin]; ok {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Vary", "Origin")
			}

			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
