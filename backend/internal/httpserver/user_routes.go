package httpserver

import (
	"encoding/json"
	"net/http"

	"github.com/FranciscoCastillo41/test-go/backend/internal/domain"
	"github.com/FranciscoCastillo41/test-go/backend/internal/services"
	"github.com/FranciscoCastillo41/test-go/backend/pkg/respond"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// userRoutes sets up user-related routes with JWT authentication
func userRoutes(r chi.Router, userSvc *services.UserService, jwtMiddleware func(http.Handler) http.Handler) {
	r.Route("/users", func(r chi.Router) {
		// All user routes require JWT authentication
		r.Use(jwtMiddleware)

		// GET /api/users/profile - Get current user's profile
		r.Get("/profile", getUserProfile(userSvc))

		// PUT /api/users/profile - Update current user's profile  
		r.Put("/profile", updateUserProfile(userSvc))

		// DELETE /api/users/profile - Delete current user's account
		r.Delete("/profile", deleteUserProfile(userSvc))

		// POST /api/users/sync - Create/sync user from Supabase Auth (auto-called on first login)
		r.Post("/sync", syncUser(userSvc))
	})
}

// getUserProfile handles GET /api/users/profile
func getUserProfile(userSvc *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get user ID from JWT middleware context
		userID, ok := GetUserIDFromContext(r)
		if !ok {
			respond.ErrorJSON(w, http.StatusUnauthorized, "User ID not found in token", "")
			return
		}

		authUserID, err := uuid.Parse(userID)
		if err != nil {
			respond.ErrorJSON(w, http.StatusBadRequest, "Invalid user ID format", "")
			return
		}

		// Get user profile
		user, err := userSvc.GetUserProfile(r.Context(), authUserID)
		if err != nil {
			respond.ErrorJSON(w, http.StatusNotFound, err.Error(), "")
			return
		}

		respond.JSON(w, http.StatusOK, user)
	}
}

// updateUserProfile handles PUT /api/users/profile
func updateUserProfile(userSvc *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get user ID from JWT middleware context
		userID, ok := GetUserIDFromContext(r)
		if !ok {
			respond.ErrorJSON(w, http.StatusUnauthorized, "User ID not found in token", "")
			return
		}

		authUserID, err := uuid.Parse(userID)
		if err != nil {
			respond.ErrorJSON(w, http.StatusBadRequest, "Invalid user ID format", "")
			return
		}

		// Parse request body
		var req domain.UpdateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respond.ErrorJSON(w, http.StatusBadRequest, "Invalid JSON", "")
			return
		}

		// Update user profile
		user, err := userSvc.UpdateUserProfile(r.Context(), authUserID, req)
		if err != nil {
			respond.ErrorJSON(w, http.StatusInternalServerError, err.Error(), "")
			return
		}

		respond.JSON(w, http.StatusOK, user)
	}
}

// deleteUserProfile handles DELETE /api/users/profile
func deleteUserProfile(userSvc *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get user ID from JWT middleware context
		userID, ok := GetUserIDFromContext(r)
		if !ok {
			respond.ErrorJSON(w, http.StatusUnauthorized, "User ID not found in token", "")
			return
		}

		authUserID, err := uuid.Parse(userID)
		if err != nil {
			respond.ErrorJSON(w, http.StatusBadRequest, "Invalid user ID format", "")
			return
		}

		// Delete user
		if err := userSvc.DeleteUser(r.Context(), authUserID); err != nil {
			respond.ErrorJSON(w, http.StatusInternalServerError, err.Error(), "")
			return
		}

		respond.JSON(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
	}
}

// syncUser handles POST /api/users/sync - Creates user record from Supabase Auth
func syncUser(userSvc *services.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get user ID from JWT middleware context
		userID, ok := GetUserIDFromContext(r)
		if !ok {
			respond.ErrorJSON(w, http.StatusUnauthorized, "User ID not found in token", "")
			return
		}

		authUserID, err := uuid.Parse(userID)
		if err != nil {
			respond.ErrorJSON(w, http.StatusBadRequest, "Invalid user ID format", "")
			return
		}

		// Parse request body to get email
		var req struct {
			Email string `json:"email"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respond.ErrorJSON(w, http.StatusBadRequest, "Invalid JSON", "")
			return
		}

		if req.Email == "" {
			respond.ErrorJSON(w, http.StatusBadRequest, "Email is required", "")
			return
		}

		// Create or get user
		user, err := userSvc.CreateOrGetUser(r.Context(), authUserID, req.Email)
		if err != nil {
			respond.ErrorJSON(w, http.StatusInternalServerError, err.Error(), "")
			return
		}

		respond.JSON(w, http.StatusOK, user)
	}
}