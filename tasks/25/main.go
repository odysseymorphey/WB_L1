package main

import (
	"time"
)

func sleep(seconds int) {
	// Сохраняем время старта функции
	startTime := time.Now()

	for {
		// На каждой итерации фиксируем текущее время
		currentTime := time.Now()

		// Отнимаем от текущего времени время старта и сравниваем разницу с параметром функции
		if currentTime.Second()-startTime.Second() >= seconds {
			break
		}
	}
}

func main() {
	sleep(5)
}
