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

	//Print Even Numbers
	for even := 2; even <= 100; even = even + 2 {
		fmt.Printf("%v\n", even)
	}
	fmt.Printf("\n")
	//Print Odd Numbers
	for odd := 1; odd <= 100; odd = odd + 2 {
		fmt.Printf("%v\n", odd)
	}

	//Print Even Numbers
	for x := 1; x <= 100; x++ {
		if x%2 == 0 {
			fmt.Printf("%v\n", x)
		}
	}
	fmt.Printf("\n")
	//Print Odd Numbers
	for y := 1; y <= 100; y++ {
		if y%2 != 0 {
			fmt.Printf("%v\n", y)
		}
	}
}
