/*
d - decimal integer
o - octal integer
O - octal integer with 0o prefix
b - binary integer
x - hexadecimal integer lowercase
X - hexadecimal integer uppercase
f - decimal floating point, lowercase
F - decimal floating point, uppercase
e - scientific notation (mantissa/exponent), lowercase
E - scientific notation (mantissa/exponent), uppercase
g - the shortest representation of %e or %f
G - the shortest representation of %E or %F
c - a character represented by the corresponding Unicode code point
q - a quoted character
U - Unicode escape sequence
t - the word true or false
s - a string
v - default format
#v - Go-syntax representation of the value
T - a Go-syntax representation of the type of the value
p - pointer address
% - a double %% prints a single %
*/

//Number conversion calculator

package main

import "fmt"

func main() {
	var decimalNumber int

	fmt.Printf("Decimal Number=")
	fmt.Scanf("%v", &decimalNumber)

	fmt.Printf("Binary Number= %b\n", decimalNumber)
	fmt.Printf("Octal Number= %o\n", decimalNumber)
	fmt.Printf("Hexadecimal Number= %x\n", decimalNumber)

	f_number := 3.141624588888888889999999999666666666333333
	fmt.Printf("Float Number= %f\n", f_number)
	fmt.Printf("Float Number= %.3f\n", f_number)

}
