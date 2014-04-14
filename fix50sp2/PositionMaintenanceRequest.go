package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//PositionMaintenanceRequest msg type = AL.
type PositionMaintenanceRequest struct {
	message.Message
}

//PositionMaintenanceRequestBuilder builds PositionMaintenanceRequest messages.
type PositionMaintenanceRequestBuilder struct {
	message.MessageBuilder
}

//NewPositionMaintenanceRequestBuilder returns an initialized PositionMaintenanceRequestBuilder with specified required fields.
func NewPositionMaintenanceRequestBuilder(
	postranstype field.PosTransType,
	posmaintaction field.PosMaintAction,
	clearingbusinessdate field.ClearingBusinessDate) *PositionMaintenanceRequestBuilder {
	builder := new(PositionMaintenanceRequestBuilder)
	builder.Body.Set(postranstype)
	builder.Body.Set(posmaintaction)
	builder.Body.Set(clearingbusinessdate)
	return builder
}

//PosReqID is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) PosReqID() (*field.PosReqID, error) {
	f := new(field.PosReqID)
	err := m.Body.Get(f)
	return f, err
}

//PosTransType is a required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) PosTransType() (*field.PosTransType, error) {
	f := new(field.PosTransType)
	err := m.Body.Get(f)
	return f, err
}

//PosMaintAction is a required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) PosMaintAction() (*field.PosMaintAction, error) {
	f := new(field.PosMaintAction)
	err := m.Body.Get(f)
	return f, err
}

//OrigPosReqRefID is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) OrigPosReqRefID() (*field.OrigPosReqRefID, error) {
	f := new(field.OrigPosReqRefID)
	err := m.Body.Get(f)
	return f, err
}

//PosMaintRptRefID is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) PosMaintRptRefID() (*field.PosMaintRptRefID, error) {
	f := new(field.PosMaintRptRefID)
	err := m.Body.Get(f)
	return f, err
}

//ClearingBusinessDate is a required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//SettlSessID is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//SettlSessSubID is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//AcctIDSource is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//AccountType is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//SecurityGroup is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SecurityGroup() (*field.SecurityGroup, error) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrementAmount is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, error) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasureQty is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) UnitOfMeasureQty() (*field.UnitOfMeasureQty, error) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLLen is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SecurityXMLLen() (*field.SecurityXMLLen, error) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXML is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SecurityXML() (*field.SecurityXML, error) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//SecurityXMLSchema is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SecurityXMLSchema() (*field.SecurityXMLSchema, error) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//ProductComplex is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) ProductComplex() (*field.ProductComplex, error) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasure is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, error) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//PriceUnitOfMeasureQty is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, error) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//SettlMethod is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SettlMethod() (*field.SettlMethod, error) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//ExerciseStyle is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) ExerciseStyle() (*field.ExerciseStyle, error) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//OptPayoutAmount is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) OptPayoutAmount() (*field.OptPayoutAmount, error) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}

//PriceQuoteMethod is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) PriceQuoteMethod() (*field.PriceQuoteMethod, error) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//ListMethod is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) ListMethod() (*field.ListMethod, error) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//CapPrice is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) CapPrice() (*field.CapPrice, error) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//FloorPrice is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) FloorPrice() (*field.FloorPrice, error) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//PutOrCall is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) PutOrCall() (*field.PutOrCall, error) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//FlexibleIndicator is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) FlexibleIndicator() (*field.FlexibleIndicator, error) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//FlexProductEligibilityIndicator is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, error) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//ValuationMethod is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) ValuationMethod() (*field.ValuationMethod, error) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplierUnit is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) ContractMultiplierUnit() (*field.ContractMultiplierUnit, error) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}

//FlowScheduleType is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) FlowScheduleType() (*field.FlowScheduleType, error) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}

//RestructuringType is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) RestructuringType() (*field.RestructuringType, error) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}

//Seniority is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) Seniority() (*field.Seniority, error) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}

//NotionalPercentageOutstanding is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, error) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//OriginalNotionalPercentageOutstanding is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, error) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//AttachmentPoint is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) AttachmentPoint() (*field.AttachmentPoint, error) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//DetachmentPoint is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) DetachmentPoint() (*field.DetachmentPoint, error) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//StrikePriceDeterminationMethod is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, error) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikePriceBoundaryMethod is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, error) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikePriceBoundaryPrecision is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, error) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}

//UnderlyingPriceDeterminationMethod is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, error) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//OptPayoutType is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) OptPayoutType() (*field.OptPayoutType, error) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}

//NoComplexEvents is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) NoComplexEvents() (*field.NoComplexEvents, error) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//NoTradingSessions is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) NoTradingSessions() (*field.NoTradingSessions, error) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//NoPositions is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) NoPositions() (*field.NoPositions, error) {
	f := new(field.NoPositions)
	err := m.Body.Get(f)
	return f, err
}

//AdjustmentType is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) AdjustmentType() (*field.AdjustmentType, error) {
	f := new(field.AdjustmentType)
	err := m.Body.Get(f)
	return f, err
}

//ContraryInstructionIndicator is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) ContraryInstructionIndicator() (*field.ContraryInstructionIndicator, error) {
	f := new(field.ContraryInstructionIndicator)
	err := m.Body.Get(f)
	return f, err
}

//PriorSpreadIndicator is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) PriorSpreadIndicator() (*field.PriorSpreadIndicator, error) {
	f := new(field.PriorSpreadIndicator)
	err := m.Body.Get(f)
	return f, err
}

//ThresholdAmount is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) ThresholdAmount() (*field.ThresholdAmount, error) {
	f := new(field.ThresholdAmount)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//NoPosAmt is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) NoPosAmt() (*field.NoPosAmt, error) {
	f := new(field.NoPosAmt)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrency is a non-required field for PositionMaintenanceRequest.
func (m *PositionMaintenanceRequest) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}
