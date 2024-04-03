package main

import (
	"fmt"
	"strings"
)

func checkUnique(m map[rune]int) bool {
	// Итерируемся по мапе и ищем ключи значения которых больше 1
	for k := range m {
		if m[k] > 1 {
			return false
		}
	}

	return true
}

func main() {
	str := ""
	fmt.Println("Enter string")
	fmt.Scan(&str)

	// Так как нет зависимости от регистра, 
	// можно перевести все символы строки в нижний региср
	str = strings.ToLower(str)

	// Создаем мапу которая будет хранить символы строки в качестве ключа
	// и количество этих символов в качестве значения
	m := make(map[rune]int)
	for _, v := range str {
		m[v] += 1
	}

	fmt.Println(checkUnique(m))

}