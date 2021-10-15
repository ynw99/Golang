package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
}

func main() {
	// Struct method pada umumnya diakses menggunakan pass by
	// value.
	// Direkomendasikan menggunakan pointer agar hemat memory.
	yusuf := Man{"Yusuf"}

	yusuf.Married()

	fmt.Println("After married ->", yusuf.Name)
}
