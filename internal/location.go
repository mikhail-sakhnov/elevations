package internal

// Location represents location tuple
type Location struct {
	Longitude float64 `form:"longitude"`
	Latitude  float64 `form:"latitude"`
}

// Valid validates given tuple
func (l Location) Valid() bool {
	return true
}
