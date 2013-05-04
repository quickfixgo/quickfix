package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type MarketDefinition struct {
	quickfixgo.Message
}

func (m *MarketDefinition) MarketReportID() (*field.MarketReportID, error) {
	f := new(field.MarketReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) MarketReqID() (*field.MarketReqID, error) {
	f := new(field.MarketReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) MarketSegmentDesc() (*field.MarketSegmentDesc, error) {
	f := new(field.MarketSegmentDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) EncodedMktSegmDescLen() (*field.EncodedMktSegmDescLen, error) {
	f := new(field.EncodedMktSegmDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) EncodedMktSegmDesc() (*field.EncodedMktSegmDesc, error) {
	f := new(field.EncodedMktSegmDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) ParentMktSegmID() (*field.ParentMktSegmID, error) {
	f := new(field.ParentMktSegmID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinition) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
