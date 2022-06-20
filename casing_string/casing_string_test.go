package casing_string_test

import (
	"testing"

	"merge_lists/casing_string"
)

func TestCasingStringEasySolution(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{input: "how can mirrors be real if our eyes aren't real", want: "How Can Mirrors Be Real If Our Eyes Aren't Real"},
		{input: "How Can mirrors be real if Our eyes aren't real", want: "How Can Mirrors Be Real If Our Eyes Aren't Real"},
	}

	for _, tc := range testCases {
		t.Run("should return the string where the first letter of each word is capital", func(t *testing.T) {
			result := casing_string.CasingStringEasySolution(tc.input)
			if result != tc.want {
				t.Errorf("got %v want %v", result, tc.want)
			}
		})
	}
}

func TestCasingStringManual(t *testing.T) {
	testCases := []struct {
		input string
		want  string
	}{
		{input: "how can mirrors be real if our eyes aren't real", want: "How Can Mirrors Be Real If Our Eyes Aren't Real"},
		{input: "How Can mirrors be real if Our eyes aren't real", want: "How Can Mirrors Be Real If Our Eyes Aren't Real"},
		{input: "How Can mirrors be REAL if Our eyes aren't real", want: "How Can Mirrors Be Real If Our Eyes Aren't Real"},
	}

	for _, tc := range testCases {
		t.Run("should return the string where the first letter of each word is capital", func(t *testing.T) {
			result := casing_string.CasingStringManual(tc.input)
			if result != tc.want {
				t.Errorf("got %v want %v", result, tc.want)
			}
		})
	}
}
