package models

import "github.com/google/uuid"

// Vehicle holds the fields shared by all vehicle types.
// Embed it in a concrete type (Car, Motorcycle) to reuse its fields and methods.
type Vehicle struct {
	ID    uuid.UUID `json:"id" swaggertype:"string" format:"uuid" example:"3f2b8c1e-9d4a-4e6b-8a7f-1c2d3e4f5a6b"`
	Brand string    `json:"brand" example:"Volkswagen"`
	Model string    `json:"model" example:"Golf"`
	Year  int       `json:"year" example:"2020"`
}

// NewVehicle creates a Vehicle with a new random (v4) GUID.
func NewVehicle(brand, model string, year int) Vehicle {
	return Vehicle{
		ID:    uuid.New(),
		Brand: brand,
		Model: model,
		Year:  year,
	}
}

// Name returns the brand and model, e.g. "Volkswagen Golf".
func (v Vehicle) Name() string {
	return v.Brand + " " + v.Model
}
