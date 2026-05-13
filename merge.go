package main

func MergeBanners(base map[rune][]string, priority map[rune][]string) map[rune][]string {
	result := make(map[rune][]string)
	for i, value := range base {
		result[i] = value
	}
	for i, value := range priority {
		result[i] = value
	}
	return result
}
