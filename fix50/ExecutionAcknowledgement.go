package fix50

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type ExecutionAcknowledgement struct {
	quickfixgo.Message
}

func (m *ExecutionAcknowledgement) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionAcknowledgement) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionAcknowledgement) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionAcknowledgement) ExecAckStatus() (*field.ExecAckStatus, error) {
	f := new(field.ExecAckStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionAcknowledgement) ExecID() (*field.ExecID, error) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionAcknowledgement) DKReason() (*field.DKReason, error) {
	f := new(field.DKReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionAcknowledgement) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionAcknowledgement) LastQty() (*field.LastQty, error) {
	f := new(field.LastQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionAcknowledgement) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionAcknowledgement) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionAcknowledgement) LastParPx() (*field.LastParPx, error) {
	f := new(field.LastParPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionAcknowledgement) CumQty() (*field.CumQty, error) {
	f := new(field.CumQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionAcknowledgement) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionAcknowledgement) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionAcknowledgement) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *ExecutionAcknowledgement) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
