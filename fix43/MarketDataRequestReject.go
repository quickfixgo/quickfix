package fix43

import (
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

//NewMarketDataRequestRejectBuilder returns an initialized MarketDataRequestRejectBuilder with specified required fields.
func NewMarketDataRequestRejectBuilder(
	mdreqid field.MDReqID) *MarketDataRequestRejectBuilder {
	builder := new(MarketDataRequestRejectBuilder)
	builder.Body.Set(mdreqid)
	return builder
}

//MDReqID is a required field for MarketDataRequestReject.
func (m *MarketDataRequestReject) MDReqID() (*field.MDReqID, error) {
	f := new(field.MDReqID)
	err := m.Body.Get(f)
	return f, err
}

//MDReqRejReason is a non-required field for MarketDataRequestReject.
func (m *MarketDataRequestReject) MDReqRejReason() (*field.MDReqRejReason, error) {
	f := new(field.MDReqRejReason)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for MarketDataRequestReject.
func (m *MarketDataRequestReject) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for MarketDataRequestReject.
func (m *MarketDataRequestReject) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for MarketDataRequestReject.
func (m *MarketDataRequestReject) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
