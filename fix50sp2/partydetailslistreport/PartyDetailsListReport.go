//Package partydetailslistreport msg type = CG.
package partydetailslistreport

import (
	"github.com/quickfixgo/quickfix"
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
func (m Message) ApplID() (*field.ApplIDField, quickfix.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from PartyDetailsListReport.
func (m Message) GetApplID(f *field.ApplIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for PartyDetailsListReport.
func (m Message) ApplSeqNum() (*field.ApplSeqNumField, quickfix.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from PartyDetailsListReport.
func (m Message) GetApplSeqNum(f *field.ApplSeqNumField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for PartyDetailsListReport.
func (m Message) ApplLastSeqNum() (*field.ApplLastSeqNumField, quickfix.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from PartyDetailsListReport.
func (m Message) GetApplLastSeqNum(f *field.ApplLastSeqNumField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for PartyDetailsListReport.
func (m Message) ApplResendFlag() (*field.ApplResendFlagField, quickfix.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from PartyDetailsListReport.
func (m Message) GetApplResendFlag(f *field.ApplResendFlagField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PartyDetailsListReportID is a required field for PartyDetailsListReport.
func (m Message) PartyDetailsListReportID() (*field.PartyDetailsListReportIDField, quickfix.MessageRejectError) {
	f := &field.PartyDetailsListReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPartyDetailsListReportID reads a PartyDetailsListReportID from PartyDetailsListReport.
func (m Message) GetPartyDetailsListReportID(f *field.PartyDetailsListReportIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PartyDetailsListRequestID is a non-required field for PartyDetailsListReport.
func (m Message) PartyDetailsListRequestID() (*field.PartyDetailsListRequestIDField, quickfix.MessageRejectError) {
	f := &field.PartyDetailsListRequestIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPartyDetailsListRequestID reads a PartyDetailsListRequestID from PartyDetailsListReport.
func (m Message) GetPartyDetailsListRequestID(f *field.PartyDetailsListRequestIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PartyDetailsRequestResult is a non-required field for PartyDetailsListReport.
func (m Message) PartyDetailsRequestResult() (*field.PartyDetailsRequestResultField, quickfix.MessageRejectError) {
	f := &field.PartyDetailsRequestResultField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPartyDetailsRequestResult reads a PartyDetailsRequestResult from PartyDetailsListReport.
func (m Message) GetPartyDetailsRequestResult(f *field.PartyDetailsRequestResultField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoPartyList is a non-required field for PartyDetailsListReport.
func (m Message) TotNoPartyList() (*field.TotNoPartyListField, quickfix.MessageRejectError) {
	f := &field.TotNoPartyListField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoPartyList reads a TotNoPartyList from PartyDetailsListReport.
func (m Message) GetTotNoPartyList(f *field.TotNoPartyListField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for PartyDetailsListReport.
func (m Message) LastFragment() (*field.LastFragmentField, quickfix.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from PartyDetailsListReport.
func (m Message) GetLastFragment(f *field.LastFragmentField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyList is a non-required field for PartyDetailsListReport.
func (m Message) NoPartyList() (*field.NoPartyListField, quickfix.MessageRejectError) {
	f := &field.NoPartyListField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyList reads a NoPartyList from PartyDetailsListReport.
func (m Message) GetNoPartyList(f *field.NoPartyListField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for PartyDetailsListReport.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from PartyDetailsListReport.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for PartyDetailsListReport.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from PartyDetailsListReport.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for PartyDetailsListReport.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from PartyDetailsListReport.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
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
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "CG", r
}
