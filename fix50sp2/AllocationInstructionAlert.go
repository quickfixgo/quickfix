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

//AllocationInstructionAlert msg type = BM.
type AllocationInstructionAlert struct {
	message.Message
}

//AllocationInstructionAlertBuilder builds AllocationInstructionAlert messages.
type AllocationInstructionAlertBuilder struct {
	message.MessageBuilder
}

//CreateAllocationInstructionAlertBuilder returns an initialized AllocationInstructionAlertBuilder with specified required fields.
func CreateAllocationInstructionAlertBuilder(
	allocid *field.AllocIDField,
	alloctranstype *field.AllocTransTypeField,
	alloctype *field.AllocTypeField,
	side *field.SideField,
	quantity *field.QuantityField,
	tradedate *field.TradeDateField) AllocationInstructionAlertBuilder {
	var builder AllocationInstructionAlertBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("BM"))
	builder.Body().Set(allocid)
	builder.Body().Set(alloctranstype)
	builder.Body().Set(alloctype)
	builder.Body().Set(side)
	builder.Body().Set(quantity)
	builder.Body().Set(tradedate)
	return builder
}

//AllocID is a required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AllocID() (*field.AllocIDField, errors.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAllocID(f *field.AllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocTransType is a required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AllocTransType() (*field.AllocTransTypeField, errors.MessageRejectError) {
	f := &field.AllocTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocTransType reads a AllocTransType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAllocTransType(f *field.AllocTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocType is a required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AllocType() (*field.AllocTypeField, errors.MessageRejectError) {
	f := &field.AllocTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocType reads a AllocType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAllocType(f *field.AllocTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryAllocID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecondaryAllocID() (*field.SecondaryAllocIDField, errors.MessageRejectError) {
	f := &field.SecondaryAllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryAllocID reads a SecondaryAllocID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecondaryAllocID(f *field.SecondaryAllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefAllocID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) RefAllocID() (*field.RefAllocIDField, errors.MessageRejectError) {
	f := &field.RefAllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefAllocID reads a RefAllocID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetRefAllocID(f *field.RefAllocIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocCancReplaceReason is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AllocCancReplaceReason() (*field.AllocCancReplaceReasonField, errors.MessageRejectError) {
	f := &field.AllocCancReplaceReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocCancReplaceReason reads a AllocCancReplaceReason from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAllocCancReplaceReason(f *field.AllocCancReplaceReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocIntermedReqType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AllocIntermedReqType() (*field.AllocIntermedReqTypeField, errors.MessageRejectError) {
	f := &field.AllocIntermedReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocIntermedReqType reads a AllocIntermedReqType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAllocIntermedReqType(f *field.AllocIntermedReqTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocLinkID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AllocLinkID() (*field.AllocLinkIDField, errors.MessageRejectError) {
	f := &field.AllocLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocLinkID reads a AllocLinkID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAllocLinkID(f *field.AllocLinkIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocLinkType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AllocLinkType() (*field.AllocLinkTypeField, errors.MessageRejectError) {
	f := &field.AllocLinkTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocLinkType reads a AllocLinkType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAllocLinkType(f *field.AllocLinkTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingRefID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BookingRefID() (*field.BookingRefIDField, errors.MessageRejectError) {
	f := &field.BookingRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingRefID reads a BookingRefID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBookingRefID(f *field.BookingRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocNoOrdersType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AllocNoOrdersType() (*field.AllocNoOrdersTypeField, errors.MessageRejectError) {
	f := &field.AllocNoOrdersTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocNoOrdersType reads a AllocNoOrdersType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAllocNoOrdersType(f *field.AllocNoOrdersTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoOrders() (*field.NoOrdersField, errors.MessageRejectError) {
	f := &field.NoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoOrders(f *field.NoOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecs is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoExecs() (*field.NoExecsField, errors.MessageRejectError) {
	f := &field.NoExecsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecs reads a NoExecs from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoExecs(f *field.NoExecsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreviouslyReported is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) PreviouslyReported() (*field.PreviouslyReportedField, errors.MessageRejectError) {
	f := &field.PreviouslyReportedField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreviouslyReported reads a PreviouslyReported from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetPreviouslyReported(f *field.PreviouslyReportedField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ReversalIndicator is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) ReversalIndicator() (*field.ReversalIndicatorField, errors.MessageRejectError) {
	f := &field.ReversalIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetReversalIndicator reads a ReversalIndicator from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetReversalIndicator(f *field.ReversalIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MatchType() (*field.MatchTypeField, errors.MessageRejectError) {
	f := &field.MatchTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchType reads a MatchType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMatchType(f *field.MatchTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecurityStatus() (*field.SecurityStatusField, errors.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecurityStatus(f *field.SecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StrikeMultiplier() (*field.StrikeMultiplierField, errors.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStrikeMultiplier(f *field.StrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StrikeValue() (*field.StrikeValueField, errors.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStrikeValue(f *field.StrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MinPriceIncrement() (*field.MinPriceIncrementField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMinPriceIncrement(f *field.MinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) PositionLimit() (*field.PositionLimitField, errors.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetPositionLimit(f *field.PositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NTPositionLimit() (*field.NTPositionLimitField, errors.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNTPositionLimit(f *field.NTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoInstrumentParties() (*field.NoInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) UnitOfMeasure() (*field.UnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetUnitOfMeasure(f *field.UnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TimeUnit() (*field.TimeUnitField, errors.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTimeUnit(f *field.TimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MaturityTime() (*field.MaturityTimeField, errors.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMaturityTime(f *field.MaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecurityGroup() (*field.SecurityGroupField, errors.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecurityGroup(f *field.SecurityGroupField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecurityXMLLen() (*field.SecurityXMLLenField, errors.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecurityXMLLen(f *field.SecurityXMLLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecurityXML() (*field.SecurityXMLField, errors.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecurityXML(f *field.SecurityXMLField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecurityXMLSchema() (*field.SecurityXMLSchemaField, errors.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) ProductComplex() (*field.ProductComplexField, errors.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetProductComplex(f *field.ProductComplexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SettlMethod() (*field.SettlMethodField, errors.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSettlMethod(f *field.SettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) ExerciseStyle() (*field.ExerciseStyleField, errors.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetExerciseStyle(f *field.ExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) OptPayoutAmount() (*field.OptPayoutAmountField, errors.MessageRejectError) {
	f := &field.OptPayoutAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetOptPayoutAmount(f *field.OptPayoutAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) PriceQuoteMethod() (*field.PriceQuoteMethodField, errors.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) ListMethod() (*field.ListMethodField, errors.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetListMethod(f *field.ListMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CapPrice() (*field.CapPriceField, errors.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCapPrice(f *field.CapPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) FloorPrice() (*field.FloorPriceField, errors.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetFloorPrice(f *field.FloorPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) FlexibleIndicator() (*field.FlexibleIndicatorField, errors.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetFlexibleIndicator(f *field.FlexibleIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, errors.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) ValuationMethod() (*field.ValuationMethodField, errors.MessageRejectError) {
	f := &field.ValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetValuationMethod(f *field.ValuationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) ContractMultiplierUnit() (*field.ContractMultiplierUnitField, errors.MessageRejectError) {
	f := &field.ContractMultiplierUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetContractMultiplierUnit(f *field.ContractMultiplierUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) FlowScheduleType() (*field.FlowScheduleTypeField, errors.MessageRejectError) {
	f := &field.FlowScheduleTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetFlowScheduleType(f *field.FlowScheduleTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) RestructuringType() (*field.RestructuringTypeField, errors.MessageRejectError) {
	f := &field.RestructuringTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetRestructuringType(f *field.RestructuringTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Seniority() (*field.SeniorityField, errors.MessageRejectError) {
	f := &field.SeniorityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSeniority(f *field.SeniorityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstandingField, errors.MessageRejectError) {
	f := &field.NotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstandingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstandingField, errors.MessageRejectError) {
	f := &field.OriginalNotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstandingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AttachmentPoint() (*field.AttachmentPointField, errors.MessageRejectError) {
	f := &field.AttachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAttachmentPoint(f *field.AttachmentPointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) DetachmentPoint() (*field.DetachmentPointField, errors.MessageRejectError) {
	f := &field.DetachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetDetachmentPoint(f *field.DetachmentPointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethodField, errors.MessageRejectError) {
	f := &field.StrikePriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethodField, errors.MessageRejectError) {
	f := &field.StrikePriceBoundaryMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecisionField, errors.MessageRejectError) {
	f := &field.StrikePriceBoundaryPrecisionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecisionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethodField, errors.MessageRejectError) {
	f := &field.UnderlyingPriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) OptPayoutType() (*field.OptPayoutTypeField, errors.MessageRejectError) {
	f := &field.OptPayoutTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetOptPayoutType(f *field.OptPayoutTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoComplexEvents() (*field.NoComplexEventsField, errors.MessageRejectError) {
	f := &field.NoComplexEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoComplexEvents(f *field.NoComplexEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryForm is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) DeliveryForm() (*field.DeliveryFormField, errors.MessageRejectError) {
	f := &field.DeliveryFormField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryForm reads a DeliveryForm from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetDeliveryForm(f *field.DeliveryFormField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PctAtRisk is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) PctAtRisk() (*field.PctAtRiskField, errors.MessageRejectError) {
	f := &field.PctAtRiskField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPctAtRisk reads a PctAtRisk from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetPctAtRisk(f *field.PctAtRiskField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrAttrib is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoInstrAttrib() (*field.NoInstrAttribField, errors.MessageRejectError) {
	f := &field.NoInstrAttribField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrAttrib reads a NoInstrAttrib from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoInstrAttrib(f *field.NoInstrAttribField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AgreementDesc() (*field.AgreementDescField, errors.MessageRejectError) {
	f := &field.AgreementDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAgreementDesc(f *field.AgreementDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AgreementID() (*field.AgreementIDField, errors.MessageRejectError) {
	f := &field.AgreementIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAgreementID(f *field.AgreementIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AgreementDate() (*field.AgreementDateField, errors.MessageRejectError) {
	f := &field.AgreementDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAgreementDate(f *field.AgreementDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AgreementCurrency() (*field.AgreementCurrencyField, errors.MessageRejectError) {
	f := &field.AgreementCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAgreementCurrency(f *field.AgreementCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TerminationType() (*field.TerminationTypeField, errors.MessageRejectError) {
	f := &field.TerminationTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTerminationType(f *field.TerminationTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StartDate() (*field.StartDateField, errors.MessageRejectError) {
	f := &field.StartDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStartDate(f *field.StartDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EndDate() (*field.EndDateField, errors.MessageRejectError) {
	f := &field.EndDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEndDate(f *field.EndDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) DeliveryType() (*field.DeliveryTypeField, errors.MessageRejectError) {
	f := &field.DeliveryTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetDeliveryType(f *field.DeliveryTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MarginRatio() (*field.MarginRatioField, errors.MessageRejectError) {
	f := &field.MarginRatioField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMarginRatio(f *field.MarginRatioField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Quantity is a required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Quantity() (*field.QuantityField, errors.MessageRejectError) {
	f := &field.QuantityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuantity reads a Quantity from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetQuantity(f *field.QuantityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) QtyType() (*field.QtyTypeField, errors.MessageRejectError) {
	f := &field.QtyTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetQtyType(f *field.QtyTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) LastMkt() (*field.LastMktField, errors.MessageRejectError) {
	f := &field.LastMktField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetLastMkt(f *field.LastMktField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TradeOriginationDate() (*field.TradeOriginationDateField, errors.MessageRejectError) {
	f := &field.TradeOriginationDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTradeOriginationDate(f *field.TradeOriginationDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TradingSessionSubID() (*field.TradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTradingSessionSubID(f *field.TradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AvgPx() (*field.AvgPxField, errors.MessageRejectError) {
	f := &field.AvgPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAvgPx(f *field.AvgPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgParPx is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AvgParPx() (*field.AvgParPxField, errors.MessageRejectError) {
	f := &field.AvgParPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgParPx reads a AvgParPx from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAvgParPx(f *field.AvgParPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Spread() (*field.SpreadField, errors.MessageRejectError) {
	f := &field.SpreadField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSpread(f *field.SpreadField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrencyField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BenchmarkCurveName() (*field.BenchmarkCurveNameField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBenchmarkCurveName(f *field.BenchmarkCurveNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BenchmarkCurvePoint() (*field.BenchmarkCurvePointField, errors.MessageRejectError) {
	f := &field.BenchmarkCurvePointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BenchmarkPrice() (*field.BenchmarkPriceField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBenchmarkPrice(f *field.BenchmarkPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BenchmarkPriceType() (*field.BenchmarkPriceTypeField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBenchmarkPriceType(f *field.BenchmarkPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BenchmarkSecurityID() (*field.BenchmarkSecurityIDField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBenchmarkSecurityID(f *field.BenchmarkSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPxPrecision is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AvgPxPrecision() (*field.AvgPxPrecisionField, errors.MessageRejectError) {
	f := &field.AvgPxPrecisionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPxPrecision reads a AvgPxPrecision from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAvgPxPrecision(f *field.AvgPxPrecisionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TradeDate() (*field.TradeDateField, errors.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTradeDate(f *field.TradeDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SettlType() (*field.SettlTypeField, errors.MessageRejectError) {
	f := &field.SettlTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSettlType(f *field.SettlTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SettlDate() (*field.SettlDateField, errors.MessageRejectError) {
	f := &field.SettlDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSettlDate(f *field.SettlDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BookingType() (*field.BookingTypeField, errors.MessageRejectError) {
	f := &field.BookingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingType reads a BookingType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBookingType(f *field.BookingTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GrossTradeAmt is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) GrossTradeAmt() (*field.GrossTradeAmtField, errors.MessageRejectError) {
	f := &field.GrossTradeAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGrossTradeAmt reads a GrossTradeAmt from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetGrossTradeAmt(f *field.GrossTradeAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Concession is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Concession() (*field.ConcessionField, errors.MessageRejectError) {
	f := &field.ConcessionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConcession reads a Concession from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetConcession(f *field.ConcessionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalTakedown is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TotalTakedown() (*field.TotalTakedownField, errors.MessageRejectError) {
	f := &field.TotalTakedownField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalTakedown reads a TotalTakedown from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTotalTakedown(f *field.TotalTakedownField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NetMoney() (*field.NetMoneyField, errors.MessageRejectError) {
	f := &field.NetMoneyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNetMoney(f *field.NetMoneyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) PositionEffect() (*field.PositionEffectField, errors.MessageRejectError) {
	f := &field.PositionEffectField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetPositionEffect(f *field.PositionEffectField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AutoAcceptIndicator is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AutoAcceptIndicator() (*field.AutoAcceptIndicatorField, errors.MessageRejectError) {
	f := &field.AutoAcceptIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAutoAcceptIndicator reads a AutoAcceptIndicator from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAutoAcceptIndicator(f *field.AutoAcceptIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NumDaysInterest is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NumDaysInterest() (*field.NumDaysInterestField, errors.MessageRejectError) {
	f := &field.NumDaysInterestField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNumDaysInterest reads a NumDaysInterest from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNumDaysInterest(f *field.NumDaysInterestField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestRate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AccruedInterestRate() (*field.AccruedInterestRateField, errors.MessageRejectError) {
	f := &field.AccruedInterestRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestRate reads a AccruedInterestRate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAccruedInterestRate(f *field.AccruedInterestRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestAmt is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AccruedInterestAmt() (*field.AccruedInterestAmtField, errors.MessageRejectError) {
	f := &field.AccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestAmt reads a AccruedInterestAmt from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAccruedInterestAmt(f *field.AccruedInterestAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalAccruedInterestAmt is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TotalAccruedInterestAmt() (*field.TotalAccruedInterestAmtField, errors.MessageRejectError) {
	f := &field.TotalAccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalAccruedInterestAmt reads a TotalAccruedInterestAmt from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTotalAccruedInterestAmt(f *field.TotalAccruedInterestAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAtMaturity is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) InterestAtMaturity() (*field.InterestAtMaturityField, errors.MessageRejectError) {
	f := &field.InterestAtMaturityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAtMaturity reads a InterestAtMaturity from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetInterestAtMaturity(f *field.InterestAtMaturityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndAccruedInterestAmt is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EndAccruedInterestAmt() (*field.EndAccruedInterestAmtField, errors.MessageRejectError) {
	f := &field.EndAccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndAccruedInterestAmt reads a EndAccruedInterestAmt from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEndAccruedInterestAmt(f *field.EndAccruedInterestAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartCash is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StartCash() (*field.StartCashField, errors.MessageRejectError) {
	f := &field.StartCashField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartCash reads a StartCash from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStartCash(f *field.StartCashField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndCash is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EndCash() (*field.EndCashField, errors.MessageRejectError) {
	f := &field.EndCashField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndCash reads a EndCash from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEndCash(f *field.EndCashField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LegalConfirm is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) LegalConfirm() (*field.LegalConfirmField, errors.MessageRejectError) {
	f := &field.LegalConfirmField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLegalConfirm reads a LegalConfirm from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetLegalConfirm(f *field.LegalConfirmField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoStipulations() (*field.NoStipulationsField, errors.MessageRejectError) {
	f := &field.NoStipulationsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoStipulations(f *field.NoStipulationsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) YieldType() (*field.YieldTypeField, errors.MessageRejectError) {
	f := &field.YieldTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetYieldType(f *field.YieldTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Yield() (*field.YieldField, errors.MessageRejectError) {
	f := &field.YieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetYield(f *field.YieldField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) YieldCalcDate() (*field.YieldCalcDateField, errors.MessageRejectError) {
	f := &field.YieldCalcDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetYieldCalcDate(f *field.YieldCalcDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) YieldRedemptionDate() (*field.YieldRedemptionDateField, errors.MessageRejectError) {
	f := &field.YieldRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetYieldRedemptionDate(f *field.YieldRedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) YieldRedemptionPrice() (*field.YieldRedemptionPriceField, errors.MessageRejectError) {
	f := &field.YieldRedemptionPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetYieldRedemptionPrice(f *field.YieldRedemptionPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) YieldRedemptionPriceType() (*field.YieldRedemptionPriceTypeField, errors.MessageRejectError) {
	f := &field.YieldRedemptionPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPosAmt is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoPosAmt() (*field.NoPosAmtField, errors.MessageRejectError) {
	f := &field.NoPosAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPosAmt reads a NoPosAmt from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoPosAmt(f *field.NoPosAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoAllocs is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TotNoAllocs() (*field.TotNoAllocsField, errors.MessageRejectError) {
	f := &field.TotNoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoAllocs reads a TotNoAllocs from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTotNoAllocs(f *field.TotNoAllocsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) LastFragment() (*field.LastFragmentField, errors.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetLastFragment(f *field.LastFragmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoAllocs() (*field.NoAllocsField, errors.MessageRejectError) {
	f := &field.NoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoAllocs(f *field.NoAllocsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPxIndicator is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AvgPxIndicator() (*field.AvgPxIndicatorField, errors.MessageRejectError) {
	f := &field.AvgPxIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPxIndicator reads a AvgPxIndicator from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAvgPxIndicator(f *field.AvgPxIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) ClearingBusinessDate() (*field.ClearingBusinessDateField, errors.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetClearingBusinessDate(f *field.ClearingBusinessDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TrdType() (*field.TrdTypeField, errors.MessageRejectError) {
	f := &field.TrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdType reads a TrdType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTrdType(f *field.TrdTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdSubType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TrdSubType() (*field.TrdSubTypeField, errors.MessageRejectError) {
	f := &field.TrdSubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdSubType reads a TrdSubType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTrdSubType(f *field.TrdSubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CustOrderCapacity() (*field.CustOrderCapacityField, errors.MessageRejectError) {
	f := &field.CustOrderCapacityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCustOrderCapacity(f *field.CustOrderCapacityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeInputSource is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TradeInputSource() (*field.TradeInputSourceField, errors.MessageRejectError) {
	f := &field.TradeInputSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeInputSource reads a TradeInputSource from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTradeInputSource(f *field.TradeInputSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegReportingType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MultiLegReportingType() (*field.MultiLegReportingTypeField, errors.MessageRejectError) {
	f := &field.MultiLegReportingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegReportingType reads a MultiLegReportingType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMultiLegReportingType(f *field.MultiLegReportingTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageEventSource is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MessageEventSource() (*field.MessageEventSourceField, errors.MessageRejectError) {
	f := &field.MessageEventSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMessageEventSource reads a MessageEventSource from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMessageEventSource(f *field.MessageEventSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RndPx is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) RndPx() (*field.RndPxField, errors.MessageRejectError) {
	f := &field.RndPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRndPx reads a RndPx from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetRndPx(f *field.RndPxField) errors.MessageRejectError {
	return m.Body.Get(f)
}
