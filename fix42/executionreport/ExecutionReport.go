//Package executionreport msg type = 8.
package executionreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//NoContraBrokers is a repeating group in ExecutionReport
type NoContraBrokers struct {
	//ContraBroker is a non-required field for NoContraBrokers.
	ContraBroker *string `fix:"375"`
	//ContraTrader is a non-required field for NoContraBrokers.
	ContraTrader *string `fix:"337"`
	//ContraTradeQty is a non-required field for NoContraBrokers.
	ContraTradeQty *float64 `fix:"437"`
	//ContraTradeTime is a non-required field for NoContraBrokers.
	ContraTradeTime *time.Time `fix:"438"`
}

//Message is a ExecutionReport FIX Message
type Message struct {
	FIXMsgType string `fix:"8"`
	Header     fix42.Header
	//OrderID is a required field for ExecutionReport.
	OrderID string `fix:"37"`
	//SecondaryOrderID is a non-required field for ExecutionReport.
	SecondaryOrderID *string `fix:"198"`
	//ClOrdID is a non-required field for ExecutionReport.
	ClOrdID *string `fix:"11"`
	//OrigClOrdID is a non-required field for ExecutionReport.
	OrigClOrdID *string `fix:"41"`
	//ClientID is a non-required field for ExecutionReport.
	ClientID *string `fix:"109"`
	//ExecBroker is a non-required field for ExecutionReport.
	ExecBroker *string `fix:"76"`
	//NoContraBrokers is a non-required field for ExecutionReport.
	NoContraBrokers []NoContraBrokers `fix:"382,omitempty"`
	//ListID is a non-required field for ExecutionReport.
	ListID *string `fix:"66"`
	//ExecID is a required field for ExecutionReport.
	ExecID string `fix:"17"`
	//ExecTransType is a required field for ExecutionReport.
	ExecTransType string `fix:"20"`
	//ExecRefID is a non-required field for ExecutionReport.
	ExecRefID *string `fix:"19"`
	//ExecType is a required field for ExecutionReport.
	ExecType string `fix:"150"`
	//OrdStatus is a required field for ExecutionReport.
	OrdStatus string `fix:"39"`
	//OrdRejReason is a non-required field for ExecutionReport.
	OrdRejReason *int `fix:"103"`
	//ExecRestatementReason is a non-required field for ExecutionReport.
	ExecRestatementReason *int `fix:"378"`
	//Account is a non-required field for ExecutionReport.
	Account *string `fix:"1"`
	//SettlmntTyp is a non-required field for ExecutionReport.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for ExecutionReport.
	FutSettDate *string `fix:"64"`
	//Symbol is a required field for ExecutionReport.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for ExecutionReport.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for ExecutionReport.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for ExecutionReport.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for ExecutionReport.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for ExecutionReport.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for ExecutionReport.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for ExecutionReport.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for ExecutionReport.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for ExecutionReport.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for ExecutionReport.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for ExecutionReport.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for ExecutionReport.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for ExecutionReport.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for ExecutionReport.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for ExecutionReport.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for ExecutionReport.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for ExecutionReport.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for ExecutionReport.
	EncodedSecurityDesc *string `fix:"351"`
	//Side is a required field for ExecutionReport.
	Side string `fix:"54"`
	//OrderQty is a non-required field for ExecutionReport.
	OrderQty *float64 `fix:"38"`
	//CashOrderQty is a non-required field for ExecutionReport.
	CashOrderQty *float64 `fix:"152"`
	//OrdType is a non-required field for ExecutionReport.
	OrdType *string `fix:"40"`
	//Price is a non-required field for ExecutionReport.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for ExecutionReport.
	StopPx *float64 `fix:"99"`
	//PegDifference is a non-required field for ExecutionReport.
	PegDifference *float64 `fix:"211"`
	//DiscretionInst is a non-required field for ExecutionReport.
	DiscretionInst *string `fix:"388"`
	//DiscretionOffset is a non-required field for ExecutionReport.
	DiscretionOffset *float64 `fix:"389"`
	//Currency is a non-required field for ExecutionReport.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for ExecutionReport.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for ExecutionReport.
	SolicitedFlag *bool `fix:"377"`
	//TimeInForce is a non-required field for ExecutionReport.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for ExecutionReport.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for ExecutionReport.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for ExecutionReport.
	ExpireTime *time.Time `fix:"126"`
	//ExecInst is a non-required field for ExecutionReport.
	ExecInst *string `fix:"18"`
	//Rule80A is a non-required field for ExecutionReport.
	Rule80A *string `fix:"47"`
	//LastShares is a non-required field for ExecutionReport.
	LastShares *float64 `fix:"32"`
	//LastPx is a non-required field for ExecutionReport.
	LastPx *float64 `fix:"31"`
	//LastSpotRate is a non-required field for ExecutionReport.
	LastSpotRate *float64 `fix:"194"`
	//LastForwardPoints is a non-required field for ExecutionReport.
	LastForwardPoints *float64 `fix:"195"`
	//LastMkt is a non-required field for ExecutionReport.
	LastMkt *string `fix:"30"`
	//TradingSessionID is a non-required field for ExecutionReport.
	TradingSessionID *string `fix:"336"`
	//LastCapacity is a non-required field for ExecutionReport.
	LastCapacity *string `fix:"29"`
	//LeavesQty is a required field for ExecutionReport.
	LeavesQty float64 `fix:"151"`
	//CumQty is a required field for ExecutionReport.
	CumQty float64 `fix:"14"`
	//AvgPx is a required field for ExecutionReport.
	AvgPx float64 `fix:"6"`
	//DayOrderQty is a non-required field for ExecutionReport.
	DayOrderQty *float64 `fix:"424"`
	//DayCumQty is a non-required field for ExecutionReport.
	DayCumQty *float64 `fix:"425"`
	//DayAvgPx is a non-required field for ExecutionReport.
	DayAvgPx *float64 `fix:"426"`
	//GTBookingInst is a non-required field for ExecutionReport.
	GTBookingInst *int `fix:"427"`
	//TradeDate is a non-required field for ExecutionReport.
	TradeDate *string `fix:"75"`
	//TransactTime is a non-required field for ExecutionReport.
	TransactTime *time.Time `fix:"60"`
	//ReportToExch is a non-required field for ExecutionReport.
	ReportToExch *bool `fix:"113"`
	//Commission is a non-required field for ExecutionReport.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for ExecutionReport.
	CommType *string `fix:"13"`
	//GrossTradeAmt is a non-required field for ExecutionReport.
	GrossTradeAmt *float64 `fix:"381"`
	//SettlCurrAmt is a non-required field for ExecutionReport.
	SettlCurrAmt *float64 `fix:"119"`
	//SettlCurrency is a non-required field for ExecutionReport.
	SettlCurrency *string `fix:"120"`
	//SettlCurrFxRate is a non-required field for ExecutionReport.
	SettlCurrFxRate *float64 `fix:"155"`
	//SettlCurrFxRateCalc is a non-required field for ExecutionReport.
	SettlCurrFxRateCalc *string `fix:"156"`
	//HandlInst is a non-required field for ExecutionReport.
	HandlInst *string `fix:"21"`
	//MinQty is a non-required field for ExecutionReport.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for ExecutionReport.
	MaxFloor *float64 `fix:"111"`
	//OpenClose is a non-required field for ExecutionReport.
	OpenClose *string `fix:"77"`
	//MaxShow is a non-required field for ExecutionReport.
	MaxShow *float64 `fix:"210"`
	//Text is a non-required field for ExecutionReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for ExecutionReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for ExecutionReport.
	EncodedText *string `fix:"355"`
	//FutSettDate2 is a non-required field for ExecutionReport.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for ExecutionReport.
	OrderQty2 *float64 `fix:"192"`
	//ClearingFirm is a non-required field for ExecutionReport.
	ClearingFirm *string `fix:"439"`
	//ClearingAccount is a non-required field for ExecutionReport.
	ClearingAccount *string `fix:"440"`
	//MultiLegReportingType is a non-required field for ExecutionReport.
	MultiLegReportingType *string `fix:"442"`
	Trailer               fix42.Trailer
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
	return enum.BeginStringFIX42, "8", r
}
