package main

import (
	"fmt"
	"sync"
)

// Здесь используются те же приемы, что и во втором задании
func main() {
	arr := []int{2, 4, 6, 8, 10}
	sum := 0

	wg := sync.WaitGroup{}

	for _, v := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum += v * v
		}()
		wg.Wait()
	}

	fmt.Println(sum)
}
