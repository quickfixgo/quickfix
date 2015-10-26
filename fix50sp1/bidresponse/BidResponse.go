//Package bidresponse msg type = l.
package bidresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix/enum"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a BidResponse wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//BidID is a non-required field for BidResponse.
func (m Message) BidID() (*field.BidIDField, quickfix.MessageRejectError) {
	f := &field.BidIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidID reads a BidID from BidResponse.
func (m Message) GetBidID(f *field.BidIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClientBidID is a non-required field for BidResponse.
func (m Message) ClientBidID() (*field.ClientBidIDField, quickfix.MessageRejectError) {
	f := &field.ClientBidIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientBidID reads a ClientBidID from BidResponse.
func (m Message) GetClientBidID(f *field.ClientBidIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoBidComponents is a required field for BidResponse.
func (m Message) NoBidComponents() (*field.NoBidComponentsField, quickfix.MessageRejectError) {
	f := &field.NoBidComponentsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoBidComponents reads a NoBidComponents from BidResponse.
func (m Message) GetNoBidComponents(f *field.NoBidComponentsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for BidResponse.
func New(
	nobidcomponents *field.NoBidComponentsField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("l"))
	builder.Body.Set(nobidcomponents)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "l", r
}
