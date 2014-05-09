package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	origclordid *field.OrigClOrdIDField,
	clordid *field.ClOrdIDField,
	handlinst *field.HandlInstField,
	side *field.SideField,
	transacttime *field.TransactTimeField,
	ordtype *field.OrdTypeField) OrderCancelReplaceRequestBuilder {
	var builder OrderCancelReplaceRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX43))
	builder.Header.Set(field.NewMsgType("G"))
	builder.Body.Set(origclordid)
	builder.Body.Set(clordid)
	builder.Body.Set(handlinst)
	builder.Body.Set(side)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//OrderID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TradeOriginationDate() (*field.TradeOriginationDateField, errors.MessageRejectError) {
	f := &field.TradeOriginationDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTradeOriginationDate(f *field.TradeOriginationDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrigClOrdID() (*field.OrigClOrdIDField, errors.MessageRejectError) {
	f := &field.OrigClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrigClOrdID(f *field.OrigClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecondaryClOrdID() (*field.SecondaryClOrdIDField, errors.MessageRejectError) {
	f := &field.SecondaryClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecondaryClOrdID(f *field.SecondaryClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ClOrdLinkID() (*field.ClOrdLinkIDField, errors.MessageRejectError) {
	f := &field.ClOrdLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetClOrdLinkID(f *field.ClOrdLinkIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigOrdModTime is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrigOrdModTime() (*field.OrigOrdModTimeField, errors.MessageRejectError) {
	f := &field.OrigOrdModTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigOrdModTime reads a OrigOrdModTime from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrigOrdModTime(f *field.OrigOrdModTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayBookingInst is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DayBookingInst() (*field.DayBookingInstField, errors.MessageRejectError) {
	f := &field.DayBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayBookingInst reads a DayBookingInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDayBookingInst(f *field.DayBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingUnit is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BookingUnit() (*field.BookingUnitField, errors.MessageRejectError) {
	f := &field.BookingUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingUnit reads a BookingUnit from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBookingUnit(f *field.BookingUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreallocMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PreallocMethod() (*field.PreallocMethodField, errors.MessageRejectError) {
	f := &field.PreallocMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreallocMethod reads a PreallocMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPreallocMethod(f *field.PreallocMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoAllocs() (*field.NoAllocsField, errors.MessageRejectError) {
	f := &field.NoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoAllocs(f *field.NoAllocsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettlmntTyp() (*field.SettlmntTypField, errors.MessageRejectError) {
	f := &field.SettlmntTypField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettlmntTyp(f *field.SettlmntTypField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FutSettDate() (*field.FutSettDateField, errors.MessageRejectError) {
	f := &field.FutSettDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFutSettDate(f *field.FutSettDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashMargin is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CashMargin() (*field.CashMarginField, errors.MessageRejectError) {
	f := &field.CashMarginField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashMargin reads a CashMargin from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCashMargin(f *field.CashMarginField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFeeIndicator is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ClearingFeeIndicator() (*field.ClearingFeeIndicatorField, errors.MessageRejectError) {
	f := &field.ClearingFeeIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFeeIndicator reads a ClearingFeeIndicator from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetClearingFeeIndicator(f *field.ClearingFeeIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) HandlInst() (*field.HandlInstField, errors.MessageRejectError) {
	f := &field.HandlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetHandlInst(f *field.HandlInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExecInst() (*field.ExecInstField, errors.MessageRejectError) {
	f := &field.ExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExecInst(f *field.ExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MinQty() (*field.MinQtyField, errors.MessageRejectError) {
	f := &field.MinQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMinQty(f *field.MinQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaxFloor() (*field.MaxFloorField, errors.MessageRejectError) {
	f := &field.MaxFloorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaxFloor(f *field.MaxFloorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExDestination() (*field.ExDestinationField, errors.MessageRejectError) {
	f := &field.ExDestinationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExDestination(f *field.ExDestinationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoTradingSessions() (*field.NoTradingSessionsField, errors.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoTradingSessions(f *field.NoTradingSessionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuantityType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) QuantityType() (*field.QuantityTypeField, errors.MessageRejectError) {
	f := &field.QuantityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuantityType reads a QuantityType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetQuantityType(f *field.QuantityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderQty() (*field.OrderQtyField, errors.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderQty(f *field.OrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CashOrderQty() (*field.CashOrderQtyField, errors.MessageRejectError) {
	f := &field.CashOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCashOrderQty(f *field.CashOrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderPercent() (*field.OrderPercentField, errors.MessageRejectError) {
	f := &field.OrderPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderPercent(f *field.OrderPercentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RoundingDirection() (*field.RoundingDirectionField, errors.MessageRejectError) {
	f := &field.RoundingDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRoundingDirection(f *field.RoundingDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RoundingModulus() (*field.RoundingModulusField, errors.MessageRejectError) {
	f := &field.RoundingModulusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRoundingModulus(f *field.RoundingModulusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrdType() (*field.OrdTypeField, errors.MessageRejectError) {
	f := &field.OrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrdType(f *field.OrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StopPx() (*field.StopPxField, errors.MessageRejectError) {
	f := &field.StopPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStopPx(f *field.StopPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Spread() (*field.SpreadField, errors.MessageRejectError) {
	f := &field.SpreadField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSpread(f *field.SpreadField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrencyField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BenchmarkCurveName() (*field.BenchmarkCurveNameField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBenchmarkCurveName(f *field.BenchmarkCurveNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BenchmarkCurvePoint() (*field.BenchmarkCurvePointField, errors.MessageRejectError) {
	f := &field.BenchmarkCurvePointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) YieldType() (*field.YieldTypeField, errors.MessageRejectError) {
	f := &field.YieldTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetYieldType(f *field.YieldTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Yield() (*field.YieldField, errors.MessageRejectError) {
	f := &field.YieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetYield(f *field.YieldField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegDifference is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegDifference() (*field.PegDifferenceField, errors.MessageRejectError) {
	f := &field.PegDifferenceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegDifference reads a PegDifference from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegDifference(f *field.PegDifferenceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionInst() (*field.DiscretionInstField, errors.MessageRejectError) {
	f := &field.DiscretionInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionInst(f *field.DiscretionInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffset is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionOffset() (*field.DiscretionOffsetField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffset reads a DiscretionOffset from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionOffset(f *field.DiscretionOffsetField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ComplianceID() (*field.ComplianceIDField, errors.MessageRejectError) {
	f := &field.ComplianceIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetComplianceID(f *field.ComplianceIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SolicitedFlag is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SolicitedFlag() (*field.SolicitedFlagField, errors.MessageRejectError) {
	f := &field.SolicitedFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSolicitedFlag reads a SolicitedFlag from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSolicitedFlag(f *field.SolicitedFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TimeInForce() (*field.TimeInForceField, errors.MessageRejectError) {
	f := &field.TimeInForceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTimeInForce(f *field.TimeInForceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EffectiveTime() (*field.EffectiveTimeField, errors.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEffectiveTime(f *field.EffectiveTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExpireDate() (*field.ExpireDateField, errors.MessageRejectError) {
	f := &field.ExpireDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExpireDate(f *field.ExpireDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExpireTime() (*field.ExpireTimeField, errors.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExpireTime(f *field.ExpireTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GTBookingInst() (*field.GTBookingInstField, errors.MessageRejectError) {
	f := &field.GTBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetGTBookingInst(f *field.GTBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Commission() (*field.CommissionField, errors.MessageRejectError) {
	f := &field.CommissionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCommission(f *field.CommissionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CommType() (*field.CommTypeField, errors.MessageRejectError) {
	f := &field.CommTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCommType(f *field.CommTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommCurrency is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CommCurrency() (*field.CommCurrencyField, errors.MessageRejectError) {
	f := &field.CommCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommCurrency reads a CommCurrency from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCommCurrency(f *field.CommCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FundRenewWaiv is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FundRenewWaiv() (*field.FundRenewWaivField, errors.MessageRejectError) {
	f := &field.FundRenewWaivField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFundRenewWaiv reads a FundRenewWaiv from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFundRenewWaiv(f *field.FundRenewWaivField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderCapacity() (*field.OrderCapacityField, errors.MessageRejectError) {
	f := &field.OrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderCapacity(f *field.OrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderRestrictions() (*field.OrderRestrictionsField, errors.MessageRejectError) {
	f := &field.OrderRestrictionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderRestrictions(f *field.OrderRestrictionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CustOrderCapacity() (*field.CustOrderCapacityField, errors.MessageRejectError) {
	f := &field.CustOrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCustOrderCapacity(f *field.CustOrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Rule80A is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Rule80A() (*field.Rule80AField, errors.MessageRejectError) {
	f := &field.Rule80AField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRule80A reads a Rule80A from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRule80A(f *field.Rule80AField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ForexReq() (*field.ForexReqField, errors.MessageRejectError) {
	f := &field.ForexReqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetForexReq(f *field.ForexReqField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettlCurrency() (*field.SettlCurrencyField, errors.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettlCurrency(f *field.SettlCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate2 is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FutSettDate2() (*field.FutSettDate2Field, errors.MessageRejectError) {
	f := &field.FutSettDate2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate2 reads a FutSettDate2 from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFutSettDate2(f *field.FutSettDate2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty2 is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderQty2() (*field.OrderQty2Field, errors.MessageRejectError) {
	f := &field.OrderQty2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty2 reads a OrderQty2 from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderQty2(f *field.OrderQty2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price2 is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Price2() (*field.Price2Field, errors.MessageRejectError) {
	f := &field.Price2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice2 reads a Price2 from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPrice2(f *field.Price2Field) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PositionEffect() (*field.PositionEffectField, errors.MessageRejectError) {
	f := &field.PositionEffectField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPositionEffect(f *field.PositionEffectField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CoveredOrUncovered is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CoveredOrUncovered() (*field.CoveredOrUncoveredField, errors.MessageRejectError) {
	f := &field.CoveredOrUncoveredField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCoveredOrUncovered reads a CoveredOrUncovered from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCoveredOrUncovered(f *field.CoveredOrUncoveredField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaxShow() (*field.MaxShowField, errors.MessageRejectError) {
	f := &field.MaxShowField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaxShow(f *field.MaxShowField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) LocateReqd() (*field.LocateReqdField, errors.MessageRejectError) {
	f := &field.LocateReqdField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetLocateReqd(f *field.LocateReqdField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CancellationRights() (*field.CancellationRightsField, errors.MessageRejectError) {
	f := &field.CancellationRightsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCancellationRights(f *field.CancellationRightsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MoneyLaunderingStatus() (*field.MoneyLaunderingStatusField, errors.MessageRejectError) {
	f := &field.MoneyLaunderingStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RegistID() (*field.RegistIDField, errors.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRegistID(f *field.RegistIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Designation() (*field.DesignationField, errors.MessageRejectError) {
	f := &field.DesignationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDesignation(f *field.DesignationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestRate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AccruedInterestRate() (*field.AccruedInterestRateField, errors.MessageRejectError) {
	f := &field.AccruedInterestRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestRate reads a AccruedInterestRate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAccruedInterestRate(f *field.AccruedInterestRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestAmt is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AccruedInterestAmt() (*field.AccruedInterestAmtField, errors.MessageRejectError) {
	f := &field.AccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestAmt reads a AccruedInterestAmt from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAccruedInterestAmt(f *field.AccruedInterestAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NetMoney() (*field.NetMoneyField, errors.MessageRejectError) {
	f := &field.NetMoneyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNetMoney(f *field.NetMoneyField) errors.MessageRejectError {
	return m.Body.Get(f)
}
