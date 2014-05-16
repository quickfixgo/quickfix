//Package registrationinstructionsresponse msg type = p.
package registrationinstructionsresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a RegistrationInstructionsResponse wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//RegistID is a required field for RegistrationInstructionsResponse.
func (m Message) RegistID() (*field.RegistIDField, errors.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from RegistrationInstructionsResponse.
func (m Message) GetRegistID(f *field.RegistIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistTransType is a required field for RegistrationInstructionsResponse.
func (m Message) RegistTransType() (*field.RegistTransTypeField, errors.MessageRejectError) {
	f := &field.RegistTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistTransType reads a RegistTransType from RegistrationInstructionsResponse.
func (m Message) GetRegistTransType(f *field.RegistTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistRefID is a required field for RegistrationInstructionsResponse.
func (m Message) RegistRefID() (*field.RegistRefIDField, errors.MessageRejectError) {
	f := &field.RegistRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistRefID reads a RegistRefID from RegistrationInstructionsResponse.
func (m Message) GetRegistRefID(f *field.RegistRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for RegistrationInstructionsResponse.
func (m Message) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from RegistrationInstructionsResponse.
func (m Message) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for RegistrationInstructionsResponse.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from RegistrationInstructionsResponse.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for RegistrationInstructionsResponse.
func (m Message) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from RegistrationInstructionsResponse.
func (m Message) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistStatus is a required field for RegistrationInstructionsResponse.
func (m Message) RegistStatus() (*field.RegistStatusField, errors.MessageRejectError) {
	f := &field.RegistStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistStatus reads a RegistStatus from RegistrationInstructionsResponse.
func (m Message) GetRegistStatus(f *field.RegistStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistRejReasonCode is a non-required field for RegistrationInstructionsResponse.
func (m Message) RegistRejReasonCode() (*field.RegistRejReasonCodeField, errors.MessageRejectError) {
	f := &field.RegistRejReasonCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistRejReasonCode reads a RegistRejReasonCode from RegistrationInstructionsResponse.
func (m Message) GetRegistRejReasonCode(f *field.RegistRejReasonCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistRejReasonText is a non-required field for RegistrationInstructionsResponse.
func (m Message) RegistRejReasonText() (*field.RegistRejReasonTextField, errors.MessageRejectError) {
	f := &field.RegistRejReasonTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistRejReasonText reads a RegistRejReasonText from RegistrationInstructionsResponse.
func (m Message) GetRegistRejReasonText(f *field.RegistRejReasonTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds RegistrationInstructionsResponse messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for RegistrationInstructionsResponse.
func Builder(
	registid *field.RegistIDField,
	registtranstype *field.RegistTransTypeField,
	registrefid *field.RegistRefIDField,
	registstatus *field.RegistStatusField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX43))
	builder.Header().Set(field.NewMsgType("p"))
	builder.Body().Set(registid)
	builder.Body().Set(registtranstype)
	builder.Body().Set(registrefid)
	builder.Body().Set(registstatus)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX43, "p", r
}
