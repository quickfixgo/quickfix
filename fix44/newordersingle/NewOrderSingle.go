//Package newordersingle msg type = D.
package newordersingle

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/commissiondata"
	"github.com/quickfixgo/quickfix/fix44/discretioninstructions"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/nestedparties"
	"github.com/quickfixgo/quickfix/fix44/orderqtydata"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/peginstructions"
	"github.com/quickfixgo/quickfix/fix44/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix44/stipulations"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fix44/yielddata"
	"time"
)

//NoAllocs is a repeating group in NewOrderSingle
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for NoAllocs.
	AllocAcctIDSource *int `fix:"661"`
	//AllocSettlCurrency is a non-required field for NoAllocs.
	AllocSettlCurrency *string `fix:"736"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//NestedParties is a non-required component for NoAllocs.
	NestedParties *nestedparties.NestedParties
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
}

//NewNoAllocs returns an initialized NoAllocs instance
func NewNoAllocs() *NoAllocs {
	var m NoAllocs
	return &m
}

func (m *NoAllocs) SetAllocAccount(v string)                       { m.AllocAccount = &v }
func (m *NoAllocs) SetAllocAcctIDSource(v int)                     { m.AllocAcctIDSource = &v }
func (m *NoAllocs) SetAllocSettlCurrency(v string)                 { m.AllocSettlCurrency = &v }
func (m *NoAllocs) SetIndividualAllocID(v string)                  { m.IndividualAllocID = &v }
func (m *NoAllocs) SetNestedParties(v nestedparties.NestedParties) { m.NestedParties = &v }
func (m *NoAllocs) SetAllocQty(v float64)                          { m.AllocQty = &v }

//NoTradingSessions is a repeating group in NewOrderSingle
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

//NewNoTradingSessions returns an initialized NoTradingSessions instance
func NewNoTradingSessions() *NoTradingSessions {
	var m NoTradingSessions
	return &m
}

func (m *NoTradingSessions) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoTradingSessions) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }

//NoUnderlyings is a repeating group in NewOrderSingle
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

//Message is a NewOrderSingle FIX Message
type Message struct {
	FIXMsgType string `fix:"D"`
	fix44.Header
	//ClOrdID is a required field for NewOrderSingle.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NewOrderSingle.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for NewOrderSingle.
	ClOrdLinkID *string `fix:"583"`
	//Parties is a non-required component for NewOrderSingle.
	Parties *parties.Parties
	//TradeOriginationDate is a non-required field for NewOrderSingle.
	TradeOriginationDate *string `fix:"229"`
	//TradeDate is a non-required field for NewOrderSingle.
	TradeDate *string `fix:"75"`
	//Account is a non-required field for NewOrderSingle.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for NewOrderSingle.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for NewOrderSingle.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for NewOrderSingle.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for NewOrderSingle.
	BookingUnit *string `fix:"590"`
	//PreallocMethod is a non-required field for NewOrderSingle.
	PreallocMethod *string `fix:"591"`
	//AllocID is a non-required field for NewOrderSingle.
	AllocID *string `fix:"70"`
	//NoAllocs is a non-required field for NewOrderSingle.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	//SettlType is a non-required field for NewOrderSingle.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for NewOrderSingle.
	SettlDate *string `fix:"64"`
	//CashMargin is a non-required field for NewOrderSingle.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for NewOrderSingle.
	ClearingFeeIndicator *string `fix:"635"`
	//HandlInst is a non-required field for NewOrderSingle.
	HandlInst *string `fix:"21"`
	//ExecInst is a non-required field for NewOrderSingle.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for NewOrderSingle.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for NewOrderSingle.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for NewOrderSingle.
	ExDestination *string `fix:"100"`
	//NoTradingSessions is a non-required field for NewOrderSingle.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//ProcessCode is a non-required field for NewOrderSingle.
	ProcessCode *string `fix:"81"`
	//Instrument is a required component for NewOrderSingle.
	instrument.Instrument
	//FinancingDetails is a non-required component for NewOrderSingle.
	FinancingDetails *financingdetails.FinancingDetails
	//NoUnderlyings is a non-required field for NewOrderSingle.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//PrevClosePx is a non-required field for NewOrderSingle.
	PrevClosePx *float64 `fix:"140"`
	//Side is a required field for NewOrderSingle.
	Side string `fix:"54"`
	//LocateReqd is a non-required field for NewOrderSingle.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a required field for NewOrderSingle.
	TransactTime time.Time `fix:"60"`
	//Stipulations is a non-required component for NewOrderSingle.
	Stipulations *stipulations.Stipulations
	//QtyType is a non-required field for NewOrderSingle.
	QtyType *int `fix:"854"`
	//OrderQtyData is a required component for NewOrderSingle.
	orderqtydata.OrderQtyData
	//OrdType is a required field for NewOrderSingle.
	OrdType string `fix:"40"`
	//PriceType is a non-required field for NewOrderSingle.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NewOrderSingle.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NewOrderSingle.
	StopPx *float64 `fix:"99"`
	//SpreadOrBenchmarkCurveData is a non-required component for NewOrderSingle.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData is a non-required component for NewOrderSingle.
	YieldData *yielddata.YieldData
	//Currency is a non-required field for NewOrderSingle.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for NewOrderSingle.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for NewOrderSingle.
	SolicitedFlag *bool `fix:"377"`
	//IOIID is a non-required field for NewOrderSingle.
	IOIID *string `fix:"23"`
	//QuoteID is a non-required field for NewOrderSingle.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for NewOrderSingle.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for NewOrderSingle.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for NewOrderSingle.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for NewOrderSingle.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for NewOrderSingle.
	GTBookingInst *int `fix:"427"`
	//CommissionData is a non-required component for NewOrderSingle.
	CommissionData *commissiondata.CommissionData
	//OrderCapacity is a non-required field for NewOrderSingle.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for NewOrderSingle.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for NewOrderSingle.
	CustOrderCapacity *int `fix:"582"`
	//ForexReq is a non-required field for NewOrderSingle.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for NewOrderSingle.
	SettlCurrency *string `fix:"120"`
	//BookingType is a non-required field for NewOrderSingle.
	BookingType *int `fix:"775"`
	//Text is a non-required field for NewOrderSingle.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NewOrderSingle.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NewOrderSingle.
	EncodedText *string `fix:"355"`
	//SettlDate2 is a non-required field for NewOrderSingle.
	SettlDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NewOrderSingle.
	OrderQty2 *float64 `fix:"192"`
	//Price2 is a non-required field for NewOrderSingle.
	Price2 *float64 `fix:"640"`
	//PositionEffect is a non-required field for NewOrderSingle.
	PositionEffect *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for NewOrderSingle.
	CoveredOrUncovered *int `fix:"203"`
	//MaxShow is a non-required field for NewOrderSingle.
	MaxShow *float64 `fix:"210"`
	//PegInstructions is a non-required component for NewOrderSingle.
	PegInstructions *peginstructions.PegInstructions
	//DiscretionInstructions is a non-required component for NewOrderSingle.
	DiscretionInstructions *discretioninstructions.DiscretionInstructions
	//TargetStrategy is a non-required field for NewOrderSingle.
	TargetStrategy *int `fix:"847"`
	//TargetStrategyParameters is a non-required field for NewOrderSingle.
	TargetStrategyParameters *string `fix:"848"`
	//ParticipationRate is a non-required field for NewOrderSingle.
	ParticipationRate *float64 `fix:"849"`
	//CancellationRights is a non-required field for NewOrderSingle.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for NewOrderSingle.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for NewOrderSingle.
	RegistID *string `fix:"513"`
	//Designation is a non-required field for NewOrderSingle.
	Designation *string `fix:"494"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetClOrdID(v string)                                     { m.ClOrdID = v }
func (m *Message) SetSecondaryClOrdID(v string)                            { m.SecondaryClOrdID = &v }
func (m *Message) SetClOrdLinkID(v string)                                 { m.ClOrdLinkID = &v }
func (m *Message) SetParties(v parties.Parties)                            { m.Parties = &v }
func (m *Message) SetTradeOriginationDate(v string)                        { m.TradeOriginationDate = &v }
func (m *Message) SetTradeDate(v string)                                   { m.TradeDate = &v }
func (m *Message) SetAccount(v string)                                     { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)                                   { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)                                    { m.AccountType = &v }
func (m *Message) SetDayBookingInst(v string)                              { m.DayBookingInst = &v }
func (m *Message) SetBookingUnit(v string)                                 { m.BookingUnit = &v }
func (m *Message) SetPreallocMethod(v string)                              { m.PreallocMethod = &v }
func (m *Message) SetAllocID(v string)                                     { m.AllocID = &v }
func (m *Message) SetNoAllocs(v []NoAllocs)                                { m.NoAllocs = v }
func (m *Message) SetSettlType(v string)                                   { m.SettlType = &v }
func (m *Message) SetSettlDate(v string)                                   { m.SettlDate = &v }
func (m *Message) SetCashMargin(v string)                                  { m.CashMargin = &v }
func (m *Message) SetClearingFeeIndicator(v string)                        { m.ClearingFeeIndicator = &v }
func (m *Message) SetHandlInst(v string)                                   { m.HandlInst = &v }
func (m *Message) SetExecInst(v string)                                    { m.ExecInst = &v }
func (m *Message) SetMinQty(v float64)                                     { m.MinQty = &v }
func (m *Message) SetMaxFloor(v float64)                                   { m.MaxFloor = &v }
func (m *Message) SetExDestination(v string)                               { m.ExDestination = &v }
func (m *Message) SetNoTradingSessions(v []NoTradingSessions)              { m.NoTradingSessions = v }
func (m *Message) SetProcessCode(v string)                                 { m.ProcessCode = &v }
func (m *Message) SetInstrument(v instrument.Instrument)                   { m.Instrument = v }
func (m *Message) SetFinancingDetails(v financingdetails.FinancingDetails) { m.FinancingDetails = &v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings)                      { m.NoUnderlyings = v }
func (m *Message) SetPrevClosePx(v float64)                                { m.PrevClosePx = &v }
func (m *Message) SetSide(v string)                                        { m.Side = v }
func (m *Message) SetLocateReqd(v bool)                                    { m.LocateReqd = &v }
func (m *Message) SetTransactTime(v time.Time)                             { m.TransactTime = v }
func (m *Message) SetStipulations(v stipulations.Stipulations)             { m.Stipulations = &v }
func (m *Message) SetQtyType(v int)                                        { m.QtyType = &v }
func (m *Message) SetOrderQtyData(v orderqtydata.OrderQtyData)             { m.OrderQtyData = v }
func (m *Message) SetOrdType(v string)                                     { m.OrdType = v }
func (m *Message) SetPriceType(v int)                                      { m.PriceType = &v }
func (m *Message) SetPrice(v float64)                                      { m.Price = &v }
func (m *Message) SetStopPx(v float64)                                     { m.StopPx = &v }
func (m *Message) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *Message) SetYieldData(v yielddata.YieldData)                   { m.YieldData = &v }
func (m *Message) SetCurrency(v string)                                 { m.Currency = &v }
func (m *Message) SetComplianceID(v string)                             { m.ComplianceID = &v }
func (m *Message) SetSolicitedFlag(v bool)                              { m.SolicitedFlag = &v }
func (m *Message) SetIOIID(v string)                                    { m.IOIID = &v }
func (m *Message) SetQuoteID(v string)                                  { m.QuoteID = &v }
func (m *Message) SetTimeInForce(v string)                              { m.TimeInForce = &v }
func (m *Message) SetEffectiveTime(v time.Time)                         { m.EffectiveTime = &v }
func (m *Message) SetExpireDate(v string)                               { m.ExpireDate = &v }
func (m *Message) SetExpireTime(v time.Time)                            { m.ExpireTime = &v }
func (m *Message) SetGTBookingInst(v int)                               { m.GTBookingInst = &v }
func (m *Message) SetCommissionData(v commissiondata.CommissionData)    { m.CommissionData = &v }
func (m *Message) SetOrderCapacity(v string)                            { m.OrderCapacity = &v }
func (m *Message) SetOrderRestrictions(v string)                        { m.OrderRestrictions = &v }
func (m *Message) SetCustOrderCapacity(v int)                           { m.CustOrderCapacity = &v }
func (m *Message) SetForexReq(v bool)                                   { m.ForexReq = &v }
func (m *Message) SetSettlCurrency(v string)                            { m.SettlCurrency = &v }
func (m *Message) SetBookingType(v int)                                 { m.BookingType = &v }
func (m *Message) SetText(v string)                                     { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                              { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                              { m.EncodedText = &v }
func (m *Message) SetSettlDate2(v string)                               { m.SettlDate2 = &v }
func (m *Message) SetOrderQty2(v float64)                               { m.OrderQty2 = &v }
func (m *Message) SetPrice2(v float64)                                  { m.Price2 = &v }
func (m *Message) SetPositionEffect(v string)                           { m.PositionEffect = &v }
func (m *Message) SetCoveredOrUncovered(v int)                          { m.CoveredOrUncovered = &v }
func (m *Message) SetMaxShow(v float64)                                 { m.MaxShow = &v }
func (m *Message) SetPegInstructions(v peginstructions.PegInstructions) { m.PegInstructions = &v }
func (m *Message) SetDiscretionInstructions(v discretioninstructions.DiscretionInstructions) {
	m.DiscretionInstructions = &v
}
func (m *Message) SetTargetStrategy(v int)              { m.TargetStrategy = &v }
func (m *Message) SetTargetStrategyParameters(v string) { m.TargetStrategyParameters = &v }
func (m *Message) SetParticipationRate(v float64)       { m.ParticipationRate = &v }
func (m *Message) SetCancellationRights(v string)       { m.CancellationRights = &v }
func (m *Message) SetMoneyLaunderingStatus(v string)    { m.MoneyLaunderingStatus = &v }
func (m *Message) SetRegistID(v string)                 { m.RegistID = &v }
func (m *Message) SetDesignation(v string)              { m.Designation = &v }

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
	return enum.BeginStringFIX44, "D", r
}
