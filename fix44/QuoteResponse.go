package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//QuoteResponse msg type = AJ.
type QuoteResponse struct {
	message.Message
}

//QuoteResponseBuilder builds QuoteResponse messages.
type QuoteResponseBuilder struct {
	message.MessageBuilder
}

//CreateQuoteResponseBuilder returns an initialized QuoteResponseBuilder with specified required fields.
func CreateQuoteResponseBuilder(
	quoterespid field.QuoteRespID,
	quoteresptype field.QuoteRespType) QuoteResponseBuilder {
	var builder QuoteResponseBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("AJ"))
	builder.Body.Set(quoterespid)
	builder.Body.Set(quoteresptype)
	return builder
}

//QuoteRespID is a required field for QuoteResponse.
func (m QuoteResponse) QuoteRespID() (*field.QuoteRespID, errors.MessageRejectError) {
	f := new(field.QuoteRespID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteRespID reads a QuoteRespID from QuoteResponse.
func (m QuoteResponse) GetQuoteRespID(f *field.QuoteRespID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for QuoteResponse.
func (m QuoteResponse) QuoteID() (*field.QuoteID, errors.MessageRejectError) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from QuoteResponse.
func (m QuoteResponse) GetQuoteID(f *field.QuoteID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteRespType is a required field for QuoteResponse.
func (m QuoteResponse) QuoteRespType() (*field.QuoteRespType, errors.MessageRejectError) {
	f := new(field.QuoteRespType)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteRespType reads a QuoteRespType from QuoteResponse.
func (m QuoteResponse) GetQuoteRespType(f *field.QuoteRespType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for QuoteResponse.
func (m QuoteResponse) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from QuoteResponse.
func (m QuoteResponse) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for QuoteResponse.
func (m QuoteResponse) OrderCapacity() (*field.OrderCapacity, errors.MessageRejectError) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from QuoteResponse.
func (m QuoteResponse) GetOrderCapacity(f *field.OrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIID is a non-required field for QuoteResponse.
func (m QuoteResponse) IOIID() (*field.IOIID, errors.MessageRejectError) {
	f := new(field.IOIID)
	err := m.Body.Get(f)
	return f, err
}

//GetIOIID reads a IOIID from QuoteResponse.
func (m QuoteResponse) GetIOIID(f *field.IOIID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteType is a non-required field for QuoteResponse.
func (m QuoteResponse) QuoteType() (*field.QuoteType, errors.MessageRejectError) {
	f := new(field.QuoteType)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteType reads a QuoteType from QuoteResponse.
func (m QuoteResponse) GetQuoteType(f *field.QuoteType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoQuoteQualifiers is a non-required field for QuoteResponse.
func (m QuoteResponse) NoQuoteQualifiers() (*field.NoQuoteQualifiers, errors.MessageRejectError) {
	f := new(field.NoQuoteQualifiers)
	err := m.Body.Get(f)
	return f, err
}

//GetNoQuoteQualifiers reads a NoQuoteQualifiers from QuoteResponse.
func (m QuoteResponse) GetNoQuoteQualifiers(f *field.NoQuoteQualifiers) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for QuoteResponse.
func (m QuoteResponse) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from QuoteResponse.
func (m QuoteResponse) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for QuoteResponse.
func (m QuoteResponse) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from QuoteResponse.
func (m QuoteResponse) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for QuoteResponse.
func (m QuoteResponse) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from QuoteResponse.
func (m QuoteResponse) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for QuoteResponse.
func (m QuoteResponse) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from QuoteResponse.
func (m QuoteResponse) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for QuoteResponse.
func (m QuoteResponse) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from QuoteResponse.
func (m QuoteResponse) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for QuoteResponse.
func (m QuoteResponse) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from QuoteResponse.
func (m QuoteResponse) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for QuoteResponse.
func (m QuoteResponse) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from QuoteResponse.
func (m QuoteResponse) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for QuoteResponse.
func (m QuoteResponse) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from QuoteResponse.
func (m QuoteResponse) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for QuoteResponse.
func (m QuoteResponse) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from QuoteResponse.
func (m QuoteResponse) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for QuoteResponse.
func (m QuoteResponse) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from QuoteResponse.
func (m QuoteResponse) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for QuoteResponse.
func (m QuoteResponse) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from QuoteResponse.
func (m QuoteResponse) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for QuoteResponse.
func (m QuoteResponse) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from QuoteResponse.
func (m QuoteResponse) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for QuoteResponse.
func (m QuoteResponse) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from QuoteResponse.
func (m QuoteResponse) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for QuoteResponse.
func (m QuoteResponse) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from QuoteResponse.
func (m QuoteResponse) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for QuoteResponse.
func (m QuoteResponse) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from QuoteResponse.
func (m QuoteResponse) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for QuoteResponse.
func (m QuoteResponse) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from QuoteResponse.
func (m QuoteResponse) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for QuoteResponse.
func (m QuoteResponse) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from QuoteResponse.
func (m QuoteResponse) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for QuoteResponse.
func (m QuoteResponse) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from QuoteResponse.
func (m QuoteResponse) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for QuoteResponse.
func (m QuoteResponse) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from QuoteResponse.
func (m QuoteResponse) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for QuoteResponse.
func (m QuoteResponse) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from QuoteResponse.
func (m QuoteResponse) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for QuoteResponse.
func (m QuoteResponse) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from QuoteResponse.
func (m QuoteResponse) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for QuoteResponse.
func (m QuoteResponse) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from QuoteResponse.
func (m QuoteResponse) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for QuoteResponse.
func (m QuoteResponse) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from QuoteResponse.
func (m QuoteResponse) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for QuoteResponse.
func (m QuoteResponse) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from QuoteResponse.
func (m QuoteResponse) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for QuoteResponse.
func (m QuoteResponse) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from QuoteResponse.
func (m QuoteResponse) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for QuoteResponse.
func (m QuoteResponse) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from QuoteResponse.
func (m QuoteResponse) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for QuoteResponse.
func (m QuoteResponse) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from QuoteResponse.
func (m QuoteResponse) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for QuoteResponse.
func (m QuoteResponse) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from QuoteResponse.
func (m QuoteResponse) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for QuoteResponse.
func (m QuoteResponse) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from QuoteResponse.
func (m QuoteResponse) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for QuoteResponse.
func (m QuoteResponse) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from QuoteResponse.
func (m QuoteResponse) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for QuoteResponse.
func (m QuoteResponse) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from QuoteResponse.
func (m QuoteResponse) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for QuoteResponse.
func (m QuoteResponse) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from QuoteResponse.
func (m QuoteResponse) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for QuoteResponse.
func (m QuoteResponse) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from QuoteResponse.
func (m QuoteResponse) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for QuoteResponse.
func (m QuoteResponse) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from QuoteResponse.
func (m QuoteResponse) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for QuoteResponse.
func (m QuoteResponse) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from QuoteResponse.
func (m QuoteResponse) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for QuoteResponse.
func (m QuoteResponse) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from QuoteResponse.
func (m QuoteResponse) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for QuoteResponse.
func (m QuoteResponse) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from QuoteResponse.
func (m QuoteResponse) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for QuoteResponse.
func (m QuoteResponse) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from QuoteResponse.
func (m QuoteResponse) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for QuoteResponse.
func (m QuoteResponse) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from QuoteResponse.
func (m QuoteResponse) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for QuoteResponse.
func (m QuoteResponse) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from QuoteResponse.
func (m QuoteResponse) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for QuoteResponse.
func (m QuoteResponse) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from QuoteResponse.
func (m QuoteResponse) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for QuoteResponse.
func (m QuoteResponse) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from QuoteResponse.
func (m QuoteResponse) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for QuoteResponse.
func (m QuoteResponse) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from QuoteResponse.
func (m QuoteResponse) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for QuoteResponse.
func (m QuoteResponse) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from QuoteResponse.
func (m QuoteResponse) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for QuoteResponse.
func (m QuoteResponse) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from QuoteResponse.
func (m QuoteResponse) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for QuoteResponse.
func (m QuoteResponse) AgreementDesc() (*field.AgreementDesc, errors.MessageRejectError) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from QuoteResponse.
func (m QuoteResponse) GetAgreementDesc(f *field.AgreementDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for QuoteResponse.
func (m QuoteResponse) AgreementID() (*field.AgreementID, errors.MessageRejectError) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from QuoteResponse.
func (m QuoteResponse) GetAgreementID(f *field.AgreementID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for QuoteResponse.
func (m QuoteResponse) AgreementDate() (*field.AgreementDate, errors.MessageRejectError) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from QuoteResponse.
func (m QuoteResponse) GetAgreementDate(f *field.AgreementDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for QuoteResponse.
func (m QuoteResponse) AgreementCurrency() (*field.AgreementCurrency, errors.MessageRejectError) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from QuoteResponse.
func (m QuoteResponse) GetAgreementCurrency(f *field.AgreementCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for QuoteResponse.
func (m QuoteResponse) TerminationType() (*field.TerminationType, errors.MessageRejectError) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from QuoteResponse.
func (m QuoteResponse) GetTerminationType(f *field.TerminationType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for QuoteResponse.
func (m QuoteResponse) StartDate() (*field.StartDate, errors.MessageRejectError) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from QuoteResponse.
func (m QuoteResponse) GetStartDate(f *field.StartDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for QuoteResponse.
func (m QuoteResponse) EndDate() (*field.EndDate, errors.MessageRejectError) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from QuoteResponse.
func (m QuoteResponse) GetEndDate(f *field.EndDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for QuoteResponse.
func (m QuoteResponse) DeliveryType() (*field.DeliveryType, errors.MessageRejectError) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from QuoteResponse.
func (m QuoteResponse) GetDeliveryType(f *field.DeliveryType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for QuoteResponse.
func (m QuoteResponse) MarginRatio() (*field.MarginRatio, errors.MessageRejectError) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from QuoteResponse.
func (m QuoteResponse) GetMarginRatio(f *field.MarginRatio) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for QuoteResponse.
func (m QuoteResponse) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from QuoteResponse.
func (m QuoteResponse) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for QuoteResponse.
func (m QuoteResponse) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from QuoteResponse.
func (m QuoteResponse) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for QuoteResponse.
func (m QuoteResponse) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from QuoteResponse.
func (m QuoteResponse) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for QuoteResponse.
func (m QuoteResponse) CashOrderQty() (*field.CashOrderQty, errors.MessageRejectError) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from QuoteResponse.
func (m QuoteResponse) GetCashOrderQty(f *field.CashOrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for QuoteResponse.
func (m QuoteResponse) OrderPercent() (*field.OrderPercent, errors.MessageRejectError) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from QuoteResponse.
func (m QuoteResponse) GetOrderPercent(f *field.OrderPercent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for QuoteResponse.
func (m QuoteResponse) RoundingDirection() (*field.RoundingDirection, errors.MessageRejectError) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from QuoteResponse.
func (m QuoteResponse) GetRoundingDirection(f *field.RoundingDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for QuoteResponse.
func (m QuoteResponse) RoundingModulus() (*field.RoundingModulus, errors.MessageRejectError) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from QuoteResponse.
func (m QuoteResponse) GetRoundingModulus(f *field.RoundingModulus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for QuoteResponse.
func (m QuoteResponse) SettlType() (*field.SettlType, errors.MessageRejectError) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from QuoteResponse.
func (m QuoteResponse) GetSettlType(f *field.SettlType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for QuoteResponse.
func (m QuoteResponse) SettlDate() (*field.SettlDate, errors.MessageRejectError) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from QuoteResponse.
func (m QuoteResponse) GetSettlDate(f *field.SettlDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate2 is a non-required field for QuoteResponse.
func (m QuoteResponse) SettlDate2() (*field.SettlDate2, errors.MessageRejectError) {
	f := new(field.SettlDate2)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate2 reads a SettlDate2 from QuoteResponse.
func (m QuoteResponse) GetSettlDate2(f *field.SettlDate2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty2 is a non-required field for QuoteResponse.
func (m QuoteResponse) OrderQty2() (*field.OrderQty2, errors.MessageRejectError) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty2 reads a OrderQty2 from QuoteResponse.
func (m QuoteResponse) GetOrderQty2(f *field.OrderQty2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for QuoteResponse.
func (m QuoteResponse) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from QuoteResponse.
func (m QuoteResponse) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for QuoteResponse.
func (m QuoteResponse) NoStipulations() (*field.NoStipulations, errors.MessageRejectError) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from QuoteResponse.
func (m QuoteResponse) GetNoStipulations(f *field.NoStipulations) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for QuoteResponse.
func (m QuoteResponse) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from QuoteResponse.
func (m QuoteResponse) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for QuoteResponse.
func (m QuoteResponse) AcctIDSource() (*field.AcctIDSource, errors.MessageRejectError) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from QuoteResponse.
func (m QuoteResponse) GetAcctIDSource(f *field.AcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for QuoteResponse.
func (m QuoteResponse) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from QuoteResponse.
func (m QuoteResponse) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for QuoteResponse.
func (m QuoteResponse) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from QuoteResponse.
func (m QuoteResponse) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidPx is a non-required field for QuoteResponse.
func (m QuoteResponse) BidPx() (*field.BidPx, errors.MessageRejectError) {
	f := new(field.BidPx)
	err := m.Body.Get(f)
	return f, err
}

//GetBidPx reads a BidPx from QuoteResponse.
func (m QuoteResponse) GetBidPx(f *field.BidPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferPx is a non-required field for QuoteResponse.
func (m QuoteResponse) OfferPx() (*field.OfferPx, errors.MessageRejectError) {
	f := new(field.OfferPx)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferPx reads a OfferPx from QuoteResponse.
func (m QuoteResponse) GetOfferPx(f *field.OfferPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MktBidPx is a non-required field for QuoteResponse.
func (m QuoteResponse) MktBidPx() (*field.MktBidPx, errors.MessageRejectError) {
	f := new(field.MktBidPx)
	err := m.Body.Get(f)
	return f, err
}

//GetMktBidPx reads a MktBidPx from QuoteResponse.
func (m QuoteResponse) GetMktBidPx(f *field.MktBidPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MktOfferPx is a non-required field for QuoteResponse.
func (m QuoteResponse) MktOfferPx() (*field.MktOfferPx, errors.MessageRejectError) {
	f := new(field.MktOfferPx)
	err := m.Body.Get(f)
	return f, err
}

//GetMktOfferPx reads a MktOfferPx from QuoteResponse.
func (m QuoteResponse) GetMktOfferPx(f *field.MktOfferPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinBidSize is a non-required field for QuoteResponse.
func (m QuoteResponse) MinBidSize() (*field.MinBidSize, errors.MessageRejectError) {
	f := new(field.MinBidSize)
	err := m.Body.Get(f)
	return f, err
}

//GetMinBidSize reads a MinBidSize from QuoteResponse.
func (m QuoteResponse) GetMinBidSize(f *field.MinBidSize) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidSize is a non-required field for QuoteResponse.
func (m QuoteResponse) BidSize() (*field.BidSize, errors.MessageRejectError) {
	f := new(field.BidSize)
	err := m.Body.Get(f)
	return f, err
}

//GetBidSize reads a BidSize from QuoteResponse.
func (m QuoteResponse) GetBidSize(f *field.BidSize) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinOfferSize is a non-required field for QuoteResponse.
func (m QuoteResponse) MinOfferSize() (*field.MinOfferSize, errors.MessageRejectError) {
	f := new(field.MinOfferSize)
	err := m.Body.Get(f)
	return f, err
}

//GetMinOfferSize reads a MinOfferSize from QuoteResponse.
func (m QuoteResponse) GetMinOfferSize(f *field.MinOfferSize) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferSize is a non-required field for QuoteResponse.
func (m QuoteResponse) OfferSize() (*field.OfferSize, errors.MessageRejectError) {
	f := new(field.OfferSize)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferSize reads a OfferSize from QuoteResponse.
func (m QuoteResponse) GetOfferSize(f *field.OfferSize) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValidUntilTime is a non-required field for QuoteResponse.
func (m QuoteResponse) ValidUntilTime() (*field.ValidUntilTime, errors.MessageRejectError) {
	f := new(field.ValidUntilTime)
	err := m.Body.Get(f)
	return f, err
}

//GetValidUntilTime reads a ValidUntilTime from QuoteResponse.
func (m QuoteResponse) GetValidUntilTime(f *field.ValidUntilTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidSpotRate is a non-required field for QuoteResponse.
func (m QuoteResponse) BidSpotRate() (*field.BidSpotRate, errors.MessageRejectError) {
	f := new(field.BidSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//GetBidSpotRate reads a BidSpotRate from QuoteResponse.
func (m QuoteResponse) GetBidSpotRate(f *field.BidSpotRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferSpotRate is a non-required field for QuoteResponse.
func (m QuoteResponse) OfferSpotRate() (*field.OfferSpotRate, errors.MessageRejectError) {
	f := new(field.OfferSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferSpotRate reads a OfferSpotRate from QuoteResponse.
func (m QuoteResponse) GetOfferSpotRate(f *field.OfferSpotRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidForwardPoints is a non-required field for QuoteResponse.
func (m QuoteResponse) BidForwardPoints() (*field.BidForwardPoints, errors.MessageRejectError) {
	f := new(field.BidForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//GetBidForwardPoints reads a BidForwardPoints from QuoteResponse.
func (m QuoteResponse) GetBidForwardPoints(f *field.BidForwardPoints) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferForwardPoints is a non-required field for QuoteResponse.
func (m QuoteResponse) OfferForwardPoints() (*field.OfferForwardPoints, errors.MessageRejectError) {
	f := new(field.OfferForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferForwardPoints reads a OfferForwardPoints from QuoteResponse.
func (m QuoteResponse) GetOfferForwardPoints(f *field.OfferForwardPoints) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MidPx is a non-required field for QuoteResponse.
func (m QuoteResponse) MidPx() (*field.MidPx, errors.MessageRejectError) {
	f := new(field.MidPx)
	err := m.Body.Get(f)
	return f, err
}

//GetMidPx reads a MidPx from QuoteResponse.
func (m QuoteResponse) GetMidPx(f *field.MidPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidYield is a non-required field for QuoteResponse.
func (m QuoteResponse) BidYield() (*field.BidYield, errors.MessageRejectError) {
	f := new(field.BidYield)
	err := m.Body.Get(f)
	return f, err
}

//GetBidYield reads a BidYield from QuoteResponse.
func (m QuoteResponse) GetBidYield(f *field.BidYield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MidYield is a non-required field for QuoteResponse.
func (m QuoteResponse) MidYield() (*field.MidYield, errors.MessageRejectError) {
	f := new(field.MidYield)
	err := m.Body.Get(f)
	return f, err
}

//GetMidYield reads a MidYield from QuoteResponse.
func (m QuoteResponse) GetMidYield(f *field.MidYield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferYield is a non-required field for QuoteResponse.
func (m QuoteResponse) OfferYield() (*field.OfferYield, errors.MessageRejectError) {
	f := new(field.OfferYield)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferYield reads a OfferYield from QuoteResponse.
func (m QuoteResponse) GetOfferYield(f *field.OfferYield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for QuoteResponse.
func (m QuoteResponse) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from QuoteResponse.
func (m QuoteResponse) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a non-required field for QuoteResponse.
func (m QuoteResponse) OrdType() (*field.OrdType, errors.MessageRejectError) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from QuoteResponse.
func (m QuoteResponse) GetOrdType(f *field.OrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BidForwardPoints2 is a non-required field for QuoteResponse.
func (m QuoteResponse) BidForwardPoints2() (*field.BidForwardPoints2, errors.MessageRejectError) {
	f := new(field.BidForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}

//GetBidForwardPoints2 reads a BidForwardPoints2 from QuoteResponse.
func (m QuoteResponse) GetBidForwardPoints2(f *field.BidForwardPoints2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OfferForwardPoints2 is a non-required field for QuoteResponse.
func (m QuoteResponse) OfferForwardPoints2() (*field.OfferForwardPoints2, errors.MessageRejectError) {
	f := new(field.OfferForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}

//GetOfferForwardPoints2 reads a OfferForwardPoints2 from QuoteResponse.
func (m QuoteResponse) GetOfferForwardPoints2(f *field.OfferForwardPoints2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrBidFxRate is a non-required field for QuoteResponse.
func (m QuoteResponse) SettlCurrBidFxRate() (*field.SettlCurrBidFxRate, errors.MessageRejectError) {
	f := new(field.SettlCurrBidFxRate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrBidFxRate reads a SettlCurrBidFxRate from QuoteResponse.
func (m QuoteResponse) GetSettlCurrBidFxRate(f *field.SettlCurrBidFxRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrOfferFxRate is a non-required field for QuoteResponse.
func (m QuoteResponse) SettlCurrOfferFxRate() (*field.SettlCurrOfferFxRate, errors.MessageRejectError) {
	f := new(field.SettlCurrOfferFxRate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrOfferFxRate reads a SettlCurrOfferFxRate from QuoteResponse.
func (m QuoteResponse) GetSettlCurrOfferFxRate(f *field.SettlCurrOfferFxRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrFxRateCalc is a non-required field for QuoteResponse.
func (m QuoteResponse) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalc, errors.MessageRejectError) {
	f := new(field.SettlCurrFxRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrFxRateCalc reads a SettlCurrFxRateCalc from QuoteResponse.
func (m QuoteResponse) GetSettlCurrFxRateCalc(f *field.SettlCurrFxRateCalc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for QuoteResponse.
func (m QuoteResponse) Commission() (*field.Commission, errors.MessageRejectError) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from QuoteResponse.
func (m QuoteResponse) GetCommission(f *field.Commission) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for QuoteResponse.
func (m QuoteResponse) CommType() (*field.CommType, errors.MessageRejectError) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from QuoteResponse.
func (m QuoteResponse) GetCommType(f *field.CommType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for QuoteResponse.
func (m QuoteResponse) CustOrderCapacity() (*field.CustOrderCapacity, errors.MessageRejectError) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from QuoteResponse.
func (m QuoteResponse) GetCustOrderCapacity(f *field.CustOrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for QuoteResponse.
func (m QuoteResponse) ExDestination() (*field.ExDestination, errors.MessageRejectError) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from QuoteResponse.
func (m QuoteResponse) GetExDestination(f *field.ExDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for QuoteResponse.
func (m QuoteResponse) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from QuoteResponse.
func (m QuoteResponse) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for QuoteResponse.
func (m QuoteResponse) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from QuoteResponse.
func (m QuoteResponse) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for QuoteResponse.
func (m QuoteResponse) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from QuoteResponse.
func (m QuoteResponse) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for QuoteResponse.
func (m QuoteResponse) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from QuoteResponse.
func (m QuoteResponse) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for QuoteResponse.
func (m QuoteResponse) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from QuoteResponse.
func (m QuoteResponse) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for QuoteResponse.
func (m QuoteResponse) Spread() (*field.Spread, errors.MessageRejectError) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from QuoteResponse.
func (m QuoteResponse) GetSpread(f *field.Spread) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for QuoteResponse.
func (m QuoteResponse) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from QuoteResponse.
func (m QuoteResponse) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for QuoteResponse.
func (m QuoteResponse) BenchmarkCurveName() (*field.BenchmarkCurveName, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from QuoteResponse.
func (m QuoteResponse) GetBenchmarkCurveName(f *field.BenchmarkCurveName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for QuoteResponse.
func (m QuoteResponse) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, errors.MessageRejectError) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from QuoteResponse.
func (m QuoteResponse) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for QuoteResponse.
func (m QuoteResponse) BenchmarkPrice() (*field.BenchmarkPrice, errors.MessageRejectError) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from QuoteResponse.
func (m QuoteResponse) GetBenchmarkPrice(f *field.BenchmarkPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for QuoteResponse.
func (m QuoteResponse) BenchmarkPriceType() (*field.BenchmarkPriceType, errors.MessageRejectError) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from QuoteResponse.
func (m QuoteResponse) GetBenchmarkPriceType(f *field.BenchmarkPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for QuoteResponse.
func (m QuoteResponse) BenchmarkSecurityID() (*field.BenchmarkSecurityID, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from QuoteResponse.
func (m QuoteResponse) GetBenchmarkSecurityID(f *field.BenchmarkSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for QuoteResponse.
func (m QuoteResponse) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from QuoteResponse.
func (m QuoteResponse) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for QuoteResponse.
func (m QuoteResponse) YieldType() (*field.YieldType, errors.MessageRejectError) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from QuoteResponse.
func (m QuoteResponse) GetYieldType(f *field.YieldType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for QuoteResponse.
func (m QuoteResponse) Yield() (*field.Yield, errors.MessageRejectError) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from QuoteResponse.
func (m QuoteResponse) GetYield(f *field.Yield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for QuoteResponse.
func (m QuoteResponse) YieldCalcDate() (*field.YieldCalcDate, errors.MessageRejectError) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from QuoteResponse.
func (m QuoteResponse) GetYieldCalcDate(f *field.YieldCalcDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for QuoteResponse.
func (m QuoteResponse) YieldRedemptionDate() (*field.YieldRedemptionDate, errors.MessageRejectError) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from QuoteResponse.
func (m QuoteResponse) GetYieldRedemptionDate(f *field.YieldRedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for QuoteResponse.
func (m QuoteResponse) YieldRedemptionPrice() (*field.YieldRedemptionPrice, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from QuoteResponse.
func (m QuoteResponse) GetYieldRedemptionPrice(f *field.YieldRedemptionPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for QuoteResponse.
func (m QuoteResponse) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from QuoteResponse.
func (m QuoteResponse) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}
