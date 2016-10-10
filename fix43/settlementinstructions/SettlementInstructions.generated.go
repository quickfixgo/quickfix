package settlementinstructions

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//SettlementInstructions is the fix43 SettlementInstructions type, MsgType = T
type SettlementInstructions struct {
	fix43.Header
	*quickfix.Body
	fix43.Trailer
	Message *quickfix.Message
}

//FromMessage creates a SettlementInstructions from a quickfix.Message instance
func FromMessage(m *quickfix.Message) SettlementInstructions {
	return SettlementInstructions{
		Header:  fix43.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix43.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m SettlementInstructions) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a SettlementInstructions initialized with the required fields for SettlementInstructions
func New(settlinstid field.SettlInstIDField, settlinsttranstype field.SettlInstTransTypeField, settlinstrefid field.SettlInstRefIDField, settlinstmode field.SettlInstModeField, settlinstsource field.SettlInstSourceField, allocaccount field.AllocAccountField, transacttime field.TransactTimeField) (m SettlementInstructions) {
	m.Message = quickfix.NewMessage()
	m.Header = fix43.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("T"))
	m.Set(settlinstid)
	m.Set(settlinsttranstype)
	m.Set(settlinstrefid)
	m.Set(settlinstmode)
	m.Set(settlinstsource)
	m.Set(allocaccount)
	m.Set(transacttime)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg SettlementInstructions, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "T", r
}

//SetClOrdID sets ClOrdID, Tag 11
func (m SettlementInstructions) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetLastMkt sets LastMkt, Tag 30
func (m SettlementInstructions) SetLastMkt(v string) {
	m.Set(field.NewLastMkt(v))
}

//SetSide sets Side, Tag 54
func (m SettlementInstructions) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m SettlementInstructions) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetAllocID sets AllocID, Tag 70
func (m SettlementInstructions) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m SettlementInstructions) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetAllocAccount sets AllocAccount, Tag 79
func (m SettlementInstructions) SetAllocAccount(v string) {
	m.Set(field.NewAllocAccount(v))
}

//SetSettlInstMode sets SettlInstMode, Tag 160
func (m SettlementInstructions) SetSettlInstMode(v enum.SettlInstMode) {
	m.Set(field.NewSettlInstMode(v))
}

//SetSettlInstID sets SettlInstID, Tag 162
func (m SettlementInstructions) SetSettlInstID(v string) {
	m.Set(field.NewSettlInstID(v))
}

//SetSettlInstTransType sets SettlInstTransType, Tag 163
func (m SettlementInstructions) SetSettlInstTransType(v enum.SettlInstTransType) {
	m.Set(field.NewSettlInstTransType(v))
}

//SetSettlInstSource sets SettlInstSource, Tag 165
func (m SettlementInstructions) SetSettlInstSource(v enum.SettlInstSource) {
	m.Set(field.NewSettlInstSource(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m SettlementInstructions) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

//SetEffectiveTime sets EffectiveTime, Tag 168
func (m SettlementInstructions) SetEffectiveTime(v time.Time) {
	m.Set(field.NewEffectiveTime(v))
}

//SetStandInstDbType sets StandInstDbType, Tag 169
func (m SettlementInstructions) SetStandInstDbType(v enum.StandInstDbType) {
	m.Set(field.NewStandInstDbType(v))
}

//SetStandInstDbName sets StandInstDbName, Tag 170
func (m SettlementInstructions) SetStandInstDbName(v string) {
	m.Set(field.NewStandInstDbName(v))
}

//SetStandInstDbID sets StandInstDbID, Tag 171
func (m SettlementInstructions) SetStandInstDbID(v string) {
	m.Set(field.NewStandInstDbID(v))
}

//SetSettlDeliveryType sets SettlDeliveryType, Tag 172
func (m SettlementInstructions) SetSettlDeliveryType(v enum.SettlDeliveryType) {
	m.Set(field.NewSettlDeliveryType(v))
}

//SetSettlDepositoryCode sets SettlDepositoryCode, Tag 173
func (m SettlementInstructions) SetSettlDepositoryCode(v string) {
	m.Set(field.NewSettlDepositoryCode(v))
}

//SetSettlBrkrCode sets SettlBrkrCode, Tag 174
func (m SettlementInstructions) SetSettlBrkrCode(v string) {
	m.Set(field.NewSettlBrkrCode(v))
}

//SetSettlInstCode sets SettlInstCode, Tag 175
func (m SettlementInstructions) SetSettlInstCode(v string) {
	m.Set(field.NewSettlInstCode(v))
}

//SetSecuritySettlAgentName sets SecuritySettlAgentName, Tag 176
func (m SettlementInstructions) SetSecuritySettlAgentName(v string) {
	m.Set(field.NewSecuritySettlAgentName(v))
}

//SetSecuritySettlAgentCode sets SecuritySettlAgentCode, Tag 177
func (m SettlementInstructions) SetSecuritySettlAgentCode(v string) {
	m.Set(field.NewSecuritySettlAgentCode(v))
}

//SetSecuritySettlAgentAcctNum sets SecuritySettlAgentAcctNum, Tag 178
func (m SettlementInstructions) SetSecuritySettlAgentAcctNum(v string) {
	m.Set(field.NewSecuritySettlAgentAcctNum(v))
}

//SetSecuritySettlAgentAcctName sets SecuritySettlAgentAcctName, Tag 179
func (m SettlementInstructions) SetSecuritySettlAgentAcctName(v string) {
	m.Set(field.NewSecuritySettlAgentAcctName(v))
}

//SetSecuritySettlAgentContactName sets SecuritySettlAgentContactName, Tag 180
func (m SettlementInstructions) SetSecuritySettlAgentContactName(v string) {
	m.Set(field.NewSecuritySettlAgentContactName(v))
}

//SetSecuritySettlAgentContactPhone sets SecuritySettlAgentContactPhone, Tag 181
func (m SettlementInstructions) SetSecuritySettlAgentContactPhone(v string) {
	m.Set(field.NewSecuritySettlAgentContactPhone(v))
}

//SetCashSettlAgentName sets CashSettlAgentName, Tag 182
func (m SettlementInstructions) SetCashSettlAgentName(v string) {
	m.Set(field.NewCashSettlAgentName(v))
}

//SetCashSettlAgentCode sets CashSettlAgentCode, Tag 183
func (m SettlementInstructions) SetCashSettlAgentCode(v string) {
	m.Set(field.NewCashSettlAgentCode(v))
}

//SetCashSettlAgentAcctNum sets CashSettlAgentAcctNum, Tag 184
func (m SettlementInstructions) SetCashSettlAgentAcctNum(v string) {
	m.Set(field.NewCashSettlAgentAcctNum(v))
}

//SetCashSettlAgentAcctName sets CashSettlAgentAcctName, Tag 185
func (m SettlementInstructions) SetCashSettlAgentAcctName(v string) {
	m.Set(field.NewCashSettlAgentAcctName(v))
}

//SetCashSettlAgentContactName sets CashSettlAgentContactName, Tag 186
func (m SettlementInstructions) SetCashSettlAgentContactName(v string) {
	m.Set(field.NewCashSettlAgentContactName(v))
}

//SetCashSettlAgentContactPhone sets CashSettlAgentContactPhone, Tag 187
func (m SettlementInstructions) SetCashSettlAgentContactPhone(v string) {
	m.Set(field.NewCashSettlAgentContactPhone(v))
}

//SetSettlInstRefID sets SettlInstRefID, Tag 214
func (m SettlementInstructions) SetSettlInstRefID(v string) {
	m.Set(field.NewSettlInstRefID(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m SettlementInstructions) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m SettlementInstructions) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetIndividualAllocID sets IndividualAllocID, Tag 467
func (m SettlementInstructions) SetIndividualAllocID(v string) {
	m.Set(field.NewIndividualAllocID(v))
}

//SetPaymentRef sets PaymentRef, Tag 476
func (m SettlementInstructions) SetPaymentRef(v string) {
	m.Set(field.NewPaymentRef(v))
}

//SetCardHolderName sets CardHolderName, Tag 488
func (m SettlementInstructions) SetCardHolderName(v string) {
	m.Set(field.NewCardHolderName(v))
}

//SetCardNumber sets CardNumber, Tag 489
func (m SettlementInstructions) SetCardNumber(v string) {
	m.Set(field.NewCardNumber(v))
}

//SetCardExpDate sets CardExpDate, Tag 490
func (m SettlementInstructions) SetCardExpDate(v string) {
	m.Set(field.NewCardExpDate(v))
}

//SetCardIssNo sets CardIssNo, Tag 491
func (m SettlementInstructions) SetCardIssNo(v string) {
	m.Set(field.NewCardIssNo(v))
}

//SetPaymentMethod sets PaymentMethod, Tag 492
func (m SettlementInstructions) SetPaymentMethod(v enum.PaymentMethod) {
	m.Set(field.NewPaymentMethod(v))
}

//SetCardStartDate sets CardStartDate, Tag 503
func (m SettlementInstructions) SetCardStartDate(v string) {
	m.Set(field.NewCardStartDate(v))
}

//SetPaymentDate sets PaymentDate, Tag 504
func (m SettlementInstructions) SetPaymentDate(v string) {
	m.Set(field.NewPaymentDate(v))
}

//SetPaymentRemitterID sets PaymentRemitterID, Tag 505
func (m SettlementInstructions) SetPaymentRemitterID(v string) {
	m.Set(field.NewPaymentRemitterID(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m SettlementInstructions) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

//GetClOrdID gets ClOrdID, Tag 11
func (m SettlementInstructions) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastMkt gets LastMkt, Tag 30
func (m SettlementInstructions) GetLastMkt() (v string, err quickfix.MessageRejectError) {
	var f field.LastMktField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSide gets Side, Tag 54
func (m SettlementInstructions) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m SettlementInstructions) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocID gets AllocID, Tag 70
func (m SettlementInstructions) GetAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.AllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m SettlementInstructions) GetTradeDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocAccount gets AllocAccount, Tag 79
func (m SettlementInstructions) GetAllocAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AllocAccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlInstMode gets SettlInstMode, Tag 160
func (m SettlementInstructions) GetSettlInstMode() (v enum.SettlInstMode, err quickfix.MessageRejectError) {
	var f field.SettlInstModeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlInstID gets SettlInstID, Tag 162
func (m SettlementInstructions) GetSettlInstID() (v string, err quickfix.MessageRejectError) {
	var f field.SettlInstIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlInstTransType gets SettlInstTransType, Tag 163
func (m SettlementInstructions) GetSettlInstTransType() (v enum.SettlInstTransType, err quickfix.MessageRejectError) {
	var f field.SettlInstTransTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlInstSource gets SettlInstSource, Tag 165
func (m SettlementInstructions) GetSettlInstSource() (v enum.SettlInstSource, err quickfix.MessageRejectError) {
	var f field.SettlInstSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m SettlementInstructions) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEffectiveTime gets EffectiveTime, Tag 168
func (m SettlementInstructions) GetEffectiveTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.EffectiveTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStandInstDbType gets StandInstDbType, Tag 169
func (m SettlementInstructions) GetStandInstDbType() (v enum.StandInstDbType, err quickfix.MessageRejectError) {
	var f field.StandInstDbTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStandInstDbName gets StandInstDbName, Tag 170
func (m SettlementInstructions) GetStandInstDbName() (v string, err quickfix.MessageRejectError) {
	var f field.StandInstDbNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStandInstDbID gets StandInstDbID, Tag 171
func (m SettlementInstructions) GetStandInstDbID() (v string, err quickfix.MessageRejectError) {
	var f field.StandInstDbIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlDeliveryType gets SettlDeliveryType, Tag 172
func (m SettlementInstructions) GetSettlDeliveryType() (v enum.SettlDeliveryType, err quickfix.MessageRejectError) {
	var f field.SettlDeliveryTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlDepositoryCode gets SettlDepositoryCode, Tag 173
func (m SettlementInstructions) GetSettlDepositoryCode() (v string, err quickfix.MessageRejectError) {
	var f field.SettlDepositoryCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlBrkrCode gets SettlBrkrCode, Tag 174
func (m SettlementInstructions) GetSettlBrkrCode() (v string, err quickfix.MessageRejectError) {
	var f field.SettlBrkrCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlInstCode gets SettlInstCode, Tag 175
func (m SettlementInstructions) GetSettlInstCode() (v string, err quickfix.MessageRejectError) {
	var f field.SettlInstCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecuritySettlAgentName gets SecuritySettlAgentName, Tag 176
func (m SettlementInstructions) GetSecuritySettlAgentName() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySettlAgentNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecuritySettlAgentCode gets SecuritySettlAgentCode, Tag 177
func (m SettlementInstructions) GetSecuritySettlAgentCode() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySettlAgentCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecuritySettlAgentAcctNum gets SecuritySettlAgentAcctNum, Tag 178
func (m SettlementInstructions) GetSecuritySettlAgentAcctNum() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySettlAgentAcctNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecuritySettlAgentAcctName gets SecuritySettlAgentAcctName, Tag 179
func (m SettlementInstructions) GetSecuritySettlAgentAcctName() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySettlAgentAcctNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecuritySettlAgentContactName gets SecuritySettlAgentContactName, Tag 180
func (m SettlementInstructions) GetSecuritySettlAgentContactName() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySettlAgentContactNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecuritySettlAgentContactPhone gets SecuritySettlAgentContactPhone, Tag 181
func (m SettlementInstructions) GetSecuritySettlAgentContactPhone() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySettlAgentContactPhoneField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashSettlAgentName gets CashSettlAgentName, Tag 182
func (m SettlementInstructions) GetCashSettlAgentName() (v string, err quickfix.MessageRejectError) {
	var f field.CashSettlAgentNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashSettlAgentCode gets CashSettlAgentCode, Tag 183
func (m SettlementInstructions) GetCashSettlAgentCode() (v string, err quickfix.MessageRejectError) {
	var f field.CashSettlAgentCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashSettlAgentAcctNum gets CashSettlAgentAcctNum, Tag 184
func (m SettlementInstructions) GetCashSettlAgentAcctNum() (v string, err quickfix.MessageRejectError) {
	var f field.CashSettlAgentAcctNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashSettlAgentAcctName gets CashSettlAgentAcctName, Tag 185
func (m SettlementInstructions) GetCashSettlAgentAcctName() (v string, err quickfix.MessageRejectError) {
	var f field.CashSettlAgentAcctNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashSettlAgentContactName gets CashSettlAgentContactName, Tag 186
func (m SettlementInstructions) GetCashSettlAgentContactName() (v string, err quickfix.MessageRejectError) {
	var f field.CashSettlAgentContactNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashSettlAgentContactPhone gets CashSettlAgentContactPhone, Tag 187
func (m SettlementInstructions) GetCashSettlAgentContactPhone() (v string, err quickfix.MessageRejectError) {
	var f field.CashSettlAgentContactPhoneField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlInstRefID gets SettlInstRefID, Tag 214
func (m SettlementInstructions) GetSettlInstRefID() (v string, err quickfix.MessageRejectError) {
	var f field.SettlInstRefIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m SettlementInstructions) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m SettlementInstructions) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetIndividualAllocID gets IndividualAllocID, Tag 467
func (m SettlementInstructions) GetIndividualAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.IndividualAllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPaymentRef gets PaymentRef, Tag 476
func (m SettlementInstructions) GetPaymentRef() (v string, err quickfix.MessageRejectError) {
	var f field.PaymentRefField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCardHolderName gets CardHolderName, Tag 488
func (m SettlementInstructions) GetCardHolderName() (v string, err quickfix.MessageRejectError) {
	var f field.CardHolderNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCardNumber gets CardNumber, Tag 489
func (m SettlementInstructions) GetCardNumber() (v string, err quickfix.MessageRejectError) {
	var f field.CardNumberField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCardExpDate gets CardExpDate, Tag 490
func (m SettlementInstructions) GetCardExpDate() (v string, err quickfix.MessageRejectError) {
	var f field.CardExpDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCardIssNo gets CardIssNo, Tag 491
func (m SettlementInstructions) GetCardIssNo() (v string, err quickfix.MessageRejectError) {
	var f field.CardIssNoField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPaymentMethod gets PaymentMethod, Tag 492
func (m SettlementInstructions) GetPaymentMethod() (v enum.PaymentMethod, err quickfix.MessageRejectError) {
	var f field.PaymentMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCardStartDate gets CardStartDate, Tag 503
func (m SettlementInstructions) GetCardStartDate() (v string, err quickfix.MessageRejectError) {
	var f field.CardStartDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPaymentDate gets PaymentDate, Tag 504
func (m SettlementInstructions) GetPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.PaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPaymentRemitterID gets PaymentRemitterID, Tag 505
func (m SettlementInstructions) GetPaymentRemitterID() (v string, err quickfix.MessageRejectError) {
	var f field.PaymentRemitterIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m SettlementInstructions) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m SettlementInstructions) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasLastMkt returns true if LastMkt is present, Tag 30
func (m SettlementInstructions) HasLastMkt() bool {
	return m.Has(tag.LastMkt)
}

//HasSide returns true if Side is present, Tag 54
func (m SettlementInstructions) HasSide() bool {
	return m.Has(tag.Side)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m SettlementInstructions) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasAllocID returns true if AllocID is present, Tag 70
func (m SettlementInstructions) HasAllocID() bool {
	return m.Has(tag.AllocID)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m SettlementInstructions) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasAllocAccount returns true if AllocAccount is present, Tag 79
func (m SettlementInstructions) HasAllocAccount() bool {
	return m.Has(tag.AllocAccount)
}

//HasSettlInstMode returns true if SettlInstMode is present, Tag 160
func (m SettlementInstructions) HasSettlInstMode() bool {
	return m.Has(tag.SettlInstMode)
}

//HasSettlInstID returns true if SettlInstID is present, Tag 162
func (m SettlementInstructions) HasSettlInstID() bool {
	return m.Has(tag.SettlInstID)
}

//HasSettlInstTransType returns true if SettlInstTransType is present, Tag 163
func (m SettlementInstructions) HasSettlInstTransType() bool {
	return m.Has(tag.SettlInstTransType)
}

//HasSettlInstSource returns true if SettlInstSource is present, Tag 165
func (m SettlementInstructions) HasSettlInstSource() bool {
	return m.Has(tag.SettlInstSource)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m SettlementInstructions) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasEffectiveTime returns true if EffectiveTime is present, Tag 168
func (m SettlementInstructions) HasEffectiveTime() bool {
	return m.Has(tag.EffectiveTime)
}

//HasStandInstDbType returns true if StandInstDbType is present, Tag 169
func (m SettlementInstructions) HasStandInstDbType() bool {
	return m.Has(tag.StandInstDbType)
}

//HasStandInstDbName returns true if StandInstDbName is present, Tag 170
func (m SettlementInstructions) HasStandInstDbName() bool {
	return m.Has(tag.StandInstDbName)
}

//HasStandInstDbID returns true if StandInstDbID is present, Tag 171
func (m SettlementInstructions) HasStandInstDbID() bool {
	return m.Has(tag.StandInstDbID)
}

//HasSettlDeliveryType returns true if SettlDeliveryType is present, Tag 172
func (m SettlementInstructions) HasSettlDeliveryType() bool {
	return m.Has(tag.SettlDeliveryType)
}

//HasSettlDepositoryCode returns true if SettlDepositoryCode is present, Tag 173
func (m SettlementInstructions) HasSettlDepositoryCode() bool {
	return m.Has(tag.SettlDepositoryCode)
}

//HasSettlBrkrCode returns true if SettlBrkrCode is present, Tag 174
func (m SettlementInstructions) HasSettlBrkrCode() bool {
	return m.Has(tag.SettlBrkrCode)
}

//HasSettlInstCode returns true if SettlInstCode is present, Tag 175
func (m SettlementInstructions) HasSettlInstCode() bool {
	return m.Has(tag.SettlInstCode)
}

//HasSecuritySettlAgentName returns true if SecuritySettlAgentName is present, Tag 176
func (m SettlementInstructions) HasSecuritySettlAgentName() bool {
	return m.Has(tag.SecuritySettlAgentName)
}

//HasSecuritySettlAgentCode returns true if SecuritySettlAgentCode is present, Tag 177
func (m SettlementInstructions) HasSecuritySettlAgentCode() bool {
	return m.Has(tag.SecuritySettlAgentCode)
}

//HasSecuritySettlAgentAcctNum returns true if SecuritySettlAgentAcctNum is present, Tag 178
func (m SettlementInstructions) HasSecuritySettlAgentAcctNum() bool {
	return m.Has(tag.SecuritySettlAgentAcctNum)
}

//HasSecuritySettlAgentAcctName returns true if SecuritySettlAgentAcctName is present, Tag 179
func (m SettlementInstructions) HasSecuritySettlAgentAcctName() bool {
	return m.Has(tag.SecuritySettlAgentAcctName)
}

//HasSecuritySettlAgentContactName returns true if SecuritySettlAgentContactName is present, Tag 180
func (m SettlementInstructions) HasSecuritySettlAgentContactName() bool {
	return m.Has(tag.SecuritySettlAgentContactName)
}

//HasSecuritySettlAgentContactPhone returns true if SecuritySettlAgentContactPhone is present, Tag 181
func (m SettlementInstructions) HasSecuritySettlAgentContactPhone() bool {
	return m.Has(tag.SecuritySettlAgentContactPhone)
}

//HasCashSettlAgentName returns true if CashSettlAgentName is present, Tag 182
func (m SettlementInstructions) HasCashSettlAgentName() bool {
	return m.Has(tag.CashSettlAgentName)
}

//HasCashSettlAgentCode returns true if CashSettlAgentCode is present, Tag 183
func (m SettlementInstructions) HasCashSettlAgentCode() bool {
	return m.Has(tag.CashSettlAgentCode)
}

//HasCashSettlAgentAcctNum returns true if CashSettlAgentAcctNum is present, Tag 184
func (m SettlementInstructions) HasCashSettlAgentAcctNum() bool {
	return m.Has(tag.CashSettlAgentAcctNum)
}

//HasCashSettlAgentAcctName returns true if CashSettlAgentAcctName is present, Tag 185
func (m SettlementInstructions) HasCashSettlAgentAcctName() bool {
	return m.Has(tag.CashSettlAgentAcctName)
}

//HasCashSettlAgentContactName returns true if CashSettlAgentContactName is present, Tag 186
func (m SettlementInstructions) HasCashSettlAgentContactName() bool {
	return m.Has(tag.CashSettlAgentContactName)
}

//HasCashSettlAgentContactPhone returns true if CashSettlAgentContactPhone is present, Tag 187
func (m SettlementInstructions) HasCashSettlAgentContactPhone() bool {
	return m.Has(tag.CashSettlAgentContactPhone)
}

//HasSettlInstRefID returns true if SettlInstRefID is present, Tag 214
func (m SettlementInstructions) HasSettlInstRefID() bool {
	return m.Has(tag.SettlInstRefID)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m SettlementInstructions) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m SettlementInstructions) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasIndividualAllocID returns true if IndividualAllocID is present, Tag 467
func (m SettlementInstructions) HasIndividualAllocID() bool {
	return m.Has(tag.IndividualAllocID)
}

//HasPaymentRef returns true if PaymentRef is present, Tag 476
func (m SettlementInstructions) HasPaymentRef() bool {
	return m.Has(tag.PaymentRef)
}

//HasCardHolderName returns true if CardHolderName is present, Tag 488
func (m SettlementInstructions) HasCardHolderName() bool {
	return m.Has(tag.CardHolderName)
}

//HasCardNumber returns true if CardNumber is present, Tag 489
func (m SettlementInstructions) HasCardNumber() bool {
	return m.Has(tag.CardNumber)
}

//HasCardExpDate returns true if CardExpDate is present, Tag 490
func (m SettlementInstructions) HasCardExpDate() bool {
	return m.Has(tag.CardExpDate)
}

//HasCardIssNo returns true if CardIssNo is present, Tag 491
func (m SettlementInstructions) HasCardIssNo() bool {
	return m.Has(tag.CardIssNo)
}

//HasPaymentMethod returns true if PaymentMethod is present, Tag 492
func (m SettlementInstructions) HasPaymentMethod() bool {
	return m.Has(tag.PaymentMethod)
}

//HasCardStartDate returns true if CardStartDate is present, Tag 503
func (m SettlementInstructions) HasCardStartDate() bool {
	return m.Has(tag.CardStartDate)
}

//HasPaymentDate returns true if PaymentDate is present, Tag 504
func (m SettlementInstructions) HasPaymentDate() bool {
	return m.Has(tag.PaymentDate)
}

//HasPaymentRemitterID returns true if PaymentRemitterID is present, Tag 505
func (m SettlementInstructions) HasPaymentRemitterID() bool {
	return m.Has(tag.PaymentRemitterID)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m SettlementInstructions) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//NoPartyIDs is a repeating group element, Tag 453
type NoPartyIDs struct {
	*quickfix.Group
}

//SetPartyID sets PartyID, Tag 448
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

//SetPartyIDSource sets PartyIDSource, Tag 447
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

//SetPartyRole sets PartyRole, Tag 452
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartyIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//GetPartyID gets PartyID, Tag 448
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyIDSource gets PartyIDSource, Tag 447
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyRole gets PartyRole, Tag 452
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartyIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasPartyID returns true if PartyID is present, Tag 448
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

//HasPartyIDSource returns true if PartyIDSource is present, Tag 447
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

//HasPartyRole returns true if PartyRole is present, Tag 452
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

//HasPartySubID returns true if PartySubID is present, Tag 523
func (m NoPartyIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

//NoPartyIDsRepeatingGroup is a repeating group, Tag 453
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), quickfix.GroupElement(tag.PartySubID)})}
}

//Add create and append a new NoPartyIDs to this group
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

//Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}
