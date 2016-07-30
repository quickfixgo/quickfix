package liststatus

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//ListStatus is the fix43 ListStatus type, MsgType = N
type ListStatus struct {
	fix43.Header
	quickfix.Body
	fix43.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a ListStatus from a quickfix.Message instance
func FromMessage(m quickfix.Message) ListStatus {
	return ListStatus{
		Header:      fix43.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix43.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ListStatus) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a ListStatus initialized with the required fields for ListStatus
func New(listid field.ListIDField, liststatustype field.ListStatusTypeField, norpts field.NoRptsField, listorderstatus field.ListOrderStatusField, rptseq field.RptSeqField, totnoorders field.TotNoOrdersField) (m ListStatus) {
	m.Header = fix43.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("N"))
	m.Header.Set(field.NewBeginString("FIX.4.3"))
	m.Set(listid)
	m.Set(liststatustype)
	m.Set(norpts)
	m.Set(listorderstatus)
	m.Set(rptseq)
	m.Set(totnoorders)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ListStatus, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "N", r
}

//SetTransactTime sets TransactTime, Tag 60
func (m ListStatus) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetListID sets ListID, Tag 66
func (m ListStatus) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetTotNoOrders sets TotNoOrders, Tag 68
func (m ListStatus) SetTotNoOrders(v int) {
	m.Set(field.NewTotNoOrders(v))
}

//SetNoOrders sets NoOrders, Tag 73
func (m ListStatus) SetNoOrders(f NoOrdersRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoRpts sets NoRpts, Tag 82
func (m ListStatus) SetNoRpts(v int) {
	m.Set(field.NewNoRpts(v))
}

//SetRptSeq sets RptSeq, Tag 83
func (m ListStatus) SetRptSeq(v int) {
	m.Set(field.NewRptSeq(v))
}

//SetListStatusType sets ListStatusType, Tag 429
func (m ListStatus) SetListStatusType(v int) {
	m.Set(field.NewListStatusType(v))
}

//SetListOrderStatus sets ListOrderStatus, Tag 431
func (m ListStatus) SetListOrderStatus(v int) {
	m.Set(field.NewListOrderStatus(v))
}

//SetListStatusText sets ListStatusText, Tag 444
func (m ListStatus) SetListStatusText(v string) {
	m.Set(field.NewListStatusText(v))
}

//SetEncodedListStatusTextLen sets EncodedListStatusTextLen, Tag 445
func (m ListStatus) SetEncodedListStatusTextLen(v int) {
	m.Set(field.NewEncodedListStatusTextLen(v))
}

//SetEncodedListStatusText sets EncodedListStatusText, Tag 446
func (m ListStatus) SetEncodedListStatusText(v string) {
	m.Set(field.NewEncodedListStatusText(v))
}

//GetTransactTime gets TransactTime, Tag 60
func (m ListStatus) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListID gets ListID, Tag 66
func (m ListStatus) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTotNoOrders gets TotNoOrders, Tag 68
func (m ListStatus) GetTotNoOrders() (f field.TotNoOrdersField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoOrders gets NoOrders, Tag 73
func (m ListStatus) GetNoOrders() (NoOrdersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoOrdersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoRpts gets NoRpts, Tag 82
func (m ListStatus) GetNoRpts() (f field.NoRptsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRptSeq gets RptSeq, Tag 83
func (m ListStatus) GetRptSeq() (f field.RptSeqField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListStatusType gets ListStatusType, Tag 429
func (m ListStatus) GetListStatusType() (f field.ListStatusTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListOrderStatus gets ListOrderStatus, Tag 431
func (m ListStatus) GetListOrderStatus() (f field.ListOrderStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListStatusText gets ListStatusText, Tag 444
func (m ListStatus) GetListStatusText() (f field.ListStatusTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedListStatusTextLen gets EncodedListStatusTextLen, Tag 445
func (m ListStatus) GetEncodedListStatusTextLen() (f field.EncodedListStatusTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedListStatusText gets EncodedListStatusText, Tag 446
func (m ListStatus) GetEncodedListStatusText() (f field.EncodedListStatusTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m ListStatus) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasListID returns true if ListID is present, Tag 66
func (m ListStatus) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasTotNoOrders returns true if TotNoOrders is present, Tag 68
func (m ListStatus) HasTotNoOrders() bool {
	return m.Has(tag.TotNoOrders)
}

//HasNoOrders returns true if NoOrders is present, Tag 73
func (m ListStatus) HasNoOrders() bool {
	return m.Has(tag.NoOrders)
}

//HasNoRpts returns true if NoRpts is present, Tag 82
func (m ListStatus) HasNoRpts() bool {
	return m.Has(tag.NoRpts)
}

//HasRptSeq returns true if RptSeq is present, Tag 83
func (m ListStatus) HasRptSeq() bool {
	return m.Has(tag.RptSeq)
}

//HasListStatusType returns true if ListStatusType is present, Tag 429
func (m ListStatus) HasListStatusType() bool {
	return m.Has(tag.ListStatusType)
}

//HasListOrderStatus returns true if ListOrderStatus is present, Tag 431
func (m ListStatus) HasListOrderStatus() bool {
	return m.Has(tag.ListOrderStatus)
}

//HasListStatusText returns true if ListStatusText is present, Tag 444
func (m ListStatus) HasListStatusText() bool {
	return m.Has(tag.ListStatusText)
}

//HasEncodedListStatusTextLen returns true if EncodedListStatusTextLen is present, Tag 445
func (m ListStatus) HasEncodedListStatusTextLen() bool {
	return m.Has(tag.EncodedListStatusTextLen)
}

//HasEncodedListStatusText returns true if EncodedListStatusText is present, Tag 446
func (m ListStatus) HasEncodedListStatusText() bool {
	return m.Has(tag.EncodedListStatusText)
}

//NoOrders is a repeating group element, Tag 73
type NoOrders struct {
	quickfix.Group
}

//SetClOrdID sets ClOrdID, Tag 11
func (m NoOrders) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526
func (m NoOrders) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

//SetCumQty sets CumQty, Tag 14
func (m NoOrders) SetCumQty(v float64) {
	m.Set(field.NewCumQty(v))
}

//SetOrdStatus sets OrdStatus, Tag 39
func (m NoOrders) SetOrdStatus(v string) {
	m.Set(field.NewOrdStatus(v))
}

//SetWorkingIndicator sets WorkingIndicator, Tag 636
func (m NoOrders) SetWorkingIndicator(v bool) {
	m.Set(field.NewWorkingIndicator(v))
}

//SetLeavesQty sets LeavesQty, Tag 151
func (m NoOrders) SetLeavesQty(v float64) {
	m.Set(field.NewLeavesQty(v))
}

//SetCxlQty sets CxlQty, Tag 84
func (m NoOrders) SetCxlQty(v float64) {
	m.Set(field.NewCxlQty(v))
}

//SetAvgPx sets AvgPx, Tag 6
func (m NoOrders) SetAvgPx(v float64) {
	m.Set(field.NewAvgPx(v))
}

//SetOrdRejReason sets OrdRejReason, Tag 103
func (m NoOrders) SetOrdRejReason(v int) {
	m.Set(field.NewOrdRejReason(v))
}

//SetText sets Text, Tag 58
func (m NoOrders) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m NoOrders) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m NoOrders) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//GetClOrdID gets ClOrdID, Tag 11
func (m NoOrders) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526
func (m NoOrders) GetSecondaryClOrdID() (f field.SecondaryClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCumQty gets CumQty, Tag 14
func (m NoOrders) GetCumQty() (f field.CumQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdStatus gets OrdStatus, Tag 39
func (m NoOrders) GetOrdStatus() (f field.OrdStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetWorkingIndicator gets WorkingIndicator, Tag 636
func (m NoOrders) GetWorkingIndicator() (f field.WorkingIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLeavesQty gets LeavesQty, Tag 151
func (m NoOrders) GetLeavesQty() (f field.LeavesQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCxlQty gets CxlQty, Tag 84
func (m NoOrders) GetCxlQty() (f field.CxlQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAvgPx gets AvgPx, Tag 6
func (m NoOrders) GetAvgPx() (f field.AvgPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdRejReason gets OrdRejReason, Tag 103
func (m NoOrders) GetOrdRejReason() (f field.OrdRejReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m NoOrders) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m NoOrders) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m NoOrders) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m NoOrders) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526
func (m NoOrders) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

//HasCumQty returns true if CumQty is present, Tag 14
func (m NoOrders) HasCumQty() bool {
	return m.Has(tag.CumQty)
}

//HasOrdStatus returns true if OrdStatus is present, Tag 39
func (m NoOrders) HasOrdStatus() bool {
	return m.Has(tag.OrdStatus)
}

//HasWorkingIndicator returns true if WorkingIndicator is present, Tag 636
func (m NoOrders) HasWorkingIndicator() bool {
	return m.Has(tag.WorkingIndicator)
}

//HasLeavesQty returns true if LeavesQty is present, Tag 151
func (m NoOrders) HasLeavesQty() bool {
	return m.Has(tag.LeavesQty)
}

//HasCxlQty returns true if CxlQty is present, Tag 84
func (m NoOrders) HasCxlQty() bool {
	return m.Has(tag.CxlQty)
}

//HasAvgPx returns true if AvgPx is present, Tag 6
func (m NoOrders) HasAvgPx() bool {
	return m.Has(tag.AvgPx)
}

//HasOrdRejReason returns true if OrdRejReason is present, Tag 103
func (m NoOrders) HasOrdRejReason() bool {
	return m.Has(tag.OrdRejReason)
}

//HasText returns true if Text is present, Tag 58
func (m NoOrders) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m NoOrders) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m NoOrders) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//NoOrdersRepeatingGroup is a repeating group, Tag 73
type NoOrdersRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoOrdersRepeatingGroup returns an initialized, NoOrdersRepeatingGroup
func NewNoOrdersRepeatingGroup() NoOrdersRepeatingGroup {
	return NoOrdersRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoOrders,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ClOrdID), quickfix.GroupElement(tag.SecondaryClOrdID), quickfix.GroupElement(tag.CumQty), quickfix.GroupElement(tag.OrdStatus), quickfix.GroupElement(tag.WorkingIndicator), quickfix.GroupElement(tag.LeavesQty), quickfix.GroupElement(tag.CxlQty), quickfix.GroupElement(tag.AvgPx), quickfix.GroupElement(tag.OrdRejReason), quickfix.GroupElement(tag.Text), quickfix.GroupElement(tag.EncodedTextLen), quickfix.GroupElement(tag.EncodedText)})}
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
