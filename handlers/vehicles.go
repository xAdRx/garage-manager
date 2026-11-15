package handlers

import (
	"encoding/json"
	"net/http"

	"restapi/models"
)

// VehicleDoc documents one item of the GET /vehicles response.
// It is used only by Swagger: the real items are models.Car or models.Motorcycle.
type VehicleDoc struct {
	models.Vehicle
	Doors      *int  `json:"doors,omitempty" example:"5"`          // only when type is car
	HasSidecar *bool `json:"hasSidecar,omitempty" example:"false"` // only when type is motorcycle
}

// VehiclesHandler groups the handlers for the /vehicles endpoint.
type VehiclesHandler struct {
	vehicles []models.VehicleItem
}

// NewVehiclesHandler creates a VehiclesHandler with the default in-memory vehicles.
func NewVehiclesHandler() *VehiclesHandler {
	return &VehiclesHandler{
		vehicles: []models.VehicleItem{
			models.NewCar("Volkswagen", "Golf", 2020, 5),
			models.NewCar("Toyota", "Corolla", 2022, 4),
			models.NewMotorcycle("Honda", "CB500F", 2021, false),
		},
	}
}

// Get handles GET /vehicles.
//
//	@Summary		List vehicles
//	@Description	Returns all vehicles. Every filter is optional; filters combine with AND.
//	@Tags			vehicles
//	@Produce		json
//	@Param			type	query		string	false	"Vehicle type"	Enums(car, motorcycle)
//	@Param			brand	query		string	false	"Brand, case-insensitive exact match"
//	@Param			model	query		string	false	"Model, case-insensitive exact match"
//	@Param			year	query		int		false	"Exact production year"
//	@Param			minYear	query		int		false	"Production year from (inclusive)"
//	@Param			maxYear	query		int		false	"Production year to (inclusive)"
//	@Success		200		{array}		VehicleDoc
//	@Failure		400		{string}	string	"Invalid filter value"
//	@Router			/vehicles [get]
func (h *VehiclesHandler) Get(w http.ResponseWriter, r *http.Request) {
	filter, err := parseVehicleFilter(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := make([]models.VehicleItem, 0, len(h.vehicles))
	for _, v := range h.vehicles {
		if filter.matches(v.Base()) {
			result = append(result, v)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
