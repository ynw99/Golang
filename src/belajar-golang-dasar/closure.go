package main

import "fmt"

func main() {
	// Closure -> kemampuan sebuah fungsi untuk berinteraksi
	// dengan data-data di sekitarnya dalam scope yang sama
	counter := 0
	name := "Yusuf"

	// Fungsi yang ada di bawahnya bisa mengakses data yang
	// lebih global maka perlu deklarasi ulang variabel yang
	// ada di atasnya
	increment := func() {
		name := "Ucup"
		fmt.Println("Increment func called")
		counter++
		fmt.Println("Counter di dalam anonym func =", counter)
		fmt.Println("Name di dalam anonym func =", name)
	}

	increment()
	increment()

	fmt.Println()
	fmt.Println("Counter di dalam main func =", counter)
	fmt.Println("Name di dalam main func =", name)
}
