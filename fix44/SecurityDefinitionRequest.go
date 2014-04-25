package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityDefinitionRequest msg type = c.
type SecurityDefinitionRequest struct {
	message.Message
}

//SecurityDefinitionRequestBuilder builds SecurityDefinitionRequest messages.
type SecurityDefinitionRequestBuilder struct {
	message.MessageBuilder
}

//CreateSecurityDefinitionRequestBuilder returns an initialized SecurityDefinitionRequestBuilder with specified required fields.
func CreateSecurityDefinitionRequestBuilder(
	securityreqid field.SecurityReqID,
	securityrequesttype field.SecurityRequestType) SecurityDefinitionRequestBuilder {
	var builder SecurityDefinitionRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("c"))
	builder.Body.Set(securityreqid)
	builder.Body.Set(securityrequesttype)
	return builder
}

//SecurityReqID is a required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityReqID() (*field.SecurityReqID, errors.MessageRejectError) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetSecurityReqID(f *field.SecurityReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityRequestType is a required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityRequestType() (*field.SecurityRequestType, errors.MessageRejectError) {
	f := new(field.SecurityRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityRequestType reads a SecurityRequestType from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetSecurityRequestType(f *field.SecurityRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryForm is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) DeliveryForm() (*field.DeliveryForm, errors.MessageRejectError) {
	f := new(field.DeliveryForm)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryForm reads a DeliveryForm from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetDeliveryForm(f *field.DeliveryForm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PctAtRisk is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) PctAtRisk() (*field.PctAtRisk, errors.MessageRejectError) {
	f := new(field.PctAtRisk)
	err := m.Body.Get(f)
	return f, err
}

//GetPctAtRisk reads a PctAtRisk from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetPctAtRisk(f *field.PctAtRisk) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrAttrib is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NoInstrAttrib() (*field.NoInstrAttrib, errors.MessageRejectError) {
	f := new(field.NoInstrAttrib)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrAttrib reads a NoInstrAttrib from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetNoInstrAttrib(f *field.NoInstrAttrib) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpirationCycle is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) ExpirationCycle() (*field.ExpirationCycle, errors.MessageRejectError) {
	f := new(field.ExpirationCycle)
	err := m.Body.Get(f)
	return f, err
}

//GetExpirationCycle reads a ExpirationCycle from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetExpirationCycle(f *field.ExpirationCycle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, errors.MessageRejectError) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from SecurityDefinitionRequest.
func (m SecurityDefinitionRequest) GetSubscriptionRequestType(f *field.SubscriptionRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}
