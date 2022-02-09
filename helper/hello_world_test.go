package helper

import (
	"fmt"
	"testing"
)

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
