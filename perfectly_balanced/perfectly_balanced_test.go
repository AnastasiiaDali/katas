package perfectly_balanced_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"merge_lists/perfectly_balanced"
)

func TestPerfectlyBalanced(t *testing.T) {
	testCases := []struct {
		input string
		want  bool
	}{
		{input: "{[{}{}]}[()]", want: true},
		{input: "{{}{}}", want: true},
		{input: "[]{}()", want: true},
		{input: "{()}[)", want: false},
		{input: "{(})", want: false},
	}

	for _, tc := range testCases {
		output := perfectly_balanced.PerfectlyBalanced(tc.input)
		assert.Equal(t, tc.want, output)
	}
}
