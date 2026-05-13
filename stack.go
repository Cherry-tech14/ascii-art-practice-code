package main

func StackTwo(top, bottom []string) []string {
	result := []string{}
	for _, row := range top {
		result = append(result, row)
	}
	for _, row := range bottom {
		result = append(result, row)
	}
	return result
}

func StackAll(blocks [][]string) []string {
	result := []string{}
	for _, block := range blocks {
		for _, row := range block {
			result = append(result, row)
		}
	}
	return result
}
