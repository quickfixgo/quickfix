package fix50sp2

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
