//Package registrationinstructionsresponse msg type = p.
package registrationinstructionsresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a RegistrationInstructionsResponse wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//RegistID is a required field for RegistrationInstructionsResponse.
func (m Message) RegistID() (*field.RegistIDField, quickfix.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from RegistrationInstructionsResponse.
func (m Message) GetRegistID(f *field.RegistIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RegistTransType is a required field for RegistrationInstructionsResponse.
func (m Message) RegistTransType() (*field.RegistTransTypeField, quickfix.MessageRejectError) {
	f := &field.RegistTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistTransType reads a RegistTransType from RegistrationInstructionsResponse.
func (m Message) GetRegistTransType(f *field.RegistTransTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RegistRefID is a required field for RegistrationInstructionsResponse.
func (m Message) RegistRefID() (*field.RegistRefIDField, quickfix.MessageRejectError) {
	f := &field.RegistRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistRefID reads a RegistRefID from RegistrationInstructionsResponse.
func (m Message) GetRegistRefID(f *field.RegistRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for RegistrationInstructionsResponse.
func (m Message) ClOrdID() (*field.ClOrdIDField, quickfix.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from RegistrationInstructionsResponse.
func (m Message) GetClOrdID(f *field.ClOrdIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for RegistrationInstructionsResponse.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from RegistrationInstructionsResponse.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for RegistrationInstructionsResponse.
func (m Message) Account() (*field.AccountField, quickfix.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from RegistrationInstructionsResponse.
func (m Message) GetAccount(f *field.AccountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for RegistrationInstructionsResponse.
func (m Message) AcctIDSource() (*field.AcctIDSourceField, quickfix.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from RegistrationInstructionsResponse.
func (m Message) GetAcctIDSource(f *field.AcctIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RegistStatus is a required field for RegistrationInstructionsResponse.
func (m Message) RegistStatus() (*field.RegistStatusField, quickfix.MessageRejectError) {
	f := &field.RegistStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistStatus reads a RegistStatus from RegistrationInstructionsResponse.
func (m Message) GetRegistStatus(f *field.RegistStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RegistRejReasonCode is a non-required field for RegistrationInstructionsResponse.
func (m Message) RegistRejReasonCode() (*field.RegistRejReasonCodeField, quickfix.MessageRejectError) {
	f := &field.RegistRejReasonCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistRejReasonCode reads a RegistRejReasonCode from RegistrationInstructionsResponse.
func (m Message) GetRegistRejReasonCode(f *field.RegistRejReasonCodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RegistRejReasonText is a non-required field for RegistrationInstructionsResponse.
func (m Message) RegistRejReasonText() (*field.RegistRejReasonTextField, quickfix.MessageRejectError) {
	f := &field.RegistRejReasonTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistRejReasonText reads a RegistRejReasonText from RegistrationInstructionsResponse.
func (m Message) GetRegistRejReasonText(f *field.RegistRejReasonTextField) quickfix.MessageRejectError {
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
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("p"))
	builder.Body().Set(registid)
	builder.Body().Set(registtranstype)
	builder.Body().Set(registrefid)
	builder.Body().Set(registstatus)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "p", r
}
