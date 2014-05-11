package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	nobidcomponents *field.NoBidComponentsField) BidResponseBuilder {
	var builder BidResponseBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header().Set(field.NewMsgType("l"))
	builder.Body().Set(nobidcomponents)
	return builder
}

//BidID is a non-required field for BidResponse.
func (m BidResponse) BidID() (*field.BidIDField, errors.MessageRejectError) {
	f := &field.BidIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidID reads a BidID from BidResponse.
func (m BidResponse) GetBidID(f *field.BidIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientBidID is a non-required field for BidResponse.
func (m BidResponse) ClientBidID() (*field.ClientBidIDField, errors.MessageRejectError) {
	f := &field.ClientBidIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientBidID reads a ClientBidID from BidResponse.
func (m BidResponse) GetClientBidID(f *field.ClientBidIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoBidComponents is a required field for BidResponse.
func (m BidResponse) NoBidComponents() (*field.NoBidComponentsField, errors.MessageRejectError) {
	f := &field.NoBidComponentsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoBidComponents reads a NoBidComponents from BidResponse.
func (m BidResponse) GetNoBidComponents(f *field.NoBidComponentsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
