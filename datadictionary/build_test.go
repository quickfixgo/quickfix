package datadictionary

import (
	"encoding/xml"
	"github.com/cbusbey/quickfixgo/tag"
	. "launchpad.net/gocheck"
)

var _ = Suite(&BuildTests{})

type BuildTests struct{}

func (s *BuildTests) TestBuildField(c *C) {
	xmlField := &XMLComponentMember{XMLName: xml.Name{Local: "field"}, Name: "myfield", Members: []*XMLComponentMember{}}

	fieldTypeByName := make(map[string]*FieldType)
	fieldTypeByName["myfield"] = &FieldType{Tag: tag.ClOrdID}
	dict := &DataDictionary{FieldTypeByName: fieldTypeByName}

	f, err := buildFieldInterface(xmlField, dict)
	c.Check(err, IsNil)

	c.Check(f.Tag(), Equals, tag.ClOrdID)
}

func (s *BuildTests) TestBuildFieldGroup(c *C) {
	xmlField := &XMLComponentMember{XMLName: xml.Name{Local: "group"}, Name: "myfield", Members: []*XMLComponentMember{}}

	fieldTypeByName := make(map[string]*FieldType)
	fieldTypeByName["myfield"] = &FieldType{Tag: tag.ClOrdID}
	dict := &DataDictionary{FieldTypeByName: fieldTypeByName}

	f, err := buildFieldInterface(xmlField, dict)
	c.Check(err, IsNil)

	c.Check(f.Tag(), Equals, tag.ClOrdID)
	c.Check(len(f.MemberTags()), Equals, 0)
}
