package perfectly_spherical_houses_in_a_vacuum

import (
	"strings"
)

//input: "^>v<", want: 4
//axis x and y
// ^ = y + 1
// v = y - 1
// > = x + 1
// < = x - 1

type Coordinates struct {
	X int
	Y int
}

func HouseCounter(direction string) int {
	directions := strings.Split(direction, "")
	santaLocation := Coordinates{X: 0, Y: 0}
	visitedHousesCoordinatesXY := map[Coordinates]int{}

	// on starting point Santa already delivers present to one house
	visitedHousesCoordinatesXY[santaLocation] = 1

	if len(directions) == 0 {
		return 0
	}

	for _, dir := range directions {
		switch dir {
		case "^":
			santaLocation.Y += 1
		case "v":
			santaLocation.Y -= 1
		case ">":
			santaLocation.X += 1
		case "<":
			santaLocation.X -= 1
		}
		visitedHousesCoordinatesXY[santaLocation] += 1
	}
	return len(visitedHousesCoordinatesXY)
}
