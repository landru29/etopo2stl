package xyz

import "github.com/landru29/etopo2stl/earth"

// Above ...
func Above(xyzAngle []earth.VectorA, plan float64) (result []earth.VectorA) {
	for _, xyzVector := range xyzAngle {
		above := earth.VectorA{
			Lon:      xyzVector.Lon,
			Lat:      xyzVector.Lat,
			Altitude: xyzVector.Altitude,
		}
		if above.Altitude >= plan {
			result = append(result, above)
		}
	}
	return
}
