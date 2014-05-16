//Package partydetailslistreport msg type = CG.
package partydetailslistreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a PartyDetailsListReport wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//ApplID is a non-required field for PartyDetailsListReport.
func (m Message) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from PartyDetailsListReport.
func (m Message) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for PartyDetailsListReport.
func (m Message) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from PartyDetailsListReport.
func (m Message) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for PartyDetailsListReport.
func (m Message) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from PartyDetailsListReport.
func (m Message) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for PartyDetailsListReport.
func (m Message) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from PartyDetailsListReport.
func (m Message) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PartyDetailsListReportID is a required field for PartyDetailsListReport.
func (m Message) PartyDetailsListReportID() (*field.PartyDetailsListReportIDField, errors.MessageRejectError) {
	f := &field.PartyDetailsListReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPartyDetailsListReportID reads a PartyDetailsListReportID from PartyDetailsListReport.
func (m Message) GetPartyDetailsListReportID(f *field.PartyDetailsListReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PartyDetailsListRequestID is a non-required field for PartyDetailsListReport.
func (m Message) PartyDetailsListRequestID() (*field.PartyDetailsListRequestIDField, errors.MessageRejectError) {
	f := &field.PartyDetailsListRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPartyDetailsListRequestID reads a PartyDetailsListRequestID from PartyDetailsListReport.
func (m Message) GetPartyDetailsListRequestID(f *field.PartyDetailsListRequestIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PartyDetailsRequestResult is a non-required field for PartyDetailsListReport.
func (m Message) PartyDetailsRequestResult() (*field.PartyDetailsRequestResultField, errors.MessageRejectError) {
	f := &field.PartyDetailsRequestResultField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPartyDetailsRequestResult reads a PartyDetailsRequestResult from PartyDetailsListReport.
func (m Message) GetPartyDetailsRequestResult(f *field.PartyDetailsRequestResultField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoPartyList is a non-required field for PartyDetailsListReport.
func (m Message) TotNoPartyList() (*field.TotNoPartyListField, errors.MessageRejectError) {
	f := &field.TotNoPartyListField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoPartyList reads a TotNoPartyList from PartyDetailsListReport.
func (m Message) GetTotNoPartyList(f *field.TotNoPartyListField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for PartyDetailsListReport.
func (m Message) LastFragment() (*field.LastFragmentField, errors.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from PartyDetailsListReport.
func (m Message) GetLastFragment(f *field.LastFragmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyList is a non-required field for PartyDetailsListReport.
func (m Message) NoPartyList() (*field.NoPartyListField, errors.MessageRejectError) {
	f := &field.NoPartyListField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyList reads a NoPartyList from PartyDetailsListReport.
func (m Message) GetNoPartyList(f *field.NoPartyListField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for PartyDetailsListReport.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from PartyDetailsListReport.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for PartyDetailsListReport.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from PartyDetailsListReport.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for PartyDetailsListReport.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from PartyDetailsListReport.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds PartyDetailsListReport messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for PartyDetailsListReport.
func Builder(
	partydetailslistreportid *field.PartyDetailsListReportIDField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("CG"))
	builder.Body().Set(partydetailslistreportid)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "CG", r
}
