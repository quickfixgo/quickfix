package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	listid *field.ListIDField,
	bidtype *field.BidTypeField,
	totnoorders *field.TotNoOrdersField,
	noorders *field.NoOrdersField) NewOrderListBuilder {
	var builder NewOrderListBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("E"))
	builder.Body().Set(listid)
	builder.Body().Set(bidtype)
	builder.Body().Set(totnoorders)
	builder.Body().Set(noorders)
	return builder
}

//ListID is a required field for NewOrderList.
func (m NewOrderList) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from NewOrderList.
func (m NewOrderList) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidID is a non-required field for NewOrderList.
func (m NewOrderList) BidID() (*field.BidIDField, errors.MessageRejectError) {
	f := &field.BidIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidID reads a BidID from NewOrderList.
func (m NewOrderList) GetBidID(f *field.BidIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientBidID is a non-required field for NewOrderList.
func (m NewOrderList) ClientBidID() (*field.ClientBidIDField, errors.MessageRejectError) {
	f := &field.ClientBidIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientBidID reads a ClientBidID from NewOrderList.
func (m NewOrderList) GetClientBidID(f *field.ClientBidIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProgRptReqs is a non-required field for NewOrderList.
func (m NewOrderList) ProgRptReqs() (*field.ProgRptReqsField, errors.MessageRejectError) {
	f := &field.ProgRptReqsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProgRptReqs reads a ProgRptReqs from NewOrderList.
func (m NewOrderList) GetProgRptReqs(f *field.ProgRptReqsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidType is a required field for NewOrderList.
func (m NewOrderList) BidType() (*field.BidTypeField, errors.MessageRejectError) {
	f := &field.BidTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidType reads a BidType from NewOrderList.
func (m NewOrderList) GetBidType(f *field.BidTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProgPeriodInterval is a non-required field for NewOrderList.
func (m NewOrderList) ProgPeriodInterval() (*field.ProgPeriodIntervalField, errors.MessageRejectError) {
	f := &field.ProgPeriodIntervalField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProgPeriodInterval reads a ProgPeriodInterval from NewOrderList.
func (m NewOrderList) GetProgPeriodInterval(f *field.ProgPeriodIntervalField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for NewOrderList.
func (m NewOrderList) CancellationRights() (*field.CancellationRightsField, errors.MessageRejectError) {
	f := &field.CancellationRightsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from NewOrderList.
func (m NewOrderList) GetCancellationRights(f *field.CancellationRightsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for NewOrderList.
func (m NewOrderList) MoneyLaunderingStatus() (*field.MoneyLaunderingStatusField, errors.MessageRejectError) {
	f := &field.MoneyLaunderingStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from NewOrderList.
func (m NewOrderList) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for NewOrderList.
func (m NewOrderList) RegistID() (*field.RegistIDField, errors.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from NewOrderList.
func (m NewOrderList) GetRegistID(f *field.RegistIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListExecInstType is a non-required field for NewOrderList.
func (m NewOrderList) ListExecInstType() (*field.ListExecInstTypeField, errors.MessageRejectError) {
	f := &field.ListExecInstTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListExecInstType reads a ListExecInstType from NewOrderList.
func (m NewOrderList) GetListExecInstType(f *field.ListExecInstTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListExecInst is a non-required field for NewOrderList.
func (m NewOrderList) ListExecInst() (*field.ListExecInstField, errors.MessageRejectError) {
	f := &field.ListExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListExecInst reads a ListExecInst from NewOrderList.
func (m NewOrderList) GetListExecInst(f *field.ListExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedListExecInstLen is a non-required field for NewOrderList.
func (m NewOrderList) EncodedListExecInstLen() (*field.EncodedListExecInstLenField, errors.MessageRejectError) {
	f := &field.EncodedListExecInstLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedListExecInstLen reads a EncodedListExecInstLen from NewOrderList.
func (m NewOrderList) GetEncodedListExecInstLen(f *field.EncodedListExecInstLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedListExecInst is a non-required field for NewOrderList.
func (m NewOrderList) EncodedListExecInst() (*field.EncodedListExecInstField, errors.MessageRejectError) {
	f := &field.EncodedListExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedListExecInst reads a EncodedListExecInst from NewOrderList.
func (m NewOrderList) GetEncodedListExecInst(f *field.EncodedListExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllowableOneSidednessPct is a non-required field for NewOrderList.
func (m NewOrderList) AllowableOneSidednessPct() (*field.AllowableOneSidednessPctField, errors.MessageRejectError) {
	f := &field.AllowableOneSidednessPctField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllowableOneSidednessPct reads a AllowableOneSidednessPct from NewOrderList.
func (m NewOrderList) GetAllowableOneSidednessPct(f *field.AllowableOneSidednessPctField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllowableOneSidednessValue is a non-required field for NewOrderList.
func (m NewOrderList) AllowableOneSidednessValue() (*field.AllowableOneSidednessValueField, errors.MessageRejectError) {
	f := &field.AllowableOneSidednessValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllowableOneSidednessValue reads a AllowableOneSidednessValue from NewOrderList.
func (m NewOrderList) GetAllowableOneSidednessValue(f *field.AllowableOneSidednessValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllowableOneSidednessCurr is a non-required field for NewOrderList.
func (m NewOrderList) AllowableOneSidednessCurr() (*field.AllowableOneSidednessCurrField, errors.MessageRejectError) {
	f := &field.AllowableOneSidednessCurrField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllowableOneSidednessCurr reads a AllowableOneSidednessCurr from NewOrderList.
func (m NewOrderList) GetAllowableOneSidednessCurr(f *field.AllowableOneSidednessCurrField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoOrders is a required field for NewOrderList.
func (m NewOrderList) TotNoOrders() (*field.TotNoOrdersField, errors.MessageRejectError) {
	f := &field.TotNoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoOrders reads a TotNoOrders from NewOrderList.
func (m NewOrderList) GetTotNoOrders(f *field.TotNoOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for NewOrderList.
func (m NewOrderList) LastFragment() (*field.LastFragmentField, errors.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from NewOrderList.
func (m NewOrderList) GetLastFragment(f *field.LastFragmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a required field for NewOrderList.
func (m NewOrderList) NoOrders() (*field.NoOrdersField, errors.MessageRejectError) {
	f := &field.NoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from NewOrderList.
func (m NewOrderList) GetNoOrders(f *field.NoOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRootPartyIDs is a non-required field for NewOrderList.
func (m NewOrderList) NoRootPartyIDs() (*field.NoRootPartyIDsField, errors.MessageRejectError) {
	f := &field.NoRootPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRootPartyIDs reads a NoRootPartyIDs from NewOrderList.
func (m NewOrderList) GetNoRootPartyIDs(f *field.NoRootPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContingencyType is a non-required field for NewOrderList.
func (m NewOrderList) ContingencyType() (*field.ContingencyTypeField, errors.MessageRejectError) {
	f := &field.ContingencyTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContingencyType reads a ContingencyType from NewOrderList.
func (m NewOrderList) GetContingencyType(f *field.ContingencyTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}
