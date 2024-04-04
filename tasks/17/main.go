package main

import (
	"fmt"
)

func binarySearch(data []int, target int) (int, bool) {
	// Сохраняем минимальный индекс
	low := 0
	
	// Сохраняем максимальный индекс
	high := len(data) - 1


	for low <= high {
		// Вычисляем индекс центрального элемента
		mid := (low + high) / 2

		// Если таргет равен центральному элементу, возвращаем индекс центрального элемента и индикатор true
		// В противном случае, в зависимости от условия, уменьшаем максимальный
		// или увеличиваем минимальный индекс
		if target == data[mid] {
			return mid, true
		} else if target < data[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	// Если ничего не найдено возвращаем 0 и индикатор false
	return 0, false
}

func main() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	i, ok := binarySearch(data, 4)
	if !ok {
		fmt.Println("Sorry")
	} else {
		fmt.Println(i)
	}
}
