package main

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var n int
		fmt.Scan(&n)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
		}

		count := make([]int, n+1)
		left, validIntervals := 0, 0
		uniqueCount := 0

		for right := 0; right < n; right++ {
			if count[a[right]] == 0 {
				uniqueCount++
			}
			count[a[right]]++

			for uniqueCount == (right - left + 1) {
				validIntervals++

				count[a[left]]--
				if count[a[left]] == 0 {
					uniqueCount--
				}
				left++
			}
		}

		fmt.Println(validIntervals)
	}
}
