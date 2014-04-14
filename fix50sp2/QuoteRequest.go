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

//NewQuoteRequestBuilder returns an initialized QuoteRequestBuilder with specified required fields.
func NewQuoteRequestBuilder(
	quotereqid field.QuoteReqID,
	norelatedsym field.NoRelatedSym) *QuoteRequestBuilder {
	builder := new(QuoteRequestBuilder)
	builder.Body.Set(quotereqid)
	builder.Body.Set(norelatedsym)
	return builder
}

//QuoteReqID is a required field for QuoteRequest.
func (m *QuoteRequest) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//RFQReqID is a non-required field for QuoteRequest.
func (m *QuoteRequest) RFQReqID() (*field.RFQReqID, error) {
	f := new(field.RFQReqID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a non-required field for QuoteRequest.
func (m *QuoteRequest) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//OrderCapacity is a non-required field for QuoteRequest.
func (m *QuoteRequest) OrderCapacity() (*field.OrderCapacity, error) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//NoRelatedSym is a required field for QuoteRequest.
func (m *QuoteRequest) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for QuoteRequest.
func (m *QuoteRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for QuoteRequest.
func (m *QuoteRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for QuoteRequest.
func (m *QuoteRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//NoRootPartyIDs is a non-required field for QuoteRequest.
func (m *QuoteRequest) NoRootPartyIDs() (*field.NoRootPartyIDs, error) {
	f := new(field.NoRootPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//PrivateQuote is a non-required field for QuoteRequest.
func (m *QuoteRequest) PrivateQuote() (*field.PrivateQuote, error) {
	f := new(field.PrivateQuote)
	err := m.Body.Get(f)
	return f, err
}

//RespondentType is a non-required field for QuoteRequest.
func (m *QuoteRequest) RespondentType() (*field.RespondentType, error) {
	f := new(field.RespondentType)
	err := m.Body.Get(f)
	return f, err
}

//PreTradeAnonymity is a non-required field for QuoteRequest.
func (m *QuoteRequest) PreTradeAnonymity() (*field.PreTradeAnonymity, error) {
	f := new(field.PreTradeAnonymity)
	err := m.Body.Get(f)
	return f, err
}

//BookingType is a non-required field for QuoteRequest.
func (m *QuoteRequest) BookingType() (*field.BookingType, error) {
	f := new(field.BookingType)
	err := m.Body.Get(f)
	return f, err
}

//OrderRestrictions is a non-required field for QuoteRequest.
func (m *QuoteRequest) OrderRestrictions() (*field.OrderRestrictions, error) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}
