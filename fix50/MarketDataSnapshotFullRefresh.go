package fix50

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type MarketDataSnapshotFullRefresh struct {
	quickfixgo.Message
}

func (m *MarketDataSnapshotFullRefresh) MDReqID() (*field.MDReqID, error) {
	f := new(field.MDReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) FinancialStatus() (*field.FinancialStatus, error) {
	f := new(field.FinancialStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) CorporateAction() (*field.CorporateAction, error) {
	f := new(field.CorporateAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) NetChgPrevDay() (*field.NetChgPrevDay, error) {
	f := new(field.NetChgPrevDay)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) ApplQueueDepth() (*field.ApplQueueDepth, error) {
	f := new(field.ApplQueueDepth)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) ApplQueueResolution() (*field.ApplQueueResolution, error) {
	f := new(field.ApplQueueResolution)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) MDReportID() (*field.MDReportID, error) {
	f := new(field.MDReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) MDBookType() (*field.MDBookType, error) {
	f := new(field.MDBookType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) MDFeedType() (*field.MDFeedType, error) {
	f := new(field.MDFeedType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
