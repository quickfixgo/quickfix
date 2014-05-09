package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//MarketDataRequestReject msg type = Y.
type MarketDataRequestReject struct {
	message.Message
}

//MarketDataRequestRejectBuilder builds MarketDataRequestReject messages.
type MarketDataRequestRejectBuilder struct {
	message.MessageBuilder
}

//CreateMarketDataRequestRejectBuilder returns an initialized MarketDataRequestRejectBuilder with specified required fields.
func CreateMarketDataRequestRejectBuilder(
	mdreqid *field.MDReqIDField) MarketDataRequestRejectBuilder {
	var builder MarketDataRequestRejectBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX43))
	builder.Header.Set(field.NewMsgType("Y"))
	builder.Body.Set(mdreqid)
	return builder
}

//MDReqID is a required field for MarketDataRequestReject.
func (m MarketDataRequestReject) MDReqID() (*field.MDReqIDField, errors.MessageRejectError) {
	f := &field.MDReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqID reads a MDReqID from MarketDataRequestReject.
func (m MarketDataRequestReject) GetMDReqID(f *field.MDReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDReqRejReason is a non-required field for MarketDataRequestReject.
func (m MarketDataRequestReject) MDReqRejReason() (*field.MDReqRejReasonField, errors.MessageRejectError) {
	f := &field.MDReqRejReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqRejReason reads a MDReqRejReason from MarketDataRequestReject.
func (m MarketDataRequestReject) GetMDReqRejReason(f *field.MDReqRejReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for MarketDataRequestReject.
func (m MarketDataRequestReject) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from MarketDataRequestReject.
func (m MarketDataRequestReject) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for MarketDataRequestReject.
func (m MarketDataRequestReject) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from MarketDataRequestReject.
func (m MarketDataRequestReject) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for MarketDataRequestReject.
func (m MarketDataRequestReject) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from MarketDataRequestReject.
func (m MarketDataRequestReject) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
