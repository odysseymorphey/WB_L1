package main

import "fmt"

func main() {
	// Последовательность
	seq := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем мапу в которой ключом будет элемент последовательности,
	// а значением количетсво этого элемента в последовательности
	set := make(map[string]int)

	// Итерируемся по последовательности добавляя в мапу каждый элемент
	for _, str := range seq {
		set[str] += 1
	}

	// Выводим результат
	fmt.Println(set)
}
