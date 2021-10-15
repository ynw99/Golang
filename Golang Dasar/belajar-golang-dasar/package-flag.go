package main

import (
	"flag"
	"fmt"
)

func main() {
	// Package flag
	var host *string = flag.String("host", "localhost", "Put your database host")
	var username *string = flag.String("username", "root", "Put your database username")
	var password *string = flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 100, "Put your number")

	flag.Parse()

	fmt.Println("Host:", *host)
	fmt.Println("User:", *username)
	fmt.Println("Password:", *password)
	fmt.Println("Number:", *number)
}
