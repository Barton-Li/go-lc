package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var T int
	fmt.Fscan(reader, &T)
	for i := 0; i < T; i++ {
		var m, n int
		fmt.Fscan(reader, &m, &n)
		s, _ := reader.ReadString('\n')
		tStr, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		tStr = strings.TrimSpace(tStr)
		if canTrans(s, tStr) {
			fmt.Fprintln(writer, "YES")
		} else {
			fmt.Fprintln(writer, "NO")
		}

	}
}
func canTrans(s string, t string) bool {
	n, m := len(s), len(t)
	if m > n {
		return false
	}
	for k := 0; k <= m; k++ {
		left := t[:k] == s[:k]
		right := t[k:] == reverse(s[k:k+(m-k)])
		if left && right {
			return true
		}
	}
	return false

}
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
