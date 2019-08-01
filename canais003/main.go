package main

import "fmt"

func main() {

	mensagens := make(chan string, 1)

	mensagens <- "buferizado"
	mensagens <- "canal"

	fmt.Println(<-mensagens)
	fmt.Println(<-mensagens)
}
