package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CrossOrderCancelReplaceRequest msg type = t.
type CrossOrderCancelReplaceRequest struct {
	message.Message
}

//CrossOrderCancelReplaceRequestBuilder builds CrossOrderCancelReplaceRequest messages.
type CrossOrderCancelReplaceRequestBuilder struct {
	message.MessageBuilder
}

//CreateCrossOrderCancelReplaceRequestBuilder returns an initialized CrossOrderCancelReplaceRequestBuilder with specified required fields.
func CreateCrossOrderCancelReplaceRequestBuilder(
	crossid field.CrossID,
	origcrossid field.OrigCrossID,
	crosstype field.CrossType,
	crossprioritization field.CrossPrioritization,
	nosides field.NoSides,
	handlinst field.HandlInst,
	transacttime field.TransactTime,
	ordtype field.OrdType) CrossOrderCancelReplaceRequestBuilder {
	var builder CrossOrderCancelReplaceRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIX43))
	builder.Header.Set(field.BuildMsgType("t"))
	builder.Body.Set(crossid)
	builder.Body.Set(origcrossid)
	builder.Body.Set(crosstype)
	builder.Body.Set(crossprioritization)
	builder.Body.Set(nosides)
	builder.Body.Set(handlinst)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//OrderID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossID is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CrossID() (*field.CrossID, errors.MessageRejectError) {
	f := new(field.CrossID)
	err := m.Body.Get(f)
	return f, err
}

//GetCrossID reads a CrossID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCrossID(f *field.CrossID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigCrossID is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OrigCrossID() (*field.OrigCrossID, errors.MessageRejectError) {
	f := new(field.OrigCrossID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigCrossID reads a OrigCrossID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetOrigCrossID(f *field.OrigCrossID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossType is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CrossType() (*field.CrossType, errors.MessageRejectError) {
	f := new(field.CrossType)
	err := m.Body.Get(f)
	return f, err
}

//GetCrossType reads a CrossType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCrossType(f *field.CrossType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossPrioritization is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CrossPrioritization() (*field.CrossPrioritization, errors.MessageRejectError) {
	f := new(field.CrossPrioritization)
	err := m.Body.Get(f)
	return f, err
}

//GetCrossPrioritization reads a CrossPrioritization from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCrossPrioritization(f *field.CrossPrioritization) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSides is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoSides() (*field.NoSides, errors.MessageRejectError) {
	f := new(field.NoSides)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSides reads a NoSides from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetNoSides(f *field.NoSides) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) SettlmntTyp() (*field.SettlmntTyp, errors.MessageRejectError) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSettlmntTyp(f *field.SettlmntTyp) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) FutSettDate() (*field.FutSettDate, errors.MessageRejectError) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetFutSettDate(f *field.FutSettDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) HandlInst() (*field.HandlInst, errors.MessageRejectError) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetHandlInst(f *field.HandlInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExecInst() (*field.ExecInst, errors.MessageRejectError) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetExecInst(f *field.ExecInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MinQty() (*field.MinQty, errors.MessageRejectError) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetMinQty(f *field.MinQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaxFloor() (*field.MaxFloor, errors.MessageRejectError) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetMaxFloor(f *field.MaxFloor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExDestination() (*field.ExDestination, errors.MessageRejectError) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetExDestination(f *field.ExDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoTradingSessions() (*field.NoTradingSessions, errors.MessageRejectError) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetNoTradingSessions(f *field.NoTradingSessions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ProcessCode() (*field.ProcessCode, errors.MessageRejectError) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetProcessCode(f *field.ProcessCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PrevClosePx() (*field.PrevClosePx, errors.MessageRejectError) {
	f := new(field.PrevClosePx)
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetPrevClosePx(f *field.PrevClosePx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) LocateReqd() (*field.LocateReqd, errors.MessageRejectError) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetLocateReqd(f *field.LocateReqd) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NoStipulations() (*field.NoStipulations, errors.MessageRejectError) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetNoStipulations(f *field.NoStipulations) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) OrdType() (*field.OrdType, errors.MessageRejectError) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetOrdType(f *field.OrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) StopPx() (*field.StopPx, errors.MessageRejectError) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetStopPx(f *field.StopPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Spread() (*field.Spread, errors.MessageRejectError) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetSpread(f *field.Spread) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkCurveName() (*field.BenchmarkCurveName, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetBenchmarkCurveName(f *field.BenchmarkCurveName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, errors.MessageRejectError) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) YieldType() (*field.YieldType, errors.MessageRejectError) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetYieldType(f *field.YieldType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Yield() (*field.Yield, errors.MessageRejectError) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetYield(f *field.Yield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ComplianceID() (*field.ComplianceID, errors.MessageRejectError) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetComplianceID(f *field.ComplianceID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIid is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) IOIid() (*field.IOIid, errors.MessageRejectError) {
	f := new(field.IOIid)
	err := m.Body.Get(f)
	return f, err
}

//GetIOIid reads a IOIid from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetIOIid(f *field.IOIid) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) QuoteID() (*field.QuoteID, errors.MessageRejectError) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetQuoteID(f *field.QuoteID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) TimeInForce() (*field.TimeInForce, errors.MessageRejectError) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetTimeInForce(f *field.TimeInForce) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) EffectiveTime() (*field.EffectiveTime, errors.MessageRejectError) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetEffectiveTime(f *field.EffectiveTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExpireDate() (*field.ExpireDate, errors.MessageRejectError) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetExpireDate(f *field.ExpireDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) ExpireTime() (*field.ExpireTime, errors.MessageRejectError) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetExpireTime(f *field.ExpireTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GTBookingInst() (*field.GTBookingInst, errors.MessageRejectError) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetGTBookingInst(f *field.GTBookingInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MaxShow() (*field.MaxShow, errors.MessageRejectError) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetMaxShow(f *field.MaxShow) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegDifference is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) PegDifference() (*field.PegDifference, errors.MessageRejectError) {
	f := new(field.PegDifference)
	err := m.Body.Get(f)
	return f, err
}

//GetPegDifference reads a PegDifference from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetPegDifference(f *field.PegDifference) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionInst() (*field.DiscretionInst, errors.MessageRejectError) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetDiscretionInst(f *field.DiscretionInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffset is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) DiscretionOffset() (*field.DiscretionOffset, errors.MessageRejectError) {
	f := new(field.DiscretionOffset)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffset reads a DiscretionOffset from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetDiscretionOffset(f *field.DiscretionOffset) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) CancellationRights() (*field.CancellationRights, errors.MessageRejectError) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetCancellationRights(f *field.CancellationRights) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, errors.MessageRejectError) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) RegistID() (*field.RegistID, errors.MessageRejectError) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetRegistID(f *field.RegistID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) Designation() (*field.Designation, errors.MessageRejectError) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetDesignation(f *field.Designation) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestRate is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) AccruedInterestRate() (*field.AccruedInterestRate, errors.MessageRejectError) {
	f := new(field.AccruedInterestRate)
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestRate reads a AccruedInterestRate from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetAccruedInterestRate(f *field.AccruedInterestRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestAmt is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) AccruedInterestAmt() (*field.AccruedInterestAmt, errors.MessageRejectError) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestAmt reads a AccruedInterestAmt from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetAccruedInterestAmt(f *field.AccruedInterestAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a non-required field for CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) NetMoney() (*field.NetMoney, errors.MessageRejectError) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from CrossOrderCancelReplaceRequest.
func (m CrossOrderCancelReplaceRequest) GetNetMoney(f *field.NetMoney) errors.MessageRejectError {
	return m.Body.Get(f)
}
