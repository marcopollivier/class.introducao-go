package main

import "fmt"

func main() {

	var i = 1
	for i <= 3 {
		fmt.Print(i)
		i++
	}

	for j := 1; j < 9; j++ {
		fmt.Print(j)
	}

	var k = 1
	for {
		k++
		fmt.Print(k)
		if k == 9 {
			break
		}
	}

}
