//Package businessmessagereject msg type = j.
package businessmessagereject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix/enum"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a BusinessMessageReject wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//RefSeqNum is a non-required field for BusinessMessageReject.
func (m Message) RefSeqNum() (*field.RefSeqNumField, quickfix.MessageRejectError) {
	f := &field.RefSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefSeqNum reads a RefSeqNum from BusinessMessageReject.
func (m Message) GetRefSeqNum(f *field.RefSeqNumField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RefMsgType is a required field for BusinessMessageReject.
func (m Message) RefMsgType() (*field.RefMsgTypeField, quickfix.MessageRejectError) {
	f := &field.RefMsgTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefMsgType reads a RefMsgType from BusinessMessageReject.
func (m Message) GetRefMsgType(f *field.RefMsgTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BusinessRejectRefID is a non-required field for BusinessMessageReject.
func (m Message) BusinessRejectRefID() (*field.BusinessRejectRefIDField, quickfix.MessageRejectError) {
	f := &field.BusinessRejectRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBusinessRejectRefID reads a BusinessRejectRefID from BusinessMessageReject.
func (m Message) GetBusinessRejectRefID(f *field.BusinessRejectRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BusinessRejectReason is a required field for BusinessMessageReject.
func (m Message) BusinessRejectReason() (*field.BusinessRejectReasonField, quickfix.MessageRejectError) {
	f := &field.BusinessRejectReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBusinessRejectReason reads a BusinessRejectReason from BusinessMessageReject.
func (m Message) GetBusinessRejectReason(f *field.BusinessRejectReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for BusinessMessageReject.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from BusinessMessageReject.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for BusinessMessageReject.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from BusinessMessageReject.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for BusinessMessageReject.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from BusinessMessageReject.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for BusinessMessageReject.
func New(
	refmsgtype *field.RefMsgTypeField,
	businessrejectreason *field.BusinessRejectReasonField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIX44))
	builder.Header.Set(field.NewMsgType("j"))
	builder.Body.Set(refmsgtype)
	builder.Body.Set(businessrejectreason)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.BeginStringFIX44, "j", r
}
