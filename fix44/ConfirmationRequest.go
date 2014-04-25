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
	builder.Header.Set(field.BuildMsgType("BH"))
	builder.Body.Set(confirmreqid)
	builder.Body.Set(confirmtype)
	builder.Body.Set(transacttime)
	return builder
}

//ConfirmReqID is a required field for ConfirmationRequest.
func (m ConfirmationRequest) ConfirmReqID() (*field.ConfirmReqID, errors.MessageRejectError) {
	f := new(field.ConfirmReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmReqID reads a ConfirmReqID from ConfirmationRequest.
func (m ConfirmationRequest) GetConfirmReqID(f *field.ConfirmReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ConfirmType is a required field for ConfirmationRequest.
func (m ConfirmationRequest) ConfirmType() (*field.ConfirmType, errors.MessageRejectError) {
	f := new(field.ConfirmType)
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmType reads a ConfirmType from ConfirmationRequest.
func (m ConfirmationRequest) GetConfirmType(f *field.ConfirmType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) NoOrders() (*field.NoOrders, errors.MessageRejectError) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from ConfirmationRequest.
func (m ConfirmationRequest) GetNoOrders(f *field.NoOrders) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) AllocID() (*field.AllocID, errors.MessageRejectError) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from ConfirmationRequest.
func (m ConfirmationRequest) GetAllocID(f *field.AllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryAllocID is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) SecondaryAllocID() (*field.SecondaryAllocID, errors.MessageRejectError) {
	f := new(field.SecondaryAllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryAllocID reads a SecondaryAllocID from ConfirmationRequest.
func (m ConfirmationRequest) GetSecondaryAllocID(f *field.SecondaryAllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IndividualAllocID is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) IndividualAllocID() (*field.IndividualAllocID, errors.MessageRejectError) {
	f := new(field.IndividualAllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetIndividualAllocID reads a IndividualAllocID from ConfirmationRequest.
func (m ConfirmationRequest) GetIndividualAllocID(f *field.IndividualAllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for ConfirmationRequest.
func (m ConfirmationRequest) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ConfirmationRequest.
func (m ConfirmationRequest) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAccount is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) AllocAccount() (*field.AllocAccount, errors.MessageRejectError) {
	f := new(field.AllocAccount)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAccount reads a AllocAccount from ConfirmationRequest.
func (m ConfirmationRequest) GetAllocAccount(f *field.AllocAccount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAcctIDSource is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) AllocAcctIDSource() (*field.AllocAcctIDSource, errors.MessageRejectError) {
	f := new(field.AllocAcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAcctIDSource reads a AllocAcctIDSource from ConfirmationRequest.
func (m ConfirmationRequest) GetAllocAcctIDSource(f *field.AllocAcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAccountType is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) AllocAccountType() (*field.AllocAccountType, errors.MessageRejectError) {
	f := new(field.AllocAccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAccountType reads a AllocAccountType from ConfirmationRequest.
func (m ConfirmationRequest) GetAllocAccountType(f *field.AllocAccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ConfirmationRequest.
func (m ConfirmationRequest) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ConfirmationRequest.
func (m ConfirmationRequest) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ConfirmationRequest.
func (m ConfirmationRequest) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ConfirmationRequest.
func (m ConfirmationRequest) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
