//Package listexecute msg type = L.
package listexecute

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a ListExecute wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//ListID is a required field for ListExecute.
func (m Message) ListID() (*field.ListIDField, quickfix.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ListExecute.
func (m Message) GetListID(f *field.ListIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//WaveNo is a non-required field for ListExecute.
func (m Message) WaveNo() (*field.WaveNoField, quickfix.MessageRejectError) {
	f := &field.WaveNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetWaveNo reads a WaveNo from ListExecute.
func (m Message) GetWaveNo(f *field.WaveNoField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ListExecute.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ListExecute.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for ListExecute.
func New(
	listid *field.ListIDField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIX40))
	builder.Header.Set(field.NewMsgType("L"))
	builder.Body.Set(listid)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.BeginStringFIX40, "L", r
}
