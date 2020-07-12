package mapbox

import (
	"bytes"
	"context"
	"fmt"
	"github.com/soider/elevations/internal/geo"
	"image"
	"image/color"
	"image/png"
)

// ElevationDecoder decodes elevation data from pngraw format
type ElevationDecoder struct{}

// Decode decodes elevation data from pngraw format
// Every rawpng file is png 256x256 size
// To avoid projecting real lat\long pair to the actual pixel in the tile
// we take 4 points (middle points of each quadrant of the tile)
// and calculate average tile elevation based on the elevations at those points
// Not production ready
// To be production ready:
// - metrics for cache hit
// - metrics for decode duration
// - cache for tiles elevation
// -
func (m ElevationDecoder) Decode(ctx context.Context, data EncodedElevationData) (geo.RouteElevation, error) {
	var result geo.RouteElevation
	for tileCoord, rawPng := range data.png {
		image, err := png.Decode(bytes.NewBuffer(rawPng))
		if err != nil {
			return result, fmt.Errorf("broken png file from the mapbox")
		}
		result = append(result,
			geo.Elevation{
				Location: tileCoord.From,
				Elevation: getAverageElevation(
					getColorAtTheMiddleOfTopLeftQuadrant(image),
					getColorAtTheMiddleOfTopRightQuadrant(image),
					getColorAtTheMiddleOfBottomLeftQuadrant(image),
					getColorAtTheMiddleOfBottomRightQuadrant(image),
				),
			},
		)
	}
	return result, nil
}

func getColorAtTheMiddleOfTopLeftQuadrant(i image.Image) color.Color {
	return i.At(64, 64)
}

func getColorAtTheMiddleOfTopRightQuadrant(i image.Image) color.Color {
	return i.At(64+128, 64)
}

func getColorAtTheMiddleOfBottomLeftQuadrant(i image.Image) color.Color {
	return i.At(64, 64+128)
}

func getColorAtTheMiddleOfBottomRightQuadrant(i image.Image) color.Color {
	return i.At(64+128, 64+128)
}

func getAverageElevation(a, b, c, d color.Color) float64 {
	return (getElevationAtThePoint(a) + getElevationAtThePoint(a) + getElevationAtThePoint(a) + getElevationAtThePoint(a)) / 4
}

func getElevationAtThePoint(a color.Color) float64 {
	r, g, b, _ := a.RGBA()
	return -10000 + ((float64(r)*256 + float64(g)*256 + float64(b)) * 0.1)
}

// NewMapboxElevationDecoder constructor
func NewMapboxElevationDecoder() *ElevationDecoder {
	return &ElevationDecoder{}
}
