package main

import "fmt"

func add(a int, b int) int {
	result := a + b
	return result
}

func sub(a int, b int) int {
	result := a - b
	return result
}

func main() {
	fmt.Printf("%v\n", add(10, 20))
	fmt.Printf("%v\n", sub(10, 20))
}
