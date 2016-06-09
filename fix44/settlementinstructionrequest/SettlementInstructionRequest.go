package settlementinstructionrequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/tag"
)

//SettlementInstructionRequest is the fix44 SettlementInstructionRequest type, MsgType = AV
type SettlementInstructionRequest struct {
	fix44.Header
	quickfix.Body
	fix44.Trailer
}

//FromMessage creates a SettlementInstructionRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) SettlementInstructionRequest {
	return SettlementInstructionRequest{
		Header:  fix44.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fix44.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m SettlementInstructionRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a SettlementInstructionRequest initialized with the required fields for SettlementInstructionRequest
func New(settlinstreqid field.SettlInstReqIDField, transacttime field.TransactTimeField) (m SettlementInstructionRequest) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("AV"))
	m.Set(settlinstreqid)
	m.Set(transacttime)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg SettlementInstructionRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "AV", r
}

//SetSide sets Side, Tag 54
func (m SettlementInstructionRequest) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m SettlementInstructionRequest) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetAllocAccount sets AllocAccount, Tag 79
func (m SettlementInstructionRequest) SetAllocAccount(v string) {
	m.Set(field.NewAllocAccount(v))
}

//SetExpireTime sets ExpireTime, Tag 126
func (m SettlementInstructionRequest) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m SettlementInstructionRequest) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetEffectiveTime sets EffectiveTime, Tag 168
func (m SettlementInstructionRequest) SetEffectiveTime(v time.Time) {
	m.Set(field.NewEffectiveTime(v))
}

//SetStandInstDbType sets StandInstDbType, Tag 169
func (m SettlementInstructionRequest) SetStandInstDbType(v int) {
	m.Set(field.NewStandInstDbType(v))
}

//SetStandInstDbName sets StandInstDbName, Tag 170
func (m SettlementInstructionRequest) SetStandInstDbName(v string) {
	m.Set(field.NewStandInstDbName(v))
}

//SetStandInstDbID sets StandInstDbID, Tag 171
func (m SettlementInstructionRequest) SetStandInstDbID(v string) {
	m.Set(field.NewStandInstDbID(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m SettlementInstructionRequest) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m SettlementInstructionRequest) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m SettlementInstructionRequest) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetAllocAcctIDSource sets AllocAcctIDSource, Tag 661
func (m SettlementInstructionRequest) SetAllocAcctIDSource(v int) {
	m.Set(field.NewAllocAcctIDSource(v))
}

//SetLastUpdateTime sets LastUpdateTime, Tag 779
func (m SettlementInstructionRequest) SetLastUpdateTime(v time.Time) {
	m.Set(field.NewLastUpdateTime(v))
}

//SetSettlInstReqID sets SettlInstReqID, Tag 791
func (m SettlementInstructionRequest) SetSettlInstReqID(v string) {
	m.Set(field.NewSettlInstReqID(v))
}

//GetSide gets Side, Tag 54
func (m SettlementInstructionRequest) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m SettlementInstructionRequest) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocAccount gets AllocAccount, Tag 79
func (m SettlementInstructionRequest) GetAllocAccount() (f field.AllocAccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m SettlementInstructionRequest) GetExpireTime() (f field.ExpireTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m SettlementInstructionRequest) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEffectiveTime gets EffectiveTime, Tag 168
func (m SettlementInstructionRequest) GetEffectiveTime() (f field.EffectiveTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStandInstDbType gets StandInstDbType, Tag 169
func (m SettlementInstructionRequest) GetStandInstDbType() (f field.StandInstDbTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStandInstDbName gets StandInstDbName, Tag 170
func (m SettlementInstructionRequest) GetStandInstDbName() (f field.StandInstDbNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStandInstDbID gets StandInstDbID, Tag 171
func (m SettlementInstructionRequest) GetStandInstDbID() (f field.StandInstDbIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m SettlementInstructionRequest) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m SettlementInstructionRequest) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m SettlementInstructionRequest) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocAcctIDSource gets AllocAcctIDSource, Tag 661
func (m SettlementInstructionRequest) GetAllocAcctIDSource() (f field.AllocAcctIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastUpdateTime gets LastUpdateTime, Tag 779
func (m SettlementInstructionRequest) GetLastUpdateTime() (f field.LastUpdateTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlInstReqID gets SettlInstReqID, Tag 791
func (m SettlementInstructionRequest) GetSettlInstReqID() (f field.SettlInstReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasSide returns true if Side is present, Tag 54
func (m SettlementInstructionRequest) HasSide() bool {
	return m.Has(tag.Side)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m SettlementInstructionRequest) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasAllocAccount returns true if AllocAccount is present, Tag 79
func (m SettlementInstructionRequest) HasAllocAccount() bool {
	return m.Has(tag.AllocAccount)
}

//HasExpireTime returns true if ExpireTime is present, Tag 126
func (m SettlementInstructionRequest) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m SettlementInstructionRequest) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasEffectiveTime returns true if EffectiveTime is present, Tag 168
func (m SettlementInstructionRequest) HasEffectiveTime() bool {
	return m.Has(tag.EffectiveTime)
}

//HasStandInstDbType returns true if StandInstDbType is present, Tag 169
func (m SettlementInstructionRequest) HasStandInstDbType() bool {
	return m.Has(tag.StandInstDbType)
}

//HasStandInstDbName returns true if StandInstDbName is present, Tag 170
func (m SettlementInstructionRequest) HasStandInstDbName() bool {
	return m.Has(tag.StandInstDbName)
}

//HasStandInstDbID returns true if StandInstDbID is present, Tag 171
func (m SettlementInstructionRequest) HasStandInstDbID() bool {
	return m.Has(tag.StandInstDbID)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m SettlementInstructionRequest) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasProduct returns true if Product is present, Tag 460
func (m SettlementInstructionRequest) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m SettlementInstructionRequest) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasAllocAcctIDSource returns true if AllocAcctIDSource is present, Tag 661
func (m SettlementInstructionRequest) HasAllocAcctIDSource() bool {
	return m.Has(tag.AllocAcctIDSource)
}

//HasLastUpdateTime returns true if LastUpdateTime is present, Tag 779
func (m SettlementInstructionRequest) HasLastUpdateTime() bool {
	return m.Has(tag.LastUpdateTime)
}

//HasSettlInstReqID returns true if SettlInstReqID is present, Tag 791
func (m SettlementInstructionRequest) HasSettlInstReqID() bool {
	return m.Has(tag.SettlInstReqID)
}

//NoPartyIDs is a repeating group element, Tag 453
type NoPartyIDs struct {
	quickfix.Group
}

//SetPartyID sets PartyID, Tag 448
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

//SetPartyIDSource sets PartyIDSource, Tag 447
func (m NoPartyIDs) SetPartyIDSource(v string) {
	m.Set(field.NewPartyIDSource(v))
}

//SetPartyRole sets PartyRole, Tag 452
func (m NoPartyIDs) SetPartyRole(v int) {
	m.Set(field.NewPartyRole(v))
}

//SetNoPartySubIDs sets NoPartySubIDs, Tag 802
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetPartyID gets PartyID, Tag 448
func (m NoPartyIDs) GetPartyID() (f field.PartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartyIDSource gets PartyIDSource, Tag 447
func (m NoPartyIDs) GetPartyIDSource() (f field.PartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartyRole gets PartyRole, Tag 452
func (m NoPartyIDs) GetPartyRole() (f field.PartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
	quickfix.Group
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//SetPartySubIDType sets PartySubIDType, Tag 803
func (m NoPartySubIDs) SetPartySubIDType(v int) {
	m.Set(field.NewPartySubIDType(v))
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartySubIDs) GetPartySubID() (f field.PartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartySubIDType gets PartySubIDType, Tag 803
func (m NoPartySubIDs) GetPartySubIDType() (f field.PartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), quickfix.GroupElement(tag.NoPartySubIDs)})}
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
