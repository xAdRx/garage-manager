package handlers

import (
	"encoding/json"
	"net/http"
)

// CarsResponse is the JSON body returned by GET /cars.
type CarsResponse struct {
	Message string `json:"message"`
}

// CarsHandler groups the handlers for the /cars endpoint.
type CarsHandler struct{}

// NewCarsHandler creates an CarsHandler.
func NewCarsHandler() *CarsHandler {
	return &CarsHandler{}
}

// Get handles GET /cars.
//
//	@Summary	Get cars
//	@Tags		cars
//	@Produce	json
//	@Success	200	{object}	CarsResponse
//	@Router		/cars [get]
func (h *CarsHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(CarsResponse{Message: "Hello from /cars"})
}
