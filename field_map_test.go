package quickfix

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldMap_Clear(t *testing.T) {
	var fMap FieldMap
	fMap.init()

	fMap.SetField(1, FIXString("hello"))
	fMap.SetField(2, FIXString("world"))

	fMap.Clear()

	if fMap.Has(1) || fMap.Has(2) {
		t.Error("All fields should be cleared")
	}
}

func TestFieldMap_SetAndGet(t *testing.T) {
	var fMap FieldMap
	fMap.init()

	fMap.SetField(1, FIXString("hello"))
	fMap.SetField(2, FIXString("world"))

	var testCases = []struct {
		tag         Tag
		expectErr   bool
		expectValue string
	}{
		{tag: 1, expectValue: "hello"},
		{tag: 2, expectValue: "world"},
		{tag: 44, expectErr: true},
	}

	for _, tc := range testCases {
		var testField FIXString
		err := fMap.GetField(tc.tag, &testField)

		if tc.expectErr {
			assert.NotNil(t, err, "Expected Error")
			continue
		}

		assert.Nil(t, err, "Unexpected error")
		assert.Equal(t, tc.expectValue, string(testField))
	}
}

func TestFieldMap_Length(t *testing.T) {
	var fMap FieldMap
	fMap.init()
	fMap.SetField(1, FIXString("hello"))
	fMap.SetField(2, FIXString("world"))
	fMap.SetField(8, FIXString("FIX.4.4"))
	fMap.SetField(9, FIXInt(100))
	fMap.SetField(10, FIXString("100"))
	assert.Equal(t, 16, fMap.length(), "Length should include all fields but beginString, bodyLength, and checkSum")
}

func TestFieldMap_Total(t *testing.T) {

	var fMap FieldMap
	fMap.init()
	fMap.SetField(1, FIXString("hello"))
	fMap.SetField(2, FIXString("world"))
	fMap.SetField(8, FIXString("FIX.4.4"))
	fMap.SetField(Tag(9), FIXInt(100))
	fMap.SetField(10, FIXString("100"))

	assert.Equal(t, 2116, fMap.total(), "Total should includes all fields but checkSum")
}

func TestFieldMap_TypedSetAndGet(t *testing.T) {
	var fMap FieldMap
	fMap.init()

	fMap.SetString(1, "hello")
	fMap.SetInt(2, 256)

	s, err := fMap.GetString(1)
	assert.Nil(t, err)
	assert.Equal(t, "hello", s)

	i, err := fMap.GetInt(2)
	assert.Nil(t, err)
	assert.Equal(t, 256, i)

	_, err = fMap.GetInt(1)
	assert.NotNil(t, err, "Type mismatch should occur error")

	s, err = fMap.GetString(2)
	assert.Nil(t, err)
	assert.Equal(t, "256", s)

	b, err := fMap.GetBytes(1)
	assert.Nil(t, err)
	assert.True(t, bytes.Equal([]byte("hello"), b))
}

func TestFieldMap_BoolTypedSetAndGet(t *testing.T) {
	var fMap FieldMap
	fMap.init()

	fMap.SetBool(1, true)
	v, err := fMap.GetBool(1)
	assert.Nil(t, err)
	assert.True(t, v)

	s, err := fMap.GetString(1)
	assert.Nil(t, err)
	assert.Equal(t, "Y", s)

	fMap.SetBool(2, false)
	v, err = fMap.GetBool(2)
	assert.Nil(t, err)
	assert.False(t, v)

	s, err = fMap.GetString(2)
	assert.Nil(t, err)
	assert.Equal(t, "N", s)
}
