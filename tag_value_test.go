package quickfix

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTagValue_init(t *testing.T) {
	var tv TagValue
	tv.init(Tag(8), []byte("blahblah"))

	expectedData := []byte("8=blahblah")
	if !bytes.Equal(tv.bytes, expectedData) {
		t.Errorf("Expected %v, got %v", expectedData, tv.bytes)
	}
	expectedValue := []byte("blahblah")
	if !bytes.Equal(tv.value, expectedValue) {
		t.Errorf("Expected %v, got %v", expectedValue, tv.value)
	}
}

func TestTagValue_parse(t *testing.T) {
	stringField := "8=FIX.4.0"
	tv := TagValue{}
	err := tv.parse([]byte(stringField))
	assert.Nil(t, err)
	assert.Equal(t, Tag(8), tv.tag)

	if !bytes.Equal(tv.bytes, []byte(stringField)) {
		t.Errorf("Expected %v got %v", stringField, tv.bytes)
	}

	if !bytes.Equal(tv.value, []byte("FIX.4.0")) {
		t.Error("Unxpected value", tv.value)
	}
}

func TestTagValue_parseFail(t *testing.T) {
	stringField := "not_tag_equal_value"
	var tv TagValue

	assert.NotNil(t, tv.parse([]byte(stringField)))

	stringField = "tag_not_an_int=uhoh"
	assert.NotNil(t, tv.parse([]byte(stringField)))

	stringField = "=notag"
	assert.NotNil(t, tv.parse([]byte(stringField)))
}

func TestTagValue_String(t *testing.T) {
	stringField := "8=FIX.4.0"
	var tv TagValue

	require.Nil(t, tv.parse([]byte(stringField)))
	assert.Equal(t, "8=FIX.4.0", tv.String())
}

func TestTagValue_length(t *testing.T) {
	stringField := "8=FIX.4.0"
	var tv TagValue
	require.Nil(t, tv.parse([]byte(stringField)))
	assert.Equal(t, len(stringField), tv.length())
}

func TestTagValue_total(t *testing.T) {
	stringField := "1=hello"
	var tv TagValue
	require.Nil(t, tv.parse([]byte(stringField)))
	assert.Equal(t, 643, tv.total(), "Total is the summation of the ascii byte values of the field string")
}
