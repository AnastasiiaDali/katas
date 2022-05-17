package main

import (
	"fmt"
	"math"
)

func main() {
	array := []int{4, 3, 9, 7, 2, 1}
	result := SquareRoot(array)

	fmt.Printf("Success %v\n", result)

}

func SquareRoot(array []int) []int {
	results := []int{}

	for _, num := range array {
		sqrt := math.Sqrt(float64(num))
		if math.Mod(sqrt, 1.0) == 0 {
			results = append(results, int(sqrt))
		} else {
			results = append(results, num*num)
		}
	}

	return results
}
