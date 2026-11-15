package handlers

import (
	"encoding/json"
	"net/http"

	"restapi/models"
)

// CarsHandler groups the handlers for the /cars endpoint.
type CarsHandler struct {
	cars []models.Car
}

// NewCarsHandler creates a CarsHandler with sample in-memory data.
func NewCarsHandler() *CarsHandler {
	return &CarsHandler{
		cars: []models.Car{
			models.NewCar("Volkswagen", "Golf", 2020, 5),
			models.NewCar("Toyota", "Corolla", 2022, 4),
			models.NewCar("Mazda", "MX-5", 2019, 2),
		},
	}
}

// Get handles GET /cars.
//
//	@Summary	Get cars
//	@Tags		cars
//	@Produce	json
//	@Success	200	{array}	models.Car
//	@Router		/cars [get]
func (h *CarsHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.cars)
}
