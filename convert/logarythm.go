package convert

import (
	"math"

	"github.com/landru29/xyz2stl/earth"
)

// GetMaxAltitude ...
func GetMaxAltitude(xyz []earth.VectorA) (alt float64) {
	for _, xyzVector := range xyz {
		alt = math.Max(alt, xyzVector.Altitude)
	}
	return
}

// GetMinAltitude ...
func GetMinAltitude(xyz []earth.VectorA) (alt float64) {
	for _, xyzVector := range xyz {
		alt = math.Min(alt, xyzVector.Altitude)
	}
	return
}

// Logarythm ...
func Logarythm(xyz []earth.VectorA) (xyzOut []earth.VectorA) {
	maxAltitude := GetMaxAltitude(xyz)
	minAltitude := GetMinAltitude(xyz)
	positiveFactor := float64(1)
	negativeFactor := float64(-1)
	if maxAltitude != 0 {
		positiveFactor = maxAltitude / math.Log10(maxAltitude)
	}
	if minAltitude != 0 {
		negativeFactor = minAltitude / math.Log10(-minAltitude)
	}
	for _, xyzVector := range xyz {
		altitude := float64(0)
		if xyzVector.Altitude > 0 {
			altitude = math.Log10(xyzVector.Altitude) * positiveFactor
		}
		if xyzVector.Altitude < 0 {
			altitude = math.Log10(-xyzVector.Altitude) * negativeFactor
		}
		xyzOut = append(
			xyzOut,
			earth.VectorA{
				Lon:      xyzVector.Lon,
				Lat:      xyzVector.Lat,
				Altitude: altitude,
			},
		)
	}
	return
}
