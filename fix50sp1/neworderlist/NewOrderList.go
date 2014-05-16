//Package neworderlist msg type = E.
package neworderlist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a NewOrderList wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//ListID is a required field for NewOrderList.
func (m Message) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from NewOrderList.
func (m Message) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidID is a non-required field for NewOrderList.
func (m Message) BidID() (*field.BidIDField, errors.MessageRejectError) {
	f := &field.BidIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidID reads a BidID from NewOrderList.
func (m Message) GetBidID(f *field.BidIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientBidID is a non-required field for NewOrderList.
func (m Message) ClientBidID() (*field.ClientBidIDField, errors.MessageRejectError) {
	f := &field.ClientBidIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientBidID reads a ClientBidID from NewOrderList.
func (m Message) GetClientBidID(f *field.ClientBidIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProgRptReqs is a non-required field for NewOrderList.
func (m Message) ProgRptReqs() (*field.ProgRptReqsField, errors.MessageRejectError) {
	f := &field.ProgRptReqsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProgRptReqs reads a ProgRptReqs from NewOrderList.
func (m Message) GetProgRptReqs(f *field.ProgRptReqsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidType is a required field for NewOrderList.
func (m Message) BidType() (*field.BidTypeField, errors.MessageRejectError) {
	f := &field.BidTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidType reads a BidType from NewOrderList.
func (m Message) GetBidType(f *field.BidTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProgPeriodInterval is a non-required field for NewOrderList.
func (m Message) ProgPeriodInterval() (*field.ProgPeriodIntervalField, errors.MessageRejectError) {
	f := &field.ProgPeriodIntervalField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProgPeriodInterval reads a ProgPeriodInterval from NewOrderList.
func (m Message) GetProgPeriodInterval(f *field.ProgPeriodIntervalField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for NewOrderList.
func (m Message) CancellationRights() (*field.CancellationRightsField, errors.MessageRejectError) {
	f := &field.CancellationRightsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from NewOrderList.
func (m Message) GetCancellationRights(f *field.CancellationRightsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for NewOrderList.
func (m Message) MoneyLaunderingStatus() (*field.MoneyLaunderingStatusField, errors.MessageRejectError) {
	f := &field.MoneyLaunderingStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from NewOrderList.
func (m Message) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for NewOrderList.
func (m Message) RegistID() (*field.RegistIDField, errors.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from NewOrderList.
func (m Message) GetRegistID(f *field.RegistIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListExecInstType is a non-required field for NewOrderList.
func (m Message) ListExecInstType() (*field.ListExecInstTypeField, errors.MessageRejectError) {
	f := &field.ListExecInstTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListExecInstType reads a ListExecInstType from NewOrderList.
func (m Message) GetListExecInstType(f *field.ListExecInstTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListExecInst is a non-required field for NewOrderList.
func (m Message) ListExecInst() (*field.ListExecInstField, errors.MessageRejectError) {
	f := &field.ListExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListExecInst reads a ListExecInst from NewOrderList.
func (m Message) GetListExecInst(f *field.ListExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedListExecInstLen is a non-required field for NewOrderList.
func (m Message) EncodedListExecInstLen() (*field.EncodedListExecInstLenField, errors.MessageRejectError) {
	f := &field.EncodedListExecInstLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedListExecInstLen reads a EncodedListExecInstLen from NewOrderList.
func (m Message) GetEncodedListExecInstLen(f *field.EncodedListExecInstLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedListExecInst is a non-required field for NewOrderList.
func (m Message) EncodedListExecInst() (*field.EncodedListExecInstField, errors.MessageRejectError) {
	f := &field.EncodedListExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedListExecInst reads a EncodedListExecInst from NewOrderList.
func (m Message) GetEncodedListExecInst(f *field.EncodedListExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllowableOneSidednessPct is a non-required field for NewOrderList.
func (m Message) AllowableOneSidednessPct() (*field.AllowableOneSidednessPctField, errors.MessageRejectError) {
	f := &field.AllowableOneSidednessPctField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllowableOneSidednessPct reads a AllowableOneSidednessPct from NewOrderList.
func (m Message) GetAllowableOneSidednessPct(f *field.AllowableOneSidednessPctField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllowableOneSidednessValue is a non-required field for NewOrderList.
func (m Message) AllowableOneSidednessValue() (*field.AllowableOneSidednessValueField, errors.MessageRejectError) {
	f := &field.AllowableOneSidednessValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllowableOneSidednessValue reads a AllowableOneSidednessValue from NewOrderList.
func (m Message) GetAllowableOneSidednessValue(f *field.AllowableOneSidednessValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllowableOneSidednessCurr is a non-required field for NewOrderList.
func (m Message) AllowableOneSidednessCurr() (*field.AllowableOneSidednessCurrField, errors.MessageRejectError) {
	f := &field.AllowableOneSidednessCurrField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllowableOneSidednessCurr reads a AllowableOneSidednessCurr from NewOrderList.
func (m Message) GetAllowableOneSidednessCurr(f *field.AllowableOneSidednessCurrField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoOrders is a required field for NewOrderList.
func (m Message) TotNoOrders() (*field.TotNoOrdersField, errors.MessageRejectError) {
	f := &field.TotNoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoOrders reads a TotNoOrders from NewOrderList.
func (m Message) GetTotNoOrders(f *field.TotNoOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for NewOrderList.
func (m Message) LastFragment() (*field.LastFragmentField, errors.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from NewOrderList.
func (m Message) GetLastFragment(f *field.LastFragmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a required field for NewOrderList.
func (m Message) NoOrders() (*field.NoOrdersField, errors.MessageRejectError) {
	f := &field.NoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from NewOrderList.
func (m Message) GetNoOrders(f *field.NoOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRootPartyIDs is a non-required field for NewOrderList.
func (m Message) NoRootPartyIDs() (*field.NoRootPartyIDsField, errors.MessageRejectError) {
	f := &field.NoRootPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRootPartyIDs reads a NoRootPartyIDs from NewOrderList.
func (m Message) GetNoRootPartyIDs(f *field.NoRootPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContingencyType is a non-required field for NewOrderList.
func (m Message) ContingencyType() (*field.ContingencyTypeField, errors.MessageRejectError) {
	f := &field.ContingencyTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContingencyType reads a ContingencyType from NewOrderList.
func (m Message) GetContingencyType(f *field.ContingencyTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds NewOrderList messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for NewOrderList.
func Builder(
	listid *field.ListIDField,
	bidtype *field.BidTypeField,
	totnoorders *field.TotNoOrdersField,
	noorders *field.NoOrdersField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("E"))
	builder.Body().Set(listid)
	builder.Body().Set(bidtype)
	builder.Body().Set(totnoorders)
	builder.Body().Set(noorders)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "E", r
}
