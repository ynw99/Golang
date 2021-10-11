package main

import (
	"belajar-golang-dasar/database"
	"fmt"
)

func main() {
	// Ketika kita membuat sebuah package kita bisa membuat
	// sebuah fungsi yang akan dipanggil ketika package kita
	// diakses.
	result := database.GetDatabase()
	fmt.Println(result)
}
