package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderCancelRequest msg type = F.
type OrderCancelRequest struct {
	message.Message
}

//OrderCancelRequestBuilder builds OrderCancelRequest messages.
type OrderCancelRequestBuilder struct {
	message.MessageBuilder
}

//CreateOrderCancelRequestBuilder returns an initialized OrderCancelRequestBuilder with specified required fields.
func CreateOrderCancelRequestBuilder(
	origclordid field.OrigClOrdID,
	clordid field.ClOrdID,
	side field.Side,
	transacttime field.TransactTime) OrderCancelRequestBuilder {
	var builder OrderCancelRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("F"))
	builder.Body.Set(origclordid)
	builder.Body.Set(clordid)
	builder.Body.Set(side)
	builder.Body.Set(transacttime)
	return builder
}

//OrigClOrdID is a required field for OrderCancelRequest.
func (m OrderCancelRequest) OrigClOrdID() (*field.OrigClOrdID, errors.MessageRejectError) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from OrderCancelRequest.
func (m OrderCancelRequest) GetOrigClOrdID(f *field.OrigClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from OrderCancelRequest.
func (m OrderCancelRequest) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for OrderCancelRequest.
func (m OrderCancelRequest) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderCancelRequest.
func (m OrderCancelRequest) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecondaryClOrdID() (*field.SecondaryClOrdID, errors.MessageRejectError) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from OrderCancelRequest.
func (m OrderCancelRequest) GetSecondaryClOrdID(f *field.SecondaryClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ClOrdLinkID() (*field.ClOrdLinkID, errors.MessageRejectError) {
	f := new(field.ClOrdLinkID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from OrderCancelRequest.
func (m OrderCancelRequest) GetClOrdLinkID(f *field.ClOrdLinkID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ListID() (*field.ListID, errors.MessageRejectError) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from OrderCancelRequest.
func (m OrderCancelRequest) GetListID(f *field.ListID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigOrdModTime is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) OrigOrdModTime() (*field.OrigOrdModTime, errors.MessageRejectError) {
	f := new(field.OrigOrdModTime)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigOrdModTime reads a OrigOrdModTime from OrderCancelRequest.
func (m OrderCancelRequest) GetOrigOrdModTime(f *field.OrigOrdModTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from OrderCancelRequest.
func (m OrderCancelRequest) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) AcctIDSource() (*field.AcctIDSource, errors.MessageRejectError) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from OrderCancelRequest.
func (m OrderCancelRequest) GetAcctIDSource(f *field.AcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from OrderCancelRequest.
func (m OrderCancelRequest) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from OrderCancelRequest.
func (m OrderCancelRequest) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from OrderCancelRequest.
func (m OrderCancelRequest) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from OrderCancelRequest.
func (m OrderCancelRequest) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from OrderCancelRequest.
func (m OrderCancelRequest) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from OrderCancelRequest.
func (m OrderCancelRequest) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from OrderCancelRequest.
func (m OrderCancelRequest) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from OrderCancelRequest.
func (m OrderCancelRequest) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from OrderCancelRequest.
func (m OrderCancelRequest) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from OrderCancelRequest.
func (m OrderCancelRequest) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from OrderCancelRequest.
func (m OrderCancelRequest) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from OrderCancelRequest.
func (m OrderCancelRequest) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from OrderCancelRequest.
func (m OrderCancelRequest) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from OrderCancelRequest.
func (m OrderCancelRequest) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from OrderCancelRequest.
func (m OrderCancelRequest) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from OrderCancelRequest.
func (m OrderCancelRequest) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from OrderCancelRequest.
func (m OrderCancelRequest) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from OrderCancelRequest.
func (m OrderCancelRequest) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from OrderCancelRequest.
func (m OrderCancelRequest) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from OrderCancelRequest.
func (m OrderCancelRequest) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from OrderCancelRequest.
func (m OrderCancelRequest) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from OrderCancelRequest.
func (m OrderCancelRequest) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from OrderCancelRequest.
func (m OrderCancelRequest) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from OrderCancelRequest.
func (m OrderCancelRequest) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from OrderCancelRequest.
func (m OrderCancelRequest) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from OrderCancelRequest.
func (m OrderCancelRequest) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from OrderCancelRequest.
func (m OrderCancelRequest) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from OrderCancelRequest.
func (m OrderCancelRequest) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from OrderCancelRequest.
func (m OrderCancelRequest) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from OrderCancelRequest.
func (m OrderCancelRequest) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from OrderCancelRequest.
func (m OrderCancelRequest) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from OrderCancelRequest.
func (m OrderCancelRequest) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from OrderCancelRequest.
func (m OrderCancelRequest) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from OrderCancelRequest.
func (m OrderCancelRequest) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from OrderCancelRequest.
func (m OrderCancelRequest) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from OrderCancelRequest.
func (m OrderCancelRequest) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from OrderCancelRequest.
func (m OrderCancelRequest) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from OrderCancelRequest.
func (m OrderCancelRequest) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from OrderCancelRequest.
func (m OrderCancelRequest) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from OrderCancelRequest.
func (m OrderCancelRequest) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from OrderCancelRequest.
func (m OrderCancelRequest) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from OrderCancelRequest.
func (m OrderCancelRequest) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from OrderCancelRequest.
func (m OrderCancelRequest) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from OrderCancelRequest.
func (m OrderCancelRequest) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from OrderCancelRequest.
func (m OrderCancelRequest) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from OrderCancelRequest.
func (m OrderCancelRequest) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from OrderCancelRequest.
func (m OrderCancelRequest) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from OrderCancelRequest.
func (m OrderCancelRequest) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from OrderCancelRequest.
func (m OrderCancelRequest) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from OrderCancelRequest.
func (m OrderCancelRequest) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from OrderCancelRequest.
func (m OrderCancelRequest) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from OrderCancelRequest.
func (m OrderCancelRequest) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from OrderCancelRequest.
func (m OrderCancelRequest) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from OrderCancelRequest.
func (m OrderCancelRequest) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from OrderCancelRequest.
func (m OrderCancelRequest) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from OrderCancelRequest.
func (m OrderCancelRequest) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) AgreementDesc() (*field.AgreementDesc, errors.MessageRejectError) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from OrderCancelRequest.
func (m OrderCancelRequest) GetAgreementDesc(f *field.AgreementDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) AgreementID() (*field.AgreementID, errors.MessageRejectError) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from OrderCancelRequest.
func (m OrderCancelRequest) GetAgreementID(f *field.AgreementID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) AgreementDate() (*field.AgreementDate, errors.MessageRejectError) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from OrderCancelRequest.
func (m OrderCancelRequest) GetAgreementDate(f *field.AgreementDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) AgreementCurrency() (*field.AgreementCurrency, errors.MessageRejectError) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from OrderCancelRequest.
func (m OrderCancelRequest) GetAgreementCurrency(f *field.AgreementCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) TerminationType() (*field.TerminationType, errors.MessageRejectError) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from OrderCancelRequest.
func (m OrderCancelRequest) GetTerminationType(f *field.TerminationType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) StartDate() (*field.StartDate, errors.MessageRejectError) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from OrderCancelRequest.
func (m OrderCancelRequest) GetStartDate(f *field.StartDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) EndDate() (*field.EndDate, errors.MessageRejectError) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from OrderCancelRequest.
func (m OrderCancelRequest) GetEndDate(f *field.EndDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) DeliveryType() (*field.DeliveryType, errors.MessageRejectError) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from OrderCancelRequest.
func (m OrderCancelRequest) GetDeliveryType(f *field.DeliveryType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) MarginRatio() (*field.MarginRatio, errors.MessageRejectError) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from OrderCancelRequest.
func (m OrderCancelRequest) GetMarginRatio(f *field.MarginRatio) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from OrderCancelRequest.
func (m OrderCancelRequest) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for OrderCancelRequest.
func (m OrderCancelRequest) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from OrderCancelRequest.
func (m OrderCancelRequest) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for OrderCancelRequest.
func (m OrderCancelRequest) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from OrderCancelRequest.
func (m OrderCancelRequest) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from OrderCancelRequest.
func (m OrderCancelRequest) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) CashOrderQty() (*field.CashOrderQty, errors.MessageRejectError) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from OrderCancelRequest.
func (m OrderCancelRequest) GetCashOrderQty(f *field.CashOrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) OrderPercent() (*field.OrderPercent, errors.MessageRejectError) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from OrderCancelRequest.
func (m OrderCancelRequest) GetOrderPercent(f *field.OrderPercent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) RoundingDirection() (*field.RoundingDirection, errors.MessageRejectError) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from OrderCancelRequest.
func (m OrderCancelRequest) GetRoundingDirection(f *field.RoundingDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) RoundingModulus() (*field.RoundingModulus, errors.MessageRejectError) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from OrderCancelRequest.
func (m OrderCancelRequest) GetRoundingModulus(f *field.RoundingModulus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) ComplianceID() (*field.ComplianceID, errors.MessageRejectError) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from OrderCancelRequest.
func (m OrderCancelRequest) GetComplianceID(f *field.ComplianceID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from OrderCancelRequest.
func (m OrderCancelRequest) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from OrderCancelRequest.
func (m OrderCancelRequest) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for OrderCancelRequest.
func (m OrderCancelRequest) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from OrderCancelRequest.
func (m OrderCancelRequest) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
