package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	tradsesreqid *field.TradSesReqIDField,
	subscriptionrequesttype *field.SubscriptionRequestTypeField) TradingSessionListRequestBuilder {
	var builder TradingSessionListRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("BI"))
	builder.Body.Set(tradsesreqid)
	builder.Body.Set(subscriptionrequesttype)
	return builder
}

//TradSesReqID is a required field for TradingSessionListRequest.
func (m TradingSessionListRequest) TradSesReqID() (*field.TradSesReqIDField, errors.MessageRejectError) {
	f := &field.TradSesReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesReqID reads a TradSesReqID from TradingSessionListRequest.
func (m TradingSessionListRequest) GetTradSesReqID(f *field.TradSesReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from TradingSessionListRequest.
func (m TradingSessionListRequest) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) TradingSessionSubID() (*field.TradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from TradingSessionListRequest.
func (m TradingSessionListRequest) GetTradingSessionSubID(f *field.TradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from TradingSessionListRequest.
func (m TradingSessionListRequest) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesMethod is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) TradSesMethod() (*field.TradSesMethodField, errors.MessageRejectError) {
	f := &field.TradSesMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesMethod reads a TradSesMethod from TradingSessionListRequest.
func (m TradingSessionListRequest) GetTradSesMethod(f *field.TradSesMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesMode is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) TradSesMode() (*field.TradSesModeField, errors.MessageRejectError) {
	f := &field.TradSesModeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesMode reads a TradSesMode from TradingSessionListRequest.
func (m TradingSessionListRequest) GetTradSesMode(f *field.TradSesModeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a required field for TradingSessionListRequest.
func (m TradingSessionListRequest) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, errors.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from TradingSessionListRequest.
func (m TradingSessionListRequest) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) MarketID() (*field.MarketIDField, errors.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from TradingSessionListRequest.
func (m TradingSessionListRequest) GetMarketID(f *field.MarketIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for TradingSessionListRequest.
func (m TradingSessionListRequest) MarketSegmentID() (*field.MarketSegmentIDField, errors.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from TradingSessionListRequest.
func (m TradingSessionListRequest) GetMarketSegmentID(f *field.MarketSegmentIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}
