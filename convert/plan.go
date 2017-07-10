package convert

import "github.com/landru29/xyz2stl/earth"

// Above ...
func Above(xyz []earth.VectorA, plan float64) (result []earth.VectorA) {
	for _, xyzVector := range xyz {
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
