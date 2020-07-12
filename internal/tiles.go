package internal

// TileCoordinatesPair represents mercator projection tile location
type TileCoordinatesPair struct {
	X int
	Y int

	From Location
}

// Tiles collection of tiles
type Tiles []TileCoordinatesPair
