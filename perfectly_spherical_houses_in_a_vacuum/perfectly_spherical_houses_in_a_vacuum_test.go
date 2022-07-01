package perfectly_spherical_houses_in_a_vacuum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHouseCounter(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{input: "^v^v^v^v^v", want: 2},
		{input: "^>v<", want: 4},
		{input: ">", want: 2},
	}

	for _, test := range testCases {
		t.Run("should return correct number of houses", func(t *testing.T) {
			got := HouseCounter(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
