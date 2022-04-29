package module

import (
	"github.com/test-go/testify/assert"
	"testing"
)

func TestDo(t *testing.T) {
	assert.Equal(t, "hello", Do())
}
