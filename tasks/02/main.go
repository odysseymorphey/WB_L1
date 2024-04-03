package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}

	// Создаем вейм группу для ожидания завершения горутины
	wg := sync.WaitGroup{}

	for _, v := range arr {
		// Увеличиваем счетчик вейтгруппы на один перед запуском горутины
		wg.Add(1)
		go func() {
			// Уменьшаем счечик вейтгруппы на 1 после заверщения работы функции
			defer wg.Done()

			// Выводим результат 
			fmt.Println(v * 2)
		}()
		wg.Wait()
	}
}
