package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Umumnya di Golang semua variable itu passed by value,
	// bukan by reference.
	// Dengan kata lain data yang dikirimkan adalah duplikasi
	// nilainya.
	address1 := Address{"Malang", "Jawa Timur", "Indonesia"}
	address2 := address1

	address2.City = "Surabaya"

	fmt.Println("Pass by value:")
	fmt.Println(address1)
	fmt.Println(address2)

	// Pointer -> kemampuan membuat reference ke lokasi data di
	// memory yang sama tanpa menduplikasi data yang sudah ada.
	// Pointer -> bisa membuat pass by reference.
	// Asterisk (*) digunakan untuk deklarasi
	var address3 Address = Address{"Banyuwangi", "Jawa Timur", "Indonesia"}
	var address4 *Address = &address3

	address4.City = "Jember"
	fmt.Println()
	fmt.Println("Pass by reference:")
	fmt.Println(address3)
	fmt.Println(*address4)

	// Memindahkan referensi awal ke referensi kedua, sehingga
	// nilai dari variabel rujukan pertama ikut berubah menjadi
	// variabel kedua menggunakan operator asterisk (*)
	*address4 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println()
	fmt.Println("Memindahkan kedua referensi kepada referensi dari nilai yang baru:")
	fmt.Println(address3)
	fmt.Println(address4)

	// Melakukan re-assigning value kepada nilai dari variabel
	// pointer
	address4 = &Address{"Yogyakarta", "D.I. Yogyakarta", "Indonesia"}

	fmt.Println()
	fmt.Println("Melakukan re-assigning nilai kepada variabel pointer:")
	fmt.Println(address3)
	fmt.Println(address4)

	// Function new -> membuat pointer hanya mengembalikan pointer
	// ke data kosong, artinya tidak ada data awal.
	alamat1 := new(Address)
	alamat2 := alamat1

	// Kedua variabel pointer di atas memiliki referensi yang sama
	alamat2.Country = "Indonesia"

	fmt.Println()
	fmt.Println("Function new:")
	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
