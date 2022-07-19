package word_wrap

import "strings"

func WordWrap(text string, width int) string {
	var finalText string
	var characters []rune

	for _, char := range text {
		characters = append(characters, char)
	}

	for len(characters) >= 1 {
		value := strings.Contains(string(characters[:width]), "\n")
		if value {
			i := strings.Index(string(characters), "\n")
			finalText = finalText + string(characters[:i+1]) + "\r"
			characters = characters[i+1:]
		} else {
			finalText = finalText + string(characters[:width]) + "\r\n"
			characters = characters[width:]
		}

		if len(characters) < width {
			width = len(characters)
		}
	}

	return finalText
}
