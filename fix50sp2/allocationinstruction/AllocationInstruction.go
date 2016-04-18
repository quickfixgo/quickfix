//Package allocationinstruction msg type = J.
package allocationinstruction

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

//Message is a AllocationInstruction FIX Message
type Message struct {
	FIXMsgType string `fix:"J"`
	fixt11.Header
	//AllocID is a required field for AllocationInstruction.
	AllocID string `fix:"70"`
	//AllocTransType is a required field for AllocationInstruction.
	AllocTransType string `fix:"71"`
	//AllocType is a required field for AllocationInstruction.
	AllocType int `fix:"626"`
	//SecondaryAllocID is a non-required field for AllocationInstruction.
	SecondaryAllocID *string `fix:"793"`
	//RefAllocID is a non-required field for AllocationInstruction.
	RefAllocID *string `fix:"72"`
	//AllocCancReplaceReason is a non-required field for AllocationInstruction.
	AllocCancReplaceReason *int `fix:"796"`
	//AllocIntermedReqType is a non-required field for AllocationInstruction.
	AllocIntermedReqType *int `fix:"808"`
	//AllocLinkID is a non-required field for AllocationInstruction.
	AllocLinkID *string `fix:"196"`
	//AllocLinkType is a non-required field for AllocationInstruction.
	AllocLinkType *int `fix:"197"`
	//BookingRefID is a non-required field for AllocationInstruction.
	BookingRefID *string `fix:"466"`
	//AllocNoOrdersType is a non-required field for AllocationInstruction.
	AllocNoOrdersType *int `fix:"857"`
	//OrdAllocGrp is a non-required component for AllocationInstruction.
	OrdAllocGrp *ordallocgrp.OrdAllocGrp
	//ExecAllocGrp is a non-required component for AllocationInstruction.
	ExecAllocGrp *execallocgrp.ExecAllocGrp
	//PreviouslyReported is a non-required field for AllocationInstruction.
	PreviouslyReported *bool `fix:"570"`
	//ReversalIndicator is a non-required field for AllocationInstruction.
	ReversalIndicator *bool `fix:"700"`
	//MatchType is a non-required field for AllocationInstruction.
	MatchType *string `fix:"574"`
	//Side is a required field for AllocationInstruction.
	Side string `fix:"54"`
	//Instrument is a required component for AllocationInstruction.
	instrument.Instrument
	//InstrumentExtension is a non-required component for AllocationInstruction.
	InstrumentExtension *instrumentextension.InstrumentExtension
	//FinancingDetails is a non-required component for AllocationInstruction.
	FinancingDetails *financingdetails.FinancingDetails
	//UndInstrmtGrp is a non-required component for AllocationInstruction.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//InstrmtLegGrp is a non-required component for AllocationInstruction.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//Quantity is a required field for AllocationInstruction.
	Quantity float64 `fix:"53"`
	//QtyType is a non-required field for AllocationInstruction.
	QtyType *int `fix:"854"`
	//LastMkt is a non-required field for AllocationInstruction.
	LastMkt *string `fix:"30"`
	//TradeOriginationDate is a non-required field for AllocationInstruction.
	TradeOriginationDate *string `fix:"229"`
	//TradingSessionID is a non-required field for AllocationInstruction.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for AllocationInstruction.
	TradingSessionSubID *string `fix:"625"`
	//PriceType is a non-required field for AllocationInstruction.
	PriceType *int `fix:"423"`
	//AvgPx is a non-required field for AllocationInstruction.
	AvgPx *float64 `fix:"6"`
	//AvgParPx is a non-required field for AllocationInstruction.
	AvgParPx *float64 `fix:"860"`
	//SpreadOrBenchmarkCurveData is a non-required component for AllocationInstruction.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//Currency is a non-required field for AllocationInstruction.
	Currency *string `fix:"15"`
	//AvgPxPrecision is a non-required field for AllocationInstruction.
	AvgPxPrecision *int `fix:"74"`
	//Parties is a non-required component for AllocationInstruction.
	Parties *parties.Parties
	//TradeDate is a required field for AllocationInstruction.
	TradeDate string `fix:"75"`
	//TransactTime is a non-required field for AllocationInstruction.
	TransactTime *time.Time `fix:"60"`
	//SettlType is a non-required field for AllocationInstruction.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for AllocationInstruction.
	SettlDate *string `fix:"64"`
	//BookingType is a non-required field for AllocationInstruction.
	BookingType *int `fix:"775"`
	//GrossTradeAmt is a non-required field for AllocationInstruction.
	GrossTradeAmt *float64 `fix:"381"`
	//Concession is a non-required field for AllocationInstruction.
	Concession *float64 `fix:"238"`
	//TotalTakedown is a non-required field for AllocationInstruction.
	TotalTakedown *float64 `fix:"237"`
	//NetMoney is a non-required field for AllocationInstruction.
	NetMoney *float64 `fix:"118"`
	//PositionEffect is a non-required field for AllocationInstruction.
	PositionEffect *string `fix:"77"`
	//AutoAcceptIndicator is a non-required field for AllocationInstruction.
	AutoAcceptIndicator *bool `fix:"754"`
	//Text is a non-required field for AllocationInstruction.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for AllocationInstruction.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for AllocationInstruction.
	EncodedText *string `fix:"355"`
	//NumDaysInterest is a non-required field for AllocationInstruction.
	NumDaysInterest *int `fix:"157"`
	//AccruedInterestRate is a non-required field for AllocationInstruction.
	AccruedInterestRate *float64 `fix:"158"`
	//AccruedInterestAmt is a non-required field for AllocationInstruction.
	AccruedInterestAmt *float64 `fix:"159"`
	//TotalAccruedInterestAmt is a non-required field for AllocationInstruction.
	TotalAccruedInterestAmt *float64 `fix:"540"`
	//InterestAtMaturity is a non-required field for AllocationInstruction.
	InterestAtMaturity *float64 `fix:"738"`
	//EndAccruedInterestAmt is a non-required field for AllocationInstruction.
	EndAccruedInterestAmt *float64 `fix:"920"`
	//StartCash is a non-required field for AllocationInstruction.
	StartCash *float64 `fix:"921"`
	//EndCash is a non-required field for AllocationInstruction.
	EndCash *float64 `fix:"922"`
	//LegalConfirm is a non-required field for AllocationInstruction.
	LegalConfirm *bool `fix:"650"`
	//Stipulations is a non-required component for AllocationInstruction.
	Stipulations *stipulations.Stipulations
	//YieldData is a non-required component for AllocationInstruction.
	YieldData *yielddata.YieldData
	//TotNoAllocs is a non-required field for AllocationInstruction.
	TotNoAllocs *int `fix:"892"`
	//LastFragment is a non-required field for AllocationInstruction.
	LastFragment *bool `fix:"893"`
	//AllocGrp is a non-required component for AllocationInstruction.
	AllocGrp *allocgrp.AllocGrp
	//PositionAmountData is a non-required component for AllocationInstruction.
	PositionAmountData *positionamountdata.PositionAmountData
	//AvgPxIndicator is a non-required field for AllocationInstruction.
	AvgPxIndicator *int `fix:"819"`
	//ClearingBusinessDate is a non-required field for AllocationInstruction.
	ClearingBusinessDate *string `fix:"715"`
	//TrdType is a non-required field for AllocationInstruction.
	TrdType *int `fix:"828"`
	//TrdSubType is a non-required field for AllocationInstruction.
	TrdSubType *int `fix:"829"`
	//CustOrderCapacity is a non-required field for AllocationInstruction.
	CustOrderCapacity *int `fix:"582"`
	//TradeInputSource is a non-required field for AllocationInstruction.
	TradeInputSource *string `fix:"578"`
	//MultiLegReportingType is a non-required field for AllocationInstruction.
	MultiLegReportingType *string `fix:"442"`
	//MessageEventSource is a non-required field for AllocationInstruction.
	MessageEventSource *string `fix:"1011"`
	//RndPx is a non-required field for AllocationInstruction.
	RndPx *float64 `fix:"991"`
	//RateSource is a non-required component for AllocationInstruction.
	RateSource *ratesource.RateSource
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized AllocationInstruction instance
func New(allocid string, alloctranstype string, alloctype int, side string, instrument instrument.Instrument, quantity float64, tradedate string) *Message {
	var m Message
	m.SetAllocID(allocid)
	m.SetAllocTransType(alloctranstype)
	m.SetAllocType(alloctype)
	m.SetSide(side)
	m.SetInstrument(instrument)
	m.SetQuantity(quantity)
	m.SetTradeDate(tradedate)
	return &m
}

func (m *Message) SetAllocID(v string)                         { m.AllocID = v }
func (m *Message) SetAllocTransType(v string)                  { m.AllocTransType = v }
func (m *Message) SetAllocType(v int)                          { m.AllocType = v }
func (m *Message) SetSecondaryAllocID(v string)                { m.SecondaryAllocID = &v }
func (m *Message) SetRefAllocID(v string)                      { m.RefAllocID = &v }
func (m *Message) SetAllocCancReplaceReason(v int)             { m.AllocCancReplaceReason = &v }
func (m *Message) SetAllocIntermedReqType(v int)               { m.AllocIntermedReqType = &v }
func (m *Message) SetAllocLinkID(v string)                     { m.AllocLinkID = &v }
func (m *Message) SetAllocLinkType(v int)                      { m.AllocLinkType = &v }
func (m *Message) SetBookingRefID(v string)                    { m.BookingRefID = &v }
func (m *Message) SetAllocNoOrdersType(v int)                  { m.AllocNoOrdersType = &v }
func (m *Message) SetOrdAllocGrp(v ordallocgrp.OrdAllocGrp)    { m.OrdAllocGrp = &v }
func (m *Message) SetExecAllocGrp(v execallocgrp.ExecAllocGrp) { m.ExecAllocGrp = &v }
func (m *Message) SetPreviouslyReported(v bool)                { m.PreviouslyReported = &v }
func (m *Message) SetReversalIndicator(v bool)                 { m.ReversalIndicator = &v }
func (m *Message) SetMatchType(v string)                       { m.MatchType = &v }
func (m *Message) SetSide(v string)                            { m.Side = v }
func (m *Message) SetInstrument(v instrument.Instrument)       { m.Instrument = v }
func (m *Message) SetInstrumentExtension(v instrumentextension.InstrumentExtension) {
	m.InstrumentExtension = &v
}
func (m *Message) SetFinancingDetails(v financingdetails.FinancingDetails) { m.FinancingDetails = &v }
func (m *Message) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp)          { m.UndInstrmtGrp = &v }
func (m *Message) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp)          { m.InstrmtLegGrp = &v }
func (m *Message) SetQuantity(v float64)                                   { m.Quantity = v }
func (m *Message) SetQtyType(v int)                                        { m.QtyType = &v }
func (m *Message) SetLastMkt(v string)                                     { m.LastMkt = &v }
func (m *Message) SetTradeOriginationDate(v string)                        { m.TradeOriginationDate = &v }
func (m *Message) SetTradingSessionID(v string)                            { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)                         { m.TradingSessionSubID = &v }
func (m *Message) SetPriceType(v int)                                      { m.PriceType = &v }
func (m *Message) SetAvgPx(v float64)                                      { m.AvgPx = &v }
func (m *Message) SetAvgParPx(v float64)                                   { m.AvgParPx = &v }
func (m *Message) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *Message) SetCurrency(v string)                        { m.Currency = &v }
func (m *Message) SetAvgPxPrecision(v int)                     { m.AvgPxPrecision = &v }
func (m *Message) SetParties(v parties.Parties)                { m.Parties = &v }
func (m *Message) SetTradeDate(v string)                       { m.TradeDate = v }
func (m *Message) SetTransactTime(v time.Time)                 { m.TransactTime = &v }
func (m *Message) SetSettlType(v string)                       { m.SettlType = &v }
func (m *Message) SetSettlDate(v string)                       { m.SettlDate = &v }
func (m *Message) SetBookingType(v int)                        { m.BookingType = &v }
func (m *Message) SetGrossTradeAmt(v float64)                  { m.GrossTradeAmt = &v }
func (m *Message) SetConcession(v float64)                     { m.Concession = &v }
func (m *Message) SetTotalTakedown(v float64)                  { m.TotalTakedown = &v }
func (m *Message) SetNetMoney(v float64)                       { m.NetMoney = &v }
func (m *Message) SetPositionEffect(v string)                  { m.PositionEffect = &v }
func (m *Message) SetAutoAcceptIndicator(v bool)               { m.AutoAcceptIndicator = &v }
func (m *Message) SetText(v string)                            { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                     { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                     { m.EncodedText = &v }
func (m *Message) SetNumDaysInterest(v int)                    { m.NumDaysInterest = &v }
func (m *Message) SetAccruedInterestRate(v float64)            { m.AccruedInterestRate = &v }
func (m *Message) SetAccruedInterestAmt(v float64)             { m.AccruedInterestAmt = &v }
func (m *Message) SetTotalAccruedInterestAmt(v float64)        { m.TotalAccruedInterestAmt = &v }
func (m *Message) SetInterestAtMaturity(v float64)             { m.InterestAtMaturity = &v }
func (m *Message) SetEndAccruedInterestAmt(v float64)          { m.EndAccruedInterestAmt = &v }
func (m *Message) SetStartCash(v float64)                      { m.StartCash = &v }
func (m *Message) SetEndCash(v float64)                        { m.EndCash = &v }
func (m *Message) SetLegalConfirm(v bool)                      { m.LegalConfirm = &v }
func (m *Message) SetStipulations(v stipulations.Stipulations) { m.Stipulations = &v }
func (m *Message) SetYieldData(v yielddata.YieldData)          { m.YieldData = &v }
func (m *Message) SetTotNoAllocs(v int)                        { m.TotNoAllocs = &v }
func (m *Message) SetLastFragment(v bool)                      { m.LastFragment = &v }
func (m *Message) SetAllocGrp(v allocgrp.AllocGrp)             { m.AllocGrp = &v }
func (m *Message) SetPositionAmountData(v positionamountdata.PositionAmountData) {
	m.PositionAmountData = &v
}
func (m *Message) SetAvgPxIndicator(v int)               { m.AvgPxIndicator = &v }
func (m *Message) SetClearingBusinessDate(v string)      { m.ClearingBusinessDate = &v }
func (m *Message) SetTrdType(v int)                      { m.TrdType = &v }
func (m *Message) SetTrdSubType(v int)                   { m.TrdSubType = &v }
func (m *Message) SetCustOrderCapacity(v int)            { m.CustOrderCapacity = &v }
func (m *Message) SetTradeInputSource(v string)          { m.TradeInputSource = &v }
func (m *Message) SetMultiLegReportingType(v string)     { m.MultiLegReportingType = &v }
func (m *Message) SetMessageEventSource(v string)        { m.MessageEventSource = &v }
func (m *Message) SetRndPx(v float64)                    { m.RndPx = &v }
func (m *Message) SetRateSource(v ratesource.RateSource) { m.RateSource = &v }

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
	return enum.ApplVerID_FIX50SP2, "J", r
}
