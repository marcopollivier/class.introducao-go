package main

import "fmt"

func main() {

	var i = 1
	for i <= 3 {
		fmt.Print(i)
		i++
	}

	for j := 4; j < 7; j++ {
		fmt.Print(j)
	}

	var k = 7
	for {
		fmt.Print(k)
		k++
		if k > 9 {
			break
		}
	}

}
