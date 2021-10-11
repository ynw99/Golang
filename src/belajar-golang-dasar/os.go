package main

import (
	"fmt"
	"os"
)

func main() {
	// Package OS
	// Get system argument
	arg := os.Args
	fmt.Println(arg[1])

	// Get hostname
	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname:", hostname)
	} else {
		fmt.Println("Error:", err.Error())
	}

	// Get environment
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
