package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//MultilegOrderCancelReplaceRequest msg type = AC.
type MultilegOrderCancelReplaceRequest struct {
	message.Message
}

//MultilegOrderCancelReplaceRequestBuilder builds MultilegOrderCancelReplaceRequest messages.
type MultilegOrderCancelReplaceRequestBuilder struct {
	message.MessageBuilder
}

//CreateMultilegOrderCancelReplaceRequestBuilder returns an initialized MultilegOrderCancelReplaceRequestBuilder with specified required fields.
func CreateMultilegOrderCancelReplaceRequestBuilder(
	origclordid field.OrigClOrdID,
	clordid field.ClOrdID,
	handlinst field.HandlInst,
	side field.Side,
	nolegs field.NoLegs,
	transacttime field.TransactTime,
	ordtype field.OrdType) MultilegOrderCancelReplaceRequestBuilder {
	var builder MultilegOrderCancelReplaceRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(origclordid)
	builder.Body.Set(clordid)
	builder.Body.Set(handlinst)
	builder.Body.Set(side)
	builder.Body.Set(nolegs)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//OrderID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrigClOrdID() (*field.OrigClOrdID, errors.MessageRejectError) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOrigClOrdID(f *field.OrigClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecondaryClOrdID() (*field.SecondaryClOrdID, errors.MessageRejectError) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSecondaryClOrdID(f *field.SecondaryClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ClOrdLinkID() (*field.ClOrdLinkID, errors.MessageRejectError) {
	f := new(field.ClOrdLinkID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetClOrdLinkID(f *field.ClOrdLinkID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigOrdModTime is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrigOrdModTime() (*field.OrigOrdModTime, errors.MessageRejectError) {
	f := new(field.OrigOrdModTime)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigOrdModTime reads a OrigOrdModTime from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOrigOrdModTime(f *field.OrigOrdModTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayBookingInst is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) DayBookingInst() (*field.DayBookingInst, errors.MessageRejectError) {
	f := new(field.DayBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//GetDayBookingInst reads a DayBookingInst from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetDayBookingInst(f *field.DayBookingInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingUnit is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) BookingUnit() (*field.BookingUnit, errors.MessageRejectError) {
	f := new(field.BookingUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetBookingUnit reads a BookingUnit from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetBookingUnit(f *field.BookingUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreallocMethod is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) PreallocMethod() (*field.PreallocMethod, errors.MessageRejectError) {
	f := new(field.PreallocMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPreallocMethod reads a PreallocMethod from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetPreallocMethod(f *field.PreallocMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NoAllocs() (*field.NoAllocs, errors.MessageRejectError) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetNoAllocs(f *field.NoAllocs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SettlmntTyp() (*field.SettlmntTyp, errors.MessageRejectError) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSettlmntTyp(f *field.SettlmntTyp) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) FutSettDate() (*field.FutSettDate, errors.MessageRejectError) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetFutSettDate(f *field.FutSettDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashMargin is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CashMargin() (*field.CashMargin, errors.MessageRejectError) {
	f := new(field.CashMargin)
	err := m.Body.Get(f)
	return f, err
}

//GetCashMargin reads a CashMargin from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCashMargin(f *field.CashMargin) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFeeIndicator is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ClearingFeeIndicator() (*field.ClearingFeeIndicator, errors.MessageRejectError) {
	f := new(field.ClearingFeeIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFeeIndicator reads a ClearingFeeIndicator from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetClearingFeeIndicator(f *field.ClearingFeeIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) HandlInst() (*field.HandlInst, errors.MessageRejectError) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetHandlInst(f *field.HandlInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ExecInst() (*field.ExecInst, errors.MessageRejectError) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetExecInst(f *field.ExecInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MinQty() (*field.MinQty, errors.MessageRejectError) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetMinQty(f *field.MinQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MaxFloor() (*field.MaxFloor, errors.MessageRejectError) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetMaxFloor(f *field.MaxFloor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ExDestination() (*field.ExDestination, errors.MessageRejectError) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetExDestination(f *field.ExDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NoTradingSessions() (*field.NoTradingSessions, errors.MessageRejectError) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetNoTradingSessions(f *field.NoTradingSessions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ProcessCode() (*field.ProcessCode, errors.MessageRejectError) {
	f := new(field.ProcessCode)
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetProcessCode(f *field.ProcessCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) PrevClosePx() (*field.PrevClosePx, errors.MessageRejectError) {
	f := new(field.PrevClosePx)
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetPrevClosePx(f *field.PrevClosePx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) LocateReqd() (*field.LocateReqd, errors.MessageRejectError) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetLocateReqd(f *field.LocateReqd) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuantityType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) QuantityType() (*field.QuantityType, errors.MessageRejectError) {
	f := new(field.QuantityType)
	err := m.Body.Get(f)
	return f, err
}

//GetQuantityType reads a QuantityType from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetQuantityType(f *field.QuantityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CashOrderQty() (*field.CashOrderQty, errors.MessageRejectError) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCashOrderQty(f *field.CashOrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrderPercent() (*field.OrderPercent, errors.MessageRejectError) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOrderPercent(f *field.OrderPercent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RoundingDirection() (*field.RoundingDirection, errors.MessageRejectError) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetRoundingDirection(f *field.RoundingDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RoundingModulus() (*field.RoundingModulus, errors.MessageRejectError) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetRoundingModulus(f *field.RoundingModulus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrdType() (*field.OrdType, errors.MessageRejectError) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOrdType(f *field.OrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) StopPx() (*field.StopPx, errors.MessageRejectError) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetStopPx(f *field.StopPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ComplianceID() (*field.ComplianceID, errors.MessageRejectError) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetComplianceID(f *field.ComplianceID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SolicitedFlag is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SolicitedFlag() (*field.SolicitedFlag, errors.MessageRejectError) {
	f := new(field.SolicitedFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSolicitedFlag reads a SolicitedFlag from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSolicitedFlag(f *field.SolicitedFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIid is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) IOIid() (*field.IOIid, errors.MessageRejectError) {
	f := new(field.IOIid)
	err := m.Body.Get(f)
	return f, err
}

//GetIOIid reads a IOIid from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetIOIid(f *field.IOIid) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) QuoteID() (*field.QuoteID, errors.MessageRejectError) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetQuoteID(f *field.QuoteID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) TimeInForce() (*field.TimeInForce, errors.MessageRejectError) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetTimeInForce(f *field.TimeInForce) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EffectiveTime() (*field.EffectiveTime, errors.MessageRejectError) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetEffectiveTime(f *field.EffectiveTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ExpireDate() (*field.ExpireDate, errors.MessageRejectError) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetExpireDate(f *field.ExpireDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ExpireTime() (*field.ExpireTime, errors.MessageRejectError) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetExpireTime(f *field.ExpireTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GTBookingInst() (*field.GTBookingInst, errors.MessageRejectError) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetGTBookingInst(f *field.GTBookingInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Commission() (*field.Commission, errors.MessageRejectError) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCommission(f *field.Commission) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CommType() (*field.CommType, errors.MessageRejectError) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCommType(f *field.CommType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommCurrency is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CommCurrency() (*field.CommCurrency, errors.MessageRejectError) {
	f := new(field.CommCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetCommCurrency reads a CommCurrency from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCommCurrency(f *field.CommCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FundRenewWaiv is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) FundRenewWaiv() (*field.FundRenewWaiv, errors.MessageRejectError) {
	f := new(field.FundRenewWaiv)
	err := m.Body.Get(f)
	return f, err
}

//GetFundRenewWaiv reads a FundRenewWaiv from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetFundRenewWaiv(f *field.FundRenewWaiv) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrderCapacity() (*field.OrderCapacity, errors.MessageRejectError) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOrderCapacity(f *field.OrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrderRestrictions() (*field.OrderRestrictions, errors.MessageRejectError) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOrderRestrictions(f *field.OrderRestrictions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CustOrderCapacity() (*field.CustOrderCapacity, errors.MessageRejectError) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCustOrderCapacity(f *field.CustOrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ForexReq() (*field.ForexReq, errors.MessageRejectError) {
	f := new(field.ForexReq)
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetForexReq(f *field.ForexReq) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SettlCurrency() (*field.SettlCurrency, errors.MessageRejectError) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSettlCurrency(f *field.SettlCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) PositionEffect() (*field.PositionEffect, errors.MessageRejectError) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetPositionEffect(f *field.PositionEffect) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CoveredOrUncovered is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CoveredOrUncovered() (*field.CoveredOrUncovered, errors.MessageRejectError) {
	f := new(field.CoveredOrUncovered)
	err := m.Body.Get(f)
	return f, err
}

//GetCoveredOrUncovered reads a CoveredOrUncovered from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCoveredOrUncovered(f *field.CoveredOrUncovered) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MaxShow() (*field.MaxShow, errors.MessageRejectError) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetMaxShow(f *field.MaxShow) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegDifference is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) PegDifference() (*field.PegDifference, errors.MessageRejectError) {
	f := new(field.PegDifference)
	err := m.Body.Get(f)
	return f, err
}

//GetPegDifference reads a PegDifference from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetPegDifference(f *field.PegDifference) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) DiscretionInst() (*field.DiscretionInst, errors.MessageRejectError) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetDiscretionInst(f *field.DiscretionInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffset is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) DiscretionOffset() (*field.DiscretionOffset, errors.MessageRejectError) {
	f := new(field.DiscretionOffset)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffset reads a DiscretionOffset from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetDiscretionOffset(f *field.DiscretionOffset) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CancellationRights() (*field.CancellationRights, errors.MessageRejectError) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCancellationRights(f *field.CancellationRights) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, errors.MessageRejectError) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RegistID() (*field.RegistID, errors.MessageRejectError) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetRegistID(f *field.RegistID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Designation() (*field.Designation, errors.MessageRejectError) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetDesignation(f *field.Designation) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegRptTypeReq is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MultiLegRptTypeReq() (*field.MultiLegRptTypeReq, errors.MessageRejectError) {
	f := new(field.MultiLegRptTypeReq)
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegRptTypeReq reads a MultiLegRptTypeReq from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetMultiLegRptTypeReq(f *field.MultiLegRptTypeReq) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NetMoney() (*field.NetMoney, errors.MessageRejectError) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetNetMoney(f *field.NetMoney) errors.MessageRejectError {
	return m.Body.Get(f)
}
