package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Package regexp -> RE2
	var pattern *regexp.Regexp = regexp.MustCompile("e[a-z]o")

	fmt.Println(pattern.MatchString("eko"))
	fmt.Println(pattern.MatchString("eDo"))
	fmt.Println(pattern.MatchString("eto"))

	search := pattern.FindAllString("eko eka edo eto eyo eki", -1)
	fmt.Println(search)
}
