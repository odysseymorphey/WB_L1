package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ch_count := 0
	fmt.Println("Enter workers count")
	fmt.Scan(&ch_count)
	ch := make(chan string, 0)

	sc := bufio.NewScanner(os.Stdin)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < ch_count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				mu.Lock()
				fmt.Println("Worker ID:", i, "Msg:", v)
				mu.Unlock()
			}
		}()
	}

	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-s
		close(ch)
		os.Exit(0)
	}()

	for {
		fmt.Println("Enter message")
		sc.Scan()
		msg := sc.Text()
		ch <- msg
	}
}
