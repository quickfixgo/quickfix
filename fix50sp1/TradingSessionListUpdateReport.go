package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type TradingSessionListUpdateReport struct {
	quickfixgo.Message
}

func (m *TradingSessionListUpdateReport) TradSesReqID() (*field.TradSesReqID, error) {
	f := new(field.TradSesReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionListUpdateReport) TradSesUpdateAction() (*field.TradSesUpdateAction, error) {
	f := new(field.TradSesUpdateAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionListUpdateReport) NoTradingSessions() (*field.NoTradingSessions, error) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionListUpdateReport) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionListUpdateReport) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionListUpdateReport) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionListUpdateReport) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}
