package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	listid *field.ListIDField,
	totnostrikes *field.TotNoStrikesField,
	nostrikes *field.NoStrikesField) ListStrikePriceBuilder {
	var builder ListStrikePriceBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header().Set(field.NewMsgType("m"))
	builder.Body().Set(listid)
	builder.Body().Set(totnostrikes)
	builder.Body().Set(nostrikes)
	return builder
}

//ListID is a required field for ListStrikePrice.
func (m ListStrikePrice) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ListStrikePrice.
func (m ListStrikePrice) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoStrikes is a required field for ListStrikePrice.
func (m ListStrikePrice) TotNoStrikes() (*field.TotNoStrikesField, errors.MessageRejectError) {
	f := &field.TotNoStrikesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoStrikes reads a TotNoStrikes from ListStrikePrice.
func (m ListStrikePrice) GetTotNoStrikes(f *field.TotNoStrikesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for ListStrikePrice.
func (m ListStrikePrice) LastFragment() (*field.LastFragmentField, errors.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from ListStrikePrice.
func (m ListStrikePrice) GetLastFragment(f *field.LastFragmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStrikes is a required field for ListStrikePrice.
func (m ListStrikePrice) NoStrikes() (*field.NoStrikesField, errors.MessageRejectError) {
	f := &field.NoStrikesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStrikes reads a NoStrikes from ListStrikePrice.
func (m ListStrikePrice) GetNoStrikes(f *field.NoStrikesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for ListStrikePrice.
func (m ListStrikePrice) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from ListStrikePrice.
func (m ListStrikePrice) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
