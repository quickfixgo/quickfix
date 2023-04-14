package allocationinstructionack

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/gen/enum"
	"github.com/quickfixgo/quickfix/gen/field"
	"github.com/quickfixgo/quickfix/gen/fixt11"
	"github.com/quickfixgo/quickfix/gen/tag"
)

// AllocationInstructionAck is the fix50sp2 AllocationInstructionAck type, MsgType = P.
type AllocationInstructionAck struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a AllocationInstructionAck from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) AllocationInstructionAck {
	return AllocationInstructionAck{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m AllocationInstructionAck) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a AllocationInstructionAck initialized with the required fields for AllocationInstructionAck.
func New(allocid field.AllocIDField, allocstatus field.AllocStatusField) (m AllocationInstructionAck) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("P"))
	m.Set(allocid)
	m.Set(allocstatus)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg AllocationInstructionAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "P", r
}

// SetText sets Text, Tag 58.
func (m AllocationInstructionAck) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m AllocationInstructionAck) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetAllocID sets AllocID, Tag 70.
func (m AllocationInstructionAck) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

// SetTradeDate sets TradeDate, Tag 75.
func (m AllocationInstructionAck) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

// SetNoAllocs sets NoAllocs, Tag 78.
func (m AllocationInstructionAck) SetNoAllocs(f NoAllocsRepeatingGroup) {
	m.SetGroup(f)
}

// SetAllocStatus sets AllocStatus, Tag 87.
func (m AllocationInstructionAck) SetAllocStatus(v enum.AllocStatus) {
	m.Set(field.NewAllocStatus(v))
}

// SetAllocRejCode sets AllocRejCode, Tag 88.
func (m AllocationInstructionAck) SetAllocRejCode(v enum.AllocRejCode) {
	m.Set(field.NewAllocRejCode(v))
}

// SetSecurityType sets SecurityType, Tag 167.
func (m AllocationInstructionAck) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m AllocationInstructionAck) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m AllocationInstructionAck) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m AllocationInstructionAck) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetProduct sets Product, Tag 460.
func (m AllocationInstructionAck) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

// SetMatchStatus sets MatchStatus, Tag 573.
func (m AllocationInstructionAck) SetMatchStatus(v enum.MatchStatus) {
	m.Set(field.NewMatchStatus(v))
}

// SetAllocType sets AllocType, Tag 626.
func (m AllocationInstructionAck) SetAllocType(v enum.AllocType) {
	m.Set(field.NewAllocType(v))
}

// SetSecondaryAllocID sets SecondaryAllocID, Tag 793.
func (m AllocationInstructionAck) SetSecondaryAllocID(v string) {
	m.Set(field.NewSecondaryAllocID(v))
}

// SetAllocIntermedReqType sets AllocIntermedReqType, Tag 808.
func (m AllocationInstructionAck) SetAllocIntermedReqType(v enum.AllocIntermedReqType) {
	m.Set(field.NewAllocIntermedReqType(v))
}

// GetText gets Text, Tag 58.
func (m AllocationInstructionAck) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m AllocationInstructionAck) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocID gets AllocID, Tag 70.
func (m AllocationInstructionAck) GetAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.AllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradeDate gets TradeDate, Tag 75.
func (m AllocationInstructionAck) GetTradeDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoAllocs gets NoAllocs, Tag 78.
func (m AllocationInstructionAck) GetNoAllocs() (NoAllocsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAllocsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetAllocStatus gets AllocStatus, Tag 87.
func (m AllocationInstructionAck) GetAllocStatus() (v enum.AllocStatus, err quickfix.MessageRejectError) {
	var f field.AllocStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocRejCode gets AllocRejCode, Tag 88.
func (m AllocationInstructionAck) GetAllocRejCode() (v enum.AllocRejCode, err quickfix.MessageRejectError) {
	var f field.AllocRejCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityType gets SecurityType, Tag 167.
func (m AllocationInstructionAck) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m AllocationInstructionAck) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m AllocationInstructionAck) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m AllocationInstructionAck) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetProduct gets Product, Tag 460.
func (m AllocationInstructionAck) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMatchStatus gets MatchStatus, Tag 573.
func (m AllocationInstructionAck) GetMatchStatus() (v enum.MatchStatus, err quickfix.MessageRejectError) {
	var f field.MatchStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocType gets AllocType, Tag 626.
func (m AllocationInstructionAck) GetAllocType() (v enum.AllocType, err quickfix.MessageRejectError) {
	var f field.AllocTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecondaryAllocID gets SecondaryAllocID, Tag 793.
func (m AllocationInstructionAck) GetSecondaryAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryAllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocIntermedReqType gets AllocIntermedReqType, Tag 808.
func (m AllocationInstructionAck) GetAllocIntermedReqType() (v enum.AllocIntermedReqType, err quickfix.MessageRejectError) {
	var f field.AllocIntermedReqTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasText returns true if Text is present, Tag 58.
func (m AllocationInstructionAck) HasText() bool {
	return m.Has(tag.Text)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m AllocationInstructionAck) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasAllocID returns true if AllocID is present, Tag 70.
func (m AllocationInstructionAck) HasAllocID() bool {
	return m.Has(tag.AllocID)
}

// HasTradeDate returns true if TradeDate is present, Tag 75.
func (m AllocationInstructionAck) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

// HasNoAllocs returns true if NoAllocs is present, Tag 78.
func (m AllocationInstructionAck) HasNoAllocs() bool {
	return m.Has(tag.NoAllocs)
}

// HasAllocStatus returns true if AllocStatus is present, Tag 87.
func (m AllocationInstructionAck) HasAllocStatus() bool {
	return m.Has(tag.AllocStatus)
}

// HasAllocRejCode returns true if AllocRejCode is present, Tag 88.
func (m AllocationInstructionAck) HasAllocRejCode() bool {
	return m.Has(tag.AllocRejCode)
}

// HasSecurityType returns true if SecurityType is present, Tag 167.
func (m AllocationInstructionAck) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m AllocationInstructionAck) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m AllocationInstructionAck) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m AllocationInstructionAck) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasProduct returns true if Product is present, Tag 460.
func (m AllocationInstructionAck) HasProduct() bool {
	return m.Has(tag.Product)
}

// HasMatchStatus returns true if MatchStatus is present, Tag 573.
func (m AllocationInstructionAck) HasMatchStatus() bool {
	return m.Has(tag.MatchStatus)
}

// HasAllocType returns true if AllocType is present, Tag 626.
func (m AllocationInstructionAck) HasAllocType() bool {
	return m.Has(tag.AllocType)
}

// HasSecondaryAllocID returns true if SecondaryAllocID is present, Tag 793.
func (m AllocationInstructionAck) HasSecondaryAllocID() bool {
	return m.Has(tag.SecondaryAllocID)
}

// HasAllocIntermedReqType returns true if AllocIntermedReqType is present, Tag 808.
func (m AllocationInstructionAck) HasAllocIntermedReqType() bool {
	return m.Has(tag.AllocIntermedReqType)
}

// NoAllocs is a repeating group element, Tag 78.
type NoAllocs struct {
	*quickfix.Group
}

// SetAllocAccount sets AllocAccount, Tag 79.
func (m NoAllocs) SetAllocAccount(v string) {
	m.Set(field.NewAllocAccount(v))
}

// SetAllocAcctIDSource sets AllocAcctIDSource, Tag 661.
func (m NoAllocs) SetAllocAcctIDSource(v int) {
	m.Set(field.NewAllocAcctIDSource(v))
}

// SetAllocPrice sets AllocPrice, Tag 366.
func (m NoAllocs) SetAllocPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewAllocPrice(value, scale))
}

// SetIndividualAllocID sets IndividualAllocID, Tag 467.
func (m NoAllocs) SetIndividualAllocID(v string) {
	m.Set(field.NewIndividualAllocID(v))
}

// SetIndividualAllocRejCode sets IndividualAllocRejCode, Tag 776.
func (m NoAllocs) SetIndividualAllocRejCode(v int) {
	m.Set(field.NewIndividualAllocRejCode(v))
}

// SetAllocText sets AllocText, Tag 161.
func (m NoAllocs) SetAllocText(v string) {
	m.Set(field.NewAllocText(v))
}

// SetEncodedAllocTextLen sets EncodedAllocTextLen, Tag 360.
func (m NoAllocs) SetEncodedAllocTextLen(v int) {
	m.Set(field.NewEncodedAllocTextLen(v))
}

// SetEncodedAllocText sets EncodedAllocText, Tag 361.
func (m NoAllocs) SetEncodedAllocText(v string) {
	m.Set(field.NewEncodedAllocText(v))
}

// SetSecondaryIndividualAllocID sets SecondaryIndividualAllocID, Tag 989.
func (m NoAllocs) SetSecondaryIndividualAllocID(v string) {
	m.Set(field.NewSecondaryIndividualAllocID(v))
}

// SetAllocCustomerCapacity sets AllocCustomerCapacity, Tag 993.
func (m NoAllocs) SetAllocCustomerCapacity(v string) {
	m.Set(field.NewAllocCustomerCapacity(v))
}

// SetIndividualAllocType sets IndividualAllocType, Tag 992.
func (m NoAllocs) SetIndividualAllocType(v enum.IndividualAllocType) {
	m.Set(field.NewIndividualAllocType(v))
}

// SetAllocQty sets AllocQty, Tag 80.
func (m NoAllocs) SetAllocQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewAllocQty(value, scale))
}

// SetNoNestedPartyIDs sets NoNestedPartyIDs, Tag 539.
func (m NoAllocs) SetNoNestedPartyIDs(f NoNestedPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetAllocPositionEffect sets AllocPositionEffect, Tag 1047.
func (m NoAllocs) SetAllocPositionEffect(v enum.AllocPositionEffect) {
	m.Set(field.NewAllocPositionEffect(v))
}

// GetAllocAccount gets AllocAccount, Tag 79.
func (m NoAllocs) GetAllocAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AllocAccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocAcctIDSource gets AllocAcctIDSource, Tag 661.
func (m NoAllocs) GetAllocAcctIDSource() (v int, err quickfix.MessageRejectError) {
	var f field.AllocAcctIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocPrice gets AllocPrice, Tag 366.
func (m NoAllocs) GetAllocPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AllocPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIndividualAllocID gets IndividualAllocID, Tag 467.
func (m NoAllocs) GetIndividualAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.IndividualAllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIndividualAllocRejCode gets IndividualAllocRejCode, Tag 776.
func (m NoAllocs) GetIndividualAllocRejCode() (v int, err quickfix.MessageRejectError) {
	var f field.IndividualAllocRejCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocText gets AllocText, Tag 161.
func (m NoAllocs) GetAllocText() (v string, err quickfix.MessageRejectError) {
	var f field.AllocTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedAllocTextLen gets EncodedAllocTextLen, Tag 360.
func (m NoAllocs) GetEncodedAllocTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedAllocTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedAllocText gets EncodedAllocText, Tag 361.
func (m NoAllocs) GetEncodedAllocText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedAllocTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecondaryIndividualAllocID gets SecondaryIndividualAllocID, Tag 989.
func (m NoAllocs) GetSecondaryIndividualAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryIndividualAllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocCustomerCapacity gets AllocCustomerCapacity, Tag 993.
func (m NoAllocs) GetAllocCustomerCapacity() (v string, err quickfix.MessageRejectError) {
	var f field.AllocCustomerCapacityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIndividualAllocType gets IndividualAllocType, Tag 992.
func (m NoAllocs) GetIndividualAllocType() (v enum.IndividualAllocType, err quickfix.MessageRejectError) {
	var f field.IndividualAllocTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAllocQty gets AllocQty, Tag 80.
func (m NoAllocs) GetAllocQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AllocQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoNestedPartyIDs gets NoNestedPartyIDs, Tag 539.
func (m NoAllocs) GetNoNestedPartyIDs() (NoNestedPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetAllocPositionEffect gets AllocPositionEffect, Tag 1047.
func (m NoAllocs) GetAllocPositionEffect() (v enum.AllocPositionEffect, err quickfix.MessageRejectError) {
	var f field.AllocPositionEffectField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasAllocAccount returns true if AllocAccount is present, Tag 79.
func (m NoAllocs) HasAllocAccount() bool {
	return m.Has(tag.AllocAccount)
}

// HasAllocAcctIDSource returns true if AllocAcctIDSource is present, Tag 661.
func (m NoAllocs) HasAllocAcctIDSource() bool {
	return m.Has(tag.AllocAcctIDSource)
}

// HasAllocPrice returns true if AllocPrice is present, Tag 366.
func (m NoAllocs) HasAllocPrice() bool {
	return m.Has(tag.AllocPrice)
}

// HasIndividualAllocID returns true if IndividualAllocID is present, Tag 467.
func (m NoAllocs) HasIndividualAllocID() bool {
	return m.Has(tag.IndividualAllocID)
}

// HasIndividualAllocRejCode returns true if IndividualAllocRejCode is present, Tag 776.
func (m NoAllocs) HasIndividualAllocRejCode() bool {
	return m.Has(tag.IndividualAllocRejCode)
}

// HasAllocText returns true if AllocText is present, Tag 161.
func (m NoAllocs) HasAllocText() bool {
	return m.Has(tag.AllocText)
}

// HasEncodedAllocTextLen returns true if EncodedAllocTextLen is present, Tag 360.
func (m NoAllocs) HasEncodedAllocTextLen() bool {
	return m.Has(tag.EncodedAllocTextLen)
}

// HasEncodedAllocText returns true if EncodedAllocText is present, Tag 361.
func (m NoAllocs) HasEncodedAllocText() bool {
	return m.Has(tag.EncodedAllocText)
}

// HasSecondaryIndividualAllocID returns true if SecondaryIndividualAllocID is present, Tag 989.
func (m NoAllocs) HasSecondaryIndividualAllocID() bool {
	return m.Has(tag.SecondaryIndividualAllocID)
}

// HasAllocCustomerCapacity returns true if AllocCustomerCapacity is present, Tag 993.
func (m NoAllocs) HasAllocCustomerCapacity() bool {
	return m.Has(tag.AllocCustomerCapacity)
}

// HasIndividualAllocType returns true if IndividualAllocType is present, Tag 992.
func (m NoAllocs) HasIndividualAllocType() bool {
	return m.Has(tag.IndividualAllocType)
}

// HasAllocQty returns true if AllocQty is present, Tag 80.
func (m NoAllocs) HasAllocQty() bool {
	return m.Has(tag.AllocQty)
}

// HasNoNestedPartyIDs returns true if NoNestedPartyIDs is present, Tag 539.
func (m NoAllocs) HasNoNestedPartyIDs() bool {
	return m.Has(tag.NoNestedPartyIDs)
}

// HasAllocPositionEffect returns true if AllocPositionEffect is present, Tag 1047.
func (m NoAllocs) HasAllocPositionEffect() bool {
	return m.Has(tag.AllocPositionEffect)
}

// NoNestedPartyIDs is a repeating group element, Tag 539.
type NoNestedPartyIDs struct {
	*quickfix.Group
}

// SetNestedPartyID sets NestedPartyID, Tag 524.
func (m NoNestedPartyIDs) SetNestedPartyID(v string) {
	m.Set(field.NewNestedPartyID(v))
}

// SetNestedPartyIDSource sets NestedPartyIDSource, Tag 525.
func (m NoNestedPartyIDs) SetNestedPartyIDSource(v string) {
	m.Set(field.NewNestedPartyIDSource(v))
}

// SetNestedPartyRole sets NestedPartyRole, Tag 538.
func (m NoNestedPartyIDs) SetNestedPartyRole(v int) {
	m.Set(field.NewNestedPartyRole(v))
}

// SetNoNestedPartySubIDs sets NoNestedPartySubIDs, Tag 804.
func (m NoNestedPartyIDs) SetNoNestedPartySubIDs(f NoNestedPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetNestedPartyID gets NestedPartyID, Tag 524.
func (m NoNestedPartyIDs) GetNestedPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNestedPartyIDSource gets NestedPartyIDSource, Tag 525.
func (m NoNestedPartyIDs) GetNestedPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNestedPartyRole gets NestedPartyRole, Tag 538.
func (m NoNestedPartyIDs) GetNestedPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.NestedPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoNestedPartySubIDs gets NoNestedPartySubIDs, Tag 804.
func (m NoNestedPartyIDs) GetNoNestedPartySubIDs() (NoNestedPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasNestedPartyID returns true if NestedPartyID is present, Tag 524.
func (m NoNestedPartyIDs) HasNestedPartyID() bool {
	return m.Has(tag.NestedPartyID)
}

// HasNestedPartyIDSource returns true if NestedPartyIDSource is present, Tag 525.
func (m NoNestedPartyIDs) HasNestedPartyIDSource() bool {
	return m.Has(tag.NestedPartyIDSource)
}

// HasNestedPartyRole returns true if NestedPartyRole is present, Tag 538.
func (m NoNestedPartyIDs) HasNestedPartyRole() bool {
	return m.Has(tag.NestedPartyRole)
}

// HasNoNestedPartySubIDs returns true if NoNestedPartySubIDs is present, Tag 804.
func (m NoNestedPartyIDs) HasNoNestedPartySubIDs() bool {
	return m.Has(tag.NoNestedPartySubIDs)
}

// NoNestedPartySubIDs is a repeating group element, Tag 804.
type NoNestedPartySubIDs struct {
	*quickfix.Group
}

// SetNestedPartySubID sets NestedPartySubID, Tag 545.
func (m NoNestedPartySubIDs) SetNestedPartySubID(v string) {
	m.Set(field.NewNestedPartySubID(v))
}

// SetNestedPartySubIDType sets NestedPartySubIDType, Tag 805.
func (m NoNestedPartySubIDs) SetNestedPartySubIDType(v int) {
	m.Set(field.NewNestedPartySubIDType(v))
}

// GetNestedPartySubID gets NestedPartySubID, Tag 545.
func (m NoNestedPartySubIDs) GetNestedPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNestedPartySubIDType gets NestedPartySubIDType, Tag 805.
func (m NoNestedPartySubIDs) GetNestedPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.NestedPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasNestedPartySubID returns true if NestedPartySubID is present, Tag 545.
func (m NoNestedPartySubIDs) HasNestedPartySubID() bool {
	return m.Has(tag.NestedPartySubID)
}

// HasNestedPartySubIDType returns true if NestedPartySubIDType is present, Tag 805.
func (m NoNestedPartySubIDs) HasNestedPartySubIDType() bool {
	return m.Has(tag.NestedPartySubIDType)
}

// NoNestedPartySubIDsRepeatingGroup is a repeating group, Tag 804.
type NoNestedPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoNestedPartySubIDsRepeatingGroup returns an initialized, NoNestedPartySubIDsRepeatingGroup.
func NewNoNestedPartySubIDsRepeatingGroup() NoNestedPartySubIDsRepeatingGroup {
	return NoNestedPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNestedPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NestedPartySubID), quickfix.GroupElement(tag.NestedPartySubIDType)})}
}

// Add create and append a new NoNestedPartySubIDs to this group.
func (m NoNestedPartySubIDsRepeatingGroup) Add() NoNestedPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoNestedPartySubIDs{g}
}

// Get returns the ith NoNestedPartySubIDs in the NoNestedPartySubIDsRepeatinGroup.
func (m NoNestedPartySubIDsRepeatingGroup) Get(i int) NoNestedPartySubIDs {
	return NoNestedPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoNestedPartyIDsRepeatingGroup is a repeating group, Tag 539.
type NoNestedPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoNestedPartyIDsRepeatingGroup returns an initialized, NoNestedPartyIDsRepeatingGroup.
func NewNoNestedPartyIDsRepeatingGroup() NoNestedPartyIDsRepeatingGroup {
	return NoNestedPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNestedPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NestedPartyID), quickfix.GroupElement(tag.NestedPartyIDSource), quickfix.GroupElement(tag.NestedPartyRole), NewNoNestedPartySubIDsRepeatingGroup()})}
}

// Add create and append a new NoNestedPartyIDs to this group.
func (m NoNestedPartyIDsRepeatingGroup) Add() NoNestedPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoNestedPartyIDs{g}
}

// Get returns the ith NoNestedPartyIDs in the NoNestedPartyIDsRepeatinGroup.
func (m NoNestedPartyIDsRepeatingGroup) Get(i int) NoNestedPartyIDs {
	return NoNestedPartyIDs{m.RepeatingGroup.Get(i)}
}

// NoAllocsRepeatingGroup is a repeating group, Tag 78.
type NoAllocsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoAllocsRepeatingGroup returns an initialized, NoAllocsRepeatingGroup.
func NewNoAllocsRepeatingGroup() NoAllocsRepeatingGroup {
	return NoAllocsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAllocs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AllocAccount), quickfix.GroupElement(tag.AllocAcctIDSource), quickfix.GroupElement(tag.AllocPrice), quickfix.GroupElement(tag.IndividualAllocID), quickfix.GroupElement(tag.IndividualAllocRejCode), quickfix.GroupElement(tag.AllocText), quickfix.GroupElement(tag.EncodedAllocTextLen), quickfix.GroupElement(tag.EncodedAllocText), quickfix.GroupElement(tag.SecondaryIndividualAllocID), quickfix.GroupElement(tag.AllocCustomerCapacity), quickfix.GroupElement(tag.IndividualAllocType), quickfix.GroupElement(tag.AllocQty), NewNoNestedPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.AllocPositionEffect)})}
}

// Add create and append a new NoAllocs to this group.
func (m NoAllocsRepeatingGroup) Add() NoAllocs {
	g := m.RepeatingGroup.Add()
	return NoAllocs{g}
}

// Get returns the ith NoAllocs in the NoAllocsRepeatinGroup.
func (m NoAllocsRepeatingGroup) Get(i int) NoAllocs {
	return NoAllocs{m.RepeatingGroup.Get(i)}
}

// NoPartyIDs is a repeating group element, Tag 453.
type NoPartyIDs struct {
	*quickfix.Group
}

// SetPartyID sets PartyID, Tag 448.
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

// SetPartyIDSource sets PartyIDSource, Tag 447.
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

// SetPartyRole sets PartyRole, Tag 452.
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

// SetNoPartySubIDs sets NoPartySubIDs, Tag 802.
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetPartyID gets PartyID, Tag 448.
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyIDSource gets PartyIDSource, Tag 447.
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyRole gets PartyRole, Tag 452.
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartySubIDs gets NoPartySubIDs, Tag 802.
func (m NoPartyIDs) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasPartyID returns true if PartyID is present, Tag 448.
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

// HasPartyIDSource returns true if PartyIDSource is present, Tag 447.
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

// HasPartyRole returns true if PartyRole is present, Tag 452.
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

// HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802.
func (m NoPartyIDs) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

// NoPartySubIDs is a repeating group element, Tag 802.
type NoPartySubIDs struct {
	*quickfix.Group
}

// SetPartySubID sets PartySubID, Tag 523.
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

// SetPartySubIDType sets PartySubIDType, Tag 803.
func (m NoPartySubIDs) SetPartySubIDType(v enum.PartySubIDType) {
	m.Set(field.NewPartySubIDType(v))
}

// GetPartySubID gets PartySubID, Tag 523.
func (m NoPartySubIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartySubIDType gets PartySubIDType, Tag 803.
func (m NoPartySubIDs) GetPartySubIDType() (v enum.PartySubIDType, err quickfix.MessageRejectError) {
	var f field.PartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartySubID returns true if PartySubID is present, Tag 523.
func (m NoPartySubIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

// HasPartySubIDType returns true if PartySubIDType is present, Tag 803.
func (m NoPartySubIDs) HasPartySubIDType() bool {
	return m.Has(tag.PartySubIDType)
}

// NoPartySubIDsRepeatingGroup is a repeating group, Tag 802.
type NoPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartySubIDsRepeatingGroup returns an initialized, NoPartySubIDsRepeatingGroup.
func NewNoPartySubIDsRepeatingGroup() NoPartySubIDsRepeatingGroup {
	return NoPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartySubID), quickfix.GroupElement(tag.PartySubIDType)})}
}

// Add create and append a new NoPartySubIDs to this group.
func (m NoPartySubIDsRepeatingGroup) Add() NoPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartySubIDs{g}
}

// Get returns the ith NoPartySubIDs in the NoPartySubIDsRepeatinGroup.
func (m NoPartySubIDsRepeatingGroup) Get(i int) NoPartySubIDs {
	return NoPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoPartyIDsRepeatingGroup is a repeating group, Tag 453.
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup.
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), NewNoPartySubIDsRepeatingGroup()})}
}

// Add create and append a new NoPartyIDs to this group.
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

// Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup.
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}
