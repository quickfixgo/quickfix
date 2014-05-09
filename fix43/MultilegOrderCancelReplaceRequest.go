package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	origclordid *field.OrigClOrdIDField,
	clordid *field.ClOrdIDField,
	handlinst *field.HandlInstField,
	side *field.SideField,
	nolegs *field.NoLegsField,
	transacttime *field.TransactTimeField,
	ordtype *field.OrdTypeField) MultilegOrderCancelReplaceRequestBuilder {
	var builder MultilegOrderCancelReplaceRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX43))
	builder.Header.Set(field.NewMsgType("AC"))
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
func (m MultilegOrderCancelReplaceRequest) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigClOrdID is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrigClOrdID() (*field.OrigClOrdIDField, errors.MessageRejectError) {
	f := &field.OrigClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigClOrdID reads a OrigClOrdID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOrigClOrdID(f *field.OrigClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecondaryClOrdID() (*field.SecondaryClOrdIDField, errors.MessageRejectError) {
	f := &field.SecondaryClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSecondaryClOrdID(f *field.SecondaryClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ClOrdLinkID() (*field.ClOrdLinkIDField, errors.MessageRejectError) {
	f := &field.ClOrdLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetClOrdLinkID(f *field.ClOrdLinkIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrigOrdModTime is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrigOrdModTime() (*field.OrigOrdModTimeField, errors.MessageRejectError) {
	f := &field.OrigOrdModTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigOrdModTime reads a OrigOrdModTime from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOrigOrdModTime(f *field.OrigOrdModTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayBookingInst is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) DayBookingInst() (*field.DayBookingInstField, errors.MessageRejectError) {
	f := &field.DayBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayBookingInst reads a DayBookingInst from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetDayBookingInst(f *field.DayBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingUnit is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) BookingUnit() (*field.BookingUnitField, errors.MessageRejectError) {
	f := &field.BookingUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingUnit reads a BookingUnit from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetBookingUnit(f *field.BookingUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreallocMethod is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) PreallocMethod() (*field.PreallocMethodField, errors.MessageRejectError) {
	f := &field.PreallocMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreallocMethod reads a PreallocMethod from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetPreallocMethod(f *field.PreallocMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NoAllocs() (*field.NoAllocsField, errors.MessageRejectError) {
	f := &field.NoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetNoAllocs(f *field.NoAllocsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SettlmntTyp() (*field.SettlmntTypField, errors.MessageRejectError) {
	f := &field.SettlmntTypField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSettlmntTyp(f *field.SettlmntTypField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) FutSettDate() (*field.FutSettDateField, errors.MessageRejectError) {
	f := &field.FutSettDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetFutSettDate(f *field.FutSettDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashMargin is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CashMargin() (*field.CashMarginField, errors.MessageRejectError) {
	f := &field.CashMarginField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashMargin reads a CashMargin from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCashMargin(f *field.CashMarginField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFeeIndicator is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ClearingFeeIndicator() (*field.ClearingFeeIndicatorField, errors.MessageRejectError) {
	f := &field.ClearingFeeIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFeeIndicator reads a ClearingFeeIndicator from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetClearingFeeIndicator(f *field.ClearingFeeIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) HandlInst() (*field.HandlInstField, errors.MessageRejectError) {
	f := &field.HandlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetHandlInst(f *field.HandlInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ExecInst() (*field.ExecInstField, errors.MessageRejectError) {
	f := &field.ExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetExecInst(f *field.ExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MinQty() (*field.MinQtyField, errors.MessageRejectError) {
	f := &field.MinQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetMinQty(f *field.MinQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MaxFloor() (*field.MaxFloorField, errors.MessageRejectError) {
	f := &field.MaxFloorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetMaxFloor(f *field.MaxFloorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ExDestination() (*field.ExDestinationField, errors.MessageRejectError) {
	f := &field.ExDestinationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetExDestination(f *field.ExDestinationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NoTradingSessions() (*field.NoTradingSessionsField, errors.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetNoTradingSessions(f *field.NoTradingSessionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ProcessCode() (*field.ProcessCodeField, errors.MessageRejectError) {
	f := &field.ProcessCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetProcessCode(f *field.ProcessCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) PrevClosePx() (*field.PrevClosePxField, errors.MessageRejectError) {
	f := &field.PrevClosePxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetPrevClosePx(f *field.PrevClosePxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) LocateReqd() (*field.LocateReqdField, errors.MessageRejectError) {
	f := &field.LocateReqdField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetLocateReqd(f *field.LocateReqdField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuantityType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) QuantityType() (*field.QuantityTypeField, errors.MessageRejectError) {
	f := &field.QuantityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuantityType reads a QuantityType from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetQuantityType(f *field.QuantityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrderQty() (*field.OrderQtyField, errors.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOrderQty(f *field.OrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CashOrderQty() (*field.CashOrderQtyField, errors.MessageRejectError) {
	f := &field.CashOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCashOrderQty(f *field.CashOrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrderPercent() (*field.OrderPercentField, errors.MessageRejectError) {
	f := &field.OrderPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOrderPercent(f *field.OrderPercentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RoundingDirection() (*field.RoundingDirectionField, errors.MessageRejectError) {
	f := &field.RoundingDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetRoundingDirection(f *field.RoundingDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RoundingModulus() (*field.RoundingModulusField, errors.MessageRejectError) {
	f := &field.RoundingModulusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetRoundingModulus(f *field.RoundingModulusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrdType() (*field.OrdTypeField, errors.MessageRejectError) {
	f := &field.OrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOrdType(f *field.OrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) StopPx() (*field.StopPxField, errors.MessageRejectError) {
	f := &field.StopPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetStopPx(f *field.StopPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ComplianceID() (*field.ComplianceIDField, errors.MessageRejectError) {
	f := &field.ComplianceIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetComplianceID(f *field.ComplianceIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SolicitedFlag is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SolicitedFlag() (*field.SolicitedFlagField, errors.MessageRejectError) {
	f := &field.SolicitedFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSolicitedFlag reads a SolicitedFlag from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSolicitedFlag(f *field.SolicitedFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIid is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) IOIid() (*field.IOIidField, errors.MessageRejectError) {
	f := &field.IOIidField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIid reads a IOIid from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetIOIid(f *field.IOIidField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) QuoteID() (*field.QuoteIDField, errors.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetQuoteID(f *field.QuoteIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) TimeInForce() (*field.TimeInForceField, errors.MessageRejectError) {
	f := &field.TimeInForceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetTimeInForce(f *field.TimeInForceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EffectiveTime() (*field.EffectiveTimeField, errors.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetEffectiveTime(f *field.EffectiveTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ExpireDate() (*field.ExpireDateField, errors.MessageRejectError) {
	f := &field.ExpireDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetExpireDate(f *field.ExpireDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ExpireTime() (*field.ExpireTimeField, errors.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetExpireTime(f *field.ExpireTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GTBookingInst() (*field.GTBookingInstField, errors.MessageRejectError) {
	f := &field.GTBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetGTBookingInst(f *field.GTBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Commission() (*field.CommissionField, errors.MessageRejectError) {
	f := &field.CommissionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCommission(f *field.CommissionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CommType() (*field.CommTypeField, errors.MessageRejectError) {
	f := &field.CommTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCommType(f *field.CommTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommCurrency is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CommCurrency() (*field.CommCurrencyField, errors.MessageRejectError) {
	f := &field.CommCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommCurrency reads a CommCurrency from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCommCurrency(f *field.CommCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FundRenewWaiv is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) FundRenewWaiv() (*field.FundRenewWaivField, errors.MessageRejectError) {
	f := &field.FundRenewWaivField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFundRenewWaiv reads a FundRenewWaiv from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetFundRenewWaiv(f *field.FundRenewWaivField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrderCapacity() (*field.OrderCapacityField, errors.MessageRejectError) {
	f := &field.OrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOrderCapacity(f *field.OrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) OrderRestrictions() (*field.OrderRestrictionsField, errors.MessageRejectError) {
	f := &field.OrderRestrictionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetOrderRestrictions(f *field.OrderRestrictionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CustOrderCapacity() (*field.CustOrderCapacityField, errors.MessageRejectError) {
	f := &field.CustOrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCustOrderCapacity(f *field.CustOrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) ForexReq() (*field.ForexReqField, errors.MessageRejectError) {
	f := &field.ForexReqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetForexReq(f *field.ForexReqField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) SettlCurrency() (*field.SettlCurrencyField, errors.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetSettlCurrency(f *field.SettlCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) PositionEffect() (*field.PositionEffectField, errors.MessageRejectError) {
	f := &field.PositionEffectField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetPositionEffect(f *field.PositionEffectField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CoveredOrUncovered is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CoveredOrUncovered() (*field.CoveredOrUncoveredField, errors.MessageRejectError) {
	f := &field.CoveredOrUncoveredField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCoveredOrUncovered reads a CoveredOrUncovered from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCoveredOrUncovered(f *field.CoveredOrUncoveredField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MaxShow() (*field.MaxShowField, errors.MessageRejectError) {
	f := &field.MaxShowField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetMaxShow(f *field.MaxShowField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegDifference is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) PegDifference() (*field.PegDifferenceField, errors.MessageRejectError) {
	f := &field.PegDifferenceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegDifference reads a PegDifference from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetPegDifference(f *field.PegDifferenceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) DiscretionInst() (*field.DiscretionInstField, errors.MessageRejectError) {
	f := &field.DiscretionInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetDiscretionInst(f *field.DiscretionInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffset is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) DiscretionOffset() (*field.DiscretionOffsetField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffset reads a DiscretionOffset from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetDiscretionOffset(f *field.DiscretionOffsetField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) CancellationRights() (*field.CancellationRightsField, errors.MessageRejectError) {
	f := &field.CancellationRightsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetCancellationRights(f *field.CancellationRightsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MoneyLaunderingStatus() (*field.MoneyLaunderingStatusField, errors.MessageRejectError) {
	f := &field.MoneyLaunderingStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) RegistID() (*field.RegistIDField, errors.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetRegistID(f *field.RegistIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) Designation() (*field.DesignationField, errors.MessageRejectError) {
	f := &field.DesignationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetDesignation(f *field.DesignationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegRptTypeReq is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) MultiLegRptTypeReq() (*field.MultiLegRptTypeReqField, errors.MessageRejectError) {
	f := &field.MultiLegRptTypeReqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegRptTypeReq reads a MultiLegRptTypeReq from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetMultiLegRptTypeReq(f *field.MultiLegRptTypeReqField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a non-required field for MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) NetMoney() (*field.NetMoneyField, errors.MessageRejectError) {
	f := &field.NetMoneyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from MultilegOrderCancelReplaceRequest.
func (m MultilegOrderCancelReplaceRequest) GetNetMoney(f *field.NetMoneyField) errors.MessageRejectError {
	return m.Body.Get(f)
}
