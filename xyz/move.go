package xyz

import (
	"math"

	"github.com/landru29/etopo2stl/earth"
)

// GetOrigin ...
func GetOrigin(xyzAngle []earth.VectorA) (result earth.VectorA) {
	result = earth.VectorA{
		Lat:      xyzAngle[0].Lat,
		Lon:      xyzAngle[0].Lon,
		Altitude: xyzAngle[0].Altitude,
	}

	for _, xyzVector := range xyzAngle {
		result = earth.VectorA{
			Lat:      math.Min(result.Lat, xyzVector.Lat),
			Lon:      math.Min(result.Lon, xyzVector.Lon),
			Altitude: math.Min(result.Altitude, xyzVector.Altitude),
		}
	}

	return
}

// MoveToOrigin ...
func MoveToOrigin(xyzAngle []earth.VectorA, withAltitude bool) (result []earth.VectorA) {
	orig := GetOrigin(xyzAngle)
	for _, xyzVector := range xyzAngle {
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
