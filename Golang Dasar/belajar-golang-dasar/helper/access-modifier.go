package helper

import "fmt"

// Access modifier di Golang hanya menggunakan penamaan awal
// huruf kapital dan huruf kecil.
// Huruf kapital -> bisa diakses dari package lain
// Huruf kecil -> tidak bisa diakses dari package lain

var version = 1
var Application = "Belajar Golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayHello(name string) {
	fmt.Println("Goodbye", name)
}
