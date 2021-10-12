package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type ContohLagi struct {
	Name        string `required:"true"`
	Description string `required:"true"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			return reflect.ValueOf(data).Field(i).Interface() != ""
		}
	}
	return true
}

func main() {
	// Di dalam bahasa pemrograman biasanya ada fitur Reflection.
	// Sebuah fitur yang memungkinkan kita untuk bisa melihat
	// struktur kode kita saat aplikasi sedang berjalan.
	// Hal ini bisa dilakukan di Golang menggunakan package reflect.
	sample := Sample{"Yusuf"}
	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println("Jumlah field dari struct:", sampleType.NumField())
	fmt.Println("Nama field ke-0:", sampleType.Field(0).Name)
	fmt.Println("Required field?:", sampleType.Field(0).Tag.Get("required"))
	fmt.Println("Max field?:", sampleType.Field(0).Tag.Get("max"))
	fmt.Println("Min field?:", sampleType.Field(0).Tag.Get("min"))

	// Check if sample is valid
	sample.Name = "Ucup"
	fmt.Println("Is sample valid?", IsValid(sample))

	contoh := ContohLagi{"", ""}
	fmt.Println("Is contoh valid?", IsValid(contoh))
}
