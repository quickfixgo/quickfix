//Package applicationmessagerequestack msg type = BX.
package applicationmessagerequestack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a ApplicationMessageRequestAck wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//ApplResponseID is a required field for ApplicationMessageRequestAck.
func (m Message) ApplResponseID() (*field.ApplResponseIDField, quickfix.MessageRejectError) {
	f := &field.ApplResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResponseID reads a ApplResponseID from ApplicationMessageRequestAck.
func (m Message) GetApplResponseID(f *field.ApplResponseIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplReqID is a non-required field for ApplicationMessageRequestAck.
func (m Message) ApplReqID() (*field.ApplReqIDField, quickfix.MessageRejectError) {
	f := &field.ApplReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplReqID reads a ApplReqID from ApplicationMessageRequestAck.
func (m Message) GetApplReqID(f *field.ApplReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplReqType is a non-required field for ApplicationMessageRequestAck.
func (m Message) ApplReqType() (*field.ApplReqTypeField, quickfix.MessageRejectError) {
	f := &field.ApplReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplReqType reads a ApplReqType from ApplicationMessageRequestAck.
func (m Message) GetApplReqType(f *field.ApplReqTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResponseType is a non-required field for ApplicationMessageRequestAck.
func (m Message) ApplResponseType() (*field.ApplResponseTypeField, quickfix.MessageRejectError) {
	f := &field.ApplResponseTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResponseType reads a ApplResponseType from ApplicationMessageRequestAck.
func (m Message) GetApplResponseType(f *field.ApplResponseTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplTotalMessageCount is a non-required field for ApplicationMessageRequestAck.
func (m Message) ApplTotalMessageCount() (*field.ApplTotalMessageCountField, quickfix.MessageRejectError) {
	f := &field.ApplTotalMessageCountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplTotalMessageCount reads a ApplTotalMessageCount from ApplicationMessageRequestAck.
func (m Message) GetApplTotalMessageCount(f *field.ApplTotalMessageCountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoApplIDs is a non-required field for ApplicationMessageRequestAck.
func (m Message) NoApplIDs() (*field.NoApplIDsField, quickfix.MessageRejectError) {
	f := &field.NoApplIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoApplIDs reads a NoApplIDs from ApplicationMessageRequestAck.
func (m Message) GetNoApplIDs(f *field.NoApplIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ApplicationMessageRequestAck.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ApplicationMessageRequestAck.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ApplicationMessageRequestAck.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ApplicationMessageRequestAck.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ApplicationMessageRequestAck.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ApplicationMessageRequestAck.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized MessageBuilder with specified required fields for ApplicationMessageRequestAck.
func New(
	applresponseid *field.ApplResponseIDField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("BX"))
	builder.Body.Set(applresponseid)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "BX", r
}
