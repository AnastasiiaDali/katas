package main

import "fmt"

func main() {
	Triangle()
}

func Triangle() {
	var rows int
	temp := 1
	fmt.Print("Please enter the number of rows ")
	fmt.Scan(&rows)

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
