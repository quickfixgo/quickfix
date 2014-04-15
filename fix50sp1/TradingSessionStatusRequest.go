package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//TradingSessionStatusRequest msg type = g.
type TradingSessionStatusRequest struct {
	message.Message
}

//TradingSessionStatusRequestBuilder builds TradingSessionStatusRequest messages.
type TradingSessionStatusRequestBuilder struct {
	message.MessageBuilder
}

//CreateTradingSessionStatusRequestBuilder returns an initialized TradingSessionStatusRequestBuilder with specified required fields.
func CreateTradingSessionStatusRequestBuilder(
	tradsesreqid field.TradSesReqID,
	subscriptionrequesttype field.SubscriptionRequestType) TradingSessionStatusRequestBuilder {
	var builder TradingSessionStatusRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(tradsesreqid)
	builder.Body.Set(subscriptionrequesttype)
	return builder
}

//TradSesReqID is a required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) TradSesReqID() (field.TradSesReqID, error) {
	var f field.TradSesReqID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//TradSesMethod is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) TradSesMethod() (field.TradSesMethod, error) {
	var f field.TradSesMethod
	err := m.Body.Get(&f)
	return f, err
}

//TradSesMode is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) TradSesMode() (field.TradSesMode, error) {
	var f field.TradSesMode
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) SubscriptionRequestType() (field.SubscriptionRequestType, error) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//MarketID is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) MarketID() (field.MarketID, error) {
	var f field.MarketID
	err := m.Body.Get(&f)
	return f, err
}

//MarketSegmentID is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) MarketSegmentID() (field.MarketSegmentID, error) {
	var f field.MarketSegmentID
	err := m.Body.Get(&f)
	return f, err
}
