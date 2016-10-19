package marketdatarequestreject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//MarketDataRequestReject is the fix50sp2 MarketDataRequestReject type, MsgType = Y
type MarketDataRequestReject struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

//FromMessage creates a MarketDataRequestReject from a quickfix.Message instance
func FromMessage(m *quickfix.Message) MarketDataRequestReject {
	return MarketDataRequestReject{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m MarketDataRequestReject) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a MarketDataRequestReject initialized with the required fields for MarketDataRequestReject
func New(mdreqid field.MDReqIDField) (m MarketDataRequestReject) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("Y"))
	m.Set(mdreqid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg MarketDataRequestReject, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "Y", r
}

//SetText sets Text, Tag 58
func (m MarketDataRequestReject) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetMDReqID sets MDReqID, Tag 262
func (m MarketDataRequestReject) SetMDReqID(v string) {
	m.Set(field.NewMDReqID(v))
}

//SetMDReqRejReason sets MDReqRejReason, Tag 281
func (m MarketDataRequestReject) SetMDReqRejReason(v enum.MDReqRejReason) {
	m.Set(field.NewMDReqRejReason(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m MarketDataRequestReject) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m MarketDataRequestReject) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m MarketDataRequestReject) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoAltMDSource sets NoAltMDSource, Tag 816
func (m MarketDataRequestReject) SetNoAltMDSource(f NoAltMDSourceRepeatingGroup) {
	m.SetGroup(f)
}

//GetText gets Text, Tag 58
func (m MarketDataRequestReject) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDReqID gets MDReqID, Tag 262
func (m MarketDataRequestReject) GetMDReqID() (v string, err quickfix.MessageRejectError) {
	var f field.MDReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDReqRejReason gets MDReqRejReason, Tag 281
func (m MarketDataRequestReject) GetMDReqRejReason() (v enum.MDReqRejReason, err quickfix.MessageRejectError) {
	var f field.MDReqRejReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m MarketDataRequestReject) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m MarketDataRequestReject) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m MarketDataRequestReject) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoAltMDSource gets NoAltMDSource, Tag 816
func (m MarketDataRequestReject) GetNoAltMDSource() (NoAltMDSourceRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAltMDSourceRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasText returns true if Text is present, Tag 58
func (m MarketDataRequestReject) HasText() bool {
	return m.Has(tag.Text)
}

//HasMDReqID returns true if MDReqID is present, Tag 262
func (m MarketDataRequestReject) HasMDReqID() bool {
	return m.Has(tag.MDReqID)
}

//HasMDReqRejReason returns true if MDReqRejReason is present, Tag 281
func (m MarketDataRequestReject) HasMDReqRejReason() bool {
	return m.Has(tag.MDReqRejReason)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m MarketDataRequestReject) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m MarketDataRequestReject) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m MarketDataRequestReject) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasNoAltMDSource returns true if NoAltMDSource is present, Tag 816
func (m MarketDataRequestReject) HasNoAltMDSource() bool {
	return m.Has(tag.NoAltMDSource)
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

//NoAltMDSource is a repeating group element, Tag 816
type NoAltMDSource struct {
	*quickfix.Group
}

//SetAltMDSourceID sets AltMDSourceID, Tag 817
func (m NoAltMDSource) SetAltMDSourceID(v string) {
	m.Set(field.NewAltMDSourceID(v))
}

//GetAltMDSourceID gets AltMDSourceID, Tag 817
func (m NoAltMDSource) GetAltMDSourceID() (v string, err quickfix.MessageRejectError) {
	var f field.AltMDSourceIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasAltMDSourceID returns true if AltMDSourceID is present, Tag 817
func (m NoAltMDSource) HasAltMDSourceID() bool {
	return m.Has(tag.AltMDSourceID)
}

//NoAltMDSourceRepeatingGroup is a repeating group, Tag 816
type NoAltMDSourceRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoAltMDSourceRepeatingGroup returns an initialized, NoAltMDSourceRepeatingGroup
func NewNoAltMDSourceRepeatingGroup() NoAltMDSourceRepeatingGroup {
	return NoAltMDSourceRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAltMDSource,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AltMDSourceID)})}
}

//Add create and append a new NoAltMDSource to this group
func (m NoAltMDSourceRepeatingGroup) Add() NoAltMDSource {
	g := m.RepeatingGroup.Add()
	return NoAltMDSource{g}
}

//Get returns the ith NoAltMDSource in the NoAltMDSourceRepeatinGroup
func (m NoAltMDSourceRepeatingGroup) Get(i int) NoAltMDSource {
	return NoAltMDSource{m.RepeatingGroup.Get(i)}
}
