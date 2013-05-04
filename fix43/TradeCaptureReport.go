package fix43

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type TradeCaptureReport struct {
	quickfixgo.Message
}

func (m *TradeCaptureReport) TradeReportID() (*field.TradeReportID, error) {
	f := new(field.TradeReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TradeReportTransType() (*field.TradeReportTransType, error) {
	f := new(field.TradeReportTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TradeRequestID() (*field.TradeRequestID, error) {
	f := new(field.TradeRequestID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) ExecType() (*field.ExecType, error) {
	f := new(field.ExecType)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TradeReportRefID() (*field.TradeReportRefID, error) {
	f := new(field.TradeReportRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) ExecID() (*field.ExecID, error) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) SecondaryExecID() (*field.SecondaryExecID, error) {
	f := new(field.SecondaryExecID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) ExecRestatementReason() (*field.ExecRestatementReason, error) {
	f := new(field.ExecRestatementReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) PreviouslyReported() (*field.PreviouslyReported, error) {
	f := new(field.PreviouslyReported)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) LastQty() (*field.LastQty, error) {
	f := new(field.LastQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) LastSpotRate() (*field.LastSpotRate, error) {
	f := new(field.LastSpotRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) LastForwardPoints() (*field.LastForwardPoints, error) {
	f := new(field.LastForwardPoints)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) SettlmntTyp() (*field.SettlmntTyp, error) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) FutSettDate() (*field.FutSettDate, error) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) MatchStatus() (*field.MatchStatus, error) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradeCaptureReport) MatchType() (*field.MatchType, error) {
	f := new(field.MatchType)
	err := m.Body.Get(f)
	return f, err
}
