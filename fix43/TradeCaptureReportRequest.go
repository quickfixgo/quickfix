package fix43

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type TradeCaptureReportRequest struct {
	quickfixgo.Message
}

func (m *TradeCaptureReportRequest) TradeRequestID() (*field.TradeRequestID, error) {
	f := new(field.TradeRequestID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequest) TradeRequestType() (*field.TradeRequestType, error) {
	f := new(field.TradeRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequest) ExecID() (*field.ExecID, error) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequest) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequest) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequest) MatchStatus() (*field.MatchStatus, error) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequest) TradeInputSource() (*field.TradeInputSource, error) {
	f := new(field.TradeInputSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequest) TradeInputDevice() (*field.TradeInputDevice, error) {
	f := new(field.TradeInputDevice)
	err := m.Body.Get(f)
	return f, err
}
