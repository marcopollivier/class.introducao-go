package main

import "fmt"

func main() {

	var x = 1
	fmt.Println(x)
	fmt.Println(&x)

	var p = &x
	fmt.Println(*p)
	fmt.Println(p)

	var c = x
	//fmt.Println(*c) nem deixa isso acontecer
	fmt.Println(c)
	fmt.Println(&c)

	c = 99
	fmt.Println(c)
	fmt.Println(x)

}
