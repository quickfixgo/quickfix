package fix50sp2

import (
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

//NewConfirmationAckBuilder returns an initialized ConfirmationAckBuilder with specified required fields.
func NewConfirmationAckBuilder(
	confirmid field.ConfirmID,
	tradedate field.TradeDate,
	transacttime field.TransactTime,
	affirmstatus field.AffirmStatus) *ConfirmationAckBuilder {
	builder := new(ConfirmationAckBuilder)
	builder.Body.Set(confirmid)
	builder.Body.Set(tradedate)
	builder.Body.Set(transacttime)
	builder.Body.Set(affirmstatus)
	return builder
}

//ConfirmID is a required field for ConfirmationAck.
func (m *ConfirmationAck) ConfirmID() (*field.ConfirmID, error) {
	f := new(field.ConfirmID)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a required field for ConfirmationAck.
func (m *ConfirmationAck) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a required field for ConfirmationAck.
func (m *ConfirmationAck) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//AffirmStatus is a required field for ConfirmationAck.
func (m *ConfirmationAck) AffirmStatus() (*field.AffirmStatus, error) {
	f := new(field.AffirmStatus)
	err := m.Body.Get(f)
	return f, err
}

//ConfirmRejReason is a non-required field for ConfirmationAck.
func (m *ConfirmationAck) ConfirmRejReason() (*field.ConfirmRejReason, error) {
	f := new(field.ConfirmRejReason)
	err := m.Body.Get(f)
	return f, err
}

//MatchStatus is a non-required field for ConfirmationAck.
func (m *ConfirmationAck) MatchStatus() (*field.MatchStatus, error) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for ConfirmationAck.
func (m *ConfirmationAck) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for ConfirmationAck.
func (m *ConfirmationAck) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for ConfirmationAck.
func (m *ConfirmationAck) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
