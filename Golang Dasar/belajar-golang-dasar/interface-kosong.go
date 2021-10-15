package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func any(apapun interface{}) interface{} {
	return apapun
}

func main() {
	// Interface kosong -> interface yang tidak memiliki
	// deklarasi method satupun yang membuat secara otomatis
	// semua tipe data akan menjadi implementasinya
	var data interface{} = Ups(3)

	fmt.Println(data)

	data = any("haha")
	fmt.Println(data)

	data = any(3)
	fmt.Println(data)
}
