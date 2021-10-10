package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blackList Blacklist) {
	if blackList(name) {
		fmt.Println("You're blocked", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

// func blackListAdmin(name string) bool {
// 	return name == "admin"
// }

// func blackListRoot(name string) bool {
// 	return name == "root"
// }

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("root", func(name string) bool {
		return name == "root"
	})

}
