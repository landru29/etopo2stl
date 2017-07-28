package convert

import (
	"fmt"

	"github.com/landru29/etopo2stl/earth"
)

// GenerateStl ...
func GenerateStl(xyzMeter [][]earth.VectorL, sizeX int, sizeY int) {
	numYLinesMaxIndex := len(xyzMeter) - 1
	for indexY, currentYLine := range xyzMeter {
		if indexY < numYLinesMaxIndex {
			nextYLine := xyzMeter[indexY+1]
			numXLinesMaxIndex := len(currentYLine) - 1
			for indexX, currentPoint := range currentYLine {
				if indexX < numXLinesMaxIndex {
					pointRight := currentYLine[indexX+1]
					pointBottom := nextYLine[indexX]
					pointBottomRight := nextYLine[indexX+1]

					// Process here
					fmt.Printf("%v\n%v\n%v\n%v\n", currentPoint, pointRight, pointBottom, pointBottomRight)
				}
			}
		}
	}
}
