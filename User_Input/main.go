package main

import "fmt"

func main() {
	var name string
	var num1, num2 int

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter two numbers: ")
	fmt.Scan(&num1, &num2)

	fmt.Printf("My name is %v \n", name)
	fmt.Printf("Number1 is %v and Number2 is %v \n", num1, num2)

}
