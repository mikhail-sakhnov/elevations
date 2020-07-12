package internal

import "github.com/go-spatial/proj"

// LatLonToMercator converts Location to mercator x,y tuple
func LatLonToMercator(l Location) MercatorCoordinates {
	xy, err := proj.Convert(proj.WebMercator, []float64{l.Longitude, l.Latitude})
	if err != nil {
		// almost impossible to do anything
		panic(err)
	}
	return MercatorCoordinates{
		X: xy[0],
		Y: xy[1],
	}
}
