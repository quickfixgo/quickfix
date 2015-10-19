//Package marketdataincrementalrefresh msg type = X.
package marketdataincrementalrefresh

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
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

//MessageBuilder builds MarketDataIncrementalRefresh messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for MarketDataIncrementalRefresh.
func Builder(
	nomdentries *field.NoMDEntriesField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX42))
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
	return fix.BeginString_FIX42, "X", r
}
