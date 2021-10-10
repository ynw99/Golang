package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpUcup NoKTP = "123456789"
	var marriedStatus Married = false

	fmt.Println(noKtpUcup)
	fmt.Println(marriedStatus)

}
