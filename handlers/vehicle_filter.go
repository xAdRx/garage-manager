package handlers

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"restapi/models"
)

// vehicleFilter holds the optional query filters for GET /vehicles.
// Empty or zero fields mean "no filter".
type vehicleFilter struct {
	vehicleType models.VehicleType
	brand       string
	model       string
	year        int
	minYear     int
	maxYear     int
}

// parseVehicleFilter reads filters from the query string.
// It returns an error for an unknown type or a non-numeric year value.
func parseVehicleFilter(q url.Values) (vehicleFilter, error) {
	f := vehicleFilter{
		vehicleType: models.VehicleType(strings.ToLower(q.Get("type"))),
		brand:       q.Get("brand"),
		model:       q.Get("model"),
	}

	if f.vehicleType != "" && !f.vehicleType.Valid() {
		return f, fmt.Errorf("invalid type %q: use car or motorcycle", q.Get("type"))
	}

	var err error
	if f.year, err = parseIntParam(q, "year"); err != nil {
		return f, err
	}
	if f.minYear, err = parseIntParam(q, "minYear"); err != nil {
		return f, err
	}
	if f.maxYear, err = parseIntParam(q, "maxYear"); err != nil {
		return f, err
	}
	return f, nil
}

// parseIntParam returns 0 when the parameter is missing.
func parseIntParam(q url.Values, name string) (int, error) {
	raw := q.Get(name)
	if raw == "" {
		return 0, nil
	}
	n, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("invalid %s %q: must be a number", name, raw)
	}
	return n, nil
}

// matches reports whether v passes every set filter.
// Brand and model match case-insensitively on the full value.
func (f vehicleFilter) matches(v models.Vehicle) bool {
	switch {
	case f.vehicleType != "" && v.Type != f.vehicleType:
		return false
	case f.brand != "" && !strings.EqualFold(v.Brand, f.brand):
		return false
	case f.model != "" && !strings.EqualFold(v.Model, f.model):
		return false
	case f.year != 0 && v.Year != f.year:
		return false
	case f.minYear != 0 && v.Year < f.minYear:
		return false
	case f.maxYear != 0 && v.Year > f.maxYear:
		return false
	}
	return true
}
