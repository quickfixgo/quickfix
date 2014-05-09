package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	clordid *field.ClOrdIDField,
	handlinst *field.HandlInstField,
	side *field.SideField,
	nolegs *field.NoLegsField,
	transacttime *field.TransactTimeField,
	ordtype *field.OrdTypeField) NewOrderMultilegBuilder {
	var builder NewOrderMultilegBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX43))
	builder.Header.Set(field.NewMsgType("AB"))
	builder.Body.Set(clordid)
	builder.Body.Set(handlinst)
	builder.Body.Set(side)
	builder.Body.Set(nolegs)
	builder.Body.Set(transacttime)
	builder.Body.Set(ordtype)
	return builder
}

//ClOrdID is a required field for NewOrderMultileg.
func (m NewOrderMultileg) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from NewOrderMultileg.
func (m NewOrderMultileg) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecondaryClOrdID() (*field.SecondaryClOrdIDField, errors.MessageRejectError) {
	f := &field.SecondaryClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from NewOrderMultileg.
func (m NewOrderMultileg) GetSecondaryClOrdID(f *field.SecondaryClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdLinkID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ClOrdLinkID() (*field.ClOrdLinkIDField, errors.MessageRejectError) {
	f := &field.ClOrdLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdLinkID reads a ClOrdLinkID from NewOrderMultileg.
func (m NewOrderMultileg) GetClOrdLinkID(f *field.ClOrdLinkIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from NewOrderMultileg.
func (m NewOrderMultileg) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from NewOrderMultileg.
func (m NewOrderMultileg) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from NewOrderMultileg.
func (m NewOrderMultileg) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DayBookingInst is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DayBookingInst() (*field.DayBookingInstField, errors.MessageRejectError) {
	f := &field.DayBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDayBookingInst reads a DayBookingInst from NewOrderMultileg.
func (m NewOrderMultileg) GetDayBookingInst(f *field.DayBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingUnit is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) BookingUnit() (*field.BookingUnitField, errors.MessageRejectError) {
	f := &field.BookingUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingUnit reads a BookingUnit from NewOrderMultileg.
func (m NewOrderMultileg) GetBookingUnit(f *field.BookingUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreallocMethod is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PreallocMethod() (*field.PreallocMethodField, errors.MessageRejectError) {
	f := &field.PreallocMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreallocMethod reads a PreallocMethod from NewOrderMultileg.
func (m NewOrderMultileg) GetPreallocMethod(f *field.PreallocMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoAllocs() (*field.NoAllocsField, errors.MessageRejectError) {
	f := &field.NoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from NewOrderMultileg.
func (m NewOrderMultileg) GetNoAllocs(f *field.NoAllocsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SettlmntTyp() (*field.SettlmntTypField, errors.MessageRejectError) {
	f := &field.SettlmntTypField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from NewOrderMultileg.
func (m NewOrderMultileg) GetSettlmntTyp(f *field.SettlmntTypField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) FutSettDate() (*field.FutSettDateField, errors.MessageRejectError) {
	f := &field.FutSettDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from NewOrderMultileg.
func (m NewOrderMultileg) GetFutSettDate(f *field.FutSettDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashMargin is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CashMargin() (*field.CashMarginField, errors.MessageRejectError) {
	f := &field.CashMarginField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashMargin reads a CashMargin from NewOrderMultileg.
func (m NewOrderMultileg) GetCashMargin(f *field.CashMarginField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFeeIndicator is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ClearingFeeIndicator() (*field.ClearingFeeIndicatorField, errors.MessageRejectError) {
	f := &field.ClearingFeeIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFeeIndicator reads a ClearingFeeIndicator from NewOrderMultileg.
func (m NewOrderMultileg) GetClearingFeeIndicator(f *field.ClearingFeeIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//HandlInst is a required field for NewOrderMultileg.
func (m NewOrderMultileg) HandlInst() (*field.HandlInstField, errors.MessageRejectError) {
	f := &field.HandlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHandlInst reads a HandlInst from NewOrderMultileg.
func (m NewOrderMultileg) GetHandlInst(f *field.HandlInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecInst is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ExecInst() (*field.ExecInstField, errors.MessageRejectError) {
	f := &field.ExecInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecInst reads a ExecInst from NewOrderMultileg.
func (m NewOrderMultileg) GetExecInst(f *field.ExecInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MinQty() (*field.MinQtyField, errors.MessageRejectError) {
	f := &field.MinQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinQty reads a MinQty from NewOrderMultileg.
func (m NewOrderMultileg) GetMinQty(f *field.MinQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxFloor is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MaxFloor() (*field.MaxFloorField, errors.MessageRejectError) {
	f := &field.MaxFloorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxFloor reads a MaxFloor from NewOrderMultileg.
func (m NewOrderMultileg) GetMaxFloor(f *field.MaxFloorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExDestination is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ExDestination() (*field.ExDestinationField, errors.MessageRejectError) {
	f := &field.ExDestinationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExDestination reads a ExDestination from NewOrderMultileg.
func (m NewOrderMultileg) GetExDestination(f *field.ExDestinationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoTradingSessions() (*field.NoTradingSessionsField, errors.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from NewOrderMultileg.
func (m NewOrderMultileg) GetNoTradingSessions(f *field.NoTradingSessionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProcessCode is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ProcessCode() (*field.ProcessCodeField, errors.MessageRejectError) {
	f := &field.ProcessCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProcessCode reads a ProcessCode from NewOrderMultileg.
func (m NewOrderMultileg) GetProcessCode(f *field.ProcessCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for NewOrderMultileg.
func (m NewOrderMultileg) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from NewOrderMultileg.
func (m NewOrderMultileg) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from NewOrderMultileg.
func (m NewOrderMultileg) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from NewOrderMultileg.
func (m NewOrderMultileg) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from NewOrderMultileg.
func (m NewOrderMultileg) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from NewOrderMultileg.
func (m NewOrderMultileg) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from NewOrderMultileg.
func (m NewOrderMultileg) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from NewOrderMultileg.
func (m NewOrderMultileg) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from NewOrderMultileg.
func (m NewOrderMultileg) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from NewOrderMultileg.
func (m NewOrderMultileg) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from NewOrderMultileg.
func (m NewOrderMultileg) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from NewOrderMultileg.
func (m NewOrderMultileg) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from NewOrderMultileg.
func (m NewOrderMultileg) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from NewOrderMultileg.
func (m NewOrderMultileg) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from NewOrderMultileg.
func (m NewOrderMultileg) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from NewOrderMultileg.
func (m NewOrderMultileg) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from NewOrderMultileg.
func (m NewOrderMultileg) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from NewOrderMultileg.
func (m NewOrderMultileg) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from NewOrderMultileg.
func (m NewOrderMultileg) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from NewOrderMultileg.
func (m NewOrderMultileg) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from NewOrderMultileg.
func (m NewOrderMultileg) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from NewOrderMultileg.
func (m NewOrderMultileg) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from NewOrderMultileg.
func (m NewOrderMultileg) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from NewOrderMultileg.
func (m NewOrderMultileg) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from NewOrderMultileg.
func (m NewOrderMultileg) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from NewOrderMultileg.
func (m NewOrderMultileg) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from NewOrderMultileg.
func (m NewOrderMultileg) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from NewOrderMultileg.
func (m NewOrderMultileg) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from NewOrderMultileg.
func (m NewOrderMultileg) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from NewOrderMultileg.
func (m NewOrderMultileg) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from NewOrderMultileg.
func (m NewOrderMultileg) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from NewOrderMultileg.
func (m NewOrderMultileg) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from NewOrderMultileg.
func (m NewOrderMultileg) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from NewOrderMultileg.
func (m NewOrderMultileg) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from NewOrderMultileg.
func (m NewOrderMultileg) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PrevClosePx is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PrevClosePx() (*field.PrevClosePxField, errors.MessageRejectError) {
	f := &field.PrevClosePxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrevClosePx reads a PrevClosePx from NewOrderMultileg.
func (m NewOrderMultileg) GetPrevClosePx(f *field.PrevClosePxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a required field for NewOrderMultileg.
func (m NewOrderMultileg) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from NewOrderMultileg.
func (m NewOrderMultileg) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocateReqd is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) LocateReqd() (*field.LocateReqdField, errors.MessageRejectError) {
	f := &field.LocateReqdField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocateReqd reads a LocateReqd from NewOrderMultileg.
func (m NewOrderMultileg) GetLocateReqd(f *field.LocateReqdField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for NewOrderMultileg.
func (m NewOrderMultileg) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from NewOrderMultileg.
func (m NewOrderMultileg) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuantityType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) QuantityType() (*field.QuantityTypeField, errors.MessageRejectError) {
	f := &field.QuantityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuantityType reads a QuantityType from NewOrderMultileg.
func (m NewOrderMultileg) GetQuantityType(f *field.QuantityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OrderQty() (*field.OrderQtyField, errors.MessageRejectError) {
	f := &field.OrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderQty reads a OrderQty from NewOrderMultileg.
func (m NewOrderMultileg) GetOrderQty(f *field.OrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOrderQty is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CashOrderQty() (*field.CashOrderQtyField, errors.MessageRejectError) {
	f := &field.CashOrderQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashOrderQty reads a CashOrderQty from NewOrderMultileg.
func (m NewOrderMultileg) GetCashOrderQty(f *field.CashOrderQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderPercent is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OrderPercent() (*field.OrderPercentField, errors.MessageRejectError) {
	f := &field.OrderPercentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderPercent reads a OrderPercent from NewOrderMultileg.
func (m NewOrderMultileg) GetOrderPercent(f *field.OrderPercentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingDirection is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RoundingDirection() (*field.RoundingDirectionField, errors.MessageRejectError) {
	f := &field.RoundingDirectionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingDirection reads a RoundingDirection from NewOrderMultileg.
func (m NewOrderMultileg) GetRoundingDirection(f *field.RoundingDirectionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RoundingModulus is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RoundingModulus() (*field.RoundingModulusField, errors.MessageRejectError) {
	f := &field.RoundingModulusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRoundingModulus reads a RoundingModulus from NewOrderMultileg.
func (m NewOrderMultileg) GetRoundingModulus(f *field.RoundingModulusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrdType is a required field for NewOrderMultileg.
func (m NewOrderMultileg) OrdType() (*field.OrdTypeField, errors.MessageRejectError) {
	f := &field.OrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrdType reads a OrdType from NewOrderMultileg.
func (m NewOrderMultileg) GetOrdType(f *field.OrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from NewOrderMultileg.
func (m NewOrderMultileg) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from NewOrderMultileg.
func (m NewOrderMultileg) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StopPx is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) StopPx() (*field.StopPxField, errors.MessageRejectError) {
	f := &field.StopPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStopPx reads a StopPx from NewOrderMultileg.
func (m NewOrderMultileg) GetStopPx(f *field.StopPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from NewOrderMultileg.
func (m NewOrderMultileg) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ComplianceID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ComplianceID() (*field.ComplianceIDField, errors.MessageRejectError) {
	f := &field.ComplianceIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetComplianceID reads a ComplianceID from NewOrderMultileg.
func (m NewOrderMultileg) GetComplianceID(f *field.ComplianceIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SolicitedFlag is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SolicitedFlag() (*field.SolicitedFlagField, errors.MessageRejectError) {
	f := &field.SolicitedFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSolicitedFlag reads a SolicitedFlag from NewOrderMultileg.
func (m NewOrderMultileg) GetSolicitedFlag(f *field.SolicitedFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IOIid is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) IOIid() (*field.IOIidField, errors.MessageRejectError) {
	f := &field.IOIidField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIOIid reads a IOIid from NewOrderMultileg.
func (m NewOrderMultileg) GetIOIid(f *field.IOIidField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QuoteID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) QuoteID() (*field.QuoteIDField, errors.MessageRejectError) {
	f := &field.QuoteIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuoteID reads a QuoteID from NewOrderMultileg.
func (m NewOrderMultileg) GetQuoteID(f *field.QuoteIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeInForce is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) TimeInForce() (*field.TimeInForceField, errors.MessageRejectError) {
	f := &field.TimeInForceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeInForce reads a TimeInForce from NewOrderMultileg.
func (m NewOrderMultileg) GetTimeInForce(f *field.TimeInForceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EffectiveTime() (*field.EffectiveTimeField, errors.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from NewOrderMultileg.
func (m NewOrderMultileg) GetEffectiveTime(f *field.EffectiveTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireDate is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ExpireDate() (*field.ExpireDateField, errors.MessageRejectError) {
	f := &field.ExpireDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireDate reads a ExpireDate from NewOrderMultileg.
func (m NewOrderMultileg) GetExpireDate(f *field.ExpireDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ExpireTime() (*field.ExpireTimeField, errors.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from NewOrderMultileg.
func (m NewOrderMultileg) GetExpireTime(f *field.ExpireTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GTBookingInst is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) GTBookingInst() (*field.GTBookingInstField, errors.MessageRejectError) {
	f := &field.GTBookingInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGTBookingInst reads a GTBookingInst from NewOrderMultileg.
func (m NewOrderMultileg) GetGTBookingInst(f *field.GTBookingInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Commission is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Commission() (*field.CommissionField, errors.MessageRejectError) {
	f := &field.CommissionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommission reads a Commission from NewOrderMultileg.
func (m NewOrderMultileg) GetCommission(f *field.CommissionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommType is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CommType() (*field.CommTypeField, errors.MessageRejectError) {
	f := &field.CommTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommType reads a CommType from NewOrderMultileg.
func (m NewOrderMultileg) GetCommType(f *field.CommTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CommCurrency is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CommCurrency() (*field.CommCurrencyField, errors.MessageRejectError) {
	f := &field.CommCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCommCurrency reads a CommCurrency from NewOrderMultileg.
func (m NewOrderMultileg) GetCommCurrency(f *field.CommCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FundRenewWaiv is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) FundRenewWaiv() (*field.FundRenewWaivField, errors.MessageRejectError) {
	f := &field.FundRenewWaivField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFundRenewWaiv reads a FundRenewWaiv from NewOrderMultileg.
func (m NewOrderMultileg) GetFundRenewWaiv(f *field.FundRenewWaivField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderCapacity is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OrderCapacity() (*field.OrderCapacityField, errors.MessageRejectError) {
	f := &field.OrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderCapacity reads a OrderCapacity from NewOrderMultileg.
func (m NewOrderMultileg) GetOrderCapacity(f *field.OrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderRestrictions is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) OrderRestrictions() (*field.OrderRestrictionsField, errors.MessageRejectError) {
	f := &field.OrderRestrictionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderRestrictions reads a OrderRestrictions from NewOrderMultileg.
func (m NewOrderMultileg) GetOrderRestrictions(f *field.OrderRestrictionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CustOrderCapacity() (*field.CustOrderCapacityField, errors.MessageRejectError) {
	f := &field.CustOrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from NewOrderMultileg.
func (m NewOrderMultileg) GetCustOrderCapacity(f *field.CustOrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ForexReq is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) ForexReq() (*field.ForexReqField, errors.MessageRejectError) {
	f := &field.ForexReqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetForexReq reads a ForexReq from NewOrderMultileg.
func (m NewOrderMultileg) GetForexReq(f *field.ForexReqField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) SettlCurrency() (*field.SettlCurrencyField, errors.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from NewOrderMultileg.
func (m NewOrderMultileg) GetSettlCurrency(f *field.SettlCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from NewOrderMultileg.
func (m NewOrderMultileg) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from NewOrderMultileg.
func (m NewOrderMultileg) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from NewOrderMultileg.
func (m NewOrderMultileg) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PositionEffect() (*field.PositionEffectField, errors.MessageRejectError) {
	f := &field.PositionEffectField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from NewOrderMultileg.
func (m NewOrderMultileg) GetPositionEffect(f *field.PositionEffectField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CoveredOrUncovered is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CoveredOrUncovered() (*field.CoveredOrUncoveredField, errors.MessageRejectError) {
	f := &field.CoveredOrUncoveredField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCoveredOrUncovered reads a CoveredOrUncovered from NewOrderMultileg.
func (m NewOrderMultileg) GetCoveredOrUncovered(f *field.CoveredOrUncoveredField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaxShow is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MaxShow() (*field.MaxShowField, errors.MessageRejectError) {
	f := &field.MaxShowField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaxShow reads a MaxShow from NewOrderMultileg.
func (m NewOrderMultileg) GetMaxShow(f *field.MaxShowField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PegDifference is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) PegDifference() (*field.PegDifferenceField, errors.MessageRejectError) {
	f := &field.PegDifferenceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPegDifference reads a PegDifference from NewOrderMultileg.
func (m NewOrderMultileg) GetPegDifference(f *field.PegDifferenceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionInst is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DiscretionInst() (*field.DiscretionInstField, errors.MessageRejectError) {
	f := &field.DiscretionInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionInst reads a DiscretionInst from NewOrderMultileg.
func (m NewOrderMultileg) GetDiscretionInst(f *field.DiscretionInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DiscretionOffset is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) DiscretionOffset() (*field.DiscretionOffsetField, errors.MessageRejectError) {
	f := &field.DiscretionOffsetField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDiscretionOffset reads a DiscretionOffset from NewOrderMultileg.
func (m NewOrderMultileg) GetDiscretionOffset(f *field.DiscretionOffsetField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CancellationRights is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) CancellationRights() (*field.CancellationRightsField, errors.MessageRejectError) {
	f := &field.CancellationRightsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCancellationRights reads a CancellationRights from NewOrderMultileg.
func (m NewOrderMultileg) GetCancellationRights(f *field.CancellationRightsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MoneyLaunderingStatus is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MoneyLaunderingStatus() (*field.MoneyLaunderingStatusField, errors.MessageRejectError) {
	f := &field.MoneyLaunderingStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMoneyLaunderingStatus reads a MoneyLaunderingStatus from NewOrderMultileg.
func (m NewOrderMultileg) GetMoneyLaunderingStatus(f *field.MoneyLaunderingStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RegistID is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) RegistID() (*field.RegistIDField, errors.MessageRejectError) {
	f := &field.RegistIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRegistID reads a RegistID from NewOrderMultileg.
func (m NewOrderMultileg) GetRegistID(f *field.RegistIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Designation is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) Designation() (*field.DesignationField, errors.MessageRejectError) {
	f := &field.DesignationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDesignation reads a Designation from NewOrderMultileg.
func (m NewOrderMultileg) GetDesignation(f *field.DesignationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegRptTypeReq is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) MultiLegRptTypeReq() (*field.MultiLegRptTypeReqField, errors.MessageRejectError) {
	f := &field.MultiLegRptTypeReqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegRptTypeReq reads a MultiLegRptTypeReq from NewOrderMultileg.
func (m NewOrderMultileg) GetMultiLegRptTypeReq(f *field.MultiLegRptTypeReqField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a non-required field for NewOrderMultileg.
func (m NewOrderMultileg) NetMoney() (*field.NetMoneyField, errors.MessageRejectError) {
	f := &field.NetMoneyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from NewOrderMultileg.
func (m NewOrderMultileg) GetNetMoney(f *field.NetMoneyField) errors.MessageRejectError {
	return m.Body.Get(f)
}
