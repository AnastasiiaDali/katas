package bubble

func BubbleSort(numbers []int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		for n := 0; n < len(numbers)-1-i; n++ {
			if numbers[n] > numbers[n+1] {
				numbers[n], numbers[n+1] = numbers[n+1], numbers[n]
			}
		}
	}
	return numbers
}
