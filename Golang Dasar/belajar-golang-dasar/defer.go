package main

import "fmt"

func logging() {
	fmt.Println("Func defer telah dipanggil")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result:", result)
}

func main() {
	// Defer -> fungsi yang bisa dijadwalkan untuk dieksekusi
	// setelah sebuah fungsi selesai dieksekusi
	runApplication(10)
}
