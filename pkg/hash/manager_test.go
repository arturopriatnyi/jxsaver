package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManager_Hash(t *testing.T) {
	m := NewManager()
	data := []byte("testData")

	hash := m.Hash(data)

	assert.NotEqual(t, data, hash)
}
