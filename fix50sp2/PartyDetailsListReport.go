package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	partydetailslistreportid *field.PartyDetailsListReportIDField) PartyDetailsListReportBuilder {
	var builder PartyDetailsListReportBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("CG"))
	builder.Body.Set(partydetailslistreportid)
	return builder
}

//ApplID is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from PartyDetailsListReport.
func (m PartyDetailsListReport) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from PartyDetailsListReport.
func (m PartyDetailsListReport) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from PartyDetailsListReport.
func (m PartyDetailsListReport) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from PartyDetailsListReport.
func (m PartyDetailsListReport) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PartyDetailsListReportID is a required field for PartyDetailsListReport.
func (m PartyDetailsListReport) PartyDetailsListReportID() (*field.PartyDetailsListReportIDField, errors.MessageRejectError) {
	f := &field.PartyDetailsListReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPartyDetailsListReportID reads a PartyDetailsListReportID from PartyDetailsListReport.
func (m PartyDetailsListReport) GetPartyDetailsListReportID(f *field.PartyDetailsListReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PartyDetailsListRequestID is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) PartyDetailsListRequestID() (*field.PartyDetailsListRequestIDField, errors.MessageRejectError) {
	f := &field.PartyDetailsListRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPartyDetailsListRequestID reads a PartyDetailsListRequestID from PartyDetailsListReport.
func (m PartyDetailsListReport) GetPartyDetailsListRequestID(f *field.PartyDetailsListRequestIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PartyDetailsRequestResult is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) PartyDetailsRequestResult() (*field.PartyDetailsRequestResultField, errors.MessageRejectError) {
	f := &field.PartyDetailsRequestResultField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPartyDetailsRequestResult reads a PartyDetailsRequestResult from PartyDetailsListReport.
func (m PartyDetailsListReport) GetPartyDetailsRequestResult(f *field.PartyDetailsRequestResultField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoPartyList is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) TotNoPartyList() (*field.TotNoPartyListField, errors.MessageRejectError) {
	f := &field.TotNoPartyListField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoPartyList reads a TotNoPartyList from PartyDetailsListReport.
func (m PartyDetailsListReport) GetTotNoPartyList(f *field.TotNoPartyListField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) LastFragment() (*field.LastFragmentField, errors.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from PartyDetailsListReport.
func (m PartyDetailsListReport) GetLastFragment(f *field.LastFragmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyList is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) NoPartyList() (*field.NoPartyListField, errors.MessageRejectError) {
	f := &field.NoPartyListField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyList reads a NoPartyList from PartyDetailsListReport.
func (m PartyDetailsListReport) GetNoPartyList(f *field.NoPartyListField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from PartyDetailsListReport.
func (m PartyDetailsListReport) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from PartyDetailsListReport.
func (m PartyDetailsListReport) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for PartyDetailsListReport.
func (m PartyDetailsListReport) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from PartyDetailsListReport.
func (m PartyDetailsListReport) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
