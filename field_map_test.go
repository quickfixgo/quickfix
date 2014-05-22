package quickfix

import (
	"github.com/quickfixgo/quickfix/fix"
	"testing"
)

func TestFieldMap_SetAndGet(t *testing.T) {
	fMap := fieldMap{}
	fMap.init(normalFieldOrder)

	fMap.Set(fix.NewStringField(1, "hello"))
	fMap.Set(fix.NewStringField(2, "world"))

	var testCases = []struct {
		tag         fix.Tag
		expectErr   bool
		expectValue string
	}{
		{tag: 1, expectValue: "hello"},
		{tag: 2, expectValue: "world"},
		{tag: 44, expectErr: true},
	}

	for _, tc := range testCases {
		testField := fix.NewStringField(tc.tag, "")
		err := fMap.Get(testField)

		if tc.expectErr {
			if err == nil {
				t.Error("Expected Error")
			}
			continue
		}

		if err != nil {
			t.Error("Unexpected Error", err)
		}

		if testField.Value != tc.expectValue {
			t.Errorf("Expected %v got %v", tc.expectValue, testField.Value)
		}
	}
}

func TestFieldMap_Length(t *testing.T) {
	f1 := fix.NewStringField(1, "hello")
	f2 := fix.NewStringField(2, "world")

	beginString := fix.NewStringField(8, "FIX.4.4")
	bodyLength := fix.NewIntField(9, 100)
	checkSum := fix.NewStringField(10, "100")

	fMap := fieldMap{}
	fMap.init(normalFieldOrder)
	fMap.Set(f1)
	fMap.Set(f2)
	fMap.Set(beginString)
	fMap.Set(bodyLength)
	fMap.Set(checkSum)

	if fMap.length() != 16 {
		t.Error("Length should include all fields but beginString, bodyLength, and checkSum- got ", fMap.length())
	}
}

func TestFieldMap_Total(t *testing.T) {
	f1 := fix.NewStringField(1, "hello")
	f2 := fix.NewStringField(2, "world")

	beginString := fix.NewStringField(8, "FIX.4.4")
	bodyLength := fix.NewIntField(9, 100)
	checkSum := fix.NewStringField(10, "100")

	fMap := fieldMap{}
	fMap.init(normalFieldOrder)
	fMap.Set(f1)
	fMap.Set(f2)
	fMap.Set(beginString)
	fMap.Set(bodyLength)
	fMap.Set(checkSum)

	if fMap.total() != 2116 {
		t.Error("Total should includes all fields but checkSum- got ", fMap.total())
	}
}
