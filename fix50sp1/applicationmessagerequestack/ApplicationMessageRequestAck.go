//Package applicationmessagerequestack msg type = BX.
package applicationmessagerequestack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
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
func (m Message) ApplResponseID() (*field.ApplResponseIDField, errors.MessageRejectError) {
	f := &field.ApplResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResponseID reads a ApplResponseID from ApplicationMessageRequestAck.
func (m Message) GetApplResponseID(f *field.ApplResponseIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplReqID is a non-required field for ApplicationMessageRequestAck.
func (m Message) ApplReqID() (*field.ApplReqIDField, errors.MessageRejectError) {
	f := &field.ApplReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplReqID reads a ApplReqID from ApplicationMessageRequestAck.
func (m Message) GetApplReqID(f *field.ApplReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplReqType is a non-required field for ApplicationMessageRequestAck.
func (m Message) ApplReqType() (*field.ApplReqTypeField, errors.MessageRejectError) {
	f := &field.ApplReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplReqType reads a ApplReqType from ApplicationMessageRequestAck.
func (m Message) GetApplReqType(f *field.ApplReqTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResponseType is a non-required field for ApplicationMessageRequestAck.
func (m Message) ApplResponseType() (*field.ApplResponseTypeField, errors.MessageRejectError) {
	f := &field.ApplResponseTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResponseType reads a ApplResponseType from ApplicationMessageRequestAck.
func (m Message) GetApplResponseType(f *field.ApplResponseTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplTotalMessageCount is a non-required field for ApplicationMessageRequestAck.
func (m Message) ApplTotalMessageCount() (*field.ApplTotalMessageCountField, errors.MessageRejectError) {
	f := &field.ApplTotalMessageCountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplTotalMessageCount reads a ApplTotalMessageCount from ApplicationMessageRequestAck.
func (m Message) GetApplTotalMessageCount(f *field.ApplTotalMessageCountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoApplIDs is a non-required field for ApplicationMessageRequestAck.
func (m Message) NoApplIDs() (*field.NoApplIDsField, errors.MessageRejectError) {
	f := &field.NoApplIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoApplIDs reads a NoApplIDs from ApplicationMessageRequestAck.
func (m Message) GetNoApplIDs(f *field.NoApplIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ApplicationMessageRequestAck.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ApplicationMessageRequestAck.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ApplicationMessageRequestAck.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ApplicationMessageRequestAck.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ApplicationMessageRequestAck.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ApplicationMessageRequestAck.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds ApplicationMessageRequestAck messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for ApplicationMessageRequestAck.
func Builder(
	applresponseid *field.ApplResponseIDField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("BX"))
	builder.Body().Set(applresponseid)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "BX", r
}
