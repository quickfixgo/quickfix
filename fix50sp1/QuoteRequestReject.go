package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
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
	builder.Header.Set(field.BuildMsgType("AG"))
	builder.Body.Set(quotereqid)
	builder.Body.Set(quoterequestrejectreason)
	builder.Body.Set(norelatedsym)
	return builder
}

//QuoteReqID is a required field for QuoteRequestReject.
func (m QuoteRequestReject) QuoteReqID() (*field.QuoteReqID, errors.MessageRejectError) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteReqID reads a QuoteReqID from QuoteRequestReject.
func (m QuoteRequestReject) GetQuoteReqID(f *field.QuoteReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RFQReqID is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) RFQReqID() (*field.RFQReqID, errors.MessageRejectError) {
	f := new(field.RFQReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetRFQReqID reads a RFQReqID from QuoteRequestReject.
func (m QuoteRequestReject) GetRFQReqID(f *field.RFQReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteRequestRejectReason is a required field for QuoteRequestReject.
func (m QuoteRequestReject) QuoteRequestRejectReason() (*field.QuoteRequestRejectReason, errors.MessageRejectError) {
	f := new(field.QuoteRequestRejectReason)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteRequestRejectReason reads a QuoteRequestRejectReason from QuoteRequestReject.
func (m QuoteRequestReject) GetQuoteRequestRejectReason(f *field.QuoteRequestRejectReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a required field for QuoteRequestReject.
func (m QuoteRequestReject) NoRelatedSym() (*field.NoRelatedSym, errors.MessageRejectError) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from QuoteRequestReject.
func (m QuoteRequestReject) GetNoRelatedSym(f *field.NoRelatedSym) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from QuoteRequestReject.
func (m QuoteRequestReject) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from QuoteRequestReject.
func (m QuoteRequestReject) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from QuoteRequestReject.
func (m QuoteRequestReject) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRootPartyIDs is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) NoRootPartyIDs() (*field.NoRootPartyIDs, errors.MessageRejectError) {
	f := new(field.NoRootPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRootPartyIDs reads a NoRootPartyIDs from QuoteRequestReject.
func (m QuoteRequestReject) GetNoRootPartyIDs(f *field.NoRootPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrivateQuote is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) PrivateQuote() (*field.PrivateQuote, errors.MessageRejectError) {
	f := new(field.PrivateQuote)
	err := m.Body.Get(f)
	return f, err
}

//GetPrivateQuote reads a PrivateQuote from QuoteRequestReject.
func (m QuoteRequestReject) GetPrivateQuote(f *field.PrivateQuote) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RespondentType is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) RespondentType() (*field.RespondentType, errors.MessageRejectError) {
	f := new(field.RespondentType)
	err := m.Body.Get(f)
	return f, err
}

//GetRespondentType reads a RespondentType from QuoteRequestReject.
func (m QuoteRequestReject) GetRespondentType(f *field.RespondentType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreTradeAnonymity is a non-required field for QuoteRequestReject.
func (m QuoteRequestReject) PreTradeAnonymity() (*field.PreTradeAnonymity, errors.MessageRejectError) {
	f := new(field.PreTradeAnonymity)
	err := m.Body.Get(f)
	return f, err
}

//GetPreTradeAnonymity reads a PreTradeAnonymity from QuoteRequestReject.
func (m QuoteRequestReject) GetPreTradeAnonymity(f *field.PreTradeAnonymity) errors.MessageRejectError {
	return m.Body.Get(f)
}
