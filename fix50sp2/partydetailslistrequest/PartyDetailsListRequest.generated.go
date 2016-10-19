package partydetailslistrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//PartyDetailsListRequest is the fix50sp2 PartyDetailsListRequest type, MsgType = CF
type PartyDetailsListRequest struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

//FromMessage creates a PartyDetailsListRequest from a quickfix.Message instance
func FromMessage(m *quickfix.Message) PartyDetailsListRequest {
	return PartyDetailsListRequest{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m PartyDetailsListRequest) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a PartyDetailsListRequest initialized with the required fields for PartyDetailsListRequest
func New(partydetailslistrequestid field.PartyDetailsListRequestIDField) (m PartyDetailsListRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("CF"))
	m.Set(partydetailslistrequestid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg PartyDetailsListRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "CF", r
}

//SetText sets Text, Tag 58
func (m PartyDetailsListRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263
func (m PartyDetailsListRequest) SetSubscriptionRequestType(v enum.SubscriptionRequestType) {
	m.Set(field.NewSubscriptionRequestType(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m PartyDetailsListRequest) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m PartyDetailsListRequest) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m PartyDetailsListRequest) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetPartyDetailsListRequestID sets PartyDetailsListRequestID, Tag 1505
func (m PartyDetailsListRequest) SetPartyDetailsListRequestID(v string) {
	m.Set(field.NewPartyDetailsListRequestID(v))
}

//SetNoPartyListResponseTypes sets NoPartyListResponseTypes, Tag 1506
func (m PartyDetailsListRequest) SetNoPartyListResponseTypes(f NoPartyListResponseTypesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoRequestedPartyRoles sets NoRequestedPartyRoles, Tag 1508
func (m PartyDetailsListRequest) SetNoRequestedPartyRoles(f NoRequestedPartyRolesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoPartyRelationships sets NoPartyRelationships, Tag 1514
func (m PartyDetailsListRequest) SetNoPartyRelationships(f NoPartyRelationshipsRepeatingGroup) {
	m.SetGroup(f)
}

//GetText gets Text, Tag 58
func (m PartyDetailsListRequest) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263
func (m PartyDetailsListRequest) GetSubscriptionRequestType() (v enum.SubscriptionRequestType, err quickfix.MessageRejectError) {
	var f field.SubscriptionRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m PartyDetailsListRequest) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m PartyDetailsListRequest) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m PartyDetailsListRequest) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetPartyDetailsListRequestID gets PartyDetailsListRequestID, Tag 1505
func (m PartyDetailsListRequest) GetPartyDetailsListRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyDetailsListRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartyListResponseTypes gets NoPartyListResponseTypes, Tag 1506
func (m PartyDetailsListRequest) GetNoPartyListResponseTypes() (NoPartyListResponseTypesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyListResponseTypesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoRequestedPartyRoles gets NoRequestedPartyRoles, Tag 1508
func (m PartyDetailsListRequest) GetNoRequestedPartyRoles() (NoRequestedPartyRolesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRequestedPartyRolesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoPartyRelationships gets NoPartyRelationships, Tag 1514
func (m PartyDetailsListRequest) GetNoPartyRelationships() (NoPartyRelationshipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyRelationshipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasText returns true if Text is present, Tag 58
func (m PartyDetailsListRequest) HasText() bool {
	return m.Has(tag.Text)
}

//HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263
func (m PartyDetailsListRequest) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m PartyDetailsListRequest) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m PartyDetailsListRequest) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m PartyDetailsListRequest) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasPartyDetailsListRequestID returns true if PartyDetailsListRequestID is present, Tag 1505
func (m PartyDetailsListRequest) HasPartyDetailsListRequestID() bool {
	return m.Has(tag.PartyDetailsListRequestID)
}

//HasNoPartyListResponseTypes returns true if NoPartyListResponseTypes is present, Tag 1506
func (m PartyDetailsListRequest) HasNoPartyListResponseTypes() bool {
	return m.Has(tag.NoPartyListResponseTypes)
}

//HasNoRequestedPartyRoles returns true if NoRequestedPartyRoles is present, Tag 1508
func (m PartyDetailsListRequest) HasNoRequestedPartyRoles() bool {
	return m.Has(tag.NoRequestedPartyRoles)
}

//HasNoPartyRelationships returns true if NoPartyRelationships is present, Tag 1514
func (m PartyDetailsListRequest) HasNoPartyRelationships() bool {
	return m.Has(tag.NoPartyRelationships)
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

//NoPartyListResponseTypes is a repeating group element, Tag 1506
type NoPartyListResponseTypes struct {
	*quickfix.Group
}

//SetPartyListResponseType sets PartyListResponseType, Tag 1507
func (m NoPartyListResponseTypes) SetPartyListResponseType(v enum.PartyListResponseType) {
	m.Set(field.NewPartyListResponseType(v))
}

//GetPartyListResponseType gets PartyListResponseType, Tag 1507
func (m NoPartyListResponseTypes) GetPartyListResponseType() (v enum.PartyListResponseType, err quickfix.MessageRejectError) {
	var f field.PartyListResponseTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasPartyListResponseType returns true if PartyListResponseType is present, Tag 1507
func (m NoPartyListResponseTypes) HasPartyListResponseType() bool {
	return m.Has(tag.PartyListResponseType)
}

//NoPartyListResponseTypesRepeatingGroup is a repeating group, Tag 1506
type NoPartyListResponseTypesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyListResponseTypesRepeatingGroup returns an initialized, NoPartyListResponseTypesRepeatingGroup
func NewNoPartyListResponseTypesRepeatingGroup() NoPartyListResponseTypesRepeatingGroup {
	return NoPartyListResponseTypesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyListResponseTypes,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyListResponseType)})}
}

//Add create and append a new NoPartyListResponseTypes to this group
func (m NoPartyListResponseTypesRepeatingGroup) Add() NoPartyListResponseTypes {
	g := m.RepeatingGroup.Add()
	return NoPartyListResponseTypes{g}
}

//Get returns the ith NoPartyListResponseTypes in the NoPartyListResponseTypesRepeatinGroup
func (m NoPartyListResponseTypesRepeatingGroup) Get(i int) NoPartyListResponseTypes {
	return NoPartyListResponseTypes{m.RepeatingGroup.Get(i)}
}

//NoRequestedPartyRoles is a repeating group element, Tag 1508
type NoRequestedPartyRoles struct {
	*quickfix.Group
}

//SetRequestedPartyRole sets RequestedPartyRole, Tag 1509
func (m NoRequestedPartyRoles) SetRequestedPartyRole(v int) {
	m.Set(field.NewRequestedPartyRole(v))
}

//GetRequestedPartyRole gets RequestedPartyRole, Tag 1509
func (m NoRequestedPartyRoles) GetRequestedPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.RequestedPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasRequestedPartyRole returns true if RequestedPartyRole is present, Tag 1509
func (m NoRequestedPartyRoles) HasRequestedPartyRole() bool {
	return m.Has(tag.RequestedPartyRole)
}

//NoRequestedPartyRolesRepeatingGroup is a repeating group, Tag 1508
type NoRequestedPartyRolesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRequestedPartyRolesRepeatingGroup returns an initialized, NoRequestedPartyRolesRepeatingGroup
func NewNoRequestedPartyRolesRepeatingGroup() NoRequestedPartyRolesRepeatingGroup {
	return NoRequestedPartyRolesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRequestedPartyRoles,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RequestedPartyRole)})}
}

//Add create and append a new NoRequestedPartyRoles to this group
func (m NoRequestedPartyRolesRepeatingGroup) Add() NoRequestedPartyRoles {
	g := m.RepeatingGroup.Add()
	return NoRequestedPartyRoles{g}
}

//Get returns the ith NoRequestedPartyRoles in the NoRequestedPartyRolesRepeatinGroup
func (m NoRequestedPartyRolesRepeatingGroup) Get(i int) NoRequestedPartyRoles {
	return NoRequestedPartyRoles{m.RepeatingGroup.Get(i)}
}

//NoPartyRelationships is a repeating group element, Tag 1514
type NoPartyRelationships struct {
	*quickfix.Group
}

//SetPartyRelationship sets PartyRelationship, Tag 1515
func (m NoPartyRelationships) SetPartyRelationship(v enum.PartyRelationship) {
	m.Set(field.NewPartyRelationship(v))
}

//GetPartyRelationship gets PartyRelationship, Tag 1515
func (m NoPartyRelationships) GetPartyRelationship() (v enum.PartyRelationship, err quickfix.MessageRejectError) {
	var f field.PartyRelationshipField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasPartyRelationship returns true if PartyRelationship is present, Tag 1515
func (m NoPartyRelationships) HasPartyRelationship() bool {
	return m.Has(tag.PartyRelationship)
}

//NoPartyRelationshipsRepeatingGroup is a repeating group, Tag 1514
type NoPartyRelationshipsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyRelationshipsRepeatingGroup returns an initialized, NoPartyRelationshipsRepeatingGroup
func NewNoPartyRelationshipsRepeatingGroup() NoPartyRelationshipsRepeatingGroup {
	return NoPartyRelationshipsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyRelationships,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyRelationship)})}
}

//Add create and append a new NoPartyRelationships to this group
func (m NoPartyRelationshipsRepeatingGroup) Add() NoPartyRelationships {
	g := m.RepeatingGroup.Add()
	return NoPartyRelationships{g}
}

//Get returns the ith NoPartyRelationships in the NoPartyRelationshipsRepeatinGroup
func (m NoPartyRelationshipsRepeatingGroup) Get(i int) NoPartyRelationships {
	return NoPartyRelationships{m.RepeatingGroup.Get(i)}
}
