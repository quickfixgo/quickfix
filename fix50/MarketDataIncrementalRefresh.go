package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header.Set(field.BuildMsgType("X"))
	builder.Body.Set(nomdentries)
	return builder
}

//MDReqID is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) MDReqID() (*field.MDReqID, errors.MessageRejectError) {
	f := new(field.MDReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqID reads a MDReqID from MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) GetMDReqID(f *field.MDReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMDEntries is a required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) NoMDEntries() (*field.NoMDEntries, errors.MessageRejectError) {
	f := new(field.NoMDEntries)
	err := m.Body.Get(f)
	return f, err
}

//GetNoMDEntries reads a NoMDEntries from MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) GetNoMDEntries(f *field.NoMDEntries) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueDepth is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) ApplQueueDepth() (*field.ApplQueueDepth, errors.MessageRejectError) {
	f := new(field.ApplQueueDepth)
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueDepth reads a ApplQueueDepth from MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) GetApplQueueDepth(f *field.ApplQueueDepth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueResolution is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) ApplQueueResolution() (*field.ApplQueueResolution, errors.MessageRejectError) {
	f := new(field.ApplQueueResolution)
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueResolution reads a ApplQueueResolution from MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) GetApplQueueResolution(f *field.ApplQueueResolution) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDBookType is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) MDBookType() (*field.MDBookType, errors.MessageRejectError) {
	f := new(field.MDBookType)
	err := m.Body.Get(f)
	return f, err
}

//GetMDBookType reads a MDBookType from MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) GetMDBookType(f *field.MDBookType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDFeedType is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) MDFeedType() (*field.MDFeedType, errors.MessageRejectError) {
	f := new(field.MDFeedType)
	err := m.Body.Get(f)
	return f, err
}

//GetMDFeedType reads a MDFeedType from MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) GetMDFeedType(f *field.MDFeedType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRoutingIDs is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) NoRoutingIDs() (*field.NoRoutingIDs, errors.MessageRejectError) {
	f := new(field.NoRoutingIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRoutingIDs reads a NoRoutingIDs from MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) GetNoRoutingIDs(f *field.NoRoutingIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}
