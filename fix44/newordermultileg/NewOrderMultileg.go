//Package newordermultileg msg type = AB.
package newordermultileg

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/commissiondata"
	"github.com/quickfixgo/quickfix/fix44/discretioninstructions"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/legstipulations"
	"github.com/quickfixgo/quickfix/fix44/nestedparties"
	"github.com/quickfixgo/quickfix/fix44/nestedparties2"
	"github.com/quickfixgo/quickfix/fix44/nestedparties3"
	"github.com/quickfixgo/quickfix/fix44/orderqtydata"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/peginstructions"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"time"
)

//NoAllocs is a repeating group in NewOrderMultileg
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for NoAllocs.
	AllocAcctIDSource *int `fix:"661"`
	//AllocSettlCurrency is a non-required field for NoAllocs.
	AllocSettlCurrency *string `fix:"736"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//NestedParties3 is a non-required component for NoAllocs.
	NestedParties3 *nestedparties3.NestedParties3
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
func (m *NoAllocs) SetNestedParties3(v nestedparties3.NestedParties3) { m.NestedParties3 = &v }
func (m *NoAllocs) SetAllocQty(v float64)                             { m.AllocQty = &v }

//NoTradingSessions is a repeating group in NewOrderMultileg
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

//NoUnderlyings is a repeating group in NewOrderMultileg
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

//NoLegs is a repeating group in NewOrderMultileg
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
	//LegQty is a non-required field for NoLegs.
	LegQty *float64 `fix:"687"`
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegStipulations is a non-required component for NoLegs.
	LegStipulations *legstipulations.LegStipulations
	//NoLegAllocs is a non-required field for NoLegs.
	NoLegAllocs []NoLegAllocs `fix:"670,omitempty"`
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
func (m *NoLegs) SetNoLegAllocs(v []NoLegAllocs)                       { m.NoLegAllocs = v }
func (m *NoLegs) SetLegPositionEffect(v string)                        { m.LegPositionEffect = &v }
func (m *NoLegs) SetLegCoveredOrUncovered(v int)                       { m.LegCoveredOrUncovered = &v }
func (m *NoLegs) SetNestedParties(v nestedparties.NestedParties)       { m.NestedParties = &v }
func (m *NoLegs) SetLegRefID(v string)                                 { m.LegRefID = &v }
func (m *NoLegs) SetLegPrice(v float64)                                { m.LegPrice = &v }
func (m *NoLegs) SetLegSettlType(v string)                             { m.LegSettlType = &v }
func (m *NoLegs) SetLegSettlDate(v string)                             { m.LegSettlDate = &v }

//NoLegAllocs is a repeating group in NoLegs
type NoLegAllocs struct {
	//LegAllocAccount is a non-required field for NoLegAllocs.
	LegAllocAccount *string `fix:"671"`
	//LegIndividualAllocID is a non-required field for NoLegAllocs.
	LegIndividualAllocID *string `fix:"672"`
	//NestedParties2 is a non-required component for NoLegAllocs.
	NestedParties2 *nestedparties2.NestedParties2
	//LegAllocQty is a non-required field for NoLegAllocs.
	LegAllocQty *float64 `fix:"673"`
	//LegAllocAcctIDSource is a non-required field for NoLegAllocs.
	LegAllocAcctIDSource *string `fix:"674"`
	//LegSettlCurrency is a non-required field for NoLegAllocs.
	LegSettlCurrency *string `fix:"675"`
}

//NewNoLegAllocs returns an initialized NoLegAllocs instance
func NewNoLegAllocs() *NoLegAllocs {
	var m NoLegAllocs
	return &m
}

func (m *NoLegAllocs) SetLegAllocAccount(v string)                       { m.LegAllocAccount = &v }
func (m *NoLegAllocs) SetLegIndividualAllocID(v string)                  { m.LegIndividualAllocID = &v }
func (m *NoLegAllocs) SetNestedParties2(v nestedparties2.NestedParties2) { m.NestedParties2 = &v }
func (m *NoLegAllocs) SetLegAllocQty(v float64)                          { m.LegAllocQty = &v }
func (m *NoLegAllocs) SetLegAllocAcctIDSource(v string)                  { m.LegAllocAcctIDSource = &v }
func (m *NoLegAllocs) SetLegSettlCurrency(v string)                      { m.LegSettlCurrency = &v }

//Message is a NewOrderMultileg FIX Message
type Message struct {
	FIXMsgType string `fix:"AB"`
	fix44.Header
	//ClOrdID is a required field for NewOrderMultileg.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NewOrderMultileg.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for NewOrderMultileg.
	ClOrdLinkID *string `fix:"583"`
	//Parties is a non-required component for NewOrderMultileg.
	Parties *parties.Parties
	//TradeOriginationDate is a non-required field for NewOrderMultileg.
	TradeOriginationDate *string `fix:"229"`
	//TradeDate is a non-required field for NewOrderMultileg.
	TradeDate *string `fix:"75"`
	//Account is a non-required field for NewOrderMultileg.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for NewOrderMultileg.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for NewOrderMultileg.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for NewOrderMultileg.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for NewOrderMultileg.
	BookingUnit *string `fix:"590"`
	//PreallocMethod is a non-required field for NewOrderMultileg.
	PreallocMethod *string `fix:"591"`
	//AllocID is a non-required field for NewOrderMultileg.
	AllocID *string `fix:"70"`
	//NoAllocs is a non-required field for NewOrderMultileg.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	//SettlType is a non-required field for NewOrderMultileg.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for NewOrderMultileg.
	SettlDate *string `fix:"64"`
	//CashMargin is a non-required field for NewOrderMultileg.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for NewOrderMultileg.
	ClearingFeeIndicator *string `fix:"635"`
	//HandlInst is a non-required field for NewOrderMultileg.
	HandlInst *string `fix:"21"`
	//ExecInst is a non-required field for NewOrderMultileg.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for NewOrderMultileg.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for NewOrderMultileg.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for NewOrderMultileg.
	ExDestination *string `fix:"100"`
	//NoTradingSessions is a non-required field for NewOrderMultileg.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//ProcessCode is a non-required field for NewOrderMultileg.
	ProcessCode *string `fix:"81"`
	//Side is a required field for NewOrderMultileg.
	Side string `fix:"54"`
	//Instrument is a required component for NewOrderMultileg.
	instrument.Instrument
	//NoUnderlyings is a non-required field for NewOrderMultileg.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//PrevClosePx is a non-required field for NewOrderMultileg.
	PrevClosePx *float64 `fix:"140"`
	//NoLegs is a required field for NewOrderMultileg.
	NoLegs []NoLegs `fix:"555"`
	//LocateReqd is a non-required field for NewOrderMultileg.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a required field for NewOrderMultileg.
	TransactTime time.Time `fix:"60"`
	//QtyType is a non-required field for NewOrderMultileg.
	QtyType *int `fix:"854"`
	//OrderQtyData is a required component for NewOrderMultileg.
	orderqtydata.OrderQtyData
	//OrdType is a required field for NewOrderMultileg.
	OrdType string `fix:"40"`
	//PriceType is a non-required field for NewOrderMultileg.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NewOrderMultileg.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NewOrderMultileg.
	StopPx *float64 `fix:"99"`
	//Currency is a non-required field for NewOrderMultileg.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for NewOrderMultileg.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for NewOrderMultileg.
	SolicitedFlag *bool `fix:"377"`
	//IOIID is a non-required field for NewOrderMultileg.
	IOIID *string `fix:"23"`
	//QuoteID is a non-required field for NewOrderMultileg.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for NewOrderMultileg.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for NewOrderMultileg.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for NewOrderMultileg.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for NewOrderMultileg.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for NewOrderMultileg.
	GTBookingInst *int `fix:"427"`
	//CommissionData is a non-required component for NewOrderMultileg.
	CommissionData *commissiondata.CommissionData
	//OrderCapacity is a non-required field for NewOrderMultileg.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for NewOrderMultileg.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for NewOrderMultileg.
	CustOrderCapacity *int `fix:"582"`
	//ForexReq is a non-required field for NewOrderMultileg.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for NewOrderMultileg.
	SettlCurrency *string `fix:"120"`
	//BookingType is a non-required field for NewOrderMultileg.
	BookingType *int `fix:"775"`
	//Text is a non-required field for NewOrderMultileg.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NewOrderMultileg.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NewOrderMultileg.
	EncodedText *string `fix:"355"`
	//PositionEffect is a non-required field for NewOrderMultileg.
	PositionEffect *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for NewOrderMultileg.
	CoveredOrUncovered *int `fix:"203"`
	//MaxShow is a non-required field for NewOrderMultileg.
	MaxShow *float64 `fix:"210"`
	//PegInstructions is a non-required component for NewOrderMultileg.
	PegInstructions *peginstructions.PegInstructions
	//DiscretionInstructions is a non-required component for NewOrderMultileg.
	DiscretionInstructions *discretioninstructions.DiscretionInstructions
	//TargetStrategy is a non-required field for NewOrderMultileg.
	TargetStrategy *int `fix:"847"`
	//TargetStrategyParameters is a non-required field for NewOrderMultileg.
	TargetStrategyParameters *string `fix:"848"`
	//ParticipationRate is a non-required field for NewOrderMultileg.
	ParticipationRate *float64 `fix:"849"`
	//CancellationRights is a non-required field for NewOrderMultileg.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for NewOrderMultileg.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for NewOrderMultileg.
	RegistID *string `fix:"513"`
	//Designation is a non-required field for NewOrderMultileg.
	Designation *string `fix:"494"`
	//MultiLegRptTypeReq is a non-required field for NewOrderMultileg.
	MultiLegRptTypeReq *int `fix:"563"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetClOrdID(v string)                                  { m.ClOrdID = v }
func (m *Message) SetSecondaryClOrdID(v string)                         { m.SecondaryClOrdID = &v }
func (m *Message) SetClOrdLinkID(v string)                              { m.ClOrdLinkID = &v }
func (m *Message) SetParties(v parties.Parties)                         { m.Parties = &v }
func (m *Message) SetTradeOriginationDate(v string)                     { m.TradeOriginationDate = &v }
func (m *Message) SetTradeDate(v string)                                { m.TradeDate = &v }
func (m *Message) SetAccount(v string)                                  { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)                                { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)                                 { m.AccountType = &v }
func (m *Message) SetDayBookingInst(v string)                           { m.DayBookingInst = &v }
func (m *Message) SetBookingUnit(v string)                              { m.BookingUnit = &v }
func (m *Message) SetPreallocMethod(v string)                           { m.PreallocMethod = &v }
func (m *Message) SetAllocID(v string)                                  { m.AllocID = &v }
func (m *Message) SetNoAllocs(v []NoAllocs)                             { m.NoAllocs = v }
func (m *Message) SetSettlType(v string)                                { m.SettlType = &v }
func (m *Message) SetSettlDate(v string)                                { m.SettlDate = &v }
func (m *Message) SetCashMargin(v string)                               { m.CashMargin = &v }
func (m *Message) SetClearingFeeIndicator(v string)                     { m.ClearingFeeIndicator = &v }
func (m *Message) SetHandlInst(v string)                                { m.HandlInst = &v }
func (m *Message) SetExecInst(v string)                                 { m.ExecInst = &v }
func (m *Message) SetMinQty(v float64)                                  { m.MinQty = &v }
func (m *Message) SetMaxFloor(v float64)                                { m.MaxFloor = &v }
func (m *Message) SetExDestination(v string)                            { m.ExDestination = &v }
func (m *Message) SetNoTradingSessions(v []NoTradingSessions)           { m.NoTradingSessions = v }
func (m *Message) SetProcessCode(v string)                              { m.ProcessCode = &v }
func (m *Message) SetSide(v string)                                     { m.Side = v }
func (m *Message) SetInstrument(v instrument.Instrument)                { m.Instrument = v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings)                   { m.NoUnderlyings = v }
func (m *Message) SetPrevClosePx(v float64)                             { m.PrevClosePx = &v }
func (m *Message) SetNoLegs(v []NoLegs)                                 { m.NoLegs = v }
func (m *Message) SetLocateReqd(v bool)                                 { m.LocateReqd = &v }
func (m *Message) SetTransactTime(v time.Time)                          { m.TransactTime = v }
func (m *Message) SetQtyType(v int)                                     { m.QtyType = &v }
func (m *Message) SetOrderQtyData(v orderqtydata.OrderQtyData)          { m.OrderQtyData = v }
func (m *Message) SetOrdType(v string)                                  { m.OrdType = v }
func (m *Message) SetPriceType(v int)                                   { m.PriceType = &v }
func (m *Message) SetPrice(v float64)                                   { m.Price = &v }
func (m *Message) SetStopPx(v float64)                                  { m.StopPx = &v }
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
func (m *Message) SetMultiLegRptTypeReq(v int)          { m.MultiLegRptTypeReq = &v }

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
	return enum.BeginStringFIX44, "AB", r
}
