package main

import "fmt"

func add(a int, b int) {
	result := a + b
	fmt.Printf("%v\n", result)
}

func sub(a int, b int) {
	result := a - b
	fmt.Printf("%v\n", result)
}

func main() {
	add(10, 20)
	sub(50, 13)
}
