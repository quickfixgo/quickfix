package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//AllocationInstruction msg type = J.
type AllocationInstruction struct {
	message.Message
}

//AllocationInstructionBuilder builds AllocationInstruction messages.
type AllocationInstructionBuilder struct {
	message.MessageBuilder
}

//CreateAllocationInstructionBuilder returns an initialized AllocationInstructionBuilder with specified required fields.
func CreateAllocationInstructionBuilder(
	allocid field.AllocID,
	alloctranstype field.AllocTransType,
	alloctype field.AllocType,
	allocnoorderstype field.AllocNoOrdersType,
	side field.Side,
	quantity field.Quantity,
	avgpx field.AvgPx,
	tradedate field.TradeDate) AllocationInstructionBuilder {
	var builder AllocationInstructionBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(allocid)
	builder.Body.Set(alloctranstype)
	builder.Body.Set(alloctype)
	builder.Body.Set(allocnoorderstype)
	builder.Body.Set(side)
	builder.Body.Set(quantity)
	builder.Body.Set(avgpx)
	builder.Body.Set(tradedate)
	return builder
}

//AllocID is a required field for AllocationInstruction.
func (m AllocationInstruction) AllocID() (field.AllocID, errors.MessageRejectError) {
	var f field.AllocID
	err := m.Body.Get(&f)
	return f, err
}

//AllocTransType is a required field for AllocationInstruction.
func (m AllocationInstruction) AllocTransType() (field.AllocTransType, errors.MessageRejectError) {
	var f field.AllocTransType
	err := m.Body.Get(&f)
	return f, err
}

//AllocType is a required field for AllocationInstruction.
func (m AllocationInstruction) AllocType() (field.AllocType, errors.MessageRejectError) {
	var f field.AllocType
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryAllocID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecondaryAllocID() (field.SecondaryAllocID, errors.MessageRejectError) {
	var f field.SecondaryAllocID
	err := m.Body.Get(&f)
	return f, err
}

//RefAllocID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) RefAllocID() (field.RefAllocID, errors.MessageRejectError) {
	var f field.RefAllocID
	err := m.Body.Get(&f)
	return f, err
}

//AllocCancReplaceReason is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AllocCancReplaceReason() (field.AllocCancReplaceReason, errors.MessageRejectError) {
	var f field.AllocCancReplaceReason
	err := m.Body.Get(&f)
	return f, err
}

//AllocIntermedReqType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AllocIntermedReqType() (field.AllocIntermedReqType, errors.MessageRejectError) {
	var f field.AllocIntermedReqType
	err := m.Body.Get(&f)
	return f, err
}

//AllocLinkID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AllocLinkID() (field.AllocLinkID, errors.MessageRejectError) {
	var f field.AllocLinkID
	err := m.Body.Get(&f)
	return f, err
}

//AllocLinkType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AllocLinkType() (field.AllocLinkType, errors.MessageRejectError) {
	var f field.AllocLinkType
	err := m.Body.Get(&f)
	return f, err
}

//BookingRefID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BookingRefID() (field.BookingRefID, errors.MessageRejectError) {
	var f field.BookingRefID
	err := m.Body.Get(&f)
	return f, err
}

//AllocNoOrdersType is a required field for AllocationInstruction.
func (m AllocationInstruction) AllocNoOrdersType() (field.AllocNoOrdersType, errors.MessageRejectError) {
	var f field.AllocNoOrdersType
	err := m.Body.Get(&f)
	return f, err
}

//NoOrders is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoOrders() (field.NoOrders, errors.MessageRejectError) {
	var f field.NoOrders
	err := m.Body.Get(&f)
	return f, err
}

//NoExecs is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoExecs() (field.NoExecs, errors.MessageRejectError) {
	var f field.NoExecs
	err := m.Body.Get(&f)
	return f, err
}

//PreviouslyReported is a non-required field for AllocationInstruction.
func (m AllocationInstruction) PreviouslyReported() (field.PreviouslyReported, errors.MessageRejectError) {
	var f field.PreviouslyReported
	err := m.Body.Get(&f)
	return f, err
}

//ReversalIndicator is a non-required field for AllocationInstruction.
func (m AllocationInstruction) ReversalIndicator() (field.ReversalIndicator, errors.MessageRejectError) {
	var f field.ReversalIndicator
	err := m.Body.Get(&f)
	return f, err
}

//MatchType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) MatchType() (field.MatchType, errors.MessageRejectError) {
	var f field.MatchType
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for AllocationInstruction.
func (m AllocationInstruction) Side() (field.Side, errors.MessageRejectError) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecurityIDSource() (field.SecurityIDSource, errors.MessageRejectError) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoSecurityAltID() (field.NoSecurityAltID, errors.MessageRejectError) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Product() (field.Product, errors.MessageRejectError) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for AllocationInstruction.
func (m AllocationInstruction) CFICode() (field.CFICode, errors.MessageRejectError) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecuritySubType() (field.SecuritySubType, errors.MessageRejectError) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for AllocationInstruction.
func (m AllocationInstruction) MaturityMonthYear() (field.MaturityMonthYear, errors.MessageRejectError) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) MaturityDate() (field.MaturityDate, errors.MessageRejectError) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) CouponPaymentDate() (field.CouponPaymentDate, errors.MessageRejectError) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) IssueDate() (field.IssueDate, errors.MessageRejectError) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, errors.MessageRejectError) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for AllocationInstruction.
func (m AllocationInstruction) RepurchaseTerm() (field.RepurchaseTerm, errors.MessageRejectError) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) RepurchaseRate() (field.RepurchaseRate, errors.MessageRejectError) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Factor() (field.Factor, errors.MessageRejectError) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for AllocationInstruction.
func (m AllocationInstruction) CreditRating() (field.CreditRating, errors.MessageRejectError) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for AllocationInstruction.
func (m AllocationInstruction) InstrRegistry() (field.InstrRegistry, errors.MessageRejectError) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for AllocationInstruction.
func (m AllocationInstruction) CountryOfIssue() (field.CountryOfIssue, errors.MessageRejectError) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for AllocationInstruction.
func (m AllocationInstruction) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for AllocationInstruction.
func (m AllocationInstruction) LocaleOfIssue() (field.LocaleOfIssue, errors.MessageRejectError) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) RedemptionDate() (field.RedemptionDate, errors.MessageRejectError) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for AllocationInstruction.
func (m AllocationInstruction) StrikePrice() (field.StrikePrice, errors.MessageRejectError) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for AllocationInstruction.
func (m AllocationInstruction) StrikeCurrency() (field.StrikeCurrency, errors.MessageRejectError) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for AllocationInstruction.
func (m AllocationInstruction) OptAttribute() (field.OptAttribute, errors.MessageRejectError) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for AllocationInstruction.
func (m AllocationInstruction) ContractMultiplier() (field.ContractMultiplier, errors.MessageRejectError) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) CouponRate() (field.CouponRate, errors.MessageRejectError) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EncodedIssuerLen() (field.EncodedIssuerLen, errors.MessageRejectError) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EncodedIssuer() (field.EncodedIssuer, errors.MessageRejectError) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, errors.MessageRejectError) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EncodedSecurityDesc() (field.EncodedSecurityDesc, errors.MessageRejectError) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Pool() (field.Pool, errors.MessageRejectError) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for AllocationInstruction.
func (m AllocationInstruction) ContractSettlMonth() (field.ContractSettlMonth, errors.MessageRejectError) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for AllocationInstruction.
func (m AllocationInstruction) CPProgram() (field.CPProgram, errors.MessageRejectError) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) CPRegType() (field.CPRegType, errors.MessageRejectError) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoEvents() (field.NoEvents, errors.MessageRejectError) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) DatedDate() (field.DatedDate, errors.MessageRejectError) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) InterestAccrualDate() (field.InterestAccrualDate, errors.MessageRejectError) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryForm is a non-required field for AllocationInstruction.
func (m AllocationInstruction) DeliveryForm() (field.DeliveryForm, errors.MessageRejectError) {
	var f field.DeliveryForm
	err := m.Body.Get(&f)
	return f, err
}

//PctAtRisk is a non-required field for AllocationInstruction.
func (m AllocationInstruction) PctAtRisk() (field.PctAtRisk, errors.MessageRejectError) {
	var f field.PctAtRisk
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrAttrib is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoInstrAttrib() (field.NoInstrAttrib, errors.MessageRejectError) {
	var f field.NoInstrAttrib
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AgreementDesc() (field.AgreementDesc, errors.MessageRejectError) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AgreementID() (field.AgreementID, errors.MessageRejectError) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AgreementDate() (field.AgreementDate, errors.MessageRejectError) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AgreementCurrency() (field.AgreementCurrency, errors.MessageRejectError) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TerminationType() (field.TerminationType, errors.MessageRejectError) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) StartDate() (field.StartDate, errors.MessageRejectError) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EndDate() (field.EndDate, errors.MessageRejectError) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) DeliveryType() (field.DeliveryType, errors.MessageRejectError) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for AllocationInstruction.
func (m AllocationInstruction) MarginRatio() (field.MarginRatio, errors.MessageRejectError) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoUnderlyings() (field.NoUnderlyings, errors.MessageRejectError) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoLegs() (field.NoLegs, errors.MessageRejectError) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//Quantity is a required field for AllocationInstruction.
func (m AllocationInstruction) Quantity() (field.Quantity, errors.MessageRejectError) {
	var f field.Quantity
	err := m.Body.Get(&f)
	return f, err
}

//QtyType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) QtyType() (field.QtyType, errors.MessageRejectError) {
	var f field.QtyType
	err := m.Body.Get(&f)
	return f, err
}

//LastMkt is a non-required field for AllocationInstruction.
func (m AllocationInstruction) LastMkt() (field.LastMkt, errors.MessageRejectError) {
	var f field.LastMkt
	err := m.Body.Get(&f)
	return f, err
}

//TradeOriginationDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TradeOriginationDate() (field.TradeOriginationDate, errors.MessageRejectError) {
	var f field.TradeOriginationDate
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TradingSessionID() (field.TradingSessionID, errors.MessageRejectError) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TradingSessionSubID() (field.TradingSessionSubID, errors.MessageRejectError) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) PriceType() (field.PriceType, errors.MessageRejectError) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//AvgPx is a required field for AllocationInstruction.
func (m AllocationInstruction) AvgPx() (field.AvgPx, errors.MessageRejectError) {
	var f field.AvgPx
	err := m.Body.Get(&f)
	return f, err
}

//AvgParPx is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AvgParPx() (field.AvgParPx, errors.MessageRejectError) {
	var f field.AvgParPx
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Spread() (field.Spread, errors.MessageRejectError) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BenchmarkCurveName() (field.BenchmarkCurveName, errors.MessageRejectError) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, errors.MessageRejectError) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BenchmarkPrice() (field.BenchmarkPrice, errors.MessageRejectError) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BenchmarkPriceType() (field.BenchmarkPriceType, errors.MessageRejectError) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BenchmarkSecurityID() (field.BenchmarkSecurityID, errors.MessageRejectError) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Currency() (field.Currency, errors.MessageRejectError) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//AvgPxPrecision is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AvgPxPrecision() (field.AvgPxPrecision, errors.MessageRejectError) {
	var f field.AvgPxPrecision
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoPartyIDs() (field.NoPartyIDs, errors.MessageRejectError) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a required field for AllocationInstruction.
func (m AllocationInstruction) TradeDate() (field.TradeDate, errors.MessageRejectError) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SettlType() (field.SettlType, errors.MessageRejectError) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SettlDate() (field.SettlDate, errors.MessageRejectError) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//BookingType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BookingType() (field.BookingType, errors.MessageRejectError) {
	var f field.BookingType
	err := m.Body.Get(&f)
	return f, err
}

//GrossTradeAmt is a non-required field for AllocationInstruction.
func (m AllocationInstruction) GrossTradeAmt() (field.GrossTradeAmt, errors.MessageRejectError) {
	var f field.GrossTradeAmt
	err := m.Body.Get(&f)
	return f, err
}

//Concession is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Concession() (field.Concession, errors.MessageRejectError) {
	var f field.Concession
	err := m.Body.Get(&f)
	return f, err
}

//TotalTakedown is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TotalTakedown() (field.TotalTakedown, errors.MessageRejectError) {
	var f field.TotalTakedown
	err := m.Body.Get(&f)
	return f, err
}

//NetMoney is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NetMoney() (field.NetMoney, errors.MessageRejectError) {
	var f field.NetMoney
	err := m.Body.Get(&f)
	return f, err
}

//PositionEffect is a non-required field for AllocationInstruction.
func (m AllocationInstruction) PositionEffect() (field.PositionEffect, errors.MessageRejectError) {
	var f field.PositionEffect
	err := m.Body.Get(&f)
	return f, err
}

//AutoAcceptIndicator is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AutoAcceptIndicator() (field.AutoAcceptIndicator, errors.MessageRejectError) {
	var f field.AutoAcceptIndicator
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//NumDaysInterest is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NumDaysInterest() (field.NumDaysInterest, errors.MessageRejectError) {
	var f field.NumDaysInterest
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestRate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AccruedInterestRate() (field.AccruedInterestRate, errors.MessageRejectError) {
	var f field.AccruedInterestRate
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestAmt is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AccruedInterestAmt() (field.AccruedInterestAmt, errors.MessageRejectError) {
	var f field.AccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//TotalAccruedInterestAmt is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TotalAccruedInterestAmt() (field.TotalAccruedInterestAmt, errors.MessageRejectError) {
	var f field.TotalAccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//InterestAtMaturity is a non-required field for AllocationInstruction.
func (m AllocationInstruction) InterestAtMaturity() (field.InterestAtMaturity, errors.MessageRejectError) {
	var f field.InterestAtMaturity
	err := m.Body.Get(&f)
	return f, err
}

//EndAccruedInterestAmt is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EndAccruedInterestAmt() (field.EndAccruedInterestAmt, errors.MessageRejectError) {
	var f field.EndAccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//StartCash is a non-required field for AllocationInstruction.
func (m AllocationInstruction) StartCash() (field.StartCash, errors.MessageRejectError) {
	var f field.StartCash
	err := m.Body.Get(&f)
	return f, err
}

//EndCash is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EndCash() (field.EndCash, errors.MessageRejectError) {
	var f field.EndCash
	err := m.Body.Get(&f)
	return f, err
}

//LegalConfirm is a non-required field for AllocationInstruction.
func (m AllocationInstruction) LegalConfirm() (field.LegalConfirm, errors.MessageRejectError) {
	var f field.LegalConfirm
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoStipulations() (field.NoStipulations, errors.MessageRejectError) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) YieldType() (field.YieldType, errors.MessageRejectError) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Yield() (field.Yield, errors.MessageRejectError) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//YieldCalcDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) YieldCalcDate() (field.YieldCalcDate, errors.MessageRejectError) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) YieldRedemptionDate() (field.YieldRedemptionDate, errors.MessageRejectError) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for AllocationInstruction.
func (m AllocationInstruction) YieldRedemptionPrice() (field.YieldRedemptionPrice, errors.MessageRejectError) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, errors.MessageRejectError) {
	var f field.YieldRedemptionPriceType
	err := m.Body.Get(&f)
	return f, err
}

//TotNoAllocs is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TotNoAllocs() (field.TotNoAllocs, errors.MessageRejectError) {
	var f field.TotNoAllocs
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for AllocationInstruction.
func (m AllocationInstruction) LastFragment() (field.LastFragment, errors.MessageRejectError) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoAllocs is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoAllocs() (field.NoAllocs, errors.MessageRejectError) {
	var f field.NoAllocs
	err := m.Body.Get(&f)
	return f, err
}
