package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ConfirmationRequest msg type = BH.
type ConfirmationRequest struct {
	message.Message
}

//ConfirmationRequestBuilder builds ConfirmationRequest messages.
type ConfirmationRequestBuilder struct {
	message.MessageBuilder
}

//CreateConfirmationRequestBuilder returns an initialized ConfirmationRequestBuilder with specified required fields.
func CreateConfirmationRequestBuilder(
	confirmreqid field.ConfirmReqID,
	confirmtype field.ConfirmType,
	transacttime field.TransactTime) ConfirmationRequestBuilder {
	var builder ConfirmationRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(confirmreqid)
	builder.Body.Set(confirmtype)
	builder.Body.Set(transacttime)
	return builder
}

//ConfirmReqID is a required field for ConfirmationRequest.
func (m ConfirmationRequest) ConfirmReqID() (field.ConfirmReqID, errors.MessageRejectError) {
	var f field.ConfirmReqID
	err := m.Body.Get(&f)
	return f, err
}

//ConfirmType is a required field for ConfirmationRequest.
func (m ConfirmationRequest) ConfirmType() (field.ConfirmType, errors.MessageRejectError) {
	var f field.ConfirmType
	err := m.Body.Get(&f)
	return f, err
}

//NoOrders is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) NoOrders() (field.NoOrders, errors.MessageRejectError) {
	var f field.NoOrders
	err := m.Body.Get(&f)
	return f, err
}

//AllocID is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) AllocID() (field.AllocID, errors.MessageRejectError) {
	var f field.AllocID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryAllocID is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) SecondaryAllocID() (field.SecondaryAllocID, errors.MessageRejectError) {
	var f field.SecondaryAllocID
	err := m.Body.Get(&f)
	return f, err
}

//IndividualAllocID is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) IndividualAllocID() (field.IndividualAllocID, errors.MessageRejectError) {
	var f field.IndividualAllocID
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for ConfirmationRequest.
func (m ConfirmationRequest) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//AllocAccount is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) AllocAccount() (field.AllocAccount, errors.MessageRejectError) {
	var f field.AllocAccount
	err := m.Body.Get(&f)
	return f, err
}

//AllocAcctIDSource is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) AllocAcctIDSource() (field.AllocAcctIDSource, errors.MessageRejectError) {
	var f field.AllocAcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AllocAccountType is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) AllocAccountType() (field.AllocAccountType, errors.MessageRejectError) {
	var f field.AllocAccountType
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
