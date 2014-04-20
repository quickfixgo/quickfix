package fix43

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
func (m NewOrderList) ListID() (*field.ListID, errors.MessageRejectError) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from NewOrderList.
func (m NewOrderList) GetListID(f *field.ListID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidID is a non-required field for NewOrderList.
func (m NewOrderList) BidID() (*field.BidID, errors.MessageRejectError) {
	f := new(field.BidID)
	err := m.Body.Get(f)
	return f, err
}

//GetBidID reads a BidID from NewOrderList.
func (m NewOrderList) GetBidID(f *field.BidID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientBidID is a non-required field for NewOrderList.
func (m NewOrderList) ClientBidID() (*field.ClientBidID, errors.MessageRejectError) {
	f := new(field.ClientBidID)
	err := m.Body.Get(f)
	return f, err
}

//GetClientBidID reads a ClientBidID from NewOrderList.
func (m NewOrderList) GetClientBidID(f *field.ClientBidID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProgRptReqs is a non-required field for NewOrderList.
func (m NewOrderList) ProgRptReqs() (*field.ProgRptReqs, errors.MessageRejectError) {
	f := new(field.ProgRptReqs)
	err := m.Body.Get(f)
	return f, err
}

//GetProgRptReqs reads a ProgRptReqs from NewOrderList.
func (m NewOrderList) GetProgRptReqs(f *field.ProgRptReqs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidType is a required field for NewOrderList.
func (m NewOrderList) BidType() (*field.BidType, errors.MessageRejectError) {
	f := new(field.BidType)
	err := m.Body.Get(f)
	return f, err
}

//GetBidType reads a BidType from NewOrderList.
func (m NewOrderList) GetBidType(f *field.BidType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProgPeriodInterval is a non-required field for NewOrderList.
func (m NewOrderList) ProgPeriodInterval() (*field.ProgPeriodInterval, errors.MessageRejectError) {
	f := new(field.ProgPeriodInterval)
	err := m.Body.Get(f)
	return f, err
}

//GetProgPeriodInterval reads a ProgPeriodInterval from NewOrderList.
func (m NewOrderList) GetProgPeriodInterval(f *field.ProgPeriodInterval) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for NewOrderList.
func (m NewOrderList) CancellationRights() (*field.CancellationRights, errors.MessageRejectError) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from NewOrderList.
func (m NewOrderList) GetCancellationRights(f *field.CancellationRights) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for NewOrderList.
func (m NewOrderList) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, errors.MessageRejectError) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from NewOrderList.
func (m NewOrderList) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for NewOrderList.
func (m NewOrderList) RegistID() (*field.RegistID, errors.MessageRejectError) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from NewOrderList.
func (m NewOrderList) GetRegistID(f *field.RegistID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListExecInstType is a non-required field for NewOrderList.
func (m NewOrderList) ListExecInstType() (*field.ListExecInstType, errors.MessageRejectError) {
	f := new(field.ListExecInstType)
	err := m.Body.Get(f)
	return f, err
}

//GetListExecInstType reads a ListExecInstType from NewOrderList.
func (m NewOrderList) GetListExecInstType(f *field.ListExecInstType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListExecInst is a non-required field for NewOrderList.
func (m NewOrderList) ListExecInst() (*field.ListExecInst, errors.MessageRejectError) {
	f := new(field.ListExecInst)
	err := m.Body.Get(f)
	return f, err
}

//GetListExecInst reads a ListExecInst from NewOrderList.
func (m NewOrderList) GetListExecInst(f *field.ListExecInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedListExecInstLen is a non-required field for NewOrderList.
func (m NewOrderList) EncodedListExecInstLen() (*field.EncodedListExecInstLen, errors.MessageRejectError) {
	f := new(field.EncodedListExecInstLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedListExecInstLen reads a EncodedListExecInstLen from NewOrderList.
func (m NewOrderList) GetEncodedListExecInstLen(f *field.EncodedListExecInstLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedListExecInst is a non-required field for NewOrderList.
func (m NewOrderList) EncodedListExecInst() (*field.EncodedListExecInst, errors.MessageRejectError) {
	f := new(field.EncodedListExecInst)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedListExecInst reads a EncodedListExecInst from NewOrderList.
func (m NewOrderList) GetEncodedListExecInst(f *field.EncodedListExecInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoOrders is a required field for NewOrderList.
func (m NewOrderList) TotNoOrders() (*field.TotNoOrders, errors.MessageRejectError) {
	f := new(field.TotNoOrders)
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoOrders reads a TotNoOrders from NewOrderList.
func (m NewOrderList) GetTotNoOrders(f *field.TotNoOrders) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a required field for NewOrderList.
func (m NewOrderList) NoOrders() (*field.NoOrders, errors.MessageRejectError) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from NewOrderList.
func (m NewOrderList) GetNoOrders(f *field.NoOrders) errors.MessageRejectError {
	return m.Body.Get(f)
}
