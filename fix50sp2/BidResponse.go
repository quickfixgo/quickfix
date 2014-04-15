package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//BidResponse msg type = l.
type BidResponse struct {
	message.Message
}

//BidResponseBuilder builds BidResponse messages.
type BidResponseBuilder struct {
	message.MessageBuilder
}

//CreateBidResponseBuilder returns an initialized BidResponseBuilder with specified required fields.
func CreateBidResponseBuilder(
	nobidcomponents field.NoBidComponents) BidResponseBuilder {
	var builder BidResponseBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(nobidcomponents)
	return builder
}

//BidID is a non-required field for BidResponse.
func (m BidResponse) BidID() (field.BidID, error) {
	var f field.BidID
	err := m.Body.Get(&f)
	return f, err
}

//ClientBidID is a non-required field for BidResponse.
func (m BidResponse) ClientBidID() (field.ClientBidID, error) {
	var f field.ClientBidID
	err := m.Body.Get(&f)
	return f, err
}

//NoBidComponents is a required field for BidResponse.
func (m BidResponse) NoBidComponents() (field.NoBidComponents, error) {
	var f field.NoBidComponents
	err := m.Body.Get(&f)
	return f, err
}
