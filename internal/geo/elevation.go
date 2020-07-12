package geo

// Elevation represents elevation value
type Elevation struct {
	Location  Location `json:"location"`
	Elevation float64  `json:"elevation"`
}

// RouteElevation represents elevation values for a route
type RouteElevation []Elevation
