package registrationinstructions

import (
	"github.com/shopspring/decimal"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/tag"
)

//RegistrationInstructions is the fix44 RegistrationInstructions type, MsgType = o
type RegistrationInstructions struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

//FromMessage creates a RegistrationInstructions from a quickfix.Message instance
func FromMessage(m *quickfix.Message) RegistrationInstructions {
	return RegistrationInstructions{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m RegistrationInstructions) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a RegistrationInstructions initialized with the required fields for RegistrationInstructions
func New(registid field.RegistIDField, registtranstype field.RegistTransTypeField, registrefid field.RegistRefIDField) (m RegistrationInstructions) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("o"))
	m.Set(registid)
	m.Set(registtranstype)
	m.Set(registrefid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg RegistrationInstructions, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "o", r
}

//SetAccount sets Account, Tag 1
func (m RegistrationInstructions) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m RegistrationInstructions) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m RegistrationInstructions) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoRegistDtls sets NoRegistDtls, Tag 473
func (m RegistrationInstructions) SetNoRegistDtls(f NoRegistDtlsRepeatingGroup) {
	m.SetGroup(f)
}

//SetRegistAcctType sets RegistAcctType, Tag 493
func (m RegistrationInstructions) SetRegistAcctType(v string) {
	m.Set(field.NewRegistAcctType(v))
}

//SetTaxAdvantageType sets TaxAdvantageType, Tag 495
func (m RegistrationInstructions) SetTaxAdvantageType(v enum.TaxAdvantageType) {
	m.Set(field.NewTaxAdvantageType(v))
}

//SetRegistRefID sets RegistRefID, Tag 508
func (m RegistrationInstructions) SetRegistRefID(v string) {
	m.Set(field.NewRegistRefID(v))
}

//SetNoDistribInsts sets NoDistribInsts, Tag 510
func (m RegistrationInstructions) SetNoDistribInsts(f NoDistribInstsRepeatingGroup) {
	m.SetGroup(f)
}

//SetRegistID sets RegistID, Tag 513
func (m RegistrationInstructions) SetRegistID(v string) {
	m.Set(field.NewRegistID(v))
}

//SetRegistTransType sets RegistTransType, Tag 514
func (m RegistrationInstructions) SetRegistTransType(v enum.RegistTransType) {
	m.Set(field.NewRegistTransType(v))
}

//SetOwnershipType sets OwnershipType, Tag 517
func (m RegistrationInstructions) SetOwnershipType(v enum.OwnershipType) {
	m.Set(field.NewOwnershipType(v))
}

//SetAcctIDSource sets AcctIDSource, Tag 660
func (m RegistrationInstructions) SetAcctIDSource(v enum.AcctIDSource) {
	m.Set(field.NewAcctIDSource(v))
}

//GetAccount gets Account, Tag 1
func (m RegistrationInstructions) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m RegistrationInstructions) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m RegistrationInstructions) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoRegistDtls gets NoRegistDtls, Tag 473
func (m RegistrationInstructions) GetNoRegistDtls() (NoRegistDtlsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRegistDtlsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetRegistAcctType gets RegistAcctType, Tag 493
func (m RegistrationInstructions) GetRegistAcctType() (v string, err quickfix.MessageRejectError) {
	var f field.RegistAcctTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTaxAdvantageType gets TaxAdvantageType, Tag 495
func (m RegistrationInstructions) GetTaxAdvantageType() (v enum.TaxAdvantageType, err quickfix.MessageRejectError) {
	var f field.TaxAdvantageTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRegistRefID gets RegistRefID, Tag 508
func (m RegistrationInstructions) GetRegistRefID() (v string, err quickfix.MessageRejectError) {
	var f field.RegistRefIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoDistribInsts gets NoDistribInsts, Tag 510
func (m RegistrationInstructions) GetNoDistribInsts() (NoDistribInstsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoDistribInstsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetRegistID gets RegistID, Tag 513
func (m RegistrationInstructions) GetRegistID() (v string, err quickfix.MessageRejectError) {
	var f field.RegistIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRegistTransType gets RegistTransType, Tag 514
func (m RegistrationInstructions) GetRegistTransType() (v enum.RegistTransType, err quickfix.MessageRejectError) {
	var f field.RegistTransTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOwnershipType gets OwnershipType, Tag 517
func (m RegistrationInstructions) GetOwnershipType() (v enum.OwnershipType, err quickfix.MessageRejectError) {
	var f field.OwnershipTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAcctIDSource gets AcctIDSource, Tag 660
func (m RegistrationInstructions) GetAcctIDSource() (v enum.AcctIDSource, err quickfix.MessageRejectError) {
	var f field.AcctIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasAccount returns true if Account is present, Tag 1
func (m RegistrationInstructions) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m RegistrationInstructions) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m RegistrationInstructions) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasNoRegistDtls returns true if NoRegistDtls is present, Tag 473
func (m RegistrationInstructions) HasNoRegistDtls() bool {
	return m.Has(tag.NoRegistDtls)
}

//HasRegistAcctType returns true if RegistAcctType is present, Tag 493
func (m RegistrationInstructions) HasRegistAcctType() bool {
	return m.Has(tag.RegistAcctType)
}

//HasTaxAdvantageType returns true if TaxAdvantageType is present, Tag 495
func (m RegistrationInstructions) HasTaxAdvantageType() bool {
	return m.Has(tag.TaxAdvantageType)
}

//HasRegistRefID returns true if RegistRefID is present, Tag 508
func (m RegistrationInstructions) HasRegistRefID() bool {
	return m.Has(tag.RegistRefID)
}

//HasNoDistribInsts returns true if NoDistribInsts is present, Tag 510
func (m RegistrationInstructions) HasNoDistribInsts() bool {
	return m.Has(tag.NoDistribInsts)
}

//HasRegistID returns true if RegistID is present, Tag 513
func (m RegistrationInstructions) HasRegistID() bool {
	return m.Has(tag.RegistID)
}

//HasRegistTransType returns true if RegistTransType is present, Tag 514
func (m RegistrationInstructions) HasRegistTransType() bool {
	return m.Has(tag.RegistTransType)
}

//HasOwnershipType returns true if OwnershipType is present, Tag 517
func (m RegistrationInstructions) HasOwnershipType() bool {
	return m.Has(tag.OwnershipType)
}

//HasAcctIDSource returns true if AcctIDSource is present, Tag 660
func (m RegistrationInstructions) HasAcctIDSource() bool {
	return m.Has(tag.AcctIDSource)
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

//SetNoPartySubIDs sets NoPartySubIDs, Tag 802
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
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

//GetNoPartySubIDs gets NoPartySubIDs, Tag 802
func (m NoPartyIDs) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
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

//HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802
func (m NoPartyIDs) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

//NoPartySubIDs is a repeating group element, Tag 802
type NoPartySubIDs struct {
	*quickfix.Group
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//SetPartySubIDType sets PartySubIDType, Tag 803
func (m NoPartySubIDs) SetPartySubIDType(v enum.PartySubIDType) {
	m.Set(field.NewPartySubIDType(v))
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartySubIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartySubIDType gets PartySubIDType, Tag 803
func (m NoPartySubIDs) GetPartySubIDType() (v enum.PartySubIDType, err quickfix.MessageRejectError) {
	var f field.PartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasPartySubID returns true if PartySubID is present, Tag 523
func (m NoPartySubIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

//HasPartySubIDType returns true if PartySubIDType is present, Tag 803
func (m NoPartySubIDs) HasPartySubIDType() bool {
	return m.Has(tag.PartySubIDType)
}

//NoPartySubIDsRepeatingGroup is a repeating group, Tag 802
type NoPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartySubIDsRepeatingGroup returns an initialized, NoPartySubIDsRepeatingGroup
func NewNoPartySubIDsRepeatingGroup() NoPartySubIDsRepeatingGroup {
	return NoPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartySubID), quickfix.GroupElement(tag.PartySubIDType)})}
}

//Add create and append a new NoPartySubIDs to this group
func (m NoPartySubIDsRepeatingGroup) Add() NoPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartySubIDs{g}
}

//Get returns the ith NoPartySubIDs in the NoPartySubIDsRepeatinGroup
func (m NoPartySubIDsRepeatingGroup) Get(i int) NoPartySubIDs {
	return NoPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoPartyIDsRepeatingGroup is a repeating group, Tag 453
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), NewNoPartySubIDsRepeatingGroup()})}
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

//NoRegistDtls is a repeating group element, Tag 473
type NoRegistDtls struct {
	*quickfix.Group
}

//SetRegistDtls sets RegistDtls, Tag 509
func (m NoRegistDtls) SetRegistDtls(v string) {
	m.Set(field.NewRegistDtls(v))
}

//SetRegistEmail sets RegistEmail, Tag 511
func (m NoRegistDtls) SetRegistEmail(v string) {
	m.Set(field.NewRegistEmail(v))
}

//SetMailingDtls sets MailingDtls, Tag 474
func (m NoRegistDtls) SetMailingDtls(v string) {
	m.Set(field.NewMailingDtls(v))
}

//SetMailingInst sets MailingInst, Tag 482
func (m NoRegistDtls) SetMailingInst(v string) {
	m.Set(field.NewMailingInst(v))
}

//SetNoNestedPartyIDs sets NoNestedPartyIDs, Tag 539
func (m NoRegistDtls) SetNoNestedPartyIDs(f NoNestedPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetOwnerType sets OwnerType, Tag 522
func (m NoRegistDtls) SetOwnerType(v enum.OwnerType) {
	m.Set(field.NewOwnerType(v))
}

//SetDateOfBirth sets DateOfBirth, Tag 486
func (m NoRegistDtls) SetDateOfBirth(v string) {
	m.Set(field.NewDateOfBirth(v))
}

//SetInvestorCountryOfResidence sets InvestorCountryOfResidence, Tag 475
func (m NoRegistDtls) SetInvestorCountryOfResidence(v string) {
	m.Set(field.NewInvestorCountryOfResidence(v))
}

//GetRegistDtls gets RegistDtls, Tag 509
func (m NoRegistDtls) GetRegistDtls() (v string, err quickfix.MessageRejectError) {
	var f field.RegistDtlsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRegistEmail gets RegistEmail, Tag 511
func (m NoRegistDtls) GetRegistEmail() (v string, err quickfix.MessageRejectError) {
	var f field.RegistEmailField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMailingDtls gets MailingDtls, Tag 474
func (m NoRegistDtls) GetMailingDtls() (v string, err quickfix.MessageRejectError) {
	var f field.MailingDtlsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMailingInst gets MailingInst, Tag 482
func (m NoRegistDtls) GetMailingInst() (v string, err quickfix.MessageRejectError) {
	var f field.MailingInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoNestedPartyIDs gets NoNestedPartyIDs, Tag 539
func (m NoRegistDtls) GetNoNestedPartyIDs() (NoNestedPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetOwnerType gets OwnerType, Tag 522
func (m NoRegistDtls) GetOwnerType() (v enum.OwnerType, err quickfix.MessageRejectError) {
	var f field.OwnerTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDateOfBirth gets DateOfBirth, Tag 486
func (m NoRegistDtls) GetDateOfBirth() (v string, err quickfix.MessageRejectError) {
	var f field.DateOfBirthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInvestorCountryOfResidence gets InvestorCountryOfResidence, Tag 475
func (m NoRegistDtls) GetInvestorCountryOfResidence() (v string, err quickfix.MessageRejectError) {
	var f field.InvestorCountryOfResidenceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasRegistDtls returns true if RegistDtls is present, Tag 509
func (m NoRegistDtls) HasRegistDtls() bool {
	return m.Has(tag.RegistDtls)
}

//HasRegistEmail returns true if RegistEmail is present, Tag 511
func (m NoRegistDtls) HasRegistEmail() bool {
	return m.Has(tag.RegistEmail)
}

//HasMailingDtls returns true if MailingDtls is present, Tag 474
func (m NoRegistDtls) HasMailingDtls() bool {
	return m.Has(tag.MailingDtls)
}

//HasMailingInst returns true if MailingInst is present, Tag 482
func (m NoRegistDtls) HasMailingInst() bool {
	return m.Has(tag.MailingInst)
}

//HasNoNestedPartyIDs returns true if NoNestedPartyIDs is present, Tag 539
func (m NoRegistDtls) HasNoNestedPartyIDs() bool {
	return m.Has(tag.NoNestedPartyIDs)
}

//HasOwnerType returns true if OwnerType is present, Tag 522
func (m NoRegistDtls) HasOwnerType() bool {
	return m.Has(tag.OwnerType)
}

//HasDateOfBirth returns true if DateOfBirth is present, Tag 486
func (m NoRegistDtls) HasDateOfBirth() bool {
	return m.Has(tag.DateOfBirth)
}

//HasInvestorCountryOfResidence returns true if InvestorCountryOfResidence is present, Tag 475
func (m NoRegistDtls) HasInvestorCountryOfResidence() bool {
	return m.Has(tag.InvestorCountryOfResidence)
}

//NoNestedPartyIDs is a repeating group element, Tag 539
type NoNestedPartyIDs struct {
	*quickfix.Group
}

//SetNestedPartyID sets NestedPartyID, Tag 524
func (m NoNestedPartyIDs) SetNestedPartyID(v string) {
	m.Set(field.NewNestedPartyID(v))
}

//SetNestedPartyIDSource sets NestedPartyIDSource, Tag 525
func (m NoNestedPartyIDs) SetNestedPartyIDSource(v string) {
	m.Set(field.NewNestedPartyIDSource(v))
}

//SetNestedPartyRole sets NestedPartyRole, Tag 538
func (m NoNestedPartyIDs) SetNestedPartyRole(v int) {
	m.Set(field.NewNestedPartyRole(v))
}

//SetNoNestedPartySubIDs sets NoNestedPartySubIDs, Tag 804
func (m NoNestedPartyIDs) SetNoNestedPartySubIDs(f NoNestedPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetNestedPartyID gets NestedPartyID, Tag 524
func (m NoNestedPartyIDs) GetNestedPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNestedPartyIDSource gets NestedPartyIDSource, Tag 525
func (m NoNestedPartyIDs) GetNestedPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNestedPartyRole gets NestedPartyRole, Tag 538
func (m NoNestedPartyIDs) GetNestedPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.NestedPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoNestedPartySubIDs gets NoNestedPartySubIDs, Tag 804
func (m NoNestedPartyIDs) GetNoNestedPartySubIDs() (NoNestedPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasNestedPartyID returns true if NestedPartyID is present, Tag 524
func (m NoNestedPartyIDs) HasNestedPartyID() bool {
	return m.Has(tag.NestedPartyID)
}

//HasNestedPartyIDSource returns true if NestedPartyIDSource is present, Tag 525
func (m NoNestedPartyIDs) HasNestedPartyIDSource() bool {
	return m.Has(tag.NestedPartyIDSource)
}

//HasNestedPartyRole returns true if NestedPartyRole is present, Tag 538
func (m NoNestedPartyIDs) HasNestedPartyRole() bool {
	return m.Has(tag.NestedPartyRole)
}

//HasNoNestedPartySubIDs returns true if NoNestedPartySubIDs is present, Tag 804
func (m NoNestedPartyIDs) HasNoNestedPartySubIDs() bool {
	return m.Has(tag.NoNestedPartySubIDs)
}

//NoNestedPartySubIDs is a repeating group element, Tag 804
type NoNestedPartySubIDs struct {
	*quickfix.Group
}

//SetNestedPartySubID sets NestedPartySubID, Tag 545
func (m NoNestedPartySubIDs) SetNestedPartySubID(v string) {
	m.Set(field.NewNestedPartySubID(v))
}

//SetNestedPartySubIDType sets NestedPartySubIDType, Tag 805
func (m NoNestedPartySubIDs) SetNestedPartySubIDType(v int) {
	m.Set(field.NewNestedPartySubIDType(v))
}

//GetNestedPartySubID gets NestedPartySubID, Tag 545
func (m NoNestedPartySubIDs) GetNestedPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNestedPartySubIDType gets NestedPartySubIDType, Tag 805
func (m NoNestedPartySubIDs) GetNestedPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.NestedPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasNestedPartySubID returns true if NestedPartySubID is present, Tag 545
func (m NoNestedPartySubIDs) HasNestedPartySubID() bool {
	return m.Has(tag.NestedPartySubID)
}

//HasNestedPartySubIDType returns true if NestedPartySubIDType is present, Tag 805
func (m NoNestedPartySubIDs) HasNestedPartySubIDType() bool {
	return m.Has(tag.NestedPartySubIDType)
}

//NoNestedPartySubIDsRepeatingGroup is a repeating group, Tag 804
type NoNestedPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNestedPartySubIDsRepeatingGroup returns an initialized, NoNestedPartySubIDsRepeatingGroup
func NewNoNestedPartySubIDsRepeatingGroup() NoNestedPartySubIDsRepeatingGroup {
	return NoNestedPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNestedPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NestedPartySubID), quickfix.GroupElement(tag.NestedPartySubIDType)})}
}

//Add create and append a new NoNestedPartySubIDs to this group
func (m NoNestedPartySubIDsRepeatingGroup) Add() NoNestedPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoNestedPartySubIDs{g}
}

//Get returns the ith NoNestedPartySubIDs in the NoNestedPartySubIDsRepeatinGroup
func (m NoNestedPartySubIDsRepeatingGroup) Get(i int) NoNestedPartySubIDs {
	return NoNestedPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoNestedPartyIDsRepeatingGroup is a repeating group, Tag 539
type NoNestedPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNestedPartyIDsRepeatingGroup returns an initialized, NoNestedPartyIDsRepeatingGroup
func NewNoNestedPartyIDsRepeatingGroup() NoNestedPartyIDsRepeatingGroup {
	return NoNestedPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNestedPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NestedPartyID), quickfix.GroupElement(tag.NestedPartyIDSource), quickfix.GroupElement(tag.NestedPartyRole), NewNoNestedPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoNestedPartyIDs to this group
func (m NoNestedPartyIDsRepeatingGroup) Add() NoNestedPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoNestedPartyIDs{g}
}

//Get returns the ith NoNestedPartyIDs in the NoNestedPartyIDsRepeatinGroup
func (m NoNestedPartyIDsRepeatingGroup) Get(i int) NoNestedPartyIDs {
	return NoNestedPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoRegistDtlsRepeatingGroup is a repeating group, Tag 473
type NoRegistDtlsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRegistDtlsRepeatingGroup returns an initialized, NoRegistDtlsRepeatingGroup
func NewNoRegistDtlsRepeatingGroup() NoRegistDtlsRepeatingGroup {
	return NoRegistDtlsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRegistDtls,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RegistDtls), quickfix.GroupElement(tag.RegistEmail), quickfix.GroupElement(tag.MailingDtls), quickfix.GroupElement(tag.MailingInst), NewNoNestedPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.OwnerType), quickfix.GroupElement(tag.DateOfBirth), quickfix.GroupElement(tag.InvestorCountryOfResidence)})}
}

//Add create and append a new NoRegistDtls to this group
func (m NoRegistDtlsRepeatingGroup) Add() NoRegistDtls {
	g := m.RepeatingGroup.Add()
	return NoRegistDtls{g}
}

//Get returns the ith NoRegistDtls in the NoRegistDtlsRepeatinGroup
func (m NoRegistDtlsRepeatingGroup) Get(i int) NoRegistDtls {
	return NoRegistDtls{m.RepeatingGroup.Get(i)}
}

//NoDistribInsts is a repeating group element, Tag 510
type NoDistribInsts struct {
	*quickfix.Group
}

//SetDistribPaymentMethod sets DistribPaymentMethod, Tag 477
func (m NoDistribInsts) SetDistribPaymentMethod(v enum.DistribPaymentMethod) {
	m.Set(field.NewDistribPaymentMethod(v))
}

//SetDistribPercentage sets DistribPercentage, Tag 512
func (m NoDistribInsts) SetDistribPercentage(value decimal.Decimal, scale int32) {
	m.Set(field.NewDistribPercentage(value, scale))
}

//SetCashDistribCurr sets CashDistribCurr, Tag 478
func (m NoDistribInsts) SetCashDistribCurr(v string) {
	m.Set(field.NewCashDistribCurr(v))
}

//SetCashDistribAgentName sets CashDistribAgentName, Tag 498
func (m NoDistribInsts) SetCashDistribAgentName(v string) {
	m.Set(field.NewCashDistribAgentName(v))
}

//SetCashDistribAgentCode sets CashDistribAgentCode, Tag 499
func (m NoDistribInsts) SetCashDistribAgentCode(v string) {
	m.Set(field.NewCashDistribAgentCode(v))
}

//SetCashDistribAgentAcctNumber sets CashDistribAgentAcctNumber, Tag 500
func (m NoDistribInsts) SetCashDistribAgentAcctNumber(v string) {
	m.Set(field.NewCashDistribAgentAcctNumber(v))
}

//SetCashDistribPayRef sets CashDistribPayRef, Tag 501
func (m NoDistribInsts) SetCashDistribPayRef(v string) {
	m.Set(field.NewCashDistribPayRef(v))
}

//SetCashDistribAgentAcctName sets CashDistribAgentAcctName, Tag 502
func (m NoDistribInsts) SetCashDistribAgentAcctName(v string) {
	m.Set(field.NewCashDistribAgentAcctName(v))
}

//GetDistribPaymentMethod gets DistribPaymentMethod, Tag 477
func (m NoDistribInsts) GetDistribPaymentMethod() (v enum.DistribPaymentMethod, err quickfix.MessageRejectError) {
	var f field.DistribPaymentMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDistribPercentage gets DistribPercentage, Tag 512
func (m NoDistribInsts) GetDistribPercentage() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DistribPercentageField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashDistribCurr gets CashDistribCurr, Tag 478
func (m NoDistribInsts) GetCashDistribCurr() (v string, err quickfix.MessageRejectError) {
	var f field.CashDistribCurrField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashDistribAgentName gets CashDistribAgentName, Tag 498
func (m NoDistribInsts) GetCashDistribAgentName() (v string, err quickfix.MessageRejectError) {
	var f field.CashDistribAgentNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashDistribAgentCode gets CashDistribAgentCode, Tag 499
func (m NoDistribInsts) GetCashDistribAgentCode() (v string, err quickfix.MessageRejectError) {
	var f field.CashDistribAgentCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashDistribAgentAcctNumber gets CashDistribAgentAcctNumber, Tag 500
func (m NoDistribInsts) GetCashDistribAgentAcctNumber() (v string, err quickfix.MessageRejectError) {
	var f field.CashDistribAgentAcctNumberField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashDistribPayRef gets CashDistribPayRef, Tag 501
func (m NoDistribInsts) GetCashDistribPayRef() (v string, err quickfix.MessageRejectError) {
	var f field.CashDistribPayRefField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashDistribAgentAcctName gets CashDistribAgentAcctName, Tag 502
func (m NoDistribInsts) GetCashDistribAgentAcctName() (v string, err quickfix.MessageRejectError) {
	var f field.CashDistribAgentAcctNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasDistribPaymentMethod returns true if DistribPaymentMethod is present, Tag 477
func (m NoDistribInsts) HasDistribPaymentMethod() bool {
	return m.Has(tag.DistribPaymentMethod)
}

//HasDistribPercentage returns true if DistribPercentage is present, Tag 512
func (m NoDistribInsts) HasDistribPercentage() bool {
	return m.Has(tag.DistribPercentage)
}

//HasCashDistribCurr returns true if CashDistribCurr is present, Tag 478
func (m NoDistribInsts) HasCashDistribCurr() bool {
	return m.Has(tag.CashDistribCurr)
}

//HasCashDistribAgentName returns true if CashDistribAgentName is present, Tag 498
func (m NoDistribInsts) HasCashDistribAgentName() bool {
	return m.Has(tag.CashDistribAgentName)
}

//HasCashDistribAgentCode returns true if CashDistribAgentCode is present, Tag 499
func (m NoDistribInsts) HasCashDistribAgentCode() bool {
	return m.Has(tag.CashDistribAgentCode)
}

//HasCashDistribAgentAcctNumber returns true if CashDistribAgentAcctNumber is present, Tag 500
func (m NoDistribInsts) HasCashDistribAgentAcctNumber() bool {
	return m.Has(tag.CashDistribAgentAcctNumber)
}

//HasCashDistribPayRef returns true if CashDistribPayRef is present, Tag 501
func (m NoDistribInsts) HasCashDistribPayRef() bool {
	return m.Has(tag.CashDistribPayRef)
}

//HasCashDistribAgentAcctName returns true if CashDistribAgentAcctName is present, Tag 502
func (m NoDistribInsts) HasCashDistribAgentAcctName() bool {
	return m.Has(tag.CashDistribAgentAcctName)
}

//NoDistribInstsRepeatingGroup is a repeating group, Tag 510
type NoDistribInstsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoDistribInstsRepeatingGroup returns an initialized, NoDistribInstsRepeatingGroup
func NewNoDistribInstsRepeatingGroup() NoDistribInstsRepeatingGroup {
	return NoDistribInstsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoDistribInsts,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.DistribPaymentMethod), quickfix.GroupElement(tag.DistribPercentage), quickfix.GroupElement(tag.CashDistribCurr), quickfix.GroupElement(tag.CashDistribAgentName), quickfix.GroupElement(tag.CashDistribAgentCode), quickfix.GroupElement(tag.CashDistribAgentAcctNumber), quickfix.GroupElement(tag.CashDistribPayRef), quickfix.GroupElement(tag.CashDistribAgentAcctName)})}
}

//Add create and append a new NoDistribInsts to this group
func (m NoDistribInstsRepeatingGroup) Add() NoDistribInsts {
	g := m.RepeatingGroup.Add()
	return NoDistribInsts{g}
}

//Get returns the ith NoDistribInsts in the NoDistribInstsRepeatinGroup
func (m NoDistribInstsRepeatingGroup) Get(i int) NoDistribInsts {
	return NoDistribInsts{m.RepeatingGroup.Get(i)}
}
