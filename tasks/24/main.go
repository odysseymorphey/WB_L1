package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// Конструктор. Принимает на вход координаты x и y
// и возвращает структуру с заданными параметрами
func NewPoint(x float64, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

// Возвращает результат вычисения расстояния по формуле
func (p *Point) CalcDistance(f Point, s Point) float64 {
	return math.Sqrt(((f.x - f.y) * (f.x - f.y)) - ((s.x - s.y) * (s.x - s.y)))
}

func main() {
	a := NewPoint(3.1, 5.2)
	b := NewPoint(6.4, 7.9)

	fmt.Printf("%.4f\n", a.CalcDistance(a, b))
}
