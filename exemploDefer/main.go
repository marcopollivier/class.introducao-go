package main

import "fmt"

func main() {
	fmt.Println("Primeiro")
	
	defer fmt.Println("Imprima apenas quando terminar")
	
	fmt.Println("Segundo")
	fmt.Println("Terceiro")
	fmt.Println("Quarto")
	
}