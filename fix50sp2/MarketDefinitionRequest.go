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

//NewMarketDefinitionRequestBuilder returns an initialized MarketDefinitionRequestBuilder with specified required fields.
func NewMarketDefinitionRequestBuilder(
	marketreqid field.MarketReqID,
	subscriptionrequesttype field.SubscriptionRequestType) *MarketDefinitionRequestBuilder {
	builder := new(MarketDefinitionRequestBuilder)
	builder.Body.Set(marketreqid)
	builder.Body.Set(subscriptionrequesttype)
	return builder
}

//MarketReqID is a required field for MarketDefinitionRequest.
func (m *MarketDefinitionRequest) MarketReqID() (*field.MarketReqID, error) {
	f := new(field.MarketReqID)
	err := m.Body.Get(f)
	return f, err
}

//SubscriptionRequestType is a required field for MarketDefinitionRequest.
func (m *MarketDefinitionRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//MarketID is a non-required field for MarketDefinitionRequest.
func (m *MarketDefinitionRequest) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//MarketSegmentID is a non-required field for MarketDefinitionRequest.
func (m *MarketDefinitionRequest) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}

//ParentMktSegmID is a non-required field for MarketDefinitionRequest.
func (m *MarketDefinitionRequest) ParentMktSegmID() (*field.ParentMktSegmID, error) {
	f := new(field.ParentMktSegmID)
	err := m.Body.Get(f)
	return f, err
}
