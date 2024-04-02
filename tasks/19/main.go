package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func reverseString(str string) string {
	runes := []rune(str)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	fmt.Println("Enter string")
	sc := bufio.NewScanner(os.Stdin)

    sig := make(chan os.Signal)
    signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

    go func() {
        <- sig

        os.Exit(0)
    }()

	for sc.Scan() {
        str := sc.Text()
        fmt.Println(reverseString(str))
	}

}
