package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CrossOrderCancelRequest msg type = u.
type CrossOrderCancelRequest struct {
	message.Message
}

//CrossOrderCancelRequestBuilder builds CrossOrderCancelRequest messages.
type CrossOrderCancelRequestBuilder struct {
	message.MessageBuilder
}

//CreateCrossOrderCancelRequestBuilder returns an initialized CrossOrderCancelRequestBuilder with specified required fields.
func CreateCrossOrderCancelRequestBuilder(
	crossid field.CrossID,
	origcrossid field.OrigCrossID,
	crosstype field.CrossType,
	crossprioritization field.CrossPrioritization,
	nosides field.NoSides,
	transacttime field.TransactTime) CrossOrderCancelRequestBuilder {
	var builder CrossOrderCancelRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("u"))
	builder.Body.Set(crossid)
	builder.Body.Set(origcrossid)
	builder.Body.Set(crosstype)
	builder.Body.Set(crossprioritization)
	builder.Body.Set(nosides)
	builder.Body.Set(transacttime)
	return builder
}

//OrderID is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossID is a required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CrossID() (*field.CrossID, errors.MessageRejectError) {
	f := new(field.CrossID)
	err := m.Body.Get(f)
	return f, err
}

//GetCrossID reads a CrossID from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetCrossID(f *field.CrossID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigCrossID is a required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) OrigCrossID() (*field.OrigCrossID, errors.MessageRejectError) {
	f := new(field.OrigCrossID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigCrossID reads a OrigCrossID from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetOrigCrossID(f *field.OrigCrossID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossType is a required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CrossType() (*field.CrossType, errors.MessageRejectError) {
	f := new(field.CrossType)
	err := m.Body.Get(f)
	return f, err
}

//GetCrossType reads a CrossType from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetCrossType(f *field.CrossType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossPrioritization is a required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CrossPrioritization() (*field.CrossPrioritization, errors.MessageRejectError) {
	f := new(field.CrossPrioritization)
	err := m.Body.Get(f)
	return f, err
}

//GetCrossPrioritization reads a CrossPrioritization from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetCrossPrioritization(f *field.CrossPrioritization) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSides is a required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) NoSides() (*field.NoSides, errors.MessageRejectError) {
	f := new(field.NoSides)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSides reads a NoSides from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetNoSides(f *field.NoSides) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SecurityGroup() (*field.SecurityGroup, errors.MessageRejectError) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetSecurityGroup(f *field.SecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) UnitOfMeasureQty() (*field.UnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetUnitOfMeasureQty(f *field.UnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SecurityXMLLen() (*field.SecurityXMLLen, errors.MessageRejectError) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetSecurityXMLLen(f *field.SecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SecurityXML() (*field.SecurityXML, errors.MessageRejectError) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetSecurityXML(f *field.SecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SecurityXMLSchema() (*field.SecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetSecurityXMLSchema(f *field.SecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) ProductComplex() (*field.ProductComplex, errors.MessageRejectError) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetProductComplex(f *field.ProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) SettlMethod() (*field.SettlMethod, errors.MessageRejectError) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetSettlMethod(f *field.SettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) ExerciseStyle() (*field.ExerciseStyle, errors.MessageRejectError) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetExerciseStyle(f *field.ExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) OptPayoutAmount() (*field.OptPayoutAmount, errors.MessageRejectError) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetOptPayoutAmount(f *field.OptPayoutAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) PriceQuoteMethod() (*field.PriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetPriceQuoteMethod(f *field.PriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) ListMethod() (*field.ListMethod, errors.MessageRejectError) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetListMethod(f *field.ListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) CapPrice() (*field.CapPrice, errors.MessageRejectError) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetCapPrice(f *field.CapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) FloorPrice() (*field.FloorPrice, errors.MessageRejectError) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetFloorPrice(f *field.FloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) FlexibleIndicator() (*field.FlexibleIndicator, errors.MessageRejectError) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetFlexibleIndicator(f *field.FlexibleIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) ValuationMethod() (*field.ValuationMethod, errors.MessageRejectError) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetValuationMethod(f *field.ValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) ContractMultiplierUnit() (*field.ContractMultiplierUnit, errors.MessageRejectError) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetContractMultiplierUnit(f *field.ContractMultiplierUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) FlowScheduleType() (*field.FlowScheduleType, errors.MessageRejectError) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetFlowScheduleType(f *field.FlowScheduleType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) RestructuringType() (*field.RestructuringType, errors.MessageRejectError) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetRestructuringType(f *field.RestructuringType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) Seniority() (*field.Seniority, errors.MessageRejectError) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetSeniority(f *field.Seniority) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) AttachmentPoint() (*field.AttachmentPoint, errors.MessageRejectError) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetAttachmentPoint(f *field.AttachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) DetachmentPoint() (*field.DetachmentPoint, errors.MessageRejectError) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetDetachmentPoint(f *field.DetachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecision) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) OptPayoutType() (*field.OptPayoutType, errors.MessageRejectError) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetOptPayoutType(f *field.OptPayoutType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) NoComplexEvents() (*field.NoComplexEvents, errors.MessageRejectError) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetNoComplexEvents(f *field.NoComplexEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HostCrossID is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) HostCrossID() (*field.HostCrossID, errors.MessageRejectError) {
	f := new(field.HostCrossID)
	err := m.Body.Get(f)
	return f, err
}

//GetHostCrossID reads a HostCrossID from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetHostCrossID(f *field.HostCrossID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRootPartyIDs is a non-required field for CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) NoRootPartyIDs() (*field.NoRootPartyIDs, errors.MessageRejectError) {
	f := new(field.NoRootPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRootPartyIDs reads a NoRootPartyIDs from CrossOrderCancelRequest.
func (m CrossOrderCancelRequest) GetNoRootPartyIDs(f *field.NoRootPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}
