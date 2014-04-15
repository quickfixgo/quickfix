package fix44

import (
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
	transacttime field.TransactTime,
	allocstatus field.AllocStatus) AllocationInstructionAckBuilder {
	var builder AllocationInstructionAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(allocid)
	builder.Body.Set(transacttime)
	builder.Body.Set(allocstatus)
	return builder
}

//AllocID is a required field for AllocationInstructionAck.
func (m AllocationInstructionAck) AllocID() (field.AllocID, error) {
	var f field.AllocID
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryAllocID is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) SecondaryAllocID() (field.SecondaryAllocID, error) {
	var f field.SecondaryAllocID
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for AllocationInstructionAck.
func (m AllocationInstructionAck) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//AllocStatus is a required field for AllocationInstructionAck.
func (m AllocationInstructionAck) AllocStatus() (field.AllocStatus, error) {
	var f field.AllocStatus
	err := m.Body.Get(&f)
	return f, err
}

//AllocRejCode is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) AllocRejCode() (field.AllocRejCode, error) {
	var f field.AllocRejCode
	err := m.Body.Get(&f)
	return f, err
}

//AllocType is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) AllocType() (field.AllocType, error) {
	var f field.AllocType
	err := m.Body.Get(&f)
	return f, err
}

//AllocIntermedReqType is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) AllocIntermedReqType() (field.AllocIntermedReqType, error) {
	var f field.AllocIntermedReqType
	err := m.Body.Get(&f)
	return f, err
}

//MatchStatus is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) MatchStatus() (field.MatchStatus, error) {
	var f field.MatchStatus
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//NoAllocs is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) NoAllocs() (field.NoAllocs, error) {
	var f field.NoAllocs
	err := m.Body.Get(&f)
	return f, err
}
