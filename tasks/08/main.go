package main

import "fmt"

func main() {
	var pos int
	var obj int64 = 4
	var bit int64
	var mask int64

	fmt.Println("Enter bit position")
	fmt.Scan(&pos)
	fmt.Println("Enter value: 0 or 1")
	fmt.Scan(&bit)

	// Независимо от значения бита, сначала готовим маску
	// сдвигая еденицу побитово, затем через логическое И
	// применяем маску число
	if bit == 0 {
		mask = ^(1 << pos)
		obj = obj & mask
	} else if bit == 1 {
		mask = (1 << pos)
		obj = obj | mask
	}

	fmt.Println(obj)
}