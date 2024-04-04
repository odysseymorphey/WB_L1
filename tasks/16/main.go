package main

import (
	"fmt"
)

// Сама функция сортировки. Сравниваем по условию элементы по левому и правому индексу и меняем их местами
func partition(arr []int, low, high int) ([]int, int) {
	// Создаем опорный элемент (самый правый индекс)
	pivot := arr[high]

	// Сохраняем минимальный индекс (самый левый)
	// По этому индексу в дальнейшем будет делиться массив
	i := low

	// Итерируемся по слайсу и если текущий элемент ме	ньше опорного, то меняем их местами
	// и увеличиваем индекс i
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
		// fmt.Println(arr)
	}
	
	// Меняем самый правый индекс местами с i-тым индексов
	arr[i], arr[high] = arr[high], arr[i]
	// fmt.Println()

	// Возвращаем отсортированный слайс и индекс по которому
	// будем делить слайс
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		// Сортируем
		arr, p = partition(arr, low, high)
		
		// Условно делим массив на два подмассива
		// и рекурсивно вызываем функцию quickSort
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func startSort(data []int) []int {
	return quickSort(data, 0, len(data)-1) 
}

func main() {
	data := []int{2, 56, 3, 4, 1, 3, 33, 23, 42, 5}

	fmt.Println(startSort(data))
}
