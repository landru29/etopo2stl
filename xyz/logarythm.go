package xyz

import (
	"math"

	"github.com/landru29/etopo2stl/boundary"
	"github.com/landru29/etopo2stl/earth"
)

// Logarythm ...
func Logarythm(xyzAngle []earth.VectorA) (xyzOut []earth.VectorA) {
	boundMax := boundary.GetMax(xyzAngle)
	boundMin := boundary.GetMin(xyzAngle)
	positiveFactor := float64(1)
	negativeFactor := float64(-1)
	if boundMax.Altitude != 0 {
		positiveFactor = boundMax.Altitude / math.Log10(boundMax.Altitude)
	}
	if boundMin.Altitude != 0 {
		negativeFactor = boundMin.Altitude / math.Log10(-boundMin.Altitude)
	}
	for _, xyzVector := range xyzAngle {
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
