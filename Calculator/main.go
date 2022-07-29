package main

import "fmt"

func add(a, b float32) float32 {

	return a + b
}

func sub(a, b float32) float32 {

	return a - b
}

func mul(a, b float32) float32 {

	return a * b
}

func div(a, b float32) float32 {

	return a / b
}

func main() {
	var num1, num2, result float32
	var option string

	i := true

	for i == true {
		fmt.Printf("Number 1: ")
		fmt.Scan((&num1))

		fmt.Printf("Number 2: ")
		fmt.Scan((&num2))

		fmt.Printf("Choose an option (+ - * /): ")
		fmt.Scan((&option))

		switch option {
		case "+":
			result = add(num1, num2)

		case "-":
			result = sub(num1, num2)

		case "*":
			result = mul(num1, num2)

		case "/":
			result = div(num1, num2)

		default:
			fmt.Printf("Invalid Option\n")
			continue

		}

		fmt.Printf("Result = %v\n", result)
	}
}
