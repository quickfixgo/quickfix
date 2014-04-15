package fix50sp2

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

//CreateListStrikePriceBuilder returns an initialized ListStrikePriceBuilder with specified required fields.
func CreateListStrikePriceBuilder(
	listid field.ListID,
	totnostrikes field.TotNoStrikes,
	nostrikes field.NoStrikes) ListStrikePriceBuilder {
	var builder ListStrikePriceBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(listid)
	builder.Body.Set(totnostrikes)
	builder.Body.Set(nostrikes)
	return builder
}

//ListID is a required field for ListStrikePrice.
func (m ListStrikePrice) ListID() (field.ListID, error) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//TotNoStrikes is a required field for ListStrikePrice.
func (m ListStrikePrice) TotNoStrikes() (field.TotNoStrikes, error) {
	var f field.TotNoStrikes
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for ListStrikePrice.
func (m ListStrikePrice) LastFragment() (field.LastFragment, error) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoStrikes is a required field for ListStrikePrice.
func (m ListStrikePrice) NoStrikes() (field.NoStrikes, error) {
	var f field.NoStrikes
	err := m.Body.Get(&f)
	return f, err
}
