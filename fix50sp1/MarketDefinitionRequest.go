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

//MarketDefinitionRequest msg type = BT.
type MarketDefinitionRequest struct {
	message.Message
}

//MarketDefinitionRequestBuilder builds MarketDefinitionRequest messages.
type MarketDefinitionRequestBuilder struct {
	message.MessageBuilder
}

//CreateMarketDefinitionRequestBuilder returns an initialized MarketDefinitionRequestBuilder with specified required fields.
func CreateMarketDefinitionRequestBuilder(
	marketreqid *field.MarketReqIDField,
	subscriptionrequesttype *field.SubscriptionRequestTypeField) MarketDefinitionRequestBuilder {
	var builder MarketDefinitionRequestBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("BT"))
	builder.Body.Set(marketreqid)
	builder.Body.Set(subscriptionrequesttype)
	return builder
}

//MarketReqID is a required field for MarketDefinitionRequest.
func (m MarketDefinitionRequest) MarketReqID() (*field.MarketReqIDField, errors.MessageRejectError) {
	f := &field.MarketReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketReqID reads a MarketReqID from MarketDefinitionRequest.
func (m MarketDefinitionRequest) GetMarketReqID(f *field.MarketReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a required field for MarketDefinitionRequest.
func (m MarketDefinitionRequest) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, errors.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from MarketDefinitionRequest.
func (m MarketDefinitionRequest) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for MarketDefinitionRequest.
func (m MarketDefinitionRequest) MarketID() (*field.MarketIDField, errors.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from MarketDefinitionRequest.
func (m MarketDefinitionRequest) GetMarketID(f *field.MarketIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for MarketDefinitionRequest.
func (m MarketDefinitionRequest) MarketSegmentID() (*field.MarketSegmentIDField, errors.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from MarketDefinitionRequest.
func (m MarketDefinitionRequest) GetMarketSegmentID(f *field.MarketSegmentIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParentMktSegmID is a non-required field for MarketDefinitionRequest.
func (m MarketDefinitionRequest) ParentMktSegmID() (*field.ParentMktSegmIDField, errors.MessageRejectError) {
	f := &field.ParentMktSegmIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetParentMktSegmID reads a ParentMktSegmID from MarketDefinitionRequest.
func (m MarketDefinitionRequest) GetParentMktSegmID(f *field.ParentMktSegmIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}
