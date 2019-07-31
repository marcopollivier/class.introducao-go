package main

import "fmt"

func main() {
	retorno := hello("Marco")

	fmt.Println(retorno)
}

func hello(nome string) string {
	return "Hello, " + nome
}
