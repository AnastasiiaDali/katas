package postfix_evaluation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"merge_lists/postfix_evaluation"
)

func TestEvaluatePostfix(t *testing.T) {
	testCases := []struct {
		input  string
		output int
	}{
		{input: "82/", output: 4},
		{input: "138*+", output: 25},
		{input: "545*+5/", output: 5},
	}

	for _, tc := range testCases {
		result := postfix_evaluation.EvaluatePostfix(tc.input)
		assert.Equal(t, tc.output, result)
	}
}
