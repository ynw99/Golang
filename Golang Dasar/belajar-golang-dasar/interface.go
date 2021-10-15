package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

// Implementasi
func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func main() {
	// Interface -> tipe data yang abstrak, tidak memiliki
	// implementasi langsung.
	// Interface -> berisikan definisi-definisi method dan
	// biasanya digunakan sebagai kontrak.
	yusuf := Person{Name: "Yusuf"}

	sayHello(yusuf)

	cat := Animal{Name: "Puss"}

	sayHello(cat)
}
