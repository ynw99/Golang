package main

import "fmt"

func main() {
	// Tipe data array tapi bisa berubah
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// Mengubah array/slice membuat keduanya juga berubah
	months[4] = "Diubah"
	fmt.Println(months)
	fmt.Println(slice1)
	slice1[0] = "Mei Lagi"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Ucup")
	fmt.Println(slice3)
	fmt.Println(months)

	// Create slice with make function
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Yusuf"
	newSlice[1] = "Nur"

	fmt.Println(newSlice, len(newSlice), cap(newSlice))

	// Copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// Deklarasi array vs slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
