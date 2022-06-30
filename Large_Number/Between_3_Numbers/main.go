package main

import "fmt"

func main() {
	var num1, num2, num3 int

	fmt.Printf("Enter Three Numbers: ")
	fmt.Scan(&num1, &num2, &num3)

	if num1 > num2 {
		if num1 > num3 {
			fmt.Printf("%v is the largest number", num1)
		} else {
			fmt.Printf("%v is the largest number", num3)
		}
	} else if num2 > num1 {
		if num2 > num3 {
			fmt.Printf("%v is the largest number", num2)
		} else {
			fmt.Printf("%v is the largest number", num3)
		}
	} else {
		fmt.Printf("All Numbers are Equal")
	}
}
