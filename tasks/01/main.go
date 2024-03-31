package main

import "fmt"

type Human struct {
	Name string
	Age string
}

type Action struct {
	Human
	Act string
}

func main() {
	act := Action{
		Human{
			Name: "Gigachad",
			Age: "25",
		},
		"Gigachading",
	}

	fmt.Println(act)
}