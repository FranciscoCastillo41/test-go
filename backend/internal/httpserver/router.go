package httpserver

import (
	"net/http"

	"github.com/FranciscoCastillo41/test-go/backend/internal/config"
	"github.com/FranciscoCastillo41/test-go/backend/internal/services"
	"github.com/go-chi/chi/v5"
)

// Deps collects the services your routes need.
type Deps struct {
	Widgets *services.WidgetService
	Users   *services.UserService
	// Add more later:
	// Auth   *services.AuthService
	// Orders *services.OrderService
}

// BuildRouter is the single entry point that sets global middleware,
// health checks, and mounts versioned subrouters (e.g., /v1).
func BuildRouter(deps Deps) http.Handler {
	cfg := config.Load()

	r := chi.NewRouter()

	// Global middleware
	r.Use(SimpleCORS(cfg.CORSOrigins))
	r.Use(RequestLogger)
	// r.Use(SimpleRPSLimit(50, 100)) // if you added the limiter

	// Create JWT middleware
	jwtMiddleware := SupabaseJWTVerifier(cfg.SupabaseURL, cfg.SupabaseJWTSecret)

	// Health/liveness
	r.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	// Versioned API
	r.Route("/v1", func(v1 chi.Router) {
		// Base demo endpoints (/v1/hello, /v1/echo) - public
		v1.Mount("/", BasicRoutes())

		// Feature endpoints (/v1/widgets/...) - public for now
		v1.Mount("/widgets", WidgetsRoutes(deps.Widgets))

		// User endpoints (/v1/users/...) - protected with JWT
		userRoutes(v1, deps.Users, jwtMiddleware)

		// More features later:
		// v1.Mount("/auth",   AuthRoutes(deps.Auth))
		// v1.Mount("/orders", OrdersRoutes(deps.Orders))
	})

	return r
}
