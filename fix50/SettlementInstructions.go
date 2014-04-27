package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header.Set(field.BuildMsgType("T"))
	builder.Body.Set(settlinstmsgid)
	builder.Body.Set(settlinstmode)
	builder.Body.Set(transacttime)
	return builder
}

//SettlInstMsgID is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstMsgID() (*field.SettlInstMsgID, errors.MessageRejectError) {
	f := new(field.SettlInstMsgID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstMsgID reads a SettlInstMsgID from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstMsgID(f *field.SettlInstMsgID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstReqID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstReqID() (*field.SettlInstReqID, errors.MessageRejectError) {
	f := new(field.SettlInstReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstReqID reads a SettlInstReqID from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstReqID(f *field.SettlInstReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstMode is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstMode() (*field.SettlInstMode, errors.MessageRejectError) {
	f := new(field.SettlInstMode)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstMode reads a SettlInstMode from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstMode(f *field.SettlInstMode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstReqRejCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstReqRejCode() (*field.SettlInstReqRejCode, errors.MessageRejectError) {
	f := new(field.SettlInstReqRejCode)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstReqRejCode reads a SettlInstReqRejCode from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstReqRejCode(f *field.SettlInstReqRejCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for SettlementInstructions.
func (m SettlementInstructions) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from SettlementInstructions.
func (m SettlementInstructions) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for SettlementInstructions.
func (m SettlementInstructions) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from SettlementInstructions.
func (m SettlementInstructions) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for SettlementInstructions.
func (m SettlementInstructions) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from SettlementInstructions.
func (m SettlementInstructions) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from SettlementInstructions.
func (m SettlementInstructions) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for SettlementInstructions.
func (m SettlementInstructions) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SettlementInstructions.
func (m SettlementInstructions) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSettlInst is a non-required field for SettlementInstructions.
func (m SettlementInstructions) NoSettlInst() (*field.NoSettlInst, errors.MessageRejectError) {
	f := new(field.NoSettlInst)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSettlInst reads a NoSettlInst from SettlementInstructions.
func (m SettlementInstructions) GetNoSettlInst(f *field.NoSettlInst) errors.MessageRejectError {
	return m.Body.Get(f)
}
