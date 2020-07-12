package internal

// Route represents collection of intermediate points
type Route []Location

// Valid validates route
func (r Route) Valid() bool {
	if len(r) == 0 {
		return false
	}
	for _, point := range r {
		if !point.Valid() {
			return false
		}
	}
	return true
}

// Location represents location tuple
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// Valid validates given tuple
func (l Location) Valid() bool {
	if l.Longitude > 180 {
		return false
	}
	if l.Longitude < -180 {
		return false
	}
	if l.Latitude > 90 {
		return false
	}
	if l.Latitude < -90 {
		return false
	}
	return true
}
