package fix42

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
