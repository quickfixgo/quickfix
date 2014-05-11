package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	clordid *field.ClOrdIDField,
	side *field.SideField,
	transacttime *field.TransactTimeField,
	ordtype *field.OrdTypeField) OrderCancelReplaceRequestBuilder {
	var builder OrderCancelReplaceRequestBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("G"))
	builder.Body().Set(clordid)
	builder.Body().Set(side)
	builder.Body().Set(transacttime)
	builder.Body().Set(ordtype)
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

//TradeDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a non-required field for OrderCancelReplaceRequest.
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

//AcctIDSource is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AcctIDSource() (*field.AcctIDSourceField, errors.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAcctIDSource(f *field.AcctIDSourceField) errors.MessageRejectError {
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

//AllocID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AllocID() (*field.AllocIDField, errors.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAllocID(f *field.AllocIDField) errors.MessageRejectError {
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

//SettlType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettlType() (*field.SettlTypeField, errors.MessageRejectError) {
	f := &field.SettlTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettlType(f *field.SettlTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettlDate() (*field.SettlDateField, errors.MessageRejectError) {
	f := &field.SettlDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettlDate(f *field.SettlDateField) errors.MessageRejectError {
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

//HandlInst is a non-required field for OrderCancelReplaceRequest.
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

//SecuritySubType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
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

//StrikeCurrency is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
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

//Pool is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityStatus() (*field.SecurityStatusField, errors.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityStatus(f *field.SecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StrikeMultiplier() (*field.StrikeMultiplierField, errors.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStrikeMultiplier(f *field.StrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StrikeValue() (*field.StrikeValueField, errors.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStrikeValue(f *field.StrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MinPriceIncrement() (*field.MinPriceIncrementField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMinPriceIncrement(f *field.MinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PositionLimit() (*field.PositionLimitField, errors.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPositionLimit(f *field.PositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NTPositionLimit() (*field.NTPositionLimitField, errors.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNTPositionLimit(f *field.NTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoInstrumentParties() (*field.NoInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) UnitOfMeasure() (*field.UnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetUnitOfMeasure(f *field.UnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TimeUnit() (*field.TimeUnitField, errors.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTimeUnit(f *field.TimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaturityTime() (*field.MaturityTimeField, errors.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaturityTime(f *field.MaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityGroup() (*field.SecurityGroupField, errors.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityGroup(f *field.SecurityGroupField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityXMLLen() (*field.SecurityXMLLenField, errors.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityXMLLen(f *field.SecurityXMLLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityXML() (*field.SecurityXMLField, errors.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityXML(f *field.SecurityXMLField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecurityXMLSchema() (*field.SecurityXMLSchemaField, errors.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ProductComplex() (*field.ProductComplexField, errors.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetProductComplex(f *field.ProductComplexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettlMethod() (*field.SettlMethodField, errors.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettlMethod(f *field.SettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExerciseStyle() (*field.ExerciseStyleField, errors.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExerciseStyle(f *field.ExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OptPayoutAmount() (*field.OptPayoutAmountField, errors.MessageRejectError) {
	f := &field.OptPayoutAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOptPayoutAmount(f *field.OptPayoutAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PriceQuoteMethod() (*field.PriceQuoteMethodField, errors.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ListMethod() (*field.ListMethodField, errors.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetListMethod(f *field.ListMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CapPrice() (*field.CapPriceField, errors.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCapPrice(f *field.CapPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FloorPrice() (*field.FloorPriceField, errors.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFloorPrice(f *field.FloorPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FlexibleIndicator() (*field.FlexibleIndicatorField, errors.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFlexibleIndicator(f *field.FlexibleIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, errors.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ValuationMethod() (*field.ValuationMethodField, errors.MessageRejectError) {
	f := &field.ValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetValuationMethod(f *field.ValuationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ContractMultiplierUnit() (*field.ContractMultiplierUnitField, errors.MessageRejectError) {
	f := &field.ContractMultiplierUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetContractMultiplierUnit(f *field.ContractMultiplierUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) FlowScheduleType() (*field.FlowScheduleTypeField, errors.MessageRejectError) {
	f := &field.FlowScheduleTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetFlowScheduleType(f *field.FlowScheduleTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RestructuringType() (*field.RestructuringTypeField, errors.MessageRejectError) {
	f := &field.RestructuringTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRestructuringType(f *field.RestructuringTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) Seniority() (*field.SeniorityField, errors.MessageRejectError) {
	f := &field.SeniorityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSeniority(f *field.SeniorityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstandingField, errors.MessageRejectError) {
	f := &field.NotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstandingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstandingField, errors.MessageRejectError) {
	f := &field.OriginalNotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstandingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AttachmentPoint() (*field.AttachmentPointField, errors.MessageRejectError) {
	f := &field.AttachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAttachmentPoint(f *field.AttachmentPointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DetachmentPoint() (*field.DetachmentPointField, errors.MessageRejectError) {
	f := &field.DetachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDetachmentPoint(f *field.DetachmentPointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethodField, errors.MessageRejectError) {
	f := &field.StrikePriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethodField, errors.MessageRejectError) {
	f := &field.StrikePriceBoundaryMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecisionField, errors.MessageRejectError) {
	f := &field.StrikePriceBoundaryPrecisionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecisionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethodField, errors.MessageRejectError) {
	f := &field.UnderlyingPriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OptPayoutType() (*field.OptPayoutTypeField, errors.MessageRejectError) {
	f := &field.OptPayoutTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOptPayoutType(f *field.OptPayoutTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoComplexEvents() (*field.NoComplexEventsField, errors.MessageRejectError) {
	f := &field.NoComplexEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoComplexEvents(f *field.NoComplexEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AgreementDesc() (*field.AgreementDescField, errors.MessageRejectError) {
	f := &field.AgreementDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAgreementDesc(f *field.AgreementDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AgreementID() (*field.AgreementIDField, errors.MessageRejectError) {
	f := &field.AgreementIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAgreementID(f *field.AgreementIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AgreementDate() (*field.AgreementDateField, errors.MessageRejectError) {
	f := &field.AgreementDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAgreementDate(f *field.AgreementDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) AgreementCurrency() (*field.AgreementCurrencyField, errors.MessageRejectError) {
	f := &field.AgreementCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetAgreementCurrency(f *field.AgreementCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TerminationType() (*field.TerminationTypeField, errors.MessageRejectError) {
	f := &field.TerminationTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTerminationType(f *field.TerminationTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) StartDate() (*field.StartDateField, errors.MessageRejectError) {
	f := &field.StartDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetStartDate(f *field.StartDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) EndDate() (*field.EndDateField, errors.MessageRejectError) {
	f := &field.EndDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetEndDate(f *field.EndDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DeliveryType() (*field.DeliveryTypeField, errors.MessageRejectError) {
	f := &field.DeliveryTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDeliveryType(f *field.DeliveryTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MarginRatio() (*field.MarginRatioField, errors.MessageRejectError) {
	f := &field.MarginRatioField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMarginRatio(f *field.MarginRatioField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
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

//QtyType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) QtyType() (*field.QtyTypeField, errors.MessageRejectError) {
	f := &field.QtyTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetQtyType(f *field.QtyTypeField) errors.MessageRejectError {
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

//BenchmarkPrice is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BenchmarkPrice() (*field.BenchmarkPriceField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBenchmarkPrice(f *field.BenchmarkPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BenchmarkPriceType() (*field.BenchmarkPriceTypeField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBenchmarkPriceType(f *field.BenchmarkPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BenchmarkSecurityID() (*field.BenchmarkSecurityIDField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBenchmarkSecurityID(f *field.BenchmarkSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSourceField) errors.MessageRejectError {
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

//YieldCalcDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) YieldCalcDate() (*field.YieldCalcDateField, errors.MessageRejectError) {
	f := &field.YieldCalcDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetYieldCalcDate(f *field.YieldCalcDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) YieldRedemptionDate() (*field.YieldRedemptionDateField, errors.MessageRejectError) {
	f := &field.YieldRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetYieldRedemptionDate(f *field.YieldRedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) YieldRedemptionPrice() (*field.YieldRedemptionPriceField, errors.MessageRejectError) {
	f := &field.YieldRedemptionPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetYieldRedemptionPrice(f *field.YieldRedemptionPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) YieldRedemptionPriceType() (*field.YieldRedemptionPriceTypeField, errors.MessageRejectError) {
	f := &field.YieldRedemptionPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetValue is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegOffsetValue() (*field.PegOffsetValueField, errors.MessageRejectError) {
	f := &field.PegOffsetValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetValue reads a PegOffsetValue from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegOffsetValue(f *field.PegOffsetValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegMoveType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegMoveType() (*field.PegMoveTypeField, errors.MessageRejectError) {
	f := &field.PegMoveTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegMoveType reads a PegMoveType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegMoveType(f *field.PegMoveTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegOffsetType() (*field.PegOffsetTypeField, errors.MessageRejectError) {
	f := &field.PegOffsetTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetType reads a PegOffsetType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegOffsetType(f *field.PegOffsetTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegLimitType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegLimitType() (*field.PegLimitTypeField, errors.MessageRejectError) {
	f := &field.PegLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegLimitType reads a PegLimitType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegLimitType(f *field.PegLimitTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegRoundDirection is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegRoundDirection() (*field.PegRoundDirectionField, errors.MessageRejectError) {
	f := &field.PegRoundDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegRoundDirection reads a PegRoundDirection from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegRoundDirection(f *field.PegRoundDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegScope is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegScope() (*field.PegScopeField, errors.MessageRejectError) {
	f := &field.PegScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegScope reads a PegScope from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegScope(f *field.PegScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegPriceType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegPriceType() (*field.PegPriceTypeField, errors.MessageRejectError) {
	f := &field.PegPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegPriceType reads a PegPriceType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegPriceType(f *field.PegPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityIDSource is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegSecurityIDSource() (*field.PegSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.PegSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityIDSource reads a PegSecurityIDSource from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegSecurityIDSource(f *field.PegSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegSecurityID() (*field.PegSecurityIDField, errors.MessageRejectError) {
	f := &field.PegSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityID reads a PegSecurityID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegSecurityID(f *field.PegSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSymbol is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegSymbol() (*field.PegSymbolField, errors.MessageRejectError) {
	f := &field.PegSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSymbol reads a PegSymbol from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegSymbol(f *field.PegSymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityDesc is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PegSecurityDesc() (*field.PegSecurityDescField, errors.MessageRejectError) {
	f := &field.PegSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityDesc reads a PegSecurityDesc from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPegSecurityDesc(f *field.PegSecurityDescField) errors.MessageRejectError {
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

//DiscretionOffsetValue is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionOffsetValue() (*field.DiscretionOffsetValueField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetValue reads a DiscretionOffsetValue from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionOffsetValue(f *field.DiscretionOffsetValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionMoveType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionMoveType() (*field.DiscretionMoveTypeField, errors.MessageRejectError) {
	f := &field.DiscretionMoveTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionMoveType reads a DiscretionMoveType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionMoveType(f *field.DiscretionMoveTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionOffsetType() (*field.DiscretionOffsetTypeField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetType reads a DiscretionOffsetType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionOffsetType(f *field.DiscretionOffsetTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionLimitType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionLimitType() (*field.DiscretionLimitTypeField, errors.MessageRejectError) {
	f := &field.DiscretionLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionLimitType reads a DiscretionLimitType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionLimitType(f *field.DiscretionLimitTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionRoundDirection is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionRoundDirection() (*field.DiscretionRoundDirectionField, errors.MessageRejectError) {
	f := &field.DiscretionRoundDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionRoundDirection reads a DiscretionRoundDirection from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionRoundDirection(f *field.DiscretionRoundDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionScope is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DiscretionScope() (*field.DiscretionScopeField, errors.MessageRejectError) {
	f := &field.DiscretionScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionScope reads a DiscretionScope from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDiscretionScope(f *field.DiscretionScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategy is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TargetStrategy() (*field.TargetStrategyField, errors.MessageRejectError) {
	f := &field.TargetStrategyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategy reads a TargetStrategy from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTargetStrategy(f *field.TargetStrategyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategyParameters is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TargetStrategyParameters() (*field.TargetStrategyParametersField, errors.MessageRejectError) {
	f := &field.TargetStrategyParametersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategyParameters reads a TargetStrategyParameters from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTargetStrategyParameters(f *field.TargetStrategyParametersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParticipationRate is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ParticipationRate() (*field.ParticipationRateField, errors.MessageRejectError) {
	f := &field.ParticipationRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetParticipationRate reads a ParticipationRate from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetParticipationRate(f *field.ParticipationRateField) errors.MessageRejectError {
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

//BookingType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) BookingType() (*field.BookingTypeField, errors.MessageRejectError) {
	f := &field.BookingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingType reads a BookingType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetBookingType(f *field.BookingTypeField) errors.MessageRejectError {
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

//SettlDate2 is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SettlDate2() (*field.SettlDate2Field, errors.MessageRejectError) {
	f := &field.SettlDate2Field{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate2 reads a SettlDate2 from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSettlDate2(f *field.SettlDate2Field) errors.MessageRejectError {
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

//NoStrategyParameters is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoStrategyParameters() (*field.NoStrategyParametersField, errors.MessageRejectError) {
	f := &field.NoStrategyParametersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStrategyParameters reads a NoStrategyParameters from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoStrategyParameters(f *field.NoStrategyParametersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ManualOrderIndicator is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ManualOrderIndicator() (*field.ManualOrderIndicatorField, errors.MessageRejectError) {
	f := &field.ManualOrderIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetManualOrderIndicator reads a ManualOrderIndicator from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetManualOrderIndicator(f *field.ManualOrderIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustDirectedOrder is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CustDirectedOrder() (*field.CustDirectedOrderField, errors.MessageRejectError) {
	f := &field.CustDirectedOrderField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustDirectedOrder reads a CustDirectedOrder from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCustDirectedOrder(f *field.CustDirectedOrderField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ReceivedDeptID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ReceivedDeptID() (*field.ReceivedDeptIDField, errors.MessageRejectError) {
	f := &field.ReceivedDeptIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetReceivedDeptID reads a ReceivedDeptID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetReceivedDeptID(f *field.ReceivedDeptIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderHandlingInst is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) CustOrderHandlingInst() (*field.CustOrderHandlingInstField, errors.MessageRejectError) {
	f := &field.CustOrderHandlingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderHandlingInst reads a CustOrderHandlingInst from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetCustOrderHandlingInst(f *field.CustOrderHandlingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderHandlingInstSource is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) OrderHandlingInstSource() (*field.OrderHandlingInstSourceField, errors.MessageRejectError) {
	f := &field.OrderHandlingInstSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderHandlingInstSource reads a OrderHandlingInstSource from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetOrderHandlingInstSource(f *field.OrderHandlingInstSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrdRegTimestamps is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) NoTrdRegTimestamps() (*field.NoTrdRegTimestampsField, errors.MessageRejectError) {
	f := &field.NoTrdRegTimestampsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrdRegTimestamps reads a NoTrdRegTimestamps from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetNoTrdRegTimestamps(f *field.NoTrdRegTimestampsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchIncrement is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MatchIncrement() (*field.MatchIncrementField, errors.MessageRejectError) {
	f := &field.MatchIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchIncrement reads a MatchIncrement from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMatchIncrement(f *field.MatchIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxPriceLevels is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) MaxPriceLevels() (*field.MaxPriceLevelsField, errors.MessageRejectError) {
	f := &field.MaxPriceLevelsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxPriceLevels reads a MaxPriceLevels from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetMaxPriceLevels(f *field.MaxPriceLevelsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryDisplayQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) SecondaryDisplayQty() (*field.SecondaryDisplayQtyField, errors.MessageRejectError) {
	f := &field.SecondaryDisplayQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryDisplayQty reads a SecondaryDisplayQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetSecondaryDisplayQty(f *field.SecondaryDisplayQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayWhen is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DisplayWhen() (*field.DisplayWhenField, errors.MessageRejectError) {
	f := &field.DisplayWhenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayWhen reads a DisplayWhen from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDisplayWhen(f *field.DisplayWhenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMethod is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DisplayMethod() (*field.DisplayMethodField, errors.MessageRejectError) {
	f := &field.DisplayMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMethod reads a DisplayMethod from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDisplayMethod(f *field.DisplayMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayLowQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DisplayLowQty() (*field.DisplayLowQtyField, errors.MessageRejectError) {
	f := &field.DisplayLowQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayLowQty reads a DisplayLowQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDisplayLowQty(f *field.DisplayLowQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayHighQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DisplayHighQty() (*field.DisplayHighQtyField, errors.MessageRejectError) {
	f := &field.DisplayHighQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayHighQty reads a DisplayHighQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDisplayHighQty(f *field.DisplayHighQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMinIncr is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DisplayMinIncr() (*field.DisplayMinIncrField, errors.MessageRejectError) {
	f := &field.DisplayMinIncrField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMinIncr reads a DisplayMinIncr from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDisplayMinIncr(f *field.DisplayMinIncrField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefreshQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) RefreshQty() (*field.RefreshQtyField, errors.MessageRejectError) {
	f := &field.RefreshQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefreshQty reads a RefreshQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetRefreshQty(f *field.RefreshQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) DisplayQty() (*field.DisplayQtyField, errors.MessageRejectError) {
	f := &field.DisplayQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayQty reads a DisplayQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetDisplayQty(f *field.DisplayQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceProtectionScope is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PriceProtectionScope() (*field.PriceProtectionScopeField, errors.MessageRejectError) {
	f := &field.PriceProtectionScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceProtectionScope reads a PriceProtectionScope from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPriceProtectionScope(f *field.PriceProtectionScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerType() (*field.TriggerTypeField, errors.MessageRejectError) {
	f := &field.TriggerTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerType reads a TriggerType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerType(f *field.TriggerTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerAction is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerAction() (*field.TriggerActionField, errors.MessageRejectError) {
	f := &field.TriggerActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerAction reads a TriggerAction from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerAction(f *field.TriggerActionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPrice is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerPrice() (*field.TriggerPriceField, errors.MessageRejectError) {
	f := &field.TriggerPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPrice reads a TriggerPrice from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerPrice(f *field.TriggerPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSymbol is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerSymbol() (*field.TriggerSymbolField, errors.MessageRejectError) {
	f := &field.TriggerSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSymbol reads a TriggerSymbol from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerSymbol(f *field.TriggerSymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerSecurityID() (*field.TriggerSecurityIDField, errors.MessageRejectError) {
	f := &field.TriggerSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityID reads a TriggerSecurityID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerSecurityID(f *field.TriggerSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityIDSource is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerSecurityIDSource() (*field.TriggerSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.TriggerSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityIDSource reads a TriggerSecurityIDSource from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerSecurityIDSource(f *field.TriggerSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityDesc is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerSecurityDesc() (*field.TriggerSecurityDescField, errors.MessageRejectError) {
	f := &field.TriggerSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityDesc reads a TriggerSecurityDesc from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerSecurityDesc(f *field.TriggerSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerPriceType() (*field.TriggerPriceTypeField, errors.MessageRejectError) {
	f := &field.TriggerPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceType reads a TriggerPriceType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerPriceType(f *field.TriggerPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceTypeScope is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerPriceTypeScope() (*field.TriggerPriceTypeScopeField, errors.MessageRejectError) {
	f := &field.TriggerPriceTypeScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceTypeScope reads a TriggerPriceTypeScope from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerPriceTypeScope(f *field.TriggerPriceTypeScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceDirection is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerPriceDirection() (*field.TriggerPriceDirectionField, errors.MessageRejectError) {
	f := &field.TriggerPriceDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceDirection reads a TriggerPriceDirection from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerPriceDirection(f *field.TriggerPriceDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewPrice is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerNewPrice() (*field.TriggerNewPriceField, errors.MessageRejectError) {
	f := &field.TriggerNewPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewPrice reads a TriggerNewPrice from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerNewPrice(f *field.TriggerNewPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerOrderType is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerOrderType() (*field.TriggerOrderTypeField, errors.MessageRejectError) {
	f := &field.TriggerOrderTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerOrderType reads a TriggerOrderType from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerOrderType(f *field.TriggerOrderTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewQty is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerNewQty() (*field.TriggerNewQtyField, errors.MessageRejectError) {
	f := &field.TriggerNewQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewQty reads a TriggerNewQty from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerNewQty(f *field.TriggerNewQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerTradingSessionID() (*field.TriggerTradingSessionIDField, errors.MessageRejectError) {
	f := &field.TriggerTradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionID reads a TriggerTradingSessionID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerTradingSessionID(f *field.TriggerTradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionSubID is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) TriggerTradingSessionSubID() (*field.TriggerTradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TriggerTradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionSubID reads a TriggerTradingSessionSubID from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetTriggerTradingSessionSubID(f *field.TriggerTradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreTradeAnonymity is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) PreTradeAnonymity() (*field.PreTradeAnonymityField, errors.MessageRejectError) {
	f := &field.PreTradeAnonymityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreTradeAnonymity reads a PreTradeAnonymity from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetPreTradeAnonymity(f *field.PreTradeAnonymityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestinationIDSource is a non-required field for OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) ExDestinationIDSource() (*field.ExDestinationIDSourceField, errors.MessageRejectError) {
	f := &field.ExDestinationIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDestinationIDSource reads a ExDestinationIDSource from OrderCancelReplaceRequest.
func (m OrderCancelReplaceRequest) GetExDestinationIDSource(f *field.ExDestinationIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}
