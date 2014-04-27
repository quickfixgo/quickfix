package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header.Set(field.BuildMsgType("m"))
	builder.Body.Set(listid)
	builder.Body.Set(totnostrikes)
	builder.Body.Set(nostrikes)
	return builder
}

//ListID is a required field for ListStrikePrice.
func (m ListStrikePrice) ListID() (*field.ListID, errors.MessageRejectError) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ListStrikePrice.
func (m ListStrikePrice) GetListID(f *field.ListID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoStrikes is a required field for ListStrikePrice.
func (m ListStrikePrice) TotNoStrikes() (*field.TotNoStrikes, errors.MessageRejectError) {
	f := new(field.TotNoStrikes)
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoStrikes reads a TotNoStrikes from ListStrikePrice.
func (m ListStrikePrice) GetTotNoStrikes(f *field.TotNoStrikes) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for ListStrikePrice.
func (m ListStrikePrice) LastFragment() (*field.LastFragment, errors.MessageRejectError) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from ListStrikePrice.
func (m ListStrikePrice) GetLastFragment(f *field.LastFragment) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStrikes is a required field for ListStrikePrice.
func (m ListStrikePrice) NoStrikes() (*field.NoStrikes, errors.MessageRejectError) {
	f := new(field.NoStrikes)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStrikes reads a NoStrikes from ListStrikePrice.
func (m ListStrikePrice) GetNoStrikes(f *field.NoStrikes) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for ListStrikePrice.
func (m ListStrikePrice) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from ListStrikePrice.
func (m ListStrikePrice) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}
