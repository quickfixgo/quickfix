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

//MultilegOrderCancelReplace msg type = AC.
type MultilegOrderCancelReplace struct {
	message.Message
}

//MultilegOrderCancelReplaceBuilder builds MultilegOrderCancelReplace messages.
type MultilegOrderCancelReplaceBuilder struct {
	message.MessageBuilder
}

//CreateMultilegOrderCancelReplaceBuilder returns an initialized MultilegOrderCancelReplaceBuilder with specified required fields.
func CreateMultilegOrderCancelReplaceBuilder(
	side field.Side,
	nolegs field.NoLegs,
	transacttime field.TransactTime,
	ordtype field.OrdType) MultilegOrderCancelReplaceBuilder {
	var builder MultilegOrderCancelReplaceBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.BuildMsgType("AC"))
	builder.Body.Set(side)
	builder.Body.Set(nolegs)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//OrderID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrigClOrdID() (*field.OrigClOrdID, errors.MessageRejectError) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOrigClOrdID(f *field.OrigClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecondaryClOrdID() (*field.SecondaryClOrdID, errors.MessageRejectError) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecondaryClOrdID(f *field.SecondaryClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ClOrdLinkID() (*field.ClOrdLinkID, errors.MessageRejectError) {
	f := new(field.ClOrdLinkID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetClOrdLinkID(f *field.ClOrdLinkID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigOrdModTime is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrigOrdModTime() (*field.OrigOrdModTime, errors.MessageRejectError) {
	f := new(field.OrigOrdModTime)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigOrdModTime reads a OrigOrdModTime from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOrigOrdModTime(f *field.OrigOrdModTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TradeOriginationDate() (*field.TradeOriginationDate, errors.MessageRejectError) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTradeOriginationDate(f *field.TradeOriginationDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) AcctIDSource() (*field.AcctIDSource, errors.MessageRejectError) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetAcctIDSource(f *field.AcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayBookingInst is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DayBookingInst() (*field.DayBookingInst, errors.MessageRejectError) {
	f := new(field.DayBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//GetDayBookingInst reads a DayBookingInst from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDayBookingInst(f *field.DayBookingInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingUnit is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) BookingUnit() (*field.BookingUnit, errors.MessageRejectError) {
	f := new(field.BookingUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetBookingUnit reads a BookingUnit from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetBookingUnit(f *field.BookingUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreallocMethod is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PreallocMethod() (*field.PreallocMethod, errors.MessageRejectError) {
	f := new(field.PreallocMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPreallocMethod reads a PreallocMethod from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPreallocMethod(f *field.PreallocMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) AllocID() (*field.AllocID, errors.MessageRejectError) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetAllocID(f *field.AllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoAllocs() (*field.NoAllocs, errors.MessageRejectError) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoAllocs(f *field.NoAllocs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SettlType() (*field.SettlType, errors.MessageRejectError) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSettlType(f *field.SettlType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SettlDate() (*field.SettlDate, errors.MessageRejectError) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSettlDate(f *field.SettlDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashMargin is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CashMargin() (*field.CashMargin, errors.MessageRejectError) {
	f := new(field.CashMargin)
	err := m.Body.Get(f)
	return f, err
}

//GetCashMargin reads a CashMargin from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCashMargin(f *field.CashMargin) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFeeIndicator is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ClearingFeeIndicator() (*field.ClearingFeeIndicator, errors.MessageRejectError) {
	f := new(field.ClearingFeeIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFeeIndicator reads a ClearingFeeIndicator from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetClearingFeeIndicator(f *field.ClearingFeeIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) HandlInst() (*field.HandlInst, errors.MessageRejectError) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetHandlInst(f *field.HandlInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ExecInst() (*field.ExecInst, errors.MessageRejectError) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetExecInst(f *field.ExecInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MinQty() (*field.MinQty, errors.MessageRejectError) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMinQty(f *field.MinQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MaxFloor() (*field.MaxFloor, errors.MessageRejectError) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMaxFloor(f *field.MaxFloor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ExDestination() (*field.ExDestination, errors.MessageRejectError) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetExDestination(f *field.ExDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoTradingSessions() (*field.NoTradingSessions, errors.MessageRejectError) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoTradingSessions(f *field.NoTradingSessions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ProcessCode() (*field.ProcessCode, errors.MessageRejectError) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetProcessCode(f *field.ProcessCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityGroup() (*field.SecurityGroup, errors.MessageRejectError) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecurityGroup(f *field.SecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) UnitOfMeasureQty() (*field.UnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetUnitOfMeasureQty(f *field.UnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityXMLLen() (*field.SecurityXMLLen, errors.MessageRejectError) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecurityXMLLen(f *field.SecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityXML() (*field.SecurityXML, errors.MessageRejectError) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecurityXML(f *field.SecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityXMLSchema() (*field.SecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecurityXMLSchema(f *field.SecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ProductComplex() (*field.ProductComplex, errors.MessageRejectError) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetProductComplex(f *field.ProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SettlMethod() (*field.SettlMethod, errors.MessageRejectError) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSettlMethod(f *field.SettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ExerciseStyle() (*field.ExerciseStyle, errors.MessageRejectError) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetExerciseStyle(f *field.ExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayAmount is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OptPayAmount() (*field.OptPayAmount, errors.MessageRejectError) {
	f := new(field.OptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayAmount reads a OptPayAmount from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOptPayAmount(f *field.OptPayAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PriceQuoteMethod() (*field.PriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPriceQuoteMethod(f *field.PriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ListMethod() (*field.ListMethod, errors.MessageRejectError) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetListMethod(f *field.ListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CapPrice() (*field.CapPrice, errors.MessageRejectError) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCapPrice(f *field.CapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) FloorPrice() (*field.FloorPrice, errors.MessageRejectError) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetFloorPrice(f *field.FloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) FlexibleIndicator() (*field.FlexibleIndicator, errors.MessageRejectError) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetFlexibleIndicator(f *field.FlexibleIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FuturesValuationMethod is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) FuturesValuationMethod() (*field.FuturesValuationMethod, errors.MessageRejectError) {
	f := new(field.FuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetFuturesValuationMethod reads a FuturesValuationMethod from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetFuturesValuationMethod(f *field.FuturesValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PrevClosePx() (*field.PrevClosePx, errors.MessageRejectError) {
	f := new(field.PrevClosePx)
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPrevClosePx(f *field.PrevClosePx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) LocateReqd() (*field.LocateReqd, errors.MessageRejectError) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetLocateReqd(f *field.LocateReqd) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) QtyType() (*field.QtyType, errors.MessageRejectError) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetQtyType(f *field.QtyType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CashOrderQty() (*field.CashOrderQty, errors.MessageRejectError) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCashOrderQty(f *field.CashOrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrderPercent() (*field.OrderPercent, errors.MessageRejectError) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOrderPercent(f *field.OrderPercent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RoundingDirection() (*field.RoundingDirection, errors.MessageRejectError) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRoundingDirection(f *field.RoundingDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RoundingModulus() (*field.RoundingModulus, errors.MessageRejectError) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRoundingModulus(f *field.RoundingModulus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrdType() (*field.OrdType, errors.MessageRejectError) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOrdType(f *field.OrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) StopPx() (*field.StopPx, errors.MessageRejectError) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetStopPx(f *field.StopPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ComplianceID() (*field.ComplianceID, errors.MessageRejectError) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetComplianceID(f *field.ComplianceID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SolicitedFlag is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SolicitedFlag() (*field.SolicitedFlag, errors.MessageRejectError) {
	f := new(field.SolicitedFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSolicitedFlag reads a SolicitedFlag from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSolicitedFlag(f *field.SolicitedFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) IOIID() (*field.IOIID, errors.MessageRejectError) {
	f := new(field.IOIID)
	err := m.Body.Get(f)
	return f, err
}

//GetIOIID reads a IOIID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetIOIID(f *field.IOIID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) QuoteID() (*field.QuoteID, errors.MessageRejectError) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetQuoteID(f *field.QuoteID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TimeInForce() (*field.TimeInForce, errors.MessageRejectError) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTimeInForce(f *field.TimeInForce) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EffectiveTime() (*field.EffectiveTime, errors.MessageRejectError) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetEffectiveTime(f *field.EffectiveTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ExpireDate() (*field.ExpireDate, errors.MessageRejectError) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetExpireDate(f *field.ExpireDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ExpireTime() (*field.ExpireTime, errors.MessageRejectError) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetExpireTime(f *field.ExpireTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GTBookingInst() (*field.GTBookingInst, errors.MessageRejectError) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetGTBookingInst(f *field.GTBookingInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Commission() (*field.Commission, errors.MessageRejectError) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCommission(f *field.Commission) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CommType() (*field.CommType, errors.MessageRejectError) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCommType(f *field.CommType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommCurrency is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CommCurrency() (*field.CommCurrency, errors.MessageRejectError) {
	f := new(field.CommCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetCommCurrency reads a CommCurrency from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCommCurrency(f *field.CommCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FundRenewWaiv is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) FundRenewWaiv() (*field.FundRenewWaiv, errors.MessageRejectError) {
	f := new(field.FundRenewWaiv)
	err := m.Body.Get(f)
	return f, err
}

//GetFundRenewWaiv reads a FundRenewWaiv from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetFundRenewWaiv(f *field.FundRenewWaiv) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrderCapacity() (*field.OrderCapacity, errors.MessageRejectError) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOrderCapacity(f *field.OrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrderRestrictions() (*field.OrderRestrictions, errors.MessageRejectError) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOrderRestrictions(f *field.OrderRestrictions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CustOrderCapacity() (*field.CustOrderCapacity, errors.MessageRejectError) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCustOrderCapacity(f *field.CustOrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ForexReq() (*field.ForexReq, errors.MessageRejectError) {
	f := new(field.ForexReq)
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetForexReq(f *field.ForexReq) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SettlCurrency() (*field.SettlCurrency, errors.MessageRejectError) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSettlCurrency(f *field.SettlCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) BookingType() (*field.BookingType, errors.MessageRejectError) {
	f := new(field.BookingType)
	err := m.Body.Get(f)
	return f, err
}

//GetBookingType reads a BookingType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetBookingType(f *field.BookingType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PositionEffect() (*field.PositionEffect, errors.MessageRejectError) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPositionEffect(f *field.PositionEffect) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CoveredOrUncovered is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CoveredOrUncovered() (*field.CoveredOrUncovered, errors.MessageRejectError) {
	f := new(field.CoveredOrUncovered)
	err := m.Body.Get(f)
	return f, err
}

//GetCoveredOrUncovered reads a CoveredOrUncovered from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCoveredOrUncovered(f *field.CoveredOrUncovered) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MaxShow() (*field.MaxShow, errors.MessageRejectError) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMaxShow(f *field.MaxShow) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetValue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegOffsetValue() (*field.PegOffsetValue, errors.MessageRejectError) {
	f := new(field.PegOffsetValue)
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetValue reads a PegOffsetValue from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegOffsetValue(f *field.PegOffsetValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegMoveType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegMoveType() (*field.PegMoveType, errors.MessageRejectError) {
	f := new(field.PegMoveType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegMoveType reads a PegMoveType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegMoveType(f *field.PegMoveType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegOffsetType() (*field.PegOffsetType, errors.MessageRejectError) {
	f := new(field.PegOffsetType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetType reads a PegOffsetType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegOffsetType(f *field.PegOffsetType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegLimitType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegLimitType() (*field.PegLimitType, errors.MessageRejectError) {
	f := new(field.PegLimitType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegLimitType reads a PegLimitType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegLimitType(f *field.PegLimitType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegRoundDirection is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegRoundDirection() (*field.PegRoundDirection, errors.MessageRejectError) {
	f := new(field.PegRoundDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetPegRoundDirection reads a PegRoundDirection from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegRoundDirection(f *field.PegRoundDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegScope is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegScope() (*field.PegScope, errors.MessageRejectError) {
	f := new(field.PegScope)
	err := m.Body.Get(f)
	return f, err
}

//GetPegScope reads a PegScope from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegScope(f *field.PegScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegPriceType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegPriceType() (*field.PegPriceType, errors.MessageRejectError) {
	f := new(field.PegPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegPriceType reads a PegPriceType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegPriceType(f *field.PegPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityIDSource is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegSecurityIDSource() (*field.PegSecurityIDSource, errors.MessageRejectError) {
	f := new(field.PegSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityIDSource reads a PegSecurityIDSource from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegSecurityIDSource(f *field.PegSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegSecurityID() (*field.PegSecurityID, errors.MessageRejectError) {
	f := new(field.PegSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityID reads a PegSecurityID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegSecurityID(f *field.PegSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSymbol is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegSymbol() (*field.PegSymbol, errors.MessageRejectError) {
	f := new(field.PegSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSymbol reads a PegSymbol from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegSymbol(f *field.PegSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityDesc is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegSecurityDesc() (*field.PegSecurityDesc, errors.MessageRejectError) {
	f := new(field.PegSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityDesc reads a PegSecurityDesc from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegSecurityDesc(f *field.PegSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionInst() (*field.DiscretionInst, errors.MessageRejectError) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDiscretionInst(f *field.DiscretionInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetValue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionOffsetValue() (*field.DiscretionOffsetValue, errors.MessageRejectError) {
	f := new(field.DiscretionOffsetValue)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetValue reads a DiscretionOffsetValue from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDiscretionOffsetValue(f *field.DiscretionOffsetValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionMoveType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionMoveType() (*field.DiscretionMoveType, errors.MessageRejectError) {
	f := new(field.DiscretionMoveType)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionMoveType reads a DiscretionMoveType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDiscretionMoveType(f *field.DiscretionMoveType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionOffsetType() (*field.DiscretionOffsetType, errors.MessageRejectError) {
	f := new(field.DiscretionOffsetType)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetType reads a DiscretionOffsetType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDiscretionOffsetType(f *field.DiscretionOffsetType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionLimitType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionLimitType() (*field.DiscretionLimitType, errors.MessageRejectError) {
	f := new(field.DiscretionLimitType)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionLimitType reads a DiscretionLimitType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDiscretionLimitType(f *field.DiscretionLimitType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionRoundDirection is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionRoundDirection() (*field.DiscretionRoundDirection, errors.MessageRejectError) {
	f := new(field.DiscretionRoundDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionRoundDirection reads a DiscretionRoundDirection from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDiscretionRoundDirection(f *field.DiscretionRoundDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionScope is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionScope() (*field.DiscretionScope, errors.MessageRejectError) {
	f := new(field.DiscretionScope)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionScope reads a DiscretionScope from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDiscretionScope(f *field.DiscretionScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategy is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TargetStrategy() (*field.TargetStrategy, errors.MessageRejectError) {
	f := new(field.TargetStrategy)
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategy reads a TargetStrategy from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTargetStrategy(f *field.TargetStrategy) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategyParameters is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TargetStrategyParameters() (*field.TargetStrategyParameters, errors.MessageRejectError) {
	f := new(field.TargetStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategyParameters reads a TargetStrategyParameters from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTargetStrategyParameters(f *field.TargetStrategyParameters) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParticipationRate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ParticipationRate() (*field.ParticipationRate, errors.MessageRejectError) {
	f := new(field.ParticipationRate)
	err := m.Body.Get(f)
	return f, err
}

//GetParticipationRate reads a ParticipationRate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetParticipationRate(f *field.ParticipationRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CancellationRights() (*field.CancellationRights, errors.MessageRejectError) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCancellationRights(f *field.CancellationRights) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, errors.MessageRejectError) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RegistID() (*field.RegistID, errors.MessageRejectError) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRegistID(f *field.RegistID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Designation() (*field.Designation, errors.MessageRejectError) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDesignation(f *field.Designation) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegRptTypeReq is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MultiLegRptTypeReq() (*field.MultiLegRptTypeReq, errors.MessageRejectError) {
	f := new(field.MultiLegRptTypeReq)
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegRptTypeReq reads a MultiLegRptTypeReq from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMultiLegRptTypeReq(f *field.MultiLegRptTypeReq) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStrategyParameters is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoStrategyParameters() (*field.NoStrategyParameters, errors.MessageRejectError) {
	f := new(field.NoStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStrategyParameters reads a NoStrategyParameters from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoStrategyParameters(f *field.NoStrategyParameters) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchIncrement is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MatchIncrement() (*field.MatchIncrement, errors.MessageRejectError) {
	f := new(field.MatchIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMatchIncrement reads a MatchIncrement from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMatchIncrement(f *field.MatchIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxPriceLevels is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MaxPriceLevels() (*field.MaxPriceLevels, errors.MessageRejectError) {
	f := new(field.MaxPriceLevels)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxPriceLevels reads a MaxPriceLevels from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMaxPriceLevels(f *field.MaxPriceLevels) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryDisplayQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecondaryDisplayQty() (*field.SecondaryDisplayQty, errors.MessageRejectError) {
	f := new(field.SecondaryDisplayQty)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryDisplayQty reads a SecondaryDisplayQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecondaryDisplayQty(f *field.SecondaryDisplayQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayWhen is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DisplayWhen() (*field.DisplayWhen, errors.MessageRejectError) {
	f := new(field.DisplayWhen)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayWhen reads a DisplayWhen from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDisplayWhen(f *field.DisplayWhen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMethod is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DisplayMethod() (*field.DisplayMethod, errors.MessageRejectError) {
	f := new(field.DisplayMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMethod reads a DisplayMethod from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDisplayMethod(f *field.DisplayMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayLowQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DisplayLowQty() (*field.DisplayLowQty, errors.MessageRejectError) {
	f := new(field.DisplayLowQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayLowQty reads a DisplayLowQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDisplayLowQty(f *field.DisplayLowQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayHighQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DisplayHighQty() (*field.DisplayHighQty, errors.MessageRejectError) {
	f := new(field.DisplayHighQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayHighQty reads a DisplayHighQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDisplayHighQty(f *field.DisplayHighQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMinIncr is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DisplayMinIncr() (*field.DisplayMinIncr, errors.MessageRejectError) {
	f := new(field.DisplayMinIncr)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMinIncr reads a DisplayMinIncr from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDisplayMinIncr(f *field.DisplayMinIncr) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefreshQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RefreshQty() (*field.RefreshQty, errors.MessageRejectError) {
	f := new(field.RefreshQty)
	err := m.Body.Get(f)
	return f, err
}

//GetRefreshQty reads a RefreshQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRefreshQty(f *field.RefreshQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DisplayQty() (*field.DisplayQty, errors.MessageRejectError) {
	f := new(field.DisplayQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayQty reads a DisplayQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDisplayQty(f *field.DisplayQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceProtectionScope is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PriceProtectionScope() (*field.PriceProtectionScope, errors.MessageRejectError) {
	f := new(field.PriceProtectionScope)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceProtectionScope reads a PriceProtectionScope from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPriceProtectionScope(f *field.PriceProtectionScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerType() (*field.TriggerType, errors.MessageRejectError) {
	f := new(field.TriggerType)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerType reads a TriggerType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerType(f *field.TriggerType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerAction is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerAction() (*field.TriggerAction, errors.MessageRejectError) {
	f := new(field.TriggerAction)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerAction reads a TriggerAction from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerAction(f *field.TriggerAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPrice is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerPrice() (*field.TriggerPrice, errors.MessageRejectError) {
	f := new(field.TriggerPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPrice reads a TriggerPrice from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerPrice(f *field.TriggerPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSymbol is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerSymbol() (*field.TriggerSymbol, errors.MessageRejectError) {
	f := new(field.TriggerSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSymbol reads a TriggerSymbol from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerSymbol(f *field.TriggerSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerSecurityID() (*field.TriggerSecurityID, errors.MessageRejectError) {
	f := new(field.TriggerSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityID reads a TriggerSecurityID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerSecurityID(f *field.TriggerSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityIDSource is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerSecurityIDSource() (*field.TriggerSecurityIDSource, errors.MessageRejectError) {
	f := new(field.TriggerSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityIDSource reads a TriggerSecurityIDSource from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerSecurityIDSource(f *field.TriggerSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityDesc is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerSecurityDesc() (*field.TriggerSecurityDesc, errors.MessageRejectError) {
	f := new(field.TriggerSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityDesc reads a TriggerSecurityDesc from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerSecurityDesc(f *field.TriggerSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerPriceType() (*field.TriggerPriceType, errors.MessageRejectError) {
	f := new(field.TriggerPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceType reads a TriggerPriceType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerPriceType(f *field.TriggerPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceTypeScope is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerPriceTypeScope() (*field.TriggerPriceTypeScope, errors.MessageRejectError) {
	f := new(field.TriggerPriceTypeScope)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceTypeScope reads a TriggerPriceTypeScope from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerPriceTypeScope(f *field.TriggerPriceTypeScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceDirection is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerPriceDirection() (*field.TriggerPriceDirection, errors.MessageRejectError) {
	f := new(field.TriggerPriceDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceDirection reads a TriggerPriceDirection from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerPriceDirection(f *field.TriggerPriceDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewPrice is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerNewPrice() (*field.TriggerNewPrice, errors.MessageRejectError) {
	f := new(field.TriggerNewPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewPrice reads a TriggerNewPrice from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerNewPrice(f *field.TriggerNewPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerOrderType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerOrderType() (*field.TriggerOrderType, errors.MessageRejectError) {
	f := new(field.TriggerOrderType)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerOrderType reads a TriggerOrderType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerOrderType(f *field.TriggerOrderType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerNewQty() (*field.TriggerNewQty, errors.MessageRejectError) {
	f := new(field.TriggerNewQty)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewQty reads a TriggerNewQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerNewQty(f *field.TriggerNewQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerTradingSessionID() (*field.TriggerTradingSessionID, errors.MessageRejectError) {
	f := new(field.TriggerTradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionID reads a TriggerTradingSessionID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerTradingSessionID(f *field.TriggerTradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionSubID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerTradingSessionSubID() (*field.TriggerTradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TriggerTradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionSubID reads a TriggerTradingSessionSubID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerTradingSessionSubID(f *field.TriggerTradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreTradeAnonymity is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PreTradeAnonymity() (*field.PreTradeAnonymity, errors.MessageRejectError) {
	f := new(field.PreTradeAnonymity)
	err := m.Body.Get(f)
	return f, err
}

//GetPreTradeAnonymity reads a PreTradeAnonymity from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPreTradeAnonymity(f *field.PreTradeAnonymity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestinationIDSource is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ExDestinationIDSource() (*field.ExDestinationIDSource, errors.MessageRejectError) {
	f := new(field.ExDestinationIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestinationIDSource reads a ExDestinationIDSource from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetExDestinationIDSource(f *field.ExDestinationIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SwapPoints is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SwapPoints() (*field.SwapPoints, errors.MessageRejectError) {
	f := new(field.SwapPoints)
	err := m.Body.Get(f)
	return f, err
}

//GetSwapPoints reads a SwapPoints from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSwapPoints(f *field.SwapPoints) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultilegModel is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MultilegModel() (*field.MultilegModel, errors.MessageRejectError) {
	f := new(field.MultilegModel)
	err := m.Body.Get(f)
	return f, err
}

//GetMultilegModel reads a MultilegModel from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMultilegModel(f *field.MultilegModel) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultilegPriceMethod is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MultilegPriceMethod() (*field.MultilegPriceMethod, errors.MessageRejectError) {
	f := new(field.MultilegPriceMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetMultilegPriceMethod reads a MultilegPriceMethod from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMultilegPriceMethod(f *field.MultilegPriceMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RiskFreeRate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RiskFreeRate() (*field.RiskFreeRate, errors.MessageRejectError) {
	f := new(field.RiskFreeRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRiskFreeRate reads a RiskFreeRate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRiskFreeRate(f *field.RiskFreeRate) errors.MessageRejectError {
	return m.Body.Get(f)
}
