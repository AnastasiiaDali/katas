package main

import "fmt"

func main() {
	fmt.Print("Success ", divider(100))
}

//recursive func
func divider(number int) int {
	fmt.Printf("result %v\n", number)

	if number == 1 {
		return number
	}

	var result int

	if number%3 != 0 {
		if (number-1)%3 == 0 {
			result = number - 1
		} else {
			result = number + 1
		}
	} else {
		result = number / 3
	}

	return divider(result)
}

// func return => func return => func return => number
