package datadictionary

import (
	"encoding/xml"
	"github.com/quickfixgo/quickfix/fix/tag"
	. "launchpad.net/gocheck"
)

var _ = Suite(&BuildTests{})

type BuildTests struct{}

func (s *BuildTests) TestBuildField(c *C) {
	xmlField := &XMLComponentMember{XMLName: xml.Name{Local: "field"}, Name: "myfield", Members: []*XMLComponentMember{}}

	fieldTypeByName := make(map[string]*FieldType)
	fieldTypeByName["myfield"] = &FieldType{Tag: tag.ClOrdID}
	dict := &DataDictionary{FieldTypeByName: fieldTypeByName}

	b := &builder{doc: nil, dict: dict}
	f, err := b.buildFieldDef(xmlField)
	c.Check(err, IsNil)

	c.Check(f.Tag, Equals, tag.ClOrdID)
	c.Check(len(f.childTags()), Equals, 0)
}

func (s *BuildTests) TestBuildFieldGroup(c *C) {
	xmlField := &XMLComponentMember{XMLName: xml.Name{Local: "group"}, Name: "myfield", Members: []*XMLComponentMember{}}

	fieldTypeByName := make(map[string]*FieldType)
	fieldTypeByName["myfield"] = &FieldType{Tag: tag.ClOrdID}
	dict := &DataDictionary{FieldTypeByName: fieldTypeByName}

	b := &builder{doc: nil, dict: dict}
	f, err := b.buildFieldDef(xmlField)
	c.Check(err, IsNil)

	c.Check(f.Tag, Equals, tag.ClOrdID)
	c.Check(len(f.childTags()), Equals, 0)
}
