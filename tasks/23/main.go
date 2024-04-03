package main

import "fmt"


// Функция которая принимает на вход два параметра:
// слайс и позицию элемента который нужно удалить
func deleteElement(s []int, pos int) []int {
	// Если позиция больше чем длина слайса, то просто возвращаем неизмененный слайс
	if pos >= len(s) {
		return s
	}

	// Создается новый слайс на основе дочерних слайсов:
	// Если pos = 4, то
	// s[:pos] = [1, 2, 3, 4]
	// s[pos + 1:] = [6, 7]
	return append(s[:pos], s[pos + 1:]...)
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	s = deleteElement(s, 4)
	fmt.Println(s)
}