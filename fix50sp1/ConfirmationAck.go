package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//ConfirmationAck msg type = AU.
type ConfirmationAck struct {
	message.Message
}

//ConfirmationAckBuilder builds ConfirmationAck messages.
type ConfirmationAckBuilder struct {
	message.MessageBuilder
}

//CreateConfirmationAckBuilder returns an initialized ConfirmationAckBuilder with specified required fields.
func CreateConfirmationAckBuilder(
	confirmid *field.ConfirmIDField,
	tradedate *field.TradeDateField,
	transacttime *field.TransactTimeField,
	affirmstatus *field.AffirmStatusField) ConfirmationAckBuilder {
	var builder ConfirmationAckBuilder
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

//ConfirmID is a required field for ConfirmationAck.
func (m ConfirmationAck) ConfirmID() (*field.ConfirmIDField, errors.MessageRejectError) {
	f := &field.ConfirmIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmID reads a ConfirmID from ConfirmationAck.
func (m ConfirmationAck) GetConfirmID(f *field.ConfirmIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for ConfirmationAck.
func (m ConfirmationAck) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from ConfirmationAck.
func (m ConfirmationAck) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for ConfirmationAck.
func (m ConfirmationAck) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ConfirmationAck.
func (m ConfirmationAck) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AffirmStatus is a required field for ConfirmationAck.
func (m ConfirmationAck) AffirmStatus() (*field.AffirmStatusField, errors.MessageRejectError) {
	f := &field.AffirmStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAffirmStatus reads a AffirmStatus from ConfirmationAck.
func (m ConfirmationAck) GetAffirmStatus(f *field.AffirmStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ConfirmRejReason is a non-required field for ConfirmationAck.
func (m ConfirmationAck) ConfirmRejReason() (*field.ConfirmRejReasonField, errors.MessageRejectError) {
	f := &field.ConfirmRejReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmRejReason reads a ConfirmRejReason from ConfirmationAck.
func (m ConfirmationAck) GetConfirmRejReason(f *field.ConfirmRejReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchStatus is a non-required field for ConfirmationAck.
func (m ConfirmationAck) MatchStatus() (*field.MatchStatusField, errors.MessageRejectError) {
	f := &field.MatchStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchStatus reads a MatchStatus from ConfirmationAck.
func (m ConfirmationAck) GetMatchStatus(f *field.MatchStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ConfirmationAck.
func (m ConfirmationAck) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ConfirmationAck.
func (m ConfirmationAck) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ConfirmationAck.
func (m ConfirmationAck) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ConfirmationAck.
func (m ConfirmationAck) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ConfirmationAck.
func (m ConfirmationAck) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ConfirmationAck.
func (m ConfirmationAck) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
