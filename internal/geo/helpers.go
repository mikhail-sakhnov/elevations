package geo

import (
	"github.com/buckhx/tiles"
)

const defaultZoomLevel = 15

// LatLonToTile converts Location to mercator x,y tuple
func LatLonToTile(l Location) TileCoordinatesPair {
	t := tiles.FromCoordinate(l.Latitude, l.Longitude, defaultZoomLevel)
	return TileCoordinatesPair{
		X:    t.X,
		Y:    t.Y,
		Z:    t.Z,
		From: l,
	}
}
