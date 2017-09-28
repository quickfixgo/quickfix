package datadictionary

import (
	"encoding/xml"
)

//XMLDoc is the unmarshalled root of a FIX Dictionary.
type XMLDoc struct {
	Type        string `xml:"type,attr"`
	Major       string `xml:"major,attr"`
	Minor       string `xml:"minor,attr"`
	ServicePack int    `xml:"servicepack,attr"`

	Header     *XMLComponent   `xml:"header"`
	Trailer    *XMLComponent   `xml:"trailer"`
	Messages   []*XMLComponent `xml:"messages>message"`
	Components []*XMLComponent `xml:"components>component"`
	Fields     []*XMLField     `xml:"fields>field"`
}

// XMLComponent can represent header, trailer, messages/message, or components/component xml elements.
type XMLComponent struct {
	Name    string `xml:"name,attr"`
	MsgCat  string `xml:"msgcat,attr"`
	MsgType string `xml:"msgtype,attr"`

	Members []*XMLComponentMember `xml:",any"`
}

//XMLField represents the fields/field xml element.
type XMLField struct {
	Number int         `xml:"number,attr"`
	Name   string      `xml:"name,attr"`
	Type   string      `xml:"type,attr"`
	Values []*XMLValue `xml:"value"`
}

//XMLValue represents the fields/field/value xml element.
type XMLValue struct {
	Enum        string `xml:"enum,attr"`
	Description string `xml:"description,attr"`
}

//XMLComponentMember represents child elements of header, trailer, messages/message, and components/component elements
type XMLComponentMember struct {
	XMLName  xml.Name
	Name     string `xml:"name,attr"`
	Required string `xml:"required,attr"`

	Members []*XMLComponentMember `xml:",any"`
}

func (member XMLComponentMember) isComponent() bool {
	return member.XMLName.Local == "component"
}

func (member XMLComponentMember) isGroup() bool {
	return member.XMLName.Local == "group"
}

func (member XMLComponentMember) isRequired() bool {
	return member.Required == "Y"
}
