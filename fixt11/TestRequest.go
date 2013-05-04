package fixt11

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type TestRequest struct {
	quickfixgo.Message
}

func (m *TestRequest) TestReqID() (*field.TestReqID, error) {
	f := new(field.TestReqID)
	err := m.Body.Get(f)
	return f, err
}
