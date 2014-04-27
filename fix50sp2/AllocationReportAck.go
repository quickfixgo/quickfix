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
	allocreportid field.AllocReportID) AllocationReportAckBuilder {
	var builder AllocationReportAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.BuildMsgType("AT"))
	builder.Body.Set(allocreportid)
	return builder
}

//AllocReportID is a required field for AllocationReportAck.
func (m AllocationReportAck) AllocReportID() (*field.AllocReportID, errors.MessageRejectError) {
	f := new(field.AllocReportID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocReportID reads a AllocReportID from AllocationReportAck.
func (m AllocationReportAck) GetAllocReportID(f *field.AllocReportID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AllocID() (*field.AllocID, errors.MessageRejectError) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from AllocationReportAck.
func (m AllocationReportAck) GetAllocID(f *field.AllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for AllocationReportAck.
func (m AllocationReportAck) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AllocationReportAck.
func (m AllocationReportAck) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryAllocID is a non-required field for AllocationReportAck.
func (m AllocationReportAck) SecondaryAllocID() (*field.SecondaryAllocID, errors.MessageRejectError) {
	f := new(field.SecondaryAllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryAllocID reads a SecondaryAllocID from AllocationReportAck.
func (m AllocationReportAck) GetSecondaryAllocID(f *field.SecondaryAllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for AllocationReportAck.
func (m AllocationReportAck) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from AllocationReportAck.
func (m AllocationReportAck) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for AllocationReportAck.
func (m AllocationReportAck) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from AllocationReportAck.
func (m AllocationReportAck) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocStatus is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AllocStatus() (*field.AllocStatus, errors.MessageRejectError) {
	f := new(field.AllocStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocStatus reads a AllocStatus from AllocationReportAck.
func (m AllocationReportAck) GetAllocStatus(f *field.AllocStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocRejCode is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AllocRejCode() (*field.AllocRejCode, errors.MessageRejectError) {
	f := new(field.AllocRejCode)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocRejCode reads a AllocRejCode from AllocationReportAck.
func (m AllocationReportAck) GetAllocRejCode(f *field.AllocRejCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocReportType is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AllocReportType() (*field.AllocReportType, errors.MessageRejectError) {
	f := new(field.AllocReportType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocReportType reads a AllocReportType from AllocationReportAck.
func (m AllocationReportAck) GetAllocReportType(f *field.AllocReportType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocIntermedReqType is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AllocIntermedReqType() (*field.AllocIntermedReqType, errors.MessageRejectError) {
	f := new(field.AllocIntermedReqType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocIntermedReqType reads a AllocIntermedReqType from AllocationReportAck.
func (m AllocationReportAck) GetAllocIntermedReqType(f *field.AllocIntermedReqType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchStatus is a non-required field for AllocationReportAck.
func (m AllocationReportAck) MatchStatus() (*field.MatchStatus, errors.MessageRejectError) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetMatchStatus reads a MatchStatus from AllocationReportAck.
func (m AllocationReportAck) GetMatchStatus(f *field.MatchStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for AllocationReportAck.
func (m AllocationReportAck) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from AllocationReportAck.
func (m AllocationReportAck) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for AllocationReportAck.
func (m AllocationReportAck) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from AllocationReportAck.
func (m AllocationReportAck) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for AllocationReportAck.
func (m AllocationReportAck) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from AllocationReportAck.
func (m AllocationReportAck) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for AllocationReportAck.
func (m AllocationReportAck) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from AllocationReportAck.
func (m AllocationReportAck) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for AllocationReportAck.
func (m AllocationReportAck) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from AllocationReportAck.
func (m AllocationReportAck) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for AllocationReportAck.
func (m AllocationReportAck) NoAllocs() (*field.NoAllocs, errors.MessageRejectError) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from AllocationReportAck.
func (m AllocationReportAck) GetNoAllocs(f *field.NoAllocs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for AllocationReportAck.
func (m AllocationReportAck) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from AllocationReportAck.
func (m AllocationReportAck) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPxIndicator is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AvgPxIndicator() (*field.AvgPxIndicator, errors.MessageRejectError) {
	f := new(field.AvgPxIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPxIndicator reads a AvgPxIndicator from AllocationReportAck.
func (m AllocationReportAck) GetAvgPxIndicator(f *field.AvgPxIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Quantity is a non-required field for AllocationReportAck.
func (m AllocationReportAck) Quantity() (*field.Quantity, errors.MessageRejectError) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}

//GetQuantity reads a Quantity from AllocationReportAck.
func (m AllocationReportAck) GetQuantity(f *field.Quantity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocTransType is a non-required field for AllocationReportAck.
func (m AllocationReportAck) AllocTransType() (*field.AllocTransType, errors.MessageRejectError) {
	f := new(field.AllocTransType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocTransType reads a AllocTransType from AllocationReportAck.
func (m AllocationReportAck) GetAllocTransType(f *field.AllocTransType) errors.MessageRejectError {
	return m.Body.Get(f)
}
