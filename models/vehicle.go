package models

import "github.com/google/uuid"

// VehicleType identifies the concrete kind of a vehicle.
type VehicleType string

const (
	TypeCar        VehicleType = "car"
	TypeMotorcycle VehicleType = "motorcycle"
)

// Valid reports whether t is a known vehicle type.
func (t VehicleType) Valid() bool {
	return t == TypeCar || t == TypeMotorcycle
}

// Vehicle holds the fields shared by all vehicle types.
// Embed it in a concrete type (Car, Motorcycle) to reuse its fields and methods.
type Vehicle struct {
	ID    uuid.UUID   `json:"id" swaggertype:"string" format:"uuid" example:"3f2b8c1e-9d4a-4e6b-8a7f-1c2d3e4f5a6b"`
	Type  VehicleType `json:"type" enums:"car,motorcycle" example:"car"`
	Brand string      `json:"brand" example:"Volkswagen"`
	Model string      `json:"model" example:"Golf"`
	Year  int         `json:"year" example:"2020"`
}

// NewVehicle creates a Vehicle with a new random (v4) GUID.
func NewVehicle(vehicleType VehicleType, brand, model string, year int) Vehicle {
	return Vehicle{
		ID:    uuid.New(),
		Type:  vehicleType,
		Brand: brand,
		Model: model,
		Year:  year,
	}
}

// Name returns the brand and model, e.g. "Volkswagen Golf".
func (v Vehicle) Name() string {
	return v.Brand + " " + v.Model
}

// Base returns the shared Vehicle fields. Car and Motorcycle get this method
// through embedding, so both satisfy VehicleItem.
func (v Vehicle) Base() Vehicle {
	return v
}

// VehicleItem is any concrete vehicle (Car, Motorcycle). It lets one list
// hold different vehicle types.
type VehicleItem interface {
	Base() Vehicle
}
