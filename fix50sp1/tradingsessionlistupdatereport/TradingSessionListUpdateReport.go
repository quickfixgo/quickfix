//Package tradingsessionlistupdatereport msg type = BS.
package tradingsessionlistupdatereport

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

//Message is a TradingSessionListUpdateReport wrapper for the generic Message type
type Message struct {
	message.Message
}

//TradSesReqID is a non-required field for TradingSessionListUpdateReport.
func (m Message) TradSesReqID() (*field.TradSesReqIDField, errors.MessageRejectError) {
	f := &field.TradSesReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesReqID reads a TradSesReqID from TradingSessionListUpdateReport.
func (m Message) GetTradSesReqID(f *field.TradSesReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesUpdateAction is a non-required field for TradingSessionListUpdateReport.
func (m Message) TradSesUpdateAction() (*field.TradSesUpdateActionField, errors.MessageRejectError) {
	f := &field.TradSesUpdateActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesUpdateAction reads a TradSesUpdateAction from TradingSessionListUpdateReport.
func (m Message) GetTradSesUpdateAction(f *field.TradSesUpdateActionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a required field for TradingSessionListUpdateReport.
func (m Message) NoTradingSessions() (*field.NoTradingSessionsField, errors.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from TradingSessionListUpdateReport.
func (m Message) GetNoTradingSessions(f *field.NoTradingSessionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for TradingSessionListUpdateReport.
func (m Message) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from TradingSessionListUpdateReport.
func (m Message) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for TradingSessionListUpdateReport.
func (m Message) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from TradingSessionListUpdateReport.
func (m Message) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for TradingSessionListUpdateReport.
func (m Message) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from TradingSessionListUpdateReport.
func (m Message) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for TradingSessionListUpdateReport.
func (m Message) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from TradingSessionListUpdateReport.
func (m Message) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds TradingSessionListUpdateReport messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for TradingSessionListUpdateReport.
func Builder(
	notradingsessions *field.NoTradingSessionsField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("BS"))
	builder.Body().Set(notradingsessions)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "BS", r
}
