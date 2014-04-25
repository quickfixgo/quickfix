package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//NewOrderMultileg msg type = AB.
type NewOrderMultileg struct {
	message.Message
}

//NewOrderMultilegBuilder builds NewOrderMultileg messages.
type NewOrderMultilegBuilder struct {
	message.MessageBuilder
}

//CreateNewOrderMultilegBuilder returns an initialized NewOrderMultilegBuilder with specified required fields.
func CreateNewOrderMultilegBuilder(
	clordid field.ClOrdID,
	side field.Side,
	nolegs field.NoLegs,
	transacttime field.TransactTime,
	ordtype field.OrdType) NewOrderMultilegBuilder {
	var builder NewOrderMultilegBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("AB"))
	builder.Body.Set(clordid)
	builder.Body.Set(side)
	builder.Body.Set(nolegs)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//ClOrdID is a required field for NewOrderMultileg.
func (m NewOrderMultileg) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from NewOrderMultileg.
func (m NewOrderMultileg) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecondaryClOrdID() (*field.SecondaryClOrdID, errors.MessageRejectError) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from NewOrderMultileg.
func (m NewOrderMultileg) GetSecondaryClOrdID(f *field.SecondaryClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ClOrdLinkID() (*field.ClOrdLinkID, errors.MessageRejectError) {
	f := new(field.ClOrdLinkID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from NewOrderMultileg.
func (m NewOrderMultileg) GetClOrdLinkID(f *field.ClOrdLinkID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from NewOrderMultileg.
func (m NewOrderMultileg) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TradeOriginationDate() (*field.TradeOriginationDate, errors.MessageRejectError) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from NewOrderMultileg.
func (m NewOrderMultileg) GetTradeOriginationDate(f *field.TradeOriginationDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from NewOrderMultileg.
func (m NewOrderMultileg) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from NewOrderMultileg.
func (m NewOrderMultileg) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) AcctIDSource() (*field.AcctIDSource, errors.MessageRejectError) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from NewOrderMultileg.
func (m NewOrderMultileg) GetAcctIDSource(f *field.AcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from NewOrderMultileg.
func (m NewOrderMultileg) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayBookingInst is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DayBookingInst() (*field.DayBookingInst, errors.MessageRejectError) {
	f := new(field.DayBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//GetDayBookingInst reads a DayBookingInst from NewOrderMultileg.
func (m NewOrderMultileg) GetDayBookingInst(f *field.DayBookingInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingUnit is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) BookingUnit() (*field.BookingUnit, errors.MessageRejectError) {
	f := new(field.BookingUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetBookingUnit reads a BookingUnit from NewOrderMultileg.
func (m NewOrderMultileg) GetBookingUnit(f *field.BookingUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreallocMethod is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PreallocMethod() (*field.PreallocMethod, errors.MessageRejectError) {
	f := new(field.PreallocMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPreallocMethod reads a PreallocMethod from NewOrderMultileg.
func (m NewOrderMultileg) GetPreallocMethod(f *field.PreallocMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) AllocID() (*field.AllocID, errors.MessageRejectError) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from NewOrderMultileg.
func (m NewOrderMultileg) GetAllocID(f *field.AllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoAllocs() (*field.NoAllocs, errors.MessageRejectError) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from NewOrderMultileg.
func (m NewOrderMultileg) GetNoAllocs(f *field.NoAllocs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SettlType() (*field.SettlType, errors.MessageRejectError) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from NewOrderMultileg.
func (m NewOrderMultileg) GetSettlType(f *field.SettlType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SettlDate() (*field.SettlDate, errors.MessageRejectError) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from NewOrderMultileg.
func (m NewOrderMultileg) GetSettlDate(f *field.SettlDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashMargin is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CashMargin() (*field.CashMargin, errors.MessageRejectError) {
	f := new(field.CashMargin)
	err := m.Body.Get(f)
	return f, err
}

//GetCashMargin reads a CashMargin from NewOrderMultileg.
func (m NewOrderMultileg) GetCashMargin(f *field.CashMargin) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFeeIndicator is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ClearingFeeIndicator() (*field.ClearingFeeIndicator, errors.MessageRejectError) {
	f := new(field.ClearingFeeIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFeeIndicator reads a ClearingFeeIndicator from NewOrderMultileg.
func (m NewOrderMultileg) GetClearingFeeIndicator(f *field.ClearingFeeIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) HandlInst() (*field.HandlInst, errors.MessageRejectError) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from NewOrderMultileg.
func (m NewOrderMultileg) GetHandlInst(f *field.HandlInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ExecInst() (*field.ExecInst, errors.MessageRejectError) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from NewOrderMultileg.
func (m NewOrderMultileg) GetExecInst(f *field.ExecInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MinQty() (*field.MinQty, errors.MessageRejectError) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from NewOrderMultileg.
func (m NewOrderMultileg) GetMinQty(f *field.MinQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MaxFloor() (*field.MaxFloor, errors.MessageRejectError) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from NewOrderMultileg.
func (m NewOrderMultileg) GetMaxFloor(f *field.MaxFloor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ExDestination() (*field.ExDestination, errors.MessageRejectError) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from NewOrderMultileg.
func (m NewOrderMultileg) GetExDestination(f *field.ExDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoTradingSessions() (*field.NoTradingSessions, errors.MessageRejectError) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from NewOrderMultileg.
func (m NewOrderMultileg) GetNoTradingSessions(f *field.NoTradingSessions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ProcessCode() (*field.ProcessCode, errors.MessageRejectError) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from NewOrderMultileg.
func (m NewOrderMultileg) GetProcessCode(f *field.ProcessCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for NewOrderMultileg.
func (m NewOrderMultileg) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from NewOrderMultileg.
func (m NewOrderMultileg) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from NewOrderMultileg.
func (m NewOrderMultileg) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from NewOrderMultileg.
func (m NewOrderMultileg) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from NewOrderMultileg.
func (m NewOrderMultileg) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from NewOrderMultileg.
func (m NewOrderMultileg) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from NewOrderMultileg.
func (m NewOrderMultileg) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from NewOrderMultileg.
func (m NewOrderMultileg) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from NewOrderMultileg.
func (m NewOrderMultileg) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from NewOrderMultileg.
func (m NewOrderMultileg) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from NewOrderMultileg.
func (m NewOrderMultileg) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from NewOrderMultileg.
func (m NewOrderMultileg) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from NewOrderMultileg.
func (m NewOrderMultileg) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from NewOrderMultileg.
func (m NewOrderMultileg) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from NewOrderMultileg.
func (m NewOrderMultileg) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from NewOrderMultileg.
func (m NewOrderMultileg) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from NewOrderMultileg.
func (m NewOrderMultileg) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from NewOrderMultileg.
func (m NewOrderMultileg) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from NewOrderMultileg.
func (m NewOrderMultileg) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from NewOrderMultileg.
func (m NewOrderMultileg) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from NewOrderMultileg.
func (m NewOrderMultileg) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from NewOrderMultileg.
func (m NewOrderMultileg) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from NewOrderMultileg.
func (m NewOrderMultileg) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from NewOrderMultileg.
func (m NewOrderMultileg) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from NewOrderMultileg.
func (m NewOrderMultileg) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from NewOrderMultileg.
func (m NewOrderMultileg) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from NewOrderMultileg.
func (m NewOrderMultileg) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from NewOrderMultileg.
func (m NewOrderMultileg) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from NewOrderMultileg.
func (m NewOrderMultileg) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from NewOrderMultileg.
func (m NewOrderMultileg) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from NewOrderMultileg.
func (m NewOrderMultileg) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from NewOrderMultileg.
func (m NewOrderMultileg) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from NewOrderMultileg.
func (m NewOrderMultileg) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from NewOrderMultileg.
func (m NewOrderMultileg) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from NewOrderMultileg.
func (m NewOrderMultileg) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from NewOrderMultileg.
func (m NewOrderMultileg) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from NewOrderMultileg.
func (m NewOrderMultileg) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from NewOrderMultileg.
func (m NewOrderMultileg) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from NewOrderMultileg.
func (m NewOrderMultileg) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from NewOrderMultileg.
func (m NewOrderMultileg) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from NewOrderMultileg.
func (m NewOrderMultileg) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from NewOrderMultileg.
func (m NewOrderMultileg) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from NewOrderMultileg.
func (m NewOrderMultileg) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from NewOrderMultileg.
func (m NewOrderMultileg) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from NewOrderMultileg.
func (m NewOrderMultileg) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from NewOrderMultileg.
func (m NewOrderMultileg) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from NewOrderMultileg.
func (m NewOrderMultileg) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from NewOrderMultileg.
func (m NewOrderMultileg) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from NewOrderMultileg.
func (m NewOrderMultileg) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from NewOrderMultileg.
func (m NewOrderMultileg) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from NewOrderMultileg.
func (m NewOrderMultileg) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from NewOrderMultileg.
func (m NewOrderMultileg) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from NewOrderMultileg.
func (m NewOrderMultileg) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from NewOrderMultileg.
func (m NewOrderMultileg) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from NewOrderMultileg.
func (m NewOrderMultileg) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from NewOrderMultileg.
func (m NewOrderMultileg) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from NewOrderMultileg.
func (m NewOrderMultileg) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PrevClosePx() (*field.PrevClosePx, errors.MessageRejectError) {
	f := new(field.PrevClosePx)
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from NewOrderMultileg.
func (m NewOrderMultileg) GetPrevClosePx(f *field.PrevClosePx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a required field for NewOrderMultileg.
func (m NewOrderMultileg) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from NewOrderMultileg.
func (m NewOrderMultileg) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) LocateReqd() (*field.LocateReqd, errors.MessageRejectError) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from NewOrderMultileg.
func (m NewOrderMultileg) GetLocateReqd(f *field.LocateReqd) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for NewOrderMultileg.
func (m NewOrderMultileg) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from NewOrderMultileg.
func (m NewOrderMultileg) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) QtyType() (*field.QtyType, errors.MessageRejectError) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from NewOrderMultileg.
func (m NewOrderMultileg) GetQtyType(f *field.QtyType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from NewOrderMultileg.
func (m NewOrderMultileg) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CashOrderQty() (*field.CashOrderQty, errors.MessageRejectError) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from NewOrderMultileg.
func (m NewOrderMultileg) GetCashOrderQty(f *field.CashOrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OrderPercent() (*field.OrderPercent, errors.MessageRejectError) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from NewOrderMultileg.
func (m NewOrderMultileg) GetOrderPercent(f *field.OrderPercent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RoundingDirection() (*field.RoundingDirection, errors.MessageRejectError) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from NewOrderMultileg.
func (m NewOrderMultileg) GetRoundingDirection(f *field.RoundingDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RoundingModulus() (*field.RoundingModulus, errors.MessageRejectError) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from NewOrderMultileg.
func (m NewOrderMultileg) GetRoundingModulus(f *field.RoundingModulus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for NewOrderMultileg.
func (m NewOrderMultileg) OrdType() (*field.OrdType, errors.MessageRejectError) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from NewOrderMultileg.
func (m NewOrderMultileg) GetOrdType(f *field.OrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from NewOrderMultileg.
func (m NewOrderMultileg) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from NewOrderMultileg.
func (m NewOrderMultileg) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StopPx() (*field.StopPx, errors.MessageRejectError) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from NewOrderMultileg.
func (m NewOrderMultileg) GetStopPx(f *field.StopPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from NewOrderMultileg.
func (m NewOrderMultileg) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ComplianceID() (*field.ComplianceID, errors.MessageRejectError) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from NewOrderMultileg.
func (m NewOrderMultileg) GetComplianceID(f *field.ComplianceID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SolicitedFlag is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SolicitedFlag() (*field.SolicitedFlag, errors.MessageRejectError) {
	f := new(field.SolicitedFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSolicitedFlag reads a SolicitedFlag from NewOrderMultileg.
func (m NewOrderMultileg) GetSolicitedFlag(f *field.SolicitedFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) IOIID() (*field.IOIID, errors.MessageRejectError) {
	f := new(field.IOIID)
	err := m.Body.Get(f)
	return f, err
}

//GetIOIID reads a IOIID from NewOrderMultileg.
func (m NewOrderMultileg) GetIOIID(f *field.IOIID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) QuoteID() (*field.QuoteID, errors.MessageRejectError) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from NewOrderMultileg.
func (m NewOrderMultileg) GetQuoteID(f *field.QuoteID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TimeInForce() (*field.TimeInForce, errors.MessageRejectError) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from NewOrderMultileg.
func (m NewOrderMultileg) GetTimeInForce(f *field.TimeInForce) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EffectiveTime() (*field.EffectiveTime, errors.MessageRejectError) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from NewOrderMultileg.
func (m NewOrderMultileg) GetEffectiveTime(f *field.EffectiveTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ExpireDate() (*field.ExpireDate, errors.MessageRejectError) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from NewOrderMultileg.
func (m NewOrderMultileg) GetExpireDate(f *field.ExpireDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ExpireTime() (*field.ExpireTime, errors.MessageRejectError) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from NewOrderMultileg.
func (m NewOrderMultileg) GetExpireTime(f *field.ExpireTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) GTBookingInst() (*field.GTBookingInst, errors.MessageRejectError) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from NewOrderMultileg.
func (m NewOrderMultileg) GetGTBookingInst(f *field.GTBookingInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Commission() (*field.Commission, errors.MessageRejectError) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from NewOrderMultileg.
func (m NewOrderMultileg) GetCommission(f *field.Commission) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CommType() (*field.CommType, errors.MessageRejectError) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from NewOrderMultileg.
func (m NewOrderMultileg) GetCommType(f *field.CommType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommCurrency is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CommCurrency() (*field.CommCurrency, errors.MessageRejectError) {
	f := new(field.CommCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetCommCurrency reads a CommCurrency from NewOrderMultileg.
func (m NewOrderMultileg) GetCommCurrency(f *field.CommCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FundRenewWaiv is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) FundRenewWaiv() (*field.FundRenewWaiv, errors.MessageRejectError) {
	f := new(field.FundRenewWaiv)
	err := m.Body.Get(f)
	return f, err
}

//GetFundRenewWaiv reads a FundRenewWaiv from NewOrderMultileg.
func (m NewOrderMultileg) GetFundRenewWaiv(f *field.FundRenewWaiv) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OrderCapacity() (*field.OrderCapacity, errors.MessageRejectError) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from NewOrderMultileg.
func (m NewOrderMultileg) GetOrderCapacity(f *field.OrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OrderRestrictions() (*field.OrderRestrictions, errors.MessageRejectError) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from NewOrderMultileg.
func (m NewOrderMultileg) GetOrderRestrictions(f *field.OrderRestrictions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CustOrderCapacity() (*field.CustOrderCapacity, errors.MessageRejectError) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from NewOrderMultileg.
func (m NewOrderMultileg) GetCustOrderCapacity(f *field.CustOrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ForexReq() (*field.ForexReq, errors.MessageRejectError) {
	f := new(field.ForexReq)
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from NewOrderMultileg.
func (m NewOrderMultileg) GetForexReq(f *field.ForexReq) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SettlCurrency() (*field.SettlCurrency, errors.MessageRejectError) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from NewOrderMultileg.
func (m NewOrderMultileg) GetSettlCurrency(f *field.SettlCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) BookingType() (*field.BookingType, errors.MessageRejectError) {
	f := new(field.BookingType)
	err := m.Body.Get(f)
	return f, err
}

//GetBookingType reads a BookingType from NewOrderMultileg.
func (m NewOrderMultileg) GetBookingType(f *field.BookingType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from NewOrderMultileg.
func (m NewOrderMultileg) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from NewOrderMultileg.
func (m NewOrderMultileg) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from NewOrderMultileg.
func (m NewOrderMultileg) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PositionEffect() (*field.PositionEffect, errors.MessageRejectError) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from NewOrderMultileg.
func (m NewOrderMultileg) GetPositionEffect(f *field.PositionEffect) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CoveredOrUncovered is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CoveredOrUncovered() (*field.CoveredOrUncovered, errors.MessageRejectError) {
	f := new(field.CoveredOrUncovered)
	err := m.Body.Get(f)
	return f, err
}

//GetCoveredOrUncovered reads a CoveredOrUncovered from NewOrderMultileg.
func (m NewOrderMultileg) GetCoveredOrUncovered(f *field.CoveredOrUncovered) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MaxShow() (*field.MaxShow, errors.MessageRejectError) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from NewOrderMultileg.
func (m NewOrderMultileg) GetMaxShow(f *field.MaxShow) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetValue is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegOffsetValue() (*field.PegOffsetValue, errors.MessageRejectError) {
	f := new(field.PegOffsetValue)
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetValue reads a PegOffsetValue from NewOrderMultileg.
func (m NewOrderMultileg) GetPegOffsetValue(f *field.PegOffsetValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegMoveType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegMoveType() (*field.PegMoveType, errors.MessageRejectError) {
	f := new(field.PegMoveType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegMoveType reads a PegMoveType from NewOrderMultileg.
func (m NewOrderMultileg) GetPegMoveType(f *field.PegMoveType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegOffsetType() (*field.PegOffsetType, errors.MessageRejectError) {
	f := new(field.PegOffsetType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetType reads a PegOffsetType from NewOrderMultileg.
func (m NewOrderMultileg) GetPegOffsetType(f *field.PegOffsetType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegLimitType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegLimitType() (*field.PegLimitType, errors.MessageRejectError) {
	f := new(field.PegLimitType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegLimitType reads a PegLimitType from NewOrderMultileg.
func (m NewOrderMultileg) GetPegLimitType(f *field.PegLimitType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegRoundDirection is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegRoundDirection() (*field.PegRoundDirection, errors.MessageRejectError) {
	f := new(field.PegRoundDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetPegRoundDirection reads a PegRoundDirection from NewOrderMultileg.
func (m NewOrderMultileg) GetPegRoundDirection(f *field.PegRoundDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegScope is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegScope() (*field.PegScope, errors.MessageRejectError) {
	f := new(field.PegScope)
	err := m.Body.Get(f)
	return f, err
}

//GetPegScope reads a PegScope from NewOrderMultileg.
func (m NewOrderMultileg) GetPegScope(f *field.PegScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegPriceType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegPriceType() (*field.PegPriceType, errors.MessageRejectError) {
	f := new(field.PegPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegPriceType reads a PegPriceType from NewOrderMultileg.
func (m NewOrderMultileg) GetPegPriceType(f *field.PegPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityIDSource is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegSecurityIDSource() (*field.PegSecurityIDSource, errors.MessageRejectError) {
	f := new(field.PegSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityIDSource reads a PegSecurityIDSource from NewOrderMultileg.
func (m NewOrderMultileg) GetPegSecurityIDSource(f *field.PegSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegSecurityID() (*field.PegSecurityID, errors.MessageRejectError) {
	f := new(field.PegSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityID reads a PegSecurityID from NewOrderMultileg.
func (m NewOrderMultileg) GetPegSecurityID(f *field.PegSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSymbol is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegSymbol() (*field.PegSymbol, errors.MessageRejectError) {
	f := new(field.PegSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSymbol reads a PegSymbol from NewOrderMultileg.
func (m NewOrderMultileg) GetPegSymbol(f *field.PegSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityDesc is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegSecurityDesc() (*field.PegSecurityDesc, errors.MessageRejectError) {
	f := new(field.PegSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityDesc reads a PegSecurityDesc from NewOrderMultileg.
func (m NewOrderMultileg) GetPegSecurityDesc(f *field.PegSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DiscretionInst() (*field.DiscretionInst, errors.MessageRejectError) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from NewOrderMultileg.
func (m NewOrderMultileg) GetDiscretionInst(f *field.DiscretionInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetValue is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DiscretionOffsetValue() (*field.DiscretionOffsetValue, errors.MessageRejectError) {
	f := new(field.DiscretionOffsetValue)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetValue reads a DiscretionOffsetValue from NewOrderMultileg.
func (m NewOrderMultileg) GetDiscretionOffsetValue(f *field.DiscretionOffsetValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionMoveType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DiscretionMoveType() (*field.DiscretionMoveType, errors.MessageRejectError) {
	f := new(field.DiscretionMoveType)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionMoveType reads a DiscretionMoveType from NewOrderMultileg.
func (m NewOrderMultileg) GetDiscretionMoveType(f *field.DiscretionMoveType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DiscretionOffsetType() (*field.DiscretionOffsetType, errors.MessageRejectError) {
	f := new(field.DiscretionOffsetType)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetType reads a DiscretionOffsetType from NewOrderMultileg.
func (m NewOrderMultileg) GetDiscretionOffsetType(f *field.DiscretionOffsetType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionLimitType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DiscretionLimitType() (*field.DiscretionLimitType, errors.MessageRejectError) {
	f := new(field.DiscretionLimitType)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionLimitType reads a DiscretionLimitType from NewOrderMultileg.
func (m NewOrderMultileg) GetDiscretionLimitType(f *field.DiscretionLimitType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionRoundDirection is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DiscretionRoundDirection() (*field.DiscretionRoundDirection, errors.MessageRejectError) {
	f := new(field.DiscretionRoundDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionRoundDirection reads a DiscretionRoundDirection from NewOrderMultileg.
func (m NewOrderMultileg) GetDiscretionRoundDirection(f *field.DiscretionRoundDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionScope is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DiscretionScope() (*field.DiscretionScope, errors.MessageRejectError) {
	f := new(field.DiscretionScope)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionScope reads a DiscretionScope from NewOrderMultileg.
func (m NewOrderMultileg) GetDiscretionScope(f *field.DiscretionScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategy is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TargetStrategy() (*field.TargetStrategy, errors.MessageRejectError) {
	f := new(field.TargetStrategy)
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategy reads a TargetStrategy from NewOrderMultileg.
func (m NewOrderMultileg) GetTargetStrategy(f *field.TargetStrategy) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategyParameters is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TargetStrategyParameters() (*field.TargetStrategyParameters, errors.MessageRejectError) {
	f := new(field.TargetStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategyParameters reads a TargetStrategyParameters from NewOrderMultileg.
func (m NewOrderMultileg) GetTargetStrategyParameters(f *field.TargetStrategyParameters) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParticipationRate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ParticipationRate() (*field.ParticipationRate, errors.MessageRejectError) {
	f := new(field.ParticipationRate)
	err := m.Body.Get(f)
	return f, err
}

//GetParticipationRate reads a ParticipationRate from NewOrderMultileg.
func (m NewOrderMultileg) GetParticipationRate(f *field.ParticipationRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CancellationRights() (*field.CancellationRights, errors.MessageRejectError) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from NewOrderMultileg.
func (m NewOrderMultileg) GetCancellationRights(f *field.CancellationRights) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, errors.MessageRejectError) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from NewOrderMultileg.
func (m NewOrderMultileg) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RegistID() (*field.RegistID, errors.MessageRejectError) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from NewOrderMultileg.
func (m NewOrderMultileg) GetRegistID(f *field.RegistID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Designation() (*field.Designation, errors.MessageRejectError) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from NewOrderMultileg.
func (m NewOrderMultileg) GetDesignation(f *field.Designation) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegRptTypeReq is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MultiLegRptTypeReq() (*field.MultiLegRptTypeReq, errors.MessageRejectError) {
	f := new(field.MultiLegRptTypeReq)
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegRptTypeReq reads a MultiLegRptTypeReq from NewOrderMultileg.
func (m NewOrderMultileg) GetMultiLegRptTypeReq(f *field.MultiLegRptTypeReq) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStrategyParameters is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoStrategyParameters() (*field.NoStrategyParameters, errors.MessageRejectError) {
	f := new(field.NoStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStrategyParameters reads a NoStrategyParameters from NewOrderMultileg.
func (m NewOrderMultileg) GetNoStrategyParameters(f *field.NoStrategyParameters) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SwapPoints is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SwapPoints() (*field.SwapPoints, errors.MessageRejectError) {
	f := new(field.SwapPoints)
	err := m.Body.Get(f)
	return f, err
}

//GetSwapPoints reads a SwapPoints from NewOrderMultileg.
func (m NewOrderMultileg) GetSwapPoints(f *field.SwapPoints) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchIncrement is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MatchIncrement() (*field.MatchIncrement, errors.MessageRejectError) {
	f := new(field.MatchIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMatchIncrement reads a MatchIncrement from NewOrderMultileg.
func (m NewOrderMultileg) GetMatchIncrement(f *field.MatchIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxPriceLevels is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MaxPriceLevels() (*field.MaxPriceLevels, errors.MessageRejectError) {
	f := new(field.MaxPriceLevels)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxPriceLevels reads a MaxPriceLevels from NewOrderMultileg.
func (m NewOrderMultileg) GetMaxPriceLevels(f *field.MaxPriceLevels) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryDisplayQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecondaryDisplayQty() (*field.SecondaryDisplayQty, errors.MessageRejectError) {
	f := new(field.SecondaryDisplayQty)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryDisplayQty reads a SecondaryDisplayQty from NewOrderMultileg.
func (m NewOrderMultileg) GetSecondaryDisplayQty(f *field.SecondaryDisplayQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayWhen is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DisplayWhen() (*field.DisplayWhen, errors.MessageRejectError) {
	f := new(field.DisplayWhen)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayWhen reads a DisplayWhen from NewOrderMultileg.
func (m NewOrderMultileg) GetDisplayWhen(f *field.DisplayWhen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMethod is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DisplayMethod() (*field.DisplayMethod, errors.MessageRejectError) {
	f := new(field.DisplayMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMethod reads a DisplayMethod from NewOrderMultileg.
func (m NewOrderMultileg) GetDisplayMethod(f *field.DisplayMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayLowQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DisplayLowQty() (*field.DisplayLowQty, errors.MessageRejectError) {
	f := new(field.DisplayLowQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayLowQty reads a DisplayLowQty from NewOrderMultileg.
func (m NewOrderMultileg) GetDisplayLowQty(f *field.DisplayLowQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayHighQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DisplayHighQty() (*field.DisplayHighQty, errors.MessageRejectError) {
	f := new(field.DisplayHighQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayHighQty reads a DisplayHighQty from NewOrderMultileg.
func (m NewOrderMultileg) GetDisplayHighQty(f *field.DisplayHighQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMinIncr is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DisplayMinIncr() (*field.DisplayMinIncr, errors.MessageRejectError) {
	f := new(field.DisplayMinIncr)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMinIncr reads a DisplayMinIncr from NewOrderMultileg.
func (m NewOrderMultileg) GetDisplayMinIncr(f *field.DisplayMinIncr) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefreshQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RefreshQty() (*field.RefreshQty, errors.MessageRejectError) {
	f := new(field.RefreshQty)
	err := m.Body.Get(f)
	return f, err
}

//GetRefreshQty reads a RefreshQty from NewOrderMultileg.
func (m NewOrderMultileg) GetRefreshQty(f *field.RefreshQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DisplayQty() (*field.DisplayQty, errors.MessageRejectError) {
	f := new(field.DisplayQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayQty reads a DisplayQty from NewOrderMultileg.
func (m NewOrderMultileg) GetDisplayQty(f *field.DisplayQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceProtectionScope is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PriceProtectionScope() (*field.PriceProtectionScope, errors.MessageRejectError) {
	f := new(field.PriceProtectionScope)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceProtectionScope reads a PriceProtectionScope from NewOrderMultileg.
func (m NewOrderMultileg) GetPriceProtectionScope(f *field.PriceProtectionScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerType() (*field.TriggerType, errors.MessageRejectError) {
	f := new(field.TriggerType)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerType reads a TriggerType from NewOrderMultileg.
func (m NewOrderMultileg) GetTriggerType(f *field.TriggerType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerAction is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerAction() (*field.TriggerAction, errors.MessageRejectError) {
	f := new(field.TriggerAction)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerAction reads a TriggerAction from NewOrderMultileg.
func (m NewOrderMultileg) GetTriggerAction(f *field.TriggerAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPrice is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerPrice() (*field.TriggerPrice, errors.MessageRejectError) {
	f := new(field.TriggerPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPrice reads a TriggerPrice from NewOrderMultileg.
func (m NewOrderMultileg) GetTriggerPrice(f *field.TriggerPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSymbol is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerSymbol() (*field.TriggerSymbol, errors.MessageRejectError) {
	f := new(field.TriggerSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSymbol reads a TriggerSymbol from NewOrderMultileg.
func (m NewOrderMultileg) GetTriggerSymbol(f *field.TriggerSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerSecurityID() (*field.TriggerSecurityID, errors.MessageRejectError) {
	f := new(field.TriggerSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityID reads a TriggerSecurityID from NewOrderMultileg.
func (m NewOrderMultileg) GetTriggerSecurityID(f *field.TriggerSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityIDSource is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerSecurityIDSource() (*field.TriggerSecurityIDSource, errors.MessageRejectError) {
	f := new(field.TriggerSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityIDSource reads a TriggerSecurityIDSource from NewOrderMultileg.
func (m NewOrderMultileg) GetTriggerSecurityIDSource(f *field.TriggerSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityDesc is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerSecurityDesc() (*field.TriggerSecurityDesc, errors.MessageRejectError) {
	f := new(field.TriggerSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityDesc reads a TriggerSecurityDesc from NewOrderMultileg.
func (m NewOrderMultileg) GetTriggerSecurityDesc(f *field.TriggerSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerPriceType() (*field.TriggerPriceType, errors.MessageRejectError) {
	f := new(field.TriggerPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceType reads a TriggerPriceType from NewOrderMultileg.
func (m NewOrderMultileg) GetTriggerPriceType(f *field.TriggerPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceTypeScope is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerPriceTypeScope() (*field.TriggerPriceTypeScope, errors.MessageRejectError) {
	f := new(field.TriggerPriceTypeScope)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceTypeScope reads a TriggerPriceTypeScope from NewOrderMultileg.
func (m NewOrderMultileg) GetTriggerPriceTypeScope(f *field.TriggerPriceTypeScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceDirection is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerPriceDirection() (*field.TriggerPriceDirection, errors.MessageRejectError) {
	f := new(field.TriggerPriceDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceDirection reads a TriggerPriceDirection from NewOrderMultileg.
func (m NewOrderMultileg) GetTriggerPriceDirection(f *field.TriggerPriceDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewPrice is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerNewPrice() (*field.TriggerNewPrice, errors.MessageRejectError) {
	f := new(field.TriggerNewPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewPrice reads a TriggerNewPrice from NewOrderMultileg.
func (m NewOrderMultileg) GetTriggerNewPrice(f *field.TriggerNewPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerOrderType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerOrderType() (*field.TriggerOrderType, errors.MessageRejectError) {
	f := new(field.TriggerOrderType)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerOrderType reads a TriggerOrderType from NewOrderMultileg.
func (m NewOrderMultileg) GetTriggerOrderType(f *field.TriggerOrderType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerNewQty() (*field.TriggerNewQty, errors.MessageRejectError) {
	f := new(field.TriggerNewQty)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewQty reads a TriggerNewQty from NewOrderMultileg.
func (m NewOrderMultileg) GetTriggerNewQty(f *field.TriggerNewQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerTradingSessionID() (*field.TriggerTradingSessionID, errors.MessageRejectError) {
	f := new(field.TriggerTradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionID reads a TriggerTradingSessionID from NewOrderMultileg.
func (m NewOrderMultileg) GetTriggerTradingSessionID(f *field.TriggerTradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionSubID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TriggerTradingSessionSubID() (*field.TriggerTradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TriggerTradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionSubID reads a TriggerTradingSessionSubID from NewOrderMultileg.
func (m NewOrderMultileg) GetTriggerTradingSessionSubID(f *field.TriggerTradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefOrderID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RefOrderID() (*field.RefOrderID, errors.MessageRejectError) {
	f := new(field.RefOrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetRefOrderID reads a RefOrderID from NewOrderMultileg.
func (m NewOrderMultileg) GetRefOrderID(f *field.RefOrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefOrderIDSource is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RefOrderIDSource() (*field.RefOrderIDSource, errors.MessageRejectError) {
	f := new(field.RefOrderIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetRefOrderIDSource reads a RefOrderIDSource from NewOrderMultileg.
func (m NewOrderMultileg) GetRefOrderIDSource(f *field.RefOrderIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreTradeAnonymity is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PreTradeAnonymity() (*field.PreTradeAnonymity, errors.MessageRejectError) {
	f := new(field.PreTradeAnonymity)
	err := m.Body.Get(f)
	return f, err
}

//GetPreTradeAnonymity reads a PreTradeAnonymity from NewOrderMultileg.
func (m NewOrderMultileg) GetPreTradeAnonymity(f *field.PreTradeAnonymity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestinationIDSource is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ExDestinationIDSource() (*field.ExDestinationIDSource, errors.MessageRejectError) {
	f := new(field.ExDestinationIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestinationIDSource reads a ExDestinationIDSource from NewOrderMultileg.
func (m NewOrderMultileg) GetExDestinationIDSource(f *field.ExDestinationIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}
