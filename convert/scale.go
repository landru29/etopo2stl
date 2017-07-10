package convert

import "github.com/landru29/xyz2stl/earth"

// Scale ...
func Scale(xyz []earth.VectorA, scale float64) (xyzOut []earth.VectorA) {
	for _, xyzVector := range xyz {
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
