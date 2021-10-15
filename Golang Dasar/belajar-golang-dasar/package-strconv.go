package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Package string conversion
	// Parse -> string to other data types
	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	// Format -> other data types to string
	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	// Atoi - Itoa
	valueInt, _ := strconv.Atoi("5000")
	fmt.Println(valueInt)
}
