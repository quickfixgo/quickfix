package datadictionary

import (
	"encoding/xml"
	"github.com/quickfixgo/quickfix/fix/tag"
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
		fieldTypeByName["myfield"] = &FieldType{Tag: tag.ClOrdID}
		dict := &DataDictionary{FieldTypeByName: fieldTypeByName}

		b := &builder{doc: nil, dict: dict}
		f, err := b.buildFieldDef(xmlField)
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}

		if f.Tag != tag.ClOrdID {
			t.Errorf("Unexpected tag %v", f.Tag)
		}
		if len(f.childTags()) != 0 {
			t.Errorf("Got %v children too many", len(f.childTags()))
		}
	}
}
