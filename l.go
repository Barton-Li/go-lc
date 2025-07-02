package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	oChan := make(chan struct{}, 1)
	eChan := make(chan struct{}, 1)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			<-oChan
			fmt.Println(i)
			eChan <- struct{}{}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			<-eChan
			fmt.Println(i)
			oChan <- struct{}{}
		}
	}()
	oChan <- struct{}{}
	wg.Wait()
}
