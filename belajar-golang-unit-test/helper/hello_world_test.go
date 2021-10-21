package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldYusuf(t *testing.T) {
	result := HelloWorld("Yusuf")

	if result != "Hello Yusuf" {
		// error
		// t.Fail()
		t.Error("Result must be Hello Yusuf")
	}

	fmt.Println("TestHelloWorldYusuf Done")
}

func TestHelloWorldNur(t *testing.T) {
	result := HelloWorld("Nur")

	if result != "Hello Nur" {
		// error
		// t.FailNow()
		t.Fatal("Result must be Hello Nur")
	}

	fmt.Println("TestHelloWorldNur Done")
}
