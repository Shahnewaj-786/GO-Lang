package main

import "fmt"

func main() { // GO lang has only one Loop Control Statement which is for loop

	for x := 1; x < 6; x++ {
		fmt.Println("Bangladesh")
	}

	for y := 1; y <= 100; y++ {
		fmt.Printf("%v\n", y)
	}

	z := 100

	for z >= 1 {
		fmt.Printf("%v\n", z)
		z--
	}
}
