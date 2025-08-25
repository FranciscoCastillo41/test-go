package httpserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/FranciscoCastillo41/test-go/backend/internal/domain"
	"github.com/FranciscoCastillo41/test-go/backend/internal/repository"
	"github.com/FranciscoCastillo41/test-go/backend/internal/services"
	"github.com/FranciscoCastillo41/test-go/backend/pkg/respond"
)

// WidgetsRoutes exposes /v1/widgets endpoints.
func WidgetsRoutes(svc *services.WidgetService) chi.Router {
	r := chi.NewRouter()

	// GET /v1/widgets - get all widgets
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		widgets, err := svc.GetAll(ctx)
		if err != nil {
			respond.ErrorJSON(w, http.StatusInternalServerError, "fetch failed", err.Error())
			return
		}
		respond.JSON(w, http.StatusOK, widgets)
	})

	// POST /v1/widgets
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Name  string  `json:"name"`
			Price float64 `json:"price"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			respond.ErrorJSON(w, http.StatusBadRequest, "invalid JSON", err.Error())
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		obj, err := svc.Create(ctx, body.Name, body.Price)
		if err != nil {
			if errors.Is(err, services.ErrInvalidInput) {
				respond.ErrorJSON(w, http.StatusBadRequest, "validation failed", "name required, price >= 0")
				return
			}
			respond.ErrorJSON(w, http.StatusInternalServerError, "create failed", err.Error())
			return
		}
		respond.JSON(w, http.StatusCreated, obj)
	})

	// GET /v1/widgets/{id}
	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		obj, err := svc.ByID(ctx, id)
		if err != nil {
			if errors.Is(err, repository.ErrNotFound) {
				respond.ErrorJSON(w, http.StatusNotFound, "not found", "")
				return
			}
			respond.ErrorJSON(w, http.StatusInternalServerError, "fetch failed", err.Error())
			return
		}
		respond.JSON(w, http.StatusOK, obj)
	})

	// DELETE /v1/widgets/{id} - delete one
	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		if err := svc.DeleteByID(ctx, id); err != nil {
			if errors.Is(err, services.ErrInvalidInput) {
				respond.ErrorJSON(w, http.StatusBadRequest, "invalid id", "")
				return
			}
			if errors.Is(err, repository.ErrNotFound) {
				respond.ErrorJSON(w, http.StatusNotFound, "not found", "")
				return
			}
			respond.ErrorJSON(w, http.StatusInternalServerError, "delete failed", err.Error())
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})

	// DELETE /v1/widgets - delete all
	r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		deleted, err := svc.DeleteAll(ctx)
		if err != nil {
			respond.ErrorJSON(w, http.StatusInternalServerError, "bulk delete failed", err.Error())
			return
		}
		respond.JSON(w, http.StatusOK, map[string]int64{"deleted": deleted})
	})

	// PATCH /v1/widgets/{id}
	r.Patch("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		var body struct {
			Name  *string  `json:"name,omitempty"`
			Price *float64 `json:"price,omitempty"`
		}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			respond.ErrorJSON(w, http.StatusBadRequest, "invalid JSON", err.Error())
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		obj, err := svc.UpdateByID(ctx, id, domain.WidgetUpdate{
			Name:  body.Name,
			Price: body.Price,
		})
		if err != nil {
			switch {
			case errors.Is(err, services.ErrInvalidInput):
				respond.ErrorJSON(w, http.StatusBadRequest, "validation failed", "provide non-empty name and/or price >= 0")
			case errors.Is(err, repository.ErrNotFound):
				respond.ErrorJSON(w, http.StatusNotFound, "not found", "")
			default:
				respond.ErrorJSON(w, http.StatusInternalServerError, "update failed", err.Error())
			}
			return
		}
		respond.JSON(w, http.StatusOK, obj)
	})

	return r
}
