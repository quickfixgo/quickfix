package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//TradingSessionList msg type = BJ.
type TradingSessionList struct {
	message.Message
}

//TradingSessionListBuilder builds TradingSessionList messages.
type TradingSessionListBuilder struct {
	message.MessageBuilder
}

//CreateTradingSessionListBuilder returns an initialized TradingSessionListBuilder with specified required fields.
func CreateTradingSessionListBuilder(
	notradingsessions *field.NoTradingSessionsField) TradingSessionListBuilder {
	var builder TradingSessionListBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("BJ"))
	builder.Body.Set(notradingsessions)
	return builder
}

//TradSesReqID is a non-required field for TradingSessionList.
func (m TradingSessionList) TradSesReqID() (*field.TradSesReqIDField, errors.MessageRejectError) {
	f := &field.TradSesReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesReqID reads a TradSesReqID from TradingSessionList.
func (m TradingSessionList) GetTradSesReqID(f *field.TradSesReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a required field for TradingSessionList.
func (m TradingSessionList) NoTradingSessions() (*field.NoTradingSessionsField, errors.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from TradingSessionList.
func (m TradingSessionList) GetNoTradingSessions(f *field.NoTradingSessionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for TradingSessionList.
func (m TradingSessionList) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from TradingSessionList.
func (m TradingSessionList) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for TradingSessionList.
func (m TradingSessionList) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from TradingSessionList.
func (m TradingSessionList) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for TradingSessionList.
func (m TradingSessionList) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from TradingSessionList.
func (m TradingSessionList) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for TradingSessionList.
func (m TradingSessionList) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from TradingSessionList.
func (m TradingSessionList) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}
