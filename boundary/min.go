package boundary

import (
	"math"

	"github.com/landru29/etopo2stl/earth"
)

// GetMin ...
func GetMin(xyz []earth.VectorA) earth.VectorA {
	minLon := xyz[0].Lon
	minLat := xyz[0].Lat
	minAlt := xyz[0].Altitude
	for _, xyzVector := range xyz {
		minLon = math.Min(minLon, xyzVector.Lon)
		minLat = math.Min(minLat, xyzVector.Lat)
		minAlt = math.Min(minAlt, xyzVector.Altitude)
	}
	return earth.VectorA{
		Lon:      minLon,
		Lat:      minLat,
		Altitude: minAlt,
	}
}
