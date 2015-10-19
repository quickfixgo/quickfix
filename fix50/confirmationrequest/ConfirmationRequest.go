//Package confirmationrequest msg type = BH.
package confirmationrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a ConfirmationRequest wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//ConfirmReqID is a required field for ConfirmationRequest.
func (m Message) ConfirmReqID() (*field.ConfirmReqIDField, quickfix.MessageRejectError) {
	f := &field.ConfirmReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmReqID reads a ConfirmReqID from ConfirmationRequest.
func (m Message) GetConfirmReqID(f *field.ConfirmReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ConfirmType is a required field for ConfirmationRequest.
func (m Message) ConfirmType() (*field.ConfirmTypeField, quickfix.MessageRejectError) {
	f := &field.ConfirmTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConfirmType reads a ConfirmType from ConfirmationRequest.
func (m Message) GetConfirmType(f *field.ConfirmTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a non-required field for ConfirmationRequest.
func (m Message) NoOrders() (*field.NoOrdersField, quickfix.MessageRejectError) {
	f := &field.NoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from ConfirmationRequest.
func (m Message) GetNoOrders(f *field.NoOrdersField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for ConfirmationRequest.
func (m Message) AllocID() (*field.AllocIDField, quickfix.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from ConfirmationRequest.
func (m Message) GetAllocID(f *field.AllocIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryAllocID is a non-required field for ConfirmationRequest.
func (m Message) SecondaryAllocID() (*field.SecondaryAllocIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryAllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryAllocID reads a SecondaryAllocID from ConfirmationRequest.
func (m Message) GetSecondaryAllocID(f *field.SecondaryAllocIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//IndividualAllocID is a non-required field for ConfirmationRequest.
func (m Message) IndividualAllocID() (*field.IndividualAllocIDField, quickfix.MessageRejectError) {
	f := &field.IndividualAllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIndividualAllocID reads a IndividualAllocID from ConfirmationRequest.
func (m Message) GetIndividualAllocID(f *field.IndividualAllocIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for ConfirmationRequest.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ConfirmationRequest.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAccount is a non-required field for ConfirmationRequest.
func (m Message) AllocAccount() (*field.AllocAccountField, quickfix.MessageRejectError) {
	f := &field.AllocAccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAccount reads a AllocAccount from ConfirmationRequest.
func (m Message) GetAllocAccount(f *field.AllocAccountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAcctIDSource is a non-required field for ConfirmationRequest.
func (m Message) AllocAcctIDSource() (*field.AllocAcctIDSourceField, quickfix.MessageRejectError) {
	f := &field.AllocAcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAcctIDSource reads a AllocAcctIDSource from ConfirmationRequest.
func (m Message) GetAllocAcctIDSource(f *field.AllocAcctIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAccountType is a non-required field for ConfirmationRequest.
func (m Message) AllocAccountType() (*field.AllocAccountTypeField, quickfix.MessageRejectError) {
	f := &field.AllocAccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAccountType reads a AllocAccountType from ConfirmationRequest.
func (m Message) GetAllocAccountType(f *field.AllocAccountTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ConfirmationRequest.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ConfirmationRequest.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ConfirmationRequest.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ConfirmationRequest.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ConfirmationRequest.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ConfirmationRequest.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds ConfirmationRequest messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for ConfirmationRequest.
func Builder(
	confirmreqid *field.ConfirmReqIDField,
	confirmtype *field.ConfirmTypeField,
	transacttime *field.TransactTimeField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header.Set(field.NewMsgType("BH"))
	builder.Body.Set(confirmreqid)
	builder.Body.Set(confirmtype)
	builder.Body.Set(transacttime)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX50, "BH", r
}
