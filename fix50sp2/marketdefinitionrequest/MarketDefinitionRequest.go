//Package marketdefinitionrequest msg type = BT.
package marketdefinitionrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a MarketDefinitionRequest wrapper for the generic Message type
type Message struct {
	message.Message
}

//MarketReqID is a required field for MarketDefinitionRequest.
func (m Message) MarketReqID() (*field.MarketReqIDField, errors.MessageRejectError) {
	f := &field.MarketReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketReqID reads a MarketReqID from MarketDefinitionRequest.
func (m Message) GetMarketReqID(f *field.MarketReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a required field for MarketDefinitionRequest.
func (m Message) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, errors.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from MarketDefinitionRequest.
func (m Message) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for MarketDefinitionRequest.
func (m Message) MarketID() (*field.MarketIDField, errors.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from MarketDefinitionRequest.
func (m Message) GetMarketID(f *field.MarketIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for MarketDefinitionRequest.
func (m Message) MarketSegmentID() (*field.MarketSegmentIDField, errors.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from MarketDefinitionRequest.
func (m Message) GetMarketSegmentID(f *field.MarketSegmentIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParentMktSegmID is a non-required field for MarketDefinitionRequest.
func (m Message) ParentMktSegmID() (*field.ParentMktSegmIDField, errors.MessageRejectError) {
	f := &field.ParentMktSegmIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetParentMktSegmID reads a ParentMktSegmID from MarketDefinitionRequest.
func (m Message) GetParentMktSegmID(f *field.ParentMktSegmIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds MarketDefinitionRequest messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for MarketDefinitionRequest.
func Builder(
	marketreqid *field.MarketReqIDField,
	subscriptionrequesttype *field.SubscriptionRequestTypeField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("BT"))
	builder.Body().Set(marketreqid)
	builder.Body().Set(subscriptionrequesttype)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "BT", r
}
