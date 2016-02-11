//Package ordercancelreplacerequest msg type = G.
package ordercancelreplacerequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//NoAllocs is a repeating group in OrderCancelReplaceRequest
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocShares is a non-required field for NoAllocs.
	AllocShares *float64 `fix:"80"`
}

//NoTradingSessions is a repeating group in OrderCancelReplaceRequest
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
}

//Message is a OrderCancelReplaceRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"G"`
	Header     fix42.Header
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
	//NoAllocs is a non-required field for OrderCancelReplaceRequest.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	//SettlmntTyp is a non-required field for OrderCancelReplaceRequest.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for OrderCancelReplaceRequest.
	FutSettDate *string `fix:"64"`
	//HandlInst is a required field for OrderCancelReplaceRequest.
	HandlInst string `fix:"21"`
	//ExecInst is a non-required field for OrderCancelReplaceRequest.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for OrderCancelReplaceRequest.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for OrderCancelReplaceRequest.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for OrderCancelReplaceRequest.
	ExDestination *string `fix:"100"`
	//NoTradingSessions is a non-required field for OrderCancelReplaceRequest.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
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
	//ContractMultiplier is a non-required field for OrderCancelReplaceRequest.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for OrderCancelReplaceRequest.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for OrderCancelReplaceRequest.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for OrderCancelReplaceRequest.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for OrderCancelReplaceRequest.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for OrderCancelReplaceRequest.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for OrderCancelReplaceRequest.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for OrderCancelReplaceRequest.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for OrderCancelReplaceRequest.
	EncodedSecurityDesc *string `fix:"351"`
	//Side is a required field for OrderCancelReplaceRequest.
	Side string `fix:"54"`
	//TransactTime is a required field for OrderCancelReplaceRequest.
	TransactTime time.Time `fix:"60"`
	//OrderQty is a non-required field for OrderCancelReplaceRequest.
	OrderQty *float64 `fix:"38"`
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
	//DiscretionInst is a non-required field for OrderCancelReplaceRequest.
	DiscretionInst *string `fix:"388"`
	//DiscretionOffset is a non-required field for OrderCancelReplaceRequest.
	DiscretionOffset *float64 `fix:"389"`
	//ComplianceID is a non-required field for OrderCancelReplaceRequest.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for OrderCancelReplaceRequest.
	SolicitedFlag *bool `fix:"377"`
	//Currency is a non-required field for OrderCancelReplaceRequest.
	Currency *string `fix:"15"`
	//TimeInForce is a non-required field for OrderCancelReplaceRequest.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for OrderCancelReplaceRequest.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for OrderCancelReplaceRequest.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for OrderCancelReplaceRequest.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for OrderCancelReplaceRequest.
	GTBookingInst *int `fix:"427"`
	//Commission is a non-required field for OrderCancelReplaceRequest.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for OrderCancelReplaceRequest.
	CommType *string `fix:"13"`
	//Rule80A is a non-required field for OrderCancelReplaceRequest.
	Rule80A *string `fix:"47"`
	//ForexReq is a non-required field for OrderCancelReplaceRequest.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for OrderCancelReplaceRequest.
	SettlCurrency *string `fix:"120"`
	//Text is a non-required field for OrderCancelReplaceRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for OrderCancelReplaceRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for OrderCancelReplaceRequest.
	EncodedText *string `fix:"355"`
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
	MaxShow *float64 `fix:"210"`
	//LocateReqd is a non-required field for OrderCancelReplaceRequest.
	LocateReqd *bool `fix:"114"`
	//ClearingFirm is a non-required field for OrderCancelReplaceRequest.
	ClearingFirm *string `fix:"439"`
	//ClearingAccount is a non-required field for OrderCancelReplaceRequest.
	ClearingAccount *string `fix:"440"`
	Trailer         fix42.Trailer
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
	return enum.BeginStringFIX42, "G", r
}
