package liststatus

import (
	"github.com/shopspring/decimal"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix41"
	"github.com/quickfixgo/quickfix/tag"
)

//ListStatus is the fix41 ListStatus type, MsgType = N
type ListStatus struct {
	fix41.Header
	*quickfix.Body
	fix41.Trailer
	Message *quickfix.Message
}

//FromMessage creates a ListStatus from a quickfix.Message instance
func FromMessage(m *quickfix.Message) ListStatus {
	return ListStatus{
		Header:  fix41.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix41.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ListStatus) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a ListStatus initialized with the required fields for ListStatus
func New(listid field.ListIDField, norpts field.NoRptsField, rptseq field.RptSeqField) (m ListStatus) {
	m.Message = quickfix.NewMessage()
	m.Header = fix41.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("N"))
	m.Set(listid)
	m.Set(norpts)
	m.Set(rptseq)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ListStatus, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.1", "N", r
}

//SetListID sets ListID, Tag 66
func (m ListStatus) SetListID(v string) {
	m.Set(field.NewListID(v))
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

//SetWaveNo sets WaveNo, Tag 105
func (m ListStatus) SetWaveNo(v string) {
	m.Set(field.NewWaveNo(v))
}

//GetListID gets ListID, Tag 66
func (m ListStatus) GetListID() (v string, err quickfix.MessageRejectError) {
	var f field.ListIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoOrders gets NoOrders, Tag 73
func (m ListStatus) GetNoOrders() (NoOrdersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoOrdersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoRpts gets NoRpts, Tag 82
func (m ListStatus) GetNoRpts() (v int, err quickfix.MessageRejectError) {
	var f field.NoRptsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRptSeq gets RptSeq, Tag 83
func (m ListStatus) GetRptSeq() (v int, err quickfix.MessageRejectError) {
	var f field.RptSeqField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetWaveNo gets WaveNo, Tag 105
func (m ListStatus) GetWaveNo() (v string, err quickfix.MessageRejectError) {
	var f field.WaveNoField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasListID returns true if ListID is present, Tag 66
func (m ListStatus) HasListID() bool {
	return m.Has(tag.ListID)
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

//HasWaveNo returns true if WaveNo is present, Tag 105
func (m ListStatus) HasWaveNo() bool {
	return m.Has(tag.WaveNo)
}

//NoOrders is a repeating group element, Tag 73
type NoOrders struct {
	*quickfix.Group
}

//SetClOrdID sets ClOrdID, Tag 11
func (m NoOrders) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetCumQty sets CumQty, Tag 14
func (m NoOrders) SetCumQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCumQty(value, scale))
}

//SetLeavesQty sets LeavesQty, Tag 151
func (m NoOrders) SetLeavesQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLeavesQty(value, scale))
}

//SetCxlQty sets CxlQty, Tag 84
func (m NoOrders) SetCxlQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCxlQty(value, scale))
}

//SetAvgPx sets AvgPx, Tag 6
func (m NoOrders) SetAvgPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewAvgPx(value, scale))
}

//GetClOrdID gets ClOrdID, Tag 11
func (m NoOrders) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCumQty gets CumQty, Tag 14
func (m NoOrders) GetCumQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CumQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLeavesQty gets LeavesQty, Tag 151
func (m NoOrders) GetLeavesQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LeavesQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCxlQty gets CxlQty, Tag 84
func (m NoOrders) GetCxlQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CxlQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAvgPx gets AvgPx, Tag 6
func (m NoOrders) GetAvgPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AvgPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m NoOrders) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasCumQty returns true if CumQty is present, Tag 14
func (m NoOrders) HasCumQty() bool {
	return m.Has(tag.CumQty)
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

//NoOrdersRepeatingGroup is a repeating group, Tag 73
type NoOrdersRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoOrdersRepeatingGroup returns an initialized, NoOrdersRepeatingGroup
func NewNoOrdersRepeatingGroup() NoOrdersRepeatingGroup {
	return NoOrdersRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoOrders,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ClOrdID), quickfix.GroupElement(tag.CumQty), quickfix.GroupElement(tag.LeavesQty), quickfix.GroupElement(tag.CxlQty), quickfix.GroupElement(tag.AvgPx)})}
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
