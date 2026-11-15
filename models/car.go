package models

// Car is a Vehicle with car-specific fields.
type Car struct {
	Vehicle
	Doors *int `json:"doors,omitempty" example:"5"` // nil when the number of doors is not known
}

// NewCar creates a Car with a new GUID. Pass doors = 0 when the number of doors is not known.
func NewCar(brand, model string, year, doors int) Car {
	car := Car{Vehicle: NewVehicle(TypeCar, brand, model, year)}
	if doors > 0 {
		car.Doors = &doors
	}
	return car
}
