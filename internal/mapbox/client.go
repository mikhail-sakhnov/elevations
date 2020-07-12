package mapbox

import (
	"context"
	"fmt"
	"github.com/soider/elevations/internal/geo"
	"io/ioutil"
	"log"
	"net/http"
)

const defaultTileAPIUrl = "https://api.mapbox.com/v4/mapbox.terrain-rgb"

// EncodedElevationData represents pngraw elevation data from mapbox API
type EncodedElevationData struct {
	png map[geo.TileCoordinatesPair][]byte
}

// Client default mapbox clients
type Client struct {
	token   string
	baseURL string
}

// NewClient constructs new mapbox client
func NewClient(t string) *Client {
	return &Client{
		token:   t, // TODO: make through opts...
		baseURL: defaultTileAPIUrl,
	}
}

// GetElevationPNGs loads elevation data for given set of tiles
// To be production ready:
// - do requests concurrently
// - add pooling (each png has the same size, should align perfectly with sync.Pool)
// - deduplicate tiles to avoid requesting same tile few times
// - circuit breaker and rate limiter for the outgoing requests
// - tune http clients (keep alive could be an option here, because all requests go to the same host)
// - metrics for requests duration
// - metrics for cache hit
func (m Client) GetElevationPNGs(ctx context.Context, tiles geo.Tiles) (EncodedElevationData, error) {
	result := EncodedElevationData{
		png: map[geo.TileCoordinatesPair][]byte{},
	}
	for _, t := range tiles {
		finalURL := fmt.Sprintf("%s/%d/%d/%d.pngraw?access_token=%s", defaultTileAPIUrl, t.Z, t.X, t.Y, m.token)
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
