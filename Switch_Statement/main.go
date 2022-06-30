package main

import "fmt"

func main() {
	var digit int

	fmt.Printf("Enter a digit from 0-9: ")
	fmt.Scan(&digit)

	switch digit {
	case 0:
		fmt.Printf("Zero\n")
	case 1:
		fmt.Printf("One\n")
	case 2:
		fmt.Printf("Two\n")
	case 3:
		fmt.Printf("Three\n")
	case 4:
		fmt.Printf("Four\n")
	case 5:
		fmt.Printf("Five\n")
	case 6:
		fmt.Printf("Six\n")
	case 7:
		fmt.Printf("Seven\n")
	case 8:
		fmt.Printf("Eight\n")
	case 9:
		fmt.Printf("Nine\n")
	default:
		fmt.Printf("Not a valid digit. Enter between 0 to 9\n")
	}
}
