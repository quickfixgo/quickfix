package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type SecurityDefinitionUpdateReport struct {
	quickfixgo.Message
}

func (m *SecurityDefinitionUpdateReport) SecurityReportID() (*field.SecurityReportID, error) {
	f := new(field.SecurityReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionUpdateReport) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionUpdateReport) SecurityResponseID() (*field.SecurityResponseID, error) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionUpdateReport) SecurityResponseType() (*field.SecurityResponseType, error) {
	f := new(field.SecurityResponseType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionUpdateReport) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionUpdateReport) SecurityUpdateAction() (*field.SecurityUpdateAction, error) {
	f := new(field.SecurityUpdateAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionUpdateReport) CorporateAction() (*field.CorporateAction, error) {
	f := new(field.CorporateAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionUpdateReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionUpdateReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionUpdateReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityDefinitionUpdateReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
