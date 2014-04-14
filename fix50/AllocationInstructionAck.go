package fix50

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

//NewAllocationInstructionAckBuilder returns an initialized AllocationInstructionAckBuilder with specified required fields.
func NewAllocationInstructionAckBuilder(
	allocid field.AllocID,
	allocstatus field.AllocStatus) *AllocationInstructionAckBuilder {
	builder := new(AllocationInstructionAckBuilder)
	builder.Body.Set(allocid)
	builder.Body.Set(allocstatus)
	return builder
}

//AllocID is a required field for AllocationInstructionAck.
func (m *AllocationInstructionAck) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for AllocationInstructionAck.
func (m *AllocationInstructionAck) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryAllocID is a non-required field for AllocationInstructionAck.
func (m *AllocationInstructionAck) SecondaryAllocID() (*field.SecondaryAllocID, error) {
	f := new(field.SecondaryAllocID)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a non-required field for AllocationInstructionAck.
func (m *AllocationInstructionAck) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for AllocationInstructionAck.
func (m *AllocationInstructionAck) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//AllocStatus is a required field for AllocationInstructionAck.
func (m *AllocationInstructionAck) AllocStatus() (*field.AllocStatus, error) {
	f := new(field.AllocStatus)
	err := m.Body.Get(f)
	return f, err
}

//AllocRejCode is a non-required field for AllocationInstructionAck.
func (m *AllocationInstructionAck) AllocRejCode() (*field.AllocRejCode, error) {
	f := new(field.AllocRejCode)
	err := m.Body.Get(f)
	return f, err
}

//AllocType is a non-required field for AllocationInstructionAck.
func (m *AllocationInstructionAck) AllocType() (*field.AllocType, error) {
	f := new(field.AllocType)
	err := m.Body.Get(f)
	return f, err
}

//AllocIntermedReqType is a non-required field for AllocationInstructionAck.
func (m *AllocationInstructionAck) AllocIntermedReqType() (*field.AllocIntermedReqType, error) {
	f := new(field.AllocIntermedReqType)
	err := m.Body.Get(f)
	return f, err
}

//MatchStatus is a non-required field for AllocationInstructionAck.
func (m *AllocationInstructionAck) MatchStatus() (*field.MatchStatus, error) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for AllocationInstructionAck.
func (m *AllocationInstructionAck) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for AllocationInstructionAck.
func (m *AllocationInstructionAck) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for AllocationInstructionAck.
func (m *AllocationInstructionAck) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for AllocationInstructionAck.
func (m *AllocationInstructionAck) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for AllocationInstructionAck.
func (m *AllocationInstructionAck) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//NoAllocs is a non-required field for AllocationInstructionAck.
func (m *AllocationInstructionAck) NoAllocs() (*field.NoAllocs, error) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}
