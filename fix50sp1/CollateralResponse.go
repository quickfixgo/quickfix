package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type CollateralResponse struct {
	quickfixgo.Message
}

func (m *CollateralResponse) CollRespID() (*field.CollRespID, error) {
	f := new(field.CollRespID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) CollAsgnID() (*field.CollAsgnID, error) {
	f := new(field.CollAsgnID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) CollReqID() (*field.CollReqID, error) {
	f := new(field.CollReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) CollAsgnReason() (*field.CollAsgnReason, error) {
	f := new(field.CollAsgnReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) CollAsgnTransType() (*field.CollAsgnTransType, error) {
	f := new(field.CollAsgnTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) CollAsgnRespType() (*field.CollAsgnRespType, error) {
	f := new(field.CollAsgnRespType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) CollAsgnRejectReason() (*field.CollAsgnRejectReason, error) {
	f := new(field.CollAsgnRejectReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) Quantity() (*field.Quantity, error) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) MarginExcess() (*field.MarginExcess, error) {
	f := new(field.MarginExcess)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) TotalNetValue() (*field.TotalNetValue, error) {
	f := new(field.TotalNetValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) CashOutstanding() (*field.CashOutstanding, error) {
	f := new(field.CashOutstanding)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) AccruedInterestAmt() (*field.AccruedInterestAmt, error) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) EndAccruedInterestAmt() (*field.EndAccruedInterestAmt, error) {
	f := new(field.EndAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) StartCash() (*field.StartCash, error) {
	f := new(field.StartCash)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) EndCash() (*field.EndCash, error) {
	f := new(field.EndCash)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) CollApplType() (*field.CollApplType, error) {
	f := new(field.CollApplType)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) FinancialStatus() (*field.FinancialStatus, error) {
	f := new(field.FinancialStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *CollateralResponse) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
