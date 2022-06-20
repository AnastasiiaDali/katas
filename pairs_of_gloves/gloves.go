package gloves

import (
	"math"
	"strings"
)

func GlovesAssembler(gloves []string) int {
	var theNumberOfPairs int
	temp := make(map[string]int)

	for _, color := range gloves {
		colorOfTheGlove := strings.ToLower(color)
		if _, ok := temp[colorOfTheGlove]; ok {
			temp[colorOfTheGlove] = temp[colorOfTheGlove] + 1
		} else {
			temp[colorOfTheGlove] = 1
		}
	}

	for _, v := range temp {
		value := v / 2
		floatN := float64(value)
		if v >= 2 {
			theNumberOfPairs = theNumberOfPairs + int(math.Floor(floatN))
		}
	}
	return theNumberOfPairs
}
