package neworderlist

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//NewOrderList is the fix43 NewOrderList type, MsgType = E
type NewOrderList struct {
	fix43.Header
	quickfix.Body
	fix43.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a NewOrderList from a quickfix.Message instance
func FromMessage(m quickfix.Message) NewOrderList {
	return NewOrderList{
		Header:      fix43.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix43.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m NewOrderList) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a NewOrderList initialized with the required fields for NewOrderList
func New(listid field.ListIDField, bidtype field.BidTypeField, totnoorders field.TotNoOrdersField) (m NewOrderList) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("E"))
	m.Set(listid)
	m.Set(bidtype)
	m.Set(totnoorders)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg NewOrderList, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "E", r
}

//SetListID sets ListID, Tag 66
func (m NewOrderList) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetTotNoOrders sets TotNoOrders, Tag 68
func (m NewOrderList) SetTotNoOrders(v int) {
	m.Set(field.NewTotNoOrders(v))
}

//SetListExecInst sets ListExecInst, Tag 69
func (m NewOrderList) SetListExecInst(v string) {
	m.Set(field.NewListExecInst(v))
}

//SetNoOrders sets NoOrders, Tag 73
func (m NewOrderList) SetNoOrders(f NoOrdersRepeatingGroup) {
	m.SetGroup(f)
}

//SetEncodedListExecInstLen sets EncodedListExecInstLen, Tag 352
func (m NewOrderList) SetEncodedListExecInstLen(v int) {
	m.Set(field.NewEncodedListExecInstLen(v))
}

//SetEncodedListExecInst sets EncodedListExecInst, Tag 353
func (m NewOrderList) SetEncodedListExecInst(v string) {
	m.Set(field.NewEncodedListExecInst(v))
}

//SetBidID sets BidID, Tag 390
func (m NewOrderList) SetBidID(v string) {
	m.Set(field.NewBidID(v))
}

//SetClientBidID sets ClientBidID, Tag 391
func (m NewOrderList) SetClientBidID(v string) {
	m.Set(field.NewClientBidID(v))
}

//SetBidType sets BidType, Tag 394
func (m NewOrderList) SetBidType(v int) {
	m.Set(field.NewBidType(v))
}

//SetProgRptReqs sets ProgRptReqs, Tag 414
func (m NewOrderList) SetProgRptReqs(v int) {
	m.Set(field.NewProgRptReqs(v))
}

//SetProgPeriodInterval sets ProgPeriodInterval, Tag 415
func (m NewOrderList) SetProgPeriodInterval(v int) {
	m.Set(field.NewProgPeriodInterval(v))
}

//SetListExecInstType sets ListExecInstType, Tag 433
func (m NewOrderList) SetListExecInstType(v string) {
	m.Set(field.NewListExecInstType(v))
}

//SetCancellationRights sets CancellationRights, Tag 480
func (m NewOrderList) SetCancellationRights(v string) {
	m.Set(field.NewCancellationRights(v))
}

//SetMoneyLaunderingStatus sets MoneyLaunderingStatus, Tag 481
func (m NewOrderList) SetMoneyLaunderingStatus(v string) {
	m.Set(field.NewMoneyLaunderingStatus(v))
}

//SetRegistID sets RegistID, Tag 513
func (m NewOrderList) SetRegistID(v string) {
	m.Set(field.NewRegistID(v))
}

//GetListID gets ListID, Tag 66
func (m NewOrderList) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTotNoOrders gets TotNoOrders, Tag 68
func (m NewOrderList) GetTotNoOrders() (f field.TotNoOrdersField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListExecInst gets ListExecInst, Tag 69
func (m NewOrderList) GetListExecInst() (f field.ListExecInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoOrders gets NoOrders, Tag 73
func (m NewOrderList) GetNoOrders() (NoOrdersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoOrdersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetEncodedListExecInstLen gets EncodedListExecInstLen, Tag 352
func (m NewOrderList) GetEncodedListExecInstLen() (f field.EncodedListExecInstLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedListExecInst gets EncodedListExecInst, Tag 353
func (m NewOrderList) GetEncodedListExecInst() (f field.EncodedListExecInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidID gets BidID, Tag 390
func (m NewOrderList) GetBidID() (f field.BidIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClientBidID gets ClientBidID, Tag 391
func (m NewOrderList) GetClientBidID() (f field.ClientBidIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidType gets BidType, Tag 394
func (m NewOrderList) GetBidType() (f field.BidTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProgRptReqs gets ProgRptReqs, Tag 414
func (m NewOrderList) GetProgRptReqs() (f field.ProgRptReqsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProgPeriodInterval gets ProgPeriodInterval, Tag 415
func (m NewOrderList) GetProgPeriodInterval() (f field.ProgPeriodIntervalField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListExecInstType gets ListExecInstType, Tag 433
func (m NewOrderList) GetListExecInstType() (f field.ListExecInstTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCancellationRights gets CancellationRights, Tag 480
func (m NewOrderList) GetCancellationRights() (f field.CancellationRightsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMoneyLaunderingStatus gets MoneyLaunderingStatus, Tag 481
func (m NewOrderList) GetMoneyLaunderingStatus() (f field.MoneyLaunderingStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRegistID gets RegistID, Tag 513
func (m NewOrderList) GetRegistID() (f field.RegistIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasListID returns true if ListID is present, Tag 66
func (m NewOrderList) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasTotNoOrders returns true if TotNoOrders is present, Tag 68
func (m NewOrderList) HasTotNoOrders() bool {
	return m.Has(tag.TotNoOrders)
}

//HasListExecInst returns true if ListExecInst is present, Tag 69
func (m NewOrderList) HasListExecInst() bool {
	return m.Has(tag.ListExecInst)
}

//HasNoOrders returns true if NoOrders is present, Tag 73
func (m NewOrderList) HasNoOrders() bool {
	return m.Has(tag.NoOrders)
}

//HasEncodedListExecInstLen returns true if EncodedListExecInstLen is present, Tag 352
func (m NewOrderList) HasEncodedListExecInstLen() bool {
	return m.Has(tag.EncodedListExecInstLen)
}

//HasEncodedListExecInst returns true if EncodedListExecInst is present, Tag 353
func (m NewOrderList) HasEncodedListExecInst() bool {
	return m.Has(tag.EncodedListExecInst)
}

//HasBidID returns true if BidID is present, Tag 390
func (m NewOrderList) HasBidID() bool {
	return m.Has(tag.BidID)
}

//HasClientBidID returns true if ClientBidID is present, Tag 391
func (m NewOrderList) HasClientBidID() bool {
	return m.Has(tag.ClientBidID)
}

//HasBidType returns true if BidType is present, Tag 394
func (m NewOrderList) HasBidType() bool {
	return m.Has(tag.BidType)
}

//HasProgRptReqs returns true if ProgRptReqs is present, Tag 414
func (m NewOrderList) HasProgRptReqs() bool {
	return m.Has(tag.ProgRptReqs)
}

//HasProgPeriodInterval returns true if ProgPeriodInterval is present, Tag 415
func (m NewOrderList) HasProgPeriodInterval() bool {
	return m.Has(tag.ProgPeriodInterval)
}

//HasListExecInstType returns true if ListExecInstType is present, Tag 433
func (m NewOrderList) HasListExecInstType() bool {
	return m.Has(tag.ListExecInstType)
}

//HasCancellationRights returns true if CancellationRights is present, Tag 480
func (m NewOrderList) HasCancellationRights() bool {
	return m.Has(tag.CancellationRights)
}

//HasMoneyLaunderingStatus returns true if MoneyLaunderingStatus is present, Tag 481
func (m NewOrderList) HasMoneyLaunderingStatus() bool {
	return m.Has(tag.MoneyLaunderingStatus)
}

//HasRegistID returns true if RegistID is present, Tag 513
func (m NewOrderList) HasRegistID() bool {
	return m.Has(tag.RegistID)
}

//NoOrders is a repeating group element, Tag 73
type NoOrders struct {
	quickfix.Group
}

//SetClOrdID sets ClOrdID, Tag 11
func (m NoOrders) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526
func (m NoOrders) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

//SetListSeqNo sets ListSeqNo, Tag 67
func (m NoOrders) SetListSeqNo(v int) {
	m.Set(field.NewListSeqNo(v))
}

//SetClOrdLinkID sets ClOrdLinkID, Tag 583
func (m NoOrders) SetClOrdLinkID(v string) {
	m.Set(field.NewClOrdLinkID(v))
}

//SetSettlInstMode sets SettlInstMode, Tag 160
func (m NoOrders) SetSettlInstMode(v string) {
	m.Set(field.NewSettlInstMode(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m NoOrders) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetTradeOriginationDate sets TradeOriginationDate, Tag 229
func (m NoOrders) SetTradeOriginationDate(v string) {
	m.Set(field.NewTradeOriginationDate(v))
}

//SetAccount sets Account, Tag 1
func (m NoOrders) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetAccountType sets AccountType, Tag 581
func (m NoOrders) SetAccountType(v int) {
	m.Set(field.NewAccountType(v))
}

//SetDayBookingInst sets DayBookingInst, Tag 589
func (m NoOrders) SetDayBookingInst(v string) {
	m.Set(field.NewDayBookingInst(v))
}

//SetBookingUnit sets BookingUnit, Tag 590
func (m NoOrders) SetBookingUnit(v string) {
	m.Set(field.NewBookingUnit(v))
}

//SetPreallocMethod sets PreallocMethod, Tag 591
func (m NoOrders) SetPreallocMethod(v string) {
	m.Set(field.NewPreallocMethod(v))
}

//SetNoAllocs sets NoAllocs, Tag 78
func (m NoOrders) SetNoAllocs(f NoAllocsRepeatingGroup) {
	m.SetGroup(f)
}

//SetSettlmntTyp sets SettlmntTyp, Tag 63
func (m NoOrders) SetSettlmntTyp(v string) {
	m.Set(field.NewSettlmntTyp(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m NoOrders) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetCashMargin sets CashMargin, Tag 544
func (m NoOrders) SetCashMargin(v string) {
	m.Set(field.NewCashMargin(v))
}

//SetClearingFeeIndicator sets ClearingFeeIndicator, Tag 635
func (m NoOrders) SetClearingFeeIndicator(v string) {
	m.Set(field.NewClearingFeeIndicator(v))
}

//SetHandlInst sets HandlInst, Tag 21
func (m NoOrders) SetHandlInst(v string) {
	m.Set(field.NewHandlInst(v))
}

//SetExecInst sets ExecInst, Tag 18
func (m NoOrders) SetExecInst(v string) {
	m.Set(field.NewExecInst(v))
}

//SetMinQty sets MinQty, Tag 110
func (m NoOrders) SetMinQty(v float64) {
	m.Set(field.NewMinQty(v))
}

//SetMaxFloor sets MaxFloor, Tag 111
func (m NoOrders) SetMaxFloor(v float64) {
	m.Set(field.NewMaxFloor(v))
}

//SetExDestination sets ExDestination, Tag 100
func (m NoOrders) SetExDestination(v string) {
	m.Set(field.NewExDestination(v))
}

//SetNoTradingSessions sets NoTradingSessions, Tag 386
func (m NoOrders) SetNoTradingSessions(f NoTradingSessionsRepeatingGroup) {
	m.SetGroup(f)
}

//SetProcessCode sets ProcessCode, Tag 81
func (m NoOrders) SetProcessCode(v string) {
	m.Set(field.NewProcessCode(v))
}

//SetSymbol sets Symbol, Tag 55
func (m NoOrders) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m NoOrders) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m NoOrders) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m NoOrders) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m NoOrders) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m NoOrders) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m NoOrders) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m NoOrders) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m NoOrders) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m NoOrders) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m NoOrders) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m NoOrders) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m NoOrders) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m NoOrders) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m NoOrders) SetRepurchaseRate(v float64) {
	m.Set(field.NewRepurchaseRate(v))
}

//SetFactor sets Factor, Tag 228
func (m NoOrders) SetFactor(v float64) {
	m.Set(field.NewFactor(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m NoOrders) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m NoOrders) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m NoOrders) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m NoOrders) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m NoOrders) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m NoOrders) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m NoOrders) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m NoOrders) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m NoOrders) SetContractMultiplier(v float64) {
	m.Set(field.NewContractMultiplier(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m NoOrders) SetCouponRate(v float64) {
	m.Set(field.NewCouponRate(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m NoOrders) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetIssuer sets Issuer, Tag 106
func (m NoOrders) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m NoOrders) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m NoOrders) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m NoOrders) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m NoOrders) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m NoOrders) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetPrevClosePx sets PrevClosePx, Tag 140
func (m NoOrders) SetPrevClosePx(v float64) {
	m.Set(field.NewPrevClosePx(v))
}

//SetSide sets Side, Tag 54
func (m NoOrders) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSideValueInd sets SideValueInd, Tag 401
func (m NoOrders) SetSideValueInd(v int) {
	m.Set(field.NewSideValueInd(v))
}

//SetLocateReqd sets LocateReqd, Tag 114
func (m NoOrders) SetLocateReqd(v bool) {
	m.Set(field.NewLocateReqd(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m NoOrders) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetNoStipulations sets NoStipulations, Tag 232
func (m NoOrders) SetNoStipulations(f NoStipulationsRepeatingGroup) {
	m.SetGroup(f)
}

//SetQuantityType sets QuantityType, Tag 465
func (m NoOrders) SetQuantityType(v int) {
	m.Set(field.NewQuantityType(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m NoOrders) SetOrderQty(v float64) {
	m.Set(field.NewOrderQty(v))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m NoOrders) SetCashOrderQty(v float64) {
	m.Set(field.NewCashOrderQty(v))
}

//SetOrderPercent sets OrderPercent, Tag 516
func (m NoOrders) SetOrderPercent(v float64) {
	m.Set(field.NewOrderPercent(v))
}

//SetRoundingDirection sets RoundingDirection, Tag 468
func (m NoOrders) SetRoundingDirection(v string) {
	m.Set(field.NewRoundingDirection(v))
}

//SetRoundingModulus sets RoundingModulus, Tag 469
func (m NoOrders) SetRoundingModulus(v float64) {
	m.Set(field.NewRoundingModulus(v))
}

//SetOrdType sets OrdType, Tag 40
func (m NoOrders) SetOrdType(v string) {
	m.Set(field.NewOrdType(v))
}

//SetPriceType sets PriceType, Tag 423
func (m NoOrders) SetPriceType(v int) {
	m.Set(field.NewPriceType(v))
}

//SetPrice sets Price, Tag 44
func (m NoOrders) SetPrice(v float64) {
	m.Set(field.NewPrice(v))
}

//SetStopPx sets StopPx, Tag 99
func (m NoOrders) SetStopPx(v float64) {
	m.Set(field.NewStopPx(v))
}

//SetSpread sets Spread, Tag 218
func (m NoOrders) SetSpread(v float64) {
	m.Set(field.NewSpread(v))
}

//SetBenchmarkCurveCurrency sets BenchmarkCurveCurrency, Tag 220
func (m NoOrders) SetBenchmarkCurveCurrency(v string) {
	m.Set(field.NewBenchmarkCurveCurrency(v))
}

//SetBenchmarkCurveName sets BenchmarkCurveName, Tag 221
func (m NoOrders) SetBenchmarkCurveName(v string) {
	m.Set(field.NewBenchmarkCurveName(v))
}

//SetBenchmarkCurvePoint sets BenchmarkCurvePoint, Tag 222
func (m NoOrders) SetBenchmarkCurvePoint(v string) {
	m.Set(field.NewBenchmarkCurvePoint(v))
}

//SetYieldType sets YieldType, Tag 235
func (m NoOrders) SetYieldType(v string) {
	m.Set(field.NewYieldType(v))
}

//SetYield sets Yield, Tag 236
func (m NoOrders) SetYield(v float64) {
	m.Set(field.NewYield(v))
}

//SetCurrency sets Currency, Tag 15
func (m NoOrders) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetComplianceID sets ComplianceID, Tag 376
func (m NoOrders) SetComplianceID(v string) {
	m.Set(field.NewComplianceID(v))
}

//SetSolicitedFlag sets SolicitedFlag, Tag 377
func (m NoOrders) SetSolicitedFlag(v bool) {
	m.Set(field.NewSolicitedFlag(v))
}

//SetIOIid sets IOIid, Tag 23
func (m NoOrders) SetIOIid(v string) {
	m.Set(field.NewIOIid(v))
}

//SetQuoteID sets QuoteID, Tag 117
func (m NoOrders) SetQuoteID(v string) {
	m.Set(field.NewQuoteID(v))
}

//SetTimeInForce sets TimeInForce, Tag 59
func (m NoOrders) SetTimeInForce(v string) {
	m.Set(field.NewTimeInForce(v))
}

//SetEffectiveTime sets EffectiveTime, Tag 168
func (m NoOrders) SetEffectiveTime(v time.Time) {
	m.Set(field.NewEffectiveTime(v))
}

//SetExpireDate sets ExpireDate, Tag 432
func (m NoOrders) SetExpireDate(v string) {
	m.Set(field.NewExpireDate(v))
}

//SetExpireTime sets ExpireTime, Tag 126
func (m NoOrders) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

//SetGTBookingInst sets GTBookingInst, Tag 427
func (m NoOrders) SetGTBookingInst(v int) {
	m.Set(field.NewGTBookingInst(v))
}

//SetCommission sets Commission, Tag 12
func (m NoOrders) SetCommission(v float64) {
	m.Set(field.NewCommission(v))
}

//SetCommType sets CommType, Tag 13
func (m NoOrders) SetCommType(v string) {
	m.Set(field.NewCommType(v))
}

//SetCommCurrency sets CommCurrency, Tag 479
func (m NoOrders) SetCommCurrency(v string) {
	m.Set(field.NewCommCurrency(v))
}

//SetFundRenewWaiv sets FundRenewWaiv, Tag 497
func (m NoOrders) SetFundRenewWaiv(v string) {
	m.Set(field.NewFundRenewWaiv(v))
}

//SetOrderCapacity sets OrderCapacity, Tag 528
func (m NoOrders) SetOrderCapacity(v string) {
	m.Set(field.NewOrderCapacity(v))
}

//SetOrderRestrictions sets OrderRestrictions, Tag 529
func (m NoOrders) SetOrderRestrictions(v string) {
	m.Set(field.NewOrderRestrictions(v))
}

//SetCustOrderCapacity sets CustOrderCapacity, Tag 582
func (m NoOrders) SetCustOrderCapacity(v int) {
	m.Set(field.NewCustOrderCapacity(v))
}

//SetRule80A sets Rule80A, Tag 47
func (m NoOrders) SetRule80A(v string) {
	m.Set(field.NewRule80A(v))
}

//SetForexReq sets ForexReq, Tag 121
func (m NoOrders) SetForexReq(v bool) {
	m.Set(field.NewForexReq(v))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m NoOrders) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetText sets Text, Tag 58
func (m NoOrders) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m NoOrders) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m NoOrders) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetFutSettDate2 sets FutSettDate2, Tag 193
func (m NoOrders) SetFutSettDate2(v string) {
	m.Set(field.NewFutSettDate2(v))
}

//SetOrderQty2 sets OrderQty2, Tag 192
func (m NoOrders) SetOrderQty2(v float64) {
	m.Set(field.NewOrderQty2(v))
}

//SetPrice2 sets Price2, Tag 640
func (m NoOrders) SetPrice2(v float64) {
	m.Set(field.NewPrice2(v))
}

//SetPositionEffect sets PositionEffect, Tag 77
func (m NoOrders) SetPositionEffect(v string) {
	m.Set(field.NewPositionEffect(v))
}

//SetCoveredOrUncovered sets CoveredOrUncovered, Tag 203
func (m NoOrders) SetCoveredOrUncovered(v int) {
	m.Set(field.NewCoveredOrUncovered(v))
}

//SetMaxShow sets MaxShow, Tag 210
func (m NoOrders) SetMaxShow(v float64) {
	m.Set(field.NewMaxShow(v))
}

//SetPegDifference sets PegDifference, Tag 211
func (m NoOrders) SetPegDifference(v float64) {
	m.Set(field.NewPegDifference(v))
}

//SetDiscretionInst sets DiscretionInst, Tag 388
func (m NoOrders) SetDiscretionInst(v string) {
	m.Set(field.NewDiscretionInst(v))
}

//SetDiscretionOffset sets DiscretionOffset, Tag 389
func (m NoOrders) SetDiscretionOffset(v float64) {
	m.Set(field.NewDiscretionOffset(v))
}

//SetDesignation sets Designation, Tag 494
func (m NoOrders) SetDesignation(v string) {
	m.Set(field.NewDesignation(v))
}

//SetAccruedInterestRate sets AccruedInterestRate, Tag 158
func (m NoOrders) SetAccruedInterestRate(v float64) {
	m.Set(field.NewAccruedInterestRate(v))
}

//SetAccruedInterestAmt sets AccruedInterestAmt, Tag 159
func (m NoOrders) SetAccruedInterestAmt(v float64) {
	m.Set(field.NewAccruedInterestAmt(v))
}

//SetNetMoney sets NetMoney, Tag 118
func (m NoOrders) SetNetMoney(v float64) {
	m.Set(field.NewNetMoney(v))
}

//GetClOrdID gets ClOrdID, Tag 11
func (m NoOrders) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526
func (m NoOrders) GetSecondaryClOrdID() (f field.SecondaryClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListSeqNo gets ListSeqNo, Tag 67
func (m NoOrders) GetListSeqNo() (f field.ListSeqNoField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdLinkID gets ClOrdLinkID, Tag 583
func (m NoOrders) GetClOrdLinkID() (f field.ClOrdLinkIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlInstMode gets SettlInstMode, Tag 160
func (m NoOrders) GetSettlInstMode() (f field.SettlInstModeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m NoOrders) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetTradeOriginationDate gets TradeOriginationDate, Tag 229
func (m NoOrders) GetTradeOriginationDate() (f field.TradeOriginationDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccount gets Account, Tag 1
func (m NoOrders) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccountType gets AccountType, Tag 581
func (m NoOrders) GetAccountType() (f field.AccountTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDayBookingInst gets DayBookingInst, Tag 589
func (m NoOrders) GetDayBookingInst() (f field.DayBookingInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBookingUnit gets BookingUnit, Tag 590
func (m NoOrders) GetBookingUnit() (f field.BookingUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPreallocMethod gets PreallocMethod, Tag 591
func (m NoOrders) GetPreallocMethod() (f field.PreallocMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoAllocs gets NoAllocs, Tag 78
func (m NoOrders) GetNoAllocs() (NoAllocsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAllocsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSettlmntTyp gets SettlmntTyp, Tag 63
func (m NoOrders) GetSettlmntTyp() (f field.SettlmntTypField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m NoOrders) GetFutSettDate() (f field.FutSettDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashMargin gets CashMargin, Tag 544
func (m NoOrders) GetCashMargin() (f field.CashMarginField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClearingFeeIndicator gets ClearingFeeIndicator, Tag 635
func (m NoOrders) GetClearingFeeIndicator() (f field.ClearingFeeIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetHandlInst gets HandlInst, Tag 21
func (m NoOrders) GetHandlInst() (f field.HandlInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecInst gets ExecInst, Tag 18
func (m NoOrders) GetExecInst() (f field.ExecInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinQty gets MinQty, Tag 110
func (m NoOrders) GetMinQty() (f field.MinQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxFloor gets MaxFloor, Tag 111
func (m NoOrders) GetMaxFloor() (f field.MaxFloorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExDestination gets ExDestination, Tag 100
func (m NoOrders) GetExDestination() (f field.ExDestinationField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoTradingSessions gets NoTradingSessions, Tag 386
func (m NoOrders) GetNoTradingSessions() (NoTradingSessionsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTradingSessionsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProcessCode gets ProcessCode, Tag 81
func (m NoOrders) GetProcessCode() (f field.ProcessCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m NoOrders) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m NoOrders) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m NoOrders) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m NoOrders) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m NoOrders) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m NoOrders) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m NoOrders) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m NoOrders) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m NoOrders) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m NoOrders) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m NoOrders) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m NoOrders) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m NoOrders) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m NoOrders) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m NoOrders) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m NoOrders) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m NoOrders) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m NoOrders) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m NoOrders) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m NoOrders) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m NoOrders) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m NoOrders) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m NoOrders) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m NoOrders) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m NoOrders) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m NoOrders) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m NoOrders) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m NoOrders) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m NoOrders) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m NoOrders) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m NoOrders) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m NoOrders) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m NoOrders) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrevClosePx gets PrevClosePx, Tag 140
func (m NoOrders) GetPrevClosePx() (f field.PrevClosePxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m NoOrders) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSideValueInd gets SideValueInd, Tag 401
func (m NoOrders) GetSideValueInd() (f field.SideValueIndField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocateReqd gets LocateReqd, Tag 114
func (m NoOrders) GetLocateReqd() (f field.LocateReqdField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m NoOrders) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoStipulations gets NoStipulations, Tag 232
func (m NoOrders) GetNoStipulations() (NoStipulationsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoStipulationsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetQuantityType gets QuantityType, Tag 465
func (m NoOrders) GetQuantityType() (f field.QuantityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m NoOrders) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m NoOrders) GetCashOrderQty() (f field.CashOrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderPercent gets OrderPercent, Tag 516
func (m NoOrders) GetOrderPercent() (f field.OrderPercentField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingDirection gets RoundingDirection, Tag 468
func (m NoOrders) GetRoundingDirection() (f field.RoundingDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingModulus gets RoundingModulus, Tag 469
func (m NoOrders) GetRoundingModulus() (f field.RoundingModulusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdType gets OrdType, Tag 40
func (m NoOrders) GetOrdType() (f field.OrdTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceType gets PriceType, Tag 423
func (m NoOrders) GetPriceType() (f field.PriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice gets Price, Tag 44
func (m NoOrders) GetPrice() (f field.PriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStopPx gets StopPx, Tag 99
func (m NoOrders) GetStopPx() (f field.StopPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSpread gets Spread, Tag 218
func (m NoOrders) GetSpread() (f field.SpreadField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurveCurrency gets BenchmarkCurveCurrency, Tag 220
func (m NoOrders) GetBenchmarkCurveCurrency() (f field.BenchmarkCurveCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurveName gets BenchmarkCurveName, Tag 221
func (m NoOrders) GetBenchmarkCurveName() (f field.BenchmarkCurveNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurvePoint gets BenchmarkCurvePoint, Tag 222
func (m NoOrders) GetBenchmarkCurvePoint() (f field.BenchmarkCurvePointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldType gets YieldType, Tag 235
func (m NoOrders) GetYieldType() (f field.YieldTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYield gets Yield, Tag 236
func (m NoOrders) GetYield() (f field.YieldField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m NoOrders) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplianceID gets ComplianceID, Tag 376
func (m NoOrders) GetComplianceID() (f field.ComplianceIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSolicitedFlag gets SolicitedFlag, Tag 377
func (m NoOrders) GetSolicitedFlag() (f field.SolicitedFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIid gets IOIid, Tag 23
func (m NoOrders) GetIOIid() (f field.IOIidField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteID gets QuoteID, Tag 117
func (m NoOrders) GetQuoteID() (f field.QuoteIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeInForce gets TimeInForce, Tag 59
func (m NoOrders) GetTimeInForce() (f field.TimeInForceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEffectiveTime gets EffectiveTime, Tag 168
func (m NoOrders) GetEffectiveTime() (f field.EffectiveTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireDate gets ExpireDate, Tag 432
func (m NoOrders) GetExpireDate() (f field.ExpireDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m NoOrders) GetExpireTime() (f field.ExpireTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetGTBookingInst gets GTBookingInst, Tag 427
func (m NoOrders) GetGTBookingInst() (f field.GTBookingInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommission gets Commission, Tag 12
func (m NoOrders) GetCommission() (f field.CommissionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommType gets CommType, Tag 13
func (m NoOrders) GetCommType() (f field.CommTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommCurrency gets CommCurrency, Tag 479
func (m NoOrders) GetCommCurrency() (f field.CommCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFundRenewWaiv gets FundRenewWaiv, Tag 497
func (m NoOrders) GetFundRenewWaiv() (f field.FundRenewWaivField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderCapacity gets OrderCapacity, Tag 528
func (m NoOrders) GetOrderCapacity() (f field.OrderCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderRestrictions gets OrderRestrictions, Tag 529
func (m NoOrders) GetOrderRestrictions() (f field.OrderRestrictionsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCustOrderCapacity gets CustOrderCapacity, Tag 582
func (m NoOrders) GetCustOrderCapacity() (f field.CustOrderCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRule80A gets Rule80A, Tag 47
func (m NoOrders) GetRule80A() (f field.Rule80AField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetForexReq gets ForexReq, Tag 121
func (m NoOrders) GetForexReq() (f field.ForexReqField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m NoOrders) GetSettlCurrency() (f field.SettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m NoOrders) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m NoOrders) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m NoOrders) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate2 gets FutSettDate2, Tag 193
func (m NoOrders) GetFutSettDate2() (f field.FutSettDate2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty2 gets OrderQty2, Tag 192
func (m NoOrders) GetOrderQty2() (f field.OrderQty2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice2 gets Price2, Tag 640
func (m NoOrders) GetPrice2() (f field.Price2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPositionEffect gets PositionEffect, Tag 77
func (m NoOrders) GetPositionEffect() (f field.PositionEffectField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCoveredOrUncovered gets CoveredOrUncovered, Tag 203
func (m NoOrders) GetCoveredOrUncovered() (f field.CoveredOrUncoveredField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxShow gets MaxShow, Tag 210
func (m NoOrders) GetMaxShow() (f field.MaxShowField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegDifference gets PegDifference, Tag 211
func (m NoOrders) GetPegDifference() (f field.PegDifferenceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionInst gets DiscretionInst, Tag 388
func (m NoOrders) GetDiscretionInst() (f field.DiscretionInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionOffset gets DiscretionOffset, Tag 389
func (m NoOrders) GetDiscretionOffset() (f field.DiscretionOffsetField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDesignation gets Designation, Tag 494
func (m NoOrders) GetDesignation() (f field.DesignationField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccruedInterestRate gets AccruedInterestRate, Tag 158
func (m NoOrders) GetAccruedInterestRate() (f field.AccruedInterestRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccruedInterestAmt gets AccruedInterestAmt, Tag 159
func (m NoOrders) GetAccruedInterestAmt() (f field.AccruedInterestAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNetMoney gets NetMoney, Tag 118
func (m NoOrders) GetNetMoney() (f field.NetMoneyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m NoOrders) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526
func (m NoOrders) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

//HasListSeqNo returns true if ListSeqNo is present, Tag 67
func (m NoOrders) HasListSeqNo() bool {
	return m.Has(tag.ListSeqNo)
}

//HasClOrdLinkID returns true if ClOrdLinkID is present, Tag 583
func (m NoOrders) HasClOrdLinkID() bool {
	return m.Has(tag.ClOrdLinkID)
}

//HasSettlInstMode returns true if SettlInstMode is present, Tag 160
func (m NoOrders) HasSettlInstMode() bool {
	return m.Has(tag.SettlInstMode)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m NoOrders) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasTradeOriginationDate returns true if TradeOriginationDate is present, Tag 229
func (m NoOrders) HasTradeOriginationDate() bool {
	return m.Has(tag.TradeOriginationDate)
}

//HasAccount returns true if Account is present, Tag 1
func (m NoOrders) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasAccountType returns true if AccountType is present, Tag 581
func (m NoOrders) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

//HasDayBookingInst returns true if DayBookingInst is present, Tag 589
func (m NoOrders) HasDayBookingInst() bool {
	return m.Has(tag.DayBookingInst)
}

//HasBookingUnit returns true if BookingUnit is present, Tag 590
func (m NoOrders) HasBookingUnit() bool {
	return m.Has(tag.BookingUnit)
}

//HasPreallocMethod returns true if PreallocMethod is present, Tag 591
func (m NoOrders) HasPreallocMethod() bool {
	return m.Has(tag.PreallocMethod)
}

//HasNoAllocs returns true if NoAllocs is present, Tag 78
func (m NoOrders) HasNoAllocs() bool {
	return m.Has(tag.NoAllocs)
}

//HasSettlmntTyp returns true if SettlmntTyp is present, Tag 63
func (m NoOrders) HasSettlmntTyp() bool {
	return m.Has(tag.SettlmntTyp)
}

//HasFutSettDate returns true if FutSettDate is present, Tag 64
func (m NoOrders) HasFutSettDate() bool {
	return m.Has(tag.FutSettDate)
}

//HasCashMargin returns true if CashMargin is present, Tag 544
func (m NoOrders) HasCashMargin() bool {
	return m.Has(tag.CashMargin)
}

//HasClearingFeeIndicator returns true if ClearingFeeIndicator is present, Tag 635
func (m NoOrders) HasClearingFeeIndicator() bool {
	return m.Has(tag.ClearingFeeIndicator)
}

//HasHandlInst returns true if HandlInst is present, Tag 21
func (m NoOrders) HasHandlInst() bool {
	return m.Has(tag.HandlInst)
}

//HasExecInst returns true if ExecInst is present, Tag 18
func (m NoOrders) HasExecInst() bool {
	return m.Has(tag.ExecInst)
}

//HasMinQty returns true if MinQty is present, Tag 110
func (m NoOrders) HasMinQty() bool {
	return m.Has(tag.MinQty)
}

//HasMaxFloor returns true if MaxFloor is present, Tag 111
func (m NoOrders) HasMaxFloor() bool {
	return m.Has(tag.MaxFloor)
}

//HasExDestination returns true if ExDestination is present, Tag 100
func (m NoOrders) HasExDestination() bool {
	return m.Has(tag.ExDestination)
}

//HasNoTradingSessions returns true if NoTradingSessions is present, Tag 386
func (m NoOrders) HasNoTradingSessions() bool {
	return m.Has(tag.NoTradingSessions)
}

//HasProcessCode returns true if ProcessCode is present, Tag 81
func (m NoOrders) HasProcessCode() bool {
	return m.Has(tag.ProcessCode)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m NoOrders) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m NoOrders) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m NoOrders) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m NoOrders) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m NoOrders) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m NoOrders) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m NoOrders) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m NoOrders) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m NoOrders) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m NoOrders) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m NoOrders) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m NoOrders) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m NoOrders) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m NoOrders) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m NoOrders) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m NoOrders) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m NoOrders) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m NoOrders) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m NoOrders) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m NoOrders) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m NoOrders) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m NoOrders) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m NoOrders) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m NoOrders) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m NoOrders) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m NoOrders) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m NoOrders) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m NoOrders) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m NoOrders) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m NoOrders) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m NoOrders) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m NoOrders) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m NoOrders) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasPrevClosePx returns true if PrevClosePx is present, Tag 140
func (m NoOrders) HasPrevClosePx() bool {
	return m.Has(tag.PrevClosePx)
}

//HasSide returns true if Side is present, Tag 54
func (m NoOrders) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSideValueInd returns true if SideValueInd is present, Tag 401
func (m NoOrders) HasSideValueInd() bool {
	return m.Has(tag.SideValueInd)
}

//HasLocateReqd returns true if LocateReqd is present, Tag 114
func (m NoOrders) HasLocateReqd() bool {
	return m.Has(tag.LocateReqd)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m NoOrders) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasNoStipulations returns true if NoStipulations is present, Tag 232
func (m NoOrders) HasNoStipulations() bool {
	return m.Has(tag.NoStipulations)
}

//HasQuantityType returns true if QuantityType is present, Tag 465
func (m NoOrders) HasQuantityType() bool {
	return m.Has(tag.QuantityType)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m NoOrders) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m NoOrders) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

//HasOrderPercent returns true if OrderPercent is present, Tag 516
func (m NoOrders) HasOrderPercent() bool {
	return m.Has(tag.OrderPercent)
}

//HasRoundingDirection returns true if RoundingDirection is present, Tag 468
func (m NoOrders) HasRoundingDirection() bool {
	return m.Has(tag.RoundingDirection)
}

//HasRoundingModulus returns true if RoundingModulus is present, Tag 469
func (m NoOrders) HasRoundingModulus() bool {
	return m.Has(tag.RoundingModulus)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m NoOrders) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m NoOrders) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasPrice returns true if Price is present, Tag 44
func (m NoOrders) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasStopPx returns true if StopPx is present, Tag 99
func (m NoOrders) HasStopPx() bool {
	return m.Has(tag.StopPx)
}

//HasSpread returns true if Spread is present, Tag 218
func (m NoOrders) HasSpread() bool {
	return m.Has(tag.Spread)
}

//HasBenchmarkCurveCurrency returns true if BenchmarkCurveCurrency is present, Tag 220
func (m NoOrders) HasBenchmarkCurveCurrency() bool {
	return m.Has(tag.BenchmarkCurveCurrency)
}

//HasBenchmarkCurveName returns true if BenchmarkCurveName is present, Tag 221
func (m NoOrders) HasBenchmarkCurveName() bool {
	return m.Has(tag.BenchmarkCurveName)
}

//HasBenchmarkCurvePoint returns true if BenchmarkCurvePoint is present, Tag 222
func (m NoOrders) HasBenchmarkCurvePoint() bool {
	return m.Has(tag.BenchmarkCurvePoint)
}

//HasYieldType returns true if YieldType is present, Tag 235
func (m NoOrders) HasYieldType() bool {
	return m.Has(tag.YieldType)
}

//HasYield returns true if Yield is present, Tag 236
func (m NoOrders) HasYield() bool {
	return m.Has(tag.Yield)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m NoOrders) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasComplianceID returns true if ComplianceID is present, Tag 376
func (m NoOrders) HasComplianceID() bool {
	return m.Has(tag.ComplianceID)
}

//HasSolicitedFlag returns true if SolicitedFlag is present, Tag 377
func (m NoOrders) HasSolicitedFlag() bool {
	return m.Has(tag.SolicitedFlag)
}

//HasIOIid returns true if IOIid is present, Tag 23
func (m NoOrders) HasIOIid() bool {
	return m.Has(tag.IOIid)
}

//HasQuoteID returns true if QuoteID is present, Tag 117
func (m NoOrders) HasQuoteID() bool {
	return m.Has(tag.QuoteID)
}

//HasTimeInForce returns true if TimeInForce is present, Tag 59
func (m NoOrders) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

//HasEffectiveTime returns true if EffectiveTime is present, Tag 168
func (m NoOrders) HasEffectiveTime() bool {
	return m.Has(tag.EffectiveTime)
}

//HasExpireDate returns true if ExpireDate is present, Tag 432
func (m NoOrders) HasExpireDate() bool {
	return m.Has(tag.ExpireDate)
}

//HasExpireTime returns true if ExpireTime is present, Tag 126
func (m NoOrders) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

//HasGTBookingInst returns true if GTBookingInst is present, Tag 427
func (m NoOrders) HasGTBookingInst() bool {
	return m.Has(tag.GTBookingInst)
}

//HasCommission returns true if Commission is present, Tag 12
func (m NoOrders) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m NoOrders) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasCommCurrency returns true if CommCurrency is present, Tag 479
func (m NoOrders) HasCommCurrency() bool {
	return m.Has(tag.CommCurrency)
}

//HasFundRenewWaiv returns true if FundRenewWaiv is present, Tag 497
func (m NoOrders) HasFundRenewWaiv() bool {
	return m.Has(tag.FundRenewWaiv)
}

//HasOrderCapacity returns true if OrderCapacity is present, Tag 528
func (m NoOrders) HasOrderCapacity() bool {
	return m.Has(tag.OrderCapacity)
}

//HasOrderRestrictions returns true if OrderRestrictions is present, Tag 529
func (m NoOrders) HasOrderRestrictions() bool {
	return m.Has(tag.OrderRestrictions)
}

//HasCustOrderCapacity returns true if CustOrderCapacity is present, Tag 582
func (m NoOrders) HasCustOrderCapacity() bool {
	return m.Has(tag.CustOrderCapacity)
}

//HasRule80A returns true if Rule80A is present, Tag 47
func (m NoOrders) HasRule80A() bool {
	return m.Has(tag.Rule80A)
}

//HasForexReq returns true if ForexReq is present, Tag 121
func (m NoOrders) HasForexReq() bool {
	return m.Has(tag.ForexReq)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m NoOrders) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasText returns true if Text is present, Tag 58
func (m NoOrders) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m NoOrders) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m NoOrders) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasFutSettDate2 returns true if FutSettDate2 is present, Tag 193
func (m NoOrders) HasFutSettDate2() bool {
	return m.Has(tag.FutSettDate2)
}

//HasOrderQty2 returns true if OrderQty2 is present, Tag 192
func (m NoOrders) HasOrderQty2() bool {
	return m.Has(tag.OrderQty2)
}

//HasPrice2 returns true if Price2 is present, Tag 640
func (m NoOrders) HasPrice2() bool {
	return m.Has(tag.Price2)
}

//HasPositionEffect returns true if PositionEffect is present, Tag 77
func (m NoOrders) HasPositionEffect() bool {
	return m.Has(tag.PositionEffect)
}

//HasCoveredOrUncovered returns true if CoveredOrUncovered is present, Tag 203
func (m NoOrders) HasCoveredOrUncovered() bool {
	return m.Has(tag.CoveredOrUncovered)
}

//HasMaxShow returns true if MaxShow is present, Tag 210
func (m NoOrders) HasMaxShow() bool {
	return m.Has(tag.MaxShow)
}

//HasPegDifference returns true if PegDifference is present, Tag 211
func (m NoOrders) HasPegDifference() bool {
	return m.Has(tag.PegDifference)
}

//HasDiscretionInst returns true if DiscretionInst is present, Tag 388
func (m NoOrders) HasDiscretionInst() bool {
	return m.Has(tag.DiscretionInst)
}

//HasDiscretionOffset returns true if DiscretionOffset is present, Tag 389
func (m NoOrders) HasDiscretionOffset() bool {
	return m.Has(tag.DiscretionOffset)
}

//HasDesignation returns true if Designation is present, Tag 494
func (m NoOrders) HasDesignation() bool {
	return m.Has(tag.Designation)
}

//HasAccruedInterestRate returns true if AccruedInterestRate is present, Tag 158
func (m NoOrders) HasAccruedInterestRate() bool {
	return m.Has(tag.AccruedInterestRate)
}

//HasAccruedInterestAmt returns true if AccruedInterestAmt is present, Tag 159
func (m NoOrders) HasAccruedInterestAmt() bool {
	return m.Has(tag.AccruedInterestAmt)
}

//HasNetMoney returns true if NetMoney is present, Tag 118
func (m NoOrders) HasNetMoney() bool {
	return m.Has(tag.NetMoney)
}

//NoPartyIDs is a repeating group element, Tag 453
type NoPartyIDs struct {
	quickfix.Group
}

//SetPartyID sets PartyID, Tag 448
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

//SetPartyIDSource sets PartyIDSource, Tag 447
func (m NoPartyIDs) SetPartyIDSource(v string) {
	m.Set(field.NewPartyIDSource(v))
}

//SetPartyRole sets PartyRole, Tag 452
func (m NoPartyIDs) SetPartyRole(v int) {
	m.Set(field.NewPartyRole(v))
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartyIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//GetPartyID gets PartyID, Tag 448
func (m NoPartyIDs) GetPartyID() (f field.PartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartyIDSource gets PartyIDSource, Tag 447
func (m NoPartyIDs) GetPartyIDSource() (f field.PartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartyRole gets PartyRole, Tag 452
func (m NoPartyIDs) GetPartyRole() (f field.PartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartyIDs) GetPartySubID() (f field.PartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasPartyID returns true if PartyID is present, Tag 448
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

//HasPartyIDSource returns true if PartyIDSource is present, Tag 447
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

//HasPartyRole returns true if PartyRole is present, Tag 452
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

//HasPartySubID returns true if PartySubID is present, Tag 523
func (m NoPartyIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

//NoPartyIDsRepeatingGroup is a repeating group, Tag 453
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), quickfix.GroupElement(tag.PartySubID)})}
}

//Add create and append a new NoPartyIDs to this group
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

//Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoAllocs is a repeating group element, Tag 78
type NoAllocs struct {
	quickfix.Group
}

//SetAllocAccount sets AllocAccount, Tag 79
func (m NoAllocs) SetAllocAccount(v string) {
	m.Set(field.NewAllocAccount(v))
}

//SetIndividualAllocID sets IndividualAllocID, Tag 467
func (m NoAllocs) SetIndividualAllocID(v string) {
	m.Set(field.NewIndividualAllocID(v))
}

//SetNoNestedPartyIDs sets NoNestedPartyIDs, Tag 539
func (m NoAllocs) SetNoNestedPartyIDs(f NoNestedPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetAllocQty sets AllocQty, Tag 80
func (m NoAllocs) SetAllocQty(v float64) {
	m.Set(field.NewAllocQty(v))
}

//GetAllocAccount gets AllocAccount, Tag 79
func (m NoAllocs) GetAllocAccount() (f field.AllocAccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIndividualAllocID gets IndividualAllocID, Tag 467
func (m NoAllocs) GetIndividualAllocID() (f field.IndividualAllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoNestedPartyIDs gets NoNestedPartyIDs, Tag 539
func (m NoAllocs) GetNoNestedPartyIDs() (NoNestedPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAllocQty gets AllocQty, Tag 80
func (m NoAllocs) GetAllocQty() (f field.AllocQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasAllocAccount returns true if AllocAccount is present, Tag 79
func (m NoAllocs) HasAllocAccount() bool {
	return m.Has(tag.AllocAccount)
}

//HasIndividualAllocID returns true if IndividualAllocID is present, Tag 467
func (m NoAllocs) HasIndividualAllocID() bool {
	return m.Has(tag.IndividualAllocID)
}

//HasNoNestedPartyIDs returns true if NoNestedPartyIDs is present, Tag 539
func (m NoAllocs) HasNoNestedPartyIDs() bool {
	return m.Has(tag.NoNestedPartyIDs)
}

//HasAllocQty returns true if AllocQty is present, Tag 80
func (m NoAllocs) HasAllocQty() bool {
	return m.Has(tag.AllocQty)
}

//NoNestedPartyIDs is a repeating group element, Tag 539
type NoNestedPartyIDs struct {
	quickfix.Group
}

//SetNestedPartyID sets NestedPartyID, Tag 524
func (m NoNestedPartyIDs) SetNestedPartyID(v string) {
	m.Set(field.NewNestedPartyID(v))
}

//SetNestedPartyIDSource sets NestedPartyIDSource, Tag 525
func (m NoNestedPartyIDs) SetNestedPartyIDSource(v string) {
	m.Set(field.NewNestedPartyIDSource(v))
}

//SetNestedPartyRole sets NestedPartyRole, Tag 538
func (m NoNestedPartyIDs) SetNestedPartyRole(v int) {
	m.Set(field.NewNestedPartyRole(v))
}

//SetNestedPartySubID sets NestedPartySubID, Tag 545
func (m NoNestedPartyIDs) SetNestedPartySubID(v string) {
	m.Set(field.NewNestedPartySubID(v))
}

//GetNestedPartyID gets NestedPartyID, Tag 524
func (m NoNestedPartyIDs) GetNestedPartyID() (f field.NestedPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNestedPartyIDSource gets NestedPartyIDSource, Tag 525
func (m NoNestedPartyIDs) GetNestedPartyIDSource() (f field.NestedPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNestedPartyRole gets NestedPartyRole, Tag 538
func (m NoNestedPartyIDs) GetNestedPartyRole() (f field.NestedPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNestedPartySubID gets NestedPartySubID, Tag 545
func (m NoNestedPartyIDs) GetNestedPartySubID() (f field.NestedPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasNestedPartyID returns true if NestedPartyID is present, Tag 524
func (m NoNestedPartyIDs) HasNestedPartyID() bool {
	return m.Has(tag.NestedPartyID)
}

//HasNestedPartyIDSource returns true if NestedPartyIDSource is present, Tag 525
func (m NoNestedPartyIDs) HasNestedPartyIDSource() bool {
	return m.Has(tag.NestedPartyIDSource)
}

//HasNestedPartyRole returns true if NestedPartyRole is present, Tag 538
func (m NoNestedPartyIDs) HasNestedPartyRole() bool {
	return m.Has(tag.NestedPartyRole)
}

//HasNestedPartySubID returns true if NestedPartySubID is present, Tag 545
func (m NoNestedPartyIDs) HasNestedPartySubID() bool {
	return m.Has(tag.NestedPartySubID)
}

//NoNestedPartyIDsRepeatingGroup is a repeating group, Tag 539
type NoNestedPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNestedPartyIDsRepeatingGroup returns an initialized, NoNestedPartyIDsRepeatingGroup
func NewNoNestedPartyIDsRepeatingGroup() NoNestedPartyIDsRepeatingGroup {
	return NoNestedPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNestedPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NestedPartyID), quickfix.GroupElement(tag.NestedPartyIDSource), quickfix.GroupElement(tag.NestedPartyRole), quickfix.GroupElement(tag.NestedPartySubID)})}
}

//Add create and append a new NoNestedPartyIDs to this group
func (m NoNestedPartyIDsRepeatingGroup) Add() NoNestedPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoNestedPartyIDs{g}
}

//Get returns the ith NoNestedPartyIDs in the NoNestedPartyIDsRepeatinGroup
func (m NoNestedPartyIDsRepeatingGroup) Get(i int) NoNestedPartyIDs {
	return NoNestedPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoAllocsRepeatingGroup is a repeating group, Tag 78
type NoAllocsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoAllocsRepeatingGroup returns an initialized, NoAllocsRepeatingGroup
func NewNoAllocsRepeatingGroup() NoAllocsRepeatingGroup {
	return NoAllocsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAllocs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AllocAccount), quickfix.GroupElement(tag.IndividualAllocID), NewNoNestedPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.AllocQty)})}
}

//Add create and append a new NoAllocs to this group
func (m NoAllocsRepeatingGroup) Add() NoAllocs {
	g := m.RepeatingGroup.Add()
	return NoAllocs{g}
}

//Get returns the ith NoAllocs in the NoAllocsRepeatinGroup
func (m NoAllocsRepeatingGroup) Get(i int) NoAllocs {
	return NoAllocs{m.RepeatingGroup.Get(i)}
}

//NoTradingSessions is a repeating group element, Tag 386
type NoTradingSessions struct {
	quickfix.Group
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoTradingSessions) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m NoTradingSessions) SetTradingSessionSubID(v string) {
	m.Set(field.NewTradingSessionSubID(v))
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoTradingSessions) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m NoTradingSessions) GetTradingSessionSubID() (f field.TradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m NoTradingSessions) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m NoTradingSessions) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//NoTradingSessionsRepeatingGroup is a repeating group, Tag 386
type NoTradingSessionsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoTradingSessionsRepeatingGroup returns an initialized, NoTradingSessionsRepeatingGroup
func NewNoTradingSessionsRepeatingGroup() NoTradingSessionsRepeatingGroup {
	return NoTradingSessionsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTradingSessions,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.TradingSessionSubID)})}
}

//Add create and append a new NoTradingSessions to this group
func (m NoTradingSessionsRepeatingGroup) Add() NoTradingSessions {
	g := m.RepeatingGroup.Add()
	return NoTradingSessions{g}
}

//Get returns the ith NoTradingSessions in the NoTradingSessionsRepeatinGroup
func (m NoTradingSessionsRepeatingGroup) Get(i int) NoTradingSessions {
	return NoTradingSessions{m.RepeatingGroup.Get(i)}
}

//NoSecurityAltID is a repeating group element, Tag 454
type NoSecurityAltID struct {
	quickfix.Group
}

//SetSecurityAltID sets SecurityAltID, Tag 455
func (m NoSecurityAltID) SetSecurityAltID(v string) {
	m.Set(field.NewSecurityAltID(v))
}

//SetSecurityAltIDSource sets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) SetSecurityAltIDSource(v string) {
	m.Set(field.NewSecurityAltIDSource(v))
}

//GetSecurityAltID gets SecurityAltID, Tag 455
func (m NoSecurityAltID) GetSecurityAltID() (f field.SecurityAltIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) GetSecurityAltIDSource() (f field.SecurityAltIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasSecurityAltID returns true if SecurityAltID is present, Tag 455
func (m NoSecurityAltID) HasSecurityAltID() bool {
	return m.Has(tag.SecurityAltID)
}

//HasSecurityAltIDSource returns true if SecurityAltIDSource is present, Tag 456
func (m NoSecurityAltID) HasSecurityAltIDSource() bool {
	return m.Has(tag.SecurityAltIDSource)
}

//NoSecurityAltIDRepeatingGroup is a repeating group, Tag 454
type NoSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSecurityAltIDRepeatingGroup returns an initialized, NoSecurityAltIDRepeatingGroup
func NewNoSecurityAltIDRepeatingGroup() NoSecurityAltIDRepeatingGroup {
	return NoSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SecurityAltID), quickfix.GroupElement(tag.SecurityAltIDSource)})}
}

//Add create and append a new NoSecurityAltID to this group
func (m NoSecurityAltIDRepeatingGroup) Add() NoSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoSecurityAltID{g}
}

//Get returns the ith NoSecurityAltID in the NoSecurityAltIDRepeatinGroup
func (m NoSecurityAltIDRepeatingGroup) Get(i int) NoSecurityAltID {
	return NoSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoStipulations is a repeating group element, Tag 232
type NoStipulations struct {
	quickfix.Group
}

//SetStipulationType sets StipulationType, Tag 233
func (m NoStipulations) SetStipulationType(v string) {
	m.Set(field.NewStipulationType(v))
}

//SetStipulationValue sets StipulationValue, Tag 234
func (m NoStipulations) SetStipulationValue(v string) {
	m.Set(field.NewStipulationValue(v))
}

//GetStipulationType gets StipulationType, Tag 233
func (m NoStipulations) GetStipulationType() (f field.StipulationTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStipulationValue gets StipulationValue, Tag 234
func (m NoStipulations) GetStipulationValue() (f field.StipulationValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasStipulationType returns true if StipulationType is present, Tag 233
func (m NoStipulations) HasStipulationType() bool {
	return m.Has(tag.StipulationType)
}

//HasStipulationValue returns true if StipulationValue is present, Tag 234
func (m NoStipulations) HasStipulationValue() bool {
	return m.Has(tag.StipulationValue)
}

//NoStipulationsRepeatingGroup is a repeating group, Tag 232
type NoStipulationsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoStipulationsRepeatingGroup returns an initialized, NoStipulationsRepeatingGroup
func NewNoStipulationsRepeatingGroup() NoStipulationsRepeatingGroup {
	return NoStipulationsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoStipulations,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.StipulationType), quickfix.GroupElement(tag.StipulationValue)})}
}

//Add create and append a new NoStipulations to this group
func (m NoStipulationsRepeatingGroup) Add() NoStipulations {
	g := m.RepeatingGroup.Add()
	return NoStipulations{g}
}

//Get returns the ith NoStipulations in the NoStipulationsRepeatinGroup
func (m NoStipulationsRepeatingGroup) Get(i int) NoStipulations {
	return NoStipulations{m.RepeatingGroup.Get(i)}
}

//NoOrdersRepeatingGroup is a repeating group, Tag 73
type NoOrdersRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoOrdersRepeatingGroup returns an initialized, NoOrdersRepeatingGroup
func NewNoOrdersRepeatingGroup() NoOrdersRepeatingGroup {
	return NoOrdersRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoOrders,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ClOrdID), quickfix.GroupElement(tag.SecondaryClOrdID), quickfix.GroupElement(tag.ListSeqNo), quickfix.GroupElement(tag.ClOrdLinkID), quickfix.GroupElement(tag.SettlInstMode), NewNoPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.TradeOriginationDate), quickfix.GroupElement(tag.Account), quickfix.GroupElement(tag.AccountType), quickfix.GroupElement(tag.DayBookingInst), quickfix.GroupElement(tag.BookingUnit), quickfix.GroupElement(tag.PreallocMethod), NewNoAllocsRepeatingGroup(), quickfix.GroupElement(tag.SettlmntTyp), quickfix.GroupElement(tag.FutSettDate), quickfix.GroupElement(tag.CashMargin), quickfix.GroupElement(tag.ClearingFeeIndicator), quickfix.GroupElement(tag.HandlInst), quickfix.GroupElement(tag.ExecInst), quickfix.GroupElement(tag.MinQty), quickfix.GroupElement(tag.MaxFloor), quickfix.GroupElement(tag.ExDestination), NewNoTradingSessionsRepeatingGroup(), quickfix.GroupElement(tag.ProcessCode), quickfix.GroupElement(tag.Symbol), quickfix.GroupElement(tag.SymbolSfx), quickfix.GroupElement(tag.SecurityID), quickfix.GroupElement(tag.SecurityIDSource), NewNoSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.Product), quickfix.GroupElement(tag.CFICode), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.MaturityMonthYear), quickfix.GroupElement(tag.MaturityDate), quickfix.GroupElement(tag.CouponPaymentDate), quickfix.GroupElement(tag.IssueDate), quickfix.GroupElement(tag.RepoCollateralSecurityType), quickfix.GroupElement(tag.RepurchaseTerm), quickfix.GroupElement(tag.RepurchaseRate), quickfix.GroupElement(tag.Factor), quickfix.GroupElement(tag.CreditRating), quickfix.GroupElement(tag.InstrRegistry), quickfix.GroupElement(tag.CountryOfIssue), quickfix.GroupElement(tag.StateOrProvinceOfIssue), quickfix.GroupElement(tag.LocaleOfIssue), quickfix.GroupElement(tag.RedemptionDate), quickfix.GroupElement(tag.StrikePrice), quickfix.GroupElement(tag.OptAttribute), quickfix.GroupElement(tag.ContractMultiplier), quickfix.GroupElement(tag.CouponRate), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.Issuer), quickfix.GroupElement(tag.EncodedIssuerLen), quickfix.GroupElement(tag.EncodedIssuer), quickfix.GroupElement(tag.SecurityDesc), quickfix.GroupElement(tag.EncodedSecurityDescLen), quickfix.GroupElement(tag.EncodedSecurityDesc), quickfix.GroupElement(tag.PrevClosePx), quickfix.GroupElement(tag.Side), quickfix.GroupElement(tag.SideValueInd), quickfix.GroupElement(tag.LocateReqd), quickfix.GroupElement(tag.TransactTime), NewNoStipulationsRepeatingGroup(), quickfix.GroupElement(tag.QuantityType), quickfix.GroupElement(tag.OrderQty), quickfix.GroupElement(tag.CashOrderQty), quickfix.GroupElement(tag.OrderPercent), quickfix.GroupElement(tag.RoundingDirection), quickfix.GroupElement(tag.RoundingModulus), quickfix.GroupElement(tag.OrdType), quickfix.GroupElement(tag.PriceType), quickfix.GroupElement(tag.Price), quickfix.GroupElement(tag.StopPx), quickfix.GroupElement(tag.Spread), quickfix.GroupElement(tag.BenchmarkCurveCurrency), quickfix.GroupElement(tag.BenchmarkCurveName), quickfix.GroupElement(tag.BenchmarkCurvePoint), quickfix.GroupElement(tag.YieldType), quickfix.GroupElement(tag.Yield), quickfix.GroupElement(tag.Currency), quickfix.GroupElement(tag.ComplianceID), quickfix.GroupElement(tag.SolicitedFlag), quickfix.GroupElement(tag.IOIid), quickfix.GroupElement(tag.QuoteID), quickfix.GroupElement(tag.TimeInForce), quickfix.GroupElement(tag.EffectiveTime), quickfix.GroupElement(tag.ExpireDate), quickfix.GroupElement(tag.ExpireTime), quickfix.GroupElement(tag.GTBookingInst), quickfix.GroupElement(tag.Commission), quickfix.GroupElement(tag.CommType), quickfix.GroupElement(tag.CommCurrency), quickfix.GroupElement(tag.FundRenewWaiv), quickfix.GroupElement(tag.OrderCapacity), quickfix.GroupElement(tag.OrderRestrictions), quickfix.GroupElement(tag.CustOrderCapacity), quickfix.GroupElement(tag.Rule80A), quickfix.GroupElement(tag.ForexReq), quickfix.GroupElement(tag.SettlCurrency), quickfix.GroupElement(tag.Text), quickfix.GroupElement(tag.EncodedTextLen), quickfix.GroupElement(tag.EncodedText), quickfix.GroupElement(tag.FutSettDate2), quickfix.GroupElement(tag.OrderQty2), quickfix.GroupElement(tag.Price2), quickfix.GroupElement(tag.PositionEffect), quickfix.GroupElement(tag.CoveredOrUncovered), quickfix.GroupElement(tag.MaxShow), quickfix.GroupElement(tag.PegDifference), quickfix.GroupElement(tag.DiscretionInst), quickfix.GroupElement(tag.DiscretionOffset), quickfix.GroupElement(tag.Designation), quickfix.GroupElement(tag.AccruedInterestRate), quickfix.GroupElement(tag.AccruedInterestAmt), quickfix.GroupElement(tag.NetMoney)})}
}

//Add create and append a new NoOrders to this group
func (m NoOrdersRepeatingGroup) Add() NoOrders {
	g := m.RepeatingGroup.Add()
	return NoOrders{g}
}

//Get returns the ith NoOrders in the NoOrdersRepeatinGroup
func (m NoOrdersRepeatingGroup) Get(i int) NoOrders {
	return NoOrders{m.RepeatingGroup.Get(i)}
}
