package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildMsgType("P"))
	builder.Body.Set(allocid)
	builder.Body.Set(allocstatus)
	return builder
}

//AllocID is a required field for AllocationInstructionAck.
func (m AllocationInstructionAck) AllocID() (*field.AllocID, errors.MessageRejectError) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from AllocationInstructionAck.
func (m AllocationInstructionAck) GetAllocID(f *field.AllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AllocationInstructionAck.
func (m AllocationInstructionAck) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryAllocID is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) SecondaryAllocID() (*field.SecondaryAllocID, errors.MessageRejectError) {
	f := new(field.SecondaryAllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryAllocID reads a SecondaryAllocID from AllocationInstructionAck.
func (m AllocationInstructionAck) GetSecondaryAllocID(f *field.SecondaryAllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from AllocationInstructionAck.
func (m AllocationInstructionAck) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from AllocationInstructionAck.
func (m AllocationInstructionAck) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocStatus is a required field for AllocationInstructionAck.
func (m AllocationInstructionAck) AllocStatus() (*field.AllocStatus, errors.MessageRejectError) {
	f := new(field.AllocStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocStatus reads a AllocStatus from AllocationInstructionAck.
func (m AllocationInstructionAck) GetAllocStatus(f *field.AllocStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocRejCode is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) AllocRejCode() (*field.AllocRejCode, errors.MessageRejectError) {
	f := new(field.AllocRejCode)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocRejCode reads a AllocRejCode from AllocationInstructionAck.
func (m AllocationInstructionAck) GetAllocRejCode(f *field.AllocRejCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocType is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) AllocType() (*field.AllocType, errors.MessageRejectError) {
	f := new(field.AllocType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocType reads a AllocType from AllocationInstructionAck.
func (m AllocationInstructionAck) GetAllocType(f *field.AllocType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocIntermedReqType is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) AllocIntermedReqType() (*field.AllocIntermedReqType, errors.MessageRejectError) {
	f := new(field.AllocIntermedReqType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocIntermedReqType reads a AllocIntermedReqType from AllocationInstructionAck.
func (m AllocationInstructionAck) GetAllocIntermedReqType(f *field.AllocIntermedReqType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchStatus is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) MatchStatus() (*field.MatchStatus, errors.MessageRejectError) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetMatchStatus reads a MatchStatus from AllocationInstructionAck.
func (m AllocationInstructionAck) GetMatchStatus(f *field.MatchStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from AllocationInstructionAck.
func (m AllocationInstructionAck) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from AllocationInstructionAck.
func (m AllocationInstructionAck) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from AllocationInstructionAck.
func (m AllocationInstructionAck) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from AllocationInstructionAck.
func (m AllocationInstructionAck) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from AllocationInstructionAck.
func (m AllocationInstructionAck) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for AllocationInstructionAck.
func (m AllocationInstructionAck) NoAllocs() (*field.NoAllocs, errors.MessageRejectError) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from AllocationInstructionAck.
func (m AllocationInstructionAck) GetNoAllocs(f *field.NoAllocs) errors.MessageRejectError {
	return m.Body.Get(f)
}
