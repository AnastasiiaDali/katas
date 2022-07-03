package perfectly_spherical_houses_in_a_vacuum

import (
	"strings"

	"golang.org/x/exp/slices"
)

//input: "^>v<", want: 4
//axis x and y
// ^ = y + 1
// v = y - 1
// > = x + 1
// < = x - 1

func HouseCounter(direction string) int {
	houses := 1
	coordinatesXY := [][]int{{0, 0}}
	directions := strings.Split(direction, "")

	for _, dir := range directions {
		lastCoordinates := coordinatesXY[len(coordinatesXY)-1]

		switch dir {
		case "^":
			if contains(coordinatesXY, []int{lastCoordinates[0], lastCoordinates[1] + 1}) {
				coordinatesXY = append(coordinatesXY, []int{lastCoordinates[0], lastCoordinates[1] + 1})
			} else {
				coordinatesXY = append(coordinatesXY, []int{lastCoordinates[0], lastCoordinates[1] + 1})
				houses += 1
			}
		case "v":
			if contains(coordinatesXY, []int{lastCoordinates[0], lastCoordinates[1] - 1}) {
				coordinatesXY = append(coordinatesXY, []int{lastCoordinates[0], lastCoordinates[1] - 1})
			} else {
				coordinatesXY = append(coordinatesXY, []int{lastCoordinates[0], lastCoordinates[1] - 1})
				houses += 1
			}
		case ">":
			if contains(coordinatesXY, []int{lastCoordinates[0] + 1, lastCoordinates[1]}) {
				coordinatesXY = append(coordinatesXY, []int{lastCoordinates[0] + 1, lastCoordinates[1]})
			} else {
				coordinatesXY = append(coordinatesXY, []int{lastCoordinates[0] + 1, lastCoordinates[1]})
				houses += 1
			}
		case "<":
			if contains(coordinatesXY, []int{lastCoordinates[0] - 1, lastCoordinates[1]}) {
				coordinatesXY = append(coordinatesXY, []int{lastCoordinates[0] - 1, lastCoordinates[1]})
			} else {
				coordinatesXY = append(coordinatesXY, []int{lastCoordinates[0] - 1, lastCoordinates[1]})
				houses += 1
			}
		}
	}
	return houses
}

func contains(s [][]int, e []int) bool {
	for _, a := range s {
		if slices.Equal(a, e) {
			return true
		}
	}
	return false
}
