package main

import "fmt"

func main() {

	exemplo1()

}

func exemplo1() {
	var x int
	fmt.Println(x)
	fmt.Println(&x)

	x = 1

	fmt.Println(x)
	fmt.Println(&x)
}

func exemplo2() {

	fmt.Println("Coé")
}
