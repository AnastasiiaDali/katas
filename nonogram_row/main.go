package main

import "fmt"

//expected 2,1,3
func main() {
	row := []int{0, 0, 0, 0, 1, 1, 0, 0, 1, 0, 1, 1, 1}
	fmt.Printf("Result %v", Nonogram(row))
}

func Nonogram(slice []int) []int {
	var result []int
	container := make(map[int]int)
	key := 0

	for _, number := range slice {
		if number == 1 {
			container[key] += 1
		} else {
			key++
		}
	}

	for _, i := range container {
		result = append(result, i)
	}
	return result
}
