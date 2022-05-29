package main

import "fmt"

func main() {
	num := 7
	Triangle(num)
}

func Triangle(rows int) {
	temp := 1

	for i := 0; i < rows; i++ {

		for j := 0; j <= i; j++ {

			if j == 0 || i == 0 {
				temp = 1
			} else {
				temp = temp * (i - j + 1) / j
			}

			fmt.Printf(" %v", temp)
		}
		fmt.Println("")
	}
}
