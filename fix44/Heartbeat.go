package fix44

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type Heartbeat struct {
	quickfixgo.Message
}

func (m *Heartbeat) TestReqID() (*field.TestReqID, error) {
	f := new(field.TestReqID)
	err := m.Body.Get(f)
	return f, err
}
