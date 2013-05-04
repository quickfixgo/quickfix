package fix50sp2

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type TradeCaptureReportRequestAck struct {
	quickfixgo.Message
}

func (m *TradeCaptureReportRequestAck) TradeRequestID() (*field.TradeRequestID, error) {
	f := new(field.TradeRequestID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequestAck) TradeRequestType() (*field.TradeRequestType, error) {
	f := new(field.TradeRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequestAck) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequestAck) TotNumTradeReports() (*field.TotNumTradeReports, error) {
	f := new(field.TotNumTradeReports)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequestAck) TradeRequestResult() (*field.TradeRequestResult, error) {
	f := new(field.TradeRequestResult)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequestAck) TradeRequestStatus() (*field.TradeRequestStatus, error) {
	f := new(field.TradeRequestStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequestAck) MultiLegReportingType() (*field.MultiLegReportingType, error) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequestAck) ResponseTransportType() (*field.ResponseTransportType, error) {
	f := new(field.ResponseTransportType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequestAck) ResponseDestination() (*field.ResponseDestination, error) {
	f := new(field.ResponseDestination)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequestAck) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequestAck) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequestAck) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequestAck) MessageEventSource() (*field.MessageEventSource, error) {
	f := new(field.MessageEventSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequestAck) TradeID() (*field.TradeID, error) {
	f := new(field.TradeID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequestAck) SecondaryTradeID() (*field.SecondaryTradeID, error) {
	f := new(field.SecondaryTradeID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequestAck) FirmTradeID() (*field.FirmTradeID, error) {
	f := new(field.FirmTradeID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReportRequestAck) SecondaryFirmTradeID() (*field.SecondaryFirmTradeID, error) {
	f := new(field.SecondaryFirmTradeID)
	err := m.Body.Get(f)
	return f, err
}
