//Package executionreport msg type = 8.
package executionreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/commissiondata"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/instrumentleg"
	"github.com/quickfixgo/quickfix/fix43/nestedparties"
	"github.com/quickfixgo/quickfix/fix43/orderqtydata"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"github.com/quickfixgo/quickfix/fix43/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix43/stipulations"
	"github.com/quickfixgo/quickfix/fix43/yielddata"
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
	//ContraLegRefID is a non-required field for NoContraBrokers.
	ContraLegRefID *string `fix:"655"`
}

func (m *NoContraBrokers) SetContraBroker(v string)       { m.ContraBroker = &v }
func (m *NoContraBrokers) SetContraTrader(v string)       { m.ContraTrader = &v }
func (m *NoContraBrokers) SetContraTradeQty(v float64)    { m.ContraTradeQty = &v }
func (m *NoContraBrokers) SetContraTradeTime(v time.Time) { m.ContraTradeTime = &v }
func (m *NoContraBrokers) SetContraLegRefID(v string)     { m.ContraLegRefID = &v }

//NoContAmts is a repeating group in ExecutionReport
type NoContAmts struct {
	//ContAmtType is a non-required field for NoContAmts.
	ContAmtType *int `fix:"519"`
	//ContAmtValue is a non-required field for NoContAmts.
	ContAmtValue *float64 `fix:"520"`
	//ContAmtCurr is a non-required field for NoContAmts.
	ContAmtCurr *string `fix:"521"`
}

func (m *NoContAmts) SetContAmtType(v int)      { m.ContAmtType = &v }
func (m *NoContAmts) SetContAmtValue(v float64) { m.ContAmtValue = &v }
func (m *NoContAmts) SetContAmtCurr(v string)   { m.ContAmtCurr = &v }

//NoLegs is a repeating group in ExecutionReport
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
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
	//LegSettlmntTyp is a non-required field for NoLegs.
	LegSettlmntTyp *string `fix:"587"`
	//LegFutSettDate is a non-required field for NoLegs.
	LegFutSettDate *string `fix:"588"`
	//LegLastPx is a non-required field for NoLegs.
	LegLastPx *float64 `fix:"637"`
}

func (m *NoLegs) SetLegPositionEffect(v string)  { m.LegPositionEffect = &v }
func (m *NoLegs) SetLegCoveredOrUncovered(v int) { m.LegCoveredOrUncovered = &v }
func (m *NoLegs) SetLegRefID(v string)           { m.LegRefID = &v }
func (m *NoLegs) SetLegPrice(v float64)          { m.LegPrice = &v }
func (m *NoLegs) SetLegSettlmntTyp(v string)     { m.LegSettlmntTyp = &v }
func (m *NoLegs) SetLegFutSettDate(v string)     { m.LegFutSettDate = &v }
func (m *NoLegs) SetLegLastPx(v float64)         { m.LegLastPx = &v }

//Message is a ExecutionReport FIX Message
type Message struct {
	FIXMsgType string `fix:"8"`
	Header     fix43.Header
	//OrderID is a required field for ExecutionReport.
	OrderID string `fix:"37"`
	//SecondaryOrderID is a non-required field for ExecutionReport.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for ExecutionReport.
	SecondaryClOrdID *string `fix:"526"`
	//SecondaryExecID is a non-required field for ExecutionReport.
	SecondaryExecID *string `fix:"527"`
	//ClOrdID is a non-required field for ExecutionReport.
	ClOrdID *string `fix:"11"`
	//OrigClOrdID is a non-required field for ExecutionReport.
	OrigClOrdID *string `fix:"41"`
	//ClOrdLinkID is a non-required field for ExecutionReport.
	ClOrdLinkID *string `fix:"583"`
	//Parties Component
	Parties parties.Component
	//TradeOriginationDate is a non-required field for ExecutionReport.
	TradeOriginationDate *string `fix:"229"`
	//NoContraBrokers is a non-required field for ExecutionReport.
	NoContraBrokers []NoContraBrokers `fix:"382,omitempty"`
	//ListID is a non-required field for ExecutionReport.
	ListID *string `fix:"66"`
	//CrossID is a non-required field for ExecutionReport.
	CrossID *string `fix:"548"`
	//OrigCrossID is a non-required field for ExecutionReport.
	OrigCrossID *string `fix:"551"`
	//CrossType is a non-required field for ExecutionReport.
	CrossType *int `fix:"549"`
	//ExecID is a required field for ExecutionReport.
	ExecID string `fix:"17"`
	//ExecRefID is a non-required field for ExecutionReport.
	ExecRefID *string `fix:"19"`
	//ExecType is a required field for ExecutionReport.
	ExecType string `fix:"150"`
	//OrdStatus is a required field for ExecutionReport.
	OrdStatus string `fix:"39"`
	//WorkingIndicator is a non-required field for ExecutionReport.
	WorkingIndicator *bool `fix:"636"`
	//OrdRejReason is a non-required field for ExecutionReport.
	OrdRejReason *int `fix:"103"`
	//ExecRestatementReason is a non-required field for ExecutionReport.
	ExecRestatementReason *int `fix:"378"`
	//Account is a non-required field for ExecutionReport.
	Account *string `fix:"1"`
	//AccountType is a non-required field for ExecutionReport.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for ExecutionReport.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for ExecutionReport.
	BookingUnit *string `fix:"590"`
	//PreallocMethod is a non-required field for ExecutionReport.
	PreallocMethod *string `fix:"591"`
	//SettlmntTyp is a non-required field for ExecutionReport.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for ExecutionReport.
	FutSettDate *string `fix:"64"`
	//CashMargin is a non-required field for ExecutionReport.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for ExecutionReport.
	ClearingFeeIndicator *string `fix:"635"`
	//Instrument Component
	Instrument instrument.Component
	//Side is a required field for ExecutionReport.
	Side string `fix:"54"`
	//Stipulations Component
	Stipulations stipulations.Component
	//QuantityType is a non-required field for ExecutionReport.
	QuantityType *int `fix:"465"`
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//OrdType is a non-required field for ExecutionReport.
	OrdType *string `fix:"40"`
	//PriceType is a non-required field for ExecutionReport.
	PriceType *int `fix:"423"`
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
	//OrderCapacity is a non-required field for ExecutionReport.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for ExecutionReport.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for ExecutionReport.
	CustOrderCapacity *int `fix:"582"`
	//Rule80A is a non-required field for ExecutionReport.
	Rule80A *string `fix:"47"`
	//LastQty is a non-required field for ExecutionReport.
	LastQty *float64 `fix:"32"`
	//UnderlyingLastQty is a non-required field for ExecutionReport.
	UnderlyingLastQty *float64 `fix:"652"`
	//LastPx is a non-required field for ExecutionReport.
	LastPx *float64 `fix:"31"`
	//UnderlyingLastPx is a non-required field for ExecutionReport.
	UnderlyingLastPx *float64 `fix:"651"`
	//LastSpotRate is a non-required field for ExecutionReport.
	LastSpotRate *float64 `fix:"194"`
	//LastForwardPoints is a non-required field for ExecutionReport.
	LastForwardPoints *float64 `fix:"195"`
	//LastMkt is a non-required field for ExecutionReport.
	LastMkt *string `fix:"30"`
	//TradingSessionID is a non-required field for ExecutionReport.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for ExecutionReport.
	TradingSessionSubID *string `fix:"625"`
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
	//CommissionData Component
	CommissionData commissiondata.Component
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//YieldData Component
	YieldData yielddata.Component
	//GrossTradeAmt is a non-required field for ExecutionReport.
	GrossTradeAmt *float64 `fix:"381"`
	//NumDaysInterest is a non-required field for ExecutionReport.
	NumDaysInterest *int `fix:"157"`
	//ExDate is a non-required field for ExecutionReport.
	ExDate *string `fix:"230"`
	//AccruedInterestRate is a non-required field for ExecutionReport.
	AccruedInterestRate *float64 `fix:"158"`
	//AccruedInterestAmt is a non-required field for ExecutionReport.
	AccruedInterestAmt *float64 `fix:"159"`
	//TradedFlatSwitch is a non-required field for ExecutionReport.
	TradedFlatSwitch *bool `fix:"258"`
	//BasisFeatureDate is a non-required field for ExecutionReport.
	BasisFeatureDate *string `fix:"259"`
	//BasisFeaturePrice is a non-required field for ExecutionReport.
	BasisFeaturePrice *float64 `fix:"260"`
	//Concession is a non-required field for ExecutionReport.
	Concession *float64 `fix:"238"`
	//TotalTakedown is a non-required field for ExecutionReport.
	TotalTakedown *float64 `fix:"237"`
	//NetMoney is a non-required field for ExecutionReport.
	NetMoney *float64 `fix:"118"`
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
	//PositionEffect is a non-required field for ExecutionReport.
	PositionEffect *string `fix:"77"`
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
	//LastForwardPoints2 is a non-required field for ExecutionReport.
	LastForwardPoints2 *float64 `fix:"641"`
	//MultiLegReportingType is a non-required field for ExecutionReport.
	MultiLegReportingType *string `fix:"442"`
	//CancellationRights is a non-required field for ExecutionReport.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for ExecutionReport.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for ExecutionReport.
	RegistID *string `fix:"513"`
	//Designation is a non-required field for ExecutionReport.
	Designation *string `fix:"494"`
	//TransBkdTime is a non-required field for ExecutionReport.
	TransBkdTime *time.Time `fix:"483"`
	//ExecValuationPoint is a non-required field for ExecutionReport.
	ExecValuationPoint *time.Time `fix:"515"`
	//ExecPriceType is a non-required field for ExecutionReport.
	ExecPriceType *string `fix:"484"`
	//ExecPriceAdjustment is a non-required field for ExecutionReport.
	ExecPriceAdjustment *float64 `fix:"485"`
	//PriorityIndicator is a non-required field for ExecutionReport.
	PriorityIndicator *int `fix:"638"`
	//PriceImprovement is a non-required field for ExecutionReport.
	PriceImprovement *float64 `fix:"639"`
	//NoContAmts is a non-required field for ExecutionReport.
	NoContAmts []NoContAmts `fix:"518,omitempty"`
	//NoLegs is a non-required field for ExecutionReport.
	NoLegs  []NoLegs `fix:"555,omitempty"`
	Trailer fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrderID(v string)                    { m.OrderID = v }
func (m *Message) SetSecondaryOrderID(v string)           { m.SecondaryOrderID = &v }
func (m *Message) SetSecondaryClOrdID(v string)           { m.SecondaryClOrdID = &v }
func (m *Message) SetSecondaryExecID(v string)            { m.SecondaryExecID = &v }
func (m *Message) SetClOrdID(v string)                    { m.ClOrdID = &v }
func (m *Message) SetOrigClOrdID(v string)                { m.OrigClOrdID = &v }
func (m *Message) SetClOrdLinkID(v string)                { m.ClOrdLinkID = &v }
func (m *Message) SetTradeOriginationDate(v string)       { m.TradeOriginationDate = &v }
func (m *Message) SetNoContraBrokers(v []NoContraBrokers) { m.NoContraBrokers = v }
func (m *Message) SetListID(v string)                     { m.ListID = &v }
func (m *Message) SetCrossID(v string)                    { m.CrossID = &v }
func (m *Message) SetOrigCrossID(v string)                { m.OrigCrossID = &v }
func (m *Message) SetCrossType(v int)                     { m.CrossType = &v }
func (m *Message) SetExecID(v string)                     { m.ExecID = v }
func (m *Message) SetExecRefID(v string)                  { m.ExecRefID = &v }
func (m *Message) SetExecType(v string)                   { m.ExecType = v }
func (m *Message) SetOrdStatus(v string)                  { m.OrdStatus = v }
func (m *Message) SetWorkingIndicator(v bool)             { m.WorkingIndicator = &v }
func (m *Message) SetOrdRejReason(v int)                  { m.OrdRejReason = &v }
func (m *Message) SetExecRestatementReason(v int)         { m.ExecRestatementReason = &v }
func (m *Message) SetAccount(v string)                    { m.Account = &v }
func (m *Message) SetAccountType(v int)                   { m.AccountType = &v }
func (m *Message) SetDayBookingInst(v string)             { m.DayBookingInst = &v }
func (m *Message) SetBookingUnit(v string)                { m.BookingUnit = &v }
func (m *Message) SetPreallocMethod(v string)             { m.PreallocMethod = &v }
func (m *Message) SetSettlmntTyp(v string)                { m.SettlmntTyp = &v }
func (m *Message) SetFutSettDate(v string)                { m.FutSettDate = &v }
func (m *Message) SetCashMargin(v string)                 { m.CashMargin = &v }
func (m *Message) SetClearingFeeIndicator(v string)       { m.ClearingFeeIndicator = &v }
func (m *Message) SetSide(v string)                       { m.Side = v }
func (m *Message) SetQuantityType(v int)                  { m.QuantityType = &v }
func (m *Message) SetOrdType(v string)                    { m.OrdType = &v }
func (m *Message) SetPriceType(v int)                     { m.PriceType = &v }
func (m *Message) SetPrice(v float64)                     { m.Price = &v }
func (m *Message) SetStopPx(v float64)                    { m.StopPx = &v }
func (m *Message) SetPegDifference(v float64)             { m.PegDifference = &v }
func (m *Message) SetDiscretionInst(v string)             { m.DiscretionInst = &v }
func (m *Message) SetDiscretionOffset(v float64)          { m.DiscretionOffset = &v }
func (m *Message) SetCurrency(v string)                   { m.Currency = &v }
func (m *Message) SetComplianceID(v string)               { m.ComplianceID = &v }
func (m *Message) SetSolicitedFlag(v bool)                { m.SolicitedFlag = &v }
func (m *Message) SetTimeInForce(v string)                { m.TimeInForce = &v }
func (m *Message) SetEffectiveTime(v time.Time)           { m.EffectiveTime = &v }
func (m *Message) SetExpireDate(v string)                 { m.ExpireDate = &v }
func (m *Message) SetExpireTime(v time.Time)              { m.ExpireTime = &v }
func (m *Message) SetExecInst(v string)                   { m.ExecInst = &v }
func (m *Message) SetOrderCapacity(v string)              { m.OrderCapacity = &v }
func (m *Message) SetOrderRestrictions(v string)          { m.OrderRestrictions = &v }
func (m *Message) SetCustOrderCapacity(v int)             { m.CustOrderCapacity = &v }
func (m *Message) SetRule80A(v string)                    { m.Rule80A = &v }
func (m *Message) SetLastQty(v float64)                   { m.LastQty = &v }
func (m *Message) SetUnderlyingLastQty(v float64)         { m.UnderlyingLastQty = &v }
func (m *Message) SetLastPx(v float64)                    { m.LastPx = &v }
func (m *Message) SetUnderlyingLastPx(v float64)          { m.UnderlyingLastPx = &v }
func (m *Message) SetLastSpotRate(v float64)              { m.LastSpotRate = &v }
func (m *Message) SetLastForwardPoints(v float64)         { m.LastForwardPoints = &v }
func (m *Message) SetLastMkt(v string)                    { m.LastMkt = &v }
func (m *Message) SetTradingSessionID(v string)           { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)        { m.TradingSessionSubID = &v }
func (m *Message) SetLastCapacity(v string)               { m.LastCapacity = &v }
func (m *Message) SetLeavesQty(v float64)                 { m.LeavesQty = v }
func (m *Message) SetCumQty(v float64)                    { m.CumQty = v }
func (m *Message) SetAvgPx(v float64)                     { m.AvgPx = v }
func (m *Message) SetDayOrderQty(v float64)               { m.DayOrderQty = &v }
func (m *Message) SetDayCumQty(v float64)                 { m.DayCumQty = &v }
func (m *Message) SetDayAvgPx(v float64)                  { m.DayAvgPx = &v }
func (m *Message) SetGTBookingInst(v int)                 { m.GTBookingInst = &v }
func (m *Message) SetTradeDate(v string)                  { m.TradeDate = &v }
func (m *Message) SetTransactTime(v time.Time)            { m.TransactTime = &v }
func (m *Message) SetReportToExch(v bool)                 { m.ReportToExch = &v }
func (m *Message) SetGrossTradeAmt(v float64)             { m.GrossTradeAmt = &v }
func (m *Message) SetNumDaysInterest(v int)               { m.NumDaysInterest = &v }
func (m *Message) SetExDate(v string)                     { m.ExDate = &v }
func (m *Message) SetAccruedInterestRate(v float64)       { m.AccruedInterestRate = &v }
func (m *Message) SetAccruedInterestAmt(v float64)        { m.AccruedInterestAmt = &v }
func (m *Message) SetTradedFlatSwitch(v bool)             { m.TradedFlatSwitch = &v }
func (m *Message) SetBasisFeatureDate(v string)           { m.BasisFeatureDate = &v }
func (m *Message) SetBasisFeaturePrice(v float64)         { m.BasisFeaturePrice = &v }
func (m *Message) SetConcession(v float64)                { m.Concession = &v }
func (m *Message) SetTotalTakedown(v float64)             { m.TotalTakedown = &v }
func (m *Message) SetNetMoney(v float64)                  { m.NetMoney = &v }
func (m *Message) SetSettlCurrAmt(v float64)              { m.SettlCurrAmt = &v }
func (m *Message) SetSettlCurrency(v string)              { m.SettlCurrency = &v }
func (m *Message) SetSettlCurrFxRate(v float64)           { m.SettlCurrFxRate = &v }
func (m *Message) SetSettlCurrFxRateCalc(v string)        { m.SettlCurrFxRateCalc = &v }
func (m *Message) SetHandlInst(v string)                  { m.HandlInst = &v }
func (m *Message) SetMinQty(v float64)                    { m.MinQty = &v }
func (m *Message) SetMaxFloor(v float64)                  { m.MaxFloor = &v }
func (m *Message) SetPositionEffect(v string)             { m.PositionEffect = &v }
func (m *Message) SetMaxShow(v float64)                   { m.MaxShow = &v }
func (m *Message) SetText(v string)                       { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                { m.EncodedText = &v }
func (m *Message) SetFutSettDate2(v string)               { m.FutSettDate2 = &v }
func (m *Message) SetOrderQty2(v float64)                 { m.OrderQty2 = &v }
func (m *Message) SetLastForwardPoints2(v float64)        { m.LastForwardPoints2 = &v }
func (m *Message) SetMultiLegReportingType(v string)      { m.MultiLegReportingType = &v }
func (m *Message) SetCancellationRights(v string)         { m.CancellationRights = &v }
func (m *Message) SetMoneyLaunderingStatus(v string)      { m.MoneyLaunderingStatus = &v }
func (m *Message) SetRegistID(v string)                   { m.RegistID = &v }
func (m *Message) SetDesignation(v string)                { m.Designation = &v }
func (m *Message) SetTransBkdTime(v time.Time)            { m.TransBkdTime = &v }
func (m *Message) SetExecValuationPoint(v time.Time)      { m.ExecValuationPoint = &v }
func (m *Message) SetExecPriceType(v string)              { m.ExecPriceType = &v }
func (m *Message) SetExecPriceAdjustment(v float64)       { m.ExecPriceAdjustment = &v }
func (m *Message) SetPriorityIndicator(v int)             { m.PriorityIndicator = &v }
func (m *Message) SetPriceImprovement(v float64)          { m.PriceImprovement = &v }
func (m *Message) SetNoContAmts(v []NoContAmts)           { m.NoContAmts = v }
func (m *Message) SetNoLegs(v []NoLegs)                   { m.NoLegs = v }

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
	return enum.BeginStringFIX43, "8", r
}
