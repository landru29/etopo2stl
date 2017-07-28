package xyz

import "github.com/landru29/etopo2stl/earth"

// Scale ...
func Scale(xyzAngle []earth.VectorA, scale float64) (xyzOut []earth.VectorA) {
	for _, xyzVector := range xyzAngle {
		xyzOut = append(
			xyzOut,
			earth.VectorA{
				Lon:      xyzVector.Lon,
				Lat:      xyzVector.Lat,
				Altitude: xyzVector.Altitude * scale,
			},
		)
	}
	return
}
