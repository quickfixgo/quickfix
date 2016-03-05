//Package neworderlist msg type = E.
package neworderlist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/commissiondata"
	"github.com/quickfixgo/quickfix/fix44/discretioninstructions"
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

//NoOrders is a repeating group in NewOrderList
type NoOrders struct {
	//ClOrdID is a required field for NoOrders.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoOrders.
	SecondaryClOrdID *string `fix:"526"`
	//ListSeqNo is a required field for NoOrders.
	ListSeqNo int `fix:"67"`
	//ClOrdLinkID is a non-required field for NoOrders.
	ClOrdLinkID *string `fix:"583"`
	//SettlInstMode is a non-required field for NoOrders.
	SettlInstMode *string `fix:"160"`
	//Parties Component
	parties.Parties
	//TradeOriginationDate is a non-required field for NoOrders.
	TradeOriginationDate *string `fix:"229"`
	//TradeDate is a non-required field for NoOrders.
	TradeDate *string `fix:"75"`
	//Account is a non-required field for NoOrders.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for NoOrders.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for NoOrders.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for NoOrders.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for NoOrders.
	BookingUnit *string `fix:"590"`
	//AllocID is a non-required field for NoOrders.
	AllocID *string `fix:"70"`
	//PreallocMethod is a non-required field for NoOrders.
	PreallocMethod *string `fix:"591"`
	//NoAllocs is a non-required field for NoOrders.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	//SettlType is a non-required field for NoOrders.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for NoOrders.
	SettlDate *string `fix:"64"`
	//CashMargin is a non-required field for NoOrders.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for NoOrders.
	ClearingFeeIndicator *string `fix:"635"`
	//HandlInst is a non-required field for NoOrders.
	HandlInst *string `fix:"21"`
	//ExecInst is a non-required field for NoOrders.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for NoOrders.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for NoOrders.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for NoOrders.
	ExDestination *string `fix:"100"`
	//NoTradingSessions is a non-required field for NoOrders.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//ProcessCode is a non-required field for NoOrders.
	ProcessCode *string `fix:"81"`
	//Instrument Component
	instrument.Instrument
	//NoUnderlyings is a non-required field for NoOrders.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//PrevClosePx is a non-required field for NoOrders.
	PrevClosePx *float64 `fix:"140"`
	//Side is a required field for NoOrders.
	Side string `fix:"54"`
	//SideValueInd is a non-required field for NoOrders.
	SideValueInd *int `fix:"401"`
	//LocateReqd is a non-required field for NoOrders.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a non-required field for NoOrders.
	TransactTime *time.Time `fix:"60"`
	//Stipulations Component
	stipulations.Stipulations
	//QtyType is a non-required field for NoOrders.
	QtyType *int `fix:"854"`
	//OrderQtyData Component
	orderqtydata.OrderQtyData
	//OrdType is a non-required field for NoOrders.
	OrdType *string `fix:"40"`
	//PriceType is a non-required field for NoOrders.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NoOrders.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NoOrders.
	StopPx *float64 `fix:"99"`
	//SpreadOrBenchmarkCurveData Component
	spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData Component
	yielddata.YieldData
	//Currency is a non-required field for NoOrders.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for NoOrders.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for NoOrders.
	SolicitedFlag *bool `fix:"377"`
	//IOIID is a non-required field for NoOrders.
	IOIID *string `fix:"23"`
	//QuoteID is a non-required field for NoOrders.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for NoOrders.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for NoOrders.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for NoOrders.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for NoOrders.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for NoOrders.
	GTBookingInst *int `fix:"427"`
	//CommissionData Component
	commissiondata.CommissionData
	//OrderCapacity is a non-required field for NoOrders.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for NoOrders.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for NoOrders.
	CustOrderCapacity *int `fix:"582"`
	//ForexReq is a non-required field for NoOrders.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for NoOrders.
	SettlCurrency *string `fix:"120"`
	//BookingType is a non-required field for NoOrders.
	BookingType *int `fix:"775"`
	//Text is a non-required field for NoOrders.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoOrders.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoOrders.
	EncodedText *string `fix:"355"`
	//SettlDate2 is a non-required field for NoOrders.
	SettlDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NoOrders.
	OrderQty2 *float64 `fix:"192"`
	//Price2 is a non-required field for NoOrders.
	Price2 *float64 `fix:"640"`
	//PositionEffect is a non-required field for NoOrders.
	PositionEffect *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for NoOrders.
	CoveredOrUncovered *int `fix:"203"`
	//MaxShow is a non-required field for NoOrders.
	MaxShow *float64 `fix:"210"`
	//PegInstructions Component
	peginstructions.PegInstructions
	//DiscretionInstructions Component
	discretioninstructions.DiscretionInstructions
	//TargetStrategy is a non-required field for NoOrders.
	TargetStrategy *int `fix:"847"`
	//TargetStrategyParameters is a non-required field for NoOrders.
	TargetStrategyParameters *string `fix:"848"`
	//ParticipationRate is a non-required field for NoOrders.
	ParticipationRate *float64 `fix:"849"`
	//Designation is a non-required field for NoOrders.
	Designation *string `fix:"494"`
}

func (m *NoOrders) SetClOrdID(v string)                        { m.ClOrdID = v }
func (m *NoOrders) SetSecondaryClOrdID(v string)               { m.SecondaryClOrdID = &v }
func (m *NoOrders) SetListSeqNo(v int)                         { m.ListSeqNo = v }
func (m *NoOrders) SetClOrdLinkID(v string)                    { m.ClOrdLinkID = &v }
func (m *NoOrders) SetSettlInstMode(v string)                  { m.SettlInstMode = &v }
func (m *NoOrders) SetTradeOriginationDate(v string)           { m.TradeOriginationDate = &v }
func (m *NoOrders) SetTradeDate(v string)                      { m.TradeDate = &v }
func (m *NoOrders) SetAccount(v string)                        { m.Account = &v }
func (m *NoOrders) SetAcctIDSource(v int)                      { m.AcctIDSource = &v }
func (m *NoOrders) SetAccountType(v int)                       { m.AccountType = &v }
func (m *NoOrders) SetDayBookingInst(v string)                 { m.DayBookingInst = &v }
func (m *NoOrders) SetBookingUnit(v string)                    { m.BookingUnit = &v }
func (m *NoOrders) SetAllocID(v string)                        { m.AllocID = &v }
func (m *NoOrders) SetPreallocMethod(v string)                 { m.PreallocMethod = &v }
func (m *NoOrders) SetNoAllocs(v []NoAllocs)                   { m.NoAllocs = v }
func (m *NoOrders) SetSettlType(v string)                      { m.SettlType = &v }
func (m *NoOrders) SetSettlDate(v string)                      { m.SettlDate = &v }
func (m *NoOrders) SetCashMargin(v string)                     { m.CashMargin = &v }
func (m *NoOrders) SetClearingFeeIndicator(v string)           { m.ClearingFeeIndicator = &v }
func (m *NoOrders) SetHandlInst(v string)                      { m.HandlInst = &v }
func (m *NoOrders) SetExecInst(v string)                       { m.ExecInst = &v }
func (m *NoOrders) SetMinQty(v float64)                        { m.MinQty = &v }
func (m *NoOrders) SetMaxFloor(v float64)                      { m.MaxFloor = &v }
func (m *NoOrders) SetExDestination(v string)                  { m.ExDestination = &v }
func (m *NoOrders) SetNoTradingSessions(v []NoTradingSessions) { m.NoTradingSessions = v }
func (m *NoOrders) SetProcessCode(v string)                    { m.ProcessCode = &v }
func (m *NoOrders) SetNoUnderlyings(v []NoUnderlyings)         { m.NoUnderlyings = v }
func (m *NoOrders) SetPrevClosePx(v float64)                   { m.PrevClosePx = &v }
func (m *NoOrders) SetSide(v string)                           { m.Side = v }
func (m *NoOrders) SetSideValueInd(v int)                      { m.SideValueInd = &v }
func (m *NoOrders) SetLocateReqd(v bool)                       { m.LocateReqd = &v }
func (m *NoOrders) SetTransactTime(v time.Time)                { m.TransactTime = &v }
func (m *NoOrders) SetQtyType(v int)                           { m.QtyType = &v }
func (m *NoOrders) SetOrdType(v string)                        { m.OrdType = &v }
func (m *NoOrders) SetPriceType(v int)                         { m.PriceType = &v }
func (m *NoOrders) SetPrice(v float64)                         { m.Price = &v }
func (m *NoOrders) SetStopPx(v float64)                        { m.StopPx = &v }
func (m *NoOrders) SetCurrency(v string)                       { m.Currency = &v }
func (m *NoOrders) SetComplianceID(v string)                   { m.ComplianceID = &v }
func (m *NoOrders) SetSolicitedFlag(v bool)                    { m.SolicitedFlag = &v }
func (m *NoOrders) SetIOIID(v string)                          { m.IOIID = &v }
func (m *NoOrders) SetQuoteID(v string)                        { m.QuoteID = &v }
func (m *NoOrders) SetTimeInForce(v string)                    { m.TimeInForce = &v }
func (m *NoOrders) SetEffectiveTime(v time.Time)               { m.EffectiveTime = &v }
func (m *NoOrders) SetExpireDate(v string)                     { m.ExpireDate = &v }
func (m *NoOrders) SetExpireTime(v time.Time)                  { m.ExpireTime = &v }
func (m *NoOrders) SetGTBookingInst(v int)                     { m.GTBookingInst = &v }
func (m *NoOrders) SetOrderCapacity(v string)                  { m.OrderCapacity = &v }
func (m *NoOrders) SetOrderRestrictions(v string)              { m.OrderRestrictions = &v }
func (m *NoOrders) SetCustOrderCapacity(v int)                 { m.CustOrderCapacity = &v }
func (m *NoOrders) SetForexReq(v bool)                         { m.ForexReq = &v }
func (m *NoOrders) SetSettlCurrency(v string)                  { m.SettlCurrency = &v }
func (m *NoOrders) SetBookingType(v int)                       { m.BookingType = &v }
func (m *NoOrders) SetText(v string)                           { m.Text = &v }
func (m *NoOrders) SetEncodedTextLen(v int)                    { m.EncodedTextLen = &v }
func (m *NoOrders) SetEncodedText(v string)                    { m.EncodedText = &v }
func (m *NoOrders) SetSettlDate2(v string)                     { m.SettlDate2 = &v }
func (m *NoOrders) SetOrderQty2(v float64)                     { m.OrderQty2 = &v }
func (m *NoOrders) SetPrice2(v float64)                        { m.Price2 = &v }
func (m *NoOrders) SetPositionEffect(v string)                 { m.PositionEffect = &v }
func (m *NoOrders) SetCoveredOrUncovered(v int)                { m.CoveredOrUncovered = &v }
func (m *NoOrders) SetMaxShow(v float64)                       { m.MaxShow = &v }
func (m *NoOrders) SetTargetStrategy(v int)                    { m.TargetStrategy = &v }
func (m *NoOrders) SetTargetStrategyParameters(v string)       { m.TargetStrategyParameters = &v }
func (m *NoOrders) SetParticipationRate(v float64)             { m.ParticipationRate = &v }
func (m *NoOrders) SetDesignation(v string)                    { m.Designation = &v }

//NoAllocs is a repeating group in NoOrders
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for NoAllocs.
	AllocAcctIDSource *int `fix:"661"`
	//AllocSettlCurrency is a non-required field for NoAllocs.
	AllocSettlCurrency *string `fix:"736"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//NestedParties Component
	nestedparties.NestedParties
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
}

func (m *NoAllocs) SetAllocAccount(v string)       { m.AllocAccount = &v }
func (m *NoAllocs) SetAllocAcctIDSource(v int)     { m.AllocAcctIDSource = &v }
func (m *NoAllocs) SetAllocSettlCurrency(v string) { m.AllocSettlCurrency = &v }
func (m *NoAllocs) SetIndividualAllocID(v string)  { m.IndividualAllocID = &v }
func (m *NoAllocs) SetAllocQty(v float64)          { m.AllocQty = &v }

//NoTradingSessions is a repeating group in NoOrders
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

func (m *NoTradingSessions) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoTradingSessions) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }

//NoUnderlyings is a repeating group in NoOrders
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
}

//Message is a NewOrderList FIX Message
type Message struct {
	FIXMsgType string `fix:"E"`
	fix44.Header
	//ListID is a required field for NewOrderList.
	ListID string `fix:"66"`
	//BidID is a non-required field for NewOrderList.
	BidID *string `fix:"390"`
	//ClientBidID is a non-required field for NewOrderList.
	ClientBidID *string `fix:"391"`
	//ProgRptReqs is a non-required field for NewOrderList.
	ProgRptReqs *int `fix:"414"`
	//BidType is a required field for NewOrderList.
	BidType int `fix:"394"`
	//ProgPeriodInterval is a non-required field for NewOrderList.
	ProgPeriodInterval *int `fix:"415"`
	//CancellationRights is a non-required field for NewOrderList.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for NewOrderList.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for NewOrderList.
	RegistID *string `fix:"513"`
	//ListExecInstType is a non-required field for NewOrderList.
	ListExecInstType *string `fix:"433"`
	//ListExecInst is a non-required field for NewOrderList.
	ListExecInst *string `fix:"69"`
	//EncodedListExecInstLen is a non-required field for NewOrderList.
	EncodedListExecInstLen *int `fix:"352"`
	//EncodedListExecInst is a non-required field for NewOrderList.
	EncodedListExecInst *string `fix:"353"`
	//AllowableOneSidednessPct is a non-required field for NewOrderList.
	AllowableOneSidednessPct *float64 `fix:"765"`
	//AllowableOneSidednessValue is a non-required field for NewOrderList.
	AllowableOneSidednessValue *float64 `fix:"766"`
	//AllowableOneSidednessCurr is a non-required field for NewOrderList.
	AllowableOneSidednessCurr *string `fix:"767"`
	//TotNoOrders is a required field for NewOrderList.
	TotNoOrders int `fix:"68"`
	//LastFragment is a non-required field for NewOrderList.
	LastFragment *bool `fix:"893"`
	//NoOrders is a required field for NewOrderList.
	NoOrders []NoOrders `fix:"73"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetListID(v string)                      { m.ListID = v }
func (m *Message) SetBidID(v string)                       { m.BidID = &v }
func (m *Message) SetClientBidID(v string)                 { m.ClientBidID = &v }
func (m *Message) SetProgRptReqs(v int)                    { m.ProgRptReqs = &v }
func (m *Message) SetBidType(v int)                        { m.BidType = v }
func (m *Message) SetProgPeriodInterval(v int)             { m.ProgPeriodInterval = &v }
func (m *Message) SetCancellationRights(v string)          { m.CancellationRights = &v }
func (m *Message) SetMoneyLaunderingStatus(v string)       { m.MoneyLaunderingStatus = &v }
func (m *Message) SetRegistID(v string)                    { m.RegistID = &v }
func (m *Message) SetListExecInstType(v string)            { m.ListExecInstType = &v }
func (m *Message) SetListExecInst(v string)                { m.ListExecInst = &v }
func (m *Message) SetEncodedListExecInstLen(v int)         { m.EncodedListExecInstLen = &v }
func (m *Message) SetEncodedListExecInst(v string)         { m.EncodedListExecInst = &v }
func (m *Message) SetAllowableOneSidednessPct(v float64)   { m.AllowableOneSidednessPct = &v }
func (m *Message) SetAllowableOneSidednessValue(v float64) { m.AllowableOneSidednessValue = &v }
func (m *Message) SetAllowableOneSidednessCurr(v string)   { m.AllowableOneSidednessCurr = &v }
func (m *Message) SetTotNoOrders(v int)                    { m.TotNoOrders = v }
func (m *Message) SetLastFragment(v bool)                  { m.LastFragment = &v }
func (m *Message) SetNoOrders(v []NoOrders)                { m.NoOrders = v }

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
	return enum.BeginStringFIX44, "E", r
}
