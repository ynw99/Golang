package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	// Package and import using Go Modules because GOPATH can't
	// be used.
	helper.SayHello("Ucup")
	fmt.Println(helper.Application)
}
