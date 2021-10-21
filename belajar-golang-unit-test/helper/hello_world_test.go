package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Yusuf")
	require.Equal(t, "Hello Yusuf", result, "Result must be Hello Yusuf")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Yusuf")
	assert.Equal(t, "Hello Yusuf", result, "Result must be Hello Yusuf")
	fmt.Println("TestHelloWorld with Assert Done")
}

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
