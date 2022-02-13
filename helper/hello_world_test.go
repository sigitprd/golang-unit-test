package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before TestMain")
	m.Run()
	fmt.Println("After TestMain")
}

func BenchmarkHelloWorldSigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Sigit")
	}
}

func BenchmarkHelloWorldPriadi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Priadi")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Sigit)",
			request:  "Sigit",
			expected: "Hello Sigit",
		},
		{
			name:     "HelloWorld(Priadi)",
			request:  "Priadi",
			expected: "Hello Priadi",
		},
		{
			name:     "HelloWorld(Budi)",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "HelloWorld(Joko)",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HelloWorld(tt.request)
			// assert.Equal(t, tt.expected, result)
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Sigit", func(t *testing.T) {
		result := HelloWorld("Sigit")
		assert.Equal(t, "Hello Sigit", result, "Expected Hello Sigit, but got "+result)
		fmt.Println("Test HelloWorldSigitRequire Success")
	})
	t.Run("Priadi", func(t *testing.T) {
		result := HelloWorld("Priadi")
		assert.Equal(t, "Hello Priadi", result, "Expected Hello Priadi, but got "+result)
		fmt.Println("Test HelloWorldPriadiRequire Success")
	})
}

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
