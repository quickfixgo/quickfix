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

//CreatePartyDetailsListReportBuilder returns an initialized PartyDetailsListReportBuilder with specified required fields.
func CreatePartyDetailsListReportBuilder(
	partydetailslistreportid field.PartyDetailsListReportID) PartyDetailsListReportBuilder {
	var builder PartyDetailsListReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(partydetailslistreportid)
	return builder
}

//ApplID is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) ApplID() (field.ApplID, error) {
	var f field.ApplID
	err := m.Body.Get(&f)
	return f, err
}

//ApplSeqNum is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) ApplSeqNum() (field.ApplSeqNum, error) {
	var f field.ApplSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplLastSeqNum is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) ApplLastSeqNum() (field.ApplLastSeqNum, error) {
	var f field.ApplLastSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplResendFlag is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) ApplResendFlag() (field.ApplResendFlag, error) {
	var f field.ApplResendFlag
	err := m.Body.Get(&f)
	return f, err
}

//PartyDetailsListReportID is a required field for PartyDetailsListReport.
func (m PartyDetailsListReport) PartyDetailsListReportID() (field.PartyDetailsListReportID, error) {
	var f field.PartyDetailsListReportID
	err := m.Body.Get(&f)
	return f, err
}

//PartyDetailsListRequestID is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) PartyDetailsListRequestID() (field.PartyDetailsListRequestID, error) {
	var f field.PartyDetailsListRequestID
	err := m.Body.Get(&f)
	return f, err
}

//PartyDetailsRequestResult is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) PartyDetailsRequestResult() (field.PartyDetailsRequestResult, error) {
	var f field.PartyDetailsRequestResult
	err := m.Body.Get(&f)
	return f, err
}

//TotNoPartyList is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) TotNoPartyList() (field.TotNoPartyList, error) {
	var f field.TotNoPartyList
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) LastFragment() (field.LastFragment, error) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyList is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) NoPartyList() (field.NoPartyList, error) {
	var f field.NoPartyList
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
