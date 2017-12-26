package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	ROUNDS = 1000 * 1000 * 10
)

func echo1(args []string) {
	s := ""
	for _, arg := range args[0:] {
		s += arg + " "
	}
}

func echo2(args []string) {
	strings.Join(args[0:], "")
}
func main() {
	start := time.Now()
	i := ROUNDS
	for i > 0 {
		echo1(os.Args)
		i--
	}
	fmt.Printf("echo1 %.2fs elapsed\n", time.Since(start).Seconds())
	i = ROUNDS
	start = time.Now()
	for i > 0 {
		echo2(os.Args)
		i--
	}
	fmt.Printf("echo2 %.2fs elapsed\n", time.Since(start).Seconds())
}
