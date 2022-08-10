package perfectly_spherical_houses_in_a_vacuum_part_2

import (
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

func HouseCounter(direction string) int {
	santaLocation := Coordinates{X: 0, Y: 0}
	robotLocation := Coordinates{X: 0, Y: 0}

	visitedHousesBySanta := map[Coordinates]int{}
	visitedHousesByRobot := map[Coordinates]int{}

	visitedHousesBySanta[santaLocation] = 1
	visitedHousesByRobot[robotLocation] = 0

	directions := strings.Split(direction, "")
	if len(directions) == 0 {
		return 0
	}

	for i, dir := range directions {

		if i == 0 || i%2 == 0 {
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
			visitedHousesBySanta[santaLocation] += 1
		} else {
			switch dir {
			case "^":
				robotLocation.Y += 1
			case "v":
				robotLocation.Y -= 1
			case ">":
				robotLocation.X += 1
			case "<":
				robotLocation.X -= 1
			}
			visitedHousesByRobot[robotLocation] += 1
		}
	}
	return len(visitedHousesByRobot) + len(visitedHousesBySanta) - 1
}
