package main

import (
	"fmt"
	"sync"
)

type array struct {
	data []int
	sum int
	mu sync.Mutex
}

// Здесь используются те же приемы, что и во втором задании
func main() {
	arr := array{
		data: []int{2, 4, 6, 8, 10},
	}

	wg := sync.WaitGroup{}

	for i := range arr.data {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			arr.mu.Lock()
			arr.sum += arr.data[idx] * arr.data[idx]
			arr.mu.Unlock()
		}(i)

		wg.Wait()
	}

	fmt.Println(arr.sum)

}