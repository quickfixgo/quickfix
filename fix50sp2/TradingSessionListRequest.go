package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//TradingSessionListRequest msg type = BI.
type TradingSessionListRequest struct {
	message.Message
}

//TradingSessionListRequestBuilder builds TradingSessionListRequest messages.
type TradingSessionListRequestBuilder struct {
	message.MessageBuilder
}

//CreateTradingSessionListRequestBuilder returns an initialized TradingSessionListRequestBuilder with specified required fields.
func CreateTradingSessionListRequestBuilder(
	tradsesreqid field.TradSesReqID,
	subscriptionrequesttype field.SubscriptionRequestType) TradingSessionListRequestBuilder {
	var builder TradingSessionListRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(tradsesreqid)
	builder.Body.Set(subscriptionrequesttype)
	return builder
}

//TradSesReqID is a required field for TradingSessionListRequest.
func (m TradingSessionListRequest) TradSesReqID() (field.TradSesReqID, errors.MessageRejectError) {
	var f field.TradSesReqID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) TradingSessionID() (field.TradingSessionID, errors.MessageRejectError) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) TradingSessionSubID() (field.TradingSessionSubID, errors.MessageRejectError) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//TradSesMethod is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) TradSesMethod() (field.TradSesMethod, errors.MessageRejectError) {
	var f field.TradSesMethod
	err := m.Body.Get(&f)
	return f, err
}

//TradSesMode is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) TradSesMode() (field.TradSesMode, errors.MessageRejectError) {
	var f field.TradSesMode
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a required field for TradingSessionListRequest.
func (m TradingSessionListRequest) SubscriptionRequestType() (field.SubscriptionRequestType, errors.MessageRejectError) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//MarketID is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) MarketID() (field.MarketID, errors.MessageRejectError) {
	var f field.MarketID
	err := m.Body.Get(&f)
	return f, err
}

//MarketSegmentID is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) MarketSegmentID() (field.MarketSegmentID, errors.MessageRejectError) {
	var f field.MarketSegmentID
	err := m.Body.Get(&f)
	return f, err
}
