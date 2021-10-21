package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Sub test
func TestSubTest(t *testing.T) {
	t.Run("Yusuf", func(t *testing.T) {
		result := HelloWorld("Yusuf")
		require.Equal(t, "Hello Yusuf", result, "Result must be Hello Yusuf")
	})

	t.Run("Nur", func(t *testing.T) {
		result := HelloWorld("Nur")
		require.Equal(t, "Hello Nur", result, "Result must be Hello Yusuf")
	})
}

// Main test
func TestMain(m *testing.M) {
	// Before
	fmt.Println("================")
	fmt.Println("BEFORE UNIT TEST")
	fmt.Println("================")

	// Run all unit tests
	m.Run()

	// After
	fmt.Println("================")
	fmt.Println("AFTER UNIT TEST")
	fmt.Println("================")
}

// Skip test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can not run on Linux OS")
	}

	result := HelloWorld("Yusuf")
	require.Equal(t, "Hello Yusuf", result, "Result must be Hello Yusuf")
}

// Unit tests
// Tidak disarankan menggunakan unit test
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
