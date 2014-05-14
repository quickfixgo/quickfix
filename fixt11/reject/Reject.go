//Package reject msg type = 3.
package reject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a Reject wrapper for the generic Message type
type Message struct {
	message.Message
}

//RefSeqNum is a required field for Reject.
func (m Message) RefSeqNum() (*field.RefSeqNumField, errors.MessageRejectError) {
	f := &field.RefSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefSeqNum reads a RefSeqNum from Reject.
func (m Message) GetRefSeqNum(f *field.RefSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefTagID is a non-required field for Reject.
func (m Message) RefTagID() (*field.RefTagIDField, errors.MessageRejectError) {
	f := &field.RefTagIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefTagID reads a RefTagID from Reject.
func (m Message) GetRefTagID(f *field.RefTagIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefMsgType is a non-required field for Reject.
func (m Message) RefMsgType() (*field.RefMsgTypeField, errors.MessageRejectError) {
	f := &field.RefMsgTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefMsgType reads a RefMsgType from Reject.
func (m Message) GetRefMsgType(f *field.RefMsgTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SessionRejectReason is a non-required field for Reject.
func (m Message) SessionRejectReason() (*field.SessionRejectReasonField, errors.MessageRejectError) {
	f := &field.SessionRejectReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSessionRejectReason reads a SessionRejectReason from Reject.
func (m Message) GetSessionRejectReason(f *field.SessionRejectReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for Reject.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from Reject.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for Reject.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from Reject.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for Reject.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from Reject.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds Reject messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for Reject.
func Builder(
	refseqnum *field.RefSeqNumField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewMsgType("3"))
	builder.Body().Set(refseqnum)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIXT11, "3", r
}
