package models

// Car is a Vehicle with car-specific fields.
type Car struct {
	Vehicle
	Doors int `json:"doors" example:"5"`
}

// NewCar creates a Car with a new GUID.
func NewCar(brand, model string, year, doors int) Car {
	return Car{
		Vehicle: NewVehicle(TypeCar, brand, model, year),
		Doors:   doors,
	}
}
