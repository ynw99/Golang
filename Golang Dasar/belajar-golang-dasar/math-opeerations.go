package main

import "fmt"

func main() {
	var result = 0
	var a = 10
	var b = 10

	result = a + b
	fmt.Println(result)

	result = a * b
	fmt.Println(result)

	// Augmented assignments
	result += a
	fmt.Println(result)

	// Unary operator
	result++
	fmt.Println(result)

	var negatif = -1000
	var positif = +1000
	fmt.Println(negatif)
	fmt.Println(positif)
}
