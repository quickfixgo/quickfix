package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
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
	confirmid field.ConfirmID,
	tradedate field.TradeDate,
	transacttime field.TransactTime,
	affirmstatus field.AffirmStatus) ConfirmationAckBuilder {
	var builder ConfirmationAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("AU"))
	builder.Body.Set(confirmid)
	builder.Body.Set(tradedate)
	builder.Body.Set(transacttime)
	builder.Body.Set(affirmstatus)
	return builder
}

//ConfirmID is a required field for ConfirmationAck.
func (m ConfirmationAck) ConfirmID() (*field.ConfirmID, errors.MessageRejectError) {
	f := new(field.ConfirmID)
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmID reads a ConfirmID from ConfirmationAck.
func (m ConfirmationAck) GetConfirmID(f *field.ConfirmID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for ConfirmationAck.
func (m ConfirmationAck) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from ConfirmationAck.
func (m ConfirmationAck) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for ConfirmationAck.
func (m ConfirmationAck) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ConfirmationAck.
func (m ConfirmationAck) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AffirmStatus is a required field for ConfirmationAck.
func (m ConfirmationAck) AffirmStatus() (*field.AffirmStatus, errors.MessageRejectError) {
	f := new(field.AffirmStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetAffirmStatus reads a AffirmStatus from ConfirmationAck.
func (m ConfirmationAck) GetAffirmStatus(f *field.AffirmStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ConfirmRejReason is a non-required field for ConfirmationAck.
func (m ConfirmationAck) ConfirmRejReason() (*field.ConfirmRejReason, errors.MessageRejectError) {
	f := new(field.ConfirmRejReason)
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmRejReason reads a ConfirmRejReason from ConfirmationAck.
func (m ConfirmationAck) GetConfirmRejReason(f *field.ConfirmRejReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchStatus is a non-required field for ConfirmationAck.
func (m ConfirmationAck) MatchStatus() (*field.MatchStatus, errors.MessageRejectError) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetMatchStatus reads a MatchStatus from ConfirmationAck.
func (m ConfirmationAck) GetMatchStatus(f *field.MatchStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ConfirmationAck.
func (m ConfirmationAck) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ConfirmationAck.
func (m ConfirmationAck) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ConfirmationAck.
func (m ConfirmationAck) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ConfirmationAck.
func (m ConfirmationAck) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ConfirmationAck.
func (m ConfirmationAck) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ConfirmationAck.
func (m ConfirmationAck) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
