package main

import "fmt"

func getGoodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	sayGoodBye := getGoodbye

	result := sayGoodBye("Ucup")
	fmt.Println(result)
	fmt.Println(getGoodbye("Ucup"))
}
