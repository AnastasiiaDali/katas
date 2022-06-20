package gloves_test

import (
	"testing"

	gloves "merge_lists/pairs_of_gloves"
)

func TestGlovesAssembler(t *testing.T) {

	testCases := []struct {
		input []string
		want  int
	}{
		{input: []string{"red", "green", "red", "blue", "blue"}, want: 2},
		{input: []string{"blue", "blue", "0", "green", "yellow", "red"}, want: 1},
		{input: []string{"blue", "blue", "blue", "blue", "blue", "blue"}, want: 3},
		{input: []string{"Yellow", "Green", "green", "yeLLow", "GreEn", "red"}, want: 2},
	}

	for _, tc := range testCases {
		t.Run("should return the number of pairs of gloves", func(t *testing.T) {
			result := gloves.GlovesAssembler(tc.input)
			if result != tc.want {
				t.Errorf("got %v want %v", result, tc.want)
			}
		})
	}

}
