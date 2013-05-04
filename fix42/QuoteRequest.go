package fix42

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type QuoteRequest struct {
	quickfixgo.Message
}

func (m *QuoteRequest) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}
