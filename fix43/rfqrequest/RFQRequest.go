//Package rfqrequest msg type = AH.
package rfqrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a RFQRequest wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//RFQReqID is a required field for RFQRequest.
func (m Message) RFQReqID() (*field.RFQReqIDField, quickfix.MessageRejectError) {
	f := &field.RFQReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRFQReqID reads a RFQReqID from RFQRequest.
func (m Message) GetRFQReqID(f *field.RFQReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a required field for RFQRequest.
func (m Message) NoRelatedSym() (*field.NoRelatedSymField, quickfix.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from RFQRequest.
func (m Message) GetNoRelatedSym(f *field.NoRelatedSymField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for RFQRequest.
func (m Message) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, quickfix.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from RFQRequest.
func (m Message) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for RFQRequest.
func New(
	rfqreqid *field.RFQReqIDField,
	norelatedsym *field.NoRelatedSymField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIX43))
	builder.Header.Set(field.NewMsgType("AH"))
	builder.Body.Set(rfqreqid)
	builder.Body.Set(norelatedsym)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.BeginStringFIX43, "AH", r
}
