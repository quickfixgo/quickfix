//Package marketdatarequest msg type = V.
package marketdatarequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a MarketDataRequest wrapper for the generic Message type
type Message struct {
	message.Message
}

//MDReqID is a required field for MarketDataRequest.
func (m Message) MDReqID() (*field.MDReqIDField, errors.MessageRejectError) {
	f := &field.MDReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqID reads a MDReqID from MarketDataRequest.
func (m Message) GetMDReqID(f *field.MDReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a required field for MarketDataRequest.
func (m Message) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, errors.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from MarketDataRequest.
func (m Message) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketDepth is a required field for MarketDataRequest.
func (m Message) MarketDepth() (*field.MarketDepthField, errors.MessageRejectError) {
	f := &field.MarketDepthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketDepth reads a MarketDepth from MarketDataRequest.
func (m Message) GetMarketDepth(f *field.MarketDepthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDUpdateType is a non-required field for MarketDataRequest.
func (m Message) MDUpdateType() (*field.MDUpdateTypeField, errors.MessageRejectError) {
	f := &field.MDUpdateTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDUpdateType reads a MDUpdateType from MarketDataRequest.
func (m Message) GetMDUpdateType(f *field.MDUpdateTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AggregatedBook is a non-required field for MarketDataRequest.
func (m Message) AggregatedBook() (*field.AggregatedBookField, errors.MessageRejectError) {
	f := &field.AggregatedBookField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAggregatedBook reads a AggregatedBook from MarketDataRequest.
func (m Message) GetAggregatedBook(f *field.AggregatedBookField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OpenCloseSettlFlag is a non-required field for MarketDataRequest.
func (m Message) OpenCloseSettlFlag() (*field.OpenCloseSettlFlagField, errors.MessageRejectError) {
	f := &field.OpenCloseSettlFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOpenCloseSettlFlag reads a OpenCloseSettlFlag from MarketDataRequest.
func (m Message) GetOpenCloseSettlFlag(f *field.OpenCloseSettlFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Scope is a non-required field for MarketDataRequest.
func (m Message) Scope() (*field.ScopeField, errors.MessageRejectError) {
	f := &field.ScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetScope reads a Scope from MarketDataRequest.
func (m Message) GetScope(f *field.ScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDImplicitDelete is a non-required field for MarketDataRequest.
func (m Message) MDImplicitDelete() (*field.MDImplicitDeleteField, errors.MessageRejectError) {
	f := &field.MDImplicitDeleteField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDImplicitDelete reads a MDImplicitDelete from MarketDataRequest.
func (m Message) GetMDImplicitDelete(f *field.MDImplicitDeleteField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMDEntryTypes is a required field for MarketDataRequest.
func (m Message) NoMDEntryTypes() (*field.NoMDEntryTypesField, errors.MessageRejectError) {
	f := &field.NoMDEntryTypesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoMDEntryTypes reads a NoMDEntryTypes from MarketDataRequest.
func (m Message) GetNoMDEntryTypes(f *field.NoMDEntryTypesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a required field for MarketDataRequest.
func (m Message) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from MarketDataRequest.
func (m Message) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds MarketDataRequest messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for MarketDataRequest.
func Builder(
	mdreqid *field.MDReqIDField,
	subscriptionrequesttype *field.SubscriptionRequestTypeField,
	marketdepth *field.MarketDepthField,
	nomdentrytypes *field.NoMDEntryTypesField,
	norelatedsym *field.NoRelatedSymField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header().Set(field.NewMsgType("V"))
	builder.Body().Set(mdreqid)
	builder.Body().Set(subscriptionrequesttype)
	builder.Body().Set(marketdepth)
	builder.Body().Set(nomdentrytypes)
	builder.Body().Set(norelatedsym)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX44, "V", r
}
