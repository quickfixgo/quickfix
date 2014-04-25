package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityDefinitionUpdateReport msg type = BP.
type SecurityDefinitionUpdateReport struct {
	message.Message
}

//SecurityDefinitionUpdateReportBuilder builds SecurityDefinitionUpdateReport messages.
type SecurityDefinitionUpdateReportBuilder struct {
	message.MessageBuilder
}

//CreateSecurityDefinitionUpdateReportBuilder returns an initialized SecurityDefinitionUpdateReportBuilder with specified required fields.
func CreateSecurityDefinitionUpdateReportBuilder() SecurityDefinitionUpdateReportBuilder {
	var builder SecurityDefinitionUpdateReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("BP"))
	return builder
}

//SecurityReportID is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SecurityReportID() (*field.SecurityReportID, errors.MessageRejectError) {
	f := new(field.SecurityReportID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReportID reads a SecurityReportID from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSecurityReportID(f *field.SecurityReportID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityReqID is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SecurityReqID() (*field.SecurityReqID, errors.MessageRejectError) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSecurityReqID(f *field.SecurityReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SecurityResponseID() (*field.SecurityResponseID, errors.MessageRejectError) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSecurityResponseID(f *field.SecurityResponseID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseType is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SecurityResponseType() (*field.SecurityResponseType, errors.MessageRejectError) {
	f := new(field.SecurityResponseType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseType reads a SecurityResponseType from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSecurityResponseType(f *field.SecurityResponseType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityUpdateAction is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SecurityUpdateAction() (*field.SecurityUpdateAction, errors.MessageRejectError) {
	f := new(field.SecurityUpdateAction)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityUpdateAction reads a SecurityUpdateAction from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSecurityUpdateAction(f *field.SecurityUpdateAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CorporateAction is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) CorporateAction() (*field.CorporateAction, errors.MessageRejectError) {
	f := new(field.CorporateAction)
	err := m.Body.Get(f)
	return f, err
}

//GetCorporateAction reads a CorporateAction from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetCorporateAction(f *field.CorporateAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SecurityGroup() (*field.SecurityGroup, errors.MessageRejectError) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSecurityGroup(f *field.SecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) UnitOfMeasureQty() (*field.UnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetUnitOfMeasureQty(f *field.UnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SecurityXMLLen() (*field.SecurityXMLLen, errors.MessageRejectError) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSecurityXMLLen(f *field.SecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SecurityXML() (*field.SecurityXML, errors.MessageRejectError) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSecurityXML(f *field.SecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SecurityXMLSchema() (*field.SecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSecurityXMLSchema(f *field.SecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) ProductComplex() (*field.ProductComplex, errors.MessageRejectError) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetProductComplex(f *field.ProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) SettlMethod() (*field.SettlMethod, errors.MessageRejectError) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSettlMethod(f *field.SettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) ExerciseStyle() (*field.ExerciseStyle, errors.MessageRejectError) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetExerciseStyle(f *field.ExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) OptPayoutAmount() (*field.OptPayoutAmount, errors.MessageRejectError) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetOptPayoutAmount(f *field.OptPayoutAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) PriceQuoteMethod() (*field.PriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetPriceQuoteMethod(f *field.PriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) ListMethod() (*field.ListMethod, errors.MessageRejectError) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetListMethod(f *field.ListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) CapPrice() (*field.CapPrice, errors.MessageRejectError) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetCapPrice(f *field.CapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) FloorPrice() (*field.FloorPrice, errors.MessageRejectError) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetFloorPrice(f *field.FloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) FlexibleIndicator() (*field.FlexibleIndicator, errors.MessageRejectError) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetFlexibleIndicator(f *field.FlexibleIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) ValuationMethod() (*field.ValuationMethod, errors.MessageRejectError) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetValuationMethod(f *field.ValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) ContractMultiplierUnit() (*field.ContractMultiplierUnit, errors.MessageRejectError) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetContractMultiplierUnit(f *field.ContractMultiplierUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) FlowScheduleType() (*field.FlowScheduleType, errors.MessageRejectError) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetFlowScheduleType(f *field.FlowScheduleType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) RestructuringType() (*field.RestructuringType, errors.MessageRejectError) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetRestructuringType(f *field.RestructuringType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) Seniority() (*field.Seniority, errors.MessageRejectError) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSeniority(f *field.Seniority) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) AttachmentPoint() (*field.AttachmentPoint, errors.MessageRejectError) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetAttachmentPoint(f *field.AttachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) DetachmentPoint() (*field.DetachmentPoint, errors.MessageRejectError) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetDetachmentPoint(f *field.DetachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecision) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) OptPayoutType() (*field.OptPayoutType, errors.MessageRejectError) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetOptPayoutType(f *field.OptPayoutType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) NoComplexEvents() (*field.NoComplexEvents, errors.MessageRejectError) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetNoComplexEvents(f *field.NoComplexEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryForm is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) DeliveryForm() (*field.DeliveryForm, errors.MessageRejectError) {
	f := new(field.DeliveryForm)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryForm reads a DeliveryForm from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetDeliveryForm(f *field.DeliveryForm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PctAtRisk is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) PctAtRisk() (*field.PctAtRisk, errors.MessageRejectError) {
	f := new(field.PctAtRisk)
	err := m.Body.Get(f)
	return f, err
}

//GetPctAtRisk reads a PctAtRisk from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetPctAtRisk(f *field.PctAtRisk) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrAttrib is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) NoInstrAttrib() (*field.NoInstrAttrib, errors.MessageRejectError) {
	f := new(field.NoInstrAttrib)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrAttrib reads a NoInstrAttrib from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetNoInstrAttrib(f *field.NoInstrAttrib) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) NoStipulations() (*field.NoStipulations, errors.MessageRejectError) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetNoStipulations(f *field.NoStipulations) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) Spread() (*field.Spread, errors.MessageRejectError) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetSpread(f *field.Spread) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) BenchmarkCurveName() (*field.BenchmarkCurveName, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetBenchmarkCurveName(f *field.BenchmarkCurveName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, errors.MessageRejectError) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) BenchmarkPrice() (*field.BenchmarkPrice, errors.MessageRejectError) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetBenchmarkPrice(f *field.BenchmarkPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) BenchmarkPriceType() (*field.BenchmarkPriceType, errors.MessageRejectError) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetBenchmarkPriceType(f *field.BenchmarkPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) BenchmarkSecurityID() (*field.BenchmarkSecurityID, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetBenchmarkSecurityID(f *field.BenchmarkSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) YieldType() (*field.YieldType, errors.MessageRejectError) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetYieldType(f *field.YieldType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) Yield() (*field.Yield, errors.MessageRejectError) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetYield(f *field.Yield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) YieldCalcDate() (*field.YieldCalcDate, errors.MessageRejectError) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetYieldCalcDate(f *field.YieldCalcDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) YieldRedemptionDate() (*field.YieldRedemptionDate, errors.MessageRejectError) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetYieldRedemptionDate(f *field.YieldRedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) YieldRedemptionPrice() (*field.YieldRedemptionPrice, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetYieldRedemptionPrice(f *field.YieldRedemptionPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMarketSegments is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) NoMarketSegments() (*field.NoMarketSegments, errors.MessageRejectError) {
	f := new(field.NoMarketSegments)
	err := m.Body.Get(f)
	return f, err
}

//GetNoMarketSegments reads a NoMarketSegments from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetNoMarketSegments(f *field.NoMarketSegments) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) ApplID() (*field.ApplID, errors.MessageRejectError) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetApplID(f *field.ApplID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) ApplSeqNum() (*field.ApplSeqNum, errors.MessageRejectError) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetApplSeqNum(f *field.ApplSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) ApplLastSeqNum() (*field.ApplLastSeqNum, errors.MessageRejectError) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetApplLastSeqNum(f *field.ApplLastSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) ApplResendFlag() (*field.ApplResendFlag, errors.MessageRejectError) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetApplResendFlag(f *field.ApplResendFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SecurityDefinitionUpdateReport.
func (m SecurityDefinitionUpdateReport) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}
