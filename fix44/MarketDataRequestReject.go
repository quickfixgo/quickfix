package fix44

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
func (m MarketDataRequestReject) MDReqID() (*field.MDReqID, errors.MessageRejectError) {
	f := new(field.MDReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqID reads a MDReqID from MarketDataRequestReject.
func (m MarketDataRequestReject) GetMDReqID(f *field.MDReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDReqRejReason is a non-required field for MarketDataRequestReject.
func (m MarketDataRequestReject) MDReqRejReason() (*field.MDReqRejReason, errors.MessageRejectError) {
	f := new(field.MDReqRejReason)
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqRejReason reads a MDReqRejReason from MarketDataRequestReject.
func (m MarketDataRequestReject) GetMDReqRejReason(f *field.MDReqRejReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAltMDSource is a non-required field for MarketDataRequestReject.
func (m MarketDataRequestReject) NoAltMDSource() (*field.NoAltMDSource, errors.MessageRejectError) {
	f := new(field.NoAltMDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetNoAltMDSource reads a NoAltMDSource from MarketDataRequestReject.
func (m MarketDataRequestReject) GetNoAltMDSource(f *field.NoAltMDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for MarketDataRequestReject.
func (m MarketDataRequestReject) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from MarketDataRequestReject.
func (m MarketDataRequestReject) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for MarketDataRequestReject.
func (m MarketDataRequestReject) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from MarketDataRequestReject.
func (m MarketDataRequestReject) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for MarketDataRequestReject.
func (m MarketDataRequestReject) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from MarketDataRequestReject.
func (m MarketDataRequestReject) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
