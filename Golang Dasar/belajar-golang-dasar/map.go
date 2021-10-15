package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{
		"name":    "Yusuf",
		"address": "Banyuwangi",
	}

	fmt.Println(person)
	fmt.Println(person["name"])

	person["job"] = "Programmer"
	fmt.Println(person["job"])

	// Create map with make function
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Ucup"
	book["ups"] = "Salah"

	fmt.Println(book)
	// Delete value in map
	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}
