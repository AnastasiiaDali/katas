package merge

//refactored

func Merge(arr1 []int, arr2 []int) []int {
	var mergedArr []int
	index1 := 0
	index2 := 0

	for {
		if len(arr1) == index1 {
			mergedArr = append(mergedArr, arr2[index2:]...)
			break
		}
		if len(arr2) == index2 {
			mergedArr = append(mergedArr, arr1[index1:]...)
			break
		}
		if arr1[index1] > arr2[index2] {
			mergedArr = append(mergedArr, arr2[index2])
			index2++
		} else {
			mergedArr = append(mergedArr, arr1[index1])
			index1++
		}
	}
	return mergedArr
}

// not working

//func Merge(arr1 []int, arr2 []int) []int {
//	var mergedArr []int
//	numberOfIteration := len(arr1) + len(arr2) - 1
//
//	for i := 0; i < numberOfIteration; i++ {q
//		if arr1[0] > arr2[0] {
//			mergedArr = append(mergedArr, arr2[0])
//			arr2 = append(arr2[:0], arr2[1:]...)
//		} else {
//			mergedArr = append(mergedArr, arr1[0])
//			arr1 = append(arr1[:0], arr1[1:]...)
//		}
//	}
//	if len(arr1) == 0 {
//		mergedArr = append(mergedArr, arr2[0])
//	} else if len(arr2) == 0 {
//		mergedArr = append(mergedArr, arr1[0])
//	}
//
//	return mergedArr
//}
