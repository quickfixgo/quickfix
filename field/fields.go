package field

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/tag"
)

//AccountField is a STRING field
type AccountField struct{ quickfix.StringValue }

//Tag returns tag.Account (1)
func (f AccountField) Tag() quickfix.Tag { return tag.Account }

//NewAccount returns a new AccountField initialized with val
func NewAccount(val string) *AccountField {
	field := &AccountField{}
	field.Value = val
	return field
}

//AccountTypeField is a INT field
type AccountTypeField struct{ quickfix.IntValue }

//Tag returns tag.AccountType (581)
func (f AccountTypeField) Tag() quickfix.Tag { return tag.AccountType }

//NewAccountType returns a new AccountTypeField initialized with val
func NewAccountType(val int) *AccountTypeField {
	field := &AccountTypeField{}
	field.Value = val
	return field
}

//AccruedInterestAmtField is a AMT field
type AccruedInterestAmtField struct{ quickfix.AmtValue }

//Tag returns tag.AccruedInterestAmt (159)
func (f AccruedInterestAmtField) Tag() quickfix.Tag { return tag.AccruedInterestAmt }

//NewAccruedInterestAmt returns a new AccruedInterestAmtField initialized with val
func NewAccruedInterestAmt(val float64) *AccruedInterestAmtField {
	field := &AccruedInterestAmtField{}
	field.Value = val
	return field
}

//AccruedInterestRateField is a PERCENTAGE field
type AccruedInterestRateField struct{ quickfix.PercentageValue }

//Tag returns tag.AccruedInterestRate (158)
func (f AccruedInterestRateField) Tag() quickfix.Tag { return tag.AccruedInterestRate }

//NewAccruedInterestRate returns a new AccruedInterestRateField initialized with val
func NewAccruedInterestRate(val float64) *AccruedInterestRateField {
	field := &AccruedInterestRateField{}
	field.Value = val
	return field
}

//AcctIDSourceField is a INT field
type AcctIDSourceField struct{ quickfix.IntValue }

//Tag returns tag.AcctIDSource (660)
func (f AcctIDSourceField) Tag() quickfix.Tag { return tag.AcctIDSource }

//NewAcctIDSource returns a new AcctIDSourceField initialized with val
func NewAcctIDSource(val int) *AcctIDSourceField {
	field := &AcctIDSourceField{}
	field.Value = val
	return field
}

//AdjustmentField is a INT field
type AdjustmentField struct{ quickfix.IntValue }

//Tag returns tag.Adjustment (334)
func (f AdjustmentField) Tag() quickfix.Tag { return tag.Adjustment }

//NewAdjustment returns a new AdjustmentField initialized with val
func NewAdjustment(val int) *AdjustmentField {
	field := &AdjustmentField{}
	field.Value = val
	return field
}

//AdjustmentTypeField is a INT field
type AdjustmentTypeField struct{ quickfix.IntValue }

//Tag returns tag.AdjustmentType (718)
func (f AdjustmentTypeField) Tag() quickfix.Tag { return tag.AdjustmentType }

//NewAdjustmentType returns a new AdjustmentTypeField initialized with val
func NewAdjustmentType(val int) *AdjustmentTypeField {
	field := &AdjustmentTypeField{}
	field.Value = val
	return field
}

//AdvIdField is a STRING field
type AdvIdField struct{ quickfix.StringValue }

//Tag returns tag.AdvId (2)
func (f AdvIdField) Tag() quickfix.Tag { return tag.AdvId }

//NewAdvId returns a new AdvIdField initialized with val
func NewAdvId(val string) *AdvIdField {
	field := &AdvIdField{}
	field.Value = val
	return field
}

//AdvRefIDField is a STRING field
type AdvRefIDField struct{ quickfix.StringValue }

//Tag returns tag.AdvRefID (3)
func (f AdvRefIDField) Tag() quickfix.Tag { return tag.AdvRefID }

//NewAdvRefID returns a new AdvRefIDField initialized with val
func NewAdvRefID(val string) *AdvRefIDField {
	field := &AdvRefIDField{}
	field.Value = val
	return field
}

//AdvSideField is a CHAR field
type AdvSideField struct{ quickfix.CharValue }

//Tag returns tag.AdvSide (4)
func (f AdvSideField) Tag() quickfix.Tag { return tag.AdvSide }

//NewAdvSide returns a new AdvSideField initialized with val
func NewAdvSide(val string) *AdvSideField {
	field := &AdvSideField{}
	field.Value = val
	return field
}

//AdvTransTypeField is a STRING field
type AdvTransTypeField struct{ quickfix.StringValue }

//Tag returns tag.AdvTransType (5)
func (f AdvTransTypeField) Tag() quickfix.Tag { return tag.AdvTransType }

//NewAdvTransType returns a new AdvTransTypeField initialized with val
func NewAdvTransType(val string) *AdvTransTypeField {
	field := &AdvTransTypeField{}
	field.Value = val
	return field
}

//AffectedOrderIDField is a STRING field
type AffectedOrderIDField struct{ quickfix.StringValue }

//Tag returns tag.AffectedOrderID (535)
func (f AffectedOrderIDField) Tag() quickfix.Tag { return tag.AffectedOrderID }

//NewAffectedOrderID returns a new AffectedOrderIDField initialized with val
func NewAffectedOrderID(val string) *AffectedOrderIDField {
	field := &AffectedOrderIDField{}
	field.Value = val
	return field
}

//AffectedSecondaryOrderIDField is a STRING field
type AffectedSecondaryOrderIDField struct{ quickfix.StringValue }

//Tag returns tag.AffectedSecondaryOrderID (536)
func (f AffectedSecondaryOrderIDField) Tag() quickfix.Tag { return tag.AffectedSecondaryOrderID }

//NewAffectedSecondaryOrderID returns a new AffectedSecondaryOrderIDField initialized with val
func NewAffectedSecondaryOrderID(val string) *AffectedSecondaryOrderIDField {
	field := &AffectedSecondaryOrderIDField{}
	field.Value = val
	return field
}

//AffirmStatusField is a INT field
type AffirmStatusField struct{ quickfix.IntValue }

//Tag returns tag.AffirmStatus (940)
func (f AffirmStatusField) Tag() quickfix.Tag { return tag.AffirmStatus }

//NewAffirmStatus returns a new AffirmStatusField initialized with val
func NewAffirmStatus(val int) *AffirmStatusField {
	field := &AffirmStatusField{}
	field.Value = val
	return field
}

//AggregatedBookField is a BOOLEAN field
type AggregatedBookField struct{ quickfix.BooleanValue }

//Tag returns tag.AggregatedBook (266)
func (f AggregatedBookField) Tag() quickfix.Tag { return tag.AggregatedBook }

//NewAggregatedBook returns a new AggregatedBookField initialized with val
func NewAggregatedBook(val bool) *AggregatedBookField {
	field := &AggregatedBookField{}
	field.Value = val
	return field
}

//AggressorIndicatorField is a BOOLEAN field
type AggressorIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.AggressorIndicator (1057)
func (f AggressorIndicatorField) Tag() quickfix.Tag { return tag.AggressorIndicator }

//NewAggressorIndicator returns a new AggressorIndicatorField initialized with val
func NewAggressorIndicator(val bool) *AggressorIndicatorField {
	field := &AggressorIndicatorField{}
	field.Value = val
	return field
}

//AgreementCurrencyField is a CURRENCY field
type AgreementCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.AgreementCurrency (918)
func (f AgreementCurrencyField) Tag() quickfix.Tag { return tag.AgreementCurrency }

//NewAgreementCurrency returns a new AgreementCurrencyField initialized with val
func NewAgreementCurrency(val string) *AgreementCurrencyField {
	field := &AgreementCurrencyField{}
	field.Value = val
	return field
}

//AgreementDateField is a LOCALMKTDATE field
type AgreementDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.AgreementDate (915)
func (f AgreementDateField) Tag() quickfix.Tag { return tag.AgreementDate }

//NewAgreementDate returns a new AgreementDateField initialized with val
func NewAgreementDate(val string) *AgreementDateField {
	field := &AgreementDateField{}
	field.Value = val
	return field
}

//AgreementDescField is a STRING field
type AgreementDescField struct{ quickfix.StringValue }

//Tag returns tag.AgreementDesc (913)
func (f AgreementDescField) Tag() quickfix.Tag { return tag.AgreementDesc }

//NewAgreementDesc returns a new AgreementDescField initialized with val
func NewAgreementDesc(val string) *AgreementDescField {
	field := &AgreementDescField{}
	field.Value = val
	return field
}

//AgreementIDField is a STRING field
type AgreementIDField struct{ quickfix.StringValue }

//Tag returns tag.AgreementID (914)
func (f AgreementIDField) Tag() quickfix.Tag { return tag.AgreementID }

//NewAgreementID returns a new AgreementIDField initialized with val
func NewAgreementID(val string) *AgreementIDField {
	field := &AgreementIDField{}
	field.Value = val
	return field
}

//AllocAccountField is a STRING field
type AllocAccountField struct{ quickfix.StringValue }

//Tag returns tag.AllocAccount (79)
func (f AllocAccountField) Tag() quickfix.Tag { return tag.AllocAccount }

//NewAllocAccount returns a new AllocAccountField initialized with val
func NewAllocAccount(val string) *AllocAccountField {
	field := &AllocAccountField{}
	field.Value = val
	return field
}

//AllocAccountTypeField is a INT field
type AllocAccountTypeField struct{ quickfix.IntValue }

//Tag returns tag.AllocAccountType (798)
func (f AllocAccountTypeField) Tag() quickfix.Tag { return tag.AllocAccountType }

//NewAllocAccountType returns a new AllocAccountTypeField initialized with val
func NewAllocAccountType(val int) *AllocAccountTypeField {
	field := &AllocAccountTypeField{}
	field.Value = val
	return field
}

//AllocAccruedInterestAmtField is a AMT field
type AllocAccruedInterestAmtField struct{ quickfix.AmtValue }

//Tag returns tag.AllocAccruedInterestAmt (742)
func (f AllocAccruedInterestAmtField) Tag() quickfix.Tag { return tag.AllocAccruedInterestAmt }

//NewAllocAccruedInterestAmt returns a new AllocAccruedInterestAmtField initialized with val
func NewAllocAccruedInterestAmt(val float64) *AllocAccruedInterestAmtField {
	field := &AllocAccruedInterestAmtField{}
	field.Value = val
	return field
}

//AllocAcctIDSourceField is a INT field
type AllocAcctIDSourceField struct{ quickfix.IntValue }

//Tag returns tag.AllocAcctIDSource (661)
func (f AllocAcctIDSourceField) Tag() quickfix.Tag { return tag.AllocAcctIDSource }

//NewAllocAcctIDSource returns a new AllocAcctIDSourceField initialized with val
func NewAllocAcctIDSource(val int) *AllocAcctIDSourceField {
	field := &AllocAcctIDSourceField{}
	field.Value = val
	return field
}

//AllocAvgPxField is a PRICE field
type AllocAvgPxField struct{ quickfix.PriceValue }

//Tag returns tag.AllocAvgPx (153)
func (f AllocAvgPxField) Tag() quickfix.Tag { return tag.AllocAvgPx }

//NewAllocAvgPx returns a new AllocAvgPxField initialized with val
func NewAllocAvgPx(val float64) *AllocAvgPxField {
	field := &AllocAvgPxField{}
	field.Value = val
	return field
}

//AllocCancReplaceReasonField is a INT field
type AllocCancReplaceReasonField struct{ quickfix.IntValue }

//Tag returns tag.AllocCancReplaceReason (796)
func (f AllocCancReplaceReasonField) Tag() quickfix.Tag { return tag.AllocCancReplaceReason }

//NewAllocCancReplaceReason returns a new AllocCancReplaceReasonField initialized with val
func NewAllocCancReplaceReason(val int) *AllocCancReplaceReasonField {
	field := &AllocCancReplaceReasonField{}
	field.Value = val
	return field
}

//AllocClearingFeeIndicatorField is a STRING field
type AllocClearingFeeIndicatorField struct{ quickfix.StringValue }

//Tag returns tag.AllocClearingFeeIndicator (1136)
func (f AllocClearingFeeIndicatorField) Tag() quickfix.Tag { return tag.AllocClearingFeeIndicator }

//NewAllocClearingFeeIndicator returns a new AllocClearingFeeIndicatorField initialized with val
func NewAllocClearingFeeIndicator(val string) *AllocClearingFeeIndicatorField {
	field := &AllocClearingFeeIndicatorField{}
	field.Value = val
	return field
}

//AllocCustomerCapacityField is a STRING field
type AllocCustomerCapacityField struct{ quickfix.StringValue }

//Tag returns tag.AllocCustomerCapacity (993)
func (f AllocCustomerCapacityField) Tag() quickfix.Tag { return tag.AllocCustomerCapacity }

//NewAllocCustomerCapacity returns a new AllocCustomerCapacityField initialized with val
func NewAllocCustomerCapacity(val string) *AllocCustomerCapacityField {
	field := &AllocCustomerCapacityField{}
	field.Value = val
	return field
}

//AllocHandlInstField is a INT field
type AllocHandlInstField struct{ quickfix.IntValue }

//Tag returns tag.AllocHandlInst (209)
func (f AllocHandlInstField) Tag() quickfix.Tag { return tag.AllocHandlInst }

//NewAllocHandlInst returns a new AllocHandlInstField initialized with val
func NewAllocHandlInst(val int) *AllocHandlInstField {
	field := &AllocHandlInstField{}
	field.Value = val
	return field
}

//AllocIDField is a STRING field
type AllocIDField struct{ quickfix.StringValue }

//Tag returns tag.AllocID (70)
func (f AllocIDField) Tag() quickfix.Tag { return tag.AllocID }

//NewAllocID returns a new AllocIDField initialized with val
func NewAllocID(val string) *AllocIDField {
	field := &AllocIDField{}
	field.Value = val
	return field
}

//AllocInterestAtMaturityField is a AMT field
type AllocInterestAtMaturityField struct{ quickfix.AmtValue }

//Tag returns tag.AllocInterestAtMaturity (741)
func (f AllocInterestAtMaturityField) Tag() quickfix.Tag { return tag.AllocInterestAtMaturity }

//NewAllocInterestAtMaturity returns a new AllocInterestAtMaturityField initialized with val
func NewAllocInterestAtMaturity(val float64) *AllocInterestAtMaturityField {
	field := &AllocInterestAtMaturityField{}
	field.Value = val
	return field
}

//AllocIntermedReqTypeField is a INT field
type AllocIntermedReqTypeField struct{ quickfix.IntValue }

//Tag returns tag.AllocIntermedReqType (808)
func (f AllocIntermedReqTypeField) Tag() quickfix.Tag { return tag.AllocIntermedReqType }

//NewAllocIntermedReqType returns a new AllocIntermedReqTypeField initialized with val
func NewAllocIntermedReqType(val int) *AllocIntermedReqTypeField {
	field := &AllocIntermedReqTypeField{}
	field.Value = val
	return field
}

//AllocLinkIDField is a STRING field
type AllocLinkIDField struct{ quickfix.StringValue }

//Tag returns tag.AllocLinkID (196)
func (f AllocLinkIDField) Tag() quickfix.Tag { return tag.AllocLinkID }

//NewAllocLinkID returns a new AllocLinkIDField initialized with val
func NewAllocLinkID(val string) *AllocLinkIDField {
	field := &AllocLinkIDField{}
	field.Value = val
	return field
}

//AllocLinkTypeField is a INT field
type AllocLinkTypeField struct{ quickfix.IntValue }

//Tag returns tag.AllocLinkType (197)
func (f AllocLinkTypeField) Tag() quickfix.Tag { return tag.AllocLinkType }

//NewAllocLinkType returns a new AllocLinkTypeField initialized with val
func NewAllocLinkType(val int) *AllocLinkTypeField {
	field := &AllocLinkTypeField{}
	field.Value = val
	return field
}

//AllocMethodField is a INT field
type AllocMethodField struct{ quickfix.IntValue }

//Tag returns tag.AllocMethod (1002)
func (f AllocMethodField) Tag() quickfix.Tag { return tag.AllocMethod }

//NewAllocMethod returns a new AllocMethodField initialized with val
func NewAllocMethod(val int) *AllocMethodField {
	field := &AllocMethodField{}
	field.Value = val
	return field
}

//AllocNetMoneyField is a AMT field
type AllocNetMoneyField struct{ quickfix.AmtValue }

//Tag returns tag.AllocNetMoney (154)
func (f AllocNetMoneyField) Tag() quickfix.Tag { return tag.AllocNetMoney }

//NewAllocNetMoney returns a new AllocNetMoneyField initialized with val
func NewAllocNetMoney(val float64) *AllocNetMoneyField {
	field := &AllocNetMoneyField{}
	field.Value = val
	return field
}

//AllocNoOrdersTypeField is a INT field
type AllocNoOrdersTypeField struct{ quickfix.IntValue }

//Tag returns tag.AllocNoOrdersType (857)
func (f AllocNoOrdersTypeField) Tag() quickfix.Tag { return tag.AllocNoOrdersType }

//NewAllocNoOrdersType returns a new AllocNoOrdersTypeField initialized with val
func NewAllocNoOrdersType(val int) *AllocNoOrdersTypeField {
	field := &AllocNoOrdersTypeField{}
	field.Value = val
	return field
}

//AllocPositionEffectField is a CHAR field
type AllocPositionEffectField struct{ quickfix.CharValue }

//Tag returns tag.AllocPositionEffect (1047)
func (f AllocPositionEffectField) Tag() quickfix.Tag { return tag.AllocPositionEffect }

//NewAllocPositionEffect returns a new AllocPositionEffectField initialized with val
func NewAllocPositionEffect(val string) *AllocPositionEffectField {
	field := &AllocPositionEffectField{}
	field.Value = val
	return field
}

//AllocPriceField is a PRICE field
type AllocPriceField struct{ quickfix.PriceValue }

//Tag returns tag.AllocPrice (366)
func (f AllocPriceField) Tag() quickfix.Tag { return tag.AllocPrice }

//NewAllocPrice returns a new AllocPriceField initialized with val
func NewAllocPrice(val float64) *AllocPriceField {
	field := &AllocPriceField{}
	field.Value = val
	return field
}

//AllocQtyField is a QTY field
type AllocQtyField struct{ quickfix.QtyValue }

//Tag returns tag.AllocQty (80)
func (f AllocQtyField) Tag() quickfix.Tag { return tag.AllocQty }

//NewAllocQty returns a new AllocQtyField initialized with val
func NewAllocQty(val float64) *AllocQtyField {
	field := &AllocQtyField{}
	field.Value = val
	return field
}

//AllocRejCodeField is a INT field
type AllocRejCodeField struct{ quickfix.IntValue }

//Tag returns tag.AllocRejCode (88)
func (f AllocRejCodeField) Tag() quickfix.Tag { return tag.AllocRejCode }

//NewAllocRejCode returns a new AllocRejCodeField initialized with val
func NewAllocRejCode(val int) *AllocRejCodeField {
	field := &AllocRejCodeField{}
	field.Value = val
	return field
}

//AllocReportIDField is a STRING field
type AllocReportIDField struct{ quickfix.StringValue }

//Tag returns tag.AllocReportID (755)
func (f AllocReportIDField) Tag() quickfix.Tag { return tag.AllocReportID }

//NewAllocReportID returns a new AllocReportIDField initialized with val
func NewAllocReportID(val string) *AllocReportIDField {
	field := &AllocReportIDField{}
	field.Value = val
	return field
}

//AllocReportRefIDField is a STRING field
type AllocReportRefIDField struct{ quickfix.StringValue }

//Tag returns tag.AllocReportRefID (795)
func (f AllocReportRefIDField) Tag() quickfix.Tag { return tag.AllocReportRefID }

//NewAllocReportRefID returns a new AllocReportRefIDField initialized with val
func NewAllocReportRefID(val string) *AllocReportRefIDField {
	field := &AllocReportRefIDField{}
	field.Value = val
	return field
}

//AllocReportTypeField is a INT field
type AllocReportTypeField struct{ quickfix.IntValue }

//Tag returns tag.AllocReportType (794)
func (f AllocReportTypeField) Tag() quickfix.Tag { return tag.AllocReportType }

//NewAllocReportType returns a new AllocReportTypeField initialized with val
func NewAllocReportType(val int) *AllocReportTypeField {
	field := &AllocReportTypeField{}
	field.Value = val
	return field
}

//AllocSettlCurrAmtField is a AMT field
type AllocSettlCurrAmtField struct{ quickfix.AmtValue }

//Tag returns tag.AllocSettlCurrAmt (737)
func (f AllocSettlCurrAmtField) Tag() quickfix.Tag { return tag.AllocSettlCurrAmt }

//NewAllocSettlCurrAmt returns a new AllocSettlCurrAmtField initialized with val
func NewAllocSettlCurrAmt(val float64) *AllocSettlCurrAmtField {
	field := &AllocSettlCurrAmtField{}
	field.Value = val
	return field
}

//AllocSettlCurrencyField is a CURRENCY field
type AllocSettlCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.AllocSettlCurrency (736)
func (f AllocSettlCurrencyField) Tag() quickfix.Tag { return tag.AllocSettlCurrency }

//NewAllocSettlCurrency returns a new AllocSettlCurrencyField initialized with val
func NewAllocSettlCurrency(val string) *AllocSettlCurrencyField {
	field := &AllocSettlCurrencyField{}
	field.Value = val
	return field
}

//AllocSettlInstTypeField is a INT field
type AllocSettlInstTypeField struct{ quickfix.IntValue }

//Tag returns tag.AllocSettlInstType (780)
func (f AllocSettlInstTypeField) Tag() quickfix.Tag { return tag.AllocSettlInstType }

//NewAllocSettlInstType returns a new AllocSettlInstTypeField initialized with val
func NewAllocSettlInstType(val int) *AllocSettlInstTypeField {
	field := &AllocSettlInstTypeField{}
	field.Value = val
	return field
}

//AllocSharesField is a QTY field
type AllocSharesField struct{ quickfix.QtyValue }

//Tag returns tag.AllocShares (80)
func (f AllocSharesField) Tag() quickfix.Tag { return tag.AllocShares }

//NewAllocShares returns a new AllocSharesField initialized with val
func NewAllocShares(val float64) *AllocSharesField {
	field := &AllocSharesField{}
	field.Value = val
	return field
}

//AllocStatusField is a INT field
type AllocStatusField struct{ quickfix.IntValue }

//Tag returns tag.AllocStatus (87)
func (f AllocStatusField) Tag() quickfix.Tag { return tag.AllocStatus }

//NewAllocStatus returns a new AllocStatusField initialized with val
func NewAllocStatus(val int) *AllocStatusField {
	field := &AllocStatusField{}
	field.Value = val
	return field
}

//AllocTextField is a STRING field
type AllocTextField struct{ quickfix.StringValue }

//Tag returns tag.AllocText (161)
func (f AllocTextField) Tag() quickfix.Tag { return tag.AllocText }

//NewAllocText returns a new AllocTextField initialized with val
func NewAllocText(val string) *AllocTextField {
	field := &AllocTextField{}
	field.Value = val
	return field
}

//AllocTransTypeField is a CHAR field
type AllocTransTypeField struct{ quickfix.CharValue }

//Tag returns tag.AllocTransType (71)
func (f AllocTransTypeField) Tag() quickfix.Tag { return tag.AllocTransType }

//NewAllocTransType returns a new AllocTransTypeField initialized with val
func NewAllocTransType(val string) *AllocTransTypeField {
	field := &AllocTransTypeField{}
	field.Value = val
	return field
}

//AllocTypeField is a INT field
type AllocTypeField struct{ quickfix.IntValue }

//Tag returns tag.AllocType (626)
func (f AllocTypeField) Tag() quickfix.Tag { return tag.AllocType }

//NewAllocType returns a new AllocTypeField initialized with val
func NewAllocType(val int) *AllocTypeField {
	field := &AllocTypeField{}
	field.Value = val
	return field
}

//AllowableOneSidednessCurrField is a CURRENCY field
type AllowableOneSidednessCurrField struct{ quickfix.CurrencyValue }

//Tag returns tag.AllowableOneSidednessCurr (767)
func (f AllowableOneSidednessCurrField) Tag() quickfix.Tag { return tag.AllowableOneSidednessCurr }

//NewAllowableOneSidednessCurr returns a new AllowableOneSidednessCurrField initialized with val
func NewAllowableOneSidednessCurr(val string) *AllowableOneSidednessCurrField {
	field := &AllowableOneSidednessCurrField{}
	field.Value = val
	return field
}

//AllowableOneSidednessPctField is a PERCENTAGE field
type AllowableOneSidednessPctField struct{ quickfix.PercentageValue }

//Tag returns tag.AllowableOneSidednessPct (765)
func (f AllowableOneSidednessPctField) Tag() quickfix.Tag { return tag.AllowableOneSidednessPct }

//NewAllowableOneSidednessPct returns a new AllowableOneSidednessPctField initialized with val
func NewAllowableOneSidednessPct(val float64) *AllowableOneSidednessPctField {
	field := &AllowableOneSidednessPctField{}
	field.Value = val
	return field
}

//AllowableOneSidednessValueField is a AMT field
type AllowableOneSidednessValueField struct{ quickfix.AmtValue }

//Tag returns tag.AllowableOneSidednessValue (766)
func (f AllowableOneSidednessValueField) Tag() quickfix.Tag { return tag.AllowableOneSidednessValue }

//NewAllowableOneSidednessValue returns a new AllowableOneSidednessValueField initialized with val
func NewAllowableOneSidednessValue(val float64) *AllowableOneSidednessValueField {
	field := &AllowableOneSidednessValueField{}
	field.Value = val
	return field
}

//AltMDSourceIDField is a STRING field
type AltMDSourceIDField struct{ quickfix.StringValue }

//Tag returns tag.AltMDSourceID (817)
func (f AltMDSourceIDField) Tag() quickfix.Tag { return tag.AltMDSourceID }

//NewAltMDSourceID returns a new AltMDSourceIDField initialized with val
func NewAltMDSourceID(val string) *AltMDSourceIDField {
	field := &AltMDSourceIDField{}
	field.Value = val
	return field
}

//ApplBegSeqNumField is a SEQNUM field
type ApplBegSeqNumField struct{ quickfix.SeqNumValue }

//Tag returns tag.ApplBegSeqNum (1182)
func (f ApplBegSeqNumField) Tag() quickfix.Tag { return tag.ApplBegSeqNum }

//NewApplBegSeqNum returns a new ApplBegSeqNumField initialized with val
func NewApplBegSeqNum(val int) *ApplBegSeqNumField {
	field := &ApplBegSeqNumField{}
	field.Value = val
	return field
}

//ApplEndSeqNumField is a SEQNUM field
type ApplEndSeqNumField struct{ quickfix.SeqNumValue }

//Tag returns tag.ApplEndSeqNum (1183)
func (f ApplEndSeqNumField) Tag() quickfix.Tag { return tag.ApplEndSeqNum }

//NewApplEndSeqNum returns a new ApplEndSeqNumField initialized with val
func NewApplEndSeqNum(val int) *ApplEndSeqNumField {
	field := &ApplEndSeqNumField{}
	field.Value = val
	return field
}

//ApplExtIDField is a INT field
type ApplExtIDField struct{ quickfix.IntValue }

//Tag returns tag.ApplExtID (1156)
func (f ApplExtIDField) Tag() quickfix.Tag { return tag.ApplExtID }

//NewApplExtID returns a new ApplExtIDField initialized with val
func NewApplExtID(val int) *ApplExtIDField {
	field := &ApplExtIDField{}
	field.Value = val
	return field
}

//ApplIDField is a STRING field
type ApplIDField struct{ quickfix.StringValue }

//Tag returns tag.ApplID (1180)
func (f ApplIDField) Tag() quickfix.Tag { return tag.ApplID }

//NewApplID returns a new ApplIDField initialized with val
func NewApplID(val string) *ApplIDField {
	field := &ApplIDField{}
	field.Value = val
	return field
}

//ApplLastSeqNumField is a SEQNUM field
type ApplLastSeqNumField struct{ quickfix.SeqNumValue }

//Tag returns tag.ApplLastSeqNum (1350)
func (f ApplLastSeqNumField) Tag() quickfix.Tag { return tag.ApplLastSeqNum }

//NewApplLastSeqNum returns a new ApplLastSeqNumField initialized with val
func NewApplLastSeqNum(val int) *ApplLastSeqNumField {
	field := &ApplLastSeqNumField{}
	field.Value = val
	return field
}

//ApplNewSeqNumField is a SEQNUM field
type ApplNewSeqNumField struct{ quickfix.SeqNumValue }

//Tag returns tag.ApplNewSeqNum (1399)
func (f ApplNewSeqNumField) Tag() quickfix.Tag { return tag.ApplNewSeqNum }

//NewApplNewSeqNum returns a new ApplNewSeqNumField initialized with val
func NewApplNewSeqNum(val int) *ApplNewSeqNumField {
	field := &ApplNewSeqNumField{}
	field.Value = val
	return field
}

//ApplQueueActionField is a INT field
type ApplQueueActionField struct{ quickfix.IntValue }

//Tag returns tag.ApplQueueAction (815)
func (f ApplQueueActionField) Tag() quickfix.Tag { return tag.ApplQueueAction }

//NewApplQueueAction returns a new ApplQueueActionField initialized with val
func NewApplQueueAction(val int) *ApplQueueActionField {
	field := &ApplQueueActionField{}
	field.Value = val
	return field
}

//ApplQueueDepthField is a INT field
type ApplQueueDepthField struct{ quickfix.IntValue }

//Tag returns tag.ApplQueueDepth (813)
func (f ApplQueueDepthField) Tag() quickfix.Tag { return tag.ApplQueueDepth }

//NewApplQueueDepth returns a new ApplQueueDepthField initialized with val
func NewApplQueueDepth(val int) *ApplQueueDepthField {
	field := &ApplQueueDepthField{}
	field.Value = val
	return field
}

//ApplQueueMaxField is a INT field
type ApplQueueMaxField struct{ quickfix.IntValue }

//Tag returns tag.ApplQueueMax (812)
func (f ApplQueueMaxField) Tag() quickfix.Tag { return tag.ApplQueueMax }

//NewApplQueueMax returns a new ApplQueueMaxField initialized with val
func NewApplQueueMax(val int) *ApplQueueMaxField {
	field := &ApplQueueMaxField{}
	field.Value = val
	return field
}

//ApplQueueResolutionField is a INT field
type ApplQueueResolutionField struct{ quickfix.IntValue }

//Tag returns tag.ApplQueueResolution (814)
func (f ApplQueueResolutionField) Tag() quickfix.Tag { return tag.ApplQueueResolution }

//NewApplQueueResolution returns a new ApplQueueResolutionField initialized with val
func NewApplQueueResolution(val int) *ApplQueueResolutionField {
	field := &ApplQueueResolutionField{}
	field.Value = val
	return field
}

//ApplReportIDField is a STRING field
type ApplReportIDField struct{ quickfix.StringValue }

//Tag returns tag.ApplReportID (1356)
func (f ApplReportIDField) Tag() quickfix.Tag { return tag.ApplReportID }

//NewApplReportID returns a new ApplReportIDField initialized with val
func NewApplReportID(val string) *ApplReportIDField {
	field := &ApplReportIDField{}
	field.Value = val
	return field
}

//ApplReportTypeField is a INT field
type ApplReportTypeField struct{ quickfix.IntValue }

//Tag returns tag.ApplReportType (1426)
func (f ApplReportTypeField) Tag() quickfix.Tag { return tag.ApplReportType }

//NewApplReportType returns a new ApplReportTypeField initialized with val
func NewApplReportType(val int) *ApplReportTypeField {
	field := &ApplReportTypeField{}
	field.Value = val
	return field
}

//ApplReqIDField is a STRING field
type ApplReqIDField struct{ quickfix.StringValue }

//Tag returns tag.ApplReqID (1346)
func (f ApplReqIDField) Tag() quickfix.Tag { return tag.ApplReqID }

//NewApplReqID returns a new ApplReqIDField initialized with val
func NewApplReqID(val string) *ApplReqIDField {
	field := &ApplReqIDField{}
	field.Value = val
	return field
}

//ApplReqTypeField is a INT field
type ApplReqTypeField struct{ quickfix.IntValue }

//Tag returns tag.ApplReqType (1347)
func (f ApplReqTypeField) Tag() quickfix.Tag { return tag.ApplReqType }

//NewApplReqType returns a new ApplReqTypeField initialized with val
func NewApplReqType(val int) *ApplReqTypeField {
	field := &ApplReqTypeField{}
	field.Value = val
	return field
}

//ApplResendFlagField is a BOOLEAN field
type ApplResendFlagField struct{ quickfix.BooleanValue }

//Tag returns tag.ApplResendFlag (1352)
func (f ApplResendFlagField) Tag() quickfix.Tag { return tag.ApplResendFlag }

//NewApplResendFlag returns a new ApplResendFlagField initialized with val
func NewApplResendFlag(val bool) *ApplResendFlagField {
	field := &ApplResendFlagField{}
	field.Value = val
	return field
}

//ApplResponseErrorField is a INT field
type ApplResponseErrorField struct{ quickfix.IntValue }

//Tag returns tag.ApplResponseError (1354)
func (f ApplResponseErrorField) Tag() quickfix.Tag { return tag.ApplResponseError }

//NewApplResponseError returns a new ApplResponseErrorField initialized with val
func NewApplResponseError(val int) *ApplResponseErrorField {
	field := &ApplResponseErrorField{}
	field.Value = val
	return field
}

//ApplResponseIDField is a STRING field
type ApplResponseIDField struct{ quickfix.StringValue }

//Tag returns tag.ApplResponseID (1353)
func (f ApplResponseIDField) Tag() quickfix.Tag { return tag.ApplResponseID }

//NewApplResponseID returns a new ApplResponseIDField initialized with val
func NewApplResponseID(val string) *ApplResponseIDField {
	field := &ApplResponseIDField{}
	field.Value = val
	return field
}

//ApplResponseTypeField is a INT field
type ApplResponseTypeField struct{ quickfix.IntValue }

//Tag returns tag.ApplResponseType (1348)
func (f ApplResponseTypeField) Tag() quickfix.Tag { return tag.ApplResponseType }

//NewApplResponseType returns a new ApplResponseTypeField initialized with val
func NewApplResponseType(val int) *ApplResponseTypeField {
	field := &ApplResponseTypeField{}
	field.Value = val
	return field
}

//ApplSeqNumField is a SEQNUM field
type ApplSeqNumField struct{ quickfix.SeqNumValue }

//Tag returns tag.ApplSeqNum (1181)
func (f ApplSeqNumField) Tag() quickfix.Tag { return tag.ApplSeqNum }

//NewApplSeqNum returns a new ApplSeqNumField initialized with val
func NewApplSeqNum(val int) *ApplSeqNumField {
	field := &ApplSeqNumField{}
	field.Value = val
	return field
}

//ApplTotalMessageCountField is a INT field
type ApplTotalMessageCountField struct{ quickfix.IntValue }

//Tag returns tag.ApplTotalMessageCount (1349)
func (f ApplTotalMessageCountField) Tag() quickfix.Tag { return tag.ApplTotalMessageCount }

//NewApplTotalMessageCount returns a new ApplTotalMessageCountField initialized with val
func NewApplTotalMessageCount(val int) *ApplTotalMessageCountField {
	field := &ApplTotalMessageCountField{}
	field.Value = val
	return field
}

//ApplVerIDField is a STRING field
type ApplVerIDField struct{ quickfix.StringValue }

//Tag returns tag.ApplVerID (1128)
func (f ApplVerIDField) Tag() quickfix.Tag { return tag.ApplVerID }

//NewApplVerID returns a new ApplVerIDField initialized with val
func NewApplVerID(val string) *ApplVerIDField {
	field := &ApplVerIDField{}
	field.Value = val
	return field
}

//AsOfIndicatorField is a CHAR field
type AsOfIndicatorField struct{ quickfix.CharValue }

//Tag returns tag.AsOfIndicator (1015)
func (f AsOfIndicatorField) Tag() quickfix.Tag { return tag.AsOfIndicator }

//NewAsOfIndicator returns a new AsOfIndicatorField initialized with val
func NewAsOfIndicator(val string) *AsOfIndicatorField {
	field := &AsOfIndicatorField{}
	field.Value = val
	return field
}

//AsgnReqIDField is a STRING field
type AsgnReqIDField struct{ quickfix.StringValue }

//Tag returns tag.AsgnReqID (831)
func (f AsgnReqIDField) Tag() quickfix.Tag { return tag.AsgnReqID }

//NewAsgnReqID returns a new AsgnReqIDField initialized with val
func NewAsgnReqID(val string) *AsgnReqIDField {
	field := &AsgnReqIDField{}
	field.Value = val
	return field
}

//AsgnRptIDField is a STRING field
type AsgnRptIDField struct{ quickfix.StringValue }

//Tag returns tag.AsgnRptID (833)
func (f AsgnRptIDField) Tag() quickfix.Tag { return tag.AsgnRptID }

//NewAsgnRptID returns a new AsgnRptIDField initialized with val
func NewAsgnRptID(val string) *AsgnRptIDField {
	field := &AsgnRptIDField{}
	field.Value = val
	return field
}

//AssignmentMethodField is a CHAR field
type AssignmentMethodField struct{ quickfix.CharValue }

//Tag returns tag.AssignmentMethod (744)
func (f AssignmentMethodField) Tag() quickfix.Tag { return tag.AssignmentMethod }

//NewAssignmentMethod returns a new AssignmentMethodField initialized with val
func NewAssignmentMethod(val string) *AssignmentMethodField {
	field := &AssignmentMethodField{}
	field.Value = val
	return field
}

//AssignmentUnitField is a QTY field
type AssignmentUnitField struct{ quickfix.QtyValue }

//Tag returns tag.AssignmentUnit (745)
func (f AssignmentUnitField) Tag() quickfix.Tag { return tag.AssignmentUnit }

//NewAssignmentUnit returns a new AssignmentUnitField initialized with val
func NewAssignmentUnit(val float64) *AssignmentUnitField {
	field := &AssignmentUnitField{}
	field.Value = val
	return field
}

//AttachmentPointField is a PERCENTAGE field
type AttachmentPointField struct{ quickfix.PercentageValue }

//Tag returns tag.AttachmentPoint (1457)
func (f AttachmentPointField) Tag() quickfix.Tag { return tag.AttachmentPoint }

//NewAttachmentPoint returns a new AttachmentPointField initialized with val
func NewAttachmentPoint(val float64) *AttachmentPointField {
	field := &AttachmentPointField{}
	field.Value = val
	return field
}

//AutoAcceptIndicatorField is a BOOLEAN field
type AutoAcceptIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.AutoAcceptIndicator (754)
func (f AutoAcceptIndicatorField) Tag() quickfix.Tag { return tag.AutoAcceptIndicator }

//NewAutoAcceptIndicator returns a new AutoAcceptIndicatorField initialized with val
func NewAutoAcceptIndicator(val bool) *AutoAcceptIndicatorField {
	field := &AutoAcceptIndicatorField{}
	field.Value = val
	return field
}

//AvgParPxField is a PRICE field
type AvgParPxField struct{ quickfix.PriceValue }

//Tag returns tag.AvgParPx (860)
func (f AvgParPxField) Tag() quickfix.Tag { return tag.AvgParPx }

//NewAvgParPx returns a new AvgParPxField initialized with val
func NewAvgParPx(val float64) *AvgParPxField {
	field := &AvgParPxField{}
	field.Value = val
	return field
}

//AvgPrxPrecisionField is a INT field
type AvgPrxPrecisionField struct{ quickfix.IntValue }

//Tag returns tag.AvgPrxPrecision (74)
func (f AvgPrxPrecisionField) Tag() quickfix.Tag { return tag.AvgPrxPrecision }

//NewAvgPrxPrecision returns a new AvgPrxPrecisionField initialized with val
func NewAvgPrxPrecision(val int) *AvgPrxPrecisionField {
	field := &AvgPrxPrecisionField{}
	field.Value = val
	return field
}

//AvgPxField is a PRICE field
type AvgPxField struct{ quickfix.PriceValue }

//Tag returns tag.AvgPx (6)
func (f AvgPxField) Tag() quickfix.Tag { return tag.AvgPx }

//NewAvgPx returns a new AvgPxField initialized with val
func NewAvgPx(val float64) *AvgPxField {
	field := &AvgPxField{}
	field.Value = val
	return field
}

//AvgPxIndicatorField is a INT field
type AvgPxIndicatorField struct{ quickfix.IntValue }

//Tag returns tag.AvgPxIndicator (819)
func (f AvgPxIndicatorField) Tag() quickfix.Tag { return tag.AvgPxIndicator }

//NewAvgPxIndicator returns a new AvgPxIndicatorField initialized with val
func NewAvgPxIndicator(val int) *AvgPxIndicatorField {
	field := &AvgPxIndicatorField{}
	field.Value = val
	return field
}

//AvgPxPrecisionField is a INT field
type AvgPxPrecisionField struct{ quickfix.IntValue }

//Tag returns tag.AvgPxPrecision (74)
func (f AvgPxPrecisionField) Tag() quickfix.Tag { return tag.AvgPxPrecision }

//NewAvgPxPrecision returns a new AvgPxPrecisionField initialized with val
func NewAvgPxPrecision(val int) *AvgPxPrecisionField {
	field := &AvgPxPrecisionField{}
	field.Value = val
	return field
}

//BasisFeatureDateField is a LOCALMKTDATE field
type BasisFeatureDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.BasisFeatureDate (259)
func (f BasisFeatureDateField) Tag() quickfix.Tag { return tag.BasisFeatureDate }

//NewBasisFeatureDate returns a new BasisFeatureDateField initialized with val
func NewBasisFeatureDate(val string) *BasisFeatureDateField {
	field := &BasisFeatureDateField{}
	field.Value = val
	return field
}

//BasisFeaturePriceField is a PRICE field
type BasisFeaturePriceField struct{ quickfix.PriceValue }

//Tag returns tag.BasisFeaturePrice (260)
func (f BasisFeaturePriceField) Tag() quickfix.Tag { return tag.BasisFeaturePrice }

//NewBasisFeaturePrice returns a new BasisFeaturePriceField initialized with val
func NewBasisFeaturePrice(val float64) *BasisFeaturePriceField {
	field := &BasisFeaturePriceField{}
	field.Value = val
	return field
}

//BasisPxTypeField is a CHAR field
type BasisPxTypeField struct{ quickfix.CharValue }

//Tag returns tag.BasisPxType (419)
func (f BasisPxTypeField) Tag() quickfix.Tag { return tag.BasisPxType }

//NewBasisPxType returns a new BasisPxTypeField initialized with val
func NewBasisPxType(val string) *BasisPxTypeField {
	field := &BasisPxTypeField{}
	field.Value = val
	return field
}

//BeginSeqNoField is a SEQNUM field
type BeginSeqNoField struct{ quickfix.SeqNumValue }

//Tag returns tag.BeginSeqNo (7)
func (f BeginSeqNoField) Tag() quickfix.Tag { return tag.BeginSeqNo }

//NewBeginSeqNo returns a new BeginSeqNoField initialized with val
func NewBeginSeqNo(val int) *BeginSeqNoField {
	field := &BeginSeqNoField{}
	field.Value = val
	return field
}

//BeginStringField is a STRING field
type BeginStringField struct{ quickfix.StringValue }

//Tag returns tag.BeginString (8)
func (f BeginStringField) Tag() quickfix.Tag { return tag.BeginString }

//NewBeginString returns a new BeginStringField initialized with val
func NewBeginString(val string) *BeginStringField {
	field := &BeginStringField{}
	field.Value = val
	return field
}

//BenchmarkField is a CHAR field
type BenchmarkField struct{ quickfix.CharValue }

//Tag returns tag.Benchmark (219)
func (f BenchmarkField) Tag() quickfix.Tag { return tag.Benchmark }

//NewBenchmark returns a new BenchmarkField initialized with val
func NewBenchmark(val string) *BenchmarkField {
	field := &BenchmarkField{}
	field.Value = val
	return field
}

//BenchmarkCurveCurrencyField is a CURRENCY field
type BenchmarkCurveCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.BenchmarkCurveCurrency (220)
func (f BenchmarkCurveCurrencyField) Tag() quickfix.Tag { return tag.BenchmarkCurveCurrency }

//NewBenchmarkCurveCurrency returns a new BenchmarkCurveCurrencyField initialized with val
func NewBenchmarkCurveCurrency(val string) *BenchmarkCurveCurrencyField {
	field := &BenchmarkCurveCurrencyField{}
	field.Value = val
	return field
}

//BenchmarkCurveNameField is a STRING field
type BenchmarkCurveNameField struct{ quickfix.StringValue }

//Tag returns tag.BenchmarkCurveName (221)
func (f BenchmarkCurveNameField) Tag() quickfix.Tag { return tag.BenchmarkCurveName }

//NewBenchmarkCurveName returns a new BenchmarkCurveNameField initialized with val
func NewBenchmarkCurveName(val string) *BenchmarkCurveNameField {
	field := &BenchmarkCurveNameField{}
	field.Value = val
	return field
}

//BenchmarkCurvePointField is a STRING field
type BenchmarkCurvePointField struct{ quickfix.StringValue }

//Tag returns tag.BenchmarkCurvePoint (222)
func (f BenchmarkCurvePointField) Tag() quickfix.Tag { return tag.BenchmarkCurvePoint }

//NewBenchmarkCurvePoint returns a new BenchmarkCurvePointField initialized with val
func NewBenchmarkCurvePoint(val string) *BenchmarkCurvePointField {
	field := &BenchmarkCurvePointField{}
	field.Value = val
	return field
}

//BenchmarkPriceField is a PRICE field
type BenchmarkPriceField struct{ quickfix.PriceValue }

//Tag returns tag.BenchmarkPrice (662)
func (f BenchmarkPriceField) Tag() quickfix.Tag { return tag.BenchmarkPrice }

//NewBenchmarkPrice returns a new BenchmarkPriceField initialized with val
func NewBenchmarkPrice(val float64) *BenchmarkPriceField {
	field := &BenchmarkPriceField{}
	field.Value = val
	return field
}

//BenchmarkPriceTypeField is a INT field
type BenchmarkPriceTypeField struct{ quickfix.IntValue }

//Tag returns tag.BenchmarkPriceType (663)
func (f BenchmarkPriceTypeField) Tag() quickfix.Tag { return tag.BenchmarkPriceType }

//NewBenchmarkPriceType returns a new BenchmarkPriceTypeField initialized with val
func NewBenchmarkPriceType(val int) *BenchmarkPriceTypeField {
	field := &BenchmarkPriceTypeField{}
	field.Value = val
	return field
}

//BenchmarkSecurityIDField is a STRING field
type BenchmarkSecurityIDField struct{ quickfix.StringValue }

//Tag returns tag.BenchmarkSecurityID (699)
func (f BenchmarkSecurityIDField) Tag() quickfix.Tag { return tag.BenchmarkSecurityID }

//NewBenchmarkSecurityID returns a new BenchmarkSecurityIDField initialized with val
func NewBenchmarkSecurityID(val string) *BenchmarkSecurityIDField {
	field := &BenchmarkSecurityIDField{}
	field.Value = val
	return field
}

//BenchmarkSecurityIDSourceField is a STRING field
type BenchmarkSecurityIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.BenchmarkSecurityIDSource (761)
func (f BenchmarkSecurityIDSourceField) Tag() quickfix.Tag { return tag.BenchmarkSecurityIDSource }

//NewBenchmarkSecurityIDSource returns a new BenchmarkSecurityIDSourceField initialized with val
func NewBenchmarkSecurityIDSource(val string) *BenchmarkSecurityIDSourceField {
	field := &BenchmarkSecurityIDSourceField{}
	field.Value = val
	return field
}

//BidDescriptorField is a STRING field
type BidDescriptorField struct{ quickfix.StringValue }

//Tag returns tag.BidDescriptor (400)
func (f BidDescriptorField) Tag() quickfix.Tag { return tag.BidDescriptor }

//NewBidDescriptor returns a new BidDescriptorField initialized with val
func NewBidDescriptor(val string) *BidDescriptorField {
	field := &BidDescriptorField{}
	field.Value = val
	return field
}

//BidDescriptorTypeField is a INT field
type BidDescriptorTypeField struct{ quickfix.IntValue }

//Tag returns tag.BidDescriptorType (399)
func (f BidDescriptorTypeField) Tag() quickfix.Tag { return tag.BidDescriptorType }

//NewBidDescriptorType returns a new BidDescriptorTypeField initialized with val
func NewBidDescriptorType(val int) *BidDescriptorTypeField {
	field := &BidDescriptorTypeField{}
	field.Value = val
	return field
}

//BidForwardPointsField is a PRICEOFFSET field
type BidForwardPointsField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.BidForwardPoints (189)
func (f BidForwardPointsField) Tag() quickfix.Tag { return tag.BidForwardPoints }

//NewBidForwardPoints returns a new BidForwardPointsField initialized with val
func NewBidForwardPoints(val float64) *BidForwardPointsField {
	field := &BidForwardPointsField{}
	field.Value = val
	return field
}

//BidForwardPoints2Field is a PRICEOFFSET field
type BidForwardPoints2Field struct{ quickfix.PriceOffsetValue }

//Tag returns tag.BidForwardPoints2 (642)
func (f BidForwardPoints2Field) Tag() quickfix.Tag { return tag.BidForwardPoints2 }

//NewBidForwardPoints2 returns a new BidForwardPoints2Field initialized with val
func NewBidForwardPoints2(val float64) *BidForwardPoints2Field {
	field := &BidForwardPoints2Field{}
	field.Value = val
	return field
}

//BidIDField is a STRING field
type BidIDField struct{ quickfix.StringValue }

//Tag returns tag.BidID (390)
func (f BidIDField) Tag() quickfix.Tag { return tag.BidID }

//NewBidID returns a new BidIDField initialized with val
func NewBidID(val string) *BidIDField {
	field := &BidIDField{}
	field.Value = val
	return field
}

//BidPxField is a PRICE field
type BidPxField struct{ quickfix.PriceValue }

//Tag returns tag.BidPx (132)
func (f BidPxField) Tag() quickfix.Tag { return tag.BidPx }

//NewBidPx returns a new BidPxField initialized with val
func NewBidPx(val float64) *BidPxField {
	field := &BidPxField{}
	field.Value = val
	return field
}

//BidRequestTransTypeField is a CHAR field
type BidRequestTransTypeField struct{ quickfix.CharValue }

//Tag returns tag.BidRequestTransType (374)
func (f BidRequestTransTypeField) Tag() quickfix.Tag { return tag.BidRequestTransType }

//NewBidRequestTransType returns a new BidRequestTransTypeField initialized with val
func NewBidRequestTransType(val string) *BidRequestTransTypeField {
	field := &BidRequestTransTypeField{}
	field.Value = val
	return field
}

//BidSizeField is a QTY field
type BidSizeField struct{ quickfix.QtyValue }

//Tag returns tag.BidSize (134)
func (f BidSizeField) Tag() quickfix.Tag { return tag.BidSize }

//NewBidSize returns a new BidSizeField initialized with val
func NewBidSize(val float64) *BidSizeField {
	field := &BidSizeField{}
	field.Value = val
	return field
}

//BidSpotRateField is a PRICE field
type BidSpotRateField struct{ quickfix.PriceValue }

//Tag returns tag.BidSpotRate (188)
func (f BidSpotRateField) Tag() quickfix.Tag { return tag.BidSpotRate }

//NewBidSpotRate returns a new BidSpotRateField initialized with val
func NewBidSpotRate(val float64) *BidSpotRateField {
	field := &BidSpotRateField{}
	field.Value = val
	return field
}

//BidSwapPointsField is a PRICEOFFSET field
type BidSwapPointsField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.BidSwapPoints (1065)
func (f BidSwapPointsField) Tag() quickfix.Tag { return tag.BidSwapPoints }

//NewBidSwapPoints returns a new BidSwapPointsField initialized with val
func NewBidSwapPoints(val float64) *BidSwapPointsField {
	field := &BidSwapPointsField{}
	field.Value = val
	return field
}

//BidTradeTypeField is a CHAR field
type BidTradeTypeField struct{ quickfix.CharValue }

//Tag returns tag.BidTradeType (418)
func (f BidTradeTypeField) Tag() quickfix.Tag { return tag.BidTradeType }

//NewBidTradeType returns a new BidTradeTypeField initialized with val
func NewBidTradeType(val string) *BidTradeTypeField {
	field := &BidTradeTypeField{}
	field.Value = val
	return field
}

//BidTypeField is a INT field
type BidTypeField struct{ quickfix.IntValue }

//Tag returns tag.BidType (394)
func (f BidTypeField) Tag() quickfix.Tag { return tag.BidType }

//NewBidType returns a new BidTypeField initialized with val
func NewBidType(val int) *BidTypeField {
	field := &BidTypeField{}
	field.Value = val
	return field
}

//BidYieldField is a PERCENTAGE field
type BidYieldField struct{ quickfix.PercentageValue }

//Tag returns tag.BidYield (632)
func (f BidYieldField) Tag() quickfix.Tag { return tag.BidYield }

//NewBidYield returns a new BidYieldField initialized with val
func NewBidYield(val float64) *BidYieldField {
	field := &BidYieldField{}
	field.Value = val
	return field
}

//BodyLengthField is a LENGTH field
type BodyLengthField struct{ quickfix.LengthValue }

//Tag returns tag.BodyLength (9)
func (f BodyLengthField) Tag() quickfix.Tag { return tag.BodyLength }

//NewBodyLength returns a new BodyLengthField initialized with val
func NewBodyLength(val int) *BodyLengthField {
	field := &BodyLengthField{}
	field.Value = val
	return field
}

//BookingRefIDField is a STRING field
type BookingRefIDField struct{ quickfix.StringValue }

//Tag returns tag.BookingRefID (466)
func (f BookingRefIDField) Tag() quickfix.Tag { return tag.BookingRefID }

//NewBookingRefID returns a new BookingRefIDField initialized with val
func NewBookingRefID(val string) *BookingRefIDField {
	field := &BookingRefIDField{}
	field.Value = val
	return field
}

//BookingTypeField is a INT field
type BookingTypeField struct{ quickfix.IntValue }

//Tag returns tag.BookingType (775)
func (f BookingTypeField) Tag() quickfix.Tag { return tag.BookingType }

//NewBookingType returns a new BookingTypeField initialized with val
func NewBookingType(val int) *BookingTypeField {
	field := &BookingTypeField{}
	field.Value = val
	return field
}

//BookingUnitField is a CHAR field
type BookingUnitField struct{ quickfix.CharValue }

//Tag returns tag.BookingUnit (590)
func (f BookingUnitField) Tag() quickfix.Tag { return tag.BookingUnit }

//NewBookingUnit returns a new BookingUnitField initialized with val
func NewBookingUnit(val string) *BookingUnitField {
	field := &BookingUnitField{}
	field.Value = val
	return field
}

//BrokerOfCreditField is a STRING field
type BrokerOfCreditField struct{ quickfix.StringValue }

//Tag returns tag.BrokerOfCredit (92)
func (f BrokerOfCreditField) Tag() quickfix.Tag { return tag.BrokerOfCredit }

//NewBrokerOfCredit returns a new BrokerOfCreditField initialized with val
func NewBrokerOfCredit(val string) *BrokerOfCreditField {
	field := &BrokerOfCreditField{}
	field.Value = val
	return field
}

//BusinessRejectReasonField is a INT field
type BusinessRejectReasonField struct{ quickfix.IntValue }

//Tag returns tag.BusinessRejectReason (380)
func (f BusinessRejectReasonField) Tag() quickfix.Tag { return tag.BusinessRejectReason }

//NewBusinessRejectReason returns a new BusinessRejectReasonField initialized with val
func NewBusinessRejectReason(val int) *BusinessRejectReasonField {
	field := &BusinessRejectReasonField{}
	field.Value = val
	return field
}

//BusinessRejectRefIDField is a STRING field
type BusinessRejectRefIDField struct{ quickfix.StringValue }

//Tag returns tag.BusinessRejectRefID (379)
func (f BusinessRejectRefIDField) Tag() quickfix.Tag { return tag.BusinessRejectRefID }

//NewBusinessRejectRefID returns a new BusinessRejectRefIDField initialized with val
func NewBusinessRejectRefID(val string) *BusinessRejectRefIDField {
	field := &BusinessRejectRefIDField{}
	field.Value = val
	return field
}

//BuyVolumeField is a QTY field
type BuyVolumeField struct{ quickfix.QtyValue }

//Tag returns tag.BuyVolume (330)
func (f BuyVolumeField) Tag() quickfix.Tag { return tag.BuyVolume }

//NewBuyVolume returns a new BuyVolumeField initialized with val
func NewBuyVolume(val float64) *BuyVolumeField {
	field := &BuyVolumeField{}
	field.Value = val
	return field
}

//CFICodeField is a STRING field
type CFICodeField struct{ quickfix.StringValue }

//Tag returns tag.CFICode (461)
func (f CFICodeField) Tag() quickfix.Tag { return tag.CFICode }

//NewCFICode returns a new CFICodeField initialized with val
func NewCFICode(val string) *CFICodeField {
	field := &CFICodeField{}
	field.Value = val
	return field
}

//CPProgramField is a INT field
type CPProgramField struct{ quickfix.IntValue }

//Tag returns tag.CPProgram (875)
func (f CPProgramField) Tag() quickfix.Tag { return tag.CPProgram }

//NewCPProgram returns a new CPProgramField initialized with val
func NewCPProgram(val int) *CPProgramField {
	field := &CPProgramField{}
	field.Value = val
	return field
}

//CPRegTypeField is a STRING field
type CPRegTypeField struct{ quickfix.StringValue }

//Tag returns tag.CPRegType (876)
func (f CPRegTypeField) Tag() quickfix.Tag { return tag.CPRegType }

//NewCPRegType returns a new CPRegTypeField initialized with val
func NewCPRegType(val string) *CPRegTypeField {
	field := &CPRegTypeField{}
	field.Value = val
	return field
}

//CalculatedCcyLastQtyField is a QTY field
type CalculatedCcyLastQtyField struct{ quickfix.QtyValue }

//Tag returns tag.CalculatedCcyLastQty (1056)
func (f CalculatedCcyLastQtyField) Tag() quickfix.Tag { return tag.CalculatedCcyLastQty }

//NewCalculatedCcyLastQty returns a new CalculatedCcyLastQtyField initialized with val
func NewCalculatedCcyLastQty(val float64) *CalculatedCcyLastQtyField {
	field := &CalculatedCcyLastQtyField{}
	field.Value = val
	return field
}

//CancellationRightsField is a CHAR field
type CancellationRightsField struct{ quickfix.CharValue }

//Tag returns tag.CancellationRights (480)
func (f CancellationRightsField) Tag() quickfix.Tag { return tag.CancellationRights }

//NewCancellationRights returns a new CancellationRightsField initialized with val
func NewCancellationRights(val string) *CancellationRightsField {
	field := &CancellationRightsField{}
	field.Value = val
	return field
}

//CapPriceField is a PRICE field
type CapPriceField struct{ quickfix.PriceValue }

//Tag returns tag.CapPrice (1199)
func (f CapPriceField) Tag() quickfix.Tag { return tag.CapPrice }

//NewCapPrice returns a new CapPriceField initialized with val
func NewCapPrice(val float64) *CapPriceField {
	field := &CapPriceField{}
	field.Value = val
	return field
}

//CardExpDateField is a LOCALMKTDATE field
type CardExpDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.CardExpDate (490)
func (f CardExpDateField) Tag() quickfix.Tag { return tag.CardExpDate }

//NewCardExpDate returns a new CardExpDateField initialized with val
func NewCardExpDate(val string) *CardExpDateField {
	field := &CardExpDateField{}
	field.Value = val
	return field
}

//CardHolderNameField is a STRING field
type CardHolderNameField struct{ quickfix.StringValue }

//Tag returns tag.CardHolderName (488)
func (f CardHolderNameField) Tag() quickfix.Tag { return tag.CardHolderName }

//NewCardHolderName returns a new CardHolderNameField initialized with val
func NewCardHolderName(val string) *CardHolderNameField {
	field := &CardHolderNameField{}
	field.Value = val
	return field
}

//CardIssNoField is a STRING field
type CardIssNoField struct{ quickfix.StringValue }

//Tag returns tag.CardIssNo (491)
func (f CardIssNoField) Tag() quickfix.Tag { return tag.CardIssNo }

//NewCardIssNo returns a new CardIssNoField initialized with val
func NewCardIssNo(val string) *CardIssNoField {
	field := &CardIssNoField{}
	field.Value = val
	return field
}

//CardIssNumField is a STRING field
type CardIssNumField struct{ quickfix.StringValue }

//Tag returns tag.CardIssNum (491)
func (f CardIssNumField) Tag() quickfix.Tag { return tag.CardIssNum }

//NewCardIssNum returns a new CardIssNumField initialized with val
func NewCardIssNum(val string) *CardIssNumField {
	field := &CardIssNumField{}
	field.Value = val
	return field
}

//CardNumberField is a STRING field
type CardNumberField struct{ quickfix.StringValue }

//Tag returns tag.CardNumber (489)
func (f CardNumberField) Tag() quickfix.Tag { return tag.CardNumber }

//NewCardNumber returns a new CardNumberField initialized with val
func NewCardNumber(val string) *CardNumberField {
	field := &CardNumberField{}
	field.Value = val
	return field
}

//CardStartDateField is a LOCALMKTDATE field
type CardStartDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.CardStartDate (503)
func (f CardStartDateField) Tag() quickfix.Tag { return tag.CardStartDate }

//NewCardStartDate returns a new CardStartDateField initialized with val
func NewCardStartDate(val string) *CardStartDateField {
	field := &CardStartDateField{}
	field.Value = val
	return field
}

//CashDistribAgentAcctNameField is a STRING field
type CashDistribAgentAcctNameField struct{ quickfix.StringValue }

//Tag returns tag.CashDistribAgentAcctName (502)
func (f CashDistribAgentAcctNameField) Tag() quickfix.Tag { return tag.CashDistribAgentAcctName }

//NewCashDistribAgentAcctName returns a new CashDistribAgentAcctNameField initialized with val
func NewCashDistribAgentAcctName(val string) *CashDistribAgentAcctNameField {
	field := &CashDistribAgentAcctNameField{}
	field.Value = val
	return field
}

//CashDistribAgentAcctNumberField is a STRING field
type CashDistribAgentAcctNumberField struct{ quickfix.StringValue }

//Tag returns tag.CashDistribAgentAcctNumber (500)
func (f CashDistribAgentAcctNumberField) Tag() quickfix.Tag { return tag.CashDistribAgentAcctNumber }

//NewCashDistribAgentAcctNumber returns a new CashDistribAgentAcctNumberField initialized with val
func NewCashDistribAgentAcctNumber(val string) *CashDistribAgentAcctNumberField {
	field := &CashDistribAgentAcctNumberField{}
	field.Value = val
	return field
}

//CashDistribAgentCodeField is a STRING field
type CashDistribAgentCodeField struct{ quickfix.StringValue }

//Tag returns tag.CashDistribAgentCode (499)
func (f CashDistribAgentCodeField) Tag() quickfix.Tag { return tag.CashDistribAgentCode }

//NewCashDistribAgentCode returns a new CashDistribAgentCodeField initialized with val
func NewCashDistribAgentCode(val string) *CashDistribAgentCodeField {
	field := &CashDistribAgentCodeField{}
	field.Value = val
	return field
}

//CashDistribAgentNameField is a STRING field
type CashDistribAgentNameField struct{ quickfix.StringValue }

//Tag returns tag.CashDistribAgentName (498)
func (f CashDistribAgentNameField) Tag() quickfix.Tag { return tag.CashDistribAgentName }

//NewCashDistribAgentName returns a new CashDistribAgentNameField initialized with val
func NewCashDistribAgentName(val string) *CashDistribAgentNameField {
	field := &CashDistribAgentNameField{}
	field.Value = val
	return field
}

//CashDistribCurrField is a CURRENCY field
type CashDistribCurrField struct{ quickfix.CurrencyValue }

//Tag returns tag.CashDistribCurr (478)
func (f CashDistribCurrField) Tag() quickfix.Tag { return tag.CashDistribCurr }

//NewCashDistribCurr returns a new CashDistribCurrField initialized with val
func NewCashDistribCurr(val string) *CashDistribCurrField {
	field := &CashDistribCurrField{}
	field.Value = val
	return field
}

//CashDistribPayRefField is a STRING field
type CashDistribPayRefField struct{ quickfix.StringValue }

//Tag returns tag.CashDistribPayRef (501)
func (f CashDistribPayRefField) Tag() quickfix.Tag { return tag.CashDistribPayRef }

//NewCashDistribPayRef returns a new CashDistribPayRefField initialized with val
func NewCashDistribPayRef(val string) *CashDistribPayRefField {
	field := &CashDistribPayRefField{}
	field.Value = val
	return field
}

//CashMarginField is a CHAR field
type CashMarginField struct{ quickfix.CharValue }

//Tag returns tag.CashMargin (544)
func (f CashMarginField) Tag() quickfix.Tag { return tag.CashMargin }

//NewCashMargin returns a new CashMarginField initialized with val
func NewCashMargin(val string) *CashMarginField {
	field := &CashMarginField{}
	field.Value = val
	return field
}

//CashOrderQtyField is a QTY field
type CashOrderQtyField struct{ quickfix.QtyValue }

//Tag returns tag.CashOrderQty (152)
func (f CashOrderQtyField) Tag() quickfix.Tag { return tag.CashOrderQty }

//NewCashOrderQty returns a new CashOrderQtyField initialized with val
func NewCashOrderQty(val float64) *CashOrderQtyField {
	field := &CashOrderQtyField{}
	field.Value = val
	return field
}

//CashOutstandingField is a AMT field
type CashOutstandingField struct{ quickfix.AmtValue }

//Tag returns tag.CashOutstanding (901)
func (f CashOutstandingField) Tag() quickfix.Tag { return tag.CashOutstanding }

//NewCashOutstanding returns a new CashOutstandingField initialized with val
func NewCashOutstanding(val float64) *CashOutstandingField {
	field := &CashOutstandingField{}
	field.Value = val
	return field
}

//CashSettlAgentAcctNameField is a STRING field
type CashSettlAgentAcctNameField struct{ quickfix.StringValue }

//Tag returns tag.CashSettlAgentAcctName (185)
func (f CashSettlAgentAcctNameField) Tag() quickfix.Tag { return tag.CashSettlAgentAcctName }

//NewCashSettlAgentAcctName returns a new CashSettlAgentAcctNameField initialized with val
func NewCashSettlAgentAcctName(val string) *CashSettlAgentAcctNameField {
	field := &CashSettlAgentAcctNameField{}
	field.Value = val
	return field
}

//CashSettlAgentAcctNumField is a STRING field
type CashSettlAgentAcctNumField struct{ quickfix.StringValue }

//Tag returns tag.CashSettlAgentAcctNum (184)
func (f CashSettlAgentAcctNumField) Tag() quickfix.Tag { return tag.CashSettlAgentAcctNum }

//NewCashSettlAgentAcctNum returns a new CashSettlAgentAcctNumField initialized with val
func NewCashSettlAgentAcctNum(val string) *CashSettlAgentAcctNumField {
	field := &CashSettlAgentAcctNumField{}
	field.Value = val
	return field
}

//CashSettlAgentCodeField is a STRING field
type CashSettlAgentCodeField struct{ quickfix.StringValue }

//Tag returns tag.CashSettlAgentCode (183)
func (f CashSettlAgentCodeField) Tag() quickfix.Tag { return tag.CashSettlAgentCode }

//NewCashSettlAgentCode returns a new CashSettlAgentCodeField initialized with val
func NewCashSettlAgentCode(val string) *CashSettlAgentCodeField {
	field := &CashSettlAgentCodeField{}
	field.Value = val
	return field
}

//CashSettlAgentContactNameField is a STRING field
type CashSettlAgentContactNameField struct{ quickfix.StringValue }

//Tag returns tag.CashSettlAgentContactName (186)
func (f CashSettlAgentContactNameField) Tag() quickfix.Tag { return tag.CashSettlAgentContactName }

//NewCashSettlAgentContactName returns a new CashSettlAgentContactNameField initialized with val
func NewCashSettlAgentContactName(val string) *CashSettlAgentContactNameField {
	field := &CashSettlAgentContactNameField{}
	field.Value = val
	return field
}

//CashSettlAgentContactPhoneField is a STRING field
type CashSettlAgentContactPhoneField struct{ quickfix.StringValue }

//Tag returns tag.CashSettlAgentContactPhone (187)
func (f CashSettlAgentContactPhoneField) Tag() quickfix.Tag { return tag.CashSettlAgentContactPhone }

//NewCashSettlAgentContactPhone returns a new CashSettlAgentContactPhoneField initialized with val
func NewCashSettlAgentContactPhone(val string) *CashSettlAgentContactPhoneField {
	field := &CashSettlAgentContactPhoneField{}
	field.Value = val
	return field
}

//CashSettlAgentNameField is a STRING field
type CashSettlAgentNameField struct{ quickfix.StringValue }

//Tag returns tag.CashSettlAgentName (182)
func (f CashSettlAgentNameField) Tag() quickfix.Tag { return tag.CashSettlAgentName }

//NewCashSettlAgentName returns a new CashSettlAgentNameField initialized with val
func NewCashSettlAgentName(val string) *CashSettlAgentNameField {
	field := &CashSettlAgentNameField{}
	field.Value = val
	return field
}

//CcyAmtField is a AMT field
type CcyAmtField struct{ quickfix.AmtValue }

//Tag returns tag.CcyAmt (1157)
func (f CcyAmtField) Tag() quickfix.Tag { return tag.CcyAmt }

//NewCcyAmt returns a new CcyAmtField initialized with val
func NewCcyAmt(val float64) *CcyAmtField {
	field := &CcyAmtField{}
	field.Value = val
	return field
}

//CheckSumField is a STRING field
type CheckSumField struct{ quickfix.StringValue }

//Tag returns tag.CheckSum (10)
func (f CheckSumField) Tag() quickfix.Tag { return tag.CheckSum }

//NewCheckSum returns a new CheckSumField initialized with val
func NewCheckSum(val string) *CheckSumField {
	field := &CheckSumField{}
	field.Value = val
	return field
}

//ClOrdIDField is a STRING field
type ClOrdIDField struct{ quickfix.StringValue }

//Tag returns tag.ClOrdID (11)
func (f ClOrdIDField) Tag() quickfix.Tag { return tag.ClOrdID }

//NewClOrdID returns a new ClOrdIDField initialized with val
func NewClOrdID(val string) *ClOrdIDField {
	field := &ClOrdIDField{}
	field.Value = val
	return field
}

//ClOrdLinkIDField is a STRING field
type ClOrdLinkIDField struct{ quickfix.StringValue }

//Tag returns tag.ClOrdLinkID (583)
func (f ClOrdLinkIDField) Tag() quickfix.Tag { return tag.ClOrdLinkID }

//NewClOrdLinkID returns a new ClOrdLinkIDField initialized with val
func NewClOrdLinkID(val string) *ClOrdLinkIDField {
	field := &ClOrdLinkIDField{}
	field.Value = val
	return field
}

//ClearingAccountField is a STRING field
type ClearingAccountField struct{ quickfix.StringValue }

//Tag returns tag.ClearingAccount (440)
func (f ClearingAccountField) Tag() quickfix.Tag { return tag.ClearingAccount }

//NewClearingAccount returns a new ClearingAccountField initialized with val
func NewClearingAccount(val string) *ClearingAccountField {
	field := &ClearingAccountField{}
	field.Value = val
	return field
}

//ClearingBusinessDateField is a LOCALMKTDATE field
type ClearingBusinessDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.ClearingBusinessDate (715)
func (f ClearingBusinessDateField) Tag() quickfix.Tag { return tag.ClearingBusinessDate }

//NewClearingBusinessDate returns a new ClearingBusinessDateField initialized with val
func NewClearingBusinessDate(val string) *ClearingBusinessDateField {
	field := &ClearingBusinessDateField{}
	field.Value = val
	return field
}

//ClearingFeeIndicatorField is a STRING field
type ClearingFeeIndicatorField struct{ quickfix.StringValue }

//Tag returns tag.ClearingFeeIndicator (635)
func (f ClearingFeeIndicatorField) Tag() quickfix.Tag { return tag.ClearingFeeIndicator }

//NewClearingFeeIndicator returns a new ClearingFeeIndicatorField initialized with val
func NewClearingFeeIndicator(val string) *ClearingFeeIndicatorField {
	field := &ClearingFeeIndicatorField{}
	field.Value = val
	return field
}

//ClearingFirmField is a STRING field
type ClearingFirmField struct{ quickfix.StringValue }

//Tag returns tag.ClearingFirm (439)
func (f ClearingFirmField) Tag() quickfix.Tag { return tag.ClearingFirm }

//NewClearingFirm returns a new ClearingFirmField initialized with val
func NewClearingFirm(val string) *ClearingFirmField {
	field := &ClearingFirmField{}
	field.Value = val
	return field
}

//ClearingInstructionField is a INT field
type ClearingInstructionField struct{ quickfix.IntValue }

//Tag returns tag.ClearingInstruction (577)
func (f ClearingInstructionField) Tag() quickfix.Tag { return tag.ClearingInstruction }

//NewClearingInstruction returns a new ClearingInstructionField initialized with val
func NewClearingInstruction(val int) *ClearingInstructionField {
	field := &ClearingInstructionField{}
	field.Value = val
	return field
}

//ClientBidIDField is a STRING field
type ClientBidIDField struct{ quickfix.StringValue }

//Tag returns tag.ClientBidID (391)
func (f ClientBidIDField) Tag() quickfix.Tag { return tag.ClientBidID }

//NewClientBidID returns a new ClientBidIDField initialized with val
func NewClientBidID(val string) *ClientBidIDField {
	field := &ClientBidIDField{}
	field.Value = val
	return field
}

//ClientIDField is a STRING field
type ClientIDField struct{ quickfix.StringValue }

//Tag returns tag.ClientID (109)
func (f ClientIDField) Tag() quickfix.Tag { return tag.ClientID }

//NewClientID returns a new ClientIDField initialized with val
func NewClientID(val string) *ClientIDField {
	field := &ClientIDField{}
	field.Value = val
	return field
}

//CollActionField is a INT field
type CollActionField struct{ quickfix.IntValue }

//Tag returns tag.CollAction (944)
func (f CollActionField) Tag() quickfix.Tag { return tag.CollAction }

//NewCollAction returns a new CollActionField initialized with val
func NewCollAction(val int) *CollActionField {
	field := &CollActionField{}
	field.Value = val
	return field
}

//CollApplTypeField is a INT field
type CollApplTypeField struct{ quickfix.IntValue }

//Tag returns tag.CollApplType (1043)
func (f CollApplTypeField) Tag() quickfix.Tag { return tag.CollApplType }

//NewCollApplType returns a new CollApplTypeField initialized with val
func NewCollApplType(val int) *CollApplTypeField {
	field := &CollApplTypeField{}
	field.Value = val
	return field
}

//CollAsgnIDField is a STRING field
type CollAsgnIDField struct{ quickfix.StringValue }

//Tag returns tag.CollAsgnID (902)
func (f CollAsgnIDField) Tag() quickfix.Tag { return tag.CollAsgnID }

//NewCollAsgnID returns a new CollAsgnIDField initialized with val
func NewCollAsgnID(val string) *CollAsgnIDField {
	field := &CollAsgnIDField{}
	field.Value = val
	return field
}

//CollAsgnReasonField is a INT field
type CollAsgnReasonField struct{ quickfix.IntValue }

//Tag returns tag.CollAsgnReason (895)
func (f CollAsgnReasonField) Tag() quickfix.Tag { return tag.CollAsgnReason }

//NewCollAsgnReason returns a new CollAsgnReasonField initialized with val
func NewCollAsgnReason(val int) *CollAsgnReasonField {
	field := &CollAsgnReasonField{}
	field.Value = val
	return field
}

//CollAsgnRefIDField is a STRING field
type CollAsgnRefIDField struct{ quickfix.StringValue }

//Tag returns tag.CollAsgnRefID (907)
func (f CollAsgnRefIDField) Tag() quickfix.Tag { return tag.CollAsgnRefID }

//NewCollAsgnRefID returns a new CollAsgnRefIDField initialized with val
func NewCollAsgnRefID(val string) *CollAsgnRefIDField {
	field := &CollAsgnRefIDField{}
	field.Value = val
	return field
}

//CollAsgnRejectReasonField is a INT field
type CollAsgnRejectReasonField struct{ quickfix.IntValue }

//Tag returns tag.CollAsgnRejectReason (906)
func (f CollAsgnRejectReasonField) Tag() quickfix.Tag { return tag.CollAsgnRejectReason }

//NewCollAsgnRejectReason returns a new CollAsgnRejectReasonField initialized with val
func NewCollAsgnRejectReason(val int) *CollAsgnRejectReasonField {
	field := &CollAsgnRejectReasonField{}
	field.Value = val
	return field
}

//CollAsgnRespTypeField is a INT field
type CollAsgnRespTypeField struct{ quickfix.IntValue }

//Tag returns tag.CollAsgnRespType (905)
func (f CollAsgnRespTypeField) Tag() quickfix.Tag { return tag.CollAsgnRespType }

//NewCollAsgnRespType returns a new CollAsgnRespTypeField initialized with val
func NewCollAsgnRespType(val int) *CollAsgnRespTypeField {
	field := &CollAsgnRespTypeField{}
	field.Value = val
	return field
}

//CollAsgnTransTypeField is a INT field
type CollAsgnTransTypeField struct{ quickfix.IntValue }

//Tag returns tag.CollAsgnTransType (903)
func (f CollAsgnTransTypeField) Tag() quickfix.Tag { return tag.CollAsgnTransType }

//NewCollAsgnTransType returns a new CollAsgnTransTypeField initialized with val
func NewCollAsgnTransType(val int) *CollAsgnTransTypeField {
	field := &CollAsgnTransTypeField{}
	field.Value = val
	return field
}

//CollInquiryIDField is a STRING field
type CollInquiryIDField struct{ quickfix.StringValue }

//Tag returns tag.CollInquiryID (909)
func (f CollInquiryIDField) Tag() quickfix.Tag { return tag.CollInquiryID }

//NewCollInquiryID returns a new CollInquiryIDField initialized with val
func NewCollInquiryID(val string) *CollInquiryIDField {
	field := &CollInquiryIDField{}
	field.Value = val
	return field
}

//CollInquiryQualifierField is a INT field
type CollInquiryQualifierField struct{ quickfix.IntValue }

//Tag returns tag.CollInquiryQualifier (896)
func (f CollInquiryQualifierField) Tag() quickfix.Tag { return tag.CollInquiryQualifier }

//NewCollInquiryQualifier returns a new CollInquiryQualifierField initialized with val
func NewCollInquiryQualifier(val int) *CollInquiryQualifierField {
	field := &CollInquiryQualifierField{}
	field.Value = val
	return field
}

//CollInquiryResultField is a INT field
type CollInquiryResultField struct{ quickfix.IntValue }

//Tag returns tag.CollInquiryResult (946)
func (f CollInquiryResultField) Tag() quickfix.Tag { return tag.CollInquiryResult }

//NewCollInquiryResult returns a new CollInquiryResultField initialized with val
func NewCollInquiryResult(val int) *CollInquiryResultField {
	field := &CollInquiryResultField{}
	field.Value = val
	return field
}

//CollInquiryStatusField is a INT field
type CollInquiryStatusField struct{ quickfix.IntValue }

//Tag returns tag.CollInquiryStatus (945)
func (f CollInquiryStatusField) Tag() quickfix.Tag { return tag.CollInquiryStatus }

//NewCollInquiryStatus returns a new CollInquiryStatusField initialized with val
func NewCollInquiryStatus(val int) *CollInquiryStatusField {
	field := &CollInquiryStatusField{}
	field.Value = val
	return field
}

//CollReqIDField is a STRING field
type CollReqIDField struct{ quickfix.StringValue }

//Tag returns tag.CollReqID (894)
func (f CollReqIDField) Tag() quickfix.Tag { return tag.CollReqID }

//NewCollReqID returns a new CollReqIDField initialized with val
func NewCollReqID(val string) *CollReqIDField {
	field := &CollReqIDField{}
	field.Value = val
	return field
}

//CollRespIDField is a STRING field
type CollRespIDField struct{ quickfix.StringValue }

//Tag returns tag.CollRespID (904)
func (f CollRespIDField) Tag() quickfix.Tag { return tag.CollRespID }

//NewCollRespID returns a new CollRespIDField initialized with val
func NewCollRespID(val string) *CollRespIDField {
	field := &CollRespIDField{}
	field.Value = val
	return field
}

//CollRptIDField is a STRING field
type CollRptIDField struct{ quickfix.StringValue }

//Tag returns tag.CollRptID (908)
func (f CollRptIDField) Tag() quickfix.Tag { return tag.CollRptID }

//NewCollRptID returns a new CollRptIDField initialized with val
func NewCollRptID(val string) *CollRptIDField {
	field := &CollRptIDField{}
	field.Value = val
	return field
}

//CollStatusField is a INT field
type CollStatusField struct{ quickfix.IntValue }

//Tag returns tag.CollStatus (910)
func (f CollStatusField) Tag() quickfix.Tag { return tag.CollStatus }

//NewCollStatus returns a new CollStatusField initialized with val
func NewCollStatus(val int) *CollStatusField {
	field := &CollStatusField{}
	field.Value = val
	return field
}

//CommCurrencyField is a CURRENCY field
type CommCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.CommCurrency (479)
func (f CommCurrencyField) Tag() quickfix.Tag { return tag.CommCurrency }

//NewCommCurrency returns a new CommCurrencyField initialized with val
func NewCommCurrency(val string) *CommCurrencyField {
	field := &CommCurrencyField{}
	field.Value = val
	return field
}

//CommTypeField is a CHAR field
type CommTypeField struct{ quickfix.CharValue }

//Tag returns tag.CommType (13)
func (f CommTypeField) Tag() quickfix.Tag { return tag.CommType }

//NewCommType returns a new CommTypeField initialized with val
func NewCommType(val string) *CommTypeField {
	field := &CommTypeField{}
	field.Value = val
	return field
}

//CommissionField is a AMT field
type CommissionField struct{ quickfix.AmtValue }

//Tag returns tag.Commission (12)
func (f CommissionField) Tag() quickfix.Tag { return tag.Commission }

//NewCommission returns a new CommissionField initialized with val
func NewCommission(val float64) *CommissionField {
	field := &CommissionField{}
	field.Value = val
	return field
}

//ComplexEventConditionField is a INT field
type ComplexEventConditionField struct{ quickfix.IntValue }

//Tag returns tag.ComplexEventCondition (1490)
func (f ComplexEventConditionField) Tag() quickfix.Tag { return tag.ComplexEventCondition }

//NewComplexEventCondition returns a new ComplexEventConditionField initialized with val
func NewComplexEventCondition(val int) *ComplexEventConditionField {
	field := &ComplexEventConditionField{}
	field.Value = val
	return field
}

//ComplexEventEndDateField is a UTCTIMESTAMP field
type ComplexEventEndDateField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.ComplexEventEndDate (1493)
func (f ComplexEventEndDateField) Tag() quickfix.Tag { return tag.ComplexEventEndDate }

//ComplexEventEndTimeField is a UTCTIMEONLY field
type ComplexEventEndTimeField struct{ quickfix.UTCTimeOnlyValue }

//Tag returns tag.ComplexEventEndTime (1496)
func (f ComplexEventEndTimeField) Tag() quickfix.Tag { return tag.ComplexEventEndTime }

//ComplexEventPriceField is a PRICE field
type ComplexEventPriceField struct{ quickfix.PriceValue }

//Tag returns tag.ComplexEventPrice (1486)
func (f ComplexEventPriceField) Tag() quickfix.Tag { return tag.ComplexEventPrice }

//NewComplexEventPrice returns a new ComplexEventPriceField initialized with val
func NewComplexEventPrice(val float64) *ComplexEventPriceField {
	field := &ComplexEventPriceField{}
	field.Value = val
	return field
}

//ComplexEventPriceBoundaryMethodField is a INT field
type ComplexEventPriceBoundaryMethodField struct{ quickfix.IntValue }

//Tag returns tag.ComplexEventPriceBoundaryMethod (1487)
func (f ComplexEventPriceBoundaryMethodField) Tag() quickfix.Tag {
	return tag.ComplexEventPriceBoundaryMethod
}

//NewComplexEventPriceBoundaryMethod returns a new ComplexEventPriceBoundaryMethodField initialized with val
func NewComplexEventPriceBoundaryMethod(val int) *ComplexEventPriceBoundaryMethodField {
	field := &ComplexEventPriceBoundaryMethodField{}
	field.Value = val
	return field
}

//ComplexEventPriceBoundaryPrecisionField is a PERCENTAGE field
type ComplexEventPriceBoundaryPrecisionField struct{ quickfix.PercentageValue }

//Tag returns tag.ComplexEventPriceBoundaryPrecision (1488)
func (f ComplexEventPriceBoundaryPrecisionField) Tag() quickfix.Tag {
	return tag.ComplexEventPriceBoundaryPrecision
}

//NewComplexEventPriceBoundaryPrecision returns a new ComplexEventPriceBoundaryPrecisionField initialized with val
func NewComplexEventPriceBoundaryPrecision(val float64) *ComplexEventPriceBoundaryPrecisionField {
	field := &ComplexEventPriceBoundaryPrecisionField{}
	field.Value = val
	return field
}

//ComplexEventPriceTimeTypeField is a INT field
type ComplexEventPriceTimeTypeField struct{ quickfix.IntValue }

//Tag returns tag.ComplexEventPriceTimeType (1489)
func (f ComplexEventPriceTimeTypeField) Tag() quickfix.Tag { return tag.ComplexEventPriceTimeType }

//NewComplexEventPriceTimeType returns a new ComplexEventPriceTimeTypeField initialized with val
func NewComplexEventPriceTimeType(val int) *ComplexEventPriceTimeTypeField {
	field := &ComplexEventPriceTimeTypeField{}
	field.Value = val
	return field
}

//ComplexEventStartDateField is a UTCTIMESTAMP field
type ComplexEventStartDateField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.ComplexEventStartDate (1492)
func (f ComplexEventStartDateField) Tag() quickfix.Tag { return tag.ComplexEventStartDate }

//ComplexEventStartTimeField is a UTCTIMEONLY field
type ComplexEventStartTimeField struct{ quickfix.UTCTimeOnlyValue }

//Tag returns tag.ComplexEventStartTime (1495)
func (f ComplexEventStartTimeField) Tag() quickfix.Tag { return tag.ComplexEventStartTime }

//ComplexEventTypeField is a INT field
type ComplexEventTypeField struct{ quickfix.IntValue }

//Tag returns tag.ComplexEventType (1484)
func (f ComplexEventTypeField) Tag() quickfix.Tag { return tag.ComplexEventType }

//NewComplexEventType returns a new ComplexEventTypeField initialized with val
func NewComplexEventType(val int) *ComplexEventTypeField {
	field := &ComplexEventTypeField{}
	field.Value = val
	return field
}

//ComplexOptPayoutAmountField is a AMT field
type ComplexOptPayoutAmountField struct{ quickfix.AmtValue }

//Tag returns tag.ComplexOptPayoutAmount (1485)
func (f ComplexOptPayoutAmountField) Tag() quickfix.Tag { return tag.ComplexOptPayoutAmount }

//NewComplexOptPayoutAmount returns a new ComplexOptPayoutAmountField initialized with val
func NewComplexOptPayoutAmount(val float64) *ComplexOptPayoutAmountField {
	field := &ComplexOptPayoutAmountField{}
	field.Value = val
	return field
}

//ComplianceIDField is a STRING field
type ComplianceIDField struct{ quickfix.StringValue }

//Tag returns tag.ComplianceID (376)
func (f ComplianceIDField) Tag() quickfix.Tag { return tag.ComplianceID }

//NewComplianceID returns a new ComplianceIDField initialized with val
func NewComplianceID(val string) *ComplianceIDField {
	field := &ComplianceIDField{}
	field.Value = val
	return field
}

//ConcessionField is a AMT field
type ConcessionField struct{ quickfix.AmtValue }

//Tag returns tag.Concession (238)
func (f ConcessionField) Tag() quickfix.Tag { return tag.Concession }

//NewConcession returns a new ConcessionField initialized with val
func NewConcession(val float64) *ConcessionField {
	field := &ConcessionField{}
	field.Value = val
	return field
}

//ConfirmIDField is a STRING field
type ConfirmIDField struct{ quickfix.StringValue }

//Tag returns tag.ConfirmID (664)
func (f ConfirmIDField) Tag() quickfix.Tag { return tag.ConfirmID }

//NewConfirmID returns a new ConfirmIDField initialized with val
func NewConfirmID(val string) *ConfirmIDField {
	field := &ConfirmIDField{}
	field.Value = val
	return field
}

//ConfirmRefIDField is a STRING field
type ConfirmRefIDField struct{ quickfix.StringValue }

//Tag returns tag.ConfirmRefID (772)
func (f ConfirmRefIDField) Tag() quickfix.Tag { return tag.ConfirmRefID }

//NewConfirmRefID returns a new ConfirmRefIDField initialized with val
func NewConfirmRefID(val string) *ConfirmRefIDField {
	field := &ConfirmRefIDField{}
	field.Value = val
	return field
}

//ConfirmRejReasonField is a INT field
type ConfirmRejReasonField struct{ quickfix.IntValue }

//Tag returns tag.ConfirmRejReason (774)
func (f ConfirmRejReasonField) Tag() quickfix.Tag { return tag.ConfirmRejReason }

//NewConfirmRejReason returns a new ConfirmRejReasonField initialized with val
func NewConfirmRejReason(val int) *ConfirmRejReasonField {
	field := &ConfirmRejReasonField{}
	field.Value = val
	return field
}

//ConfirmReqIDField is a STRING field
type ConfirmReqIDField struct{ quickfix.StringValue }

//Tag returns tag.ConfirmReqID (859)
func (f ConfirmReqIDField) Tag() quickfix.Tag { return tag.ConfirmReqID }

//NewConfirmReqID returns a new ConfirmReqIDField initialized with val
func NewConfirmReqID(val string) *ConfirmReqIDField {
	field := &ConfirmReqIDField{}
	field.Value = val
	return field
}

//ConfirmStatusField is a INT field
type ConfirmStatusField struct{ quickfix.IntValue }

//Tag returns tag.ConfirmStatus (665)
func (f ConfirmStatusField) Tag() quickfix.Tag { return tag.ConfirmStatus }

//NewConfirmStatus returns a new ConfirmStatusField initialized with val
func NewConfirmStatus(val int) *ConfirmStatusField {
	field := &ConfirmStatusField{}
	field.Value = val
	return field
}

//ConfirmTransTypeField is a INT field
type ConfirmTransTypeField struct{ quickfix.IntValue }

//Tag returns tag.ConfirmTransType (666)
func (f ConfirmTransTypeField) Tag() quickfix.Tag { return tag.ConfirmTransType }

//NewConfirmTransType returns a new ConfirmTransTypeField initialized with val
func NewConfirmTransType(val int) *ConfirmTransTypeField {
	field := &ConfirmTransTypeField{}
	field.Value = val
	return field
}

//ConfirmTypeField is a INT field
type ConfirmTypeField struct{ quickfix.IntValue }

//Tag returns tag.ConfirmType (773)
func (f ConfirmTypeField) Tag() quickfix.Tag { return tag.ConfirmType }

//NewConfirmType returns a new ConfirmTypeField initialized with val
func NewConfirmType(val int) *ConfirmTypeField {
	field := &ConfirmTypeField{}
	field.Value = val
	return field
}

//ContAmtCurrField is a CURRENCY field
type ContAmtCurrField struct{ quickfix.CurrencyValue }

//Tag returns tag.ContAmtCurr (521)
func (f ContAmtCurrField) Tag() quickfix.Tag { return tag.ContAmtCurr }

//NewContAmtCurr returns a new ContAmtCurrField initialized with val
func NewContAmtCurr(val string) *ContAmtCurrField {
	field := &ContAmtCurrField{}
	field.Value = val
	return field
}

//ContAmtTypeField is a INT field
type ContAmtTypeField struct{ quickfix.IntValue }

//Tag returns tag.ContAmtType (519)
func (f ContAmtTypeField) Tag() quickfix.Tag { return tag.ContAmtType }

//NewContAmtType returns a new ContAmtTypeField initialized with val
func NewContAmtType(val int) *ContAmtTypeField {
	field := &ContAmtTypeField{}
	field.Value = val
	return field
}

//ContAmtValueField is a FLOAT field
type ContAmtValueField struct{ quickfix.FloatValue }

//Tag returns tag.ContAmtValue (520)
func (f ContAmtValueField) Tag() quickfix.Tag { return tag.ContAmtValue }

//NewContAmtValue returns a new ContAmtValueField initialized with val
func NewContAmtValue(val float64) *ContAmtValueField {
	field := &ContAmtValueField{}
	field.Value = val
	return field
}

//ContIntRptIDField is a STRING field
type ContIntRptIDField struct{ quickfix.StringValue }

//Tag returns tag.ContIntRptID (977)
func (f ContIntRptIDField) Tag() quickfix.Tag { return tag.ContIntRptID }

//NewContIntRptID returns a new ContIntRptIDField initialized with val
func NewContIntRptID(val string) *ContIntRptIDField {
	field := &ContIntRptIDField{}
	field.Value = val
	return field
}

//ContextPartyIDField is a STRING field
type ContextPartyIDField struct{ quickfix.StringValue }

//Tag returns tag.ContextPartyID (1523)
func (f ContextPartyIDField) Tag() quickfix.Tag { return tag.ContextPartyID }

//NewContextPartyID returns a new ContextPartyIDField initialized with val
func NewContextPartyID(val string) *ContextPartyIDField {
	field := &ContextPartyIDField{}
	field.Value = val
	return field
}

//ContextPartyIDSourceField is a CHAR field
type ContextPartyIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.ContextPartyIDSource (1524)
func (f ContextPartyIDSourceField) Tag() quickfix.Tag { return tag.ContextPartyIDSource }

//NewContextPartyIDSource returns a new ContextPartyIDSourceField initialized with val
func NewContextPartyIDSource(val string) *ContextPartyIDSourceField {
	field := &ContextPartyIDSourceField{}
	field.Value = val
	return field
}

//ContextPartyRoleField is a INT field
type ContextPartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.ContextPartyRole (1525)
func (f ContextPartyRoleField) Tag() quickfix.Tag { return tag.ContextPartyRole }

//NewContextPartyRole returns a new ContextPartyRoleField initialized with val
func NewContextPartyRole(val int) *ContextPartyRoleField {
	field := &ContextPartyRoleField{}
	field.Value = val
	return field
}

//ContextPartySubIDField is a STRING field
type ContextPartySubIDField struct{ quickfix.StringValue }

//Tag returns tag.ContextPartySubID (1527)
func (f ContextPartySubIDField) Tag() quickfix.Tag { return tag.ContextPartySubID }

//NewContextPartySubID returns a new ContextPartySubIDField initialized with val
func NewContextPartySubID(val string) *ContextPartySubIDField {
	field := &ContextPartySubIDField{}
	field.Value = val
	return field
}

//ContextPartySubIDTypeField is a INT field
type ContextPartySubIDTypeField struct{ quickfix.IntValue }

//Tag returns tag.ContextPartySubIDType (1528)
func (f ContextPartySubIDTypeField) Tag() quickfix.Tag { return tag.ContextPartySubIDType }

//NewContextPartySubIDType returns a new ContextPartySubIDTypeField initialized with val
func NewContextPartySubIDType(val int) *ContextPartySubIDTypeField {
	field := &ContextPartySubIDTypeField{}
	field.Value = val
	return field
}

//ContingencyTypeField is a INT field
type ContingencyTypeField struct{ quickfix.IntValue }

//Tag returns tag.ContingencyType (1385)
func (f ContingencyTypeField) Tag() quickfix.Tag { return tag.ContingencyType }

//NewContingencyType returns a new ContingencyTypeField initialized with val
func NewContingencyType(val int) *ContingencyTypeField {
	field := &ContingencyTypeField{}
	field.Value = val
	return field
}

//ContraBrokerField is a STRING field
type ContraBrokerField struct{ quickfix.StringValue }

//Tag returns tag.ContraBroker (375)
func (f ContraBrokerField) Tag() quickfix.Tag { return tag.ContraBroker }

//NewContraBroker returns a new ContraBrokerField initialized with val
func NewContraBroker(val string) *ContraBrokerField {
	field := &ContraBrokerField{}
	field.Value = val
	return field
}

//ContraLegRefIDField is a STRING field
type ContraLegRefIDField struct{ quickfix.StringValue }

//Tag returns tag.ContraLegRefID (655)
func (f ContraLegRefIDField) Tag() quickfix.Tag { return tag.ContraLegRefID }

//NewContraLegRefID returns a new ContraLegRefIDField initialized with val
func NewContraLegRefID(val string) *ContraLegRefIDField {
	field := &ContraLegRefIDField{}
	field.Value = val
	return field
}

//ContraTradeQtyField is a QTY field
type ContraTradeQtyField struct{ quickfix.QtyValue }

//Tag returns tag.ContraTradeQty (437)
func (f ContraTradeQtyField) Tag() quickfix.Tag { return tag.ContraTradeQty }

//NewContraTradeQty returns a new ContraTradeQtyField initialized with val
func NewContraTradeQty(val float64) *ContraTradeQtyField {
	field := &ContraTradeQtyField{}
	field.Value = val
	return field
}

//ContraTradeTimeField is a UTCTIMESTAMP field
type ContraTradeTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.ContraTradeTime (438)
func (f ContraTradeTimeField) Tag() quickfix.Tag { return tag.ContraTradeTime }

//ContraTraderField is a STRING field
type ContraTraderField struct{ quickfix.StringValue }

//Tag returns tag.ContraTrader (337)
func (f ContraTraderField) Tag() quickfix.Tag { return tag.ContraTrader }

//NewContraTrader returns a new ContraTraderField initialized with val
func NewContraTrader(val string) *ContraTraderField {
	field := &ContraTraderField{}
	field.Value = val
	return field
}

//ContractMultiplierField is a FLOAT field
type ContractMultiplierField struct{ quickfix.FloatValue }

//Tag returns tag.ContractMultiplier (231)
func (f ContractMultiplierField) Tag() quickfix.Tag { return tag.ContractMultiplier }

//NewContractMultiplier returns a new ContractMultiplierField initialized with val
func NewContractMultiplier(val float64) *ContractMultiplierField {
	field := &ContractMultiplierField{}
	field.Value = val
	return field
}

//ContractMultiplierUnitField is a INT field
type ContractMultiplierUnitField struct{ quickfix.IntValue }

//Tag returns tag.ContractMultiplierUnit (1435)
func (f ContractMultiplierUnitField) Tag() quickfix.Tag { return tag.ContractMultiplierUnit }

//NewContractMultiplierUnit returns a new ContractMultiplierUnitField initialized with val
func NewContractMultiplierUnit(val int) *ContractMultiplierUnitField {
	field := &ContractMultiplierUnitField{}
	field.Value = val
	return field
}

//ContractSettlMonthField is a MONTHYEAR field
type ContractSettlMonthField struct{ quickfix.MonthYearValue }

//Tag returns tag.ContractSettlMonth (667)
func (f ContractSettlMonthField) Tag() quickfix.Tag { return tag.ContractSettlMonth }

//NewContractSettlMonth returns a new ContractSettlMonthField initialized with val
func NewContractSettlMonth(val string) *ContractSettlMonthField {
	field := &ContractSettlMonthField{}
	field.Value = val
	return field
}

//ContraryInstructionIndicatorField is a BOOLEAN field
type ContraryInstructionIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.ContraryInstructionIndicator (719)
func (f ContraryInstructionIndicatorField) Tag() quickfix.Tag { return tag.ContraryInstructionIndicator }

//NewContraryInstructionIndicator returns a new ContraryInstructionIndicatorField initialized with val
func NewContraryInstructionIndicator(val bool) *ContraryInstructionIndicatorField {
	field := &ContraryInstructionIndicatorField{}
	field.Value = val
	return field
}

//CopyMsgIndicatorField is a BOOLEAN field
type CopyMsgIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.CopyMsgIndicator (797)
func (f CopyMsgIndicatorField) Tag() quickfix.Tag { return tag.CopyMsgIndicator }

//NewCopyMsgIndicator returns a new CopyMsgIndicatorField initialized with val
func NewCopyMsgIndicator(val bool) *CopyMsgIndicatorField {
	field := &CopyMsgIndicatorField{}
	field.Value = val
	return field
}

//CorporateActionField is a MULTIPLECHARVALUE field
type CorporateActionField struct{ quickfix.MultipleCharValue }

//Tag returns tag.CorporateAction (292)
func (f CorporateActionField) Tag() quickfix.Tag { return tag.CorporateAction }

//NewCorporateAction returns a new CorporateActionField initialized with val
func NewCorporateAction(val string) *CorporateActionField {
	field := &CorporateActionField{}
	field.Value = val
	return field
}

//CountryField is a COUNTRY field
type CountryField struct{ quickfix.CountryValue }

//Tag returns tag.Country (421)
func (f CountryField) Tag() quickfix.Tag { return tag.Country }

//NewCountry returns a new CountryField initialized with val
func NewCountry(val string) *CountryField {
	field := &CountryField{}
	field.Value = val
	return field
}

//CountryOfIssueField is a COUNTRY field
type CountryOfIssueField struct{ quickfix.CountryValue }

//Tag returns tag.CountryOfIssue (470)
func (f CountryOfIssueField) Tag() quickfix.Tag { return tag.CountryOfIssue }

//NewCountryOfIssue returns a new CountryOfIssueField initialized with val
func NewCountryOfIssue(val string) *CountryOfIssueField {
	field := &CountryOfIssueField{}
	field.Value = val
	return field
}

//CouponPaymentDateField is a LOCALMKTDATE field
type CouponPaymentDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.CouponPaymentDate (224)
func (f CouponPaymentDateField) Tag() quickfix.Tag { return tag.CouponPaymentDate }

//NewCouponPaymentDate returns a new CouponPaymentDateField initialized with val
func NewCouponPaymentDate(val string) *CouponPaymentDateField {
	field := &CouponPaymentDateField{}
	field.Value = val
	return field
}

//CouponRateField is a PERCENTAGE field
type CouponRateField struct{ quickfix.PercentageValue }

//Tag returns tag.CouponRate (223)
func (f CouponRateField) Tag() quickfix.Tag { return tag.CouponRate }

//NewCouponRate returns a new CouponRateField initialized with val
func NewCouponRate(val float64) *CouponRateField {
	field := &CouponRateField{}
	field.Value = val
	return field
}

//CoveredOrUncoveredField is a INT field
type CoveredOrUncoveredField struct{ quickfix.IntValue }

//Tag returns tag.CoveredOrUncovered (203)
func (f CoveredOrUncoveredField) Tag() quickfix.Tag { return tag.CoveredOrUncovered }

//NewCoveredOrUncovered returns a new CoveredOrUncoveredField initialized with val
func NewCoveredOrUncovered(val int) *CoveredOrUncoveredField {
	field := &CoveredOrUncoveredField{}
	field.Value = val
	return field
}

//CreditRatingField is a STRING field
type CreditRatingField struct{ quickfix.StringValue }

//Tag returns tag.CreditRating (255)
func (f CreditRatingField) Tag() quickfix.Tag { return tag.CreditRating }

//NewCreditRating returns a new CreditRatingField initialized with val
func NewCreditRating(val string) *CreditRatingField {
	field := &CreditRatingField{}
	field.Value = val
	return field
}

//CrossIDField is a STRING field
type CrossIDField struct{ quickfix.StringValue }

//Tag returns tag.CrossID (548)
func (f CrossIDField) Tag() quickfix.Tag { return tag.CrossID }

//NewCrossID returns a new CrossIDField initialized with val
func NewCrossID(val string) *CrossIDField {
	field := &CrossIDField{}
	field.Value = val
	return field
}

//CrossPercentField is a PERCENTAGE field
type CrossPercentField struct{ quickfix.PercentageValue }

//Tag returns tag.CrossPercent (413)
func (f CrossPercentField) Tag() quickfix.Tag { return tag.CrossPercent }

//NewCrossPercent returns a new CrossPercentField initialized with val
func NewCrossPercent(val float64) *CrossPercentField {
	field := &CrossPercentField{}
	field.Value = val
	return field
}

//CrossPrioritizationField is a INT field
type CrossPrioritizationField struct{ quickfix.IntValue }

//Tag returns tag.CrossPrioritization (550)
func (f CrossPrioritizationField) Tag() quickfix.Tag { return tag.CrossPrioritization }

//NewCrossPrioritization returns a new CrossPrioritizationField initialized with val
func NewCrossPrioritization(val int) *CrossPrioritizationField {
	field := &CrossPrioritizationField{}
	field.Value = val
	return field
}

//CrossTypeField is a INT field
type CrossTypeField struct{ quickfix.IntValue }

//Tag returns tag.CrossType (549)
func (f CrossTypeField) Tag() quickfix.Tag { return tag.CrossType }

//NewCrossType returns a new CrossTypeField initialized with val
func NewCrossType(val int) *CrossTypeField {
	field := &CrossTypeField{}
	field.Value = val
	return field
}

//CstmApplVerIDField is a STRING field
type CstmApplVerIDField struct{ quickfix.StringValue }

//Tag returns tag.CstmApplVerID (1129)
func (f CstmApplVerIDField) Tag() quickfix.Tag { return tag.CstmApplVerID }

//NewCstmApplVerID returns a new CstmApplVerIDField initialized with val
func NewCstmApplVerID(val string) *CstmApplVerIDField {
	field := &CstmApplVerIDField{}
	field.Value = val
	return field
}

//CumQtyField is a QTY field
type CumQtyField struct{ quickfix.QtyValue }

//Tag returns tag.CumQty (14)
func (f CumQtyField) Tag() quickfix.Tag { return tag.CumQty }

//NewCumQty returns a new CumQtyField initialized with val
func NewCumQty(val float64) *CumQtyField {
	field := &CumQtyField{}
	field.Value = val
	return field
}

//CurrencyField is a CURRENCY field
type CurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.Currency (15)
func (f CurrencyField) Tag() quickfix.Tag { return tag.Currency }

//NewCurrency returns a new CurrencyField initialized with val
func NewCurrency(val string) *CurrencyField {
	field := &CurrencyField{}
	field.Value = val
	return field
}

//CurrencyRatioField is a FLOAT field
type CurrencyRatioField struct{ quickfix.FloatValue }

//Tag returns tag.CurrencyRatio (1382)
func (f CurrencyRatioField) Tag() quickfix.Tag { return tag.CurrencyRatio }

//NewCurrencyRatio returns a new CurrencyRatioField initialized with val
func NewCurrencyRatio(val float64) *CurrencyRatioField {
	field := &CurrencyRatioField{}
	field.Value = val
	return field
}

//CustDirectedOrderField is a BOOLEAN field
type CustDirectedOrderField struct{ quickfix.BooleanValue }

//Tag returns tag.CustDirectedOrder (1029)
func (f CustDirectedOrderField) Tag() quickfix.Tag { return tag.CustDirectedOrder }

//NewCustDirectedOrder returns a new CustDirectedOrderField initialized with val
func NewCustDirectedOrder(val bool) *CustDirectedOrderField {
	field := &CustDirectedOrderField{}
	field.Value = val
	return field
}

//CustOrderCapacityField is a INT field
type CustOrderCapacityField struct{ quickfix.IntValue }

//Tag returns tag.CustOrderCapacity (582)
func (f CustOrderCapacityField) Tag() quickfix.Tag { return tag.CustOrderCapacity }

//NewCustOrderCapacity returns a new CustOrderCapacityField initialized with val
func NewCustOrderCapacity(val int) *CustOrderCapacityField {
	field := &CustOrderCapacityField{}
	field.Value = val
	return field
}

//CustOrderHandlingInstField is a MULTIPLESTRINGVALUE field
type CustOrderHandlingInstField struct{ quickfix.MultipleStringValue }

//Tag returns tag.CustOrderHandlingInst (1031)
func (f CustOrderHandlingInstField) Tag() quickfix.Tag { return tag.CustOrderHandlingInst }

//NewCustOrderHandlingInst returns a new CustOrderHandlingInstField initialized with val
func NewCustOrderHandlingInst(val string) *CustOrderHandlingInstField {
	field := &CustOrderHandlingInstField{}
	field.Value = val
	return field
}

//CustomerOrFirmField is a INT field
type CustomerOrFirmField struct{ quickfix.IntValue }

//Tag returns tag.CustomerOrFirm (204)
func (f CustomerOrFirmField) Tag() quickfix.Tag { return tag.CustomerOrFirm }

//NewCustomerOrFirm returns a new CustomerOrFirmField initialized with val
func NewCustomerOrFirm(val int) *CustomerOrFirmField {
	field := &CustomerOrFirmField{}
	field.Value = val
	return field
}

//CxlQtyField is a QTY field
type CxlQtyField struct{ quickfix.QtyValue }

//Tag returns tag.CxlQty (84)
func (f CxlQtyField) Tag() quickfix.Tag { return tag.CxlQty }

//NewCxlQty returns a new CxlQtyField initialized with val
func NewCxlQty(val float64) *CxlQtyField {
	field := &CxlQtyField{}
	field.Value = val
	return field
}

//CxlRejReasonField is a INT field
type CxlRejReasonField struct{ quickfix.IntValue }

//Tag returns tag.CxlRejReason (102)
func (f CxlRejReasonField) Tag() quickfix.Tag { return tag.CxlRejReason }

//NewCxlRejReason returns a new CxlRejReasonField initialized with val
func NewCxlRejReason(val int) *CxlRejReasonField {
	field := &CxlRejReasonField{}
	field.Value = val
	return field
}

//CxlRejResponseToField is a CHAR field
type CxlRejResponseToField struct{ quickfix.CharValue }

//Tag returns tag.CxlRejResponseTo (434)
func (f CxlRejResponseToField) Tag() quickfix.Tag { return tag.CxlRejResponseTo }

//NewCxlRejResponseTo returns a new CxlRejResponseToField initialized with val
func NewCxlRejResponseTo(val string) *CxlRejResponseToField {
	field := &CxlRejResponseToField{}
	field.Value = val
	return field
}

//CxlTypeField is a CHAR field
type CxlTypeField struct{ quickfix.CharValue }

//Tag returns tag.CxlType (125)
func (f CxlTypeField) Tag() quickfix.Tag { return tag.CxlType }

//NewCxlType returns a new CxlTypeField initialized with val
func NewCxlType(val string) *CxlTypeField {
	field := &CxlTypeField{}
	field.Value = val
	return field
}

//DKReasonField is a CHAR field
type DKReasonField struct{ quickfix.CharValue }

//Tag returns tag.DKReason (127)
func (f DKReasonField) Tag() quickfix.Tag { return tag.DKReason }

//NewDKReason returns a new DKReasonField initialized with val
func NewDKReason(val string) *DKReasonField {
	field := &DKReasonField{}
	field.Value = val
	return field
}

//DateOfBirthField is a LOCALMKTDATE field
type DateOfBirthField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.DateOfBirth (486)
func (f DateOfBirthField) Tag() quickfix.Tag { return tag.DateOfBirth }

//NewDateOfBirth returns a new DateOfBirthField initialized with val
func NewDateOfBirth(val string) *DateOfBirthField {
	field := &DateOfBirthField{}
	field.Value = val
	return field
}

//DatedDateField is a LOCALMKTDATE field
type DatedDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.DatedDate (873)
func (f DatedDateField) Tag() quickfix.Tag { return tag.DatedDate }

//NewDatedDate returns a new DatedDateField initialized with val
func NewDatedDate(val string) *DatedDateField {
	field := &DatedDateField{}
	field.Value = val
	return field
}

//DayAvgPxField is a PRICE field
type DayAvgPxField struct{ quickfix.PriceValue }

//Tag returns tag.DayAvgPx (426)
func (f DayAvgPxField) Tag() quickfix.Tag { return tag.DayAvgPx }

//NewDayAvgPx returns a new DayAvgPxField initialized with val
func NewDayAvgPx(val float64) *DayAvgPxField {
	field := &DayAvgPxField{}
	field.Value = val
	return field
}

//DayBookingInstField is a CHAR field
type DayBookingInstField struct{ quickfix.CharValue }

//Tag returns tag.DayBookingInst (589)
func (f DayBookingInstField) Tag() quickfix.Tag { return tag.DayBookingInst }

//NewDayBookingInst returns a new DayBookingInstField initialized with val
func NewDayBookingInst(val string) *DayBookingInstField {
	field := &DayBookingInstField{}
	field.Value = val
	return field
}

//DayCumQtyField is a QTY field
type DayCumQtyField struct{ quickfix.QtyValue }

//Tag returns tag.DayCumQty (425)
func (f DayCumQtyField) Tag() quickfix.Tag { return tag.DayCumQty }

//NewDayCumQty returns a new DayCumQtyField initialized with val
func NewDayCumQty(val float64) *DayCumQtyField {
	field := &DayCumQtyField{}
	field.Value = val
	return field
}

//DayOrderQtyField is a QTY field
type DayOrderQtyField struct{ quickfix.QtyValue }

//Tag returns tag.DayOrderQty (424)
func (f DayOrderQtyField) Tag() quickfix.Tag { return tag.DayOrderQty }

//NewDayOrderQty returns a new DayOrderQtyField initialized with val
func NewDayOrderQty(val float64) *DayOrderQtyField {
	field := &DayOrderQtyField{}
	field.Value = val
	return field
}

//DealingCapacityField is a CHAR field
type DealingCapacityField struct{ quickfix.CharValue }

//Tag returns tag.DealingCapacity (1048)
func (f DealingCapacityField) Tag() quickfix.Tag { return tag.DealingCapacity }

//NewDealingCapacity returns a new DealingCapacityField initialized with val
func NewDealingCapacity(val string) *DealingCapacityField {
	field := &DealingCapacityField{}
	field.Value = val
	return field
}

//DefBidSizeField is a QTY field
type DefBidSizeField struct{ quickfix.QtyValue }

//Tag returns tag.DefBidSize (293)
func (f DefBidSizeField) Tag() quickfix.Tag { return tag.DefBidSize }

//NewDefBidSize returns a new DefBidSizeField initialized with val
func NewDefBidSize(val float64) *DefBidSizeField {
	field := &DefBidSizeField{}
	field.Value = val
	return field
}

//DefOfferSizeField is a QTY field
type DefOfferSizeField struct{ quickfix.QtyValue }

//Tag returns tag.DefOfferSize (294)
func (f DefOfferSizeField) Tag() quickfix.Tag { return tag.DefOfferSize }

//NewDefOfferSize returns a new DefOfferSizeField initialized with val
func NewDefOfferSize(val float64) *DefOfferSizeField {
	field := &DefOfferSizeField{}
	field.Value = val
	return field
}

//DefaultApplExtIDField is a INT field
type DefaultApplExtIDField struct{ quickfix.IntValue }

//Tag returns tag.DefaultApplExtID (1407)
func (f DefaultApplExtIDField) Tag() quickfix.Tag { return tag.DefaultApplExtID }

//NewDefaultApplExtID returns a new DefaultApplExtIDField initialized with val
func NewDefaultApplExtID(val int) *DefaultApplExtIDField {
	field := &DefaultApplExtIDField{}
	field.Value = val
	return field
}

//DefaultApplVerIDField is a STRING field
type DefaultApplVerIDField struct{ quickfix.StringValue }

//Tag returns tag.DefaultApplVerID (1137)
func (f DefaultApplVerIDField) Tag() quickfix.Tag { return tag.DefaultApplVerID }

//NewDefaultApplVerID returns a new DefaultApplVerIDField initialized with val
func NewDefaultApplVerID(val string) *DefaultApplVerIDField {
	field := &DefaultApplVerIDField{}
	field.Value = val
	return field
}

//DefaultCstmApplVerIDField is a STRING field
type DefaultCstmApplVerIDField struct{ quickfix.StringValue }

//Tag returns tag.DefaultCstmApplVerID (1408)
func (f DefaultCstmApplVerIDField) Tag() quickfix.Tag { return tag.DefaultCstmApplVerID }

//NewDefaultCstmApplVerID returns a new DefaultCstmApplVerIDField initialized with val
func NewDefaultCstmApplVerID(val string) *DefaultCstmApplVerIDField {
	field := &DefaultCstmApplVerIDField{}
	field.Value = val
	return field
}

//DefaultVerIndicatorField is a BOOLEAN field
type DefaultVerIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.DefaultVerIndicator (1410)
func (f DefaultVerIndicatorField) Tag() quickfix.Tag { return tag.DefaultVerIndicator }

//NewDefaultVerIndicator returns a new DefaultVerIndicatorField initialized with val
func NewDefaultVerIndicator(val bool) *DefaultVerIndicatorField {
	field := &DefaultVerIndicatorField{}
	field.Value = val
	return field
}

//DeleteReasonField is a CHAR field
type DeleteReasonField struct{ quickfix.CharValue }

//Tag returns tag.DeleteReason (285)
func (f DeleteReasonField) Tag() quickfix.Tag { return tag.DeleteReason }

//NewDeleteReason returns a new DeleteReasonField initialized with val
func NewDeleteReason(val string) *DeleteReasonField {
	field := &DeleteReasonField{}
	field.Value = val
	return field
}

//DeliverToCompIDField is a STRING field
type DeliverToCompIDField struct{ quickfix.StringValue }

//Tag returns tag.DeliverToCompID (128)
func (f DeliverToCompIDField) Tag() quickfix.Tag { return tag.DeliverToCompID }

//NewDeliverToCompID returns a new DeliverToCompIDField initialized with val
func NewDeliverToCompID(val string) *DeliverToCompIDField {
	field := &DeliverToCompIDField{}
	field.Value = val
	return field
}

//DeliverToLocationIDField is a STRING field
type DeliverToLocationIDField struct{ quickfix.StringValue }

//Tag returns tag.DeliverToLocationID (145)
func (f DeliverToLocationIDField) Tag() quickfix.Tag { return tag.DeliverToLocationID }

//NewDeliverToLocationID returns a new DeliverToLocationIDField initialized with val
func NewDeliverToLocationID(val string) *DeliverToLocationIDField {
	field := &DeliverToLocationIDField{}
	field.Value = val
	return field
}

//DeliverToSubIDField is a STRING field
type DeliverToSubIDField struct{ quickfix.StringValue }

//Tag returns tag.DeliverToSubID (129)
func (f DeliverToSubIDField) Tag() quickfix.Tag { return tag.DeliverToSubID }

//NewDeliverToSubID returns a new DeliverToSubIDField initialized with val
func NewDeliverToSubID(val string) *DeliverToSubIDField {
	field := &DeliverToSubIDField{}
	field.Value = val
	return field
}

//DeliveryDateField is a LOCALMKTDATE field
type DeliveryDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.DeliveryDate (743)
func (f DeliveryDateField) Tag() quickfix.Tag { return tag.DeliveryDate }

//NewDeliveryDate returns a new DeliveryDateField initialized with val
func NewDeliveryDate(val string) *DeliveryDateField {
	field := &DeliveryDateField{}
	field.Value = val
	return field
}

//DeliveryFormField is a INT field
type DeliveryFormField struct{ quickfix.IntValue }

//Tag returns tag.DeliveryForm (668)
func (f DeliveryFormField) Tag() quickfix.Tag { return tag.DeliveryForm }

//NewDeliveryForm returns a new DeliveryFormField initialized with val
func NewDeliveryForm(val int) *DeliveryFormField {
	field := &DeliveryFormField{}
	field.Value = val
	return field
}

//DeliveryTypeField is a INT field
type DeliveryTypeField struct{ quickfix.IntValue }

//Tag returns tag.DeliveryType (919)
func (f DeliveryTypeField) Tag() quickfix.Tag { return tag.DeliveryType }

//NewDeliveryType returns a new DeliveryTypeField initialized with val
func NewDeliveryType(val int) *DeliveryTypeField {
	field := &DeliveryTypeField{}
	field.Value = val
	return field
}

//DerivFlexProductEligibilityIndicatorField is a BOOLEAN field
type DerivFlexProductEligibilityIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.DerivFlexProductEligibilityIndicator (1243)
func (f DerivFlexProductEligibilityIndicatorField) Tag() quickfix.Tag {
	return tag.DerivFlexProductEligibilityIndicator
}

//NewDerivFlexProductEligibilityIndicator returns a new DerivFlexProductEligibilityIndicatorField initialized with val
func NewDerivFlexProductEligibilityIndicator(val bool) *DerivFlexProductEligibilityIndicatorField {
	field := &DerivFlexProductEligibilityIndicatorField{}
	field.Value = val
	return field
}

//DerivativeCFICodeField is a STRING field
type DerivativeCFICodeField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeCFICode (1248)
func (f DerivativeCFICodeField) Tag() quickfix.Tag { return tag.DerivativeCFICode }

//NewDerivativeCFICode returns a new DerivativeCFICodeField initialized with val
func NewDerivativeCFICode(val string) *DerivativeCFICodeField {
	field := &DerivativeCFICodeField{}
	field.Value = val
	return field
}

//DerivativeCapPriceField is a PRICE field
type DerivativeCapPriceField struct{ quickfix.PriceValue }

//Tag returns tag.DerivativeCapPrice (1321)
func (f DerivativeCapPriceField) Tag() quickfix.Tag { return tag.DerivativeCapPrice }

//NewDerivativeCapPrice returns a new DerivativeCapPriceField initialized with val
func NewDerivativeCapPrice(val float64) *DerivativeCapPriceField {
	field := &DerivativeCapPriceField{}
	field.Value = val
	return field
}

//DerivativeContractMultiplierField is a FLOAT field
type DerivativeContractMultiplierField struct{ quickfix.FloatValue }

//Tag returns tag.DerivativeContractMultiplier (1266)
func (f DerivativeContractMultiplierField) Tag() quickfix.Tag { return tag.DerivativeContractMultiplier }

//NewDerivativeContractMultiplier returns a new DerivativeContractMultiplierField initialized with val
func NewDerivativeContractMultiplier(val float64) *DerivativeContractMultiplierField {
	field := &DerivativeContractMultiplierField{}
	field.Value = val
	return field
}

//DerivativeContractMultiplierUnitField is a INT field
type DerivativeContractMultiplierUnitField struct{ quickfix.IntValue }

//Tag returns tag.DerivativeContractMultiplierUnit (1438)
func (f DerivativeContractMultiplierUnitField) Tag() quickfix.Tag {
	return tag.DerivativeContractMultiplierUnit
}

//NewDerivativeContractMultiplierUnit returns a new DerivativeContractMultiplierUnitField initialized with val
func NewDerivativeContractMultiplierUnit(val int) *DerivativeContractMultiplierUnitField {
	field := &DerivativeContractMultiplierUnitField{}
	field.Value = val
	return field
}

//DerivativeContractSettlMonthField is a MONTHYEAR field
type DerivativeContractSettlMonthField struct{ quickfix.MonthYearValue }

//Tag returns tag.DerivativeContractSettlMonth (1285)
func (f DerivativeContractSettlMonthField) Tag() quickfix.Tag { return tag.DerivativeContractSettlMonth }

//NewDerivativeContractSettlMonth returns a new DerivativeContractSettlMonthField initialized with val
func NewDerivativeContractSettlMonth(val string) *DerivativeContractSettlMonthField {
	field := &DerivativeContractSettlMonthField{}
	field.Value = val
	return field
}

//DerivativeCountryOfIssueField is a COUNTRY field
type DerivativeCountryOfIssueField struct{ quickfix.CountryValue }

//Tag returns tag.DerivativeCountryOfIssue (1258)
func (f DerivativeCountryOfIssueField) Tag() quickfix.Tag { return tag.DerivativeCountryOfIssue }

//NewDerivativeCountryOfIssue returns a new DerivativeCountryOfIssueField initialized with val
func NewDerivativeCountryOfIssue(val string) *DerivativeCountryOfIssueField {
	field := &DerivativeCountryOfIssueField{}
	field.Value = val
	return field
}

//DerivativeEncodedIssuerField is a DATA field
type DerivativeEncodedIssuerField struct{ quickfix.DataValue }

//Tag returns tag.DerivativeEncodedIssuer (1278)
func (f DerivativeEncodedIssuerField) Tag() quickfix.Tag { return tag.DerivativeEncodedIssuer }

//NewDerivativeEncodedIssuer returns a new DerivativeEncodedIssuerField initialized with val
func NewDerivativeEncodedIssuer(val string) *DerivativeEncodedIssuerField {
	field := &DerivativeEncodedIssuerField{}
	field.Value = val
	return field
}

//DerivativeEncodedIssuerLenField is a LENGTH field
type DerivativeEncodedIssuerLenField struct{ quickfix.LengthValue }

//Tag returns tag.DerivativeEncodedIssuerLen (1277)
func (f DerivativeEncodedIssuerLenField) Tag() quickfix.Tag { return tag.DerivativeEncodedIssuerLen }

//NewDerivativeEncodedIssuerLen returns a new DerivativeEncodedIssuerLenField initialized with val
func NewDerivativeEncodedIssuerLen(val int) *DerivativeEncodedIssuerLenField {
	field := &DerivativeEncodedIssuerLenField{}
	field.Value = val
	return field
}

//DerivativeEncodedSecurityDescField is a DATA field
type DerivativeEncodedSecurityDescField struct{ quickfix.DataValue }

//Tag returns tag.DerivativeEncodedSecurityDesc (1281)
func (f DerivativeEncodedSecurityDescField) Tag() quickfix.Tag {
	return tag.DerivativeEncodedSecurityDesc
}

//NewDerivativeEncodedSecurityDesc returns a new DerivativeEncodedSecurityDescField initialized with val
func NewDerivativeEncodedSecurityDesc(val string) *DerivativeEncodedSecurityDescField {
	field := &DerivativeEncodedSecurityDescField{}
	field.Value = val
	return field
}

//DerivativeEncodedSecurityDescLenField is a LENGTH field
type DerivativeEncodedSecurityDescLenField struct{ quickfix.LengthValue }

//Tag returns tag.DerivativeEncodedSecurityDescLen (1280)
func (f DerivativeEncodedSecurityDescLenField) Tag() quickfix.Tag {
	return tag.DerivativeEncodedSecurityDescLen
}

//NewDerivativeEncodedSecurityDescLen returns a new DerivativeEncodedSecurityDescLenField initialized with val
func NewDerivativeEncodedSecurityDescLen(val int) *DerivativeEncodedSecurityDescLenField {
	field := &DerivativeEncodedSecurityDescLenField{}
	field.Value = val
	return field
}

//DerivativeEventDateField is a LOCALMKTDATE field
type DerivativeEventDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.DerivativeEventDate (1288)
func (f DerivativeEventDateField) Tag() quickfix.Tag { return tag.DerivativeEventDate }

//NewDerivativeEventDate returns a new DerivativeEventDateField initialized with val
func NewDerivativeEventDate(val string) *DerivativeEventDateField {
	field := &DerivativeEventDateField{}
	field.Value = val
	return field
}

//DerivativeEventPxField is a PRICE field
type DerivativeEventPxField struct{ quickfix.PriceValue }

//Tag returns tag.DerivativeEventPx (1290)
func (f DerivativeEventPxField) Tag() quickfix.Tag { return tag.DerivativeEventPx }

//NewDerivativeEventPx returns a new DerivativeEventPxField initialized with val
func NewDerivativeEventPx(val float64) *DerivativeEventPxField {
	field := &DerivativeEventPxField{}
	field.Value = val
	return field
}

//DerivativeEventTextField is a STRING field
type DerivativeEventTextField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeEventText (1291)
func (f DerivativeEventTextField) Tag() quickfix.Tag { return tag.DerivativeEventText }

//NewDerivativeEventText returns a new DerivativeEventTextField initialized with val
func NewDerivativeEventText(val string) *DerivativeEventTextField {
	field := &DerivativeEventTextField{}
	field.Value = val
	return field
}

//DerivativeEventTimeField is a UTCTIMESTAMP field
type DerivativeEventTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.DerivativeEventTime (1289)
func (f DerivativeEventTimeField) Tag() quickfix.Tag { return tag.DerivativeEventTime }

//DerivativeEventTypeField is a INT field
type DerivativeEventTypeField struct{ quickfix.IntValue }

//Tag returns tag.DerivativeEventType (1287)
func (f DerivativeEventTypeField) Tag() quickfix.Tag { return tag.DerivativeEventType }

//NewDerivativeEventType returns a new DerivativeEventTypeField initialized with val
func NewDerivativeEventType(val int) *DerivativeEventTypeField {
	field := &DerivativeEventTypeField{}
	field.Value = val
	return field
}

//DerivativeExerciseStyleField is a CHAR field
type DerivativeExerciseStyleField struct{ quickfix.CharValue }

//Tag returns tag.DerivativeExerciseStyle (1299)
func (f DerivativeExerciseStyleField) Tag() quickfix.Tag { return tag.DerivativeExerciseStyle }

//NewDerivativeExerciseStyle returns a new DerivativeExerciseStyleField initialized with val
func NewDerivativeExerciseStyle(val string) *DerivativeExerciseStyleField {
	field := &DerivativeExerciseStyleField{}
	field.Value = val
	return field
}

//DerivativeFloorPriceField is a PRICE field
type DerivativeFloorPriceField struct{ quickfix.PriceValue }

//Tag returns tag.DerivativeFloorPrice (1322)
func (f DerivativeFloorPriceField) Tag() quickfix.Tag { return tag.DerivativeFloorPrice }

//NewDerivativeFloorPrice returns a new DerivativeFloorPriceField initialized with val
func NewDerivativeFloorPrice(val float64) *DerivativeFloorPriceField {
	field := &DerivativeFloorPriceField{}
	field.Value = val
	return field
}

//DerivativeFlowScheduleTypeField is a INT field
type DerivativeFlowScheduleTypeField struct{ quickfix.IntValue }

//Tag returns tag.DerivativeFlowScheduleType (1442)
func (f DerivativeFlowScheduleTypeField) Tag() quickfix.Tag { return tag.DerivativeFlowScheduleType }

//NewDerivativeFlowScheduleType returns a new DerivativeFlowScheduleTypeField initialized with val
func NewDerivativeFlowScheduleType(val int) *DerivativeFlowScheduleTypeField {
	field := &DerivativeFlowScheduleTypeField{}
	field.Value = val
	return field
}

//DerivativeFuturesValuationMethodField is a STRING field
type DerivativeFuturesValuationMethodField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeFuturesValuationMethod (1319)
func (f DerivativeFuturesValuationMethodField) Tag() quickfix.Tag {
	return tag.DerivativeFuturesValuationMethod
}

//NewDerivativeFuturesValuationMethod returns a new DerivativeFuturesValuationMethodField initialized with val
func NewDerivativeFuturesValuationMethod(val string) *DerivativeFuturesValuationMethodField {
	field := &DerivativeFuturesValuationMethodField{}
	field.Value = val
	return field
}

//DerivativeInstrAttribTypeField is a INT field
type DerivativeInstrAttribTypeField struct{ quickfix.IntValue }

//Tag returns tag.DerivativeInstrAttribType (1313)
func (f DerivativeInstrAttribTypeField) Tag() quickfix.Tag { return tag.DerivativeInstrAttribType }

//NewDerivativeInstrAttribType returns a new DerivativeInstrAttribTypeField initialized with val
func NewDerivativeInstrAttribType(val int) *DerivativeInstrAttribTypeField {
	field := &DerivativeInstrAttribTypeField{}
	field.Value = val
	return field
}

//DerivativeInstrAttribValueField is a STRING field
type DerivativeInstrAttribValueField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeInstrAttribValue (1314)
func (f DerivativeInstrAttribValueField) Tag() quickfix.Tag { return tag.DerivativeInstrAttribValue }

//NewDerivativeInstrAttribValue returns a new DerivativeInstrAttribValueField initialized with val
func NewDerivativeInstrAttribValue(val string) *DerivativeInstrAttribValueField {
	field := &DerivativeInstrAttribValueField{}
	field.Value = val
	return field
}

//DerivativeInstrRegistryField is a STRING field
type DerivativeInstrRegistryField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeInstrRegistry (1257)
func (f DerivativeInstrRegistryField) Tag() quickfix.Tag { return tag.DerivativeInstrRegistry }

//NewDerivativeInstrRegistry returns a new DerivativeInstrRegistryField initialized with val
func NewDerivativeInstrRegistry(val string) *DerivativeInstrRegistryField {
	field := &DerivativeInstrRegistryField{}
	field.Value = val
	return field
}

//DerivativeInstrmtAssignmentMethodField is a CHAR field
type DerivativeInstrmtAssignmentMethodField struct{ quickfix.CharValue }

//Tag returns tag.DerivativeInstrmtAssignmentMethod (1255)
func (f DerivativeInstrmtAssignmentMethodField) Tag() quickfix.Tag {
	return tag.DerivativeInstrmtAssignmentMethod
}

//NewDerivativeInstrmtAssignmentMethod returns a new DerivativeInstrmtAssignmentMethodField initialized with val
func NewDerivativeInstrmtAssignmentMethod(val string) *DerivativeInstrmtAssignmentMethodField {
	field := &DerivativeInstrmtAssignmentMethodField{}
	field.Value = val
	return field
}

//DerivativeInstrumentPartyIDField is a STRING field
type DerivativeInstrumentPartyIDField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeInstrumentPartyID (1293)
func (f DerivativeInstrumentPartyIDField) Tag() quickfix.Tag { return tag.DerivativeInstrumentPartyID }

//NewDerivativeInstrumentPartyID returns a new DerivativeInstrumentPartyIDField initialized with val
func NewDerivativeInstrumentPartyID(val string) *DerivativeInstrumentPartyIDField {
	field := &DerivativeInstrumentPartyIDField{}
	field.Value = val
	return field
}

//DerivativeInstrumentPartyIDSourceField is a STRING field
type DerivativeInstrumentPartyIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeInstrumentPartyIDSource (1294)
func (f DerivativeInstrumentPartyIDSourceField) Tag() quickfix.Tag {
	return tag.DerivativeInstrumentPartyIDSource
}

//NewDerivativeInstrumentPartyIDSource returns a new DerivativeInstrumentPartyIDSourceField initialized with val
func NewDerivativeInstrumentPartyIDSource(val string) *DerivativeInstrumentPartyIDSourceField {
	field := &DerivativeInstrumentPartyIDSourceField{}
	field.Value = val
	return field
}

//DerivativeInstrumentPartyRoleField is a INT field
type DerivativeInstrumentPartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.DerivativeInstrumentPartyRole (1295)
func (f DerivativeInstrumentPartyRoleField) Tag() quickfix.Tag {
	return tag.DerivativeInstrumentPartyRole
}

//NewDerivativeInstrumentPartyRole returns a new DerivativeInstrumentPartyRoleField initialized with val
func NewDerivativeInstrumentPartyRole(val int) *DerivativeInstrumentPartyRoleField {
	field := &DerivativeInstrumentPartyRoleField{}
	field.Value = val
	return field
}

//DerivativeInstrumentPartySubIDField is a STRING field
type DerivativeInstrumentPartySubIDField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeInstrumentPartySubID (1297)
func (f DerivativeInstrumentPartySubIDField) Tag() quickfix.Tag {
	return tag.DerivativeInstrumentPartySubID
}

//NewDerivativeInstrumentPartySubID returns a new DerivativeInstrumentPartySubIDField initialized with val
func NewDerivativeInstrumentPartySubID(val string) *DerivativeInstrumentPartySubIDField {
	field := &DerivativeInstrumentPartySubIDField{}
	field.Value = val
	return field
}

//DerivativeInstrumentPartySubIDTypeField is a INT field
type DerivativeInstrumentPartySubIDTypeField struct{ quickfix.IntValue }

//Tag returns tag.DerivativeInstrumentPartySubIDType (1298)
func (f DerivativeInstrumentPartySubIDTypeField) Tag() quickfix.Tag {
	return tag.DerivativeInstrumentPartySubIDType
}

//NewDerivativeInstrumentPartySubIDType returns a new DerivativeInstrumentPartySubIDTypeField initialized with val
func NewDerivativeInstrumentPartySubIDType(val int) *DerivativeInstrumentPartySubIDTypeField {
	field := &DerivativeInstrumentPartySubIDTypeField{}
	field.Value = val
	return field
}

//DerivativeIssueDateField is a LOCALMKTDATE field
type DerivativeIssueDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.DerivativeIssueDate (1276)
func (f DerivativeIssueDateField) Tag() quickfix.Tag { return tag.DerivativeIssueDate }

//NewDerivativeIssueDate returns a new DerivativeIssueDateField initialized with val
func NewDerivativeIssueDate(val string) *DerivativeIssueDateField {
	field := &DerivativeIssueDateField{}
	field.Value = val
	return field
}

//DerivativeIssuerField is a STRING field
type DerivativeIssuerField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeIssuer (1275)
func (f DerivativeIssuerField) Tag() quickfix.Tag { return tag.DerivativeIssuer }

//NewDerivativeIssuer returns a new DerivativeIssuerField initialized with val
func NewDerivativeIssuer(val string) *DerivativeIssuerField {
	field := &DerivativeIssuerField{}
	field.Value = val
	return field
}

//DerivativeListMethodField is a INT field
type DerivativeListMethodField struct{ quickfix.IntValue }

//Tag returns tag.DerivativeListMethod (1320)
func (f DerivativeListMethodField) Tag() quickfix.Tag { return tag.DerivativeListMethod }

//NewDerivativeListMethod returns a new DerivativeListMethodField initialized with val
func NewDerivativeListMethod(val int) *DerivativeListMethodField {
	field := &DerivativeListMethodField{}
	field.Value = val
	return field
}

//DerivativeLocaleOfIssueField is a STRING field
type DerivativeLocaleOfIssueField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeLocaleOfIssue (1260)
func (f DerivativeLocaleOfIssueField) Tag() quickfix.Tag { return tag.DerivativeLocaleOfIssue }

//NewDerivativeLocaleOfIssue returns a new DerivativeLocaleOfIssueField initialized with val
func NewDerivativeLocaleOfIssue(val string) *DerivativeLocaleOfIssueField {
	field := &DerivativeLocaleOfIssueField{}
	field.Value = val
	return field
}

//DerivativeMaturityDateField is a LOCALMKTDATE field
type DerivativeMaturityDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.DerivativeMaturityDate (1252)
func (f DerivativeMaturityDateField) Tag() quickfix.Tag { return tag.DerivativeMaturityDate }

//NewDerivativeMaturityDate returns a new DerivativeMaturityDateField initialized with val
func NewDerivativeMaturityDate(val string) *DerivativeMaturityDateField {
	field := &DerivativeMaturityDateField{}
	field.Value = val
	return field
}

//DerivativeMaturityMonthYearField is a MONTHYEAR field
type DerivativeMaturityMonthYearField struct{ quickfix.MonthYearValue }

//Tag returns tag.DerivativeMaturityMonthYear (1251)
func (f DerivativeMaturityMonthYearField) Tag() quickfix.Tag { return tag.DerivativeMaturityMonthYear }

//NewDerivativeMaturityMonthYear returns a new DerivativeMaturityMonthYearField initialized with val
func NewDerivativeMaturityMonthYear(val string) *DerivativeMaturityMonthYearField {
	field := &DerivativeMaturityMonthYearField{}
	field.Value = val
	return field
}

//DerivativeMaturityTimeField is a TZTIMEONLY field
type DerivativeMaturityTimeField struct{ quickfix.TZTimeOnlyValue }

//Tag returns tag.DerivativeMaturityTime (1253)
func (f DerivativeMaturityTimeField) Tag() quickfix.Tag { return tag.DerivativeMaturityTime }

//DerivativeMinPriceIncrementField is a FLOAT field
type DerivativeMinPriceIncrementField struct{ quickfix.FloatValue }

//Tag returns tag.DerivativeMinPriceIncrement (1267)
func (f DerivativeMinPriceIncrementField) Tag() quickfix.Tag { return tag.DerivativeMinPriceIncrement }

//NewDerivativeMinPriceIncrement returns a new DerivativeMinPriceIncrementField initialized with val
func NewDerivativeMinPriceIncrement(val float64) *DerivativeMinPriceIncrementField {
	field := &DerivativeMinPriceIncrementField{}
	field.Value = val
	return field
}

//DerivativeMinPriceIncrementAmountField is a AMT field
type DerivativeMinPriceIncrementAmountField struct{ quickfix.AmtValue }

//Tag returns tag.DerivativeMinPriceIncrementAmount (1268)
func (f DerivativeMinPriceIncrementAmountField) Tag() quickfix.Tag {
	return tag.DerivativeMinPriceIncrementAmount
}

//NewDerivativeMinPriceIncrementAmount returns a new DerivativeMinPriceIncrementAmountField initialized with val
func NewDerivativeMinPriceIncrementAmount(val float64) *DerivativeMinPriceIncrementAmountField {
	field := &DerivativeMinPriceIncrementAmountField{}
	field.Value = val
	return field
}

//DerivativeNTPositionLimitField is a INT field
type DerivativeNTPositionLimitField struct{ quickfix.IntValue }

//Tag returns tag.DerivativeNTPositionLimit (1274)
func (f DerivativeNTPositionLimitField) Tag() quickfix.Tag { return tag.DerivativeNTPositionLimit }

//NewDerivativeNTPositionLimit returns a new DerivativeNTPositionLimitField initialized with val
func NewDerivativeNTPositionLimit(val int) *DerivativeNTPositionLimitField {
	field := &DerivativeNTPositionLimitField{}
	field.Value = val
	return field
}

//DerivativeOptAttributeField is a CHAR field
type DerivativeOptAttributeField struct{ quickfix.CharValue }

//Tag returns tag.DerivativeOptAttribute (1265)
func (f DerivativeOptAttributeField) Tag() quickfix.Tag { return tag.DerivativeOptAttribute }

//NewDerivativeOptAttribute returns a new DerivativeOptAttributeField initialized with val
func NewDerivativeOptAttribute(val string) *DerivativeOptAttributeField {
	field := &DerivativeOptAttributeField{}
	field.Value = val
	return field
}

//DerivativeOptPayAmountField is a AMT field
type DerivativeOptPayAmountField struct{ quickfix.AmtValue }

//Tag returns tag.DerivativeOptPayAmount (1225)
func (f DerivativeOptPayAmountField) Tag() quickfix.Tag { return tag.DerivativeOptPayAmount }

//NewDerivativeOptPayAmount returns a new DerivativeOptPayAmountField initialized with val
func NewDerivativeOptPayAmount(val float64) *DerivativeOptPayAmountField {
	field := &DerivativeOptPayAmountField{}
	field.Value = val
	return field
}

//DerivativePositionLimitField is a INT field
type DerivativePositionLimitField struct{ quickfix.IntValue }

//Tag returns tag.DerivativePositionLimit (1273)
func (f DerivativePositionLimitField) Tag() quickfix.Tag { return tag.DerivativePositionLimit }

//NewDerivativePositionLimit returns a new DerivativePositionLimitField initialized with val
func NewDerivativePositionLimit(val int) *DerivativePositionLimitField {
	field := &DerivativePositionLimitField{}
	field.Value = val
	return field
}

//DerivativePriceQuoteMethodField is a STRING field
type DerivativePriceQuoteMethodField struct{ quickfix.StringValue }

//Tag returns tag.DerivativePriceQuoteMethod (1318)
func (f DerivativePriceQuoteMethodField) Tag() quickfix.Tag { return tag.DerivativePriceQuoteMethod }

//NewDerivativePriceQuoteMethod returns a new DerivativePriceQuoteMethodField initialized with val
func NewDerivativePriceQuoteMethod(val string) *DerivativePriceQuoteMethodField {
	field := &DerivativePriceQuoteMethodField{}
	field.Value = val
	return field
}

//DerivativePriceUnitOfMeasureField is a STRING field
type DerivativePriceUnitOfMeasureField struct{ quickfix.StringValue }

//Tag returns tag.DerivativePriceUnitOfMeasure (1315)
func (f DerivativePriceUnitOfMeasureField) Tag() quickfix.Tag { return tag.DerivativePriceUnitOfMeasure }

//NewDerivativePriceUnitOfMeasure returns a new DerivativePriceUnitOfMeasureField initialized with val
func NewDerivativePriceUnitOfMeasure(val string) *DerivativePriceUnitOfMeasureField {
	field := &DerivativePriceUnitOfMeasureField{}
	field.Value = val
	return field
}

//DerivativePriceUnitOfMeasureQtyField is a QTY field
type DerivativePriceUnitOfMeasureQtyField struct{ quickfix.QtyValue }

//Tag returns tag.DerivativePriceUnitOfMeasureQty (1316)
func (f DerivativePriceUnitOfMeasureQtyField) Tag() quickfix.Tag {
	return tag.DerivativePriceUnitOfMeasureQty
}

//NewDerivativePriceUnitOfMeasureQty returns a new DerivativePriceUnitOfMeasureQtyField initialized with val
func NewDerivativePriceUnitOfMeasureQty(val float64) *DerivativePriceUnitOfMeasureQtyField {
	field := &DerivativePriceUnitOfMeasureQtyField{}
	field.Value = val
	return field
}

//DerivativeProductField is a INT field
type DerivativeProductField struct{ quickfix.IntValue }

//Tag returns tag.DerivativeProduct (1246)
func (f DerivativeProductField) Tag() quickfix.Tag { return tag.DerivativeProduct }

//NewDerivativeProduct returns a new DerivativeProductField initialized with val
func NewDerivativeProduct(val int) *DerivativeProductField {
	field := &DerivativeProductField{}
	field.Value = val
	return field
}

//DerivativeProductComplexField is a STRING field
type DerivativeProductComplexField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeProductComplex (1228)
func (f DerivativeProductComplexField) Tag() quickfix.Tag { return tag.DerivativeProductComplex }

//NewDerivativeProductComplex returns a new DerivativeProductComplexField initialized with val
func NewDerivativeProductComplex(val string) *DerivativeProductComplexField {
	field := &DerivativeProductComplexField{}
	field.Value = val
	return field
}

//DerivativePutOrCallField is a INT field
type DerivativePutOrCallField struct{ quickfix.IntValue }

//Tag returns tag.DerivativePutOrCall (1323)
func (f DerivativePutOrCallField) Tag() quickfix.Tag { return tag.DerivativePutOrCall }

//NewDerivativePutOrCall returns a new DerivativePutOrCallField initialized with val
func NewDerivativePutOrCall(val int) *DerivativePutOrCallField {
	field := &DerivativePutOrCallField{}
	field.Value = val
	return field
}

//DerivativeSecurityAltIDField is a STRING field
type DerivativeSecurityAltIDField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeSecurityAltID (1219)
func (f DerivativeSecurityAltIDField) Tag() quickfix.Tag { return tag.DerivativeSecurityAltID }

//NewDerivativeSecurityAltID returns a new DerivativeSecurityAltIDField initialized with val
func NewDerivativeSecurityAltID(val string) *DerivativeSecurityAltIDField {
	field := &DerivativeSecurityAltIDField{}
	field.Value = val
	return field
}

//DerivativeSecurityAltIDSourceField is a STRING field
type DerivativeSecurityAltIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeSecurityAltIDSource (1220)
func (f DerivativeSecurityAltIDSourceField) Tag() quickfix.Tag {
	return tag.DerivativeSecurityAltIDSource
}

//NewDerivativeSecurityAltIDSource returns a new DerivativeSecurityAltIDSourceField initialized with val
func NewDerivativeSecurityAltIDSource(val string) *DerivativeSecurityAltIDSourceField {
	field := &DerivativeSecurityAltIDSourceField{}
	field.Value = val
	return field
}

//DerivativeSecurityDescField is a STRING field
type DerivativeSecurityDescField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeSecurityDesc (1279)
func (f DerivativeSecurityDescField) Tag() quickfix.Tag { return tag.DerivativeSecurityDesc }

//NewDerivativeSecurityDesc returns a new DerivativeSecurityDescField initialized with val
func NewDerivativeSecurityDesc(val string) *DerivativeSecurityDescField {
	field := &DerivativeSecurityDescField{}
	field.Value = val
	return field
}

//DerivativeSecurityExchangeField is a EXCHANGE field
type DerivativeSecurityExchangeField struct{ quickfix.ExchangeValue }

//Tag returns tag.DerivativeSecurityExchange (1272)
func (f DerivativeSecurityExchangeField) Tag() quickfix.Tag { return tag.DerivativeSecurityExchange }

//NewDerivativeSecurityExchange returns a new DerivativeSecurityExchangeField initialized with val
func NewDerivativeSecurityExchange(val string) *DerivativeSecurityExchangeField {
	field := &DerivativeSecurityExchangeField{}
	field.Value = val
	return field
}

//DerivativeSecurityGroupField is a STRING field
type DerivativeSecurityGroupField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeSecurityGroup (1247)
func (f DerivativeSecurityGroupField) Tag() quickfix.Tag { return tag.DerivativeSecurityGroup }

//NewDerivativeSecurityGroup returns a new DerivativeSecurityGroupField initialized with val
func NewDerivativeSecurityGroup(val string) *DerivativeSecurityGroupField {
	field := &DerivativeSecurityGroupField{}
	field.Value = val
	return field
}

//DerivativeSecurityIDField is a STRING field
type DerivativeSecurityIDField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeSecurityID (1216)
func (f DerivativeSecurityIDField) Tag() quickfix.Tag { return tag.DerivativeSecurityID }

//NewDerivativeSecurityID returns a new DerivativeSecurityIDField initialized with val
func NewDerivativeSecurityID(val string) *DerivativeSecurityIDField {
	field := &DerivativeSecurityIDField{}
	field.Value = val
	return field
}

//DerivativeSecurityIDSourceField is a STRING field
type DerivativeSecurityIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeSecurityIDSource (1217)
func (f DerivativeSecurityIDSourceField) Tag() quickfix.Tag { return tag.DerivativeSecurityIDSource }

//NewDerivativeSecurityIDSource returns a new DerivativeSecurityIDSourceField initialized with val
func NewDerivativeSecurityIDSource(val string) *DerivativeSecurityIDSourceField {
	field := &DerivativeSecurityIDSourceField{}
	field.Value = val
	return field
}

//DerivativeSecurityListRequestTypeField is a INT field
type DerivativeSecurityListRequestTypeField struct{ quickfix.IntValue }

//Tag returns tag.DerivativeSecurityListRequestType (1307)
func (f DerivativeSecurityListRequestTypeField) Tag() quickfix.Tag {
	return tag.DerivativeSecurityListRequestType
}

//NewDerivativeSecurityListRequestType returns a new DerivativeSecurityListRequestTypeField initialized with val
func NewDerivativeSecurityListRequestType(val int) *DerivativeSecurityListRequestTypeField {
	field := &DerivativeSecurityListRequestTypeField{}
	field.Value = val
	return field
}

//DerivativeSecurityStatusField is a STRING field
type DerivativeSecurityStatusField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeSecurityStatus (1256)
func (f DerivativeSecurityStatusField) Tag() quickfix.Tag { return tag.DerivativeSecurityStatus }

//NewDerivativeSecurityStatus returns a new DerivativeSecurityStatusField initialized with val
func NewDerivativeSecurityStatus(val string) *DerivativeSecurityStatusField {
	field := &DerivativeSecurityStatusField{}
	field.Value = val
	return field
}

//DerivativeSecuritySubTypeField is a STRING field
type DerivativeSecuritySubTypeField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeSecuritySubType (1250)
func (f DerivativeSecuritySubTypeField) Tag() quickfix.Tag { return tag.DerivativeSecuritySubType }

//NewDerivativeSecuritySubType returns a new DerivativeSecuritySubTypeField initialized with val
func NewDerivativeSecuritySubType(val string) *DerivativeSecuritySubTypeField {
	field := &DerivativeSecuritySubTypeField{}
	field.Value = val
	return field
}

//DerivativeSecurityTypeField is a STRING field
type DerivativeSecurityTypeField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeSecurityType (1249)
func (f DerivativeSecurityTypeField) Tag() quickfix.Tag { return tag.DerivativeSecurityType }

//NewDerivativeSecurityType returns a new DerivativeSecurityTypeField initialized with val
func NewDerivativeSecurityType(val string) *DerivativeSecurityTypeField {
	field := &DerivativeSecurityTypeField{}
	field.Value = val
	return field
}

//DerivativeSecurityXMLField is a DATA field
type DerivativeSecurityXMLField struct{ quickfix.DataValue }

//Tag returns tag.DerivativeSecurityXML (1283)
func (f DerivativeSecurityXMLField) Tag() quickfix.Tag { return tag.DerivativeSecurityXML }

//NewDerivativeSecurityXML returns a new DerivativeSecurityXMLField initialized with val
func NewDerivativeSecurityXML(val string) *DerivativeSecurityXMLField {
	field := &DerivativeSecurityXMLField{}
	field.Value = val
	return field
}

//DerivativeSecurityXMLLenField is a LENGTH field
type DerivativeSecurityXMLLenField struct{ quickfix.LengthValue }

//Tag returns tag.DerivativeSecurityXMLLen (1282)
func (f DerivativeSecurityXMLLenField) Tag() quickfix.Tag { return tag.DerivativeSecurityXMLLen }

//NewDerivativeSecurityXMLLen returns a new DerivativeSecurityXMLLenField initialized with val
func NewDerivativeSecurityXMLLen(val int) *DerivativeSecurityXMLLenField {
	field := &DerivativeSecurityXMLLenField{}
	field.Value = val
	return field
}

//DerivativeSecurityXMLSchemaField is a STRING field
type DerivativeSecurityXMLSchemaField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeSecurityXMLSchema (1284)
func (f DerivativeSecurityXMLSchemaField) Tag() quickfix.Tag { return tag.DerivativeSecurityXMLSchema }

//NewDerivativeSecurityXMLSchema returns a new DerivativeSecurityXMLSchemaField initialized with val
func NewDerivativeSecurityXMLSchema(val string) *DerivativeSecurityXMLSchemaField {
	field := &DerivativeSecurityXMLSchemaField{}
	field.Value = val
	return field
}

//DerivativeSettlMethodField is a CHAR field
type DerivativeSettlMethodField struct{ quickfix.CharValue }

//Tag returns tag.DerivativeSettlMethod (1317)
func (f DerivativeSettlMethodField) Tag() quickfix.Tag { return tag.DerivativeSettlMethod }

//NewDerivativeSettlMethod returns a new DerivativeSettlMethodField initialized with val
func NewDerivativeSettlMethod(val string) *DerivativeSettlMethodField {
	field := &DerivativeSettlMethodField{}
	field.Value = val
	return field
}

//DerivativeSettleOnOpenFlagField is a STRING field
type DerivativeSettleOnOpenFlagField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeSettleOnOpenFlag (1254)
func (f DerivativeSettleOnOpenFlagField) Tag() quickfix.Tag { return tag.DerivativeSettleOnOpenFlag }

//NewDerivativeSettleOnOpenFlag returns a new DerivativeSettleOnOpenFlagField initialized with val
func NewDerivativeSettleOnOpenFlag(val string) *DerivativeSettleOnOpenFlagField {
	field := &DerivativeSettleOnOpenFlagField{}
	field.Value = val
	return field
}

//DerivativeStateOrProvinceOfIssueField is a STRING field
type DerivativeStateOrProvinceOfIssueField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeStateOrProvinceOfIssue (1259)
func (f DerivativeStateOrProvinceOfIssueField) Tag() quickfix.Tag {
	return tag.DerivativeStateOrProvinceOfIssue
}

//NewDerivativeStateOrProvinceOfIssue returns a new DerivativeStateOrProvinceOfIssueField initialized with val
func NewDerivativeStateOrProvinceOfIssue(val string) *DerivativeStateOrProvinceOfIssueField {
	field := &DerivativeStateOrProvinceOfIssueField{}
	field.Value = val
	return field
}

//DerivativeStrikeCurrencyField is a CURRENCY field
type DerivativeStrikeCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.DerivativeStrikeCurrency (1262)
func (f DerivativeStrikeCurrencyField) Tag() quickfix.Tag { return tag.DerivativeStrikeCurrency }

//NewDerivativeStrikeCurrency returns a new DerivativeStrikeCurrencyField initialized with val
func NewDerivativeStrikeCurrency(val string) *DerivativeStrikeCurrencyField {
	field := &DerivativeStrikeCurrencyField{}
	field.Value = val
	return field
}

//DerivativeStrikeMultiplierField is a FLOAT field
type DerivativeStrikeMultiplierField struct{ quickfix.FloatValue }

//Tag returns tag.DerivativeStrikeMultiplier (1263)
func (f DerivativeStrikeMultiplierField) Tag() quickfix.Tag { return tag.DerivativeStrikeMultiplier }

//NewDerivativeStrikeMultiplier returns a new DerivativeStrikeMultiplierField initialized with val
func NewDerivativeStrikeMultiplier(val float64) *DerivativeStrikeMultiplierField {
	field := &DerivativeStrikeMultiplierField{}
	field.Value = val
	return field
}

//DerivativeStrikePriceField is a PRICE field
type DerivativeStrikePriceField struct{ quickfix.PriceValue }

//Tag returns tag.DerivativeStrikePrice (1261)
func (f DerivativeStrikePriceField) Tag() quickfix.Tag { return tag.DerivativeStrikePrice }

//NewDerivativeStrikePrice returns a new DerivativeStrikePriceField initialized with val
func NewDerivativeStrikePrice(val float64) *DerivativeStrikePriceField {
	field := &DerivativeStrikePriceField{}
	field.Value = val
	return field
}

//DerivativeStrikeValueField is a FLOAT field
type DerivativeStrikeValueField struct{ quickfix.FloatValue }

//Tag returns tag.DerivativeStrikeValue (1264)
func (f DerivativeStrikeValueField) Tag() quickfix.Tag { return tag.DerivativeStrikeValue }

//NewDerivativeStrikeValue returns a new DerivativeStrikeValueField initialized with val
func NewDerivativeStrikeValue(val float64) *DerivativeStrikeValueField {
	field := &DerivativeStrikeValueField{}
	field.Value = val
	return field
}

//DerivativeSymbolField is a STRING field
type DerivativeSymbolField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeSymbol (1214)
func (f DerivativeSymbolField) Tag() quickfix.Tag { return tag.DerivativeSymbol }

//NewDerivativeSymbol returns a new DerivativeSymbolField initialized with val
func NewDerivativeSymbol(val string) *DerivativeSymbolField {
	field := &DerivativeSymbolField{}
	field.Value = val
	return field
}

//DerivativeSymbolSfxField is a STRING field
type DerivativeSymbolSfxField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeSymbolSfx (1215)
func (f DerivativeSymbolSfxField) Tag() quickfix.Tag { return tag.DerivativeSymbolSfx }

//NewDerivativeSymbolSfx returns a new DerivativeSymbolSfxField initialized with val
func NewDerivativeSymbolSfx(val string) *DerivativeSymbolSfxField {
	field := &DerivativeSymbolSfxField{}
	field.Value = val
	return field
}

//DerivativeTimeUnitField is a STRING field
type DerivativeTimeUnitField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeTimeUnit (1271)
func (f DerivativeTimeUnitField) Tag() quickfix.Tag { return tag.DerivativeTimeUnit }

//NewDerivativeTimeUnit returns a new DerivativeTimeUnitField initialized with val
func NewDerivativeTimeUnit(val string) *DerivativeTimeUnitField {
	field := &DerivativeTimeUnitField{}
	field.Value = val
	return field
}

//DerivativeUnitOfMeasureField is a STRING field
type DerivativeUnitOfMeasureField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeUnitOfMeasure (1269)
func (f DerivativeUnitOfMeasureField) Tag() quickfix.Tag { return tag.DerivativeUnitOfMeasure }

//NewDerivativeUnitOfMeasure returns a new DerivativeUnitOfMeasureField initialized with val
func NewDerivativeUnitOfMeasure(val string) *DerivativeUnitOfMeasureField {
	field := &DerivativeUnitOfMeasureField{}
	field.Value = val
	return field
}

//DerivativeUnitOfMeasureQtyField is a QTY field
type DerivativeUnitOfMeasureQtyField struct{ quickfix.QtyValue }

//Tag returns tag.DerivativeUnitOfMeasureQty (1270)
func (f DerivativeUnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.DerivativeUnitOfMeasureQty }

//NewDerivativeUnitOfMeasureQty returns a new DerivativeUnitOfMeasureQtyField initialized with val
func NewDerivativeUnitOfMeasureQty(val float64) *DerivativeUnitOfMeasureQtyField {
	field := &DerivativeUnitOfMeasureQtyField{}
	field.Value = val
	return field
}

//DerivativeValuationMethodField is a STRING field
type DerivativeValuationMethodField struct{ quickfix.StringValue }

//Tag returns tag.DerivativeValuationMethod (1319)
func (f DerivativeValuationMethodField) Tag() quickfix.Tag { return tag.DerivativeValuationMethod }

//NewDerivativeValuationMethod returns a new DerivativeValuationMethodField initialized with val
func NewDerivativeValuationMethod(val string) *DerivativeValuationMethodField {
	field := &DerivativeValuationMethodField{}
	field.Value = val
	return field
}

//DesignationField is a STRING field
type DesignationField struct{ quickfix.StringValue }

//Tag returns tag.Designation (494)
func (f DesignationField) Tag() quickfix.Tag { return tag.Designation }

//NewDesignation returns a new DesignationField initialized with val
func NewDesignation(val string) *DesignationField {
	field := &DesignationField{}
	field.Value = val
	return field
}

//DeskIDField is a STRING field
type DeskIDField struct{ quickfix.StringValue }

//Tag returns tag.DeskID (284)
func (f DeskIDField) Tag() quickfix.Tag { return tag.DeskID }

//NewDeskID returns a new DeskIDField initialized with val
func NewDeskID(val string) *DeskIDField {
	field := &DeskIDField{}
	field.Value = val
	return field
}

//DeskOrderHandlingInstField is a MULTIPLESTRINGVALUE field
type DeskOrderHandlingInstField struct{ quickfix.MultipleStringValue }

//Tag returns tag.DeskOrderHandlingInst (1035)
func (f DeskOrderHandlingInstField) Tag() quickfix.Tag { return tag.DeskOrderHandlingInst }

//NewDeskOrderHandlingInst returns a new DeskOrderHandlingInstField initialized with val
func NewDeskOrderHandlingInst(val string) *DeskOrderHandlingInstField {
	field := &DeskOrderHandlingInstField{}
	field.Value = val
	return field
}

//DeskTypeField is a STRING field
type DeskTypeField struct{ quickfix.StringValue }

//Tag returns tag.DeskType (1033)
func (f DeskTypeField) Tag() quickfix.Tag { return tag.DeskType }

//NewDeskType returns a new DeskTypeField initialized with val
func NewDeskType(val string) *DeskTypeField {
	field := &DeskTypeField{}
	field.Value = val
	return field
}

//DeskTypeSourceField is a INT field
type DeskTypeSourceField struct{ quickfix.IntValue }

//Tag returns tag.DeskTypeSource (1034)
func (f DeskTypeSourceField) Tag() quickfix.Tag { return tag.DeskTypeSource }

//NewDeskTypeSource returns a new DeskTypeSourceField initialized with val
func NewDeskTypeSource(val int) *DeskTypeSourceField {
	field := &DeskTypeSourceField{}
	field.Value = val
	return field
}

//DetachmentPointField is a PERCENTAGE field
type DetachmentPointField struct{ quickfix.PercentageValue }

//Tag returns tag.DetachmentPoint (1458)
func (f DetachmentPointField) Tag() quickfix.Tag { return tag.DetachmentPoint }

//NewDetachmentPoint returns a new DetachmentPointField initialized with val
func NewDetachmentPoint(val float64) *DetachmentPointField {
	field := &DetachmentPointField{}
	field.Value = val
	return field
}

//DiscretionInstField is a CHAR field
type DiscretionInstField struct{ quickfix.CharValue }

//Tag returns tag.DiscretionInst (388)
func (f DiscretionInstField) Tag() quickfix.Tag { return tag.DiscretionInst }

//NewDiscretionInst returns a new DiscretionInstField initialized with val
func NewDiscretionInst(val string) *DiscretionInstField {
	field := &DiscretionInstField{}
	field.Value = val
	return field
}

//DiscretionLimitTypeField is a INT field
type DiscretionLimitTypeField struct{ quickfix.IntValue }

//Tag returns tag.DiscretionLimitType (843)
func (f DiscretionLimitTypeField) Tag() quickfix.Tag { return tag.DiscretionLimitType }

//NewDiscretionLimitType returns a new DiscretionLimitTypeField initialized with val
func NewDiscretionLimitType(val int) *DiscretionLimitTypeField {
	field := &DiscretionLimitTypeField{}
	field.Value = val
	return field
}

//DiscretionMoveTypeField is a INT field
type DiscretionMoveTypeField struct{ quickfix.IntValue }

//Tag returns tag.DiscretionMoveType (841)
func (f DiscretionMoveTypeField) Tag() quickfix.Tag { return tag.DiscretionMoveType }

//NewDiscretionMoveType returns a new DiscretionMoveTypeField initialized with val
func NewDiscretionMoveType(val int) *DiscretionMoveTypeField {
	field := &DiscretionMoveTypeField{}
	field.Value = val
	return field
}

//DiscretionOffsetField is a PRICEOFFSET field
type DiscretionOffsetField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.DiscretionOffset (389)
func (f DiscretionOffsetField) Tag() quickfix.Tag { return tag.DiscretionOffset }

//NewDiscretionOffset returns a new DiscretionOffsetField initialized with val
func NewDiscretionOffset(val float64) *DiscretionOffsetField {
	field := &DiscretionOffsetField{}
	field.Value = val
	return field
}

//DiscretionOffsetTypeField is a INT field
type DiscretionOffsetTypeField struct{ quickfix.IntValue }

//Tag returns tag.DiscretionOffsetType (842)
func (f DiscretionOffsetTypeField) Tag() quickfix.Tag { return tag.DiscretionOffsetType }

//NewDiscretionOffsetType returns a new DiscretionOffsetTypeField initialized with val
func NewDiscretionOffsetType(val int) *DiscretionOffsetTypeField {
	field := &DiscretionOffsetTypeField{}
	field.Value = val
	return field
}

//DiscretionOffsetValueField is a FLOAT field
type DiscretionOffsetValueField struct{ quickfix.FloatValue }

//Tag returns tag.DiscretionOffsetValue (389)
func (f DiscretionOffsetValueField) Tag() quickfix.Tag { return tag.DiscretionOffsetValue }

//NewDiscretionOffsetValue returns a new DiscretionOffsetValueField initialized with val
func NewDiscretionOffsetValue(val float64) *DiscretionOffsetValueField {
	field := &DiscretionOffsetValueField{}
	field.Value = val
	return field
}

//DiscretionPriceField is a PRICE field
type DiscretionPriceField struct{ quickfix.PriceValue }

//Tag returns tag.DiscretionPrice (845)
func (f DiscretionPriceField) Tag() quickfix.Tag { return tag.DiscretionPrice }

//NewDiscretionPrice returns a new DiscretionPriceField initialized with val
func NewDiscretionPrice(val float64) *DiscretionPriceField {
	field := &DiscretionPriceField{}
	field.Value = val
	return field
}

//DiscretionRoundDirectionField is a INT field
type DiscretionRoundDirectionField struct{ quickfix.IntValue }

//Tag returns tag.DiscretionRoundDirection (844)
func (f DiscretionRoundDirectionField) Tag() quickfix.Tag { return tag.DiscretionRoundDirection }

//NewDiscretionRoundDirection returns a new DiscretionRoundDirectionField initialized with val
func NewDiscretionRoundDirection(val int) *DiscretionRoundDirectionField {
	field := &DiscretionRoundDirectionField{}
	field.Value = val
	return field
}

//DiscretionScopeField is a INT field
type DiscretionScopeField struct{ quickfix.IntValue }

//Tag returns tag.DiscretionScope (846)
func (f DiscretionScopeField) Tag() quickfix.Tag { return tag.DiscretionScope }

//NewDiscretionScope returns a new DiscretionScopeField initialized with val
func NewDiscretionScope(val int) *DiscretionScopeField {
	field := &DiscretionScopeField{}
	field.Value = val
	return field
}

//DisplayHighQtyField is a QTY field
type DisplayHighQtyField struct{ quickfix.QtyValue }

//Tag returns tag.DisplayHighQty (1086)
func (f DisplayHighQtyField) Tag() quickfix.Tag { return tag.DisplayHighQty }

//NewDisplayHighQty returns a new DisplayHighQtyField initialized with val
func NewDisplayHighQty(val float64) *DisplayHighQtyField {
	field := &DisplayHighQtyField{}
	field.Value = val
	return field
}

//DisplayLowQtyField is a QTY field
type DisplayLowQtyField struct{ quickfix.QtyValue }

//Tag returns tag.DisplayLowQty (1085)
func (f DisplayLowQtyField) Tag() quickfix.Tag { return tag.DisplayLowQty }

//NewDisplayLowQty returns a new DisplayLowQtyField initialized with val
func NewDisplayLowQty(val float64) *DisplayLowQtyField {
	field := &DisplayLowQtyField{}
	field.Value = val
	return field
}

//DisplayMethodField is a CHAR field
type DisplayMethodField struct{ quickfix.CharValue }

//Tag returns tag.DisplayMethod (1084)
func (f DisplayMethodField) Tag() quickfix.Tag { return tag.DisplayMethod }

//NewDisplayMethod returns a new DisplayMethodField initialized with val
func NewDisplayMethod(val string) *DisplayMethodField {
	field := &DisplayMethodField{}
	field.Value = val
	return field
}

//DisplayMinIncrField is a QTY field
type DisplayMinIncrField struct{ quickfix.QtyValue }

//Tag returns tag.DisplayMinIncr (1087)
func (f DisplayMinIncrField) Tag() quickfix.Tag { return tag.DisplayMinIncr }

//NewDisplayMinIncr returns a new DisplayMinIncrField initialized with val
func NewDisplayMinIncr(val float64) *DisplayMinIncrField {
	field := &DisplayMinIncrField{}
	field.Value = val
	return field
}

//DisplayQtyField is a QTY field
type DisplayQtyField struct{ quickfix.QtyValue }

//Tag returns tag.DisplayQty (1138)
func (f DisplayQtyField) Tag() quickfix.Tag { return tag.DisplayQty }

//NewDisplayQty returns a new DisplayQtyField initialized with val
func NewDisplayQty(val float64) *DisplayQtyField {
	field := &DisplayQtyField{}
	field.Value = val
	return field
}

//DisplayWhenField is a CHAR field
type DisplayWhenField struct{ quickfix.CharValue }

//Tag returns tag.DisplayWhen (1083)
func (f DisplayWhenField) Tag() quickfix.Tag { return tag.DisplayWhen }

//NewDisplayWhen returns a new DisplayWhenField initialized with val
func NewDisplayWhen(val string) *DisplayWhenField {
	field := &DisplayWhenField{}
	field.Value = val
	return field
}

//DistribPaymentMethodField is a INT field
type DistribPaymentMethodField struct{ quickfix.IntValue }

//Tag returns tag.DistribPaymentMethod (477)
func (f DistribPaymentMethodField) Tag() quickfix.Tag { return tag.DistribPaymentMethod }

//NewDistribPaymentMethod returns a new DistribPaymentMethodField initialized with val
func NewDistribPaymentMethod(val int) *DistribPaymentMethodField {
	field := &DistribPaymentMethodField{}
	field.Value = val
	return field
}

//DistribPercentageField is a PERCENTAGE field
type DistribPercentageField struct{ quickfix.PercentageValue }

//Tag returns tag.DistribPercentage (512)
func (f DistribPercentageField) Tag() quickfix.Tag { return tag.DistribPercentage }

//NewDistribPercentage returns a new DistribPercentageField initialized with val
func NewDistribPercentage(val float64) *DistribPercentageField {
	field := &DistribPercentageField{}
	field.Value = val
	return field
}

//DividendYieldField is a PERCENTAGE field
type DividendYieldField struct{ quickfix.PercentageValue }

//Tag returns tag.DividendYield (1380)
func (f DividendYieldField) Tag() quickfix.Tag { return tag.DividendYield }

//NewDividendYield returns a new DividendYieldField initialized with val
func NewDividendYield(val float64) *DividendYieldField {
	field := &DividendYieldField{}
	field.Value = val
	return field
}

//DlvyInstField is a STRING field
type DlvyInstField struct{ quickfix.StringValue }

//Tag returns tag.DlvyInst (86)
func (f DlvyInstField) Tag() quickfix.Tag { return tag.DlvyInst }

//NewDlvyInst returns a new DlvyInstField initialized with val
func NewDlvyInst(val string) *DlvyInstField {
	field := &DlvyInstField{}
	field.Value = val
	return field
}

//DlvyInstTypeField is a CHAR field
type DlvyInstTypeField struct{ quickfix.CharValue }

//Tag returns tag.DlvyInstType (787)
func (f DlvyInstTypeField) Tag() quickfix.Tag { return tag.DlvyInstType }

//NewDlvyInstType returns a new DlvyInstTypeField initialized with val
func NewDlvyInstType(val string) *DlvyInstTypeField {
	field := &DlvyInstTypeField{}
	field.Value = val
	return field
}

//DueToRelatedField is a BOOLEAN field
type DueToRelatedField struct{ quickfix.BooleanValue }

//Tag returns tag.DueToRelated (329)
func (f DueToRelatedField) Tag() quickfix.Tag { return tag.DueToRelated }

//NewDueToRelated returns a new DueToRelatedField initialized with val
func NewDueToRelated(val bool) *DueToRelatedField {
	field := &DueToRelatedField{}
	field.Value = val
	return field
}

//EFPTrackingErrorField is a PERCENTAGE field
type EFPTrackingErrorField struct{ quickfix.PercentageValue }

//Tag returns tag.EFPTrackingError (405)
func (f EFPTrackingErrorField) Tag() quickfix.Tag { return tag.EFPTrackingError }

//NewEFPTrackingError returns a new EFPTrackingErrorField initialized with val
func NewEFPTrackingError(val float64) *EFPTrackingErrorField {
	field := &EFPTrackingErrorField{}
	field.Value = val
	return field
}

//EffectiveTimeField is a UTCTIMESTAMP field
type EffectiveTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.EffectiveTime (168)
func (f EffectiveTimeField) Tag() quickfix.Tag { return tag.EffectiveTime }

//EmailThreadIDField is a STRING field
type EmailThreadIDField struct{ quickfix.StringValue }

//Tag returns tag.EmailThreadID (164)
func (f EmailThreadIDField) Tag() quickfix.Tag { return tag.EmailThreadID }

//NewEmailThreadID returns a new EmailThreadIDField initialized with val
func NewEmailThreadID(val string) *EmailThreadIDField {
	field := &EmailThreadIDField{}
	field.Value = val
	return field
}

//EmailTypeField is a CHAR field
type EmailTypeField struct{ quickfix.CharValue }

//Tag returns tag.EmailType (94)
func (f EmailTypeField) Tag() quickfix.Tag { return tag.EmailType }

//NewEmailType returns a new EmailTypeField initialized with val
func NewEmailType(val string) *EmailTypeField {
	field := &EmailTypeField{}
	field.Value = val
	return field
}

//EncodedAllocTextField is a DATA field
type EncodedAllocTextField struct{ quickfix.DataValue }

//Tag returns tag.EncodedAllocText (361)
func (f EncodedAllocTextField) Tag() quickfix.Tag { return tag.EncodedAllocText }

//NewEncodedAllocText returns a new EncodedAllocTextField initialized with val
func NewEncodedAllocText(val string) *EncodedAllocTextField {
	field := &EncodedAllocTextField{}
	field.Value = val
	return field
}

//EncodedAllocTextLenField is a LENGTH field
type EncodedAllocTextLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncodedAllocTextLen (360)
func (f EncodedAllocTextLenField) Tag() quickfix.Tag { return tag.EncodedAllocTextLen }

//NewEncodedAllocTextLen returns a new EncodedAllocTextLenField initialized with val
func NewEncodedAllocTextLen(val int) *EncodedAllocTextLenField {
	field := &EncodedAllocTextLenField{}
	field.Value = val
	return field
}

//EncodedHeadlineField is a DATA field
type EncodedHeadlineField struct{ quickfix.DataValue }

//Tag returns tag.EncodedHeadline (359)
func (f EncodedHeadlineField) Tag() quickfix.Tag { return tag.EncodedHeadline }

//NewEncodedHeadline returns a new EncodedHeadlineField initialized with val
func NewEncodedHeadline(val string) *EncodedHeadlineField {
	field := &EncodedHeadlineField{}
	field.Value = val
	return field
}

//EncodedHeadlineLenField is a LENGTH field
type EncodedHeadlineLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncodedHeadlineLen (358)
func (f EncodedHeadlineLenField) Tag() quickfix.Tag { return tag.EncodedHeadlineLen }

//NewEncodedHeadlineLen returns a new EncodedHeadlineLenField initialized with val
func NewEncodedHeadlineLen(val int) *EncodedHeadlineLenField {
	field := &EncodedHeadlineLenField{}
	field.Value = val
	return field
}

//EncodedIssuerField is a DATA field
type EncodedIssuerField struct{ quickfix.DataValue }

//Tag returns tag.EncodedIssuer (349)
func (f EncodedIssuerField) Tag() quickfix.Tag { return tag.EncodedIssuer }

//NewEncodedIssuer returns a new EncodedIssuerField initialized with val
func NewEncodedIssuer(val string) *EncodedIssuerField {
	field := &EncodedIssuerField{}
	field.Value = val
	return field
}

//EncodedIssuerLenField is a LENGTH field
type EncodedIssuerLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncodedIssuerLen (348)
func (f EncodedIssuerLenField) Tag() quickfix.Tag { return tag.EncodedIssuerLen }

//NewEncodedIssuerLen returns a new EncodedIssuerLenField initialized with val
func NewEncodedIssuerLen(val int) *EncodedIssuerLenField {
	field := &EncodedIssuerLenField{}
	field.Value = val
	return field
}

//EncodedLegIssuerField is a DATA field
type EncodedLegIssuerField struct{ quickfix.DataValue }

//Tag returns tag.EncodedLegIssuer (619)
func (f EncodedLegIssuerField) Tag() quickfix.Tag { return tag.EncodedLegIssuer }

//NewEncodedLegIssuer returns a new EncodedLegIssuerField initialized with val
func NewEncodedLegIssuer(val string) *EncodedLegIssuerField {
	field := &EncodedLegIssuerField{}
	field.Value = val
	return field
}

//EncodedLegIssuerLenField is a LENGTH field
type EncodedLegIssuerLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncodedLegIssuerLen (618)
func (f EncodedLegIssuerLenField) Tag() quickfix.Tag { return tag.EncodedLegIssuerLen }

//NewEncodedLegIssuerLen returns a new EncodedLegIssuerLenField initialized with val
func NewEncodedLegIssuerLen(val int) *EncodedLegIssuerLenField {
	field := &EncodedLegIssuerLenField{}
	field.Value = val
	return field
}

//EncodedLegSecurityDescField is a DATA field
type EncodedLegSecurityDescField struct{ quickfix.DataValue }

//Tag returns tag.EncodedLegSecurityDesc (622)
func (f EncodedLegSecurityDescField) Tag() quickfix.Tag { return tag.EncodedLegSecurityDesc }

//NewEncodedLegSecurityDesc returns a new EncodedLegSecurityDescField initialized with val
func NewEncodedLegSecurityDesc(val string) *EncodedLegSecurityDescField {
	field := &EncodedLegSecurityDescField{}
	field.Value = val
	return field
}

//EncodedLegSecurityDescLenField is a LENGTH field
type EncodedLegSecurityDescLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncodedLegSecurityDescLen (621)
func (f EncodedLegSecurityDescLenField) Tag() quickfix.Tag { return tag.EncodedLegSecurityDescLen }

//NewEncodedLegSecurityDescLen returns a new EncodedLegSecurityDescLenField initialized with val
func NewEncodedLegSecurityDescLen(val int) *EncodedLegSecurityDescLenField {
	field := &EncodedLegSecurityDescLenField{}
	field.Value = val
	return field
}

//EncodedListExecInstField is a DATA field
type EncodedListExecInstField struct{ quickfix.DataValue }

//Tag returns tag.EncodedListExecInst (353)
func (f EncodedListExecInstField) Tag() quickfix.Tag { return tag.EncodedListExecInst }

//NewEncodedListExecInst returns a new EncodedListExecInstField initialized with val
func NewEncodedListExecInst(val string) *EncodedListExecInstField {
	field := &EncodedListExecInstField{}
	field.Value = val
	return field
}

//EncodedListExecInstLenField is a LENGTH field
type EncodedListExecInstLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncodedListExecInstLen (352)
func (f EncodedListExecInstLenField) Tag() quickfix.Tag { return tag.EncodedListExecInstLen }

//NewEncodedListExecInstLen returns a new EncodedListExecInstLenField initialized with val
func NewEncodedListExecInstLen(val int) *EncodedListExecInstLenField {
	field := &EncodedListExecInstLenField{}
	field.Value = val
	return field
}

//EncodedListStatusTextField is a DATA field
type EncodedListStatusTextField struct{ quickfix.DataValue }

//Tag returns tag.EncodedListStatusText (446)
func (f EncodedListStatusTextField) Tag() quickfix.Tag { return tag.EncodedListStatusText }

//NewEncodedListStatusText returns a new EncodedListStatusTextField initialized with val
func NewEncodedListStatusText(val string) *EncodedListStatusTextField {
	field := &EncodedListStatusTextField{}
	field.Value = val
	return field
}

//EncodedListStatusTextLenField is a LENGTH field
type EncodedListStatusTextLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncodedListStatusTextLen (445)
func (f EncodedListStatusTextLenField) Tag() quickfix.Tag { return tag.EncodedListStatusTextLen }

//NewEncodedListStatusTextLen returns a new EncodedListStatusTextLenField initialized with val
func NewEncodedListStatusTextLen(val int) *EncodedListStatusTextLenField {
	field := &EncodedListStatusTextLenField{}
	field.Value = val
	return field
}

//EncodedMktSegmDescField is a DATA field
type EncodedMktSegmDescField struct{ quickfix.DataValue }

//Tag returns tag.EncodedMktSegmDesc (1398)
func (f EncodedMktSegmDescField) Tag() quickfix.Tag { return tag.EncodedMktSegmDesc }

//NewEncodedMktSegmDesc returns a new EncodedMktSegmDescField initialized with val
func NewEncodedMktSegmDesc(val string) *EncodedMktSegmDescField {
	field := &EncodedMktSegmDescField{}
	field.Value = val
	return field
}

//EncodedMktSegmDescLenField is a LENGTH field
type EncodedMktSegmDescLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncodedMktSegmDescLen (1397)
func (f EncodedMktSegmDescLenField) Tag() quickfix.Tag { return tag.EncodedMktSegmDescLen }

//NewEncodedMktSegmDescLen returns a new EncodedMktSegmDescLenField initialized with val
func NewEncodedMktSegmDescLen(val int) *EncodedMktSegmDescLenField {
	field := &EncodedMktSegmDescLenField{}
	field.Value = val
	return field
}

//EncodedSecurityDescField is a DATA field
type EncodedSecurityDescField struct{ quickfix.DataValue }

//Tag returns tag.EncodedSecurityDesc (351)
func (f EncodedSecurityDescField) Tag() quickfix.Tag { return tag.EncodedSecurityDesc }

//NewEncodedSecurityDesc returns a new EncodedSecurityDescField initialized with val
func NewEncodedSecurityDesc(val string) *EncodedSecurityDescField {
	field := &EncodedSecurityDescField{}
	field.Value = val
	return field
}

//EncodedSecurityDescLenField is a LENGTH field
type EncodedSecurityDescLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncodedSecurityDescLen (350)
func (f EncodedSecurityDescLenField) Tag() quickfix.Tag { return tag.EncodedSecurityDescLen }

//NewEncodedSecurityDescLen returns a new EncodedSecurityDescLenField initialized with val
func NewEncodedSecurityDescLen(val int) *EncodedSecurityDescLenField {
	field := &EncodedSecurityDescLenField{}
	field.Value = val
	return field
}

//EncodedSecurityListDescField is a DATA field
type EncodedSecurityListDescField struct{ quickfix.DataValue }

//Tag returns tag.EncodedSecurityListDesc (1469)
func (f EncodedSecurityListDescField) Tag() quickfix.Tag { return tag.EncodedSecurityListDesc }

//NewEncodedSecurityListDesc returns a new EncodedSecurityListDescField initialized with val
func NewEncodedSecurityListDesc(val string) *EncodedSecurityListDescField {
	field := &EncodedSecurityListDescField{}
	field.Value = val
	return field
}

//EncodedSecurityListDescLenField is a LENGTH field
type EncodedSecurityListDescLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncodedSecurityListDescLen (1468)
func (f EncodedSecurityListDescLenField) Tag() quickfix.Tag { return tag.EncodedSecurityListDescLen }

//NewEncodedSecurityListDescLen returns a new EncodedSecurityListDescLenField initialized with val
func NewEncodedSecurityListDescLen(val int) *EncodedSecurityListDescLenField {
	field := &EncodedSecurityListDescLenField{}
	field.Value = val
	return field
}

//EncodedSubjectField is a DATA field
type EncodedSubjectField struct{ quickfix.DataValue }

//Tag returns tag.EncodedSubject (357)
func (f EncodedSubjectField) Tag() quickfix.Tag { return tag.EncodedSubject }

//NewEncodedSubject returns a new EncodedSubjectField initialized with val
func NewEncodedSubject(val string) *EncodedSubjectField {
	field := &EncodedSubjectField{}
	field.Value = val
	return field
}

//EncodedSubjectLenField is a LENGTH field
type EncodedSubjectLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncodedSubjectLen (356)
func (f EncodedSubjectLenField) Tag() quickfix.Tag { return tag.EncodedSubjectLen }

//NewEncodedSubjectLen returns a new EncodedSubjectLenField initialized with val
func NewEncodedSubjectLen(val int) *EncodedSubjectLenField {
	field := &EncodedSubjectLenField{}
	field.Value = val
	return field
}

//EncodedSymbolField is a DATA field
type EncodedSymbolField struct{ quickfix.DataValue }

//Tag returns tag.EncodedSymbol (1360)
func (f EncodedSymbolField) Tag() quickfix.Tag { return tag.EncodedSymbol }

//NewEncodedSymbol returns a new EncodedSymbolField initialized with val
func NewEncodedSymbol(val string) *EncodedSymbolField {
	field := &EncodedSymbolField{}
	field.Value = val
	return field
}

//EncodedSymbolLenField is a LENGTH field
type EncodedSymbolLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncodedSymbolLen (1359)
func (f EncodedSymbolLenField) Tag() quickfix.Tag { return tag.EncodedSymbolLen }

//NewEncodedSymbolLen returns a new EncodedSymbolLenField initialized with val
func NewEncodedSymbolLen(val int) *EncodedSymbolLenField {
	field := &EncodedSymbolLenField{}
	field.Value = val
	return field
}

//EncodedTextField is a DATA field
type EncodedTextField struct{ quickfix.DataValue }

//Tag returns tag.EncodedText (355)
func (f EncodedTextField) Tag() quickfix.Tag { return tag.EncodedText }

//NewEncodedText returns a new EncodedTextField initialized with val
func NewEncodedText(val string) *EncodedTextField {
	field := &EncodedTextField{}
	field.Value = val
	return field
}

//EncodedTextLenField is a LENGTH field
type EncodedTextLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncodedTextLen (354)
func (f EncodedTextLenField) Tag() quickfix.Tag { return tag.EncodedTextLen }

//NewEncodedTextLen returns a new EncodedTextLenField initialized with val
func NewEncodedTextLen(val int) *EncodedTextLenField {
	field := &EncodedTextLenField{}
	field.Value = val
	return field
}

//EncodedUnderlyingIssuerField is a DATA field
type EncodedUnderlyingIssuerField struct{ quickfix.DataValue }

//Tag returns tag.EncodedUnderlyingIssuer (363)
func (f EncodedUnderlyingIssuerField) Tag() quickfix.Tag { return tag.EncodedUnderlyingIssuer }

//NewEncodedUnderlyingIssuer returns a new EncodedUnderlyingIssuerField initialized with val
func NewEncodedUnderlyingIssuer(val string) *EncodedUnderlyingIssuerField {
	field := &EncodedUnderlyingIssuerField{}
	field.Value = val
	return field
}

//EncodedUnderlyingIssuerLenField is a LENGTH field
type EncodedUnderlyingIssuerLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncodedUnderlyingIssuerLen (362)
func (f EncodedUnderlyingIssuerLenField) Tag() quickfix.Tag { return tag.EncodedUnderlyingIssuerLen }

//NewEncodedUnderlyingIssuerLen returns a new EncodedUnderlyingIssuerLenField initialized with val
func NewEncodedUnderlyingIssuerLen(val int) *EncodedUnderlyingIssuerLenField {
	field := &EncodedUnderlyingIssuerLenField{}
	field.Value = val
	return field
}

//EncodedUnderlyingSecurityDescField is a DATA field
type EncodedUnderlyingSecurityDescField struct{ quickfix.DataValue }

//Tag returns tag.EncodedUnderlyingSecurityDesc (365)
func (f EncodedUnderlyingSecurityDescField) Tag() quickfix.Tag {
	return tag.EncodedUnderlyingSecurityDesc
}

//NewEncodedUnderlyingSecurityDesc returns a new EncodedUnderlyingSecurityDescField initialized with val
func NewEncodedUnderlyingSecurityDesc(val string) *EncodedUnderlyingSecurityDescField {
	field := &EncodedUnderlyingSecurityDescField{}
	field.Value = val
	return field
}

//EncodedUnderlyingSecurityDescLenField is a LENGTH field
type EncodedUnderlyingSecurityDescLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncodedUnderlyingSecurityDescLen (364)
func (f EncodedUnderlyingSecurityDescLenField) Tag() quickfix.Tag {
	return tag.EncodedUnderlyingSecurityDescLen
}

//NewEncodedUnderlyingSecurityDescLen returns a new EncodedUnderlyingSecurityDescLenField initialized with val
func NewEncodedUnderlyingSecurityDescLen(val int) *EncodedUnderlyingSecurityDescLenField {
	field := &EncodedUnderlyingSecurityDescLenField{}
	field.Value = val
	return field
}

//EncryptMethodField is a INT field
type EncryptMethodField struct{ quickfix.IntValue }

//Tag returns tag.EncryptMethod (98)
func (f EncryptMethodField) Tag() quickfix.Tag { return tag.EncryptMethod }

//NewEncryptMethod returns a new EncryptMethodField initialized with val
func NewEncryptMethod(val int) *EncryptMethodField {
	field := &EncryptMethodField{}
	field.Value = val
	return field
}

//EncryptedNewPasswordField is a DATA field
type EncryptedNewPasswordField struct{ quickfix.DataValue }

//Tag returns tag.EncryptedNewPassword (1404)
func (f EncryptedNewPasswordField) Tag() quickfix.Tag { return tag.EncryptedNewPassword }

//NewEncryptedNewPassword returns a new EncryptedNewPasswordField initialized with val
func NewEncryptedNewPassword(val string) *EncryptedNewPasswordField {
	field := &EncryptedNewPasswordField{}
	field.Value = val
	return field
}

//EncryptedNewPasswordLenField is a LENGTH field
type EncryptedNewPasswordLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncryptedNewPasswordLen (1403)
func (f EncryptedNewPasswordLenField) Tag() quickfix.Tag { return tag.EncryptedNewPasswordLen }

//NewEncryptedNewPasswordLen returns a new EncryptedNewPasswordLenField initialized with val
func NewEncryptedNewPasswordLen(val int) *EncryptedNewPasswordLenField {
	field := &EncryptedNewPasswordLenField{}
	field.Value = val
	return field
}

//EncryptedPasswordField is a DATA field
type EncryptedPasswordField struct{ quickfix.DataValue }

//Tag returns tag.EncryptedPassword (1402)
func (f EncryptedPasswordField) Tag() quickfix.Tag { return tag.EncryptedPassword }

//NewEncryptedPassword returns a new EncryptedPasswordField initialized with val
func NewEncryptedPassword(val string) *EncryptedPasswordField {
	field := &EncryptedPasswordField{}
	field.Value = val
	return field
}

//EncryptedPasswordLenField is a LENGTH field
type EncryptedPasswordLenField struct{ quickfix.LengthValue }

//Tag returns tag.EncryptedPasswordLen (1401)
func (f EncryptedPasswordLenField) Tag() quickfix.Tag { return tag.EncryptedPasswordLen }

//NewEncryptedPasswordLen returns a new EncryptedPasswordLenField initialized with val
func NewEncryptedPasswordLen(val int) *EncryptedPasswordLenField {
	field := &EncryptedPasswordLenField{}
	field.Value = val
	return field
}

//EncryptedPasswordMethodField is a INT field
type EncryptedPasswordMethodField struct{ quickfix.IntValue }

//Tag returns tag.EncryptedPasswordMethod (1400)
func (f EncryptedPasswordMethodField) Tag() quickfix.Tag { return tag.EncryptedPasswordMethod }

//NewEncryptedPasswordMethod returns a new EncryptedPasswordMethodField initialized with val
func NewEncryptedPasswordMethod(val int) *EncryptedPasswordMethodField {
	field := &EncryptedPasswordMethodField{}
	field.Value = val
	return field
}

//EndAccruedInterestAmtField is a AMT field
type EndAccruedInterestAmtField struct{ quickfix.AmtValue }

//Tag returns tag.EndAccruedInterestAmt (920)
func (f EndAccruedInterestAmtField) Tag() quickfix.Tag { return tag.EndAccruedInterestAmt }

//NewEndAccruedInterestAmt returns a new EndAccruedInterestAmtField initialized with val
func NewEndAccruedInterestAmt(val float64) *EndAccruedInterestAmtField {
	field := &EndAccruedInterestAmtField{}
	field.Value = val
	return field
}

//EndCashField is a AMT field
type EndCashField struct{ quickfix.AmtValue }

//Tag returns tag.EndCash (922)
func (f EndCashField) Tag() quickfix.Tag { return tag.EndCash }

//NewEndCash returns a new EndCashField initialized with val
func NewEndCash(val float64) *EndCashField {
	field := &EndCashField{}
	field.Value = val
	return field
}

//EndDateField is a LOCALMKTDATE field
type EndDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.EndDate (917)
func (f EndDateField) Tag() quickfix.Tag { return tag.EndDate }

//NewEndDate returns a new EndDateField initialized with val
func NewEndDate(val string) *EndDateField {
	field := &EndDateField{}
	field.Value = val
	return field
}

//EndMaturityMonthYearField is a MONTHYEAR field
type EndMaturityMonthYearField struct{ quickfix.MonthYearValue }

//Tag returns tag.EndMaturityMonthYear (1226)
func (f EndMaturityMonthYearField) Tag() quickfix.Tag { return tag.EndMaturityMonthYear }

//NewEndMaturityMonthYear returns a new EndMaturityMonthYearField initialized with val
func NewEndMaturityMonthYear(val string) *EndMaturityMonthYearField {
	field := &EndMaturityMonthYearField{}
	field.Value = val
	return field
}

//EndSeqNoField is a SEQNUM field
type EndSeqNoField struct{ quickfix.SeqNumValue }

//Tag returns tag.EndSeqNo (16)
func (f EndSeqNoField) Tag() quickfix.Tag { return tag.EndSeqNo }

//NewEndSeqNo returns a new EndSeqNoField initialized with val
func NewEndSeqNo(val int) *EndSeqNoField {
	field := &EndSeqNoField{}
	field.Value = val
	return field
}

//EndStrikePxRangeField is a PRICE field
type EndStrikePxRangeField struct{ quickfix.PriceValue }

//Tag returns tag.EndStrikePxRange (1203)
func (f EndStrikePxRangeField) Tag() quickfix.Tag { return tag.EndStrikePxRange }

//NewEndStrikePxRange returns a new EndStrikePxRangeField initialized with val
func NewEndStrikePxRange(val float64) *EndStrikePxRangeField {
	field := &EndStrikePxRangeField{}
	field.Value = val
	return field
}

//EndTickPriceRangeField is a PRICE field
type EndTickPriceRangeField struct{ quickfix.PriceValue }

//Tag returns tag.EndTickPriceRange (1207)
func (f EndTickPriceRangeField) Tag() quickfix.Tag { return tag.EndTickPriceRange }

//NewEndTickPriceRange returns a new EndTickPriceRangeField initialized with val
func NewEndTickPriceRange(val float64) *EndTickPriceRangeField {
	field := &EndTickPriceRangeField{}
	field.Value = val
	return field
}

//EventDateField is a LOCALMKTDATE field
type EventDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.EventDate (866)
func (f EventDateField) Tag() quickfix.Tag { return tag.EventDate }

//NewEventDate returns a new EventDateField initialized with val
func NewEventDate(val string) *EventDateField {
	field := &EventDateField{}
	field.Value = val
	return field
}

//EventPxField is a PRICE field
type EventPxField struct{ quickfix.PriceValue }

//Tag returns tag.EventPx (867)
func (f EventPxField) Tag() quickfix.Tag { return tag.EventPx }

//NewEventPx returns a new EventPxField initialized with val
func NewEventPx(val float64) *EventPxField {
	field := &EventPxField{}
	field.Value = val
	return field
}

//EventTextField is a STRING field
type EventTextField struct{ quickfix.StringValue }

//Tag returns tag.EventText (868)
func (f EventTextField) Tag() quickfix.Tag { return tag.EventText }

//NewEventText returns a new EventTextField initialized with val
func NewEventText(val string) *EventTextField {
	field := &EventTextField{}
	field.Value = val
	return field
}

//EventTimeField is a UTCTIMESTAMP field
type EventTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.EventTime (1145)
func (f EventTimeField) Tag() quickfix.Tag { return tag.EventTime }

//EventTypeField is a INT field
type EventTypeField struct{ quickfix.IntValue }

//Tag returns tag.EventType (865)
func (f EventTypeField) Tag() quickfix.Tag { return tag.EventType }

//NewEventType returns a new EventTypeField initialized with val
func NewEventType(val int) *EventTypeField {
	field := &EventTypeField{}
	field.Value = val
	return field
}

//ExDateField is a LOCALMKTDATE field
type ExDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.ExDate (230)
func (f ExDateField) Tag() quickfix.Tag { return tag.ExDate }

//NewExDate returns a new ExDateField initialized with val
func NewExDate(val string) *ExDateField {
	field := &ExDateField{}
	field.Value = val
	return field
}

//ExDestinationField is a EXCHANGE field
type ExDestinationField struct{ quickfix.ExchangeValue }

//Tag returns tag.ExDestination (100)
func (f ExDestinationField) Tag() quickfix.Tag { return tag.ExDestination }

//NewExDestination returns a new ExDestinationField initialized with val
func NewExDestination(val string) *ExDestinationField {
	field := &ExDestinationField{}
	field.Value = val
	return field
}

//ExDestinationIDSourceField is a CHAR field
type ExDestinationIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.ExDestinationIDSource (1133)
func (f ExDestinationIDSourceField) Tag() quickfix.Tag { return tag.ExDestinationIDSource }

//NewExDestinationIDSource returns a new ExDestinationIDSourceField initialized with val
func NewExDestinationIDSource(val string) *ExDestinationIDSourceField {
	field := &ExDestinationIDSourceField{}
	field.Value = val
	return field
}

//ExchangeForPhysicalField is a BOOLEAN field
type ExchangeForPhysicalField struct{ quickfix.BooleanValue }

//Tag returns tag.ExchangeForPhysical (411)
func (f ExchangeForPhysicalField) Tag() quickfix.Tag { return tag.ExchangeForPhysical }

//NewExchangeForPhysical returns a new ExchangeForPhysicalField initialized with val
func NewExchangeForPhysical(val bool) *ExchangeForPhysicalField {
	field := &ExchangeForPhysicalField{}
	field.Value = val
	return field
}

//ExchangeRuleField is a STRING field
type ExchangeRuleField struct{ quickfix.StringValue }

//Tag returns tag.ExchangeRule (825)
func (f ExchangeRuleField) Tag() quickfix.Tag { return tag.ExchangeRule }

//NewExchangeRule returns a new ExchangeRuleField initialized with val
func NewExchangeRule(val string) *ExchangeRuleField {
	field := &ExchangeRuleField{}
	field.Value = val
	return field
}

//ExchangeSpecialInstructionsField is a STRING field
type ExchangeSpecialInstructionsField struct{ quickfix.StringValue }

//Tag returns tag.ExchangeSpecialInstructions (1139)
func (f ExchangeSpecialInstructionsField) Tag() quickfix.Tag { return tag.ExchangeSpecialInstructions }

//NewExchangeSpecialInstructions returns a new ExchangeSpecialInstructionsField initialized with val
func NewExchangeSpecialInstructions(val string) *ExchangeSpecialInstructionsField {
	field := &ExchangeSpecialInstructionsField{}
	field.Value = val
	return field
}

//ExecAckStatusField is a CHAR field
type ExecAckStatusField struct{ quickfix.CharValue }

//Tag returns tag.ExecAckStatus (1036)
func (f ExecAckStatusField) Tag() quickfix.Tag { return tag.ExecAckStatus }

//NewExecAckStatus returns a new ExecAckStatusField initialized with val
func NewExecAckStatus(val string) *ExecAckStatusField {
	field := &ExecAckStatusField{}
	field.Value = val
	return field
}

//ExecBrokerField is a STRING field
type ExecBrokerField struct{ quickfix.StringValue }

//Tag returns tag.ExecBroker (76)
func (f ExecBrokerField) Tag() quickfix.Tag { return tag.ExecBroker }

//NewExecBroker returns a new ExecBrokerField initialized with val
func NewExecBroker(val string) *ExecBrokerField {
	field := &ExecBrokerField{}
	field.Value = val
	return field
}

//ExecIDField is a STRING field
type ExecIDField struct{ quickfix.StringValue }

//Tag returns tag.ExecID (17)
func (f ExecIDField) Tag() quickfix.Tag { return tag.ExecID }

//NewExecID returns a new ExecIDField initialized with val
func NewExecID(val string) *ExecIDField {
	field := &ExecIDField{}
	field.Value = val
	return field
}

//ExecInstField is a MULTIPLECHARVALUE field
type ExecInstField struct{ quickfix.MultipleCharValue }

//Tag returns tag.ExecInst (18)
func (f ExecInstField) Tag() quickfix.Tag { return tag.ExecInst }

//NewExecInst returns a new ExecInstField initialized with val
func NewExecInst(val string) *ExecInstField {
	field := &ExecInstField{}
	field.Value = val
	return field
}

//ExecInstValueField is a CHAR field
type ExecInstValueField struct{ quickfix.CharValue }

//Tag returns tag.ExecInstValue (1308)
func (f ExecInstValueField) Tag() quickfix.Tag { return tag.ExecInstValue }

//NewExecInstValue returns a new ExecInstValueField initialized with val
func NewExecInstValue(val string) *ExecInstValueField {
	field := &ExecInstValueField{}
	field.Value = val
	return field
}

//ExecPriceAdjustmentField is a FLOAT field
type ExecPriceAdjustmentField struct{ quickfix.FloatValue }

//Tag returns tag.ExecPriceAdjustment (485)
func (f ExecPriceAdjustmentField) Tag() quickfix.Tag { return tag.ExecPriceAdjustment }

//NewExecPriceAdjustment returns a new ExecPriceAdjustmentField initialized with val
func NewExecPriceAdjustment(val float64) *ExecPriceAdjustmentField {
	field := &ExecPriceAdjustmentField{}
	field.Value = val
	return field
}

//ExecPriceTypeField is a CHAR field
type ExecPriceTypeField struct{ quickfix.CharValue }

//Tag returns tag.ExecPriceType (484)
func (f ExecPriceTypeField) Tag() quickfix.Tag { return tag.ExecPriceType }

//NewExecPriceType returns a new ExecPriceTypeField initialized with val
func NewExecPriceType(val string) *ExecPriceTypeField {
	field := &ExecPriceTypeField{}
	field.Value = val
	return field
}

//ExecRefIDField is a STRING field
type ExecRefIDField struct{ quickfix.StringValue }

//Tag returns tag.ExecRefID (19)
func (f ExecRefIDField) Tag() quickfix.Tag { return tag.ExecRefID }

//NewExecRefID returns a new ExecRefIDField initialized with val
func NewExecRefID(val string) *ExecRefIDField {
	field := &ExecRefIDField{}
	field.Value = val
	return field
}

//ExecRestatementReasonField is a INT field
type ExecRestatementReasonField struct{ quickfix.IntValue }

//Tag returns tag.ExecRestatementReason (378)
func (f ExecRestatementReasonField) Tag() quickfix.Tag { return tag.ExecRestatementReason }

//NewExecRestatementReason returns a new ExecRestatementReasonField initialized with val
func NewExecRestatementReason(val int) *ExecRestatementReasonField {
	field := &ExecRestatementReasonField{}
	field.Value = val
	return field
}

//ExecTransTypeField is a CHAR field
type ExecTransTypeField struct{ quickfix.CharValue }

//Tag returns tag.ExecTransType (20)
func (f ExecTransTypeField) Tag() quickfix.Tag { return tag.ExecTransType }

//NewExecTransType returns a new ExecTransTypeField initialized with val
func NewExecTransType(val string) *ExecTransTypeField {
	field := &ExecTransTypeField{}
	field.Value = val
	return field
}

//ExecTypeField is a CHAR field
type ExecTypeField struct{ quickfix.CharValue }

//Tag returns tag.ExecType (150)
func (f ExecTypeField) Tag() quickfix.Tag { return tag.ExecType }

//NewExecType returns a new ExecTypeField initialized with val
func NewExecType(val string) *ExecTypeField {
	field := &ExecTypeField{}
	field.Value = val
	return field
}

//ExecValuationPointField is a UTCTIMESTAMP field
type ExecValuationPointField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.ExecValuationPoint (515)
func (f ExecValuationPointField) Tag() quickfix.Tag { return tag.ExecValuationPoint }

//ExerciseMethodField is a CHAR field
type ExerciseMethodField struct{ quickfix.CharValue }

//Tag returns tag.ExerciseMethod (747)
func (f ExerciseMethodField) Tag() quickfix.Tag { return tag.ExerciseMethod }

//NewExerciseMethod returns a new ExerciseMethodField initialized with val
func NewExerciseMethod(val string) *ExerciseMethodField {
	field := &ExerciseMethodField{}
	field.Value = val
	return field
}

//ExerciseStyleField is a INT field
type ExerciseStyleField struct{ quickfix.IntValue }

//Tag returns tag.ExerciseStyle (1194)
func (f ExerciseStyleField) Tag() quickfix.Tag { return tag.ExerciseStyle }

//NewExerciseStyle returns a new ExerciseStyleField initialized with val
func NewExerciseStyle(val int) *ExerciseStyleField {
	field := &ExerciseStyleField{}
	field.Value = val
	return field
}

//ExpQtyField is a QTY field
type ExpQtyField struct{ quickfix.QtyValue }

//Tag returns tag.ExpQty (983)
func (f ExpQtyField) Tag() quickfix.Tag { return tag.ExpQty }

//NewExpQty returns a new ExpQtyField initialized with val
func NewExpQty(val float64) *ExpQtyField {
	field := &ExpQtyField{}
	field.Value = val
	return field
}

//ExpTypeField is a INT field
type ExpTypeField struct{ quickfix.IntValue }

//Tag returns tag.ExpType (982)
func (f ExpTypeField) Tag() quickfix.Tag { return tag.ExpType }

//NewExpType returns a new ExpTypeField initialized with val
func NewExpType(val int) *ExpTypeField {
	field := &ExpTypeField{}
	field.Value = val
	return field
}

//ExpirationCycleField is a INT field
type ExpirationCycleField struct{ quickfix.IntValue }

//Tag returns tag.ExpirationCycle (827)
func (f ExpirationCycleField) Tag() quickfix.Tag { return tag.ExpirationCycle }

//NewExpirationCycle returns a new ExpirationCycleField initialized with val
func NewExpirationCycle(val int) *ExpirationCycleField {
	field := &ExpirationCycleField{}
	field.Value = val
	return field
}

//ExpirationQtyTypeField is a INT field
type ExpirationQtyTypeField struct{ quickfix.IntValue }

//Tag returns tag.ExpirationQtyType (982)
func (f ExpirationQtyTypeField) Tag() quickfix.Tag { return tag.ExpirationQtyType }

//NewExpirationQtyType returns a new ExpirationQtyTypeField initialized with val
func NewExpirationQtyType(val int) *ExpirationQtyTypeField {
	field := &ExpirationQtyTypeField{}
	field.Value = val
	return field
}

//ExpireDateField is a LOCALMKTDATE field
type ExpireDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.ExpireDate (432)
func (f ExpireDateField) Tag() quickfix.Tag { return tag.ExpireDate }

//NewExpireDate returns a new ExpireDateField initialized with val
func NewExpireDate(val string) *ExpireDateField {
	field := &ExpireDateField{}
	field.Value = val
	return field
}

//ExpireTimeField is a UTCTIMESTAMP field
type ExpireTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.ExpireTime (126)
func (f ExpireTimeField) Tag() quickfix.Tag { return tag.ExpireTime }

//FactorField is a FLOAT field
type FactorField struct{ quickfix.FloatValue }

//Tag returns tag.Factor (228)
func (f FactorField) Tag() quickfix.Tag { return tag.Factor }

//NewFactor returns a new FactorField initialized with val
func NewFactor(val float64) *FactorField {
	field := &FactorField{}
	field.Value = val
	return field
}

//FairValueField is a AMT field
type FairValueField struct{ quickfix.AmtValue }

//Tag returns tag.FairValue (406)
func (f FairValueField) Tag() quickfix.Tag { return tag.FairValue }

//NewFairValue returns a new FairValueField initialized with val
func NewFairValue(val float64) *FairValueField {
	field := &FairValueField{}
	field.Value = val
	return field
}

//FeeMultiplierField is a FLOAT field
type FeeMultiplierField struct{ quickfix.FloatValue }

//Tag returns tag.FeeMultiplier (1329)
func (f FeeMultiplierField) Tag() quickfix.Tag { return tag.FeeMultiplier }

//NewFeeMultiplier returns a new FeeMultiplierField initialized with val
func NewFeeMultiplier(val float64) *FeeMultiplierField {
	field := &FeeMultiplierField{}
	field.Value = val
	return field
}

//FillExecIDField is a STRING field
type FillExecIDField struct{ quickfix.StringValue }

//Tag returns tag.FillExecID (1363)
func (f FillExecIDField) Tag() quickfix.Tag { return tag.FillExecID }

//NewFillExecID returns a new FillExecIDField initialized with val
func NewFillExecID(val string) *FillExecIDField {
	field := &FillExecIDField{}
	field.Value = val
	return field
}

//FillLiquidityIndField is a INT field
type FillLiquidityIndField struct{ quickfix.IntValue }

//Tag returns tag.FillLiquidityInd (1443)
func (f FillLiquidityIndField) Tag() quickfix.Tag { return tag.FillLiquidityInd }

//NewFillLiquidityInd returns a new FillLiquidityIndField initialized with val
func NewFillLiquidityInd(val int) *FillLiquidityIndField {
	field := &FillLiquidityIndField{}
	field.Value = val
	return field
}

//FillPxField is a PRICE field
type FillPxField struct{ quickfix.PriceValue }

//Tag returns tag.FillPx (1364)
func (f FillPxField) Tag() quickfix.Tag { return tag.FillPx }

//NewFillPx returns a new FillPxField initialized with val
func NewFillPx(val float64) *FillPxField {
	field := &FillPxField{}
	field.Value = val
	return field
}

//FillQtyField is a QTY field
type FillQtyField struct{ quickfix.QtyValue }

//Tag returns tag.FillQty (1365)
func (f FillQtyField) Tag() quickfix.Tag { return tag.FillQty }

//NewFillQty returns a new FillQtyField initialized with val
func NewFillQty(val float64) *FillQtyField {
	field := &FillQtyField{}
	field.Value = val
	return field
}

//FinancialStatusField is a MULTIPLECHARVALUE field
type FinancialStatusField struct{ quickfix.MultipleCharValue }

//Tag returns tag.FinancialStatus (291)
func (f FinancialStatusField) Tag() quickfix.Tag { return tag.FinancialStatus }

//NewFinancialStatus returns a new FinancialStatusField initialized with val
func NewFinancialStatus(val string) *FinancialStatusField {
	field := &FinancialStatusField{}
	field.Value = val
	return field
}

//FirmTradeIDField is a STRING field
type FirmTradeIDField struct{ quickfix.StringValue }

//Tag returns tag.FirmTradeID (1041)
func (f FirmTradeIDField) Tag() quickfix.Tag { return tag.FirmTradeID }

//NewFirmTradeID returns a new FirmTradeIDField initialized with val
func NewFirmTradeID(val string) *FirmTradeIDField {
	field := &FirmTradeIDField{}
	field.Value = val
	return field
}

//FirstPxField is a PRICE field
type FirstPxField struct{ quickfix.PriceValue }

//Tag returns tag.FirstPx (1025)
func (f FirstPxField) Tag() quickfix.Tag { return tag.FirstPx }

//NewFirstPx returns a new FirstPxField initialized with val
func NewFirstPx(val float64) *FirstPxField {
	field := &FirstPxField{}
	field.Value = val
	return field
}

//FlexProductEligibilityIndicatorField is a BOOLEAN field
type FlexProductEligibilityIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.FlexProductEligibilityIndicator (1242)
func (f FlexProductEligibilityIndicatorField) Tag() quickfix.Tag {
	return tag.FlexProductEligibilityIndicator
}

//NewFlexProductEligibilityIndicator returns a new FlexProductEligibilityIndicatorField initialized with val
func NewFlexProductEligibilityIndicator(val bool) *FlexProductEligibilityIndicatorField {
	field := &FlexProductEligibilityIndicatorField{}
	field.Value = val
	return field
}

//FlexibleIndicatorField is a BOOLEAN field
type FlexibleIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.FlexibleIndicator (1244)
func (f FlexibleIndicatorField) Tag() quickfix.Tag { return tag.FlexibleIndicator }

//NewFlexibleIndicator returns a new FlexibleIndicatorField initialized with val
func NewFlexibleIndicator(val bool) *FlexibleIndicatorField {
	field := &FlexibleIndicatorField{}
	field.Value = val
	return field
}

//FloorPriceField is a PRICE field
type FloorPriceField struct{ quickfix.PriceValue }

//Tag returns tag.FloorPrice (1200)
func (f FloorPriceField) Tag() quickfix.Tag { return tag.FloorPrice }

//NewFloorPrice returns a new FloorPriceField initialized with val
func NewFloorPrice(val float64) *FloorPriceField {
	field := &FloorPriceField{}
	field.Value = val
	return field
}

//FlowScheduleTypeField is a INT field
type FlowScheduleTypeField struct{ quickfix.IntValue }

//Tag returns tag.FlowScheduleType (1439)
func (f FlowScheduleTypeField) Tag() quickfix.Tag { return tag.FlowScheduleType }

//NewFlowScheduleType returns a new FlowScheduleTypeField initialized with val
func NewFlowScheduleType(val int) *FlowScheduleTypeField {
	field := &FlowScheduleTypeField{}
	field.Value = val
	return field
}

//ForexReqField is a BOOLEAN field
type ForexReqField struct{ quickfix.BooleanValue }

//Tag returns tag.ForexReq (121)
func (f ForexReqField) Tag() quickfix.Tag { return tag.ForexReq }

//NewForexReq returns a new ForexReqField initialized with val
func NewForexReq(val bool) *ForexReqField {
	field := &ForexReqField{}
	field.Value = val
	return field
}

//FundRenewWaivField is a CHAR field
type FundRenewWaivField struct{ quickfix.CharValue }

//Tag returns tag.FundRenewWaiv (497)
func (f FundRenewWaivField) Tag() quickfix.Tag { return tag.FundRenewWaiv }

//NewFundRenewWaiv returns a new FundRenewWaivField initialized with val
func NewFundRenewWaiv(val string) *FundRenewWaivField {
	field := &FundRenewWaivField{}
	field.Value = val
	return field
}

//FutSettDateField is a LOCALMKTDATE field
type FutSettDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.FutSettDate (64)
func (f FutSettDateField) Tag() quickfix.Tag { return tag.FutSettDate }

//NewFutSettDate returns a new FutSettDateField initialized with val
func NewFutSettDate(val string) *FutSettDateField {
	field := &FutSettDateField{}
	field.Value = val
	return field
}

//FutSettDate2Field is a LOCALMKTDATE field
type FutSettDate2Field struct{ quickfix.LocalMktDateValue }

//Tag returns tag.FutSettDate2 (193)
func (f FutSettDate2Field) Tag() quickfix.Tag { return tag.FutSettDate2 }

//NewFutSettDate2 returns a new FutSettDate2Field initialized with val
func NewFutSettDate2(val string) *FutSettDate2Field {
	field := &FutSettDate2Field{}
	field.Value = val
	return field
}

//FuturesValuationMethodField is a STRING field
type FuturesValuationMethodField struct{ quickfix.StringValue }

//Tag returns tag.FuturesValuationMethod (1197)
func (f FuturesValuationMethodField) Tag() quickfix.Tag { return tag.FuturesValuationMethod }

//NewFuturesValuationMethod returns a new FuturesValuationMethodField initialized with val
func NewFuturesValuationMethod(val string) *FuturesValuationMethodField {
	field := &FuturesValuationMethodField{}
	field.Value = val
	return field
}

//GTBookingInstField is a INT field
type GTBookingInstField struct{ quickfix.IntValue }

//Tag returns tag.GTBookingInst (427)
func (f GTBookingInstField) Tag() quickfix.Tag { return tag.GTBookingInst }

//NewGTBookingInst returns a new GTBookingInstField initialized with val
func NewGTBookingInst(val int) *GTBookingInstField {
	field := &GTBookingInstField{}
	field.Value = val
	return field
}

//GapFillFlagField is a BOOLEAN field
type GapFillFlagField struct{ quickfix.BooleanValue }

//Tag returns tag.GapFillFlag (123)
func (f GapFillFlagField) Tag() quickfix.Tag { return tag.GapFillFlag }

//NewGapFillFlag returns a new GapFillFlagField initialized with val
func NewGapFillFlag(val bool) *GapFillFlagField {
	field := &GapFillFlagField{}
	field.Value = val
	return field
}

//GrossTradeAmtField is a AMT field
type GrossTradeAmtField struct{ quickfix.AmtValue }

//Tag returns tag.GrossTradeAmt (381)
func (f GrossTradeAmtField) Tag() quickfix.Tag { return tag.GrossTradeAmt }

//NewGrossTradeAmt returns a new GrossTradeAmtField initialized with val
func NewGrossTradeAmt(val float64) *GrossTradeAmtField {
	field := &GrossTradeAmtField{}
	field.Value = val
	return field
}

//HaltReasonCharField is a CHAR field
type HaltReasonCharField struct{ quickfix.CharValue }

//Tag returns tag.HaltReasonChar (327)
func (f HaltReasonCharField) Tag() quickfix.Tag { return tag.HaltReasonChar }

//NewHaltReasonChar returns a new HaltReasonCharField initialized with val
func NewHaltReasonChar(val string) *HaltReasonCharField {
	field := &HaltReasonCharField{}
	field.Value = val
	return field
}

//HaltReasonIntField is a INT field
type HaltReasonIntField struct{ quickfix.IntValue }

//Tag returns tag.HaltReasonInt (327)
func (f HaltReasonIntField) Tag() quickfix.Tag { return tag.HaltReasonInt }

//NewHaltReasonInt returns a new HaltReasonIntField initialized with val
func NewHaltReasonInt(val int) *HaltReasonIntField {
	field := &HaltReasonIntField{}
	field.Value = val
	return field
}

//HandlInstField is a CHAR field
type HandlInstField struct{ quickfix.CharValue }

//Tag returns tag.HandlInst (21)
func (f HandlInstField) Tag() quickfix.Tag { return tag.HandlInst }

//NewHandlInst returns a new HandlInstField initialized with val
func NewHandlInst(val string) *HandlInstField {
	field := &HandlInstField{}
	field.Value = val
	return field
}

//HeadlineField is a STRING field
type HeadlineField struct{ quickfix.StringValue }

//Tag returns tag.Headline (148)
func (f HeadlineField) Tag() quickfix.Tag { return tag.Headline }

//NewHeadline returns a new HeadlineField initialized with val
func NewHeadline(val string) *HeadlineField {
	field := &HeadlineField{}
	field.Value = val
	return field
}

//HeartBtIntField is a INT field
type HeartBtIntField struct{ quickfix.IntValue }

//Tag returns tag.HeartBtInt (108)
func (f HeartBtIntField) Tag() quickfix.Tag { return tag.HeartBtInt }

//NewHeartBtInt returns a new HeartBtIntField initialized with val
func NewHeartBtInt(val int) *HeartBtIntField {
	field := &HeartBtIntField{}
	field.Value = val
	return field
}

//HighLimitPriceField is a PRICE field
type HighLimitPriceField struct{ quickfix.PriceValue }

//Tag returns tag.HighLimitPrice (1149)
func (f HighLimitPriceField) Tag() quickfix.Tag { return tag.HighLimitPrice }

//NewHighLimitPrice returns a new HighLimitPriceField initialized with val
func NewHighLimitPrice(val float64) *HighLimitPriceField {
	field := &HighLimitPriceField{}
	field.Value = val
	return field
}

//HighPxField is a PRICE field
type HighPxField struct{ quickfix.PriceValue }

//Tag returns tag.HighPx (332)
func (f HighPxField) Tag() quickfix.Tag { return tag.HighPx }

//NewHighPx returns a new HighPxField initialized with val
func NewHighPx(val float64) *HighPxField {
	field := &HighPxField{}
	field.Value = val
	return field
}

//HopCompIDField is a STRING field
type HopCompIDField struct{ quickfix.StringValue }

//Tag returns tag.HopCompID (628)
func (f HopCompIDField) Tag() quickfix.Tag { return tag.HopCompID }

//NewHopCompID returns a new HopCompIDField initialized with val
func NewHopCompID(val string) *HopCompIDField {
	field := &HopCompIDField{}
	field.Value = val
	return field
}

//HopRefIDField is a SEQNUM field
type HopRefIDField struct{ quickfix.SeqNumValue }

//Tag returns tag.HopRefID (630)
func (f HopRefIDField) Tag() quickfix.Tag { return tag.HopRefID }

//NewHopRefID returns a new HopRefIDField initialized with val
func NewHopRefID(val int) *HopRefIDField {
	field := &HopRefIDField{}
	field.Value = val
	return field
}

//HopSendingTimeField is a UTCTIMESTAMP field
type HopSendingTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.HopSendingTime (629)
func (f HopSendingTimeField) Tag() quickfix.Tag { return tag.HopSendingTime }

//HostCrossIDField is a STRING field
type HostCrossIDField struct{ quickfix.StringValue }

//Tag returns tag.HostCrossID (961)
func (f HostCrossIDField) Tag() quickfix.Tag { return tag.HostCrossID }

//NewHostCrossID returns a new HostCrossIDField initialized with val
func NewHostCrossID(val string) *HostCrossIDField {
	field := &HostCrossIDField{}
	field.Value = val
	return field
}

//IDSourceField is a STRING field
type IDSourceField struct{ quickfix.StringValue }

//Tag returns tag.IDSource (22)
func (f IDSourceField) Tag() quickfix.Tag { return tag.IDSource }

//NewIDSource returns a new IDSourceField initialized with val
func NewIDSource(val string) *IDSourceField {
	field := &IDSourceField{}
	field.Value = val
	return field
}

//IOIIDField is a STRING field
type IOIIDField struct{ quickfix.StringValue }

//Tag returns tag.IOIID (23)
func (f IOIIDField) Tag() quickfix.Tag { return tag.IOIID }

//NewIOIID returns a new IOIIDField initialized with val
func NewIOIID(val string) *IOIIDField {
	field := &IOIIDField{}
	field.Value = val
	return field
}

//IOINaturalFlagField is a BOOLEAN field
type IOINaturalFlagField struct{ quickfix.BooleanValue }

//Tag returns tag.IOINaturalFlag (130)
func (f IOINaturalFlagField) Tag() quickfix.Tag { return tag.IOINaturalFlag }

//NewIOINaturalFlag returns a new IOINaturalFlagField initialized with val
func NewIOINaturalFlag(val bool) *IOINaturalFlagField {
	field := &IOINaturalFlagField{}
	field.Value = val
	return field
}

//IOIOthSvcField is a CHAR field
type IOIOthSvcField struct{ quickfix.CharValue }

//Tag returns tag.IOIOthSvc (24)
func (f IOIOthSvcField) Tag() quickfix.Tag { return tag.IOIOthSvc }

//NewIOIOthSvc returns a new IOIOthSvcField initialized with val
func NewIOIOthSvc(val string) *IOIOthSvcField {
	field := &IOIOthSvcField{}
	field.Value = val
	return field
}

//IOIQltyIndField is a CHAR field
type IOIQltyIndField struct{ quickfix.CharValue }

//Tag returns tag.IOIQltyInd (25)
func (f IOIQltyIndField) Tag() quickfix.Tag { return tag.IOIQltyInd }

//NewIOIQltyInd returns a new IOIQltyIndField initialized with val
func NewIOIQltyInd(val string) *IOIQltyIndField {
	field := &IOIQltyIndField{}
	field.Value = val
	return field
}

//IOIQtyField is a STRING field
type IOIQtyField struct{ quickfix.StringValue }

//Tag returns tag.IOIQty (27)
func (f IOIQtyField) Tag() quickfix.Tag { return tag.IOIQty }

//NewIOIQty returns a new IOIQtyField initialized with val
func NewIOIQty(val string) *IOIQtyField {
	field := &IOIQtyField{}
	field.Value = val
	return field
}

//IOIQualifierField is a CHAR field
type IOIQualifierField struct{ quickfix.CharValue }

//Tag returns tag.IOIQualifier (104)
func (f IOIQualifierField) Tag() quickfix.Tag { return tag.IOIQualifier }

//NewIOIQualifier returns a new IOIQualifierField initialized with val
func NewIOIQualifier(val string) *IOIQualifierField {
	field := &IOIQualifierField{}
	field.Value = val
	return field
}

//IOIRefIDField is a STRING field
type IOIRefIDField struct{ quickfix.StringValue }

//Tag returns tag.IOIRefID (26)
func (f IOIRefIDField) Tag() quickfix.Tag { return tag.IOIRefID }

//NewIOIRefID returns a new IOIRefIDField initialized with val
func NewIOIRefID(val string) *IOIRefIDField {
	field := &IOIRefIDField{}
	field.Value = val
	return field
}

//IOISharesField is a STRING field
type IOISharesField struct{ quickfix.StringValue }

//Tag returns tag.IOIShares (27)
func (f IOISharesField) Tag() quickfix.Tag { return tag.IOIShares }

//NewIOIShares returns a new IOISharesField initialized with val
func NewIOIShares(val string) *IOISharesField {
	field := &IOISharesField{}
	field.Value = val
	return field
}

//IOITransTypeField is a CHAR field
type IOITransTypeField struct{ quickfix.CharValue }

//Tag returns tag.IOITransType (28)
func (f IOITransTypeField) Tag() quickfix.Tag { return tag.IOITransType }

//NewIOITransType returns a new IOITransTypeField initialized with val
func NewIOITransType(val string) *IOITransTypeField {
	field := &IOITransTypeField{}
	field.Value = val
	return field
}

//IOIidField is a STRING field
type IOIidField struct{ quickfix.StringValue }

//Tag returns tag.IOIid (23)
func (f IOIidField) Tag() quickfix.Tag { return tag.IOIid }

//NewIOIid returns a new IOIidField initialized with val
func NewIOIid(val string) *IOIidField {
	field := &IOIidField{}
	field.Value = val
	return field
}

//ImpliedMarketIndicatorField is a INT field
type ImpliedMarketIndicatorField struct{ quickfix.IntValue }

//Tag returns tag.ImpliedMarketIndicator (1144)
func (f ImpliedMarketIndicatorField) Tag() quickfix.Tag { return tag.ImpliedMarketIndicator }

//NewImpliedMarketIndicator returns a new ImpliedMarketIndicatorField initialized with val
func NewImpliedMarketIndicator(val int) *ImpliedMarketIndicatorField {
	field := &ImpliedMarketIndicatorField{}
	field.Value = val
	return field
}

//InViewOfCommonField is a BOOLEAN field
type InViewOfCommonField struct{ quickfix.BooleanValue }

//Tag returns tag.InViewOfCommon (328)
func (f InViewOfCommonField) Tag() quickfix.Tag { return tag.InViewOfCommon }

//NewInViewOfCommon returns a new InViewOfCommonField initialized with val
func NewInViewOfCommon(val bool) *InViewOfCommonField {
	field := &InViewOfCommonField{}
	field.Value = val
	return field
}

//IncTaxIndField is a INT field
type IncTaxIndField struct{ quickfix.IntValue }

//Tag returns tag.IncTaxInd (416)
func (f IncTaxIndField) Tag() quickfix.Tag { return tag.IncTaxInd }

//NewIncTaxInd returns a new IncTaxIndField initialized with val
func NewIncTaxInd(val int) *IncTaxIndField {
	field := &IncTaxIndField{}
	field.Value = val
	return field
}

//IndividualAllocIDField is a STRING field
type IndividualAllocIDField struct{ quickfix.StringValue }

//Tag returns tag.IndividualAllocID (467)
func (f IndividualAllocIDField) Tag() quickfix.Tag { return tag.IndividualAllocID }

//NewIndividualAllocID returns a new IndividualAllocIDField initialized with val
func NewIndividualAllocID(val string) *IndividualAllocIDField {
	field := &IndividualAllocIDField{}
	field.Value = val
	return field
}

//IndividualAllocRejCodeField is a INT field
type IndividualAllocRejCodeField struct{ quickfix.IntValue }

//Tag returns tag.IndividualAllocRejCode (776)
func (f IndividualAllocRejCodeField) Tag() quickfix.Tag { return tag.IndividualAllocRejCode }

//NewIndividualAllocRejCode returns a new IndividualAllocRejCodeField initialized with val
func NewIndividualAllocRejCode(val int) *IndividualAllocRejCodeField {
	field := &IndividualAllocRejCodeField{}
	field.Value = val
	return field
}

//IndividualAllocTypeField is a INT field
type IndividualAllocTypeField struct{ quickfix.IntValue }

//Tag returns tag.IndividualAllocType (992)
func (f IndividualAllocTypeField) Tag() quickfix.Tag { return tag.IndividualAllocType }

//NewIndividualAllocType returns a new IndividualAllocTypeField initialized with val
func NewIndividualAllocType(val int) *IndividualAllocTypeField {
	field := &IndividualAllocTypeField{}
	field.Value = val
	return field
}

//InputSourceField is a STRING field
type InputSourceField struct{ quickfix.StringValue }

//Tag returns tag.InputSource (979)
func (f InputSourceField) Tag() quickfix.Tag { return tag.InputSource }

//NewInputSource returns a new InputSourceField initialized with val
func NewInputSource(val string) *InputSourceField {
	field := &InputSourceField{}
	field.Value = val
	return field
}

//InstrAttribTypeField is a INT field
type InstrAttribTypeField struct{ quickfix.IntValue }

//Tag returns tag.InstrAttribType (871)
func (f InstrAttribTypeField) Tag() quickfix.Tag { return tag.InstrAttribType }

//NewInstrAttribType returns a new InstrAttribTypeField initialized with val
func NewInstrAttribType(val int) *InstrAttribTypeField {
	field := &InstrAttribTypeField{}
	field.Value = val
	return field
}

//InstrAttribValueField is a STRING field
type InstrAttribValueField struct{ quickfix.StringValue }

//Tag returns tag.InstrAttribValue (872)
func (f InstrAttribValueField) Tag() quickfix.Tag { return tag.InstrAttribValue }

//NewInstrAttribValue returns a new InstrAttribValueField initialized with val
func NewInstrAttribValue(val string) *InstrAttribValueField {
	field := &InstrAttribValueField{}
	field.Value = val
	return field
}

//InstrRegistryField is a STRING field
type InstrRegistryField struct{ quickfix.StringValue }

//Tag returns tag.InstrRegistry (543)
func (f InstrRegistryField) Tag() quickfix.Tag { return tag.InstrRegistry }

//NewInstrRegistry returns a new InstrRegistryField initialized with val
func NewInstrRegistry(val string) *InstrRegistryField {
	field := &InstrRegistryField{}
	field.Value = val
	return field
}

//InstrmtAssignmentMethodField is a CHAR field
type InstrmtAssignmentMethodField struct{ quickfix.CharValue }

//Tag returns tag.InstrmtAssignmentMethod (1049)
func (f InstrmtAssignmentMethodField) Tag() quickfix.Tag { return tag.InstrmtAssignmentMethod }

//NewInstrmtAssignmentMethod returns a new InstrmtAssignmentMethodField initialized with val
func NewInstrmtAssignmentMethod(val string) *InstrmtAssignmentMethodField {
	field := &InstrmtAssignmentMethodField{}
	field.Value = val
	return field
}

//InstrumentPartyIDField is a STRING field
type InstrumentPartyIDField struct{ quickfix.StringValue }

//Tag returns tag.InstrumentPartyID (1019)
func (f InstrumentPartyIDField) Tag() quickfix.Tag { return tag.InstrumentPartyID }

//NewInstrumentPartyID returns a new InstrumentPartyIDField initialized with val
func NewInstrumentPartyID(val string) *InstrumentPartyIDField {
	field := &InstrumentPartyIDField{}
	field.Value = val
	return field
}

//InstrumentPartyIDSourceField is a CHAR field
type InstrumentPartyIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.InstrumentPartyIDSource (1050)
func (f InstrumentPartyIDSourceField) Tag() quickfix.Tag { return tag.InstrumentPartyIDSource }

//NewInstrumentPartyIDSource returns a new InstrumentPartyIDSourceField initialized with val
func NewInstrumentPartyIDSource(val string) *InstrumentPartyIDSourceField {
	field := &InstrumentPartyIDSourceField{}
	field.Value = val
	return field
}

//InstrumentPartyRoleField is a INT field
type InstrumentPartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.InstrumentPartyRole (1051)
func (f InstrumentPartyRoleField) Tag() quickfix.Tag { return tag.InstrumentPartyRole }

//NewInstrumentPartyRole returns a new InstrumentPartyRoleField initialized with val
func NewInstrumentPartyRole(val int) *InstrumentPartyRoleField {
	field := &InstrumentPartyRoleField{}
	field.Value = val
	return field
}

//InstrumentPartySubIDField is a STRING field
type InstrumentPartySubIDField struct{ quickfix.StringValue }

//Tag returns tag.InstrumentPartySubID (1053)
func (f InstrumentPartySubIDField) Tag() quickfix.Tag { return tag.InstrumentPartySubID }

//NewInstrumentPartySubID returns a new InstrumentPartySubIDField initialized with val
func NewInstrumentPartySubID(val string) *InstrumentPartySubIDField {
	field := &InstrumentPartySubIDField{}
	field.Value = val
	return field
}

//InstrumentPartySubIDTypeField is a INT field
type InstrumentPartySubIDTypeField struct{ quickfix.IntValue }

//Tag returns tag.InstrumentPartySubIDType (1054)
func (f InstrumentPartySubIDTypeField) Tag() quickfix.Tag { return tag.InstrumentPartySubIDType }

//NewInstrumentPartySubIDType returns a new InstrumentPartySubIDTypeField initialized with val
func NewInstrumentPartySubIDType(val int) *InstrumentPartySubIDTypeField {
	field := &InstrumentPartySubIDTypeField{}
	field.Value = val
	return field
}

//InterestAccrualDateField is a LOCALMKTDATE field
type InterestAccrualDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.InterestAccrualDate (874)
func (f InterestAccrualDateField) Tag() quickfix.Tag { return tag.InterestAccrualDate }

//NewInterestAccrualDate returns a new InterestAccrualDateField initialized with val
func NewInterestAccrualDate(val string) *InterestAccrualDateField {
	field := &InterestAccrualDateField{}
	field.Value = val
	return field
}

//InterestAtMaturityField is a AMT field
type InterestAtMaturityField struct{ quickfix.AmtValue }

//Tag returns tag.InterestAtMaturity (738)
func (f InterestAtMaturityField) Tag() quickfix.Tag { return tag.InterestAtMaturity }

//NewInterestAtMaturity returns a new InterestAtMaturityField initialized with val
func NewInterestAtMaturity(val float64) *InterestAtMaturityField {
	field := &InterestAtMaturityField{}
	field.Value = val
	return field
}

//InvestorCountryOfResidenceField is a COUNTRY field
type InvestorCountryOfResidenceField struct{ quickfix.CountryValue }

//Tag returns tag.InvestorCountryOfResidence (475)
func (f InvestorCountryOfResidenceField) Tag() quickfix.Tag { return tag.InvestorCountryOfResidence }

//NewInvestorCountryOfResidence returns a new InvestorCountryOfResidenceField initialized with val
func NewInvestorCountryOfResidence(val string) *InvestorCountryOfResidenceField {
	field := &InvestorCountryOfResidenceField{}
	field.Value = val
	return field
}

//IssueDateField is a LOCALMKTDATE field
type IssueDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.IssueDate (225)
func (f IssueDateField) Tag() quickfix.Tag { return tag.IssueDate }

//NewIssueDate returns a new IssueDateField initialized with val
func NewIssueDate(val string) *IssueDateField {
	field := &IssueDateField{}
	field.Value = val
	return field
}

//IssuerField is a STRING field
type IssuerField struct{ quickfix.StringValue }

//Tag returns tag.Issuer (106)
func (f IssuerField) Tag() quickfix.Tag { return tag.Issuer }

//NewIssuer returns a new IssuerField initialized with val
func NewIssuer(val string) *IssuerField {
	field := &IssuerField{}
	field.Value = val
	return field
}

//LanguageCodeField is a LANGUAGE field
type LanguageCodeField struct{ quickfix.LanguageValue }

//Tag returns tag.LanguageCode (1474)
func (f LanguageCodeField) Tag() quickfix.Tag { return tag.LanguageCode }

//NewLanguageCode returns a new LanguageCodeField initialized with val
func NewLanguageCode(val string) *LanguageCodeField {
	field := &LanguageCodeField{}
	field.Value = val
	return field
}

//LastCapacityField is a CHAR field
type LastCapacityField struct{ quickfix.CharValue }

//Tag returns tag.LastCapacity (29)
func (f LastCapacityField) Tag() quickfix.Tag { return tag.LastCapacity }

//NewLastCapacity returns a new LastCapacityField initialized with val
func NewLastCapacity(val string) *LastCapacityField {
	field := &LastCapacityField{}
	field.Value = val
	return field
}

//LastForwardPointsField is a PRICEOFFSET field
type LastForwardPointsField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.LastForwardPoints (195)
func (f LastForwardPointsField) Tag() quickfix.Tag { return tag.LastForwardPoints }

//NewLastForwardPoints returns a new LastForwardPointsField initialized with val
func NewLastForwardPoints(val float64) *LastForwardPointsField {
	field := &LastForwardPointsField{}
	field.Value = val
	return field
}

//LastForwardPoints2Field is a PRICEOFFSET field
type LastForwardPoints2Field struct{ quickfix.PriceOffsetValue }

//Tag returns tag.LastForwardPoints2 (641)
func (f LastForwardPoints2Field) Tag() quickfix.Tag { return tag.LastForwardPoints2 }

//NewLastForwardPoints2 returns a new LastForwardPoints2Field initialized with val
func NewLastForwardPoints2(val float64) *LastForwardPoints2Field {
	field := &LastForwardPoints2Field{}
	field.Value = val
	return field
}

//LastFragmentField is a BOOLEAN field
type LastFragmentField struct{ quickfix.BooleanValue }

//Tag returns tag.LastFragment (893)
func (f LastFragmentField) Tag() quickfix.Tag { return tag.LastFragment }

//NewLastFragment returns a new LastFragmentField initialized with val
func NewLastFragment(val bool) *LastFragmentField {
	field := &LastFragmentField{}
	field.Value = val
	return field
}

//LastLiquidityIndField is a INT field
type LastLiquidityIndField struct{ quickfix.IntValue }

//Tag returns tag.LastLiquidityInd (851)
func (f LastLiquidityIndField) Tag() quickfix.Tag { return tag.LastLiquidityInd }

//NewLastLiquidityInd returns a new LastLiquidityIndField initialized with val
func NewLastLiquidityInd(val int) *LastLiquidityIndField {
	field := &LastLiquidityIndField{}
	field.Value = val
	return field
}

//LastMktField is a EXCHANGE field
type LastMktField struct{ quickfix.ExchangeValue }

//Tag returns tag.LastMkt (30)
func (f LastMktField) Tag() quickfix.Tag { return tag.LastMkt }

//NewLastMkt returns a new LastMktField initialized with val
func NewLastMkt(val string) *LastMktField {
	field := &LastMktField{}
	field.Value = val
	return field
}

//LastMsgSeqNumProcessedField is a SEQNUM field
type LastMsgSeqNumProcessedField struct{ quickfix.SeqNumValue }

//Tag returns tag.LastMsgSeqNumProcessed (369)
func (f LastMsgSeqNumProcessedField) Tag() quickfix.Tag { return tag.LastMsgSeqNumProcessed }

//NewLastMsgSeqNumProcessed returns a new LastMsgSeqNumProcessedField initialized with val
func NewLastMsgSeqNumProcessed(val int) *LastMsgSeqNumProcessedField {
	field := &LastMsgSeqNumProcessedField{}
	field.Value = val
	return field
}

//LastNetworkResponseIDField is a STRING field
type LastNetworkResponseIDField struct{ quickfix.StringValue }

//Tag returns tag.LastNetworkResponseID (934)
func (f LastNetworkResponseIDField) Tag() quickfix.Tag { return tag.LastNetworkResponseID }

//NewLastNetworkResponseID returns a new LastNetworkResponseIDField initialized with val
func NewLastNetworkResponseID(val string) *LastNetworkResponseIDField {
	field := &LastNetworkResponseIDField{}
	field.Value = val
	return field
}

//LastParPxField is a PRICE field
type LastParPxField struct{ quickfix.PriceValue }

//Tag returns tag.LastParPx (669)
func (f LastParPxField) Tag() quickfix.Tag { return tag.LastParPx }

//NewLastParPx returns a new LastParPxField initialized with val
func NewLastParPx(val float64) *LastParPxField {
	field := &LastParPxField{}
	field.Value = val
	return field
}

//LastPxField is a PRICE field
type LastPxField struct{ quickfix.PriceValue }

//Tag returns tag.LastPx (31)
func (f LastPxField) Tag() quickfix.Tag { return tag.LastPx }

//NewLastPx returns a new LastPxField initialized with val
func NewLastPx(val float64) *LastPxField {
	field := &LastPxField{}
	field.Value = val
	return field
}

//LastQtyField is a QTY field
type LastQtyField struct{ quickfix.QtyValue }

//Tag returns tag.LastQty (32)
func (f LastQtyField) Tag() quickfix.Tag { return tag.LastQty }

//NewLastQty returns a new LastQtyField initialized with val
func NewLastQty(val float64) *LastQtyField {
	field := &LastQtyField{}
	field.Value = val
	return field
}

//LastRptRequestedField is a BOOLEAN field
type LastRptRequestedField struct{ quickfix.BooleanValue }

//Tag returns tag.LastRptRequested (912)
func (f LastRptRequestedField) Tag() quickfix.Tag { return tag.LastRptRequested }

//NewLastRptRequested returns a new LastRptRequestedField initialized with val
func NewLastRptRequested(val bool) *LastRptRequestedField {
	field := &LastRptRequestedField{}
	field.Value = val
	return field
}

//LastSharesField is a QTY field
type LastSharesField struct{ quickfix.QtyValue }

//Tag returns tag.LastShares (32)
func (f LastSharesField) Tag() quickfix.Tag { return tag.LastShares }

//NewLastShares returns a new LastSharesField initialized with val
func NewLastShares(val float64) *LastSharesField {
	field := &LastSharesField{}
	field.Value = val
	return field
}

//LastSpotRateField is a PRICE field
type LastSpotRateField struct{ quickfix.PriceValue }

//Tag returns tag.LastSpotRate (194)
func (f LastSpotRateField) Tag() quickfix.Tag { return tag.LastSpotRate }

//NewLastSpotRate returns a new LastSpotRateField initialized with val
func NewLastSpotRate(val float64) *LastSpotRateField {
	field := &LastSpotRateField{}
	field.Value = val
	return field
}

//LastSwapPointsField is a PRICEOFFSET field
type LastSwapPointsField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.LastSwapPoints (1071)
func (f LastSwapPointsField) Tag() quickfix.Tag { return tag.LastSwapPoints }

//NewLastSwapPoints returns a new LastSwapPointsField initialized with val
func NewLastSwapPoints(val float64) *LastSwapPointsField {
	field := &LastSwapPointsField{}
	field.Value = val
	return field
}

//LastUpdateTimeField is a UTCTIMESTAMP field
type LastUpdateTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.LastUpdateTime (779)
func (f LastUpdateTimeField) Tag() quickfix.Tag { return tag.LastUpdateTime }

//LateIndicatorField is a BOOLEAN field
type LateIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.LateIndicator (978)
func (f LateIndicatorField) Tag() quickfix.Tag { return tag.LateIndicator }

//NewLateIndicator returns a new LateIndicatorField initialized with val
func NewLateIndicator(val bool) *LateIndicatorField {
	field := &LateIndicatorField{}
	field.Value = val
	return field
}

//LeavesQtyField is a QTY field
type LeavesQtyField struct{ quickfix.QtyValue }

//Tag returns tag.LeavesQty (151)
func (f LeavesQtyField) Tag() quickfix.Tag { return tag.LeavesQty }

//NewLeavesQty returns a new LeavesQtyField initialized with val
func NewLeavesQty(val float64) *LeavesQtyField {
	field := &LeavesQtyField{}
	field.Value = val
	return field
}

//LegAllocAccountField is a STRING field
type LegAllocAccountField struct{ quickfix.StringValue }

//Tag returns tag.LegAllocAccount (671)
func (f LegAllocAccountField) Tag() quickfix.Tag { return tag.LegAllocAccount }

//NewLegAllocAccount returns a new LegAllocAccountField initialized with val
func NewLegAllocAccount(val string) *LegAllocAccountField {
	field := &LegAllocAccountField{}
	field.Value = val
	return field
}

//LegAllocAcctIDSourceField is a STRING field
type LegAllocAcctIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.LegAllocAcctIDSource (674)
func (f LegAllocAcctIDSourceField) Tag() quickfix.Tag { return tag.LegAllocAcctIDSource }

//NewLegAllocAcctIDSource returns a new LegAllocAcctIDSourceField initialized with val
func NewLegAllocAcctIDSource(val string) *LegAllocAcctIDSourceField {
	field := &LegAllocAcctIDSourceField{}
	field.Value = val
	return field
}

//LegAllocIDField is a STRING field
type LegAllocIDField struct{ quickfix.StringValue }

//Tag returns tag.LegAllocID (1366)
func (f LegAllocIDField) Tag() quickfix.Tag { return tag.LegAllocID }

//NewLegAllocID returns a new LegAllocIDField initialized with val
func NewLegAllocID(val string) *LegAllocIDField {
	field := &LegAllocIDField{}
	field.Value = val
	return field
}

//LegAllocQtyField is a QTY field
type LegAllocQtyField struct{ quickfix.QtyValue }

//Tag returns tag.LegAllocQty (673)
func (f LegAllocQtyField) Tag() quickfix.Tag { return tag.LegAllocQty }

//NewLegAllocQty returns a new LegAllocQtyField initialized with val
func NewLegAllocQty(val float64) *LegAllocQtyField {
	field := &LegAllocQtyField{}
	field.Value = val
	return field
}

//LegAllocSettlCurrencyField is a CURRENCY field
type LegAllocSettlCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.LegAllocSettlCurrency (1367)
func (f LegAllocSettlCurrencyField) Tag() quickfix.Tag { return tag.LegAllocSettlCurrency }

//NewLegAllocSettlCurrency returns a new LegAllocSettlCurrencyField initialized with val
func NewLegAllocSettlCurrency(val string) *LegAllocSettlCurrencyField {
	field := &LegAllocSettlCurrencyField{}
	field.Value = val
	return field
}

//LegBenchmarkCurveCurrencyField is a CURRENCY field
type LegBenchmarkCurveCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.LegBenchmarkCurveCurrency (676)
func (f LegBenchmarkCurveCurrencyField) Tag() quickfix.Tag { return tag.LegBenchmarkCurveCurrency }

//NewLegBenchmarkCurveCurrency returns a new LegBenchmarkCurveCurrencyField initialized with val
func NewLegBenchmarkCurveCurrency(val string) *LegBenchmarkCurveCurrencyField {
	field := &LegBenchmarkCurveCurrencyField{}
	field.Value = val
	return field
}

//LegBenchmarkCurveNameField is a STRING field
type LegBenchmarkCurveNameField struct{ quickfix.StringValue }

//Tag returns tag.LegBenchmarkCurveName (677)
func (f LegBenchmarkCurveNameField) Tag() quickfix.Tag { return tag.LegBenchmarkCurveName }

//NewLegBenchmarkCurveName returns a new LegBenchmarkCurveNameField initialized with val
func NewLegBenchmarkCurveName(val string) *LegBenchmarkCurveNameField {
	field := &LegBenchmarkCurveNameField{}
	field.Value = val
	return field
}

//LegBenchmarkCurvePointField is a STRING field
type LegBenchmarkCurvePointField struct{ quickfix.StringValue }

//Tag returns tag.LegBenchmarkCurvePoint (678)
func (f LegBenchmarkCurvePointField) Tag() quickfix.Tag { return tag.LegBenchmarkCurvePoint }

//NewLegBenchmarkCurvePoint returns a new LegBenchmarkCurvePointField initialized with val
func NewLegBenchmarkCurvePoint(val string) *LegBenchmarkCurvePointField {
	field := &LegBenchmarkCurvePointField{}
	field.Value = val
	return field
}

//LegBenchmarkPriceField is a PRICE field
type LegBenchmarkPriceField struct{ quickfix.PriceValue }

//Tag returns tag.LegBenchmarkPrice (679)
func (f LegBenchmarkPriceField) Tag() quickfix.Tag { return tag.LegBenchmarkPrice }

//NewLegBenchmarkPrice returns a new LegBenchmarkPriceField initialized with val
func NewLegBenchmarkPrice(val float64) *LegBenchmarkPriceField {
	field := &LegBenchmarkPriceField{}
	field.Value = val
	return field
}

//LegBenchmarkPriceTypeField is a INT field
type LegBenchmarkPriceTypeField struct{ quickfix.IntValue }

//Tag returns tag.LegBenchmarkPriceType (680)
func (f LegBenchmarkPriceTypeField) Tag() quickfix.Tag { return tag.LegBenchmarkPriceType }

//NewLegBenchmarkPriceType returns a new LegBenchmarkPriceTypeField initialized with val
func NewLegBenchmarkPriceType(val int) *LegBenchmarkPriceTypeField {
	field := &LegBenchmarkPriceTypeField{}
	field.Value = val
	return field
}

//LegBidForwardPointsField is a PRICEOFFSET field
type LegBidForwardPointsField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.LegBidForwardPoints (1067)
func (f LegBidForwardPointsField) Tag() quickfix.Tag { return tag.LegBidForwardPoints }

//NewLegBidForwardPoints returns a new LegBidForwardPointsField initialized with val
func NewLegBidForwardPoints(val float64) *LegBidForwardPointsField {
	field := &LegBidForwardPointsField{}
	field.Value = val
	return field
}

//LegBidPxField is a PRICE field
type LegBidPxField struct{ quickfix.PriceValue }

//Tag returns tag.LegBidPx (681)
func (f LegBidPxField) Tag() quickfix.Tag { return tag.LegBidPx }

//NewLegBidPx returns a new LegBidPxField initialized with val
func NewLegBidPx(val float64) *LegBidPxField {
	field := &LegBidPxField{}
	field.Value = val
	return field
}

//LegCFICodeField is a STRING field
type LegCFICodeField struct{ quickfix.StringValue }

//Tag returns tag.LegCFICode (608)
func (f LegCFICodeField) Tag() quickfix.Tag { return tag.LegCFICode }

//NewLegCFICode returns a new LegCFICodeField initialized with val
func NewLegCFICode(val string) *LegCFICodeField {
	field := &LegCFICodeField{}
	field.Value = val
	return field
}

//LegCalculatedCcyLastQtyField is a QTY field
type LegCalculatedCcyLastQtyField struct{ quickfix.QtyValue }

//Tag returns tag.LegCalculatedCcyLastQty (1074)
func (f LegCalculatedCcyLastQtyField) Tag() quickfix.Tag { return tag.LegCalculatedCcyLastQty }

//NewLegCalculatedCcyLastQty returns a new LegCalculatedCcyLastQtyField initialized with val
func NewLegCalculatedCcyLastQty(val float64) *LegCalculatedCcyLastQtyField {
	field := &LegCalculatedCcyLastQtyField{}
	field.Value = val
	return field
}

//LegContractMultiplierField is a FLOAT field
type LegContractMultiplierField struct{ quickfix.FloatValue }

//Tag returns tag.LegContractMultiplier (614)
func (f LegContractMultiplierField) Tag() quickfix.Tag { return tag.LegContractMultiplier }

//NewLegContractMultiplier returns a new LegContractMultiplierField initialized with val
func NewLegContractMultiplier(val float64) *LegContractMultiplierField {
	field := &LegContractMultiplierField{}
	field.Value = val
	return field
}

//LegContractMultiplierUnitField is a INT field
type LegContractMultiplierUnitField struct{ quickfix.IntValue }

//Tag returns tag.LegContractMultiplierUnit (1436)
func (f LegContractMultiplierUnitField) Tag() quickfix.Tag { return tag.LegContractMultiplierUnit }

//NewLegContractMultiplierUnit returns a new LegContractMultiplierUnitField initialized with val
func NewLegContractMultiplierUnit(val int) *LegContractMultiplierUnitField {
	field := &LegContractMultiplierUnitField{}
	field.Value = val
	return field
}

//LegContractSettlMonthField is a MONTHYEAR field
type LegContractSettlMonthField struct{ quickfix.MonthYearValue }

//Tag returns tag.LegContractSettlMonth (955)
func (f LegContractSettlMonthField) Tag() quickfix.Tag { return tag.LegContractSettlMonth }

//NewLegContractSettlMonth returns a new LegContractSettlMonthField initialized with val
func NewLegContractSettlMonth(val string) *LegContractSettlMonthField {
	field := &LegContractSettlMonthField{}
	field.Value = val
	return field
}

//LegCountryOfIssueField is a COUNTRY field
type LegCountryOfIssueField struct{ quickfix.CountryValue }

//Tag returns tag.LegCountryOfIssue (596)
func (f LegCountryOfIssueField) Tag() quickfix.Tag { return tag.LegCountryOfIssue }

//NewLegCountryOfIssue returns a new LegCountryOfIssueField initialized with val
func NewLegCountryOfIssue(val string) *LegCountryOfIssueField {
	field := &LegCountryOfIssueField{}
	field.Value = val
	return field
}

//LegCouponPaymentDateField is a LOCALMKTDATE field
type LegCouponPaymentDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.LegCouponPaymentDate (248)
func (f LegCouponPaymentDateField) Tag() quickfix.Tag { return tag.LegCouponPaymentDate }

//NewLegCouponPaymentDate returns a new LegCouponPaymentDateField initialized with val
func NewLegCouponPaymentDate(val string) *LegCouponPaymentDateField {
	field := &LegCouponPaymentDateField{}
	field.Value = val
	return field
}

//LegCouponRateField is a PERCENTAGE field
type LegCouponRateField struct{ quickfix.PercentageValue }

//Tag returns tag.LegCouponRate (615)
func (f LegCouponRateField) Tag() quickfix.Tag { return tag.LegCouponRate }

//NewLegCouponRate returns a new LegCouponRateField initialized with val
func NewLegCouponRate(val float64) *LegCouponRateField {
	field := &LegCouponRateField{}
	field.Value = val
	return field
}

//LegCoveredOrUncoveredField is a INT field
type LegCoveredOrUncoveredField struct{ quickfix.IntValue }

//Tag returns tag.LegCoveredOrUncovered (565)
func (f LegCoveredOrUncoveredField) Tag() quickfix.Tag { return tag.LegCoveredOrUncovered }

//NewLegCoveredOrUncovered returns a new LegCoveredOrUncoveredField initialized with val
func NewLegCoveredOrUncovered(val int) *LegCoveredOrUncoveredField {
	field := &LegCoveredOrUncoveredField{}
	field.Value = val
	return field
}

//LegCreditRatingField is a STRING field
type LegCreditRatingField struct{ quickfix.StringValue }

//Tag returns tag.LegCreditRating (257)
func (f LegCreditRatingField) Tag() quickfix.Tag { return tag.LegCreditRating }

//NewLegCreditRating returns a new LegCreditRatingField initialized with val
func NewLegCreditRating(val string) *LegCreditRatingField {
	field := &LegCreditRatingField{}
	field.Value = val
	return field
}

//LegCurrencyField is a CURRENCY field
type LegCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.LegCurrency (556)
func (f LegCurrencyField) Tag() quickfix.Tag { return tag.LegCurrency }

//NewLegCurrency returns a new LegCurrencyField initialized with val
func NewLegCurrency(val string) *LegCurrencyField {
	field := &LegCurrencyField{}
	field.Value = val
	return field
}

//LegCurrencyRatioField is a FLOAT field
type LegCurrencyRatioField struct{ quickfix.FloatValue }

//Tag returns tag.LegCurrencyRatio (1383)
func (f LegCurrencyRatioField) Tag() quickfix.Tag { return tag.LegCurrencyRatio }

//NewLegCurrencyRatio returns a new LegCurrencyRatioField initialized with val
func NewLegCurrencyRatio(val float64) *LegCurrencyRatioField {
	field := &LegCurrencyRatioField{}
	field.Value = val
	return field
}

//LegDatedDateField is a LOCALMKTDATE field
type LegDatedDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.LegDatedDate (739)
func (f LegDatedDateField) Tag() quickfix.Tag { return tag.LegDatedDate }

//NewLegDatedDate returns a new LegDatedDateField initialized with val
func NewLegDatedDate(val string) *LegDatedDateField {
	field := &LegDatedDateField{}
	field.Value = val
	return field
}

//LegDividendYieldField is a PERCENTAGE field
type LegDividendYieldField struct{ quickfix.PercentageValue }

//Tag returns tag.LegDividendYield (1381)
func (f LegDividendYieldField) Tag() quickfix.Tag { return tag.LegDividendYield }

//NewLegDividendYield returns a new LegDividendYieldField initialized with val
func NewLegDividendYield(val float64) *LegDividendYieldField {
	field := &LegDividendYieldField{}
	field.Value = val
	return field
}

//LegExecInstField is a MULTIPLECHARVALUE field
type LegExecInstField struct{ quickfix.MultipleCharValue }

//Tag returns tag.LegExecInst (1384)
func (f LegExecInstField) Tag() quickfix.Tag { return tag.LegExecInst }

//NewLegExecInst returns a new LegExecInstField initialized with val
func NewLegExecInst(val string) *LegExecInstField {
	field := &LegExecInstField{}
	field.Value = val
	return field
}

//LegExerciseStyleField is a INT field
type LegExerciseStyleField struct{ quickfix.IntValue }

//Tag returns tag.LegExerciseStyle (1420)
func (f LegExerciseStyleField) Tag() quickfix.Tag { return tag.LegExerciseStyle }

//NewLegExerciseStyle returns a new LegExerciseStyleField initialized with val
func NewLegExerciseStyle(val int) *LegExerciseStyleField {
	field := &LegExerciseStyleField{}
	field.Value = val
	return field
}

//LegFactorField is a FLOAT field
type LegFactorField struct{ quickfix.FloatValue }

//Tag returns tag.LegFactor (253)
func (f LegFactorField) Tag() quickfix.Tag { return tag.LegFactor }

//NewLegFactor returns a new LegFactorField initialized with val
func NewLegFactor(val float64) *LegFactorField {
	field := &LegFactorField{}
	field.Value = val
	return field
}

//LegFlowScheduleTypeField is a INT field
type LegFlowScheduleTypeField struct{ quickfix.IntValue }

//Tag returns tag.LegFlowScheduleType (1440)
func (f LegFlowScheduleTypeField) Tag() quickfix.Tag { return tag.LegFlowScheduleType }

//NewLegFlowScheduleType returns a new LegFlowScheduleTypeField initialized with val
func NewLegFlowScheduleType(val int) *LegFlowScheduleTypeField {
	field := &LegFlowScheduleTypeField{}
	field.Value = val
	return field
}

//LegFutSettDateField is a LOCALMKTDATE field
type LegFutSettDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.LegFutSettDate (588)
func (f LegFutSettDateField) Tag() quickfix.Tag { return tag.LegFutSettDate }

//NewLegFutSettDate returns a new LegFutSettDateField initialized with val
func NewLegFutSettDate(val string) *LegFutSettDateField {
	field := &LegFutSettDateField{}
	field.Value = val
	return field
}

//LegGrossTradeAmtField is a AMT field
type LegGrossTradeAmtField struct{ quickfix.AmtValue }

//Tag returns tag.LegGrossTradeAmt (1075)
func (f LegGrossTradeAmtField) Tag() quickfix.Tag { return tag.LegGrossTradeAmt }

//NewLegGrossTradeAmt returns a new LegGrossTradeAmtField initialized with val
func NewLegGrossTradeAmt(val float64) *LegGrossTradeAmtField {
	field := &LegGrossTradeAmtField{}
	field.Value = val
	return field
}

//LegIOIQtyField is a STRING field
type LegIOIQtyField struct{ quickfix.StringValue }

//Tag returns tag.LegIOIQty (682)
func (f LegIOIQtyField) Tag() quickfix.Tag { return tag.LegIOIQty }

//NewLegIOIQty returns a new LegIOIQtyField initialized with val
func NewLegIOIQty(val string) *LegIOIQtyField {
	field := &LegIOIQtyField{}
	field.Value = val
	return field
}

//LegIndividualAllocIDField is a STRING field
type LegIndividualAllocIDField struct{ quickfix.StringValue }

//Tag returns tag.LegIndividualAllocID (672)
func (f LegIndividualAllocIDField) Tag() quickfix.Tag { return tag.LegIndividualAllocID }

//NewLegIndividualAllocID returns a new LegIndividualAllocIDField initialized with val
func NewLegIndividualAllocID(val string) *LegIndividualAllocIDField {
	field := &LegIndividualAllocIDField{}
	field.Value = val
	return field
}

//LegInstrRegistryField is a STRING field
type LegInstrRegistryField struct{ quickfix.StringValue }

//Tag returns tag.LegInstrRegistry (599)
func (f LegInstrRegistryField) Tag() quickfix.Tag { return tag.LegInstrRegistry }

//NewLegInstrRegistry returns a new LegInstrRegistryField initialized with val
func NewLegInstrRegistry(val string) *LegInstrRegistryField {
	field := &LegInstrRegistryField{}
	field.Value = val
	return field
}

//LegInterestAccrualDateField is a LOCALMKTDATE field
type LegInterestAccrualDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.LegInterestAccrualDate (956)
func (f LegInterestAccrualDateField) Tag() quickfix.Tag { return tag.LegInterestAccrualDate }

//NewLegInterestAccrualDate returns a new LegInterestAccrualDateField initialized with val
func NewLegInterestAccrualDate(val string) *LegInterestAccrualDateField {
	field := &LegInterestAccrualDateField{}
	field.Value = val
	return field
}

//LegIssueDateField is a LOCALMKTDATE field
type LegIssueDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.LegIssueDate (249)
func (f LegIssueDateField) Tag() quickfix.Tag { return tag.LegIssueDate }

//NewLegIssueDate returns a new LegIssueDateField initialized with val
func NewLegIssueDate(val string) *LegIssueDateField {
	field := &LegIssueDateField{}
	field.Value = val
	return field
}

//LegIssuerField is a STRING field
type LegIssuerField struct{ quickfix.StringValue }

//Tag returns tag.LegIssuer (617)
func (f LegIssuerField) Tag() quickfix.Tag { return tag.LegIssuer }

//NewLegIssuer returns a new LegIssuerField initialized with val
func NewLegIssuer(val string) *LegIssuerField {
	field := &LegIssuerField{}
	field.Value = val
	return field
}

//LegLastForwardPointsField is a PRICEOFFSET field
type LegLastForwardPointsField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.LegLastForwardPoints (1073)
func (f LegLastForwardPointsField) Tag() quickfix.Tag { return tag.LegLastForwardPoints }

//NewLegLastForwardPoints returns a new LegLastForwardPointsField initialized with val
func NewLegLastForwardPoints(val float64) *LegLastForwardPointsField {
	field := &LegLastForwardPointsField{}
	field.Value = val
	return field
}

//LegLastPxField is a PRICE field
type LegLastPxField struct{ quickfix.PriceValue }

//Tag returns tag.LegLastPx (637)
func (f LegLastPxField) Tag() quickfix.Tag { return tag.LegLastPx }

//NewLegLastPx returns a new LegLastPxField initialized with val
func NewLegLastPx(val float64) *LegLastPxField {
	field := &LegLastPxField{}
	field.Value = val
	return field
}

//LegLastQtyField is a QTY field
type LegLastQtyField struct{ quickfix.QtyValue }

//Tag returns tag.LegLastQty (1418)
func (f LegLastQtyField) Tag() quickfix.Tag { return tag.LegLastQty }

//NewLegLastQty returns a new LegLastQtyField initialized with val
func NewLegLastQty(val float64) *LegLastQtyField {
	field := &LegLastQtyField{}
	field.Value = val
	return field
}

//LegLocaleOfIssueField is a STRING field
type LegLocaleOfIssueField struct{ quickfix.StringValue }

//Tag returns tag.LegLocaleOfIssue (598)
func (f LegLocaleOfIssueField) Tag() quickfix.Tag { return tag.LegLocaleOfIssue }

//NewLegLocaleOfIssue returns a new LegLocaleOfIssueField initialized with val
func NewLegLocaleOfIssue(val string) *LegLocaleOfIssueField {
	field := &LegLocaleOfIssueField{}
	field.Value = val
	return field
}

//LegMaturityDateField is a LOCALMKTDATE field
type LegMaturityDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.LegMaturityDate (611)
func (f LegMaturityDateField) Tag() quickfix.Tag { return tag.LegMaturityDate }

//NewLegMaturityDate returns a new LegMaturityDateField initialized with val
func NewLegMaturityDate(val string) *LegMaturityDateField {
	field := &LegMaturityDateField{}
	field.Value = val
	return field
}

//LegMaturityMonthYearField is a MONTHYEAR field
type LegMaturityMonthYearField struct{ quickfix.MonthYearValue }

//Tag returns tag.LegMaturityMonthYear (610)
func (f LegMaturityMonthYearField) Tag() quickfix.Tag { return tag.LegMaturityMonthYear }

//NewLegMaturityMonthYear returns a new LegMaturityMonthYearField initialized with val
func NewLegMaturityMonthYear(val string) *LegMaturityMonthYearField {
	field := &LegMaturityMonthYearField{}
	field.Value = val
	return field
}

//LegMaturityTimeField is a TZTIMEONLY field
type LegMaturityTimeField struct{ quickfix.TZTimeOnlyValue }

//Tag returns tag.LegMaturityTime (1212)
func (f LegMaturityTimeField) Tag() quickfix.Tag { return tag.LegMaturityTime }

//LegNumberField is a INT field
type LegNumberField struct{ quickfix.IntValue }

//Tag returns tag.LegNumber (1152)
func (f LegNumberField) Tag() quickfix.Tag { return tag.LegNumber }

//NewLegNumber returns a new LegNumberField initialized with val
func NewLegNumber(val int) *LegNumberField {
	field := &LegNumberField{}
	field.Value = val
	return field
}

//LegOfferForwardPointsField is a PRICEOFFSET field
type LegOfferForwardPointsField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.LegOfferForwardPoints (1068)
func (f LegOfferForwardPointsField) Tag() quickfix.Tag { return tag.LegOfferForwardPoints }

//NewLegOfferForwardPoints returns a new LegOfferForwardPointsField initialized with val
func NewLegOfferForwardPoints(val float64) *LegOfferForwardPointsField {
	field := &LegOfferForwardPointsField{}
	field.Value = val
	return field
}

//LegOfferPxField is a PRICE field
type LegOfferPxField struct{ quickfix.PriceValue }

//Tag returns tag.LegOfferPx (684)
func (f LegOfferPxField) Tag() quickfix.Tag { return tag.LegOfferPx }

//NewLegOfferPx returns a new LegOfferPxField initialized with val
func NewLegOfferPx(val float64) *LegOfferPxField {
	field := &LegOfferPxField{}
	field.Value = val
	return field
}

//LegOptAttributeField is a CHAR field
type LegOptAttributeField struct{ quickfix.CharValue }

//Tag returns tag.LegOptAttribute (613)
func (f LegOptAttributeField) Tag() quickfix.Tag { return tag.LegOptAttribute }

//NewLegOptAttribute returns a new LegOptAttributeField initialized with val
func NewLegOptAttribute(val string) *LegOptAttributeField {
	field := &LegOptAttributeField{}
	field.Value = val
	return field
}

//LegOptionRatioField is a FLOAT field
type LegOptionRatioField struct{ quickfix.FloatValue }

//Tag returns tag.LegOptionRatio (1017)
func (f LegOptionRatioField) Tag() quickfix.Tag { return tag.LegOptionRatio }

//NewLegOptionRatio returns a new LegOptionRatioField initialized with val
func NewLegOptionRatio(val float64) *LegOptionRatioField {
	field := &LegOptionRatioField{}
	field.Value = val
	return field
}

//LegOrderQtyField is a QTY field
type LegOrderQtyField struct{ quickfix.QtyValue }

//Tag returns tag.LegOrderQty (685)
func (f LegOrderQtyField) Tag() quickfix.Tag { return tag.LegOrderQty }

//NewLegOrderQty returns a new LegOrderQtyField initialized with val
func NewLegOrderQty(val float64) *LegOrderQtyField {
	field := &LegOrderQtyField{}
	field.Value = val
	return field
}

//LegPoolField is a STRING field
type LegPoolField struct{ quickfix.StringValue }

//Tag returns tag.LegPool (740)
func (f LegPoolField) Tag() quickfix.Tag { return tag.LegPool }

//NewLegPool returns a new LegPoolField initialized with val
func NewLegPool(val string) *LegPoolField {
	field := &LegPoolField{}
	field.Value = val
	return field
}

//LegPositionEffectField is a CHAR field
type LegPositionEffectField struct{ quickfix.CharValue }

//Tag returns tag.LegPositionEffect (564)
func (f LegPositionEffectField) Tag() quickfix.Tag { return tag.LegPositionEffect }

//NewLegPositionEffect returns a new LegPositionEffectField initialized with val
func NewLegPositionEffect(val string) *LegPositionEffectField {
	field := &LegPositionEffectField{}
	field.Value = val
	return field
}

//LegPriceField is a PRICE field
type LegPriceField struct{ quickfix.PriceValue }

//Tag returns tag.LegPrice (566)
func (f LegPriceField) Tag() quickfix.Tag { return tag.LegPrice }

//NewLegPrice returns a new LegPriceField initialized with val
func NewLegPrice(val float64) *LegPriceField {
	field := &LegPriceField{}
	field.Value = val
	return field
}

//LegPriceTypeField is a INT field
type LegPriceTypeField struct{ quickfix.IntValue }

//Tag returns tag.LegPriceType (686)
func (f LegPriceTypeField) Tag() quickfix.Tag { return tag.LegPriceType }

//NewLegPriceType returns a new LegPriceTypeField initialized with val
func NewLegPriceType(val int) *LegPriceTypeField {
	field := &LegPriceTypeField{}
	field.Value = val
	return field
}

//LegPriceUnitOfMeasureField is a STRING field
type LegPriceUnitOfMeasureField struct{ quickfix.StringValue }

//Tag returns tag.LegPriceUnitOfMeasure (1421)
func (f LegPriceUnitOfMeasureField) Tag() quickfix.Tag { return tag.LegPriceUnitOfMeasure }

//NewLegPriceUnitOfMeasure returns a new LegPriceUnitOfMeasureField initialized with val
func NewLegPriceUnitOfMeasure(val string) *LegPriceUnitOfMeasureField {
	field := &LegPriceUnitOfMeasureField{}
	field.Value = val
	return field
}

//LegPriceUnitOfMeasureQtyField is a QTY field
type LegPriceUnitOfMeasureQtyField struct{ quickfix.QtyValue }

//Tag returns tag.LegPriceUnitOfMeasureQty (1422)
func (f LegPriceUnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.LegPriceUnitOfMeasureQty }

//NewLegPriceUnitOfMeasureQty returns a new LegPriceUnitOfMeasureQtyField initialized with val
func NewLegPriceUnitOfMeasureQty(val float64) *LegPriceUnitOfMeasureQtyField {
	field := &LegPriceUnitOfMeasureQtyField{}
	field.Value = val
	return field
}

//LegProductField is a INT field
type LegProductField struct{ quickfix.IntValue }

//Tag returns tag.LegProduct (607)
func (f LegProductField) Tag() quickfix.Tag { return tag.LegProduct }

//NewLegProduct returns a new LegProductField initialized with val
func NewLegProduct(val int) *LegProductField {
	field := &LegProductField{}
	field.Value = val
	return field
}

//LegPutOrCallField is a INT field
type LegPutOrCallField struct{ quickfix.IntValue }

//Tag returns tag.LegPutOrCall (1358)
func (f LegPutOrCallField) Tag() quickfix.Tag { return tag.LegPutOrCall }

//NewLegPutOrCall returns a new LegPutOrCallField initialized with val
func NewLegPutOrCall(val int) *LegPutOrCallField {
	field := &LegPutOrCallField{}
	field.Value = val
	return field
}

//LegQtyField is a QTY field
type LegQtyField struct{ quickfix.QtyValue }

//Tag returns tag.LegQty (687)
func (f LegQtyField) Tag() quickfix.Tag { return tag.LegQty }

//NewLegQty returns a new LegQtyField initialized with val
func NewLegQty(val float64) *LegQtyField {
	field := &LegQtyField{}
	field.Value = val
	return field
}

//LegRatioQtyField is a FLOAT field
type LegRatioQtyField struct{ quickfix.FloatValue }

//Tag returns tag.LegRatioQty (623)
func (f LegRatioQtyField) Tag() quickfix.Tag { return tag.LegRatioQty }

//NewLegRatioQty returns a new LegRatioQtyField initialized with val
func NewLegRatioQty(val float64) *LegRatioQtyField {
	field := &LegRatioQtyField{}
	field.Value = val
	return field
}

//LegRedemptionDateField is a LOCALMKTDATE field
type LegRedemptionDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.LegRedemptionDate (254)
func (f LegRedemptionDateField) Tag() quickfix.Tag { return tag.LegRedemptionDate }

//NewLegRedemptionDate returns a new LegRedemptionDateField initialized with val
func NewLegRedemptionDate(val string) *LegRedemptionDateField {
	field := &LegRedemptionDateField{}
	field.Value = val
	return field
}

//LegRefIDField is a STRING field
type LegRefIDField struct{ quickfix.StringValue }

//Tag returns tag.LegRefID (654)
func (f LegRefIDField) Tag() quickfix.Tag { return tag.LegRefID }

//NewLegRefID returns a new LegRefIDField initialized with val
func NewLegRefID(val string) *LegRefIDField {
	field := &LegRefIDField{}
	field.Value = val
	return field
}

//LegRepoCollateralSecurityTypeField is a INT field
type LegRepoCollateralSecurityTypeField struct{ quickfix.IntValue }

//Tag returns tag.LegRepoCollateralSecurityType (250)
func (f LegRepoCollateralSecurityTypeField) Tag() quickfix.Tag {
	return tag.LegRepoCollateralSecurityType
}

//NewLegRepoCollateralSecurityType returns a new LegRepoCollateralSecurityTypeField initialized with val
func NewLegRepoCollateralSecurityType(val int) *LegRepoCollateralSecurityTypeField {
	field := &LegRepoCollateralSecurityTypeField{}
	field.Value = val
	return field
}

//LegReportIDField is a STRING field
type LegReportIDField struct{ quickfix.StringValue }

//Tag returns tag.LegReportID (990)
func (f LegReportIDField) Tag() quickfix.Tag { return tag.LegReportID }

//NewLegReportID returns a new LegReportIDField initialized with val
func NewLegReportID(val string) *LegReportIDField {
	field := &LegReportIDField{}
	field.Value = val
	return field
}

//LegRepurchaseRateField is a PERCENTAGE field
type LegRepurchaseRateField struct{ quickfix.PercentageValue }

//Tag returns tag.LegRepurchaseRate (252)
func (f LegRepurchaseRateField) Tag() quickfix.Tag { return tag.LegRepurchaseRate }

//NewLegRepurchaseRate returns a new LegRepurchaseRateField initialized with val
func NewLegRepurchaseRate(val float64) *LegRepurchaseRateField {
	field := &LegRepurchaseRateField{}
	field.Value = val
	return field
}

//LegRepurchaseTermField is a INT field
type LegRepurchaseTermField struct{ quickfix.IntValue }

//Tag returns tag.LegRepurchaseTerm (251)
func (f LegRepurchaseTermField) Tag() quickfix.Tag { return tag.LegRepurchaseTerm }

//NewLegRepurchaseTerm returns a new LegRepurchaseTermField initialized with val
func NewLegRepurchaseTerm(val int) *LegRepurchaseTermField {
	field := &LegRepurchaseTermField{}
	field.Value = val
	return field
}

//LegSecurityAltIDField is a STRING field
type LegSecurityAltIDField struct{ quickfix.StringValue }

//Tag returns tag.LegSecurityAltID (605)
func (f LegSecurityAltIDField) Tag() quickfix.Tag { return tag.LegSecurityAltID }

//NewLegSecurityAltID returns a new LegSecurityAltIDField initialized with val
func NewLegSecurityAltID(val string) *LegSecurityAltIDField {
	field := &LegSecurityAltIDField{}
	field.Value = val
	return field
}

//LegSecurityAltIDSourceField is a STRING field
type LegSecurityAltIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.LegSecurityAltIDSource (606)
func (f LegSecurityAltIDSourceField) Tag() quickfix.Tag { return tag.LegSecurityAltIDSource }

//NewLegSecurityAltIDSource returns a new LegSecurityAltIDSourceField initialized with val
func NewLegSecurityAltIDSource(val string) *LegSecurityAltIDSourceField {
	field := &LegSecurityAltIDSourceField{}
	field.Value = val
	return field
}

//LegSecurityDescField is a STRING field
type LegSecurityDescField struct{ quickfix.StringValue }

//Tag returns tag.LegSecurityDesc (620)
func (f LegSecurityDescField) Tag() quickfix.Tag { return tag.LegSecurityDesc }

//NewLegSecurityDesc returns a new LegSecurityDescField initialized with val
func NewLegSecurityDesc(val string) *LegSecurityDescField {
	field := &LegSecurityDescField{}
	field.Value = val
	return field
}

//LegSecurityExchangeField is a EXCHANGE field
type LegSecurityExchangeField struct{ quickfix.ExchangeValue }

//Tag returns tag.LegSecurityExchange (616)
func (f LegSecurityExchangeField) Tag() quickfix.Tag { return tag.LegSecurityExchange }

//NewLegSecurityExchange returns a new LegSecurityExchangeField initialized with val
func NewLegSecurityExchange(val string) *LegSecurityExchangeField {
	field := &LegSecurityExchangeField{}
	field.Value = val
	return field
}

//LegSecurityIDField is a STRING field
type LegSecurityIDField struct{ quickfix.StringValue }

//Tag returns tag.LegSecurityID (602)
func (f LegSecurityIDField) Tag() quickfix.Tag { return tag.LegSecurityID }

//NewLegSecurityID returns a new LegSecurityIDField initialized with val
func NewLegSecurityID(val string) *LegSecurityIDField {
	field := &LegSecurityIDField{}
	field.Value = val
	return field
}

//LegSecurityIDSourceField is a STRING field
type LegSecurityIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.LegSecurityIDSource (603)
func (f LegSecurityIDSourceField) Tag() quickfix.Tag { return tag.LegSecurityIDSource }

//NewLegSecurityIDSource returns a new LegSecurityIDSourceField initialized with val
func NewLegSecurityIDSource(val string) *LegSecurityIDSourceField {
	field := &LegSecurityIDSourceField{}
	field.Value = val
	return field
}

//LegSecuritySubTypeField is a STRING field
type LegSecuritySubTypeField struct{ quickfix.StringValue }

//Tag returns tag.LegSecuritySubType (764)
func (f LegSecuritySubTypeField) Tag() quickfix.Tag { return tag.LegSecuritySubType }

//NewLegSecuritySubType returns a new LegSecuritySubTypeField initialized with val
func NewLegSecuritySubType(val string) *LegSecuritySubTypeField {
	field := &LegSecuritySubTypeField{}
	field.Value = val
	return field
}

//LegSecurityTypeField is a STRING field
type LegSecurityTypeField struct{ quickfix.StringValue }

//Tag returns tag.LegSecurityType (609)
func (f LegSecurityTypeField) Tag() quickfix.Tag { return tag.LegSecurityType }

//NewLegSecurityType returns a new LegSecurityTypeField initialized with val
func NewLegSecurityType(val string) *LegSecurityTypeField {
	field := &LegSecurityTypeField{}
	field.Value = val
	return field
}

//LegSettlCurrencyField is a CURRENCY field
type LegSettlCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.LegSettlCurrency (675)
func (f LegSettlCurrencyField) Tag() quickfix.Tag { return tag.LegSettlCurrency }

//NewLegSettlCurrency returns a new LegSettlCurrencyField initialized with val
func NewLegSettlCurrency(val string) *LegSettlCurrencyField {
	field := &LegSettlCurrencyField{}
	field.Value = val
	return field
}

//LegSettlDateField is a LOCALMKTDATE field
type LegSettlDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.LegSettlDate (588)
func (f LegSettlDateField) Tag() quickfix.Tag { return tag.LegSettlDate }

//NewLegSettlDate returns a new LegSettlDateField initialized with val
func NewLegSettlDate(val string) *LegSettlDateField {
	field := &LegSettlDateField{}
	field.Value = val
	return field
}

//LegSettlTypeField is a CHAR field
type LegSettlTypeField struct{ quickfix.CharValue }

//Tag returns tag.LegSettlType (587)
func (f LegSettlTypeField) Tag() quickfix.Tag { return tag.LegSettlType }

//NewLegSettlType returns a new LegSettlTypeField initialized with val
func NewLegSettlType(val string) *LegSettlTypeField {
	field := &LegSettlTypeField{}
	field.Value = val
	return field
}

//LegSettlmntTypField is a CHAR field
type LegSettlmntTypField struct{ quickfix.CharValue }

//Tag returns tag.LegSettlmntTyp (587)
func (f LegSettlmntTypField) Tag() quickfix.Tag { return tag.LegSettlmntTyp }

//NewLegSettlmntTyp returns a new LegSettlmntTypField initialized with val
func NewLegSettlmntTyp(val string) *LegSettlmntTypField {
	field := &LegSettlmntTypField{}
	field.Value = val
	return field
}

//LegSideField is a CHAR field
type LegSideField struct{ quickfix.CharValue }

//Tag returns tag.LegSide (624)
func (f LegSideField) Tag() quickfix.Tag { return tag.LegSide }

//NewLegSide returns a new LegSideField initialized with val
func NewLegSide(val string) *LegSideField {
	field := &LegSideField{}
	field.Value = val
	return field
}

//LegStateOrProvinceOfIssueField is a STRING field
type LegStateOrProvinceOfIssueField struct{ quickfix.StringValue }

//Tag returns tag.LegStateOrProvinceOfIssue (597)
func (f LegStateOrProvinceOfIssueField) Tag() quickfix.Tag { return tag.LegStateOrProvinceOfIssue }

//NewLegStateOrProvinceOfIssue returns a new LegStateOrProvinceOfIssueField initialized with val
func NewLegStateOrProvinceOfIssue(val string) *LegStateOrProvinceOfIssueField {
	field := &LegStateOrProvinceOfIssueField{}
	field.Value = val
	return field
}

//LegStipulationTypeField is a STRING field
type LegStipulationTypeField struct{ quickfix.StringValue }

//Tag returns tag.LegStipulationType (688)
func (f LegStipulationTypeField) Tag() quickfix.Tag { return tag.LegStipulationType }

//NewLegStipulationType returns a new LegStipulationTypeField initialized with val
func NewLegStipulationType(val string) *LegStipulationTypeField {
	field := &LegStipulationTypeField{}
	field.Value = val
	return field
}

//LegStipulationValueField is a STRING field
type LegStipulationValueField struct{ quickfix.StringValue }

//Tag returns tag.LegStipulationValue (689)
func (f LegStipulationValueField) Tag() quickfix.Tag { return tag.LegStipulationValue }

//NewLegStipulationValue returns a new LegStipulationValueField initialized with val
func NewLegStipulationValue(val string) *LegStipulationValueField {
	field := &LegStipulationValueField{}
	field.Value = val
	return field
}

//LegStrikeCurrencyField is a CURRENCY field
type LegStrikeCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.LegStrikeCurrency (942)
func (f LegStrikeCurrencyField) Tag() quickfix.Tag { return tag.LegStrikeCurrency }

//NewLegStrikeCurrency returns a new LegStrikeCurrencyField initialized with val
func NewLegStrikeCurrency(val string) *LegStrikeCurrencyField {
	field := &LegStrikeCurrencyField{}
	field.Value = val
	return field
}

//LegStrikePriceField is a PRICE field
type LegStrikePriceField struct{ quickfix.PriceValue }

//Tag returns tag.LegStrikePrice (612)
func (f LegStrikePriceField) Tag() quickfix.Tag { return tag.LegStrikePrice }

//NewLegStrikePrice returns a new LegStrikePriceField initialized with val
func NewLegStrikePrice(val float64) *LegStrikePriceField {
	field := &LegStrikePriceField{}
	field.Value = val
	return field
}

//LegSwapTypeField is a INT field
type LegSwapTypeField struct{ quickfix.IntValue }

//Tag returns tag.LegSwapType (690)
func (f LegSwapTypeField) Tag() quickfix.Tag { return tag.LegSwapType }

//NewLegSwapType returns a new LegSwapTypeField initialized with val
func NewLegSwapType(val int) *LegSwapTypeField {
	field := &LegSwapTypeField{}
	field.Value = val
	return field
}

//LegSymbolField is a STRING field
type LegSymbolField struct{ quickfix.StringValue }

//Tag returns tag.LegSymbol (600)
func (f LegSymbolField) Tag() quickfix.Tag { return tag.LegSymbol }

//NewLegSymbol returns a new LegSymbolField initialized with val
func NewLegSymbol(val string) *LegSymbolField {
	field := &LegSymbolField{}
	field.Value = val
	return field
}

//LegSymbolSfxField is a STRING field
type LegSymbolSfxField struct{ quickfix.StringValue }

//Tag returns tag.LegSymbolSfx (601)
func (f LegSymbolSfxField) Tag() quickfix.Tag { return tag.LegSymbolSfx }

//NewLegSymbolSfx returns a new LegSymbolSfxField initialized with val
func NewLegSymbolSfx(val string) *LegSymbolSfxField {
	field := &LegSymbolSfxField{}
	field.Value = val
	return field
}

//LegTimeUnitField is a STRING field
type LegTimeUnitField struct{ quickfix.StringValue }

//Tag returns tag.LegTimeUnit (1001)
func (f LegTimeUnitField) Tag() quickfix.Tag { return tag.LegTimeUnit }

//NewLegTimeUnit returns a new LegTimeUnitField initialized with val
func NewLegTimeUnit(val string) *LegTimeUnitField {
	field := &LegTimeUnitField{}
	field.Value = val
	return field
}

//LegUnitOfMeasureField is a STRING field
type LegUnitOfMeasureField struct{ quickfix.StringValue }

//Tag returns tag.LegUnitOfMeasure (999)
func (f LegUnitOfMeasureField) Tag() quickfix.Tag { return tag.LegUnitOfMeasure }

//NewLegUnitOfMeasure returns a new LegUnitOfMeasureField initialized with val
func NewLegUnitOfMeasure(val string) *LegUnitOfMeasureField {
	field := &LegUnitOfMeasureField{}
	field.Value = val
	return field
}

//LegUnitOfMeasureQtyField is a QTY field
type LegUnitOfMeasureQtyField struct{ quickfix.QtyValue }

//Tag returns tag.LegUnitOfMeasureQty (1224)
func (f LegUnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.LegUnitOfMeasureQty }

//NewLegUnitOfMeasureQty returns a new LegUnitOfMeasureQtyField initialized with val
func NewLegUnitOfMeasureQty(val float64) *LegUnitOfMeasureQtyField {
	field := &LegUnitOfMeasureQtyField{}
	field.Value = val
	return field
}

//LegVolatilityField is a FLOAT field
type LegVolatilityField struct{ quickfix.FloatValue }

//Tag returns tag.LegVolatility (1379)
func (f LegVolatilityField) Tag() quickfix.Tag { return tag.LegVolatility }

//NewLegVolatility returns a new LegVolatilityField initialized with val
func NewLegVolatility(val float64) *LegVolatilityField {
	field := &LegVolatilityField{}
	field.Value = val
	return field
}

//LegalConfirmField is a BOOLEAN field
type LegalConfirmField struct{ quickfix.BooleanValue }

//Tag returns tag.LegalConfirm (650)
func (f LegalConfirmField) Tag() quickfix.Tag { return tag.LegalConfirm }

//NewLegalConfirm returns a new LegalConfirmField initialized with val
func NewLegalConfirm(val bool) *LegalConfirmField {
	field := &LegalConfirmField{}
	field.Value = val
	return field
}

//LinesOfTextField is a NUMINGROUP field
type LinesOfTextField struct{ quickfix.NumInGroupValue }

//Tag returns tag.LinesOfText (33)
func (f LinesOfTextField) Tag() quickfix.Tag { return tag.LinesOfText }

//NewLinesOfText returns a new LinesOfTextField initialized with val
func NewLinesOfText(val int) *LinesOfTextField {
	field := &LinesOfTextField{}
	field.Value = val
	return field
}

//LiquidityIndTypeField is a INT field
type LiquidityIndTypeField struct{ quickfix.IntValue }

//Tag returns tag.LiquidityIndType (409)
func (f LiquidityIndTypeField) Tag() quickfix.Tag { return tag.LiquidityIndType }

//NewLiquidityIndType returns a new LiquidityIndTypeField initialized with val
func NewLiquidityIndType(val int) *LiquidityIndTypeField {
	field := &LiquidityIndTypeField{}
	field.Value = val
	return field
}

//LiquidityNumSecuritiesField is a INT field
type LiquidityNumSecuritiesField struct{ quickfix.IntValue }

//Tag returns tag.LiquidityNumSecurities (441)
func (f LiquidityNumSecuritiesField) Tag() quickfix.Tag { return tag.LiquidityNumSecurities }

//NewLiquidityNumSecurities returns a new LiquidityNumSecuritiesField initialized with val
func NewLiquidityNumSecurities(val int) *LiquidityNumSecuritiesField {
	field := &LiquidityNumSecuritiesField{}
	field.Value = val
	return field
}

//LiquidityPctHighField is a PERCENTAGE field
type LiquidityPctHighField struct{ quickfix.PercentageValue }

//Tag returns tag.LiquidityPctHigh (403)
func (f LiquidityPctHighField) Tag() quickfix.Tag { return tag.LiquidityPctHigh }

//NewLiquidityPctHigh returns a new LiquidityPctHighField initialized with val
func NewLiquidityPctHigh(val float64) *LiquidityPctHighField {
	field := &LiquidityPctHighField{}
	field.Value = val
	return field
}

//LiquidityPctLowField is a PERCENTAGE field
type LiquidityPctLowField struct{ quickfix.PercentageValue }

//Tag returns tag.LiquidityPctLow (402)
func (f LiquidityPctLowField) Tag() quickfix.Tag { return tag.LiquidityPctLow }

//NewLiquidityPctLow returns a new LiquidityPctLowField initialized with val
func NewLiquidityPctLow(val float64) *LiquidityPctLowField {
	field := &LiquidityPctLowField{}
	field.Value = val
	return field
}

//LiquidityValueField is a AMT field
type LiquidityValueField struct{ quickfix.AmtValue }

//Tag returns tag.LiquidityValue (404)
func (f LiquidityValueField) Tag() quickfix.Tag { return tag.LiquidityValue }

//NewLiquidityValue returns a new LiquidityValueField initialized with val
func NewLiquidityValue(val float64) *LiquidityValueField {
	field := &LiquidityValueField{}
	field.Value = val
	return field
}

//ListExecInstField is a STRING field
type ListExecInstField struct{ quickfix.StringValue }

//Tag returns tag.ListExecInst (69)
func (f ListExecInstField) Tag() quickfix.Tag { return tag.ListExecInst }

//NewListExecInst returns a new ListExecInstField initialized with val
func NewListExecInst(val string) *ListExecInstField {
	field := &ListExecInstField{}
	field.Value = val
	return field
}

//ListExecInstTypeField is a CHAR field
type ListExecInstTypeField struct{ quickfix.CharValue }

//Tag returns tag.ListExecInstType (433)
func (f ListExecInstTypeField) Tag() quickfix.Tag { return tag.ListExecInstType }

//NewListExecInstType returns a new ListExecInstTypeField initialized with val
func NewListExecInstType(val string) *ListExecInstTypeField {
	field := &ListExecInstTypeField{}
	field.Value = val
	return field
}

//ListIDField is a STRING field
type ListIDField struct{ quickfix.StringValue }

//Tag returns tag.ListID (66)
func (f ListIDField) Tag() quickfix.Tag { return tag.ListID }

//NewListID returns a new ListIDField initialized with val
func NewListID(val string) *ListIDField {
	field := &ListIDField{}
	field.Value = val
	return field
}

//ListMethodField is a INT field
type ListMethodField struct{ quickfix.IntValue }

//Tag returns tag.ListMethod (1198)
func (f ListMethodField) Tag() quickfix.Tag { return tag.ListMethod }

//NewListMethod returns a new ListMethodField initialized with val
func NewListMethod(val int) *ListMethodField {
	field := &ListMethodField{}
	field.Value = val
	return field
}

//ListNameField is a STRING field
type ListNameField struct{ quickfix.StringValue }

//Tag returns tag.ListName (392)
func (f ListNameField) Tag() quickfix.Tag { return tag.ListName }

//NewListName returns a new ListNameField initialized with val
func NewListName(val string) *ListNameField {
	field := &ListNameField{}
	field.Value = val
	return field
}

//ListNoOrdsField is a INT field
type ListNoOrdsField struct{ quickfix.IntValue }

//Tag returns tag.ListNoOrds (68)
func (f ListNoOrdsField) Tag() quickfix.Tag { return tag.ListNoOrds }

//NewListNoOrds returns a new ListNoOrdsField initialized with val
func NewListNoOrds(val int) *ListNoOrdsField {
	field := &ListNoOrdsField{}
	field.Value = val
	return field
}

//ListOrderStatusField is a INT field
type ListOrderStatusField struct{ quickfix.IntValue }

//Tag returns tag.ListOrderStatus (431)
func (f ListOrderStatusField) Tag() quickfix.Tag { return tag.ListOrderStatus }

//NewListOrderStatus returns a new ListOrderStatusField initialized with val
func NewListOrderStatus(val int) *ListOrderStatusField {
	field := &ListOrderStatusField{}
	field.Value = val
	return field
}

//ListRejectReasonField is a INT field
type ListRejectReasonField struct{ quickfix.IntValue }

//Tag returns tag.ListRejectReason (1386)
func (f ListRejectReasonField) Tag() quickfix.Tag { return tag.ListRejectReason }

//NewListRejectReason returns a new ListRejectReasonField initialized with val
func NewListRejectReason(val int) *ListRejectReasonField {
	field := &ListRejectReasonField{}
	field.Value = val
	return field
}

//ListSeqNoField is a INT field
type ListSeqNoField struct{ quickfix.IntValue }

//Tag returns tag.ListSeqNo (67)
func (f ListSeqNoField) Tag() quickfix.Tag { return tag.ListSeqNo }

//NewListSeqNo returns a new ListSeqNoField initialized with val
func NewListSeqNo(val int) *ListSeqNoField {
	field := &ListSeqNoField{}
	field.Value = val
	return field
}

//ListStatusTextField is a STRING field
type ListStatusTextField struct{ quickfix.StringValue }

//Tag returns tag.ListStatusText (444)
func (f ListStatusTextField) Tag() quickfix.Tag { return tag.ListStatusText }

//NewListStatusText returns a new ListStatusTextField initialized with val
func NewListStatusText(val string) *ListStatusTextField {
	field := &ListStatusTextField{}
	field.Value = val
	return field
}

//ListStatusTypeField is a INT field
type ListStatusTypeField struct{ quickfix.IntValue }

//Tag returns tag.ListStatusType (429)
func (f ListStatusTypeField) Tag() quickfix.Tag { return tag.ListStatusType }

//NewListStatusType returns a new ListStatusTypeField initialized with val
func NewListStatusType(val int) *ListStatusTypeField {
	field := &ListStatusTypeField{}
	field.Value = val
	return field
}

//ListUpdateActionField is a CHAR field
type ListUpdateActionField struct{ quickfix.CharValue }

//Tag returns tag.ListUpdateAction (1324)
func (f ListUpdateActionField) Tag() quickfix.Tag { return tag.ListUpdateAction }

//NewListUpdateAction returns a new ListUpdateActionField initialized with val
func NewListUpdateAction(val string) *ListUpdateActionField {
	field := &ListUpdateActionField{}
	field.Value = val
	return field
}

//LocaleOfIssueField is a STRING field
type LocaleOfIssueField struct{ quickfix.StringValue }

//Tag returns tag.LocaleOfIssue (472)
func (f LocaleOfIssueField) Tag() quickfix.Tag { return tag.LocaleOfIssue }

//NewLocaleOfIssue returns a new LocaleOfIssueField initialized with val
func NewLocaleOfIssue(val string) *LocaleOfIssueField {
	field := &LocaleOfIssueField{}
	field.Value = val
	return field
}

//LocateReqdField is a BOOLEAN field
type LocateReqdField struct{ quickfix.BooleanValue }

//Tag returns tag.LocateReqd (114)
func (f LocateReqdField) Tag() quickfix.Tag { return tag.LocateReqd }

//NewLocateReqd returns a new LocateReqdField initialized with val
func NewLocateReqd(val bool) *LocateReqdField {
	field := &LocateReqdField{}
	field.Value = val
	return field
}

//LocationIDField is a STRING field
type LocationIDField struct{ quickfix.StringValue }

//Tag returns tag.LocationID (283)
func (f LocationIDField) Tag() quickfix.Tag { return tag.LocationID }

//NewLocationID returns a new LocationIDField initialized with val
func NewLocationID(val string) *LocationIDField {
	field := &LocationIDField{}
	field.Value = val
	return field
}

//LongQtyField is a QTY field
type LongQtyField struct{ quickfix.QtyValue }

//Tag returns tag.LongQty (704)
func (f LongQtyField) Tag() quickfix.Tag { return tag.LongQty }

//NewLongQty returns a new LongQtyField initialized with val
func NewLongQty(val float64) *LongQtyField {
	field := &LongQtyField{}
	field.Value = val
	return field
}

//LotTypeField is a CHAR field
type LotTypeField struct{ quickfix.CharValue }

//Tag returns tag.LotType (1093)
func (f LotTypeField) Tag() quickfix.Tag { return tag.LotType }

//NewLotType returns a new LotTypeField initialized with val
func NewLotType(val string) *LotTypeField {
	field := &LotTypeField{}
	field.Value = val
	return field
}

//LowLimitPriceField is a PRICE field
type LowLimitPriceField struct{ quickfix.PriceValue }

//Tag returns tag.LowLimitPrice (1148)
func (f LowLimitPriceField) Tag() quickfix.Tag { return tag.LowLimitPrice }

//NewLowLimitPrice returns a new LowLimitPriceField initialized with val
func NewLowLimitPrice(val float64) *LowLimitPriceField {
	field := &LowLimitPriceField{}
	field.Value = val
	return field
}

//LowPxField is a PRICE field
type LowPxField struct{ quickfix.PriceValue }

//Tag returns tag.LowPx (333)
func (f LowPxField) Tag() quickfix.Tag { return tag.LowPx }

//NewLowPx returns a new LowPxField initialized with val
func NewLowPx(val float64) *LowPxField {
	field := &LowPxField{}
	field.Value = val
	return field
}

//MDBookTypeField is a INT field
type MDBookTypeField struct{ quickfix.IntValue }

//Tag returns tag.MDBookType (1021)
func (f MDBookTypeField) Tag() quickfix.Tag { return tag.MDBookType }

//NewMDBookType returns a new MDBookTypeField initialized with val
func NewMDBookType(val int) *MDBookTypeField {
	field := &MDBookTypeField{}
	field.Value = val
	return field
}

//MDEntryBuyerField is a STRING field
type MDEntryBuyerField struct{ quickfix.StringValue }

//Tag returns tag.MDEntryBuyer (288)
func (f MDEntryBuyerField) Tag() quickfix.Tag { return tag.MDEntryBuyer }

//NewMDEntryBuyer returns a new MDEntryBuyerField initialized with val
func NewMDEntryBuyer(val string) *MDEntryBuyerField {
	field := &MDEntryBuyerField{}
	field.Value = val
	return field
}

//MDEntryDateField is a UTCDATEONLY field
type MDEntryDateField struct{ quickfix.UTCDateOnlyValue }

//Tag returns tag.MDEntryDate (272)
func (f MDEntryDateField) Tag() quickfix.Tag { return tag.MDEntryDate }

//MDEntryForwardPointsField is a PRICEOFFSET field
type MDEntryForwardPointsField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.MDEntryForwardPoints (1027)
func (f MDEntryForwardPointsField) Tag() quickfix.Tag { return tag.MDEntryForwardPoints }

//NewMDEntryForwardPoints returns a new MDEntryForwardPointsField initialized with val
func NewMDEntryForwardPoints(val float64) *MDEntryForwardPointsField {
	field := &MDEntryForwardPointsField{}
	field.Value = val
	return field
}

//MDEntryIDField is a STRING field
type MDEntryIDField struct{ quickfix.StringValue }

//Tag returns tag.MDEntryID (278)
func (f MDEntryIDField) Tag() quickfix.Tag { return tag.MDEntryID }

//NewMDEntryID returns a new MDEntryIDField initialized with val
func NewMDEntryID(val string) *MDEntryIDField {
	field := &MDEntryIDField{}
	field.Value = val
	return field
}

//MDEntryOriginatorField is a STRING field
type MDEntryOriginatorField struct{ quickfix.StringValue }

//Tag returns tag.MDEntryOriginator (282)
func (f MDEntryOriginatorField) Tag() quickfix.Tag { return tag.MDEntryOriginator }

//NewMDEntryOriginator returns a new MDEntryOriginatorField initialized with val
func NewMDEntryOriginator(val string) *MDEntryOriginatorField {
	field := &MDEntryOriginatorField{}
	field.Value = val
	return field
}

//MDEntryPositionNoField is a INT field
type MDEntryPositionNoField struct{ quickfix.IntValue }

//Tag returns tag.MDEntryPositionNo (290)
func (f MDEntryPositionNoField) Tag() quickfix.Tag { return tag.MDEntryPositionNo }

//NewMDEntryPositionNo returns a new MDEntryPositionNoField initialized with val
func NewMDEntryPositionNo(val int) *MDEntryPositionNoField {
	field := &MDEntryPositionNoField{}
	field.Value = val
	return field
}

//MDEntryPxField is a PRICE field
type MDEntryPxField struct{ quickfix.PriceValue }

//Tag returns tag.MDEntryPx (270)
func (f MDEntryPxField) Tag() quickfix.Tag { return tag.MDEntryPx }

//NewMDEntryPx returns a new MDEntryPxField initialized with val
func NewMDEntryPx(val float64) *MDEntryPxField {
	field := &MDEntryPxField{}
	field.Value = val
	return field
}

//MDEntryRefIDField is a STRING field
type MDEntryRefIDField struct{ quickfix.StringValue }

//Tag returns tag.MDEntryRefID (280)
func (f MDEntryRefIDField) Tag() quickfix.Tag { return tag.MDEntryRefID }

//NewMDEntryRefID returns a new MDEntryRefIDField initialized with val
func NewMDEntryRefID(val string) *MDEntryRefIDField {
	field := &MDEntryRefIDField{}
	field.Value = val
	return field
}

//MDEntrySellerField is a STRING field
type MDEntrySellerField struct{ quickfix.StringValue }

//Tag returns tag.MDEntrySeller (289)
func (f MDEntrySellerField) Tag() quickfix.Tag { return tag.MDEntrySeller }

//NewMDEntrySeller returns a new MDEntrySellerField initialized with val
func NewMDEntrySeller(val string) *MDEntrySellerField {
	field := &MDEntrySellerField{}
	field.Value = val
	return field
}

//MDEntrySizeField is a QTY field
type MDEntrySizeField struct{ quickfix.QtyValue }

//Tag returns tag.MDEntrySize (271)
func (f MDEntrySizeField) Tag() quickfix.Tag { return tag.MDEntrySize }

//NewMDEntrySize returns a new MDEntrySizeField initialized with val
func NewMDEntrySize(val float64) *MDEntrySizeField {
	field := &MDEntrySizeField{}
	field.Value = val
	return field
}

//MDEntrySpotRateField is a FLOAT field
type MDEntrySpotRateField struct{ quickfix.FloatValue }

//Tag returns tag.MDEntrySpotRate (1026)
func (f MDEntrySpotRateField) Tag() quickfix.Tag { return tag.MDEntrySpotRate }

//NewMDEntrySpotRate returns a new MDEntrySpotRateField initialized with val
func NewMDEntrySpotRate(val float64) *MDEntrySpotRateField {
	field := &MDEntrySpotRateField{}
	field.Value = val
	return field
}

//MDEntryTimeField is a UTCTIMEONLY field
type MDEntryTimeField struct{ quickfix.UTCTimeOnlyValue }

//Tag returns tag.MDEntryTime (273)
func (f MDEntryTimeField) Tag() quickfix.Tag { return tag.MDEntryTime }

//MDEntryTypeField is a CHAR field
type MDEntryTypeField struct{ quickfix.CharValue }

//Tag returns tag.MDEntryType (269)
func (f MDEntryTypeField) Tag() quickfix.Tag { return tag.MDEntryType }

//NewMDEntryType returns a new MDEntryTypeField initialized with val
func NewMDEntryType(val string) *MDEntryTypeField {
	field := &MDEntryTypeField{}
	field.Value = val
	return field
}

//MDFeedTypeField is a STRING field
type MDFeedTypeField struct{ quickfix.StringValue }

//Tag returns tag.MDFeedType (1022)
func (f MDFeedTypeField) Tag() quickfix.Tag { return tag.MDFeedType }

//NewMDFeedType returns a new MDFeedTypeField initialized with val
func NewMDFeedType(val string) *MDFeedTypeField {
	field := &MDFeedTypeField{}
	field.Value = val
	return field
}

//MDImplicitDeleteField is a BOOLEAN field
type MDImplicitDeleteField struct{ quickfix.BooleanValue }

//Tag returns tag.MDImplicitDelete (547)
func (f MDImplicitDeleteField) Tag() quickfix.Tag { return tag.MDImplicitDelete }

//NewMDImplicitDelete returns a new MDImplicitDeleteField initialized with val
func NewMDImplicitDelete(val bool) *MDImplicitDeleteField {
	field := &MDImplicitDeleteField{}
	field.Value = val
	return field
}

//MDMktField is a EXCHANGE field
type MDMktField struct{ quickfix.ExchangeValue }

//Tag returns tag.MDMkt (275)
func (f MDMktField) Tag() quickfix.Tag { return tag.MDMkt }

//NewMDMkt returns a new MDMktField initialized with val
func NewMDMkt(val string) *MDMktField {
	field := &MDMktField{}
	field.Value = val
	return field
}

//MDOriginTypeField is a INT field
type MDOriginTypeField struct{ quickfix.IntValue }

//Tag returns tag.MDOriginType (1024)
func (f MDOriginTypeField) Tag() quickfix.Tag { return tag.MDOriginType }

//NewMDOriginType returns a new MDOriginTypeField initialized with val
func NewMDOriginType(val int) *MDOriginTypeField {
	field := &MDOriginTypeField{}
	field.Value = val
	return field
}

//MDPriceLevelField is a INT field
type MDPriceLevelField struct{ quickfix.IntValue }

//Tag returns tag.MDPriceLevel (1023)
func (f MDPriceLevelField) Tag() quickfix.Tag { return tag.MDPriceLevel }

//NewMDPriceLevel returns a new MDPriceLevelField initialized with val
func NewMDPriceLevel(val int) *MDPriceLevelField {
	field := &MDPriceLevelField{}
	field.Value = val
	return field
}

//MDQuoteTypeField is a INT field
type MDQuoteTypeField struct{ quickfix.IntValue }

//Tag returns tag.MDQuoteType (1070)
func (f MDQuoteTypeField) Tag() quickfix.Tag { return tag.MDQuoteType }

//NewMDQuoteType returns a new MDQuoteTypeField initialized with val
func NewMDQuoteType(val int) *MDQuoteTypeField {
	field := &MDQuoteTypeField{}
	field.Value = val
	return field
}

//MDReportIDField is a INT field
type MDReportIDField struct{ quickfix.IntValue }

//Tag returns tag.MDReportID (963)
func (f MDReportIDField) Tag() quickfix.Tag { return tag.MDReportID }

//NewMDReportID returns a new MDReportIDField initialized with val
func NewMDReportID(val int) *MDReportIDField {
	field := &MDReportIDField{}
	field.Value = val
	return field
}

//MDReqIDField is a STRING field
type MDReqIDField struct{ quickfix.StringValue }

//Tag returns tag.MDReqID (262)
func (f MDReqIDField) Tag() quickfix.Tag { return tag.MDReqID }

//NewMDReqID returns a new MDReqIDField initialized with val
func NewMDReqID(val string) *MDReqIDField {
	field := &MDReqIDField{}
	field.Value = val
	return field
}

//MDReqRejReasonField is a CHAR field
type MDReqRejReasonField struct{ quickfix.CharValue }

//Tag returns tag.MDReqRejReason (281)
func (f MDReqRejReasonField) Tag() quickfix.Tag { return tag.MDReqRejReason }

//NewMDReqRejReason returns a new MDReqRejReasonField initialized with val
func NewMDReqRejReason(val string) *MDReqRejReasonField {
	field := &MDReqRejReasonField{}
	field.Value = val
	return field
}

//MDSecSizeField is a QTY field
type MDSecSizeField struct{ quickfix.QtyValue }

//Tag returns tag.MDSecSize (1179)
func (f MDSecSizeField) Tag() quickfix.Tag { return tag.MDSecSize }

//NewMDSecSize returns a new MDSecSizeField initialized with val
func NewMDSecSize(val float64) *MDSecSizeField {
	field := &MDSecSizeField{}
	field.Value = val
	return field
}

//MDSecSizeTypeField is a INT field
type MDSecSizeTypeField struct{ quickfix.IntValue }

//Tag returns tag.MDSecSizeType (1178)
func (f MDSecSizeTypeField) Tag() quickfix.Tag { return tag.MDSecSizeType }

//NewMDSecSizeType returns a new MDSecSizeTypeField initialized with val
func NewMDSecSizeType(val int) *MDSecSizeTypeField {
	field := &MDSecSizeTypeField{}
	field.Value = val
	return field
}

//MDStreamIDField is a STRING field
type MDStreamIDField struct{ quickfix.StringValue }

//Tag returns tag.MDStreamID (1500)
func (f MDStreamIDField) Tag() quickfix.Tag { return tag.MDStreamID }

//NewMDStreamID returns a new MDStreamIDField initialized with val
func NewMDStreamID(val string) *MDStreamIDField {
	field := &MDStreamIDField{}
	field.Value = val
	return field
}

//MDSubBookTypeField is a INT field
type MDSubBookTypeField struct{ quickfix.IntValue }

//Tag returns tag.MDSubBookType (1173)
func (f MDSubBookTypeField) Tag() quickfix.Tag { return tag.MDSubBookType }

//NewMDSubBookType returns a new MDSubBookTypeField initialized with val
func NewMDSubBookType(val int) *MDSubBookTypeField {
	field := &MDSubBookTypeField{}
	field.Value = val
	return field
}

//MDUpdateActionField is a CHAR field
type MDUpdateActionField struct{ quickfix.CharValue }

//Tag returns tag.MDUpdateAction (279)
func (f MDUpdateActionField) Tag() quickfix.Tag { return tag.MDUpdateAction }

//NewMDUpdateAction returns a new MDUpdateActionField initialized with val
func NewMDUpdateAction(val string) *MDUpdateActionField {
	field := &MDUpdateActionField{}
	field.Value = val
	return field
}

//MDUpdateTypeField is a INT field
type MDUpdateTypeField struct{ quickfix.IntValue }

//Tag returns tag.MDUpdateType (265)
func (f MDUpdateTypeField) Tag() quickfix.Tag { return tag.MDUpdateType }

//NewMDUpdateType returns a new MDUpdateTypeField initialized with val
func NewMDUpdateType(val int) *MDUpdateTypeField {
	field := &MDUpdateTypeField{}
	field.Value = val
	return field
}

//MailingDtlsField is a STRING field
type MailingDtlsField struct{ quickfix.StringValue }

//Tag returns tag.MailingDtls (474)
func (f MailingDtlsField) Tag() quickfix.Tag { return tag.MailingDtls }

//NewMailingDtls returns a new MailingDtlsField initialized with val
func NewMailingDtls(val string) *MailingDtlsField {
	field := &MailingDtlsField{}
	field.Value = val
	return field
}

//MailingInstField is a STRING field
type MailingInstField struct{ quickfix.StringValue }

//Tag returns tag.MailingInst (482)
func (f MailingInstField) Tag() quickfix.Tag { return tag.MailingInst }

//NewMailingInst returns a new MailingInstField initialized with val
func NewMailingInst(val string) *MailingInstField {
	field := &MailingInstField{}
	field.Value = val
	return field
}

//ManualOrderIndicatorField is a BOOLEAN field
type ManualOrderIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.ManualOrderIndicator (1028)
func (f ManualOrderIndicatorField) Tag() quickfix.Tag { return tag.ManualOrderIndicator }

//NewManualOrderIndicator returns a new ManualOrderIndicatorField initialized with val
func NewManualOrderIndicator(val bool) *ManualOrderIndicatorField {
	field := &ManualOrderIndicatorField{}
	field.Value = val
	return field
}

//MarginExcessField is a AMT field
type MarginExcessField struct{ quickfix.AmtValue }

//Tag returns tag.MarginExcess (899)
func (f MarginExcessField) Tag() quickfix.Tag { return tag.MarginExcess }

//NewMarginExcess returns a new MarginExcessField initialized with val
func NewMarginExcess(val float64) *MarginExcessField {
	field := &MarginExcessField{}
	field.Value = val
	return field
}

//MarginRatioField is a PERCENTAGE field
type MarginRatioField struct{ quickfix.PercentageValue }

//Tag returns tag.MarginRatio (898)
func (f MarginRatioField) Tag() quickfix.Tag { return tag.MarginRatio }

//NewMarginRatio returns a new MarginRatioField initialized with val
func NewMarginRatio(val float64) *MarginRatioField {
	field := &MarginRatioField{}
	field.Value = val
	return field
}

//MarketDepthField is a INT field
type MarketDepthField struct{ quickfix.IntValue }

//Tag returns tag.MarketDepth (264)
func (f MarketDepthField) Tag() quickfix.Tag { return tag.MarketDepth }

//NewMarketDepth returns a new MarketDepthField initialized with val
func NewMarketDepth(val int) *MarketDepthField {
	field := &MarketDepthField{}
	field.Value = val
	return field
}

//MarketIDField is a EXCHANGE field
type MarketIDField struct{ quickfix.ExchangeValue }

//Tag returns tag.MarketID (1301)
func (f MarketIDField) Tag() quickfix.Tag { return tag.MarketID }

//NewMarketID returns a new MarketIDField initialized with val
func NewMarketID(val string) *MarketIDField {
	field := &MarketIDField{}
	field.Value = val
	return field
}

//MarketReportIDField is a STRING field
type MarketReportIDField struct{ quickfix.StringValue }

//Tag returns tag.MarketReportID (1394)
func (f MarketReportIDField) Tag() quickfix.Tag { return tag.MarketReportID }

//NewMarketReportID returns a new MarketReportIDField initialized with val
func NewMarketReportID(val string) *MarketReportIDField {
	field := &MarketReportIDField{}
	field.Value = val
	return field
}

//MarketReqIDField is a STRING field
type MarketReqIDField struct{ quickfix.StringValue }

//Tag returns tag.MarketReqID (1393)
func (f MarketReqIDField) Tag() quickfix.Tag { return tag.MarketReqID }

//NewMarketReqID returns a new MarketReqIDField initialized with val
func NewMarketReqID(val string) *MarketReqIDField {
	field := &MarketReqIDField{}
	field.Value = val
	return field
}

//MarketSegmentDescField is a STRING field
type MarketSegmentDescField struct{ quickfix.StringValue }

//Tag returns tag.MarketSegmentDesc (1396)
func (f MarketSegmentDescField) Tag() quickfix.Tag { return tag.MarketSegmentDesc }

//NewMarketSegmentDesc returns a new MarketSegmentDescField initialized with val
func NewMarketSegmentDesc(val string) *MarketSegmentDescField {
	field := &MarketSegmentDescField{}
	field.Value = val
	return field
}

//MarketSegmentIDField is a STRING field
type MarketSegmentIDField struct{ quickfix.StringValue }

//Tag returns tag.MarketSegmentID (1300)
func (f MarketSegmentIDField) Tag() quickfix.Tag { return tag.MarketSegmentID }

//NewMarketSegmentID returns a new MarketSegmentIDField initialized with val
func NewMarketSegmentID(val string) *MarketSegmentIDField {
	field := &MarketSegmentIDField{}
	field.Value = val
	return field
}

//MarketUpdateActionField is a CHAR field
type MarketUpdateActionField struct{ quickfix.CharValue }

//Tag returns tag.MarketUpdateAction (1395)
func (f MarketUpdateActionField) Tag() quickfix.Tag { return tag.MarketUpdateAction }

//NewMarketUpdateAction returns a new MarketUpdateActionField initialized with val
func NewMarketUpdateAction(val string) *MarketUpdateActionField {
	field := &MarketUpdateActionField{}
	field.Value = val
	return field
}

//MassActionRejectReasonField is a INT field
type MassActionRejectReasonField struct{ quickfix.IntValue }

//Tag returns tag.MassActionRejectReason (1376)
func (f MassActionRejectReasonField) Tag() quickfix.Tag { return tag.MassActionRejectReason }

//NewMassActionRejectReason returns a new MassActionRejectReasonField initialized with val
func NewMassActionRejectReason(val int) *MassActionRejectReasonField {
	field := &MassActionRejectReasonField{}
	field.Value = val
	return field
}

//MassActionReportIDField is a STRING field
type MassActionReportIDField struct{ quickfix.StringValue }

//Tag returns tag.MassActionReportID (1369)
func (f MassActionReportIDField) Tag() quickfix.Tag { return tag.MassActionReportID }

//NewMassActionReportID returns a new MassActionReportIDField initialized with val
func NewMassActionReportID(val string) *MassActionReportIDField {
	field := &MassActionReportIDField{}
	field.Value = val
	return field
}

//MassActionResponseField is a INT field
type MassActionResponseField struct{ quickfix.IntValue }

//Tag returns tag.MassActionResponse (1375)
func (f MassActionResponseField) Tag() quickfix.Tag { return tag.MassActionResponse }

//NewMassActionResponse returns a new MassActionResponseField initialized with val
func NewMassActionResponse(val int) *MassActionResponseField {
	field := &MassActionResponseField{}
	field.Value = val
	return field
}

//MassActionScopeField is a INT field
type MassActionScopeField struct{ quickfix.IntValue }

//Tag returns tag.MassActionScope (1374)
func (f MassActionScopeField) Tag() quickfix.Tag { return tag.MassActionScope }

//NewMassActionScope returns a new MassActionScopeField initialized with val
func NewMassActionScope(val int) *MassActionScopeField {
	field := &MassActionScopeField{}
	field.Value = val
	return field
}

//MassActionTypeField is a INT field
type MassActionTypeField struct{ quickfix.IntValue }

//Tag returns tag.MassActionType (1373)
func (f MassActionTypeField) Tag() quickfix.Tag { return tag.MassActionType }

//NewMassActionType returns a new MassActionTypeField initialized with val
func NewMassActionType(val int) *MassActionTypeField {
	field := &MassActionTypeField{}
	field.Value = val
	return field
}

//MassCancelRejectReasonField is a INT field
type MassCancelRejectReasonField struct{ quickfix.IntValue }

//Tag returns tag.MassCancelRejectReason (532)
func (f MassCancelRejectReasonField) Tag() quickfix.Tag { return tag.MassCancelRejectReason }

//NewMassCancelRejectReason returns a new MassCancelRejectReasonField initialized with val
func NewMassCancelRejectReason(val int) *MassCancelRejectReasonField {
	field := &MassCancelRejectReasonField{}
	field.Value = val
	return field
}

//MassCancelRequestTypeField is a CHAR field
type MassCancelRequestTypeField struct{ quickfix.CharValue }

//Tag returns tag.MassCancelRequestType (530)
func (f MassCancelRequestTypeField) Tag() quickfix.Tag { return tag.MassCancelRequestType }

//NewMassCancelRequestType returns a new MassCancelRequestTypeField initialized with val
func NewMassCancelRequestType(val string) *MassCancelRequestTypeField {
	field := &MassCancelRequestTypeField{}
	field.Value = val
	return field
}

//MassCancelResponseField is a CHAR field
type MassCancelResponseField struct{ quickfix.CharValue }

//Tag returns tag.MassCancelResponse (531)
func (f MassCancelResponseField) Tag() quickfix.Tag { return tag.MassCancelResponse }

//NewMassCancelResponse returns a new MassCancelResponseField initialized with val
func NewMassCancelResponse(val string) *MassCancelResponseField {
	field := &MassCancelResponseField{}
	field.Value = val
	return field
}

//MassStatusReqIDField is a STRING field
type MassStatusReqIDField struct{ quickfix.StringValue }

//Tag returns tag.MassStatusReqID (584)
func (f MassStatusReqIDField) Tag() quickfix.Tag { return tag.MassStatusReqID }

//NewMassStatusReqID returns a new MassStatusReqIDField initialized with val
func NewMassStatusReqID(val string) *MassStatusReqIDField {
	field := &MassStatusReqIDField{}
	field.Value = val
	return field
}

//MassStatusReqTypeField is a INT field
type MassStatusReqTypeField struct{ quickfix.IntValue }

//Tag returns tag.MassStatusReqType (585)
func (f MassStatusReqTypeField) Tag() quickfix.Tag { return tag.MassStatusReqType }

//NewMassStatusReqType returns a new MassStatusReqTypeField initialized with val
func NewMassStatusReqType(val int) *MassStatusReqTypeField {
	field := &MassStatusReqTypeField{}
	field.Value = val
	return field
}

//MatchAlgorithmField is a STRING field
type MatchAlgorithmField struct{ quickfix.StringValue }

//Tag returns tag.MatchAlgorithm (1142)
func (f MatchAlgorithmField) Tag() quickfix.Tag { return tag.MatchAlgorithm }

//NewMatchAlgorithm returns a new MatchAlgorithmField initialized with val
func NewMatchAlgorithm(val string) *MatchAlgorithmField {
	field := &MatchAlgorithmField{}
	field.Value = val
	return field
}

//MatchIncrementField is a QTY field
type MatchIncrementField struct{ quickfix.QtyValue }

//Tag returns tag.MatchIncrement (1089)
func (f MatchIncrementField) Tag() quickfix.Tag { return tag.MatchIncrement }

//NewMatchIncrement returns a new MatchIncrementField initialized with val
func NewMatchIncrement(val float64) *MatchIncrementField {
	field := &MatchIncrementField{}
	field.Value = val
	return field
}

//MatchStatusField is a CHAR field
type MatchStatusField struct{ quickfix.CharValue }

//Tag returns tag.MatchStatus (573)
func (f MatchStatusField) Tag() quickfix.Tag { return tag.MatchStatus }

//NewMatchStatus returns a new MatchStatusField initialized with val
func NewMatchStatus(val string) *MatchStatusField {
	field := &MatchStatusField{}
	field.Value = val
	return field
}

//MatchTypeField is a STRING field
type MatchTypeField struct{ quickfix.StringValue }

//Tag returns tag.MatchType (574)
func (f MatchTypeField) Tag() quickfix.Tag { return tag.MatchType }

//NewMatchType returns a new MatchTypeField initialized with val
func NewMatchType(val string) *MatchTypeField {
	field := &MatchTypeField{}
	field.Value = val
	return field
}

//MaturityDateField is a LOCALMKTDATE field
type MaturityDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.MaturityDate (541)
func (f MaturityDateField) Tag() quickfix.Tag { return tag.MaturityDate }

//NewMaturityDate returns a new MaturityDateField initialized with val
func NewMaturityDate(val string) *MaturityDateField {
	field := &MaturityDateField{}
	field.Value = val
	return field
}

//MaturityDayField is a DAYOFMONTH field
type MaturityDayField struct{ quickfix.DayOfMonthValue }

//Tag returns tag.MaturityDay (205)
func (f MaturityDayField) Tag() quickfix.Tag { return tag.MaturityDay }

//NewMaturityDay returns a new MaturityDayField initialized with val
func NewMaturityDay(val int) *MaturityDayField {
	field := &MaturityDayField{}
	field.Value = val
	return field
}

//MaturityMonthYearField is a MONTHYEAR field
type MaturityMonthYearField struct{ quickfix.MonthYearValue }

//Tag returns tag.MaturityMonthYear (200)
func (f MaturityMonthYearField) Tag() quickfix.Tag { return tag.MaturityMonthYear }

//NewMaturityMonthYear returns a new MaturityMonthYearField initialized with val
func NewMaturityMonthYear(val string) *MaturityMonthYearField {
	field := &MaturityMonthYearField{}
	field.Value = val
	return field
}

//MaturityMonthYearFormatField is a INT field
type MaturityMonthYearFormatField struct{ quickfix.IntValue }

//Tag returns tag.MaturityMonthYearFormat (1303)
func (f MaturityMonthYearFormatField) Tag() quickfix.Tag { return tag.MaturityMonthYearFormat }

//NewMaturityMonthYearFormat returns a new MaturityMonthYearFormatField initialized with val
func NewMaturityMonthYearFormat(val int) *MaturityMonthYearFormatField {
	field := &MaturityMonthYearFormatField{}
	field.Value = val
	return field
}

//MaturityMonthYearIncrementField is a INT field
type MaturityMonthYearIncrementField struct{ quickfix.IntValue }

//Tag returns tag.MaturityMonthYearIncrement (1229)
func (f MaturityMonthYearIncrementField) Tag() quickfix.Tag { return tag.MaturityMonthYearIncrement }

//NewMaturityMonthYearIncrement returns a new MaturityMonthYearIncrementField initialized with val
func NewMaturityMonthYearIncrement(val int) *MaturityMonthYearIncrementField {
	field := &MaturityMonthYearIncrementField{}
	field.Value = val
	return field
}

//MaturityMonthYearIncrementUnitsField is a INT field
type MaturityMonthYearIncrementUnitsField struct{ quickfix.IntValue }

//Tag returns tag.MaturityMonthYearIncrementUnits (1302)
func (f MaturityMonthYearIncrementUnitsField) Tag() quickfix.Tag {
	return tag.MaturityMonthYearIncrementUnits
}

//NewMaturityMonthYearIncrementUnits returns a new MaturityMonthYearIncrementUnitsField initialized with val
func NewMaturityMonthYearIncrementUnits(val int) *MaturityMonthYearIncrementUnitsField {
	field := &MaturityMonthYearIncrementUnitsField{}
	field.Value = val
	return field
}

//MaturityNetMoneyField is a AMT field
type MaturityNetMoneyField struct{ quickfix.AmtValue }

//Tag returns tag.MaturityNetMoney (890)
func (f MaturityNetMoneyField) Tag() quickfix.Tag { return tag.MaturityNetMoney }

//NewMaturityNetMoney returns a new MaturityNetMoneyField initialized with val
func NewMaturityNetMoney(val float64) *MaturityNetMoneyField {
	field := &MaturityNetMoneyField{}
	field.Value = val
	return field
}

//MaturityRuleIDField is a STRING field
type MaturityRuleIDField struct{ quickfix.StringValue }

//Tag returns tag.MaturityRuleID (1222)
func (f MaturityRuleIDField) Tag() quickfix.Tag { return tag.MaturityRuleID }

//NewMaturityRuleID returns a new MaturityRuleIDField initialized with val
func NewMaturityRuleID(val string) *MaturityRuleIDField {
	field := &MaturityRuleIDField{}
	field.Value = val
	return field
}

//MaturityTimeField is a TZTIMEONLY field
type MaturityTimeField struct{ quickfix.TZTimeOnlyValue }

//Tag returns tag.MaturityTime (1079)
func (f MaturityTimeField) Tag() quickfix.Tag { return tag.MaturityTime }

//MaxFloorField is a QTY field
type MaxFloorField struct{ quickfix.QtyValue }

//Tag returns tag.MaxFloor (111)
func (f MaxFloorField) Tag() quickfix.Tag { return tag.MaxFloor }

//NewMaxFloor returns a new MaxFloorField initialized with val
func NewMaxFloor(val float64) *MaxFloorField {
	field := &MaxFloorField{}
	field.Value = val
	return field
}

//MaxMessageSizeField is a LENGTH field
type MaxMessageSizeField struct{ quickfix.LengthValue }

//Tag returns tag.MaxMessageSize (383)
func (f MaxMessageSizeField) Tag() quickfix.Tag { return tag.MaxMessageSize }

//NewMaxMessageSize returns a new MaxMessageSizeField initialized with val
func NewMaxMessageSize(val int) *MaxMessageSizeField {
	field := &MaxMessageSizeField{}
	field.Value = val
	return field
}

//MaxPriceLevelsField is a INT field
type MaxPriceLevelsField struct{ quickfix.IntValue }

//Tag returns tag.MaxPriceLevels (1090)
func (f MaxPriceLevelsField) Tag() quickfix.Tag { return tag.MaxPriceLevels }

//NewMaxPriceLevels returns a new MaxPriceLevelsField initialized with val
func NewMaxPriceLevels(val int) *MaxPriceLevelsField {
	field := &MaxPriceLevelsField{}
	field.Value = val
	return field
}

//MaxPriceVariationField is a FLOAT field
type MaxPriceVariationField struct{ quickfix.FloatValue }

//Tag returns tag.MaxPriceVariation (1143)
func (f MaxPriceVariationField) Tag() quickfix.Tag { return tag.MaxPriceVariation }

//NewMaxPriceVariation returns a new MaxPriceVariationField initialized with val
func NewMaxPriceVariation(val float64) *MaxPriceVariationField {
	field := &MaxPriceVariationField{}
	field.Value = val
	return field
}

//MaxShowField is a QTY field
type MaxShowField struct{ quickfix.QtyValue }

//Tag returns tag.MaxShow (210)
func (f MaxShowField) Tag() quickfix.Tag { return tag.MaxShow }

//NewMaxShow returns a new MaxShowField initialized with val
func NewMaxShow(val float64) *MaxShowField {
	field := &MaxShowField{}
	field.Value = val
	return field
}

//MaxTradeVolField is a QTY field
type MaxTradeVolField struct{ quickfix.QtyValue }

//Tag returns tag.MaxTradeVol (1140)
func (f MaxTradeVolField) Tag() quickfix.Tag { return tag.MaxTradeVol }

//NewMaxTradeVol returns a new MaxTradeVolField initialized with val
func NewMaxTradeVol(val float64) *MaxTradeVolField {
	field := &MaxTradeVolField{}
	field.Value = val
	return field
}

//MessageEncodingField is a STRING field
type MessageEncodingField struct{ quickfix.StringValue }

//Tag returns tag.MessageEncoding (347)
func (f MessageEncodingField) Tag() quickfix.Tag { return tag.MessageEncoding }

//NewMessageEncoding returns a new MessageEncodingField initialized with val
func NewMessageEncoding(val string) *MessageEncodingField {
	field := &MessageEncodingField{}
	field.Value = val
	return field
}

//MessageEventSourceField is a STRING field
type MessageEventSourceField struct{ quickfix.StringValue }

//Tag returns tag.MessageEventSource (1011)
func (f MessageEventSourceField) Tag() quickfix.Tag { return tag.MessageEventSource }

//NewMessageEventSource returns a new MessageEventSourceField initialized with val
func NewMessageEventSource(val string) *MessageEventSourceField {
	field := &MessageEventSourceField{}
	field.Value = val
	return field
}

//MidPxField is a PRICE field
type MidPxField struct{ quickfix.PriceValue }

//Tag returns tag.MidPx (631)
func (f MidPxField) Tag() quickfix.Tag { return tag.MidPx }

//NewMidPx returns a new MidPxField initialized with val
func NewMidPx(val float64) *MidPxField {
	field := &MidPxField{}
	field.Value = val
	return field
}

//MidYieldField is a PERCENTAGE field
type MidYieldField struct{ quickfix.PercentageValue }

//Tag returns tag.MidYield (633)
func (f MidYieldField) Tag() quickfix.Tag { return tag.MidYield }

//NewMidYield returns a new MidYieldField initialized with val
func NewMidYield(val float64) *MidYieldField {
	field := &MidYieldField{}
	field.Value = val
	return field
}

//MinBidSizeField is a QTY field
type MinBidSizeField struct{ quickfix.QtyValue }

//Tag returns tag.MinBidSize (647)
func (f MinBidSizeField) Tag() quickfix.Tag { return tag.MinBidSize }

//NewMinBidSize returns a new MinBidSizeField initialized with val
func NewMinBidSize(val float64) *MinBidSizeField {
	field := &MinBidSizeField{}
	field.Value = val
	return field
}

//MinLotSizeField is a QTY field
type MinLotSizeField struct{ quickfix.QtyValue }

//Tag returns tag.MinLotSize (1231)
func (f MinLotSizeField) Tag() quickfix.Tag { return tag.MinLotSize }

//NewMinLotSize returns a new MinLotSizeField initialized with val
func NewMinLotSize(val float64) *MinLotSizeField {
	field := &MinLotSizeField{}
	field.Value = val
	return field
}

//MinOfferSizeField is a QTY field
type MinOfferSizeField struct{ quickfix.QtyValue }

//Tag returns tag.MinOfferSize (648)
func (f MinOfferSizeField) Tag() quickfix.Tag { return tag.MinOfferSize }

//NewMinOfferSize returns a new MinOfferSizeField initialized with val
func NewMinOfferSize(val float64) *MinOfferSizeField {
	field := &MinOfferSizeField{}
	field.Value = val
	return field
}

//MinPriceIncrementField is a FLOAT field
type MinPriceIncrementField struct{ quickfix.FloatValue }

//Tag returns tag.MinPriceIncrement (969)
func (f MinPriceIncrementField) Tag() quickfix.Tag { return tag.MinPriceIncrement }

//NewMinPriceIncrement returns a new MinPriceIncrementField initialized with val
func NewMinPriceIncrement(val float64) *MinPriceIncrementField {
	field := &MinPriceIncrementField{}
	field.Value = val
	return field
}

//MinPriceIncrementAmountField is a AMT field
type MinPriceIncrementAmountField struct{ quickfix.AmtValue }

//Tag returns tag.MinPriceIncrementAmount (1146)
func (f MinPriceIncrementAmountField) Tag() quickfix.Tag { return tag.MinPriceIncrementAmount }

//NewMinPriceIncrementAmount returns a new MinPriceIncrementAmountField initialized with val
func NewMinPriceIncrementAmount(val float64) *MinPriceIncrementAmountField {
	field := &MinPriceIncrementAmountField{}
	field.Value = val
	return field
}

//MinQtyField is a QTY field
type MinQtyField struct{ quickfix.QtyValue }

//Tag returns tag.MinQty (110)
func (f MinQtyField) Tag() quickfix.Tag { return tag.MinQty }

//NewMinQty returns a new MinQtyField initialized with val
func NewMinQty(val float64) *MinQtyField {
	field := &MinQtyField{}
	field.Value = val
	return field
}

//MinTradeVolField is a QTY field
type MinTradeVolField struct{ quickfix.QtyValue }

//Tag returns tag.MinTradeVol (562)
func (f MinTradeVolField) Tag() quickfix.Tag { return tag.MinTradeVol }

//NewMinTradeVol returns a new MinTradeVolField initialized with val
func NewMinTradeVol(val float64) *MinTradeVolField {
	field := &MinTradeVolField{}
	field.Value = val
	return field
}

//MiscFeeAmtField is a AMT field
type MiscFeeAmtField struct{ quickfix.AmtValue }

//Tag returns tag.MiscFeeAmt (137)
func (f MiscFeeAmtField) Tag() quickfix.Tag { return tag.MiscFeeAmt }

//NewMiscFeeAmt returns a new MiscFeeAmtField initialized with val
func NewMiscFeeAmt(val float64) *MiscFeeAmtField {
	field := &MiscFeeAmtField{}
	field.Value = val
	return field
}

//MiscFeeBasisField is a INT field
type MiscFeeBasisField struct{ quickfix.IntValue }

//Tag returns tag.MiscFeeBasis (891)
func (f MiscFeeBasisField) Tag() quickfix.Tag { return tag.MiscFeeBasis }

//NewMiscFeeBasis returns a new MiscFeeBasisField initialized with val
func NewMiscFeeBasis(val int) *MiscFeeBasisField {
	field := &MiscFeeBasisField{}
	field.Value = val
	return field
}

//MiscFeeCurrField is a CURRENCY field
type MiscFeeCurrField struct{ quickfix.CurrencyValue }

//Tag returns tag.MiscFeeCurr (138)
func (f MiscFeeCurrField) Tag() quickfix.Tag { return tag.MiscFeeCurr }

//NewMiscFeeCurr returns a new MiscFeeCurrField initialized with val
func NewMiscFeeCurr(val string) *MiscFeeCurrField {
	field := &MiscFeeCurrField{}
	field.Value = val
	return field
}

//MiscFeeTypeField is a STRING field
type MiscFeeTypeField struct{ quickfix.StringValue }

//Tag returns tag.MiscFeeType (139)
func (f MiscFeeTypeField) Tag() quickfix.Tag { return tag.MiscFeeType }

//NewMiscFeeType returns a new MiscFeeTypeField initialized with val
func NewMiscFeeType(val string) *MiscFeeTypeField {
	field := &MiscFeeTypeField{}
	field.Value = val
	return field
}

//MktBidPxField is a PRICE field
type MktBidPxField struct{ quickfix.PriceValue }

//Tag returns tag.MktBidPx (645)
func (f MktBidPxField) Tag() quickfix.Tag { return tag.MktBidPx }

//NewMktBidPx returns a new MktBidPxField initialized with val
func NewMktBidPx(val float64) *MktBidPxField {
	field := &MktBidPxField{}
	field.Value = val
	return field
}

//MktOfferPxField is a PRICE field
type MktOfferPxField struct{ quickfix.PriceValue }

//Tag returns tag.MktOfferPx (646)
func (f MktOfferPxField) Tag() quickfix.Tag { return tag.MktOfferPx }

//NewMktOfferPx returns a new MktOfferPxField initialized with val
func NewMktOfferPx(val float64) *MktOfferPxField {
	field := &MktOfferPxField{}
	field.Value = val
	return field
}

//ModelTypeField is a INT field
type ModelTypeField struct{ quickfix.IntValue }

//Tag returns tag.ModelType (1434)
func (f ModelTypeField) Tag() quickfix.Tag { return tag.ModelType }

//NewModelType returns a new ModelTypeField initialized with val
func NewModelType(val int) *ModelTypeField {
	field := &ModelTypeField{}
	field.Value = val
	return field
}

//MoneyLaunderingStatusField is a CHAR field
type MoneyLaunderingStatusField struct{ quickfix.CharValue }

//Tag returns tag.MoneyLaunderingStatus (481)
func (f MoneyLaunderingStatusField) Tag() quickfix.Tag { return tag.MoneyLaunderingStatus }

//NewMoneyLaunderingStatus returns a new MoneyLaunderingStatusField initialized with val
func NewMoneyLaunderingStatus(val string) *MoneyLaunderingStatusField {
	field := &MoneyLaunderingStatusField{}
	field.Value = val
	return field
}

//MsgDirectionField is a CHAR field
type MsgDirectionField struct{ quickfix.CharValue }

//Tag returns tag.MsgDirection (385)
func (f MsgDirectionField) Tag() quickfix.Tag { return tag.MsgDirection }

//NewMsgDirection returns a new MsgDirectionField initialized with val
func NewMsgDirection(val string) *MsgDirectionField {
	field := &MsgDirectionField{}
	field.Value = val
	return field
}

//MsgSeqNumField is a SEQNUM field
type MsgSeqNumField struct{ quickfix.SeqNumValue }

//Tag returns tag.MsgSeqNum (34)
func (f MsgSeqNumField) Tag() quickfix.Tag { return tag.MsgSeqNum }

//NewMsgSeqNum returns a new MsgSeqNumField initialized with val
func NewMsgSeqNum(val int) *MsgSeqNumField {
	field := &MsgSeqNumField{}
	field.Value = val
	return field
}

//MsgTypeField is a STRING field
type MsgTypeField struct{ quickfix.StringValue }

//Tag returns tag.MsgType (35)
func (f MsgTypeField) Tag() quickfix.Tag { return tag.MsgType }

//NewMsgType returns a new MsgTypeField initialized with val
func NewMsgType(val string) *MsgTypeField {
	field := &MsgTypeField{}
	field.Value = val
	return field
}

//MultiLegReportingTypeField is a CHAR field
type MultiLegReportingTypeField struct{ quickfix.CharValue }

//Tag returns tag.MultiLegReportingType (442)
func (f MultiLegReportingTypeField) Tag() quickfix.Tag { return tag.MultiLegReportingType }

//NewMultiLegReportingType returns a new MultiLegReportingTypeField initialized with val
func NewMultiLegReportingType(val string) *MultiLegReportingTypeField {
	field := &MultiLegReportingTypeField{}
	field.Value = val
	return field
}

//MultiLegRptTypeReqField is a INT field
type MultiLegRptTypeReqField struct{ quickfix.IntValue }

//Tag returns tag.MultiLegRptTypeReq (563)
func (f MultiLegRptTypeReqField) Tag() quickfix.Tag { return tag.MultiLegRptTypeReq }

//NewMultiLegRptTypeReq returns a new MultiLegRptTypeReqField initialized with val
func NewMultiLegRptTypeReq(val int) *MultiLegRptTypeReqField {
	field := &MultiLegRptTypeReqField{}
	field.Value = val
	return field
}

//MultilegModelField is a INT field
type MultilegModelField struct{ quickfix.IntValue }

//Tag returns tag.MultilegModel (1377)
func (f MultilegModelField) Tag() quickfix.Tag { return tag.MultilegModel }

//NewMultilegModel returns a new MultilegModelField initialized with val
func NewMultilegModel(val int) *MultilegModelField {
	field := &MultilegModelField{}
	field.Value = val
	return field
}

//MultilegPriceMethodField is a INT field
type MultilegPriceMethodField struct{ quickfix.IntValue }

//Tag returns tag.MultilegPriceMethod (1378)
func (f MultilegPriceMethodField) Tag() quickfix.Tag { return tag.MultilegPriceMethod }

//NewMultilegPriceMethod returns a new MultilegPriceMethodField initialized with val
func NewMultilegPriceMethod(val int) *MultilegPriceMethodField {
	field := &MultilegPriceMethodField{}
	field.Value = val
	return field
}

//NTPositionLimitField is a INT field
type NTPositionLimitField struct{ quickfix.IntValue }

//Tag returns tag.NTPositionLimit (971)
func (f NTPositionLimitField) Tag() quickfix.Tag { return tag.NTPositionLimit }

//NewNTPositionLimit returns a new NTPositionLimitField initialized with val
func NewNTPositionLimit(val int) *NTPositionLimitField {
	field := &NTPositionLimitField{}
	field.Value = val
	return field
}

//Nested2PartyIDField is a STRING field
type Nested2PartyIDField struct{ quickfix.StringValue }

//Tag returns tag.Nested2PartyID (757)
func (f Nested2PartyIDField) Tag() quickfix.Tag { return tag.Nested2PartyID }

//NewNested2PartyID returns a new Nested2PartyIDField initialized with val
func NewNested2PartyID(val string) *Nested2PartyIDField {
	field := &Nested2PartyIDField{}
	field.Value = val
	return field
}

//Nested2PartyIDSourceField is a CHAR field
type Nested2PartyIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.Nested2PartyIDSource (758)
func (f Nested2PartyIDSourceField) Tag() quickfix.Tag { return tag.Nested2PartyIDSource }

//NewNested2PartyIDSource returns a new Nested2PartyIDSourceField initialized with val
func NewNested2PartyIDSource(val string) *Nested2PartyIDSourceField {
	field := &Nested2PartyIDSourceField{}
	field.Value = val
	return field
}

//Nested2PartyRoleField is a INT field
type Nested2PartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.Nested2PartyRole (759)
func (f Nested2PartyRoleField) Tag() quickfix.Tag { return tag.Nested2PartyRole }

//NewNested2PartyRole returns a new Nested2PartyRoleField initialized with val
func NewNested2PartyRole(val int) *Nested2PartyRoleField {
	field := &Nested2PartyRoleField{}
	field.Value = val
	return field
}

//Nested2PartySubIDField is a STRING field
type Nested2PartySubIDField struct{ quickfix.StringValue }

//Tag returns tag.Nested2PartySubID (760)
func (f Nested2PartySubIDField) Tag() quickfix.Tag { return tag.Nested2PartySubID }

//NewNested2PartySubID returns a new Nested2PartySubIDField initialized with val
func NewNested2PartySubID(val string) *Nested2PartySubIDField {
	field := &Nested2PartySubIDField{}
	field.Value = val
	return field
}

//Nested2PartySubIDTypeField is a INT field
type Nested2PartySubIDTypeField struct{ quickfix.IntValue }

//Tag returns tag.Nested2PartySubIDType (807)
func (f Nested2PartySubIDTypeField) Tag() quickfix.Tag { return tag.Nested2PartySubIDType }

//NewNested2PartySubIDType returns a new Nested2PartySubIDTypeField initialized with val
func NewNested2PartySubIDType(val int) *Nested2PartySubIDTypeField {
	field := &Nested2PartySubIDTypeField{}
	field.Value = val
	return field
}

//Nested3PartyIDField is a STRING field
type Nested3PartyIDField struct{ quickfix.StringValue }

//Tag returns tag.Nested3PartyID (949)
func (f Nested3PartyIDField) Tag() quickfix.Tag { return tag.Nested3PartyID }

//NewNested3PartyID returns a new Nested3PartyIDField initialized with val
func NewNested3PartyID(val string) *Nested3PartyIDField {
	field := &Nested3PartyIDField{}
	field.Value = val
	return field
}

//Nested3PartyIDSourceField is a CHAR field
type Nested3PartyIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.Nested3PartyIDSource (950)
func (f Nested3PartyIDSourceField) Tag() quickfix.Tag { return tag.Nested3PartyIDSource }

//NewNested3PartyIDSource returns a new Nested3PartyIDSourceField initialized with val
func NewNested3PartyIDSource(val string) *Nested3PartyIDSourceField {
	field := &Nested3PartyIDSourceField{}
	field.Value = val
	return field
}

//Nested3PartyRoleField is a INT field
type Nested3PartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.Nested3PartyRole (951)
func (f Nested3PartyRoleField) Tag() quickfix.Tag { return tag.Nested3PartyRole }

//NewNested3PartyRole returns a new Nested3PartyRoleField initialized with val
func NewNested3PartyRole(val int) *Nested3PartyRoleField {
	field := &Nested3PartyRoleField{}
	field.Value = val
	return field
}

//Nested3PartySubIDField is a STRING field
type Nested3PartySubIDField struct{ quickfix.StringValue }

//Tag returns tag.Nested3PartySubID (953)
func (f Nested3PartySubIDField) Tag() quickfix.Tag { return tag.Nested3PartySubID }

//NewNested3PartySubID returns a new Nested3PartySubIDField initialized with val
func NewNested3PartySubID(val string) *Nested3PartySubIDField {
	field := &Nested3PartySubIDField{}
	field.Value = val
	return field
}

//Nested3PartySubIDTypeField is a INT field
type Nested3PartySubIDTypeField struct{ quickfix.IntValue }

//Tag returns tag.Nested3PartySubIDType (954)
func (f Nested3PartySubIDTypeField) Tag() quickfix.Tag { return tag.Nested3PartySubIDType }

//NewNested3PartySubIDType returns a new Nested3PartySubIDTypeField initialized with val
func NewNested3PartySubIDType(val int) *Nested3PartySubIDTypeField {
	field := &Nested3PartySubIDTypeField{}
	field.Value = val
	return field
}

//Nested4PartyIDField is a STRING field
type Nested4PartyIDField struct{ quickfix.StringValue }

//Tag returns tag.Nested4PartyID (1415)
func (f Nested4PartyIDField) Tag() quickfix.Tag { return tag.Nested4PartyID }

//NewNested4PartyID returns a new Nested4PartyIDField initialized with val
func NewNested4PartyID(val string) *Nested4PartyIDField {
	field := &Nested4PartyIDField{}
	field.Value = val
	return field
}

//Nested4PartyIDSourceField is a CHAR field
type Nested4PartyIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.Nested4PartyIDSource (1416)
func (f Nested4PartyIDSourceField) Tag() quickfix.Tag { return tag.Nested4PartyIDSource }

//NewNested4PartyIDSource returns a new Nested4PartyIDSourceField initialized with val
func NewNested4PartyIDSource(val string) *Nested4PartyIDSourceField {
	field := &Nested4PartyIDSourceField{}
	field.Value = val
	return field
}

//Nested4PartyRoleField is a INT field
type Nested4PartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.Nested4PartyRole (1417)
func (f Nested4PartyRoleField) Tag() quickfix.Tag { return tag.Nested4PartyRole }

//NewNested4PartyRole returns a new Nested4PartyRoleField initialized with val
func NewNested4PartyRole(val int) *Nested4PartyRoleField {
	field := &Nested4PartyRoleField{}
	field.Value = val
	return field
}

//Nested4PartySubIDField is a STRING field
type Nested4PartySubIDField struct{ quickfix.StringValue }

//Tag returns tag.Nested4PartySubID (1412)
func (f Nested4PartySubIDField) Tag() quickfix.Tag { return tag.Nested4PartySubID }

//NewNested4PartySubID returns a new Nested4PartySubIDField initialized with val
func NewNested4PartySubID(val string) *Nested4PartySubIDField {
	field := &Nested4PartySubIDField{}
	field.Value = val
	return field
}

//Nested4PartySubIDTypeField is a INT field
type Nested4PartySubIDTypeField struct{ quickfix.IntValue }

//Tag returns tag.Nested4PartySubIDType (1411)
func (f Nested4PartySubIDTypeField) Tag() quickfix.Tag { return tag.Nested4PartySubIDType }

//NewNested4PartySubIDType returns a new Nested4PartySubIDTypeField initialized with val
func NewNested4PartySubIDType(val int) *Nested4PartySubIDTypeField {
	field := &Nested4PartySubIDTypeField{}
	field.Value = val
	return field
}

//NestedInstrAttribTypeField is a INT field
type NestedInstrAttribTypeField struct{ quickfix.IntValue }

//Tag returns tag.NestedInstrAttribType (1210)
func (f NestedInstrAttribTypeField) Tag() quickfix.Tag { return tag.NestedInstrAttribType }

//NewNestedInstrAttribType returns a new NestedInstrAttribTypeField initialized with val
func NewNestedInstrAttribType(val int) *NestedInstrAttribTypeField {
	field := &NestedInstrAttribTypeField{}
	field.Value = val
	return field
}

//NestedInstrAttribValueField is a STRING field
type NestedInstrAttribValueField struct{ quickfix.StringValue }

//Tag returns tag.NestedInstrAttribValue (1211)
func (f NestedInstrAttribValueField) Tag() quickfix.Tag { return tag.NestedInstrAttribValue }

//NewNestedInstrAttribValue returns a new NestedInstrAttribValueField initialized with val
func NewNestedInstrAttribValue(val string) *NestedInstrAttribValueField {
	field := &NestedInstrAttribValueField{}
	field.Value = val
	return field
}

//NestedPartyIDField is a STRING field
type NestedPartyIDField struct{ quickfix.StringValue }

//Tag returns tag.NestedPartyID (524)
func (f NestedPartyIDField) Tag() quickfix.Tag { return tag.NestedPartyID }

//NewNestedPartyID returns a new NestedPartyIDField initialized with val
func NewNestedPartyID(val string) *NestedPartyIDField {
	field := &NestedPartyIDField{}
	field.Value = val
	return field
}

//NestedPartyIDSourceField is a CHAR field
type NestedPartyIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.NestedPartyIDSource (525)
func (f NestedPartyIDSourceField) Tag() quickfix.Tag { return tag.NestedPartyIDSource }

//NewNestedPartyIDSource returns a new NestedPartyIDSourceField initialized with val
func NewNestedPartyIDSource(val string) *NestedPartyIDSourceField {
	field := &NestedPartyIDSourceField{}
	field.Value = val
	return field
}

//NestedPartyRoleField is a INT field
type NestedPartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.NestedPartyRole (538)
func (f NestedPartyRoleField) Tag() quickfix.Tag { return tag.NestedPartyRole }

//NewNestedPartyRole returns a new NestedPartyRoleField initialized with val
func NewNestedPartyRole(val int) *NestedPartyRoleField {
	field := &NestedPartyRoleField{}
	field.Value = val
	return field
}

//NestedPartySubIDField is a STRING field
type NestedPartySubIDField struct{ quickfix.StringValue }

//Tag returns tag.NestedPartySubID (545)
func (f NestedPartySubIDField) Tag() quickfix.Tag { return tag.NestedPartySubID }

//NewNestedPartySubID returns a new NestedPartySubIDField initialized with val
func NewNestedPartySubID(val string) *NestedPartySubIDField {
	field := &NestedPartySubIDField{}
	field.Value = val
	return field
}

//NestedPartySubIDTypeField is a INT field
type NestedPartySubIDTypeField struct{ quickfix.IntValue }

//Tag returns tag.NestedPartySubIDType (805)
func (f NestedPartySubIDTypeField) Tag() quickfix.Tag { return tag.NestedPartySubIDType }

//NewNestedPartySubIDType returns a new NestedPartySubIDTypeField initialized with val
func NewNestedPartySubIDType(val int) *NestedPartySubIDTypeField {
	field := &NestedPartySubIDTypeField{}
	field.Value = val
	return field
}

//NetChgPrevDayField is a PRICEOFFSET field
type NetChgPrevDayField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.NetChgPrevDay (451)
func (f NetChgPrevDayField) Tag() quickfix.Tag { return tag.NetChgPrevDay }

//NewNetChgPrevDay returns a new NetChgPrevDayField initialized with val
func NewNetChgPrevDay(val float64) *NetChgPrevDayField {
	field := &NetChgPrevDayField{}
	field.Value = val
	return field
}

//NetGrossIndField is a INT field
type NetGrossIndField struct{ quickfix.IntValue }

//Tag returns tag.NetGrossInd (430)
func (f NetGrossIndField) Tag() quickfix.Tag { return tag.NetGrossInd }

//NewNetGrossInd returns a new NetGrossIndField initialized with val
func NewNetGrossInd(val int) *NetGrossIndField {
	field := &NetGrossIndField{}
	field.Value = val
	return field
}

//NetMoneyField is a AMT field
type NetMoneyField struct{ quickfix.AmtValue }

//Tag returns tag.NetMoney (118)
func (f NetMoneyField) Tag() quickfix.Tag { return tag.NetMoney }

//NewNetMoney returns a new NetMoneyField initialized with val
func NewNetMoney(val float64) *NetMoneyField {
	field := &NetMoneyField{}
	field.Value = val
	return field
}

//NetworkRequestIDField is a STRING field
type NetworkRequestIDField struct{ quickfix.StringValue }

//Tag returns tag.NetworkRequestID (933)
func (f NetworkRequestIDField) Tag() quickfix.Tag { return tag.NetworkRequestID }

//NewNetworkRequestID returns a new NetworkRequestIDField initialized with val
func NewNetworkRequestID(val string) *NetworkRequestIDField {
	field := &NetworkRequestIDField{}
	field.Value = val
	return field
}

//NetworkRequestTypeField is a INT field
type NetworkRequestTypeField struct{ quickfix.IntValue }

//Tag returns tag.NetworkRequestType (935)
func (f NetworkRequestTypeField) Tag() quickfix.Tag { return tag.NetworkRequestType }

//NewNetworkRequestType returns a new NetworkRequestTypeField initialized with val
func NewNetworkRequestType(val int) *NetworkRequestTypeField {
	field := &NetworkRequestTypeField{}
	field.Value = val
	return field
}

//NetworkResponseIDField is a STRING field
type NetworkResponseIDField struct{ quickfix.StringValue }

//Tag returns tag.NetworkResponseID (932)
func (f NetworkResponseIDField) Tag() quickfix.Tag { return tag.NetworkResponseID }

//NewNetworkResponseID returns a new NetworkResponseIDField initialized with val
func NewNetworkResponseID(val string) *NetworkResponseIDField {
	field := &NetworkResponseIDField{}
	field.Value = val
	return field
}

//NetworkStatusResponseTypeField is a INT field
type NetworkStatusResponseTypeField struct{ quickfix.IntValue }

//Tag returns tag.NetworkStatusResponseType (937)
func (f NetworkStatusResponseTypeField) Tag() quickfix.Tag { return tag.NetworkStatusResponseType }

//NewNetworkStatusResponseType returns a new NetworkStatusResponseTypeField initialized with val
func NewNetworkStatusResponseType(val int) *NetworkStatusResponseTypeField {
	field := &NetworkStatusResponseTypeField{}
	field.Value = val
	return field
}

//NewPasswordField is a STRING field
type NewPasswordField struct{ quickfix.StringValue }

//Tag returns tag.NewPassword (925)
func (f NewPasswordField) Tag() quickfix.Tag { return tag.NewPassword }

//NewNewPassword returns a new NewPasswordField initialized with val
func NewNewPassword(val string) *NewPasswordField {
	field := &NewPasswordField{}
	field.Value = val
	return field
}

//NewSeqNoField is a SEQNUM field
type NewSeqNoField struct{ quickfix.SeqNumValue }

//Tag returns tag.NewSeqNo (36)
func (f NewSeqNoField) Tag() quickfix.Tag { return tag.NewSeqNo }

//NewNewSeqNo returns a new NewSeqNoField initialized with val
func NewNewSeqNo(val int) *NewSeqNoField {
	field := &NewSeqNoField{}
	field.Value = val
	return field
}

//NewsCategoryField is a INT field
type NewsCategoryField struct{ quickfix.IntValue }

//Tag returns tag.NewsCategory (1473)
func (f NewsCategoryField) Tag() quickfix.Tag { return tag.NewsCategory }

//NewNewsCategory returns a new NewsCategoryField initialized with val
func NewNewsCategory(val int) *NewsCategoryField {
	field := &NewsCategoryField{}
	field.Value = val
	return field
}

//NewsIDField is a STRING field
type NewsIDField struct{ quickfix.StringValue }

//Tag returns tag.NewsID (1472)
func (f NewsIDField) Tag() quickfix.Tag { return tag.NewsID }

//NewNewsID returns a new NewsIDField initialized with val
func NewNewsID(val string) *NewsIDField {
	field := &NewsIDField{}
	field.Value = val
	return field
}

//NewsRefIDField is a STRING field
type NewsRefIDField struct{ quickfix.StringValue }

//Tag returns tag.NewsRefID (1476)
func (f NewsRefIDField) Tag() quickfix.Tag { return tag.NewsRefID }

//NewNewsRefID returns a new NewsRefIDField initialized with val
func NewNewsRefID(val string) *NewsRefIDField {
	field := &NewsRefIDField{}
	field.Value = val
	return field
}

//NewsRefTypeField is a INT field
type NewsRefTypeField struct{ quickfix.IntValue }

//Tag returns tag.NewsRefType (1477)
func (f NewsRefTypeField) Tag() quickfix.Tag { return tag.NewsRefType }

//NewNewsRefType returns a new NewsRefTypeField initialized with val
func NewNewsRefType(val int) *NewsRefTypeField {
	field := &NewsRefTypeField{}
	field.Value = val
	return field
}

//NextExpectedMsgSeqNumField is a SEQNUM field
type NextExpectedMsgSeqNumField struct{ quickfix.SeqNumValue }

//Tag returns tag.NextExpectedMsgSeqNum (789)
func (f NextExpectedMsgSeqNumField) Tag() quickfix.Tag { return tag.NextExpectedMsgSeqNum }

//NewNextExpectedMsgSeqNum returns a new NextExpectedMsgSeqNumField initialized with val
func NewNextExpectedMsgSeqNum(val int) *NextExpectedMsgSeqNumField {
	field := &NextExpectedMsgSeqNumField{}
	field.Value = val
	return field
}

//NoAffectedOrdersField is a NUMINGROUP field
type NoAffectedOrdersField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoAffectedOrders (534)
func (f NoAffectedOrdersField) Tag() quickfix.Tag { return tag.NoAffectedOrders }

//NewNoAffectedOrders returns a new NoAffectedOrdersField initialized with val
func NewNoAffectedOrders(val int) *NoAffectedOrdersField {
	field := &NoAffectedOrdersField{}
	field.Value = val
	return field
}

//NoAllocsField is a NUMINGROUP field
type NoAllocsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoAllocs (78)
func (f NoAllocsField) Tag() quickfix.Tag { return tag.NoAllocs }

//NewNoAllocs returns a new NoAllocsField initialized with val
func NewNoAllocs(val int) *NoAllocsField {
	field := &NoAllocsField{}
	field.Value = val
	return field
}

//NoAltMDSourceField is a NUMINGROUP field
type NoAltMDSourceField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoAltMDSource (816)
func (f NoAltMDSourceField) Tag() quickfix.Tag { return tag.NoAltMDSource }

//NewNoAltMDSource returns a new NoAltMDSourceField initialized with val
func NewNoAltMDSource(val int) *NoAltMDSourceField {
	field := &NoAltMDSourceField{}
	field.Value = val
	return field
}

//NoApplIDsField is a NUMINGROUP field
type NoApplIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoApplIDs (1351)
func (f NoApplIDsField) Tag() quickfix.Tag { return tag.NoApplIDs }

//NewNoApplIDs returns a new NoApplIDsField initialized with val
func NewNoApplIDs(val int) *NoApplIDsField {
	field := &NoApplIDsField{}
	field.Value = val
	return field
}

//NoAsgnReqsField is a NUMINGROUP field
type NoAsgnReqsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoAsgnReqs (1499)
func (f NoAsgnReqsField) Tag() quickfix.Tag { return tag.NoAsgnReqs }

//NewNoAsgnReqs returns a new NoAsgnReqsField initialized with val
func NewNoAsgnReqs(val int) *NoAsgnReqsField {
	field := &NoAsgnReqsField{}
	field.Value = val
	return field
}

//NoBidComponentsField is a NUMINGROUP field
type NoBidComponentsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoBidComponents (420)
func (f NoBidComponentsField) Tag() quickfix.Tag { return tag.NoBidComponents }

//NewNoBidComponents returns a new NoBidComponentsField initialized with val
func NewNoBidComponents(val int) *NoBidComponentsField {
	field := &NoBidComponentsField{}
	field.Value = val
	return field
}

//NoBidDescriptorsField is a NUMINGROUP field
type NoBidDescriptorsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoBidDescriptors (398)
func (f NoBidDescriptorsField) Tag() quickfix.Tag { return tag.NoBidDescriptors }

//NewNoBidDescriptors returns a new NoBidDescriptorsField initialized with val
func NewNoBidDescriptors(val int) *NoBidDescriptorsField {
	field := &NoBidDescriptorsField{}
	field.Value = val
	return field
}

//NoCapacitiesField is a NUMINGROUP field
type NoCapacitiesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoCapacities (862)
func (f NoCapacitiesField) Tag() quickfix.Tag { return tag.NoCapacities }

//NewNoCapacities returns a new NoCapacitiesField initialized with val
func NewNoCapacities(val int) *NoCapacitiesField {
	field := &NoCapacitiesField{}
	field.Value = val
	return field
}

//NoClearingInstructionsField is a NUMINGROUP field
type NoClearingInstructionsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoClearingInstructions (576)
func (f NoClearingInstructionsField) Tag() quickfix.Tag { return tag.NoClearingInstructions }

//NewNoClearingInstructions returns a new NoClearingInstructionsField initialized with val
func NewNoClearingInstructions(val int) *NoClearingInstructionsField {
	field := &NoClearingInstructionsField{}
	field.Value = val
	return field
}

//NoCollInquiryQualifierField is a NUMINGROUP field
type NoCollInquiryQualifierField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoCollInquiryQualifier (938)
func (f NoCollInquiryQualifierField) Tag() quickfix.Tag { return tag.NoCollInquiryQualifier }

//NewNoCollInquiryQualifier returns a new NoCollInquiryQualifierField initialized with val
func NewNoCollInquiryQualifier(val int) *NoCollInquiryQualifierField {
	field := &NoCollInquiryQualifierField{}
	field.Value = val
	return field
}

//NoCompIDsField is a NUMINGROUP field
type NoCompIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoCompIDs (936)
func (f NoCompIDsField) Tag() quickfix.Tag { return tag.NoCompIDs }

//NewNoCompIDs returns a new NoCompIDsField initialized with val
func NewNoCompIDs(val int) *NoCompIDsField {
	field := &NoCompIDsField{}
	field.Value = val
	return field
}

//NoComplexEventDatesField is a NUMINGROUP field
type NoComplexEventDatesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoComplexEventDates (1491)
func (f NoComplexEventDatesField) Tag() quickfix.Tag { return tag.NoComplexEventDates }

//NewNoComplexEventDates returns a new NoComplexEventDatesField initialized with val
func NewNoComplexEventDates(val int) *NoComplexEventDatesField {
	field := &NoComplexEventDatesField{}
	field.Value = val
	return field
}

//NoComplexEventTimesField is a NUMINGROUP field
type NoComplexEventTimesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoComplexEventTimes (1494)
func (f NoComplexEventTimesField) Tag() quickfix.Tag { return tag.NoComplexEventTimes }

//NewNoComplexEventTimes returns a new NoComplexEventTimesField initialized with val
func NewNoComplexEventTimes(val int) *NoComplexEventTimesField {
	field := &NoComplexEventTimesField{}
	field.Value = val
	return field
}

//NoComplexEventsField is a NUMINGROUP field
type NoComplexEventsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoComplexEvents (1483)
func (f NoComplexEventsField) Tag() quickfix.Tag { return tag.NoComplexEvents }

//NewNoComplexEvents returns a new NoComplexEventsField initialized with val
func NewNoComplexEvents(val int) *NoComplexEventsField {
	field := &NoComplexEventsField{}
	field.Value = val
	return field
}

//NoContAmtsField is a NUMINGROUP field
type NoContAmtsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoContAmts (518)
func (f NoContAmtsField) Tag() quickfix.Tag { return tag.NoContAmts }

//NewNoContAmts returns a new NoContAmtsField initialized with val
func NewNoContAmts(val int) *NoContAmtsField {
	field := &NoContAmtsField{}
	field.Value = val
	return field
}

//NoContextPartyIDsField is a NUMINGROUP field
type NoContextPartyIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoContextPartyIDs (1522)
func (f NoContextPartyIDsField) Tag() quickfix.Tag { return tag.NoContextPartyIDs }

//NewNoContextPartyIDs returns a new NoContextPartyIDsField initialized with val
func NewNoContextPartyIDs(val int) *NoContextPartyIDsField {
	field := &NoContextPartyIDsField{}
	field.Value = val
	return field
}

//NoContextPartySubIDsField is a NUMINGROUP field
type NoContextPartySubIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoContextPartySubIDs (1526)
func (f NoContextPartySubIDsField) Tag() quickfix.Tag { return tag.NoContextPartySubIDs }

//NewNoContextPartySubIDs returns a new NoContextPartySubIDsField initialized with val
func NewNoContextPartySubIDs(val int) *NoContextPartySubIDsField {
	field := &NoContextPartySubIDsField{}
	field.Value = val
	return field
}

//NoContraBrokersField is a NUMINGROUP field
type NoContraBrokersField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoContraBrokers (382)
func (f NoContraBrokersField) Tag() quickfix.Tag { return tag.NoContraBrokers }

//NewNoContraBrokers returns a new NoContraBrokersField initialized with val
func NewNoContraBrokers(val int) *NoContraBrokersField {
	field := &NoContraBrokersField{}
	field.Value = val
	return field
}

//NoDatesField is a INT field
type NoDatesField struct{ quickfix.IntValue }

//Tag returns tag.NoDates (580)
func (f NoDatesField) Tag() quickfix.Tag { return tag.NoDates }

//NewNoDates returns a new NoDatesField initialized with val
func NewNoDates(val int) *NoDatesField {
	field := &NoDatesField{}
	field.Value = val
	return field
}

//NoDerivativeEventsField is a NUMINGROUP field
type NoDerivativeEventsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoDerivativeEvents (1286)
func (f NoDerivativeEventsField) Tag() quickfix.Tag { return tag.NoDerivativeEvents }

//NewNoDerivativeEvents returns a new NoDerivativeEventsField initialized with val
func NewNoDerivativeEvents(val int) *NoDerivativeEventsField {
	field := &NoDerivativeEventsField{}
	field.Value = val
	return field
}

//NoDerivativeInstrAttribField is a NUMINGROUP field
type NoDerivativeInstrAttribField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoDerivativeInstrAttrib (1311)
func (f NoDerivativeInstrAttribField) Tag() quickfix.Tag { return tag.NoDerivativeInstrAttrib }

//NewNoDerivativeInstrAttrib returns a new NoDerivativeInstrAttribField initialized with val
func NewNoDerivativeInstrAttrib(val int) *NoDerivativeInstrAttribField {
	field := &NoDerivativeInstrAttribField{}
	field.Value = val
	return field
}

//NoDerivativeInstrumentPartiesField is a NUMINGROUP field
type NoDerivativeInstrumentPartiesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoDerivativeInstrumentParties (1292)
func (f NoDerivativeInstrumentPartiesField) Tag() quickfix.Tag {
	return tag.NoDerivativeInstrumentParties
}

//NewNoDerivativeInstrumentParties returns a new NoDerivativeInstrumentPartiesField initialized with val
func NewNoDerivativeInstrumentParties(val int) *NoDerivativeInstrumentPartiesField {
	field := &NoDerivativeInstrumentPartiesField{}
	field.Value = val
	return field
}

//NoDerivativeInstrumentPartySubIDsField is a NUMINGROUP field
type NoDerivativeInstrumentPartySubIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoDerivativeInstrumentPartySubIDs (1296)
func (f NoDerivativeInstrumentPartySubIDsField) Tag() quickfix.Tag {
	return tag.NoDerivativeInstrumentPartySubIDs
}

//NewNoDerivativeInstrumentPartySubIDs returns a new NoDerivativeInstrumentPartySubIDsField initialized with val
func NewNoDerivativeInstrumentPartySubIDs(val int) *NoDerivativeInstrumentPartySubIDsField {
	field := &NoDerivativeInstrumentPartySubIDsField{}
	field.Value = val
	return field
}

//NoDerivativeSecurityAltIDField is a NUMINGROUP field
type NoDerivativeSecurityAltIDField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoDerivativeSecurityAltID (1218)
func (f NoDerivativeSecurityAltIDField) Tag() quickfix.Tag { return tag.NoDerivativeSecurityAltID }

//NewNoDerivativeSecurityAltID returns a new NoDerivativeSecurityAltIDField initialized with val
func NewNoDerivativeSecurityAltID(val int) *NoDerivativeSecurityAltIDField {
	field := &NoDerivativeSecurityAltIDField{}
	field.Value = val
	return field
}

//NoDistribInstsField is a NUMINGROUP field
type NoDistribInstsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoDistribInsts (510)
func (f NoDistribInstsField) Tag() quickfix.Tag { return tag.NoDistribInsts }

//NewNoDistribInsts returns a new NoDistribInstsField initialized with val
func NewNoDistribInsts(val int) *NoDistribInstsField {
	field := &NoDistribInstsField{}
	field.Value = val
	return field
}

//NoDlvyInstField is a NUMINGROUP field
type NoDlvyInstField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoDlvyInst (85)
func (f NoDlvyInstField) Tag() quickfix.Tag { return tag.NoDlvyInst }

//NewNoDlvyInst returns a new NoDlvyInstField initialized with val
func NewNoDlvyInst(val int) *NoDlvyInstField {
	field := &NoDlvyInstField{}
	field.Value = val
	return field
}

//NoEventsField is a NUMINGROUP field
type NoEventsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoEvents (864)
func (f NoEventsField) Tag() quickfix.Tag { return tag.NoEvents }

//NewNoEvents returns a new NoEventsField initialized with val
func NewNoEvents(val int) *NoEventsField {
	field := &NoEventsField{}
	field.Value = val
	return field
}

//NoExecInstRulesField is a NUMINGROUP field
type NoExecInstRulesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoExecInstRules (1232)
func (f NoExecInstRulesField) Tag() quickfix.Tag { return tag.NoExecInstRules }

//NewNoExecInstRules returns a new NoExecInstRulesField initialized with val
func NewNoExecInstRules(val int) *NoExecInstRulesField {
	field := &NoExecInstRulesField{}
	field.Value = val
	return field
}

//NoExecsField is a NUMINGROUP field
type NoExecsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoExecs (124)
func (f NoExecsField) Tag() quickfix.Tag { return tag.NoExecs }

//NewNoExecs returns a new NoExecsField initialized with val
func NewNoExecs(val int) *NoExecsField {
	field := &NoExecsField{}
	field.Value = val
	return field
}

//NoExpirationField is a NUMINGROUP field
type NoExpirationField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoExpiration (981)
func (f NoExpirationField) Tag() quickfix.Tag { return tag.NoExpiration }

//NewNoExpiration returns a new NoExpirationField initialized with val
func NewNoExpiration(val int) *NoExpirationField {
	field := &NoExpirationField{}
	field.Value = val
	return field
}

//NoFillsField is a NUMINGROUP field
type NoFillsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoFills (1362)
func (f NoFillsField) Tag() quickfix.Tag { return tag.NoFills }

//NewNoFills returns a new NoFillsField initialized with val
func NewNoFills(val int) *NoFillsField {
	field := &NoFillsField{}
	field.Value = val
	return field
}

//NoHopsField is a NUMINGROUP field
type NoHopsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoHops (627)
func (f NoHopsField) Tag() quickfix.Tag { return tag.NoHops }

//NewNoHops returns a new NoHopsField initialized with val
func NewNoHops(val int) *NoHopsField {
	field := &NoHopsField{}
	field.Value = val
	return field
}

//NoIOIQualifiersField is a NUMINGROUP field
type NoIOIQualifiersField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoIOIQualifiers (199)
func (f NoIOIQualifiersField) Tag() quickfix.Tag { return tag.NoIOIQualifiers }

//NewNoIOIQualifiers returns a new NoIOIQualifiersField initialized with val
func NewNoIOIQualifiers(val int) *NoIOIQualifiersField {
	field := &NoIOIQualifiersField{}
	field.Value = val
	return field
}

//NoInstrAttribField is a NUMINGROUP field
type NoInstrAttribField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoInstrAttrib (870)
func (f NoInstrAttribField) Tag() quickfix.Tag { return tag.NoInstrAttrib }

//NewNoInstrAttrib returns a new NoInstrAttribField initialized with val
func NewNoInstrAttrib(val int) *NoInstrAttribField {
	field := &NoInstrAttribField{}
	field.Value = val
	return field
}

//NoInstrumentPartiesField is a NUMINGROUP field
type NoInstrumentPartiesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoInstrumentParties (1018)
func (f NoInstrumentPartiesField) Tag() quickfix.Tag { return tag.NoInstrumentParties }

//NewNoInstrumentParties returns a new NoInstrumentPartiesField initialized with val
func NewNoInstrumentParties(val int) *NoInstrumentPartiesField {
	field := &NoInstrumentPartiesField{}
	field.Value = val
	return field
}

//NoInstrumentPartySubIDsField is a NUMINGROUP field
type NoInstrumentPartySubIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoInstrumentPartySubIDs (1052)
func (f NoInstrumentPartySubIDsField) Tag() quickfix.Tag { return tag.NoInstrumentPartySubIDs }

//NewNoInstrumentPartySubIDs returns a new NoInstrumentPartySubIDsField initialized with val
func NewNoInstrumentPartySubIDs(val int) *NoInstrumentPartySubIDsField {
	field := &NoInstrumentPartySubIDsField{}
	field.Value = val
	return field
}

//NoLegAllocsField is a NUMINGROUP field
type NoLegAllocsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoLegAllocs (670)
func (f NoLegAllocsField) Tag() quickfix.Tag { return tag.NoLegAllocs }

//NewNoLegAllocs returns a new NoLegAllocsField initialized with val
func NewNoLegAllocs(val int) *NoLegAllocsField {
	field := &NoLegAllocsField{}
	field.Value = val
	return field
}

//NoLegSecurityAltIDField is a STRING field
type NoLegSecurityAltIDField struct{ quickfix.StringValue }

//Tag returns tag.NoLegSecurityAltID (604)
func (f NoLegSecurityAltIDField) Tag() quickfix.Tag { return tag.NoLegSecurityAltID }

//NewNoLegSecurityAltID returns a new NoLegSecurityAltIDField initialized with val
func NewNoLegSecurityAltID(val string) *NoLegSecurityAltIDField {
	field := &NoLegSecurityAltIDField{}
	field.Value = val
	return field
}

//NoLegStipulationsField is a NUMINGROUP field
type NoLegStipulationsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoLegStipulations (683)
func (f NoLegStipulationsField) Tag() quickfix.Tag { return tag.NoLegStipulations }

//NewNoLegStipulations returns a new NoLegStipulationsField initialized with val
func NewNoLegStipulations(val int) *NoLegStipulationsField {
	field := &NoLegStipulationsField{}
	field.Value = val
	return field
}

//NoLegsField is a NUMINGROUP field
type NoLegsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoLegs (555)
func (f NoLegsField) Tag() quickfix.Tag { return tag.NoLegs }

//NewNoLegs returns a new NoLegsField initialized with val
func NewNoLegs(val int) *NoLegsField {
	field := &NoLegsField{}
	field.Value = val
	return field
}

//NoLinesOfTextField is a NUMINGROUP field
type NoLinesOfTextField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoLinesOfText (33)
func (f NoLinesOfTextField) Tag() quickfix.Tag { return tag.NoLinesOfText }

//NewNoLinesOfText returns a new NoLinesOfTextField initialized with val
func NewNoLinesOfText(val int) *NoLinesOfTextField {
	field := &NoLinesOfTextField{}
	field.Value = val
	return field
}

//NoLotTypeRulesField is a NUMINGROUP field
type NoLotTypeRulesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoLotTypeRules (1234)
func (f NoLotTypeRulesField) Tag() quickfix.Tag { return tag.NoLotTypeRules }

//NewNoLotTypeRules returns a new NoLotTypeRulesField initialized with val
func NewNoLotTypeRules(val int) *NoLotTypeRulesField {
	field := &NoLotTypeRulesField{}
	field.Value = val
	return field
}

//NoMDEntriesField is a NUMINGROUP field
type NoMDEntriesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoMDEntries (268)
func (f NoMDEntriesField) Tag() quickfix.Tag { return tag.NoMDEntries }

//NewNoMDEntries returns a new NoMDEntriesField initialized with val
func NewNoMDEntries(val int) *NoMDEntriesField {
	field := &NoMDEntriesField{}
	field.Value = val
	return field
}

//NoMDEntryTypesField is a NUMINGROUP field
type NoMDEntryTypesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoMDEntryTypes (267)
func (f NoMDEntryTypesField) Tag() quickfix.Tag { return tag.NoMDEntryTypes }

//NewNoMDEntryTypes returns a new NoMDEntryTypesField initialized with val
func NewNoMDEntryTypes(val int) *NoMDEntryTypesField {
	field := &NoMDEntryTypesField{}
	field.Value = val
	return field
}

//NoMDFeedTypesField is a NUMINGROUP field
type NoMDFeedTypesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoMDFeedTypes (1141)
func (f NoMDFeedTypesField) Tag() quickfix.Tag { return tag.NoMDFeedTypes }

//NewNoMDFeedTypes returns a new NoMDFeedTypesField initialized with val
func NewNoMDFeedTypes(val int) *NoMDFeedTypesField {
	field := &NoMDFeedTypesField{}
	field.Value = val
	return field
}

//NoMarketSegmentsField is a NUMINGROUP field
type NoMarketSegmentsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoMarketSegments (1310)
func (f NoMarketSegmentsField) Tag() quickfix.Tag { return tag.NoMarketSegments }

//NewNoMarketSegments returns a new NoMarketSegmentsField initialized with val
func NewNoMarketSegments(val int) *NoMarketSegmentsField {
	field := &NoMarketSegmentsField{}
	field.Value = val
	return field
}

//NoMatchRulesField is a NUMINGROUP field
type NoMatchRulesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoMatchRules (1235)
func (f NoMatchRulesField) Tag() quickfix.Tag { return tag.NoMatchRules }

//NewNoMatchRules returns a new NoMatchRulesField initialized with val
func NewNoMatchRules(val int) *NoMatchRulesField {
	field := &NoMatchRulesField{}
	field.Value = val
	return field
}

//NoMaturityRulesField is a NUMINGROUP field
type NoMaturityRulesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoMaturityRules (1236)
func (f NoMaturityRulesField) Tag() quickfix.Tag { return tag.NoMaturityRules }

//NewNoMaturityRules returns a new NoMaturityRulesField initialized with val
func NewNoMaturityRules(val int) *NoMaturityRulesField {
	field := &NoMaturityRulesField{}
	field.Value = val
	return field
}

//NoMiscFeesField is a NUMINGROUP field
type NoMiscFeesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoMiscFees (136)
func (f NoMiscFeesField) Tag() quickfix.Tag { return tag.NoMiscFees }

//NewNoMiscFees returns a new NoMiscFeesField initialized with val
func NewNoMiscFees(val int) *NoMiscFeesField {
	field := &NoMiscFeesField{}
	field.Value = val
	return field
}

//NoMsgTypesField is a NUMINGROUP field
type NoMsgTypesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoMsgTypes (384)
func (f NoMsgTypesField) Tag() quickfix.Tag { return tag.NoMsgTypes }

//NewNoMsgTypes returns a new NoMsgTypesField initialized with val
func NewNoMsgTypes(val int) *NoMsgTypesField {
	field := &NoMsgTypesField{}
	field.Value = val
	return field
}

//NoNested2PartyIDsField is a NUMINGROUP field
type NoNested2PartyIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoNested2PartyIDs (756)
func (f NoNested2PartyIDsField) Tag() quickfix.Tag { return tag.NoNested2PartyIDs }

//NewNoNested2PartyIDs returns a new NoNested2PartyIDsField initialized with val
func NewNoNested2PartyIDs(val int) *NoNested2PartyIDsField {
	field := &NoNested2PartyIDsField{}
	field.Value = val
	return field
}

//NoNested2PartySubIDsField is a NUMINGROUP field
type NoNested2PartySubIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoNested2PartySubIDs (806)
func (f NoNested2PartySubIDsField) Tag() quickfix.Tag { return tag.NoNested2PartySubIDs }

//NewNoNested2PartySubIDs returns a new NoNested2PartySubIDsField initialized with val
func NewNoNested2PartySubIDs(val int) *NoNested2PartySubIDsField {
	field := &NoNested2PartySubIDsField{}
	field.Value = val
	return field
}

//NoNested3PartyIDsField is a NUMINGROUP field
type NoNested3PartyIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoNested3PartyIDs (948)
func (f NoNested3PartyIDsField) Tag() quickfix.Tag { return tag.NoNested3PartyIDs }

//NewNoNested3PartyIDs returns a new NoNested3PartyIDsField initialized with val
func NewNoNested3PartyIDs(val int) *NoNested3PartyIDsField {
	field := &NoNested3PartyIDsField{}
	field.Value = val
	return field
}

//NoNested3PartySubIDsField is a NUMINGROUP field
type NoNested3PartySubIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoNested3PartySubIDs (952)
func (f NoNested3PartySubIDsField) Tag() quickfix.Tag { return tag.NoNested3PartySubIDs }

//NewNoNested3PartySubIDs returns a new NoNested3PartySubIDsField initialized with val
func NewNoNested3PartySubIDs(val int) *NoNested3PartySubIDsField {
	field := &NoNested3PartySubIDsField{}
	field.Value = val
	return field
}

//NoNested4PartyIDsField is a NUMINGROUP field
type NoNested4PartyIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoNested4PartyIDs (1414)
func (f NoNested4PartyIDsField) Tag() quickfix.Tag { return tag.NoNested4PartyIDs }

//NewNoNested4PartyIDs returns a new NoNested4PartyIDsField initialized with val
func NewNoNested4PartyIDs(val int) *NoNested4PartyIDsField {
	field := &NoNested4PartyIDsField{}
	field.Value = val
	return field
}

//NoNested4PartySubIDsField is a NUMINGROUP field
type NoNested4PartySubIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoNested4PartySubIDs (1413)
func (f NoNested4PartySubIDsField) Tag() quickfix.Tag { return tag.NoNested4PartySubIDs }

//NewNoNested4PartySubIDs returns a new NoNested4PartySubIDsField initialized with val
func NewNoNested4PartySubIDs(val int) *NoNested4PartySubIDsField {
	field := &NoNested4PartySubIDsField{}
	field.Value = val
	return field
}

//NoNestedInstrAttribField is a NUMINGROUP field
type NoNestedInstrAttribField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoNestedInstrAttrib (1312)
func (f NoNestedInstrAttribField) Tag() quickfix.Tag { return tag.NoNestedInstrAttrib }

//NewNoNestedInstrAttrib returns a new NoNestedInstrAttribField initialized with val
func NewNoNestedInstrAttrib(val int) *NoNestedInstrAttribField {
	field := &NoNestedInstrAttribField{}
	field.Value = val
	return field
}

//NoNestedPartyIDsField is a NUMINGROUP field
type NoNestedPartyIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoNestedPartyIDs (539)
func (f NoNestedPartyIDsField) Tag() quickfix.Tag { return tag.NoNestedPartyIDs }

//NewNoNestedPartyIDs returns a new NoNestedPartyIDsField initialized with val
func NewNoNestedPartyIDs(val int) *NoNestedPartyIDsField {
	field := &NoNestedPartyIDsField{}
	field.Value = val
	return field
}

//NoNestedPartySubIDsField is a NUMINGROUP field
type NoNestedPartySubIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoNestedPartySubIDs (804)
func (f NoNestedPartySubIDsField) Tag() quickfix.Tag { return tag.NoNestedPartySubIDs }

//NewNoNestedPartySubIDs returns a new NoNestedPartySubIDsField initialized with val
func NewNoNestedPartySubIDs(val int) *NoNestedPartySubIDsField {
	field := &NoNestedPartySubIDsField{}
	field.Value = val
	return field
}

//NoNewsRefIDsField is a NUMINGROUP field
type NoNewsRefIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoNewsRefIDs (1475)
func (f NoNewsRefIDsField) Tag() quickfix.Tag { return tag.NoNewsRefIDs }

//NewNoNewsRefIDs returns a new NoNewsRefIDsField initialized with val
func NewNoNewsRefIDs(val int) *NoNewsRefIDsField {
	field := &NoNewsRefIDsField{}
	field.Value = val
	return field
}

//NoNotAffectedOrdersField is a NUMINGROUP field
type NoNotAffectedOrdersField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoNotAffectedOrders (1370)
func (f NoNotAffectedOrdersField) Tag() quickfix.Tag { return tag.NoNotAffectedOrders }

//NewNoNotAffectedOrders returns a new NoNotAffectedOrdersField initialized with val
func NewNoNotAffectedOrders(val int) *NoNotAffectedOrdersField {
	field := &NoNotAffectedOrdersField{}
	field.Value = val
	return field
}

//NoOfLegUnderlyingsField is a NUMINGROUP field
type NoOfLegUnderlyingsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoOfLegUnderlyings (1342)
func (f NoOfLegUnderlyingsField) Tag() quickfix.Tag { return tag.NoOfLegUnderlyings }

//NewNoOfLegUnderlyings returns a new NoOfLegUnderlyingsField initialized with val
func NewNoOfLegUnderlyings(val int) *NoOfLegUnderlyingsField {
	field := &NoOfLegUnderlyingsField{}
	field.Value = val
	return field
}

//NoOfSecSizesField is a NUMINGROUP field
type NoOfSecSizesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoOfSecSizes (1177)
func (f NoOfSecSizesField) Tag() quickfix.Tag { return tag.NoOfSecSizes }

//NewNoOfSecSizes returns a new NoOfSecSizesField initialized with val
func NewNoOfSecSizes(val int) *NoOfSecSizesField {
	field := &NoOfSecSizesField{}
	field.Value = val
	return field
}

//NoOrdTypeRulesField is a NUMINGROUP field
type NoOrdTypeRulesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoOrdTypeRules (1237)
func (f NoOrdTypeRulesField) Tag() quickfix.Tag { return tag.NoOrdTypeRules }

//NewNoOrdTypeRules returns a new NoOrdTypeRulesField initialized with val
func NewNoOrdTypeRules(val int) *NoOrdTypeRulesField {
	field := &NoOrdTypeRulesField{}
	field.Value = val
	return field
}

//NoOrdersField is a NUMINGROUP field
type NoOrdersField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoOrders (73)
func (f NoOrdersField) Tag() quickfix.Tag { return tag.NoOrders }

//NewNoOrders returns a new NoOrdersField initialized with val
func NewNoOrders(val int) *NoOrdersField {
	field := &NoOrdersField{}
	field.Value = val
	return field
}

//NoPartyAltIDsField is a NUMINGROUP field
type NoPartyAltIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoPartyAltIDs (1516)
func (f NoPartyAltIDsField) Tag() quickfix.Tag { return tag.NoPartyAltIDs }

//NewNoPartyAltIDs returns a new NoPartyAltIDsField initialized with val
func NewNoPartyAltIDs(val int) *NoPartyAltIDsField {
	field := &NoPartyAltIDsField{}
	field.Value = val
	return field
}

//NoPartyAltSubIDsField is a NUMINGROUP field
type NoPartyAltSubIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoPartyAltSubIDs (1519)
func (f NoPartyAltSubIDsField) Tag() quickfix.Tag { return tag.NoPartyAltSubIDs }

//NewNoPartyAltSubIDs returns a new NoPartyAltSubIDsField initialized with val
func NewNoPartyAltSubIDs(val int) *NoPartyAltSubIDsField {
	field := &NoPartyAltSubIDsField{}
	field.Value = val
	return field
}

//NoPartyIDsField is a NUMINGROUP field
type NoPartyIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoPartyIDs (453)
func (f NoPartyIDsField) Tag() quickfix.Tag { return tag.NoPartyIDs }

//NewNoPartyIDs returns a new NoPartyIDsField initialized with val
func NewNoPartyIDs(val int) *NoPartyIDsField {
	field := &NoPartyIDsField{}
	field.Value = val
	return field
}

//NoPartyListField is a NUMINGROUP field
type NoPartyListField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoPartyList (1513)
func (f NoPartyListField) Tag() quickfix.Tag { return tag.NoPartyList }

//NewNoPartyList returns a new NoPartyListField initialized with val
func NewNoPartyList(val int) *NoPartyListField {
	field := &NoPartyListField{}
	field.Value = val
	return field
}

//NoPartyListResponseTypesField is a NUMINGROUP field
type NoPartyListResponseTypesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoPartyListResponseTypes (1506)
func (f NoPartyListResponseTypesField) Tag() quickfix.Tag { return tag.NoPartyListResponseTypes }

//NewNoPartyListResponseTypes returns a new NoPartyListResponseTypesField initialized with val
func NewNoPartyListResponseTypes(val int) *NoPartyListResponseTypesField {
	field := &NoPartyListResponseTypesField{}
	field.Value = val
	return field
}

//NoPartyRelationshipsField is a NUMINGROUP field
type NoPartyRelationshipsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoPartyRelationships (1514)
func (f NoPartyRelationshipsField) Tag() quickfix.Tag { return tag.NoPartyRelationships }

//NewNoPartyRelationships returns a new NoPartyRelationshipsField initialized with val
func NewNoPartyRelationships(val int) *NoPartyRelationshipsField {
	field := &NoPartyRelationshipsField{}
	field.Value = val
	return field
}

//NoPartySubIDsField is a NUMINGROUP field
type NoPartySubIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoPartySubIDs (802)
func (f NoPartySubIDsField) Tag() quickfix.Tag { return tag.NoPartySubIDs }

//NewNoPartySubIDs returns a new NoPartySubIDsField initialized with val
func NewNoPartySubIDs(val int) *NoPartySubIDsField {
	field := &NoPartySubIDsField{}
	field.Value = val
	return field
}

//NoPosAmtField is a NUMINGROUP field
type NoPosAmtField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoPosAmt (753)
func (f NoPosAmtField) Tag() quickfix.Tag { return tag.NoPosAmt }

//NewNoPosAmt returns a new NoPosAmtField initialized with val
func NewNoPosAmt(val int) *NoPosAmtField {
	field := &NoPosAmtField{}
	field.Value = val
	return field
}

//NoPositionsField is a NUMINGROUP field
type NoPositionsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoPositions (702)
func (f NoPositionsField) Tag() quickfix.Tag { return tag.NoPositions }

//NewNoPositions returns a new NoPositionsField initialized with val
func NewNoPositions(val int) *NoPositionsField {
	field := &NoPositionsField{}
	field.Value = val
	return field
}

//NoQuoteEntriesField is a NUMINGROUP field
type NoQuoteEntriesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoQuoteEntries (295)
func (f NoQuoteEntriesField) Tag() quickfix.Tag { return tag.NoQuoteEntries }

//NewNoQuoteEntries returns a new NoQuoteEntriesField initialized with val
func NewNoQuoteEntries(val int) *NoQuoteEntriesField {
	field := &NoQuoteEntriesField{}
	field.Value = val
	return field
}

//NoQuoteQualifiersField is a NUMINGROUP field
type NoQuoteQualifiersField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoQuoteQualifiers (735)
func (f NoQuoteQualifiersField) Tag() quickfix.Tag { return tag.NoQuoteQualifiers }

//NewNoQuoteQualifiers returns a new NoQuoteQualifiersField initialized with val
func NewNoQuoteQualifiers(val int) *NoQuoteQualifiersField {
	field := &NoQuoteQualifiersField{}
	field.Value = val
	return field
}

//NoQuoteSetsField is a NUMINGROUP field
type NoQuoteSetsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoQuoteSets (296)
func (f NoQuoteSetsField) Tag() quickfix.Tag { return tag.NoQuoteSets }

//NewNoQuoteSets returns a new NoQuoteSetsField initialized with val
func NewNoQuoteSets(val int) *NoQuoteSetsField {
	field := &NoQuoteSetsField{}
	field.Value = val
	return field
}

//NoRateSourcesField is a NUMINGROUP field
type NoRateSourcesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRateSources (1445)
func (f NoRateSourcesField) Tag() quickfix.Tag { return tag.NoRateSources }

//NewNoRateSources returns a new NoRateSourcesField initialized with val
func NewNoRateSources(val int) *NoRateSourcesField {
	field := &NoRateSourcesField{}
	field.Value = val
	return field
}

//NoRegistDtlsField is a NUMINGROUP field
type NoRegistDtlsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRegistDtls (473)
func (f NoRegistDtlsField) Tag() quickfix.Tag { return tag.NoRegistDtls }

//NewNoRegistDtls returns a new NoRegistDtlsField initialized with val
func NewNoRegistDtls(val int) *NoRegistDtlsField {
	field := &NoRegistDtlsField{}
	field.Value = val
	return field
}

//NoRelatedContextPartyIDsField is a NUMINGROUP field
type NoRelatedContextPartyIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRelatedContextPartyIDs (1575)
func (f NoRelatedContextPartyIDsField) Tag() quickfix.Tag { return tag.NoRelatedContextPartyIDs }

//NewNoRelatedContextPartyIDs returns a new NoRelatedContextPartyIDsField initialized with val
func NewNoRelatedContextPartyIDs(val int) *NoRelatedContextPartyIDsField {
	field := &NoRelatedContextPartyIDsField{}
	field.Value = val
	return field
}

//NoRelatedContextPartySubIDsField is a NUMINGROUP field
type NoRelatedContextPartySubIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRelatedContextPartySubIDs (1579)
func (f NoRelatedContextPartySubIDsField) Tag() quickfix.Tag { return tag.NoRelatedContextPartySubIDs }

//NewNoRelatedContextPartySubIDs returns a new NoRelatedContextPartySubIDsField initialized with val
func NewNoRelatedContextPartySubIDs(val int) *NoRelatedContextPartySubIDsField {
	field := &NoRelatedContextPartySubIDsField{}
	field.Value = val
	return field
}

//NoRelatedPartyAltIDsField is a NUMINGROUP field
type NoRelatedPartyAltIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRelatedPartyAltIDs (1569)
func (f NoRelatedPartyAltIDsField) Tag() quickfix.Tag { return tag.NoRelatedPartyAltIDs }

//NewNoRelatedPartyAltIDs returns a new NoRelatedPartyAltIDsField initialized with val
func NewNoRelatedPartyAltIDs(val int) *NoRelatedPartyAltIDsField {
	field := &NoRelatedPartyAltIDsField{}
	field.Value = val
	return field
}

//NoRelatedPartyAltSubIDsField is a NUMINGROUP field
type NoRelatedPartyAltSubIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRelatedPartyAltSubIDs (1572)
func (f NoRelatedPartyAltSubIDsField) Tag() quickfix.Tag { return tag.NoRelatedPartyAltSubIDs }

//NewNoRelatedPartyAltSubIDs returns a new NoRelatedPartyAltSubIDsField initialized with val
func NewNoRelatedPartyAltSubIDs(val int) *NoRelatedPartyAltSubIDsField {
	field := &NoRelatedPartyAltSubIDsField{}
	field.Value = val
	return field
}

//NoRelatedPartyIDsField is a NUMINGROUP field
type NoRelatedPartyIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRelatedPartyIDs (1562)
func (f NoRelatedPartyIDsField) Tag() quickfix.Tag { return tag.NoRelatedPartyIDs }

//NewNoRelatedPartyIDs returns a new NoRelatedPartyIDsField initialized with val
func NewNoRelatedPartyIDs(val int) *NoRelatedPartyIDsField {
	field := &NoRelatedPartyIDsField{}
	field.Value = val
	return field
}

//NoRelatedPartySubIDsField is a NUMINGROUP field
type NoRelatedPartySubIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRelatedPartySubIDs (1566)
func (f NoRelatedPartySubIDsField) Tag() quickfix.Tag { return tag.NoRelatedPartySubIDs }

//NewNoRelatedPartySubIDs returns a new NoRelatedPartySubIDsField initialized with val
func NewNoRelatedPartySubIDs(val int) *NoRelatedPartySubIDsField {
	field := &NoRelatedPartySubIDsField{}
	field.Value = val
	return field
}

//NoRelatedSymField is a NUMINGROUP field
type NoRelatedSymField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRelatedSym (146)
func (f NoRelatedSymField) Tag() quickfix.Tag { return tag.NoRelatedSym }

//NewNoRelatedSym returns a new NoRelatedSymField initialized with val
func NewNoRelatedSym(val int) *NoRelatedSymField {
	field := &NoRelatedSymField{}
	field.Value = val
	return field
}

//NoRelationshipRiskInstrumentsField is a NUMINGROUP field
type NoRelationshipRiskInstrumentsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRelationshipRiskInstruments (1587)
func (f NoRelationshipRiskInstrumentsField) Tag() quickfix.Tag {
	return tag.NoRelationshipRiskInstruments
}

//NewNoRelationshipRiskInstruments returns a new NoRelationshipRiskInstrumentsField initialized with val
func NewNoRelationshipRiskInstruments(val int) *NoRelationshipRiskInstrumentsField {
	field := &NoRelationshipRiskInstrumentsField{}
	field.Value = val
	return field
}

//NoRelationshipRiskLimitsField is a NUMINGROUP field
type NoRelationshipRiskLimitsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRelationshipRiskLimits (1582)
func (f NoRelationshipRiskLimitsField) Tag() quickfix.Tag { return tag.NoRelationshipRiskLimits }

//NewNoRelationshipRiskLimits returns a new NoRelationshipRiskLimitsField initialized with val
func NewNoRelationshipRiskLimits(val int) *NoRelationshipRiskLimitsField {
	field := &NoRelationshipRiskLimitsField{}
	field.Value = val
	return field
}

//NoRelationshipRiskSecurityAltIDField is a NUMINGROUP field
type NoRelationshipRiskSecurityAltIDField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRelationshipRiskSecurityAltID (1593)
func (f NoRelationshipRiskSecurityAltIDField) Tag() quickfix.Tag {
	return tag.NoRelationshipRiskSecurityAltID
}

//NewNoRelationshipRiskSecurityAltID returns a new NoRelationshipRiskSecurityAltIDField initialized with val
func NewNoRelationshipRiskSecurityAltID(val int) *NoRelationshipRiskSecurityAltIDField {
	field := &NoRelationshipRiskSecurityAltIDField{}
	field.Value = val
	return field
}

//NoRelationshipRiskWarningLevelsField is a NUMINGROUP field
type NoRelationshipRiskWarningLevelsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRelationshipRiskWarningLevels (1613)
func (f NoRelationshipRiskWarningLevelsField) Tag() quickfix.Tag {
	return tag.NoRelationshipRiskWarningLevels
}

//NewNoRelationshipRiskWarningLevels returns a new NoRelationshipRiskWarningLevelsField initialized with val
func NewNoRelationshipRiskWarningLevels(val int) *NoRelationshipRiskWarningLevelsField {
	field := &NoRelationshipRiskWarningLevelsField{}
	field.Value = val
	return field
}

//NoRequestedPartyRolesField is a NUMINGROUP field
type NoRequestedPartyRolesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRequestedPartyRoles (1508)
func (f NoRequestedPartyRolesField) Tag() quickfix.Tag { return tag.NoRequestedPartyRoles }

//NewNoRequestedPartyRoles returns a new NoRequestedPartyRolesField initialized with val
func NewNoRequestedPartyRoles(val int) *NoRequestedPartyRolesField {
	field := &NoRequestedPartyRolesField{}
	field.Value = val
	return field
}

//NoRiskInstrumentsField is a NUMINGROUP field
type NoRiskInstrumentsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRiskInstruments (1534)
func (f NoRiskInstrumentsField) Tag() quickfix.Tag { return tag.NoRiskInstruments }

//NewNoRiskInstruments returns a new NoRiskInstrumentsField initialized with val
func NewNoRiskInstruments(val int) *NoRiskInstrumentsField {
	field := &NoRiskInstrumentsField{}
	field.Value = val
	return field
}

//NoRiskLimitsField is a NUMINGROUP field
type NoRiskLimitsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRiskLimits (1529)
func (f NoRiskLimitsField) Tag() quickfix.Tag { return tag.NoRiskLimits }

//NewNoRiskLimits returns a new NoRiskLimitsField initialized with val
func NewNoRiskLimits(val int) *NoRiskLimitsField {
	field := &NoRiskLimitsField{}
	field.Value = val
	return field
}

//NoRiskSecurityAltIDField is a NUMINGROUP field
type NoRiskSecurityAltIDField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRiskSecurityAltID (1540)
func (f NoRiskSecurityAltIDField) Tag() quickfix.Tag { return tag.NoRiskSecurityAltID }

//NewNoRiskSecurityAltID returns a new NoRiskSecurityAltIDField initialized with val
func NewNoRiskSecurityAltID(val int) *NoRiskSecurityAltIDField {
	field := &NoRiskSecurityAltIDField{}
	field.Value = val
	return field
}

//NoRiskWarningLevelsField is a NUMINGROUP field
type NoRiskWarningLevelsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRiskWarningLevels (1559)
func (f NoRiskWarningLevelsField) Tag() quickfix.Tag { return tag.NoRiskWarningLevels }

//NewNoRiskWarningLevels returns a new NoRiskWarningLevelsField initialized with val
func NewNoRiskWarningLevels(val int) *NoRiskWarningLevelsField {
	field := &NoRiskWarningLevelsField{}
	field.Value = val
	return field
}

//NoRootPartyIDsField is a NUMINGROUP field
type NoRootPartyIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRootPartyIDs (1116)
func (f NoRootPartyIDsField) Tag() quickfix.Tag { return tag.NoRootPartyIDs }

//NewNoRootPartyIDs returns a new NoRootPartyIDsField initialized with val
func NewNoRootPartyIDs(val int) *NoRootPartyIDsField {
	field := &NoRootPartyIDsField{}
	field.Value = val
	return field
}

//NoRootPartySubIDsField is a NUMINGROUP field
type NoRootPartySubIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRootPartySubIDs (1120)
func (f NoRootPartySubIDsField) Tag() quickfix.Tag { return tag.NoRootPartySubIDs }

//NewNoRootPartySubIDs returns a new NoRootPartySubIDsField initialized with val
func NewNoRootPartySubIDs(val int) *NoRootPartySubIDsField {
	field := &NoRootPartySubIDsField{}
	field.Value = val
	return field
}

//NoRoutingIDsField is a NUMINGROUP field
type NoRoutingIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoRoutingIDs (215)
func (f NoRoutingIDsField) Tag() quickfix.Tag { return tag.NoRoutingIDs }

//NewNoRoutingIDs returns a new NoRoutingIDsField initialized with val
func NewNoRoutingIDs(val int) *NoRoutingIDsField {
	field := &NoRoutingIDsField{}
	field.Value = val
	return field
}

//NoRptsField is a INT field
type NoRptsField struct{ quickfix.IntValue }

//Tag returns tag.NoRpts (82)
func (f NoRptsField) Tag() quickfix.Tag { return tag.NoRpts }

//NewNoRpts returns a new NoRptsField initialized with val
func NewNoRpts(val int) *NoRptsField {
	field := &NoRptsField{}
	field.Value = val
	return field
}

//NoSecurityAltIDField is a NUMINGROUP field
type NoSecurityAltIDField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoSecurityAltID (454)
func (f NoSecurityAltIDField) Tag() quickfix.Tag { return tag.NoSecurityAltID }

//NewNoSecurityAltID returns a new NoSecurityAltIDField initialized with val
func NewNoSecurityAltID(val int) *NoSecurityAltIDField {
	field := &NoSecurityAltIDField{}
	field.Value = val
	return field
}

//NoSecurityTypesField is a NUMINGROUP field
type NoSecurityTypesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoSecurityTypes (558)
func (f NoSecurityTypesField) Tag() quickfix.Tag { return tag.NoSecurityTypes }

//NewNoSecurityTypes returns a new NoSecurityTypesField initialized with val
func NewNoSecurityTypes(val int) *NoSecurityTypesField {
	field := &NoSecurityTypesField{}
	field.Value = val
	return field
}

//NoSettlDetailsField is a NUMINGROUP field
type NoSettlDetailsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoSettlDetails (1158)
func (f NoSettlDetailsField) Tag() quickfix.Tag { return tag.NoSettlDetails }

//NewNoSettlDetails returns a new NoSettlDetailsField initialized with val
func NewNoSettlDetails(val int) *NoSettlDetailsField {
	field := &NoSettlDetailsField{}
	field.Value = val
	return field
}

//NoSettlInstField is a NUMINGROUP field
type NoSettlInstField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoSettlInst (778)
func (f NoSettlInstField) Tag() quickfix.Tag { return tag.NoSettlInst }

//NewNoSettlInst returns a new NoSettlInstField initialized with val
func NewNoSettlInst(val int) *NoSettlInstField {
	field := &NoSettlInstField{}
	field.Value = val
	return field
}

//NoSettlObligField is a NUMINGROUP field
type NoSettlObligField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoSettlOblig (1165)
func (f NoSettlObligField) Tag() quickfix.Tag { return tag.NoSettlOblig }

//NewNoSettlOblig returns a new NoSettlObligField initialized with val
func NewNoSettlOblig(val int) *NoSettlObligField {
	field := &NoSettlObligField{}
	field.Value = val
	return field
}

//NoSettlPartyIDsField is a NUMINGROUP field
type NoSettlPartyIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoSettlPartyIDs (781)
func (f NoSettlPartyIDsField) Tag() quickfix.Tag { return tag.NoSettlPartyIDs }

//NewNoSettlPartyIDs returns a new NoSettlPartyIDsField initialized with val
func NewNoSettlPartyIDs(val int) *NoSettlPartyIDsField {
	field := &NoSettlPartyIDsField{}
	field.Value = val
	return field
}

//NoSettlPartySubIDsField is a NUMINGROUP field
type NoSettlPartySubIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoSettlPartySubIDs (801)
func (f NoSettlPartySubIDsField) Tag() quickfix.Tag { return tag.NoSettlPartySubIDs }

//NewNoSettlPartySubIDs returns a new NoSettlPartySubIDsField initialized with val
func NewNoSettlPartySubIDs(val int) *NoSettlPartySubIDsField {
	field := &NoSettlPartySubIDsField{}
	field.Value = val
	return field
}

//NoSideTrdRegTSField is a NUMINGROUP field
type NoSideTrdRegTSField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoSideTrdRegTS (1016)
func (f NoSideTrdRegTSField) Tag() quickfix.Tag { return tag.NoSideTrdRegTS }

//NewNoSideTrdRegTS returns a new NoSideTrdRegTSField initialized with val
func NewNoSideTrdRegTS(val int) *NoSideTrdRegTSField {
	field := &NoSideTrdRegTSField{}
	field.Value = val
	return field
}

//NoSidesField is a NUMINGROUP field
type NoSidesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoSides (552)
func (f NoSidesField) Tag() quickfix.Tag { return tag.NoSides }

//NewNoSides returns a new NoSidesField initialized with val
func NewNoSides(val int) *NoSidesField {
	field := &NoSidesField{}
	field.Value = val
	return field
}

//NoStatsIndicatorsField is a NUMINGROUP field
type NoStatsIndicatorsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoStatsIndicators (1175)
func (f NoStatsIndicatorsField) Tag() quickfix.Tag { return tag.NoStatsIndicators }

//NewNoStatsIndicators returns a new NoStatsIndicatorsField initialized with val
func NewNoStatsIndicators(val int) *NoStatsIndicatorsField {
	field := &NoStatsIndicatorsField{}
	field.Value = val
	return field
}

//NoStipulationsField is a NUMINGROUP field
type NoStipulationsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoStipulations (232)
func (f NoStipulationsField) Tag() quickfix.Tag { return tag.NoStipulations }

//NewNoStipulations returns a new NoStipulationsField initialized with val
func NewNoStipulations(val int) *NoStipulationsField {
	field := &NoStipulationsField{}
	field.Value = val
	return field
}

//NoStrategyParametersField is a NUMINGROUP field
type NoStrategyParametersField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoStrategyParameters (957)
func (f NoStrategyParametersField) Tag() quickfix.Tag { return tag.NoStrategyParameters }

//NewNoStrategyParameters returns a new NoStrategyParametersField initialized with val
func NewNoStrategyParameters(val int) *NoStrategyParametersField {
	field := &NoStrategyParametersField{}
	field.Value = val
	return field
}

//NoStrikeRulesField is a NUMINGROUP field
type NoStrikeRulesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoStrikeRules (1201)
func (f NoStrikeRulesField) Tag() quickfix.Tag { return tag.NoStrikeRules }

//NewNoStrikeRules returns a new NoStrikeRulesField initialized with val
func NewNoStrikeRules(val int) *NoStrikeRulesField {
	field := &NoStrikeRulesField{}
	field.Value = val
	return field
}

//NoStrikesField is a NUMINGROUP field
type NoStrikesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoStrikes (428)
func (f NoStrikesField) Tag() quickfix.Tag { return tag.NoStrikes }

//NewNoStrikes returns a new NoStrikesField initialized with val
func NewNoStrikes(val int) *NoStrikesField {
	field := &NoStrikesField{}
	field.Value = val
	return field
}

//NoTargetPartyIDsField is a NUMINGROUP field
type NoTargetPartyIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoTargetPartyIDs (1461)
func (f NoTargetPartyIDsField) Tag() quickfix.Tag { return tag.NoTargetPartyIDs }

//NewNoTargetPartyIDs returns a new NoTargetPartyIDsField initialized with val
func NewNoTargetPartyIDs(val int) *NoTargetPartyIDsField {
	field := &NoTargetPartyIDsField{}
	field.Value = val
	return field
}

//NoTickRulesField is a NUMINGROUP field
type NoTickRulesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoTickRules (1205)
func (f NoTickRulesField) Tag() quickfix.Tag { return tag.NoTickRules }

//NewNoTickRules returns a new NoTickRulesField initialized with val
func NewNoTickRules(val int) *NoTickRulesField {
	field := &NoTickRulesField{}
	field.Value = val
	return field
}

//NoTimeInForceRulesField is a NUMINGROUP field
type NoTimeInForceRulesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoTimeInForceRules (1239)
func (f NoTimeInForceRulesField) Tag() quickfix.Tag { return tag.NoTimeInForceRules }

//NewNoTimeInForceRules returns a new NoTimeInForceRulesField initialized with val
func NewNoTimeInForceRules(val int) *NoTimeInForceRulesField {
	field := &NoTimeInForceRulesField{}
	field.Value = val
	return field
}

//NoTradesField is a NUMINGROUP field
type NoTradesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoTrades (897)
func (f NoTradesField) Tag() quickfix.Tag { return tag.NoTrades }

//NewNoTrades returns a new NoTradesField initialized with val
func NewNoTrades(val int) *NoTradesField {
	field := &NoTradesField{}
	field.Value = val
	return field
}

//NoTradingSessionRulesField is a NUMINGROUP field
type NoTradingSessionRulesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoTradingSessionRules (1309)
func (f NoTradingSessionRulesField) Tag() quickfix.Tag { return tag.NoTradingSessionRules }

//NewNoTradingSessionRules returns a new NoTradingSessionRulesField initialized with val
func NewNoTradingSessionRules(val int) *NoTradingSessionRulesField {
	field := &NoTradingSessionRulesField{}
	field.Value = val
	return field
}

//NoTradingSessionsField is a NUMINGROUP field
type NoTradingSessionsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoTradingSessions (386)
func (f NoTradingSessionsField) Tag() quickfix.Tag { return tag.NoTradingSessions }

//NewNoTradingSessions returns a new NoTradingSessionsField initialized with val
func NewNoTradingSessions(val int) *NoTradingSessionsField {
	field := &NoTradingSessionsField{}
	field.Value = val
	return field
}

//NoTrdRegTimestampsField is a NUMINGROUP field
type NoTrdRegTimestampsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoTrdRegTimestamps (768)
func (f NoTrdRegTimestampsField) Tag() quickfix.Tag { return tag.NoTrdRegTimestamps }

//NewNoTrdRegTimestamps returns a new NoTrdRegTimestampsField initialized with val
func NewNoTrdRegTimestamps(val int) *NoTrdRegTimestampsField {
	field := &NoTrdRegTimestampsField{}
	field.Value = val
	return field
}

//NoTrdRepIndicatorsField is a NUMINGROUP field
type NoTrdRepIndicatorsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoTrdRepIndicators (1387)
func (f NoTrdRepIndicatorsField) Tag() quickfix.Tag { return tag.NoTrdRepIndicators }

//NewNoTrdRepIndicators returns a new NoTrdRepIndicatorsField initialized with val
func NewNoTrdRepIndicators(val int) *NoTrdRepIndicatorsField {
	field := &NoTrdRepIndicatorsField{}
	field.Value = val
	return field
}

//NoUnderlyingAmountsField is a NUMINGROUP field
type NoUnderlyingAmountsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoUnderlyingAmounts (984)
func (f NoUnderlyingAmountsField) Tag() quickfix.Tag { return tag.NoUnderlyingAmounts }

//NewNoUnderlyingAmounts returns a new NoUnderlyingAmountsField initialized with val
func NewNoUnderlyingAmounts(val int) *NoUnderlyingAmountsField {
	field := &NoUnderlyingAmountsField{}
	field.Value = val
	return field
}

//NoUnderlyingLegSecurityAltIDField is a NUMINGROUP field
type NoUnderlyingLegSecurityAltIDField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoUnderlyingLegSecurityAltID (1334)
func (f NoUnderlyingLegSecurityAltIDField) Tag() quickfix.Tag { return tag.NoUnderlyingLegSecurityAltID }

//NewNoUnderlyingLegSecurityAltID returns a new NoUnderlyingLegSecurityAltIDField initialized with val
func NewNoUnderlyingLegSecurityAltID(val int) *NoUnderlyingLegSecurityAltIDField {
	field := &NoUnderlyingLegSecurityAltIDField{}
	field.Value = val
	return field
}

//NoUnderlyingSecurityAltIDField is a NUMINGROUP field
type NoUnderlyingSecurityAltIDField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoUnderlyingSecurityAltID (457)
func (f NoUnderlyingSecurityAltIDField) Tag() quickfix.Tag { return tag.NoUnderlyingSecurityAltID }

//NewNoUnderlyingSecurityAltID returns a new NoUnderlyingSecurityAltIDField initialized with val
func NewNoUnderlyingSecurityAltID(val int) *NoUnderlyingSecurityAltIDField {
	field := &NoUnderlyingSecurityAltIDField{}
	field.Value = val
	return field
}

//NoUnderlyingStipsField is a NUMINGROUP field
type NoUnderlyingStipsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoUnderlyingStips (887)
func (f NoUnderlyingStipsField) Tag() quickfix.Tag { return tag.NoUnderlyingStips }

//NewNoUnderlyingStips returns a new NoUnderlyingStipsField initialized with val
func NewNoUnderlyingStips(val int) *NoUnderlyingStipsField {
	field := &NoUnderlyingStipsField{}
	field.Value = val
	return field
}

//NoUnderlyingsField is a NUMINGROUP field
type NoUnderlyingsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoUnderlyings (711)
func (f NoUnderlyingsField) Tag() quickfix.Tag { return tag.NoUnderlyings }

//NewNoUnderlyings returns a new NoUnderlyingsField initialized with val
func NewNoUnderlyings(val int) *NoUnderlyingsField {
	field := &NoUnderlyingsField{}
	field.Value = val
	return field
}

//NoUndlyInstrumentPartiesField is a NUMINGROUP field
type NoUndlyInstrumentPartiesField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoUndlyInstrumentParties (1058)
func (f NoUndlyInstrumentPartiesField) Tag() quickfix.Tag { return tag.NoUndlyInstrumentParties }

//NewNoUndlyInstrumentParties returns a new NoUndlyInstrumentPartiesField initialized with val
func NewNoUndlyInstrumentParties(val int) *NoUndlyInstrumentPartiesField {
	field := &NoUndlyInstrumentPartiesField{}
	field.Value = val
	return field
}

//NoUndlyInstrumentPartySubIDsField is a NUMINGROUP field
type NoUndlyInstrumentPartySubIDsField struct{ quickfix.NumInGroupValue }

//Tag returns tag.NoUndlyInstrumentPartySubIDs (1062)
func (f NoUndlyInstrumentPartySubIDsField) Tag() quickfix.Tag { return tag.NoUndlyInstrumentPartySubIDs }

//NewNoUndlyInstrumentPartySubIDs returns a new NoUndlyInstrumentPartySubIDsField initialized with val
func NewNoUndlyInstrumentPartySubIDs(val int) *NoUndlyInstrumentPartySubIDsField {
	field := &NoUndlyInstrumentPartySubIDsField{}
	field.Value = val
	return field
}

//NotAffOrigClOrdIDField is a STRING field
type NotAffOrigClOrdIDField struct{ quickfix.StringValue }

//Tag returns tag.NotAffOrigClOrdID (1372)
func (f NotAffOrigClOrdIDField) Tag() quickfix.Tag { return tag.NotAffOrigClOrdID }

//NewNotAffOrigClOrdID returns a new NotAffOrigClOrdIDField initialized with val
func NewNotAffOrigClOrdID(val string) *NotAffOrigClOrdIDField {
	field := &NotAffOrigClOrdIDField{}
	field.Value = val
	return field
}

//NotAffectedOrderIDField is a STRING field
type NotAffectedOrderIDField struct{ quickfix.StringValue }

//Tag returns tag.NotAffectedOrderID (1371)
func (f NotAffectedOrderIDField) Tag() quickfix.Tag { return tag.NotAffectedOrderID }

//NewNotAffectedOrderID returns a new NotAffectedOrderIDField initialized with val
func NewNotAffectedOrderID(val string) *NotAffectedOrderIDField {
	field := &NotAffectedOrderIDField{}
	field.Value = val
	return field
}

//NotifyBrokerOfCreditField is a BOOLEAN field
type NotifyBrokerOfCreditField struct{ quickfix.BooleanValue }

//Tag returns tag.NotifyBrokerOfCredit (208)
func (f NotifyBrokerOfCreditField) Tag() quickfix.Tag { return tag.NotifyBrokerOfCredit }

//NewNotifyBrokerOfCredit returns a new NotifyBrokerOfCreditField initialized with val
func NewNotifyBrokerOfCredit(val bool) *NotifyBrokerOfCreditField {
	field := &NotifyBrokerOfCreditField{}
	field.Value = val
	return field
}

//NotionalPercentageOutstandingField is a PERCENTAGE field
type NotionalPercentageOutstandingField struct{ quickfix.PercentageValue }

//Tag returns tag.NotionalPercentageOutstanding (1451)
func (f NotionalPercentageOutstandingField) Tag() quickfix.Tag {
	return tag.NotionalPercentageOutstanding
}

//NewNotionalPercentageOutstanding returns a new NotionalPercentageOutstandingField initialized with val
func NewNotionalPercentageOutstanding(val float64) *NotionalPercentageOutstandingField {
	field := &NotionalPercentageOutstandingField{}
	field.Value = val
	return field
}

//NumBiddersField is a INT field
type NumBiddersField struct{ quickfix.IntValue }

//Tag returns tag.NumBidders (417)
func (f NumBiddersField) Tag() quickfix.Tag { return tag.NumBidders }

//NewNumBidders returns a new NumBiddersField initialized with val
func NewNumBidders(val int) *NumBiddersField {
	field := &NumBiddersField{}
	field.Value = val
	return field
}

//NumDaysInterestField is a INT field
type NumDaysInterestField struct{ quickfix.IntValue }

//Tag returns tag.NumDaysInterest (157)
func (f NumDaysInterestField) Tag() quickfix.Tag { return tag.NumDaysInterest }

//NewNumDaysInterest returns a new NumDaysInterestField initialized with val
func NewNumDaysInterest(val int) *NumDaysInterestField {
	field := &NumDaysInterestField{}
	field.Value = val
	return field
}

//NumTicketsField is a INT field
type NumTicketsField struct{ quickfix.IntValue }

//Tag returns tag.NumTickets (395)
func (f NumTicketsField) Tag() quickfix.Tag { return tag.NumTickets }

//NewNumTickets returns a new NumTicketsField initialized with val
func NewNumTickets(val int) *NumTicketsField {
	field := &NumTicketsField{}
	field.Value = val
	return field
}

//NumberOfOrdersField is a INT field
type NumberOfOrdersField struct{ quickfix.IntValue }

//Tag returns tag.NumberOfOrders (346)
func (f NumberOfOrdersField) Tag() quickfix.Tag { return tag.NumberOfOrders }

//NewNumberOfOrders returns a new NumberOfOrdersField initialized with val
func NewNumberOfOrders(val int) *NumberOfOrdersField {
	field := &NumberOfOrdersField{}
	field.Value = val
	return field
}

//OddLotField is a BOOLEAN field
type OddLotField struct{ quickfix.BooleanValue }

//Tag returns tag.OddLot (575)
func (f OddLotField) Tag() quickfix.Tag { return tag.OddLot }

//NewOddLot returns a new OddLotField initialized with val
func NewOddLot(val bool) *OddLotField {
	field := &OddLotField{}
	field.Value = val
	return field
}

//OfferForwardPointsField is a PRICEOFFSET field
type OfferForwardPointsField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.OfferForwardPoints (191)
func (f OfferForwardPointsField) Tag() quickfix.Tag { return tag.OfferForwardPoints }

//NewOfferForwardPoints returns a new OfferForwardPointsField initialized with val
func NewOfferForwardPoints(val float64) *OfferForwardPointsField {
	field := &OfferForwardPointsField{}
	field.Value = val
	return field
}

//OfferForwardPoints2Field is a PRICEOFFSET field
type OfferForwardPoints2Field struct{ quickfix.PriceOffsetValue }

//Tag returns tag.OfferForwardPoints2 (643)
func (f OfferForwardPoints2Field) Tag() quickfix.Tag { return tag.OfferForwardPoints2 }

//NewOfferForwardPoints2 returns a new OfferForwardPoints2Field initialized with val
func NewOfferForwardPoints2(val float64) *OfferForwardPoints2Field {
	field := &OfferForwardPoints2Field{}
	field.Value = val
	return field
}

//OfferPxField is a PRICE field
type OfferPxField struct{ quickfix.PriceValue }

//Tag returns tag.OfferPx (133)
func (f OfferPxField) Tag() quickfix.Tag { return tag.OfferPx }

//NewOfferPx returns a new OfferPxField initialized with val
func NewOfferPx(val float64) *OfferPxField {
	field := &OfferPxField{}
	field.Value = val
	return field
}

//OfferSizeField is a QTY field
type OfferSizeField struct{ quickfix.QtyValue }

//Tag returns tag.OfferSize (135)
func (f OfferSizeField) Tag() quickfix.Tag { return tag.OfferSize }

//NewOfferSize returns a new OfferSizeField initialized with val
func NewOfferSize(val float64) *OfferSizeField {
	field := &OfferSizeField{}
	field.Value = val
	return field
}

//OfferSpotRateField is a PRICE field
type OfferSpotRateField struct{ quickfix.PriceValue }

//Tag returns tag.OfferSpotRate (190)
func (f OfferSpotRateField) Tag() quickfix.Tag { return tag.OfferSpotRate }

//NewOfferSpotRate returns a new OfferSpotRateField initialized with val
func NewOfferSpotRate(val float64) *OfferSpotRateField {
	field := &OfferSpotRateField{}
	field.Value = val
	return field
}

//OfferSwapPointsField is a PRICEOFFSET field
type OfferSwapPointsField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.OfferSwapPoints (1066)
func (f OfferSwapPointsField) Tag() quickfix.Tag { return tag.OfferSwapPoints }

//NewOfferSwapPoints returns a new OfferSwapPointsField initialized with val
func NewOfferSwapPoints(val float64) *OfferSwapPointsField {
	field := &OfferSwapPointsField{}
	field.Value = val
	return field
}

//OfferYieldField is a PERCENTAGE field
type OfferYieldField struct{ quickfix.PercentageValue }

//Tag returns tag.OfferYield (634)
func (f OfferYieldField) Tag() quickfix.Tag { return tag.OfferYield }

//NewOfferYield returns a new OfferYieldField initialized with val
func NewOfferYield(val float64) *OfferYieldField {
	field := &OfferYieldField{}
	field.Value = val
	return field
}

//OnBehalfOfCompIDField is a STRING field
type OnBehalfOfCompIDField struct{ quickfix.StringValue }

//Tag returns tag.OnBehalfOfCompID (115)
func (f OnBehalfOfCompIDField) Tag() quickfix.Tag { return tag.OnBehalfOfCompID }

//NewOnBehalfOfCompID returns a new OnBehalfOfCompIDField initialized with val
func NewOnBehalfOfCompID(val string) *OnBehalfOfCompIDField {
	field := &OnBehalfOfCompIDField{}
	field.Value = val
	return field
}

//OnBehalfOfLocationIDField is a STRING field
type OnBehalfOfLocationIDField struct{ quickfix.StringValue }

//Tag returns tag.OnBehalfOfLocationID (144)
func (f OnBehalfOfLocationIDField) Tag() quickfix.Tag { return tag.OnBehalfOfLocationID }

//NewOnBehalfOfLocationID returns a new OnBehalfOfLocationIDField initialized with val
func NewOnBehalfOfLocationID(val string) *OnBehalfOfLocationIDField {
	field := &OnBehalfOfLocationIDField{}
	field.Value = val
	return field
}

//OnBehalfOfSendingTimeField is a UTCTIMESTAMP field
type OnBehalfOfSendingTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.OnBehalfOfSendingTime (370)
func (f OnBehalfOfSendingTimeField) Tag() quickfix.Tag { return tag.OnBehalfOfSendingTime }

//OnBehalfOfSubIDField is a STRING field
type OnBehalfOfSubIDField struct{ quickfix.StringValue }

//Tag returns tag.OnBehalfOfSubID (116)
func (f OnBehalfOfSubIDField) Tag() quickfix.Tag { return tag.OnBehalfOfSubID }

//NewOnBehalfOfSubID returns a new OnBehalfOfSubIDField initialized with val
func NewOnBehalfOfSubID(val string) *OnBehalfOfSubIDField {
	field := &OnBehalfOfSubIDField{}
	field.Value = val
	return field
}

//OpenCloseField is a CHAR field
type OpenCloseField struct{ quickfix.CharValue }

//Tag returns tag.OpenClose (77)
func (f OpenCloseField) Tag() quickfix.Tag { return tag.OpenClose }

//NewOpenClose returns a new OpenCloseField initialized with val
func NewOpenClose(val string) *OpenCloseField {
	field := &OpenCloseField{}
	field.Value = val
	return field
}

//OpenCloseSettlFlagField is a MULTIPLECHARVALUE field
type OpenCloseSettlFlagField struct{ quickfix.MultipleCharValue }

//Tag returns tag.OpenCloseSettlFlag (286)
func (f OpenCloseSettlFlagField) Tag() quickfix.Tag { return tag.OpenCloseSettlFlag }

//NewOpenCloseSettlFlag returns a new OpenCloseSettlFlagField initialized with val
func NewOpenCloseSettlFlag(val string) *OpenCloseSettlFlagField {
	field := &OpenCloseSettlFlagField{}
	field.Value = val
	return field
}

//OpenCloseSettleFlagField is a MULTIPLEVALUESTRING field
type OpenCloseSettleFlagField struct{ quickfix.MultipleStringValue }

//Tag returns tag.OpenCloseSettleFlag (286)
func (f OpenCloseSettleFlagField) Tag() quickfix.Tag { return tag.OpenCloseSettleFlag }

//NewOpenCloseSettleFlag returns a new OpenCloseSettleFlagField initialized with val
func NewOpenCloseSettleFlag(val string) *OpenCloseSettleFlagField {
	field := &OpenCloseSettleFlagField{}
	field.Value = val
	return field
}

//OpenInterestField is a AMT field
type OpenInterestField struct{ quickfix.AmtValue }

//Tag returns tag.OpenInterest (746)
func (f OpenInterestField) Tag() quickfix.Tag { return tag.OpenInterest }

//NewOpenInterest returns a new OpenInterestField initialized with val
func NewOpenInterest(val float64) *OpenInterestField {
	field := &OpenInterestField{}
	field.Value = val
	return field
}

//OptAttributeField is a CHAR field
type OptAttributeField struct{ quickfix.CharValue }

//Tag returns tag.OptAttribute (206)
func (f OptAttributeField) Tag() quickfix.Tag { return tag.OptAttribute }

//NewOptAttribute returns a new OptAttributeField initialized with val
func NewOptAttribute(val string) *OptAttributeField {
	field := &OptAttributeField{}
	field.Value = val
	return field
}

//OptPayAmountField is a AMT field
type OptPayAmountField struct{ quickfix.AmtValue }

//Tag returns tag.OptPayAmount (1195)
func (f OptPayAmountField) Tag() quickfix.Tag { return tag.OptPayAmount }

//NewOptPayAmount returns a new OptPayAmountField initialized with val
func NewOptPayAmount(val float64) *OptPayAmountField {
	field := &OptPayAmountField{}
	field.Value = val
	return field
}

//OptPayoutAmountField is a AMT field
type OptPayoutAmountField struct{ quickfix.AmtValue }

//Tag returns tag.OptPayoutAmount (1195)
func (f OptPayoutAmountField) Tag() quickfix.Tag { return tag.OptPayoutAmount }

//NewOptPayoutAmount returns a new OptPayoutAmountField initialized with val
func NewOptPayoutAmount(val float64) *OptPayoutAmountField {
	field := &OptPayoutAmountField{}
	field.Value = val
	return field
}

//OptPayoutTypeField is a INT field
type OptPayoutTypeField struct{ quickfix.IntValue }

//Tag returns tag.OptPayoutType (1482)
func (f OptPayoutTypeField) Tag() quickfix.Tag { return tag.OptPayoutType }

//NewOptPayoutType returns a new OptPayoutTypeField initialized with val
func NewOptPayoutType(val int) *OptPayoutTypeField {
	field := &OptPayoutTypeField{}
	field.Value = val
	return field
}

//OrdRejReasonField is a INT field
type OrdRejReasonField struct{ quickfix.IntValue }

//Tag returns tag.OrdRejReason (103)
func (f OrdRejReasonField) Tag() quickfix.Tag { return tag.OrdRejReason }

//NewOrdRejReason returns a new OrdRejReasonField initialized with val
func NewOrdRejReason(val int) *OrdRejReasonField {
	field := &OrdRejReasonField{}
	field.Value = val
	return field
}

//OrdStatusField is a CHAR field
type OrdStatusField struct{ quickfix.CharValue }

//Tag returns tag.OrdStatus (39)
func (f OrdStatusField) Tag() quickfix.Tag { return tag.OrdStatus }

//NewOrdStatus returns a new OrdStatusField initialized with val
func NewOrdStatus(val string) *OrdStatusField {
	field := &OrdStatusField{}
	field.Value = val
	return field
}

//OrdStatusReqIDField is a STRING field
type OrdStatusReqIDField struct{ quickfix.StringValue }

//Tag returns tag.OrdStatusReqID (790)
func (f OrdStatusReqIDField) Tag() quickfix.Tag { return tag.OrdStatusReqID }

//NewOrdStatusReqID returns a new OrdStatusReqIDField initialized with val
func NewOrdStatusReqID(val string) *OrdStatusReqIDField {
	field := &OrdStatusReqIDField{}
	field.Value = val
	return field
}

//OrdTypeField is a CHAR field
type OrdTypeField struct{ quickfix.CharValue }

//Tag returns tag.OrdType (40)
func (f OrdTypeField) Tag() quickfix.Tag { return tag.OrdType }

//NewOrdType returns a new OrdTypeField initialized with val
func NewOrdType(val string) *OrdTypeField {
	field := &OrdTypeField{}
	field.Value = val
	return field
}

//OrderAvgPxField is a PRICE field
type OrderAvgPxField struct{ quickfix.PriceValue }

//Tag returns tag.OrderAvgPx (799)
func (f OrderAvgPxField) Tag() quickfix.Tag { return tag.OrderAvgPx }

//NewOrderAvgPx returns a new OrderAvgPxField initialized with val
func NewOrderAvgPx(val float64) *OrderAvgPxField {
	field := &OrderAvgPxField{}
	field.Value = val
	return field
}

//OrderBookingQtyField is a QTY field
type OrderBookingQtyField struct{ quickfix.QtyValue }

//Tag returns tag.OrderBookingQty (800)
func (f OrderBookingQtyField) Tag() quickfix.Tag { return tag.OrderBookingQty }

//NewOrderBookingQty returns a new OrderBookingQtyField initialized with val
func NewOrderBookingQty(val float64) *OrderBookingQtyField {
	field := &OrderBookingQtyField{}
	field.Value = val
	return field
}

//OrderCapacityField is a CHAR field
type OrderCapacityField struct{ quickfix.CharValue }

//Tag returns tag.OrderCapacity (528)
func (f OrderCapacityField) Tag() quickfix.Tag { return tag.OrderCapacity }

//NewOrderCapacity returns a new OrderCapacityField initialized with val
func NewOrderCapacity(val string) *OrderCapacityField {
	field := &OrderCapacityField{}
	field.Value = val
	return field
}

//OrderCapacityQtyField is a QTY field
type OrderCapacityQtyField struct{ quickfix.QtyValue }

//Tag returns tag.OrderCapacityQty (863)
func (f OrderCapacityQtyField) Tag() quickfix.Tag { return tag.OrderCapacityQty }

//NewOrderCapacityQty returns a new OrderCapacityQtyField initialized with val
func NewOrderCapacityQty(val float64) *OrderCapacityQtyField {
	field := &OrderCapacityQtyField{}
	field.Value = val
	return field
}

//OrderCategoryField is a CHAR field
type OrderCategoryField struct{ quickfix.CharValue }

//Tag returns tag.OrderCategory (1115)
func (f OrderCategoryField) Tag() quickfix.Tag { return tag.OrderCategory }

//NewOrderCategory returns a new OrderCategoryField initialized with val
func NewOrderCategory(val string) *OrderCategoryField {
	field := &OrderCategoryField{}
	field.Value = val
	return field
}

//OrderDelayField is a INT field
type OrderDelayField struct{ quickfix.IntValue }

//Tag returns tag.OrderDelay (1428)
func (f OrderDelayField) Tag() quickfix.Tag { return tag.OrderDelay }

//NewOrderDelay returns a new OrderDelayField initialized with val
func NewOrderDelay(val int) *OrderDelayField {
	field := &OrderDelayField{}
	field.Value = val
	return field
}

//OrderDelayUnitField is a INT field
type OrderDelayUnitField struct{ quickfix.IntValue }

//Tag returns tag.OrderDelayUnit (1429)
func (f OrderDelayUnitField) Tag() quickfix.Tag { return tag.OrderDelayUnit }

//NewOrderDelayUnit returns a new OrderDelayUnitField initialized with val
func NewOrderDelayUnit(val int) *OrderDelayUnitField {
	field := &OrderDelayUnitField{}
	field.Value = val
	return field
}

//OrderHandlingInstSourceField is a INT field
type OrderHandlingInstSourceField struct{ quickfix.IntValue }

//Tag returns tag.OrderHandlingInstSource (1032)
func (f OrderHandlingInstSourceField) Tag() quickfix.Tag { return tag.OrderHandlingInstSource }

//NewOrderHandlingInstSource returns a new OrderHandlingInstSourceField initialized with val
func NewOrderHandlingInstSource(val int) *OrderHandlingInstSourceField {
	field := &OrderHandlingInstSourceField{}
	field.Value = val
	return field
}

//OrderIDField is a STRING field
type OrderIDField struct{ quickfix.StringValue }

//Tag returns tag.OrderID (37)
func (f OrderIDField) Tag() quickfix.Tag { return tag.OrderID }

//NewOrderID returns a new OrderIDField initialized with val
func NewOrderID(val string) *OrderIDField {
	field := &OrderIDField{}
	field.Value = val
	return field
}

//OrderInputDeviceField is a STRING field
type OrderInputDeviceField struct{ quickfix.StringValue }

//Tag returns tag.OrderInputDevice (821)
func (f OrderInputDeviceField) Tag() quickfix.Tag { return tag.OrderInputDevice }

//NewOrderInputDevice returns a new OrderInputDeviceField initialized with val
func NewOrderInputDevice(val string) *OrderInputDeviceField {
	field := &OrderInputDeviceField{}
	field.Value = val
	return field
}

//OrderPercentField is a PERCENTAGE field
type OrderPercentField struct{ quickfix.PercentageValue }

//Tag returns tag.OrderPercent (516)
func (f OrderPercentField) Tag() quickfix.Tag { return tag.OrderPercent }

//NewOrderPercent returns a new OrderPercentField initialized with val
func NewOrderPercent(val float64) *OrderPercentField {
	field := &OrderPercentField{}
	field.Value = val
	return field
}

//OrderQtyField is a QTY field
type OrderQtyField struct{ quickfix.QtyValue }

//Tag returns tag.OrderQty (38)
func (f OrderQtyField) Tag() quickfix.Tag { return tag.OrderQty }

//NewOrderQty returns a new OrderQtyField initialized with val
func NewOrderQty(val float64) *OrderQtyField {
	field := &OrderQtyField{}
	field.Value = val
	return field
}

//OrderQty2Field is a QTY field
type OrderQty2Field struct{ quickfix.QtyValue }

//Tag returns tag.OrderQty2 (192)
func (f OrderQty2Field) Tag() quickfix.Tag { return tag.OrderQty2 }

//NewOrderQty2 returns a new OrderQty2Field initialized with val
func NewOrderQty2(val float64) *OrderQty2Field {
	field := &OrderQty2Field{}
	field.Value = val
	return field
}

//OrderRestrictionsField is a MULTIPLECHARVALUE field
type OrderRestrictionsField struct{ quickfix.MultipleCharValue }

//Tag returns tag.OrderRestrictions (529)
func (f OrderRestrictionsField) Tag() quickfix.Tag { return tag.OrderRestrictions }

//NewOrderRestrictions returns a new OrderRestrictionsField initialized with val
func NewOrderRestrictions(val string) *OrderRestrictionsField {
	field := &OrderRestrictionsField{}
	field.Value = val
	return field
}

//OrigClOrdIDField is a STRING field
type OrigClOrdIDField struct{ quickfix.StringValue }

//Tag returns tag.OrigClOrdID (41)
func (f OrigClOrdIDField) Tag() quickfix.Tag { return tag.OrigClOrdID }

//NewOrigClOrdID returns a new OrigClOrdIDField initialized with val
func NewOrigClOrdID(val string) *OrigClOrdIDField {
	field := &OrigClOrdIDField{}
	field.Value = val
	return field
}

//OrigCrossIDField is a STRING field
type OrigCrossIDField struct{ quickfix.StringValue }

//Tag returns tag.OrigCrossID (551)
func (f OrigCrossIDField) Tag() quickfix.Tag { return tag.OrigCrossID }

//NewOrigCrossID returns a new OrigCrossIDField initialized with val
func NewOrigCrossID(val string) *OrigCrossIDField {
	field := &OrigCrossIDField{}
	field.Value = val
	return field
}

//OrigCustOrderCapacityField is a INT field
type OrigCustOrderCapacityField struct{ quickfix.IntValue }

//Tag returns tag.OrigCustOrderCapacity (1432)
func (f OrigCustOrderCapacityField) Tag() quickfix.Tag { return tag.OrigCustOrderCapacity }

//NewOrigCustOrderCapacity returns a new OrigCustOrderCapacityField initialized with val
func NewOrigCustOrderCapacity(val int) *OrigCustOrderCapacityField {
	field := &OrigCustOrderCapacityField{}
	field.Value = val
	return field
}

//OrigOrdModTimeField is a UTCTIMESTAMP field
type OrigOrdModTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.OrigOrdModTime (586)
func (f OrigOrdModTimeField) Tag() quickfix.Tag { return tag.OrigOrdModTime }

//OrigPosReqRefIDField is a STRING field
type OrigPosReqRefIDField struct{ quickfix.StringValue }

//Tag returns tag.OrigPosReqRefID (713)
func (f OrigPosReqRefIDField) Tag() quickfix.Tag { return tag.OrigPosReqRefID }

//NewOrigPosReqRefID returns a new OrigPosReqRefIDField initialized with val
func NewOrigPosReqRefID(val string) *OrigPosReqRefIDField {
	field := &OrigPosReqRefIDField{}
	field.Value = val
	return field
}

//OrigSecondaryTradeIDField is a STRING field
type OrigSecondaryTradeIDField struct{ quickfix.StringValue }

//Tag returns tag.OrigSecondaryTradeID (1127)
func (f OrigSecondaryTradeIDField) Tag() quickfix.Tag { return tag.OrigSecondaryTradeID }

//NewOrigSecondaryTradeID returns a new OrigSecondaryTradeIDField initialized with val
func NewOrigSecondaryTradeID(val string) *OrigSecondaryTradeIDField {
	field := &OrigSecondaryTradeIDField{}
	field.Value = val
	return field
}

//OrigSendingTimeField is a UTCTIMESTAMP field
type OrigSendingTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.OrigSendingTime (122)
func (f OrigSendingTimeField) Tag() quickfix.Tag { return tag.OrigSendingTime }

//OrigTimeField is a UTCTIMESTAMP field
type OrigTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.OrigTime (42)
func (f OrigTimeField) Tag() quickfix.Tag { return tag.OrigTime }

//OrigTradeDateField is a LOCALMKTDATE field
type OrigTradeDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.OrigTradeDate (1125)
func (f OrigTradeDateField) Tag() quickfix.Tag { return tag.OrigTradeDate }

//NewOrigTradeDate returns a new OrigTradeDateField initialized with val
func NewOrigTradeDate(val string) *OrigTradeDateField {
	field := &OrigTradeDateField{}
	field.Value = val
	return field
}

//OrigTradeHandlingInstrField is a CHAR field
type OrigTradeHandlingInstrField struct{ quickfix.CharValue }

//Tag returns tag.OrigTradeHandlingInstr (1124)
func (f OrigTradeHandlingInstrField) Tag() quickfix.Tag { return tag.OrigTradeHandlingInstr }

//NewOrigTradeHandlingInstr returns a new OrigTradeHandlingInstrField initialized with val
func NewOrigTradeHandlingInstr(val string) *OrigTradeHandlingInstrField {
	field := &OrigTradeHandlingInstrField{}
	field.Value = val
	return field
}

//OrigTradeIDField is a STRING field
type OrigTradeIDField struct{ quickfix.StringValue }

//Tag returns tag.OrigTradeID (1126)
func (f OrigTradeIDField) Tag() quickfix.Tag { return tag.OrigTradeID }

//NewOrigTradeID returns a new OrigTradeIDField initialized with val
func NewOrigTradeID(val string) *OrigTradeIDField {
	field := &OrigTradeIDField{}
	field.Value = val
	return field
}

//OriginalNotionalPercentageOutstandingField is a PERCENTAGE field
type OriginalNotionalPercentageOutstandingField struct{ quickfix.PercentageValue }

//Tag returns tag.OriginalNotionalPercentageOutstanding (1452)
func (f OriginalNotionalPercentageOutstandingField) Tag() quickfix.Tag {
	return tag.OriginalNotionalPercentageOutstanding
}

//NewOriginalNotionalPercentageOutstanding returns a new OriginalNotionalPercentageOutstandingField initialized with val
func NewOriginalNotionalPercentageOutstanding(val float64) *OriginalNotionalPercentageOutstandingField {
	field := &OriginalNotionalPercentageOutstandingField{}
	field.Value = val
	return field
}

//OutMainCntryUIndexField is a AMT field
type OutMainCntryUIndexField struct{ quickfix.AmtValue }

//Tag returns tag.OutMainCntryUIndex (412)
func (f OutMainCntryUIndexField) Tag() quickfix.Tag { return tag.OutMainCntryUIndex }

//NewOutMainCntryUIndex returns a new OutMainCntryUIndexField initialized with val
func NewOutMainCntryUIndex(val float64) *OutMainCntryUIndexField {
	field := &OutMainCntryUIndexField{}
	field.Value = val
	return field
}

//OutsideIndexPctField is a PERCENTAGE field
type OutsideIndexPctField struct{ quickfix.PercentageValue }

//Tag returns tag.OutsideIndexPct (407)
func (f OutsideIndexPctField) Tag() quickfix.Tag { return tag.OutsideIndexPct }

//NewOutsideIndexPct returns a new OutsideIndexPctField initialized with val
func NewOutsideIndexPct(val float64) *OutsideIndexPctField {
	field := &OutsideIndexPctField{}
	field.Value = val
	return field
}

//OwnerTypeField is a INT field
type OwnerTypeField struct{ quickfix.IntValue }

//Tag returns tag.OwnerType (522)
func (f OwnerTypeField) Tag() quickfix.Tag { return tag.OwnerType }

//NewOwnerType returns a new OwnerTypeField initialized with val
func NewOwnerType(val int) *OwnerTypeField {
	field := &OwnerTypeField{}
	field.Value = val
	return field
}

//OwnershipTypeField is a CHAR field
type OwnershipTypeField struct{ quickfix.CharValue }

//Tag returns tag.OwnershipType (517)
func (f OwnershipTypeField) Tag() quickfix.Tag { return tag.OwnershipType }

//NewOwnershipType returns a new OwnershipTypeField initialized with val
func NewOwnershipType(val string) *OwnershipTypeField {
	field := &OwnershipTypeField{}
	field.Value = val
	return field
}

//ParentMktSegmIDField is a STRING field
type ParentMktSegmIDField struct{ quickfix.StringValue }

//Tag returns tag.ParentMktSegmID (1325)
func (f ParentMktSegmIDField) Tag() quickfix.Tag { return tag.ParentMktSegmID }

//NewParentMktSegmID returns a new ParentMktSegmIDField initialized with val
func NewParentMktSegmID(val string) *ParentMktSegmIDField {
	field := &ParentMktSegmIDField{}
	field.Value = val
	return field
}

//ParticipationRateField is a PERCENTAGE field
type ParticipationRateField struct{ quickfix.PercentageValue }

//Tag returns tag.ParticipationRate (849)
func (f ParticipationRateField) Tag() quickfix.Tag { return tag.ParticipationRate }

//NewParticipationRate returns a new ParticipationRateField initialized with val
func NewParticipationRate(val float64) *ParticipationRateField {
	field := &ParticipationRateField{}
	field.Value = val
	return field
}

//PartyAltIDField is a STRING field
type PartyAltIDField struct{ quickfix.StringValue }

//Tag returns tag.PartyAltID (1517)
func (f PartyAltIDField) Tag() quickfix.Tag { return tag.PartyAltID }

//NewPartyAltID returns a new PartyAltIDField initialized with val
func NewPartyAltID(val string) *PartyAltIDField {
	field := &PartyAltIDField{}
	field.Value = val
	return field
}

//PartyAltIDSourceField is a CHAR field
type PartyAltIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.PartyAltIDSource (1518)
func (f PartyAltIDSourceField) Tag() quickfix.Tag { return tag.PartyAltIDSource }

//NewPartyAltIDSource returns a new PartyAltIDSourceField initialized with val
func NewPartyAltIDSource(val string) *PartyAltIDSourceField {
	field := &PartyAltIDSourceField{}
	field.Value = val
	return field
}

//PartyAltSubIDField is a STRING field
type PartyAltSubIDField struct{ quickfix.StringValue }

//Tag returns tag.PartyAltSubID (1520)
func (f PartyAltSubIDField) Tag() quickfix.Tag { return tag.PartyAltSubID }

//NewPartyAltSubID returns a new PartyAltSubIDField initialized with val
func NewPartyAltSubID(val string) *PartyAltSubIDField {
	field := &PartyAltSubIDField{}
	field.Value = val
	return field
}

//PartyAltSubIDTypeField is a INT field
type PartyAltSubIDTypeField struct{ quickfix.IntValue }

//Tag returns tag.PartyAltSubIDType (1521)
func (f PartyAltSubIDTypeField) Tag() quickfix.Tag { return tag.PartyAltSubIDType }

//NewPartyAltSubIDType returns a new PartyAltSubIDTypeField initialized with val
func NewPartyAltSubIDType(val int) *PartyAltSubIDTypeField {
	field := &PartyAltSubIDTypeField{}
	field.Value = val
	return field
}

//PartyDetailsListReportIDField is a STRING field
type PartyDetailsListReportIDField struct{ quickfix.StringValue }

//Tag returns tag.PartyDetailsListReportID (1510)
func (f PartyDetailsListReportIDField) Tag() quickfix.Tag { return tag.PartyDetailsListReportID }

//NewPartyDetailsListReportID returns a new PartyDetailsListReportIDField initialized with val
func NewPartyDetailsListReportID(val string) *PartyDetailsListReportIDField {
	field := &PartyDetailsListReportIDField{}
	field.Value = val
	return field
}

//PartyDetailsListRequestIDField is a STRING field
type PartyDetailsListRequestIDField struct{ quickfix.StringValue }

//Tag returns tag.PartyDetailsListRequestID (1505)
func (f PartyDetailsListRequestIDField) Tag() quickfix.Tag { return tag.PartyDetailsListRequestID }

//NewPartyDetailsListRequestID returns a new PartyDetailsListRequestIDField initialized with val
func NewPartyDetailsListRequestID(val string) *PartyDetailsListRequestIDField {
	field := &PartyDetailsListRequestIDField{}
	field.Value = val
	return field
}

//PartyDetailsRequestResultField is a INT field
type PartyDetailsRequestResultField struct{ quickfix.IntValue }

//Tag returns tag.PartyDetailsRequestResult (1511)
func (f PartyDetailsRequestResultField) Tag() quickfix.Tag { return tag.PartyDetailsRequestResult }

//NewPartyDetailsRequestResult returns a new PartyDetailsRequestResultField initialized with val
func NewPartyDetailsRequestResult(val int) *PartyDetailsRequestResultField {
	field := &PartyDetailsRequestResultField{}
	field.Value = val
	return field
}

//PartyIDField is a STRING field
type PartyIDField struct{ quickfix.StringValue }

//Tag returns tag.PartyID (448)
func (f PartyIDField) Tag() quickfix.Tag { return tag.PartyID }

//NewPartyID returns a new PartyIDField initialized with val
func NewPartyID(val string) *PartyIDField {
	field := &PartyIDField{}
	field.Value = val
	return field
}

//PartyIDSourceField is a CHAR field
type PartyIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.PartyIDSource (447)
func (f PartyIDSourceField) Tag() quickfix.Tag { return tag.PartyIDSource }

//NewPartyIDSource returns a new PartyIDSourceField initialized with val
func NewPartyIDSource(val string) *PartyIDSourceField {
	field := &PartyIDSourceField{}
	field.Value = val
	return field
}

//PartyListResponseTypeField is a INT field
type PartyListResponseTypeField struct{ quickfix.IntValue }

//Tag returns tag.PartyListResponseType (1507)
func (f PartyListResponseTypeField) Tag() quickfix.Tag { return tag.PartyListResponseType }

//NewPartyListResponseType returns a new PartyListResponseTypeField initialized with val
func NewPartyListResponseType(val int) *PartyListResponseTypeField {
	field := &PartyListResponseTypeField{}
	field.Value = val
	return field
}

//PartyRelationshipField is a INT field
type PartyRelationshipField struct{ quickfix.IntValue }

//Tag returns tag.PartyRelationship (1515)
func (f PartyRelationshipField) Tag() quickfix.Tag { return tag.PartyRelationship }

//NewPartyRelationship returns a new PartyRelationshipField initialized with val
func NewPartyRelationship(val int) *PartyRelationshipField {
	field := &PartyRelationshipField{}
	field.Value = val
	return field
}

//PartyRoleField is a INT field
type PartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.PartyRole (452)
func (f PartyRoleField) Tag() quickfix.Tag { return tag.PartyRole }

//NewPartyRole returns a new PartyRoleField initialized with val
func NewPartyRole(val int) *PartyRoleField {
	field := &PartyRoleField{}
	field.Value = val
	return field
}

//PartySubIDField is a STRING field
type PartySubIDField struct{ quickfix.StringValue }

//Tag returns tag.PartySubID (523)
func (f PartySubIDField) Tag() quickfix.Tag { return tag.PartySubID }

//NewPartySubID returns a new PartySubIDField initialized with val
func NewPartySubID(val string) *PartySubIDField {
	field := &PartySubIDField{}
	field.Value = val
	return field
}

//PartySubIDTypeField is a INT field
type PartySubIDTypeField struct{ quickfix.IntValue }

//Tag returns tag.PartySubIDType (803)
func (f PartySubIDTypeField) Tag() quickfix.Tag { return tag.PartySubIDType }

//NewPartySubIDType returns a new PartySubIDTypeField initialized with val
func NewPartySubIDType(val int) *PartySubIDTypeField {
	field := &PartySubIDTypeField{}
	field.Value = val
	return field
}

//PasswordField is a STRING field
type PasswordField struct{ quickfix.StringValue }

//Tag returns tag.Password (554)
func (f PasswordField) Tag() quickfix.Tag { return tag.Password }

//NewPassword returns a new PasswordField initialized with val
func NewPassword(val string) *PasswordField {
	field := &PasswordField{}
	field.Value = val
	return field
}

//PaymentDateField is a LOCALMKTDATE field
type PaymentDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.PaymentDate (504)
func (f PaymentDateField) Tag() quickfix.Tag { return tag.PaymentDate }

//NewPaymentDate returns a new PaymentDateField initialized with val
func NewPaymentDate(val string) *PaymentDateField {
	field := &PaymentDateField{}
	field.Value = val
	return field
}

//PaymentMethodField is a INT field
type PaymentMethodField struct{ quickfix.IntValue }

//Tag returns tag.PaymentMethod (492)
func (f PaymentMethodField) Tag() quickfix.Tag { return tag.PaymentMethod }

//NewPaymentMethod returns a new PaymentMethodField initialized with val
func NewPaymentMethod(val int) *PaymentMethodField {
	field := &PaymentMethodField{}
	field.Value = val
	return field
}

//PaymentRefField is a STRING field
type PaymentRefField struct{ quickfix.StringValue }

//Tag returns tag.PaymentRef (476)
func (f PaymentRefField) Tag() quickfix.Tag { return tag.PaymentRef }

//NewPaymentRef returns a new PaymentRefField initialized with val
func NewPaymentRef(val string) *PaymentRefField {
	field := &PaymentRefField{}
	field.Value = val
	return field
}

//PaymentRemitterIDField is a STRING field
type PaymentRemitterIDField struct{ quickfix.StringValue }

//Tag returns tag.PaymentRemitterID (505)
func (f PaymentRemitterIDField) Tag() quickfix.Tag { return tag.PaymentRemitterID }

//NewPaymentRemitterID returns a new PaymentRemitterIDField initialized with val
func NewPaymentRemitterID(val string) *PaymentRemitterIDField {
	field := &PaymentRemitterIDField{}
	field.Value = val
	return field
}

//PctAtRiskField is a PERCENTAGE field
type PctAtRiskField struct{ quickfix.PercentageValue }

//Tag returns tag.PctAtRisk (869)
func (f PctAtRiskField) Tag() quickfix.Tag { return tag.PctAtRisk }

//NewPctAtRisk returns a new PctAtRiskField initialized with val
func NewPctAtRisk(val float64) *PctAtRiskField {
	field := &PctAtRiskField{}
	field.Value = val
	return field
}

//PegDifferenceField is a PRICEOFFSET field
type PegDifferenceField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.PegDifference (211)
func (f PegDifferenceField) Tag() quickfix.Tag { return tag.PegDifference }

//NewPegDifference returns a new PegDifferenceField initialized with val
func NewPegDifference(val float64) *PegDifferenceField {
	field := &PegDifferenceField{}
	field.Value = val
	return field
}

//PegLimitTypeField is a INT field
type PegLimitTypeField struct{ quickfix.IntValue }

//Tag returns tag.PegLimitType (837)
func (f PegLimitTypeField) Tag() quickfix.Tag { return tag.PegLimitType }

//NewPegLimitType returns a new PegLimitTypeField initialized with val
func NewPegLimitType(val int) *PegLimitTypeField {
	field := &PegLimitTypeField{}
	field.Value = val
	return field
}

//PegMoveTypeField is a INT field
type PegMoveTypeField struct{ quickfix.IntValue }

//Tag returns tag.PegMoveType (835)
func (f PegMoveTypeField) Tag() quickfix.Tag { return tag.PegMoveType }

//NewPegMoveType returns a new PegMoveTypeField initialized with val
func NewPegMoveType(val int) *PegMoveTypeField {
	field := &PegMoveTypeField{}
	field.Value = val
	return field
}

//PegOffsetTypeField is a INT field
type PegOffsetTypeField struct{ quickfix.IntValue }

//Tag returns tag.PegOffsetType (836)
func (f PegOffsetTypeField) Tag() quickfix.Tag { return tag.PegOffsetType }

//NewPegOffsetType returns a new PegOffsetTypeField initialized with val
func NewPegOffsetType(val int) *PegOffsetTypeField {
	field := &PegOffsetTypeField{}
	field.Value = val
	return field
}

//PegOffsetValueField is a FLOAT field
type PegOffsetValueField struct{ quickfix.FloatValue }

//Tag returns tag.PegOffsetValue (211)
func (f PegOffsetValueField) Tag() quickfix.Tag { return tag.PegOffsetValue }

//NewPegOffsetValue returns a new PegOffsetValueField initialized with val
func NewPegOffsetValue(val float64) *PegOffsetValueField {
	field := &PegOffsetValueField{}
	field.Value = val
	return field
}

//PegPriceTypeField is a INT field
type PegPriceTypeField struct{ quickfix.IntValue }

//Tag returns tag.PegPriceType (1094)
func (f PegPriceTypeField) Tag() quickfix.Tag { return tag.PegPriceType }

//NewPegPriceType returns a new PegPriceTypeField initialized with val
func NewPegPriceType(val int) *PegPriceTypeField {
	field := &PegPriceTypeField{}
	field.Value = val
	return field
}

//PegRoundDirectionField is a INT field
type PegRoundDirectionField struct{ quickfix.IntValue }

//Tag returns tag.PegRoundDirection (838)
func (f PegRoundDirectionField) Tag() quickfix.Tag { return tag.PegRoundDirection }

//NewPegRoundDirection returns a new PegRoundDirectionField initialized with val
func NewPegRoundDirection(val int) *PegRoundDirectionField {
	field := &PegRoundDirectionField{}
	field.Value = val
	return field
}

//PegScopeField is a INT field
type PegScopeField struct{ quickfix.IntValue }

//Tag returns tag.PegScope (840)
func (f PegScopeField) Tag() quickfix.Tag { return tag.PegScope }

//NewPegScope returns a new PegScopeField initialized with val
func NewPegScope(val int) *PegScopeField {
	field := &PegScopeField{}
	field.Value = val
	return field
}

//PegSecurityDescField is a STRING field
type PegSecurityDescField struct{ quickfix.StringValue }

//Tag returns tag.PegSecurityDesc (1099)
func (f PegSecurityDescField) Tag() quickfix.Tag { return tag.PegSecurityDesc }

//NewPegSecurityDesc returns a new PegSecurityDescField initialized with val
func NewPegSecurityDesc(val string) *PegSecurityDescField {
	field := &PegSecurityDescField{}
	field.Value = val
	return field
}

//PegSecurityIDField is a STRING field
type PegSecurityIDField struct{ quickfix.StringValue }

//Tag returns tag.PegSecurityID (1097)
func (f PegSecurityIDField) Tag() quickfix.Tag { return tag.PegSecurityID }

//NewPegSecurityID returns a new PegSecurityIDField initialized with val
func NewPegSecurityID(val string) *PegSecurityIDField {
	field := &PegSecurityIDField{}
	field.Value = val
	return field
}

//PegSecurityIDSourceField is a STRING field
type PegSecurityIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.PegSecurityIDSource (1096)
func (f PegSecurityIDSourceField) Tag() quickfix.Tag { return tag.PegSecurityIDSource }

//NewPegSecurityIDSource returns a new PegSecurityIDSourceField initialized with val
func NewPegSecurityIDSource(val string) *PegSecurityIDSourceField {
	field := &PegSecurityIDSourceField{}
	field.Value = val
	return field
}

//PegSymbolField is a STRING field
type PegSymbolField struct{ quickfix.StringValue }

//Tag returns tag.PegSymbol (1098)
func (f PegSymbolField) Tag() quickfix.Tag { return tag.PegSymbol }

//NewPegSymbol returns a new PegSymbolField initialized with val
func NewPegSymbol(val string) *PegSymbolField {
	field := &PegSymbolField{}
	field.Value = val
	return field
}

//PeggedPriceField is a PRICE field
type PeggedPriceField struct{ quickfix.PriceValue }

//Tag returns tag.PeggedPrice (839)
func (f PeggedPriceField) Tag() quickfix.Tag { return tag.PeggedPrice }

//NewPeggedPrice returns a new PeggedPriceField initialized with val
func NewPeggedPrice(val float64) *PeggedPriceField {
	field := &PeggedPriceField{}
	field.Value = val
	return field
}

//PeggedRefPriceField is a PRICE field
type PeggedRefPriceField struct{ quickfix.PriceValue }

//Tag returns tag.PeggedRefPrice (1095)
func (f PeggedRefPriceField) Tag() quickfix.Tag { return tag.PeggedRefPrice }

//NewPeggedRefPrice returns a new PeggedRefPriceField initialized with val
func NewPeggedRefPrice(val float64) *PeggedRefPriceField {
	field := &PeggedRefPriceField{}
	field.Value = val
	return field
}

//PoolField is a STRING field
type PoolField struct{ quickfix.StringValue }

//Tag returns tag.Pool (691)
func (f PoolField) Tag() quickfix.Tag { return tag.Pool }

//NewPool returns a new PoolField initialized with val
func NewPool(val string) *PoolField {
	field := &PoolField{}
	field.Value = val
	return field
}

//PosAmtField is a AMT field
type PosAmtField struct{ quickfix.AmtValue }

//Tag returns tag.PosAmt (708)
func (f PosAmtField) Tag() quickfix.Tag { return tag.PosAmt }

//NewPosAmt returns a new PosAmtField initialized with val
func NewPosAmt(val float64) *PosAmtField {
	field := &PosAmtField{}
	field.Value = val
	return field
}

//PosAmtTypeField is a STRING field
type PosAmtTypeField struct{ quickfix.StringValue }

//Tag returns tag.PosAmtType (707)
func (f PosAmtTypeField) Tag() quickfix.Tag { return tag.PosAmtType }

//NewPosAmtType returns a new PosAmtTypeField initialized with val
func NewPosAmtType(val string) *PosAmtTypeField {
	field := &PosAmtTypeField{}
	field.Value = val
	return field
}

//PosMaintActionField is a INT field
type PosMaintActionField struct{ quickfix.IntValue }

//Tag returns tag.PosMaintAction (712)
func (f PosMaintActionField) Tag() quickfix.Tag { return tag.PosMaintAction }

//NewPosMaintAction returns a new PosMaintActionField initialized with val
func NewPosMaintAction(val int) *PosMaintActionField {
	field := &PosMaintActionField{}
	field.Value = val
	return field
}

//PosMaintResultField is a INT field
type PosMaintResultField struct{ quickfix.IntValue }

//Tag returns tag.PosMaintResult (723)
func (f PosMaintResultField) Tag() quickfix.Tag { return tag.PosMaintResult }

//NewPosMaintResult returns a new PosMaintResultField initialized with val
func NewPosMaintResult(val int) *PosMaintResultField {
	field := &PosMaintResultField{}
	field.Value = val
	return field
}

//PosMaintRptIDField is a STRING field
type PosMaintRptIDField struct{ quickfix.StringValue }

//Tag returns tag.PosMaintRptID (721)
func (f PosMaintRptIDField) Tag() quickfix.Tag { return tag.PosMaintRptID }

//NewPosMaintRptID returns a new PosMaintRptIDField initialized with val
func NewPosMaintRptID(val string) *PosMaintRptIDField {
	field := &PosMaintRptIDField{}
	field.Value = val
	return field
}

//PosMaintRptRefIDField is a STRING field
type PosMaintRptRefIDField struct{ quickfix.StringValue }

//Tag returns tag.PosMaintRptRefID (714)
func (f PosMaintRptRefIDField) Tag() quickfix.Tag { return tag.PosMaintRptRefID }

//NewPosMaintRptRefID returns a new PosMaintRptRefIDField initialized with val
func NewPosMaintRptRefID(val string) *PosMaintRptRefIDField {
	field := &PosMaintRptRefIDField{}
	field.Value = val
	return field
}

//PosMaintStatusField is a INT field
type PosMaintStatusField struct{ quickfix.IntValue }

//Tag returns tag.PosMaintStatus (722)
func (f PosMaintStatusField) Tag() quickfix.Tag { return tag.PosMaintStatus }

//NewPosMaintStatus returns a new PosMaintStatusField initialized with val
func NewPosMaintStatus(val int) *PosMaintStatusField {
	field := &PosMaintStatusField{}
	field.Value = val
	return field
}

//PosQtyStatusField is a INT field
type PosQtyStatusField struct{ quickfix.IntValue }

//Tag returns tag.PosQtyStatus (706)
func (f PosQtyStatusField) Tag() quickfix.Tag { return tag.PosQtyStatus }

//NewPosQtyStatus returns a new PosQtyStatusField initialized with val
func NewPosQtyStatus(val int) *PosQtyStatusField {
	field := &PosQtyStatusField{}
	field.Value = val
	return field
}

//PosReqIDField is a STRING field
type PosReqIDField struct{ quickfix.StringValue }

//Tag returns tag.PosReqID (710)
func (f PosReqIDField) Tag() quickfix.Tag { return tag.PosReqID }

//NewPosReqID returns a new PosReqIDField initialized with val
func NewPosReqID(val string) *PosReqIDField {
	field := &PosReqIDField{}
	field.Value = val
	return field
}

//PosReqResultField is a INT field
type PosReqResultField struct{ quickfix.IntValue }

//Tag returns tag.PosReqResult (728)
func (f PosReqResultField) Tag() quickfix.Tag { return tag.PosReqResult }

//NewPosReqResult returns a new PosReqResultField initialized with val
func NewPosReqResult(val int) *PosReqResultField {
	field := &PosReqResultField{}
	field.Value = val
	return field
}

//PosReqStatusField is a INT field
type PosReqStatusField struct{ quickfix.IntValue }

//Tag returns tag.PosReqStatus (729)
func (f PosReqStatusField) Tag() quickfix.Tag { return tag.PosReqStatus }

//NewPosReqStatus returns a new PosReqStatusField initialized with val
func NewPosReqStatus(val int) *PosReqStatusField {
	field := &PosReqStatusField{}
	field.Value = val
	return field
}

//PosReqTypeField is a INT field
type PosReqTypeField struct{ quickfix.IntValue }

//Tag returns tag.PosReqType (724)
func (f PosReqTypeField) Tag() quickfix.Tag { return tag.PosReqType }

//NewPosReqType returns a new PosReqTypeField initialized with val
func NewPosReqType(val int) *PosReqTypeField {
	field := &PosReqTypeField{}
	field.Value = val
	return field
}

//PosTransTypeField is a INT field
type PosTransTypeField struct{ quickfix.IntValue }

//Tag returns tag.PosTransType (709)
func (f PosTransTypeField) Tag() quickfix.Tag { return tag.PosTransType }

//NewPosTransType returns a new PosTransTypeField initialized with val
func NewPosTransType(val int) *PosTransTypeField {
	field := &PosTransTypeField{}
	field.Value = val
	return field
}

//PosTypeField is a STRING field
type PosTypeField struct{ quickfix.StringValue }

//Tag returns tag.PosType (703)
func (f PosTypeField) Tag() quickfix.Tag { return tag.PosType }

//NewPosType returns a new PosTypeField initialized with val
func NewPosType(val string) *PosTypeField {
	field := &PosTypeField{}
	field.Value = val
	return field
}

//PositionCurrencyField is a STRING field
type PositionCurrencyField struct{ quickfix.StringValue }

//Tag returns tag.PositionCurrency (1055)
func (f PositionCurrencyField) Tag() quickfix.Tag { return tag.PositionCurrency }

//NewPositionCurrency returns a new PositionCurrencyField initialized with val
func NewPositionCurrency(val string) *PositionCurrencyField {
	field := &PositionCurrencyField{}
	field.Value = val
	return field
}

//PositionEffectField is a CHAR field
type PositionEffectField struct{ quickfix.CharValue }

//Tag returns tag.PositionEffect (77)
func (f PositionEffectField) Tag() quickfix.Tag { return tag.PositionEffect }

//NewPositionEffect returns a new PositionEffectField initialized with val
func NewPositionEffect(val string) *PositionEffectField {
	field := &PositionEffectField{}
	field.Value = val
	return field
}

//PositionLimitField is a INT field
type PositionLimitField struct{ quickfix.IntValue }

//Tag returns tag.PositionLimit (970)
func (f PositionLimitField) Tag() quickfix.Tag { return tag.PositionLimit }

//NewPositionLimit returns a new PositionLimitField initialized with val
func NewPositionLimit(val int) *PositionLimitField {
	field := &PositionLimitField{}
	field.Value = val
	return field
}

//PossDupFlagField is a BOOLEAN field
type PossDupFlagField struct{ quickfix.BooleanValue }

//Tag returns tag.PossDupFlag (43)
func (f PossDupFlagField) Tag() quickfix.Tag { return tag.PossDupFlag }

//NewPossDupFlag returns a new PossDupFlagField initialized with val
func NewPossDupFlag(val bool) *PossDupFlagField {
	field := &PossDupFlagField{}
	field.Value = val
	return field
}

//PossResendField is a BOOLEAN field
type PossResendField struct{ quickfix.BooleanValue }

//Tag returns tag.PossResend (97)
func (f PossResendField) Tag() quickfix.Tag { return tag.PossResend }

//NewPossResend returns a new PossResendField initialized with val
func NewPossResend(val bool) *PossResendField {
	field := &PossResendField{}
	field.Value = val
	return field
}

//PreTradeAnonymityField is a BOOLEAN field
type PreTradeAnonymityField struct{ quickfix.BooleanValue }

//Tag returns tag.PreTradeAnonymity (1091)
func (f PreTradeAnonymityField) Tag() quickfix.Tag { return tag.PreTradeAnonymity }

//NewPreTradeAnonymity returns a new PreTradeAnonymityField initialized with val
func NewPreTradeAnonymity(val bool) *PreTradeAnonymityField {
	field := &PreTradeAnonymityField{}
	field.Value = val
	return field
}

//PreallocMethodField is a CHAR field
type PreallocMethodField struct{ quickfix.CharValue }

//Tag returns tag.PreallocMethod (591)
func (f PreallocMethodField) Tag() quickfix.Tag { return tag.PreallocMethod }

//NewPreallocMethod returns a new PreallocMethodField initialized with val
func NewPreallocMethod(val string) *PreallocMethodField {
	field := &PreallocMethodField{}
	field.Value = val
	return field
}

//PrevClosePxField is a PRICE field
type PrevClosePxField struct{ quickfix.PriceValue }

//Tag returns tag.PrevClosePx (140)
func (f PrevClosePxField) Tag() quickfix.Tag { return tag.PrevClosePx }

//NewPrevClosePx returns a new PrevClosePxField initialized with val
func NewPrevClosePx(val float64) *PrevClosePxField {
	field := &PrevClosePxField{}
	field.Value = val
	return field
}

//PreviouslyReportedField is a BOOLEAN field
type PreviouslyReportedField struct{ quickfix.BooleanValue }

//Tag returns tag.PreviouslyReported (570)
func (f PreviouslyReportedField) Tag() quickfix.Tag { return tag.PreviouslyReported }

//NewPreviouslyReported returns a new PreviouslyReportedField initialized with val
func NewPreviouslyReported(val bool) *PreviouslyReportedField {
	field := &PreviouslyReportedField{}
	field.Value = val
	return field
}

//PriceField is a PRICE field
type PriceField struct{ quickfix.PriceValue }

//Tag returns tag.Price (44)
func (f PriceField) Tag() quickfix.Tag { return tag.Price }

//NewPrice returns a new PriceField initialized with val
func NewPrice(val float64) *PriceField {
	field := &PriceField{}
	field.Value = val
	return field
}

//Price2Field is a PRICE field
type Price2Field struct{ quickfix.PriceValue }

//Tag returns tag.Price2 (640)
func (f Price2Field) Tag() quickfix.Tag { return tag.Price2 }

//NewPrice2 returns a new Price2Field initialized with val
func NewPrice2(val float64) *Price2Field {
	field := &Price2Field{}
	field.Value = val
	return field
}

//PriceDeltaField is a FLOAT field
type PriceDeltaField struct{ quickfix.FloatValue }

//Tag returns tag.PriceDelta (811)
func (f PriceDeltaField) Tag() quickfix.Tag { return tag.PriceDelta }

//NewPriceDelta returns a new PriceDeltaField initialized with val
func NewPriceDelta(val float64) *PriceDeltaField {
	field := &PriceDeltaField{}
	field.Value = val
	return field
}

//PriceImprovementField is a PRICEOFFSET field
type PriceImprovementField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.PriceImprovement (639)
func (f PriceImprovementField) Tag() quickfix.Tag { return tag.PriceImprovement }

//NewPriceImprovement returns a new PriceImprovementField initialized with val
func NewPriceImprovement(val float64) *PriceImprovementField {
	field := &PriceImprovementField{}
	field.Value = val
	return field
}

//PriceLimitTypeField is a INT field
type PriceLimitTypeField struct{ quickfix.IntValue }

//Tag returns tag.PriceLimitType (1306)
func (f PriceLimitTypeField) Tag() quickfix.Tag { return tag.PriceLimitType }

//NewPriceLimitType returns a new PriceLimitTypeField initialized with val
func NewPriceLimitType(val int) *PriceLimitTypeField {
	field := &PriceLimitTypeField{}
	field.Value = val
	return field
}

//PriceProtectionScopeField is a CHAR field
type PriceProtectionScopeField struct{ quickfix.CharValue }

//Tag returns tag.PriceProtectionScope (1092)
func (f PriceProtectionScopeField) Tag() quickfix.Tag { return tag.PriceProtectionScope }

//NewPriceProtectionScope returns a new PriceProtectionScopeField initialized with val
func NewPriceProtectionScope(val string) *PriceProtectionScopeField {
	field := &PriceProtectionScopeField{}
	field.Value = val
	return field
}

//PriceQuoteMethodField is a STRING field
type PriceQuoteMethodField struct{ quickfix.StringValue }

//Tag returns tag.PriceQuoteMethod (1196)
func (f PriceQuoteMethodField) Tag() quickfix.Tag { return tag.PriceQuoteMethod }

//NewPriceQuoteMethod returns a new PriceQuoteMethodField initialized with val
func NewPriceQuoteMethod(val string) *PriceQuoteMethodField {
	field := &PriceQuoteMethodField{}
	field.Value = val
	return field
}

//PriceTypeField is a INT field
type PriceTypeField struct{ quickfix.IntValue }

//Tag returns tag.PriceType (423)
func (f PriceTypeField) Tag() quickfix.Tag { return tag.PriceType }

//NewPriceType returns a new PriceTypeField initialized with val
func NewPriceType(val int) *PriceTypeField {
	field := &PriceTypeField{}
	field.Value = val
	return field
}

//PriceUnitOfMeasureField is a STRING field
type PriceUnitOfMeasureField struct{ quickfix.StringValue }

//Tag returns tag.PriceUnitOfMeasure (1191)
func (f PriceUnitOfMeasureField) Tag() quickfix.Tag { return tag.PriceUnitOfMeasure }

//NewPriceUnitOfMeasure returns a new PriceUnitOfMeasureField initialized with val
func NewPriceUnitOfMeasure(val string) *PriceUnitOfMeasureField {
	field := &PriceUnitOfMeasureField{}
	field.Value = val
	return field
}

//PriceUnitOfMeasureQtyField is a QTY field
type PriceUnitOfMeasureQtyField struct{ quickfix.QtyValue }

//Tag returns tag.PriceUnitOfMeasureQty (1192)
func (f PriceUnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.PriceUnitOfMeasureQty }

//NewPriceUnitOfMeasureQty returns a new PriceUnitOfMeasureQtyField initialized with val
func NewPriceUnitOfMeasureQty(val float64) *PriceUnitOfMeasureQtyField {
	field := &PriceUnitOfMeasureQtyField{}
	field.Value = val
	return field
}

//PriorSettlPriceField is a PRICE field
type PriorSettlPriceField struct{ quickfix.PriceValue }

//Tag returns tag.PriorSettlPrice (734)
func (f PriorSettlPriceField) Tag() quickfix.Tag { return tag.PriorSettlPrice }

//NewPriorSettlPrice returns a new PriorSettlPriceField initialized with val
func NewPriorSettlPrice(val float64) *PriorSettlPriceField {
	field := &PriorSettlPriceField{}
	field.Value = val
	return field
}

//PriorSpreadIndicatorField is a BOOLEAN field
type PriorSpreadIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.PriorSpreadIndicator (720)
func (f PriorSpreadIndicatorField) Tag() quickfix.Tag { return tag.PriorSpreadIndicator }

//NewPriorSpreadIndicator returns a new PriorSpreadIndicatorField initialized with val
func NewPriorSpreadIndicator(val bool) *PriorSpreadIndicatorField {
	field := &PriorSpreadIndicatorField{}
	field.Value = val
	return field
}

//PriorityIndicatorField is a INT field
type PriorityIndicatorField struct{ quickfix.IntValue }

//Tag returns tag.PriorityIndicator (638)
func (f PriorityIndicatorField) Tag() quickfix.Tag { return tag.PriorityIndicator }

//NewPriorityIndicator returns a new PriorityIndicatorField initialized with val
func NewPriorityIndicator(val int) *PriorityIndicatorField {
	field := &PriorityIndicatorField{}
	field.Value = val
	return field
}

//PrivateQuoteField is a BOOLEAN field
type PrivateQuoteField struct{ quickfix.BooleanValue }

//Tag returns tag.PrivateQuote (1171)
func (f PrivateQuoteField) Tag() quickfix.Tag { return tag.PrivateQuote }

//NewPrivateQuote returns a new PrivateQuoteField initialized with val
func NewPrivateQuote(val bool) *PrivateQuoteField {
	field := &PrivateQuoteField{}
	field.Value = val
	return field
}

//ProcessCodeField is a CHAR field
type ProcessCodeField struct{ quickfix.CharValue }

//Tag returns tag.ProcessCode (81)
func (f ProcessCodeField) Tag() quickfix.Tag { return tag.ProcessCode }

//NewProcessCode returns a new ProcessCodeField initialized with val
func NewProcessCode(val string) *ProcessCodeField {
	field := &ProcessCodeField{}
	field.Value = val
	return field
}

//ProductField is a INT field
type ProductField struct{ quickfix.IntValue }

//Tag returns tag.Product (460)
func (f ProductField) Tag() quickfix.Tag { return tag.Product }

//NewProduct returns a new ProductField initialized with val
func NewProduct(val int) *ProductField {
	field := &ProductField{}
	field.Value = val
	return field
}

//ProductComplexField is a STRING field
type ProductComplexField struct{ quickfix.StringValue }

//Tag returns tag.ProductComplex (1227)
func (f ProductComplexField) Tag() quickfix.Tag { return tag.ProductComplex }

//NewProductComplex returns a new ProductComplexField initialized with val
func NewProductComplex(val string) *ProductComplexField {
	field := &ProductComplexField{}
	field.Value = val
	return field
}

//ProgPeriodIntervalField is a INT field
type ProgPeriodIntervalField struct{ quickfix.IntValue }

//Tag returns tag.ProgPeriodInterval (415)
func (f ProgPeriodIntervalField) Tag() quickfix.Tag { return tag.ProgPeriodInterval }

//NewProgPeriodInterval returns a new ProgPeriodIntervalField initialized with val
func NewProgPeriodInterval(val int) *ProgPeriodIntervalField {
	field := &ProgPeriodIntervalField{}
	field.Value = val
	return field
}

//ProgRptReqsField is a INT field
type ProgRptReqsField struct{ quickfix.IntValue }

//Tag returns tag.ProgRptReqs (414)
func (f ProgRptReqsField) Tag() quickfix.Tag { return tag.ProgRptReqs }

//NewProgRptReqs returns a new ProgRptReqsField initialized with val
func NewProgRptReqs(val int) *ProgRptReqsField {
	field := &ProgRptReqsField{}
	field.Value = val
	return field
}

//PublishTrdIndicatorField is a BOOLEAN field
type PublishTrdIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.PublishTrdIndicator (852)
func (f PublishTrdIndicatorField) Tag() quickfix.Tag { return tag.PublishTrdIndicator }

//NewPublishTrdIndicator returns a new PublishTrdIndicatorField initialized with val
func NewPublishTrdIndicator(val bool) *PublishTrdIndicatorField {
	field := &PublishTrdIndicatorField{}
	field.Value = val
	return field
}

//PutOrCallField is a INT field
type PutOrCallField struct{ quickfix.IntValue }

//Tag returns tag.PutOrCall (201)
func (f PutOrCallField) Tag() quickfix.Tag { return tag.PutOrCall }

//NewPutOrCall returns a new PutOrCallField initialized with val
func NewPutOrCall(val int) *PutOrCallField {
	field := &PutOrCallField{}
	field.Value = val
	return field
}

//QtyTypeField is a INT field
type QtyTypeField struct{ quickfix.IntValue }

//Tag returns tag.QtyType (854)
func (f QtyTypeField) Tag() quickfix.Tag { return tag.QtyType }

//NewQtyType returns a new QtyTypeField initialized with val
func NewQtyType(val int) *QtyTypeField {
	field := &QtyTypeField{}
	field.Value = val
	return field
}

//QuantityField is a QTY field
type QuantityField struct{ quickfix.QtyValue }

//Tag returns tag.Quantity (53)
func (f QuantityField) Tag() quickfix.Tag { return tag.Quantity }

//NewQuantity returns a new QuantityField initialized with val
func NewQuantity(val float64) *QuantityField {
	field := &QuantityField{}
	field.Value = val
	return field
}

//QuantityDateField is a LOCALMKTDATE field
type QuantityDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.QuantityDate (976)
func (f QuantityDateField) Tag() quickfix.Tag { return tag.QuantityDate }

//NewQuantityDate returns a new QuantityDateField initialized with val
func NewQuantityDate(val string) *QuantityDateField {
	field := &QuantityDateField{}
	field.Value = val
	return field
}

//QuantityTypeField is a INT field
type QuantityTypeField struct{ quickfix.IntValue }

//Tag returns tag.QuantityType (465)
func (f QuantityTypeField) Tag() quickfix.Tag { return tag.QuantityType }

//NewQuantityType returns a new QuantityTypeField initialized with val
func NewQuantityType(val int) *QuantityTypeField {
	field := &QuantityTypeField{}
	field.Value = val
	return field
}

//QuoteAckStatusField is a INT field
type QuoteAckStatusField struct{ quickfix.IntValue }

//Tag returns tag.QuoteAckStatus (297)
func (f QuoteAckStatusField) Tag() quickfix.Tag { return tag.QuoteAckStatus }

//NewQuoteAckStatus returns a new QuoteAckStatusField initialized with val
func NewQuoteAckStatus(val int) *QuoteAckStatusField {
	field := &QuoteAckStatusField{}
	field.Value = val
	return field
}

//QuoteCancelTypeField is a INT field
type QuoteCancelTypeField struct{ quickfix.IntValue }

//Tag returns tag.QuoteCancelType (298)
func (f QuoteCancelTypeField) Tag() quickfix.Tag { return tag.QuoteCancelType }

//NewQuoteCancelType returns a new QuoteCancelTypeField initialized with val
func NewQuoteCancelType(val int) *QuoteCancelTypeField {
	field := &QuoteCancelTypeField{}
	field.Value = val
	return field
}

//QuoteConditionField is a MULTIPLESTRINGVALUE field
type QuoteConditionField struct{ quickfix.MultipleStringValue }

//Tag returns tag.QuoteCondition (276)
func (f QuoteConditionField) Tag() quickfix.Tag { return tag.QuoteCondition }

//NewQuoteCondition returns a new QuoteConditionField initialized with val
func NewQuoteCondition(val string) *QuoteConditionField {
	field := &QuoteConditionField{}
	field.Value = val
	return field
}

//QuoteEntryIDField is a STRING field
type QuoteEntryIDField struct{ quickfix.StringValue }

//Tag returns tag.QuoteEntryID (299)
func (f QuoteEntryIDField) Tag() quickfix.Tag { return tag.QuoteEntryID }

//NewQuoteEntryID returns a new QuoteEntryIDField initialized with val
func NewQuoteEntryID(val string) *QuoteEntryIDField {
	field := &QuoteEntryIDField{}
	field.Value = val
	return field
}

//QuoteEntryRejectReasonField is a INT field
type QuoteEntryRejectReasonField struct{ quickfix.IntValue }

//Tag returns tag.QuoteEntryRejectReason (368)
func (f QuoteEntryRejectReasonField) Tag() quickfix.Tag { return tag.QuoteEntryRejectReason }

//NewQuoteEntryRejectReason returns a new QuoteEntryRejectReasonField initialized with val
func NewQuoteEntryRejectReason(val int) *QuoteEntryRejectReasonField {
	field := &QuoteEntryRejectReasonField{}
	field.Value = val
	return field
}

//QuoteEntryStatusField is a INT field
type QuoteEntryStatusField struct{ quickfix.IntValue }

//Tag returns tag.QuoteEntryStatus (1167)
func (f QuoteEntryStatusField) Tag() quickfix.Tag { return tag.QuoteEntryStatus }

//NewQuoteEntryStatus returns a new QuoteEntryStatusField initialized with val
func NewQuoteEntryStatus(val int) *QuoteEntryStatusField {
	field := &QuoteEntryStatusField{}
	field.Value = val
	return field
}

//QuoteIDField is a STRING field
type QuoteIDField struct{ quickfix.StringValue }

//Tag returns tag.QuoteID (117)
func (f QuoteIDField) Tag() quickfix.Tag { return tag.QuoteID }

//NewQuoteID returns a new QuoteIDField initialized with val
func NewQuoteID(val string) *QuoteIDField {
	field := &QuoteIDField{}
	field.Value = val
	return field
}

//QuoteMsgIDField is a STRING field
type QuoteMsgIDField struct{ quickfix.StringValue }

//Tag returns tag.QuoteMsgID (1166)
func (f QuoteMsgIDField) Tag() quickfix.Tag { return tag.QuoteMsgID }

//NewQuoteMsgID returns a new QuoteMsgIDField initialized with val
func NewQuoteMsgID(val string) *QuoteMsgIDField {
	field := &QuoteMsgIDField{}
	field.Value = val
	return field
}

//QuotePriceTypeField is a INT field
type QuotePriceTypeField struct{ quickfix.IntValue }

//Tag returns tag.QuotePriceType (692)
func (f QuotePriceTypeField) Tag() quickfix.Tag { return tag.QuotePriceType }

//NewQuotePriceType returns a new QuotePriceTypeField initialized with val
func NewQuotePriceType(val int) *QuotePriceTypeField {
	field := &QuotePriceTypeField{}
	field.Value = val
	return field
}

//QuoteQualifierField is a CHAR field
type QuoteQualifierField struct{ quickfix.CharValue }

//Tag returns tag.QuoteQualifier (695)
func (f QuoteQualifierField) Tag() quickfix.Tag { return tag.QuoteQualifier }

//NewQuoteQualifier returns a new QuoteQualifierField initialized with val
func NewQuoteQualifier(val string) *QuoteQualifierField {
	field := &QuoteQualifierField{}
	field.Value = val
	return field
}

//QuoteRejectReasonField is a INT field
type QuoteRejectReasonField struct{ quickfix.IntValue }

//Tag returns tag.QuoteRejectReason (300)
func (f QuoteRejectReasonField) Tag() quickfix.Tag { return tag.QuoteRejectReason }

//NewQuoteRejectReason returns a new QuoteRejectReasonField initialized with val
func NewQuoteRejectReason(val int) *QuoteRejectReasonField {
	field := &QuoteRejectReasonField{}
	field.Value = val
	return field
}

//QuoteReqIDField is a STRING field
type QuoteReqIDField struct{ quickfix.StringValue }

//Tag returns tag.QuoteReqID (131)
func (f QuoteReqIDField) Tag() quickfix.Tag { return tag.QuoteReqID }

//NewQuoteReqID returns a new QuoteReqIDField initialized with val
func NewQuoteReqID(val string) *QuoteReqIDField {
	field := &QuoteReqIDField{}
	field.Value = val
	return field
}

//QuoteRequestRejectReasonField is a INT field
type QuoteRequestRejectReasonField struct{ quickfix.IntValue }

//Tag returns tag.QuoteRequestRejectReason (658)
func (f QuoteRequestRejectReasonField) Tag() quickfix.Tag { return tag.QuoteRequestRejectReason }

//NewQuoteRequestRejectReason returns a new QuoteRequestRejectReasonField initialized with val
func NewQuoteRequestRejectReason(val int) *QuoteRequestRejectReasonField {
	field := &QuoteRequestRejectReasonField{}
	field.Value = val
	return field
}

//QuoteRequestTypeField is a INT field
type QuoteRequestTypeField struct{ quickfix.IntValue }

//Tag returns tag.QuoteRequestType (303)
func (f QuoteRequestTypeField) Tag() quickfix.Tag { return tag.QuoteRequestType }

//NewQuoteRequestType returns a new QuoteRequestTypeField initialized with val
func NewQuoteRequestType(val int) *QuoteRequestTypeField {
	field := &QuoteRequestTypeField{}
	field.Value = val
	return field
}

//QuoteRespIDField is a STRING field
type QuoteRespIDField struct{ quickfix.StringValue }

//Tag returns tag.QuoteRespID (693)
func (f QuoteRespIDField) Tag() quickfix.Tag { return tag.QuoteRespID }

//NewQuoteRespID returns a new QuoteRespIDField initialized with val
func NewQuoteRespID(val string) *QuoteRespIDField {
	field := &QuoteRespIDField{}
	field.Value = val
	return field
}

//QuoteRespTypeField is a INT field
type QuoteRespTypeField struct{ quickfix.IntValue }

//Tag returns tag.QuoteRespType (694)
func (f QuoteRespTypeField) Tag() quickfix.Tag { return tag.QuoteRespType }

//NewQuoteRespType returns a new QuoteRespTypeField initialized with val
func NewQuoteRespType(val int) *QuoteRespTypeField {
	field := &QuoteRespTypeField{}
	field.Value = val
	return field
}

//QuoteResponseLevelField is a INT field
type QuoteResponseLevelField struct{ quickfix.IntValue }

//Tag returns tag.QuoteResponseLevel (301)
func (f QuoteResponseLevelField) Tag() quickfix.Tag { return tag.QuoteResponseLevel }

//NewQuoteResponseLevel returns a new QuoteResponseLevelField initialized with val
func NewQuoteResponseLevel(val int) *QuoteResponseLevelField {
	field := &QuoteResponseLevelField{}
	field.Value = val
	return field
}

//QuoteSetIDField is a STRING field
type QuoteSetIDField struct{ quickfix.StringValue }

//Tag returns tag.QuoteSetID (302)
func (f QuoteSetIDField) Tag() quickfix.Tag { return tag.QuoteSetID }

//NewQuoteSetID returns a new QuoteSetIDField initialized with val
func NewQuoteSetID(val string) *QuoteSetIDField {
	field := &QuoteSetIDField{}
	field.Value = val
	return field
}

//QuoteSetValidUntilTimeField is a UTCTIMESTAMP field
type QuoteSetValidUntilTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.QuoteSetValidUntilTime (367)
func (f QuoteSetValidUntilTimeField) Tag() quickfix.Tag { return tag.QuoteSetValidUntilTime }

//QuoteStatusField is a INT field
type QuoteStatusField struct{ quickfix.IntValue }

//Tag returns tag.QuoteStatus (297)
func (f QuoteStatusField) Tag() quickfix.Tag { return tag.QuoteStatus }

//NewQuoteStatus returns a new QuoteStatusField initialized with val
func NewQuoteStatus(val int) *QuoteStatusField {
	field := &QuoteStatusField{}
	field.Value = val
	return field
}

//QuoteStatusReqIDField is a STRING field
type QuoteStatusReqIDField struct{ quickfix.StringValue }

//Tag returns tag.QuoteStatusReqID (649)
func (f QuoteStatusReqIDField) Tag() quickfix.Tag { return tag.QuoteStatusReqID }

//NewQuoteStatusReqID returns a new QuoteStatusReqIDField initialized with val
func NewQuoteStatusReqID(val string) *QuoteStatusReqIDField {
	field := &QuoteStatusReqIDField{}
	field.Value = val
	return field
}

//QuoteTypeField is a INT field
type QuoteTypeField struct{ quickfix.IntValue }

//Tag returns tag.QuoteType (537)
func (f QuoteTypeField) Tag() quickfix.Tag { return tag.QuoteType }

//NewQuoteType returns a new QuoteTypeField initialized with val
func NewQuoteType(val int) *QuoteTypeField {
	field := &QuoteTypeField{}
	field.Value = val
	return field
}

//RFQReqIDField is a STRING field
type RFQReqIDField struct{ quickfix.StringValue }

//Tag returns tag.RFQReqID (644)
func (f RFQReqIDField) Tag() quickfix.Tag { return tag.RFQReqID }

//NewRFQReqID returns a new RFQReqIDField initialized with val
func NewRFQReqID(val string) *RFQReqIDField {
	field := &RFQReqIDField{}
	field.Value = val
	return field
}

//RateSourceField is a INT field
type RateSourceField struct{ quickfix.IntValue }

//Tag returns tag.RateSource (1446)
func (f RateSourceField) Tag() quickfix.Tag { return tag.RateSource }

//NewRateSource returns a new RateSourceField initialized with val
func NewRateSource(val int) *RateSourceField {
	field := &RateSourceField{}
	field.Value = val
	return field
}

//RateSourceTypeField is a INT field
type RateSourceTypeField struct{ quickfix.IntValue }

//Tag returns tag.RateSourceType (1447)
func (f RateSourceTypeField) Tag() quickfix.Tag { return tag.RateSourceType }

//NewRateSourceType returns a new RateSourceTypeField initialized with val
func NewRateSourceType(val int) *RateSourceTypeField {
	field := &RateSourceTypeField{}
	field.Value = val
	return field
}

//RatioQtyField is a QTY field
type RatioQtyField struct{ quickfix.QtyValue }

//Tag returns tag.RatioQty (319)
func (f RatioQtyField) Tag() quickfix.Tag { return tag.RatioQty }

//NewRatioQty returns a new RatioQtyField initialized with val
func NewRatioQty(val float64) *RatioQtyField {
	field := &RatioQtyField{}
	field.Value = val
	return field
}

//RawDataField is a DATA field
type RawDataField struct{ quickfix.DataValue }

//Tag returns tag.RawData (96)
func (f RawDataField) Tag() quickfix.Tag { return tag.RawData }

//NewRawData returns a new RawDataField initialized with val
func NewRawData(val string) *RawDataField {
	field := &RawDataField{}
	field.Value = val
	return field
}

//RawDataLengthField is a LENGTH field
type RawDataLengthField struct{ quickfix.LengthValue }

//Tag returns tag.RawDataLength (95)
func (f RawDataLengthField) Tag() quickfix.Tag { return tag.RawDataLength }

//NewRawDataLength returns a new RawDataLengthField initialized with val
func NewRawDataLength(val int) *RawDataLengthField {
	field := &RawDataLengthField{}
	field.Value = val
	return field
}

//ReceivedDeptIDField is a STRING field
type ReceivedDeptIDField struct{ quickfix.StringValue }

//Tag returns tag.ReceivedDeptID (1030)
func (f ReceivedDeptIDField) Tag() quickfix.Tag { return tag.ReceivedDeptID }

//NewReceivedDeptID returns a new ReceivedDeptIDField initialized with val
func NewReceivedDeptID(val string) *ReceivedDeptIDField {
	field := &ReceivedDeptIDField{}
	field.Value = val
	return field
}

//RedemptionDateField is a LOCALMKTDATE field
type RedemptionDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.RedemptionDate (240)
func (f RedemptionDateField) Tag() quickfix.Tag { return tag.RedemptionDate }

//NewRedemptionDate returns a new RedemptionDateField initialized with val
func NewRedemptionDate(val string) *RedemptionDateField {
	field := &RedemptionDateField{}
	field.Value = val
	return field
}

//RefAllocIDField is a STRING field
type RefAllocIDField struct{ quickfix.StringValue }

//Tag returns tag.RefAllocID (72)
func (f RefAllocIDField) Tag() quickfix.Tag { return tag.RefAllocID }

//NewRefAllocID returns a new RefAllocIDField initialized with val
func NewRefAllocID(val string) *RefAllocIDField {
	field := &RefAllocIDField{}
	field.Value = val
	return field
}

//RefApplExtIDField is a INT field
type RefApplExtIDField struct{ quickfix.IntValue }

//Tag returns tag.RefApplExtID (1406)
func (f RefApplExtIDField) Tag() quickfix.Tag { return tag.RefApplExtID }

//NewRefApplExtID returns a new RefApplExtIDField initialized with val
func NewRefApplExtID(val int) *RefApplExtIDField {
	field := &RefApplExtIDField{}
	field.Value = val
	return field
}

//RefApplIDField is a STRING field
type RefApplIDField struct{ quickfix.StringValue }

//Tag returns tag.RefApplID (1355)
func (f RefApplIDField) Tag() quickfix.Tag { return tag.RefApplID }

//NewRefApplID returns a new RefApplIDField initialized with val
func NewRefApplID(val string) *RefApplIDField {
	field := &RefApplIDField{}
	field.Value = val
	return field
}

//RefApplLastSeqNumField is a SEQNUM field
type RefApplLastSeqNumField struct{ quickfix.SeqNumValue }

//Tag returns tag.RefApplLastSeqNum (1357)
func (f RefApplLastSeqNumField) Tag() quickfix.Tag { return tag.RefApplLastSeqNum }

//NewRefApplLastSeqNum returns a new RefApplLastSeqNumField initialized with val
func NewRefApplLastSeqNum(val int) *RefApplLastSeqNumField {
	field := &RefApplLastSeqNumField{}
	field.Value = val
	return field
}

//RefApplReqIDField is a STRING field
type RefApplReqIDField struct{ quickfix.StringValue }

//Tag returns tag.RefApplReqID (1433)
func (f RefApplReqIDField) Tag() quickfix.Tag { return tag.RefApplReqID }

//NewRefApplReqID returns a new RefApplReqIDField initialized with val
func NewRefApplReqID(val string) *RefApplReqIDField {
	field := &RefApplReqIDField{}
	field.Value = val
	return field
}

//RefApplVerIDField is a STRING field
type RefApplVerIDField struct{ quickfix.StringValue }

//Tag returns tag.RefApplVerID (1130)
func (f RefApplVerIDField) Tag() quickfix.Tag { return tag.RefApplVerID }

//NewRefApplVerID returns a new RefApplVerIDField initialized with val
func NewRefApplVerID(val string) *RefApplVerIDField {
	field := &RefApplVerIDField{}
	field.Value = val
	return field
}

//RefCompIDField is a STRING field
type RefCompIDField struct{ quickfix.StringValue }

//Tag returns tag.RefCompID (930)
func (f RefCompIDField) Tag() quickfix.Tag { return tag.RefCompID }

//NewRefCompID returns a new RefCompIDField initialized with val
func NewRefCompID(val string) *RefCompIDField {
	field := &RefCompIDField{}
	field.Value = val
	return field
}

//RefCstmApplVerIDField is a STRING field
type RefCstmApplVerIDField struct{ quickfix.StringValue }

//Tag returns tag.RefCstmApplVerID (1131)
func (f RefCstmApplVerIDField) Tag() quickfix.Tag { return tag.RefCstmApplVerID }

//NewRefCstmApplVerID returns a new RefCstmApplVerIDField initialized with val
func NewRefCstmApplVerID(val string) *RefCstmApplVerIDField {
	field := &RefCstmApplVerIDField{}
	field.Value = val
	return field
}

//RefMsgTypeField is a STRING field
type RefMsgTypeField struct{ quickfix.StringValue }

//Tag returns tag.RefMsgType (372)
func (f RefMsgTypeField) Tag() quickfix.Tag { return tag.RefMsgType }

//NewRefMsgType returns a new RefMsgTypeField initialized with val
func NewRefMsgType(val string) *RefMsgTypeField {
	field := &RefMsgTypeField{}
	field.Value = val
	return field
}

//RefOrdIDReasonField is a INT field
type RefOrdIDReasonField struct{ quickfix.IntValue }

//Tag returns tag.RefOrdIDReason (1431)
func (f RefOrdIDReasonField) Tag() quickfix.Tag { return tag.RefOrdIDReason }

//NewRefOrdIDReason returns a new RefOrdIDReasonField initialized with val
func NewRefOrdIDReason(val int) *RefOrdIDReasonField {
	field := &RefOrdIDReasonField{}
	field.Value = val
	return field
}

//RefOrderIDField is a STRING field
type RefOrderIDField struct{ quickfix.StringValue }

//Tag returns tag.RefOrderID (1080)
func (f RefOrderIDField) Tag() quickfix.Tag { return tag.RefOrderID }

//NewRefOrderID returns a new RefOrderIDField initialized with val
func NewRefOrderID(val string) *RefOrderIDField {
	field := &RefOrderIDField{}
	field.Value = val
	return field
}

//RefOrderIDSourceField is a CHAR field
type RefOrderIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.RefOrderIDSource (1081)
func (f RefOrderIDSourceField) Tag() quickfix.Tag { return tag.RefOrderIDSource }

//NewRefOrderIDSource returns a new RefOrderIDSourceField initialized with val
func NewRefOrderIDSource(val string) *RefOrderIDSourceField {
	field := &RefOrderIDSourceField{}
	field.Value = val
	return field
}

//RefSeqNumField is a SEQNUM field
type RefSeqNumField struct{ quickfix.SeqNumValue }

//Tag returns tag.RefSeqNum (45)
func (f RefSeqNumField) Tag() quickfix.Tag { return tag.RefSeqNum }

//NewRefSeqNum returns a new RefSeqNumField initialized with val
func NewRefSeqNum(val int) *RefSeqNumField {
	field := &RefSeqNumField{}
	field.Value = val
	return field
}

//RefSubIDField is a STRING field
type RefSubIDField struct{ quickfix.StringValue }

//Tag returns tag.RefSubID (931)
func (f RefSubIDField) Tag() quickfix.Tag { return tag.RefSubID }

//NewRefSubID returns a new RefSubIDField initialized with val
func NewRefSubID(val string) *RefSubIDField {
	field := &RefSubIDField{}
	field.Value = val
	return field
}

//RefTagIDField is a INT field
type RefTagIDField struct{ quickfix.IntValue }

//Tag returns tag.RefTagID (371)
func (f RefTagIDField) Tag() quickfix.Tag { return tag.RefTagID }

//NewRefTagID returns a new RefTagIDField initialized with val
func NewRefTagID(val int) *RefTagIDField {
	field := &RefTagIDField{}
	field.Value = val
	return field
}

//ReferencePageField is a STRING field
type ReferencePageField struct{ quickfix.StringValue }

//Tag returns tag.ReferencePage (1448)
func (f ReferencePageField) Tag() quickfix.Tag { return tag.ReferencePage }

//NewReferencePage returns a new ReferencePageField initialized with val
func NewReferencePage(val string) *ReferencePageField {
	field := &ReferencePageField{}
	field.Value = val
	return field
}

//RefreshIndicatorField is a BOOLEAN field
type RefreshIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.RefreshIndicator (1187)
func (f RefreshIndicatorField) Tag() quickfix.Tag { return tag.RefreshIndicator }

//NewRefreshIndicator returns a new RefreshIndicatorField initialized with val
func NewRefreshIndicator(val bool) *RefreshIndicatorField {
	field := &RefreshIndicatorField{}
	field.Value = val
	return field
}

//RefreshQtyField is a QTY field
type RefreshQtyField struct{ quickfix.QtyValue }

//Tag returns tag.RefreshQty (1088)
func (f RefreshQtyField) Tag() quickfix.Tag { return tag.RefreshQty }

//NewRefreshQty returns a new RefreshQtyField initialized with val
func NewRefreshQty(val float64) *RefreshQtyField {
	field := &RefreshQtyField{}
	field.Value = val
	return field
}

//RegistAcctTypeField is a STRING field
type RegistAcctTypeField struct{ quickfix.StringValue }

//Tag returns tag.RegistAcctType (493)
func (f RegistAcctTypeField) Tag() quickfix.Tag { return tag.RegistAcctType }

//NewRegistAcctType returns a new RegistAcctTypeField initialized with val
func NewRegistAcctType(val string) *RegistAcctTypeField {
	field := &RegistAcctTypeField{}
	field.Value = val
	return field
}

//RegistDetlsField is a STRING field
type RegistDetlsField struct{ quickfix.StringValue }

//Tag returns tag.RegistDetls (509)
func (f RegistDetlsField) Tag() quickfix.Tag { return tag.RegistDetls }

//NewRegistDetls returns a new RegistDetlsField initialized with val
func NewRegistDetls(val string) *RegistDetlsField {
	field := &RegistDetlsField{}
	field.Value = val
	return field
}

//RegistDtlsField is a STRING field
type RegistDtlsField struct{ quickfix.StringValue }

//Tag returns tag.RegistDtls (509)
func (f RegistDtlsField) Tag() quickfix.Tag { return tag.RegistDtls }

//NewRegistDtls returns a new RegistDtlsField initialized with val
func NewRegistDtls(val string) *RegistDtlsField {
	field := &RegistDtlsField{}
	field.Value = val
	return field
}

//RegistEmailField is a STRING field
type RegistEmailField struct{ quickfix.StringValue }

//Tag returns tag.RegistEmail (511)
func (f RegistEmailField) Tag() quickfix.Tag { return tag.RegistEmail }

//NewRegistEmail returns a new RegistEmailField initialized with val
func NewRegistEmail(val string) *RegistEmailField {
	field := &RegistEmailField{}
	field.Value = val
	return field
}

//RegistIDField is a STRING field
type RegistIDField struct{ quickfix.StringValue }

//Tag returns tag.RegistID (513)
func (f RegistIDField) Tag() quickfix.Tag { return tag.RegistID }

//NewRegistID returns a new RegistIDField initialized with val
func NewRegistID(val string) *RegistIDField {
	field := &RegistIDField{}
	field.Value = val
	return field
}

//RegistRefIDField is a STRING field
type RegistRefIDField struct{ quickfix.StringValue }

//Tag returns tag.RegistRefID (508)
func (f RegistRefIDField) Tag() quickfix.Tag { return tag.RegistRefID }

//NewRegistRefID returns a new RegistRefIDField initialized with val
func NewRegistRefID(val string) *RegistRefIDField {
	field := &RegistRefIDField{}
	field.Value = val
	return field
}

//RegistRejReasonCodeField is a INT field
type RegistRejReasonCodeField struct{ quickfix.IntValue }

//Tag returns tag.RegistRejReasonCode (507)
func (f RegistRejReasonCodeField) Tag() quickfix.Tag { return tag.RegistRejReasonCode }

//NewRegistRejReasonCode returns a new RegistRejReasonCodeField initialized with val
func NewRegistRejReasonCode(val int) *RegistRejReasonCodeField {
	field := &RegistRejReasonCodeField{}
	field.Value = val
	return field
}

//RegistRejReasonTextField is a STRING field
type RegistRejReasonTextField struct{ quickfix.StringValue }

//Tag returns tag.RegistRejReasonText (496)
func (f RegistRejReasonTextField) Tag() quickfix.Tag { return tag.RegistRejReasonText }

//NewRegistRejReasonText returns a new RegistRejReasonTextField initialized with val
func NewRegistRejReasonText(val string) *RegistRejReasonTextField {
	field := &RegistRejReasonTextField{}
	field.Value = val
	return field
}

//RegistStatusField is a CHAR field
type RegistStatusField struct{ quickfix.CharValue }

//Tag returns tag.RegistStatus (506)
func (f RegistStatusField) Tag() quickfix.Tag { return tag.RegistStatus }

//NewRegistStatus returns a new RegistStatusField initialized with val
func NewRegistStatus(val string) *RegistStatusField {
	field := &RegistStatusField{}
	field.Value = val
	return field
}

//RegistTransTypeField is a CHAR field
type RegistTransTypeField struct{ quickfix.CharValue }

//Tag returns tag.RegistTransType (514)
func (f RegistTransTypeField) Tag() quickfix.Tag { return tag.RegistTransType }

//NewRegistTransType returns a new RegistTransTypeField initialized with val
func NewRegistTransType(val string) *RegistTransTypeField {
	field := &RegistTransTypeField{}
	field.Value = val
	return field
}

//RejectTextField is a STRING field
type RejectTextField struct{ quickfix.StringValue }

//Tag returns tag.RejectText (1328)
func (f RejectTextField) Tag() quickfix.Tag { return tag.RejectText }

//NewRejectText returns a new RejectTextField initialized with val
func NewRejectText(val string) *RejectTextField {
	field := &RejectTextField{}
	field.Value = val
	return field
}

//RelSymTransactTimeField is a UTCTIMESTAMP field
type RelSymTransactTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.RelSymTransactTime (1504)
func (f RelSymTransactTimeField) Tag() quickfix.Tag { return tag.RelSymTransactTime }

//RelatdSymField is a STRING field
type RelatdSymField struct{ quickfix.StringValue }

//Tag returns tag.RelatdSym (46)
func (f RelatdSymField) Tag() quickfix.Tag { return tag.RelatdSym }

//NewRelatdSym returns a new RelatdSymField initialized with val
func NewRelatdSym(val string) *RelatdSymField {
	field := &RelatdSymField{}
	field.Value = val
	return field
}

//RelatedContextPartyIDField is a STRING field
type RelatedContextPartyIDField struct{ quickfix.StringValue }

//Tag returns tag.RelatedContextPartyID (1576)
func (f RelatedContextPartyIDField) Tag() quickfix.Tag { return tag.RelatedContextPartyID }

//NewRelatedContextPartyID returns a new RelatedContextPartyIDField initialized with val
func NewRelatedContextPartyID(val string) *RelatedContextPartyIDField {
	field := &RelatedContextPartyIDField{}
	field.Value = val
	return field
}

//RelatedContextPartyIDSourceField is a CHAR field
type RelatedContextPartyIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.RelatedContextPartyIDSource (1577)
func (f RelatedContextPartyIDSourceField) Tag() quickfix.Tag { return tag.RelatedContextPartyIDSource }

//NewRelatedContextPartyIDSource returns a new RelatedContextPartyIDSourceField initialized with val
func NewRelatedContextPartyIDSource(val string) *RelatedContextPartyIDSourceField {
	field := &RelatedContextPartyIDSourceField{}
	field.Value = val
	return field
}

//RelatedContextPartyRoleField is a INT field
type RelatedContextPartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.RelatedContextPartyRole (1578)
func (f RelatedContextPartyRoleField) Tag() quickfix.Tag { return tag.RelatedContextPartyRole }

//NewRelatedContextPartyRole returns a new RelatedContextPartyRoleField initialized with val
func NewRelatedContextPartyRole(val int) *RelatedContextPartyRoleField {
	field := &RelatedContextPartyRoleField{}
	field.Value = val
	return field
}

//RelatedContextPartySubIDField is a STRING field
type RelatedContextPartySubIDField struct{ quickfix.StringValue }

//Tag returns tag.RelatedContextPartySubID (1580)
func (f RelatedContextPartySubIDField) Tag() quickfix.Tag { return tag.RelatedContextPartySubID }

//NewRelatedContextPartySubID returns a new RelatedContextPartySubIDField initialized with val
func NewRelatedContextPartySubID(val string) *RelatedContextPartySubIDField {
	field := &RelatedContextPartySubIDField{}
	field.Value = val
	return field
}

//RelatedContextPartySubIDTypeField is a INT field
type RelatedContextPartySubIDTypeField struct{ quickfix.IntValue }

//Tag returns tag.RelatedContextPartySubIDType (1581)
func (f RelatedContextPartySubIDTypeField) Tag() quickfix.Tag { return tag.RelatedContextPartySubIDType }

//NewRelatedContextPartySubIDType returns a new RelatedContextPartySubIDTypeField initialized with val
func NewRelatedContextPartySubIDType(val int) *RelatedContextPartySubIDTypeField {
	field := &RelatedContextPartySubIDTypeField{}
	field.Value = val
	return field
}

//RelatedPartyAltIDField is a STRING field
type RelatedPartyAltIDField struct{ quickfix.StringValue }

//Tag returns tag.RelatedPartyAltID (1570)
func (f RelatedPartyAltIDField) Tag() quickfix.Tag { return tag.RelatedPartyAltID }

//NewRelatedPartyAltID returns a new RelatedPartyAltIDField initialized with val
func NewRelatedPartyAltID(val string) *RelatedPartyAltIDField {
	field := &RelatedPartyAltIDField{}
	field.Value = val
	return field
}

//RelatedPartyAltIDSourceField is a CHAR field
type RelatedPartyAltIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.RelatedPartyAltIDSource (1571)
func (f RelatedPartyAltIDSourceField) Tag() quickfix.Tag { return tag.RelatedPartyAltIDSource }

//NewRelatedPartyAltIDSource returns a new RelatedPartyAltIDSourceField initialized with val
func NewRelatedPartyAltIDSource(val string) *RelatedPartyAltIDSourceField {
	field := &RelatedPartyAltIDSourceField{}
	field.Value = val
	return field
}

//RelatedPartyAltSubIDField is a STRING field
type RelatedPartyAltSubIDField struct{ quickfix.StringValue }

//Tag returns tag.RelatedPartyAltSubID (1573)
func (f RelatedPartyAltSubIDField) Tag() quickfix.Tag { return tag.RelatedPartyAltSubID }

//NewRelatedPartyAltSubID returns a new RelatedPartyAltSubIDField initialized with val
func NewRelatedPartyAltSubID(val string) *RelatedPartyAltSubIDField {
	field := &RelatedPartyAltSubIDField{}
	field.Value = val
	return field
}

//RelatedPartyAltSubIDTypeField is a INT field
type RelatedPartyAltSubIDTypeField struct{ quickfix.IntValue }

//Tag returns tag.RelatedPartyAltSubIDType (1574)
func (f RelatedPartyAltSubIDTypeField) Tag() quickfix.Tag { return tag.RelatedPartyAltSubIDType }

//NewRelatedPartyAltSubIDType returns a new RelatedPartyAltSubIDTypeField initialized with val
func NewRelatedPartyAltSubIDType(val int) *RelatedPartyAltSubIDTypeField {
	field := &RelatedPartyAltSubIDTypeField{}
	field.Value = val
	return field
}

//RelatedPartyIDField is a STRING field
type RelatedPartyIDField struct{ quickfix.StringValue }

//Tag returns tag.RelatedPartyID (1563)
func (f RelatedPartyIDField) Tag() quickfix.Tag { return tag.RelatedPartyID }

//NewRelatedPartyID returns a new RelatedPartyIDField initialized with val
func NewRelatedPartyID(val string) *RelatedPartyIDField {
	field := &RelatedPartyIDField{}
	field.Value = val
	return field
}

//RelatedPartyIDSourceField is a CHAR field
type RelatedPartyIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.RelatedPartyIDSource (1564)
func (f RelatedPartyIDSourceField) Tag() quickfix.Tag { return tag.RelatedPartyIDSource }

//NewRelatedPartyIDSource returns a new RelatedPartyIDSourceField initialized with val
func NewRelatedPartyIDSource(val string) *RelatedPartyIDSourceField {
	field := &RelatedPartyIDSourceField{}
	field.Value = val
	return field
}

//RelatedPartyRoleField is a INT field
type RelatedPartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.RelatedPartyRole (1565)
func (f RelatedPartyRoleField) Tag() quickfix.Tag { return tag.RelatedPartyRole }

//NewRelatedPartyRole returns a new RelatedPartyRoleField initialized with val
func NewRelatedPartyRole(val int) *RelatedPartyRoleField {
	field := &RelatedPartyRoleField{}
	field.Value = val
	return field
}

//RelatedPartySubIDField is a STRING field
type RelatedPartySubIDField struct{ quickfix.StringValue }

//Tag returns tag.RelatedPartySubID (1567)
func (f RelatedPartySubIDField) Tag() quickfix.Tag { return tag.RelatedPartySubID }

//NewRelatedPartySubID returns a new RelatedPartySubIDField initialized with val
func NewRelatedPartySubID(val string) *RelatedPartySubIDField {
	field := &RelatedPartySubIDField{}
	field.Value = val
	return field
}

//RelatedPartySubIDTypeField is a INT field
type RelatedPartySubIDTypeField struct{ quickfix.IntValue }

//Tag returns tag.RelatedPartySubIDType (1568)
func (f RelatedPartySubIDTypeField) Tag() quickfix.Tag { return tag.RelatedPartySubIDType }

//NewRelatedPartySubIDType returns a new RelatedPartySubIDTypeField initialized with val
func NewRelatedPartySubIDType(val int) *RelatedPartySubIDTypeField {
	field := &RelatedPartySubIDTypeField{}
	field.Value = val
	return field
}

//RelationshipRiskCFICodeField is a STRING field
type RelationshipRiskCFICodeField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskCFICode (1599)
func (f RelationshipRiskCFICodeField) Tag() quickfix.Tag { return tag.RelationshipRiskCFICode }

//NewRelationshipRiskCFICode returns a new RelationshipRiskCFICodeField initialized with val
func NewRelationshipRiskCFICode(val string) *RelationshipRiskCFICodeField {
	field := &RelationshipRiskCFICodeField{}
	field.Value = val
	return field
}

//RelationshipRiskCouponRateField is a PERCENTAGE field
type RelationshipRiskCouponRateField struct{ quickfix.PercentageValue }

//Tag returns tag.RelationshipRiskCouponRate (1608)
func (f RelationshipRiskCouponRateField) Tag() quickfix.Tag { return tag.RelationshipRiskCouponRate }

//NewRelationshipRiskCouponRate returns a new RelationshipRiskCouponRateField initialized with val
func NewRelationshipRiskCouponRate(val float64) *RelationshipRiskCouponRateField {
	field := &RelationshipRiskCouponRateField{}
	field.Value = val
	return field
}

//RelationshipRiskEncodedSecurityDescField is a DATA field
type RelationshipRiskEncodedSecurityDescField struct{ quickfix.DataValue }

//Tag returns tag.RelationshipRiskEncodedSecurityDesc (1619)
func (f RelationshipRiskEncodedSecurityDescField) Tag() quickfix.Tag {
	return tag.RelationshipRiskEncodedSecurityDesc
}

//NewRelationshipRiskEncodedSecurityDesc returns a new RelationshipRiskEncodedSecurityDescField initialized with val
func NewRelationshipRiskEncodedSecurityDesc(val string) *RelationshipRiskEncodedSecurityDescField {
	field := &RelationshipRiskEncodedSecurityDescField{}
	field.Value = val
	return field
}

//RelationshipRiskEncodedSecurityDescLenField is a LENGTH field
type RelationshipRiskEncodedSecurityDescLenField struct{ quickfix.LengthValue }

//Tag returns tag.RelationshipRiskEncodedSecurityDescLen (1618)
func (f RelationshipRiskEncodedSecurityDescLenField) Tag() quickfix.Tag {
	return tag.RelationshipRiskEncodedSecurityDescLen
}

//NewRelationshipRiskEncodedSecurityDescLen returns a new RelationshipRiskEncodedSecurityDescLenField initialized with val
func NewRelationshipRiskEncodedSecurityDescLen(val int) *RelationshipRiskEncodedSecurityDescLenField {
	field := &RelationshipRiskEncodedSecurityDescLenField{}
	field.Value = val
	return field
}

//RelationshipRiskFlexibleIndicatorField is a BOOLEAN field
type RelationshipRiskFlexibleIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.RelationshipRiskFlexibleIndicator (1607)
func (f RelationshipRiskFlexibleIndicatorField) Tag() quickfix.Tag {
	return tag.RelationshipRiskFlexibleIndicator
}

//NewRelationshipRiskFlexibleIndicator returns a new RelationshipRiskFlexibleIndicatorField initialized with val
func NewRelationshipRiskFlexibleIndicator(val bool) *RelationshipRiskFlexibleIndicatorField {
	field := &RelationshipRiskFlexibleIndicatorField{}
	field.Value = val
	return field
}

//RelationshipRiskInstrumentMultiplierField is a FLOAT field
type RelationshipRiskInstrumentMultiplierField struct{ quickfix.FloatValue }

//Tag returns tag.RelationshipRiskInstrumentMultiplier (1612)
func (f RelationshipRiskInstrumentMultiplierField) Tag() quickfix.Tag {
	return tag.RelationshipRiskInstrumentMultiplier
}

//NewRelationshipRiskInstrumentMultiplier returns a new RelationshipRiskInstrumentMultiplierField initialized with val
func NewRelationshipRiskInstrumentMultiplier(val float64) *RelationshipRiskInstrumentMultiplierField {
	field := &RelationshipRiskInstrumentMultiplierField{}
	field.Value = val
	return field
}

//RelationshipRiskInstrumentOperatorField is a INT field
type RelationshipRiskInstrumentOperatorField struct{ quickfix.IntValue }

//Tag returns tag.RelationshipRiskInstrumentOperator (1588)
func (f RelationshipRiskInstrumentOperatorField) Tag() quickfix.Tag {
	return tag.RelationshipRiskInstrumentOperator
}

//NewRelationshipRiskInstrumentOperator returns a new RelationshipRiskInstrumentOperatorField initialized with val
func NewRelationshipRiskInstrumentOperator(val int) *RelationshipRiskInstrumentOperatorField {
	field := &RelationshipRiskInstrumentOperatorField{}
	field.Value = val
	return field
}

//RelationshipRiskInstrumentSettlTypeField is a STRING field
type RelationshipRiskInstrumentSettlTypeField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskInstrumentSettlType (1611)
func (f RelationshipRiskInstrumentSettlTypeField) Tag() quickfix.Tag {
	return tag.RelationshipRiskInstrumentSettlType
}

//NewRelationshipRiskInstrumentSettlType returns a new RelationshipRiskInstrumentSettlTypeField initialized with val
func NewRelationshipRiskInstrumentSettlType(val string) *RelationshipRiskInstrumentSettlTypeField {
	field := &RelationshipRiskInstrumentSettlTypeField{}
	field.Value = val
	return field
}

//RelationshipRiskLimitAmountField is a AMT field
type RelationshipRiskLimitAmountField struct{ quickfix.AmtValue }

//Tag returns tag.RelationshipRiskLimitAmount (1584)
func (f RelationshipRiskLimitAmountField) Tag() quickfix.Tag { return tag.RelationshipRiskLimitAmount }

//NewRelationshipRiskLimitAmount returns a new RelationshipRiskLimitAmountField initialized with val
func NewRelationshipRiskLimitAmount(val float64) *RelationshipRiskLimitAmountField {
	field := &RelationshipRiskLimitAmountField{}
	field.Value = val
	return field
}

//RelationshipRiskLimitCurrencyField is a CURRENCY field
type RelationshipRiskLimitCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.RelationshipRiskLimitCurrency (1585)
func (f RelationshipRiskLimitCurrencyField) Tag() quickfix.Tag {
	return tag.RelationshipRiskLimitCurrency
}

//NewRelationshipRiskLimitCurrency returns a new RelationshipRiskLimitCurrencyField initialized with val
func NewRelationshipRiskLimitCurrency(val string) *RelationshipRiskLimitCurrencyField {
	field := &RelationshipRiskLimitCurrencyField{}
	field.Value = val
	return field
}

//RelationshipRiskLimitPlatformField is a STRING field
type RelationshipRiskLimitPlatformField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskLimitPlatform (1586)
func (f RelationshipRiskLimitPlatformField) Tag() quickfix.Tag {
	return tag.RelationshipRiskLimitPlatform
}

//NewRelationshipRiskLimitPlatform returns a new RelationshipRiskLimitPlatformField initialized with val
func NewRelationshipRiskLimitPlatform(val string) *RelationshipRiskLimitPlatformField {
	field := &RelationshipRiskLimitPlatformField{}
	field.Value = val
	return field
}

//RelationshipRiskLimitTypeField is a INT field
type RelationshipRiskLimitTypeField struct{ quickfix.IntValue }

//Tag returns tag.RelationshipRiskLimitType (1583)
func (f RelationshipRiskLimitTypeField) Tag() quickfix.Tag { return tag.RelationshipRiskLimitType }

//NewRelationshipRiskLimitType returns a new RelationshipRiskLimitTypeField initialized with val
func NewRelationshipRiskLimitType(val int) *RelationshipRiskLimitTypeField {
	field := &RelationshipRiskLimitTypeField{}
	field.Value = val
	return field
}

//RelationshipRiskMaturityMonthYearField is a MONTHYEAR field
type RelationshipRiskMaturityMonthYearField struct{ quickfix.MonthYearValue }

//Tag returns tag.RelationshipRiskMaturityMonthYear (1602)
func (f RelationshipRiskMaturityMonthYearField) Tag() quickfix.Tag {
	return tag.RelationshipRiskMaturityMonthYear
}

//NewRelationshipRiskMaturityMonthYear returns a new RelationshipRiskMaturityMonthYearField initialized with val
func NewRelationshipRiskMaturityMonthYear(val string) *RelationshipRiskMaturityMonthYearField {
	field := &RelationshipRiskMaturityMonthYearField{}
	field.Value = val
	return field
}

//RelationshipRiskMaturityTimeField is a TZTIMEONLY field
type RelationshipRiskMaturityTimeField struct{ quickfix.TZTimeOnlyValue }

//Tag returns tag.RelationshipRiskMaturityTime (1603)
func (f RelationshipRiskMaturityTimeField) Tag() quickfix.Tag { return tag.RelationshipRiskMaturityTime }

//RelationshipRiskProductField is a INT field
type RelationshipRiskProductField struct{ quickfix.IntValue }

//Tag returns tag.RelationshipRiskProduct (1596)
func (f RelationshipRiskProductField) Tag() quickfix.Tag { return tag.RelationshipRiskProduct }

//NewRelationshipRiskProduct returns a new RelationshipRiskProductField initialized with val
func NewRelationshipRiskProduct(val int) *RelationshipRiskProductField {
	field := &RelationshipRiskProductField{}
	field.Value = val
	return field
}

//RelationshipRiskProductComplexField is a STRING field
type RelationshipRiskProductComplexField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskProductComplex (1597)
func (f RelationshipRiskProductComplexField) Tag() quickfix.Tag {
	return tag.RelationshipRiskProductComplex
}

//NewRelationshipRiskProductComplex returns a new RelationshipRiskProductComplexField initialized with val
func NewRelationshipRiskProductComplex(val string) *RelationshipRiskProductComplexField {
	field := &RelationshipRiskProductComplexField{}
	field.Value = val
	return field
}

//RelationshipRiskPutOrCallField is a INT field
type RelationshipRiskPutOrCallField struct{ quickfix.IntValue }

//Tag returns tag.RelationshipRiskPutOrCall (1606)
func (f RelationshipRiskPutOrCallField) Tag() quickfix.Tag { return tag.RelationshipRiskPutOrCall }

//NewRelationshipRiskPutOrCall returns a new RelationshipRiskPutOrCallField initialized with val
func NewRelationshipRiskPutOrCall(val int) *RelationshipRiskPutOrCallField {
	field := &RelationshipRiskPutOrCallField{}
	field.Value = val
	return field
}

//RelationshipRiskRestructuringTypeField is a STRING field
type RelationshipRiskRestructuringTypeField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskRestructuringType (1604)
func (f RelationshipRiskRestructuringTypeField) Tag() quickfix.Tag {
	return tag.RelationshipRiskRestructuringType
}

//NewRelationshipRiskRestructuringType returns a new RelationshipRiskRestructuringTypeField initialized with val
func NewRelationshipRiskRestructuringType(val string) *RelationshipRiskRestructuringTypeField {
	field := &RelationshipRiskRestructuringTypeField{}
	field.Value = val
	return field
}

//RelationshipRiskSecurityAltIDField is a STRING field
type RelationshipRiskSecurityAltIDField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskSecurityAltID (1594)
func (f RelationshipRiskSecurityAltIDField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecurityAltID
}

//NewRelationshipRiskSecurityAltID returns a new RelationshipRiskSecurityAltIDField initialized with val
func NewRelationshipRiskSecurityAltID(val string) *RelationshipRiskSecurityAltIDField {
	field := &RelationshipRiskSecurityAltIDField{}
	field.Value = val
	return field
}

//RelationshipRiskSecurityAltIDSourceField is a STRING field
type RelationshipRiskSecurityAltIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskSecurityAltIDSource (1595)
func (f RelationshipRiskSecurityAltIDSourceField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecurityAltIDSource
}

//NewRelationshipRiskSecurityAltIDSource returns a new RelationshipRiskSecurityAltIDSourceField initialized with val
func NewRelationshipRiskSecurityAltIDSource(val string) *RelationshipRiskSecurityAltIDSourceField {
	field := &RelationshipRiskSecurityAltIDSourceField{}
	field.Value = val
	return field
}

//RelationshipRiskSecurityDescField is a STRING field
type RelationshipRiskSecurityDescField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskSecurityDesc (1610)
func (f RelationshipRiskSecurityDescField) Tag() quickfix.Tag { return tag.RelationshipRiskSecurityDesc }

//NewRelationshipRiskSecurityDesc returns a new RelationshipRiskSecurityDescField initialized with val
func NewRelationshipRiskSecurityDesc(val string) *RelationshipRiskSecurityDescField {
	field := &RelationshipRiskSecurityDescField{}
	field.Value = val
	return field
}

//RelationshipRiskSecurityExchangeField is a EXCHANGE field
type RelationshipRiskSecurityExchangeField struct{ quickfix.ExchangeValue }

//Tag returns tag.RelationshipRiskSecurityExchange (1609)
func (f RelationshipRiskSecurityExchangeField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecurityExchange
}

//NewRelationshipRiskSecurityExchange returns a new RelationshipRiskSecurityExchangeField initialized with val
func NewRelationshipRiskSecurityExchange(val string) *RelationshipRiskSecurityExchangeField {
	field := &RelationshipRiskSecurityExchangeField{}
	field.Value = val
	return field
}

//RelationshipRiskSecurityGroupField is a STRING field
type RelationshipRiskSecurityGroupField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskSecurityGroup (1598)
func (f RelationshipRiskSecurityGroupField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecurityGroup
}

//NewRelationshipRiskSecurityGroup returns a new RelationshipRiskSecurityGroupField initialized with val
func NewRelationshipRiskSecurityGroup(val string) *RelationshipRiskSecurityGroupField {
	field := &RelationshipRiskSecurityGroupField{}
	field.Value = val
	return field
}

//RelationshipRiskSecurityIDField is a STRING field
type RelationshipRiskSecurityIDField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskSecurityID (1591)
func (f RelationshipRiskSecurityIDField) Tag() quickfix.Tag { return tag.RelationshipRiskSecurityID }

//NewRelationshipRiskSecurityID returns a new RelationshipRiskSecurityIDField initialized with val
func NewRelationshipRiskSecurityID(val string) *RelationshipRiskSecurityIDField {
	field := &RelationshipRiskSecurityIDField{}
	field.Value = val
	return field
}

//RelationshipRiskSecurityIDSourceField is a STRING field
type RelationshipRiskSecurityIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskSecurityIDSource (1592)
func (f RelationshipRiskSecurityIDSourceField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecurityIDSource
}

//NewRelationshipRiskSecurityIDSource returns a new RelationshipRiskSecurityIDSourceField initialized with val
func NewRelationshipRiskSecurityIDSource(val string) *RelationshipRiskSecurityIDSourceField {
	field := &RelationshipRiskSecurityIDSourceField{}
	field.Value = val
	return field
}

//RelationshipRiskSecuritySubTypeField is a STRING field
type RelationshipRiskSecuritySubTypeField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskSecuritySubType (1601)
func (f RelationshipRiskSecuritySubTypeField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecuritySubType
}

//NewRelationshipRiskSecuritySubType returns a new RelationshipRiskSecuritySubTypeField initialized with val
func NewRelationshipRiskSecuritySubType(val string) *RelationshipRiskSecuritySubTypeField {
	field := &RelationshipRiskSecuritySubTypeField{}
	field.Value = val
	return field
}

//RelationshipRiskSecurityTypeField is a STRING field
type RelationshipRiskSecurityTypeField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskSecurityType (1600)
func (f RelationshipRiskSecurityTypeField) Tag() quickfix.Tag { return tag.RelationshipRiskSecurityType }

//NewRelationshipRiskSecurityType returns a new RelationshipRiskSecurityTypeField initialized with val
func NewRelationshipRiskSecurityType(val string) *RelationshipRiskSecurityTypeField {
	field := &RelationshipRiskSecurityTypeField{}
	field.Value = val
	return field
}

//RelationshipRiskSeniorityField is a STRING field
type RelationshipRiskSeniorityField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskSeniority (1605)
func (f RelationshipRiskSeniorityField) Tag() quickfix.Tag { return tag.RelationshipRiskSeniority }

//NewRelationshipRiskSeniority returns a new RelationshipRiskSeniorityField initialized with val
func NewRelationshipRiskSeniority(val string) *RelationshipRiskSeniorityField {
	field := &RelationshipRiskSeniorityField{}
	field.Value = val
	return field
}

//RelationshipRiskSymbolField is a STRING field
type RelationshipRiskSymbolField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskSymbol (1589)
func (f RelationshipRiskSymbolField) Tag() quickfix.Tag { return tag.RelationshipRiskSymbol }

//NewRelationshipRiskSymbol returns a new RelationshipRiskSymbolField initialized with val
func NewRelationshipRiskSymbol(val string) *RelationshipRiskSymbolField {
	field := &RelationshipRiskSymbolField{}
	field.Value = val
	return field
}

//RelationshipRiskSymbolSfxField is a STRING field
type RelationshipRiskSymbolSfxField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskSymbolSfx (1590)
func (f RelationshipRiskSymbolSfxField) Tag() quickfix.Tag { return tag.RelationshipRiskSymbolSfx }

//NewRelationshipRiskSymbolSfx returns a new RelationshipRiskSymbolSfxField initialized with val
func NewRelationshipRiskSymbolSfx(val string) *RelationshipRiskSymbolSfxField {
	field := &RelationshipRiskSymbolSfxField{}
	field.Value = val
	return field
}

//RelationshipRiskWarningLevelNameField is a STRING field
type RelationshipRiskWarningLevelNameField struct{ quickfix.StringValue }

//Tag returns tag.RelationshipRiskWarningLevelName (1615)
func (f RelationshipRiskWarningLevelNameField) Tag() quickfix.Tag {
	return tag.RelationshipRiskWarningLevelName
}

//NewRelationshipRiskWarningLevelName returns a new RelationshipRiskWarningLevelNameField initialized with val
func NewRelationshipRiskWarningLevelName(val string) *RelationshipRiskWarningLevelNameField {
	field := &RelationshipRiskWarningLevelNameField{}
	field.Value = val
	return field
}

//RelationshipRiskWarningLevelPercentField is a PERCENTAGE field
type RelationshipRiskWarningLevelPercentField struct{ quickfix.PercentageValue }

//Tag returns tag.RelationshipRiskWarningLevelPercent (1614)
func (f RelationshipRiskWarningLevelPercentField) Tag() quickfix.Tag {
	return tag.RelationshipRiskWarningLevelPercent
}

//NewRelationshipRiskWarningLevelPercent returns a new RelationshipRiskWarningLevelPercentField initialized with val
func NewRelationshipRiskWarningLevelPercent(val float64) *RelationshipRiskWarningLevelPercentField {
	field := &RelationshipRiskWarningLevelPercentField{}
	field.Value = val
	return field
}

//RepoCollateralSecurityTypeField is a INT field
type RepoCollateralSecurityTypeField struct{ quickfix.IntValue }

//Tag returns tag.RepoCollateralSecurityType (239)
func (f RepoCollateralSecurityTypeField) Tag() quickfix.Tag { return tag.RepoCollateralSecurityType }

//NewRepoCollateralSecurityType returns a new RepoCollateralSecurityTypeField initialized with val
func NewRepoCollateralSecurityType(val int) *RepoCollateralSecurityTypeField {
	field := &RepoCollateralSecurityTypeField{}
	field.Value = val
	return field
}

//ReportToExchField is a BOOLEAN field
type ReportToExchField struct{ quickfix.BooleanValue }

//Tag returns tag.ReportToExch (113)
func (f ReportToExchField) Tag() quickfix.Tag { return tag.ReportToExch }

//NewReportToExch returns a new ReportToExchField initialized with val
func NewReportToExch(val bool) *ReportToExchField {
	field := &ReportToExchField{}
	field.Value = val
	return field
}

//ReportedPxField is a PRICE field
type ReportedPxField struct{ quickfix.PriceValue }

//Tag returns tag.ReportedPx (861)
func (f ReportedPxField) Tag() quickfix.Tag { return tag.ReportedPx }

//NewReportedPx returns a new ReportedPxField initialized with val
func NewReportedPx(val float64) *ReportedPxField {
	field := &ReportedPxField{}
	field.Value = val
	return field
}

//ReportedPxDiffField is a BOOLEAN field
type ReportedPxDiffField struct{ quickfix.BooleanValue }

//Tag returns tag.ReportedPxDiff (1134)
func (f ReportedPxDiffField) Tag() quickfix.Tag { return tag.ReportedPxDiff }

//NewReportedPxDiff returns a new ReportedPxDiffField initialized with val
func NewReportedPxDiff(val bool) *ReportedPxDiffField {
	field := &ReportedPxDiffField{}
	field.Value = val
	return field
}

//RepurchaseRateField is a PERCENTAGE field
type RepurchaseRateField struct{ quickfix.PercentageValue }

//Tag returns tag.RepurchaseRate (227)
func (f RepurchaseRateField) Tag() quickfix.Tag { return tag.RepurchaseRate }

//NewRepurchaseRate returns a new RepurchaseRateField initialized with val
func NewRepurchaseRate(val float64) *RepurchaseRateField {
	field := &RepurchaseRateField{}
	field.Value = val
	return field
}

//RepurchaseTermField is a INT field
type RepurchaseTermField struct{ quickfix.IntValue }

//Tag returns tag.RepurchaseTerm (226)
func (f RepurchaseTermField) Tag() quickfix.Tag { return tag.RepurchaseTerm }

//NewRepurchaseTerm returns a new RepurchaseTermField initialized with val
func NewRepurchaseTerm(val int) *RepurchaseTermField {
	field := &RepurchaseTermField{}
	field.Value = val
	return field
}

//RequestedPartyRoleField is a INT field
type RequestedPartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.RequestedPartyRole (1509)
func (f RequestedPartyRoleField) Tag() quickfix.Tag { return tag.RequestedPartyRole }

//NewRequestedPartyRole returns a new RequestedPartyRoleField initialized with val
func NewRequestedPartyRole(val int) *RequestedPartyRoleField {
	field := &RequestedPartyRoleField{}
	field.Value = val
	return field
}

//ResetSeqNumFlagField is a BOOLEAN field
type ResetSeqNumFlagField struct{ quickfix.BooleanValue }

//Tag returns tag.ResetSeqNumFlag (141)
func (f ResetSeqNumFlagField) Tag() quickfix.Tag { return tag.ResetSeqNumFlag }

//NewResetSeqNumFlag returns a new ResetSeqNumFlagField initialized with val
func NewResetSeqNumFlag(val bool) *ResetSeqNumFlagField {
	field := &ResetSeqNumFlagField{}
	field.Value = val
	return field
}

//RespondentTypeField is a INT field
type RespondentTypeField struct{ quickfix.IntValue }

//Tag returns tag.RespondentType (1172)
func (f RespondentTypeField) Tag() quickfix.Tag { return tag.RespondentType }

//NewRespondentType returns a new RespondentTypeField initialized with val
func NewRespondentType(val int) *RespondentTypeField {
	field := &RespondentTypeField{}
	field.Value = val
	return field
}

//ResponseDestinationField is a STRING field
type ResponseDestinationField struct{ quickfix.StringValue }

//Tag returns tag.ResponseDestination (726)
func (f ResponseDestinationField) Tag() quickfix.Tag { return tag.ResponseDestination }

//NewResponseDestination returns a new ResponseDestinationField initialized with val
func NewResponseDestination(val string) *ResponseDestinationField {
	field := &ResponseDestinationField{}
	field.Value = val
	return field
}

//ResponseTransportTypeField is a INT field
type ResponseTransportTypeField struct{ quickfix.IntValue }

//Tag returns tag.ResponseTransportType (725)
func (f ResponseTransportTypeField) Tag() quickfix.Tag { return tag.ResponseTransportType }

//NewResponseTransportType returns a new ResponseTransportTypeField initialized with val
func NewResponseTransportType(val int) *ResponseTransportTypeField {
	field := &ResponseTransportTypeField{}
	field.Value = val
	return field
}

//RestructuringTypeField is a STRING field
type RestructuringTypeField struct{ quickfix.StringValue }

//Tag returns tag.RestructuringType (1449)
func (f RestructuringTypeField) Tag() quickfix.Tag { return tag.RestructuringType }

//NewRestructuringType returns a new RestructuringTypeField initialized with val
func NewRestructuringType(val string) *RestructuringTypeField {
	field := &RestructuringTypeField{}
	field.Value = val
	return field
}

//ReversalIndicatorField is a BOOLEAN field
type ReversalIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.ReversalIndicator (700)
func (f ReversalIndicatorField) Tag() quickfix.Tag { return tag.ReversalIndicator }

//NewReversalIndicator returns a new ReversalIndicatorField initialized with val
func NewReversalIndicator(val bool) *ReversalIndicatorField {
	field := &ReversalIndicatorField{}
	field.Value = val
	return field
}

//RiskCFICodeField is a STRING field
type RiskCFICodeField struct{ quickfix.StringValue }

//Tag returns tag.RiskCFICode (1546)
func (f RiskCFICodeField) Tag() quickfix.Tag { return tag.RiskCFICode }

//NewRiskCFICode returns a new RiskCFICodeField initialized with val
func NewRiskCFICode(val string) *RiskCFICodeField {
	field := &RiskCFICodeField{}
	field.Value = val
	return field
}

//RiskCouponRateField is a PERCENTAGE field
type RiskCouponRateField struct{ quickfix.PercentageValue }

//Tag returns tag.RiskCouponRate (1555)
func (f RiskCouponRateField) Tag() quickfix.Tag { return tag.RiskCouponRate }

//NewRiskCouponRate returns a new RiskCouponRateField initialized with val
func NewRiskCouponRate(val float64) *RiskCouponRateField {
	field := &RiskCouponRateField{}
	field.Value = val
	return field
}

//RiskEncodedSecurityDescField is a DATA field
type RiskEncodedSecurityDescField struct{ quickfix.DataValue }

//Tag returns tag.RiskEncodedSecurityDesc (1621)
func (f RiskEncodedSecurityDescField) Tag() quickfix.Tag { return tag.RiskEncodedSecurityDesc }

//NewRiskEncodedSecurityDesc returns a new RiskEncodedSecurityDescField initialized with val
func NewRiskEncodedSecurityDesc(val string) *RiskEncodedSecurityDescField {
	field := &RiskEncodedSecurityDescField{}
	field.Value = val
	return field
}

//RiskEncodedSecurityDescLenField is a LENGTH field
type RiskEncodedSecurityDescLenField struct{ quickfix.LengthValue }

//Tag returns tag.RiskEncodedSecurityDescLen (1620)
func (f RiskEncodedSecurityDescLenField) Tag() quickfix.Tag { return tag.RiskEncodedSecurityDescLen }

//NewRiskEncodedSecurityDescLen returns a new RiskEncodedSecurityDescLenField initialized with val
func NewRiskEncodedSecurityDescLen(val int) *RiskEncodedSecurityDescLenField {
	field := &RiskEncodedSecurityDescLenField{}
	field.Value = val
	return field
}

//RiskFlexibleIndicatorField is a BOOLEAN field
type RiskFlexibleIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.RiskFlexibleIndicator (1554)
func (f RiskFlexibleIndicatorField) Tag() quickfix.Tag { return tag.RiskFlexibleIndicator }

//NewRiskFlexibleIndicator returns a new RiskFlexibleIndicatorField initialized with val
func NewRiskFlexibleIndicator(val bool) *RiskFlexibleIndicatorField {
	field := &RiskFlexibleIndicatorField{}
	field.Value = val
	return field
}

//RiskFreeRateField is a FLOAT field
type RiskFreeRateField struct{ quickfix.FloatValue }

//Tag returns tag.RiskFreeRate (1190)
func (f RiskFreeRateField) Tag() quickfix.Tag { return tag.RiskFreeRate }

//NewRiskFreeRate returns a new RiskFreeRateField initialized with val
func NewRiskFreeRate(val float64) *RiskFreeRateField {
	field := &RiskFreeRateField{}
	field.Value = val
	return field
}

//RiskInstrumentMultiplierField is a FLOAT field
type RiskInstrumentMultiplierField struct{ quickfix.FloatValue }

//Tag returns tag.RiskInstrumentMultiplier (1558)
func (f RiskInstrumentMultiplierField) Tag() quickfix.Tag { return tag.RiskInstrumentMultiplier }

//NewRiskInstrumentMultiplier returns a new RiskInstrumentMultiplierField initialized with val
func NewRiskInstrumentMultiplier(val float64) *RiskInstrumentMultiplierField {
	field := &RiskInstrumentMultiplierField{}
	field.Value = val
	return field
}

//RiskInstrumentOperatorField is a INT field
type RiskInstrumentOperatorField struct{ quickfix.IntValue }

//Tag returns tag.RiskInstrumentOperator (1535)
func (f RiskInstrumentOperatorField) Tag() quickfix.Tag { return tag.RiskInstrumentOperator }

//NewRiskInstrumentOperator returns a new RiskInstrumentOperatorField initialized with val
func NewRiskInstrumentOperator(val int) *RiskInstrumentOperatorField {
	field := &RiskInstrumentOperatorField{}
	field.Value = val
	return field
}

//RiskInstrumentSettlTypeField is a STRING field
type RiskInstrumentSettlTypeField struct{ quickfix.StringValue }

//Tag returns tag.RiskInstrumentSettlType (1557)
func (f RiskInstrumentSettlTypeField) Tag() quickfix.Tag { return tag.RiskInstrumentSettlType }

//NewRiskInstrumentSettlType returns a new RiskInstrumentSettlTypeField initialized with val
func NewRiskInstrumentSettlType(val string) *RiskInstrumentSettlTypeField {
	field := &RiskInstrumentSettlTypeField{}
	field.Value = val
	return field
}

//RiskLimitAmountField is a AMT field
type RiskLimitAmountField struct{ quickfix.AmtValue }

//Tag returns tag.RiskLimitAmount (1531)
func (f RiskLimitAmountField) Tag() quickfix.Tag { return tag.RiskLimitAmount }

//NewRiskLimitAmount returns a new RiskLimitAmountField initialized with val
func NewRiskLimitAmount(val float64) *RiskLimitAmountField {
	field := &RiskLimitAmountField{}
	field.Value = val
	return field
}

//RiskLimitCurrencyField is a CURRENCY field
type RiskLimitCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.RiskLimitCurrency (1532)
func (f RiskLimitCurrencyField) Tag() quickfix.Tag { return tag.RiskLimitCurrency }

//NewRiskLimitCurrency returns a new RiskLimitCurrencyField initialized with val
func NewRiskLimitCurrency(val string) *RiskLimitCurrencyField {
	field := &RiskLimitCurrencyField{}
	field.Value = val
	return field
}

//RiskLimitPlatformField is a STRING field
type RiskLimitPlatformField struct{ quickfix.StringValue }

//Tag returns tag.RiskLimitPlatform (1533)
func (f RiskLimitPlatformField) Tag() quickfix.Tag { return tag.RiskLimitPlatform }

//NewRiskLimitPlatform returns a new RiskLimitPlatformField initialized with val
func NewRiskLimitPlatform(val string) *RiskLimitPlatformField {
	field := &RiskLimitPlatformField{}
	field.Value = val
	return field
}

//RiskLimitTypeField is a INT field
type RiskLimitTypeField struct{ quickfix.IntValue }

//Tag returns tag.RiskLimitType (1530)
func (f RiskLimitTypeField) Tag() quickfix.Tag { return tag.RiskLimitType }

//NewRiskLimitType returns a new RiskLimitTypeField initialized with val
func NewRiskLimitType(val int) *RiskLimitTypeField {
	field := &RiskLimitTypeField{}
	field.Value = val
	return field
}

//RiskMaturityMonthYearField is a MONTHYEAR field
type RiskMaturityMonthYearField struct{ quickfix.MonthYearValue }

//Tag returns tag.RiskMaturityMonthYear (1549)
func (f RiskMaturityMonthYearField) Tag() quickfix.Tag { return tag.RiskMaturityMonthYear }

//NewRiskMaturityMonthYear returns a new RiskMaturityMonthYearField initialized with val
func NewRiskMaturityMonthYear(val string) *RiskMaturityMonthYearField {
	field := &RiskMaturityMonthYearField{}
	field.Value = val
	return field
}

//RiskMaturityTimeField is a TZTIMEONLY field
type RiskMaturityTimeField struct{ quickfix.TZTimeOnlyValue }

//Tag returns tag.RiskMaturityTime (1550)
func (f RiskMaturityTimeField) Tag() quickfix.Tag { return tag.RiskMaturityTime }

//RiskProductField is a INT field
type RiskProductField struct{ quickfix.IntValue }

//Tag returns tag.RiskProduct (1543)
func (f RiskProductField) Tag() quickfix.Tag { return tag.RiskProduct }

//NewRiskProduct returns a new RiskProductField initialized with val
func NewRiskProduct(val int) *RiskProductField {
	field := &RiskProductField{}
	field.Value = val
	return field
}

//RiskProductComplexField is a STRING field
type RiskProductComplexField struct{ quickfix.StringValue }

//Tag returns tag.RiskProductComplex (1544)
func (f RiskProductComplexField) Tag() quickfix.Tag { return tag.RiskProductComplex }

//NewRiskProductComplex returns a new RiskProductComplexField initialized with val
func NewRiskProductComplex(val string) *RiskProductComplexField {
	field := &RiskProductComplexField{}
	field.Value = val
	return field
}

//RiskPutOrCallField is a INT field
type RiskPutOrCallField struct{ quickfix.IntValue }

//Tag returns tag.RiskPutOrCall (1553)
func (f RiskPutOrCallField) Tag() quickfix.Tag { return tag.RiskPutOrCall }

//NewRiskPutOrCall returns a new RiskPutOrCallField initialized with val
func NewRiskPutOrCall(val int) *RiskPutOrCallField {
	field := &RiskPutOrCallField{}
	field.Value = val
	return field
}

//RiskRestructuringTypeField is a STRING field
type RiskRestructuringTypeField struct{ quickfix.StringValue }

//Tag returns tag.RiskRestructuringType (1551)
func (f RiskRestructuringTypeField) Tag() quickfix.Tag { return tag.RiskRestructuringType }

//NewRiskRestructuringType returns a new RiskRestructuringTypeField initialized with val
func NewRiskRestructuringType(val string) *RiskRestructuringTypeField {
	field := &RiskRestructuringTypeField{}
	field.Value = val
	return field
}

//RiskSecurityAltIDField is a STRING field
type RiskSecurityAltIDField struct{ quickfix.StringValue }

//Tag returns tag.RiskSecurityAltID (1541)
func (f RiskSecurityAltIDField) Tag() quickfix.Tag { return tag.RiskSecurityAltID }

//NewRiskSecurityAltID returns a new RiskSecurityAltIDField initialized with val
func NewRiskSecurityAltID(val string) *RiskSecurityAltIDField {
	field := &RiskSecurityAltIDField{}
	field.Value = val
	return field
}

//RiskSecurityAltIDSourceField is a STRING field
type RiskSecurityAltIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.RiskSecurityAltIDSource (1542)
func (f RiskSecurityAltIDSourceField) Tag() quickfix.Tag { return tag.RiskSecurityAltIDSource }

//NewRiskSecurityAltIDSource returns a new RiskSecurityAltIDSourceField initialized with val
func NewRiskSecurityAltIDSource(val string) *RiskSecurityAltIDSourceField {
	field := &RiskSecurityAltIDSourceField{}
	field.Value = val
	return field
}

//RiskSecurityDescField is a STRING field
type RiskSecurityDescField struct{ quickfix.StringValue }

//Tag returns tag.RiskSecurityDesc (1556)
func (f RiskSecurityDescField) Tag() quickfix.Tag { return tag.RiskSecurityDesc }

//NewRiskSecurityDesc returns a new RiskSecurityDescField initialized with val
func NewRiskSecurityDesc(val string) *RiskSecurityDescField {
	field := &RiskSecurityDescField{}
	field.Value = val
	return field
}

//RiskSecurityExchangeField is a EXCHANGE field
type RiskSecurityExchangeField struct{ quickfix.ExchangeValue }

//Tag returns tag.RiskSecurityExchange (1616)
func (f RiskSecurityExchangeField) Tag() quickfix.Tag { return tag.RiskSecurityExchange }

//NewRiskSecurityExchange returns a new RiskSecurityExchangeField initialized with val
func NewRiskSecurityExchange(val string) *RiskSecurityExchangeField {
	field := &RiskSecurityExchangeField{}
	field.Value = val
	return field
}

//RiskSecurityGroupField is a STRING field
type RiskSecurityGroupField struct{ quickfix.StringValue }

//Tag returns tag.RiskSecurityGroup (1545)
func (f RiskSecurityGroupField) Tag() quickfix.Tag { return tag.RiskSecurityGroup }

//NewRiskSecurityGroup returns a new RiskSecurityGroupField initialized with val
func NewRiskSecurityGroup(val string) *RiskSecurityGroupField {
	field := &RiskSecurityGroupField{}
	field.Value = val
	return field
}

//RiskSecurityIDField is a STRING field
type RiskSecurityIDField struct{ quickfix.StringValue }

//Tag returns tag.RiskSecurityID (1538)
func (f RiskSecurityIDField) Tag() quickfix.Tag { return tag.RiskSecurityID }

//NewRiskSecurityID returns a new RiskSecurityIDField initialized with val
func NewRiskSecurityID(val string) *RiskSecurityIDField {
	field := &RiskSecurityIDField{}
	field.Value = val
	return field
}

//RiskSecurityIDSourceField is a STRING field
type RiskSecurityIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.RiskSecurityIDSource (1539)
func (f RiskSecurityIDSourceField) Tag() quickfix.Tag { return tag.RiskSecurityIDSource }

//NewRiskSecurityIDSource returns a new RiskSecurityIDSourceField initialized with val
func NewRiskSecurityIDSource(val string) *RiskSecurityIDSourceField {
	field := &RiskSecurityIDSourceField{}
	field.Value = val
	return field
}

//RiskSecuritySubTypeField is a STRING field
type RiskSecuritySubTypeField struct{ quickfix.StringValue }

//Tag returns tag.RiskSecuritySubType (1548)
func (f RiskSecuritySubTypeField) Tag() quickfix.Tag { return tag.RiskSecuritySubType }

//NewRiskSecuritySubType returns a new RiskSecuritySubTypeField initialized with val
func NewRiskSecuritySubType(val string) *RiskSecuritySubTypeField {
	field := &RiskSecuritySubTypeField{}
	field.Value = val
	return field
}

//RiskSecurityTypeField is a STRING field
type RiskSecurityTypeField struct{ quickfix.StringValue }

//Tag returns tag.RiskSecurityType (1547)
func (f RiskSecurityTypeField) Tag() quickfix.Tag { return tag.RiskSecurityType }

//NewRiskSecurityType returns a new RiskSecurityTypeField initialized with val
func NewRiskSecurityType(val string) *RiskSecurityTypeField {
	field := &RiskSecurityTypeField{}
	field.Value = val
	return field
}

//RiskSeniorityField is a STRING field
type RiskSeniorityField struct{ quickfix.StringValue }

//Tag returns tag.RiskSeniority (1552)
func (f RiskSeniorityField) Tag() quickfix.Tag { return tag.RiskSeniority }

//NewRiskSeniority returns a new RiskSeniorityField initialized with val
func NewRiskSeniority(val string) *RiskSeniorityField {
	field := &RiskSeniorityField{}
	field.Value = val
	return field
}

//RiskSymbolField is a STRING field
type RiskSymbolField struct{ quickfix.StringValue }

//Tag returns tag.RiskSymbol (1536)
func (f RiskSymbolField) Tag() quickfix.Tag { return tag.RiskSymbol }

//NewRiskSymbol returns a new RiskSymbolField initialized with val
func NewRiskSymbol(val string) *RiskSymbolField {
	field := &RiskSymbolField{}
	field.Value = val
	return field
}

//RiskSymbolSfxField is a STRING field
type RiskSymbolSfxField struct{ quickfix.StringValue }

//Tag returns tag.RiskSymbolSfx (1537)
func (f RiskSymbolSfxField) Tag() quickfix.Tag { return tag.RiskSymbolSfx }

//NewRiskSymbolSfx returns a new RiskSymbolSfxField initialized with val
func NewRiskSymbolSfx(val string) *RiskSymbolSfxField {
	field := &RiskSymbolSfxField{}
	field.Value = val
	return field
}

//RiskWarningLevelNameField is a STRING field
type RiskWarningLevelNameField struct{ quickfix.StringValue }

//Tag returns tag.RiskWarningLevelName (1561)
func (f RiskWarningLevelNameField) Tag() quickfix.Tag { return tag.RiskWarningLevelName }

//NewRiskWarningLevelName returns a new RiskWarningLevelNameField initialized with val
func NewRiskWarningLevelName(val string) *RiskWarningLevelNameField {
	field := &RiskWarningLevelNameField{}
	field.Value = val
	return field
}

//RiskWarningLevelPercentField is a PERCENTAGE field
type RiskWarningLevelPercentField struct{ quickfix.PercentageValue }

//Tag returns tag.RiskWarningLevelPercent (1560)
func (f RiskWarningLevelPercentField) Tag() quickfix.Tag { return tag.RiskWarningLevelPercent }

//NewRiskWarningLevelPercent returns a new RiskWarningLevelPercentField initialized with val
func NewRiskWarningLevelPercent(val float64) *RiskWarningLevelPercentField {
	field := &RiskWarningLevelPercentField{}
	field.Value = val
	return field
}

//RndPxField is a PRICE field
type RndPxField struct{ quickfix.PriceValue }

//Tag returns tag.RndPx (991)
func (f RndPxField) Tag() quickfix.Tag { return tag.RndPx }

//NewRndPx returns a new RndPxField initialized with val
func NewRndPx(val float64) *RndPxField {
	field := &RndPxField{}
	field.Value = val
	return field
}

//RootPartyIDField is a STRING field
type RootPartyIDField struct{ quickfix.StringValue }

//Tag returns tag.RootPartyID (1117)
func (f RootPartyIDField) Tag() quickfix.Tag { return tag.RootPartyID }

//NewRootPartyID returns a new RootPartyIDField initialized with val
func NewRootPartyID(val string) *RootPartyIDField {
	field := &RootPartyIDField{}
	field.Value = val
	return field
}

//RootPartyIDSourceField is a CHAR field
type RootPartyIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.RootPartyIDSource (1118)
func (f RootPartyIDSourceField) Tag() quickfix.Tag { return tag.RootPartyIDSource }

//NewRootPartyIDSource returns a new RootPartyIDSourceField initialized with val
func NewRootPartyIDSource(val string) *RootPartyIDSourceField {
	field := &RootPartyIDSourceField{}
	field.Value = val
	return field
}

//RootPartyRoleField is a INT field
type RootPartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.RootPartyRole (1119)
func (f RootPartyRoleField) Tag() quickfix.Tag { return tag.RootPartyRole }

//NewRootPartyRole returns a new RootPartyRoleField initialized with val
func NewRootPartyRole(val int) *RootPartyRoleField {
	field := &RootPartyRoleField{}
	field.Value = val
	return field
}

//RootPartySubIDField is a STRING field
type RootPartySubIDField struct{ quickfix.StringValue }

//Tag returns tag.RootPartySubID (1121)
func (f RootPartySubIDField) Tag() quickfix.Tag { return tag.RootPartySubID }

//NewRootPartySubID returns a new RootPartySubIDField initialized with val
func NewRootPartySubID(val string) *RootPartySubIDField {
	field := &RootPartySubIDField{}
	field.Value = val
	return field
}

//RootPartySubIDTypeField is a INT field
type RootPartySubIDTypeField struct{ quickfix.IntValue }

//Tag returns tag.RootPartySubIDType (1122)
func (f RootPartySubIDTypeField) Tag() quickfix.Tag { return tag.RootPartySubIDType }

//NewRootPartySubIDType returns a new RootPartySubIDTypeField initialized with val
func NewRootPartySubIDType(val int) *RootPartySubIDTypeField {
	field := &RootPartySubIDTypeField{}
	field.Value = val
	return field
}

//RoundLotField is a QTY field
type RoundLotField struct{ quickfix.QtyValue }

//Tag returns tag.RoundLot (561)
func (f RoundLotField) Tag() quickfix.Tag { return tag.RoundLot }

//NewRoundLot returns a new RoundLotField initialized with val
func NewRoundLot(val float64) *RoundLotField {
	field := &RoundLotField{}
	field.Value = val
	return field
}

//RoundingDirectionField is a CHAR field
type RoundingDirectionField struct{ quickfix.CharValue }

//Tag returns tag.RoundingDirection (468)
func (f RoundingDirectionField) Tag() quickfix.Tag { return tag.RoundingDirection }

//NewRoundingDirection returns a new RoundingDirectionField initialized with val
func NewRoundingDirection(val string) *RoundingDirectionField {
	field := &RoundingDirectionField{}
	field.Value = val
	return field
}

//RoundingModulusField is a FLOAT field
type RoundingModulusField struct{ quickfix.FloatValue }

//Tag returns tag.RoundingModulus (469)
func (f RoundingModulusField) Tag() quickfix.Tag { return tag.RoundingModulus }

//NewRoundingModulus returns a new RoundingModulusField initialized with val
func NewRoundingModulus(val float64) *RoundingModulusField {
	field := &RoundingModulusField{}
	field.Value = val
	return field
}

//RoutingIDField is a STRING field
type RoutingIDField struct{ quickfix.StringValue }

//Tag returns tag.RoutingID (217)
func (f RoutingIDField) Tag() quickfix.Tag { return tag.RoutingID }

//NewRoutingID returns a new RoutingIDField initialized with val
func NewRoutingID(val string) *RoutingIDField {
	field := &RoutingIDField{}
	field.Value = val
	return field
}

//RoutingTypeField is a INT field
type RoutingTypeField struct{ quickfix.IntValue }

//Tag returns tag.RoutingType (216)
func (f RoutingTypeField) Tag() quickfix.Tag { return tag.RoutingType }

//NewRoutingType returns a new RoutingTypeField initialized with val
func NewRoutingType(val int) *RoutingTypeField {
	field := &RoutingTypeField{}
	field.Value = val
	return field
}

//RptSeqField is a INT field
type RptSeqField struct{ quickfix.IntValue }

//Tag returns tag.RptSeq (83)
func (f RptSeqField) Tag() quickfix.Tag { return tag.RptSeq }

//NewRptSeq returns a new RptSeqField initialized with val
func NewRptSeq(val int) *RptSeqField {
	field := &RptSeqField{}
	field.Value = val
	return field
}

//RptSysField is a STRING field
type RptSysField struct{ quickfix.StringValue }

//Tag returns tag.RptSys (1135)
func (f RptSysField) Tag() quickfix.Tag { return tag.RptSys }

//NewRptSys returns a new RptSysField initialized with val
func NewRptSys(val string) *RptSysField {
	field := &RptSysField{}
	field.Value = val
	return field
}

//Rule80AField is a CHAR field
type Rule80AField struct{ quickfix.CharValue }

//Tag returns tag.Rule80A (47)
func (f Rule80AField) Tag() quickfix.Tag { return tag.Rule80A }

//NewRule80A returns a new Rule80AField initialized with val
func NewRule80A(val string) *Rule80AField {
	field := &Rule80AField{}
	field.Value = val
	return field
}

//ScopeField is a MULTIPLECHARVALUE field
type ScopeField struct{ quickfix.MultipleCharValue }

//Tag returns tag.Scope (546)
func (f ScopeField) Tag() quickfix.Tag { return tag.Scope }

//NewScope returns a new ScopeField initialized with val
func NewScope(val string) *ScopeField {
	field := &ScopeField{}
	field.Value = val
	return field
}

//SecDefStatusField is a INT field
type SecDefStatusField struct{ quickfix.IntValue }

//Tag returns tag.SecDefStatus (653)
func (f SecDefStatusField) Tag() quickfix.Tag { return tag.SecDefStatus }

//NewSecDefStatus returns a new SecDefStatusField initialized with val
func NewSecDefStatus(val int) *SecDefStatusField {
	field := &SecDefStatusField{}
	field.Value = val
	return field
}

//SecondaryAllocIDField is a STRING field
type SecondaryAllocIDField struct{ quickfix.StringValue }

//Tag returns tag.SecondaryAllocID (793)
func (f SecondaryAllocIDField) Tag() quickfix.Tag { return tag.SecondaryAllocID }

//NewSecondaryAllocID returns a new SecondaryAllocIDField initialized with val
func NewSecondaryAllocID(val string) *SecondaryAllocIDField {
	field := &SecondaryAllocIDField{}
	field.Value = val
	return field
}

//SecondaryClOrdIDField is a STRING field
type SecondaryClOrdIDField struct{ quickfix.StringValue }

//Tag returns tag.SecondaryClOrdID (526)
func (f SecondaryClOrdIDField) Tag() quickfix.Tag { return tag.SecondaryClOrdID }

//NewSecondaryClOrdID returns a new SecondaryClOrdIDField initialized with val
func NewSecondaryClOrdID(val string) *SecondaryClOrdIDField {
	field := &SecondaryClOrdIDField{}
	field.Value = val
	return field
}

//SecondaryDisplayQtyField is a QTY field
type SecondaryDisplayQtyField struct{ quickfix.QtyValue }

//Tag returns tag.SecondaryDisplayQty (1082)
func (f SecondaryDisplayQtyField) Tag() quickfix.Tag { return tag.SecondaryDisplayQty }

//NewSecondaryDisplayQty returns a new SecondaryDisplayQtyField initialized with val
func NewSecondaryDisplayQty(val float64) *SecondaryDisplayQtyField {
	field := &SecondaryDisplayQtyField{}
	field.Value = val
	return field
}

//SecondaryExecIDField is a STRING field
type SecondaryExecIDField struct{ quickfix.StringValue }

//Tag returns tag.SecondaryExecID (527)
func (f SecondaryExecIDField) Tag() quickfix.Tag { return tag.SecondaryExecID }

//NewSecondaryExecID returns a new SecondaryExecIDField initialized with val
func NewSecondaryExecID(val string) *SecondaryExecIDField {
	field := &SecondaryExecIDField{}
	field.Value = val
	return field
}

//SecondaryFirmTradeIDField is a STRING field
type SecondaryFirmTradeIDField struct{ quickfix.StringValue }

//Tag returns tag.SecondaryFirmTradeID (1042)
func (f SecondaryFirmTradeIDField) Tag() quickfix.Tag { return tag.SecondaryFirmTradeID }

//NewSecondaryFirmTradeID returns a new SecondaryFirmTradeIDField initialized with val
func NewSecondaryFirmTradeID(val string) *SecondaryFirmTradeIDField {
	field := &SecondaryFirmTradeIDField{}
	field.Value = val
	return field
}

//SecondaryHighLimitPriceField is a PRICE field
type SecondaryHighLimitPriceField struct{ quickfix.PriceValue }

//Tag returns tag.SecondaryHighLimitPrice (1230)
func (f SecondaryHighLimitPriceField) Tag() quickfix.Tag { return tag.SecondaryHighLimitPrice }

//NewSecondaryHighLimitPrice returns a new SecondaryHighLimitPriceField initialized with val
func NewSecondaryHighLimitPrice(val float64) *SecondaryHighLimitPriceField {
	field := &SecondaryHighLimitPriceField{}
	field.Value = val
	return field
}

//SecondaryIndividualAllocIDField is a STRING field
type SecondaryIndividualAllocIDField struct{ quickfix.StringValue }

//Tag returns tag.SecondaryIndividualAllocID (989)
func (f SecondaryIndividualAllocIDField) Tag() quickfix.Tag { return tag.SecondaryIndividualAllocID }

//NewSecondaryIndividualAllocID returns a new SecondaryIndividualAllocIDField initialized with val
func NewSecondaryIndividualAllocID(val string) *SecondaryIndividualAllocIDField {
	field := &SecondaryIndividualAllocIDField{}
	field.Value = val
	return field
}

//SecondaryLowLimitPriceField is a PRICE field
type SecondaryLowLimitPriceField struct{ quickfix.PriceValue }

//Tag returns tag.SecondaryLowLimitPrice (1221)
func (f SecondaryLowLimitPriceField) Tag() quickfix.Tag { return tag.SecondaryLowLimitPrice }

//NewSecondaryLowLimitPrice returns a new SecondaryLowLimitPriceField initialized with val
func NewSecondaryLowLimitPrice(val float64) *SecondaryLowLimitPriceField {
	field := &SecondaryLowLimitPriceField{}
	field.Value = val
	return field
}

//SecondaryOrderIDField is a STRING field
type SecondaryOrderIDField struct{ quickfix.StringValue }

//Tag returns tag.SecondaryOrderID (198)
func (f SecondaryOrderIDField) Tag() quickfix.Tag { return tag.SecondaryOrderID }

//NewSecondaryOrderID returns a new SecondaryOrderIDField initialized with val
func NewSecondaryOrderID(val string) *SecondaryOrderIDField {
	field := &SecondaryOrderIDField{}
	field.Value = val
	return field
}

//SecondaryPriceLimitTypeField is a INT field
type SecondaryPriceLimitTypeField struct{ quickfix.IntValue }

//Tag returns tag.SecondaryPriceLimitType (1305)
func (f SecondaryPriceLimitTypeField) Tag() quickfix.Tag { return tag.SecondaryPriceLimitType }

//NewSecondaryPriceLimitType returns a new SecondaryPriceLimitTypeField initialized with val
func NewSecondaryPriceLimitType(val int) *SecondaryPriceLimitTypeField {
	field := &SecondaryPriceLimitTypeField{}
	field.Value = val
	return field
}

//SecondaryTradeIDField is a STRING field
type SecondaryTradeIDField struct{ quickfix.StringValue }

//Tag returns tag.SecondaryTradeID (1040)
func (f SecondaryTradeIDField) Tag() quickfix.Tag { return tag.SecondaryTradeID }

//NewSecondaryTradeID returns a new SecondaryTradeIDField initialized with val
func NewSecondaryTradeID(val string) *SecondaryTradeIDField {
	field := &SecondaryTradeIDField{}
	field.Value = val
	return field
}

//SecondaryTradeReportIDField is a STRING field
type SecondaryTradeReportIDField struct{ quickfix.StringValue }

//Tag returns tag.SecondaryTradeReportID (818)
func (f SecondaryTradeReportIDField) Tag() quickfix.Tag { return tag.SecondaryTradeReportID }

//NewSecondaryTradeReportID returns a new SecondaryTradeReportIDField initialized with val
func NewSecondaryTradeReportID(val string) *SecondaryTradeReportIDField {
	field := &SecondaryTradeReportIDField{}
	field.Value = val
	return field
}

//SecondaryTradeReportRefIDField is a STRING field
type SecondaryTradeReportRefIDField struct{ quickfix.StringValue }

//Tag returns tag.SecondaryTradeReportRefID (881)
func (f SecondaryTradeReportRefIDField) Tag() quickfix.Tag { return tag.SecondaryTradeReportRefID }

//NewSecondaryTradeReportRefID returns a new SecondaryTradeReportRefIDField initialized with val
func NewSecondaryTradeReportRefID(val string) *SecondaryTradeReportRefIDField {
	field := &SecondaryTradeReportRefIDField{}
	field.Value = val
	return field
}

//SecondaryTradingReferencePriceField is a PRICE field
type SecondaryTradingReferencePriceField struct{ quickfix.PriceValue }

//Tag returns tag.SecondaryTradingReferencePrice (1240)
func (f SecondaryTradingReferencePriceField) Tag() quickfix.Tag {
	return tag.SecondaryTradingReferencePrice
}

//NewSecondaryTradingReferencePrice returns a new SecondaryTradingReferencePriceField initialized with val
func NewSecondaryTradingReferencePrice(val float64) *SecondaryTradingReferencePriceField {
	field := &SecondaryTradingReferencePriceField{}
	field.Value = val
	return field
}

//SecondaryTrdTypeField is a INT field
type SecondaryTrdTypeField struct{ quickfix.IntValue }

//Tag returns tag.SecondaryTrdType (855)
func (f SecondaryTrdTypeField) Tag() quickfix.Tag { return tag.SecondaryTrdType }

//NewSecondaryTrdType returns a new SecondaryTrdTypeField initialized with val
func NewSecondaryTrdType(val int) *SecondaryTrdTypeField {
	field := &SecondaryTrdTypeField{}
	field.Value = val
	return field
}

//SecureDataField is a DATA field
type SecureDataField struct{ quickfix.DataValue }

//Tag returns tag.SecureData (91)
func (f SecureDataField) Tag() quickfix.Tag { return tag.SecureData }

//NewSecureData returns a new SecureDataField initialized with val
func NewSecureData(val string) *SecureDataField {
	field := &SecureDataField{}
	field.Value = val
	return field
}

//SecureDataLenField is a LENGTH field
type SecureDataLenField struct{ quickfix.LengthValue }

//Tag returns tag.SecureDataLen (90)
func (f SecureDataLenField) Tag() quickfix.Tag { return tag.SecureDataLen }

//NewSecureDataLen returns a new SecureDataLenField initialized with val
func NewSecureDataLen(val int) *SecureDataLenField {
	field := &SecureDataLenField{}
	field.Value = val
	return field
}

//SecurityAltIDField is a STRING field
type SecurityAltIDField struct{ quickfix.StringValue }

//Tag returns tag.SecurityAltID (455)
func (f SecurityAltIDField) Tag() quickfix.Tag { return tag.SecurityAltID }

//NewSecurityAltID returns a new SecurityAltIDField initialized with val
func NewSecurityAltID(val string) *SecurityAltIDField {
	field := &SecurityAltIDField{}
	field.Value = val
	return field
}

//SecurityAltIDSourceField is a STRING field
type SecurityAltIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.SecurityAltIDSource (456)
func (f SecurityAltIDSourceField) Tag() quickfix.Tag { return tag.SecurityAltIDSource }

//NewSecurityAltIDSource returns a new SecurityAltIDSourceField initialized with val
func NewSecurityAltIDSource(val string) *SecurityAltIDSourceField {
	field := &SecurityAltIDSourceField{}
	field.Value = val
	return field
}

//SecurityDescField is a STRING field
type SecurityDescField struct{ quickfix.StringValue }

//Tag returns tag.SecurityDesc (107)
func (f SecurityDescField) Tag() quickfix.Tag { return tag.SecurityDesc }

//NewSecurityDesc returns a new SecurityDescField initialized with val
func NewSecurityDesc(val string) *SecurityDescField {
	field := &SecurityDescField{}
	field.Value = val
	return field
}

//SecurityExchangeField is a EXCHANGE field
type SecurityExchangeField struct{ quickfix.ExchangeValue }

//Tag returns tag.SecurityExchange (207)
func (f SecurityExchangeField) Tag() quickfix.Tag { return tag.SecurityExchange }

//NewSecurityExchange returns a new SecurityExchangeField initialized with val
func NewSecurityExchange(val string) *SecurityExchangeField {
	field := &SecurityExchangeField{}
	field.Value = val
	return field
}

//SecurityGroupField is a STRING field
type SecurityGroupField struct{ quickfix.StringValue }

//Tag returns tag.SecurityGroup (1151)
func (f SecurityGroupField) Tag() quickfix.Tag { return tag.SecurityGroup }

//NewSecurityGroup returns a new SecurityGroupField initialized with val
func NewSecurityGroup(val string) *SecurityGroupField {
	field := &SecurityGroupField{}
	field.Value = val
	return field
}

//SecurityIDField is a STRING field
type SecurityIDField struct{ quickfix.StringValue }

//Tag returns tag.SecurityID (48)
func (f SecurityIDField) Tag() quickfix.Tag { return tag.SecurityID }

//NewSecurityID returns a new SecurityIDField initialized with val
func NewSecurityID(val string) *SecurityIDField {
	field := &SecurityIDField{}
	field.Value = val
	return field
}

//SecurityIDSourceField is a STRING field
type SecurityIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.SecurityIDSource (22)
func (f SecurityIDSourceField) Tag() quickfix.Tag { return tag.SecurityIDSource }

//NewSecurityIDSource returns a new SecurityIDSourceField initialized with val
func NewSecurityIDSource(val string) *SecurityIDSourceField {
	field := &SecurityIDSourceField{}
	field.Value = val
	return field
}

//SecurityListDescField is a STRING field
type SecurityListDescField struct{ quickfix.StringValue }

//Tag returns tag.SecurityListDesc (1467)
func (f SecurityListDescField) Tag() quickfix.Tag { return tag.SecurityListDesc }

//NewSecurityListDesc returns a new SecurityListDescField initialized with val
func NewSecurityListDesc(val string) *SecurityListDescField {
	field := &SecurityListDescField{}
	field.Value = val
	return field
}

//SecurityListIDField is a STRING field
type SecurityListIDField struct{ quickfix.StringValue }

//Tag returns tag.SecurityListID (1465)
func (f SecurityListIDField) Tag() quickfix.Tag { return tag.SecurityListID }

//NewSecurityListID returns a new SecurityListIDField initialized with val
func NewSecurityListID(val string) *SecurityListIDField {
	field := &SecurityListIDField{}
	field.Value = val
	return field
}

//SecurityListRefIDField is a STRING field
type SecurityListRefIDField struct{ quickfix.StringValue }

//Tag returns tag.SecurityListRefID (1466)
func (f SecurityListRefIDField) Tag() quickfix.Tag { return tag.SecurityListRefID }

//NewSecurityListRefID returns a new SecurityListRefIDField initialized with val
func NewSecurityListRefID(val string) *SecurityListRefIDField {
	field := &SecurityListRefIDField{}
	field.Value = val
	return field
}

//SecurityListRequestTypeField is a INT field
type SecurityListRequestTypeField struct{ quickfix.IntValue }

//Tag returns tag.SecurityListRequestType (559)
func (f SecurityListRequestTypeField) Tag() quickfix.Tag { return tag.SecurityListRequestType }

//NewSecurityListRequestType returns a new SecurityListRequestTypeField initialized with val
func NewSecurityListRequestType(val int) *SecurityListRequestTypeField {
	field := &SecurityListRequestTypeField{}
	field.Value = val
	return field
}

//SecurityListTypeField is a INT field
type SecurityListTypeField struct{ quickfix.IntValue }

//Tag returns tag.SecurityListType (1470)
func (f SecurityListTypeField) Tag() quickfix.Tag { return tag.SecurityListType }

//NewSecurityListType returns a new SecurityListTypeField initialized with val
func NewSecurityListType(val int) *SecurityListTypeField {
	field := &SecurityListTypeField{}
	field.Value = val
	return field
}

//SecurityListTypeSourceField is a INT field
type SecurityListTypeSourceField struct{ quickfix.IntValue }

//Tag returns tag.SecurityListTypeSource (1471)
func (f SecurityListTypeSourceField) Tag() quickfix.Tag { return tag.SecurityListTypeSource }

//NewSecurityListTypeSource returns a new SecurityListTypeSourceField initialized with val
func NewSecurityListTypeSource(val int) *SecurityListTypeSourceField {
	field := &SecurityListTypeSourceField{}
	field.Value = val
	return field
}

//SecurityReportIDField is a INT field
type SecurityReportIDField struct{ quickfix.IntValue }

//Tag returns tag.SecurityReportID (964)
func (f SecurityReportIDField) Tag() quickfix.Tag { return tag.SecurityReportID }

//NewSecurityReportID returns a new SecurityReportIDField initialized with val
func NewSecurityReportID(val int) *SecurityReportIDField {
	field := &SecurityReportIDField{}
	field.Value = val
	return field
}

//SecurityReqIDField is a STRING field
type SecurityReqIDField struct{ quickfix.StringValue }

//Tag returns tag.SecurityReqID (320)
func (f SecurityReqIDField) Tag() quickfix.Tag { return tag.SecurityReqID }

//NewSecurityReqID returns a new SecurityReqIDField initialized with val
func NewSecurityReqID(val string) *SecurityReqIDField {
	field := &SecurityReqIDField{}
	field.Value = val
	return field
}

//SecurityRequestResultField is a INT field
type SecurityRequestResultField struct{ quickfix.IntValue }

//Tag returns tag.SecurityRequestResult (560)
func (f SecurityRequestResultField) Tag() quickfix.Tag { return tag.SecurityRequestResult }

//NewSecurityRequestResult returns a new SecurityRequestResultField initialized with val
func NewSecurityRequestResult(val int) *SecurityRequestResultField {
	field := &SecurityRequestResultField{}
	field.Value = val
	return field
}

//SecurityRequestTypeField is a INT field
type SecurityRequestTypeField struct{ quickfix.IntValue }

//Tag returns tag.SecurityRequestType (321)
func (f SecurityRequestTypeField) Tag() quickfix.Tag { return tag.SecurityRequestType }

//NewSecurityRequestType returns a new SecurityRequestTypeField initialized with val
func NewSecurityRequestType(val int) *SecurityRequestTypeField {
	field := &SecurityRequestTypeField{}
	field.Value = val
	return field
}

//SecurityResponseIDField is a STRING field
type SecurityResponseIDField struct{ quickfix.StringValue }

//Tag returns tag.SecurityResponseID (322)
func (f SecurityResponseIDField) Tag() quickfix.Tag { return tag.SecurityResponseID }

//NewSecurityResponseID returns a new SecurityResponseIDField initialized with val
func NewSecurityResponseID(val string) *SecurityResponseIDField {
	field := &SecurityResponseIDField{}
	field.Value = val
	return field
}

//SecurityResponseTypeField is a INT field
type SecurityResponseTypeField struct{ quickfix.IntValue }

//Tag returns tag.SecurityResponseType (323)
func (f SecurityResponseTypeField) Tag() quickfix.Tag { return tag.SecurityResponseType }

//NewSecurityResponseType returns a new SecurityResponseTypeField initialized with val
func NewSecurityResponseType(val int) *SecurityResponseTypeField {
	field := &SecurityResponseTypeField{}
	field.Value = val
	return field
}

//SecuritySettlAgentAcctNameField is a STRING field
type SecuritySettlAgentAcctNameField struct{ quickfix.StringValue }

//Tag returns tag.SecuritySettlAgentAcctName (179)
func (f SecuritySettlAgentAcctNameField) Tag() quickfix.Tag { return tag.SecuritySettlAgentAcctName }

//NewSecuritySettlAgentAcctName returns a new SecuritySettlAgentAcctNameField initialized with val
func NewSecuritySettlAgentAcctName(val string) *SecuritySettlAgentAcctNameField {
	field := &SecuritySettlAgentAcctNameField{}
	field.Value = val
	return field
}

//SecuritySettlAgentAcctNumField is a STRING field
type SecuritySettlAgentAcctNumField struct{ quickfix.StringValue }

//Tag returns tag.SecuritySettlAgentAcctNum (178)
func (f SecuritySettlAgentAcctNumField) Tag() quickfix.Tag { return tag.SecuritySettlAgentAcctNum }

//NewSecuritySettlAgentAcctNum returns a new SecuritySettlAgentAcctNumField initialized with val
func NewSecuritySettlAgentAcctNum(val string) *SecuritySettlAgentAcctNumField {
	field := &SecuritySettlAgentAcctNumField{}
	field.Value = val
	return field
}

//SecuritySettlAgentCodeField is a STRING field
type SecuritySettlAgentCodeField struct{ quickfix.StringValue }

//Tag returns tag.SecuritySettlAgentCode (177)
func (f SecuritySettlAgentCodeField) Tag() quickfix.Tag { return tag.SecuritySettlAgentCode }

//NewSecuritySettlAgentCode returns a new SecuritySettlAgentCodeField initialized with val
func NewSecuritySettlAgentCode(val string) *SecuritySettlAgentCodeField {
	field := &SecuritySettlAgentCodeField{}
	field.Value = val
	return field
}

//SecuritySettlAgentContactNameField is a STRING field
type SecuritySettlAgentContactNameField struct{ quickfix.StringValue }

//Tag returns tag.SecuritySettlAgentContactName (180)
func (f SecuritySettlAgentContactNameField) Tag() quickfix.Tag {
	return tag.SecuritySettlAgentContactName
}

//NewSecuritySettlAgentContactName returns a new SecuritySettlAgentContactNameField initialized with val
func NewSecuritySettlAgentContactName(val string) *SecuritySettlAgentContactNameField {
	field := &SecuritySettlAgentContactNameField{}
	field.Value = val
	return field
}

//SecuritySettlAgentContactPhoneField is a STRING field
type SecuritySettlAgentContactPhoneField struct{ quickfix.StringValue }

//Tag returns tag.SecuritySettlAgentContactPhone (181)
func (f SecuritySettlAgentContactPhoneField) Tag() quickfix.Tag {
	return tag.SecuritySettlAgentContactPhone
}

//NewSecuritySettlAgentContactPhone returns a new SecuritySettlAgentContactPhoneField initialized with val
func NewSecuritySettlAgentContactPhone(val string) *SecuritySettlAgentContactPhoneField {
	field := &SecuritySettlAgentContactPhoneField{}
	field.Value = val
	return field
}

//SecuritySettlAgentNameField is a STRING field
type SecuritySettlAgentNameField struct{ quickfix.StringValue }

//Tag returns tag.SecuritySettlAgentName (176)
func (f SecuritySettlAgentNameField) Tag() quickfix.Tag { return tag.SecuritySettlAgentName }

//NewSecuritySettlAgentName returns a new SecuritySettlAgentNameField initialized with val
func NewSecuritySettlAgentName(val string) *SecuritySettlAgentNameField {
	field := &SecuritySettlAgentNameField{}
	field.Value = val
	return field
}

//SecurityStatusField is a STRING field
type SecurityStatusField struct{ quickfix.StringValue }

//Tag returns tag.SecurityStatus (965)
func (f SecurityStatusField) Tag() quickfix.Tag { return tag.SecurityStatus }

//NewSecurityStatus returns a new SecurityStatusField initialized with val
func NewSecurityStatus(val string) *SecurityStatusField {
	field := &SecurityStatusField{}
	field.Value = val
	return field
}

//SecurityStatusReqIDField is a STRING field
type SecurityStatusReqIDField struct{ quickfix.StringValue }

//Tag returns tag.SecurityStatusReqID (324)
func (f SecurityStatusReqIDField) Tag() quickfix.Tag { return tag.SecurityStatusReqID }

//NewSecurityStatusReqID returns a new SecurityStatusReqIDField initialized with val
func NewSecurityStatusReqID(val string) *SecurityStatusReqIDField {
	field := &SecurityStatusReqIDField{}
	field.Value = val
	return field
}

//SecuritySubTypeField is a STRING field
type SecuritySubTypeField struct{ quickfix.StringValue }

//Tag returns tag.SecuritySubType (762)
func (f SecuritySubTypeField) Tag() quickfix.Tag { return tag.SecuritySubType }

//NewSecuritySubType returns a new SecuritySubTypeField initialized with val
func NewSecuritySubType(val string) *SecuritySubTypeField {
	field := &SecuritySubTypeField{}
	field.Value = val
	return field
}

//SecurityTradingEventField is a INT field
type SecurityTradingEventField struct{ quickfix.IntValue }

//Tag returns tag.SecurityTradingEvent (1174)
func (f SecurityTradingEventField) Tag() quickfix.Tag { return tag.SecurityTradingEvent }

//NewSecurityTradingEvent returns a new SecurityTradingEventField initialized with val
func NewSecurityTradingEvent(val int) *SecurityTradingEventField {
	field := &SecurityTradingEventField{}
	field.Value = val
	return field
}

//SecurityTradingStatusField is a INT field
type SecurityTradingStatusField struct{ quickfix.IntValue }

//Tag returns tag.SecurityTradingStatus (326)
func (f SecurityTradingStatusField) Tag() quickfix.Tag { return tag.SecurityTradingStatus }

//NewSecurityTradingStatus returns a new SecurityTradingStatusField initialized with val
func NewSecurityTradingStatus(val int) *SecurityTradingStatusField {
	field := &SecurityTradingStatusField{}
	field.Value = val
	return field
}

//SecurityTypeField is a STRING field
type SecurityTypeField struct{ quickfix.StringValue }

//Tag returns tag.SecurityType (167)
func (f SecurityTypeField) Tag() quickfix.Tag { return tag.SecurityType }

//NewSecurityType returns a new SecurityTypeField initialized with val
func NewSecurityType(val string) *SecurityTypeField {
	field := &SecurityTypeField{}
	field.Value = val
	return field
}

//SecurityUpdateActionField is a CHAR field
type SecurityUpdateActionField struct{ quickfix.CharValue }

//Tag returns tag.SecurityUpdateAction (980)
func (f SecurityUpdateActionField) Tag() quickfix.Tag { return tag.SecurityUpdateAction }

//NewSecurityUpdateAction returns a new SecurityUpdateActionField initialized with val
func NewSecurityUpdateAction(val string) *SecurityUpdateActionField {
	field := &SecurityUpdateActionField{}
	field.Value = val
	return field
}

//SecurityXMLField is a XMLDATA field
type SecurityXMLField struct{ quickfix.XMLDataValue }

//Tag returns tag.SecurityXML (1185)
func (f SecurityXMLField) Tag() quickfix.Tag { return tag.SecurityXML }

//NewSecurityXML returns a new SecurityXMLField initialized with val
func NewSecurityXML(val string) *SecurityXMLField {
	field := &SecurityXMLField{}
	field.Value = val
	return field
}

//SecurityXMLLenField is a LENGTH field
type SecurityXMLLenField struct{ quickfix.LengthValue }

//Tag returns tag.SecurityXMLLen (1184)
func (f SecurityXMLLenField) Tag() quickfix.Tag { return tag.SecurityXMLLen }

//NewSecurityXMLLen returns a new SecurityXMLLenField initialized with val
func NewSecurityXMLLen(val int) *SecurityXMLLenField {
	field := &SecurityXMLLenField{}
	field.Value = val
	return field
}

//SecurityXMLSchemaField is a STRING field
type SecurityXMLSchemaField struct{ quickfix.StringValue }

//Tag returns tag.SecurityXMLSchema (1186)
func (f SecurityXMLSchemaField) Tag() quickfix.Tag { return tag.SecurityXMLSchema }

//NewSecurityXMLSchema returns a new SecurityXMLSchemaField initialized with val
func NewSecurityXMLSchema(val string) *SecurityXMLSchemaField {
	field := &SecurityXMLSchemaField{}
	field.Value = val
	return field
}

//SellVolumeField is a QTY field
type SellVolumeField struct{ quickfix.QtyValue }

//Tag returns tag.SellVolume (331)
func (f SellVolumeField) Tag() quickfix.Tag { return tag.SellVolume }

//NewSellVolume returns a new SellVolumeField initialized with val
func NewSellVolume(val float64) *SellVolumeField {
	field := &SellVolumeField{}
	field.Value = val
	return field
}

//SellerDaysField is a INT field
type SellerDaysField struct{ quickfix.IntValue }

//Tag returns tag.SellerDays (287)
func (f SellerDaysField) Tag() quickfix.Tag { return tag.SellerDays }

//NewSellerDays returns a new SellerDaysField initialized with val
func NewSellerDays(val int) *SellerDaysField {
	field := &SellerDaysField{}
	field.Value = val
	return field
}

//SenderCompIDField is a STRING field
type SenderCompIDField struct{ quickfix.StringValue }

//Tag returns tag.SenderCompID (49)
func (f SenderCompIDField) Tag() quickfix.Tag { return tag.SenderCompID }

//NewSenderCompID returns a new SenderCompIDField initialized with val
func NewSenderCompID(val string) *SenderCompIDField {
	field := &SenderCompIDField{}
	field.Value = val
	return field
}

//SenderLocationIDField is a STRING field
type SenderLocationIDField struct{ quickfix.StringValue }

//Tag returns tag.SenderLocationID (142)
func (f SenderLocationIDField) Tag() quickfix.Tag { return tag.SenderLocationID }

//NewSenderLocationID returns a new SenderLocationIDField initialized with val
func NewSenderLocationID(val string) *SenderLocationIDField {
	field := &SenderLocationIDField{}
	field.Value = val
	return field
}

//SenderSubIDField is a STRING field
type SenderSubIDField struct{ quickfix.StringValue }

//Tag returns tag.SenderSubID (50)
func (f SenderSubIDField) Tag() quickfix.Tag { return tag.SenderSubID }

//NewSenderSubID returns a new SenderSubIDField initialized with val
func NewSenderSubID(val string) *SenderSubIDField {
	field := &SenderSubIDField{}
	field.Value = val
	return field
}

//SendingDateField is a LOCALMKTDATE field
type SendingDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.SendingDate (51)
func (f SendingDateField) Tag() quickfix.Tag { return tag.SendingDate }

//NewSendingDate returns a new SendingDateField initialized with val
func NewSendingDate(val string) *SendingDateField {
	field := &SendingDateField{}
	field.Value = val
	return field
}

//SendingTimeField is a UTCTIMESTAMP field
type SendingTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.SendingTime (52)
func (f SendingTimeField) Tag() quickfix.Tag { return tag.SendingTime }

//SeniorityField is a STRING field
type SeniorityField struct{ quickfix.StringValue }

//Tag returns tag.Seniority (1450)
func (f SeniorityField) Tag() quickfix.Tag { return tag.Seniority }

//NewSeniority returns a new SeniorityField initialized with val
func NewSeniority(val string) *SeniorityField {
	field := &SeniorityField{}
	field.Value = val
	return field
}

//SessionRejectReasonField is a INT field
type SessionRejectReasonField struct{ quickfix.IntValue }

//Tag returns tag.SessionRejectReason (373)
func (f SessionRejectReasonField) Tag() quickfix.Tag { return tag.SessionRejectReason }

//NewSessionRejectReason returns a new SessionRejectReasonField initialized with val
func NewSessionRejectReason(val int) *SessionRejectReasonField {
	field := &SessionRejectReasonField{}
	field.Value = val
	return field
}

//SessionStatusField is a INT field
type SessionStatusField struct{ quickfix.IntValue }

//Tag returns tag.SessionStatus (1409)
func (f SessionStatusField) Tag() quickfix.Tag { return tag.SessionStatus }

//NewSessionStatus returns a new SessionStatusField initialized with val
func NewSessionStatus(val int) *SessionStatusField {
	field := &SessionStatusField{}
	field.Value = val
	return field
}

//SettlBrkrCodeField is a STRING field
type SettlBrkrCodeField struct{ quickfix.StringValue }

//Tag returns tag.SettlBrkrCode (174)
func (f SettlBrkrCodeField) Tag() quickfix.Tag { return tag.SettlBrkrCode }

//NewSettlBrkrCode returns a new SettlBrkrCodeField initialized with val
func NewSettlBrkrCode(val string) *SettlBrkrCodeField {
	field := &SettlBrkrCodeField{}
	field.Value = val
	return field
}

//SettlCurrAmtField is a AMT field
type SettlCurrAmtField struct{ quickfix.AmtValue }

//Tag returns tag.SettlCurrAmt (119)
func (f SettlCurrAmtField) Tag() quickfix.Tag { return tag.SettlCurrAmt }

//NewSettlCurrAmt returns a new SettlCurrAmtField initialized with val
func NewSettlCurrAmt(val float64) *SettlCurrAmtField {
	field := &SettlCurrAmtField{}
	field.Value = val
	return field
}

//SettlCurrBidFxRateField is a FLOAT field
type SettlCurrBidFxRateField struct{ quickfix.FloatValue }

//Tag returns tag.SettlCurrBidFxRate (656)
func (f SettlCurrBidFxRateField) Tag() quickfix.Tag { return tag.SettlCurrBidFxRate }

//NewSettlCurrBidFxRate returns a new SettlCurrBidFxRateField initialized with val
func NewSettlCurrBidFxRate(val float64) *SettlCurrBidFxRateField {
	field := &SettlCurrBidFxRateField{}
	field.Value = val
	return field
}

//SettlCurrFxRateField is a FLOAT field
type SettlCurrFxRateField struct{ quickfix.FloatValue }

//Tag returns tag.SettlCurrFxRate (155)
func (f SettlCurrFxRateField) Tag() quickfix.Tag { return tag.SettlCurrFxRate }

//NewSettlCurrFxRate returns a new SettlCurrFxRateField initialized with val
func NewSettlCurrFxRate(val float64) *SettlCurrFxRateField {
	field := &SettlCurrFxRateField{}
	field.Value = val
	return field
}

//SettlCurrFxRateCalcField is a CHAR field
type SettlCurrFxRateCalcField struct{ quickfix.CharValue }

//Tag returns tag.SettlCurrFxRateCalc (156)
func (f SettlCurrFxRateCalcField) Tag() quickfix.Tag { return tag.SettlCurrFxRateCalc }

//NewSettlCurrFxRateCalc returns a new SettlCurrFxRateCalcField initialized with val
func NewSettlCurrFxRateCalc(val string) *SettlCurrFxRateCalcField {
	field := &SettlCurrFxRateCalcField{}
	field.Value = val
	return field
}

//SettlCurrOfferFxRateField is a FLOAT field
type SettlCurrOfferFxRateField struct{ quickfix.FloatValue }

//Tag returns tag.SettlCurrOfferFxRate (657)
func (f SettlCurrOfferFxRateField) Tag() quickfix.Tag { return tag.SettlCurrOfferFxRate }

//NewSettlCurrOfferFxRate returns a new SettlCurrOfferFxRateField initialized with val
func NewSettlCurrOfferFxRate(val float64) *SettlCurrOfferFxRateField {
	field := &SettlCurrOfferFxRateField{}
	field.Value = val
	return field
}

//SettlCurrencyField is a CURRENCY field
type SettlCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.SettlCurrency (120)
func (f SettlCurrencyField) Tag() quickfix.Tag { return tag.SettlCurrency }

//NewSettlCurrency returns a new SettlCurrencyField initialized with val
func NewSettlCurrency(val string) *SettlCurrencyField {
	field := &SettlCurrencyField{}
	field.Value = val
	return field
}

//SettlDateField is a LOCALMKTDATE field
type SettlDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.SettlDate (64)
func (f SettlDateField) Tag() quickfix.Tag { return tag.SettlDate }

//NewSettlDate returns a new SettlDateField initialized with val
func NewSettlDate(val string) *SettlDateField {
	field := &SettlDateField{}
	field.Value = val
	return field
}

//SettlDate2Field is a LOCALMKTDATE field
type SettlDate2Field struct{ quickfix.LocalMktDateValue }

//Tag returns tag.SettlDate2 (193)
func (f SettlDate2Field) Tag() quickfix.Tag { return tag.SettlDate2 }

//NewSettlDate2 returns a new SettlDate2Field initialized with val
func NewSettlDate2(val string) *SettlDate2Field {
	field := &SettlDate2Field{}
	field.Value = val
	return field
}

//SettlDeliveryTypeField is a INT field
type SettlDeliveryTypeField struct{ quickfix.IntValue }

//Tag returns tag.SettlDeliveryType (172)
func (f SettlDeliveryTypeField) Tag() quickfix.Tag { return tag.SettlDeliveryType }

//NewSettlDeliveryType returns a new SettlDeliveryTypeField initialized with val
func NewSettlDeliveryType(val int) *SettlDeliveryTypeField {
	field := &SettlDeliveryTypeField{}
	field.Value = val
	return field
}

//SettlDepositoryCodeField is a STRING field
type SettlDepositoryCodeField struct{ quickfix.StringValue }

//Tag returns tag.SettlDepositoryCode (173)
func (f SettlDepositoryCodeField) Tag() quickfix.Tag { return tag.SettlDepositoryCode }

//NewSettlDepositoryCode returns a new SettlDepositoryCodeField initialized with val
func NewSettlDepositoryCode(val string) *SettlDepositoryCodeField {
	field := &SettlDepositoryCodeField{}
	field.Value = val
	return field
}

//SettlInstCodeField is a STRING field
type SettlInstCodeField struct{ quickfix.StringValue }

//Tag returns tag.SettlInstCode (175)
func (f SettlInstCodeField) Tag() quickfix.Tag { return tag.SettlInstCode }

//NewSettlInstCode returns a new SettlInstCodeField initialized with val
func NewSettlInstCode(val string) *SettlInstCodeField {
	field := &SettlInstCodeField{}
	field.Value = val
	return field
}

//SettlInstIDField is a STRING field
type SettlInstIDField struct{ quickfix.StringValue }

//Tag returns tag.SettlInstID (162)
func (f SettlInstIDField) Tag() quickfix.Tag { return tag.SettlInstID }

//NewSettlInstID returns a new SettlInstIDField initialized with val
func NewSettlInstID(val string) *SettlInstIDField {
	field := &SettlInstIDField{}
	field.Value = val
	return field
}

//SettlInstModeField is a CHAR field
type SettlInstModeField struct{ quickfix.CharValue }

//Tag returns tag.SettlInstMode (160)
func (f SettlInstModeField) Tag() quickfix.Tag { return tag.SettlInstMode }

//NewSettlInstMode returns a new SettlInstModeField initialized with val
func NewSettlInstMode(val string) *SettlInstModeField {
	field := &SettlInstModeField{}
	field.Value = val
	return field
}

//SettlInstMsgIDField is a STRING field
type SettlInstMsgIDField struct{ quickfix.StringValue }

//Tag returns tag.SettlInstMsgID (777)
func (f SettlInstMsgIDField) Tag() quickfix.Tag { return tag.SettlInstMsgID }

//NewSettlInstMsgID returns a new SettlInstMsgIDField initialized with val
func NewSettlInstMsgID(val string) *SettlInstMsgIDField {
	field := &SettlInstMsgIDField{}
	field.Value = val
	return field
}

//SettlInstRefIDField is a STRING field
type SettlInstRefIDField struct{ quickfix.StringValue }

//Tag returns tag.SettlInstRefID (214)
func (f SettlInstRefIDField) Tag() quickfix.Tag { return tag.SettlInstRefID }

//NewSettlInstRefID returns a new SettlInstRefIDField initialized with val
func NewSettlInstRefID(val string) *SettlInstRefIDField {
	field := &SettlInstRefIDField{}
	field.Value = val
	return field
}

//SettlInstReqIDField is a STRING field
type SettlInstReqIDField struct{ quickfix.StringValue }

//Tag returns tag.SettlInstReqID (791)
func (f SettlInstReqIDField) Tag() quickfix.Tag { return tag.SettlInstReqID }

//NewSettlInstReqID returns a new SettlInstReqIDField initialized with val
func NewSettlInstReqID(val string) *SettlInstReqIDField {
	field := &SettlInstReqIDField{}
	field.Value = val
	return field
}

//SettlInstReqRejCodeField is a INT field
type SettlInstReqRejCodeField struct{ quickfix.IntValue }

//Tag returns tag.SettlInstReqRejCode (792)
func (f SettlInstReqRejCodeField) Tag() quickfix.Tag { return tag.SettlInstReqRejCode }

//NewSettlInstReqRejCode returns a new SettlInstReqRejCodeField initialized with val
func NewSettlInstReqRejCode(val int) *SettlInstReqRejCodeField {
	field := &SettlInstReqRejCodeField{}
	field.Value = val
	return field
}

//SettlInstSourceField is a CHAR field
type SettlInstSourceField struct{ quickfix.CharValue }

//Tag returns tag.SettlInstSource (165)
func (f SettlInstSourceField) Tag() quickfix.Tag { return tag.SettlInstSource }

//NewSettlInstSource returns a new SettlInstSourceField initialized with val
func NewSettlInstSource(val string) *SettlInstSourceField {
	field := &SettlInstSourceField{}
	field.Value = val
	return field
}

//SettlInstTransTypeField is a CHAR field
type SettlInstTransTypeField struct{ quickfix.CharValue }

//Tag returns tag.SettlInstTransType (163)
func (f SettlInstTransTypeField) Tag() quickfix.Tag { return tag.SettlInstTransType }

//NewSettlInstTransType returns a new SettlInstTransTypeField initialized with val
func NewSettlInstTransType(val string) *SettlInstTransTypeField {
	field := &SettlInstTransTypeField{}
	field.Value = val
	return field
}

//SettlLocationField is a STRING field
type SettlLocationField struct{ quickfix.StringValue }

//Tag returns tag.SettlLocation (166)
func (f SettlLocationField) Tag() quickfix.Tag { return tag.SettlLocation }

//NewSettlLocation returns a new SettlLocationField initialized with val
func NewSettlLocation(val string) *SettlLocationField {
	field := &SettlLocationField{}
	field.Value = val
	return field
}

//SettlMethodField is a CHAR field
type SettlMethodField struct{ quickfix.CharValue }

//Tag returns tag.SettlMethod (1193)
func (f SettlMethodField) Tag() quickfix.Tag { return tag.SettlMethod }

//NewSettlMethod returns a new SettlMethodField initialized with val
func NewSettlMethod(val string) *SettlMethodField {
	field := &SettlMethodField{}
	field.Value = val
	return field
}

//SettlObligIDField is a STRING field
type SettlObligIDField struct{ quickfix.StringValue }

//Tag returns tag.SettlObligID (1161)
func (f SettlObligIDField) Tag() quickfix.Tag { return tag.SettlObligID }

//NewSettlObligID returns a new SettlObligIDField initialized with val
func NewSettlObligID(val string) *SettlObligIDField {
	field := &SettlObligIDField{}
	field.Value = val
	return field
}

//SettlObligModeField is a INT field
type SettlObligModeField struct{ quickfix.IntValue }

//Tag returns tag.SettlObligMode (1159)
func (f SettlObligModeField) Tag() quickfix.Tag { return tag.SettlObligMode }

//NewSettlObligMode returns a new SettlObligModeField initialized with val
func NewSettlObligMode(val int) *SettlObligModeField {
	field := &SettlObligModeField{}
	field.Value = val
	return field
}

//SettlObligMsgIDField is a STRING field
type SettlObligMsgIDField struct{ quickfix.StringValue }

//Tag returns tag.SettlObligMsgID (1160)
func (f SettlObligMsgIDField) Tag() quickfix.Tag { return tag.SettlObligMsgID }

//NewSettlObligMsgID returns a new SettlObligMsgIDField initialized with val
func NewSettlObligMsgID(val string) *SettlObligMsgIDField {
	field := &SettlObligMsgIDField{}
	field.Value = val
	return field
}

//SettlObligRefIDField is a STRING field
type SettlObligRefIDField struct{ quickfix.StringValue }

//Tag returns tag.SettlObligRefID (1163)
func (f SettlObligRefIDField) Tag() quickfix.Tag { return tag.SettlObligRefID }

//NewSettlObligRefID returns a new SettlObligRefIDField initialized with val
func NewSettlObligRefID(val string) *SettlObligRefIDField {
	field := &SettlObligRefIDField{}
	field.Value = val
	return field
}

//SettlObligSourceField is a CHAR field
type SettlObligSourceField struct{ quickfix.CharValue }

//Tag returns tag.SettlObligSource (1164)
func (f SettlObligSourceField) Tag() quickfix.Tag { return tag.SettlObligSource }

//NewSettlObligSource returns a new SettlObligSourceField initialized with val
func NewSettlObligSource(val string) *SettlObligSourceField {
	field := &SettlObligSourceField{}
	field.Value = val
	return field
}

//SettlObligTransTypeField is a CHAR field
type SettlObligTransTypeField struct{ quickfix.CharValue }

//Tag returns tag.SettlObligTransType (1162)
func (f SettlObligTransTypeField) Tag() quickfix.Tag { return tag.SettlObligTransType }

//NewSettlObligTransType returns a new SettlObligTransTypeField initialized with val
func NewSettlObligTransType(val string) *SettlObligTransTypeField {
	field := &SettlObligTransTypeField{}
	field.Value = val
	return field
}

//SettlPartyIDField is a STRING field
type SettlPartyIDField struct{ quickfix.StringValue }

//Tag returns tag.SettlPartyID (782)
func (f SettlPartyIDField) Tag() quickfix.Tag { return tag.SettlPartyID }

//NewSettlPartyID returns a new SettlPartyIDField initialized with val
func NewSettlPartyID(val string) *SettlPartyIDField {
	field := &SettlPartyIDField{}
	field.Value = val
	return field
}

//SettlPartyIDSourceField is a CHAR field
type SettlPartyIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.SettlPartyIDSource (783)
func (f SettlPartyIDSourceField) Tag() quickfix.Tag { return tag.SettlPartyIDSource }

//NewSettlPartyIDSource returns a new SettlPartyIDSourceField initialized with val
func NewSettlPartyIDSource(val string) *SettlPartyIDSourceField {
	field := &SettlPartyIDSourceField{}
	field.Value = val
	return field
}

//SettlPartyRoleField is a INT field
type SettlPartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.SettlPartyRole (784)
func (f SettlPartyRoleField) Tag() quickfix.Tag { return tag.SettlPartyRole }

//NewSettlPartyRole returns a new SettlPartyRoleField initialized with val
func NewSettlPartyRole(val int) *SettlPartyRoleField {
	field := &SettlPartyRoleField{}
	field.Value = val
	return field
}

//SettlPartySubIDField is a STRING field
type SettlPartySubIDField struct{ quickfix.StringValue }

//Tag returns tag.SettlPartySubID (785)
func (f SettlPartySubIDField) Tag() quickfix.Tag { return tag.SettlPartySubID }

//NewSettlPartySubID returns a new SettlPartySubIDField initialized with val
func NewSettlPartySubID(val string) *SettlPartySubIDField {
	field := &SettlPartySubIDField{}
	field.Value = val
	return field
}

//SettlPartySubIDTypeField is a INT field
type SettlPartySubIDTypeField struct{ quickfix.IntValue }

//Tag returns tag.SettlPartySubIDType (786)
func (f SettlPartySubIDTypeField) Tag() quickfix.Tag { return tag.SettlPartySubIDType }

//NewSettlPartySubIDType returns a new SettlPartySubIDTypeField initialized with val
func NewSettlPartySubIDType(val int) *SettlPartySubIDTypeField {
	field := &SettlPartySubIDTypeField{}
	field.Value = val
	return field
}

//SettlPriceField is a PRICE field
type SettlPriceField struct{ quickfix.PriceValue }

//Tag returns tag.SettlPrice (730)
func (f SettlPriceField) Tag() quickfix.Tag { return tag.SettlPrice }

//NewSettlPrice returns a new SettlPriceField initialized with val
func NewSettlPrice(val float64) *SettlPriceField {
	field := &SettlPriceField{}
	field.Value = val
	return field
}

//SettlPriceTypeField is a INT field
type SettlPriceTypeField struct{ quickfix.IntValue }

//Tag returns tag.SettlPriceType (731)
func (f SettlPriceTypeField) Tag() quickfix.Tag { return tag.SettlPriceType }

//NewSettlPriceType returns a new SettlPriceTypeField initialized with val
func NewSettlPriceType(val int) *SettlPriceTypeField {
	field := &SettlPriceTypeField{}
	field.Value = val
	return field
}

//SettlSessIDField is a STRING field
type SettlSessIDField struct{ quickfix.StringValue }

//Tag returns tag.SettlSessID (716)
func (f SettlSessIDField) Tag() quickfix.Tag { return tag.SettlSessID }

//NewSettlSessID returns a new SettlSessIDField initialized with val
func NewSettlSessID(val string) *SettlSessIDField {
	field := &SettlSessIDField{}
	field.Value = val
	return field
}

//SettlSessSubIDField is a STRING field
type SettlSessSubIDField struct{ quickfix.StringValue }

//Tag returns tag.SettlSessSubID (717)
func (f SettlSessSubIDField) Tag() quickfix.Tag { return tag.SettlSessSubID }

//NewSettlSessSubID returns a new SettlSessSubIDField initialized with val
func NewSettlSessSubID(val string) *SettlSessSubIDField {
	field := &SettlSessSubIDField{}
	field.Value = val
	return field
}

//SettlTypeField is a STRING field
type SettlTypeField struct{ quickfix.StringValue }

//Tag returns tag.SettlType (63)
func (f SettlTypeField) Tag() quickfix.Tag { return tag.SettlType }

//NewSettlType returns a new SettlTypeField initialized with val
func NewSettlType(val string) *SettlTypeField {
	field := &SettlTypeField{}
	field.Value = val
	return field
}

//SettleOnOpenFlagField is a STRING field
type SettleOnOpenFlagField struct{ quickfix.StringValue }

//Tag returns tag.SettleOnOpenFlag (966)
func (f SettleOnOpenFlagField) Tag() quickfix.Tag { return tag.SettleOnOpenFlag }

//NewSettleOnOpenFlag returns a new SettleOnOpenFlagField initialized with val
func NewSettleOnOpenFlag(val string) *SettleOnOpenFlagField {
	field := &SettleOnOpenFlagField{}
	field.Value = val
	return field
}

//SettlementCycleNoField is a INT field
type SettlementCycleNoField struct{ quickfix.IntValue }

//Tag returns tag.SettlementCycleNo (1153)
func (f SettlementCycleNoField) Tag() quickfix.Tag { return tag.SettlementCycleNo }

//NewSettlementCycleNo returns a new SettlementCycleNoField initialized with val
func NewSettlementCycleNo(val int) *SettlementCycleNoField {
	field := &SettlementCycleNoField{}
	field.Value = val
	return field
}

//SettlmntTypField is a CHAR field
type SettlmntTypField struct{ quickfix.CharValue }

//Tag returns tag.SettlmntTyp (63)
func (f SettlmntTypField) Tag() quickfix.Tag { return tag.SettlmntTyp }

//NewSettlmntTyp returns a new SettlmntTypField initialized with val
func NewSettlmntTyp(val string) *SettlmntTypField {
	field := &SettlmntTypField{}
	field.Value = val
	return field
}

//SharedCommissionField is a AMT field
type SharedCommissionField struct{ quickfix.AmtValue }

//Tag returns tag.SharedCommission (858)
func (f SharedCommissionField) Tag() quickfix.Tag { return tag.SharedCommission }

//NewSharedCommission returns a new SharedCommissionField initialized with val
func NewSharedCommission(val float64) *SharedCommissionField {
	field := &SharedCommissionField{}
	field.Value = val
	return field
}

//SharesField is a QTY field
type SharesField struct{ quickfix.QtyValue }

//Tag returns tag.Shares (53)
func (f SharesField) Tag() quickfix.Tag { return tag.Shares }

//NewShares returns a new SharesField initialized with val
func NewShares(val float64) *SharesField {
	field := &SharesField{}
	field.Value = val
	return field
}

//ShortQtyField is a QTY field
type ShortQtyField struct{ quickfix.QtyValue }

//Tag returns tag.ShortQty (705)
func (f ShortQtyField) Tag() quickfix.Tag { return tag.ShortQty }

//NewShortQty returns a new ShortQtyField initialized with val
func NewShortQty(val float64) *ShortQtyField {
	field := &ShortQtyField{}
	field.Value = val
	return field
}

//ShortSaleReasonField is a INT field
type ShortSaleReasonField struct{ quickfix.IntValue }

//Tag returns tag.ShortSaleReason (853)
func (f ShortSaleReasonField) Tag() quickfix.Tag { return tag.ShortSaleReason }

//NewShortSaleReason returns a new ShortSaleReasonField initialized with val
func NewShortSaleReason(val int) *ShortSaleReasonField {
	field := &ShortSaleReasonField{}
	field.Value = val
	return field
}

//SideField is a CHAR field
type SideField struct{ quickfix.CharValue }

//Tag returns tag.Side (54)
func (f SideField) Tag() quickfix.Tag { return tag.Side }

//NewSide returns a new SideField initialized with val
func NewSide(val string) *SideField {
	field := &SideField{}
	field.Value = val
	return field
}

//SideComplianceIDField is a STRING field
type SideComplianceIDField struct{ quickfix.StringValue }

//Tag returns tag.SideComplianceID (659)
func (f SideComplianceIDField) Tag() quickfix.Tag { return tag.SideComplianceID }

//NewSideComplianceID returns a new SideComplianceIDField initialized with val
func NewSideComplianceID(val string) *SideComplianceIDField {
	field := &SideComplianceIDField{}
	field.Value = val
	return field
}

//SideCurrencyField is a CURRENCY field
type SideCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.SideCurrency (1154)
func (f SideCurrencyField) Tag() quickfix.Tag { return tag.SideCurrency }

//NewSideCurrency returns a new SideCurrencyField initialized with val
func NewSideCurrency(val string) *SideCurrencyField {
	field := &SideCurrencyField{}
	field.Value = val
	return field
}

//SideExecIDField is a STRING field
type SideExecIDField struct{ quickfix.StringValue }

//Tag returns tag.SideExecID (1427)
func (f SideExecIDField) Tag() quickfix.Tag { return tag.SideExecID }

//NewSideExecID returns a new SideExecIDField initialized with val
func NewSideExecID(val string) *SideExecIDField {
	field := &SideExecIDField{}
	field.Value = val
	return field
}

//SideFillStationCdField is a STRING field
type SideFillStationCdField struct{ quickfix.StringValue }

//Tag returns tag.SideFillStationCd (1006)
func (f SideFillStationCdField) Tag() quickfix.Tag { return tag.SideFillStationCd }

//NewSideFillStationCd returns a new SideFillStationCdField initialized with val
func NewSideFillStationCd(val string) *SideFillStationCdField {
	field := &SideFillStationCdField{}
	field.Value = val
	return field
}

//SideGrossTradeAmtField is a AMT field
type SideGrossTradeAmtField struct{ quickfix.AmtValue }

//Tag returns tag.SideGrossTradeAmt (1072)
func (f SideGrossTradeAmtField) Tag() quickfix.Tag { return tag.SideGrossTradeAmt }

//NewSideGrossTradeAmt returns a new SideGrossTradeAmtField initialized with val
func NewSideGrossTradeAmt(val float64) *SideGrossTradeAmtField {
	field := &SideGrossTradeAmtField{}
	field.Value = val
	return field
}

//SideLastQtyField is a INT field
type SideLastQtyField struct{ quickfix.IntValue }

//Tag returns tag.SideLastQty (1009)
func (f SideLastQtyField) Tag() quickfix.Tag { return tag.SideLastQty }

//NewSideLastQty returns a new SideLastQtyField initialized with val
func NewSideLastQty(val int) *SideLastQtyField {
	field := &SideLastQtyField{}
	field.Value = val
	return field
}

//SideLiquidityIndField is a INT field
type SideLiquidityIndField struct{ quickfix.IntValue }

//Tag returns tag.SideLiquidityInd (1444)
func (f SideLiquidityIndField) Tag() quickfix.Tag { return tag.SideLiquidityInd }

//NewSideLiquidityInd returns a new SideLiquidityIndField initialized with val
func NewSideLiquidityInd(val int) *SideLiquidityIndField {
	field := &SideLiquidityIndField{}
	field.Value = val
	return field
}

//SideMultiLegReportingTypeField is a INT field
type SideMultiLegReportingTypeField struct{ quickfix.IntValue }

//Tag returns tag.SideMultiLegReportingType (752)
func (f SideMultiLegReportingTypeField) Tag() quickfix.Tag { return tag.SideMultiLegReportingType }

//NewSideMultiLegReportingType returns a new SideMultiLegReportingTypeField initialized with val
func NewSideMultiLegReportingType(val int) *SideMultiLegReportingTypeField {
	field := &SideMultiLegReportingTypeField{}
	field.Value = val
	return field
}

//SideQtyField is a INT field
type SideQtyField struct{ quickfix.IntValue }

//Tag returns tag.SideQty (1009)
func (f SideQtyField) Tag() quickfix.Tag { return tag.SideQty }

//NewSideQty returns a new SideQtyField initialized with val
func NewSideQty(val int) *SideQtyField {
	field := &SideQtyField{}
	field.Value = val
	return field
}

//SideReasonCdField is a STRING field
type SideReasonCdField struct{ quickfix.StringValue }

//Tag returns tag.SideReasonCd (1007)
func (f SideReasonCdField) Tag() quickfix.Tag { return tag.SideReasonCd }

//NewSideReasonCd returns a new SideReasonCdField initialized with val
func NewSideReasonCd(val string) *SideReasonCdField {
	field := &SideReasonCdField{}
	field.Value = val
	return field
}

//SideSettlCurrencyField is a CURRENCY field
type SideSettlCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.SideSettlCurrency (1155)
func (f SideSettlCurrencyField) Tag() quickfix.Tag { return tag.SideSettlCurrency }

//NewSideSettlCurrency returns a new SideSettlCurrencyField initialized with val
func NewSideSettlCurrency(val string) *SideSettlCurrencyField {
	field := &SideSettlCurrencyField{}
	field.Value = val
	return field
}

//SideTimeInForceField is a UTCTIMESTAMP field
type SideTimeInForceField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.SideTimeInForce (962)
func (f SideTimeInForceField) Tag() quickfix.Tag { return tag.SideTimeInForce }

//SideTradeReportIDField is a STRING field
type SideTradeReportIDField struct{ quickfix.StringValue }

//Tag returns tag.SideTradeReportID (1005)
func (f SideTradeReportIDField) Tag() quickfix.Tag { return tag.SideTradeReportID }

//NewSideTradeReportID returns a new SideTradeReportIDField initialized with val
func NewSideTradeReportID(val string) *SideTradeReportIDField {
	field := &SideTradeReportIDField{}
	field.Value = val
	return field
}

//SideTrdRegTimestampField is a UTCTIMESTAMP field
type SideTrdRegTimestampField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.SideTrdRegTimestamp (1012)
func (f SideTrdRegTimestampField) Tag() quickfix.Tag { return tag.SideTrdRegTimestamp }

//SideTrdRegTimestampSrcField is a STRING field
type SideTrdRegTimestampSrcField struct{ quickfix.StringValue }

//Tag returns tag.SideTrdRegTimestampSrc (1014)
func (f SideTrdRegTimestampSrcField) Tag() quickfix.Tag { return tag.SideTrdRegTimestampSrc }

//NewSideTrdRegTimestampSrc returns a new SideTrdRegTimestampSrcField initialized with val
func NewSideTrdRegTimestampSrc(val string) *SideTrdRegTimestampSrcField {
	field := &SideTrdRegTimestampSrcField{}
	field.Value = val
	return field
}

//SideTrdRegTimestampTypeField is a INT field
type SideTrdRegTimestampTypeField struct{ quickfix.IntValue }

//Tag returns tag.SideTrdRegTimestampType (1013)
func (f SideTrdRegTimestampTypeField) Tag() quickfix.Tag { return tag.SideTrdRegTimestampType }

//NewSideTrdRegTimestampType returns a new SideTrdRegTimestampTypeField initialized with val
func NewSideTrdRegTimestampType(val int) *SideTrdRegTimestampTypeField {
	field := &SideTrdRegTimestampTypeField{}
	field.Value = val
	return field
}

//SideTrdSubTypField is a INT field
type SideTrdSubTypField struct{ quickfix.IntValue }

//Tag returns tag.SideTrdSubTyp (1008)
func (f SideTrdSubTypField) Tag() quickfix.Tag { return tag.SideTrdSubTyp }

//NewSideTrdSubTyp returns a new SideTrdSubTypField initialized with val
func NewSideTrdSubTyp(val int) *SideTrdSubTypField {
	field := &SideTrdSubTypField{}
	field.Value = val
	return field
}

//SideValue1Field is a AMT field
type SideValue1Field struct{ quickfix.AmtValue }

//Tag returns tag.SideValue1 (396)
func (f SideValue1Field) Tag() quickfix.Tag { return tag.SideValue1 }

//NewSideValue1 returns a new SideValue1Field initialized with val
func NewSideValue1(val float64) *SideValue1Field {
	field := &SideValue1Field{}
	field.Value = val
	return field
}

//SideValue2Field is a AMT field
type SideValue2Field struct{ quickfix.AmtValue }

//Tag returns tag.SideValue2 (397)
func (f SideValue2Field) Tag() quickfix.Tag { return tag.SideValue2 }

//NewSideValue2 returns a new SideValue2Field initialized with val
func NewSideValue2(val float64) *SideValue2Field {
	field := &SideValue2Field{}
	field.Value = val
	return field
}

//SideValueIndField is a INT field
type SideValueIndField struct{ quickfix.IntValue }

//Tag returns tag.SideValueInd (401)
func (f SideValueIndField) Tag() quickfix.Tag { return tag.SideValueInd }

//NewSideValueInd returns a new SideValueIndField initialized with val
func NewSideValueInd(val int) *SideValueIndField {
	field := &SideValueIndField{}
	field.Value = val
	return field
}

//SignatureField is a DATA field
type SignatureField struct{ quickfix.DataValue }

//Tag returns tag.Signature (89)
func (f SignatureField) Tag() quickfix.Tag { return tag.Signature }

//NewSignature returns a new SignatureField initialized with val
func NewSignature(val string) *SignatureField {
	field := &SignatureField{}
	field.Value = val
	return field
}

//SignatureLengthField is a LENGTH field
type SignatureLengthField struct{ quickfix.LengthValue }

//Tag returns tag.SignatureLength (93)
func (f SignatureLengthField) Tag() quickfix.Tag { return tag.SignatureLength }

//NewSignatureLength returns a new SignatureLengthField initialized with val
func NewSignatureLength(val int) *SignatureLengthField {
	field := &SignatureLengthField{}
	field.Value = val
	return field
}

//SolicitedFlagField is a BOOLEAN field
type SolicitedFlagField struct{ quickfix.BooleanValue }

//Tag returns tag.SolicitedFlag (377)
func (f SolicitedFlagField) Tag() quickfix.Tag { return tag.SolicitedFlag }

//NewSolicitedFlag returns a new SolicitedFlagField initialized with val
func NewSolicitedFlag(val bool) *SolicitedFlagField {
	field := &SolicitedFlagField{}
	field.Value = val
	return field
}

//SpreadField is a PRICEOFFSET field
type SpreadField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.Spread (218)
func (f SpreadField) Tag() quickfix.Tag { return tag.Spread }

//NewSpread returns a new SpreadField initialized with val
func NewSpread(val float64) *SpreadField {
	field := &SpreadField{}
	field.Value = val
	return field
}

//SpreadToBenchmarkField is a PRICEOFFSET field
type SpreadToBenchmarkField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.SpreadToBenchmark (218)
func (f SpreadToBenchmarkField) Tag() quickfix.Tag { return tag.SpreadToBenchmark }

//NewSpreadToBenchmark returns a new SpreadToBenchmarkField initialized with val
func NewSpreadToBenchmark(val float64) *SpreadToBenchmarkField {
	field := &SpreadToBenchmarkField{}
	field.Value = val
	return field
}

//StandInstDbIDField is a STRING field
type StandInstDbIDField struct{ quickfix.StringValue }

//Tag returns tag.StandInstDbID (171)
func (f StandInstDbIDField) Tag() quickfix.Tag { return tag.StandInstDbID }

//NewStandInstDbID returns a new StandInstDbIDField initialized with val
func NewStandInstDbID(val string) *StandInstDbIDField {
	field := &StandInstDbIDField{}
	field.Value = val
	return field
}

//StandInstDbNameField is a STRING field
type StandInstDbNameField struct{ quickfix.StringValue }

//Tag returns tag.StandInstDbName (170)
func (f StandInstDbNameField) Tag() quickfix.Tag { return tag.StandInstDbName }

//NewStandInstDbName returns a new StandInstDbNameField initialized with val
func NewStandInstDbName(val string) *StandInstDbNameField {
	field := &StandInstDbNameField{}
	field.Value = val
	return field
}

//StandInstDbTypeField is a INT field
type StandInstDbTypeField struct{ quickfix.IntValue }

//Tag returns tag.StandInstDbType (169)
func (f StandInstDbTypeField) Tag() quickfix.Tag { return tag.StandInstDbType }

//NewStandInstDbType returns a new StandInstDbTypeField initialized with val
func NewStandInstDbType(val int) *StandInstDbTypeField {
	field := &StandInstDbTypeField{}
	field.Value = val
	return field
}

//StartCashField is a AMT field
type StartCashField struct{ quickfix.AmtValue }

//Tag returns tag.StartCash (921)
func (f StartCashField) Tag() quickfix.Tag { return tag.StartCash }

//NewStartCash returns a new StartCashField initialized with val
func NewStartCash(val float64) *StartCashField {
	field := &StartCashField{}
	field.Value = val
	return field
}

//StartDateField is a LOCALMKTDATE field
type StartDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.StartDate (916)
func (f StartDateField) Tag() quickfix.Tag { return tag.StartDate }

//NewStartDate returns a new StartDateField initialized with val
func NewStartDate(val string) *StartDateField {
	field := &StartDateField{}
	field.Value = val
	return field
}

//StartMaturityMonthYearField is a MONTHYEAR field
type StartMaturityMonthYearField struct{ quickfix.MonthYearValue }

//Tag returns tag.StartMaturityMonthYear (1241)
func (f StartMaturityMonthYearField) Tag() quickfix.Tag { return tag.StartMaturityMonthYear }

//NewStartMaturityMonthYear returns a new StartMaturityMonthYearField initialized with val
func NewStartMaturityMonthYear(val string) *StartMaturityMonthYearField {
	field := &StartMaturityMonthYearField{}
	field.Value = val
	return field
}

//StartStrikePxRangeField is a PRICE field
type StartStrikePxRangeField struct{ quickfix.PriceValue }

//Tag returns tag.StartStrikePxRange (1202)
func (f StartStrikePxRangeField) Tag() quickfix.Tag { return tag.StartStrikePxRange }

//NewStartStrikePxRange returns a new StartStrikePxRangeField initialized with val
func NewStartStrikePxRange(val float64) *StartStrikePxRangeField {
	field := &StartStrikePxRangeField{}
	field.Value = val
	return field
}

//StartTickPriceRangeField is a PRICE field
type StartTickPriceRangeField struct{ quickfix.PriceValue }

//Tag returns tag.StartTickPriceRange (1206)
func (f StartTickPriceRangeField) Tag() quickfix.Tag { return tag.StartTickPriceRange }

//NewStartTickPriceRange returns a new StartTickPriceRangeField initialized with val
func NewStartTickPriceRange(val float64) *StartTickPriceRangeField {
	field := &StartTickPriceRangeField{}
	field.Value = val
	return field
}

//StateOrProvinceOfIssueField is a STRING field
type StateOrProvinceOfIssueField struct{ quickfix.StringValue }

//Tag returns tag.StateOrProvinceOfIssue (471)
func (f StateOrProvinceOfIssueField) Tag() quickfix.Tag { return tag.StateOrProvinceOfIssue }

//NewStateOrProvinceOfIssue returns a new StateOrProvinceOfIssueField initialized with val
func NewStateOrProvinceOfIssue(val string) *StateOrProvinceOfIssueField {
	field := &StateOrProvinceOfIssueField{}
	field.Value = val
	return field
}

//StatsTypeField is a INT field
type StatsTypeField struct{ quickfix.IntValue }

//Tag returns tag.StatsType (1176)
func (f StatsTypeField) Tag() quickfix.Tag { return tag.StatsType }

//NewStatsType returns a new StatsTypeField initialized with val
func NewStatsType(val int) *StatsTypeField {
	field := &StatsTypeField{}
	field.Value = val
	return field
}

//StatusTextField is a STRING field
type StatusTextField struct{ quickfix.StringValue }

//Tag returns tag.StatusText (929)
func (f StatusTextField) Tag() quickfix.Tag { return tag.StatusText }

//NewStatusText returns a new StatusTextField initialized with val
func NewStatusText(val string) *StatusTextField {
	field := &StatusTextField{}
	field.Value = val
	return field
}

//StatusValueField is a INT field
type StatusValueField struct{ quickfix.IntValue }

//Tag returns tag.StatusValue (928)
func (f StatusValueField) Tag() quickfix.Tag { return tag.StatusValue }

//NewStatusValue returns a new StatusValueField initialized with val
func NewStatusValue(val int) *StatusValueField {
	field := &StatusValueField{}
	field.Value = val
	return field
}

//StipulationTypeField is a STRING field
type StipulationTypeField struct{ quickfix.StringValue }

//Tag returns tag.StipulationType (233)
func (f StipulationTypeField) Tag() quickfix.Tag { return tag.StipulationType }

//NewStipulationType returns a new StipulationTypeField initialized with val
func NewStipulationType(val string) *StipulationTypeField {
	field := &StipulationTypeField{}
	field.Value = val
	return field
}

//StipulationValueField is a STRING field
type StipulationValueField struct{ quickfix.StringValue }

//Tag returns tag.StipulationValue (234)
func (f StipulationValueField) Tag() quickfix.Tag { return tag.StipulationValue }

//NewStipulationValue returns a new StipulationValueField initialized with val
func NewStipulationValue(val string) *StipulationValueField {
	field := &StipulationValueField{}
	field.Value = val
	return field
}

//StopPxField is a PRICE field
type StopPxField struct{ quickfix.PriceValue }

//Tag returns tag.StopPx (99)
func (f StopPxField) Tag() quickfix.Tag { return tag.StopPx }

//NewStopPx returns a new StopPxField initialized with val
func NewStopPx(val float64) *StopPxField {
	field := &StopPxField{}
	field.Value = val
	return field
}

//StrategyParameterNameField is a STRING field
type StrategyParameterNameField struct{ quickfix.StringValue }

//Tag returns tag.StrategyParameterName (958)
func (f StrategyParameterNameField) Tag() quickfix.Tag { return tag.StrategyParameterName }

//NewStrategyParameterName returns a new StrategyParameterNameField initialized with val
func NewStrategyParameterName(val string) *StrategyParameterNameField {
	field := &StrategyParameterNameField{}
	field.Value = val
	return field
}

//StrategyParameterTypeField is a INT field
type StrategyParameterTypeField struct{ quickfix.IntValue }

//Tag returns tag.StrategyParameterType (959)
func (f StrategyParameterTypeField) Tag() quickfix.Tag { return tag.StrategyParameterType }

//NewStrategyParameterType returns a new StrategyParameterTypeField initialized with val
func NewStrategyParameterType(val int) *StrategyParameterTypeField {
	field := &StrategyParameterTypeField{}
	field.Value = val
	return field
}

//StrategyParameterValueField is a STRING field
type StrategyParameterValueField struct{ quickfix.StringValue }

//Tag returns tag.StrategyParameterValue (960)
func (f StrategyParameterValueField) Tag() quickfix.Tag { return tag.StrategyParameterValue }

//NewStrategyParameterValue returns a new StrategyParameterValueField initialized with val
func NewStrategyParameterValue(val string) *StrategyParameterValueField {
	field := &StrategyParameterValueField{}
	field.Value = val
	return field
}

//StreamAsgnAckTypeField is a INT field
type StreamAsgnAckTypeField struct{ quickfix.IntValue }

//Tag returns tag.StreamAsgnAckType (1503)
func (f StreamAsgnAckTypeField) Tag() quickfix.Tag { return tag.StreamAsgnAckType }

//NewStreamAsgnAckType returns a new StreamAsgnAckTypeField initialized with val
func NewStreamAsgnAckType(val int) *StreamAsgnAckTypeField {
	field := &StreamAsgnAckTypeField{}
	field.Value = val
	return field
}

//StreamAsgnRejReasonField is a INT field
type StreamAsgnRejReasonField struct{ quickfix.IntValue }

//Tag returns tag.StreamAsgnRejReason (1502)
func (f StreamAsgnRejReasonField) Tag() quickfix.Tag { return tag.StreamAsgnRejReason }

//NewStreamAsgnRejReason returns a new StreamAsgnRejReasonField initialized with val
func NewStreamAsgnRejReason(val int) *StreamAsgnRejReasonField {
	field := &StreamAsgnRejReasonField{}
	field.Value = val
	return field
}

//StreamAsgnReqIDField is a STRING field
type StreamAsgnReqIDField struct{ quickfix.StringValue }

//Tag returns tag.StreamAsgnReqID (1497)
func (f StreamAsgnReqIDField) Tag() quickfix.Tag { return tag.StreamAsgnReqID }

//NewStreamAsgnReqID returns a new StreamAsgnReqIDField initialized with val
func NewStreamAsgnReqID(val string) *StreamAsgnReqIDField {
	field := &StreamAsgnReqIDField{}
	field.Value = val
	return field
}

//StreamAsgnReqTypeField is a INT field
type StreamAsgnReqTypeField struct{ quickfix.IntValue }

//Tag returns tag.StreamAsgnReqType (1498)
func (f StreamAsgnReqTypeField) Tag() quickfix.Tag { return tag.StreamAsgnReqType }

//NewStreamAsgnReqType returns a new StreamAsgnReqTypeField initialized with val
func NewStreamAsgnReqType(val int) *StreamAsgnReqTypeField {
	field := &StreamAsgnReqTypeField{}
	field.Value = val
	return field
}

//StreamAsgnRptIDField is a STRING field
type StreamAsgnRptIDField struct{ quickfix.StringValue }

//Tag returns tag.StreamAsgnRptID (1501)
func (f StreamAsgnRptIDField) Tag() quickfix.Tag { return tag.StreamAsgnRptID }

//NewStreamAsgnRptID returns a new StreamAsgnRptIDField initialized with val
func NewStreamAsgnRptID(val string) *StreamAsgnRptIDField {
	field := &StreamAsgnRptIDField{}
	field.Value = val
	return field
}

//StreamAsgnTypeField is a INT field
type StreamAsgnTypeField struct{ quickfix.IntValue }

//Tag returns tag.StreamAsgnType (1617)
func (f StreamAsgnTypeField) Tag() quickfix.Tag { return tag.StreamAsgnType }

//NewStreamAsgnType returns a new StreamAsgnTypeField initialized with val
func NewStreamAsgnType(val int) *StreamAsgnTypeField {
	field := &StreamAsgnTypeField{}
	field.Value = val
	return field
}

//StrikeCurrencyField is a CURRENCY field
type StrikeCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.StrikeCurrency (947)
func (f StrikeCurrencyField) Tag() quickfix.Tag { return tag.StrikeCurrency }

//NewStrikeCurrency returns a new StrikeCurrencyField initialized with val
func NewStrikeCurrency(val string) *StrikeCurrencyField {
	field := &StrikeCurrencyField{}
	field.Value = val
	return field
}

//StrikeExerciseStyleField is a INT field
type StrikeExerciseStyleField struct{ quickfix.IntValue }

//Tag returns tag.StrikeExerciseStyle (1304)
func (f StrikeExerciseStyleField) Tag() quickfix.Tag { return tag.StrikeExerciseStyle }

//NewStrikeExerciseStyle returns a new StrikeExerciseStyleField initialized with val
func NewStrikeExerciseStyle(val int) *StrikeExerciseStyleField {
	field := &StrikeExerciseStyleField{}
	field.Value = val
	return field
}

//StrikeIncrementField is a FLOAT field
type StrikeIncrementField struct{ quickfix.FloatValue }

//Tag returns tag.StrikeIncrement (1204)
func (f StrikeIncrementField) Tag() quickfix.Tag { return tag.StrikeIncrement }

//NewStrikeIncrement returns a new StrikeIncrementField initialized with val
func NewStrikeIncrement(val float64) *StrikeIncrementField {
	field := &StrikeIncrementField{}
	field.Value = val
	return field
}

//StrikeMultiplierField is a FLOAT field
type StrikeMultiplierField struct{ quickfix.FloatValue }

//Tag returns tag.StrikeMultiplier (967)
func (f StrikeMultiplierField) Tag() quickfix.Tag { return tag.StrikeMultiplier }

//NewStrikeMultiplier returns a new StrikeMultiplierField initialized with val
func NewStrikeMultiplier(val float64) *StrikeMultiplierField {
	field := &StrikeMultiplierField{}
	field.Value = val
	return field
}

//StrikePriceField is a PRICE field
type StrikePriceField struct{ quickfix.PriceValue }

//Tag returns tag.StrikePrice (202)
func (f StrikePriceField) Tag() quickfix.Tag { return tag.StrikePrice }

//NewStrikePrice returns a new StrikePriceField initialized with val
func NewStrikePrice(val float64) *StrikePriceField {
	field := &StrikePriceField{}
	field.Value = val
	return field
}

//StrikePriceBoundaryMethodField is a INT field
type StrikePriceBoundaryMethodField struct{ quickfix.IntValue }

//Tag returns tag.StrikePriceBoundaryMethod (1479)
func (f StrikePriceBoundaryMethodField) Tag() quickfix.Tag { return tag.StrikePriceBoundaryMethod }

//NewStrikePriceBoundaryMethod returns a new StrikePriceBoundaryMethodField initialized with val
func NewStrikePriceBoundaryMethod(val int) *StrikePriceBoundaryMethodField {
	field := &StrikePriceBoundaryMethodField{}
	field.Value = val
	return field
}

//StrikePriceBoundaryPrecisionField is a PERCENTAGE field
type StrikePriceBoundaryPrecisionField struct{ quickfix.PercentageValue }

//Tag returns tag.StrikePriceBoundaryPrecision (1480)
func (f StrikePriceBoundaryPrecisionField) Tag() quickfix.Tag { return tag.StrikePriceBoundaryPrecision }

//NewStrikePriceBoundaryPrecision returns a new StrikePriceBoundaryPrecisionField initialized with val
func NewStrikePriceBoundaryPrecision(val float64) *StrikePriceBoundaryPrecisionField {
	field := &StrikePriceBoundaryPrecisionField{}
	field.Value = val
	return field
}

//StrikePriceDeterminationMethodField is a INT field
type StrikePriceDeterminationMethodField struct{ quickfix.IntValue }

//Tag returns tag.StrikePriceDeterminationMethod (1478)
func (f StrikePriceDeterminationMethodField) Tag() quickfix.Tag {
	return tag.StrikePriceDeterminationMethod
}

//NewStrikePriceDeterminationMethod returns a new StrikePriceDeterminationMethodField initialized with val
func NewStrikePriceDeterminationMethod(val int) *StrikePriceDeterminationMethodField {
	field := &StrikePriceDeterminationMethodField{}
	field.Value = val
	return field
}

//StrikeRuleIDField is a STRING field
type StrikeRuleIDField struct{ quickfix.StringValue }

//Tag returns tag.StrikeRuleID (1223)
func (f StrikeRuleIDField) Tag() quickfix.Tag { return tag.StrikeRuleID }

//NewStrikeRuleID returns a new StrikeRuleIDField initialized with val
func NewStrikeRuleID(val string) *StrikeRuleIDField {
	field := &StrikeRuleIDField{}
	field.Value = val
	return field
}

//StrikeTimeField is a UTCTIMESTAMP field
type StrikeTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.StrikeTime (443)
func (f StrikeTimeField) Tag() quickfix.Tag { return tag.StrikeTime }

//StrikeValueField is a FLOAT field
type StrikeValueField struct{ quickfix.FloatValue }

//Tag returns tag.StrikeValue (968)
func (f StrikeValueField) Tag() quickfix.Tag { return tag.StrikeValue }

//NewStrikeValue returns a new StrikeValueField initialized with val
func NewStrikeValue(val float64) *StrikeValueField {
	field := &StrikeValueField{}
	field.Value = val
	return field
}

//SubjectField is a STRING field
type SubjectField struct{ quickfix.StringValue }

//Tag returns tag.Subject (147)
func (f SubjectField) Tag() quickfix.Tag { return tag.Subject }

//NewSubject returns a new SubjectField initialized with val
func NewSubject(val string) *SubjectField {
	field := &SubjectField{}
	field.Value = val
	return field
}

//SubscriptionRequestTypeField is a CHAR field
type SubscriptionRequestTypeField struct{ quickfix.CharValue }

//Tag returns tag.SubscriptionRequestType (263)
func (f SubscriptionRequestTypeField) Tag() quickfix.Tag { return tag.SubscriptionRequestType }

//NewSubscriptionRequestType returns a new SubscriptionRequestTypeField initialized with val
func NewSubscriptionRequestType(val string) *SubscriptionRequestTypeField {
	field := &SubscriptionRequestTypeField{}
	field.Value = val
	return field
}

//SwapPointsField is a PRICEOFFSET field
type SwapPointsField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.SwapPoints (1069)
func (f SwapPointsField) Tag() quickfix.Tag { return tag.SwapPoints }

//NewSwapPoints returns a new SwapPointsField initialized with val
func NewSwapPoints(val float64) *SwapPointsField {
	field := &SwapPointsField{}
	field.Value = val
	return field
}

//SymbolField is a STRING field
type SymbolField struct{ quickfix.StringValue }

//Tag returns tag.Symbol (55)
func (f SymbolField) Tag() quickfix.Tag { return tag.Symbol }

//NewSymbol returns a new SymbolField initialized with val
func NewSymbol(val string) *SymbolField {
	field := &SymbolField{}
	field.Value = val
	return field
}

//SymbolSfxField is a STRING field
type SymbolSfxField struct{ quickfix.StringValue }

//Tag returns tag.SymbolSfx (65)
func (f SymbolSfxField) Tag() quickfix.Tag { return tag.SymbolSfx }

//NewSymbolSfx returns a new SymbolSfxField initialized with val
func NewSymbolSfx(val string) *SymbolSfxField {
	field := &SymbolSfxField{}
	field.Value = val
	return field
}

//TZTransactTimeField is a TZTIMESTAMP field
type TZTransactTimeField struct{ quickfix.TZTimestampValue }

//Tag returns tag.TZTransactTime (1132)
func (f TZTransactTimeField) Tag() quickfix.Tag { return tag.TZTransactTime }

//TargetCompIDField is a STRING field
type TargetCompIDField struct{ quickfix.StringValue }

//Tag returns tag.TargetCompID (56)
func (f TargetCompIDField) Tag() quickfix.Tag { return tag.TargetCompID }

//NewTargetCompID returns a new TargetCompIDField initialized with val
func NewTargetCompID(val string) *TargetCompIDField {
	field := &TargetCompIDField{}
	field.Value = val
	return field
}

//TargetLocationIDField is a STRING field
type TargetLocationIDField struct{ quickfix.StringValue }

//Tag returns tag.TargetLocationID (143)
func (f TargetLocationIDField) Tag() quickfix.Tag { return tag.TargetLocationID }

//NewTargetLocationID returns a new TargetLocationIDField initialized with val
func NewTargetLocationID(val string) *TargetLocationIDField {
	field := &TargetLocationIDField{}
	field.Value = val
	return field
}

//TargetPartyIDField is a STRING field
type TargetPartyIDField struct{ quickfix.StringValue }

//Tag returns tag.TargetPartyID (1462)
func (f TargetPartyIDField) Tag() quickfix.Tag { return tag.TargetPartyID }

//NewTargetPartyID returns a new TargetPartyIDField initialized with val
func NewTargetPartyID(val string) *TargetPartyIDField {
	field := &TargetPartyIDField{}
	field.Value = val
	return field
}

//TargetPartyIDSourceField is a CHAR field
type TargetPartyIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.TargetPartyIDSource (1463)
func (f TargetPartyIDSourceField) Tag() quickfix.Tag { return tag.TargetPartyIDSource }

//NewTargetPartyIDSource returns a new TargetPartyIDSourceField initialized with val
func NewTargetPartyIDSource(val string) *TargetPartyIDSourceField {
	field := &TargetPartyIDSourceField{}
	field.Value = val
	return field
}

//TargetPartyRoleField is a INT field
type TargetPartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.TargetPartyRole (1464)
func (f TargetPartyRoleField) Tag() quickfix.Tag { return tag.TargetPartyRole }

//NewTargetPartyRole returns a new TargetPartyRoleField initialized with val
func NewTargetPartyRole(val int) *TargetPartyRoleField {
	field := &TargetPartyRoleField{}
	field.Value = val
	return field
}

//TargetStrategyField is a INT field
type TargetStrategyField struct{ quickfix.IntValue }

//Tag returns tag.TargetStrategy (847)
func (f TargetStrategyField) Tag() quickfix.Tag { return tag.TargetStrategy }

//NewTargetStrategy returns a new TargetStrategyField initialized with val
func NewTargetStrategy(val int) *TargetStrategyField {
	field := &TargetStrategyField{}
	field.Value = val
	return field
}

//TargetStrategyParametersField is a STRING field
type TargetStrategyParametersField struct{ quickfix.StringValue }

//Tag returns tag.TargetStrategyParameters (848)
func (f TargetStrategyParametersField) Tag() quickfix.Tag { return tag.TargetStrategyParameters }

//NewTargetStrategyParameters returns a new TargetStrategyParametersField initialized with val
func NewTargetStrategyParameters(val string) *TargetStrategyParametersField {
	field := &TargetStrategyParametersField{}
	field.Value = val
	return field
}

//TargetStrategyPerformanceField is a FLOAT field
type TargetStrategyPerformanceField struct{ quickfix.FloatValue }

//Tag returns tag.TargetStrategyPerformance (850)
func (f TargetStrategyPerformanceField) Tag() quickfix.Tag { return tag.TargetStrategyPerformance }

//NewTargetStrategyPerformance returns a new TargetStrategyPerformanceField initialized with val
func NewTargetStrategyPerformance(val float64) *TargetStrategyPerformanceField {
	field := &TargetStrategyPerformanceField{}
	field.Value = val
	return field
}

//TargetSubIDField is a STRING field
type TargetSubIDField struct{ quickfix.StringValue }

//Tag returns tag.TargetSubID (57)
func (f TargetSubIDField) Tag() quickfix.Tag { return tag.TargetSubID }

//NewTargetSubID returns a new TargetSubIDField initialized with val
func NewTargetSubID(val string) *TargetSubIDField {
	field := &TargetSubIDField{}
	field.Value = val
	return field
}

//TaxAdvantageTypeField is a INT field
type TaxAdvantageTypeField struct{ quickfix.IntValue }

//Tag returns tag.TaxAdvantageType (495)
func (f TaxAdvantageTypeField) Tag() quickfix.Tag { return tag.TaxAdvantageType }

//NewTaxAdvantageType returns a new TaxAdvantageTypeField initialized with val
func NewTaxAdvantageType(val int) *TaxAdvantageTypeField {
	field := &TaxAdvantageTypeField{}
	field.Value = val
	return field
}

//TerminationTypeField is a INT field
type TerminationTypeField struct{ quickfix.IntValue }

//Tag returns tag.TerminationType (788)
func (f TerminationTypeField) Tag() quickfix.Tag { return tag.TerminationType }

//NewTerminationType returns a new TerminationTypeField initialized with val
func NewTerminationType(val int) *TerminationTypeField {
	field := &TerminationTypeField{}
	field.Value = val
	return field
}

//TestMessageIndicatorField is a BOOLEAN field
type TestMessageIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.TestMessageIndicator (464)
func (f TestMessageIndicatorField) Tag() quickfix.Tag { return tag.TestMessageIndicator }

//NewTestMessageIndicator returns a new TestMessageIndicatorField initialized with val
func NewTestMessageIndicator(val bool) *TestMessageIndicatorField {
	field := &TestMessageIndicatorField{}
	field.Value = val
	return field
}

//TestReqIDField is a STRING field
type TestReqIDField struct{ quickfix.StringValue }

//Tag returns tag.TestReqID (112)
func (f TestReqIDField) Tag() quickfix.Tag { return tag.TestReqID }

//NewTestReqID returns a new TestReqIDField initialized with val
func NewTestReqID(val string) *TestReqIDField {
	field := &TestReqIDField{}
	field.Value = val
	return field
}

//TextField is a STRING field
type TextField struct{ quickfix.StringValue }

//Tag returns tag.Text (58)
func (f TextField) Tag() quickfix.Tag { return tag.Text }

//NewText returns a new TextField initialized with val
func NewText(val string) *TextField {
	field := &TextField{}
	field.Value = val
	return field
}

//ThresholdAmountField is a PRICEOFFSET field
type ThresholdAmountField struct{ quickfix.PriceOffsetValue }

//Tag returns tag.ThresholdAmount (834)
func (f ThresholdAmountField) Tag() quickfix.Tag { return tag.ThresholdAmount }

//NewThresholdAmount returns a new ThresholdAmountField initialized with val
func NewThresholdAmount(val float64) *ThresholdAmountField {
	field := &ThresholdAmountField{}
	field.Value = val
	return field
}

//TickDirectionField is a CHAR field
type TickDirectionField struct{ quickfix.CharValue }

//Tag returns tag.TickDirection (274)
func (f TickDirectionField) Tag() quickfix.Tag { return tag.TickDirection }

//NewTickDirection returns a new TickDirectionField initialized with val
func NewTickDirection(val string) *TickDirectionField {
	field := &TickDirectionField{}
	field.Value = val
	return field
}

//TickIncrementField is a PRICE field
type TickIncrementField struct{ quickfix.PriceValue }

//Tag returns tag.TickIncrement (1208)
func (f TickIncrementField) Tag() quickfix.Tag { return tag.TickIncrement }

//NewTickIncrement returns a new TickIncrementField initialized with val
func NewTickIncrement(val float64) *TickIncrementField {
	field := &TickIncrementField{}
	field.Value = val
	return field
}

//TickRuleTypeField is a INT field
type TickRuleTypeField struct{ quickfix.IntValue }

//Tag returns tag.TickRuleType (1209)
func (f TickRuleTypeField) Tag() quickfix.Tag { return tag.TickRuleType }

//NewTickRuleType returns a new TickRuleTypeField initialized with val
func NewTickRuleType(val int) *TickRuleTypeField {
	field := &TickRuleTypeField{}
	field.Value = val
	return field
}

//TierCodeField is a STRING field
type TierCodeField struct{ quickfix.StringValue }

//Tag returns tag.TierCode (994)
func (f TierCodeField) Tag() quickfix.Tag { return tag.TierCode }

//NewTierCode returns a new TierCodeField initialized with val
func NewTierCode(val string) *TierCodeField {
	field := &TierCodeField{}
	field.Value = val
	return field
}

//TimeBracketField is a STRING field
type TimeBracketField struct{ quickfix.StringValue }

//Tag returns tag.TimeBracket (943)
func (f TimeBracketField) Tag() quickfix.Tag { return tag.TimeBracket }

//NewTimeBracket returns a new TimeBracketField initialized with val
func NewTimeBracket(val string) *TimeBracketField {
	field := &TimeBracketField{}
	field.Value = val
	return field
}

//TimeInForceField is a CHAR field
type TimeInForceField struct{ quickfix.CharValue }

//Tag returns tag.TimeInForce (59)
func (f TimeInForceField) Tag() quickfix.Tag { return tag.TimeInForce }

//NewTimeInForce returns a new TimeInForceField initialized with val
func NewTimeInForce(val string) *TimeInForceField {
	field := &TimeInForceField{}
	field.Value = val
	return field
}

//TimeToExpirationField is a FLOAT field
type TimeToExpirationField struct{ quickfix.FloatValue }

//Tag returns tag.TimeToExpiration (1189)
func (f TimeToExpirationField) Tag() quickfix.Tag { return tag.TimeToExpiration }

//NewTimeToExpiration returns a new TimeToExpirationField initialized with val
func NewTimeToExpiration(val float64) *TimeToExpirationField {
	field := &TimeToExpirationField{}
	field.Value = val
	return field
}

//TimeUnitField is a STRING field
type TimeUnitField struct{ quickfix.StringValue }

//Tag returns tag.TimeUnit (997)
func (f TimeUnitField) Tag() quickfix.Tag { return tag.TimeUnit }

//NewTimeUnit returns a new TimeUnitField initialized with val
func NewTimeUnit(val string) *TimeUnitField {
	field := &TimeUnitField{}
	field.Value = val
	return field
}

//TotNoAccQuotesField is a INT field
type TotNoAccQuotesField struct{ quickfix.IntValue }

//Tag returns tag.TotNoAccQuotes (1169)
func (f TotNoAccQuotesField) Tag() quickfix.Tag { return tag.TotNoAccQuotes }

//NewTotNoAccQuotes returns a new TotNoAccQuotesField initialized with val
func NewTotNoAccQuotes(val int) *TotNoAccQuotesField {
	field := &TotNoAccQuotesField{}
	field.Value = val
	return field
}

//TotNoAllocsField is a INT field
type TotNoAllocsField struct{ quickfix.IntValue }

//Tag returns tag.TotNoAllocs (892)
func (f TotNoAllocsField) Tag() quickfix.Tag { return tag.TotNoAllocs }

//NewTotNoAllocs returns a new TotNoAllocsField initialized with val
func NewTotNoAllocs(val int) *TotNoAllocsField {
	field := &TotNoAllocsField{}
	field.Value = val
	return field
}

//TotNoCxldQuotesField is a INT field
type TotNoCxldQuotesField struct{ quickfix.IntValue }

//Tag returns tag.TotNoCxldQuotes (1168)
func (f TotNoCxldQuotesField) Tag() quickfix.Tag { return tag.TotNoCxldQuotes }

//NewTotNoCxldQuotes returns a new TotNoCxldQuotesField initialized with val
func NewTotNoCxldQuotes(val int) *TotNoCxldQuotesField {
	field := &TotNoCxldQuotesField{}
	field.Value = val
	return field
}

//TotNoFillsField is a INT field
type TotNoFillsField struct{ quickfix.IntValue }

//Tag returns tag.TotNoFills (1361)
func (f TotNoFillsField) Tag() quickfix.Tag { return tag.TotNoFills }

//NewTotNoFills returns a new TotNoFillsField initialized with val
func NewTotNoFills(val int) *TotNoFillsField {
	field := &TotNoFillsField{}
	field.Value = val
	return field
}

//TotNoOrdersField is a INT field
type TotNoOrdersField struct{ quickfix.IntValue }

//Tag returns tag.TotNoOrders (68)
func (f TotNoOrdersField) Tag() quickfix.Tag { return tag.TotNoOrders }

//NewTotNoOrders returns a new TotNoOrdersField initialized with val
func NewTotNoOrders(val int) *TotNoOrdersField {
	field := &TotNoOrdersField{}
	field.Value = val
	return field
}

//TotNoPartyListField is a INT field
type TotNoPartyListField struct{ quickfix.IntValue }

//Tag returns tag.TotNoPartyList (1512)
func (f TotNoPartyListField) Tag() quickfix.Tag { return tag.TotNoPartyList }

//NewTotNoPartyList returns a new TotNoPartyListField initialized with val
func NewTotNoPartyList(val int) *TotNoPartyListField {
	field := &TotNoPartyListField{}
	field.Value = val
	return field
}

//TotNoQuoteEntriesField is a INT field
type TotNoQuoteEntriesField struct{ quickfix.IntValue }

//Tag returns tag.TotNoQuoteEntries (304)
func (f TotNoQuoteEntriesField) Tag() quickfix.Tag { return tag.TotNoQuoteEntries }

//NewTotNoQuoteEntries returns a new TotNoQuoteEntriesField initialized with val
func NewTotNoQuoteEntries(val int) *TotNoQuoteEntriesField {
	field := &TotNoQuoteEntriesField{}
	field.Value = val
	return field
}

//TotNoRejQuotesField is a INT field
type TotNoRejQuotesField struct{ quickfix.IntValue }

//Tag returns tag.TotNoRejQuotes (1170)
func (f TotNoRejQuotesField) Tag() quickfix.Tag { return tag.TotNoRejQuotes }

//NewTotNoRejQuotes returns a new TotNoRejQuotesField initialized with val
func NewTotNoRejQuotes(val int) *TotNoRejQuotesField {
	field := &TotNoRejQuotesField{}
	field.Value = val
	return field
}

//TotNoRelatedSymField is a INT field
type TotNoRelatedSymField struct{ quickfix.IntValue }

//Tag returns tag.TotNoRelatedSym (393)
func (f TotNoRelatedSymField) Tag() quickfix.Tag { return tag.TotNoRelatedSym }

//NewTotNoRelatedSym returns a new TotNoRelatedSymField initialized with val
func NewTotNoRelatedSym(val int) *TotNoRelatedSymField {
	field := &TotNoRelatedSymField{}
	field.Value = val
	return field
}

//TotNoSecurityTypesField is a INT field
type TotNoSecurityTypesField struct{ quickfix.IntValue }

//Tag returns tag.TotNoSecurityTypes (557)
func (f TotNoSecurityTypesField) Tag() quickfix.Tag { return tag.TotNoSecurityTypes }

//NewTotNoSecurityTypes returns a new TotNoSecurityTypesField initialized with val
func NewTotNoSecurityTypes(val int) *TotNoSecurityTypesField {
	field := &TotNoSecurityTypesField{}
	field.Value = val
	return field
}

//TotNoStrikesField is a INT field
type TotNoStrikesField struct{ quickfix.IntValue }

//Tag returns tag.TotNoStrikes (422)
func (f TotNoStrikesField) Tag() quickfix.Tag { return tag.TotNoStrikes }

//NewTotNoStrikes returns a new TotNoStrikesField initialized with val
func NewTotNoStrikes(val int) *TotNoStrikesField {
	field := &TotNoStrikesField{}
	field.Value = val
	return field
}

//TotNumAssignmentReportsField is a INT field
type TotNumAssignmentReportsField struct{ quickfix.IntValue }

//Tag returns tag.TotNumAssignmentReports (832)
func (f TotNumAssignmentReportsField) Tag() quickfix.Tag { return tag.TotNumAssignmentReports }

//NewTotNumAssignmentReports returns a new TotNumAssignmentReportsField initialized with val
func NewTotNumAssignmentReports(val int) *TotNumAssignmentReportsField {
	field := &TotNumAssignmentReportsField{}
	field.Value = val
	return field
}

//TotNumReportsField is a INT field
type TotNumReportsField struct{ quickfix.IntValue }

//Tag returns tag.TotNumReports (911)
func (f TotNumReportsField) Tag() quickfix.Tag { return tag.TotNumReports }

//NewTotNumReports returns a new TotNumReportsField initialized with val
func NewTotNumReports(val int) *TotNumReportsField {
	field := &TotNumReportsField{}
	field.Value = val
	return field
}

//TotNumTradeReportsField is a INT field
type TotNumTradeReportsField struct{ quickfix.IntValue }

//Tag returns tag.TotNumTradeReports (748)
func (f TotNumTradeReportsField) Tag() quickfix.Tag { return tag.TotNumTradeReports }

//NewTotNumTradeReports returns a new TotNumTradeReportsField initialized with val
func NewTotNumTradeReports(val int) *TotNumTradeReportsField {
	field := &TotNumTradeReportsField{}
	field.Value = val
	return field
}

//TotQuoteEntriesField is a INT field
type TotQuoteEntriesField struct{ quickfix.IntValue }

//Tag returns tag.TotQuoteEntries (304)
func (f TotQuoteEntriesField) Tag() quickfix.Tag { return tag.TotQuoteEntries }

//NewTotQuoteEntries returns a new TotQuoteEntriesField initialized with val
func NewTotQuoteEntries(val int) *TotQuoteEntriesField {
	field := &TotQuoteEntriesField{}
	field.Value = val
	return field
}

//TotalAccruedInterestAmtField is a AMT field
type TotalAccruedInterestAmtField struct{ quickfix.AmtValue }

//Tag returns tag.TotalAccruedInterestAmt (540)
func (f TotalAccruedInterestAmtField) Tag() quickfix.Tag { return tag.TotalAccruedInterestAmt }

//NewTotalAccruedInterestAmt returns a new TotalAccruedInterestAmtField initialized with val
func NewTotalAccruedInterestAmt(val float64) *TotalAccruedInterestAmtField {
	field := &TotalAccruedInterestAmtField{}
	field.Value = val
	return field
}

//TotalAffectedOrdersField is a INT field
type TotalAffectedOrdersField struct{ quickfix.IntValue }

//Tag returns tag.TotalAffectedOrders (533)
func (f TotalAffectedOrdersField) Tag() quickfix.Tag { return tag.TotalAffectedOrders }

//NewTotalAffectedOrders returns a new TotalAffectedOrdersField initialized with val
func NewTotalAffectedOrders(val int) *TotalAffectedOrdersField {
	field := &TotalAffectedOrdersField{}
	field.Value = val
	return field
}

//TotalNetValueField is a AMT field
type TotalNetValueField struct{ quickfix.AmtValue }

//Tag returns tag.TotalNetValue (900)
func (f TotalNetValueField) Tag() quickfix.Tag { return tag.TotalNetValue }

//NewTotalNetValue returns a new TotalNetValueField initialized with val
func NewTotalNetValue(val float64) *TotalNetValueField {
	field := &TotalNetValueField{}
	field.Value = val
	return field
}

//TotalNumPosReportsField is a INT field
type TotalNumPosReportsField struct{ quickfix.IntValue }

//Tag returns tag.TotalNumPosReports (727)
func (f TotalNumPosReportsField) Tag() quickfix.Tag { return tag.TotalNumPosReports }

//NewTotalNumPosReports returns a new TotalNumPosReportsField initialized with val
func NewTotalNumPosReports(val int) *TotalNumPosReportsField {
	field := &TotalNumPosReportsField{}
	field.Value = val
	return field
}

//TotalNumSecuritiesField is a INT field
type TotalNumSecuritiesField struct{ quickfix.IntValue }

//Tag returns tag.TotalNumSecurities (393)
func (f TotalNumSecuritiesField) Tag() quickfix.Tag { return tag.TotalNumSecurities }

//NewTotalNumSecurities returns a new TotalNumSecuritiesField initialized with val
func NewTotalNumSecurities(val int) *TotalNumSecuritiesField {
	field := &TotalNumSecuritiesField{}
	field.Value = val
	return field
}

//TotalNumSecurityTypesField is a INT field
type TotalNumSecurityTypesField struct{ quickfix.IntValue }

//Tag returns tag.TotalNumSecurityTypes (557)
func (f TotalNumSecurityTypesField) Tag() quickfix.Tag { return tag.TotalNumSecurityTypes }

//NewTotalNumSecurityTypes returns a new TotalNumSecurityTypesField initialized with val
func NewTotalNumSecurityTypes(val int) *TotalNumSecurityTypesField {
	field := &TotalNumSecurityTypesField{}
	field.Value = val
	return field
}

//TotalTakedownField is a AMT field
type TotalTakedownField struct{ quickfix.AmtValue }

//Tag returns tag.TotalTakedown (237)
func (f TotalTakedownField) Tag() quickfix.Tag { return tag.TotalTakedown }

//NewTotalTakedown returns a new TotalTakedownField initialized with val
func NewTotalTakedown(val float64) *TotalTakedownField {
	field := &TotalTakedownField{}
	field.Value = val
	return field
}

//TotalVolumeTradedField is a QTY field
type TotalVolumeTradedField struct{ quickfix.QtyValue }

//Tag returns tag.TotalVolumeTraded (387)
func (f TotalVolumeTradedField) Tag() quickfix.Tag { return tag.TotalVolumeTraded }

//NewTotalVolumeTraded returns a new TotalVolumeTradedField initialized with val
func NewTotalVolumeTraded(val float64) *TotalVolumeTradedField {
	field := &TotalVolumeTradedField{}
	field.Value = val
	return field
}

//TotalVolumeTradedDateField is a UTCDATEONLY field
type TotalVolumeTradedDateField struct{ quickfix.UTCDateOnlyValue }

//Tag returns tag.TotalVolumeTradedDate (449)
func (f TotalVolumeTradedDateField) Tag() quickfix.Tag { return tag.TotalVolumeTradedDate }

//TotalVolumeTradedTimeField is a UTCTIMEONLY field
type TotalVolumeTradedTimeField struct{ quickfix.UTCTimeOnlyValue }

//Tag returns tag.TotalVolumeTradedTime (450)
func (f TotalVolumeTradedTimeField) Tag() quickfix.Tag { return tag.TotalVolumeTradedTime }

//TradSesCloseTimeField is a UTCTIMESTAMP field
type TradSesCloseTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.TradSesCloseTime (344)
func (f TradSesCloseTimeField) Tag() quickfix.Tag { return tag.TradSesCloseTime }

//TradSesEndTimeField is a UTCTIMESTAMP field
type TradSesEndTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.TradSesEndTime (345)
func (f TradSesEndTimeField) Tag() quickfix.Tag { return tag.TradSesEndTime }

//TradSesEventField is a INT field
type TradSesEventField struct{ quickfix.IntValue }

//Tag returns tag.TradSesEvent (1368)
func (f TradSesEventField) Tag() quickfix.Tag { return tag.TradSesEvent }

//NewTradSesEvent returns a new TradSesEventField initialized with val
func NewTradSesEvent(val int) *TradSesEventField {
	field := &TradSesEventField{}
	field.Value = val
	return field
}

//TradSesMethodField is a INT field
type TradSesMethodField struct{ quickfix.IntValue }

//Tag returns tag.TradSesMethod (338)
func (f TradSesMethodField) Tag() quickfix.Tag { return tag.TradSesMethod }

//NewTradSesMethod returns a new TradSesMethodField initialized with val
func NewTradSesMethod(val int) *TradSesMethodField {
	field := &TradSesMethodField{}
	field.Value = val
	return field
}

//TradSesModeField is a INT field
type TradSesModeField struct{ quickfix.IntValue }

//Tag returns tag.TradSesMode (339)
func (f TradSesModeField) Tag() quickfix.Tag { return tag.TradSesMode }

//NewTradSesMode returns a new TradSesModeField initialized with val
func NewTradSesMode(val int) *TradSesModeField {
	field := &TradSesModeField{}
	field.Value = val
	return field
}

//TradSesOpenTimeField is a UTCTIMESTAMP field
type TradSesOpenTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.TradSesOpenTime (342)
func (f TradSesOpenTimeField) Tag() quickfix.Tag { return tag.TradSesOpenTime }

//TradSesPreCloseTimeField is a UTCTIMESTAMP field
type TradSesPreCloseTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.TradSesPreCloseTime (343)
func (f TradSesPreCloseTimeField) Tag() quickfix.Tag { return tag.TradSesPreCloseTime }

//TradSesReqIDField is a STRING field
type TradSesReqIDField struct{ quickfix.StringValue }

//Tag returns tag.TradSesReqID (335)
func (f TradSesReqIDField) Tag() quickfix.Tag { return tag.TradSesReqID }

//NewTradSesReqID returns a new TradSesReqIDField initialized with val
func NewTradSesReqID(val string) *TradSesReqIDField {
	field := &TradSesReqIDField{}
	field.Value = val
	return field
}

//TradSesStartTimeField is a UTCTIMESTAMP field
type TradSesStartTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.TradSesStartTime (341)
func (f TradSesStartTimeField) Tag() quickfix.Tag { return tag.TradSesStartTime }

//TradSesStatusField is a INT field
type TradSesStatusField struct{ quickfix.IntValue }

//Tag returns tag.TradSesStatus (340)
func (f TradSesStatusField) Tag() quickfix.Tag { return tag.TradSesStatus }

//NewTradSesStatus returns a new TradSesStatusField initialized with val
func NewTradSesStatus(val int) *TradSesStatusField {
	field := &TradSesStatusField{}
	field.Value = val
	return field
}

//TradSesStatusRejReasonField is a INT field
type TradSesStatusRejReasonField struct{ quickfix.IntValue }

//Tag returns tag.TradSesStatusRejReason (567)
func (f TradSesStatusRejReasonField) Tag() quickfix.Tag { return tag.TradSesStatusRejReason }

//NewTradSesStatusRejReason returns a new TradSesStatusRejReasonField initialized with val
func NewTradSesStatusRejReason(val int) *TradSesStatusRejReasonField {
	field := &TradSesStatusRejReasonField{}
	field.Value = val
	return field
}

//TradSesUpdateActionField is a CHAR field
type TradSesUpdateActionField struct{ quickfix.CharValue }

//Tag returns tag.TradSesUpdateAction (1327)
func (f TradSesUpdateActionField) Tag() quickfix.Tag { return tag.TradSesUpdateAction }

//NewTradSesUpdateAction returns a new TradSesUpdateActionField initialized with val
func NewTradSesUpdateAction(val string) *TradSesUpdateActionField {
	field := &TradSesUpdateActionField{}
	field.Value = val
	return field
}

//TradeAllocIndicatorField is a INT field
type TradeAllocIndicatorField struct{ quickfix.IntValue }

//Tag returns tag.TradeAllocIndicator (826)
func (f TradeAllocIndicatorField) Tag() quickfix.Tag { return tag.TradeAllocIndicator }

//NewTradeAllocIndicator returns a new TradeAllocIndicatorField initialized with val
func NewTradeAllocIndicator(val int) *TradeAllocIndicatorField {
	field := &TradeAllocIndicatorField{}
	field.Value = val
	return field
}

//TradeConditionField is a MULTIPLESTRINGVALUE field
type TradeConditionField struct{ quickfix.MultipleStringValue }

//Tag returns tag.TradeCondition (277)
func (f TradeConditionField) Tag() quickfix.Tag { return tag.TradeCondition }

//NewTradeCondition returns a new TradeConditionField initialized with val
func NewTradeCondition(val string) *TradeConditionField {
	field := &TradeConditionField{}
	field.Value = val
	return field
}

//TradeDateField is a LOCALMKTDATE field
type TradeDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.TradeDate (75)
func (f TradeDateField) Tag() quickfix.Tag { return tag.TradeDate }

//NewTradeDate returns a new TradeDateField initialized with val
func NewTradeDate(val string) *TradeDateField {
	field := &TradeDateField{}
	field.Value = val
	return field
}

//TradeHandlingInstrField is a CHAR field
type TradeHandlingInstrField struct{ quickfix.CharValue }

//Tag returns tag.TradeHandlingInstr (1123)
func (f TradeHandlingInstrField) Tag() quickfix.Tag { return tag.TradeHandlingInstr }

//NewTradeHandlingInstr returns a new TradeHandlingInstrField initialized with val
func NewTradeHandlingInstr(val string) *TradeHandlingInstrField {
	field := &TradeHandlingInstrField{}
	field.Value = val
	return field
}

//TradeIDField is a STRING field
type TradeIDField struct{ quickfix.StringValue }

//Tag returns tag.TradeID (1003)
func (f TradeIDField) Tag() quickfix.Tag { return tag.TradeID }

//NewTradeID returns a new TradeIDField initialized with val
func NewTradeID(val string) *TradeIDField {
	field := &TradeIDField{}
	field.Value = val
	return field
}

//TradeInputDeviceField is a STRING field
type TradeInputDeviceField struct{ quickfix.StringValue }

//Tag returns tag.TradeInputDevice (579)
func (f TradeInputDeviceField) Tag() quickfix.Tag { return tag.TradeInputDevice }

//NewTradeInputDevice returns a new TradeInputDeviceField initialized with val
func NewTradeInputDevice(val string) *TradeInputDeviceField {
	field := &TradeInputDeviceField{}
	field.Value = val
	return field
}

//TradeInputSourceField is a STRING field
type TradeInputSourceField struct{ quickfix.StringValue }

//Tag returns tag.TradeInputSource (578)
func (f TradeInputSourceField) Tag() quickfix.Tag { return tag.TradeInputSource }

//NewTradeInputSource returns a new TradeInputSourceField initialized with val
func NewTradeInputSource(val string) *TradeInputSourceField {
	field := &TradeInputSourceField{}
	field.Value = val
	return field
}

//TradeLegRefIDField is a STRING field
type TradeLegRefIDField struct{ quickfix.StringValue }

//Tag returns tag.TradeLegRefID (824)
func (f TradeLegRefIDField) Tag() quickfix.Tag { return tag.TradeLegRefID }

//NewTradeLegRefID returns a new TradeLegRefIDField initialized with val
func NewTradeLegRefID(val string) *TradeLegRefIDField {
	field := &TradeLegRefIDField{}
	field.Value = val
	return field
}

//TradeLinkIDField is a STRING field
type TradeLinkIDField struct{ quickfix.StringValue }

//Tag returns tag.TradeLinkID (820)
func (f TradeLinkIDField) Tag() quickfix.Tag { return tag.TradeLinkID }

//NewTradeLinkID returns a new TradeLinkIDField initialized with val
func NewTradeLinkID(val string) *TradeLinkIDField {
	field := &TradeLinkIDField{}
	field.Value = val
	return field
}

//TradeOriginationDateField is a LOCALMKTDATE field
type TradeOriginationDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.TradeOriginationDate (229)
func (f TradeOriginationDateField) Tag() quickfix.Tag { return tag.TradeOriginationDate }

//NewTradeOriginationDate returns a new TradeOriginationDateField initialized with val
func NewTradeOriginationDate(val string) *TradeOriginationDateField {
	field := &TradeOriginationDateField{}
	field.Value = val
	return field
}

//TradePublishIndicatorField is a INT field
type TradePublishIndicatorField struct{ quickfix.IntValue }

//Tag returns tag.TradePublishIndicator (1390)
func (f TradePublishIndicatorField) Tag() quickfix.Tag { return tag.TradePublishIndicator }

//NewTradePublishIndicator returns a new TradePublishIndicatorField initialized with val
func NewTradePublishIndicator(val int) *TradePublishIndicatorField {
	field := &TradePublishIndicatorField{}
	field.Value = val
	return field
}

//TradeReportIDField is a STRING field
type TradeReportIDField struct{ quickfix.StringValue }

//Tag returns tag.TradeReportID (571)
func (f TradeReportIDField) Tag() quickfix.Tag { return tag.TradeReportID }

//NewTradeReportID returns a new TradeReportIDField initialized with val
func NewTradeReportID(val string) *TradeReportIDField {
	field := &TradeReportIDField{}
	field.Value = val
	return field
}

//TradeReportRefIDField is a STRING field
type TradeReportRefIDField struct{ quickfix.StringValue }

//Tag returns tag.TradeReportRefID (572)
func (f TradeReportRefIDField) Tag() quickfix.Tag { return tag.TradeReportRefID }

//NewTradeReportRefID returns a new TradeReportRefIDField initialized with val
func NewTradeReportRefID(val string) *TradeReportRefIDField {
	field := &TradeReportRefIDField{}
	field.Value = val
	return field
}

//TradeReportRejectReasonField is a INT field
type TradeReportRejectReasonField struct{ quickfix.IntValue }

//Tag returns tag.TradeReportRejectReason (751)
func (f TradeReportRejectReasonField) Tag() quickfix.Tag { return tag.TradeReportRejectReason }

//NewTradeReportRejectReason returns a new TradeReportRejectReasonField initialized with val
func NewTradeReportRejectReason(val int) *TradeReportRejectReasonField {
	field := &TradeReportRejectReasonField{}
	field.Value = val
	return field
}

//TradeReportTransTypeField is a INT field
type TradeReportTransTypeField struct{ quickfix.IntValue }

//Tag returns tag.TradeReportTransType (487)
func (f TradeReportTransTypeField) Tag() quickfix.Tag { return tag.TradeReportTransType }

//NewTradeReportTransType returns a new TradeReportTransTypeField initialized with val
func NewTradeReportTransType(val int) *TradeReportTransTypeField {
	field := &TradeReportTransTypeField{}
	field.Value = val
	return field
}

//TradeReportTypeField is a INT field
type TradeReportTypeField struct{ quickfix.IntValue }

//Tag returns tag.TradeReportType (856)
func (f TradeReportTypeField) Tag() quickfix.Tag { return tag.TradeReportType }

//NewTradeReportType returns a new TradeReportTypeField initialized with val
func NewTradeReportType(val int) *TradeReportTypeField {
	field := &TradeReportTypeField{}
	field.Value = val
	return field
}

//TradeRequestIDField is a STRING field
type TradeRequestIDField struct{ quickfix.StringValue }

//Tag returns tag.TradeRequestID (568)
func (f TradeRequestIDField) Tag() quickfix.Tag { return tag.TradeRequestID }

//NewTradeRequestID returns a new TradeRequestIDField initialized with val
func NewTradeRequestID(val string) *TradeRequestIDField {
	field := &TradeRequestIDField{}
	field.Value = val
	return field
}

//TradeRequestResultField is a INT field
type TradeRequestResultField struct{ quickfix.IntValue }

//Tag returns tag.TradeRequestResult (749)
func (f TradeRequestResultField) Tag() quickfix.Tag { return tag.TradeRequestResult }

//NewTradeRequestResult returns a new TradeRequestResultField initialized with val
func NewTradeRequestResult(val int) *TradeRequestResultField {
	field := &TradeRequestResultField{}
	field.Value = val
	return field
}

//TradeRequestStatusField is a INT field
type TradeRequestStatusField struct{ quickfix.IntValue }

//Tag returns tag.TradeRequestStatus (750)
func (f TradeRequestStatusField) Tag() quickfix.Tag { return tag.TradeRequestStatus }

//NewTradeRequestStatus returns a new TradeRequestStatusField initialized with val
func NewTradeRequestStatus(val int) *TradeRequestStatusField {
	field := &TradeRequestStatusField{}
	field.Value = val
	return field
}

//TradeRequestTypeField is a INT field
type TradeRequestTypeField struct{ quickfix.IntValue }

//Tag returns tag.TradeRequestType (569)
func (f TradeRequestTypeField) Tag() quickfix.Tag { return tag.TradeRequestType }

//NewTradeRequestType returns a new TradeRequestTypeField initialized with val
func NewTradeRequestType(val int) *TradeRequestTypeField {
	field := &TradeRequestTypeField{}
	field.Value = val
	return field
}

//TradeTypeField is a CHAR field
type TradeTypeField struct{ quickfix.CharValue }

//Tag returns tag.TradeType (418)
func (f TradeTypeField) Tag() quickfix.Tag { return tag.TradeType }

//NewTradeType returns a new TradeTypeField initialized with val
func NewTradeType(val string) *TradeTypeField {
	field := &TradeTypeField{}
	field.Value = val
	return field
}

//TradeVolumeField is a QTY field
type TradeVolumeField struct{ quickfix.QtyValue }

//Tag returns tag.TradeVolume (1020)
func (f TradeVolumeField) Tag() quickfix.Tag { return tag.TradeVolume }

//NewTradeVolume returns a new TradeVolumeField initialized with val
func NewTradeVolume(val float64) *TradeVolumeField {
	field := &TradeVolumeField{}
	field.Value = val
	return field
}

//TradedFlatSwitchField is a BOOLEAN field
type TradedFlatSwitchField struct{ quickfix.BooleanValue }

//Tag returns tag.TradedFlatSwitch (258)
func (f TradedFlatSwitchField) Tag() quickfix.Tag { return tag.TradedFlatSwitch }

//NewTradedFlatSwitch returns a new TradedFlatSwitchField initialized with val
func NewTradedFlatSwitch(val bool) *TradedFlatSwitchField {
	field := &TradedFlatSwitchField{}
	field.Value = val
	return field
}

//TradingCurrencyField is a CURRENCY field
type TradingCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.TradingCurrency (1245)
func (f TradingCurrencyField) Tag() quickfix.Tag { return tag.TradingCurrency }

//NewTradingCurrency returns a new TradingCurrencyField initialized with val
func NewTradingCurrency(val string) *TradingCurrencyField {
	field := &TradingCurrencyField{}
	field.Value = val
	return field
}

//TradingReferencePriceField is a PRICE field
type TradingReferencePriceField struct{ quickfix.PriceValue }

//Tag returns tag.TradingReferencePrice (1150)
func (f TradingReferencePriceField) Tag() quickfix.Tag { return tag.TradingReferencePrice }

//NewTradingReferencePrice returns a new TradingReferencePriceField initialized with val
func NewTradingReferencePrice(val float64) *TradingReferencePriceField {
	field := &TradingReferencePriceField{}
	field.Value = val
	return field
}

//TradingSessionDescField is a STRING field
type TradingSessionDescField struct{ quickfix.StringValue }

//Tag returns tag.TradingSessionDesc (1326)
func (f TradingSessionDescField) Tag() quickfix.Tag { return tag.TradingSessionDesc }

//NewTradingSessionDesc returns a new TradingSessionDescField initialized with val
func NewTradingSessionDesc(val string) *TradingSessionDescField {
	field := &TradingSessionDescField{}
	field.Value = val
	return field
}

//TradingSessionIDField is a STRING field
type TradingSessionIDField struct{ quickfix.StringValue }

//Tag returns tag.TradingSessionID (336)
func (f TradingSessionIDField) Tag() quickfix.Tag { return tag.TradingSessionID }

//NewTradingSessionID returns a new TradingSessionIDField initialized with val
func NewTradingSessionID(val string) *TradingSessionIDField {
	field := &TradingSessionIDField{}
	field.Value = val
	return field
}

//TradingSessionSubIDField is a STRING field
type TradingSessionSubIDField struct{ quickfix.StringValue }

//Tag returns tag.TradingSessionSubID (625)
func (f TradingSessionSubIDField) Tag() quickfix.Tag { return tag.TradingSessionSubID }

//NewTradingSessionSubID returns a new TradingSessionSubIDField initialized with val
func NewTradingSessionSubID(val string) *TradingSessionSubIDField {
	field := &TradingSessionSubIDField{}
	field.Value = val
	return field
}

//TransBkdTimeField is a UTCTIMESTAMP field
type TransBkdTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.TransBkdTime (483)
func (f TransBkdTimeField) Tag() quickfix.Tag { return tag.TransBkdTime }

//TransactTimeField is a UTCTIMESTAMP field
type TransactTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.TransactTime (60)
func (f TransactTimeField) Tag() quickfix.Tag { return tag.TransactTime }

//TransferReasonField is a STRING field
type TransferReasonField struct{ quickfix.StringValue }

//Tag returns tag.TransferReason (830)
func (f TransferReasonField) Tag() quickfix.Tag { return tag.TransferReason }

//NewTransferReason returns a new TransferReasonField initialized with val
func NewTransferReason(val string) *TransferReasonField {
	field := &TransferReasonField{}
	field.Value = val
	return field
}

//TrdMatchIDField is a STRING field
type TrdMatchIDField struct{ quickfix.StringValue }

//Tag returns tag.TrdMatchID (880)
func (f TrdMatchIDField) Tag() quickfix.Tag { return tag.TrdMatchID }

//NewTrdMatchID returns a new TrdMatchIDField initialized with val
func NewTrdMatchID(val string) *TrdMatchIDField {
	field := &TrdMatchIDField{}
	field.Value = val
	return field
}

//TrdRegTimestampField is a UTCTIMESTAMP field
type TrdRegTimestampField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.TrdRegTimestamp (769)
func (f TrdRegTimestampField) Tag() quickfix.Tag { return tag.TrdRegTimestamp }

//TrdRegTimestampOriginField is a STRING field
type TrdRegTimestampOriginField struct{ quickfix.StringValue }

//Tag returns tag.TrdRegTimestampOrigin (771)
func (f TrdRegTimestampOriginField) Tag() quickfix.Tag { return tag.TrdRegTimestampOrigin }

//NewTrdRegTimestampOrigin returns a new TrdRegTimestampOriginField initialized with val
func NewTrdRegTimestampOrigin(val string) *TrdRegTimestampOriginField {
	field := &TrdRegTimestampOriginField{}
	field.Value = val
	return field
}

//TrdRegTimestampTypeField is a INT field
type TrdRegTimestampTypeField struct{ quickfix.IntValue }

//Tag returns tag.TrdRegTimestampType (770)
func (f TrdRegTimestampTypeField) Tag() quickfix.Tag { return tag.TrdRegTimestampType }

//NewTrdRegTimestampType returns a new TrdRegTimestampTypeField initialized with val
func NewTrdRegTimestampType(val int) *TrdRegTimestampTypeField {
	field := &TrdRegTimestampTypeField{}
	field.Value = val
	return field
}

//TrdRepIndicatorField is a BOOLEAN field
type TrdRepIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.TrdRepIndicator (1389)
func (f TrdRepIndicatorField) Tag() quickfix.Tag { return tag.TrdRepIndicator }

//NewTrdRepIndicator returns a new TrdRepIndicatorField initialized with val
func NewTrdRepIndicator(val bool) *TrdRepIndicatorField {
	field := &TrdRepIndicatorField{}
	field.Value = val
	return field
}

//TrdRepPartyRoleField is a INT field
type TrdRepPartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.TrdRepPartyRole (1388)
func (f TrdRepPartyRoleField) Tag() quickfix.Tag { return tag.TrdRepPartyRole }

//NewTrdRepPartyRole returns a new TrdRepPartyRoleField initialized with val
func NewTrdRepPartyRole(val int) *TrdRepPartyRoleField {
	field := &TrdRepPartyRoleField{}
	field.Value = val
	return field
}

//TrdRptStatusField is a INT field
type TrdRptStatusField struct{ quickfix.IntValue }

//Tag returns tag.TrdRptStatus (939)
func (f TrdRptStatusField) Tag() quickfix.Tag { return tag.TrdRptStatus }

//NewTrdRptStatus returns a new TrdRptStatusField initialized with val
func NewTrdRptStatus(val int) *TrdRptStatusField {
	field := &TrdRptStatusField{}
	field.Value = val
	return field
}

//TrdSubTypeField is a INT field
type TrdSubTypeField struct{ quickfix.IntValue }

//Tag returns tag.TrdSubType (829)
func (f TrdSubTypeField) Tag() quickfix.Tag { return tag.TrdSubType }

//NewTrdSubType returns a new TrdSubTypeField initialized with val
func NewTrdSubType(val int) *TrdSubTypeField {
	field := &TrdSubTypeField{}
	field.Value = val
	return field
}

//TrdTypeField is a INT field
type TrdTypeField struct{ quickfix.IntValue }

//Tag returns tag.TrdType (828)
func (f TrdTypeField) Tag() quickfix.Tag { return tag.TrdType }

//NewTrdType returns a new TrdTypeField initialized with val
func NewTrdType(val int) *TrdTypeField {
	field := &TrdTypeField{}
	field.Value = val
	return field
}

//TriggerActionField is a CHAR field
type TriggerActionField struct{ quickfix.CharValue }

//Tag returns tag.TriggerAction (1101)
func (f TriggerActionField) Tag() quickfix.Tag { return tag.TriggerAction }

//NewTriggerAction returns a new TriggerActionField initialized with val
func NewTriggerAction(val string) *TriggerActionField {
	field := &TriggerActionField{}
	field.Value = val
	return field
}

//TriggerNewPriceField is a PRICE field
type TriggerNewPriceField struct{ quickfix.PriceValue }

//Tag returns tag.TriggerNewPrice (1110)
func (f TriggerNewPriceField) Tag() quickfix.Tag { return tag.TriggerNewPrice }

//NewTriggerNewPrice returns a new TriggerNewPriceField initialized with val
func NewTriggerNewPrice(val float64) *TriggerNewPriceField {
	field := &TriggerNewPriceField{}
	field.Value = val
	return field
}

//TriggerNewQtyField is a QTY field
type TriggerNewQtyField struct{ quickfix.QtyValue }

//Tag returns tag.TriggerNewQty (1112)
func (f TriggerNewQtyField) Tag() quickfix.Tag { return tag.TriggerNewQty }

//NewTriggerNewQty returns a new TriggerNewQtyField initialized with val
func NewTriggerNewQty(val float64) *TriggerNewQtyField {
	field := &TriggerNewQtyField{}
	field.Value = val
	return field
}

//TriggerOrderTypeField is a CHAR field
type TriggerOrderTypeField struct{ quickfix.CharValue }

//Tag returns tag.TriggerOrderType (1111)
func (f TriggerOrderTypeField) Tag() quickfix.Tag { return tag.TriggerOrderType }

//NewTriggerOrderType returns a new TriggerOrderTypeField initialized with val
func NewTriggerOrderType(val string) *TriggerOrderTypeField {
	field := &TriggerOrderTypeField{}
	field.Value = val
	return field
}

//TriggerPriceField is a PRICE field
type TriggerPriceField struct{ quickfix.PriceValue }

//Tag returns tag.TriggerPrice (1102)
func (f TriggerPriceField) Tag() quickfix.Tag { return tag.TriggerPrice }

//NewTriggerPrice returns a new TriggerPriceField initialized with val
func NewTriggerPrice(val float64) *TriggerPriceField {
	field := &TriggerPriceField{}
	field.Value = val
	return field
}

//TriggerPriceDirectionField is a CHAR field
type TriggerPriceDirectionField struct{ quickfix.CharValue }

//Tag returns tag.TriggerPriceDirection (1109)
func (f TriggerPriceDirectionField) Tag() quickfix.Tag { return tag.TriggerPriceDirection }

//NewTriggerPriceDirection returns a new TriggerPriceDirectionField initialized with val
func NewTriggerPriceDirection(val string) *TriggerPriceDirectionField {
	field := &TriggerPriceDirectionField{}
	field.Value = val
	return field
}

//TriggerPriceTypeField is a CHAR field
type TriggerPriceTypeField struct{ quickfix.CharValue }

//Tag returns tag.TriggerPriceType (1107)
func (f TriggerPriceTypeField) Tag() quickfix.Tag { return tag.TriggerPriceType }

//NewTriggerPriceType returns a new TriggerPriceTypeField initialized with val
func NewTriggerPriceType(val string) *TriggerPriceTypeField {
	field := &TriggerPriceTypeField{}
	field.Value = val
	return field
}

//TriggerPriceTypeScopeField is a CHAR field
type TriggerPriceTypeScopeField struct{ quickfix.CharValue }

//Tag returns tag.TriggerPriceTypeScope (1108)
func (f TriggerPriceTypeScopeField) Tag() quickfix.Tag { return tag.TriggerPriceTypeScope }

//NewTriggerPriceTypeScope returns a new TriggerPriceTypeScopeField initialized with val
func NewTriggerPriceTypeScope(val string) *TriggerPriceTypeScopeField {
	field := &TriggerPriceTypeScopeField{}
	field.Value = val
	return field
}

//TriggerSecurityDescField is a STRING field
type TriggerSecurityDescField struct{ quickfix.StringValue }

//Tag returns tag.TriggerSecurityDesc (1106)
func (f TriggerSecurityDescField) Tag() quickfix.Tag { return tag.TriggerSecurityDesc }

//NewTriggerSecurityDesc returns a new TriggerSecurityDescField initialized with val
func NewTriggerSecurityDesc(val string) *TriggerSecurityDescField {
	field := &TriggerSecurityDescField{}
	field.Value = val
	return field
}

//TriggerSecurityIDField is a STRING field
type TriggerSecurityIDField struct{ quickfix.StringValue }

//Tag returns tag.TriggerSecurityID (1104)
func (f TriggerSecurityIDField) Tag() quickfix.Tag { return tag.TriggerSecurityID }

//NewTriggerSecurityID returns a new TriggerSecurityIDField initialized with val
func NewTriggerSecurityID(val string) *TriggerSecurityIDField {
	field := &TriggerSecurityIDField{}
	field.Value = val
	return field
}

//TriggerSecurityIDSourceField is a STRING field
type TriggerSecurityIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.TriggerSecurityIDSource (1105)
func (f TriggerSecurityIDSourceField) Tag() quickfix.Tag { return tag.TriggerSecurityIDSource }

//NewTriggerSecurityIDSource returns a new TriggerSecurityIDSourceField initialized with val
func NewTriggerSecurityIDSource(val string) *TriggerSecurityIDSourceField {
	field := &TriggerSecurityIDSourceField{}
	field.Value = val
	return field
}

//TriggerSymbolField is a STRING field
type TriggerSymbolField struct{ quickfix.StringValue }

//Tag returns tag.TriggerSymbol (1103)
func (f TriggerSymbolField) Tag() quickfix.Tag { return tag.TriggerSymbol }

//NewTriggerSymbol returns a new TriggerSymbolField initialized with val
func NewTriggerSymbol(val string) *TriggerSymbolField {
	field := &TriggerSymbolField{}
	field.Value = val
	return field
}

//TriggerTradingSessionIDField is a STRING field
type TriggerTradingSessionIDField struct{ quickfix.StringValue }

//Tag returns tag.TriggerTradingSessionID (1113)
func (f TriggerTradingSessionIDField) Tag() quickfix.Tag { return tag.TriggerTradingSessionID }

//NewTriggerTradingSessionID returns a new TriggerTradingSessionIDField initialized with val
func NewTriggerTradingSessionID(val string) *TriggerTradingSessionIDField {
	field := &TriggerTradingSessionIDField{}
	field.Value = val
	return field
}

//TriggerTradingSessionSubIDField is a STRING field
type TriggerTradingSessionSubIDField struct{ quickfix.StringValue }

//Tag returns tag.TriggerTradingSessionSubID (1114)
func (f TriggerTradingSessionSubIDField) Tag() quickfix.Tag { return tag.TriggerTradingSessionSubID }

//NewTriggerTradingSessionSubID returns a new TriggerTradingSessionSubIDField initialized with val
func NewTriggerTradingSessionSubID(val string) *TriggerTradingSessionSubIDField {
	field := &TriggerTradingSessionSubIDField{}
	field.Value = val
	return field
}

//TriggerTypeField is a CHAR field
type TriggerTypeField struct{ quickfix.CharValue }

//Tag returns tag.TriggerType (1100)
func (f TriggerTypeField) Tag() quickfix.Tag { return tag.TriggerType }

//NewTriggerType returns a new TriggerTypeField initialized with val
func NewTriggerType(val string) *TriggerTypeField {
	field := &TriggerTypeField{}
	field.Value = val
	return field
}

//URLLinkField is a STRING field
type URLLinkField struct{ quickfix.StringValue }

//Tag returns tag.URLLink (149)
func (f URLLinkField) Tag() quickfix.Tag { return tag.URLLink }

//NewURLLink returns a new URLLinkField initialized with val
func NewURLLink(val string) *URLLinkField {
	field := &URLLinkField{}
	field.Value = val
	return field
}

//UnderlyingAdjustedQuantityField is a QTY field
type UnderlyingAdjustedQuantityField struct{ quickfix.QtyValue }

//Tag returns tag.UnderlyingAdjustedQuantity (1044)
func (f UnderlyingAdjustedQuantityField) Tag() quickfix.Tag { return tag.UnderlyingAdjustedQuantity }

//NewUnderlyingAdjustedQuantity returns a new UnderlyingAdjustedQuantityField initialized with val
func NewUnderlyingAdjustedQuantity(val float64) *UnderlyingAdjustedQuantityField {
	field := &UnderlyingAdjustedQuantityField{}
	field.Value = val
	return field
}

//UnderlyingAllocationPercentField is a PERCENTAGE field
type UnderlyingAllocationPercentField struct{ quickfix.PercentageValue }

//Tag returns tag.UnderlyingAllocationPercent (972)
func (f UnderlyingAllocationPercentField) Tag() quickfix.Tag { return tag.UnderlyingAllocationPercent }

//NewUnderlyingAllocationPercent returns a new UnderlyingAllocationPercentField initialized with val
func NewUnderlyingAllocationPercent(val float64) *UnderlyingAllocationPercentField {
	field := &UnderlyingAllocationPercentField{}
	field.Value = val
	return field
}

//UnderlyingAttachmentPointField is a PERCENTAGE field
type UnderlyingAttachmentPointField struct{ quickfix.PercentageValue }

//Tag returns tag.UnderlyingAttachmentPoint (1459)
func (f UnderlyingAttachmentPointField) Tag() quickfix.Tag { return tag.UnderlyingAttachmentPoint }

//NewUnderlyingAttachmentPoint returns a new UnderlyingAttachmentPointField initialized with val
func NewUnderlyingAttachmentPoint(val float64) *UnderlyingAttachmentPointField {
	field := &UnderlyingAttachmentPointField{}
	field.Value = val
	return field
}

//UnderlyingCFICodeField is a STRING field
type UnderlyingCFICodeField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingCFICode (463)
func (f UnderlyingCFICodeField) Tag() quickfix.Tag { return tag.UnderlyingCFICode }

//NewUnderlyingCFICode returns a new UnderlyingCFICodeField initialized with val
func NewUnderlyingCFICode(val string) *UnderlyingCFICodeField {
	field := &UnderlyingCFICodeField{}
	field.Value = val
	return field
}

//UnderlyingCPProgramField is a STRING field
type UnderlyingCPProgramField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingCPProgram (877)
func (f UnderlyingCPProgramField) Tag() quickfix.Tag { return tag.UnderlyingCPProgram }

//NewUnderlyingCPProgram returns a new UnderlyingCPProgramField initialized with val
func NewUnderlyingCPProgram(val string) *UnderlyingCPProgramField {
	field := &UnderlyingCPProgramField{}
	field.Value = val
	return field
}

//UnderlyingCPRegTypeField is a STRING field
type UnderlyingCPRegTypeField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingCPRegType (878)
func (f UnderlyingCPRegTypeField) Tag() quickfix.Tag { return tag.UnderlyingCPRegType }

//NewUnderlyingCPRegType returns a new UnderlyingCPRegTypeField initialized with val
func NewUnderlyingCPRegType(val string) *UnderlyingCPRegTypeField {
	field := &UnderlyingCPRegTypeField{}
	field.Value = val
	return field
}

//UnderlyingCapValueField is a AMT field
type UnderlyingCapValueField struct{ quickfix.AmtValue }

//Tag returns tag.UnderlyingCapValue (1038)
func (f UnderlyingCapValueField) Tag() quickfix.Tag { return tag.UnderlyingCapValue }

//NewUnderlyingCapValue returns a new UnderlyingCapValueField initialized with val
func NewUnderlyingCapValue(val float64) *UnderlyingCapValueField {
	field := &UnderlyingCapValueField{}
	field.Value = val
	return field
}

//UnderlyingCashAmountField is a AMT field
type UnderlyingCashAmountField struct{ quickfix.AmtValue }

//Tag returns tag.UnderlyingCashAmount (973)
func (f UnderlyingCashAmountField) Tag() quickfix.Tag { return tag.UnderlyingCashAmount }

//NewUnderlyingCashAmount returns a new UnderlyingCashAmountField initialized with val
func NewUnderlyingCashAmount(val float64) *UnderlyingCashAmountField {
	field := &UnderlyingCashAmountField{}
	field.Value = val
	return field
}

//UnderlyingCashTypeField is a STRING field
type UnderlyingCashTypeField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingCashType (974)
func (f UnderlyingCashTypeField) Tag() quickfix.Tag { return tag.UnderlyingCashType }

//NewUnderlyingCashType returns a new UnderlyingCashTypeField initialized with val
func NewUnderlyingCashType(val string) *UnderlyingCashTypeField {
	field := &UnderlyingCashTypeField{}
	field.Value = val
	return field
}

//UnderlyingCollectAmountField is a AMT field
type UnderlyingCollectAmountField struct{ quickfix.AmtValue }

//Tag returns tag.UnderlyingCollectAmount (986)
func (f UnderlyingCollectAmountField) Tag() quickfix.Tag { return tag.UnderlyingCollectAmount }

//NewUnderlyingCollectAmount returns a new UnderlyingCollectAmountField initialized with val
func NewUnderlyingCollectAmount(val float64) *UnderlyingCollectAmountField {
	field := &UnderlyingCollectAmountField{}
	field.Value = val
	return field
}

//UnderlyingContractMultiplierField is a FLOAT field
type UnderlyingContractMultiplierField struct{ quickfix.FloatValue }

//Tag returns tag.UnderlyingContractMultiplier (436)
func (f UnderlyingContractMultiplierField) Tag() quickfix.Tag { return tag.UnderlyingContractMultiplier }

//NewUnderlyingContractMultiplier returns a new UnderlyingContractMultiplierField initialized with val
func NewUnderlyingContractMultiplier(val float64) *UnderlyingContractMultiplierField {
	field := &UnderlyingContractMultiplierField{}
	field.Value = val
	return field
}

//UnderlyingContractMultiplierUnitField is a INT field
type UnderlyingContractMultiplierUnitField struct{ quickfix.IntValue }

//Tag returns tag.UnderlyingContractMultiplierUnit (1437)
func (f UnderlyingContractMultiplierUnitField) Tag() quickfix.Tag {
	return tag.UnderlyingContractMultiplierUnit
}

//NewUnderlyingContractMultiplierUnit returns a new UnderlyingContractMultiplierUnitField initialized with val
func NewUnderlyingContractMultiplierUnit(val int) *UnderlyingContractMultiplierUnitField {
	field := &UnderlyingContractMultiplierUnitField{}
	field.Value = val
	return field
}

//UnderlyingCountryOfIssueField is a COUNTRY field
type UnderlyingCountryOfIssueField struct{ quickfix.CountryValue }

//Tag returns tag.UnderlyingCountryOfIssue (592)
func (f UnderlyingCountryOfIssueField) Tag() quickfix.Tag { return tag.UnderlyingCountryOfIssue }

//NewUnderlyingCountryOfIssue returns a new UnderlyingCountryOfIssueField initialized with val
func NewUnderlyingCountryOfIssue(val string) *UnderlyingCountryOfIssueField {
	field := &UnderlyingCountryOfIssueField{}
	field.Value = val
	return field
}

//UnderlyingCouponPaymentDateField is a LOCALMKTDATE field
type UnderlyingCouponPaymentDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.UnderlyingCouponPaymentDate (241)
func (f UnderlyingCouponPaymentDateField) Tag() quickfix.Tag { return tag.UnderlyingCouponPaymentDate }

//NewUnderlyingCouponPaymentDate returns a new UnderlyingCouponPaymentDateField initialized with val
func NewUnderlyingCouponPaymentDate(val string) *UnderlyingCouponPaymentDateField {
	field := &UnderlyingCouponPaymentDateField{}
	field.Value = val
	return field
}

//UnderlyingCouponRateField is a PERCENTAGE field
type UnderlyingCouponRateField struct{ quickfix.PercentageValue }

//Tag returns tag.UnderlyingCouponRate (435)
func (f UnderlyingCouponRateField) Tag() quickfix.Tag { return tag.UnderlyingCouponRate }

//NewUnderlyingCouponRate returns a new UnderlyingCouponRateField initialized with val
func NewUnderlyingCouponRate(val float64) *UnderlyingCouponRateField {
	field := &UnderlyingCouponRateField{}
	field.Value = val
	return field
}

//UnderlyingCreditRatingField is a STRING field
type UnderlyingCreditRatingField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingCreditRating (256)
func (f UnderlyingCreditRatingField) Tag() quickfix.Tag { return tag.UnderlyingCreditRating }

//NewUnderlyingCreditRating returns a new UnderlyingCreditRatingField initialized with val
func NewUnderlyingCreditRating(val string) *UnderlyingCreditRatingField {
	field := &UnderlyingCreditRatingField{}
	field.Value = val
	return field
}

//UnderlyingCurrencyField is a CURRENCY field
type UnderlyingCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.UnderlyingCurrency (318)
func (f UnderlyingCurrencyField) Tag() quickfix.Tag { return tag.UnderlyingCurrency }

//NewUnderlyingCurrency returns a new UnderlyingCurrencyField initialized with val
func NewUnderlyingCurrency(val string) *UnderlyingCurrencyField {
	field := &UnderlyingCurrencyField{}
	field.Value = val
	return field
}

//UnderlyingCurrentValueField is a AMT field
type UnderlyingCurrentValueField struct{ quickfix.AmtValue }

//Tag returns tag.UnderlyingCurrentValue (885)
func (f UnderlyingCurrentValueField) Tag() quickfix.Tag { return tag.UnderlyingCurrentValue }

//NewUnderlyingCurrentValue returns a new UnderlyingCurrentValueField initialized with val
func NewUnderlyingCurrentValue(val float64) *UnderlyingCurrentValueField {
	field := &UnderlyingCurrentValueField{}
	field.Value = val
	return field
}

//UnderlyingDeliveryAmountField is a AMT field
type UnderlyingDeliveryAmountField struct{ quickfix.AmtValue }

//Tag returns tag.UnderlyingDeliveryAmount (1037)
func (f UnderlyingDeliveryAmountField) Tag() quickfix.Tag { return tag.UnderlyingDeliveryAmount }

//NewUnderlyingDeliveryAmount returns a new UnderlyingDeliveryAmountField initialized with val
func NewUnderlyingDeliveryAmount(val float64) *UnderlyingDeliveryAmountField {
	field := &UnderlyingDeliveryAmountField{}
	field.Value = val
	return field
}

//UnderlyingDetachmentPointField is a PERCENTAGE field
type UnderlyingDetachmentPointField struct{ quickfix.PercentageValue }

//Tag returns tag.UnderlyingDetachmentPoint (1460)
func (f UnderlyingDetachmentPointField) Tag() quickfix.Tag { return tag.UnderlyingDetachmentPoint }

//NewUnderlyingDetachmentPoint returns a new UnderlyingDetachmentPointField initialized with val
func NewUnderlyingDetachmentPoint(val float64) *UnderlyingDetachmentPointField {
	field := &UnderlyingDetachmentPointField{}
	field.Value = val
	return field
}

//UnderlyingDirtyPriceField is a PRICE field
type UnderlyingDirtyPriceField struct{ quickfix.PriceValue }

//Tag returns tag.UnderlyingDirtyPrice (882)
func (f UnderlyingDirtyPriceField) Tag() quickfix.Tag { return tag.UnderlyingDirtyPrice }

//NewUnderlyingDirtyPrice returns a new UnderlyingDirtyPriceField initialized with val
func NewUnderlyingDirtyPrice(val float64) *UnderlyingDirtyPriceField {
	field := &UnderlyingDirtyPriceField{}
	field.Value = val
	return field
}

//UnderlyingEndPriceField is a PRICE field
type UnderlyingEndPriceField struct{ quickfix.PriceValue }

//Tag returns tag.UnderlyingEndPrice (883)
func (f UnderlyingEndPriceField) Tag() quickfix.Tag { return tag.UnderlyingEndPrice }

//NewUnderlyingEndPrice returns a new UnderlyingEndPriceField initialized with val
func NewUnderlyingEndPrice(val float64) *UnderlyingEndPriceField {
	field := &UnderlyingEndPriceField{}
	field.Value = val
	return field
}

//UnderlyingEndValueField is a AMT field
type UnderlyingEndValueField struct{ quickfix.AmtValue }

//Tag returns tag.UnderlyingEndValue (886)
func (f UnderlyingEndValueField) Tag() quickfix.Tag { return tag.UnderlyingEndValue }

//NewUnderlyingEndValue returns a new UnderlyingEndValueField initialized with val
func NewUnderlyingEndValue(val float64) *UnderlyingEndValueField {
	field := &UnderlyingEndValueField{}
	field.Value = val
	return field
}

//UnderlyingExerciseStyleField is a INT field
type UnderlyingExerciseStyleField struct{ quickfix.IntValue }

//Tag returns tag.UnderlyingExerciseStyle (1419)
func (f UnderlyingExerciseStyleField) Tag() quickfix.Tag { return tag.UnderlyingExerciseStyle }

//NewUnderlyingExerciseStyle returns a new UnderlyingExerciseStyleField initialized with val
func NewUnderlyingExerciseStyle(val int) *UnderlyingExerciseStyleField {
	field := &UnderlyingExerciseStyleField{}
	field.Value = val
	return field
}

//UnderlyingFXRateField is a FLOAT field
type UnderlyingFXRateField struct{ quickfix.FloatValue }

//Tag returns tag.UnderlyingFXRate (1045)
func (f UnderlyingFXRateField) Tag() quickfix.Tag { return tag.UnderlyingFXRate }

//NewUnderlyingFXRate returns a new UnderlyingFXRateField initialized with val
func NewUnderlyingFXRate(val float64) *UnderlyingFXRateField {
	field := &UnderlyingFXRateField{}
	field.Value = val
	return field
}

//UnderlyingFXRateCalcField is a CHAR field
type UnderlyingFXRateCalcField struct{ quickfix.CharValue }

//Tag returns tag.UnderlyingFXRateCalc (1046)
func (f UnderlyingFXRateCalcField) Tag() quickfix.Tag { return tag.UnderlyingFXRateCalc }

//NewUnderlyingFXRateCalc returns a new UnderlyingFXRateCalcField initialized with val
func NewUnderlyingFXRateCalc(val string) *UnderlyingFXRateCalcField {
	field := &UnderlyingFXRateCalcField{}
	field.Value = val
	return field
}

//UnderlyingFactorField is a FLOAT field
type UnderlyingFactorField struct{ quickfix.FloatValue }

//Tag returns tag.UnderlyingFactor (246)
func (f UnderlyingFactorField) Tag() quickfix.Tag { return tag.UnderlyingFactor }

//NewUnderlyingFactor returns a new UnderlyingFactorField initialized with val
func NewUnderlyingFactor(val float64) *UnderlyingFactorField {
	field := &UnderlyingFactorField{}
	field.Value = val
	return field
}

//UnderlyingFlowScheduleTypeField is a INT field
type UnderlyingFlowScheduleTypeField struct{ quickfix.IntValue }

//Tag returns tag.UnderlyingFlowScheduleType (1441)
func (f UnderlyingFlowScheduleTypeField) Tag() quickfix.Tag { return tag.UnderlyingFlowScheduleType }

//NewUnderlyingFlowScheduleType returns a new UnderlyingFlowScheduleTypeField initialized with val
func NewUnderlyingFlowScheduleType(val int) *UnderlyingFlowScheduleTypeField {
	field := &UnderlyingFlowScheduleTypeField{}
	field.Value = val
	return field
}

//UnderlyingIDSourceField is a STRING field
type UnderlyingIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingIDSource (305)
func (f UnderlyingIDSourceField) Tag() quickfix.Tag { return tag.UnderlyingIDSource }

//NewUnderlyingIDSource returns a new UnderlyingIDSourceField initialized with val
func NewUnderlyingIDSource(val string) *UnderlyingIDSourceField {
	field := &UnderlyingIDSourceField{}
	field.Value = val
	return field
}

//UnderlyingInstrRegistryField is a STRING field
type UnderlyingInstrRegistryField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingInstrRegistry (595)
func (f UnderlyingInstrRegistryField) Tag() quickfix.Tag { return tag.UnderlyingInstrRegistry }

//NewUnderlyingInstrRegistry returns a new UnderlyingInstrRegistryField initialized with val
func NewUnderlyingInstrRegistry(val string) *UnderlyingInstrRegistryField {
	field := &UnderlyingInstrRegistryField{}
	field.Value = val
	return field
}

//UnderlyingInstrumentPartyIDField is a STRING field
type UnderlyingInstrumentPartyIDField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingInstrumentPartyID (1059)
func (f UnderlyingInstrumentPartyIDField) Tag() quickfix.Tag { return tag.UnderlyingInstrumentPartyID }

//NewUnderlyingInstrumentPartyID returns a new UnderlyingInstrumentPartyIDField initialized with val
func NewUnderlyingInstrumentPartyID(val string) *UnderlyingInstrumentPartyIDField {
	field := &UnderlyingInstrumentPartyIDField{}
	field.Value = val
	return field
}

//UnderlyingInstrumentPartyIDSourceField is a CHAR field
type UnderlyingInstrumentPartyIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.UnderlyingInstrumentPartyIDSource (1060)
func (f UnderlyingInstrumentPartyIDSourceField) Tag() quickfix.Tag {
	return tag.UnderlyingInstrumentPartyIDSource
}

//NewUnderlyingInstrumentPartyIDSource returns a new UnderlyingInstrumentPartyIDSourceField initialized with val
func NewUnderlyingInstrumentPartyIDSource(val string) *UnderlyingInstrumentPartyIDSourceField {
	field := &UnderlyingInstrumentPartyIDSourceField{}
	field.Value = val
	return field
}

//UnderlyingInstrumentPartyRoleField is a INT field
type UnderlyingInstrumentPartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.UnderlyingInstrumentPartyRole (1061)
func (f UnderlyingInstrumentPartyRoleField) Tag() quickfix.Tag {
	return tag.UnderlyingInstrumentPartyRole
}

//NewUnderlyingInstrumentPartyRole returns a new UnderlyingInstrumentPartyRoleField initialized with val
func NewUnderlyingInstrumentPartyRole(val int) *UnderlyingInstrumentPartyRoleField {
	field := &UnderlyingInstrumentPartyRoleField{}
	field.Value = val
	return field
}

//UnderlyingInstrumentPartySubIDField is a STRING field
type UnderlyingInstrumentPartySubIDField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingInstrumentPartySubID (1063)
func (f UnderlyingInstrumentPartySubIDField) Tag() quickfix.Tag {
	return tag.UnderlyingInstrumentPartySubID
}

//NewUnderlyingInstrumentPartySubID returns a new UnderlyingInstrumentPartySubIDField initialized with val
func NewUnderlyingInstrumentPartySubID(val string) *UnderlyingInstrumentPartySubIDField {
	field := &UnderlyingInstrumentPartySubIDField{}
	field.Value = val
	return field
}

//UnderlyingInstrumentPartySubIDTypeField is a INT field
type UnderlyingInstrumentPartySubIDTypeField struct{ quickfix.IntValue }

//Tag returns tag.UnderlyingInstrumentPartySubIDType (1064)
func (f UnderlyingInstrumentPartySubIDTypeField) Tag() quickfix.Tag {
	return tag.UnderlyingInstrumentPartySubIDType
}

//NewUnderlyingInstrumentPartySubIDType returns a new UnderlyingInstrumentPartySubIDTypeField initialized with val
func NewUnderlyingInstrumentPartySubIDType(val int) *UnderlyingInstrumentPartySubIDTypeField {
	field := &UnderlyingInstrumentPartySubIDTypeField{}
	field.Value = val
	return field
}

//UnderlyingIssueDateField is a LOCALMKTDATE field
type UnderlyingIssueDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.UnderlyingIssueDate (242)
func (f UnderlyingIssueDateField) Tag() quickfix.Tag { return tag.UnderlyingIssueDate }

//NewUnderlyingIssueDate returns a new UnderlyingIssueDateField initialized with val
func NewUnderlyingIssueDate(val string) *UnderlyingIssueDateField {
	field := &UnderlyingIssueDateField{}
	field.Value = val
	return field
}

//UnderlyingIssuerField is a STRING field
type UnderlyingIssuerField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingIssuer (306)
func (f UnderlyingIssuerField) Tag() quickfix.Tag { return tag.UnderlyingIssuer }

//NewUnderlyingIssuer returns a new UnderlyingIssuerField initialized with val
func NewUnderlyingIssuer(val string) *UnderlyingIssuerField {
	field := &UnderlyingIssuerField{}
	field.Value = val
	return field
}

//UnderlyingLastPxField is a PRICE field
type UnderlyingLastPxField struct{ quickfix.PriceValue }

//Tag returns tag.UnderlyingLastPx (651)
func (f UnderlyingLastPxField) Tag() quickfix.Tag { return tag.UnderlyingLastPx }

//NewUnderlyingLastPx returns a new UnderlyingLastPxField initialized with val
func NewUnderlyingLastPx(val float64) *UnderlyingLastPxField {
	field := &UnderlyingLastPxField{}
	field.Value = val
	return field
}

//UnderlyingLastQtyField is a QTY field
type UnderlyingLastQtyField struct{ quickfix.QtyValue }

//Tag returns tag.UnderlyingLastQty (652)
func (f UnderlyingLastQtyField) Tag() quickfix.Tag { return tag.UnderlyingLastQty }

//NewUnderlyingLastQty returns a new UnderlyingLastQtyField initialized with val
func NewUnderlyingLastQty(val float64) *UnderlyingLastQtyField {
	field := &UnderlyingLastQtyField{}
	field.Value = val
	return field
}

//UnderlyingLegCFICodeField is a STRING field
type UnderlyingLegCFICodeField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingLegCFICode (1344)
func (f UnderlyingLegCFICodeField) Tag() quickfix.Tag { return tag.UnderlyingLegCFICode }

//NewUnderlyingLegCFICode returns a new UnderlyingLegCFICodeField initialized with val
func NewUnderlyingLegCFICode(val string) *UnderlyingLegCFICodeField {
	field := &UnderlyingLegCFICodeField{}
	field.Value = val
	return field
}

//UnderlyingLegMaturityDateField is a LOCALMKTDATE field
type UnderlyingLegMaturityDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.UnderlyingLegMaturityDate (1345)
func (f UnderlyingLegMaturityDateField) Tag() quickfix.Tag { return tag.UnderlyingLegMaturityDate }

//NewUnderlyingLegMaturityDate returns a new UnderlyingLegMaturityDateField initialized with val
func NewUnderlyingLegMaturityDate(val string) *UnderlyingLegMaturityDateField {
	field := &UnderlyingLegMaturityDateField{}
	field.Value = val
	return field
}

//UnderlyingLegMaturityMonthYearField is a MONTHYEAR field
type UnderlyingLegMaturityMonthYearField struct{ quickfix.MonthYearValue }

//Tag returns tag.UnderlyingLegMaturityMonthYear (1339)
func (f UnderlyingLegMaturityMonthYearField) Tag() quickfix.Tag {
	return tag.UnderlyingLegMaturityMonthYear
}

//NewUnderlyingLegMaturityMonthYear returns a new UnderlyingLegMaturityMonthYearField initialized with val
func NewUnderlyingLegMaturityMonthYear(val string) *UnderlyingLegMaturityMonthYearField {
	field := &UnderlyingLegMaturityMonthYearField{}
	field.Value = val
	return field
}

//UnderlyingLegMaturityTimeField is a TZTIMEONLY field
type UnderlyingLegMaturityTimeField struct{ quickfix.TZTimeOnlyValue }

//Tag returns tag.UnderlyingLegMaturityTime (1405)
func (f UnderlyingLegMaturityTimeField) Tag() quickfix.Tag { return tag.UnderlyingLegMaturityTime }

//UnderlyingLegOptAttributeField is a CHAR field
type UnderlyingLegOptAttributeField struct{ quickfix.CharValue }

//Tag returns tag.UnderlyingLegOptAttribute (1391)
func (f UnderlyingLegOptAttributeField) Tag() quickfix.Tag { return tag.UnderlyingLegOptAttribute }

//NewUnderlyingLegOptAttribute returns a new UnderlyingLegOptAttributeField initialized with val
func NewUnderlyingLegOptAttribute(val string) *UnderlyingLegOptAttributeField {
	field := &UnderlyingLegOptAttributeField{}
	field.Value = val
	return field
}

//UnderlyingLegPutOrCallField is a INT field
type UnderlyingLegPutOrCallField struct{ quickfix.IntValue }

//Tag returns tag.UnderlyingLegPutOrCall (1343)
func (f UnderlyingLegPutOrCallField) Tag() quickfix.Tag { return tag.UnderlyingLegPutOrCall }

//NewUnderlyingLegPutOrCall returns a new UnderlyingLegPutOrCallField initialized with val
func NewUnderlyingLegPutOrCall(val int) *UnderlyingLegPutOrCallField {
	field := &UnderlyingLegPutOrCallField{}
	field.Value = val
	return field
}

//UnderlyingLegSecurityAltIDField is a STRING field
type UnderlyingLegSecurityAltIDField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingLegSecurityAltID (1335)
func (f UnderlyingLegSecurityAltIDField) Tag() quickfix.Tag { return tag.UnderlyingLegSecurityAltID }

//NewUnderlyingLegSecurityAltID returns a new UnderlyingLegSecurityAltIDField initialized with val
func NewUnderlyingLegSecurityAltID(val string) *UnderlyingLegSecurityAltIDField {
	field := &UnderlyingLegSecurityAltIDField{}
	field.Value = val
	return field
}

//UnderlyingLegSecurityAltIDSourceField is a STRING field
type UnderlyingLegSecurityAltIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingLegSecurityAltIDSource (1336)
func (f UnderlyingLegSecurityAltIDSourceField) Tag() quickfix.Tag {
	return tag.UnderlyingLegSecurityAltIDSource
}

//NewUnderlyingLegSecurityAltIDSource returns a new UnderlyingLegSecurityAltIDSourceField initialized with val
func NewUnderlyingLegSecurityAltIDSource(val string) *UnderlyingLegSecurityAltIDSourceField {
	field := &UnderlyingLegSecurityAltIDSourceField{}
	field.Value = val
	return field
}

//UnderlyingLegSecurityDescField is a STRING field
type UnderlyingLegSecurityDescField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingLegSecurityDesc (1392)
func (f UnderlyingLegSecurityDescField) Tag() quickfix.Tag { return tag.UnderlyingLegSecurityDesc }

//NewUnderlyingLegSecurityDesc returns a new UnderlyingLegSecurityDescField initialized with val
func NewUnderlyingLegSecurityDesc(val string) *UnderlyingLegSecurityDescField {
	field := &UnderlyingLegSecurityDescField{}
	field.Value = val
	return field
}

//UnderlyingLegSecurityExchangeField is a STRING field
type UnderlyingLegSecurityExchangeField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingLegSecurityExchange (1341)
func (f UnderlyingLegSecurityExchangeField) Tag() quickfix.Tag {
	return tag.UnderlyingLegSecurityExchange
}

//NewUnderlyingLegSecurityExchange returns a new UnderlyingLegSecurityExchangeField initialized with val
func NewUnderlyingLegSecurityExchange(val string) *UnderlyingLegSecurityExchangeField {
	field := &UnderlyingLegSecurityExchangeField{}
	field.Value = val
	return field
}

//UnderlyingLegSecurityIDField is a STRING field
type UnderlyingLegSecurityIDField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingLegSecurityID (1332)
func (f UnderlyingLegSecurityIDField) Tag() quickfix.Tag { return tag.UnderlyingLegSecurityID }

//NewUnderlyingLegSecurityID returns a new UnderlyingLegSecurityIDField initialized with val
func NewUnderlyingLegSecurityID(val string) *UnderlyingLegSecurityIDField {
	field := &UnderlyingLegSecurityIDField{}
	field.Value = val
	return field
}

//UnderlyingLegSecurityIDSourceField is a STRING field
type UnderlyingLegSecurityIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingLegSecurityIDSource (1333)
func (f UnderlyingLegSecurityIDSourceField) Tag() quickfix.Tag {
	return tag.UnderlyingLegSecurityIDSource
}

//NewUnderlyingLegSecurityIDSource returns a new UnderlyingLegSecurityIDSourceField initialized with val
func NewUnderlyingLegSecurityIDSource(val string) *UnderlyingLegSecurityIDSourceField {
	field := &UnderlyingLegSecurityIDSourceField{}
	field.Value = val
	return field
}

//UnderlyingLegSecuritySubTypeField is a STRING field
type UnderlyingLegSecuritySubTypeField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingLegSecuritySubType (1338)
func (f UnderlyingLegSecuritySubTypeField) Tag() quickfix.Tag { return tag.UnderlyingLegSecuritySubType }

//NewUnderlyingLegSecuritySubType returns a new UnderlyingLegSecuritySubTypeField initialized with val
func NewUnderlyingLegSecuritySubType(val string) *UnderlyingLegSecuritySubTypeField {
	field := &UnderlyingLegSecuritySubTypeField{}
	field.Value = val
	return field
}

//UnderlyingLegSecurityTypeField is a STRING field
type UnderlyingLegSecurityTypeField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingLegSecurityType (1337)
func (f UnderlyingLegSecurityTypeField) Tag() quickfix.Tag { return tag.UnderlyingLegSecurityType }

//NewUnderlyingLegSecurityType returns a new UnderlyingLegSecurityTypeField initialized with val
func NewUnderlyingLegSecurityType(val string) *UnderlyingLegSecurityTypeField {
	field := &UnderlyingLegSecurityTypeField{}
	field.Value = val
	return field
}

//UnderlyingLegStrikePriceField is a PRICE field
type UnderlyingLegStrikePriceField struct{ quickfix.PriceValue }

//Tag returns tag.UnderlyingLegStrikePrice (1340)
func (f UnderlyingLegStrikePriceField) Tag() quickfix.Tag { return tag.UnderlyingLegStrikePrice }

//NewUnderlyingLegStrikePrice returns a new UnderlyingLegStrikePriceField initialized with val
func NewUnderlyingLegStrikePrice(val float64) *UnderlyingLegStrikePriceField {
	field := &UnderlyingLegStrikePriceField{}
	field.Value = val
	return field
}

//UnderlyingLegSymbolField is a STRING field
type UnderlyingLegSymbolField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingLegSymbol (1330)
func (f UnderlyingLegSymbolField) Tag() quickfix.Tag { return tag.UnderlyingLegSymbol }

//NewUnderlyingLegSymbol returns a new UnderlyingLegSymbolField initialized with val
func NewUnderlyingLegSymbol(val string) *UnderlyingLegSymbolField {
	field := &UnderlyingLegSymbolField{}
	field.Value = val
	return field
}

//UnderlyingLegSymbolSfxField is a STRING field
type UnderlyingLegSymbolSfxField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingLegSymbolSfx (1331)
func (f UnderlyingLegSymbolSfxField) Tag() quickfix.Tag { return tag.UnderlyingLegSymbolSfx }

//NewUnderlyingLegSymbolSfx returns a new UnderlyingLegSymbolSfxField initialized with val
func NewUnderlyingLegSymbolSfx(val string) *UnderlyingLegSymbolSfxField {
	field := &UnderlyingLegSymbolSfxField{}
	field.Value = val
	return field
}

//UnderlyingLocaleOfIssueField is a STRING field
type UnderlyingLocaleOfIssueField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingLocaleOfIssue (594)
func (f UnderlyingLocaleOfIssueField) Tag() quickfix.Tag { return tag.UnderlyingLocaleOfIssue }

//NewUnderlyingLocaleOfIssue returns a new UnderlyingLocaleOfIssueField initialized with val
func NewUnderlyingLocaleOfIssue(val string) *UnderlyingLocaleOfIssueField {
	field := &UnderlyingLocaleOfIssueField{}
	field.Value = val
	return field
}

//UnderlyingMaturityDateField is a LOCALMKTDATE field
type UnderlyingMaturityDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.UnderlyingMaturityDate (542)
func (f UnderlyingMaturityDateField) Tag() quickfix.Tag { return tag.UnderlyingMaturityDate }

//NewUnderlyingMaturityDate returns a new UnderlyingMaturityDateField initialized with val
func NewUnderlyingMaturityDate(val string) *UnderlyingMaturityDateField {
	field := &UnderlyingMaturityDateField{}
	field.Value = val
	return field
}

//UnderlyingMaturityDayField is a DAYOFMONTH field
type UnderlyingMaturityDayField struct{ quickfix.DayOfMonthValue }

//Tag returns tag.UnderlyingMaturityDay (314)
func (f UnderlyingMaturityDayField) Tag() quickfix.Tag { return tag.UnderlyingMaturityDay }

//NewUnderlyingMaturityDay returns a new UnderlyingMaturityDayField initialized with val
func NewUnderlyingMaturityDay(val int) *UnderlyingMaturityDayField {
	field := &UnderlyingMaturityDayField{}
	field.Value = val
	return field
}

//UnderlyingMaturityMonthYearField is a MONTHYEAR field
type UnderlyingMaturityMonthYearField struct{ quickfix.MonthYearValue }

//Tag returns tag.UnderlyingMaturityMonthYear (313)
func (f UnderlyingMaturityMonthYearField) Tag() quickfix.Tag { return tag.UnderlyingMaturityMonthYear }

//NewUnderlyingMaturityMonthYear returns a new UnderlyingMaturityMonthYearField initialized with val
func NewUnderlyingMaturityMonthYear(val string) *UnderlyingMaturityMonthYearField {
	field := &UnderlyingMaturityMonthYearField{}
	field.Value = val
	return field
}

//UnderlyingMaturityTimeField is a TZTIMEONLY field
type UnderlyingMaturityTimeField struct{ quickfix.TZTimeOnlyValue }

//Tag returns tag.UnderlyingMaturityTime (1213)
func (f UnderlyingMaturityTimeField) Tag() quickfix.Tag { return tag.UnderlyingMaturityTime }

//UnderlyingNotionalPercentageOutstandingField is a PERCENTAGE field
type UnderlyingNotionalPercentageOutstandingField struct{ quickfix.PercentageValue }

//Tag returns tag.UnderlyingNotionalPercentageOutstanding (1455)
func (f UnderlyingNotionalPercentageOutstandingField) Tag() quickfix.Tag {
	return tag.UnderlyingNotionalPercentageOutstanding
}

//NewUnderlyingNotionalPercentageOutstanding returns a new UnderlyingNotionalPercentageOutstandingField initialized with val
func NewUnderlyingNotionalPercentageOutstanding(val float64) *UnderlyingNotionalPercentageOutstandingField {
	field := &UnderlyingNotionalPercentageOutstandingField{}
	field.Value = val
	return field
}

//UnderlyingOptAttributeField is a CHAR field
type UnderlyingOptAttributeField struct{ quickfix.CharValue }

//Tag returns tag.UnderlyingOptAttribute (317)
func (f UnderlyingOptAttributeField) Tag() quickfix.Tag { return tag.UnderlyingOptAttribute }

//NewUnderlyingOptAttribute returns a new UnderlyingOptAttributeField initialized with val
func NewUnderlyingOptAttribute(val string) *UnderlyingOptAttributeField {
	field := &UnderlyingOptAttributeField{}
	field.Value = val
	return field
}

//UnderlyingOriginalNotionalPercentageOutstandingField is a PERCENTAGE field
type UnderlyingOriginalNotionalPercentageOutstandingField struct{ quickfix.PercentageValue }

//Tag returns tag.UnderlyingOriginalNotionalPercentageOutstanding (1456)
func (f UnderlyingOriginalNotionalPercentageOutstandingField) Tag() quickfix.Tag {
	return tag.UnderlyingOriginalNotionalPercentageOutstanding
}

//NewUnderlyingOriginalNotionalPercentageOutstanding returns a new UnderlyingOriginalNotionalPercentageOutstandingField initialized with val
func NewUnderlyingOriginalNotionalPercentageOutstanding(val float64) *UnderlyingOriginalNotionalPercentageOutstandingField {
	field := &UnderlyingOriginalNotionalPercentageOutstandingField{}
	field.Value = val
	return field
}

//UnderlyingPayAmountField is a AMT field
type UnderlyingPayAmountField struct{ quickfix.AmtValue }

//Tag returns tag.UnderlyingPayAmount (985)
func (f UnderlyingPayAmountField) Tag() quickfix.Tag { return tag.UnderlyingPayAmount }

//NewUnderlyingPayAmount returns a new UnderlyingPayAmountField initialized with val
func NewUnderlyingPayAmount(val float64) *UnderlyingPayAmountField {
	field := &UnderlyingPayAmountField{}
	field.Value = val
	return field
}

//UnderlyingPriceDeterminationMethodField is a INT field
type UnderlyingPriceDeterminationMethodField struct{ quickfix.IntValue }

//Tag returns tag.UnderlyingPriceDeterminationMethod (1481)
func (f UnderlyingPriceDeterminationMethodField) Tag() quickfix.Tag {
	return tag.UnderlyingPriceDeterminationMethod
}

//NewUnderlyingPriceDeterminationMethod returns a new UnderlyingPriceDeterminationMethodField initialized with val
func NewUnderlyingPriceDeterminationMethod(val int) *UnderlyingPriceDeterminationMethodField {
	field := &UnderlyingPriceDeterminationMethodField{}
	field.Value = val
	return field
}

//UnderlyingPriceUnitOfMeasureField is a STRING field
type UnderlyingPriceUnitOfMeasureField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingPriceUnitOfMeasure (1424)
func (f UnderlyingPriceUnitOfMeasureField) Tag() quickfix.Tag { return tag.UnderlyingPriceUnitOfMeasure }

//NewUnderlyingPriceUnitOfMeasure returns a new UnderlyingPriceUnitOfMeasureField initialized with val
func NewUnderlyingPriceUnitOfMeasure(val string) *UnderlyingPriceUnitOfMeasureField {
	field := &UnderlyingPriceUnitOfMeasureField{}
	field.Value = val
	return field
}

//UnderlyingPriceUnitOfMeasureQtyField is a QTY field
type UnderlyingPriceUnitOfMeasureQtyField struct{ quickfix.QtyValue }

//Tag returns tag.UnderlyingPriceUnitOfMeasureQty (1425)
func (f UnderlyingPriceUnitOfMeasureQtyField) Tag() quickfix.Tag {
	return tag.UnderlyingPriceUnitOfMeasureQty
}

//NewUnderlyingPriceUnitOfMeasureQty returns a new UnderlyingPriceUnitOfMeasureQtyField initialized with val
func NewUnderlyingPriceUnitOfMeasureQty(val float64) *UnderlyingPriceUnitOfMeasureQtyField {
	field := &UnderlyingPriceUnitOfMeasureQtyField{}
	field.Value = val
	return field
}

//UnderlyingProductField is a INT field
type UnderlyingProductField struct{ quickfix.IntValue }

//Tag returns tag.UnderlyingProduct (462)
func (f UnderlyingProductField) Tag() quickfix.Tag { return tag.UnderlyingProduct }

//NewUnderlyingProduct returns a new UnderlyingProductField initialized with val
func NewUnderlyingProduct(val int) *UnderlyingProductField {
	field := &UnderlyingProductField{}
	field.Value = val
	return field
}

//UnderlyingPutOrCallField is a INT field
type UnderlyingPutOrCallField struct{ quickfix.IntValue }

//Tag returns tag.UnderlyingPutOrCall (315)
func (f UnderlyingPutOrCallField) Tag() quickfix.Tag { return tag.UnderlyingPutOrCall }

//NewUnderlyingPutOrCall returns a new UnderlyingPutOrCallField initialized with val
func NewUnderlyingPutOrCall(val int) *UnderlyingPutOrCallField {
	field := &UnderlyingPutOrCallField{}
	field.Value = val
	return field
}

//UnderlyingPxField is a PRICE field
type UnderlyingPxField struct{ quickfix.PriceValue }

//Tag returns tag.UnderlyingPx (810)
func (f UnderlyingPxField) Tag() quickfix.Tag { return tag.UnderlyingPx }

//NewUnderlyingPx returns a new UnderlyingPxField initialized with val
func NewUnderlyingPx(val float64) *UnderlyingPxField {
	field := &UnderlyingPxField{}
	field.Value = val
	return field
}

//UnderlyingQtyField is a QTY field
type UnderlyingQtyField struct{ quickfix.QtyValue }

//Tag returns tag.UnderlyingQty (879)
func (f UnderlyingQtyField) Tag() quickfix.Tag { return tag.UnderlyingQty }

//NewUnderlyingQty returns a new UnderlyingQtyField initialized with val
func NewUnderlyingQty(val float64) *UnderlyingQtyField {
	field := &UnderlyingQtyField{}
	field.Value = val
	return field
}

//UnderlyingRedemptionDateField is a LOCALMKTDATE field
type UnderlyingRedemptionDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.UnderlyingRedemptionDate (247)
func (f UnderlyingRedemptionDateField) Tag() quickfix.Tag { return tag.UnderlyingRedemptionDate }

//NewUnderlyingRedemptionDate returns a new UnderlyingRedemptionDateField initialized with val
func NewUnderlyingRedemptionDate(val string) *UnderlyingRedemptionDateField {
	field := &UnderlyingRedemptionDateField{}
	field.Value = val
	return field
}

//UnderlyingRepoCollateralSecurityTypeField is a INT field
type UnderlyingRepoCollateralSecurityTypeField struct{ quickfix.IntValue }

//Tag returns tag.UnderlyingRepoCollateralSecurityType (243)
func (f UnderlyingRepoCollateralSecurityTypeField) Tag() quickfix.Tag {
	return tag.UnderlyingRepoCollateralSecurityType
}

//NewUnderlyingRepoCollateralSecurityType returns a new UnderlyingRepoCollateralSecurityTypeField initialized with val
func NewUnderlyingRepoCollateralSecurityType(val int) *UnderlyingRepoCollateralSecurityTypeField {
	field := &UnderlyingRepoCollateralSecurityTypeField{}
	field.Value = val
	return field
}

//UnderlyingRepurchaseRateField is a PERCENTAGE field
type UnderlyingRepurchaseRateField struct{ quickfix.PercentageValue }

//Tag returns tag.UnderlyingRepurchaseRate (245)
func (f UnderlyingRepurchaseRateField) Tag() quickfix.Tag { return tag.UnderlyingRepurchaseRate }

//NewUnderlyingRepurchaseRate returns a new UnderlyingRepurchaseRateField initialized with val
func NewUnderlyingRepurchaseRate(val float64) *UnderlyingRepurchaseRateField {
	field := &UnderlyingRepurchaseRateField{}
	field.Value = val
	return field
}

//UnderlyingRepurchaseTermField is a INT field
type UnderlyingRepurchaseTermField struct{ quickfix.IntValue }

//Tag returns tag.UnderlyingRepurchaseTerm (244)
func (f UnderlyingRepurchaseTermField) Tag() quickfix.Tag { return tag.UnderlyingRepurchaseTerm }

//NewUnderlyingRepurchaseTerm returns a new UnderlyingRepurchaseTermField initialized with val
func NewUnderlyingRepurchaseTerm(val int) *UnderlyingRepurchaseTermField {
	field := &UnderlyingRepurchaseTermField{}
	field.Value = val
	return field
}

//UnderlyingRestructuringTypeField is a STRING field
type UnderlyingRestructuringTypeField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingRestructuringType (1453)
func (f UnderlyingRestructuringTypeField) Tag() quickfix.Tag { return tag.UnderlyingRestructuringType }

//NewUnderlyingRestructuringType returns a new UnderlyingRestructuringTypeField initialized with val
func NewUnderlyingRestructuringType(val string) *UnderlyingRestructuringTypeField {
	field := &UnderlyingRestructuringTypeField{}
	field.Value = val
	return field
}

//UnderlyingSecurityAltIDField is a STRING field
type UnderlyingSecurityAltIDField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingSecurityAltID (458)
func (f UnderlyingSecurityAltIDField) Tag() quickfix.Tag { return tag.UnderlyingSecurityAltID }

//NewUnderlyingSecurityAltID returns a new UnderlyingSecurityAltIDField initialized with val
func NewUnderlyingSecurityAltID(val string) *UnderlyingSecurityAltIDField {
	field := &UnderlyingSecurityAltIDField{}
	field.Value = val
	return field
}

//UnderlyingSecurityAltIDSourceField is a STRING field
type UnderlyingSecurityAltIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingSecurityAltIDSource (459)
func (f UnderlyingSecurityAltIDSourceField) Tag() quickfix.Tag {
	return tag.UnderlyingSecurityAltIDSource
}

//NewUnderlyingSecurityAltIDSource returns a new UnderlyingSecurityAltIDSourceField initialized with val
func NewUnderlyingSecurityAltIDSource(val string) *UnderlyingSecurityAltIDSourceField {
	field := &UnderlyingSecurityAltIDSourceField{}
	field.Value = val
	return field
}

//UnderlyingSecurityDescField is a STRING field
type UnderlyingSecurityDescField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingSecurityDesc (307)
func (f UnderlyingSecurityDescField) Tag() quickfix.Tag { return tag.UnderlyingSecurityDesc }

//NewUnderlyingSecurityDesc returns a new UnderlyingSecurityDescField initialized with val
func NewUnderlyingSecurityDesc(val string) *UnderlyingSecurityDescField {
	field := &UnderlyingSecurityDescField{}
	field.Value = val
	return field
}

//UnderlyingSecurityExchangeField is a EXCHANGE field
type UnderlyingSecurityExchangeField struct{ quickfix.ExchangeValue }

//Tag returns tag.UnderlyingSecurityExchange (308)
func (f UnderlyingSecurityExchangeField) Tag() quickfix.Tag { return tag.UnderlyingSecurityExchange }

//NewUnderlyingSecurityExchange returns a new UnderlyingSecurityExchangeField initialized with val
func NewUnderlyingSecurityExchange(val string) *UnderlyingSecurityExchangeField {
	field := &UnderlyingSecurityExchangeField{}
	field.Value = val
	return field
}

//UnderlyingSecurityIDField is a STRING field
type UnderlyingSecurityIDField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingSecurityID (309)
func (f UnderlyingSecurityIDField) Tag() quickfix.Tag { return tag.UnderlyingSecurityID }

//NewUnderlyingSecurityID returns a new UnderlyingSecurityIDField initialized with val
func NewUnderlyingSecurityID(val string) *UnderlyingSecurityIDField {
	field := &UnderlyingSecurityIDField{}
	field.Value = val
	return field
}

//UnderlyingSecurityIDSourceField is a STRING field
type UnderlyingSecurityIDSourceField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingSecurityIDSource (305)
func (f UnderlyingSecurityIDSourceField) Tag() quickfix.Tag { return tag.UnderlyingSecurityIDSource }

//NewUnderlyingSecurityIDSource returns a new UnderlyingSecurityIDSourceField initialized with val
func NewUnderlyingSecurityIDSource(val string) *UnderlyingSecurityIDSourceField {
	field := &UnderlyingSecurityIDSourceField{}
	field.Value = val
	return field
}

//UnderlyingSecuritySubTypeField is a STRING field
type UnderlyingSecuritySubTypeField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingSecuritySubType (763)
func (f UnderlyingSecuritySubTypeField) Tag() quickfix.Tag { return tag.UnderlyingSecuritySubType }

//NewUnderlyingSecuritySubType returns a new UnderlyingSecuritySubTypeField initialized with val
func NewUnderlyingSecuritySubType(val string) *UnderlyingSecuritySubTypeField {
	field := &UnderlyingSecuritySubTypeField{}
	field.Value = val
	return field
}

//UnderlyingSecurityTypeField is a STRING field
type UnderlyingSecurityTypeField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingSecurityType (310)
func (f UnderlyingSecurityTypeField) Tag() quickfix.Tag { return tag.UnderlyingSecurityType }

//NewUnderlyingSecurityType returns a new UnderlyingSecurityTypeField initialized with val
func NewUnderlyingSecurityType(val string) *UnderlyingSecurityTypeField {
	field := &UnderlyingSecurityTypeField{}
	field.Value = val
	return field
}

//UnderlyingSeniorityField is a STRING field
type UnderlyingSeniorityField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingSeniority (1454)
func (f UnderlyingSeniorityField) Tag() quickfix.Tag { return tag.UnderlyingSeniority }

//NewUnderlyingSeniority returns a new UnderlyingSeniorityField initialized with val
func NewUnderlyingSeniority(val string) *UnderlyingSeniorityField {
	field := &UnderlyingSeniorityField{}
	field.Value = val
	return field
}

//UnderlyingSettlMethodField is a STRING field
type UnderlyingSettlMethodField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingSettlMethod (1039)
func (f UnderlyingSettlMethodField) Tag() quickfix.Tag { return tag.UnderlyingSettlMethod }

//NewUnderlyingSettlMethod returns a new UnderlyingSettlMethodField initialized with val
func NewUnderlyingSettlMethod(val string) *UnderlyingSettlMethodField {
	field := &UnderlyingSettlMethodField{}
	field.Value = val
	return field
}

//UnderlyingSettlPriceField is a PRICE field
type UnderlyingSettlPriceField struct{ quickfix.PriceValue }

//Tag returns tag.UnderlyingSettlPrice (732)
func (f UnderlyingSettlPriceField) Tag() quickfix.Tag { return tag.UnderlyingSettlPrice }

//NewUnderlyingSettlPrice returns a new UnderlyingSettlPriceField initialized with val
func NewUnderlyingSettlPrice(val float64) *UnderlyingSettlPriceField {
	field := &UnderlyingSettlPriceField{}
	field.Value = val
	return field
}

//UnderlyingSettlPriceTypeField is a INT field
type UnderlyingSettlPriceTypeField struct{ quickfix.IntValue }

//Tag returns tag.UnderlyingSettlPriceType (733)
func (f UnderlyingSettlPriceTypeField) Tag() quickfix.Tag { return tag.UnderlyingSettlPriceType }

//NewUnderlyingSettlPriceType returns a new UnderlyingSettlPriceTypeField initialized with val
func NewUnderlyingSettlPriceType(val int) *UnderlyingSettlPriceTypeField {
	field := &UnderlyingSettlPriceTypeField{}
	field.Value = val
	return field
}

//UnderlyingSettlementDateField is a LOCALMKTDATE field
type UnderlyingSettlementDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.UnderlyingSettlementDate (987)
func (f UnderlyingSettlementDateField) Tag() quickfix.Tag { return tag.UnderlyingSettlementDate }

//NewUnderlyingSettlementDate returns a new UnderlyingSettlementDateField initialized with val
func NewUnderlyingSettlementDate(val string) *UnderlyingSettlementDateField {
	field := &UnderlyingSettlementDateField{}
	field.Value = val
	return field
}

//UnderlyingSettlementStatusField is a STRING field
type UnderlyingSettlementStatusField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingSettlementStatus (988)
func (f UnderlyingSettlementStatusField) Tag() quickfix.Tag { return tag.UnderlyingSettlementStatus }

//NewUnderlyingSettlementStatus returns a new UnderlyingSettlementStatusField initialized with val
func NewUnderlyingSettlementStatus(val string) *UnderlyingSettlementStatusField {
	field := &UnderlyingSettlementStatusField{}
	field.Value = val
	return field
}

//UnderlyingSettlementTypeField is a INT field
type UnderlyingSettlementTypeField struct{ quickfix.IntValue }

//Tag returns tag.UnderlyingSettlementType (975)
func (f UnderlyingSettlementTypeField) Tag() quickfix.Tag { return tag.UnderlyingSettlementType }

//NewUnderlyingSettlementType returns a new UnderlyingSettlementTypeField initialized with val
func NewUnderlyingSettlementType(val int) *UnderlyingSettlementTypeField {
	field := &UnderlyingSettlementTypeField{}
	field.Value = val
	return field
}

//UnderlyingStartValueField is a AMT field
type UnderlyingStartValueField struct{ quickfix.AmtValue }

//Tag returns tag.UnderlyingStartValue (884)
func (f UnderlyingStartValueField) Tag() quickfix.Tag { return tag.UnderlyingStartValue }

//NewUnderlyingStartValue returns a new UnderlyingStartValueField initialized with val
func NewUnderlyingStartValue(val float64) *UnderlyingStartValueField {
	field := &UnderlyingStartValueField{}
	field.Value = val
	return field
}

//UnderlyingStateOrProvinceOfIssueField is a STRING field
type UnderlyingStateOrProvinceOfIssueField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingStateOrProvinceOfIssue (593)
func (f UnderlyingStateOrProvinceOfIssueField) Tag() quickfix.Tag {
	return tag.UnderlyingStateOrProvinceOfIssue
}

//NewUnderlyingStateOrProvinceOfIssue returns a new UnderlyingStateOrProvinceOfIssueField initialized with val
func NewUnderlyingStateOrProvinceOfIssue(val string) *UnderlyingStateOrProvinceOfIssueField {
	field := &UnderlyingStateOrProvinceOfIssueField{}
	field.Value = val
	return field
}

//UnderlyingStipTypeField is a STRING field
type UnderlyingStipTypeField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingStipType (888)
func (f UnderlyingStipTypeField) Tag() quickfix.Tag { return tag.UnderlyingStipType }

//NewUnderlyingStipType returns a new UnderlyingStipTypeField initialized with val
func NewUnderlyingStipType(val string) *UnderlyingStipTypeField {
	field := &UnderlyingStipTypeField{}
	field.Value = val
	return field
}

//UnderlyingStipValueField is a STRING field
type UnderlyingStipValueField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingStipValue (889)
func (f UnderlyingStipValueField) Tag() quickfix.Tag { return tag.UnderlyingStipValue }

//NewUnderlyingStipValue returns a new UnderlyingStipValueField initialized with val
func NewUnderlyingStipValue(val string) *UnderlyingStipValueField {
	field := &UnderlyingStipValueField{}
	field.Value = val
	return field
}

//UnderlyingStrikeCurrencyField is a CURRENCY field
type UnderlyingStrikeCurrencyField struct{ quickfix.CurrencyValue }

//Tag returns tag.UnderlyingStrikeCurrency (941)
func (f UnderlyingStrikeCurrencyField) Tag() quickfix.Tag { return tag.UnderlyingStrikeCurrency }

//NewUnderlyingStrikeCurrency returns a new UnderlyingStrikeCurrencyField initialized with val
func NewUnderlyingStrikeCurrency(val string) *UnderlyingStrikeCurrencyField {
	field := &UnderlyingStrikeCurrencyField{}
	field.Value = val
	return field
}

//UnderlyingStrikePriceField is a PRICE field
type UnderlyingStrikePriceField struct{ quickfix.PriceValue }

//Tag returns tag.UnderlyingStrikePrice (316)
func (f UnderlyingStrikePriceField) Tag() quickfix.Tag { return tag.UnderlyingStrikePrice }

//NewUnderlyingStrikePrice returns a new UnderlyingStrikePriceField initialized with val
func NewUnderlyingStrikePrice(val float64) *UnderlyingStrikePriceField {
	field := &UnderlyingStrikePriceField{}
	field.Value = val
	return field
}

//UnderlyingSymbolField is a STRING field
type UnderlyingSymbolField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingSymbol (311)
func (f UnderlyingSymbolField) Tag() quickfix.Tag { return tag.UnderlyingSymbol }

//NewUnderlyingSymbol returns a new UnderlyingSymbolField initialized with val
func NewUnderlyingSymbol(val string) *UnderlyingSymbolField {
	field := &UnderlyingSymbolField{}
	field.Value = val
	return field
}

//UnderlyingSymbolSfxField is a STRING field
type UnderlyingSymbolSfxField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingSymbolSfx (312)
func (f UnderlyingSymbolSfxField) Tag() quickfix.Tag { return tag.UnderlyingSymbolSfx }

//NewUnderlyingSymbolSfx returns a new UnderlyingSymbolSfxField initialized with val
func NewUnderlyingSymbolSfx(val string) *UnderlyingSymbolSfxField {
	field := &UnderlyingSymbolSfxField{}
	field.Value = val
	return field
}

//UnderlyingTimeUnitField is a STRING field
type UnderlyingTimeUnitField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingTimeUnit (1000)
func (f UnderlyingTimeUnitField) Tag() quickfix.Tag { return tag.UnderlyingTimeUnit }

//NewUnderlyingTimeUnit returns a new UnderlyingTimeUnitField initialized with val
func NewUnderlyingTimeUnit(val string) *UnderlyingTimeUnitField {
	field := &UnderlyingTimeUnitField{}
	field.Value = val
	return field
}

//UnderlyingTradingSessionIDField is a STRING field
type UnderlyingTradingSessionIDField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingTradingSessionID (822)
func (f UnderlyingTradingSessionIDField) Tag() quickfix.Tag { return tag.UnderlyingTradingSessionID }

//NewUnderlyingTradingSessionID returns a new UnderlyingTradingSessionIDField initialized with val
func NewUnderlyingTradingSessionID(val string) *UnderlyingTradingSessionIDField {
	field := &UnderlyingTradingSessionIDField{}
	field.Value = val
	return field
}

//UnderlyingTradingSessionSubIDField is a STRING field
type UnderlyingTradingSessionSubIDField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingTradingSessionSubID (823)
func (f UnderlyingTradingSessionSubIDField) Tag() quickfix.Tag {
	return tag.UnderlyingTradingSessionSubID
}

//NewUnderlyingTradingSessionSubID returns a new UnderlyingTradingSessionSubIDField initialized with val
func NewUnderlyingTradingSessionSubID(val string) *UnderlyingTradingSessionSubIDField {
	field := &UnderlyingTradingSessionSubIDField{}
	field.Value = val
	return field
}

//UnderlyingUnitOfMeasureField is a STRING field
type UnderlyingUnitOfMeasureField struct{ quickfix.StringValue }

//Tag returns tag.UnderlyingUnitOfMeasure (998)
func (f UnderlyingUnitOfMeasureField) Tag() quickfix.Tag { return tag.UnderlyingUnitOfMeasure }

//NewUnderlyingUnitOfMeasure returns a new UnderlyingUnitOfMeasureField initialized with val
func NewUnderlyingUnitOfMeasure(val string) *UnderlyingUnitOfMeasureField {
	field := &UnderlyingUnitOfMeasureField{}
	field.Value = val
	return field
}

//UnderlyingUnitOfMeasureQtyField is a QTY field
type UnderlyingUnitOfMeasureQtyField struct{ quickfix.QtyValue }

//Tag returns tag.UnderlyingUnitOfMeasureQty (1423)
func (f UnderlyingUnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.UnderlyingUnitOfMeasureQty }

//NewUnderlyingUnitOfMeasureQty returns a new UnderlyingUnitOfMeasureQtyField initialized with val
func NewUnderlyingUnitOfMeasureQty(val float64) *UnderlyingUnitOfMeasureQtyField {
	field := &UnderlyingUnitOfMeasureQtyField{}
	field.Value = val
	return field
}

//UndlyInstrumentPartyIDField is a STRING field
type UndlyInstrumentPartyIDField struct{ quickfix.StringValue }

//Tag returns tag.UndlyInstrumentPartyID (1059)
func (f UndlyInstrumentPartyIDField) Tag() quickfix.Tag { return tag.UndlyInstrumentPartyID }

//NewUndlyInstrumentPartyID returns a new UndlyInstrumentPartyIDField initialized with val
func NewUndlyInstrumentPartyID(val string) *UndlyInstrumentPartyIDField {
	field := &UndlyInstrumentPartyIDField{}
	field.Value = val
	return field
}

//UndlyInstrumentPartyIDSourceField is a CHAR field
type UndlyInstrumentPartyIDSourceField struct{ quickfix.CharValue }

//Tag returns tag.UndlyInstrumentPartyIDSource (1060)
func (f UndlyInstrumentPartyIDSourceField) Tag() quickfix.Tag { return tag.UndlyInstrumentPartyIDSource }

//NewUndlyInstrumentPartyIDSource returns a new UndlyInstrumentPartyIDSourceField initialized with val
func NewUndlyInstrumentPartyIDSource(val string) *UndlyInstrumentPartyIDSourceField {
	field := &UndlyInstrumentPartyIDSourceField{}
	field.Value = val
	return field
}

//UndlyInstrumentPartyRoleField is a INT field
type UndlyInstrumentPartyRoleField struct{ quickfix.IntValue }

//Tag returns tag.UndlyInstrumentPartyRole (1061)
func (f UndlyInstrumentPartyRoleField) Tag() quickfix.Tag { return tag.UndlyInstrumentPartyRole }

//NewUndlyInstrumentPartyRole returns a new UndlyInstrumentPartyRoleField initialized with val
func NewUndlyInstrumentPartyRole(val int) *UndlyInstrumentPartyRoleField {
	field := &UndlyInstrumentPartyRoleField{}
	field.Value = val
	return field
}

//UndlyInstrumentPartySubIDField is a STRING field
type UndlyInstrumentPartySubIDField struct{ quickfix.StringValue }

//Tag returns tag.UndlyInstrumentPartySubID (1063)
func (f UndlyInstrumentPartySubIDField) Tag() quickfix.Tag { return tag.UndlyInstrumentPartySubID }

//NewUndlyInstrumentPartySubID returns a new UndlyInstrumentPartySubIDField initialized with val
func NewUndlyInstrumentPartySubID(val string) *UndlyInstrumentPartySubIDField {
	field := &UndlyInstrumentPartySubIDField{}
	field.Value = val
	return field
}

//UndlyInstrumentPartySubIDTypeField is a INT field
type UndlyInstrumentPartySubIDTypeField struct{ quickfix.IntValue }

//Tag returns tag.UndlyInstrumentPartySubIDType (1064)
func (f UndlyInstrumentPartySubIDTypeField) Tag() quickfix.Tag {
	return tag.UndlyInstrumentPartySubIDType
}

//NewUndlyInstrumentPartySubIDType returns a new UndlyInstrumentPartySubIDTypeField initialized with val
func NewUndlyInstrumentPartySubIDType(val int) *UndlyInstrumentPartySubIDTypeField {
	field := &UndlyInstrumentPartySubIDTypeField{}
	field.Value = val
	return field
}

//UnitOfMeasureField is a STRING field
type UnitOfMeasureField struct{ quickfix.StringValue }

//Tag returns tag.UnitOfMeasure (996)
func (f UnitOfMeasureField) Tag() quickfix.Tag { return tag.UnitOfMeasure }

//NewUnitOfMeasure returns a new UnitOfMeasureField initialized with val
func NewUnitOfMeasure(val string) *UnitOfMeasureField {
	field := &UnitOfMeasureField{}
	field.Value = val
	return field
}

//UnitOfMeasureQtyField is a QTY field
type UnitOfMeasureQtyField struct{ quickfix.QtyValue }

//Tag returns tag.UnitOfMeasureQty (1147)
func (f UnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.UnitOfMeasureQty }

//NewUnitOfMeasureQty returns a new UnitOfMeasureQtyField initialized with val
func NewUnitOfMeasureQty(val float64) *UnitOfMeasureQtyField {
	field := &UnitOfMeasureQtyField{}
	field.Value = val
	return field
}

//UnsolicitedIndicatorField is a BOOLEAN field
type UnsolicitedIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.UnsolicitedIndicator (325)
func (f UnsolicitedIndicatorField) Tag() quickfix.Tag { return tag.UnsolicitedIndicator }

//NewUnsolicitedIndicator returns a new UnsolicitedIndicatorField initialized with val
func NewUnsolicitedIndicator(val bool) *UnsolicitedIndicatorField {
	field := &UnsolicitedIndicatorField{}
	field.Value = val
	return field
}

//UrgencyField is a CHAR field
type UrgencyField struct{ quickfix.CharValue }

//Tag returns tag.Urgency (61)
func (f UrgencyField) Tag() quickfix.Tag { return tag.Urgency }

//NewUrgency returns a new UrgencyField initialized with val
func NewUrgency(val string) *UrgencyField {
	field := &UrgencyField{}
	field.Value = val
	return field
}

//UserRequestIDField is a STRING field
type UserRequestIDField struct{ quickfix.StringValue }

//Tag returns tag.UserRequestID (923)
func (f UserRequestIDField) Tag() quickfix.Tag { return tag.UserRequestID }

//NewUserRequestID returns a new UserRequestIDField initialized with val
func NewUserRequestID(val string) *UserRequestIDField {
	field := &UserRequestIDField{}
	field.Value = val
	return field
}

//UserRequestTypeField is a INT field
type UserRequestTypeField struct{ quickfix.IntValue }

//Tag returns tag.UserRequestType (924)
func (f UserRequestTypeField) Tag() quickfix.Tag { return tag.UserRequestType }

//NewUserRequestType returns a new UserRequestTypeField initialized with val
func NewUserRequestType(val int) *UserRequestTypeField {
	field := &UserRequestTypeField{}
	field.Value = val
	return field
}

//UserStatusField is a INT field
type UserStatusField struct{ quickfix.IntValue }

//Tag returns tag.UserStatus (926)
func (f UserStatusField) Tag() quickfix.Tag { return tag.UserStatus }

//NewUserStatus returns a new UserStatusField initialized with val
func NewUserStatus(val int) *UserStatusField {
	field := &UserStatusField{}
	field.Value = val
	return field
}

//UserStatusTextField is a STRING field
type UserStatusTextField struct{ quickfix.StringValue }

//Tag returns tag.UserStatusText (927)
func (f UserStatusTextField) Tag() quickfix.Tag { return tag.UserStatusText }

//NewUserStatusText returns a new UserStatusTextField initialized with val
func NewUserStatusText(val string) *UserStatusTextField {
	field := &UserStatusTextField{}
	field.Value = val
	return field
}

//UsernameField is a STRING field
type UsernameField struct{ quickfix.StringValue }

//Tag returns tag.Username (553)
func (f UsernameField) Tag() quickfix.Tag { return tag.Username }

//NewUsername returns a new UsernameField initialized with val
func NewUsername(val string) *UsernameField {
	field := &UsernameField{}
	field.Value = val
	return field
}

//ValidUntilTimeField is a UTCTIMESTAMP field
type ValidUntilTimeField struct{ quickfix.UTCTimestampValue }

//Tag returns tag.ValidUntilTime (62)
func (f ValidUntilTimeField) Tag() quickfix.Tag { return tag.ValidUntilTime }

//ValuationMethodField is a STRING field
type ValuationMethodField struct{ quickfix.StringValue }

//Tag returns tag.ValuationMethod (1197)
func (f ValuationMethodField) Tag() quickfix.Tag { return tag.ValuationMethod }

//NewValuationMethod returns a new ValuationMethodField initialized with val
func NewValuationMethod(val string) *ValuationMethodField {
	field := &ValuationMethodField{}
	field.Value = val
	return field
}

//ValueOfFuturesField is a AMT field
type ValueOfFuturesField struct{ quickfix.AmtValue }

//Tag returns tag.ValueOfFutures (408)
func (f ValueOfFuturesField) Tag() quickfix.Tag { return tag.ValueOfFutures }

//NewValueOfFutures returns a new ValueOfFuturesField initialized with val
func NewValueOfFutures(val float64) *ValueOfFuturesField {
	field := &ValueOfFuturesField{}
	field.Value = val
	return field
}

//VenueTypeField is a CHAR field
type VenueTypeField struct{ quickfix.CharValue }

//Tag returns tag.VenueType (1430)
func (f VenueTypeField) Tag() quickfix.Tag { return tag.VenueType }

//NewVenueType returns a new VenueTypeField initialized with val
func NewVenueType(val string) *VenueTypeField {
	field := &VenueTypeField{}
	field.Value = val
	return field
}

//VolatilityField is a FLOAT field
type VolatilityField struct{ quickfix.FloatValue }

//Tag returns tag.Volatility (1188)
func (f VolatilityField) Tag() quickfix.Tag { return tag.Volatility }

//NewVolatility returns a new VolatilityField initialized with val
func NewVolatility(val float64) *VolatilityField {
	field := &VolatilityField{}
	field.Value = val
	return field
}

//WaveNoField is a STRING field
type WaveNoField struct{ quickfix.StringValue }

//Tag returns tag.WaveNo (105)
func (f WaveNoField) Tag() quickfix.Tag { return tag.WaveNo }

//NewWaveNo returns a new WaveNoField initialized with val
func NewWaveNo(val string) *WaveNoField {
	field := &WaveNoField{}
	field.Value = val
	return field
}

//WorkingIndicatorField is a BOOLEAN field
type WorkingIndicatorField struct{ quickfix.BooleanValue }

//Tag returns tag.WorkingIndicator (636)
func (f WorkingIndicatorField) Tag() quickfix.Tag { return tag.WorkingIndicator }

//NewWorkingIndicator returns a new WorkingIndicatorField initialized with val
func NewWorkingIndicator(val bool) *WorkingIndicatorField {
	field := &WorkingIndicatorField{}
	field.Value = val
	return field
}

//WtAverageLiquidityField is a PERCENTAGE field
type WtAverageLiquidityField struct{ quickfix.PercentageValue }

//Tag returns tag.WtAverageLiquidity (410)
func (f WtAverageLiquidityField) Tag() quickfix.Tag { return tag.WtAverageLiquidity }

//NewWtAverageLiquidity returns a new WtAverageLiquidityField initialized with val
func NewWtAverageLiquidity(val float64) *WtAverageLiquidityField {
	field := &WtAverageLiquidityField{}
	field.Value = val
	return field
}

//XmlDataField is a DATA field
type XmlDataField struct{ quickfix.DataValue }

//Tag returns tag.XmlData (213)
func (f XmlDataField) Tag() quickfix.Tag { return tag.XmlData }

//NewXmlData returns a new XmlDataField initialized with val
func NewXmlData(val string) *XmlDataField {
	field := &XmlDataField{}
	field.Value = val
	return field
}

//XmlDataLenField is a LENGTH field
type XmlDataLenField struct{ quickfix.LengthValue }

//Tag returns tag.XmlDataLen (212)
func (f XmlDataLenField) Tag() quickfix.Tag { return tag.XmlDataLen }

//NewXmlDataLen returns a new XmlDataLenField initialized with val
func NewXmlDataLen(val int) *XmlDataLenField {
	field := &XmlDataLenField{}
	field.Value = val
	return field
}

//YieldField is a PERCENTAGE field
type YieldField struct{ quickfix.PercentageValue }

//Tag returns tag.Yield (236)
func (f YieldField) Tag() quickfix.Tag { return tag.Yield }

//NewYield returns a new YieldField initialized with val
func NewYield(val float64) *YieldField {
	field := &YieldField{}
	field.Value = val
	return field
}

//YieldCalcDateField is a LOCALMKTDATE field
type YieldCalcDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.YieldCalcDate (701)
func (f YieldCalcDateField) Tag() quickfix.Tag { return tag.YieldCalcDate }

//NewYieldCalcDate returns a new YieldCalcDateField initialized with val
func NewYieldCalcDate(val string) *YieldCalcDateField {
	field := &YieldCalcDateField{}
	field.Value = val
	return field
}

//YieldRedemptionDateField is a LOCALMKTDATE field
type YieldRedemptionDateField struct{ quickfix.LocalMktDateValue }

//Tag returns tag.YieldRedemptionDate (696)
func (f YieldRedemptionDateField) Tag() quickfix.Tag { return tag.YieldRedemptionDate }

//NewYieldRedemptionDate returns a new YieldRedemptionDateField initialized with val
func NewYieldRedemptionDate(val string) *YieldRedemptionDateField {
	field := &YieldRedemptionDateField{}
	field.Value = val
	return field
}

//YieldRedemptionPriceField is a PRICE field
type YieldRedemptionPriceField struct{ quickfix.PriceValue }

//Tag returns tag.YieldRedemptionPrice (697)
func (f YieldRedemptionPriceField) Tag() quickfix.Tag { return tag.YieldRedemptionPrice }

//NewYieldRedemptionPrice returns a new YieldRedemptionPriceField initialized with val
func NewYieldRedemptionPrice(val float64) *YieldRedemptionPriceField {
	field := &YieldRedemptionPriceField{}
	field.Value = val
	return field
}

//YieldRedemptionPriceTypeField is a INT field
type YieldRedemptionPriceTypeField struct{ quickfix.IntValue }

//Tag returns tag.YieldRedemptionPriceType (698)
func (f YieldRedemptionPriceTypeField) Tag() quickfix.Tag { return tag.YieldRedemptionPriceType }

//NewYieldRedemptionPriceType returns a new YieldRedemptionPriceTypeField initialized with val
func NewYieldRedemptionPriceType(val int) *YieldRedemptionPriceTypeField {
	field := &YieldRedemptionPriceTypeField{}
	field.Value = val
	return field
}

//YieldTypeField is a STRING field
type YieldTypeField struct{ quickfix.StringValue }

//Tag returns tag.YieldType (235)
func (f YieldTypeField) Tag() quickfix.Tag { return tag.YieldType }

//NewYieldType returns a new YieldTypeField initialized with val
func NewYieldType(val string) *YieldTypeField {
	field := &YieldTypeField{}
	field.Value = val
	return field
}
