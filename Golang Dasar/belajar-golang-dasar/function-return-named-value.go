package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Yusuf"
	middleName = "Nur"
	lastName = "Wahid"

	return firstName, middleName, lastName
}

func getPerson() (fullName string, isGanteng bool) {
	fullName = "Yusuf Nur Wahid"
	isGanteng = true

	return fullName, isGanteng
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println(firstName, middleName, lastName)

	fullName, isGanteng := getPerson()
	fmt.Println(fullName, isGanteng)
}
