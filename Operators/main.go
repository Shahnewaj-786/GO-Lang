/*
1.Arithmetic --> + - * / %
2.Assignment --> = += -= *= /= %=
3.Unary --> ++ --
4.Relational --> < > == != !< !>
5.Logical --> && || !
6.Bitwise --> & | ^ >> <<
7.Spical
*/

package main

import "fmt"

func main() {

	var A, B int

	fmt.Printf(" A = ")
	fmt.Scan(&A)

	fmt.Printf(" B = ")
	fmt.Scan(&B)

	result := A + B
	fmt.Printf("SUM = %v + %v = %v\n", A, B, result)

	result = A - B
	fmt.Printf("SUB = %v - %v = %v\n", A, B, result)

	result = A * B
	fmt.Printf("MULT = %v * %v =% v\n", A, B, result)

	result = A / B
	fmt.Printf("DIV = %v / %v = %v\n", A, B, result)

	result = A % B
	fmt.Printf("Reminder = %v  %%  %v = %v\n", A, B, result)

	//Assignment
	x := 4

	x += 5 // here x = x+5 represents by x += 5

	fmt.Printf("X = %v\n", x)

	//Unary
	u_x := 4

	u_x++ // here x = x+1 represents by x ++

	fmt.Printf("X = %v\n", u_x)

	//Relational
	r_x := 5
	r_y := 9

	fmt.Printf("x = %v\n", r_x == r_y)

	//Logical

	l_x := 16

	l_z := l_x > 5 && l_x > 9

	fmt.Printf("l_z = %v\n", l_z)

	//Bitwise

	b_x := 18 // 18 = 10010 (binary)
	b_y := 17 // 17 = 10001 (binary)
	b_z := b_x & b_y

	fmt.Printf("Bitwise AND= %v\n", b_z)

}
