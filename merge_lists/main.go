package main

import "fmt"

func main() {
	arr1 := []int{1, 4, 6}
	arr2 := []int{2, 3, 5}
	fmt.Print("Success ", merge(arr1, arr2))
}

func merge(arr1 []int, arr2 []int) []int {
	var mergedArr []int

	for i := 0; i < len(arr1)+len(arr2)+4; i++ {
		if arr1[0] > arr2[0] {
			mergedArr = append(mergedArr, arr2[0])
			arr2 = append(arr2[:0], arr2[1:]...)
		} else {
			mergedArr = append(mergedArr, arr1[0])
			arr1 = append(arr1[:0], arr1[1:]...)
		}
	}
	if len(arr1) == 0 {
		mergedArr = append(mergedArr, arr2[0])
	} else if len(arr2) == 0 {
		mergedArr = append(mergedArr, arr1[0])
	}

	return mergedArr
}
