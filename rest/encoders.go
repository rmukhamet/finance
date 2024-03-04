package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// encode example err := encode(w, r, http.StatusOK, obj)
func encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}

// decode example decoded, err := decode[CreateSomethingRequest](r)
func decode[T any](r *http.Request) (T, error) {
	var v T

	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}

	return v, nil
}
