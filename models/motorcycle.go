package models

// Motorcycle is a Vehicle with motorcycle-specific fields.
type Motorcycle struct {
	Vehicle
	HasSidecar bool `json:"hasSidecar" example:"false"`
}

// NewMotorcycle creates a Motorcycle with a new GUID.
func NewMotorcycle(brand, model string, year int, hasSidecar bool) Motorcycle {
	return Motorcycle{
		Vehicle:    NewVehicle(brand, model, year),
		HasSidecar: hasSidecar,
	}
}
