//Package tradecapturereport msg type = AE.
package tradecapturereport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/commissiondata"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/legstipulations"
	"github.com/quickfixgo/quickfix/fix44/nestedparties"
	"github.com/quickfixgo/quickfix/fix44/nestedparties2"
	"github.com/quickfixgo/quickfix/fix44/orderqtydata"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/positionamountdata"
	"github.com/quickfixgo/quickfix/fix44/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix44/stipulations"
	"github.com/quickfixgo/quickfix/fix44/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fix44/yielddata"
	"time"
)

//NoUnderlyings is a repeating group in TradeCaptureReport
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//NoLegs is a repeating group in TradeCaptureReport
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
	//LegQty is a non-required field for NoLegs.
	LegQty *float64 `fix:"687"`
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegStipulations Component
	LegStipulations legstipulations.Component
	//LegPositionEffect is a non-required field for NoLegs.
	LegPositionEffect *string `fix:"564"`
	//LegCoveredOrUncovered is a non-required field for NoLegs.
	LegCoveredOrUncovered *int `fix:"565"`
	//NestedParties Component
	NestedParties nestedparties.Component
	//LegRefID is a non-required field for NoLegs.
	LegRefID *string `fix:"654"`
	//LegPrice is a non-required field for NoLegs.
	LegPrice *float64 `fix:"566"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegSettlDate is a non-required field for NoLegs.
	LegSettlDate *string `fix:"588"`
	//LegLastPx is a non-required field for NoLegs.
	LegLastPx *float64 `fix:"637"`
}

//NoSides is a repeating group in TradeCaptureReport
type NoSides struct {
	//Side is a required field for NoSides.
	Side string `fix:"54"`
	//OrderID is a required field for NoSides.
	OrderID string `fix:"37"`
	//SecondaryOrderID is a non-required field for NoSides.
	SecondaryOrderID *string `fix:"198"`
	//ClOrdID is a non-required field for NoSides.
	ClOrdID *string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoSides.
	SecondaryClOrdID *string `fix:"526"`
	//ListID is a non-required field for NoSides.
	ListID *string `fix:"66"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for NoSides.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for NoSides.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for NoSides.
	AccountType *int `fix:"581"`
	//ProcessCode is a non-required field for NoSides.
	ProcessCode *string `fix:"81"`
	//OddLot is a non-required field for NoSides.
	OddLot *bool `fix:"575"`
	//NoClearingInstructions is a non-required field for NoSides.
	NoClearingInstructions []NoClearingInstructions `fix:"576,omitempty"`
	//ClearingFeeIndicator is a non-required field for NoSides.
	ClearingFeeIndicator *string `fix:"635"`
	//TradeInputSource is a non-required field for NoSides.
	TradeInputSource *string `fix:"578"`
	//TradeInputDevice is a non-required field for NoSides.
	TradeInputDevice *string `fix:"579"`
	//OrderInputDevice is a non-required field for NoSides.
	OrderInputDevice *string `fix:"821"`
	//Currency is a non-required field for NoSides.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for NoSides.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for NoSides.
	SolicitedFlag *bool `fix:"377"`
	//OrderCapacity is a non-required field for NoSides.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for NoSides.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for NoSides.
	CustOrderCapacity *int `fix:"582"`
	//OrdType is a non-required field for NoSides.
	OrdType *string `fix:"40"`
	//ExecInst is a non-required field for NoSides.
	ExecInst *string `fix:"18"`
	//TransBkdTime is a non-required field for NoSides.
	TransBkdTime *time.Time `fix:"483"`
	//TradingSessionID is a non-required field for NoSides.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoSides.
	TradingSessionSubID *string `fix:"625"`
	//TimeBracket is a non-required field for NoSides.
	TimeBracket *string `fix:"943"`
	//CommissionData Component
	CommissionData commissiondata.Component
	//GrossTradeAmt is a non-required field for NoSides.
	GrossTradeAmt *float64 `fix:"381"`
	//NumDaysInterest is a non-required field for NoSides.
	NumDaysInterest *int `fix:"157"`
	//ExDate is a non-required field for NoSides.
	ExDate *string `fix:"230"`
	//AccruedInterestRate is a non-required field for NoSides.
	AccruedInterestRate *float64 `fix:"158"`
	//AccruedInterestAmt is a non-required field for NoSides.
	AccruedInterestAmt *float64 `fix:"159"`
	//InterestAtMaturity is a non-required field for NoSides.
	InterestAtMaturity *float64 `fix:"738"`
	//EndAccruedInterestAmt is a non-required field for NoSides.
	EndAccruedInterestAmt *float64 `fix:"920"`
	//StartCash is a non-required field for NoSides.
	StartCash *float64 `fix:"921"`
	//EndCash is a non-required field for NoSides.
	EndCash *float64 `fix:"922"`
	//Concession is a non-required field for NoSides.
	Concession *float64 `fix:"238"`
	//TotalTakedown is a non-required field for NoSides.
	TotalTakedown *float64 `fix:"237"`
	//NetMoney is a non-required field for NoSides.
	NetMoney *float64 `fix:"118"`
	//SettlCurrAmt is a non-required field for NoSides.
	SettlCurrAmt *float64 `fix:"119"`
	//SettlCurrency is a non-required field for NoSides.
	SettlCurrency *string `fix:"120"`
	//SettlCurrFxRate is a non-required field for NoSides.
	SettlCurrFxRate *float64 `fix:"155"`
	//SettlCurrFxRateCalc is a non-required field for NoSides.
	SettlCurrFxRateCalc *string `fix:"156"`
	//PositionEffect is a non-required field for NoSides.
	PositionEffect *string `fix:"77"`
	//Text is a non-required field for NoSides.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoSides.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoSides.
	EncodedText *string `fix:"355"`
	//SideMultiLegReportingType is a non-required field for NoSides.
	SideMultiLegReportingType *int `fix:"752"`
	//NoContAmts is a non-required field for NoSides.
	NoContAmts []NoContAmts `fix:"518,omitempty"`
	//Stipulations Component
	Stipulations stipulations.Component
	//NoMiscFees is a non-required field for NoSides.
	NoMiscFees []NoMiscFees `fix:"136,omitempty"`
	//ExchangeRule is a non-required field for NoSides.
	ExchangeRule *string `fix:"825"`
	//TradeAllocIndicator is a non-required field for NoSides.
	TradeAllocIndicator *int `fix:"826"`
	//PreallocMethod is a non-required field for NoSides.
	PreallocMethod *string `fix:"591"`
	//AllocID is a non-required field for NoSides.
	AllocID *string `fix:"70"`
	//NoAllocs is a non-required field for NoSides.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	//CopyMsgIndicator is a non-required field for NoSides.
	CopyMsgIndicator *bool `fix:"797"`
	//PublishTrdIndicator is a non-required field for NoSides.
	PublishTrdIndicator *bool `fix:"852"`
	//ShortSaleReason is a non-required field for NoSides.
	ShortSaleReason *int `fix:"853"`
}

//NoClearingInstructions is a repeating group in NoSides
type NoClearingInstructions struct {
	//ClearingInstruction is a non-required field for NoClearingInstructions.
	ClearingInstruction *int `fix:"577"`
}

//NoContAmts is a repeating group in NoSides
type NoContAmts struct {
	//ContAmtType is a non-required field for NoContAmts.
	ContAmtType *int `fix:"519"`
	//ContAmtValue is a non-required field for NoContAmts.
	ContAmtValue *float64 `fix:"520"`
	//ContAmtCurr is a non-required field for NoContAmts.
	ContAmtCurr *string `fix:"521"`
}

//NoMiscFees is a repeating group in NoSides
type NoMiscFees struct {
	//MiscFeeAmt is a non-required field for NoMiscFees.
	MiscFeeAmt *float64 `fix:"137"`
	//MiscFeeCurr is a non-required field for NoMiscFees.
	MiscFeeCurr *string `fix:"138"`
	//MiscFeeType is a non-required field for NoMiscFees.
	MiscFeeType *string `fix:"139"`
	//MiscFeeBasis is a non-required field for NoMiscFees.
	MiscFeeBasis *int `fix:"891"`
}

//NoAllocs is a repeating group in NoSides
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for NoAllocs.
	AllocAcctIDSource *int `fix:"661"`
	//AllocSettlCurrency is a non-required field for NoAllocs.
	AllocSettlCurrency *string `fix:"736"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//NestedParties2 Component
	NestedParties2 nestedparties2.Component
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
}

//Message is a TradeCaptureReport FIX Message
type Message struct {
	FIXMsgType string `fix:"AE"`
	Header     fix44.Header
	//TradeReportID is a required field for TradeCaptureReport.
	TradeReportID string `fix:"571"`
	//TradeReportTransType is a non-required field for TradeCaptureReport.
	TradeReportTransType *int `fix:"487"`
	//TradeReportType is a non-required field for TradeCaptureReport.
	TradeReportType *int `fix:"856"`
	//TradeRequestID is a non-required field for TradeCaptureReport.
	TradeRequestID *string `fix:"568"`
	//TrdType is a non-required field for TradeCaptureReport.
	TrdType *int `fix:"828"`
	//TrdSubType is a non-required field for TradeCaptureReport.
	TrdSubType *int `fix:"829"`
	//SecondaryTrdType is a non-required field for TradeCaptureReport.
	SecondaryTrdType *int `fix:"855"`
	//TransferReason is a non-required field for TradeCaptureReport.
	TransferReason *string `fix:"830"`
	//ExecType is a non-required field for TradeCaptureReport.
	ExecType *string `fix:"150"`
	//TotNumTradeReports is a non-required field for TradeCaptureReport.
	TotNumTradeReports *int `fix:"748"`
	//LastRptRequested is a non-required field for TradeCaptureReport.
	LastRptRequested *bool `fix:"912"`
	//UnsolicitedIndicator is a non-required field for TradeCaptureReport.
	UnsolicitedIndicator *bool `fix:"325"`
	//SubscriptionRequestType is a non-required field for TradeCaptureReport.
	SubscriptionRequestType *string `fix:"263"`
	//TradeReportRefID is a non-required field for TradeCaptureReport.
	TradeReportRefID *string `fix:"572"`
	//SecondaryTradeReportRefID is a non-required field for TradeCaptureReport.
	SecondaryTradeReportRefID *string `fix:"881"`
	//SecondaryTradeReportID is a non-required field for TradeCaptureReport.
	SecondaryTradeReportID *string `fix:"818"`
	//TradeLinkID is a non-required field for TradeCaptureReport.
	TradeLinkID *string `fix:"820"`
	//TrdMatchID is a non-required field for TradeCaptureReport.
	TrdMatchID *string `fix:"880"`
	//ExecID is a non-required field for TradeCaptureReport.
	ExecID *string `fix:"17"`
	//OrdStatus is a non-required field for TradeCaptureReport.
	OrdStatus *string `fix:"39"`
	//SecondaryExecID is a non-required field for TradeCaptureReport.
	SecondaryExecID *string `fix:"527"`
	//ExecRestatementReason is a non-required field for TradeCaptureReport.
	ExecRestatementReason *int `fix:"378"`
	//PreviouslyReported is a required field for TradeCaptureReport.
	PreviouslyReported bool `fix:"570"`
	//PriceType is a non-required field for TradeCaptureReport.
	PriceType *int `fix:"423"`
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//QtyType is a non-required field for TradeCaptureReport.
	QtyType *int `fix:"854"`
	//YieldData Component
	YieldData yielddata.Component
	//NoUnderlyings is a non-required field for TradeCaptureReport.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//UnderlyingTradingSessionID is a non-required field for TradeCaptureReport.
	UnderlyingTradingSessionID *string `fix:"822"`
	//UnderlyingTradingSessionSubID is a non-required field for TradeCaptureReport.
	UnderlyingTradingSessionSubID *string `fix:"823"`
	//LastQty is a required field for TradeCaptureReport.
	LastQty float64 `fix:"32"`
	//LastPx is a required field for TradeCaptureReport.
	LastPx float64 `fix:"31"`
	//LastParPx is a non-required field for TradeCaptureReport.
	LastParPx *float64 `fix:"669"`
	//LastSpotRate is a non-required field for TradeCaptureReport.
	LastSpotRate *float64 `fix:"194"`
	//LastForwardPoints is a non-required field for TradeCaptureReport.
	LastForwardPoints *float64 `fix:"195"`
	//LastMkt is a non-required field for TradeCaptureReport.
	LastMkt *string `fix:"30"`
	//TradeDate is a required field for TradeCaptureReport.
	TradeDate string `fix:"75"`
	//ClearingBusinessDate is a non-required field for TradeCaptureReport.
	ClearingBusinessDate *string `fix:"715"`
	//AvgPx is a non-required field for TradeCaptureReport.
	AvgPx *float64 `fix:"6"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//AvgPxIndicator is a non-required field for TradeCaptureReport.
	AvgPxIndicator *int `fix:"819"`
	//PositionAmountData Component
	PositionAmountData positionamountdata.Component
	//MultiLegReportingType is a non-required field for TradeCaptureReport.
	MultiLegReportingType *string `fix:"442"`
	//TradeLegRefID is a non-required field for TradeCaptureReport.
	TradeLegRefID *string `fix:"824"`
	//NoLegs is a non-required field for TradeCaptureReport.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//TransactTime is a required field for TradeCaptureReport.
	TransactTime time.Time `fix:"60"`
	//TrdRegTimestamps Component
	TrdRegTimestamps trdregtimestamps.Component
	//SettlType is a non-required field for TradeCaptureReport.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for TradeCaptureReport.
	SettlDate *string `fix:"64"`
	//MatchStatus is a non-required field for TradeCaptureReport.
	MatchStatus *string `fix:"573"`
	//MatchType is a non-required field for TradeCaptureReport.
	MatchType *string `fix:"574"`
	//NoSides is a required field for TradeCaptureReport.
	NoSides []NoSides `fix:"552"`
	Trailer fix44.Trailer
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
	return enum.BeginStringFIX44, "AE", r
}
