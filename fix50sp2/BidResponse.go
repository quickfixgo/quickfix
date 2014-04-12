package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type BidResponse struct {
	message.Message
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
func (m *BidResponse) NoBidComponents() (*field.NoBidComponents, error) {
	f := new(field.NoBidComponents)
	err := m.Body.Get(f)
	return f, err
}
