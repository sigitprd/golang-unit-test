package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Sigit")
	if result != "Hello Sigit" {
		panic("Expected Hello Sigit, but got " + result)
	}
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
		panic("Expected Hello Priadi, but got " + result)
	}
}
