package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("Success %v\n", DieRoll())
}

func RandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 7
	return rand.Intn(max-min) + min
}

func DieRoll() []int {
	index1, index2, index3, index4, index5, index6 := 0, 0, 0, 0, 0, 0

	for i := 0; i < 1000000; i++ {
		randomNumber := RandomNumber()
		switch randomNumber {
		case 1:
			index1++
		case 2:
			index2++
		case 3:
			index3++
		case 4:
			index4++
		case 5:
			index5++
		default:
			index6++
		}
	}
	arr := []int{index1, index2, index3, index4, index5, index6}
	return arr
}
