package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SettlementInstructions msg type = T.
type SettlementInstructions struct {
	message.Message
}

//SettlementInstructionsBuilder builds SettlementInstructions messages.
type SettlementInstructionsBuilder struct {
	message.MessageBuilder
}

//NewSettlementInstructionsBuilder returns an initialized SettlementInstructionsBuilder with specified required fields.
func NewSettlementInstructionsBuilder(
	settlinstmsgid field.SettlInstMsgID,
	settlinstmode field.SettlInstMode,
	transacttime field.TransactTime) *SettlementInstructionsBuilder {
	builder := new(SettlementInstructionsBuilder)
	builder.Body.Set(settlinstmsgid)
	builder.Body.Set(settlinstmode)
	builder.Body.Set(transacttime)
	return builder
}

//SettlInstMsgID is a required field for SettlementInstructions.
func (m *SettlementInstructions) SettlInstMsgID() (*field.SettlInstMsgID, error) {
	f := new(field.SettlInstMsgID)
	err := m.Body.Get(f)
	return f, err
}

//SettlInstReqID is a non-required field for SettlementInstructions.
func (m *SettlementInstructions) SettlInstReqID() (*field.SettlInstReqID, error) {
	f := new(field.SettlInstReqID)
	err := m.Body.Get(f)
	return f, err
}

//SettlInstMode is a required field for SettlementInstructions.
func (m *SettlementInstructions) SettlInstMode() (*field.SettlInstMode, error) {
	f := new(field.SettlInstMode)
	err := m.Body.Get(f)
	return f, err
}

//SettlInstReqRejCode is a non-required field for SettlementInstructions.
func (m *SettlementInstructions) SettlInstReqRejCode() (*field.SettlInstReqRejCode, error) {
	f := new(field.SettlInstReqRejCode)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for SettlementInstructions.
func (m *SettlementInstructions) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for SettlementInstructions.
func (m *SettlementInstructions) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for SettlementInstructions.
func (m *SettlementInstructions) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a non-required field for SettlementInstructions.
func (m *SettlementInstructions) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a required field for SettlementInstructions.
func (m *SettlementInstructions) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//NoSettlInst is a non-required field for SettlementInstructions.
func (m *SettlementInstructions) NoSettlInst() (*field.NoSettlInst, error) {
	f := new(field.NoSettlInst)
	err := m.Body.Get(f)
	return f, err
}
