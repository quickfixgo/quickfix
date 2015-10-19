//Package listcancelrequest msg type = K.
package listcancelrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a ListCancelRequest wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//ListID is a required field for ListCancelRequest.
func (m Message) ListID() (*field.ListIDField, quickfix.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ListCancelRequest.
func (m Message) GetListID(f *field.ListIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for ListCancelRequest.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ListCancelRequest.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for ListCancelRequest.
func (m Message) TradeOriginationDate() (*field.TradeOriginationDateField, quickfix.MessageRejectError) {
	f := &field.TradeOriginationDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from ListCancelRequest.
func (m Message) GetTradeOriginationDate(f *field.TradeOriginationDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for ListCancelRequest.
func (m Message) TradeDate() (*field.TradeDateField, quickfix.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from ListCancelRequest.
func (m Message) GetTradeDate(f *field.TradeDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ListCancelRequest.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ListCancelRequest.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ListCancelRequest.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ListCancelRequest.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ListCancelRequest.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ListCancelRequest.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for ListCancelRequest.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from ListCancelRequest.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds ListCancelRequest messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for ListCancelRequest.
func Builder(
	listid *field.ListIDField,
	transacttime *field.TransactTimeField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("K"))
	builder.Body.Set(listid)
	builder.Body.Set(transacttime)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "K", r
}
