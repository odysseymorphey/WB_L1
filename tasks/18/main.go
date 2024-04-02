package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	cntr int
	mu sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.cntr++
	c.mu.Unlock()
}

func main() {
	c := &Counter{
		cntr: 0,
	}

	wg := &sync.WaitGroup{}

	for i := 0; i < 22; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	wg.Wait()

	fmt.Println(c.cntr)
}