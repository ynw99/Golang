package main

import "fmt"

func endApp() {
	// Recover disimpan di dalam defer function
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("defer: Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR!")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	/*
		Recover -> function yang bisa digunakan untuk menangkap
		data panic.
		Dengan recover proses panic akan terhenti, sehingga
		program akan tetap berjalan.
	*/
	runApp(false)
	fmt.Println("Ucup")
}
