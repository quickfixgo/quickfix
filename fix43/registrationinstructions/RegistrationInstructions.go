//Package registrationinstructions msg type = o.
package registrationinstructions

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix/enum"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a RegistrationInstructions wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//RegistID is a required field for RegistrationInstructions.
func (m Message) RegistID() (*field.RegistIDField, quickfix.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from RegistrationInstructions.
func (m Message) GetRegistID(f *field.RegistIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RegistTransType is a required field for RegistrationInstructions.
func (m Message) RegistTransType() (*field.RegistTransTypeField, quickfix.MessageRejectError) {
	f := &field.RegistTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistTransType reads a RegistTransType from RegistrationInstructions.
func (m Message) GetRegistTransType(f *field.RegistTransTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RegistRefID is a required field for RegistrationInstructions.
func (m Message) RegistRefID() (*field.RegistRefIDField, quickfix.MessageRejectError) {
	f := &field.RegistRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistRefID reads a RegistRefID from RegistrationInstructions.
func (m Message) GetRegistRefID(f *field.RegistRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for RegistrationInstructions.
func (m Message) ClOrdID() (*field.ClOrdIDField, quickfix.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from RegistrationInstructions.
func (m Message) GetClOrdID(f *field.ClOrdIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for RegistrationInstructions.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from RegistrationInstructions.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for RegistrationInstructions.
func (m Message) Account() (*field.AccountField, quickfix.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from RegistrationInstructions.
func (m Message) GetAccount(f *field.AccountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RegistAcctType is a non-required field for RegistrationInstructions.
func (m Message) RegistAcctType() (*field.RegistAcctTypeField, quickfix.MessageRejectError) {
	f := &field.RegistAcctTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistAcctType reads a RegistAcctType from RegistrationInstructions.
func (m Message) GetRegistAcctType(f *field.RegistAcctTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TaxAdvantageType is a non-required field for RegistrationInstructions.
func (m Message) TaxAdvantageType() (*field.TaxAdvantageTypeField, quickfix.MessageRejectError) {
	f := &field.TaxAdvantageTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTaxAdvantageType reads a TaxAdvantageType from RegistrationInstructions.
func (m Message) GetTaxAdvantageType(f *field.TaxAdvantageTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OwnershipType is a non-required field for RegistrationInstructions.
func (m Message) OwnershipType() (*field.OwnershipTypeField, quickfix.MessageRejectError) {
	f := &field.OwnershipTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOwnershipType reads a OwnershipType from RegistrationInstructions.
func (m Message) GetOwnershipType(f *field.OwnershipTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRegistDtls is a non-required field for RegistrationInstructions.
func (m Message) NoRegistDtls() (*field.NoRegistDtlsField, quickfix.MessageRejectError) {
	f := &field.NoRegistDtlsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRegistDtls reads a NoRegistDtls from RegistrationInstructions.
func (m Message) GetNoRegistDtls(f *field.NoRegistDtlsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoDistribInsts is a non-required field for RegistrationInstructions.
func (m Message) NoDistribInsts() (*field.NoDistribInstsField, quickfix.MessageRejectError) {
	f := &field.NoDistribInstsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoDistribInsts reads a NoDistribInsts from RegistrationInstructions.
func (m Message) GetNoDistribInsts(f *field.NoDistribInstsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for RegistrationInstructions.
func New(
	registid *field.RegistIDField,
	registtranstype *field.RegistTransTypeField,
	registrefid *field.RegistRefIDField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIX43))
	builder.Header.Set(field.NewMsgType("o"))
	builder.Body.Set(registid)
	builder.Body.Set(registtranstype)
	builder.Body.Set(registrefid)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.BeginStringFIX43, "o", r
}
