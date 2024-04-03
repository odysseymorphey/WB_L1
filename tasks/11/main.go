package main

import "fmt"

func main() {
	set1 := []int{1, 4, 3, 2, 5}
	set2 := []int{3, 8, 5, 2}
	intersect := []int{}

	// Создаем мапу под первое множество для более быстрого обращения
	// к ее элементам
	set1Map := make(map[int]bool)
	
	// Итерируемся по одному множеству и добавляем его элементыы в мапу
	for _, v  := range set1 {
		set1Map[v] = true
    }
	
	// Итерируемся по другому можеству и проверяем наличие его елементов в мапе,
	// которую мы создали выше и добавляем в слайс элемент второго множества
    for _, v := range set2 {
        if set1Map[v] {
            intersect = append(intersect, v)
        }
    }

	// Выводим результат
	fmt.Println(intersect)
}