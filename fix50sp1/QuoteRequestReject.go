package fix50sp1

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

//CreateQuoteRequestRejectBuilder returns an initialized QuoteRequestRejectBuilder with specified required fields.
func CreateQuoteRequestRejectBuilder(
	quotereqid field.QuoteReqID,
	quoterequestrejectreason field.QuoteRequestRejectReason,
	norelatedsym field.NoRelatedSym) QuoteRequestRejectBuilder {
	var builder QuoteRequestRejectBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(quotereqid)
	builder.Body.Set(quoterequestrejectreason)
	builder.Body.Set(norelatedsym)
	return builder
}

//QuoteReqID is a required field for QuoteRequestReject.
func (m QuoteRequestReject) QuoteReqID() (field.QuoteReqID, error) {
	var f field.QuoteReqID
	err := m.Body.Get(&f)
	return f, err
}

//RFQReqID is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) RFQReqID() (field.RFQReqID, error) {
	var f field.RFQReqID
	err := m.Body.Get(&f)
	return f, err
}

//QuoteRequestRejectReason is a required field for QuoteRequestReject.
func (m QuoteRequestReject) QuoteRequestRejectReason() (field.QuoteRequestRejectReason, error) {
	var f field.QuoteRequestRejectReason
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a required field for QuoteRequestReject.
func (m QuoteRequestReject) NoRelatedSym() (field.NoRelatedSym, error) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//NoRootPartyIDs is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) NoRootPartyIDs() (field.NoRootPartyIDs, error) {
	var f field.NoRootPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//PrivateQuote is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) PrivateQuote() (field.PrivateQuote, error) {
	var f field.PrivateQuote
	err := m.Body.Get(&f)
	return f, err
}

//RespondentType is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) RespondentType() (field.RespondentType, error) {
	var f field.RespondentType
	err := m.Body.Get(&f)
	return f, err
}

//PreTradeAnonymity is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) PreTradeAnonymity() (field.PreTradeAnonymity, error) {
	var f field.PreTradeAnonymity
	err := m.Body.Get(&f)
	return f, err
}
