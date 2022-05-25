package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := 93
	fmt.Printf("Success %v\n", Fibonacci(n))
}

func Fibonacci(n int) string {
	a := big.NewInt(0)
	b := big.NewInt(1)

	for i := 0; i < n; i++ {
		a, b = b, b.Add(a, b)
	}

	return a.String()
}
