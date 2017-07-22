package boundary

import (
	"math"

	"github.com/landru29/etopo2stl/earth"
)

// GetMax ...
func GetMax(xyz []earth.VectorA) earth.VectorA {
	maxLon := xyz[0].Lon
	maxLat := xyz[0].Lat
	maxAlt := xyz[0].Altitude
	for _, xyzVector := range xyz {
		maxLon = math.Max(maxLon, xyzVector.Lon)
		maxLat = math.Max(maxLat, xyzVector.Lat)
		maxAlt = math.Max(maxAlt, xyzVector.Altitude)
	}
	return earth.VectorA{
		Lon:      maxLon,
		Lat:      maxLat,
		Altitude: maxAlt,
	}
}
