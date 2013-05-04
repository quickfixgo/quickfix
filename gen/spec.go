package gen

import (
	"encoding/xml"
	"os"
)

type Field struct {
	Name     string `xml:"name,attr"`
	Required string `xml:"required,attr"`
}

type Component struct {
	Name     string `xml:"name,attr"`
	Required string `xml:"required,attr"`
}

type FieldValue struct {
	Enum        string `xml:"enum,attr"`
	Description string `xml:"description,attr"`
}

type FieldType struct {
	Number      int          `xml:"number,attr"`
	Name        string       `xml:"name,attr"`
	Type        string       `xml:"type,attr"`
	FieldValues []FieldValue `xml:"value"`
}

type ComponentType struct {
	Fields     []Field     `xml:"field"`
	Components []Component `xml:"component"`
	Groups     []Group     `xml:"group"`
}

type Group struct {
	Name       string      `xml:"name,attr"`
	Required   string      `xml:"required,attr"`
	Fields     []Field     `xml:"field"`
	Components []Component `xml:"component"`
}

type Message struct {
	Name       string      `xml:"name,attr"`
	MsgCat     string      `xml:"msgcat,attr"`
	MsgType    string      `xml:"msgtype,attr"`
	Fields     []Field     `xml:"field"`
	Groups     []Group     `xml:"group"`
	Components []Component `xml:"component"`
}

type FixSpec struct {
	FixType     string `xml:"type,attr"`
	Major       int    `xml:"major,attr"`
	Minor       int    `xml:"minor,attr"`
	ServicePack int    `xml:"servicepack,attr"`

	Header         []Field         `xml:"header>field"`
	Messages       []Message       `xml:"messages>message"`
	Trailer        []Field         `xml:"trailer>field"`
	FieldTypes     []FieldType     `xml:"fields>field"`
	ComponentTypes []ComponentType `xml:"components>component"`

	FieldTypeMap map[string]FieldType
}

func (spec *FixSpec) init() {
	spec.FieldTypeMap = make(map[string]FieldType)

	for _, f := range spec.FieldTypes {
		spec.FieldTypeMap[f.Name] = f
	}
}

func ParseFixSpec(path string) (*FixSpec, error) {
	xmlFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	var spec FixSpec
	decoder := xml.NewDecoder(xmlFile)
	if err := decoder.Decode(&spec); err != nil {
		return nil, err
	}

	spec.init()

	return &spec, nil
}
