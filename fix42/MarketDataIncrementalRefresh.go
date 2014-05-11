package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	nomdentries *field.NoMDEntriesField) MarketDataIncrementalRefreshBuilder {
	var builder MarketDataIncrementalRefreshBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX42))
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
