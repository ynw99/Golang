package main

import "golang/belajar-golang-dasar/helper"

func main() {
	// Package and import using Go Modules because GOPATH can't
	// be used.
	helper.SayHello("Ucup")
}
