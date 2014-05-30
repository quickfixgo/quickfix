package quickfix

import (
	"bytes"
	"github.com/quickfixgo/quickfix/fix"
	"testing"
)

func TestFieldBytes_NewField(t *testing.T) {
	f := newFieldBytes(fix.Tag(8), []byte("blahblah"))
	if f == nil {
		t.Error("f should not be nil")
	}

	expectedData := []byte("8=blahblah")
	if !bytes.Equal(f.Data, expectedData) {
		t.Errorf("Expected %v, got %v", expectedData, f.Data)
	}
	expectedValue := []byte("blahblah")
	if !bytes.Equal(f.Value, expectedValue) {
		t.Errorf("Expected %v, got %v", expectedValue, f.Value)
	}
}

func TestFieldBytes_ParseField(t *testing.T) {
	stringField := "8=FIX.4.0"

	f, err := parseField([]byte(stringField))

	if err != nil {
		t.Error("Unexpected error", err)
	}

	if f.Tag != fix.Tag(8) {
		t.Error("Unexpected tag", f.Tag)
	}

	if !bytes.Equal(f.Data, []byte(stringField)) {
		t.Errorf("Expected %v got %v", stringField, f.Data)
	}

	if !bytes.Equal(f.Value, []byte("FIX.4.0")) {
		t.Error("Unxpected value", f.Value)
	}
}

func TestFieldBytes_ParseFieldFail(t *testing.T) {
	stringField := "not_tag_equal_value"

	_, err := parseField([]byte(stringField))

	if err == nil {
		t.Error("Expected Error")
	}

	stringField = "tag_not_an_int=uhoh"
	_, err = parseField([]byte(stringField))
	if err == nil {
		t.Error("Expected Error")
	}
}

func TestFieldBytes_String(t *testing.T) {
	stringField := "8=FIX.4.0"
	f, _ := parseField([]byte(stringField))

	if f.String() != "8=FIX.4.0" {
		t.Error("Unexpected string value", f.String())
	}
}

func TestFieldBytes_Length(t *testing.T) {
	stringField := "8=FIX.4.0"
	f, _ := parseField([]byte(stringField))

	if f.Length() != len(stringField) {
		t.Error("Unexpected Length", f.Length())
	}
}

func TestFieldBytes_Total(t *testing.T) {
	stringField := "1=hello"
	f, _ := parseField([]byte(stringField))
	if f.Total() != 643 {
		t.Error("Total is the summation of the ascii byte values of the field string, got ", f.Total())
	}
}
