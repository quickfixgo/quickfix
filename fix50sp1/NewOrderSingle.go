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

//NewOrderSingle msg type = D.
type NewOrderSingle struct {
	message.Message
}

//NewOrderSingleBuilder builds NewOrderSingle messages.
type NewOrderSingleBuilder struct {
	message.MessageBuilder
}

//CreateNewOrderSingleBuilder returns an initialized NewOrderSingleBuilder with specified required fields.
func CreateNewOrderSingleBuilder(
	clordid field.ClOrdID,
	side field.Side,
	transacttime field.TransactTime,
	ordtype field.OrdType) NewOrderSingleBuilder {
	var builder NewOrderSingleBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.BuildMsgType("D"))
	builder.Body.Set(clordid)
	builder.Body.Set(side)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//ClOrdID is a required field for NewOrderSingle.
func (m NewOrderSingle) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from NewOrderSingle.
func (m NewOrderSingle) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecondaryClOrdID() (*field.SecondaryClOrdID, errors.MessageRejectError) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from NewOrderSingle.
func (m NewOrderSingle) GetSecondaryClOrdID(f *field.SecondaryClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ClOrdLinkID() (*field.ClOrdLinkID, errors.MessageRejectError) {
	f := new(field.ClOrdLinkID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from NewOrderSingle.
func (m NewOrderSingle) GetClOrdLinkID(f *field.ClOrdLinkID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from NewOrderSingle.
func (m NewOrderSingle) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TradeOriginationDate() (*field.TradeOriginationDate, errors.MessageRejectError) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from NewOrderSingle.
func (m NewOrderSingle) GetTradeOriginationDate(f *field.TradeOriginationDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from NewOrderSingle.
func (m NewOrderSingle) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from NewOrderSingle.
func (m NewOrderSingle) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AcctIDSource() (*field.AcctIDSource, errors.MessageRejectError) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from NewOrderSingle.
func (m NewOrderSingle) GetAcctIDSource(f *field.AcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from NewOrderSingle.
func (m NewOrderSingle) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayBookingInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DayBookingInst() (*field.DayBookingInst, errors.MessageRejectError) {
	f := new(field.DayBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//GetDayBookingInst reads a DayBookingInst from NewOrderSingle.
func (m NewOrderSingle) GetDayBookingInst(f *field.DayBookingInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingUnit is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BookingUnit() (*field.BookingUnit, errors.MessageRejectError) {
	f := new(field.BookingUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetBookingUnit reads a BookingUnit from NewOrderSingle.
func (m NewOrderSingle) GetBookingUnit(f *field.BookingUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreallocMethod is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PreallocMethod() (*field.PreallocMethod, errors.MessageRejectError) {
	f := new(field.PreallocMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPreallocMethod reads a PreallocMethod from NewOrderSingle.
func (m NewOrderSingle) GetPreallocMethod(f *field.PreallocMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AllocID() (*field.AllocID, errors.MessageRejectError) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from NewOrderSingle.
func (m NewOrderSingle) GetAllocID(f *field.AllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoAllocs() (*field.NoAllocs, errors.MessageRejectError) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from NewOrderSingle.
func (m NewOrderSingle) GetNoAllocs(f *field.NoAllocs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SettlType() (*field.SettlType, errors.MessageRejectError) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from NewOrderSingle.
func (m NewOrderSingle) GetSettlType(f *field.SettlType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SettlDate() (*field.SettlDate, errors.MessageRejectError) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from NewOrderSingle.
func (m NewOrderSingle) GetSettlDate(f *field.SettlDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashMargin is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CashMargin() (*field.CashMargin, errors.MessageRejectError) {
	f := new(field.CashMargin)
	err := m.Body.Get(f)
	return f, err
}

//GetCashMargin reads a CashMargin from NewOrderSingle.
func (m NewOrderSingle) GetCashMargin(f *field.CashMargin) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFeeIndicator is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ClearingFeeIndicator() (*field.ClearingFeeIndicator, errors.MessageRejectError) {
	f := new(field.ClearingFeeIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFeeIndicator reads a ClearingFeeIndicator from NewOrderSingle.
func (m NewOrderSingle) GetClearingFeeIndicator(f *field.ClearingFeeIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) HandlInst() (*field.HandlInst, errors.MessageRejectError) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from NewOrderSingle.
func (m NewOrderSingle) GetHandlInst(f *field.HandlInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExecInst() (*field.ExecInst, errors.MessageRejectError) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from NewOrderSingle.
func (m NewOrderSingle) GetExecInst(f *field.ExecInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MinQty() (*field.MinQty, errors.MessageRejectError) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from NewOrderSingle.
func (m NewOrderSingle) GetMinQty(f *field.MinQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaxFloor() (*field.MaxFloor, errors.MessageRejectError) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from NewOrderSingle.
func (m NewOrderSingle) GetMaxFloor(f *field.MaxFloor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExDestination() (*field.ExDestination, errors.MessageRejectError) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from NewOrderSingle.
func (m NewOrderSingle) GetExDestination(f *field.ExDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoTradingSessions() (*field.NoTradingSessions, errors.MessageRejectError) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from NewOrderSingle.
func (m NewOrderSingle) GetNoTradingSessions(f *field.NoTradingSessions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ProcessCode() (*field.ProcessCode, errors.MessageRejectError) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from NewOrderSingle.
func (m NewOrderSingle) GetProcessCode(f *field.ProcessCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from NewOrderSingle.
func (m NewOrderSingle) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from NewOrderSingle.
func (m NewOrderSingle) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from NewOrderSingle.
func (m NewOrderSingle) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from NewOrderSingle.
func (m NewOrderSingle) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from NewOrderSingle.
func (m NewOrderSingle) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from NewOrderSingle.
func (m NewOrderSingle) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from NewOrderSingle.
func (m NewOrderSingle) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from NewOrderSingle.
func (m NewOrderSingle) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from NewOrderSingle.
func (m NewOrderSingle) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from NewOrderSingle.
func (m NewOrderSingle) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from NewOrderSingle.
func (m NewOrderSingle) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from NewOrderSingle.
func (m NewOrderSingle) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from NewOrderSingle.
func (m NewOrderSingle) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from NewOrderSingle.
func (m NewOrderSingle) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from NewOrderSingle.
func (m NewOrderSingle) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from NewOrderSingle.
func (m NewOrderSingle) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from NewOrderSingle.
func (m NewOrderSingle) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from NewOrderSingle.
func (m NewOrderSingle) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for NewOrderSingle.
func (m NewOrderSingle) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from NewOrderSingle.
func (m NewOrderSingle) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from NewOrderSingle.
func (m NewOrderSingle) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from NewOrderSingle.
func (m NewOrderSingle) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for NewOrderSingle.
func (m NewOrderSingle) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from NewOrderSingle.
func (m NewOrderSingle) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from NewOrderSingle.
func (m NewOrderSingle) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from NewOrderSingle.
func (m NewOrderSingle) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from NewOrderSingle.
func (m NewOrderSingle) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from NewOrderSingle.
func (m NewOrderSingle) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from NewOrderSingle.
func (m NewOrderSingle) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from NewOrderSingle.
func (m NewOrderSingle) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from NewOrderSingle.
func (m NewOrderSingle) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from NewOrderSingle.
func (m NewOrderSingle) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from NewOrderSingle.
func (m NewOrderSingle) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from NewOrderSingle.
func (m NewOrderSingle) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from NewOrderSingle.
func (m NewOrderSingle) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from NewOrderSingle.
func (m NewOrderSingle) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from NewOrderSingle.
func (m NewOrderSingle) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from NewOrderSingle.
func (m NewOrderSingle) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from NewOrderSingle.
func (m NewOrderSingle) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from NewOrderSingle.
func (m NewOrderSingle) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from NewOrderSingle.
func (m NewOrderSingle) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from NewOrderSingle.
func (m NewOrderSingle) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from NewOrderSingle.
func (m NewOrderSingle) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from NewOrderSingle.
func (m NewOrderSingle) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from NewOrderSingle.
func (m NewOrderSingle) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from NewOrderSingle.
func (m NewOrderSingle) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for NewOrderSingle.
func (m NewOrderSingle) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from NewOrderSingle.
func (m NewOrderSingle) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from NewOrderSingle.
func (m NewOrderSingle) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from NewOrderSingle.
func (m NewOrderSingle) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from NewOrderSingle.
func (m NewOrderSingle) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from NewOrderSingle.
func (m NewOrderSingle) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from NewOrderSingle.
func (m NewOrderSingle) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from NewOrderSingle.
func (m NewOrderSingle) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for NewOrderSingle.
func (m NewOrderSingle) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from NewOrderSingle.
func (m NewOrderSingle) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from NewOrderSingle.
func (m NewOrderSingle) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from NewOrderSingle.
func (m NewOrderSingle) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityGroup() (*field.SecurityGroup, errors.MessageRejectError) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from NewOrderSingle.
func (m NewOrderSingle) GetSecurityGroup(f *field.SecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from NewOrderSingle.
func (m NewOrderSingle) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) UnitOfMeasureQty() (*field.UnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from NewOrderSingle.
func (m NewOrderSingle) GetUnitOfMeasureQty(f *field.UnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityXMLLen() (*field.SecurityXMLLen, errors.MessageRejectError) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from NewOrderSingle.
func (m NewOrderSingle) GetSecurityXMLLen(f *field.SecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityXML() (*field.SecurityXML, errors.MessageRejectError) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from NewOrderSingle.
func (m NewOrderSingle) GetSecurityXML(f *field.SecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecurityXMLSchema() (*field.SecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from NewOrderSingle.
func (m NewOrderSingle) GetSecurityXMLSchema(f *field.SecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ProductComplex() (*field.ProductComplex, errors.MessageRejectError) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from NewOrderSingle.
func (m NewOrderSingle) GetProductComplex(f *field.ProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from NewOrderSingle.
func (m NewOrderSingle) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from NewOrderSingle.
func (m NewOrderSingle) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SettlMethod() (*field.SettlMethod, errors.MessageRejectError) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from NewOrderSingle.
func (m NewOrderSingle) GetSettlMethod(f *field.SettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExerciseStyle() (*field.ExerciseStyle, errors.MessageRejectError) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from NewOrderSingle.
func (m NewOrderSingle) GetExerciseStyle(f *field.ExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayAmount is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OptPayAmount() (*field.OptPayAmount, errors.MessageRejectError) {
	f := new(field.OptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayAmount reads a OptPayAmount from NewOrderSingle.
func (m NewOrderSingle) GetOptPayAmount(f *field.OptPayAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PriceQuoteMethod() (*field.PriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from NewOrderSingle.
func (m NewOrderSingle) GetPriceQuoteMethod(f *field.PriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ListMethod() (*field.ListMethod, errors.MessageRejectError) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from NewOrderSingle.
func (m NewOrderSingle) GetListMethod(f *field.ListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CapPrice() (*field.CapPrice, errors.MessageRejectError) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from NewOrderSingle.
func (m NewOrderSingle) GetCapPrice(f *field.CapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for NewOrderSingle.
func (m NewOrderSingle) FloorPrice() (*field.FloorPrice, errors.MessageRejectError) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from NewOrderSingle.
func (m NewOrderSingle) GetFloorPrice(f *field.FloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from NewOrderSingle.
func (m NewOrderSingle) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for NewOrderSingle.
func (m NewOrderSingle) FlexibleIndicator() (*field.FlexibleIndicator, errors.MessageRejectError) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from NewOrderSingle.
func (m NewOrderSingle) GetFlexibleIndicator(f *field.FlexibleIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for NewOrderSingle.
func (m NewOrderSingle) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from NewOrderSingle.
func (m NewOrderSingle) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FuturesValuationMethod is a non-required field for NewOrderSingle.
func (m NewOrderSingle) FuturesValuationMethod() (*field.FuturesValuationMethod, errors.MessageRejectError) {
	f := new(field.FuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetFuturesValuationMethod reads a FuturesValuationMethod from NewOrderSingle.
func (m NewOrderSingle) GetFuturesValuationMethod(f *field.FuturesValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AgreementDesc() (*field.AgreementDesc, errors.MessageRejectError) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from NewOrderSingle.
func (m NewOrderSingle) GetAgreementDesc(f *field.AgreementDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AgreementID() (*field.AgreementID, errors.MessageRejectError) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from NewOrderSingle.
func (m NewOrderSingle) GetAgreementID(f *field.AgreementID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AgreementDate() (*field.AgreementDate, errors.MessageRejectError) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from NewOrderSingle.
func (m NewOrderSingle) GetAgreementDate(f *field.AgreementDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) AgreementCurrency() (*field.AgreementCurrency, errors.MessageRejectError) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from NewOrderSingle.
func (m NewOrderSingle) GetAgreementCurrency(f *field.AgreementCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TerminationType() (*field.TerminationType, errors.MessageRejectError) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from NewOrderSingle.
func (m NewOrderSingle) GetTerminationType(f *field.TerminationType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StartDate() (*field.StartDate, errors.MessageRejectError) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from NewOrderSingle.
func (m NewOrderSingle) GetStartDate(f *field.StartDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EndDate() (*field.EndDate, errors.MessageRejectError) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from NewOrderSingle.
func (m NewOrderSingle) GetEndDate(f *field.EndDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DeliveryType() (*field.DeliveryType, errors.MessageRejectError) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from NewOrderSingle.
func (m NewOrderSingle) GetDeliveryType(f *field.DeliveryType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MarginRatio() (*field.MarginRatio, errors.MessageRejectError) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from NewOrderSingle.
func (m NewOrderSingle) GetMarginRatio(f *field.MarginRatio) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from NewOrderSingle.
func (m NewOrderSingle) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PrevClosePx() (*field.PrevClosePx, errors.MessageRejectError) {
	f := new(field.PrevClosePx)
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from NewOrderSingle.
func (m NewOrderSingle) GetPrevClosePx(f *field.PrevClosePx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for NewOrderSingle.
func (m NewOrderSingle) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from NewOrderSingle.
func (m NewOrderSingle) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for NewOrderSingle.
func (m NewOrderSingle) LocateReqd() (*field.LocateReqd, errors.MessageRejectError) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from NewOrderSingle.
func (m NewOrderSingle) GetLocateReqd(f *field.LocateReqd) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for NewOrderSingle.
func (m NewOrderSingle) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from NewOrderSingle.
func (m NewOrderSingle) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoStipulations() (*field.NoStipulations, errors.MessageRejectError) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from NewOrderSingle.
func (m NewOrderSingle) GetNoStipulations(f *field.NoStipulations) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) QtyType() (*field.QtyType, errors.MessageRejectError) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from NewOrderSingle.
func (m NewOrderSingle) GetQtyType(f *field.QtyType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from NewOrderSingle.
func (m NewOrderSingle) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CashOrderQty() (*field.CashOrderQty, errors.MessageRejectError) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from NewOrderSingle.
func (m NewOrderSingle) GetCashOrderQty(f *field.CashOrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderPercent() (*field.OrderPercent, errors.MessageRejectError) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from NewOrderSingle.
func (m NewOrderSingle) GetOrderPercent(f *field.OrderPercent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RoundingDirection() (*field.RoundingDirection, errors.MessageRejectError) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from NewOrderSingle.
func (m NewOrderSingle) GetRoundingDirection(f *field.RoundingDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RoundingModulus() (*field.RoundingModulus, errors.MessageRejectError) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from NewOrderSingle.
func (m NewOrderSingle) GetRoundingModulus(f *field.RoundingModulus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for NewOrderSingle.
func (m NewOrderSingle) OrdType() (*field.OrdType, errors.MessageRejectError) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from NewOrderSingle.
func (m NewOrderSingle) GetOrdType(f *field.OrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from NewOrderSingle.
func (m NewOrderSingle) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from NewOrderSingle.
func (m NewOrderSingle) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for NewOrderSingle.
func (m NewOrderSingle) StopPx() (*field.StopPx, errors.MessageRejectError) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from NewOrderSingle.
func (m NewOrderSingle) GetStopPx(f *field.StopPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Spread() (*field.Spread, errors.MessageRejectError) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from NewOrderSingle.
func (m NewOrderSingle) GetSpread(f *field.Spread) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from NewOrderSingle.
func (m NewOrderSingle) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkCurveName() (*field.BenchmarkCurveName, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from NewOrderSingle.
func (m NewOrderSingle) GetBenchmarkCurveName(f *field.BenchmarkCurveName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, errors.MessageRejectError) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from NewOrderSingle.
func (m NewOrderSingle) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkPrice() (*field.BenchmarkPrice, errors.MessageRejectError) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from NewOrderSingle.
func (m NewOrderSingle) GetBenchmarkPrice(f *field.BenchmarkPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkPriceType() (*field.BenchmarkPriceType, errors.MessageRejectError) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from NewOrderSingle.
func (m NewOrderSingle) GetBenchmarkPriceType(f *field.BenchmarkPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkSecurityID() (*field.BenchmarkSecurityID, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from NewOrderSingle.
func (m NewOrderSingle) GetBenchmarkSecurityID(f *field.BenchmarkSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from NewOrderSingle.
func (m NewOrderSingle) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) YieldType() (*field.YieldType, errors.MessageRejectError) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from NewOrderSingle.
func (m NewOrderSingle) GetYieldType(f *field.YieldType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Yield() (*field.Yield, errors.MessageRejectError) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from NewOrderSingle.
func (m NewOrderSingle) GetYield(f *field.Yield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) YieldCalcDate() (*field.YieldCalcDate, errors.MessageRejectError) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from NewOrderSingle.
func (m NewOrderSingle) GetYieldCalcDate(f *field.YieldCalcDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) YieldRedemptionDate() (*field.YieldRedemptionDate, errors.MessageRejectError) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from NewOrderSingle.
func (m NewOrderSingle) GetYieldRedemptionDate(f *field.YieldRedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for NewOrderSingle.
func (m NewOrderSingle) YieldRedemptionPrice() (*field.YieldRedemptionPrice, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from NewOrderSingle.
func (m NewOrderSingle) GetYieldRedemptionPrice(f *field.YieldRedemptionPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from NewOrderSingle.
func (m NewOrderSingle) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from NewOrderSingle.
func (m NewOrderSingle) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ComplianceID() (*field.ComplianceID, errors.MessageRejectError) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from NewOrderSingle.
func (m NewOrderSingle) GetComplianceID(f *field.ComplianceID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SolicitedFlag is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SolicitedFlag() (*field.SolicitedFlag, errors.MessageRejectError) {
	f := new(field.SolicitedFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSolicitedFlag reads a SolicitedFlag from NewOrderSingle.
func (m NewOrderSingle) GetSolicitedFlag(f *field.SolicitedFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) IOIID() (*field.IOIID, errors.MessageRejectError) {
	f := new(field.IOIID)
	err := m.Body.Get(f)
	return f, err
}

//GetIOIID reads a IOIID from NewOrderSingle.
func (m NewOrderSingle) GetIOIID(f *field.IOIID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) QuoteID() (*field.QuoteID, errors.MessageRejectError) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from NewOrderSingle.
func (m NewOrderSingle) GetQuoteID(f *field.QuoteID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TimeInForce() (*field.TimeInForce, errors.MessageRejectError) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from NewOrderSingle.
func (m NewOrderSingle) GetTimeInForce(f *field.TimeInForce) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EffectiveTime() (*field.EffectiveTime, errors.MessageRejectError) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from NewOrderSingle.
func (m NewOrderSingle) GetEffectiveTime(f *field.EffectiveTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExpireDate() (*field.ExpireDate, errors.MessageRejectError) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from NewOrderSingle.
func (m NewOrderSingle) GetExpireDate(f *field.ExpireDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExpireTime() (*field.ExpireTime, errors.MessageRejectError) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from NewOrderSingle.
func (m NewOrderSingle) GetExpireTime(f *field.ExpireTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) GTBookingInst() (*field.GTBookingInst, errors.MessageRejectError) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from NewOrderSingle.
func (m NewOrderSingle) GetGTBookingInst(f *field.GTBookingInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Commission() (*field.Commission, errors.MessageRejectError) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from NewOrderSingle.
func (m NewOrderSingle) GetCommission(f *field.Commission) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CommType() (*field.CommType, errors.MessageRejectError) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from NewOrderSingle.
func (m NewOrderSingle) GetCommType(f *field.CommType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommCurrency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CommCurrency() (*field.CommCurrency, errors.MessageRejectError) {
	f := new(field.CommCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetCommCurrency reads a CommCurrency from NewOrderSingle.
func (m NewOrderSingle) GetCommCurrency(f *field.CommCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FundRenewWaiv is a non-required field for NewOrderSingle.
func (m NewOrderSingle) FundRenewWaiv() (*field.FundRenewWaiv, errors.MessageRejectError) {
	f := new(field.FundRenewWaiv)
	err := m.Body.Get(f)
	return f, err
}

//GetFundRenewWaiv reads a FundRenewWaiv from NewOrderSingle.
func (m NewOrderSingle) GetFundRenewWaiv(f *field.FundRenewWaiv) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderCapacity() (*field.OrderCapacity, errors.MessageRejectError) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from NewOrderSingle.
func (m NewOrderSingle) GetOrderCapacity(f *field.OrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderRestrictions() (*field.OrderRestrictions, errors.MessageRejectError) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from NewOrderSingle.
func (m NewOrderSingle) GetOrderRestrictions(f *field.OrderRestrictions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CustOrderCapacity() (*field.CustOrderCapacity, errors.MessageRejectError) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from NewOrderSingle.
func (m NewOrderSingle) GetCustOrderCapacity(f *field.CustOrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ForexReq() (*field.ForexReq, errors.MessageRejectError) {
	f := new(field.ForexReq)
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from NewOrderSingle.
func (m NewOrderSingle) GetForexReq(f *field.ForexReq) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SettlCurrency() (*field.SettlCurrency, errors.MessageRejectError) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from NewOrderSingle.
func (m NewOrderSingle) GetSettlCurrency(f *field.SettlCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) BookingType() (*field.BookingType, errors.MessageRejectError) {
	f := new(field.BookingType)
	err := m.Body.Get(f)
	return f, err
}

//GetBookingType reads a BookingType from NewOrderSingle.
func (m NewOrderSingle) GetBookingType(f *field.BookingType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from NewOrderSingle.
func (m NewOrderSingle) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from NewOrderSingle.
func (m NewOrderSingle) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for NewOrderSingle.
func (m NewOrderSingle) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from NewOrderSingle.
func (m NewOrderSingle) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate2 is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SettlDate2() (*field.SettlDate2, errors.MessageRejectError) {
	f := new(field.SettlDate2)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate2 reads a SettlDate2 from NewOrderSingle.
func (m NewOrderSingle) GetSettlDate2(f *field.SettlDate2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty2 is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderQty2() (*field.OrderQty2, errors.MessageRejectError) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty2 reads a OrderQty2 from NewOrderSingle.
func (m NewOrderSingle) GetOrderQty2(f *field.OrderQty2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price2 is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Price2() (*field.Price2, errors.MessageRejectError) {
	f := new(field.Price2)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice2 reads a Price2 from NewOrderSingle.
func (m NewOrderSingle) GetPrice2(f *field.Price2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PositionEffect() (*field.PositionEffect, errors.MessageRejectError) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from NewOrderSingle.
func (m NewOrderSingle) GetPositionEffect(f *field.PositionEffect) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CoveredOrUncovered is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CoveredOrUncovered() (*field.CoveredOrUncovered, errors.MessageRejectError) {
	f := new(field.CoveredOrUncovered)
	err := m.Body.Get(f)
	return f, err
}

//GetCoveredOrUncovered reads a CoveredOrUncovered from NewOrderSingle.
func (m NewOrderSingle) GetCoveredOrUncovered(f *field.CoveredOrUncovered) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaxShow() (*field.MaxShow, errors.MessageRejectError) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from NewOrderSingle.
func (m NewOrderSingle) GetMaxShow(f *field.MaxShow) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetValue is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegOffsetValue() (*field.PegOffsetValue, errors.MessageRejectError) {
	f := new(field.PegOffsetValue)
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetValue reads a PegOffsetValue from NewOrderSingle.
func (m NewOrderSingle) GetPegOffsetValue(f *field.PegOffsetValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegMoveType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegMoveType() (*field.PegMoveType, errors.MessageRejectError) {
	f := new(field.PegMoveType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegMoveType reads a PegMoveType from NewOrderSingle.
func (m NewOrderSingle) GetPegMoveType(f *field.PegMoveType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegOffsetType() (*field.PegOffsetType, errors.MessageRejectError) {
	f := new(field.PegOffsetType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetType reads a PegOffsetType from NewOrderSingle.
func (m NewOrderSingle) GetPegOffsetType(f *field.PegOffsetType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegLimitType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegLimitType() (*field.PegLimitType, errors.MessageRejectError) {
	f := new(field.PegLimitType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegLimitType reads a PegLimitType from NewOrderSingle.
func (m NewOrderSingle) GetPegLimitType(f *field.PegLimitType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegRoundDirection is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegRoundDirection() (*field.PegRoundDirection, errors.MessageRejectError) {
	f := new(field.PegRoundDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetPegRoundDirection reads a PegRoundDirection from NewOrderSingle.
func (m NewOrderSingle) GetPegRoundDirection(f *field.PegRoundDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegScope is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegScope() (*field.PegScope, errors.MessageRejectError) {
	f := new(field.PegScope)
	err := m.Body.Get(f)
	return f, err
}

//GetPegScope reads a PegScope from NewOrderSingle.
func (m NewOrderSingle) GetPegScope(f *field.PegScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegPriceType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegPriceType() (*field.PegPriceType, errors.MessageRejectError) {
	f := new(field.PegPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegPriceType reads a PegPriceType from NewOrderSingle.
func (m NewOrderSingle) GetPegPriceType(f *field.PegPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityIDSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegSecurityIDSource() (*field.PegSecurityIDSource, errors.MessageRejectError) {
	f := new(field.PegSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityIDSource reads a PegSecurityIDSource from NewOrderSingle.
func (m NewOrderSingle) GetPegSecurityIDSource(f *field.PegSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegSecurityID() (*field.PegSecurityID, errors.MessageRejectError) {
	f := new(field.PegSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityID reads a PegSecurityID from NewOrderSingle.
func (m NewOrderSingle) GetPegSecurityID(f *field.PegSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSymbol is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegSymbol() (*field.PegSymbol, errors.MessageRejectError) {
	f := new(field.PegSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSymbol reads a PegSymbol from NewOrderSingle.
func (m NewOrderSingle) GetPegSymbol(f *field.PegSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityDesc is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PegSecurityDesc() (*field.PegSecurityDesc, errors.MessageRejectError) {
	f := new(field.PegSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityDesc reads a PegSecurityDesc from NewOrderSingle.
func (m NewOrderSingle) GetPegSecurityDesc(f *field.PegSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DiscretionInst() (*field.DiscretionInst, errors.MessageRejectError) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from NewOrderSingle.
func (m NewOrderSingle) GetDiscretionInst(f *field.DiscretionInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetValue is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DiscretionOffsetValue() (*field.DiscretionOffsetValue, errors.MessageRejectError) {
	f := new(field.DiscretionOffsetValue)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetValue reads a DiscretionOffsetValue from NewOrderSingle.
func (m NewOrderSingle) GetDiscretionOffsetValue(f *field.DiscretionOffsetValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionMoveType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DiscretionMoveType() (*field.DiscretionMoveType, errors.MessageRejectError) {
	f := new(field.DiscretionMoveType)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionMoveType reads a DiscretionMoveType from NewOrderSingle.
func (m NewOrderSingle) GetDiscretionMoveType(f *field.DiscretionMoveType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DiscretionOffsetType() (*field.DiscretionOffsetType, errors.MessageRejectError) {
	f := new(field.DiscretionOffsetType)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetType reads a DiscretionOffsetType from NewOrderSingle.
func (m NewOrderSingle) GetDiscretionOffsetType(f *field.DiscretionOffsetType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionLimitType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DiscretionLimitType() (*field.DiscretionLimitType, errors.MessageRejectError) {
	f := new(field.DiscretionLimitType)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionLimitType reads a DiscretionLimitType from NewOrderSingle.
func (m NewOrderSingle) GetDiscretionLimitType(f *field.DiscretionLimitType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionRoundDirection is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DiscretionRoundDirection() (*field.DiscretionRoundDirection, errors.MessageRejectError) {
	f := new(field.DiscretionRoundDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionRoundDirection reads a DiscretionRoundDirection from NewOrderSingle.
func (m NewOrderSingle) GetDiscretionRoundDirection(f *field.DiscretionRoundDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionScope is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DiscretionScope() (*field.DiscretionScope, errors.MessageRejectError) {
	f := new(field.DiscretionScope)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionScope reads a DiscretionScope from NewOrderSingle.
func (m NewOrderSingle) GetDiscretionScope(f *field.DiscretionScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategy is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TargetStrategy() (*field.TargetStrategy, errors.MessageRejectError) {
	f := new(field.TargetStrategy)
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategy reads a TargetStrategy from NewOrderSingle.
func (m NewOrderSingle) GetTargetStrategy(f *field.TargetStrategy) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategyParameters is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TargetStrategyParameters() (*field.TargetStrategyParameters, errors.MessageRejectError) {
	f := new(field.TargetStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategyParameters reads a TargetStrategyParameters from NewOrderSingle.
func (m NewOrderSingle) GetTargetStrategyParameters(f *field.TargetStrategyParameters) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParticipationRate is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ParticipationRate() (*field.ParticipationRate, errors.MessageRejectError) {
	f := new(field.ParticipationRate)
	err := m.Body.Get(f)
	return f, err
}

//GetParticipationRate reads a ParticipationRate from NewOrderSingle.
func (m NewOrderSingle) GetParticipationRate(f *field.ParticipationRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CancellationRights() (*field.CancellationRights, errors.MessageRejectError) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from NewOrderSingle.
func (m NewOrderSingle) GetCancellationRights(f *field.CancellationRights) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, errors.MessageRejectError) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from NewOrderSingle.
func (m NewOrderSingle) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RegistID() (*field.RegistID, errors.MessageRejectError) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from NewOrderSingle.
func (m NewOrderSingle) GetRegistID(f *field.RegistID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for NewOrderSingle.
func (m NewOrderSingle) Designation() (*field.Designation, errors.MessageRejectError) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from NewOrderSingle.
func (m NewOrderSingle) GetDesignation(f *field.Designation) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStrategyParameters is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoStrategyParameters() (*field.NoStrategyParameters, errors.MessageRejectError) {
	f := new(field.NoStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStrategyParameters reads a NoStrategyParameters from NewOrderSingle.
func (m NewOrderSingle) GetNoStrategyParameters(f *field.NoStrategyParameters) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ManualOrderIndicator is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ManualOrderIndicator() (*field.ManualOrderIndicator, errors.MessageRejectError) {
	f := new(field.ManualOrderIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetManualOrderIndicator reads a ManualOrderIndicator from NewOrderSingle.
func (m NewOrderSingle) GetManualOrderIndicator(f *field.ManualOrderIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustDirectedOrder is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CustDirectedOrder() (*field.CustDirectedOrder, errors.MessageRejectError) {
	f := new(field.CustDirectedOrder)
	err := m.Body.Get(f)
	return f, err
}

//GetCustDirectedOrder reads a CustDirectedOrder from NewOrderSingle.
func (m NewOrderSingle) GetCustDirectedOrder(f *field.CustDirectedOrder) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ReceivedDeptID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ReceivedDeptID() (*field.ReceivedDeptID, errors.MessageRejectError) {
	f := new(field.ReceivedDeptID)
	err := m.Body.Get(f)
	return f, err
}

//GetReceivedDeptID reads a ReceivedDeptID from NewOrderSingle.
func (m NewOrderSingle) GetReceivedDeptID(f *field.ReceivedDeptID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderHandlingInst is a non-required field for NewOrderSingle.
func (m NewOrderSingle) CustOrderHandlingInst() (*field.CustOrderHandlingInst, errors.MessageRejectError) {
	f := new(field.CustOrderHandlingInst)
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderHandlingInst reads a CustOrderHandlingInst from NewOrderSingle.
func (m NewOrderSingle) GetCustOrderHandlingInst(f *field.CustOrderHandlingInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderHandlingInstSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) OrderHandlingInstSource() (*field.OrderHandlingInstSource, errors.MessageRejectError) {
	f := new(field.OrderHandlingInstSource)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderHandlingInstSource reads a OrderHandlingInstSource from NewOrderSingle.
func (m NewOrderSingle) GetOrderHandlingInstSource(f *field.OrderHandlingInstSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrdRegTimestamps is a non-required field for NewOrderSingle.
func (m NewOrderSingle) NoTrdRegTimestamps() (*field.NoTrdRegTimestamps, errors.MessageRejectError) {
	f := new(field.NoTrdRegTimestamps)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrdRegTimestamps reads a NoTrdRegTimestamps from NewOrderSingle.
func (m NewOrderSingle) GetNoTrdRegTimestamps(f *field.NoTrdRegTimestamps) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchIncrement is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MatchIncrement() (*field.MatchIncrement, errors.MessageRejectError) {
	f := new(field.MatchIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMatchIncrement reads a MatchIncrement from NewOrderSingle.
func (m NewOrderSingle) GetMatchIncrement(f *field.MatchIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxPriceLevels is a non-required field for NewOrderSingle.
func (m NewOrderSingle) MaxPriceLevels() (*field.MaxPriceLevels, errors.MessageRejectError) {
	f := new(field.MaxPriceLevels)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxPriceLevels reads a MaxPriceLevels from NewOrderSingle.
func (m NewOrderSingle) GetMaxPriceLevels(f *field.MaxPriceLevels) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryDisplayQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) SecondaryDisplayQty() (*field.SecondaryDisplayQty, errors.MessageRejectError) {
	f := new(field.SecondaryDisplayQty)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryDisplayQty reads a SecondaryDisplayQty from NewOrderSingle.
func (m NewOrderSingle) GetSecondaryDisplayQty(f *field.SecondaryDisplayQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayWhen is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DisplayWhen() (*field.DisplayWhen, errors.MessageRejectError) {
	f := new(field.DisplayWhen)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayWhen reads a DisplayWhen from NewOrderSingle.
func (m NewOrderSingle) GetDisplayWhen(f *field.DisplayWhen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMethod is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DisplayMethod() (*field.DisplayMethod, errors.MessageRejectError) {
	f := new(field.DisplayMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMethod reads a DisplayMethod from NewOrderSingle.
func (m NewOrderSingle) GetDisplayMethod(f *field.DisplayMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayLowQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DisplayLowQty() (*field.DisplayLowQty, errors.MessageRejectError) {
	f := new(field.DisplayLowQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayLowQty reads a DisplayLowQty from NewOrderSingle.
func (m NewOrderSingle) GetDisplayLowQty(f *field.DisplayLowQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayHighQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DisplayHighQty() (*field.DisplayHighQty, errors.MessageRejectError) {
	f := new(field.DisplayHighQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayHighQty reads a DisplayHighQty from NewOrderSingle.
func (m NewOrderSingle) GetDisplayHighQty(f *field.DisplayHighQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMinIncr is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DisplayMinIncr() (*field.DisplayMinIncr, errors.MessageRejectError) {
	f := new(field.DisplayMinIncr)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMinIncr reads a DisplayMinIncr from NewOrderSingle.
func (m NewOrderSingle) GetDisplayMinIncr(f *field.DisplayMinIncr) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefreshQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RefreshQty() (*field.RefreshQty, errors.MessageRejectError) {
	f := new(field.RefreshQty)
	err := m.Body.Get(f)
	return f, err
}

//GetRefreshQty reads a RefreshQty from NewOrderSingle.
func (m NewOrderSingle) GetRefreshQty(f *field.RefreshQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) DisplayQty() (*field.DisplayQty, errors.MessageRejectError) {
	f := new(field.DisplayQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayQty reads a DisplayQty from NewOrderSingle.
func (m NewOrderSingle) GetDisplayQty(f *field.DisplayQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceProtectionScope is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PriceProtectionScope() (*field.PriceProtectionScope, errors.MessageRejectError) {
	f := new(field.PriceProtectionScope)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceProtectionScope reads a PriceProtectionScope from NewOrderSingle.
func (m NewOrderSingle) GetPriceProtectionScope(f *field.PriceProtectionScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerType() (*field.TriggerType, errors.MessageRejectError) {
	f := new(field.TriggerType)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerType reads a TriggerType from NewOrderSingle.
func (m NewOrderSingle) GetTriggerType(f *field.TriggerType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerAction is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerAction() (*field.TriggerAction, errors.MessageRejectError) {
	f := new(field.TriggerAction)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerAction reads a TriggerAction from NewOrderSingle.
func (m NewOrderSingle) GetTriggerAction(f *field.TriggerAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPrice is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerPrice() (*field.TriggerPrice, errors.MessageRejectError) {
	f := new(field.TriggerPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPrice reads a TriggerPrice from NewOrderSingle.
func (m NewOrderSingle) GetTriggerPrice(f *field.TriggerPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSymbol is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerSymbol() (*field.TriggerSymbol, errors.MessageRejectError) {
	f := new(field.TriggerSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSymbol reads a TriggerSymbol from NewOrderSingle.
func (m NewOrderSingle) GetTriggerSymbol(f *field.TriggerSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerSecurityID() (*field.TriggerSecurityID, errors.MessageRejectError) {
	f := new(field.TriggerSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityID reads a TriggerSecurityID from NewOrderSingle.
func (m NewOrderSingle) GetTriggerSecurityID(f *field.TriggerSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityIDSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerSecurityIDSource() (*field.TriggerSecurityIDSource, errors.MessageRejectError) {
	f := new(field.TriggerSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityIDSource reads a TriggerSecurityIDSource from NewOrderSingle.
func (m NewOrderSingle) GetTriggerSecurityIDSource(f *field.TriggerSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityDesc is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerSecurityDesc() (*field.TriggerSecurityDesc, errors.MessageRejectError) {
	f := new(field.TriggerSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityDesc reads a TriggerSecurityDesc from NewOrderSingle.
func (m NewOrderSingle) GetTriggerSecurityDesc(f *field.TriggerSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerPriceType() (*field.TriggerPriceType, errors.MessageRejectError) {
	f := new(field.TriggerPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceType reads a TriggerPriceType from NewOrderSingle.
func (m NewOrderSingle) GetTriggerPriceType(f *field.TriggerPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceTypeScope is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerPriceTypeScope() (*field.TriggerPriceTypeScope, errors.MessageRejectError) {
	f := new(field.TriggerPriceTypeScope)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceTypeScope reads a TriggerPriceTypeScope from NewOrderSingle.
func (m NewOrderSingle) GetTriggerPriceTypeScope(f *field.TriggerPriceTypeScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceDirection is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerPriceDirection() (*field.TriggerPriceDirection, errors.MessageRejectError) {
	f := new(field.TriggerPriceDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceDirection reads a TriggerPriceDirection from NewOrderSingle.
func (m NewOrderSingle) GetTriggerPriceDirection(f *field.TriggerPriceDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewPrice is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerNewPrice() (*field.TriggerNewPrice, errors.MessageRejectError) {
	f := new(field.TriggerNewPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewPrice reads a TriggerNewPrice from NewOrderSingle.
func (m NewOrderSingle) GetTriggerNewPrice(f *field.TriggerNewPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerOrderType is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerOrderType() (*field.TriggerOrderType, errors.MessageRejectError) {
	f := new(field.TriggerOrderType)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerOrderType reads a TriggerOrderType from NewOrderSingle.
func (m NewOrderSingle) GetTriggerOrderType(f *field.TriggerOrderType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewQty is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerNewQty() (*field.TriggerNewQty, errors.MessageRejectError) {
	f := new(field.TriggerNewQty)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewQty reads a TriggerNewQty from NewOrderSingle.
func (m NewOrderSingle) GetTriggerNewQty(f *field.TriggerNewQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerTradingSessionID() (*field.TriggerTradingSessionID, errors.MessageRejectError) {
	f := new(field.TriggerTradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionID reads a TriggerTradingSessionID from NewOrderSingle.
func (m NewOrderSingle) GetTriggerTradingSessionID(f *field.TriggerTradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionSubID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) TriggerTradingSessionSubID() (*field.TriggerTradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TriggerTradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionSubID reads a TriggerTradingSessionSubID from NewOrderSingle.
func (m NewOrderSingle) GetTriggerTradingSessionSubID(f *field.TriggerTradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreTradeAnonymity is a non-required field for NewOrderSingle.
func (m NewOrderSingle) PreTradeAnonymity() (*field.PreTradeAnonymity, errors.MessageRejectError) {
	f := new(field.PreTradeAnonymity)
	err := m.Body.Get(f)
	return f, err
}

//GetPreTradeAnonymity reads a PreTradeAnonymity from NewOrderSingle.
func (m NewOrderSingle) GetPreTradeAnonymity(f *field.PreTradeAnonymity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefOrderID is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RefOrderID() (*field.RefOrderID, errors.MessageRejectError) {
	f := new(field.RefOrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetRefOrderID reads a RefOrderID from NewOrderSingle.
func (m NewOrderSingle) GetRefOrderID(f *field.RefOrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefOrderIDSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) RefOrderIDSource() (*field.RefOrderIDSource, errors.MessageRejectError) {
	f := new(field.RefOrderIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetRefOrderIDSource reads a RefOrderIDSource from NewOrderSingle.
func (m NewOrderSingle) GetRefOrderIDSource(f *field.RefOrderIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestinationIDSource is a non-required field for NewOrderSingle.
func (m NewOrderSingle) ExDestinationIDSource() (*field.ExDestinationIDSource, errors.MessageRejectError) {
	f := new(field.ExDestinationIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestinationIDSource reads a ExDestinationIDSource from NewOrderSingle.
func (m NewOrderSingle) GetExDestinationIDSource(f *field.ExDestinationIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}
