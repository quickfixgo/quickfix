package fix50sp2

import (
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

//NewTradingSessionListBuilder returns an initialized TradingSessionListBuilder with specified required fields.
func NewTradingSessionListBuilder(
	notradingsessions field.NoTradingSessions) *TradingSessionListBuilder {
	builder := new(TradingSessionListBuilder)
	builder.Body.Set(notradingsessions)
	return builder
}

//TradSesReqID is a non-required field for TradingSessionList.
func (m *TradingSessionList) TradSesReqID() (*field.TradSesReqID, error) {
	f := new(field.TradSesReqID)
	err := m.Body.Get(f)
	return f, err
}

//NoTradingSessions is a required field for TradingSessionList.
func (m *TradingSessionList) NoTradingSessions() (*field.NoTradingSessions, error) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//ApplID is a non-required field for TradingSessionList.
func (m *TradingSessionList) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//ApplSeqNum is a non-required field for TradingSessionList.
func (m *TradingSessionList) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplLastSeqNum is a non-required field for TradingSessionList.
func (m *TradingSessionList) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplResendFlag is a non-required field for TradingSessionList.
func (m *TradingSessionList) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}
