package main

import "fmt"

func main() {

	//Variable Declaration

	var name, uni string
	var age int
	var CG float32

	// Variable initialization

	name = "Bappi"
	uni = "AIUB"
	age = 22
	CG = 3.3

	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("My University is", uni)
	fmt.Println("My CGPA is", CG)

	fmt.Println(" \n ")

	fmt.Println("My name is", name, "I am ", age, "years old. \n", "My University is", uni, "and my CGPA is", CG)

}
