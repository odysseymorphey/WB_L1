package main

import (
	"fmt"
	"sync"
)

// Структура-счетчик
type Counter struct {
	cntr int
	mu sync.Mutex
}

// При вызове инкремента доступ к переменной будет блокироваться,
// пока вызвавшая функцию не завершит работу
func (c *Counter) Increment() {
	c.mu.Lock()
	c.cntr++
	c.mu.Unlock()
}

func main() {
	// Создаем структуру Counter
	c := &Counter{
		cntr: 0,
	}

	// Создаем группу для ожидания завершения работы всез горутин
	wg := &sync.WaitGroup{}

	// Вызывем метод Increment() структуры Counter
	// в отдельной горутине
	for i := 0; i < 22; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим счетчик
	fmt.Println(c.cntr)
}