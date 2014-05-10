package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	confirmreqid *field.ConfirmReqIDField,
	confirmtype *field.ConfirmTypeField,
	transacttime *field.TransactTimeField) ConfirmationRequestBuilder {
	var builder ConfirmationRequestBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header.Set(field.NewMsgType("BH"))
	builder.Body.Set(confirmreqid)
	builder.Body.Set(confirmtype)
	builder.Body.Set(transacttime)
	return builder
}

//ConfirmReqID is a required field for ConfirmationRequest.
func (m ConfirmationRequest) ConfirmReqID() (*field.ConfirmReqIDField, errors.MessageRejectError) {
	f := &field.ConfirmReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmReqID reads a ConfirmReqID from ConfirmationRequest.
func (m ConfirmationRequest) GetConfirmReqID(f *field.ConfirmReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ConfirmType is a required field for ConfirmationRequest.
func (m ConfirmationRequest) ConfirmType() (*field.ConfirmTypeField, errors.MessageRejectError) {
	f := &field.ConfirmTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmType reads a ConfirmType from ConfirmationRequest.
func (m ConfirmationRequest) GetConfirmType(f *field.ConfirmTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) NoOrders() (*field.NoOrdersField, errors.MessageRejectError) {
	f := &field.NoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from ConfirmationRequest.
func (m ConfirmationRequest) GetNoOrders(f *field.NoOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) AllocID() (*field.AllocIDField, errors.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from ConfirmationRequest.
func (m ConfirmationRequest) GetAllocID(f *field.AllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryAllocID is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) SecondaryAllocID() (*field.SecondaryAllocIDField, errors.MessageRejectError) {
	f := &field.SecondaryAllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryAllocID reads a SecondaryAllocID from ConfirmationRequest.
func (m ConfirmationRequest) GetSecondaryAllocID(f *field.SecondaryAllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IndividualAllocID is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) IndividualAllocID() (*field.IndividualAllocIDField, errors.MessageRejectError) {
	f := &field.IndividualAllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIndividualAllocID reads a IndividualAllocID from ConfirmationRequest.
func (m ConfirmationRequest) GetIndividualAllocID(f *field.IndividualAllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for ConfirmationRequest.
func (m ConfirmationRequest) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ConfirmationRequest.
func (m ConfirmationRequest) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAccount is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) AllocAccount() (*field.AllocAccountField, errors.MessageRejectError) {
	f := &field.AllocAccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAccount reads a AllocAccount from ConfirmationRequest.
func (m ConfirmationRequest) GetAllocAccount(f *field.AllocAccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAcctIDSource is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) AllocAcctIDSource() (*field.AllocAcctIDSourceField, errors.MessageRejectError) {
	f := &field.AllocAcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAcctIDSource reads a AllocAcctIDSource from ConfirmationRequest.
func (m ConfirmationRequest) GetAllocAcctIDSource(f *field.AllocAcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAccountType is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) AllocAccountType() (*field.AllocAccountTypeField, errors.MessageRejectError) {
	f := &field.AllocAccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAccountType reads a AllocAccountType from ConfirmationRequest.
func (m ConfirmationRequest) GetAllocAccountType(f *field.AllocAccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ConfirmationRequest.
func (m ConfirmationRequest) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ConfirmationRequest.
func (m ConfirmationRequest) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ConfirmationRequest.
func (m ConfirmationRequest) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
