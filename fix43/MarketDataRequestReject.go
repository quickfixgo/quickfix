package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
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
	mdreqid field.MDReqID) MarketDataRequestRejectBuilder {
	var builder MarketDataRequestRejectBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(mdreqid)
	return builder
}

//MDReqID is a required field for MarketDataRequestReject.
func (m MarketDataRequestReject) MDReqID() (field.MDReqID, errors.MessageRejectError) {
	var f field.MDReqID
	err := m.Body.Get(&f)
	return f, err
}

//MDReqRejReason is a non-required field for MarketDataRequestReject.
func (m MarketDataRequestReject) MDReqRejReason() (field.MDReqRejReason, errors.MessageRejectError) {
	var f field.MDReqRejReason
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for MarketDataRequestReject.
func (m MarketDataRequestReject) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for MarketDataRequestReject.
func (m MarketDataRequestReject) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for MarketDataRequestReject.
func (m MarketDataRequestReject) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
