package fix41

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type SettlementInstructions struct {
	quickfixgo.Message
}

func (m *SettlementInstructions) SettlInstID() (*field.SettlInstID, error) {
	f := new(field.SettlInstID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SettlInstTransType() (*field.SettlInstTransType, error) {
	f := new(field.SettlInstTransType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SettlInstMode() (*field.SettlInstMode, error) {
	f := new(field.SettlInstMode)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SettlInstSource() (*field.SettlInstSource, error) {
	f := new(field.SettlInstSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) AllocAccount() (*field.AllocAccount, error) {
	f := new(field.AllocAccount)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SettlLocation() (*field.SettlLocation, error) {
	f := new(field.SettlLocation)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) EffectiveTime() (*field.EffectiveTime, error) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) ClientID() (*field.ClientID, error) {
	f := new(field.ClientID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) ExecBroker() (*field.ExecBroker, error) {
	f := new(field.ExecBroker)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) StandInstDbType() (*field.StandInstDbType, error) {
	f := new(field.StandInstDbType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) StandInstDbName() (*field.StandInstDbName, error) {
	f := new(field.StandInstDbName)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) StandInstDbID() (*field.StandInstDbID, error) {
	f := new(field.StandInstDbID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SettlDeliveryType() (*field.SettlDeliveryType, error) {
	f := new(field.SettlDeliveryType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SettlDepositoryCode() (*field.SettlDepositoryCode, error) {
	f := new(field.SettlDepositoryCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SettlBrkrCode() (*field.SettlBrkrCode, error) {
	f := new(field.SettlBrkrCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SettlInstCode() (*field.SettlInstCode, error) {
	f := new(field.SettlInstCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SecuritySettlAgentName() (*field.SecuritySettlAgentName, error) {
	f := new(field.SecuritySettlAgentName)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SecuritySettlAgentCode() (*field.SecuritySettlAgentCode, error) {
	f := new(field.SecuritySettlAgentCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SecuritySettlAgentAcctNum() (*field.SecuritySettlAgentAcctNum, error) {
	f := new(field.SecuritySettlAgentAcctNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SecuritySettlAgentAcctName() (*field.SecuritySettlAgentAcctName, error) {
	f := new(field.SecuritySettlAgentAcctName)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SecuritySettlAgentContactName() (*field.SecuritySettlAgentContactName, error) {
	f := new(field.SecuritySettlAgentContactName)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) SecuritySettlAgentContactPhone() (*field.SecuritySettlAgentContactPhone, error) {
	f := new(field.SecuritySettlAgentContactPhone)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) CashSettlAgentName() (*field.CashSettlAgentName, error) {
	f := new(field.CashSettlAgentName)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) CashSettlAgentCode() (*field.CashSettlAgentCode, error) {
	f := new(field.CashSettlAgentCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) CashSettlAgentAcctNum() (*field.CashSettlAgentAcctNum, error) {
	f := new(field.CashSettlAgentAcctNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) CashSettlAgentAcctName() (*field.CashSettlAgentAcctName, error) {
	f := new(field.CashSettlAgentAcctName)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) CashSettlAgentContactName() (*field.CashSettlAgentContactName, error) {
	f := new(field.CashSettlAgentContactName)
	err := m.Body.Get(f)
	return f, err
}
func (m *SettlementInstructions) CashSettlAgentContactPhone() (*field.CashSettlAgentContactPhone, error) {
	f := new(field.CashSettlAgentContactPhone)
	err := m.Body.Get(f)
	return f, err
}
