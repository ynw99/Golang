package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello ", name, "my name is", customer.Name)
}

func (customer Customer) sayHu() {
	fmt.Println("Huuuu from", customer.Name)
}

func main() {
	ucup := Customer{Name: "Ucup"}
	ucup.sayHello("Yusuf")
	ucup.sayHu()
}
