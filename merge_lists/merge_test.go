package merge

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {

	t.Run("if we pass arrays with the length of 3 integers, should return sorted array", func(t *testing.T) {
		arr1 := []int{1, 4, 6}
		arr2 := []int{2, 3, 5}

		want := []int{1, 2, 3, 4, 5, 6}
		got := Merge(arr2, arr1)

		if reflect.DeepEqual(want, got) != true {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run(" if we pass arrays with the length of 5 integers, should return sorted array", func(t *testing.T) {
		arr1 := []int{1, 4, 6, 8}
		arr2 := []int{2, 3, 5, 7}

		want := []int{1, 2, 3, 4, 5, 6, 7, 8}
		got := Merge(arr2, arr1)

		if reflect.DeepEqual(want, got) != true {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
