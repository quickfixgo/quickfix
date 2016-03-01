//Package ordercancelreplacerequest msg type = G.
package ordercancelreplacerequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix41"
	"time"
)

//Message is a OrderCancelReplaceRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"G"`
	Header     fix41.Header
	//OrderID is a non-required field for OrderCancelReplaceRequest.
	OrderID *string `fix:"37"`
	//ClientID is a non-required field for OrderCancelReplaceRequest.
	ClientID *string `fix:"109"`
	//ExecBroker is a non-required field for OrderCancelReplaceRequest.
	ExecBroker *string `fix:"76"`
	//OrigClOrdID is a required field for OrderCancelReplaceRequest.
	OrigClOrdID string `fix:"41"`
	//ClOrdID is a required field for OrderCancelReplaceRequest.
	ClOrdID string `fix:"11"`
	//ListID is a non-required field for OrderCancelReplaceRequest.
	ListID *string `fix:"66"`
	//Account is a non-required field for OrderCancelReplaceRequest.
	Account *string `fix:"1"`
	//SettlmntTyp is a non-required field for OrderCancelReplaceRequest.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for OrderCancelReplaceRequest.
	FutSettDate *string `fix:"64"`
	//HandlInst is a required field for OrderCancelReplaceRequest.
	HandlInst string `fix:"21"`
	//ExecInst is a non-required field for OrderCancelReplaceRequest.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for OrderCancelReplaceRequest.
	MinQty *int `fix:"110"`
	//MaxFloor is a non-required field for OrderCancelReplaceRequest.
	MaxFloor *int `fix:"111"`
	//ExDestination is a non-required field for OrderCancelReplaceRequest.
	ExDestination *string `fix:"100"`
	//Symbol is a required field for OrderCancelReplaceRequest.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for OrderCancelReplaceRequest.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for OrderCancelReplaceRequest.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for OrderCancelReplaceRequest.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for OrderCancelReplaceRequest.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for OrderCancelReplaceRequest.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for OrderCancelReplaceRequest.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for OrderCancelReplaceRequest.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for OrderCancelReplaceRequest.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for OrderCancelReplaceRequest.
	OptAttribute *string `fix:"206"`
	//SecurityExchange is a non-required field for OrderCancelReplaceRequest.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for OrderCancelReplaceRequest.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for OrderCancelReplaceRequest.
	SecurityDesc *string `fix:"107"`
	//Side is a required field for OrderCancelReplaceRequest.
	Side string `fix:"54"`
	//OrderQty is a non-required field for OrderCancelReplaceRequest.
	OrderQty *int `fix:"38"`
	//CashOrderQty is a non-required field for OrderCancelReplaceRequest.
	CashOrderQty *float64 `fix:"152"`
	//OrdType is a required field for OrderCancelReplaceRequest.
	OrdType string `fix:"40"`
	//Price is a non-required field for OrderCancelReplaceRequest.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for OrderCancelReplaceRequest.
	StopPx *float64 `fix:"99"`
	//PegDifference is a non-required field for OrderCancelReplaceRequest.
	PegDifference *float64 `fix:"211"`
	//Currency is a non-required field for OrderCancelReplaceRequest.
	Currency *string `fix:"15"`
	//TimeInForce is a non-required field for OrderCancelReplaceRequest.
	TimeInForce *string `fix:"59"`
	//ExpireTime is a non-required field for OrderCancelReplaceRequest.
	ExpireTime *time.Time `fix:"126"`
	//Commission is a non-required field for OrderCancelReplaceRequest.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for OrderCancelReplaceRequest.
	CommType *string `fix:"13"`
	//Rule80A is a non-required field for OrderCancelReplaceRequest.
	Rule80A *string `fix:"47"`
	//ForexReq is a non-required field for OrderCancelReplaceRequest.
	ForexReq *string `fix:"121"`
	//SettlCurrency is a non-required field for OrderCancelReplaceRequest.
	SettlCurrency *string `fix:"120"`
	//Text is a non-required field for OrderCancelReplaceRequest.
	Text *string `fix:"58"`
	//FutSettDate2 is a non-required field for OrderCancelReplaceRequest.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for OrderCancelReplaceRequest.
	OrderQty2 *float64 `fix:"192"`
	//OpenClose is a non-required field for OrderCancelReplaceRequest.
	OpenClose *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for OrderCancelReplaceRequest.
	CoveredOrUncovered *int `fix:"203"`
	//CustomerOrFirm is a non-required field for OrderCancelReplaceRequest.
	CustomerOrFirm *int `fix:"204"`
	//MaxShow is a non-required field for OrderCancelReplaceRequest.
	MaxShow *int `fix:"210"`
	//LocateReqd is a non-required field for OrderCancelReplaceRequest.
	LocateReqd *string `fix:"114"`
	Trailer    fix41.Trailer
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
	return enum.BeginStringFIX41, "G", r
}
