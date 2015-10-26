//Package neworderlist msg type = E.
package neworderlist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a NewOrderList wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//ListID is a required field for NewOrderList.
func (m Message) ListID() (*field.ListIDField, quickfix.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from NewOrderList.
func (m Message) GetListID(f *field.ListIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BidID is a non-required field for NewOrderList.
func (m Message) BidID() (*field.BidIDField, quickfix.MessageRejectError) {
	f := &field.BidIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidID reads a BidID from NewOrderList.
func (m Message) GetBidID(f *field.BidIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClientBidID is a non-required field for NewOrderList.
func (m Message) ClientBidID() (*field.ClientBidIDField, quickfix.MessageRejectError) {
	f := &field.ClientBidIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientBidID reads a ClientBidID from NewOrderList.
func (m Message) GetClientBidID(f *field.ClientBidIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ProgRptReqs is a non-required field for NewOrderList.
func (m Message) ProgRptReqs() (*field.ProgRptReqsField, quickfix.MessageRejectError) {
	f := &field.ProgRptReqsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProgRptReqs reads a ProgRptReqs from NewOrderList.
func (m Message) GetProgRptReqs(f *field.ProgRptReqsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BidType is a required field for NewOrderList.
func (m Message) BidType() (*field.BidTypeField, quickfix.MessageRejectError) {
	f := &field.BidTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBidType reads a BidType from NewOrderList.
func (m Message) GetBidType(f *field.BidTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ProgPeriodInterval is a non-required field for NewOrderList.
func (m Message) ProgPeriodInterval() (*field.ProgPeriodIntervalField, quickfix.MessageRejectError) {
	f := &field.ProgPeriodIntervalField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProgPeriodInterval reads a ProgPeriodInterval from NewOrderList.
func (m Message) GetProgPeriodInterval(f *field.ProgPeriodIntervalField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for NewOrderList.
func (m Message) CancellationRights() (*field.CancellationRightsField, quickfix.MessageRejectError) {
	f := &field.CancellationRightsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from NewOrderList.
func (m Message) GetCancellationRights(f *field.CancellationRightsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for NewOrderList.
func (m Message) MoneyLaunderingStatus() (*field.MoneyLaunderingStatusField, quickfix.MessageRejectError) {
	f := &field.MoneyLaunderingStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from NewOrderList.
func (m Message) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for NewOrderList.
func (m Message) RegistID() (*field.RegistIDField, quickfix.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from NewOrderList.
func (m Message) GetRegistID(f *field.RegistIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ListExecInstType is a non-required field for NewOrderList.
func (m Message) ListExecInstType() (*field.ListExecInstTypeField, quickfix.MessageRejectError) {
	f := &field.ListExecInstTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListExecInstType reads a ListExecInstType from NewOrderList.
func (m Message) GetListExecInstType(f *field.ListExecInstTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ListExecInst is a non-required field for NewOrderList.
func (m Message) ListExecInst() (*field.ListExecInstField, quickfix.MessageRejectError) {
	f := &field.ListExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListExecInst reads a ListExecInst from NewOrderList.
func (m Message) GetListExecInst(f *field.ListExecInstField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedListExecInstLen is a non-required field for NewOrderList.
func (m Message) EncodedListExecInstLen() (*field.EncodedListExecInstLenField, quickfix.MessageRejectError) {
	f := &field.EncodedListExecInstLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedListExecInstLen reads a EncodedListExecInstLen from NewOrderList.
func (m Message) GetEncodedListExecInstLen(f *field.EncodedListExecInstLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedListExecInst is a non-required field for NewOrderList.
func (m Message) EncodedListExecInst() (*field.EncodedListExecInstField, quickfix.MessageRejectError) {
	f := &field.EncodedListExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedListExecInst reads a EncodedListExecInst from NewOrderList.
func (m Message) GetEncodedListExecInst(f *field.EncodedListExecInstField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllowableOneSidednessPct is a non-required field for NewOrderList.
func (m Message) AllowableOneSidednessPct() (*field.AllowableOneSidednessPctField, quickfix.MessageRejectError) {
	f := &field.AllowableOneSidednessPctField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllowableOneSidednessPct reads a AllowableOneSidednessPct from NewOrderList.
func (m Message) GetAllowableOneSidednessPct(f *field.AllowableOneSidednessPctField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllowableOneSidednessValue is a non-required field for NewOrderList.
func (m Message) AllowableOneSidednessValue() (*field.AllowableOneSidednessValueField, quickfix.MessageRejectError) {
	f := &field.AllowableOneSidednessValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllowableOneSidednessValue reads a AllowableOneSidednessValue from NewOrderList.
func (m Message) GetAllowableOneSidednessValue(f *field.AllowableOneSidednessValueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllowableOneSidednessCurr is a non-required field for NewOrderList.
func (m Message) AllowableOneSidednessCurr() (*field.AllowableOneSidednessCurrField, quickfix.MessageRejectError) {
	f := &field.AllowableOneSidednessCurrField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllowableOneSidednessCurr reads a AllowableOneSidednessCurr from NewOrderList.
func (m Message) GetAllowableOneSidednessCurr(f *field.AllowableOneSidednessCurrField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoOrders is a required field for NewOrderList.
func (m Message) TotNoOrders() (*field.TotNoOrdersField, quickfix.MessageRejectError) {
	f := &field.TotNoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoOrders reads a TotNoOrders from NewOrderList.
func (m Message) GetTotNoOrders(f *field.TotNoOrdersField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for NewOrderList.
func (m Message) LastFragment() (*field.LastFragmentField, quickfix.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from NewOrderList.
func (m Message) GetLastFragment(f *field.LastFragmentField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a required field for NewOrderList.
func (m Message) NoOrders() (*field.NoOrdersField, quickfix.MessageRejectError) {
	f := &field.NoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from NewOrderList.
func (m Message) GetNoOrders(f *field.NoOrdersField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for NewOrderList.
func New(
	listid *field.ListIDField,
	bidtype *field.BidTypeField,
	totnoorders *field.TotNoOrdersField,
	noorders *field.NoOrdersField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIX44))
	builder.Header.Set(field.NewMsgType("E"))
	builder.Body.Set(listid)
	builder.Body.Set(bidtype)
	builder.Body.Set(totnoorders)
	builder.Body.Set(noorders)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.BeginStringFIX44, "E", r
}
