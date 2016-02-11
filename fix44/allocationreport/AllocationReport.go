//Package allocationreport msg type = AS.
package allocationreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/commissiondata"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentextension"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/nestedparties"
	"github.com/quickfixgo/quickfix/fix44/nestedparties2"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/settlinstructionsdata"
	"github.com/quickfixgo/quickfix/fix44/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix44/stipulations"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fix44/yielddata"
	"time"
)

//NoOrders is a repeating group in AllocationReport
type NoOrders struct {
	//ClOrdID is a non-required field for NoOrders.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for NoOrders.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for NoOrders.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for NoOrders.
	SecondaryClOrdID *string `fix:"526"`
	//ListID is a non-required field for NoOrders.
	ListID *string `fix:"66"`
	//NestedParties2 Component
	NestedParties2 nestedparties2.Component
	//OrderQty is a non-required field for NoOrders.
	OrderQty *float64 `fix:"38"`
	//OrderAvgPx is a non-required field for NoOrders.
	OrderAvgPx *float64 `fix:"799"`
	//OrderBookingQty is a non-required field for NoOrders.
	OrderBookingQty *float64 `fix:"800"`
}

//NoExecs is a repeating group in AllocationReport
type NoExecs struct {
	//LastQty is a non-required field for NoExecs.
	LastQty *float64 `fix:"32"`
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
	//SecondaryExecID is a non-required field for NoExecs.
	SecondaryExecID *string `fix:"527"`
	//LastPx is a non-required field for NoExecs.
	LastPx *float64 `fix:"31"`
	//LastParPx is a non-required field for NoExecs.
	LastParPx *float64 `fix:"669"`
	//LastCapacity is a non-required field for NoExecs.
	LastCapacity *string `fix:"29"`
}

//NoUnderlyings is a repeating group in AllocationReport
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//NoLegs is a repeating group in AllocationReport
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
}

//NoAllocs is a repeating group in AllocationReport
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for NoAllocs.
	AllocAcctIDSource *int `fix:"661"`
	//MatchStatus is a non-required field for NoAllocs.
	MatchStatus *string `fix:"573"`
	//AllocPrice is a non-required field for NoAllocs.
	AllocPrice *float64 `fix:"366"`
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//ProcessCode is a non-required field for NoAllocs.
	ProcessCode *string `fix:"81"`
	//NestedParties Component
	NestedParties nestedparties.Component
	//NotifyBrokerOfCredit is a non-required field for NoAllocs.
	NotifyBrokerOfCredit *bool `fix:"208"`
	//AllocHandlInst is a non-required field for NoAllocs.
	AllocHandlInst *int `fix:"209"`
	//AllocText is a non-required field for NoAllocs.
	AllocText *string `fix:"161"`
	//EncodedAllocTextLen is a non-required field for NoAllocs.
	EncodedAllocTextLen *int `fix:"360"`
	//EncodedAllocText is a non-required field for NoAllocs.
	EncodedAllocText *string `fix:"361"`
	//CommissionData Component
	CommissionData commissiondata.Component
	//AllocAvgPx is a non-required field for NoAllocs.
	AllocAvgPx *float64 `fix:"153"`
	//AllocNetMoney is a non-required field for NoAllocs.
	AllocNetMoney *float64 `fix:"154"`
	//SettlCurrAmt is a non-required field for NoAllocs.
	SettlCurrAmt *float64 `fix:"119"`
	//AllocSettlCurrAmt is a non-required field for NoAllocs.
	AllocSettlCurrAmt *float64 `fix:"737"`
	//SettlCurrency is a non-required field for NoAllocs.
	SettlCurrency *string `fix:"120"`
	//AllocSettlCurrency is a non-required field for NoAllocs.
	AllocSettlCurrency *string `fix:"736"`
	//SettlCurrFxRate is a non-required field for NoAllocs.
	SettlCurrFxRate *float64 `fix:"155"`
	//SettlCurrFxRateCalc is a non-required field for NoAllocs.
	SettlCurrFxRateCalc *string `fix:"156"`
	//AllocAccruedInterestAmt is a non-required field for NoAllocs.
	AllocAccruedInterestAmt *float64 `fix:"742"`
	//AllocInterestAtMaturity is a non-required field for NoAllocs.
	AllocInterestAtMaturity *float64 `fix:"741"`
	//NoMiscFees is a non-required field for NoAllocs.
	NoMiscFees []NoMiscFees `fix:"136,omitempty"`
	//NoClearingInstructions is a non-required field for NoAllocs.
	NoClearingInstructions []NoClearingInstructions `fix:"576,omitempty"`
	//ClearingFeeIndicator is a non-required field for NoAllocs.
	ClearingFeeIndicator *string `fix:"635"`
	//AllocSettlInstType is a non-required field for NoAllocs.
	AllocSettlInstType *int `fix:"780"`
	//SettlInstructionsData Component
	SettlInstructionsData settlinstructionsdata.Component
}

//NoMiscFees is a repeating group in NoAllocs
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

//NoClearingInstructions is a repeating group in NoAllocs
type NoClearingInstructions struct {
	//ClearingInstruction is a non-required field for NoClearingInstructions.
	ClearingInstruction *int `fix:"577"`
}

//Message is a AllocationReport FIX Message
type Message struct {
	FIXMsgType string `fix:"AS"`
	Header     fix44.Header
	//AllocReportID is a required field for AllocationReport.
	AllocReportID string `fix:"755"`
	//AllocID is a non-required field for AllocationReport.
	AllocID *string `fix:"70"`
	//AllocTransType is a required field for AllocationReport.
	AllocTransType string `fix:"71"`
	//AllocReportRefID is a non-required field for AllocationReport.
	AllocReportRefID *string `fix:"795"`
	//AllocCancReplaceReason is a non-required field for AllocationReport.
	AllocCancReplaceReason *int `fix:"796"`
	//SecondaryAllocID is a non-required field for AllocationReport.
	SecondaryAllocID *string `fix:"793"`
	//AllocReportType is a required field for AllocationReport.
	AllocReportType int `fix:"794"`
	//AllocStatus is a required field for AllocationReport.
	AllocStatus int `fix:"87"`
	//AllocRejCode is a non-required field for AllocationReport.
	AllocRejCode *int `fix:"88"`
	//RefAllocID is a non-required field for AllocationReport.
	RefAllocID *string `fix:"72"`
	//AllocIntermedReqType is a non-required field for AllocationReport.
	AllocIntermedReqType *int `fix:"808"`
	//AllocLinkID is a non-required field for AllocationReport.
	AllocLinkID *string `fix:"196"`
	//AllocLinkType is a non-required field for AllocationReport.
	AllocLinkType *int `fix:"197"`
	//BookingRefID is a non-required field for AllocationReport.
	BookingRefID *string `fix:"466"`
	//AllocNoOrdersType is a required field for AllocationReport.
	AllocNoOrdersType int `fix:"857"`
	//NoOrders is a non-required field for AllocationReport.
	NoOrders []NoOrders `fix:"73,omitempty"`
	//NoExecs is a non-required field for AllocationReport.
	NoExecs []NoExecs `fix:"124,omitempty"`
	//PreviouslyReported is a non-required field for AllocationReport.
	PreviouslyReported *bool `fix:"570"`
	//ReversalIndicator is a non-required field for AllocationReport.
	ReversalIndicator *bool `fix:"700"`
	//MatchType is a non-required field for AllocationReport.
	MatchType *string `fix:"574"`
	//Side is a required field for AllocationReport.
	Side string `fix:"54"`
	//Instrument Component
	Instrument instrument.Component
	//InstrumentExtension Component
	InstrumentExtension instrumentextension.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//NoUnderlyings is a non-required field for AllocationReport.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//NoLegs is a non-required field for AllocationReport.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//Quantity is a required field for AllocationReport.
	Quantity float64 `fix:"53"`
	//QtyType is a non-required field for AllocationReport.
	QtyType *int `fix:"854"`
	//LastMkt is a non-required field for AllocationReport.
	LastMkt *string `fix:"30"`
	//TradeOriginationDate is a non-required field for AllocationReport.
	TradeOriginationDate *string `fix:"229"`
	//TradingSessionID is a non-required field for AllocationReport.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for AllocationReport.
	TradingSessionSubID *string `fix:"625"`
	//PriceType is a non-required field for AllocationReport.
	PriceType *int `fix:"423"`
	//AvgPx is a required field for AllocationReport.
	AvgPx float64 `fix:"6"`
	//AvgParPx is a non-required field for AllocationReport.
	AvgParPx *float64 `fix:"860"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//Currency is a non-required field for AllocationReport.
	Currency *string `fix:"15"`
	//AvgPxPrecision is a non-required field for AllocationReport.
	AvgPxPrecision *int `fix:"74"`
	//Parties Component
	Parties parties.Component
	//TradeDate is a required field for AllocationReport.
	TradeDate string `fix:"75"`
	//TransactTime is a non-required field for AllocationReport.
	TransactTime *time.Time `fix:"60"`
	//SettlType is a non-required field for AllocationReport.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for AllocationReport.
	SettlDate *string `fix:"64"`
	//BookingType is a non-required field for AllocationReport.
	BookingType *int `fix:"775"`
	//GrossTradeAmt is a non-required field for AllocationReport.
	GrossTradeAmt *float64 `fix:"381"`
	//Concession is a non-required field for AllocationReport.
	Concession *float64 `fix:"238"`
	//TotalTakedown is a non-required field for AllocationReport.
	TotalTakedown *float64 `fix:"237"`
	//NetMoney is a non-required field for AllocationReport.
	NetMoney *float64 `fix:"118"`
	//PositionEffect is a non-required field for AllocationReport.
	PositionEffect *string `fix:"77"`
	//AutoAcceptIndicator is a non-required field for AllocationReport.
	AutoAcceptIndicator *bool `fix:"754"`
	//Text is a non-required field for AllocationReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for AllocationReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for AllocationReport.
	EncodedText *string `fix:"355"`
	//NumDaysInterest is a non-required field for AllocationReport.
	NumDaysInterest *int `fix:"157"`
	//AccruedInterestRate is a non-required field for AllocationReport.
	AccruedInterestRate *float64 `fix:"158"`
	//AccruedInterestAmt is a non-required field for AllocationReport.
	AccruedInterestAmt *float64 `fix:"159"`
	//TotalAccruedInterestAmt is a non-required field for AllocationReport.
	TotalAccruedInterestAmt *float64 `fix:"540"`
	//InterestAtMaturity is a non-required field for AllocationReport.
	InterestAtMaturity *float64 `fix:"738"`
	//EndAccruedInterestAmt is a non-required field for AllocationReport.
	EndAccruedInterestAmt *float64 `fix:"920"`
	//StartCash is a non-required field for AllocationReport.
	StartCash *float64 `fix:"921"`
	//EndCash is a non-required field for AllocationReport.
	EndCash *float64 `fix:"922"`
	//LegalConfirm is a non-required field for AllocationReport.
	LegalConfirm *bool `fix:"650"`
	//Stipulations Component
	Stipulations stipulations.Component
	//YieldData Component
	YieldData yielddata.Component
	//TotNoAllocs is a non-required field for AllocationReport.
	TotNoAllocs *int `fix:"892"`
	//LastFragment is a non-required field for AllocationReport.
	LastFragment *bool `fix:"893"`
	//NoAllocs is a non-required field for AllocationReport.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	Trailer  fix44.Trailer
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
	return enum.BeginStringFIX44, "AS", r
}
