package main

import "strings"

func PadArtRows(rows []string, width int) []string {
	result := make([]string, len(rows))
	for i, row := range rows {
		if len(row) < width {
			row += strings.Repeat(" ", width-len(row))
		}
		result[i] = row
	}
	return result
}
