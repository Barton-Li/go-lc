package main

import (
	"fmt"
)

func longOfsub(s string) int {
	indexMap := make(map[rune]int)
	maxLen := 0
	start := 0
	for end, char := range s {
		if index, exists := indexMap[char]; exists && index >= start {
			start = index + 1
		}
		indexMap[char] = end
		curLen := end - start + 1
		if curLen > maxLen {
			maxLen = curLen
		}

	}
	return maxLen
}
func main() {
	s := "abbbccbc"
	fmt.Println(longOfsub(s))
}
