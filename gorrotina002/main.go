package main

import (
	"fmt"
	"time"
)

func main() {
	go print(45)
	go print(10)
	go print(9)
	spinner(100 * time.Millisecond) // usando agora para travar a execução
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func print(n int) {
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fib(n))
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
