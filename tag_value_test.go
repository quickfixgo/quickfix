package quickfix

import (
	"bytes"
	"testing"
)

func TestTagValue_init(t *testing.T) {
	var tv tagValue
	tv.init(Tag(8), []byte("blahblah"))

	expectedData := []byte("8=blahblah")
	if !bytes.Equal(tv.bytes, expectedData) {
		t.Errorf("Expected %v, got %v", expectedData, tv.bytes)
	}
	expectedValue := []byte("blahblah")
	if !bytes.Equal(tv.Value, expectedValue) {
		t.Errorf("Expected %v, got %v", expectedValue, tv.Value)
	}
}

func TestTagValue_parse(t *testing.T) {
	stringField := "8=FIX.4.0"
	tv := tagValue{}
	err := tv.parse([]byte(stringField))

	if err != nil {
		t.Error("Unexpected error", err)
	}

	if tv.Tag != Tag(8) {
		t.Error("Unexpected tag", tv.Tag)
	}

	if !bytes.Equal(tv.bytes, []byte(stringField)) {
		t.Errorf("Expected %v got %v", stringField, tv.bytes)
	}

	if !bytes.Equal(tv.Value, []byte("FIX.4.0")) {
		t.Error("Unxpected value", tv.Value)
	}
}

func TestTagValue_parseFail(t *testing.T) {
	stringField := "not_tag_equal_value"
	tv := tagValue{}

	err := tv.parse([]byte(stringField))

	if err == nil {
		t.Error("Expected Error")
	}

	stringField = "tag_not_an_int=uhoh"
	err = tv.parse([]byte(stringField))
	if err == nil {
		t.Error("Expected Error")
	}
}

func TestTagValue_String(t *testing.T) {
	stringField := "8=FIX.4.0"
	tv := tagValue{}
	tv.parse([]byte(stringField))

	if tv.String() != "8=FIX.4.0" {
		t.Error("Unexpected string value", tv.String())
	}
}

func TestTagValue_length(t *testing.T) {
	stringField := "8=FIX.4.0"
	tv := tagValue{}
	tv.parse([]byte(stringField))

	if tv.length() != len(stringField) {
		t.Error("Unexpected length", tv.length())
	}
}

func TestTagValue_total(t *testing.T) {
	stringField := "1=hello"
	tv := tagValue{}
	tv.parse([]byte(stringField))
	if tv.total() != 643 {
		t.Error("Total is the summation of the ascii byte values of the field string, got ", tv.total())
	}
}
