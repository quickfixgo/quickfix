package fix50sp2

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type TradingSessionListRequest struct {
	quickfixgo.Message
}

func (m *TradingSessionListRequest) TradSesReqID() (*field.TradSesReqID, error) {
	f := new(field.TradSesReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionListRequest) TradSesMethod() (*field.TradSesMethod, error) {
	f := new(field.TradSesMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionListRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionListRequest) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionListRequest) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionListRequest) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionListRequest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionListRequest) TradSesMode() (*field.TradSesMode, error) {
	f := new(field.TradSesMode)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionListRequest) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}
