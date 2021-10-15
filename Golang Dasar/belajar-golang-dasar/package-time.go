package main

import (
	"fmt"
	"time"
)

func main() {
	// Package time
	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// Membuat waktu secara manual
	utc := time.Date(2021, 10, 12, 18, 30, 30, 0, time.UTC)

	fmt.Println(utc)

	// Parsing string to date
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2021-10-13")
	fmt.Println(parse)
}
