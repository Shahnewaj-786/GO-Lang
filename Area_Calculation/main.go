package main

import "fmt"

func main() {
	var base, height, radius float32

	fmt.Printf("Base = ")
	fmt.Scan(&base)

	fmt.Printf("Height = ")
	fmt.Scan(&height)

	area := 0.5 * base * height

	fmt.Printf("Area of Triangel = %v\n", area)

	fmt.Printf("Radius = ")
	fmt.Scan(&radius)

	r_area := 3.1416 * radius * radius

	fmt.Printf("Area of Circule = %v\n", r_area)

}
