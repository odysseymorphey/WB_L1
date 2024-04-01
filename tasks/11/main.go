package main

import "fmt"

func main() {
	set1 := []int{1, 4, 3, 2, 5}
	set2 := []int{3, 8, 5, 2}
	set1Map := make(map[int]bool)
	intersect := []int{}
	
	for _, val := range set1 {
		set1Map[val] = true
    }
	
    for _, val := range set2 {
        if set1Map[val] {
            intersect = append(intersect, val)
        }
    }

	fmt.Println(intersect)
}