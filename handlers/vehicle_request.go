package handlers

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"restapi/models"
)

// firstCarYear is the year of the first production automobile (Benz Patent-Motorwagen).
const firstCarYear = 1886

// CreateVehicleRequest is the JSON body for POST /vehicles.
// The server generates the id. Set doors only for a car, and hasSidecar only for a motorcycle.
type CreateVehicleRequest struct {
	Type       models.VehicleType `json:"type" enums:"car,motorcycle" example:"car" validate:"required"`
	Brand      string             `json:"brand" example:"Skoda" validate:"required"`
	Model      string             `json:"model" example:"Octavia" validate:"required"`
	Year       int                `json:"year" example:"2023" validate:"required"`
	Doors      *int               `json:"doors,omitempty" example:"5"`          // required when type is car
	HasSidecar *bool              `json:"hasSidecar,omitempty" example:"false"` // optional when type is motorcycle, default false
}

// toVehicle validates the request and builds a new Car or Motorcycle with a new GUID.
func (req CreateVehicleRequest) toVehicle() (models.VehicleItem, error) {
	req.Type = models.VehicleType(strings.ToLower(string(req.Type)))
	req.Brand = strings.TrimSpace(req.Brand)
	req.Model = strings.TrimSpace(req.Model)

	if !req.Type.Valid() {
		return nil, fmt.Errorf("invalid type %q: use car or motorcycle", req.Type)
	}
	if req.Brand == "" {
		return nil, errors.New("brand is required")
	}
	if req.Model == "" {
		return nil, errors.New("model is required")
	}
	if maxYear := time.Now().Year() + 1; req.Year < firstCarYear || req.Year > maxYear {
		return nil, fmt.Errorf("year must be between %d and %d", firstCarYear, maxYear)
	}

	switch req.Type {
	case models.TypeCar:
		if req.HasSidecar != nil {
			return nil, errors.New("hasSidecar is only allowed for a motorcycle")
		}
		if req.Doors == nil {
			return nil, errors.New("doors is required for a car")
		}
		if *req.Doors < 1 || *req.Doors > 10 {
			return nil, errors.New("doors must be between 1 and 10")
		}
		return models.NewCar(req.Brand, req.Model, req.Year, *req.Doors), nil

	default: // models.TypeMotorcycle
		if req.Doors != nil {
			return nil, errors.New("doors is only allowed for a car")
		}
		hasSidecar := req.HasSidecar != nil && *req.HasSidecar
		return models.NewMotorcycle(req.Brand, req.Model, req.Year, hasSidecar), nil
	}
}
