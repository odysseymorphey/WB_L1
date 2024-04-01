package main

import (
	"fmt"
	"sync"
)

type ConcMap struct {
	m  map[string]int
	mu sync.RWMutex
}

func NewConcMap() *ConcMap {
	return &ConcMap{
		m: make(map[string]int),
	}
}

func (c *ConcMap) Load(key string) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.m[key]

	return val, ok
}

func (c *ConcMap) Store(key string, val int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.m[key] = val
}

type kv struct {
	key string
	val int
}

func main() {
	m := NewConcMap()
	sampleData := []kv{{"SPAS", 12}, {"MAG", 7}, {"Kalak", 12}, {"M", 16}}
	ch := make(chan kv)
	
	wg := &sync.WaitGroup{}
	
	go func() {
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for v := range ch {
					m.Store(v.key, v.val)
				}
			}()
		}
	}()

	for _, v := range sampleData {
		ch <- v
	}

	close(ch)
	wg.Wait()

	for k := range m.m {
		fmt.Print(k, " ")
		fmt.Println(m.Load(k))
	}
}
