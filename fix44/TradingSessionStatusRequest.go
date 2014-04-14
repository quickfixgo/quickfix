package fix44

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

//NewTradingSessionStatusRequestBuilder returns an initialized TradingSessionStatusRequestBuilder with specified required fields.
func NewTradingSessionStatusRequestBuilder(
	tradsesreqid field.TradSesReqID,
	subscriptionrequesttype field.SubscriptionRequestType) *TradingSessionStatusRequestBuilder {
	builder := new(TradingSessionStatusRequestBuilder)
	builder.Body.Set(tradsesreqid)
	builder.Body.Set(subscriptionrequesttype)
	return builder
}

//TradSesReqID is a required field for TradingSessionStatusRequest.
func (m *TradingSessionStatusRequest) TradSesReqID() (*field.TradSesReqID, error) {
	f := new(field.TradSesReqID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for TradingSessionStatusRequest.
func (m *TradingSessionStatusRequest) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for TradingSessionStatusRequest.
func (m *TradingSessionStatusRequest) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//TradSesMethod is a non-required field for TradingSessionStatusRequest.
func (m *TradingSessionStatusRequest) TradSesMethod() (*field.TradSesMethod, error) {
	f := new(field.TradSesMethod)
	err := m.Body.Get(f)
	return f, err
}

//TradSesMode is a non-required field for TradingSessionStatusRequest.
func (m *TradingSessionStatusRequest) TradSesMode() (*field.TradSesMode, error) {
	f := new(field.TradSesMode)
	err := m.Body.Get(f)
	return f, err
}

//SubscriptionRequestType is a required field for TradingSessionStatusRequest.
func (m *TradingSessionStatusRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
