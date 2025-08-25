package respond

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}

func JSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func ErrorJSON(w http.ResponseWriter, code int, msg, details string) {
	JSON(w, code, Error{Error: msg, Details: details})
}
