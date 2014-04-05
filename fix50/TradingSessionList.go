package fix50

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type TradingSessionList struct {
	quickfixgo.Message
}

func (m *TradingSessionList) TradSesReqID() (*field.TradSesReqID, error) {
	f := new(field.TradSesReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *TradingSessionList) NoTradingSessions() (*field.NoTradingSessions, error) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}
