package fix50sp1

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

//NewBidResponseBuilder returns an initialized BidResponseBuilder with specified required fields.
func NewBidResponseBuilder(
	nobidcomponents field.NoBidComponents) *BidResponseBuilder {
	builder := new(BidResponseBuilder)
	builder.Body.Set(nobidcomponents)
	return builder
}

//BidID is a non-required field for BidResponse.
func (m *BidResponse) BidID() (*field.BidID, error) {
	f := new(field.BidID)
	err := m.Body.Get(f)
	return f, err
}

//ClientBidID is a non-required field for BidResponse.
func (m *BidResponse) ClientBidID() (*field.ClientBidID, error) {
	f := new(field.ClientBidID)
	err := m.Body.Get(f)
	return f, err
}

//NoBidComponents is a required field for BidResponse.
func (m *BidResponse) NoBidComponents() (*field.NoBidComponents, error) {
	f := new(field.NoBidComponents)
	err := m.Body.Get(f)
	return f, err
}
