package datadictionary

import (
	"encoding/xml"
	"testing"
)

func TestBuildFieldDef(t *testing.T) {
	var tests = []struct {
		element string
	}{
		{"field"},
		{"group"},
	}

	for _, test := range tests {
		xmlField := &XMLComponentMember{XMLName: xml.Name{Local: test.element}, Name: "myfield", Members: []*XMLComponentMember{}}

		fieldTypeByName := make(map[string]*FieldType)
		fieldTypeByName["myfield"] = NewFieldType("some field", 11, "INT")
		dict := &DataDictionary{FieldTypeByName: fieldTypeByName}

		b := &builder{doc: nil, dict: dict}
		f, err := b.buildFieldDef(xmlField)
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		if f.Tag() != 11 {
			t.Errorf("Unexpected tag %v", f.Tag())
		}
		if len(f.childTags()) != 0 {
			t.Errorf("Got %v children too many", len(f.childTags()))
		}
	}
}
