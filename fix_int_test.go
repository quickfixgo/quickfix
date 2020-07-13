package quickfix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFIXInt_Write(t *testing.T) {
	field := FIXInt(5)

	assert.Equal(t, "5", string(field.Write()))
}

func TestFIXInt_Read(t *testing.T) {
	var field FIXInt
	err := field.Read([]byte("15"))
	assert.Nil(t, err, "Unexpected error")
	assert.Equal(t, 15, int(field))

	err = field.Read([]byte("blah"))
	assert.NotNil(t, err, "Unexpected error")
}

func TestFIXInt_Int(t *testing.T) {
	f := FIXInt(4)
	assert.Equal(t, 4, f.Int())
}

func BenchmarkFIXInt_Read(b *testing.B) {
	intBytes := []byte("1500")
	var field FIXInt

	for i := 0; i < b.N; i++ {
		_ = field.Read(intBytes)
	}
}
