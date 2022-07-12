package cash_machine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBreakIntoChange(t *testing.T) {
	testsCases := []struct {
		name   string
		input  float64
		output map[string]int
	}{
		{name: " should return the minimum number of coins and notes representating 3.45", input: 3.45, output: map[string]int{"£2": 1, "£1": 1, "20p": 2, "5p": 1}},
		{name: " should return the minimum number of coins and notes representating 160", input: 160, output: map[string]int{"£50": 3, "£10": 1}},
		{name: " should return the minimum number of coins and notes representating 0.13", input: 0.13, output: map[string]int{"10p": 1, "2p": 1, "1p": 1}},
	}

	for _, tc := range testsCases {
		t.Run(tc.name, func(t *testing.T) {
			result := BreakIntoChange(tc.input)
			assert.Equal(t, tc.output, result)
		})
	}
}
