package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type SettlementInstructionRequest struct {
	message.Message
}

func (m *SettlementInstructionRequest) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructionRequest) AllocAcctIDSource() (*field.AllocAcctIDSource, error) {
	f := new(field.AllocAcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructionRequest) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructionRequest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructionRequest) SettlInstReqID() (*field.SettlInstReqID, error) {
	f := new(field.SettlInstReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructionRequest) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructionRequest) EffectiveTime() (*field.EffectiveTime, error) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructionRequest) StandInstDbID() (*field.StandInstDbID, error) {
	f := new(field.StandInstDbID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructionRequest) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructionRequest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructionRequest) StandInstDbType() (*field.StandInstDbType, error) {
	f := new(field.StandInstDbType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructionRequest) StandInstDbName() (*field.StandInstDbName, error) {
	f := new(field.StandInstDbName)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructionRequest) AllocAccount() (*field.AllocAccount, error) {
	f := new(field.AllocAccount)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructionRequest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructionRequest) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructionRequest) LastUpdateTime() (*field.LastUpdateTime, error) {
	f := new(field.LastUpdateTime)
	err := m.Body.Get(f)
	return f, err
}
