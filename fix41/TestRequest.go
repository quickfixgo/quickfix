package fix41

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type TestRequest struct {
	quickfix.Message
}

func (m *TestRequest) TestReqID() (*field.TestReqID, error) {
	f := new(field.TestReqID)
	err := m.Body.Get(f)
	return f, err
}
