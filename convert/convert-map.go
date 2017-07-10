package convert

import (
	"math"

	"github.com/landru29/xyz2stl/earth"
)

// GetOrigin ...
func GetOrigin(xyz []earth.VectorA) (result earth.VectorA) {
	result = earth.VectorA{
		Lat:      xyz[0].Lat,
		Lon:      xyz[0].Lon,
		Altitude: xyz[0].Altitude,
	}

	for _, xyzVector := range xyz {
		result = earth.VectorA{
			Lat:      math.Min(result.Lat, xyzVector.Lat),
			Lon:      math.Min(result.Lon, xyzVector.Lon),
			Altitude: math.Min(result.Altitude, xyzVector.Altitude),
		}
	}

	return
}

// MoveToOrigin ...
func MoveToOrigin(xyz []earth.VectorA, withAltitude bool) (result []earth.VectorA) {
	orig := GetOrigin(xyz)
	for _, xyzVector := range xyz {
		moved := earth.VectorA{
			Lon:      xyzVector.Lon - orig.Lon,
			Lat:      xyzVector.Lat - orig.Lat,
			Altitude: xyzVector.Altitude,
		}
		if withAltitude {
			moved.Altitude -= orig.Altitude
		}
		result = append(result, moved)
	}
	return
}

// Map ...
func PolarToLength(xyz []earth.VectorA) (xyzLength []earth.VectorL) {
	xyzOrigin := MoveToOrigin(xyz, true)
	for _, xyzVector := range xyzOrigin {
		xyzLength = append(xyzLength, earth.Angle2length(xyzVector, xyzVector.Lon))
	}
	return
}
