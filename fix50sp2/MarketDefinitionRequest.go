package fix50sp2

import (
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
func (m MarketDefinitionRequest) MarketReqID() (field.MarketReqID, error) {
	var f field.MarketReqID
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a required field for MarketDefinitionRequest.
func (m MarketDefinitionRequest) SubscriptionRequestType() (field.SubscriptionRequestType, error) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//MarketID is a non-required field for MarketDefinitionRequest.
func (m MarketDefinitionRequest) MarketID() (field.MarketID, error) {
	var f field.MarketID
	err := m.Body.Get(&f)
	return f, err
}

//MarketSegmentID is a non-required field for MarketDefinitionRequest.
func (m MarketDefinitionRequest) MarketSegmentID() (field.MarketSegmentID, error) {
	var f field.MarketSegmentID
	err := m.Body.Get(&f)
	return f, err
}

//ParentMktSegmID is a non-required field for MarketDefinitionRequest.
func (m MarketDefinitionRequest) ParentMktSegmID() (field.ParentMktSegmID, error) {
	var f field.ParentMktSegmID
	err := m.Body.Get(&f)
	return f, err
}
