package main

import "fmt"

type Filter func(string) string

func sayhelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Kotor" {
		return "..."
	} else {
		return name
	}
}

func main() {
	name := "Yusuf"
	kotor := "Kotor"

	sayHelloWithFilter(name, spamFilter)
	sayhelloWithFilter(kotor, spamFilter)
}
