package main

import "fmt"

func main() {
	// Var keyword with data type
	var name string

	name = "Yusuf Nur Wahid"
	fmt.Println(name)

	name = "Wahid Nur Yusuf"
	fmt.Println(name)

	// Var keyword without data type
	var friendName = "Ucup"

	fmt.Println(friendName)

	var age int8 = 23
	fmt.Println(age)

	// Without var keyword declaration
	nama := "Cakcup"
	fmt.Println(nama)

	country := "Indonesia"
	fmt.Println(country)

	// Multivariable declaration
	var (
		firstName = "Yusuf"
		lastName  = "Nur Wahid"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
