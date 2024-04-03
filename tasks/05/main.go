package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем канал для сообщений
	ch := make(chan int)

	// В этой горутине в бесконечном цикле закидываем сообщения в канал
	// Бесконечного цикла не боимся потому что программа завершается через N секунд
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	// А в этой горутине выводим сообщения из канала
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	// Ждем 5 секунд прежде чем горутина main() завершит свою работу
	time.Sleep(time.Second * 5)
}