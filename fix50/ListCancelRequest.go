package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ListCancelRequest msg type = K.
type ListCancelRequest struct {
	message.Message
}

//ListCancelRequestBuilder builds ListCancelRequest messages.
type ListCancelRequestBuilder struct {
	message.MessageBuilder
}

//CreateListCancelRequestBuilder returns an initialized ListCancelRequestBuilder with specified required fields.
func CreateListCancelRequestBuilder(
	listid field.ListID,
	transacttime field.TransactTime) ListCancelRequestBuilder {
	var builder ListCancelRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(listid)
	builder.Body.Set(transacttime)
	return builder
}

//ListID is a required field for ListCancelRequest.
func (m ListCancelRequest) ListID() (field.ListID, error) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for ListCancelRequest.
func (m ListCancelRequest) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//TradeOriginationDate is a non-required field for ListCancelRequest.
func (m ListCancelRequest) TradeOriginationDate() (field.TradeOriginationDate, error) {
	var f field.TradeOriginationDate
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for ListCancelRequest.
func (m ListCancelRequest) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for ListCancelRequest.
func (m ListCancelRequest) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for ListCancelRequest.
func (m ListCancelRequest) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for ListCancelRequest.
func (m ListCancelRequest) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for ListCancelRequest.
func (m ListCancelRequest) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}
