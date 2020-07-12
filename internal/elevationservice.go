package internal

import (
	"context"
	"fmt"
)

type mapboxClient interface {
	GetElevationPNGs(ctx context.Context, tiles Tiles) (EncodedElevationData, error)
}

type decoder interface {
	Decode(ctx context.Context, data EncodedElevationData) (RouteElevation, error)
}

// Elevation represents elevation value
type Elevation struct {
	Location  Location `json:"location"`
	Elevation float64  `json:"elevation"`
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
func (es ElevationService) GetElevation(ctx context.Context, route Route) (RouteElevation, error) {
	var routeElevation RouteElevation
	if !route.Valid() {
		return routeElevation, fmt.Errorf("invalid route %v", route)
	}
	tiles := make(Tiles, len(route))
	for i, loc := range route {
		tiles[i] = LatLonToTile(loc)
	}

	rawPngElevation, err := es.mapbox.GetElevationPNGs(ctx, tiles)

	if err != nil {
		return routeElevation, fmt.Errorf("mapbox api error: %w", err)
	}

	routeElevation, err = es.decoder.Decode(ctx, rawPngElevation)

	if err != nil {
		return routeElevation, fmt.Errorf("decode error for route: %w", err)
	}

	return routeElevation, nil
}
