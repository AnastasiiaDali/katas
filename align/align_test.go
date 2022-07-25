package align_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"merge_lists/align"
)

func TestAlign(t *testing.T) {

	testCases := []struct {
		str            string
		length         int
		alignment      string
		expectedOutput string
	}{
		{str: "iAmTheString", length: 20, alignment: "center", expectedOutput: "....iAmTheString...."},
		{str: "iAmTheString", length: 17, alignment: "center", expectedOutput: "...iAmTheString.."},
		{str: "iAmTheString", length: 15, alignment: "left", expectedOutput: "iAmTheString..."},
		{str: "iAmTheString", length: 15, alignment: "right", expectedOutput: "...iAmTheString"},
		{str: "iAmTheString", length: 0, alignment: "right", expectedOutput: "iAmTheString"},
		{str: "", length: 0, alignment: "right", expectedOutput: ""},
		{str: "iAmTheString", length: 10, alignment: "", expectedOutput: "iAmTheString"},
	}

	for _, tc := range testCases {
		t.Run("should return correct alignment", func(t *testing.T) {
			output := align.Align(tc.str, tc.length, tc.alignment)
			assert.Equal(t, tc.expectedOutput, output)
		})
	}
}
