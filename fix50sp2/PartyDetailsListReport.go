package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildMsgType("CG"))
	builder.Body.Set(partydetailslistreportid)
	return builder
}

//ApplID is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) ApplID() (*field.ApplID, errors.MessageRejectError) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from PartyDetailsListReport.
func (m PartyDetailsListReport) GetApplID(f *field.ApplID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) ApplSeqNum() (*field.ApplSeqNum, errors.MessageRejectError) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from PartyDetailsListReport.
func (m PartyDetailsListReport) GetApplSeqNum(f *field.ApplSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) ApplLastSeqNum() (*field.ApplLastSeqNum, errors.MessageRejectError) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from PartyDetailsListReport.
func (m PartyDetailsListReport) GetApplLastSeqNum(f *field.ApplLastSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) ApplResendFlag() (*field.ApplResendFlag, errors.MessageRejectError) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from PartyDetailsListReport.
func (m PartyDetailsListReport) GetApplResendFlag(f *field.ApplResendFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PartyDetailsListReportID is a required field for PartyDetailsListReport.
func (m PartyDetailsListReport) PartyDetailsListReportID() (*field.PartyDetailsListReportID, errors.MessageRejectError) {
	f := new(field.PartyDetailsListReportID)
	err := m.Body.Get(f)
	return f, err
}

//GetPartyDetailsListReportID reads a PartyDetailsListReportID from PartyDetailsListReport.
func (m PartyDetailsListReport) GetPartyDetailsListReportID(f *field.PartyDetailsListReportID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PartyDetailsListRequestID is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) PartyDetailsListRequestID() (*field.PartyDetailsListRequestID, errors.MessageRejectError) {
	f := new(field.PartyDetailsListRequestID)
	err := m.Body.Get(f)
	return f, err
}

//GetPartyDetailsListRequestID reads a PartyDetailsListRequestID from PartyDetailsListReport.
func (m PartyDetailsListReport) GetPartyDetailsListRequestID(f *field.PartyDetailsListRequestID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PartyDetailsRequestResult is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) PartyDetailsRequestResult() (*field.PartyDetailsRequestResult, errors.MessageRejectError) {
	f := new(field.PartyDetailsRequestResult)
	err := m.Body.Get(f)
	return f, err
}

//GetPartyDetailsRequestResult reads a PartyDetailsRequestResult from PartyDetailsListReport.
func (m PartyDetailsListReport) GetPartyDetailsRequestResult(f *field.PartyDetailsRequestResult) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoPartyList is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) TotNoPartyList() (*field.TotNoPartyList, errors.MessageRejectError) {
	f := new(field.TotNoPartyList)
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoPartyList reads a TotNoPartyList from PartyDetailsListReport.
func (m PartyDetailsListReport) GetTotNoPartyList(f *field.TotNoPartyList) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) LastFragment() (*field.LastFragment, errors.MessageRejectError) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from PartyDetailsListReport.
func (m PartyDetailsListReport) GetLastFragment(f *field.LastFragment) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyList is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) NoPartyList() (*field.NoPartyList, errors.MessageRejectError) {
	f := new(field.NoPartyList)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyList reads a NoPartyList from PartyDetailsListReport.
func (m PartyDetailsListReport) GetNoPartyList(f *field.NoPartyList) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from PartyDetailsListReport.
func (m PartyDetailsListReport) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from PartyDetailsListReport.
func (m PartyDetailsListReport) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from PartyDetailsListReport.
func (m PartyDetailsListReport) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
