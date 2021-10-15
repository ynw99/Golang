package main

import "fmt"

func main() {
	name1 := "Yusuf"
	name2 := "Yusuf"
	name3 := "Ucup"

	var result1 bool = name1 == name2
	var result2 bool = name1 == name3
	var result3 bool = name1 > name3

	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)

	value1 := 100
	var value2 = 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
}
