package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	count := 0
	maxCount := 50
	current := 1
	ch := make(chan int, 1)
	printNum := func(num int) {
		defer wg.Done()
		for count < maxCount {
			<-ch
			if num == current {
				fmt.Print(num)
				count++
				current = (current % 4) + 1
				ch <- current
			} else {
				ch <- num
			}
		}
	}
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go printNum(i)
	}
	ch <- 1
	wg.Wait()
}
