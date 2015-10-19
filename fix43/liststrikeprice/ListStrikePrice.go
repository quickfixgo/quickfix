//Package liststrikeprice msg type = m.
package liststrikeprice

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a ListStrikePrice wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//ListID is a required field for ListStrikePrice.
func (m Message) ListID() (*field.ListIDField, quickfix.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ListStrikePrice.
func (m Message) GetListID(f *field.ListIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoStrikes is a required field for ListStrikePrice.
func (m Message) TotNoStrikes() (*field.TotNoStrikesField, quickfix.MessageRejectError) {
	f := &field.TotNoStrikesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoStrikes reads a TotNoStrikes from ListStrikePrice.
func (m Message) GetTotNoStrikes(f *field.TotNoStrikesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoStrikes is a required field for ListStrikePrice.
func (m Message) NoStrikes() (*field.NoStrikesField, quickfix.MessageRejectError) {
	f := &field.NoStrikesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStrikes reads a NoStrikes from ListStrikePrice.
func (m Message) GetNoStrikes(f *field.NoStrikesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds ListStrikePrice messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for ListStrikePrice.
func Builder(
	listid *field.ListIDField,
	totnostrikes *field.TotNoStrikesField,
	nostrikes *field.NoStrikesField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX43))
	builder.Header.Set(field.NewMsgType("m"))
	builder.Body.Set(listid)
	builder.Body.Set(totnostrikes)
	builder.Body.Set(nostrikes)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX43, "m", r
}
