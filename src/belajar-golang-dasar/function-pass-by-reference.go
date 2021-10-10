package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// Pass by reference untuk parameter fungsi
	address1 := Address{"Bandung", "Jawa Barat", ""}

	changeAddressToIndonesia(&address1)

	fmt.Println(address1)
}
