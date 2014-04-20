package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
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
	notradingsessions field.NoTradingSessions) TradingSessionListBuilder {
	var builder TradingSessionListBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(notradingsessions)
	return builder
}

//TradSesReqID is a non-required field for TradingSessionList.
func (m TradingSessionList) TradSesReqID() (*field.TradSesReqID, errors.MessageRejectError) {
	f := new(field.TradSesReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesReqID reads a TradSesReqID from TradingSessionList.
func (m TradingSessionList) GetTradSesReqID(f *field.TradSesReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a required field for TradingSessionList.
func (m TradingSessionList) NoTradingSessions() (*field.NoTradingSessions, errors.MessageRejectError) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from TradingSessionList.
func (m TradingSessionList) GetNoTradingSessions(f *field.NoTradingSessions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for TradingSessionList.
func (m TradingSessionList) ApplID() (*field.ApplID, errors.MessageRejectError) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from TradingSessionList.
func (m TradingSessionList) GetApplID(f *field.ApplID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for TradingSessionList.
func (m TradingSessionList) ApplSeqNum() (*field.ApplSeqNum, errors.MessageRejectError) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from TradingSessionList.
func (m TradingSessionList) GetApplSeqNum(f *field.ApplSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for TradingSessionList.
func (m TradingSessionList) ApplLastSeqNum() (*field.ApplLastSeqNum, errors.MessageRejectError) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from TradingSessionList.
func (m TradingSessionList) GetApplLastSeqNum(f *field.ApplLastSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for TradingSessionList.
func (m TradingSessionList) ApplResendFlag() (*field.ApplResendFlag, errors.MessageRejectError) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from TradingSessionList.
func (m TradingSessionList) GetApplResendFlag(f *field.ApplResendFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}
