package quickfix

import (
	"testing"
)

func TestIntField_NewField(t *testing.T) {
	field := NewIntField(Tag(1), 5)
	if field.Tag() != Tag(1) {
		t.Error("Unexpected tag ", field.Tag())
	}

	if field.Value != 5 {
		t.Error("Unexpected value", field.Value)
	}
}

func TestIntField_Write(t *testing.T) {
	field := NewIntField(Tag(1), 5)
	bytes := field.Write()

	if string(bytes) != "5" {
		t.Error("Unexpected bytes ", bytes)
	}
}

func TestIntField_Read(t *testing.T) {
	field := new(IntField)
	err := field.Read([]TagValue{TagValue{Value: []byte("15")}})

	if err != nil {
		t.Error("Unexpected error", err)
	}

	if field.Value != 15 {
		t.Error("unexpected value", field.Value)
	}

	err = field.Read([]TagValue{TagValue{Value: []byte("blah")}})

	if err == nil {
		t.Error("expected error")
	}
}

func BenchmarkIntField_Read(b *testing.B) {
	intBytes := []byte("1500")
	field := &IntField{}

	for i := 0; i < b.N; i++ {
		field.Read([]TagValue{TagValue{Value: intBytes}})
	}
}
