package main

import (
	"fmt"
	"reflect"
)

func varType(v interface{}) {
	fmt.Println(reflect.TypeOf(v).String())
}

func main() {
	var a int = 10
	var b float64 = 10.1
	var c string = "aboba"
	var d chan int
	var e bool = true

	varType(a)
	varType(b)
	varType(c)
	varType(d)
	varType(e)
}
