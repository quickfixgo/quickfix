package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//OrderCancelReject msg type = 9.
type OrderCancelReject struct {
	message.Message
}

//OrderCancelRejectBuilder builds OrderCancelReject messages.
type OrderCancelRejectBuilder struct {
	message.MessageBuilder
}

//CreateOrderCancelRejectBuilder returns an initialized OrderCancelRejectBuilder with specified required fields.
func CreateOrderCancelRejectBuilder(
	orderid *field.OrderIDField,
	clordid *field.ClOrdIDField,
	ordstatus *field.OrdStatusField,
	cxlrejresponseto *field.CxlRejResponseToField) OrderCancelRejectBuilder {
	var builder OrderCancelRejectBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("9"))
	builder.Body().Set(orderid)
	builder.Body().Set(clordid)
	builder.Body().Set(ordstatus)
	builder.Body().Set(cxlrejresponseto)
	return builder
}

//OrderID is a required field for OrderCancelReject.
func (m OrderCancelReject) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from OrderCancelReject.
func (m OrderCancelReject) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for OrderCancelReject.
func (m OrderCancelReject) SecondaryOrderID() (*field.SecondaryOrderIDField, errors.MessageRejectError) {
	f := &field.SecondaryOrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from OrderCancelReject.
func (m OrderCancelReject) GetSecondaryOrderID(f *field.SecondaryOrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for OrderCancelReject.
func (m OrderCancelReject) SecondaryClOrdID() (*field.SecondaryClOrdIDField, errors.MessageRejectError) {
	f := &field.SecondaryClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from OrderCancelReject.
func (m OrderCancelReject) GetSecondaryClOrdID(f *field.SecondaryClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for OrderCancelReject.
func (m OrderCancelReject) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderCancelReject.
func (m OrderCancelReject) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for OrderCancelReject.
func (m OrderCancelReject) ClOrdLinkID() (*field.ClOrdLinkIDField, errors.MessageRejectError) {
	f := &field.ClOrdLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from OrderCancelReject.
func (m OrderCancelReject) GetClOrdLinkID(f *field.ClOrdLinkIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a non-required field for OrderCancelReject.
func (m OrderCancelReject) OrigClOrdID() (*field.OrigClOrdIDField, errors.MessageRejectError) {
	f := &field.OrigClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from OrderCancelReject.
func (m OrderCancelReject) GetOrigClOrdID(f *field.OrigClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdStatus is a required field for OrderCancelReject.
func (m OrderCancelReject) OrdStatus() (*field.OrdStatusField, errors.MessageRejectError) {
	f := &field.OrdStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdStatus reads a OrdStatus from OrderCancelReject.
func (m OrderCancelReject) GetOrdStatus(f *field.OrdStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//WorkingIndicator is a non-required field for OrderCancelReject.
func (m OrderCancelReject) WorkingIndicator() (*field.WorkingIndicatorField, errors.MessageRejectError) {
	f := &field.WorkingIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetWorkingIndicator reads a WorkingIndicator from OrderCancelReject.
func (m OrderCancelReject) GetWorkingIndicator(f *field.WorkingIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigOrdModTime is a non-required field for OrderCancelReject.
func (m OrderCancelReject) OrigOrdModTime() (*field.OrigOrdModTimeField, errors.MessageRejectError) {
	f := &field.OrigOrdModTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigOrdModTime reads a OrigOrdModTime from OrderCancelReject.
func (m OrderCancelReject) GetOrigOrdModTime(f *field.OrigOrdModTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListID is a non-required field for OrderCancelReject.
func (m OrderCancelReject) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from OrderCancelReject.
func (m OrderCancelReject) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for OrderCancelReject.
func (m OrderCancelReject) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from OrderCancelReject.
func (m OrderCancelReject) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for OrderCancelReject.
func (m OrderCancelReject) AcctIDSource() (*field.AcctIDSourceField, errors.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from OrderCancelReject.
func (m OrderCancelReject) GetAcctIDSource(f *field.AcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for OrderCancelReject.
func (m OrderCancelReject) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from OrderCancelReject.
func (m OrderCancelReject) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for OrderCancelReject.
func (m OrderCancelReject) TradeOriginationDate() (*field.TradeOriginationDateField, errors.MessageRejectError) {
	f := &field.TradeOriginationDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from OrderCancelReject.
func (m OrderCancelReject) GetTradeOriginationDate(f *field.TradeOriginationDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for OrderCancelReject.
func (m OrderCancelReject) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from OrderCancelReject.
func (m OrderCancelReject) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for OrderCancelReject.
func (m OrderCancelReject) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from OrderCancelReject.
func (m OrderCancelReject) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CxlRejResponseTo is a required field for OrderCancelReject.
func (m OrderCancelReject) CxlRejResponseTo() (*field.CxlRejResponseToField, errors.MessageRejectError) {
	f := &field.CxlRejResponseToField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCxlRejResponseTo reads a CxlRejResponseTo from OrderCancelReject.
func (m OrderCancelReject) GetCxlRejResponseTo(f *field.CxlRejResponseToField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CxlRejReason is a non-required field for OrderCancelReject.
func (m OrderCancelReject) CxlRejReason() (*field.CxlRejReasonField, errors.MessageRejectError) {
	f := &field.CxlRejReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCxlRejReason reads a CxlRejReason from OrderCancelReject.
func (m OrderCancelReject) GetCxlRejReason(f *field.CxlRejReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for OrderCancelReject.
func (m OrderCancelReject) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from OrderCancelReject.
func (m OrderCancelReject) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for OrderCancelReject.
func (m OrderCancelReject) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from OrderCancelReject.
func (m OrderCancelReject) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for OrderCancelReject.
func (m OrderCancelReject) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from OrderCancelReject.
func (m OrderCancelReject) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
