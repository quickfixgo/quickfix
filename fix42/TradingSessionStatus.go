package fix42

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type TradingSessionStatus struct {
	message.Message
}

func (m *TradingSessionStatus) TradSesReqID() (*field.TradSesReqID, error) {
	f := new(field.TradSesReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionStatus) TradSesMode() (*field.TradSesMode, error) {
	f := new(field.TradSesMode)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionStatus) TradSesStatus() (*field.TradSesStatus, error) {
	f := new(field.TradSesStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionStatus) TradSesPreCloseTime() (*field.TradSesPreCloseTime, error) {
	f := new(field.TradSesPreCloseTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionStatus) TotalVolumeTraded() (*field.TotalVolumeTraded, error) {
	f := new(field.TotalVolumeTraded)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionStatus) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionStatus) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionStatus) TradSesCloseTime() (*field.TradSesCloseTime, error) {
	f := new(field.TradSesCloseTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionStatus) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionStatus) TradSesMethod() (*field.TradSesMethod, error) {
	f := new(field.TradSesMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionStatus) UnsolicitedIndicator() (*field.UnsolicitedIndicator, error) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionStatus) TradSesStartTime() (*field.TradSesStartTime, error) {
	f := new(field.TradSesStartTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionStatus) TradSesEndTime() (*field.TradSesEndTime, error) {
	f := new(field.TradSesEndTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionStatus) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionStatus) TradSesOpenTime() (*field.TradSesOpenTime, error) {
	f := new(field.TradSesOpenTime)
	err := m.Body.Get(f)
	return f, err
}
