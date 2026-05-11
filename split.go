package main

import (
	"strings"
)

func SplitInput(input string) []string {
	text := strings.Split(input, "\\n")
	return text
}
