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
