package fix50sp2

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type MarketDefinitionUpdateReport struct {
	quickfixgo.Message
}

func (m *MarketDefinitionUpdateReport) MarketReportID() (*field.MarketReportID, error) {
	f := new(field.MarketReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionUpdateReport) MarketReqID() (*field.MarketReqID, error) {
	f := new(field.MarketReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionUpdateReport) MarketUpdateAction() (*field.MarketUpdateAction, error) {
	f := new(field.MarketUpdateAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionUpdateReport) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionUpdateReport) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionUpdateReport) MarketSegmentDesc() (*field.MarketSegmentDesc, error) {
	f := new(field.MarketSegmentDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionUpdateReport) EncodedMktSegmDescLen() (*field.EncodedMktSegmDescLen, error) {
	f := new(field.EncodedMktSegmDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionUpdateReport) EncodedMktSegmDesc() (*field.EncodedMktSegmDesc, error) {
	f := new(field.EncodedMktSegmDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionUpdateReport) ParentMktSegmID() (*field.ParentMktSegmID, error) {
	f := new(field.ParentMktSegmID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionUpdateReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionUpdateReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionUpdateReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionUpdateReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionUpdateReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
