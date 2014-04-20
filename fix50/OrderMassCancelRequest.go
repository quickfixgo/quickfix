package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderMassCancelRequest msg type = q.
type OrderMassCancelRequest struct {
	message.Message
}

//OrderMassCancelRequestBuilder builds OrderMassCancelRequest messages.
type OrderMassCancelRequestBuilder struct {
	message.MessageBuilder
}

//CreateOrderMassCancelRequestBuilder returns an initialized OrderMassCancelRequestBuilder with specified required fields.
func CreateOrderMassCancelRequestBuilder(
	clordid field.ClOrdID,
	masscancelrequesttype field.MassCancelRequestType,
	transacttime field.TransactTime) OrderMassCancelRequestBuilder {
	var builder OrderMassCancelRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(clordid)
	builder.Body.Set(masscancelrequesttype)
	builder.Body.Set(transacttime)
	return builder
}

//ClOrdID is a required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecondaryClOrdID() (*field.SecondaryClOrdID, errors.MessageRejectError) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecondaryClOrdID(f *field.SecondaryClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MassCancelRequestType is a required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) MassCancelRequestType() (*field.MassCancelRequestType, errors.MessageRejectError) {
	f := new(field.MassCancelRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetMassCancelRequestType reads a MassCancelRequestType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetMassCancelRequestType(f *field.MassCancelRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbol is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSymbol() (*field.UnderlyingSymbol, errors.MessageRejectError) {
	f := new(field.UnderlyingSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbol reads a UnderlyingSymbol from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSymbol(f *field.UnderlyingSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSymbolSfx is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSymbolSfx() (*field.UnderlyingSymbolSfx, errors.MessageRejectError) {
	f := new(field.UnderlyingSymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSymbolSfx reads a UnderlyingSymbolSfx from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSymbolSfx(f *field.UnderlyingSymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityID is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSecurityID() (*field.UnderlyingSecurityID, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityID reads a UnderlyingSecurityID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSecurityID(f *field.UnderlyingSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityIDSource is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSecurityIDSource() (*field.UnderlyingSecurityIDSource, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityIDSource reads a UnderlyingSecurityIDSource from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSecurityIDSource(f *field.UnderlyingSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingSecurityAltID is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) NoUnderlyingSecurityAltID() (*field.NoUnderlyingSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoUnderlyingSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingSecurityAltID reads a NoUnderlyingSecurityAltID from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetNoUnderlyingSecurityAltID(f *field.NoUnderlyingSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingProduct is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingProduct() (*field.UnderlyingProduct, errors.MessageRejectError) {
	f := new(field.UnderlyingProduct)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingProduct reads a UnderlyingProduct from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingProduct(f *field.UnderlyingProduct) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCFICode is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCFICode() (*field.UnderlyingCFICode, errors.MessageRejectError) {
	f := new(field.UnderlyingCFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCFICode reads a UnderlyingCFICode from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCFICode(f *field.UnderlyingCFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSecurityType() (*field.UnderlyingSecurityType, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityType reads a UnderlyingSecurityType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSecurityType(f *field.UnderlyingSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecuritySubType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSecuritySubType() (*field.UnderlyingSecuritySubType, errors.MessageRejectError) {
	f := new(field.UnderlyingSecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecuritySubType reads a UnderlyingSecuritySubType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSecuritySubType(f *field.UnderlyingSecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityMonthYear is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingMaturityMonthYear() (*field.UnderlyingMaturityMonthYear, errors.MessageRejectError) {
	f := new(field.UnderlyingMaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityMonthYear reads a UnderlyingMaturityMonthYear from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingMaturityMonthYear(f *field.UnderlyingMaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingMaturityDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingMaturityDate() (*field.UnderlyingMaturityDate, errors.MessageRejectError) {
	f := new(field.UnderlyingMaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingMaturityDate reads a UnderlyingMaturityDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingMaturityDate(f *field.UnderlyingMaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponPaymentDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCouponPaymentDate() (*field.UnderlyingCouponPaymentDate, errors.MessageRejectError) {
	f := new(field.UnderlyingCouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponPaymentDate reads a UnderlyingCouponPaymentDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCouponPaymentDate(f *field.UnderlyingCouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssueDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingIssueDate() (*field.UnderlyingIssueDate, errors.MessageRejectError) {
	f := new(field.UnderlyingIssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssueDate reads a UnderlyingIssueDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingIssueDate(f *field.UnderlyingIssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepoCollateralSecurityType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingRepoCollateralSecurityType() (*field.UnderlyingRepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.UnderlyingRepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepoCollateralSecurityType reads a UnderlyingRepoCollateralSecurityType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingRepoCollateralSecurityType(f *field.UnderlyingRepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseTerm is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingRepurchaseTerm() (*field.UnderlyingRepurchaseTerm, errors.MessageRejectError) {
	f := new(field.UnderlyingRepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseTerm reads a UnderlyingRepurchaseTerm from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingRepurchaseTerm(f *field.UnderlyingRepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRepurchaseRate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingRepurchaseRate() (*field.UnderlyingRepurchaseRate, errors.MessageRejectError) {
	f := new(field.UnderlyingRepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRepurchaseRate reads a UnderlyingRepurchaseRate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingRepurchaseRate(f *field.UnderlyingRepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFactor is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingFactor() (*field.UnderlyingFactor, errors.MessageRejectError) {
	f := new(field.UnderlyingFactor)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFactor reads a UnderlyingFactor from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingFactor(f *field.UnderlyingFactor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCreditRating is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCreditRating() (*field.UnderlyingCreditRating, errors.MessageRejectError) {
	f := new(field.UnderlyingCreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCreditRating reads a UnderlyingCreditRating from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCreditRating(f *field.UnderlyingCreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingInstrRegistry is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingInstrRegistry() (*field.UnderlyingInstrRegistry, errors.MessageRejectError) {
	f := new(field.UnderlyingInstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingInstrRegistry reads a UnderlyingInstrRegistry from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingInstrRegistry(f *field.UnderlyingInstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCountryOfIssue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCountryOfIssue() (*field.UnderlyingCountryOfIssue, errors.MessageRejectError) {
	f := new(field.UnderlyingCountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCountryOfIssue reads a UnderlyingCountryOfIssue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCountryOfIssue(f *field.UnderlyingCountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStateOrProvinceOfIssue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingStateOrProvinceOfIssue() (*field.UnderlyingStateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.UnderlyingStateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStateOrProvinceOfIssue reads a UnderlyingStateOrProvinceOfIssue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingStateOrProvinceOfIssue(f *field.UnderlyingStateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingLocaleOfIssue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingLocaleOfIssue() (*field.UnderlyingLocaleOfIssue, errors.MessageRejectError) {
	f := new(field.UnderlyingLocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingLocaleOfIssue reads a UnderlyingLocaleOfIssue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingLocaleOfIssue(f *field.UnderlyingLocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingRedemptionDate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingRedemptionDate() (*field.UnderlyingRedemptionDate, errors.MessageRejectError) {
	f := new(field.UnderlyingRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingRedemptionDate reads a UnderlyingRedemptionDate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingRedemptionDate(f *field.UnderlyingRedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikePrice is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingStrikePrice() (*field.UnderlyingStrikePrice, errors.MessageRejectError) {
	f := new(field.UnderlyingStrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikePrice reads a UnderlyingStrikePrice from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingStrikePrice(f *field.UnderlyingStrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStrikeCurrency is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingStrikeCurrency() (*field.UnderlyingStrikeCurrency, errors.MessageRejectError) {
	f := new(field.UnderlyingStrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStrikeCurrency reads a UnderlyingStrikeCurrency from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingStrikeCurrency(f *field.UnderlyingStrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingOptAttribute is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingOptAttribute() (*field.UnderlyingOptAttribute, errors.MessageRejectError) {
	f := new(field.UnderlyingOptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingOptAttribute reads a UnderlyingOptAttribute from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingOptAttribute(f *field.UnderlyingOptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingContractMultiplier is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingContractMultiplier() (*field.UnderlyingContractMultiplier, errors.MessageRejectError) {
	f := new(field.UnderlyingContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingContractMultiplier reads a UnderlyingContractMultiplier from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingContractMultiplier(f *field.UnderlyingContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCouponRate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCouponRate() (*field.UnderlyingCouponRate, errors.MessageRejectError) {
	f := new(field.UnderlyingCouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCouponRate reads a UnderlyingCouponRate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCouponRate(f *field.UnderlyingCouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityExchange is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSecurityExchange() (*field.UnderlyingSecurityExchange, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityExchange reads a UnderlyingSecurityExchange from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSecurityExchange(f *field.UnderlyingSecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingIssuer is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingIssuer() (*field.UnderlyingIssuer, errors.MessageRejectError) {
	f := new(field.UnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingIssuer reads a UnderlyingIssuer from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingIssuer(f *field.UnderlyingIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuerLen is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedUnderlyingIssuerLen() (*field.EncodedUnderlyingIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuerLen reads a EncodedUnderlyingIssuerLen from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedUnderlyingIssuerLen(f *field.EncodedUnderlyingIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingIssuer is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedUnderlyingIssuer() (*field.EncodedUnderlyingIssuer, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingIssuer reads a EncodedUnderlyingIssuer from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedUnderlyingIssuer(f *field.EncodedUnderlyingIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSecurityDesc is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSecurityDesc() (*field.UnderlyingSecurityDesc, errors.MessageRejectError) {
	f := new(field.UnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSecurityDesc reads a UnderlyingSecurityDesc from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSecurityDesc(f *field.UnderlyingSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDescLen is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedUnderlyingSecurityDescLen() (*field.EncodedUnderlyingSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDescLen reads a EncodedUnderlyingSecurityDescLen from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedUnderlyingSecurityDescLen(f *field.EncodedUnderlyingSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedUnderlyingSecurityDesc is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedUnderlyingSecurityDesc() (*field.EncodedUnderlyingSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedUnderlyingSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedUnderlyingSecurityDesc reads a EncodedUnderlyingSecurityDesc from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedUnderlyingSecurityDesc(f *field.EncodedUnderlyingSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPProgram is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCPProgram() (*field.UnderlyingCPProgram, errors.MessageRejectError) {
	f := new(field.UnderlyingCPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPProgram reads a UnderlyingCPProgram from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCPProgram(f *field.UnderlyingCPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCPRegType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCPRegType() (*field.UnderlyingCPRegType, errors.MessageRejectError) {
	f := new(field.UnderlyingCPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCPRegType reads a UnderlyingCPRegType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCPRegType(f *field.UnderlyingCPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrency is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCurrency() (*field.UnderlyingCurrency, errors.MessageRejectError) {
	f := new(field.UnderlyingCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrency reads a UnderlyingCurrency from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCurrency(f *field.UnderlyingCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingQty is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingQty() (*field.UnderlyingQty, errors.MessageRejectError) {
	f := new(field.UnderlyingQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingQty reads a UnderlyingQty from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingQty(f *field.UnderlyingQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPx is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingPx() (*field.UnderlyingPx, errors.MessageRejectError) {
	f := new(field.UnderlyingPx)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPx reads a UnderlyingPx from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingPx(f *field.UnderlyingPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingDirtyPrice is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingDirtyPrice() (*field.UnderlyingDirtyPrice, errors.MessageRejectError) {
	f := new(field.UnderlyingDirtyPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingDirtyPrice reads a UnderlyingDirtyPrice from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingDirtyPrice(f *field.UnderlyingDirtyPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndPrice is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingEndPrice() (*field.UnderlyingEndPrice, errors.MessageRejectError) {
	f := new(field.UnderlyingEndPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndPrice reads a UnderlyingEndPrice from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingEndPrice(f *field.UnderlyingEndPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingStartValue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingStartValue() (*field.UnderlyingStartValue, errors.MessageRejectError) {
	f := new(field.UnderlyingStartValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingStartValue reads a UnderlyingStartValue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingStartValue(f *field.UnderlyingStartValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCurrentValue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCurrentValue() (*field.UnderlyingCurrentValue, errors.MessageRejectError) {
	f := new(field.UnderlyingCurrentValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCurrentValue reads a UnderlyingCurrentValue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCurrentValue(f *field.UnderlyingCurrentValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingEndValue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingEndValue() (*field.UnderlyingEndValue, errors.MessageRejectError) {
	f := new(field.UnderlyingEndValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingEndValue reads a UnderlyingEndValue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingEndValue(f *field.UnderlyingEndValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyingStips is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) NoUnderlyingStips() (*field.NoUnderlyingStips, errors.MessageRejectError) {
	f := new(field.NoUnderlyingStips)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyingStips reads a NoUnderlyingStips from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetNoUnderlyingStips(f *field.NoUnderlyingStips) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAllocationPercent is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingAllocationPercent() (*field.UnderlyingAllocationPercent, errors.MessageRejectError) {
	f := new(field.UnderlyingAllocationPercent)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAllocationPercent reads a UnderlyingAllocationPercent from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingAllocationPercent(f *field.UnderlyingAllocationPercent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlementType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSettlementType() (*field.UnderlyingSettlementType, errors.MessageRejectError) {
	f := new(field.UnderlyingSettlementType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlementType reads a UnderlyingSettlementType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSettlementType(f *field.UnderlyingSettlementType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashAmount is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCashAmount() (*field.UnderlyingCashAmount, errors.MessageRejectError) {
	f := new(field.UnderlyingCashAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashAmount reads a UnderlyingCashAmount from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCashAmount(f *field.UnderlyingCashAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCashType is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCashType() (*field.UnderlyingCashType, errors.MessageRejectError) {
	f := new(field.UnderlyingCashType)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCashType reads a UnderlyingCashType from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCashType(f *field.UnderlyingCashType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingUnitOfMeasure is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingUnitOfMeasure() (*field.UnderlyingUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnderlyingUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingUnitOfMeasure reads a UnderlyingUnitOfMeasure from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingUnitOfMeasure(f *field.UnderlyingUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTimeUnit is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingTimeUnit() (*field.UnderlyingTimeUnit, errors.MessageRejectError) {
	f := new(field.UnderlyingTimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTimeUnit reads a UnderlyingTimeUnit from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingTimeUnit(f *field.UnderlyingTimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingCapValue is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingCapValue() (*field.UnderlyingCapValue, errors.MessageRejectError) {
	f := new(field.UnderlyingCapValue)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingCapValue reads a UnderlyingCapValue from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingCapValue(f *field.UnderlyingCapValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUndlyInstrumentParties is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) NoUndlyInstrumentParties() (*field.NoUndlyInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoUndlyInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUndlyInstrumentParties reads a NoUndlyInstrumentParties from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetNoUndlyInstrumentParties(f *field.NoUndlyInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingSettlMethod is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingSettlMethod() (*field.UnderlyingSettlMethod, errors.MessageRejectError) {
	f := new(field.UnderlyingSettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingSettlMethod reads a UnderlyingSettlMethod from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingSettlMethod(f *field.UnderlyingSettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingAdjustedQuantity is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingAdjustedQuantity() (*field.UnderlyingAdjustedQuantity, errors.MessageRejectError) {
	f := new(field.UnderlyingAdjustedQuantity)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingAdjustedQuantity reads a UnderlyingAdjustedQuantity from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingAdjustedQuantity(f *field.UnderlyingAdjustedQuantity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRate is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingFXRate() (*field.UnderlyingFXRate, errors.MessageRejectError) {
	f := new(field.UnderlyingFXRate)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRate reads a UnderlyingFXRate from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingFXRate(f *field.UnderlyingFXRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingFXRateCalc is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) UnderlyingFXRateCalc() (*field.UnderlyingFXRateCalc, errors.MessageRejectError) {
	f := new(field.UnderlyingFXRateCalc)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingFXRateCalc reads a UnderlyingFXRateCalc from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetUnderlyingFXRateCalc(f *field.UnderlyingFXRateCalc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for OrderMassCancelRequest.
func (m OrderMassCancelRequest) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from OrderMassCancelRequest.
func (m OrderMassCancelRequest) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}
