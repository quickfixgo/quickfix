//Package marketdataincrementalrefresh msg type = X.
package marketdataincrementalrefresh

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a MarketDataIncrementalRefresh wrapper for the generic Message type
type Message struct {
	message.Message
}

//MDReqID is a non-required field for MarketDataIncrementalRefresh.
func (m Message) MDReqID() (*field.MDReqIDField, errors.MessageRejectError) {
	f := &field.MDReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqID reads a MDReqID from MarketDataIncrementalRefresh.
func (m Message) GetMDReqID(f *field.MDReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMDEntries is a required field for MarketDataIncrementalRefresh.
func (m Message) NoMDEntries() (*field.NoMDEntriesField, errors.MessageRejectError) {
	f := &field.NoMDEntriesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoMDEntries reads a NoMDEntries from MarketDataIncrementalRefresh.
func (m Message) GetNoMDEntries(f *field.NoMDEntriesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueDepth is a non-required field for MarketDataIncrementalRefresh.
func (m Message) ApplQueueDepth() (*field.ApplQueueDepthField, errors.MessageRejectError) {
	f := &field.ApplQueueDepthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueDepth reads a ApplQueueDepth from MarketDataIncrementalRefresh.
func (m Message) GetApplQueueDepth(f *field.ApplQueueDepthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueResolution is a non-required field for MarketDataIncrementalRefresh.
func (m Message) ApplQueueResolution() (*field.ApplQueueResolutionField, errors.MessageRejectError) {
	f := &field.ApplQueueResolutionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueResolution reads a ApplQueueResolution from MarketDataIncrementalRefresh.
func (m Message) GetApplQueueResolution(f *field.ApplQueueResolutionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds MarketDataIncrementalRefresh messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for MarketDataIncrementalRefresh.
func Builder(
	nomdentries *field.NoMDEntriesField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header().Set(field.NewMsgType("X"))
	builder.Body().Set(nomdentries)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX44, "X", r
}
