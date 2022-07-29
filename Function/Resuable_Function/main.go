package main

import "fmt"

func displayMessage(countryName string) {
	fmt.Printf("I love my motherland: %v\n", countryName) // Passing aurgument or parameter for resue function
}

func square(number int) {
	res := number * number
	fmt.Printf("%v\n", res)
}

func main() {
	displayMessage("Bangladesh.")
	displayMessage("Germany.")
	displayMessage("Canada.")
	square(5)
	square(6)
	square(4)
	square(55)
}
