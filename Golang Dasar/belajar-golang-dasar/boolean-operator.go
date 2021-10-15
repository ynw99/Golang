package main

import "fmt"

func main() {
	ujian := 90
	absensi := 80

	var lulusUjian bool = ujian > 80
	var lulusAbsensi bool = absensi > 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)
}
