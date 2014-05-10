package fix50sp2

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
	settlinstmsgid *field.SettlInstMsgIDField,
	settlinstmode *field.SettlInstModeField,
	transacttime *field.TransactTimeField) SettlementInstructionsBuilder {
	var builder SettlementInstructionsBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("T"))
	builder.Body.Set(settlinstmsgid)
	builder.Body.Set(settlinstmode)
	builder.Body.Set(transacttime)
	return builder
}

//SettlInstMsgID is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstMsgID() (*field.SettlInstMsgIDField, errors.MessageRejectError) {
	f := &field.SettlInstMsgIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstMsgID reads a SettlInstMsgID from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstMsgID(f *field.SettlInstMsgIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstReqID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstReqID() (*field.SettlInstReqIDField, errors.MessageRejectError) {
	f := &field.SettlInstReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstReqID reads a SettlInstReqID from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstReqID(f *field.SettlInstReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstMode is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstMode() (*field.SettlInstModeField, errors.MessageRejectError) {
	f := &field.SettlInstModeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstMode reads a SettlInstMode from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstMode(f *field.SettlInstModeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstReqRejCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstReqRejCode() (*field.SettlInstReqRejCodeField, errors.MessageRejectError) {
	f := &field.SettlInstReqRejCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstReqRejCode reads a SettlInstReqRejCode from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstReqRejCode(f *field.SettlInstReqRejCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for SettlementInstructions.
func (m SettlementInstructions) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from SettlementInstructions.
func (m SettlementInstructions) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for SettlementInstructions.
func (m SettlementInstructions) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from SettlementInstructions.
func (m SettlementInstructions) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for SettlementInstructions.
func (m SettlementInstructions) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from SettlementInstructions.
func (m SettlementInstructions) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from SettlementInstructions.
func (m SettlementInstructions) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for SettlementInstructions.
func (m SettlementInstructions) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SettlementInstructions.
func (m SettlementInstructions) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSettlInst is a non-required field for SettlementInstructions.
func (m SettlementInstructions) NoSettlInst() (*field.NoSettlInstField, errors.MessageRejectError) {
	f := &field.NoSettlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSettlInst reads a NoSettlInst from SettlementInstructions.
func (m SettlementInstructions) GetNoSettlInst(f *field.NoSettlInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}
