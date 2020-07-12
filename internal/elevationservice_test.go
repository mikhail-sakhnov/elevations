package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mapboxMockedClient func(ctx context.Context, route Tiles) (EncodedElevationData, error)

func (m mapboxMockedClient) GetElevationData(ctx context.Context, route Tiles) (EncodedElevationData, error) {
	return m(ctx, route)
}

type decoderMock func(ctx context.Context, tiles Tiles, data EncodedElevationData) (Elevation, error)

func (d decoderMock) Decode(ctx context.Context, tiles Tiles, data EncodedElevationData) (Elevation, error) {
	return d(ctx, tiles, data)
}

func TestElevationService(t *testing.T) {
	t.Run("must_return_error_if_location_invalid", func(t *testing.T) {
		svc := ElevationService{}
		_, err := svc.GetElevation(context.Background(), []Location{{
			Longitude: 200,
			Latitude:  200,
		},
		})
		assert.Error(t, err)
	})

	t.Run("must_return_error_if_mapbox_failed", func(t *testing.T) {
		expectedErr := errors.New("random error from remote api")
		svc := ElevationService{
			mapbox: mapboxMockedClient(func(ctx context.Context, route Tiles) (EncodedElevationData, error) {
				return EncodedElevationData{}, fmt.Errorf("some error from mapbox: %w", expectedErr)
			}),
		}

		_, err := svc.GetElevation(context.Background(), []Location{{
			Longitude: 70,
			Latitude:  70,
		},
		})

		assert.True(t, errors.Is(err, expectedErr))
	})

	t.Run("must_return_error_if_mapbox_decoding_failed", func(t *testing.T) {
		expectedErr := errors.New("decoder error")
		svc := ElevationService{
			mapbox: mapboxMockedClient(func(ctx context.Context, route Tiles) (EncodedElevationData, error) {
				return EncodedElevationData{}, nil
			}),
			decoder: decoderMock(func(ctx context.Context, tiles Tiles, data EncodedElevationData) (Elevation, error) {
				return Elevation{}, fmt.Errorf("some error while decoding data: %w", expectedErr)
			}),
		}

		_, err := svc.GetElevation(context.Background(), []Location{
			{
				Longitude: 70,
				Latitude:  70,
			},
		})

		assert.True(t, errors.Is(err, expectedErr))
	})

	t.Run("must_return_elevation_from_the_decoded_data", func(t *testing.T) {
		panic("test is not implemented")
	})
}
