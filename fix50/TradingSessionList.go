package fix50

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type TradingSessionList struct {
	quickfix.Message
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
