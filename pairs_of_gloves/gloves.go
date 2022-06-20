package gloves

import (
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
		if v >= 2 {
			theNumberOfPairs = theNumberOfPairs + value
		}
	}
	return theNumberOfPairs
}
