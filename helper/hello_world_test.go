package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skip on windows")
	}

	result := HelloWorld("Sigit")
	require.Equal(t, "Hello Sigit", result, "Expected Hello Sigit, but got "+result)
}

func TestHelloWorldSigit(t *testing.T) {
	result := HelloWorld("Sigit")
	if result != "Hello Sigit" {
		// panic("Expected Hello Sigit, but got " + result)
		t.Error("Expected Hello Sigit, but got " + result)
		// t.Fail()
	}

	fmt.Println("Test HelloWorldSigit Success")
	// type args struct {
	// 	name string
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want string
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := HelloWorld(tt.args.name); got != tt.want {
	// 			t.Errorf("HelloWorld() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}

func TestHelloWorldPriadi(t *testing.T) {
	result := HelloWorld("Priadi")
	if result != "Hello Priadi" {
		t.Fatal("Expected Hello Priadi, but got " + result)
		// t.FailNow()
		// panic("Expected Hello Priadi, but got " + result)
	}
	fmt.Println("Test HelloWorldPriadi Success")
}

func TestHelloWorldSigitAssertion(t *testing.T) {
	result := HelloWorld("Sigit")
	assert.Equal(t, "Hello Sigit", result)
	fmt.Println("Test HelloWorldSigitAssertion Success")
}

func TestHelloWorldPriadiAssertion(t *testing.T) {
	result := HelloWorld("Priadi")
	assert.Equal(t, "Hello Priadi", result)
	fmt.Println("Test HelloWorldPriadiAssertion Success")
}

func TestHelloWorldSigitRequire(t *testing.T) {
	result := HelloWorld("Sigit")
	require.Equal(t, "Hello Sigit", result)
	fmt.Println("Test HelloWorldSigitRequire Success")
}

func TestHelloWorldPriadiRequire(t *testing.T) {
	result := HelloWorld("Priadi")
	require.Equal(t, "Hello Priadi", result)
	fmt.Println("Test HelloWorldPriadiRequire Success")
}
