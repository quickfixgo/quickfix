//Package marketdatarequest msg type = V.
package marketdatarequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a MarketDataRequest wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//MDReqID is a required field for MarketDataRequest.
func (m Message) MDReqID() (*field.MDReqIDField, quickfix.MessageRejectError) {
	f := &field.MDReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqID reads a MDReqID from MarketDataRequest.
func (m Message) GetMDReqID(f *field.MDReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a required field for MarketDataRequest.
func (m Message) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, quickfix.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from MarketDataRequest.
func (m Message) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarketDepth is a required field for MarketDataRequest.
func (m Message) MarketDepth() (*field.MarketDepthField, quickfix.MessageRejectError) {
	f := &field.MarketDepthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketDepth reads a MarketDepth from MarketDataRequest.
func (m Message) GetMarketDepth(f *field.MarketDepthField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MDUpdateType is a non-required field for MarketDataRequest.
func (m Message) MDUpdateType() (*field.MDUpdateTypeField, quickfix.MessageRejectError) {
	f := &field.MDUpdateTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDUpdateType reads a MDUpdateType from MarketDataRequest.
func (m Message) GetMDUpdateType(f *field.MDUpdateTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AggregatedBook is a non-required field for MarketDataRequest.
func (m Message) AggregatedBook() (*field.AggregatedBookField, quickfix.MessageRejectError) {
	f := &field.AggregatedBookField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAggregatedBook reads a AggregatedBook from MarketDataRequest.
func (m Message) GetAggregatedBook(f *field.AggregatedBookField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OpenCloseSettlFlag is a non-required field for MarketDataRequest.
func (m Message) OpenCloseSettlFlag() (*field.OpenCloseSettlFlagField, quickfix.MessageRejectError) {
	f := &field.OpenCloseSettlFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOpenCloseSettlFlag reads a OpenCloseSettlFlag from MarketDataRequest.
func (m Message) GetOpenCloseSettlFlag(f *field.OpenCloseSettlFlagField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Scope is a non-required field for MarketDataRequest.
func (m Message) Scope() (*field.ScopeField, quickfix.MessageRejectError) {
	f := &field.ScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetScope reads a Scope from MarketDataRequest.
func (m Message) GetScope(f *field.ScopeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MDImplicitDelete is a non-required field for MarketDataRequest.
func (m Message) MDImplicitDelete() (*field.MDImplicitDeleteField, quickfix.MessageRejectError) {
	f := &field.MDImplicitDeleteField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDImplicitDelete reads a MDImplicitDelete from MarketDataRequest.
func (m Message) GetMDImplicitDelete(f *field.MDImplicitDeleteField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoMDEntryTypes is a required field for MarketDataRequest.
func (m Message) NoMDEntryTypes() (*field.NoMDEntryTypesField, quickfix.MessageRejectError) {
	f := &field.NoMDEntryTypesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoMDEntryTypes reads a NoMDEntryTypes from MarketDataRequest.
func (m Message) GetNoMDEntryTypes(f *field.NoMDEntryTypesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a required field for MarketDataRequest.
func (m Message) NoRelatedSym() (*field.NoRelatedSymField, quickfix.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from MarketDataRequest.
func (m Message) GetNoRelatedSym(f *field.NoRelatedSymField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for MarketDataRequest.
func (m Message) NoTradingSessions() (*field.NoTradingSessionsField, quickfix.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from MarketDataRequest.
func (m Message) GetNoTradingSessions(f *field.NoTradingSessionsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueAction is a non-required field for MarketDataRequest.
func (m Message) ApplQueueAction() (*field.ApplQueueActionField, quickfix.MessageRejectError) {
	f := &field.ApplQueueActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueAction reads a ApplQueueAction from MarketDataRequest.
func (m Message) GetApplQueueAction(f *field.ApplQueueActionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueMax is a non-required field for MarketDataRequest.
func (m Message) ApplQueueMax() (*field.ApplQueueMaxField, quickfix.MessageRejectError) {
	f := &field.ApplQueueMaxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueMax reads a ApplQueueMax from MarketDataRequest.
func (m Message) GetApplQueueMax(f *field.ApplQueueMaxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MDQuoteType is a non-required field for MarketDataRequest.
func (m Message) MDQuoteType() (*field.MDQuoteTypeField, quickfix.MessageRejectError) {
	f := &field.MDQuoteTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDQuoteType reads a MDQuoteType from MarketDataRequest.
func (m Message) GetMDQuoteType(f *field.MDQuoteTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for MarketDataRequest.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from MarketDataRequest.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for MarketDataRequest.
func New(
	mdreqid *field.MDReqIDField,
	subscriptionrequesttype *field.SubscriptionRequestTypeField,
	marketdepth *field.MarketDepthField,
	nomdentrytypes *field.NoMDEntryTypesField,
	norelatedsym *field.NoRelatedSymField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("V"))
	builder.Body.Set(mdreqid)
	builder.Body.Set(subscriptionrequesttype)
	builder.Body.Set(marketdepth)
	builder.Body.Set(nomdentrytypes)
	builder.Body.Set(norelatedsym)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "V", r
}
