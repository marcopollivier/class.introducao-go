package main

import (
	"fmt"
	"time"
)

func main() {
	naturais := make(chan int)
	squares := make(chan int)

	go func() { //counter
		for x := 1; x < 20; x++ {
			naturais <- x
		}
	}()

	go func() { //squarer
		for {
			x := <-naturais
			squares <- x * x
		}
	}()

	go func() { //printer
		for {
			fmt.Println(<-squares)
		}
	}()

	spinner()
}

func spinner() {
	for {
		time.Sleep(100 * time.Millisecond)
	}
}
