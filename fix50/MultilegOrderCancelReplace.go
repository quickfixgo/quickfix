package fix50

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
	origclordid *field.OrigClOrdIDField,
	clordid *field.ClOrdIDField,
	side *field.SideField,
	nolegs *field.NoLegsField,
	transacttime *field.TransactTimeField,
	ordtype *field.OrdTypeField) MultilegOrderCancelReplaceBuilder {
	var builder MultilegOrderCancelReplaceBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header().Set(field.NewMsgType("AC"))
	builder.Body().Set(origclordid)
	builder.Body().Set(clordid)
	builder.Body().Set(side)
	builder.Body().Set(nolegs)
	builder.Body().Set(transacttime)
	builder.Body().Set(ordtype)
	return builder
}

//OrderID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrigClOrdID() (*field.OrigClOrdIDField, errors.MessageRejectError) {
	f := &field.OrigClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOrigClOrdID(f *field.OrigClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecondaryClOrdID() (*field.SecondaryClOrdIDField, errors.MessageRejectError) {
	f := &field.SecondaryClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecondaryClOrdID(f *field.SecondaryClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ClOrdLinkID() (*field.ClOrdLinkIDField, errors.MessageRejectError) {
	f := &field.ClOrdLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetClOrdLinkID(f *field.ClOrdLinkIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigOrdModTime is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrigOrdModTime() (*field.OrigOrdModTimeField, errors.MessageRejectError) {
	f := &field.OrigOrdModTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigOrdModTime reads a OrigOrdModTime from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOrigOrdModTime(f *field.OrigOrdModTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TradeOriginationDate() (*field.TradeOriginationDateField, errors.MessageRejectError) {
	f := &field.TradeOriginationDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTradeOriginationDate(f *field.TradeOriginationDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AcctIDSource is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) AcctIDSource() (*field.AcctIDSourceField, errors.MessageRejectError) {
	f := &field.AcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAcctIDSource reads a AcctIDSource from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetAcctIDSource(f *field.AcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayBookingInst is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DayBookingInst() (*field.DayBookingInstField, errors.MessageRejectError) {
	f := &field.DayBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayBookingInst reads a DayBookingInst from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDayBookingInst(f *field.DayBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingUnit is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) BookingUnit() (*field.BookingUnitField, errors.MessageRejectError) {
	f := &field.BookingUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingUnit reads a BookingUnit from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetBookingUnit(f *field.BookingUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreallocMethod is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PreallocMethod() (*field.PreallocMethodField, errors.MessageRejectError) {
	f := &field.PreallocMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreallocMethod reads a PreallocMethod from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPreallocMethod(f *field.PreallocMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) AllocID() (*field.AllocIDField, errors.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetAllocID(f *field.AllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoAllocs() (*field.NoAllocsField, errors.MessageRejectError) {
	f := &field.NoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoAllocs(f *field.NoAllocsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SettlType() (*field.SettlTypeField, errors.MessageRejectError) {
	f := &field.SettlTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSettlType(f *field.SettlTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SettlDate() (*field.SettlDateField, errors.MessageRejectError) {
	f := &field.SettlDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSettlDate(f *field.SettlDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashMargin is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CashMargin() (*field.CashMarginField, errors.MessageRejectError) {
	f := &field.CashMarginField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashMargin reads a CashMargin from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCashMargin(f *field.CashMarginField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFeeIndicator is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ClearingFeeIndicator() (*field.ClearingFeeIndicatorField, errors.MessageRejectError) {
	f := &field.ClearingFeeIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFeeIndicator reads a ClearingFeeIndicator from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetClearingFeeIndicator(f *field.ClearingFeeIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) HandlInst() (*field.HandlInstField, errors.MessageRejectError) {
	f := &field.HandlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetHandlInst(f *field.HandlInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ExecInst() (*field.ExecInstField, errors.MessageRejectError) {
	f := &field.ExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetExecInst(f *field.ExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MinQty() (*field.MinQtyField, errors.MessageRejectError) {
	f := &field.MinQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMinQty(f *field.MinQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MaxFloor() (*field.MaxFloorField, errors.MessageRejectError) {
	f := &field.MaxFloorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMaxFloor(f *field.MaxFloorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ExDestination() (*field.ExDestinationField, errors.MessageRejectError) {
	f := &field.ExDestinationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetExDestination(f *field.ExDestinationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoTradingSessions() (*field.NoTradingSessionsField, errors.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoTradingSessions(f *field.NoTradingSessionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ProcessCode() (*field.ProcessCodeField, errors.MessageRejectError) {
	f := &field.ProcessCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetProcessCode(f *field.ProcessCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecurityStatus() (*field.SecurityStatusField, errors.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecurityStatus(f *field.SecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) StrikeMultiplier() (*field.StrikeMultiplierField, errors.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetStrikeMultiplier(f *field.StrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) StrikeValue() (*field.StrikeValueField, errors.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetStrikeValue(f *field.StrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MinPriceIncrement() (*field.MinPriceIncrementField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMinPriceIncrement(f *field.MinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PositionLimit() (*field.PositionLimitField, errors.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPositionLimit(f *field.PositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NTPositionLimit() (*field.NTPositionLimitField, errors.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNTPositionLimit(f *field.NTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoInstrumentParties() (*field.NoInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) UnitOfMeasure() (*field.UnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetUnitOfMeasure(f *field.UnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TimeUnit() (*field.TimeUnitField, errors.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTimeUnit(f *field.TimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MaturityTime() (*field.MaturityTimeField, errors.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMaturityTime(f *field.MaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PrevClosePx() (*field.PrevClosePxField, errors.MessageRejectError) {
	f := &field.PrevClosePxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPrevClosePx(f *field.PrevClosePxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) LocateReqd() (*field.LocateReqdField, errors.MessageRejectError) {
	f := &field.LocateReqdField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetLocateReqd(f *field.LocateReqdField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) QtyType() (*field.QtyTypeField, errors.MessageRejectError) {
	f := &field.QtyTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetQtyType(f *field.QtyTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrderQty() (*field.OrderQtyField, errors.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOrderQty(f *field.OrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CashOrderQty() (*field.CashOrderQtyField, errors.MessageRejectError) {
	f := &field.CashOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCashOrderQty(f *field.CashOrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrderPercent() (*field.OrderPercentField, errors.MessageRejectError) {
	f := &field.OrderPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOrderPercent(f *field.OrderPercentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RoundingDirection() (*field.RoundingDirectionField, errors.MessageRejectError) {
	f := &field.RoundingDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRoundingDirection(f *field.RoundingDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RoundingModulus() (*field.RoundingModulusField, errors.MessageRejectError) {
	f := &field.RoundingModulusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRoundingModulus(f *field.RoundingModulusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrdType() (*field.OrdTypeField, errors.MessageRejectError) {
	f := &field.OrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOrdType(f *field.OrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) StopPx() (*field.StopPxField, errors.MessageRejectError) {
	f := &field.StopPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetStopPx(f *field.StopPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ComplianceID() (*field.ComplianceIDField, errors.MessageRejectError) {
	f := &field.ComplianceIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetComplianceID(f *field.ComplianceIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SolicitedFlag is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SolicitedFlag() (*field.SolicitedFlagField, errors.MessageRejectError) {
	f := &field.SolicitedFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSolicitedFlag reads a SolicitedFlag from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSolicitedFlag(f *field.SolicitedFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) IOIID() (*field.IOIIDField, errors.MessageRejectError) {
	f := &field.IOIIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIID reads a IOIID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetIOIID(f *field.IOIIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) QuoteID() (*field.QuoteIDField, errors.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetQuoteID(f *field.QuoteIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TimeInForce() (*field.TimeInForceField, errors.MessageRejectError) {
	f := &field.TimeInForceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTimeInForce(f *field.TimeInForceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EffectiveTime() (*field.EffectiveTimeField, errors.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetEffectiveTime(f *field.EffectiveTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ExpireDate() (*field.ExpireDateField, errors.MessageRejectError) {
	f := &field.ExpireDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetExpireDate(f *field.ExpireDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ExpireTime() (*field.ExpireTimeField, errors.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetExpireTime(f *field.ExpireTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GTBookingInst() (*field.GTBookingInstField, errors.MessageRejectError) {
	f := &field.GTBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetGTBookingInst(f *field.GTBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Commission() (*field.CommissionField, errors.MessageRejectError) {
	f := &field.CommissionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCommission(f *field.CommissionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CommType() (*field.CommTypeField, errors.MessageRejectError) {
	f := &field.CommTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCommType(f *field.CommTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommCurrency is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CommCurrency() (*field.CommCurrencyField, errors.MessageRejectError) {
	f := &field.CommCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommCurrency reads a CommCurrency from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCommCurrency(f *field.CommCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FundRenewWaiv is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) FundRenewWaiv() (*field.FundRenewWaivField, errors.MessageRejectError) {
	f := &field.FundRenewWaivField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFundRenewWaiv reads a FundRenewWaiv from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetFundRenewWaiv(f *field.FundRenewWaivField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrderCapacity() (*field.OrderCapacityField, errors.MessageRejectError) {
	f := &field.OrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOrderCapacity(f *field.OrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) OrderRestrictions() (*field.OrderRestrictionsField, errors.MessageRejectError) {
	f := &field.OrderRestrictionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetOrderRestrictions(f *field.OrderRestrictionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CustOrderCapacity() (*field.CustOrderCapacityField, errors.MessageRejectError) {
	f := &field.CustOrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCustOrderCapacity(f *field.CustOrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ForexReq() (*field.ForexReqField, errors.MessageRejectError) {
	f := &field.ForexReqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetForexReq(f *field.ForexReqField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SettlCurrency() (*field.SettlCurrencyField, errors.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSettlCurrency(f *field.SettlCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) BookingType() (*field.BookingTypeField, errors.MessageRejectError) {
	f := &field.BookingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingType reads a BookingType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetBookingType(f *field.BookingTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PositionEffect() (*field.PositionEffectField, errors.MessageRejectError) {
	f := &field.PositionEffectField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPositionEffect(f *field.PositionEffectField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CoveredOrUncovered is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CoveredOrUncovered() (*field.CoveredOrUncoveredField, errors.MessageRejectError) {
	f := &field.CoveredOrUncoveredField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCoveredOrUncovered reads a CoveredOrUncovered from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCoveredOrUncovered(f *field.CoveredOrUncoveredField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MaxShow() (*field.MaxShowField, errors.MessageRejectError) {
	f := &field.MaxShowField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMaxShow(f *field.MaxShowField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetValue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegOffsetValue() (*field.PegOffsetValueField, errors.MessageRejectError) {
	f := &field.PegOffsetValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetValue reads a PegOffsetValue from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegOffsetValue(f *field.PegOffsetValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegMoveType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegMoveType() (*field.PegMoveTypeField, errors.MessageRejectError) {
	f := &field.PegMoveTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegMoveType reads a PegMoveType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegMoveType(f *field.PegMoveTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegOffsetType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegOffsetType() (*field.PegOffsetTypeField, errors.MessageRejectError) {
	f := &field.PegOffsetTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegOffsetType reads a PegOffsetType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegOffsetType(f *field.PegOffsetTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegLimitType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegLimitType() (*field.PegLimitTypeField, errors.MessageRejectError) {
	f := &field.PegLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegLimitType reads a PegLimitType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegLimitType(f *field.PegLimitTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegRoundDirection is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegRoundDirection() (*field.PegRoundDirectionField, errors.MessageRejectError) {
	f := &field.PegRoundDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegRoundDirection reads a PegRoundDirection from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegRoundDirection(f *field.PegRoundDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegScope is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegScope() (*field.PegScopeField, errors.MessageRejectError) {
	f := &field.PegScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegScope reads a PegScope from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegScope(f *field.PegScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegPriceType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegPriceType() (*field.PegPriceTypeField, errors.MessageRejectError) {
	f := &field.PegPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegPriceType reads a PegPriceType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegPriceType(f *field.PegPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityIDSource is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegSecurityIDSource() (*field.PegSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.PegSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityIDSource reads a PegSecurityIDSource from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegSecurityIDSource(f *field.PegSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegSecurityID() (*field.PegSecurityIDField, errors.MessageRejectError) {
	f := &field.PegSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityID reads a PegSecurityID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegSecurityID(f *field.PegSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSymbol is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegSymbol() (*field.PegSymbolField, errors.MessageRejectError) {
	f := &field.PegSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSymbol reads a PegSymbol from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegSymbol(f *field.PegSymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegSecurityDesc is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PegSecurityDesc() (*field.PegSecurityDescField, errors.MessageRejectError) {
	f := &field.PegSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegSecurityDesc reads a PegSecurityDesc from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPegSecurityDesc(f *field.PegSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionInst() (*field.DiscretionInstField, errors.MessageRejectError) {
	f := &field.DiscretionInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDiscretionInst(f *field.DiscretionInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetValue is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionOffsetValue() (*field.DiscretionOffsetValueField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetValue reads a DiscretionOffsetValue from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDiscretionOffsetValue(f *field.DiscretionOffsetValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionMoveType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionMoveType() (*field.DiscretionMoveTypeField, errors.MessageRejectError) {
	f := &field.DiscretionMoveTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionMoveType reads a DiscretionMoveType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDiscretionMoveType(f *field.DiscretionMoveTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffsetType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionOffsetType() (*field.DiscretionOffsetTypeField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffsetType reads a DiscretionOffsetType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDiscretionOffsetType(f *field.DiscretionOffsetTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionLimitType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionLimitType() (*field.DiscretionLimitTypeField, errors.MessageRejectError) {
	f := &field.DiscretionLimitTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionLimitType reads a DiscretionLimitType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDiscretionLimitType(f *field.DiscretionLimitTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionRoundDirection is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionRoundDirection() (*field.DiscretionRoundDirectionField, errors.MessageRejectError) {
	f := &field.DiscretionRoundDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionRoundDirection reads a DiscretionRoundDirection from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDiscretionRoundDirection(f *field.DiscretionRoundDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionScope is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DiscretionScope() (*field.DiscretionScopeField, errors.MessageRejectError) {
	f := &field.DiscretionScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionScope reads a DiscretionScope from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDiscretionScope(f *field.DiscretionScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategy is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TargetStrategy() (*field.TargetStrategyField, errors.MessageRejectError) {
	f := &field.TargetStrategyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategy reads a TargetStrategy from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTargetStrategy(f *field.TargetStrategyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TargetStrategyParameters is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TargetStrategyParameters() (*field.TargetStrategyParametersField, errors.MessageRejectError) {
	f := &field.TargetStrategyParametersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTargetStrategyParameters reads a TargetStrategyParameters from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTargetStrategyParameters(f *field.TargetStrategyParametersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ParticipationRate is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ParticipationRate() (*field.ParticipationRateField, errors.MessageRejectError) {
	f := &field.ParticipationRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetParticipationRate reads a ParticipationRate from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetParticipationRate(f *field.ParticipationRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) CancellationRights() (*field.CancellationRightsField, errors.MessageRejectError) {
	f := &field.CancellationRightsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetCancellationRights(f *field.CancellationRightsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MoneyLaunderingStatus() (*field.MoneyLaunderingStatusField, errors.MessageRejectError) {
	f := &field.MoneyLaunderingStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RegistID() (*field.RegistIDField, errors.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRegistID(f *field.RegistIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) Designation() (*field.DesignationField, errors.MessageRejectError) {
	f := &field.DesignationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDesignation(f *field.DesignationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegRptTypeReq is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MultiLegRptTypeReq() (*field.MultiLegRptTypeReqField, errors.MessageRejectError) {
	f := &field.MultiLegRptTypeReqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegRptTypeReq reads a MultiLegRptTypeReq from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMultiLegRptTypeReq(f *field.MultiLegRptTypeReqField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStrategyParameters is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) NoStrategyParameters() (*field.NoStrategyParametersField, errors.MessageRejectError) {
	f := &field.NoStrategyParametersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStrategyParameters reads a NoStrategyParameters from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetNoStrategyParameters(f *field.NoStrategyParametersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchIncrement is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MatchIncrement() (*field.MatchIncrementField, errors.MessageRejectError) {
	f := &field.MatchIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchIncrement reads a MatchIncrement from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMatchIncrement(f *field.MatchIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxPriceLevels is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) MaxPriceLevels() (*field.MaxPriceLevelsField, errors.MessageRejectError) {
	f := &field.MaxPriceLevelsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxPriceLevels reads a MaxPriceLevels from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetMaxPriceLevels(f *field.MaxPriceLevelsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryDisplayQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SecondaryDisplayQty() (*field.SecondaryDisplayQtyField, errors.MessageRejectError) {
	f := &field.SecondaryDisplayQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryDisplayQty reads a SecondaryDisplayQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSecondaryDisplayQty(f *field.SecondaryDisplayQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayWhen is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DisplayWhen() (*field.DisplayWhenField, errors.MessageRejectError) {
	f := &field.DisplayWhenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayWhen reads a DisplayWhen from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDisplayWhen(f *field.DisplayWhenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMethod is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DisplayMethod() (*field.DisplayMethodField, errors.MessageRejectError) {
	f := &field.DisplayMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMethod reads a DisplayMethod from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDisplayMethod(f *field.DisplayMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayLowQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DisplayLowQty() (*field.DisplayLowQtyField, errors.MessageRejectError) {
	f := &field.DisplayLowQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayLowQty reads a DisplayLowQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDisplayLowQty(f *field.DisplayLowQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayHighQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DisplayHighQty() (*field.DisplayHighQtyField, errors.MessageRejectError) {
	f := &field.DisplayHighQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayHighQty reads a DisplayHighQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDisplayHighQty(f *field.DisplayHighQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayMinIncr is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DisplayMinIncr() (*field.DisplayMinIncrField, errors.MessageRejectError) {
	f := &field.DisplayMinIncrField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayMinIncr reads a DisplayMinIncr from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDisplayMinIncr(f *field.DisplayMinIncrField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefreshQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) RefreshQty() (*field.RefreshQtyField, errors.MessageRejectError) {
	f := &field.RefreshQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefreshQty reads a RefreshQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetRefreshQty(f *field.RefreshQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DisplayQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) DisplayQty() (*field.DisplayQtyField, errors.MessageRejectError) {
	f := &field.DisplayQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDisplayQty reads a DisplayQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetDisplayQty(f *field.DisplayQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceProtectionScope is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PriceProtectionScope() (*field.PriceProtectionScopeField, errors.MessageRejectError) {
	f := &field.PriceProtectionScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceProtectionScope reads a PriceProtectionScope from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPriceProtectionScope(f *field.PriceProtectionScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerType() (*field.TriggerTypeField, errors.MessageRejectError) {
	f := &field.TriggerTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerType reads a TriggerType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerType(f *field.TriggerTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerAction is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerAction() (*field.TriggerActionField, errors.MessageRejectError) {
	f := &field.TriggerActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerAction reads a TriggerAction from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerAction(f *field.TriggerActionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPrice is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerPrice() (*field.TriggerPriceField, errors.MessageRejectError) {
	f := &field.TriggerPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPrice reads a TriggerPrice from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerPrice(f *field.TriggerPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSymbol is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerSymbol() (*field.TriggerSymbolField, errors.MessageRejectError) {
	f := &field.TriggerSymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSymbol reads a TriggerSymbol from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerSymbol(f *field.TriggerSymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerSecurityID() (*field.TriggerSecurityIDField, errors.MessageRejectError) {
	f := &field.TriggerSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityID reads a TriggerSecurityID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerSecurityID(f *field.TriggerSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityIDSource is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerSecurityIDSource() (*field.TriggerSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.TriggerSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityIDSource reads a TriggerSecurityIDSource from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerSecurityIDSource(f *field.TriggerSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerSecurityDesc is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerSecurityDesc() (*field.TriggerSecurityDescField, errors.MessageRejectError) {
	f := &field.TriggerSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerSecurityDesc reads a TriggerSecurityDesc from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerSecurityDesc(f *field.TriggerSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerPriceType() (*field.TriggerPriceTypeField, errors.MessageRejectError) {
	f := &field.TriggerPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceType reads a TriggerPriceType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerPriceType(f *field.TriggerPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceTypeScope is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerPriceTypeScope() (*field.TriggerPriceTypeScopeField, errors.MessageRejectError) {
	f := &field.TriggerPriceTypeScopeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceTypeScope reads a TriggerPriceTypeScope from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerPriceTypeScope(f *field.TriggerPriceTypeScopeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerPriceDirection is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerPriceDirection() (*field.TriggerPriceDirectionField, errors.MessageRejectError) {
	f := &field.TriggerPriceDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerPriceDirection reads a TriggerPriceDirection from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerPriceDirection(f *field.TriggerPriceDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewPrice is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerNewPrice() (*field.TriggerNewPriceField, errors.MessageRejectError) {
	f := &field.TriggerNewPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewPrice reads a TriggerNewPrice from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerNewPrice(f *field.TriggerNewPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerOrderType is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerOrderType() (*field.TriggerOrderTypeField, errors.MessageRejectError) {
	f := &field.TriggerOrderTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerOrderType reads a TriggerOrderType from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerOrderType(f *field.TriggerOrderTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerNewQty is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerNewQty() (*field.TriggerNewQtyField, errors.MessageRejectError) {
	f := &field.TriggerNewQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerNewQty reads a TriggerNewQty from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerNewQty(f *field.TriggerNewQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerTradingSessionID() (*field.TriggerTradingSessionIDField, errors.MessageRejectError) {
	f := &field.TriggerTradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionID reads a TriggerTradingSessionID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerTradingSessionID(f *field.TriggerTradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TriggerTradingSessionSubID is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) TriggerTradingSessionSubID() (*field.TriggerTradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TriggerTradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTriggerTradingSessionSubID reads a TriggerTradingSessionSubID from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetTriggerTradingSessionSubID(f *field.TriggerTradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreTradeAnonymity is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) PreTradeAnonymity() (*field.PreTradeAnonymityField, errors.MessageRejectError) {
	f := &field.PreTradeAnonymityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreTradeAnonymity reads a PreTradeAnonymity from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetPreTradeAnonymity(f *field.PreTradeAnonymityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestinationIDSource is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) ExDestinationIDSource() (*field.ExDestinationIDSourceField, errors.MessageRejectError) {
	f := &field.ExDestinationIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDestinationIDSource reads a ExDestinationIDSource from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetExDestinationIDSource(f *field.ExDestinationIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SwapPoints is a non-required field for MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) SwapPoints() (*field.SwapPointsField, errors.MessageRejectError) {
	f := &field.SwapPointsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSwapPoints reads a SwapPoints from MultilegOrderCancelReplace.
func (m MultilegOrderCancelReplace) GetSwapPoints(f *field.SwapPointsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
