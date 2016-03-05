//Package settlementinstructions msg type = T.
package settlementinstructions

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix41"
	"time"
)

//Message is a SettlementInstructions FIX Message
type Message struct {
	FIXMsgType string `fix:"T"`
	fix41.Header
	//SettlInstID is a required field for SettlementInstructions.
	SettlInstID string `fix:"162"`
	//SettlInstTransType is a required field for SettlementInstructions.
	SettlInstTransType string `fix:"163"`
	//SettlInstMode is a required field for SettlementInstructions.
	SettlInstMode string `fix:"160"`
	//SettlInstSource is a required field for SettlementInstructions.
	SettlInstSource string `fix:"165"`
	//AllocAccount is a required field for SettlementInstructions.
	AllocAccount string `fix:"79"`
	//SettlLocation is a non-required field for SettlementInstructions.
	SettlLocation *string `fix:"166"`
	//TradeDate is a non-required field for SettlementInstructions.
	TradeDate *string `fix:"75"`
	//AllocID is a non-required field for SettlementInstructions.
	AllocID *string `fix:"70"`
	//LastMkt is a non-required field for SettlementInstructions.
	LastMkt *string `fix:"30"`
	//Side is a non-required field for SettlementInstructions.
	Side *string `fix:"54"`
	//SecurityType is a non-required field for SettlementInstructions.
	SecurityType *string `fix:"167"`
	//EffectiveTime is a non-required field for SettlementInstructions.
	EffectiveTime *time.Time `fix:"168"`
	//TransactTime is a required field for SettlementInstructions.
	TransactTime time.Time `fix:"60"`
	//ClientID is a non-required field for SettlementInstructions.
	ClientID *string `fix:"109"`
	//ExecBroker is a non-required field for SettlementInstructions.
	ExecBroker *string `fix:"76"`
	//StandInstDbType is a non-required field for SettlementInstructions.
	StandInstDbType *int `fix:"169"`
	//StandInstDbName is a non-required field for SettlementInstructions.
	StandInstDbName *string `fix:"170"`
	//StandInstDbID is a non-required field for SettlementInstructions.
	StandInstDbID *string `fix:"171"`
	//SettlDeliveryType is a non-required field for SettlementInstructions.
	SettlDeliveryType *int `fix:"172"`
	//SettlDepositoryCode is a non-required field for SettlementInstructions.
	SettlDepositoryCode *string `fix:"173"`
	//SettlBrkrCode is a non-required field for SettlementInstructions.
	SettlBrkrCode *string `fix:"174"`
	//SettlInstCode is a non-required field for SettlementInstructions.
	SettlInstCode *string `fix:"175"`
	//SecuritySettlAgentName is a non-required field for SettlementInstructions.
	SecuritySettlAgentName *string `fix:"176"`
	//SecuritySettlAgentCode is a non-required field for SettlementInstructions.
	SecuritySettlAgentCode *string `fix:"177"`
	//SecuritySettlAgentAcctNum is a non-required field for SettlementInstructions.
	SecuritySettlAgentAcctNum *string `fix:"178"`
	//SecuritySettlAgentAcctName is a non-required field for SettlementInstructions.
	SecuritySettlAgentAcctName *string `fix:"179"`
	//SecuritySettlAgentContactName is a non-required field for SettlementInstructions.
	SecuritySettlAgentContactName *string `fix:"180"`
	//SecuritySettlAgentContactPhone is a non-required field for SettlementInstructions.
	SecuritySettlAgentContactPhone *string `fix:"181"`
	//CashSettlAgentName is a non-required field for SettlementInstructions.
	CashSettlAgentName *string `fix:"182"`
	//CashSettlAgentCode is a non-required field for SettlementInstructions.
	CashSettlAgentCode *string `fix:"183"`
	//CashSettlAgentAcctNum is a non-required field for SettlementInstructions.
	CashSettlAgentAcctNum *string `fix:"184"`
	//CashSettlAgentAcctName is a non-required field for SettlementInstructions.
	CashSettlAgentAcctName *string `fix:"185"`
	//CashSettlAgentContactName is a non-required field for SettlementInstructions.
	CashSettlAgentContactName *string `fix:"186"`
	//CashSettlAgentContactPhone is a non-required field for SettlementInstructions.
	CashSettlAgentContactPhone *string `fix:"187"`
	fix41.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetSettlInstID(v string)                    { m.SettlInstID = v }
func (m *Message) SetSettlInstTransType(v string)             { m.SettlInstTransType = v }
func (m *Message) SetSettlInstMode(v string)                  { m.SettlInstMode = v }
func (m *Message) SetSettlInstSource(v string)                { m.SettlInstSource = v }
func (m *Message) SetAllocAccount(v string)                   { m.AllocAccount = v }
func (m *Message) SetSettlLocation(v string)                  { m.SettlLocation = &v }
func (m *Message) SetTradeDate(v string)                      { m.TradeDate = &v }
func (m *Message) SetAllocID(v string)                        { m.AllocID = &v }
func (m *Message) SetLastMkt(v string)                        { m.LastMkt = &v }
func (m *Message) SetSide(v string)                           { m.Side = &v }
func (m *Message) SetSecurityType(v string)                   { m.SecurityType = &v }
func (m *Message) SetEffectiveTime(v time.Time)               { m.EffectiveTime = &v }
func (m *Message) SetTransactTime(v time.Time)                { m.TransactTime = v }
func (m *Message) SetClientID(v string)                       { m.ClientID = &v }
func (m *Message) SetExecBroker(v string)                     { m.ExecBroker = &v }
func (m *Message) SetStandInstDbType(v int)                   { m.StandInstDbType = &v }
func (m *Message) SetStandInstDbName(v string)                { m.StandInstDbName = &v }
func (m *Message) SetStandInstDbID(v string)                  { m.StandInstDbID = &v }
func (m *Message) SetSettlDeliveryType(v int)                 { m.SettlDeliveryType = &v }
func (m *Message) SetSettlDepositoryCode(v string)            { m.SettlDepositoryCode = &v }
func (m *Message) SetSettlBrkrCode(v string)                  { m.SettlBrkrCode = &v }
func (m *Message) SetSettlInstCode(v string)                  { m.SettlInstCode = &v }
func (m *Message) SetSecuritySettlAgentName(v string)         { m.SecuritySettlAgentName = &v }
func (m *Message) SetSecuritySettlAgentCode(v string)         { m.SecuritySettlAgentCode = &v }
func (m *Message) SetSecuritySettlAgentAcctNum(v string)      { m.SecuritySettlAgentAcctNum = &v }
func (m *Message) SetSecuritySettlAgentAcctName(v string)     { m.SecuritySettlAgentAcctName = &v }
func (m *Message) SetSecuritySettlAgentContactName(v string)  { m.SecuritySettlAgentContactName = &v }
func (m *Message) SetSecuritySettlAgentContactPhone(v string) { m.SecuritySettlAgentContactPhone = &v }
func (m *Message) SetCashSettlAgentName(v string)             { m.CashSettlAgentName = &v }
func (m *Message) SetCashSettlAgentCode(v string)             { m.CashSettlAgentCode = &v }
func (m *Message) SetCashSettlAgentAcctNum(v string)          { m.CashSettlAgentAcctNum = &v }
func (m *Message) SetCashSettlAgentAcctName(v string)         { m.CashSettlAgentAcctName = &v }
func (m *Message) SetCashSettlAgentContactName(v string)      { m.CashSettlAgentContactName = &v }
func (m *Message) SetCashSettlAgentContactPhone(v string)     { m.CashSettlAgentContactPhone = &v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX41, "T", r
}
