package convert

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/landru29/xyz2stl/earth"
)

// ReadXyz ...
func ReadXyz(file io.Reader) (data []earth.VectorA, err error) {
	//data = make([]earth.VectorA, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		xyz := strings.Split(scanner.Text(), " ")
		if len(xyz) == 3 {
			lon, err := strconv.ParseFloat(xyz[0], 64)
			if err != nil {
				log.Fatal(err)
			}

			lat, err := strconv.ParseFloat(xyz[1], 64)
			if err != nil {
				log.Fatal(err)
			}

			altitude, err := strconv.ParseFloat(xyz[2], 64)
			if err != nil {
				log.Fatal(err)
			}

			data = append(data, earth.VectorA{
				Lon:      lon,
				Lat:      lat,
				Altitude: altitude,
			})
		}
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return

}
