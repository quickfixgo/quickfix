package collateralrequest

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/gen/enum"
	"github.com/quickfixgo/quickfix/gen/field"
	"github.com/quickfixgo/quickfix/gen/fixt11"
	"github.com/quickfixgo/quickfix/gen/tag"
)

// CollateralRequest is the fix50sp2 CollateralRequest type, MsgType = AX.
type CollateralRequest struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a CollateralRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) CollateralRequest {
	return CollateralRequest{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m CollateralRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a CollateralRequest initialized with the required fields for CollateralRequest.
func New(collreqid field.CollReqIDField, collasgnreason field.CollAsgnReasonField, transacttime field.TransactTimeField) (m CollateralRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("AX"))
	m.Set(collreqid)
	m.Set(collasgnreason)
	m.Set(transacttime)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg CollateralRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "AX", r
}

// SetAccount sets Account, Tag 1.
func (m CollateralRequest) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

// SetClOrdID sets ClOrdID, Tag 11.
func (m CollateralRequest) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

// SetCurrency sets Currency, Tag 15.
func (m CollateralRequest) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

// SetSecurityIDSource sets SecurityIDSource, Tag 22.
func (m CollateralRequest) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

// SetOrderID sets OrderID, Tag 37.
func (m CollateralRequest) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

// SetPrice sets Price, Tag 44.
func (m CollateralRequest) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

// SetSecurityID sets SecurityID, Tag 48.
func (m CollateralRequest) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetQuantity sets Quantity, Tag 53.
func (m CollateralRequest) SetQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewQuantity(value, scale))
}

// SetSide sets Side, Tag 54.
func (m CollateralRequest) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

// SetSymbol sets Symbol, Tag 55.
func (m CollateralRequest) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetText sets Text, Tag 58.
func (m CollateralRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

// SetTransactTime sets TransactTime, Tag 60.
func (m CollateralRequest) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

// SetSettlDate sets SettlDate, Tag 64.
func (m CollateralRequest) SetSettlDate(v string) {
	m.Set(field.NewSettlDate(v))
}

// SetSymbolSfx sets SymbolSfx, Tag 65.
func (m CollateralRequest) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

// SetIssuer sets Issuer, Tag 106.
func (m CollateralRequest) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

// SetSecurityDesc sets SecurityDesc, Tag 107.
func (m CollateralRequest) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

// SetNoExecs sets NoExecs, Tag 124.
func (m CollateralRequest) SetNoExecs(f NoExecsRepeatingGroup) {
	m.SetGroup(f)
}

// SetExpireTime sets ExpireTime, Tag 126.
func (m CollateralRequest) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

// SetNoMiscFees sets NoMiscFees, Tag 136.
func (m CollateralRequest) SetNoMiscFees(f NoMiscFeesRepeatingGroup) {
	m.SetGroup(f)
}

// SetAccruedInterestAmt sets AccruedInterestAmt, Tag 159.
func (m CollateralRequest) SetAccruedInterestAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewAccruedInterestAmt(value, scale))
}

// SetSecurityType sets SecurityType, Tag 167.
func (m CollateralRequest) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

// SetSecondaryOrderID sets SecondaryOrderID, Tag 198.
func (m CollateralRequest) SetSecondaryOrderID(v string) {
	m.Set(field.NewSecondaryOrderID(v))
}

// SetMaturityMonthYear sets MaturityMonthYear, Tag 200.
func (m CollateralRequest) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

// SetPutOrCall sets PutOrCall, Tag 201.
func (m CollateralRequest) SetPutOrCall(v enum.PutOrCall) {
	m.Set(field.NewPutOrCall(v))
}

// SetStrikePrice sets StrikePrice, Tag 202.
func (m CollateralRequest) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

// SetOptAttribute sets OptAttribute, Tag 206.
func (m CollateralRequest) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

// SetSecurityExchange sets SecurityExchange, Tag 207.
func (m CollateralRequest) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

// SetSpread sets Spread, Tag 218.
func (m CollateralRequest) SetSpread(value decimal.Decimal, scale int32) {
	m.Set(field.NewSpread(value, scale))
}

// SetBenchmarkCurveCurrency sets BenchmarkCurveCurrency, Tag 220.
func (m CollateralRequest) SetBenchmarkCurveCurrency(v string) {
	m.Set(field.NewBenchmarkCurveCurrency(v))
}

// SetBenchmarkCurveName sets BenchmarkCurveName, Tag 221.
func (m CollateralRequest) SetBenchmarkCurveName(v enum.BenchmarkCurveName) {
	m.Set(field.NewBenchmarkCurveName(v))
}

// SetBenchmarkCurvePoint sets BenchmarkCurvePoint, Tag 222.
func (m CollateralRequest) SetBenchmarkCurvePoint(v string) {
	m.Set(field.NewBenchmarkCurvePoint(v))
}

// SetCouponRate sets CouponRate, Tag 223.
func (m CollateralRequest) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

// SetCouponPaymentDate sets CouponPaymentDate, Tag 224.
func (m CollateralRequest) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

// SetIssueDate sets IssueDate, Tag 225.
func (m CollateralRequest) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

// SetRepurchaseTerm sets RepurchaseTerm, Tag 226.
func (m CollateralRequest) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

// SetRepurchaseRate sets RepurchaseRate, Tag 227.
func (m CollateralRequest) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

// SetFactor sets Factor, Tag 228.
func (m CollateralRequest) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

// SetContractMultiplier sets ContractMultiplier, Tag 231.
func (m CollateralRequest) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

// SetNoStipulations sets NoStipulations, Tag 232.
func (m CollateralRequest) SetNoStipulations(f NoStipulationsRepeatingGroup) {
	m.SetGroup(f)
}

// SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239.
func (m CollateralRequest) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

// SetRedemptionDate sets RedemptionDate, Tag 240.
func (m CollateralRequest) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

// SetCreditRating sets CreditRating, Tag 255.
func (m CollateralRequest) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m CollateralRequest) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348.
func (m CollateralRequest) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

// SetEncodedIssuer sets EncodedIssuer, Tag 349.
func (m CollateralRequest) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

// SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350.
func (m CollateralRequest) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

// SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351.
func (m CollateralRequest) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

// SetEncodedTextLen sets EncodedTextLen, Tag 354.
func (m CollateralRequest) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

// SetEncodedText sets EncodedText, Tag 355.
func (m CollateralRequest) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

// SetPriceType sets PriceType, Tag 423.
func (m CollateralRequest) SetPriceType(v enum.PriceType) {
	m.Set(field.NewPriceType(v))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m CollateralRequest) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoSecurityAltID sets NoSecurityAltID, Tag 454.
func (m CollateralRequest) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetProduct sets Product, Tag 460.
func (m CollateralRequest) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

// SetCFICode sets CFICode, Tag 461.
func (m CollateralRequest) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

// SetCountryOfIssue sets CountryOfIssue, Tag 470.
func (m CollateralRequest) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

// SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471.
func (m CollateralRequest) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

// SetLocaleOfIssue sets LocaleOfIssue, Tag 472.
func (m CollateralRequest) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

// SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526.
func (m CollateralRequest) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

// SetMaturityDate sets MaturityDate, Tag 541.
func (m CollateralRequest) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

// SetInstrRegistry sets InstrRegistry, Tag 543.
func (m CollateralRequest) SetInstrRegistry(v enum.InstrRegistry) {
	m.Set(field.NewInstrRegistry(v))
}

// SetNoLegs sets NoLegs, Tag 555.
func (m CollateralRequest) SetNoLegs(f NoLegsRepeatingGroup) {
	m.SetGroup(f)
}

// SetAccountType sets AccountType, Tag 581.
func (m CollateralRequest) SetAccountType(v enum.AccountType) {
	m.Set(field.NewAccountType(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625.
func (m CollateralRequest) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// SetBenchmarkPrice sets BenchmarkPrice, Tag 662.
func (m CollateralRequest) SetBenchmarkPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewBenchmarkPrice(value, scale))
}

// SetBenchmarkPriceType sets BenchmarkPriceType, Tag 663.
func (m CollateralRequest) SetBenchmarkPriceType(v int) {
	m.Set(field.NewBenchmarkPriceType(v))
}

// SetContractSettlMonth sets ContractSettlMonth, Tag 667.
func (m CollateralRequest) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

// SetPool sets Pool, Tag 691.
func (m CollateralRequest) SetPool(v string) {
	m.Set(field.NewPool(v))
}

// SetBenchmarkSecurityID sets BenchmarkSecurityID, Tag 699.
func (m CollateralRequest) SetBenchmarkSecurityID(v string) {
	m.Set(field.NewBenchmarkSecurityID(v))
}

// SetNoUnderlyings sets NoUnderlyings, Tag 711.
func (m CollateralRequest) SetNoUnderlyings(f NoUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

// SetClearingBusinessDate sets ClearingBusinessDate, Tag 715.
func (m CollateralRequest) SetClearingBusinessDate(v string) {
	m.Set(field.NewClearingBusinessDate(v))
}

// SetSettlSessID sets SettlSessID, Tag 716.
func (m CollateralRequest) SetSettlSessID(v enum.SettlSessID) {
	m.Set(field.NewSettlSessID(v))
}

// SetSettlSessSubID sets SettlSessSubID, Tag 717.
func (m CollateralRequest) SetSettlSessSubID(v string) {
	m.Set(field.NewSettlSessSubID(v))
}

// SetBenchmarkSecurityIDSource sets BenchmarkSecurityIDSource, Tag 761.
func (m CollateralRequest) SetBenchmarkSecurityIDSource(v string) {
	m.Set(field.NewBenchmarkSecurityIDSource(v))
}

// SetSecuritySubType sets SecuritySubType, Tag 762.
func (m CollateralRequest) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

// SetNoTrdRegTimestamps sets NoTrdRegTimestamps, Tag 768.
func (m CollateralRequest) SetNoTrdRegTimestamps(f NoTrdRegTimestampsRepeatingGroup) {
	m.SetGroup(f)
}

// SetTerminationType sets TerminationType, Tag 788.
func (m CollateralRequest) SetTerminationType(v enum.TerminationType) {
	m.Set(field.NewTerminationType(v))
}

// SetQtyType sets QtyType, Tag 854.
func (m CollateralRequest) SetQtyType(v enum.QtyType) {
	m.Set(field.NewQtyType(v))
}

// SetNoEvents sets NoEvents, Tag 864.
func (m CollateralRequest) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

// SetDatedDate sets DatedDate, Tag 873.
func (m CollateralRequest) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

// SetInterestAccrualDate sets InterestAccrualDate, Tag 874.
func (m CollateralRequest) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

// SetCPProgram sets CPProgram, Tag 875.
func (m CollateralRequest) SetCPProgram(v enum.CPProgram) {
	m.Set(field.NewCPProgram(v))
}

// SetCPRegType sets CPRegType, Tag 876.
func (m CollateralRequest) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

// SetCollReqID sets CollReqID, Tag 894.
func (m CollateralRequest) SetCollReqID(v string) {
	m.Set(field.NewCollReqID(v))
}

// SetCollAsgnReason sets CollAsgnReason, Tag 895.
func (m CollateralRequest) SetCollAsgnReason(v enum.CollAsgnReason) {
	m.Set(field.NewCollAsgnReason(v))
}

// SetNoTrades sets NoTrades, Tag 897.
func (m CollateralRequest) SetNoTrades(f NoTradesRepeatingGroup) {
	m.SetGroup(f)
}

// SetMarginRatio sets MarginRatio, Tag 898.
func (m CollateralRequest) SetMarginRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewMarginRatio(value, scale))
}

// SetMarginExcess sets MarginExcess, Tag 899.
func (m CollateralRequest) SetMarginExcess(value decimal.Decimal, scale int32) {
	m.Set(field.NewMarginExcess(value, scale))
}

// SetTotalNetValue sets TotalNetValue, Tag 900.
func (m CollateralRequest) SetTotalNetValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewTotalNetValue(value, scale))
}

// SetCashOutstanding sets CashOutstanding, Tag 901.
func (m CollateralRequest) SetCashOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewCashOutstanding(value, scale))
}

// SetAgreementDesc sets AgreementDesc, Tag 913.
func (m CollateralRequest) SetAgreementDesc(v string) {
	m.Set(field.NewAgreementDesc(v))
}

// SetAgreementID sets AgreementID, Tag 914.
func (m CollateralRequest) SetAgreementID(v string) {
	m.Set(field.NewAgreementID(v))
}

// SetAgreementDate sets AgreementDate, Tag 915.
func (m CollateralRequest) SetAgreementDate(v string) {
	m.Set(field.NewAgreementDate(v))
}

// SetStartDate sets StartDate, Tag 916.
func (m CollateralRequest) SetStartDate(v string) {
	m.Set(field.NewStartDate(v))
}

// SetEndDate sets EndDate, Tag 917.
func (m CollateralRequest) SetEndDate(v string) {
	m.Set(field.NewEndDate(v))
}

// SetAgreementCurrency sets AgreementCurrency, Tag 918.
func (m CollateralRequest) SetAgreementCurrency(v string) {
	m.Set(field.NewAgreementCurrency(v))
}

// SetDeliveryType sets DeliveryType, Tag 919.
func (m CollateralRequest) SetDeliveryType(v enum.DeliveryType) {
	m.Set(field.NewDeliveryType(v))
}

// SetEndAccruedInterestAmt sets EndAccruedInterestAmt, Tag 920.
func (m CollateralRequest) SetEndAccruedInterestAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewEndAccruedInterestAmt(value, scale))
}

// SetStartCash sets StartCash, Tag 921.
func (m CollateralRequest) SetStartCash(value decimal.Decimal, scale int32) {
	m.Set(field.NewStartCash(value, scale))
}

// SetEndCash sets EndCash, Tag 922.
func (m CollateralRequest) SetEndCash(value decimal.Decimal, scale int32) {
	m.Set(field.NewEndCash(value, scale))
}

// SetStrikeCurrency sets StrikeCurrency, Tag 947.
func (m CollateralRequest) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

// SetSecurityStatus sets SecurityStatus, Tag 965.
func (m CollateralRequest) SetSecurityStatus(v enum.SecurityStatus) {
	m.Set(field.NewSecurityStatus(v))
}

// SetSettleOnOpenFlag sets SettleOnOpenFlag, Tag 966.
func (m CollateralRequest) SetSettleOnOpenFlag(v string) {
	m.Set(field.NewSettleOnOpenFlag(v))
}

// SetStrikeMultiplier sets StrikeMultiplier, Tag 967.
func (m CollateralRequest) SetStrikeMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeMultiplier(value, scale))
}

// SetStrikeValue sets StrikeValue, Tag 968.
func (m CollateralRequest) SetStrikeValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeValue(value, scale))
}

// SetMinPriceIncrement sets MinPriceIncrement, Tag 969.
func (m CollateralRequest) SetMinPriceIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrement(value, scale))
}

// SetPositionLimit sets PositionLimit, Tag 970.
func (m CollateralRequest) SetPositionLimit(v int) {
	m.Set(field.NewPositionLimit(v))
}

// SetNTPositionLimit sets NTPositionLimit, Tag 971.
func (m CollateralRequest) SetNTPositionLimit(v int) {
	m.Set(field.NewNTPositionLimit(v))
}

// SetUnitOfMeasure sets UnitOfMeasure, Tag 996.
func (m CollateralRequest) SetUnitOfMeasure(v enum.UnitOfMeasure) {
	m.Set(field.NewUnitOfMeasure(v))
}

// SetTimeUnit sets TimeUnit, Tag 997.
func (m CollateralRequest) SetTimeUnit(v enum.TimeUnit) {
	m.Set(field.NewTimeUnit(v))
}

// SetNoInstrumentParties sets NoInstrumentParties, Tag 1018.
func (m CollateralRequest) SetNoInstrumentParties(f NoInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

// SetInstrmtAssignmentMethod sets InstrmtAssignmentMethod, Tag 1049.
func (m CollateralRequest) SetInstrmtAssignmentMethod(v string) {
	m.Set(field.NewInstrmtAssignmentMethod(v))
}

// SetMaturityTime sets MaturityTime, Tag 1079.
func (m CollateralRequest) SetMaturityTime(v string) {
	m.Set(field.NewMaturityTime(v))
}

// SetMinPriceIncrementAmount sets MinPriceIncrementAmount, Tag 1146.
func (m CollateralRequest) SetMinPriceIncrementAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrementAmount(value, scale))
}

// SetUnitOfMeasureQty sets UnitOfMeasureQty, Tag 1147.
func (m CollateralRequest) SetUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnitOfMeasureQty(value, scale))
}

// SetSecurityGroup sets SecurityGroup, Tag 1151.
func (m CollateralRequest) SetSecurityGroup(v string) {
	m.Set(field.NewSecurityGroup(v))
}

// SetSecurityXMLLen sets SecurityXMLLen, Tag 1184.
func (m CollateralRequest) SetSecurityXMLLen(v int) {
	m.Set(field.NewSecurityXMLLen(v))
}

// SetSecurityXML sets SecurityXML, Tag 1185.
func (m CollateralRequest) SetSecurityXML(v string) {
	m.Set(field.NewSecurityXML(v))
}

// SetSecurityXMLSchema sets SecurityXMLSchema, Tag 1186.
func (m CollateralRequest) SetSecurityXMLSchema(v string) {
	m.Set(field.NewSecurityXMLSchema(v))
}

// SetPriceUnitOfMeasure sets PriceUnitOfMeasure, Tag 1191.
func (m CollateralRequest) SetPriceUnitOfMeasure(v string) {
	m.Set(field.NewPriceUnitOfMeasure(v))
}

// SetPriceUnitOfMeasureQty sets PriceUnitOfMeasureQty, Tag 1192.
func (m CollateralRequest) SetPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewPriceUnitOfMeasureQty(value, scale))
}

// SetSettlMethod sets SettlMethod, Tag 1193.
func (m CollateralRequest) SetSettlMethod(v enum.SettlMethod) {
	m.Set(field.NewSettlMethod(v))
}

// SetExerciseStyle sets ExerciseStyle, Tag 1194.
func (m CollateralRequest) SetExerciseStyle(v enum.ExerciseStyle) {
	m.Set(field.NewExerciseStyle(v))
}

// SetOptPayoutAmount sets OptPayoutAmount, Tag 1195.
func (m CollateralRequest) SetOptPayoutAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewOptPayoutAmount(value, scale))
}

// SetPriceQuoteMethod sets PriceQuoteMethod, Tag 1196.
func (m CollateralRequest) SetPriceQuoteMethod(v enum.PriceQuoteMethod) {
	m.Set(field.NewPriceQuoteMethod(v))
}

// SetValuationMethod sets ValuationMethod, Tag 1197.
func (m CollateralRequest) SetValuationMethod(v enum.ValuationMethod) {
	m.Set(field.NewValuationMethod(v))
}

// SetListMethod sets ListMethod, Tag 1198.
func (m CollateralRequest) SetListMethod(v enum.ListMethod) {
	m.Set(field.NewListMethod(v))
}

// SetCapPrice sets CapPrice, Tag 1199.
func (m CollateralRequest) SetCapPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewCapPrice(value, scale))
}

// SetFloorPrice sets FloorPrice, Tag 1200.
func (m CollateralRequest) SetFloorPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewFloorPrice(value, scale))
}

// SetProductComplex sets ProductComplex, Tag 1227.
func (m CollateralRequest) SetProductComplex(v string) {
	m.Set(field.NewProductComplex(v))
}

// SetFlexProductEligibilityIndicator sets FlexProductEligibilityIndicator, Tag 1242.
func (m CollateralRequest) SetFlexProductEligibilityIndicator(v bool) {
	m.Set(field.NewFlexProductEligibilityIndicator(v))
}

// SetFlexibleIndicator sets FlexibleIndicator, Tag 1244.
func (m CollateralRequest) SetFlexibleIndicator(v bool) {
	m.Set(field.NewFlexibleIndicator(v))
}

// SetContractMultiplierUnit sets ContractMultiplierUnit, Tag 1435.
func (m CollateralRequest) SetContractMultiplierUnit(v enum.ContractMultiplierUnit) {
	m.Set(field.NewContractMultiplierUnit(v))
}

// SetFlowScheduleType sets FlowScheduleType, Tag 1439.
func (m CollateralRequest) SetFlowScheduleType(v enum.FlowScheduleType) {
	m.Set(field.NewFlowScheduleType(v))
}

// SetRestructuringType sets RestructuringType, Tag 1449.
func (m CollateralRequest) SetRestructuringType(v enum.RestructuringType) {
	m.Set(field.NewRestructuringType(v))
}

// SetSeniority sets Seniority, Tag 1450.
func (m CollateralRequest) SetSeniority(v enum.Seniority) {
	m.Set(field.NewSeniority(v))
}

// SetNotionalPercentageOutstanding sets NotionalPercentageOutstanding, Tag 1451.
func (m CollateralRequest) SetNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewNotionalPercentageOutstanding(value, scale))
}

// SetOriginalNotionalPercentageOutstanding sets OriginalNotionalPercentageOutstanding, Tag 1452.
func (m CollateralRequest) SetOriginalNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewOriginalNotionalPercentageOutstanding(value, scale))
}

// SetAttachmentPoint sets AttachmentPoint, Tag 1457.
func (m CollateralRequest) SetAttachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewAttachmentPoint(value, scale))
}

// SetDetachmentPoint sets DetachmentPoint, Tag 1458.
func (m CollateralRequest) SetDetachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewDetachmentPoint(value, scale))
}

// SetStrikePriceDeterminationMethod sets StrikePriceDeterminationMethod, Tag 1478.
func (m CollateralRequest) SetStrikePriceDeterminationMethod(v enum.StrikePriceDeterminationMethod) {
	m.Set(field.NewStrikePriceDeterminationMethod(v))
}

// SetStrikePriceBoundaryMethod sets StrikePriceBoundaryMethod, Tag 1479.
func (m CollateralRequest) SetStrikePriceBoundaryMethod(v enum.StrikePriceBoundaryMethod) {
	m.Set(field.NewStrikePriceBoundaryMethod(v))
}

// SetStrikePriceBoundaryPrecision sets StrikePriceBoundaryPrecision, Tag 1480.
func (m CollateralRequest) SetStrikePriceBoundaryPrecision(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePriceBoundaryPrecision(value, scale))
}

// SetUnderlyingPriceDeterminationMethod sets UnderlyingPriceDeterminationMethod, Tag 1481.
func (m CollateralRequest) SetUnderlyingPriceDeterminationMethod(v enum.UnderlyingPriceDeterminationMethod) {
	m.Set(field.NewUnderlyingPriceDeterminationMethod(v))
}

// SetOptPayoutType sets OptPayoutType, Tag 1482.
func (m CollateralRequest) SetOptPayoutType(v enum.OptPayoutType) {
	m.Set(field.NewOptPayoutType(v))
}

// SetNoComplexEvents sets NoComplexEvents, Tag 1483.
func (m CollateralRequest) SetNoComplexEvents(f NoComplexEventsRepeatingGroup) {
	m.SetGroup(f)
}

// GetAccount gets Account, Tag 1.
func (m CollateralRequest) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetClOrdID gets ClOrdID, Tag 11.
func (m CollateralRequest) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCurrency gets Currency, Tag 15.
func (m CollateralRequest) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityIDSource gets SecurityIDSource, Tag 22.
func (m CollateralRequest) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOrderID gets OrderID, Tag 37.
func (m CollateralRequest) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPrice gets Price, Tag 44.
func (m CollateralRequest) GetPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m CollateralRequest) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuantity gets Quantity, Tag 53.
func (m CollateralRequest) GetQuantity() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.QuantityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSide gets Side, Tag 54.
func (m CollateralRequest) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55.
func (m CollateralRequest) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetText gets Text, Tag 58.
func (m CollateralRequest) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTransactTime gets TransactTime, Tag 60.
func (m CollateralRequest) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlDate gets SettlDate, Tag 64.
func (m CollateralRequest) GetSettlDate() (v string, err quickfix.MessageRejectError) {
	var f field.SettlDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbolSfx gets SymbolSfx, Tag 65.
func (m CollateralRequest) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssuer gets Issuer, Tag 106.
func (m CollateralRequest) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityDesc gets SecurityDesc, Tag 107.
func (m CollateralRequest) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoExecs gets NoExecs, Tag 124.
func (m CollateralRequest) GetNoExecs() (NoExecsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoExecsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetExpireTime gets ExpireTime, Tag 126.
func (m CollateralRequest) GetExpireTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ExpireTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoMiscFees gets NoMiscFees, Tag 136.
func (m CollateralRequest) GetNoMiscFees() (NoMiscFeesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMiscFeesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetAccruedInterestAmt gets AccruedInterestAmt, Tag 159.
func (m CollateralRequest) GetAccruedInterestAmt() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AccruedInterestAmtField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityType gets SecurityType, Tag 167.
func (m CollateralRequest) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecondaryOrderID gets SecondaryOrderID, Tag 198.
func (m CollateralRequest) GetSecondaryOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryOrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityMonthYear gets MaturityMonthYear, Tag 200.
func (m CollateralRequest) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPutOrCall gets PutOrCall, Tag 201.
func (m CollateralRequest) GetPutOrCall() (v enum.PutOrCall, err quickfix.MessageRejectError) {
	var f field.PutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePrice gets StrikePrice, Tag 202.
func (m CollateralRequest) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptAttribute gets OptAttribute, Tag 206.
func (m CollateralRequest) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityExchange gets SecurityExchange, Tag 207.
func (m CollateralRequest) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSpread gets Spread, Tag 218.
func (m CollateralRequest) GetSpread() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SpreadField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkCurveCurrency gets BenchmarkCurveCurrency, Tag 220.
func (m CollateralRequest) GetBenchmarkCurveCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkCurveCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkCurveName gets BenchmarkCurveName, Tag 221.
func (m CollateralRequest) GetBenchmarkCurveName() (v enum.BenchmarkCurveName, err quickfix.MessageRejectError) {
	var f field.BenchmarkCurveNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkCurvePoint gets BenchmarkCurvePoint, Tag 222.
func (m CollateralRequest) GetBenchmarkCurvePoint() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkCurvePointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponRate gets CouponRate, Tag 223.
func (m CollateralRequest) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponPaymentDate gets CouponPaymentDate, Tag 224.
func (m CollateralRequest) GetCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.CouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssueDate gets IssueDate, Tag 225.
func (m CollateralRequest) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseTerm gets RepurchaseTerm, Tag 226.
func (m CollateralRequest) GetRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.RepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseRate gets RepurchaseRate, Tag 227.
func (m CollateralRequest) GetRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFactor gets Factor, Tag 228.
func (m CollateralRequest) GetFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractMultiplier gets ContractMultiplier, Tag 231.
func (m CollateralRequest) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoStipulations gets NoStipulations, Tag 232.
func (m CollateralRequest) GetNoStipulations() (NoStipulationsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoStipulationsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239.
func (m CollateralRequest) GetRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.RepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRedemptionDate gets RedemptionDate, Tag 240.
func (m CollateralRequest) GetRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.RedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCreditRating gets CreditRating, Tag 255.
func (m CollateralRequest) GetCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.CreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m CollateralRequest) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348.
func (m CollateralRequest) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuer gets EncodedIssuer, Tag 349.
func (m CollateralRequest) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350.
func (m CollateralRequest) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351.
func (m CollateralRequest) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedTextLen gets EncodedTextLen, Tag 354.
func (m CollateralRequest) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedText gets EncodedText, Tag 355.
func (m CollateralRequest) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceType gets PriceType, Tag 423.
func (m CollateralRequest) GetPriceType() (v enum.PriceType, err quickfix.MessageRejectError) {
	var f field.PriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m CollateralRequest) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoSecurityAltID gets NoSecurityAltID, Tag 454.
func (m CollateralRequest) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetProduct gets Product, Tag 460.
func (m CollateralRequest) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCFICode gets CFICode, Tag 461.
func (m CollateralRequest) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCountryOfIssue gets CountryOfIssue, Tag 470.
func (m CollateralRequest) GetCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.CountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471.
func (m CollateralRequest) GetStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.StateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLocaleOfIssue gets LocaleOfIssue, Tag 472.
func (m CollateralRequest) GetLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526.
func (m CollateralRequest) GetSecondaryClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityDate gets MaturityDate, Tag 541.
func (m CollateralRequest) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrRegistry gets InstrRegistry, Tag 543.
func (m CollateralRequest) GetInstrRegistry() (v enum.InstrRegistry, err quickfix.MessageRejectError) {
	var f field.InstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoLegs gets NoLegs, Tag 555.
func (m CollateralRequest) GetNoLegs() (NoLegsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetAccountType gets AccountType, Tag 581.
func (m CollateralRequest) GetAccountType() (v enum.AccountType, err quickfix.MessageRejectError) {
	var f field.AccountTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625.
func (m CollateralRequest) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkPrice gets BenchmarkPrice, Tag 662.
func (m CollateralRequest) GetBenchmarkPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BenchmarkPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkPriceType gets BenchmarkPriceType, Tag 663.
func (m CollateralRequest) GetBenchmarkPriceType() (v int, err quickfix.MessageRejectError) {
	var f field.BenchmarkPriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractSettlMonth gets ContractSettlMonth, Tag 667.
func (m CollateralRequest) GetContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.ContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPool gets Pool, Tag 691.
func (m CollateralRequest) GetPool() (v string, err quickfix.MessageRejectError) {
	var f field.PoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkSecurityID gets BenchmarkSecurityID, Tag 699.
func (m CollateralRequest) GetBenchmarkSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUnderlyings gets NoUnderlyings, Tag 711.
func (m CollateralRequest) GetNoUnderlyings() (NoUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetClearingBusinessDate gets ClearingBusinessDate, Tag 715.
func (m CollateralRequest) GetClearingBusinessDate() (v string, err quickfix.MessageRejectError) {
	var f field.ClearingBusinessDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlSessID gets SettlSessID, Tag 716.
func (m CollateralRequest) GetSettlSessID() (v enum.SettlSessID, err quickfix.MessageRejectError) {
	var f field.SettlSessIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlSessSubID gets SettlSessSubID, Tag 717.
func (m CollateralRequest) GetSettlSessSubID() (v string, err quickfix.MessageRejectError) {
	var f field.SettlSessSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetBenchmarkSecurityIDSource gets BenchmarkSecurityIDSource, Tag 761.
func (m CollateralRequest) GetBenchmarkSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.BenchmarkSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecuritySubType gets SecuritySubType, Tag 762.
func (m CollateralRequest) GetSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoTrdRegTimestamps gets NoTrdRegTimestamps, Tag 768.
func (m CollateralRequest) GetNoTrdRegTimestamps() (NoTrdRegTimestampsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTrdRegTimestampsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetTerminationType gets TerminationType, Tag 788.
func (m CollateralRequest) GetTerminationType() (v enum.TerminationType, err quickfix.MessageRejectError) {
	var f field.TerminationTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQtyType gets QtyType, Tag 854.
func (m CollateralRequest) GetQtyType() (v enum.QtyType, err quickfix.MessageRejectError) {
	var f field.QtyTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoEvents gets NoEvents, Tag 864.
func (m CollateralRequest) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetDatedDate gets DatedDate, Tag 873.
func (m CollateralRequest) GetDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.DatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInterestAccrualDate gets InterestAccrualDate, Tag 874.
func (m CollateralRequest) GetInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.InterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCPProgram gets CPProgram, Tag 875.
func (m CollateralRequest) GetCPProgram() (v enum.CPProgram, err quickfix.MessageRejectError) {
	var f field.CPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCPRegType gets CPRegType, Tag 876.
func (m CollateralRequest) GetCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.CPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollReqID gets CollReqID, Tag 894.
func (m CollateralRequest) GetCollReqID() (v string, err quickfix.MessageRejectError) {
	var f field.CollReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollAsgnReason gets CollAsgnReason, Tag 895.
func (m CollateralRequest) GetCollAsgnReason() (v enum.CollAsgnReason, err quickfix.MessageRejectError) {
	var f field.CollAsgnReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoTrades gets NoTrades, Tag 897.
func (m CollateralRequest) GetNoTrades() (NoTradesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTradesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetMarginRatio gets MarginRatio, Tag 898.
func (m CollateralRequest) GetMarginRatio() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MarginRatioField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarginExcess gets MarginExcess, Tag 899.
func (m CollateralRequest) GetMarginExcess() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MarginExcessField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTotalNetValue gets TotalNetValue, Tag 900.
func (m CollateralRequest) GetTotalNetValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TotalNetValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCashOutstanding gets CashOutstanding, Tag 901.
func (m CollateralRequest) GetCashOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CashOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAgreementDesc gets AgreementDesc, Tag 913.
func (m CollateralRequest) GetAgreementDesc() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAgreementID gets AgreementID, Tag 914.
func (m CollateralRequest) GetAgreementID() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAgreementDate gets AgreementDate, Tag 915.
func (m CollateralRequest) GetAgreementDate() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStartDate gets StartDate, Tag 916.
func (m CollateralRequest) GetStartDate() (v string, err quickfix.MessageRejectError) {
	var f field.StartDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEndDate gets EndDate, Tag 917.
func (m CollateralRequest) GetEndDate() (v string, err quickfix.MessageRejectError) {
	var f field.EndDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAgreementCurrency gets AgreementCurrency, Tag 918.
func (m CollateralRequest) GetAgreementCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDeliveryType gets DeliveryType, Tag 919.
func (m CollateralRequest) GetDeliveryType() (v enum.DeliveryType, err quickfix.MessageRejectError) {
	var f field.DeliveryTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEndAccruedInterestAmt gets EndAccruedInterestAmt, Tag 920.
func (m CollateralRequest) GetEndAccruedInterestAmt() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.EndAccruedInterestAmtField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStartCash gets StartCash, Tag 921.
func (m CollateralRequest) GetStartCash() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StartCashField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEndCash gets EndCash, Tag 922.
func (m CollateralRequest) GetEndCash() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.EndCashField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikeCurrency gets StrikeCurrency, Tag 947.
func (m CollateralRequest) GetStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.StrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityStatus gets SecurityStatus, Tag 965.
func (m CollateralRequest) GetSecurityStatus() (v enum.SecurityStatus, err quickfix.MessageRejectError) {
	var f field.SecurityStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettleOnOpenFlag gets SettleOnOpenFlag, Tag 966.
func (m CollateralRequest) GetSettleOnOpenFlag() (v string, err quickfix.MessageRejectError) {
	var f field.SettleOnOpenFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikeMultiplier gets StrikeMultiplier, Tag 967.
func (m CollateralRequest) GetStrikeMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikeMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikeValue gets StrikeValue, Tag 968.
func (m CollateralRequest) GetStrikeValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikeValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMinPriceIncrement gets MinPriceIncrement, Tag 969.
func (m CollateralRequest) GetMinPriceIncrement() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinPriceIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPositionLimit gets PositionLimit, Tag 970.
func (m CollateralRequest) GetPositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.PositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNTPositionLimit gets NTPositionLimit, Tag 971.
func (m CollateralRequest) GetNTPositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.NTPositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnitOfMeasure gets UnitOfMeasure, Tag 996.
func (m CollateralRequest) GetUnitOfMeasure() (v enum.UnitOfMeasure, err quickfix.MessageRejectError) {
	var f field.UnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTimeUnit gets TimeUnit, Tag 997.
func (m CollateralRequest) GetTimeUnit() (v enum.TimeUnit, err quickfix.MessageRejectError) {
	var f field.TimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoInstrumentParties gets NoInstrumentParties, Tag 1018.
func (m CollateralRequest) GetNoInstrumentParties() (NoInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetInstrmtAssignmentMethod gets InstrmtAssignmentMethod, Tag 1049.
func (m CollateralRequest) GetInstrmtAssignmentMethod() (v string, err quickfix.MessageRejectError) {
	var f field.InstrmtAssignmentMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityTime gets MaturityTime, Tag 1079.
func (m CollateralRequest) GetMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMinPriceIncrementAmount gets MinPriceIncrementAmount, Tag 1146.
func (m CollateralRequest) GetMinPriceIncrementAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinPriceIncrementAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnitOfMeasureQty gets UnitOfMeasureQty, Tag 1147.
func (m CollateralRequest) GetUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityGroup gets SecurityGroup, Tag 1151.
func (m CollateralRequest) GetSecurityGroup() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityGroupField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityXMLLen gets SecurityXMLLen, Tag 1184.
func (m CollateralRequest) GetSecurityXMLLen() (v int, err quickfix.MessageRejectError) {
	var f field.SecurityXMLLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityXML gets SecurityXML, Tag 1185.
func (m CollateralRequest) GetSecurityXML() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityXMLField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityXMLSchema gets SecurityXMLSchema, Tag 1186.
func (m CollateralRequest) GetSecurityXMLSchema() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityXMLSchemaField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceUnitOfMeasure gets PriceUnitOfMeasure, Tag 1191.
func (m CollateralRequest) GetPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.PriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceUnitOfMeasureQty gets PriceUnitOfMeasureQty, Tag 1192.
func (m CollateralRequest) GetPriceUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlMethod gets SettlMethod, Tag 1193.
func (m CollateralRequest) GetSettlMethod() (v enum.SettlMethod, err quickfix.MessageRejectError) {
	var f field.SettlMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExerciseStyle gets ExerciseStyle, Tag 1194.
func (m CollateralRequest) GetExerciseStyle() (v enum.ExerciseStyle, err quickfix.MessageRejectError) {
	var f field.ExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptPayoutAmount gets OptPayoutAmount, Tag 1195.
func (m CollateralRequest) GetOptPayoutAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OptPayoutAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceQuoteMethod gets PriceQuoteMethod, Tag 1196.
func (m CollateralRequest) GetPriceQuoteMethod() (v enum.PriceQuoteMethod, err quickfix.MessageRejectError) {
	var f field.PriceQuoteMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetValuationMethod gets ValuationMethod, Tag 1197.
func (m CollateralRequest) GetValuationMethod() (v enum.ValuationMethod, err quickfix.MessageRejectError) {
	var f field.ValuationMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetListMethod gets ListMethod, Tag 1198.
func (m CollateralRequest) GetListMethod() (v enum.ListMethod, err quickfix.MessageRejectError) {
	var f field.ListMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCapPrice gets CapPrice, Tag 1199.
func (m CollateralRequest) GetCapPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CapPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFloorPrice gets FloorPrice, Tag 1200.
func (m CollateralRequest) GetFloorPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FloorPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetProductComplex gets ProductComplex, Tag 1227.
func (m CollateralRequest) GetProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.ProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFlexProductEligibilityIndicator gets FlexProductEligibilityIndicator, Tag 1242.
func (m CollateralRequest) GetFlexProductEligibilityIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.FlexProductEligibilityIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFlexibleIndicator gets FlexibleIndicator, Tag 1244.
func (m CollateralRequest) GetFlexibleIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.FlexibleIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractMultiplierUnit gets ContractMultiplierUnit, Tag 1435.
func (m CollateralRequest) GetContractMultiplierUnit() (v enum.ContractMultiplierUnit, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFlowScheduleType gets FlowScheduleType, Tag 1439.
func (m CollateralRequest) GetFlowScheduleType() (v enum.FlowScheduleType, err quickfix.MessageRejectError) {
	var f field.FlowScheduleTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRestructuringType gets RestructuringType, Tag 1449.
func (m CollateralRequest) GetRestructuringType() (v enum.RestructuringType, err quickfix.MessageRejectError) {
	var f field.RestructuringTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSeniority gets Seniority, Tag 1450.
func (m CollateralRequest) GetSeniority() (v enum.Seniority, err quickfix.MessageRejectError) {
	var f field.SeniorityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNotionalPercentageOutstanding gets NotionalPercentageOutstanding, Tag 1451.
func (m CollateralRequest) GetNotionalPercentageOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.NotionalPercentageOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOriginalNotionalPercentageOutstanding gets OriginalNotionalPercentageOutstanding, Tag 1452.
func (m CollateralRequest) GetOriginalNotionalPercentageOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OriginalNotionalPercentageOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAttachmentPoint gets AttachmentPoint, Tag 1457.
func (m CollateralRequest) GetAttachmentPoint() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AttachmentPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDetachmentPoint gets DetachmentPoint, Tag 1458.
func (m CollateralRequest) GetDetachmentPoint() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DetachmentPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePriceDeterminationMethod gets StrikePriceDeterminationMethod, Tag 1478.
func (m CollateralRequest) GetStrikePriceDeterminationMethod() (v enum.StrikePriceDeterminationMethod, err quickfix.MessageRejectError) {
	var f field.StrikePriceDeterminationMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePriceBoundaryMethod gets StrikePriceBoundaryMethod, Tag 1479.
func (m CollateralRequest) GetStrikePriceBoundaryMethod() (v enum.StrikePriceBoundaryMethod, err quickfix.MessageRejectError) {
	var f field.StrikePriceBoundaryMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePriceBoundaryPrecision gets StrikePriceBoundaryPrecision, Tag 1480.
func (m CollateralRequest) GetStrikePriceBoundaryPrecision() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceBoundaryPrecisionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPriceDeterminationMethod gets UnderlyingPriceDeterminationMethod, Tag 1481.
func (m CollateralRequest) GetUnderlyingPriceDeterminationMethod() (v enum.UnderlyingPriceDeterminationMethod, err quickfix.MessageRejectError) {
	var f field.UnderlyingPriceDeterminationMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptPayoutType gets OptPayoutType, Tag 1482.
func (m CollateralRequest) GetOptPayoutType() (v enum.OptPayoutType, err quickfix.MessageRejectError) {
	var f field.OptPayoutTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoComplexEvents gets NoComplexEvents, Tag 1483.
func (m CollateralRequest) GetNoComplexEvents() (NoComplexEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasAccount returns true if Account is present, Tag 1.
func (m CollateralRequest) HasAccount() bool {
	return m.Has(tag.Account)
}

// HasClOrdID returns true if ClOrdID is present, Tag 11.
func (m CollateralRequest) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

// HasCurrency returns true if Currency is present, Tag 15.
func (m CollateralRequest) HasCurrency() bool {
	return m.Has(tag.Currency)
}

// HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22.
func (m CollateralRequest) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

// HasOrderID returns true if OrderID is present, Tag 37.
func (m CollateralRequest) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

// HasPrice returns true if Price is present, Tag 44.
func (m CollateralRequest) HasPrice() bool {
	return m.Has(tag.Price)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m CollateralRequest) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasQuantity returns true if Quantity is present, Tag 53.
func (m CollateralRequest) HasQuantity() bool {
	return m.Has(tag.Quantity)
}

// HasSide returns true if Side is present, Tag 54.
func (m CollateralRequest) HasSide() bool {
	return m.Has(tag.Side)
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m CollateralRequest) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasText returns true if Text is present, Tag 58.
func (m CollateralRequest) HasText() bool {
	return m.Has(tag.Text)
}

// HasTransactTime returns true if TransactTime is present, Tag 60.
func (m CollateralRequest) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

// HasSettlDate returns true if SettlDate is present, Tag 64.
func (m CollateralRequest) HasSettlDate() bool {
	return m.Has(tag.SettlDate)
}

// HasSymbolSfx returns true if SymbolSfx is present, Tag 65.
func (m CollateralRequest) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

// HasIssuer returns true if Issuer is present, Tag 106.
func (m CollateralRequest) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

// HasSecurityDesc returns true if SecurityDesc is present, Tag 107.
func (m CollateralRequest) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

// HasNoExecs returns true if NoExecs is present, Tag 124.
func (m CollateralRequest) HasNoExecs() bool {
	return m.Has(tag.NoExecs)
}

// HasExpireTime returns true if ExpireTime is present, Tag 126.
func (m CollateralRequest) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

// HasNoMiscFees returns true if NoMiscFees is present, Tag 136.
func (m CollateralRequest) HasNoMiscFees() bool {
	return m.Has(tag.NoMiscFees)
}

// HasAccruedInterestAmt returns true if AccruedInterestAmt is present, Tag 159.
func (m CollateralRequest) HasAccruedInterestAmt() bool {
	return m.Has(tag.AccruedInterestAmt)
}

// HasSecurityType returns true if SecurityType is present, Tag 167.
func (m CollateralRequest) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

// HasSecondaryOrderID returns true if SecondaryOrderID is present, Tag 198.
func (m CollateralRequest) HasSecondaryOrderID() bool {
	return m.Has(tag.SecondaryOrderID)
}

// HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200.
func (m CollateralRequest) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

// HasPutOrCall returns true if PutOrCall is present, Tag 201.
func (m CollateralRequest) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

// HasStrikePrice returns true if StrikePrice is present, Tag 202.
func (m CollateralRequest) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

// HasOptAttribute returns true if OptAttribute is present, Tag 206.
func (m CollateralRequest) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

// HasSecurityExchange returns true if SecurityExchange is present, Tag 207.
func (m CollateralRequest) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

// HasSpread returns true if Spread is present, Tag 218.
func (m CollateralRequest) HasSpread() bool {
	return m.Has(tag.Spread)
}

// HasBenchmarkCurveCurrency returns true if BenchmarkCurveCurrency is present, Tag 220.
func (m CollateralRequest) HasBenchmarkCurveCurrency() bool {
	return m.Has(tag.BenchmarkCurveCurrency)
}

// HasBenchmarkCurveName returns true if BenchmarkCurveName is present, Tag 221.
func (m CollateralRequest) HasBenchmarkCurveName() bool {
	return m.Has(tag.BenchmarkCurveName)
}

// HasBenchmarkCurvePoint returns true if BenchmarkCurvePoint is present, Tag 222.
func (m CollateralRequest) HasBenchmarkCurvePoint() bool {
	return m.Has(tag.BenchmarkCurvePoint)
}

// HasCouponRate returns true if CouponRate is present, Tag 223.
func (m CollateralRequest) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

// HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224.
func (m CollateralRequest) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

// HasIssueDate returns true if IssueDate is present, Tag 225.
func (m CollateralRequest) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

// HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226.
func (m CollateralRequest) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

// HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227.
func (m CollateralRequest) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

// HasFactor returns true if Factor is present, Tag 228.
func (m CollateralRequest) HasFactor() bool {
	return m.Has(tag.Factor)
}

// HasContractMultiplier returns true if ContractMultiplier is present, Tag 231.
func (m CollateralRequest) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

// HasNoStipulations returns true if NoStipulations is present, Tag 232.
func (m CollateralRequest) HasNoStipulations() bool {
	return m.Has(tag.NoStipulations)
}

// HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239.
func (m CollateralRequest) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

// HasRedemptionDate returns true if RedemptionDate is present, Tag 240.
func (m CollateralRequest) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

// HasCreditRating returns true if CreditRating is present, Tag 255.
func (m CollateralRequest) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m CollateralRequest) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348.
func (m CollateralRequest) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

// HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349.
func (m CollateralRequest) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

// HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350.
func (m CollateralRequest) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

// HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351.
func (m CollateralRequest) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

// HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354.
func (m CollateralRequest) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

// HasEncodedText returns true if EncodedText is present, Tag 355.
func (m CollateralRequest) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

// HasPriceType returns true if PriceType is present, Tag 423.
func (m CollateralRequest) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m CollateralRequest) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454.
func (m CollateralRequest) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

// HasProduct returns true if Product is present, Tag 460.
func (m CollateralRequest) HasProduct() bool {
	return m.Has(tag.Product)
}

// HasCFICode returns true if CFICode is present, Tag 461.
func (m CollateralRequest) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

// HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470.
func (m CollateralRequest) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

// HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471.
func (m CollateralRequest) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

// HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472.
func (m CollateralRequest) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

// HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526.
func (m CollateralRequest) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

// HasMaturityDate returns true if MaturityDate is present, Tag 541.
func (m CollateralRequest) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

// HasInstrRegistry returns true if InstrRegistry is present, Tag 543.
func (m CollateralRequest) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

// HasNoLegs returns true if NoLegs is present, Tag 555.
func (m CollateralRequest) HasNoLegs() bool {
	return m.Has(tag.NoLegs)
}

// HasAccountType returns true if AccountType is present, Tag 581.
func (m CollateralRequest) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625.
func (m CollateralRequest) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

// HasBenchmarkPrice returns true if BenchmarkPrice is present, Tag 662.
func (m CollateralRequest) HasBenchmarkPrice() bool {
	return m.Has(tag.BenchmarkPrice)
}

// HasBenchmarkPriceType returns true if BenchmarkPriceType is present, Tag 663.
func (m CollateralRequest) HasBenchmarkPriceType() bool {
	return m.Has(tag.BenchmarkPriceType)
}

// HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667.
func (m CollateralRequest) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

// HasPool returns true if Pool is present, Tag 691.
func (m CollateralRequest) HasPool() bool {
	return m.Has(tag.Pool)
}

// HasBenchmarkSecurityID returns true if BenchmarkSecurityID is present, Tag 699.
func (m CollateralRequest) HasBenchmarkSecurityID() bool {
	return m.Has(tag.BenchmarkSecurityID)
}

// HasNoUnderlyings returns true if NoUnderlyings is present, Tag 711.
func (m CollateralRequest) HasNoUnderlyings() bool {
	return m.Has(tag.NoUnderlyings)
}

// HasClearingBusinessDate returns true if ClearingBusinessDate is present, Tag 715.
func (m CollateralRequest) HasClearingBusinessDate() bool {
	return m.Has(tag.ClearingBusinessDate)
}

// HasSettlSessID returns true if SettlSessID is present, Tag 716.
func (m CollateralRequest) HasSettlSessID() bool {
	return m.Has(tag.SettlSessID)
}

// HasSettlSessSubID returns true if SettlSessSubID is present, Tag 717.
func (m CollateralRequest) HasSettlSessSubID() bool {
	return m.Has(tag.SettlSessSubID)
}

// HasBenchmarkSecurityIDSource returns true if BenchmarkSecurityIDSource is present, Tag 761.
func (m CollateralRequest) HasBenchmarkSecurityIDSource() bool {
	return m.Has(tag.BenchmarkSecurityIDSource)
}

// HasSecuritySubType returns true if SecuritySubType is present, Tag 762.
func (m CollateralRequest) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

// HasNoTrdRegTimestamps returns true if NoTrdRegTimestamps is present, Tag 768.
func (m CollateralRequest) HasNoTrdRegTimestamps() bool {
	return m.Has(tag.NoTrdRegTimestamps)
}

// HasTerminationType returns true if TerminationType is present, Tag 788.
func (m CollateralRequest) HasTerminationType() bool {
	return m.Has(tag.TerminationType)
}

// HasQtyType returns true if QtyType is present, Tag 854.
func (m CollateralRequest) HasQtyType() bool {
	return m.Has(tag.QtyType)
}

// HasNoEvents returns true if NoEvents is present, Tag 864.
func (m CollateralRequest) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

// HasDatedDate returns true if DatedDate is present, Tag 873.
func (m CollateralRequest) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

// HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874.
func (m CollateralRequest) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

// HasCPProgram returns true if CPProgram is present, Tag 875.
func (m CollateralRequest) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

// HasCPRegType returns true if CPRegType is present, Tag 876.
func (m CollateralRequest) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

// HasCollReqID returns true if CollReqID is present, Tag 894.
func (m CollateralRequest) HasCollReqID() bool {
	return m.Has(tag.CollReqID)
}

// HasCollAsgnReason returns true if CollAsgnReason is present, Tag 895.
func (m CollateralRequest) HasCollAsgnReason() bool {
	return m.Has(tag.CollAsgnReason)
}

// HasNoTrades returns true if NoTrades is present, Tag 897.
func (m CollateralRequest) HasNoTrades() bool {
	return m.Has(tag.NoTrades)
}

// HasMarginRatio returns true if MarginRatio is present, Tag 898.
func (m CollateralRequest) HasMarginRatio() bool {
	return m.Has(tag.MarginRatio)
}

// HasMarginExcess returns true if MarginExcess is present, Tag 899.
func (m CollateralRequest) HasMarginExcess() bool {
	return m.Has(tag.MarginExcess)
}

// HasTotalNetValue returns true if TotalNetValue is present, Tag 900.
func (m CollateralRequest) HasTotalNetValue() bool {
	return m.Has(tag.TotalNetValue)
}

// HasCashOutstanding returns true if CashOutstanding is present, Tag 901.
func (m CollateralRequest) HasCashOutstanding() bool {
	return m.Has(tag.CashOutstanding)
}

// HasAgreementDesc returns true if AgreementDesc is present, Tag 913.
func (m CollateralRequest) HasAgreementDesc() bool {
	return m.Has(tag.AgreementDesc)
}

// HasAgreementID returns true if AgreementID is present, Tag 914.
func (m CollateralRequest) HasAgreementID() bool {
	return m.Has(tag.AgreementID)
}

// HasAgreementDate returns true if AgreementDate is present, Tag 915.
func (m CollateralRequest) HasAgreementDate() bool {
	return m.Has(tag.AgreementDate)
}

// HasStartDate returns true if StartDate is present, Tag 916.
func (m CollateralRequest) HasStartDate() bool {
	return m.Has(tag.StartDate)
}

// HasEndDate returns true if EndDate is present, Tag 917.
func (m CollateralRequest) HasEndDate() bool {
	return m.Has(tag.EndDate)
}

// HasAgreementCurrency returns true if AgreementCurrency is present, Tag 918.
func (m CollateralRequest) HasAgreementCurrency() bool {
	return m.Has(tag.AgreementCurrency)
}

// HasDeliveryType returns true if DeliveryType is present, Tag 919.
func (m CollateralRequest) HasDeliveryType() bool {
	return m.Has(tag.DeliveryType)
}

// HasEndAccruedInterestAmt returns true if EndAccruedInterestAmt is present, Tag 920.
func (m CollateralRequest) HasEndAccruedInterestAmt() bool {
	return m.Has(tag.EndAccruedInterestAmt)
}

// HasStartCash returns true if StartCash is present, Tag 921.
func (m CollateralRequest) HasStartCash() bool {
	return m.Has(tag.StartCash)
}

// HasEndCash returns true if EndCash is present, Tag 922.
func (m CollateralRequest) HasEndCash() bool {
	return m.Has(tag.EndCash)
}

// HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947.
func (m CollateralRequest) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

// HasSecurityStatus returns true if SecurityStatus is present, Tag 965.
func (m CollateralRequest) HasSecurityStatus() bool {
	return m.Has(tag.SecurityStatus)
}

// HasSettleOnOpenFlag returns true if SettleOnOpenFlag is present, Tag 966.
func (m CollateralRequest) HasSettleOnOpenFlag() bool {
	return m.Has(tag.SettleOnOpenFlag)
}

// HasStrikeMultiplier returns true if StrikeMultiplier is present, Tag 967.
func (m CollateralRequest) HasStrikeMultiplier() bool {
	return m.Has(tag.StrikeMultiplier)
}

// HasStrikeValue returns true if StrikeValue is present, Tag 968.
func (m CollateralRequest) HasStrikeValue() bool {
	return m.Has(tag.StrikeValue)
}

// HasMinPriceIncrement returns true if MinPriceIncrement is present, Tag 969.
func (m CollateralRequest) HasMinPriceIncrement() bool {
	return m.Has(tag.MinPriceIncrement)
}

// HasPositionLimit returns true if PositionLimit is present, Tag 970.
func (m CollateralRequest) HasPositionLimit() bool {
	return m.Has(tag.PositionLimit)
}

// HasNTPositionLimit returns true if NTPositionLimit is present, Tag 971.
func (m CollateralRequest) HasNTPositionLimit() bool {
	return m.Has(tag.NTPositionLimit)
}

// HasUnitOfMeasure returns true if UnitOfMeasure is present, Tag 996.
func (m CollateralRequest) HasUnitOfMeasure() bool {
	return m.Has(tag.UnitOfMeasure)
}

// HasTimeUnit returns true if TimeUnit is present, Tag 997.
func (m CollateralRequest) HasTimeUnit() bool {
	return m.Has(tag.TimeUnit)
}

// HasNoInstrumentParties returns true if NoInstrumentParties is present, Tag 1018.
func (m CollateralRequest) HasNoInstrumentParties() bool {
	return m.Has(tag.NoInstrumentParties)
}

// HasInstrmtAssignmentMethod returns true if InstrmtAssignmentMethod is present, Tag 1049.
func (m CollateralRequest) HasInstrmtAssignmentMethod() bool {
	return m.Has(tag.InstrmtAssignmentMethod)
}

// HasMaturityTime returns true if MaturityTime is present, Tag 1079.
func (m CollateralRequest) HasMaturityTime() bool {
	return m.Has(tag.MaturityTime)
}

// HasMinPriceIncrementAmount returns true if MinPriceIncrementAmount is present, Tag 1146.
func (m CollateralRequest) HasMinPriceIncrementAmount() bool {
	return m.Has(tag.MinPriceIncrementAmount)
}

// HasUnitOfMeasureQty returns true if UnitOfMeasureQty is present, Tag 1147.
func (m CollateralRequest) HasUnitOfMeasureQty() bool {
	return m.Has(tag.UnitOfMeasureQty)
}

// HasSecurityGroup returns true if SecurityGroup is present, Tag 1151.
func (m CollateralRequest) HasSecurityGroup() bool {
	return m.Has(tag.SecurityGroup)
}

// HasSecurityXMLLen returns true if SecurityXMLLen is present, Tag 1184.
func (m CollateralRequest) HasSecurityXMLLen() bool {
	return m.Has(tag.SecurityXMLLen)
}

// HasSecurityXML returns true if SecurityXML is present, Tag 1185.
func (m CollateralRequest) HasSecurityXML() bool {
	return m.Has(tag.SecurityXML)
}

// HasSecurityXMLSchema returns true if SecurityXMLSchema is present, Tag 1186.
func (m CollateralRequest) HasSecurityXMLSchema() bool {
	return m.Has(tag.SecurityXMLSchema)
}

// HasPriceUnitOfMeasure returns true if PriceUnitOfMeasure is present, Tag 1191.
func (m CollateralRequest) HasPriceUnitOfMeasure() bool {
	return m.Has(tag.PriceUnitOfMeasure)
}

// HasPriceUnitOfMeasureQty returns true if PriceUnitOfMeasureQty is present, Tag 1192.
func (m CollateralRequest) HasPriceUnitOfMeasureQty() bool {
	return m.Has(tag.PriceUnitOfMeasureQty)
}

// HasSettlMethod returns true if SettlMethod is present, Tag 1193.
func (m CollateralRequest) HasSettlMethod() bool {
	return m.Has(tag.SettlMethod)
}

// HasExerciseStyle returns true if ExerciseStyle is present, Tag 1194.
func (m CollateralRequest) HasExerciseStyle() bool {
	return m.Has(tag.ExerciseStyle)
}

// HasOptPayoutAmount returns true if OptPayoutAmount is present, Tag 1195.
func (m CollateralRequest) HasOptPayoutAmount() bool {
	return m.Has(tag.OptPayoutAmount)
}

// HasPriceQuoteMethod returns true if PriceQuoteMethod is present, Tag 1196.
func (m CollateralRequest) HasPriceQuoteMethod() bool {
	return m.Has(tag.PriceQuoteMethod)
}

// HasValuationMethod returns true if ValuationMethod is present, Tag 1197.
func (m CollateralRequest) HasValuationMethod() bool {
	return m.Has(tag.ValuationMethod)
}

// HasListMethod returns true if ListMethod is present, Tag 1198.
func (m CollateralRequest) HasListMethod() bool {
	return m.Has(tag.ListMethod)
}

// HasCapPrice returns true if CapPrice is present, Tag 1199.
func (m CollateralRequest) HasCapPrice() bool {
	return m.Has(tag.CapPrice)
}

// HasFloorPrice returns true if FloorPrice is present, Tag 1200.
func (m CollateralRequest) HasFloorPrice() bool {
	return m.Has(tag.FloorPrice)
}

// HasProductComplex returns true if ProductComplex is present, Tag 1227.
func (m CollateralRequest) HasProductComplex() bool {
	return m.Has(tag.ProductComplex)
}

// HasFlexProductEligibilityIndicator returns true if FlexProductEligibilityIndicator is present, Tag 1242.
func (m CollateralRequest) HasFlexProductEligibilityIndicator() bool {
	return m.Has(tag.FlexProductEligibilityIndicator)
}

// HasFlexibleIndicator returns true if FlexibleIndicator is present, Tag 1244.
func (m CollateralRequest) HasFlexibleIndicator() bool {
	return m.Has(tag.FlexibleIndicator)
}

// HasContractMultiplierUnit returns true if ContractMultiplierUnit is present, Tag 1435.
func (m CollateralRequest) HasContractMultiplierUnit() bool {
	return m.Has(tag.ContractMultiplierUnit)
}

// HasFlowScheduleType returns true if FlowScheduleType is present, Tag 1439.
func (m CollateralRequest) HasFlowScheduleType() bool {
	return m.Has(tag.FlowScheduleType)
}

// HasRestructuringType returns true if RestructuringType is present, Tag 1449.
func (m CollateralRequest) HasRestructuringType() bool {
	return m.Has(tag.RestructuringType)
}

// HasSeniority returns true if Seniority is present, Tag 1450.
func (m CollateralRequest) HasSeniority() bool {
	return m.Has(tag.Seniority)
}

// HasNotionalPercentageOutstanding returns true if NotionalPercentageOutstanding is present, Tag 1451.
func (m CollateralRequest) HasNotionalPercentageOutstanding() bool {
	return m.Has(tag.NotionalPercentageOutstanding)
}

// HasOriginalNotionalPercentageOutstanding returns true if OriginalNotionalPercentageOutstanding is present, Tag 1452.
func (m CollateralRequest) HasOriginalNotionalPercentageOutstanding() bool {
	return m.Has(tag.OriginalNotionalPercentageOutstanding)
}

// HasAttachmentPoint returns true if AttachmentPoint is present, Tag 1457.
func (m CollateralRequest) HasAttachmentPoint() bool {
	return m.Has(tag.AttachmentPoint)
}

// HasDetachmentPoint returns true if DetachmentPoint is present, Tag 1458.
func (m CollateralRequest) HasDetachmentPoint() bool {
	return m.Has(tag.DetachmentPoint)
}

// HasStrikePriceDeterminationMethod returns true if StrikePriceDeterminationMethod is present, Tag 1478.
func (m CollateralRequest) HasStrikePriceDeterminationMethod() bool {
	return m.Has(tag.StrikePriceDeterminationMethod)
}

// HasStrikePriceBoundaryMethod returns true if StrikePriceBoundaryMethod is present, Tag 1479.
func (m CollateralRequest) HasStrikePriceBoundaryMethod() bool {
	return m.Has(tag.StrikePriceBoundaryMethod)
}

// HasStrikePriceBoundaryPrecision returns true if StrikePriceBoundaryPrecision is present, Tag 1480.
func (m CollateralRequest) HasStrikePriceBoundaryPrecision() bool {
	return m.Has(tag.StrikePriceBoundaryPrecision)
}

// HasUnderlyingPriceDeterminationMethod returns true if UnderlyingPriceDeterminationMethod is present, Tag 1481.
func (m CollateralRequest) HasUnderlyingPriceDeterminationMethod() bool {
	return m.Has(tag.UnderlyingPriceDeterminationMethod)
}

// HasOptPayoutType returns true if OptPayoutType is present, Tag 1482.
func (m CollateralRequest) HasOptPayoutType() bool {
	return m.Has(tag.OptPayoutType)
}

// HasNoComplexEvents returns true if NoComplexEvents is present, Tag 1483.
func (m CollateralRequest) HasNoComplexEvents() bool {
	return m.Has(tag.NoComplexEvents)
}

// NoExecs is a repeating group element, Tag 124.
type NoExecs struct {
	*quickfix.Group
}

// SetExecID sets ExecID, Tag 17.
func (m NoExecs) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

// GetExecID gets ExecID, Tag 17.
func (m NoExecs) GetExecID() (v string, err quickfix.MessageRejectError) {
	var f field.ExecIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasExecID returns true if ExecID is present, Tag 17.
func (m NoExecs) HasExecID() bool {
	return m.Has(tag.ExecID)
}

// NoExecsRepeatingGroup is a repeating group, Tag 124.
type NoExecsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoExecsRepeatingGroup returns an initialized, NoExecsRepeatingGroup.
func NewNoExecsRepeatingGroup() NoExecsRepeatingGroup {
	return NoExecsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoExecs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ExecID)})}
}

// Add create and append a new NoExecs to this group.
func (m NoExecsRepeatingGroup) Add() NoExecs {
	g := m.RepeatingGroup.Add()
	return NoExecs{g}
}

// Get returns the ith NoExecs in the NoExecsRepeatinGroup.
func (m NoExecsRepeatingGroup) Get(i int) NoExecs {
	return NoExecs{m.RepeatingGroup.Get(i)}
}

// NoMiscFees is a repeating group element, Tag 136.
type NoMiscFees struct {
	*quickfix.Group
}

// SetMiscFeeAmt sets MiscFeeAmt, Tag 137.
func (m NoMiscFees) SetMiscFeeAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewMiscFeeAmt(value, scale))
}

// SetMiscFeeCurr sets MiscFeeCurr, Tag 138.
func (m NoMiscFees) SetMiscFeeCurr(v string) {
	m.Set(field.NewMiscFeeCurr(v))
}

// SetMiscFeeType sets MiscFeeType, Tag 139.
func (m NoMiscFees) SetMiscFeeType(v enum.MiscFeeType) {
	m.Set(field.NewMiscFeeType(v))
}

// SetMiscFeeBasis sets MiscFeeBasis, Tag 891.
func (m NoMiscFees) SetMiscFeeBasis(v enum.MiscFeeBasis) {
	m.Set(field.NewMiscFeeBasis(v))
}

// GetMiscFeeAmt gets MiscFeeAmt, Tag 137.
func (m NoMiscFees) GetMiscFeeAmt() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MiscFeeAmtField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMiscFeeCurr gets MiscFeeCurr, Tag 138.
func (m NoMiscFees) GetMiscFeeCurr() (v string, err quickfix.MessageRejectError) {
	var f field.MiscFeeCurrField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMiscFeeType gets MiscFeeType, Tag 139.
func (m NoMiscFees) GetMiscFeeType() (v enum.MiscFeeType, err quickfix.MessageRejectError) {
	var f field.MiscFeeTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMiscFeeBasis gets MiscFeeBasis, Tag 891.
func (m NoMiscFees) GetMiscFeeBasis() (v enum.MiscFeeBasis, err quickfix.MessageRejectError) {
	var f field.MiscFeeBasisField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasMiscFeeAmt returns true if MiscFeeAmt is present, Tag 137.
func (m NoMiscFees) HasMiscFeeAmt() bool {
	return m.Has(tag.MiscFeeAmt)
}

// HasMiscFeeCurr returns true if MiscFeeCurr is present, Tag 138.
func (m NoMiscFees) HasMiscFeeCurr() bool {
	return m.Has(tag.MiscFeeCurr)
}

// HasMiscFeeType returns true if MiscFeeType is present, Tag 139.
func (m NoMiscFees) HasMiscFeeType() bool {
	return m.Has(tag.MiscFeeType)
}

// HasMiscFeeBasis returns true if MiscFeeBasis is present, Tag 891.
func (m NoMiscFees) HasMiscFeeBasis() bool {
	return m.Has(tag.MiscFeeBasis)
}

// NoMiscFeesRepeatingGroup is a repeating group, Tag 136.
type NoMiscFeesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoMiscFeesRepeatingGroup returns an initialized, NoMiscFeesRepeatingGroup.
func NewNoMiscFeesRepeatingGroup() NoMiscFeesRepeatingGroup {
	return NoMiscFeesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMiscFees,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MiscFeeAmt), quickfix.GroupElement(tag.MiscFeeCurr), quickfix.GroupElement(tag.MiscFeeType), quickfix.GroupElement(tag.MiscFeeBasis)})}
}

// Add create and append a new NoMiscFees to this group.
func (m NoMiscFeesRepeatingGroup) Add() NoMiscFees {
	g := m.RepeatingGroup.Add()
	return NoMiscFees{g}
}

// Get returns the ith NoMiscFees in the NoMiscFeesRepeatinGroup.
func (m NoMiscFeesRepeatingGroup) Get(i int) NoMiscFees {
	return NoMiscFees{m.RepeatingGroup.Get(i)}
}

// NoStipulations is a repeating group element, Tag 232.
type NoStipulations struct {
	*quickfix.Group
}

// SetStipulationType sets StipulationType, Tag 233.
func (m NoStipulations) SetStipulationType(v enum.StipulationType) {
	m.Set(field.NewStipulationType(v))
}

// SetStipulationValue sets StipulationValue, Tag 234.
func (m NoStipulations) SetStipulationValue(v string) {
	m.Set(field.NewStipulationValue(v))
}

// GetStipulationType gets StipulationType, Tag 233.
func (m NoStipulations) GetStipulationType() (v enum.StipulationType, err quickfix.MessageRejectError) {
	var f field.StipulationTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStipulationValue gets StipulationValue, Tag 234.
func (m NoStipulations) GetStipulationValue() (v string, err quickfix.MessageRejectError) {
	var f field.StipulationValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasStipulationType returns true if StipulationType is present, Tag 233.
func (m NoStipulations) HasStipulationType() bool {
	return m.Has(tag.StipulationType)
}

// HasStipulationValue returns true if StipulationValue is present, Tag 234.
func (m NoStipulations) HasStipulationValue() bool {
	return m.Has(tag.StipulationValue)
}

// NoStipulationsRepeatingGroup is a repeating group, Tag 232.
type NoStipulationsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoStipulationsRepeatingGroup returns an initialized, NoStipulationsRepeatingGroup.
func NewNoStipulationsRepeatingGroup() NoStipulationsRepeatingGroup {
	return NoStipulationsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoStipulations,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.StipulationType), quickfix.GroupElement(tag.StipulationValue)})}
}

// Add create and append a new NoStipulations to this group.
func (m NoStipulationsRepeatingGroup) Add() NoStipulations {
	g := m.RepeatingGroup.Add()
	return NoStipulations{g}
}

// Get returns the ith NoStipulations in the NoStipulationsRepeatinGroup.
func (m NoStipulationsRepeatingGroup) Get(i int) NoStipulations {
	return NoStipulations{m.RepeatingGroup.Get(i)}
}

// NoPartyIDs is a repeating group element, Tag 453.
type NoPartyIDs struct {
	*quickfix.Group
}

// SetPartyID sets PartyID, Tag 448.
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

// SetPartyIDSource sets PartyIDSource, Tag 447.
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

// SetPartyRole sets PartyRole, Tag 452.
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

// SetNoPartySubIDs sets NoPartySubIDs, Tag 802.
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetPartyID gets PartyID, Tag 448.
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyIDSource gets PartyIDSource, Tag 447.
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyRole gets PartyRole, Tag 452.
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartySubIDs gets NoPartySubIDs, Tag 802.
func (m NoPartyIDs) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasPartyID returns true if PartyID is present, Tag 448.
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

// HasPartyIDSource returns true if PartyIDSource is present, Tag 447.
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

// HasPartyRole returns true if PartyRole is present, Tag 452.
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

// HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802.
func (m NoPartyIDs) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

// NoPartySubIDs is a repeating group element, Tag 802.
type NoPartySubIDs struct {
	*quickfix.Group
}

// SetPartySubID sets PartySubID, Tag 523.
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

// SetPartySubIDType sets PartySubIDType, Tag 803.
func (m NoPartySubIDs) SetPartySubIDType(v enum.PartySubIDType) {
	m.Set(field.NewPartySubIDType(v))
}

// GetPartySubID gets PartySubID, Tag 523.
func (m NoPartySubIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartySubIDType gets PartySubIDType, Tag 803.
func (m NoPartySubIDs) GetPartySubIDType() (v enum.PartySubIDType, err quickfix.MessageRejectError) {
	var f field.PartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartySubID returns true if PartySubID is present, Tag 523.
func (m NoPartySubIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

// HasPartySubIDType returns true if PartySubIDType is present, Tag 803.
func (m NoPartySubIDs) HasPartySubIDType() bool {
	return m.Has(tag.PartySubIDType)
}

// NoPartySubIDsRepeatingGroup is a repeating group, Tag 802.
type NoPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartySubIDsRepeatingGroup returns an initialized, NoPartySubIDsRepeatingGroup.
func NewNoPartySubIDsRepeatingGroup() NoPartySubIDsRepeatingGroup {
	return NoPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartySubID), quickfix.GroupElement(tag.PartySubIDType)})}
}

// Add create and append a new NoPartySubIDs to this group.
func (m NoPartySubIDsRepeatingGroup) Add() NoPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartySubIDs{g}
}

// Get returns the ith NoPartySubIDs in the NoPartySubIDsRepeatinGroup.
func (m NoPartySubIDsRepeatingGroup) Get(i int) NoPartySubIDs {
	return NoPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoPartyIDsRepeatingGroup is a repeating group, Tag 453.
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup.
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), NewNoPartySubIDsRepeatingGroup()})}
}

// Add create and append a new NoPartyIDs to this group.
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

// Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup.
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}

// NoSecurityAltID is a repeating group element, Tag 454.
type NoSecurityAltID struct {
	*quickfix.Group
}

// SetSecurityAltID sets SecurityAltID, Tag 455.
func (m NoSecurityAltID) SetSecurityAltID(v string) {
	m.Set(field.NewSecurityAltID(v))
}

// SetSecurityAltIDSource sets SecurityAltIDSource, Tag 456.
func (m NoSecurityAltID) SetSecurityAltIDSource(v string) {
	m.Set(field.NewSecurityAltIDSource(v))
}

// GetSecurityAltID gets SecurityAltID, Tag 455.
func (m NoSecurityAltID) GetSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456.
func (m NoSecurityAltID) GetSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasSecurityAltID returns true if SecurityAltID is present, Tag 455.
func (m NoSecurityAltID) HasSecurityAltID() bool {
	return m.Has(tag.SecurityAltID)
}

// HasSecurityAltIDSource returns true if SecurityAltIDSource is present, Tag 456.
func (m NoSecurityAltID) HasSecurityAltIDSource() bool {
	return m.Has(tag.SecurityAltIDSource)
}

// NoSecurityAltIDRepeatingGroup is a repeating group, Tag 454.
type NoSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoSecurityAltIDRepeatingGroup returns an initialized, NoSecurityAltIDRepeatingGroup.
func NewNoSecurityAltIDRepeatingGroup() NoSecurityAltIDRepeatingGroup {
	return NoSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SecurityAltID), quickfix.GroupElement(tag.SecurityAltIDSource)})}
}

// Add create and append a new NoSecurityAltID to this group.
func (m NoSecurityAltIDRepeatingGroup) Add() NoSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoSecurityAltID{g}
}

// Get returns the ith NoSecurityAltID in the NoSecurityAltIDRepeatinGroup.
func (m NoSecurityAltIDRepeatingGroup) Get(i int) NoSecurityAltID {
	return NoSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoLegs is a repeating group element, Tag 555.
type NoLegs struct {
	*quickfix.Group
}

// SetLegSymbol sets LegSymbol, Tag 600.
func (m NoLegs) SetLegSymbol(v string) {
	m.Set(field.NewLegSymbol(v))
}

// SetLegSymbolSfx sets LegSymbolSfx, Tag 601.
func (m NoLegs) SetLegSymbolSfx(v string) {
	m.Set(field.NewLegSymbolSfx(v))
}

// SetLegSecurityID sets LegSecurityID, Tag 602.
func (m NoLegs) SetLegSecurityID(v string) {
	m.Set(field.NewLegSecurityID(v))
}

// SetLegSecurityIDSource sets LegSecurityIDSource, Tag 603.
func (m NoLegs) SetLegSecurityIDSource(v string) {
	m.Set(field.NewLegSecurityIDSource(v))
}

// SetNoLegSecurityAltID sets NoLegSecurityAltID, Tag 604.
func (m NoLegs) SetNoLegSecurityAltID(f NoLegSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetLegProduct sets LegProduct, Tag 607.
func (m NoLegs) SetLegProduct(v int) {
	m.Set(field.NewLegProduct(v))
}

// SetLegCFICode sets LegCFICode, Tag 608.
func (m NoLegs) SetLegCFICode(v string) {
	m.Set(field.NewLegCFICode(v))
}

// SetLegSecurityType sets LegSecurityType, Tag 609.
func (m NoLegs) SetLegSecurityType(v string) {
	m.Set(field.NewLegSecurityType(v))
}

// SetLegSecuritySubType sets LegSecuritySubType, Tag 764.
func (m NoLegs) SetLegSecuritySubType(v string) {
	m.Set(field.NewLegSecuritySubType(v))
}

// SetLegMaturityMonthYear sets LegMaturityMonthYear, Tag 610.
func (m NoLegs) SetLegMaturityMonthYear(v string) {
	m.Set(field.NewLegMaturityMonthYear(v))
}

// SetLegMaturityDate sets LegMaturityDate, Tag 611.
func (m NoLegs) SetLegMaturityDate(v string) {
	m.Set(field.NewLegMaturityDate(v))
}

// SetLegCouponPaymentDate sets LegCouponPaymentDate, Tag 248.
func (m NoLegs) SetLegCouponPaymentDate(v string) {
	m.Set(field.NewLegCouponPaymentDate(v))
}

// SetLegIssueDate sets LegIssueDate, Tag 249.
func (m NoLegs) SetLegIssueDate(v string) {
	m.Set(field.NewLegIssueDate(v))
}

// SetLegRepoCollateralSecurityType sets LegRepoCollateralSecurityType, Tag 250.
func (m NoLegs) SetLegRepoCollateralSecurityType(v int) {
	m.Set(field.NewLegRepoCollateralSecurityType(v))
}

// SetLegRepurchaseTerm sets LegRepurchaseTerm, Tag 251.
func (m NoLegs) SetLegRepurchaseTerm(v int) {
	m.Set(field.NewLegRepurchaseTerm(v))
}

// SetLegRepurchaseRate sets LegRepurchaseRate, Tag 252.
func (m NoLegs) SetLegRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegRepurchaseRate(value, scale))
}

// SetLegFactor sets LegFactor, Tag 253.
func (m NoLegs) SetLegFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegFactor(value, scale))
}

// SetLegCreditRating sets LegCreditRating, Tag 257.
func (m NoLegs) SetLegCreditRating(v string) {
	m.Set(field.NewLegCreditRating(v))
}

// SetLegInstrRegistry sets LegInstrRegistry, Tag 599.
func (m NoLegs) SetLegInstrRegistry(v string) {
	m.Set(field.NewLegInstrRegistry(v))
}

// SetLegCountryOfIssue sets LegCountryOfIssue, Tag 596.
func (m NoLegs) SetLegCountryOfIssue(v string) {
	m.Set(field.NewLegCountryOfIssue(v))
}

// SetLegStateOrProvinceOfIssue sets LegStateOrProvinceOfIssue, Tag 597.
func (m NoLegs) SetLegStateOrProvinceOfIssue(v string) {
	m.Set(field.NewLegStateOrProvinceOfIssue(v))
}

// SetLegLocaleOfIssue sets LegLocaleOfIssue, Tag 598.
func (m NoLegs) SetLegLocaleOfIssue(v string) {
	m.Set(field.NewLegLocaleOfIssue(v))
}

// SetLegRedemptionDate sets LegRedemptionDate, Tag 254.
func (m NoLegs) SetLegRedemptionDate(v string) {
	m.Set(field.NewLegRedemptionDate(v))
}

// SetLegStrikePrice sets LegStrikePrice, Tag 612.
func (m NoLegs) SetLegStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegStrikePrice(value, scale))
}

// SetLegStrikeCurrency sets LegStrikeCurrency, Tag 942.
func (m NoLegs) SetLegStrikeCurrency(v string) {
	m.Set(field.NewLegStrikeCurrency(v))
}

// SetLegOptAttribute sets LegOptAttribute, Tag 613.
func (m NoLegs) SetLegOptAttribute(v string) {
	m.Set(field.NewLegOptAttribute(v))
}

// SetLegContractMultiplier sets LegContractMultiplier, Tag 614.
func (m NoLegs) SetLegContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegContractMultiplier(value, scale))
}

// SetLegCouponRate sets LegCouponRate, Tag 615.
func (m NoLegs) SetLegCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegCouponRate(value, scale))
}

// SetLegSecurityExchange sets LegSecurityExchange, Tag 616.
func (m NoLegs) SetLegSecurityExchange(v string) {
	m.Set(field.NewLegSecurityExchange(v))
}

// SetLegIssuer sets LegIssuer, Tag 617.
func (m NoLegs) SetLegIssuer(v string) {
	m.Set(field.NewLegIssuer(v))
}

// SetEncodedLegIssuerLen sets EncodedLegIssuerLen, Tag 618.
func (m NoLegs) SetEncodedLegIssuerLen(v int) {
	m.Set(field.NewEncodedLegIssuerLen(v))
}

// SetEncodedLegIssuer sets EncodedLegIssuer, Tag 619.
func (m NoLegs) SetEncodedLegIssuer(v string) {
	m.Set(field.NewEncodedLegIssuer(v))
}

// SetLegSecurityDesc sets LegSecurityDesc, Tag 620.
func (m NoLegs) SetLegSecurityDesc(v string) {
	m.Set(field.NewLegSecurityDesc(v))
}

// SetEncodedLegSecurityDescLen sets EncodedLegSecurityDescLen, Tag 621.
func (m NoLegs) SetEncodedLegSecurityDescLen(v int) {
	m.Set(field.NewEncodedLegSecurityDescLen(v))
}

// SetEncodedLegSecurityDesc sets EncodedLegSecurityDesc, Tag 622.
func (m NoLegs) SetEncodedLegSecurityDesc(v string) {
	m.Set(field.NewEncodedLegSecurityDesc(v))
}

// SetLegRatioQty sets LegRatioQty, Tag 623.
func (m NoLegs) SetLegRatioQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegRatioQty(value, scale))
}

// SetLegSide sets LegSide, Tag 624.
func (m NoLegs) SetLegSide(v string) {
	m.Set(field.NewLegSide(v))
}

// SetLegCurrency sets LegCurrency, Tag 556.
func (m NoLegs) SetLegCurrency(v string) {
	m.Set(field.NewLegCurrency(v))
}

// SetLegPool sets LegPool, Tag 740.
func (m NoLegs) SetLegPool(v string) {
	m.Set(field.NewLegPool(v))
}

// SetLegDatedDate sets LegDatedDate, Tag 739.
func (m NoLegs) SetLegDatedDate(v string) {
	m.Set(field.NewLegDatedDate(v))
}

// SetLegContractSettlMonth sets LegContractSettlMonth, Tag 955.
func (m NoLegs) SetLegContractSettlMonth(v string) {
	m.Set(field.NewLegContractSettlMonth(v))
}

// SetLegInterestAccrualDate sets LegInterestAccrualDate, Tag 956.
func (m NoLegs) SetLegInterestAccrualDate(v string) {
	m.Set(field.NewLegInterestAccrualDate(v))
}

// SetLegUnitOfMeasure sets LegUnitOfMeasure, Tag 999.
func (m NoLegs) SetLegUnitOfMeasure(v string) {
	m.Set(field.NewLegUnitOfMeasure(v))
}

// SetLegTimeUnit sets LegTimeUnit, Tag 1001.
func (m NoLegs) SetLegTimeUnit(v string) {
	m.Set(field.NewLegTimeUnit(v))
}

// SetLegOptionRatio sets LegOptionRatio, Tag 1017.
func (m NoLegs) SetLegOptionRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegOptionRatio(value, scale))
}

// SetLegPrice sets LegPrice, Tag 566.
func (m NoLegs) SetLegPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegPrice(value, scale))
}

// SetLegMaturityTime sets LegMaturityTime, Tag 1212.
func (m NoLegs) SetLegMaturityTime(v string) {
	m.Set(field.NewLegMaturityTime(v))
}

// SetLegPutOrCall sets LegPutOrCall, Tag 1358.
func (m NoLegs) SetLegPutOrCall(v int) {
	m.Set(field.NewLegPutOrCall(v))
}

// SetLegExerciseStyle sets LegExerciseStyle, Tag 1420.
func (m NoLegs) SetLegExerciseStyle(v int) {
	m.Set(field.NewLegExerciseStyle(v))
}

// SetLegUnitOfMeasureQty sets LegUnitOfMeasureQty, Tag 1224.
func (m NoLegs) SetLegUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegUnitOfMeasureQty(value, scale))
}

// SetLegPriceUnitOfMeasure sets LegPriceUnitOfMeasure, Tag 1421.
func (m NoLegs) SetLegPriceUnitOfMeasure(v string) {
	m.Set(field.NewLegPriceUnitOfMeasure(v))
}

// SetLegPriceUnitOfMeasureQty sets LegPriceUnitOfMeasureQty, Tag 1422.
func (m NoLegs) SetLegPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegPriceUnitOfMeasureQty(value, scale))
}

// SetLegContractMultiplierUnit sets LegContractMultiplierUnit, Tag 1436.
func (m NoLegs) SetLegContractMultiplierUnit(v int) {
	m.Set(field.NewLegContractMultiplierUnit(v))
}

// SetLegFlowScheduleType sets LegFlowScheduleType, Tag 1440.
func (m NoLegs) SetLegFlowScheduleType(v int) {
	m.Set(field.NewLegFlowScheduleType(v))
}

// GetLegSymbol gets LegSymbol, Tag 600.
func (m NoLegs) GetLegSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.LegSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSymbolSfx gets LegSymbolSfx, Tag 601.
func (m NoLegs) GetLegSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.LegSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityID gets LegSecurityID, Tag 602.
func (m NoLegs) GetLegSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityIDSource gets LegSecurityIDSource, Tag 603.
func (m NoLegs) GetLegSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoLegSecurityAltID gets NoLegSecurityAltID, Tag 604.
func (m NoLegs) GetNoLegSecurityAltID() (NoLegSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetLegProduct gets LegProduct, Tag 607.
func (m NoLegs) GetLegProduct() (v int, err quickfix.MessageRejectError) {
	var f field.LegProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCFICode gets LegCFICode, Tag 608.
func (m NoLegs) GetLegCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.LegCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityType gets LegSecurityType, Tag 609.
func (m NoLegs) GetLegSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecuritySubType gets LegSecuritySubType, Tag 764.
func (m NoLegs) GetLegSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegMaturityMonthYear gets LegMaturityMonthYear, Tag 610.
func (m NoLegs) GetLegMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegMaturityDate gets LegMaturityDate, Tag 611.
func (m NoLegs) GetLegMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCouponPaymentDate gets LegCouponPaymentDate, Tag 248.
func (m NoLegs) GetLegCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegCouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegIssueDate gets LegIssueDate, Tag 249.
func (m NoLegs) GetLegIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRepoCollateralSecurityType gets LegRepoCollateralSecurityType, Tag 250.
func (m NoLegs) GetLegRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.LegRepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRepurchaseTerm gets LegRepurchaseTerm, Tag 251.
func (m NoLegs) GetLegRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.LegRepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRepurchaseRate gets LegRepurchaseRate, Tag 252.
func (m NoLegs) GetLegRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegRepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegFactor gets LegFactor, Tag 253.
func (m NoLegs) GetLegFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegFactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCreditRating gets LegCreditRating, Tag 257.
func (m NoLegs) GetLegCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.LegCreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegInstrRegistry gets LegInstrRegistry, Tag 599.
func (m NoLegs) GetLegInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.LegInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCountryOfIssue gets LegCountryOfIssue, Tag 596.
func (m NoLegs) GetLegCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegStateOrProvinceOfIssue gets LegStateOrProvinceOfIssue, Tag 597.
func (m NoLegs) GetLegStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegLocaleOfIssue gets LegLocaleOfIssue, Tag 598.
func (m NoLegs) GetLegLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRedemptionDate gets LegRedemptionDate, Tag 254.
func (m NoLegs) GetLegRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegStrikePrice gets LegStrikePrice, Tag 612.
func (m NoLegs) GetLegStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegStrikeCurrency gets LegStrikeCurrency, Tag 942.
func (m NoLegs) GetLegStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.LegStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegOptAttribute gets LegOptAttribute, Tag 613.
func (m NoLegs) GetLegOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.LegOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegContractMultiplier gets LegContractMultiplier, Tag 614.
func (m NoLegs) GetLegContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCouponRate gets LegCouponRate, Tag 615.
func (m NoLegs) GetLegCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityExchange gets LegSecurityExchange, Tag 616.
func (m NoLegs) GetLegSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegIssuer gets LegIssuer, Tag 617.
func (m NoLegs) GetLegIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.LegIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedLegIssuerLen gets EncodedLegIssuerLen, Tag 618.
func (m NoLegs) GetEncodedLegIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedLegIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedLegIssuer gets EncodedLegIssuer, Tag 619.
func (m NoLegs) GetEncodedLegIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedLegIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityDesc gets LegSecurityDesc, Tag 620.
func (m NoLegs) GetLegSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedLegSecurityDescLen gets EncodedLegSecurityDescLen, Tag 621.
func (m NoLegs) GetEncodedLegSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedLegSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedLegSecurityDesc gets EncodedLegSecurityDesc, Tag 622.
func (m NoLegs) GetEncodedLegSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedLegSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRatioQty gets LegRatioQty, Tag 623.
func (m NoLegs) GetLegRatioQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegRatioQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSide gets LegSide, Tag 624.
func (m NoLegs) GetLegSide() (v string, err quickfix.MessageRejectError) {
	var f field.LegSideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCurrency gets LegCurrency, Tag 556.
func (m NoLegs) GetLegCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.LegCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegPool gets LegPool, Tag 740.
func (m NoLegs) GetLegPool() (v string, err quickfix.MessageRejectError) {
	var f field.LegPoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegDatedDate gets LegDatedDate, Tag 739.
func (m NoLegs) GetLegDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegDatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegContractSettlMonth gets LegContractSettlMonth, Tag 955.
func (m NoLegs) GetLegContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.LegContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegInterestAccrualDate gets LegInterestAccrualDate, Tag 956.
func (m NoLegs) GetLegInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegInterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegUnitOfMeasure gets LegUnitOfMeasure, Tag 999.
func (m NoLegs) GetLegUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.LegUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegTimeUnit gets LegTimeUnit, Tag 1001.
func (m NoLegs) GetLegTimeUnit() (v string, err quickfix.MessageRejectError) {
	var f field.LegTimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegOptionRatio gets LegOptionRatio, Tag 1017.
func (m NoLegs) GetLegOptionRatio() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegOptionRatioField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegPrice gets LegPrice, Tag 566.
func (m NoLegs) GetLegPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegMaturityTime gets LegMaturityTime, Tag 1212.
func (m NoLegs) GetLegMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegPutOrCall gets LegPutOrCall, Tag 1358.
func (m NoLegs) GetLegPutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.LegPutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegExerciseStyle gets LegExerciseStyle, Tag 1420.
func (m NoLegs) GetLegExerciseStyle() (v int, err quickfix.MessageRejectError) {
	var f field.LegExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegUnitOfMeasureQty gets LegUnitOfMeasureQty, Tag 1224.
func (m NoLegs) GetLegUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegPriceUnitOfMeasure gets LegPriceUnitOfMeasure, Tag 1421.
func (m NoLegs) GetLegPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.LegPriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegPriceUnitOfMeasureQty gets LegPriceUnitOfMeasureQty, Tag 1422.
func (m NoLegs) GetLegPriceUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegPriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegContractMultiplierUnit gets LegContractMultiplierUnit, Tag 1436.
func (m NoLegs) GetLegContractMultiplierUnit() (v int, err quickfix.MessageRejectError) {
	var f field.LegContractMultiplierUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegFlowScheduleType gets LegFlowScheduleType, Tag 1440.
func (m NoLegs) GetLegFlowScheduleType() (v int, err quickfix.MessageRejectError) {
	var f field.LegFlowScheduleTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasLegSymbol returns true if LegSymbol is present, Tag 600.
func (m NoLegs) HasLegSymbol() bool {
	return m.Has(tag.LegSymbol)
}

// HasLegSymbolSfx returns true if LegSymbolSfx is present, Tag 601.
func (m NoLegs) HasLegSymbolSfx() bool {
	return m.Has(tag.LegSymbolSfx)
}

// HasLegSecurityID returns true if LegSecurityID is present, Tag 602.
func (m NoLegs) HasLegSecurityID() bool {
	return m.Has(tag.LegSecurityID)
}

// HasLegSecurityIDSource returns true if LegSecurityIDSource is present, Tag 603.
func (m NoLegs) HasLegSecurityIDSource() bool {
	return m.Has(tag.LegSecurityIDSource)
}

// HasNoLegSecurityAltID returns true if NoLegSecurityAltID is present, Tag 604.
func (m NoLegs) HasNoLegSecurityAltID() bool {
	return m.Has(tag.NoLegSecurityAltID)
}

// HasLegProduct returns true if LegProduct is present, Tag 607.
func (m NoLegs) HasLegProduct() bool {
	return m.Has(tag.LegProduct)
}

// HasLegCFICode returns true if LegCFICode is present, Tag 608.
func (m NoLegs) HasLegCFICode() bool {
	return m.Has(tag.LegCFICode)
}

// HasLegSecurityType returns true if LegSecurityType is present, Tag 609.
func (m NoLegs) HasLegSecurityType() bool {
	return m.Has(tag.LegSecurityType)
}

// HasLegSecuritySubType returns true if LegSecuritySubType is present, Tag 764.
func (m NoLegs) HasLegSecuritySubType() bool {
	return m.Has(tag.LegSecuritySubType)
}

// HasLegMaturityMonthYear returns true if LegMaturityMonthYear is present, Tag 610.
func (m NoLegs) HasLegMaturityMonthYear() bool {
	return m.Has(tag.LegMaturityMonthYear)
}

// HasLegMaturityDate returns true if LegMaturityDate is present, Tag 611.
func (m NoLegs) HasLegMaturityDate() bool {
	return m.Has(tag.LegMaturityDate)
}

// HasLegCouponPaymentDate returns true if LegCouponPaymentDate is present, Tag 248.
func (m NoLegs) HasLegCouponPaymentDate() bool {
	return m.Has(tag.LegCouponPaymentDate)
}

// HasLegIssueDate returns true if LegIssueDate is present, Tag 249.
func (m NoLegs) HasLegIssueDate() bool {
	return m.Has(tag.LegIssueDate)
}

// HasLegRepoCollateralSecurityType returns true if LegRepoCollateralSecurityType is present, Tag 250.
func (m NoLegs) HasLegRepoCollateralSecurityType() bool {
	return m.Has(tag.LegRepoCollateralSecurityType)
}

// HasLegRepurchaseTerm returns true if LegRepurchaseTerm is present, Tag 251.
func (m NoLegs) HasLegRepurchaseTerm() bool {
	return m.Has(tag.LegRepurchaseTerm)
}

// HasLegRepurchaseRate returns true if LegRepurchaseRate is present, Tag 252.
func (m NoLegs) HasLegRepurchaseRate() bool {
	return m.Has(tag.LegRepurchaseRate)
}

// HasLegFactor returns true if LegFactor is present, Tag 253.
func (m NoLegs) HasLegFactor() bool {
	return m.Has(tag.LegFactor)
}

// HasLegCreditRating returns true if LegCreditRating is present, Tag 257.
func (m NoLegs) HasLegCreditRating() bool {
	return m.Has(tag.LegCreditRating)
}

// HasLegInstrRegistry returns true if LegInstrRegistry is present, Tag 599.
func (m NoLegs) HasLegInstrRegistry() bool {
	return m.Has(tag.LegInstrRegistry)
}

// HasLegCountryOfIssue returns true if LegCountryOfIssue is present, Tag 596.
func (m NoLegs) HasLegCountryOfIssue() bool {
	return m.Has(tag.LegCountryOfIssue)
}

// HasLegStateOrProvinceOfIssue returns true if LegStateOrProvinceOfIssue is present, Tag 597.
func (m NoLegs) HasLegStateOrProvinceOfIssue() bool {
	return m.Has(tag.LegStateOrProvinceOfIssue)
}

// HasLegLocaleOfIssue returns true if LegLocaleOfIssue is present, Tag 598.
func (m NoLegs) HasLegLocaleOfIssue() bool {
	return m.Has(tag.LegLocaleOfIssue)
}

// HasLegRedemptionDate returns true if LegRedemptionDate is present, Tag 254.
func (m NoLegs) HasLegRedemptionDate() bool {
	return m.Has(tag.LegRedemptionDate)
}

// HasLegStrikePrice returns true if LegStrikePrice is present, Tag 612.
func (m NoLegs) HasLegStrikePrice() bool {
	return m.Has(tag.LegStrikePrice)
}

// HasLegStrikeCurrency returns true if LegStrikeCurrency is present, Tag 942.
func (m NoLegs) HasLegStrikeCurrency() bool {
	return m.Has(tag.LegStrikeCurrency)
}

// HasLegOptAttribute returns true if LegOptAttribute is present, Tag 613.
func (m NoLegs) HasLegOptAttribute() bool {
	return m.Has(tag.LegOptAttribute)
}

// HasLegContractMultiplier returns true if LegContractMultiplier is present, Tag 614.
func (m NoLegs) HasLegContractMultiplier() bool {
	return m.Has(tag.LegContractMultiplier)
}

// HasLegCouponRate returns true if LegCouponRate is present, Tag 615.
func (m NoLegs) HasLegCouponRate() bool {
	return m.Has(tag.LegCouponRate)
}

// HasLegSecurityExchange returns true if LegSecurityExchange is present, Tag 616.
func (m NoLegs) HasLegSecurityExchange() bool {
	return m.Has(tag.LegSecurityExchange)
}

// HasLegIssuer returns true if LegIssuer is present, Tag 617.
func (m NoLegs) HasLegIssuer() bool {
	return m.Has(tag.LegIssuer)
}

// HasEncodedLegIssuerLen returns true if EncodedLegIssuerLen is present, Tag 618.
func (m NoLegs) HasEncodedLegIssuerLen() bool {
	return m.Has(tag.EncodedLegIssuerLen)
}

// HasEncodedLegIssuer returns true if EncodedLegIssuer is present, Tag 619.
func (m NoLegs) HasEncodedLegIssuer() bool {
	return m.Has(tag.EncodedLegIssuer)
}

// HasLegSecurityDesc returns true if LegSecurityDesc is present, Tag 620.
func (m NoLegs) HasLegSecurityDesc() bool {
	return m.Has(tag.LegSecurityDesc)
}

// HasEncodedLegSecurityDescLen returns true if EncodedLegSecurityDescLen is present, Tag 621.
func (m NoLegs) HasEncodedLegSecurityDescLen() bool {
	return m.Has(tag.EncodedLegSecurityDescLen)
}

// HasEncodedLegSecurityDesc returns true if EncodedLegSecurityDesc is present, Tag 622.
func (m NoLegs) HasEncodedLegSecurityDesc() bool {
	return m.Has(tag.EncodedLegSecurityDesc)
}

// HasLegRatioQty returns true if LegRatioQty is present, Tag 623.
func (m NoLegs) HasLegRatioQty() bool {
	return m.Has(tag.LegRatioQty)
}

// HasLegSide returns true if LegSide is present, Tag 624.
func (m NoLegs) HasLegSide() bool {
	return m.Has(tag.LegSide)
}

// HasLegCurrency returns true if LegCurrency is present, Tag 556.
func (m NoLegs) HasLegCurrency() bool {
	return m.Has(tag.LegCurrency)
}

// HasLegPool returns true if LegPool is present, Tag 740.
func (m NoLegs) HasLegPool() bool {
	return m.Has(tag.LegPool)
}

// HasLegDatedDate returns true if LegDatedDate is present, Tag 739.
func (m NoLegs) HasLegDatedDate() bool {
	return m.Has(tag.LegDatedDate)
}

// HasLegContractSettlMonth returns true if LegContractSettlMonth is present, Tag 955.
func (m NoLegs) HasLegContractSettlMonth() bool {
	return m.Has(tag.LegContractSettlMonth)
}

// HasLegInterestAccrualDate returns true if LegInterestAccrualDate is present, Tag 956.
func (m NoLegs) HasLegInterestAccrualDate() bool {
	return m.Has(tag.LegInterestAccrualDate)
}

// HasLegUnitOfMeasure returns true if LegUnitOfMeasure is present, Tag 999.
func (m NoLegs) HasLegUnitOfMeasure() bool {
	return m.Has(tag.LegUnitOfMeasure)
}

// HasLegTimeUnit returns true if LegTimeUnit is present, Tag 1001.
func (m NoLegs) HasLegTimeUnit() bool {
	return m.Has(tag.LegTimeUnit)
}

// HasLegOptionRatio returns true if LegOptionRatio is present, Tag 1017.
func (m NoLegs) HasLegOptionRatio() bool {
	return m.Has(tag.LegOptionRatio)
}

// HasLegPrice returns true if LegPrice is present, Tag 566.
func (m NoLegs) HasLegPrice() bool {
	return m.Has(tag.LegPrice)
}

// HasLegMaturityTime returns true if LegMaturityTime is present, Tag 1212.
func (m NoLegs) HasLegMaturityTime() bool {
	return m.Has(tag.LegMaturityTime)
}

// HasLegPutOrCall returns true if LegPutOrCall is present, Tag 1358.
func (m NoLegs) HasLegPutOrCall() bool {
	return m.Has(tag.LegPutOrCall)
}

// HasLegExerciseStyle returns true if LegExerciseStyle is present, Tag 1420.
func (m NoLegs) HasLegExerciseStyle() bool {
	return m.Has(tag.LegExerciseStyle)
}

// HasLegUnitOfMeasureQty returns true if LegUnitOfMeasureQty is present, Tag 1224.
func (m NoLegs) HasLegUnitOfMeasureQty() bool {
	return m.Has(tag.LegUnitOfMeasureQty)
}

// HasLegPriceUnitOfMeasure returns true if LegPriceUnitOfMeasure is present, Tag 1421.
func (m NoLegs) HasLegPriceUnitOfMeasure() bool {
	return m.Has(tag.LegPriceUnitOfMeasure)
}

// HasLegPriceUnitOfMeasureQty returns true if LegPriceUnitOfMeasureQty is present, Tag 1422.
func (m NoLegs) HasLegPriceUnitOfMeasureQty() bool {
	return m.Has(tag.LegPriceUnitOfMeasureQty)
}

// HasLegContractMultiplierUnit returns true if LegContractMultiplierUnit is present, Tag 1436.
func (m NoLegs) HasLegContractMultiplierUnit() bool {
	return m.Has(tag.LegContractMultiplierUnit)
}

// HasLegFlowScheduleType returns true if LegFlowScheduleType is present, Tag 1440.
func (m NoLegs) HasLegFlowScheduleType() bool {
	return m.Has(tag.LegFlowScheduleType)
}

// NoLegSecurityAltID is a repeating group element, Tag 604.
type NoLegSecurityAltID struct {
	*quickfix.Group
}

// SetLegSecurityAltID sets LegSecurityAltID, Tag 605.
func (m NoLegSecurityAltID) SetLegSecurityAltID(v string) {
	m.Set(field.NewLegSecurityAltID(v))
}

// SetLegSecurityAltIDSource sets LegSecurityAltIDSource, Tag 606.
func (m NoLegSecurityAltID) SetLegSecurityAltIDSource(v string) {
	m.Set(field.NewLegSecurityAltIDSource(v))
}

// GetLegSecurityAltID gets LegSecurityAltID, Tag 605.
func (m NoLegSecurityAltID) GetLegSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityAltIDSource gets LegSecurityAltIDSource, Tag 606.
func (m NoLegSecurityAltID) GetLegSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasLegSecurityAltID returns true if LegSecurityAltID is present, Tag 605.
func (m NoLegSecurityAltID) HasLegSecurityAltID() bool {
	return m.Has(tag.LegSecurityAltID)
}

// HasLegSecurityAltIDSource returns true if LegSecurityAltIDSource is present, Tag 606.
func (m NoLegSecurityAltID) HasLegSecurityAltIDSource() bool {
	return m.Has(tag.LegSecurityAltIDSource)
}

// NoLegSecurityAltIDRepeatingGroup is a repeating group, Tag 604.
type NoLegSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoLegSecurityAltIDRepeatingGroup returns an initialized, NoLegSecurityAltIDRepeatingGroup.
func NewNoLegSecurityAltIDRepeatingGroup() NoLegSecurityAltIDRepeatingGroup {
	return NoLegSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSecurityAltID), quickfix.GroupElement(tag.LegSecurityAltIDSource)})}
}

// Add create and append a new NoLegSecurityAltID to this group.
func (m NoLegSecurityAltIDRepeatingGroup) Add() NoLegSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoLegSecurityAltID{g}
}

// Get returns the ith NoLegSecurityAltID in the NoLegSecurityAltIDRepeatinGroup.
func (m NoLegSecurityAltIDRepeatingGroup) Get(i int) NoLegSecurityAltID {
	return NoLegSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoLegsRepeatingGroup is a repeating group, Tag 555.
type NoLegsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoLegsRepeatingGroup returns an initialized, NoLegsRepeatingGroup.
func NewNoLegsRepeatingGroup() NoLegsRepeatingGroup {
	return NoLegsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSymbol), quickfix.GroupElement(tag.LegSymbolSfx), quickfix.GroupElement(tag.LegSecurityID), quickfix.GroupElement(tag.LegSecurityIDSource), NewNoLegSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.LegProduct), quickfix.GroupElement(tag.LegCFICode), quickfix.GroupElement(tag.LegSecurityType), quickfix.GroupElement(tag.LegSecuritySubType), quickfix.GroupElement(tag.LegMaturityMonthYear), quickfix.GroupElement(tag.LegMaturityDate), quickfix.GroupElement(tag.LegCouponPaymentDate), quickfix.GroupElement(tag.LegIssueDate), quickfix.GroupElement(tag.LegRepoCollateralSecurityType), quickfix.GroupElement(tag.LegRepurchaseTerm), quickfix.GroupElement(tag.LegRepurchaseRate), quickfix.GroupElement(tag.LegFactor), quickfix.GroupElement(tag.LegCreditRating), quickfix.GroupElement(tag.LegInstrRegistry), quickfix.GroupElement(tag.LegCountryOfIssue), quickfix.GroupElement(tag.LegStateOrProvinceOfIssue), quickfix.GroupElement(tag.LegLocaleOfIssue), quickfix.GroupElement(tag.LegRedemptionDate), quickfix.GroupElement(tag.LegStrikePrice), quickfix.GroupElement(tag.LegStrikeCurrency), quickfix.GroupElement(tag.LegOptAttribute), quickfix.GroupElement(tag.LegContractMultiplier), quickfix.GroupElement(tag.LegCouponRate), quickfix.GroupElement(tag.LegSecurityExchange), quickfix.GroupElement(tag.LegIssuer), quickfix.GroupElement(tag.EncodedLegIssuerLen), quickfix.GroupElement(tag.EncodedLegIssuer), quickfix.GroupElement(tag.LegSecurityDesc), quickfix.GroupElement(tag.EncodedLegSecurityDescLen), quickfix.GroupElement(tag.EncodedLegSecurityDesc), quickfix.GroupElement(tag.LegRatioQty), quickfix.GroupElement(tag.LegSide), quickfix.GroupElement(tag.LegCurrency), quickfix.GroupElement(tag.LegPool), quickfix.GroupElement(tag.LegDatedDate), quickfix.GroupElement(tag.LegContractSettlMonth), quickfix.GroupElement(tag.LegInterestAccrualDate), quickfix.GroupElement(tag.LegUnitOfMeasure), quickfix.GroupElement(tag.LegTimeUnit), quickfix.GroupElement(tag.LegOptionRatio), quickfix.GroupElement(tag.LegPrice), quickfix.GroupElement(tag.LegMaturityTime), quickfix.GroupElement(tag.LegPutOrCall), quickfix.GroupElement(tag.LegExerciseStyle), quickfix.GroupElement(tag.LegUnitOfMeasureQty), quickfix.GroupElement(tag.LegPriceUnitOfMeasure), quickfix.GroupElement(tag.LegPriceUnitOfMeasureQty), quickfix.GroupElement(tag.LegContractMultiplierUnit), quickfix.GroupElement(tag.LegFlowScheduleType)})}
}

// Add create and append a new NoLegs to this group.
func (m NoLegsRepeatingGroup) Add() NoLegs {
	g := m.RepeatingGroup.Add()
	return NoLegs{g}
}

// Get returns the ith NoLegs in the NoLegsRepeatinGroup.
func (m NoLegsRepeatingGroup) Get(i int) NoLegs {
	return NoLegs{m.RepeatingGroup.Get(i)}
}

// NoUnderlyings is a repeating group element, Tag 711.
type NoUnderlyings struct {
	*quickfix.Group
}

// SetUnderlyingSymbol sets UnderlyingSymbol, Tag 311.
func (m NoUnderlyings) SetUnderlyingSymbol(v string) {
	m.Set(field.NewUnderlyingSymbol(v))
}

// SetUnderlyingSymbolSfx sets UnderlyingSymbolSfx, Tag 312.
func (m NoUnderlyings) SetUnderlyingSymbolSfx(v string) {
	m.Set(field.NewUnderlyingSymbolSfx(v))
}

// SetUnderlyingSecurityID sets UnderlyingSecurityID, Tag 309.
func (m NoUnderlyings) SetUnderlyingSecurityID(v string) {
	m.Set(field.NewUnderlyingSecurityID(v))
}

// SetUnderlyingSecurityIDSource sets UnderlyingSecurityIDSource, Tag 305.
func (m NoUnderlyings) SetUnderlyingSecurityIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityIDSource(v))
}

// SetNoUnderlyingSecurityAltID sets NoUnderlyingSecurityAltID, Tag 457.
func (m NoUnderlyings) SetNoUnderlyingSecurityAltID(f NoUnderlyingSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetUnderlyingProduct sets UnderlyingProduct, Tag 462.
func (m NoUnderlyings) SetUnderlyingProduct(v int) {
	m.Set(field.NewUnderlyingProduct(v))
}

// SetUnderlyingCFICode sets UnderlyingCFICode, Tag 463.
func (m NoUnderlyings) SetUnderlyingCFICode(v string) {
	m.Set(field.NewUnderlyingCFICode(v))
}

// SetUnderlyingSecurityType sets UnderlyingSecurityType, Tag 310.
func (m NoUnderlyings) SetUnderlyingSecurityType(v string) {
	m.Set(field.NewUnderlyingSecurityType(v))
}

// SetUnderlyingSecuritySubType sets UnderlyingSecuritySubType, Tag 763.
func (m NoUnderlyings) SetUnderlyingSecuritySubType(v string) {
	m.Set(field.NewUnderlyingSecuritySubType(v))
}

// SetUnderlyingMaturityMonthYear sets UnderlyingMaturityMonthYear, Tag 313.
func (m NoUnderlyings) SetUnderlyingMaturityMonthYear(v string) {
	m.Set(field.NewUnderlyingMaturityMonthYear(v))
}

// SetUnderlyingMaturityDate sets UnderlyingMaturityDate, Tag 542.
func (m NoUnderlyings) SetUnderlyingMaturityDate(v string) {
	m.Set(field.NewUnderlyingMaturityDate(v))
}

// SetUnderlyingCouponPaymentDate sets UnderlyingCouponPaymentDate, Tag 241.
func (m NoUnderlyings) SetUnderlyingCouponPaymentDate(v string) {
	m.Set(field.NewUnderlyingCouponPaymentDate(v))
}

// SetUnderlyingIssueDate sets UnderlyingIssueDate, Tag 242.
func (m NoUnderlyings) SetUnderlyingIssueDate(v string) {
	m.Set(field.NewUnderlyingIssueDate(v))
}

// SetUnderlyingRepoCollateralSecurityType sets UnderlyingRepoCollateralSecurityType, Tag 243.
func (m NoUnderlyings) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.Set(field.NewUnderlyingRepoCollateralSecurityType(v))
}

// SetUnderlyingRepurchaseTerm sets UnderlyingRepurchaseTerm, Tag 244.
func (m NoUnderlyings) SetUnderlyingRepurchaseTerm(v int) {
	m.Set(field.NewUnderlyingRepurchaseTerm(v))
}

// SetUnderlyingRepurchaseRate sets UnderlyingRepurchaseRate, Tag 245.
func (m NoUnderlyings) SetUnderlyingRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingRepurchaseRate(value, scale))
}

// SetUnderlyingFactor sets UnderlyingFactor, Tag 246.
func (m NoUnderlyings) SetUnderlyingFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFactor(value, scale))
}

// SetUnderlyingCreditRating sets UnderlyingCreditRating, Tag 256.
func (m NoUnderlyings) SetUnderlyingCreditRating(v string) {
	m.Set(field.NewUnderlyingCreditRating(v))
}

// SetUnderlyingInstrRegistry sets UnderlyingInstrRegistry, Tag 595.
func (m NoUnderlyings) SetUnderlyingInstrRegistry(v string) {
	m.Set(field.NewUnderlyingInstrRegistry(v))
}

// SetUnderlyingCountryOfIssue sets UnderlyingCountryOfIssue, Tag 592.
func (m NoUnderlyings) SetUnderlyingCountryOfIssue(v string) {
	m.Set(field.NewUnderlyingCountryOfIssue(v))
}

// SetUnderlyingStateOrProvinceOfIssue sets UnderlyingStateOrProvinceOfIssue, Tag 593.
func (m NoUnderlyings) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.Set(field.NewUnderlyingStateOrProvinceOfIssue(v))
}

// SetUnderlyingLocaleOfIssue sets UnderlyingLocaleOfIssue, Tag 594.
func (m NoUnderlyings) SetUnderlyingLocaleOfIssue(v string) {
	m.Set(field.NewUnderlyingLocaleOfIssue(v))
}

// SetUnderlyingRedemptionDate sets UnderlyingRedemptionDate, Tag 247.
func (m NoUnderlyings) SetUnderlyingRedemptionDate(v string) {
	m.Set(field.NewUnderlyingRedemptionDate(v))
}

// SetUnderlyingStrikePrice sets UnderlyingStrikePrice, Tag 316.
func (m NoUnderlyings) SetUnderlyingStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStrikePrice(value, scale))
}

// SetUnderlyingStrikeCurrency sets UnderlyingStrikeCurrency, Tag 941.
func (m NoUnderlyings) SetUnderlyingStrikeCurrency(v string) {
	m.Set(field.NewUnderlyingStrikeCurrency(v))
}

// SetUnderlyingOptAttribute sets UnderlyingOptAttribute, Tag 317.
func (m NoUnderlyings) SetUnderlyingOptAttribute(v string) {
	m.Set(field.NewUnderlyingOptAttribute(v))
}

// SetUnderlyingContractMultiplier sets UnderlyingContractMultiplier, Tag 436.
func (m NoUnderlyings) SetUnderlyingContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingContractMultiplier(value, scale))
}

// SetUnderlyingCouponRate sets UnderlyingCouponRate, Tag 435.
func (m NoUnderlyings) SetUnderlyingCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCouponRate(value, scale))
}

// SetUnderlyingSecurityExchange sets UnderlyingSecurityExchange, Tag 308.
func (m NoUnderlyings) SetUnderlyingSecurityExchange(v string) {
	m.Set(field.NewUnderlyingSecurityExchange(v))
}

// SetUnderlyingIssuer sets UnderlyingIssuer, Tag 306.
func (m NoUnderlyings) SetUnderlyingIssuer(v string) {
	m.Set(field.NewUnderlyingIssuer(v))
}

// SetEncodedUnderlyingIssuerLen sets EncodedUnderlyingIssuerLen, Tag 362.
func (m NoUnderlyings) SetEncodedUnderlyingIssuerLen(v int) {
	m.Set(field.NewEncodedUnderlyingIssuerLen(v))
}

// SetEncodedUnderlyingIssuer sets EncodedUnderlyingIssuer, Tag 363.
func (m NoUnderlyings) SetEncodedUnderlyingIssuer(v string) {
	m.Set(field.NewEncodedUnderlyingIssuer(v))
}

// SetUnderlyingSecurityDesc sets UnderlyingSecurityDesc, Tag 307.
func (m NoUnderlyings) SetUnderlyingSecurityDesc(v string) {
	m.Set(field.NewUnderlyingSecurityDesc(v))
}

// SetEncodedUnderlyingSecurityDescLen sets EncodedUnderlyingSecurityDescLen, Tag 364.
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.Set(field.NewEncodedUnderlyingSecurityDescLen(v))
}

// SetEncodedUnderlyingSecurityDesc sets EncodedUnderlyingSecurityDesc, Tag 365.
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDesc(v string) {
	m.Set(field.NewEncodedUnderlyingSecurityDesc(v))
}

// SetUnderlyingCPProgram sets UnderlyingCPProgram, Tag 877.
func (m NoUnderlyings) SetUnderlyingCPProgram(v string) {
	m.Set(field.NewUnderlyingCPProgram(v))
}

// SetUnderlyingCPRegType sets UnderlyingCPRegType, Tag 878.
func (m NoUnderlyings) SetUnderlyingCPRegType(v string) {
	m.Set(field.NewUnderlyingCPRegType(v))
}

// SetUnderlyingCurrency sets UnderlyingCurrency, Tag 318.
func (m NoUnderlyings) SetUnderlyingCurrency(v string) {
	m.Set(field.NewUnderlyingCurrency(v))
}

// SetUnderlyingQty sets UnderlyingQty, Tag 879.
func (m NoUnderlyings) SetUnderlyingQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingQty(value, scale))
}

// SetUnderlyingPx sets UnderlyingPx, Tag 810.
func (m NoUnderlyings) SetUnderlyingPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPx(value, scale))
}

// SetUnderlyingDirtyPrice sets UnderlyingDirtyPrice, Tag 882.
func (m NoUnderlyings) SetUnderlyingDirtyPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDirtyPrice(value, scale))
}

// SetUnderlyingEndPrice sets UnderlyingEndPrice, Tag 883.
func (m NoUnderlyings) SetUnderlyingEndPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndPrice(value, scale))
}

// SetUnderlyingStartValue sets UnderlyingStartValue, Tag 884.
func (m NoUnderlyings) SetUnderlyingStartValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStartValue(value, scale))
}

// SetUnderlyingCurrentValue sets UnderlyingCurrentValue, Tag 885.
func (m NoUnderlyings) SetUnderlyingCurrentValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCurrentValue(value, scale))
}

// SetUnderlyingEndValue sets UnderlyingEndValue, Tag 886.
func (m NoUnderlyings) SetUnderlyingEndValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndValue(value, scale))
}

// SetNoUnderlyingStips sets NoUnderlyingStips, Tag 887.
func (m NoUnderlyings) SetNoUnderlyingStips(f NoUnderlyingStipsRepeatingGroup) {
	m.SetGroup(f)
}

// SetUnderlyingAllocationPercent sets UnderlyingAllocationPercent, Tag 972.
func (m NoUnderlyings) SetUnderlyingAllocationPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAllocationPercent(value, scale))
}

// SetUnderlyingSettlementType sets UnderlyingSettlementType, Tag 975.
func (m NoUnderlyings) SetUnderlyingSettlementType(v enum.UnderlyingSettlementType) {
	m.Set(field.NewUnderlyingSettlementType(v))
}

// SetUnderlyingCashAmount sets UnderlyingCashAmount, Tag 973.
func (m NoUnderlyings) SetUnderlyingCashAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCashAmount(value, scale))
}

// SetUnderlyingCashType sets UnderlyingCashType, Tag 974.
func (m NoUnderlyings) SetUnderlyingCashType(v enum.UnderlyingCashType) {
	m.Set(field.NewUnderlyingCashType(v))
}

// SetUnderlyingUnitOfMeasure sets UnderlyingUnitOfMeasure, Tag 998.
func (m NoUnderlyings) SetUnderlyingUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingUnitOfMeasure(v))
}

// SetUnderlyingTimeUnit sets UnderlyingTimeUnit, Tag 1000.
func (m NoUnderlyings) SetUnderlyingTimeUnit(v string) {
	m.Set(field.NewUnderlyingTimeUnit(v))
}

// SetUnderlyingCapValue sets UnderlyingCapValue, Tag 1038.
func (m NoUnderlyings) SetUnderlyingCapValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCapValue(value, scale))
}

// SetNoUndlyInstrumentParties sets NoUndlyInstrumentParties, Tag 1058.
func (m NoUnderlyings) SetNoUndlyInstrumentParties(f NoUndlyInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

// SetUnderlyingSettlMethod sets UnderlyingSettlMethod, Tag 1039.
func (m NoUnderlyings) SetUnderlyingSettlMethod(v string) {
	m.Set(field.NewUnderlyingSettlMethod(v))
}

// SetUnderlyingAdjustedQuantity sets UnderlyingAdjustedQuantity, Tag 1044.
func (m NoUnderlyings) SetUnderlyingAdjustedQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAdjustedQuantity(value, scale))
}

// SetUnderlyingFXRate sets UnderlyingFXRate, Tag 1045.
func (m NoUnderlyings) SetUnderlyingFXRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFXRate(value, scale))
}

// SetUnderlyingFXRateCalc sets UnderlyingFXRateCalc, Tag 1046.
func (m NoUnderlyings) SetUnderlyingFXRateCalc(v enum.UnderlyingFXRateCalc) {
	m.Set(field.NewUnderlyingFXRateCalc(v))
}

// SetUnderlyingMaturityTime sets UnderlyingMaturityTime, Tag 1213.
func (m NoUnderlyings) SetUnderlyingMaturityTime(v string) {
	m.Set(field.NewUnderlyingMaturityTime(v))
}

// SetUnderlyingPutOrCall sets UnderlyingPutOrCall, Tag 315.
func (m NoUnderlyings) SetUnderlyingPutOrCall(v int) {
	m.Set(field.NewUnderlyingPutOrCall(v))
}

// SetUnderlyingExerciseStyle sets UnderlyingExerciseStyle, Tag 1419.
func (m NoUnderlyings) SetUnderlyingExerciseStyle(v int) {
	m.Set(field.NewUnderlyingExerciseStyle(v))
}

// SetUnderlyingUnitOfMeasureQty sets UnderlyingUnitOfMeasureQty, Tag 1423.
func (m NoUnderlyings) SetUnderlyingUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingUnitOfMeasureQty(value, scale))
}

// SetUnderlyingPriceUnitOfMeasure sets UnderlyingPriceUnitOfMeasure, Tag 1424.
func (m NoUnderlyings) SetUnderlyingPriceUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasure(v))
}

// SetUnderlyingPriceUnitOfMeasureQty sets UnderlyingPriceUnitOfMeasureQty, Tag 1425.
func (m NoUnderlyings) SetUnderlyingPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasureQty(value, scale))
}

// SetUnderlyingContractMultiplierUnit sets UnderlyingContractMultiplierUnit, Tag 1437.
func (m NoUnderlyings) SetUnderlyingContractMultiplierUnit(v int) {
	m.Set(field.NewUnderlyingContractMultiplierUnit(v))
}

// SetUnderlyingFlowScheduleType sets UnderlyingFlowScheduleType, Tag 1441.
func (m NoUnderlyings) SetUnderlyingFlowScheduleType(v int) {
	m.Set(field.NewUnderlyingFlowScheduleType(v))
}

// SetUnderlyingRestructuringType sets UnderlyingRestructuringType, Tag 1453.
func (m NoUnderlyings) SetUnderlyingRestructuringType(v string) {
	m.Set(field.NewUnderlyingRestructuringType(v))
}

// SetUnderlyingSeniority sets UnderlyingSeniority, Tag 1454.
func (m NoUnderlyings) SetUnderlyingSeniority(v string) {
	m.Set(field.NewUnderlyingSeniority(v))
}

// SetUnderlyingNotionalPercentageOutstanding sets UnderlyingNotionalPercentageOutstanding, Tag 1455.
func (m NoUnderlyings) SetUnderlyingNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingNotionalPercentageOutstanding(value, scale))
}

// SetUnderlyingOriginalNotionalPercentageOutstanding sets UnderlyingOriginalNotionalPercentageOutstanding, Tag 1456.
func (m NoUnderlyings) SetUnderlyingOriginalNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingOriginalNotionalPercentageOutstanding(value, scale))
}

// SetUnderlyingAttachmentPoint sets UnderlyingAttachmentPoint, Tag 1459.
func (m NoUnderlyings) SetUnderlyingAttachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAttachmentPoint(value, scale))
}

// SetUnderlyingDetachmentPoint sets UnderlyingDetachmentPoint, Tag 1460.
func (m NoUnderlyings) SetUnderlyingDetachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDetachmentPoint(value, scale))
}

// SetCollAction sets CollAction, Tag 944.
func (m NoUnderlyings) SetCollAction(v enum.CollAction) {
	m.Set(field.NewCollAction(v))
}

// GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311.
func (m NoUnderlyings) GetUnderlyingSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312.
func (m NoUnderlyings) GetUnderlyingSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309.
func (m NoUnderlyings) GetUnderlyingSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityIDSource gets UnderlyingSecurityIDSource, Tag 305.
func (m NoUnderlyings) GetUnderlyingSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUnderlyingSecurityAltID gets NoUnderlyingSecurityAltID, Tag 457.
func (m NoUnderlyings) GetNoUnderlyingSecurityAltID() (NoUnderlyingSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetUnderlyingProduct gets UnderlyingProduct, Tag 462.
func (m NoUnderlyings) GetUnderlyingProduct() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCFICode gets UnderlyingCFICode, Tag 463.
func (m NoUnderlyings) GetUnderlyingCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310.
func (m NoUnderlyings) GetUnderlyingSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecuritySubType gets UnderlyingSecuritySubType, Tag 763.
func (m NoUnderlyings) GetUnderlyingSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313.
func (m NoUnderlyings) GetUnderlyingMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingMaturityDate gets UnderlyingMaturityDate, Tag 542.
func (m NoUnderlyings) GetUnderlyingMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCouponPaymentDate gets UnderlyingCouponPaymentDate, Tag 241.
func (m NoUnderlyings) GetUnderlyingCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingIssueDate gets UnderlyingIssueDate, Tag 242.
func (m NoUnderlyings) GetUnderlyingIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepoCollateralSecurityType gets UnderlyingRepoCollateralSecurityType, Tag 243.
func (m NoUnderlyings) GetUnderlyingRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepurchaseTerm gets UnderlyingRepurchaseTerm, Tag 244.
func (m NoUnderlyings) GetUnderlyingRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepurchaseRate gets UnderlyingRepurchaseRate, Tag 245.
func (m NoUnderlyings) GetUnderlyingRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingFactor gets UnderlyingFactor, Tag 246.
func (m NoUnderlyings) GetUnderlyingFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingFactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCreditRating gets UnderlyingCreditRating, Tag 256.
func (m NoUnderlyings) GetUnderlyingCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingInstrRegistry gets UnderlyingInstrRegistry, Tag 595.
func (m NoUnderlyings) GetUnderlyingInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCountryOfIssue gets UnderlyingCountryOfIssue, Tag 592.
func (m NoUnderlyings) GetUnderlyingCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStateOrProvinceOfIssue gets UnderlyingStateOrProvinceOfIssue, Tag 593.
func (m NoUnderlyings) GetUnderlyingStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingLocaleOfIssue gets UnderlyingLocaleOfIssue, Tag 594.
func (m NoUnderlyings) GetUnderlyingLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRedemptionDate gets UnderlyingRedemptionDate, Tag 247.
func (m NoUnderlyings) GetUnderlyingRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316.
func (m NoUnderlyings) GetUnderlyingStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStrikeCurrency gets UnderlyingStrikeCurrency, Tag 941.
func (m NoUnderlyings) GetUnderlyingStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317.
func (m NoUnderlyings) GetUnderlyingOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436.
func (m NoUnderlyings) GetUnderlyingContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435.
func (m NoUnderlyings) GetUnderlyingCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308.
func (m NoUnderlyings) GetUnderlyingSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306.
func (m NoUnderlyings) GetUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362.
func (m NoUnderlyings) GetEncodedUnderlyingIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363.
func (m NoUnderlyings) GetEncodedUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307.
func (m NoUnderlyings) GetUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364.
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365.
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCPProgram gets UnderlyingCPProgram, Tag 877.
func (m NoUnderlyings) GetUnderlyingCPProgram() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCPRegType gets UnderlyingCPRegType, Tag 878.
func (m NoUnderlyings) GetUnderlyingCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCurrency gets UnderlyingCurrency, Tag 318.
func (m NoUnderlyings) GetUnderlyingCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingQty gets UnderlyingQty, Tag 879.
func (m NoUnderlyings) GetUnderlyingQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPx gets UnderlyingPx, Tag 810.
func (m NoUnderlyings) GetUnderlyingPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingDirtyPrice gets UnderlyingDirtyPrice, Tag 882.
func (m NoUnderlyings) GetUnderlyingDirtyPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingDirtyPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingEndPrice gets UnderlyingEndPrice, Tag 883.
func (m NoUnderlyings) GetUnderlyingEndPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStartValue gets UnderlyingStartValue, Tag 884.
func (m NoUnderlyings) GetUnderlyingStartValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStartValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCurrentValue gets UnderlyingCurrentValue, Tag 885.
func (m NoUnderlyings) GetUnderlyingCurrentValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrentValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingEndValue gets UnderlyingEndValue, Tag 886.
func (m NoUnderlyings) GetUnderlyingEndValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUnderlyingStips gets NoUnderlyingStips, Tag 887.
func (m NoUnderlyings) GetNoUnderlyingStips() (NoUnderlyingStipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingStipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetUnderlyingAllocationPercent gets UnderlyingAllocationPercent, Tag 972.
func (m NoUnderlyings) GetUnderlyingAllocationPercent() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingAllocationPercentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSettlementType gets UnderlyingSettlementType, Tag 975.
func (m NoUnderlyings) GetUnderlyingSettlementType() (v enum.UnderlyingSettlementType, err quickfix.MessageRejectError) {
	var f field.UnderlyingSettlementTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCashAmount gets UnderlyingCashAmount, Tag 973.
func (m NoUnderlyings) GetUnderlyingCashAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCashAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCashType gets UnderlyingCashType, Tag 974.
func (m NoUnderlyings) GetUnderlyingCashType() (v enum.UnderlyingCashType, err quickfix.MessageRejectError) {
	var f field.UnderlyingCashTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingUnitOfMeasure gets UnderlyingUnitOfMeasure, Tag 998.
func (m NoUnderlyings) GetUnderlyingUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingTimeUnit gets UnderlyingTimeUnit, Tag 1000.
func (m NoUnderlyings) GetUnderlyingTimeUnit() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingTimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCapValue gets UnderlyingCapValue, Tag 1038.
func (m NoUnderlyings) GetUnderlyingCapValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCapValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUndlyInstrumentParties gets NoUndlyInstrumentParties, Tag 1058.
func (m NoUnderlyings) GetNoUndlyInstrumentParties() (NoUndlyInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetUnderlyingSettlMethod gets UnderlyingSettlMethod, Tag 1039.
func (m NoUnderlyings) GetUnderlyingSettlMethod() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSettlMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingAdjustedQuantity gets UnderlyingAdjustedQuantity, Tag 1044.
func (m NoUnderlyings) GetUnderlyingAdjustedQuantity() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingAdjustedQuantityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingFXRate gets UnderlyingFXRate, Tag 1045.
func (m NoUnderlyings) GetUnderlyingFXRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingFXRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingFXRateCalc gets UnderlyingFXRateCalc, Tag 1046.
func (m NoUnderlyings) GetUnderlyingFXRateCalc() (v enum.UnderlyingFXRateCalc, err quickfix.MessageRejectError) {
	var f field.UnderlyingFXRateCalcField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingMaturityTime gets UnderlyingMaturityTime, Tag 1213.
func (m NoUnderlyings) GetUnderlyingMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPutOrCall gets UnderlyingPutOrCall, Tag 315.
func (m NoUnderlyings) GetUnderlyingPutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingPutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingExerciseStyle gets UnderlyingExerciseStyle, Tag 1419.
func (m NoUnderlyings) GetUnderlyingExerciseStyle() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingUnitOfMeasureQty gets UnderlyingUnitOfMeasureQty, Tag 1423.
func (m NoUnderlyings) GetUnderlyingUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPriceUnitOfMeasure gets UnderlyingPriceUnitOfMeasure, Tag 1424.
func (m NoUnderlyings) GetUnderlyingPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingPriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPriceUnitOfMeasureQty gets UnderlyingPriceUnitOfMeasureQty, Tag 1425.
func (m NoUnderlyings) GetUnderlyingPriceUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingPriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingContractMultiplierUnit gets UnderlyingContractMultiplierUnit, Tag 1437.
func (m NoUnderlyings) GetUnderlyingContractMultiplierUnit() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingContractMultiplierUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingFlowScheduleType gets UnderlyingFlowScheduleType, Tag 1441.
func (m NoUnderlyings) GetUnderlyingFlowScheduleType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingFlowScheduleTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRestructuringType gets UnderlyingRestructuringType, Tag 1453.
func (m NoUnderlyings) GetUnderlyingRestructuringType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingRestructuringTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSeniority gets UnderlyingSeniority, Tag 1454.
func (m NoUnderlyings) GetUnderlyingSeniority() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSeniorityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingNotionalPercentageOutstanding gets UnderlyingNotionalPercentageOutstanding, Tag 1455.
func (m NoUnderlyings) GetUnderlyingNotionalPercentageOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingNotionalPercentageOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingOriginalNotionalPercentageOutstanding gets UnderlyingOriginalNotionalPercentageOutstanding, Tag 1456.
func (m NoUnderlyings) GetUnderlyingOriginalNotionalPercentageOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingOriginalNotionalPercentageOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingAttachmentPoint gets UnderlyingAttachmentPoint, Tag 1459.
func (m NoUnderlyings) GetUnderlyingAttachmentPoint() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingAttachmentPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingDetachmentPoint gets UnderlyingDetachmentPoint, Tag 1460.
func (m NoUnderlyings) GetUnderlyingDetachmentPoint() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingDetachmentPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCollAction gets CollAction, Tag 944.
func (m NoUnderlyings) GetCollAction() (v enum.CollAction, err quickfix.MessageRejectError) {
	var f field.CollActionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingSymbol returns true if UnderlyingSymbol is present, Tag 311.
func (m NoUnderlyings) HasUnderlyingSymbol() bool {
	return m.Has(tag.UnderlyingSymbol)
}

// HasUnderlyingSymbolSfx returns true if UnderlyingSymbolSfx is present, Tag 312.
func (m NoUnderlyings) HasUnderlyingSymbolSfx() bool {
	return m.Has(tag.UnderlyingSymbolSfx)
}

// HasUnderlyingSecurityID returns true if UnderlyingSecurityID is present, Tag 309.
func (m NoUnderlyings) HasUnderlyingSecurityID() bool {
	return m.Has(tag.UnderlyingSecurityID)
}

// HasUnderlyingSecurityIDSource returns true if UnderlyingSecurityIDSource is present, Tag 305.
func (m NoUnderlyings) HasUnderlyingSecurityIDSource() bool {
	return m.Has(tag.UnderlyingSecurityIDSource)
}

// HasNoUnderlyingSecurityAltID returns true if NoUnderlyingSecurityAltID is present, Tag 457.
func (m NoUnderlyings) HasNoUnderlyingSecurityAltID() bool {
	return m.Has(tag.NoUnderlyingSecurityAltID)
}

// HasUnderlyingProduct returns true if UnderlyingProduct is present, Tag 462.
func (m NoUnderlyings) HasUnderlyingProduct() bool {
	return m.Has(tag.UnderlyingProduct)
}

// HasUnderlyingCFICode returns true if UnderlyingCFICode is present, Tag 463.
func (m NoUnderlyings) HasUnderlyingCFICode() bool {
	return m.Has(tag.UnderlyingCFICode)
}

// HasUnderlyingSecurityType returns true if UnderlyingSecurityType is present, Tag 310.
func (m NoUnderlyings) HasUnderlyingSecurityType() bool {
	return m.Has(tag.UnderlyingSecurityType)
}

// HasUnderlyingSecuritySubType returns true if UnderlyingSecuritySubType is present, Tag 763.
func (m NoUnderlyings) HasUnderlyingSecuritySubType() bool {
	return m.Has(tag.UnderlyingSecuritySubType)
}

// HasUnderlyingMaturityMonthYear returns true if UnderlyingMaturityMonthYear is present, Tag 313.
func (m NoUnderlyings) HasUnderlyingMaturityMonthYear() bool {
	return m.Has(tag.UnderlyingMaturityMonthYear)
}

// HasUnderlyingMaturityDate returns true if UnderlyingMaturityDate is present, Tag 542.
func (m NoUnderlyings) HasUnderlyingMaturityDate() bool {
	return m.Has(tag.UnderlyingMaturityDate)
}

// HasUnderlyingCouponPaymentDate returns true if UnderlyingCouponPaymentDate is present, Tag 241.
func (m NoUnderlyings) HasUnderlyingCouponPaymentDate() bool {
	return m.Has(tag.UnderlyingCouponPaymentDate)
}

// HasUnderlyingIssueDate returns true if UnderlyingIssueDate is present, Tag 242.
func (m NoUnderlyings) HasUnderlyingIssueDate() bool {
	return m.Has(tag.UnderlyingIssueDate)
}

// HasUnderlyingRepoCollateralSecurityType returns true if UnderlyingRepoCollateralSecurityType is present, Tag 243.
func (m NoUnderlyings) HasUnderlyingRepoCollateralSecurityType() bool {
	return m.Has(tag.UnderlyingRepoCollateralSecurityType)
}

// HasUnderlyingRepurchaseTerm returns true if UnderlyingRepurchaseTerm is present, Tag 244.
func (m NoUnderlyings) HasUnderlyingRepurchaseTerm() bool {
	return m.Has(tag.UnderlyingRepurchaseTerm)
}

// HasUnderlyingRepurchaseRate returns true if UnderlyingRepurchaseRate is present, Tag 245.
func (m NoUnderlyings) HasUnderlyingRepurchaseRate() bool {
	return m.Has(tag.UnderlyingRepurchaseRate)
}

// HasUnderlyingFactor returns true if UnderlyingFactor is present, Tag 246.
func (m NoUnderlyings) HasUnderlyingFactor() bool {
	return m.Has(tag.UnderlyingFactor)
}

// HasUnderlyingCreditRating returns true if UnderlyingCreditRating is present, Tag 256.
func (m NoUnderlyings) HasUnderlyingCreditRating() bool {
	return m.Has(tag.UnderlyingCreditRating)
}

// HasUnderlyingInstrRegistry returns true if UnderlyingInstrRegistry is present, Tag 595.
func (m NoUnderlyings) HasUnderlyingInstrRegistry() bool {
	return m.Has(tag.UnderlyingInstrRegistry)
}

// HasUnderlyingCountryOfIssue returns true if UnderlyingCountryOfIssue is present, Tag 592.
func (m NoUnderlyings) HasUnderlyingCountryOfIssue() bool {
	return m.Has(tag.UnderlyingCountryOfIssue)
}

// HasUnderlyingStateOrProvinceOfIssue returns true if UnderlyingStateOrProvinceOfIssue is present, Tag 593.
func (m NoUnderlyings) HasUnderlyingStateOrProvinceOfIssue() bool {
	return m.Has(tag.UnderlyingStateOrProvinceOfIssue)
}

// HasUnderlyingLocaleOfIssue returns true if UnderlyingLocaleOfIssue is present, Tag 594.
func (m NoUnderlyings) HasUnderlyingLocaleOfIssue() bool {
	return m.Has(tag.UnderlyingLocaleOfIssue)
}

// HasUnderlyingRedemptionDate returns true if UnderlyingRedemptionDate is present, Tag 247.
func (m NoUnderlyings) HasUnderlyingRedemptionDate() bool {
	return m.Has(tag.UnderlyingRedemptionDate)
}

// HasUnderlyingStrikePrice returns true if UnderlyingStrikePrice is present, Tag 316.
func (m NoUnderlyings) HasUnderlyingStrikePrice() bool {
	return m.Has(tag.UnderlyingStrikePrice)
}

// HasUnderlyingStrikeCurrency returns true if UnderlyingStrikeCurrency is present, Tag 941.
func (m NoUnderlyings) HasUnderlyingStrikeCurrency() bool {
	return m.Has(tag.UnderlyingStrikeCurrency)
}

// HasUnderlyingOptAttribute returns true if UnderlyingOptAttribute is present, Tag 317.
func (m NoUnderlyings) HasUnderlyingOptAttribute() bool {
	return m.Has(tag.UnderlyingOptAttribute)
}

// HasUnderlyingContractMultiplier returns true if UnderlyingContractMultiplier is present, Tag 436.
func (m NoUnderlyings) HasUnderlyingContractMultiplier() bool {
	return m.Has(tag.UnderlyingContractMultiplier)
}

// HasUnderlyingCouponRate returns true if UnderlyingCouponRate is present, Tag 435.
func (m NoUnderlyings) HasUnderlyingCouponRate() bool {
	return m.Has(tag.UnderlyingCouponRate)
}

// HasUnderlyingSecurityExchange returns true if UnderlyingSecurityExchange is present, Tag 308.
func (m NoUnderlyings) HasUnderlyingSecurityExchange() bool {
	return m.Has(tag.UnderlyingSecurityExchange)
}

// HasUnderlyingIssuer returns true if UnderlyingIssuer is present, Tag 306.
func (m NoUnderlyings) HasUnderlyingIssuer() bool {
	return m.Has(tag.UnderlyingIssuer)
}

// HasEncodedUnderlyingIssuerLen returns true if EncodedUnderlyingIssuerLen is present, Tag 362.
func (m NoUnderlyings) HasEncodedUnderlyingIssuerLen() bool {
	return m.Has(tag.EncodedUnderlyingIssuerLen)
}

// HasEncodedUnderlyingIssuer returns true if EncodedUnderlyingIssuer is present, Tag 363.
func (m NoUnderlyings) HasEncodedUnderlyingIssuer() bool {
	return m.Has(tag.EncodedUnderlyingIssuer)
}

// HasUnderlyingSecurityDesc returns true if UnderlyingSecurityDesc is present, Tag 307.
func (m NoUnderlyings) HasUnderlyingSecurityDesc() bool {
	return m.Has(tag.UnderlyingSecurityDesc)
}

// HasEncodedUnderlyingSecurityDescLen returns true if EncodedUnderlyingSecurityDescLen is present, Tag 364.
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDescLen() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDescLen)
}

// HasEncodedUnderlyingSecurityDesc returns true if EncodedUnderlyingSecurityDesc is present, Tag 365.
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDesc() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDesc)
}

// HasUnderlyingCPProgram returns true if UnderlyingCPProgram is present, Tag 877.
func (m NoUnderlyings) HasUnderlyingCPProgram() bool {
	return m.Has(tag.UnderlyingCPProgram)
}

// HasUnderlyingCPRegType returns true if UnderlyingCPRegType is present, Tag 878.
func (m NoUnderlyings) HasUnderlyingCPRegType() bool {
	return m.Has(tag.UnderlyingCPRegType)
}

// HasUnderlyingCurrency returns true if UnderlyingCurrency is present, Tag 318.
func (m NoUnderlyings) HasUnderlyingCurrency() bool {
	return m.Has(tag.UnderlyingCurrency)
}

// HasUnderlyingQty returns true if UnderlyingQty is present, Tag 879.
func (m NoUnderlyings) HasUnderlyingQty() bool {
	return m.Has(tag.UnderlyingQty)
}

// HasUnderlyingPx returns true if UnderlyingPx is present, Tag 810.
func (m NoUnderlyings) HasUnderlyingPx() bool {
	return m.Has(tag.UnderlyingPx)
}

// HasUnderlyingDirtyPrice returns true if UnderlyingDirtyPrice is present, Tag 882.
func (m NoUnderlyings) HasUnderlyingDirtyPrice() bool {
	return m.Has(tag.UnderlyingDirtyPrice)
}

// HasUnderlyingEndPrice returns true if UnderlyingEndPrice is present, Tag 883.
func (m NoUnderlyings) HasUnderlyingEndPrice() bool {
	return m.Has(tag.UnderlyingEndPrice)
}

// HasUnderlyingStartValue returns true if UnderlyingStartValue is present, Tag 884.
func (m NoUnderlyings) HasUnderlyingStartValue() bool {
	return m.Has(tag.UnderlyingStartValue)
}

// HasUnderlyingCurrentValue returns true if UnderlyingCurrentValue is present, Tag 885.
func (m NoUnderlyings) HasUnderlyingCurrentValue() bool {
	return m.Has(tag.UnderlyingCurrentValue)
}

// HasUnderlyingEndValue returns true if UnderlyingEndValue is present, Tag 886.
func (m NoUnderlyings) HasUnderlyingEndValue() bool {
	return m.Has(tag.UnderlyingEndValue)
}

// HasNoUnderlyingStips returns true if NoUnderlyingStips is present, Tag 887.
func (m NoUnderlyings) HasNoUnderlyingStips() bool {
	return m.Has(tag.NoUnderlyingStips)
}

// HasUnderlyingAllocationPercent returns true if UnderlyingAllocationPercent is present, Tag 972.
func (m NoUnderlyings) HasUnderlyingAllocationPercent() bool {
	return m.Has(tag.UnderlyingAllocationPercent)
}

// HasUnderlyingSettlementType returns true if UnderlyingSettlementType is present, Tag 975.
func (m NoUnderlyings) HasUnderlyingSettlementType() bool {
	return m.Has(tag.UnderlyingSettlementType)
}

// HasUnderlyingCashAmount returns true if UnderlyingCashAmount is present, Tag 973.
func (m NoUnderlyings) HasUnderlyingCashAmount() bool {
	return m.Has(tag.UnderlyingCashAmount)
}

// HasUnderlyingCashType returns true if UnderlyingCashType is present, Tag 974.
func (m NoUnderlyings) HasUnderlyingCashType() bool {
	return m.Has(tag.UnderlyingCashType)
}

// HasUnderlyingUnitOfMeasure returns true if UnderlyingUnitOfMeasure is present, Tag 998.
func (m NoUnderlyings) HasUnderlyingUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingUnitOfMeasure)
}

// HasUnderlyingTimeUnit returns true if UnderlyingTimeUnit is present, Tag 1000.
func (m NoUnderlyings) HasUnderlyingTimeUnit() bool {
	return m.Has(tag.UnderlyingTimeUnit)
}

// HasUnderlyingCapValue returns true if UnderlyingCapValue is present, Tag 1038.
func (m NoUnderlyings) HasUnderlyingCapValue() bool {
	return m.Has(tag.UnderlyingCapValue)
}

// HasNoUndlyInstrumentParties returns true if NoUndlyInstrumentParties is present, Tag 1058.
func (m NoUnderlyings) HasNoUndlyInstrumentParties() bool {
	return m.Has(tag.NoUndlyInstrumentParties)
}

// HasUnderlyingSettlMethod returns true if UnderlyingSettlMethod is present, Tag 1039.
func (m NoUnderlyings) HasUnderlyingSettlMethod() bool {
	return m.Has(tag.UnderlyingSettlMethod)
}

// HasUnderlyingAdjustedQuantity returns true if UnderlyingAdjustedQuantity is present, Tag 1044.
func (m NoUnderlyings) HasUnderlyingAdjustedQuantity() bool {
	return m.Has(tag.UnderlyingAdjustedQuantity)
}

// HasUnderlyingFXRate returns true if UnderlyingFXRate is present, Tag 1045.
func (m NoUnderlyings) HasUnderlyingFXRate() bool {
	return m.Has(tag.UnderlyingFXRate)
}

// HasUnderlyingFXRateCalc returns true if UnderlyingFXRateCalc is present, Tag 1046.
func (m NoUnderlyings) HasUnderlyingFXRateCalc() bool {
	return m.Has(tag.UnderlyingFXRateCalc)
}

// HasUnderlyingMaturityTime returns true if UnderlyingMaturityTime is present, Tag 1213.
func (m NoUnderlyings) HasUnderlyingMaturityTime() bool {
	return m.Has(tag.UnderlyingMaturityTime)
}

// HasUnderlyingPutOrCall returns true if UnderlyingPutOrCall is present, Tag 315.
func (m NoUnderlyings) HasUnderlyingPutOrCall() bool {
	return m.Has(tag.UnderlyingPutOrCall)
}

// HasUnderlyingExerciseStyle returns true if UnderlyingExerciseStyle is present, Tag 1419.
func (m NoUnderlyings) HasUnderlyingExerciseStyle() bool {
	return m.Has(tag.UnderlyingExerciseStyle)
}

// HasUnderlyingUnitOfMeasureQty returns true if UnderlyingUnitOfMeasureQty is present, Tag 1423.
func (m NoUnderlyings) HasUnderlyingUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingUnitOfMeasureQty)
}

// HasUnderlyingPriceUnitOfMeasure returns true if UnderlyingPriceUnitOfMeasure is present, Tag 1424.
func (m NoUnderlyings) HasUnderlyingPriceUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasure)
}

// HasUnderlyingPriceUnitOfMeasureQty returns true if UnderlyingPriceUnitOfMeasureQty is present, Tag 1425.
func (m NoUnderlyings) HasUnderlyingPriceUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasureQty)
}

// HasUnderlyingContractMultiplierUnit returns true if UnderlyingContractMultiplierUnit is present, Tag 1437.
func (m NoUnderlyings) HasUnderlyingContractMultiplierUnit() bool {
	return m.Has(tag.UnderlyingContractMultiplierUnit)
}

// HasUnderlyingFlowScheduleType returns true if UnderlyingFlowScheduleType is present, Tag 1441.
func (m NoUnderlyings) HasUnderlyingFlowScheduleType() bool {
	return m.Has(tag.UnderlyingFlowScheduleType)
}

// HasUnderlyingRestructuringType returns true if UnderlyingRestructuringType is present, Tag 1453.
func (m NoUnderlyings) HasUnderlyingRestructuringType() bool {
	return m.Has(tag.UnderlyingRestructuringType)
}

// HasUnderlyingSeniority returns true if UnderlyingSeniority is present, Tag 1454.
func (m NoUnderlyings) HasUnderlyingSeniority() bool {
	return m.Has(tag.UnderlyingSeniority)
}

// HasUnderlyingNotionalPercentageOutstanding returns true if UnderlyingNotionalPercentageOutstanding is present, Tag 1455.
func (m NoUnderlyings) HasUnderlyingNotionalPercentageOutstanding() bool {
	return m.Has(tag.UnderlyingNotionalPercentageOutstanding)
}

// HasUnderlyingOriginalNotionalPercentageOutstanding returns true if UnderlyingOriginalNotionalPercentageOutstanding is present, Tag 1456.
func (m NoUnderlyings) HasUnderlyingOriginalNotionalPercentageOutstanding() bool {
	return m.Has(tag.UnderlyingOriginalNotionalPercentageOutstanding)
}

// HasUnderlyingAttachmentPoint returns true if UnderlyingAttachmentPoint is present, Tag 1459.
func (m NoUnderlyings) HasUnderlyingAttachmentPoint() bool {
	return m.Has(tag.UnderlyingAttachmentPoint)
}

// HasUnderlyingDetachmentPoint returns true if UnderlyingDetachmentPoint is present, Tag 1460.
func (m NoUnderlyings) HasUnderlyingDetachmentPoint() bool {
	return m.Has(tag.UnderlyingDetachmentPoint)
}

// HasCollAction returns true if CollAction is present, Tag 944.
func (m NoUnderlyings) HasCollAction() bool {
	return m.Has(tag.CollAction)
}

// NoUnderlyingSecurityAltID is a repeating group element, Tag 457.
type NoUnderlyingSecurityAltID struct {
	*quickfix.Group
}

// SetUnderlyingSecurityAltID sets UnderlyingSecurityAltID, Tag 458.
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltID(v string) {
	m.Set(field.NewUnderlyingSecurityAltID(v))
}

// SetUnderlyingSecurityAltIDSource sets UnderlyingSecurityAltIDSource, Tag 459.
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityAltIDSource(v))
}

// GetUnderlyingSecurityAltID gets UnderlyingSecurityAltID, Tag 458.
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityAltIDSource gets UnderlyingSecurityAltIDSource, Tag 459.
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingSecurityAltID returns true if UnderlyingSecurityAltID is present, Tag 458.
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltID() bool {
	return m.Has(tag.UnderlyingSecurityAltID)
}

// HasUnderlyingSecurityAltIDSource returns true if UnderlyingSecurityAltIDSource is present, Tag 459.
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltIDSource() bool {
	return m.Has(tag.UnderlyingSecurityAltIDSource)
}

// NoUnderlyingSecurityAltIDRepeatingGroup is a repeating group, Tag 457.
type NoUnderlyingSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUnderlyingSecurityAltIDRepeatingGroup returns an initialized, NoUnderlyingSecurityAltIDRepeatingGroup.
func NewNoUnderlyingSecurityAltIDRepeatingGroup() NoUnderlyingSecurityAltIDRepeatingGroup {
	return NoUnderlyingSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSecurityAltID), quickfix.GroupElement(tag.UnderlyingSecurityAltIDSource)})}
}

// Add create and append a new NoUnderlyingSecurityAltID to this group.
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Add() NoUnderlyingSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingSecurityAltID{g}
}

// Get returns the ith NoUnderlyingSecurityAltID in the NoUnderlyingSecurityAltIDRepeatinGroup.
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Get(i int) NoUnderlyingSecurityAltID {
	return NoUnderlyingSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoUnderlyingStips is a repeating group element, Tag 887.
type NoUnderlyingStips struct {
	*quickfix.Group
}

// SetUnderlyingStipType sets UnderlyingStipType, Tag 888.
func (m NoUnderlyingStips) SetUnderlyingStipType(v string) {
	m.Set(field.NewUnderlyingStipType(v))
}

// SetUnderlyingStipValue sets UnderlyingStipValue, Tag 889.
func (m NoUnderlyingStips) SetUnderlyingStipValue(v string) {
	m.Set(field.NewUnderlyingStipValue(v))
}

// GetUnderlyingStipType gets UnderlyingStipType, Tag 888.
func (m NoUnderlyingStips) GetUnderlyingStipType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStipTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStipValue gets UnderlyingStipValue, Tag 889.
func (m NoUnderlyingStips) GetUnderlyingStipValue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStipValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingStipType returns true if UnderlyingStipType is present, Tag 888.
func (m NoUnderlyingStips) HasUnderlyingStipType() bool {
	return m.Has(tag.UnderlyingStipType)
}

// HasUnderlyingStipValue returns true if UnderlyingStipValue is present, Tag 889.
func (m NoUnderlyingStips) HasUnderlyingStipValue() bool {
	return m.Has(tag.UnderlyingStipValue)
}

// NoUnderlyingStipsRepeatingGroup is a repeating group, Tag 887.
type NoUnderlyingStipsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUnderlyingStipsRepeatingGroup returns an initialized, NoUnderlyingStipsRepeatingGroup.
func NewNoUnderlyingStipsRepeatingGroup() NoUnderlyingStipsRepeatingGroup {
	return NoUnderlyingStipsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingStips,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingStipType), quickfix.GroupElement(tag.UnderlyingStipValue)})}
}

// Add create and append a new NoUnderlyingStips to this group.
func (m NoUnderlyingStipsRepeatingGroup) Add() NoUnderlyingStips {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingStips{g}
}

// Get returns the ith NoUnderlyingStips in the NoUnderlyingStipsRepeatinGroup.
func (m NoUnderlyingStipsRepeatingGroup) Get(i int) NoUnderlyingStips {
	return NoUnderlyingStips{m.RepeatingGroup.Get(i)}
}

// NoUndlyInstrumentParties is a repeating group element, Tag 1058.
type NoUndlyInstrumentParties struct {
	*quickfix.Group
}

// SetUnderlyingInstrumentPartyID sets UnderlyingInstrumentPartyID, Tag 1059.
func (m NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyID(v string) {
	m.Set(field.NewUnderlyingInstrumentPartyID(v))
}

// SetUnderlyingInstrumentPartyIDSource sets UnderlyingInstrumentPartyIDSource, Tag 1060.
func (m NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyIDSource(v string) {
	m.Set(field.NewUnderlyingInstrumentPartyIDSource(v))
}

// SetUnderlyingInstrumentPartyRole sets UnderlyingInstrumentPartyRole, Tag 1061.
func (m NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyRole(v int) {
	m.Set(field.NewUnderlyingInstrumentPartyRole(v))
}

// SetNoUndlyInstrumentPartySubIDs sets NoUndlyInstrumentPartySubIDs, Tag 1062.
func (m NoUndlyInstrumentParties) SetNoUndlyInstrumentPartySubIDs(f NoUndlyInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetUnderlyingInstrumentPartyID gets UnderlyingInstrumentPartyID, Tag 1059.
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingInstrumentPartyIDSource gets UnderlyingInstrumentPartyIDSource, Tag 1060.
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingInstrumentPartyRole gets UnderlyingInstrumentPartyRole, Tag 1061.
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUndlyInstrumentPartySubIDs gets NoUndlyInstrumentPartySubIDs, Tag 1062.
func (m NoUndlyInstrumentParties) GetNoUndlyInstrumentPartySubIDs() (NoUndlyInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasUnderlyingInstrumentPartyID returns true if UnderlyingInstrumentPartyID is present, Tag 1059.
func (m NoUndlyInstrumentParties) HasUnderlyingInstrumentPartyID() bool {
	return m.Has(tag.UnderlyingInstrumentPartyID)
}

// HasUnderlyingInstrumentPartyIDSource returns true if UnderlyingInstrumentPartyIDSource is present, Tag 1060.
func (m NoUndlyInstrumentParties) HasUnderlyingInstrumentPartyIDSource() bool {
	return m.Has(tag.UnderlyingInstrumentPartyIDSource)
}

// HasUnderlyingInstrumentPartyRole returns true if UnderlyingInstrumentPartyRole is present, Tag 1061.
func (m NoUndlyInstrumentParties) HasUnderlyingInstrumentPartyRole() bool {
	return m.Has(tag.UnderlyingInstrumentPartyRole)
}

// HasNoUndlyInstrumentPartySubIDs returns true if NoUndlyInstrumentPartySubIDs is present, Tag 1062.
func (m NoUndlyInstrumentParties) HasNoUndlyInstrumentPartySubIDs() bool {
	return m.Has(tag.NoUndlyInstrumentPartySubIDs)
}

// NoUndlyInstrumentPartySubIDs is a repeating group element, Tag 1062.
type NoUndlyInstrumentPartySubIDs struct {
	*quickfix.Group
}

// SetUnderlyingInstrumentPartySubID sets UnderlyingInstrumentPartySubID, Tag 1063.
func (m NoUndlyInstrumentPartySubIDs) SetUnderlyingInstrumentPartySubID(v string) {
	m.Set(field.NewUnderlyingInstrumentPartySubID(v))
}

// SetUnderlyingInstrumentPartySubIDType sets UnderlyingInstrumentPartySubIDType, Tag 1064.
func (m NoUndlyInstrumentPartySubIDs) SetUnderlyingInstrumentPartySubIDType(v int) {
	m.Set(field.NewUnderlyingInstrumentPartySubIDType(v))
}

// GetUnderlyingInstrumentPartySubID gets UnderlyingInstrumentPartySubID, Tag 1063.
func (m NoUndlyInstrumentPartySubIDs) GetUnderlyingInstrumentPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingInstrumentPartySubIDType gets UnderlyingInstrumentPartySubIDType, Tag 1064.
func (m NoUndlyInstrumentPartySubIDs) GetUnderlyingInstrumentPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingInstrumentPartySubID returns true if UnderlyingInstrumentPartySubID is present, Tag 1063.
func (m NoUndlyInstrumentPartySubIDs) HasUnderlyingInstrumentPartySubID() bool {
	return m.Has(tag.UnderlyingInstrumentPartySubID)
}

// HasUnderlyingInstrumentPartySubIDType returns true if UnderlyingInstrumentPartySubIDType is present, Tag 1064.
func (m NoUndlyInstrumentPartySubIDs) HasUnderlyingInstrumentPartySubIDType() bool {
	return m.Has(tag.UnderlyingInstrumentPartySubIDType)
}

// NoUndlyInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1062.
type NoUndlyInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUndlyInstrumentPartySubIDsRepeatingGroup returns an initialized, NoUndlyInstrumentPartySubIDsRepeatingGroup.
func NewNoUndlyInstrumentPartySubIDsRepeatingGroup() NoUndlyInstrumentPartySubIDsRepeatingGroup {
	return NoUndlyInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingInstrumentPartySubID), quickfix.GroupElement(tag.UnderlyingInstrumentPartySubIDType)})}
}

// Add create and append a new NoUndlyInstrumentPartySubIDs to this group.
func (m NoUndlyInstrumentPartySubIDsRepeatingGroup) Add() NoUndlyInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoUndlyInstrumentPartySubIDs{g}
}

// Get returns the ith NoUndlyInstrumentPartySubIDs in the NoUndlyInstrumentPartySubIDsRepeatinGroup.
func (m NoUndlyInstrumentPartySubIDsRepeatingGroup) Get(i int) NoUndlyInstrumentPartySubIDs {
	return NoUndlyInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoUndlyInstrumentPartiesRepeatingGroup is a repeating group, Tag 1058.
type NoUndlyInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUndlyInstrumentPartiesRepeatingGroup returns an initialized, NoUndlyInstrumentPartiesRepeatingGroup.
func NewNoUndlyInstrumentPartiesRepeatingGroup() NoUndlyInstrumentPartiesRepeatingGroup {
	return NoUndlyInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingInstrumentPartyID), quickfix.GroupElement(tag.UnderlyingInstrumentPartyIDSource), quickfix.GroupElement(tag.UnderlyingInstrumentPartyRole), NewNoUndlyInstrumentPartySubIDsRepeatingGroup()})}
}

// Add create and append a new NoUndlyInstrumentParties to this group.
func (m NoUndlyInstrumentPartiesRepeatingGroup) Add() NoUndlyInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoUndlyInstrumentParties{g}
}

// Get returns the ith NoUndlyInstrumentParties in the NoUndlyInstrumentPartiesRepeatinGroup.
func (m NoUndlyInstrumentPartiesRepeatingGroup) Get(i int) NoUndlyInstrumentParties {
	return NoUndlyInstrumentParties{m.RepeatingGroup.Get(i)}
}

// NoUnderlyingsRepeatingGroup is a repeating group, Tag 711.
type NoUnderlyingsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUnderlyingsRepeatingGroup returns an initialized, NoUnderlyingsRepeatingGroup.
func NewNoUnderlyingsRepeatingGroup() NoUnderlyingsRepeatingGroup {
	return NoUnderlyingsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyings,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingSecurityIDSource), NewNoUnderlyingSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingProduct), quickfix.GroupElement(tag.UnderlyingCFICode), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingSecuritySubType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDate), quickfix.GroupElement(tag.UnderlyingCouponPaymentDate), quickfix.GroupElement(tag.UnderlyingIssueDate), quickfix.GroupElement(tag.UnderlyingRepoCollateralSecurityType), quickfix.GroupElement(tag.UnderlyingRepurchaseTerm), quickfix.GroupElement(tag.UnderlyingRepurchaseRate), quickfix.GroupElement(tag.UnderlyingFactor), quickfix.GroupElement(tag.UnderlyingCreditRating), quickfix.GroupElement(tag.UnderlyingInstrRegistry), quickfix.GroupElement(tag.UnderlyingCountryOfIssue), quickfix.GroupElement(tag.UnderlyingStateOrProvinceOfIssue), quickfix.GroupElement(tag.UnderlyingLocaleOfIssue), quickfix.GroupElement(tag.UnderlyingRedemptionDate), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingStrikeCurrency), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.UnderlyingCPProgram), quickfix.GroupElement(tag.UnderlyingCPRegType), quickfix.GroupElement(tag.UnderlyingCurrency), quickfix.GroupElement(tag.UnderlyingQty), quickfix.GroupElement(tag.UnderlyingPx), quickfix.GroupElement(tag.UnderlyingDirtyPrice), quickfix.GroupElement(tag.UnderlyingEndPrice), quickfix.GroupElement(tag.UnderlyingStartValue), quickfix.GroupElement(tag.UnderlyingCurrentValue), quickfix.GroupElement(tag.UnderlyingEndValue), NewNoUnderlyingStipsRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingAllocationPercent), quickfix.GroupElement(tag.UnderlyingSettlementType), quickfix.GroupElement(tag.UnderlyingCashAmount), quickfix.GroupElement(tag.UnderlyingCashType), quickfix.GroupElement(tag.UnderlyingUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingTimeUnit), quickfix.GroupElement(tag.UnderlyingCapValue), NewNoUndlyInstrumentPartiesRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingSettlMethod), quickfix.GroupElement(tag.UnderlyingAdjustedQuantity), quickfix.GroupElement(tag.UnderlyingFXRate), quickfix.GroupElement(tag.UnderlyingFXRateCalc), quickfix.GroupElement(tag.UnderlyingMaturityTime), quickfix.GroupElement(tag.UnderlyingPutOrCall), quickfix.GroupElement(tag.UnderlyingExerciseStyle), quickfix.GroupElement(tag.UnderlyingUnitOfMeasureQty), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasureQty), quickfix.GroupElement(tag.UnderlyingContractMultiplierUnit), quickfix.GroupElement(tag.UnderlyingFlowScheduleType), quickfix.GroupElement(tag.UnderlyingRestructuringType), quickfix.GroupElement(tag.UnderlyingSeniority), quickfix.GroupElement(tag.UnderlyingNotionalPercentageOutstanding), quickfix.GroupElement(tag.UnderlyingOriginalNotionalPercentageOutstanding), quickfix.GroupElement(tag.UnderlyingAttachmentPoint), quickfix.GroupElement(tag.UnderlyingDetachmentPoint), quickfix.GroupElement(tag.CollAction)})}
}

// Add create and append a new NoUnderlyings to this group.
func (m NoUnderlyingsRepeatingGroup) Add() NoUnderlyings {
	g := m.RepeatingGroup.Add()
	return NoUnderlyings{g}
}

// Get returns the ith NoUnderlyings in the NoUnderlyingsRepeatinGroup.
func (m NoUnderlyingsRepeatingGroup) Get(i int) NoUnderlyings {
	return NoUnderlyings{m.RepeatingGroup.Get(i)}
}

// NoTrdRegTimestamps is a repeating group element, Tag 768.
type NoTrdRegTimestamps struct {
	*quickfix.Group
}

// SetTrdRegTimestamp sets TrdRegTimestamp, Tag 769.
func (m NoTrdRegTimestamps) SetTrdRegTimestamp(v time.Time) {
	m.Set(field.NewTrdRegTimestamp(v))
}

// SetTrdRegTimestampType sets TrdRegTimestampType, Tag 770.
func (m NoTrdRegTimestamps) SetTrdRegTimestampType(v enum.TrdRegTimestampType) {
	m.Set(field.NewTrdRegTimestampType(v))
}

// SetTrdRegTimestampOrigin sets TrdRegTimestampOrigin, Tag 771.
func (m NoTrdRegTimestamps) SetTrdRegTimestampOrigin(v string) {
	m.Set(field.NewTrdRegTimestampOrigin(v))
}

// SetDeskType sets DeskType, Tag 1033.
func (m NoTrdRegTimestamps) SetDeskType(v enum.DeskType) {
	m.Set(field.NewDeskType(v))
}

// SetDeskTypeSource sets DeskTypeSource, Tag 1034.
func (m NoTrdRegTimestamps) SetDeskTypeSource(v enum.DeskTypeSource) {
	m.Set(field.NewDeskTypeSource(v))
}

// SetDeskOrderHandlingInst sets DeskOrderHandlingInst, Tag 1035.
func (m NoTrdRegTimestamps) SetDeskOrderHandlingInst(v enum.DeskOrderHandlingInst) {
	m.Set(field.NewDeskOrderHandlingInst(v))
}

// GetTrdRegTimestamp gets TrdRegTimestamp, Tag 769.
func (m NoTrdRegTimestamps) GetTrdRegTimestamp() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TrdRegTimestampField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTrdRegTimestampType gets TrdRegTimestampType, Tag 770.
func (m NoTrdRegTimestamps) GetTrdRegTimestampType() (v enum.TrdRegTimestampType, err quickfix.MessageRejectError) {
	var f field.TrdRegTimestampTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTrdRegTimestampOrigin gets TrdRegTimestampOrigin, Tag 771.
func (m NoTrdRegTimestamps) GetTrdRegTimestampOrigin() (v string, err quickfix.MessageRejectError) {
	var f field.TrdRegTimestampOriginField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDeskType gets DeskType, Tag 1033.
func (m NoTrdRegTimestamps) GetDeskType() (v enum.DeskType, err quickfix.MessageRejectError) {
	var f field.DeskTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDeskTypeSource gets DeskTypeSource, Tag 1034.
func (m NoTrdRegTimestamps) GetDeskTypeSource() (v enum.DeskTypeSource, err quickfix.MessageRejectError) {
	var f field.DeskTypeSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDeskOrderHandlingInst gets DeskOrderHandlingInst, Tag 1035.
func (m NoTrdRegTimestamps) GetDeskOrderHandlingInst() (v enum.DeskOrderHandlingInst, err quickfix.MessageRejectError) {
	var f field.DeskOrderHandlingInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasTrdRegTimestamp returns true if TrdRegTimestamp is present, Tag 769.
func (m NoTrdRegTimestamps) HasTrdRegTimestamp() bool {
	return m.Has(tag.TrdRegTimestamp)
}

// HasTrdRegTimestampType returns true if TrdRegTimestampType is present, Tag 770.
func (m NoTrdRegTimestamps) HasTrdRegTimestampType() bool {
	return m.Has(tag.TrdRegTimestampType)
}

// HasTrdRegTimestampOrigin returns true if TrdRegTimestampOrigin is present, Tag 771.
func (m NoTrdRegTimestamps) HasTrdRegTimestampOrigin() bool {
	return m.Has(tag.TrdRegTimestampOrigin)
}

// HasDeskType returns true if DeskType is present, Tag 1033.
func (m NoTrdRegTimestamps) HasDeskType() bool {
	return m.Has(tag.DeskType)
}

// HasDeskTypeSource returns true if DeskTypeSource is present, Tag 1034.
func (m NoTrdRegTimestamps) HasDeskTypeSource() bool {
	return m.Has(tag.DeskTypeSource)
}

// HasDeskOrderHandlingInst returns true if DeskOrderHandlingInst is present, Tag 1035.
func (m NoTrdRegTimestamps) HasDeskOrderHandlingInst() bool {
	return m.Has(tag.DeskOrderHandlingInst)
}

// NoTrdRegTimestampsRepeatingGroup is a repeating group, Tag 768.
type NoTrdRegTimestampsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoTrdRegTimestampsRepeatingGroup returns an initialized, NoTrdRegTimestampsRepeatingGroup.
func NewNoTrdRegTimestampsRepeatingGroup() NoTrdRegTimestampsRepeatingGroup {
	return NoTrdRegTimestampsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTrdRegTimestamps,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TrdRegTimestamp), quickfix.GroupElement(tag.TrdRegTimestampType), quickfix.GroupElement(tag.TrdRegTimestampOrigin), quickfix.GroupElement(tag.DeskType), quickfix.GroupElement(tag.DeskTypeSource), quickfix.GroupElement(tag.DeskOrderHandlingInst)})}
}

// Add create and append a new NoTrdRegTimestamps to this group.
func (m NoTrdRegTimestampsRepeatingGroup) Add() NoTrdRegTimestamps {
	g := m.RepeatingGroup.Add()
	return NoTrdRegTimestamps{g}
}

// Get returns the ith NoTrdRegTimestamps in the NoTrdRegTimestampsRepeatinGroup.
func (m NoTrdRegTimestampsRepeatingGroup) Get(i int) NoTrdRegTimestamps {
	return NoTrdRegTimestamps{m.RepeatingGroup.Get(i)}
}

// NoEvents is a repeating group element, Tag 864.
type NoEvents struct {
	*quickfix.Group
}

// SetEventType sets EventType, Tag 865.
func (m NoEvents) SetEventType(v enum.EventType) {
	m.Set(field.NewEventType(v))
}

// SetEventDate sets EventDate, Tag 866.
func (m NoEvents) SetEventDate(v string) {
	m.Set(field.NewEventDate(v))
}

// SetEventPx sets EventPx, Tag 867.
func (m NoEvents) SetEventPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewEventPx(value, scale))
}

// SetEventText sets EventText, Tag 868.
func (m NoEvents) SetEventText(v string) {
	m.Set(field.NewEventText(v))
}

// SetEventTime sets EventTime, Tag 1145.
func (m NoEvents) SetEventTime(v time.Time) {
	m.Set(field.NewEventTime(v))
}

// GetEventType gets EventType, Tag 865.
func (m NoEvents) GetEventType() (v enum.EventType, err quickfix.MessageRejectError) {
	var f field.EventTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventDate gets EventDate, Tag 866.
func (m NoEvents) GetEventDate() (v string, err quickfix.MessageRejectError) {
	var f field.EventDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventPx gets EventPx, Tag 867.
func (m NoEvents) GetEventPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.EventPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventText gets EventText, Tag 868.
func (m NoEvents) GetEventText() (v string, err quickfix.MessageRejectError) {
	var f field.EventTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventTime gets EventTime, Tag 1145.
func (m NoEvents) GetEventTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.EventTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasEventType returns true if EventType is present, Tag 865.
func (m NoEvents) HasEventType() bool {
	return m.Has(tag.EventType)
}

// HasEventDate returns true if EventDate is present, Tag 866.
func (m NoEvents) HasEventDate() bool {
	return m.Has(tag.EventDate)
}

// HasEventPx returns true if EventPx is present, Tag 867.
func (m NoEvents) HasEventPx() bool {
	return m.Has(tag.EventPx)
}

// HasEventText returns true if EventText is present, Tag 868.
func (m NoEvents) HasEventText() bool {
	return m.Has(tag.EventText)
}

// HasEventTime returns true if EventTime is present, Tag 1145.
func (m NoEvents) HasEventTime() bool {
	return m.Has(tag.EventTime)
}

// NoEventsRepeatingGroup is a repeating group, Tag 864.
type NoEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoEventsRepeatingGroup returns an initialized, NoEventsRepeatingGroup.
func NewNoEventsRepeatingGroup() NoEventsRepeatingGroup {
	return NoEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.EventType), quickfix.GroupElement(tag.EventDate), quickfix.GroupElement(tag.EventPx), quickfix.GroupElement(tag.EventText), quickfix.GroupElement(tag.EventTime)})}
}

// Add create and append a new NoEvents to this group.
func (m NoEventsRepeatingGroup) Add() NoEvents {
	g := m.RepeatingGroup.Add()
	return NoEvents{g}
}

// Get returns the ith NoEvents in the NoEventsRepeatinGroup.
func (m NoEventsRepeatingGroup) Get(i int) NoEvents {
	return NoEvents{m.RepeatingGroup.Get(i)}
}

// NoTrades is a repeating group element, Tag 897.
type NoTrades struct {
	*quickfix.Group
}

// SetTradeReportID sets TradeReportID, Tag 571.
func (m NoTrades) SetTradeReportID(v string) {
	m.Set(field.NewTradeReportID(v))
}

// SetSecondaryTradeReportID sets SecondaryTradeReportID, Tag 818.
func (m NoTrades) SetSecondaryTradeReportID(v string) {
	m.Set(field.NewSecondaryTradeReportID(v))
}

// GetTradeReportID gets TradeReportID, Tag 571.
func (m NoTrades) GetTradeReportID() (v string, err quickfix.MessageRejectError) {
	var f field.TradeReportIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecondaryTradeReportID gets SecondaryTradeReportID, Tag 818.
func (m NoTrades) GetSecondaryTradeReportID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryTradeReportIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasTradeReportID returns true if TradeReportID is present, Tag 571.
func (m NoTrades) HasTradeReportID() bool {
	return m.Has(tag.TradeReportID)
}

// HasSecondaryTradeReportID returns true if SecondaryTradeReportID is present, Tag 818.
func (m NoTrades) HasSecondaryTradeReportID() bool {
	return m.Has(tag.SecondaryTradeReportID)
}

// NoTradesRepeatingGroup is a repeating group, Tag 897.
type NoTradesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoTradesRepeatingGroup returns an initialized, NoTradesRepeatingGroup.
func NewNoTradesRepeatingGroup() NoTradesRepeatingGroup {
	return NoTradesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTrades,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TradeReportID), quickfix.GroupElement(tag.SecondaryTradeReportID)})}
}

// Add create and append a new NoTrades to this group.
func (m NoTradesRepeatingGroup) Add() NoTrades {
	g := m.RepeatingGroup.Add()
	return NoTrades{g}
}

// Get returns the ith NoTrades in the NoTradesRepeatinGroup.
func (m NoTradesRepeatingGroup) Get(i int) NoTrades {
	return NoTrades{m.RepeatingGroup.Get(i)}
}

// NoInstrumentParties is a repeating group element, Tag 1018.
type NoInstrumentParties struct {
	*quickfix.Group
}

// SetInstrumentPartyID sets InstrumentPartyID, Tag 1019.
func (m NoInstrumentParties) SetInstrumentPartyID(v string) {
	m.Set(field.NewInstrumentPartyID(v))
}

// SetInstrumentPartyIDSource sets InstrumentPartyIDSource, Tag 1050.
func (m NoInstrumentParties) SetInstrumentPartyIDSource(v string) {
	m.Set(field.NewInstrumentPartyIDSource(v))
}

// SetInstrumentPartyRole sets InstrumentPartyRole, Tag 1051.
func (m NoInstrumentParties) SetInstrumentPartyRole(v int) {
	m.Set(field.NewInstrumentPartyRole(v))
}

// SetNoInstrumentPartySubIDs sets NoInstrumentPartySubIDs, Tag 1052.
func (m NoInstrumentParties) SetNoInstrumentPartySubIDs(f NoInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetInstrumentPartyID gets InstrumentPartyID, Tag 1019.
func (m NoInstrumentParties) GetInstrumentPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentPartyIDSource gets InstrumentPartyIDSource, Tag 1050.
func (m NoInstrumentParties) GetInstrumentPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentPartyRole gets InstrumentPartyRole, Tag 1051.
func (m NoInstrumentParties) GetInstrumentPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoInstrumentPartySubIDs gets NoInstrumentPartySubIDs, Tag 1052.
func (m NoInstrumentParties) GetNoInstrumentPartySubIDs() (NoInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasInstrumentPartyID returns true if InstrumentPartyID is present, Tag 1019.
func (m NoInstrumentParties) HasInstrumentPartyID() bool {
	return m.Has(tag.InstrumentPartyID)
}

// HasInstrumentPartyIDSource returns true if InstrumentPartyIDSource is present, Tag 1050.
func (m NoInstrumentParties) HasInstrumentPartyIDSource() bool {
	return m.Has(tag.InstrumentPartyIDSource)
}

// HasInstrumentPartyRole returns true if InstrumentPartyRole is present, Tag 1051.
func (m NoInstrumentParties) HasInstrumentPartyRole() bool {
	return m.Has(tag.InstrumentPartyRole)
}

// HasNoInstrumentPartySubIDs returns true if NoInstrumentPartySubIDs is present, Tag 1052.
func (m NoInstrumentParties) HasNoInstrumentPartySubIDs() bool {
	return m.Has(tag.NoInstrumentPartySubIDs)
}

// NoInstrumentPartySubIDs is a repeating group element, Tag 1052.
type NoInstrumentPartySubIDs struct {
	*quickfix.Group
}

// SetInstrumentPartySubID sets InstrumentPartySubID, Tag 1053.
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubID(v string) {
	m.Set(field.NewInstrumentPartySubID(v))
}

// SetInstrumentPartySubIDType sets InstrumentPartySubIDType, Tag 1054.
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubIDType(v int) {
	m.Set(field.NewInstrumentPartySubIDType(v))
}

// GetInstrumentPartySubID gets InstrumentPartySubID, Tag 1053.
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentPartySubIDType gets InstrumentPartySubIDType, Tag 1054.
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasInstrumentPartySubID returns true if InstrumentPartySubID is present, Tag 1053.
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubID() bool {
	return m.Has(tag.InstrumentPartySubID)
}

// HasInstrumentPartySubIDType returns true if InstrumentPartySubIDType is present, Tag 1054.
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubIDType() bool {
	return m.Has(tag.InstrumentPartySubIDType)
}

// NoInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1052.
type NoInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoInstrumentPartySubIDsRepeatingGroup returns an initialized, NoInstrumentPartySubIDsRepeatingGroup.
func NewNoInstrumentPartySubIDsRepeatingGroup() NoInstrumentPartySubIDsRepeatingGroup {
	return NoInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartySubID), quickfix.GroupElement(tag.InstrumentPartySubIDType)})}
}

// Add create and append a new NoInstrumentPartySubIDs to this group.
func (m NoInstrumentPartySubIDsRepeatingGroup) Add() NoInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoInstrumentPartySubIDs{g}
}

// Get returns the ith NoInstrumentPartySubIDs in the NoInstrumentPartySubIDsRepeatinGroup.
func (m NoInstrumentPartySubIDsRepeatingGroup) Get(i int) NoInstrumentPartySubIDs {
	return NoInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoInstrumentPartiesRepeatingGroup is a repeating group, Tag 1018.
type NoInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoInstrumentPartiesRepeatingGroup returns an initialized, NoInstrumentPartiesRepeatingGroup.
func NewNoInstrumentPartiesRepeatingGroup() NoInstrumentPartiesRepeatingGroup {
	return NoInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartyID), quickfix.GroupElement(tag.InstrumentPartyIDSource), quickfix.GroupElement(tag.InstrumentPartyRole), NewNoInstrumentPartySubIDsRepeatingGroup()})}
}

// Add create and append a new NoInstrumentParties to this group.
func (m NoInstrumentPartiesRepeatingGroup) Add() NoInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoInstrumentParties{g}
}

// Get returns the ith NoInstrumentParties in the NoInstrumentPartiesRepeatinGroup.
func (m NoInstrumentPartiesRepeatingGroup) Get(i int) NoInstrumentParties {
	return NoInstrumentParties{m.RepeatingGroup.Get(i)}
}

// NoComplexEvents is a repeating group element, Tag 1483.
type NoComplexEvents struct {
	*quickfix.Group
}

// SetComplexEventType sets ComplexEventType, Tag 1484.
func (m NoComplexEvents) SetComplexEventType(v enum.ComplexEventType) {
	m.Set(field.NewComplexEventType(v))
}

// SetComplexOptPayoutAmount sets ComplexOptPayoutAmount, Tag 1485.
func (m NoComplexEvents) SetComplexOptPayoutAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexOptPayoutAmount(value, scale))
}

// SetComplexEventPrice sets ComplexEventPrice, Tag 1486.
func (m NoComplexEvents) SetComplexEventPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexEventPrice(value, scale))
}

// SetComplexEventPriceBoundaryMethod sets ComplexEventPriceBoundaryMethod, Tag 1487.
func (m NoComplexEvents) SetComplexEventPriceBoundaryMethod(v enum.ComplexEventPriceBoundaryMethod) {
	m.Set(field.NewComplexEventPriceBoundaryMethod(v))
}

// SetComplexEventPriceBoundaryPrecision sets ComplexEventPriceBoundaryPrecision, Tag 1488.
func (m NoComplexEvents) SetComplexEventPriceBoundaryPrecision(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexEventPriceBoundaryPrecision(value, scale))
}

// SetComplexEventPriceTimeType sets ComplexEventPriceTimeType, Tag 1489.
func (m NoComplexEvents) SetComplexEventPriceTimeType(v enum.ComplexEventPriceTimeType) {
	m.Set(field.NewComplexEventPriceTimeType(v))
}

// SetComplexEventCondition sets ComplexEventCondition, Tag 1490.
func (m NoComplexEvents) SetComplexEventCondition(v enum.ComplexEventCondition) {
	m.Set(field.NewComplexEventCondition(v))
}

// SetNoComplexEventDates sets NoComplexEventDates, Tag 1491.
func (m NoComplexEvents) SetNoComplexEventDates(f NoComplexEventDatesRepeatingGroup) {
	m.SetGroup(f)
}

// GetComplexEventType gets ComplexEventType, Tag 1484.
func (m NoComplexEvents) GetComplexEventType() (v enum.ComplexEventType, err quickfix.MessageRejectError) {
	var f field.ComplexEventTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexOptPayoutAmount gets ComplexOptPayoutAmount, Tag 1485.
func (m NoComplexEvents) GetComplexOptPayoutAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ComplexOptPayoutAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventPrice gets ComplexEventPrice, Tag 1486.
func (m NoComplexEvents) GetComplexEventPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ComplexEventPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventPriceBoundaryMethod gets ComplexEventPriceBoundaryMethod, Tag 1487.
func (m NoComplexEvents) GetComplexEventPriceBoundaryMethod() (v enum.ComplexEventPriceBoundaryMethod, err quickfix.MessageRejectError) {
	var f field.ComplexEventPriceBoundaryMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventPriceBoundaryPrecision gets ComplexEventPriceBoundaryPrecision, Tag 1488.
func (m NoComplexEvents) GetComplexEventPriceBoundaryPrecision() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ComplexEventPriceBoundaryPrecisionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventPriceTimeType gets ComplexEventPriceTimeType, Tag 1489.
func (m NoComplexEvents) GetComplexEventPriceTimeType() (v enum.ComplexEventPriceTimeType, err quickfix.MessageRejectError) {
	var f field.ComplexEventPriceTimeTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventCondition gets ComplexEventCondition, Tag 1490.
func (m NoComplexEvents) GetComplexEventCondition() (v enum.ComplexEventCondition, err quickfix.MessageRejectError) {
	var f field.ComplexEventConditionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoComplexEventDates gets NoComplexEventDates, Tag 1491.
func (m NoComplexEvents) GetNoComplexEventDates() (NoComplexEventDatesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventDatesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasComplexEventType returns true if ComplexEventType is present, Tag 1484.
func (m NoComplexEvents) HasComplexEventType() bool {
	return m.Has(tag.ComplexEventType)
}

// HasComplexOptPayoutAmount returns true if ComplexOptPayoutAmount is present, Tag 1485.
func (m NoComplexEvents) HasComplexOptPayoutAmount() bool {
	return m.Has(tag.ComplexOptPayoutAmount)
}

// HasComplexEventPrice returns true if ComplexEventPrice is present, Tag 1486.
func (m NoComplexEvents) HasComplexEventPrice() bool {
	return m.Has(tag.ComplexEventPrice)
}

// HasComplexEventPriceBoundaryMethod returns true if ComplexEventPriceBoundaryMethod is present, Tag 1487.
func (m NoComplexEvents) HasComplexEventPriceBoundaryMethod() bool {
	return m.Has(tag.ComplexEventPriceBoundaryMethod)
}

// HasComplexEventPriceBoundaryPrecision returns true if ComplexEventPriceBoundaryPrecision is present, Tag 1488.
func (m NoComplexEvents) HasComplexEventPriceBoundaryPrecision() bool {
	return m.Has(tag.ComplexEventPriceBoundaryPrecision)
}

// HasComplexEventPriceTimeType returns true if ComplexEventPriceTimeType is present, Tag 1489.
func (m NoComplexEvents) HasComplexEventPriceTimeType() bool {
	return m.Has(tag.ComplexEventPriceTimeType)
}

// HasComplexEventCondition returns true if ComplexEventCondition is present, Tag 1490.
func (m NoComplexEvents) HasComplexEventCondition() bool {
	return m.Has(tag.ComplexEventCondition)
}

// HasNoComplexEventDates returns true if NoComplexEventDates is present, Tag 1491.
func (m NoComplexEvents) HasNoComplexEventDates() bool {
	return m.Has(tag.NoComplexEventDates)
}

// NoComplexEventDates is a repeating group element, Tag 1491.
type NoComplexEventDates struct {
	*quickfix.Group
}

// SetComplexEventStartDate sets ComplexEventStartDate, Tag 1492.
func (m NoComplexEventDates) SetComplexEventStartDate(v time.Time) {
	m.Set(field.NewComplexEventStartDate(v))
}

// SetComplexEventEndDate sets ComplexEventEndDate, Tag 1493.
func (m NoComplexEventDates) SetComplexEventEndDate(v time.Time) {
	m.Set(field.NewComplexEventEndDate(v))
}

// SetNoComplexEventTimes sets NoComplexEventTimes, Tag 1494.
func (m NoComplexEventDates) SetNoComplexEventTimes(f NoComplexEventTimesRepeatingGroup) {
	m.SetGroup(f)
}

// GetComplexEventStartDate gets ComplexEventStartDate, Tag 1492.
func (m NoComplexEventDates) GetComplexEventStartDate() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ComplexEventStartDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventEndDate gets ComplexEventEndDate, Tag 1493.
func (m NoComplexEventDates) GetComplexEventEndDate() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ComplexEventEndDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoComplexEventTimes gets NoComplexEventTimes, Tag 1494.
func (m NoComplexEventDates) GetNoComplexEventTimes() (NoComplexEventTimesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventTimesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasComplexEventStartDate returns true if ComplexEventStartDate is present, Tag 1492.
func (m NoComplexEventDates) HasComplexEventStartDate() bool {
	return m.Has(tag.ComplexEventStartDate)
}

// HasComplexEventEndDate returns true if ComplexEventEndDate is present, Tag 1493.
func (m NoComplexEventDates) HasComplexEventEndDate() bool {
	return m.Has(tag.ComplexEventEndDate)
}

// HasNoComplexEventTimes returns true if NoComplexEventTimes is present, Tag 1494.
func (m NoComplexEventDates) HasNoComplexEventTimes() bool {
	return m.Has(tag.NoComplexEventTimes)
}

// NoComplexEventTimes is a repeating group element, Tag 1494.
type NoComplexEventTimes struct {
	*quickfix.Group
}

// SetComplexEventStartTime sets ComplexEventStartTime, Tag 1495.
func (m NoComplexEventTimes) SetComplexEventStartTime(v string) {
	m.Set(field.NewComplexEventStartTime(v))
}

// SetComplexEventEndTime sets ComplexEventEndTime, Tag 1496.
func (m NoComplexEventTimes) SetComplexEventEndTime(v string) {
	m.Set(field.NewComplexEventEndTime(v))
}

// GetComplexEventStartTime gets ComplexEventStartTime, Tag 1495.
func (m NoComplexEventTimes) GetComplexEventStartTime() (v string, err quickfix.MessageRejectError) {
	var f field.ComplexEventStartTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventEndTime gets ComplexEventEndTime, Tag 1496.
func (m NoComplexEventTimes) GetComplexEventEndTime() (v string, err quickfix.MessageRejectError) {
	var f field.ComplexEventEndTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasComplexEventStartTime returns true if ComplexEventStartTime is present, Tag 1495.
func (m NoComplexEventTimes) HasComplexEventStartTime() bool {
	return m.Has(tag.ComplexEventStartTime)
}

// HasComplexEventEndTime returns true if ComplexEventEndTime is present, Tag 1496.
func (m NoComplexEventTimes) HasComplexEventEndTime() bool {
	return m.Has(tag.ComplexEventEndTime)
}

// NoComplexEventTimesRepeatingGroup is a repeating group, Tag 1494.
type NoComplexEventTimesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoComplexEventTimesRepeatingGroup returns an initialized, NoComplexEventTimesRepeatingGroup.
func NewNoComplexEventTimesRepeatingGroup() NoComplexEventTimesRepeatingGroup {
	return NoComplexEventTimesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEventTimes,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventStartTime), quickfix.GroupElement(tag.ComplexEventEndTime)})}
}

// Add create and append a new NoComplexEventTimes to this group.
func (m NoComplexEventTimesRepeatingGroup) Add() NoComplexEventTimes {
	g := m.RepeatingGroup.Add()
	return NoComplexEventTimes{g}
}

// Get returns the ith NoComplexEventTimes in the NoComplexEventTimesRepeatinGroup.
func (m NoComplexEventTimesRepeatingGroup) Get(i int) NoComplexEventTimes {
	return NoComplexEventTimes{m.RepeatingGroup.Get(i)}
}

// NoComplexEventDatesRepeatingGroup is a repeating group, Tag 1491.
type NoComplexEventDatesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoComplexEventDatesRepeatingGroup returns an initialized, NoComplexEventDatesRepeatingGroup.
func NewNoComplexEventDatesRepeatingGroup() NoComplexEventDatesRepeatingGroup {
	return NoComplexEventDatesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEventDates,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventStartDate), quickfix.GroupElement(tag.ComplexEventEndDate), NewNoComplexEventTimesRepeatingGroup()})}
}

// Add create and append a new NoComplexEventDates to this group.
func (m NoComplexEventDatesRepeatingGroup) Add() NoComplexEventDates {
	g := m.RepeatingGroup.Add()
	return NoComplexEventDates{g}
}

// Get returns the ith NoComplexEventDates in the NoComplexEventDatesRepeatinGroup.
func (m NoComplexEventDatesRepeatingGroup) Get(i int) NoComplexEventDates {
	return NoComplexEventDates{m.RepeatingGroup.Get(i)}
}

// NoComplexEventsRepeatingGroup is a repeating group, Tag 1483.
type NoComplexEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoComplexEventsRepeatingGroup returns an initialized, NoComplexEventsRepeatingGroup.
func NewNoComplexEventsRepeatingGroup() NoComplexEventsRepeatingGroup {
	return NoComplexEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventType), quickfix.GroupElement(tag.ComplexOptPayoutAmount), quickfix.GroupElement(tag.ComplexEventPrice), quickfix.GroupElement(tag.ComplexEventPriceBoundaryMethod), quickfix.GroupElement(tag.ComplexEventPriceBoundaryPrecision), quickfix.GroupElement(tag.ComplexEventPriceTimeType), quickfix.GroupElement(tag.ComplexEventCondition), NewNoComplexEventDatesRepeatingGroup()})}
}

// Add create and append a new NoComplexEvents to this group.
func (m NoComplexEventsRepeatingGroup) Add() NoComplexEvents {
	g := m.RepeatingGroup.Add()
	return NoComplexEvents{g}
}

// Get returns the ith NoComplexEvents in the NoComplexEventsRepeatinGroup.
func (m NoComplexEventsRepeatingGroup) Get(i int) NoComplexEvents {
	return NoComplexEvents{m.RepeatingGroup.Get(i)}
}
