package main

import "fmt"

func getFullName() (string, string, string) {
	return "Yusuf", "Nur", "Wahid"
}

func getFirstLast() (string, string, string) {
	return "Yusuf", "Nur", "Wahid"
}

func main() {
	firstName, midName, lastName := getFullName()
	fmt.Println(firstName, midName, lastName)

	// Mengabaikan salah satu/beberapa return
	first, _, last := getFirstLast()
	fmt.Println(first, last)
}
