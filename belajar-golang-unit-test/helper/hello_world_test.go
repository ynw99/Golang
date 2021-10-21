package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Yusuf")

	if result != "Hello Yusuf" {
		// error
		panic("Result is not Hello Yusuf")
	}
}

func TestHelloWorldNur(t *testing.T) {
	result := HelloWorld("Nur")

	if result != "Hello Nur" {
		// error
		panic("Result is not Hello Nur")
	}
}
