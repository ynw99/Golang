package main

import "fmt"

func main() {
	// Langsung wajib memberi nilai variabel const
	const (
		firstName = "Yusuf"
		lastName  = "Nur Wahid"
		value     = 100
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	// Const tidak akan dikomplain jika tidak dipakai
	const nickName = "Ucup"

}
