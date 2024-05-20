package utils

import (
	"fmt"
	"io"
	"net/http"

	"github.com/goccy/go-json"
)

func ParseJSON(r *http.Request, payload any) error {
	err := json.NewDecoder(r.Body).Decode(payload)
	if err == io.EOF {
		return fmt.Errorf("missing request body")
	}
	return nil
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}
