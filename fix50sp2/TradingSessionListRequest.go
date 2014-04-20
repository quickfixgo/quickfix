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
func (m TradingSessionListRequest) TradSesReqID() (*field.TradSesReqID, errors.MessageRejectError) {
	f := new(field.TradSesReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesReqID reads a TradSesReqID from TradingSessionListRequest.
func (m TradingSessionListRequest) GetTradSesReqID(f *field.TradSesReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from TradingSessionListRequest.
func (m TradingSessionListRequest) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from TradingSessionListRequest.
func (m TradingSessionListRequest) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from TradingSessionListRequest.
func (m TradingSessionListRequest) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesMethod is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) TradSesMethod() (*field.TradSesMethod, errors.MessageRejectError) {
	f := new(field.TradSesMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesMethod reads a TradSesMethod from TradingSessionListRequest.
func (m TradingSessionListRequest) GetTradSesMethod(f *field.TradSesMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesMode is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) TradSesMode() (*field.TradSesMode, errors.MessageRejectError) {
	f := new(field.TradSesMode)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesMode reads a TradSesMode from TradingSessionListRequest.
func (m TradingSessionListRequest) GetTradSesMode(f *field.TradSesMode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a required field for TradingSessionListRequest.
func (m TradingSessionListRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, errors.MessageRejectError) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from TradingSessionListRequest.
func (m TradingSessionListRequest) GetSubscriptionRequestType(f *field.SubscriptionRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) MarketID() (*field.MarketID, errors.MessageRejectError) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from TradingSessionListRequest.
func (m TradingSessionListRequest) GetMarketID(f *field.MarketID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) MarketSegmentID() (*field.MarketSegmentID, errors.MessageRejectError) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from TradingSessionListRequest.
func (m TradingSessionListRequest) GetMarketSegmentID(f *field.MarketSegmentID) errors.MessageRejectError {
	return m.Body.Get(f)
}
