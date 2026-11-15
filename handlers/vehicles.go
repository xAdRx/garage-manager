package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"restapi/models"
)

// maxBodyBytes limits the size of a request body.
const maxBodyBytes = 1 << 20 // 1 MB

// VehicleDoc documents one vehicle in a response.
// It is used only by Swagger: the real items are models.Car or models.Motorcycle.
type VehicleDoc struct {
	models.Vehicle
	Doors      *int  `json:"doors,omitempty" example:"5"`          // only when type is car
	HasSidecar *bool `json:"hasSidecar,omitempty" example:"false"` // only when type is motorcycle
}

// VehiclesHandler groups the handlers for the /vehicles endpoint.
type VehiclesHandler struct {
	mu       sync.RWMutex // guards vehicles: GET and POST can run at the same time
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

	h.mu.RLock()
	result := make([]models.VehicleItem, 0, len(h.vehicles))
	for _, v := range h.vehicles {
		if filter.matches(v.Base()) {
			result = append(result, v)
		}
	}
	h.mu.RUnlock()

	writeJSON(w, http.StatusOK, result)
}

// Create handles POST /vehicles.
//
//	@Summary		Add a vehicle
//	@Description	Adds a car or motorcycle. The server generates the id. Send doors only for a car (required, 1-10) and hasSidecar only for a motorcycle (optional, default false).
//	@Tags			vehicles
//	@Accept			json
//	@Produce		json
//	@Param			vehicle	body		CreateVehicleRequest	true	"Vehicle to add"
//	@Success		201		{object}	VehicleDoc
//	@Failure		400		{string}	string	"Invalid JSON or validation error"
//	@Failure		413		{string}	string	"Request body larger than 1 MB"
//	@Router			/vehicles [post]
func (h *VehiclesHandler) Create(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var req CreateVehicleRequest
	if err := dec.Decode(&req); err != nil {
		var tooLarge *http.MaxBytesError
		if errors.As(err, &tooLarge) {
			http.Error(w, "request body larger than 1 MB", http.StatusRequestEntityTooLarge)
			return
		}
		http.Error(w, fmt.Sprintf("invalid JSON: %v", err), http.StatusBadRequest)
		return
	}
	if dec.More() {
		http.Error(w, "invalid JSON: body must contain a single object", http.StatusBadRequest)
		return
	}

	vehicle, err := req.toVehicle()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.mu.Lock()
	h.vehicles = append(h.vehicles, vehicle)
	h.mu.Unlock()

	writeJSON(w, http.StatusCreated, vehicle)
}

// writeJSON sends v as a JSON response with the given status code.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
