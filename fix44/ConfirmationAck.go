package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type ConfirmationAck struct {
	message.Message
}

func (m *ConfirmationAck) ConfirmID() (*field.ConfirmID, error) {
	f := new(field.ConfirmID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationAck) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationAck) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationAck) AffirmStatus() (*field.AffirmStatus, error) {
	f := new(field.AffirmStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationAck) MatchStatus() (*field.MatchStatus, error) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationAck) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationAck) ConfirmRejReason() (*field.ConfirmRejReason, error) {
	f := new(field.ConfirmRejReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationAck) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *ConfirmationAck) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
