//Package confirmationack msg type = AU.
package confirmationack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a ConfirmationAck wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//ConfirmID is a required field for ConfirmationAck.
func (m Message) ConfirmID() (*field.ConfirmIDField, quickfix.MessageRejectError) {
	f := &field.ConfirmIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmID reads a ConfirmID from ConfirmationAck.
func (m Message) GetConfirmID(f *field.ConfirmIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for ConfirmationAck.
func (m Message) TradeDate() (*field.TradeDateField, quickfix.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from ConfirmationAck.
func (m Message) GetTradeDate(f *field.TradeDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for ConfirmationAck.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ConfirmationAck.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AffirmStatus is a required field for ConfirmationAck.
func (m Message) AffirmStatus() (*field.AffirmStatusField, quickfix.MessageRejectError) {
	f := &field.AffirmStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAffirmStatus reads a AffirmStatus from ConfirmationAck.
func (m Message) GetAffirmStatus(f *field.AffirmStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ConfirmRejReason is a non-required field for ConfirmationAck.
func (m Message) ConfirmRejReason() (*field.ConfirmRejReasonField, quickfix.MessageRejectError) {
	f := &field.ConfirmRejReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmRejReason reads a ConfirmRejReason from ConfirmationAck.
func (m Message) GetConfirmRejReason(f *field.ConfirmRejReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MatchStatus is a non-required field for ConfirmationAck.
func (m Message) MatchStatus() (*field.MatchStatusField, quickfix.MessageRejectError) {
	f := &field.MatchStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchStatus reads a MatchStatus from ConfirmationAck.
func (m Message) GetMatchStatus(f *field.MatchStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ConfirmationAck.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ConfirmationAck.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ConfirmationAck.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ConfirmationAck.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ConfirmationAck.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ConfirmationAck.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds ConfirmationAck messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for ConfirmationAck.
func Builder(
	confirmid *field.ConfirmIDField,
	tradedate *field.TradeDateField,
	transacttime *field.TransactTimeField,
	affirmstatus *field.AffirmStatusField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header.Set(field.NewMsgType("AU"))
	builder.Body.Set(confirmid)
	builder.Body.Set(tradedate)
	builder.Body.Set(transacttime)
	builder.Body.Set(affirmstatus)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX44, "AU", r
}
