//Package reject msg type = 3.
package reject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a Reject wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//RefSeqNum is a required field for Reject.
func (m Message) RefSeqNum() (*field.RefSeqNumField, quickfix.MessageRejectError) {
	f := &field.RefSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefSeqNum reads a RefSeqNum from Reject.
func (m Message) GetRefSeqNum(f *field.RefSeqNumField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RefTagID is a non-required field for Reject.
func (m Message) RefTagID() (*field.RefTagIDField, quickfix.MessageRejectError) {
	f := &field.RefTagIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefTagID reads a RefTagID from Reject.
func (m Message) GetRefTagID(f *field.RefTagIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RefMsgType is a non-required field for Reject.
func (m Message) RefMsgType() (*field.RefMsgTypeField, quickfix.MessageRejectError) {
	f := &field.RefMsgTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefMsgType reads a RefMsgType from Reject.
func (m Message) GetRefMsgType(f *field.RefMsgTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SessionRejectReason is a non-required field for Reject.
func (m Message) SessionRejectReason() (*field.SessionRejectReasonField, quickfix.MessageRejectError) {
	f := &field.SessionRejectReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSessionRejectReason reads a SessionRejectReason from Reject.
func (m Message) GetSessionRejectReason(f *field.SessionRejectReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for Reject.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from Reject.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for Reject.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from Reject.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for Reject.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from Reject.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for Reject.
func New(
	refseqnum *field.RefSeqNumField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIXT11))
	builder.Header.Set(field.NewMsgType("3"))
	builder.Body.Set(refseqnum)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.BeginStringFIXT11, "3", r
}
