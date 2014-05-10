package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//QuoteRequestReject msg type = AG.
type QuoteRequestReject struct {
	message.Message
}

//QuoteRequestRejectBuilder builds QuoteRequestReject messages.
type QuoteRequestRejectBuilder struct {
	message.MessageBuilder
}

//CreateQuoteRequestRejectBuilder returns an initialized QuoteRequestRejectBuilder with specified required fields.
func CreateQuoteRequestRejectBuilder(
	quotereqid *field.QuoteReqIDField,
	quoterequestrejectreason *field.QuoteRequestRejectReasonField,
	norelatedsym *field.NoRelatedSymField) QuoteRequestRejectBuilder {
	var builder QuoteRequestRejectBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header.Set(field.NewMsgType("AG"))
	builder.Body.Set(quotereqid)
	builder.Body.Set(quoterequestrejectreason)
	builder.Body.Set(norelatedsym)
	return builder
}

//QuoteReqID is a required field for QuoteRequestReject.
func (m QuoteRequestReject) QuoteReqID() (*field.QuoteReqIDField, errors.MessageRejectError) {
	f := &field.QuoteReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from QuoteRequestReject.
func (m QuoteRequestReject) GetQuoteReqID(f *field.QuoteReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RFQReqID is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) RFQReqID() (*field.RFQReqIDField, errors.MessageRejectError) {
	f := &field.RFQReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRFQReqID reads a RFQReqID from QuoteRequestReject.
func (m QuoteRequestReject) GetRFQReqID(f *field.RFQReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteRequestRejectReason is a required field for QuoteRequestReject.
func (m QuoteRequestReject) QuoteRequestRejectReason() (*field.QuoteRequestRejectReasonField, errors.MessageRejectError) {
	f := &field.QuoteRequestRejectReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteRequestRejectReason reads a QuoteRequestRejectReason from QuoteRequestReject.
func (m QuoteRequestReject) GetQuoteRequestRejectReason(f *field.QuoteRequestRejectReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a required field for QuoteRequestReject.
func (m QuoteRequestReject) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from QuoteRequestReject.
func (m QuoteRequestReject) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from QuoteRequestReject.
func (m QuoteRequestReject) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from QuoteRequestReject.
func (m QuoteRequestReject) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from QuoteRequestReject.
func (m QuoteRequestReject) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
