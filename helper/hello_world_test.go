package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Abdul")
	if result != "Hello Abdul" {
		panic("Result is not 'Hello Abdul")
	}
}
