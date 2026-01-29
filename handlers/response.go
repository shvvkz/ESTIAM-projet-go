package handlers

import (
	"encoding/json"
	"net/http"
)

// WriteJSON(w http.ResponseWriter, status int, payload any)
//
// A helper function to write a JSON response.
//
// Parameters:
//   - w: The HTTP response writer.
//   - status: The HTTP status code to set for the response.
//   - payload: The data to be encoded as JSON and sent in the response.
func WriteJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}