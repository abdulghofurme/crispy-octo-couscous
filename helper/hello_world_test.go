package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// before
	fmt.Println("AFTER UNIT TEST")
}

func TestAllHelloWorld(t *testing.T) {
	tests := []struct {
		name, request, expected string
	}{
		{
			name:     "HelloWorld(Abdul)",
			request:  "Abdul",
			expected: "Hello Abdul",
		},
		{
			name:     "HelloWorld(Rizki)",
			request:  "Rizki",
			expected: "Hello Rizki",
		},
		{
			name:     "HelloWorld(Hafshoh)",
			request:  "Hafshoh",
			expected: "Hello Hafshoh",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result, "Result is not 'Hello Abdul")

		})
	}
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Abdul")
	assert.Equal(t, "Hello Abdul", result, "Result is not 'Hello Abdul")
}

func TestSkipHelloWorld(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}
	result := HelloWorld("OS")
	assert.Equal(t, "Hello OS", result, "Result is not 'Hello OS")
}
