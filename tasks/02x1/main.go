package main

import (
	"fmt"
	"sync"
)

type array struct {
	data []int
	mu sync.Mutex
}

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
			fmt.Println(arr.data[idx] * arr.data[idx])
			arr.mu.Unlock()
		}(i)

		wg.Wait()
	}

}