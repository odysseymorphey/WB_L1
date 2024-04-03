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

// Берем значение из мапы
// Блокируем чтение для других горутини с помощью функции RLock()
func (c *ConcMap) Load(key string) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.m[key]

	return val, ok
}

// Записываем в мапу
// Блокируем запись для других горутин с помощью функции Lock()
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
	// Создаем конкурнтную мапу
	m := NewConcMap()

	// Подготавливаем набор данных
	sampleData := []kv{{"SPAS", 12}, {"MAG", 7}, {"Kalak", 12}, {"M", 16}}

	// Канал куда будем писать
	ch := make(chan kv)
	
	wg := &sync.WaitGroup{}
	

	// В отдельной горутине вызываем другие горутины,
	// которые слушают один канал и записывают и записывают
	// данные из него в мапу
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

	// Пишем в канал данные
	for _, v := range sampleData {
		ch <- v
	}

	// Закрываем канал и жедм пока все горутины отработают
	close(ch)
	wg.Wait()

	// Просто вывод
	for k := range m.m {
		fmt.Print(k, " ")
		fmt.Println(m.Load(k))
	}
}
