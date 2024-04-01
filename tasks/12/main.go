package main

import "fmt"

func main() {
	seq := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]int)

	for _, str := range seq {
		set[str] += 1
	}

	fmt.Println(set)
}
