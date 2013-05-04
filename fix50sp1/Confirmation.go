package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type Confirmation struct {
	quickfixgo.Message
}

func (m *Confirmation) ConfirmID() (*field.ConfirmID, error) {
	f := new(field.ConfirmID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) ConfirmRefID() (*field.ConfirmRefID, error) {
	f := new(field.ConfirmRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) ConfirmReqID() (*field.ConfirmReqID, error) {
	f := new(field.ConfirmReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) ConfirmTransType() (*field.ConfirmTransType, error) {
	f := new(field.ConfirmTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) ConfirmType() (*field.ConfirmType, error) {
	f := new(field.ConfirmType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) CopyMsgIndicator() (*field.CopyMsgIndicator, error) {
	f := new(field.CopyMsgIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) LegalConfirm() (*field.LegalConfirm, error) {
	f := new(field.LegalConfirm)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) ConfirmStatus() (*field.ConfirmStatus, error) {
	f := new(field.ConfirmStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) SecondaryAllocID() (*field.SecondaryAllocID, error) {
	f := new(field.SecondaryAllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) IndividualAllocID() (*field.IndividualAllocID, error) {
	f := new(field.IndividualAllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) AllocQty() (*field.AllocQty, error) {
	f := new(field.AllocQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) AllocAccount() (*field.AllocAccount, error) {
	f := new(field.AllocAccount)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) AllocAcctIDSource() (*field.AllocAcctIDSource, error) {
	f := new(field.AllocAcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) AllocAccountType() (*field.AllocAccountType, error) {
	f := new(field.AllocAccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) AvgPxPrecision() (*field.AvgPxPrecision, error) {
	f := new(field.AvgPxPrecision)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) AvgParPx() (*field.AvgParPx, error) {
	f := new(field.AvgParPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) ReportedPx() (*field.ReportedPx, error) {
	f := new(field.ReportedPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) ProcessCode() (*field.ProcessCode, error) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) GrossTradeAmt() (*field.GrossTradeAmt, error) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) NumDaysInterest() (*field.NumDaysInterest, error) {
	f := new(field.NumDaysInterest)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) ExDate() (*field.ExDate, error) {
	f := new(field.ExDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) AccruedInterestRate() (*field.AccruedInterestRate, error) {
	f := new(field.AccruedInterestRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) AccruedInterestAmt() (*field.AccruedInterestAmt, error) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) InterestAtMaturity() (*field.InterestAtMaturity, error) {
	f := new(field.InterestAtMaturity)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) EndAccruedInterestAmt() (*field.EndAccruedInterestAmt, error) {
	f := new(field.EndAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) StartCash() (*field.StartCash, error) {
	f := new(field.StartCash)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) EndCash() (*field.EndCash, error) {
	f := new(field.EndCash)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) Concession() (*field.Concession, error) {
	f := new(field.Concession)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) TotalTakedown() (*field.TotalTakedown, error) {
	f := new(field.TotalTakedown)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) NetMoney() (*field.NetMoney, error) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) MaturityNetMoney() (*field.MaturityNetMoney, error) {
	f := new(field.MaturityNetMoney)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) SettlCurrAmt() (*field.SettlCurrAmt, error) {
	f := new(field.SettlCurrAmt)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) SettlCurrFxRate() (*field.SettlCurrFxRate, error) {
	f := new(field.SettlCurrFxRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalc, error) {
	f := new(field.SettlCurrFxRateCalc)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *Confirmation) SharedCommission() (*field.SharedCommission, error) {
	f := new(field.SharedCommission)
	err := m.Body.Get(f)
	return f, err
}
