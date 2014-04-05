package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type MarketDataRequest struct {
	quickfixgo.Message
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
func (m *MarketDataRequest) ApplQueueMax() (*field.ApplQueueMax, error) {
	f := new(field.ApplQueueMax)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) MDQuoteType() (*field.MDQuoteType, error) {
	f := new(field.MDQuoteType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) MDUpdateType() (*field.MDUpdateType, error) {
	f := new(field.MDUpdateType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) NoMDEntryTypes() (*field.NoMDEntryTypes, error) {
	f := new(field.NoMDEntryTypes)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) NoTradingSessions() (*field.NoTradingSessions, error) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) ApplQueueAction() (*field.ApplQueueAction, error) {
	f := new(field.ApplQueueAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) Scope() (*field.Scope, error) {
	f := new(field.Scope)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) MDImplicitDelete() (*field.MDImplicitDelete, error) {
	f := new(field.MDImplicitDelete)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) MDReqID() (*field.MDReqID, error) {
	f := new(field.MDReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) AggregatedBook() (*field.AggregatedBook, error) {
	f := new(field.AggregatedBook)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataRequest) OpenCloseSettlFlag() (*field.OpenCloseSettlFlag, error) {
	f := new(field.OpenCloseSettlFlag)
	err := m.Body.Get(f)
	return f, err
}
