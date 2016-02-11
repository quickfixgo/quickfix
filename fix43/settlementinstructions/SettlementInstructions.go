//Package settlementinstructions msg type = T.
package settlementinstructions

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"time"
)

//Message is a SettlementInstructions FIX Message
type Message struct {
	FIXMsgType string `fix:"T"`
	Header     fix43.Header
	//SettlInstID is a required field for SettlementInstructions.
	SettlInstID string `fix:"162"`
	//SettlInstTransType is a required field for SettlementInstructions.
	SettlInstTransType string `fix:"163"`
	//SettlInstRefID is a required field for SettlementInstructions.
	SettlInstRefID string `fix:"214"`
	//SettlInstMode is a required field for SettlementInstructions.
	SettlInstMode string `fix:"160"`
	//SettlInstSource is a required field for SettlementInstructions.
	SettlInstSource string `fix:"165"`
	//AllocAccount is a required field for SettlementInstructions.
	AllocAccount string `fix:"79"`
	//IndividualAllocID is a non-required field for SettlementInstructions.
	IndividualAllocID *string `fix:"467"`
	//ClOrdID is a non-required field for SettlementInstructions.
	ClOrdID *string `fix:"11"`
	//TradeDate is a non-required field for SettlementInstructions.
	TradeDate *string `fix:"75"`
	//AllocID is a non-required field for SettlementInstructions.
	AllocID *string `fix:"70"`
	//LastMkt is a non-required field for SettlementInstructions.
	LastMkt *string `fix:"30"`
	//TradingSessionID is a non-required field for SettlementInstructions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for SettlementInstructions.
	TradingSessionSubID *string `fix:"625"`
	//Side is a non-required field for SettlementInstructions.
	Side *string `fix:"54"`
	//SecurityType is a non-required field for SettlementInstructions.
	SecurityType *string `fix:"167"`
	//EffectiveTime is a non-required field for SettlementInstructions.
	EffectiveTime *time.Time `fix:"168"`
	//TransactTime is a required field for SettlementInstructions.
	TransactTime time.Time `fix:"60"`
	//Parties Component
	Parties parties.Component
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
	//PaymentMethod is a non-required field for SettlementInstructions.
	PaymentMethod *int `fix:"492"`
	//PaymentRef is a non-required field for SettlementInstructions.
	PaymentRef *string `fix:"476"`
	//CardHolderName is a non-required field for SettlementInstructions.
	CardHolderName *string `fix:"488"`
	//CardNumber is a non-required field for SettlementInstructions.
	CardNumber *string `fix:"489"`
	//CardStartDate is a non-required field for SettlementInstructions.
	CardStartDate *string `fix:"503"`
	//CardExpDate is a non-required field for SettlementInstructions.
	CardExpDate *string `fix:"490"`
	//CardIssNo is a non-required field for SettlementInstructions.
	CardIssNo *string `fix:"491"`
	//PaymentDate is a non-required field for SettlementInstructions.
	PaymentDate *string `fix:"504"`
	//PaymentRemitterID is a non-required field for SettlementInstructions.
	PaymentRemitterID *string `fix:"505"`
	Trailer           fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

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
	return enum.BeginStringFIX43, "T", r
}
