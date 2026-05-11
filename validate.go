package main

import (
	"fmt"
)

func ValidateInput(text string) (rune, error) {
	for _, char := range text {
		if char < ' ' || char > '~' {
			return char, fmt.Errorf("incorrect file : %v", char)
		}
	}
	return 0, nil
}
