package fix50sp1

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type AllocationReportAck struct {
	quickfix.Message
}

func (m *AllocationReportAck) AllocReportID() (*field.AllocReportID, error) {
	f := new(field.AllocReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) MatchStatus() (*field.MatchStatus, error) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) AllocTransType() (*field.AllocTransType, error) {
	f := new(field.AllocTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) SecondaryAllocID() (*field.SecondaryAllocID, error) {
	f := new(field.SecondaryAllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) AllocIntermedReqType() (*field.AllocIntermedReqType, error) {
	f := new(field.AllocIntermedReqType)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) AvgPxIndicator() (*field.AvgPxIndicator, error) {
	f := new(field.AvgPxIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) Quantity() (*field.Quantity, error) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) AllocReportType() (*field.AllocReportType, error) {
	f := new(field.AllocReportType)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) NoAllocs() (*field.NoAllocs, error) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) AllocStatus() (*field.AllocStatus, error) {
	f := new(field.AllocStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) AllocRejCode() (*field.AllocRejCode, error) {
	f := new(field.AllocRejCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}
func (m *AllocationReportAck) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
