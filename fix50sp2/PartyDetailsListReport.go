package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
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
func (m PartyDetailsListReport) ApplID() (field.ApplID, errors.MessageRejectError) {
	var f field.ApplID
	err := m.Body.Get(&f)
	return f, err
}

//ApplSeqNum is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) ApplSeqNum() (field.ApplSeqNum, errors.MessageRejectError) {
	var f field.ApplSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplLastSeqNum is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) ApplLastSeqNum() (field.ApplLastSeqNum, errors.MessageRejectError) {
	var f field.ApplLastSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplResendFlag is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) ApplResendFlag() (field.ApplResendFlag, errors.MessageRejectError) {
	var f field.ApplResendFlag
	err := m.Body.Get(&f)
	return f, err
}

//PartyDetailsListReportID is a required field for PartyDetailsListReport.
func (m PartyDetailsListReport) PartyDetailsListReportID() (field.PartyDetailsListReportID, errors.MessageRejectError) {
	var f field.PartyDetailsListReportID
	err := m.Body.Get(&f)
	return f, err
}

//PartyDetailsListRequestID is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) PartyDetailsListRequestID() (field.PartyDetailsListRequestID, errors.MessageRejectError) {
	var f field.PartyDetailsListRequestID
	err := m.Body.Get(&f)
	return f, err
}

//PartyDetailsRequestResult is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) PartyDetailsRequestResult() (field.PartyDetailsRequestResult, errors.MessageRejectError) {
	var f field.PartyDetailsRequestResult
	err := m.Body.Get(&f)
	return f, err
}

//TotNoPartyList is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) TotNoPartyList() (field.TotNoPartyList, errors.MessageRejectError) {
	var f field.TotNoPartyList
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) LastFragment() (field.LastFragment, errors.MessageRejectError) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyList is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) NoPartyList() (field.NoPartyList, errors.MessageRejectError) {
	var f field.NoPartyList
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
