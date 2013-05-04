package fix50

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type NewOrderCross struct {
	quickfixgo.Message
}

func (m *NewOrderCross) CrossID() (*field.CrossID, error) {
	f := new(field.CrossID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) CrossType() (*field.CrossType, error) {
	f := new(field.CrossType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) CrossPrioritization() (*field.CrossPrioritization, error) {
	f := new(field.CrossPrioritization)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) HandlInst() (*field.HandlInst, error) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ExecInst() (*field.ExecInst, error) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MinQty() (*field.MinQty, error) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MaxFloor() (*field.MaxFloor, error) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ExDestination() (*field.ExDestination, error) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ProcessCode() (*field.ProcessCode, error) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PrevClosePx() (*field.PrevClosePx, error) {
	f := new(field.PrevClosePx)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) LocateReqd() (*field.LocateReqd, error) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) StopPx() (*field.StopPx, error) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ComplianceID() (*field.ComplianceID, error) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) IOIID() (*field.IOIID, error) {
	f := new(field.IOIID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TimeInForce() (*field.TimeInForce, error) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) EffectiveTime() (*field.EffectiveTime, error) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ExpireDate() (*field.ExpireDate, error) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) GTBookingInst() (*field.GTBookingInst, error) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MaxShow() (*field.MaxShow, error) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TargetStrategy() (*field.TargetStrategy, error) {
	f := new(field.TargetStrategy)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TargetStrategyParameters() (*field.TargetStrategyParameters, error) {
	f := new(field.TargetStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ParticipationRate() (*field.ParticipationRate, error) {
	f := new(field.ParticipationRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) CancellationRights() (*field.CancellationRights, error) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, error) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) RegistID() (*field.RegistID, error) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) Designation() (*field.Designation, error) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) TransBkdTime() (*field.TransBkdTime, error) {
	f := new(field.TransBkdTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MatchIncrement() (*field.MatchIncrement, error) {
	f := new(field.MatchIncrement)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) MaxPriceLevels() (*field.MaxPriceLevels, error) {
	f := new(field.MaxPriceLevels)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) PriceProtectionScope() (*field.PriceProtectionScope, error) {
	f := new(field.PriceProtectionScope)
	err := m.Body.Get(f)
	return f, err
}
func (m *NewOrderCross) ExDestinationIDSource() (*field.ExDestinationIDSource, error) {
	f := new(field.ExDestinationIDSource)
	err := m.Body.Get(f)
	return f, err
}
