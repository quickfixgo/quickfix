package fix44

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type RequestForPositionsAck struct {
	quickfixgo.Message
}

func (m *RequestForPositionsAck) PosMaintRptID() (*field.PosMaintRptID, error) {
	f := new(field.PosMaintRptID)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) PosReqID() (*field.PosReqID, error) {
	f := new(field.PosReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) TotalNumPosReports() (*field.TotalNumPosReports, error) {
	f := new(field.TotalNumPosReports)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) UnsolicitedIndicator() (*field.UnsolicitedIndicator, error) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) PosReqResult() (*field.PosReqResult, error) {
	f := new(field.PosReqResult)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) PosReqStatus() (*field.PosReqStatus, error) {
	f := new(field.PosReqStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) ResponseTransportType() (*field.ResponseTransportType, error) {
	f := new(field.ResponseTransportType)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) ResponseDestination() (*field.ResponseDestination, error) {
	f := new(field.ResponseDestination)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
