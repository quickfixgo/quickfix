package fix44

import (
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

//NewQuoteRequestRejectBuilder returns an initialized QuoteRequestRejectBuilder with specified required fields.
func NewQuoteRequestRejectBuilder(
	quotereqid field.QuoteReqID,
	quoterequestrejectreason field.QuoteRequestRejectReason,
	norelatedsym field.NoRelatedSym) *QuoteRequestRejectBuilder {
	builder := new(QuoteRequestRejectBuilder)
	builder.Body.Set(quotereqid)
	builder.Body.Set(quoterequestrejectreason)
	builder.Body.Set(norelatedsym)
	return builder
}

//QuoteReqID is a required field for QuoteRequestReject.
func (m *QuoteRequestReject) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//RFQReqID is a non-required field for QuoteRequestReject.
func (m *QuoteRequestReject) RFQReqID() (*field.RFQReqID, error) {
	f := new(field.RFQReqID)
	err := m.Body.Get(f)
	return f, err
}

//QuoteRequestRejectReason is a required field for QuoteRequestReject.
func (m *QuoteRequestReject) QuoteRequestRejectReason() (*field.QuoteRequestRejectReason, error) {
	f := new(field.QuoteRequestRejectReason)
	err := m.Body.Get(f)
	return f, err
}

//NoRelatedSym is a required field for QuoteRequestReject.
func (m *QuoteRequestReject) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for QuoteRequestReject.
func (m *QuoteRequestReject) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for QuoteRequestReject.
func (m *QuoteRequestReject) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for QuoteRequestReject.
func (m *QuoteRequestReject) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
