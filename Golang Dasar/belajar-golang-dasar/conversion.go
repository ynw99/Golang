package main

import "fmt"

func main() {
	// Konversi tipe data
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)

	// harus hati-hati ketika konversi kepada ukuran tipe data
	// yang lebih kecil
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai8)

	// Konversi string
	var name = "Yusuf"
	var e byte = name[0] // byte adalah alias int8
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}
