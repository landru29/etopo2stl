package xyz

import "github.com/landru29/etopo2stl/earth"

// Offset ...
func Offset(xyzAngle []earth.VectorA, offset float64) (xyzOut []earth.VectorA) {

	for _, xyzVector := range xyzAngle {
		xyzOut = append(
			xyzOut,
			earth.VectorA{
				Lon:      xyzVector.Lon,
				Lat:      xyzVector.Lat,
				Altitude: xyzVector.Altitude + offset,
			},
		)
	}
	return
}
