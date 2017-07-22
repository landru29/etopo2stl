package convert

import (
	"github.com/landru29/etopo2stl/boundary"
	"github.com/landru29/etopo2stl/earth"
)

// Pedestal ...
func Pedestal(xyz []earth.VectorA, thickness float64) (xyzOut []earth.VectorA) {
	min := boundary.GetMin(xyz)
	max := boundary.GetMax(xyz)
	for _, xyzVector := range xyz {
		xyzOut = append(
			xyzOut,
			xyzVector,
		)
	}
	xyzOut = append(
		xyzOut,
		earth.VectorA{
			Lon:      min.Lon,
			Lat:      min.Lat,
			Altitude: min.Altitude - thickness,
		},
	)
	xyzOut = append(
		xyzOut,
		earth.VectorA{
			Lon:      min.Lon,
			Lat:      max.Lat,
			Altitude: min.Altitude - thickness,
		},
	)
	xyzOut = append(
		xyzOut,
		earth.VectorA{
			Lon:      max.Lon,
			Lat:      max.Lat,
			Altitude: min.Altitude - thickness,
		},
	)
	xyzOut = append(
		xyzOut,
		earth.VectorA{
			Lon:      max.Lon,
			Lat:      min.Lat,
			Altitude: min.Altitude - thickness,
		},
	)
	return
}
