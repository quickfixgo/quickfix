package fix50

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CollateralInquiryAck msg type = BG.
type CollateralInquiryAck struct {
	message.Message
}

//CollateralInquiryAckBuilder builds CollateralInquiryAck messages.
type CollateralInquiryAckBuilder struct {
	message.MessageBuilder
}

//NewCollateralInquiryAckBuilder returns an initialized CollateralInquiryAckBuilder with specified required fields.
func NewCollateralInquiryAckBuilder(
	collinquiryid field.CollInquiryID,
	collinquirystatus field.CollInquiryStatus) *CollateralInquiryAckBuilder {
	builder := new(CollateralInquiryAckBuilder)
	builder.Body.Set(collinquiryid)
	builder.Body.Set(collinquirystatus)
	return builder
}

//CollInquiryID is a required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) CollInquiryID() (*field.CollInquiryID, error) {
	f := new(field.CollInquiryID)
	err := m.Body.Get(f)
	return f, err
}

//CollInquiryStatus is a required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) CollInquiryStatus() (*field.CollInquiryStatus, error) {
	f := new(field.CollInquiryStatus)
	err := m.Body.Get(f)
	return f, err
}

//CollInquiryResult is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) CollInquiryResult() (*field.CollInquiryResult, error) {
	f := new(field.CollInquiryResult)
	err := m.Body.Get(f)
	return f, err
}

//NoCollInquiryQualifier is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) NoCollInquiryQualifier() (*field.NoCollInquiryQualifier, error) {
	f := new(field.NoCollInquiryQualifier)
	err := m.Body.Get(f)
	return f, err
}

//TotNumReports is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) TotNumReports() (*field.TotNumReports, error) {
	f := new(field.TotNumReports)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//Account is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//AccountType is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//OrderID is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryOrderID is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) SecondaryOrderID() (*field.SecondaryOrderID, error) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryClOrdID is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) SecondaryClOrdID() (*field.SecondaryClOrdID, error) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//NoExecs is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) NoExecs() (*field.NoExecs, error) {
	f := new(field.NoExecs)
	err := m.Body.Get(f)
	return f, err
}

//NoTrades is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) NoTrades() (*field.NoTrades, error) {
	f := new(field.NoTrades)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityStatus is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//StrikeMultiplier is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//StrikeValue is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//MinPriceIncrement is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//PositionLimit is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NTPositionLimit is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrumentParties is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//UnitOfMeasure is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//TimeUnit is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//MaturityTime is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDesc is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) AgreementDesc() (*field.AgreementDesc, error) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//AgreementID is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) AgreementID() (*field.AgreementID, error) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDate is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) AgreementDate() (*field.AgreementDate, error) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//AgreementCurrency is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) AgreementCurrency() (*field.AgreementCurrency, error) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//TerminationType is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) TerminationType() (*field.TerminationType, error) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//StartDate is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) StartDate() (*field.StartDate, error) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//EndDate is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) EndDate() (*field.EndDate, error) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryType is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) DeliveryType() (*field.DeliveryType, error) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//MarginRatio is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) MarginRatio() (*field.MarginRatio, error) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//SettlDate is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//Quantity is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) Quantity() (*field.Quantity, error) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}

//QtyType is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) QtyType() (*field.QtyType, error) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//Currency is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//SettlSessID is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//SettlSessSubID is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}

//ClearingBusinessDate is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//ResponseTransportType is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) ResponseTransportType() (*field.ResponseTransportType, error) {
	f := new(field.ResponseTransportType)
	err := m.Body.Get(f)
	return f, err
}

//ResponseDestination is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) ResponseDestination() (*field.ResponseDestination, error) {
	f := new(field.ResponseDestination)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for CollateralInquiryAck.
func (m *CollateralInquiryAck) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
