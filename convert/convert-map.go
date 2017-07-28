package convert

import (
	"github.com/landru29/etopo2stl/earth"
	"github.com/landru29/etopo2stl/xyz"
)

// PolarToLength ...
func PolarToLength(xyzAngle []earth.VectorA) (xyzMeter [][]earth.VectorL, sizeX int, sizeY int) {
	xyzOrigin := xyz.MoveToOrigin(xyzAngle, true)
	lastPoint := xyzOrigin[0]
	var x int
	for _, xyzVector := range xyzOrigin {
		if (lastPoint.Lat != xyzVector.Lat) && (lastPoint.Lon != xyzVector.Lon) {
			sizeY++
			var line []earth.VectorL
			xyzMeter = append(xyzMeter, line)
			sizeX = x
			x = 0
		}
		x++
		xyzMeter[sizeY] = append(xyzMeter[sizeY], earth.Angle2length(xyzVector, xyzVector.Lon))
	}
	sizeY++
	return
}
