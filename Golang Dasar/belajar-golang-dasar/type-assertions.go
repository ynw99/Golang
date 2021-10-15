package main

import "fmt"

func random() interface{} {
	return 100
}

func main() {
	// Type assertions -> kemampuan untuk mengubah sebuah tipe
	// data menjadi tipe data lain.
	// Sering digunakan ketika bertemu dengan data interface
	// kosong
	// var result interface{} = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// Agar lebih aman menggunakan switch expression
	var hasil interface{} = random()
	switch value := hasil.(type) {
	case string:
		fmt.Println("Value", value, "is string.")
	case int:
		fmt.Println("Value", value, "is integer.")
	default:
		fmt.Println("Unknown type")
	}
}
