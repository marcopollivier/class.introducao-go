package main

import "fmt"

var mensagens = make(chan string)

func main() {

	go f()

	msg := <-mensagens
	fmt.Println(msg)
	fmt.Println(msg)

}

func f() {
	mensagens <- "ping"
	mensagens <- "pong"
}
