package fix50sp2

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type PartyDetailsListReport struct {
	quickfix.Message
}

func (m *PartyDetailsListReport) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
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
func (m *PartyDetailsListReport) TotNoPartyList() (*field.TotNoPartyList, error) {
	f := new(field.TotNoPartyList)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListReport) NoPartyList() (*field.NoPartyList, error) {
	f := new(field.NoPartyList)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListReport) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListReport) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListReport) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *PartyDetailsListReport) PartyDetailsRequestResult() (*field.PartyDetailsRequestResult, error) {
	f := new(field.PartyDetailsRequestResult)
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
