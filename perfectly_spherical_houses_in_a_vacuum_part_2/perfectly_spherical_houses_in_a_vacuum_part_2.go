package perfectly_spherical_houses_in_a_vacuum_part_2

import (
	"strings"

	"golang.org/x/exp/slices"
)

func HouseCounter(direction string) int {
	housesRobotSantaVisited := 1
	housesRealSantaVisited := 0
	coordinatesXYRobotSanta := [][]int{{0, 0}}
	coordinatesXYRealSanta := [][]int{{0, 0}}
	directions := strings.Split(direction, "")

	for i, dir := range directions {

		if i == 0 || i%2 == 0 {
			lastCoordinates := coordinatesXYRobotSanta[len(coordinatesXYRobotSanta)-1]
			switch dir {
			case "^":
				if contains(coordinatesXYRobotSanta, []int{lastCoordinates[0], lastCoordinates[1] + 1}) {
					coordinatesXYRobotSanta = append(coordinatesXYRobotSanta, []int{lastCoordinates[0], lastCoordinates[1] + 1})
				} else {
					coordinatesXYRobotSanta = append(coordinatesXYRobotSanta, []int{lastCoordinates[0], lastCoordinates[1] + 1})
					housesRobotSantaVisited += 1
				}
			case "v":
				if contains(coordinatesXYRobotSanta, []int{lastCoordinates[0], lastCoordinates[1] - 1}) {
					coordinatesXYRobotSanta = append(coordinatesXYRobotSanta, []int{lastCoordinates[0], lastCoordinates[1] - 1})
				} else {
					coordinatesXYRobotSanta = append(coordinatesXYRobotSanta, []int{lastCoordinates[0], lastCoordinates[1] - 1})
					housesRobotSantaVisited += 1
				}
			case ">":
				if contains(coordinatesXYRobotSanta, []int{lastCoordinates[0] + 1, lastCoordinates[1]}) {
					coordinatesXYRobotSanta = append(coordinatesXYRobotSanta, []int{lastCoordinates[0] + 1, lastCoordinates[1]})
				} else {
					coordinatesXYRobotSanta = append(coordinatesXYRobotSanta, []int{lastCoordinates[0] + 1, lastCoordinates[1]})
					housesRobotSantaVisited += 1
				}
			case "<":
				if contains(coordinatesXYRobotSanta, []int{lastCoordinates[0] - 1, lastCoordinates[1]}) {
					coordinatesXYRobotSanta = append(coordinatesXYRobotSanta, []int{lastCoordinates[0] - 1, lastCoordinates[1]})
				} else {
					coordinatesXYRobotSanta = append(coordinatesXYRobotSanta, []int{lastCoordinates[0] - 1, lastCoordinates[1]})
					housesRobotSantaVisited += 1
				}
			}
		} else {
			lastCoordinates := coordinatesXYRealSanta[len(coordinatesXYRealSanta)-1]
			switch dir {
			case "^":
				if contains(coordinatesXYRealSanta, []int{lastCoordinates[0], lastCoordinates[1] + 1}) {
					coordinatesXYRealSanta = append(coordinatesXYRealSanta, []int{lastCoordinates[0], lastCoordinates[1] + 1})
				} else {
					coordinatesXYRealSanta = append(coordinatesXYRealSanta, []int{lastCoordinates[0], lastCoordinates[1] + 1})
					housesRealSantaVisited += 1
				}
			case "v":
				if contains(coordinatesXYRealSanta, []int{lastCoordinates[0], lastCoordinates[1] - 1}) {
					coordinatesXYRealSanta = append(coordinatesXYRealSanta, []int{lastCoordinates[0], lastCoordinates[1] - 1})
				} else {
					coordinatesXYRealSanta = append(coordinatesXYRealSanta, []int{lastCoordinates[0], lastCoordinates[1] - 1})
					housesRealSantaVisited += 1
				}
			case ">":
				if contains(coordinatesXYRealSanta, []int{lastCoordinates[0] + 1, lastCoordinates[1]}) {
					coordinatesXYRealSanta = append(coordinatesXYRealSanta, []int{lastCoordinates[0] + 1, lastCoordinates[1]})
				} else {
					coordinatesXYRealSanta = append(coordinatesXYRealSanta, []int{lastCoordinates[0] + 1, lastCoordinates[1]})
					housesRealSantaVisited += 1
				}
			case "<":
				if contains(coordinatesXYRealSanta, []int{lastCoordinates[0] - 1, lastCoordinates[1]}) {
					coordinatesXYRealSanta = append(coordinatesXYRealSanta, []int{lastCoordinates[0] - 1, lastCoordinates[1]})
				} else {
					coordinatesXYRealSanta = append(coordinatesXYRealSanta, []int{lastCoordinates[0] - 1, lastCoordinates[1]})
					housesRealSantaVisited += 1
				}
			}
		}

	}

	return housesRobotSantaVisited + housesRealSantaVisited
}

func contains(s [][]int, e []int) bool {
	for _, a := range s {
		if slices.Equal(a, e) {
			return true
		}
	}
	return false
}
