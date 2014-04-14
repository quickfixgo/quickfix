package fix42

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderStatusRequest msg type = H.
type OrderStatusRequest struct {
	message.Message
}

//OrderStatusRequestBuilder builds OrderStatusRequest messages.
type OrderStatusRequestBuilder struct {
	message.MessageBuilder
}

//NewOrderStatusRequestBuilder returns an initialized OrderStatusRequestBuilder with specified required fields.
func NewOrderStatusRequestBuilder(
	clordid field.ClOrdID,
	symbol field.Symbol,
	side field.Side) *OrderStatusRequestBuilder {
	builder := new(OrderStatusRequestBuilder)
	builder.Body.Set(clordid)
	builder.Body.Set(symbol)
	builder.Body.Set(side)
	return builder
}

//OrderID is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a required field for OrderStatusRequest.
func (m *OrderStatusRequest) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//ClientID is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) ClientID() (*field.ClientID, error) {
	f := new(field.ClientID)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//ExecBroker is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) ExecBroker() (*field.ExecBroker, error) {
	f := new(field.ExecBroker)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a required field for OrderStatusRequest.
func (m *OrderStatusRequest) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//IDSource is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) IDSource() (*field.IDSource, error) {
	f := new(field.IDSource)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDay is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) MaturityDay() (*field.MaturityDay, error) {
	f := new(field.MaturityDay)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for OrderStatusRequest.
func (m *OrderStatusRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Side is a required field for OrderStatusRequest.
func (m *OrderStatusRequest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
