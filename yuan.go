package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isVowel(ch byte) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	for _, vowel := range vowels {
		if ch == vowel {
			return true
		}
	}
	return false
}

func countVowels(s string, start, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		if isVowel(s[i]) {
			count++
		}
	}
	return count
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	s, k := "", 0
	fmt.Sscanf(line, "%s %d", &s, &k)

	maxVowels := countVowels(s, 0, k-1)

	for i := 1; i <= len(s)-k; i++ {
		vowels := countVowels(s, i, i+k-1)
		if vowels > maxVowels {
			maxVowels = vowels
		}
	}

	fmt.Println(maxVowels)
}
