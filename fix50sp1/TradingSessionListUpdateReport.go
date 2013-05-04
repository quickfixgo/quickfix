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
