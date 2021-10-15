package main

import (
	"errors"
	"fmt"
)

// Error interface memiliki bentuk seperti di bawah, namun Golang
// memiliki error interface bawaan dari package error
// type error interface {
// 	Error() string
// }

func Pembagian(pembilang int, penyebut int) (int, error) {
	if penyebut == 0 {
		return 0, errors.New("Penyebut tidak boleh nol")
	} else {
		result := pembilang / penyebut
		return result, nil
	}
}

// The error built-in interface type is the conventional
// interface to represent an error condition, with the nil value
// representing no error
func main() {
	// Error interface -> digunakan sebagai kontrak untuk
	// membuat error, nama interface nya adalah error
	// var contohError error = errors.New("Ups Error")
	// fmt.Println(contohError.Error())
	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil:", hasil)
	} else {
		fmt.Println("Error", err)
	}
}
