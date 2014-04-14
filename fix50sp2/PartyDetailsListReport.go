package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//PartyDetailsListReport msg type = CG.
type PartyDetailsListReport struct {
	message.Message
}

//PartyDetailsListReportBuilder builds PartyDetailsListReport messages.
type PartyDetailsListReportBuilder struct {
	message.MessageBuilder
}

//NewPartyDetailsListReportBuilder returns an initialized PartyDetailsListReportBuilder with specified required fields.
func NewPartyDetailsListReportBuilder(
	partydetailslistreportid field.PartyDetailsListReportID) *PartyDetailsListReportBuilder {
	builder := new(PartyDetailsListReportBuilder)
	builder.Body.Set(partydetailslistreportid)
	return builder
}

//ApplID is a non-required field for PartyDetailsListReport.
func (m *PartyDetailsListReport) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//ApplSeqNum is a non-required field for PartyDetailsListReport.
func (m *PartyDetailsListReport) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplLastSeqNum is a non-required field for PartyDetailsListReport.
func (m *PartyDetailsListReport) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplResendFlag is a non-required field for PartyDetailsListReport.
func (m *PartyDetailsListReport) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}

//PartyDetailsListReportID is a required field for PartyDetailsListReport.
func (m *PartyDetailsListReport) PartyDetailsListReportID() (*field.PartyDetailsListReportID, error) {
	f := new(field.PartyDetailsListReportID)
	err := m.Body.Get(f)
	return f, err
}

//PartyDetailsListRequestID is a non-required field for PartyDetailsListReport.
func (m *PartyDetailsListReport) PartyDetailsListRequestID() (*field.PartyDetailsListRequestID, error) {
	f := new(field.PartyDetailsListRequestID)
	err := m.Body.Get(f)
	return f, err
}

//PartyDetailsRequestResult is a non-required field for PartyDetailsListReport.
func (m *PartyDetailsListReport) PartyDetailsRequestResult() (*field.PartyDetailsRequestResult, error) {
	f := new(field.PartyDetailsRequestResult)
	err := m.Body.Get(f)
	return f, err
}

//TotNoPartyList is a non-required field for PartyDetailsListReport.
func (m *PartyDetailsListReport) TotNoPartyList() (*field.TotNoPartyList, error) {
	f := new(field.TotNoPartyList)
	err := m.Body.Get(f)
	return f, err
}

//LastFragment is a non-required field for PartyDetailsListReport.
func (m *PartyDetailsListReport) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyList is a non-required field for PartyDetailsListReport.
func (m *PartyDetailsListReport) NoPartyList() (*field.NoPartyList, error) {
	f := new(field.NoPartyList)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for PartyDetailsListReport.
func (m *PartyDetailsListReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for PartyDetailsListReport.
func (m *PartyDetailsListReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for PartyDetailsListReport.
func (m *PartyDetailsListReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
