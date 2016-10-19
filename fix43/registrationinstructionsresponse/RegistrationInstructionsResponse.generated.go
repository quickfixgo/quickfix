package registrationinstructionsresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//RegistrationInstructionsResponse is the fix43 RegistrationInstructionsResponse type, MsgType = p
type RegistrationInstructionsResponse struct {
	fix43.Header
	*quickfix.Body
	fix43.Trailer
	Message *quickfix.Message
}

//FromMessage creates a RegistrationInstructionsResponse from a quickfix.Message instance
func FromMessage(m *quickfix.Message) RegistrationInstructionsResponse {
	return RegistrationInstructionsResponse{
		Header:  fix43.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix43.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m RegistrationInstructionsResponse) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a RegistrationInstructionsResponse initialized with the required fields for RegistrationInstructionsResponse
func New(registid field.RegistIDField, registtranstype field.RegistTransTypeField, registrefid field.RegistRefIDField, registstatus field.RegistStatusField) (m RegistrationInstructionsResponse) {
	m.Message = quickfix.NewMessage()
	m.Header = fix43.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("p"))
	m.Set(registid)
	m.Set(registtranstype)
	m.Set(registrefid)
	m.Set(registstatus)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg RegistrationInstructionsResponse, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "p", r
}

//SetAccount sets Account, Tag 1
func (m RegistrationInstructionsResponse) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m RegistrationInstructionsResponse) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m RegistrationInstructionsResponse) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetRegistRejReasonText sets RegistRejReasonText, Tag 496
func (m RegistrationInstructionsResponse) SetRegistRejReasonText(v string) {
	m.Set(field.NewRegistRejReasonText(v))
}

//SetRegistStatus sets RegistStatus, Tag 506
func (m RegistrationInstructionsResponse) SetRegistStatus(v enum.RegistStatus) {
	m.Set(field.NewRegistStatus(v))
}

//SetRegistRejReasonCode sets RegistRejReasonCode, Tag 507
func (m RegistrationInstructionsResponse) SetRegistRejReasonCode(v enum.RegistRejReasonCode) {
	m.Set(field.NewRegistRejReasonCode(v))
}

//SetRegistRefID sets RegistRefID, Tag 508
func (m RegistrationInstructionsResponse) SetRegistRefID(v string) {
	m.Set(field.NewRegistRefID(v))
}

//SetRegistID sets RegistID, Tag 513
func (m RegistrationInstructionsResponse) SetRegistID(v string) {
	m.Set(field.NewRegistID(v))
}

//SetRegistTransType sets RegistTransType, Tag 514
func (m RegistrationInstructionsResponse) SetRegistTransType(v enum.RegistTransType) {
	m.Set(field.NewRegistTransType(v))
}

//GetAccount gets Account, Tag 1
func (m RegistrationInstructionsResponse) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m RegistrationInstructionsResponse) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m RegistrationInstructionsResponse) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetRegistRejReasonText gets RegistRejReasonText, Tag 496
func (m RegistrationInstructionsResponse) GetRegistRejReasonText() (v string, err quickfix.MessageRejectError) {
	var f field.RegistRejReasonTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRegistStatus gets RegistStatus, Tag 506
func (m RegistrationInstructionsResponse) GetRegistStatus() (v enum.RegistStatus, err quickfix.MessageRejectError) {
	var f field.RegistStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRegistRejReasonCode gets RegistRejReasonCode, Tag 507
func (m RegistrationInstructionsResponse) GetRegistRejReasonCode() (v enum.RegistRejReasonCode, err quickfix.MessageRejectError) {
	var f field.RegistRejReasonCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRegistRefID gets RegistRefID, Tag 508
func (m RegistrationInstructionsResponse) GetRegistRefID() (v string, err quickfix.MessageRejectError) {
	var f field.RegistRefIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRegistID gets RegistID, Tag 513
func (m RegistrationInstructionsResponse) GetRegistID() (v string, err quickfix.MessageRejectError) {
	var f field.RegistIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRegistTransType gets RegistTransType, Tag 514
func (m RegistrationInstructionsResponse) GetRegistTransType() (v enum.RegistTransType, err quickfix.MessageRejectError) {
	var f field.RegistTransTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasAccount returns true if Account is present, Tag 1
func (m RegistrationInstructionsResponse) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m RegistrationInstructionsResponse) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m RegistrationInstructionsResponse) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasRegistRejReasonText returns true if RegistRejReasonText is present, Tag 496
func (m RegistrationInstructionsResponse) HasRegistRejReasonText() bool {
	return m.Has(tag.RegistRejReasonText)
}

//HasRegistStatus returns true if RegistStatus is present, Tag 506
func (m RegistrationInstructionsResponse) HasRegistStatus() bool {
	return m.Has(tag.RegistStatus)
}

//HasRegistRejReasonCode returns true if RegistRejReasonCode is present, Tag 507
func (m RegistrationInstructionsResponse) HasRegistRejReasonCode() bool {
	return m.Has(tag.RegistRejReasonCode)
}

//HasRegistRefID returns true if RegistRefID is present, Tag 508
func (m RegistrationInstructionsResponse) HasRegistRefID() bool {
	return m.Has(tag.RegistRefID)
}

//HasRegistID returns true if RegistID is present, Tag 513
func (m RegistrationInstructionsResponse) HasRegistID() bool {
	return m.Has(tag.RegistID)
}

//HasRegistTransType returns true if RegistTransType is present, Tag 514
func (m RegistrationInstructionsResponse) HasRegistTransType() bool {
	return m.Has(tag.RegistTransType)
}

//NoPartyIDs is a repeating group element, Tag 453
type NoPartyIDs struct {
	*quickfix.Group
}

//SetPartyID sets PartyID, Tag 448
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

//SetPartyIDSource sets PartyIDSource, Tag 447
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

//SetPartyRole sets PartyRole, Tag 452
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartyIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//GetPartyID gets PartyID, Tag 448
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyIDSource gets PartyIDSource, Tag 447
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyRole gets PartyRole, Tag 452
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartyIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasPartyID returns true if PartyID is present, Tag 448
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

//HasPartyIDSource returns true if PartyIDSource is present, Tag 447
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

//HasPartyRole returns true if PartyRole is present, Tag 452
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

//HasPartySubID returns true if PartySubID is present, Tag 523
func (m NoPartyIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

//NoPartyIDsRepeatingGroup is a repeating group, Tag 453
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), quickfix.GroupElement(tag.PartySubID)})}
}

//Add create and append a new NoPartyIDs to this group
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

//Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}
