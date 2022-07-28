package main

import "fmt"

func main() {

	//series --> 1+2+3+...+10

	sum := 0

	for j := 1; j <= 10; j++ {
		sum = sum + j
	}

	fmt.Printf("series total = %v\n", sum)

	//series --> taking start and end number from user

	var startNumber, endNumber int
	series_total := 0

	fmt.Printf("Enter the first number of the series= ")
	fmt.Scan(&startNumber)

	fmt.Printf("Enter the last number of the series= ")
	fmt.Scan(&endNumber)

	for i := startNumber; i <= endNumber; i = i + 1 {
		series_total = series_total + i
	}

	fmt.Printf("Series Total is: %v\n", series_total)

}
