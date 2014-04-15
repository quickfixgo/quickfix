package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//QuoteRequest msg type = R.
type QuoteRequest struct {
	message.Message
}

//QuoteRequestBuilder builds QuoteRequest messages.
type QuoteRequestBuilder struct {
	message.MessageBuilder
}

//CreateQuoteRequestBuilder returns an initialized QuoteRequestBuilder with specified required fields.
func CreateQuoteRequestBuilder(
	quotereqid field.QuoteReqID,
	norelatedsym field.NoRelatedSym) QuoteRequestBuilder {
	var builder QuoteRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(quotereqid)
	builder.Body.Set(norelatedsym)
	return builder
}

//QuoteReqID is a required field for QuoteRequest.
func (m QuoteRequest) QuoteReqID() (field.QuoteReqID, error) {
	var f field.QuoteReqID
	err := m.Body.Get(&f)
	return f, err
}

//RFQReqID is a non-required field for QuoteRequest.
func (m QuoteRequest) RFQReqID() (field.RFQReqID, error) {
	var f field.RFQReqID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for QuoteRequest.
func (m QuoteRequest) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrderCapacity is a non-required field for QuoteRequest.
func (m QuoteRequest) OrderCapacity() (field.OrderCapacity, error) {
	var f field.OrderCapacity
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a required field for QuoteRequest.
func (m QuoteRequest) NoRelatedSym() (field.NoRelatedSym, error) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for QuoteRequest.
func (m QuoteRequest) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for QuoteRequest.
func (m QuoteRequest) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for QuoteRequest.
func (m QuoteRequest) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//NoRootPartyIDs is a non-required field for QuoteRequest.
func (m QuoteRequest) NoRootPartyIDs() (field.NoRootPartyIDs, error) {
	var f field.NoRootPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//PrivateQuote is a non-required field for QuoteRequest.
func (m QuoteRequest) PrivateQuote() (field.PrivateQuote, error) {
	var f field.PrivateQuote
	err := m.Body.Get(&f)
	return f, err
}

//RespondentType is a non-required field for QuoteRequest.
func (m QuoteRequest) RespondentType() (field.RespondentType, error) {
	var f field.RespondentType
	err := m.Body.Get(&f)
	return f, err
}

//PreTradeAnonymity is a non-required field for QuoteRequest.
func (m QuoteRequest) PreTradeAnonymity() (field.PreTradeAnonymity, error) {
	var f field.PreTradeAnonymity
	err := m.Body.Get(&f)
	return f, err
}

//BookingType is a non-required field for QuoteRequest.
func (m QuoteRequest) BookingType() (field.BookingType, error) {
	var f field.BookingType
	err := m.Body.Get(&f)
	return f, err
}

//OrderRestrictions is a non-required field for QuoteRequest.
func (m QuoteRequest) OrderRestrictions() (field.OrderRestrictions, error) {
	var f field.OrderRestrictions
	err := m.Body.Get(&f)
	return f, err
}
