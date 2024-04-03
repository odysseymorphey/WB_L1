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

	// Создаем небуферезированный канал куда будем посылать сообщения
	ch := make(chan string)

	// Создаем сканнер чтобы читывать строку целиком
	sc := bufio.NewScanner(os.Stdin)

	// Создаем вейт группу чтобы для ожидания завершения всез горутин
	wg := sync.WaitGroup{}
	
	// Создаем мьютекс чтобы из канала могла читать только одна горутина
	mu := sync.Mutex{}

	// В этом цикле все как и в предыдущих заданиях
	for i := 0; i < ch_count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				mu.Lock()

				// Выводим сообщение и айди воркера, что его обработал
				fmt.Println("Worker ID:", i, "Msg:", v)
				
				mu.Unlock()
			}
		}()
	}

	// Создаем канал типа сигнал для отлова системных сигналов
	s := make(chan os.Signal)
	signal.Notify(s, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	// В этой горуте ловим нужный сигнал и завершаем работу
	go func() {
		<-s
		close(ch)
		os.Exit(0)
	}()

	// В этом цикле передаем данные из stdout в канал для сообщений
	for sc.Scan() {
		fmt.Println("Enter message")
		msg := sc.Text()
		ch <- msg
	}
}
