package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%v\n", i)
	}

	fmt.Printf("\n")

	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			break
		}
		fmt.Printf("%v\n", j)
	}

}
