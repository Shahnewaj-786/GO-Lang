/* Package is the main thing for GO lang


Escape Sequence: \n --> newline; \t --> tab;

*/

/*

Data types and Keywords:

Datatypes: boolean, string, numeric - integer, floating, dreived types- pointer, array, slice, map, interface ect.


Variations of intwger --> int8 (8 bits), int16, int32, int64

Variations of float --> float32, float64


*/

package main

import "fmt"

//fmt means format

func main() { // func means function

	//int8 --> 2^8 = 256 which is -128 to 127
	//int16 -- 2^16
	//uint8 --> 2^8 = 256 which is 0 to 255

	/*fmt.Println("First program on GO")
	fmt.Print("Learning new things is so much fun.")
	fmt.Print("Go is great for multithread programming. \n")*/

	//Variables

	/*variable naming rules: 1. letter, digits, _
	                             2. keywords can not used as veriable
								 3. can not have space
								 4. digits can not the first letter in variable
								 5. Case sensetive


	*/
	/*
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

		newvariable := "We can declear variable without var key word. We just have to use _:_ before _=_ like \" CG := 2.0 \" \n"

		fmt.Print(newvariable)

		// Use of %v and Printf and Constant

		address := "Kalabagan"
		mobile := "0123333333"
		const HOME = "BD"

		fmt.Printf("My address is %v and my number is %v and my country is %v.", address, mobile, HOME)

		//Getting User input

		var fullname_ui string
		var num1_ui, num2_ui int

		fmt.Print("Enter your name: ")
		fmt.Scan(&fullname_ui)

		fmt.Print("Enter two numbers: ")
		fmt.Scan(&num1_ui, &num2_ui)

		fmt.Printf("My name is %v \n", fullname_ui)
		fmt.Printf("num1 is %v and num2 is %v \n", num1_ui, num2_ui)*/

	/*var float float32
	fmt.Print("Enter float number: ")
	fmt.Scanf("%v", &float)
	fmt.Printf("Float numver is %v", float)
	*/

	// Number formatting

	/* var decimalNumber int

	fmt.Printf("Decimal Number= ")
	fmt.Scanf("%v", &decimalNumber)

	fmt.Printf("Binary Number = %b\n", decimalNumber)
	fmt.Printf("Octal Number = %o\n", decimalNumber)
	fmt.Printf("Hexadecimal Number = %x\n", decimalNumber) */

	/* //String formatting

	var name = "Nafiz"

	fmt.Printf("%s\n", name)
	fmt.Printf("%q\n", name) */

	//Operators
	//Arithmetic
	/* var A, B int

	fmt.Printf(" A = ")
	fmt.Scan(&A)

	fmt.Printf(" B = ")
	fmt.Scan(&B)

	result := A + B
	//fmt.Printf(" Sum = %v ", result)
	fmt.Printf("Equn= %v+%v=%v\n", A, B, result) */

	/* //Assignment
	x := 4

	x += 5 // here x = x+5 represents by x += 5

	fmt.Printf("X = %v\n", x) */

	//Control Statement  --> Conditional CS, Loop CS

	//Conditional --> if, else if, else, switch

	//number := -590

	/* if number > 0 {
		fmt.Printf("Positive\n")
	}

	if number < 0 {
		fmt.Printf("Negative\n")
	}

	if number == 0 {
		fmt.Printf("Zero\n")
	} */

	/*if number > 0 {
		fmt.Printf("Positive\n")
	} else if number < 0 {
		fmt.Printf("Negative\n")
	} else {
		fmt.Printf("Zero\n")
	}*/

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
