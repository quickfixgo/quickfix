//Package confirmationack msg type = AU.
package confirmationack

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

//Message is a ConfirmationAck wrapper for the generic Message type
type Message struct {
	message.Message
}

//ConfirmID is a required field for ConfirmationAck.
func (m Message) ConfirmID() (*field.ConfirmIDField, errors.MessageRejectError) {
	f := &field.ConfirmIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmID reads a ConfirmID from ConfirmationAck.
func (m Message) GetConfirmID(f *field.ConfirmIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for ConfirmationAck.
func (m Message) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from ConfirmationAck.
func (m Message) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for ConfirmationAck.
func (m Message) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ConfirmationAck.
func (m Message) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AffirmStatus is a required field for ConfirmationAck.
func (m Message) AffirmStatus() (*field.AffirmStatusField, errors.MessageRejectError) {
	f := &field.AffirmStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAffirmStatus reads a AffirmStatus from ConfirmationAck.
func (m Message) GetAffirmStatus(f *field.AffirmStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ConfirmRejReason is a non-required field for ConfirmationAck.
func (m Message) ConfirmRejReason() (*field.ConfirmRejReasonField, errors.MessageRejectError) {
	f := &field.ConfirmRejReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmRejReason reads a ConfirmRejReason from ConfirmationAck.
func (m Message) GetConfirmRejReason(f *field.ConfirmRejReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchStatus is a non-required field for ConfirmationAck.
func (m Message) MatchStatus() (*field.MatchStatusField, errors.MessageRejectError) {
	f := &field.MatchStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchStatus reads a MatchStatus from ConfirmationAck.
func (m Message) GetMatchStatus(f *field.MatchStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ConfirmationAck.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ConfirmationAck.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ConfirmationAck.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ConfirmationAck.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ConfirmationAck.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ConfirmationAck.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds ConfirmationAck messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for ConfirmationAck.
func Builder(
	confirmid *field.ConfirmIDField,
	tradedate *field.TradeDateField,
	transacttime *field.TransactTimeField,
	affirmstatus *field.AffirmStatusField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("AU"))
	builder.Body().Set(confirmid)
	builder.Body().Set(tradedate)
	builder.Body().Set(transacttime)
	builder.Body().Set(affirmstatus)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "AU", r
}
