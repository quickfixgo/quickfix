package fix44

import (
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

//NewConfirmationRequestBuilder returns an initialized ConfirmationRequestBuilder with specified required fields.
func NewConfirmationRequestBuilder(
	confirmreqid field.ConfirmReqID,
	confirmtype field.ConfirmType,
	transacttime field.TransactTime) *ConfirmationRequestBuilder {
	builder := new(ConfirmationRequestBuilder)
	builder.Body.Set(confirmreqid)
	builder.Body.Set(confirmtype)
	builder.Body.Set(transacttime)
	return builder
}

//ConfirmReqID is a required field for ConfirmationRequest.
func (m *ConfirmationRequest) ConfirmReqID() (*field.ConfirmReqID, error) {
	f := new(field.ConfirmReqID)
	err := m.Body.Get(f)
	return f, err
}

//ConfirmType is a required field for ConfirmationRequest.
func (m *ConfirmationRequest) ConfirmType() (*field.ConfirmType, error) {
	f := new(field.ConfirmType)
	err := m.Body.Get(f)
	return f, err
}

//NoOrders is a non-required field for ConfirmationRequest.
func (m *ConfirmationRequest) NoOrders() (*field.NoOrders, error) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}

//AllocID is a non-required field for ConfirmationRequest.
func (m *ConfirmationRequest) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryAllocID is a non-required field for ConfirmationRequest.
func (m *ConfirmationRequest) SecondaryAllocID() (*field.SecondaryAllocID, error) {
	f := new(field.SecondaryAllocID)
	err := m.Body.Get(f)
	return f, err
}

//IndividualAllocID is a non-required field for ConfirmationRequest.
func (m *ConfirmationRequest) IndividualAllocID() (*field.IndividualAllocID, error) {
	f := new(field.IndividualAllocID)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a required field for ConfirmationRequest.
func (m *ConfirmationRequest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//AllocAccount is a non-required field for ConfirmationRequest.
func (m *ConfirmationRequest) AllocAccount() (*field.AllocAccount, error) {
	f := new(field.AllocAccount)
	err := m.Body.Get(f)
	return f, err
}

//AllocAcctIDSource is a non-required field for ConfirmationRequest.
func (m *ConfirmationRequest) AllocAcctIDSource() (*field.AllocAcctIDSource, error) {
	f := new(field.AllocAcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//AllocAccountType is a non-required field for ConfirmationRequest.
func (m *ConfirmationRequest) AllocAccountType() (*field.AllocAccountType, error) {
	f := new(field.AllocAccountType)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for ConfirmationRequest.
func (m *ConfirmationRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for ConfirmationRequest.
func (m *ConfirmationRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for ConfirmationRequest.
func (m *ConfirmationRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
