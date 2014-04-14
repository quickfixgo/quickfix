package fix50sp1

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

//NewConfirmationBuilder returns an initialized ConfirmationBuilder with specified required fields.
func NewConfirmationBuilder(
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
	netmoney field.NetMoney) *ConfirmationBuilder {
	builder := new(ConfirmationBuilder)
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
func (m *Confirmation) ConfirmID() (*field.ConfirmID, error) {
	f := new(field.ConfirmID)
	err := m.Body.Get(f)
	return f, err
}

//ConfirmRefID is a non-required field for Confirmation.
func (m *Confirmation) ConfirmRefID() (*field.ConfirmRefID, error) {
	f := new(field.ConfirmRefID)
	err := m.Body.Get(f)
	return f, err
}

//ConfirmReqID is a non-required field for Confirmation.
func (m *Confirmation) ConfirmReqID() (*field.ConfirmReqID, error) {
	f := new(field.ConfirmReqID)
	err := m.Body.Get(f)
	return f, err
}

//ConfirmTransType is a required field for Confirmation.
func (m *Confirmation) ConfirmTransType() (*field.ConfirmTransType, error) {
	f := new(field.ConfirmTransType)
	err := m.Body.Get(f)
	return f, err
}

//ConfirmType is a required field for Confirmation.
func (m *Confirmation) ConfirmType() (*field.ConfirmType, error) {
	f := new(field.ConfirmType)
	err := m.Body.Get(f)
	return f, err
}

//CopyMsgIndicator is a non-required field for Confirmation.
func (m *Confirmation) CopyMsgIndicator() (*field.CopyMsgIndicator, error) {
	f := new(field.CopyMsgIndicator)
	err := m.Body.Get(f)
	return f, err
}

//LegalConfirm is a non-required field for Confirmation.
func (m *Confirmation) LegalConfirm() (*field.LegalConfirm, error) {
	f := new(field.LegalConfirm)
	err := m.Body.Get(f)
	return f, err
}

//ConfirmStatus is a required field for Confirmation.
func (m *Confirmation) ConfirmStatus() (*field.ConfirmStatus, error) {
	f := new(field.ConfirmStatus)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for Confirmation.
func (m *Confirmation) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//NoOrders is a non-required field for Confirmation.
func (m *Confirmation) NoOrders() (*field.NoOrders, error) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}

//AllocID is a non-required field for Confirmation.
func (m *Confirmation) AllocID() (*field.AllocID, error) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryAllocID is a non-required field for Confirmation.
func (m *Confirmation) SecondaryAllocID() (*field.SecondaryAllocID, error) {
	f := new(field.SecondaryAllocID)
	err := m.Body.Get(f)
	return f, err
}

//IndividualAllocID is a non-required field for Confirmation.
func (m *Confirmation) IndividualAllocID() (*field.IndividualAllocID, error) {
	f := new(field.IndividualAllocID)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a required field for Confirmation.
func (m *Confirmation) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a required field for Confirmation.
func (m *Confirmation) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//NoTrdRegTimestamps is a non-required field for Confirmation.
func (m *Confirmation) NoTrdRegTimestamps() (*field.NoTrdRegTimestamps, error) {
	f := new(field.NoTrdRegTimestamps)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for Confirmation.
func (m *Confirmation) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for Confirmation.
func (m *Confirmation) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for Confirmation.
func (m *Confirmation) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for Confirmation.
func (m *Confirmation) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for Confirmation.
func (m *Confirmation) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for Confirmation.
func (m *Confirmation) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for Confirmation.
func (m *Confirmation) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for Confirmation.
func (m *Confirmation) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for Confirmation.
func (m *Confirmation) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for Confirmation.
func (m *Confirmation) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for Confirmation.
func (m *Confirmation) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for Confirmation.
func (m *Confirmation) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for Confirmation.
func (m *Confirmation) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for Confirmation.
func (m *Confirmation) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for Confirmation.
func (m *Confirmation) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for Confirmation.
func (m *Confirmation) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for Confirmation.
func (m *Confirmation) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for Confirmation.
func (m *Confirmation) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for Confirmation.
func (m *Confirmation) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for Confirmation.
func (m *Confirmation) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for Confirmation.
func (m *Confirmation) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for Confirmation.
func (m *Confirmation) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for Confirmation.
func (m *Confirmation) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for Confirmation.
func (m *Confirmation) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for Confirmation.
func (m *Confirmation) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for Confirmation.
func (m *Confirmation) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for Confirmation.
func (m *Confirmation) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for Confirmation.
func (m *Confirmation) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for Confirmation.
func (m *Confirmation) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for Confirmation.
func (m *Confirmation) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for Confirmation.
func (m *Confirmation) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for Confirmation.
func (m *Confirmation) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for Confirmation.
func (m *Confirmation) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for Confirmation.
func (m *Confirmation) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for Confirmation.
func (m *Confirmation) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for Confirmation.
func (m *Confirmation) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for Confirmation.
func (m *Confirmation) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for Confirmation.
func (m *Confirmation) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for Confirmation.
func (m *Confirmation) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for Confirmation.
func (m *Confirmation) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for Confirmation.
func (m *Confirmation) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for Confirmation.
func (m *Confirmation) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for Confirmation.
func (m *Confirmation) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for Confirmation.
func (m *Confirmation) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for Confirmation.
func (m *Confirmation) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for Confirmation.
func (m *Confirmation) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for Confirmation.
func (m *Confirmation) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for Confirmation.
func (m *Confirmation) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for Confirmation.
func (m *Confirmation) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for Confirmation.
func (m *Confirmation) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for Confirmation.
func (m *Confirmation) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for Confirmation.
func (m *Confirmation) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for Confirmation.
func (m *Confirmation) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for Confirmation.
func (m *Confirmation) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//SecurityGroup is a non-required field for Confirmation.
func (m *Confirmation) SecurityGroup() (*field.SecurityGroup, error) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for Confirmation.
func (m *Confirmation) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, error) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for Confirmation.
func (m *Confirmation) UnitOfMeasureQty() (*field.UnitOfMeasureQty, error) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLLen is a non-required field for Confirmation.
func (m *Confirmation) SecurityXMLLen() (*field.SecurityXMLLen, error) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXML is a non-required field for Confirmation.
func (m *Confirmation) SecurityXML() (*field.SecurityXML, error) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLSchema is a non-required field for Confirmation.
func (m *Confirmation) SecurityXMLSchema() (*field.SecurityXMLSchema, error) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//ProductComplex is a non-required field for Confirmation.
func (m *Confirmation) ProductComplex() (*field.ProductComplex, error) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for Confirmation.
func (m *Confirmation) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, error) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for Confirmation.
func (m *Confirmation) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, error) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SettlMethod is a non-required field for Confirmation.
func (m *Confirmation) SettlMethod() (*field.SettlMethod, error) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//ExerciseStyle is a non-required field for Confirmation.
func (m *Confirmation) ExerciseStyle() (*field.ExerciseStyle, error) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//OptPayAmount is a non-required field for Confirmation.
func (m *Confirmation) OptPayAmount() (*field.OptPayAmount, error) {
	f := new(field.OptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//PriceQuoteMethod is a non-required field for Confirmation.
func (m *Confirmation) PriceQuoteMethod() (*field.PriceQuoteMethod, error) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//ListMethod is a non-required field for Confirmation.
func (m *Confirmation) ListMethod() (*field.ListMethod, error) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//CapPrice is a non-required field for Confirmation.
func (m *Confirmation) CapPrice() (*field.CapPrice, error) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//FloorPrice is a non-required field for Confirmation.
func (m *Confirmation) FloorPrice() (*field.FloorPrice, error) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for Confirmation.
func (m *Confirmation) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//FlexibleIndicator is a non-required field for Confirmation.
func (m *Confirmation) FlexibleIndicator() (*field.FlexibleIndicator, error) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for Confirmation.
func (m *Confirmation) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, error) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//FuturesValuationMethod is a non-required field for Confirmation.
func (m *Confirmation) FuturesValuationMethod() (*field.FuturesValuationMethod, error) {
	f := new(field.FuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryForm is a non-required field for Confirmation.
func (m *Confirmation) DeliveryForm() (*field.DeliveryForm, error) {
	f := new(field.DeliveryForm)
	err := m.Body.Get(f)
	return f, err
}

//PctAtRisk is a non-required field for Confirmation.
func (m *Confirmation) PctAtRisk() (*field.PctAtRisk, error) {
	f := new(field.PctAtRisk)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrAttrib is a non-required field for Confirmation.
func (m *Confirmation) NoInstrAttrib() (*field.NoInstrAttrib, error) {
	f := new(field.NoInstrAttrib)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDesc is a non-required field for Confirmation.
func (m *Confirmation) AgreementDesc() (*field.AgreementDesc, error) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//AgreementID is a non-required field for Confirmation.
func (m *Confirmation) AgreementID() (*field.AgreementID, error) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDate is a non-required field for Confirmation.
func (m *Confirmation) AgreementDate() (*field.AgreementDate, error) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//AgreementCurrency is a non-required field for Confirmation.
func (m *Confirmation) AgreementCurrency() (*field.AgreementCurrency, error) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//TerminationType is a non-required field for Confirmation.
func (m *Confirmation) TerminationType() (*field.TerminationType, error) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//StartDate is a non-required field for Confirmation.
func (m *Confirmation) StartDate() (*field.StartDate, error) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//EndDate is a non-required field for Confirmation.
func (m *Confirmation) EndDate() (*field.EndDate, error) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryType is a non-required field for Confirmation.
func (m *Confirmation) DeliveryType() (*field.DeliveryType, error) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//MarginRatio is a non-required field for Confirmation.
func (m *Confirmation) MarginRatio() (*field.MarginRatio, error) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for Confirmation.
func (m *Confirmation) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for Confirmation.
func (m *Confirmation) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//YieldType is a non-required field for Confirmation.
func (m *Confirmation) YieldType() (*field.YieldType, error) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//Yield is a non-required field for Confirmation.
func (m *Confirmation) Yield() (*field.Yield, error) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//YieldCalcDate is a non-required field for Confirmation.
func (m *Confirmation) YieldCalcDate() (*field.YieldCalcDate, error) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionDate is a non-required field for Confirmation.
func (m *Confirmation) YieldRedemptionDate() (*field.YieldRedemptionDate, error) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionPrice is a non-required field for Confirmation.
func (m *Confirmation) YieldRedemptionPrice() (*field.YieldRedemptionPrice, error) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//YieldRedemptionPriceType is a non-required field for Confirmation.
func (m *Confirmation) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, error) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//AllocQty is a required field for Confirmation.
func (m *Confirmation) AllocQty() (*field.AllocQty, error) {
	f := new(field.AllocQty)
	err := m.Body.Get(f)
	return f, err
}

//QtyType is a non-required field for Confirmation.
func (m *Confirmation) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//Side is a required field for Confirmation.
func (m *Confirmation) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for Confirmation.
func (m *Confirmation) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//LastMkt is a non-required field for Confirmation.
func (m *Confirmation) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//NoCapacities is a required field for Confirmation.
func (m *Confirmation) NoCapacities() (*field.NoCapacities, error) {
	f := new(field.NoCapacities)
	err := m.Body.Get(f)
	return f, err
}

//AllocAccount is a required field for Confirmation.
func (m *Confirmation) AllocAccount() (*field.AllocAccount, error) {
	f := new(field.AllocAccount)
	err := m.Body.Get(f)
	return f, err
}

//AllocAcctIDSource is a non-required field for Confirmation.
func (m *Confirmation) AllocAcctIDSource() (*field.AllocAcctIDSource, error) {
	f := new(field.AllocAcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//AllocAccountType is a non-required field for Confirmation.
func (m *Confirmation) AllocAccountType() (*field.AllocAccountType, error) {
	f := new(field.AllocAccountType)
	err := m.Body.Get(f)
	return f, err
}

//AvgPx is a required field for Confirmation.
func (m *Confirmation) AvgPx() (*field.AvgPx, error) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}

//AvgPxPrecision is a non-required field for Confirmation.
func (m *Confirmation) AvgPxPrecision() (*field.AvgPxPrecision, error) {
	f := new(field.AvgPxPrecision)
	err := m.Body.Get(f)
	return f, err
}

//PriceType is a non-required field for Confirmation.
func (m *Confirmation) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//AvgParPx is a non-required field for Confirmation.
func (m *Confirmation) AvgParPx() (*field.AvgParPx, error) {
	f := new(field.AvgParPx)
	err := m.Body.Get(f)
	return f, err
}

//Spread is a non-required field for Confirmation.
func (m *Confirmation) Spread() (*field.Spread, error) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveCurrency is a non-required field for Confirmation.
func (m *Confirmation) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, error) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurveName is a non-required field for Confirmation.
func (m *Confirmation) BenchmarkCurveName() (*field.BenchmarkCurveName, error) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkCurvePoint is a non-required field for Confirmation.
func (m *Confirmation) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, error) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPrice is a non-required field for Confirmation.
func (m *Confirmation) BenchmarkPrice() (*field.BenchmarkPrice, error) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkPriceType is a non-required field for Confirmation.
func (m *Confirmation) BenchmarkPriceType() (*field.BenchmarkPriceType, error) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityID is a non-required field for Confirmation.
func (m *Confirmation) BenchmarkSecurityID() (*field.BenchmarkSecurityID, error) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//BenchmarkSecurityIDSource is a non-required field for Confirmation.
func (m *Confirmation) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, error) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//ReportedPx is a non-required field for Confirmation.
func (m *Confirmation) ReportedPx() (*field.ReportedPx, error) {
	f := new(field.ReportedPx)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for Confirmation.
func (m *Confirmation) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for Confirmation.
func (m *Confirmation) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for Confirmation.
func (m *Confirmation) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//ProcessCode is a non-required field for Confirmation.
func (m *Confirmation) ProcessCode() (*field.ProcessCode, error) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}

//GrossTradeAmt is a required field for Confirmation.
func (m *Confirmation) GrossTradeAmt() (*field.GrossTradeAmt, error) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}

//NumDaysInterest is a non-required field for Confirmation.
func (m *Confirmation) NumDaysInterest() (*field.NumDaysInterest, error) {
	f := new(field.NumDaysInterest)
	err := m.Body.Get(f)
	return f, err
}

//ExDate is a non-required field for Confirmation.
func (m *Confirmation) ExDate() (*field.ExDate, error) {
	f := new(field.ExDate)
	err := m.Body.Get(f)
	return f, err
}

//AccruedInterestRate is a non-required field for Confirmation.
func (m *Confirmation) AccruedInterestRate() (*field.AccruedInterestRate, error) {
	f := new(field.AccruedInterestRate)
	err := m.Body.Get(f)
	return f, err
}

//AccruedInterestAmt is a non-required field for Confirmation.
func (m *Confirmation) AccruedInterestAmt() (*field.AccruedInterestAmt, error) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//InterestAtMaturity is a non-required field for Confirmation.
func (m *Confirmation) InterestAtMaturity() (*field.InterestAtMaturity, error) {
	f := new(field.InterestAtMaturity)
	err := m.Body.Get(f)
	return f, err
}

//EndAccruedInterestAmt is a non-required field for Confirmation.
func (m *Confirmation) EndAccruedInterestAmt() (*field.EndAccruedInterestAmt, error) {
	f := new(field.EndAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//StartCash is a non-required field for Confirmation.
func (m *Confirmation) StartCash() (*field.StartCash, error) {
	f := new(field.StartCash)
	err := m.Body.Get(f)
	return f, err
}

//EndCash is a non-required field for Confirmation.
func (m *Confirmation) EndCash() (*field.EndCash, error) {
	f := new(field.EndCash)
	err := m.Body.Get(f)
	return f, err
}

//Concession is a non-required field for Confirmation.
func (m *Confirmation) Concession() (*field.Concession, error) {
	f := new(field.Concession)
	err := m.Body.Get(f)
	return f, err
}

//TotalTakedown is a non-required field for Confirmation.
func (m *Confirmation) TotalTakedown() (*field.TotalTakedown, error) {
	f := new(field.TotalTakedown)
	err := m.Body.Get(f)
	return f, err
}

//NetMoney is a required field for Confirmation.
func (m *Confirmation) NetMoney() (*field.NetMoney, error) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}

//MaturityNetMoney is a non-required field for Confirmation.
func (m *Confirmation) MaturityNetMoney() (*field.MaturityNetMoney, error) {
	f := new(field.MaturityNetMoney)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrAmt is a non-required field for Confirmation.
func (m *Confirmation) SettlCurrAmt() (*field.SettlCurrAmt, error) {
	f := new(field.SettlCurrAmt)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrency is a non-required field for Confirmation.
func (m *Confirmation) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrFxRate is a non-required field for Confirmation.
func (m *Confirmation) SettlCurrFxRate() (*field.SettlCurrFxRate, error) {
	f := new(field.SettlCurrFxRate)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrFxRateCalc is a non-required field for Confirmation.
func (m *Confirmation) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalc, error) {
	f := new(field.SettlCurrFxRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//SettlType is a non-required field for Confirmation.
func (m *Confirmation) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//SettlDate is a non-required field for Confirmation.
func (m *Confirmation) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//SettlDeliveryType is a non-required field for Confirmation.
func (m *Confirmation) SettlDeliveryType() (*field.SettlDeliveryType, error) {
	f := new(field.SettlDeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//StandInstDbType is a non-required field for Confirmation.
func (m *Confirmation) StandInstDbType() (*field.StandInstDbType, error) {
	f := new(field.StandInstDbType)
	err := m.Body.Get(f)
	return f, err
}

//StandInstDbName is a non-required field for Confirmation.
func (m *Confirmation) StandInstDbName() (*field.StandInstDbName, error) {
	f := new(field.StandInstDbName)
	err := m.Body.Get(f)
	return f, err
}

//StandInstDbID is a non-required field for Confirmation.
func (m *Confirmation) StandInstDbID() (*field.StandInstDbID, error) {
	f := new(field.StandInstDbID)
	err := m.Body.Get(f)
	return f, err
}

//NoDlvyInst is a non-required field for Confirmation.
func (m *Confirmation) NoDlvyInst() (*field.NoDlvyInst, error) {
	f := new(field.NoDlvyInst)
	err := m.Body.Get(f)
	return f, err
}

//Commission is a non-required field for Confirmation.
func (m *Confirmation) Commission() (*field.Commission, error) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//CommType is a non-required field for Confirmation.
func (m *Confirmation) CommType() (*field.CommType, error) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//CommCurrency is a non-required field for Confirmation.
func (m *Confirmation) CommCurrency() (*field.CommCurrency, error) {
	f := new(field.CommCurrency)
	err := m.Body.Get(f)
	return f, err
}

//FundRenewWaiv is a non-required field for Confirmation.
func (m *Confirmation) FundRenewWaiv() (*field.FundRenewWaiv, error) {
	f := new(field.FundRenewWaiv)
	err := m.Body.Get(f)
	return f, err
}

//SharedCommission is a non-required field for Confirmation.
func (m *Confirmation) SharedCommission() (*field.SharedCommission, error) {
	f := new(field.SharedCommission)
	err := m.Body.Get(f)
	return f, err
}

//NoStipulations is a non-required field for Confirmation.
func (m *Confirmation) NoStipulations() (*field.NoStipulations, error) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//NoMiscFees is a non-required field for Confirmation.
func (m *Confirmation) NoMiscFees() (*field.NoMiscFees, error) {
	f := new(field.NoMiscFees)
	err := m.Body.Get(f)
	return f, err
}
