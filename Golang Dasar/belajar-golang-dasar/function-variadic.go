package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	total := sumAll(3, 3, 3, 4, 5, 4, 2)
	fmt.Println(total)

	// Slice sebagai parameter variadic function
	numbers := []int{10, 10, 10, 10, 10, 10}
	total = sumAll(numbers...)
	fmt.Println(total)
}
