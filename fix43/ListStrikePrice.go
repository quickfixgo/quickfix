package fix43

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ListStrikePrice msg type = m.
type ListStrikePrice struct {
	message.Message
}

//ListStrikePriceBuilder builds ListStrikePrice messages.
type ListStrikePriceBuilder struct {
	message.MessageBuilder
}

//NewListStrikePriceBuilder returns an initialized ListStrikePriceBuilder with specified required fields.
func NewListStrikePriceBuilder(
	listid field.ListID,
	totnostrikes field.TotNoStrikes,
	nostrikes field.NoStrikes) *ListStrikePriceBuilder {
	builder := new(ListStrikePriceBuilder)
	builder.Body.Set(listid)
	builder.Body.Set(totnostrikes)
	builder.Body.Set(nostrikes)
	return builder
}

//ListID is a required field for ListStrikePrice.
func (m *ListStrikePrice) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//TotNoStrikes is a required field for ListStrikePrice.
func (m *ListStrikePrice) TotNoStrikes() (*field.TotNoStrikes, error) {
	f := new(field.TotNoStrikes)
	err := m.Body.Get(f)
	return f, err
}

//NoStrikes is a required field for ListStrikePrice.
func (m *ListStrikePrice) NoStrikes() (*field.NoStrikes, error) {
	f := new(field.NoStrikes)
	err := m.Body.Get(f)
	return f, err
}
