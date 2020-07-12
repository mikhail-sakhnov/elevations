package internal

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type mapboxClient interface {
	GetTileElevationData(ctx context.Context, l []TileCoordinatesPair) (EncodedElevationData, error)
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
	tiles := make([]TileCoordinatesPair, len(route))
	for i, loc := range route {
		tiles[i] = LatLonToTile(loc)
	}
	spew.Dump(tiles)
	//
	elevationData, err := es.mapbox.GetTileElevationData(ctx, tiles)
	spew.Dump(elevationData, err)
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
