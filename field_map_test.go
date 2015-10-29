package quickfix

import (
	"testing"
)

func TestFieldMap_SetAndGet(t *testing.T) {
	fMap := FieldMap{}
	fMap.init(normalFieldOrder)

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
			if err == nil {
				t.Error("Expected Error")
			}
			continue
		}

		if err != nil {
			t.Error("Unexpected Error", err)
		}

		if string(testField) != tc.expectValue {
			t.Errorf("Expected %v got %v", tc.expectValue, testField)
		}
	}
}

func TestFieldMap_Length(t *testing.T) {
	fMap := FieldMap{}
	fMap.init(normalFieldOrder)
	fMap.SetField(1, FIXString("hello"))
	fMap.SetField(2, FIXString("world"))
	fMap.SetField(8, FIXString("FIX.4.4"))
	fMap.SetField(9, FIXInt(100))
	fMap.SetField(10, FIXString("100"))

	if fMap.length() != 16 {
		t.Error("Length should include all fields but beginString, bodyLength, and checkSum- got ", fMap.length())
	}
}

func TestFieldMap_Total(t *testing.T) {

	fMap := FieldMap{}
	fMap.init(normalFieldOrder)
	fMap.SetField(1, FIXString("hello"))
	fMap.SetField(2, FIXString("world"))
	fMap.SetField(8, FIXString("FIX.4.4"))
	fMap.SetField(Tag(9), FIXInt(100))
	fMap.SetField(10, FIXString("100"))

	if fMap.total() != 2116 {
		t.Error("Total should includes all fields but checkSum- got ", fMap.total())
	}
}
