package fix50sp2

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type RFQRequest struct {
	quickfixgo.Message
}

func (m *RFQRequest) RFQReqID() (*field.RFQReqID, error) {
	f := new(field.RFQReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *RFQRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *RFQRequest) PrivateQuote() (*field.PrivateQuote, error) {
	f := new(field.PrivateQuote)
	err := m.Body.Get(f)
	return f, err
}
