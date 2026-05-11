package main

import (
	"strings"
)

func GenerateArt(text string, banner map[rune][]string) string {
	var output strings.Builder
	word := SplitInput(text)
	for i, words := range word {
		if words == "" {
			if i != len(word)-1 {
				output.WriteString("\n")
			}
			continue
		}
		for row := 0; row < 8; row++ {
			for _, char := range words {
				output.WriteString(banner[char][row])

			}
			output.WriteString("\n")

		}
		if word[len(word)-1] == "" {
			output.WriteString("\n")
			output.WriteString("\n")
			output.WriteString("\n")
			output.WriteString("\n")
			output.WriteString("\n")
			output.WriteString("\n")
			output.WriteString("\n")
			output.WriteString("\n")
		}

	}
	return output.String()
}
