package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocation_Valid(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input Location
		valid bool
	}{
		{
			name: "valid_location",
			input: Location{
				Latitude:  70,
				Longitude: 10,
			},
			valid: true,
		},
		{
			name: "invalid_lat",
			input: Location{
				Latitude:  91,
				Longitude: 0,
			},
			valid: false,
		},
		{
			name: "invalid_lat",
			input: Location{
				Latitude:  -91,
				Longitude: 0,
			},
			valid: false,
		},
		{
			name: "invalid_long",
			input: Location{
				Latitude:  0,
				Longitude: 182,
			},
			valid: false,
		},
		{
			name: "invalid_long",
			input: Location{
				Latitude:  0,
				Longitude: -182,
			},
			valid: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.valid, tc.input.Valid())
		})
	}
}
