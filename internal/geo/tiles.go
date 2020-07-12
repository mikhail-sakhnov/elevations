package geo

// TileCoordinatesPair represents mercator projection tile location
type TileCoordinatesPair struct {
	X int
	Y int
	Z int

	From Location
}

// Tiles collection of tiles
type Tiles []TileCoordinatesPair
