package fix50sp2

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type PartyDetailsListReport struct {
	quickfixgo.Message
}

func (m *PartyDetailsListReport) PartyDetailsListReportID() (*field.PartyDetailsListReportID, error) {
	f := new(field.PartyDetailsListReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListReport) PartyDetailsListRequestID() (*field.PartyDetailsListRequestID, error) {
	f := new(field.PartyDetailsListRequestID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListReport) PartyDetailsRequestResult() (*field.PartyDetailsRequestResult, error) {
	f := new(field.PartyDetailsRequestResult)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListReport) TotNoPartyList() (*field.TotNoPartyList, error) {
	f := new(field.TotNoPartyList)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListReport) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
