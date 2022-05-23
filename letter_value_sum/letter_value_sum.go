package main

import "fmt"

func main() {
	str := "microspectrophotometries"
	fmt.Printf("Success %v\n", Calculator(str))
}

func Calculator(str string) int {
	var sum int
	sum = 0
	for i := 0; i < len(str); i++ {
		sum = sum + int(str[i]) - 96
	}
	return sum
}
