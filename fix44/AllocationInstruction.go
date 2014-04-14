package fix44

import (
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

//NewAllocationInstructionBuilder returns an initialized AllocationInstructionBuilder with specified required fields.
func NewAllocationInstructionBuilder(
	allocid field.AllocID,
	alloctranstype field.AllocTransType,
	alloctype field.AllocType,
	allocnoorderstype field.AllocNoOrdersType,
	side field.Side,
	quantity field.Quantity,
	avgpx field.AvgPx,
	tradedate field.TradeDate) *AllocationInstructionBuilder {
	builder := new(AllocationInstructionBuilder)
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
func (m *AllocationInstruction) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//AllocTransType is a required field for AllocationInstruction.
func (m *AllocationInstruction) AllocTransType() (*field.AllocTransType, error) {
	f := new(field.AllocTransType)
	err := m.Body.Get(f)
	return f, err
}

//AllocType is a required field for AllocationInstruction.
func (m *AllocationInstruction) AllocType() (*field.AllocType, error) {
	f := new(field.AllocType)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryAllocID is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) SecondaryAllocID() (*field.SecondaryAllocID, error) {
	f := new(field.SecondaryAllocID)
	err := m.Body.Get(f)
	return f, err
}

//RefAllocID is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) RefAllocID() (*field.RefAllocID, error) {
	f := new(field.RefAllocID)
	err := m.Body.Get(f)
	return f, err
}

//AllocCancReplaceReason is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) AllocCancReplaceReason() (*field.AllocCancReplaceReason, error) {
	f := new(field.AllocCancReplaceReason)
	err := m.Body.Get(f)
	return f, err
}

//AllocIntermedReqType is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) AllocIntermedReqType() (*field.AllocIntermedReqType, error) {
	f := new(field.AllocIntermedReqType)
	err := m.Body.Get(f)
	return f, err
}

//AllocLinkID is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) AllocLinkID() (*field.AllocLinkID, error) {
	f := new(field.AllocLinkID)
	err := m.Body.Get(f)
	return f, err
}

//AllocLinkType is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) AllocLinkType() (*field.AllocLinkType, error) {
	f := new(field.AllocLinkType)
	err := m.Body.Get(f)
	return f, err
}

//BookingRefID is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) BookingRefID() (*field.BookingRefID, error) {
	f := new(field.BookingRefID)
	err := m.Body.Get(f)
	return f, err
}

//AllocNoOrdersType is a required field for AllocationInstruction.
func (m *AllocationInstruction) AllocNoOrdersType() (*field.AllocNoOrdersType, error) {
	f := new(field.AllocNoOrdersType)
	err := m.Body.Get(f)
	return f, err
}

//NoOrders is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) NoOrders() (*field.NoOrders, error) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}

//NoExecs is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) NoExecs() (*field.NoExecs, error) {
	f := new(field.NoExecs)
	err := m.Body.Get(f)
	return f, err
}

//PreviouslyReported is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) PreviouslyReported() (*field.PreviouslyReported, error) {
	f := new(field.PreviouslyReported)
	err := m.Body.Get(f)
	return f, err
}

//ReversalIndicator is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) ReversalIndicator() (*field.ReversalIndicator, error) {
	f := new(field.ReversalIndicator)
	err := m.Body.Get(f)
	return f, err
}

//MatchType is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) MatchType() (*field.MatchType, error) {
	f := new(field.MatchType)
	err := m.Body.Get(f)
	return f, err
}

//Side is a required field for AllocationInstruction.
func (m *AllocationInstruction) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryForm is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) DeliveryForm() (*field.DeliveryForm, error) {
	f := new(field.DeliveryForm)
	err := m.Body.Get(f)
	return f, err
}

//PctAtRisk is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) PctAtRisk() (*field.PctAtRisk, error) {
	f := new(field.PctAtRisk)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrAttrib is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) NoInstrAttrib() (*field.NoInstrAttrib, error) {
	f := new(field.NoInstrAttrib)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDesc is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) AgreementDesc() (*field.AgreementDesc, error) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//AgreementID is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) AgreementID() (*field.AgreementID, error) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDate is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) AgreementDate() (*field.AgreementDate, error) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//AgreementCurrency is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) AgreementCurrency() (*field.AgreementCurrency, error) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//TerminationType is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) TerminationType() (*field.TerminationType, error) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//StartDate is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) StartDate() (*field.StartDate, error) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//EndDate is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) EndDate() (*field.EndDate, error) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryType is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) DeliveryType() (*field.DeliveryType, error) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//MarginRatio is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) MarginRatio() (*field.MarginRatio, error) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//Quantity is a required field for AllocationInstruction.
func (m *AllocationInstruction) Quantity() (*field.Quantity, error) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}

//QtyType is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//LastMkt is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//TradeOriginationDate is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) TradeOriginationDate() (*field.TradeOriginationDate, error) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//PriceType is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//AvgPx is a required field for AllocationInstruction.
func (m *AllocationInstruction) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}

//AvgParPx is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) AvgParPx() (*field.AvgParPx, error) {
	f := new(field.AvgParPx)
	err := m.Body.Get(f)
	return f, err
}

//Spread is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) Spread() (*field.Spread, error) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, error) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveName is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) BenchmarkCurveName() (*field.BenchmarkCurveName, error) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, error) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPrice is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) BenchmarkPrice() (*field.BenchmarkPrice, error) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPriceType is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) BenchmarkPriceType() (*field.BenchmarkPriceType, error) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) BenchmarkSecurityID() (*field.BenchmarkSecurityID, error) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, error) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//AvgPxPrecision is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) AvgPxPrecision() (*field.AvgPxPrecision, error) {
	f := new(field.AvgPxPrecision)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a required field for AllocationInstruction.
func (m *AllocationInstruction) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//SettlType is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//SettlDate is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//BookingType is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) BookingType() (*field.BookingType, error) {
	f := new(field.BookingType)
	err := m.Body.Get(f)
	return f, err
}

//GrossTradeAmt is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) GrossTradeAmt() (*field.GrossTradeAmt, error) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}

//Concession is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) Concession() (*field.Concession, error) {
	f := new(field.Concession)
	err := m.Body.Get(f)
	return f, err
}

//TotalTakedown is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) TotalTakedown() (*field.TotalTakedown, error) {
	f := new(field.TotalTakedown)
	err := m.Body.Get(f)
	return f, err
}

//NetMoney is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) NetMoney() (*field.NetMoney, error) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}

//PositionEffect is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) PositionEffect() (*field.PositionEffect, error) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}

//AutoAcceptIndicator is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) AutoAcceptIndicator() (*field.AutoAcceptIndicator, error) {
	f := new(field.AutoAcceptIndicator)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//NumDaysInterest is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) NumDaysInterest() (*field.NumDaysInterest, error) {
	f := new(field.NumDaysInterest)
	err := m.Body.Get(f)
	return f, err
}

//AccruedInterestRate is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) AccruedInterestRate() (*field.AccruedInterestRate, error) {
	f := new(field.AccruedInterestRate)
	err := m.Body.Get(f)
	return f, err
}

//AccruedInterestAmt is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) AccruedInterestAmt() (*field.AccruedInterestAmt, error) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//TotalAccruedInterestAmt is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) TotalAccruedInterestAmt() (*field.TotalAccruedInterestAmt, error) {
	f := new(field.TotalAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//InterestAtMaturity is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) InterestAtMaturity() (*field.InterestAtMaturity, error) {
	f := new(field.InterestAtMaturity)
	err := m.Body.Get(f)
	return f, err
}

//EndAccruedInterestAmt is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) EndAccruedInterestAmt() (*field.EndAccruedInterestAmt, error) {
	f := new(field.EndAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//StartCash is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) StartCash() (*field.StartCash, error) {
	f := new(field.StartCash)
	err := m.Body.Get(f)
	return f, err
}

//EndCash is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) EndCash() (*field.EndCash, error) {
	f := new(field.EndCash)
	err := m.Body.Get(f)
	return f, err
}

//LegalConfirm is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) LegalConfirm() (*field.LegalConfirm, error) {
	f := new(field.LegalConfirm)
	err := m.Body.Get(f)
	return f, err
}

//NoStipulations is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) NoStipulations() (*field.NoStipulations, error) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//YieldType is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) YieldType() (*field.YieldType, error) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//Yield is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) Yield() (*field.Yield, error) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//YieldCalcDate is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) YieldCalcDate() (*field.YieldCalcDate, error) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionDate is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) YieldRedemptionDate() (*field.YieldRedemptionDate, error) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) YieldRedemptionPrice() (*field.YieldRedemptionPrice, error) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, error) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//TotNoAllocs is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) TotNoAllocs() (*field.TotNoAllocs, error) {
	f := new(field.TotNoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//LastFragment is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//NoAllocs is a non-required field for AllocationInstruction.
func (m *AllocationInstruction) NoAllocs() (*field.NoAllocs, error) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}
