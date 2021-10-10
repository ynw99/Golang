package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	// Struct -> template data yang digunakan untuk menggabungkan
	// satu atau lebih tipe data lainnya di dalam satu kesatuan.
	// Struct -> kumpulan dari field
	var ucup Customer
	ucup.Name = "Yusuf Nur Wahid"
	ucup.Address = "Indonesia"
	ucup.Age = 23

	fmt.Println(ucup)
	fmt.Println(ucup.Name)
	fmt.Println(ucup.Address)
	fmt.Println(ucup.Age)

	// Struct literals
	// juga bisa -> var joko Customer = Customer{name, address,
	// age}
	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(joko)

	// Untuk cara ini harus urut dan sesuai jumlah fieldnya
	budi := Customer{"Budi", "Indonesia", 33}
	fmt.Println(budi)
}
