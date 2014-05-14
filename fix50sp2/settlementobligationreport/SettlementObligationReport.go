//Package settlementobligationreport msg type = BQ.
package settlementobligationreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a SettlementObligationReport wrapper for the generic Message type
type Message struct {
	message.Message
}

//ClearingBusinessDate is a non-required field for SettlementObligationReport.
func (m Message) ClearingBusinessDate() (*field.ClearingBusinessDateField, errors.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from SettlementObligationReport.
func (m Message) GetClearingBusinessDate(f *field.ClearingBusinessDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlementCycleNo is a non-required field for SettlementObligationReport.
func (m Message) SettlementCycleNo() (*field.SettlementCycleNoField, errors.MessageRejectError) {
	f := &field.SettlementCycleNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlementCycleNo reads a SettlementCycleNo from SettlementObligationReport.
func (m Message) GetSettlementCycleNo(f *field.SettlementCycleNoField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlObligMsgID is a required field for SettlementObligationReport.
func (m Message) SettlObligMsgID() (*field.SettlObligMsgIDField, errors.MessageRejectError) {
	f := &field.SettlObligMsgIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlObligMsgID reads a SettlObligMsgID from SettlementObligationReport.
func (m Message) GetSettlObligMsgID(f *field.SettlObligMsgIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlObligMode is a required field for SettlementObligationReport.
func (m Message) SettlObligMode() (*field.SettlObligModeField, errors.MessageRejectError) {
	f := &field.SettlObligModeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlObligMode reads a SettlObligMode from SettlementObligationReport.
func (m Message) GetSettlObligMode(f *field.SettlObligModeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for SettlementObligationReport.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from SettlementObligationReport.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for SettlementObligationReport.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from SettlementObligationReport.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for SettlementObligationReport.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from SettlementObligationReport.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for SettlementObligationReport.
func (m Message) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SettlementObligationReport.
func (m Message) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSettlOblig is a non-required field for SettlementObligationReport.
func (m Message) NoSettlOblig() (*field.NoSettlObligField, errors.MessageRejectError) {
	f := &field.NoSettlObligField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSettlOblig reads a NoSettlOblig from SettlementObligationReport.
func (m Message) GetNoSettlOblig(f *field.NoSettlObligField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for SettlementObligationReport.
func (m Message) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from SettlementObligationReport.
func (m Message) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for SettlementObligationReport.
func (m Message) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from SettlementObligationReport.
func (m Message) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for SettlementObligationReport.
func (m Message) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from SettlementObligationReport.
func (m Message) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for SettlementObligationReport.
func (m Message) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from SettlementObligationReport.
func (m Message) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds SettlementObligationReport messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for SettlementObligationReport.
func Builder(
	settlobligmsgid *field.SettlObligMsgIDField,
	settlobligmode *field.SettlObligModeField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("BQ"))
	builder.Body().Set(settlobligmsgid)
	builder.Body().Set(settlobligmode)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "BQ", r
}
