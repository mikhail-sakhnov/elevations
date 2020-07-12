package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/soider/elevations/internal/geo"
	"github.com/soider/elevations/internal/mapbox"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mapboxMockedClient func(ctx context.Context, route geo.Tiles) (mapbox.EncodedElevationData, error)

func (m mapboxMockedClient) GetElevationPNGs(ctx context.Context, route geo.Tiles) (mapbox.EncodedElevationData, error) {
	return m(ctx, route)
}

type decoderMock func(ctx context.Context, data mapbox.EncodedElevationData) (geo.RouteElevation, error)

func (d decoderMock) Decode(ctx context.Context, data mapbox.EncodedElevationData) (geo.RouteElevation, error) {
	return d(ctx, data)
}

func TestElevationService(t *testing.T) {
	t.Run("must_return_error_if_location_invalid", func(t *testing.T) {
		svc := ElevationService{}
		_, err := svc.GetElevation(context.Background(), []geo.Location{{
			Longitude: 200,
			Latitude:  200,
		},
		})
		assert.Error(t, err)
	})

	t.Run("must_return_error_if_mapbox_failed", func(t *testing.T) {
		expectedErr := errors.New("random error from remote api")
		svc := ElevationService{
			mapbox: mapboxMockedClient(func(ctx context.Context, route geo.Tiles) (mapbox.EncodedElevationData, error) {
				return mapbox.EncodedElevationData{}, fmt.Errorf("some error from mapbox: %w", expectedErr)
			}),
		}

		_, err := svc.GetElevation(context.Background(), []geo.Location{{
			Longitude: 70,
			Latitude:  70,
		},
		})

		assert.True(t, errors.Is(err, expectedErr))
	})

	t.Run("must_return_error_if_mapbox_decoding_failed", func(t *testing.T) {
		expectedErr := errors.New("decoder error")
		svc := ElevationService{
			mapbox: mapboxMockedClient(func(ctx context.Context, route geo.Tiles) (mapbox.EncodedElevationData, error) {
				return mapbox.EncodedElevationData{}, nil
			}),
			decoder: decoderMock(func(ctx context.Context, data mapbox.EncodedElevationData) (geo.RouteElevation, error) {
				return geo.RouteElevation{}, fmt.Errorf("some error while decoding data: %w", expectedErr)
			}),
		}

		_, err := svc.GetElevation(context.Background(), []geo.Location{
			{
				Longitude: 70,
				Latitude:  70,
			},
		})

		assert.True(t, errors.Is(err, expectedErr))
	})

	t.Run("must_return_elevation_from_the_decoded_data", func(t *testing.T) {
		expectedResult := geo.RouteElevation{
			{
				Location: geo.Location{
					Longitude: 50,
					Latitude:  50,
				},
				Elevation: 8001,
			},
		}
		svc := ElevationService{
			mapbox: mapboxMockedClient(func(ctx context.Context, route geo.Tiles) (mapbox.EncodedElevationData, error) {
				return mapbox.EncodedElevationData{}, nil
			}),
			decoder: decoderMock(func(ctx context.Context, data mapbox.EncodedElevationData) (geo.RouteElevation, error) {
				return expectedResult, nil
			}),
		}
		result, err := svc.GetElevation(context.Background(), geo.Route{
			geo.Location{
				Longitude: 50,
				Latitude:  50,
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})
}
