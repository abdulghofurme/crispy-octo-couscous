package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Abdul")
	if result != "Hi Abdul" {
		t.Fatal("Result is not 'Hello Abdul")
	}
}
