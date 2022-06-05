package main

import "fmt"

func main() {
	string := "IX"
	romanConverter(string)
	fmt.Printf("Success: %v", romanConverter(string))
}

var arabic = map[uint8]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanConverter(roman string) int {
	result := 0
	for i := 0; i < len(roman); i++ {
		if i+1 < len(roman) && arabic[roman[i]] < arabic[roman[i+1]] {
			result += arabic[roman[i+1]] - arabic[roman[i]]
			i++
		} else {
			result += arabic[roman[i]]
		}
	}
	return result
}
