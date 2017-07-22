package convert

import (
	"math"

	"github.com/landru29/etopo2stl/earth"
)

// GetOrigin ...
func GetOrigin(xyz []earth.VectorA) (result earth.VectorA) {
	result = earth.VectorA{
		Lat:      xyz[0].Lat,
		Lon:      xyz[0].Lon,
		Altitude: xyz[0].Altitude,
	}

	for _, xyzVector := range xyz {
		result = earth.VectorA{
			Lat:      math.Min(result.Lat, xyzVector.Lat),
			Lon:      math.Min(result.Lon, xyzVector.Lon),
			Altitude: math.Min(result.Altitude, xyzVector.Altitude),
		}
	}

	return
}

// MoveToOrigin ...
func MoveToOrigin(xyz []earth.VectorA, withAltitude bool) (result []earth.VectorA) {
	orig := GetOrigin(xyz)
	for _, xyzVector := range xyz {
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

// PolarToLength ...
func PolarToLength(xyz []earth.VectorA) (xyzLength [][]earth.VectorL, sizeX int, sizeY int) {
	xyzOrigin := MoveToOrigin(xyz, true)
	lastPoint := xyzOrigin[0]
	var x int
	for _, xyzVector := range xyzOrigin {
		if (lastPoint.Lat != xyzVector.Lat) && (lastPoint.Lon != xyzVector.Lon) {
			sizeY++
			var line []earth.VectorL
			xyzLength = append(xyzLength, line)
			sizeX = x
			x = 0
		}
		x++
		xyzLength[sizeY] = append(xyzLength[sizeY], earth.Angle2length(xyzVector, xyzVector.Lon))
	}
	sizeY++
	return
}
