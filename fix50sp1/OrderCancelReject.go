package fix50sp1

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

//NewOrderCancelRejectBuilder returns an initialized OrderCancelRejectBuilder with specified required fields.
func NewOrderCancelRejectBuilder(
	orderid field.OrderID,
	clordid field.ClOrdID,
	ordstatus field.OrdStatus,
	cxlrejresponseto field.CxlRejResponseTo) *OrderCancelRejectBuilder {
	builder := new(OrderCancelRejectBuilder)
	builder.Body.Set(orderid)
	builder.Body.Set(clordid)
	builder.Body.Set(ordstatus)
	builder.Body.Set(cxlrejresponseto)
	return builder
}

//OrderID is a required field for OrderCancelReject.
func (m *OrderCancelReject) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryOrderID is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryClOrdID is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a required field for OrderCancelReject.
func (m *OrderCancelReject) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdLinkID is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) ClOrdLinkID() (*field.ClOrdLinkID, error) {
	f := new(field.ClOrdLinkID)
	err := m.Body.Get(f)
	return f, err
}

//OrigClOrdID is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) OrigClOrdID() (*field.OrigClOrdID, error) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//OrdStatus is a required field for OrderCancelReject.
func (m *OrderCancelReject) OrdStatus() (*field.OrdStatus, error) {
	f := new(field.OrdStatus)
	err := m.Body.Get(f)
	return f, err
}

//WorkingIndicator is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) WorkingIndicator() (*field.WorkingIndicator, error) {
	f := new(field.WorkingIndicator)
	err := m.Body.Get(f)
	return f, err
}

//OrigOrdModTime is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) OrigOrdModTime() (*field.OrigOrdModTime, error) {
	f := new(field.OrigOrdModTime)
	err := m.Body.Get(f)
	return f, err
}

//ListID is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//AcctIDSource is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//AccountType is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//TradeOriginationDate is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) TradeOriginationDate() (*field.TradeOriginationDate, error) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//CxlRejResponseTo is a required field for OrderCancelReject.
func (m *OrderCancelReject) CxlRejResponseTo() (*field.CxlRejResponseTo, error) {
	f := new(field.CxlRejResponseTo)
	err := m.Body.Get(f)
	return f, err
}

//CxlRejReason is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) CxlRejReason() (*field.CxlRejReason, error) {
	f := new(field.CxlRejReason)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for OrderCancelReject.
func (m *OrderCancelReject) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
