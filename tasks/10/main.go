package main

import "fmt"

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Для группировки по множеству исользуем мапу в которой ключ это множество,
	// а слайс из float64, которое хранит в себе элементы множества, это значение
	groups := make(map[int][]float64)


	for _, v := range data {
		// Количество десятков помноженое на 10 даст нам множество с шагом в 10
		set := (int(v) / 10) * 10
		groups[set] = append(groups[set], v)
	}

	fmt.Println(groups)
}
