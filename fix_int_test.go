package quickfix

import (
	"testing"
)

func TestInt_Write(t *testing.T) {
	field := FIXInt(5)
	bytes := field.Write()

	if string(bytes) != "5" {
		t.Error("Unexpected bytes ", bytes)
	}
}

func TestInt_Read(t *testing.T) {
	var field FIXInt
	err := field.Read([]byte("15"))

	if err != nil {
		t.Error("Unexpected error", err)
	}

	if int(field) != 15 {
		t.Error("unexpected value", field)
	}

	err = field.Read([]byte("blah"))

	if err == nil {
		t.Error("expected error")
	}
}

func BenchmarkInt_Read(b *testing.B) {
	intBytes := []byte("1500")
	var field FIXInt

	for i := 0; i < b.N; i++ {
		field.Read(intBytes)
	}
}
