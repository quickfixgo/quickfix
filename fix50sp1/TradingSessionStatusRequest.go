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
	tradsesreqid *field.TradSesReqIDField,
	subscriptionrequesttype *field.SubscriptionRequestTypeField) TradingSessionStatusRequestBuilder {
	var builder TradingSessionStatusRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("g"))
	builder.Body.Set(tradsesreqid)
	builder.Body.Set(subscriptionrequesttype)
	return builder
}

//TradSesReqID is a required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) TradSesReqID() (*field.TradSesReqIDField, errors.MessageRejectError) {
	f := &field.TradSesReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesReqID reads a TradSesReqID from TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) GetTradSesReqID(f *field.TradSesReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) TradingSessionSubID() (*field.TradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) GetTradingSessionSubID(f *field.TradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesMethod is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) TradSesMethod() (*field.TradSesMethodField, errors.MessageRejectError) {
	f := &field.TradSesMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesMethod reads a TradSesMethod from TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) GetTradSesMethod(f *field.TradSesMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesMode is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) TradSesMode() (*field.TradSesModeField, errors.MessageRejectError) {
	f := &field.TradSesModeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesMode reads a TradSesMode from TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) GetTradSesMode(f *field.TradSesModeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, errors.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) MarketID() (*field.MarketIDField, errors.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) GetMarketID(f *field.MarketIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) MarketSegmentID() (*field.MarketSegmentIDField, errors.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from TradingSessionStatusRequest.
func (m TradingSessionStatusRequest) GetMarketSegmentID(f *field.MarketSegmentIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}
