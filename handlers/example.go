package handlers

import (
	"encoding/json"
	"net/http"
)

// ExampleResponse is the JSON body returned by GET /example.
type ExampleResponse struct {
	Message string `json:"message"`
}

// ExampleHandler groups the handlers for the /example endpoint.
type ExampleHandler struct{}

// NewExampleHandler creates an ExampleHandler.
func NewExampleHandler() *ExampleHandler {
	return &ExampleHandler{}
}

// Get handles GET /example.
func (h *ExampleHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ExampleResponse{Message: "Hello from /example"})
}
