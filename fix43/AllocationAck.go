package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//AllocationAck msg type = P.
type AllocationAck struct {
	message.Message
}

//AllocationAckBuilder builds AllocationAck messages.
type AllocationAckBuilder struct {
	message.MessageBuilder
}

//CreateAllocationAckBuilder returns an initialized AllocationAckBuilder with specified required fields.
func CreateAllocationAckBuilder(
	allocid *field.AllocIDField,
	tradedate *field.TradeDateField,
	allocstatus *field.AllocStatusField) AllocationAckBuilder {
	var builder AllocationAckBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX43))
	builder.Header().Set(field.NewMsgType("P"))
	builder.Body().Set(allocid)
	builder.Body().Set(tradedate)
	builder.Body().Set(allocstatus)
	return builder
}

//NoPartyIDs is a non-required field for AllocationAck.
func (m AllocationAck) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AllocationAck.
func (m AllocationAck) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a required field for AllocationAck.
func (m AllocationAck) AllocID() (*field.AllocIDField, errors.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from AllocationAck.
func (m AllocationAck) GetAllocID(f *field.AllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for AllocationAck.
func (m AllocationAck) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from AllocationAck.
func (m AllocationAck) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for AllocationAck.
func (m AllocationAck) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from AllocationAck.
func (m AllocationAck) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocStatus is a required field for AllocationAck.
func (m AllocationAck) AllocStatus() (*field.AllocStatusField, errors.MessageRejectError) {
	f := &field.AllocStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocStatus reads a AllocStatus from AllocationAck.
func (m AllocationAck) GetAllocStatus(f *field.AllocStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocRejCode is a non-required field for AllocationAck.
func (m AllocationAck) AllocRejCode() (*field.AllocRejCodeField, errors.MessageRejectError) {
	f := &field.AllocRejCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocRejCode reads a AllocRejCode from AllocationAck.
func (m AllocationAck) GetAllocRejCode(f *field.AllocRejCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for AllocationAck.
func (m AllocationAck) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from AllocationAck.
func (m AllocationAck) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for AllocationAck.
func (m AllocationAck) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from AllocationAck.
func (m AllocationAck) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for AllocationAck.
func (m AllocationAck) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from AllocationAck.
func (m AllocationAck) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LegalConfirm is a non-required field for AllocationAck.
func (m AllocationAck) LegalConfirm() (*field.LegalConfirmField, errors.MessageRejectError) {
	f := &field.LegalConfirmField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLegalConfirm reads a LegalConfirm from AllocationAck.
func (m AllocationAck) GetLegalConfirm(f *field.LegalConfirmField) errors.MessageRejectError {
	return m.Body.Get(f)
}
