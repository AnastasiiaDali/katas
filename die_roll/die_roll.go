package main

import (
	"fmt"
	"math/rand"
	"time"
)

const die = 100

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("Success %v\n", DieRoll())
}

func RandomNumber() int {
	min := 1
	max := die + 1
	return rand.Intn(max-min) + min
}

func DieRoll() [die]int {

	arr := [die]int{}

	for i := 0; i < 1000000; i++ {
		randomNumber := RandomNumber()
		arr[randomNumber-1] = arr[randomNumber-1] + 1
	}
	return arr
}
