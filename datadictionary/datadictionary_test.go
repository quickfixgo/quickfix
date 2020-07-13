package datadictionary

import (
	"testing"
)

func TestParseBadPath(t *testing.T) {
	if _, err := Parse("../spec/bogus.xml"); err == nil {
		t.Error("Expected err")
	}
}

func TestParseRecursiveComponents(t *testing.T) {
	dict, err := Parse("../spec/FIX44.xml")

	if dict == nil {
		t.Error("Dictionary is nil")
	}

	if err != nil {
		t.Errorf("Unexpected err: %v", err)
	}
}

var cachedDataDictionary *DataDictionary

func dict() (*DataDictionary, error) {
	if cachedDataDictionary != nil {
		return cachedDataDictionary, nil
	}

	var err error
	cachedDataDictionary, err = Parse("../spec/FIX43.xml")
	return cachedDataDictionary, err
}

func TestComponents(t *testing.T) {
	d, _ := dict()
	if _, ok := d.ComponentTypes["SpreadOrBenchmarkCurveData"]; ok != true {
		t.Error("Component not found")
	}
}

func TestFieldsByTag(t *testing.T) {
	d, _ := dict()

	var tests = []struct {
		Tag         int
		Name        string
		Type        string
		EnumsAreNil bool
	}{
		{655, "ContraLegRefID", "STRING", true},
		{658, "QuoteRequestRejectReason", "INT", false},
	}

	for _, test := range tests {
		f, ok := d.FieldTypeByTag[test.Tag]

		if !ok {
			t.Errorf("%v not found", test.Tag)
		}

		if f.Name() != test.Name {
			t.Errorf("Expected %v got %v", test.Name, f.Name())
		}

		if f.Type != test.Type {
			t.Errorf("Expected %v got %v", test.Type, f.Type)
		}

		switch {

		case f.Enums != nil && test.EnumsAreNil:
			t.Errorf("Expected no enums")
		case f.Enums == nil && !test.EnumsAreNil:
			t.Errorf("Expected enums")
		}
	}
}

func TestEnumFieldsByTag(t *testing.T) {

	d, _ := dict()
	f := d.FieldTypeByTag[658]

	var tests = []struct {
		Value       string
		Description string
	}{
		{"1", "UNKNOWN_SYMBOL"},
		{"2", "EXCHANGE"},
		{"3", "QUOTE_REQUEST_EXCEEDS_LIMIT"},
		{"4", "TOO_LATE_TO_ENTER"},
		{"5", "INVALID_PRICE"},
		{"6", "NOT_AUTHORIZED_TO_REQUEST_QUOTE"},
	}

	if len(f.Enums) != len(tests) {
		t.Errorf("Expected %v enums got %v", len(tests), len(f.Enums))
	}

	for _, test := range tests {
		if enum, ok := f.Enums[test.Value]; !ok {
			t.Errorf("Expected Enum %v", test.Value)
		} else {

			if enum.Value != test.Value {
				t.Errorf("Expected value %v got %v", test.Value, enum.Value)
			}

			if enum.Description != test.Description {
				t.Errorf("Expected description %v got %v", test.Description, enum.Description)
			}
		}
	}
}

func TestDataDictionaryMessages(t *testing.T) {
	d, _ := dict()
	_, ok := d.Messages["D"]
	if !ok {
		t.Error("Did not find message")
	}
}

func TestDataDictionaryHeader(t *testing.T) {
	d, _ := dict()
	if d.Header == nil {
		t.Error("Did not find header")
	}
}

func TestDataDictionaryTrailer(t *testing.T) {
	d, _ := dict()
	if d.Trailer == nil {
		t.Error("Did not find trailer")
	}
}

func TestMessageRequiredTags(t *testing.T) {
	d, _ := dict()

	nos := d.Messages["D"]

	var tests = []struct {
		*MessageDef
		Tag      int
		Required bool
	}{
		{nos, 11, true},
		{nos, 526, false},
		{d.Header, 8, true},
		{d.Header, 115, false},
		{d.Trailer, 10, true},
		{d.Trailer, 89, false},
	}

	for _, test := range tests {
		switch _, required := test.MessageDef.RequiredTags[test.Tag]; {
		case required && !test.Required:
			t.Errorf("%v should not be required", test.Tag)
		case !required && test.Required:
			t.Errorf("%v should not required", test.Tag)
		}
	}
}

func TestMessageTags(t *testing.T) {
	d, _ := dict()

	nos := d.Messages["D"]

	var tests = []struct {
		*MessageDef
		Tag int
	}{
		{nos, 11},
		{nos, 526},
		{d.Header, 8},
		{d.Header, 115},
		{d.Trailer, 10},
		{d.Trailer, 89},
	}

	for _, test := range tests {
		if _, known := test.MessageDef.Tags[test.Tag]; !known {
			t.Errorf("%v is not known", test.Tag)
		}
	}
}
