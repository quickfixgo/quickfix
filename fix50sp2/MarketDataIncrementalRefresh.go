package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//MarketDataIncrementalRefresh msg type = X.
type MarketDataIncrementalRefresh struct {
	message.Message
}

//MarketDataIncrementalRefreshBuilder builds MarketDataIncrementalRefresh messages.
type MarketDataIncrementalRefreshBuilder struct {
	message.MessageBuilder
}

//NewMarketDataIncrementalRefreshBuilder returns an initialized MarketDataIncrementalRefreshBuilder with specified required fields.
func NewMarketDataIncrementalRefreshBuilder(
	nomdentries field.NoMDEntries) *MarketDataIncrementalRefreshBuilder {
	builder := new(MarketDataIncrementalRefreshBuilder)
	builder.Body.Set(nomdentries)
	return builder
}

//MDReqID is a non-required field for MarketDataIncrementalRefresh.
func (m *MarketDataIncrementalRefresh) MDReqID() (*field.MDReqID, error) {
	f := new(field.MDReqID)
	err := m.Body.Get(f)
	return f, err
}

//NoMDEntries is a required field for MarketDataIncrementalRefresh.
func (m *MarketDataIncrementalRefresh) NoMDEntries() (*field.NoMDEntries, error) {
	f := new(field.NoMDEntries)
	err := m.Body.Get(f)
	return f, err
}

//ApplQueueDepth is a non-required field for MarketDataIncrementalRefresh.
func (m *MarketDataIncrementalRefresh) ApplQueueDepth() (*field.ApplQueueDepth, error) {
	f := new(field.ApplQueueDepth)
	err := m.Body.Get(f)
	return f, err
}

//ApplQueueResolution is a non-required field for MarketDataIncrementalRefresh.
func (m *MarketDataIncrementalRefresh) ApplQueueResolution() (*field.ApplQueueResolution, error) {
	f := new(field.ApplQueueResolution)
	err := m.Body.Get(f)
	return f, err
}

//MDBookType is a non-required field for MarketDataIncrementalRefresh.
func (m *MarketDataIncrementalRefresh) MDBookType() (*field.MDBookType, error) {
	f := new(field.MDBookType)
	err := m.Body.Get(f)
	return f, err
}

//MDFeedType is a non-required field for MarketDataIncrementalRefresh.
func (m *MarketDataIncrementalRefresh) MDFeedType() (*field.MDFeedType, error) {
	f := new(field.MDFeedType)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a non-required field for MarketDataIncrementalRefresh.
func (m *MarketDataIncrementalRefresh) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//NoRoutingIDs is a non-required field for MarketDataIncrementalRefresh.
func (m *MarketDataIncrementalRefresh) NoRoutingIDs() (*field.NoRoutingIDs, error) {
	f := new(field.NoRoutingIDs)
	err := m.Body.Get(f)
	return f, err
}

//ApplID is a non-required field for MarketDataIncrementalRefresh.
func (m *MarketDataIncrementalRefresh) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//ApplSeqNum is a non-required field for MarketDataIncrementalRefresh.
func (m *MarketDataIncrementalRefresh) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplLastSeqNum is a non-required field for MarketDataIncrementalRefresh.
func (m *MarketDataIncrementalRefresh) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplResendFlag is a non-required field for MarketDataIncrementalRefresh.
func (m *MarketDataIncrementalRefresh) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}
