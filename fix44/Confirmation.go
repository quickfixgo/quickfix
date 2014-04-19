package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Confirmation msg type = AK.
type Confirmation struct {
	message.Message
}

//ConfirmationBuilder builds Confirmation messages.
type ConfirmationBuilder struct {
	message.MessageBuilder
}

//CreateConfirmationBuilder returns an initialized ConfirmationBuilder with specified required fields.
func CreateConfirmationBuilder(
	confirmid field.ConfirmID,
	confirmtranstype field.ConfirmTransType,
	confirmtype field.ConfirmType,
	confirmstatus field.ConfirmStatus,
	transacttime field.TransactTime,
	tradedate field.TradeDate,
	nounderlyings field.NoUnderlyings,
	nolegs field.NoLegs,
	allocqty field.AllocQty,
	side field.Side,
	nocapacities field.NoCapacities,
	allocaccount field.AllocAccount,
	avgpx field.AvgPx,
	grosstradeamt field.GrossTradeAmt,
	netmoney field.NetMoney) ConfirmationBuilder {
	var builder ConfirmationBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(confirmid)
	builder.Body.Set(confirmtranstype)
	builder.Body.Set(confirmtype)
	builder.Body.Set(confirmstatus)
	builder.Body.Set(transacttime)
	builder.Body.Set(tradedate)
	builder.Body.Set(nounderlyings)
	builder.Body.Set(nolegs)
	builder.Body.Set(allocqty)
	builder.Body.Set(side)
	builder.Body.Set(nocapacities)
	builder.Body.Set(allocaccount)
	builder.Body.Set(avgpx)
	builder.Body.Set(grosstradeamt)
	builder.Body.Set(netmoney)
	return builder
}

//ConfirmID is a required field for Confirmation.
func (m Confirmation) ConfirmID() (field.ConfirmID, errors.MessageRejectError) {
	var f field.ConfirmID
	err := m.Body.Get(&f)
	return f, err
}

//ConfirmRefID is a non-required field for Confirmation.
func (m Confirmation) ConfirmRefID() (field.ConfirmRefID, errors.MessageRejectError) {
	var f field.ConfirmRefID
	err := m.Body.Get(&f)
	return f, err
}

//ConfirmReqID is a non-required field for Confirmation.
func (m Confirmation) ConfirmReqID() (field.ConfirmReqID, errors.MessageRejectError) {
	var f field.ConfirmReqID
	err := m.Body.Get(&f)
	return f, err
}

//ConfirmTransType is a required field for Confirmation.
func (m Confirmation) ConfirmTransType() (field.ConfirmTransType, errors.MessageRejectError) {
	var f field.ConfirmTransType
	err := m.Body.Get(&f)
	return f, err
}

//ConfirmType is a required field for Confirmation.
func (m Confirmation) ConfirmType() (field.ConfirmType, errors.MessageRejectError) {
	var f field.ConfirmType
	err := m.Body.Get(&f)
	return f, err
}

//CopyMsgIndicator is a non-required field for Confirmation.
func (m Confirmation) CopyMsgIndicator() (field.CopyMsgIndicator, errors.MessageRejectError) {
	var f field.CopyMsgIndicator
	err := m.Body.Get(&f)
	return f, err
}

//LegalConfirm is a non-required field for Confirmation.
func (m Confirmation) LegalConfirm() (field.LegalConfirm, errors.MessageRejectError) {
	var f field.LegalConfirm
	err := m.Body.Get(&f)
	return f, err
}

//ConfirmStatus is a required field for Confirmation.
func (m Confirmation) ConfirmStatus() (field.ConfirmStatus, errors.MessageRejectError) {
	var f field.ConfirmStatus
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for Confirmation.
func (m Confirmation) NoPartyIDs() (field.NoPartyIDs, errors.MessageRejectError) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//NoOrders is a non-required field for Confirmation.
func (m Confirmation) NoOrders() (field.NoOrders, errors.MessageRejectError) {
	var f field.NoOrders
	err := m.Body.Get(&f)
	return f, err
}

//AllocID is a non-required field for Confirmation.
func (m Confirmation) AllocID() (field.AllocID, errors.MessageRejectError) {
	var f field.AllocID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryAllocID is a non-required field for Confirmation.
func (m Confirmation) SecondaryAllocID() (field.SecondaryAllocID, errors.MessageRejectError) {
	var f field.SecondaryAllocID
	err := m.Body.Get(&f)
	return f, err
}

//IndividualAllocID is a non-required field for Confirmation.
func (m Confirmation) IndividualAllocID() (field.IndividualAllocID, errors.MessageRejectError) {
	var f field.IndividualAllocID
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for Confirmation.
func (m Confirmation) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a required field for Confirmation.
func (m Confirmation) TradeDate() (field.TradeDate, errors.MessageRejectError) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for Confirmation.
func (m Confirmation) NoTrdRegTimestamps() (field.NoTrdRegTimestamps, errors.MessageRejectError) {
	var f field.NoTrdRegTimestamps
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for Confirmation.
func (m Confirmation) Symbol() (field.Symbol, errors.MessageRejectError) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for Confirmation.
func (m Confirmation) SymbolSfx() (field.SymbolSfx, errors.MessageRejectError) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for Confirmation.
func (m Confirmation) SecurityID() (field.SecurityID, errors.MessageRejectError) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for Confirmation.
func (m Confirmation) SecurityIDSource() (field.SecurityIDSource, errors.MessageRejectError) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for Confirmation.
func (m Confirmation) NoSecurityAltID() (field.NoSecurityAltID, errors.MessageRejectError) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for Confirmation.
func (m Confirmation) Product() (field.Product, errors.MessageRejectError) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for Confirmation.
func (m Confirmation) CFICode() (field.CFICode, errors.MessageRejectError) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for Confirmation.
func (m Confirmation) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for Confirmation.
func (m Confirmation) SecuritySubType() (field.SecuritySubType, errors.MessageRejectError) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for Confirmation.
func (m Confirmation) MaturityMonthYear() (field.MaturityMonthYear, errors.MessageRejectError) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for Confirmation.
func (m Confirmation) MaturityDate() (field.MaturityDate, errors.MessageRejectError) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for Confirmation.
func (m Confirmation) CouponPaymentDate() (field.CouponPaymentDate, errors.MessageRejectError) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for Confirmation.
func (m Confirmation) IssueDate() (field.IssueDate, errors.MessageRejectError) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for Confirmation.
func (m Confirmation) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, errors.MessageRejectError) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for Confirmation.
func (m Confirmation) RepurchaseTerm() (field.RepurchaseTerm, errors.MessageRejectError) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for Confirmation.
func (m Confirmation) RepurchaseRate() (field.RepurchaseRate, errors.MessageRejectError) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for Confirmation.
func (m Confirmation) Factor() (field.Factor, errors.MessageRejectError) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for Confirmation.
func (m Confirmation) CreditRating() (field.CreditRating, errors.MessageRejectError) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for Confirmation.
func (m Confirmation) InstrRegistry() (field.InstrRegistry, errors.MessageRejectError) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for Confirmation.
func (m Confirmation) CountryOfIssue() (field.CountryOfIssue, errors.MessageRejectError) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for Confirmation.
func (m Confirmation) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for Confirmation.
func (m Confirmation) LocaleOfIssue() (field.LocaleOfIssue, errors.MessageRejectError) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for Confirmation.
func (m Confirmation) RedemptionDate() (field.RedemptionDate, errors.MessageRejectError) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for Confirmation.
func (m Confirmation) StrikePrice() (field.StrikePrice, errors.MessageRejectError) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for Confirmation.
func (m Confirmation) StrikeCurrency() (field.StrikeCurrency, errors.MessageRejectError) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for Confirmation.
func (m Confirmation) OptAttribute() (field.OptAttribute, errors.MessageRejectError) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for Confirmation.
func (m Confirmation) ContractMultiplier() (field.ContractMultiplier, errors.MessageRejectError) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for Confirmation.
func (m Confirmation) CouponRate() (field.CouponRate, errors.MessageRejectError) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for Confirmation.
func (m Confirmation) SecurityExchange() (field.SecurityExchange, errors.MessageRejectError) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for Confirmation.
func (m Confirmation) Issuer() (field.Issuer, errors.MessageRejectError) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for Confirmation.
func (m Confirmation) EncodedIssuerLen() (field.EncodedIssuerLen, errors.MessageRejectError) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for Confirmation.
func (m Confirmation) EncodedIssuer() (field.EncodedIssuer, errors.MessageRejectError) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for Confirmation.
func (m Confirmation) SecurityDesc() (field.SecurityDesc, errors.MessageRejectError) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for Confirmation.
func (m Confirmation) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, errors.MessageRejectError) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for Confirmation.
func (m Confirmation) EncodedSecurityDesc() (field.EncodedSecurityDesc, errors.MessageRejectError) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for Confirmation.
func (m Confirmation) Pool() (field.Pool, errors.MessageRejectError) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for Confirmation.
func (m Confirmation) ContractSettlMonth() (field.ContractSettlMonth, errors.MessageRejectError) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for Confirmation.
func (m Confirmation) CPProgram() (field.CPProgram, errors.MessageRejectError) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for Confirmation.
func (m Confirmation) CPRegType() (field.CPRegType, errors.MessageRejectError) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for Confirmation.
func (m Confirmation) NoEvents() (field.NoEvents, errors.MessageRejectError) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for Confirmation.
func (m Confirmation) DatedDate() (field.DatedDate, errors.MessageRejectError) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for Confirmation.
func (m Confirmation) InterestAccrualDate() (field.InterestAccrualDate, errors.MessageRejectError) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryForm is a non-required field for Confirmation.
func (m Confirmation) DeliveryForm() (field.DeliveryForm, errors.MessageRejectError) {
	var f field.DeliveryForm
	err := m.Body.Get(&f)
	return f, err
}

//PctAtRisk is a non-required field for Confirmation.
func (m Confirmation) PctAtRisk() (field.PctAtRisk, errors.MessageRejectError) {
	var f field.PctAtRisk
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrAttrib is a non-required field for Confirmation.
func (m Confirmation) NoInstrAttrib() (field.NoInstrAttrib, errors.MessageRejectError) {
	var f field.NoInstrAttrib
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for Confirmation.
func (m Confirmation) AgreementDesc() (field.AgreementDesc, errors.MessageRejectError) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for Confirmation.
func (m Confirmation) AgreementID() (field.AgreementID, errors.MessageRejectError) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for Confirmation.
func (m Confirmation) AgreementDate() (field.AgreementDate, errors.MessageRejectError) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for Confirmation.
func (m Confirmation) AgreementCurrency() (field.AgreementCurrency, errors.MessageRejectError) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for Confirmation.
func (m Confirmation) TerminationType() (field.TerminationType, errors.MessageRejectError) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for Confirmation.
func (m Confirmation) StartDate() (field.StartDate, errors.MessageRejectError) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for Confirmation.
func (m Confirmation) EndDate() (field.EndDate, errors.MessageRejectError) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for Confirmation.
func (m Confirmation) DeliveryType() (field.DeliveryType, errors.MessageRejectError) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for Confirmation.
func (m Confirmation) MarginRatio() (field.MarginRatio, errors.MessageRejectError) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a required field for Confirmation.
func (m Confirmation) NoUnderlyings() (field.NoUnderlyings, errors.MessageRejectError) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a required field for Confirmation.
func (m Confirmation) NoLegs() (field.NoLegs, errors.MessageRejectError) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for Confirmation.
func (m Confirmation) YieldType() (field.YieldType, errors.MessageRejectError) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for Confirmation.
func (m Confirmation) Yield() (field.Yield, errors.MessageRejectError) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//YieldCalcDate is a non-required field for Confirmation.
func (m Confirmation) YieldCalcDate() (field.YieldCalcDate, errors.MessageRejectError) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for Confirmation.
func (m Confirmation) YieldRedemptionDate() (field.YieldRedemptionDate, errors.MessageRejectError) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for Confirmation.
func (m Confirmation) YieldRedemptionPrice() (field.YieldRedemptionPrice, errors.MessageRejectError) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for Confirmation.
func (m Confirmation) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, errors.MessageRejectError) {
	var f field.YieldRedemptionPriceType
	err := m.Body.Get(&f)
	return f, err
}

//AllocQty is a required field for Confirmation.
func (m Confirmation) AllocQty() (field.AllocQty, errors.MessageRejectError) {
	var f field.AllocQty
	err := m.Body.Get(&f)
	return f, err
}

//QtyType is a non-required field for Confirmation.
func (m Confirmation) QtyType() (field.QtyType, errors.MessageRejectError) {
	var f field.QtyType
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for Confirmation.
func (m Confirmation) Side() (field.Side, errors.MessageRejectError) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for Confirmation.
func (m Confirmation) Currency() (field.Currency, errors.MessageRejectError) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//LastMkt is a non-required field for Confirmation.
func (m Confirmation) LastMkt() (field.LastMkt, errors.MessageRejectError) {
	var f field.LastMkt
	err := m.Body.Get(&f)
	return f, err
}

//NoCapacities is a required field for Confirmation.
func (m Confirmation) NoCapacities() (field.NoCapacities, errors.MessageRejectError) {
	var f field.NoCapacities
	err := m.Body.Get(&f)
	return f, err
}

//AllocAccount is a required field for Confirmation.
func (m Confirmation) AllocAccount() (field.AllocAccount, errors.MessageRejectError) {
	var f field.AllocAccount
	err := m.Body.Get(&f)
	return f, err
}

//AllocAcctIDSource is a non-required field for Confirmation.
func (m Confirmation) AllocAcctIDSource() (field.AllocAcctIDSource, errors.MessageRejectError) {
	var f field.AllocAcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AllocAccountType is a non-required field for Confirmation.
func (m Confirmation) AllocAccountType() (field.AllocAccountType, errors.MessageRejectError) {
	var f field.AllocAccountType
	err := m.Body.Get(&f)
	return f, err
}

//AvgPx is a required field for Confirmation.
func (m Confirmation) AvgPx() (field.AvgPx, errors.MessageRejectError) {
	var f field.AvgPx
	err := m.Body.Get(&f)
	return f, err
}

//AvgPxPrecision is a non-required field for Confirmation.
func (m Confirmation) AvgPxPrecision() (field.AvgPxPrecision, errors.MessageRejectError) {
	var f field.AvgPxPrecision
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for Confirmation.
func (m Confirmation) PriceType() (field.PriceType, errors.MessageRejectError) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//AvgParPx is a non-required field for Confirmation.
func (m Confirmation) AvgParPx() (field.AvgParPx, errors.MessageRejectError) {
	var f field.AvgParPx
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for Confirmation.
func (m Confirmation) Spread() (field.Spread, errors.MessageRejectError) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for Confirmation.
func (m Confirmation) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for Confirmation.
func (m Confirmation) BenchmarkCurveName() (field.BenchmarkCurveName, errors.MessageRejectError) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for Confirmation.
func (m Confirmation) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, errors.MessageRejectError) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for Confirmation.
func (m Confirmation) BenchmarkPrice() (field.BenchmarkPrice, errors.MessageRejectError) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for Confirmation.
func (m Confirmation) BenchmarkPriceType() (field.BenchmarkPriceType, errors.MessageRejectError) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for Confirmation.
func (m Confirmation) BenchmarkSecurityID() (field.BenchmarkSecurityID, errors.MessageRejectError) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for Confirmation.
func (m Confirmation) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//ReportedPx is a non-required field for Confirmation.
func (m Confirmation) ReportedPx() (field.ReportedPx, errors.MessageRejectError) {
	var f field.ReportedPx
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for Confirmation.
func (m Confirmation) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for Confirmation.
func (m Confirmation) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for Confirmation.
func (m Confirmation) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//ProcessCode is a non-required field for Confirmation.
func (m Confirmation) ProcessCode() (field.ProcessCode, errors.MessageRejectError) {
	var f field.ProcessCode
	err := m.Body.Get(&f)
	return f, err
}

//GrossTradeAmt is a required field for Confirmation.
func (m Confirmation) GrossTradeAmt() (field.GrossTradeAmt, errors.MessageRejectError) {
	var f field.GrossTradeAmt
	err := m.Body.Get(&f)
	return f, err
}

//NumDaysInterest is a non-required field for Confirmation.
func (m Confirmation) NumDaysInterest() (field.NumDaysInterest, errors.MessageRejectError) {
	var f field.NumDaysInterest
	err := m.Body.Get(&f)
	return f, err
}

//ExDate is a non-required field for Confirmation.
func (m Confirmation) ExDate() (field.ExDate, errors.MessageRejectError) {
	var f field.ExDate
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestRate is a non-required field for Confirmation.
func (m Confirmation) AccruedInterestRate() (field.AccruedInterestRate, errors.MessageRejectError) {
	var f field.AccruedInterestRate
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestAmt is a non-required field for Confirmation.
func (m Confirmation) AccruedInterestAmt() (field.AccruedInterestAmt, errors.MessageRejectError) {
	var f field.AccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//InterestAtMaturity is a non-required field for Confirmation.
func (m Confirmation) InterestAtMaturity() (field.InterestAtMaturity, errors.MessageRejectError) {
	var f field.InterestAtMaturity
	err := m.Body.Get(&f)
	return f, err
}

//EndAccruedInterestAmt is a non-required field for Confirmation.
func (m Confirmation) EndAccruedInterestAmt() (field.EndAccruedInterestAmt, errors.MessageRejectError) {
	var f field.EndAccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//StartCash is a non-required field for Confirmation.
func (m Confirmation) StartCash() (field.StartCash, errors.MessageRejectError) {
	var f field.StartCash
	err := m.Body.Get(&f)
	return f, err
}

//EndCash is a non-required field for Confirmation.
func (m Confirmation) EndCash() (field.EndCash, errors.MessageRejectError) {
	var f field.EndCash
	err := m.Body.Get(&f)
	return f, err
}

//Concession is a non-required field for Confirmation.
func (m Confirmation) Concession() (field.Concession, errors.MessageRejectError) {
	var f field.Concession
	err := m.Body.Get(&f)
	return f, err
}

//TotalTakedown is a non-required field for Confirmation.
func (m Confirmation) TotalTakedown() (field.TotalTakedown, errors.MessageRejectError) {
	var f field.TotalTakedown
	err := m.Body.Get(&f)
	return f, err
}

//NetMoney is a required field for Confirmation.
func (m Confirmation) NetMoney() (field.NetMoney, errors.MessageRejectError) {
	var f field.NetMoney
	err := m.Body.Get(&f)
	return f, err
}

//MaturityNetMoney is a non-required field for Confirmation.
func (m Confirmation) MaturityNetMoney() (field.MaturityNetMoney, errors.MessageRejectError) {
	var f field.MaturityNetMoney
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrAmt is a non-required field for Confirmation.
func (m Confirmation) SettlCurrAmt() (field.SettlCurrAmt, errors.MessageRejectError) {
	var f field.SettlCurrAmt
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for Confirmation.
func (m Confirmation) SettlCurrency() (field.SettlCurrency, errors.MessageRejectError) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrFxRate is a non-required field for Confirmation.
func (m Confirmation) SettlCurrFxRate() (field.SettlCurrFxRate, errors.MessageRejectError) {
	var f field.SettlCurrFxRate
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrFxRateCalc is a non-required field for Confirmation.
func (m Confirmation) SettlCurrFxRateCalc() (field.SettlCurrFxRateCalc, errors.MessageRejectError) {
	var f field.SettlCurrFxRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for Confirmation.
func (m Confirmation) SettlType() (field.SettlType, errors.MessageRejectError) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for Confirmation.
func (m Confirmation) SettlDate() (field.SettlDate, errors.MessageRejectError) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//SettlDeliveryType is a non-required field for Confirmation.
func (m Confirmation) SettlDeliveryType() (field.SettlDeliveryType, errors.MessageRejectError) {
	var f field.SettlDeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbType is a non-required field for Confirmation.
func (m Confirmation) StandInstDbType() (field.StandInstDbType, errors.MessageRejectError) {
	var f field.StandInstDbType
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbName is a non-required field for Confirmation.
func (m Confirmation) StandInstDbName() (field.StandInstDbName, errors.MessageRejectError) {
	var f field.StandInstDbName
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbID is a non-required field for Confirmation.
func (m Confirmation) StandInstDbID() (field.StandInstDbID, errors.MessageRejectError) {
	var f field.StandInstDbID
	err := m.Body.Get(&f)
	return f, err
}

//NoDlvyInst is a non-required field for Confirmation.
func (m Confirmation) NoDlvyInst() (field.NoDlvyInst, errors.MessageRejectError) {
	var f field.NoDlvyInst
	err := m.Body.Get(&f)
	return f, err
}

//Commission is a non-required field for Confirmation.
func (m Confirmation) Commission() (field.Commission, errors.MessageRejectError) {
	var f field.Commission
	err := m.Body.Get(&f)
	return f, err
}

//CommType is a non-required field for Confirmation.
func (m Confirmation) CommType() (field.CommType, errors.MessageRejectError) {
	var f field.CommType
	err := m.Body.Get(&f)
	return f, err
}

//CommCurrency is a non-required field for Confirmation.
func (m Confirmation) CommCurrency() (field.CommCurrency, errors.MessageRejectError) {
	var f field.CommCurrency
	err := m.Body.Get(&f)
	return f, err
}

//FundRenewWaiv is a non-required field for Confirmation.
func (m Confirmation) FundRenewWaiv() (field.FundRenewWaiv, errors.MessageRejectError) {
	var f field.FundRenewWaiv
	err := m.Body.Get(&f)
	return f, err
}

//SharedCommission is a non-required field for Confirmation.
func (m Confirmation) SharedCommission() (field.SharedCommission, errors.MessageRejectError) {
	var f field.SharedCommission
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for Confirmation.
func (m Confirmation) NoStipulations() (field.NoStipulations, errors.MessageRejectError) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//NoMiscFees is a non-required field for Confirmation.
func (m Confirmation) NoMiscFees() (field.NoMiscFees, errors.MessageRejectError) {
	var f field.NoMiscFees
	err := m.Body.Get(&f)
	return f, err
}
