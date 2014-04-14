package fix50sp2

import (
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

//NewTradingSessionListRequestBuilder returns an initialized TradingSessionListRequestBuilder with specified required fields.
func NewTradingSessionListRequestBuilder(
	tradsesreqid field.TradSesReqID,
	subscriptionrequesttype field.SubscriptionRequestType) *TradingSessionListRequestBuilder {
	builder := new(TradingSessionListRequestBuilder)
	builder.Body.Set(tradsesreqid)
	builder.Body.Set(subscriptionrequesttype)
	return builder
}

//TradSesReqID is a required field for TradingSessionListRequest.
func (m *TradingSessionListRequest) TradSesReqID() (*field.TradSesReqID, error) {
	f := new(field.TradSesReqID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for TradingSessionListRequest.
func (m *TradingSessionListRequest) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for TradingSessionListRequest.
func (m *TradingSessionListRequest) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for TradingSessionListRequest.
func (m *TradingSessionListRequest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//TradSesMethod is a non-required field for TradingSessionListRequest.
func (m *TradingSessionListRequest) TradSesMethod() (*field.TradSesMethod, error) {
	f := new(field.TradSesMethod)
	err := m.Body.Get(f)
	return f, err
}

//TradSesMode is a non-required field for TradingSessionListRequest.
func (m *TradingSessionListRequest) TradSesMode() (*field.TradSesMode, error) {
	f := new(field.TradSesMode)
	err := m.Body.Get(f)
	return f, err
}

//SubscriptionRequestType is a required field for TradingSessionListRequest.
func (m *TradingSessionListRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//MarketID is a non-required field for TradingSessionListRequest.
func (m *TradingSessionListRequest) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//MarketSegmentID is a non-required field for TradingSessionListRequest.
func (m *TradingSessionListRequest) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}
