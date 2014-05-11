package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	quotereqid *field.QuoteReqIDField,
	norelatedsym *field.NoRelatedSymField) QuoteRequestBuilder {
	var builder QuoteRequestBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("R"))
	builder.Body().Set(quotereqid)
	builder.Body().Set(norelatedsym)
	return builder
}

//QuoteReqID is a required field for QuoteRequest.
func (m QuoteRequest) QuoteReqID() (*field.QuoteReqIDField, errors.MessageRejectError) {
	f := &field.QuoteReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from QuoteRequest.
func (m QuoteRequest) GetQuoteReqID(f *field.QuoteReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RFQReqID is a non-required field for QuoteRequest.
func (m QuoteRequest) RFQReqID() (*field.RFQReqIDField, errors.MessageRejectError) {
	f := &field.RFQReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRFQReqID reads a RFQReqID from QuoteRequest.
func (m QuoteRequest) GetRFQReqID(f *field.RFQReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for QuoteRequest.
func (m QuoteRequest) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from QuoteRequest.
func (m QuoteRequest) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for QuoteRequest.
func (m QuoteRequest) OrderCapacity() (*field.OrderCapacityField, errors.MessageRejectError) {
	f := &field.OrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from QuoteRequest.
func (m QuoteRequest) GetOrderCapacity(f *field.OrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a required field for QuoteRequest.
func (m QuoteRequest) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from QuoteRequest.
func (m QuoteRequest) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for QuoteRequest.
func (m QuoteRequest) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from QuoteRequest.
func (m QuoteRequest) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for QuoteRequest.
func (m QuoteRequest) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from QuoteRequest.
func (m QuoteRequest) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for QuoteRequest.
func (m QuoteRequest) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from QuoteRequest.
func (m QuoteRequest) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRootPartyIDs is a non-required field for QuoteRequest.
func (m QuoteRequest) NoRootPartyIDs() (*field.NoRootPartyIDsField, errors.MessageRejectError) {
	f := &field.NoRootPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRootPartyIDs reads a NoRootPartyIDs from QuoteRequest.
func (m QuoteRequest) GetNoRootPartyIDs(f *field.NoRootPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrivateQuote is a non-required field for QuoteRequest.
func (m QuoteRequest) PrivateQuote() (*field.PrivateQuoteField, errors.MessageRejectError) {
	f := &field.PrivateQuoteField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrivateQuote reads a PrivateQuote from QuoteRequest.
func (m QuoteRequest) GetPrivateQuote(f *field.PrivateQuoteField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RespondentType is a non-required field for QuoteRequest.
func (m QuoteRequest) RespondentType() (*field.RespondentTypeField, errors.MessageRejectError) {
	f := &field.RespondentTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRespondentType reads a RespondentType from QuoteRequest.
func (m QuoteRequest) GetRespondentType(f *field.RespondentTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreTradeAnonymity is a non-required field for QuoteRequest.
func (m QuoteRequest) PreTradeAnonymity() (*field.PreTradeAnonymityField, errors.MessageRejectError) {
	f := &field.PreTradeAnonymityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreTradeAnonymity reads a PreTradeAnonymity from QuoteRequest.
func (m QuoteRequest) GetPreTradeAnonymity(f *field.PreTradeAnonymityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingType is a non-required field for QuoteRequest.
func (m QuoteRequest) BookingType() (*field.BookingTypeField, errors.MessageRejectError) {
	f := &field.BookingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingType reads a BookingType from QuoteRequest.
func (m QuoteRequest) GetBookingType(f *field.BookingTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for QuoteRequest.
func (m QuoteRequest) OrderRestrictions() (*field.OrderRestrictionsField, errors.MessageRejectError) {
	f := &field.OrderRestrictionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from QuoteRequest.
func (m QuoteRequest) GetOrderRestrictions(f *field.OrderRestrictionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
