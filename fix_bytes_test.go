package quickfix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFIXBytesWrite(t *testing.T) {
	val := []byte("hello")
	field := FIXBytes(val)
	b := field.Write()

	assert.Equal(t, val, b)
}

func TestFIXBytesRead(t *testing.T) {
	val := []byte("world")
	var field FIXBytes
	assert.Nil(t, field.Read(val))
	assert.Equal(t, val, []byte(field))
}
