package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Package container/list -> implementasi double linked list
	data := list.New()
	data.PushBack("Yusuf")
	data.PushBack("Nur")
	data.PushBack("Wahid")

	// Dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// Dari belakang ke depan
	for e := data.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}

	fmt.Println("Front data:", data.Front().Value)
	fmt.Println("Back data:", data.Back().Value)
}
