package fix50

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

//CreateMarketDataIncrementalRefreshBuilder returns an initialized MarketDataIncrementalRefreshBuilder with specified required fields.
func CreateMarketDataIncrementalRefreshBuilder(
	nomdentries field.NoMDEntries) MarketDataIncrementalRefreshBuilder {
	var builder MarketDataIncrementalRefreshBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(nomdentries)
	return builder
}

//MDReqID is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) MDReqID() (field.MDReqID, error) {
	var f field.MDReqID
	err := m.Body.Get(&f)
	return f, err
}

//NoMDEntries is a required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) NoMDEntries() (field.NoMDEntries, error) {
	var f field.NoMDEntries
	err := m.Body.Get(&f)
	return f, err
}

//ApplQueueDepth is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) ApplQueueDepth() (field.ApplQueueDepth, error) {
	var f field.ApplQueueDepth
	err := m.Body.Get(&f)
	return f, err
}

//ApplQueueResolution is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) ApplQueueResolution() (field.ApplQueueResolution, error) {
	var f field.ApplQueueResolution
	err := m.Body.Get(&f)
	return f, err
}

//MDBookType is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) MDBookType() (field.MDBookType, error) {
	var f field.MDBookType
	err := m.Body.Get(&f)
	return f, err
}

//MDFeedType is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) MDFeedType() (field.MDFeedType, error) {
	var f field.MDFeedType
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//NoRoutingIDs is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) NoRoutingIDs() (field.NoRoutingIDs, error) {
	var f field.NoRoutingIDs
	err := m.Body.Get(&f)
	return f, err
}
