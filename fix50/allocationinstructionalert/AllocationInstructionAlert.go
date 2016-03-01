//Package allocationinstructionalert msg type = BM.
package allocationinstructionalert

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/allocgrp"
	"github.com/quickfixgo/quickfix/fix50/execallocgrp"
	"github.com/quickfixgo/quickfix/fix50/financingdetails"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50/ordallocgrp"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/positionamountdata"
	"github.com/quickfixgo/quickfix/fix50/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/stipulations"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a AllocationInstructionAlert FIX Message
type Message struct {
	FIXMsgType string `fix:"BM"`
	Header     fixt11.Header
	//AllocID is a required field for AllocationInstructionAlert.
	AllocID string `fix:"70"`
	//AllocTransType is a required field for AllocationInstructionAlert.
	AllocTransType string `fix:"71"`
	//AllocType is a required field for AllocationInstructionAlert.
	AllocType int `fix:"626"`
	//SecondaryAllocID is a non-required field for AllocationInstructionAlert.
	SecondaryAllocID *string `fix:"793"`
	//RefAllocID is a non-required field for AllocationInstructionAlert.
	RefAllocID *string `fix:"72"`
	//AllocCancReplaceReason is a non-required field for AllocationInstructionAlert.
	AllocCancReplaceReason *int `fix:"796"`
	//AllocIntermedReqType is a non-required field for AllocationInstructionAlert.
	AllocIntermedReqType *int `fix:"808"`
	//AllocLinkID is a non-required field for AllocationInstructionAlert.
	AllocLinkID *string `fix:"196"`
	//AllocLinkType is a non-required field for AllocationInstructionAlert.
	AllocLinkType *int `fix:"197"`
	//BookingRefID is a non-required field for AllocationInstructionAlert.
	BookingRefID *string `fix:"466"`
	//AllocNoOrdersType is a non-required field for AllocationInstructionAlert.
	AllocNoOrdersType *int `fix:"857"`
	//OrdAllocGrp Component
	OrdAllocGrp ordallocgrp.Component
	//ExecAllocGrp Component
	ExecAllocGrp execallocgrp.Component
	//PreviouslyReported is a non-required field for AllocationInstructionAlert.
	PreviouslyReported *bool `fix:"570"`
	//ReversalIndicator is a non-required field for AllocationInstructionAlert.
	ReversalIndicator *bool `fix:"700"`
	//MatchType is a non-required field for AllocationInstructionAlert.
	MatchType *string `fix:"574"`
	//Side is a required field for AllocationInstructionAlert.
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
	//Quantity is a required field for AllocationInstructionAlert.
	Quantity float64 `fix:"53"`
	//QtyType is a non-required field for AllocationInstructionAlert.
	QtyType *int `fix:"854"`
	//LastMkt is a non-required field for AllocationInstructionAlert.
	LastMkt *string `fix:"30"`
	//TradeOriginationDate is a non-required field for AllocationInstructionAlert.
	TradeOriginationDate *string `fix:"229"`
	//TradingSessionID is a non-required field for AllocationInstructionAlert.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for AllocationInstructionAlert.
	TradingSessionSubID *string `fix:"625"`
	//PriceType is a non-required field for AllocationInstructionAlert.
	PriceType *int `fix:"423"`
	//AvgPx is a non-required field for AllocationInstructionAlert.
	AvgPx *float64 `fix:"6"`
	//AvgParPx is a non-required field for AllocationInstructionAlert.
	AvgParPx *float64 `fix:"860"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//Currency is a non-required field for AllocationInstructionAlert.
	Currency *string `fix:"15"`
	//AvgPxPrecision is a non-required field for AllocationInstructionAlert.
	AvgPxPrecision *int `fix:"74"`
	//Parties Component
	Parties parties.Component
	//TradeDate is a required field for AllocationInstructionAlert.
	TradeDate string `fix:"75"`
	//TransactTime is a non-required field for AllocationInstructionAlert.
	TransactTime *time.Time `fix:"60"`
	//SettlType is a non-required field for AllocationInstructionAlert.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for AllocationInstructionAlert.
	SettlDate *string `fix:"64"`
	//BookingType is a non-required field for AllocationInstructionAlert.
	BookingType *int `fix:"775"`
	//GrossTradeAmt is a non-required field for AllocationInstructionAlert.
	GrossTradeAmt *float64 `fix:"381"`
	//Concession is a non-required field for AllocationInstructionAlert.
	Concession *float64 `fix:"238"`
	//TotalTakedown is a non-required field for AllocationInstructionAlert.
	TotalTakedown *float64 `fix:"237"`
	//NetMoney is a non-required field for AllocationInstructionAlert.
	NetMoney *float64 `fix:"118"`
	//PositionEffect is a non-required field for AllocationInstructionAlert.
	PositionEffect *string `fix:"77"`
	//AutoAcceptIndicator is a non-required field for AllocationInstructionAlert.
	AutoAcceptIndicator *bool `fix:"754"`
	//Text is a non-required field for AllocationInstructionAlert.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for AllocationInstructionAlert.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for AllocationInstructionAlert.
	EncodedText *string `fix:"355"`
	//NumDaysInterest is a non-required field for AllocationInstructionAlert.
	NumDaysInterest *int `fix:"157"`
	//AccruedInterestRate is a non-required field for AllocationInstructionAlert.
	AccruedInterestRate *float64 `fix:"158"`
	//AccruedInterestAmt is a non-required field for AllocationInstructionAlert.
	AccruedInterestAmt *float64 `fix:"159"`
	//TotalAccruedInterestAmt is a non-required field for AllocationInstructionAlert.
	TotalAccruedInterestAmt *float64 `fix:"540"`
	//InterestAtMaturity is a non-required field for AllocationInstructionAlert.
	InterestAtMaturity *float64 `fix:"738"`
	//EndAccruedInterestAmt is a non-required field for AllocationInstructionAlert.
	EndAccruedInterestAmt *float64 `fix:"920"`
	//StartCash is a non-required field for AllocationInstructionAlert.
	StartCash *float64 `fix:"921"`
	//EndCash is a non-required field for AllocationInstructionAlert.
	EndCash *float64 `fix:"922"`
	//LegalConfirm is a non-required field for AllocationInstructionAlert.
	LegalConfirm *bool `fix:"650"`
	//Stipulations Component
	Stipulations stipulations.Component
	//YieldData Component
	YieldData yielddata.Component
	//PositionAmountData Component
	PositionAmountData positionamountdata.Component
	//TotNoAllocs is a non-required field for AllocationInstructionAlert.
	TotNoAllocs *int `fix:"892"`
	//LastFragment is a non-required field for AllocationInstructionAlert.
	LastFragment *bool `fix:"893"`
	//AllocGrp Component
	AllocGrp allocgrp.Component
	//AvgPxIndicator is a non-required field for AllocationInstructionAlert.
	AvgPxIndicator *int `fix:"819"`
	//ClearingBusinessDate is a non-required field for AllocationInstructionAlert.
	ClearingBusinessDate *string `fix:"715"`
	//TrdType is a non-required field for AllocationInstructionAlert.
	TrdType *int `fix:"828"`
	//TrdSubType is a non-required field for AllocationInstructionAlert.
	TrdSubType *int `fix:"829"`
	//CustOrderCapacity is a non-required field for AllocationInstructionAlert.
	CustOrderCapacity *int `fix:"582"`
	//TradeInputSource is a non-required field for AllocationInstructionAlert.
	TradeInputSource *string `fix:"578"`
	//MultiLegReportingType is a non-required field for AllocationInstructionAlert.
	MultiLegReportingType *string `fix:"442"`
	//MessageEventSource is a non-required field for AllocationInstructionAlert.
	MessageEventSource *string `fix:"1011"`
	//RndPx is a non-required field for AllocationInstructionAlert.
	RndPx   *float64 `fix:"991"`
	Trailer fixt11.Trailer
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
	return enum.BeginStringFIX50, "BM", r
}
