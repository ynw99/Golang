package main

import "fmt"

func main() {
	name := "Yusuf"

	switch name {
	case "Yusuf":
		fmt.Println("Hello Yusuf.")
	case "Nur":
		fmt.Println("Hello Nur.")
	default:
		fmt.Println("Hello stranger.")
	}

	// Short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang.")
	case false:
		fmt.Println("Nama sudah benar.")
	}

	// Switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang.")
	case length > 5:
		fmt.Println("Nama agak panjang.")
	default:
		fmt.Println("Nama sudah benar.")
	}
}
