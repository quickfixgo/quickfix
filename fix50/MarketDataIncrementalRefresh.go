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
	nomdentries *field.NoMDEntriesField) MarketDataIncrementalRefreshBuilder {
	var builder MarketDataIncrementalRefreshBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header().Set(field.NewMsgType("X"))
	builder.Body().Set(nomdentries)
	return builder
}

//MDReqID is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) MDReqID() (*field.MDReqIDField, errors.MessageRejectError) {
	f := &field.MDReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqID reads a MDReqID from MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) GetMDReqID(f *field.MDReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMDEntries is a required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) NoMDEntries() (*field.NoMDEntriesField, errors.MessageRejectError) {
	f := &field.NoMDEntriesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoMDEntries reads a NoMDEntries from MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) GetNoMDEntries(f *field.NoMDEntriesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueDepth is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) ApplQueueDepth() (*field.ApplQueueDepthField, errors.MessageRejectError) {
	f := &field.ApplQueueDepthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueDepth reads a ApplQueueDepth from MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) GetApplQueueDepth(f *field.ApplQueueDepthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueResolution is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) ApplQueueResolution() (*field.ApplQueueResolutionField, errors.MessageRejectError) {
	f := &field.ApplQueueResolutionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueResolution reads a ApplQueueResolution from MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) GetApplQueueResolution(f *field.ApplQueueResolutionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDBookType is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) MDBookType() (*field.MDBookTypeField, errors.MessageRejectError) {
	f := &field.MDBookTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDBookType reads a MDBookType from MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) GetMDBookType(f *field.MDBookTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDFeedType is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) MDFeedType() (*field.MDFeedTypeField, errors.MessageRejectError) {
	f := &field.MDFeedTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDFeedType reads a MDFeedType from MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) GetMDFeedType(f *field.MDFeedTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRoutingIDs is a non-required field for MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) NoRoutingIDs() (*field.NoRoutingIDsField, errors.MessageRejectError) {
	f := &field.NoRoutingIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRoutingIDs reads a NoRoutingIDs from MarketDataIncrementalRefresh.
func (m MarketDataIncrementalRefresh) GetNoRoutingIDs(f *field.NoRoutingIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
