//Package settlementinstructionrequest msg type = AV.
package settlementinstructionrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a SettlementInstructionRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"AV"`
	Header     fixt11.Header
	//SettlInstReqID is a required field for SettlementInstructionRequest.
	SettlInstReqID string `fix:"791"`
	//TransactTime is a required field for SettlementInstructionRequest.
	TransactTime time.Time `fix:"60"`
	//Parties Component
	Parties parties.Component
	//AllocAccount is a non-required field for SettlementInstructionRequest.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for SettlementInstructionRequest.
	AllocAcctIDSource *int `fix:"661"`
	//Side is a non-required field for SettlementInstructionRequest.
	Side *string `fix:"54"`
	//Product is a non-required field for SettlementInstructionRequest.
	Product *int `fix:"460"`
	//SecurityType is a non-required field for SettlementInstructionRequest.
	SecurityType *string `fix:"167"`
	//CFICode is a non-required field for SettlementInstructionRequest.
	CFICode *string `fix:"461"`
	//EffectiveTime is a non-required field for SettlementInstructionRequest.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireTime is a non-required field for SettlementInstructionRequest.
	ExpireTime *time.Time `fix:"126"`
	//LastUpdateTime is a non-required field for SettlementInstructionRequest.
	LastUpdateTime *time.Time `fix:"779"`
	//StandInstDbType is a non-required field for SettlementInstructionRequest.
	StandInstDbType *int `fix:"169"`
	//StandInstDbName is a non-required field for SettlementInstructionRequest.
	StandInstDbName *string `fix:"170"`
	//StandInstDbID is a non-required field for SettlementInstructionRequest.
	StandInstDbID *string `fix:"171"`
	//SettlCurrency is a non-required field for SettlementInstructionRequest.
	SettlCurrency *string `fix:"120"`
	Trailer       fixt11.Trailer
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
	return enum.BeginStringFIX50, "AV", r
}
