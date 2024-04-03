package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func reverseString(str string) string {
	// Делим строку на отдельные слова
	sep := strings.Split(str, " ")

	// Меняем слова местами по индексу
	for i, j := 0, len(sep)-1; i < j; i, j = i+1, j-1 {
		sep[i], sep[j] = sep[j], sep[i]
	}

	// Возвращаем сконкотенированнуб строку 
	return strings.Join(sep, " ")
}

func main() {
	fmt.Println("Enter string")
	// Создаем сканнер чтобы считывать считывать строку целиком
	sc := bufio.NewScanner(os.Stdin)

	// Канал чтобы считывать посылаемые сигналы и завершать работу
	//  потому что мы имеем бесконечный цикл
    sig := make(chan os.Signal)
    signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

    go func() {
        <- sig

        os.Exit(0)
    }()

	// Считываем строку и отправляем ее функцю
	for sc.Scan() {
        str := sc.Text()
        fmt.Println(reverseString(str))
	}

}
