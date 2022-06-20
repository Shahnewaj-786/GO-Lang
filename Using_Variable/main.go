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

	// Variable Declaration and Variable initialization togather

	var name_1 = "Shahnewaj"
	var uni_1 = "AIUB"
	var age_1 = 24
	var CG_1 = 2.9

	fmt.Println("My name is", name_1, "I am ", age_1, "years old. \n", "My University is", uni_1, "and my CGPA is", CG_1)

	newvariable := "We can declear variable without var key word. We just have to use _:_ before _=_ like \" CG := 2.0 \""

	fmt.Print(newvariable)

}
