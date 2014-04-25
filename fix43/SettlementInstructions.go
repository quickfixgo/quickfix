package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
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
	settlinstid field.SettlInstID,
	settlinsttranstype field.SettlInstTransType,
	settlinstrefid field.SettlInstRefID,
	settlinstmode field.SettlInstMode,
	settlinstsource field.SettlInstSource,
	allocaccount field.AllocAccount,
	transacttime field.TransactTime) SettlementInstructionsBuilder {
	var builder SettlementInstructionsBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("T"))
	builder.Body.Set(settlinstid)
	builder.Body.Set(settlinsttranstype)
	builder.Body.Set(settlinstrefid)
	builder.Body.Set(settlinstmode)
	builder.Body.Set(settlinstsource)
	builder.Body.Set(allocaccount)
	builder.Body.Set(transacttime)
	return builder
}

//SettlInstID is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstID() (*field.SettlInstID, errors.MessageRejectError) {
	f := new(field.SettlInstID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstID reads a SettlInstID from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstID(f *field.SettlInstID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstTransType is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstTransType() (*field.SettlInstTransType, errors.MessageRejectError) {
	f := new(field.SettlInstTransType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstTransType reads a SettlInstTransType from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstTransType(f *field.SettlInstTransType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstRefID is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstRefID() (*field.SettlInstRefID, errors.MessageRejectError) {
	f := new(field.SettlInstRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstRefID reads a SettlInstRefID from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstRefID(f *field.SettlInstRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstMode is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstMode() (*field.SettlInstMode, errors.MessageRejectError) {
	f := new(field.SettlInstMode)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstMode reads a SettlInstMode from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstMode(f *field.SettlInstMode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstSource is a required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstSource() (*field.SettlInstSource, errors.MessageRejectError) {
	f := new(field.SettlInstSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstSource reads a SettlInstSource from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstSource(f *field.SettlInstSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAccount is a required field for SettlementInstructions.
func (m SettlementInstructions) AllocAccount() (*field.AllocAccount, errors.MessageRejectError) {
	f := new(field.AllocAccount)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAccount reads a AllocAccount from SettlementInstructions.
func (m SettlementInstructions) GetAllocAccount(f *field.AllocAccount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IndividualAllocID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) IndividualAllocID() (*field.IndividualAllocID, errors.MessageRejectError) {
	f := new(field.IndividualAllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetIndividualAllocID reads a IndividualAllocID from SettlementInstructions.
func (m SettlementInstructions) GetIndividualAllocID(f *field.IndividualAllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from SettlementInstructions.
func (m SettlementInstructions) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for SettlementInstructions.
func (m SettlementInstructions) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from SettlementInstructions.
func (m SettlementInstructions) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) AllocID() (*field.AllocID, errors.MessageRejectError) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from SettlementInstructions.
func (m SettlementInstructions) GetAllocID(f *field.AllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for SettlementInstructions.
func (m SettlementInstructions) LastMkt() (*field.LastMkt, errors.MessageRejectError) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from SettlementInstructions.
func (m SettlementInstructions) GetLastMkt(f *field.LastMkt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from SettlementInstructions.
func (m SettlementInstructions) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from SettlementInstructions.
func (m SettlementInstructions) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for SettlementInstructions.
func (m SettlementInstructions) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from SettlementInstructions.
func (m SettlementInstructions) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from SettlementInstructions.
func (m SettlementInstructions) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for SettlementInstructions.
func (m SettlementInstructions) EffectiveTime() (*field.EffectiveTime, errors.MessageRejectError) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from SettlementInstructions.
func (m SettlementInstructions) GetEffectiveTime(f *field.EffectiveTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for SettlementInstructions.
func (m SettlementInstructions) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SettlementInstructions.
func (m SettlementInstructions) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for SettlementInstructions.
func (m SettlementInstructions) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from SettlementInstructions.
func (m SettlementInstructions) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbType is a non-required field for SettlementInstructions.
func (m SettlementInstructions) StandInstDbType() (*field.StandInstDbType, errors.MessageRejectError) {
	f := new(field.StandInstDbType)
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbType reads a StandInstDbType from SettlementInstructions.
func (m SettlementInstructions) GetStandInstDbType(f *field.StandInstDbType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) StandInstDbName() (*field.StandInstDbName, errors.MessageRejectError) {
	f := new(field.StandInstDbName)
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbName reads a StandInstDbName from SettlementInstructions.
func (m SettlementInstructions) GetStandInstDbName(f *field.StandInstDbName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) StandInstDbID() (*field.StandInstDbID, errors.MessageRejectError) {
	f := new(field.StandInstDbID)
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbID reads a StandInstDbID from SettlementInstructions.
func (m SettlementInstructions) GetStandInstDbID(f *field.StandInstDbID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDeliveryType is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlDeliveryType() (*field.SettlDeliveryType, errors.MessageRejectError) {
	f := new(field.SettlDeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDeliveryType reads a SettlDeliveryType from SettlementInstructions.
func (m SettlementInstructions) GetSettlDeliveryType(f *field.SettlDeliveryType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDepositoryCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlDepositoryCode() (*field.SettlDepositoryCode, errors.MessageRejectError) {
	f := new(field.SettlDepositoryCode)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDepositoryCode reads a SettlDepositoryCode from SettlementInstructions.
func (m SettlementInstructions) GetSettlDepositoryCode(f *field.SettlDepositoryCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlBrkrCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlBrkrCode() (*field.SettlBrkrCode, errors.MessageRejectError) {
	f := new(field.SettlBrkrCode)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlBrkrCode reads a SettlBrkrCode from SettlementInstructions.
func (m SettlementInstructions) GetSettlBrkrCode(f *field.SettlBrkrCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SettlInstCode() (*field.SettlInstCode, errors.MessageRejectError) {
	f := new(field.SettlInstCode)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstCode reads a SettlInstCode from SettlementInstructions.
func (m SettlementInstructions) GetSettlInstCode(f *field.SettlInstCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentName() (*field.SecuritySettlAgentName, errors.MessageRejectError) {
	f := new(field.SecuritySettlAgentName)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentName reads a SecuritySettlAgentName from SettlementInstructions.
func (m SettlementInstructions) GetSecuritySettlAgentName(f *field.SecuritySettlAgentName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentCode() (*field.SecuritySettlAgentCode, errors.MessageRejectError) {
	f := new(field.SecuritySettlAgentCode)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentCode reads a SecuritySettlAgentCode from SettlementInstructions.
func (m SettlementInstructions) GetSecuritySettlAgentCode(f *field.SecuritySettlAgentCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentAcctNum is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentAcctNum() (*field.SecuritySettlAgentAcctNum, errors.MessageRejectError) {
	f := new(field.SecuritySettlAgentAcctNum)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentAcctNum reads a SecuritySettlAgentAcctNum from SettlementInstructions.
func (m SettlementInstructions) GetSecuritySettlAgentAcctNum(f *field.SecuritySettlAgentAcctNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentAcctName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentAcctName() (*field.SecuritySettlAgentAcctName, errors.MessageRejectError) {
	f := new(field.SecuritySettlAgentAcctName)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentAcctName reads a SecuritySettlAgentAcctName from SettlementInstructions.
func (m SettlementInstructions) GetSecuritySettlAgentAcctName(f *field.SecuritySettlAgentAcctName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentContactName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentContactName() (*field.SecuritySettlAgentContactName, errors.MessageRejectError) {
	f := new(field.SecuritySettlAgentContactName)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentContactName reads a SecuritySettlAgentContactName from SettlementInstructions.
func (m SettlementInstructions) GetSecuritySettlAgentContactName(f *field.SecuritySettlAgentContactName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentContactPhone is a non-required field for SettlementInstructions.
func (m SettlementInstructions) SecuritySettlAgentContactPhone() (*field.SecuritySettlAgentContactPhone, errors.MessageRejectError) {
	f := new(field.SecuritySettlAgentContactPhone)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentContactPhone reads a SecuritySettlAgentContactPhone from SettlementInstructions.
func (m SettlementInstructions) GetSecuritySettlAgentContactPhone(f *field.SecuritySettlAgentContactPhone) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentName() (*field.CashSettlAgentName, errors.MessageRejectError) {
	f := new(field.CashSettlAgentName)
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentName reads a CashSettlAgentName from SettlementInstructions.
func (m SettlementInstructions) GetCashSettlAgentName(f *field.CashSettlAgentName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentCode is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentCode() (*field.CashSettlAgentCode, errors.MessageRejectError) {
	f := new(field.CashSettlAgentCode)
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentCode reads a CashSettlAgentCode from SettlementInstructions.
func (m SettlementInstructions) GetCashSettlAgentCode(f *field.CashSettlAgentCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentAcctNum is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentAcctNum() (*field.CashSettlAgentAcctNum, errors.MessageRejectError) {
	f := new(field.CashSettlAgentAcctNum)
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentAcctNum reads a CashSettlAgentAcctNum from SettlementInstructions.
func (m SettlementInstructions) GetCashSettlAgentAcctNum(f *field.CashSettlAgentAcctNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentAcctName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentAcctName() (*field.CashSettlAgentAcctName, errors.MessageRejectError) {
	f := new(field.CashSettlAgentAcctName)
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentAcctName reads a CashSettlAgentAcctName from SettlementInstructions.
func (m SettlementInstructions) GetCashSettlAgentAcctName(f *field.CashSettlAgentAcctName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentContactName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentContactName() (*field.CashSettlAgentContactName, errors.MessageRejectError) {
	f := new(field.CashSettlAgentContactName)
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentContactName reads a CashSettlAgentContactName from SettlementInstructions.
func (m SettlementInstructions) GetCashSettlAgentContactName(f *field.CashSettlAgentContactName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentContactPhone is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CashSettlAgentContactPhone() (*field.CashSettlAgentContactPhone, errors.MessageRejectError) {
	f := new(field.CashSettlAgentContactPhone)
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentContactPhone reads a CashSettlAgentContactPhone from SettlementInstructions.
func (m SettlementInstructions) GetCashSettlAgentContactPhone(f *field.CashSettlAgentContactPhone) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PaymentMethod is a non-required field for SettlementInstructions.
func (m SettlementInstructions) PaymentMethod() (*field.PaymentMethod, errors.MessageRejectError) {
	f := new(field.PaymentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPaymentMethod reads a PaymentMethod from SettlementInstructions.
func (m SettlementInstructions) GetPaymentMethod(f *field.PaymentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PaymentRef is a non-required field for SettlementInstructions.
func (m SettlementInstructions) PaymentRef() (*field.PaymentRef, errors.MessageRejectError) {
	f := new(field.PaymentRef)
	err := m.Body.Get(f)
	return f, err
}

//GetPaymentRef reads a PaymentRef from SettlementInstructions.
func (m SettlementInstructions) GetPaymentRef(f *field.PaymentRef) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CardHolderName is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CardHolderName() (*field.CardHolderName, errors.MessageRejectError) {
	f := new(field.CardHolderName)
	err := m.Body.Get(f)
	return f, err
}

//GetCardHolderName reads a CardHolderName from SettlementInstructions.
func (m SettlementInstructions) GetCardHolderName(f *field.CardHolderName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CardNumber is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CardNumber() (*field.CardNumber, errors.MessageRejectError) {
	f := new(field.CardNumber)
	err := m.Body.Get(f)
	return f, err
}

//GetCardNumber reads a CardNumber from SettlementInstructions.
func (m SettlementInstructions) GetCardNumber(f *field.CardNumber) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CardStartDate is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CardStartDate() (*field.CardStartDate, errors.MessageRejectError) {
	f := new(field.CardStartDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCardStartDate reads a CardStartDate from SettlementInstructions.
func (m SettlementInstructions) GetCardStartDate(f *field.CardStartDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CardExpDate is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CardExpDate() (*field.CardExpDate, errors.MessageRejectError) {
	f := new(field.CardExpDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCardExpDate reads a CardExpDate from SettlementInstructions.
func (m SettlementInstructions) GetCardExpDate(f *field.CardExpDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CardIssNo is a non-required field for SettlementInstructions.
func (m SettlementInstructions) CardIssNo() (*field.CardIssNo, errors.MessageRejectError) {
	f := new(field.CardIssNo)
	err := m.Body.Get(f)
	return f, err
}

//GetCardIssNo reads a CardIssNo from SettlementInstructions.
func (m SettlementInstructions) GetCardIssNo(f *field.CardIssNo) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PaymentDate is a non-required field for SettlementInstructions.
func (m SettlementInstructions) PaymentDate() (*field.PaymentDate, errors.MessageRejectError) {
	f := new(field.PaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetPaymentDate reads a PaymentDate from SettlementInstructions.
func (m SettlementInstructions) GetPaymentDate(f *field.PaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PaymentRemitterID is a non-required field for SettlementInstructions.
func (m SettlementInstructions) PaymentRemitterID() (*field.PaymentRemitterID, errors.MessageRejectError) {
	f := new(field.PaymentRemitterID)
	err := m.Body.Get(f)
	return f, err
}

//GetPaymentRemitterID reads a PaymentRemitterID from SettlementInstructions.
func (m SettlementInstructions) GetPaymentRemitterID(f *field.PaymentRemitterID) errors.MessageRejectError {
	return m.Body.Get(f)
}
