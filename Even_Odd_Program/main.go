package main

import "fmt"

func main() {

	var number int

	fmt.Printf("Enter an integer number: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Printf("Even Number")
	} else {
		fmt.Printf("Odd Number")
	}
}
