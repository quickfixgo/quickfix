package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//RFQRequest msg type = AH.
type RFQRequest struct {
	message.Message
}

//RFQRequestBuilder builds RFQRequest messages.
type RFQRequestBuilder struct {
	message.MessageBuilder
}

//CreateRFQRequestBuilder returns an initialized RFQRequestBuilder with specified required fields.
func CreateRFQRequestBuilder(
	rfqreqid *field.RFQReqIDField,
	norelatedsym *field.NoRelatedSymField) RFQRequestBuilder {
	var builder RFQRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("AH"))
	builder.Body.Set(rfqreqid)
	builder.Body.Set(norelatedsym)
	return builder
}

//RFQReqID is a required field for RFQRequest.
func (m RFQRequest) RFQReqID() (*field.RFQReqIDField, errors.MessageRejectError) {
	f := &field.RFQReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRFQReqID reads a RFQReqID from RFQRequest.
func (m RFQRequest) GetRFQReqID(f *field.RFQReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a required field for RFQRequest.
func (m RFQRequest) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from RFQRequest.
func (m RFQRequest) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for RFQRequest.
func (m RFQRequest) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, errors.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from RFQRequest.
func (m RFQRequest) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrivateQuote is a non-required field for RFQRequest.
func (m RFQRequest) PrivateQuote() (*field.PrivateQuoteField, errors.MessageRejectError) {
	f := &field.PrivateQuoteField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrivateQuote reads a PrivateQuote from RFQRequest.
func (m RFQRequest) GetPrivateQuote(f *field.PrivateQuoteField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for RFQRequest.
func (m RFQRequest) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from RFQRequest.
func (m RFQRequest) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
