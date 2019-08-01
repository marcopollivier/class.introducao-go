package main

import "fmt"

func main() {

	var x = 1

	fmt.Println(x)
	fmt.Println(&x)

	// p, do tipo *int, aponta para x
	var p = &x

	fmt.Println(*p) // *p nesse momento vale 1
	fmt.Println(p)  // p nesse momento é o endereço de x

	*p = 99
	fmt.Println(*p)
	fmt.Println(x)

}
