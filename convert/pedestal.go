package convert

import (
	"math"

	"github.com/landru29/xyz2stl/earth"
)

// GetMax ...
func GetMax(xyz []earth.VectorA) earth.VectorA {
	maxLon := xyz[0].Lon
	maxLat := xyz[0].Lat
	maxAlt := xyz[0].Altitude
	for _, xyzVector := range xyz {
		maxLon = math.Max(maxLon, xyzVector.Lon)
		maxLat = math.Max(maxLat, xyzVector.Lat)
		maxAlt = math.Max(maxAlt, xyzVector.Altitude)
	}
	return earth.VectorA{
		Lon:      maxLon,
		Lat:      maxLat,
		Altitude: maxAlt,
	}
}

// GetMin ...
func GetMin(xyz []earth.VectorA) earth.VectorA {
	minLon := xyz[0].Lon
	minLat := xyz[0].Lat
	minAlt := xyz[0].Altitude
	for _, xyzVector := range xyz {
		minLon = math.Min(minLon, xyzVector.Lon)
		minLat = math.Min(minLat, xyzVector.Lat)
		minAlt = math.Min(minAlt, xyzVector.Altitude)
	}
	return earth.VectorA{
		Lon:      minLon,
		Lat:      minLat,
		Altitude: minAlt,
	}
}

// Pedestal ...
func Pedestal(xyz []earth.VectorA, thickness float64) (xyzOut []earth.VectorA) {
	min := GetMin(xyz)
	max := GetMax(xyz)
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
