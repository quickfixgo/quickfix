package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//MarketDataRequest msg type = V.
type MarketDataRequest struct {
	message.Message
}

//MarketDataRequestBuilder builds MarketDataRequest messages.
type MarketDataRequestBuilder struct {
	message.MessageBuilder
}

//CreateMarketDataRequestBuilder returns an initialized MarketDataRequestBuilder with specified required fields.
func CreateMarketDataRequestBuilder(
	mdreqid *field.MDReqIDField,
	subscriptionrequesttype *field.SubscriptionRequestTypeField,
	marketdepth *field.MarketDepthField,
	nomdentrytypes *field.NoMDEntryTypesField,
	norelatedsym *field.NoRelatedSymField) MarketDataRequestBuilder {
	var builder MarketDataRequestBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("V"))
	builder.Body.Set(mdreqid)
	builder.Body.Set(subscriptionrequesttype)
	builder.Body.Set(marketdepth)
	builder.Body.Set(nomdentrytypes)
	builder.Body.Set(norelatedsym)
	return builder
}

//MDReqID is a required field for MarketDataRequest.
func (m MarketDataRequest) MDReqID() (*field.MDReqIDField, errors.MessageRejectError) {
	f := &field.MDReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqID reads a MDReqID from MarketDataRequest.
func (m MarketDataRequest) GetMDReqID(f *field.MDReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a required field for MarketDataRequest.
func (m MarketDataRequest) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, errors.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from MarketDataRequest.
func (m MarketDataRequest) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketDepth is a required field for MarketDataRequest.
func (m MarketDataRequest) MarketDepth() (*field.MarketDepthField, errors.MessageRejectError) {
	f := &field.MarketDepthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketDepth reads a MarketDepth from MarketDataRequest.
func (m MarketDataRequest) GetMarketDepth(f *field.MarketDepthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDUpdateType is a non-required field for MarketDataRequest.
func (m MarketDataRequest) MDUpdateType() (*field.MDUpdateTypeField, errors.MessageRejectError) {
	f := &field.MDUpdateTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDUpdateType reads a MDUpdateType from MarketDataRequest.
func (m MarketDataRequest) GetMDUpdateType(f *field.MDUpdateTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AggregatedBook is a non-required field for MarketDataRequest.
func (m MarketDataRequest) AggregatedBook() (*field.AggregatedBookField, errors.MessageRejectError) {
	f := &field.AggregatedBookField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAggregatedBook reads a AggregatedBook from MarketDataRequest.
func (m MarketDataRequest) GetAggregatedBook(f *field.AggregatedBookField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OpenCloseSettlFlag is a non-required field for MarketDataRequest.
func (m MarketDataRequest) OpenCloseSettlFlag() (*field.OpenCloseSettlFlagField, errors.MessageRejectError) {
	f := &field.OpenCloseSettlFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOpenCloseSettlFlag reads a OpenCloseSettlFlag from MarketDataRequest.
func (m MarketDataRequest) GetOpenCloseSettlFlag(f *field.OpenCloseSettlFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Scope is a non-required field for MarketDataRequest.
func (m MarketDataRequest) Scope() (*field.ScopeField, errors.MessageRejectError) {
	f := &field.ScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetScope reads a Scope from MarketDataRequest.
func (m MarketDataRequest) GetScope(f *field.ScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDImplicitDelete is a non-required field for MarketDataRequest.
func (m MarketDataRequest) MDImplicitDelete() (*field.MDImplicitDeleteField, errors.MessageRejectError) {
	f := &field.MDImplicitDeleteField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDImplicitDelete reads a MDImplicitDelete from MarketDataRequest.
func (m MarketDataRequest) GetMDImplicitDelete(f *field.MDImplicitDeleteField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMDEntryTypes is a required field for MarketDataRequest.
func (m MarketDataRequest) NoMDEntryTypes() (*field.NoMDEntryTypesField, errors.MessageRejectError) {
	f := &field.NoMDEntryTypesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoMDEntryTypes reads a NoMDEntryTypes from MarketDataRequest.
func (m MarketDataRequest) GetNoMDEntryTypes(f *field.NoMDEntryTypesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a required field for MarketDataRequest.
func (m MarketDataRequest) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from MarketDataRequest.
func (m MarketDataRequest) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for MarketDataRequest.
func (m MarketDataRequest) NoTradingSessions() (*field.NoTradingSessionsField, errors.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from MarketDataRequest.
func (m MarketDataRequest) GetNoTradingSessions(f *field.NoTradingSessionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueAction is a non-required field for MarketDataRequest.
func (m MarketDataRequest) ApplQueueAction() (*field.ApplQueueActionField, errors.MessageRejectError) {
	f := &field.ApplQueueActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueAction reads a ApplQueueAction from MarketDataRequest.
func (m MarketDataRequest) GetApplQueueAction(f *field.ApplQueueActionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueMax is a non-required field for MarketDataRequest.
func (m MarketDataRequest) ApplQueueMax() (*field.ApplQueueMaxField, errors.MessageRejectError) {
	f := &field.ApplQueueMaxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueMax reads a ApplQueueMax from MarketDataRequest.
func (m MarketDataRequest) GetApplQueueMax(f *field.ApplQueueMaxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDQuoteType is a non-required field for MarketDataRequest.
func (m MarketDataRequest) MDQuoteType() (*field.MDQuoteTypeField, errors.MessageRejectError) {
	f := &field.MDQuoteTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDQuoteType reads a MDQuoteType from MarketDataRequest.
func (m MarketDataRequest) GetMDQuoteType(f *field.MDQuoteTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for MarketDataRequest.
func (m MarketDataRequest) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from MarketDataRequest.
func (m MarketDataRequest) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
