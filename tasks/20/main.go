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
	sep := strings.Split(str, " ")

	for i, j := 0, len(sep)-1; i < j; i, j = i+1, j-1 {
		sep[i], sep[j] = sep[j], sep[i]
	}

	return strings.Join(sep, " ")
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
