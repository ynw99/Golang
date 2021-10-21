package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldYusuf(t *testing.T) {
	result := HelloWorld("Yusuf")

	if result != "Hello Yusuf" {
		// error
		t.Fail()
	}

	fmt.Println("TestHelloWorldYusuf Done")
}

func TestHelloWorldNur(t *testing.T) {
	result := HelloWorld("Nur")

	if result != "Hello Nur" {
		// error
		t.FailNow()
	}

	fmt.Println("TestHelloWorldNur Done")
}
