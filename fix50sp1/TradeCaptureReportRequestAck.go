package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//TradeCaptureReportRequestAck msg type = AQ.
type TradeCaptureReportRequestAck struct {
	message.Message
}

//TradeCaptureReportRequestAckBuilder builds TradeCaptureReportRequestAck messages.
type TradeCaptureReportRequestAckBuilder struct {
	message.MessageBuilder
}

//CreateTradeCaptureReportRequestAckBuilder returns an initialized TradeCaptureReportRequestAckBuilder with specified required fields.
func CreateTradeCaptureReportRequestAckBuilder(
	traderequestid field.TradeRequestID,
	traderequesttype field.TradeRequestType,
	traderequestresult field.TradeRequestResult,
	traderequeststatus field.TradeRequestStatus) TradeCaptureReportRequestAckBuilder {
	var builder TradeCaptureReportRequestAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.BuildMsgType("AQ"))
	builder.Body.Set(traderequestid)
	builder.Body.Set(traderequesttype)
	builder.Body.Set(traderequestresult)
	builder.Body.Set(traderequeststatus)
	return builder
}

//TradeRequestID is a required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) TradeRequestID() (*field.TradeRequestID, errors.MessageRejectError) {
	f := new(field.TradeRequestID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeRequestID reads a TradeRequestID from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetTradeRequestID(f *field.TradeRequestID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeRequestType is a required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) TradeRequestType() (*field.TradeRequestType, errors.MessageRejectError) {
	f := new(field.TradeRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeRequestType reads a TradeRequestType from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetTradeRequestType(f *field.TradeRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SubscriptionRequestType() (*field.SubscriptionRequestType, errors.MessageRejectError) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSubscriptionRequestType(f *field.SubscriptionRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNumTradeReports is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) TotNumTradeReports() (*field.TotNumTradeReports, errors.MessageRejectError) {
	f := new(field.TotNumTradeReports)
	err := m.Body.Get(f)
	return f, err
}

//GetTotNumTradeReports reads a TotNumTradeReports from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetTotNumTradeReports(f *field.TotNumTradeReports) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeRequestResult is a required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) TradeRequestResult() (*field.TradeRequestResult, errors.MessageRejectError) {
	f := new(field.TradeRequestResult)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeRequestResult reads a TradeRequestResult from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetTradeRequestResult(f *field.TradeRequestResult) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeRequestStatus is a required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) TradeRequestStatus() (*field.TradeRequestStatus, errors.MessageRejectError) {
	f := new(field.TradeRequestStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeRequestStatus reads a TradeRequestStatus from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetTradeRequestStatus(f *field.TradeRequestStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityGroup() (*field.SecurityGroup, errors.MessageRejectError) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSecurityGroup(f *field.SecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) UnitOfMeasureQty() (*field.UnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetUnitOfMeasureQty(f *field.UnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityXMLLen() (*field.SecurityXMLLen, errors.MessageRejectError) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSecurityXMLLen(f *field.SecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityXML() (*field.SecurityXML, errors.MessageRejectError) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSecurityXML(f *field.SecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecurityXMLSchema() (*field.SecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSecurityXMLSchema(f *field.SecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) ProductComplex() (*field.ProductComplex, errors.MessageRejectError) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetProductComplex(f *field.ProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SettlMethod() (*field.SettlMethod, errors.MessageRejectError) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSettlMethod(f *field.SettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) ExerciseStyle() (*field.ExerciseStyle, errors.MessageRejectError) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetExerciseStyle(f *field.ExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayAmount is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) OptPayAmount() (*field.OptPayAmount, errors.MessageRejectError) {
	f := new(field.OptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayAmount reads a OptPayAmount from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetOptPayAmount(f *field.OptPayAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) PriceQuoteMethod() (*field.PriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetPriceQuoteMethod(f *field.PriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) ListMethod() (*field.ListMethod, errors.MessageRejectError) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetListMethod(f *field.ListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) CapPrice() (*field.CapPrice, errors.MessageRejectError) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetCapPrice(f *field.CapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) FloorPrice() (*field.FloorPrice, errors.MessageRejectError) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetFloorPrice(f *field.FloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) FlexibleIndicator() (*field.FlexibleIndicator, errors.MessageRejectError) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetFlexibleIndicator(f *field.FlexibleIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FuturesValuationMethod is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) FuturesValuationMethod() (*field.FuturesValuationMethod, errors.MessageRejectError) {
	f := new(field.FuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetFuturesValuationMethod reads a FuturesValuationMethod from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetFuturesValuationMethod(f *field.FuturesValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegReportingType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) MultiLegReportingType() (*field.MultiLegReportingType, errors.MessageRejectError) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegReportingType reads a MultiLegReportingType from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetMultiLegReportingType(f *field.MultiLegReportingType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseTransportType is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) ResponseTransportType() (*field.ResponseTransportType, errors.MessageRejectError) {
	f := new(field.ResponseTransportType)
	err := m.Body.Get(f)
	return f, err
}

//GetResponseTransportType reads a ResponseTransportType from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetResponseTransportType(f *field.ResponseTransportType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseDestination is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) ResponseDestination() (*field.ResponseDestination, errors.MessageRejectError) {
	f := new(field.ResponseDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetResponseDestination reads a ResponseDestination from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetResponseDestination(f *field.ResponseDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageEventSource is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) MessageEventSource() (*field.MessageEventSource, errors.MessageRejectError) {
	f := new(field.MessageEventSource)
	err := m.Body.Get(f)
	return f, err
}

//GetMessageEventSource reads a MessageEventSource from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetMessageEventSource(f *field.MessageEventSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeID is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) TradeID() (*field.TradeID, errors.MessageRejectError) {
	f := new(field.TradeID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeID reads a TradeID from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetTradeID(f *field.TradeID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTradeID is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecondaryTradeID() (*field.SecondaryTradeID, errors.MessageRejectError) {
	f := new(field.SecondaryTradeID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTradeID reads a SecondaryTradeID from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSecondaryTradeID(f *field.SecondaryTradeID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FirmTradeID is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) FirmTradeID() (*field.FirmTradeID, errors.MessageRejectError) {
	f := new(field.FirmTradeID)
	err := m.Body.Get(f)
	return f, err
}

//GetFirmTradeID reads a FirmTradeID from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetFirmTradeID(f *field.FirmTradeID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryFirmTradeID is a non-required field for TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) SecondaryFirmTradeID() (*field.SecondaryFirmTradeID, errors.MessageRejectError) {
	f := new(field.SecondaryFirmTradeID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryFirmTradeID reads a SecondaryFirmTradeID from TradeCaptureReportRequestAck.
func (m TradeCaptureReportRequestAck) GetSecondaryFirmTradeID(f *field.SecondaryFirmTradeID) errors.MessageRejectError {
	return m.Body.Get(f)
}
