package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CollateralInquiry msg type = BB.
type CollateralInquiry struct {
	message.Message
}

//CollateralInquiryBuilder builds CollateralInquiry messages.
type CollateralInquiryBuilder struct {
	message.MessageBuilder
}

//CreateCollateralInquiryBuilder returns an initialized CollateralInquiryBuilder with specified required fields.
func CreateCollateralInquiryBuilder() CollateralInquiryBuilder {
	var builder CollateralInquiryBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	return builder
}

//CollInquiryID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CollInquiryID() (*field.CollInquiryID, errors.MessageRejectError) {
	f := new(field.CollInquiryID)
	err := m.Body.Get(f)
	return f, err
}

//GetCollInquiryID reads a CollInquiryID from CollateralInquiry.
func (m CollateralInquiry) GetCollInquiryID(f *field.CollInquiryID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoCollInquiryQualifier is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoCollInquiryQualifier() (*field.NoCollInquiryQualifier, errors.MessageRejectError) {
	f := new(field.NoCollInquiryQualifier)
	err := m.Body.Get(f)
	return f, err
}

//GetNoCollInquiryQualifier reads a NoCollInquiryQualifier from CollateralInquiry.
func (m CollateralInquiry) GetNoCollInquiryQualifier(f *field.NoCollInquiryQualifier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SubscriptionRequestType() (*field.SubscriptionRequestType, errors.MessageRejectError) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from CollateralInquiry.
func (m CollateralInquiry) GetSubscriptionRequestType(f *field.SubscriptionRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseTransportType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) ResponseTransportType() (*field.ResponseTransportType, errors.MessageRejectError) {
	f := new(field.ResponseTransportType)
	err := m.Body.Get(f)
	return f, err
}

//GetResponseTransportType reads a ResponseTransportType from CollateralInquiry.
func (m CollateralInquiry) GetResponseTransportType(f *field.ResponseTransportType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseDestination is a non-required field for CollateralInquiry.
func (m CollateralInquiry) ResponseDestination() (*field.ResponseDestination, errors.MessageRejectError) {
	f := new(field.ResponseDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetResponseDestination reads a ResponseDestination from CollateralInquiry.
func (m CollateralInquiry) GetResponseDestination(f *field.ResponseDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from CollateralInquiry.
func (m CollateralInquiry) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from CollateralInquiry.
func (m CollateralInquiry) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from CollateralInquiry.
func (m CollateralInquiry) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from CollateralInquiry.
func (m CollateralInquiry) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from CollateralInquiry.
func (m CollateralInquiry) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecondaryOrderID() (*field.SecondaryOrderID, errors.MessageRejectError) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from CollateralInquiry.
func (m CollateralInquiry) GetSecondaryOrderID(f *field.SecondaryOrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecondaryClOrdID() (*field.SecondaryClOrdID, errors.MessageRejectError) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from CollateralInquiry.
func (m CollateralInquiry) GetSecondaryClOrdID(f *field.SecondaryClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecs is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoExecs() (*field.NoExecs, errors.MessageRejectError) {
	f := new(field.NoExecs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecs reads a NoExecs from CollateralInquiry.
func (m CollateralInquiry) GetNoExecs(f *field.NoExecs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrades is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoTrades() (*field.NoTrades, errors.MessageRejectError) {
	f := new(field.NoTrades)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrades reads a NoTrades from CollateralInquiry.
func (m CollateralInquiry) GetNoTrades(f *field.NoTrades) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from CollateralInquiry.
func (m CollateralInquiry) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from CollateralInquiry.
func (m CollateralInquiry) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from CollateralInquiry.
func (m CollateralInquiry) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from CollateralInquiry.
func (m CollateralInquiry) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from CollateralInquiry.
func (m CollateralInquiry) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from CollateralInquiry.
func (m CollateralInquiry) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from CollateralInquiry.
func (m CollateralInquiry) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from CollateralInquiry.
func (m CollateralInquiry) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from CollateralInquiry.
func (m CollateralInquiry) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for CollateralInquiry.
func (m CollateralInquiry) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from CollateralInquiry.
func (m CollateralInquiry) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from CollateralInquiry.
func (m CollateralInquiry) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from CollateralInquiry.
func (m CollateralInquiry) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from CollateralInquiry.
func (m CollateralInquiry) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from CollateralInquiry.
func (m CollateralInquiry) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for CollateralInquiry.
func (m CollateralInquiry) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from CollateralInquiry.
func (m CollateralInquiry) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from CollateralInquiry.
func (m CollateralInquiry) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from CollateralInquiry.
func (m CollateralInquiry) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from CollateralInquiry.
func (m CollateralInquiry) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for CollateralInquiry.
func (m CollateralInquiry) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from CollateralInquiry.
func (m CollateralInquiry) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from CollateralInquiry.
func (m CollateralInquiry) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from CollateralInquiry.
func (m CollateralInquiry) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for CollateralInquiry.
func (m CollateralInquiry) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from CollateralInquiry.
func (m CollateralInquiry) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from CollateralInquiry.
func (m CollateralInquiry) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from CollateralInquiry.
func (m CollateralInquiry) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from CollateralInquiry.
func (m CollateralInquiry) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for CollateralInquiry.
func (m CollateralInquiry) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from CollateralInquiry.
func (m CollateralInquiry) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for CollateralInquiry.
func (m CollateralInquiry) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from CollateralInquiry.
func (m CollateralInquiry) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from CollateralInquiry.
func (m CollateralInquiry) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from CollateralInquiry.
func (m CollateralInquiry) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from CollateralInquiry.
func (m CollateralInquiry) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from CollateralInquiry.
func (m CollateralInquiry) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from CollateralInquiry.
func (m CollateralInquiry) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from CollateralInquiry.
func (m CollateralInquiry) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from CollateralInquiry.
func (m CollateralInquiry) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from CollateralInquiry.
func (m CollateralInquiry) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from CollateralInquiry.
func (m CollateralInquiry) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for CollateralInquiry.
func (m CollateralInquiry) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from CollateralInquiry.
func (m CollateralInquiry) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from CollateralInquiry.
func (m CollateralInquiry) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from CollateralInquiry.
func (m CollateralInquiry) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from CollateralInquiry.
func (m CollateralInquiry) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from CollateralInquiry.
func (m CollateralInquiry) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from CollateralInquiry.
func (m CollateralInquiry) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from CollateralInquiry.
func (m CollateralInquiry) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from CollateralInquiry.
func (m CollateralInquiry) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for CollateralInquiry.
func (m CollateralInquiry) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from CollateralInquiry.
func (m CollateralInquiry) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from CollateralInquiry.
func (m CollateralInquiry) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from CollateralInquiry.
func (m CollateralInquiry) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for CollateralInquiry.
func (m CollateralInquiry) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from CollateralInquiry.
func (m CollateralInquiry) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for CollateralInquiry.
func (m CollateralInquiry) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from CollateralInquiry.
func (m CollateralInquiry) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from CollateralInquiry.
func (m CollateralInquiry) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from CollateralInquiry.
func (m CollateralInquiry) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for CollateralInquiry.
func (m CollateralInquiry) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from CollateralInquiry.
func (m CollateralInquiry) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for CollateralInquiry.
func (m CollateralInquiry) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from CollateralInquiry.
func (m CollateralInquiry) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for CollateralInquiry.
func (m CollateralInquiry) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from CollateralInquiry.
func (m CollateralInquiry) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecurityGroup() (*field.SecurityGroup, errors.MessageRejectError) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from CollateralInquiry.
func (m CollateralInquiry) GetSecurityGroup(f *field.SecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for CollateralInquiry.
func (m CollateralInquiry) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from CollateralInquiry.
func (m CollateralInquiry) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for CollateralInquiry.
func (m CollateralInquiry) UnitOfMeasureQty() (*field.UnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from CollateralInquiry.
func (m CollateralInquiry) GetUnitOfMeasureQty(f *field.UnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecurityXMLLen() (*field.SecurityXMLLen, errors.MessageRejectError) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from CollateralInquiry.
func (m CollateralInquiry) GetSecurityXMLLen(f *field.SecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecurityXML() (*field.SecurityXML, errors.MessageRejectError) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from CollateralInquiry.
func (m CollateralInquiry) GetSecurityXML(f *field.SecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SecurityXMLSchema() (*field.SecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from CollateralInquiry.
func (m CollateralInquiry) GetSecurityXMLSchema(f *field.SecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for CollateralInquiry.
func (m CollateralInquiry) ProductComplex() (*field.ProductComplex, errors.MessageRejectError) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from CollateralInquiry.
func (m CollateralInquiry) GetProductComplex(f *field.ProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for CollateralInquiry.
func (m CollateralInquiry) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from CollateralInquiry.
func (m CollateralInquiry) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for CollateralInquiry.
func (m CollateralInquiry) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from CollateralInquiry.
func (m CollateralInquiry) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SettlMethod() (*field.SettlMethod, errors.MessageRejectError) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from CollateralInquiry.
func (m CollateralInquiry) GetSettlMethod(f *field.SettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for CollateralInquiry.
func (m CollateralInquiry) ExerciseStyle() (*field.ExerciseStyle, errors.MessageRejectError) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from CollateralInquiry.
func (m CollateralInquiry) GetExerciseStyle(f *field.ExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayAmount is a non-required field for CollateralInquiry.
func (m CollateralInquiry) OptPayAmount() (*field.OptPayAmount, errors.MessageRejectError) {
	f := new(field.OptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayAmount reads a OptPayAmount from CollateralInquiry.
func (m CollateralInquiry) GetOptPayAmount(f *field.OptPayAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for CollateralInquiry.
func (m CollateralInquiry) PriceQuoteMethod() (*field.PriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from CollateralInquiry.
func (m CollateralInquiry) GetPriceQuoteMethod(f *field.PriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for CollateralInquiry.
func (m CollateralInquiry) ListMethod() (*field.ListMethod, errors.MessageRejectError) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from CollateralInquiry.
func (m CollateralInquiry) GetListMethod(f *field.ListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CapPrice() (*field.CapPrice, errors.MessageRejectError) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from CollateralInquiry.
func (m CollateralInquiry) GetCapPrice(f *field.CapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for CollateralInquiry.
func (m CollateralInquiry) FloorPrice() (*field.FloorPrice, errors.MessageRejectError) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from CollateralInquiry.
func (m CollateralInquiry) GetFloorPrice(f *field.FloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for CollateralInquiry.
func (m CollateralInquiry) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from CollateralInquiry.
func (m CollateralInquiry) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for CollateralInquiry.
func (m CollateralInquiry) FlexibleIndicator() (*field.FlexibleIndicator, errors.MessageRejectError) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from CollateralInquiry.
func (m CollateralInquiry) GetFlexibleIndicator(f *field.FlexibleIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for CollateralInquiry.
func (m CollateralInquiry) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from CollateralInquiry.
func (m CollateralInquiry) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FuturesValuationMethod is a non-required field for CollateralInquiry.
func (m CollateralInquiry) FuturesValuationMethod() (*field.FuturesValuationMethod, errors.MessageRejectError) {
	f := new(field.FuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetFuturesValuationMethod reads a FuturesValuationMethod from CollateralInquiry.
func (m CollateralInquiry) GetFuturesValuationMethod(f *field.FuturesValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for CollateralInquiry.
func (m CollateralInquiry) AgreementDesc() (*field.AgreementDesc, errors.MessageRejectError) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from CollateralInquiry.
func (m CollateralInquiry) GetAgreementDesc(f *field.AgreementDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) AgreementID() (*field.AgreementID, errors.MessageRejectError) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from CollateralInquiry.
func (m CollateralInquiry) GetAgreementID(f *field.AgreementID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) AgreementDate() (*field.AgreementDate, errors.MessageRejectError) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from CollateralInquiry.
func (m CollateralInquiry) GetAgreementDate(f *field.AgreementDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for CollateralInquiry.
func (m CollateralInquiry) AgreementCurrency() (*field.AgreementCurrency, errors.MessageRejectError) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from CollateralInquiry.
func (m CollateralInquiry) GetAgreementCurrency(f *field.AgreementCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) TerminationType() (*field.TerminationType, errors.MessageRejectError) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from CollateralInquiry.
func (m CollateralInquiry) GetTerminationType(f *field.TerminationType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StartDate() (*field.StartDate, errors.MessageRejectError) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from CollateralInquiry.
func (m CollateralInquiry) GetStartDate(f *field.StartDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EndDate() (*field.EndDate, errors.MessageRejectError) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from CollateralInquiry.
func (m CollateralInquiry) GetEndDate(f *field.EndDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) DeliveryType() (*field.DeliveryType, errors.MessageRejectError) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from CollateralInquiry.
func (m CollateralInquiry) GetDeliveryType(f *field.DeliveryType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for CollateralInquiry.
func (m CollateralInquiry) MarginRatio() (*field.MarginRatio, errors.MessageRejectError) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from CollateralInquiry.
func (m CollateralInquiry) GetMarginRatio(f *field.MarginRatio) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SettlDate() (*field.SettlDate, errors.MessageRejectError) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from CollateralInquiry.
func (m CollateralInquiry) GetSettlDate(f *field.SettlDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Quantity is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Quantity() (*field.Quantity, errors.MessageRejectError) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}

//GetQuantity reads a Quantity from CollateralInquiry.
func (m CollateralInquiry) GetQuantity(f *field.Quantity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) QtyType() (*field.QtyType, errors.MessageRejectError) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from CollateralInquiry.
func (m CollateralInquiry) GetQtyType(f *field.QtyType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from CollateralInquiry.
func (m CollateralInquiry) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from CollateralInquiry.
func (m CollateralInquiry) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from CollateralInquiry.
func (m CollateralInquiry) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginExcess is a non-required field for CollateralInquiry.
func (m CollateralInquiry) MarginExcess() (*field.MarginExcess, errors.MessageRejectError) {
	f := new(field.MarginExcess)
	err := m.Body.Get(f)
	return f, err
}

//GetMarginExcess reads a MarginExcess from CollateralInquiry.
func (m CollateralInquiry) GetMarginExcess(f *field.MarginExcess) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalNetValue is a non-required field for CollateralInquiry.
func (m CollateralInquiry) TotalNetValue() (*field.TotalNetValue, errors.MessageRejectError) {
	f := new(field.TotalNetValue)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalNetValue reads a TotalNetValue from CollateralInquiry.
func (m CollateralInquiry) GetTotalNetValue(f *field.TotalNetValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOutstanding is a non-required field for CollateralInquiry.
func (m CollateralInquiry) CashOutstanding() (*field.CashOutstanding, errors.MessageRejectError) {
	f := new(field.CashOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetCashOutstanding reads a CashOutstanding from CollateralInquiry.
func (m CollateralInquiry) GetCashOutstanding(f *field.CashOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrdRegTimestamps is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoTrdRegTimestamps() (*field.NoTrdRegTimestamps, errors.MessageRejectError) {
	f := new(field.NoTrdRegTimestamps)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrdRegTimestamps reads a NoTrdRegTimestamps from CollateralInquiry.
func (m CollateralInquiry) GetNoTrdRegTimestamps(f *field.NoTrdRegTimestamps) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from CollateralInquiry.
func (m CollateralInquiry) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from CollateralInquiry.
func (m CollateralInquiry) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from CollateralInquiry.
func (m CollateralInquiry) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestAmt is a non-required field for CollateralInquiry.
func (m CollateralInquiry) AccruedInterestAmt() (*field.AccruedInterestAmt, errors.MessageRejectError) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestAmt reads a AccruedInterestAmt from CollateralInquiry.
func (m CollateralInquiry) GetAccruedInterestAmt(f *field.AccruedInterestAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndAccruedInterestAmt is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EndAccruedInterestAmt() (*field.EndAccruedInterestAmt, errors.MessageRejectError) {
	f := new(field.EndAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetEndAccruedInterestAmt reads a EndAccruedInterestAmt from CollateralInquiry.
func (m CollateralInquiry) GetEndAccruedInterestAmt(f *field.EndAccruedInterestAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartCash is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StartCash() (*field.StartCash, errors.MessageRejectError) {
	f := new(field.StartCash)
	err := m.Body.Get(f)
	return f, err
}

//GetStartCash reads a StartCash from CollateralInquiry.
func (m CollateralInquiry) GetStartCash(f *field.StartCash) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndCash is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EndCash() (*field.EndCash, errors.MessageRejectError) {
	f := new(field.EndCash)
	err := m.Body.Get(f)
	return f, err
}

//GetEndCash reads a EndCash from CollateralInquiry.
func (m CollateralInquiry) GetEndCash(f *field.EndCash) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Spread() (*field.Spread, errors.MessageRejectError) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from CollateralInquiry.
func (m CollateralInquiry) GetSpread(f *field.Spread) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for CollateralInquiry.
func (m CollateralInquiry) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from CollateralInquiry.
func (m CollateralInquiry) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for CollateralInquiry.
func (m CollateralInquiry) BenchmarkCurveName() (*field.BenchmarkCurveName, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from CollateralInquiry.
func (m CollateralInquiry) GetBenchmarkCurveName(f *field.BenchmarkCurveName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for CollateralInquiry.
func (m CollateralInquiry) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, errors.MessageRejectError) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from CollateralInquiry.
func (m CollateralInquiry) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for CollateralInquiry.
func (m CollateralInquiry) BenchmarkPrice() (*field.BenchmarkPrice, errors.MessageRejectError) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from CollateralInquiry.
func (m CollateralInquiry) GetBenchmarkPrice(f *field.BenchmarkPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) BenchmarkPriceType() (*field.BenchmarkPriceType, errors.MessageRejectError) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from CollateralInquiry.
func (m CollateralInquiry) GetBenchmarkPriceType(f *field.BenchmarkPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) BenchmarkSecurityID() (*field.BenchmarkSecurityID, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from CollateralInquiry.
func (m CollateralInquiry) GetBenchmarkSecurityID(f *field.BenchmarkSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for CollateralInquiry.
func (m CollateralInquiry) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from CollateralInquiry.
func (m CollateralInquiry) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoStipulations() (*field.NoStipulations, errors.MessageRejectError) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from CollateralInquiry.
func (m CollateralInquiry) GetNoStipulations(f *field.NoStipulations) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDeliveryType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SettlDeliveryType() (*field.SettlDeliveryType, errors.MessageRejectError) {
	f := new(field.SettlDeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDeliveryType reads a SettlDeliveryType from CollateralInquiry.
func (m CollateralInquiry) GetSettlDeliveryType(f *field.SettlDeliveryType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbType is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StandInstDbType() (*field.StandInstDbType, errors.MessageRejectError) {
	f := new(field.StandInstDbType)
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbType reads a StandInstDbType from CollateralInquiry.
func (m CollateralInquiry) GetStandInstDbType(f *field.StandInstDbType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbName is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StandInstDbName() (*field.StandInstDbName, errors.MessageRejectError) {
	f := new(field.StandInstDbName)
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbName reads a StandInstDbName from CollateralInquiry.
func (m CollateralInquiry) GetStandInstDbName(f *field.StandInstDbName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) StandInstDbID() (*field.StandInstDbID, errors.MessageRejectError) {
	f := new(field.StandInstDbID)
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbID reads a StandInstDbID from CollateralInquiry.
func (m CollateralInquiry) GetStandInstDbID(f *field.StandInstDbID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDlvyInst is a non-required field for CollateralInquiry.
func (m CollateralInquiry) NoDlvyInst() (*field.NoDlvyInst, errors.MessageRejectError) {
	f := new(field.NoDlvyInst)
	err := m.Body.Get(f)
	return f, err
}

//GetNoDlvyInst reads a NoDlvyInst from CollateralInquiry.
func (m CollateralInquiry) GetNoDlvyInst(f *field.NoDlvyInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from CollateralInquiry.
func (m CollateralInquiry) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from CollateralInquiry.
func (m CollateralInquiry) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SettlSessID() (*field.SettlSessID, errors.MessageRejectError) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from CollateralInquiry.
func (m CollateralInquiry) GetSettlSessID(f *field.SettlSessID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for CollateralInquiry.
func (m CollateralInquiry) SettlSessSubID() (*field.SettlSessSubID, errors.MessageRejectError) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from CollateralInquiry.
func (m CollateralInquiry) GetSettlSessSubID(f *field.SettlSessSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for CollateralInquiry.
func (m CollateralInquiry) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from CollateralInquiry.
func (m CollateralInquiry) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for CollateralInquiry.
func (m CollateralInquiry) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from CollateralInquiry.
func (m CollateralInquiry) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from CollateralInquiry.
func (m CollateralInquiry) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for CollateralInquiry.
func (m CollateralInquiry) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from CollateralInquiry.
func (m CollateralInquiry) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
