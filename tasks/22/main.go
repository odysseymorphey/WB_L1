package main

import (
	"fmt"
	"math/big"
)

func main() {
	//  Используем тип big.Int для хранения больших чисел
	a := new(big.Int)
	b := new(big.Int)
	res := new(big.Int)

	// Задаем значения переменыыми. Значения больше чем 2^20
	a.SetString("10000000", 10)
	b.SetString("5000000", 10)

	// Сложение
	fmt.Println(res.Add(a, b))

	// Вычитание
	fmt.Println(res.Sub(a, b))
	
	// Умножение
	fmt.Println(res.Mul(a, b))
	
	// Деление
	fmt.Println(res.Div(a, b))
}
