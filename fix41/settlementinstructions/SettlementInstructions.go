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
	Header     fix41.Header
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
	Trailer                    fix41.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
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
