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
	t.Run("Hello Abdul", func(t *testing.T) {
		result := HelloWorld("Abdul")
		assert.Equal(t, "Hello Abdul", result, "Result is not 'Hello Abdul")

	})

	t.Run("Hello Rizki", func(t *testing.T) {
		result := HelloWorld("Rizki")
		assert.Equal(t, "Hello Rizki", result, "Result is not 'Hello Rizki")

	})
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
