package fix50

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
