package settlementinstructions

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//SettlementInstructions is the fix42 SettlementInstructions type, MsgType = T
type SettlementInstructions struct {
	fix42.Header
	quickfix.Body
	fix42.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a SettlementInstructions from a quickfix.Message instance
func FromMessage(m quickfix.Message) SettlementInstructions {
	return SettlementInstructions{
		Header:      fix42.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix42.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m SettlementInstructions) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a SettlementInstructions initialized with the required fields for SettlementInstructions
func New(settlinstid field.SettlInstIDField, settlinsttranstype field.SettlInstTransTypeField, settlinstrefid field.SettlInstRefIDField, settlinstmode field.SettlInstModeField, settlinstsource field.SettlInstSourceField, allocaccount field.AllocAccountField, transacttime field.TransactTimeField) (m SettlementInstructions) {
	m.Header = fix42.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("T"))
	m.Header.Set(field.NewBeginString("FIX.4.2"))
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
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "T", r
}

//SetLastMkt sets LastMkt, Tag 30
func (m SettlementInstructions) SetLastMkt(v string) {
	m.Set(field.NewLastMkt(v))
}

//SetSide sets Side, Tag 54
func (m SettlementInstructions) SetSide(v string) {
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

//SetExecBroker sets ExecBroker, Tag 76
func (m SettlementInstructions) SetExecBroker(v string) {
	m.Set(field.NewExecBroker(v))
}

//SetAllocAccount sets AllocAccount, Tag 79
func (m SettlementInstructions) SetAllocAccount(v string) {
	m.Set(field.NewAllocAccount(v))
}

//SetClientID sets ClientID, Tag 109
func (m SettlementInstructions) SetClientID(v string) {
	m.Set(field.NewClientID(v))
}

//SetSettlInstMode sets SettlInstMode, Tag 160
func (m SettlementInstructions) SetSettlInstMode(v string) {
	m.Set(field.NewSettlInstMode(v))
}

//SetSettlInstID sets SettlInstID, Tag 162
func (m SettlementInstructions) SetSettlInstID(v string) {
	m.Set(field.NewSettlInstID(v))
}

//SetSettlInstTransType sets SettlInstTransType, Tag 163
func (m SettlementInstructions) SetSettlInstTransType(v string) {
	m.Set(field.NewSettlInstTransType(v))
}

//SetSettlInstSource sets SettlInstSource, Tag 165
func (m SettlementInstructions) SetSettlInstSource(v string) {
	m.Set(field.NewSettlInstSource(v))
}

//SetSettlLocation sets SettlLocation, Tag 166
func (m SettlementInstructions) SetSettlLocation(v string) {
	m.Set(field.NewSettlLocation(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m SettlementInstructions) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetEffectiveTime sets EffectiveTime, Tag 168
func (m SettlementInstructions) SetEffectiveTime(v time.Time) {
	m.Set(field.NewEffectiveTime(v))
}

//SetStandInstDbType sets StandInstDbType, Tag 169
func (m SettlementInstructions) SetStandInstDbType(v int) {
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
func (m SettlementInstructions) SetSettlDeliveryType(v int) {
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
func (m SettlementInstructions) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//GetLastMkt gets LastMkt, Tag 30
func (m SettlementInstructions) GetLastMkt() (f field.LastMktField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m SettlementInstructions) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m SettlementInstructions) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocID gets AllocID, Tag 70
func (m SettlementInstructions) GetAllocID() (f field.AllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m SettlementInstructions) GetTradeDate() (f field.TradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecBroker gets ExecBroker, Tag 76
func (m SettlementInstructions) GetExecBroker() (f field.ExecBrokerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocAccount gets AllocAccount, Tag 79
func (m SettlementInstructions) GetAllocAccount() (f field.AllocAccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClientID gets ClientID, Tag 109
func (m SettlementInstructions) GetClientID() (f field.ClientIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlInstMode gets SettlInstMode, Tag 160
func (m SettlementInstructions) GetSettlInstMode() (f field.SettlInstModeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlInstID gets SettlInstID, Tag 162
func (m SettlementInstructions) GetSettlInstID() (f field.SettlInstIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlInstTransType gets SettlInstTransType, Tag 163
func (m SettlementInstructions) GetSettlInstTransType() (f field.SettlInstTransTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlInstSource gets SettlInstSource, Tag 165
func (m SettlementInstructions) GetSettlInstSource() (f field.SettlInstSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlLocation gets SettlLocation, Tag 166
func (m SettlementInstructions) GetSettlLocation() (f field.SettlLocationField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m SettlementInstructions) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEffectiveTime gets EffectiveTime, Tag 168
func (m SettlementInstructions) GetEffectiveTime() (f field.EffectiveTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStandInstDbType gets StandInstDbType, Tag 169
func (m SettlementInstructions) GetStandInstDbType() (f field.StandInstDbTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStandInstDbName gets StandInstDbName, Tag 170
func (m SettlementInstructions) GetStandInstDbName() (f field.StandInstDbNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStandInstDbID gets StandInstDbID, Tag 171
func (m SettlementInstructions) GetStandInstDbID() (f field.StandInstDbIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlDeliveryType gets SettlDeliveryType, Tag 172
func (m SettlementInstructions) GetSettlDeliveryType() (f field.SettlDeliveryTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlDepositoryCode gets SettlDepositoryCode, Tag 173
func (m SettlementInstructions) GetSettlDepositoryCode() (f field.SettlDepositoryCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlBrkrCode gets SettlBrkrCode, Tag 174
func (m SettlementInstructions) GetSettlBrkrCode() (f field.SettlBrkrCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlInstCode gets SettlInstCode, Tag 175
func (m SettlementInstructions) GetSettlInstCode() (f field.SettlInstCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecuritySettlAgentName gets SecuritySettlAgentName, Tag 176
func (m SettlementInstructions) GetSecuritySettlAgentName() (f field.SecuritySettlAgentNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecuritySettlAgentCode gets SecuritySettlAgentCode, Tag 177
func (m SettlementInstructions) GetSecuritySettlAgentCode() (f field.SecuritySettlAgentCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecuritySettlAgentAcctNum gets SecuritySettlAgentAcctNum, Tag 178
func (m SettlementInstructions) GetSecuritySettlAgentAcctNum() (f field.SecuritySettlAgentAcctNumField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecuritySettlAgentAcctName gets SecuritySettlAgentAcctName, Tag 179
func (m SettlementInstructions) GetSecuritySettlAgentAcctName() (f field.SecuritySettlAgentAcctNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecuritySettlAgentContactName gets SecuritySettlAgentContactName, Tag 180
func (m SettlementInstructions) GetSecuritySettlAgentContactName() (f field.SecuritySettlAgentContactNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecuritySettlAgentContactPhone gets SecuritySettlAgentContactPhone, Tag 181
func (m SettlementInstructions) GetSecuritySettlAgentContactPhone() (f field.SecuritySettlAgentContactPhoneField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashSettlAgentName gets CashSettlAgentName, Tag 182
func (m SettlementInstructions) GetCashSettlAgentName() (f field.CashSettlAgentNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashSettlAgentCode gets CashSettlAgentCode, Tag 183
func (m SettlementInstructions) GetCashSettlAgentCode() (f field.CashSettlAgentCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashSettlAgentAcctNum gets CashSettlAgentAcctNum, Tag 184
func (m SettlementInstructions) GetCashSettlAgentAcctNum() (f field.CashSettlAgentAcctNumField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashSettlAgentAcctName gets CashSettlAgentAcctName, Tag 185
func (m SettlementInstructions) GetCashSettlAgentAcctName() (f field.CashSettlAgentAcctNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashSettlAgentContactName gets CashSettlAgentContactName, Tag 186
func (m SettlementInstructions) GetCashSettlAgentContactName() (f field.CashSettlAgentContactNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashSettlAgentContactPhone gets CashSettlAgentContactPhone, Tag 187
func (m SettlementInstructions) GetCashSettlAgentContactPhone() (f field.CashSettlAgentContactPhoneField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlInstRefID gets SettlInstRefID, Tag 214
func (m SettlementInstructions) GetSettlInstRefID() (f field.SettlInstRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m SettlementInstructions) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
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

//HasExecBroker returns true if ExecBroker is present, Tag 76
func (m SettlementInstructions) HasExecBroker() bool {
	return m.Has(tag.ExecBroker)
}

//HasAllocAccount returns true if AllocAccount is present, Tag 79
func (m SettlementInstructions) HasAllocAccount() bool {
	return m.Has(tag.AllocAccount)
}

//HasClientID returns true if ClientID is present, Tag 109
func (m SettlementInstructions) HasClientID() bool {
	return m.Has(tag.ClientID)
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

//HasSettlLocation returns true if SettlLocation is present, Tag 166
func (m SettlementInstructions) HasSettlLocation() bool {
	return m.Has(tag.SettlLocation)
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
