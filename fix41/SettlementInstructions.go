package fix41

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SettlementInstructions msg type = T.
type SettlementInstructions struct {
	message.Message
}

//SettlementInstructionsBuilder builds SettlementInstructions messages.
type SettlementInstructionsBuilder struct {
	message.MessageBuilder
}

//CreateSettlementInstructionsBuilder returns an initialized SettlementInstructionsBuilder with specified required fields.
func CreateSettlementInstructionsBuilder(
	settlinstid *field.SettlInstIDField,
	settlinsttranstype *field.SettlInstTransTypeField,
	settlinstmode *field.SettlInstModeField,
	settlinstsource *field.SettlInstSourceField,
	allocaccount *field.AllocAccountField,
	transacttime *field.TransactTimeField) SettlementInstructionsBuilder {
	var builder SettlementInstructionsBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX41))
	builder.Header().Set(field.NewMsgType("T"))
	builder.Body().Set(settlinstid)
	builder.Body().Set(settlinsttranstype)
	builder.Body().Set(settlinstmode)
	builder.Body().Set(settlinstsource)
	builder.Body().Set(allocaccount)
	builder.Body().Set(transacttime)
	return builder
}

//SettlInstID is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstID() (*field.SettlInstIDField, errors.MessageRejectError) {
	f := &field.SettlInstIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstID reads a SettlInstID from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstID(f *field.SettlInstIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstTransType is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstTransType() (*field.SettlInstTransTypeField, errors.MessageRejectError) {
	f := &field.SettlInstTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstTransType reads a SettlInstTransType from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstTransType(f *field.SettlInstTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstMode is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstMode() (*field.SettlInstModeField, errors.MessageRejectError) {
	f := &field.SettlInstModeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstMode reads a SettlInstMode from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstMode(f *field.SettlInstModeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstSource is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstSource() (*field.SettlInstSourceField, errors.MessageRejectError) {
	f := &field.SettlInstSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstSource reads a SettlInstSource from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstSource(f *field.SettlInstSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAccount is a required field for SettlementInstructions.
func (m SettlementInstructions) AllocAccount() (*field.AllocAccountField, errors.MessageRejectError) {
	f := &field.AllocAccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAccount reads a AllocAccount from SettlementInstructions.
func (m SettlementInstructions) GetAllocAccount(f *field.AllocAccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlLocation is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlLocation() (*field.SettlLocationField, errors.MessageRejectError) {
	f := &field.SettlLocationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlLocation reads a SettlLocation from SettlementInstructions.
func (m SettlementInstructions) GetSettlLocation(f *field.SettlLocationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for SettlementInstructions.
func (m SettlementInstructions) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from SettlementInstructions.
func (m SettlementInstructions) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) AllocID() (*field.AllocIDField, errors.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from SettlementInstructions.
func (m SettlementInstructions) GetAllocID(f *field.AllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for SettlementInstructions.
func (m SettlementInstructions) LastMkt() (*field.LastMktField, errors.MessageRejectError) {
	f := &field.LastMktField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from SettlementInstructions.
func (m SettlementInstructions) GetLastMkt(f *field.LastMktField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for SettlementInstructions.
func (m SettlementInstructions) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from SettlementInstructions.
func (m SettlementInstructions) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from SettlementInstructions.
func (m SettlementInstructions) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for SettlementInstructions.
func (m SettlementInstructions) EffectiveTime() (*field.EffectiveTimeField, errors.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from SettlementInstructions.
func (m SettlementInstructions) GetEffectiveTime(f *field.EffectiveTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for SettlementInstructions.
func (m SettlementInstructions) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SettlementInstructions.
func (m SettlementInstructions) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClientID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) ClientID() (*field.ClientIDField, errors.MessageRejectError) {
	f := &field.ClientIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClientID reads a ClientID from SettlementInstructions.
func (m SettlementInstructions) GetClientID(f *field.ClientIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecBroker is a non-required field for SettlementInstructions.
func (m SettlementInstructions) ExecBroker() (*field.ExecBrokerField, errors.MessageRejectError) {
	f := &field.ExecBrokerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecBroker reads a ExecBroker from SettlementInstructions.
func (m SettlementInstructions) GetExecBroker(f *field.ExecBrokerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbType is a non-required field for SettlementInstructions.
func (m SettlementInstructions) StandInstDbType() (*field.StandInstDbTypeField, errors.MessageRejectError) {
	f := &field.StandInstDbTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbType reads a StandInstDbType from SettlementInstructions.
func (m SettlementInstructions) GetStandInstDbType(f *field.StandInstDbTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) StandInstDbName() (*field.StandInstDbNameField, errors.MessageRejectError) {
	f := &field.StandInstDbNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbName reads a StandInstDbName from SettlementInstructions.
func (m SettlementInstructions) GetStandInstDbName(f *field.StandInstDbNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) StandInstDbID() (*field.StandInstDbIDField, errors.MessageRejectError) {
	f := &field.StandInstDbIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbID reads a StandInstDbID from SettlementInstructions.
func (m SettlementInstructions) GetStandInstDbID(f *field.StandInstDbIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDeliveryType is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlDeliveryType() (*field.SettlDeliveryTypeField, errors.MessageRejectError) {
	f := &field.SettlDeliveryTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDeliveryType reads a SettlDeliveryType from SettlementInstructions.
func (m SettlementInstructions) GetSettlDeliveryType(f *field.SettlDeliveryTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDepositoryCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlDepositoryCode() (*field.SettlDepositoryCodeField, errors.MessageRejectError) {
	f := &field.SettlDepositoryCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDepositoryCode reads a SettlDepositoryCode from SettlementInstructions.
func (m SettlementInstructions) GetSettlDepositoryCode(f *field.SettlDepositoryCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlBrkrCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlBrkrCode() (*field.SettlBrkrCodeField, errors.MessageRejectError) {
	f := &field.SettlBrkrCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlBrkrCode reads a SettlBrkrCode from SettlementInstructions.
func (m SettlementInstructions) GetSettlBrkrCode(f *field.SettlBrkrCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstCode() (*field.SettlInstCodeField, errors.MessageRejectError) {
	f := &field.SettlInstCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstCode reads a SettlInstCode from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstCode(f *field.SettlInstCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentName() (*field.SecuritySettlAgentNameField, errors.MessageRejectError) {
	f := &field.SecuritySettlAgentNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentName reads a SecuritySettlAgentName from SettlementInstructions.
func (m SettlementInstructions) GetSecuritySettlAgentName(f *field.SecuritySettlAgentNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentCode() (*field.SecuritySettlAgentCodeField, errors.MessageRejectError) {
	f := &field.SecuritySettlAgentCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentCode reads a SecuritySettlAgentCode from SettlementInstructions.
func (m SettlementInstructions) GetSecuritySettlAgentCode(f *field.SecuritySettlAgentCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentAcctNum is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentAcctNum() (*field.SecuritySettlAgentAcctNumField, errors.MessageRejectError) {
	f := &field.SecuritySettlAgentAcctNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentAcctNum reads a SecuritySettlAgentAcctNum from SettlementInstructions.
func (m SettlementInstructions) GetSecuritySettlAgentAcctNum(f *field.SecuritySettlAgentAcctNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentAcctName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentAcctName() (*field.SecuritySettlAgentAcctNameField, errors.MessageRejectError) {
	f := &field.SecuritySettlAgentAcctNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentAcctName reads a SecuritySettlAgentAcctName from SettlementInstructions.
func (m SettlementInstructions) GetSecuritySettlAgentAcctName(f *field.SecuritySettlAgentAcctNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentContactName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentContactName() (*field.SecuritySettlAgentContactNameField, errors.MessageRejectError) {
	f := &field.SecuritySettlAgentContactNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentContactName reads a SecuritySettlAgentContactName from SettlementInstructions.
func (m SettlementInstructions) GetSecuritySettlAgentContactName(f *field.SecuritySettlAgentContactNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentContactPhone is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentContactPhone() (*field.SecuritySettlAgentContactPhoneField, errors.MessageRejectError) {
	f := &field.SecuritySettlAgentContactPhoneField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentContactPhone reads a SecuritySettlAgentContactPhone from SettlementInstructions.
func (m SettlementInstructions) GetSecuritySettlAgentContactPhone(f *field.SecuritySettlAgentContactPhoneField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentName() (*field.CashSettlAgentNameField, errors.MessageRejectError) {
	f := &field.CashSettlAgentNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentName reads a CashSettlAgentName from SettlementInstructions.
func (m SettlementInstructions) GetCashSettlAgentName(f *field.CashSettlAgentNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentCode() (*field.CashSettlAgentCodeField, errors.MessageRejectError) {
	f := &field.CashSettlAgentCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentCode reads a CashSettlAgentCode from SettlementInstructions.
func (m SettlementInstructions) GetCashSettlAgentCode(f *field.CashSettlAgentCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentAcctNum is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentAcctNum() (*field.CashSettlAgentAcctNumField, errors.MessageRejectError) {
	f := &field.CashSettlAgentAcctNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentAcctNum reads a CashSettlAgentAcctNum from SettlementInstructions.
func (m SettlementInstructions) GetCashSettlAgentAcctNum(f *field.CashSettlAgentAcctNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentAcctName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentAcctName() (*field.CashSettlAgentAcctNameField, errors.MessageRejectError) {
	f := &field.CashSettlAgentAcctNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentAcctName reads a CashSettlAgentAcctName from SettlementInstructions.
func (m SettlementInstructions) GetCashSettlAgentAcctName(f *field.CashSettlAgentAcctNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentContactName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentContactName() (*field.CashSettlAgentContactNameField, errors.MessageRejectError) {
	f := &field.CashSettlAgentContactNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentContactName reads a CashSettlAgentContactName from SettlementInstructions.
func (m SettlementInstructions) GetCashSettlAgentContactName(f *field.CashSettlAgentContactNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentContactPhone is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentContactPhone() (*field.CashSettlAgentContactPhoneField, errors.MessageRejectError) {
	f := &field.CashSettlAgentContactPhoneField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentContactPhone reads a CashSettlAgentContactPhone from SettlementInstructions.
func (m SettlementInstructions) GetCashSettlAgentContactPhone(f *field.CashSettlAgentContactPhoneField) errors.MessageRejectError {
	return m.Body.Get(f)
}
