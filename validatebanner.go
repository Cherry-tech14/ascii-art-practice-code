package main

import (
	"fmt"
)

func ValidateBanner(banner map[rune][]string) error {
	if len(banner) != 95 {
		return fmt.Errorf("invalid file")
	}
	for r := rune(32); r <= rune(126); r++ {
		if len(banner[r]) != 8 {
			return fmt.Errorf("invalid character %c got %v expected 8", rune(r), len(banner[r]))
		}

	}
	return nil
}
