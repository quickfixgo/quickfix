//Package neworderlist msg type = E.
package neworderlist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//NoOrders is a repeating group in NewOrderList
type NoOrders struct {
	//ClOrdID is a required field for NoOrders.
	ClOrdID string `fix:"11"`
	//ListSeqNo is a required field for NoOrders.
	ListSeqNo int `fix:"67"`
	//SettlInstMode is a non-required field for NoOrders.
	SettlInstMode *string `fix:"160"`
	//ClientID is a non-required field for NoOrders.
	ClientID *string `fix:"109"`
	//ExecBroker is a non-required field for NoOrders.
	ExecBroker *string `fix:"76"`
	//Account is a non-required field for NoOrders.
	Account *string `fix:"1"`
	//NoAllocs is a non-required field for NoOrders.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	//SettlmntTyp is a non-required field for NoOrders.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for NoOrders.
	FutSettDate *string `fix:"64"`
	//HandlInst is a non-required field for NoOrders.
	HandlInst *string `fix:"21"`
	//ExecInst is a non-required field for NoOrders.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for NoOrders.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for NoOrders.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for NoOrders.
	ExDestination *string `fix:"100"`
	//NoTradingSessions is a non-required field for NoOrders.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//ProcessCode is a non-required field for NoOrders.
	ProcessCode *string `fix:"81"`
	//Symbol is a required field for NoOrders.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for NoOrders.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for NoOrders.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for NoOrders.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for NoOrders.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for NoOrders.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for NoOrders.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for NoOrders.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for NoOrders.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for NoOrders.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for NoOrders.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for NoOrders.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for NoOrders.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for NoOrders.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for NoOrders.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for NoOrders.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for NoOrders.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for NoOrders.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for NoOrders.
	EncodedSecurityDesc *string `fix:"351"`
	//PrevClosePx is a non-required field for NoOrders.
	PrevClosePx *float64 `fix:"140"`
	//Side is a required field for NoOrders.
	Side string `fix:"54"`
	//SideValueInd is a non-required field for NoOrders.
	SideValueInd *int `fix:"401"`
	//LocateReqd is a non-required field for NoOrders.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a non-required field for NoOrders.
	TransactTime *time.Time `fix:"60"`
	//OrderQty is a non-required field for NoOrders.
	OrderQty *float64 `fix:"38"`
	//CashOrderQty is a non-required field for NoOrders.
	CashOrderQty *float64 `fix:"152"`
	//OrdType is a non-required field for NoOrders.
	OrdType *string `fix:"40"`
	//Price is a non-required field for NoOrders.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NoOrders.
	StopPx *float64 `fix:"99"`
	//Currency is a non-required field for NoOrders.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for NoOrders.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for NoOrders.
	SolicitedFlag *bool `fix:"377"`
	//IOIid is a non-required field for NoOrders.
	IOIid *string `fix:"23"`
	//QuoteID is a non-required field for NoOrders.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for NoOrders.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for NoOrders.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for NoOrders.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for NoOrders.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for NoOrders.
	GTBookingInst *int `fix:"427"`
	//Commission is a non-required field for NoOrders.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for NoOrders.
	CommType *string `fix:"13"`
	//Rule80A is a non-required field for NoOrders.
	Rule80A *string `fix:"47"`
	//ForexReq is a non-required field for NoOrders.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for NoOrders.
	SettlCurrency *string `fix:"120"`
	//Text is a non-required field for NoOrders.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoOrders.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoOrders.
	EncodedText *string `fix:"355"`
	//FutSettDate2 is a non-required field for NoOrders.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NoOrders.
	OrderQty2 *float64 `fix:"192"`
	//OpenClose is a non-required field for NoOrders.
	OpenClose *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for NoOrders.
	CoveredOrUncovered *int `fix:"203"`
	//CustomerOrFirm is a non-required field for NoOrders.
	CustomerOrFirm *int `fix:"204"`
	//MaxShow is a non-required field for NoOrders.
	MaxShow *float64 `fix:"210"`
	//PegDifference is a non-required field for NoOrders.
	PegDifference *float64 `fix:"211"`
	//DiscretionInst is a non-required field for NoOrders.
	DiscretionInst *string `fix:"388"`
	//DiscretionOffset is a non-required field for NoOrders.
	DiscretionOffset *float64 `fix:"389"`
	//ClearingFirm is a non-required field for NoOrders.
	ClearingFirm *string `fix:"439"`
	//ClearingAccount is a non-required field for NoOrders.
	ClearingAccount *string `fix:"440"`
}

//NoAllocs is a repeating group in NoOrders
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocShares is a non-required field for NoAllocs.
	AllocShares *float64 `fix:"80"`
}

//NoTradingSessions is a repeating group in NoOrders
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
}

//Message is a NewOrderList FIX Message
type Message struct {
	FIXMsgType string `fix:"E"`
	Header     fix42.Header
	//ListID is a required field for NewOrderList.
	ListID string `fix:"66"`
	//BidID is a non-required field for NewOrderList.
	BidID *string `fix:"390"`
	//ClientBidID is a non-required field for NewOrderList.
	ClientBidID *string `fix:"391"`
	//ProgRptReqs is a non-required field for NewOrderList.
	ProgRptReqs *int `fix:"414"`
	//BidType is a required field for NewOrderList.
	BidType int `fix:"394"`
	//ProgPeriodInterval is a non-required field for NewOrderList.
	ProgPeriodInterval *int `fix:"415"`
	//ListExecInstType is a non-required field for NewOrderList.
	ListExecInstType *string `fix:"433"`
	//ListExecInst is a non-required field for NewOrderList.
	ListExecInst *string `fix:"69"`
	//EncodedListExecInstLen is a non-required field for NewOrderList.
	EncodedListExecInstLen *int `fix:"352"`
	//EncodedListExecInst is a non-required field for NewOrderList.
	EncodedListExecInst *string `fix:"353"`
	//TotNoOrders is a required field for NewOrderList.
	TotNoOrders int `fix:"68"`
	//NoOrders is a required field for NewOrderList.
	NoOrders []NoOrders `fix:"73"`
	Trailer  fix42.Trailer
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
	return enum.BeginStringFIX42, "E", r
}
