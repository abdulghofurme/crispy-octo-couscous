package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Abdul")
	assert.Equal(t, "Hello Abdul", result, "Result is not 'Hello Abdul")
}
