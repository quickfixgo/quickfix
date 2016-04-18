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
	//UnderlyingInstrument is a non-required component for NoUnderlyings.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
}

//NewNoUnderlyings returns an initialized NoUnderlyings instance
func NewNoUnderlyings() *NoUnderlyings {
	var m NoUnderlyings
	return &m
}

func (m *NoUnderlyings) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}

//NoLegs is a repeating group in TradeCaptureReport
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
	//LegQty is a non-required field for NoLegs.
	LegQty *float64 `fix:"687"`
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegStipulations is a non-required component for NoLegs.
	LegStipulations *legstipulations.LegStipulations
	//LegPositionEffect is a non-required field for NoLegs.
	LegPositionEffect *string `fix:"564"`
	//LegCoveredOrUncovered is a non-required field for NoLegs.
	LegCoveredOrUncovered *int `fix:"565"`
	//NestedParties is a non-required component for NoLegs.
	NestedParties *nestedparties.NestedParties
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

//NewNoLegs returns an initialized NoLegs instance
func NewNoLegs() *NoLegs {
	var m NoLegs
	return &m
}

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg)       { m.InstrumentLeg = &v }
func (m *NoLegs) SetLegQty(v float64)                                  { m.LegQty = &v }
func (m *NoLegs) SetLegSwapType(v int)                                 { m.LegSwapType = &v }
func (m *NoLegs) SetLegStipulations(v legstipulations.LegStipulations) { m.LegStipulations = &v }
func (m *NoLegs) SetLegPositionEffect(v string)                        { m.LegPositionEffect = &v }
func (m *NoLegs) SetLegCoveredOrUncovered(v int)                       { m.LegCoveredOrUncovered = &v }
func (m *NoLegs) SetNestedParties(v nestedparties.NestedParties)       { m.NestedParties = &v }
func (m *NoLegs) SetLegRefID(v string)                                 { m.LegRefID = &v }
func (m *NoLegs) SetLegPrice(v float64)                                { m.LegPrice = &v }
func (m *NoLegs) SetLegSettlType(v string)                             { m.LegSettlType = &v }
func (m *NoLegs) SetLegSettlDate(v string)                             { m.LegSettlDate = &v }
func (m *NoLegs) SetLegLastPx(v float64)                               { m.LegLastPx = &v }

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
	//Parties is a non-required component for NoSides.
	Parties *parties.Parties
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
	//CommissionData is a non-required component for NoSides.
	CommissionData *commissiondata.CommissionData
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
	//Stipulations is a non-required component for NoSides.
	Stipulations *stipulations.Stipulations
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

//NewNoSides returns an initialized NoSides instance
func NewNoSides(side string, orderid string) *NoSides {
	var m NoSides
	m.SetSide(side)
	m.SetOrderID(orderid)
	return &m
}

func (m *NoSides) SetSide(v string)                                     { m.Side = v }
func (m *NoSides) SetOrderID(v string)                                  { m.OrderID = v }
func (m *NoSides) SetSecondaryOrderID(v string)                         { m.SecondaryOrderID = &v }
func (m *NoSides) SetClOrdID(v string)                                  { m.ClOrdID = &v }
func (m *NoSides) SetSecondaryClOrdID(v string)                         { m.SecondaryClOrdID = &v }
func (m *NoSides) SetListID(v string)                                   { m.ListID = &v }
func (m *NoSides) SetParties(v parties.Parties)                         { m.Parties = &v }
func (m *NoSides) SetAccount(v string)                                  { m.Account = &v }
func (m *NoSides) SetAcctIDSource(v int)                                { m.AcctIDSource = &v }
func (m *NoSides) SetAccountType(v int)                                 { m.AccountType = &v }
func (m *NoSides) SetProcessCode(v string)                              { m.ProcessCode = &v }
func (m *NoSides) SetOddLot(v bool)                                     { m.OddLot = &v }
func (m *NoSides) SetNoClearingInstructions(v []NoClearingInstructions) { m.NoClearingInstructions = v }
func (m *NoSides) SetClearingFeeIndicator(v string)                     { m.ClearingFeeIndicator = &v }
func (m *NoSides) SetTradeInputSource(v string)                         { m.TradeInputSource = &v }
func (m *NoSides) SetTradeInputDevice(v string)                         { m.TradeInputDevice = &v }
func (m *NoSides) SetOrderInputDevice(v string)                         { m.OrderInputDevice = &v }
func (m *NoSides) SetCurrency(v string)                                 { m.Currency = &v }
func (m *NoSides) SetComplianceID(v string)                             { m.ComplianceID = &v }
func (m *NoSides) SetSolicitedFlag(v bool)                              { m.SolicitedFlag = &v }
func (m *NoSides) SetOrderCapacity(v string)                            { m.OrderCapacity = &v }
func (m *NoSides) SetOrderRestrictions(v string)                        { m.OrderRestrictions = &v }
func (m *NoSides) SetCustOrderCapacity(v int)                           { m.CustOrderCapacity = &v }
func (m *NoSides) SetOrdType(v string)                                  { m.OrdType = &v }
func (m *NoSides) SetExecInst(v string)                                 { m.ExecInst = &v }
func (m *NoSides) SetTransBkdTime(v time.Time)                          { m.TransBkdTime = &v }
func (m *NoSides) SetTradingSessionID(v string)                         { m.TradingSessionID = &v }
func (m *NoSides) SetTradingSessionSubID(v string)                      { m.TradingSessionSubID = &v }
func (m *NoSides) SetTimeBracket(v string)                              { m.TimeBracket = &v }
func (m *NoSides) SetCommissionData(v commissiondata.CommissionData)    { m.CommissionData = &v }
func (m *NoSides) SetGrossTradeAmt(v float64)                           { m.GrossTradeAmt = &v }
func (m *NoSides) SetNumDaysInterest(v int)                             { m.NumDaysInterest = &v }
func (m *NoSides) SetExDate(v string)                                   { m.ExDate = &v }
func (m *NoSides) SetAccruedInterestRate(v float64)                     { m.AccruedInterestRate = &v }
func (m *NoSides) SetAccruedInterestAmt(v float64)                      { m.AccruedInterestAmt = &v }
func (m *NoSides) SetInterestAtMaturity(v float64)                      { m.InterestAtMaturity = &v }
func (m *NoSides) SetEndAccruedInterestAmt(v float64)                   { m.EndAccruedInterestAmt = &v }
func (m *NoSides) SetStartCash(v float64)                               { m.StartCash = &v }
func (m *NoSides) SetEndCash(v float64)                                 { m.EndCash = &v }
func (m *NoSides) SetConcession(v float64)                              { m.Concession = &v }
func (m *NoSides) SetTotalTakedown(v float64)                           { m.TotalTakedown = &v }
func (m *NoSides) SetNetMoney(v float64)                                { m.NetMoney = &v }
func (m *NoSides) SetSettlCurrAmt(v float64)                            { m.SettlCurrAmt = &v }
func (m *NoSides) SetSettlCurrency(v string)                            { m.SettlCurrency = &v }
func (m *NoSides) SetSettlCurrFxRate(v float64)                         { m.SettlCurrFxRate = &v }
func (m *NoSides) SetSettlCurrFxRateCalc(v string)                      { m.SettlCurrFxRateCalc = &v }
func (m *NoSides) SetPositionEffect(v string)                           { m.PositionEffect = &v }
func (m *NoSides) SetText(v string)                                     { m.Text = &v }
func (m *NoSides) SetEncodedTextLen(v int)                              { m.EncodedTextLen = &v }
func (m *NoSides) SetEncodedText(v string)                              { m.EncodedText = &v }
func (m *NoSides) SetSideMultiLegReportingType(v int)                   { m.SideMultiLegReportingType = &v }
func (m *NoSides) SetNoContAmts(v []NoContAmts)                         { m.NoContAmts = v }
func (m *NoSides) SetStipulations(v stipulations.Stipulations)          { m.Stipulations = &v }
func (m *NoSides) SetNoMiscFees(v []NoMiscFees)                         { m.NoMiscFees = v }
func (m *NoSides) SetExchangeRule(v string)                             { m.ExchangeRule = &v }
func (m *NoSides) SetTradeAllocIndicator(v int)                         { m.TradeAllocIndicator = &v }
func (m *NoSides) SetPreallocMethod(v string)                           { m.PreallocMethod = &v }
func (m *NoSides) SetAllocID(v string)                                  { m.AllocID = &v }
func (m *NoSides) SetNoAllocs(v []NoAllocs)                             { m.NoAllocs = v }
func (m *NoSides) SetCopyMsgIndicator(v bool)                           { m.CopyMsgIndicator = &v }
func (m *NoSides) SetPublishTrdIndicator(v bool)                        { m.PublishTrdIndicator = &v }
func (m *NoSides) SetShortSaleReason(v int)                             { m.ShortSaleReason = &v }

//NoClearingInstructions is a repeating group in NoSides
type NoClearingInstructions struct {
	//ClearingInstruction is a non-required field for NoClearingInstructions.
	ClearingInstruction *int `fix:"577"`
}

//NewNoClearingInstructions returns an initialized NoClearingInstructions instance
func NewNoClearingInstructions() *NoClearingInstructions {
	var m NoClearingInstructions
	return &m
}

func (m *NoClearingInstructions) SetClearingInstruction(v int) { m.ClearingInstruction = &v }

//NoContAmts is a repeating group in NoSides
type NoContAmts struct {
	//ContAmtType is a non-required field for NoContAmts.
	ContAmtType *int `fix:"519"`
	//ContAmtValue is a non-required field for NoContAmts.
	ContAmtValue *float64 `fix:"520"`
	//ContAmtCurr is a non-required field for NoContAmts.
	ContAmtCurr *string `fix:"521"`
}

//NewNoContAmts returns an initialized NoContAmts instance
func NewNoContAmts() *NoContAmts {
	var m NoContAmts
	return &m
}

func (m *NoContAmts) SetContAmtType(v int)      { m.ContAmtType = &v }
func (m *NoContAmts) SetContAmtValue(v float64) { m.ContAmtValue = &v }
func (m *NoContAmts) SetContAmtCurr(v string)   { m.ContAmtCurr = &v }

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

//NewNoMiscFees returns an initialized NoMiscFees instance
func NewNoMiscFees() *NoMiscFees {
	var m NoMiscFees
	return &m
}

func (m *NoMiscFees) SetMiscFeeAmt(v float64) { m.MiscFeeAmt = &v }
func (m *NoMiscFees) SetMiscFeeCurr(v string) { m.MiscFeeCurr = &v }
func (m *NoMiscFees) SetMiscFeeType(v string) { m.MiscFeeType = &v }
func (m *NoMiscFees) SetMiscFeeBasis(v int)   { m.MiscFeeBasis = &v }

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
	//NestedParties2 is a non-required component for NoAllocs.
	NestedParties2 *nestedparties2.NestedParties2
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
}

//NewNoAllocs returns an initialized NoAllocs instance
func NewNoAllocs() *NoAllocs {
	var m NoAllocs
	return &m
}

func (m *NoAllocs) SetAllocAccount(v string)                          { m.AllocAccount = &v }
func (m *NoAllocs) SetAllocAcctIDSource(v int)                        { m.AllocAcctIDSource = &v }
func (m *NoAllocs) SetAllocSettlCurrency(v string)                    { m.AllocSettlCurrency = &v }
func (m *NoAllocs) SetIndividualAllocID(v string)                     { m.IndividualAllocID = &v }
func (m *NoAllocs) SetNestedParties2(v nestedparties2.NestedParties2) { m.NestedParties2 = &v }
func (m *NoAllocs) SetAllocQty(v float64)                             { m.AllocQty = &v }

//Message is a TradeCaptureReport FIX Message
type Message struct {
	FIXMsgType string `fix:"AE"`
	fix44.Header
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
	//Instrument is a required component for TradeCaptureReport.
	instrument.Instrument
	//FinancingDetails is a non-required component for TradeCaptureReport.
	FinancingDetails *financingdetails.FinancingDetails
	//OrderQtyData is a non-required component for TradeCaptureReport.
	OrderQtyData *orderqtydata.OrderQtyData
	//QtyType is a non-required field for TradeCaptureReport.
	QtyType *int `fix:"854"`
	//YieldData is a non-required component for TradeCaptureReport.
	YieldData *yielddata.YieldData
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
	//SpreadOrBenchmarkCurveData is a non-required component for TradeCaptureReport.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//AvgPxIndicator is a non-required field for TradeCaptureReport.
	AvgPxIndicator *int `fix:"819"`
	//PositionAmountData is a non-required component for TradeCaptureReport.
	PositionAmountData *positionamountdata.PositionAmountData
	//MultiLegReportingType is a non-required field for TradeCaptureReport.
	MultiLegReportingType *string `fix:"442"`
	//TradeLegRefID is a non-required field for TradeCaptureReport.
	TradeLegRefID *string `fix:"824"`
	//NoLegs is a non-required field for TradeCaptureReport.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//TransactTime is a required field for TradeCaptureReport.
	TransactTime time.Time `fix:"60"`
	//TrdRegTimestamps is a non-required component for TradeCaptureReport.
	TrdRegTimestamps *trdregtimestamps.TrdRegTimestamps
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
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized TradeCaptureReport instance
func New(tradereportid string, previouslyreported bool, instrument instrument.Instrument, lastqty float64, lastpx float64, tradedate string, transacttime time.Time, nosides []NoSides) *Message {
	var m Message
	m.SetTradeReportID(tradereportid)
	m.SetPreviouslyReported(previouslyreported)
	m.SetInstrument(instrument)
	m.SetLastQty(lastqty)
	m.SetLastPx(lastpx)
	m.SetTradeDate(tradedate)
	m.SetTransactTime(transacttime)
	m.SetNoSides(nosides)
	return &m
}

func (m *Message) SetTradeReportID(v string)                               { m.TradeReportID = v }
func (m *Message) SetTradeReportTransType(v int)                           { m.TradeReportTransType = &v }
func (m *Message) SetTradeReportType(v int)                                { m.TradeReportType = &v }
func (m *Message) SetTradeRequestID(v string)                              { m.TradeRequestID = &v }
func (m *Message) SetTrdType(v int)                                        { m.TrdType = &v }
func (m *Message) SetTrdSubType(v int)                                     { m.TrdSubType = &v }
func (m *Message) SetSecondaryTrdType(v int)                               { m.SecondaryTrdType = &v }
func (m *Message) SetTransferReason(v string)                              { m.TransferReason = &v }
func (m *Message) SetExecType(v string)                                    { m.ExecType = &v }
func (m *Message) SetTotNumTradeReports(v int)                             { m.TotNumTradeReports = &v }
func (m *Message) SetLastRptRequested(v bool)                              { m.LastRptRequested = &v }
func (m *Message) SetUnsolicitedIndicator(v bool)                          { m.UnsolicitedIndicator = &v }
func (m *Message) SetSubscriptionRequestType(v string)                     { m.SubscriptionRequestType = &v }
func (m *Message) SetTradeReportRefID(v string)                            { m.TradeReportRefID = &v }
func (m *Message) SetSecondaryTradeReportRefID(v string)                   { m.SecondaryTradeReportRefID = &v }
func (m *Message) SetSecondaryTradeReportID(v string)                      { m.SecondaryTradeReportID = &v }
func (m *Message) SetTradeLinkID(v string)                                 { m.TradeLinkID = &v }
func (m *Message) SetTrdMatchID(v string)                                  { m.TrdMatchID = &v }
func (m *Message) SetExecID(v string)                                      { m.ExecID = &v }
func (m *Message) SetOrdStatus(v string)                                   { m.OrdStatus = &v }
func (m *Message) SetSecondaryExecID(v string)                             { m.SecondaryExecID = &v }
func (m *Message) SetExecRestatementReason(v int)                          { m.ExecRestatementReason = &v }
func (m *Message) SetPreviouslyReported(v bool)                            { m.PreviouslyReported = v }
func (m *Message) SetPriceType(v int)                                      { m.PriceType = &v }
func (m *Message) SetInstrument(v instrument.Instrument)                   { m.Instrument = v }
func (m *Message) SetFinancingDetails(v financingdetails.FinancingDetails) { m.FinancingDetails = &v }
func (m *Message) SetOrderQtyData(v orderqtydata.OrderQtyData)             { m.OrderQtyData = &v }
func (m *Message) SetQtyType(v int)                                        { m.QtyType = &v }
func (m *Message) SetYieldData(v yielddata.YieldData)                      { m.YieldData = &v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings)                      { m.NoUnderlyings = v }
func (m *Message) SetUnderlyingTradingSessionID(v string)                  { m.UnderlyingTradingSessionID = &v }
func (m *Message) SetUnderlyingTradingSessionSubID(v string)               { m.UnderlyingTradingSessionSubID = &v }
func (m *Message) SetLastQty(v float64)                                    { m.LastQty = v }
func (m *Message) SetLastPx(v float64)                                     { m.LastPx = v }
func (m *Message) SetLastParPx(v float64)                                  { m.LastParPx = &v }
func (m *Message) SetLastSpotRate(v float64)                               { m.LastSpotRate = &v }
func (m *Message) SetLastForwardPoints(v float64)                          { m.LastForwardPoints = &v }
func (m *Message) SetLastMkt(v string)                                     { m.LastMkt = &v }
func (m *Message) SetTradeDate(v string)                                   { m.TradeDate = v }
func (m *Message) SetClearingBusinessDate(v string)                        { m.ClearingBusinessDate = &v }
func (m *Message) SetAvgPx(v float64)                                      { m.AvgPx = &v }
func (m *Message) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *Message) SetAvgPxIndicator(v int) { m.AvgPxIndicator = &v }
func (m *Message) SetPositionAmountData(v positionamountdata.PositionAmountData) {
	m.PositionAmountData = &v
}
func (m *Message) SetMultiLegReportingType(v string)                       { m.MultiLegReportingType = &v }
func (m *Message) SetTradeLegRefID(v string)                               { m.TradeLegRefID = &v }
func (m *Message) SetNoLegs(v []NoLegs)                                    { m.NoLegs = v }
func (m *Message) SetTransactTime(v time.Time)                             { m.TransactTime = v }
func (m *Message) SetTrdRegTimestamps(v trdregtimestamps.TrdRegTimestamps) { m.TrdRegTimestamps = &v }
func (m *Message) SetSettlType(v string)                                   { m.SettlType = &v }
func (m *Message) SetSettlDate(v string)                                   { m.SettlDate = &v }
func (m *Message) SetMatchStatus(v string)                                 { m.MatchStatus = &v }
func (m *Message) SetMatchType(v string)                                   { m.MatchType = &v }
func (m *Message) SetNoSides(v []NoSides)                                  { m.NoSides = v }

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
	return enum.BeginStringFIX44, "AE", r
}
