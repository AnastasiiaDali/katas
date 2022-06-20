package casing_string

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CasingStringEasySolution(input string) string {
	result := cases.Title(language.Und, cases.NoLower).String(input)
	return result
}

func CasingStringManual(input string) string {
	arrayOfStrings := strings.Split(input, " ")
	var newArray []string

	for _, str := range arrayOfStrings {
		lowerCaseStr := strings.ToLower(str)

		firstLetterRune := []rune(strings.ToUpper(lowerCaseStr[0:1]))

		out := []rune(lowerCaseStr)
		out[0] = firstLetterRune[0]

		str = string(out)

		newArray = append(newArray, str)
	}

	result := strings.Join(newArray, " ")

	return result
}
