package datadictionary

import (
	"encoding/xml"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BuildSuite struct {
	suite.Suite
	doc *XMLDoc
	*builder
}

func TestBuildSuite(t *testing.T) {
	suite.Run(t, new(BuildSuite))
}

func (s *BuildSuite) SetupTest() {
	s.doc = &XMLDoc{Type: "FIX", Major: "4", Minor: "5"}
	s.builder = new(builder)
}

func (s *BuildSuite) TestValidTypes() {
	var tests = []string{
		"FIX",
		"FIXT",
	}

	for _, test := range tests {
		s.doc.Type = test
		dict, err := s.builder.build(s.doc)

		s.Nil(err)
		s.NotNil(dict)
		s.Equal(test, dict.FIXType)
	}
}

func (s *BuildSuite) TestInvalidTypes() {
	var tests = []string{
		"",
		"invalid",
	}

	for _, test := range tests {
		s.doc.Type = test
		_, err := s.builder.build(s.doc)

		s.NotNil(err)
	}
}

func (s *BuildSuite) TestValidMajor() {
	var tests = []int{
		4,
		5,
	}

	for _, test := range tests {
		s.doc.Major = strconv.Itoa(test)
		dict, err := s.builder.build(s.doc)

		s.Nil(err)
		s.NotNil(dict)
		s.Equal(test, dict.Major)
	}
}

func (s *BuildSuite) TestInvalidMajor() {
	var tests = []string{
		"",
		"notanumber",
	}

	for _, test := range tests {
		s.doc.Major = test
		_, err := s.builder.build(s.doc)

		s.NotNil(err)
	}
}

func (s *BuildSuite) TestValidMinor() {
	var tests = []int{
		4,
		5,
	}

	for _, test := range tests {
		s.doc.Minor = strconv.Itoa(test)
		dict, err := s.builder.build(s.doc)

		s.Nil(err)
		s.NotNil(dict)
		s.Equal(test, dict.Minor)
	}
}

func (s *BuildSuite) TestInvalidMinor() {
	var tests = []string{
		"",
		"notanumber",
	}

	for _, test := range tests {
		s.doc.Minor = test
		_, err := s.builder.build(s.doc)

		s.NotNil(err)
	}
}

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

		assert.Nil(t, err)
		assert.Equal(t, 11, f.Tag())
		assert.Empty(t, f.childTags())
	}
}
