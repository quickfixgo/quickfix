package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ExecutionAcknowledgement msg type = BN.
type ExecutionAcknowledgement struct {
	message.Message
}

//ExecutionAcknowledgementBuilder builds ExecutionAcknowledgement messages.
type ExecutionAcknowledgementBuilder struct {
	message.MessageBuilder
}

//CreateExecutionAcknowledgementBuilder returns an initialized ExecutionAcknowledgementBuilder with specified required fields.
func CreateExecutionAcknowledgementBuilder(
	orderid field.OrderID,
	execackstatus field.ExecAckStatus,
	execid field.ExecID,
	side field.Side) ExecutionAcknowledgementBuilder {
	var builder ExecutionAcknowledgementBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(orderid)
	builder.Body.Set(execackstatus)
	builder.Body.Set(execid)
	builder.Body.Set(side)
	return builder
}

//OrderID is a required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecondaryOrderID() (*field.SecondaryOrderID, errors.MessageRejectError) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetSecondaryOrderID(f *field.SecondaryOrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecAckStatus is a required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) ExecAckStatus() (*field.ExecAckStatus, errors.MessageRejectError) {
	f := new(field.ExecAckStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetExecAckStatus reads a ExecAckStatus from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetExecAckStatus(f *field.ExecAckStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecID is a required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) ExecID() (*field.ExecID, errors.MessageRejectError) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}

//GetExecID reads a ExecID from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetExecID(f *field.ExecID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DKReason is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) DKReason() (*field.DKReason, errors.MessageRejectError) {
	f := new(field.DKReason)
	err := m.Body.Get(f)
	return f, err
}

//GetDKReason reads a DKReason from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetDKReason(f *field.DKReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CashOrderQty() (*field.CashOrderQty, errors.MessageRejectError) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetCashOrderQty(f *field.CashOrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) OrderPercent() (*field.OrderPercent, errors.MessageRejectError) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetOrderPercent(f *field.OrderPercent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) RoundingDirection() (*field.RoundingDirection, errors.MessageRejectError) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetRoundingDirection(f *field.RoundingDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) RoundingModulus() (*field.RoundingModulus, errors.MessageRejectError) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetRoundingModulus(f *field.RoundingModulus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastQty is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) LastQty() (*field.LastQty, errors.MessageRejectError) {
	f := new(field.LastQty)
	err := m.Body.Get(f)
	return f, err
}

//GetLastQty reads a LastQty from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetLastQty(f *field.LastQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) LastPx() (*field.LastPx, errors.MessageRejectError) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetLastPx(f *field.LastPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastParPx is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) LastParPx() (*field.LastParPx, errors.MessageRejectError) {
	f := new(field.LastParPx)
	err := m.Body.Get(f)
	return f, err
}

//GetLastParPx reads a LastParPx from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetLastParPx(f *field.LastParPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CumQty is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) CumQty() (*field.CumQty, errors.MessageRejectError) {
	f := new(field.CumQty)
	err := m.Body.Get(f)
	return f, err
}

//GetCumQty reads a CumQty from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetCumQty(f *field.CumQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) AvgPx() (*field.AvgPx, errors.MessageRejectError) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetAvgPx(f *field.AvgPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ExecutionAcknowledgement.
func (m ExecutionAcknowledgement) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
