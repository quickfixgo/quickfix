package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type TestRequest struct {
	message.Message
}

func (m *TestRequest) TestReqID() (*field.TestReqID, error) {
	f := new(field.TestReqID)
	err := m.Body.Get(f)
	return f, err
}
