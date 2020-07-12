package internal

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"log"
	"net/http"
)

const defaultZoomLevel = 15
const defaultTileAPIUrl = "https://api.mapbox.com/v4/mapbox.terrain-rgb"

// EncodedElevationData represents pngraw elevation data from mapbox API
type EncodedElevationData struct {
	png map[TileCoordinatesPair][]byte
}

// MapboxClient default mapbox clients
type MapboxClient struct {
	token   string
	baseURL string
}

// NewMapboxClient constructs new mapbox client
func NewMapboxClient(t string) *MapboxClient {
	return &MapboxClient{
		token:   t, // TODO: make through opts...
		baseURL: defaultTileAPIUrl,
	}
}

// GetElevationPNGs loads elevation data for given set of tiles
func (m MapboxClient) GetElevationPNGs(ctx context.Context, tiles Tiles) (EncodedElevationData, error) {
	// TODO: super naive and stupid implementation, do it another way
	result := EncodedElevationData{
		png: map[TileCoordinatesPair][]byte{},
	}
	for _, t := range tiles {
		finalURL := fmt.Sprintf("%s/%d/%d/%d.pngraw?access_token=%s", defaultTileAPIUrl, defaultZoomLevel, t.X, t.Y, m.token)
		spew.Dump(finalURL)
		fmt.Println(finalURL)
		resp, err := http.DefaultClient.Get(finalURL)
		if err != nil {
			log.Printf("error while requesting mapbox api: %s", err)
		}
		if resp.Body != nil {
			defer resp.Body.Close()
		}
		png, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("error while requesting mapbox api: %s", err)
		}
		result.png[t] = png
	}
	return result, nil
}
