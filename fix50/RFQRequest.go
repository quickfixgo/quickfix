package fix50

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
func (m *RFQRequest) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}
func (m *RFQRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
