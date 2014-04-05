package fix50

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type MarketDataIncrementalRefresh struct {
	quickfixgo.Message
}

func (m *MarketDataIncrementalRefresh) MDReqID() (*field.MDReqID, error) {
	f := new(field.MDReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataIncrementalRefresh) NoMDEntries() (*field.NoMDEntries, error) {
	f := new(field.NoMDEntries)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataIncrementalRefresh) ApplQueueDepth() (*field.ApplQueueDepth, error) {
	f := new(field.ApplQueueDepth)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataIncrementalRefresh) ApplQueueResolution() (*field.ApplQueueResolution, error) {
	f := new(field.ApplQueueResolution)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataIncrementalRefresh) MDBookType() (*field.MDBookType, error) {
	f := new(field.MDBookType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataIncrementalRefresh) MDFeedType() (*field.MDFeedType, error) {
	f := new(field.MDFeedType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataIncrementalRefresh) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataIncrementalRefresh) NoRoutingIDs() (*field.NoRoutingIDs, error) {
	f := new(field.NoRoutingIDs)
	err := m.Body.Get(f)
	return f, err
}
