package reverse_string_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"merge_lists/reverse_string"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{input: "reverse me", output: "em esrever"},
		{input: "rats live on no evil star", output: "rats live on no evil star"},
	}

	for _, tc := range testCases {
		t.Run("should return reversed string", func(t *testing.T) {
			result := reverse_string.ReverseString(tc.input)
			assert.Equal(t, tc.output, result)
		})
	}
}
