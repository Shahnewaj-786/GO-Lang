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
	fmt.Print("Go is great for multithread programming.")*/

	//Variables

	/*variable naming rules: 1. letter, digits, _
	                             2. keywords can not used as veriable
								 3. can not have space
								 4. digits can not the first letter in variable
								 5. Case sensetive


	*/
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
