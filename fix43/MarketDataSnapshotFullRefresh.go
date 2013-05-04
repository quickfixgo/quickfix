package fix43

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
func (m *MarketDataSnapshotFullRefresh) TotalVolumeTraded() (*field.TotalVolumeTraded, error) {
	f := new(field.TotalVolumeTraded)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) TotalVolumeTradedDate() (*field.TotalVolumeTradedDate, error) {
	f := new(field.TotalVolumeTradedDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) TotalVolumeTradedTime() (*field.TotalVolumeTradedTime, error) {
	f := new(field.TotalVolumeTradedTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDataSnapshotFullRefresh) NetChgPrevDay() (*field.NetChgPrevDay, error) {
	f := new(field.NetChgPrevDay)
	err := m.Body.Get(f)
	return f, err
}
