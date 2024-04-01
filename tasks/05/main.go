package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	time.Sleep(time.Second * 5)
}