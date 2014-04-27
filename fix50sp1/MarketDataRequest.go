package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
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
	mdreqid field.MDReqID,
	subscriptionrequesttype field.SubscriptionRequestType,
	marketdepth field.MarketDepth,
	nomdentrytypes field.NoMDEntryTypes,
	norelatedsym field.NoRelatedSym) MarketDataRequestBuilder {
	var builder MarketDataRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildMsgType("V"))
	builder.Body.Set(mdreqid)
	builder.Body.Set(subscriptionrequesttype)
	builder.Body.Set(marketdepth)
	builder.Body.Set(nomdentrytypes)
	builder.Body.Set(norelatedsym)
	return builder
}

//MDReqID is a required field for MarketDataRequest.
func (m MarketDataRequest) MDReqID() (*field.MDReqID, errors.MessageRejectError) {
	f := new(field.MDReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqID reads a MDReqID from MarketDataRequest.
func (m MarketDataRequest) GetMDReqID(f *field.MDReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a required field for MarketDataRequest.
func (m MarketDataRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, errors.MessageRejectError) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from MarketDataRequest.
func (m MarketDataRequest) GetSubscriptionRequestType(f *field.SubscriptionRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketDepth is a required field for MarketDataRequest.
func (m MarketDataRequest) MarketDepth() (*field.MarketDepth, errors.MessageRejectError) {
	f := new(field.MarketDepth)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketDepth reads a MarketDepth from MarketDataRequest.
func (m MarketDataRequest) GetMarketDepth(f *field.MarketDepth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDUpdateType is a non-required field for MarketDataRequest.
func (m MarketDataRequest) MDUpdateType() (*field.MDUpdateType, errors.MessageRejectError) {
	f := new(field.MDUpdateType)
	err := m.Body.Get(f)
	return f, err
}

//GetMDUpdateType reads a MDUpdateType from MarketDataRequest.
func (m MarketDataRequest) GetMDUpdateType(f *field.MDUpdateType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AggregatedBook is a non-required field for MarketDataRequest.
func (m MarketDataRequest) AggregatedBook() (*field.AggregatedBook, errors.MessageRejectError) {
	f := new(field.AggregatedBook)
	err := m.Body.Get(f)
	return f, err
}

//GetAggregatedBook reads a AggregatedBook from MarketDataRequest.
func (m MarketDataRequest) GetAggregatedBook(f *field.AggregatedBook) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OpenCloseSettlFlag is a non-required field for MarketDataRequest.
func (m MarketDataRequest) OpenCloseSettlFlag() (*field.OpenCloseSettlFlag, errors.MessageRejectError) {
	f := new(field.OpenCloseSettlFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetOpenCloseSettlFlag reads a OpenCloseSettlFlag from MarketDataRequest.
func (m MarketDataRequest) GetOpenCloseSettlFlag(f *field.OpenCloseSettlFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Scope is a non-required field for MarketDataRequest.
func (m MarketDataRequest) Scope() (*field.Scope, errors.MessageRejectError) {
	f := new(field.Scope)
	err := m.Body.Get(f)
	return f, err
}

//GetScope reads a Scope from MarketDataRequest.
func (m MarketDataRequest) GetScope(f *field.Scope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDImplicitDelete is a non-required field for MarketDataRequest.
func (m MarketDataRequest) MDImplicitDelete() (*field.MDImplicitDelete, errors.MessageRejectError) {
	f := new(field.MDImplicitDelete)
	err := m.Body.Get(f)
	return f, err
}

//GetMDImplicitDelete reads a MDImplicitDelete from MarketDataRequest.
func (m MarketDataRequest) GetMDImplicitDelete(f *field.MDImplicitDelete) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMDEntryTypes is a required field for MarketDataRequest.
func (m MarketDataRequest) NoMDEntryTypes() (*field.NoMDEntryTypes, errors.MessageRejectError) {
	f := new(field.NoMDEntryTypes)
	err := m.Body.Get(f)
	return f, err
}

//GetNoMDEntryTypes reads a NoMDEntryTypes from MarketDataRequest.
func (m MarketDataRequest) GetNoMDEntryTypes(f *field.NoMDEntryTypes) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a required field for MarketDataRequest.
func (m MarketDataRequest) NoRelatedSym() (*field.NoRelatedSym, errors.MessageRejectError) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from MarketDataRequest.
func (m MarketDataRequest) GetNoRelatedSym(f *field.NoRelatedSym) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for MarketDataRequest.
func (m MarketDataRequest) NoTradingSessions() (*field.NoTradingSessions, errors.MessageRejectError) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from MarketDataRequest.
func (m MarketDataRequest) GetNoTradingSessions(f *field.NoTradingSessions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueAction is a non-required field for MarketDataRequest.
func (m MarketDataRequest) ApplQueueAction() (*field.ApplQueueAction, errors.MessageRejectError) {
	f := new(field.ApplQueueAction)
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueAction reads a ApplQueueAction from MarketDataRequest.
func (m MarketDataRequest) GetApplQueueAction(f *field.ApplQueueAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueMax is a non-required field for MarketDataRequest.
func (m MarketDataRequest) ApplQueueMax() (*field.ApplQueueMax, errors.MessageRejectError) {
	f := new(field.ApplQueueMax)
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueMax reads a ApplQueueMax from MarketDataRequest.
func (m MarketDataRequest) GetApplQueueMax(f *field.ApplQueueMax) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDQuoteType is a non-required field for MarketDataRequest.
func (m MarketDataRequest) MDQuoteType() (*field.MDQuoteType, errors.MessageRejectError) {
	f := new(field.MDQuoteType)
	err := m.Body.Get(f)
	return f, err
}

//GetMDQuoteType reads a MDQuoteType from MarketDataRequest.
func (m MarketDataRequest) GetMDQuoteType(f *field.MDQuoteType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for MarketDataRequest.
func (m MarketDataRequest) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from MarketDataRequest.
func (m MarketDataRequest) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}
