package boundary

import (
	"math"

	"github.com/landru29/etopo2stl/earth"
)

// GetStep ...
func GetStep(xyz []earth.VectorA) float64 {
	return math.Max(
		math.Abs(xyz[0].Lon-xyz[1].Lon),
		math.Abs(xyz[0].Lat-xyz[1].Lat),
	)
}
