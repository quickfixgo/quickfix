//Package neworderlist msg type = E.
package neworderlist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/commissiondata"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/nestedparties"
	"github.com/quickfixgo/quickfix/fix43/orderqtydata"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"github.com/quickfixgo/quickfix/fix43/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix43/stipulations"
	"github.com/quickfixgo/quickfix/fix43/yielddata"
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
	Parties parties.Component
	//TradeOriginationDate is a non-required field for NoOrders.
	TradeOriginationDate *string `fix:"229"`
	//Account is a non-required field for NoOrders.
	Account *string `fix:"1"`
	//AccountType is a non-required field for NoOrders.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for NoOrders.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for NoOrders.
	BookingUnit *string `fix:"590"`
	//PreallocMethod is a non-required field for NoOrders.
	PreallocMethod *string `fix:"591"`
	//NoAllocs is a non-required field for NoOrders.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	//SettlmntTyp is a non-required field for NoOrders.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for NoOrders.
	FutSettDate *string `fix:"64"`
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
	Instrument instrument.Component
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
	Stipulations stipulations.Component
	//QuantityType is a non-required field for NoOrders.
	QuantityType *int `fix:"465"`
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//OrdType is a non-required field for NoOrders.
	OrdType *string `fix:"40"`
	//PriceType is a non-required field for NoOrders.
	PriceType *int `fix:"423"`
	//Price is a non-required field for NoOrders.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NoOrders.
	StopPx *float64 `fix:"99"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//YieldData Component
	YieldData yielddata.Component
	//Currency is a non-required field for NoOrders.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for NoOrders.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for NoOrders.
	SolicitedFlag *bool `fix:"377"`
	//IOIid is a non-required field for NoOrders.
	IOIid *string `fix:"23"`
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
	CommissionData commissiondata.Component
	//OrderCapacity is a non-required field for NoOrders.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for NoOrders.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for NoOrders.
	CustOrderCapacity *int `fix:"582"`
	//Rule80A is a non-required field for NoOrders.
	Rule80A *string `fix:"47"`
	//ForexReq is a non-required field for NoOrders.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for NoOrders.
	SettlCurrency *string `fix:"120"`
	//Text is a non-required field for NoOrders.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoOrders.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoOrders.
	EncodedText *string `fix:"355"`
	//FutSettDate2 is a non-required field for NoOrders.
	FutSettDate2 *string `fix:"193"`
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
	//PegDifference is a non-required field for NoOrders.
	PegDifference *float64 `fix:"211"`
	//DiscretionInst is a non-required field for NoOrders.
	DiscretionInst *string `fix:"388"`
	//DiscretionOffset is a non-required field for NoOrders.
	DiscretionOffset *float64 `fix:"389"`
	//Designation is a non-required field for NoOrders.
	Designation *string `fix:"494"`
	//AccruedInterestRate is a non-required field for NoOrders.
	AccruedInterestRate *float64 `fix:"158"`
	//AccruedInterestAmt is a non-required field for NoOrders.
	AccruedInterestAmt *float64 `fix:"159"`
	//NetMoney is a non-required field for NoOrders.
	NetMoney *float64 `fix:"118"`
}

func (m *NoOrders) SetClOrdID(v string)                        { m.ClOrdID = v }
func (m *NoOrders) SetSecondaryClOrdID(v string)               { m.SecondaryClOrdID = &v }
func (m *NoOrders) SetListSeqNo(v int)                         { m.ListSeqNo = v }
func (m *NoOrders) SetClOrdLinkID(v string)                    { m.ClOrdLinkID = &v }
func (m *NoOrders) SetSettlInstMode(v string)                  { m.SettlInstMode = &v }
func (m *NoOrders) SetTradeOriginationDate(v string)           { m.TradeOriginationDate = &v }
func (m *NoOrders) SetAccount(v string)                        { m.Account = &v }
func (m *NoOrders) SetAccountType(v int)                       { m.AccountType = &v }
func (m *NoOrders) SetDayBookingInst(v string)                 { m.DayBookingInst = &v }
func (m *NoOrders) SetBookingUnit(v string)                    { m.BookingUnit = &v }
func (m *NoOrders) SetPreallocMethod(v string)                 { m.PreallocMethod = &v }
func (m *NoOrders) SetNoAllocs(v []NoAllocs)                   { m.NoAllocs = v }
func (m *NoOrders) SetSettlmntTyp(v string)                    { m.SettlmntTyp = &v }
func (m *NoOrders) SetFutSettDate(v string)                    { m.FutSettDate = &v }
func (m *NoOrders) SetCashMargin(v string)                     { m.CashMargin = &v }
func (m *NoOrders) SetClearingFeeIndicator(v string)           { m.ClearingFeeIndicator = &v }
func (m *NoOrders) SetHandlInst(v string)                      { m.HandlInst = &v }
func (m *NoOrders) SetExecInst(v string)                       { m.ExecInst = &v }
func (m *NoOrders) SetMinQty(v float64)                        { m.MinQty = &v }
func (m *NoOrders) SetMaxFloor(v float64)                      { m.MaxFloor = &v }
func (m *NoOrders) SetExDestination(v string)                  { m.ExDestination = &v }
func (m *NoOrders) SetNoTradingSessions(v []NoTradingSessions) { m.NoTradingSessions = v }
func (m *NoOrders) SetProcessCode(v string)                    { m.ProcessCode = &v }
func (m *NoOrders) SetPrevClosePx(v float64)                   { m.PrevClosePx = &v }
func (m *NoOrders) SetSide(v string)                           { m.Side = v }
func (m *NoOrders) SetSideValueInd(v int)                      { m.SideValueInd = &v }
func (m *NoOrders) SetLocateReqd(v bool)                       { m.LocateReqd = &v }
func (m *NoOrders) SetTransactTime(v time.Time)                { m.TransactTime = &v }
func (m *NoOrders) SetQuantityType(v int)                      { m.QuantityType = &v }
func (m *NoOrders) SetOrdType(v string)                        { m.OrdType = &v }
func (m *NoOrders) SetPriceType(v int)                         { m.PriceType = &v }
func (m *NoOrders) SetPrice(v float64)                         { m.Price = &v }
func (m *NoOrders) SetStopPx(v float64)                        { m.StopPx = &v }
func (m *NoOrders) SetCurrency(v string)                       { m.Currency = &v }
func (m *NoOrders) SetComplianceID(v string)                   { m.ComplianceID = &v }
func (m *NoOrders) SetSolicitedFlag(v bool)                    { m.SolicitedFlag = &v }
func (m *NoOrders) SetIOIid(v string)                          { m.IOIid = &v }
func (m *NoOrders) SetQuoteID(v string)                        { m.QuoteID = &v }
func (m *NoOrders) SetTimeInForce(v string)                    { m.TimeInForce = &v }
func (m *NoOrders) SetEffectiveTime(v time.Time)               { m.EffectiveTime = &v }
func (m *NoOrders) SetExpireDate(v string)                     { m.ExpireDate = &v }
func (m *NoOrders) SetExpireTime(v time.Time)                  { m.ExpireTime = &v }
func (m *NoOrders) SetGTBookingInst(v int)                     { m.GTBookingInst = &v }
func (m *NoOrders) SetOrderCapacity(v string)                  { m.OrderCapacity = &v }
func (m *NoOrders) SetOrderRestrictions(v string)              { m.OrderRestrictions = &v }
func (m *NoOrders) SetCustOrderCapacity(v int)                 { m.CustOrderCapacity = &v }
func (m *NoOrders) SetRule80A(v string)                        { m.Rule80A = &v }
func (m *NoOrders) SetForexReq(v bool)                         { m.ForexReq = &v }
func (m *NoOrders) SetSettlCurrency(v string)                  { m.SettlCurrency = &v }
func (m *NoOrders) SetText(v string)                           { m.Text = &v }
func (m *NoOrders) SetEncodedTextLen(v int)                    { m.EncodedTextLen = &v }
func (m *NoOrders) SetEncodedText(v string)                    { m.EncodedText = &v }
func (m *NoOrders) SetFutSettDate2(v string)                   { m.FutSettDate2 = &v }
func (m *NoOrders) SetOrderQty2(v float64)                     { m.OrderQty2 = &v }
func (m *NoOrders) SetPrice2(v float64)                        { m.Price2 = &v }
func (m *NoOrders) SetPositionEffect(v string)                 { m.PositionEffect = &v }
func (m *NoOrders) SetCoveredOrUncovered(v int)                { m.CoveredOrUncovered = &v }
func (m *NoOrders) SetMaxShow(v float64)                       { m.MaxShow = &v }
func (m *NoOrders) SetPegDifference(v float64)                 { m.PegDifference = &v }
func (m *NoOrders) SetDiscretionInst(v string)                 { m.DiscretionInst = &v }
func (m *NoOrders) SetDiscretionOffset(v float64)              { m.DiscretionOffset = &v }
func (m *NoOrders) SetDesignation(v string)                    { m.Designation = &v }
func (m *NoOrders) SetAccruedInterestRate(v float64)           { m.AccruedInterestRate = &v }
func (m *NoOrders) SetAccruedInterestAmt(v float64)            { m.AccruedInterestAmt = &v }
func (m *NoOrders) SetNetMoney(v float64)                      { m.NetMoney = &v }

//NoAllocs is a repeating group in NoOrders
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//NestedParties Component
	NestedParties nestedparties.Component
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
}

func (m *NoAllocs) SetAllocAccount(v string)      { m.AllocAccount = &v }
func (m *NoAllocs) SetIndividualAllocID(v string) { m.IndividualAllocID = &v }
func (m *NoAllocs) SetAllocQty(v float64)         { m.AllocQty = &v }

//NoTradingSessions is a repeating group in NoOrders
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

func (m *NoTradingSessions) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoTradingSessions) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }

//Message is a NewOrderList FIX Message
type Message struct {
	FIXMsgType string `fix:"E"`
	Header     fix43.Header
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
	//TotNoOrders is a required field for NewOrderList.
	TotNoOrders int `fix:"68"`
	//NoOrders is a required field for NewOrderList.
	NoOrders []NoOrders `fix:"73"`
	Trailer  fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetListID(v string)                { m.ListID = v }
func (m *Message) SetBidID(v string)                 { m.BidID = &v }
func (m *Message) SetClientBidID(v string)           { m.ClientBidID = &v }
func (m *Message) SetProgRptReqs(v int)              { m.ProgRptReqs = &v }
func (m *Message) SetBidType(v int)                  { m.BidType = v }
func (m *Message) SetProgPeriodInterval(v int)       { m.ProgPeriodInterval = &v }
func (m *Message) SetCancellationRights(v string)    { m.CancellationRights = &v }
func (m *Message) SetMoneyLaunderingStatus(v string) { m.MoneyLaunderingStatus = &v }
func (m *Message) SetRegistID(v string)              { m.RegistID = &v }
func (m *Message) SetListExecInstType(v string)      { m.ListExecInstType = &v }
func (m *Message) SetListExecInst(v string)          { m.ListExecInst = &v }
func (m *Message) SetEncodedListExecInstLen(v int)   { m.EncodedListExecInstLen = &v }
func (m *Message) SetEncodedListExecInst(v string)   { m.EncodedListExecInst = &v }
func (m *Message) SetTotNoOrders(v int)              { m.TotNoOrders = v }
func (m *Message) SetNoOrders(v []NoOrders)          { m.NoOrders = v }

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
	return enum.BeginStringFIX43, "E", r
}
