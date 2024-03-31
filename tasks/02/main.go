package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}

	for _, v := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(v * 2)
		}()
		wg.Wait()
	}
}
