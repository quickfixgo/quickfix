package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//AllocationReportAck msg type = AT.
type AllocationReportAck struct {
	message.Message
}

//AllocationReportAckBuilder builds AllocationReportAck messages.
type AllocationReportAckBuilder struct {
	message.MessageBuilder
}

//CreateAllocationReportAckBuilder returns an initialized AllocationReportAckBuilder with specified required fields.
func CreateAllocationReportAckBuilder(
	allocreportid field.AllocReportID,
	allocid field.AllocID) AllocationReportAckBuilder {
	var builder AllocationReportAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(allocreportid)
	builder.Body.Set(allocid)
	return builder
}

//AllocReportID is a required field for AllocationReportAck.
func (m AllocationReportAck) AllocReportID() (field.AllocReportID, error) {
	var f field.AllocReportID
	err := m.Body.Get(&f)
	return f, err
}

//AllocID is a required field for AllocationReportAck.
func (m AllocationReportAck) AllocID() (field.AllocID, error) {
	var f field.AllocID
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for AllocationReportAck.
func (m AllocationReportAck) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryAllocID is a non-required field for AllocationReportAck.
func (m AllocationReportAck) SecondaryAllocID() (field.SecondaryAllocID, error) {
	var f field.SecondaryAllocID
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for AllocationReportAck.
func (m AllocationReportAck) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for AllocationReportAck.
func (m AllocationReportAck) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//AllocStatus is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AllocStatus() (field.AllocStatus, error) {
	var f field.AllocStatus
	err := m.Body.Get(&f)
	return f, err
}

//AllocRejCode is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AllocRejCode() (field.AllocRejCode, error) {
	var f field.AllocRejCode
	err := m.Body.Get(&f)
	return f, err
}

//AllocReportType is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AllocReportType() (field.AllocReportType, error) {
	var f field.AllocReportType
	err := m.Body.Get(&f)
	return f, err
}

//AllocIntermedReqType is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AllocIntermedReqType() (field.AllocIntermedReqType, error) {
	var f field.AllocIntermedReqType
	err := m.Body.Get(&f)
	return f, err
}

//MatchStatus is a non-required field for AllocationReportAck.
func (m AllocationReportAck) MatchStatus() (field.MatchStatus, error) {
	var f field.MatchStatus
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for AllocationReportAck.
func (m AllocationReportAck) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for AllocationReportAck.
func (m AllocationReportAck) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for AllocationReportAck.
func (m AllocationReportAck) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for AllocationReportAck.
func (m AllocationReportAck) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for AllocationReportAck.
func (m AllocationReportAck) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//NoAllocs is a non-required field for AllocationReportAck.
func (m AllocationReportAck) NoAllocs() (field.NoAllocs, error) {
	var f field.NoAllocs
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for AllocationReportAck.
func (m AllocationReportAck) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//AvgPxIndicator is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AvgPxIndicator() (field.AvgPxIndicator, error) {
	var f field.AvgPxIndicator
	err := m.Body.Get(&f)
	return f, err
}

//Quantity is a non-required field for AllocationReportAck.
func (m AllocationReportAck) Quantity() (field.Quantity, error) {
	var f field.Quantity
	err := m.Body.Get(&f)
	return f, err
}

//AllocTransType is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AllocTransType() (field.AllocTransType, error) {
	var f field.AllocTransType
	err := m.Body.Get(&f)
	return f, err
}
