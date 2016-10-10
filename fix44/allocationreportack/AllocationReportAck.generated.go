package allocationreportack

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/tag"
)

//AllocationReportAck is the fix44 AllocationReportAck type, MsgType = AT
type AllocationReportAck struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

//FromMessage creates a AllocationReportAck from a quickfix.Message instance
func FromMessage(m *quickfix.Message) AllocationReportAck {
	return AllocationReportAck{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m AllocationReportAck) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a AllocationReportAck initialized with the required fields for AllocationReportAck
func New(allocreportid field.AllocReportIDField, allocid field.AllocIDField, transacttime field.TransactTimeField, allocstatus field.AllocStatusField) (m AllocationReportAck) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("AT"))
	m.Set(allocreportid)
	m.Set(allocid)
	m.Set(transacttime)
	m.Set(allocstatus)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg AllocationReportAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "AT", r
}

//SetText sets Text, Tag 58
func (m AllocationReportAck) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m AllocationReportAck) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetAllocID sets AllocID, Tag 70
func (m AllocationReportAck) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m AllocationReportAck) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetNoAllocs sets NoAllocs, Tag 78
func (m AllocationReportAck) SetNoAllocs(f NoAllocsRepeatingGroup) {
	m.SetGroup(f)
}

//SetAllocStatus sets AllocStatus, Tag 87
func (m AllocationReportAck) SetAllocStatus(v enum.AllocStatus) {
	m.Set(field.NewAllocStatus(v))
}

//SetAllocRejCode sets AllocRejCode, Tag 88
func (m AllocationReportAck) SetAllocRejCode(v enum.AllocRejCode) {
	m.Set(field.NewAllocRejCode(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m AllocationReportAck) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m AllocationReportAck) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m AllocationReportAck) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m AllocationReportAck) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m AllocationReportAck) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

//SetMatchStatus sets MatchStatus, Tag 573
func (m AllocationReportAck) SetMatchStatus(v enum.MatchStatus) {
	m.Set(field.NewMatchStatus(v))
}

//SetAllocReportID sets AllocReportID, Tag 755
func (m AllocationReportAck) SetAllocReportID(v string) {
	m.Set(field.NewAllocReportID(v))
}

//SetSecondaryAllocID sets SecondaryAllocID, Tag 793
func (m AllocationReportAck) SetSecondaryAllocID(v string) {
	m.Set(field.NewSecondaryAllocID(v))
}

//SetAllocReportType sets AllocReportType, Tag 794
func (m AllocationReportAck) SetAllocReportType(v enum.AllocReportType) {
	m.Set(field.NewAllocReportType(v))
}

//SetAllocIntermedReqType sets AllocIntermedReqType, Tag 808
func (m AllocationReportAck) SetAllocIntermedReqType(v enum.AllocIntermedReqType) {
	m.Set(field.NewAllocIntermedReqType(v))
}

//GetText gets Text, Tag 58
func (m AllocationReportAck) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m AllocationReportAck) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocID gets AllocID, Tag 70
func (m AllocationReportAck) GetAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.AllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m AllocationReportAck) GetTradeDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoAllocs gets NoAllocs, Tag 78
func (m AllocationReportAck) GetNoAllocs() (NoAllocsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAllocsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAllocStatus gets AllocStatus, Tag 87
func (m AllocationReportAck) GetAllocStatus() (v enum.AllocStatus, err quickfix.MessageRejectError) {
	var f field.AllocStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocRejCode gets AllocRejCode, Tag 88
func (m AllocationReportAck) GetAllocRejCode() (v enum.AllocRejCode, err quickfix.MessageRejectError) {
	var f field.AllocRejCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m AllocationReportAck) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m AllocationReportAck) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m AllocationReportAck) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m AllocationReportAck) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m AllocationReportAck) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMatchStatus gets MatchStatus, Tag 573
func (m AllocationReportAck) GetMatchStatus() (v enum.MatchStatus, err quickfix.MessageRejectError) {
	var f field.MatchStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocReportID gets AllocReportID, Tag 755
func (m AllocationReportAck) GetAllocReportID() (v string, err quickfix.MessageRejectError) {
	var f field.AllocReportIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecondaryAllocID gets SecondaryAllocID, Tag 793
func (m AllocationReportAck) GetSecondaryAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryAllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocReportType gets AllocReportType, Tag 794
func (m AllocationReportAck) GetAllocReportType() (v enum.AllocReportType, err quickfix.MessageRejectError) {
	var f field.AllocReportTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocIntermedReqType gets AllocIntermedReqType, Tag 808
func (m AllocationReportAck) GetAllocIntermedReqType() (v enum.AllocIntermedReqType, err quickfix.MessageRejectError) {
	var f field.AllocIntermedReqTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasText returns true if Text is present, Tag 58
func (m AllocationReportAck) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m AllocationReportAck) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasAllocID returns true if AllocID is present, Tag 70
func (m AllocationReportAck) HasAllocID() bool {
	return m.Has(tag.AllocID)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m AllocationReportAck) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasNoAllocs returns true if NoAllocs is present, Tag 78
func (m AllocationReportAck) HasNoAllocs() bool {
	return m.Has(tag.NoAllocs)
}

//HasAllocStatus returns true if AllocStatus is present, Tag 87
func (m AllocationReportAck) HasAllocStatus() bool {
	return m.Has(tag.AllocStatus)
}

//HasAllocRejCode returns true if AllocRejCode is present, Tag 88
func (m AllocationReportAck) HasAllocRejCode() bool {
	return m.Has(tag.AllocRejCode)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m AllocationReportAck) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m AllocationReportAck) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m AllocationReportAck) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m AllocationReportAck) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasProduct returns true if Product is present, Tag 460
func (m AllocationReportAck) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasMatchStatus returns true if MatchStatus is present, Tag 573
func (m AllocationReportAck) HasMatchStatus() bool {
	return m.Has(tag.MatchStatus)
}

//HasAllocReportID returns true if AllocReportID is present, Tag 755
func (m AllocationReportAck) HasAllocReportID() bool {
	return m.Has(tag.AllocReportID)
}

//HasSecondaryAllocID returns true if SecondaryAllocID is present, Tag 793
func (m AllocationReportAck) HasSecondaryAllocID() bool {
	return m.Has(tag.SecondaryAllocID)
}

//HasAllocReportType returns true if AllocReportType is present, Tag 794
func (m AllocationReportAck) HasAllocReportType() bool {
	return m.Has(tag.AllocReportType)
}

//HasAllocIntermedReqType returns true if AllocIntermedReqType is present, Tag 808
func (m AllocationReportAck) HasAllocIntermedReqType() bool {
	return m.Has(tag.AllocIntermedReqType)
}

//NoAllocs is a repeating group element, Tag 78
type NoAllocs struct {
	*quickfix.Group
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
func (m NoAllocs) SetAllocPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewAllocPrice(value, scale))
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

//GetAllocAccount gets AllocAccount, Tag 79
func (m NoAllocs) GetAllocAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AllocAccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocAcctIDSource gets AllocAcctIDSource, Tag 661
func (m NoAllocs) GetAllocAcctIDSource() (v int, err quickfix.MessageRejectError) {
	var f field.AllocAcctIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocPrice gets AllocPrice, Tag 366
func (m NoAllocs) GetAllocPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AllocPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIndividualAllocID gets IndividualAllocID, Tag 467
func (m NoAllocs) GetIndividualAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.IndividualAllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIndividualAllocRejCode gets IndividualAllocRejCode, Tag 776
func (m NoAllocs) GetIndividualAllocRejCode() (v int, err quickfix.MessageRejectError) {
	var f field.IndividualAllocRejCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocText gets AllocText, Tag 161
func (m NoAllocs) GetAllocText() (v string, err quickfix.MessageRejectError) {
	var f field.AllocTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedAllocTextLen gets EncodedAllocTextLen, Tag 360
func (m NoAllocs) GetEncodedAllocTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedAllocTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedAllocText gets EncodedAllocText, Tag 361
func (m NoAllocs) GetEncodedAllocText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedAllocTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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

//NoAllocsRepeatingGroup is a repeating group, Tag 78
type NoAllocsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoAllocsRepeatingGroup returns an initialized, NoAllocsRepeatingGroup
func NewNoAllocsRepeatingGroup() NoAllocsRepeatingGroup {
	return NoAllocsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAllocs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AllocAccount), quickfix.GroupElement(tag.AllocAcctIDSource), quickfix.GroupElement(tag.AllocPrice), quickfix.GroupElement(tag.IndividualAllocID), quickfix.GroupElement(tag.IndividualAllocRejCode), quickfix.GroupElement(tag.AllocText), quickfix.GroupElement(tag.EncodedAllocTextLen), quickfix.GroupElement(tag.EncodedAllocText)})}
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
