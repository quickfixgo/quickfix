package fix44

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type SettlementInstructions struct {
	quickfix.Message
}

func (m *SettlementInstructions) SettlInstReqID() (*field.SettlInstReqID, error) {
	f := new(field.SettlInstReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SettlInstMode() (*field.SettlInstMode, error) {
	f := new(field.SettlInstMode)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) NoSettlInst() (*field.NoSettlInst, error) {
	f := new(field.NoSettlInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SettlInstMsgID() (*field.SettlInstMsgID, error) {
	f := new(field.SettlInstMsgID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SettlInstReqRejCode() (*field.SettlInstReqRejCode, error) {
	f := new(field.SettlInstReqRejCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
