package main

import "fmt"

func endApp() {
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
		Panic function -> function yang bisa digunakan untuk
		menghentikan program dan biasa dipanggil ketika terjadi
		error pada saat program berjalan.
		Saat panic function dipanggil program akan berhenti,
		namun defer func akan tetap dieksekusi.
	*/
	runApp(true)
}
