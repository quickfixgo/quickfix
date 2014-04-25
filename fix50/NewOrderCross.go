package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//NewOrderCross msg type = s.
type NewOrderCross struct {
	message.Message
}

//NewOrderCrossBuilder builds NewOrderCross messages.
type NewOrderCrossBuilder struct {
	message.MessageBuilder
}

//CreateNewOrderCrossBuilder returns an initialized NewOrderCrossBuilder with specified required fields.
func CreateNewOrderCrossBuilder(
	crossid field.CrossID,
	crosstype field.CrossType,
	crossprioritization field.CrossPrioritization,
	nosides field.NoSides,
	transacttime field.TransactTime,
	ordtype field.OrdType) NewOrderCrossBuilder {
	var builder NewOrderCrossBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("s"))
	builder.Body.Set(crossid)
	builder.Body.Set(crosstype)
	builder.Body.Set(crossprioritization)
	builder.Body.Set(nosides)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//CrossID is a required field for NewOrderCross.
func (m NewOrderCross) CrossID() (*field.CrossID, errors.MessageRejectError) {
	f := new(field.CrossID)
	err := m.Body.Get(f)
	return f, err
}

//GetCrossID reads a CrossID from NewOrderCross.
func (m NewOrderCross) GetCrossID(f *field.CrossID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossType is a required field for NewOrderCross.
func (m NewOrderCross) CrossType() (*field.CrossType, errors.MessageRejectError) {
	f := new(field.CrossType)
	err := m.Body.Get(f)
	return f, err
}

//GetCrossType reads a CrossType from NewOrderCross.
func (m NewOrderCross) GetCrossType(f *field.CrossType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CrossPrioritization is a required field for NewOrderCross.
func (m NewOrderCross) CrossPrioritization() (*field.CrossPrioritization, errors.MessageRejectError) {
	f := new(field.CrossPrioritization)
	err := m.Body.Get(f)
	return f, err
}

//GetCrossPrioritization reads a CrossPrioritization from NewOrderCross.
func (m NewOrderCross) GetCrossPrioritization(f *field.CrossPrioritization) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSides is a required field for NewOrderCross.
func (m NewOrderCross) NoSides() (*field.NoSides, errors.MessageRejectError) {
	f := new(field.NoSides)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSides reads a NoSides from NewOrderCross.
func (m NewOrderCross) GetNoSides(f *field.NoSides) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for NewOrderCross.
func (m NewOrderCross) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from NewOrderCross.
func (m NewOrderCross) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for NewOrderCross.
func (m NewOrderCross) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from NewOrderCross.
func (m NewOrderCross) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for NewOrderCross.
func (m NewOrderCross) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from NewOrderCross.
func (m NewOrderCross) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for NewOrderCross.
func (m NewOrderCross) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from NewOrderCross.
func (m NewOrderCross) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for NewOrderCross.
func (m NewOrderCross) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from NewOrderCross.
func (m NewOrderCross) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for NewOrderCross.
func (m NewOrderCross) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from NewOrderCross.
func (m NewOrderCross) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for NewOrderCross.
func (m NewOrderCross) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from NewOrderCross.
func (m NewOrderCross) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for NewOrderCross.
func (m NewOrderCross) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from NewOrderCross.
func (m NewOrderCross) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for NewOrderCross.
func (m NewOrderCross) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from NewOrderCross.
func (m NewOrderCross) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for NewOrderCross.
func (m NewOrderCross) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from NewOrderCross.
func (m NewOrderCross) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for NewOrderCross.
func (m NewOrderCross) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from NewOrderCross.
func (m NewOrderCross) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for NewOrderCross.
func (m NewOrderCross) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from NewOrderCross.
func (m NewOrderCross) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for NewOrderCross.
func (m NewOrderCross) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from NewOrderCross.
func (m NewOrderCross) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for NewOrderCross.
func (m NewOrderCross) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from NewOrderCross.
func (m NewOrderCross) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for NewOrderCross.
func (m NewOrderCross) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from NewOrderCross.
func (m NewOrderCross) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for NewOrderCross.
func (m NewOrderCross) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from NewOrderCross.
func (m NewOrderCross) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for NewOrderCross.
func (m NewOrderCross) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from NewOrderCross.
func (m NewOrderCross) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for NewOrderCross.
func (m NewOrderCross) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from NewOrderCross.
func (m NewOrderCross) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for NewOrderCross.
func (m NewOrderCross) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from NewOrderCross.
func (m NewOrderCross) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for NewOrderCross.
func (m NewOrderCross) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from NewOrderCross.
func (m NewOrderCross) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for NewOrderCross.
func (m NewOrderCross) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from NewOrderCross.
func (m NewOrderCross) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for NewOrderCross.
func (m NewOrderCross) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from NewOrderCross.
func (m NewOrderCross) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for NewOrderCross.
func (m NewOrderCross) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from NewOrderCross.
func (m NewOrderCross) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for NewOrderCross.
func (m NewOrderCross) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from NewOrderCross.
func (m NewOrderCross) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for NewOrderCross.
func (m NewOrderCross) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from NewOrderCross.
func (m NewOrderCross) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for NewOrderCross.
func (m NewOrderCross) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from NewOrderCross.
func (m NewOrderCross) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for NewOrderCross.
func (m NewOrderCross) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from NewOrderCross.
func (m NewOrderCross) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for NewOrderCross.
func (m NewOrderCross) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from NewOrderCross.
func (m NewOrderCross) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for NewOrderCross.
func (m NewOrderCross) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from NewOrderCross.
func (m NewOrderCross) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for NewOrderCross.
func (m NewOrderCross) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from NewOrderCross.
func (m NewOrderCross) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for NewOrderCross.
func (m NewOrderCross) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from NewOrderCross.
func (m NewOrderCross) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for NewOrderCross.
func (m NewOrderCross) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from NewOrderCross.
func (m NewOrderCross) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for NewOrderCross.
func (m NewOrderCross) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from NewOrderCross.
func (m NewOrderCross) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for NewOrderCross.
func (m NewOrderCross) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from NewOrderCross.
func (m NewOrderCross) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for NewOrderCross.
func (m NewOrderCross) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from NewOrderCross.
func (m NewOrderCross) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for NewOrderCross.
func (m NewOrderCross) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from NewOrderCross.
func (m NewOrderCross) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for NewOrderCross.
func (m NewOrderCross) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from NewOrderCross.
func (m NewOrderCross) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for NewOrderCross.
func (m NewOrderCross) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from NewOrderCross.
func (m NewOrderCross) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for NewOrderCross.
func (m NewOrderCross) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from NewOrderCross.
func (m NewOrderCross) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for NewOrderCross.
func (m NewOrderCross) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from NewOrderCross.
func (m NewOrderCross) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for NewOrderCross.
func (m NewOrderCross) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from NewOrderCross.
func (m NewOrderCross) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for NewOrderCross.
func (m NewOrderCross) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from NewOrderCross.
func (m NewOrderCross) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for NewOrderCross.
func (m NewOrderCross) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from NewOrderCross.
func (m NewOrderCross) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for NewOrderCross.
func (m NewOrderCross) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from NewOrderCross.
func (m NewOrderCross) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for NewOrderCross.
func (m NewOrderCross) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from NewOrderCross.
func (m NewOrderCross) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for NewOrderCross.
func (m NewOrderCross) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from NewOrderCross.
func (m NewOrderCross) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for NewOrderCross.
func (m NewOrderCross) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from NewOrderCross.
func (m NewOrderCross) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for NewOrderCross.
func (m NewOrderCross) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from NewOrderCross.
func (m NewOrderCross) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for NewOrderCross.
func (m NewOrderCross) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from NewOrderCross.
func (m NewOrderCross) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for NewOrderCross.
func (m NewOrderCross) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from NewOrderCross.
func (m NewOrderCross) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for NewOrderCross.
func (m NewOrderCross) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from NewOrderCross.
func (m NewOrderCross) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for NewOrderCross.
func (m NewOrderCross) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from NewOrderCross.
func (m NewOrderCross) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for NewOrderCross.
func (m NewOrderCross) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from NewOrderCross.
func (m NewOrderCross) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for NewOrderCross.
func (m NewOrderCross) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from NewOrderCross.
func (m NewOrderCross) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for NewOrderCross.
func (m NewOrderCross) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from NewOrderCross.
func (m NewOrderCross) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for NewOrderCross.
func (m NewOrderCross) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from NewOrderCross.
func (m NewOrderCross) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for NewOrderCross.
func (m NewOrderCross) SettlType() (*field.SettlType, errors.MessageRejectError) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from NewOrderCross.
func (m NewOrderCross) GetSettlType(f *field.SettlType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for NewOrderCross.
func (m NewOrderCross) SettlDate() (*field.SettlDate, errors.MessageRejectError) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from NewOrderCross.
func (m NewOrderCross) GetSettlDate(f *field.SettlDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a non-required field for NewOrderCross.
func (m NewOrderCross) HandlInst() (*field.HandlInst, errors.MessageRejectError) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from NewOrderCross.
func (m NewOrderCross) GetHandlInst(f *field.HandlInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for NewOrderCross.
func (m NewOrderCross) ExecInst() (*field.ExecInst, errors.MessageRejectError) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from NewOrderCross.
func (m NewOrderCross) GetExecInst(f *field.ExecInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for NewOrderCross.
func (m NewOrderCross) MinQty() (*field.MinQty, errors.MessageRejectError) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from NewOrderCross.
func (m NewOrderCross) GetMinQty(f *field.MinQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for NewOrderCross.
func (m NewOrderCross) MaxFloor() (*field.MaxFloor, errors.MessageRejectError) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from NewOrderCross.
func (m NewOrderCross) GetMaxFloor(f *field.MaxFloor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for NewOrderCross.
func (m NewOrderCross) ExDestination() (*field.ExDestination, errors.MessageRejectError) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from NewOrderCross.
func (m NewOrderCross) GetExDestination(f *field.ExDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for NewOrderCross.
func (m NewOrderCross) NoTradingSessions() (*field.NoTradingSessions, errors.MessageRejectError) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from NewOrderCross.
func (m NewOrderCross) GetNoTradingSessions(f *field.NoTradingSessions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for NewOrderCross.
func (m NewOrderCross) ProcessCode() (*field.ProcessCode, errors.MessageRejectError) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from NewOrderCross.
func (m NewOrderCross) GetProcessCode(f *field.ProcessCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for NewOrderCross.
func (m NewOrderCross) PrevClosePx() (*field.PrevClosePx, errors.MessageRejectError) {
	f := new(field.PrevClosePx)
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from NewOrderCross.
func (m NewOrderCross) GetPrevClosePx(f *field.PrevClosePx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for NewOrderCross.
func (m NewOrderCross) LocateReqd() (*field.LocateReqd, errors.MessageRejectError) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from NewOrderCross.
func (m NewOrderCross) GetLocateReqd(f *field.LocateReqd) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for NewOrderCross.
func (m NewOrderCross) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from NewOrderCross.
func (m NewOrderCross) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for NewOrderCross.
func (m NewOrderCross) NoStipulations() (*field.NoStipulations, errors.MessageRejectError) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from NewOrderCross.
func (m NewOrderCross) GetNoStipulations(f *field.NoStipulations) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for NewOrderCross.
func (m NewOrderCross) OrdType() (*field.OrdType, errors.MessageRejectError) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from NewOrderCross.
func (m NewOrderCross) GetOrdType(f *field.OrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for NewOrderCross.
func (m NewOrderCross) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from NewOrderCross.
func (m NewOrderCross) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for NewOrderCross.
func (m NewOrderCross) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from NewOrderCross.
func (m NewOrderCross) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for NewOrderCross.
func (m NewOrderCross) StopPx() (*field.StopPx, errors.MessageRejectError) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from NewOrderCross.
func (m NewOrderCross) GetStopPx(f *field.StopPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for NewOrderCross.
func (m NewOrderCross) Spread() (*field.Spread, errors.MessageRejectError) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from NewOrderCross.
func (m NewOrderCross) GetSpread(f *field.Spread) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for NewOrderCross.
func (m NewOrderCross) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from NewOrderCross.
func (m NewOrderCross) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for NewOrderCross.
func (m NewOrderCross) BenchmarkCurveName() (*field.BenchmarkCurveName, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from NewOrderCross.
func (m NewOrderCross) GetBenchmarkCurveName(f *field.BenchmarkCurveName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for NewOrderCross.
func (m NewOrderCross) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, errors.MessageRejectError) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from NewOrderCross.
func (m NewOrderCross) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for NewOrderCross.
func (m NewOrderCross) BenchmarkPrice() (*field.BenchmarkPrice, errors.MessageRejectError) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from NewOrderCross.
func (m NewOrderCross) GetBenchmarkPrice(f *field.BenchmarkPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for NewOrderCross.
func (m NewOrderCross) BenchmarkPriceType() (*field.BenchmarkPriceType, errors.MessageRejectError) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from NewOrderCross.
func (m NewOrderCross) GetBenchmarkPriceType(f *field.BenchmarkPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for NewOrderCross.
func (m NewOrderCross) BenchmarkSecurityID() (*field.BenchmarkSecurityID, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from NewOrderCross.
func (m NewOrderCross) GetBenchmarkSecurityID(f *field.BenchmarkSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for NewOrderCross.
func (m NewOrderCross) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from NewOrderCross.
func (m NewOrderCross) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for NewOrderCross.
func (m NewOrderCross) YieldType() (*field.YieldType, errors.MessageRejectError) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from NewOrderCross.
func (m NewOrderCross) GetYieldType(f *field.YieldType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for NewOrderCross.
func (m NewOrderCross) Yield() (*field.Yield, errors.MessageRejectError) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from NewOrderCross.
func (m NewOrderCross) GetYield(f *field.Yield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for NewOrderCross.
func (m NewOrderCross) YieldCalcDate() (*field.YieldCalcDate, errors.MessageRejectError) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from NewOrderCross.
func (m NewOrderCross) GetYieldCalcDate(f *field.YieldCalcDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for NewOrderCross.
func (m NewOrderCross) YieldRedemptionDate() (*field.YieldRedemptionDate, errors.MessageRejectError) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from NewOrderCross.
func (m NewOrderCross) GetYieldRedemptionDate(f *field.YieldRedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for NewOrderCross.
func (m NewOrderCross) YieldRedemptionPrice() (*field.YieldRedemptionPrice, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from NewOrderCross.
func (m NewOrderCross) GetYieldRedemptionPrice(f *field.YieldRedemptionPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for NewOrderCross.
func (m NewOrderCross) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from NewOrderCross.
func (m NewOrderCross) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for NewOrderCross.
func (m NewOrderCross) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from NewOrderCross.
func (m NewOrderCross) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for NewOrderCross.
func (m NewOrderCross) ComplianceID() (*field.ComplianceID, errors.MessageRejectError) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from NewOrderCross.
func (m NewOrderCross) GetComplianceID(f *field.ComplianceID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIID is a non-required field for NewOrderCross.
func (m NewOrderCross) IOIID() (*field.IOIID, errors.MessageRejectError) {
	f := new(field.IOIID)
	err := m.Body.Get(f)
	return f, err
}

//GetIOIID reads a IOIID from NewOrderCross.
func (m NewOrderCross) GetIOIID(f *field.IOIID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for NewOrderCross.
func (m NewOrderCross) QuoteID() (*field.QuoteID, errors.MessageRejectError) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from NewOrderCross.
func (m NewOrderCross) GetQuoteID(f *field.QuoteID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for NewOrderCross.
func (m NewOrderCross) TimeInForce() (*field.TimeInForce, errors.MessageRejectError) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from NewOrderCross.
func (m NewOrderCross) GetTimeInForce(f *field.TimeInForce) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for NewOrderCross.
func (m NewOrderCross) EffectiveTime() (*field.EffectiveTime, errors.MessageRejectError) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from NewOrderCross.
func (m NewOrderCross) GetEffectiveTime(f *field.EffectiveTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for NewOrderCross.
func (m NewOrderCross) ExpireDate() (*field.ExpireDate, errors.MessageRejectError) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from NewOrderCross.
func (m NewOrderCross) GetExpireDate(f *field.ExpireDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for NewOrderCross.
func (m NewOrderCross) ExpireTime() (*field.ExpireTime, errors.MessageRejectError) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from NewOrderCross.
func (m NewOrderCross) GetExpireTime(f *field.ExpireTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for NewOrderCross.
func (m NewOrderCross) GTBookingInst() (*field.GTBookingInst, errors.MessageRejectError) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from NewOrderCross.
func (m NewOrderCross) GetGTBookingInst(f *field.GTBookingInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for NewOrderCross.
func (m NewOrderCross) MaxShow() (*field.MaxShow, errors.MessageRejectError) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from NewOrderCross.
func (m NewOrderCross) GetMaxShow(f *field.MaxShow) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetValue is a non-required field for NewOrderCross.
func (m NewOrderCross) PegOffsetValue() (*field.PegOffsetValue, errors.MessageRejectError) {
	f := new(field.PegOffsetValue)
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetValue reads a PegOffsetValue from NewOrderCross.
func (m NewOrderCross) GetPegOffsetValue(f *field.PegOffsetValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegMoveType is a non-required field for NewOrderCross.
func (m NewOrderCross) PegMoveType() (*field.PegMoveType, errors.MessageRejectError) {
	f := new(field.PegMoveType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegMoveType reads a PegMoveType from NewOrderCross.
func (m NewOrderCross) GetPegMoveType(f *field.PegMoveType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetType is a non-required field for NewOrderCross.
func (m NewOrderCross) PegOffsetType() (*field.PegOffsetType, errors.MessageRejectError) {
	f := new(field.PegOffsetType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetType reads a PegOffsetType from NewOrderCross.
func (m NewOrderCross) GetPegOffsetType(f *field.PegOffsetType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegLimitType is a non-required field for NewOrderCross.
func (m NewOrderCross) PegLimitType() (*field.PegLimitType, errors.MessageRejectError) {
	f := new(field.PegLimitType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegLimitType reads a PegLimitType from NewOrderCross.
func (m NewOrderCross) GetPegLimitType(f *field.PegLimitType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegRoundDirection is a non-required field for NewOrderCross.
func (m NewOrderCross) PegRoundDirection() (*field.PegRoundDirection, errors.MessageRejectError) {
	f := new(field.PegRoundDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetPegRoundDirection reads a PegRoundDirection from NewOrderCross.
func (m NewOrderCross) GetPegRoundDirection(f *field.PegRoundDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegScope is a non-required field for NewOrderCross.
func (m NewOrderCross) PegScope() (*field.PegScope, errors.MessageRejectError) {
	f := new(field.PegScope)
	err := m.Body.Get(f)
	return f, err
}

//GetPegScope reads a PegScope from NewOrderCross.
func (m NewOrderCross) GetPegScope(f *field.PegScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegPriceType is a non-required field for NewOrderCross.
func (m NewOrderCross) PegPriceType() (*field.PegPriceType, errors.MessageRejectError) {
	f := new(field.PegPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegPriceType reads a PegPriceType from NewOrderCross.
func (m NewOrderCross) GetPegPriceType(f *field.PegPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityIDSource is a non-required field for NewOrderCross.
func (m NewOrderCross) PegSecurityIDSource() (*field.PegSecurityIDSource, errors.MessageRejectError) {
	f := new(field.PegSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityIDSource reads a PegSecurityIDSource from NewOrderCross.
func (m NewOrderCross) GetPegSecurityIDSource(f *field.PegSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityID is a non-required field for NewOrderCross.
func (m NewOrderCross) PegSecurityID() (*field.PegSecurityID, errors.MessageRejectError) {
	f := new(field.PegSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityID reads a PegSecurityID from NewOrderCross.
func (m NewOrderCross) GetPegSecurityID(f *field.PegSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSymbol is a non-required field for NewOrderCross.
func (m NewOrderCross) PegSymbol() (*field.PegSymbol, errors.MessageRejectError) {
	f := new(field.PegSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSymbol reads a PegSymbol from NewOrderCross.
func (m NewOrderCross) GetPegSymbol(f *field.PegSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityDesc is a non-required field for NewOrderCross.
func (m NewOrderCross) PegSecurityDesc() (*field.PegSecurityDesc, errors.MessageRejectError) {
	f := new(field.PegSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityDesc reads a PegSecurityDesc from NewOrderCross.
func (m NewOrderCross) GetPegSecurityDesc(f *field.PegSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for NewOrderCross.
func (m NewOrderCross) DiscretionInst() (*field.DiscretionInst, errors.MessageRejectError) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from NewOrderCross.
func (m NewOrderCross) GetDiscretionInst(f *field.DiscretionInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetValue is a non-required field for NewOrderCross.
func (m NewOrderCross) DiscretionOffsetValue() (*field.DiscretionOffsetValue, errors.MessageRejectError) {
	f := new(field.DiscretionOffsetValue)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetValue reads a DiscretionOffsetValue from NewOrderCross.
func (m NewOrderCross) GetDiscretionOffsetValue(f *field.DiscretionOffsetValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionMoveType is a non-required field for NewOrderCross.
func (m NewOrderCross) DiscretionMoveType() (*field.DiscretionMoveType, errors.MessageRejectError) {
	f := new(field.DiscretionMoveType)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionMoveType reads a DiscretionMoveType from NewOrderCross.
func (m NewOrderCross) GetDiscretionMoveType(f *field.DiscretionMoveType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetType is a non-required field for NewOrderCross.
func (m NewOrderCross) DiscretionOffsetType() (*field.DiscretionOffsetType, errors.MessageRejectError) {
	f := new(field.DiscretionOffsetType)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetType reads a DiscretionOffsetType from NewOrderCross.
func (m NewOrderCross) GetDiscretionOffsetType(f *field.DiscretionOffsetType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionLimitType is a non-required field for NewOrderCross.
func (m NewOrderCross) DiscretionLimitType() (*field.DiscretionLimitType, errors.MessageRejectError) {
	f := new(field.DiscretionLimitType)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionLimitType reads a DiscretionLimitType from NewOrderCross.
func (m NewOrderCross) GetDiscretionLimitType(f *field.DiscretionLimitType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionRoundDirection is a non-required field for NewOrderCross.
func (m NewOrderCross) DiscretionRoundDirection() (*field.DiscretionRoundDirection, errors.MessageRejectError) {
	f := new(field.DiscretionRoundDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionRoundDirection reads a DiscretionRoundDirection from NewOrderCross.
func (m NewOrderCross) GetDiscretionRoundDirection(f *field.DiscretionRoundDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionScope is a non-required field for NewOrderCross.
func (m NewOrderCross) DiscretionScope() (*field.DiscretionScope, errors.MessageRejectError) {
	f := new(field.DiscretionScope)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionScope reads a DiscretionScope from NewOrderCross.
func (m NewOrderCross) GetDiscretionScope(f *field.DiscretionScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategy is a non-required field for NewOrderCross.
func (m NewOrderCross) TargetStrategy() (*field.TargetStrategy, errors.MessageRejectError) {
	f := new(field.TargetStrategy)
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategy reads a TargetStrategy from NewOrderCross.
func (m NewOrderCross) GetTargetStrategy(f *field.TargetStrategy) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategyParameters is a non-required field for NewOrderCross.
func (m NewOrderCross) TargetStrategyParameters() (*field.TargetStrategyParameters, errors.MessageRejectError) {
	f := new(field.TargetStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategyParameters reads a TargetStrategyParameters from NewOrderCross.
func (m NewOrderCross) GetTargetStrategyParameters(f *field.TargetStrategyParameters) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParticipationRate is a non-required field for NewOrderCross.
func (m NewOrderCross) ParticipationRate() (*field.ParticipationRate, errors.MessageRejectError) {
	f := new(field.ParticipationRate)
	err := m.Body.Get(f)
	return f, err
}

//GetParticipationRate reads a ParticipationRate from NewOrderCross.
func (m NewOrderCross) GetParticipationRate(f *field.ParticipationRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for NewOrderCross.
func (m NewOrderCross) CancellationRights() (*field.CancellationRights, errors.MessageRejectError) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from NewOrderCross.
func (m NewOrderCross) GetCancellationRights(f *field.CancellationRights) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for NewOrderCross.
func (m NewOrderCross) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, errors.MessageRejectError) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from NewOrderCross.
func (m NewOrderCross) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for NewOrderCross.
func (m NewOrderCross) RegistID() (*field.RegistID, errors.MessageRejectError) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from NewOrderCross.
func (m NewOrderCross) GetRegistID(f *field.RegistID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for NewOrderCross.
func (m NewOrderCross) Designation() (*field.Designation, errors.MessageRejectError) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from NewOrderCross.
func (m NewOrderCross) GetDesignation(f *field.Designation) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStrategyParameters is a non-required field for NewOrderCross.
func (m NewOrderCross) NoStrategyParameters() (*field.NoStrategyParameters, errors.MessageRejectError) {
	f := new(field.NoStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStrategyParameters reads a NoStrategyParameters from NewOrderCross.
func (m NewOrderCross) GetNoStrategyParameters(f *field.NoStrategyParameters) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransBkdTime is a non-required field for NewOrderCross.
func (m NewOrderCross) TransBkdTime() (*field.TransBkdTime, errors.MessageRejectError) {
	f := new(field.TransBkdTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransBkdTime reads a TransBkdTime from NewOrderCross.
func (m NewOrderCross) GetTransBkdTime(f *field.TransBkdTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRootPartyIDs is a non-required field for NewOrderCross.
func (m NewOrderCross) NoRootPartyIDs() (*field.NoRootPartyIDs, errors.MessageRejectError) {
	f := new(field.NoRootPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRootPartyIDs reads a NoRootPartyIDs from NewOrderCross.
func (m NewOrderCross) GetNoRootPartyIDs(f *field.NoRootPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchIncrement is a non-required field for NewOrderCross.
func (m NewOrderCross) MatchIncrement() (*field.MatchIncrement, errors.MessageRejectError) {
	f := new(field.MatchIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMatchIncrement reads a MatchIncrement from NewOrderCross.
func (m NewOrderCross) GetMatchIncrement(f *field.MatchIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxPriceLevels is a non-required field for NewOrderCross.
func (m NewOrderCross) MaxPriceLevels() (*field.MaxPriceLevels, errors.MessageRejectError) {
	f := new(field.MaxPriceLevels)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxPriceLevels reads a MaxPriceLevels from NewOrderCross.
func (m NewOrderCross) GetMaxPriceLevels(f *field.MaxPriceLevels) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryDisplayQty is a non-required field for NewOrderCross.
func (m NewOrderCross) SecondaryDisplayQty() (*field.SecondaryDisplayQty, errors.MessageRejectError) {
	f := new(field.SecondaryDisplayQty)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryDisplayQty reads a SecondaryDisplayQty from NewOrderCross.
func (m NewOrderCross) GetSecondaryDisplayQty(f *field.SecondaryDisplayQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayWhen is a non-required field for NewOrderCross.
func (m NewOrderCross) DisplayWhen() (*field.DisplayWhen, errors.MessageRejectError) {
	f := new(field.DisplayWhen)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayWhen reads a DisplayWhen from NewOrderCross.
func (m NewOrderCross) GetDisplayWhen(f *field.DisplayWhen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMethod is a non-required field for NewOrderCross.
func (m NewOrderCross) DisplayMethod() (*field.DisplayMethod, errors.MessageRejectError) {
	f := new(field.DisplayMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMethod reads a DisplayMethod from NewOrderCross.
func (m NewOrderCross) GetDisplayMethod(f *field.DisplayMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayLowQty is a non-required field for NewOrderCross.
func (m NewOrderCross) DisplayLowQty() (*field.DisplayLowQty, errors.MessageRejectError) {
	f := new(field.DisplayLowQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayLowQty reads a DisplayLowQty from NewOrderCross.
func (m NewOrderCross) GetDisplayLowQty(f *field.DisplayLowQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayHighQty is a non-required field for NewOrderCross.
func (m NewOrderCross) DisplayHighQty() (*field.DisplayHighQty, errors.MessageRejectError) {
	f := new(field.DisplayHighQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayHighQty reads a DisplayHighQty from NewOrderCross.
func (m NewOrderCross) GetDisplayHighQty(f *field.DisplayHighQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMinIncr is a non-required field for NewOrderCross.
func (m NewOrderCross) DisplayMinIncr() (*field.DisplayMinIncr, errors.MessageRejectError) {
	f := new(field.DisplayMinIncr)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMinIncr reads a DisplayMinIncr from NewOrderCross.
func (m NewOrderCross) GetDisplayMinIncr(f *field.DisplayMinIncr) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefreshQty is a non-required field for NewOrderCross.
func (m NewOrderCross) RefreshQty() (*field.RefreshQty, errors.MessageRejectError) {
	f := new(field.RefreshQty)
	err := m.Body.Get(f)
	return f, err
}

//GetRefreshQty reads a RefreshQty from NewOrderCross.
func (m NewOrderCross) GetRefreshQty(f *field.RefreshQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayQty is a non-required field for NewOrderCross.
func (m NewOrderCross) DisplayQty() (*field.DisplayQty, errors.MessageRejectError) {
	f := new(field.DisplayQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayQty reads a DisplayQty from NewOrderCross.
func (m NewOrderCross) GetDisplayQty(f *field.DisplayQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceProtectionScope is a non-required field for NewOrderCross.
func (m NewOrderCross) PriceProtectionScope() (*field.PriceProtectionScope, errors.MessageRejectError) {
	f := new(field.PriceProtectionScope)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceProtectionScope reads a PriceProtectionScope from NewOrderCross.
func (m NewOrderCross) GetPriceProtectionScope(f *field.PriceProtectionScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerType is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerType() (*field.TriggerType, errors.MessageRejectError) {
	f := new(field.TriggerType)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerType reads a TriggerType from NewOrderCross.
func (m NewOrderCross) GetTriggerType(f *field.TriggerType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerAction is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerAction() (*field.TriggerAction, errors.MessageRejectError) {
	f := new(field.TriggerAction)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerAction reads a TriggerAction from NewOrderCross.
func (m NewOrderCross) GetTriggerAction(f *field.TriggerAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPrice is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerPrice() (*field.TriggerPrice, errors.MessageRejectError) {
	f := new(field.TriggerPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPrice reads a TriggerPrice from NewOrderCross.
func (m NewOrderCross) GetTriggerPrice(f *field.TriggerPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSymbol is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerSymbol() (*field.TriggerSymbol, errors.MessageRejectError) {
	f := new(field.TriggerSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSymbol reads a TriggerSymbol from NewOrderCross.
func (m NewOrderCross) GetTriggerSymbol(f *field.TriggerSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityID is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerSecurityID() (*field.TriggerSecurityID, errors.MessageRejectError) {
	f := new(field.TriggerSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityID reads a TriggerSecurityID from NewOrderCross.
func (m NewOrderCross) GetTriggerSecurityID(f *field.TriggerSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityIDSource is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerSecurityIDSource() (*field.TriggerSecurityIDSource, errors.MessageRejectError) {
	f := new(field.TriggerSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityIDSource reads a TriggerSecurityIDSource from NewOrderCross.
func (m NewOrderCross) GetTriggerSecurityIDSource(f *field.TriggerSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityDesc is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerSecurityDesc() (*field.TriggerSecurityDesc, errors.MessageRejectError) {
	f := new(field.TriggerSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityDesc reads a TriggerSecurityDesc from NewOrderCross.
func (m NewOrderCross) GetTriggerSecurityDesc(f *field.TriggerSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceType is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerPriceType() (*field.TriggerPriceType, errors.MessageRejectError) {
	f := new(field.TriggerPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceType reads a TriggerPriceType from NewOrderCross.
func (m NewOrderCross) GetTriggerPriceType(f *field.TriggerPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceTypeScope is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerPriceTypeScope() (*field.TriggerPriceTypeScope, errors.MessageRejectError) {
	f := new(field.TriggerPriceTypeScope)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceTypeScope reads a TriggerPriceTypeScope from NewOrderCross.
func (m NewOrderCross) GetTriggerPriceTypeScope(f *field.TriggerPriceTypeScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceDirection is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerPriceDirection() (*field.TriggerPriceDirection, errors.MessageRejectError) {
	f := new(field.TriggerPriceDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceDirection reads a TriggerPriceDirection from NewOrderCross.
func (m NewOrderCross) GetTriggerPriceDirection(f *field.TriggerPriceDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewPrice is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerNewPrice() (*field.TriggerNewPrice, errors.MessageRejectError) {
	f := new(field.TriggerNewPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewPrice reads a TriggerNewPrice from NewOrderCross.
func (m NewOrderCross) GetTriggerNewPrice(f *field.TriggerNewPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerOrderType is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerOrderType() (*field.TriggerOrderType, errors.MessageRejectError) {
	f := new(field.TriggerOrderType)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerOrderType reads a TriggerOrderType from NewOrderCross.
func (m NewOrderCross) GetTriggerOrderType(f *field.TriggerOrderType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewQty is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerNewQty() (*field.TriggerNewQty, errors.MessageRejectError) {
	f := new(field.TriggerNewQty)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewQty reads a TriggerNewQty from NewOrderCross.
func (m NewOrderCross) GetTriggerNewQty(f *field.TriggerNewQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionID is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerTradingSessionID() (*field.TriggerTradingSessionID, errors.MessageRejectError) {
	f := new(field.TriggerTradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionID reads a TriggerTradingSessionID from NewOrderCross.
func (m NewOrderCross) GetTriggerTradingSessionID(f *field.TriggerTradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionSubID is a non-required field for NewOrderCross.
func (m NewOrderCross) TriggerTradingSessionSubID() (*field.TriggerTradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TriggerTradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionSubID reads a TriggerTradingSessionSubID from NewOrderCross.
func (m NewOrderCross) GetTriggerTradingSessionSubID(f *field.TriggerTradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestinationIDSource is a non-required field for NewOrderCross.
func (m NewOrderCross) ExDestinationIDSource() (*field.ExDestinationIDSource, errors.MessageRejectError) {
	f := new(field.ExDestinationIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestinationIDSource reads a ExDestinationIDSource from NewOrderCross.
func (m NewOrderCross) GetExDestinationIDSource(f *field.ExDestinationIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}
