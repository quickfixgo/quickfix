//Package quoterequestreject msg type = AG.
package quoterequestreject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a QuoteRequestReject wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//QuoteReqID is a required field for QuoteRequestReject.
func (m Message) QuoteReqID() (*field.QuoteReqIDField, quickfix.MessageRejectError) {
	f := &field.QuoteReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from QuoteRequestReject.
func (m Message) GetQuoteReqID(f *field.QuoteReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RFQReqID is a non-required field for QuoteRequestReject.
func (m Message) RFQReqID() (*field.RFQReqIDField, quickfix.MessageRejectError) {
	f := &field.RFQReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRFQReqID reads a RFQReqID from QuoteRequestReject.
func (m Message) GetRFQReqID(f *field.RFQReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteRequestRejectReason is a required field for QuoteRequestReject.
func (m Message) QuoteRequestRejectReason() (*field.QuoteRequestRejectReasonField, quickfix.MessageRejectError) {
	f := &field.QuoteRequestRejectReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteRequestRejectReason reads a QuoteRequestRejectReason from QuoteRequestReject.
func (m Message) GetQuoteRequestRejectReason(f *field.QuoteRequestRejectReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a required field for QuoteRequestReject.
func (m Message) NoRelatedSym() (*field.NoRelatedSymField, quickfix.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from QuoteRequestReject.
func (m Message) GetNoRelatedSym(f *field.NoRelatedSymField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for QuoteRequestReject.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from QuoteRequestReject.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for QuoteRequestReject.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from QuoteRequestReject.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for QuoteRequestReject.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from QuoteRequestReject.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds QuoteRequestReject messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for QuoteRequestReject.
func Builder(
	quotereqid *field.QuoteReqIDField,
	quoterequestrejectreason *field.QuoteRequestRejectReasonField,
	norelatedsym *field.NoRelatedSymField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header.Set(field.NewMsgType("AG"))
	builder.Body.Set(quotereqid)
	builder.Body.Set(quoterequestrejectreason)
	builder.Body.Set(norelatedsym)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX44, "AG", r
}
