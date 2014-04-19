package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//AllocationInstructionAck msg type = P.
type AllocationInstructionAck struct {
	message.Message
}

//AllocationInstructionAckBuilder builds AllocationInstructionAck messages.
type AllocationInstructionAckBuilder struct {
	message.MessageBuilder
}

//CreateAllocationInstructionAckBuilder returns an initialized AllocationInstructionAckBuilder with specified required fields.
func CreateAllocationInstructionAckBuilder(
	allocid field.AllocID,
	allocstatus field.AllocStatus) AllocationInstructionAckBuilder {
	var builder AllocationInstructionAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(allocid)
	builder.Body.Set(allocstatus)
	return builder
}

//AllocID is a required field for AllocationInstructionAck.
func (m AllocationInstructionAck) AllocID() (field.AllocID, errors.MessageRejectError) {
	var f field.AllocID
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) NoPartyIDs() (field.NoPartyIDs, errors.MessageRejectError) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryAllocID is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) SecondaryAllocID() (field.SecondaryAllocID, errors.MessageRejectError) {
	var f field.SecondaryAllocID
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) TradeDate() (field.TradeDate, errors.MessageRejectError) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//AllocStatus is a required field for AllocationInstructionAck.
func (m AllocationInstructionAck) AllocStatus() (field.AllocStatus, errors.MessageRejectError) {
	var f field.AllocStatus
	err := m.Body.Get(&f)
	return f, err
}

//AllocRejCode is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) AllocRejCode() (field.AllocRejCode, errors.MessageRejectError) {
	var f field.AllocRejCode
	err := m.Body.Get(&f)
	return f, err
}

//AllocType is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) AllocType() (field.AllocType, errors.MessageRejectError) {
	var f field.AllocType
	err := m.Body.Get(&f)
	return f, err
}

//AllocIntermedReqType is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) AllocIntermedReqType() (field.AllocIntermedReqType, errors.MessageRejectError) {
	var f field.AllocIntermedReqType
	err := m.Body.Get(&f)
	return f, err
}

//MatchStatus is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) MatchStatus() (field.MatchStatus, errors.MessageRejectError) {
	var f field.MatchStatus
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) Product() (field.Product, errors.MessageRejectError) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//NoAllocs is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) NoAllocs() (field.NoAllocs, errors.MessageRejectError) {
	var f field.NoAllocs
	err := m.Body.Get(&f)
	return f, err
}
