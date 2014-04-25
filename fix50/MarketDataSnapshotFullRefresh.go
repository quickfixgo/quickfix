package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//MarketDataSnapshotFullRefresh msg type = W.
type MarketDataSnapshotFullRefresh struct {
	message.Message
}

//MarketDataSnapshotFullRefreshBuilder builds MarketDataSnapshotFullRefresh messages.
type MarketDataSnapshotFullRefreshBuilder struct {
	message.MessageBuilder
}

//CreateMarketDataSnapshotFullRefreshBuilder returns an initialized MarketDataSnapshotFullRefreshBuilder with specified required fields.
func CreateMarketDataSnapshotFullRefreshBuilder(
	nomdentries field.NoMDEntries) MarketDataSnapshotFullRefreshBuilder {
	var builder MarketDataSnapshotFullRefreshBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("W"))
	builder.Body.Set(nomdentries)
	return builder
}

//MDReqID is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MDReqID() (*field.MDReqID, errors.MessageRejectError) {
	f := new(field.MDReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqID reads a MDReqID from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMDReqID(f *field.MDReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FinancialStatus is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) FinancialStatus() (*field.FinancialStatus, errors.MessageRejectError) {
	f := new(field.FinancialStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetFinancialStatus reads a FinancialStatus from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetFinancialStatus(f *field.FinancialStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CorporateAction is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CorporateAction() (*field.CorporateAction, errors.MessageRejectError) {
	f := new(field.CorporateAction)
	err := m.Body.Get(f)
	return f, err
}

//GetCorporateAction reads a CorporateAction from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCorporateAction(f *field.CorporateAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetChgPrevDay is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NetChgPrevDay() (*field.NetChgPrevDay, errors.MessageRejectError) {
	f := new(field.NetChgPrevDay)
	err := m.Body.Get(f)
	return f, err
}

//GetNetChgPrevDay reads a NetChgPrevDay from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNetChgPrevDay(f *field.NetChgPrevDay) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMDEntries is a required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoMDEntries() (*field.NoMDEntries, errors.MessageRejectError) {
	f := new(field.NoMDEntries)
	err := m.Body.Get(f)
	return f, err
}

//GetNoMDEntries reads a NoMDEntries from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNoMDEntries(f *field.NoMDEntries) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueDepth is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ApplQueueDepth() (*field.ApplQueueDepth, errors.MessageRejectError) {
	f := new(field.ApplQueueDepth)
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueDepth reads a ApplQueueDepth from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetApplQueueDepth(f *field.ApplQueueDepth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplQueueResolution is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ApplQueueResolution() (*field.ApplQueueResolution, errors.MessageRejectError) {
	f := new(field.ApplQueueResolution)
	err := m.Body.Get(f)
	return f, err
}

//GetApplQueueResolution reads a ApplQueueResolution from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetApplQueueResolution(f *field.ApplQueueResolution) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDReportID is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MDReportID() (*field.MDReportID, errors.MessageRejectError) {
	f := new(field.MDReportID)
	err := m.Body.Get(f)
	return f, err
}

//GetMDReportID reads a MDReportID from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMDReportID(f *field.MDReportID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDBookType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MDBookType() (*field.MDBookType, errors.MessageRejectError) {
	f := new(field.MDBookType)
	err := m.Body.Get(f)
	return f, err
}

//GetMDBookType reads a MDBookType from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMDBookType(f *field.MDBookType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MDFeedType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MDFeedType() (*field.MDFeedType, errors.MessageRejectError) {
	f := new(field.MDFeedType)
	err := m.Body.Get(f)
	return f, err
}

//GetMDFeedType reads a MDFeedType from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMDFeedType(f *field.MDFeedType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRoutingIDs is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoRoutingIDs() (*field.NoRoutingIDs, errors.MessageRejectError) {
	f := new(field.NoRoutingIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRoutingIDs reads a NoRoutingIDs from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNoRoutingIDs(f *field.NoRoutingIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}
