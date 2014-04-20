package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
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
	allocid field.AllocID,
	tradedate field.TradeDate,
	allocstatus field.AllocStatus) AllocationAckBuilder {
	var builder AllocationAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(allocid)
	builder.Body.Set(tradedate)
	builder.Body.Set(allocstatus)
	return builder
}

//NoPartyIDs is a non-required field for AllocationAck.
func (m AllocationAck) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AllocationAck.
func (m AllocationAck) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a required field for AllocationAck.
func (m AllocationAck) AllocID() (*field.AllocID, errors.MessageRejectError) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from AllocationAck.
func (m AllocationAck) GetAllocID(f *field.AllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for AllocationAck.
func (m AllocationAck) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from AllocationAck.
func (m AllocationAck) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for AllocationAck.
func (m AllocationAck) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from AllocationAck.
func (m AllocationAck) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocStatus is a required field for AllocationAck.
func (m AllocationAck) AllocStatus() (*field.AllocStatus, errors.MessageRejectError) {
	f := new(field.AllocStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocStatus reads a AllocStatus from AllocationAck.
func (m AllocationAck) GetAllocStatus(f *field.AllocStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocRejCode is a non-required field for AllocationAck.
func (m AllocationAck) AllocRejCode() (*field.AllocRejCode, errors.MessageRejectError) {
	f := new(field.AllocRejCode)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocRejCode reads a AllocRejCode from AllocationAck.
func (m AllocationAck) GetAllocRejCode(f *field.AllocRejCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for AllocationAck.
func (m AllocationAck) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from AllocationAck.
func (m AllocationAck) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for AllocationAck.
func (m AllocationAck) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from AllocationAck.
func (m AllocationAck) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for AllocationAck.
func (m AllocationAck) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from AllocationAck.
func (m AllocationAck) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LegalConfirm is a non-required field for AllocationAck.
func (m AllocationAck) LegalConfirm() (*field.LegalConfirm, errors.MessageRejectError) {
	f := new(field.LegalConfirm)
	err := m.Body.Get(f)
	return f, err
}

//GetLegalConfirm reads a LegalConfirm from AllocationAck.
func (m AllocationAck) GetLegalConfirm(f *field.LegalConfirm) errors.MessageRejectError {
	return m.Body.Get(f)
}
