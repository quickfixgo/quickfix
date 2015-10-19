//Package settlementinstructions msg type = T.
package settlementinstructions

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a SettlementInstructions wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//SettlInstID is a required field for SettlementInstructions.
func (m Message) SettlInstID() (*field.SettlInstIDField, quickfix.MessageRejectError) {
	f := &field.SettlInstIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstID reads a SettlInstID from SettlementInstructions.
func (m Message) GetSettlInstID(f *field.SettlInstIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstTransType is a required field for SettlementInstructions.
func (m Message) SettlInstTransType() (*field.SettlInstTransTypeField, quickfix.MessageRejectError) {
	f := &field.SettlInstTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstTransType reads a SettlInstTransType from SettlementInstructions.
func (m Message) GetSettlInstTransType(f *field.SettlInstTransTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstRefID is a required field for SettlementInstructions.
func (m Message) SettlInstRefID() (*field.SettlInstRefIDField, quickfix.MessageRejectError) {
	f := &field.SettlInstRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstRefID reads a SettlInstRefID from SettlementInstructions.
func (m Message) GetSettlInstRefID(f *field.SettlInstRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstMode is a required field for SettlementInstructions.
func (m Message) SettlInstMode() (*field.SettlInstModeField, quickfix.MessageRejectError) {
	f := &field.SettlInstModeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstMode reads a SettlInstMode from SettlementInstructions.
func (m Message) GetSettlInstMode(f *field.SettlInstModeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstSource is a required field for SettlementInstructions.
func (m Message) SettlInstSource() (*field.SettlInstSourceField, quickfix.MessageRejectError) {
	f := &field.SettlInstSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstSource reads a SettlInstSource from SettlementInstructions.
func (m Message) GetSettlInstSource(f *field.SettlInstSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAccount is a required field for SettlementInstructions.
func (m Message) AllocAccount() (*field.AllocAccountField, quickfix.MessageRejectError) {
	f := &field.AllocAccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAccount reads a AllocAccount from SettlementInstructions.
func (m Message) GetAllocAccount(f *field.AllocAccountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//IndividualAllocID is a non-required field for SettlementInstructions.
func (m Message) IndividualAllocID() (*field.IndividualAllocIDField, quickfix.MessageRejectError) {
	f := &field.IndividualAllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIndividualAllocID reads a IndividualAllocID from SettlementInstructions.
func (m Message) GetIndividualAllocID(f *field.IndividualAllocIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for SettlementInstructions.
func (m Message) ClOrdID() (*field.ClOrdIDField, quickfix.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from SettlementInstructions.
func (m Message) GetClOrdID(f *field.ClOrdIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for SettlementInstructions.
func (m Message) TradeDate() (*field.TradeDateField, quickfix.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from SettlementInstructions.
func (m Message) GetTradeDate(f *field.TradeDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for SettlementInstructions.
func (m Message) AllocID() (*field.AllocIDField, quickfix.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from SettlementInstructions.
func (m Message) GetAllocID(f *field.AllocIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for SettlementInstructions.
func (m Message) LastMkt() (*field.LastMktField, quickfix.MessageRejectError) {
	f := &field.LastMktField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from SettlementInstructions.
func (m Message) GetLastMkt(f *field.LastMktField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for SettlementInstructions.
func (m Message) TradingSessionID() (*field.TradingSessionIDField, quickfix.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from SettlementInstructions.
func (m Message) GetTradingSessionID(f *field.TradingSessionIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for SettlementInstructions.
func (m Message) TradingSessionSubID() (*field.TradingSessionSubIDField, quickfix.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from SettlementInstructions.
func (m Message) GetTradingSessionSubID(f *field.TradingSessionSubIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for SettlementInstructions.
func (m Message) Side() (*field.SideField, quickfix.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from SettlementInstructions.
func (m Message) GetSide(f *field.SideField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for SettlementInstructions.
func (m Message) SecurityType() (*field.SecurityTypeField, quickfix.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from SettlementInstructions.
func (m Message) GetSecurityType(f *field.SecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for SettlementInstructions.
func (m Message) EffectiveTime() (*field.EffectiveTimeField, quickfix.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from SettlementInstructions.
func (m Message) GetEffectiveTime(f *field.EffectiveTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for SettlementInstructions.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SettlementInstructions.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for SettlementInstructions.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from SettlementInstructions.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbType is a non-required field for SettlementInstructions.
func (m Message) StandInstDbType() (*field.StandInstDbTypeField, quickfix.MessageRejectError) {
	f := &field.StandInstDbTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbType reads a StandInstDbType from SettlementInstructions.
func (m Message) GetStandInstDbType(f *field.StandInstDbTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbName is a non-required field for SettlementInstructions.
func (m Message) StandInstDbName() (*field.StandInstDbNameField, quickfix.MessageRejectError) {
	f := &field.StandInstDbNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbName reads a StandInstDbName from SettlementInstructions.
func (m Message) GetStandInstDbName(f *field.StandInstDbNameField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbID is a non-required field for SettlementInstructions.
func (m Message) StandInstDbID() (*field.StandInstDbIDField, quickfix.MessageRejectError) {
	f := &field.StandInstDbIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbID reads a StandInstDbID from SettlementInstructions.
func (m Message) GetStandInstDbID(f *field.StandInstDbIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDeliveryType is a non-required field for SettlementInstructions.
func (m Message) SettlDeliveryType() (*field.SettlDeliveryTypeField, quickfix.MessageRejectError) {
	f := &field.SettlDeliveryTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDeliveryType reads a SettlDeliveryType from SettlementInstructions.
func (m Message) GetSettlDeliveryType(f *field.SettlDeliveryTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDepositoryCode is a non-required field for SettlementInstructions.
func (m Message) SettlDepositoryCode() (*field.SettlDepositoryCodeField, quickfix.MessageRejectError) {
	f := &field.SettlDepositoryCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDepositoryCode reads a SettlDepositoryCode from SettlementInstructions.
func (m Message) GetSettlDepositoryCode(f *field.SettlDepositoryCodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlBrkrCode is a non-required field for SettlementInstructions.
func (m Message) SettlBrkrCode() (*field.SettlBrkrCodeField, quickfix.MessageRejectError) {
	f := &field.SettlBrkrCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlBrkrCode reads a SettlBrkrCode from SettlementInstructions.
func (m Message) GetSettlBrkrCode(f *field.SettlBrkrCodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstCode is a non-required field for SettlementInstructions.
func (m Message) SettlInstCode() (*field.SettlInstCodeField, quickfix.MessageRejectError) {
	f := &field.SettlInstCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstCode reads a SettlInstCode from SettlementInstructions.
func (m Message) GetSettlInstCode(f *field.SettlInstCodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentName is a non-required field for SettlementInstructions.
func (m Message) SecuritySettlAgentName() (*field.SecuritySettlAgentNameField, quickfix.MessageRejectError) {
	f := &field.SecuritySettlAgentNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentName reads a SecuritySettlAgentName from SettlementInstructions.
func (m Message) GetSecuritySettlAgentName(f *field.SecuritySettlAgentNameField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentCode is a non-required field for SettlementInstructions.
func (m Message) SecuritySettlAgentCode() (*field.SecuritySettlAgentCodeField, quickfix.MessageRejectError) {
	f := &field.SecuritySettlAgentCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentCode reads a SecuritySettlAgentCode from SettlementInstructions.
func (m Message) GetSecuritySettlAgentCode(f *field.SecuritySettlAgentCodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentAcctNum is a non-required field for SettlementInstructions.
func (m Message) SecuritySettlAgentAcctNum() (*field.SecuritySettlAgentAcctNumField, quickfix.MessageRejectError) {
	f := &field.SecuritySettlAgentAcctNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentAcctNum reads a SecuritySettlAgentAcctNum from SettlementInstructions.
func (m Message) GetSecuritySettlAgentAcctNum(f *field.SecuritySettlAgentAcctNumField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentAcctName is a non-required field for SettlementInstructions.
func (m Message) SecuritySettlAgentAcctName() (*field.SecuritySettlAgentAcctNameField, quickfix.MessageRejectError) {
	f := &field.SecuritySettlAgentAcctNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentAcctName reads a SecuritySettlAgentAcctName from SettlementInstructions.
func (m Message) GetSecuritySettlAgentAcctName(f *field.SecuritySettlAgentAcctNameField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentContactName is a non-required field for SettlementInstructions.
func (m Message) SecuritySettlAgentContactName() (*field.SecuritySettlAgentContactNameField, quickfix.MessageRejectError) {
	f := &field.SecuritySettlAgentContactNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentContactName reads a SecuritySettlAgentContactName from SettlementInstructions.
func (m Message) GetSecuritySettlAgentContactName(f *field.SecuritySettlAgentContactNameField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySettlAgentContactPhone is a non-required field for SettlementInstructions.
func (m Message) SecuritySettlAgentContactPhone() (*field.SecuritySettlAgentContactPhoneField, quickfix.MessageRejectError) {
	f := &field.SecuritySettlAgentContactPhoneField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySettlAgentContactPhone reads a SecuritySettlAgentContactPhone from SettlementInstructions.
func (m Message) GetSecuritySettlAgentContactPhone(f *field.SecuritySettlAgentContactPhoneField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentName is a non-required field for SettlementInstructions.
func (m Message) CashSettlAgentName() (*field.CashSettlAgentNameField, quickfix.MessageRejectError) {
	f := &field.CashSettlAgentNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentName reads a CashSettlAgentName from SettlementInstructions.
func (m Message) GetCashSettlAgentName(f *field.CashSettlAgentNameField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentCode is a non-required field for SettlementInstructions.
func (m Message) CashSettlAgentCode() (*field.CashSettlAgentCodeField, quickfix.MessageRejectError) {
	f := &field.CashSettlAgentCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentCode reads a CashSettlAgentCode from SettlementInstructions.
func (m Message) GetCashSettlAgentCode(f *field.CashSettlAgentCodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentAcctNum is a non-required field for SettlementInstructions.
func (m Message) CashSettlAgentAcctNum() (*field.CashSettlAgentAcctNumField, quickfix.MessageRejectError) {
	f := &field.CashSettlAgentAcctNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentAcctNum reads a CashSettlAgentAcctNum from SettlementInstructions.
func (m Message) GetCashSettlAgentAcctNum(f *field.CashSettlAgentAcctNumField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentAcctName is a non-required field for SettlementInstructions.
func (m Message) CashSettlAgentAcctName() (*field.CashSettlAgentAcctNameField, quickfix.MessageRejectError) {
	f := &field.CashSettlAgentAcctNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentAcctName reads a CashSettlAgentAcctName from SettlementInstructions.
func (m Message) GetCashSettlAgentAcctName(f *field.CashSettlAgentAcctNameField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentContactName is a non-required field for SettlementInstructions.
func (m Message) CashSettlAgentContactName() (*field.CashSettlAgentContactNameField, quickfix.MessageRejectError) {
	f := &field.CashSettlAgentContactNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentContactName reads a CashSettlAgentContactName from SettlementInstructions.
func (m Message) GetCashSettlAgentContactName(f *field.CashSettlAgentContactNameField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CashSettlAgentContactPhone is a non-required field for SettlementInstructions.
func (m Message) CashSettlAgentContactPhone() (*field.CashSettlAgentContactPhoneField, quickfix.MessageRejectError) {
	f := &field.CashSettlAgentContactPhoneField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashSettlAgentContactPhone reads a CashSettlAgentContactPhone from SettlementInstructions.
func (m Message) GetCashSettlAgentContactPhone(f *field.CashSettlAgentContactPhoneField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PaymentMethod is a non-required field for SettlementInstructions.
func (m Message) PaymentMethod() (*field.PaymentMethodField, quickfix.MessageRejectError) {
	f := &field.PaymentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPaymentMethod reads a PaymentMethod from SettlementInstructions.
func (m Message) GetPaymentMethod(f *field.PaymentMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PaymentRef is a non-required field for SettlementInstructions.
func (m Message) PaymentRef() (*field.PaymentRefField, quickfix.MessageRejectError) {
	f := &field.PaymentRefField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPaymentRef reads a PaymentRef from SettlementInstructions.
func (m Message) GetPaymentRef(f *field.PaymentRefField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CardHolderName is a non-required field for SettlementInstructions.
func (m Message) CardHolderName() (*field.CardHolderNameField, quickfix.MessageRejectError) {
	f := &field.CardHolderNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCardHolderName reads a CardHolderName from SettlementInstructions.
func (m Message) GetCardHolderName(f *field.CardHolderNameField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CardNumber is a non-required field for SettlementInstructions.
func (m Message) CardNumber() (*field.CardNumberField, quickfix.MessageRejectError) {
	f := &field.CardNumberField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCardNumber reads a CardNumber from SettlementInstructions.
func (m Message) GetCardNumber(f *field.CardNumberField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CardStartDate is a non-required field for SettlementInstructions.
func (m Message) CardStartDate() (*field.CardStartDateField, quickfix.MessageRejectError) {
	f := &field.CardStartDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCardStartDate reads a CardStartDate from SettlementInstructions.
func (m Message) GetCardStartDate(f *field.CardStartDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CardExpDate is a non-required field for SettlementInstructions.
func (m Message) CardExpDate() (*field.CardExpDateField, quickfix.MessageRejectError) {
	f := &field.CardExpDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCardExpDate reads a CardExpDate from SettlementInstructions.
func (m Message) GetCardExpDate(f *field.CardExpDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CardIssNo is a non-required field for SettlementInstructions.
func (m Message) CardIssNo() (*field.CardIssNoField, quickfix.MessageRejectError) {
	f := &field.CardIssNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCardIssNo reads a CardIssNo from SettlementInstructions.
func (m Message) GetCardIssNo(f *field.CardIssNoField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PaymentDate is a non-required field for SettlementInstructions.
func (m Message) PaymentDate() (*field.PaymentDateField, quickfix.MessageRejectError) {
	f := &field.PaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPaymentDate reads a PaymentDate from SettlementInstructions.
func (m Message) GetPaymentDate(f *field.PaymentDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PaymentRemitterID is a non-required field for SettlementInstructions.
func (m Message) PaymentRemitterID() (*field.PaymentRemitterIDField, quickfix.MessageRejectError) {
	f := &field.PaymentRemitterIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPaymentRemitterID reads a PaymentRemitterID from SettlementInstructions.
func (m Message) GetPaymentRemitterID(f *field.PaymentRemitterIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized MessageBuilder with specified required fields for SettlementInstructions.
func New(
	settlinstid *field.SettlInstIDField,
	settlinsttranstype *field.SettlInstTransTypeField,
	settlinstrefid *field.SettlInstRefIDField,
	settlinstmode *field.SettlInstModeField,
	settlinstsource *field.SettlInstSourceField,
	allocaccount *field.AllocAccountField,
	transacttime *field.TransactTimeField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX43))
	builder.Header.Set(field.NewMsgType("T"))
	builder.Body.Set(settlinstid)
	builder.Body.Set(settlinsttranstype)
	builder.Body.Set(settlinstrefid)
	builder.Body.Set(settlinstmode)
	builder.Body.Set(settlinstsource)
	builder.Body.Set(allocaccount)
	builder.Body.Set(transacttime)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX43, "T", r
}
