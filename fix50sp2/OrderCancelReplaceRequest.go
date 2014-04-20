package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//OrderCancelReplaceRequest msg type = G.
type OrderCancelReplaceRequest struct {
	message.Message
}

//OrderCancelReplaceRequestBuilder builds OrderCancelReplaceRequest messages.
type OrderCancelReplaceRequestBuilder struct {
	message.MessageBuilder
}

//CreateOrderCancelReplaceRequestBuilder returns an initialized OrderCancelReplaceRequestBuilder with specified required fields.
func CreateOrderCancelReplaceRequestBuilder(
	clordid field.ClOrdID,
	side field.Side,
	transacttime field.TransactTime,
	ordtype field.OrdType) OrderCancelReplaceRequestBuilder {
	var builder OrderCancelReplaceRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(clordid)
	builder.Body.Set(side)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//OrderID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderID() (*field.OrderID, errors.MessageRejectError) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderID(f *field.OrderID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TradeOriginationDate() (*field.TradeOriginationDate, errors.MessageRejectError) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTradeOriginationDate(f *field.TradeOriginationDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrigClOrdID() (*field.OrigClOrdID, errors.MessageRejectError) {
	f := new(field.OrigClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrigClOrdID(f *field.OrigClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ClOrdID() (*field.ClOrdID, errors.MessageRejectError) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetClOrdID(f *field.ClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecondaryClOrdID() (*field.SecondaryClOrdID, errors.MessageRejectError) {
	f := new(field.SecondaryClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecondaryClOrdID(f *field.SecondaryClOrdID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ClOrdLinkID() (*field.ClOrdLinkID, errors.MessageRejectError) {
	f := new(field.ClOrdLinkID)
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetClOrdLinkID(f *field.ClOrdLinkID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ListID() (*field.ListID, errors.MessageRejectError) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetListID(f *field.ListID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigOrdModTime is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrigOrdModTime() (*field.OrigOrdModTime, errors.MessageRejectError) {
	f := new(field.OrigOrdModTime)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigOrdModTime reads a OrigOrdModTime from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrigOrdModTime(f *field.OrigOrdModTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Account() (*field.Account, errors.MessageRejectError) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAccount(f *field.Account) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AcctIDSource() (*field.AcctIDSource, errors.MessageRejectError) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAcctIDSource(f *field.AcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AccountType() (*field.AccountType, errors.MessageRejectError) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAccountType(f *field.AccountType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayBookingInst is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DayBookingInst() (*field.DayBookingInst, errors.MessageRejectError) {
	f := new(field.DayBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//GetDayBookingInst reads a DayBookingInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDayBookingInst(f *field.DayBookingInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingUnit is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BookingUnit() (*field.BookingUnit, errors.MessageRejectError) {
	f := new(field.BookingUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetBookingUnit reads a BookingUnit from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBookingUnit(f *field.BookingUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreallocMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PreallocMethod() (*field.PreallocMethod, errors.MessageRejectError) {
	f := new(field.PreallocMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPreallocMethod reads a PreallocMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPreallocMethod(f *field.PreallocMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AllocID() (*field.AllocID, errors.MessageRejectError) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAllocID(f *field.AllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoAllocs() (*field.NoAllocs, errors.MessageRejectError) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoAllocs(f *field.NoAllocs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettlType() (*field.SettlType, errors.MessageRejectError) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettlType(f *field.SettlType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettlDate() (*field.SettlDate, errors.MessageRejectError) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettlDate(f *field.SettlDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashMargin is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CashMargin() (*field.CashMargin, errors.MessageRejectError) {
	f := new(field.CashMargin)
	err := m.Body.Get(f)
	return f, err
}

//GetCashMargin reads a CashMargin from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCashMargin(f *field.CashMargin) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFeeIndicator is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ClearingFeeIndicator() (*field.ClearingFeeIndicator, errors.MessageRejectError) {
	f := new(field.ClearingFeeIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFeeIndicator reads a ClearingFeeIndicator from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetClearingFeeIndicator(f *field.ClearingFeeIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) HandlInst() (*field.HandlInst, errors.MessageRejectError) {
	f := new(field.HandlInst)
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetHandlInst(f *field.HandlInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExecInst() (*field.ExecInst, errors.MessageRejectError) {
	f := new(field.ExecInst)
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExecInst(f *field.ExecInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MinQty() (*field.MinQty, errors.MessageRejectError) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMinQty(f *field.MinQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaxFloor() (*field.MaxFloor, errors.MessageRejectError) {
	f := new(field.MaxFloor)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaxFloor(f *field.MaxFloor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExDestination() (*field.ExDestination, errors.MessageRejectError) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExDestination(f *field.ExDestination) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoTradingSessions() (*field.NoTradingSessions, errors.MessageRejectError) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoTradingSessions(f *field.NoTradingSessions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityGroup() (*field.SecurityGroup, errors.MessageRejectError) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityGroup(f *field.SecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) UnitOfMeasureQty() (*field.UnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetUnitOfMeasureQty(f *field.UnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityXMLLen() (*field.SecurityXMLLen, errors.MessageRejectError) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityXMLLen(f *field.SecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityXML() (*field.SecurityXML, errors.MessageRejectError) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityXML(f *field.SecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityXMLSchema() (*field.SecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityXMLSchema(f *field.SecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ProductComplex() (*field.ProductComplex, errors.MessageRejectError) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetProductComplex(f *field.ProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettlMethod() (*field.SettlMethod, errors.MessageRejectError) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettlMethod(f *field.SettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExerciseStyle() (*field.ExerciseStyle, errors.MessageRejectError) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExerciseStyle(f *field.ExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OptPayoutAmount() (*field.OptPayoutAmount, errors.MessageRejectError) {
	f := new(field.OptPayoutAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOptPayoutAmount(f *field.OptPayoutAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PriceQuoteMethod() (*field.PriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPriceQuoteMethod(f *field.PriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ListMethod() (*field.ListMethod, errors.MessageRejectError) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetListMethod(f *field.ListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CapPrice() (*field.CapPrice, errors.MessageRejectError) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCapPrice(f *field.CapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FloorPrice() (*field.FloorPrice, errors.MessageRejectError) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFloorPrice(f *field.FloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FlexibleIndicator() (*field.FlexibleIndicator, errors.MessageRejectError) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFlexibleIndicator(f *field.FlexibleIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ValuationMethod() (*field.ValuationMethod, errors.MessageRejectError) {
	f := new(field.ValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetValuationMethod(f *field.ValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ContractMultiplierUnit() (*field.ContractMultiplierUnit, errors.MessageRejectError) {
	f := new(field.ContractMultiplierUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetContractMultiplierUnit(f *field.ContractMultiplierUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FlowScheduleType() (*field.FlowScheduleType, errors.MessageRejectError) {
	f := new(field.FlowScheduleType)
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFlowScheduleType(f *field.FlowScheduleType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RestructuringType() (*field.RestructuringType, errors.MessageRejectError) {
	f := new(field.RestructuringType)
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRestructuringType(f *field.RestructuringType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Seniority() (*field.Seniority, errors.MessageRejectError) {
	f := new(field.Seniority)
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSeniority(f *field.Seniority) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.NotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstanding, errors.MessageRejectError) {
	f := new(field.OriginalNotionalPercentageOutstanding)
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstanding) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AttachmentPoint() (*field.AttachmentPoint, errors.MessageRejectError) {
	f := new(field.AttachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAttachmentPoint(f *field.AttachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DetachmentPoint() (*field.DetachmentPoint, errors.MessageRejectError) {
	f := new(field.DetachmentPoint)
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDetachmentPoint(f *field.DetachmentPoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethod, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecision, errors.MessageRejectError) {
	f := new(field.StrikePriceBoundaryPrecision)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecision) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethod, errors.MessageRejectError) {
	f := new(field.UnderlyingPriceDeterminationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OptPayoutType() (*field.OptPayoutType, errors.MessageRejectError) {
	f := new(field.OptPayoutType)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOptPayoutType(f *field.OptPayoutType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoComplexEvents() (*field.NoComplexEvents, errors.MessageRejectError) {
	f := new(field.NoComplexEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoComplexEvents(f *field.NoComplexEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AgreementDesc() (*field.AgreementDesc, errors.MessageRejectError) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAgreementDesc(f *field.AgreementDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AgreementID() (*field.AgreementID, errors.MessageRejectError) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAgreementID(f *field.AgreementID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AgreementDate() (*field.AgreementDate, errors.MessageRejectError) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAgreementDate(f *field.AgreementDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AgreementCurrency() (*field.AgreementCurrency, errors.MessageRejectError) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAgreementCurrency(f *field.AgreementCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TerminationType() (*field.TerminationType, errors.MessageRejectError) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTerminationType(f *field.TerminationType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StartDate() (*field.StartDate, errors.MessageRejectError) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStartDate(f *field.StartDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EndDate() (*field.EndDate, errors.MessageRejectError) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEndDate(f *field.EndDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DeliveryType() (*field.DeliveryType, errors.MessageRejectError) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDeliveryType(f *field.DeliveryType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MarginRatio() (*field.MarginRatio, errors.MessageRejectError) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMarginRatio(f *field.MarginRatio) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) QtyType() (*field.QtyType, errors.MessageRejectError) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetQtyType(f *field.QtyType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderQty() (*field.OrderQty, errors.MessageRejectError) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderQty(f *field.OrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CashOrderQty() (*field.CashOrderQty, errors.MessageRejectError) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCashOrderQty(f *field.CashOrderQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderPercent() (*field.OrderPercent, errors.MessageRejectError) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderPercent(f *field.OrderPercent) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RoundingDirection() (*field.RoundingDirection, errors.MessageRejectError) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRoundingDirection(f *field.RoundingDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RoundingModulus() (*field.RoundingModulus, errors.MessageRejectError) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRoundingModulus(f *field.RoundingModulus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrdType() (*field.OrdType, errors.MessageRejectError) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrdType(f *field.OrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Price() (*field.Price, errors.MessageRejectError) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPrice(f *field.Price) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StopPx() (*field.StopPx, errors.MessageRejectError) {
	f := new(field.StopPx)
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStopPx(f *field.StopPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Spread() (*field.Spread, errors.MessageRejectError) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSpread(f *field.Spread) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BenchmarkCurveName() (*field.BenchmarkCurveName, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBenchmarkCurveName(f *field.BenchmarkCurveName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, errors.MessageRejectError) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BenchmarkPrice() (*field.BenchmarkPrice, errors.MessageRejectError) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBenchmarkPrice(f *field.BenchmarkPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BenchmarkPriceType() (*field.BenchmarkPriceType, errors.MessageRejectError) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBenchmarkPriceType(f *field.BenchmarkPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BenchmarkSecurityID() (*field.BenchmarkSecurityID, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBenchmarkSecurityID(f *field.BenchmarkSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) YieldType() (*field.YieldType, errors.MessageRejectError) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetYieldType(f *field.YieldType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Yield() (*field.Yield, errors.MessageRejectError) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetYield(f *field.Yield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) YieldCalcDate() (*field.YieldCalcDate, errors.MessageRejectError) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetYieldCalcDate(f *field.YieldCalcDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) YieldRedemptionDate() (*field.YieldRedemptionDate, errors.MessageRejectError) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetYieldRedemptionDate(f *field.YieldRedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) YieldRedemptionPrice() (*field.YieldRedemptionPrice, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetYieldRedemptionPrice(f *field.YieldRedemptionPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetValue is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegOffsetValue() (*field.PegOffsetValue, errors.MessageRejectError) {
	f := new(field.PegOffsetValue)
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetValue reads a PegOffsetValue from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegOffsetValue(f *field.PegOffsetValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegMoveType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegMoveType() (*field.PegMoveType, errors.MessageRejectError) {
	f := new(field.PegMoveType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegMoveType reads a PegMoveType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegMoveType(f *field.PegMoveType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegOffsetType() (*field.PegOffsetType, errors.MessageRejectError) {
	f := new(field.PegOffsetType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetType reads a PegOffsetType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegOffsetType(f *field.PegOffsetType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegLimitType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegLimitType() (*field.PegLimitType, errors.MessageRejectError) {
	f := new(field.PegLimitType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegLimitType reads a PegLimitType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegLimitType(f *field.PegLimitType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegRoundDirection is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegRoundDirection() (*field.PegRoundDirection, errors.MessageRejectError) {
	f := new(field.PegRoundDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetPegRoundDirection reads a PegRoundDirection from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegRoundDirection(f *field.PegRoundDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegScope is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegScope() (*field.PegScope, errors.MessageRejectError) {
	f := new(field.PegScope)
	err := m.Body.Get(f)
	return f, err
}

//GetPegScope reads a PegScope from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegScope(f *field.PegScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegPriceType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegPriceType() (*field.PegPriceType, errors.MessageRejectError) {
	f := new(field.PegPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPegPriceType reads a PegPriceType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegPriceType(f *field.PegPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityIDSource is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegSecurityIDSource() (*field.PegSecurityIDSource, errors.MessageRejectError) {
	f := new(field.PegSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityIDSource reads a PegSecurityIDSource from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegSecurityIDSource(f *field.PegSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegSecurityID() (*field.PegSecurityID, errors.MessageRejectError) {
	f := new(field.PegSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityID reads a PegSecurityID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegSecurityID(f *field.PegSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSymbol is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegSymbol() (*field.PegSymbol, errors.MessageRejectError) {
	f := new(field.PegSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSymbol reads a PegSymbol from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegSymbol(f *field.PegSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityDesc is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegSecurityDesc() (*field.PegSecurityDesc, errors.MessageRejectError) {
	f := new(field.PegSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityDesc reads a PegSecurityDesc from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegSecurityDesc(f *field.PegSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionInst() (*field.DiscretionInst, errors.MessageRejectError) {
	f := new(field.DiscretionInst)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionInst(f *field.DiscretionInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetValue is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionOffsetValue() (*field.DiscretionOffsetValue, errors.MessageRejectError) {
	f := new(field.DiscretionOffsetValue)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetValue reads a DiscretionOffsetValue from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionOffsetValue(f *field.DiscretionOffsetValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionMoveType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionMoveType() (*field.DiscretionMoveType, errors.MessageRejectError) {
	f := new(field.DiscretionMoveType)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionMoveType reads a DiscretionMoveType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionMoveType(f *field.DiscretionMoveType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionOffsetType() (*field.DiscretionOffsetType, errors.MessageRejectError) {
	f := new(field.DiscretionOffsetType)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetType reads a DiscretionOffsetType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionOffsetType(f *field.DiscretionOffsetType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionLimitType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionLimitType() (*field.DiscretionLimitType, errors.MessageRejectError) {
	f := new(field.DiscretionLimitType)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionLimitType reads a DiscretionLimitType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionLimitType(f *field.DiscretionLimitType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionRoundDirection is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionRoundDirection() (*field.DiscretionRoundDirection, errors.MessageRejectError) {
	f := new(field.DiscretionRoundDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionRoundDirection reads a DiscretionRoundDirection from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionRoundDirection(f *field.DiscretionRoundDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionScope is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionScope() (*field.DiscretionScope, errors.MessageRejectError) {
	f := new(field.DiscretionScope)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionScope reads a DiscretionScope from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionScope(f *field.DiscretionScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategy is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TargetStrategy() (*field.TargetStrategy, errors.MessageRejectError) {
	f := new(field.TargetStrategy)
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategy reads a TargetStrategy from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTargetStrategy(f *field.TargetStrategy) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategyParameters is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TargetStrategyParameters() (*field.TargetStrategyParameters, errors.MessageRejectError) {
	f := new(field.TargetStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategyParameters reads a TargetStrategyParameters from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTargetStrategyParameters(f *field.TargetStrategyParameters) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParticipationRate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ParticipationRate() (*field.ParticipationRate, errors.MessageRejectError) {
	f := new(field.ParticipationRate)
	err := m.Body.Get(f)
	return f, err
}

//GetParticipationRate reads a ParticipationRate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetParticipationRate(f *field.ParticipationRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ComplianceID() (*field.ComplianceID, errors.MessageRejectError) {
	f := new(field.ComplianceID)
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetComplianceID(f *field.ComplianceID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SolicitedFlag is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SolicitedFlag() (*field.SolicitedFlag, errors.MessageRejectError) {
	f := new(field.SolicitedFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSolicitedFlag reads a SolicitedFlag from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSolicitedFlag(f *field.SolicitedFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TimeInForce() (*field.TimeInForce, errors.MessageRejectError) {
	f := new(field.TimeInForce)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTimeInForce(f *field.TimeInForce) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EffectiveTime() (*field.EffectiveTime, errors.MessageRejectError) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEffectiveTime(f *field.EffectiveTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExpireDate() (*field.ExpireDate, errors.MessageRejectError) {
	f := new(field.ExpireDate)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExpireDate(f *field.ExpireDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExpireTime() (*field.ExpireTime, errors.MessageRejectError) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExpireTime(f *field.ExpireTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GTBookingInst() (*field.GTBookingInst, errors.MessageRejectError) {
	f := new(field.GTBookingInst)
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetGTBookingInst(f *field.GTBookingInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Commission() (*field.Commission, errors.MessageRejectError) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCommission(f *field.Commission) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CommType() (*field.CommType, errors.MessageRejectError) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCommType(f *field.CommType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommCurrency is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CommCurrency() (*field.CommCurrency, errors.MessageRejectError) {
	f := new(field.CommCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetCommCurrency reads a CommCurrency from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCommCurrency(f *field.CommCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FundRenewWaiv is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FundRenewWaiv() (*field.FundRenewWaiv, errors.MessageRejectError) {
	f := new(field.FundRenewWaiv)
	err := m.Body.Get(f)
	return f, err
}

//GetFundRenewWaiv reads a FundRenewWaiv from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFundRenewWaiv(f *field.FundRenewWaiv) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderCapacity() (*field.OrderCapacity, errors.MessageRejectError) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderCapacity(f *field.OrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderRestrictions() (*field.OrderRestrictions, errors.MessageRejectError) {
	f := new(field.OrderRestrictions)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderRestrictions(f *field.OrderRestrictions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CustOrderCapacity() (*field.CustOrderCapacity, errors.MessageRejectError) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCustOrderCapacity(f *field.CustOrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ForexReq() (*field.ForexReq, errors.MessageRejectError) {
	f := new(field.ForexReq)
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetForexReq(f *field.ForexReq) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettlCurrency() (*field.SettlCurrency, errors.MessageRejectError) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettlCurrency(f *field.SettlCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BookingType() (*field.BookingType, errors.MessageRejectError) {
	f := new(field.BookingType)
	err := m.Body.Get(f)
	return f, err
}

//GetBookingType reads a BookingType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBookingType(f *field.BookingType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate2 is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettlDate2() (*field.SettlDate2, errors.MessageRejectError) {
	f := new(field.SettlDate2)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate2 reads a SettlDate2 from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettlDate2(f *field.SettlDate2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty2 is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderQty2() (*field.OrderQty2, errors.MessageRejectError) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty2 reads a OrderQty2 from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderQty2(f *field.OrderQty2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price2 is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Price2() (*field.Price2, errors.MessageRejectError) {
	f := new(field.Price2)
	err := m.Body.Get(f)
	return f, err
}

//GetPrice2 reads a Price2 from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPrice2(f *field.Price2) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PositionEffect() (*field.PositionEffect, errors.MessageRejectError) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPositionEffect(f *field.PositionEffect) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CoveredOrUncovered is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CoveredOrUncovered() (*field.CoveredOrUncovered, errors.MessageRejectError) {
	f := new(field.CoveredOrUncovered)
	err := m.Body.Get(f)
	return f, err
}

//GetCoveredOrUncovered reads a CoveredOrUncovered from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCoveredOrUncovered(f *field.CoveredOrUncovered) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaxShow() (*field.MaxShow, errors.MessageRejectError) {
	f := new(field.MaxShow)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaxShow(f *field.MaxShow) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) LocateReqd() (*field.LocateReqd, errors.MessageRejectError) {
	f := new(field.LocateReqd)
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetLocateReqd(f *field.LocateReqd) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CancellationRights() (*field.CancellationRights, errors.MessageRejectError) {
	f := new(field.CancellationRights)
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCancellationRights(f *field.CancellationRights) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MoneyLaunderingStatus() (*field.MoneyLaunderingStatus, errors.MessageRejectError) {
	f := new(field.MoneyLaunderingStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RegistID() (*field.RegistID, errors.MessageRejectError) {
	f := new(field.RegistID)
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRegistID(f *field.RegistID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Designation() (*field.Designation, errors.MessageRejectError) {
	f := new(field.Designation)
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDesignation(f *field.Designation) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStrategyParameters is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoStrategyParameters() (*field.NoStrategyParameters, errors.MessageRejectError) {
	f := new(field.NoStrategyParameters)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStrategyParameters reads a NoStrategyParameters from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoStrategyParameters(f *field.NoStrategyParameters) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ManualOrderIndicator is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ManualOrderIndicator() (*field.ManualOrderIndicator, errors.MessageRejectError) {
	f := new(field.ManualOrderIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetManualOrderIndicator reads a ManualOrderIndicator from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetManualOrderIndicator(f *field.ManualOrderIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustDirectedOrder is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CustDirectedOrder() (*field.CustDirectedOrder, errors.MessageRejectError) {
	f := new(field.CustDirectedOrder)
	err := m.Body.Get(f)
	return f, err
}

//GetCustDirectedOrder reads a CustDirectedOrder from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCustDirectedOrder(f *field.CustDirectedOrder) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ReceivedDeptID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ReceivedDeptID() (*field.ReceivedDeptID, errors.MessageRejectError) {
	f := new(field.ReceivedDeptID)
	err := m.Body.Get(f)
	return f, err
}

//GetReceivedDeptID reads a ReceivedDeptID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetReceivedDeptID(f *field.ReceivedDeptID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderHandlingInst is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CustOrderHandlingInst() (*field.CustOrderHandlingInst, errors.MessageRejectError) {
	f := new(field.CustOrderHandlingInst)
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderHandlingInst reads a CustOrderHandlingInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCustOrderHandlingInst(f *field.CustOrderHandlingInst) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderHandlingInstSource is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderHandlingInstSource() (*field.OrderHandlingInstSource, errors.MessageRejectError) {
	f := new(field.OrderHandlingInstSource)
	err := m.Body.Get(f)
	return f, err
}

//GetOrderHandlingInstSource reads a OrderHandlingInstSource from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderHandlingInstSource(f *field.OrderHandlingInstSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrdRegTimestamps is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoTrdRegTimestamps() (*field.NoTrdRegTimestamps, errors.MessageRejectError) {
	f := new(field.NoTrdRegTimestamps)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrdRegTimestamps reads a NoTrdRegTimestamps from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoTrdRegTimestamps(f *field.NoTrdRegTimestamps) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchIncrement is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MatchIncrement() (*field.MatchIncrement, errors.MessageRejectError) {
	f := new(field.MatchIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMatchIncrement reads a MatchIncrement from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMatchIncrement(f *field.MatchIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxPriceLevels is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaxPriceLevels() (*field.MaxPriceLevels, errors.MessageRejectError) {
	f := new(field.MaxPriceLevels)
	err := m.Body.Get(f)
	return f, err
}

//GetMaxPriceLevels reads a MaxPriceLevels from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaxPriceLevels(f *field.MaxPriceLevels) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryDisplayQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecondaryDisplayQty() (*field.SecondaryDisplayQty, errors.MessageRejectError) {
	f := new(field.SecondaryDisplayQty)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryDisplayQty reads a SecondaryDisplayQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecondaryDisplayQty(f *field.SecondaryDisplayQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayWhen is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DisplayWhen() (*field.DisplayWhen, errors.MessageRejectError) {
	f := new(field.DisplayWhen)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayWhen reads a DisplayWhen from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDisplayWhen(f *field.DisplayWhen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DisplayMethod() (*field.DisplayMethod, errors.MessageRejectError) {
	f := new(field.DisplayMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMethod reads a DisplayMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDisplayMethod(f *field.DisplayMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayLowQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DisplayLowQty() (*field.DisplayLowQty, errors.MessageRejectError) {
	f := new(field.DisplayLowQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayLowQty reads a DisplayLowQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDisplayLowQty(f *field.DisplayLowQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayHighQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DisplayHighQty() (*field.DisplayHighQty, errors.MessageRejectError) {
	f := new(field.DisplayHighQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayHighQty reads a DisplayHighQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDisplayHighQty(f *field.DisplayHighQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMinIncr is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DisplayMinIncr() (*field.DisplayMinIncr, errors.MessageRejectError) {
	f := new(field.DisplayMinIncr)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMinIncr reads a DisplayMinIncr from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDisplayMinIncr(f *field.DisplayMinIncr) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefreshQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RefreshQty() (*field.RefreshQty, errors.MessageRejectError) {
	f := new(field.RefreshQty)
	err := m.Body.Get(f)
	return f, err
}

//GetRefreshQty reads a RefreshQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRefreshQty(f *field.RefreshQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DisplayQty() (*field.DisplayQty, errors.MessageRejectError) {
	f := new(field.DisplayQty)
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayQty reads a DisplayQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDisplayQty(f *field.DisplayQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceProtectionScope is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PriceProtectionScope() (*field.PriceProtectionScope, errors.MessageRejectError) {
	f := new(field.PriceProtectionScope)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceProtectionScope reads a PriceProtectionScope from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPriceProtectionScope(f *field.PriceProtectionScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerType() (*field.TriggerType, errors.MessageRejectError) {
	f := new(field.TriggerType)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerType reads a TriggerType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerType(f *field.TriggerType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerAction is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerAction() (*field.TriggerAction, errors.MessageRejectError) {
	f := new(field.TriggerAction)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerAction reads a TriggerAction from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerAction(f *field.TriggerAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPrice is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerPrice() (*field.TriggerPrice, errors.MessageRejectError) {
	f := new(field.TriggerPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPrice reads a TriggerPrice from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerPrice(f *field.TriggerPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSymbol is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerSymbol() (*field.TriggerSymbol, errors.MessageRejectError) {
	f := new(field.TriggerSymbol)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSymbol reads a TriggerSymbol from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerSymbol(f *field.TriggerSymbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerSecurityID() (*field.TriggerSecurityID, errors.MessageRejectError) {
	f := new(field.TriggerSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityID reads a TriggerSecurityID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerSecurityID(f *field.TriggerSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityIDSource is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerSecurityIDSource() (*field.TriggerSecurityIDSource, errors.MessageRejectError) {
	f := new(field.TriggerSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityIDSource reads a TriggerSecurityIDSource from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerSecurityIDSource(f *field.TriggerSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityDesc is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerSecurityDesc() (*field.TriggerSecurityDesc, errors.MessageRejectError) {
	f := new(field.TriggerSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityDesc reads a TriggerSecurityDesc from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerSecurityDesc(f *field.TriggerSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerPriceType() (*field.TriggerPriceType, errors.MessageRejectError) {
	f := new(field.TriggerPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceType reads a TriggerPriceType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerPriceType(f *field.TriggerPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceTypeScope is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerPriceTypeScope() (*field.TriggerPriceTypeScope, errors.MessageRejectError) {
	f := new(field.TriggerPriceTypeScope)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceTypeScope reads a TriggerPriceTypeScope from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerPriceTypeScope(f *field.TriggerPriceTypeScope) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceDirection is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerPriceDirection() (*field.TriggerPriceDirection, errors.MessageRejectError) {
	f := new(field.TriggerPriceDirection)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceDirection reads a TriggerPriceDirection from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerPriceDirection(f *field.TriggerPriceDirection) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewPrice is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerNewPrice() (*field.TriggerNewPrice, errors.MessageRejectError) {
	f := new(field.TriggerNewPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewPrice reads a TriggerNewPrice from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerNewPrice(f *field.TriggerNewPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerOrderType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerOrderType() (*field.TriggerOrderType, errors.MessageRejectError) {
	f := new(field.TriggerOrderType)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerOrderType reads a TriggerOrderType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerOrderType(f *field.TriggerOrderType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerNewQty() (*field.TriggerNewQty, errors.MessageRejectError) {
	f := new(field.TriggerNewQty)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewQty reads a TriggerNewQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerNewQty(f *field.TriggerNewQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerTradingSessionID() (*field.TriggerTradingSessionID, errors.MessageRejectError) {
	f := new(field.TriggerTradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionID reads a TriggerTradingSessionID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerTradingSessionID(f *field.TriggerTradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionSubID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerTradingSessionSubID() (*field.TriggerTradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TriggerTradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionSubID reads a TriggerTradingSessionSubID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerTradingSessionSubID(f *field.TriggerTradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreTradeAnonymity is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PreTradeAnonymity() (*field.PreTradeAnonymity, errors.MessageRejectError) {
	f := new(field.PreTradeAnonymity)
	err := m.Body.Get(f)
	return f, err
}

//GetPreTradeAnonymity reads a PreTradeAnonymity from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPreTradeAnonymity(f *field.PreTradeAnonymity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestinationIDSource is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExDestinationIDSource() (*field.ExDestinationIDSource, errors.MessageRejectError) {
	f := new(field.ExDestinationIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetExDestinationIDSource reads a ExDestinationIDSource from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExDestinationIDSource(f *field.ExDestinationIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}
