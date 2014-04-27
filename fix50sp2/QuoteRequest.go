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
	quotereqid field.QuoteReqID,
	norelatedsym field.NoRelatedSym) QuoteRequestBuilder {
	var builder QuoteRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.BuildMsgType("R"))
	builder.Body.Set(quotereqid)
	builder.Body.Set(norelatedsym)
	return builder
}

//QuoteReqID is a required field for QuoteRequest.
func (m QuoteRequest) QuoteReqID() (*field.QuoteReqID, errors.MessageRejectError) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from QuoteRequest.
func (m QuoteRequest) GetQuoteReqID(f *field.QuoteReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RFQReqID is a non-required field for QuoteRequest.
func (m QuoteRequest) RFQReqID() (*field.RFQReqID, errors.MessageRejectError) {
	f := new(field.RFQReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetRFQReqID reads a RFQReqID from QuoteRequest.
func (m QuoteRequest) GetRFQReqID(f *field.RFQReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for QuoteRequest.
func (m QuoteRequest) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from QuoteRequest.
func (m QuoteRequest) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for QuoteRequest.
func (m QuoteRequest) OrderCapacity() (*field.OrderCapacity, errors.MessageRejectError) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from QuoteRequest.
func (m QuoteRequest) GetOrderCapacity(f *field.OrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a required field for QuoteRequest.
func (m QuoteRequest) NoRelatedSym() (*field.NoRelatedSym, errors.MessageRejectError) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from QuoteRequest.
func (m QuoteRequest) GetNoRelatedSym(f *field.NoRelatedSym) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for QuoteRequest.
func (m QuoteRequest) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from QuoteRequest.
func (m QuoteRequest) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for QuoteRequest.
func (m QuoteRequest) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from QuoteRequest.
func (m QuoteRequest) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for QuoteRequest.
func (m QuoteRequest) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from QuoteRequest.
func (m QuoteRequest) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRootPartyIDs is a non-required field for QuoteRequest.
func (m QuoteRequest) NoRootPartyIDs() (*field.NoRootPartyIDs, errors.MessageRejectError) {
	f := new(field.NoRootPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRootPartyIDs reads a NoRootPartyIDs from QuoteRequest.
func (m QuoteRequest) GetNoRootPartyIDs(f *field.NoRootPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrivateQuote is a non-required field for QuoteRequest.
func (m QuoteRequest) PrivateQuote() (*field.PrivateQuote, errors.MessageRejectError) {
	f := new(field.PrivateQuote)
	err := m.Body.Get(f)
	return f, err
}

//GetPrivateQuote reads a PrivateQuote from QuoteRequest.
func (m QuoteRequest) GetPrivateQuote(f *field.PrivateQuote) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RespondentType is a non-required field for QuoteRequest.
func (m QuoteRequest) RespondentType() (*field.RespondentType, errors.MessageRejectError) {
	f := new(field.RespondentType)
	err := m.Body.Get(f)
	return f, err
}

//GetRespondentType reads a RespondentType from QuoteRequest.
func (m QuoteRequest) GetRespondentType(f *field.RespondentType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreTradeAnonymity is a non-required field for QuoteRequest.
func (m QuoteRequest) PreTradeAnonymity() (*field.PreTradeAnonymity, errors.MessageRejectError) {
	f := new(field.PreTradeAnonymity)
	err := m.Body.Get(f)
	return f, err
}

//GetPreTradeAnonymity reads a PreTradeAnonymity from QuoteRequest.
func (m QuoteRequest) GetPreTradeAnonymity(f *field.PreTradeAnonymity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingType is a non-required field for QuoteRequest.
func (m QuoteRequest) BookingType() (*field.BookingType, errors.MessageRejectError) {
	f := new(field.BookingType)
	err := m.Body.Get(f)
	return f, err
}

//GetBookingType reads a BookingType from QuoteRequest.
func (m QuoteRequest) GetBookingType(f *field.BookingType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for QuoteRequest.
func (m QuoteRequest) OrderRestrictions() (*field.OrderRestrictions, errors.MessageRejectError) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from QuoteRequest.
func (m QuoteRequest) GetOrderRestrictions(f *field.OrderRestrictions) errors.MessageRejectError {
	return m.Body.Get(f)
}
