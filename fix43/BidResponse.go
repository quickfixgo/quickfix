package fix43

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type BidResponse struct {
	quickfixgo.Message
}

func (m *BidResponse) BidID() (*field.BidID, error) {
	f := new(field.BidID)
	err := m.Body.Get(f)
	return f, err
}
func (m *BidResponse) ClientBidID() (*field.ClientBidID, error) {
	f := new(field.ClientBidID)
	err := m.Body.Get(f)
	return f, err
}
