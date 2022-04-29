package module_b

import (
	"github.com/test-go/testify/assert"
	"testing"
)

func TestDo(t *testing.T) {
	assert.Equal(t, "3yZe7d", Do())
}
