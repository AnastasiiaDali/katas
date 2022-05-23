package main

import "fmt"

func main() {
	n := 15
	fmt.Printf("Success %v\n", Fibonacci(n))
}

func Fibonacci(n int) int {
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		prevB := b
		b = a + b
		a = prevB
	}
	return a
}
