package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderMassCancelReport msg type = r.
type OrderMassCancelReport struct {
	message.Message
}

//OrderMassCancelReportBuilder builds OrderMassCancelReport messages.
type OrderMassCancelReportBuilder struct {
	message.MessageBuilder
}

//CreateOrderMassCancelReportBuilder returns an initialized OrderMassCancelReportBuilder with specified required fields.
func CreateOrderMassCancelReportBuilder(
	orderid field.OrderID,
	masscancelrequesttype field.MassCancelRequestType,
	masscancelresponse field.MassCancelResponse) OrderMassCancelReportBuilder {
	var builder OrderMassCancelReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(orderid)
	builder.Body.Set(masscancelrequesttype)
	builder.Body.Set(masscancelresponse)
	return builder
}

//ClOrdID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderMassCancelReport.
func (m OrderMassCancelReport) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecondaryClOrdID() (*field.SecondaryClOrdID, errors.MessageRejectError) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from OrderMassCancelReport.
func (m OrderMassCancelReport) GetSecondaryClOrdID(f *field.SecondaryClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderID is a required field for OrderMassCancelReport.
func (m OrderMassCancelReport) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from OrderMassCancelReport.
func (m OrderMassCancelReport) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecondaryOrderID() (*field.SecondaryOrderID, errors.MessageRejectError) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from OrderMassCancelReport.
func (m OrderMassCancelReport) GetSecondaryOrderID(f *field.SecondaryOrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MassCancelRequestType is a required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MassCancelRequestType() (*field.MassCancelRequestType, errors.MessageRejectError) {
	f := new(field.MassCancelRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetMassCancelRequestType reads a MassCancelRequestType from OrderMassCancelReport.
func (m OrderMassCancelReport) GetMassCancelRequestType(f *field.MassCancelRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MassCancelResponse is a required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MassCancelResponse() (*field.MassCancelResponse, errors.MessageRejectError) {
	f := new(field.MassCancelResponse)
	err := m.Body.Get(f)
	return f, err
}

//GetMassCancelResponse reads a MassCancelResponse from OrderMassCancelReport.
func (m OrderMassCancelReport) GetMassCancelResponse(f *field.MassCancelResponse) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MassCancelRejectReason is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MassCancelRejectReason() (*field.MassCancelRejectReason, errors.MessageRejectError) {
	f := new(field.MassCancelRejectReason)
	err := m.Body.Get(f)
	return f, err
}

//GetMassCancelRejectReason reads a MassCancelRejectReason from OrderMassCancelReport.
func (m OrderMassCancelReport) GetMassCancelRejectReason(f *field.MassCancelRejectReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalAffectedOrders is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) TotalAffectedOrders() (*field.TotalAffectedOrders, errors.MessageRejectError) {
	f := new(field.TotalAffectedOrders)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalAffectedOrders reads a TotalAffectedOrders from OrderMassCancelReport.
func (m OrderMassCancelReport) GetTotalAffectedOrders(f *field.TotalAffectedOrders) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAffectedOrders is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoAffectedOrders() (*field.NoAffectedOrders, errors.MessageRejectError) {
	f := new(field.NoAffectedOrders)
	err := m.Body.Get(f)
	return f, err
}

//GetNoAffectedOrders reads a NoAffectedOrders from OrderMassCancelReport.
func (m OrderMassCancelReport) GetNoAffectedOrders(f *field.NoAffectedOrders) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from OrderMassCancelReport.
func (m OrderMassCancelReport) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from OrderMassCancelReport.
func (m OrderMassCancelReport) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from OrderMassCancelReport.
func (m OrderMassCancelReport) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from OrderMassCancelReport.
func (m OrderMassCancelReport) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from OrderMassCancelReport.
func (m OrderMassCancelReport) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from OrderMassCancelReport.
func (m OrderMassCancelReport) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from OrderMassCancelReport.
func (m OrderMassCancelReport) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from OrderMassCancelReport.
func (m OrderMassCancelReport) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from OrderMassCancelReport.
func (m OrderMassCancelReport) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from OrderMassCancelReport.
func (m OrderMassCancelReport) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from OrderMassCancelReport.
func (m OrderMassCancelReport) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from OrderMassCancelReport.
func (m OrderMassCancelReport) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from OrderMassCancelReport.
func (m OrderMassCancelReport) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from OrderMassCancelReport.
func (m OrderMassCancelReport) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from OrderMassCancelReport.
func (m OrderMassCancelReport) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from OrderMassCancelReport.
func (m OrderMassCancelReport) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from OrderMassCancelReport.
func (m OrderMassCancelReport) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from OrderMassCancelReport.
func (m OrderMassCancelReport) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from OrderMassCancelReport.
func (m OrderMassCancelReport) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from OrderMassCancelReport.
func (m OrderMassCancelReport) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from OrderMassCancelReport.
func (m OrderMassCancelReport) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from OrderMassCancelReport.
func (m OrderMassCancelReport) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from OrderMassCancelReport.
func (m OrderMassCancelReport) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from OrderMassCancelReport.
func (m OrderMassCancelReport) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from OrderMassCancelReport.
func (m OrderMassCancelReport) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from OrderMassCancelReport.
func (m OrderMassCancelReport) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from OrderMassCancelReport.
func (m OrderMassCancelReport) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from OrderMassCancelReport.
func (m OrderMassCancelReport) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from OrderMassCancelReport.
func (m OrderMassCancelReport) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from OrderMassCancelReport.
func (m OrderMassCancelReport) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from OrderMassCancelReport.
func (m OrderMassCancelReport) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from OrderMassCancelReport.
func (m OrderMassCancelReport) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from OrderMassCancelReport.
func (m OrderMassCancelReport) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from OrderMassCancelReport.
func (m OrderMassCancelReport) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from OrderMassCancelReport.
func (m OrderMassCancelReport) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from OrderMassCancelReport.
func (m OrderMassCancelReport) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from OrderMassCancelReport.
func (m OrderMassCancelReport) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from OrderMassCancelReport.
func (m OrderMassCancelReport) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from OrderMassCancelReport.
func (m OrderMassCancelReport) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from OrderMassCancelReport.
func (m OrderMassCancelReport) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from OrderMassCancelReport.
func (m OrderMassCancelReport) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from OrderMassCancelReport.
func (m OrderMassCancelReport) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from OrderMassCancelReport.
func (m OrderMassCancelReport) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from OrderMassCancelReport.
func (m OrderMassCancelReport) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from OrderMassCancelReport.
func (m OrderMassCancelReport) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from OrderMassCancelReport.
func (m OrderMassCancelReport) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from OrderMassCancelReport.
func (m OrderMassCancelReport) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from OrderMassCancelReport.
func (m OrderMassCancelReport) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from OrderMassCancelReport.
func (m OrderMassCancelReport) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from OrderMassCancelReport.
func (m OrderMassCancelReport) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from OrderMassCancelReport.
func (m OrderMassCancelReport) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from OrderMassCancelReport.
func (m OrderMassCancelReport) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from OrderMassCancelReport.
func (m OrderMassCancelReport) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from OrderMassCancelReport.
func (m OrderMassCancelReport) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from OrderMassCancelReport.
func (m OrderMassCancelReport) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbol is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSymbol() (*field.UnderlyingSymbol, errors.MessageRejectError) {
	f := new(field.UnderlyingSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbol reads a UnderlyingSymbol from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingSymbol(f *field.UnderlyingSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbolSfx is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfx, errors.MessageRejectError) {
	f := new(field.UnderlyingSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbolSfx reads a UnderlyingSymbolSfx from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingSymbolSfx(f *field.UnderlyingSymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecurityID() (*field.UnderlyingSecurityID, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityID reads a UnderlyingSecurityID from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingSecurityID(f *field.UnderlyingSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityIDSource is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSource, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityIDSource reads a UnderlyingSecurityIDSource from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingSecurityIDSource(f *field.UnderlyingSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingSecurityAltID is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoUnderlyingSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingSecurityAltID reads a NoUnderlyingSecurityAltID from OrderMassCancelReport.
func (m OrderMassCancelReport) GetNoUnderlyingSecurityAltID(f *field.NoUnderlyingSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingProduct is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingProduct() (*field.UnderlyingProduct, errors.MessageRejectError) {
	f := new(field.UnderlyingProduct)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingProduct reads a UnderlyingProduct from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingProduct(f *field.UnderlyingProduct) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCFICode is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCFICode() (*field.UnderlyingCFICode, errors.MessageRejectError) {
	f := new(field.UnderlyingCFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCFICode reads a UnderlyingCFICode from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingCFICode(f *field.UnderlyingCFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecurityType() (*field.UnderlyingSecurityType, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityType reads a UnderlyingSecurityType from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingSecurityType(f *field.UnderlyingSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecuritySubType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubType, errors.MessageRejectError) {
	f := new(field.UnderlyingSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecuritySubType reads a UnderlyingSecuritySubType from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingSecuritySubType(f *field.UnderlyingSecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityMonthYear is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYear, errors.MessageRejectError) {
	f := new(field.UnderlyingMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityMonthYear reads a UnderlyingMaturityMonthYear from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingMaturityMonthYear(f *field.UnderlyingMaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingMaturityDate() (*field.UnderlyingMaturityDate, errors.MessageRejectError) {
	f := new(field.UnderlyingMaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityDate reads a UnderlyingMaturityDate from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingMaturityDate(f *field.UnderlyingMaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponPaymentDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDate, errors.MessageRejectError) {
	f := new(field.UnderlyingCouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponPaymentDate reads a UnderlyingCouponPaymentDate from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingCouponPaymentDate(f *field.UnderlyingCouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssueDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingIssueDate() (*field.UnderlyingIssueDate, errors.MessageRejectError) {
	f := new(field.UnderlyingIssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssueDate reads a UnderlyingIssueDate from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingIssueDate(f *field.UnderlyingIssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepoCollateralSecurityType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.UnderlyingRepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepoCollateralSecurityType reads a UnderlyingRepoCollateralSecurityType from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingRepoCollateralSecurityType(f *field.UnderlyingRepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseTerm is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTerm, errors.MessageRejectError) {
	f := new(field.UnderlyingRepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseTerm reads a UnderlyingRepurchaseTerm from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingRepurchaseTerm(f *field.UnderlyingRepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseRate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRate, errors.MessageRejectError) {
	f := new(field.UnderlyingRepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseRate reads a UnderlyingRepurchaseRate from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingRepurchaseRate(f *field.UnderlyingRepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFactor is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingFactor() (*field.UnderlyingFactor, errors.MessageRejectError) {
	f := new(field.UnderlyingFactor)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFactor reads a UnderlyingFactor from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingFactor(f *field.UnderlyingFactor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCreditRating is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCreditRating() (*field.UnderlyingCreditRating, errors.MessageRejectError) {
	f := new(field.UnderlyingCreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCreditRating reads a UnderlyingCreditRating from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingCreditRating(f *field.UnderlyingCreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingInstrRegistry is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistry, errors.MessageRejectError) {
	f := new(field.UnderlyingInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingInstrRegistry reads a UnderlyingInstrRegistry from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingInstrRegistry(f *field.UnderlyingInstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCountryOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssue, errors.MessageRejectError) {
	f := new(field.UnderlyingCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCountryOfIssue reads a UnderlyingCountryOfIssue from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingCountryOfIssue(f *field.UnderlyingCountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.UnderlyingStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStateOrProvinceOfIssue reads a UnderlyingStateOrProvinceOfIssue from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingStateOrProvinceOfIssue(f *field.UnderlyingStateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLocaleOfIssue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssue, errors.MessageRejectError) {
	f := new(field.UnderlyingLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLocaleOfIssue reads a UnderlyingLocaleOfIssue from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingLocaleOfIssue(f *field.UnderlyingLocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRedemptionDate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDate, errors.MessageRejectError) {
	f := new(field.UnderlyingRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRedemptionDate reads a UnderlyingRedemptionDate from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingRedemptionDate(f *field.UnderlyingRedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikePrice is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingStrikePrice() (*field.UnderlyingStrikePrice, errors.MessageRejectError) {
	f := new(field.UnderlyingStrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikePrice reads a UnderlyingStrikePrice from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingStrikePrice(f *field.UnderlyingStrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikeCurrency is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrency, errors.MessageRejectError) {
	f := new(field.UnderlyingStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikeCurrency reads a UnderlyingStrikeCurrency from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingStrikeCurrency(f *field.UnderlyingStrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingOptAttribute is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingOptAttribute() (*field.UnderlyingOptAttribute, errors.MessageRejectError) {
	f := new(field.UnderlyingOptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingOptAttribute reads a UnderlyingOptAttribute from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingOptAttribute(f *field.UnderlyingOptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingContractMultiplier is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplier, errors.MessageRejectError) {
	f := new(field.UnderlyingContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingContractMultiplier reads a UnderlyingContractMultiplier from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingContractMultiplier(f *field.UnderlyingContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponRate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCouponRate() (*field.UnderlyingCouponRate, errors.MessageRejectError) {
	f := new(field.UnderlyingCouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponRate reads a UnderlyingCouponRate from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingCouponRate(f *field.UnderlyingCouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityExchange is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchange, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityExchange reads a UnderlyingSecurityExchange from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingSecurityExchange(f *field.UnderlyingSecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssuer is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingIssuer() (*field.UnderlyingIssuer, errors.MessageRejectError) {
	f := new(field.UnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssuer reads a UnderlyingIssuer from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingIssuer(f *field.UnderlyingIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuerLen is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuerLen reads a EncodedUnderlyingIssuerLen from OrderMassCancelReport.
func (m OrderMassCancelReport) GetEncodedUnderlyingIssuerLen(f *field.EncodedUnderlyingIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuer is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuer, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuer reads a EncodedUnderlyingIssuer from OrderMassCancelReport.
func (m OrderMassCancelReport) GetEncodedUnderlyingIssuer(f *field.EncodedUnderlyingIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityDesc is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDesc, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityDesc reads a UnderlyingSecurityDesc from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingSecurityDesc(f *field.UnderlyingSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDescLen is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDescLen reads a EncodedUnderlyingSecurityDescLen from OrderMassCancelReport.
func (m OrderMassCancelReport) GetEncodedUnderlyingSecurityDescLen(f *field.EncodedUnderlyingSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDesc is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDesc reads a EncodedUnderlyingSecurityDesc from OrderMassCancelReport.
func (m OrderMassCancelReport) GetEncodedUnderlyingSecurityDesc(f *field.EncodedUnderlyingSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPProgram is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCPProgram() (*field.UnderlyingCPProgram, errors.MessageRejectError) {
	f := new(field.UnderlyingCPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPProgram reads a UnderlyingCPProgram from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingCPProgram(f *field.UnderlyingCPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPRegType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCPRegType() (*field.UnderlyingCPRegType, errors.MessageRejectError) {
	f := new(field.UnderlyingCPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPRegType reads a UnderlyingCPRegType from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingCPRegType(f *field.UnderlyingCPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrency is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCurrency() (*field.UnderlyingCurrency, errors.MessageRejectError) {
	f := new(field.UnderlyingCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrency reads a UnderlyingCurrency from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingCurrency(f *field.UnderlyingCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingQty is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingQty() (*field.UnderlyingQty, errors.MessageRejectError) {
	f := new(field.UnderlyingQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingQty reads a UnderlyingQty from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingQty(f *field.UnderlyingQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPx is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingPx() (*field.UnderlyingPx, errors.MessageRejectError) {
	f := new(field.UnderlyingPx)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPx reads a UnderlyingPx from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingPx(f *field.UnderlyingPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingDirtyPrice is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPrice, errors.MessageRejectError) {
	f := new(field.UnderlyingDirtyPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingDirtyPrice reads a UnderlyingDirtyPrice from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingDirtyPrice(f *field.UnderlyingDirtyPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndPrice is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingEndPrice() (*field.UnderlyingEndPrice, errors.MessageRejectError) {
	f := new(field.UnderlyingEndPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndPrice reads a UnderlyingEndPrice from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingEndPrice(f *field.UnderlyingEndPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStartValue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingStartValue() (*field.UnderlyingStartValue, errors.MessageRejectError) {
	f := new(field.UnderlyingStartValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStartValue reads a UnderlyingStartValue from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingStartValue(f *field.UnderlyingStartValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrentValue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCurrentValue() (*field.UnderlyingCurrentValue, errors.MessageRejectError) {
	f := new(field.UnderlyingCurrentValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrentValue reads a UnderlyingCurrentValue from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingCurrentValue(f *field.UnderlyingCurrentValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndValue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingEndValue() (*field.UnderlyingEndValue, errors.MessageRejectError) {
	f := new(field.UnderlyingEndValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndValue reads a UnderlyingEndValue from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingEndValue(f *field.UnderlyingEndValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingStips is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoUnderlyingStips() (*field.NoUnderlyingStips, errors.MessageRejectError) {
	f := new(field.NoUnderlyingStips)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingStips reads a NoUnderlyingStips from OrderMassCancelReport.
func (m OrderMassCancelReport) GetNoUnderlyingStips(f *field.NoUnderlyingStips) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAllocationPercent is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercent, errors.MessageRejectError) {
	f := new(field.UnderlyingAllocationPercent)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAllocationPercent reads a UnderlyingAllocationPercent from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingAllocationPercent(f *field.UnderlyingAllocationPercent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlementType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSettlementType() (*field.UnderlyingSettlementType, errors.MessageRejectError) {
	f := new(field.UnderlyingSettlementType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlementType reads a UnderlyingSettlementType from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingSettlementType(f *field.UnderlyingSettlementType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashAmount is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCashAmount() (*field.UnderlyingCashAmount, errors.MessageRejectError) {
	f := new(field.UnderlyingCashAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashAmount reads a UnderlyingCashAmount from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingCashAmount(f *field.UnderlyingCashAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashType is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCashType() (*field.UnderlyingCashType, errors.MessageRejectError) {
	f := new(field.UnderlyingCashType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashType reads a UnderlyingCashType from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingCashType(f *field.UnderlyingCashType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasure is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnderlyingUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasure reads a UnderlyingUnitOfMeasure from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingUnitOfMeasure(f *field.UnderlyingUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTimeUnit is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingTimeUnit() (*field.UnderlyingTimeUnit, errors.MessageRejectError) {
	f := new(field.UnderlyingTimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTimeUnit reads a UnderlyingTimeUnit from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingTimeUnit(f *field.UnderlyingTimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCapValue is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingCapValue() (*field.UnderlyingCapValue, errors.MessageRejectError) {
	f := new(field.UnderlyingCapValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCapValue reads a UnderlyingCapValue from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingCapValue(f *field.UnderlyingCapValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUndlyInstrumentParties is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoUndlyInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUndlyInstrumentParties reads a NoUndlyInstrumentParties from OrderMassCancelReport.
func (m OrderMassCancelReport) GetNoUndlyInstrumentParties(f *field.NoUndlyInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlMethod is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingSettlMethod() (*field.UnderlyingSettlMethod, errors.MessageRejectError) {
	f := new(field.UnderlyingSettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlMethod reads a UnderlyingSettlMethod from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingSettlMethod(f *field.UnderlyingSettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAdjustedQuantity is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantity, errors.MessageRejectError) {
	f := new(field.UnderlyingAdjustedQuantity)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAdjustedQuantity reads a UnderlyingAdjustedQuantity from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingAdjustedQuantity(f *field.UnderlyingAdjustedQuantity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRate is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingFXRate() (*field.UnderlyingFXRate, errors.MessageRejectError) {
	f := new(field.UnderlyingFXRate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRate reads a UnderlyingFXRate from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingFXRate(f *field.UnderlyingFXRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRateCalc is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalc, errors.MessageRejectError) {
	f := new(field.UnderlyingFXRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRateCalc reads a UnderlyingFXRateCalc from OrderMassCancelReport.
func (m OrderMassCancelReport) GetUnderlyingFXRateCalc(f *field.UnderlyingFXRateCalc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from OrderMassCancelReport.
func (m OrderMassCancelReport) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from OrderMassCancelReport.
func (m OrderMassCancelReport) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from OrderMassCancelReport.
func (m OrderMassCancelReport) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from OrderMassCancelReport.
func (m OrderMassCancelReport) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from OrderMassCancelReport.
func (m OrderMassCancelReport) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for OrderMassCancelReport.
func (m OrderMassCancelReport) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from OrderMassCancelReport.
func (m OrderMassCancelReport) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}
