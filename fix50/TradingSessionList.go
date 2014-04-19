package fix50

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
func (m TradingSessionList) TradSesReqID() (field.TradSesReqID, errors.MessageRejectError) {
	var f field.TradSesReqID
	err := m.Body.Get(&f)
	return f, err
}

//NoTradingSessions is a required field for TradingSessionList.
func (m TradingSessionList) NoTradingSessions() (field.NoTradingSessions, errors.MessageRejectError) {
	var f field.NoTradingSessions
	err := m.Body.Get(&f)
	return f, err
}
