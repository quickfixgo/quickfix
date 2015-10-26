//Package marketdataincrementalrefresh msg type = X.
package marketdataincrementalrefresh

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a MarketDataIncrementalRefresh wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//MDReqID is a non-required field for MarketDataIncrementalRefresh.
func (m Message) MDReqID() (*field.MDReqIDField, quickfix.MessageRejectError) {
	f := &field.MDReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqID reads a MDReqID from MarketDataIncrementalRefresh.
func (m Message) GetMDReqID(f *field.MDReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoMDEntries is a required field for MarketDataIncrementalRefresh.
func (m Message) NoMDEntries() (*field.NoMDEntriesField, quickfix.MessageRejectError) {
	f := &field.NoMDEntriesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoMDEntries reads a NoMDEntries from MarketDataIncrementalRefresh.
func (m Message) GetNoMDEntries(f *field.NoMDEntriesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueDepth is a non-required field for MarketDataIncrementalRefresh.
func (m Message) ApplQueueDepth() (*field.ApplQueueDepthField, quickfix.MessageRejectError) {
	f := &field.ApplQueueDepthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueDepth reads a ApplQueueDepth from MarketDataIncrementalRefresh.
func (m Message) GetApplQueueDepth(f *field.ApplQueueDepthField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueResolution is a non-required field for MarketDataIncrementalRefresh.
func (m Message) ApplQueueResolution() (*field.ApplQueueResolutionField, quickfix.MessageRejectError) {
	f := &field.ApplQueueResolutionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueResolution reads a ApplQueueResolution from MarketDataIncrementalRefresh.
func (m Message) GetApplQueueResolution(f *field.ApplQueueResolutionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MDBookType is a non-required field for MarketDataIncrementalRefresh.
func (m Message) MDBookType() (*field.MDBookTypeField, quickfix.MessageRejectError) {
	f := &field.MDBookTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDBookType reads a MDBookType from MarketDataIncrementalRefresh.
func (m Message) GetMDBookType(f *field.MDBookTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MDFeedType is a non-required field for MarketDataIncrementalRefresh.
func (m Message) MDFeedType() (*field.MDFeedTypeField, quickfix.MessageRejectError) {
	f := &field.MDFeedTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDFeedType reads a MDFeedType from MarketDataIncrementalRefresh.
func (m Message) GetMDFeedType(f *field.MDFeedTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for MarketDataIncrementalRefresh.
func (m Message) TradeDate() (*field.TradeDateField, quickfix.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from MarketDataIncrementalRefresh.
func (m Message) GetTradeDate(f *field.TradeDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRoutingIDs is a non-required field for MarketDataIncrementalRefresh.
func (m Message) NoRoutingIDs() (*field.NoRoutingIDsField, quickfix.MessageRejectError) {
	f := &field.NoRoutingIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRoutingIDs reads a NoRoutingIDs from MarketDataIncrementalRefresh.
func (m Message) GetNoRoutingIDs(f *field.NoRoutingIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for MarketDataIncrementalRefresh.
func (m Message) ApplID() (*field.ApplIDField, quickfix.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from MarketDataIncrementalRefresh.
func (m Message) GetApplID(f *field.ApplIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for MarketDataIncrementalRefresh.
func (m Message) ApplSeqNum() (*field.ApplSeqNumField, quickfix.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from MarketDataIncrementalRefresh.
func (m Message) GetApplSeqNum(f *field.ApplSeqNumField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for MarketDataIncrementalRefresh.
func (m Message) ApplLastSeqNum() (*field.ApplLastSeqNumField, quickfix.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from MarketDataIncrementalRefresh.
func (m Message) GetApplLastSeqNum(f *field.ApplLastSeqNumField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for MarketDataIncrementalRefresh.
func (m Message) ApplResendFlag() (*field.ApplResendFlagField, quickfix.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from MarketDataIncrementalRefresh.
func (m Message) GetApplResendFlag(f *field.ApplResendFlagField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for MarketDataIncrementalRefresh.
func New(
	nomdentries *field.NoMDEntriesField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("X"))
	builder.Body.Set(nomdentries)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "X", r
}
