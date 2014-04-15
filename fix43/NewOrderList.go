package fix43

import (
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
func (m NewOrderList) ListID() (field.ListID, error) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//BidID is a non-required field for NewOrderList.
func (m NewOrderList) BidID() (field.BidID, error) {
	var f field.BidID
	err := m.Body.Get(&f)
	return f, err
}

//ClientBidID is a non-required field for NewOrderList.
func (m NewOrderList) ClientBidID() (field.ClientBidID, error) {
	var f field.ClientBidID
	err := m.Body.Get(&f)
	return f, err
}

//ProgRptReqs is a non-required field for NewOrderList.
func (m NewOrderList) ProgRptReqs() (field.ProgRptReqs, error) {
	var f field.ProgRptReqs
	err := m.Body.Get(&f)
	return f, err
}

//BidType is a required field for NewOrderList.
func (m NewOrderList) BidType() (field.BidType, error) {
	var f field.BidType
	err := m.Body.Get(&f)
	return f, err
}

//ProgPeriodInterval is a non-required field for NewOrderList.
func (m NewOrderList) ProgPeriodInterval() (field.ProgPeriodInterval, error) {
	var f field.ProgPeriodInterval
	err := m.Body.Get(&f)
	return f, err
}

//CancellationRights is a non-required field for NewOrderList.
func (m NewOrderList) CancellationRights() (field.CancellationRights, error) {
	var f field.CancellationRights
	err := m.Body.Get(&f)
	return f, err
}

//MoneyLaunderingStatus is a non-required field for NewOrderList.
func (m NewOrderList) MoneyLaunderingStatus() (field.MoneyLaunderingStatus, error) {
	var f field.MoneyLaunderingStatus
	err := m.Body.Get(&f)
	return f, err
}

//RegistID is a non-required field for NewOrderList.
func (m NewOrderList) RegistID() (field.RegistID, error) {
	var f field.RegistID
	err := m.Body.Get(&f)
	return f, err
}

//ListExecInstType is a non-required field for NewOrderList.
func (m NewOrderList) ListExecInstType() (field.ListExecInstType, error) {
	var f field.ListExecInstType
	err := m.Body.Get(&f)
	return f, err
}

//ListExecInst is a non-required field for NewOrderList.
func (m NewOrderList) ListExecInst() (field.ListExecInst, error) {
	var f field.ListExecInst
	err := m.Body.Get(&f)
	return f, err
}

//EncodedListExecInstLen is a non-required field for NewOrderList.
func (m NewOrderList) EncodedListExecInstLen() (field.EncodedListExecInstLen, error) {
	var f field.EncodedListExecInstLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedListExecInst is a non-required field for NewOrderList.
func (m NewOrderList) EncodedListExecInst() (field.EncodedListExecInst, error) {
	var f field.EncodedListExecInst
	err := m.Body.Get(&f)
	return f, err
}

//TotNoOrders is a required field for NewOrderList.
func (m NewOrderList) TotNoOrders() (field.TotNoOrders, error) {
	var f field.TotNoOrders
	err := m.Body.Get(&f)
	return f, err
}

//NoOrders is a required field for NewOrderList.
func (m NewOrderList) NoOrders() (field.NoOrders, error) {
	var f field.NoOrders
	err := m.Body.Get(&f)
	return f, err
}
