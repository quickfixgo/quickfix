package fix50sp2

import (
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
func (m Confirmation) ConfirmID() (field.ConfirmID, error) {
	var f field.ConfirmID
	err := m.Body.Get(&f)
	return f, err
}

//ConfirmRefID is a non-required field for Confirmation.
func (m Confirmation) ConfirmRefID() (field.ConfirmRefID, error) {
	var f field.ConfirmRefID
	err := m.Body.Get(&f)
	return f, err
}

//ConfirmReqID is a non-required field for Confirmation.
func (m Confirmation) ConfirmReqID() (field.ConfirmReqID, error) {
	var f field.ConfirmReqID
	err := m.Body.Get(&f)
	return f, err
}

//ConfirmTransType is a required field for Confirmation.
func (m Confirmation) ConfirmTransType() (field.ConfirmTransType, error) {
	var f field.ConfirmTransType
	err := m.Body.Get(&f)
	return f, err
}

//ConfirmType is a required field for Confirmation.
func (m Confirmation) ConfirmType() (field.ConfirmType, error) {
	var f field.ConfirmType
	err := m.Body.Get(&f)
	return f, err
}

//CopyMsgIndicator is a non-required field for Confirmation.
func (m Confirmation) CopyMsgIndicator() (field.CopyMsgIndicator, error) {
	var f field.CopyMsgIndicator
	err := m.Body.Get(&f)
	return f, err
}

//LegalConfirm is a non-required field for Confirmation.
func (m Confirmation) LegalConfirm() (field.LegalConfirm, error) {
	var f field.LegalConfirm
	err := m.Body.Get(&f)
	return f, err
}

//ConfirmStatus is a required field for Confirmation.
func (m Confirmation) ConfirmStatus() (field.ConfirmStatus, error) {
	var f field.ConfirmStatus
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for Confirmation.
func (m Confirmation) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//NoOrders is a non-required field for Confirmation.
func (m Confirmation) NoOrders() (field.NoOrders, error) {
	var f field.NoOrders
	err := m.Body.Get(&f)
	return f, err
}

//AllocID is a non-required field for Confirmation.
func (m Confirmation) AllocID() (field.AllocID, error) {
	var f field.AllocID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryAllocID is a non-required field for Confirmation.
func (m Confirmation) SecondaryAllocID() (field.SecondaryAllocID, error) {
	var f field.SecondaryAllocID
	err := m.Body.Get(&f)
	return f, err
}

//IndividualAllocID is a non-required field for Confirmation.
func (m Confirmation) IndividualAllocID() (field.IndividualAllocID, error) {
	var f field.IndividualAllocID
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for Confirmation.
func (m Confirmation) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a required field for Confirmation.
func (m Confirmation) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for Confirmation.
func (m Confirmation) NoTrdRegTimestamps() (field.NoTrdRegTimestamps, error) {
	var f field.NoTrdRegTimestamps
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for Confirmation.
func (m Confirmation) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for Confirmation.
func (m Confirmation) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for Confirmation.
func (m Confirmation) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for Confirmation.
func (m Confirmation) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for Confirmation.
func (m Confirmation) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for Confirmation.
func (m Confirmation) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for Confirmation.
func (m Confirmation) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for Confirmation.
func (m Confirmation) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for Confirmation.
func (m Confirmation) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for Confirmation.
func (m Confirmation) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for Confirmation.
func (m Confirmation) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for Confirmation.
func (m Confirmation) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for Confirmation.
func (m Confirmation) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for Confirmation.
func (m Confirmation) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for Confirmation.
func (m Confirmation) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for Confirmation.
func (m Confirmation) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for Confirmation.
func (m Confirmation) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for Confirmation.
func (m Confirmation) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for Confirmation.
func (m Confirmation) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for Confirmation.
func (m Confirmation) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for Confirmation.
func (m Confirmation) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for Confirmation.
func (m Confirmation) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for Confirmation.
func (m Confirmation) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for Confirmation.
func (m Confirmation) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for Confirmation.
func (m Confirmation) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for Confirmation.
func (m Confirmation) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for Confirmation.
func (m Confirmation) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for Confirmation.
func (m Confirmation) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for Confirmation.
func (m Confirmation) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for Confirmation.
func (m Confirmation) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for Confirmation.
func (m Confirmation) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for Confirmation.
func (m Confirmation) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for Confirmation.
func (m Confirmation) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for Confirmation.
func (m Confirmation) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for Confirmation.
func (m Confirmation) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for Confirmation.
func (m Confirmation) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for Confirmation.
func (m Confirmation) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for Confirmation.
func (m Confirmation) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for Confirmation.
func (m Confirmation) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for Confirmation.
func (m Confirmation) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for Confirmation.
func (m Confirmation) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for Confirmation.
func (m Confirmation) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for Confirmation.
func (m Confirmation) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for Confirmation.
func (m Confirmation) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for Confirmation.
func (m Confirmation) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for Confirmation.
func (m Confirmation) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for Confirmation.
func (m Confirmation) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for Confirmation.
func (m Confirmation) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for Confirmation.
func (m Confirmation) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for Confirmation.
func (m Confirmation) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for Confirmation.
func (m Confirmation) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for Confirmation.
func (m Confirmation) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for Confirmation.
func (m Confirmation) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for Confirmation.
func (m Confirmation) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//SecurityGroup is a non-required field for Confirmation.
func (m Confirmation) SecurityGroup() (field.SecurityGroup, error) {
	var f field.SecurityGroup
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for Confirmation.
func (m Confirmation) MinPriceIncrementAmount() (field.MinPriceIncrementAmount, error) {
	var f field.MinPriceIncrementAmount
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for Confirmation.
func (m Confirmation) UnitOfMeasureQty() (field.UnitOfMeasureQty, error) {
	var f field.UnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLLen is a non-required field for Confirmation.
func (m Confirmation) SecurityXMLLen() (field.SecurityXMLLen, error) {
	var f field.SecurityXMLLen
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXML is a non-required field for Confirmation.
func (m Confirmation) SecurityXML() (field.SecurityXML, error) {
	var f field.SecurityXML
	err := m.Body.Get(&f)
	return f, err
}

//SecurityXMLSchema is a non-required field for Confirmation.
func (m Confirmation) SecurityXMLSchema() (field.SecurityXMLSchema, error) {
	var f field.SecurityXMLSchema
	err := m.Body.Get(&f)
	return f, err
}

//ProductComplex is a non-required field for Confirmation.
func (m Confirmation) ProductComplex() (field.ProductComplex, error) {
	var f field.ProductComplex
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for Confirmation.
func (m Confirmation) PriceUnitOfMeasure() (field.PriceUnitOfMeasure, error) {
	var f field.PriceUnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for Confirmation.
func (m Confirmation) PriceUnitOfMeasureQty() (field.PriceUnitOfMeasureQty, error) {
	var f field.PriceUnitOfMeasureQty
	err := m.Body.Get(&f)
	return f, err
}

//SettlMethod is a non-required field for Confirmation.
func (m Confirmation) SettlMethod() (field.SettlMethod, error) {
	var f field.SettlMethod
	err := m.Body.Get(&f)
	return f, err
}

//ExerciseStyle is a non-required field for Confirmation.
func (m Confirmation) ExerciseStyle() (field.ExerciseStyle, error) {
	var f field.ExerciseStyle
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutAmount is a non-required field for Confirmation.
func (m Confirmation) OptPayoutAmount() (field.OptPayoutAmount, error) {
	var f field.OptPayoutAmount
	err := m.Body.Get(&f)
	return f, err
}

//PriceQuoteMethod is a non-required field for Confirmation.
func (m Confirmation) PriceQuoteMethod() (field.PriceQuoteMethod, error) {
	var f field.PriceQuoteMethod
	err := m.Body.Get(&f)
	return f, err
}

//ListMethod is a non-required field for Confirmation.
func (m Confirmation) ListMethod() (field.ListMethod, error) {
	var f field.ListMethod
	err := m.Body.Get(&f)
	return f, err
}

//CapPrice is a non-required field for Confirmation.
func (m Confirmation) CapPrice() (field.CapPrice, error) {
	var f field.CapPrice
	err := m.Body.Get(&f)
	return f, err
}

//FloorPrice is a non-required field for Confirmation.
func (m Confirmation) FloorPrice() (field.FloorPrice, error) {
	var f field.FloorPrice
	err := m.Body.Get(&f)
	return f, err
}

//PutOrCall is a non-required field for Confirmation.
func (m Confirmation) PutOrCall() (field.PutOrCall, error) {
	var f field.PutOrCall
	err := m.Body.Get(&f)
	return f, err
}

//FlexibleIndicator is a non-required field for Confirmation.
func (m Confirmation) FlexibleIndicator() (field.FlexibleIndicator, error) {
	var f field.FlexibleIndicator
	err := m.Body.Get(&f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for Confirmation.
func (m Confirmation) FlexProductEligibilityIndicator() (field.FlexProductEligibilityIndicator, error) {
	var f field.FlexProductEligibilityIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ValuationMethod is a non-required field for Confirmation.
func (m Confirmation) ValuationMethod() (field.ValuationMethod, error) {
	var f field.ValuationMethod
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for Confirmation.
func (m Confirmation) ContractMultiplierUnit() (field.ContractMultiplierUnit, error) {
	var f field.ContractMultiplierUnit
	err := m.Body.Get(&f)
	return f, err
}

//FlowScheduleType is a non-required field for Confirmation.
func (m Confirmation) FlowScheduleType() (field.FlowScheduleType, error) {
	var f field.FlowScheduleType
	err := m.Body.Get(&f)
	return f, err
}

//RestructuringType is a non-required field for Confirmation.
func (m Confirmation) RestructuringType() (field.RestructuringType, error) {
	var f field.RestructuringType
	err := m.Body.Get(&f)
	return f, err
}

//Seniority is a non-required field for Confirmation.
func (m Confirmation) Seniority() (field.Seniority, error) {
	var f field.Seniority
	err := m.Body.Get(&f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for Confirmation.
func (m Confirmation) NotionalPercentageOutstanding() (field.NotionalPercentageOutstanding, error) {
	var f field.NotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for Confirmation.
func (m Confirmation) OriginalNotionalPercentageOutstanding() (field.OriginalNotionalPercentageOutstanding, error) {
	var f field.OriginalNotionalPercentageOutstanding
	err := m.Body.Get(&f)
	return f, err
}

//AttachmentPoint is a non-required field for Confirmation.
func (m Confirmation) AttachmentPoint() (field.AttachmentPoint, error) {
	var f field.AttachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//DetachmentPoint is a non-required field for Confirmation.
func (m Confirmation) DetachmentPoint() (field.DetachmentPoint, error) {
	var f field.DetachmentPoint
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for Confirmation.
func (m Confirmation) StrikePriceDeterminationMethod() (field.StrikePriceDeterminationMethod, error) {
	var f field.StrikePriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for Confirmation.
func (m Confirmation) StrikePriceBoundaryMethod() (field.StrikePriceBoundaryMethod, error) {
	var f field.StrikePriceBoundaryMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for Confirmation.
func (m Confirmation) StrikePriceBoundaryPrecision() (field.StrikePriceBoundaryPrecision, error) {
	var f field.StrikePriceBoundaryPrecision
	err := m.Body.Get(&f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for Confirmation.
func (m Confirmation) UnderlyingPriceDeterminationMethod() (field.UnderlyingPriceDeterminationMethod, error) {
	var f field.UnderlyingPriceDeterminationMethod
	err := m.Body.Get(&f)
	return f, err
}

//OptPayoutType is a non-required field for Confirmation.
func (m Confirmation) OptPayoutType() (field.OptPayoutType, error) {
	var f field.OptPayoutType
	err := m.Body.Get(&f)
	return f, err
}

//NoComplexEvents is a non-required field for Confirmation.
func (m Confirmation) NoComplexEvents() (field.NoComplexEvents, error) {
	var f field.NoComplexEvents
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryForm is a non-required field for Confirmation.
func (m Confirmation) DeliveryForm() (field.DeliveryForm, error) {
	var f field.DeliveryForm
	err := m.Body.Get(&f)
	return f, err
}

//PctAtRisk is a non-required field for Confirmation.
func (m Confirmation) PctAtRisk() (field.PctAtRisk, error) {
	var f field.PctAtRisk
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrAttrib is a non-required field for Confirmation.
func (m Confirmation) NoInstrAttrib() (field.NoInstrAttrib, error) {
	var f field.NoInstrAttrib
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for Confirmation.
func (m Confirmation) AgreementDesc() (field.AgreementDesc, error) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for Confirmation.
func (m Confirmation) AgreementID() (field.AgreementID, error) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for Confirmation.
func (m Confirmation) AgreementDate() (field.AgreementDate, error) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for Confirmation.
func (m Confirmation) AgreementCurrency() (field.AgreementCurrency, error) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for Confirmation.
func (m Confirmation) TerminationType() (field.TerminationType, error) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for Confirmation.
func (m Confirmation) StartDate() (field.StartDate, error) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for Confirmation.
func (m Confirmation) EndDate() (field.EndDate, error) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for Confirmation.
func (m Confirmation) DeliveryType() (field.DeliveryType, error) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for Confirmation.
func (m Confirmation) MarginRatio() (field.MarginRatio, error) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for Confirmation.
func (m Confirmation) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for Confirmation.
func (m Confirmation) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//YieldType is a non-required field for Confirmation.
func (m Confirmation) YieldType() (field.YieldType, error) {
	var f field.YieldType
	err := m.Body.Get(&f)
	return f, err
}

//Yield is a non-required field for Confirmation.
func (m Confirmation) Yield() (field.Yield, error) {
	var f field.Yield
	err := m.Body.Get(&f)
	return f, err
}

//YieldCalcDate is a non-required field for Confirmation.
func (m Confirmation) YieldCalcDate() (field.YieldCalcDate, error) {
	var f field.YieldCalcDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionDate is a non-required field for Confirmation.
func (m Confirmation) YieldRedemptionDate() (field.YieldRedemptionDate, error) {
	var f field.YieldRedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for Confirmation.
func (m Confirmation) YieldRedemptionPrice() (field.YieldRedemptionPrice, error) {
	var f field.YieldRedemptionPrice
	err := m.Body.Get(&f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for Confirmation.
func (m Confirmation) YieldRedemptionPriceType() (field.YieldRedemptionPriceType, error) {
	var f field.YieldRedemptionPriceType
	err := m.Body.Get(&f)
	return f, err
}

//AllocQty is a required field for Confirmation.
func (m Confirmation) AllocQty() (field.AllocQty, error) {
	var f field.AllocQty
	err := m.Body.Get(&f)
	return f, err
}

//QtyType is a non-required field for Confirmation.
func (m Confirmation) QtyType() (field.QtyType, error) {
	var f field.QtyType
	err := m.Body.Get(&f)
	return f, err
}

//Side is a required field for Confirmation.
func (m Confirmation) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for Confirmation.
func (m Confirmation) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//LastMkt is a non-required field for Confirmation.
func (m Confirmation) LastMkt() (field.LastMkt, error) {
	var f field.LastMkt
	err := m.Body.Get(&f)
	return f, err
}

//NoCapacities is a required field for Confirmation.
func (m Confirmation) NoCapacities() (field.NoCapacities, error) {
	var f field.NoCapacities
	err := m.Body.Get(&f)
	return f, err
}

//AllocAccount is a required field for Confirmation.
func (m Confirmation) AllocAccount() (field.AllocAccount, error) {
	var f field.AllocAccount
	err := m.Body.Get(&f)
	return f, err
}

//AllocAcctIDSource is a non-required field for Confirmation.
func (m Confirmation) AllocAcctIDSource() (field.AllocAcctIDSource, error) {
	var f field.AllocAcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AllocAccountType is a non-required field for Confirmation.
func (m Confirmation) AllocAccountType() (field.AllocAccountType, error) {
	var f field.AllocAccountType
	err := m.Body.Get(&f)
	return f, err
}

//AvgPx is a required field for Confirmation.
func (m Confirmation) AvgPx() (field.AvgPx, error) {
	var f field.AvgPx
	err := m.Body.Get(&f)
	return f, err
}

//AvgPxPrecision is a non-required field for Confirmation.
func (m Confirmation) AvgPxPrecision() (field.AvgPxPrecision, error) {
	var f field.AvgPxPrecision
	err := m.Body.Get(&f)
	return f, err
}

//PriceType is a non-required field for Confirmation.
func (m Confirmation) PriceType() (field.PriceType, error) {
	var f field.PriceType
	err := m.Body.Get(&f)
	return f, err
}

//AvgParPx is a non-required field for Confirmation.
func (m Confirmation) AvgParPx() (field.AvgParPx, error) {
	var f field.AvgParPx
	err := m.Body.Get(&f)
	return f, err
}

//Spread is a non-required field for Confirmation.
func (m Confirmation) Spread() (field.Spread, error) {
	var f field.Spread
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for Confirmation.
func (m Confirmation) BenchmarkCurveCurrency() (field.BenchmarkCurveCurrency, error) {
	var f field.BenchmarkCurveCurrency
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurveName is a non-required field for Confirmation.
func (m Confirmation) BenchmarkCurveName() (field.BenchmarkCurveName, error) {
	var f field.BenchmarkCurveName
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for Confirmation.
func (m Confirmation) BenchmarkCurvePoint() (field.BenchmarkCurvePoint, error) {
	var f field.BenchmarkCurvePoint
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPrice is a non-required field for Confirmation.
func (m Confirmation) BenchmarkPrice() (field.BenchmarkPrice, error) {
	var f field.BenchmarkPrice
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkPriceType is a non-required field for Confirmation.
func (m Confirmation) BenchmarkPriceType() (field.BenchmarkPriceType, error) {
	var f field.BenchmarkPriceType
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for Confirmation.
func (m Confirmation) BenchmarkSecurityID() (field.BenchmarkSecurityID, error) {
	var f field.BenchmarkSecurityID
	err := m.Body.Get(&f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for Confirmation.
func (m Confirmation) BenchmarkSecurityIDSource() (field.BenchmarkSecurityIDSource, error) {
	var f field.BenchmarkSecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//ReportedPx is a non-required field for Confirmation.
func (m Confirmation) ReportedPx() (field.ReportedPx, error) {
	var f field.ReportedPx
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for Confirmation.
func (m Confirmation) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for Confirmation.
func (m Confirmation) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for Confirmation.
func (m Confirmation) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//ProcessCode is a non-required field for Confirmation.
func (m Confirmation) ProcessCode() (field.ProcessCode, error) {
	var f field.ProcessCode
	err := m.Body.Get(&f)
	return f, err
}

//GrossTradeAmt is a required field for Confirmation.
func (m Confirmation) GrossTradeAmt() (field.GrossTradeAmt, error) {
	var f field.GrossTradeAmt
	err := m.Body.Get(&f)
	return f, err
}

//NumDaysInterest is a non-required field for Confirmation.
func (m Confirmation) NumDaysInterest() (field.NumDaysInterest, error) {
	var f field.NumDaysInterest
	err := m.Body.Get(&f)
	return f, err
}

//ExDate is a non-required field for Confirmation.
func (m Confirmation) ExDate() (field.ExDate, error) {
	var f field.ExDate
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestRate is a non-required field for Confirmation.
func (m Confirmation) AccruedInterestRate() (field.AccruedInterestRate, error) {
	var f field.AccruedInterestRate
	err := m.Body.Get(&f)
	return f, err
}

//AccruedInterestAmt is a non-required field for Confirmation.
func (m Confirmation) AccruedInterestAmt() (field.AccruedInterestAmt, error) {
	var f field.AccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//InterestAtMaturity is a non-required field for Confirmation.
func (m Confirmation) InterestAtMaturity() (field.InterestAtMaturity, error) {
	var f field.InterestAtMaturity
	err := m.Body.Get(&f)
	return f, err
}

//EndAccruedInterestAmt is a non-required field for Confirmation.
func (m Confirmation) EndAccruedInterestAmt() (field.EndAccruedInterestAmt, error) {
	var f field.EndAccruedInterestAmt
	err := m.Body.Get(&f)
	return f, err
}

//StartCash is a non-required field for Confirmation.
func (m Confirmation) StartCash() (field.StartCash, error) {
	var f field.StartCash
	err := m.Body.Get(&f)
	return f, err
}

//EndCash is a non-required field for Confirmation.
func (m Confirmation) EndCash() (field.EndCash, error) {
	var f field.EndCash
	err := m.Body.Get(&f)
	return f, err
}

//Concession is a non-required field for Confirmation.
func (m Confirmation) Concession() (field.Concession, error) {
	var f field.Concession
	err := m.Body.Get(&f)
	return f, err
}

//TotalTakedown is a non-required field for Confirmation.
func (m Confirmation) TotalTakedown() (field.TotalTakedown, error) {
	var f field.TotalTakedown
	err := m.Body.Get(&f)
	return f, err
}

//NetMoney is a required field for Confirmation.
func (m Confirmation) NetMoney() (field.NetMoney, error) {
	var f field.NetMoney
	err := m.Body.Get(&f)
	return f, err
}

//MaturityNetMoney is a non-required field for Confirmation.
func (m Confirmation) MaturityNetMoney() (field.MaturityNetMoney, error) {
	var f field.MaturityNetMoney
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrAmt is a non-required field for Confirmation.
func (m Confirmation) SettlCurrAmt() (field.SettlCurrAmt, error) {
	var f field.SettlCurrAmt
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for Confirmation.
func (m Confirmation) SettlCurrency() (field.SettlCurrency, error) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrFxRate is a non-required field for Confirmation.
func (m Confirmation) SettlCurrFxRate() (field.SettlCurrFxRate, error) {
	var f field.SettlCurrFxRate
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrFxRateCalc is a non-required field for Confirmation.
func (m Confirmation) SettlCurrFxRateCalc() (field.SettlCurrFxRateCalc, error) {
	var f field.SettlCurrFxRateCalc
	err := m.Body.Get(&f)
	return f, err
}

//SettlType is a non-required field for Confirmation.
func (m Confirmation) SettlType() (field.SettlType, error) {
	var f field.SettlType
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for Confirmation.
func (m Confirmation) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//SettlDeliveryType is a non-required field for Confirmation.
func (m Confirmation) SettlDeliveryType() (field.SettlDeliveryType, error) {
	var f field.SettlDeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbType is a non-required field for Confirmation.
func (m Confirmation) StandInstDbType() (field.StandInstDbType, error) {
	var f field.StandInstDbType
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbName is a non-required field for Confirmation.
func (m Confirmation) StandInstDbName() (field.StandInstDbName, error) {
	var f field.StandInstDbName
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbID is a non-required field for Confirmation.
func (m Confirmation) StandInstDbID() (field.StandInstDbID, error) {
	var f field.StandInstDbID
	err := m.Body.Get(&f)
	return f, err
}

//NoDlvyInst is a non-required field for Confirmation.
func (m Confirmation) NoDlvyInst() (field.NoDlvyInst, error) {
	var f field.NoDlvyInst
	err := m.Body.Get(&f)
	return f, err
}

//Commission is a non-required field for Confirmation.
func (m Confirmation) Commission() (field.Commission, error) {
	var f field.Commission
	err := m.Body.Get(&f)
	return f, err
}

//CommType is a non-required field for Confirmation.
func (m Confirmation) CommType() (field.CommType, error) {
	var f field.CommType
	err := m.Body.Get(&f)
	return f, err
}

//CommCurrency is a non-required field for Confirmation.
func (m Confirmation) CommCurrency() (field.CommCurrency, error) {
	var f field.CommCurrency
	err := m.Body.Get(&f)
	return f, err
}

//FundRenewWaiv is a non-required field for Confirmation.
func (m Confirmation) FundRenewWaiv() (field.FundRenewWaiv, error) {
	var f field.FundRenewWaiv
	err := m.Body.Get(&f)
	return f, err
}

//SharedCommission is a non-required field for Confirmation.
func (m Confirmation) SharedCommission() (field.SharedCommission, error) {
	var f field.SharedCommission
	err := m.Body.Get(&f)
	return f, err
}

//NoStipulations is a non-required field for Confirmation.
func (m Confirmation) NoStipulations() (field.NoStipulations, error) {
	var f field.NoStipulations
	err := m.Body.Get(&f)
	return f, err
}

//NoMiscFees is a non-required field for Confirmation.
func (m Confirmation) NoMiscFees() (field.NoMiscFees, error) {
	var f field.NoMiscFees
	err := m.Body.Get(&f)
	return f, err
}
