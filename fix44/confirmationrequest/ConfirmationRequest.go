package confirmationrequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/tag"
)

//ConfirmationRequest is the fix44 ConfirmationRequest type, MsgType = BH
type ConfirmationRequest struct {
	fix44.Header
	quickfix.Body
	fix44.Trailer
}

//FromMessage creates a ConfirmationRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) ConfirmationRequest {
	return ConfirmationRequest{
		Header:  fix44.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fix44.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m ConfirmationRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a ConfirmationRequest initialized with the required fields for ConfirmationRequest
func New(confirmreqid field.ConfirmReqIDField, confirmtype field.ConfirmTypeField, transacttime field.TransactTimeField) (m ConfirmationRequest) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("BH"))
	m.Set(confirmreqid)
	m.Set(confirmtype)
	m.Set(transacttime)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ConfirmationRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "BH", r
}

//SetText sets Text, Tag 58
func (m ConfirmationRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m ConfirmationRequest) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetAllocID sets AllocID, Tag 70
func (m ConfirmationRequest) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

//SetNoOrders sets NoOrders, Tag 73
func (m ConfirmationRequest) SetNoOrders(f NoOrdersRepeatingGroup) {
	m.SetGroup(f)
}

//SetAllocAccount sets AllocAccount, Tag 79
func (m ConfirmationRequest) SetAllocAccount(v string) {
	m.Set(field.NewAllocAccount(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m ConfirmationRequest) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m ConfirmationRequest) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetIndividualAllocID sets IndividualAllocID, Tag 467
func (m ConfirmationRequest) SetIndividualAllocID(v string) {
	m.Set(field.NewIndividualAllocID(v))
}

//SetAllocAcctIDSource sets AllocAcctIDSource, Tag 661
func (m ConfirmationRequest) SetAllocAcctIDSource(v int) {
	m.Set(field.NewAllocAcctIDSource(v))
}

//SetConfirmType sets ConfirmType, Tag 773
func (m ConfirmationRequest) SetConfirmType(v int) {
	m.Set(field.NewConfirmType(v))
}

//SetSecondaryAllocID sets SecondaryAllocID, Tag 793
func (m ConfirmationRequest) SetSecondaryAllocID(v string) {
	m.Set(field.NewSecondaryAllocID(v))
}

//SetAllocAccountType sets AllocAccountType, Tag 798
func (m ConfirmationRequest) SetAllocAccountType(v int) {
	m.Set(field.NewAllocAccountType(v))
}

//SetConfirmReqID sets ConfirmReqID, Tag 859
func (m ConfirmationRequest) SetConfirmReqID(v string) {
	m.Set(field.NewConfirmReqID(v))
}

//GetText gets Text, Tag 58
func (m ConfirmationRequest) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m ConfirmationRequest) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocID gets AllocID, Tag 70
func (m ConfirmationRequest) GetAllocID() (f field.AllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoOrders gets NoOrders, Tag 73
func (m ConfirmationRequest) GetNoOrders() (NoOrdersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoOrdersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAllocAccount gets AllocAccount, Tag 79
func (m ConfirmationRequest) GetAllocAccount() (f field.AllocAccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m ConfirmationRequest) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m ConfirmationRequest) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIndividualAllocID gets IndividualAllocID, Tag 467
func (m ConfirmationRequest) GetIndividualAllocID() (f field.IndividualAllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocAcctIDSource gets AllocAcctIDSource, Tag 661
func (m ConfirmationRequest) GetAllocAcctIDSource() (f field.AllocAcctIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetConfirmType gets ConfirmType, Tag 773
func (m ConfirmationRequest) GetConfirmType() (f field.ConfirmTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryAllocID gets SecondaryAllocID, Tag 793
func (m ConfirmationRequest) GetSecondaryAllocID() (f field.SecondaryAllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocAccountType gets AllocAccountType, Tag 798
func (m ConfirmationRequest) GetAllocAccountType() (f field.AllocAccountTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetConfirmReqID gets ConfirmReqID, Tag 859
func (m ConfirmationRequest) GetConfirmReqID() (f field.ConfirmReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasText returns true if Text is present, Tag 58
func (m ConfirmationRequest) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m ConfirmationRequest) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasAllocID returns true if AllocID is present, Tag 70
func (m ConfirmationRequest) HasAllocID() bool {
	return m.Has(tag.AllocID)
}

//HasNoOrders returns true if NoOrders is present, Tag 73
func (m ConfirmationRequest) HasNoOrders() bool {
	return m.Has(tag.NoOrders)
}

//HasAllocAccount returns true if AllocAccount is present, Tag 79
func (m ConfirmationRequest) HasAllocAccount() bool {
	return m.Has(tag.AllocAccount)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m ConfirmationRequest) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m ConfirmationRequest) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasIndividualAllocID returns true if IndividualAllocID is present, Tag 467
func (m ConfirmationRequest) HasIndividualAllocID() bool {
	return m.Has(tag.IndividualAllocID)
}

//HasAllocAcctIDSource returns true if AllocAcctIDSource is present, Tag 661
func (m ConfirmationRequest) HasAllocAcctIDSource() bool {
	return m.Has(tag.AllocAcctIDSource)
}

//HasConfirmType returns true if ConfirmType is present, Tag 773
func (m ConfirmationRequest) HasConfirmType() bool {
	return m.Has(tag.ConfirmType)
}

//HasSecondaryAllocID returns true if SecondaryAllocID is present, Tag 793
func (m ConfirmationRequest) HasSecondaryAllocID() bool {
	return m.Has(tag.SecondaryAllocID)
}

//HasAllocAccountType returns true if AllocAccountType is present, Tag 798
func (m ConfirmationRequest) HasAllocAccountType() bool {
	return m.Has(tag.AllocAccountType)
}

//HasConfirmReqID returns true if ConfirmReqID is present, Tag 859
func (m ConfirmationRequest) HasConfirmReqID() bool {
	return m.Has(tag.ConfirmReqID)
}

//NoOrders is a repeating group element, Tag 73
type NoOrders struct {
	quickfix.Group
}

//SetClOrdID sets ClOrdID, Tag 11
func (m NoOrders) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetOrderID sets OrderID, Tag 37
func (m NoOrders) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetSecondaryOrderID sets SecondaryOrderID, Tag 198
func (m NoOrders) SetSecondaryOrderID(v string) {
	m.Set(field.NewSecondaryOrderID(v))
}

//SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526
func (m NoOrders) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

//SetListID sets ListID, Tag 66
func (m NoOrders) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetNoNested2PartyIDs sets NoNested2PartyIDs, Tag 756
func (m NoOrders) SetNoNested2PartyIDs(f NoNested2PartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetOrderQty sets OrderQty, Tag 38
func (m NoOrders) SetOrderQty(v float64) {
	m.Set(field.NewOrderQty(v))
}

//SetOrderAvgPx sets OrderAvgPx, Tag 799
func (m NoOrders) SetOrderAvgPx(v float64) {
	m.Set(field.NewOrderAvgPx(v))
}

//SetOrderBookingQty sets OrderBookingQty, Tag 800
func (m NoOrders) SetOrderBookingQty(v float64) {
	m.Set(field.NewOrderBookingQty(v))
}

//GetClOrdID gets ClOrdID, Tag 11
func (m NoOrders) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m NoOrders) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryOrderID gets SecondaryOrderID, Tag 198
func (m NoOrders) GetSecondaryOrderID() (f field.SecondaryOrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526
func (m NoOrders) GetSecondaryClOrdID() (f field.SecondaryClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListID gets ListID, Tag 66
func (m NoOrders) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoNested2PartyIDs gets NoNested2PartyIDs, Tag 756
func (m NoOrders) GetNoNested2PartyIDs() (NoNested2PartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNested2PartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetOrderQty gets OrderQty, Tag 38
func (m NoOrders) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderAvgPx gets OrderAvgPx, Tag 799
func (m NoOrders) GetOrderAvgPx() (f field.OrderAvgPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderBookingQty gets OrderBookingQty, Tag 800
func (m NoOrders) GetOrderBookingQty() (f field.OrderBookingQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m NoOrders) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m NoOrders) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasSecondaryOrderID returns true if SecondaryOrderID is present, Tag 198
func (m NoOrders) HasSecondaryOrderID() bool {
	return m.Has(tag.SecondaryOrderID)
}

//HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526
func (m NoOrders) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

//HasListID returns true if ListID is present, Tag 66
func (m NoOrders) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasNoNested2PartyIDs returns true if NoNested2PartyIDs is present, Tag 756
func (m NoOrders) HasNoNested2PartyIDs() bool {
	return m.Has(tag.NoNested2PartyIDs)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m NoOrders) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasOrderAvgPx returns true if OrderAvgPx is present, Tag 799
func (m NoOrders) HasOrderAvgPx() bool {
	return m.Has(tag.OrderAvgPx)
}

//HasOrderBookingQty returns true if OrderBookingQty is present, Tag 800
func (m NoOrders) HasOrderBookingQty() bool {
	return m.Has(tag.OrderBookingQty)
}

//NoNested2PartyIDs is a repeating group element, Tag 756
type NoNested2PartyIDs struct {
	quickfix.Group
}

//SetNested2PartyID sets Nested2PartyID, Tag 757
func (m NoNested2PartyIDs) SetNested2PartyID(v string) {
	m.Set(field.NewNested2PartyID(v))
}

//SetNested2PartyIDSource sets Nested2PartyIDSource, Tag 758
func (m NoNested2PartyIDs) SetNested2PartyIDSource(v string) {
	m.Set(field.NewNested2PartyIDSource(v))
}

//SetNested2PartyRole sets Nested2PartyRole, Tag 759
func (m NoNested2PartyIDs) SetNested2PartyRole(v int) {
	m.Set(field.NewNested2PartyRole(v))
}

//SetNoNested2PartySubIDs sets NoNested2PartySubIDs, Tag 806
func (m NoNested2PartyIDs) SetNoNested2PartySubIDs(f NoNested2PartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetNested2PartyID gets Nested2PartyID, Tag 757
func (m NoNested2PartyIDs) GetNested2PartyID() (f field.Nested2PartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNested2PartyIDSource gets Nested2PartyIDSource, Tag 758
func (m NoNested2PartyIDs) GetNested2PartyIDSource() (f field.Nested2PartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNested2PartyRole gets Nested2PartyRole, Tag 759
func (m NoNested2PartyIDs) GetNested2PartyRole() (f field.Nested2PartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoNested2PartySubIDs gets NoNested2PartySubIDs, Tag 806
func (m NoNested2PartyIDs) GetNoNested2PartySubIDs() (NoNested2PartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNested2PartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasNested2PartyID returns true if Nested2PartyID is present, Tag 757
func (m NoNested2PartyIDs) HasNested2PartyID() bool {
	return m.Has(tag.Nested2PartyID)
}

//HasNested2PartyIDSource returns true if Nested2PartyIDSource is present, Tag 758
func (m NoNested2PartyIDs) HasNested2PartyIDSource() bool {
	return m.Has(tag.Nested2PartyIDSource)
}

//HasNested2PartyRole returns true if Nested2PartyRole is present, Tag 759
func (m NoNested2PartyIDs) HasNested2PartyRole() bool {
	return m.Has(tag.Nested2PartyRole)
}

//HasNoNested2PartySubIDs returns true if NoNested2PartySubIDs is present, Tag 806
func (m NoNested2PartyIDs) HasNoNested2PartySubIDs() bool {
	return m.Has(tag.NoNested2PartySubIDs)
}

//NoNested2PartySubIDs is a repeating group element, Tag 806
type NoNested2PartySubIDs struct {
	quickfix.Group
}

//SetNested2PartySubID sets Nested2PartySubID, Tag 760
func (m NoNested2PartySubIDs) SetNested2PartySubID(v string) {
	m.Set(field.NewNested2PartySubID(v))
}

//SetNested2PartySubIDType sets Nested2PartySubIDType, Tag 807
func (m NoNested2PartySubIDs) SetNested2PartySubIDType(v int) {
	m.Set(field.NewNested2PartySubIDType(v))
}

//GetNested2PartySubID gets Nested2PartySubID, Tag 760
func (m NoNested2PartySubIDs) GetNested2PartySubID() (f field.Nested2PartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNested2PartySubIDType gets Nested2PartySubIDType, Tag 807
func (m NoNested2PartySubIDs) GetNested2PartySubIDType() (f field.Nested2PartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasNested2PartySubID returns true if Nested2PartySubID is present, Tag 760
func (m NoNested2PartySubIDs) HasNested2PartySubID() bool {
	return m.Has(tag.Nested2PartySubID)
}

//HasNested2PartySubIDType returns true if Nested2PartySubIDType is present, Tag 807
func (m NoNested2PartySubIDs) HasNested2PartySubIDType() bool {
	return m.Has(tag.Nested2PartySubIDType)
}

//NoNested2PartySubIDsRepeatingGroup is a repeating group, Tag 806
type NoNested2PartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNested2PartySubIDsRepeatingGroup returns an initialized, NoNested2PartySubIDsRepeatingGroup
func NewNoNested2PartySubIDsRepeatingGroup() NoNested2PartySubIDsRepeatingGroup {
	return NoNested2PartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNested2PartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Nested2PartySubID), quickfix.GroupElement(tag.Nested2PartySubIDType)})}
}

//Add create and append a new NoNested2PartySubIDs to this group
func (m NoNested2PartySubIDsRepeatingGroup) Add() NoNested2PartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoNested2PartySubIDs{g}
}

//Get returns the ith NoNested2PartySubIDs in the NoNested2PartySubIDsRepeatinGroup
func (m NoNested2PartySubIDsRepeatingGroup) Get(i int) NoNested2PartySubIDs {
	return NoNested2PartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoNested2PartyIDsRepeatingGroup is a repeating group, Tag 756
type NoNested2PartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNested2PartyIDsRepeatingGroup returns an initialized, NoNested2PartyIDsRepeatingGroup
func NewNoNested2PartyIDsRepeatingGroup() NoNested2PartyIDsRepeatingGroup {
	return NoNested2PartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNested2PartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Nested2PartyID), quickfix.GroupElement(tag.Nested2PartyIDSource), quickfix.GroupElement(tag.Nested2PartyRole), quickfix.GroupElement(tag.NoNested2PartySubIDs)})}
}

//Add create and append a new NoNested2PartyIDs to this group
func (m NoNested2PartyIDsRepeatingGroup) Add() NoNested2PartyIDs {
	g := m.RepeatingGroup.Add()
	return NoNested2PartyIDs{g}
}

//Get returns the ith NoNested2PartyIDs in the NoNested2PartyIDsRepeatinGroup
func (m NoNested2PartyIDsRepeatingGroup) Get(i int) NoNested2PartyIDs {
	return NoNested2PartyIDs{m.RepeatingGroup.Get(i)}
}

//NoOrdersRepeatingGroup is a repeating group, Tag 73
type NoOrdersRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoOrdersRepeatingGroup returns an initialized, NoOrdersRepeatingGroup
func NewNoOrdersRepeatingGroup() NoOrdersRepeatingGroup {
	return NoOrdersRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoOrders,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ClOrdID), quickfix.GroupElement(tag.OrderID), quickfix.GroupElement(tag.SecondaryOrderID), quickfix.GroupElement(tag.SecondaryClOrdID), quickfix.GroupElement(tag.ListID), quickfix.GroupElement(tag.NoNested2PartyIDs), quickfix.GroupElement(tag.OrderQty), quickfix.GroupElement(tag.OrderAvgPx), quickfix.GroupElement(tag.OrderBookingQty)})}
}

//Add create and append a new NoOrders to this group
func (m NoOrdersRepeatingGroup) Add() NoOrders {
	g := m.RepeatingGroup.Add()
	return NoOrders{g}
}

//Get returns the ith NoOrders in the NoOrdersRepeatinGroup
func (m NoOrdersRepeatingGroup) Get(i int) NoOrders {
	return NoOrders{m.RepeatingGroup.Get(i)}
}
