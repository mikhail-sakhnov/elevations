package internal

import (
	"context"
	"fmt"
)

type mapboxClient interface {
	GetElevationData(ctx context.Context, tiles Tiles) (EncodedElevationData, error)
}

type decoder interface {
	Decode(ctx context.Context, tiles Tiles, data EncodedElevationData) (Elevation, error)
}

// Elevation represents elevation value
type Elevation struct {
}

// RouteElevation represents elevation values for a route
type RouteElevation []Elevation

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
	tiles := make(Tiles, len(route))
	for i, loc := range route {
		tiles[i] = LatLonToTile(loc)
	}

	elevationData, err := es.mapbox.GetElevationData(ctx, tiles)

	if err != nil {
		return elevation, fmt.Errorf("mapbox api error: %w", err)
	}

	elevation, err = es.decoder.Decode(ctx, tiles, elevationData)

	if err != nil {
		return elevation, fmt.Errorf("decode error for route: %w", err)
	}

	return elevation, nil
}
