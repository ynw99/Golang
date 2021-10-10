package main

import "fmt"

func main() {
	name := "Ucup"

	if name == "Yusuf" {
		fmt.Printf("Hello Yusuf")
	} else if name == "Nur" {
		fmt.Println("Hello Nur")
	} else {
		fmt.Println("Hi", name, ".")
	}

	// Short statement
	if length := len(name); length > 4 {
		fmt.Println("Terlalu panjang.")
	} else {
		fmt.Println("Nama sudah benar.")
	}
}
