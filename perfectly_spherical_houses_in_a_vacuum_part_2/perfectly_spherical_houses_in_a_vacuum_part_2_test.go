package perfectly_spherical_houses_in_a_vacuum_part_2_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"merge_lists/perfectly_spherical_houses_in_a_vacuum_part_2"
)

func TestHouseCounter(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{input: "^v", want: 3},
		{input: "^>v<", want: 3},
		{input: "^v^v^v^v^v", want: 11},
	}

	for _, test := range testCases {
		t.Run("should return correct number of houses", func(t *testing.T) {
			got := perfectly_spherical_houses_in_a_vacuum_part_2.HouseCounter(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}
