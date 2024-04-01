package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	wg := &sync.WaitGroup{}

	go func() {
		wg.Add(1)
		defer wg.Done()
		for data := range ch1 {
			ch2 <- data * data
		}
		close(ch2)
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		for data := range ch2 {
			fmt.Println(data)
		}
	}()

	for _, v := range nums {
		ch1 <- v
	}

	close(ch1)
	wg.Wait()
}
