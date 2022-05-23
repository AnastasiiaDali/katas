package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("Success %v\n", DieRoll())
}

type indexes struct {
	i1 int
	i2 int
	i3 int
	i4 int
	i5 int
	i6 int
}

func RandomNumber() int {
	min := 1
	max := 7
	return rand.Intn(max-min) + min
}

func DieRoll() indexes {
	var newIndexes indexes
	index1 := 0
	index2 := 0
	index3 := 0
	index4 := 0
	index5 := 0
	index6 := 0
	for i := 0; i < 1000000; i++ {
		randomNumber := RandomNumber()
		switch randomNumber {
		case 1:
			index1++
			newIndexes.i1 = index1
		case 2:
			index2++
			newIndexes.i2 = index2
		case 3:
			index3++
			newIndexes.i3 = index3
		case 4:
			index4++
			newIndexes.i4 = index4
		case 5:
			index5++
			newIndexes.i5 = index5
		default:
			index6++
			newIndexes.i6 = index6
		}
	}
	return newIndexes
}

// {167026 167143 166459 166068 166559 166745}
