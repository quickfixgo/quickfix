//Package quoteacknowledgement msg type = b.
package quoteacknowledgement

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a QuoteAcknowledgement wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//QuoteReqID is a non-required field for QuoteAcknowledgement.
func (m Message) QuoteReqID() (*field.QuoteReqIDField, quickfix.MessageRejectError) {
	f := &field.QuoteReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from QuoteAcknowledgement.
func (m Message) GetQuoteReqID(f *field.QuoteReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for QuoteAcknowledgement.
func (m Message) QuoteID() (*field.QuoteIDField, quickfix.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from QuoteAcknowledgement.
func (m Message) GetQuoteID(f *field.QuoteIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteAckStatus is a required field for QuoteAcknowledgement.
func (m Message) QuoteAckStatus() (*field.QuoteAckStatusField, quickfix.MessageRejectError) {
	f := &field.QuoteAckStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteAckStatus reads a QuoteAckStatus from QuoteAcknowledgement.
func (m Message) GetQuoteAckStatus(f *field.QuoteAckStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteRejectReason is a non-required field for QuoteAcknowledgement.
func (m Message) QuoteRejectReason() (*field.QuoteRejectReasonField, quickfix.MessageRejectError) {
	f := &field.QuoteRejectReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteRejectReason reads a QuoteRejectReason from QuoteAcknowledgement.
func (m Message) GetQuoteRejectReason(f *field.QuoteRejectReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteResponseLevel is a non-required field for QuoteAcknowledgement.
func (m Message) QuoteResponseLevel() (*field.QuoteResponseLevelField, quickfix.MessageRejectError) {
	f := &field.QuoteResponseLevelField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteResponseLevel reads a QuoteResponseLevel from QuoteAcknowledgement.
func (m Message) GetQuoteResponseLevel(f *field.QuoteResponseLevelField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for QuoteAcknowledgement.
func (m Message) TradingSessionID() (*field.TradingSessionIDField, quickfix.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from QuoteAcknowledgement.
func (m Message) GetTradingSessionID(f *field.TradingSessionIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for QuoteAcknowledgement.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from QuoteAcknowledgement.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteSets is a non-required field for QuoteAcknowledgement.
func (m Message) NoQuoteSets() (*field.NoQuoteSetsField, quickfix.MessageRejectError) {
	f := &field.NoQuoteSetsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteSets reads a NoQuoteSets from QuoteAcknowledgement.
func (m Message) GetNoQuoteSets(f *field.NoQuoteSetsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds QuoteAcknowledgement messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for QuoteAcknowledgement.
func Builder(
	quoteackstatus *field.QuoteAckStatusField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header.Set(field.NewMsgType("b"))
	builder.Body.Set(quoteackstatus)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX42, "b", r
}
