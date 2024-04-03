package main

import (
	"fmt"
	"sync"
)

// Создаем структуру с мьютексом чтобы контролировать доступ
// к данным среди горутин
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

			// Блокирум мьютекс, чтобы на время дальнейших оперций
			// у других горутин не было доступа к данным
			arr.mu.Lock()
			fmt.Println(arr.data[idx] * arr.data[idx])

			// После завершения операций разблокируем мьютекс
			arr.mu.Unlock()
		}(i)
		
		// Ждем завершения горутины
		wg.Wait()
	}

}