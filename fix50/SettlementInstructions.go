package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
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

//CreateSettlementInstructionsBuilder returns an initialized SettlementInstructionsBuilder with specified required fields.
func CreateSettlementInstructionsBuilder(
	settlinstmsgid field.SettlInstMsgID,
	settlinstmode field.SettlInstMode,
	transacttime field.TransactTime) SettlementInstructionsBuilder {
	var builder SettlementInstructionsBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(settlinstmsgid)
	builder.Body.Set(settlinstmode)
	builder.Body.Set(transacttime)
	return builder
}

//SettlInstMsgID is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstMsgID() (field.SettlInstMsgID, errors.MessageRejectError) {
	var f field.SettlInstMsgID
	err := m.Body.Get(&f)
	return f, err
}

//SettlInstReqID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstReqID() (field.SettlInstReqID, errors.MessageRejectError) {
	var f field.SettlInstReqID
	err := m.Body.Get(&f)
	return f, err
}

//SettlInstMode is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstMode() (field.SettlInstMode, errors.MessageRejectError) {
	var f field.SettlInstMode
	err := m.Body.Get(&f)
	return f, err
}

//SettlInstReqRejCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstReqRejCode() (field.SettlInstReqRejCode, errors.MessageRejectError) {
	var f field.SettlInstReqRejCode
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for SettlementInstructions.
func (m SettlementInstructions) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for SettlementInstructions.
func (m SettlementInstructions) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for SettlementInstructions.
func (m SettlementInstructions) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) ClOrdID() (field.ClOrdID, errors.MessageRejectError) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for SettlementInstructions.
func (m SettlementInstructions) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoSettlInst is a non-required field for SettlementInstructions.
func (m SettlementInstructions) NoSettlInst() (field.NoSettlInst, errors.MessageRejectError) {
	var f field.NoSettlInst
	err := m.Body.Get(&f)
	return f, err
}
