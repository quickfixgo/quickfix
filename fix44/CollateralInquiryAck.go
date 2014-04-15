package fix44

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

//CreateCollateralInquiryAckBuilder returns an initialized CollateralInquiryAckBuilder with specified required fields.
func CreateCollateralInquiryAckBuilder(
	collinquiryid field.CollInquiryID,
	collinquirystatus field.CollInquiryStatus) CollateralInquiryAckBuilder {
	var builder CollateralInquiryAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(collinquiryid)
	builder.Body.Set(collinquirystatus)
	return builder
}

//CollInquiryID is a required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CollInquiryID() (field.CollInquiryID, error) {
	var f field.CollInquiryID
	err := m.Body.Get(&f)
	return f, err
}

//CollInquiryStatus is a required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CollInquiryStatus() (field.CollInquiryStatus, error) {
	var f field.CollInquiryStatus
	err := m.Body.Get(&f)
	return f, err
}

//CollInquiryResult is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CollInquiryResult() (field.CollInquiryResult, error) {
	var f field.CollInquiryResult
	err := m.Body.Get(&f)
	return f, err
}

//NoCollInquiryQualifier is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoCollInquiryQualifier() (field.NoCollInquiryQualifier, error) {
	var f field.NoCollInquiryQualifier
	err := m.Body.Get(&f)
	return f, err
}

//TotNumReports is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) TotNumReports() (field.TotNumReports, error) {
	var f field.TotNumReports
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//OrderID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryOrderID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecondaryOrderID() (field.SecondaryOrderID, error) {
	var f field.SecondaryOrderID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryClOrdID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecondaryClOrdID() (field.SecondaryClOrdID, error) {
	var f field.SecondaryClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//NoExecs is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoExecs() (field.NoExecs, error) {
	var f field.NoExecs
	err := m.Body.Get(&f)
	return f, err
}

//NoTrades is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoTrades() (field.NoTrades, error) {
	var f field.NoTrades
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) AgreementDesc() (field.AgreementDesc, error) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) AgreementID() (field.AgreementID, error) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) AgreementDate() (field.AgreementDate, error) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) AgreementCurrency() (field.AgreementCurrency, error) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) TerminationType() (field.TerminationType, error) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StartDate() (field.StartDate, error) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EndDate() (field.EndDate, error) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) DeliveryType() (field.DeliveryType, error) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) MarginRatio() (field.MarginRatio, error) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//SettlDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SettlDate() (field.SettlDate, error) {
	var f field.SettlDate
	err := m.Body.Get(&f)
	return f, err
}

//Quantity is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Quantity() (field.Quantity, error) {
	var f field.Quantity
	err := m.Body.Get(&f)
	return f, err
}

//QtyType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) QtyType() (field.QtyType, error) {
	var f field.QtyType
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SettlSessID() (field.SettlSessID, error) {
	var f field.SettlSessID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessSubID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SettlSessSubID() (field.SettlSessSubID, error) {
	var f field.SettlSessSubID
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//ResponseTransportType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ResponseTransportType() (field.ResponseTransportType, error) {
	var f field.ResponseTransportType
	err := m.Body.Get(&f)
	return f, err
}

//ResponseDestination is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ResponseDestination() (field.ResponseDestination, error) {
	var f field.ResponseDestination
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
