package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
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
	marketreqid field.MarketReqID,
	subscriptionrequesttype field.SubscriptionRequestType) MarketDefinitionRequestBuilder {
	var builder MarketDefinitionRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(marketreqid)
	builder.Body.Set(subscriptionrequesttype)
	return builder
}

//MarketReqID is a required field for MarketDefinitionRequest.
func (m MarketDefinitionRequest) MarketReqID() (*field.MarketReqID, errors.MessageRejectError) {
	f := new(field.MarketReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketReqID reads a MarketReqID from MarketDefinitionRequest.
func (m MarketDefinitionRequest) GetMarketReqID(f *field.MarketReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a required field for MarketDefinitionRequest.
func (m MarketDefinitionRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, errors.MessageRejectError) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from MarketDefinitionRequest.
func (m MarketDefinitionRequest) GetSubscriptionRequestType(f *field.SubscriptionRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for MarketDefinitionRequest.
func (m MarketDefinitionRequest) MarketID() (*field.MarketID, errors.MessageRejectError) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from MarketDefinitionRequest.
func (m MarketDefinitionRequest) GetMarketID(f *field.MarketID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for MarketDefinitionRequest.
func (m MarketDefinitionRequest) MarketSegmentID() (*field.MarketSegmentID, errors.MessageRejectError) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from MarketDefinitionRequest.
func (m MarketDefinitionRequest) GetMarketSegmentID(f *field.MarketSegmentID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParentMktSegmID is a non-required field for MarketDefinitionRequest.
func (m MarketDefinitionRequest) ParentMktSegmID() (*field.ParentMktSegmID, errors.MessageRejectError) {
	f := new(field.ParentMktSegmID)
	err := m.Body.Get(f)
	return f, err
}

//GetParentMktSegmID reads a ParentMktSegmID from MarketDefinitionRequest.
func (m MarketDefinitionRequest) GetParentMktSegmID(f *field.ParentMktSegmID) errors.MessageRejectError {
	return m.Body.Get(f)
}
