//Package newordersingle msg type = D.
package newordersingle

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
	"time"
)

//Message is a NewOrderSingle FIX Message
type Message struct {
	FIXMsgType string `fix:"D"`
	Header     fix40.Header
	//ClOrdID is a required field for NewOrderSingle.
	ClOrdID string `fix:"11"`
	//ClientID is a non-required field for NewOrderSingle.
	ClientID *string `fix:"109"`
	//ExecBroker is a non-required field for NewOrderSingle.
	ExecBroker *string `fix:"76"`
	//Account is a non-required field for NewOrderSingle.
	Account *string `fix:"1"`
	//SettlmntTyp is a non-required field for NewOrderSingle.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for NewOrderSingle.
	FutSettDate *string `fix:"64"`
	//HandlInst is a required field for NewOrderSingle.
	HandlInst string `fix:"21"`
	//ExecInst is a non-required field for NewOrderSingle.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for NewOrderSingle.
	MinQty *int `fix:"110"`
	//MaxFloor is a non-required field for NewOrderSingle.
	MaxFloor *int `fix:"111"`
	//ExDestination is a non-required field for NewOrderSingle.
	ExDestination *string `fix:"100"`
	//ProcessCode is a non-required field for NewOrderSingle.
	ProcessCode *string `fix:"81"`
	//Symbol is a required field for NewOrderSingle.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for NewOrderSingle.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for NewOrderSingle.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for NewOrderSingle.
	IDSource *string `fix:"22"`
	//Issuer is a non-required field for NewOrderSingle.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for NewOrderSingle.
	SecurityDesc *string `fix:"107"`
	//PrevClosePx is a non-required field for NewOrderSingle.
	PrevClosePx *float64 `fix:"140"`
	//Side is a required field for NewOrderSingle.
	Side string `fix:"54"`
	//LocateReqd is a non-required field for NewOrderSingle.
	LocateReqd *string `fix:"114"`
	//OrderQty is a required field for NewOrderSingle.
	OrderQty int `fix:"38"`
	//OrdType is a required field for NewOrderSingle.
	OrdType string `fix:"40"`
	//Price is a non-required field for NewOrderSingle.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NewOrderSingle.
	StopPx *float64 `fix:"99"`
	//Currency is a non-required field for NewOrderSingle.
	Currency *string `fix:"15"`
	//IOIid is a non-required field for NewOrderSingle.
	IOIid *int `fix:"23"`
	//QuoteID is a non-required field for NewOrderSingle.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for NewOrderSingle.
	TimeInForce *string `fix:"59"`
	//ExpireTime is a non-required field for NewOrderSingle.
	ExpireTime *time.Time `fix:"126"`
	//Commission is a non-required field for NewOrderSingle.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for NewOrderSingle.
	CommType *string `fix:"13"`
	//Rule80A is a non-required field for NewOrderSingle.
	Rule80A *string `fix:"47"`
	//ForexReq is a non-required field for NewOrderSingle.
	ForexReq *string `fix:"121"`
	//SettlCurrency is a non-required field for NewOrderSingle.
	SettlCurrency *string `fix:"120"`
	//Text is a non-required field for NewOrderSingle.
	Text    *string `fix:"58"`
	Trailer fix40.Trailer
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
	return enum.BeginStringFIX40, "D", r
}
