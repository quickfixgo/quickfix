//Package liststrikeprice msg type = m.
package liststrikeprice

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a ListStrikePrice wrapper for the generic Message type
type Message struct {
	message.Message
}

//ListID is a required field for ListStrikePrice.
func (m Message) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ListStrikePrice.
func (m Message) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoStrikes is a required field for ListStrikePrice.
func (m Message) TotNoStrikes() (*field.TotNoStrikesField, errors.MessageRejectError) {
	f := &field.TotNoStrikesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoStrikes reads a TotNoStrikes from ListStrikePrice.
func (m Message) GetTotNoStrikes(f *field.TotNoStrikesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for ListStrikePrice.
func (m Message) LastFragment() (*field.LastFragmentField, errors.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from ListStrikePrice.
func (m Message) GetLastFragment(f *field.LastFragmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStrikes is a required field for ListStrikePrice.
func (m Message) NoStrikes() (*field.NoStrikesField, errors.MessageRejectError) {
	f := &field.NoStrikesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStrikes reads a NoStrikes from ListStrikePrice.
func (m Message) GetNoStrikes(f *field.NoStrikesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds ListStrikePrice messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for ListStrikePrice.
func Builder(
	listid *field.ListIDField,
	totnostrikes *field.TotNoStrikesField,
	nostrikes *field.NoStrikesField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("m"))
	builder.Body().Set(listid)
	builder.Body().Set(totnostrikes)
	builder.Body().Set(nostrikes)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "m", r
}
