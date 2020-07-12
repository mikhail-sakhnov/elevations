package internal

import (
	"context"
)

// MapboxElevationDecoder 
type MapboxElevationDecoder struct{}

// Decode decodes elevation data from pngraw format
func (m MapboxElevationDecoder) Decode(ctx context.Context, location Location, data EncodedElevationData) (Elevation, error) {
	panic("implement me")
}

// NewMapboxElevationDecoder constructor
func NewMapboxElevationDecoder() *MapboxElevationDecoder {
	return &MapboxElevationDecoder{}
}
