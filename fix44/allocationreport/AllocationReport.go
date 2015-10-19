//Package allocationreport msg type = AS.
package allocationreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a AllocationReport wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//AllocReportID is a required field for AllocationReport.
func (m Message) AllocReportID() (*field.AllocReportIDField, quickfix.MessageRejectError) {
	f := &field.AllocReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocReportID reads a AllocReportID from AllocationReport.
func (m Message) GetAllocReportID(f *field.AllocReportIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a non-required field for AllocationReport.
func (m Message) AllocID() (*field.AllocIDField, quickfix.MessageRejectError) {
	f := &field.AllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from AllocationReport.
func (m Message) GetAllocID(f *field.AllocIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocTransType is a required field for AllocationReport.
func (m Message) AllocTransType() (*field.AllocTransTypeField, quickfix.MessageRejectError) {
	f := &field.AllocTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocTransType reads a AllocTransType from AllocationReport.
func (m Message) GetAllocTransType(f *field.AllocTransTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocReportRefID is a non-required field for AllocationReport.
func (m Message) AllocReportRefID() (*field.AllocReportRefIDField, quickfix.MessageRejectError) {
	f := &field.AllocReportRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocReportRefID reads a AllocReportRefID from AllocationReport.
func (m Message) GetAllocReportRefID(f *field.AllocReportRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocCancReplaceReason is a non-required field for AllocationReport.
func (m Message) AllocCancReplaceReason() (*field.AllocCancReplaceReasonField, quickfix.MessageRejectError) {
	f := &field.AllocCancReplaceReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocCancReplaceReason reads a AllocCancReplaceReason from AllocationReport.
func (m Message) GetAllocCancReplaceReason(f *field.AllocCancReplaceReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryAllocID is a non-required field for AllocationReport.
func (m Message) SecondaryAllocID() (*field.SecondaryAllocIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryAllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryAllocID reads a SecondaryAllocID from AllocationReport.
func (m Message) GetSecondaryAllocID(f *field.SecondaryAllocIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocReportType is a required field for AllocationReport.
func (m Message) AllocReportType() (*field.AllocReportTypeField, quickfix.MessageRejectError) {
	f := &field.AllocReportTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocReportType reads a AllocReportType from AllocationReport.
func (m Message) GetAllocReportType(f *field.AllocReportTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocStatus is a required field for AllocationReport.
func (m Message) AllocStatus() (*field.AllocStatusField, quickfix.MessageRejectError) {
	f := &field.AllocStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocStatus reads a AllocStatus from AllocationReport.
func (m Message) GetAllocStatus(f *field.AllocStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocRejCode is a non-required field for AllocationReport.
func (m Message) AllocRejCode() (*field.AllocRejCodeField, quickfix.MessageRejectError) {
	f := &field.AllocRejCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocRejCode reads a AllocRejCode from AllocationReport.
func (m Message) GetAllocRejCode(f *field.AllocRejCodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RefAllocID is a non-required field for AllocationReport.
func (m Message) RefAllocID() (*field.RefAllocIDField, quickfix.MessageRejectError) {
	f := &field.RefAllocIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRefAllocID reads a RefAllocID from AllocationReport.
func (m Message) GetRefAllocID(f *field.RefAllocIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocIntermedReqType is a non-required field for AllocationReport.
func (m Message) AllocIntermedReqType() (*field.AllocIntermedReqTypeField, quickfix.MessageRejectError) {
	f := &field.AllocIntermedReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocIntermedReqType reads a AllocIntermedReqType from AllocationReport.
func (m Message) GetAllocIntermedReqType(f *field.AllocIntermedReqTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocLinkID is a non-required field for AllocationReport.
func (m Message) AllocLinkID() (*field.AllocLinkIDField, quickfix.MessageRejectError) {
	f := &field.AllocLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocLinkID reads a AllocLinkID from AllocationReport.
func (m Message) GetAllocLinkID(f *field.AllocLinkIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocLinkType is a non-required field for AllocationReport.
func (m Message) AllocLinkType() (*field.AllocLinkTypeField, quickfix.MessageRejectError) {
	f := &field.AllocLinkTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocLinkType reads a AllocLinkType from AllocationReport.
func (m Message) GetAllocLinkType(f *field.AllocLinkTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BookingRefID is a non-required field for AllocationReport.
func (m Message) BookingRefID() (*field.BookingRefIDField, quickfix.MessageRejectError) {
	f := &field.BookingRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingRefID reads a BookingRefID from AllocationReport.
func (m Message) GetBookingRefID(f *field.BookingRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AllocNoOrdersType is a required field for AllocationReport.
func (m Message) AllocNoOrdersType() (*field.AllocNoOrdersTypeField, quickfix.MessageRejectError) {
	f := &field.AllocNoOrdersTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocNoOrdersType reads a AllocNoOrdersType from AllocationReport.
func (m Message) GetAllocNoOrdersType(f *field.AllocNoOrdersTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a non-required field for AllocationReport.
func (m Message) NoOrders() (*field.NoOrdersField, quickfix.MessageRejectError) {
	f := &field.NoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from AllocationReport.
func (m Message) GetNoOrders(f *field.NoOrdersField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecs is a non-required field for AllocationReport.
func (m Message) NoExecs() (*field.NoExecsField, quickfix.MessageRejectError) {
	f := &field.NoExecsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecs reads a NoExecs from AllocationReport.
func (m Message) GetNoExecs(f *field.NoExecsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PreviouslyReported is a non-required field for AllocationReport.
func (m Message) PreviouslyReported() (*field.PreviouslyReportedField, quickfix.MessageRejectError) {
	f := &field.PreviouslyReportedField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreviouslyReported reads a PreviouslyReported from AllocationReport.
func (m Message) GetPreviouslyReported(f *field.PreviouslyReportedField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ReversalIndicator is a non-required field for AllocationReport.
func (m Message) ReversalIndicator() (*field.ReversalIndicatorField, quickfix.MessageRejectError) {
	f := &field.ReversalIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetReversalIndicator reads a ReversalIndicator from AllocationReport.
func (m Message) GetReversalIndicator(f *field.ReversalIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MatchType is a non-required field for AllocationReport.
func (m Message) MatchType() (*field.MatchTypeField, quickfix.MessageRejectError) {
	f := &field.MatchTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchType reads a MatchType from AllocationReport.
func (m Message) GetMatchType(f *field.MatchTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for AllocationReport.
func (m Message) Side() (*field.SideField, quickfix.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from AllocationReport.
func (m Message) GetSide(f *field.SideField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for AllocationReport.
func (m Message) Symbol() (*field.SymbolField, quickfix.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from AllocationReport.
func (m Message) GetSymbol(f *field.SymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for AllocationReport.
func (m Message) SymbolSfx() (*field.SymbolSfxField, quickfix.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from AllocationReport.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for AllocationReport.
func (m Message) SecurityID() (*field.SecurityIDField, quickfix.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from AllocationReport.
func (m Message) GetSecurityID(f *field.SecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for AllocationReport.
func (m Message) SecurityIDSource() (*field.SecurityIDSourceField, quickfix.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from AllocationReport.
func (m Message) GetSecurityIDSource(f *field.SecurityIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for AllocationReport.
func (m Message) NoSecurityAltID() (*field.NoSecurityAltIDField, quickfix.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from AllocationReport.
func (m Message) GetNoSecurityAltID(f *field.NoSecurityAltIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for AllocationReport.
func (m Message) Product() (*field.ProductField, quickfix.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from AllocationReport.
func (m Message) GetProduct(f *field.ProductField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for AllocationReport.
func (m Message) CFICode() (*field.CFICodeField, quickfix.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from AllocationReport.
func (m Message) GetCFICode(f *field.CFICodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for AllocationReport.
func (m Message) SecurityType() (*field.SecurityTypeField, quickfix.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from AllocationReport.
func (m Message) GetSecurityType(f *field.SecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for AllocationReport.
func (m Message) SecuritySubType() (*field.SecuritySubTypeField, quickfix.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from AllocationReport.
func (m Message) GetSecuritySubType(f *field.SecuritySubTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for AllocationReport.
func (m Message) MaturityMonthYear() (*field.MaturityMonthYearField, quickfix.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from AllocationReport.
func (m Message) GetMaturityMonthYear(f *field.MaturityMonthYearField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for AllocationReport.
func (m Message) MaturityDate() (*field.MaturityDateField, quickfix.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from AllocationReport.
func (m Message) GetMaturityDate(f *field.MaturityDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for AllocationReport.
func (m Message) CouponPaymentDate() (*field.CouponPaymentDateField, quickfix.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from AllocationReport.
func (m Message) GetCouponPaymentDate(f *field.CouponPaymentDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for AllocationReport.
func (m Message) IssueDate() (*field.IssueDateField, quickfix.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from AllocationReport.
func (m Message) GetIssueDate(f *field.IssueDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for AllocationReport.
func (m Message) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, quickfix.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from AllocationReport.
func (m Message) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for AllocationReport.
func (m Message) RepurchaseTerm() (*field.RepurchaseTermField, quickfix.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from AllocationReport.
func (m Message) GetRepurchaseTerm(f *field.RepurchaseTermField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for AllocationReport.
func (m Message) RepurchaseRate() (*field.RepurchaseRateField, quickfix.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from AllocationReport.
func (m Message) GetRepurchaseRate(f *field.RepurchaseRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for AllocationReport.
func (m Message) Factor() (*field.FactorField, quickfix.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from AllocationReport.
func (m Message) GetFactor(f *field.FactorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for AllocationReport.
func (m Message) CreditRating() (*field.CreditRatingField, quickfix.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from AllocationReport.
func (m Message) GetCreditRating(f *field.CreditRatingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for AllocationReport.
func (m Message) InstrRegistry() (*field.InstrRegistryField, quickfix.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from AllocationReport.
func (m Message) GetInstrRegistry(f *field.InstrRegistryField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for AllocationReport.
func (m Message) CountryOfIssue() (*field.CountryOfIssueField, quickfix.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from AllocationReport.
func (m Message) GetCountryOfIssue(f *field.CountryOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for AllocationReport.
func (m Message) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, quickfix.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from AllocationReport.
func (m Message) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for AllocationReport.
func (m Message) LocaleOfIssue() (*field.LocaleOfIssueField, quickfix.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from AllocationReport.
func (m Message) GetLocaleOfIssue(f *field.LocaleOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for AllocationReport.
func (m Message) RedemptionDate() (*field.RedemptionDateField, quickfix.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from AllocationReport.
func (m Message) GetRedemptionDate(f *field.RedemptionDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for AllocationReport.
func (m Message) StrikePrice() (*field.StrikePriceField, quickfix.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from AllocationReport.
func (m Message) GetStrikePrice(f *field.StrikePriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for AllocationReport.
func (m Message) StrikeCurrency() (*field.StrikeCurrencyField, quickfix.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from AllocationReport.
func (m Message) GetStrikeCurrency(f *field.StrikeCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for AllocationReport.
func (m Message) OptAttribute() (*field.OptAttributeField, quickfix.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from AllocationReport.
func (m Message) GetOptAttribute(f *field.OptAttributeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for AllocationReport.
func (m Message) ContractMultiplier() (*field.ContractMultiplierField, quickfix.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from AllocationReport.
func (m Message) GetContractMultiplier(f *field.ContractMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for AllocationReport.
func (m Message) CouponRate() (*field.CouponRateField, quickfix.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from AllocationReport.
func (m Message) GetCouponRate(f *field.CouponRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for AllocationReport.
func (m Message) SecurityExchange() (*field.SecurityExchangeField, quickfix.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from AllocationReport.
func (m Message) GetSecurityExchange(f *field.SecurityExchangeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for AllocationReport.
func (m Message) Issuer() (*field.IssuerField, quickfix.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from AllocationReport.
func (m Message) GetIssuer(f *field.IssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for AllocationReport.
func (m Message) EncodedIssuerLen() (*field.EncodedIssuerLenField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from AllocationReport.
func (m Message) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for AllocationReport.
func (m Message) EncodedIssuer() (*field.EncodedIssuerField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from AllocationReport.
func (m Message) GetEncodedIssuer(f *field.EncodedIssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for AllocationReport.
func (m Message) SecurityDesc() (*field.SecurityDescField, quickfix.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from AllocationReport.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for AllocationReport.
func (m Message) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from AllocationReport.
func (m Message) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for AllocationReport.
func (m Message) EncodedSecurityDesc() (*field.EncodedSecurityDescField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from AllocationReport.
func (m Message) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for AllocationReport.
func (m Message) Pool() (*field.PoolField, quickfix.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from AllocationReport.
func (m Message) GetPool(f *field.PoolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for AllocationReport.
func (m Message) ContractSettlMonth() (*field.ContractSettlMonthField, quickfix.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from AllocationReport.
func (m Message) GetContractSettlMonth(f *field.ContractSettlMonthField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for AllocationReport.
func (m Message) CPProgram() (*field.CPProgramField, quickfix.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from AllocationReport.
func (m Message) GetCPProgram(f *field.CPProgramField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for AllocationReport.
func (m Message) CPRegType() (*field.CPRegTypeField, quickfix.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from AllocationReport.
func (m Message) GetCPRegType(f *field.CPRegTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for AllocationReport.
func (m Message) NoEvents() (*field.NoEventsField, quickfix.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from AllocationReport.
func (m Message) GetNoEvents(f *field.NoEventsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for AllocationReport.
func (m Message) DatedDate() (*field.DatedDateField, quickfix.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from AllocationReport.
func (m Message) GetDatedDate(f *field.DatedDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for AllocationReport.
func (m Message) InterestAccrualDate() (*field.InterestAccrualDateField, quickfix.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from AllocationReport.
func (m Message) GetInterestAccrualDate(f *field.InterestAccrualDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryForm is a non-required field for AllocationReport.
func (m Message) DeliveryForm() (*field.DeliveryFormField, quickfix.MessageRejectError) {
	f := &field.DeliveryFormField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryForm reads a DeliveryForm from AllocationReport.
func (m Message) GetDeliveryForm(f *field.DeliveryFormField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PctAtRisk is a non-required field for AllocationReport.
func (m Message) PctAtRisk() (*field.PctAtRiskField, quickfix.MessageRejectError) {
	f := &field.PctAtRiskField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPctAtRisk reads a PctAtRisk from AllocationReport.
func (m Message) GetPctAtRisk(f *field.PctAtRiskField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrAttrib is a non-required field for AllocationReport.
func (m Message) NoInstrAttrib() (*field.NoInstrAttribField, quickfix.MessageRejectError) {
	f := &field.NoInstrAttribField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrAttrib reads a NoInstrAttrib from AllocationReport.
func (m Message) GetNoInstrAttrib(f *field.NoInstrAttribField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for AllocationReport.
func (m Message) AgreementDesc() (*field.AgreementDescField, quickfix.MessageRejectError) {
	f := &field.AgreementDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from AllocationReport.
func (m Message) GetAgreementDesc(f *field.AgreementDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for AllocationReport.
func (m Message) AgreementID() (*field.AgreementIDField, quickfix.MessageRejectError) {
	f := &field.AgreementIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from AllocationReport.
func (m Message) GetAgreementID(f *field.AgreementIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for AllocationReport.
func (m Message) AgreementDate() (*field.AgreementDateField, quickfix.MessageRejectError) {
	f := &field.AgreementDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from AllocationReport.
func (m Message) GetAgreementDate(f *field.AgreementDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for AllocationReport.
func (m Message) AgreementCurrency() (*field.AgreementCurrencyField, quickfix.MessageRejectError) {
	f := &field.AgreementCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from AllocationReport.
func (m Message) GetAgreementCurrency(f *field.AgreementCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for AllocationReport.
func (m Message) TerminationType() (*field.TerminationTypeField, quickfix.MessageRejectError) {
	f := &field.TerminationTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from AllocationReport.
func (m Message) GetTerminationType(f *field.TerminationTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for AllocationReport.
func (m Message) StartDate() (*field.StartDateField, quickfix.MessageRejectError) {
	f := &field.StartDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from AllocationReport.
func (m Message) GetStartDate(f *field.StartDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for AllocationReport.
func (m Message) EndDate() (*field.EndDateField, quickfix.MessageRejectError) {
	f := &field.EndDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from AllocationReport.
func (m Message) GetEndDate(f *field.EndDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for AllocationReport.
func (m Message) DeliveryType() (*field.DeliveryTypeField, quickfix.MessageRejectError) {
	f := &field.DeliveryTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from AllocationReport.
func (m Message) GetDeliveryType(f *field.DeliveryTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for AllocationReport.
func (m Message) MarginRatio() (*field.MarginRatioField, quickfix.MessageRejectError) {
	f := &field.MarginRatioField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from AllocationReport.
func (m Message) GetMarginRatio(f *field.MarginRatioField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for AllocationReport.
func (m Message) NoUnderlyings() (*field.NoUnderlyingsField, quickfix.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from AllocationReport.
func (m Message) GetNoUnderlyings(f *field.NoUnderlyingsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for AllocationReport.
func (m Message) NoLegs() (*field.NoLegsField, quickfix.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from AllocationReport.
func (m Message) GetNoLegs(f *field.NoLegsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Quantity is a required field for AllocationReport.
func (m Message) Quantity() (*field.QuantityField, quickfix.MessageRejectError) {
	f := &field.QuantityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuantity reads a Quantity from AllocationReport.
func (m Message) GetQuantity(f *field.QuantityField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for AllocationReport.
func (m Message) QtyType() (*field.QtyTypeField, quickfix.MessageRejectError) {
	f := &field.QtyTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from AllocationReport.
func (m Message) GetQtyType(f *field.QtyTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for AllocationReport.
func (m Message) LastMkt() (*field.LastMktField, quickfix.MessageRejectError) {
	f := &field.LastMktField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from AllocationReport.
func (m Message) GetLastMkt(f *field.LastMktField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for AllocationReport.
func (m Message) TradeOriginationDate() (*field.TradeOriginationDateField, quickfix.MessageRejectError) {
	f := &field.TradeOriginationDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from AllocationReport.
func (m Message) GetTradeOriginationDate(f *field.TradeOriginationDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for AllocationReport.
func (m Message) TradingSessionID() (*field.TradingSessionIDField, quickfix.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from AllocationReport.
func (m Message) GetTradingSessionID(f *field.TradingSessionIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for AllocationReport.
func (m Message) TradingSessionSubID() (*field.TradingSessionSubIDField, quickfix.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from AllocationReport.
func (m Message) GetTradingSessionSubID(f *field.TradingSessionSubIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for AllocationReport.
func (m Message) PriceType() (*field.PriceTypeField, quickfix.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from AllocationReport.
func (m Message) GetPriceType(f *field.PriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a required field for AllocationReport.
func (m Message) AvgPx() (*field.AvgPxField, quickfix.MessageRejectError) {
	f := &field.AvgPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from AllocationReport.
func (m Message) GetAvgPx(f *field.AvgPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AvgParPx is a non-required field for AllocationReport.
func (m Message) AvgParPx() (*field.AvgParPxField, quickfix.MessageRejectError) {
	f := &field.AvgParPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgParPx reads a AvgParPx from AllocationReport.
func (m Message) GetAvgParPx(f *field.AvgParPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for AllocationReport.
func (m Message) Spread() (*field.SpreadField, quickfix.MessageRejectError) {
	f := &field.SpreadField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from AllocationReport.
func (m Message) GetSpread(f *field.SpreadField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for AllocationReport.
func (m Message) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrencyField, quickfix.MessageRejectError) {
	f := &field.BenchmarkCurveCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from AllocationReport.
func (m Message) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for AllocationReport.
func (m Message) BenchmarkCurveName() (*field.BenchmarkCurveNameField, quickfix.MessageRejectError) {
	f := &field.BenchmarkCurveNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from AllocationReport.
func (m Message) GetBenchmarkCurveName(f *field.BenchmarkCurveNameField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for AllocationReport.
func (m Message) BenchmarkCurvePoint() (*field.BenchmarkCurvePointField, quickfix.MessageRejectError) {
	f := &field.BenchmarkCurvePointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from AllocationReport.
func (m Message) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePointField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for AllocationReport.
func (m Message) BenchmarkPrice() (*field.BenchmarkPriceField, quickfix.MessageRejectError) {
	f := &field.BenchmarkPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from AllocationReport.
func (m Message) GetBenchmarkPrice(f *field.BenchmarkPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for AllocationReport.
func (m Message) BenchmarkPriceType() (*field.BenchmarkPriceTypeField, quickfix.MessageRejectError) {
	f := &field.BenchmarkPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from AllocationReport.
func (m Message) GetBenchmarkPriceType(f *field.BenchmarkPriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for AllocationReport.
func (m Message) BenchmarkSecurityID() (*field.BenchmarkSecurityIDField, quickfix.MessageRejectError) {
	f := &field.BenchmarkSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from AllocationReport.
func (m Message) GetBenchmarkSecurityID(f *field.BenchmarkSecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for AllocationReport.
func (m Message) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSourceField, quickfix.MessageRejectError) {
	f := &field.BenchmarkSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from AllocationReport.
func (m Message) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for AllocationReport.
func (m Message) Currency() (*field.CurrencyField, quickfix.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from AllocationReport.
func (m Message) GetCurrency(f *field.CurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPxPrecision is a non-required field for AllocationReport.
func (m Message) AvgPxPrecision() (*field.AvgPxPrecisionField, quickfix.MessageRejectError) {
	f := &field.AvgPxPrecisionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPxPrecision reads a AvgPxPrecision from AllocationReport.
func (m Message) GetAvgPxPrecision(f *field.AvgPxPrecisionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for AllocationReport.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AllocationReport.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for AllocationReport.
func (m Message) TradeDate() (*field.TradeDateField, quickfix.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from AllocationReport.
func (m Message) GetTradeDate(f *field.TradeDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for AllocationReport.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from AllocationReport.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for AllocationReport.
func (m Message) SettlType() (*field.SettlTypeField, quickfix.MessageRejectError) {
	f := &field.SettlTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from AllocationReport.
func (m Message) GetSettlType(f *field.SettlTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for AllocationReport.
func (m Message) SettlDate() (*field.SettlDateField, quickfix.MessageRejectError) {
	f := &field.SettlDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from AllocationReport.
func (m Message) GetSettlDate(f *field.SettlDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//BookingType is a non-required field for AllocationReport.
func (m Message) BookingType() (*field.BookingTypeField, quickfix.MessageRejectError) {
	f := &field.BookingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBookingType reads a BookingType from AllocationReport.
func (m Message) GetBookingType(f *field.BookingTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//GrossTradeAmt is a non-required field for AllocationReport.
func (m Message) GrossTradeAmt() (*field.GrossTradeAmtField, quickfix.MessageRejectError) {
	f := &field.GrossTradeAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGrossTradeAmt reads a GrossTradeAmt from AllocationReport.
func (m Message) GetGrossTradeAmt(f *field.GrossTradeAmtField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Concession is a non-required field for AllocationReport.
func (m Message) Concession() (*field.ConcessionField, quickfix.MessageRejectError) {
	f := &field.ConcessionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetConcession reads a Concession from AllocationReport.
func (m Message) GetConcession(f *field.ConcessionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TotalTakedown is a non-required field for AllocationReport.
func (m Message) TotalTakedown() (*field.TotalTakedownField, quickfix.MessageRejectError) {
	f := &field.TotalTakedownField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalTakedown reads a TotalTakedown from AllocationReport.
func (m Message) GetTotalTakedown(f *field.TotalTakedownField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a non-required field for AllocationReport.
func (m Message) NetMoney() (*field.NetMoneyField, quickfix.MessageRejectError) {
	f := &field.NetMoneyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from AllocationReport.
func (m Message) GetNetMoney(f *field.NetMoneyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for AllocationReport.
func (m Message) PositionEffect() (*field.PositionEffectField, quickfix.MessageRejectError) {
	f := &field.PositionEffectField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from AllocationReport.
func (m Message) GetPositionEffect(f *field.PositionEffectField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AutoAcceptIndicator is a non-required field for AllocationReport.
func (m Message) AutoAcceptIndicator() (*field.AutoAcceptIndicatorField, quickfix.MessageRejectError) {
	f := &field.AutoAcceptIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAutoAcceptIndicator reads a AutoAcceptIndicator from AllocationReport.
func (m Message) GetAutoAcceptIndicator(f *field.AutoAcceptIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for AllocationReport.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from AllocationReport.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for AllocationReport.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from AllocationReport.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for AllocationReport.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from AllocationReport.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NumDaysInterest is a non-required field for AllocationReport.
func (m Message) NumDaysInterest() (*field.NumDaysInterestField, quickfix.MessageRejectError) {
	f := &field.NumDaysInterestField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNumDaysInterest reads a NumDaysInterest from AllocationReport.
func (m Message) GetNumDaysInterest(f *field.NumDaysInterestField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestRate is a non-required field for AllocationReport.
func (m Message) AccruedInterestRate() (*field.AccruedInterestRateField, quickfix.MessageRejectError) {
	f := &field.AccruedInterestRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestRate reads a AccruedInterestRate from AllocationReport.
func (m Message) GetAccruedInterestRate(f *field.AccruedInterestRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestAmt is a non-required field for AllocationReport.
func (m Message) AccruedInterestAmt() (*field.AccruedInterestAmtField, quickfix.MessageRejectError) {
	f := &field.AccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestAmt reads a AccruedInterestAmt from AllocationReport.
func (m Message) GetAccruedInterestAmt(f *field.AccruedInterestAmtField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TotalAccruedInterestAmt is a non-required field for AllocationReport.
func (m Message) TotalAccruedInterestAmt() (*field.TotalAccruedInterestAmtField, quickfix.MessageRejectError) {
	f := &field.TotalAccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalAccruedInterestAmt reads a TotalAccruedInterestAmt from AllocationReport.
func (m Message) GetTotalAccruedInterestAmt(f *field.TotalAccruedInterestAmtField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAtMaturity is a non-required field for AllocationReport.
func (m Message) InterestAtMaturity() (*field.InterestAtMaturityField, quickfix.MessageRejectError) {
	f := &field.InterestAtMaturityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAtMaturity reads a InterestAtMaturity from AllocationReport.
func (m Message) GetInterestAtMaturity(f *field.InterestAtMaturityField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EndAccruedInterestAmt is a non-required field for AllocationReport.
func (m Message) EndAccruedInterestAmt() (*field.EndAccruedInterestAmtField, quickfix.MessageRejectError) {
	f := &field.EndAccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndAccruedInterestAmt reads a EndAccruedInterestAmt from AllocationReport.
func (m Message) GetEndAccruedInterestAmt(f *field.EndAccruedInterestAmtField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StartCash is a non-required field for AllocationReport.
func (m Message) StartCash() (*field.StartCashField, quickfix.MessageRejectError) {
	f := &field.StartCashField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartCash reads a StartCash from AllocationReport.
func (m Message) GetStartCash(f *field.StartCashField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EndCash is a non-required field for AllocationReport.
func (m Message) EndCash() (*field.EndCashField, quickfix.MessageRejectError) {
	f := &field.EndCashField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndCash reads a EndCash from AllocationReport.
func (m Message) GetEndCash(f *field.EndCashField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LegalConfirm is a non-required field for AllocationReport.
func (m Message) LegalConfirm() (*field.LegalConfirmField, quickfix.MessageRejectError) {
	f := &field.LegalConfirmField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLegalConfirm reads a LegalConfirm from AllocationReport.
func (m Message) GetLegalConfirm(f *field.LegalConfirmField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for AllocationReport.
func (m Message) NoStipulations() (*field.NoStipulationsField, quickfix.MessageRejectError) {
	f := &field.NoStipulationsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from AllocationReport.
func (m Message) GetNoStipulations(f *field.NoStipulationsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for AllocationReport.
func (m Message) YieldType() (*field.YieldTypeField, quickfix.MessageRejectError) {
	f := &field.YieldTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from AllocationReport.
func (m Message) GetYieldType(f *field.YieldTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for AllocationReport.
func (m Message) Yield() (*field.YieldField, quickfix.MessageRejectError) {
	f := &field.YieldField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from AllocationReport.
func (m Message) GetYield(f *field.YieldField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for AllocationReport.
func (m Message) YieldCalcDate() (*field.YieldCalcDateField, quickfix.MessageRejectError) {
	f := &field.YieldCalcDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from AllocationReport.
func (m Message) GetYieldCalcDate(f *field.YieldCalcDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for AllocationReport.
func (m Message) YieldRedemptionDate() (*field.YieldRedemptionDateField, quickfix.MessageRejectError) {
	f := &field.YieldRedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from AllocationReport.
func (m Message) GetYieldRedemptionDate(f *field.YieldRedemptionDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for AllocationReport.
func (m Message) YieldRedemptionPrice() (*field.YieldRedemptionPriceField, quickfix.MessageRejectError) {
	f := &field.YieldRedemptionPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from AllocationReport.
func (m Message) GetYieldRedemptionPrice(f *field.YieldRedemptionPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for AllocationReport.
func (m Message) YieldRedemptionPriceType() (*field.YieldRedemptionPriceTypeField, quickfix.MessageRejectError) {
	f := &field.YieldRedemptionPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from AllocationReport.
func (m Message) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoAllocs is a non-required field for AllocationReport.
func (m Message) TotNoAllocs() (*field.TotNoAllocsField, quickfix.MessageRejectError) {
	f := &field.TotNoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoAllocs reads a TotNoAllocs from AllocationReport.
func (m Message) GetTotNoAllocs(f *field.TotNoAllocsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for AllocationReport.
func (m Message) LastFragment() (*field.LastFragmentField, quickfix.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from AllocationReport.
func (m Message) GetLastFragment(f *field.LastFragmentField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for AllocationReport.
func (m Message) NoAllocs() (*field.NoAllocsField, quickfix.MessageRejectError) {
	f := &field.NoAllocsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from AllocationReport.
func (m Message) GetNoAllocs(f *field.NoAllocsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds AllocationReport messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for AllocationReport.
func Builder(
	allocreportid *field.AllocReportIDField,
	alloctranstype *field.AllocTransTypeField,
	allocreporttype *field.AllocReportTypeField,
	allocstatus *field.AllocStatusField,
	allocnoorderstype *field.AllocNoOrdersTypeField,
	side *field.SideField,
	quantity *field.QuantityField,
	avgpx *field.AvgPxField,
	tradedate *field.TradeDateField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header.Set(field.NewMsgType("AS"))
	builder.Body.Set(allocreportid)
	builder.Body.Set(alloctranstype)
	builder.Body.Set(allocreporttype)
	builder.Body.Set(allocstatus)
	builder.Body.Set(allocnoorderstype)
	builder.Body.Set(side)
	builder.Body.Set(quantity)
	builder.Body.Set(avgpx)
	builder.Body.Set(tradedate)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX44, "AS", r
}
