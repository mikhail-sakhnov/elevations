package internal

import "context"

// Elevation represents elevation value
type Elevation struct {
}

// NewElevationService builds new elevation service
func NewElevationService() *ElevationService {
	return &ElevationService{}
}

// ElevationService service to get elevation data for a location
type ElevationService struct {
}

// GetElevation gets elevation data for the given location
func (es ElevationService) GetElevation(ctx context.Context, location Location) (Elevation, error) {
	return Elevation{}, nil
}
