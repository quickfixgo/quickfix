//Package allocationreport msg type = AS.
package allocationreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/allocgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/execallocgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp2/ordallocgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/positionamountdata"
	"github.com/quickfixgo/quickfix/fix50sp2/ratesource"
	"github.com/quickfixgo/quickfix/fix50sp2/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp2/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a AllocationReport FIX Message
type Message struct {
	FIXMsgType string `fix:"AS"`
	Header     fixt11.Header
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
	//AllocNoOrdersType is a non-required field for AllocationReport.
	AllocNoOrdersType *int `fix:"857"`
	//OrdAllocGrp Component
	OrdAllocGrp ordallocgrp.Component
	//ExecAllocGrp Component
	ExecAllocGrp execallocgrp.Component
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
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
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
	//AllocGrp Component
	AllocGrp allocgrp.Component
	//ClearingBusinessDate is a non-required field for AllocationReport.
	ClearingBusinessDate *string `fix:"715"`
	//TrdType is a non-required field for AllocationReport.
	TrdType *int `fix:"828"`
	//TrdSubType is a non-required field for AllocationReport.
	TrdSubType *int `fix:"829"`
	//MultiLegReportingType is a non-required field for AllocationReport.
	MultiLegReportingType *string `fix:"442"`
	//CustOrderCapacity is a non-required field for AllocationReport.
	CustOrderCapacity *int `fix:"582"`
	//TradeInputSource is a non-required field for AllocationReport.
	TradeInputSource *string `fix:"578"`
	//RndPx is a non-required field for AllocationReport.
	RndPx *float64 `fix:"991"`
	//MessageEventSource is a non-required field for AllocationReport.
	MessageEventSource *string `fix:"1011"`
	//TradeInputDevice is a non-required field for AllocationReport.
	TradeInputDevice *string `fix:"579"`
	//AvgPxIndicator is a non-required field for AllocationReport.
	AvgPxIndicator *int `fix:"819"`
	//PositionAmountData Component
	PositionAmountData positionamountdata.Component
	//RateSource Component
	RateSource ratesource.Component
	Trailer    fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP2, "AS", r
}
