package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
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
func (m BidResponse) BidID() (*field.BidID, errors.MessageRejectError) {
	f := new(field.BidID)
	err := m.Body.Get(f)
	return f, err
}

//GetBidID reads a BidID from BidResponse.
func (m BidResponse) GetBidID(f *field.BidID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientBidID is a non-required field for BidResponse.
func (m BidResponse) ClientBidID() (*field.ClientBidID, errors.MessageRejectError) {
	f := new(field.ClientBidID)
	err := m.Body.Get(f)
	return f, err
}

//GetClientBidID reads a ClientBidID from BidResponse.
func (m BidResponse) GetClientBidID(f *field.ClientBidID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoBidComponents is a required field for BidResponse.
func (m BidResponse) NoBidComponents() (*field.NoBidComponents, errors.MessageRejectError) {
	f := new(field.NoBidComponents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoBidComponents reads a NoBidComponents from BidResponse.
func (m BidResponse) GetNoBidComponents(f *field.NoBidComponents) errors.MessageRejectError {
	return m.Body.Get(f)
}
