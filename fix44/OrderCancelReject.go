package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
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
	orderid field.OrderID,
	clordid field.ClOrdID,
	origclordid field.OrigClOrdID,
	ordstatus field.OrdStatus,
	cxlrejresponseto field.CxlRejResponseTo) OrderCancelRejectBuilder {
	var builder OrderCancelRejectBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(orderid)
	builder.Body.Set(clordid)
	builder.Body.Set(origclordid)
	builder.Body.Set(ordstatus)
	builder.Body.Set(cxlrejresponseto)
	return builder
}

//OrderID is a required field for OrderCancelReject.
func (m OrderCancelReject) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryOrderID is a non-required field for OrderCancelReject.
func (m OrderCancelReject) SecondaryOrderID() (field.SecondaryOrderID, error) {
	var f field.SecondaryOrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryClOrdID is a non-required field for OrderCancelReject.
func (m OrderCancelReject) SecondaryClOrdID() (field.SecondaryClOrdID, error) {
	var f field.SecondaryClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a required field for OrderCancelReject.
func (m OrderCancelReject) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdLinkID is a non-required field for OrderCancelReject.
func (m OrderCancelReject) ClOrdLinkID() (field.ClOrdLinkID, error) {
	var f field.ClOrdLinkID
	err := m.Body.Get(&f)
	return f, err
}

//OrigClOrdID is a required field for OrderCancelReject.
func (m OrderCancelReject) OrigClOrdID() (field.OrigClOrdID, error) {
	var f field.OrigClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrdStatus is a required field for OrderCancelReject.
func (m OrderCancelReject) OrdStatus() (field.OrdStatus, error) {
	var f field.OrdStatus
	err := m.Body.Get(&f)
	return f, err
}

//WorkingIndicator is a non-required field for OrderCancelReject.
func (m OrderCancelReject) WorkingIndicator() (field.WorkingIndicator, error) {
	var f field.WorkingIndicator
	err := m.Body.Get(&f)
	return f, err
}

//OrigOrdModTime is a non-required field for OrderCancelReject.
func (m OrderCancelReject) OrigOrdModTime() (field.OrigOrdModTime, error) {
	var f field.OrigOrdModTime
	err := m.Body.Get(&f)
	return f, err
}

//ListID is a non-required field for OrderCancelReject.
func (m OrderCancelReject) ListID() (field.ListID, error) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for OrderCancelReject.
func (m OrderCancelReject) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for OrderCancelReject.
func (m OrderCancelReject) AcctIDSource() (field.AcctIDSource, error) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for OrderCancelReject.
func (m OrderCancelReject) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//TradeOriginationDate is a non-required field for OrderCancelReject.
func (m OrderCancelReject) TradeOriginationDate() (field.TradeOriginationDate, error) {
	var f field.TradeOriginationDate
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a non-required field for OrderCancelReject.
func (m OrderCancelReject) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for OrderCancelReject.
func (m OrderCancelReject) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//CxlRejResponseTo is a required field for OrderCancelReject.
func (m OrderCancelReject) CxlRejResponseTo() (field.CxlRejResponseTo, error) {
	var f field.CxlRejResponseTo
	err := m.Body.Get(&f)
	return f, err
}

//CxlRejReason is a non-required field for OrderCancelReject.
func (m OrderCancelReject) CxlRejReason() (field.CxlRejReason, error) {
	var f field.CxlRejReason
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for OrderCancelReject.
func (m OrderCancelReject) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for OrderCancelReject.
func (m OrderCancelReject) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for OrderCancelReject.
func (m OrderCancelReject) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
