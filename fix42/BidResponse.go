package fix42

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type BidResponse struct {
	quickfix.Message
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
