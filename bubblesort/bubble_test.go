package bubble_test

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"

	bubble "merge_lists/bubblesort"
)

func TestBubbleSort(t *testing.T) {
	testCases := []struct {
		input []int
		want  []int
	}{
		{input: []int{23, 11, 45, 8, 6, 12}, want: []int{6, 8, 11, 12, 23, 45}},
		{input: []int{1, 3, 6, 7, 100, 101}, want: []int{1, 3, 6, 7, 100, 101}},
		{input: []int{10, 9, 8, 7, 6, 5}, want: []int{5, 6, 7, 8, 9, 10}},
		{input: []int{-10, 22, -3, 4, 7, 3}, want: []int{-10, -3, 3, 4, 7, 22}},
	}

	for _, tc := range testCases {
		t.Run("should sorts the integers in ascending order and return sorted array of integers", func(t *testing.T) {
			result := bubble.BubbleSort(tc.input)
			if reflect.DeepEqual(result, tc.want) != true {
				t.Errorf("got %d want %d", result, tc.want)
			}
		})
	}

}

func BenchmarkBubbleSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	numbers := oneMillionIntegers()

	result := bubble.BubbleSort(numbers)
	fmt.Println(result)
}

func oneMillionIntegers() []int {
	var numbers = make([]int, 10)
	for i := 0; i <= 1000000; i++ {
		n := randomNum()
		numbers = append(numbers, n)
	}
	return numbers
}

func randomNum() int {
	randomNumber := rand.Intn(100)
	return randomNumber
}
