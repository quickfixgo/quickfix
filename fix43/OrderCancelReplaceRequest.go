package fix43

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
	origclordid field.OrigClOrdID,
	clordid field.ClOrdID,
	handlinst field.HandlInst,
	side field.Side,
	transacttime field.TransactTime,
	ordtype field.OrdType) OrderCancelReplaceRequestBuilder {
	var builder OrderCancelReplaceRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(origclordid)
	builder.Body.Set(clordid)
	builder.Body.Set(handlinst)
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

//OrigClOrdID is a required field for OrderCancelReplaceRequest.
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

//SettlmntTyp is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettlmntTyp() (*field.SettlmntTyp, errors.MessageRejectError) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettlmntTyp(f *field.SettlmntTyp) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FutSettDate() (*field.FutSettDate, errors.MessageRejectError) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFutSettDate(f *field.FutSettDate) errors.MessageRejectError {
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

//HandlInst is a required field for OrderCancelReplaceRequest.
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

//QuantityType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) QuantityType() (*field.QuantityType, errors.MessageRejectError) {
	f := new(field.QuantityType)
	err := m.Body.Get(f)
	return f, err
}

//GetQuantityType reads a QuantityType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetQuantityType(f *field.QuantityType) errors.MessageRejectError {
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

//PegDifference is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegDifference() (*field.PegDifference, errors.MessageRejectError) {
	f := new(field.PegDifference)
	err := m.Body.Get(f)
	return f, err
}

//GetPegDifference reads a PegDifference from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegDifference(f *field.PegDifference) errors.MessageRejectError {
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

//DiscretionOffset is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionOffset() (*field.DiscretionOffset, errors.MessageRejectError) {
	f := new(field.DiscretionOffset)
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffset reads a DiscretionOffset from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionOffset(f *field.DiscretionOffset) errors.MessageRejectError {
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

//Rule80A is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Rule80A() (*field.Rule80A, errors.MessageRejectError) {
	f := new(field.Rule80A)
	err := m.Body.Get(f)
	return f, err
}

//GetRule80A reads a Rule80A from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRule80A(f *field.Rule80A) errors.MessageRejectError {
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

//FutSettDate2 is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FutSettDate2() (*field.FutSettDate2, errors.MessageRejectError) {
	f := new(field.FutSettDate2)
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate2 reads a FutSettDate2 from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFutSettDate2(f *field.FutSettDate2) errors.MessageRejectError {
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

//AccruedInterestRate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AccruedInterestRate() (*field.AccruedInterestRate, errors.MessageRejectError) {
	f := new(field.AccruedInterestRate)
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestRate reads a AccruedInterestRate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAccruedInterestRate(f *field.AccruedInterestRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestAmt is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AccruedInterestAmt() (*field.AccruedInterestAmt, errors.MessageRejectError) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestAmt reads a AccruedInterestAmt from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAccruedInterestAmt(f *field.AccruedInterestAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NetMoney() (*field.NetMoney, errors.MessageRejectError) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNetMoney(f *field.NetMoney) errors.MessageRejectError {
	return m.Body.Get(f)
}
