package fix42

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type MarketDataRequest struct {
	quickfixgo.Message
}

func (m *MarketDataRequest) MDReqID() (*field.MDReqID, error) {
	f := new(field.MDReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) MarketDepth() (*field.MarketDepth, error) {
	f := new(field.MarketDepth)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) MDUpdateType() (*field.MDUpdateType, error) {
	f := new(field.MDUpdateType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) AggregatedBook() (*field.AggregatedBook, error) {
	f := new(field.AggregatedBook)
	err := m.Body.Get(f)
	return f, err
}
