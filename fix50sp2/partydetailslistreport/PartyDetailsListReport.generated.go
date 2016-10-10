package partydetailslistreport

import (
	"github.com/shopspring/decimal"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//PartyDetailsListReport is the fix50sp2 PartyDetailsListReport type, MsgType = CG
type PartyDetailsListReport struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

//FromMessage creates a PartyDetailsListReport from a quickfix.Message instance
func FromMessage(m *quickfix.Message) PartyDetailsListReport {
	return PartyDetailsListReport{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m PartyDetailsListReport) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a PartyDetailsListReport initialized with the required fields for PartyDetailsListReport
func New(partydetailslistreportid field.PartyDetailsListReportIDField) (m PartyDetailsListReport) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("CG"))
	m.Set(partydetailslistreportid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg PartyDetailsListReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "CG", r
}

//SetText sets Text, Tag 58
func (m PartyDetailsListReport) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m PartyDetailsListReport) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m PartyDetailsListReport) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetLastFragment sets LastFragment, Tag 893
func (m PartyDetailsListReport) SetLastFragment(v bool) {
	m.Set(field.NewLastFragment(v))
}

//SetApplID sets ApplID, Tag 1180
func (m PartyDetailsListReport) SetApplID(v string) {
	m.Set(field.NewApplID(v))
}

//SetApplSeqNum sets ApplSeqNum, Tag 1181
func (m PartyDetailsListReport) SetApplSeqNum(v int) {
	m.Set(field.NewApplSeqNum(v))
}

//SetApplLastSeqNum sets ApplLastSeqNum, Tag 1350
func (m PartyDetailsListReport) SetApplLastSeqNum(v int) {
	m.Set(field.NewApplLastSeqNum(v))
}

//SetApplResendFlag sets ApplResendFlag, Tag 1352
func (m PartyDetailsListReport) SetApplResendFlag(v bool) {
	m.Set(field.NewApplResendFlag(v))
}

//SetPartyDetailsListRequestID sets PartyDetailsListRequestID, Tag 1505
func (m PartyDetailsListReport) SetPartyDetailsListRequestID(v string) {
	m.Set(field.NewPartyDetailsListRequestID(v))
}

//SetPartyDetailsListReportID sets PartyDetailsListReportID, Tag 1510
func (m PartyDetailsListReport) SetPartyDetailsListReportID(v string) {
	m.Set(field.NewPartyDetailsListReportID(v))
}

//SetPartyDetailsRequestResult sets PartyDetailsRequestResult, Tag 1511
func (m PartyDetailsListReport) SetPartyDetailsRequestResult(v enum.PartyDetailsRequestResult) {
	m.Set(field.NewPartyDetailsRequestResult(v))
}

//SetTotNoPartyList sets TotNoPartyList, Tag 1512
func (m PartyDetailsListReport) SetTotNoPartyList(v int) {
	m.Set(field.NewTotNoPartyList(v))
}

//SetNoPartyList sets NoPartyList, Tag 1513
func (m PartyDetailsListReport) SetNoPartyList(f NoPartyListRepeatingGroup) {
	m.SetGroup(f)
}

//GetText gets Text, Tag 58
func (m PartyDetailsListReport) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m PartyDetailsListReport) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m PartyDetailsListReport) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastFragment gets LastFragment, Tag 893
func (m PartyDetailsListReport) GetLastFragment() (v bool, err quickfix.MessageRejectError) {
	var f field.LastFragmentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplID gets ApplID, Tag 1180
func (m PartyDetailsListReport) GetApplID() (v string, err quickfix.MessageRejectError) {
	var f field.ApplIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplSeqNum gets ApplSeqNum, Tag 1181
func (m PartyDetailsListReport) GetApplSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplLastSeqNum gets ApplLastSeqNum, Tag 1350
func (m PartyDetailsListReport) GetApplLastSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplLastSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplResendFlag gets ApplResendFlag, Tag 1352
func (m PartyDetailsListReport) GetApplResendFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.ApplResendFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyDetailsListRequestID gets PartyDetailsListRequestID, Tag 1505
func (m PartyDetailsListReport) GetPartyDetailsListRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyDetailsListRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyDetailsListReportID gets PartyDetailsListReportID, Tag 1510
func (m PartyDetailsListReport) GetPartyDetailsListReportID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyDetailsListReportIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyDetailsRequestResult gets PartyDetailsRequestResult, Tag 1511
func (m PartyDetailsListReport) GetPartyDetailsRequestResult() (v enum.PartyDetailsRequestResult, err quickfix.MessageRejectError) {
	var f field.PartyDetailsRequestResultField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTotNoPartyList gets TotNoPartyList, Tag 1512
func (m PartyDetailsListReport) GetTotNoPartyList() (v int, err quickfix.MessageRejectError) {
	var f field.TotNoPartyListField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartyList gets NoPartyList, Tag 1513
func (m PartyDetailsListReport) GetNoPartyList() (NoPartyListRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyListRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasText returns true if Text is present, Tag 58
func (m PartyDetailsListReport) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m PartyDetailsListReport) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m PartyDetailsListReport) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasLastFragment returns true if LastFragment is present, Tag 893
func (m PartyDetailsListReport) HasLastFragment() bool {
	return m.Has(tag.LastFragment)
}

//HasApplID returns true if ApplID is present, Tag 1180
func (m PartyDetailsListReport) HasApplID() bool {
	return m.Has(tag.ApplID)
}

//HasApplSeqNum returns true if ApplSeqNum is present, Tag 1181
func (m PartyDetailsListReport) HasApplSeqNum() bool {
	return m.Has(tag.ApplSeqNum)
}

//HasApplLastSeqNum returns true if ApplLastSeqNum is present, Tag 1350
func (m PartyDetailsListReport) HasApplLastSeqNum() bool {
	return m.Has(tag.ApplLastSeqNum)
}

//HasApplResendFlag returns true if ApplResendFlag is present, Tag 1352
func (m PartyDetailsListReport) HasApplResendFlag() bool {
	return m.Has(tag.ApplResendFlag)
}

//HasPartyDetailsListRequestID returns true if PartyDetailsListRequestID is present, Tag 1505
func (m PartyDetailsListReport) HasPartyDetailsListRequestID() bool {
	return m.Has(tag.PartyDetailsListRequestID)
}

//HasPartyDetailsListReportID returns true if PartyDetailsListReportID is present, Tag 1510
func (m PartyDetailsListReport) HasPartyDetailsListReportID() bool {
	return m.Has(tag.PartyDetailsListReportID)
}

//HasPartyDetailsRequestResult returns true if PartyDetailsRequestResult is present, Tag 1511
func (m PartyDetailsListReport) HasPartyDetailsRequestResult() bool {
	return m.Has(tag.PartyDetailsRequestResult)
}

//HasTotNoPartyList returns true if TotNoPartyList is present, Tag 1512
func (m PartyDetailsListReport) HasTotNoPartyList() bool {
	return m.Has(tag.TotNoPartyList)
}

//HasNoPartyList returns true if NoPartyList is present, Tag 1513
func (m PartyDetailsListReport) HasNoPartyList() bool {
	return m.Has(tag.NoPartyList)
}

//NoPartyList is a repeating group element, Tag 1513
type NoPartyList struct {
	*quickfix.Group
}

//SetPartyID sets PartyID, Tag 448
func (m NoPartyList) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

//SetPartyIDSource sets PartyIDSource, Tag 447
func (m NoPartyList) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

//SetPartyRole sets PartyRole, Tag 452
func (m NoPartyList) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

//SetNoPartySubIDs sets NoPartySubIDs, Tag 802
func (m NoPartyList) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoPartyAltIDs sets NoPartyAltIDs, Tag 1516
func (m NoPartyList) SetNoPartyAltIDs(f NoPartyAltIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoContextPartyIDs sets NoContextPartyIDs, Tag 1522
func (m NoPartyList) SetNoContextPartyIDs(f NoContextPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoRiskLimits sets NoRiskLimits, Tag 1529
func (m NoPartyList) SetNoRiskLimits(f NoRiskLimitsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoRelatedPartyIDs sets NoRelatedPartyIDs, Tag 1562
func (m NoPartyList) SetNoRelatedPartyIDs(f NoRelatedPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetPartyID gets PartyID, Tag 448
func (m NoPartyList) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyIDSource gets PartyIDSource, Tag 447
func (m NoPartyList) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyRole gets PartyRole, Tag 452
func (m NoPartyList) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartySubIDs gets NoPartySubIDs, Tag 802
func (m NoPartyList) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoPartyAltIDs gets NoPartyAltIDs, Tag 1516
func (m NoPartyList) GetNoPartyAltIDs() (NoPartyAltIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyAltIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoContextPartyIDs gets NoContextPartyIDs, Tag 1522
func (m NoPartyList) GetNoContextPartyIDs() (NoContextPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoContextPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoRiskLimits gets NoRiskLimits, Tag 1529
func (m NoPartyList) GetNoRiskLimits() (NoRiskLimitsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRiskLimitsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoRelatedPartyIDs gets NoRelatedPartyIDs, Tag 1562
func (m NoPartyList) GetNoRelatedPartyIDs() (NoRelatedPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasPartyID returns true if PartyID is present, Tag 448
func (m NoPartyList) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

//HasPartyIDSource returns true if PartyIDSource is present, Tag 447
func (m NoPartyList) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

//HasPartyRole returns true if PartyRole is present, Tag 452
func (m NoPartyList) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

//HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802
func (m NoPartyList) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

//HasNoPartyAltIDs returns true if NoPartyAltIDs is present, Tag 1516
func (m NoPartyList) HasNoPartyAltIDs() bool {
	return m.Has(tag.NoPartyAltIDs)
}

//HasNoContextPartyIDs returns true if NoContextPartyIDs is present, Tag 1522
func (m NoPartyList) HasNoContextPartyIDs() bool {
	return m.Has(tag.NoContextPartyIDs)
}

//HasNoRiskLimits returns true if NoRiskLimits is present, Tag 1529
func (m NoPartyList) HasNoRiskLimits() bool {
	return m.Has(tag.NoRiskLimits)
}

//HasNoRelatedPartyIDs returns true if NoRelatedPartyIDs is present, Tag 1562
func (m NoPartyList) HasNoRelatedPartyIDs() bool {
	return m.Has(tag.NoRelatedPartyIDs)
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

//NoPartyAltIDs is a repeating group element, Tag 1516
type NoPartyAltIDs struct {
	*quickfix.Group
}

//SetPartyAltID sets PartyAltID, Tag 1517
func (m NoPartyAltIDs) SetPartyAltID(v string) {
	m.Set(field.NewPartyAltID(v))
}

//SetPartyAltIDSource sets PartyAltIDSource, Tag 1518
func (m NoPartyAltIDs) SetPartyAltIDSource(v string) {
	m.Set(field.NewPartyAltIDSource(v))
}

//SetNoPartyAltSubIDs sets NoPartyAltSubIDs, Tag 1519
func (m NoPartyAltIDs) SetNoPartyAltSubIDs(f NoPartyAltSubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetPartyAltID gets PartyAltID, Tag 1517
func (m NoPartyAltIDs) GetPartyAltID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyAltIDSource gets PartyAltIDSource, Tag 1518
func (m NoPartyAltIDs) GetPartyAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.PartyAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartyAltSubIDs gets NoPartyAltSubIDs, Tag 1519
func (m NoPartyAltIDs) GetNoPartyAltSubIDs() (NoPartyAltSubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyAltSubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasPartyAltID returns true if PartyAltID is present, Tag 1517
func (m NoPartyAltIDs) HasPartyAltID() bool {
	return m.Has(tag.PartyAltID)
}

//HasPartyAltIDSource returns true if PartyAltIDSource is present, Tag 1518
func (m NoPartyAltIDs) HasPartyAltIDSource() bool {
	return m.Has(tag.PartyAltIDSource)
}

//HasNoPartyAltSubIDs returns true if NoPartyAltSubIDs is present, Tag 1519
func (m NoPartyAltIDs) HasNoPartyAltSubIDs() bool {
	return m.Has(tag.NoPartyAltSubIDs)
}

//NoPartyAltSubIDs is a repeating group element, Tag 1519
type NoPartyAltSubIDs struct {
	*quickfix.Group
}

//SetPartyAltSubID sets PartyAltSubID, Tag 1520
func (m NoPartyAltSubIDs) SetPartyAltSubID(v string) {
	m.Set(field.NewPartyAltSubID(v))
}

//SetPartyAltSubIDType sets PartyAltSubIDType, Tag 1521
func (m NoPartyAltSubIDs) SetPartyAltSubIDType(v int) {
	m.Set(field.NewPartyAltSubIDType(v))
}

//GetPartyAltSubID gets PartyAltSubID, Tag 1520
func (m NoPartyAltSubIDs) GetPartyAltSubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyAltSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyAltSubIDType gets PartyAltSubIDType, Tag 1521
func (m NoPartyAltSubIDs) GetPartyAltSubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.PartyAltSubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasPartyAltSubID returns true if PartyAltSubID is present, Tag 1520
func (m NoPartyAltSubIDs) HasPartyAltSubID() bool {
	return m.Has(tag.PartyAltSubID)
}

//HasPartyAltSubIDType returns true if PartyAltSubIDType is present, Tag 1521
func (m NoPartyAltSubIDs) HasPartyAltSubIDType() bool {
	return m.Has(tag.PartyAltSubIDType)
}

//NoPartyAltSubIDsRepeatingGroup is a repeating group, Tag 1519
type NoPartyAltSubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyAltSubIDsRepeatingGroup returns an initialized, NoPartyAltSubIDsRepeatingGroup
func NewNoPartyAltSubIDsRepeatingGroup() NoPartyAltSubIDsRepeatingGroup {
	return NoPartyAltSubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyAltSubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyAltSubID), quickfix.GroupElement(tag.PartyAltSubIDType)})}
}

//Add create and append a new NoPartyAltSubIDs to this group
func (m NoPartyAltSubIDsRepeatingGroup) Add() NoPartyAltSubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyAltSubIDs{g}
}

//Get returns the ith NoPartyAltSubIDs in the NoPartyAltSubIDsRepeatinGroup
func (m NoPartyAltSubIDsRepeatingGroup) Get(i int) NoPartyAltSubIDs {
	return NoPartyAltSubIDs{m.RepeatingGroup.Get(i)}
}

//NoPartyAltIDsRepeatingGroup is a repeating group, Tag 1516
type NoPartyAltIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyAltIDsRepeatingGroup returns an initialized, NoPartyAltIDsRepeatingGroup
func NewNoPartyAltIDsRepeatingGroup() NoPartyAltIDsRepeatingGroup {
	return NoPartyAltIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyAltIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyAltID), quickfix.GroupElement(tag.PartyAltIDSource), NewNoPartyAltSubIDsRepeatingGroup()})}
}

//Add create and append a new NoPartyAltIDs to this group
func (m NoPartyAltIDsRepeatingGroup) Add() NoPartyAltIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyAltIDs{g}
}

//Get returns the ith NoPartyAltIDs in the NoPartyAltIDsRepeatinGroup
func (m NoPartyAltIDsRepeatingGroup) Get(i int) NoPartyAltIDs {
	return NoPartyAltIDs{m.RepeatingGroup.Get(i)}
}

//NoContextPartyIDs is a repeating group element, Tag 1522
type NoContextPartyIDs struct {
	*quickfix.Group
}

//SetContextPartyID sets ContextPartyID, Tag 1523
func (m NoContextPartyIDs) SetContextPartyID(v string) {
	m.Set(field.NewContextPartyID(v))
}

//SetContextPartyIDSource sets ContextPartyIDSource, Tag 1524
func (m NoContextPartyIDs) SetContextPartyIDSource(v string) {
	m.Set(field.NewContextPartyIDSource(v))
}

//SetContextPartyRole sets ContextPartyRole, Tag 1525
func (m NoContextPartyIDs) SetContextPartyRole(v int) {
	m.Set(field.NewContextPartyRole(v))
}

//SetNoContextPartySubIDs sets NoContextPartySubIDs, Tag 1526
func (m NoContextPartyIDs) SetNoContextPartySubIDs(f NoContextPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetContextPartyID gets ContextPartyID, Tag 1523
func (m NoContextPartyIDs) GetContextPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.ContextPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContextPartyIDSource gets ContextPartyIDSource, Tag 1524
func (m NoContextPartyIDs) GetContextPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.ContextPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContextPartyRole gets ContextPartyRole, Tag 1525
func (m NoContextPartyIDs) GetContextPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.ContextPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoContextPartySubIDs gets NoContextPartySubIDs, Tag 1526
func (m NoContextPartyIDs) GetNoContextPartySubIDs() (NoContextPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoContextPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasContextPartyID returns true if ContextPartyID is present, Tag 1523
func (m NoContextPartyIDs) HasContextPartyID() bool {
	return m.Has(tag.ContextPartyID)
}

//HasContextPartyIDSource returns true if ContextPartyIDSource is present, Tag 1524
func (m NoContextPartyIDs) HasContextPartyIDSource() bool {
	return m.Has(tag.ContextPartyIDSource)
}

//HasContextPartyRole returns true if ContextPartyRole is present, Tag 1525
func (m NoContextPartyIDs) HasContextPartyRole() bool {
	return m.Has(tag.ContextPartyRole)
}

//HasNoContextPartySubIDs returns true if NoContextPartySubIDs is present, Tag 1526
func (m NoContextPartyIDs) HasNoContextPartySubIDs() bool {
	return m.Has(tag.NoContextPartySubIDs)
}

//NoContextPartySubIDs is a repeating group element, Tag 1526
type NoContextPartySubIDs struct {
	*quickfix.Group
}

//SetContextPartySubID sets ContextPartySubID, Tag 1527
func (m NoContextPartySubIDs) SetContextPartySubID(v string) {
	m.Set(field.NewContextPartySubID(v))
}

//SetContextPartySubIDType sets ContextPartySubIDType, Tag 1528
func (m NoContextPartySubIDs) SetContextPartySubIDType(v int) {
	m.Set(field.NewContextPartySubIDType(v))
}

//GetContextPartySubID gets ContextPartySubID, Tag 1527
func (m NoContextPartySubIDs) GetContextPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.ContextPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContextPartySubIDType gets ContextPartySubIDType, Tag 1528
func (m NoContextPartySubIDs) GetContextPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.ContextPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasContextPartySubID returns true if ContextPartySubID is present, Tag 1527
func (m NoContextPartySubIDs) HasContextPartySubID() bool {
	return m.Has(tag.ContextPartySubID)
}

//HasContextPartySubIDType returns true if ContextPartySubIDType is present, Tag 1528
func (m NoContextPartySubIDs) HasContextPartySubIDType() bool {
	return m.Has(tag.ContextPartySubIDType)
}

//NoContextPartySubIDsRepeatingGroup is a repeating group, Tag 1526
type NoContextPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoContextPartySubIDsRepeatingGroup returns an initialized, NoContextPartySubIDsRepeatingGroup
func NewNoContextPartySubIDsRepeatingGroup() NoContextPartySubIDsRepeatingGroup {
	return NoContextPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoContextPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ContextPartySubID), quickfix.GroupElement(tag.ContextPartySubIDType)})}
}

//Add create and append a new NoContextPartySubIDs to this group
func (m NoContextPartySubIDsRepeatingGroup) Add() NoContextPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoContextPartySubIDs{g}
}

//Get returns the ith NoContextPartySubIDs in the NoContextPartySubIDsRepeatinGroup
func (m NoContextPartySubIDsRepeatingGroup) Get(i int) NoContextPartySubIDs {
	return NoContextPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoContextPartyIDsRepeatingGroup is a repeating group, Tag 1522
type NoContextPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoContextPartyIDsRepeatingGroup returns an initialized, NoContextPartyIDsRepeatingGroup
func NewNoContextPartyIDsRepeatingGroup() NoContextPartyIDsRepeatingGroup {
	return NoContextPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoContextPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ContextPartyID), quickfix.GroupElement(tag.ContextPartyIDSource), quickfix.GroupElement(tag.ContextPartyRole), NewNoContextPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoContextPartyIDs to this group
func (m NoContextPartyIDsRepeatingGroup) Add() NoContextPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoContextPartyIDs{g}
}

//Get returns the ith NoContextPartyIDs in the NoContextPartyIDsRepeatinGroup
func (m NoContextPartyIDsRepeatingGroup) Get(i int) NoContextPartyIDs {
	return NoContextPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoRiskLimits is a repeating group element, Tag 1529
type NoRiskLimits struct {
	*quickfix.Group
}

//SetRiskLimitType sets RiskLimitType, Tag 1530
func (m NoRiskLimits) SetRiskLimitType(v enum.RiskLimitType) {
	m.Set(field.NewRiskLimitType(v))
}

//SetRiskLimitAmount sets RiskLimitAmount, Tag 1531
func (m NoRiskLimits) SetRiskLimitAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewRiskLimitAmount(value, scale))
}

//SetRiskLimitCurrency sets RiskLimitCurrency, Tag 1532
func (m NoRiskLimits) SetRiskLimitCurrency(v string) {
	m.Set(field.NewRiskLimitCurrency(v))
}

//SetRiskLimitPlatform sets RiskLimitPlatform, Tag 1533
func (m NoRiskLimits) SetRiskLimitPlatform(v string) {
	m.Set(field.NewRiskLimitPlatform(v))
}

//SetNoRiskInstruments sets NoRiskInstruments, Tag 1534
func (m NoRiskLimits) SetNoRiskInstruments(f NoRiskInstrumentsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoRiskWarningLevels sets NoRiskWarningLevels, Tag 1559
func (m NoRiskLimits) SetNoRiskWarningLevels(f NoRiskWarningLevelsRepeatingGroup) {
	m.SetGroup(f)
}

//GetRiskLimitType gets RiskLimitType, Tag 1530
func (m NoRiskLimits) GetRiskLimitType() (v enum.RiskLimitType, err quickfix.MessageRejectError) {
	var f field.RiskLimitTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskLimitAmount gets RiskLimitAmount, Tag 1531
func (m NoRiskLimits) GetRiskLimitAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RiskLimitAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskLimitCurrency gets RiskLimitCurrency, Tag 1532
func (m NoRiskLimits) GetRiskLimitCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.RiskLimitCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskLimitPlatform gets RiskLimitPlatform, Tag 1533
func (m NoRiskLimits) GetRiskLimitPlatform() (v string, err quickfix.MessageRejectError) {
	var f field.RiskLimitPlatformField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoRiskInstruments gets NoRiskInstruments, Tag 1534
func (m NoRiskLimits) GetNoRiskInstruments() (NoRiskInstrumentsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRiskInstrumentsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoRiskWarningLevels gets NoRiskWarningLevels, Tag 1559
func (m NoRiskLimits) GetNoRiskWarningLevels() (NoRiskWarningLevelsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRiskWarningLevelsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasRiskLimitType returns true if RiskLimitType is present, Tag 1530
func (m NoRiskLimits) HasRiskLimitType() bool {
	return m.Has(tag.RiskLimitType)
}

//HasRiskLimitAmount returns true if RiskLimitAmount is present, Tag 1531
func (m NoRiskLimits) HasRiskLimitAmount() bool {
	return m.Has(tag.RiskLimitAmount)
}

//HasRiskLimitCurrency returns true if RiskLimitCurrency is present, Tag 1532
func (m NoRiskLimits) HasRiskLimitCurrency() bool {
	return m.Has(tag.RiskLimitCurrency)
}

//HasRiskLimitPlatform returns true if RiskLimitPlatform is present, Tag 1533
func (m NoRiskLimits) HasRiskLimitPlatform() bool {
	return m.Has(tag.RiskLimitPlatform)
}

//HasNoRiskInstruments returns true if NoRiskInstruments is present, Tag 1534
func (m NoRiskLimits) HasNoRiskInstruments() bool {
	return m.Has(tag.NoRiskInstruments)
}

//HasNoRiskWarningLevels returns true if NoRiskWarningLevels is present, Tag 1559
func (m NoRiskLimits) HasNoRiskWarningLevels() bool {
	return m.Has(tag.NoRiskWarningLevels)
}

//NoRiskInstruments is a repeating group element, Tag 1534
type NoRiskInstruments struct {
	*quickfix.Group
}

//SetRiskInstrumentOperator sets RiskInstrumentOperator, Tag 1535
func (m NoRiskInstruments) SetRiskInstrumentOperator(v enum.RiskInstrumentOperator) {
	m.Set(field.NewRiskInstrumentOperator(v))
}

//SetRiskSymbol sets RiskSymbol, Tag 1536
func (m NoRiskInstruments) SetRiskSymbol(v string) {
	m.Set(field.NewRiskSymbol(v))
}

//SetRiskSymbolSfx sets RiskSymbolSfx, Tag 1537
func (m NoRiskInstruments) SetRiskSymbolSfx(v string) {
	m.Set(field.NewRiskSymbolSfx(v))
}

//SetRiskSecurityID sets RiskSecurityID, Tag 1538
func (m NoRiskInstruments) SetRiskSecurityID(v string) {
	m.Set(field.NewRiskSecurityID(v))
}

//SetRiskSecurityIDSource sets RiskSecurityIDSource, Tag 1539
func (m NoRiskInstruments) SetRiskSecurityIDSource(v string) {
	m.Set(field.NewRiskSecurityIDSource(v))
}

//SetNoRiskSecurityAltID sets NoRiskSecurityAltID, Tag 1540
func (m NoRiskInstruments) SetNoRiskSecurityAltID(f NoRiskSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetRiskProduct sets RiskProduct, Tag 1543
func (m NoRiskInstruments) SetRiskProduct(v int) {
	m.Set(field.NewRiskProduct(v))
}

//SetRiskProductComplex sets RiskProductComplex, Tag 1544
func (m NoRiskInstruments) SetRiskProductComplex(v string) {
	m.Set(field.NewRiskProductComplex(v))
}

//SetRiskSecurityGroup sets RiskSecurityGroup, Tag 1545
func (m NoRiskInstruments) SetRiskSecurityGroup(v string) {
	m.Set(field.NewRiskSecurityGroup(v))
}

//SetRiskCFICode sets RiskCFICode, Tag 1546
func (m NoRiskInstruments) SetRiskCFICode(v string) {
	m.Set(field.NewRiskCFICode(v))
}

//SetRiskSecurityType sets RiskSecurityType, Tag 1547
func (m NoRiskInstruments) SetRiskSecurityType(v string) {
	m.Set(field.NewRiskSecurityType(v))
}

//SetRiskSecuritySubType sets RiskSecuritySubType, Tag 1548
func (m NoRiskInstruments) SetRiskSecuritySubType(v string) {
	m.Set(field.NewRiskSecuritySubType(v))
}

//SetRiskMaturityMonthYear sets RiskMaturityMonthYear, Tag 1549
func (m NoRiskInstruments) SetRiskMaturityMonthYear(v string) {
	m.Set(field.NewRiskMaturityMonthYear(v))
}

//SetRiskMaturityTime sets RiskMaturityTime, Tag 1550
func (m NoRiskInstruments) SetRiskMaturityTime(v string) {
	m.Set(field.NewRiskMaturityTime(v))
}

//SetRiskRestructuringType sets RiskRestructuringType, Tag 1551
func (m NoRiskInstruments) SetRiskRestructuringType(v string) {
	m.Set(field.NewRiskRestructuringType(v))
}

//SetRiskSeniority sets RiskSeniority, Tag 1552
func (m NoRiskInstruments) SetRiskSeniority(v string) {
	m.Set(field.NewRiskSeniority(v))
}

//SetRiskPutOrCall sets RiskPutOrCall, Tag 1553
func (m NoRiskInstruments) SetRiskPutOrCall(v int) {
	m.Set(field.NewRiskPutOrCall(v))
}

//SetRiskFlexibleIndicator sets RiskFlexibleIndicator, Tag 1554
func (m NoRiskInstruments) SetRiskFlexibleIndicator(v bool) {
	m.Set(field.NewRiskFlexibleIndicator(v))
}

//SetRiskCouponRate sets RiskCouponRate, Tag 1555
func (m NoRiskInstruments) SetRiskCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRiskCouponRate(value, scale))
}

//SetRiskSecurityExchange sets RiskSecurityExchange, Tag 1616
func (m NoRiskInstruments) SetRiskSecurityExchange(v string) {
	m.Set(field.NewRiskSecurityExchange(v))
}

//SetRiskSecurityDesc sets RiskSecurityDesc, Tag 1556
func (m NoRiskInstruments) SetRiskSecurityDesc(v string) {
	m.Set(field.NewRiskSecurityDesc(v))
}

//SetRiskEncodedSecurityDescLen sets RiskEncodedSecurityDescLen, Tag 1620
func (m NoRiskInstruments) SetRiskEncodedSecurityDescLen(v int) {
	m.Set(field.NewRiskEncodedSecurityDescLen(v))
}

//SetRiskEncodedSecurityDesc sets RiskEncodedSecurityDesc, Tag 1621
func (m NoRiskInstruments) SetRiskEncodedSecurityDesc(v string) {
	m.Set(field.NewRiskEncodedSecurityDesc(v))
}

//SetRiskInstrumentSettlType sets RiskInstrumentSettlType, Tag 1557
func (m NoRiskInstruments) SetRiskInstrumentSettlType(v string) {
	m.Set(field.NewRiskInstrumentSettlType(v))
}

//SetRiskInstrumentMultiplier sets RiskInstrumentMultiplier, Tag 1558
func (m NoRiskInstruments) SetRiskInstrumentMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewRiskInstrumentMultiplier(value, scale))
}

//GetRiskInstrumentOperator gets RiskInstrumentOperator, Tag 1535
func (m NoRiskInstruments) GetRiskInstrumentOperator() (v enum.RiskInstrumentOperator, err quickfix.MessageRejectError) {
	var f field.RiskInstrumentOperatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskSymbol gets RiskSymbol, Tag 1536
func (m NoRiskInstruments) GetRiskSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.RiskSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskSymbolSfx gets RiskSymbolSfx, Tag 1537
func (m NoRiskInstruments) GetRiskSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.RiskSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskSecurityID gets RiskSecurityID, Tag 1538
func (m NoRiskInstruments) GetRiskSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.RiskSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskSecurityIDSource gets RiskSecurityIDSource, Tag 1539
func (m NoRiskInstruments) GetRiskSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.RiskSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoRiskSecurityAltID gets NoRiskSecurityAltID, Tag 1540
func (m NoRiskInstruments) GetNoRiskSecurityAltID() (NoRiskSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRiskSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetRiskProduct gets RiskProduct, Tag 1543
func (m NoRiskInstruments) GetRiskProduct() (v int, err quickfix.MessageRejectError) {
	var f field.RiskProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskProductComplex gets RiskProductComplex, Tag 1544
func (m NoRiskInstruments) GetRiskProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.RiskProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskSecurityGroup gets RiskSecurityGroup, Tag 1545
func (m NoRiskInstruments) GetRiskSecurityGroup() (v string, err quickfix.MessageRejectError) {
	var f field.RiskSecurityGroupField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskCFICode gets RiskCFICode, Tag 1546
func (m NoRiskInstruments) GetRiskCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.RiskCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskSecurityType gets RiskSecurityType, Tag 1547
func (m NoRiskInstruments) GetRiskSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.RiskSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskSecuritySubType gets RiskSecuritySubType, Tag 1548
func (m NoRiskInstruments) GetRiskSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.RiskSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskMaturityMonthYear gets RiskMaturityMonthYear, Tag 1549
func (m NoRiskInstruments) GetRiskMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.RiskMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskMaturityTime gets RiskMaturityTime, Tag 1550
func (m NoRiskInstruments) GetRiskMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.RiskMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskRestructuringType gets RiskRestructuringType, Tag 1551
func (m NoRiskInstruments) GetRiskRestructuringType() (v string, err quickfix.MessageRejectError) {
	var f field.RiskRestructuringTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskSeniority gets RiskSeniority, Tag 1552
func (m NoRiskInstruments) GetRiskSeniority() (v string, err quickfix.MessageRejectError) {
	var f field.RiskSeniorityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskPutOrCall gets RiskPutOrCall, Tag 1553
func (m NoRiskInstruments) GetRiskPutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.RiskPutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskFlexibleIndicator gets RiskFlexibleIndicator, Tag 1554
func (m NoRiskInstruments) GetRiskFlexibleIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.RiskFlexibleIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskCouponRate gets RiskCouponRate, Tag 1555
func (m NoRiskInstruments) GetRiskCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RiskCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskSecurityExchange gets RiskSecurityExchange, Tag 1616
func (m NoRiskInstruments) GetRiskSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.RiskSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskSecurityDesc gets RiskSecurityDesc, Tag 1556
func (m NoRiskInstruments) GetRiskSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.RiskSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskEncodedSecurityDescLen gets RiskEncodedSecurityDescLen, Tag 1620
func (m NoRiskInstruments) GetRiskEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.RiskEncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskEncodedSecurityDesc gets RiskEncodedSecurityDesc, Tag 1621
func (m NoRiskInstruments) GetRiskEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.RiskEncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskInstrumentSettlType gets RiskInstrumentSettlType, Tag 1557
func (m NoRiskInstruments) GetRiskInstrumentSettlType() (v string, err quickfix.MessageRejectError) {
	var f field.RiskInstrumentSettlTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskInstrumentMultiplier gets RiskInstrumentMultiplier, Tag 1558
func (m NoRiskInstruments) GetRiskInstrumentMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RiskInstrumentMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasRiskInstrumentOperator returns true if RiskInstrumentOperator is present, Tag 1535
func (m NoRiskInstruments) HasRiskInstrumentOperator() bool {
	return m.Has(tag.RiskInstrumentOperator)
}

//HasRiskSymbol returns true if RiskSymbol is present, Tag 1536
func (m NoRiskInstruments) HasRiskSymbol() bool {
	return m.Has(tag.RiskSymbol)
}

//HasRiskSymbolSfx returns true if RiskSymbolSfx is present, Tag 1537
func (m NoRiskInstruments) HasRiskSymbolSfx() bool {
	return m.Has(tag.RiskSymbolSfx)
}

//HasRiskSecurityID returns true if RiskSecurityID is present, Tag 1538
func (m NoRiskInstruments) HasRiskSecurityID() bool {
	return m.Has(tag.RiskSecurityID)
}

//HasRiskSecurityIDSource returns true if RiskSecurityIDSource is present, Tag 1539
func (m NoRiskInstruments) HasRiskSecurityIDSource() bool {
	return m.Has(tag.RiskSecurityIDSource)
}

//HasNoRiskSecurityAltID returns true if NoRiskSecurityAltID is present, Tag 1540
func (m NoRiskInstruments) HasNoRiskSecurityAltID() bool {
	return m.Has(tag.NoRiskSecurityAltID)
}

//HasRiskProduct returns true if RiskProduct is present, Tag 1543
func (m NoRiskInstruments) HasRiskProduct() bool {
	return m.Has(tag.RiskProduct)
}

//HasRiskProductComplex returns true if RiskProductComplex is present, Tag 1544
func (m NoRiskInstruments) HasRiskProductComplex() bool {
	return m.Has(tag.RiskProductComplex)
}

//HasRiskSecurityGroup returns true if RiskSecurityGroup is present, Tag 1545
func (m NoRiskInstruments) HasRiskSecurityGroup() bool {
	return m.Has(tag.RiskSecurityGroup)
}

//HasRiskCFICode returns true if RiskCFICode is present, Tag 1546
func (m NoRiskInstruments) HasRiskCFICode() bool {
	return m.Has(tag.RiskCFICode)
}

//HasRiskSecurityType returns true if RiskSecurityType is present, Tag 1547
func (m NoRiskInstruments) HasRiskSecurityType() bool {
	return m.Has(tag.RiskSecurityType)
}

//HasRiskSecuritySubType returns true if RiskSecuritySubType is present, Tag 1548
func (m NoRiskInstruments) HasRiskSecuritySubType() bool {
	return m.Has(tag.RiskSecuritySubType)
}

//HasRiskMaturityMonthYear returns true if RiskMaturityMonthYear is present, Tag 1549
func (m NoRiskInstruments) HasRiskMaturityMonthYear() bool {
	return m.Has(tag.RiskMaturityMonthYear)
}

//HasRiskMaturityTime returns true if RiskMaturityTime is present, Tag 1550
func (m NoRiskInstruments) HasRiskMaturityTime() bool {
	return m.Has(tag.RiskMaturityTime)
}

//HasRiskRestructuringType returns true if RiskRestructuringType is present, Tag 1551
func (m NoRiskInstruments) HasRiskRestructuringType() bool {
	return m.Has(tag.RiskRestructuringType)
}

//HasRiskSeniority returns true if RiskSeniority is present, Tag 1552
func (m NoRiskInstruments) HasRiskSeniority() bool {
	return m.Has(tag.RiskSeniority)
}

//HasRiskPutOrCall returns true if RiskPutOrCall is present, Tag 1553
func (m NoRiskInstruments) HasRiskPutOrCall() bool {
	return m.Has(tag.RiskPutOrCall)
}

//HasRiskFlexibleIndicator returns true if RiskFlexibleIndicator is present, Tag 1554
func (m NoRiskInstruments) HasRiskFlexibleIndicator() bool {
	return m.Has(tag.RiskFlexibleIndicator)
}

//HasRiskCouponRate returns true if RiskCouponRate is present, Tag 1555
func (m NoRiskInstruments) HasRiskCouponRate() bool {
	return m.Has(tag.RiskCouponRate)
}

//HasRiskSecurityExchange returns true if RiskSecurityExchange is present, Tag 1616
func (m NoRiskInstruments) HasRiskSecurityExchange() bool {
	return m.Has(tag.RiskSecurityExchange)
}

//HasRiskSecurityDesc returns true if RiskSecurityDesc is present, Tag 1556
func (m NoRiskInstruments) HasRiskSecurityDesc() bool {
	return m.Has(tag.RiskSecurityDesc)
}

//HasRiskEncodedSecurityDescLen returns true if RiskEncodedSecurityDescLen is present, Tag 1620
func (m NoRiskInstruments) HasRiskEncodedSecurityDescLen() bool {
	return m.Has(tag.RiskEncodedSecurityDescLen)
}

//HasRiskEncodedSecurityDesc returns true if RiskEncodedSecurityDesc is present, Tag 1621
func (m NoRiskInstruments) HasRiskEncodedSecurityDesc() bool {
	return m.Has(tag.RiskEncodedSecurityDesc)
}

//HasRiskInstrumentSettlType returns true if RiskInstrumentSettlType is present, Tag 1557
func (m NoRiskInstruments) HasRiskInstrumentSettlType() bool {
	return m.Has(tag.RiskInstrumentSettlType)
}

//HasRiskInstrumentMultiplier returns true if RiskInstrumentMultiplier is present, Tag 1558
func (m NoRiskInstruments) HasRiskInstrumentMultiplier() bool {
	return m.Has(tag.RiskInstrumentMultiplier)
}

//NoRiskSecurityAltID is a repeating group element, Tag 1540
type NoRiskSecurityAltID struct {
	*quickfix.Group
}

//SetRiskSecurityAltID sets RiskSecurityAltID, Tag 1541
func (m NoRiskSecurityAltID) SetRiskSecurityAltID(v string) {
	m.Set(field.NewRiskSecurityAltID(v))
}

//SetRiskSecurityAltIDSource sets RiskSecurityAltIDSource, Tag 1542
func (m NoRiskSecurityAltID) SetRiskSecurityAltIDSource(v string) {
	m.Set(field.NewRiskSecurityAltIDSource(v))
}

//GetRiskSecurityAltID gets RiskSecurityAltID, Tag 1541
func (m NoRiskSecurityAltID) GetRiskSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.RiskSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskSecurityAltIDSource gets RiskSecurityAltIDSource, Tag 1542
func (m NoRiskSecurityAltID) GetRiskSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.RiskSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasRiskSecurityAltID returns true if RiskSecurityAltID is present, Tag 1541
func (m NoRiskSecurityAltID) HasRiskSecurityAltID() bool {
	return m.Has(tag.RiskSecurityAltID)
}

//HasRiskSecurityAltIDSource returns true if RiskSecurityAltIDSource is present, Tag 1542
func (m NoRiskSecurityAltID) HasRiskSecurityAltIDSource() bool {
	return m.Has(tag.RiskSecurityAltIDSource)
}

//NoRiskSecurityAltIDRepeatingGroup is a repeating group, Tag 1540
type NoRiskSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRiskSecurityAltIDRepeatingGroup returns an initialized, NoRiskSecurityAltIDRepeatingGroup
func NewNoRiskSecurityAltIDRepeatingGroup() NoRiskSecurityAltIDRepeatingGroup {
	return NoRiskSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRiskSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RiskSecurityAltID), quickfix.GroupElement(tag.RiskSecurityAltIDSource)})}
}

//Add create and append a new NoRiskSecurityAltID to this group
func (m NoRiskSecurityAltIDRepeatingGroup) Add() NoRiskSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoRiskSecurityAltID{g}
}

//Get returns the ith NoRiskSecurityAltID in the NoRiskSecurityAltIDRepeatinGroup
func (m NoRiskSecurityAltIDRepeatingGroup) Get(i int) NoRiskSecurityAltID {
	return NoRiskSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoRiskInstrumentsRepeatingGroup is a repeating group, Tag 1534
type NoRiskInstrumentsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRiskInstrumentsRepeatingGroup returns an initialized, NoRiskInstrumentsRepeatingGroup
func NewNoRiskInstrumentsRepeatingGroup() NoRiskInstrumentsRepeatingGroup {
	return NoRiskInstrumentsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRiskInstruments,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RiskInstrumentOperator), quickfix.GroupElement(tag.RiskSymbol), quickfix.GroupElement(tag.RiskSymbolSfx), quickfix.GroupElement(tag.RiskSecurityID), quickfix.GroupElement(tag.RiskSecurityIDSource), NewNoRiskSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.RiskProduct), quickfix.GroupElement(tag.RiskProductComplex), quickfix.GroupElement(tag.RiskSecurityGroup), quickfix.GroupElement(tag.RiskCFICode), quickfix.GroupElement(tag.RiskSecurityType), quickfix.GroupElement(tag.RiskSecuritySubType), quickfix.GroupElement(tag.RiskMaturityMonthYear), quickfix.GroupElement(tag.RiskMaturityTime), quickfix.GroupElement(tag.RiskRestructuringType), quickfix.GroupElement(tag.RiskSeniority), quickfix.GroupElement(tag.RiskPutOrCall), quickfix.GroupElement(tag.RiskFlexibleIndicator), quickfix.GroupElement(tag.RiskCouponRate), quickfix.GroupElement(tag.RiskSecurityExchange), quickfix.GroupElement(tag.RiskSecurityDesc), quickfix.GroupElement(tag.RiskEncodedSecurityDescLen), quickfix.GroupElement(tag.RiskEncodedSecurityDesc), quickfix.GroupElement(tag.RiskInstrumentSettlType), quickfix.GroupElement(tag.RiskInstrumentMultiplier)})}
}

//Add create and append a new NoRiskInstruments to this group
func (m NoRiskInstrumentsRepeatingGroup) Add() NoRiskInstruments {
	g := m.RepeatingGroup.Add()
	return NoRiskInstruments{g}
}

//Get returns the ith NoRiskInstruments in the NoRiskInstrumentsRepeatinGroup
func (m NoRiskInstrumentsRepeatingGroup) Get(i int) NoRiskInstruments {
	return NoRiskInstruments{m.RepeatingGroup.Get(i)}
}

//NoRiskWarningLevels is a repeating group element, Tag 1559
type NoRiskWarningLevels struct {
	*quickfix.Group
}

//SetRiskWarningLevelPercent sets RiskWarningLevelPercent, Tag 1560
func (m NoRiskWarningLevels) SetRiskWarningLevelPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewRiskWarningLevelPercent(value, scale))
}

//SetRiskWarningLevelName sets RiskWarningLevelName, Tag 1561
func (m NoRiskWarningLevels) SetRiskWarningLevelName(v string) {
	m.Set(field.NewRiskWarningLevelName(v))
}

//GetRiskWarningLevelPercent gets RiskWarningLevelPercent, Tag 1560
func (m NoRiskWarningLevels) GetRiskWarningLevelPercent() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RiskWarningLevelPercentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRiskWarningLevelName gets RiskWarningLevelName, Tag 1561
func (m NoRiskWarningLevels) GetRiskWarningLevelName() (v string, err quickfix.MessageRejectError) {
	var f field.RiskWarningLevelNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasRiskWarningLevelPercent returns true if RiskWarningLevelPercent is present, Tag 1560
func (m NoRiskWarningLevels) HasRiskWarningLevelPercent() bool {
	return m.Has(tag.RiskWarningLevelPercent)
}

//HasRiskWarningLevelName returns true if RiskWarningLevelName is present, Tag 1561
func (m NoRiskWarningLevels) HasRiskWarningLevelName() bool {
	return m.Has(tag.RiskWarningLevelName)
}

//NoRiskWarningLevelsRepeatingGroup is a repeating group, Tag 1559
type NoRiskWarningLevelsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRiskWarningLevelsRepeatingGroup returns an initialized, NoRiskWarningLevelsRepeatingGroup
func NewNoRiskWarningLevelsRepeatingGroup() NoRiskWarningLevelsRepeatingGroup {
	return NoRiskWarningLevelsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRiskWarningLevels,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RiskWarningLevelPercent), quickfix.GroupElement(tag.RiskWarningLevelName)})}
}

//Add create and append a new NoRiskWarningLevels to this group
func (m NoRiskWarningLevelsRepeatingGroup) Add() NoRiskWarningLevels {
	g := m.RepeatingGroup.Add()
	return NoRiskWarningLevels{g}
}

//Get returns the ith NoRiskWarningLevels in the NoRiskWarningLevelsRepeatinGroup
func (m NoRiskWarningLevelsRepeatingGroup) Get(i int) NoRiskWarningLevels {
	return NoRiskWarningLevels{m.RepeatingGroup.Get(i)}
}

//NoRiskLimitsRepeatingGroup is a repeating group, Tag 1529
type NoRiskLimitsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRiskLimitsRepeatingGroup returns an initialized, NoRiskLimitsRepeatingGroup
func NewNoRiskLimitsRepeatingGroup() NoRiskLimitsRepeatingGroup {
	return NoRiskLimitsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRiskLimits,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RiskLimitType), quickfix.GroupElement(tag.RiskLimitAmount), quickfix.GroupElement(tag.RiskLimitCurrency), quickfix.GroupElement(tag.RiskLimitPlatform), NewNoRiskInstrumentsRepeatingGroup(), NewNoRiskWarningLevelsRepeatingGroup()})}
}

//Add create and append a new NoRiskLimits to this group
func (m NoRiskLimitsRepeatingGroup) Add() NoRiskLimits {
	g := m.RepeatingGroup.Add()
	return NoRiskLimits{g}
}

//Get returns the ith NoRiskLimits in the NoRiskLimitsRepeatinGroup
func (m NoRiskLimitsRepeatingGroup) Get(i int) NoRiskLimits {
	return NoRiskLimits{m.RepeatingGroup.Get(i)}
}

//NoRelatedPartyIDs is a repeating group element, Tag 1562
type NoRelatedPartyIDs struct {
	*quickfix.Group
}

//SetRelatedPartyID sets RelatedPartyID, Tag 1563
func (m NoRelatedPartyIDs) SetRelatedPartyID(v string) {
	m.Set(field.NewRelatedPartyID(v))
}

//SetRelatedPartyIDSource sets RelatedPartyIDSource, Tag 1564
func (m NoRelatedPartyIDs) SetRelatedPartyIDSource(v string) {
	m.Set(field.NewRelatedPartyIDSource(v))
}

//SetRelatedPartyRole sets RelatedPartyRole, Tag 1565
func (m NoRelatedPartyIDs) SetRelatedPartyRole(v int) {
	m.Set(field.NewRelatedPartyRole(v))
}

//SetNoRelatedPartySubIDs sets NoRelatedPartySubIDs, Tag 1566
func (m NoRelatedPartyIDs) SetNoRelatedPartySubIDs(f NoRelatedPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoRelatedPartyAltIDs sets NoRelatedPartyAltIDs, Tag 1569
func (m NoRelatedPartyIDs) SetNoRelatedPartyAltIDs(f NoRelatedPartyAltIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoRelatedContextPartyIDs sets NoRelatedContextPartyIDs, Tag 1575
func (m NoRelatedPartyIDs) SetNoRelatedContextPartyIDs(f NoRelatedContextPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoRelationshipRiskLimits sets NoRelationshipRiskLimits, Tag 1582
func (m NoRelatedPartyIDs) SetNoRelationshipRiskLimits(f NoRelationshipRiskLimitsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoPartyRelationships sets NoPartyRelationships, Tag 1514
func (m NoRelatedPartyIDs) SetNoPartyRelationships(f NoPartyRelationshipsRepeatingGroup) {
	m.SetGroup(f)
}

//GetRelatedPartyID gets RelatedPartyID, Tag 1563
func (m NoRelatedPartyIDs) GetRelatedPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelatedPartyIDSource gets RelatedPartyIDSource, Tag 1564
func (m NoRelatedPartyIDs) GetRelatedPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelatedPartyRole gets RelatedPartyRole, Tag 1565
func (m NoRelatedPartyIDs) GetRelatedPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.RelatedPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoRelatedPartySubIDs gets NoRelatedPartySubIDs, Tag 1566
func (m NoRelatedPartyIDs) GetNoRelatedPartySubIDs() (NoRelatedPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoRelatedPartyAltIDs gets NoRelatedPartyAltIDs, Tag 1569
func (m NoRelatedPartyIDs) GetNoRelatedPartyAltIDs() (NoRelatedPartyAltIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedPartyAltIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoRelatedContextPartyIDs gets NoRelatedContextPartyIDs, Tag 1575
func (m NoRelatedPartyIDs) GetNoRelatedContextPartyIDs() (NoRelatedContextPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedContextPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoRelationshipRiskLimits gets NoRelationshipRiskLimits, Tag 1582
func (m NoRelatedPartyIDs) GetNoRelationshipRiskLimits() (NoRelationshipRiskLimitsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelationshipRiskLimitsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoPartyRelationships gets NoPartyRelationships, Tag 1514
func (m NoRelatedPartyIDs) GetNoPartyRelationships() (NoPartyRelationshipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyRelationshipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasRelatedPartyID returns true if RelatedPartyID is present, Tag 1563
func (m NoRelatedPartyIDs) HasRelatedPartyID() bool {
	return m.Has(tag.RelatedPartyID)
}

//HasRelatedPartyIDSource returns true if RelatedPartyIDSource is present, Tag 1564
func (m NoRelatedPartyIDs) HasRelatedPartyIDSource() bool {
	return m.Has(tag.RelatedPartyIDSource)
}

//HasRelatedPartyRole returns true if RelatedPartyRole is present, Tag 1565
func (m NoRelatedPartyIDs) HasRelatedPartyRole() bool {
	return m.Has(tag.RelatedPartyRole)
}

//HasNoRelatedPartySubIDs returns true if NoRelatedPartySubIDs is present, Tag 1566
func (m NoRelatedPartyIDs) HasNoRelatedPartySubIDs() bool {
	return m.Has(tag.NoRelatedPartySubIDs)
}

//HasNoRelatedPartyAltIDs returns true if NoRelatedPartyAltIDs is present, Tag 1569
func (m NoRelatedPartyIDs) HasNoRelatedPartyAltIDs() bool {
	return m.Has(tag.NoRelatedPartyAltIDs)
}

//HasNoRelatedContextPartyIDs returns true if NoRelatedContextPartyIDs is present, Tag 1575
func (m NoRelatedPartyIDs) HasNoRelatedContextPartyIDs() bool {
	return m.Has(tag.NoRelatedContextPartyIDs)
}

//HasNoRelationshipRiskLimits returns true if NoRelationshipRiskLimits is present, Tag 1582
func (m NoRelatedPartyIDs) HasNoRelationshipRiskLimits() bool {
	return m.Has(tag.NoRelationshipRiskLimits)
}

//HasNoPartyRelationships returns true if NoPartyRelationships is present, Tag 1514
func (m NoRelatedPartyIDs) HasNoPartyRelationships() bool {
	return m.Has(tag.NoPartyRelationships)
}

//NoRelatedPartySubIDs is a repeating group element, Tag 1566
type NoRelatedPartySubIDs struct {
	*quickfix.Group
}

//SetRelatedPartySubID sets RelatedPartySubID, Tag 1567
func (m NoRelatedPartySubIDs) SetRelatedPartySubID(v string) {
	m.Set(field.NewRelatedPartySubID(v))
}

//SetRelatedPartySubIDType sets RelatedPartySubIDType, Tag 1568
func (m NoRelatedPartySubIDs) SetRelatedPartySubIDType(v int) {
	m.Set(field.NewRelatedPartySubIDType(v))
}

//GetRelatedPartySubID gets RelatedPartySubID, Tag 1567
func (m NoRelatedPartySubIDs) GetRelatedPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelatedPartySubIDType gets RelatedPartySubIDType, Tag 1568
func (m NoRelatedPartySubIDs) GetRelatedPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.RelatedPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasRelatedPartySubID returns true if RelatedPartySubID is present, Tag 1567
func (m NoRelatedPartySubIDs) HasRelatedPartySubID() bool {
	return m.Has(tag.RelatedPartySubID)
}

//HasRelatedPartySubIDType returns true if RelatedPartySubIDType is present, Tag 1568
func (m NoRelatedPartySubIDs) HasRelatedPartySubIDType() bool {
	return m.Has(tag.RelatedPartySubIDType)
}

//NoRelatedPartySubIDsRepeatingGroup is a repeating group, Tag 1566
type NoRelatedPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelatedPartySubIDsRepeatingGroup returns an initialized, NoRelatedPartySubIDsRepeatingGroup
func NewNoRelatedPartySubIDsRepeatingGroup() NoRelatedPartySubIDsRepeatingGroup {
	return NoRelatedPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RelatedPartySubID), quickfix.GroupElement(tag.RelatedPartySubIDType)})}
}

//Add create and append a new NoRelatedPartySubIDs to this group
func (m NoRelatedPartySubIDsRepeatingGroup) Add() NoRelatedPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoRelatedPartySubIDs{g}
}

//Get returns the ith NoRelatedPartySubIDs in the NoRelatedPartySubIDsRepeatinGroup
func (m NoRelatedPartySubIDsRepeatingGroup) Get(i int) NoRelatedPartySubIDs {
	return NoRelatedPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoRelatedPartyAltIDs is a repeating group element, Tag 1569
type NoRelatedPartyAltIDs struct {
	*quickfix.Group
}

//SetRelatedPartyAltID sets RelatedPartyAltID, Tag 1570
func (m NoRelatedPartyAltIDs) SetRelatedPartyAltID(v string) {
	m.Set(field.NewRelatedPartyAltID(v))
}

//SetRelatedPartyAltIDSource sets RelatedPartyAltIDSource, Tag 1571
func (m NoRelatedPartyAltIDs) SetRelatedPartyAltIDSource(v string) {
	m.Set(field.NewRelatedPartyAltIDSource(v))
}

//SetNoRelatedPartyAltSubIDs sets NoRelatedPartyAltSubIDs, Tag 1572
func (m NoRelatedPartyAltIDs) SetNoRelatedPartyAltSubIDs(f NoRelatedPartyAltSubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetRelatedPartyAltID gets RelatedPartyAltID, Tag 1570
func (m NoRelatedPartyAltIDs) GetRelatedPartyAltID() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelatedPartyAltIDSource gets RelatedPartyAltIDSource, Tag 1571
func (m NoRelatedPartyAltIDs) GetRelatedPartyAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoRelatedPartyAltSubIDs gets NoRelatedPartyAltSubIDs, Tag 1572
func (m NoRelatedPartyAltIDs) GetNoRelatedPartyAltSubIDs() (NoRelatedPartyAltSubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedPartyAltSubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasRelatedPartyAltID returns true if RelatedPartyAltID is present, Tag 1570
func (m NoRelatedPartyAltIDs) HasRelatedPartyAltID() bool {
	return m.Has(tag.RelatedPartyAltID)
}

//HasRelatedPartyAltIDSource returns true if RelatedPartyAltIDSource is present, Tag 1571
func (m NoRelatedPartyAltIDs) HasRelatedPartyAltIDSource() bool {
	return m.Has(tag.RelatedPartyAltIDSource)
}

//HasNoRelatedPartyAltSubIDs returns true if NoRelatedPartyAltSubIDs is present, Tag 1572
func (m NoRelatedPartyAltIDs) HasNoRelatedPartyAltSubIDs() bool {
	return m.Has(tag.NoRelatedPartyAltSubIDs)
}

//NoRelatedPartyAltSubIDs is a repeating group element, Tag 1572
type NoRelatedPartyAltSubIDs struct {
	*quickfix.Group
}

//SetRelatedPartyAltSubID sets RelatedPartyAltSubID, Tag 1573
func (m NoRelatedPartyAltSubIDs) SetRelatedPartyAltSubID(v string) {
	m.Set(field.NewRelatedPartyAltSubID(v))
}

//SetRelatedPartyAltSubIDType sets RelatedPartyAltSubIDType, Tag 1574
func (m NoRelatedPartyAltSubIDs) SetRelatedPartyAltSubIDType(v int) {
	m.Set(field.NewRelatedPartyAltSubIDType(v))
}

//GetRelatedPartyAltSubID gets RelatedPartyAltSubID, Tag 1573
func (m NoRelatedPartyAltSubIDs) GetRelatedPartyAltSubID() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedPartyAltSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelatedPartyAltSubIDType gets RelatedPartyAltSubIDType, Tag 1574
func (m NoRelatedPartyAltSubIDs) GetRelatedPartyAltSubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.RelatedPartyAltSubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasRelatedPartyAltSubID returns true if RelatedPartyAltSubID is present, Tag 1573
func (m NoRelatedPartyAltSubIDs) HasRelatedPartyAltSubID() bool {
	return m.Has(tag.RelatedPartyAltSubID)
}

//HasRelatedPartyAltSubIDType returns true if RelatedPartyAltSubIDType is present, Tag 1574
func (m NoRelatedPartyAltSubIDs) HasRelatedPartyAltSubIDType() bool {
	return m.Has(tag.RelatedPartyAltSubIDType)
}

//NoRelatedPartyAltSubIDsRepeatingGroup is a repeating group, Tag 1572
type NoRelatedPartyAltSubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelatedPartyAltSubIDsRepeatingGroup returns an initialized, NoRelatedPartyAltSubIDsRepeatingGroup
func NewNoRelatedPartyAltSubIDsRepeatingGroup() NoRelatedPartyAltSubIDsRepeatingGroup {
	return NoRelatedPartyAltSubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedPartyAltSubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RelatedPartyAltSubID), quickfix.GroupElement(tag.RelatedPartyAltSubIDType)})}
}

//Add create and append a new NoRelatedPartyAltSubIDs to this group
func (m NoRelatedPartyAltSubIDsRepeatingGroup) Add() NoRelatedPartyAltSubIDs {
	g := m.RepeatingGroup.Add()
	return NoRelatedPartyAltSubIDs{g}
}

//Get returns the ith NoRelatedPartyAltSubIDs in the NoRelatedPartyAltSubIDsRepeatinGroup
func (m NoRelatedPartyAltSubIDsRepeatingGroup) Get(i int) NoRelatedPartyAltSubIDs {
	return NoRelatedPartyAltSubIDs{m.RepeatingGroup.Get(i)}
}

//NoRelatedPartyAltIDsRepeatingGroup is a repeating group, Tag 1569
type NoRelatedPartyAltIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelatedPartyAltIDsRepeatingGroup returns an initialized, NoRelatedPartyAltIDsRepeatingGroup
func NewNoRelatedPartyAltIDsRepeatingGroup() NoRelatedPartyAltIDsRepeatingGroup {
	return NoRelatedPartyAltIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedPartyAltIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RelatedPartyAltID), quickfix.GroupElement(tag.RelatedPartyAltIDSource), NewNoRelatedPartyAltSubIDsRepeatingGroup()})}
}

//Add create and append a new NoRelatedPartyAltIDs to this group
func (m NoRelatedPartyAltIDsRepeatingGroup) Add() NoRelatedPartyAltIDs {
	g := m.RepeatingGroup.Add()
	return NoRelatedPartyAltIDs{g}
}

//Get returns the ith NoRelatedPartyAltIDs in the NoRelatedPartyAltIDsRepeatinGroup
func (m NoRelatedPartyAltIDsRepeatingGroup) Get(i int) NoRelatedPartyAltIDs {
	return NoRelatedPartyAltIDs{m.RepeatingGroup.Get(i)}
}

//NoRelatedContextPartyIDs is a repeating group element, Tag 1575
type NoRelatedContextPartyIDs struct {
	*quickfix.Group
}

//SetRelatedContextPartyID sets RelatedContextPartyID, Tag 1576
func (m NoRelatedContextPartyIDs) SetRelatedContextPartyID(v string) {
	m.Set(field.NewRelatedContextPartyID(v))
}

//SetRelatedContextPartyIDSource sets RelatedContextPartyIDSource, Tag 1577
func (m NoRelatedContextPartyIDs) SetRelatedContextPartyIDSource(v string) {
	m.Set(field.NewRelatedContextPartyIDSource(v))
}

//SetRelatedContextPartyRole sets RelatedContextPartyRole, Tag 1578
func (m NoRelatedContextPartyIDs) SetRelatedContextPartyRole(v int) {
	m.Set(field.NewRelatedContextPartyRole(v))
}

//SetNoRelatedContextPartySubIDs sets NoRelatedContextPartySubIDs, Tag 1579
func (m NoRelatedContextPartyIDs) SetNoRelatedContextPartySubIDs(f NoRelatedContextPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetRelatedContextPartyID gets RelatedContextPartyID, Tag 1576
func (m NoRelatedContextPartyIDs) GetRelatedContextPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedContextPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelatedContextPartyIDSource gets RelatedContextPartyIDSource, Tag 1577
func (m NoRelatedContextPartyIDs) GetRelatedContextPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedContextPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelatedContextPartyRole gets RelatedContextPartyRole, Tag 1578
func (m NoRelatedContextPartyIDs) GetRelatedContextPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.RelatedContextPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoRelatedContextPartySubIDs gets NoRelatedContextPartySubIDs, Tag 1579
func (m NoRelatedContextPartyIDs) GetNoRelatedContextPartySubIDs() (NoRelatedContextPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedContextPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasRelatedContextPartyID returns true if RelatedContextPartyID is present, Tag 1576
func (m NoRelatedContextPartyIDs) HasRelatedContextPartyID() bool {
	return m.Has(tag.RelatedContextPartyID)
}

//HasRelatedContextPartyIDSource returns true if RelatedContextPartyIDSource is present, Tag 1577
func (m NoRelatedContextPartyIDs) HasRelatedContextPartyIDSource() bool {
	return m.Has(tag.RelatedContextPartyIDSource)
}

//HasRelatedContextPartyRole returns true if RelatedContextPartyRole is present, Tag 1578
func (m NoRelatedContextPartyIDs) HasRelatedContextPartyRole() bool {
	return m.Has(tag.RelatedContextPartyRole)
}

//HasNoRelatedContextPartySubIDs returns true if NoRelatedContextPartySubIDs is present, Tag 1579
func (m NoRelatedContextPartyIDs) HasNoRelatedContextPartySubIDs() bool {
	return m.Has(tag.NoRelatedContextPartySubIDs)
}

//NoRelatedContextPartySubIDs is a repeating group element, Tag 1579
type NoRelatedContextPartySubIDs struct {
	*quickfix.Group
}

//SetRelatedContextPartySubID sets RelatedContextPartySubID, Tag 1580
func (m NoRelatedContextPartySubIDs) SetRelatedContextPartySubID(v string) {
	m.Set(field.NewRelatedContextPartySubID(v))
}

//SetRelatedContextPartySubIDType sets RelatedContextPartySubIDType, Tag 1581
func (m NoRelatedContextPartySubIDs) SetRelatedContextPartySubIDType(v int) {
	m.Set(field.NewRelatedContextPartySubIDType(v))
}

//GetRelatedContextPartySubID gets RelatedContextPartySubID, Tag 1580
func (m NoRelatedContextPartySubIDs) GetRelatedContextPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.RelatedContextPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelatedContextPartySubIDType gets RelatedContextPartySubIDType, Tag 1581
func (m NoRelatedContextPartySubIDs) GetRelatedContextPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.RelatedContextPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasRelatedContextPartySubID returns true if RelatedContextPartySubID is present, Tag 1580
func (m NoRelatedContextPartySubIDs) HasRelatedContextPartySubID() bool {
	return m.Has(tag.RelatedContextPartySubID)
}

//HasRelatedContextPartySubIDType returns true if RelatedContextPartySubIDType is present, Tag 1581
func (m NoRelatedContextPartySubIDs) HasRelatedContextPartySubIDType() bool {
	return m.Has(tag.RelatedContextPartySubIDType)
}

//NoRelatedContextPartySubIDsRepeatingGroup is a repeating group, Tag 1579
type NoRelatedContextPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelatedContextPartySubIDsRepeatingGroup returns an initialized, NoRelatedContextPartySubIDsRepeatingGroup
func NewNoRelatedContextPartySubIDsRepeatingGroup() NoRelatedContextPartySubIDsRepeatingGroup {
	return NoRelatedContextPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedContextPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RelatedContextPartySubID), quickfix.GroupElement(tag.RelatedContextPartySubIDType)})}
}

//Add create and append a new NoRelatedContextPartySubIDs to this group
func (m NoRelatedContextPartySubIDsRepeatingGroup) Add() NoRelatedContextPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoRelatedContextPartySubIDs{g}
}

//Get returns the ith NoRelatedContextPartySubIDs in the NoRelatedContextPartySubIDsRepeatinGroup
func (m NoRelatedContextPartySubIDsRepeatingGroup) Get(i int) NoRelatedContextPartySubIDs {
	return NoRelatedContextPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoRelatedContextPartyIDsRepeatingGroup is a repeating group, Tag 1575
type NoRelatedContextPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelatedContextPartyIDsRepeatingGroup returns an initialized, NoRelatedContextPartyIDsRepeatingGroup
func NewNoRelatedContextPartyIDsRepeatingGroup() NoRelatedContextPartyIDsRepeatingGroup {
	return NoRelatedContextPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedContextPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RelatedContextPartyID), quickfix.GroupElement(tag.RelatedContextPartyIDSource), quickfix.GroupElement(tag.RelatedContextPartyRole), NewNoRelatedContextPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoRelatedContextPartyIDs to this group
func (m NoRelatedContextPartyIDsRepeatingGroup) Add() NoRelatedContextPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoRelatedContextPartyIDs{g}
}

//Get returns the ith NoRelatedContextPartyIDs in the NoRelatedContextPartyIDsRepeatinGroup
func (m NoRelatedContextPartyIDsRepeatingGroup) Get(i int) NoRelatedContextPartyIDs {
	return NoRelatedContextPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoRelationshipRiskLimits is a repeating group element, Tag 1582
type NoRelationshipRiskLimits struct {
	*quickfix.Group
}

//SetRelationshipRiskLimitType sets RelationshipRiskLimitType, Tag 1583
func (m NoRelationshipRiskLimits) SetRelationshipRiskLimitType(v int) {
	m.Set(field.NewRelationshipRiskLimitType(v))
}

//SetRelationshipRiskLimitAmount sets RelationshipRiskLimitAmount, Tag 1584
func (m NoRelationshipRiskLimits) SetRelationshipRiskLimitAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewRelationshipRiskLimitAmount(value, scale))
}

//SetRelationshipRiskLimitCurrency sets RelationshipRiskLimitCurrency, Tag 1585
func (m NoRelationshipRiskLimits) SetRelationshipRiskLimitCurrency(v string) {
	m.Set(field.NewRelationshipRiskLimitCurrency(v))
}

//SetRelationshipRiskLimitPlatform sets RelationshipRiskLimitPlatform, Tag 1586
func (m NoRelationshipRiskLimits) SetRelationshipRiskLimitPlatform(v string) {
	m.Set(field.NewRelationshipRiskLimitPlatform(v))
}

//SetNoRelationshipRiskInstruments sets NoRelationshipRiskInstruments, Tag 1587
func (m NoRelationshipRiskLimits) SetNoRelationshipRiskInstruments(f NoRelationshipRiskInstrumentsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoRelationshipRiskWarningLevels sets NoRelationshipRiskWarningLevels, Tag 1613
func (m NoRelationshipRiskLimits) SetNoRelationshipRiskWarningLevels(f NoRelationshipRiskWarningLevelsRepeatingGroup) {
	m.SetGroup(f)
}

//GetRelationshipRiskLimitType gets RelationshipRiskLimitType, Tag 1583
func (m NoRelationshipRiskLimits) GetRelationshipRiskLimitType() (v int, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskLimitTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskLimitAmount gets RelationshipRiskLimitAmount, Tag 1584
func (m NoRelationshipRiskLimits) GetRelationshipRiskLimitAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskLimitAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskLimitCurrency gets RelationshipRiskLimitCurrency, Tag 1585
func (m NoRelationshipRiskLimits) GetRelationshipRiskLimitCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskLimitCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskLimitPlatform gets RelationshipRiskLimitPlatform, Tag 1586
func (m NoRelationshipRiskLimits) GetRelationshipRiskLimitPlatform() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskLimitPlatformField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoRelationshipRiskInstruments gets NoRelationshipRiskInstruments, Tag 1587
func (m NoRelationshipRiskLimits) GetNoRelationshipRiskInstruments() (NoRelationshipRiskInstrumentsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelationshipRiskInstrumentsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoRelationshipRiskWarningLevels gets NoRelationshipRiskWarningLevels, Tag 1613
func (m NoRelationshipRiskLimits) GetNoRelationshipRiskWarningLevels() (NoRelationshipRiskWarningLevelsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelationshipRiskWarningLevelsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasRelationshipRiskLimitType returns true if RelationshipRiskLimitType is present, Tag 1583
func (m NoRelationshipRiskLimits) HasRelationshipRiskLimitType() bool {
	return m.Has(tag.RelationshipRiskLimitType)
}

//HasRelationshipRiskLimitAmount returns true if RelationshipRiskLimitAmount is present, Tag 1584
func (m NoRelationshipRiskLimits) HasRelationshipRiskLimitAmount() bool {
	return m.Has(tag.RelationshipRiskLimitAmount)
}

//HasRelationshipRiskLimitCurrency returns true if RelationshipRiskLimitCurrency is present, Tag 1585
func (m NoRelationshipRiskLimits) HasRelationshipRiskLimitCurrency() bool {
	return m.Has(tag.RelationshipRiskLimitCurrency)
}

//HasRelationshipRiskLimitPlatform returns true if RelationshipRiskLimitPlatform is present, Tag 1586
func (m NoRelationshipRiskLimits) HasRelationshipRiskLimitPlatform() bool {
	return m.Has(tag.RelationshipRiskLimitPlatform)
}

//HasNoRelationshipRiskInstruments returns true if NoRelationshipRiskInstruments is present, Tag 1587
func (m NoRelationshipRiskLimits) HasNoRelationshipRiskInstruments() bool {
	return m.Has(tag.NoRelationshipRiskInstruments)
}

//HasNoRelationshipRiskWarningLevels returns true if NoRelationshipRiskWarningLevels is present, Tag 1613
func (m NoRelationshipRiskLimits) HasNoRelationshipRiskWarningLevels() bool {
	return m.Has(tag.NoRelationshipRiskWarningLevels)
}

//NoRelationshipRiskInstruments is a repeating group element, Tag 1587
type NoRelationshipRiskInstruments struct {
	*quickfix.Group
}

//SetRelationshipRiskInstrumentOperator sets RelationshipRiskInstrumentOperator, Tag 1588
func (m NoRelationshipRiskInstruments) SetRelationshipRiskInstrumentOperator(v int) {
	m.Set(field.NewRelationshipRiskInstrumentOperator(v))
}

//SetRelationshipRiskSymbol sets RelationshipRiskSymbol, Tag 1589
func (m NoRelationshipRiskInstruments) SetRelationshipRiskSymbol(v string) {
	m.Set(field.NewRelationshipRiskSymbol(v))
}

//SetRelationshipRiskSymbolSfx sets RelationshipRiskSymbolSfx, Tag 1590
func (m NoRelationshipRiskInstruments) SetRelationshipRiskSymbolSfx(v string) {
	m.Set(field.NewRelationshipRiskSymbolSfx(v))
}

//SetRelationshipRiskSecurityID sets RelationshipRiskSecurityID, Tag 1591
func (m NoRelationshipRiskInstruments) SetRelationshipRiskSecurityID(v string) {
	m.Set(field.NewRelationshipRiskSecurityID(v))
}

//SetRelationshipRiskSecurityIDSource sets RelationshipRiskSecurityIDSource, Tag 1592
func (m NoRelationshipRiskInstruments) SetRelationshipRiskSecurityIDSource(v string) {
	m.Set(field.NewRelationshipRiskSecurityIDSource(v))
}

//SetNoRelationshipRiskSecurityAltID sets NoRelationshipRiskSecurityAltID, Tag 1593
func (m NoRelationshipRiskInstruments) SetNoRelationshipRiskSecurityAltID(f NoRelationshipRiskSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetRelationshipRiskProduct sets RelationshipRiskProduct, Tag 1596
func (m NoRelationshipRiskInstruments) SetRelationshipRiskProduct(v int) {
	m.Set(field.NewRelationshipRiskProduct(v))
}

//SetRelationshipRiskProductComplex sets RelationshipRiskProductComplex, Tag 1597
func (m NoRelationshipRiskInstruments) SetRelationshipRiskProductComplex(v string) {
	m.Set(field.NewRelationshipRiskProductComplex(v))
}

//SetRelationshipRiskSecurityGroup sets RelationshipRiskSecurityGroup, Tag 1598
func (m NoRelationshipRiskInstruments) SetRelationshipRiskSecurityGroup(v string) {
	m.Set(field.NewRelationshipRiskSecurityGroup(v))
}

//SetRelationshipRiskCFICode sets RelationshipRiskCFICode, Tag 1599
func (m NoRelationshipRiskInstruments) SetRelationshipRiskCFICode(v string) {
	m.Set(field.NewRelationshipRiskCFICode(v))
}

//SetRelationshipRiskSecurityType sets RelationshipRiskSecurityType, Tag 1600
func (m NoRelationshipRiskInstruments) SetRelationshipRiskSecurityType(v string) {
	m.Set(field.NewRelationshipRiskSecurityType(v))
}

//SetRelationshipRiskSecuritySubType sets RelationshipRiskSecuritySubType, Tag 1601
func (m NoRelationshipRiskInstruments) SetRelationshipRiskSecuritySubType(v string) {
	m.Set(field.NewRelationshipRiskSecuritySubType(v))
}

//SetRelationshipRiskMaturityMonthYear sets RelationshipRiskMaturityMonthYear, Tag 1602
func (m NoRelationshipRiskInstruments) SetRelationshipRiskMaturityMonthYear(v string) {
	m.Set(field.NewRelationshipRiskMaturityMonthYear(v))
}

//SetRelationshipRiskMaturityTime sets RelationshipRiskMaturityTime, Tag 1603
func (m NoRelationshipRiskInstruments) SetRelationshipRiskMaturityTime(v string) {
	m.Set(field.NewRelationshipRiskMaturityTime(v))
}

//SetRelationshipRiskRestructuringType sets RelationshipRiskRestructuringType, Tag 1604
func (m NoRelationshipRiskInstruments) SetRelationshipRiskRestructuringType(v string) {
	m.Set(field.NewRelationshipRiskRestructuringType(v))
}

//SetRelationshipRiskSeniority sets RelationshipRiskSeniority, Tag 1605
func (m NoRelationshipRiskInstruments) SetRelationshipRiskSeniority(v string) {
	m.Set(field.NewRelationshipRiskSeniority(v))
}

//SetRelationshipRiskPutOrCall sets RelationshipRiskPutOrCall, Tag 1606
func (m NoRelationshipRiskInstruments) SetRelationshipRiskPutOrCall(v int) {
	m.Set(field.NewRelationshipRiskPutOrCall(v))
}

//SetRelationshipRiskFlexibleIndicator sets RelationshipRiskFlexibleIndicator, Tag 1607
func (m NoRelationshipRiskInstruments) SetRelationshipRiskFlexibleIndicator(v bool) {
	m.Set(field.NewRelationshipRiskFlexibleIndicator(v))
}

//SetRelationshipRiskCouponRate sets RelationshipRiskCouponRate, Tag 1608
func (m NoRelationshipRiskInstruments) SetRelationshipRiskCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRelationshipRiskCouponRate(value, scale))
}

//SetRelationshipRiskSecurityExchange sets RelationshipRiskSecurityExchange, Tag 1609
func (m NoRelationshipRiskInstruments) SetRelationshipRiskSecurityExchange(v string) {
	m.Set(field.NewRelationshipRiskSecurityExchange(v))
}

//SetRelationshipRiskSecurityDesc sets RelationshipRiskSecurityDesc, Tag 1610
func (m NoRelationshipRiskInstruments) SetRelationshipRiskSecurityDesc(v string) {
	m.Set(field.NewRelationshipRiskSecurityDesc(v))
}

//SetRelationshipRiskEncodedSecurityDescLen sets RelationshipRiskEncodedSecurityDescLen, Tag 1618
func (m NoRelationshipRiskInstruments) SetRelationshipRiskEncodedSecurityDescLen(v int) {
	m.Set(field.NewRelationshipRiskEncodedSecurityDescLen(v))
}

//SetRelationshipRiskEncodedSecurityDesc sets RelationshipRiskEncodedSecurityDesc, Tag 1619
func (m NoRelationshipRiskInstruments) SetRelationshipRiskEncodedSecurityDesc(v string) {
	m.Set(field.NewRelationshipRiskEncodedSecurityDesc(v))
}

//SetRelationshipRiskInstrumentSettlType sets RelationshipRiskInstrumentSettlType, Tag 1611
func (m NoRelationshipRiskInstruments) SetRelationshipRiskInstrumentSettlType(v string) {
	m.Set(field.NewRelationshipRiskInstrumentSettlType(v))
}

//SetRelationshipRiskInstrumentMultiplier sets RelationshipRiskInstrumentMultiplier, Tag 1612
func (m NoRelationshipRiskInstruments) SetRelationshipRiskInstrumentMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewRelationshipRiskInstrumentMultiplier(value, scale))
}

//GetRelationshipRiskInstrumentOperator gets RelationshipRiskInstrumentOperator, Tag 1588
func (m NoRelationshipRiskInstruments) GetRelationshipRiskInstrumentOperator() (v int, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskInstrumentOperatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskSymbol gets RelationshipRiskSymbol, Tag 1589
func (m NoRelationshipRiskInstruments) GetRelationshipRiskSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskSymbolSfx gets RelationshipRiskSymbolSfx, Tag 1590
func (m NoRelationshipRiskInstruments) GetRelationshipRiskSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskSecurityID gets RelationshipRiskSecurityID, Tag 1591
func (m NoRelationshipRiskInstruments) GetRelationshipRiskSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskSecurityIDSource gets RelationshipRiskSecurityIDSource, Tag 1592
func (m NoRelationshipRiskInstruments) GetRelationshipRiskSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoRelationshipRiskSecurityAltID gets NoRelationshipRiskSecurityAltID, Tag 1593
func (m NoRelationshipRiskInstruments) GetNoRelationshipRiskSecurityAltID() (NoRelationshipRiskSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelationshipRiskSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetRelationshipRiskProduct gets RelationshipRiskProduct, Tag 1596
func (m NoRelationshipRiskInstruments) GetRelationshipRiskProduct() (v int, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskProductComplex gets RelationshipRiskProductComplex, Tag 1597
func (m NoRelationshipRiskInstruments) GetRelationshipRiskProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskSecurityGroup gets RelationshipRiskSecurityGroup, Tag 1598
func (m NoRelationshipRiskInstruments) GetRelationshipRiskSecurityGroup() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskSecurityGroupField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskCFICode gets RelationshipRiskCFICode, Tag 1599
func (m NoRelationshipRiskInstruments) GetRelationshipRiskCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskSecurityType gets RelationshipRiskSecurityType, Tag 1600
func (m NoRelationshipRiskInstruments) GetRelationshipRiskSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskSecuritySubType gets RelationshipRiskSecuritySubType, Tag 1601
func (m NoRelationshipRiskInstruments) GetRelationshipRiskSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskMaturityMonthYear gets RelationshipRiskMaturityMonthYear, Tag 1602
func (m NoRelationshipRiskInstruments) GetRelationshipRiskMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskMaturityTime gets RelationshipRiskMaturityTime, Tag 1603
func (m NoRelationshipRiskInstruments) GetRelationshipRiskMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskRestructuringType gets RelationshipRiskRestructuringType, Tag 1604
func (m NoRelationshipRiskInstruments) GetRelationshipRiskRestructuringType() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskRestructuringTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskSeniority gets RelationshipRiskSeniority, Tag 1605
func (m NoRelationshipRiskInstruments) GetRelationshipRiskSeniority() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskSeniorityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskPutOrCall gets RelationshipRiskPutOrCall, Tag 1606
func (m NoRelationshipRiskInstruments) GetRelationshipRiskPutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskPutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskFlexibleIndicator gets RelationshipRiskFlexibleIndicator, Tag 1607
func (m NoRelationshipRiskInstruments) GetRelationshipRiskFlexibleIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskFlexibleIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskCouponRate gets RelationshipRiskCouponRate, Tag 1608
func (m NoRelationshipRiskInstruments) GetRelationshipRiskCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskSecurityExchange gets RelationshipRiskSecurityExchange, Tag 1609
func (m NoRelationshipRiskInstruments) GetRelationshipRiskSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskSecurityDesc gets RelationshipRiskSecurityDesc, Tag 1610
func (m NoRelationshipRiskInstruments) GetRelationshipRiskSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskEncodedSecurityDescLen gets RelationshipRiskEncodedSecurityDescLen, Tag 1618
func (m NoRelationshipRiskInstruments) GetRelationshipRiskEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskEncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskEncodedSecurityDesc gets RelationshipRiskEncodedSecurityDesc, Tag 1619
func (m NoRelationshipRiskInstruments) GetRelationshipRiskEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskEncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskInstrumentSettlType gets RelationshipRiskInstrumentSettlType, Tag 1611
func (m NoRelationshipRiskInstruments) GetRelationshipRiskInstrumentSettlType() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskInstrumentSettlTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskInstrumentMultiplier gets RelationshipRiskInstrumentMultiplier, Tag 1612
func (m NoRelationshipRiskInstruments) GetRelationshipRiskInstrumentMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskInstrumentMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasRelationshipRiskInstrumentOperator returns true if RelationshipRiskInstrumentOperator is present, Tag 1588
func (m NoRelationshipRiskInstruments) HasRelationshipRiskInstrumentOperator() bool {
	return m.Has(tag.RelationshipRiskInstrumentOperator)
}

//HasRelationshipRiskSymbol returns true if RelationshipRiskSymbol is present, Tag 1589
func (m NoRelationshipRiskInstruments) HasRelationshipRiskSymbol() bool {
	return m.Has(tag.RelationshipRiskSymbol)
}

//HasRelationshipRiskSymbolSfx returns true if RelationshipRiskSymbolSfx is present, Tag 1590
func (m NoRelationshipRiskInstruments) HasRelationshipRiskSymbolSfx() bool {
	return m.Has(tag.RelationshipRiskSymbolSfx)
}

//HasRelationshipRiskSecurityID returns true if RelationshipRiskSecurityID is present, Tag 1591
func (m NoRelationshipRiskInstruments) HasRelationshipRiskSecurityID() bool {
	return m.Has(tag.RelationshipRiskSecurityID)
}

//HasRelationshipRiskSecurityIDSource returns true if RelationshipRiskSecurityIDSource is present, Tag 1592
func (m NoRelationshipRiskInstruments) HasRelationshipRiskSecurityIDSource() bool {
	return m.Has(tag.RelationshipRiskSecurityIDSource)
}

//HasNoRelationshipRiskSecurityAltID returns true if NoRelationshipRiskSecurityAltID is present, Tag 1593
func (m NoRelationshipRiskInstruments) HasNoRelationshipRiskSecurityAltID() bool {
	return m.Has(tag.NoRelationshipRiskSecurityAltID)
}

//HasRelationshipRiskProduct returns true if RelationshipRiskProduct is present, Tag 1596
func (m NoRelationshipRiskInstruments) HasRelationshipRiskProduct() bool {
	return m.Has(tag.RelationshipRiskProduct)
}

//HasRelationshipRiskProductComplex returns true if RelationshipRiskProductComplex is present, Tag 1597
func (m NoRelationshipRiskInstruments) HasRelationshipRiskProductComplex() bool {
	return m.Has(tag.RelationshipRiskProductComplex)
}

//HasRelationshipRiskSecurityGroup returns true if RelationshipRiskSecurityGroup is present, Tag 1598
func (m NoRelationshipRiskInstruments) HasRelationshipRiskSecurityGroup() bool {
	return m.Has(tag.RelationshipRiskSecurityGroup)
}

//HasRelationshipRiskCFICode returns true if RelationshipRiskCFICode is present, Tag 1599
func (m NoRelationshipRiskInstruments) HasRelationshipRiskCFICode() bool {
	return m.Has(tag.RelationshipRiskCFICode)
}

//HasRelationshipRiskSecurityType returns true if RelationshipRiskSecurityType is present, Tag 1600
func (m NoRelationshipRiskInstruments) HasRelationshipRiskSecurityType() bool {
	return m.Has(tag.RelationshipRiskSecurityType)
}

//HasRelationshipRiskSecuritySubType returns true if RelationshipRiskSecuritySubType is present, Tag 1601
func (m NoRelationshipRiskInstruments) HasRelationshipRiskSecuritySubType() bool {
	return m.Has(tag.RelationshipRiskSecuritySubType)
}

//HasRelationshipRiskMaturityMonthYear returns true if RelationshipRiskMaturityMonthYear is present, Tag 1602
func (m NoRelationshipRiskInstruments) HasRelationshipRiskMaturityMonthYear() bool {
	return m.Has(tag.RelationshipRiskMaturityMonthYear)
}

//HasRelationshipRiskMaturityTime returns true if RelationshipRiskMaturityTime is present, Tag 1603
func (m NoRelationshipRiskInstruments) HasRelationshipRiskMaturityTime() bool {
	return m.Has(tag.RelationshipRiskMaturityTime)
}

//HasRelationshipRiskRestructuringType returns true if RelationshipRiskRestructuringType is present, Tag 1604
func (m NoRelationshipRiskInstruments) HasRelationshipRiskRestructuringType() bool {
	return m.Has(tag.RelationshipRiskRestructuringType)
}

//HasRelationshipRiskSeniority returns true if RelationshipRiskSeniority is present, Tag 1605
func (m NoRelationshipRiskInstruments) HasRelationshipRiskSeniority() bool {
	return m.Has(tag.RelationshipRiskSeniority)
}

//HasRelationshipRiskPutOrCall returns true if RelationshipRiskPutOrCall is present, Tag 1606
func (m NoRelationshipRiskInstruments) HasRelationshipRiskPutOrCall() bool {
	return m.Has(tag.RelationshipRiskPutOrCall)
}

//HasRelationshipRiskFlexibleIndicator returns true if RelationshipRiskFlexibleIndicator is present, Tag 1607
func (m NoRelationshipRiskInstruments) HasRelationshipRiskFlexibleIndicator() bool {
	return m.Has(tag.RelationshipRiskFlexibleIndicator)
}

//HasRelationshipRiskCouponRate returns true if RelationshipRiskCouponRate is present, Tag 1608
func (m NoRelationshipRiskInstruments) HasRelationshipRiskCouponRate() bool {
	return m.Has(tag.RelationshipRiskCouponRate)
}

//HasRelationshipRiskSecurityExchange returns true if RelationshipRiskSecurityExchange is present, Tag 1609
func (m NoRelationshipRiskInstruments) HasRelationshipRiskSecurityExchange() bool {
	return m.Has(tag.RelationshipRiskSecurityExchange)
}

//HasRelationshipRiskSecurityDesc returns true if RelationshipRiskSecurityDesc is present, Tag 1610
func (m NoRelationshipRiskInstruments) HasRelationshipRiskSecurityDesc() bool {
	return m.Has(tag.RelationshipRiskSecurityDesc)
}

//HasRelationshipRiskEncodedSecurityDescLen returns true if RelationshipRiskEncodedSecurityDescLen is present, Tag 1618
func (m NoRelationshipRiskInstruments) HasRelationshipRiskEncodedSecurityDescLen() bool {
	return m.Has(tag.RelationshipRiskEncodedSecurityDescLen)
}

//HasRelationshipRiskEncodedSecurityDesc returns true if RelationshipRiskEncodedSecurityDesc is present, Tag 1619
func (m NoRelationshipRiskInstruments) HasRelationshipRiskEncodedSecurityDesc() bool {
	return m.Has(tag.RelationshipRiskEncodedSecurityDesc)
}

//HasRelationshipRiskInstrumentSettlType returns true if RelationshipRiskInstrumentSettlType is present, Tag 1611
func (m NoRelationshipRiskInstruments) HasRelationshipRiskInstrumentSettlType() bool {
	return m.Has(tag.RelationshipRiskInstrumentSettlType)
}

//HasRelationshipRiskInstrumentMultiplier returns true if RelationshipRiskInstrumentMultiplier is present, Tag 1612
func (m NoRelationshipRiskInstruments) HasRelationshipRiskInstrumentMultiplier() bool {
	return m.Has(tag.RelationshipRiskInstrumentMultiplier)
}

//NoRelationshipRiskSecurityAltID is a repeating group element, Tag 1593
type NoRelationshipRiskSecurityAltID struct {
	*quickfix.Group
}

//SetRelationshipRiskSecurityAltID sets RelationshipRiskSecurityAltID, Tag 1594
func (m NoRelationshipRiskSecurityAltID) SetRelationshipRiskSecurityAltID(v string) {
	m.Set(field.NewRelationshipRiskSecurityAltID(v))
}

//SetRelationshipRiskSecurityAltIDSource sets RelationshipRiskSecurityAltIDSource, Tag 1595
func (m NoRelationshipRiskSecurityAltID) SetRelationshipRiskSecurityAltIDSource(v string) {
	m.Set(field.NewRelationshipRiskSecurityAltIDSource(v))
}

//GetRelationshipRiskSecurityAltID gets RelationshipRiskSecurityAltID, Tag 1594
func (m NoRelationshipRiskSecurityAltID) GetRelationshipRiskSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskSecurityAltIDSource gets RelationshipRiskSecurityAltIDSource, Tag 1595
func (m NoRelationshipRiskSecurityAltID) GetRelationshipRiskSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasRelationshipRiskSecurityAltID returns true if RelationshipRiskSecurityAltID is present, Tag 1594
func (m NoRelationshipRiskSecurityAltID) HasRelationshipRiskSecurityAltID() bool {
	return m.Has(tag.RelationshipRiskSecurityAltID)
}

//HasRelationshipRiskSecurityAltIDSource returns true if RelationshipRiskSecurityAltIDSource is present, Tag 1595
func (m NoRelationshipRiskSecurityAltID) HasRelationshipRiskSecurityAltIDSource() bool {
	return m.Has(tag.RelationshipRiskSecurityAltIDSource)
}

//NoRelationshipRiskSecurityAltIDRepeatingGroup is a repeating group, Tag 1593
type NoRelationshipRiskSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelationshipRiskSecurityAltIDRepeatingGroup returns an initialized, NoRelationshipRiskSecurityAltIDRepeatingGroup
func NewNoRelationshipRiskSecurityAltIDRepeatingGroup() NoRelationshipRiskSecurityAltIDRepeatingGroup {
	return NoRelationshipRiskSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelationshipRiskSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RelationshipRiskSecurityAltID), quickfix.GroupElement(tag.RelationshipRiskSecurityAltIDSource)})}
}

//Add create and append a new NoRelationshipRiskSecurityAltID to this group
func (m NoRelationshipRiskSecurityAltIDRepeatingGroup) Add() NoRelationshipRiskSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoRelationshipRiskSecurityAltID{g}
}

//Get returns the ith NoRelationshipRiskSecurityAltID in the NoRelationshipRiskSecurityAltIDRepeatinGroup
func (m NoRelationshipRiskSecurityAltIDRepeatingGroup) Get(i int) NoRelationshipRiskSecurityAltID {
	return NoRelationshipRiskSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoRelationshipRiskInstrumentsRepeatingGroup is a repeating group, Tag 1587
type NoRelationshipRiskInstrumentsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelationshipRiskInstrumentsRepeatingGroup returns an initialized, NoRelationshipRiskInstrumentsRepeatingGroup
func NewNoRelationshipRiskInstrumentsRepeatingGroup() NoRelationshipRiskInstrumentsRepeatingGroup {
	return NoRelationshipRiskInstrumentsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelationshipRiskInstruments,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RelationshipRiskInstrumentOperator), quickfix.GroupElement(tag.RelationshipRiskSymbol), quickfix.GroupElement(tag.RelationshipRiskSymbolSfx), quickfix.GroupElement(tag.RelationshipRiskSecurityID), quickfix.GroupElement(tag.RelationshipRiskSecurityIDSource), NewNoRelationshipRiskSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.RelationshipRiskProduct), quickfix.GroupElement(tag.RelationshipRiskProductComplex), quickfix.GroupElement(tag.RelationshipRiskSecurityGroup), quickfix.GroupElement(tag.RelationshipRiskCFICode), quickfix.GroupElement(tag.RelationshipRiskSecurityType), quickfix.GroupElement(tag.RelationshipRiskSecuritySubType), quickfix.GroupElement(tag.RelationshipRiskMaturityMonthYear), quickfix.GroupElement(tag.RelationshipRiskMaturityTime), quickfix.GroupElement(tag.RelationshipRiskRestructuringType), quickfix.GroupElement(tag.RelationshipRiskSeniority), quickfix.GroupElement(tag.RelationshipRiskPutOrCall), quickfix.GroupElement(tag.RelationshipRiskFlexibleIndicator), quickfix.GroupElement(tag.RelationshipRiskCouponRate), quickfix.GroupElement(tag.RelationshipRiskSecurityExchange), quickfix.GroupElement(tag.RelationshipRiskSecurityDesc), quickfix.GroupElement(tag.RelationshipRiskEncodedSecurityDescLen), quickfix.GroupElement(tag.RelationshipRiskEncodedSecurityDesc), quickfix.GroupElement(tag.RelationshipRiskInstrumentSettlType), quickfix.GroupElement(tag.RelationshipRiskInstrumentMultiplier)})}
}

//Add create and append a new NoRelationshipRiskInstruments to this group
func (m NoRelationshipRiskInstrumentsRepeatingGroup) Add() NoRelationshipRiskInstruments {
	g := m.RepeatingGroup.Add()
	return NoRelationshipRiskInstruments{g}
}

//Get returns the ith NoRelationshipRiskInstruments in the NoRelationshipRiskInstrumentsRepeatinGroup
func (m NoRelationshipRiskInstrumentsRepeatingGroup) Get(i int) NoRelationshipRiskInstruments {
	return NoRelationshipRiskInstruments{m.RepeatingGroup.Get(i)}
}

//NoRelationshipRiskWarningLevels is a repeating group element, Tag 1613
type NoRelationshipRiskWarningLevels struct {
	*quickfix.Group
}

//SetRelationshipRiskWarningLevelPercent sets RelationshipRiskWarningLevelPercent, Tag 1614
func (m NoRelationshipRiskWarningLevels) SetRelationshipRiskWarningLevelPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewRelationshipRiskWarningLevelPercent(value, scale))
}

//SetRelationshipRiskWarningLevelName sets RelationshipRiskWarningLevelName, Tag 1615
func (m NoRelationshipRiskWarningLevels) SetRelationshipRiskWarningLevelName(v string) {
	m.Set(field.NewRelationshipRiskWarningLevelName(v))
}

//GetRelationshipRiskWarningLevelPercent gets RelationshipRiskWarningLevelPercent, Tag 1614
func (m NoRelationshipRiskWarningLevels) GetRelationshipRiskWarningLevelPercent() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskWarningLevelPercentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelationshipRiskWarningLevelName gets RelationshipRiskWarningLevelName, Tag 1615
func (m NoRelationshipRiskWarningLevels) GetRelationshipRiskWarningLevelName() (v string, err quickfix.MessageRejectError) {
	var f field.RelationshipRiskWarningLevelNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasRelationshipRiskWarningLevelPercent returns true if RelationshipRiskWarningLevelPercent is present, Tag 1614
func (m NoRelationshipRiskWarningLevels) HasRelationshipRiskWarningLevelPercent() bool {
	return m.Has(tag.RelationshipRiskWarningLevelPercent)
}

//HasRelationshipRiskWarningLevelName returns true if RelationshipRiskWarningLevelName is present, Tag 1615
func (m NoRelationshipRiskWarningLevels) HasRelationshipRiskWarningLevelName() bool {
	return m.Has(tag.RelationshipRiskWarningLevelName)
}

//NoRelationshipRiskWarningLevelsRepeatingGroup is a repeating group, Tag 1613
type NoRelationshipRiskWarningLevelsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelationshipRiskWarningLevelsRepeatingGroup returns an initialized, NoRelationshipRiskWarningLevelsRepeatingGroup
func NewNoRelationshipRiskWarningLevelsRepeatingGroup() NoRelationshipRiskWarningLevelsRepeatingGroup {
	return NoRelationshipRiskWarningLevelsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelationshipRiskWarningLevels,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RelationshipRiskWarningLevelPercent), quickfix.GroupElement(tag.RelationshipRiskWarningLevelName)})}
}

//Add create and append a new NoRelationshipRiskWarningLevels to this group
func (m NoRelationshipRiskWarningLevelsRepeatingGroup) Add() NoRelationshipRiskWarningLevels {
	g := m.RepeatingGroup.Add()
	return NoRelationshipRiskWarningLevels{g}
}

//Get returns the ith NoRelationshipRiskWarningLevels in the NoRelationshipRiskWarningLevelsRepeatinGroup
func (m NoRelationshipRiskWarningLevelsRepeatingGroup) Get(i int) NoRelationshipRiskWarningLevels {
	return NoRelationshipRiskWarningLevels{m.RepeatingGroup.Get(i)}
}

//NoRelationshipRiskLimitsRepeatingGroup is a repeating group, Tag 1582
type NoRelationshipRiskLimitsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelationshipRiskLimitsRepeatingGroup returns an initialized, NoRelationshipRiskLimitsRepeatingGroup
func NewNoRelationshipRiskLimitsRepeatingGroup() NoRelationshipRiskLimitsRepeatingGroup {
	return NoRelationshipRiskLimitsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelationshipRiskLimits,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RelationshipRiskLimitType), quickfix.GroupElement(tag.RelationshipRiskLimitAmount), quickfix.GroupElement(tag.RelationshipRiskLimitCurrency), quickfix.GroupElement(tag.RelationshipRiskLimitPlatform), NewNoRelationshipRiskInstrumentsRepeatingGroup(), NewNoRelationshipRiskWarningLevelsRepeatingGroup()})}
}

//Add create and append a new NoRelationshipRiskLimits to this group
func (m NoRelationshipRiskLimitsRepeatingGroup) Add() NoRelationshipRiskLimits {
	g := m.RepeatingGroup.Add()
	return NoRelationshipRiskLimits{g}
}

//Get returns the ith NoRelationshipRiskLimits in the NoRelationshipRiskLimitsRepeatinGroup
func (m NoRelationshipRiskLimitsRepeatingGroup) Get(i int) NoRelationshipRiskLimits {
	return NoRelationshipRiskLimits{m.RepeatingGroup.Get(i)}
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

//NoRelatedPartyIDsRepeatingGroup is a repeating group, Tag 1562
type NoRelatedPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelatedPartyIDsRepeatingGroup returns an initialized, NoRelatedPartyIDsRepeatingGroup
func NewNoRelatedPartyIDsRepeatingGroup() NoRelatedPartyIDsRepeatingGroup {
	return NoRelatedPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RelatedPartyID), quickfix.GroupElement(tag.RelatedPartyIDSource), quickfix.GroupElement(tag.RelatedPartyRole), NewNoRelatedPartySubIDsRepeatingGroup(), NewNoRelatedPartyAltIDsRepeatingGroup(), NewNoRelatedContextPartyIDsRepeatingGroup(), NewNoRelationshipRiskLimitsRepeatingGroup(), NewNoPartyRelationshipsRepeatingGroup()})}
}

//Add create and append a new NoRelatedPartyIDs to this group
func (m NoRelatedPartyIDsRepeatingGroup) Add() NoRelatedPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoRelatedPartyIDs{g}
}

//Get returns the ith NoRelatedPartyIDs in the NoRelatedPartyIDsRepeatinGroup
func (m NoRelatedPartyIDsRepeatingGroup) Get(i int) NoRelatedPartyIDs {
	return NoRelatedPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoPartyListRepeatingGroup is a repeating group, Tag 1513
type NoPartyListRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyListRepeatingGroup returns an initialized, NoPartyListRepeatingGroup
func NewNoPartyListRepeatingGroup() NoPartyListRepeatingGroup {
	return NoPartyListRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyList,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), NewNoPartySubIDsRepeatingGroup(), NewNoPartyAltIDsRepeatingGroup(), NewNoContextPartyIDsRepeatingGroup(), NewNoRiskLimitsRepeatingGroup(), NewNoRelatedPartyIDsRepeatingGroup()})}
}

//Add create and append a new NoPartyList to this group
func (m NoPartyListRepeatingGroup) Add() NoPartyList {
	g := m.RepeatingGroup.Add()
	return NoPartyList{g}
}

//Get returns the ith NoPartyList in the NoPartyListRepeatinGroup
func (m NoPartyListRepeatingGroup) Get(i int) NoPartyList {
	return NoPartyList{m.RepeatingGroup.Get(i)}
}
