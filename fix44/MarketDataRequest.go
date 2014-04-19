package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
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
	builder.Body.Set(mdreqid)
	builder.Body.Set(subscriptionrequesttype)
	builder.Body.Set(marketdepth)
	builder.Body.Set(nomdentrytypes)
	builder.Body.Set(norelatedsym)
	return builder
}

//MDReqID is a required field for MarketDataRequest.
func (m MarketDataRequest) MDReqID() (field.MDReqID, errors.MessageRejectError) {
	var f field.MDReqID
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a required field for MarketDataRequest.
func (m MarketDataRequest) SubscriptionRequestType() (field.SubscriptionRequestType, errors.MessageRejectError) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//MarketDepth is a required field for MarketDataRequest.
func (m MarketDataRequest) MarketDepth() (field.MarketDepth, errors.MessageRejectError) {
	var f field.MarketDepth
	err := m.Body.Get(&f)
	return f, err
}

//MDUpdateType is a non-required field for MarketDataRequest.
func (m MarketDataRequest) MDUpdateType() (field.MDUpdateType, errors.MessageRejectError) {
	var f field.MDUpdateType
	err := m.Body.Get(&f)
	return f, err
}

//AggregatedBook is a non-required field for MarketDataRequest.
func (m MarketDataRequest) AggregatedBook() (field.AggregatedBook, errors.MessageRejectError) {
	var f field.AggregatedBook
	err := m.Body.Get(&f)
	return f, err
}

//OpenCloseSettlFlag is a non-required field for MarketDataRequest.
func (m MarketDataRequest) OpenCloseSettlFlag() (field.OpenCloseSettlFlag, errors.MessageRejectError) {
	var f field.OpenCloseSettlFlag
	err := m.Body.Get(&f)
	return f, err
}

//Scope is a non-required field for MarketDataRequest.
func (m MarketDataRequest) Scope() (field.Scope, errors.MessageRejectError) {
	var f field.Scope
	err := m.Body.Get(&f)
	return f, err
}

//MDImplicitDelete is a non-required field for MarketDataRequest.
func (m MarketDataRequest) MDImplicitDelete() (field.MDImplicitDelete, errors.MessageRejectError) {
	var f field.MDImplicitDelete
	err := m.Body.Get(&f)
	return f, err
}

//NoMDEntryTypes is a required field for MarketDataRequest.
func (m MarketDataRequest) NoMDEntryTypes() (field.NoMDEntryTypes, errors.MessageRejectError) {
	var f field.NoMDEntryTypes
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a required field for MarketDataRequest.
func (m MarketDataRequest) NoRelatedSym() (field.NoRelatedSym, errors.MessageRejectError) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}
