package ioi

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//IOI is the fix50 IOI type, MsgType = 6
type IOI struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
}

//FromMessage creates a IOI from a quickfix.Message instance
func FromMessage(m quickfix.Message) IOI {
	return IOI{
		Header:  fixt11.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fixt11.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m IOI) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a IOI initialized with the required fields for IOI
func New(ioiid field.IOIIDField, ioitranstype field.IOITransTypeField, side field.SideField, ioiqty field.IOIQtyField) (m IOI) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("6"))
	m.Set(ioiid)
	m.Set(ioitranstype)
	m.Set(side)
	m.Set(ioiqty)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg IOI, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "7", "6", r
}

//SetCurrency sets Currency, Tag 15
func (m IOI) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m IOI) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetIOIID sets IOIID, Tag 23
func (m IOI) SetIOIID(v string) {
	m.Set(field.NewIOIID(v))
}

//SetIOIQltyInd sets IOIQltyInd, Tag 25
func (m IOI) SetIOIQltyInd(v string) {
	m.Set(field.NewIOIQltyInd(v))
}

//SetIOIRefID sets IOIRefID, Tag 26
func (m IOI) SetIOIRefID(v string) {
	m.Set(field.NewIOIRefID(v))
}

//SetIOIQty sets IOIQty, Tag 27
func (m IOI) SetIOIQty(v string) {
	m.Set(field.NewIOIQty(v))
}

//SetIOITransType sets IOITransType, Tag 28
func (m IOI) SetIOITransType(v string) {
	m.Set(field.NewIOITransType(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m IOI) SetOrderQty(v float64) {
	m.Set(field.NewOrderQty(v))
}

//SetPrice sets Price, Tag 44
func (m IOI) SetPrice(v float64) {
	m.Set(field.NewPrice(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m IOI) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m IOI) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m IOI) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m IOI) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m IOI) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetValidUntilTime sets ValidUntilTime, Tag 62
func (m IOI) SetValidUntilTime(v time.Time) {
	m.Set(field.NewValidUntilTime(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m IOI) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetIssuer sets Issuer, Tag 106
func (m IOI) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m IOI) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetIOINaturalFlag sets IOINaturalFlag, Tag 130
func (m IOI) SetIOINaturalFlag(v bool) {
	m.Set(field.NewIOINaturalFlag(v))
}

//SetURLLink sets URLLink, Tag 149
func (m IOI) SetURLLink(v string) {
	m.Set(field.NewURLLink(v))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m IOI) SetCashOrderQty(v float64) {
	m.Set(field.NewCashOrderQty(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m IOI) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetNoIOIQualifiers sets NoIOIQualifiers, Tag 199
func (m IOI) SetNoIOIQualifiers(f NoIOIQualifiersRepeatingGroup) {
	m.SetGroup(f)
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m IOI) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m IOI) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m IOI) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m IOI) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetNoRoutingIDs sets NoRoutingIDs, Tag 215
func (m IOI) SetNoRoutingIDs(f NoRoutingIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetSpread sets Spread, Tag 218
func (m IOI) SetSpread(v float64) {
	m.Set(field.NewSpread(v))
}

//SetBenchmarkCurveCurrency sets BenchmarkCurveCurrency, Tag 220
func (m IOI) SetBenchmarkCurveCurrency(v string) {
	m.Set(field.NewBenchmarkCurveCurrency(v))
}

//SetBenchmarkCurveName sets BenchmarkCurveName, Tag 221
func (m IOI) SetBenchmarkCurveName(v string) {
	m.Set(field.NewBenchmarkCurveName(v))
}

//SetBenchmarkCurvePoint sets BenchmarkCurvePoint, Tag 222
func (m IOI) SetBenchmarkCurvePoint(v string) {
	m.Set(field.NewBenchmarkCurvePoint(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m IOI) SetCouponRate(v float64) {
	m.Set(field.NewCouponRate(v))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m IOI) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m IOI) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m IOI) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m IOI) SetRepurchaseRate(v float64) {
	m.Set(field.NewRepurchaseRate(v))
}

//SetFactor sets Factor, Tag 228
func (m IOI) SetFactor(v float64) {
	m.Set(field.NewFactor(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m IOI) SetContractMultiplier(v float64) {
	m.Set(field.NewContractMultiplier(v))
}

//SetNoStipulations sets NoStipulations, Tag 232
func (m IOI) SetNoStipulations(f NoStipulationsRepeatingGroup) {
	m.SetGroup(f)
}

//SetYieldType sets YieldType, Tag 235
func (m IOI) SetYieldType(v string) {
	m.Set(field.NewYieldType(v))
}

//SetYield sets Yield, Tag 236
func (m IOI) SetYield(v float64) {
	m.Set(field.NewYield(v))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m IOI) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m IOI) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m IOI) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m IOI) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m IOI) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m IOI) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m IOI) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m IOI) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m IOI) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetPriceType sets PriceType, Tag 423
func (m IOI) SetPriceType(v int) {
	m.Set(field.NewPriceType(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m IOI) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m IOI) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m IOI) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m IOI) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetRoundingDirection sets RoundingDirection, Tag 468
func (m IOI) SetRoundingDirection(v string) {
	m.Set(field.NewRoundingDirection(v))
}

//SetRoundingModulus sets RoundingModulus, Tag 469
func (m IOI) SetRoundingModulus(v float64) {
	m.Set(field.NewRoundingModulus(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m IOI) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m IOI) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m IOI) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetOrderPercent sets OrderPercent, Tag 516
func (m IOI) SetOrderPercent(v float64) {
	m.Set(field.NewOrderPercent(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m IOI) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m IOI) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetNoLegs sets NoLegs, Tag 555
func (m IOI) SetNoLegs(f NoLegsRepeatingGroup) {
	m.SetGroup(f)
}

//SetBenchmarkPrice sets BenchmarkPrice, Tag 662
func (m IOI) SetBenchmarkPrice(v float64) {
	m.Set(field.NewBenchmarkPrice(v))
}

//SetBenchmarkPriceType sets BenchmarkPriceType, Tag 663
func (m IOI) SetBenchmarkPriceType(v int) {
	m.Set(field.NewBenchmarkPriceType(v))
}

//SetContractSettlMonth sets ContractSettlMonth, Tag 667
func (m IOI) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

//SetPool sets Pool, Tag 691
func (m IOI) SetPool(v string) {
	m.Set(field.NewPool(v))
}

//SetYieldRedemptionDate sets YieldRedemptionDate, Tag 696
func (m IOI) SetYieldRedemptionDate(v string) {
	m.Set(field.NewYieldRedemptionDate(v))
}

//SetYieldRedemptionPrice sets YieldRedemptionPrice, Tag 697
func (m IOI) SetYieldRedemptionPrice(v float64) {
	m.Set(field.NewYieldRedemptionPrice(v))
}

//SetYieldRedemptionPriceType sets YieldRedemptionPriceType, Tag 698
func (m IOI) SetYieldRedemptionPriceType(v int) {
	m.Set(field.NewYieldRedemptionPriceType(v))
}

//SetBenchmarkSecurityID sets BenchmarkSecurityID, Tag 699
func (m IOI) SetBenchmarkSecurityID(v string) {
	m.Set(field.NewBenchmarkSecurityID(v))
}

//SetYieldCalcDate sets YieldCalcDate, Tag 701
func (m IOI) SetYieldCalcDate(v string) {
	m.Set(field.NewYieldCalcDate(v))
}

//SetNoUnderlyings sets NoUnderlyings, Tag 711
func (m IOI) SetNoUnderlyings(f NoUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

//SetBenchmarkSecurityIDSource sets BenchmarkSecurityIDSource, Tag 761
func (m IOI) SetBenchmarkSecurityIDSource(v string) {
	m.Set(field.NewBenchmarkSecurityIDSource(v))
}

//SetSecuritySubType sets SecuritySubType, Tag 762
func (m IOI) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

//SetTerminationType sets TerminationType, Tag 788
func (m IOI) SetTerminationType(v int) {
	m.Set(field.NewTerminationType(v))
}

//SetQtyType sets QtyType, Tag 854
func (m IOI) SetQtyType(v int) {
	m.Set(field.NewQtyType(v))
}

//SetNoEvents sets NoEvents, Tag 864
func (m IOI) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDatedDate sets DatedDate, Tag 873
func (m IOI) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

//SetInterestAccrualDate sets InterestAccrualDate, Tag 874
func (m IOI) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

//SetCPProgram sets CPProgram, Tag 875
func (m IOI) SetCPProgram(v int) {
	m.Set(field.NewCPProgram(v))
}

//SetCPRegType sets CPRegType, Tag 876
func (m IOI) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

//SetMarginRatio sets MarginRatio, Tag 898
func (m IOI) SetMarginRatio(v float64) {
	m.Set(field.NewMarginRatio(v))
}

//SetAgreementDesc sets AgreementDesc, Tag 913
func (m IOI) SetAgreementDesc(v string) {
	m.Set(field.NewAgreementDesc(v))
}

//SetAgreementID sets AgreementID, Tag 914
func (m IOI) SetAgreementID(v string) {
	m.Set(field.NewAgreementID(v))
}

//SetAgreementDate sets AgreementDate, Tag 915
func (m IOI) SetAgreementDate(v string) {
	m.Set(field.NewAgreementDate(v))
}

//SetStartDate sets StartDate, Tag 916
func (m IOI) SetStartDate(v string) {
	m.Set(field.NewStartDate(v))
}

//SetEndDate sets EndDate, Tag 917
func (m IOI) SetEndDate(v string) {
	m.Set(field.NewEndDate(v))
}

//SetAgreementCurrency sets AgreementCurrency, Tag 918
func (m IOI) SetAgreementCurrency(v string) {
	m.Set(field.NewAgreementCurrency(v))
}

//SetDeliveryType sets DeliveryType, Tag 919
func (m IOI) SetDeliveryType(v int) {
	m.Set(field.NewDeliveryType(v))
}

//SetStrikeCurrency sets StrikeCurrency, Tag 947
func (m IOI) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

//SetSecurityStatus sets SecurityStatus, Tag 965
func (m IOI) SetSecurityStatus(v string) {
	m.Set(field.NewSecurityStatus(v))
}

//SetSettleOnOpenFlag sets SettleOnOpenFlag, Tag 966
func (m IOI) SetSettleOnOpenFlag(v string) {
	m.Set(field.NewSettleOnOpenFlag(v))
}

//SetStrikeMultiplier sets StrikeMultiplier, Tag 967
func (m IOI) SetStrikeMultiplier(v float64) {
	m.Set(field.NewStrikeMultiplier(v))
}

//SetStrikeValue sets StrikeValue, Tag 968
func (m IOI) SetStrikeValue(v float64) {
	m.Set(field.NewStrikeValue(v))
}

//SetMinPriceIncrement sets MinPriceIncrement, Tag 969
func (m IOI) SetMinPriceIncrement(v float64) {
	m.Set(field.NewMinPriceIncrement(v))
}

//SetPositionLimit sets PositionLimit, Tag 970
func (m IOI) SetPositionLimit(v int) {
	m.Set(field.NewPositionLimit(v))
}

//SetNTPositionLimit sets NTPositionLimit, Tag 971
func (m IOI) SetNTPositionLimit(v int) {
	m.Set(field.NewNTPositionLimit(v))
}

//SetUnitOfMeasure sets UnitOfMeasure, Tag 996
func (m IOI) SetUnitOfMeasure(v string) {
	m.Set(field.NewUnitOfMeasure(v))
}

//SetTimeUnit sets TimeUnit, Tag 997
func (m IOI) SetTimeUnit(v string) {
	m.Set(field.NewTimeUnit(v))
}

//SetNoInstrumentParties sets NoInstrumentParties, Tag 1018
func (m IOI) SetNoInstrumentParties(f NoInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetInstrmtAssignmentMethod sets InstrmtAssignmentMethod, Tag 1049
func (m IOI) SetInstrmtAssignmentMethod(v string) {
	m.Set(field.NewInstrmtAssignmentMethod(v))
}

//SetMaturityTime sets MaturityTime, Tag 1079
func (m IOI) SetMaturityTime(v string) {
	m.Set(field.NewMaturityTime(v))
}

//GetCurrency gets Currency, Tag 15
func (m IOI) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m IOI) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIID gets IOIID, Tag 23
func (m IOI) GetIOIID() (f field.IOIIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIQltyInd gets IOIQltyInd, Tag 25
func (m IOI) GetIOIQltyInd() (f field.IOIQltyIndField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIRefID gets IOIRefID, Tag 26
func (m IOI) GetIOIRefID() (f field.IOIRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIQty gets IOIQty, Tag 27
func (m IOI) GetIOIQty() (f field.IOIQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOITransType gets IOITransType, Tag 28
func (m IOI) GetIOITransType() (f field.IOITransTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m IOI) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice gets Price, Tag 44
func (m IOI) GetPrice() (f field.PriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m IOI) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m IOI) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m IOI) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m IOI) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m IOI) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetValidUntilTime gets ValidUntilTime, Tag 62
func (m IOI) GetValidUntilTime() (f field.ValidUntilTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m IOI) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m IOI) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m IOI) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOINaturalFlag gets IOINaturalFlag, Tag 130
func (m IOI) GetIOINaturalFlag() (f field.IOINaturalFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetURLLink gets URLLink, Tag 149
func (m IOI) GetURLLink() (f field.URLLinkField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m IOI) GetCashOrderQty() (f field.CashOrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m IOI) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoIOIQualifiers gets NoIOIQualifiers, Tag 199
func (m IOI) GetNoIOIQualifiers() (NoIOIQualifiersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoIOIQualifiersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m IOI) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m IOI) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m IOI) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m IOI) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoRoutingIDs gets NoRoutingIDs, Tag 215
func (m IOI) GetNoRoutingIDs() (NoRoutingIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRoutingIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSpread gets Spread, Tag 218
func (m IOI) GetSpread() (f field.SpreadField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurveCurrency gets BenchmarkCurveCurrency, Tag 220
func (m IOI) GetBenchmarkCurveCurrency() (f field.BenchmarkCurveCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurveName gets BenchmarkCurveName, Tag 221
func (m IOI) GetBenchmarkCurveName() (f field.BenchmarkCurveNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurvePoint gets BenchmarkCurvePoint, Tag 222
func (m IOI) GetBenchmarkCurvePoint() (f field.BenchmarkCurvePointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m IOI) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m IOI) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m IOI) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m IOI) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m IOI) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m IOI) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m IOI) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoStipulations gets NoStipulations, Tag 232
func (m IOI) GetNoStipulations() (NoStipulationsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoStipulationsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetYieldType gets YieldType, Tag 235
func (m IOI) GetYieldType() (f field.YieldTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYield gets Yield, Tag 236
func (m IOI) GetYield() (f field.YieldField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m IOI) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m IOI) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m IOI) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m IOI) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m IOI) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m IOI) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m IOI) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m IOI) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m IOI) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceType gets PriceType, Tag 423
func (m IOI) GetPriceType() (f field.PriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m IOI) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m IOI) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m IOI) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m IOI) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingDirection gets RoundingDirection, Tag 468
func (m IOI) GetRoundingDirection() (f field.RoundingDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingModulus gets RoundingModulus, Tag 469
func (m IOI) GetRoundingModulus() (f field.RoundingModulusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m IOI) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m IOI) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m IOI) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderPercent gets OrderPercent, Tag 516
func (m IOI) GetOrderPercent() (f field.OrderPercentField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m IOI) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m IOI) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoLegs gets NoLegs, Tag 555
func (m IOI) GetNoLegs() (NoLegsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetBenchmarkPrice gets BenchmarkPrice, Tag 662
func (m IOI) GetBenchmarkPrice() (f field.BenchmarkPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkPriceType gets BenchmarkPriceType, Tag 663
func (m IOI) GetBenchmarkPriceType() (f field.BenchmarkPriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractSettlMonth gets ContractSettlMonth, Tag 667
func (m IOI) GetContractSettlMonth() (f field.ContractSettlMonthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPool gets Pool, Tag 691
func (m IOI) GetPool() (f field.PoolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldRedemptionDate gets YieldRedemptionDate, Tag 696
func (m IOI) GetYieldRedemptionDate() (f field.YieldRedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldRedemptionPrice gets YieldRedemptionPrice, Tag 697
func (m IOI) GetYieldRedemptionPrice() (f field.YieldRedemptionPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldRedemptionPriceType gets YieldRedemptionPriceType, Tag 698
func (m IOI) GetYieldRedemptionPriceType() (f field.YieldRedemptionPriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkSecurityID gets BenchmarkSecurityID, Tag 699
func (m IOI) GetBenchmarkSecurityID() (f field.BenchmarkSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldCalcDate gets YieldCalcDate, Tag 701
func (m IOI) GetYieldCalcDate() (f field.YieldCalcDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyings gets NoUnderlyings, Tag 711
func (m IOI) GetNoUnderlyings() (NoUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetBenchmarkSecurityIDSource gets BenchmarkSecurityIDSource, Tag 761
func (m IOI) GetBenchmarkSecurityIDSource() (f field.BenchmarkSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m IOI) GetSecuritySubType() (f field.SecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTerminationType gets TerminationType, Tag 788
func (m IOI) GetTerminationType() (f field.TerminationTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQtyType gets QtyType, Tag 854
func (m IOI) GetQtyType() (f field.QtyTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoEvents gets NoEvents, Tag 864
func (m IOI) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDatedDate gets DatedDate, Tag 873
func (m IOI) GetDatedDate() (f field.DatedDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInterestAccrualDate gets InterestAccrualDate, Tag 874
func (m IOI) GetInterestAccrualDate() (f field.InterestAccrualDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPProgram gets CPProgram, Tag 875
func (m IOI) GetCPProgram() (f field.CPProgramField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPRegType gets CPRegType, Tag 876
func (m IOI) GetCPRegType() (f field.CPRegTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMarginRatio gets MarginRatio, Tag 898
func (m IOI) GetMarginRatio() (f field.MarginRatioField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementDesc gets AgreementDesc, Tag 913
func (m IOI) GetAgreementDesc() (f field.AgreementDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementID gets AgreementID, Tag 914
func (m IOI) GetAgreementID() (f field.AgreementIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementDate gets AgreementDate, Tag 915
func (m IOI) GetAgreementDate() (f field.AgreementDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStartDate gets StartDate, Tag 916
func (m IOI) GetStartDate() (f field.StartDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEndDate gets EndDate, Tag 917
func (m IOI) GetEndDate() (f field.EndDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAgreementCurrency gets AgreementCurrency, Tag 918
func (m IOI) GetAgreementCurrency() (f field.AgreementCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDeliveryType gets DeliveryType, Tag 919
func (m IOI) GetDeliveryType() (f field.DeliveryTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeCurrency gets StrikeCurrency, Tag 947
func (m IOI) GetStrikeCurrency() (f field.StrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityStatus gets SecurityStatus, Tag 965
func (m IOI) GetSecurityStatus() (f field.SecurityStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettleOnOpenFlag gets SettleOnOpenFlag, Tag 966
func (m IOI) GetSettleOnOpenFlag() (f field.SettleOnOpenFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeMultiplier gets StrikeMultiplier, Tag 967
func (m IOI) GetStrikeMultiplier() (f field.StrikeMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeValue gets StrikeValue, Tag 968
func (m IOI) GetStrikeValue() (f field.StrikeValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinPriceIncrement gets MinPriceIncrement, Tag 969
func (m IOI) GetMinPriceIncrement() (f field.MinPriceIncrementField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPositionLimit gets PositionLimit, Tag 970
func (m IOI) GetPositionLimit() (f field.PositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNTPositionLimit gets NTPositionLimit, Tag 971
func (m IOI) GetNTPositionLimit() (f field.NTPositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnitOfMeasure gets UnitOfMeasure, Tag 996
func (m IOI) GetUnitOfMeasure() (f field.UnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeUnit gets TimeUnit, Tag 997
func (m IOI) GetTimeUnit() (f field.TimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoInstrumentParties gets NoInstrumentParties, Tag 1018
func (m IOI) GetNoInstrumentParties() (NoInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetInstrmtAssignmentMethod gets InstrmtAssignmentMethod, Tag 1049
func (m IOI) GetInstrmtAssignmentMethod() (f field.InstrmtAssignmentMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityTime gets MaturityTime, Tag 1079
func (m IOI) GetMaturityTime() (f field.MaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasCurrency returns true if Currency is present, Tag 15
func (m IOI) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m IOI) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasIOIID returns true if IOIID is present, Tag 23
func (m IOI) HasIOIID() bool {
	return m.Has(tag.IOIID)
}

//HasIOIQltyInd returns true if IOIQltyInd is present, Tag 25
func (m IOI) HasIOIQltyInd() bool {
	return m.Has(tag.IOIQltyInd)
}

//HasIOIRefID returns true if IOIRefID is present, Tag 26
func (m IOI) HasIOIRefID() bool {
	return m.Has(tag.IOIRefID)
}

//HasIOIQty returns true if IOIQty is present, Tag 27
func (m IOI) HasIOIQty() bool {
	return m.Has(tag.IOIQty)
}

//HasIOITransType returns true if IOITransType is present, Tag 28
func (m IOI) HasIOITransType() bool {
	return m.Has(tag.IOITransType)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m IOI) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasPrice returns true if Price is present, Tag 44
func (m IOI) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m IOI) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSide returns true if Side is present, Tag 54
func (m IOI) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m IOI) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m IOI) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m IOI) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasValidUntilTime returns true if ValidUntilTime is present, Tag 62
func (m IOI) HasValidUntilTime() bool {
	return m.Has(tag.ValidUntilTime)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m IOI) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m IOI) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m IOI) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasIOINaturalFlag returns true if IOINaturalFlag is present, Tag 130
func (m IOI) HasIOINaturalFlag() bool {
	return m.Has(tag.IOINaturalFlag)
}

//HasURLLink returns true if URLLink is present, Tag 149
func (m IOI) HasURLLink() bool {
	return m.Has(tag.URLLink)
}

//HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m IOI) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m IOI) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasNoIOIQualifiers returns true if NoIOIQualifiers is present, Tag 199
func (m IOI) HasNoIOIQualifiers() bool {
	return m.Has(tag.NoIOIQualifiers)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m IOI) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m IOI) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m IOI) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m IOI) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasNoRoutingIDs returns true if NoRoutingIDs is present, Tag 215
func (m IOI) HasNoRoutingIDs() bool {
	return m.Has(tag.NoRoutingIDs)
}

//HasSpread returns true if Spread is present, Tag 218
func (m IOI) HasSpread() bool {
	return m.Has(tag.Spread)
}

//HasBenchmarkCurveCurrency returns true if BenchmarkCurveCurrency is present, Tag 220
func (m IOI) HasBenchmarkCurveCurrency() bool {
	return m.Has(tag.BenchmarkCurveCurrency)
}

//HasBenchmarkCurveName returns true if BenchmarkCurveName is present, Tag 221
func (m IOI) HasBenchmarkCurveName() bool {
	return m.Has(tag.BenchmarkCurveName)
}

//HasBenchmarkCurvePoint returns true if BenchmarkCurvePoint is present, Tag 222
func (m IOI) HasBenchmarkCurvePoint() bool {
	return m.Has(tag.BenchmarkCurvePoint)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m IOI) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m IOI) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m IOI) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m IOI) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m IOI) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m IOI) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m IOI) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasNoStipulations returns true if NoStipulations is present, Tag 232
func (m IOI) HasNoStipulations() bool {
	return m.Has(tag.NoStipulations)
}

//HasYieldType returns true if YieldType is present, Tag 235
func (m IOI) HasYieldType() bool {
	return m.Has(tag.YieldType)
}

//HasYield returns true if Yield is present, Tag 236
func (m IOI) HasYield() bool {
	return m.Has(tag.Yield)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m IOI) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m IOI) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m IOI) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m IOI) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m IOI) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m IOI) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m IOI) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m IOI) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m IOI) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m IOI) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m IOI) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m IOI) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m IOI) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m IOI) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasRoundingDirection returns true if RoundingDirection is present, Tag 468
func (m IOI) HasRoundingDirection() bool {
	return m.Has(tag.RoundingDirection)
}

//HasRoundingModulus returns true if RoundingModulus is present, Tag 469
func (m IOI) HasRoundingModulus() bool {
	return m.Has(tag.RoundingModulus)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m IOI) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m IOI) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m IOI) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasOrderPercent returns true if OrderPercent is present, Tag 516
func (m IOI) HasOrderPercent() bool {
	return m.Has(tag.OrderPercent)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m IOI) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m IOI) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasNoLegs returns true if NoLegs is present, Tag 555
func (m IOI) HasNoLegs() bool {
	return m.Has(tag.NoLegs)
}

//HasBenchmarkPrice returns true if BenchmarkPrice is present, Tag 662
func (m IOI) HasBenchmarkPrice() bool {
	return m.Has(tag.BenchmarkPrice)
}

//HasBenchmarkPriceType returns true if BenchmarkPriceType is present, Tag 663
func (m IOI) HasBenchmarkPriceType() bool {
	return m.Has(tag.BenchmarkPriceType)
}

//HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667
func (m IOI) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

//HasPool returns true if Pool is present, Tag 691
func (m IOI) HasPool() bool {
	return m.Has(tag.Pool)
}

//HasYieldRedemptionDate returns true if YieldRedemptionDate is present, Tag 696
func (m IOI) HasYieldRedemptionDate() bool {
	return m.Has(tag.YieldRedemptionDate)
}

//HasYieldRedemptionPrice returns true if YieldRedemptionPrice is present, Tag 697
func (m IOI) HasYieldRedemptionPrice() bool {
	return m.Has(tag.YieldRedemptionPrice)
}

//HasYieldRedemptionPriceType returns true if YieldRedemptionPriceType is present, Tag 698
func (m IOI) HasYieldRedemptionPriceType() bool {
	return m.Has(tag.YieldRedemptionPriceType)
}

//HasBenchmarkSecurityID returns true if BenchmarkSecurityID is present, Tag 699
func (m IOI) HasBenchmarkSecurityID() bool {
	return m.Has(tag.BenchmarkSecurityID)
}

//HasYieldCalcDate returns true if YieldCalcDate is present, Tag 701
func (m IOI) HasYieldCalcDate() bool {
	return m.Has(tag.YieldCalcDate)
}

//HasNoUnderlyings returns true if NoUnderlyings is present, Tag 711
func (m IOI) HasNoUnderlyings() bool {
	return m.Has(tag.NoUnderlyings)
}

//HasBenchmarkSecurityIDSource returns true if BenchmarkSecurityIDSource is present, Tag 761
func (m IOI) HasBenchmarkSecurityIDSource() bool {
	return m.Has(tag.BenchmarkSecurityIDSource)
}

//HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m IOI) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

//HasTerminationType returns true if TerminationType is present, Tag 788
func (m IOI) HasTerminationType() bool {
	return m.Has(tag.TerminationType)
}

//HasQtyType returns true if QtyType is present, Tag 854
func (m IOI) HasQtyType() bool {
	return m.Has(tag.QtyType)
}

//HasNoEvents returns true if NoEvents is present, Tag 864
func (m IOI) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

//HasDatedDate returns true if DatedDate is present, Tag 873
func (m IOI) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

//HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874
func (m IOI) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

//HasCPProgram returns true if CPProgram is present, Tag 875
func (m IOI) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

//HasCPRegType returns true if CPRegType is present, Tag 876
func (m IOI) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

//HasMarginRatio returns true if MarginRatio is present, Tag 898
func (m IOI) HasMarginRatio() bool {
	return m.Has(tag.MarginRatio)
}

//HasAgreementDesc returns true if AgreementDesc is present, Tag 913
func (m IOI) HasAgreementDesc() bool {
	return m.Has(tag.AgreementDesc)
}

//HasAgreementID returns true if AgreementID is present, Tag 914
func (m IOI) HasAgreementID() bool {
	return m.Has(tag.AgreementID)
}

//HasAgreementDate returns true if AgreementDate is present, Tag 915
func (m IOI) HasAgreementDate() bool {
	return m.Has(tag.AgreementDate)
}

//HasStartDate returns true if StartDate is present, Tag 916
func (m IOI) HasStartDate() bool {
	return m.Has(tag.StartDate)
}

//HasEndDate returns true if EndDate is present, Tag 917
func (m IOI) HasEndDate() bool {
	return m.Has(tag.EndDate)
}

//HasAgreementCurrency returns true if AgreementCurrency is present, Tag 918
func (m IOI) HasAgreementCurrency() bool {
	return m.Has(tag.AgreementCurrency)
}

//HasDeliveryType returns true if DeliveryType is present, Tag 919
func (m IOI) HasDeliveryType() bool {
	return m.Has(tag.DeliveryType)
}

//HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947
func (m IOI) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

//HasSecurityStatus returns true if SecurityStatus is present, Tag 965
func (m IOI) HasSecurityStatus() bool {
	return m.Has(tag.SecurityStatus)
}

//HasSettleOnOpenFlag returns true if SettleOnOpenFlag is present, Tag 966
func (m IOI) HasSettleOnOpenFlag() bool {
	return m.Has(tag.SettleOnOpenFlag)
}

//HasStrikeMultiplier returns true if StrikeMultiplier is present, Tag 967
func (m IOI) HasStrikeMultiplier() bool {
	return m.Has(tag.StrikeMultiplier)
}

//HasStrikeValue returns true if StrikeValue is present, Tag 968
func (m IOI) HasStrikeValue() bool {
	return m.Has(tag.StrikeValue)
}

//HasMinPriceIncrement returns true if MinPriceIncrement is present, Tag 969
func (m IOI) HasMinPriceIncrement() bool {
	return m.Has(tag.MinPriceIncrement)
}

//HasPositionLimit returns true if PositionLimit is present, Tag 970
func (m IOI) HasPositionLimit() bool {
	return m.Has(tag.PositionLimit)
}

//HasNTPositionLimit returns true if NTPositionLimit is present, Tag 971
func (m IOI) HasNTPositionLimit() bool {
	return m.Has(tag.NTPositionLimit)
}

//HasUnitOfMeasure returns true if UnitOfMeasure is present, Tag 996
func (m IOI) HasUnitOfMeasure() bool {
	return m.Has(tag.UnitOfMeasure)
}

//HasTimeUnit returns true if TimeUnit is present, Tag 997
func (m IOI) HasTimeUnit() bool {
	return m.Has(tag.TimeUnit)
}

//HasNoInstrumentParties returns true if NoInstrumentParties is present, Tag 1018
func (m IOI) HasNoInstrumentParties() bool {
	return m.Has(tag.NoInstrumentParties)
}

//HasInstrmtAssignmentMethod returns true if InstrmtAssignmentMethod is present, Tag 1049
func (m IOI) HasInstrmtAssignmentMethod() bool {
	return m.Has(tag.InstrmtAssignmentMethod)
}

//HasMaturityTime returns true if MaturityTime is present, Tag 1079
func (m IOI) HasMaturityTime() bool {
	return m.Has(tag.MaturityTime)
}

//NoIOIQualifiers is a repeating group element, Tag 199
type NoIOIQualifiers struct {
	quickfix.Group
}

//SetIOIQualifier sets IOIQualifier, Tag 104
func (m NoIOIQualifiers) SetIOIQualifier(v string) {
	m.Set(field.NewIOIQualifier(v))
}

//GetIOIQualifier gets IOIQualifier, Tag 104
func (m NoIOIQualifiers) GetIOIQualifier() (f field.IOIQualifierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasIOIQualifier returns true if IOIQualifier is present, Tag 104
func (m NoIOIQualifiers) HasIOIQualifier() bool {
	return m.Has(tag.IOIQualifier)
}

//NoIOIQualifiersRepeatingGroup is a repeating group, Tag 199
type NoIOIQualifiersRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoIOIQualifiersRepeatingGroup returns an initialized, NoIOIQualifiersRepeatingGroup
func NewNoIOIQualifiersRepeatingGroup() NoIOIQualifiersRepeatingGroup {
	return NoIOIQualifiersRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoIOIQualifiers,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.IOIQualifier)})}
}

//Add create and append a new NoIOIQualifiers to this group
func (m NoIOIQualifiersRepeatingGroup) Add() NoIOIQualifiers {
	g := m.RepeatingGroup.Add()
	return NoIOIQualifiers{g}
}

//Get returns the ith NoIOIQualifiers in the NoIOIQualifiersRepeatinGroup
func (m NoIOIQualifiersRepeatingGroup) Get(i int) NoIOIQualifiers {
	return NoIOIQualifiers{m.RepeatingGroup.Get(i)}
}

//NoRoutingIDs is a repeating group element, Tag 215
type NoRoutingIDs struct {
	quickfix.Group
}

//SetRoutingType sets RoutingType, Tag 216
func (m NoRoutingIDs) SetRoutingType(v int) {
	m.Set(field.NewRoutingType(v))
}

//SetRoutingID sets RoutingID, Tag 217
func (m NoRoutingIDs) SetRoutingID(v string) {
	m.Set(field.NewRoutingID(v))
}

//GetRoutingType gets RoutingType, Tag 216
func (m NoRoutingIDs) GetRoutingType() (f field.RoutingTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoutingID gets RoutingID, Tag 217
func (m NoRoutingIDs) GetRoutingID() (f field.RoutingIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasRoutingType returns true if RoutingType is present, Tag 216
func (m NoRoutingIDs) HasRoutingType() bool {
	return m.Has(tag.RoutingType)
}

//HasRoutingID returns true if RoutingID is present, Tag 217
func (m NoRoutingIDs) HasRoutingID() bool {
	return m.Has(tag.RoutingID)
}

//NoRoutingIDsRepeatingGroup is a repeating group, Tag 215
type NoRoutingIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRoutingIDsRepeatingGroup returns an initialized, NoRoutingIDsRepeatingGroup
func NewNoRoutingIDsRepeatingGroup() NoRoutingIDsRepeatingGroup {
	return NoRoutingIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRoutingIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RoutingType), quickfix.GroupElement(tag.RoutingID)})}
}

//Add create and append a new NoRoutingIDs to this group
func (m NoRoutingIDsRepeatingGroup) Add() NoRoutingIDs {
	g := m.RepeatingGroup.Add()
	return NoRoutingIDs{g}
}

//Get returns the ith NoRoutingIDs in the NoRoutingIDsRepeatinGroup
func (m NoRoutingIDsRepeatingGroup) Get(i int) NoRoutingIDs {
	return NoRoutingIDs{m.RepeatingGroup.Get(i)}
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

//SetNoPartySubIDs sets NoPartySubIDs, Tag 802
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
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

//GetNoPartySubIDs gets NoPartySubIDs, Tag 802
func (m NoPartyIDs) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
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

//HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802
func (m NoPartyIDs) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

//NoPartySubIDs is a repeating group element, Tag 802
type NoPartySubIDs struct {
	quickfix.Group
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//SetPartySubIDType sets PartySubIDType, Tag 803
func (m NoPartySubIDs) SetPartySubIDType(v int) {
	m.Set(field.NewPartySubIDType(v))
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartySubIDs) GetPartySubID() (f field.PartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartySubIDType gets PartySubIDType, Tag 803
func (m NoPartySubIDs) GetPartySubIDType() (f field.PartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasPartySubID returns true if PartySubID is present, Tag 523
func (m NoPartySubIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

//HasPartySubIDType returns true if PartySubIDType is present, Tag 803
func (m NoPartySubIDs) HasPartySubIDType() bool {
	return m.Has(tag.PartySubIDType)
}

//NoPartySubIDsRepeatingGroup is a repeating group, Tag 802
type NoPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartySubIDsRepeatingGroup returns an initialized, NoPartySubIDsRepeatingGroup
func NewNoPartySubIDsRepeatingGroup() NoPartySubIDsRepeatingGroup {
	return NoPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartySubID), quickfix.GroupElement(tag.PartySubIDType)})}
}

//Add create and append a new NoPartySubIDs to this group
func (m NoPartySubIDsRepeatingGroup) Add() NoPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartySubIDs{g}
}

//Get returns the ith NoPartySubIDs in the NoPartySubIDsRepeatinGroup
func (m NoPartySubIDsRepeatingGroup) Get(i int) NoPartySubIDs {
	return NoPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoPartyIDsRepeatingGroup is a repeating group, Tag 453
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), quickfix.GroupElement(tag.NoPartySubIDs)})}
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

//NoLegs is a repeating group element, Tag 555
type NoLegs struct {
	quickfix.Group
}

//SetLegSymbol sets LegSymbol, Tag 600
func (m NoLegs) SetLegSymbol(v string) {
	m.Set(field.NewLegSymbol(v))
}

//SetLegSymbolSfx sets LegSymbolSfx, Tag 601
func (m NoLegs) SetLegSymbolSfx(v string) {
	m.Set(field.NewLegSymbolSfx(v))
}

//SetLegSecurityID sets LegSecurityID, Tag 602
func (m NoLegs) SetLegSecurityID(v string) {
	m.Set(field.NewLegSecurityID(v))
}

//SetLegSecurityIDSource sets LegSecurityIDSource, Tag 603
func (m NoLegs) SetLegSecurityIDSource(v string) {
	m.Set(field.NewLegSecurityIDSource(v))
}

//SetNoLegSecurityAltID sets NoLegSecurityAltID, Tag 604
func (m NoLegs) SetNoLegSecurityAltID(f NoLegSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetLegProduct sets LegProduct, Tag 607
func (m NoLegs) SetLegProduct(v int) {
	m.Set(field.NewLegProduct(v))
}

//SetLegCFICode sets LegCFICode, Tag 608
func (m NoLegs) SetLegCFICode(v string) {
	m.Set(field.NewLegCFICode(v))
}

//SetLegSecurityType sets LegSecurityType, Tag 609
func (m NoLegs) SetLegSecurityType(v string) {
	m.Set(field.NewLegSecurityType(v))
}

//SetLegSecuritySubType sets LegSecuritySubType, Tag 764
func (m NoLegs) SetLegSecuritySubType(v string) {
	m.Set(field.NewLegSecuritySubType(v))
}

//SetLegMaturityMonthYear sets LegMaturityMonthYear, Tag 610
func (m NoLegs) SetLegMaturityMonthYear(v string) {
	m.Set(field.NewLegMaturityMonthYear(v))
}

//SetLegMaturityDate sets LegMaturityDate, Tag 611
func (m NoLegs) SetLegMaturityDate(v string) {
	m.Set(field.NewLegMaturityDate(v))
}

//SetLegCouponPaymentDate sets LegCouponPaymentDate, Tag 248
func (m NoLegs) SetLegCouponPaymentDate(v string) {
	m.Set(field.NewLegCouponPaymentDate(v))
}

//SetLegIssueDate sets LegIssueDate, Tag 249
func (m NoLegs) SetLegIssueDate(v string) {
	m.Set(field.NewLegIssueDate(v))
}

//SetLegRepoCollateralSecurityType sets LegRepoCollateralSecurityType, Tag 250
func (m NoLegs) SetLegRepoCollateralSecurityType(v int) {
	m.Set(field.NewLegRepoCollateralSecurityType(v))
}

//SetLegRepurchaseTerm sets LegRepurchaseTerm, Tag 251
func (m NoLegs) SetLegRepurchaseTerm(v int) {
	m.Set(field.NewLegRepurchaseTerm(v))
}

//SetLegRepurchaseRate sets LegRepurchaseRate, Tag 252
func (m NoLegs) SetLegRepurchaseRate(v float64) {
	m.Set(field.NewLegRepurchaseRate(v))
}

//SetLegFactor sets LegFactor, Tag 253
func (m NoLegs) SetLegFactor(v float64) {
	m.Set(field.NewLegFactor(v))
}

//SetLegCreditRating sets LegCreditRating, Tag 257
func (m NoLegs) SetLegCreditRating(v string) {
	m.Set(field.NewLegCreditRating(v))
}

//SetLegInstrRegistry sets LegInstrRegistry, Tag 599
func (m NoLegs) SetLegInstrRegistry(v string) {
	m.Set(field.NewLegInstrRegistry(v))
}

//SetLegCountryOfIssue sets LegCountryOfIssue, Tag 596
func (m NoLegs) SetLegCountryOfIssue(v string) {
	m.Set(field.NewLegCountryOfIssue(v))
}

//SetLegStateOrProvinceOfIssue sets LegStateOrProvinceOfIssue, Tag 597
func (m NoLegs) SetLegStateOrProvinceOfIssue(v string) {
	m.Set(field.NewLegStateOrProvinceOfIssue(v))
}

//SetLegLocaleOfIssue sets LegLocaleOfIssue, Tag 598
func (m NoLegs) SetLegLocaleOfIssue(v string) {
	m.Set(field.NewLegLocaleOfIssue(v))
}

//SetLegRedemptionDate sets LegRedemptionDate, Tag 254
func (m NoLegs) SetLegRedemptionDate(v string) {
	m.Set(field.NewLegRedemptionDate(v))
}

//SetLegStrikePrice sets LegStrikePrice, Tag 612
func (m NoLegs) SetLegStrikePrice(v float64) {
	m.Set(field.NewLegStrikePrice(v))
}

//SetLegStrikeCurrency sets LegStrikeCurrency, Tag 942
func (m NoLegs) SetLegStrikeCurrency(v string) {
	m.Set(field.NewLegStrikeCurrency(v))
}

//SetLegOptAttribute sets LegOptAttribute, Tag 613
func (m NoLegs) SetLegOptAttribute(v string) {
	m.Set(field.NewLegOptAttribute(v))
}

//SetLegContractMultiplier sets LegContractMultiplier, Tag 614
func (m NoLegs) SetLegContractMultiplier(v float64) {
	m.Set(field.NewLegContractMultiplier(v))
}

//SetLegCouponRate sets LegCouponRate, Tag 615
func (m NoLegs) SetLegCouponRate(v float64) {
	m.Set(field.NewLegCouponRate(v))
}

//SetLegSecurityExchange sets LegSecurityExchange, Tag 616
func (m NoLegs) SetLegSecurityExchange(v string) {
	m.Set(field.NewLegSecurityExchange(v))
}

//SetLegIssuer sets LegIssuer, Tag 617
func (m NoLegs) SetLegIssuer(v string) {
	m.Set(field.NewLegIssuer(v))
}

//SetEncodedLegIssuerLen sets EncodedLegIssuerLen, Tag 618
func (m NoLegs) SetEncodedLegIssuerLen(v int) {
	m.Set(field.NewEncodedLegIssuerLen(v))
}

//SetEncodedLegIssuer sets EncodedLegIssuer, Tag 619
func (m NoLegs) SetEncodedLegIssuer(v string) {
	m.Set(field.NewEncodedLegIssuer(v))
}

//SetLegSecurityDesc sets LegSecurityDesc, Tag 620
func (m NoLegs) SetLegSecurityDesc(v string) {
	m.Set(field.NewLegSecurityDesc(v))
}

//SetEncodedLegSecurityDescLen sets EncodedLegSecurityDescLen, Tag 621
func (m NoLegs) SetEncodedLegSecurityDescLen(v int) {
	m.Set(field.NewEncodedLegSecurityDescLen(v))
}

//SetEncodedLegSecurityDesc sets EncodedLegSecurityDesc, Tag 622
func (m NoLegs) SetEncodedLegSecurityDesc(v string) {
	m.Set(field.NewEncodedLegSecurityDesc(v))
}

//SetLegRatioQty sets LegRatioQty, Tag 623
func (m NoLegs) SetLegRatioQty(v float64) {
	m.Set(field.NewLegRatioQty(v))
}

//SetLegSide sets LegSide, Tag 624
func (m NoLegs) SetLegSide(v string) {
	m.Set(field.NewLegSide(v))
}

//SetLegCurrency sets LegCurrency, Tag 556
func (m NoLegs) SetLegCurrency(v string) {
	m.Set(field.NewLegCurrency(v))
}

//SetLegPool sets LegPool, Tag 740
func (m NoLegs) SetLegPool(v string) {
	m.Set(field.NewLegPool(v))
}

//SetLegDatedDate sets LegDatedDate, Tag 739
func (m NoLegs) SetLegDatedDate(v string) {
	m.Set(field.NewLegDatedDate(v))
}

//SetLegContractSettlMonth sets LegContractSettlMonth, Tag 955
func (m NoLegs) SetLegContractSettlMonth(v string) {
	m.Set(field.NewLegContractSettlMonth(v))
}

//SetLegInterestAccrualDate sets LegInterestAccrualDate, Tag 956
func (m NoLegs) SetLegInterestAccrualDate(v string) {
	m.Set(field.NewLegInterestAccrualDate(v))
}

//SetLegUnitOfMeasure sets LegUnitOfMeasure, Tag 999
func (m NoLegs) SetLegUnitOfMeasure(v string) {
	m.Set(field.NewLegUnitOfMeasure(v))
}

//SetLegTimeUnit sets LegTimeUnit, Tag 1001
func (m NoLegs) SetLegTimeUnit(v string) {
	m.Set(field.NewLegTimeUnit(v))
}

//SetLegIOIQty sets LegIOIQty, Tag 682
func (m NoLegs) SetLegIOIQty(v string) {
	m.Set(field.NewLegIOIQty(v))
}

//SetNoLegStipulations sets NoLegStipulations, Tag 683
func (m NoLegs) SetNoLegStipulations(f NoLegStipulationsRepeatingGroup) {
	m.SetGroup(f)
}

//GetLegSymbol gets LegSymbol, Tag 600
func (m NoLegs) GetLegSymbol() (f field.LegSymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSymbolSfx gets LegSymbolSfx, Tag 601
func (m NoLegs) GetLegSymbolSfx() (f field.LegSymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityID gets LegSecurityID, Tag 602
func (m NoLegs) GetLegSecurityID() (f field.LegSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityIDSource gets LegSecurityIDSource, Tag 603
func (m NoLegs) GetLegSecurityIDSource() (f field.LegSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoLegSecurityAltID gets NoLegSecurityAltID, Tag 604
func (m NoLegs) GetNoLegSecurityAltID() (NoLegSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLegProduct gets LegProduct, Tag 607
func (m NoLegs) GetLegProduct() (f field.LegProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCFICode gets LegCFICode, Tag 608
func (m NoLegs) GetLegCFICode() (f field.LegCFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityType gets LegSecurityType, Tag 609
func (m NoLegs) GetLegSecurityType() (f field.LegSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecuritySubType gets LegSecuritySubType, Tag 764
func (m NoLegs) GetLegSecuritySubType() (f field.LegSecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegMaturityMonthYear gets LegMaturityMonthYear, Tag 610
func (m NoLegs) GetLegMaturityMonthYear() (f field.LegMaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegMaturityDate gets LegMaturityDate, Tag 611
func (m NoLegs) GetLegMaturityDate() (f field.LegMaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCouponPaymentDate gets LegCouponPaymentDate, Tag 248
func (m NoLegs) GetLegCouponPaymentDate() (f field.LegCouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegIssueDate gets LegIssueDate, Tag 249
func (m NoLegs) GetLegIssueDate() (f field.LegIssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRepoCollateralSecurityType gets LegRepoCollateralSecurityType, Tag 250
func (m NoLegs) GetLegRepoCollateralSecurityType() (f field.LegRepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRepurchaseTerm gets LegRepurchaseTerm, Tag 251
func (m NoLegs) GetLegRepurchaseTerm() (f field.LegRepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRepurchaseRate gets LegRepurchaseRate, Tag 252
func (m NoLegs) GetLegRepurchaseRate() (f field.LegRepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegFactor gets LegFactor, Tag 253
func (m NoLegs) GetLegFactor() (f field.LegFactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCreditRating gets LegCreditRating, Tag 257
func (m NoLegs) GetLegCreditRating() (f field.LegCreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegInstrRegistry gets LegInstrRegistry, Tag 599
func (m NoLegs) GetLegInstrRegistry() (f field.LegInstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCountryOfIssue gets LegCountryOfIssue, Tag 596
func (m NoLegs) GetLegCountryOfIssue() (f field.LegCountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegStateOrProvinceOfIssue gets LegStateOrProvinceOfIssue, Tag 597
func (m NoLegs) GetLegStateOrProvinceOfIssue() (f field.LegStateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegLocaleOfIssue gets LegLocaleOfIssue, Tag 598
func (m NoLegs) GetLegLocaleOfIssue() (f field.LegLocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRedemptionDate gets LegRedemptionDate, Tag 254
func (m NoLegs) GetLegRedemptionDate() (f field.LegRedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegStrikePrice gets LegStrikePrice, Tag 612
func (m NoLegs) GetLegStrikePrice() (f field.LegStrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegStrikeCurrency gets LegStrikeCurrency, Tag 942
func (m NoLegs) GetLegStrikeCurrency() (f field.LegStrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegOptAttribute gets LegOptAttribute, Tag 613
func (m NoLegs) GetLegOptAttribute() (f field.LegOptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegContractMultiplier gets LegContractMultiplier, Tag 614
func (m NoLegs) GetLegContractMultiplier() (f field.LegContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCouponRate gets LegCouponRate, Tag 615
func (m NoLegs) GetLegCouponRate() (f field.LegCouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityExchange gets LegSecurityExchange, Tag 616
func (m NoLegs) GetLegSecurityExchange() (f field.LegSecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegIssuer gets LegIssuer, Tag 617
func (m NoLegs) GetLegIssuer() (f field.LegIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedLegIssuerLen gets EncodedLegIssuerLen, Tag 618
func (m NoLegs) GetEncodedLegIssuerLen() (f field.EncodedLegIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedLegIssuer gets EncodedLegIssuer, Tag 619
func (m NoLegs) GetEncodedLegIssuer() (f field.EncodedLegIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityDesc gets LegSecurityDesc, Tag 620
func (m NoLegs) GetLegSecurityDesc() (f field.LegSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedLegSecurityDescLen gets EncodedLegSecurityDescLen, Tag 621
func (m NoLegs) GetEncodedLegSecurityDescLen() (f field.EncodedLegSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedLegSecurityDesc gets EncodedLegSecurityDesc, Tag 622
func (m NoLegs) GetEncodedLegSecurityDesc() (f field.EncodedLegSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRatioQty gets LegRatioQty, Tag 623
func (m NoLegs) GetLegRatioQty() (f field.LegRatioQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSide gets LegSide, Tag 624
func (m NoLegs) GetLegSide() (f field.LegSideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCurrency gets LegCurrency, Tag 556
func (m NoLegs) GetLegCurrency() (f field.LegCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPool gets LegPool, Tag 740
func (m NoLegs) GetLegPool() (f field.LegPoolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegDatedDate gets LegDatedDate, Tag 739
func (m NoLegs) GetLegDatedDate() (f field.LegDatedDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegContractSettlMonth gets LegContractSettlMonth, Tag 955
func (m NoLegs) GetLegContractSettlMonth() (f field.LegContractSettlMonthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegInterestAccrualDate gets LegInterestAccrualDate, Tag 956
func (m NoLegs) GetLegInterestAccrualDate() (f field.LegInterestAccrualDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegUnitOfMeasure gets LegUnitOfMeasure, Tag 999
func (m NoLegs) GetLegUnitOfMeasure() (f field.LegUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegTimeUnit gets LegTimeUnit, Tag 1001
func (m NoLegs) GetLegTimeUnit() (f field.LegTimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegIOIQty gets LegIOIQty, Tag 682
func (m NoLegs) GetLegIOIQty() (f field.LegIOIQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoLegStipulations gets NoLegStipulations, Tag 683
func (m NoLegs) GetNoLegStipulations() (NoLegStipulationsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegStipulationsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasLegSymbol returns true if LegSymbol is present, Tag 600
func (m NoLegs) HasLegSymbol() bool {
	return m.Has(tag.LegSymbol)
}

//HasLegSymbolSfx returns true if LegSymbolSfx is present, Tag 601
func (m NoLegs) HasLegSymbolSfx() bool {
	return m.Has(tag.LegSymbolSfx)
}

//HasLegSecurityID returns true if LegSecurityID is present, Tag 602
func (m NoLegs) HasLegSecurityID() bool {
	return m.Has(tag.LegSecurityID)
}

//HasLegSecurityIDSource returns true if LegSecurityIDSource is present, Tag 603
func (m NoLegs) HasLegSecurityIDSource() bool {
	return m.Has(tag.LegSecurityIDSource)
}

//HasNoLegSecurityAltID returns true if NoLegSecurityAltID is present, Tag 604
func (m NoLegs) HasNoLegSecurityAltID() bool {
	return m.Has(tag.NoLegSecurityAltID)
}

//HasLegProduct returns true if LegProduct is present, Tag 607
func (m NoLegs) HasLegProduct() bool {
	return m.Has(tag.LegProduct)
}

//HasLegCFICode returns true if LegCFICode is present, Tag 608
func (m NoLegs) HasLegCFICode() bool {
	return m.Has(tag.LegCFICode)
}

//HasLegSecurityType returns true if LegSecurityType is present, Tag 609
func (m NoLegs) HasLegSecurityType() bool {
	return m.Has(tag.LegSecurityType)
}

//HasLegSecuritySubType returns true if LegSecuritySubType is present, Tag 764
func (m NoLegs) HasLegSecuritySubType() bool {
	return m.Has(tag.LegSecuritySubType)
}

//HasLegMaturityMonthYear returns true if LegMaturityMonthYear is present, Tag 610
func (m NoLegs) HasLegMaturityMonthYear() bool {
	return m.Has(tag.LegMaturityMonthYear)
}

//HasLegMaturityDate returns true if LegMaturityDate is present, Tag 611
func (m NoLegs) HasLegMaturityDate() bool {
	return m.Has(tag.LegMaturityDate)
}

//HasLegCouponPaymentDate returns true if LegCouponPaymentDate is present, Tag 248
func (m NoLegs) HasLegCouponPaymentDate() bool {
	return m.Has(tag.LegCouponPaymentDate)
}

//HasLegIssueDate returns true if LegIssueDate is present, Tag 249
func (m NoLegs) HasLegIssueDate() bool {
	return m.Has(tag.LegIssueDate)
}

//HasLegRepoCollateralSecurityType returns true if LegRepoCollateralSecurityType is present, Tag 250
func (m NoLegs) HasLegRepoCollateralSecurityType() bool {
	return m.Has(tag.LegRepoCollateralSecurityType)
}

//HasLegRepurchaseTerm returns true if LegRepurchaseTerm is present, Tag 251
func (m NoLegs) HasLegRepurchaseTerm() bool {
	return m.Has(tag.LegRepurchaseTerm)
}

//HasLegRepurchaseRate returns true if LegRepurchaseRate is present, Tag 252
func (m NoLegs) HasLegRepurchaseRate() bool {
	return m.Has(tag.LegRepurchaseRate)
}

//HasLegFactor returns true if LegFactor is present, Tag 253
func (m NoLegs) HasLegFactor() bool {
	return m.Has(tag.LegFactor)
}

//HasLegCreditRating returns true if LegCreditRating is present, Tag 257
func (m NoLegs) HasLegCreditRating() bool {
	return m.Has(tag.LegCreditRating)
}

//HasLegInstrRegistry returns true if LegInstrRegistry is present, Tag 599
func (m NoLegs) HasLegInstrRegistry() bool {
	return m.Has(tag.LegInstrRegistry)
}

//HasLegCountryOfIssue returns true if LegCountryOfIssue is present, Tag 596
func (m NoLegs) HasLegCountryOfIssue() bool {
	return m.Has(tag.LegCountryOfIssue)
}

//HasLegStateOrProvinceOfIssue returns true if LegStateOrProvinceOfIssue is present, Tag 597
func (m NoLegs) HasLegStateOrProvinceOfIssue() bool {
	return m.Has(tag.LegStateOrProvinceOfIssue)
}

//HasLegLocaleOfIssue returns true if LegLocaleOfIssue is present, Tag 598
func (m NoLegs) HasLegLocaleOfIssue() bool {
	return m.Has(tag.LegLocaleOfIssue)
}

//HasLegRedemptionDate returns true if LegRedemptionDate is present, Tag 254
func (m NoLegs) HasLegRedemptionDate() bool {
	return m.Has(tag.LegRedemptionDate)
}

//HasLegStrikePrice returns true if LegStrikePrice is present, Tag 612
func (m NoLegs) HasLegStrikePrice() bool {
	return m.Has(tag.LegStrikePrice)
}

//HasLegStrikeCurrency returns true if LegStrikeCurrency is present, Tag 942
func (m NoLegs) HasLegStrikeCurrency() bool {
	return m.Has(tag.LegStrikeCurrency)
}

//HasLegOptAttribute returns true if LegOptAttribute is present, Tag 613
func (m NoLegs) HasLegOptAttribute() bool {
	return m.Has(tag.LegOptAttribute)
}

//HasLegContractMultiplier returns true if LegContractMultiplier is present, Tag 614
func (m NoLegs) HasLegContractMultiplier() bool {
	return m.Has(tag.LegContractMultiplier)
}

//HasLegCouponRate returns true if LegCouponRate is present, Tag 615
func (m NoLegs) HasLegCouponRate() bool {
	return m.Has(tag.LegCouponRate)
}

//HasLegSecurityExchange returns true if LegSecurityExchange is present, Tag 616
func (m NoLegs) HasLegSecurityExchange() bool {
	return m.Has(tag.LegSecurityExchange)
}

//HasLegIssuer returns true if LegIssuer is present, Tag 617
func (m NoLegs) HasLegIssuer() bool {
	return m.Has(tag.LegIssuer)
}

//HasEncodedLegIssuerLen returns true if EncodedLegIssuerLen is present, Tag 618
func (m NoLegs) HasEncodedLegIssuerLen() bool {
	return m.Has(tag.EncodedLegIssuerLen)
}

//HasEncodedLegIssuer returns true if EncodedLegIssuer is present, Tag 619
func (m NoLegs) HasEncodedLegIssuer() bool {
	return m.Has(tag.EncodedLegIssuer)
}

//HasLegSecurityDesc returns true if LegSecurityDesc is present, Tag 620
func (m NoLegs) HasLegSecurityDesc() bool {
	return m.Has(tag.LegSecurityDesc)
}

//HasEncodedLegSecurityDescLen returns true if EncodedLegSecurityDescLen is present, Tag 621
func (m NoLegs) HasEncodedLegSecurityDescLen() bool {
	return m.Has(tag.EncodedLegSecurityDescLen)
}

//HasEncodedLegSecurityDesc returns true if EncodedLegSecurityDesc is present, Tag 622
func (m NoLegs) HasEncodedLegSecurityDesc() bool {
	return m.Has(tag.EncodedLegSecurityDesc)
}

//HasLegRatioQty returns true if LegRatioQty is present, Tag 623
func (m NoLegs) HasLegRatioQty() bool {
	return m.Has(tag.LegRatioQty)
}

//HasLegSide returns true if LegSide is present, Tag 624
func (m NoLegs) HasLegSide() bool {
	return m.Has(tag.LegSide)
}

//HasLegCurrency returns true if LegCurrency is present, Tag 556
func (m NoLegs) HasLegCurrency() bool {
	return m.Has(tag.LegCurrency)
}

//HasLegPool returns true if LegPool is present, Tag 740
func (m NoLegs) HasLegPool() bool {
	return m.Has(tag.LegPool)
}

//HasLegDatedDate returns true if LegDatedDate is present, Tag 739
func (m NoLegs) HasLegDatedDate() bool {
	return m.Has(tag.LegDatedDate)
}

//HasLegContractSettlMonth returns true if LegContractSettlMonth is present, Tag 955
func (m NoLegs) HasLegContractSettlMonth() bool {
	return m.Has(tag.LegContractSettlMonth)
}

//HasLegInterestAccrualDate returns true if LegInterestAccrualDate is present, Tag 956
func (m NoLegs) HasLegInterestAccrualDate() bool {
	return m.Has(tag.LegInterestAccrualDate)
}

//HasLegUnitOfMeasure returns true if LegUnitOfMeasure is present, Tag 999
func (m NoLegs) HasLegUnitOfMeasure() bool {
	return m.Has(tag.LegUnitOfMeasure)
}

//HasLegTimeUnit returns true if LegTimeUnit is present, Tag 1001
func (m NoLegs) HasLegTimeUnit() bool {
	return m.Has(tag.LegTimeUnit)
}

//HasLegIOIQty returns true if LegIOIQty is present, Tag 682
func (m NoLegs) HasLegIOIQty() bool {
	return m.Has(tag.LegIOIQty)
}

//HasNoLegStipulations returns true if NoLegStipulations is present, Tag 683
func (m NoLegs) HasNoLegStipulations() bool {
	return m.Has(tag.NoLegStipulations)
}

//NoLegSecurityAltID is a repeating group element, Tag 604
type NoLegSecurityAltID struct {
	quickfix.Group
}

//SetLegSecurityAltID sets LegSecurityAltID, Tag 605
func (m NoLegSecurityAltID) SetLegSecurityAltID(v string) {
	m.Set(field.NewLegSecurityAltID(v))
}

//SetLegSecurityAltIDSource sets LegSecurityAltIDSource, Tag 606
func (m NoLegSecurityAltID) SetLegSecurityAltIDSource(v string) {
	m.Set(field.NewLegSecurityAltIDSource(v))
}

//GetLegSecurityAltID gets LegSecurityAltID, Tag 605
func (m NoLegSecurityAltID) GetLegSecurityAltID() (f field.LegSecurityAltIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityAltIDSource gets LegSecurityAltIDSource, Tag 606
func (m NoLegSecurityAltID) GetLegSecurityAltIDSource() (f field.LegSecurityAltIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasLegSecurityAltID returns true if LegSecurityAltID is present, Tag 605
func (m NoLegSecurityAltID) HasLegSecurityAltID() bool {
	return m.Has(tag.LegSecurityAltID)
}

//HasLegSecurityAltIDSource returns true if LegSecurityAltIDSource is present, Tag 606
func (m NoLegSecurityAltID) HasLegSecurityAltIDSource() bool {
	return m.Has(tag.LegSecurityAltIDSource)
}

//NoLegSecurityAltIDRepeatingGroup is a repeating group, Tag 604
type NoLegSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegSecurityAltIDRepeatingGroup returns an initialized, NoLegSecurityAltIDRepeatingGroup
func NewNoLegSecurityAltIDRepeatingGroup() NoLegSecurityAltIDRepeatingGroup {
	return NoLegSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSecurityAltID), quickfix.GroupElement(tag.LegSecurityAltIDSource)})}
}

//Add create and append a new NoLegSecurityAltID to this group
func (m NoLegSecurityAltIDRepeatingGroup) Add() NoLegSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoLegSecurityAltID{g}
}

//Get returns the ith NoLegSecurityAltID in the NoLegSecurityAltIDRepeatinGroup
func (m NoLegSecurityAltIDRepeatingGroup) Get(i int) NoLegSecurityAltID {
	return NoLegSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoLegStipulations is a repeating group element, Tag 683
type NoLegStipulations struct {
	quickfix.Group
}

//SetLegStipulationType sets LegStipulationType, Tag 688
func (m NoLegStipulations) SetLegStipulationType(v string) {
	m.Set(field.NewLegStipulationType(v))
}

//SetLegStipulationValue sets LegStipulationValue, Tag 689
func (m NoLegStipulations) SetLegStipulationValue(v string) {
	m.Set(field.NewLegStipulationValue(v))
}

//GetLegStipulationType gets LegStipulationType, Tag 688
func (m NoLegStipulations) GetLegStipulationType() (f field.LegStipulationTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegStipulationValue gets LegStipulationValue, Tag 689
func (m NoLegStipulations) GetLegStipulationValue() (f field.LegStipulationValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasLegStipulationType returns true if LegStipulationType is present, Tag 688
func (m NoLegStipulations) HasLegStipulationType() bool {
	return m.Has(tag.LegStipulationType)
}

//HasLegStipulationValue returns true if LegStipulationValue is present, Tag 689
func (m NoLegStipulations) HasLegStipulationValue() bool {
	return m.Has(tag.LegStipulationValue)
}

//NoLegStipulationsRepeatingGroup is a repeating group, Tag 683
type NoLegStipulationsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegStipulationsRepeatingGroup returns an initialized, NoLegStipulationsRepeatingGroup
func NewNoLegStipulationsRepeatingGroup() NoLegStipulationsRepeatingGroup {
	return NoLegStipulationsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegStipulations,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegStipulationType), quickfix.GroupElement(tag.LegStipulationValue)})}
}

//Add create and append a new NoLegStipulations to this group
func (m NoLegStipulationsRepeatingGroup) Add() NoLegStipulations {
	g := m.RepeatingGroup.Add()
	return NoLegStipulations{g}
}

//Get returns the ith NoLegStipulations in the NoLegStipulationsRepeatinGroup
func (m NoLegStipulationsRepeatingGroup) Get(i int) NoLegStipulations {
	return NoLegStipulations{m.RepeatingGroup.Get(i)}
}

//NoLegsRepeatingGroup is a repeating group, Tag 555
type NoLegsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegsRepeatingGroup returns an initialized, NoLegsRepeatingGroup
func NewNoLegsRepeatingGroup() NoLegsRepeatingGroup {
	return NoLegsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSymbol), quickfix.GroupElement(tag.LegSymbolSfx), quickfix.GroupElement(tag.LegSecurityID), quickfix.GroupElement(tag.LegSecurityIDSource), quickfix.GroupElement(tag.NoLegSecurityAltID), quickfix.GroupElement(tag.LegProduct), quickfix.GroupElement(tag.LegCFICode), quickfix.GroupElement(tag.LegSecurityType), quickfix.GroupElement(tag.LegSecuritySubType), quickfix.GroupElement(tag.LegMaturityMonthYear), quickfix.GroupElement(tag.LegMaturityDate), quickfix.GroupElement(tag.LegCouponPaymentDate), quickfix.GroupElement(tag.LegIssueDate), quickfix.GroupElement(tag.LegRepoCollateralSecurityType), quickfix.GroupElement(tag.LegRepurchaseTerm), quickfix.GroupElement(tag.LegRepurchaseRate), quickfix.GroupElement(tag.LegFactor), quickfix.GroupElement(tag.LegCreditRating), quickfix.GroupElement(tag.LegInstrRegistry), quickfix.GroupElement(tag.LegCountryOfIssue), quickfix.GroupElement(tag.LegStateOrProvinceOfIssue), quickfix.GroupElement(tag.LegLocaleOfIssue), quickfix.GroupElement(tag.LegRedemptionDate), quickfix.GroupElement(tag.LegStrikePrice), quickfix.GroupElement(tag.LegStrikeCurrency), quickfix.GroupElement(tag.LegOptAttribute), quickfix.GroupElement(tag.LegContractMultiplier), quickfix.GroupElement(tag.LegCouponRate), quickfix.GroupElement(tag.LegSecurityExchange), quickfix.GroupElement(tag.LegIssuer), quickfix.GroupElement(tag.EncodedLegIssuerLen), quickfix.GroupElement(tag.EncodedLegIssuer), quickfix.GroupElement(tag.LegSecurityDesc), quickfix.GroupElement(tag.EncodedLegSecurityDescLen), quickfix.GroupElement(tag.EncodedLegSecurityDesc), quickfix.GroupElement(tag.LegRatioQty), quickfix.GroupElement(tag.LegSide), quickfix.GroupElement(tag.LegCurrency), quickfix.GroupElement(tag.LegPool), quickfix.GroupElement(tag.LegDatedDate), quickfix.GroupElement(tag.LegContractSettlMonth), quickfix.GroupElement(tag.LegInterestAccrualDate), quickfix.GroupElement(tag.LegUnitOfMeasure), quickfix.GroupElement(tag.LegTimeUnit), quickfix.GroupElement(tag.LegIOIQty), quickfix.GroupElement(tag.NoLegStipulations)})}
}

//Add create and append a new NoLegs to this group
func (m NoLegsRepeatingGroup) Add() NoLegs {
	g := m.RepeatingGroup.Add()
	return NoLegs{g}
}

//Get returns the ith NoLegs in the NoLegsRepeatinGroup
func (m NoLegsRepeatingGroup) Get(i int) NoLegs {
	return NoLegs{m.RepeatingGroup.Get(i)}
}

//NoUnderlyings is a repeating group element, Tag 711
type NoUnderlyings struct {
	quickfix.Group
}

//SetUnderlyingSymbol sets UnderlyingSymbol, Tag 311
func (m NoUnderlyings) SetUnderlyingSymbol(v string) {
	m.Set(field.NewUnderlyingSymbol(v))
}

//SetUnderlyingSymbolSfx sets UnderlyingSymbolSfx, Tag 312
func (m NoUnderlyings) SetUnderlyingSymbolSfx(v string) {
	m.Set(field.NewUnderlyingSymbolSfx(v))
}

//SetUnderlyingSecurityID sets UnderlyingSecurityID, Tag 309
func (m NoUnderlyings) SetUnderlyingSecurityID(v string) {
	m.Set(field.NewUnderlyingSecurityID(v))
}

//SetUnderlyingSecurityIDSource sets UnderlyingSecurityIDSource, Tag 305
func (m NoUnderlyings) SetUnderlyingSecurityIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityIDSource(v))
}

//SetNoUnderlyingSecurityAltID sets NoUnderlyingSecurityAltID, Tag 457
func (m NoUnderlyings) SetNoUnderlyingSecurityAltID(f NoUnderlyingSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingProduct sets UnderlyingProduct, Tag 462
func (m NoUnderlyings) SetUnderlyingProduct(v int) {
	m.Set(field.NewUnderlyingProduct(v))
}

//SetUnderlyingCFICode sets UnderlyingCFICode, Tag 463
func (m NoUnderlyings) SetUnderlyingCFICode(v string) {
	m.Set(field.NewUnderlyingCFICode(v))
}

//SetUnderlyingSecurityType sets UnderlyingSecurityType, Tag 310
func (m NoUnderlyings) SetUnderlyingSecurityType(v string) {
	m.Set(field.NewUnderlyingSecurityType(v))
}

//SetUnderlyingSecuritySubType sets UnderlyingSecuritySubType, Tag 763
func (m NoUnderlyings) SetUnderlyingSecuritySubType(v string) {
	m.Set(field.NewUnderlyingSecuritySubType(v))
}

//SetUnderlyingMaturityMonthYear sets UnderlyingMaturityMonthYear, Tag 313
func (m NoUnderlyings) SetUnderlyingMaturityMonthYear(v string) {
	m.Set(field.NewUnderlyingMaturityMonthYear(v))
}

//SetUnderlyingMaturityDate sets UnderlyingMaturityDate, Tag 542
func (m NoUnderlyings) SetUnderlyingMaturityDate(v string) {
	m.Set(field.NewUnderlyingMaturityDate(v))
}

//SetUnderlyingCouponPaymentDate sets UnderlyingCouponPaymentDate, Tag 241
func (m NoUnderlyings) SetUnderlyingCouponPaymentDate(v string) {
	m.Set(field.NewUnderlyingCouponPaymentDate(v))
}

//SetUnderlyingIssueDate sets UnderlyingIssueDate, Tag 242
func (m NoUnderlyings) SetUnderlyingIssueDate(v string) {
	m.Set(field.NewUnderlyingIssueDate(v))
}

//SetUnderlyingRepoCollateralSecurityType sets UnderlyingRepoCollateralSecurityType, Tag 243
func (m NoUnderlyings) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.Set(field.NewUnderlyingRepoCollateralSecurityType(v))
}

//SetUnderlyingRepurchaseTerm sets UnderlyingRepurchaseTerm, Tag 244
func (m NoUnderlyings) SetUnderlyingRepurchaseTerm(v int) {
	m.Set(field.NewUnderlyingRepurchaseTerm(v))
}

//SetUnderlyingRepurchaseRate sets UnderlyingRepurchaseRate, Tag 245
func (m NoUnderlyings) SetUnderlyingRepurchaseRate(v float64) {
	m.Set(field.NewUnderlyingRepurchaseRate(v))
}

//SetUnderlyingFactor sets UnderlyingFactor, Tag 246
func (m NoUnderlyings) SetUnderlyingFactor(v float64) {
	m.Set(field.NewUnderlyingFactor(v))
}

//SetUnderlyingCreditRating sets UnderlyingCreditRating, Tag 256
func (m NoUnderlyings) SetUnderlyingCreditRating(v string) {
	m.Set(field.NewUnderlyingCreditRating(v))
}

//SetUnderlyingInstrRegistry sets UnderlyingInstrRegistry, Tag 595
func (m NoUnderlyings) SetUnderlyingInstrRegistry(v string) {
	m.Set(field.NewUnderlyingInstrRegistry(v))
}

//SetUnderlyingCountryOfIssue sets UnderlyingCountryOfIssue, Tag 592
func (m NoUnderlyings) SetUnderlyingCountryOfIssue(v string) {
	m.Set(field.NewUnderlyingCountryOfIssue(v))
}

//SetUnderlyingStateOrProvinceOfIssue sets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m NoUnderlyings) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.Set(field.NewUnderlyingStateOrProvinceOfIssue(v))
}

//SetUnderlyingLocaleOfIssue sets UnderlyingLocaleOfIssue, Tag 594
func (m NoUnderlyings) SetUnderlyingLocaleOfIssue(v string) {
	m.Set(field.NewUnderlyingLocaleOfIssue(v))
}

//SetUnderlyingRedemptionDate sets UnderlyingRedemptionDate, Tag 247
func (m NoUnderlyings) SetUnderlyingRedemptionDate(v string) {
	m.Set(field.NewUnderlyingRedemptionDate(v))
}

//SetUnderlyingStrikePrice sets UnderlyingStrikePrice, Tag 316
func (m NoUnderlyings) SetUnderlyingStrikePrice(v float64) {
	m.Set(field.NewUnderlyingStrikePrice(v))
}

//SetUnderlyingStrikeCurrency sets UnderlyingStrikeCurrency, Tag 941
func (m NoUnderlyings) SetUnderlyingStrikeCurrency(v string) {
	m.Set(field.NewUnderlyingStrikeCurrency(v))
}

//SetUnderlyingOptAttribute sets UnderlyingOptAttribute, Tag 317
func (m NoUnderlyings) SetUnderlyingOptAttribute(v string) {
	m.Set(field.NewUnderlyingOptAttribute(v))
}

//SetUnderlyingContractMultiplier sets UnderlyingContractMultiplier, Tag 436
func (m NoUnderlyings) SetUnderlyingContractMultiplier(v float64) {
	m.Set(field.NewUnderlyingContractMultiplier(v))
}

//SetUnderlyingCouponRate sets UnderlyingCouponRate, Tag 435
func (m NoUnderlyings) SetUnderlyingCouponRate(v float64) {
	m.Set(field.NewUnderlyingCouponRate(v))
}

//SetUnderlyingSecurityExchange sets UnderlyingSecurityExchange, Tag 308
func (m NoUnderlyings) SetUnderlyingSecurityExchange(v string) {
	m.Set(field.NewUnderlyingSecurityExchange(v))
}

//SetUnderlyingIssuer sets UnderlyingIssuer, Tag 306
func (m NoUnderlyings) SetUnderlyingIssuer(v string) {
	m.Set(field.NewUnderlyingIssuer(v))
}

//SetEncodedUnderlyingIssuerLen sets EncodedUnderlyingIssuerLen, Tag 362
func (m NoUnderlyings) SetEncodedUnderlyingIssuerLen(v int) {
	m.Set(field.NewEncodedUnderlyingIssuerLen(v))
}

//SetEncodedUnderlyingIssuer sets EncodedUnderlyingIssuer, Tag 363
func (m NoUnderlyings) SetEncodedUnderlyingIssuer(v string) {
	m.Set(field.NewEncodedUnderlyingIssuer(v))
}

//SetUnderlyingSecurityDesc sets UnderlyingSecurityDesc, Tag 307
func (m NoUnderlyings) SetUnderlyingSecurityDesc(v string) {
	m.Set(field.NewUnderlyingSecurityDesc(v))
}

//SetEncodedUnderlyingSecurityDescLen sets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.Set(field.NewEncodedUnderlyingSecurityDescLen(v))
}

//SetEncodedUnderlyingSecurityDesc sets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDesc(v string) {
	m.Set(field.NewEncodedUnderlyingSecurityDesc(v))
}

//SetUnderlyingCPProgram sets UnderlyingCPProgram, Tag 877
func (m NoUnderlyings) SetUnderlyingCPProgram(v string) {
	m.Set(field.NewUnderlyingCPProgram(v))
}

//SetUnderlyingCPRegType sets UnderlyingCPRegType, Tag 878
func (m NoUnderlyings) SetUnderlyingCPRegType(v string) {
	m.Set(field.NewUnderlyingCPRegType(v))
}

//SetUnderlyingCurrency sets UnderlyingCurrency, Tag 318
func (m NoUnderlyings) SetUnderlyingCurrency(v string) {
	m.Set(field.NewUnderlyingCurrency(v))
}

//SetUnderlyingQty sets UnderlyingQty, Tag 879
func (m NoUnderlyings) SetUnderlyingQty(v float64) {
	m.Set(field.NewUnderlyingQty(v))
}

//SetUnderlyingPx sets UnderlyingPx, Tag 810
func (m NoUnderlyings) SetUnderlyingPx(v float64) {
	m.Set(field.NewUnderlyingPx(v))
}

//SetUnderlyingDirtyPrice sets UnderlyingDirtyPrice, Tag 882
func (m NoUnderlyings) SetUnderlyingDirtyPrice(v float64) {
	m.Set(field.NewUnderlyingDirtyPrice(v))
}

//SetUnderlyingEndPrice sets UnderlyingEndPrice, Tag 883
func (m NoUnderlyings) SetUnderlyingEndPrice(v float64) {
	m.Set(field.NewUnderlyingEndPrice(v))
}

//SetUnderlyingStartValue sets UnderlyingStartValue, Tag 884
func (m NoUnderlyings) SetUnderlyingStartValue(v float64) {
	m.Set(field.NewUnderlyingStartValue(v))
}

//SetUnderlyingCurrentValue sets UnderlyingCurrentValue, Tag 885
func (m NoUnderlyings) SetUnderlyingCurrentValue(v float64) {
	m.Set(field.NewUnderlyingCurrentValue(v))
}

//SetUnderlyingEndValue sets UnderlyingEndValue, Tag 886
func (m NoUnderlyings) SetUnderlyingEndValue(v float64) {
	m.Set(field.NewUnderlyingEndValue(v))
}

//SetNoUnderlyingStips sets NoUnderlyingStips, Tag 887
func (m NoUnderlyings) SetNoUnderlyingStips(f NoUnderlyingStipsRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingAllocationPercent sets UnderlyingAllocationPercent, Tag 972
func (m NoUnderlyings) SetUnderlyingAllocationPercent(v float64) {
	m.Set(field.NewUnderlyingAllocationPercent(v))
}

//SetUnderlyingSettlementType sets UnderlyingSettlementType, Tag 975
func (m NoUnderlyings) SetUnderlyingSettlementType(v int) {
	m.Set(field.NewUnderlyingSettlementType(v))
}

//SetUnderlyingCashAmount sets UnderlyingCashAmount, Tag 973
func (m NoUnderlyings) SetUnderlyingCashAmount(v float64) {
	m.Set(field.NewUnderlyingCashAmount(v))
}

//SetUnderlyingCashType sets UnderlyingCashType, Tag 974
func (m NoUnderlyings) SetUnderlyingCashType(v string) {
	m.Set(field.NewUnderlyingCashType(v))
}

//SetUnderlyingUnitOfMeasure sets UnderlyingUnitOfMeasure, Tag 998
func (m NoUnderlyings) SetUnderlyingUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingUnitOfMeasure(v))
}

//SetUnderlyingTimeUnit sets UnderlyingTimeUnit, Tag 1000
func (m NoUnderlyings) SetUnderlyingTimeUnit(v string) {
	m.Set(field.NewUnderlyingTimeUnit(v))
}

//SetUnderlyingCapValue sets UnderlyingCapValue, Tag 1038
func (m NoUnderlyings) SetUnderlyingCapValue(v float64) {
	m.Set(field.NewUnderlyingCapValue(v))
}

//SetNoUndlyInstrumentParties sets NoUndlyInstrumentParties, Tag 1058
func (m NoUnderlyings) SetNoUndlyInstrumentParties(f NoUndlyInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingSettlMethod sets UnderlyingSettlMethod, Tag 1039
func (m NoUnderlyings) SetUnderlyingSettlMethod(v string) {
	m.Set(field.NewUnderlyingSettlMethod(v))
}

//SetUnderlyingAdjustedQuantity sets UnderlyingAdjustedQuantity, Tag 1044
func (m NoUnderlyings) SetUnderlyingAdjustedQuantity(v float64) {
	m.Set(field.NewUnderlyingAdjustedQuantity(v))
}

//SetUnderlyingFXRate sets UnderlyingFXRate, Tag 1045
func (m NoUnderlyings) SetUnderlyingFXRate(v float64) {
	m.Set(field.NewUnderlyingFXRate(v))
}

//SetUnderlyingFXRateCalc sets UnderlyingFXRateCalc, Tag 1046
func (m NoUnderlyings) SetUnderlyingFXRateCalc(v string) {
	m.Set(field.NewUnderlyingFXRateCalc(v))
}

//GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311
func (m NoUnderlyings) GetUnderlyingSymbol() (f field.UnderlyingSymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312
func (m NoUnderlyings) GetUnderlyingSymbolSfx() (f field.UnderlyingSymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309
func (m NoUnderlyings) GetUnderlyingSecurityID() (f field.UnderlyingSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityIDSource gets UnderlyingSecurityIDSource, Tag 305
func (m NoUnderlyings) GetUnderlyingSecurityIDSource() (f field.UnderlyingSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyingSecurityAltID gets NoUnderlyingSecurityAltID, Tag 457
func (m NoUnderlyings) GetNoUnderlyingSecurityAltID() (NoUnderlyingSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingProduct gets UnderlyingProduct, Tag 462
func (m NoUnderlyings) GetUnderlyingProduct() (f field.UnderlyingProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCFICode gets UnderlyingCFICode, Tag 463
func (m NoUnderlyings) GetUnderlyingCFICode() (f field.UnderlyingCFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310
func (m NoUnderlyings) GetUnderlyingSecurityType() (f field.UnderlyingSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecuritySubType gets UnderlyingSecuritySubType, Tag 763
func (m NoUnderlyings) GetUnderlyingSecuritySubType() (f field.UnderlyingSecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313
func (m NoUnderlyings) GetUnderlyingMaturityMonthYear() (f field.UnderlyingMaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingMaturityDate gets UnderlyingMaturityDate, Tag 542
func (m NoUnderlyings) GetUnderlyingMaturityDate() (f field.UnderlyingMaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCouponPaymentDate gets UnderlyingCouponPaymentDate, Tag 241
func (m NoUnderlyings) GetUnderlyingCouponPaymentDate() (f field.UnderlyingCouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingIssueDate gets UnderlyingIssueDate, Tag 242
func (m NoUnderlyings) GetUnderlyingIssueDate() (f field.UnderlyingIssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRepoCollateralSecurityType gets UnderlyingRepoCollateralSecurityType, Tag 243
func (m NoUnderlyings) GetUnderlyingRepoCollateralSecurityType() (f field.UnderlyingRepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRepurchaseTerm gets UnderlyingRepurchaseTerm, Tag 244
func (m NoUnderlyings) GetUnderlyingRepurchaseTerm() (f field.UnderlyingRepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRepurchaseRate gets UnderlyingRepurchaseRate, Tag 245
func (m NoUnderlyings) GetUnderlyingRepurchaseRate() (f field.UnderlyingRepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFactor gets UnderlyingFactor, Tag 246
func (m NoUnderlyings) GetUnderlyingFactor() (f field.UnderlyingFactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCreditRating gets UnderlyingCreditRating, Tag 256
func (m NoUnderlyings) GetUnderlyingCreditRating() (f field.UnderlyingCreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingInstrRegistry gets UnderlyingInstrRegistry, Tag 595
func (m NoUnderlyings) GetUnderlyingInstrRegistry() (f field.UnderlyingInstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCountryOfIssue gets UnderlyingCountryOfIssue, Tag 592
func (m NoUnderlyings) GetUnderlyingCountryOfIssue() (f field.UnderlyingCountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStateOrProvinceOfIssue gets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m NoUnderlyings) GetUnderlyingStateOrProvinceOfIssue() (f field.UnderlyingStateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLocaleOfIssue gets UnderlyingLocaleOfIssue, Tag 594
func (m NoUnderlyings) GetUnderlyingLocaleOfIssue() (f field.UnderlyingLocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRedemptionDate gets UnderlyingRedemptionDate, Tag 247
func (m NoUnderlyings) GetUnderlyingRedemptionDate() (f field.UnderlyingRedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316
func (m NoUnderlyings) GetUnderlyingStrikePrice() (f field.UnderlyingStrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStrikeCurrency gets UnderlyingStrikeCurrency, Tag 941
func (m NoUnderlyings) GetUnderlyingStrikeCurrency() (f field.UnderlyingStrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317
func (m NoUnderlyings) GetUnderlyingOptAttribute() (f field.UnderlyingOptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436
func (m NoUnderlyings) GetUnderlyingContractMultiplier() (f field.UnderlyingContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435
func (m NoUnderlyings) GetUnderlyingCouponRate() (f field.UnderlyingCouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308
func (m NoUnderlyings) GetUnderlyingSecurityExchange() (f field.UnderlyingSecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306
func (m NoUnderlyings) GetUnderlyingIssuer() (f field.UnderlyingIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362
func (m NoUnderlyings) GetEncodedUnderlyingIssuerLen() (f field.EncodedUnderlyingIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363
func (m NoUnderlyings) GetEncodedUnderlyingIssuer() (f field.EncodedUnderlyingIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307
func (m NoUnderlyings) GetUnderlyingSecurityDesc() (f field.UnderlyingSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDescLen() (f field.EncodedUnderlyingSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDesc() (f field.EncodedUnderlyingSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCPProgram gets UnderlyingCPProgram, Tag 877
func (m NoUnderlyings) GetUnderlyingCPProgram() (f field.UnderlyingCPProgramField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCPRegType gets UnderlyingCPRegType, Tag 878
func (m NoUnderlyings) GetUnderlyingCPRegType() (f field.UnderlyingCPRegTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCurrency gets UnderlyingCurrency, Tag 318
func (m NoUnderlyings) GetUnderlyingCurrency() (f field.UnderlyingCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingQty gets UnderlyingQty, Tag 879
func (m NoUnderlyings) GetUnderlyingQty() (f field.UnderlyingQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPx gets UnderlyingPx, Tag 810
func (m NoUnderlyings) GetUnderlyingPx() (f field.UnderlyingPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingDirtyPrice gets UnderlyingDirtyPrice, Tag 882
func (m NoUnderlyings) GetUnderlyingDirtyPrice() (f field.UnderlyingDirtyPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingEndPrice gets UnderlyingEndPrice, Tag 883
func (m NoUnderlyings) GetUnderlyingEndPrice() (f field.UnderlyingEndPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStartValue gets UnderlyingStartValue, Tag 884
func (m NoUnderlyings) GetUnderlyingStartValue() (f field.UnderlyingStartValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCurrentValue gets UnderlyingCurrentValue, Tag 885
func (m NoUnderlyings) GetUnderlyingCurrentValue() (f field.UnderlyingCurrentValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingEndValue gets UnderlyingEndValue, Tag 886
func (m NoUnderlyings) GetUnderlyingEndValue() (f field.UnderlyingEndValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyingStips gets NoUnderlyingStips, Tag 887
func (m NoUnderlyings) GetNoUnderlyingStips() (NoUnderlyingStipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingStipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingAllocationPercent gets UnderlyingAllocationPercent, Tag 972
func (m NoUnderlyings) GetUnderlyingAllocationPercent() (f field.UnderlyingAllocationPercentField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSettlementType gets UnderlyingSettlementType, Tag 975
func (m NoUnderlyings) GetUnderlyingSettlementType() (f field.UnderlyingSettlementTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCashAmount gets UnderlyingCashAmount, Tag 973
func (m NoUnderlyings) GetUnderlyingCashAmount() (f field.UnderlyingCashAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCashType gets UnderlyingCashType, Tag 974
func (m NoUnderlyings) GetUnderlyingCashType() (f field.UnderlyingCashTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingUnitOfMeasure gets UnderlyingUnitOfMeasure, Tag 998
func (m NoUnderlyings) GetUnderlyingUnitOfMeasure() (f field.UnderlyingUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingTimeUnit gets UnderlyingTimeUnit, Tag 1000
func (m NoUnderlyings) GetUnderlyingTimeUnit() (f field.UnderlyingTimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCapValue gets UnderlyingCapValue, Tag 1038
func (m NoUnderlyings) GetUnderlyingCapValue() (f field.UnderlyingCapValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUndlyInstrumentParties gets NoUndlyInstrumentParties, Tag 1058
func (m NoUnderlyings) GetNoUndlyInstrumentParties() (NoUndlyInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingSettlMethod gets UnderlyingSettlMethod, Tag 1039
func (m NoUnderlyings) GetUnderlyingSettlMethod() (f field.UnderlyingSettlMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingAdjustedQuantity gets UnderlyingAdjustedQuantity, Tag 1044
func (m NoUnderlyings) GetUnderlyingAdjustedQuantity() (f field.UnderlyingAdjustedQuantityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFXRate gets UnderlyingFXRate, Tag 1045
func (m NoUnderlyings) GetUnderlyingFXRate() (f field.UnderlyingFXRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFXRateCalc gets UnderlyingFXRateCalc, Tag 1046
func (m NoUnderlyings) GetUnderlyingFXRateCalc() (f field.UnderlyingFXRateCalcField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingSymbol returns true if UnderlyingSymbol is present, Tag 311
func (m NoUnderlyings) HasUnderlyingSymbol() bool {
	return m.Has(tag.UnderlyingSymbol)
}

//HasUnderlyingSymbolSfx returns true if UnderlyingSymbolSfx is present, Tag 312
func (m NoUnderlyings) HasUnderlyingSymbolSfx() bool {
	return m.Has(tag.UnderlyingSymbolSfx)
}

//HasUnderlyingSecurityID returns true if UnderlyingSecurityID is present, Tag 309
func (m NoUnderlyings) HasUnderlyingSecurityID() bool {
	return m.Has(tag.UnderlyingSecurityID)
}

//HasUnderlyingSecurityIDSource returns true if UnderlyingSecurityIDSource is present, Tag 305
func (m NoUnderlyings) HasUnderlyingSecurityIDSource() bool {
	return m.Has(tag.UnderlyingSecurityIDSource)
}

//HasNoUnderlyingSecurityAltID returns true if NoUnderlyingSecurityAltID is present, Tag 457
func (m NoUnderlyings) HasNoUnderlyingSecurityAltID() bool {
	return m.Has(tag.NoUnderlyingSecurityAltID)
}

//HasUnderlyingProduct returns true if UnderlyingProduct is present, Tag 462
func (m NoUnderlyings) HasUnderlyingProduct() bool {
	return m.Has(tag.UnderlyingProduct)
}

//HasUnderlyingCFICode returns true if UnderlyingCFICode is present, Tag 463
func (m NoUnderlyings) HasUnderlyingCFICode() bool {
	return m.Has(tag.UnderlyingCFICode)
}

//HasUnderlyingSecurityType returns true if UnderlyingSecurityType is present, Tag 310
func (m NoUnderlyings) HasUnderlyingSecurityType() bool {
	return m.Has(tag.UnderlyingSecurityType)
}

//HasUnderlyingSecuritySubType returns true if UnderlyingSecuritySubType is present, Tag 763
func (m NoUnderlyings) HasUnderlyingSecuritySubType() bool {
	return m.Has(tag.UnderlyingSecuritySubType)
}

//HasUnderlyingMaturityMonthYear returns true if UnderlyingMaturityMonthYear is present, Tag 313
func (m NoUnderlyings) HasUnderlyingMaturityMonthYear() bool {
	return m.Has(tag.UnderlyingMaturityMonthYear)
}

//HasUnderlyingMaturityDate returns true if UnderlyingMaturityDate is present, Tag 542
func (m NoUnderlyings) HasUnderlyingMaturityDate() bool {
	return m.Has(tag.UnderlyingMaturityDate)
}

//HasUnderlyingCouponPaymentDate returns true if UnderlyingCouponPaymentDate is present, Tag 241
func (m NoUnderlyings) HasUnderlyingCouponPaymentDate() bool {
	return m.Has(tag.UnderlyingCouponPaymentDate)
}

//HasUnderlyingIssueDate returns true if UnderlyingIssueDate is present, Tag 242
func (m NoUnderlyings) HasUnderlyingIssueDate() bool {
	return m.Has(tag.UnderlyingIssueDate)
}

//HasUnderlyingRepoCollateralSecurityType returns true if UnderlyingRepoCollateralSecurityType is present, Tag 243
func (m NoUnderlyings) HasUnderlyingRepoCollateralSecurityType() bool {
	return m.Has(tag.UnderlyingRepoCollateralSecurityType)
}

//HasUnderlyingRepurchaseTerm returns true if UnderlyingRepurchaseTerm is present, Tag 244
func (m NoUnderlyings) HasUnderlyingRepurchaseTerm() bool {
	return m.Has(tag.UnderlyingRepurchaseTerm)
}

//HasUnderlyingRepurchaseRate returns true if UnderlyingRepurchaseRate is present, Tag 245
func (m NoUnderlyings) HasUnderlyingRepurchaseRate() bool {
	return m.Has(tag.UnderlyingRepurchaseRate)
}

//HasUnderlyingFactor returns true if UnderlyingFactor is present, Tag 246
func (m NoUnderlyings) HasUnderlyingFactor() bool {
	return m.Has(tag.UnderlyingFactor)
}

//HasUnderlyingCreditRating returns true if UnderlyingCreditRating is present, Tag 256
func (m NoUnderlyings) HasUnderlyingCreditRating() bool {
	return m.Has(tag.UnderlyingCreditRating)
}

//HasUnderlyingInstrRegistry returns true if UnderlyingInstrRegistry is present, Tag 595
func (m NoUnderlyings) HasUnderlyingInstrRegistry() bool {
	return m.Has(tag.UnderlyingInstrRegistry)
}

//HasUnderlyingCountryOfIssue returns true if UnderlyingCountryOfIssue is present, Tag 592
func (m NoUnderlyings) HasUnderlyingCountryOfIssue() bool {
	return m.Has(tag.UnderlyingCountryOfIssue)
}

//HasUnderlyingStateOrProvinceOfIssue returns true if UnderlyingStateOrProvinceOfIssue is present, Tag 593
func (m NoUnderlyings) HasUnderlyingStateOrProvinceOfIssue() bool {
	return m.Has(tag.UnderlyingStateOrProvinceOfIssue)
}

//HasUnderlyingLocaleOfIssue returns true if UnderlyingLocaleOfIssue is present, Tag 594
func (m NoUnderlyings) HasUnderlyingLocaleOfIssue() bool {
	return m.Has(tag.UnderlyingLocaleOfIssue)
}

//HasUnderlyingRedemptionDate returns true if UnderlyingRedemptionDate is present, Tag 247
func (m NoUnderlyings) HasUnderlyingRedemptionDate() bool {
	return m.Has(tag.UnderlyingRedemptionDate)
}

//HasUnderlyingStrikePrice returns true if UnderlyingStrikePrice is present, Tag 316
func (m NoUnderlyings) HasUnderlyingStrikePrice() bool {
	return m.Has(tag.UnderlyingStrikePrice)
}

//HasUnderlyingStrikeCurrency returns true if UnderlyingStrikeCurrency is present, Tag 941
func (m NoUnderlyings) HasUnderlyingStrikeCurrency() bool {
	return m.Has(tag.UnderlyingStrikeCurrency)
}

//HasUnderlyingOptAttribute returns true if UnderlyingOptAttribute is present, Tag 317
func (m NoUnderlyings) HasUnderlyingOptAttribute() bool {
	return m.Has(tag.UnderlyingOptAttribute)
}

//HasUnderlyingContractMultiplier returns true if UnderlyingContractMultiplier is present, Tag 436
func (m NoUnderlyings) HasUnderlyingContractMultiplier() bool {
	return m.Has(tag.UnderlyingContractMultiplier)
}

//HasUnderlyingCouponRate returns true if UnderlyingCouponRate is present, Tag 435
func (m NoUnderlyings) HasUnderlyingCouponRate() bool {
	return m.Has(tag.UnderlyingCouponRate)
}

//HasUnderlyingSecurityExchange returns true if UnderlyingSecurityExchange is present, Tag 308
func (m NoUnderlyings) HasUnderlyingSecurityExchange() bool {
	return m.Has(tag.UnderlyingSecurityExchange)
}

//HasUnderlyingIssuer returns true if UnderlyingIssuer is present, Tag 306
func (m NoUnderlyings) HasUnderlyingIssuer() bool {
	return m.Has(tag.UnderlyingIssuer)
}

//HasEncodedUnderlyingIssuerLen returns true if EncodedUnderlyingIssuerLen is present, Tag 362
func (m NoUnderlyings) HasEncodedUnderlyingIssuerLen() bool {
	return m.Has(tag.EncodedUnderlyingIssuerLen)
}

//HasEncodedUnderlyingIssuer returns true if EncodedUnderlyingIssuer is present, Tag 363
func (m NoUnderlyings) HasEncodedUnderlyingIssuer() bool {
	return m.Has(tag.EncodedUnderlyingIssuer)
}

//HasUnderlyingSecurityDesc returns true if UnderlyingSecurityDesc is present, Tag 307
func (m NoUnderlyings) HasUnderlyingSecurityDesc() bool {
	return m.Has(tag.UnderlyingSecurityDesc)
}

//HasEncodedUnderlyingSecurityDescLen returns true if EncodedUnderlyingSecurityDescLen is present, Tag 364
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDescLen() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDescLen)
}

//HasEncodedUnderlyingSecurityDesc returns true if EncodedUnderlyingSecurityDesc is present, Tag 365
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDesc() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDesc)
}

//HasUnderlyingCPProgram returns true if UnderlyingCPProgram is present, Tag 877
func (m NoUnderlyings) HasUnderlyingCPProgram() bool {
	return m.Has(tag.UnderlyingCPProgram)
}

//HasUnderlyingCPRegType returns true if UnderlyingCPRegType is present, Tag 878
func (m NoUnderlyings) HasUnderlyingCPRegType() bool {
	return m.Has(tag.UnderlyingCPRegType)
}

//HasUnderlyingCurrency returns true if UnderlyingCurrency is present, Tag 318
func (m NoUnderlyings) HasUnderlyingCurrency() bool {
	return m.Has(tag.UnderlyingCurrency)
}

//HasUnderlyingQty returns true if UnderlyingQty is present, Tag 879
func (m NoUnderlyings) HasUnderlyingQty() bool {
	return m.Has(tag.UnderlyingQty)
}

//HasUnderlyingPx returns true if UnderlyingPx is present, Tag 810
func (m NoUnderlyings) HasUnderlyingPx() bool {
	return m.Has(tag.UnderlyingPx)
}

//HasUnderlyingDirtyPrice returns true if UnderlyingDirtyPrice is present, Tag 882
func (m NoUnderlyings) HasUnderlyingDirtyPrice() bool {
	return m.Has(tag.UnderlyingDirtyPrice)
}

//HasUnderlyingEndPrice returns true if UnderlyingEndPrice is present, Tag 883
func (m NoUnderlyings) HasUnderlyingEndPrice() bool {
	return m.Has(tag.UnderlyingEndPrice)
}

//HasUnderlyingStartValue returns true if UnderlyingStartValue is present, Tag 884
func (m NoUnderlyings) HasUnderlyingStartValue() bool {
	return m.Has(tag.UnderlyingStartValue)
}

//HasUnderlyingCurrentValue returns true if UnderlyingCurrentValue is present, Tag 885
func (m NoUnderlyings) HasUnderlyingCurrentValue() bool {
	return m.Has(tag.UnderlyingCurrentValue)
}

//HasUnderlyingEndValue returns true if UnderlyingEndValue is present, Tag 886
func (m NoUnderlyings) HasUnderlyingEndValue() bool {
	return m.Has(tag.UnderlyingEndValue)
}

//HasNoUnderlyingStips returns true if NoUnderlyingStips is present, Tag 887
func (m NoUnderlyings) HasNoUnderlyingStips() bool {
	return m.Has(tag.NoUnderlyingStips)
}

//HasUnderlyingAllocationPercent returns true if UnderlyingAllocationPercent is present, Tag 972
func (m NoUnderlyings) HasUnderlyingAllocationPercent() bool {
	return m.Has(tag.UnderlyingAllocationPercent)
}

//HasUnderlyingSettlementType returns true if UnderlyingSettlementType is present, Tag 975
func (m NoUnderlyings) HasUnderlyingSettlementType() bool {
	return m.Has(tag.UnderlyingSettlementType)
}

//HasUnderlyingCashAmount returns true if UnderlyingCashAmount is present, Tag 973
func (m NoUnderlyings) HasUnderlyingCashAmount() bool {
	return m.Has(tag.UnderlyingCashAmount)
}

//HasUnderlyingCashType returns true if UnderlyingCashType is present, Tag 974
func (m NoUnderlyings) HasUnderlyingCashType() bool {
	return m.Has(tag.UnderlyingCashType)
}

//HasUnderlyingUnitOfMeasure returns true if UnderlyingUnitOfMeasure is present, Tag 998
func (m NoUnderlyings) HasUnderlyingUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingUnitOfMeasure)
}

//HasUnderlyingTimeUnit returns true if UnderlyingTimeUnit is present, Tag 1000
func (m NoUnderlyings) HasUnderlyingTimeUnit() bool {
	return m.Has(tag.UnderlyingTimeUnit)
}

//HasUnderlyingCapValue returns true if UnderlyingCapValue is present, Tag 1038
func (m NoUnderlyings) HasUnderlyingCapValue() bool {
	return m.Has(tag.UnderlyingCapValue)
}

//HasNoUndlyInstrumentParties returns true if NoUndlyInstrumentParties is present, Tag 1058
func (m NoUnderlyings) HasNoUndlyInstrumentParties() bool {
	return m.Has(tag.NoUndlyInstrumentParties)
}

//HasUnderlyingSettlMethod returns true if UnderlyingSettlMethod is present, Tag 1039
func (m NoUnderlyings) HasUnderlyingSettlMethod() bool {
	return m.Has(tag.UnderlyingSettlMethod)
}

//HasUnderlyingAdjustedQuantity returns true if UnderlyingAdjustedQuantity is present, Tag 1044
func (m NoUnderlyings) HasUnderlyingAdjustedQuantity() bool {
	return m.Has(tag.UnderlyingAdjustedQuantity)
}

//HasUnderlyingFXRate returns true if UnderlyingFXRate is present, Tag 1045
func (m NoUnderlyings) HasUnderlyingFXRate() bool {
	return m.Has(tag.UnderlyingFXRate)
}

//HasUnderlyingFXRateCalc returns true if UnderlyingFXRateCalc is present, Tag 1046
func (m NoUnderlyings) HasUnderlyingFXRateCalc() bool {
	return m.Has(tag.UnderlyingFXRateCalc)
}

//NoUnderlyingSecurityAltID is a repeating group element, Tag 457
type NoUnderlyingSecurityAltID struct {
	quickfix.Group
}

//SetUnderlyingSecurityAltID sets UnderlyingSecurityAltID, Tag 458
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltID(v string) {
	m.Set(field.NewUnderlyingSecurityAltID(v))
}

//SetUnderlyingSecurityAltIDSource sets UnderlyingSecurityAltIDSource, Tag 459
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityAltIDSource(v))
}

//GetUnderlyingSecurityAltID gets UnderlyingSecurityAltID, Tag 458
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltID() (f field.UnderlyingSecurityAltIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityAltIDSource gets UnderlyingSecurityAltIDSource, Tag 459
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltIDSource() (f field.UnderlyingSecurityAltIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingSecurityAltID returns true if UnderlyingSecurityAltID is present, Tag 458
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltID() bool {
	return m.Has(tag.UnderlyingSecurityAltID)
}

//HasUnderlyingSecurityAltIDSource returns true if UnderlyingSecurityAltIDSource is present, Tag 459
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltIDSource() bool {
	return m.Has(tag.UnderlyingSecurityAltIDSource)
}

//NoUnderlyingSecurityAltIDRepeatingGroup is a repeating group, Tag 457
type NoUnderlyingSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingSecurityAltIDRepeatingGroup returns an initialized, NoUnderlyingSecurityAltIDRepeatingGroup
func NewNoUnderlyingSecurityAltIDRepeatingGroup() NoUnderlyingSecurityAltIDRepeatingGroup {
	return NoUnderlyingSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSecurityAltID), quickfix.GroupElement(tag.UnderlyingSecurityAltIDSource)})}
}

//Add create and append a new NoUnderlyingSecurityAltID to this group
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Add() NoUnderlyingSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingSecurityAltID{g}
}

//Get returns the ith NoUnderlyingSecurityAltID in the NoUnderlyingSecurityAltIDRepeatinGroup
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Get(i int) NoUnderlyingSecurityAltID {
	return NoUnderlyingSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoUnderlyingStips is a repeating group element, Tag 887
type NoUnderlyingStips struct {
	quickfix.Group
}

//SetUnderlyingStipType sets UnderlyingStipType, Tag 888
func (m NoUnderlyingStips) SetUnderlyingStipType(v string) {
	m.Set(field.NewUnderlyingStipType(v))
}

//SetUnderlyingStipValue sets UnderlyingStipValue, Tag 889
func (m NoUnderlyingStips) SetUnderlyingStipValue(v string) {
	m.Set(field.NewUnderlyingStipValue(v))
}

//GetUnderlyingStipType gets UnderlyingStipType, Tag 888
func (m NoUnderlyingStips) GetUnderlyingStipType() (f field.UnderlyingStipTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStipValue gets UnderlyingStipValue, Tag 889
func (m NoUnderlyingStips) GetUnderlyingStipValue() (f field.UnderlyingStipValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingStipType returns true if UnderlyingStipType is present, Tag 888
func (m NoUnderlyingStips) HasUnderlyingStipType() bool {
	return m.Has(tag.UnderlyingStipType)
}

//HasUnderlyingStipValue returns true if UnderlyingStipValue is present, Tag 889
func (m NoUnderlyingStips) HasUnderlyingStipValue() bool {
	return m.Has(tag.UnderlyingStipValue)
}

//NoUnderlyingStipsRepeatingGroup is a repeating group, Tag 887
type NoUnderlyingStipsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingStipsRepeatingGroup returns an initialized, NoUnderlyingStipsRepeatingGroup
func NewNoUnderlyingStipsRepeatingGroup() NoUnderlyingStipsRepeatingGroup {
	return NoUnderlyingStipsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingStips,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingStipType), quickfix.GroupElement(tag.UnderlyingStipValue)})}
}

//Add create and append a new NoUnderlyingStips to this group
func (m NoUnderlyingStipsRepeatingGroup) Add() NoUnderlyingStips {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingStips{g}
}

//Get returns the ith NoUnderlyingStips in the NoUnderlyingStipsRepeatinGroup
func (m NoUnderlyingStipsRepeatingGroup) Get(i int) NoUnderlyingStips {
	return NoUnderlyingStips{m.RepeatingGroup.Get(i)}
}

//NoUndlyInstrumentParties is a repeating group element, Tag 1058
type NoUndlyInstrumentParties struct {
	quickfix.Group
}

//SetUndlyInstrumentPartyID sets UndlyInstrumentPartyID, Tag 1059
func (m NoUndlyInstrumentParties) SetUndlyInstrumentPartyID(v string) {
	m.Set(field.NewUndlyInstrumentPartyID(v))
}

//SetUndlyInstrumentPartyIDSource sets UndlyInstrumentPartyIDSource, Tag 1060
func (m NoUndlyInstrumentParties) SetUndlyInstrumentPartyIDSource(v string) {
	m.Set(field.NewUndlyInstrumentPartyIDSource(v))
}

//SetUndlyInstrumentPartyRole sets UndlyInstrumentPartyRole, Tag 1061
func (m NoUndlyInstrumentParties) SetUndlyInstrumentPartyRole(v int) {
	m.Set(field.NewUndlyInstrumentPartyRole(v))
}

//SetNoUndlyInstrumentPartySubIDs sets NoUndlyInstrumentPartySubIDs, Tag 1062
func (m NoUndlyInstrumentParties) SetNoUndlyInstrumentPartySubIDs(f NoUndlyInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetUndlyInstrumentPartyID gets UndlyInstrumentPartyID, Tag 1059
func (m NoUndlyInstrumentParties) GetUndlyInstrumentPartyID() (f field.UndlyInstrumentPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUndlyInstrumentPartyIDSource gets UndlyInstrumentPartyIDSource, Tag 1060
func (m NoUndlyInstrumentParties) GetUndlyInstrumentPartyIDSource() (f field.UndlyInstrumentPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUndlyInstrumentPartyRole gets UndlyInstrumentPartyRole, Tag 1061
func (m NoUndlyInstrumentParties) GetUndlyInstrumentPartyRole() (f field.UndlyInstrumentPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUndlyInstrumentPartySubIDs gets NoUndlyInstrumentPartySubIDs, Tag 1062
func (m NoUndlyInstrumentParties) GetNoUndlyInstrumentPartySubIDs() (NoUndlyInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasUndlyInstrumentPartyID returns true if UndlyInstrumentPartyID is present, Tag 1059
func (m NoUndlyInstrumentParties) HasUndlyInstrumentPartyID() bool {
	return m.Has(tag.UndlyInstrumentPartyID)
}

//HasUndlyInstrumentPartyIDSource returns true if UndlyInstrumentPartyIDSource is present, Tag 1060
func (m NoUndlyInstrumentParties) HasUndlyInstrumentPartyIDSource() bool {
	return m.Has(tag.UndlyInstrumentPartyIDSource)
}

//HasUndlyInstrumentPartyRole returns true if UndlyInstrumentPartyRole is present, Tag 1061
func (m NoUndlyInstrumentParties) HasUndlyInstrumentPartyRole() bool {
	return m.Has(tag.UndlyInstrumentPartyRole)
}

//HasNoUndlyInstrumentPartySubIDs returns true if NoUndlyInstrumentPartySubIDs is present, Tag 1062
func (m NoUndlyInstrumentParties) HasNoUndlyInstrumentPartySubIDs() bool {
	return m.Has(tag.NoUndlyInstrumentPartySubIDs)
}

//NoUndlyInstrumentPartySubIDs is a repeating group element, Tag 1062
type NoUndlyInstrumentPartySubIDs struct {
	quickfix.Group
}

//SetUndlyInstrumentPartySubID sets UndlyInstrumentPartySubID, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) SetUndlyInstrumentPartySubID(v string) {
	m.Set(field.NewUndlyInstrumentPartySubID(v))
}

//SetUndlyInstrumentPartySubIDType sets UndlyInstrumentPartySubIDType, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) SetUndlyInstrumentPartySubIDType(v int) {
	m.Set(field.NewUndlyInstrumentPartySubIDType(v))
}

//GetUndlyInstrumentPartySubID gets UndlyInstrumentPartySubID, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) GetUndlyInstrumentPartySubID() (f field.UndlyInstrumentPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUndlyInstrumentPartySubIDType gets UndlyInstrumentPartySubIDType, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) GetUndlyInstrumentPartySubIDType() (f field.UndlyInstrumentPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUndlyInstrumentPartySubID returns true if UndlyInstrumentPartySubID is present, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) HasUndlyInstrumentPartySubID() bool {
	return m.Has(tag.UndlyInstrumentPartySubID)
}

//HasUndlyInstrumentPartySubIDType returns true if UndlyInstrumentPartySubIDType is present, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) HasUndlyInstrumentPartySubIDType() bool {
	return m.Has(tag.UndlyInstrumentPartySubIDType)
}

//NoUndlyInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1062
type NoUndlyInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUndlyInstrumentPartySubIDsRepeatingGroup returns an initialized, NoUndlyInstrumentPartySubIDsRepeatingGroup
func NewNoUndlyInstrumentPartySubIDsRepeatingGroup() NoUndlyInstrumentPartySubIDsRepeatingGroup {
	return NoUndlyInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UndlyInstrumentPartySubID), quickfix.GroupElement(tag.UndlyInstrumentPartySubIDType)})}
}

//Add create and append a new NoUndlyInstrumentPartySubIDs to this group
func (m NoUndlyInstrumentPartySubIDsRepeatingGroup) Add() NoUndlyInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoUndlyInstrumentPartySubIDs{g}
}

//Get returns the ith NoUndlyInstrumentPartySubIDs in the NoUndlyInstrumentPartySubIDsRepeatinGroup
func (m NoUndlyInstrumentPartySubIDsRepeatingGroup) Get(i int) NoUndlyInstrumentPartySubIDs {
	return NoUndlyInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoUndlyInstrumentPartiesRepeatingGroup is a repeating group, Tag 1058
type NoUndlyInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUndlyInstrumentPartiesRepeatingGroup returns an initialized, NoUndlyInstrumentPartiesRepeatingGroup
func NewNoUndlyInstrumentPartiesRepeatingGroup() NoUndlyInstrumentPartiesRepeatingGroup {
	return NoUndlyInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UndlyInstrumentPartyID), quickfix.GroupElement(tag.UndlyInstrumentPartyIDSource), quickfix.GroupElement(tag.UndlyInstrumentPartyRole), quickfix.GroupElement(tag.NoUndlyInstrumentPartySubIDs)})}
}

//Add create and append a new NoUndlyInstrumentParties to this group
func (m NoUndlyInstrumentPartiesRepeatingGroup) Add() NoUndlyInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoUndlyInstrumentParties{g}
}

//Get returns the ith NoUndlyInstrumentParties in the NoUndlyInstrumentPartiesRepeatinGroup
func (m NoUndlyInstrumentPartiesRepeatingGroup) Get(i int) NoUndlyInstrumentParties {
	return NoUndlyInstrumentParties{m.RepeatingGroup.Get(i)}
}

//NoUnderlyingsRepeatingGroup is a repeating group, Tag 711
type NoUnderlyingsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingsRepeatingGroup returns an initialized, NoUnderlyingsRepeatingGroup
func NewNoUnderlyingsRepeatingGroup() NoUnderlyingsRepeatingGroup {
	return NoUnderlyingsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyings,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingSecurityIDSource), quickfix.GroupElement(tag.NoUnderlyingSecurityAltID), quickfix.GroupElement(tag.UnderlyingProduct), quickfix.GroupElement(tag.UnderlyingCFICode), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingSecuritySubType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDate), quickfix.GroupElement(tag.UnderlyingCouponPaymentDate), quickfix.GroupElement(tag.UnderlyingIssueDate), quickfix.GroupElement(tag.UnderlyingRepoCollateralSecurityType), quickfix.GroupElement(tag.UnderlyingRepurchaseTerm), quickfix.GroupElement(tag.UnderlyingRepurchaseRate), quickfix.GroupElement(tag.UnderlyingFactor), quickfix.GroupElement(tag.UnderlyingCreditRating), quickfix.GroupElement(tag.UnderlyingInstrRegistry), quickfix.GroupElement(tag.UnderlyingCountryOfIssue), quickfix.GroupElement(tag.UnderlyingStateOrProvinceOfIssue), quickfix.GroupElement(tag.UnderlyingLocaleOfIssue), quickfix.GroupElement(tag.UnderlyingRedemptionDate), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingStrikeCurrency), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.UnderlyingCPProgram), quickfix.GroupElement(tag.UnderlyingCPRegType), quickfix.GroupElement(tag.UnderlyingCurrency), quickfix.GroupElement(tag.UnderlyingQty), quickfix.GroupElement(tag.UnderlyingPx), quickfix.GroupElement(tag.UnderlyingDirtyPrice), quickfix.GroupElement(tag.UnderlyingEndPrice), quickfix.GroupElement(tag.UnderlyingStartValue), quickfix.GroupElement(tag.UnderlyingCurrentValue), quickfix.GroupElement(tag.UnderlyingEndValue), quickfix.GroupElement(tag.NoUnderlyingStips), quickfix.GroupElement(tag.UnderlyingAllocationPercent), quickfix.GroupElement(tag.UnderlyingSettlementType), quickfix.GroupElement(tag.UnderlyingCashAmount), quickfix.GroupElement(tag.UnderlyingCashType), quickfix.GroupElement(tag.UnderlyingUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingTimeUnit), quickfix.GroupElement(tag.UnderlyingCapValue), quickfix.GroupElement(tag.NoUndlyInstrumentParties), quickfix.GroupElement(tag.UnderlyingSettlMethod), quickfix.GroupElement(tag.UnderlyingAdjustedQuantity), quickfix.GroupElement(tag.UnderlyingFXRate), quickfix.GroupElement(tag.UnderlyingFXRateCalc)})}
}

//Add create and append a new NoUnderlyings to this group
func (m NoUnderlyingsRepeatingGroup) Add() NoUnderlyings {
	g := m.RepeatingGroup.Add()
	return NoUnderlyings{g}
}

//Get returns the ith NoUnderlyings in the NoUnderlyingsRepeatinGroup
func (m NoUnderlyingsRepeatingGroup) Get(i int) NoUnderlyings {
	return NoUnderlyings{m.RepeatingGroup.Get(i)}
}

//NoEvents is a repeating group element, Tag 864
type NoEvents struct {
	quickfix.Group
}

//SetEventType sets EventType, Tag 865
func (m NoEvents) SetEventType(v int) {
	m.Set(field.NewEventType(v))
}

//SetEventDate sets EventDate, Tag 866
func (m NoEvents) SetEventDate(v string) {
	m.Set(field.NewEventDate(v))
}

//SetEventPx sets EventPx, Tag 867
func (m NoEvents) SetEventPx(v float64) {
	m.Set(field.NewEventPx(v))
}

//SetEventText sets EventText, Tag 868
func (m NoEvents) SetEventText(v string) {
	m.Set(field.NewEventText(v))
}

//GetEventType gets EventType, Tag 865
func (m NoEvents) GetEventType() (f field.EventTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEventDate gets EventDate, Tag 866
func (m NoEvents) GetEventDate() (f field.EventDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEventPx gets EventPx, Tag 867
func (m NoEvents) GetEventPx() (f field.EventPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEventText gets EventText, Tag 868
func (m NoEvents) GetEventText() (f field.EventTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasEventType returns true if EventType is present, Tag 865
func (m NoEvents) HasEventType() bool {
	return m.Has(tag.EventType)
}

//HasEventDate returns true if EventDate is present, Tag 866
func (m NoEvents) HasEventDate() bool {
	return m.Has(tag.EventDate)
}

//HasEventPx returns true if EventPx is present, Tag 867
func (m NoEvents) HasEventPx() bool {
	return m.Has(tag.EventPx)
}

//HasEventText returns true if EventText is present, Tag 868
func (m NoEvents) HasEventText() bool {
	return m.Has(tag.EventText)
}

//NoEventsRepeatingGroup is a repeating group, Tag 864
type NoEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoEventsRepeatingGroup returns an initialized, NoEventsRepeatingGroup
func NewNoEventsRepeatingGroup() NoEventsRepeatingGroup {
	return NoEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.EventType), quickfix.GroupElement(tag.EventDate), quickfix.GroupElement(tag.EventPx), quickfix.GroupElement(tag.EventText)})}
}

//Add create and append a new NoEvents to this group
func (m NoEventsRepeatingGroup) Add() NoEvents {
	g := m.RepeatingGroup.Add()
	return NoEvents{g}
}

//Get returns the ith NoEvents in the NoEventsRepeatinGroup
func (m NoEventsRepeatingGroup) Get(i int) NoEvents {
	return NoEvents{m.RepeatingGroup.Get(i)}
}

//NoInstrumentParties is a repeating group element, Tag 1018
type NoInstrumentParties struct {
	quickfix.Group
}

//SetInstrumentPartyID sets InstrumentPartyID, Tag 1019
func (m NoInstrumentParties) SetInstrumentPartyID(v string) {
	m.Set(field.NewInstrumentPartyID(v))
}

//SetInstrumentPartyIDSource sets InstrumentPartyIDSource, Tag 1050
func (m NoInstrumentParties) SetInstrumentPartyIDSource(v string) {
	m.Set(field.NewInstrumentPartyIDSource(v))
}

//SetInstrumentPartyRole sets InstrumentPartyRole, Tag 1051
func (m NoInstrumentParties) SetInstrumentPartyRole(v int) {
	m.Set(field.NewInstrumentPartyRole(v))
}

//SetNoInstrumentPartySubIDs sets NoInstrumentPartySubIDs, Tag 1052
func (m NoInstrumentParties) SetNoInstrumentPartySubIDs(f NoInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetInstrumentPartyID gets InstrumentPartyID, Tag 1019
func (m NoInstrumentParties) GetInstrumentPartyID() (f field.InstrumentPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrumentPartyIDSource gets InstrumentPartyIDSource, Tag 1050
func (m NoInstrumentParties) GetInstrumentPartyIDSource() (f field.InstrumentPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrumentPartyRole gets InstrumentPartyRole, Tag 1051
func (m NoInstrumentParties) GetInstrumentPartyRole() (f field.InstrumentPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoInstrumentPartySubIDs gets NoInstrumentPartySubIDs, Tag 1052
func (m NoInstrumentParties) GetNoInstrumentPartySubIDs() (NoInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasInstrumentPartyID returns true if InstrumentPartyID is present, Tag 1019
func (m NoInstrumentParties) HasInstrumentPartyID() bool {
	return m.Has(tag.InstrumentPartyID)
}

//HasInstrumentPartyIDSource returns true if InstrumentPartyIDSource is present, Tag 1050
func (m NoInstrumentParties) HasInstrumentPartyIDSource() bool {
	return m.Has(tag.InstrumentPartyIDSource)
}

//HasInstrumentPartyRole returns true if InstrumentPartyRole is present, Tag 1051
func (m NoInstrumentParties) HasInstrumentPartyRole() bool {
	return m.Has(tag.InstrumentPartyRole)
}

//HasNoInstrumentPartySubIDs returns true if NoInstrumentPartySubIDs is present, Tag 1052
func (m NoInstrumentParties) HasNoInstrumentPartySubIDs() bool {
	return m.Has(tag.NoInstrumentPartySubIDs)
}

//NoInstrumentPartySubIDs is a repeating group element, Tag 1052
type NoInstrumentPartySubIDs struct {
	quickfix.Group
}

//SetInstrumentPartySubID sets InstrumentPartySubID, Tag 1053
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubID(v string) {
	m.Set(field.NewInstrumentPartySubID(v))
}

//SetInstrumentPartySubIDType sets InstrumentPartySubIDType, Tag 1054
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubIDType(v int) {
	m.Set(field.NewInstrumentPartySubIDType(v))
}

//GetInstrumentPartySubID gets InstrumentPartySubID, Tag 1053
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubID() (f field.InstrumentPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrumentPartySubIDType gets InstrumentPartySubIDType, Tag 1054
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubIDType() (f field.InstrumentPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasInstrumentPartySubID returns true if InstrumentPartySubID is present, Tag 1053
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubID() bool {
	return m.Has(tag.InstrumentPartySubID)
}

//HasInstrumentPartySubIDType returns true if InstrumentPartySubIDType is present, Tag 1054
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubIDType() bool {
	return m.Has(tag.InstrumentPartySubIDType)
}

//NoInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1052
type NoInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoInstrumentPartySubIDsRepeatingGroup returns an initialized, NoInstrumentPartySubIDsRepeatingGroup
func NewNoInstrumentPartySubIDsRepeatingGroup() NoInstrumentPartySubIDsRepeatingGroup {
	return NoInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartySubID), quickfix.GroupElement(tag.InstrumentPartySubIDType)})}
}

//Add create and append a new NoInstrumentPartySubIDs to this group
func (m NoInstrumentPartySubIDsRepeatingGroup) Add() NoInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoInstrumentPartySubIDs{g}
}

//Get returns the ith NoInstrumentPartySubIDs in the NoInstrumentPartySubIDsRepeatinGroup
func (m NoInstrumentPartySubIDsRepeatingGroup) Get(i int) NoInstrumentPartySubIDs {
	return NoInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoInstrumentPartiesRepeatingGroup is a repeating group, Tag 1018
type NoInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoInstrumentPartiesRepeatingGroup returns an initialized, NoInstrumentPartiesRepeatingGroup
func NewNoInstrumentPartiesRepeatingGroup() NoInstrumentPartiesRepeatingGroup {
	return NoInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartyID), quickfix.GroupElement(tag.InstrumentPartyIDSource), quickfix.GroupElement(tag.InstrumentPartyRole), quickfix.GroupElement(tag.NoInstrumentPartySubIDs)})}
}

//Add create and append a new NoInstrumentParties to this group
func (m NoInstrumentPartiesRepeatingGroup) Add() NoInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoInstrumentParties{g}
}

//Get returns the ith NoInstrumentParties in the NoInstrumentPartiesRepeatinGroup
func (m NoInstrumentPartiesRepeatingGroup) Get(i int) NoInstrumentParties {
	return NoInstrumentParties{m.RepeatingGroup.Get(i)}
}
