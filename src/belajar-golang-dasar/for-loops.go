package main

import "fmt"

func main() {
	counter := 1

	// Mirip while loops
	for counter <= 10 {
		fmt.Println("Perulangan ke-", counter)
		counter++
	}

	// For loops biasa
	for i := 1; i < counter; i++ {
		fmt.Println("i ke-", i)
	}

	// For range (mirip for each)
	slice := []string{"Yusuf", "Nur", "Wahid"}
	for index, value := range slice {
		fmt.Println("Index", index, "=", value)
	}

	for _, value := range slice {
		fmt.Println(value)
	}

	// for range bisa juga dipakai di map
	person := make(map[string]string)
	person["name"] = "Yusuf"
	person["job"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
