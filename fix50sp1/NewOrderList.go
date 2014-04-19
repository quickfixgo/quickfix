package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//NewOrderList msg type = E.
type NewOrderList struct {
	message.Message
}

//NewOrderListBuilder builds NewOrderList messages.
type NewOrderListBuilder struct {
	message.MessageBuilder
}

//CreateNewOrderListBuilder returns an initialized NewOrderListBuilder with specified required fields.
func CreateNewOrderListBuilder(
	listid field.ListID,
	bidtype field.BidType,
	totnoorders field.TotNoOrders,
	noorders field.NoOrders) NewOrderListBuilder {
	var builder NewOrderListBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(listid)
	builder.Body.Set(bidtype)
	builder.Body.Set(totnoorders)
	builder.Body.Set(noorders)
	return builder
}

//ListID is a required field for NewOrderList.
func (m NewOrderList) ListID() (field.ListID, errors.MessageRejectError) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//BidID is a non-required field for NewOrderList.
func (m NewOrderList) BidID() (field.BidID, errors.MessageRejectError) {
	var f field.BidID
	err := m.Body.Get(&f)
	return f, err
}

//ClientBidID is a non-required field for NewOrderList.
func (m NewOrderList) ClientBidID() (field.ClientBidID, errors.MessageRejectError) {
	var f field.ClientBidID
	err := m.Body.Get(&f)
	return f, err
}

//ProgRptReqs is a non-required field for NewOrderList.
func (m NewOrderList) ProgRptReqs() (field.ProgRptReqs, errors.MessageRejectError) {
	var f field.ProgRptReqs
	err := m.Body.Get(&f)
	return f, err
}

//BidType is a required field for NewOrderList.
func (m NewOrderList) BidType() (field.BidType, errors.MessageRejectError) {
	var f field.BidType
	err := m.Body.Get(&f)
	return f, err
}

//ProgPeriodInterval is a non-required field for NewOrderList.
func (m NewOrderList) ProgPeriodInterval() (field.ProgPeriodInterval, errors.MessageRejectError) {
	var f field.ProgPeriodInterval
	err := m.Body.Get(&f)
	return f, err
}

//CancellationRights is a non-required field for NewOrderList.
func (m NewOrderList) CancellationRights() (field.CancellationRights, errors.MessageRejectError) {
	var f field.CancellationRights
	err := m.Body.Get(&f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for NewOrderList.
func (m NewOrderList) MoneyLaunderingStatus() (field.MoneyLaunderingStatus, errors.MessageRejectError) {
	var f field.MoneyLaunderingStatus
	err := m.Body.Get(&f)
	return f, err
}

//RegistID is a non-required field for NewOrderList.
func (m NewOrderList) RegistID() (field.RegistID, errors.MessageRejectError) {
	var f field.RegistID
	err := m.Body.Get(&f)
	return f, err
}

//ListExecInstType is a non-required field for NewOrderList.
func (m NewOrderList) ListExecInstType() (field.ListExecInstType, errors.MessageRejectError) {
	var f field.ListExecInstType
	err := m.Body.Get(&f)
	return f, err
}

//ListExecInst is a non-required field for NewOrderList.
func (m NewOrderList) ListExecInst() (field.ListExecInst, errors.MessageRejectError) {
	var f field.ListExecInst
	err := m.Body.Get(&f)
	return f, err
}

//EncodedListExecInstLen is a non-required field for NewOrderList.
func (m NewOrderList) EncodedListExecInstLen() (field.EncodedListExecInstLen, errors.MessageRejectError) {
	var f field.EncodedListExecInstLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedListExecInst is a non-required field for NewOrderList.
func (m NewOrderList) EncodedListExecInst() (field.EncodedListExecInst, errors.MessageRejectError) {
	var f field.EncodedListExecInst
	err := m.Body.Get(&f)
	return f, err
}

//AllowableOneSidednessPct is a non-required field for NewOrderList.
func (m NewOrderList) AllowableOneSidednessPct() (field.AllowableOneSidednessPct, errors.MessageRejectError) {
	var f field.AllowableOneSidednessPct
	err := m.Body.Get(&f)
	return f, err
}

//AllowableOneSidednessValue is a non-required field for NewOrderList.
func (m NewOrderList) AllowableOneSidednessValue() (field.AllowableOneSidednessValue, errors.MessageRejectError) {
	var f field.AllowableOneSidednessValue
	err := m.Body.Get(&f)
	return f, err
}

//AllowableOneSidednessCurr is a non-required field for NewOrderList.
func (m NewOrderList) AllowableOneSidednessCurr() (field.AllowableOneSidednessCurr, errors.MessageRejectError) {
	var f field.AllowableOneSidednessCurr
	err := m.Body.Get(&f)
	return f, err
}

//TotNoOrders is a required field for NewOrderList.
func (m NewOrderList) TotNoOrders() (field.TotNoOrders, errors.MessageRejectError) {
	var f field.TotNoOrders
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for NewOrderList.
func (m NewOrderList) LastFragment() (field.LastFragment, errors.MessageRejectError) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoOrders is a required field for NewOrderList.
func (m NewOrderList) NoOrders() (field.NoOrders, errors.MessageRejectError) {
	var f field.NoOrders
	err := m.Body.Get(&f)
	return f, err
}

//NoRootPartyIDs is a non-required field for NewOrderList.
func (m NewOrderList) NoRootPartyIDs() (field.NoRootPartyIDs, errors.MessageRejectError) {
	var f field.NoRootPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//ContingencyType is a non-required field for NewOrderList.
func (m NewOrderList) ContingencyType() (field.ContingencyType, errors.MessageRejectError) {
	var f field.ContingencyType
	err := m.Body.Get(&f)
	return f, err
}
