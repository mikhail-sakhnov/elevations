package internal

import (
	"context"
	"fmt"
	"github.com/soider/elevations/internal/geo"
	"github.com/soider/elevations/internal/mapbox"
)

type mapboxClient interface {
	GetElevationPNGs(ctx context.Context, tiles geo.Tiles) (mapbox.EncodedElevationData, error)
}

type decoder interface {
	Decode(ctx context.Context, data mapbox.EncodedElevationData) (geo.RouteElevation, error)
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
func (es ElevationService) GetElevation(ctx context.Context, route geo.Route) (geo.RouteElevation, error) {
	var routeElevation geo.RouteElevation
	if !route.Valid() {
		return routeElevation, fmt.Errorf("invalid route %v", route)
	}
	tiles := make(geo.Tiles, len(route))
	for i, loc := range route {
		tiles[i] = geo.LatLonToTile(loc)
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
