package internal

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

const defaultZoomLevel = 15
const defaultTileAPIUrl = "https://api.mapbox.com/v4/mapbox.terrain-rgb"

// EncodedElevationData represents pngraw elevation data from mapbox API
type EncodedElevationData struct {
}

// MapboxClient default mapbox clients
type MapboxClient struct {
	token   string
	baseURL string
}

func NewMapboxClient(t string) *MapboxClient {
	return &MapboxClient{
		token:   t, // TODO: make through opts...
		baseURL: defaultTileAPIUrl,
	}
}

// GetTile loads tile from mapbox api for the given location
func (m MapboxClient) GetTileElevationData(ctx context.Context, l MercatorCoordinates) (EncodedElevationData, error) {
	finalURL := fmt.Sprintf("%s/%d/%f/%f.pngraw?access_token=%s", defaultTileAPIUrl, defaultZoomLevel, l.X, l.Y, m.token)
	spew.Dump(finalURL)
	fmt.Println(finalURL)
	panic("implement me")
}
