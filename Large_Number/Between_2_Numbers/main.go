package main

import "fmt"

func main() {
	var num1, num2 int

	fmt.Printf("Enter Two Numbers: ")
	fmt.Scan(&num1, &num2)

	if num1 > num2 {
		fmt.Printf("Largest Number is %v\n", num1)
	} else if num1 < num2 {
		fmt.Printf("Largest Number is %v\n", num2)
	} else {
		fmt.Printf("Both Numbers are Equal")
	}
}
