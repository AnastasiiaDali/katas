package main

import "fmt"

//func that accepts k and array
// 123456 => 345612

func main() {
	arrayOfNum := []int{1, 2, 3, 4, 5, 6}
	k := 2
	fmt.Printf("Success %v", rotate(arrayOfNum, k))
}

func rotate(arrayOfNum []int, k int) []int {
	for i := 0; i < k; i++ {
		arrayOfNum = append(arrayOfNum, arrayOfNum[0])
		arrayOfNum = append(arrayOfNum[:0], arrayOfNum[1:]...)
	}
	return arrayOfNum
}
