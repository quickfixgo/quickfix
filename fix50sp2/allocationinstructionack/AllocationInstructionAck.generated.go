package allocationinstructionack

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//AllocationInstructionAck is the fix50sp2 AllocationInstructionAck type, MsgType = P
type AllocationInstructionAck struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a AllocationInstructionAck from a quickfix.Message instance
func FromMessage(m quickfix.Message) AllocationInstructionAck {
	return AllocationInstructionAck{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m AllocationInstructionAck) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a AllocationInstructionAck initialized with the required fields for AllocationInstructionAck
func New(allocid field.AllocIDField, allocstatus field.AllocStatusField) (m AllocationInstructionAck) {
	m.Header = fixt11.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("P"))
	m.Set(allocid)
	m.Set(allocstatus)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg AllocationInstructionAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "P", r
}

//SetText sets Text, Tag 58
func (m AllocationInstructionAck) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m AllocationInstructionAck) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetAllocID sets AllocID, Tag 70
func (m AllocationInstructionAck) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m AllocationInstructionAck) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetNoAllocs sets NoAllocs, Tag 78
func (m AllocationInstructionAck) SetNoAllocs(f NoAllocsRepeatingGroup) {
	m.SetGroup(f)
}

//SetAllocStatus sets AllocStatus, Tag 87
func (m AllocationInstructionAck) SetAllocStatus(v int) {
	m.Set(field.NewAllocStatus(v))
}

//SetAllocRejCode sets AllocRejCode, Tag 88
func (m AllocationInstructionAck) SetAllocRejCode(v int) {
	m.Set(field.NewAllocRejCode(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m AllocationInstructionAck) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m AllocationInstructionAck) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m AllocationInstructionAck) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m AllocationInstructionAck) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m AllocationInstructionAck) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetMatchStatus sets MatchStatus, Tag 573
func (m AllocationInstructionAck) SetMatchStatus(v string) {
	m.Set(field.NewMatchStatus(v))
}

//SetAllocType sets AllocType, Tag 626
func (m AllocationInstructionAck) SetAllocType(v int) {
	m.Set(field.NewAllocType(v))
}

//SetSecondaryAllocID sets SecondaryAllocID, Tag 793
func (m AllocationInstructionAck) SetSecondaryAllocID(v string) {
	m.Set(field.NewSecondaryAllocID(v))
}

//SetAllocIntermedReqType sets AllocIntermedReqType, Tag 808
func (m AllocationInstructionAck) SetAllocIntermedReqType(v int) {
	m.Set(field.NewAllocIntermedReqType(v))
}

//GetText gets Text, Tag 58
func (m AllocationInstructionAck) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m AllocationInstructionAck) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocID gets AllocID, Tag 70
func (m AllocationInstructionAck) GetAllocID() (f field.AllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m AllocationInstructionAck) GetTradeDate() (f field.TradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoAllocs gets NoAllocs, Tag 78
func (m AllocationInstructionAck) GetNoAllocs() (NoAllocsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAllocsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAllocStatus gets AllocStatus, Tag 87
func (m AllocationInstructionAck) GetAllocStatus() (f field.AllocStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocRejCode gets AllocRejCode, Tag 88
func (m AllocationInstructionAck) GetAllocRejCode() (f field.AllocRejCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m AllocationInstructionAck) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m AllocationInstructionAck) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m AllocationInstructionAck) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m AllocationInstructionAck) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m AllocationInstructionAck) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMatchStatus gets MatchStatus, Tag 573
func (m AllocationInstructionAck) GetMatchStatus() (f field.MatchStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocType gets AllocType, Tag 626
func (m AllocationInstructionAck) GetAllocType() (f field.AllocTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryAllocID gets SecondaryAllocID, Tag 793
func (m AllocationInstructionAck) GetSecondaryAllocID() (f field.SecondaryAllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocIntermedReqType gets AllocIntermedReqType, Tag 808
func (m AllocationInstructionAck) GetAllocIntermedReqType() (f field.AllocIntermedReqTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasText returns true if Text is present, Tag 58
func (m AllocationInstructionAck) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m AllocationInstructionAck) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasAllocID returns true if AllocID is present, Tag 70
func (m AllocationInstructionAck) HasAllocID() bool {
	return m.Has(tag.AllocID)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m AllocationInstructionAck) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasNoAllocs returns true if NoAllocs is present, Tag 78
func (m AllocationInstructionAck) HasNoAllocs() bool {
	return m.Has(tag.NoAllocs)
}

//HasAllocStatus returns true if AllocStatus is present, Tag 87
func (m AllocationInstructionAck) HasAllocStatus() bool {
	return m.Has(tag.AllocStatus)
}

//HasAllocRejCode returns true if AllocRejCode is present, Tag 88
func (m AllocationInstructionAck) HasAllocRejCode() bool {
	return m.Has(tag.AllocRejCode)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m AllocationInstructionAck) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m AllocationInstructionAck) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m AllocationInstructionAck) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m AllocationInstructionAck) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasProduct returns true if Product is present, Tag 460
func (m AllocationInstructionAck) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasMatchStatus returns true if MatchStatus is present, Tag 573
func (m AllocationInstructionAck) HasMatchStatus() bool {
	return m.Has(tag.MatchStatus)
}

//HasAllocType returns true if AllocType is present, Tag 626
func (m AllocationInstructionAck) HasAllocType() bool {
	return m.Has(tag.AllocType)
}

//HasSecondaryAllocID returns true if SecondaryAllocID is present, Tag 793
func (m AllocationInstructionAck) HasSecondaryAllocID() bool {
	return m.Has(tag.SecondaryAllocID)
}

//HasAllocIntermedReqType returns true if AllocIntermedReqType is present, Tag 808
func (m AllocationInstructionAck) HasAllocIntermedReqType() bool {
	return m.Has(tag.AllocIntermedReqType)
}

//NoAllocs is a repeating group element, Tag 78
type NoAllocs struct {
	quickfix.Group
}

//SetAllocAccount sets AllocAccount, Tag 79
func (m NoAllocs) SetAllocAccount(v string) {
	m.Set(field.NewAllocAccount(v))
}

//SetAllocAcctIDSource sets AllocAcctIDSource, Tag 661
func (m NoAllocs) SetAllocAcctIDSource(v int) {
	m.Set(field.NewAllocAcctIDSource(v))
}

//SetAllocPrice sets AllocPrice, Tag 366
func (m NoAllocs) SetAllocPrice(v float64) {
	m.Set(field.NewAllocPrice(v))
}

//SetIndividualAllocID sets IndividualAllocID, Tag 467
func (m NoAllocs) SetIndividualAllocID(v string) {
	m.Set(field.NewIndividualAllocID(v))
}

//SetIndividualAllocRejCode sets IndividualAllocRejCode, Tag 776
func (m NoAllocs) SetIndividualAllocRejCode(v int) {
	m.Set(field.NewIndividualAllocRejCode(v))
}

//SetAllocText sets AllocText, Tag 161
func (m NoAllocs) SetAllocText(v string) {
	m.Set(field.NewAllocText(v))
}

//SetEncodedAllocTextLen sets EncodedAllocTextLen, Tag 360
func (m NoAllocs) SetEncodedAllocTextLen(v int) {
	m.Set(field.NewEncodedAllocTextLen(v))
}

//SetEncodedAllocText sets EncodedAllocText, Tag 361
func (m NoAllocs) SetEncodedAllocText(v string) {
	m.Set(field.NewEncodedAllocText(v))
}

//SetSecondaryIndividualAllocID sets SecondaryIndividualAllocID, Tag 989
func (m NoAllocs) SetSecondaryIndividualAllocID(v string) {
	m.Set(field.NewSecondaryIndividualAllocID(v))
}

//SetAllocCustomerCapacity sets AllocCustomerCapacity, Tag 993
func (m NoAllocs) SetAllocCustomerCapacity(v string) {
	m.Set(field.NewAllocCustomerCapacity(v))
}

//SetIndividualAllocType sets IndividualAllocType, Tag 992
func (m NoAllocs) SetIndividualAllocType(v int) {
	m.Set(field.NewIndividualAllocType(v))
}

//SetAllocQty sets AllocQty, Tag 80
func (m NoAllocs) SetAllocQty(v float64) {
	m.Set(field.NewAllocQty(v))
}

//SetNoNestedPartyIDs sets NoNestedPartyIDs, Tag 539
func (m NoAllocs) SetNoNestedPartyIDs(f NoNestedPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetAllocPositionEffect sets AllocPositionEffect, Tag 1047
func (m NoAllocs) SetAllocPositionEffect(v string) {
	m.Set(field.NewAllocPositionEffect(v))
}

//GetAllocAccount gets AllocAccount, Tag 79
func (m NoAllocs) GetAllocAccount() (f field.AllocAccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocAcctIDSource gets AllocAcctIDSource, Tag 661
func (m NoAllocs) GetAllocAcctIDSource() (f field.AllocAcctIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocPrice gets AllocPrice, Tag 366
func (m NoAllocs) GetAllocPrice() (f field.AllocPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIndividualAllocID gets IndividualAllocID, Tag 467
func (m NoAllocs) GetIndividualAllocID() (f field.IndividualAllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIndividualAllocRejCode gets IndividualAllocRejCode, Tag 776
func (m NoAllocs) GetIndividualAllocRejCode() (f field.IndividualAllocRejCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocText gets AllocText, Tag 161
func (m NoAllocs) GetAllocText() (f field.AllocTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedAllocTextLen gets EncodedAllocTextLen, Tag 360
func (m NoAllocs) GetEncodedAllocTextLen() (f field.EncodedAllocTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedAllocText gets EncodedAllocText, Tag 361
func (m NoAllocs) GetEncodedAllocText() (f field.EncodedAllocTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryIndividualAllocID gets SecondaryIndividualAllocID, Tag 989
func (m NoAllocs) GetSecondaryIndividualAllocID() (f field.SecondaryIndividualAllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocCustomerCapacity gets AllocCustomerCapacity, Tag 993
func (m NoAllocs) GetAllocCustomerCapacity() (f field.AllocCustomerCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIndividualAllocType gets IndividualAllocType, Tag 992
func (m NoAllocs) GetIndividualAllocType() (f field.IndividualAllocTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocQty gets AllocQty, Tag 80
func (m NoAllocs) GetAllocQty() (f field.AllocQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoNestedPartyIDs gets NoNestedPartyIDs, Tag 539
func (m NoAllocs) GetNoNestedPartyIDs() (NoNestedPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAllocPositionEffect gets AllocPositionEffect, Tag 1047
func (m NoAllocs) GetAllocPositionEffect() (f field.AllocPositionEffectField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasAllocAccount returns true if AllocAccount is present, Tag 79
func (m NoAllocs) HasAllocAccount() bool {
	return m.Has(tag.AllocAccount)
}

//HasAllocAcctIDSource returns true if AllocAcctIDSource is present, Tag 661
func (m NoAllocs) HasAllocAcctIDSource() bool {
	return m.Has(tag.AllocAcctIDSource)
}

//HasAllocPrice returns true if AllocPrice is present, Tag 366
func (m NoAllocs) HasAllocPrice() bool {
	return m.Has(tag.AllocPrice)
}

//HasIndividualAllocID returns true if IndividualAllocID is present, Tag 467
func (m NoAllocs) HasIndividualAllocID() bool {
	return m.Has(tag.IndividualAllocID)
}

//HasIndividualAllocRejCode returns true if IndividualAllocRejCode is present, Tag 776
func (m NoAllocs) HasIndividualAllocRejCode() bool {
	return m.Has(tag.IndividualAllocRejCode)
}

//HasAllocText returns true if AllocText is present, Tag 161
func (m NoAllocs) HasAllocText() bool {
	return m.Has(tag.AllocText)
}

//HasEncodedAllocTextLen returns true if EncodedAllocTextLen is present, Tag 360
func (m NoAllocs) HasEncodedAllocTextLen() bool {
	return m.Has(tag.EncodedAllocTextLen)
}

//HasEncodedAllocText returns true if EncodedAllocText is present, Tag 361
func (m NoAllocs) HasEncodedAllocText() bool {
	return m.Has(tag.EncodedAllocText)
}

//HasSecondaryIndividualAllocID returns true if SecondaryIndividualAllocID is present, Tag 989
func (m NoAllocs) HasSecondaryIndividualAllocID() bool {
	return m.Has(tag.SecondaryIndividualAllocID)
}

//HasAllocCustomerCapacity returns true if AllocCustomerCapacity is present, Tag 993
func (m NoAllocs) HasAllocCustomerCapacity() bool {
	return m.Has(tag.AllocCustomerCapacity)
}

//HasIndividualAllocType returns true if IndividualAllocType is present, Tag 992
func (m NoAllocs) HasIndividualAllocType() bool {
	return m.Has(tag.IndividualAllocType)
}

//HasAllocQty returns true if AllocQty is present, Tag 80
func (m NoAllocs) HasAllocQty() bool {
	return m.Has(tag.AllocQty)
}

//HasNoNestedPartyIDs returns true if NoNestedPartyIDs is present, Tag 539
func (m NoAllocs) HasNoNestedPartyIDs() bool {
	return m.Has(tag.NoNestedPartyIDs)
}

//HasAllocPositionEffect returns true if AllocPositionEffect is present, Tag 1047
func (m NoAllocs) HasAllocPositionEffect() bool {
	return m.Has(tag.AllocPositionEffect)
}

//NoNestedPartyIDs is a repeating group element, Tag 539
type NoNestedPartyIDs struct {
	quickfix.Group
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
func (m NoNestedPartyIDs) GetNestedPartyID() (f field.NestedPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNestedPartyIDSource gets NestedPartyIDSource, Tag 525
func (m NoNestedPartyIDs) GetNestedPartyIDSource() (f field.NestedPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNestedPartyRole gets NestedPartyRole, Tag 538
func (m NoNestedPartyIDs) GetNestedPartyRole() (f field.NestedPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
	quickfix.Group
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
func (m NoNestedPartySubIDs) GetNestedPartySubID() (f field.NestedPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNestedPartySubIDType gets NestedPartySubIDType, Tag 805
func (m NoNestedPartySubIDs) GetNestedPartySubIDType() (f field.NestedPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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

//NoAllocsRepeatingGroup is a repeating group, Tag 78
type NoAllocsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoAllocsRepeatingGroup returns an initialized, NoAllocsRepeatingGroup
func NewNoAllocsRepeatingGroup() NoAllocsRepeatingGroup {
	return NoAllocsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAllocs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AllocAccount), quickfix.GroupElement(tag.AllocAcctIDSource), quickfix.GroupElement(tag.AllocPrice), quickfix.GroupElement(tag.IndividualAllocID), quickfix.GroupElement(tag.IndividualAllocRejCode), quickfix.GroupElement(tag.AllocText), quickfix.GroupElement(tag.EncodedAllocTextLen), quickfix.GroupElement(tag.EncodedAllocText), quickfix.GroupElement(tag.SecondaryIndividualAllocID), quickfix.GroupElement(tag.AllocCustomerCapacity), quickfix.GroupElement(tag.IndividualAllocType), quickfix.GroupElement(tag.AllocQty), NewNoNestedPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.AllocPositionEffect)})}
}

//Add create and append a new NoAllocs to this group
func (m NoAllocsRepeatingGroup) Add() NoAllocs {
	g := m.RepeatingGroup.Add()
	return NoAllocs{g}
}

//Get returns the ith NoAllocs in the NoAllocsRepeatinGroup
func (m NoAllocsRepeatingGroup) Get(i int) NoAllocs {
	return NoAllocs{m.RepeatingGroup.Get(i)}
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
