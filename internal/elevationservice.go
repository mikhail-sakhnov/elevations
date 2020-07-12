package internal

import (
	"context"
	"fmt"
)

type mapboxClient interface {
	GetTileElevationData(ctx context.Context, l MercatorCoordinates) (EncodedElevationData, error)
}

type decoder interface {
	Decode(ctx context.Context, location Location, data EncodedElevationData) (Elevation, error)
}

// Elevation represents elevation value
type Elevation struct {
}

// NewElevationService builds new elevation service
func NewElevationService(m mapboxClient, d decoder) *ElevationService {
	return &ElevationService{
		mapbox:  m,
		decoder: d,
	}
}

// ElevationService service to get elevation data for a location
type ElevationService struct {
	mapbox  mapboxClient
	decoder decoder
}

// GetElevation gets elevation data for the given location
func (es ElevationService) GetElevation(ctx context.Context, route Route) (Elevation, error) {
	var elevation Elevation
	if !route.Valid() {
		return elevation, fmt.Errorf("invalid route %v", route)
	}
	//coord := LatLonToMercator(location)
	//
	//elevationData, err := es.mapbox.GetTileElevationData(ctx, coord)
	//if err != nil {
	//	return elevation, fmt.Errorf("mapbox api error: %w", err)
	//}
	//
	//elevation, err = es.decoder.Decode(ctx, location, elevationData)
	//if err != nil {
	//	return elevation, fmt.Errorf("decode error for location `%v`: %w", location, err)
	//}

	return elevation, nil
}
