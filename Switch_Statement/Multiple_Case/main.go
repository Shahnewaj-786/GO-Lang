package main

import "fmt"

func main() {
	var classNumber int

	fmt.Printf("Enter study level: ")
	fmt.Scan(&classNumber)

	switch classNumber {
	case 1, 2, 3, 4, 5:
		fmt.Printf("PSC\n")
	case 6, 7, 8, 9, 10:
		fmt.Printf("SSC\n")
	case 11, 12:
		fmt.Printf("HSC\n")
	default:
		fmt.Printf("Not in Exam group")
	}
}
