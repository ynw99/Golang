package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	// Nil (Null) -> nilai default untuk sebuah objek jika di
	// bahasa pemrograman lain.
	// Nil -> nilai default di beberapa tipe data seperti:
	// interface, function, map, slice, pointer, dan channel
	var person map[string]string = newMap("")

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}

	person = newMap("Ucup")

	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person)
	}
}
