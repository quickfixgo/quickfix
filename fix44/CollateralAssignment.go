package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CollateralAssignment msg type = AY.
type CollateralAssignment struct {
	message.Message
}

//CollateralAssignmentBuilder builds CollateralAssignment messages.
type CollateralAssignmentBuilder struct {
	message.MessageBuilder
}

//CreateCollateralAssignmentBuilder returns an initialized CollateralAssignmentBuilder with specified required fields.
func CreateCollateralAssignmentBuilder(
	collasgnid field.CollAsgnID,
	collasgnreason field.CollAsgnReason,
	collasgntranstype field.CollAsgnTransType,
	transacttime field.TransactTime) CollateralAssignmentBuilder {
	var builder CollateralAssignmentBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("AY"))
	builder.Body.Set(collasgnid)
	builder.Body.Set(collasgnreason)
	builder.Body.Set(collasgntranstype)
	builder.Body.Set(transacttime)
	return builder
}

//CollAsgnID is a required field for CollateralAssignment.
func (m CollateralAssignment) CollAsgnID() (*field.CollAsgnID, errors.MessageRejectError) {
	f := new(field.CollAsgnID)
	err := m.Body.Get(f)
	return f, err
}

//GetCollAsgnID reads a CollAsgnID from CollateralAssignment.
func (m CollateralAssignment) GetCollAsgnID(f *field.CollAsgnID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CollReqID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CollReqID() (*field.CollReqID, errors.MessageRejectError) {
	f := new(field.CollReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetCollReqID reads a CollReqID from CollateralAssignment.
func (m CollateralAssignment) GetCollReqID(f *field.CollReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CollAsgnReason is a required field for CollateralAssignment.
func (m CollateralAssignment) CollAsgnReason() (*field.CollAsgnReason, errors.MessageRejectError) {
	f := new(field.CollAsgnReason)
	err := m.Body.Get(f)
	return f, err
}

//GetCollAsgnReason reads a CollAsgnReason from CollateralAssignment.
func (m CollateralAssignment) GetCollAsgnReason(f *field.CollAsgnReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CollAsgnTransType is a required field for CollateralAssignment.
func (m CollateralAssignment) CollAsgnTransType() (*field.CollAsgnTransType, errors.MessageRejectError) {
	f := new(field.CollAsgnTransType)
	err := m.Body.Get(f)
	return f, err
}

//GetCollAsgnTransType reads a CollAsgnTransType from CollateralAssignment.
func (m CollateralAssignment) GetCollAsgnTransType(f *field.CollAsgnTransType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CollAsgnRefID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CollAsgnRefID() (*field.CollAsgnRefID, errors.MessageRejectError) {
	f := new(field.CollAsgnRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetCollAsgnRefID reads a CollAsgnRefID from CollateralAssignment.
func (m CollateralAssignment) GetCollAsgnRefID(f *field.CollAsgnRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for CollateralAssignment.
func (m CollateralAssignment) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from CollateralAssignment.
func (m CollateralAssignment) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ExpireTime() (*field.ExpireTime, errors.MessageRejectError) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from CollateralAssignment.
func (m CollateralAssignment) GetExpireTime(f *field.ExpireTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from CollateralAssignment.
func (m CollateralAssignment) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from CollateralAssignment.
func (m CollateralAssignment) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from CollateralAssignment.
func (m CollateralAssignment) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from CollateralAssignment.
func (m CollateralAssignment) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from CollateralAssignment.
func (m CollateralAssignment) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecondaryOrderID() (*field.SecondaryOrderID, errors.MessageRejectError) {
	f := new(field.SecondaryOrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from CollateralAssignment.
func (m CollateralAssignment) GetSecondaryOrderID(f *field.SecondaryOrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecondaryClOrdID() (*field.SecondaryClOrdID, errors.MessageRejectError) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from CollateralAssignment.
func (m CollateralAssignment) GetSecondaryClOrdID(f *field.SecondaryClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecs is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoExecs() (*field.NoExecs, errors.MessageRejectError) {
	f := new(field.NoExecs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecs reads a NoExecs from CollateralAssignment.
func (m CollateralAssignment) GetNoExecs(f *field.NoExecs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrades is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoTrades() (*field.NoTrades, errors.MessageRejectError) {
	f := new(field.NoTrades)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrades reads a NoTrades from CollateralAssignment.
func (m CollateralAssignment) GetNoTrades(f *field.NoTrades) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from CollateralAssignment.
func (m CollateralAssignment) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from CollateralAssignment.
func (m CollateralAssignment) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from CollateralAssignment.
func (m CollateralAssignment) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from CollateralAssignment.
func (m CollateralAssignment) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from CollateralAssignment.
func (m CollateralAssignment) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from CollateralAssignment.
func (m CollateralAssignment) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from CollateralAssignment.
func (m CollateralAssignment) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from CollateralAssignment.
func (m CollateralAssignment) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from CollateralAssignment.
func (m CollateralAssignment) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from CollateralAssignment.
func (m CollateralAssignment) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from CollateralAssignment.
func (m CollateralAssignment) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from CollateralAssignment.
func (m CollateralAssignment) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from CollateralAssignment.
func (m CollateralAssignment) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from CollateralAssignment.
func (m CollateralAssignment) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for CollateralAssignment.
func (m CollateralAssignment) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from CollateralAssignment.
func (m CollateralAssignment) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from CollateralAssignment.
func (m CollateralAssignment) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from CollateralAssignment.
func (m CollateralAssignment) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from CollateralAssignment.
func (m CollateralAssignment) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for CollateralAssignment.
func (m CollateralAssignment) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from CollateralAssignment.
func (m CollateralAssignment) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from CollateralAssignment.
func (m CollateralAssignment) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from CollateralAssignment.
func (m CollateralAssignment) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for CollateralAssignment.
func (m CollateralAssignment) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from CollateralAssignment.
func (m CollateralAssignment) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from CollateralAssignment.
func (m CollateralAssignment) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from CollateralAssignment.
func (m CollateralAssignment) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from CollateralAssignment.
func (m CollateralAssignment) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for CollateralAssignment.
func (m CollateralAssignment) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from CollateralAssignment.
func (m CollateralAssignment) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from CollateralAssignment.
func (m CollateralAssignment) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from CollateralAssignment.
func (m CollateralAssignment) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from CollateralAssignment.
func (m CollateralAssignment) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from CollateralAssignment.
func (m CollateralAssignment) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from CollateralAssignment.
func (m CollateralAssignment) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from CollateralAssignment.
func (m CollateralAssignment) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from CollateralAssignment.
func (m CollateralAssignment) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from CollateralAssignment.
func (m CollateralAssignment) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from CollateralAssignment.
func (m CollateralAssignment) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from CollateralAssignment.
func (m CollateralAssignment) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from CollateralAssignment.
func (m CollateralAssignment) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from CollateralAssignment.
func (m CollateralAssignment) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from CollateralAssignment.
func (m CollateralAssignment) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from CollateralAssignment.
func (m CollateralAssignment) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from CollateralAssignment.
func (m CollateralAssignment) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from CollateralAssignment.
func (m CollateralAssignment) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AgreementDesc() (*field.AgreementDesc, errors.MessageRejectError) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from CollateralAssignment.
func (m CollateralAssignment) GetAgreementDesc(f *field.AgreementDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AgreementID() (*field.AgreementID, errors.MessageRejectError) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from CollateralAssignment.
func (m CollateralAssignment) GetAgreementID(f *field.AgreementID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AgreementDate() (*field.AgreementDate, errors.MessageRejectError) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from CollateralAssignment.
func (m CollateralAssignment) GetAgreementDate(f *field.AgreementDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AgreementCurrency() (*field.AgreementCurrency, errors.MessageRejectError) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from CollateralAssignment.
func (m CollateralAssignment) GetAgreementCurrency(f *field.AgreementCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) TerminationType() (*field.TerminationType, errors.MessageRejectError) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from CollateralAssignment.
func (m CollateralAssignment) GetTerminationType(f *field.TerminationType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StartDate() (*field.StartDate, errors.MessageRejectError) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from CollateralAssignment.
func (m CollateralAssignment) GetStartDate(f *field.StartDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EndDate() (*field.EndDate, errors.MessageRejectError) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from CollateralAssignment.
func (m CollateralAssignment) GetEndDate(f *field.EndDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) DeliveryType() (*field.DeliveryType, errors.MessageRejectError) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from CollateralAssignment.
func (m CollateralAssignment) GetDeliveryType(f *field.DeliveryType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MarginRatio() (*field.MarginRatio, errors.MessageRejectError) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from CollateralAssignment.
func (m CollateralAssignment) GetMarginRatio(f *field.MarginRatio) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SettlDate() (*field.SettlDate, errors.MessageRejectError) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from CollateralAssignment.
func (m CollateralAssignment) GetSettlDate(f *field.SettlDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Quantity is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Quantity() (*field.Quantity, errors.MessageRejectError) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}

//GetQuantity reads a Quantity from CollateralAssignment.
func (m CollateralAssignment) GetQuantity(f *field.Quantity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) QtyType() (*field.QtyType, errors.MessageRejectError) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from CollateralAssignment.
func (m CollateralAssignment) GetQtyType(f *field.QtyType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from CollateralAssignment.
func (m CollateralAssignment) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from CollateralAssignment.
func (m CollateralAssignment) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from CollateralAssignment.
func (m CollateralAssignment) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginExcess is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MarginExcess() (*field.MarginExcess, errors.MessageRejectError) {
	f := new(field.MarginExcess)
	err := m.Body.Get(f)
	return f, err
}

//GetMarginExcess reads a MarginExcess from CollateralAssignment.
func (m CollateralAssignment) GetMarginExcess(f *field.MarginExcess) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalNetValue is a non-required field for CollateralAssignment.
func (m CollateralAssignment) TotalNetValue() (*field.TotalNetValue, errors.MessageRejectError) {
	f := new(field.TotalNetValue)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalNetValue reads a TotalNetValue from CollateralAssignment.
func (m CollateralAssignment) GetTotalNetValue(f *field.TotalNetValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOutstanding is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CashOutstanding() (*field.CashOutstanding, errors.MessageRejectError) {
	f := new(field.CashOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetCashOutstanding reads a CashOutstanding from CollateralAssignment.
func (m CollateralAssignment) GetCashOutstanding(f *field.CashOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrdRegTimestamps is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoTrdRegTimestamps() (*field.NoTrdRegTimestamps, errors.MessageRejectError) {
	f := new(field.NoTrdRegTimestamps)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrdRegTimestamps reads a NoTrdRegTimestamps from CollateralAssignment.
func (m CollateralAssignment) GetNoTrdRegTimestamps(f *field.NoTrdRegTimestamps) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from CollateralAssignment.
func (m CollateralAssignment) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMiscFees is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoMiscFees() (*field.NoMiscFees, errors.MessageRejectError) {
	f := new(field.NoMiscFees)
	err := m.Body.Get(f)
	return f, err
}

//GetNoMiscFees reads a NoMiscFees from CollateralAssignment.
func (m CollateralAssignment) GetNoMiscFees(f *field.NoMiscFees) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from CollateralAssignment.
func (m CollateralAssignment) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from CollateralAssignment.
func (m CollateralAssignment) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestAmt is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AccruedInterestAmt() (*field.AccruedInterestAmt, errors.MessageRejectError) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestAmt reads a AccruedInterestAmt from CollateralAssignment.
func (m CollateralAssignment) GetAccruedInterestAmt(f *field.AccruedInterestAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndAccruedInterestAmt is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EndAccruedInterestAmt() (*field.EndAccruedInterestAmt, errors.MessageRejectError) {
	f := new(field.EndAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetEndAccruedInterestAmt reads a EndAccruedInterestAmt from CollateralAssignment.
func (m CollateralAssignment) GetEndAccruedInterestAmt(f *field.EndAccruedInterestAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartCash is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StartCash() (*field.StartCash, errors.MessageRejectError) {
	f := new(field.StartCash)
	err := m.Body.Get(f)
	return f, err
}

//GetStartCash reads a StartCash from CollateralAssignment.
func (m CollateralAssignment) GetStartCash(f *field.StartCash) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndCash is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EndCash() (*field.EndCash, errors.MessageRejectError) {
	f := new(field.EndCash)
	err := m.Body.Get(f)
	return f, err
}

//GetEndCash reads a EndCash from CollateralAssignment.
func (m CollateralAssignment) GetEndCash(f *field.EndCash) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Spread() (*field.Spread, errors.MessageRejectError) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from CollateralAssignment.
func (m CollateralAssignment) GetSpread(f *field.Spread) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from CollateralAssignment.
func (m CollateralAssignment) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkCurveName() (*field.BenchmarkCurveName, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from CollateralAssignment.
func (m CollateralAssignment) GetBenchmarkCurveName(f *field.BenchmarkCurveName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, errors.MessageRejectError) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from CollateralAssignment.
func (m CollateralAssignment) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkPrice() (*field.BenchmarkPrice, errors.MessageRejectError) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from CollateralAssignment.
func (m CollateralAssignment) GetBenchmarkPrice(f *field.BenchmarkPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkPriceType() (*field.BenchmarkPriceType, errors.MessageRejectError) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from CollateralAssignment.
func (m CollateralAssignment) GetBenchmarkPriceType(f *field.BenchmarkPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkSecurityID() (*field.BenchmarkSecurityID, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from CollateralAssignment.
func (m CollateralAssignment) GetBenchmarkSecurityID(f *field.BenchmarkSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from CollateralAssignment.
func (m CollateralAssignment) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoStipulations() (*field.NoStipulations, errors.MessageRejectError) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from CollateralAssignment.
func (m CollateralAssignment) GetNoStipulations(f *field.NoStipulations) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDeliveryType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SettlDeliveryType() (*field.SettlDeliveryType, errors.MessageRejectError) {
	f := new(field.SettlDeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDeliveryType reads a SettlDeliveryType from CollateralAssignment.
func (m CollateralAssignment) GetSettlDeliveryType(f *field.SettlDeliveryType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StandInstDbType() (*field.StandInstDbType, errors.MessageRejectError) {
	f := new(field.StandInstDbType)
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbType reads a StandInstDbType from CollateralAssignment.
func (m CollateralAssignment) GetStandInstDbType(f *field.StandInstDbType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbName is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StandInstDbName() (*field.StandInstDbName, errors.MessageRejectError) {
	f := new(field.StandInstDbName)
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbName reads a StandInstDbName from CollateralAssignment.
func (m CollateralAssignment) GetStandInstDbName(f *field.StandInstDbName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StandInstDbID() (*field.StandInstDbID, errors.MessageRejectError) {
	f := new(field.StandInstDbID)
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbID reads a StandInstDbID from CollateralAssignment.
func (m CollateralAssignment) GetStandInstDbID(f *field.StandInstDbID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDlvyInst is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoDlvyInst() (*field.NoDlvyInst, errors.MessageRejectError) {
	f := new(field.NoDlvyInst)
	err := m.Body.Get(f)
	return f, err
}

//GetNoDlvyInst reads a NoDlvyInst from CollateralAssignment.
func (m CollateralAssignment) GetNoDlvyInst(f *field.NoDlvyInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from CollateralAssignment.
func (m CollateralAssignment) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from CollateralAssignment.
func (m CollateralAssignment) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SettlSessID() (*field.SettlSessID, errors.MessageRejectError) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from CollateralAssignment.
func (m CollateralAssignment) GetSettlSessID(f *field.SettlSessID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SettlSessSubID() (*field.SettlSessSubID, errors.MessageRejectError) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from CollateralAssignment.
func (m CollateralAssignment) GetSettlSessSubID(f *field.SettlSessSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from CollateralAssignment.
func (m CollateralAssignment) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from CollateralAssignment.
func (m CollateralAssignment) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from CollateralAssignment.
func (m CollateralAssignment) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from CollateralAssignment.
func (m CollateralAssignment) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
