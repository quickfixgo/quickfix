package field

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/tag"

	"time"
) //AccountField is a STRING field
type AccountField struct{ quickfix.FIXString }

//Tag returns tag.Account (1)
func (f AccountField) Tag() quickfix.Tag { return tag.Account }

//NewAccount returns a new AccountField initialized with val
func NewAccount(val quickfix.FIXString) AccountField {
	return AccountField{val}
}

//AccountTypeField is a INT field
type AccountTypeField struct{ quickfix.FIXInt }

//Tag returns tag.AccountType (581)
func (f AccountTypeField) Tag() quickfix.Tag { return tag.AccountType }

//NewAccountType returns a new AccountTypeField initialized with val
func NewAccountType(val quickfix.FIXInt) AccountTypeField {
	return AccountTypeField{val}
}

//AccruedInterestAmtField is a AMT field
type AccruedInterestAmtField struct{ quickfix.FIXFloat }

//Tag returns tag.AccruedInterestAmt (159)
func (f AccruedInterestAmtField) Tag() quickfix.Tag { return tag.AccruedInterestAmt }

//NewAccruedInterestAmt returns a new AccruedInterestAmtField initialized with val
func NewAccruedInterestAmt(val quickfix.FIXFloat) AccruedInterestAmtField {
	return AccruedInterestAmtField{val}
}

//AccruedInterestRateField is a PERCENTAGE field
type AccruedInterestRateField struct{ quickfix.FIXFloat }

//Tag returns tag.AccruedInterestRate (158)
func (f AccruedInterestRateField) Tag() quickfix.Tag { return tag.AccruedInterestRate }

//NewAccruedInterestRate returns a new AccruedInterestRateField initialized with val
func NewAccruedInterestRate(val quickfix.FIXFloat) AccruedInterestRateField {
	return AccruedInterestRateField{val}
}

//AcctIDSourceField is a INT field
type AcctIDSourceField struct{ quickfix.FIXInt }

//Tag returns tag.AcctIDSource (660)
func (f AcctIDSourceField) Tag() quickfix.Tag { return tag.AcctIDSource }

//NewAcctIDSource returns a new AcctIDSourceField initialized with val
func NewAcctIDSource(val quickfix.FIXInt) AcctIDSourceField {
	return AcctIDSourceField{val}
}

//AdjustmentField is a INT field
type AdjustmentField struct{ quickfix.FIXInt }

//Tag returns tag.Adjustment (334)
func (f AdjustmentField) Tag() quickfix.Tag { return tag.Adjustment }

//NewAdjustment returns a new AdjustmentField initialized with val
func NewAdjustment(val quickfix.FIXInt) AdjustmentField {
	return AdjustmentField{val}
}

//AdjustmentTypeField is a INT field
type AdjustmentTypeField struct{ quickfix.FIXInt }

//Tag returns tag.AdjustmentType (718)
func (f AdjustmentTypeField) Tag() quickfix.Tag { return tag.AdjustmentType }

//NewAdjustmentType returns a new AdjustmentTypeField initialized with val
func NewAdjustmentType(val quickfix.FIXInt) AdjustmentTypeField {
	return AdjustmentTypeField{val}
}

//AdvIdField is a STRING field
type AdvIdField struct{ quickfix.FIXString }

//Tag returns tag.AdvId (2)
func (f AdvIdField) Tag() quickfix.Tag { return tag.AdvId }

//NewAdvId returns a new AdvIdField initialized with val
func NewAdvId(val quickfix.FIXString) AdvIdField {
	return AdvIdField{val}
}

//AdvRefIDField is a STRING field
type AdvRefIDField struct{ quickfix.FIXString }

//Tag returns tag.AdvRefID (3)
func (f AdvRefIDField) Tag() quickfix.Tag { return tag.AdvRefID }

//NewAdvRefID returns a new AdvRefIDField initialized with val
func NewAdvRefID(val quickfix.FIXString) AdvRefIDField {
	return AdvRefIDField{val}
}

//AdvSideField is a CHAR field
type AdvSideField struct{ quickfix.FIXString }

//Tag returns tag.AdvSide (4)
func (f AdvSideField) Tag() quickfix.Tag { return tag.AdvSide }

//NewAdvSide returns a new AdvSideField initialized with val
func NewAdvSide(val quickfix.FIXString) AdvSideField {
	return AdvSideField{val}
}

//AdvTransTypeField is a STRING field
type AdvTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.AdvTransType (5)
func (f AdvTransTypeField) Tag() quickfix.Tag { return tag.AdvTransType }

//NewAdvTransType returns a new AdvTransTypeField initialized with val
func NewAdvTransType(val quickfix.FIXString) AdvTransTypeField {
	return AdvTransTypeField{val}
}

//AffectedOrderIDField is a STRING field
type AffectedOrderIDField struct{ quickfix.FIXString }

//Tag returns tag.AffectedOrderID (535)
func (f AffectedOrderIDField) Tag() quickfix.Tag { return tag.AffectedOrderID }

//NewAffectedOrderID returns a new AffectedOrderIDField initialized with val
func NewAffectedOrderID(val quickfix.FIXString) AffectedOrderIDField {
	return AffectedOrderIDField{val}
}

//AffectedSecondaryOrderIDField is a STRING field
type AffectedSecondaryOrderIDField struct{ quickfix.FIXString }

//Tag returns tag.AffectedSecondaryOrderID (536)
func (f AffectedSecondaryOrderIDField) Tag() quickfix.Tag { return tag.AffectedSecondaryOrderID }

//NewAffectedSecondaryOrderID returns a new AffectedSecondaryOrderIDField initialized with val
func NewAffectedSecondaryOrderID(val quickfix.FIXString) AffectedSecondaryOrderIDField {
	return AffectedSecondaryOrderIDField{val}
}

//AffirmStatusField is a INT field
type AffirmStatusField struct{ quickfix.FIXInt }

//Tag returns tag.AffirmStatus (940)
func (f AffirmStatusField) Tag() quickfix.Tag { return tag.AffirmStatus }

//NewAffirmStatus returns a new AffirmStatusField initialized with val
func NewAffirmStatus(val quickfix.FIXInt) AffirmStatusField {
	return AffirmStatusField{val}
}

//AggregatedBookField is a BOOLEAN field
type AggregatedBookField struct{ quickfix.FIXBoolean }

//Tag returns tag.AggregatedBook (266)
func (f AggregatedBookField) Tag() quickfix.Tag { return tag.AggregatedBook }

//NewAggregatedBook returns a new AggregatedBookField initialized with val
func NewAggregatedBook(val quickfix.FIXBoolean) AggregatedBookField {
	return AggregatedBookField{val}
}

//AggressorIndicatorField is a BOOLEAN field
type AggressorIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.AggressorIndicator (1057)
func (f AggressorIndicatorField) Tag() quickfix.Tag { return tag.AggressorIndicator }

//NewAggressorIndicator returns a new AggressorIndicatorField initialized with val
func NewAggressorIndicator(val quickfix.FIXBoolean) AggressorIndicatorField {
	return AggressorIndicatorField{val}
}

//AgreementCurrencyField is a CURRENCY field
type AgreementCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.AgreementCurrency (918)
func (f AgreementCurrencyField) Tag() quickfix.Tag { return tag.AgreementCurrency }

//NewAgreementCurrency returns a new AgreementCurrencyField initialized with val
func NewAgreementCurrency(val quickfix.FIXString) AgreementCurrencyField {
	return AgreementCurrencyField{val}
}

//AgreementDateField is a LOCALMKTDATE field
type AgreementDateField struct{ quickfix.FIXString }

//Tag returns tag.AgreementDate (915)
func (f AgreementDateField) Tag() quickfix.Tag { return tag.AgreementDate }

//NewAgreementDate returns a new AgreementDateField initialized with val
func NewAgreementDate(val quickfix.FIXString) AgreementDateField {
	return AgreementDateField{val}
}

//AgreementDescField is a STRING field
type AgreementDescField struct{ quickfix.FIXString }

//Tag returns tag.AgreementDesc (913)
func (f AgreementDescField) Tag() quickfix.Tag { return tag.AgreementDesc }

//NewAgreementDesc returns a new AgreementDescField initialized with val
func NewAgreementDesc(val quickfix.FIXString) AgreementDescField {
	return AgreementDescField{val}
}

//AgreementIDField is a STRING field
type AgreementIDField struct{ quickfix.FIXString }

//Tag returns tag.AgreementID (914)
func (f AgreementIDField) Tag() quickfix.Tag { return tag.AgreementID }

//NewAgreementID returns a new AgreementIDField initialized with val
func NewAgreementID(val quickfix.FIXString) AgreementIDField {
	return AgreementIDField{val}
}

//AllocAccountField is a STRING field
type AllocAccountField struct{ quickfix.FIXString }

//Tag returns tag.AllocAccount (79)
func (f AllocAccountField) Tag() quickfix.Tag { return tag.AllocAccount }

//NewAllocAccount returns a new AllocAccountField initialized with val
func NewAllocAccount(val quickfix.FIXString) AllocAccountField {
	return AllocAccountField{val}
}

//AllocAccountTypeField is a INT field
type AllocAccountTypeField struct{ quickfix.FIXInt }

//Tag returns tag.AllocAccountType (798)
func (f AllocAccountTypeField) Tag() quickfix.Tag { return tag.AllocAccountType }

//NewAllocAccountType returns a new AllocAccountTypeField initialized with val
func NewAllocAccountType(val quickfix.FIXInt) AllocAccountTypeField {
	return AllocAccountTypeField{val}
}

//AllocAccruedInterestAmtField is a AMT field
type AllocAccruedInterestAmtField struct{ quickfix.FIXFloat }

//Tag returns tag.AllocAccruedInterestAmt (742)
func (f AllocAccruedInterestAmtField) Tag() quickfix.Tag { return tag.AllocAccruedInterestAmt }

//NewAllocAccruedInterestAmt returns a new AllocAccruedInterestAmtField initialized with val
func NewAllocAccruedInterestAmt(val quickfix.FIXFloat) AllocAccruedInterestAmtField {
	return AllocAccruedInterestAmtField{val}
}

//AllocAcctIDSourceField is a INT field
type AllocAcctIDSourceField struct{ quickfix.FIXInt }

//Tag returns tag.AllocAcctIDSource (661)
func (f AllocAcctIDSourceField) Tag() quickfix.Tag { return tag.AllocAcctIDSource }

//NewAllocAcctIDSource returns a new AllocAcctIDSourceField initialized with val
func NewAllocAcctIDSource(val quickfix.FIXInt) AllocAcctIDSourceField {
	return AllocAcctIDSourceField{val}
}

//AllocAvgPxField is a PRICE field
type AllocAvgPxField struct{ quickfix.FIXFloat }

//Tag returns tag.AllocAvgPx (153)
func (f AllocAvgPxField) Tag() quickfix.Tag { return tag.AllocAvgPx }

//NewAllocAvgPx returns a new AllocAvgPxField initialized with val
func NewAllocAvgPx(val quickfix.FIXFloat) AllocAvgPxField {
	return AllocAvgPxField{val}
}

//AllocCancReplaceReasonField is a INT field
type AllocCancReplaceReasonField struct{ quickfix.FIXInt }

//Tag returns tag.AllocCancReplaceReason (796)
func (f AllocCancReplaceReasonField) Tag() quickfix.Tag { return tag.AllocCancReplaceReason }

//NewAllocCancReplaceReason returns a new AllocCancReplaceReasonField initialized with val
func NewAllocCancReplaceReason(val quickfix.FIXInt) AllocCancReplaceReasonField {
	return AllocCancReplaceReasonField{val}
}

//AllocClearingFeeIndicatorField is a STRING field
type AllocClearingFeeIndicatorField struct{ quickfix.FIXString }

//Tag returns tag.AllocClearingFeeIndicator (1136)
func (f AllocClearingFeeIndicatorField) Tag() quickfix.Tag { return tag.AllocClearingFeeIndicator }

//NewAllocClearingFeeIndicator returns a new AllocClearingFeeIndicatorField initialized with val
func NewAllocClearingFeeIndicator(val quickfix.FIXString) AllocClearingFeeIndicatorField {
	return AllocClearingFeeIndicatorField{val}
}

//AllocCustomerCapacityField is a STRING field
type AllocCustomerCapacityField struct{ quickfix.FIXString }

//Tag returns tag.AllocCustomerCapacity (993)
func (f AllocCustomerCapacityField) Tag() quickfix.Tag { return tag.AllocCustomerCapacity }

//NewAllocCustomerCapacity returns a new AllocCustomerCapacityField initialized with val
func NewAllocCustomerCapacity(val quickfix.FIXString) AllocCustomerCapacityField {
	return AllocCustomerCapacityField{val}
}

//AllocHandlInstField is a INT field
type AllocHandlInstField struct{ quickfix.FIXInt }

//Tag returns tag.AllocHandlInst (209)
func (f AllocHandlInstField) Tag() quickfix.Tag { return tag.AllocHandlInst }

//NewAllocHandlInst returns a new AllocHandlInstField initialized with val
func NewAllocHandlInst(val quickfix.FIXInt) AllocHandlInstField {
	return AllocHandlInstField{val}
}

//AllocIDField is a STRING field
type AllocIDField struct{ quickfix.FIXString }

//Tag returns tag.AllocID (70)
func (f AllocIDField) Tag() quickfix.Tag { return tag.AllocID }

//NewAllocID returns a new AllocIDField initialized with val
func NewAllocID(val quickfix.FIXString) AllocIDField {
	return AllocIDField{val}
}

//AllocInterestAtMaturityField is a AMT field
type AllocInterestAtMaturityField struct{ quickfix.FIXFloat }

//Tag returns tag.AllocInterestAtMaturity (741)
func (f AllocInterestAtMaturityField) Tag() quickfix.Tag { return tag.AllocInterestAtMaturity }

//NewAllocInterestAtMaturity returns a new AllocInterestAtMaturityField initialized with val
func NewAllocInterestAtMaturity(val quickfix.FIXFloat) AllocInterestAtMaturityField {
	return AllocInterestAtMaturityField{val}
}

//AllocIntermedReqTypeField is a INT field
type AllocIntermedReqTypeField struct{ quickfix.FIXInt }

//Tag returns tag.AllocIntermedReqType (808)
func (f AllocIntermedReqTypeField) Tag() quickfix.Tag { return tag.AllocIntermedReqType }

//NewAllocIntermedReqType returns a new AllocIntermedReqTypeField initialized with val
func NewAllocIntermedReqType(val quickfix.FIXInt) AllocIntermedReqTypeField {
	return AllocIntermedReqTypeField{val}
}

//AllocLinkIDField is a STRING field
type AllocLinkIDField struct{ quickfix.FIXString }

//Tag returns tag.AllocLinkID (196)
func (f AllocLinkIDField) Tag() quickfix.Tag { return tag.AllocLinkID }

//NewAllocLinkID returns a new AllocLinkIDField initialized with val
func NewAllocLinkID(val quickfix.FIXString) AllocLinkIDField {
	return AllocLinkIDField{val}
}

//AllocLinkTypeField is a INT field
type AllocLinkTypeField struct{ quickfix.FIXInt }

//Tag returns tag.AllocLinkType (197)
func (f AllocLinkTypeField) Tag() quickfix.Tag { return tag.AllocLinkType }

//NewAllocLinkType returns a new AllocLinkTypeField initialized with val
func NewAllocLinkType(val quickfix.FIXInt) AllocLinkTypeField {
	return AllocLinkTypeField{val}
}

//AllocMethodField is a INT field
type AllocMethodField struct{ quickfix.FIXInt }

//Tag returns tag.AllocMethod (1002)
func (f AllocMethodField) Tag() quickfix.Tag { return tag.AllocMethod }

//NewAllocMethod returns a new AllocMethodField initialized with val
func NewAllocMethod(val quickfix.FIXInt) AllocMethodField {
	return AllocMethodField{val}
}

//AllocNetMoneyField is a AMT field
type AllocNetMoneyField struct{ quickfix.FIXFloat }

//Tag returns tag.AllocNetMoney (154)
func (f AllocNetMoneyField) Tag() quickfix.Tag { return tag.AllocNetMoney }

//NewAllocNetMoney returns a new AllocNetMoneyField initialized with val
func NewAllocNetMoney(val quickfix.FIXFloat) AllocNetMoneyField {
	return AllocNetMoneyField{val}
}

//AllocNoOrdersTypeField is a INT field
type AllocNoOrdersTypeField struct{ quickfix.FIXInt }

//Tag returns tag.AllocNoOrdersType (857)
func (f AllocNoOrdersTypeField) Tag() quickfix.Tag { return tag.AllocNoOrdersType }

//NewAllocNoOrdersType returns a new AllocNoOrdersTypeField initialized with val
func NewAllocNoOrdersType(val quickfix.FIXInt) AllocNoOrdersTypeField {
	return AllocNoOrdersTypeField{val}
}

//AllocPositionEffectField is a CHAR field
type AllocPositionEffectField struct{ quickfix.FIXString }

//Tag returns tag.AllocPositionEffect (1047)
func (f AllocPositionEffectField) Tag() quickfix.Tag { return tag.AllocPositionEffect }

//NewAllocPositionEffect returns a new AllocPositionEffectField initialized with val
func NewAllocPositionEffect(val quickfix.FIXString) AllocPositionEffectField {
	return AllocPositionEffectField{val}
}

//AllocPriceField is a PRICE field
type AllocPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.AllocPrice (366)
func (f AllocPriceField) Tag() quickfix.Tag { return tag.AllocPrice }

//NewAllocPrice returns a new AllocPriceField initialized with val
func NewAllocPrice(val quickfix.FIXFloat) AllocPriceField {
	return AllocPriceField{val}
}

//AllocQtyField is a QTY field
type AllocQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.AllocQty (80)
func (f AllocQtyField) Tag() quickfix.Tag { return tag.AllocQty }

//NewAllocQty returns a new AllocQtyField initialized with val
func NewAllocQty(val quickfix.FIXFloat) AllocQtyField {
	return AllocQtyField{val}
}

//AllocRejCodeField is a INT field
type AllocRejCodeField struct{ quickfix.FIXInt }

//Tag returns tag.AllocRejCode (88)
func (f AllocRejCodeField) Tag() quickfix.Tag { return tag.AllocRejCode }

//NewAllocRejCode returns a new AllocRejCodeField initialized with val
func NewAllocRejCode(val quickfix.FIXInt) AllocRejCodeField {
	return AllocRejCodeField{val}
}

//AllocReportIDField is a STRING field
type AllocReportIDField struct{ quickfix.FIXString }

//Tag returns tag.AllocReportID (755)
func (f AllocReportIDField) Tag() quickfix.Tag { return tag.AllocReportID }

//NewAllocReportID returns a new AllocReportIDField initialized with val
func NewAllocReportID(val quickfix.FIXString) AllocReportIDField {
	return AllocReportIDField{val}
}

//AllocReportRefIDField is a STRING field
type AllocReportRefIDField struct{ quickfix.FIXString }

//Tag returns tag.AllocReportRefID (795)
func (f AllocReportRefIDField) Tag() quickfix.Tag { return tag.AllocReportRefID }

//NewAllocReportRefID returns a new AllocReportRefIDField initialized with val
func NewAllocReportRefID(val quickfix.FIXString) AllocReportRefIDField {
	return AllocReportRefIDField{val}
}

//AllocReportTypeField is a INT field
type AllocReportTypeField struct{ quickfix.FIXInt }

//Tag returns tag.AllocReportType (794)
func (f AllocReportTypeField) Tag() quickfix.Tag { return tag.AllocReportType }

//NewAllocReportType returns a new AllocReportTypeField initialized with val
func NewAllocReportType(val quickfix.FIXInt) AllocReportTypeField {
	return AllocReportTypeField{val}
}

//AllocSettlCurrAmtField is a AMT field
type AllocSettlCurrAmtField struct{ quickfix.FIXFloat }

//Tag returns tag.AllocSettlCurrAmt (737)
func (f AllocSettlCurrAmtField) Tag() quickfix.Tag { return tag.AllocSettlCurrAmt }

//NewAllocSettlCurrAmt returns a new AllocSettlCurrAmtField initialized with val
func NewAllocSettlCurrAmt(val quickfix.FIXFloat) AllocSettlCurrAmtField {
	return AllocSettlCurrAmtField{val}
}

//AllocSettlCurrencyField is a CURRENCY field
type AllocSettlCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.AllocSettlCurrency (736)
func (f AllocSettlCurrencyField) Tag() quickfix.Tag { return tag.AllocSettlCurrency }

//NewAllocSettlCurrency returns a new AllocSettlCurrencyField initialized with val
func NewAllocSettlCurrency(val quickfix.FIXString) AllocSettlCurrencyField {
	return AllocSettlCurrencyField{val}
}

//AllocSettlInstTypeField is a INT field
type AllocSettlInstTypeField struct{ quickfix.FIXInt }

//Tag returns tag.AllocSettlInstType (780)
func (f AllocSettlInstTypeField) Tag() quickfix.Tag { return tag.AllocSettlInstType }

//NewAllocSettlInstType returns a new AllocSettlInstTypeField initialized with val
func NewAllocSettlInstType(val quickfix.FIXInt) AllocSettlInstTypeField {
	return AllocSettlInstTypeField{val}
}

//AllocSharesField is a QTY field
type AllocSharesField struct{ quickfix.FIXFloat }

//Tag returns tag.AllocShares (80)
func (f AllocSharesField) Tag() quickfix.Tag { return tag.AllocShares }

//NewAllocShares returns a new AllocSharesField initialized with val
func NewAllocShares(val quickfix.FIXFloat) AllocSharesField {
	return AllocSharesField{val}
}

//AllocStatusField is a INT field
type AllocStatusField struct{ quickfix.FIXInt }

//Tag returns tag.AllocStatus (87)
func (f AllocStatusField) Tag() quickfix.Tag { return tag.AllocStatus }

//NewAllocStatus returns a new AllocStatusField initialized with val
func NewAllocStatus(val quickfix.FIXInt) AllocStatusField {
	return AllocStatusField{val}
}

//AllocTextField is a STRING field
type AllocTextField struct{ quickfix.FIXString }

//Tag returns tag.AllocText (161)
func (f AllocTextField) Tag() quickfix.Tag { return tag.AllocText }

//NewAllocText returns a new AllocTextField initialized with val
func NewAllocText(val quickfix.FIXString) AllocTextField {
	return AllocTextField{val}
}

//AllocTransTypeField is a CHAR field
type AllocTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.AllocTransType (71)
func (f AllocTransTypeField) Tag() quickfix.Tag { return tag.AllocTransType }

//NewAllocTransType returns a new AllocTransTypeField initialized with val
func NewAllocTransType(val quickfix.FIXString) AllocTransTypeField {
	return AllocTransTypeField{val}
}

//AllocTypeField is a INT field
type AllocTypeField struct{ quickfix.FIXInt }

//Tag returns tag.AllocType (626)
func (f AllocTypeField) Tag() quickfix.Tag { return tag.AllocType }

//NewAllocType returns a new AllocTypeField initialized with val
func NewAllocType(val quickfix.FIXInt) AllocTypeField {
	return AllocTypeField{val}
}

//AllowableOneSidednessCurrField is a CURRENCY field
type AllowableOneSidednessCurrField struct{ quickfix.FIXString }

//Tag returns tag.AllowableOneSidednessCurr (767)
func (f AllowableOneSidednessCurrField) Tag() quickfix.Tag { return tag.AllowableOneSidednessCurr }

//NewAllowableOneSidednessCurr returns a new AllowableOneSidednessCurrField initialized with val
func NewAllowableOneSidednessCurr(val quickfix.FIXString) AllowableOneSidednessCurrField {
	return AllowableOneSidednessCurrField{val}
}

//AllowableOneSidednessPctField is a PERCENTAGE field
type AllowableOneSidednessPctField struct{ quickfix.FIXFloat }

//Tag returns tag.AllowableOneSidednessPct (765)
func (f AllowableOneSidednessPctField) Tag() quickfix.Tag { return tag.AllowableOneSidednessPct }

//NewAllowableOneSidednessPct returns a new AllowableOneSidednessPctField initialized with val
func NewAllowableOneSidednessPct(val quickfix.FIXFloat) AllowableOneSidednessPctField {
	return AllowableOneSidednessPctField{val}
}

//AllowableOneSidednessValueField is a AMT field
type AllowableOneSidednessValueField struct{ quickfix.FIXFloat }

//Tag returns tag.AllowableOneSidednessValue (766)
func (f AllowableOneSidednessValueField) Tag() quickfix.Tag { return tag.AllowableOneSidednessValue }

//NewAllowableOneSidednessValue returns a new AllowableOneSidednessValueField initialized with val
func NewAllowableOneSidednessValue(val quickfix.FIXFloat) AllowableOneSidednessValueField {
	return AllowableOneSidednessValueField{val}
}

//AltMDSourceIDField is a STRING field
type AltMDSourceIDField struct{ quickfix.FIXString }

//Tag returns tag.AltMDSourceID (817)
func (f AltMDSourceIDField) Tag() quickfix.Tag { return tag.AltMDSourceID }

//NewAltMDSourceID returns a new AltMDSourceIDField initialized with val
func NewAltMDSourceID(val quickfix.FIXString) AltMDSourceIDField {
	return AltMDSourceIDField{val}
}

//ApplBegSeqNumField is a SEQNUM field
type ApplBegSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.ApplBegSeqNum (1182)
func (f ApplBegSeqNumField) Tag() quickfix.Tag { return tag.ApplBegSeqNum }

//NewApplBegSeqNum returns a new ApplBegSeqNumField initialized with val
func NewApplBegSeqNum(val quickfix.FIXInt) ApplBegSeqNumField {
	return ApplBegSeqNumField{val}
}

//ApplEndSeqNumField is a SEQNUM field
type ApplEndSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.ApplEndSeqNum (1183)
func (f ApplEndSeqNumField) Tag() quickfix.Tag { return tag.ApplEndSeqNum }

//NewApplEndSeqNum returns a new ApplEndSeqNumField initialized with val
func NewApplEndSeqNum(val quickfix.FIXInt) ApplEndSeqNumField {
	return ApplEndSeqNumField{val}
}

//ApplExtIDField is a INT field
type ApplExtIDField struct{ quickfix.FIXInt }

//Tag returns tag.ApplExtID (1156)
func (f ApplExtIDField) Tag() quickfix.Tag { return tag.ApplExtID }

//NewApplExtID returns a new ApplExtIDField initialized with val
func NewApplExtID(val quickfix.FIXInt) ApplExtIDField {
	return ApplExtIDField{val}
}

//ApplIDField is a STRING field
type ApplIDField struct{ quickfix.FIXString }

//Tag returns tag.ApplID (1180)
func (f ApplIDField) Tag() quickfix.Tag { return tag.ApplID }

//NewApplID returns a new ApplIDField initialized with val
func NewApplID(val quickfix.FIXString) ApplIDField {
	return ApplIDField{val}
}

//ApplLastSeqNumField is a SEQNUM field
type ApplLastSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.ApplLastSeqNum (1350)
func (f ApplLastSeqNumField) Tag() quickfix.Tag { return tag.ApplLastSeqNum }

//NewApplLastSeqNum returns a new ApplLastSeqNumField initialized with val
func NewApplLastSeqNum(val quickfix.FIXInt) ApplLastSeqNumField {
	return ApplLastSeqNumField{val}
}

//ApplNewSeqNumField is a SEQNUM field
type ApplNewSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.ApplNewSeqNum (1399)
func (f ApplNewSeqNumField) Tag() quickfix.Tag { return tag.ApplNewSeqNum }

//NewApplNewSeqNum returns a new ApplNewSeqNumField initialized with val
func NewApplNewSeqNum(val quickfix.FIXInt) ApplNewSeqNumField {
	return ApplNewSeqNumField{val}
}

//ApplQueueActionField is a INT field
type ApplQueueActionField struct{ quickfix.FIXInt }

//Tag returns tag.ApplQueueAction (815)
func (f ApplQueueActionField) Tag() quickfix.Tag { return tag.ApplQueueAction }

//NewApplQueueAction returns a new ApplQueueActionField initialized with val
func NewApplQueueAction(val quickfix.FIXInt) ApplQueueActionField {
	return ApplQueueActionField{val}
}

//ApplQueueDepthField is a INT field
type ApplQueueDepthField struct{ quickfix.FIXInt }

//Tag returns tag.ApplQueueDepth (813)
func (f ApplQueueDepthField) Tag() quickfix.Tag { return tag.ApplQueueDepth }

//NewApplQueueDepth returns a new ApplQueueDepthField initialized with val
func NewApplQueueDepth(val quickfix.FIXInt) ApplQueueDepthField {
	return ApplQueueDepthField{val}
}

//ApplQueueMaxField is a INT field
type ApplQueueMaxField struct{ quickfix.FIXInt }

//Tag returns tag.ApplQueueMax (812)
func (f ApplQueueMaxField) Tag() quickfix.Tag { return tag.ApplQueueMax }

//NewApplQueueMax returns a new ApplQueueMaxField initialized with val
func NewApplQueueMax(val quickfix.FIXInt) ApplQueueMaxField {
	return ApplQueueMaxField{val}
}

//ApplQueueResolutionField is a INT field
type ApplQueueResolutionField struct{ quickfix.FIXInt }

//Tag returns tag.ApplQueueResolution (814)
func (f ApplQueueResolutionField) Tag() quickfix.Tag { return tag.ApplQueueResolution }

//NewApplQueueResolution returns a new ApplQueueResolutionField initialized with val
func NewApplQueueResolution(val quickfix.FIXInt) ApplQueueResolutionField {
	return ApplQueueResolutionField{val}
}

//ApplReportIDField is a STRING field
type ApplReportIDField struct{ quickfix.FIXString }

//Tag returns tag.ApplReportID (1356)
func (f ApplReportIDField) Tag() quickfix.Tag { return tag.ApplReportID }

//NewApplReportID returns a new ApplReportIDField initialized with val
func NewApplReportID(val quickfix.FIXString) ApplReportIDField {
	return ApplReportIDField{val}
}

//ApplReportTypeField is a INT field
type ApplReportTypeField struct{ quickfix.FIXInt }

//Tag returns tag.ApplReportType (1426)
func (f ApplReportTypeField) Tag() quickfix.Tag { return tag.ApplReportType }

//NewApplReportType returns a new ApplReportTypeField initialized with val
func NewApplReportType(val quickfix.FIXInt) ApplReportTypeField {
	return ApplReportTypeField{val}
}

//ApplReqIDField is a STRING field
type ApplReqIDField struct{ quickfix.FIXString }

//Tag returns tag.ApplReqID (1346)
func (f ApplReqIDField) Tag() quickfix.Tag { return tag.ApplReqID }

//NewApplReqID returns a new ApplReqIDField initialized with val
func NewApplReqID(val quickfix.FIXString) ApplReqIDField {
	return ApplReqIDField{val}
}

//ApplReqTypeField is a INT field
type ApplReqTypeField struct{ quickfix.FIXInt }

//Tag returns tag.ApplReqType (1347)
func (f ApplReqTypeField) Tag() quickfix.Tag { return tag.ApplReqType }

//NewApplReqType returns a new ApplReqTypeField initialized with val
func NewApplReqType(val quickfix.FIXInt) ApplReqTypeField {
	return ApplReqTypeField{val}
}

//ApplResendFlagField is a BOOLEAN field
type ApplResendFlagField struct{ quickfix.FIXBoolean }

//Tag returns tag.ApplResendFlag (1352)
func (f ApplResendFlagField) Tag() quickfix.Tag { return tag.ApplResendFlag }

//NewApplResendFlag returns a new ApplResendFlagField initialized with val
func NewApplResendFlag(val quickfix.FIXBoolean) ApplResendFlagField {
	return ApplResendFlagField{val}
}

//ApplResponseErrorField is a INT field
type ApplResponseErrorField struct{ quickfix.FIXInt }

//Tag returns tag.ApplResponseError (1354)
func (f ApplResponseErrorField) Tag() quickfix.Tag { return tag.ApplResponseError }

//NewApplResponseError returns a new ApplResponseErrorField initialized with val
func NewApplResponseError(val quickfix.FIXInt) ApplResponseErrorField {
	return ApplResponseErrorField{val}
}

//ApplResponseIDField is a STRING field
type ApplResponseIDField struct{ quickfix.FIXString }

//Tag returns tag.ApplResponseID (1353)
func (f ApplResponseIDField) Tag() quickfix.Tag { return tag.ApplResponseID }

//NewApplResponseID returns a new ApplResponseIDField initialized with val
func NewApplResponseID(val quickfix.FIXString) ApplResponseIDField {
	return ApplResponseIDField{val}
}

//ApplResponseTypeField is a INT field
type ApplResponseTypeField struct{ quickfix.FIXInt }

//Tag returns tag.ApplResponseType (1348)
func (f ApplResponseTypeField) Tag() quickfix.Tag { return tag.ApplResponseType }

//NewApplResponseType returns a new ApplResponseTypeField initialized with val
func NewApplResponseType(val quickfix.FIXInt) ApplResponseTypeField {
	return ApplResponseTypeField{val}
}

//ApplSeqNumField is a SEQNUM field
type ApplSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.ApplSeqNum (1181)
func (f ApplSeqNumField) Tag() quickfix.Tag { return tag.ApplSeqNum }

//NewApplSeqNum returns a new ApplSeqNumField initialized with val
func NewApplSeqNum(val quickfix.FIXInt) ApplSeqNumField {
	return ApplSeqNumField{val}
}

//ApplTotalMessageCountField is a INT field
type ApplTotalMessageCountField struct{ quickfix.FIXInt }

//Tag returns tag.ApplTotalMessageCount (1349)
func (f ApplTotalMessageCountField) Tag() quickfix.Tag { return tag.ApplTotalMessageCount }

//NewApplTotalMessageCount returns a new ApplTotalMessageCountField initialized with val
func NewApplTotalMessageCount(val quickfix.FIXInt) ApplTotalMessageCountField {
	return ApplTotalMessageCountField{val}
}

//ApplVerIDField is a STRING field
type ApplVerIDField struct{ quickfix.FIXString }

//Tag returns tag.ApplVerID (1128)
func (f ApplVerIDField) Tag() quickfix.Tag { return tag.ApplVerID }

//NewApplVerID returns a new ApplVerIDField initialized with val
func NewApplVerID(val quickfix.FIXString) ApplVerIDField {
	return ApplVerIDField{val}
}

//AsOfIndicatorField is a CHAR field
type AsOfIndicatorField struct{ quickfix.FIXString }

//Tag returns tag.AsOfIndicator (1015)
func (f AsOfIndicatorField) Tag() quickfix.Tag { return tag.AsOfIndicator }

//NewAsOfIndicator returns a new AsOfIndicatorField initialized with val
func NewAsOfIndicator(val quickfix.FIXString) AsOfIndicatorField {
	return AsOfIndicatorField{val}
}

//AsgnReqIDField is a STRING field
type AsgnReqIDField struct{ quickfix.FIXString }

//Tag returns tag.AsgnReqID (831)
func (f AsgnReqIDField) Tag() quickfix.Tag { return tag.AsgnReqID }

//NewAsgnReqID returns a new AsgnReqIDField initialized with val
func NewAsgnReqID(val quickfix.FIXString) AsgnReqIDField {
	return AsgnReqIDField{val}
}

//AsgnRptIDField is a STRING field
type AsgnRptIDField struct{ quickfix.FIXString }

//Tag returns tag.AsgnRptID (833)
func (f AsgnRptIDField) Tag() quickfix.Tag { return tag.AsgnRptID }

//NewAsgnRptID returns a new AsgnRptIDField initialized with val
func NewAsgnRptID(val quickfix.FIXString) AsgnRptIDField {
	return AsgnRptIDField{val}
}

//AssignmentMethodField is a CHAR field
type AssignmentMethodField struct{ quickfix.FIXString }

//Tag returns tag.AssignmentMethod (744)
func (f AssignmentMethodField) Tag() quickfix.Tag { return tag.AssignmentMethod }

//NewAssignmentMethod returns a new AssignmentMethodField initialized with val
func NewAssignmentMethod(val quickfix.FIXString) AssignmentMethodField {
	return AssignmentMethodField{val}
}

//AssignmentUnitField is a QTY field
type AssignmentUnitField struct{ quickfix.FIXFloat }

//Tag returns tag.AssignmentUnit (745)
func (f AssignmentUnitField) Tag() quickfix.Tag { return tag.AssignmentUnit }

//NewAssignmentUnit returns a new AssignmentUnitField initialized with val
func NewAssignmentUnit(val quickfix.FIXFloat) AssignmentUnitField {
	return AssignmentUnitField{val}
}

//AttachmentPointField is a PERCENTAGE field
type AttachmentPointField struct{ quickfix.FIXFloat }

//Tag returns tag.AttachmentPoint (1457)
func (f AttachmentPointField) Tag() quickfix.Tag { return tag.AttachmentPoint }

//NewAttachmentPoint returns a new AttachmentPointField initialized with val
func NewAttachmentPoint(val quickfix.FIXFloat) AttachmentPointField {
	return AttachmentPointField{val}
}

//AutoAcceptIndicatorField is a BOOLEAN field
type AutoAcceptIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.AutoAcceptIndicator (754)
func (f AutoAcceptIndicatorField) Tag() quickfix.Tag { return tag.AutoAcceptIndicator }

//NewAutoAcceptIndicator returns a new AutoAcceptIndicatorField initialized with val
func NewAutoAcceptIndicator(val quickfix.FIXBoolean) AutoAcceptIndicatorField {
	return AutoAcceptIndicatorField{val}
}

//AvgParPxField is a PRICE field
type AvgParPxField struct{ quickfix.FIXFloat }

//Tag returns tag.AvgParPx (860)
func (f AvgParPxField) Tag() quickfix.Tag { return tag.AvgParPx }

//NewAvgParPx returns a new AvgParPxField initialized with val
func NewAvgParPx(val quickfix.FIXFloat) AvgParPxField {
	return AvgParPxField{val}
}

//AvgPrxPrecisionField is a INT field
type AvgPrxPrecisionField struct{ quickfix.FIXInt }

//Tag returns tag.AvgPrxPrecision (74)
func (f AvgPrxPrecisionField) Tag() quickfix.Tag { return tag.AvgPrxPrecision }

//NewAvgPrxPrecision returns a new AvgPrxPrecisionField initialized with val
func NewAvgPrxPrecision(val quickfix.FIXInt) AvgPrxPrecisionField {
	return AvgPrxPrecisionField{val}
}

//AvgPxField is a PRICE field
type AvgPxField struct{ quickfix.FIXFloat }

//Tag returns tag.AvgPx (6)
func (f AvgPxField) Tag() quickfix.Tag { return tag.AvgPx }

//NewAvgPx returns a new AvgPxField initialized with val
func NewAvgPx(val quickfix.FIXFloat) AvgPxField {
	return AvgPxField{val}
}

//AvgPxIndicatorField is a INT field
type AvgPxIndicatorField struct{ quickfix.FIXInt }

//Tag returns tag.AvgPxIndicator (819)
func (f AvgPxIndicatorField) Tag() quickfix.Tag { return tag.AvgPxIndicator }

//NewAvgPxIndicator returns a new AvgPxIndicatorField initialized with val
func NewAvgPxIndicator(val quickfix.FIXInt) AvgPxIndicatorField {
	return AvgPxIndicatorField{val}
}

//AvgPxPrecisionField is a INT field
type AvgPxPrecisionField struct{ quickfix.FIXInt }

//Tag returns tag.AvgPxPrecision (74)
func (f AvgPxPrecisionField) Tag() quickfix.Tag { return tag.AvgPxPrecision }

//NewAvgPxPrecision returns a new AvgPxPrecisionField initialized with val
func NewAvgPxPrecision(val quickfix.FIXInt) AvgPxPrecisionField {
	return AvgPxPrecisionField{val}
}

//BasisFeatureDateField is a LOCALMKTDATE field
type BasisFeatureDateField struct{ quickfix.FIXString }

//Tag returns tag.BasisFeatureDate (259)
func (f BasisFeatureDateField) Tag() quickfix.Tag { return tag.BasisFeatureDate }

//NewBasisFeatureDate returns a new BasisFeatureDateField initialized with val
func NewBasisFeatureDate(val quickfix.FIXString) BasisFeatureDateField {
	return BasisFeatureDateField{val}
}

//BasisFeaturePriceField is a PRICE field
type BasisFeaturePriceField struct{ quickfix.FIXFloat }

//Tag returns tag.BasisFeaturePrice (260)
func (f BasisFeaturePriceField) Tag() quickfix.Tag { return tag.BasisFeaturePrice }

//NewBasisFeaturePrice returns a new BasisFeaturePriceField initialized with val
func NewBasisFeaturePrice(val quickfix.FIXFloat) BasisFeaturePriceField {
	return BasisFeaturePriceField{val}
}

//BasisPxTypeField is a CHAR field
type BasisPxTypeField struct{ quickfix.FIXString }

//Tag returns tag.BasisPxType (419)
func (f BasisPxTypeField) Tag() quickfix.Tag { return tag.BasisPxType }

//NewBasisPxType returns a new BasisPxTypeField initialized with val
func NewBasisPxType(val quickfix.FIXString) BasisPxTypeField {
	return BasisPxTypeField{val}
}

//BeginSeqNoField is a SEQNUM field
type BeginSeqNoField struct{ quickfix.FIXInt }

//Tag returns tag.BeginSeqNo (7)
func (f BeginSeqNoField) Tag() quickfix.Tag { return tag.BeginSeqNo }

//NewBeginSeqNo returns a new BeginSeqNoField initialized with val
func NewBeginSeqNo(val quickfix.FIXInt) BeginSeqNoField {
	return BeginSeqNoField{val}
}

//BeginStringField is a STRING field
type BeginStringField struct{ quickfix.FIXString }

//Tag returns tag.BeginString (8)
func (f BeginStringField) Tag() quickfix.Tag { return tag.BeginString }

//NewBeginString returns a new BeginStringField initialized with val
func NewBeginString(val quickfix.FIXString) BeginStringField {
	return BeginStringField{val}
}

//BenchmarkField is a CHAR field
type BenchmarkField struct{ quickfix.FIXString }

//Tag returns tag.Benchmark (219)
func (f BenchmarkField) Tag() quickfix.Tag { return tag.Benchmark }

//NewBenchmark returns a new BenchmarkField initialized with val
func NewBenchmark(val quickfix.FIXString) BenchmarkField {
	return BenchmarkField{val}
}

//BenchmarkCurveCurrencyField is a CURRENCY field
type BenchmarkCurveCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.BenchmarkCurveCurrency (220)
func (f BenchmarkCurveCurrencyField) Tag() quickfix.Tag { return tag.BenchmarkCurveCurrency }

//NewBenchmarkCurveCurrency returns a new BenchmarkCurveCurrencyField initialized with val
func NewBenchmarkCurveCurrency(val quickfix.FIXString) BenchmarkCurveCurrencyField {
	return BenchmarkCurveCurrencyField{val}
}

//BenchmarkCurveNameField is a STRING field
type BenchmarkCurveNameField struct{ quickfix.FIXString }

//Tag returns tag.BenchmarkCurveName (221)
func (f BenchmarkCurveNameField) Tag() quickfix.Tag { return tag.BenchmarkCurveName }

//NewBenchmarkCurveName returns a new BenchmarkCurveNameField initialized with val
func NewBenchmarkCurveName(val quickfix.FIXString) BenchmarkCurveNameField {
	return BenchmarkCurveNameField{val}
}

//BenchmarkCurvePointField is a STRING field
type BenchmarkCurvePointField struct{ quickfix.FIXString }

//Tag returns tag.BenchmarkCurvePoint (222)
func (f BenchmarkCurvePointField) Tag() quickfix.Tag { return tag.BenchmarkCurvePoint }

//NewBenchmarkCurvePoint returns a new BenchmarkCurvePointField initialized with val
func NewBenchmarkCurvePoint(val quickfix.FIXString) BenchmarkCurvePointField {
	return BenchmarkCurvePointField{val}
}

//BenchmarkPriceField is a PRICE field
type BenchmarkPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.BenchmarkPrice (662)
func (f BenchmarkPriceField) Tag() quickfix.Tag { return tag.BenchmarkPrice }

//NewBenchmarkPrice returns a new BenchmarkPriceField initialized with val
func NewBenchmarkPrice(val quickfix.FIXFloat) BenchmarkPriceField {
	return BenchmarkPriceField{val}
}

//BenchmarkPriceTypeField is a INT field
type BenchmarkPriceTypeField struct{ quickfix.FIXInt }

//Tag returns tag.BenchmarkPriceType (663)
func (f BenchmarkPriceTypeField) Tag() quickfix.Tag { return tag.BenchmarkPriceType }

//NewBenchmarkPriceType returns a new BenchmarkPriceTypeField initialized with val
func NewBenchmarkPriceType(val quickfix.FIXInt) BenchmarkPriceTypeField {
	return BenchmarkPriceTypeField{val}
}

//BenchmarkSecurityIDField is a STRING field
type BenchmarkSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.BenchmarkSecurityID (699)
func (f BenchmarkSecurityIDField) Tag() quickfix.Tag { return tag.BenchmarkSecurityID }

//NewBenchmarkSecurityID returns a new BenchmarkSecurityIDField initialized with val
func NewBenchmarkSecurityID(val quickfix.FIXString) BenchmarkSecurityIDField {
	return BenchmarkSecurityIDField{val}
}

//BenchmarkSecurityIDSourceField is a STRING field
type BenchmarkSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.BenchmarkSecurityIDSource (761)
func (f BenchmarkSecurityIDSourceField) Tag() quickfix.Tag { return tag.BenchmarkSecurityIDSource }

//NewBenchmarkSecurityIDSource returns a new BenchmarkSecurityIDSourceField initialized with val
func NewBenchmarkSecurityIDSource(val quickfix.FIXString) BenchmarkSecurityIDSourceField {
	return BenchmarkSecurityIDSourceField{val}
}

//BidDescriptorField is a STRING field
type BidDescriptorField struct{ quickfix.FIXString }

//Tag returns tag.BidDescriptor (400)
func (f BidDescriptorField) Tag() quickfix.Tag { return tag.BidDescriptor }

//NewBidDescriptor returns a new BidDescriptorField initialized with val
func NewBidDescriptor(val quickfix.FIXString) BidDescriptorField {
	return BidDescriptorField{val}
}

//BidDescriptorTypeField is a INT field
type BidDescriptorTypeField struct{ quickfix.FIXInt }

//Tag returns tag.BidDescriptorType (399)
func (f BidDescriptorTypeField) Tag() quickfix.Tag { return tag.BidDescriptorType }

//NewBidDescriptorType returns a new BidDescriptorTypeField initialized with val
func NewBidDescriptorType(val quickfix.FIXInt) BidDescriptorTypeField {
	return BidDescriptorTypeField{val}
}

//BidForwardPointsField is a PRICEOFFSET field
type BidForwardPointsField struct{ quickfix.FIXFloat }

//Tag returns tag.BidForwardPoints (189)
func (f BidForwardPointsField) Tag() quickfix.Tag { return tag.BidForwardPoints }

//NewBidForwardPoints returns a new BidForwardPointsField initialized with val
func NewBidForwardPoints(val quickfix.FIXFloat) BidForwardPointsField {
	return BidForwardPointsField{val}
}

//BidForwardPoints2Field is a PRICEOFFSET field
type BidForwardPoints2Field struct{ quickfix.FIXFloat }

//Tag returns tag.BidForwardPoints2 (642)
func (f BidForwardPoints2Field) Tag() quickfix.Tag { return tag.BidForwardPoints2 }

//NewBidForwardPoints2 returns a new BidForwardPoints2Field initialized with val
func NewBidForwardPoints2(val quickfix.FIXFloat) BidForwardPoints2Field {
	return BidForwardPoints2Field{val}
}

//BidIDField is a STRING field
type BidIDField struct{ quickfix.FIXString }

//Tag returns tag.BidID (390)
func (f BidIDField) Tag() quickfix.Tag { return tag.BidID }

//NewBidID returns a new BidIDField initialized with val
func NewBidID(val quickfix.FIXString) BidIDField {
	return BidIDField{val}
}

//BidPxField is a PRICE field
type BidPxField struct{ quickfix.FIXFloat }

//Tag returns tag.BidPx (132)
func (f BidPxField) Tag() quickfix.Tag { return tag.BidPx }

//NewBidPx returns a new BidPxField initialized with val
func NewBidPx(val quickfix.FIXFloat) BidPxField {
	return BidPxField{val}
}

//BidRequestTransTypeField is a CHAR field
type BidRequestTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.BidRequestTransType (374)
func (f BidRequestTransTypeField) Tag() quickfix.Tag { return tag.BidRequestTransType }

//NewBidRequestTransType returns a new BidRequestTransTypeField initialized with val
func NewBidRequestTransType(val quickfix.FIXString) BidRequestTransTypeField {
	return BidRequestTransTypeField{val}
}

//BidSizeField is a QTY field
type BidSizeField struct{ quickfix.FIXFloat }

//Tag returns tag.BidSize (134)
func (f BidSizeField) Tag() quickfix.Tag { return tag.BidSize }

//NewBidSize returns a new BidSizeField initialized with val
func NewBidSize(val quickfix.FIXFloat) BidSizeField {
	return BidSizeField{val}
}

//BidSpotRateField is a PRICE field
type BidSpotRateField struct{ quickfix.FIXFloat }

//Tag returns tag.BidSpotRate (188)
func (f BidSpotRateField) Tag() quickfix.Tag { return tag.BidSpotRate }

//NewBidSpotRate returns a new BidSpotRateField initialized with val
func NewBidSpotRate(val quickfix.FIXFloat) BidSpotRateField {
	return BidSpotRateField{val}
}

//BidSwapPointsField is a PRICEOFFSET field
type BidSwapPointsField struct{ quickfix.FIXFloat }

//Tag returns tag.BidSwapPoints (1065)
func (f BidSwapPointsField) Tag() quickfix.Tag { return tag.BidSwapPoints }

//NewBidSwapPoints returns a new BidSwapPointsField initialized with val
func NewBidSwapPoints(val quickfix.FIXFloat) BidSwapPointsField {
	return BidSwapPointsField{val}
}

//BidTradeTypeField is a CHAR field
type BidTradeTypeField struct{ quickfix.FIXString }

//Tag returns tag.BidTradeType (418)
func (f BidTradeTypeField) Tag() quickfix.Tag { return tag.BidTradeType }

//NewBidTradeType returns a new BidTradeTypeField initialized with val
func NewBidTradeType(val quickfix.FIXString) BidTradeTypeField {
	return BidTradeTypeField{val}
}

//BidTypeField is a INT field
type BidTypeField struct{ quickfix.FIXInt }

//Tag returns tag.BidType (394)
func (f BidTypeField) Tag() quickfix.Tag { return tag.BidType }

//NewBidType returns a new BidTypeField initialized with val
func NewBidType(val quickfix.FIXInt) BidTypeField {
	return BidTypeField{val}
}

//BidYieldField is a PERCENTAGE field
type BidYieldField struct{ quickfix.FIXFloat }

//Tag returns tag.BidYield (632)
func (f BidYieldField) Tag() quickfix.Tag { return tag.BidYield }

//NewBidYield returns a new BidYieldField initialized with val
func NewBidYield(val quickfix.FIXFloat) BidYieldField {
	return BidYieldField{val}
}

//BodyLengthField is a LENGTH field
type BodyLengthField struct{ quickfix.FIXInt }

//Tag returns tag.BodyLength (9)
func (f BodyLengthField) Tag() quickfix.Tag { return tag.BodyLength }

//NewBodyLength returns a new BodyLengthField initialized with val
func NewBodyLength(val quickfix.FIXInt) BodyLengthField {
	return BodyLengthField{val}
}

//BookingRefIDField is a STRING field
type BookingRefIDField struct{ quickfix.FIXString }

//Tag returns tag.BookingRefID (466)
func (f BookingRefIDField) Tag() quickfix.Tag { return tag.BookingRefID }

//NewBookingRefID returns a new BookingRefIDField initialized with val
func NewBookingRefID(val quickfix.FIXString) BookingRefIDField {
	return BookingRefIDField{val}
}

//BookingTypeField is a INT field
type BookingTypeField struct{ quickfix.FIXInt }

//Tag returns tag.BookingType (775)
func (f BookingTypeField) Tag() quickfix.Tag { return tag.BookingType }

//NewBookingType returns a new BookingTypeField initialized with val
func NewBookingType(val quickfix.FIXInt) BookingTypeField {
	return BookingTypeField{val}
}

//BookingUnitField is a CHAR field
type BookingUnitField struct{ quickfix.FIXString }

//Tag returns tag.BookingUnit (590)
func (f BookingUnitField) Tag() quickfix.Tag { return tag.BookingUnit }

//NewBookingUnit returns a new BookingUnitField initialized with val
func NewBookingUnit(val quickfix.FIXString) BookingUnitField {
	return BookingUnitField{val}
}

//BrokerOfCreditField is a STRING field
type BrokerOfCreditField struct{ quickfix.FIXString }

//Tag returns tag.BrokerOfCredit (92)
func (f BrokerOfCreditField) Tag() quickfix.Tag { return tag.BrokerOfCredit }

//NewBrokerOfCredit returns a new BrokerOfCreditField initialized with val
func NewBrokerOfCredit(val quickfix.FIXString) BrokerOfCreditField {
	return BrokerOfCreditField{val}
}

//BusinessRejectReasonField is a INT field
type BusinessRejectReasonField struct{ quickfix.FIXInt }

//Tag returns tag.BusinessRejectReason (380)
func (f BusinessRejectReasonField) Tag() quickfix.Tag { return tag.BusinessRejectReason }

//NewBusinessRejectReason returns a new BusinessRejectReasonField initialized with val
func NewBusinessRejectReason(val quickfix.FIXInt) BusinessRejectReasonField {
	return BusinessRejectReasonField{val}
}

//BusinessRejectRefIDField is a STRING field
type BusinessRejectRefIDField struct{ quickfix.FIXString }

//Tag returns tag.BusinessRejectRefID (379)
func (f BusinessRejectRefIDField) Tag() quickfix.Tag { return tag.BusinessRejectRefID }

//NewBusinessRejectRefID returns a new BusinessRejectRefIDField initialized with val
func NewBusinessRejectRefID(val quickfix.FIXString) BusinessRejectRefIDField {
	return BusinessRejectRefIDField{val}
}

//BuyVolumeField is a QTY field
type BuyVolumeField struct{ quickfix.FIXFloat }

//Tag returns tag.BuyVolume (330)
func (f BuyVolumeField) Tag() quickfix.Tag { return tag.BuyVolume }

//NewBuyVolume returns a new BuyVolumeField initialized with val
func NewBuyVolume(val quickfix.FIXFloat) BuyVolumeField {
	return BuyVolumeField{val}
}

//CFICodeField is a STRING field
type CFICodeField struct{ quickfix.FIXString }

//Tag returns tag.CFICode (461)
func (f CFICodeField) Tag() quickfix.Tag { return tag.CFICode }

//NewCFICode returns a new CFICodeField initialized with val
func NewCFICode(val quickfix.FIXString) CFICodeField {
	return CFICodeField{val}
}

//CPProgramField is a INT field
type CPProgramField struct{ quickfix.FIXInt }

//Tag returns tag.CPProgram (875)
func (f CPProgramField) Tag() quickfix.Tag { return tag.CPProgram }

//NewCPProgram returns a new CPProgramField initialized with val
func NewCPProgram(val quickfix.FIXInt) CPProgramField {
	return CPProgramField{val}
}

//CPRegTypeField is a STRING field
type CPRegTypeField struct{ quickfix.FIXString }

//Tag returns tag.CPRegType (876)
func (f CPRegTypeField) Tag() quickfix.Tag { return tag.CPRegType }

//NewCPRegType returns a new CPRegTypeField initialized with val
func NewCPRegType(val quickfix.FIXString) CPRegTypeField {
	return CPRegTypeField{val}
}

//CalculatedCcyLastQtyField is a QTY field
type CalculatedCcyLastQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.CalculatedCcyLastQty (1056)
func (f CalculatedCcyLastQtyField) Tag() quickfix.Tag { return tag.CalculatedCcyLastQty }

//NewCalculatedCcyLastQty returns a new CalculatedCcyLastQtyField initialized with val
func NewCalculatedCcyLastQty(val quickfix.FIXFloat) CalculatedCcyLastQtyField {
	return CalculatedCcyLastQtyField{val}
}

//CancellationRightsField is a CHAR field
type CancellationRightsField struct{ quickfix.FIXString }

//Tag returns tag.CancellationRights (480)
func (f CancellationRightsField) Tag() quickfix.Tag { return tag.CancellationRights }

//NewCancellationRights returns a new CancellationRightsField initialized with val
func NewCancellationRights(val quickfix.FIXString) CancellationRightsField {
	return CancellationRightsField{val}
}

//CapPriceField is a PRICE field
type CapPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.CapPrice (1199)
func (f CapPriceField) Tag() quickfix.Tag { return tag.CapPrice }

//NewCapPrice returns a new CapPriceField initialized with val
func NewCapPrice(val quickfix.FIXFloat) CapPriceField {
	return CapPriceField{val}
}

//CardExpDateField is a LOCALMKTDATE field
type CardExpDateField struct{ quickfix.FIXString }

//Tag returns tag.CardExpDate (490)
func (f CardExpDateField) Tag() quickfix.Tag { return tag.CardExpDate }

//NewCardExpDate returns a new CardExpDateField initialized with val
func NewCardExpDate(val quickfix.FIXString) CardExpDateField {
	return CardExpDateField{val}
}

//CardHolderNameField is a STRING field
type CardHolderNameField struct{ quickfix.FIXString }

//Tag returns tag.CardHolderName (488)
func (f CardHolderNameField) Tag() quickfix.Tag { return tag.CardHolderName }

//NewCardHolderName returns a new CardHolderNameField initialized with val
func NewCardHolderName(val quickfix.FIXString) CardHolderNameField {
	return CardHolderNameField{val}
}

//CardIssNoField is a STRING field
type CardIssNoField struct{ quickfix.FIXString }

//Tag returns tag.CardIssNo (491)
func (f CardIssNoField) Tag() quickfix.Tag { return tag.CardIssNo }

//NewCardIssNo returns a new CardIssNoField initialized with val
func NewCardIssNo(val quickfix.FIXString) CardIssNoField {
	return CardIssNoField{val}
}

//CardIssNumField is a STRING field
type CardIssNumField struct{ quickfix.FIXString }

//Tag returns tag.CardIssNum (491)
func (f CardIssNumField) Tag() quickfix.Tag { return tag.CardIssNum }

//NewCardIssNum returns a new CardIssNumField initialized with val
func NewCardIssNum(val quickfix.FIXString) CardIssNumField {
	return CardIssNumField{val}
}

//CardNumberField is a STRING field
type CardNumberField struct{ quickfix.FIXString }

//Tag returns tag.CardNumber (489)
func (f CardNumberField) Tag() quickfix.Tag { return tag.CardNumber }

//NewCardNumber returns a new CardNumberField initialized with val
func NewCardNumber(val quickfix.FIXString) CardNumberField {
	return CardNumberField{val}
}

//CardStartDateField is a LOCALMKTDATE field
type CardStartDateField struct{ quickfix.FIXString }

//Tag returns tag.CardStartDate (503)
func (f CardStartDateField) Tag() quickfix.Tag { return tag.CardStartDate }

//NewCardStartDate returns a new CardStartDateField initialized with val
func NewCardStartDate(val quickfix.FIXString) CardStartDateField {
	return CardStartDateField{val}
}

//CashDistribAgentAcctNameField is a STRING field
type CashDistribAgentAcctNameField struct{ quickfix.FIXString }

//Tag returns tag.CashDistribAgentAcctName (502)
func (f CashDistribAgentAcctNameField) Tag() quickfix.Tag { return tag.CashDistribAgentAcctName }

//NewCashDistribAgentAcctName returns a new CashDistribAgentAcctNameField initialized with val
func NewCashDistribAgentAcctName(val quickfix.FIXString) CashDistribAgentAcctNameField {
	return CashDistribAgentAcctNameField{val}
}

//CashDistribAgentAcctNumberField is a STRING field
type CashDistribAgentAcctNumberField struct{ quickfix.FIXString }

//Tag returns tag.CashDistribAgentAcctNumber (500)
func (f CashDistribAgentAcctNumberField) Tag() quickfix.Tag { return tag.CashDistribAgentAcctNumber }

//NewCashDistribAgentAcctNumber returns a new CashDistribAgentAcctNumberField initialized with val
func NewCashDistribAgentAcctNumber(val quickfix.FIXString) CashDistribAgentAcctNumberField {
	return CashDistribAgentAcctNumberField{val}
}

//CashDistribAgentCodeField is a STRING field
type CashDistribAgentCodeField struct{ quickfix.FIXString }

//Tag returns tag.CashDistribAgentCode (499)
func (f CashDistribAgentCodeField) Tag() quickfix.Tag { return tag.CashDistribAgentCode }

//NewCashDistribAgentCode returns a new CashDistribAgentCodeField initialized with val
func NewCashDistribAgentCode(val quickfix.FIXString) CashDistribAgentCodeField {
	return CashDistribAgentCodeField{val}
}

//CashDistribAgentNameField is a STRING field
type CashDistribAgentNameField struct{ quickfix.FIXString }

//Tag returns tag.CashDistribAgentName (498)
func (f CashDistribAgentNameField) Tag() quickfix.Tag { return tag.CashDistribAgentName }

//NewCashDistribAgentName returns a new CashDistribAgentNameField initialized with val
func NewCashDistribAgentName(val quickfix.FIXString) CashDistribAgentNameField {
	return CashDistribAgentNameField{val}
}

//CashDistribCurrField is a CURRENCY field
type CashDistribCurrField struct{ quickfix.FIXString }

//Tag returns tag.CashDistribCurr (478)
func (f CashDistribCurrField) Tag() quickfix.Tag { return tag.CashDistribCurr }

//NewCashDistribCurr returns a new CashDistribCurrField initialized with val
func NewCashDistribCurr(val quickfix.FIXString) CashDistribCurrField {
	return CashDistribCurrField{val}
}

//CashDistribPayRefField is a STRING field
type CashDistribPayRefField struct{ quickfix.FIXString }

//Tag returns tag.CashDistribPayRef (501)
func (f CashDistribPayRefField) Tag() quickfix.Tag { return tag.CashDistribPayRef }

//NewCashDistribPayRef returns a new CashDistribPayRefField initialized with val
func NewCashDistribPayRef(val quickfix.FIXString) CashDistribPayRefField {
	return CashDistribPayRefField{val}
}

//CashMarginField is a CHAR field
type CashMarginField struct{ quickfix.FIXString }

//Tag returns tag.CashMargin (544)
func (f CashMarginField) Tag() quickfix.Tag { return tag.CashMargin }

//NewCashMargin returns a new CashMarginField initialized with val
func NewCashMargin(val quickfix.FIXString) CashMarginField {
	return CashMarginField{val}
}

//CashOrderQtyField is a QTY field
type CashOrderQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.CashOrderQty (152)
func (f CashOrderQtyField) Tag() quickfix.Tag { return tag.CashOrderQty }

//NewCashOrderQty returns a new CashOrderQtyField initialized with val
func NewCashOrderQty(val quickfix.FIXFloat) CashOrderQtyField {
	return CashOrderQtyField{val}
}

//CashOutstandingField is a AMT field
type CashOutstandingField struct{ quickfix.FIXFloat }

//Tag returns tag.CashOutstanding (901)
func (f CashOutstandingField) Tag() quickfix.Tag { return tag.CashOutstanding }

//NewCashOutstanding returns a new CashOutstandingField initialized with val
func NewCashOutstanding(val quickfix.FIXFloat) CashOutstandingField {
	return CashOutstandingField{val}
}

//CashSettlAgentAcctNameField is a STRING field
type CashSettlAgentAcctNameField struct{ quickfix.FIXString }

//Tag returns tag.CashSettlAgentAcctName (185)
func (f CashSettlAgentAcctNameField) Tag() quickfix.Tag { return tag.CashSettlAgentAcctName }

//NewCashSettlAgentAcctName returns a new CashSettlAgentAcctNameField initialized with val
func NewCashSettlAgentAcctName(val quickfix.FIXString) CashSettlAgentAcctNameField {
	return CashSettlAgentAcctNameField{val}
}

//CashSettlAgentAcctNumField is a STRING field
type CashSettlAgentAcctNumField struct{ quickfix.FIXString }

//Tag returns tag.CashSettlAgentAcctNum (184)
func (f CashSettlAgentAcctNumField) Tag() quickfix.Tag { return tag.CashSettlAgentAcctNum }

//NewCashSettlAgentAcctNum returns a new CashSettlAgentAcctNumField initialized with val
func NewCashSettlAgentAcctNum(val quickfix.FIXString) CashSettlAgentAcctNumField {
	return CashSettlAgentAcctNumField{val}
}

//CashSettlAgentCodeField is a STRING field
type CashSettlAgentCodeField struct{ quickfix.FIXString }

//Tag returns tag.CashSettlAgentCode (183)
func (f CashSettlAgentCodeField) Tag() quickfix.Tag { return tag.CashSettlAgentCode }

//NewCashSettlAgentCode returns a new CashSettlAgentCodeField initialized with val
func NewCashSettlAgentCode(val quickfix.FIXString) CashSettlAgentCodeField {
	return CashSettlAgentCodeField{val}
}

//CashSettlAgentContactNameField is a STRING field
type CashSettlAgentContactNameField struct{ quickfix.FIXString }

//Tag returns tag.CashSettlAgentContactName (186)
func (f CashSettlAgentContactNameField) Tag() quickfix.Tag { return tag.CashSettlAgentContactName }

//NewCashSettlAgentContactName returns a new CashSettlAgentContactNameField initialized with val
func NewCashSettlAgentContactName(val quickfix.FIXString) CashSettlAgentContactNameField {
	return CashSettlAgentContactNameField{val}
}

//CashSettlAgentContactPhoneField is a STRING field
type CashSettlAgentContactPhoneField struct{ quickfix.FIXString }

//Tag returns tag.CashSettlAgentContactPhone (187)
func (f CashSettlAgentContactPhoneField) Tag() quickfix.Tag { return tag.CashSettlAgentContactPhone }

//NewCashSettlAgentContactPhone returns a new CashSettlAgentContactPhoneField initialized with val
func NewCashSettlAgentContactPhone(val quickfix.FIXString) CashSettlAgentContactPhoneField {
	return CashSettlAgentContactPhoneField{val}
}

//CashSettlAgentNameField is a STRING field
type CashSettlAgentNameField struct{ quickfix.FIXString }

//Tag returns tag.CashSettlAgentName (182)
func (f CashSettlAgentNameField) Tag() quickfix.Tag { return tag.CashSettlAgentName }

//NewCashSettlAgentName returns a new CashSettlAgentNameField initialized with val
func NewCashSettlAgentName(val quickfix.FIXString) CashSettlAgentNameField {
	return CashSettlAgentNameField{val}
}

//CcyAmtField is a AMT field
type CcyAmtField struct{ quickfix.FIXFloat }

//Tag returns tag.CcyAmt (1157)
func (f CcyAmtField) Tag() quickfix.Tag { return tag.CcyAmt }

//NewCcyAmt returns a new CcyAmtField initialized with val
func NewCcyAmt(val quickfix.FIXFloat) CcyAmtField {
	return CcyAmtField{val}
}

//CheckSumField is a STRING field
type CheckSumField struct{ quickfix.FIXString }

//Tag returns tag.CheckSum (10)
func (f CheckSumField) Tag() quickfix.Tag { return tag.CheckSum }

//NewCheckSum returns a new CheckSumField initialized with val
func NewCheckSum(val quickfix.FIXString) CheckSumField {
	return CheckSumField{val}
}

//ClOrdIDField is a STRING field
type ClOrdIDField struct{ quickfix.FIXString }

//Tag returns tag.ClOrdID (11)
func (f ClOrdIDField) Tag() quickfix.Tag { return tag.ClOrdID }

//NewClOrdID returns a new ClOrdIDField initialized with val
func NewClOrdID(val quickfix.FIXString) ClOrdIDField {
	return ClOrdIDField{val}
}

//ClOrdLinkIDField is a STRING field
type ClOrdLinkIDField struct{ quickfix.FIXString }

//Tag returns tag.ClOrdLinkID (583)
func (f ClOrdLinkIDField) Tag() quickfix.Tag { return tag.ClOrdLinkID }

//NewClOrdLinkID returns a new ClOrdLinkIDField initialized with val
func NewClOrdLinkID(val quickfix.FIXString) ClOrdLinkIDField {
	return ClOrdLinkIDField{val}
}

//ClearingAccountField is a STRING field
type ClearingAccountField struct{ quickfix.FIXString }

//Tag returns tag.ClearingAccount (440)
func (f ClearingAccountField) Tag() quickfix.Tag { return tag.ClearingAccount }

//NewClearingAccount returns a new ClearingAccountField initialized with val
func NewClearingAccount(val quickfix.FIXString) ClearingAccountField {
	return ClearingAccountField{val}
}

//ClearingBusinessDateField is a LOCALMKTDATE field
type ClearingBusinessDateField struct{ quickfix.FIXString }

//Tag returns tag.ClearingBusinessDate (715)
func (f ClearingBusinessDateField) Tag() quickfix.Tag { return tag.ClearingBusinessDate }

//NewClearingBusinessDate returns a new ClearingBusinessDateField initialized with val
func NewClearingBusinessDate(val quickfix.FIXString) ClearingBusinessDateField {
	return ClearingBusinessDateField{val}
}

//ClearingFeeIndicatorField is a STRING field
type ClearingFeeIndicatorField struct{ quickfix.FIXString }

//Tag returns tag.ClearingFeeIndicator (635)
func (f ClearingFeeIndicatorField) Tag() quickfix.Tag { return tag.ClearingFeeIndicator }

//NewClearingFeeIndicator returns a new ClearingFeeIndicatorField initialized with val
func NewClearingFeeIndicator(val quickfix.FIXString) ClearingFeeIndicatorField {
	return ClearingFeeIndicatorField{val}
}

//ClearingFirmField is a STRING field
type ClearingFirmField struct{ quickfix.FIXString }

//Tag returns tag.ClearingFirm (439)
func (f ClearingFirmField) Tag() quickfix.Tag { return tag.ClearingFirm }

//NewClearingFirm returns a new ClearingFirmField initialized with val
func NewClearingFirm(val quickfix.FIXString) ClearingFirmField {
	return ClearingFirmField{val}
}

//ClearingInstructionField is a INT field
type ClearingInstructionField struct{ quickfix.FIXInt }

//Tag returns tag.ClearingInstruction (577)
func (f ClearingInstructionField) Tag() quickfix.Tag { return tag.ClearingInstruction }

//NewClearingInstruction returns a new ClearingInstructionField initialized with val
func NewClearingInstruction(val quickfix.FIXInt) ClearingInstructionField {
	return ClearingInstructionField{val}
}

//ClientBidIDField is a STRING field
type ClientBidIDField struct{ quickfix.FIXString }

//Tag returns tag.ClientBidID (391)
func (f ClientBidIDField) Tag() quickfix.Tag { return tag.ClientBidID }

//NewClientBidID returns a new ClientBidIDField initialized with val
func NewClientBidID(val quickfix.FIXString) ClientBidIDField {
	return ClientBidIDField{val}
}

//ClientIDField is a STRING field
type ClientIDField struct{ quickfix.FIXString }

//Tag returns tag.ClientID (109)
func (f ClientIDField) Tag() quickfix.Tag { return tag.ClientID }

//NewClientID returns a new ClientIDField initialized with val
func NewClientID(val quickfix.FIXString) ClientIDField {
	return ClientIDField{val}
}

//CollActionField is a INT field
type CollActionField struct{ quickfix.FIXInt }

//Tag returns tag.CollAction (944)
func (f CollActionField) Tag() quickfix.Tag { return tag.CollAction }

//NewCollAction returns a new CollActionField initialized with val
func NewCollAction(val quickfix.FIXInt) CollActionField {
	return CollActionField{val}
}

//CollApplTypeField is a INT field
type CollApplTypeField struct{ quickfix.FIXInt }

//Tag returns tag.CollApplType (1043)
func (f CollApplTypeField) Tag() quickfix.Tag { return tag.CollApplType }

//NewCollApplType returns a new CollApplTypeField initialized with val
func NewCollApplType(val quickfix.FIXInt) CollApplTypeField {
	return CollApplTypeField{val}
}

//CollAsgnIDField is a STRING field
type CollAsgnIDField struct{ quickfix.FIXString }

//Tag returns tag.CollAsgnID (902)
func (f CollAsgnIDField) Tag() quickfix.Tag { return tag.CollAsgnID }

//NewCollAsgnID returns a new CollAsgnIDField initialized with val
func NewCollAsgnID(val quickfix.FIXString) CollAsgnIDField {
	return CollAsgnIDField{val}
}

//CollAsgnReasonField is a INT field
type CollAsgnReasonField struct{ quickfix.FIXInt }

//Tag returns tag.CollAsgnReason (895)
func (f CollAsgnReasonField) Tag() quickfix.Tag { return tag.CollAsgnReason }

//NewCollAsgnReason returns a new CollAsgnReasonField initialized with val
func NewCollAsgnReason(val quickfix.FIXInt) CollAsgnReasonField {
	return CollAsgnReasonField{val}
}

//CollAsgnRefIDField is a STRING field
type CollAsgnRefIDField struct{ quickfix.FIXString }

//Tag returns tag.CollAsgnRefID (907)
func (f CollAsgnRefIDField) Tag() quickfix.Tag { return tag.CollAsgnRefID }

//NewCollAsgnRefID returns a new CollAsgnRefIDField initialized with val
func NewCollAsgnRefID(val quickfix.FIXString) CollAsgnRefIDField {
	return CollAsgnRefIDField{val}
}

//CollAsgnRejectReasonField is a INT field
type CollAsgnRejectReasonField struct{ quickfix.FIXInt }

//Tag returns tag.CollAsgnRejectReason (906)
func (f CollAsgnRejectReasonField) Tag() quickfix.Tag { return tag.CollAsgnRejectReason }

//NewCollAsgnRejectReason returns a new CollAsgnRejectReasonField initialized with val
func NewCollAsgnRejectReason(val quickfix.FIXInt) CollAsgnRejectReasonField {
	return CollAsgnRejectReasonField{val}
}

//CollAsgnRespTypeField is a INT field
type CollAsgnRespTypeField struct{ quickfix.FIXInt }

//Tag returns tag.CollAsgnRespType (905)
func (f CollAsgnRespTypeField) Tag() quickfix.Tag { return tag.CollAsgnRespType }

//NewCollAsgnRespType returns a new CollAsgnRespTypeField initialized with val
func NewCollAsgnRespType(val quickfix.FIXInt) CollAsgnRespTypeField {
	return CollAsgnRespTypeField{val}
}

//CollAsgnTransTypeField is a INT field
type CollAsgnTransTypeField struct{ quickfix.FIXInt }

//Tag returns tag.CollAsgnTransType (903)
func (f CollAsgnTransTypeField) Tag() quickfix.Tag { return tag.CollAsgnTransType }

//NewCollAsgnTransType returns a new CollAsgnTransTypeField initialized with val
func NewCollAsgnTransType(val quickfix.FIXInt) CollAsgnTransTypeField {
	return CollAsgnTransTypeField{val}
}

//CollInquiryIDField is a STRING field
type CollInquiryIDField struct{ quickfix.FIXString }

//Tag returns tag.CollInquiryID (909)
func (f CollInquiryIDField) Tag() quickfix.Tag { return tag.CollInquiryID }

//NewCollInquiryID returns a new CollInquiryIDField initialized with val
func NewCollInquiryID(val quickfix.FIXString) CollInquiryIDField {
	return CollInquiryIDField{val}
}

//CollInquiryQualifierField is a INT field
type CollInquiryQualifierField struct{ quickfix.FIXInt }

//Tag returns tag.CollInquiryQualifier (896)
func (f CollInquiryQualifierField) Tag() quickfix.Tag { return tag.CollInquiryQualifier }

//NewCollInquiryQualifier returns a new CollInquiryQualifierField initialized with val
func NewCollInquiryQualifier(val quickfix.FIXInt) CollInquiryQualifierField {
	return CollInquiryQualifierField{val}
}

//CollInquiryResultField is a INT field
type CollInquiryResultField struct{ quickfix.FIXInt }

//Tag returns tag.CollInquiryResult (946)
func (f CollInquiryResultField) Tag() quickfix.Tag { return tag.CollInquiryResult }

//NewCollInquiryResult returns a new CollInquiryResultField initialized with val
func NewCollInquiryResult(val quickfix.FIXInt) CollInquiryResultField {
	return CollInquiryResultField{val}
}

//CollInquiryStatusField is a INT field
type CollInquiryStatusField struct{ quickfix.FIXInt }

//Tag returns tag.CollInquiryStatus (945)
func (f CollInquiryStatusField) Tag() quickfix.Tag { return tag.CollInquiryStatus }

//NewCollInquiryStatus returns a new CollInquiryStatusField initialized with val
func NewCollInquiryStatus(val quickfix.FIXInt) CollInquiryStatusField {
	return CollInquiryStatusField{val}
}

//CollReqIDField is a STRING field
type CollReqIDField struct{ quickfix.FIXString }

//Tag returns tag.CollReqID (894)
func (f CollReqIDField) Tag() quickfix.Tag { return tag.CollReqID }

//NewCollReqID returns a new CollReqIDField initialized with val
func NewCollReqID(val quickfix.FIXString) CollReqIDField {
	return CollReqIDField{val}
}

//CollRespIDField is a STRING field
type CollRespIDField struct{ quickfix.FIXString }

//Tag returns tag.CollRespID (904)
func (f CollRespIDField) Tag() quickfix.Tag { return tag.CollRespID }

//NewCollRespID returns a new CollRespIDField initialized with val
func NewCollRespID(val quickfix.FIXString) CollRespIDField {
	return CollRespIDField{val}
}

//CollRptIDField is a STRING field
type CollRptIDField struct{ quickfix.FIXString }

//Tag returns tag.CollRptID (908)
func (f CollRptIDField) Tag() quickfix.Tag { return tag.CollRptID }

//NewCollRptID returns a new CollRptIDField initialized with val
func NewCollRptID(val quickfix.FIXString) CollRptIDField {
	return CollRptIDField{val}
}

//CollStatusField is a INT field
type CollStatusField struct{ quickfix.FIXInt }

//Tag returns tag.CollStatus (910)
func (f CollStatusField) Tag() quickfix.Tag { return tag.CollStatus }

//NewCollStatus returns a new CollStatusField initialized with val
func NewCollStatus(val quickfix.FIXInt) CollStatusField {
	return CollStatusField{val}
}

//CommCurrencyField is a CURRENCY field
type CommCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.CommCurrency (479)
func (f CommCurrencyField) Tag() quickfix.Tag { return tag.CommCurrency }

//NewCommCurrency returns a new CommCurrencyField initialized with val
func NewCommCurrency(val quickfix.FIXString) CommCurrencyField {
	return CommCurrencyField{val}
}

//CommTypeField is a CHAR field
type CommTypeField struct{ quickfix.FIXString }

//Tag returns tag.CommType (13)
func (f CommTypeField) Tag() quickfix.Tag { return tag.CommType }

//NewCommType returns a new CommTypeField initialized with val
func NewCommType(val quickfix.FIXString) CommTypeField {
	return CommTypeField{val}
}

//CommissionField is a AMT field
type CommissionField struct{ quickfix.FIXFloat }

//Tag returns tag.Commission (12)
func (f CommissionField) Tag() quickfix.Tag { return tag.Commission }

//NewCommission returns a new CommissionField initialized with val
func NewCommission(val quickfix.FIXFloat) CommissionField {
	return CommissionField{val}
}

//ComplexEventConditionField is a INT field
type ComplexEventConditionField struct{ quickfix.FIXInt }

//Tag returns tag.ComplexEventCondition (1490)
func (f ComplexEventConditionField) Tag() quickfix.Tag { return tag.ComplexEventCondition }

//NewComplexEventCondition returns a new ComplexEventConditionField initialized with val
func NewComplexEventCondition(val quickfix.FIXInt) ComplexEventConditionField {
	return ComplexEventConditionField{val}
}

//ComplexEventEndDateField is a UTCTIMESTAMP field
type ComplexEventEndDateField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.ComplexEventEndDate (1493)
func (f ComplexEventEndDateField) Tag() quickfix.Tag { return tag.ComplexEventEndDate }

//NewComplexEventEndDate returns a new ComplexEventEndDateField initialized with val
func NewComplexEventEndDate(val time.Time) ComplexEventEndDateField {
	return ComplexEventEndDateField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewComplexEventEndDateNoMillis returns a new ComplexEventEndDateField initialized with val without millisecs
func NewComplexEventEndDateNoMillis(val time.Time) ComplexEventEndDateField {
	return ComplexEventEndDateField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//ComplexEventEndTimeField is a UTCTIMEONLY field
type ComplexEventEndTimeField struct{ quickfix.FIXString }

//Tag returns tag.ComplexEventEndTime (1496)
func (f ComplexEventEndTimeField) Tag() quickfix.Tag { return tag.ComplexEventEndTime }

//NewComplexEventEndTime returns a new ComplexEventEndTimeField initialized with val
func NewComplexEventEndTime(val quickfix.FIXString) ComplexEventEndTimeField {
	return ComplexEventEndTimeField{val}
}

//ComplexEventPriceField is a PRICE field
type ComplexEventPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.ComplexEventPrice (1486)
func (f ComplexEventPriceField) Tag() quickfix.Tag { return tag.ComplexEventPrice }

//NewComplexEventPrice returns a new ComplexEventPriceField initialized with val
func NewComplexEventPrice(val quickfix.FIXFloat) ComplexEventPriceField {
	return ComplexEventPriceField{val}
}

//ComplexEventPriceBoundaryMethodField is a INT field
type ComplexEventPriceBoundaryMethodField struct{ quickfix.FIXInt }

//Tag returns tag.ComplexEventPriceBoundaryMethod (1487)
func (f ComplexEventPriceBoundaryMethodField) Tag() quickfix.Tag {
	return tag.ComplexEventPriceBoundaryMethod
}

//NewComplexEventPriceBoundaryMethod returns a new ComplexEventPriceBoundaryMethodField initialized with val
func NewComplexEventPriceBoundaryMethod(val quickfix.FIXInt) ComplexEventPriceBoundaryMethodField {
	return ComplexEventPriceBoundaryMethodField{val}
}

//ComplexEventPriceBoundaryPrecisionField is a PERCENTAGE field
type ComplexEventPriceBoundaryPrecisionField struct{ quickfix.FIXFloat }

//Tag returns tag.ComplexEventPriceBoundaryPrecision (1488)
func (f ComplexEventPriceBoundaryPrecisionField) Tag() quickfix.Tag {
	return tag.ComplexEventPriceBoundaryPrecision
}

//NewComplexEventPriceBoundaryPrecision returns a new ComplexEventPriceBoundaryPrecisionField initialized with val
func NewComplexEventPriceBoundaryPrecision(val quickfix.FIXFloat) ComplexEventPriceBoundaryPrecisionField {
	return ComplexEventPriceBoundaryPrecisionField{val}
}

//ComplexEventPriceTimeTypeField is a INT field
type ComplexEventPriceTimeTypeField struct{ quickfix.FIXInt }

//Tag returns tag.ComplexEventPriceTimeType (1489)
func (f ComplexEventPriceTimeTypeField) Tag() quickfix.Tag { return tag.ComplexEventPriceTimeType }

//NewComplexEventPriceTimeType returns a new ComplexEventPriceTimeTypeField initialized with val
func NewComplexEventPriceTimeType(val quickfix.FIXInt) ComplexEventPriceTimeTypeField {
	return ComplexEventPriceTimeTypeField{val}
}

//ComplexEventStartDateField is a UTCTIMESTAMP field
type ComplexEventStartDateField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.ComplexEventStartDate (1492)
func (f ComplexEventStartDateField) Tag() quickfix.Tag { return tag.ComplexEventStartDate }

//NewComplexEventStartDate returns a new ComplexEventStartDateField initialized with val
func NewComplexEventStartDate(val time.Time) ComplexEventStartDateField {
	return ComplexEventStartDateField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewComplexEventStartDateNoMillis returns a new ComplexEventStartDateField initialized with val without millisecs
func NewComplexEventStartDateNoMillis(val time.Time) ComplexEventStartDateField {
	return ComplexEventStartDateField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//ComplexEventStartTimeField is a UTCTIMEONLY field
type ComplexEventStartTimeField struct{ quickfix.FIXString }

//Tag returns tag.ComplexEventStartTime (1495)
func (f ComplexEventStartTimeField) Tag() quickfix.Tag { return tag.ComplexEventStartTime }

//NewComplexEventStartTime returns a new ComplexEventStartTimeField initialized with val
func NewComplexEventStartTime(val quickfix.FIXString) ComplexEventStartTimeField {
	return ComplexEventStartTimeField{val}
}

//ComplexEventTypeField is a INT field
type ComplexEventTypeField struct{ quickfix.FIXInt }

//Tag returns tag.ComplexEventType (1484)
func (f ComplexEventTypeField) Tag() quickfix.Tag { return tag.ComplexEventType }

//NewComplexEventType returns a new ComplexEventTypeField initialized with val
func NewComplexEventType(val quickfix.FIXInt) ComplexEventTypeField {
	return ComplexEventTypeField{val}
}

//ComplexOptPayoutAmountField is a AMT field
type ComplexOptPayoutAmountField struct{ quickfix.FIXFloat }

//Tag returns tag.ComplexOptPayoutAmount (1485)
func (f ComplexOptPayoutAmountField) Tag() quickfix.Tag { return tag.ComplexOptPayoutAmount }

//NewComplexOptPayoutAmount returns a new ComplexOptPayoutAmountField initialized with val
func NewComplexOptPayoutAmount(val quickfix.FIXFloat) ComplexOptPayoutAmountField {
	return ComplexOptPayoutAmountField{val}
}

//ComplianceIDField is a STRING field
type ComplianceIDField struct{ quickfix.FIXString }

//Tag returns tag.ComplianceID (376)
func (f ComplianceIDField) Tag() quickfix.Tag { return tag.ComplianceID }

//NewComplianceID returns a new ComplianceIDField initialized with val
func NewComplianceID(val quickfix.FIXString) ComplianceIDField {
	return ComplianceIDField{val}
}

//ConcessionField is a AMT field
type ConcessionField struct{ quickfix.FIXFloat }

//Tag returns tag.Concession (238)
func (f ConcessionField) Tag() quickfix.Tag { return tag.Concession }

//NewConcession returns a new ConcessionField initialized with val
func NewConcession(val quickfix.FIXFloat) ConcessionField {
	return ConcessionField{val}
}

//ConfirmIDField is a STRING field
type ConfirmIDField struct{ quickfix.FIXString }

//Tag returns tag.ConfirmID (664)
func (f ConfirmIDField) Tag() quickfix.Tag { return tag.ConfirmID }

//NewConfirmID returns a new ConfirmIDField initialized with val
func NewConfirmID(val quickfix.FIXString) ConfirmIDField {
	return ConfirmIDField{val}
}

//ConfirmRefIDField is a STRING field
type ConfirmRefIDField struct{ quickfix.FIXString }

//Tag returns tag.ConfirmRefID (772)
func (f ConfirmRefIDField) Tag() quickfix.Tag { return tag.ConfirmRefID }

//NewConfirmRefID returns a new ConfirmRefIDField initialized with val
func NewConfirmRefID(val quickfix.FIXString) ConfirmRefIDField {
	return ConfirmRefIDField{val}
}

//ConfirmRejReasonField is a INT field
type ConfirmRejReasonField struct{ quickfix.FIXInt }

//Tag returns tag.ConfirmRejReason (774)
func (f ConfirmRejReasonField) Tag() quickfix.Tag { return tag.ConfirmRejReason }

//NewConfirmRejReason returns a new ConfirmRejReasonField initialized with val
func NewConfirmRejReason(val quickfix.FIXInt) ConfirmRejReasonField {
	return ConfirmRejReasonField{val}
}

//ConfirmReqIDField is a STRING field
type ConfirmReqIDField struct{ quickfix.FIXString }

//Tag returns tag.ConfirmReqID (859)
func (f ConfirmReqIDField) Tag() quickfix.Tag { return tag.ConfirmReqID }

//NewConfirmReqID returns a new ConfirmReqIDField initialized with val
func NewConfirmReqID(val quickfix.FIXString) ConfirmReqIDField {
	return ConfirmReqIDField{val}
}

//ConfirmStatusField is a INT field
type ConfirmStatusField struct{ quickfix.FIXInt }

//Tag returns tag.ConfirmStatus (665)
func (f ConfirmStatusField) Tag() quickfix.Tag { return tag.ConfirmStatus }

//NewConfirmStatus returns a new ConfirmStatusField initialized with val
func NewConfirmStatus(val quickfix.FIXInt) ConfirmStatusField {
	return ConfirmStatusField{val}
}

//ConfirmTransTypeField is a INT field
type ConfirmTransTypeField struct{ quickfix.FIXInt }

//Tag returns tag.ConfirmTransType (666)
func (f ConfirmTransTypeField) Tag() quickfix.Tag { return tag.ConfirmTransType }

//NewConfirmTransType returns a new ConfirmTransTypeField initialized with val
func NewConfirmTransType(val quickfix.FIXInt) ConfirmTransTypeField {
	return ConfirmTransTypeField{val}
}

//ConfirmTypeField is a INT field
type ConfirmTypeField struct{ quickfix.FIXInt }

//Tag returns tag.ConfirmType (773)
func (f ConfirmTypeField) Tag() quickfix.Tag { return tag.ConfirmType }

//NewConfirmType returns a new ConfirmTypeField initialized with val
func NewConfirmType(val quickfix.FIXInt) ConfirmTypeField {
	return ConfirmTypeField{val}
}

//ContAmtCurrField is a CURRENCY field
type ContAmtCurrField struct{ quickfix.FIXString }

//Tag returns tag.ContAmtCurr (521)
func (f ContAmtCurrField) Tag() quickfix.Tag { return tag.ContAmtCurr }

//NewContAmtCurr returns a new ContAmtCurrField initialized with val
func NewContAmtCurr(val quickfix.FIXString) ContAmtCurrField {
	return ContAmtCurrField{val}
}

//ContAmtTypeField is a INT field
type ContAmtTypeField struct{ quickfix.FIXInt }

//Tag returns tag.ContAmtType (519)
func (f ContAmtTypeField) Tag() quickfix.Tag { return tag.ContAmtType }

//NewContAmtType returns a new ContAmtTypeField initialized with val
func NewContAmtType(val quickfix.FIXInt) ContAmtTypeField {
	return ContAmtTypeField{val}
}

//ContAmtValueField is a FLOAT field
type ContAmtValueField struct{ quickfix.FIXFloat }

//Tag returns tag.ContAmtValue (520)
func (f ContAmtValueField) Tag() quickfix.Tag { return tag.ContAmtValue }

//NewContAmtValue returns a new ContAmtValueField initialized with val
func NewContAmtValue(val quickfix.FIXFloat) ContAmtValueField {
	return ContAmtValueField{val}
}

//ContIntRptIDField is a STRING field
type ContIntRptIDField struct{ quickfix.FIXString }

//Tag returns tag.ContIntRptID (977)
func (f ContIntRptIDField) Tag() quickfix.Tag { return tag.ContIntRptID }

//NewContIntRptID returns a new ContIntRptIDField initialized with val
func NewContIntRptID(val quickfix.FIXString) ContIntRptIDField {
	return ContIntRptIDField{val}
}

//ContextPartyIDField is a STRING field
type ContextPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.ContextPartyID (1523)
func (f ContextPartyIDField) Tag() quickfix.Tag { return tag.ContextPartyID }

//NewContextPartyID returns a new ContextPartyIDField initialized with val
func NewContextPartyID(val quickfix.FIXString) ContextPartyIDField {
	return ContextPartyIDField{val}
}

//ContextPartyIDSourceField is a CHAR field
type ContextPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.ContextPartyIDSource (1524)
func (f ContextPartyIDSourceField) Tag() quickfix.Tag { return tag.ContextPartyIDSource }

//NewContextPartyIDSource returns a new ContextPartyIDSourceField initialized with val
func NewContextPartyIDSource(val quickfix.FIXString) ContextPartyIDSourceField {
	return ContextPartyIDSourceField{val}
}

//ContextPartyRoleField is a INT field
type ContextPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.ContextPartyRole (1525)
func (f ContextPartyRoleField) Tag() quickfix.Tag { return tag.ContextPartyRole }

//NewContextPartyRole returns a new ContextPartyRoleField initialized with val
func NewContextPartyRole(val quickfix.FIXInt) ContextPartyRoleField {
	return ContextPartyRoleField{val}
}

//ContextPartySubIDField is a STRING field
type ContextPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.ContextPartySubID (1527)
func (f ContextPartySubIDField) Tag() quickfix.Tag { return tag.ContextPartySubID }

//NewContextPartySubID returns a new ContextPartySubIDField initialized with val
func NewContextPartySubID(val quickfix.FIXString) ContextPartySubIDField {
	return ContextPartySubIDField{val}
}

//ContextPartySubIDTypeField is a INT field
type ContextPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.ContextPartySubIDType (1528)
func (f ContextPartySubIDTypeField) Tag() quickfix.Tag { return tag.ContextPartySubIDType }

//NewContextPartySubIDType returns a new ContextPartySubIDTypeField initialized with val
func NewContextPartySubIDType(val quickfix.FIXInt) ContextPartySubIDTypeField {
	return ContextPartySubIDTypeField{val}
}

//ContingencyTypeField is a INT field
type ContingencyTypeField struct{ quickfix.FIXInt }

//Tag returns tag.ContingencyType (1385)
func (f ContingencyTypeField) Tag() quickfix.Tag { return tag.ContingencyType }

//NewContingencyType returns a new ContingencyTypeField initialized with val
func NewContingencyType(val quickfix.FIXInt) ContingencyTypeField {
	return ContingencyTypeField{val}
}

//ContraBrokerField is a STRING field
type ContraBrokerField struct{ quickfix.FIXString }

//Tag returns tag.ContraBroker (375)
func (f ContraBrokerField) Tag() quickfix.Tag { return tag.ContraBroker }

//NewContraBroker returns a new ContraBrokerField initialized with val
func NewContraBroker(val quickfix.FIXString) ContraBrokerField {
	return ContraBrokerField{val}
}

//ContraLegRefIDField is a STRING field
type ContraLegRefIDField struct{ quickfix.FIXString }

//Tag returns tag.ContraLegRefID (655)
func (f ContraLegRefIDField) Tag() quickfix.Tag { return tag.ContraLegRefID }

//NewContraLegRefID returns a new ContraLegRefIDField initialized with val
func NewContraLegRefID(val quickfix.FIXString) ContraLegRefIDField {
	return ContraLegRefIDField{val}
}

//ContraTradeQtyField is a QTY field
type ContraTradeQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.ContraTradeQty (437)
func (f ContraTradeQtyField) Tag() quickfix.Tag { return tag.ContraTradeQty }

//NewContraTradeQty returns a new ContraTradeQtyField initialized with val
func NewContraTradeQty(val quickfix.FIXFloat) ContraTradeQtyField {
	return ContraTradeQtyField{val}
}

//ContraTradeTimeField is a UTCTIMESTAMP field
type ContraTradeTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.ContraTradeTime (438)
func (f ContraTradeTimeField) Tag() quickfix.Tag { return tag.ContraTradeTime }

//NewContraTradeTime returns a new ContraTradeTimeField initialized with val
func NewContraTradeTime(val time.Time) ContraTradeTimeField {
	return ContraTradeTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewContraTradeTimeNoMillis returns a new ContraTradeTimeField initialized with val without millisecs
func NewContraTradeTimeNoMillis(val time.Time) ContraTradeTimeField {
	return ContraTradeTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//ContraTraderField is a STRING field
type ContraTraderField struct{ quickfix.FIXString }

//Tag returns tag.ContraTrader (337)
func (f ContraTraderField) Tag() quickfix.Tag { return tag.ContraTrader }

//NewContraTrader returns a new ContraTraderField initialized with val
func NewContraTrader(val quickfix.FIXString) ContraTraderField {
	return ContraTraderField{val}
}

//ContractMultiplierField is a FLOAT field
type ContractMultiplierField struct{ quickfix.FIXFloat }

//Tag returns tag.ContractMultiplier (231)
func (f ContractMultiplierField) Tag() quickfix.Tag { return tag.ContractMultiplier }

//NewContractMultiplier returns a new ContractMultiplierField initialized with val
func NewContractMultiplier(val quickfix.FIXFloat) ContractMultiplierField {
	return ContractMultiplierField{val}
}

//ContractMultiplierUnitField is a INT field
type ContractMultiplierUnitField struct{ quickfix.FIXInt }

//Tag returns tag.ContractMultiplierUnit (1435)
func (f ContractMultiplierUnitField) Tag() quickfix.Tag { return tag.ContractMultiplierUnit }

//NewContractMultiplierUnit returns a new ContractMultiplierUnitField initialized with val
func NewContractMultiplierUnit(val quickfix.FIXInt) ContractMultiplierUnitField {
	return ContractMultiplierUnitField{val}
}

//ContractSettlMonthField is a MONTHYEAR field
type ContractSettlMonthField struct{ quickfix.FIXString }

//Tag returns tag.ContractSettlMonth (667)
func (f ContractSettlMonthField) Tag() quickfix.Tag { return tag.ContractSettlMonth }

//NewContractSettlMonth returns a new ContractSettlMonthField initialized with val
func NewContractSettlMonth(val quickfix.FIXString) ContractSettlMonthField {
	return ContractSettlMonthField{val}
}

//ContraryInstructionIndicatorField is a BOOLEAN field
type ContraryInstructionIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.ContraryInstructionIndicator (719)
func (f ContraryInstructionIndicatorField) Tag() quickfix.Tag { return tag.ContraryInstructionIndicator }

//NewContraryInstructionIndicator returns a new ContraryInstructionIndicatorField initialized with val
func NewContraryInstructionIndicator(val quickfix.FIXBoolean) ContraryInstructionIndicatorField {
	return ContraryInstructionIndicatorField{val}
}

//CopyMsgIndicatorField is a BOOLEAN field
type CopyMsgIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.CopyMsgIndicator (797)
func (f CopyMsgIndicatorField) Tag() quickfix.Tag { return tag.CopyMsgIndicator }

//NewCopyMsgIndicator returns a new CopyMsgIndicatorField initialized with val
func NewCopyMsgIndicator(val quickfix.FIXBoolean) CopyMsgIndicatorField {
	return CopyMsgIndicatorField{val}
}

//CorporateActionField is a MULTIPLECHARVALUE field
type CorporateActionField struct{ quickfix.FIXString }

//Tag returns tag.CorporateAction (292)
func (f CorporateActionField) Tag() quickfix.Tag { return tag.CorporateAction }

//NewCorporateAction returns a new CorporateActionField initialized with val
func NewCorporateAction(val quickfix.FIXString) CorporateActionField {
	return CorporateActionField{val}
}

//CountryField is a COUNTRY field
type CountryField struct{ quickfix.FIXString }

//Tag returns tag.Country (421)
func (f CountryField) Tag() quickfix.Tag { return tag.Country }

//NewCountry returns a new CountryField initialized with val
func NewCountry(val quickfix.FIXString) CountryField {
	return CountryField{val}
}

//CountryOfIssueField is a COUNTRY field
type CountryOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.CountryOfIssue (470)
func (f CountryOfIssueField) Tag() quickfix.Tag { return tag.CountryOfIssue }

//NewCountryOfIssue returns a new CountryOfIssueField initialized with val
func NewCountryOfIssue(val quickfix.FIXString) CountryOfIssueField {
	return CountryOfIssueField{val}
}

//CouponPaymentDateField is a LOCALMKTDATE field
type CouponPaymentDateField struct{ quickfix.FIXString }

//Tag returns tag.CouponPaymentDate (224)
func (f CouponPaymentDateField) Tag() quickfix.Tag { return tag.CouponPaymentDate }

//NewCouponPaymentDate returns a new CouponPaymentDateField initialized with val
func NewCouponPaymentDate(val quickfix.FIXString) CouponPaymentDateField {
	return CouponPaymentDateField{val}
}

//CouponRateField is a PERCENTAGE field
type CouponRateField struct{ quickfix.FIXFloat }

//Tag returns tag.CouponRate (223)
func (f CouponRateField) Tag() quickfix.Tag { return tag.CouponRate }

//NewCouponRate returns a new CouponRateField initialized with val
func NewCouponRate(val quickfix.FIXFloat) CouponRateField {
	return CouponRateField{val}
}

//CoveredOrUncoveredField is a INT field
type CoveredOrUncoveredField struct{ quickfix.FIXInt }

//Tag returns tag.CoveredOrUncovered (203)
func (f CoveredOrUncoveredField) Tag() quickfix.Tag { return tag.CoveredOrUncovered }

//NewCoveredOrUncovered returns a new CoveredOrUncoveredField initialized with val
func NewCoveredOrUncovered(val quickfix.FIXInt) CoveredOrUncoveredField {
	return CoveredOrUncoveredField{val}
}

//CreditRatingField is a STRING field
type CreditRatingField struct{ quickfix.FIXString }

//Tag returns tag.CreditRating (255)
func (f CreditRatingField) Tag() quickfix.Tag { return tag.CreditRating }

//NewCreditRating returns a new CreditRatingField initialized with val
func NewCreditRating(val quickfix.FIXString) CreditRatingField {
	return CreditRatingField{val}
}

//CrossIDField is a STRING field
type CrossIDField struct{ quickfix.FIXString }

//Tag returns tag.CrossID (548)
func (f CrossIDField) Tag() quickfix.Tag { return tag.CrossID }

//NewCrossID returns a new CrossIDField initialized with val
func NewCrossID(val quickfix.FIXString) CrossIDField {
	return CrossIDField{val}
}

//CrossPercentField is a PERCENTAGE field
type CrossPercentField struct{ quickfix.FIXFloat }

//Tag returns tag.CrossPercent (413)
func (f CrossPercentField) Tag() quickfix.Tag { return tag.CrossPercent }

//NewCrossPercent returns a new CrossPercentField initialized with val
func NewCrossPercent(val quickfix.FIXFloat) CrossPercentField {
	return CrossPercentField{val}
}

//CrossPrioritizationField is a INT field
type CrossPrioritizationField struct{ quickfix.FIXInt }

//Tag returns tag.CrossPrioritization (550)
func (f CrossPrioritizationField) Tag() quickfix.Tag { return tag.CrossPrioritization }

//NewCrossPrioritization returns a new CrossPrioritizationField initialized with val
func NewCrossPrioritization(val quickfix.FIXInt) CrossPrioritizationField {
	return CrossPrioritizationField{val}
}

//CrossTypeField is a INT field
type CrossTypeField struct{ quickfix.FIXInt }

//Tag returns tag.CrossType (549)
func (f CrossTypeField) Tag() quickfix.Tag { return tag.CrossType }

//NewCrossType returns a new CrossTypeField initialized with val
func NewCrossType(val quickfix.FIXInt) CrossTypeField {
	return CrossTypeField{val}
}

//CstmApplVerIDField is a STRING field
type CstmApplVerIDField struct{ quickfix.FIXString }

//Tag returns tag.CstmApplVerID (1129)
func (f CstmApplVerIDField) Tag() quickfix.Tag { return tag.CstmApplVerID }

//NewCstmApplVerID returns a new CstmApplVerIDField initialized with val
func NewCstmApplVerID(val quickfix.FIXString) CstmApplVerIDField {
	return CstmApplVerIDField{val}
}

//CumQtyField is a QTY field
type CumQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.CumQty (14)
func (f CumQtyField) Tag() quickfix.Tag { return tag.CumQty }

//NewCumQty returns a new CumQtyField initialized with val
func NewCumQty(val quickfix.FIXFloat) CumQtyField {
	return CumQtyField{val}
}

//CurrencyField is a CURRENCY field
type CurrencyField struct{ quickfix.FIXString }

//Tag returns tag.Currency (15)
func (f CurrencyField) Tag() quickfix.Tag { return tag.Currency }

//NewCurrency returns a new CurrencyField initialized with val
func NewCurrency(val quickfix.FIXString) CurrencyField {
	return CurrencyField{val}
}

//CurrencyRatioField is a FLOAT field
type CurrencyRatioField struct{ quickfix.FIXFloat }

//Tag returns tag.CurrencyRatio (1382)
func (f CurrencyRatioField) Tag() quickfix.Tag { return tag.CurrencyRatio }

//NewCurrencyRatio returns a new CurrencyRatioField initialized with val
func NewCurrencyRatio(val quickfix.FIXFloat) CurrencyRatioField {
	return CurrencyRatioField{val}
}

//CustDirectedOrderField is a BOOLEAN field
type CustDirectedOrderField struct{ quickfix.FIXBoolean }

//Tag returns tag.CustDirectedOrder (1029)
func (f CustDirectedOrderField) Tag() quickfix.Tag { return tag.CustDirectedOrder }

//NewCustDirectedOrder returns a new CustDirectedOrderField initialized with val
func NewCustDirectedOrder(val quickfix.FIXBoolean) CustDirectedOrderField {
	return CustDirectedOrderField{val}
}

//CustOrderCapacityField is a INT field
type CustOrderCapacityField struct{ quickfix.FIXInt }

//Tag returns tag.CustOrderCapacity (582)
func (f CustOrderCapacityField) Tag() quickfix.Tag { return tag.CustOrderCapacity }

//NewCustOrderCapacity returns a new CustOrderCapacityField initialized with val
func NewCustOrderCapacity(val quickfix.FIXInt) CustOrderCapacityField {
	return CustOrderCapacityField{val}
}

//CustOrderHandlingInstField is a MULTIPLESTRINGVALUE field
type CustOrderHandlingInstField struct{ quickfix.FIXString }

//Tag returns tag.CustOrderHandlingInst (1031)
func (f CustOrderHandlingInstField) Tag() quickfix.Tag { return tag.CustOrderHandlingInst }

//NewCustOrderHandlingInst returns a new CustOrderHandlingInstField initialized with val
func NewCustOrderHandlingInst(val quickfix.FIXString) CustOrderHandlingInstField {
	return CustOrderHandlingInstField{val}
}

//CustomerOrFirmField is a INT field
type CustomerOrFirmField struct{ quickfix.FIXInt }

//Tag returns tag.CustomerOrFirm (204)
func (f CustomerOrFirmField) Tag() quickfix.Tag { return tag.CustomerOrFirm }

//NewCustomerOrFirm returns a new CustomerOrFirmField initialized with val
func NewCustomerOrFirm(val quickfix.FIXInt) CustomerOrFirmField {
	return CustomerOrFirmField{val}
}

//CxlQtyField is a QTY field
type CxlQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.CxlQty (84)
func (f CxlQtyField) Tag() quickfix.Tag { return tag.CxlQty }

//NewCxlQty returns a new CxlQtyField initialized with val
func NewCxlQty(val quickfix.FIXFloat) CxlQtyField {
	return CxlQtyField{val}
}

//CxlRejReasonField is a INT field
type CxlRejReasonField struct{ quickfix.FIXInt }

//Tag returns tag.CxlRejReason (102)
func (f CxlRejReasonField) Tag() quickfix.Tag { return tag.CxlRejReason }

//NewCxlRejReason returns a new CxlRejReasonField initialized with val
func NewCxlRejReason(val quickfix.FIXInt) CxlRejReasonField {
	return CxlRejReasonField{val}
}

//CxlRejResponseToField is a CHAR field
type CxlRejResponseToField struct{ quickfix.FIXString }

//Tag returns tag.CxlRejResponseTo (434)
func (f CxlRejResponseToField) Tag() quickfix.Tag { return tag.CxlRejResponseTo }

//NewCxlRejResponseTo returns a new CxlRejResponseToField initialized with val
func NewCxlRejResponseTo(val quickfix.FIXString) CxlRejResponseToField {
	return CxlRejResponseToField{val}
}

//CxlTypeField is a CHAR field
type CxlTypeField struct{ quickfix.FIXString }

//Tag returns tag.CxlType (125)
func (f CxlTypeField) Tag() quickfix.Tag { return tag.CxlType }

//NewCxlType returns a new CxlTypeField initialized with val
func NewCxlType(val quickfix.FIXString) CxlTypeField {
	return CxlTypeField{val}
}

//DKReasonField is a CHAR field
type DKReasonField struct{ quickfix.FIXString }

//Tag returns tag.DKReason (127)
func (f DKReasonField) Tag() quickfix.Tag { return tag.DKReason }

//NewDKReason returns a new DKReasonField initialized with val
func NewDKReason(val quickfix.FIXString) DKReasonField {
	return DKReasonField{val}
}

//DateOfBirthField is a LOCALMKTDATE field
type DateOfBirthField struct{ quickfix.FIXString }

//Tag returns tag.DateOfBirth (486)
func (f DateOfBirthField) Tag() quickfix.Tag { return tag.DateOfBirth }

//NewDateOfBirth returns a new DateOfBirthField initialized with val
func NewDateOfBirth(val quickfix.FIXString) DateOfBirthField {
	return DateOfBirthField{val}
}

//DatedDateField is a LOCALMKTDATE field
type DatedDateField struct{ quickfix.FIXString }

//Tag returns tag.DatedDate (873)
func (f DatedDateField) Tag() quickfix.Tag { return tag.DatedDate }

//NewDatedDate returns a new DatedDateField initialized with val
func NewDatedDate(val quickfix.FIXString) DatedDateField {
	return DatedDateField{val}
}

//DayAvgPxField is a PRICE field
type DayAvgPxField struct{ quickfix.FIXFloat }

//Tag returns tag.DayAvgPx (426)
func (f DayAvgPxField) Tag() quickfix.Tag { return tag.DayAvgPx }

//NewDayAvgPx returns a new DayAvgPxField initialized with val
func NewDayAvgPx(val quickfix.FIXFloat) DayAvgPxField {
	return DayAvgPxField{val}
}

//DayBookingInstField is a CHAR field
type DayBookingInstField struct{ quickfix.FIXString }

//Tag returns tag.DayBookingInst (589)
func (f DayBookingInstField) Tag() quickfix.Tag { return tag.DayBookingInst }

//NewDayBookingInst returns a new DayBookingInstField initialized with val
func NewDayBookingInst(val quickfix.FIXString) DayBookingInstField {
	return DayBookingInstField{val}
}

//DayCumQtyField is a QTY field
type DayCumQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.DayCumQty (425)
func (f DayCumQtyField) Tag() quickfix.Tag { return tag.DayCumQty }

//NewDayCumQty returns a new DayCumQtyField initialized with val
func NewDayCumQty(val quickfix.FIXFloat) DayCumQtyField {
	return DayCumQtyField{val}
}

//DayOrderQtyField is a QTY field
type DayOrderQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.DayOrderQty (424)
func (f DayOrderQtyField) Tag() quickfix.Tag { return tag.DayOrderQty }

//NewDayOrderQty returns a new DayOrderQtyField initialized with val
func NewDayOrderQty(val quickfix.FIXFloat) DayOrderQtyField {
	return DayOrderQtyField{val}
}

//DealingCapacityField is a CHAR field
type DealingCapacityField struct{ quickfix.FIXString }

//Tag returns tag.DealingCapacity (1048)
func (f DealingCapacityField) Tag() quickfix.Tag { return tag.DealingCapacity }

//NewDealingCapacity returns a new DealingCapacityField initialized with val
func NewDealingCapacity(val quickfix.FIXString) DealingCapacityField {
	return DealingCapacityField{val}
}

//DefBidSizeField is a QTY field
type DefBidSizeField struct{ quickfix.FIXFloat }

//Tag returns tag.DefBidSize (293)
func (f DefBidSizeField) Tag() quickfix.Tag { return tag.DefBidSize }

//NewDefBidSize returns a new DefBidSizeField initialized with val
func NewDefBidSize(val quickfix.FIXFloat) DefBidSizeField {
	return DefBidSizeField{val}
}

//DefOfferSizeField is a QTY field
type DefOfferSizeField struct{ quickfix.FIXFloat }

//Tag returns tag.DefOfferSize (294)
func (f DefOfferSizeField) Tag() quickfix.Tag { return tag.DefOfferSize }

//NewDefOfferSize returns a new DefOfferSizeField initialized with val
func NewDefOfferSize(val quickfix.FIXFloat) DefOfferSizeField {
	return DefOfferSizeField{val}
}

//DefaultApplExtIDField is a INT field
type DefaultApplExtIDField struct{ quickfix.FIXInt }

//Tag returns tag.DefaultApplExtID (1407)
func (f DefaultApplExtIDField) Tag() quickfix.Tag { return tag.DefaultApplExtID }

//NewDefaultApplExtID returns a new DefaultApplExtIDField initialized with val
func NewDefaultApplExtID(val quickfix.FIXInt) DefaultApplExtIDField {
	return DefaultApplExtIDField{val}
}

//DefaultApplVerIDField is a STRING field
type DefaultApplVerIDField struct{ quickfix.FIXString }

//Tag returns tag.DefaultApplVerID (1137)
func (f DefaultApplVerIDField) Tag() quickfix.Tag { return tag.DefaultApplVerID }

//NewDefaultApplVerID returns a new DefaultApplVerIDField initialized with val
func NewDefaultApplVerID(val quickfix.FIXString) DefaultApplVerIDField {
	return DefaultApplVerIDField{val}
}

//DefaultCstmApplVerIDField is a STRING field
type DefaultCstmApplVerIDField struct{ quickfix.FIXString }

//Tag returns tag.DefaultCstmApplVerID (1408)
func (f DefaultCstmApplVerIDField) Tag() quickfix.Tag { return tag.DefaultCstmApplVerID }

//NewDefaultCstmApplVerID returns a new DefaultCstmApplVerIDField initialized with val
func NewDefaultCstmApplVerID(val quickfix.FIXString) DefaultCstmApplVerIDField {
	return DefaultCstmApplVerIDField{val}
}

//DefaultVerIndicatorField is a BOOLEAN field
type DefaultVerIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.DefaultVerIndicator (1410)
func (f DefaultVerIndicatorField) Tag() quickfix.Tag { return tag.DefaultVerIndicator }

//NewDefaultVerIndicator returns a new DefaultVerIndicatorField initialized with val
func NewDefaultVerIndicator(val quickfix.FIXBoolean) DefaultVerIndicatorField {
	return DefaultVerIndicatorField{val}
}

//DeleteReasonField is a CHAR field
type DeleteReasonField struct{ quickfix.FIXString }

//Tag returns tag.DeleteReason (285)
func (f DeleteReasonField) Tag() quickfix.Tag { return tag.DeleteReason }

//NewDeleteReason returns a new DeleteReasonField initialized with val
func NewDeleteReason(val quickfix.FIXString) DeleteReasonField {
	return DeleteReasonField{val}
}

//DeliverToCompIDField is a STRING field
type DeliverToCompIDField struct{ quickfix.FIXString }

//Tag returns tag.DeliverToCompID (128)
func (f DeliverToCompIDField) Tag() quickfix.Tag { return tag.DeliverToCompID }

//NewDeliverToCompID returns a new DeliverToCompIDField initialized with val
func NewDeliverToCompID(val quickfix.FIXString) DeliverToCompIDField {
	return DeliverToCompIDField{val}
}

//DeliverToLocationIDField is a STRING field
type DeliverToLocationIDField struct{ quickfix.FIXString }

//Tag returns tag.DeliverToLocationID (145)
func (f DeliverToLocationIDField) Tag() quickfix.Tag { return tag.DeliverToLocationID }

//NewDeliverToLocationID returns a new DeliverToLocationIDField initialized with val
func NewDeliverToLocationID(val quickfix.FIXString) DeliverToLocationIDField {
	return DeliverToLocationIDField{val}
}

//DeliverToSubIDField is a STRING field
type DeliverToSubIDField struct{ quickfix.FIXString }

//Tag returns tag.DeliverToSubID (129)
func (f DeliverToSubIDField) Tag() quickfix.Tag { return tag.DeliverToSubID }

//NewDeliverToSubID returns a new DeliverToSubIDField initialized with val
func NewDeliverToSubID(val quickfix.FIXString) DeliverToSubIDField {
	return DeliverToSubIDField{val}
}

//DeliveryDateField is a LOCALMKTDATE field
type DeliveryDateField struct{ quickfix.FIXString }

//Tag returns tag.DeliveryDate (743)
func (f DeliveryDateField) Tag() quickfix.Tag { return tag.DeliveryDate }

//NewDeliveryDate returns a new DeliveryDateField initialized with val
func NewDeliveryDate(val quickfix.FIXString) DeliveryDateField {
	return DeliveryDateField{val}
}

//DeliveryFormField is a INT field
type DeliveryFormField struct{ quickfix.FIXInt }

//Tag returns tag.DeliveryForm (668)
func (f DeliveryFormField) Tag() quickfix.Tag { return tag.DeliveryForm }

//NewDeliveryForm returns a new DeliveryFormField initialized with val
func NewDeliveryForm(val quickfix.FIXInt) DeliveryFormField {
	return DeliveryFormField{val}
}

//DeliveryTypeField is a INT field
type DeliveryTypeField struct{ quickfix.FIXInt }

//Tag returns tag.DeliveryType (919)
func (f DeliveryTypeField) Tag() quickfix.Tag { return tag.DeliveryType }

//NewDeliveryType returns a new DeliveryTypeField initialized with val
func NewDeliveryType(val quickfix.FIXInt) DeliveryTypeField {
	return DeliveryTypeField{val}
}

//DerivFlexProductEligibilityIndicatorField is a BOOLEAN field
type DerivFlexProductEligibilityIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.DerivFlexProductEligibilityIndicator (1243)
func (f DerivFlexProductEligibilityIndicatorField) Tag() quickfix.Tag {
	return tag.DerivFlexProductEligibilityIndicator
}

//NewDerivFlexProductEligibilityIndicator returns a new DerivFlexProductEligibilityIndicatorField initialized with val
func NewDerivFlexProductEligibilityIndicator(val quickfix.FIXBoolean) DerivFlexProductEligibilityIndicatorField {
	return DerivFlexProductEligibilityIndicatorField{val}
}

//DerivativeCFICodeField is a STRING field
type DerivativeCFICodeField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeCFICode (1248)
func (f DerivativeCFICodeField) Tag() quickfix.Tag { return tag.DerivativeCFICode }

//NewDerivativeCFICode returns a new DerivativeCFICodeField initialized with val
func NewDerivativeCFICode(val quickfix.FIXString) DerivativeCFICodeField {
	return DerivativeCFICodeField{val}
}

//DerivativeCapPriceField is a PRICE field
type DerivativeCapPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.DerivativeCapPrice (1321)
func (f DerivativeCapPriceField) Tag() quickfix.Tag { return tag.DerivativeCapPrice }

//NewDerivativeCapPrice returns a new DerivativeCapPriceField initialized with val
func NewDerivativeCapPrice(val quickfix.FIXFloat) DerivativeCapPriceField {
	return DerivativeCapPriceField{val}
}

//DerivativeContractMultiplierField is a FLOAT field
type DerivativeContractMultiplierField struct{ quickfix.FIXFloat }

//Tag returns tag.DerivativeContractMultiplier (1266)
func (f DerivativeContractMultiplierField) Tag() quickfix.Tag { return tag.DerivativeContractMultiplier }

//NewDerivativeContractMultiplier returns a new DerivativeContractMultiplierField initialized with val
func NewDerivativeContractMultiplier(val quickfix.FIXFloat) DerivativeContractMultiplierField {
	return DerivativeContractMultiplierField{val}
}

//DerivativeContractMultiplierUnitField is a INT field
type DerivativeContractMultiplierUnitField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeContractMultiplierUnit (1438)
func (f DerivativeContractMultiplierUnitField) Tag() quickfix.Tag {
	return tag.DerivativeContractMultiplierUnit
}

//NewDerivativeContractMultiplierUnit returns a new DerivativeContractMultiplierUnitField initialized with val
func NewDerivativeContractMultiplierUnit(val quickfix.FIXInt) DerivativeContractMultiplierUnitField {
	return DerivativeContractMultiplierUnitField{val}
}

//DerivativeContractSettlMonthField is a MONTHYEAR field
type DerivativeContractSettlMonthField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeContractSettlMonth (1285)
func (f DerivativeContractSettlMonthField) Tag() quickfix.Tag { return tag.DerivativeContractSettlMonth }

//NewDerivativeContractSettlMonth returns a new DerivativeContractSettlMonthField initialized with val
func NewDerivativeContractSettlMonth(val quickfix.FIXString) DerivativeContractSettlMonthField {
	return DerivativeContractSettlMonthField{val}
}

//DerivativeCountryOfIssueField is a COUNTRY field
type DerivativeCountryOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeCountryOfIssue (1258)
func (f DerivativeCountryOfIssueField) Tag() quickfix.Tag { return tag.DerivativeCountryOfIssue }

//NewDerivativeCountryOfIssue returns a new DerivativeCountryOfIssueField initialized with val
func NewDerivativeCountryOfIssue(val quickfix.FIXString) DerivativeCountryOfIssueField {
	return DerivativeCountryOfIssueField{val}
}

//DerivativeEncodedIssuerField is a DATA field
type DerivativeEncodedIssuerField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeEncodedIssuer (1278)
func (f DerivativeEncodedIssuerField) Tag() quickfix.Tag { return tag.DerivativeEncodedIssuer }

//NewDerivativeEncodedIssuer returns a new DerivativeEncodedIssuerField initialized with val
func NewDerivativeEncodedIssuer(val quickfix.FIXString) DerivativeEncodedIssuerField {
	return DerivativeEncodedIssuerField{val}
}

//DerivativeEncodedIssuerLenField is a LENGTH field
type DerivativeEncodedIssuerLenField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeEncodedIssuerLen (1277)
func (f DerivativeEncodedIssuerLenField) Tag() quickfix.Tag { return tag.DerivativeEncodedIssuerLen }

//NewDerivativeEncodedIssuerLen returns a new DerivativeEncodedIssuerLenField initialized with val
func NewDerivativeEncodedIssuerLen(val quickfix.FIXInt) DerivativeEncodedIssuerLenField {
	return DerivativeEncodedIssuerLenField{val}
}

//DerivativeEncodedSecurityDescField is a DATA field
type DerivativeEncodedSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeEncodedSecurityDesc (1281)
func (f DerivativeEncodedSecurityDescField) Tag() quickfix.Tag {
	return tag.DerivativeEncodedSecurityDesc
}

//NewDerivativeEncodedSecurityDesc returns a new DerivativeEncodedSecurityDescField initialized with val
func NewDerivativeEncodedSecurityDesc(val quickfix.FIXString) DerivativeEncodedSecurityDescField {
	return DerivativeEncodedSecurityDescField{val}
}

//DerivativeEncodedSecurityDescLenField is a LENGTH field
type DerivativeEncodedSecurityDescLenField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeEncodedSecurityDescLen (1280)
func (f DerivativeEncodedSecurityDescLenField) Tag() quickfix.Tag {
	return tag.DerivativeEncodedSecurityDescLen
}

//NewDerivativeEncodedSecurityDescLen returns a new DerivativeEncodedSecurityDescLenField initialized with val
func NewDerivativeEncodedSecurityDescLen(val quickfix.FIXInt) DerivativeEncodedSecurityDescLenField {
	return DerivativeEncodedSecurityDescLenField{val}
}

//DerivativeEventDateField is a LOCALMKTDATE field
type DerivativeEventDateField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeEventDate (1288)
func (f DerivativeEventDateField) Tag() quickfix.Tag { return tag.DerivativeEventDate }

//NewDerivativeEventDate returns a new DerivativeEventDateField initialized with val
func NewDerivativeEventDate(val quickfix.FIXString) DerivativeEventDateField {
	return DerivativeEventDateField{val}
}

//DerivativeEventPxField is a PRICE field
type DerivativeEventPxField struct{ quickfix.FIXFloat }

//Tag returns tag.DerivativeEventPx (1290)
func (f DerivativeEventPxField) Tag() quickfix.Tag { return tag.DerivativeEventPx }

//NewDerivativeEventPx returns a new DerivativeEventPxField initialized with val
func NewDerivativeEventPx(val quickfix.FIXFloat) DerivativeEventPxField {
	return DerivativeEventPxField{val}
}

//DerivativeEventTextField is a STRING field
type DerivativeEventTextField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeEventText (1291)
func (f DerivativeEventTextField) Tag() quickfix.Tag { return tag.DerivativeEventText }

//NewDerivativeEventText returns a new DerivativeEventTextField initialized with val
func NewDerivativeEventText(val quickfix.FIXString) DerivativeEventTextField {
	return DerivativeEventTextField{val}
}

//DerivativeEventTimeField is a UTCTIMESTAMP field
type DerivativeEventTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.DerivativeEventTime (1289)
func (f DerivativeEventTimeField) Tag() quickfix.Tag { return tag.DerivativeEventTime }

//NewDerivativeEventTime returns a new DerivativeEventTimeField initialized with val
func NewDerivativeEventTime(val time.Time) DerivativeEventTimeField {
	return DerivativeEventTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewDerivativeEventTimeNoMillis returns a new DerivativeEventTimeField initialized with val without millisecs
func NewDerivativeEventTimeNoMillis(val time.Time) DerivativeEventTimeField {
	return DerivativeEventTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//DerivativeEventTypeField is a INT field
type DerivativeEventTypeField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeEventType (1287)
func (f DerivativeEventTypeField) Tag() quickfix.Tag { return tag.DerivativeEventType }

//NewDerivativeEventType returns a new DerivativeEventTypeField initialized with val
func NewDerivativeEventType(val quickfix.FIXInt) DerivativeEventTypeField {
	return DerivativeEventTypeField{val}
}

//DerivativeExerciseStyleField is a CHAR field
type DerivativeExerciseStyleField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeExerciseStyle (1299)
func (f DerivativeExerciseStyleField) Tag() quickfix.Tag { return tag.DerivativeExerciseStyle }

//NewDerivativeExerciseStyle returns a new DerivativeExerciseStyleField initialized with val
func NewDerivativeExerciseStyle(val quickfix.FIXString) DerivativeExerciseStyleField {
	return DerivativeExerciseStyleField{val}
}

//DerivativeFloorPriceField is a PRICE field
type DerivativeFloorPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.DerivativeFloorPrice (1322)
func (f DerivativeFloorPriceField) Tag() quickfix.Tag { return tag.DerivativeFloorPrice }

//NewDerivativeFloorPrice returns a new DerivativeFloorPriceField initialized with val
func NewDerivativeFloorPrice(val quickfix.FIXFloat) DerivativeFloorPriceField {
	return DerivativeFloorPriceField{val}
}

//DerivativeFlowScheduleTypeField is a INT field
type DerivativeFlowScheduleTypeField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeFlowScheduleType (1442)
func (f DerivativeFlowScheduleTypeField) Tag() quickfix.Tag { return tag.DerivativeFlowScheduleType }

//NewDerivativeFlowScheduleType returns a new DerivativeFlowScheduleTypeField initialized with val
func NewDerivativeFlowScheduleType(val quickfix.FIXInt) DerivativeFlowScheduleTypeField {
	return DerivativeFlowScheduleTypeField{val}
}

//DerivativeFuturesValuationMethodField is a STRING field
type DerivativeFuturesValuationMethodField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeFuturesValuationMethod (1319)
func (f DerivativeFuturesValuationMethodField) Tag() quickfix.Tag {
	return tag.DerivativeFuturesValuationMethod
}

//NewDerivativeFuturesValuationMethod returns a new DerivativeFuturesValuationMethodField initialized with val
func NewDerivativeFuturesValuationMethod(val quickfix.FIXString) DerivativeFuturesValuationMethodField {
	return DerivativeFuturesValuationMethodField{val}
}

//DerivativeInstrAttribTypeField is a INT field
type DerivativeInstrAttribTypeField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeInstrAttribType (1313)
func (f DerivativeInstrAttribTypeField) Tag() quickfix.Tag { return tag.DerivativeInstrAttribType }

//NewDerivativeInstrAttribType returns a new DerivativeInstrAttribTypeField initialized with val
func NewDerivativeInstrAttribType(val quickfix.FIXInt) DerivativeInstrAttribTypeField {
	return DerivativeInstrAttribTypeField{val}
}

//DerivativeInstrAttribValueField is a STRING field
type DerivativeInstrAttribValueField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeInstrAttribValue (1314)
func (f DerivativeInstrAttribValueField) Tag() quickfix.Tag { return tag.DerivativeInstrAttribValue }

//NewDerivativeInstrAttribValue returns a new DerivativeInstrAttribValueField initialized with val
func NewDerivativeInstrAttribValue(val quickfix.FIXString) DerivativeInstrAttribValueField {
	return DerivativeInstrAttribValueField{val}
}

//DerivativeInstrRegistryField is a STRING field
type DerivativeInstrRegistryField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeInstrRegistry (1257)
func (f DerivativeInstrRegistryField) Tag() quickfix.Tag { return tag.DerivativeInstrRegistry }

//NewDerivativeInstrRegistry returns a new DerivativeInstrRegistryField initialized with val
func NewDerivativeInstrRegistry(val quickfix.FIXString) DerivativeInstrRegistryField {
	return DerivativeInstrRegistryField{val}
}

//DerivativeInstrmtAssignmentMethodField is a CHAR field
type DerivativeInstrmtAssignmentMethodField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeInstrmtAssignmentMethod (1255)
func (f DerivativeInstrmtAssignmentMethodField) Tag() quickfix.Tag {
	return tag.DerivativeInstrmtAssignmentMethod
}

//NewDerivativeInstrmtAssignmentMethod returns a new DerivativeInstrmtAssignmentMethodField initialized with val
func NewDerivativeInstrmtAssignmentMethod(val quickfix.FIXString) DerivativeInstrmtAssignmentMethodField {
	return DerivativeInstrmtAssignmentMethodField{val}
}

//DerivativeInstrumentPartyIDField is a STRING field
type DerivativeInstrumentPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeInstrumentPartyID (1293)
func (f DerivativeInstrumentPartyIDField) Tag() quickfix.Tag { return tag.DerivativeInstrumentPartyID }

//NewDerivativeInstrumentPartyID returns a new DerivativeInstrumentPartyIDField initialized with val
func NewDerivativeInstrumentPartyID(val quickfix.FIXString) DerivativeInstrumentPartyIDField {
	return DerivativeInstrumentPartyIDField{val}
}

//DerivativeInstrumentPartyIDSourceField is a STRING field
type DerivativeInstrumentPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeInstrumentPartyIDSource (1294)
func (f DerivativeInstrumentPartyIDSourceField) Tag() quickfix.Tag {
	return tag.DerivativeInstrumentPartyIDSource
}

//NewDerivativeInstrumentPartyIDSource returns a new DerivativeInstrumentPartyIDSourceField initialized with val
func NewDerivativeInstrumentPartyIDSource(val quickfix.FIXString) DerivativeInstrumentPartyIDSourceField {
	return DerivativeInstrumentPartyIDSourceField{val}
}

//DerivativeInstrumentPartyRoleField is a INT field
type DerivativeInstrumentPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeInstrumentPartyRole (1295)
func (f DerivativeInstrumentPartyRoleField) Tag() quickfix.Tag {
	return tag.DerivativeInstrumentPartyRole
}

//NewDerivativeInstrumentPartyRole returns a new DerivativeInstrumentPartyRoleField initialized with val
func NewDerivativeInstrumentPartyRole(val quickfix.FIXInt) DerivativeInstrumentPartyRoleField {
	return DerivativeInstrumentPartyRoleField{val}
}

//DerivativeInstrumentPartySubIDField is a STRING field
type DerivativeInstrumentPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeInstrumentPartySubID (1297)
func (f DerivativeInstrumentPartySubIDField) Tag() quickfix.Tag {
	return tag.DerivativeInstrumentPartySubID
}

//NewDerivativeInstrumentPartySubID returns a new DerivativeInstrumentPartySubIDField initialized with val
func NewDerivativeInstrumentPartySubID(val quickfix.FIXString) DerivativeInstrumentPartySubIDField {
	return DerivativeInstrumentPartySubIDField{val}
}

//DerivativeInstrumentPartySubIDTypeField is a INT field
type DerivativeInstrumentPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeInstrumentPartySubIDType (1298)
func (f DerivativeInstrumentPartySubIDTypeField) Tag() quickfix.Tag {
	return tag.DerivativeInstrumentPartySubIDType
}

//NewDerivativeInstrumentPartySubIDType returns a new DerivativeInstrumentPartySubIDTypeField initialized with val
func NewDerivativeInstrumentPartySubIDType(val quickfix.FIXInt) DerivativeInstrumentPartySubIDTypeField {
	return DerivativeInstrumentPartySubIDTypeField{val}
}

//DerivativeIssueDateField is a LOCALMKTDATE field
type DerivativeIssueDateField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeIssueDate (1276)
func (f DerivativeIssueDateField) Tag() quickfix.Tag { return tag.DerivativeIssueDate }

//NewDerivativeIssueDate returns a new DerivativeIssueDateField initialized with val
func NewDerivativeIssueDate(val quickfix.FIXString) DerivativeIssueDateField {
	return DerivativeIssueDateField{val}
}

//DerivativeIssuerField is a STRING field
type DerivativeIssuerField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeIssuer (1275)
func (f DerivativeIssuerField) Tag() quickfix.Tag { return tag.DerivativeIssuer }

//NewDerivativeIssuer returns a new DerivativeIssuerField initialized with val
func NewDerivativeIssuer(val quickfix.FIXString) DerivativeIssuerField {
	return DerivativeIssuerField{val}
}

//DerivativeListMethodField is a INT field
type DerivativeListMethodField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeListMethod (1320)
func (f DerivativeListMethodField) Tag() quickfix.Tag { return tag.DerivativeListMethod }

//NewDerivativeListMethod returns a new DerivativeListMethodField initialized with val
func NewDerivativeListMethod(val quickfix.FIXInt) DerivativeListMethodField {
	return DerivativeListMethodField{val}
}

//DerivativeLocaleOfIssueField is a STRING field
type DerivativeLocaleOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeLocaleOfIssue (1260)
func (f DerivativeLocaleOfIssueField) Tag() quickfix.Tag { return tag.DerivativeLocaleOfIssue }

//NewDerivativeLocaleOfIssue returns a new DerivativeLocaleOfIssueField initialized with val
func NewDerivativeLocaleOfIssue(val quickfix.FIXString) DerivativeLocaleOfIssueField {
	return DerivativeLocaleOfIssueField{val}
}

//DerivativeMaturityDateField is a LOCALMKTDATE field
type DerivativeMaturityDateField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeMaturityDate (1252)
func (f DerivativeMaturityDateField) Tag() quickfix.Tag { return tag.DerivativeMaturityDate }

//NewDerivativeMaturityDate returns a new DerivativeMaturityDateField initialized with val
func NewDerivativeMaturityDate(val quickfix.FIXString) DerivativeMaturityDateField {
	return DerivativeMaturityDateField{val}
}

//DerivativeMaturityMonthYearField is a MONTHYEAR field
type DerivativeMaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeMaturityMonthYear (1251)
func (f DerivativeMaturityMonthYearField) Tag() quickfix.Tag { return tag.DerivativeMaturityMonthYear }

//NewDerivativeMaturityMonthYear returns a new DerivativeMaturityMonthYearField initialized with val
func NewDerivativeMaturityMonthYear(val quickfix.FIXString) DerivativeMaturityMonthYearField {
	return DerivativeMaturityMonthYearField{val}
}

//DerivativeMaturityTimeField is a TZTIMEONLY field
type DerivativeMaturityTimeField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeMaturityTime (1253)
func (f DerivativeMaturityTimeField) Tag() quickfix.Tag { return tag.DerivativeMaturityTime }

//NewDerivativeMaturityTime returns a new DerivativeMaturityTimeField initialized with val
func NewDerivativeMaturityTime(val quickfix.FIXString) DerivativeMaturityTimeField {
	return DerivativeMaturityTimeField{val}
}

//DerivativeMinPriceIncrementField is a FLOAT field
type DerivativeMinPriceIncrementField struct{ quickfix.FIXFloat }

//Tag returns tag.DerivativeMinPriceIncrement (1267)
func (f DerivativeMinPriceIncrementField) Tag() quickfix.Tag { return tag.DerivativeMinPriceIncrement }

//NewDerivativeMinPriceIncrement returns a new DerivativeMinPriceIncrementField initialized with val
func NewDerivativeMinPriceIncrement(val quickfix.FIXFloat) DerivativeMinPriceIncrementField {
	return DerivativeMinPriceIncrementField{val}
}

//DerivativeMinPriceIncrementAmountField is a AMT field
type DerivativeMinPriceIncrementAmountField struct{ quickfix.FIXFloat }

//Tag returns tag.DerivativeMinPriceIncrementAmount (1268)
func (f DerivativeMinPriceIncrementAmountField) Tag() quickfix.Tag {
	return tag.DerivativeMinPriceIncrementAmount
}

//NewDerivativeMinPriceIncrementAmount returns a new DerivativeMinPriceIncrementAmountField initialized with val
func NewDerivativeMinPriceIncrementAmount(val quickfix.FIXFloat) DerivativeMinPriceIncrementAmountField {
	return DerivativeMinPriceIncrementAmountField{val}
}

//DerivativeNTPositionLimitField is a INT field
type DerivativeNTPositionLimitField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeNTPositionLimit (1274)
func (f DerivativeNTPositionLimitField) Tag() quickfix.Tag { return tag.DerivativeNTPositionLimit }

//NewDerivativeNTPositionLimit returns a new DerivativeNTPositionLimitField initialized with val
func NewDerivativeNTPositionLimit(val quickfix.FIXInt) DerivativeNTPositionLimitField {
	return DerivativeNTPositionLimitField{val}
}

//DerivativeOptAttributeField is a CHAR field
type DerivativeOptAttributeField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeOptAttribute (1265)
func (f DerivativeOptAttributeField) Tag() quickfix.Tag { return tag.DerivativeOptAttribute }

//NewDerivativeOptAttribute returns a new DerivativeOptAttributeField initialized with val
func NewDerivativeOptAttribute(val quickfix.FIXString) DerivativeOptAttributeField {
	return DerivativeOptAttributeField{val}
}

//DerivativeOptPayAmountField is a AMT field
type DerivativeOptPayAmountField struct{ quickfix.FIXFloat }

//Tag returns tag.DerivativeOptPayAmount (1225)
func (f DerivativeOptPayAmountField) Tag() quickfix.Tag { return tag.DerivativeOptPayAmount }

//NewDerivativeOptPayAmount returns a new DerivativeOptPayAmountField initialized with val
func NewDerivativeOptPayAmount(val quickfix.FIXFloat) DerivativeOptPayAmountField {
	return DerivativeOptPayAmountField{val}
}

//DerivativePositionLimitField is a INT field
type DerivativePositionLimitField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativePositionLimit (1273)
func (f DerivativePositionLimitField) Tag() quickfix.Tag { return tag.DerivativePositionLimit }

//NewDerivativePositionLimit returns a new DerivativePositionLimitField initialized with val
func NewDerivativePositionLimit(val quickfix.FIXInt) DerivativePositionLimitField {
	return DerivativePositionLimitField{val}
}

//DerivativePriceQuoteMethodField is a STRING field
type DerivativePriceQuoteMethodField struct{ quickfix.FIXString }

//Tag returns tag.DerivativePriceQuoteMethod (1318)
func (f DerivativePriceQuoteMethodField) Tag() quickfix.Tag { return tag.DerivativePriceQuoteMethod }

//NewDerivativePriceQuoteMethod returns a new DerivativePriceQuoteMethodField initialized with val
func NewDerivativePriceQuoteMethod(val quickfix.FIXString) DerivativePriceQuoteMethodField {
	return DerivativePriceQuoteMethodField{val}
}

//DerivativePriceUnitOfMeasureField is a STRING field
type DerivativePriceUnitOfMeasureField struct{ quickfix.FIXString }

//Tag returns tag.DerivativePriceUnitOfMeasure (1315)
func (f DerivativePriceUnitOfMeasureField) Tag() quickfix.Tag { return tag.DerivativePriceUnitOfMeasure }

//NewDerivativePriceUnitOfMeasure returns a new DerivativePriceUnitOfMeasureField initialized with val
func NewDerivativePriceUnitOfMeasure(val quickfix.FIXString) DerivativePriceUnitOfMeasureField {
	return DerivativePriceUnitOfMeasureField{val}
}

//DerivativePriceUnitOfMeasureQtyField is a QTY field
type DerivativePriceUnitOfMeasureQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.DerivativePriceUnitOfMeasureQty (1316)
func (f DerivativePriceUnitOfMeasureQtyField) Tag() quickfix.Tag {
	return tag.DerivativePriceUnitOfMeasureQty
}

//NewDerivativePriceUnitOfMeasureQty returns a new DerivativePriceUnitOfMeasureQtyField initialized with val
func NewDerivativePriceUnitOfMeasureQty(val quickfix.FIXFloat) DerivativePriceUnitOfMeasureQtyField {
	return DerivativePriceUnitOfMeasureQtyField{val}
}

//DerivativeProductField is a INT field
type DerivativeProductField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeProduct (1246)
func (f DerivativeProductField) Tag() quickfix.Tag { return tag.DerivativeProduct }

//NewDerivativeProduct returns a new DerivativeProductField initialized with val
func NewDerivativeProduct(val quickfix.FIXInt) DerivativeProductField {
	return DerivativeProductField{val}
}

//DerivativeProductComplexField is a STRING field
type DerivativeProductComplexField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeProductComplex (1228)
func (f DerivativeProductComplexField) Tag() quickfix.Tag { return tag.DerivativeProductComplex }

//NewDerivativeProductComplex returns a new DerivativeProductComplexField initialized with val
func NewDerivativeProductComplex(val quickfix.FIXString) DerivativeProductComplexField {
	return DerivativeProductComplexField{val}
}

//DerivativePutOrCallField is a INT field
type DerivativePutOrCallField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativePutOrCall (1323)
func (f DerivativePutOrCallField) Tag() quickfix.Tag { return tag.DerivativePutOrCall }

//NewDerivativePutOrCall returns a new DerivativePutOrCallField initialized with val
func NewDerivativePutOrCall(val quickfix.FIXInt) DerivativePutOrCallField {
	return DerivativePutOrCallField{val}
}

//DerivativeSecurityAltIDField is a STRING field
type DerivativeSecurityAltIDField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityAltID (1219)
func (f DerivativeSecurityAltIDField) Tag() quickfix.Tag { return tag.DerivativeSecurityAltID }

//NewDerivativeSecurityAltID returns a new DerivativeSecurityAltIDField initialized with val
func NewDerivativeSecurityAltID(val quickfix.FIXString) DerivativeSecurityAltIDField {
	return DerivativeSecurityAltIDField{val}
}

//DerivativeSecurityAltIDSourceField is a STRING field
type DerivativeSecurityAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityAltIDSource (1220)
func (f DerivativeSecurityAltIDSourceField) Tag() quickfix.Tag {
	return tag.DerivativeSecurityAltIDSource
}

//NewDerivativeSecurityAltIDSource returns a new DerivativeSecurityAltIDSourceField initialized with val
func NewDerivativeSecurityAltIDSource(val quickfix.FIXString) DerivativeSecurityAltIDSourceField {
	return DerivativeSecurityAltIDSourceField{val}
}

//DerivativeSecurityDescField is a STRING field
type DerivativeSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityDesc (1279)
func (f DerivativeSecurityDescField) Tag() quickfix.Tag { return tag.DerivativeSecurityDesc }

//NewDerivativeSecurityDesc returns a new DerivativeSecurityDescField initialized with val
func NewDerivativeSecurityDesc(val quickfix.FIXString) DerivativeSecurityDescField {
	return DerivativeSecurityDescField{val}
}

//DerivativeSecurityExchangeField is a EXCHANGE field
type DerivativeSecurityExchangeField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityExchange (1272)
func (f DerivativeSecurityExchangeField) Tag() quickfix.Tag { return tag.DerivativeSecurityExchange }

//NewDerivativeSecurityExchange returns a new DerivativeSecurityExchangeField initialized with val
func NewDerivativeSecurityExchange(val quickfix.FIXString) DerivativeSecurityExchangeField {
	return DerivativeSecurityExchangeField{val}
}

//DerivativeSecurityGroupField is a STRING field
type DerivativeSecurityGroupField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityGroup (1247)
func (f DerivativeSecurityGroupField) Tag() quickfix.Tag { return tag.DerivativeSecurityGroup }

//NewDerivativeSecurityGroup returns a new DerivativeSecurityGroupField initialized with val
func NewDerivativeSecurityGroup(val quickfix.FIXString) DerivativeSecurityGroupField {
	return DerivativeSecurityGroupField{val}
}

//DerivativeSecurityIDField is a STRING field
type DerivativeSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityID (1216)
func (f DerivativeSecurityIDField) Tag() quickfix.Tag { return tag.DerivativeSecurityID }

//NewDerivativeSecurityID returns a new DerivativeSecurityIDField initialized with val
func NewDerivativeSecurityID(val quickfix.FIXString) DerivativeSecurityIDField {
	return DerivativeSecurityIDField{val}
}

//DerivativeSecurityIDSourceField is a STRING field
type DerivativeSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityIDSource (1217)
func (f DerivativeSecurityIDSourceField) Tag() quickfix.Tag { return tag.DerivativeSecurityIDSource }

//NewDerivativeSecurityIDSource returns a new DerivativeSecurityIDSourceField initialized with val
func NewDerivativeSecurityIDSource(val quickfix.FIXString) DerivativeSecurityIDSourceField {
	return DerivativeSecurityIDSourceField{val}
}

//DerivativeSecurityListRequestTypeField is a INT field
type DerivativeSecurityListRequestTypeField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeSecurityListRequestType (1307)
func (f DerivativeSecurityListRequestTypeField) Tag() quickfix.Tag {
	return tag.DerivativeSecurityListRequestType
}

//NewDerivativeSecurityListRequestType returns a new DerivativeSecurityListRequestTypeField initialized with val
func NewDerivativeSecurityListRequestType(val quickfix.FIXInt) DerivativeSecurityListRequestTypeField {
	return DerivativeSecurityListRequestTypeField{val}
}

//DerivativeSecurityStatusField is a STRING field
type DerivativeSecurityStatusField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityStatus (1256)
func (f DerivativeSecurityStatusField) Tag() quickfix.Tag { return tag.DerivativeSecurityStatus }

//NewDerivativeSecurityStatus returns a new DerivativeSecurityStatusField initialized with val
func NewDerivativeSecurityStatus(val quickfix.FIXString) DerivativeSecurityStatusField {
	return DerivativeSecurityStatusField{val}
}

//DerivativeSecuritySubTypeField is a STRING field
type DerivativeSecuritySubTypeField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecuritySubType (1250)
func (f DerivativeSecuritySubTypeField) Tag() quickfix.Tag { return tag.DerivativeSecuritySubType }

//NewDerivativeSecuritySubType returns a new DerivativeSecuritySubTypeField initialized with val
func NewDerivativeSecuritySubType(val quickfix.FIXString) DerivativeSecuritySubTypeField {
	return DerivativeSecuritySubTypeField{val}
}

//DerivativeSecurityTypeField is a STRING field
type DerivativeSecurityTypeField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityType (1249)
func (f DerivativeSecurityTypeField) Tag() quickfix.Tag { return tag.DerivativeSecurityType }

//NewDerivativeSecurityType returns a new DerivativeSecurityTypeField initialized with val
func NewDerivativeSecurityType(val quickfix.FIXString) DerivativeSecurityTypeField {
	return DerivativeSecurityTypeField{val}
}

//DerivativeSecurityXMLField is a DATA field
type DerivativeSecurityXMLField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityXML (1283)
func (f DerivativeSecurityXMLField) Tag() quickfix.Tag { return tag.DerivativeSecurityXML }

//NewDerivativeSecurityXML returns a new DerivativeSecurityXMLField initialized with val
func NewDerivativeSecurityXML(val quickfix.FIXString) DerivativeSecurityXMLField {
	return DerivativeSecurityXMLField{val}
}

//DerivativeSecurityXMLLenField is a LENGTH field
type DerivativeSecurityXMLLenField struct{ quickfix.FIXInt }

//Tag returns tag.DerivativeSecurityXMLLen (1282)
func (f DerivativeSecurityXMLLenField) Tag() quickfix.Tag { return tag.DerivativeSecurityXMLLen }

//NewDerivativeSecurityXMLLen returns a new DerivativeSecurityXMLLenField initialized with val
func NewDerivativeSecurityXMLLen(val quickfix.FIXInt) DerivativeSecurityXMLLenField {
	return DerivativeSecurityXMLLenField{val}
}

//DerivativeSecurityXMLSchemaField is a STRING field
type DerivativeSecurityXMLSchemaField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSecurityXMLSchema (1284)
func (f DerivativeSecurityXMLSchemaField) Tag() quickfix.Tag { return tag.DerivativeSecurityXMLSchema }

//NewDerivativeSecurityXMLSchema returns a new DerivativeSecurityXMLSchemaField initialized with val
func NewDerivativeSecurityXMLSchema(val quickfix.FIXString) DerivativeSecurityXMLSchemaField {
	return DerivativeSecurityXMLSchemaField{val}
}

//DerivativeSettlMethodField is a CHAR field
type DerivativeSettlMethodField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSettlMethod (1317)
func (f DerivativeSettlMethodField) Tag() quickfix.Tag { return tag.DerivativeSettlMethod }

//NewDerivativeSettlMethod returns a new DerivativeSettlMethodField initialized with val
func NewDerivativeSettlMethod(val quickfix.FIXString) DerivativeSettlMethodField {
	return DerivativeSettlMethodField{val}
}

//DerivativeSettleOnOpenFlagField is a STRING field
type DerivativeSettleOnOpenFlagField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSettleOnOpenFlag (1254)
func (f DerivativeSettleOnOpenFlagField) Tag() quickfix.Tag { return tag.DerivativeSettleOnOpenFlag }

//NewDerivativeSettleOnOpenFlag returns a new DerivativeSettleOnOpenFlagField initialized with val
func NewDerivativeSettleOnOpenFlag(val quickfix.FIXString) DerivativeSettleOnOpenFlagField {
	return DerivativeSettleOnOpenFlagField{val}
}

//DerivativeStateOrProvinceOfIssueField is a STRING field
type DerivativeStateOrProvinceOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeStateOrProvinceOfIssue (1259)
func (f DerivativeStateOrProvinceOfIssueField) Tag() quickfix.Tag {
	return tag.DerivativeStateOrProvinceOfIssue
}

//NewDerivativeStateOrProvinceOfIssue returns a new DerivativeStateOrProvinceOfIssueField initialized with val
func NewDerivativeStateOrProvinceOfIssue(val quickfix.FIXString) DerivativeStateOrProvinceOfIssueField {
	return DerivativeStateOrProvinceOfIssueField{val}
}

//DerivativeStrikeCurrencyField is a CURRENCY field
type DerivativeStrikeCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeStrikeCurrency (1262)
func (f DerivativeStrikeCurrencyField) Tag() quickfix.Tag { return tag.DerivativeStrikeCurrency }

//NewDerivativeStrikeCurrency returns a new DerivativeStrikeCurrencyField initialized with val
func NewDerivativeStrikeCurrency(val quickfix.FIXString) DerivativeStrikeCurrencyField {
	return DerivativeStrikeCurrencyField{val}
}

//DerivativeStrikeMultiplierField is a FLOAT field
type DerivativeStrikeMultiplierField struct{ quickfix.FIXFloat }

//Tag returns tag.DerivativeStrikeMultiplier (1263)
func (f DerivativeStrikeMultiplierField) Tag() quickfix.Tag { return tag.DerivativeStrikeMultiplier }

//NewDerivativeStrikeMultiplier returns a new DerivativeStrikeMultiplierField initialized with val
func NewDerivativeStrikeMultiplier(val quickfix.FIXFloat) DerivativeStrikeMultiplierField {
	return DerivativeStrikeMultiplierField{val}
}

//DerivativeStrikePriceField is a PRICE field
type DerivativeStrikePriceField struct{ quickfix.FIXFloat }

//Tag returns tag.DerivativeStrikePrice (1261)
func (f DerivativeStrikePriceField) Tag() quickfix.Tag { return tag.DerivativeStrikePrice }

//NewDerivativeStrikePrice returns a new DerivativeStrikePriceField initialized with val
func NewDerivativeStrikePrice(val quickfix.FIXFloat) DerivativeStrikePriceField {
	return DerivativeStrikePriceField{val}
}

//DerivativeStrikeValueField is a FLOAT field
type DerivativeStrikeValueField struct{ quickfix.FIXFloat }

//Tag returns tag.DerivativeStrikeValue (1264)
func (f DerivativeStrikeValueField) Tag() quickfix.Tag { return tag.DerivativeStrikeValue }

//NewDerivativeStrikeValue returns a new DerivativeStrikeValueField initialized with val
func NewDerivativeStrikeValue(val quickfix.FIXFloat) DerivativeStrikeValueField {
	return DerivativeStrikeValueField{val}
}

//DerivativeSymbolField is a STRING field
type DerivativeSymbolField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSymbol (1214)
func (f DerivativeSymbolField) Tag() quickfix.Tag { return tag.DerivativeSymbol }

//NewDerivativeSymbol returns a new DerivativeSymbolField initialized with val
func NewDerivativeSymbol(val quickfix.FIXString) DerivativeSymbolField {
	return DerivativeSymbolField{val}
}

//DerivativeSymbolSfxField is a STRING field
type DerivativeSymbolSfxField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeSymbolSfx (1215)
func (f DerivativeSymbolSfxField) Tag() quickfix.Tag { return tag.DerivativeSymbolSfx }

//NewDerivativeSymbolSfx returns a new DerivativeSymbolSfxField initialized with val
func NewDerivativeSymbolSfx(val quickfix.FIXString) DerivativeSymbolSfxField {
	return DerivativeSymbolSfxField{val}
}

//DerivativeTimeUnitField is a STRING field
type DerivativeTimeUnitField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeTimeUnit (1271)
func (f DerivativeTimeUnitField) Tag() quickfix.Tag { return tag.DerivativeTimeUnit }

//NewDerivativeTimeUnit returns a new DerivativeTimeUnitField initialized with val
func NewDerivativeTimeUnit(val quickfix.FIXString) DerivativeTimeUnitField {
	return DerivativeTimeUnitField{val}
}

//DerivativeUnitOfMeasureField is a STRING field
type DerivativeUnitOfMeasureField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeUnitOfMeasure (1269)
func (f DerivativeUnitOfMeasureField) Tag() quickfix.Tag { return tag.DerivativeUnitOfMeasure }

//NewDerivativeUnitOfMeasure returns a new DerivativeUnitOfMeasureField initialized with val
func NewDerivativeUnitOfMeasure(val quickfix.FIXString) DerivativeUnitOfMeasureField {
	return DerivativeUnitOfMeasureField{val}
}

//DerivativeUnitOfMeasureQtyField is a QTY field
type DerivativeUnitOfMeasureQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.DerivativeUnitOfMeasureQty (1270)
func (f DerivativeUnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.DerivativeUnitOfMeasureQty }

//NewDerivativeUnitOfMeasureQty returns a new DerivativeUnitOfMeasureQtyField initialized with val
func NewDerivativeUnitOfMeasureQty(val quickfix.FIXFloat) DerivativeUnitOfMeasureQtyField {
	return DerivativeUnitOfMeasureQtyField{val}
}

//DerivativeValuationMethodField is a STRING field
type DerivativeValuationMethodField struct{ quickfix.FIXString }

//Tag returns tag.DerivativeValuationMethod (1319)
func (f DerivativeValuationMethodField) Tag() quickfix.Tag { return tag.DerivativeValuationMethod }

//NewDerivativeValuationMethod returns a new DerivativeValuationMethodField initialized with val
func NewDerivativeValuationMethod(val quickfix.FIXString) DerivativeValuationMethodField {
	return DerivativeValuationMethodField{val}
}

//DesignationField is a STRING field
type DesignationField struct{ quickfix.FIXString }

//Tag returns tag.Designation (494)
func (f DesignationField) Tag() quickfix.Tag { return tag.Designation }

//NewDesignation returns a new DesignationField initialized with val
func NewDesignation(val quickfix.FIXString) DesignationField {
	return DesignationField{val}
}

//DeskIDField is a STRING field
type DeskIDField struct{ quickfix.FIXString }

//Tag returns tag.DeskID (284)
func (f DeskIDField) Tag() quickfix.Tag { return tag.DeskID }

//NewDeskID returns a new DeskIDField initialized with val
func NewDeskID(val quickfix.FIXString) DeskIDField {
	return DeskIDField{val}
}

//DeskOrderHandlingInstField is a MULTIPLESTRINGVALUE field
type DeskOrderHandlingInstField struct{ quickfix.FIXString }

//Tag returns tag.DeskOrderHandlingInst (1035)
func (f DeskOrderHandlingInstField) Tag() quickfix.Tag { return tag.DeskOrderHandlingInst }

//NewDeskOrderHandlingInst returns a new DeskOrderHandlingInstField initialized with val
func NewDeskOrderHandlingInst(val quickfix.FIXString) DeskOrderHandlingInstField {
	return DeskOrderHandlingInstField{val}
}

//DeskTypeField is a STRING field
type DeskTypeField struct{ quickfix.FIXString }

//Tag returns tag.DeskType (1033)
func (f DeskTypeField) Tag() quickfix.Tag { return tag.DeskType }

//NewDeskType returns a new DeskTypeField initialized with val
func NewDeskType(val quickfix.FIXString) DeskTypeField {
	return DeskTypeField{val}
}

//DeskTypeSourceField is a INT field
type DeskTypeSourceField struct{ quickfix.FIXInt }

//Tag returns tag.DeskTypeSource (1034)
func (f DeskTypeSourceField) Tag() quickfix.Tag { return tag.DeskTypeSource }

//NewDeskTypeSource returns a new DeskTypeSourceField initialized with val
func NewDeskTypeSource(val quickfix.FIXInt) DeskTypeSourceField {
	return DeskTypeSourceField{val}
}

//DetachmentPointField is a PERCENTAGE field
type DetachmentPointField struct{ quickfix.FIXFloat }

//Tag returns tag.DetachmentPoint (1458)
func (f DetachmentPointField) Tag() quickfix.Tag { return tag.DetachmentPoint }

//NewDetachmentPoint returns a new DetachmentPointField initialized with val
func NewDetachmentPoint(val quickfix.FIXFloat) DetachmentPointField {
	return DetachmentPointField{val}
}

//DiscretionInstField is a CHAR field
type DiscretionInstField struct{ quickfix.FIXString }

//Tag returns tag.DiscretionInst (388)
func (f DiscretionInstField) Tag() quickfix.Tag { return tag.DiscretionInst }

//NewDiscretionInst returns a new DiscretionInstField initialized with val
func NewDiscretionInst(val quickfix.FIXString) DiscretionInstField {
	return DiscretionInstField{val}
}

//DiscretionLimitTypeField is a INT field
type DiscretionLimitTypeField struct{ quickfix.FIXInt }

//Tag returns tag.DiscretionLimitType (843)
func (f DiscretionLimitTypeField) Tag() quickfix.Tag { return tag.DiscretionLimitType }

//NewDiscretionLimitType returns a new DiscretionLimitTypeField initialized with val
func NewDiscretionLimitType(val quickfix.FIXInt) DiscretionLimitTypeField {
	return DiscretionLimitTypeField{val}
}

//DiscretionMoveTypeField is a INT field
type DiscretionMoveTypeField struct{ quickfix.FIXInt }

//Tag returns tag.DiscretionMoveType (841)
func (f DiscretionMoveTypeField) Tag() quickfix.Tag { return tag.DiscretionMoveType }

//NewDiscretionMoveType returns a new DiscretionMoveTypeField initialized with val
func NewDiscretionMoveType(val quickfix.FIXInt) DiscretionMoveTypeField {
	return DiscretionMoveTypeField{val}
}

//DiscretionOffsetField is a PRICEOFFSET field
type DiscretionOffsetField struct{ quickfix.FIXFloat }

//Tag returns tag.DiscretionOffset (389)
func (f DiscretionOffsetField) Tag() quickfix.Tag { return tag.DiscretionOffset }

//NewDiscretionOffset returns a new DiscretionOffsetField initialized with val
func NewDiscretionOffset(val quickfix.FIXFloat) DiscretionOffsetField {
	return DiscretionOffsetField{val}
}

//DiscretionOffsetTypeField is a INT field
type DiscretionOffsetTypeField struct{ quickfix.FIXInt }

//Tag returns tag.DiscretionOffsetType (842)
func (f DiscretionOffsetTypeField) Tag() quickfix.Tag { return tag.DiscretionOffsetType }

//NewDiscretionOffsetType returns a new DiscretionOffsetTypeField initialized with val
func NewDiscretionOffsetType(val quickfix.FIXInt) DiscretionOffsetTypeField {
	return DiscretionOffsetTypeField{val}
}

//DiscretionOffsetValueField is a FLOAT field
type DiscretionOffsetValueField struct{ quickfix.FIXFloat }

//Tag returns tag.DiscretionOffsetValue (389)
func (f DiscretionOffsetValueField) Tag() quickfix.Tag { return tag.DiscretionOffsetValue }

//NewDiscretionOffsetValue returns a new DiscretionOffsetValueField initialized with val
func NewDiscretionOffsetValue(val quickfix.FIXFloat) DiscretionOffsetValueField {
	return DiscretionOffsetValueField{val}
}

//DiscretionPriceField is a PRICE field
type DiscretionPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.DiscretionPrice (845)
func (f DiscretionPriceField) Tag() quickfix.Tag { return tag.DiscretionPrice }

//NewDiscretionPrice returns a new DiscretionPriceField initialized with val
func NewDiscretionPrice(val quickfix.FIXFloat) DiscretionPriceField {
	return DiscretionPriceField{val}
}

//DiscretionRoundDirectionField is a INT field
type DiscretionRoundDirectionField struct{ quickfix.FIXInt }

//Tag returns tag.DiscretionRoundDirection (844)
func (f DiscretionRoundDirectionField) Tag() quickfix.Tag { return tag.DiscretionRoundDirection }

//NewDiscretionRoundDirection returns a new DiscretionRoundDirectionField initialized with val
func NewDiscretionRoundDirection(val quickfix.FIXInt) DiscretionRoundDirectionField {
	return DiscretionRoundDirectionField{val}
}

//DiscretionScopeField is a INT field
type DiscretionScopeField struct{ quickfix.FIXInt }

//Tag returns tag.DiscretionScope (846)
func (f DiscretionScopeField) Tag() quickfix.Tag { return tag.DiscretionScope }

//NewDiscretionScope returns a new DiscretionScopeField initialized with val
func NewDiscretionScope(val quickfix.FIXInt) DiscretionScopeField {
	return DiscretionScopeField{val}
}

//DisplayHighQtyField is a QTY field
type DisplayHighQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.DisplayHighQty (1086)
func (f DisplayHighQtyField) Tag() quickfix.Tag { return tag.DisplayHighQty }

//NewDisplayHighQty returns a new DisplayHighQtyField initialized with val
func NewDisplayHighQty(val quickfix.FIXFloat) DisplayHighQtyField {
	return DisplayHighQtyField{val}
}

//DisplayLowQtyField is a QTY field
type DisplayLowQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.DisplayLowQty (1085)
func (f DisplayLowQtyField) Tag() quickfix.Tag { return tag.DisplayLowQty }

//NewDisplayLowQty returns a new DisplayLowQtyField initialized with val
func NewDisplayLowQty(val quickfix.FIXFloat) DisplayLowQtyField {
	return DisplayLowQtyField{val}
}

//DisplayMethodField is a CHAR field
type DisplayMethodField struct{ quickfix.FIXString }

//Tag returns tag.DisplayMethod (1084)
func (f DisplayMethodField) Tag() quickfix.Tag { return tag.DisplayMethod }

//NewDisplayMethod returns a new DisplayMethodField initialized with val
func NewDisplayMethod(val quickfix.FIXString) DisplayMethodField {
	return DisplayMethodField{val}
}

//DisplayMinIncrField is a QTY field
type DisplayMinIncrField struct{ quickfix.FIXFloat }

//Tag returns tag.DisplayMinIncr (1087)
func (f DisplayMinIncrField) Tag() quickfix.Tag { return tag.DisplayMinIncr }

//NewDisplayMinIncr returns a new DisplayMinIncrField initialized with val
func NewDisplayMinIncr(val quickfix.FIXFloat) DisplayMinIncrField {
	return DisplayMinIncrField{val}
}

//DisplayQtyField is a QTY field
type DisplayQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.DisplayQty (1138)
func (f DisplayQtyField) Tag() quickfix.Tag { return tag.DisplayQty }

//NewDisplayQty returns a new DisplayQtyField initialized with val
func NewDisplayQty(val quickfix.FIXFloat) DisplayQtyField {
	return DisplayQtyField{val}
}

//DisplayWhenField is a CHAR field
type DisplayWhenField struct{ quickfix.FIXString }

//Tag returns tag.DisplayWhen (1083)
func (f DisplayWhenField) Tag() quickfix.Tag { return tag.DisplayWhen }

//NewDisplayWhen returns a new DisplayWhenField initialized with val
func NewDisplayWhen(val quickfix.FIXString) DisplayWhenField {
	return DisplayWhenField{val}
}

//DistribPaymentMethodField is a INT field
type DistribPaymentMethodField struct{ quickfix.FIXInt }

//Tag returns tag.DistribPaymentMethod (477)
func (f DistribPaymentMethodField) Tag() quickfix.Tag { return tag.DistribPaymentMethod }

//NewDistribPaymentMethod returns a new DistribPaymentMethodField initialized with val
func NewDistribPaymentMethod(val quickfix.FIXInt) DistribPaymentMethodField {
	return DistribPaymentMethodField{val}
}

//DistribPercentageField is a PERCENTAGE field
type DistribPercentageField struct{ quickfix.FIXFloat }

//Tag returns tag.DistribPercentage (512)
func (f DistribPercentageField) Tag() quickfix.Tag { return tag.DistribPercentage }

//NewDistribPercentage returns a new DistribPercentageField initialized with val
func NewDistribPercentage(val quickfix.FIXFloat) DistribPercentageField {
	return DistribPercentageField{val}
}

//DividendYieldField is a PERCENTAGE field
type DividendYieldField struct{ quickfix.FIXFloat }

//Tag returns tag.DividendYield (1380)
func (f DividendYieldField) Tag() quickfix.Tag { return tag.DividendYield }

//NewDividendYield returns a new DividendYieldField initialized with val
func NewDividendYield(val quickfix.FIXFloat) DividendYieldField {
	return DividendYieldField{val}
}

//DlvyInstField is a STRING field
type DlvyInstField struct{ quickfix.FIXString }

//Tag returns tag.DlvyInst (86)
func (f DlvyInstField) Tag() quickfix.Tag { return tag.DlvyInst }

//NewDlvyInst returns a new DlvyInstField initialized with val
func NewDlvyInst(val quickfix.FIXString) DlvyInstField {
	return DlvyInstField{val}
}

//DlvyInstTypeField is a CHAR field
type DlvyInstTypeField struct{ quickfix.FIXString }

//Tag returns tag.DlvyInstType (787)
func (f DlvyInstTypeField) Tag() quickfix.Tag { return tag.DlvyInstType }

//NewDlvyInstType returns a new DlvyInstTypeField initialized with val
func NewDlvyInstType(val quickfix.FIXString) DlvyInstTypeField {
	return DlvyInstTypeField{val}
}

//DueToRelatedField is a BOOLEAN field
type DueToRelatedField struct{ quickfix.FIXBoolean }

//Tag returns tag.DueToRelated (329)
func (f DueToRelatedField) Tag() quickfix.Tag { return tag.DueToRelated }

//NewDueToRelated returns a new DueToRelatedField initialized with val
func NewDueToRelated(val quickfix.FIXBoolean) DueToRelatedField {
	return DueToRelatedField{val}
}

//EFPTrackingErrorField is a PERCENTAGE field
type EFPTrackingErrorField struct{ quickfix.FIXFloat }

//Tag returns tag.EFPTrackingError (405)
func (f EFPTrackingErrorField) Tag() quickfix.Tag { return tag.EFPTrackingError }

//NewEFPTrackingError returns a new EFPTrackingErrorField initialized with val
func NewEFPTrackingError(val quickfix.FIXFloat) EFPTrackingErrorField {
	return EFPTrackingErrorField{val}
}

//EffectiveTimeField is a UTCTIMESTAMP field
type EffectiveTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.EffectiveTime (168)
func (f EffectiveTimeField) Tag() quickfix.Tag { return tag.EffectiveTime }

//NewEffectiveTime returns a new EffectiveTimeField initialized with val
func NewEffectiveTime(val time.Time) EffectiveTimeField {
	return EffectiveTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewEffectiveTimeNoMillis returns a new EffectiveTimeField initialized with val without millisecs
func NewEffectiveTimeNoMillis(val time.Time) EffectiveTimeField {
	return EffectiveTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//EmailThreadIDField is a STRING field
type EmailThreadIDField struct{ quickfix.FIXString }

//Tag returns tag.EmailThreadID (164)
func (f EmailThreadIDField) Tag() quickfix.Tag { return tag.EmailThreadID }

//NewEmailThreadID returns a new EmailThreadIDField initialized with val
func NewEmailThreadID(val quickfix.FIXString) EmailThreadIDField {
	return EmailThreadIDField{val}
}

//EmailTypeField is a CHAR field
type EmailTypeField struct{ quickfix.FIXString }

//Tag returns tag.EmailType (94)
func (f EmailTypeField) Tag() quickfix.Tag { return tag.EmailType }

//NewEmailType returns a new EmailTypeField initialized with val
func NewEmailType(val quickfix.FIXString) EmailTypeField {
	return EmailTypeField{val}
}

//EncodedAllocTextField is a DATA field
type EncodedAllocTextField struct{ quickfix.FIXString }

//Tag returns tag.EncodedAllocText (361)
func (f EncodedAllocTextField) Tag() quickfix.Tag { return tag.EncodedAllocText }

//NewEncodedAllocText returns a new EncodedAllocTextField initialized with val
func NewEncodedAllocText(val quickfix.FIXString) EncodedAllocTextField {
	return EncodedAllocTextField{val}
}

//EncodedAllocTextLenField is a LENGTH field
type EncodedAllocTextLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedAllocTextLen (360)
func (f EncodedAllocTextLenField) Tag() quickfix.Tag { return tag.EncodedAllocTextLen }

//NewEncodedAllocTextLen returns a new EncodedAllocTextLenField initialized with val
func NewEncodedAllocTextLen(val quickfix.FIXInt) EncodedAllocTextLenField {
	return EncodedAllocTextLenField{val}
}

//EncodedHeadlineField is a DATA field
type EncodedHeadlineField struct{ quickfix.FIXString }

//Tag returns tag.EncodedHeadline (359)
func (f EncodedHeadlineField) Tag() quickfix.Tag { return tag.EncodedHeadline }

//NewEncodedHeadline returns a new EncodedHeadlineField initialized with val
func NewEncodedHeadline(val quickfix.FIXString) EncodedHeadlineField {
	return EncodedHeadlineField{val}
}

//EncodedHeadlineLenField is a LENGTH field
type EncodedHeadlineLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedHeadlineLen (358)
func (f EncodedHeadlineLenField) Tag() quickfix.Tag { return tag.EncodedHeadlineLen }

//NewEncodedHeadlineLen returns a new EncodedHeadlineLenField initialized with val
func NewEncodedHeadlineLen(val quickfix.FIXInt) EncodedHeadlineLenField {
	return EncodedHeadlineLenField{val}
}

//EncodedIssuerField is a DATA field
type EncodedIssuerField struct{ quickfix.FIXString }

//Tag returns tag.EncodedIssuer (349)
func (f EncodedIssuerField) Tag() quickfix.Tag { return tag.EncodedIssuer }

//NewEncodedIssuer returns a new EncodedIssuerField initialized with val
func NewEncodedIssuer(val quickfix.FIXString) EncodedIssuerField {
	return EncodedIssuerField{val}
}

//EncodedIssuerLenField is a LENGTH field
type EncodedIssuerLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedIssuerLen (348)
func (f EncodedIssuerLenField) Tag() quickfix.Tag { return tag.EncodedIssuerLen }

//NewEncodedIssuerLen returns a new EncodedIssuerLenField initialized with val
func NewEncodedIssuerLen(val quickfix.FIXInt) EncodedIssuerLenField {
	return EncodedIssuerLenField{val}
}

//EncodedLegIssuerField is a DATA field
type EncodedLegIssuerField struct{ quickfix.FIXString }

//Tag returns tag.EncodedLegIssuer (619)
func (f EncodedLegIssuerField) Tag() quickfix.Tag { return tag.EncodedLegIssuer }

//NewEncodedLegIssuer returns a new EncodedLegIssuerField initialized with val
func NewEncodedLegIssuer(val quickfix.FIXString) EncodedLegIssuerField {
	return EncodedLegIssuerField{val}
}

//EncodedLegIssuerLenField is a LENGTH field
type EncodedLegIssuerLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedLegIssuerLen (618)
func (f EncodedLegIssuerLenField) Tag() quickfix.Tag { return tag.EncodedLegIssuerLen }

//NewEncodedLegIssuerLen returns a new EncodedLegIssuerLenField initialized with val
func NewEncodedLegIssuerLen(val quickfix.FIXInt) EncodedLegIssuerLenField {
	return EncodedLegIssuerLenField{val}
}

//EncodedLegSecurityDescField is a DATA field
type EncodedLegSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.EncodedLegSecurityDesc (622)
func (f EncodedLegSecurityDescField) Tag() quickfix.Tag { return tag.EncodedLegSecurityDesc }

//NewEncodedLegSecurityDesc returns a new EncodedLegSecurityDescField initialized with val
func NewEncodedLegSecurityDesc(val quickfix.FIXString) EncodedLegSecurityDescField {
	return EncodedLegSecurityDescField{val}
}

//EncodedLegSecurityDescLenField is a LENGTH field
type EncodedLegSecurityDescLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedLegSecurityDescLen (621)
func (f EncodedLegSecurityDescLenField) Tag() quickfix.Tag { return tag.EncodedLegSecurityDescLen }

//NewEncodedLegSecurityDescLen returns a new EncodedLegSecurityDescLenField initialized with val
func NewEncodedLegSecurityDescLen(val quickfix.FIXInt) EncodedLegSecurityDescLenField {
	return EncodedLegSecurityDescLenField{val}
}

//EncodedListExecInstField is a DATA field
type EncodedListExecInstField struct{ quickfix.FIXString }

//Tag returns tag.EncodedListExecInst (353)
func (f EncodedListExecInstField) Tag() quickfix.Tag { return tag.EncodedListExecInst }

//NewEncodedListExecInst returns a new EncodedListExecInstField initialized with val
func NewEncodedListExecInst(val quickfix.FIXString) EncodedListExecInstField {
	return EncodedListExecInstField{val}
}

//EncodedListExecInstLenField is a LENGTH field
type EncodedListExecInstLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedListExecInstLen (352)
func (f EncodedListExecInstLenField) Tag() quickfix.Tag { return tag.EncodedListExecInstLen }

//NewEncodedListExecInstLen returns a new EncodedListExecInstLenField initialized with val
func NewEncodedListExecInstLen(val quickfix.FIXInt) EncodedListExecInstLenField {
	return EncodedListExecInstLenField{val}
}

//EncodedListStatusTextField is a DATA field
type EncodedListStatusTextField struct{ quickfix.FIXString }

//Tag returns tag.EncodedListStatusText (446)
func (f EncodedListStatusTextField) Tag() quickfix.Tag { return tag.EncodedListStatusText }

//NewEncodedListStatusText returns a new EncodedListStatusTextField initialized with val
func NewEncodedListStatusText(val quickfix.FIXString) EncodedListStatusTextField {
	return EncodedListStatusTextField{val}
}

//EncodedListStatusTextLenField is a LENGTH field
type EncodedListStatusTextLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedListStatusTextLen (445)
func (f EncodedListStatusTextLenField) Tag() quickfix.Tag { return tag.EncodedListStatusTextLen }

//NewEncodedListStatusTextLen returns a new EncodedListStatusTextLenField initialized with val
func NewEncodedListStatusTextLen(val quickfix.FIXInt) EncodedListStatusTextLenField {
	return EncodedListStatusTextLenField{val}
}

//EncodedMktSegmDescField is a DATA field
type EncodedMktSegmDescField struct{ quickfix.FIXString }

//Tag returns tag.EncodedMktSegmDesc (1398)
func (f EncodedMktSegmDescField) Tag() quickfix.Tag { return tag.EncodedMktSegmDesc }

//NewEncodedMktSegmDesc returns a new EncodedMktSegmDescField initialized with val
func NewEncodedMktSegmDesc(val quickfix.FIXString) EncodedMktSegmDescField {
	return EncodedMktSegmDescField{val}
}

//EncodedMktSegmDescLenField is a LENGTH field
type EncodedMktSegmDescLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedMktSegmDescLen (1397)
func (f EncodedMktSegmDescLenField) Tag() quickfix.Tag { return tag.EncodedMktSegmDescLen }

//NewEncodedMktSegmDescLen returns a new EncodedMktSegmDescLenField initialized with val
func NewEncodedMktSegmDescLen(val quickfix.FIXInt) EncodedMktSegmDescLenField {
	return EncodedMktSegmDescLenField{val}
}

//EncodedSecurityDescField is a DATA field
type EncodedSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.EncodedSecurityDesc (351)
func (f EncodedSecurityDescField) Tag() quickfix.Tag { return tag.EncodedSecurityDesc }

//NewEncodedSecurityDesc returns a new EncodedSecurityDescField initialized with val
func NewEncodedSecurityDesc(val quickfix.FIXString) EncodedSecurityDescField {
	return EncodedSecurityDescField{val}
}

//EncodedSecurityDescLenField is a LENGTH field
type EncodedSecurityDescLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedSecurityDescLen (350)
func (f EncodedSecurityDescLenField) Tag() quickfix.Tag { return tag.EncodedSecurityDescLen }

//NewEncodedSecurityDescLen returns a new EncodedSecurityDescLenField initialized with val
func NewEncodedSecurityDescLen(val quickfix.FIXInt) EncodedSecurityDescLenField {
	return EncodedSecurityDescLenField{val}
}

//EncodedSecurityListDescField is a DATA field
type EncodedSecurityListDescField struct{ quickfix.FIXString }

//Tag returns tag.EncodedSecurityListDesc (1469)
func (f EncodedSecurityListDescField) Tag() quickfix.Tag { return tag.EncodedSecurityListDesc }

//NewEncodedSecurityListDesc returns a new EncodedSecurityListDescField initialized with val
func NewEncodedSecurityListDesc(val quickfix.FIXString) EncodedSecurityListDescField {
	return EncodedSecurityListDescField{val}
}

//EncodedSecurityListDescLenField is a LENGTH field
type EncodedSecurityListDescLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedSecurityListDescLen (1468)
func (f EncodedSecurityListDescLenField) Tag() quickfix.Tag { return tag.EncodedSecurityListDescLen }

//NewEncodedSecurityListDescLen returns a new EncodedSecurityListDescLenField initialized with val
func NewEncodedSecurityListDescLen(val quickfix.FIXInt) EncodedSecurityListDescLenField {
	return EncodedSecurityListDescLenField{val}
}

//EncodedSubjectField is a DATA field
type EncodedSubjectField struct{ quickfix.FIXString }

//Tag returns tag.EncodedSubject (357)
func (f EncodedSubjectField) Tag() quickfix.Tag { return tag.EncodedSubject }

//NewEncodedSubject returns a new EncodedSubjectField initialized with val
func NewEncodedSubject(val quickfix.FIXString) EncodedSubjectField {
	return EncodedSubjectField{val}
}

//EncodedSubjectLenField is a LENGTH field
type EncodedSubjectLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedSubjectLen (356)
func (f EncodedSubjectLenField) Tag() quickfix.Tag { return tag.EncodedSubjectLen }

//NewEncodedSubjectLen returns a new EncodedSubjectLenField initialized with val
func NewEncodedSubjectLen(val quickfix.FIXInt) EncodedSubjectLenField {
	return EncodedSubjectLenField{val}
}

//EncodedSymbolField is a DATA field
type EncodedSymbolField struct{ quickfix.FIXString }

//Tag returns tag.EncodedSymbol (1360)
func (f EncodedSymbolField) Tag() quickfix.Tag { return tag.EncodedSymbol }

//NewEncodedSymbol returns a new EncodedSymbolField initialized with val
func NewEncodedSymbol(val quickfix.FIXString) EncodedSymbolField {
	return EncodedSymbolField{val}
}

//EncodedSymbolLenField is a LENGTH field
type EncodedSymbolLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedSymbolLen (1359)
func (f EncodedSymbolLenField) Tag() quickfix.Tag { return tag.EncodedSymbolLen }

//NewEncodedSymbolLen returns a new EncodedSymbolLenField initialized with val
func NewEncodedSymbolLen(val quickfix.FIXInt) EncodedSymbolLenField {
	return EncodedSymbolLenField{val}
}

//EncodedTextField is a DATA field
type EncodedTextField struct{ quickfix.FIXString }

//Tag returns tag.EncodedText (355)
func (f EncodedTextField) Tag() quickfix.Tag { return tag.EncodedText }

//NewEncodedText returns a new EncodedTextField initialized with val
func NewEncodedText(val quickfix.FIXString) EncodedTextField {
	return EncodedTextField{val}
}

//EncodedTextLenField is a LENGTH field
type EncodedTextLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedTextLen (354)
func (f EncodedTextLenField) Tag() quickfix.Tag { return tag.EncodedTextLen }

//NewEncodedTextLen returns a new EncodedTextLenField initialized with val
func NewEncodedTextLen(val quickfix.FIXInt) EncodedTextLenField {
	return EncodedTextLenField{val}
}

//EncodedUnderlyingIssuerField is a DATA field
type EncodedUnderlyingIssuerField struct{ quickfix.FIXString }

//Tag returns tag.EncodedUnderlyingIssuer (363)
func (f EncodedUnderlyingIssuerField) Tag() quickfix.Tag { return tag.EncodedUnderlyingIssuer }

//NewEncodedUnderlyingIssuer returns a new EncodedUnderlyingIssuerField initialized with val
func NewEncodedUnderlyingIssuer(val quickfix.FIXString) EncodedUnderlyingIssuerField {
	return EncodedUnderlyingIssuerField{val}
}

//EncodedUnderlyingIssuerLenField is a LENGTH field
type EncodedUnderlyingIssuerLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedUnderlyingIssuerLen (362)
func (f EncodedUnderlyingIssuerLenField) Tag() quickfix.Tag { return tag.EncodedUnderlyingIssuerLen }

//NewEncodedUnderlyingIssuerLen returns a new EncodedUnderlyingIssuerLenField initialized with val
func NewEncodedUnderlyingIssuerLen(val quickfix.FIXInt) EncodedUnderlyingIssuerLenField {
	return EncodedUnderlyingIssuerLenField{val}
}

//EncodedUnderlyingSecurityDescField is a DATA field
type EncodedUnderlyingSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.EncodedUnderlyingSecurityDesc (365)
func (f EncodedUnderlyingSecurityDescField) Tag() quickfix.Tag {
	return tag.EncodedUnderlyingSecurityDesc
}

//NewEncodedUnderlyingSecurityDesc returns a new EncodedUnderlyingSecurityDescField initialized with val
func NewEncodedUnderlyingSecurityDesc(val quickfix.FIXString) EncodedUnderlyingSecurityDescField {
	return EncodedUnderlyingSecurityDescField{val}
}

//EncodedUnderlyingSecurityDescLenField is a LENGTH field
type EncodedUnderlyingSecurityDescLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncodedUnderlyingSecurityDescLen (364)
func (f EncodedUnderlyingSecurityDescLenField) Tag() quickfix.Tag {
	return tag.EncodedUnderlyingSecurityDescLen
}

//NewEncodedUnderlyingSecurityDescLen returns a new EncodedUnderlyingSecurityDescLenField initialized with val
func NewEncodedUnderlyingSecurityDescLen(val quickfix.FIXInt) EncodedUnderlyingSecurityDescLenField {
	return EncodedUnderlyingSecurityDescLenField{val}
}

//EncryptMethodField is a INT field
type EncryptMethodField struct{ quickfix.FIXInt }

//Tag returns tag.EncryptMethod (98)
func (f EncryptMethodField) Tag() quickfix.Tag { return tag.EncryptMethod }

//NewEncryptMethod returns a new EncryptMethodField initialized with val
func NewEncryptMethod(val quickfix.FIXInt) EncryptMethodField {
	return EncryptMethodField{val}
}

//EncryptedNewPasswordField is a DATA field
type EncryptedNewPasswordField struct{ quickfix.FIXString }

//Tag returns tag.EncryptedNewPassword (1404)
func (f EncryptedNewPasswordField) Tag() quickfix.Tag { return tag.EncryptedNewPassword }

//NewEncryptedNewPassword returns a new EncryptedNewPasswordField initialized with val
func NewEncryptedNewPassword(val quickfix.FIXString) EncryptedNewPasswordField {
	return EncryptedNewPasswordField{val}
}

//EncryptedNewPasswordLenField is a LENGTH field
type EncryptedNewPasswordLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncryptedNewPasswordLen (1403)
func (f EncryptedNewPasswordLenField) Tag() quickfix.Tag { return tag.EncryptedNewPasswordLen }

//NewEncryptedNewPasswordLen returns a new EncryptedNewPasswordLenField initialized with val
func NewEncryptedNewPasswordLen(val quickfix.FIXInt) EncryptedNewPasswordLenField {
	return EncryptedNewPasswordLenField{val}
}

//EncryptedPasswordField is a DATA field
type EncryptedPasswordField struct{ quickfix.FIXString }

//Tag returns tag.EncryptedPassword (1402)
func (f EncryptedPasswordField) Tag() quickfix.Tag { return tag.EncryptedPassword }

//NewEncryptedPassword returns a new EncryptedPasswordField initialized with val
func NewEncryptedPassword(val quickfix.FIXString) EncryptedPasswordField {
	return EncryptedPasswordField{val}
}

//EncryptedPasswordLenField is a LENGTH field
type EncryptedPasswordLenField struct{ quickfix.FIXInt }

//Tag returns tag.EncryptedPasswordLen (1401)
func (f EncryptedPasswordLenField) Tag() quickfix.Tag { return tag.EncryptedPasswordLen }

//NewEncryptedPasswordLen returns a new EncryptedPasswordLenField initialized with val
func NewEncryptedPasswordLen(val quickfix.FIXInt) EncryptedPasswordLenField {
	return EncryptedPasswordLenField{val}
}

//EncryptedPasswordMethodField is a INT field
type EncryptedPasswordMethodField struct{ quickfix.FIXInt }

//Tag returns tag.EncryptedPasswordMethod (1400)
func (f EncryptedPasswordMethodField) Tag() quickfix.Tag { return tag.EncryptedPasswordMethod }

//NewEncryptedPasswordMethod returns a new EncryptedPasswordMethodField initialized with val
func NewEncryptedPasswordMethod(val quickfix.FIXInt) EncryptedPasswordMethodField {
	return EncryptedPasswordMethodField{val}
}

//EndAccruedInterestAmtField is a AMT field
type EndAccruedInterestAmtField struct{ quickfix.FIXFloat }

//Tag returns tag.EndAccruedInterestAmt (920)
func (f EndAccruedInterestAmtField) Tag() quickfix.Tag { return tag.EndAccruedInterestAmt }

//NewEndAccruedInterestAmt returns a new EndAccruedInterestAmtField initialized with val
func NewEndAccruedInterestAmt(val quickfix.FIXFloat) EndAccruedInterestAmtField {
	return EndAccruedInterestAmtField{val}
}

//EndCashField is a AMT field
type EndCashField struct{ quickfix.FIXFloat }

//Tag returns tag.EndCash (922)
func (f EndCashField) Tag() quickfix.Tag { return tag.EndCash }

//NewEndCash returns a new EndCashField initialized with val
func NewEndCash(val quickfix.FIXFloat) EndCashField {
	return EndCashField{val}
}

//EndDateField is a LOCALMKTDATE field
type EndDateField struct{ quickfix.FIXString }

//Tag returns tag.EndDate (917)
func (f EndDateField) Tag() quickfix.Tag { return tag.EndDate }

//NewEndDate returns a new EndDateField initialized with val
func NewEndDate(val quickfix.FIXString) EndDateField {
	return EndDateField{val}
}

//EndMaturityMonthYearField is a MONTHYEAR field
type EndMaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.EndMaturityMonthYear (1226)
func (f EndMaturityMonthYearField) Tag() quickfix.Tag { return tag.EndMaturityMonthYear }

//NewEndMaturityMonthYear returns a new EndMaturityMonthYearField initialized with val
func NewEndMaturityMonthYear(val quickfix.FIXString) EndMaturityMonthYearField {
	return EndMaturityMonthYearField{val}
}

//EndSeqNoField is a SEQNUM field
type EndSeqNoField struct{ quickfix.FIXInt }

//Tag returns tag.EndSeqNo (16)
func (f EndSeqNoField) Tag() quickfix.Tag { return tag.EndSeqNo }

//NewEndSeqNo returns a new EndSeqNoField initialized with val
func NewEndSeqNo(val quickfix.FIXInt) EndSeqNoField {
	return EndSeqNoField{val}
}

//EndStrikePxRangeField is a PRICE field
type EndStrikePxRangeField struct{ quickfix.FIXFloat }

//Tag returns tag.EndStrikePxRange (1203)
func (f EndStrikePxRangeField) Tag() quickfix.Tag { return tag.EndStrikePxRange }

//NewEndStrikePxRange returns a new EndStrikePxRangeField initialized with val
func NewEndStrikePxRange(val quickfix.FIXFloat) EndStrikePxRangeField {
	return EndStrikePxRangeField{val}
}

//EndTickPriceRangeField is a PRICE field
type EndTickPriceRangeField struct{ quickfix.FIXFloat }

//Tag returns tag.EndTickPriceRange (1207)
func (f EndTickPriceRangeField) Tag() quickfix.Tag { return tag.EndTickPriceRange }

//NewEndTickPriceRange returns a new EndTickPriceRangeField initialized with val
func NewEndTickPriceRange(val quickfix.FIXFloat) EndTickPriceRangeField {
	return EndTickPriceRangeField{val}
}

//EventDateField is a LOCALMKTDATE field
type EventDateField struct{ quickfix.FIXString }

//Tag returns tag.EventDate (866)
func (f EventDateField) Tag() quickfix.Tag { return tag.EventDate }

//NewEventDate returns a new EventDateField initialized with val
func NewEventDate(val quickfix.FIXString) EventDateField {
	return EventDateField{val}
}

//EventPxField is a PRICE field
type EventPxField struct{ quickfix.FIXFloat }

//Tag returns tag.EventPx (867)
func (f EventPxField) Tag() quickfix.Tag { return tag.EventPx }

//NewEventPx returns a new EventPxField initialized with val
func NewEventPx(val quickfix.FIXFloat) EventPxField {
	return EventPxField{val}
}

//EventTextField is a STRING field
type EventTextField struct{ quickfix.FIXString }

//Tag returns tag.EventText (868)
func (f EventTextField) Tag() quickfix.Tag { return tag.EventText }

//NewEventText returns a new EventTextField initialized with val
func NewEventText(val quickfix.FIXString) EventTextField {
	return EventTextField{val}
}

//EventTimeField is a UTCTIMESTAMP field
type EventTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.EventTime (1145)
func (f EventTimeField) Tag() quickfix.Tag { return tag.EventTime }

//NewEventTime returns a new EventTimeField initialized with val
func NewEventTime(val time.Time) EventTimeField {
	return EventTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewEventTimeNoMillis returns a new EventTimeField initialized with val without millisecs
func NewEventTimeNoMillis(val time.Time) EventTimeField {
	return EventTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//EventTypeField is a INT field
type EventTypeField struct{ quickfix.FIXInt }

//Tag returns tag.EventType (865)
func (f EventTypeField) Tag() quickfix.Tag { return tag.EventType }

//NewEventType returns a new EventTypeField initialized with val
func NewEventType(val quickfix.FIXInt) EventTypeField {
	return EventTypeField{val}
}

//ExDateField is a LOCALMKTDATE field
type ExDateField struct{ quickfix.FIXString }

//Tag returns tag.ExDate (230)
func (f ExDateField) Tag() quickfix.Tag { return tag.ExDate }

//NewExDate returns a new ExDateField initialized with val
func NewExDate(val quickfix.FIXString) ExDateField {
	return ExDateField{val}
}

//ExDestinationField is a EXCHANGE field
type ExDestinationField struct{ quickfix.FIXString }

//Tag returns tag.ExDestination (100)
func (f ExDestinationField) Tag() quickfix.Tag { return tag.ExDestination }

//NewExDestination returns a new ExDestinationField initialized with val
func NewExDestination(val quickfix.FIXString) ExDestinationField {
	return ExDestinationField{val}
}

//ExDestinationIDSourceField is a CHAR field
type ExDestinationIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.ExDestinationIDSource (1133)
func (f ExDestinationIDSourceField) Tag() quickfix.Tag { return tag.ExDestinationIDSource }

//NewExDestinationIDSource returns a new ExDestinationIDSourceField initialized with val
func NewExDestinationIDSource(val quickfix.FIXString) ExDestinationIDSourceField {
	return ExDestinationIDSourceField{val}
}

//ExchangeForPhysicalField is a BOOLEAN field
type ExchangeForPhysicalField struct{ quickfix.FIXBoolean }

//Tag returns tag.ExchangeForPhysical (411)
func (f ExchangeForPhysicalField) Tag() quickfix.Tag { return tag.ExchangeForPhysical }

//NewExchangeForPhysical returns a new ExchangeForPhysicalField initialized with val
func NewExchangeForPhysical(val quickfix.FIXBoolean) ExchangeForPhysicalField {
	return ExchangeForPhysicalField{val}
}

//ExchangeRuleField is a STRING field
type ExchangeRuleField struct{ quickfix.FIXString }

//Tag returns tag.ExchangeRule (825)
func (f ExchangeRuleField) Tag() quickfix.Tag { return tag.ExchangeRule }

//NewExchangeRule returns a new ExchangeRuleField initialized with val
func NewExchangeRule(val quickfix.FIXString) ExchangeRuleField {
	return ExchangeRuleField{val}
}

//ExchangeSpecialInstructionsField is a STRING field
type ExchangeSpecialInstructionsField struct{ quickfix.FIXString }

//Tag returns tag.ExchangeSpecialInstructions (1139)
func (f ExchangeSpecialInstructionsField) Tag() quickfix.Tag { return tag.ExchangeSpecialInstructions }

//NewExchangeSpecialInstructions returns a new ExchangeSpecialInstructionsField initialized with val
func NewExchangeSpecialInstructions(val quickfix.FIXString) ExchangeSpecialInstructionsField {
	return ExchangeSpecialInstructionsField{val}
}

//ExecAckStatusField is a CHAR field
type ExecAckStatusField struct{ quickfix.FIXString }

//Tag returns tag.ExecAckStatus (1036)
func (f ExecAckStatusField) Tag() quickfix.Tag { return tag.ExecAckStatus }

//NewExecAckStatus returns a new ExecAckStatusField initialized with val
func NewExecAckStatus(val quickfix.FIXString) ExecAckStatusField {
	return ExecAckStatusField{val}
}

//ExecBrokerField is a STRING field
type ExecBrokerField struct{ quickfix.FIXString }

//Tag returns tag.ExecBroker (76)
func (f ExecBrokerField) Tag() quickfix.Tag { return tag.ExecBroker }

//NewExecBroker returns a new ExecBrokerField initialized with val
func NewExecBroker(val quickfix.FIXString) ExecBrokerField {
	return ExecBrokerField{val}
}

//ExecIDField is a STRING field
type ExecIDField struct{ quickfix.FIXString }

//Tag returns tag.ExecID (17)
func (f ExecIDField) Tag() quickfix.Tag { return tag.ExecID }

//NewExecID returns a new ExecIDField initialized with val
func NewExecID(val quickfix.FIXString) ExecIDField {
	return ExecIDField{val}
}

//ExecInstField is a MULTIPLECHARVALUE field
type ExecInstField struct{ quickfix.FIXString }

//Tag returns tag.ExecInst (18)
func (f ExecInstField) Tag() quickfix.Tag { return tag.ExecInst }

//NewExecInst returns a new ExecInstField initialized with val
func NewExecInst(val quickfix.FIXString) ExecInstField {
	return ExecInstField{val}
}

//ExecInstValueField is a CHAR field
type ExecInstValueField struct{ quickfix.FIXString }

//Tag returns tag.ExecInstValue (1308)
func (f ExecInstValueField) Tag() quickfix.Tag { return tag.ExecInstValue }

//NewExecInstValue returns a new ExecInstValueField initialized with val
func NewExecInstValue(val quickfix.FIXString) ExecInstValueField {
	return ExecInstValueField{val}
}

//ExecPriceAdjustmentField is a FLOAT field
type ExecPriceAdjustmentField struct{ quickfix.FIXFloat }

//Tag returns tag.ExecPriceAdjustment (485)
func (f ExecPriceAdjustmentField) Tag() quickfix.Tag { return tag.ExecPriceAdjustment }

//NewExecPriceAdjustment returns a new ExecPriceAdjustmentField initialized with val
func NewExecPriceAdjustment(val quickfix.FIXFloat) ExecPriceAdjustmentField {
	return ExecPriceAdjustmentField{val}
}

//ExecPriceTypeField is a CHAR field
type ExecPriceTypeField struct{ quickfix.FIXString }

//Tag returns tag.ExecPriceType (484)
func (f ExecPriceTypeField) Tag() quickfix.Tag { return tag.ExecPriceType }

//NewExecPriceType returns a new ExecPriceTypeField initialized with val
func NewExecPriceType(val quickfix.FIXString) ExecPriceTypeField {
	return ExecPriceTypeField{val}
}

//ExecRefIDField is a STRING field
type ExecRefIDField struct{ quickfix.FIXString }

//Tag returns tag.ExecRefID (19)
func (f ExecRefIDField) Tag() quickfix.Tag { return tag.ExecRefID }

//NewExecRefID returns a new ExecRefIDField initialized with val
func NewExecRefID(val quickfix.FIXString) ExecRefIDField {
	return ExecRefIDField{val}
}

//ExecRestatementReasonField is a INT field
type ExecRestatementReasonField struct{ quickfix.FIXInt }

//Tag returns tag.ExecRestatementReason (378)
func (f ExecRestatementReasonField) Tag() quickfix.Tag { return tag.ExecRestatementReason }

//NewExecRestatementReason returns a new ExecRestatementReasonField initialized with val
func NewExecRestatementReason(val quickfix.FIXInt) ExecRestatementReasonField {
	return ExecRestatementReasonField{val}
}

//ExecTransTypeField is a CHAR field
type ExecTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.ExecTransType (20)
func (f ExecTransTypeField) Tag() quickfix.Tag { return tag.ExecTransType }

//NewExecTransType returns a new ExecTransTypeField initialized with val
func NewExecTransType(val quickfix.FIXString) ExecTransTypeField {
	return ExecTransTypeField{val}
}

//ExecTypeField is a CHAR field
type ExecTypeField struct{ quickfix.FIXString }

//Tag returns tag.ExecType (150)
func (f ExecTypeField) Tag() quickfix.Tag { return tag.ExecType }

//NewExecType returns a new ExecTypeField initialized with val
func NewExecType(val quickfix.FIXString) ExecTypeField {
	return ExecTypeField{val}
}

//ExecValuationPointField is a UTCTIMESTAMP field
type ExecValuationPointField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.ExecValuationPoint (515)
func (f ExecValuationPointField) Tag() quickfix.Tag { return tag.ExecValuationPoint }

//NewExecValuationPoint returns a new ExecValuationPointField initialized with val
func NewExecValuationPoint(val time.Time) ExecValuationPointField {
	return ExecValuationPointField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewExecValuationPointNoMillis returns a new ExecValuationPointField initialized with val without millisecs
func NewExecValuationPointNoMillis(val time.Time) ExecValuationPointField {
	return ExecValuationPointField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//ExerciseMethodField is a CHAR field
type ExerciseMethodField struct{ quickfix.FIXString }

//Tag returns tag.ExerciseMethod (747)
func (f ExerciseMethodField) Tag() quickfix.Tag { return tag.ExerciseMethod }

//NewExerciseMethod returns a new ExerciseMethodField initialized with val
func NewExerciseMethod(val quickfix.FIXString) ExerciseMethodField {
	return ExerciseMethodField{val}
}

//ExerciseStyleField is a INT field
type ExerciseStyleField struct{ quickfix.FIXInt }

//Tag returns tag.ExerciseStyle (1194)
func (f ExerciseStyleField) Tag() quickfix.Tag { return tag.ExerciseStyle }

//NewExerciseStyle returns a new ExerciseStyleField initialized with val
func NewExerciseStyle(val quickfix.FIXInt) ExerciseStyleField {
	return ExerciseStyleField{val}
}

//ExpQtyField is a QTY field
type ExpQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.ExpQty (983)
func (f ExpQtyField) Tag() quickfix.Tag { return tag.ExpQty }

//NewExpQty returns a new ExpQtyField initialized with val
func NewExpQty(val quickfix.FIXFloat) ExpQtyField {
	return ExpQtyField{val}
}

//ExpTypeField is a INT field
type ExpTypeField struct{ quickfix.FIXInt }

//Tag returns tag.ExpType (982)
func (f ExpTypeField) Tag() quickfix.Tag { return tag.ExpType }

//NewExpType returns a new ExpTypeField initialized with val
func NewExpType(val quickfix.FIXInt) ExpTypeField {
	return ExpTypeField{val}
}

//ExpirationCycleField is a INT field
type ExpirationCycleField struct{ quickfix.FIXInt }

//Tag returns tag.ExpirationCycle (827)
func (f ExpirationCycleField) Tag() quickfix.Tag { return tag.ExpirationCycle }

//NewExpirationCycle returns a new ExpirationCycleField initialized with val
func NewExpirationCycle(val quickfix.FIXInt) ExpirationCycleField {
	return ExpirationCycleField{val}
}

//ExpirationQtyTypeField is a INT field
type ExpirationQtyTypeField struct{ quickfix.FIXInt }

//Tag returns tag.ExpirationQtyType (982)
func (f ExpirationQtyTypeField) Tag() quickfix.Tag { return tag.ExpirationQtyType }

//NewExpirationQtyType returns a new ExpirationQtyTypeField initialized with val
func NewExpirationQtyType(val quickfix.FIXInt) ExpirationQtyTypeField {
	return ExpirationQtyTypeField{val}
}

//ExpireDateField is a LOCALMKTDATE field
type ExpireDateField struct{ quickfix.FIXString }

//Tag returns tag.ExpireDate (432)
func (f ExpireDateField) Tag() quickfix.Tag { return tag.ExpireDate }

//NewExpireDate returns a new ExpireDateField initialized with val
func NewExpireDate(val quickfix.FIXString) ExpireDateField {
	return ExpireDateField{val}
}

//ExpireTimeField is a UTCTIMESTAMP field
type ExpireTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.ExpireTime (126)
func (f ExpireTimeField) Tag() quickfix.Tag { return tag.ExpireTime }

//NewExpireTime returns a new ExpireTimeField initialized with val
func NewExpireTime(val time.Time) ExpireTimeField {
	return ExpireTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewExpireTimeNoMillis returns a new ExpireTimeField initialized with val without millisecs
func NewExpireTimeNoMillis(val time.Time) ExpireTimeField {
	return ExpireTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//FactorField is a FLOAT field
type FactorField struct{ quickfix.FIXFloat }

//Tag returns tag.Factor (228)
func (f FactorField) Tag() quickfix.Tag { return tag.Factor }

//NewFactor returns a new FactorField initialized with val
func NewFactor(val quickfix.FIXFloat) FactorField {
	return FactorField{val}
}

//FairValueField is a AMT field
type FairValueField struct{ quickfix.FIXFloat }

//Tag returns tag.FairValue (406)
func (f FairValueField) Tag() quickfix.Tag { return tag.FairValue }

//NewFairValue returns a new FairValueField initialized with val
func NewFairValue(val quickfix.FIXFloat) FairValueField {
	return FairValueField{val}
}

//FeeMultiplierField is a FLOAT field
type FeeMultiplierField struct{ quickfix.FIXFloat }

//Tag returns tag.FeeMultiplier (1329)
func (f FeeMultiplierField) Tag() quickfix.Tag { return tag.FeeMultiplier }

//NewFeeMultiplier returns a new FeeMultiplierField initialized with val
func NewFeeMultiplier(val quickfix.FIXFloat) FeeMultiplierField {
	return FeeMultiplierField{val}
}

//FillExecIDField is a STRING field
type FillExecIDField struct{ quickfix.FIXString }

//Tag returns tag.FillExecID (1363)
func (f FillExecIDField) Tag() quickfix.Tag { return tag.FillExecID }

//NewFillExecID returns a new FillExecIDField initialized with val
func NewFillExecID(val quickfix.FIXString) FillExecIDField {
	return FillExecIDField{val}
}

//FillLiquidityIndField is a INT field
type FillLiquidityIndField struct{ quickfix.FIXInt }

//Tag returns tag.FillLiquidityInd (1443)
func (f FillLiquidityIndField) Tag() quickfix.Tag { return tag.FillLiquidityInd }

//NewFillLiquidityInd returns a new FillLiquidityIndField initialized with val
func NewFillLiquidityInd(val quickfix.FIXInt) FillLiquidityIndField {
	return FillLiquidityIndField{val}
}

//FillPxField is a PRICE field
type FillPxField struct{ quickfix.FIXFloat }

//Tag returns tag.FillPx (1364)
func (f FillPxField) Tag() quickfix.Tag { return tag.FillPx }

//NewFillPx returns a new FillPxField initialized with val
func NewFillPx(val quickfix.FIXFloat) FillPxField {
	return FillPxField{val}
}

//FillQtyField is a QTY field
type FillQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.FillQty (1365)
func (f FillQtyField) Tag() quickfix.Tag { return tag.FillQty }

//NewFillQty returns a new FillQtyField initialized with val
func NewFillQty(val quickfix.FIXFloat) FillQtyField {
	return FillQtyField{val}
}

//FinancialStatusField is a MULTIPLECHARVALUE field
type FinancialStatusField struct{ quickfix.FIXString }

//Tag returns tag.FinancialStatus (291)
func (f FinancialStatusField) Tag() quickfix.Tag { return tag.FinancialStatus }

//NewFinancialStatus returns a new FinancialStatusField initialized with val
func NewFinancialStatus(val quickfix.FIXString) FinancialStatusField {
	return FinancialStatusField{val}
}

//FirmTradeIDField is a STRING field
type FirmTradeIDField struct{ quickfix.FIXString }

//Tag returns tag.FirmTradeID (1041)
func (f FirmTradeIDField) Tag() quickfix.Tag { return tag.FirmTradeID }

//NewFirmTradeID returns a new FirmTradeIDField initialized with val
func NewFirmTradeID(val quickfix.FIXString) FirmTradeIDField {
	return FirmTradeIDField{val}
}

//FirstPxField is a PRICE field
type FirstPxField struct{ quickfix.FIXFloat }

//Tag returns tag.FirstPx (1025)
func (f FirstPxField) Tag() quickfix.Tag { return tag.FirstPx }

//NewFirstPx returns a new FirstPxField initialized with val
func NewFirstPx(val quickfix.FIXFloat) FirstPxField {
	return FirstPxField{val}
}

//FlexProductEligibilityIndicatorField is a BOOLEAN field
type FlexProductEligibilityIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.FlexProductEligibilityIndicator (1242)
func (f FlexProductEligibilityIndicatorField) Tag() quickfix.Tag {
	return tag.FlexProductEligibilityIndicator
}

//NewFlexProductEligibilityIndicator returns a new FlexProductEligibilityIndicatorField initialized with val
func NewFlexProductEligibilityIndicator(val quickfix.FIXBoolean) FlexProductEligibilityIndicatorField {
	return FlexProductEligibilityIndicatorField{val}
}

//FlexibleIndicatorField is a BOOLEAN field
type FlexibleIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.FlexibleIndicator (1244)
func (f FlexibleIndicatorField) Tag() quickfix.Tag { return tag.FlexibleIndicator }

//NewFlexibleIndicator returns a new FlexibleIndicatorField initialized with val
func NewFlexibleIndicator(val quickfix.FIXBoolean) FlexibleIndicatorField {
	return FlexibleIndicatorField{val}
}

//FloorPriceField is a PRICE field
type FloorPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.FloorPrice (1200)
func (f FloorPriceField) Tag() quickfix.Tag { return tag.FloorPrice }

//NewFloorPrice returns a new FloorPriceField initialized with val
func NewFloorPrice(val quickfix.FIXFloat) FloorPriceField {
	return FloorPriceField{val}
}

//FlowScheduleTypeField is a INT field
type FlowScheduleTypeField struct{ quickfix.FIXInt }

//Tag returns tag.FlowScheduleType (1439)
func (f FlowScheduleTypeField) Tag() quickfix.Tag { return tag.FlowScheduleType }

//NewFlowScheduleType returns a new FlowScheduleTypeField initialized with val
func NewFlowScheduleType(val quickfix.FIXInt) FlowScheduleTypeField {
	return FlowScheduleTypeField{val}
}

//ForexReqField is a BOOLEAN field
type ForexReqField struct{ quickfix.FIXBoolean }

//Tag returns tag.ForexReq (121)
func (f ForexReqField) Tag() quickfix.Tag { return tag.ForexReq }

//NewForexReq returns a new ForexReqField initialized with val
func NewForexReq(val quickfix.FIXBoolean) ForexReqField {
	return ForexReqField{val}
}

//FundRenewWaivField is a CHAR field
type FundRenewWaivField struct{ quickfix.FIXString }

//Tag returns tag.FundRenewWaiv (497)
func (f FundRenewWaivField) Tag() quickfix.Tag { return tag.FundRenewWaiv }

//NewFundRenewWaiv returns a new FundRenewWaivField initialized with val
func NewFundRenewWaiv(val quickfix.FIXString) FundRenewWaivField {
	return FundRenewWaivField{val}
}

//FutSettDateField is a LOCALMKTDATE field
type FutSettDateField struct{ quickfix.FIXString }

//Tag returns tag.FutSettDate (64)
func (f FutSettDateField) Tag() quickfix.Tag { return tag.FutSettDate }

//NewFutSettDate returns a new FutSettDateField initialized with val
func NewFutSettDate(val quickfix.FIXString) FutSettDateField {
	return FutSettDateField{val}
}

//FutSettDate2Field is a LOCALMKTDATE field
type FutSettDate2Field struct{ quickfix.FIXString }

//Tag returns tag.FutSettDate2 (193)
func (f FutSettDate2Field) Tag() quickfix.Tag { return tag.FutSettDate2 }

//NewFutSettDate2 returns a new FutSettDate2Field initialized with val
func NewFutSettDate2(val quickfix.FIXString) FutSettDate2Field {
	return FutSettDate2Field{val}
}

//FuturesValuationMethodField is a STRING field
type FuturesValuationMethodField struct{ quickfix.FIXString }

//Tag returns tag.FuturesValuationMethod (1197)
func (f FuturesValuationMethodField) Tag() quickfix.Tag { return tag.FuturesValuationMethod }

//NewFuturesValuationMethod returns a new FuturesValuationMethodField initialized with val
func NewFuturesValuationMethod(val quickfix.FIXString) FuturesValuationMethodField {
	return FuturesValuationMethodField{val}
}

//GTBookingInstField is a INT field
type GTBookingInstField struct{ quickfix.FIXInt }

//Tag returns tag.GTBookingInst (427)
func (f GTBookingInstField) Tag() quickfix.Tag { return tag.GTBookingInst }

//NewGTBookingInst returns a new GTBookingInstField initialized with val
func NewGTBookingInst(val quickfix.FIXInt) GTBookingInstField {
	return GTBookingInstField{val}
}

//GapFillFlagField is a BOOLEAN field
type GapFillFlagField struct{ quickfix.FIXBoolean }

//Tag returns tag.GapFillFlag (123)
func (f GapFillFlagField) Tag() quickfix.Tag { return tag.GapFillFlag }

//NewGapFillFlag returns a new GapFillFlagField initialized with val
func NewGapFillFlag(val quickfix.FIXBoolean) GapFillFlagField {
	return GapFillFlagField{val}
}

//GrossTradeAmtField is a AMT field
type GrossTradeAmtField struct{ quickfix.FIXFloat }

//Tag returns tag.GrossTradeAmt (381)
func (f GrossTradeAmtField) Tag() quickfix.Tag { return tag.GrossTradeAmt }

//NewGrossTradeAmt returns a new GrossTradeAmtField initialized with val
func NewGrossTradeAmt(val quickfix.FIXFloat) GrossTradeAmtField {
	return GrossTradeAmtField{val}
}

//HaltReasonCharField is a CHAR field
type HaltReasonCharField struct{ quickfix.FIXString }

//Tag returns tag.HaltReasonChar (327)
func (f HaltReasonCharField) Tag() quickfix.Tag { return tag.HaltReasonChar }

//NewHaltReasonChar returns a new HaltReasonCharField initialized with val
func NewHaltReasonChar(val quickfix.FIXString) HaltReasonCharField {
	return HaltReasonCharField{val}
}

//HaltReasonIntField is a INT field
type HaltReasonIntField struct{ quickfix.FIXInt }

//Tag returns tag.HaltReasonInt (327)
func (f HaltReasonIntField) Tag() quickfix.Tag { return tag.HaltReasonInt }

//NewHaltReasonInt returns a new HaltReasonIntField initialized with val
func NewHaltReasonInt(val quickfix.FIXInt) HaltReasonIntField {
	return HaltReasonIntField{val}
}

//HandlInstField is a CHAR field
type HandlInstField struct{ quickfix.FIXString }

//Tag returns tag.HandlInst (21)
func (f HandlInstField) Tag() quickfix.Tag { return tag.HandlInst }

//NewHandlInst returns a new HandlInstField initialized with val
func NewHandlInst(val quickfix.FIXString) HandlInstField {
	return HandlInstField{val}
}

//HeadlineField is a STRING field
type HeadlineField struct{ quickfix.FIXString }

//Tag returns tag.Headline (148)
func (f HeadlineField) Tag() quickfix.Tag { return tag.Headline }

//NewHeadline returns a new HeadlineField initialized with val
func NewHeadline(val quickfix.FIXString) HeadlineField {
	return HeadlineField{val}
}

//HeartBtIntField is a INT field
type HeartBtIntField struct{ quickfix.FIXInt }

//Tag returns tag.HeartBtInt (108)
func (f HeartBtIntField) Tag() quickfix.Tag { return tag.HeartBtInt }

//NewHeartBtInt returns a new HeartBtIntField initialized with val
func NewHeartBtInt(val quickfix.FIXInt) HeartBtIntField {
	return HeartBtIntField{val}
}

//HighLimitPriceField is a PRICE field
type HighLimitPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.HighLimitPrice (1149)
func (f HighLimitPriceField) Tag() quickfix.Tag { return tag.HighLimitPrice }

//NewHighLimitPrice returns a new HighLimitPriceField initialized with val
func NewHighLimitPrice(val quickfix.FIXFloat) HighLimitPriceField {
	return HighLimitPriceField{val}
}

//HighPxField is a PRICE field
type HighPxField struct{ quickfix.FIXFloat }

//Tag returns tag.HighPx (332)
func (f HighPxField) Tag() quickfix.Tag { return tag.HighPx }

//NewHighPx returns a new HighPxField initialized with val
func NewHighPx(val quickfix.FIXFloat) HighPxField {
	return HighPxField{val}
}

//HopCompIDField is a STRING field
type HopCompIDField struct{ quickfix.FIXString }

//Tag returns tag.HopCompID (628)
func (f HopCompIDField) Tag() quickfix.Tag { return tag.HopCompID }

//NewHopCompID returns a new HopCompIDField initialized with val
func NewHopCompID(val quickfix.FIXString) HopCompIDField {
	return HopCompIDField{val}
}

//HopRefIDField is a SEQNUM field
type HopRefIDField struct{ quickfix.FIXInt }

//Tag returns tag.HopRefID (630)
func (f HopRefIDField) Tag() quickfix.Tag { return tag.HopRefID }

//NewHopRefID returns a new HopRefIDField initialized with val
func NewHopRefID(val quickfix.FIXInt) HopRefIDField {
	return HopRefIDField{val}
}

//HopSendingTimeField is a UTCTIMESTAMP field
type HopSendingTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.HopSendingTime (629)
func (f HopSendingTimeField) Tag() quickfix.Tag { return tag.HopSendingTime }

//NewHopSendingTime returns a new HopSendingTimeField initialized with val
func NewHopSendingTime(val time.Time) HopSendingTimeField {
	return HopSendingTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewHopSendingTimeNoMillis returns a new HopSendingTimeField initialized with val without millisecs
func NewHopSendingTimeNoMillis(val time.Time) HopSendingTimeField {
	return HopSendingTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//HostCrossIDField is a STRING field
type HostCrossIDField struct{ quickfix.FIXString }

//Tag returns tag.HostCrossID (961)
func (f HostCrossIDField) Tag() quickfix.Tag { return tag.HostCrossID }

//NewHostCrossID returns a new HostCrossIDField initialized with val
func NewHostCrossID(val quickfix.FIXString) HostCrossIDField {
	return HostCrossIDField{val}
}

//IDSourceField is a STRING field
type IDSourceField struct{ quickfix.FIXString }

//Tag returns tag.IDSource (22)
func (f IDSourceField) Tag() quickfix.Tag { return tag.IDSource }

//NewIDSource returns a new IDSourceField initialized with val
func NewIDSource(val quickfix.FIXString) IDSourceField {
	return IDSourceField{val}
}

//IOIIDField is a STRING field
type IOIIDField struct{ quickfix.FIXString }

//Tag returns tag.IOIID (23)
func (f IOIIDField) Tag() quickfix.Tag { return tag.IOIID }

//NewIOIID returns a new IOIIDField initialized with val
func NewIOIID(val quickfix.FIXString) IOIIDField {
	return IOIIDField{val}
}

//IOINaturalFlagField is a BOOLEAN field
type IOINaturalFlagField struct{ quickfix.FIXBoolean }

//Tag returns tag.IOINaturalFlag (130)
func (f IOINaturalFlagField) Tag() quickfix.Tag { return tag.IOINaturalFlag }

//NewIOINaturalFlag returns a new IOINaturalFlagField initialized with val
func NewIOINaturalFlag(val quickfix.FIXBoolean) IOINaturalFlagField {
	return IOINaturalFlagField{val}
}

//IOIOthSvcField is a CHAR field
type IOIOthSvcField struct{ quickfix.FIXString }

//Tag returns tag.IOIOthSvc (24)
func (f IOIOthSvcField) Tag() quickfix.Tag { return tag.IOIOthSvc }

//NewIOIOthSvc returns a new IOIOthSvcField initialized with val
func NewIOIOthSvc(val quickfix.FIXString) IOIOthSvcField {
	return IOIOthSvcField{val}
}

//IOIQltyIndField is a CHAR field
type IOIQltyIndField struct{ quickfix.FIXString }

//Tag returns tag.IOIQltyInd (25)
func (f IOIQltyIndField) Tag() quickfix.Tag { return tag.IOIQltyInd }

//NewIOIQltyInd returns a new IOIQltyIndField initialized with val
func NewIOIQltyInd(val quickfix.FIXString) IOIQltyIndField {
	return IOIQltyIndField{val}
}

//IOIQtyField is a STRING field
type IOIQtyField struct{ quickfix.FIXString }

//Tag returns tag.IOIQty (27)
func (f IOIQtyField) Tag() quickfix.Tag { return tag.IOIQty }

//NewIOIQty returns a new IOIQtyField initialized with val
func NewIOIQty(val quickfix.FIXString) IOIQtyField {
	return IOIQtyField{val}
}

//IOIQualifierField is a CHAR field
type IOIQualifierField struct{ quickfix.FIXString }

//Tag returns tag.IOIQualifier (104)
func (f IOIQualifierField) Tag() quickfix.Tag { return tag.IOIQualifier }

//NewIOIQualifier returns a new IOIQualifierField initialized with val
func NewIOIQualifier(val quickfix.FIXString) IOIQualifierField {
	return IOIQualifierField{val}
}

//IOIRefIDField is a STRING field
type IOIRefIDField struct{ quickfix.FIXString }

//Tag returns tag.IOIRefID (26)
func (f IOIRefIDField) Tag() quickfix.Tag { return tag.IOIRefID }

//NewIOIRefID returns a new IOIRefIDField initialized with val
func NewIOIRefID(val quickfix.FIXString) IOIRefIDField {
	return IOIRefIDField{val}
}

//IOISharesField is a STRING field
type IOISharesField struct{ quickfix.FIXString }

//Tag returns tag.IOIShares (27)
func (f IOISharesField) Tag() quickfix.Tag { return tag.IOIShares }

//NewIOIShares returns a new IOISharesField initialized with val
func NewIOIShares(val quickfix.FIXString) IOISharesField {
	return IOISharesField{val}
}

//IOITransTypeField is a CHAR field
type IOITransTypeField struct{ quickfix.FIXString }

//Tag returns tag.IOITransType (28)
func (f IOITransTypeField) Tag() quickfix.Tag { return tag.IOITransType }

//NewIOITransType returns a new IOITransTypeField initialized with val
func NewIOITransType(val quickfix.FIXString) IOITransTypeField {
	return IOITransTypeField{val}
}

//IOIidField is a STRING field
type IOIidField struct{ quickfix.FIXString }

//Tag returns tag.IOIid (23)
func (f IOIidField) Tag() quickfix.Tag { return tag.IOIid }

//NewIOIid returns a new IOIidField initialized with val
func NewIOIid(val quickfix.FIXString) IOIidField {
	return IOIidField{val}
}

//ImpliedMarketIndicatorField is a INT field
type ImpliedMarketIndicatorField struct{ quickfix.FIXInt }

//Tag returns tag.ImpliedMarketIndicator (1144)
func (f ImpliedMarketIndicatorField) Tag() quickfix.Tag { return tag.ImpliedMarketIndicator }

//NewImpliedMarketIndicator returns a new ImpliedMarketIndicatorField initialized with val
func NewImpliedMarketIndicator(val quickfix.FIXInt) ImpliedMarketIndicatorField {
	return ImpliedMarketIndicatorField{val}
}

//InViewOfCommonField is a BOOLEAN field
type InViewOfCommonField struct{ quickfix.FIXBoolean }

//Tag returns tag.InViewOfCommon (328)
func (f InViewOfCommonField) Tag() quickfix.Tag { return tag.InViewOfCommon }

//NewInViewOfCommon returns a new InViewOfCommonField initialized with val
func NewInViewOfCommon(val quickfix.FIXBoolean) InViewOfCommonField {
	return InViewOfCommonField{val}
}

//IncTaxIndField is a INT field
type IncTaxIndField struct{ quickfix.FIXInt }

//Tag returns tag.IncTaxInd (416)
func (f IncTaxIndField) Tag() quickfix.Tag { return tag.IncTaxInd }

//NewIncTaxInd returns a new IncTaxIndField initialized with val
func NewIncTaxInd(val quickfix.FIXInt) IncTaxIndField {
	return IncTaxIndField{val}
}

//IndividualAllocIDField is a STRING field
type IndividualAllocIDField struct{ quickfix.FIXString }

//Tag returns tag.IndividualAllocID (467)
func (f IndividualAllocIDField) Tag() quickfix.Tag { return tag.IndividualAllocID }

//NewIndividualAllocID returns a new IndividualAllocIDField initialized with val
func NewIndividualAllocID(val quickfix.FIXString) IndividualAllocIDField {
	return IndividualAllocIDField{val}
}

//IndividualAllocRejCodeField is a INT field
type IndividualAllocRejCodeField struct{ quickfix.FIXInt }

//Tag returns tag.IndividualAllocRejCode (776)
func (f IndividualAllocRejCodeField) Tag() quickfix.Tag { return tag.IndividualAllocRejCode }

//NewIndividualAllocRejCode returns a new IndividualAllocRejCodeField initialized with val
func NewIndividualAllocRejCode(val quickfix.FIXInt) IndividualAllocRejCodeField {
	return IndividualAllocRejCodeField{val}
}

//IndividualAllocTypeField is a INT field
type IndividualAllocTypeField struct{ quickfix.FIXInt }

//Tag returns tag.IndividualAllocType (992)
func (f IndividualAllocTypeField) Tag() quickfix.Tag { return tag.IndividualAllocType }

//NewIndividualAllocType returns a new IndividualAllocTypeField initialized with val
func NewIndividualAllocType(val quickfix.FIXInt) IndividualAllocTypeField {
	return IndividualAllocTypeField{val}
}

//InputSourceField is a STRING field
type InputSourceField struct{ quickfix.FIXString }

//Tag returns tag.InputSource (979)
func (f InputSourceField) Tag() quickfix.Tag { return tag.InputSource }

//NewInputSource returns a new InputSourceField initialized with val
func NewInputSource(val quickfix.FIXString) InputSourceField {
	return InputSourceField{val}
}

//InstrAttribTypeField is a INT field
type InstrAttribTypeField struct{ quickfix.FIXInt }

//Tag returns tag.InstrAttribType (871)
func (f InstrAttribTypeField) Tag() quickfix.Tag { return tag.InstrAttribType }

//NewInstrAttribType returns a new InstrAttribTypeField initialized with val
func NewInstrAttribType(val quickfix.FIXInt) InstrAttribTypeField {
	return InstrAttribTypeField{val}
}

//InstrAttribValueField is a STRING field
type InstrAttribValueField struct{ quickfix.FIXString }

//Tag returns tag.InstrAttribValue (872)
func (f InstrAttribValueField) Tag() quickfix.Tag { return tag.InstrAttribValue }

//NewInstrAttribValue returns a new InstrAttribValueField initialized with val
func NewInstrAttribValue(val quickfix.FIXString) InstrAttribValueField {
	return InstrAttribValueField{val}
}

//InstrRegistryField is a STRING field
type InstrRegistryField struct{ quickfix.FIXString }

//Tag returns tag.InstrRegistry (543)
func (f InstrRegistryField) Tag() quickfix.Tag { return tag.InstrRegistry }

//NewInstrRegistry returns a new InstrRegistryField initialized with val
func NewInstrRegistry(val quickfix.FIXString) InstrRegistryField {
	return InstrRegistryField{val}
}

//InstrmtAssignmentMethodField is a CHAR field
type InstrmtAssignmentMethodField struct{ quickfix.FIXString }

//Tag returns tag.InstrmtAssignmentMethod (1049)
func (f InstrmtAssignmentMethodField) Tag() quickfix.Tag { return tag.InstrmtAssignmentMethod }

//NewInstrmtAssignmentMethod returns a new InstrmtAssignmentMethodField initialized with val
func NewInstrmtAssignmentMethod(val quickfix.FIXString) InstrmtAssignmentMethodField {
	return InstrmtAssignmentMethodField{val}
}

//InstrumentPartyIDField is a STRING field
type InstrumentPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.InstrumentPartyID (1019)
func (f InstrumentPartyIDField) Tag() quickfix.Tag { return tag.InstrumentPartyID }

//NewInstrumentPartyID returns a new InstrumentPartyIDField initialized with val
func NewInstrumentPartyID(val quickfix.FIXString) InstrumentPartyIDField {
	return InstrumentPartyIDField{val}
}

//InstrumentPartyIDSourceField is a CHAR field
type InstrumentPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.InstrumentPartyIDSource (1050)
func (f InstrumentPartyIDSourceField) Tag() quickfix.Tag { return tag.InstrumentPartyIDSource }

//NewInstrumentPartyIDSource returns a new InstrumentPartyIDSourceField initialized with val
func NewInstrumentPartyIDSource(val quickfix.FIXString) InstrumentPartyIDSourceField {
	return InstrumentPartyIDSourceField{val}
}

//InstrumentPartyRoleField is a INT field
type InstrumentPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.InstrumentPartyRole (1051)
func (f InstrumentPartyRoleField) Tag() quickfix.Tag { return tag.InstrumentPartyRole }

//NewInstrumentPartyRole returns a new InstrumentPartyRoleField initialized with val
func NewInstrumentPartyRole(val quickfix.FIXInt) InstrumentPartyRoleField {
	return InstrumentPartyRoleField{val}
}

//InstrumentPartySubIDField is a STRING field
type InstrumentPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.InstrumentPartySubID (1053)
func (f InstrumentPartySubIDField) Tag() quickfix.Tag { return tag.InstrumentPartySubID }

//NewInstrumentPartySubID returns a new InstrumentPartySubIDField initialized with val
func NewInstrumentPartySubID(val quickfix.FIXString) InstrumentPartySubIDField {
	return InstrumentPartySubIDField{val}
}

//InstrumentPartySubIDTypeField is a INT field
type InstrumentPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.InstrumentPartySubIDType (1054)
func (f InstrumentPartySubIDTypeField) Tag() quickfix.Tag { return tag.InstrumentPartySubIDType }

//NewInstrumentPartySubIDType returns a new InstrumentPartySubIDTypeField initialized with val
func NewInstrumentPartySubIDType(val quickfix.FIXInt) InstrumentPartySubIDTypeField {
	return InstrumentPartySubIDTypeField{val}
}

//InterestAccrualDateField is a LOCALMKTDATE field
type InterestAccrualDateField struct{ quickfix.FIXString }

//Tag returns tag.InterestAccrualDate (874)
func (f InterestAccrualDateField) Tag() quickfix.Tag { return tag.InterestAccrualDate }

//NewInterestAccrualDate returns a new InterestAccrualDateField initialized with val
func NewInterestAccrualDate(val quickfix.FIXString) InterestAccrualDateField {
	return InterestAccrualDateField{val}
}

//InterestAtMaturityField is a AMT field
type InterestAtMaturityField struct{ quickfix.FIXFloat }

//Tag returns tag.InterestAtMaturity (738)
func (f InterestAtMaturityField) Tag() quickfix.Tag { return tag.InterestAtMaturity }

//NewInterestAtMaturity returns a new InterestAtMaturityField initialized with val
func NewInterestAtMaturity(val quickfix.FIXFloat) InterestAtMaturityField {
	return InterestAtMaturityField{val}
}

//InvestorCountryOfResidenceField is a COUNTRY field
type InvestorCountryOfResidenceField struct{ quickfix.FIXString }

//Tag returns tag.InvestorCountryOfResidence (475)
func (f InvestorCountryOfResidenceField) Tag() quickfix.Tag { return tag.InvestorCountryOfResidence }

//NewInvestorCountryOfResidence returns a new InvestorCountryOfResidenceField initialized with val
func NewInvestorCountryOfResidence(val quickfix.FIXString) InvestorCountryOfResidenceField {
	return InvestorCountryOfResidenceField{val}
}

//IssueDateField is a LOCALMKTDATE field
type IssueDateField struct{ quickfix.FIXString }

//Tag returns tag.IssueDate (225)
func (f IssueDateField) Tag() quickfix.Tag { return tag.IssueDate }

//NewIssueDate returns a new IssueDateField initialized with val
func NewIssueDate(val quickfix.FIXString) IssueDateField {
	return IssueDateField{val}
}

//IssuerField is a STRING field
type IssuerField struct{ quickfix.FIXString }

//Tag returns tag.Issuer (106)
func (f IssuerField) Tag() quickfix.Tag { return tag.Issuer }

//NewIssuer returns a new IssuerField initialized with val
func NewIssuer(val quickfix.FIXString) IssuerField {
	return IssuerField{val}
}

//LanguageCodeField is a LANGUAGE field
type LanguageCodeField struct{ quickfix.FIXString }

//Tag returns tag.LanguageCode (1474)
func (f LanguageCodeField) Tag() quickfix.Tag { return tag.LanguageCode }

//NewLanguageCode returns a new LanguageCodeField initialized with val
func NewLanguageCode(val quickfix.FIXString) LanguageCodeField {
	return LanguageCodeField{val}
}

//LastCapacityField is a CHAR field
type LastCapacityField struct{ quickfix.FIXString }

//Tag returns tag.LastCapacity (29)
func (f LastCapacityField) Tag() quickfix.Tag { return tag.LastCapacity }

//NewLastCapacity returns a new LastCapacityField initialized with val
func NewLastCapacity(val quickfix.FIXString) LastCapacityField {
	return LastCapacityField{val}
}

//LastForwardPointsField is a PRICEOFFSET field
type LastForwardPointsField struct{ quickfix.FIXFloat }

//Tag returns tag.LastForwardPoints (195)
func (f LastForwardPointsField) Tag() quickfix.Tag { return tag.LastForwardPoints }

//NewLastForwardPoints returns a new LastForwardPointsField initialized with val
func NewLastForwardPoints(val quickfix.FIXFloat) LastForwardPointsField {
	return LastForwardPointsField{val}
}

//LastForwardPoints2Field is a PRICEOFFSET field
type LastForwardPoints2Field struct{ quickfix.FIXFloat }

//Tag returns tag.LastForwardPoints2 (641)
func (f LastForwardPoints2Field) Tag() quickfix.Tag { return tag.LastForwardPoints2 }

//NewLastForwardPoints2 returns a new LastForwardPoints2Field initialized with val
func NewLastForwardPoints2(val quickfix.FIXFloat) LastForwardPoints2Field {
	return LastForwardPoints2Field{val}
}

//LastFragmentField is a BOOLEAN field
type LastFragmentField struct{ quickfix.FIXBoolean }

//Tag returns tag.LastFragment (893)
func (f LastFragmentField) Tag() quickfix.Tag { return tag.LastFragment }

//NewLastFragment returns a new LastFragmentField initialized with val
func NewLastFragment(val quickfix.FIXBoolean) LastFragmentField {
	return LastFragmentField{val}
}

//LastLiquidityIndField is a INT field
type LastLiquidityIndField struct{ quickfix.FIXInt }

//Tag returns tag.LastLiquidityInd (851)
func (f LastLiquidityIndField) Tag() quickfix.Tag { return tag.LastLiquidityInd }

//NewLastLiquidityInd returns a new LastLiquidityIndField initialized with val
func NewLastLiquidityInd(val quickfix.FIXInt) LastLiquidityIndField {
	return LastLiquidityIndField{val}
}

//LastMktField is a EXCHANGE field
type LastMktField struct{ quickfix.FIXString }

//Tag returns tag.LastMkt (30)
func (f LastMktField) Tag() quickfix.Tag { return tag.LastMkt }

//NewLastMkt returns a new LastMktField initialized with val
func NewLastMkt(val quickfix.FIXString) LastMktField {
	return LastMktField{val}
}

//LastMsgSeqNumProcessedField is a SEQNUM field
type LastMsgSeqNumProcessedField struct{ quickfix.FIXInt }

//Tag returns tag.LastMsgSeqNumProcessed (369)
func (f LastMsgSeqNumProcessedField) Tag() quickfix.Tag { return tag.LastMsgSeqNumProcessed }

//NewLastMsgSeqNumProcessed returns a new LastMsgSeqNumProcessedField initialized with val
func NewLastMsgSeqNumProcessed(val quickfix.FIXInt) LastMsgSeqNumProcessedField {
	return LastMsgSeqNumProcessedField{val}
}

//LastNetworkResponseIDField is a STRING field
type LastNetworkResponseIDField struct{ quickfix.FIXString }

//Tag returns tag.LastNetworkResponseID (934)
func (f LastNetworkResponseIDField) Tag() quickfix.Tag { return tag.LastNetworkResponseID }

//NewLastNetworkResponseID returns a new LastNetworkResponseIDField initialized with val
func NewLastNetworkResponseID(val quickfix.FIXString) LastNetworkResponseIDField {
	return LastNetworkResponseIDField{val}
}

//LastParPxField is a PRICE field
type LastParPxField struct{ quickfix.FIXFloat }

//Tag returns tag.LastParPx (669)
func (f LastParPxField) Tag() quickfix.Tag { return tag.LastParPx }

//NewLastParPx returns a new LastParPxField initialized with val
func NewLastParPx(val quickfix.FIXFloat) LastParPxField {
	return LastParPxField{val}
}

//LastPxField is a PRICE field
type LastPxField struct{ quickfix.FIXFloat }

//Tag returns tag.LastPx (31)
func (f LastPxField) Tag() quickfix.Tag { return tag.LastPx }

//NewLastPx returns a new LastPxField initialized with val
func NewLastPx(val quickfix.FIXFloat) LastPxField {
	return LastPxField{val}
}

//LastQtyField is a QTY field
type LastQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.LastQty (32)
func (f LastQtyField) Tag() quickfix.Tag { return tag.LastQty }

//NewLastQty returns a new LastQtyField initialized with val
func NewLastQty(val quickfix.FIXFloat) LastQtyField {
	return LastQtyField{val}
}

//LastRptRequestedField is a BOOLEAN field
type LastRptRequestedField struct{ quickfix.FIXBoolean }

//Tag returns tag.LastRptRequested (912)
func (f LastRptRequestedField) Tag() quickfix.Tag { return tag.LastRptRequested }

//NewLastRptRequested returns a new LastRptRequestedField initialized with val
func NewLastRptRequested(val quickfix.FIXBoolean) LastRptRequestedField {
	return LastRptRequestedField{val}
}

//LastSharesField is a QTY field
type LastSharesField struct{ quickfix.FIXFloat }

//Tag returns tag.LastShares (32)
func (f LastSharesField) Tag() quickfix.Tag { return tag.LastShares }

//NewLastShares returns a new LastSharesField initialized with val
func NewLastShares(val quickfix.FIXFloat) LastSharesField {
	return LastSharesField{val}
}

//LastSpotRateField is a PRICE field
type LastSpotRateField struct{ quickfix.FIXFloat }

//Tag returns tag.LastSpotRate (194)
func (f LastSpotRateField) Tag() quickfix.Tag { return tag.LastSpotRate }

//NewLastSpotRate returns a new LastSpotRateField initialized with val
func NewLastSpotRate(val quickfix.FIXFloat) LastSpotRateField {
	return LastSpotRateField{val}
}

//LastSwapPointsField is a PRICEOFFSET field
type LastSwapPointsField struct{ quickfix.FIXFloat }

//Tag returns tag.LastSwapPoints (1071)
func (f LastSwapPointsField) Tag() quickfix.Tag { return tag.LastSwapPoints }

//NewLastSwapPoints returns a new LastSwapPointsField initialized with val
func NewLastSwapPoints(val quickfix.FIXFloat) LastSwapPointsField {
	return LastSwapPointsField{val}
}

//LastUpdateTimeField is a UTCTIMESTAMP field
type LastUpdateTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.LastUpdateTime (779)
func (f LastUpdateTimeField) Tag() quickfix.Tag { return tag.LastUpdateTime }

//NewLastUpdateTime returns a new LastUpdateTimeField initialized with val
func NewLastUpdateTime(val time.Time) LastUpdateTimeField {
	return LastUpdateTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewLastUpdateTimeNoMillis returns a new LastUpdateTimeField initialized with val without millisecs
func NewLastUpdateTimeNoMillis(val time.Time) LastUpdateTimeField {
	return LastUpdateTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//LateIndicatorField is a BOOLEAN field
type LateIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.LateIndicator (978)
func (f LateIndicatorField) Tag() quickfix.Tag { return tag.LateIndicator }

//NewLateIndicator returns a new LateIndicatorField initialized with val
func NewLateIndicator(val quickfix.FIXBoolean) LateIndicatorField {
	return LateIndicatorField{val}
}

//LeavesQtyField is a QTY field
type LeavesQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.LeavesQty (151)
func (f LeavesQtyField) Tag() quickfix.Tag { return tag.LeavesQty }

//NewLeavesQty returns a new LeavesQtyField initialized with val
func NewLeavesQty(val quickfix.FIXFloat) LeavesQtyField {
	return LeavesQtyField{val}
}

//LegAllocAccountField is a STRING field
type LegAllocAccountField struct{ quickfix.FIXString }

//Tag returns tag.LegAllocAccount (671)
func (f LegAllocAccountField) Tag() quickfix.Tag { return tag.LegAllocAccount }

//NewLegAllocAccount returns a new LegAllocAccountField initialized with val
func NewLegAllocAccount(val quickfix.FIXString) LegAllocAccountField {
	return LegAllocAccountField{val}
}

//LegAllocAcctIDSourceField is a STRING field
type LegAllocAcctIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.LegAllocAcctIDSource (674)
func (f LegAllocAcctIDSourceField) Tag() quickfix.Tag { return tag.LegAllocAcctIDSource }

//NewLegAllocAcctIDSource returns a new LegAllocAcctIDSourceField initialized with val
func NewLegAllocAcctIDSource(val quickfix.FIXString) LegAllocAcctIDSourceField {
	return LegAllocAcctIDSourceField{val}
}

//LegAllocIDField is a STRING field
type LegAllocIDField struct{ quickfix.FIXString }

//Tag returns tag.LegAllocID (1366)
func (f LegAllocIDField) Tag() quickfix.Tag { return tag.LegAllocID }

//NewLegAllocID returns a new LegAllocIDField initialized with val
func NewLegAllocID(val quickfix.FIXString) LegAllocIDField {
	return LegAllocIDField{val}
}

//LegAllocQtyField is a QTY field
type LegAllocQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.LegAllocQty (673)
func (f LegAllocQtyField) Tag() quickfix.Tag { return tag.LegAllocQty }

//NewLegAllocQty returns a new LegAllocQtyField initialized with val
func NewLegAllocQty(val quickfix.FIXFloat) LegAllocQtyField {
	return LegAllocQtyField{val}
}

//LegAllocSettlCurrencyField is a CURRENCY field
type LegAllocSettlCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.LegAllocSettlCurrency (1367)
func (f LegAllocSettlCurrencyField) Tag() quickfix.Tag { return tag.LegAllocSettlCurrency }

//NewLegAllocSettlCurrency returns a new LegAllocSettlCurrencyField initialized with val
func NewLegAllocSettlCurrency(val quickfix.FIXString) LegAllocSettlCurrencyField {
	return LegAllocSettlCurrencyField{val}
}

//LegBenchmarkCurveCurrencyField is a CURRENCY field
type LegBenchmarkCurveCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.LegBenchmarkCurveCurrency (676)
func (f LegBenchmarkCurveCurrencyField) Tag() quickfix.Tag { return tag.LegBenchmarkCurveCurrency }

//NewLegBenchmarkCurveCurrency returns a new LegBenchmarkCurveCurrencyField initialized with val
func NewLegBenchmarkCurveCurrency(val quickfix.FIXString) LegBenchmarkCurveCurrencyField {
	return LegBenchmarkCurveCurrencyField{val}
}

//LegBenchmarkCurveNameField is a STRING field
type LegBenchmarkCurveNameField struct{ quickfix.FIXString }

//Tag returns tag.LegBenchmarkCurveName (677)
func (f LegBenchmarkCurveNameField) Tag() quickfix.Tag { return tag.LegBenchmarkCurveName }

//NewLegBenchmarkCurveName returns a new LegBenchmarkCurveNameField initialized with val
func NewLegBenchmarkCurveName(val quickfix.FIXString) LegBenchmarkCurveNameField {
	return LegBenchmarkCurveNameField{val}
}

//LegBenchmarkCurvePointField is a STRING field
type LegBenchmarkCurvePointField struct{ quickfix.FIXString }

//Tag returns tag.LegBenchmarkCurvePoint (678)
func (f LegBenchmarkCurvePointField) Tag() quickfix.Tag { return tag.LegBenchmarkCurvePoint }

//NewLegBenchmarkCurvePoint returns a new LegBenchmarkCurvePointField initialized with val
func NewLegBenchmarkCurvePoint(val quickfix.FIXString) LegBenchmarkCurvePointField {
	return LegBenchmarkCurvePointField{val}
}

//LegBenchmarkPriceField is a PRICE field
type LegBenchmarkPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.LegBenchmarkPrice (679)
func (f LegBenchmarkPriceField) Tag() quickfix.Tag { return tag.LegBenchmarkPrice }

//NewLegBenchmarkPrice returns a new LegBenchmarkPriceField initialized with val
func NewLegBenchmarkPrice(val quickfix.FIXFloat) LegBenchmarkPriceField {
	return LegBenchmarkPriceField{val}
}

//LegBenchmarkPriceTypeField is a INT field
type LegBenchmarkPriceTypeField struct{ quickfix.FIXInt }

//Tag returns tag.LegBenchmarkPriceType (680)
func (f LegBenchmarkPriceTypeField) Tag() quickfix.Tag { return tag.LegBenchmarkPriceType }

//NewLegBenchmarkPriceType returns a new LegBenchmarkPriceTypeField initialized with val
func NewLegBenchmarkPriceType(val quickfix.FIXInt) LegBenchmarkPriceTypeField {
	return LegBenchmarkPriceTypeField{val}
}

//LegBidForwardPointsField is a PRICEOFFSET field
type LegBidForwardPointsField struct{ quickfix.FIXFloat }

//Tag returns tag.LegBidForwardPoints (1067)
func (f LegBidForwardPointsField) Tag() quickfix.Tag { return tag.LegBidForwardPoints }

//NewLegBidForwardPoints returns a new LegBidForwardPointsField initialized with val
func NewLegBidForwardPoints(val quickfix.FIXFloat) LegBidForwardPointsField {
	return LegBidForwardPointsField{val}
}

//LegBidPxField is a PRICE field
type LegBidPxField struct{ quickfix.FIXFloat }

//Tag returns tag.LegBidPx (681)
func (f LegBidPxField) Tag() quickfix.Tag { return tag.LegBidPx }

//NewLegBidPx returns a new LegBidPxField initialized with val
func NewLegBidPx(val quickfix.FIXFloat) LegBidPxField {
	return LegBidPxField{val}
}

//LegCFICodeField is a STRING field
type LegCFICodeField struct{ quickfix.FIXString }

//Tag returns tag.LegCFICode (608)
func (f LegCFICodeField) Tag() quickfix.Tag { return tag.LegCFICode }

//NewLegCFICode returns a new LegCFICodeField initialized with val
func NewLegCFICode(val quickfix.FIXString) LegCFICodeField {
	return LegCFICodeField{val}
}

//LegCalculatedCcyLastQtyField is a QTY field
type LegCalculatedCcyLastQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.LegCalculatedCcyLastQty (1074)
func (f LegCalculatedCcyLastQtyField) Tag() quickfix.Tag { return tag.LegCalculatedCcyLastQty }

//NewLegCalculatedCcyLastQty returns a new LegCalculatedCcyLastQtyField initialized with val
func NewLegCalculatedCcyLastQty(val quickfix.FIXFloat) LegCalculatedCcyLastQtyField {
	return LegCalculatedCcyLastQtyField{val}
}

//LegContractMultiplierField is a FLOAT field
type LegContractMultiplierField struct{ quickfix.FIXFloat }

//Tag returns tag.LegContractMultiplier (614)
func (f LegContractMultiplierField) Tag() quickfix.Tag { return tag.LegContractMultiplier }

//NewLegContractMultiplier returns a new LegContractMultiplierField initialized with val
func NewLegContractMultiplier(val quickfix.FIXFloat) LegContractMultiplierField {
	return LegContractMultiplierField{val}
}

//LegContractMultiplierUnitField is a INT field
type LegContractMultiplierUnitField struct{ quickfix.FIXInt }

//Tag returns tag.LegContractMultiplierUnit (1436)
func (f LegContractMultiplierUnitField) Tag() quickfix.Tag { return tag.LegContractMultiplierUnit }

//NewLegContractMultiplierUnit returns a new LegContractMultiplierUnitField initialized with val
func NewLegContractMultiplierUnit(val quickfix.FIXInt) LegContractMultiplierUnitField {
	return LegContractMultiplierUnitField{val}
}

//LegContractSettlMonthField is a MONTHYEAR field
type LegContractSettlMonthField struct{ quickfix.FIXString }

//Tag returns tag.LegContractSettlMonth (955)
func (f LegContractSettlMonthField) Tag() quickfix.Tag { return tag.LegContractSettlMonth }

//NewLegContractSettlMonth returns a new LegContractSettlMonthField initialized with val
func NewLegContractSettlMonth(val quickfix.FIXString) LegContractSettlMonthField {
	return LegContractSettlMonthField{val}
}

//LegCountryOfIssueField is a COUNTRY field
type LegCountryOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.LegCountryOfIssue (596)
func (f LegCountryOfIssueField) Tag() quickfix.Tag { return tag.LegCountryOfIssue }

//NewLegCountryOfIssue returns a new LegCountryOfIssueField initialized with val
func NewLegCountryOfIssue(val quickfix.FIXString) LegCountryOfIssueField {
	return LegCountryOfIssueField{val}
}

//LegCouponPaymentDateField is a LOCALMKTDATE field
type LegCouponPaymentDateField struct{ quickfix.FIXString }

//Tag returns tag.LegCouponPaymentDate (248)
func (f LegCouponPaymentDateField) Tag() quickfix.Tag { return tag.LegCouponPaymentDate }

//NewLegCouponPaymentDate returns a new LegCouponPaymentDateField initialized with val
func NewLegCouponPaymentDate(val quickfix.FIXString) LegCouponPaymentDateField {
	return LegCouponPaymentDateField{val}
}

//LegCouponRateField is a PERCENTAGE field
type LegCouponRateField struct{ quickfix.FIXFloat }

//Tag returns tag.LegCouponRate (615)
func (f LegCouponRateField) Tag() quickfix.Tag { return tag.LegCouponRate }

//NewLegCouponRate returns a new LegCouponRateField initialized with val
func NewLegCouponRate(val quickfix.FIXFloat) LegCouponRateField {
	return LegCouponRateField{val}
}

//LegCoveredOrUncoveredField is a INT field
type LegCoveredOrUncoveredField struct{ quickfix.FIXInt }

//Tag returns tag.LegCoveredOrUncovered (565)
func (f LegCoveredOrUncoveredField) Tag() quickfix.Tag { return tag.LegCoveredOrUncovered }

//NewLegCoveredOrUncovered returns a new LegCoveredOrUncoveredField initialized with val
func NewLegCoveredOrUncovered(val quickfix.FIXInt) LegCoveredOrUncoveredField {
	return LegCoveredOrUncoveredField{val}
}

//LegCreditRatingField is a STRING field
type LegCreditRatingField struct{ quickfix.FIXString }

//Tag returns tag.LegCreditRating (257)
func (f LegCreditRatingField) Tag() quickfix.Tag { return tag.LegCreditRating }

//NewLegCreditRating returns a new LegCreditRatingField initialized with val
func NewLegCreditRating(val quickfix.FIXString) LegCreditRatingField {
	return LegCreditRatingField{val}
}

//LegCurrencyField is a CURRENCY field
type LegCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.LegCurrency (556)
func (f LegCurrencyField) Tag() quickfix.Tag { return tag.LegCurrency }

//NewLegCurrency returns a new LegCurrencyField initialized with val
func NewLegCurrency(val quickfix.FIXString) LegCurrencyField {
	return LegCurrencyField{val}
}

//LegCurrencyRatioField is a FLOAT field
type LegCurrencyRatioField struct{ quickfix.FIXFloat }

//Tag returns tag.LegCurrencyRatio (1383)
func (f LegCurrencyRatioField) Tag() quickfix.Tag { return tag.LegCurrencyRatio }

//NewLegCurrencyRatio returns a new LegCurrencyRatioField initialized with val
func NewLegCurrencyRatio(val quickfix.FIXFloat) LegCurrencyRatioField {
	return LegCurrencyRatioField{val}
}

//LegDatedDateField is a LOCALMKTDATE field
type LegDatedDateField struct{ quickfix.FIXString }

//Tag returns tag.LegDatedDate (739)
func (f LegDatedDateField) Tag() quickfix.Tag { return tag.LegDatedDate }

//NewLegDatedDate returns a new LegDatedDateField initialized with val
func NewLegDatedDate(val quickfix.FIXString) LegDatedDateField {
	return LegDatedDateField{val}
}

//LegDividendYieldField is a PERCENTAGE field
type LegDividendYieldField struct{ quickfix.FIXFloat }

//Tag returns tag.LegDividendYield (1381)
func (f LegDividendYieldField) Tag() quickfix.Tag { return tag.LegDividendYield }

//NewLegDividendYield returns a new LegDividendYieldField initialized with val
func NewLegDividendYield(val quickfix.FIXFloat) LegDividendYieldField {
	return LegDividendYieldField{val}
}

//LegExecInstField is a MULTIPLECHARVALUE field
type LegExecInstField struct{ quickfix.FIXString }

//Tag returns tag.LegExecInst (1384)
func (f LegExecInstField) Tag() quickfix.Tag { return tag.LegExecInst }

//NewLegExecInst returns a new LegExecInstField initialized with val
func NewLegExecInst(val quickfix.FIXString) LegExecInstField {
	return LegExecInstField{val}
}

//LegExerciseStyleField is a INT field
type LegExerciseStyleField struct{ quickfix.FIXInt }

//Tag returns tag.LegExerciseStyle (1420)
func (f LegExerciseStyleField) Tag() quickfix.Tag { return tag.LegExerciseStyle }

//NewLegExerciseStyle returns a new LegExerciseStyleField initialized with val
func NewLegExerciseStyle(val quickfix.FIXInt) LegExerciseStyleField {
	return LegExerciseStyleField{val}
}

//LegFactorField is a FLOAT field
type LegFactorField struct{ quickfix.FIXFloat }

//Tag returns tag.LegFactor (253)
func (f LegFactorField) Tag() quickfix.Tag { return tag.LegFactor }

//NewLegFactor returns a new LegFactorField initialized with val
func NewLegFactor(val quickfix.FIXFloat) LegFactorField {
	return LegFactorField{val}
}

//LegFlowScheduleTypeField is a INT field
type LegFlowScheduleTypeField struct{ quickfix.FIXInt }

//Tag returns tag.LegFlowScheduleType (1440)
func (f LegFlowScheduleTypeField) Tag() quickfix.Tag { return tag.LegFlowScheduleType }

//NewLegFlowScheduleType returns a new LegFlowScheduleTypeField initialized with val
func NewLegFlowScheduleType(val quickfix.FIXInt) LegFlowScheduleTypeField {
	return LegFlowScheduleTypeField{val}
}

//LegFutSettDateField is a LOCALMKTDATE field
type LegFutSettDateField struct{ quickfix.FIXString }

//Tag returns tag.LegFutSettDate (588)
func (f LegFutSettDateField) Tag() quickfix.Tag { return tag.LegFutSettDate }

//NewLegFutSettDate returns a new LegFutSettDateField initialized with val
func NewLegFutSettDate(val quickfix.FIXString) LegFutSettDateField {
	return LegFutSettDateField{val}
}

//LegGrossTradeAmtField is a AMT field
type LegGrossTradeAmtField struct{ quickfix.FIXFloat }

//Tag returns tag.LegGrossTradeAmt (1075)
func (f LegGrossTradeAmtField) Tag() quickfix.Tag { return tag.LegGrossTradeAmt }

//NewLegGrossTradeAmt returns a new LegGrossTradeAmtField initialized with val
func NewLegGrossTradeAmt(val quickfix.FIXFloat) LegGrossTradeAmtField {
	return LegGrossTradeAmtField{val}
}

//LegIOIQtyField is a STRING field
type LegIOIQtyField struct{ quickfix.FIXString }

//Tag returns tag.LegIOIQty (682)
func (f LegIOIQtyField) Tag() quickfix.Tag { return tag.LegIOIQty }

//NewLegIOIQty returns a new LegIOIQtyField initialized with val
func NewLegIOIQty(val quickfix.FIXString) LegIOIQtyField {
	return LegIOIQtyField{val}
}

//LegIndividualAllocIDField is a STRING field
type LegIndividualAllocIDField struct{ quickfix.FIXString }

//Tag returns tag.LegIndividualAllocID (672)
func (f LegIndividualAllocIDField) Tag() quickfix.Tag { return tag.LegIndividualAllocID }

//NewLegIndividualAllocID returns a new LegIndividualAllocIDField initialized with val
func NewLegIndividualAllocID(val quickfix.FIXString) LegIndividualAllocIDField {
	return LegIndividualAllocIDField{val}
}

//LegInstrRegistryField is a STRING field
type LegInstrRegistryField struct{ quickfix.FIXString }

//Tag returns tag.LegInstrRegistry (599)
func (f LegInstrRegistryField) Tag() quickfix.Tag { return tag.LegInstrRegistry }

//NewLegInstrRegistry returns a new LegInstrRegistryField initialized with val
func NewLegInstrRegistry(val quickfix.FIXString) LegInstrRegistryField {
	return LegInstrRegistryField{val}
}

//LegInterestAccrualDateField is a LOCALMKTDATE field
type LegInterestAccrualDateField struct{ quickfix.FIXString }

//Tag returns tag.LegInterestAccrualDate (956)
func (f LegInterestAccrualDateField) Tag() quickfix.Tag { return tag.LegInterestAccrualDate }

//NewLegInterestAccrualDate returns a new LegInterestAccrualDateField initialized with val
func NewLegInterestAccrualDate(val quickfix.FIXString) LegInterestAccrualDateField {
	return LegInterestAccrualDateField{val}
}

//LegIssueDateField is a LOCALMKTDATE field
type LegIssueDateField struct{ quickfix.FIXString }

//Tag returns tag.LegIssueDate (249)
func (f LegIssueDateField) Tag() quickfix.Tag { return tag.LegIssueDate }

//NewLegIssueDate returns a new LegIssueDateField initialized with val
func NewLegIssueDate(val quickfix.FIXString) LegIssueDateField {
	return LegIssueDateField{val}
}

//LegIssuerField is a STRING field
type LegIssuerField struct{ quickfix.FIXString }

//Tag returns tag.LegIssuer (617)
func (f LegIssuerField) Tag() quickfix.Tag { return tag.LegIssuer }

//NewLegIssuer returns a new LegIssuerField initialized with val
func NewLegIssuer(val quickfix.FIXString) LegIssuerField {
	return LegIssuerField{val}
}

//LegLastForwardPointsField is a PRICEOFFSET field
type LegLastForwardPointsField struct{ quickfix.FIXFloat }

//Tag returns tag.LegLastForwardPoints (1073)
func (f LegLastForwardPointsField) Tag() quickfix.Tag { return tag.LegLastForwardPoints }

//NewLegLastForwardPoints returns a new LegLastForwardPointsField initialized with val
func NewLegLastForwardPoints(val quickfix.FIXFloat) LegLastForwardPointsField {
	return LegLastForwardPointsField{val}
}

//LegLastPxField is a PRICE field
type LegLastPxField struct{ quickfix.FIXFloat }

//Tag returns tag.LegLastPx (637)
func (f LegLastPxField) Tag() quickfix.Tag { return tag.LegLastPx }

//NewLegLastPx returns a new LegLastPxField initialized with val
func NewLegLastPx(val quickfix.FIXFloat) LegLastPxField {
	return LegLastPxField{val}
}

//LegLastQtyField is a QTY field
type LegLastQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.LegLastQty (1418)
func (f LegLastQtyField) Tag() quickfix.Tag { return tag.LegLastQty }

//NewLegLastQty returns a new LegLastQtyField initialized with val
func NewLegLastQty(val quickfix.FIXFloat) LegLastQtyField {
	return LegLastQtyField{val}
}

//LegLocaleOfIssueField is a STRING field
type LegLocaleOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.LegLocaleOfIssue (598)
func (f LegLocaleOfIssueField) Tag() quickfix.Tag { return tag.LegLocaleOfIssue }

//NewLegLocaleOfIssue returns a new LegLocaleOfIssueField initialized with val
func NewLegLocaleOfIssue(val quickfix.FIXString) LegLocaleOfIssueField {
	return LegLocaleOfIssueField{val}
}

//LegMaturityDateField is a LOCALMKTDATE field
type LegMaturityDateField struct{ quickfix.FIXString }

//Tag returns tag.LegMaturityDate (611)
func (f LegMaturityDateField) Tag() quickfix.Tag { return tag.LegMaturityDate }

//NewLegMaturityDate returns a new LegMaturityDateField initialized with val
func NewLegMaturityDate(val quickfix.FIXString) LegMaturityDateField {
	return LegMaturityDateField{val}
}

//LegMaturityMonthYearField is a MONTHYEAR field
type LegMaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.LegMaturityMonthYear (610)
func (f LegMaturityMonthYearField) Tag() quickfix.Tag { return tag.LegMaturityMonthYear }

//NewLegMaturityMonthYear returns a new LegMaturityMonthYearField initialized with val
func NewLegMaturityMonthYear(val quickfix.FIXString) LegMaturityMonthYearField {
	return LegMaturityMonthYearField{val}
}

//LegMaturityTimeField is a TZTIMEONLY field
type LegMaturityTimeField struct{ quickfix.FIXString }

//Tag returns tag.LegMaturityTime (1212)
func (f LegMaturityTimeField) Tag() quickfix.Tag { return tag.LegMaturityTime }

//NewLegMaturityTime returns a new LegMaturityTimeField initialized with val
func NewLegMaturityTime(val quickfix.FIXString) LegMaturityTimeField {
	return LegMaturityTimeField{val}
}

//LegNumberField is a INT field
type LegNumberField struct{ quickfix.FIXInt }

//Tag returns tag.LegNumber (1152)
func (f LegNumberField) Tag() quickfix.Tag { return tag.LegNumber }

//NewLegNumber returns a new LegNumberField initialized with val
func NewLegNumber(val quickfix.FIXInt) LegNumberField {
	return LegNumberField{val}
}

//LegOfferForwardPointsField is a PRICEOFFSET field
type LegOfferForwardPointsField struct{ quickfix.FIXFloat }

//Tag returns tag.LegOfferForwardPoints (1068)
func (f LegOfferForwardPointsField) Tag() quickfix.Tag { return tag.LegOfferForwardPoints }

//NewLegOfferForwardPoints returns a new LegOfferForwardPointsField initialized with val
func NewLegOfferForwardPoints(val quickfix.FIXFloat) LegOfferForwardPointsField {
	return LegOfferForwardPointsField{val}
}

//LegOfferPxField is a PRICE field
type LegOfferPxField struct{ quickfix.FIXFloat }

//Tag returns tag.LegOfferPx (684)
func (f LegOfferPxField) Tag() quickfix.Tag { return tag.LegOfferPx }

//NewLegOfferPx returns a new LegOfferPxField initialized with val
func NewLegOfferPx(val quickfix.FIXFloat) LegOfferPxField {
	return LegOfferPxField{val}
}

//LegOptAttributeField is a CHAR field
type LegOptAttributeField struct{ quickfix.FIXString }

//Tag returns tag.LegOptAttribute (613)
func (f LegOptAttributeField) Tag() quickfix.Tag { return tag.LegOptAttribute }

//NewLegOptAttribute returns a new LegOptAttributeField initialized with val
func NewLegOptAttribute(val quickfix.FIXString) LegOptAttributeField {
	return LegOptAttributeField{val}
}

//LegOptionRatioField is a FLOAT field
type LegOptionRatioField struct{ quickfix.FIXFloat }

//Tag returns tag.LegOptionRatio (1017)
func (f LegOptionRatioField) Tag() quickfix.Tag { return tag.LegOptionRatio }

//NewLegOptionRatio returns a new LegOptionRatioField initialized with val
func NewLegOptionRatio(val quickfix.FIXFloat) LegOptionRatioField {
	return LegOptionRatioField{val}
}

//LegOrderQtyField is a QTY field
type LegOrderQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.LegOrderQty (685)
func (f LegOrderQtyField) Tag() quickfix.Tag { return tag.LegOrderQty }

//NewLegOrderQty returns a new LegOrderQtyField initialized with val
func NewLegOrderQty(val quickfix.FIXFloat) LegOrderQtyField {
	return LegOrderQtyField{val}
}

//LegPoolField is a STRING field
type LegPoolField struct{ quickfix.FIXString }

//Tag returns tag.LegPool (740)
func (f LegPoolField) Tag() quickfix.Tag { return tag.LegPool }

//NewLegPool returns a new LegPoolField initialized with val
func NewLegPool(val quickfix.FIXString) LegPoolField {
	return LegPoolField{val}
}

//LegPositionEffectField is a CHAR field
type LegPositionEffectField struct{ quickfix.FIXString }

//Tag returns tag.LegPositionEffect (564)
func (f LegPositionEffectField) Tag() quickfix.Tag { return tag.LegPositionEffect }

//NewLegPositionEffect returns a new LegPositionEffectField initialized with val
func NewLegPositionEffect(val quickfix.FIXString) LegPositionEffectField {
	return LegPositionEffectField{val}
}

//LegPriceField is a PRICE field
type LegPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.LegPrice (566)
func (f LegPriceField) Tag() quickfix.Tag { return tag.LegPrice }

//NewLegPrice returns a new LegPriceField initialized with val
func NewLegPrice(val quickfix.FIXFloat) LegPriceField {
	return LegPriceField{val}
}

//LegPriceTypeField is a INT field
type LegPriceTypeField struct{ quickfix.FIXInt }

//Tag returns tag.LegPriceType (686)
func (f LegPriceTypeField) Tag() quickfix.Tag { return tag.LegPriceType }

//NewLegPriceType returns a new LegPriceTypeField initialized with val
func NewLegPriceType(val quickfix.FIXInt) LegPriceTypeField {
	return LegPriceTypeField{val}
}

//LegPriceUnitOfMeasureField is a STRING field
type LegPriceUnitOfMeasureField struct{ quickfix.FIXString }

//Tag returns tag.LegPriceUnitOfMeasure (1421)
func (f LegPriceUnitOfMeasureField) Tag() quickfix.Tag { return tag.LegPriceUnitOfMeasure }

//NewLegPriceUnitOfMeasure returns a new LegPriceUnitOfMeasureField initialized with val
func NewLegPriceUnitOfMeasure(val quickfix.FIXString) LegPriceUnitOfMeasureField {
	return LegPriceUnitOfMeasureField{val}
}

//LegPriceUnitOfMeasureQtyField is a QTY field
type LegPriceUnitOfMeasureQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.LegPriceUnitOfMeasureQty (1422)
func (f LegPriceUnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.LegPriceUnitOfMeasureQty }

//NewLegPriceUnitOfMeasureQty returns a new LegPriceUnitOfMeasureQtyField initialized with val
func NewLegPriceUnitOfMeasureQty(val quickfix.FIXFloat) LegPriceUnitOfMeasureQtyField {
	return LegPriceUnitOfMeasureQtyField{val}
}

//LegProductField is a INT field
type LegProductField struct{ quickfix.FIXInt }

//Tag returns tag.LegProduct (607)
func (f LegProductField) Tag() quickfix.Tag { return tag.LegProduct }

//NewLegProduct returns a new LegProductField initialized with val
func NewLegProduct(val quickfix.FIXInt) LegProductField {
	return LegProductField{val}
}

//LegPutOrCallField is a INT field
type LegPutOrCallField struct{ quickfix.FIXInt }

//Tag returns tag.LegPutOrCall (1358)
func (f LegPutOrCallField) Tag() quickfix.Tag { return tag.LegPutOrCall }

//NewLegPutOrCall returns a new LegPutOrCallField initialized with val
func NewLegPutOrCall(val quickfix.FIXInt) LegPutOrCallField {
	return LegPutOrCallField{val}
}

//LegQtyField is a QTY field
type LegQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.LegQty (687)
func (f LegQtyField) Tag() quickfix.Tag { return tag.LegQty }

//NewLegQty returns a new LegQtyField initialized with val
func NewLegQty(val quickfix.FIXFloat) LegQtyField {
	return LegQtyField{val}
}

//LegRatioQtyField is a FLOAT field
type LegRatioQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.LegRatioQty (623)
func (f LegRatioQtyField) Tag() quickfix.Tag { return tag.LegRatioQty }

//NewLegRatioQty returns a new LegRatioQtyField initialized with val
func NewLegRatioQty(val quickfix.FIXFloat) LegRatioQtyField {
	return LegRatioQtyField{val}
}

//LegRedemptionDateField is a LOCALMKTDATE field
type LegRedemptionDateField struct{ quickfix.FIXString }

//Tag returns tag.LegRedemptionDate (254)
func (f LegRedemptionDateField) Tag() quickfix.Tag { return tag.LegRedemptionDate }

//NewLegRedemptionDate returns a new LegRedemptionDateField initialized with val
func NewLegRedemptionDate(val quickfix.FIXString) LegRedemptionDateField {
	return LegRedemptionDateField{val}
}

//LegRefIDField is a STRING field
type LegRefIDField struct{ quickfix.FIXString }

//Tag returns tag.LegRefID (654)
func (f LegRefIDField) Tag() quickfix.Tag { return tag.LegRefID }

//NewLegRefID returns a new LegRefIDField initialized with val
func NewLegRefID(val quickfix.FIXString) LegRefIDField {
	return LegRefIDField{val}
}

//LegRepoCollateralSecurityTypeField is a INT field
type LegRepoCollateralSecurityTypeField struct{ quickfix.FIXInt }

//Tag returns tag.LegRepoCollateralSecurityType (250)
func (f LegRepoCollateralSecurityTypeField) Tag() quickfix.Tag {
	return tag.LegRepoCollateralSecurityType
}

//NewLegRepoCollateralSecurityType returns a new LegRepoCollateralSecurityTypeField initialized with val
func NewLegRepoCollateralSecurityType(val quickfix.FIXInt) LegRepoCollateralSecurityTypeField {
	return LegRepoCollateralSecurityTypeField{val}
}

//LegReportIDField is a STRING field
type LegReportIDField struct{ quickfix.FIXString }

//Tag returns tag.LegReportID (990)
func (f LegReportIDField) Tag() quickfix.Tag { return tag.LegReportID }

//NewLegReportID returns a new LegReportIDField initialized with val
func NewLegReportID(val quickfix.FIXString) LegReportIDField {
	return LegReportIDField{val}
}

//LegRepurchaseRateField is a PERCENTAGE field
type LegRepurchaseRateField struct{ quickfix.FIXFloat }

//Tag returns tag.LegRepurchaseRate (252)
func (f LegRepurchaseRateField) Tag() quickfix.Tag { return tag.LegRepurchaseRate }

//NewLegRepurchaseRate returns a new LegRepurchaseRateField initialized with val
func NewLegRepurchaseRate(val quickfix.FIXFloat) LegRepurchaseRateField {
	return LegRepurchaseRateField{val}
}

//LegRepurchaseTermField is a INT field
type LegRepurchaseTermField struct{ quickfix.FIXInt }

//Tag returns tag.LegRepurchaseTerm (251)
func (f LegRepurchaseTermField) Tag() quickfix.Tag { return tag.LegRepurchaseTerm }

//NewLegRepurchaseTerm returns a new LegRepurchaseTermField initialized with val
func NewLegRepurchaseTerm(val quickfix.FIXInt) LegRepurchaseTermField {
	return LegRepurchaseTermField{val}
}

//LegSecurityAltIDField is a STRING field
type LegSecurityAltIDField struct{ quickfix.FIXString }

//Tag returns tag.LegSecurityAltID (605)
func (f LegSecurityAltIDField) Tag() quickfix.Tag { return tag.LegSecurityAltID }

//NewLegSecurityAltID returns a new LegSecurityAltIDField initialized with val
func NewLegSecurityAltID(val quickfix.FIXString) LegSecurityAltIDField {
	return LegSecurityAltIDField{val}
}

//LegSecurityAltIDSourceField is a STRING field
type LegSecurityAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.LegSecurityAltIDSource (606)
func (f LegSecurityAltIDSourceField) Tag() quickfix.Tag { return tag.LegSecurityAltIDSource }

//NewLegSecurityAltIDSource returns a new LegSecurityAltIDSourceField initialized with val
func NewLegSecurityAltIDSource(val quickfix.FIXString) LegSecurityAltIDSourceField {
	return LegSecurityAltIDSourceField{val}
}

//LegSecurityDescField is a STRING field
type LegSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.LegSecurityDesc (620)
func (f LegSecurityDescField) Tag() quickfix.Tag { return tag.LegSecurityDesc }

//NewLegSecurityDesc returns a new LegSecurityDescField initialized with val
func NewLegSecurityDesc(val quickfix.FIXString) LegSecurityDescField {
	return LegSecurityDescField{val}
}

//LegSecurityExchangeField is a EXCHANGE field
type LegSecurityExchangeField struct{ quickfix.FIXString }

//Tag returns tag.LegSecurityExchange (616)
func (f LegSecurityExchangeField) Tag() quickfix.Tag { return tag.LegSecurityExchange }

//NewLegSecurityExchange returns a new LegSecurityExchangeField initialized with val
func NewLegSecurityExchange(val quickfix.FIXString) LegSecurityExchangeField {
	return LegSecurityExchangeField{val}
}

//LegSecurityIDField is a STRING field
type LegSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.LegSecurityID (602)
func (f LegSecurityIDField) Tag() quickfix.Tag { return tag.LegSecurityID }

//NewLegSecurityID returns a new LegSecurityIDField initialized with val
func NewLegSecurityID(val quickfix.FIXString) LegSecurityIDField {
	return LegSecurityIDField{val}
}

//LegSecurityIDSourceField is a STRING field
type LegSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.LegSecurityIDSource (603)
func (f LegSecurityIDSourceField) Tag() quickfix.Tag { return tag.LegSecurityIDSource }

//NewLegSecurityIDSource returns a new LegSecurityIDSourceField initialized with val
func NewLegSecurityIDSource(val quickfix.FIXString) LegSecurityIDSourceField {
	return LegSecurityIDSourceField{val}
}

//LegSecuritySubTypeField is a STRING field
type LegSecuritySubTypeField struct{ quickfix.FIXString }

//Tag returns tag.LegSecuritySubType (764)
func (f LegSecuritySubTypeField) Tag() quickfix.Tag { return tag.LegSecuritySubType }

//NewLegSecuritySubType returns a new LegSecuritySubTypeField initialized with val
func NewLegSecuritySubType(val quickfix.FIXString) LegSecuritySubTypeField {
	return LegSecuritySubTypeField{val}
}

//LegSecurityTypeField is a STRING field
type LegSecurityTypeField struct{ quickfix.FIXString }

//Tag returns tag.LegSecurityType (609)
func (f LegSecurityTypeField) Tag() quickfix.Tag { return tag.LegSecurityType }

//NewLegSecurityType returns a new LegSecurityTypeField initialized with val
func NewLegSecurityType(val quickfix.FIXString) LegSecurityTypeField {
	return LegSecurityTypeField{val}
}

//LegSettlCurrencyField is a CURRENCY field
type LegSettlCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.LegSettlCurrency (675)
func (f LegSettlCurrencyField) Tag() quickfix.Tag { return tag.LegSettlCurrency }

//NewLegSettlCurrency returns a new LegSettlCurrencyField initialized with val
func NewLegSettlCurrency(val quickfix.FIXString) LegSettlCurrencyField {
	return LegSettlCurrencyField{val}
}

//LegSettlDateField is a LOCALMKTDATE field
type LegSettlDateField struct{ quickfix.FIXString }

//Tag returns tag.LegSettlDate (588)
func (f LegSettlDateField) Tag() quickfix.Tag { return tag.LegSettlDate }

//NewLegSettlDate returns a new LegSettlDateField initialized with val
func NewLegSettlDate(val quickfix.FIXString) LegSettlDateField {
	return LegSettlDateField{val}
}

//LegSettlTypeField is a CHAR field
type LegSettlTypeField struct{ quickfix.FIXString }

//Tag returns tag.LegSettlType (587)
func (f LegSettlTypeField) Tag() quickfix.Tag { return tag.LegSettlType }

//NewLegSettlType returns a new LegSettlTypeField initialized with val
func NewLegSettlType(val quickfix.FIXString) LegSettlTypeField {
	return LegSettlTypeField{val}
}

//LegSettlmntTypField is a CHAR field
type LegSettlmntTypField struct{ quickfix.FIXString }

//Tag returns tag.LegSettlmntTyp (587)
func (f LegSettlmntTypField) Tag() quickfix.Tag { return tag.LegSettlmntTyp }

//NewLegSettlmntTyp returns a new LegSettlmntTypField initialized with val
func NewLegSettlmntTyp(val quickfix.FIXString) LegSettlmntTypField {
	return LegSettlmntTypField{val}
}

//LegSideField is a CHAR field
type LegSideField struct{ quickfix.FIXString }

//Tag returns tag.LegSide (624)
func (f LegSideField) Tag() quickfix.Tag { return tag.LegSide }

//NewLegSide returns a new LegSideField initialized with val
func NewLegSide(val quickfix.FIXString) LegSideField {
	return LegSideField{val}
}

//LegStateOrProvinceOfIssueField is a STRING field
type LegStateOrProvinceOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.LegStateOrProvinceOfIssue (597)
func (f LegStateOrProvinceOfIssueField) Tag() quickfix.Tag { return tag.LegStateOrProvinceOfIssue }

//NewLegStateOrProvinceOfIssue returns a new LegStateOrProvinceOfIssueField initialized with val
func NewLegStateOrProvinceOfIssue(val quickfix.FIXString) LegStateOrProvinceOfIssueField {
	return LegStateOrProvinceOfIssueField{val}
}

//LegStipulationTypeField is a STRING field
type LegStipulationTypeField struct{ quickfix.FIXString }

//Tag returns tag.LegStipulationType (688)
func (f LegStipulationTypeField) Tag() quickfix.Tag { return tag.LegStipulationType }

//NewLegStipulationType returns a new LegStipulationTypeField initialized with val
func NewLegStipulationType(val quickfix.FIXString) LegStipulationTypeField {
	return LegStipulationTypeField{val}
}

//LegStipulationValueField is a STRING field
type LegStipulationValueField struct{ quickfix.FIXString }

//Tag returns tag.LegStipulationValue (689)
func (f LegStipulationValueField) Tag() quickfix.Tag { return tag.LegStipulationValue }

//NewLegStipulationValue returns a new LegStipulationValueField initialized with val
func NewLegStipulationValue(val quickfix.FIXString) LegStipulationValueField {
	return LegStipulationValueField{val}
}

//LegStrikeCurrencyField is a CURRENCY field
type LegStrikeCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.LegStrikeCurrency (942)
func (f LegStrikeCurrencyField) Tag() quickfix.Tag { return tag.LegStrikeCurrency }

//NewLegStrikeCurrency returns a new LegStrikeCurrencyField initialized with val
func NewLegStrikeCurrency(val quickfix.FIXString) LegStrikeCurrencyField {
	return LegStrikeCurrencyField{val}
}

//LegStrikePriceField is a PRICE field
type LegStrikePriceField struct{ quickfix.FIXFloat }

//Tag returns tag.LegStrikePrice (612)
func (f LegStrikePriceField) Tag() quickfix.Tag { return tag.LegStrikePrice }

//NewLegStrikePrice returns a new LegStrikePriceField initialized with val
func NewLegStrikePrice(val quickfix.FIXFloat) LegStrikePriceField {
	return LegStrikePriceField{val}
}

//LegSwapTypeField is a INT field
type LegSwapTypeField struct{ quickfix.FIXInt }

//Tag returns tag.LegSwapType (690)
func (f LegSwapTypeField) Tag() quickfix.Tag { return tag.LegSwapType }

//NewLegSwapType returns a new LegSwapTypeField initialized with val
func NewLegSwapType(val quickfix.FIXInt) LegSwapTypeField {
	return LegSwapTypeField{val}
}

//LegSymbolField is a STRING field
type LegSymbolField struct{ quickfix.FIXString }

//Tag returns tag.LegSymbol (600)
func (f LegSymbolField) Tag() quickfix.Tag { return tag.LegSymbol }

//NewLegSymbol returns a new LegSymbolField initialized with val
func NewLegSymbol(val quickfix.FIXString) LegSymbolField {
	return LegSymbolField{val}
}

//LegSymbolSfxField is a STRING field
type LegSymbolSfxField struct{ quickfix.FIXString }

//Tag returns tag.LegSymbolSfx (601)
func (f LegSymbolSfxField) Tag() quickfix.Tag { return tag.LegSymbolSfx }

//NewLegSymbolSfx returns a new LegSymbolSfxField initialized with val
func NewLegSymbolSfx(val quickfix.FIXString) LegSymbolSfxField {
	return LegSymbolSfxField{val}
}

//LegTimeUnitField is a STRING field
type LegTimeUnitField struct{ quickfix.FIXString }

//Tag returns tag.LegTimeUnit (1001)
func (f LegTimeUnitField) Tag() quickfix.Tag { return tag.LegTimeUnit }

//NewLegTimeUnit returns a new LegTimeUnitField initialized with val
func NewLegTimeUnit(val quickfix.FIXString) LegTimeUnitField {
	return LegTimeUnitField{val}
}

//LegUnitOfMeasureField is a STRING field
type LegUnitOfMeasureField struct{ quickfix.FIXString }

//Tag returns tag.LegUnitOfMeasure (999)
func (f LegUnitOfMeasureField) Tag() quickfix.Tag { return tag.LegUnitOfMeasure }

//NewLegUnitOfMeasure returns a new LegUnitOfMeasureField initialized with val
func NewLegUnitOfMeasure(val quickfix.FIXString) LegUnitOfMeasureField {
	return LegUnitOfMeasureField{val}
}

//LegUnitOfMeasureQtyField is a QTY field
type LegUnitOfMeasureQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.LegUnitOfMeasureQty (1224)
func (f LegUnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.LegUnitOfMeasureQty }

//NewLegUnitOfMeasureQty returns a new LegUnitOfMeasureQtyField initialized with val
func NewLegUnitOfMeasureQty(val quickfix.FIXFloat) LegUnitOfMeasureQtyField {
	return LegUnitOfMeasureQtyField{val}
}

//LegVolatilityField is a FLOAT field
type LegVolatilityField struct{ quickfix.FIXFloat }

//Tag returns tag.LegVolatility (1379)
func (f LegVolatilityField) Tag() quickfix.Tag { return tag.LegVolatility }

//NewLegVolatility returns a new LegVolatilityField initialized with val
func NewLegVolatility(val quickfix.FIXFloat) LegVolatilityField {
	return LegVolatilityField{val}
}

//LegalConfirmField is a BOOLEAN field
type LegalConfirmField struct{ quickfix.FIXBoolean }

//Tag returns tag.LegalConfirm (650)
func (f LegalConfirmField) Tag() quickfix.Tag { return tag.LegalConfirm }

//NewLegalConfirm returns a new LegalConfirmField initialized with val
func NewLegalConfirm(val quickfix.FIXBoolean) LegalConfirmField {
	return LegalConfirmField{val}
}

//LinesOfTextField is a NUMINGROUP field
type LinesOfTextField struct{ quickfix.FIXInt }

//Tag returns tag.LinesOfText (33)
func (f LinesOfTextField) Tag() quickfix.Tag { return tag.LinesOfText }

//NewLinesOfText returns a new LinesOfTextField initialized with val
func NewLinesOfText(val quickfix.FIXInt) LinesOfTextField {
	return LinesOfTextField{val}
}

//LiquidityIndTypeField is a INT field
type LiquidityIndTypeField struct{ quickfix.FIXInt }

//Tag returns tag.LiquidityIndType (409)
func (f LiquidityIndTypeField) Tag() quickfix.Tag { return tag.LiquidityIndType }

//NewLiquidityIndType returns a new LiquidityIndTypeField initialized with val
func NewLiquidityIndType(val quickfix.FIXInt) LiquidityIndTypeField {
	return LiquidityIndTypeField{val}
}

//LiquidityNumSecuritiesField is a INT field
type LiquidityNumSecuritiesField struct{ quickfix.FIXInt }

//Tag returns tag.LiquidityNumSecurities (441)
func (f LiquidityNumSecuritiesField) Tag() quickfix.Tag { return tag.LiquidityNumSecurities }

//NewLiquidityNumSecurities returns a new LiquidityNumSecuritiesField initialized with val
func NewLiquidityNumSecurities(val quickfix.FIXInt) LiquidityNumSecuritiesField {
	return LiquidityNumSecuritiesField{val}
}

//LiquidityPctHighField is a PERCENTAGE field
type LiquidityPctHighField struct{ quickfix.FIXFloat }

//Tag returns tag.LiquidityPctHigh (403)
func (f LiquidityPctHighField) Tag() quickfix.Tag { return tag.LiquidityPctHigh }

//NewLiquidityPctHigh returns a new LiquidityPctHighField initialized with val
func NewLiquidityPctHigh(val quickfix.FIXFloat) LiquidityPctHighField {
	return LiquidityPctHighField{val}
}

//LiquidityPctLowField is a PERCENTAGE field
type LiquidityPctLowField struct{ quickfix.FIXFloat }

//Tag returns tag.LiquidityPctLow (402)
func (f LiquidityPctLowField) Tag() quickfix.Tag { return tag.LiquidityPctLow }

//NewLiquidityPctLow returns a new LiquidityPctLowField initialized with val
func NewLiquidityPctLow(val quickfix.FIXFloat) LiquidityPctLowField {
	return LiquidityPctLowField{val}
}

//LiquidityValueField is a AMT field
type LiquidityValueField struct{ quickfix.FIXFloat }

//Tag returns tag.LiquidityValue (404)
func (f LiquidityValueField) Tag() quickfix.Tag { return tag.LiquidityValue }

//NewLiquidityValue returns a new LiquidityValueField initialized with val
func NewLiquidityValue(val quickfix.FIXFloat) LiquidityValueField {
	return LiquidityValueField{val}
}

//ListExecInstField is a STRING field
type ListExecInstField struct{ quickfix.FIXString }

//Tag returns tag.ListExecInst (69)
func (f ListExecInstField) Tag() quickfix.Tag { return tag.ListExecInst }

//NewListExecInst returns a new ListExecInstField initialized with val
func NewListExecInst(val quickfix.FIXString) ListExecInstField {
	return ListExecInstField{val}
}

//ListExecInstTypeField is a CHAR field
type ListExecInstTypeField struct{ quickfix.FIXString }

//Tag returns tag.ListExecInstType (433)
func (f ListExecInstTypeField) Tag() quickfix.Tag { return tag.ListExecInstType }

//NewListExecInstType returns a new ListExecInstTypeField initialized with val
func NewListExecInstType(val quickfix.FIXString) ListExecInstTypeField {
	return ListExecInstTypeField{val}
}

//ListIDField is a STRING field
type ListIDField struct{ quickfix.FIXString }

//Tag returns tag.ListID (66)
func (f ListIDField) Tag() quickfix.Tag { return tag.ListID }

//NewListID returns a new ListIDField initialized with val
func NewListID(val quickfix.FIXString) ListIDField {
	return ListIDField{val}
}

//ListMethodField is a INT field
type ListMethodField struct{ quickfix.FIXInt }

//Tag returns tag.ListMethod (1198)
func (f ListMethodField) Tag() quickfix.Tag { return tag.ListMethod }

//NewListMethod returns a new ListMethodField initialized with val
func NewListMethod(val quickfix.FIXInt) ListMethodField {
	return ListMethodField{val}
}

//ListNameField is a STRING field
type ListNameField struct{ quickfix.FIXString }

//Tag returns tag.ListName (392)
func (f ListNameField) Tag() quickfix.Tag { return tag.ListName }

//NewListName returns a new ListNameField initialized with val
func NewListName(val quickfix.FIXString) ListNameField {
	return ListNameField{val}
}

//ListNoOrdsField is a INT field
type ListNoOrdsField struct{ quickfix.FIXInt }

//Tag returns tag.ListNoOrds (68)
func (f ListNoOrdsField) Tag() quickfix.Tag { return tag.ListNoOrds }

//NewListNoOrds returns a new ListNoOrdsField initialized with val
func NewListNoOrds(val quickfix.FIXInt) ListNoOrdsField {
	return ListNoOrdsField{val}
}

//ListOrderStatusField is a INT field
type ListOrderStatusField struct{ quickfix.FIXInt }

//Tag returns tag.ListOrderStatus (431)
func (f ListOrderStatusField) Tag() quickfix.Tag { return tag.ListOrderStatus }

//NewListOrderStatus returns a new ListOrderStatusField initialized with val
func NewListOrderStatus(val quickfix.FIXInt) ListOrderStatusField {
	return ListOrderStatusField{val}
}

//ListRejectReasonField is a INT field
type ListRejectReasonField struct{ quickfix.FIXInt }

//Tag returns tag.ListRejectReason (1386)
func (f ListRejectReasonField) Tag() quickfix.Tag { return tag.ListRejectReason }

//NewListRejectReason returns a new ListRejectReasonField initialized with val
func NewListRejectReason(val quickfix.FIXInt) ListRejectReasonField {
	return ListRejectReasonField{val}
}

//ListSeqNoField is a INT field
type ListSeqNoField struct{ quickfix.FIXInt }

//Tag returns tag.ListSeqNo (67)
func (f ListSeqNoField) Tag() quickfix.Tag { return tag.ListSeqNo }

//NewListSeqNo returns a new ListSeqNoField initialized with val
func NewListSeqNo(val quickfix.FIXInt) ListSeqNoField {
	return ListSeqNoField{val}
}

//ListStatusTextField is a STRING field
type ListStatusTextField struct{ quickfix.FIXString }

//Tag returns tag.ListStatusText (444)
func (f ListStatusTextField) Tag() quickfix.Tag { return tag.ListStatusText }

//NewListStatusText returns a new ListStatusTextField initialized with val
func NewListStatusText(val quickfix.FIXString) ListStatusTextField {
	return ListStatusTextField{val}
}

//ListStatusTypeField is a INT field
type ListStatusTypeField struct{ quickfix.FIXInt }

//Tag returns tag.ListStatusType (429)
func (f ListStatusTypeField) Tag() quickfix.Tag { return tag.ListStatusType }

//NewListStatusType returns a new ListStatusTypeField initialized with val
func NewListStatusType(val quickfix.FIXInt) ListStatusTypeField {
	return ListStatusTypeField{val}
}

//ListUpdateActionField is a CHAR field
type ListUpdateActionField struct{ quickfix.FIXString }

//Tag returns tag.ListUpdateAction (1324)
func (f ListUpdateActionField) Tag() quickfix.Tag { return tag.ListUpdateAction }

//NewListUpdateAction returns a new ListUpdateActionField initialized with val
func NewListUpdateAction(val quickfix.FIXString) ListUpdateActionField {
	return ListUpdateActionField{val}
}

//LocaleOfIssueField is a STRING field
type LocaleOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.LocaleOfIssue (472)
func (f LocaleOfIssueField) Tag() quickfix.Tag { return tag.LocaleOfIssue }

//NewLocaleOfIssue returns a new LocaleOfIssueField initialized with val
func NewLocaleOfIssue(val quickfix.FIXString) LocaleOfIssueField {
	return LocaleOfIssueField{val}
}

//LocateReqdField is a BOOLEAN field
type LocateReqdField struct{ quickfix.FIXBoolean }

//Tag returns tag.LocateReqd (114)
func (f LocateReqdField) Tag() quickfix.Tag { return tag.LocateReqd }

//NewLocateReqd returns a new LocateReqdField initialized with val
func NewLocateReqd(val quickfix.FIXBoolean) LocateReqdField {
	return LocateReqdField{val}
}

//LocationIDField is a STRING field
type LocationIDField struct{ quickfix.FIXString }

//Tag returns tag.LocationID (283)
func (f LocationIDField) Tag() quickfix.Tag { return tag.LocationID }

//NewLocationID returns a new LocationIDField initialized with val
func NewLocationID(val quickfix.FIXString) LocationIDField {
	return LocationIDField{val}
}

//LongQtyField is a QTY field
type LongQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.LongQty (704)
func (f LongQtyField) Tag() quickfix.Tag { return tag.LongQty }

//NewLongQty returns a new LongQtyField initialized with val
func NewLongQty(val quickfix.FIXFloat) LongQtyField {
	return LongQtyField{val}
}

//LotTypeField is a CHAR field
type LotTypeField struct{ quickfix.FIXString }

//Tag returns tag.LotType (1093)
func (f LotTypeField) Tag() quickfix.Tag { return tag.LotType }

//NewLotType returns a new LotTypeField initialized with val
func NewLotType(val quickfix.FIXString) LotTypeField {
	return LotTypeField{val}
}

//LowLimitPriceField is a PRICE field
type LowLimitPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.LowLimitPrice (1148)
func (f LowLimitPriceField) Tag() quickfix.Tag { return tag.LowLimitPrice }

//NewLowLimitPrice returns a new LowLimitPriceField initialized with val
func NewLowLimitPrice(val quickfix.FIXFloat) LowLimitPriceField {
	return LowLimitPriceField{val}
}

//LowPxField is a PRICE field
type LowPxField struct{ quickfix.FIXFloat }

//Tag returns tag.LowPx (333)
func (f LowPxField) Tag() quickfix.Tag { return tag.LowPx }

//NewLowPx returns a new LowPxField initialized with val
func NewLowPx(val quickfix.FIXFloat) LowPxField {
	return LowPxField{val}
}

//MDBookTypeField is a INT field
type MDBookTypeField struct{ quickfix.FIXInt }

//Tag returns tag.MDBookType (1021)
func (f MDBookTypeField) Tag() quickfix.Tag { return tag.MDBookType }

//NewMDBookType returns a new MDBookTypeField initialized with val
func NewMDBookType(val quickfix.FIXInt) MDBookTypeField {
	return MDBookTypeField{val}
}

//MDEntryBuyerField is a STRING field
type MDEntryBuyerField struct{ quickfix.FIXString }

//Tag returns tag.MDEntryBuyer (288)
func (f MDEntryBuyerField) Tag() quickfix.Tag { return tag.MDEntryBuyer }

//NewMDEntryBuyer returns a new MDEntryBuyerField initialized with val
func NewMDEntryBuyer(val quickfix.FIXString) MDEntryBuyerField {
	return MDEntryBuyerField{val}
}

//MDEntryDateField is a UTCDATEONLY field
type MDEntryDateField struct{ quickfix.FIXString }

//Tag returns tag.MDEntryDate (272)
func (f MDEntryDateField) Tag() quickfix.Tag { return tag.MDEntryDate }

//NewMDEntryDate returns a new MDEntryDateField initialized with val
func NewMDEntryDate(val quickfix.FIXString) MDEntryDateField {
	return MDEntryDateField{val}
}

//MDEntryForwardPointsField is a PRICEOFFSET field
type MDEntryForwardPointsField struct{ quickfix.FIXFloat }

//Tag returns tag.MDEntryForwardPoints (1027)
func (f MDEntryForwardPointsField) Tag() quickfix.Tag { return tag.MDEntryForwardPoints }

//NewMDEntryForwardPoints returns a new MDEntryForwardPointsField initialized with val
func NewMDEntryForwardPoints(val quickfix.FIXFloat) MDEntryForwardPointsField {
	return MDEntryForwardPointsField{val}
}

//MDEntryIDField is a STRING field
type MDEntryIDField struct{ quickfix.FIXString }

//Tag returns tag.MDEntryID (278)
func (f MDEntryIDField) Tag() quickfix.Tag { return tag.MDEntryID }

//NewMDEntryID returns a new MDEntryIDField initialized with val
func NewMDEntryID(val quickfix.FIXString) MDEntryIDField {
	return MDEntryIDField{val}
}

//MDEntryOriginatorField is a STRING field
type MDEntryOriginatorField struct{ quickfix.FIXString }

//Tag returns tag.MDEntryOriginator (282)
func (f MDEntryOriginatorField) Tag() quickfix.Tag { return tag.MDEntryOriginator }

//NewMDEntryOriginator returns a new MDEntryOriginatorField initialized with val
func NewMDEntryOriginator(val quickfix.FIXString) MDEntryOriginatorField {
	return MDEntryOriginatorField{val}
}

//MDEntryPositionNoField is a INT field
type MDEntryPositionNoField struct{ quickfix.FIXInt }

//Tag returns tag.MDEntryPositionNo (290)
func (f MDEntryPositionNoField) Tag() quickfix.Tag { return tag.MDEntryPositionNo }

//NewMDEntryPositionNo returns a new MDEntryPositionNoField initialized with val
func NewMDEntryPositionNo(val quickfix.FIXInt) MDEntryPositionNoField {
	return MDEntryPositionNoField{val}
}

//MDEntryPxField is a PRICE field
type MDEntryPxField struct{ quickfix.FIXFloat }

//Tag returns tag.MDEntryPx (270)
func (f MDEntryPxField) Tag() quickfix.Tag { return tag.MDEntryPx }

//NewMDEntryPx returns a new MDEntryPxField initialized with val
func NewMDEntryPx(val quickfix.FIXFloat) MDEntryPxField {
	return MDEntryPxField{val}
}

//MDEntryRefIDField is a STRING field
type MDEntryRefIDField struct{ quickfix.FIXString }

//Tag returns tag.MDEntryRefID (280)
func (f MDEntryRefIDField) Tag() quickfix.Tag { return tag.MDEntryRefID }

//NewMDEntryRefID returns a new MDEntryRefIDField initialized with val
func NewMDEntryRefID(val quickfix.FIXString) MDEntryRefIDField {
	return MDEntryRefIDField{val}
}

//MDEntrySellerField is a STRING field
type MDEntrySellerField struct{ quickfix.FIXString }

//Tag returns tag.MDEntrySeller (289)
func (f MDEntrySellerField) Tag() quickfix.Tag { return tag.MDEntrySeller }

//NewMDEntrySeller returns a new MDEntrySellerField initialized with val
func NewMDEntrySeller(val quickfix.FIXString) MDEntrySellerField {
	return MDEntrySellerField{val}
}

//MDEntrySizeField is a QTY field
type MDEntrySizeField struct{ quickfix.FIXFloat }

//Tag returns tag.MDEntrySize (271)
func (f MDEntrySizeField) Tag() quickfix.Tag { return tag.MDEntrySize }

//NewMDEntrySize returns a new MDEntrySizeField initialized with val
func NewMDEntrySize(val quickfix.FIXFloat) MDEntrySizeField {
	return MDEntrySizeField{val}
}

//MDEntrySpotRateField is a FLOAT field
type MDEntrySpotRateField struct{ quickfix.FIXFloat }

//Tag returns tag.MDEntrySpotRate (1026)
func (f MDEntrySpotRateField) Tag() quickfix.Tag { return tag.MDEntrySpotRate }

//NewMDEntrySpotRate returns a new MDEntrySpotRateField initialized with val
func NewMDEntrySpotRate(val quickfix.FIXFloat) MDEntrySpotRateField {
	return MDEntrySpotRateField{val}
}

//MDEntryTimeField is a UTCTIMEONLY field
type MDEntryTimeField struct{ quickfix.FIXString }

//Tag returns tag.MDEntryTime (273)
func (f MDEntryTimeField) Tag() quickfix.Tag { return tag.MDEntryTime }

//NewMDEntryTime returns a new MDEntryTimeField initialized with val
func NewMDEntryTime(val quickfix.FIXString) MDEntryTimeField {
	return MDEntryTimeField{val}
}

//MDEntryTypeField is a CHAR field
type MDEntryTypeField struct{ quickfix.FIXString }

//Tag returns tag.MDEntryType (269)
func (f MDEntryTypeField) Tag() quickfix.Tag { return tag.MDEntryType }

//NewMDEntryType returns a new MDEntryTypeField initialized with val
func NewMDEntryType(val quickfix.FIXString) MDEntryTypeField {
	return MDEntryTypeField{val}
}

//MDFeedTypeField is a STRING field
type MDFeedTypeField struct{ quickfix.FIXString }

//Tag returns tag.MDFeedType (1022)
func (f MDFeedTypeField) Tag() quickfix.Tag { return tag.MDFeedType }

//NewMDFeedType returns a new MDFeedTypeField initialized with val
func NewMDFeedType(val quickfix.FIXString) MDFeedTypeField {
	return MDFeedTypeField{val}
}

//MDImplicitDeleteField is a BOOLEAN field
type MDImplicitDeleteField struct{ quickfix.FIXBoolean }

//Tag returns tag.MDImplicitDelete (547)
func (f MDImplicitDeleteField) Tag() quickfix.Tag { return tag.MDImplicitDelete }

//NewMDImplicitDelete returns a new MDImplicitDeleteField initialized with val
func NewMDImplicitDelete(val quickfix.FIXBoolean) MDImplicitDeleteField {
	return MDImplicitDeleteField{val}
}

//MDMktField is a EXCHANGE field
type MDMktField struct{ quickfix.FIXString }

//Tag returns tag.MDMkt (275)
func (f MDMktField) Tag() quickfix.Tag { return tag.MDMkt }

//NewMDMkt returns a new MDMktField initialized with val
func NewMDMkt(val quickfix.FIXString) MDMktField {
	return MDMktField{val}
}

//MDOriginTypeField is a INT field
type MDOriginTypeField struct{ quickfix.FIXInt }

//Tag returns tag.MDOriginType (1024)
func (f MDOriginTypeField) Tag() quickfix.Tag { return tag.MDOriginType }

//NewMDOriginType returns a new MDOriginTypeField initialized with val
func NewMDOriginType(val quickfix.FIXInt) MDOriginTypeField {
	return MDOriginTypeField{val}
}

//MDPriceLevelField is a INT field
type MDPriceLevelField struct{ quickfix.FIXInt }

//Tag returns tag.MDPriceLevel (1023)
func (f MDPriceLevelField) Tag() quickfix.Tag { return tag.MDPriceLevel }

//NewMDPriceLevel returns a new MDPriceLevelField initialized with val
func NewMDPriceLevel(val quickfix.FIXInt) MDPriceLevelField {
	return MDPriceLevelField{val}
}

//MDQuoteTypeField is a INT field
type MDQuoteTypeField struct{ quickfix.FIXInt }

//Tag returns tag.MDQuoteType (1070)
func (f MDQuoteTypeField) Tag() quickfix.Tag { return tag.MDQuoteType }

//NewMDQuoteType returns a new MDQuoteTypeField initialized with val
func NewMDQuoteType(val quickfix.FIXInt) MDQuoteTypeField {
	return MDQuoteTypeField{val}
}

//MDReportIDField is a INT field
type MDReportIDField struct{ quickfix.FIXInt }

//Tag returns tag.MDReportID (963)
func (f MDReportIDField) Tag() quickfix.Tag { return tag.MDReportID }

//NewMDReportID returns a new MDReportIDField initialized with val
func NewMDReportID(val quickfix.FIXInt) MDReportIDField {
	return MDReportIDField{val}
}

//MDReqIDField is a STRING field
type MDReqIDField struct{ quickfix.FIXString }

//Tag returns tag.MDReqID (262)
func (f MDReqIDField) Tag() quickfix.Tag { return tag.MDReqID }

//NewMDReqID returns a new MDReqIDField initialized with val
func NewMDReqID(val quickfix.FIXString) MDReqIDField {
	return MDReqIDField{val}
}

//MDReqRejReasonField is a CHAR field
type MDReqRejReasonField struct{ quickfix.FIXString }

//Tag returns tag.MDReqRejReason (281)
func (f MDReqRejReasonField) Tag() quickfix.Tag { return tag.MDReqRejReason }

//NewMDReqRejReason returns a new MDReqRejReasonField initialized with val
func NewMDReqRejReason(val quickfix.FIXString) MDReqRejReasonField {
	return MDReqRejReasonField{val}
}

//MDSecSizeField is a QTY field
type MDSecSizeField struct{ quickfix.FIXFloat }

//Tag returns tag.MDSecSize (1179)
func (f MDSecSizeField) Tag() quickfix.Tag { return tag.MDSecSize }

//NewMDSecSize returns a new MDSecSizeField initialized with val
func NewMDSecSize(val quickfix.FIXFloat) MDSecSizeField {
	return MDSecSizeField{val}
}

//MDSecSizeTypeField is a INT field
type MDSecSizeTypeField struct{ quickfix.FIXInt }

//Tag returns tag.MDSecSizeType (1178)
func (f MDSecSizeTypeField) Tag() quickfix.Tag { return tag.MDSecSizeType }

//NewMDSecSizeType returns a new MDSecSizeTypeField initialized with val
func NewMDSecSizeType(val quickfix.FIXInt) MDSecSizeTypeField {
	return MDSecSizeTypeField{val}
}

//MDStreamIDField is a STRING field
type MDStreamIDField struct{ quickfix.FIXString }

//Tag returns tag.MDStreamID (1500)
func (f MDStreamIDField) Tag() quickfix.Tag { return tag.MDStreamID }

//NewMDStreamID returns a new MDStreamIDField initialized with val
func NewMDStreamID(val quickfix.FIXString) MDStreamIDField {
	return MDStreamIDField{val}
}

//MDSubBookTypeField is a INT field
type MDSubBookTypeField struct{ quickfix.FIXInt }

//Tag returns tag.MDSubBookType (1173)
func (f MDSubBookTypeField) Tag() quickfix.Tag { return tag.MDSubBookType }

//NewMDSubBookType returns a new MDSubBookTypeField initialized with val
func NewMDSubBookType(val quickfix.FIXInt) MDSubBookTypeField {
	return MDSubBookTypeField{val}
}

//MDUpdateActionField is a CHAR field
type MDUpdateActionField struct{ quickfix.FIXString }

//Tag returns tag.MDUpdateAction (279)
func (f MDUpdateActionField) Tag() quickfix.Tag { return tag.MDUpdateAction }

//NewMDUpdateAction returns a new MDUpdateActionField initialized with val
func NewMDUpdateAction(val quickfix.FIXString) MDUpdateActionField {
	return MDUpdateActionField{val}
}

//MDUpdateTypeField is a INT field
type MDUpdateTypeField struct{ quickfix.FIXInt }

//Tag returns tag.MDUpdateType (265)
func (f MDUpdateTypeField) Tag() quickfix.Tag { return tag.MDUpdateType }

//NewMDUpdateType returns a new MDUpdateTypeField initialized with val
func NewMDUpdateType(val quickfix.FIXInt) MDUpdateTypeField {
	return MDUpdateTypeField{val}
}

//MailingDtlsField is a STRING field
type MailingDtlsField struct{ quickfix.FIXString }

//Tag returns tag.MailingDtls (474)
func (f MailingDtlsField) Tag() quickfix.Tag { return tag.MailingDtls }

//NewMailingDtls returns a new MailingDtlsField initialized with val
func NewMailingDtls(val quickfix.FIXString) MailingDtlsField {
	return MailingDtlsField{val}
}

//MailingInstField is a STRING field
type MailingInstField struct{ quickfix.FIXString }

//Tag returns tag.MailingInst (482)
func (f MailingInstField) Tag() quickfix.Tag { return tag.MailingInst }

//NewMailingInst returns a new MailingInstField initialized with val
func NewMailingInst(val quickfix.FIXString) MailingInstField {
	return MailingInstField{val}
}

//ManualOrderIndicatorField is a BOOLEAN field
type ManualOrderIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.ManualOrderIndicator (1028)
func (f ManualOrderIndicatorField) Tag() quickfix.Tag { return tag.ManualOrderIndicator }

//NewManualOrderIndicator returns a new ManualOrderIndicatorField initialized with val
func NewManualOrderIndicator(val quickfix.FIXBoolean) ManualOrderIndicatorField {
	return ManualOrderIndicatorField{val}
}

//MarginExcessField is a AMT field
type MarginExcessField struct{ quickfix.FIXFloat }

//Tag returns tag.MarginExcess (899)
func (f MarginExcessField) Tag() quickfix.Tag { return tag.MarginExcess }

//NewMarginExcess returns a new MarginExcessField initialized with val
func NewMarginExcess(val quickfix.FIXFloat) MarginExcessField {
	return MarginExcessField{val}
}

//MarginRatioField is a PERCENTAGE field
type MarginRatioField struct{ quickfix.FIXFloat }

//Tag returns tag.MarginRatio (898)
func (f MarginRatioField) Tag() quickfix.Tag { return tag.MarginRatio }

//NewMarginRatio returns a new MarginRatioField initialized with val
func NewMarginRatio(val quickfix.FIXFloat) MarginRatioField {
	return MarginRatioField{val}
}

//MarketDepthField is a INT field
type MarketDepthField struct{ quickfix.FIXInt }

//Tag returns tag.MarketDepth (264)
func (f MarketDepthField) Tag() quickfix.Tag { return tag.MarketDepth }

//NewMarketDepth returns a new MarketDepthField initialized with val
func NewMarketDepth(val quickfix.FIXInt) MarketDepthField {
	return MarketDepthField{val}
}

//MarketIDField is a EXCHANGE field
type MarketIDField struct{ quickfix.FIXString }

//Tag returns tag.MarketID (1301)
func (f MarketIDField) Tag() quickfix.Tag { return tag.MarketID }

//NewMarketID returns a new MarketIDField initialized with val
func NewMarketID(val quickfix.FIXString) MarketIDField {
	return MarketIDField{val}
}

//MarketReportIDField is a STRING field
type MarketReportIDField struct{ quickfix.FIXString }

//Tag returns tag.MarketReportID (1394)
func (f MarketReportIDField) Tag() quickfix.Tag { return tag.MarketReportID }

//NewMarketReportID returns a new MarketReportIDField initialized with val
func NewMarketReportID(val quickfix.FIXString) MarketReportIDField {
	return MarketReportIDField{val}
}

//MarketReqIDField is a STRING field
type MarketReqIDField struct{ quickfix.FIXString }

//Tag returns tag.MarketReqID (1393)
func (f MarketReqIDField) Tag() quickfix.Tag { return tag.MarketReqID }

//NewMarketReqID returns a new MarketReqIDField initialized with val
func NewMarketReqID(val quickfix.FIXString) MarketReqIDField {
	return MarketReqIDField{val}
}

//MarketSegmentDescField is a STRING field
type MarketSegmentDescField struct{ quickfix.FIXString }

//Tag returns tag.MarketSegmentDesc (1396)
func (f MarketSegmentDescField) Tag() quickfix.Tag { return tag.MarketSegmentDesc }

//NewMarketSegmentDesc returns a new MarketSegmentDescField initialized with val
func NewMarketSegmentDesc(val quickfix.FIXString) MarketSegmentDescField {
	return MarketSegmentDescField{val}
}

//MarketSegmentIDField is a STRING field
type MarketSegmentIDField struct{ quickfix.FIXString }

//Tag returns tag.MarketSegmentID (1300)
func (f MarketSegmentIDField) Tag() quickfix.Tag { return tag.MarketSegmentID }

//NewMarketSegmentID returns a new MarketSegmentIDField initialized with val
func NewMarketSegmentID(val quickfix.FIXString) MarketSegmentIDField {
	return MarketSegmentIDField{val}
}

//MarketUpdateActionField is a CHAR field
type MarketUpdateActionField struct{ quickfix.FIXString }

//Tag returns tag.MarketUpdateAction (1395)
func (f MarketUpdateActionField) Tag() quickfix.Tag { return tag.MarketUpdateAction }

//NewMarketUpdateAction returns a new MarketUpdateActionField initialized with val
func NewMarketUpdateAction(val quickfix.FIXString) MarketUpdateActionField {
	return MarketUpdateActionField{val}
}

//MassActionRejectReasonField is a INT field
type MassActionRejectReasonField struct{ quickfix.FIXInt }

//Tag returns tag.MassActionRejectReason (1376)
func (f MassActionRejectReasonField) Tag() quickfix.Tag { return tag.MassActionRejectReason }

//NewMassActionRejectReason returns a new MassActionRejectReasonField initialized with val
func NewMassActionRejectReason(val quickfix.FIXInt) MassActionRejectReasonField {
	return MassActionRejectReasonField{val}
}

//MassActionReportIDField is a STRING field
type MassActionReportIDField struct{ quickfix.FIXString }

//Tag returns tag.MassActionReportID (1369)
func (f MassActionReportIDField) Tag() quickfix.Tag { return tag.MassActionReportID }

//NewMassActionReportID returns a new MassActionReportIDField initialized with val
func NewMassActionReportID(val quickfix.FIXString) MassActionReportIDField {
	return MassActionReportIDField{val}
}

//MassActionResponseField is a INT field
type MassActionResponseField struct{ quickfix.FIXInt }

//Tag returns tag.MassActionResponse (1375)
func (f MassActionResponseField) Tag() quickfix.Tag { return tag.MassActionResponse }

//NewMassActionResponse returns a new MassActionResponseField initialized with val
func NewMassActionResponse(val quickfix.FIXInt) MassActionResponseField {
	return MassActionResponseField{val}
}

//MassActionScopeField is a INT field
type MassActionScopeField struct{ quickfix.FIXInt }

//Tag returns tag.MassActionScope (1374)
func (f MassActionScopeField) Tag() quickfix.Tag { return tag.MassActionScope }

//NewMassActionScope returns a new MassActionScopeField initialized with val
func NewMassActionScope(val quickfix.FIXInt) MassActionScopeField {
	return MassActionScopeField{val}
}

//MassActionTypeField is a INT field
type MassActionTypeField struct{ quickfix.FIXInt }

//Tag returns tag.MassActionType (1373)
func (f MassActionTypeField) Tag() quickfix.Tag { return tag.MassActionType }

//NewMassActionType returns a new MassActionTypeField initialized with val
func NewMassActionType(val quickfix.FIXInt) MassActionTypeField {
	return MassActionTypeField{val}
}

//MassCancelRejectReasonField is a INT field
type MassCancelRejectReasonField struct{ quickfix.FIXInt }

//Tag returns tag.MassCancelRejectReason (532)
func (f MassCancelRejectReasonField) Tag() quickfix.Tag { return tag.MassCancelRejectReason }

//NewMassCancelRejectReason returns a new MassCancelRejectReasonField initialized with val
func NewMassCancelRejectReason(val quickfix.FIXInt) MassCancelRejectReasonField {
	return MassCancelRejectReasonField{val}
}

//MassCancelRequestTypeField is a CHAR field
type MassCancelRequestTypeField struct{ quickfix.FIXString }

//Tag returns tag.MassCancelRequestType (530)
func (f MassCancelRequestTypeField) Tag() quickfix.Tag { return tag.MassCancelRequestType }

//NewMassCancelRequestType returns a new MassCancelRequestTypeField initialized with val
func NewMassCancelRequestType(val quickfix.FIXString) MassCancelRequestTypeField {
	return MassCancelRequestTypeField{val}
}

//MassCancelResponseField is a CHAR field
type MassCancelResponseField struct{ quickfix.FIXString }

//Tag returns tag.MassCancelResponse (531)
func (f MassCancelResponseField) Tag() quickfix.Tag { return tag.MassCancelResponse }

//NewMassCancelResponse returns a new MassCancelResponseField initialized with val
func NewMassCancelResponse(val quickfix.FIXString) MassCancelResponseField {
	return MassCancelResponseField{val}
}

//MassStatusReqIDField is a STRING field
type MassStatusReqIDField struct{ quickfix.FIXString }

//Tag returns tag.MassStatusReqID (584)
func (f MassStatusReqIDField) Tag() quickfix.Tag { return tag.MassStatusReqID }

//NewMassStatusReqID returns a new MassStatusReqIDField initialized with val
func NewMassStatusReqID(val quickfix.FIXString) MassStatusReqIDField {
	return MassStatusReqIDField{val}
}

//MassStatusReqTypeField is a INT field
type MassStatusReqTypeField struct{ quickfix.FIXInt }

//Tag returns tag.MassStatusReqType (585)
func (f MassStatusReqTypeField) Tag() quickfix.Tag { return tag.MassStatusReqType }

//NewMassStatusReqType returns a new MassStatusReqTypeField initialized with val
func NewMassStatusReqType(val quickfix.FIXInt) MassStatusReqTypeField {
	return MassStatusReqTypeField{val}
}

//MatchAlgorithmField is a STRING field
type MatchAlgorithmField struct{ quickfix.FIXString }

//Tag returns tag.MatchAlgorithm (1142)
func (f MatchAlgorithmField) Tag() quickfix.Tag { return tag.MatchAlgorithm }

//NewMatchAlgorithm returns a new MatchAlgorithmField initialized with val
func NewMatchAlgorithm(val quickfix.FIXString) MatchAlgorithmField {
	return MatchAlgorithmField{val}
}

//MatchIncrementField is a QTY field
type MatchIncrementField struct{ quickfix.FIXFloat }

//Tag returns tag.MatchIncrement (1089)
func (f MatchIncrementField) Tag() quickfix.Tag { return tag.MatchIncrement }

//NewMatchIncrement returns a new MatchIncrementField initialized with val
func NewMatchIncrement(val quickfix.FIXFloat) MatchIncrementField {
	return MatchIncrementField{val}
}

//MatchStatusField is a CHAR field
type MatchStatusField struct{ quickfix.FIXString }

//Tag returns tag.MatchStatus (573)
func (f MatchStatusField) Tag() quickfix.Tag { return tag.MatchStatus }

//NewMatchStatus returns a new MatchStatusField initialized with val
func NewMatchStatus(val quickfix.FIXString) MatchStatusField {
	return MatchStatusField{val}
}

//MatchTypeField is a STRING field
type MatchTypeField struct{ quickfix.FIXString }

//Tag returns tag.MatchType (574)
func (f MatchTypeField) Tag() quickfix.Tag { return tag.MatchType }

//NewMatchType returns a new MatchTypeField initialized with val
func NewMatchType(val quickfix.FIXString) MatchTypeField {
	return MatchTypeField{val}
}

//MaturityDateField is a LOCALMKTDATE field
type MaturityDateField struct{ quickfix.FIXString }

//Tag returns tag.MaturityDate (541)
func (f MaturityDateField) Tag() quickfix.Tag { return tag.MaturityDate }

//NewMaturityDate returns a new MaturityDateField initialized with val
func NewMaturityDate(val quickfix.FIXString) MaturityDateField {
	return MaturityDateField{val}
}

//MaturityDayField is a DAYOFMONTH field
type MaturityDayField struct{ quickfix.FIXInt }

//Tag returns tag.MaturityDay (205)
func (f MaturityDayField) Tag() quickfix.Tag { return tag.MaturityDay }

//NewMaturityDay returns a new MaturityDayField initialized with val
func NewMaturityDay(val quickfix.FIXInt) MaturityDayField {
	return MaturityDayField{val}
}

//MaturityMonthYearField is a MONTHYEAR field
type MaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.MaturityMonthYear (200)
func (f MaturityMonthYearField) Tag() quickfix.Tag { return tag.MaturityMonthYear }

//NewMaturityMonthYear returns a new MaturityMonthYearField initialized with val
func NewMaturityMonthYear(val quickfix.FIXString) MaturityMonthYearField {
	return MaturityMonthYearField{val}
}

//MaturityMonthYearFormatField is a INT field
type MaturityMonthYearFormatField struct{ quickfix.FIXInt }

//Tag returns tag.MaturityMonthYearFormat (1303)
func (f MaturityMonthYearFormatField) Tag() quickfix.Tag { return tag.MaturityMonthYearFormat }

//NewMaturityMonthYearFormat returns a new MaturityMonthYearFormatField initialized with val
func NewMaturityMonthYearFormat(val quickfix.FIXInt) MaturityMonthYearFormatField {
	return MaturityMonthYearFormatField{val}
}

//MaturityMonthYearIncrementField is a INT field
type MaturityMonthYearIncrementField struct{ quickfix.FIXInt }

//Tag returns tag.MaturityMonthYearIncrement (1229)
func (f MaturityMonthYearIncrementField) Tag() quickfix.Tag { return tag.MaturityMonthYearIncrement }

//NewMaturityMonthYearIncrement returns a new MaturityMonthYearIncrementField initialized with val
func NewMaturityMonthYearIncrement(val quickfix.FIXInt) MaturityMonthYearIncrementField {
	return MaturityMonthYearIncrementField{val}
}

//MaturityMonthYearIncrementUnitsField is a INT field
type MaturityMonthYearIncrementUnitsField struct{ quickfix.FIXInt }

//Tag returns tag.MaturityMonthYearIncrementUnits (1302)
func (f MaturityMonthYearIncrementUnitsField) Tag() quickfix.Tag {
	return tag.MaturityMonthYearIncrementUnits
}

//NewMaturityMonthYearIncrementUnits returns a new MaturityMonthYearIncrementUnitsField initialized with val
func NewMaturityMonthYearIncrementUnits(val quickfix.FIXInt) MaturityMonthYearIncrementUnitsField {
	return MaturityMonthYearIncrementUnitsField{val}
}

//MaturityNetMoneyField is a AMT field
type MaturityNetMoneyField struct{ quickfix.FIXFloat }

//Tag returns tag.MaturityNetMoney (890)
func (f MaturityNetMoneyField) Tag() quickfix.Tag { return tag.MaturityNetMoney }

//NewMaturityNetMoney returns a new MaturityNetMoneyField initialized with val
func NewMaturityNetMoney(val quickfix.FIXFloat) MaturityNetMoneyField {
	return MaturityNetMoneyField{val}
}

//MaturityRuleIDField is a STRING field
type MaturityRuleIDField struct{ quickfix.FIXString }

//Tag returns tag.MaturityRuleID (1222)
func (f MaturityRuleIDField) Tag() quickfix.Tag { return tag.MaturityRuleID }

//NewMaturityRuleID returns a new MaturityRuleIDField initialized with val
func NewMaturityRuleID(val quickfix.FIXString) MaturityRuleIDField {
	return MaturityRuleIDField{val}
}

//MaturityTimeField is a TZTIMEONLY field
type MaturityTimeField struct{ quickfix.FIXString }

//Tag returns tag.MaturityTime (1079)
func (f MaturityTimeField) Tag() quickfix.Tag { return tag.MaturityTime }

//NewMaturityTime returns a new MaturityTimeField initialized with val
func NewMaturityTime(val quickfix.FIXString) MaturityTimeField {
	return MaturityTimeField{val}
}

//MaxFloorField is a QTY field
type MaxFloorField struct{ quickfix.FIXFloat }

//Tag returns tag.MaxFloor (111)
func (f MaxFloorField) Tag() quickfix.Tag { return tag.MaxFloor }

//NewMaxFloor returns a new MaxFloorField initialized with val
func NewMaxFloor(val quickfix.FIXFloat) MaxFloorField {
	return MaxFloorField{val}
}

//MaxMessageSizeField is a LENGTH field
type MaxMessageSizeField struct{ quickfix.FIXInt }

//Tag returns tag.MaxMessageSize (383)
func (f MaxMessageSizeField) Tag() quickfix.Tag { return tag.MaxMessageSize }

//NewMaxMessageSize returns a new MaxMessageSizeField initialized with val
func NewMaxMessageSize(val quickfix.FIXInt) MaxMessageSizeField {
	return MaxMessageSizeField{val}
}

//MaxPriceLevelsField is a INT field
type MaxPriceLevelsField struct{ quickfix.FIXInt }

//Tag returns tag.MaxPriceLevels (1090)
func (f MaxPriceLevelsField) Tag() quickfix.Tag { return tag.MaxPriceLevels }

//NewMaxPriceLevels returns a new MaxPriceLevelsField initialized with val
func NewMaxPriceLevels(val quickfix.FIXInt) MaxPriceLevelsField {
	return MaxPriceLevelsField{val}
}

//MaxPriceVariationField is a FLOAT field
type MaxPriceVariationField struct{ quickfix.FIXFloat }

//Tag returns tag.MaxPriceVariation (1143)
func (f MaxPriceVariationField) Tag() quickfix.Tag { return tag.MaxPriceVariation }

//NewMaxPriceVariation returns a new MaxPriceVariationField initialized with val
func NewMaxPriceVariation(val quickfix.FIXFloat) MaxPriceVariationField {
	return MaxPriceVariationField{val}
}

//MaxShowField is a QTY field
type MaxShowField struct{ quickfix.FIXFloat }

//Tag returns tag.MaxShow (210)
func (f MaxShowField) Tag() quickfix.Tag { return tag.MaxShow }

//NewMaxShow returns a new MaxShowField initialized with val
func NewMaxShow(val quickfix.FIXFloat) MaxShowField {
	return MaxShowField{val}
}

//MaxTradeVolField is a QTY field
type MaxTradeVolField struct{ quickfix.FIXFloat }

//Tag returns tag.MaxTradeVol (1140)
func (f MaxTradeVolField) Tag() quickfix.Tag { return tag.MaxTradeVol }

//NewMaxTradeVol returns a new MaxTradeVolField initialized with val
func NewMaxTradeVol(val quickfix.FIXFloat) MaxTradeVolField {
	return MaxTradeVolField{val}
}

//MessageEncodingField is a STRING field
type MessageEncodingField struct{ quickfix.FIXString }

//Tag returns tag.MessageEncoding (347)
func (f MessageEncodingField) Tag() quickfix.Tag { return tag.MessageEncoding }

//NewMessageEncoding returns a new MessageEncodingField initialized with val
func NewMessageEncoding(val quickfix.FIXString) MessageEncodingField {
	return MessageEncodingField{val}
}

//MessageEventSourceField is a STRING field
type MessageEventSourceField struct{ quickfix.FIXString }

//Tag returns tag.MessageEventSource (1011)
func (f MessageEventSourceField) Tag() quickfix.Tag { return tag.MessageEventSource }

//NewMessageEventSource returns a new MessageEventSourceField initialized with val
func NewMessageEventSource(val quickfix.FIXString) MessageEventSourceField {
	return MessageEventSourceField{val}
}

//MidPxField is a PRICE field
type MidPxField struct{ quickfix.FIXFloat }

//Tag returns tag.MidPx (631)
func (f MidPxField) Tag() quickfix.Tag { return tag.MidPx }

//NewMidPx returns a new MidPxField initialized with val
func NewMidPx(val quickfix.FIXFloat) MidPxField {
	return MidPxField{val}
}

//MidYieldField is a PERCENTAGE field
type MidYieldField struct{ quickfix.FIXFloat }

//Tag returns tag.MidYield (633)
func (f MidYieldField) Tag() quickfix.Tag { return tag.MidYield }

//NewMidYield returns a new MidYieldField initialized with val
func NewMidYield(val quickfix.FIXFloat) MidYieldField {
	return MidYieldField{val}
}

//MinBidSizeField is a QTY field
type MinBidSizeField struct{ quickfix.FIXFloat }

//Tag returns tag.MinBidSize (647)
func (f MinBidSizeField) Tag() quickfix.Tag { return tag.MinBidSize }

//NewMinBidSize returns a new MinBidSizeField initialized with val
func NewMinBidSize(val quickfix.FIXFloat) MinBidSizeField {
	return MinBidSizeField{val}
}

//MinLotSizeField is a QTY field
type MinLotSizeField struct{ quickfix.FIXFloat }

//Tag returns tag.MinLotSize (1231)
func (f MinLotSizeField) Tag() quickfix.Tag { return tag.MinLotSize }

//NewMinLotSize returns a new MinLotSizeField initialized with val
func NewMinLotSize(val quickfix.FIXFloat) MinLotSizeField {
	return MinLotSizeField{val}
}

//MinOfferSizeField is a QTY field
type MinOfferSizeField struct{ quickfix.FIXFloat }

//Tag returns tag.MinOfferSize (648)
func (f MinOfferSizeField) Tag() quickfix.Tag { return tag.MinOfferSize }

//NewMinOfferSize returns a new MinOfferSizeField initialized with val
func NewMinOfferSize(val quickfix.FIXFloat) MinOfferSizeField {
	return MinOfferSizeField{val}
}

//MinPriceIncrementField is a FLOAT field
type MinPriceIncrementField struct{ quickfix.FIXFloat }

//Tag returns tag.MinPriceIncrement (969)
func (f MinPriceIncrementField) Tag() quickfix.Tag { return tag.MinPriceIncrement }

//NewMinPriceIncrement returns a new MinPriceIncrementField initialized with val
func NewMinPriceIncrement(val quickfix.FIXFloat) MinPriceIncrementField {
	return MinPriceIncrementField{val}
}

//MinPriceIncrementAmountField is a AMT field
type MinPriceIncrementAmountField struct{ quickfix.FIXFloat }

//Tag returns tag.MinPriceIncrementAmount (1146)
func (f MinPriceIncrementAmountField) Tag() quickfix.Tag { return tag.MinPriceIncrementAmount }

//NewMinPriceIncrementAmount returns a new MinPriceIncrementAmountField initialized with val
func NewMinPriceIncrementAmount(val quickfix.FIXFloat) MinPriceIncrementAmountField {
	return MinPriceIncrementAmountField{val}
}

//MinQtyField is a QTY field
type MinQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.MinQty (110)
func (f MinQtyField) Tag() quickfix.Tag { return tag.MinQty }

//NewMinQty returns a new MinQtyField initialized with val
func NewMinQty(val quickfix.FIXFloat) MinQtyField {
	return MinQtyField{val}
}

//MinTradeVolField is a QTY field
type MinTradeVolField struct{ quickfix.FIXFloat }

//Tag returns tag.MinTradeVol (562)
func (f MinTradeVolField) Tag() quickfix.Tag { return tag.MinTradeVol }

//NewMinTradeVol returns a new MinTradeVolField initialized with val
func NewMinTradeVol(val quickfix.FIXFloat) MinTradeVolField {
	return MinTradeVolField{val}
}

//MiscFeeAmtField is a AMT field
type MiscFeeAmtField struct{ quickfix.FIXFloat }

//Tag returns tag.MiscFeeAmt (137)
func (f MiscFeeAmtField) Tag() quickfix.Tag { return tag.MiscFeeAmt }

//NewMiscFeeAmt returns a new MiscFeeAmtField initialized with val
func NewMiscFeeAmt(val quickfix.FIXFloat) MiscFeeAmtField {
	return MiscFeeAmtField{val}
}

//MiscFeeBasisField is a INT field
type MiscFeeBasisField struct{ quickfix.FIXInt }

//Tag returns tag.MiscFeeBasis (891)
func (f MiscFeeBasisField) Tag() quickfix.Tag { return tag.MiscFeeBasis }

//NewMiscFeeBasis returns a new MiscFeeBasisField initialized with val
func NewMiscFeeBasis(val quickfix.FIXInt) MiscFeeBasisField {
	return MiscFeeBasisField{val}
}

//MiscFeeCurrField is a CURRENCY field
type MiscFeeCurrField struct{ quickfix.FIXString }

//Tag returns tag.MiscFeeCurr (138)
func (f MiscFeeCurrField) Tag() quickfix.Tag { return tag.MiscFeeCurr }

//NewMiscFeeCurr returns a new MiscFeeCurrField initialized with val
func NewMiscFeeCurr(val quickfix.FIXString) MiscFeeCurrField {
	return MiscFeeCurrField{val}
}

//MiscFeeTypeField is a STRING field
type MiscFeeTypeField struct{ quickfix.FIXString }

//Tag returns tag.MiscFeeType (139)
func (f MiscFeeTypeField) Tag() quickfix.Tag { return tag.MiscFeeType }

//NewMiscFeeType returns a new MiscFeeTypeField initialized with val
func NewMiscFeeType(val quickfix.FIXString) MiscFeeTypeField {
	return MiscFeeTypeField{val}
}

//MktBidPxField is a PRICE field
type MktBidPxField struct{ quickfix.FIXFloat }

//Tag returns tag.MktBidPx (645)
func (f MktBidPxField) Tag() quickfix.Tag { return tag.MktBidPx }

//NewMktBidPx returns a new MktBidPxField initialized with val
func NewMktBidPx(val quickfix.FIXFloat) MktBidPxField {
	return MktBidPxField{val}
}

//MktOfferPxField is a PRICE field
type MktOfferPxField struct{ quickfix.FIXFloat }

//Tag returns tag.MktOfferPx (646)
func (f MktOfferPxField) Tag() quickfix.Tag { return tag.MktOfferPx }

//NewMktOfferPx returns a new MktOfferPxField initialized with val
func NewMktOfferPx(val quickfix.FIXFloat) MktOfferPxField {
	return MktOfferPxField{val}
}

//ModelTypeField is a INT field
type ModelTypeField struct{ quickfix.FIXInt }

//Tag returns tag.ModelType (1434)
func (f ModelTypeField) Tag() quickfix.Tag { return tag.ModelType }

//NewModelType returns a new ModelTypeField initialized with val
func NewModelType(val quickfix.FIXInt) ModelTypeField {
	return ModelTypeField{val}
}

//MoneyLaunderingStatusField is a CHAR field
type MoneyLaunderingStatusField struct{ quickfix.FIXString }

//Tag returns tag.MoneyLaunderingStatus (481)
func (f MoneyLaunderingStatusField) Tag() quickfix.Tag { return tag.MoneyLaunderingStatus }

//NewMoneyLaunderingStatus returns a new MoneyLaunderingStatusField initialized with val
func NewMoneyLaunderingStatus(val quickfix.FIXString) MoneyLaunderingStatusField {
	return MoneyLaunderingStatusField{val}
}

//MsgDirectionField is a CHAR field
type MsgDirectionField struct{ quickfix.FIXString }

//Tag returns tag.MsgDirection (385)
func (f MsgDirectionField) Tag() quickfix.Tag { return tag.MsgDirection }

//NewMsgDirection returns a new MsgDirectionField initialized with val
func NewMsgDirection(val quickfix.FIXString) MsgDirectionField {
	return MsgDirectionField{val}
}

//MsgSeqNumField is a SEQNUM field
type MsgSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.MsgSeqNum (34)
func (f MsgSeqNumField) Tag() quickfix.Tag { return tag.MsgSeqNum }

//NewMsgSeqNum returns a new MsgSeqNumField initialized with val
func NewMsgSeqNum(val quickfix.FIXInt) MsgSeqNumField {
	return MsgSeqNumField{val}
}

//MsgTypeField is a STRING field
type MsgTypeField struct{ quickfix.FIXString }

//Tag returns tag.MsgType (35)
func (f MsgTypeField) Tag() quickfix.Tag { return tag.MsgType }

//NewMsgType returns a new MsgTypeField initialized with val
func NewMsgType(val quickfix.FIXString) MsgTypeField {
	return MsgTypeField{val}
}

//MultiLegReportingTypeField is a CHAR field
type MultiLegReportingTypeField struct{ quickfix.FIXString }

//Tag returns tag.MultiLegReportingType (442)
func (f MultiLegReportingTypeField) Tag() quickfix.Tag { return tag.MultiLegReportingType }

//NewMultiLegReportingType returns a new MultiLegReportingTypeField initialized with val
func NewMultiLegReportingType(val quickfix.FIXString) MultiLegReportingTypeField {
	return MultiLegReportingTypeField{val}
}

//MultiLegRptTypeReqField is a INT field
type MultiLegRptTypeReqField struct{ quickfix.FIXInt }

//Tag returns tag.MultiLegRptTypeReq (563)
func (f MultiLegRptTypeReqField) Tag() quickfix.Tag { return tag.MultiLegRptTypeReq }

//NewMultiLegRptTypeReq returns a new MultiLegRptTypeReqField initialized with val
func NewMultiLegRptTypeReq(val quickfix.FIXInt) MultiLegRptTypeReqField {
	return MultiLegRptTypeReqField{val}
}

//MultilegModelField is a INT field
type MultilegModelField struct{ quickfix.FIXInt }

//Tag returns tag.MultilegModel (1377)
func (f MultilegModelField) Tag() quickfix.Tag { return tag.MultilegModel }

//NewMultilegModel returns a new MultilegModelField initialized with val
func NewMultilegModel(val quickfix.FIXInt) MultilegModelField {
	return MultilegModelField{val}
}

//MultilegPriceMethodField is a INT field
type MultilegPriceMethodField struct{ quickfix.FIXInt }

//Tag returns tag.MultilegPriceMethod (1378)
func (f MultilegPriceMethodField) Tag() quickfix.Tag { return tag.MultilegPriceMethod }

//NewMultilegPriceMethod returns a new MultilegPriceMethodField initialized with val
func NewMultilegPriceMethod(val quickfix.FIXInt) MultilegPriceMethodField {
	return MultilegPriceMethodField{val}
}

//NTPositionLimitField is a INT field
type NTPositionLimitField struct{ quickfix.FIXInt }

//Tag returns tag.NTPositionLimit (971)
func (f NTPositionLimitField) Tag() quickfix.Tag { return tag.NTPositionLimit }

//NewNTPositionLimit returns a new NTPositionLimitField initialized with val
func NewNTPositionLimit(val quickfix.FIXInt) NTPositionLimitField {
	return NTPositionLimitField{val}
}

//Nested2PartyIDField is a STRING field
type Nested2PartyIDField struct{ quickfix.FIXString }

//Tag returns tag.Nested2PartyID (757)
func (f Nested2PartyIDField) Tag() quickfix.Tag { return tag.Nested2PartyID }

//NewNested2PartyID returns a new Nested2PartyIDField initialized with val
func NewNested2PartyID(val quickfix.FIXString) Nested2PartyIDField {
	return Nested2PartyIDField{val}
}

//Nested2PartyIDSourceField is a CHAR field
type Nested2PartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.Nested2PartyIDSource (758)
func (f Nested2PartyIDSourceField) Tag() quickfix.Tag { return tag.Nested2PartyIDSource }

//NewNested2PartyIDSource returns a new Nested2PartyIDSourceField initialized with val
func NewNested2PartyIDSource(val quickfix.FIXString) Nested2PartyIDSourceField {
	return Nested2PartyIDSourceField{val}
}

//Nested2PartyRoleField is a INT field
type Nested2PartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.Nested2PartyRole (759)
func (f Nested2PartyRoleField) Tag() quickfix.Tag { return tag.Nested2PartyRole }

//NewNested2PartyRole returns a new Nested2PartyRoleField initialized with val
func NewNested2PartyRole(val quickfix.FIXInt) Nested2PartyRoleField {
	return Nested2PartyRoleField{val}
}

//Nested2PartySubIDField is a STRING field
type Nested2PartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.Nested2PartySubID (760)
func (f Nested2PartySubIDField) Tag() quickfix.Tag { return tag.Nested2PartySubID }

//NewNested2PartySubID returns a new Nested2PartySubIDField initialized with val
func NewNested2PartySubID(val quickfix.FIXString) Nested2PartySubIDField {
	return Nested2PartySubIDField{val}
}

//Nested2PartySubIDTypeField is a INT field
type Nested2PartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.Nested2PartySubIDType (807)
func (f Nested2PartySubIDTypeField) Tag() quickfix.Tag { return tag.Nested2PartySubIDType }

//NewNested2PartySubIDType returns a new Nested2PartySubIDTypeField initialized with val
func NewNested2PartySubIDType(val quickfix.FIXInt) Nested2PartySubIDTypeField {
	return Nested2PartySubIDTypeField{val}
}

//Nested3PartyIDField is a STRING field
type Nested3PartyIDField struct{ quickfix.FIXString }

//Tag returns tag.Nested3PartyID (949)
func (f Nested3PartyIDField) Tag() quickfix.Tag { return tag.Nested3PartyID }

//NewNested3PartyID returns a new Nested3PartyIDField initialized with val
func NewNested3PartyID(val quickfix.FIXString) Nested3PartyIDField {
	return Nested3PartyIDField{val}
}

//Nested3PartyIDSourceField is a CHAR field
type Nested3PartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.Nested3PartyIDSource (950)
func (f Nested3PartyIDSourceField) Tag() quickfix.Tag { return tag.Nested3PartyIDSource }

//NewNested3PartyIDSource returns a new Nested3PartyIDSourceField initialized with val
func NewNested3PartyIDSource(val quickfix.FIXString) Nested3PartyIDSourceField {
	return Nested3PartyIDSourceField{val}
}

//Nested3PartyRoleField is a INT field
type Nested3PartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.Nested3PartyRole (951)
func (f Nested3PartyRoleField) Tag() quickfix.Tag { return tag.Nested3PartyRole }

//NewNested3PartyRole returns a new Nested3PartyRoleField initialized with val
func NewNested3PartyRole(val quickfix.FIXInt) Nested3PartyRoleField {
	return Nested3PartyRoleField{val}
}

//Nested3PartySubIDField is a STRING field
type Nested3PartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.Nested3PartySubID (953)
func (f Nested3PartySubIDField) Tag() quickfix.Tag { return tag.Nested3PartySubID }

//NewNested3PartySubID returns a new Nested3PartySubIDField initialized with val
func NewNested3PartySubID(val quickfix.FIXString) Nested3PartySubIDField {
	return Nested3PartySubIDField{val}
}

//Nested3PartySubIDTypeField is a INT field
type Nested3PartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.Nested3PartySubIDType (954)
func (f Nested3PartySubIDTypeField) Tag() quickfix.Tag { return tag.Nested3PartySubIDType }

//NewNested3PartySubIDType returns a new Nested3PartySubIDTypeField initialized with val
func NewNested3PartySubIDType(val quickfix.FIXInt) Nested3PartySubIDTypeField {
	return Nested3PartySubIDTypeField{val}
}

//Nested4PartyIDField is a STRING field
type Nested4PartyIDField struct{ quickfix.FIXString }

//Tag returns tag.Nested4PartyID (1415)
func (f Nested4PartyIDField) Tag() quickfix.Tag { return tag.Nested4PartyID }

//NewNested4PartyID returns a new Nested4PartyIDField initialized with val
func NewNested4PartyID(val quickfix.FIXString) Nested4PartyIDField {
	return Nested4PartyIDField{val}
}

//Nested4PartyIDSourceField is a CHAR field
type Nested4PartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.Nested4PartyIDSource (1416)
func (f Nested4PartyIDSourceField) Tag() quickfix.Tag { return tag.Nested4PartyIDSource }

//NewNested4PartyIDSource returns a new Nested4PartyIDSourceField initialized with val
func NewNested4PartyIDSource(val quickfix.FIXString) Nested4PartyIDSourceField {
	return Nested4PartyIDSourceField{val}
}

//Nested4PartyRoleField is a INT field
type Nested4PartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.Nested4PartyRole (1417)
func (f Nested4PartyRoleField) Tag() quickfix.Tag { return tag.Nested4PartyRole }

//NewNested4PartyRole returns a new Nested4PartyRoleField initialized with val
func NewNested4PartyRole(val quickfix.FIXInt) Nested4PartyRoleField {
	return Nested4PartyRoleField{val}
}

//Nested4PartySubIDField is a STRING field
type Nested4PartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.Nested4PartySubID (1412)
func (f Nested4PartySubIDField) Tag() quickfix.Tag { return tag.Nested4PartySubID }

//NewNested4PartySubID returns a new Nested4PartySubIDField initialized with val
func NewNested4PartySubID(val quickfix.FIXString) Nested4PartySubIDField {
	return Nested4PartySubIDField{val}
}

//Nested4PartySubIDTypeField is a INT field
type Nested4PartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.Nested4PartySubIDType (1411)
func (f Nested4PartySubIDTypeField) Tag() quickfix.Tag { return tag.Nested4PartySubIDType }

//NewNested4PartySubIDType returns a new Nested4PartySubIDTypeField initialized with val
func NewNested4PartySubIDType(val quickfix.FIXInt) Nested4PartySubIDTypeField {
	return Nested4PartySubIDTypeField{val}
}

//NestedInstrAttribTypeField is a INT field
type NestedInstrAttribTypeField struct{ quickfix.FIXInt }

//Tag returns tag.NestedInstrAttribType (1210)
func (f NestedInstrAttribTypeField) Tag() quickfix.Tag { return tag.NestedInstrAttribType }

//NewNestedInstrAttribType returns a new NestedInstrAttribTypeField initialized with val
func NewNestedInstrAttribType(val quickfix.FIXInt) NestedInstrAttribTypeField {
	return NestedInstrAttribTypeField{val}
}

//NestedInstrAttribValueField is a STRING field
type NestedInstrAttribValueField struct{ quickfix.FIXString }

//Tag returns tag.NestedInstrAttribValue (1211)
func (f NestedInstrAttribValueField) Tag() quickfix.Tag { return tag.NestedInstrAttribValue }

//NewNestedInstrAttribValue returns a new NestedInstrAttribValueField initialized with val
func NewNestedInstrAttribValue(val quickfix.FIXString) NestedInstrAttribValueField {
	return NestedInstrAttribValueField{val}
}

//NestedPartyIDField is a STRING field
type NestedPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.NestedPartyID (524)
func (f NestedPartyIDField) Tag() quickfix.Tag { return tag.NestedPartyID }

//NewNestedPartyID returns a new NestedPartyIDField initialized with val
func NewNestedPartyID(val quickfix.FIXString) NestedPartyIDField {
	return NestedPartyIDField{val}
}

//NestedPartyIDSourceField is a CHAR field
type NestedPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.NestedPartyIDSource (525)
func (f NestedPartyIDSourceField) Tag() quickfix.Tag { return tag.NestedPartyIDSource }

//NewNestedPartyIDSource returns a new NestedPartyIDSourceField initialized with val
func NewNestedPartyIDSource(val quickfix.FIXString) NestedPartyIDSourceField {
	return NestedPartyIDSourceField{val}
}

//NestedPartyRoleField is a INT field
type NestedPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.NestedPartyRole (538)
func (f NestedPartyRoleField) Tag() quickfix.Tag { return tag.NestedPartyRole }

//NewNestedPartyRole returns a new NestedPartyRoleField initialized with val
func NewNestedPartyRole(val quickfix.FIXInt) NestedPartyRoleField {
	return NestedPartyRoleField{val}
}

//NestedPartySubIDField is a STRING field
type NestedPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.NestedPartySubID (545)
func (f NestedPartySubIDField) Tag() quickfix.Tag { return tag.NestedPartySubID }

//NewNestedPartySubID returns a new NestedPartySubIDField initialized with val
func NewNestedPartySubID(val quickfix.FIXString) NestedPartySubIDField {
	return NestedPartySubIDField{val}
}

//NestedPartySubIDTypeField is a INT field
type NestedPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.NestedPartySubIDType (805)
func (f NestedPartySubIDTypeField) Tag() quickfix.Tag { return tag.NestedPartySubIDType }

//NewNestedPartySubIDType returns a new NestedPartySubIDTypeField initialized with val
func NewNestedPartySubIDType(val quickfix.FIXInt) NestedPartySubIDTypeField {
	return NestedPartySubIDTypeField{val}
}

//NetChgPrevDayField is a PRICEOFFSET field
type NetChgPrevDayField struct{ quickfix.FIXFloat }

//Tag returns tag.NetChgPrevDay (451)
func (f NetChgPrevDayField) Tag() quickfix.Tag { return tag.NetChgPrevDay }

//NewNetChgPrevDay returns a new NetChgPrevDayField initialized with val
func NewNetChgPrevDay(val quickfix.FIXFloat) NetChgPrevDayField {
	return NetChgPrevDayField{val}
}

//NetGrossIndField is a INT field
type NetGrossIndField struct{ quickfix.FIXInt }

//Tag returns tag.NetGrossInd (430)
func (f NetGrossIndField) Tag() quickfix.Tag { return tag.NetGrossInd }

//NewNetGrossInd returns a new NetGrossIndField initialized with val
func NewNetGrossInd(val quickfix.FIXInt) NetGrossIndField {
	return NetGrossIndField{val}
}

//NetMoneyField is a AMT field
type NetMoneyField struct{ quickfix.FIXFloat }

//Tag returns tag.NetMoney (118)
func (f NetMoneyField) Tag() quickfix.Tag { return tag.NetMoney }

//NewNetMoney returns a new NetMoneyField initialized with val
func NewNetMoney(val quickfix.FIXFloat) NetMoneyField {
	return NetMoneyField{val}
}

//NetworkRequestIDField is a STRING field
type NetworkRequestIDField struct{ quickfix.FIXString }

//Tag returns tag.NetworkRequestID (933)
func (f NetworkRequestIDField) Tag() quickfix.Tag { return tag.NetworkRequestID }

//NewNetworkRequestID returns a new NetworkRequestIDField initialized with val
func NewNetworkRequestID(val quickfix.FIXString) NetworkRequestIDField {
	return NetworkRequestIDField{val}
}

//NetworkRequestTypeField is a INT field
type NetworkRequestTypeField struct{ quickfix.FIXInt }

//Tag returns tag.NetworkRequestType (935)
func (f NetworkRequestTypeField) Tag() quickfix.Tag { return tag.NetworkRequestType }

//NewNetworkRequestType returns a new NetworkRequestTypeField initialized with val
func NewNetworkRequestType(val quickfix.FIXInt) NetworkRequestTypeField {
	return NetworkRequestTypeField{val}
}

//NetworkResponseIDField is a STRING field
type NetworkResponseIDField struct{ quickfix.FIXString }

//Tag returns tag.NetworkResponseID (932)
func (f NetworkResponseIDField) Tag() quickfix.Tag { return tag.NetworkResponseID }

//NewNetworkResponseID returns a new NetworkResponseIDField initialized with val
func NewNetworkResponseID(val quickfix.FIXString) NetworkResponseIDField {
	return NetworkResponseIDField{val}
}

//NetworkStatusResponseTypeField is a INT field
type NetworkStatusResponseTypeField struct{ quickfix.FIXInt }

//Tag returns tag.NetworkStatusResponseType (937)
func (f NetworkStatusResponseTypeField) Tag() quickfix.Tag { return tag.NetworkStatusResponseType }

//NewNetworkStatusResponseType returns a new NetworkStatusResponseTypeField initialized with val
func NewNetworkStatusResponseType(val quickfix.FIXInt) NetworkStatusResponseTypeField {
	return NetworkStatusResponseTypeField{val}
}

//NewPasswordField is a STRING field
type NewPasswordField struct{ quickfix.FIXString }

//Tag returns tag.NewPassword (925)
func (f NewPasswordField) Tag() quickfix.Tag { return tag.NewPassword }

//NewNewPassword returns a new NewPasswordField initialized with val
func NewNewPassword(val quickfix.FIXString) NewPasswordField {
	return NewPasswordField{val}
}

//NewSeqNoField is a SEQNUM field
type NewSeqNoField struct{ quickfix.FIXInt }

//Tag returns tag.NewSeqNo (36)
func (f NewSeqNoField) Tag() quickfix.Tag { return tag.NewSeqNo }

//NewNewSeqNo returns a new NewSeqNoField initialized with val
func NewNewSeqNo(val quickfix.FIXInt) NewSeqNoField {
	return NewSeqNoField{val}
}

//NewsCategoryField is a INT field
type NewsCategoryField struct{ quickfix.FIXInt }

//Tag returns tag.NewsCategory (1473)
func (f NewsCategoryField) Tag() quickfix.Tag { return tag.NewsCategory }

//NewNewsCategory returns a new NewsCategoryField initialized with val
func NewNewsCategory(val quickfix.FIXInt) NewsCategoryField {
	return NewsCategoryField{val}
}

//NewsIDField is a STRING field
type NewsIDField struct{ quickfix.FIXString }

//Tag returns tag.NewsID (1472)
func (f NewsIDField) Tag() quickfix.Tag { return tag.NewsID }

//NewNewsID returns a new NewsIDField initialized with val
func NewNewsID(val quickfix.FIXString) NewsIDField {
	return NewsIDField{val}
}

//NewsRefIDField is a STRING field
type NewsRefIDField struct{ quickfix.FIXString }

//Tag returns tag.NewsRefID (1476)
func (f NewsRefIDField) Tag() quickfix.Tag { return tag.NewsRefID }

//NewNewsRefID returns a new NewsRefIDField initialized with val
func NewNewsRefID(val quickfix.FIXString) NewsRefIDField {
	return NewsRefIDField{val}
}

//NewsRefTypeField is a INT field
type NewsRefTypeField struct{ quickfix.FIXInt }

//Tag returns tag.NewsRefType (1477)
func (f NewsRefTypeField) Tag() quickfix.Tag { return tag.NewsRefType }

//NewNewsRefType returns a new NewsRefTypeField initialized with val
func NewNewsRefType(val quickfix.FIXInt) NewsRefTypeField {
	return NewsRefTypeField{val}
}

//NextExpectedMsgSeqNumField is a SEQNUM field
type NextExpectedMsgSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.NextExpectedMsgSeqNum (789)
func (f NextExpectedMsgSeqNumField) Tag() quickfix.Tag { return tag.NextExpectedMsgSeqNum }

//NewNextExpectedMsgSeqNum returns a new NextExpectedMsgSeqNumField initialized with val
func NewNextExpectedMsgSeqNum(val quickfix.FIXInt) NextExpectedMsgSeqNumField {
	return NextExpectedMsgSeqNumField{val}
}

//NoAffectedOrdersField is a NUMINGROUP field
type NoAffectedOrdersField struct{ quickfix.FIXInt }

//Tag returns tag.NoAffectedOrders (534)
func (f NoAffectedOrdersField) Tag() quickfix.Tag { return tag.NoAffectedOrders }

//NewNoAffectedOrders returns a new NoAffectedOrdersField initialized with val
func NewNoAffectedOrders(val quickfix.FIXInt) NoAffectedOrdersField {
	return NoAffectedOrdersField{val}
}

//NoAllocsField is a NUMINGROUP field
type NoAllocsField struct{ quickfix.FIXInt }

//Tag returns tag.NoAllocs (78)
func (f NoAllocsField) Tag() quickfix.Tag { return tag.NoAllocs }

//NewNoAllocs returns a new NoAllocsField initialized with val
func NewNoAllocs(val quickfix.FIXInt) NoAllocsField {
	return NoAllocsField{val}
}

//NoAltMDSourceField is a NUMINGROUP field
type NoAltMDSourceField struct{ quickfix.FIXInt }

//Tag returns tag.NoAltMDSource (816)
func (f NoAltMDSourceField) Tag() quickfix.Tag { return tag.NoAltMDSource }

//NewNoAltMDSource returns a new NoAltMDSourceField initialized with val
func NewNoAltMDSource(val quickfix.FIXInt) NoAltMDSourceField {
	return NoAltMDSourceField{val}
}

//NoApplIDsField is a NUMINGROUP field
type NoApplIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoApplIDs (1351)
func (f NoApplIDsField) Tag() quickfix.Tag { return tag.NoApplIDs }

//NewNoApplIDs returns a new NoApplIDsField initialized with val
func NewNoApplIDs(val quickfix.FIXInt) NoApplIDsField {
	return NoApplIDsField{val}
}

//NoAsgnReqsField is a NUMINGROUP field
type NoAsgnReqsField struct{ quickfix.FIXInt }

//Tag returns tag.NoAsgnReqs (1499)
func (f NoAsgnReqsField) Tag() quickfix.Tag { return tag.NoAsgnReqs }

//NewNoAsgnReqs returns a new NoAsgnReqsField initialized with val
func NewNoAsgnReqs(val quickfix.FIXInt) NoAsgnReqsField {
	return NoAsgnReqsField{val}
}

//NoBidComponentsField is a NUMINGROUP field
type NoBidComponentsField struct{ quickfix.FIXInt }

//Tag returns tag.NoBidComponents (420)
func (f NoBidComponentsField) Tag() quickfix.Tag { return tag.NoBidComponents }

//NewNoBidComponents returns a new NoBidComponentsField initialized with val
func NewNoBidComponents(val quickfix.FIXInt) NoBidComponentsField {
	return NoBidComponentsField{val}
}

//NoBidDescriptorsField is a NUMINGROUP field
type NoBidDescriptorsField struct{ quickfix.FIXInt }

//Tag returns tag.NoBidDescriptors (398)
func (f NoBidDescriptorsField) Tag() quickfix.Tag { return tag.NoBidDescriptors }

//NewNoBidDescriptors returns a new NoBidDescriptorsField initialized with val
func NewNoBidDescriptors(val quickfix.FIXInt) NoBidDescriptorsField {
	return NoBidDescriptorsField{val}
}

//NoCapacitiesField is a NUMINGROUP field
type NoCapacitiesField struct{ quickfix.FIXInt }

//Tag returns tag.NoCapacities (862)
func (f NoCapacitiesField) Tag() quickfix.Tag { return tag.NoCapacities }

//NewNoCapacities returns a new NoCapacitiesField initialized with val
func NewNoCapacities(val quickfix.FIXInt) NoCapacitiesField {
	return NoCapacitiesField{val}
}

//NoClearingInstructionsField is a NUMINGROUP field
type NoClearingInstructionsField struct{ quickfix.FIXInt }

//Tag returns tag.NoClearingInstructions (576)
func (f NoClearingInstructionsField) Tag() quickfix.Tag { return tag.NoClearingInstructions }

//NewNoClearingInstructions returns a new NoClearingInstructionsField initialized with val
func NewNoClearingInstructions(val quickfix.FIXInt) NoClearingInstructionsField {
	return NoClearingInstructionsField{val}
}

//NoCollInquiryQualifierField is a NUMINGROUP field
type NoCollInquiryQualifierField struct{ quickfix.FIXInt }

//Tag returns tag.NoCollInquiryQualifier (938)
func (f NoCollInquiryQualifierField) Tag() quickfix.Tag { return tag.NoCollInquiryQualifier }

//NewNoCollInquiryQualifier returns a new NoCollInquiryQualifierField initialized with val
func NewNoCollInquiryQualifier(val quickfix.FIXInt) NoCollInquiryQualifierField {
	return NoCollInquiryQualifierField{val}
}

//NoCompIDsField is a NUMINGROUP field
type NoCompIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoCompIDs (936)
func (f NoCompIDsField) Tag() quickfix.Tag { return tag.NoCompIDs }

//NewNoCompIDs returns a new NoCompIDsField initialized with val
func NewNoCompIDs(val quickfix.FIXInt) NoCompIDsField {
	return NoCompIDsField{val}
}

//NoComplexEventDatesField is a NUMINGROUP field
type NoComplexEventDatesField struct{ quickfix.FIXInt }

//Tag returns tag.NoComplexEventDates (1491)
func (f NoComplexEventDatesField) Tag() quickfix.Tag { return tag.NoComplexEventDates }

//NewNoComplexEventDates returns a new NoComplexEventDatesField initialized with val
func NewNoComplexEventDates(val quickfix.FIXInt) NoComplexEventDatesField {
	return NoComplexEventDatesField{val}
}

//NoComplexEventTimesField is a NUMINGROUP field
type NoComplexEventTimesField struct{ quickfix.FIXInt }

//Tag returns tag.NoComplexEventTimes (1494)
func (f NoComplexEventTimesField) Tag() quickfix.Tag { return tag.NoComplexEventTimes }

//NewNoComplexEventTimes returns a new NoComplexEventTimesField initialized with val
func NewNoComplexEventTimes(val quickfix.FIXInt) NoComplexEventTimesField {
	return NoComplexEventTimesField{val}
}

//NoComplexEventsField is a NUMINGROUP field
type NoComplexEventsField struct{ quickfix.FIXInt }

//Tag returns tag.NoComplexEvents (1483)
func (f NoComplexEventsField) Tag() quickfix.Tag { return tag.NoComplexEvents }

//NewNoComplexEvents returns a new NoComplexEventsField initialized with val
func NewNoComplexEvents(val quickfix.FIXInt) NoComplexEventsField {
	return NoComplexEventsField{val}
}

//NoContAmtsField is a NUMINGROUP field
type NoContAmtsField struct{ quickfix.FIXInt }

//Tag returns tag.NoContAmts (518)
func (f NoContAmtsField) Tag() quickfix.Tag { return tag.NoContAmts }

//NewNoContAmts returns a new NoContAmtsField initialized with val
func NewNoContAmts(val quickfix.FIXInt) NoContAmtsField {
	return NoContAmtsField{val}
}

//NoContextPartyIDsField is a NUMINGROUP field
type NoContextPartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoContextPartyIDs (1522)
func (f NoContextPartyIDsField) Tag() quickfix.Tag { return tag.NoContextPartyIDs }

//NewNoContextPartyIDs returns a new NoContextPartyIDsField initialized with val
func NewNoContextPartyIDs(val quickfix.FIXInt) NoContextPartyIDsField {
	return NoContextPartyIDsField{val}
}

//NoContextPartySubIDsField is a NUMINGROUP field
type NoContextPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoContextPartySubIDs (1526)
func (f NoContextPartySubIDsField) Tag() quickfix.Tag { return tag.NoContextPartySubIDs }

//NewNoContextPartySubIDs returns a new NoContextPartySubIDsField initialized with val
func NewNoContextPartySubIDs(val quickfix.FIXInt) NoContextPartySubIDsField {
	return NoContextPartySubIDsField{val}
}

//NoContraBrokersField is a NUMINGROUP field
type NoContraBrokersField struct{ quickfix.FIXInt }

//Tag returns tag.NoContraBrokers (382)
func (f NoContraBrokersField) Tag() quickfix.Tag { return tag.NoContraBrokers }

//NewNoContraBrokers returns a new NoContraBrokersField initialized with val
func NewNoContraBrokers(val quickfix.FIXInt) NoContraBrokersField {
	return NoContraBrokersField{val}
}

//NoDatesField is a INT field
type NoDatesField struct{ quickfix.FIXInt }

//Tag returns tag.NoDates (580)
func (f NoDatesField) Tag() quickfix.Tag { return tag.NoDates }

//NewNoDates returns a new NoDatesField initialized with val
func NewNoDates(val quickfix.FIXInt) NoDatesField {
	return NoDatesField{val}
}

//NoDerivativeEventsField is a NUMINGROUP field
type NoDerivativeEventsField struct{ quickfix.FIXInt }

//Tag returns tag.NoDerivativeEvents (1286)
func (f NoDerivativeEventsField) Tag() quickfix.Tag { return tag.NoDerivativeEvents }

//NewNoDerivativeEvents returns a new NoDerivativeEventsField initialized with val
func NewNoDerivativeEvents(val quickfix.FIXInt) NoDerivativeEventsField {
	return NoDerivativeEventsField{val}
}

//NoDerivativeInstrAttribField is a NUMINGROUP field
type NoDerivativeInstrAttribField struct{ quickfix.FIXInt }

//Tag returns tag.NoDerivativeInstrAttrib (1311)
func (f NoDerivativeInstrAttribField) Tag() quickfix.Tag { return tag.NoDerivativeInstrAttrib }

//NewNoDerivativeInstrAttrib returns a new NoDerivativeInstrAttribField initialized with val
func NewNoDerivativeInstrAttrib(val quickfix.FIXInt) NoDerivativeInstrAttribField {
	return NoDerivativeInstrAttribField{val}
}

//NoDerivativeInstrumentPartiesField is a NUMINGROUP field
type NoDerivativeInstrumentPartiesField struct{ quickfix.FIXInt }

//Tag returns tag.NoDerivativeInstrumentParties (1292)
func (f NoDerivativeInstrumentPartiesField) Tag() quickfix.Tag {
	return tag.NoDerivativeInstrumentParties
}

//NewNoDerivativeInstrumentParties returns a new NoDerivativeInstrumentPartiesField initialized with val
func NewNoDerivativeInstrumentParties(val quickfix.FIXInt) NoDerivativeInstrumentPartiesField {
	return NoDerivativeInstrumentPartiesField{val}
}

//NoDerivativeInstrumentPartySubIDsField is a NUMINGROUP field
type NoDerivativeInstrumentPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoDerivativeInstrumentPartySubIDs (1296)
func (f NoDerivativeInstrumentPartySubIDsField) Tag() quickfix.Tag {
	return tag.NoDerivativeInstrumentPartySubIDs
}

//NewNoDerivativeInstrumentPartySubIDs returns a new NoDerivativeInstrumentPartySubIDsField initialized with val
func NewNoDerivativeInstrumentPartySubIDs(val quickfix.FIXInt) NoDerivativeInstrumentPartySubIDsField {
	return NoDerivativeInstrumentPartySubIDsField{val}
}

//NoDerivativeSecurityAltIDField is a NUMINGROUP field
type NoDerivativeSecurityAltIDField struct{ quickfix.FIXInt }

//Tag returns tag.NoDerivativeSecurityAltID (1218)
func (f NoDerivativeSecurityAltIDField) Tag() quickfix.Tag { return tag.NoDerivativeSecurityAltID }

//NewNoDerivativeSecurityAltID returns a new NoDerivativeSecurityAltIDField initialized with val
func NewNoDerivativeSecurityAltID(val quickfix.FIXInt) NoDerivativeSecurityAltIDField {
	return NoDerivativeSecurityAltIDField{val}
}

//NoDistribInstsField is a NUMINGROUP field
type NoDistribInstsField struct{ quickfix.FIXInt }

//Tag returns tag.NoDistribInsts (510)
func (f NoDistribInstsField) Tag() quickfix.Tag { return tag.NoDistribInsts }

//NewNoDistribInsts returns a new NoDistribInstsField initialized with val
func NewNoDistribInsts(val quickfix.FIXInt) NoDistribInstsField {
	return NoDistribInstsField{val}
}

//NoDlvyInstField is a NUMINGROUP field
type NoDlvyInstField struct{ quickfix.FIXInt }

//Tag returns tag.NoDlvyInst (85)
func (f NoDlvyInstField) Tag() quickfix.Tag { return tag.NoDlvyInst }

//NewNoDlvyInst returns a new NoDlvyInstField initialized with val
func NewNoDlvyInst(val quickfix.FIXInt) NoDlvyInstField {
	return NoDlvyInstField{val}
}

//NoEventsField is a NUMINGROUP field
type NoEventsField struct{ quickfix.FIXInt }

//Tag returns tag.NoEvents (864)
func (f NoEventsField) Tag() quickfix.Tag { return tag.NoEvents }

//NewNoEvents returns a new NoEventsField initialized with val
func NewNoEvents(val quickfix.FIXInt) NoEventsField {
	return NoEventsField{val}
}

//NoExecInstRulesField is a NUMINGROUP field
type NoExecInstRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoExecInstRules (1232)
func (f NoExecInstRulesField) Tag() quickfix.Tag { return tag.NoExecInstRules }

//NewNoExecInstRules returns a new NoExecInstRulesField initialized with val
func NewNoExecInstRules(val quickfix.FIXInt) NoExecInstRulesField {
	return NoExecInstRulesField{val}
}

//NoExecsField is a NUMINGROUP field
type NoExecsField struct{ quickfix.FIXInt }

//Tag returns tag.NoExecs (124)
func (f NoExecsField) Tag() quickfix.Tag { return tag.NoExecs }

//NewNoExecs returns a new NoExecsField initialized with val
func NewNoExecs(val quickfix.FIXInt) NoExecsField {
	return NoExecsField{val}
}

//NoExpirationField is a NUMINGROUP field
type NoExpirationField struct{ quickfix.FIXInt }

//Tag returns tag.NoExpiration (981)
func (f NoExpirationField) Tag() quickfix.Tag { return tag.NoExpiration }

//NewNoExpiration returns a new NoExpirationField initialized with val
func NewNoExpiration(val quickfix.FIXInt) NoExpirationField {
	return NoExpirationField{val}
}

//NoFillsField is a NUMINGROUP field
type NoFillsField struct{ quickfix.FIXInt }

//Tag returns tag.NoFills (1362)
func (f NoFillsField) Tag() quickfix.Tag { return tag.NoFills }

//NewNoFills returns a new NoFillsField initialized with val
func NewNoFills(val quickfix.FIXInt) NoFillsField {
	return NoFillsField{val}
}

//NoHopsField is a NUMINGROUP field
type NoHopsField struct{ quickfix.FIXInt }

//Tag returns tag.NoHops (627)
func (f NoHopsField) Tag() quickfix.Tag { return tag.NoHops }

//NewNoHops returns a new NoHopsField initialized with val
func NewNoHops(val quickfix.FIXInt) NoHopsField {
	return NoHopsField{val}
}

//NoIOIQualifiersField is a NUMINGROUP field
type NoIOIQualifiersField struct{ quickfix.FIXInt }

//Tag returns tag.NoIOIQualifiers (199)
func (f NoIOIQualifiersField) Tag() quickfix.Tag { return tag.NoIOIQualifiers }

//NewNoIOIQualifiers returns a new NoIOIQualifiersField initialized with val
func NewNoIOIQualifiers(val quickfix.FIXInt) NoIOIQualifiersField {
	return NoIOIQualifiersField{val}
}

//NoInstrAttribField is a NUMINGROUP field
type NoInstrAttribField struct{ quickfix.FIXInt }

//Tag returns tag.NoInstrAttrib (870)
func (f NoInstrAttribField) Tag() quickfix.Tag { return tag.NoInstrAttrib }

//NewNoInstrAttrib returns a new NoInstrAttribField initialized with val
func NewNoInstrAttrib(val quickfix.FIXInt) NoInstrAttribField {
	return NoInstrAttribField{val}
}

//NoInstrumentPartiesField is a NUMINGROUP field
type NoInstrumentPartiesField struct{ quickfix.FIXInt }

//Tag returns tag.NoInstrumentParties (1018)
func (f NoInstrumentPartiesField) Tag() quickfix.Tag { return tag.NoInstrumentParties }

//NewNoInstrumentParties returns a new NoInstrumentPartiesField initialized with val
func NewNoInstrumentParties(val quickfix.FIXInt) NoInstrumentPartiesField {
	return NoInstrumentPartiesField{val}
}

//NoInstrumentPartySubIDsField is a NUMINGROUP field
type NoInstrumentPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoInstrumentPartySubIDs (1052)
func (f NoInstrumentPartySubIDsField) Tag() quickfix.Tag { return tag.NoInstrumentPartySubIDs }

//NewNoInstrumentPartySubIDs returns a new NoInstrumentPartySubIDsField initialized with val
func NewNoInstrumentPartySubIDs(val quickfix.FIXInt) NoInstrumentPartySubIDsField {
	return NoInstrumentPartySubIDsField{val}
}

//NoLegAllocsField is a NUMINGROUP field
type NoLegAllocsField struct{ quickfix.FIXInt }

//Tag returns tag.NoLegAllocs (670)
func (f NoLegAllocsField) Tag() quickfix.Tag { return tag.NoLegAllocs }

//NewNoLegAllocs returns a new NoLegAllocsField initialized with val
func NewNoLegAllocs(val quickfix.FIXInt) NoLegAllocsField {
	return NoLegAllocsField{val}
}

//NoLegSecurityAltIDField is a STRING field
type NoLegSecurityAltIDField struct{ quickfix.FIXString }

//Tag returns tag.NoLegSecurityAltID (604)
func (f NoLegSecurityAltIDField) Tag() quickfix.Tag { return tag.NoLegSecurityAltID }

//NewNoLegSecurityAltID returns a new NoLegSecurityAltIDField initialized with val
func NewNoLegSecurityAltID(val quickfix.FIXString) NoLegSecurityAltIDField {
	return NoLegSecurityAltIDField{val}
}

//NoLegStipulationsField is a NUMINGROUP field
type NoLegStipulationsField struct{ quickfix.FIXInt }

//Tag returns tag.NoLegStipulations (683)
func (f NoLegStipulationsField) Tag() quickfix.Tag { return tag.NoLegStipulations }

//NewNoLegStipulations returns a new NoLegStipulationsField initialized with val
func NewNoLegStipulations(val quickfix.FIXInt) NoLegStipulationsField {
	return NoLegStipulationsField{val}
}

//NoLegsField is a NUMINGROUP field
type NoLegsField struct{ quickfix.FIXInt }

//Tag returns tag.NoLegs (555)
func (f NoLegsField) Tag() quickfix.Tag { return tag.NoLegs }

//NewNoLegs returns a new NoLegsField initialized with val
func NewNoLegs(val quickfix.FIXInt) NoLegsField {
	return NoLegsField{val}
}

//NoLinesOfTextField is a NUMINGROUP field
type NoLinesOfTextField struct{ quickfix.FIXInt }

//Tag returns tag.NoLinesOfText (33)
func (f NoLinesOfTextField) Tag() quickfix.Tag { return tag.NoLinesOfText }

//NewNoLinesOfText returns a new NoLinesOfTextField initialized with val
func NewNoLinesOfText(val quickfix.FIXInt) NoLinesOfTextField {
	return NoLinesOfTextField{val}
}

//NoLotTypeRulesField is a NUMINGROUP field
type NoLotTypeRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoLotTypeRules (1234)
func (f NoLotTypeRulesField) Tag() quickfix.Tag { return tag.NoLotTypeRules }

//NewNoLotTypeRules returns a new NoLotTypeRulesField initialized with val
func NewNoLotTypeRules(val quickfix.FIXInt) NoLotTypeRulesField {
	return NoLotTypeRulesField{val}
}

//NoMDEntriesField is a NUMINGROUP field
type NoMDEntriesField struct{ quickfix.FIXInt }

//Tag returns tag.NoMDEntries (268)
func (f NoMDEntriesField) Tag() quickfix.Tag { return tag.NoMDEntries }

//NewNoMDEntries returns a new NoMDEntriesField initialized with val
func NewNoMDEntries(val quickfix.FIXInt) NoMDEntriesField {
	return NoMDEntriesField{val}
}

//NoMDEntryTypesField is a NUMINGROUP field
type NoMDEntryTypesField struct{ quickfix.FIXInt }

//Tag returns tag.NoMDEntryTypes (267)
func (f NoMDEntryTypesField) Tag() quickfix.Tag { return tag.NoMDEntryTypes }

//NewNoMDEntryTypes returns a new NoMDEntryTypesField initialized with val
func NewNoMDEntryTypes(val quickfix.FIXInt) NoMDEntryTypesField {
	return NoMDEntryTypesField{val}
}

//NoMDFeedTypesField is a NUMINGROUP field
type NoMDFeedTypesField struct{ quickfix.FIXInt }

//Tag returns tag.NoMDFeedTypes (1141)
func (f NoMDFeedTypesField) Tag() quickfix.Tag { return tag.NoMDFeedTypes }

//NewNoMDFeedTypes returns a new NoMDFeedTypesField initialized with val
func NewNoMDFeedTypes(val quickfix.FIXInt) NoMDFeedTypesField {
	return NoMDFeedTypesField{val}
}

//NoMarketSegmentsField is a NUMINGROUP field
type NoMarketSegmentsField struct{ quickfix.FIXInt }

//Tag returns tag.NoMarketSegments (1310)
func (f NoMarketSegmentsField) Tag() quickfix.Tag { return tag.NoMarketSegments }

//NewNoMarketSegments returns a new NoMarketSegmentsField initialized with val
func NewNoMarketSegments(val quickfix.FIXInt) NoMarketSegmentsField {
	return NoMarketSegmentsField{val}
}

//NoMatchRulesField is a NUMINGROUP field
type NoMatchRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoMatchRules (1235)
func (f NoMatchRulesField) Tag() quickfix.Tag { return tag.NoMatchRules }

//NewNoMatchRules returns a new NoMatchRulesField initialized with val
func NewNoMatchRules(val quickfix.FIXInt) NoMatchRulesField {
	return NoMatchRulesField{val}
}

//NoMaturityRulesField is a NUMINGROUP field
type NoMaturityRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoMaturityRules (1236)
func (f NoMaturityRulesField) Tag() quickfix.Tag { return tag.NoMaturityRules }

//NewNoMaturityRules returns a new NoMaturityRulesField initialized with val
func NewNoMaturityRules(val quickfix.FIXInt) NoMaturityRulesField {
	return NoMaturityRulesField{val}
}

//NoMiscFeesField is a NUMINGROUP field
type NoMiscFeesField struct{ quickfix.FIXInt }

//Tag returns tag.NoMiscFees (136)
func (f NoMiscFeesField) Tag() quickfix.Tag { return tag.NoMiscFees }

//NewNoMiscFees returns a new NoMiscFeesField initialized with val
func NewNoMiscFees(val quickfix.FIXInt) NoMiscFeesField {
	return NoMiscFeesField{val}
}

//NoMsgTypesField is a NUMINGROUP field
type NoMsgTypesField struct{ quickfix.FIXInt }

//Tag returns tag.NoMsgTypes (384)
func (f NoMsgTypesField) Tag() quickfix.Tag { return tag.NoMsgTypes }

//NewNoMsgTypes returns a new NoMsgTypesField initialized with val
func NewNoMsgTypes(val quickfix.FIXInt) NoMsgTypesField {
	return NoMsgTypesField{val}
}

//NoNested2PartyIDsField is a NUMINGROUP field
type NoNested2PartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNested2PartyIDs (756)
func (f NoNested2PartyIDsField) Tag() quickfix.Tag { return tag.NoNested2PartyIDs }

//NewNoNested2PartyIDs returns a new NoNested2PartyIDsField initialized with val
func NewNoNested2PartyIDs(val quickfix.FIXInt) NoNested2PartyIDsField {
	return NoNested2PartyIDsField{val}
}

//NoNested2PartySubIDsField is a NUMINGROUP field
type NoNested2PartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNested2PartySubIDs (806)
func (f NoNested2PartySubIDsField) Tag() quickfix.Tag { return tag.NoNested2PartySubIDs }

//NewNoNested2PartySubIDs returns a new NoNested2PartySubIDsField initialized with val
func NewNoNested2PartySubIDs(val quickfix.FIXInt) NoNested2PartySubIDsField {
	return NoNested2PartySubIDsField{val}
}

//NoNested3PartyIDsField is a NUMINGROUP field
type NoNested3PartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNested3PartyIDs (948)
func (f NoNested3PartyIDsField) Tag() quickfix.Tag { return tag.NoNested3PartyIDs }

//NewNoNested3PartyIDs returns a new NoNested3PartyIDsField initialized with val
func NewNoNested3PartyIDs(val quickfix.FIXInt) NoNested3PartyIDsField {
	return NoNested3PartyIDsField{val}
}

//NoNested3PartySubIDsField is a NUMINGROUP field
type NoNested3PartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNested3PartySubIDs (952)
func (f NoNested3PartySubIDsField) Tag() quickfix.Tag { return tag.NoNested3PartySubIDs }

//NewNoNested3PartySubIDs returns a new NoNested3PartySubIDsField initialized with val
func NewNoNested3PartySubIDs(val quickfix.FIXInt) NoNested3PartySubIDsField {
	return NoNested3PartySubIDsField{val}
}

//NoNested4PartyIDsField is a NUMINGROUP field
type NoNested4PartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNested4PartyIDs (1414)
func (f NoNested4PartyIDsField) Tag() quickfix.Tag { return tag.NoNested4PartyIDs }

//NewNoNested4PartyIDs returns a new NoNested4PartyIDsField initialized with val
func NewNoNested4PartyIDs(val quickfix.FIXInt) NoNested4PartyIDsField {
	return NoNested4PartyIDsField{val}
}

//NoNested4PartySubIDsField is a NUMINGROUP field
type NoNested4PartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNested4PartySubIDs (1413)
func (f NoNested4PartySubIDsField) Tag() quickfix.Tag { return tag.NoNested4PartySubIDs }

//NewNoNested4PartySubIDs returns a new NoNested4PartySubIDsField initialized with val
func NewNoNested4PartySubIDs(val quickfix.FIXInt) NoNested4PartySubIDsField {
	return NoNested4PartySubIDsField{val}
}

//NoNestedInstrAttribField is a NUMINGROUP field
type NoNestedInstrAttribField struct{ quickfix.FIXInt }

//Tag returns tag.NoNestedInstrAttrib (1312)
func (f NoNestedInstrAttribField) Tag() quickfix.Tag { return tag.NoNestedInstrAttrib }

//NewNoNestedInstrAttrib returns a new NoNestedInstrAttribField initialized with val
func NewNoNestedInstrAttrib(val quickfix.FIXInt) NoNestedInstrAttribField {
	return NoNestedInstrAttribField{val}
}

//NoNestedPartyIDsField is a NUMINGROUP field
type NoNestedPartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNestedPartyIDs (539)
func (f NoNestedPartyIDsField) Tag() quickfix.Tag { return tag.NoNestedPartyIDs }

//NewNoNestedPartyIDs returns a new NoNestedPartyIDsField initialized with val
func NewNoNestedPartyIDs(val quickfix.FIXInt) NoNestedPartyIDsField {
	return NoNestedPartyIDsField{val}
}

//NoNestedPartySubIDsField is a NUMINGROUP field
type NoNestedPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNestedPartySubIDs (804)
func (f NoNestedPartySubIDsField) Tag() quickfix.Tag { return tag.NoNestedPartySubIDs }

//NewNoNestedPartySubIDs returns a new NoNestedPartySubIDsField initialized with val
func NewNoNestedPartySubIDs(val quickfix.FIXInt) NoNestedPartySubIDsField {
	return NoNestedPartySubIDsField{val}
}

//NoNewsRefIDsField is a NUMINGROUP field
type NoNewsRefIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoNewsRefIDs (1475)
func (f NoNewsRefIDsField) Tag() quickfix.Tag { return tag.NoNewsRefIDs }

//NewNoNewsRefIDs returns a new NoNewsRefIDsField initialized with val
func NewNoNewsRefIDs(val quickfix.FIXInt) NoNewsRefIDsField {
	return NoNewsRefIDsField{val}
}

//NoNotAffectedOrdersField is a NUMINGROUP field
type NoNotAffectedOrdersField struct{ quickfix.FIXInt }

//Tag returns tag.NoNotAffectedOrders (1370)
func (f NoNotAffectedOrdersField) Tag() quickfix.Tag { return tag.NoNotAffectedOrders }

//NewNoNotAffectedOrders returns a new NoNotAffectedOrdersField initialized with val
func NewNoNotAffectedOrders(val quickfix.FIXInt) NoNotAffectedOrdersField {
	return NoNotAffectedOrdersField{val}
}

//NoOfLegUnderlyingsField is a NUMINGROUP field
type NoOfLegUnderlyingsField struct{ quickfix.FIXInt }

//Tag returns tag.NoOfLegUnderlyings (1342)
func (f NoOfLegUnderlyingsField) Tag() quickfix.Tag { return tag.NoOfLegUnderlyings }

//NewNoOfLegUnderlyings returns a new NoOfLegUnderlyingsField initialized with val
func NewNoOfLegUnderlyings(val quickfix.FIXInt) NoOfLegUnderlyingsField {
	return NoOfLegUnderlyingsField{val}
}

//NoOfSecSizesField is a NUMINGROUP field
type NoOfSecSizesField struct{ quickfix.FIXInt }

//Tag returns tag.NoOfSecSizes (1177)
func (f NoOfSecSizesField) Tag() quickfix.Tag { return tag.NoOfSecSizes }

//NewNoOfSecSizes returns a new NoOfSecSizesField initialized with val
func NewNoOfSecSizes(val quickfix.FIXInt) NoOfSecSizesField {
	return NoOfSecSizesField{val}
}

//NoOrdTypeRulesField is a NUMINGROUP field
type NoOrdTypeRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoOrdTypeRules (1237)
func (f NoOrdTypeRulesField) Tag() quickfix.Tag { return tag.NoOrdTypeRules }

//NewNoOrdTypeRules returns a new NoOrdTypeRulesField initialized with val
func NewNoOrdTypeRules(val quickfix.FIXInt) NoOrdTypeRulesField {
	return NoOrdTypeRulesField{val}
}

//NoOrdersField is a NUMINGROUP field
type NoOrdersField struct{ quickfix.FIXInt }

//Tag returns tag.NoOrders (73)
func (f NoOrdersField) Tag() quickfix.Tag { return tag.NoOrders }

//NewNoOrders returns a new NoOrdersField initialized with val
func NewNoOrders(val quickfix.FIXInt) NoOrdersField {
	return NoOrdersField{val}
}

//NoPartyAltIDsField is a NUMINGROUP field
type NoPartyAltIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoPartyAltIDs (1516)
func (f NoPartyAltIDsField) Tag() quickfix.Tag { return tag.NoPartyAltIDs }

//NewNoPartyAltIDs returns a new NoPartyAltIDsField initialized with val
func NewNoPartyAltIDs(val quickfix.FIXInt) NoPartyAltIDsField {
	return NoPartyAltIDsField{val}
}

//NoPartyAltSubIDsField is a NUMINGROUP field
type NoPartyAltSubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoPartyAltSubIDs (1519)
func (f NoPartyAltSubIDsField) Tag() quickfix.Tag { return tag.NoPartyAltSubIDs }

//NewNoPartyAltSubIDs returns a new NoPartyAltSubIDsField initialized with val
func NewNoPartyAltSubIDs(val quickfix.FIXInt) NoPartyAltSubIDsField {
	return NoPartyAltSubIDsField{val}
}

//NoPartyIDsField is a NUMINGROUP field
type NoPartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoPartyIDs (453)
func (f NoPartyIDsField) Tag() quickfix.Tag { return tag.NoPartyIDs }

//NewNoPartyIDs returns a new NoPartyIDsField initialized with val
func NewNoPartyIDs(val quickfix.FIXInt) NoPartyIDsField {
	return NoPartyIDsField{val}
}

//NoPartyListField is a NUMINGROUP field
type NoPartyListField struct{ quickfix.FIXInt }

//Tag returns tag.NoPartyList (1513)
func (f NoPartyListField) Tag() quickfix.Tag { return tag.NoPartyList }

//NewNoPartyList returns a new NoPartyListField initialized with val
func NewNoPartyList(val quickfix.FIXInt) NoPartyListField {
	return NoPartyListField{val}
}

//NoPartyListResponseTypesField is a NUMINGROUP field
type NoPartyListResponseTypesField struct{ quickfix.FIXInt }

//Tag returns tag.NoPartyListResponseTypes (1506)
func (f NoPartyListResponseTypesField) Tag() quickfix.Tag { return tag.NoPartyListResponseTypes }

//NewNoPartyListResponseTypes returns a new NoPartyListResponseTypesField initialized with val
func NewNoPartyListResponseTypes(val quickfix.FIXInt) NoPartyListResponseTypesField {
	return NoPartyListResponseTypesField{val}
}

//NoPartyRelationshipsField is a NUMINGROUP field
type NoPartyRelationshipsField struct{ quickfix.FIXInt }

//Tag returns tag.NoPartyRelationships (1514)
func (f NoPartyRelationshipsField) Tag() quickfix.Tag { return tag.NoPartyRelationships }

//NewNoPartyRelationships returns a new NoPartyRelationshipsField initialized with val
func NewNoPartyRelationships(val quickfix.FIXInt) NoPartyRelationshipsField {
	return NoPartyRelationshipsField{val}
}

//NoPartySubIDsField is a NUMINGROUP field
type NoPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoPartySubIDs (802)
func (f NoPartySubIDsField) Tag() quickfix.Tag { return tag.NoPartySubIDs }

//NewNoPartySubIDs returns a new NoPartySubIDsField initialized with val
func NewNoPartySubIDs(val quickfix.FIXInt) NoPartySubIDsField {
	return NoPartySubIDsField{val}
}

//NoPosAmtField is a NUMINGROUP field
type NoPosAmtField struct{ quickfix.FIXInt }

//Tag returns tag.NoPosAmt (753)
func (f NoPosAmtField) Tag() quickfix.Tag { return tag.NoPosAmt }

//NewNoPosAmt returns a new NoPosAmtField initialized with val
func NewNoPosAmt(val quickfix.FIXInt) NoPosAmtField {
	return NoPosAmtField{val}
}

//NoPositionsField is a NUMINGROUP field
type NoPositionsField struct{ quickfix.FIXInt }

//Tag returns tag.NoPositions (702)
func (f NoPositionsField) Tag() quickfix.Tag { return tag.NoPositions }

//NewNoPositions returns a new NoPositionsField initialized with val
func NewNoPositions(val quickfix.FIXInt) NoPositionsField {
	return NoPositionsField{val}
}

//NoQuoteEntriesField is a NUMINGROUP field
type NoQuoteEntriesField struct{ quickfix.FIXInt }

//Tag returns tag.NoQuoteEntries (295)
func (f NoQuoteEntriesField) Tag() quickfix.Tag { return tag.NoQuoteEntries }

//NewNoQuoteEntries returns a new NoQuoteEntriesField initialized with val
func NewNoQuoteEntries(val quickfix.FIXInt) NoQuoteEntriesField {
	return NoQuoteEntriesField{val}
}

//NoQuoteQualifiersField is a NUMINGROUP field
type NoQuoteQualifiersField struct{ quickfix.FIXInt }

//Tag returns tag.NoQuoteQualifiers (735)
func (f NoQuoteQualifiersField) Tag() quickfix.Tag { return tag.NoQuoteQualifiers }

//NewNoQuoteQualifiers returns a new NoQuoteQualifiersField initialized with val
func NewNoQuoteQualifiers(val quickfix.FIXInt) NoQuoteQualifiersField {
	return NoQuoteQualifiersField{val}
}

//NoQuoteSetsField is a NUMINGROUP field
type NoQuoteSetsField struct{ quickfix.FIXInt }

//Tag returns tag.NoQuoteSets (296)
func (f NoQuoteSetsField) Tag() quickfix.Tag { return tag.NoQuoteSets }

//NewNoQuoteSets returns a new NoQuoteSetsField initialized with val
func NewNoQuoteSets(val quickfix.FIXInt) NoQuoteSetsField {
	return NoQuoteSetsField{val}
}

//NoRateSourcesField is a NUMINGROUP field
type NoRateSourcesField struct{ quickfix.FIXInt }

//Tag returns tag.NoRateSources (1445)
func (f NoRateSourcesField) Tag() quickfix.Tag { return tag.NoRateSources }

//NewNoRateSources returns a new NoRateSourcesField initialized with val
func NewNoRateSources(val quickfix.FIXInt) NoRateSourcesField {
	return NoRateSourcesField{val}
}

//NoRegistDtlsField is a NUMINGROUP field
type NoRegistDtlsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRegistDtls (473)
func (f NoRegistDtlsField) Tag() quickfix.Tag { return tag.NoRegistDtls }

//NewNoRegistDtls returns a new NoRegistDtlsField initialized with val
func NewNoRegistDtls(val quickfix.FIXInt) NoRegistDtlsField {
	return NoRegistDtlsField{val}
}

//NoRelatedContextPartyIDsField is a NUMINGROUP field
type NoRelatedContextPartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelatedContextPartyIDs (1575)
func (f NoRelatedContextPartyIDsField) Tag() quickfix.Tag { return tag.NoRelatedContextPartyIDs }

//NewNoRelatedContextPartyIDs returns a new NoRelatedContextPartyIDsField initialized with val
func NewNoRelatedContextPartyIDs(val quickfix.FIXInt) NoRelatedContextPartyIDsField {
	return NoRelatedContextPartyIDsField{val}
}

//NoRelatedContextPartySubIDsField is a NUMINGROUP field
type NoRelatedContextPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelatedContextPartySubIDs (1579)
func (f NoRelatedContextPartySubIDsField) Tag() quickfix.Tag { return tag.NoRelatedContextPartySubIDs }

//NewNoRelatedContextPartySubIDs returns a new NoRelatedContextPartySubIDsField initialized with val
func NewNoRelatedContextPartySubIDs(val quickfix.FIXInt) NoRelatedContextPartySubIDsField {
	return NoRelatedContextPartySubIDsField{val}
}

//NoRelatedPartyAltIDsField is a NUMINGROUP field
type NoRelatedPartyAltIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelatedPartyAltIDs (1569)
func (f NoRelatedPartyAltIDsField) Tag() quickfix.Tag { return tag.NoRelatedPartyAltIDs }

//NewNoRelatedPartyAltIDs returns a new NoRelatedPartyAltIDsField initialized with val
func NewNoRelatedPartyAltIDs(val quickfix.FIXInt) NoRelatedPartyAltIDsField {
	return NoRelatedPartyAltIDsField{val}
}

//NoRelatedPartyAltSubIDsField is a NUMINGROUP field
type NoRelatedPartyAltSubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelatedPartyAltSubIDs (1572)
func (f NoRelatedPartyAltSubIDsField) Tag() quickfix.Tag { return tag.NoRelatedPartyAltSubIDs }

//NewNoRelatedPartyAltSubIDs returns a new NoRelatedPartyAltSubIDsField initialized with val
func NewNoRelatedPartyAltSubIDs(val quickfix.FIXInt) NoRelatedPartyAltSubIDsField {
	return NoRelatedPartyAltSubIDsField{val}
}

//NoRelatedPartyIDsField is a NUMINGROUP field
type NoRelatedPartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelatedPartyIDs (1562)
func (f NoRelatedPartyIDsField) Tag() quickfix.Tag { return tag.NoRelatedPartyIDs }

//NewNoRelatedPartyIDs returns a new NoRelatedPartyIDsField initialized with val
func NewNoRelatedPartyIDs(val quickfix.FIXInt) NoRelatedPartyIDsField {
	return NoRelatedPartyIDsField{val}
}

//NoRelatedPartySubIDsField is a NUMINGROUP field
type NoRelatedPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelatedPartySubIDs (1566)
func (f NoRelatedPartySubIDsField) Tag() quickfix.Tag { return tag.NoRelatedPartySubIDs }

//NewNoRelatedPartySubIDs returns a new NoRelatedPartySubIDsField initialized with val
func NewNoRelatedPartySubIDs(val quickfix.FIXInt) NoRelatedPartySubIDsField {
	return NoRelatedPartySubIDsField{val}
}

//NoRelatedSymField is a NUMINGROUP field
type NoRelatedSymField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelatedSym (146)
func (f NoRelatedSymField) Tag() quickfix.Tag { return tag.NoRelatedSym }

//NewNoRelatedSym returns a new NoRelatedSymField initialized with val
func NewNoRelatedSym(val quickfix.FIXInt) NoRelatedSymField {
	return NoRelatedSymField{val}
}

//NoRelationshipRiskInstrumentsField is a NUMINGROUP field
type NoRelationshipRiskInstrumentsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelationshipRiskInstruments (1587)
func (f NoRelationshipRiskInstrumentsField) Tag() quickfix.Tag {
	return tag.NoRelationshipRiskInstruments
}

//NewNoRelationshipRiskInstruments returns a new NoRelationshipRiskInstrumentsField initialized with val
func NewNoRelationshipRiskInstruments(val quickfix.FIXInt) NoRelationshipRiskInstrumentsField {
	return NoRelationshipRiskInstrumentsField{val}
}

//NoRelationshipRiskLimitsField is a NUMINGROUP field
type NoRelationshipRiskLimitsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelationshipRiskLimits (1582)
func (f NoRelationshipRiskLimitsField) Tag() quickfix.Tag { return tag.NoRelationshipRiskLimits }

//NewNoRelationshipRiskLimits returns a new NoRelationshipRiskLimitsField initialized with val
func NewNoRelationshipRiskLimits(val quickfix.FIXInt) NoRelationshipRiskLimitsField {
	return NoRelationshipRiskLimitsField{val}
}

//NoRelationshipRiskSecurityAltIDField is a NUMINGROUP field
type NoRelationshipRiskSecurityAltIDField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelationshipRiskSecurityAltID (1593)
func (f NoRelationshipRiskSecurityAltIDField) Tag() quickfix.Tag {
	return tag.NoRelationshipRiskSecurityAltID
}

//NewNoRelationshipRiskSecurityAltID returns a new NoRelationshipRiskSecurityAltIDField initialized with val
func NewNoRelationshipRiskSecurityAltID(val quickfix.FIXInt) NoRelationshipRiskSecurityAltIDField {
	return NoRelationshipRiskSecurityAltIDField{val}
}

//NoRelationshipRiskWarningLevelsField is a NUMINGROUP field
type NoRelationshipRiskWarningLevelsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRelationshipRiskWarningLevels (1613)
func (f NoRelationshipRiskWarningLevelsField) Tag() quickfix.Tag {
	return tag.NoRelationshipRiskWarningLevels
}

//NewNoRelationshipRiskWarningLevels returns a new NoRelationshipRiskWarningLevelsField initialized with val
func NewNoRelationshipRiskWarningLevels(val quickfix.FIXInt) NoRelationshipRiskWarningLevelsField {
	return NoRelationshipRiskWarningLevelsField{val}
}

//NoRequestedPartyRolesField is a NUMINGROUP field
type NoRequestedPartyRolesField struct{ quickfix.FIXInt }

//Tag returns tag.NoRequestedPartyRoles (1508)
func (f NoRequestedPartyRolesField) Tag() quickfix.Tag { return tag.NoRequestedPartyRoles }

//NewNoRequestedPartyRoles returns a new NoRequestedPartyRolesField initialized with val
func NewNoRequestedPartyRoles(val quickfix.FIXInt) NoRequestedPartyRolesField {
	return NoRequestedPartyRolesField{val}
}

//NoRiskInstrumentsField is a NUMINGROUP field
type NoRiskInstrumentsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRiskInstruments (1534)
func (f NoRiskInstrumentsField) Tag() quickfix.Tag { return tag.NoRiskInstruments }

//NewNoRiskInstruments returns a new NoRiskInstrumentsField initialized with val
func NewNoRiskInstruments(val quickfix.FIXInt) NoRiskInstrumentsField {
	return NoRiskInstrumentsField{val}
}

//NoRiskLimitsField is a NUMINGROUP field
type NoRiskLimitsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRiskLimits (1529)
func (f NoRiskLimitsField) Tag() quickfix.Tag { return tag.NoRiskLimits }

//NewNoRiskLimits returns a new NoRiskLimitsField initialized with val
func NewNoRiskLimits(val quickfix.FIXInt) NoRiskLimitsField {
	return NoRiskLimitsField{val}
}

//NoRiskSecurityAltIDField is a NUMINGROUP field
type NoRiskSecurityAltIDField struct{ quickfix.FIXInt }

//Tag returns tag.NoRiskSecurityAltID (1540)
func (f NoRiskSecurityAltIDField) Tag() quickfix.Tag { return tag.NoRiskSecurityAltID }

//NewNoRiskSecurityAltID returns a new NoRiskSecurityAltIDField initialized with val
func NewNoRiskSecurityAltID(val quickfix.FIXInt) NoRiskSecurityAltIDField {
	return NoRiskSecurityAltIDField{val}
}

//NoRiskWarningLevelsField is a NUMINGROUP field
type NoRiskWarningLevelsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRiskWarningLevels (1559)
func (f NoRiskWarningLevelsField) Tag() quickfix.Tag { return tag.NoRiskWarningLevels }

//NewNoRiskWarningLevels returns a new NoRiskWarningLevelsField initialized with val
func NewNoRiskWarningLevels(val quickfix.FIXInt) NoRiskWarningLevelsField {
	return NoRiskWarningLevelsField{val}
}

//NoRootPartyIDsField is a NUMINGROUP field
type NoRootPartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRootPartyIDs (1116)
func (f NoRootPartyIDsField) Tag() quickfix.Tag { return tag.NoRootPartyIDs }

//NewNoRootPartyIDs returns a new NoRootPartyIDsField initialized with val
func NewNoRootPartyIDs(val quickfix.FIXInt) NoRootPartyIDsField {
	return NoRootPartyIDsField{val}
}

//NoRootPartySubIDsField is a NUMINGROUP field
type NoRootPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRootPartySubIDs (1120)
func (f NoRootPartySubIDsField) Tag() quickfix.Tag { return tag.NoRootPartySubIDs }

//NewNoRootPartySubIDs returns a new NoRootPartySubIDsField initialized with val
func NewNoRootPartySubIDs(val quickfix.FIXInt) NoRootPartySubIDsField {
	return NoRootPartySubIDsField{val}
}

//NoRoutingIDsField is a NUMINGROUP field
type NoRoutingIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRoutingIDs (215)
func (f NoRoutingIDsField) Tag() quickfix.Tag { return tag.NoRoutingIDs }

//NewNoRoutingIDs returns a new NoRoutingIDsField initialized with val
func NewNoRoutingIDs(val quickfix.FIXInt) NoRoutingIDsField {
	return NoRoutingIDsField{val}
}

//NoRptsField is a INT field
type NoRptsField struct{ quickfix.FIXInt }

//Tag returns tag.NoRpts (82)
func (f NoRptsField) Tag() quickfix.Tag { return tag.NoRpts }

//NewNoRpts returns a new NoRptsField initialized with val
func NewNoRpts(val quickfix.FIXInt) NoRptsField {
	return NoRptsField{val}
}

//NoSecurityAltIDField is a NUMINGROUP field
type NoSecurityAltIDField struct{ quickfix.FIXInt }

//Tag returns tag.NoSecurityAltID (454)
func (f NoSecurityAltIDField) Tag() quickfix.Tag { return tag.NoSecurityAltID }

//NewNoSecurityAltID returns a new NoSecurityAltIDField initialized with val
func NewNoSecurityAltID(val quickfix.FIXInt) NoSecurityAltIDField {
	return NoSecurityAltIDField{val}
}

//NoSecurityTypesField is a NUMINGROUP field
type NoSecurityTypesField struct{ quickfix.FIXInt }

//Tag returns tag.NoSecurityTypes (558)
func (f NoSecurityTypesField) Tag() quickfix.Tag { return tag.NoSecurityTypes }

//NewNoSecurityTypes returns a new NoSecurityTypesField initialized with val
func NewNoSecurityTypes(val quickfix.FIXInt) NoSecurityTypesField {
	return NoSecurityTypesField{val}
}

//NoSettlDetailsField is a NUMINGROUP field
type NoSettlDetailsField struct{ quickfix.FIXInt }

//Tag returns tag.NoSettlDetails (1158)
func (f NoSettlDetailsField) Tag() quickfix.Tag { return tag.NoSettlDetails }

//NewNoSettlDetails returns a new NoSettlDetailsField initialized with val
func NewNoSettlDetails(val quickfix.FIXInt) NoSettlDetailsField {
	return NoSettlDetailsField{val}
}

//NoSettlInstField is a NUMINGROUP field
type NoSettlInstField struct{ quickfix.FIXInt }

//Tag returns tag.NoSettlInst (778)
func (f NoSettlInstField) Tag() quickfix.Tag { return tag.NoSettlInst }

//NewNoSettlInst returns a new NoSettlInstField initialized with val
func NewNoSettlInst(val quickfix.FIXInt) NoSettlInstField {
	return NoSettlInstField{val}
}

//NoSettlObligField is a NUMINGROUP field
type NoSettlObligField struct{ quickfix.FIXInt }

//Tag returns tag.NoSettlOblig (1165)
func (f NoSettlObligField) Tag() quickfix.Tag { return tag.NoSettlOblig }

//NewNoSettlOblig returns a new NoSettlObligField initialized with val
func NewNoSettlOblig(val quickfix.FIXInt) NoSettlObligField {
	return NoSettlObligField{val}
}

//NoSettlPartyIDsField is a NUMINGROUP field
type NoSettlPartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoSettlPartyIDs (781)
func (f NoSettlPartyIDsField) Tag() quickfix.Tag { return tag.NoSettlPartyIDs }

//NewNoSettlPartyIDs returns a new NoSettlPartyIDsField initialized with val
func NewNoSettlPartyIDs(val quickfix.FIXInt) NoSettlPartyIDsField {
	return NoSettlPartyIDsField{val}
}

//NoSettlPartySubIDsField is a NUMINGROUP field
type NoSettlPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoSettlPartySubIDs (801)
func (f NoSettlPartySubIDsField) Tag() quickfix.Tag { return tag.NoSettlPartySubIDs }

//NewNoSettlPartySubIDs returns a new NoSettlPartySubIDsField initialized with val
func NewNoSettlPartySubIDs(val quickfix.FIXInt) NoSettlPartySubIDsField {
	return NoSettlPartySubIDsField{val}
}

//NoSideTrdRegTSField is a NUMINGROUP field
type NoSideTrdRegTSField struct{ quickfix.FIXInt }

//Tag returns tag.NoSideTrdRegTS (1016)
func (f NoSideTrdRegTSField) Tag() quickfix.Tag { return tag.NoSideTrdRegTS }

//NewNoSideTrdRegTS returns a new NoSideTrdRegTSField initialized with val
func NewNoSideTrdRegTS(val quickfix.FIXInt) NoSideTrdRegTSField {
	return NoSideTrdRegTSField{val}
}

//NoSidesField is a NUMINGROUP field
type NoSidesField struct{ quickfix.FIXInt }

//Tag returns tag.NoSides (552)
func (f NoSidesField) Tag() quickfix.Tag { return tag.NoSides }

//NewNoSides returns a new NoSidesField initialized with val
func NewNoSides(val quickfix.FIXInt) NoSidesField {
	return NoSidesField{val}
}

//NoStatsIndicatorsField is a NUMINGROUP field
type NoStatsIndicatorsField struct{ quickfix.FIXInt }

//Tag returns tag.NoStatsIndicators (1175)
func (f NoStatsIndicatorsField) Tag() quickfix.Tag { return tag.NoStatsIndicators }

//NewNoStatsIndicators returns a new NoStatsIndicatorsField initialized with val
func NewNoStatsIndicators(val quickfix.FIXInt) NoStatsIndicatorsField {
	return NoStatsIndicatorsField{val}
}

//NoStipulationsField is a NUMINGROUP field
type NoStipulationsField struct{ quickfix.FIXInt }

//Tag returns tag.NoStipulations (232)
func (f NoStipulationsField) Tag() quickfix.Tag { return tag.NoStipulations }

//NewNoStipulations returns a new NoStipulationsField initialized with val
func NewNoStipulations(val quickfix.FIXInt) NoStipulationsField {
	return NoStipulationsField{val}
}

//NoStrategyParametersField is a NUMINGROUP field
type NoStrategyParametersField struct{ quickfix.FIXInt }

//Tag returns tag.NoStrategyParameters (957)
func (f NoStrategyParametersField) Tag() quickfix.Tag { return tag.NoStrategyParameters }

//NewNoStrategyParameters returns a new NoStrategyParametersField initialized with val
func NewNoStrategyParameters(val quickfix.FIXInt) NoStrategyParametersField {
	return NoStrategyParametersField{val}
}

//NoStrikeRulesField is a NUMINGROUP field
type NoStrikeRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoStrikeRules (1201)
func (f NoStrikeRulesField) Tag() quickfix.Tag { return tag.NoStrikeRules }

//NewNoStrikeRules returns a new NoStrikeRulesField initialized with val
func NewNoStrikeRules(val quickfix.FIXInt) NoStrikeRulesField {
	return NoStrikeRulesField{val}
}

//NoStrikesField is a NUMINGROUP field
type NoStrikesField struct{ quickfix.FIXInt }

//Tag returns tag.NoStrikes (428)
func (f NoStrikesField) Tag() quickfix.Tag { return tag.NoStrikes }

//NewNoStrikes returns a new NoStrikesField initialized with val
func NewNoStrikes(val quickfix.FIXInt) NoStrikesField {
	return NoStrikesField{val}
}

//NoTargetPartyIDsField is a NUMINGROUP field
type NoTargetPartyIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoTargetPartyIDs (1461)
func (f NoTargetPartyIDsField) Tag() quickfix.Tag { return tag.NoTargetPartyIDs }

//NewNoTargetPartyIDs returns a new NoTargetPartyIDsField initialized with val
func NewNoTargetPartyIDs(val quickfix.FIXInt) NoTargetPartyIDsField {
	return NoTargetPartyIDsField{val}
}

//NoTickRulesField is a NUMINGROUP field
type NoTickRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoTickRules (1205)
func (f NoTickRulesField) Tag() quickfix.Tag { return tag.NoTickRules }

//NewNoTickRules returns a new NoTickRulesField initialized with val
func NewNoTickRules(val quickfix.FIXInt) NoTickRulesField {
	return NoTickRulesField{val}
}

//NoTimeInForceRulesField is a NUMINGROUP field
type NoTimeInForceRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoTimeInForceRules (1239)
func (f NoTimeInForceRulesField) Tag() quickfix.Tag { return tag.NoTimeInForceRules }

//NewNoTimeInForceRules returns a new NoTimeInForceRulesField initialized with val
func NewNoTimeInForceRules(val quickfix.FIXInt) NoTimeInForceRulesField {
	return NoTimeInForceRulesField{val}
}

//NoTradesField is a NUMINGROUP field
type NoTradesField struct{ quickfix.FIXInt }

//Tag returns tag.NoTrades (897)
func (f NoTradesField) Tag() quickfix.Tag { return tag.NoTrades }

//NewNoTrades returns a new NoTradesField initialized with val
func NewNoTrades(val quickfix.FIXInt) NoTradesField {
	return NoTradesField{val}
}

//NoTradingSessionRulesField is a NUMINGROUP field
type NoTradingSessionRulesField struct{ quickfix.FIXInt }

//Tag returns tag.NoTradingSessionRules (1309)
func (f NoTradingSessionRulesField) Tag() quickfix.Tag { return tag.NoTradingSessionRules }

//NewNoTradingSessionRules returns a new NoTradingSessionRulesField initialized with val
func NewNoTradingSessionRules(val quickfix.FIXInt) NoTradingSessionRulesField {
	return NoTradingSessionRulesField{val}
}

//NoTradingSessionsField is a NUMINGROUP field
type NoTradingSessionsField struct{ quickfix.FIXInt }

//Tag returns tag.NoTradingSessions (386)
func (f NoTradingSessionsField) Tag() quickfix.Tag { return tag.NoTradingSessions }

//NewNoTradingSessions returns a new NoTradingSessionsField initialized with val
func NewNoTradingSessions(val quickfix.FIXInt) NoTradingSessionsField {
	return NoTradingSessionsField{val}
}

//NoTrdRegTimestampsField is a NUMINGROUP field
type NoTrdRegTimestampsField struct{ quickfix.FIXInt }

//Tag returns tag.NoTrdRegTimestamps (768)
func (f NoTrdRegTimestampsField) Tag() quickfix.Tag { return tag.NoTrdRegTimestamps }

//NewNoTrdRegTimestamps returns a new NoTrdRegTimestampsField initialized with val
func NewNoTrdRegTimestamps(val quickfix.FIXInt) NoTrdRegTimestampsField {
	return NoTrdRegTimestampsField{val}
}

//NoTrdRepIndicatorsField is a NUMINGROUP field
type NoTrdRepIndicatorsField struct{ quickfix.FIXInt }

//Tag returns tag.NoTrdRepIndicators (1387)
func (f NoTrdRepIndicatorsField) Tag() quickfix.Tag { return tag.NoTrdRepIndicators }

//NewNoTrdRepIndicators returns a new NoTrdRepIndicatorsField initialized with val
func NewNoTrdRepIndicators(val quickfix.FIXInt) NoTrdRepIndicatorsField {
	return NoTrdRepIndicatorsField{val}
}

//NoUnderlyingAmountsField is a NUMINGROUP field
type NoUnderlyingAmountsField struct{ quickfix.FIXInt }

//Tag returns tag.NoUnderlyingAmounts (984)
func (f NoUnderlyingAmountsField) Tag() quickfix.Tag { return tag.NoUnderlyingAmounts }

//NewNoUnderlyingAmounts returns a new NoUnderlyingAmountsField initialized with val
func NewNoUnderlyingAmounts(val quickfix.FIXInt) NoUnderlyingAmountsField {
	return NoUnderlyingAmountsField{val}
}

//NoUnderlyingLegSecurityAltIDField is a NUMINGROUP field
type NoUnderlyingLegSecurityAltIDField struct{ quickfix.FIXInt }

//Tag returns tag.NoUnderlyingLegSecurityAltID (1334)
func (f NoUnderlyingLegSecurityAltIDField) Tag() quickfix.Tag { return tag.NoUnderlyingLegSecurityAltID }

//NewNoUnderlyingLegSecurityAltID returns a new NoUnderlyingLegSecurityAltIDField initialized with val
func NewNoUnderlyingLegSecurityAltID(val quickfix.FIXInt) NoUnderlyingLegSecurityAltIDField {
	return NoUnderlyingLegSecurityAltIDField{val}
}

//NoUnderlyingSecurityAltIDField is a NUMINGROUP field
type NoUnderlyingSecurityAltIDField struct{ quickfix.FIXInt }

//Tag returns tag.NoUnderlyingSecurityAltID (457)
func (f NoUnderlyingSecurityAltIDField) Tag() quickfix.Tag { return tag.NoUnderlyingSecurityAltID }

//NewNoUnderlyingSecurityAltID returns a new NoUnderlyingSecurityAltIDField initialized with val
func NewNoUnderlyingSecurityAltID(val quickfix.FIXInt) NoUnderlyingSecurityAltIDField {
	return NoUnderlyingSecurityAltIDField{val}
}

//NoUnderlyingStipsField is a NUMINGROUP field
type NoUnderlyingStipsField struct{ quickfix.FIXInt }

//Tag returns tag.NoUnderlyingStips (887)
func (f NoUnderlyingStipsField) Tag() quickfix.Tag { return tag.NoUnderlyingStips }

//NewNoUnderlyingStips returns a new NoUnderlyingStipsField initialized with val
func NewNoUnderlyingStips(val quickfix.FIXInt) NoUnderlyingStipsField {
	return NoUnderlyingStipsField{val}
}

//NoUnderlyingsField is a NUMINGROUP field
type NoUnderlyingsField struct{ quickfix.FIXInt }

//Tag returns tag.NoUnderlyings (711)
func (f NoUnderlyingsField) Tag() quickfix.Tag { return tag.NoUnderlyings }

//NewNoUnderlyings returns a new NoUnderlyingsField initialized with val
func NewNoUnderlyings(val quickfix.FIXInt) NoUnderlyingsField {
	return NoUnderlyingsField{val}
}

//NoUndlyInstrumentPartiesField is a NUMINGROUP field
type NoUndlyInstrumentPartiesField struct{ quickfix.FIXInt }

//Tag returns tag.NoUndlyInstrumentParties (1058)
func (f NoUndlyInstrumentPartiesField) Tag() quickfix.Tag { return tag.NoUndlyInstrumentParties }

//NewNoUndlyInstrumentParties returns a new NoUndlyInstrumentPartiesField initialized with val
func NewNoUndlyInstrumentParties(val quickfix.FIXInt) NoUndlyInstrumentPartiesField {
	return NoUndlyInstrumentPartiesField{val}
}

//NoUndlyInstrumentPartySubIDsField is a NUMINGROUP field
type NoUndlyInstrumentPartySubIDsField struct{ quickfix.FIXInt }

//Tag returns tag.NoUndlyInstrumentPartySubIDs (1062)
func (f NoUndlyInstrumentPartySubIDsField) Tag() quickfix.Tag { return tag.NoUndlyInstrumentPartySubIDs }

//NewNoUndlyInstrumentPartySubIDs returns a new NoUndlyInstrumentPartySubIDsField initialized with val
func NewNoUndlyInstrumentPartySubIDs(val quickfix.FIXInt) NoUndlyInstrumentPartySubIDsField {
	return NoUndlyInstrumentPartySubIDsField{val}
}

//NotAffOrigClOrdIDField is a STRING field
type NotAffOrigClOrdIDField struct{ quickfix.FIXString }

//Tag returns tag.NotAffOrigClOrdID (1372)
func (f NotAffOrigClOrdIDField) Tag() quickfix.Tag { return tag.NotAffOrigClOrdID }

//NewNotAffOrigClOrdID returns a new NotAffOrigClOrdIDField initialized with val
func NewNotAffOrigClOrdID(val quickfix.FIXString) NotAffOrigClOrdIDField {
	return NotAffOrigClOrdIDField{val}
}

//NotAffectedOrderIDField is a STRING field
type NotAffectedOrderIDField struct{ quickfix.FIXString }

//Tag returns tag.NotAffectedOrderID (1371)
func (f NotAffectedOrderIDField) Tag() quickfix.Tag { return tag.NotAffectedOrderID }

//NewNotAffectedOrderID returns a new NotAffectedOrderIDField initialized with val
func NewNotAffectedOrderID(val quickfix.FIXString) NotAffectedOrderIDField {
	return NotAffectedOrderIDField{val}
}

//NotifyBrokerOfCreditField is a BOOLEAN field
type NotifyBrokerOfCreditField struct{ quickfix.FIXBoolean }

//Tag returns tag.NotifyBrokerOfCredit (208)
func (f NotifyBrokerOfCreditField) Tag() quickfix.Tag { return tag.NotifyBrokerOfCredit }

//NewNotifyBrokerOfCredit returns a new NotifyBrokerOfCreditField initialized with val
func NewNotifyBrokerOfCredit(val quickfix.FIXBoolean) NotifyBrokerOfCreditField {
	return NotifyBrokerOfCreditField{val}
}

//NotionalPercentageOutstandingField is a PERCENTAGE field
type NotionalPercentageOutstandingField struct{ quickfix.FIXFloat }

//Tag returns tag.NotionalPercentageOutstanding (1451)
func (f NotionalPercentageOutstandingField) Tag() quickfix.Tag {
	return tag.NotionalPercentageOutstanding
}

//NewNotionalPercentageOutstanding returns a new NotionalPercentageOutstandingField initialized with val
func NewNotionalPercentageOutstanding(val quickfix.FIXFloat) NotionalPercentageOutstandingField {
	return NotionalPercentageOutstandingField{val}
}

//NumBiddersField is a INT field
type NumBiddersField struct{ quickfix.FIXInt }

//Tag returns tag.NumBidders (417)
func (f NumBiddersField) Tag() quickfix.Tag { return tag.NumBidders }

//NewNumBidders returns a new NumBiddersField initialized with val
func NewNumBidders(val quickfix.FIXInt) NumBiddersField {
	return NumBiddersField{val}
}

//NumDaysInterestField is a INT field
type NumDaysInterestField struct{ quickfix.FIXInt }

//Tag returns tag.NumDaysInterest (157)
func (f NumDaysInterestField) Tag() quickfix.Tag { return tag.NumDaysInterest }

//NewNumDaysInterest returns a new NumDaysInterestField initialized with val
func NewNumDaysInterest(val quickfix.FIXInt) NumDaysInterestField {
	return NumDaysInterestField{val}
}

//NumTicketsField is a INT field
type NumTicketsField struct{ quickfix.FIXInt }

//Tag returns tag.NumTickets (395)
func (f NumTicketsField) Tag() quickfix.Tag { return tag.NumTickets }

//NewNumTickets returns a new NumTicketsField initialized with val
func NewNumTickets(val quickfix.FIXInt) NumTicketsField {
	return NumTicketsField{val}
}

//NumberOfOrdersField is a INT field
type NumberOfOrdersField struct{ quickfix.FIXInt }

//Tag returns tag.NumberOfOrders (346)
func (f NumberOfOrdersField) Tag() quickfix.Tag { return tag.NumberOfOrders }

//NewNumberOfOrders returns a new NumberOfOrdersField initialized with val
func NewNumberOfOrders(val quickfix.FIXInt) NumberOfOrdersField {
	return NumberOfOrdersField{val}
}

//OddLotField is a BOOLEAN field
type OddLotField struct{ quickfix.FIXBoolean }

//Tag returns tag.OddLot (575)
func (f OddLotField) Tag() quickfix.Tag { return tag.OddLot }

//NewOddLot returns a new OddLotField initialized with val
func NewOddLot(val quickfix.FIXBoolean) OddLotField {
	return OddLotField{val}
}

//OfferForwardPointsField is a PRICEOFFSET field
type OfferForwardPointsField struct{ quickfix.FIXFloat }

//Tag returns tag.OfferForwardPoints (191)
func (f OfferForwardPointsField) Tag() quickfix.Tag { return tag.OfferForwardPoints }

//NewOfferForwardPoints returns a new OfferForwardPointsField initialized with val
func NewOfferForwardPoints(val quickfix.FIXFloat) OfferForwardPointsField {
	return OfferForwardPointsField{val}
}

//OfferForwardPoints2Field is a PRICEOFFSET field
type OfferForwardPoints2Field struct{ quickfix.FIXFloat }

//Tag returns tag.OfferForwardPoints2 (643)
func (f OfferForwardPoints2Field) Tag() quickfix.Tag { return tag.OfferForwardPoints2 }

//NewOfferForwardPoints2 returns a new OfferForwardPoints2Field initialized with val
func NewOfferForwardPoints2(val quickfix.FIXFloat) OfferForwardPoints2Field {
	return OfferForwardPoints2Field{val}
}

//OfferPxField is a PRICE field
type OfferPxField struct{ quickfix.FIXFloat }

//Tag returns tag.OfferPx (133)
func (f OfferPxField) Tag() quickfix.Tag { return tag.OfferPx }

//NewOfferPx returns a new OfferPxField initialized with val
func NewOfferPx(val quickfix.FIXFloat) OfferPxField {
	return OfferPxField{val}
}

//OfferSizeField is a QTY field
type OfferSizeField struct{ quickfix.FIXFloat }

//Tag returns tag.OfferSize (135)
func (f OfferSizeField) Tag() quickfix.Tag { return tag.OfferSize }

//NewOfferSize returns a new OfferSizeField initialized with val
func NewOfferSize(val quickfix.FIXFloat) OfferSizeField {
	return OfferSizeField{val}
}

//OfferSpotRateField is a PRICE field
type OfferSpotRateField struct{ quickfix.FIXFloat }

//Tag returns tag.OfferSpotRate (190)
func (f OfferSpotRateField) Tag() quickfix.Tag { return tag.OfferSpotRate }

//NewOfferSpotRate returns a new OfferSpotRateField initialized with val
func NewOfferSpotRate(val quickfix.FIXFloat) OfferSpotRateField {
	return OfferSpotRateField{val}
}

//OfferSwapPointsField is a PRICEOFFSET field
type OfferSwapPointsField struct{ quickfix.FIXFloat }

//Tag returns tag.OfferSwapPoints (1066)
func (f OfferSwapPointsField) Tag() quickfix.Tag { return tag.OfferSwapPoints }

//NewOfferSwapPoints returns a new OfferSwapPointsField initialized with val
func NewOfferSwapPoints(val quickfix.FIXFloat) OfferSwapPointsField {
	return OfferSwapPointsField{val}
}

//OfferYieldField is a PERCENTAGE field
type OfferYieldField struct{ quickfix.FIXFloat }

//Tag returns tag.OfferYield (634)
func (f OfferYieldField) Tag() quickfix.Tag { return tag.OfferYield }

//NewOfferYield returns a new OfferYieldField initialized with val
func NewOfferYield(val quickfix.FIXFloat) OfferYieldField {
	return OfferYieldField{val}
}

//OnBehalfOfCompIDField is a STRING field
type OnBehalfOfCompIDField struct{ quickfix.FIXString }

//Tag returns tag.OnBehalfOfCompID (115)
func (f OnBehalfOfCompIDField) Tag() quickfix.Tag { return tag.OnBehalfOfCompID }

//NewOnBehalfOfCompID returns a new OnBehalfOfCompIDField initialized with val
func NewOnBehalfOfCompID(val quickfix.FIXString) OnBehalfOfCompIDField {
	return OnBehalfOfCompIDField{val}
}

//OnBehalfOfLocationIDField is a STRING field
type OnBehalfOfLocationIDField struct{ quickfix.FIXString }

//Tag returns tag.OnBehalfOfLocationID (144)
func (f OnBehalfOfLocationIDField) Tag() quickfix.Tag { return tag.OnBehalfOfLocationID }

//NewOnBehalfOfLocationID returns a new OnBehalfOfLocationIDField initialized with val
func NewOnBehalfOfLocationID(val quickfix.FIXString) OnBehalfOfLocationIDField {
	return OnBehalfOfLocationIDField{val}
}

//OnBehalfOfSendingTimeField is a UTCTIMESTAMP field
type OnBehalfOfSendingTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.OnBehalfOfSendingTime (370)
func (f OnBehalfOfSendingTimeField) Tag() quickfix.Tag { return tag.OnBehalfOfSendingTime }

//NewOnBehalfOfSendingTime returns a new OnBehalfOfSendingTimeField initialized with val
func NewOnBehalfOfSendingTime(val time.Time) OnBehalfOfSendingTimeField {
	return OnBehalfOfSendingTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewOnBehalfOfSendingTimeNoMillis returns a new OnBehalfOfSendingTimeField initialized with val without millisecs
func NewOnBehalfOfSendingTimeNoMillis(val time.Time) OnBehalfOfSendingTimeField {
	return OnBehalfOfSendingTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//OnBehalfOfSubIDField is a STRING field
type OnBehalfOfSubIDField struct{ quickfix.FIXString }

//Tag returns tag.OnBehalfOfSubID (116)
func (f OnBehalfOfSubIDField) Tag() quickfix.Tag { return tag.OnBehalfOfSubID }

//NewOnBehalfOfSubID returns a new OnBehalfOfSubIDField initialized with val
func NewOnBehalfOfSubID(val quickfix.FIXString) OnBehalfOfSubIDField {
	return OnBehalfOfSubIDField{val}
}

//OpenCloseField is a CHAR field
type OpenCloseField struct{ quickfix.FIXString }

//Tag returns tag.OpenClose (77)
func (f OpenCloseField) Tag() quickfix.Tag { return tag.OpenClose }

//NewOpenClose returns a new OpenCloseField initialized with val
func NewOpenClose(val quickfix.FIXString) OpenCloseField {
	return OpenCloseField{val}
}

//OpenCloseSettlFlagField is a MULTIPLECHARVALUE field
type OpenCloseSettlFlagField struct{ quickfix.FIXString }

//Tag returns tag.OpenCloseSettlFlag (286)
func (f OpenCloseSettlFlagField) Tag() quickfix.Tag { return tag.OpenCloseSettlFlag }

//NewOpenCloseSettlFlag returns a new OpenCloseSettlFlagField initialized with val
func NewOpenCloseSettlFlag(val quickfix.FIXString) OpenCloseSettlFlagField {
	return OpenCloseSettlFlagField{val}
}

//OpenCloseSettleFlagField is a MULTIPLEVALUESTRING field
type OpenCloseSettleFlagField struct{ quickfix.FIXString }

//Tag returns tag.OpenCloseSettleFlag (286)
func (f OpenCloseSettleFlagField) Tag() quickfix.Tag { return tag.OpenCloseSettleFlag }

//NewOpenCloseSettleFlag returns a new OpenCloseSettleFlagField initialized with val
func NewOpenCloseSettleFlag(val quickfix.FIXString) OpenCloseSettleFlagField {
	return OpenCloseSettleFlagField{val}
}

//OpenInterestField is a AMT field
type OpenInterestField struct{ quickfix.FIXFloat }

//Tag returns tag.OpenInterest (746)
func (f OpenInterestField) Tag() quickfix.Tag { return tag.OpenInterest }

//NewOpenInterest returns a new OpenInterestField initialized with val
func NewOpenInterest(val quickfix.FIXFloat) OpenInterestField {
	return OpenInterestField{val}
}

//OptAttributeField is a CHAR field
type OptAttributeField struct{ quickfix.FIXString }

//Tag returns tag.OptAttribute (206)
func (f OptAttributeField) Tag() quickfix.Tag { return tag.OptAttribute }

//NewOptAttribute returns a new OptAttributeField initialized with val
func NewOptAttribute(val quickfix.FIXString) OptAttributeField {
	return OptAttributeField{val}
}

//OptPayAmountField is a AMT field
type OptPayAmountField struct{ quickfix.FIXFloat }

//Tag returns tag.OptPayAmount (1195)
func (f OptPayAmountField) Tag() quickfix.Tag { return tag.OptPayAmount }

//NewOptPayAmount returns a new OptPayAmountField initialized with val
func NewOptPayAmount(val quickfix.FIXFloat) OptPayAmountField {
	return OptPayAmountField{val}
}

//OptPayoutAmountField is a AMT field
type OptPayoutAmountField struct{ quickfix.FIXFloat }

//Tag returns tag.OptPayoutAmount (1195)
func (f OptPayoutAmountField) Tag() quickfix.Tag { return tag.OptPayoutAmount }

//NewOptPayoutAmount returns a new OptPayoutAmountField initialized with val
func NewOptPayoutAmount(val quickfix.FIXFloat) OptPayoutAmountField {
	return OptPayoutAmountField{val}
}

//OptPayoutTypeField is a INT field
type OptPayoutTypeField struct{ quickfix.FIXInt }

//Tag returns tag.OptPayoutType (1482)
func (f OptPayoutTypeField) Tag() quickfix.Tag { return tag.OptPayoutType }

//NewOptPayoutType returns a new OptPayoutTypeField initialized with val
func NewOptPayoutType(val quickfix.FIXInt) OptPayoutTypeField {
	return OptPayoutTypeField{val}
}

//OrdRejReasonField is a INT field
type OrdRejReasonField struct{ quickfix.FIXInt }

//Tag returns tag.OrdRejReason (103)
func (f OrdRejReasonField) Tag() quickfix.Tag { return tag.OrdRejReason }

//NewOrdRejReason returns a new OrdRejReasonField initialized with val
func NewOrdRejReason(val quickfix.FIXInt) OrdRejReasonField {
	return OrdRejReasonField{val}
}

//OrdStatusField is a CHAR field
type OrdStatusField struct{ quickfix.FIXString }

//Tag returns tag.OrdStatus (39)
func (f OrdStatusField) Tag() quickfix.Tag { return tag.OrdStatus }

//NewOrdStatus returns a new OrdStatusField initialized with val
func NewOrdStatus(val quickfix.FIXString) OrdStatusField {
	return OrdStatusField{val}
}

//OrdStatusReqIDField is a STRING field
type OrdStatusReqIDField struct{ quickfix.FIXString }

//Tag returns tag.OrdStatusReqID (790)
func (f OrdStatusReqIDField) Tag() quickfix.Tag { return tag.OrdStatusReqID }

//NewOrdStatusReqID returns a new OrdStatusReqIDField initialized with val
func NewOrdStatusReqID(val quickfix.FIXString) OrdStatusReqIDField {
	return OrdStatusReqIDField{val}
}

//OrdTypeField is a CHAR field
type OrdTypeField struct{ quickfix.FIXString }

//Tag returns tag.OrdType (40)
func (f OrdTypeField) Tag() quickfix.Tag { return tag.OrdType }

//NewOrdType returns a new OrdTypeField initialized with val
func NewOrdType(val quickfix.FIXString) OrdTypeField {
	return OrdTypeField{val}
}

//OrderAvgPxField is a PRICE field
type OrderAvgPxField struct{ quickfix.FIXFloat }

//Tag returns tag.OrderAvgPx (799)
func (f OrderAvgPxField) Tag() quickfix.Tag { return tag.OrderAvgPx }

//NewOrderAvgPx returns a new OrderAvgPxField initialized with val
func NewOrderAvgPx(val quickfix.FIXFloat) OrderAvgPxField {
	return OrderAvgPxField{val}
}

//OrderBookingQtyField is a QTY field
type OrderBookingQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.OrderBookingQty (800)
func (f OrderBookingQtyField) Tag() quickfix.Tag { return tag.OrderBookingQty }

//NewOrderBookingQty returns a new OrderBookingQtyField initialized with val
func NewOrderBookingQty(val quickfix.FIXFloat) OrderBookingQtyField {
	return OrderBookingQtyField{val}
}

//OrderCapacityField is a CHAR field
type OrderCapacityField struct{ quickfix.FIXString }

//Tag returns tag.OrderCapacity (528)
func (f OrderCapacityField) Tag() quickfix.Tag { return tag.OrderCapacity }

//NewOrderCapacity returns a new OrderCapacityField initialized with val
func NewOrderCapacity(val quickfix.FIXString) OrderCapacityField {
	return OrderCapacityField{val}
}

//OrderCapacityQtyField is a QTY field
type OrderCapacityQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.OrderCapacityQty (863)
func (f OrderCapacityQtyField) Tag() quickfix.Tag { return tag.OrderCapacityQty }

//NewOrderCapacityQty returns a new OrderCapacityQtyField initialized with val
func NewOrderCapacityQty(val quickfix.FIXFloat) OrderCapacityQtyField {
	return OrderCapacityQtyField{val}
}

//OrderCategoryField is a CHAR field
type OrderCategoryField struct{ quickfix.FIXString }

//Tag returns tag.OrderCategory (1115)
func (f OrderCategoryField) Tag() quickfix.Tag { return tag.OrderCategory }

//NewOrderCategory returns a new OrderCategoryField initialized with val
func NewOrderCategory(val quickfix.FIXString) OrderCategoryField {
	return OrderCategoryField{val}
}

//OrderDelayField is a INT field
type OrderDelayField struct{ quickfix.FIXInt }

//Tag returns tag.OrderDelay (1428)
func (f OrderDelayField) Tag() quickfix.Tag { return tag.OrderDelay }

//NewOrderDelay returns a new OrderDelayField initialized with val
func NewOrderDelay(val quickfix.FIXInt) OrderDelayField {
	return OrderDelayField{val}
}

//OrderDelayUnitField is a INT field
type OrderDelayUnitField struct{ quickfix.FIXInt }

//Tag returns tag.OrderDelayUnit (1429)
func (f OrderDelayUnitField) Tag() quickfix.Tag { return tag.OrderDelayUnit }

//NewOrderDelayUnit returns a new OrderDelayUnitField initialized with val
func NewOrderDelayUnit(val quickfix.FIXInt) OrderDelayUnitField {
	return OrderDelayUnitField{val}
}

//OrderHandlingInstSourceField is a INT field
type OrderHandlingInstSourceField struct{ quickfix.FIXInt }

//Tag returns tag.OrderHandlingInstSource (1032)
func (f OrderHandlingInstSourceField) Tag() quickfix.Tag { return tag.OrderHandlingInstSource }

//NewOrderHandlingInstSource returns a new OrderHandlingInstSourceField initialized with val
func NewOrderHandlingInstSource(val quickfix.FIXInt) OrderHandlingInstSourceField {
	return OrderHandlingInstSourceField{val}
}

//OrderIDField is a STRING field
type OrderIDField struct{ quickfix.FIXString }

//Tag returns tag.OrderID (37)
func (f OrderIDField) Tag() quickfix.Tag { return tag.OrderID }

//NewOrderID returns a new OrderIDField initialized with val
func NewOrderID(val quickfix.FIXString) OrderIDField {
	return OrderIDField{val}
}

//OrderInputDeviceField is a STRING field
type OrderInputDeviceField struct{ quickfix.FIXString }

//Tag returns tag.OrderInputDevice (821)
func (f OrderInputDeviceField) Tag() quickfix.Tag { return tag.OrderInputDevice }

//NewOrderInputDevice returns a new OrderInputDeviceField initialized with val
func NewOrderInputDevice(val quickfix.FIXString) OrderInputDeviceField {
	return OrderInputDeviceField{val}
}

//OrderPercentField is a PERCENTAGE field
type OrderPercentField struct{ quickfix.FIXFloat }

//Tag returns tag.OrderPercent (516)
func (f OrderPercentField) Tag() quickfix.Tag { return tag.OrderPercent }

//NewOrderPercent returns a new OrderPercentField initialized with val
func NewOrderPercent(val quickfix.FIXFloat) OrderPercentField {
	return OrderPercentField{val}
}

//OrderQtyField is a QTY field
type OrderQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.OrderQty (38)
func (f OrderQtyField) Tag() quickfix.Tag { return tag.OrderQty }

//NewOrderQty returns a new OrderQtyField initialized with val
func NewOrderQty(val quickfix.FIXFloat) OrderQtyField {
	return OrderQtyField{val}
}

//OrderQty2Field is a QTY field
type OrderQty2Field struct{ quickfix.FIXFloat }

//Tag returns tag.OrderQty2 (192)
func (f OrderQty2Field) Tag() quickfix.Tag { return tag.OrderQty2 }

//NewOrderQty2 returns a new OrderQty2Field initialized with val
func NewOrderQty2(val quickfix.FIXFloat) OrderQty2Field {
	return OrderQty2Field{val}
}

//OrderRestrictionsField is a MULTIPLECHARVALUE field
type OrderRestrictionsField struct{ quickfix.FIXString }

//Tag returns tag.OrderRestrictions (529)
func (f OrderRestrictionsField) Tag() quickfix.Tag { return tag.OrderRestrictions }

//NewOrderRestrictions returns a new OrderRestrictionsField initialized with val
func NewOrderRestrictions(val quickfix.FIXString) OrderRestrictionsField {
	return OrderRestrictionsField{val}
}

//OrigClOrdIDField is a STRING field
type OrigClOrdIDField struct{ quickfix.FIXString }

//Tag returns tag.OrigClOrdID (41)
func (f OrigClOrdIDField) Tag() quickfix.Tag { return tag.OrigClOrdID }

//NewOrigClOrdID returns a new OrigClOrdIDField initialized with val
func NewOrigClOrdID(val quickfix.FIXString) OrigClOrdIDField {
	return OrigClOrdIDField{val}
}

//OrigCrossIDField is a STRING field
type OrigCrossIDField struct{ quickfix.FIXString }

//Tag returns tag.OrigCrossID (551)
func (f OrigCrossIDField) Tag() quickfix.Tag { return tag.OrigCrossID }

//NewOrigCrossID returns a new OrigCrossIDField initialized with val
func NewOrigCrossID(val quickfix.FIXString) OrigCrossIDField {
	return OrigCrossIDField{val}
}

//OrigCustOrderCapacityField is a INT field
type OrigCustOrderCapacityField struct{ quickfix.FIXInt }

//Tag returns tag.OrigCustOrderCapacity (1432)
func (f OrigCustOrderCapacityField) Tag() quickfix.Tag { return tag.OrigCustOrderCapacity }

//NewOrigCustOrderCapacity returns a new OrigCustOrderCapacityField initialized with val
func NewOrigCustOrderCapacity(val quickfix.FIXInt) OrigCustOrderCapacityField {
	return OrigCustOrderCapacityField{val}
}

//OrigOrdModTimeField is a UTCTIMESTAMP field
type OrigOrdModTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.OrigOrdModTime (586)
func (f OrigOrdModTimeField) Tag() quickfix.Tag { return tag.OrigOrdModTime }

//NewOrigOrdModTime returns a new OrigOrdModTimeField initialized with val
func NewOrigOrdModTime(val time.Time) OrigOrdModTimeField {
	return OrigOrdModTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewOrigOrdModTimeNoMillis returns a new OrigOrdModTimeField initialized with val without millisecs
func NewOrigOrdModTimeNoMillis(val time.Time) OrigOrdModTimeField {
	return OrigOrdModTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//OrigPosReqRefIDField is a STRING field
type OrigPosReqRefIDField struct{ quickfix.FIXString }

//Tag returns tag.OrigPosReqRefID (713)
func (f OrigPosReqRefIDField) Tag() quickfix.Tag { return tag.OrigPosReqRefID }

//NewOrigPosReqRefID returns a new OrigPosReqRefIDField initialized with val
func NewOrigPosReqRefID(val quickfix.FIXString) OrigPosReqRefIDField {
	return OrigPosReqRefIDField{val}
}

//OrigSecondaryTradeIDField is a STRING field
type OrigSecondaryTradeIDField struct{ quickfix.FIXString }

//Tag returns tag.OrigSecondaryTradeID (1127)
func (f OrigSecondaryTradeIDField) Tag() quickfix.Tag { return tag.OrigSecondaryTradeID }

//NewOrigSecondaryTradeID returns a new OrigSecondaryTradeIDField initialized with val
func NewOrigSecondaryTradeID(val quickfix.FIXString) OrigSecondaryTradeIDField {
	return OrigSecondaryTradeIDField{val}
}

//OrigSendingTimeField is a UTCTIMESTAMP field
type OrigSendingTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.OrigSendingTime (122)
func (f OrigSendingTimeField) Tag() quickfix.Tag { return tag.OrigSendingTime }

//NewOrigSendingTime returns a new OrigSendingTimeField initialized with val
func NewOrigSendingTime(val time.Time) OrigSendingTimeField {
	return OrigSendingTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewOrigSendingTimeNoMillis returns a new OrigSendingTimeField initialized with val without millisecs
func NewOrigSendingTimeNoMillis(val time.Time) OrigSendingTimeField {
	return OrigSendingTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//OrigTimeField is a UTCTIMESTAMP field
type OrigTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.OrigTime (42)
func (f OrigTimeField) Tag() quickfix.Tag { return tag.OrigTime }

//NewOrigTime returns a new OrigTimeField initialized with val
func NewOrigTime(val time.Time) OrigTimeField {
	return OrigTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewOrigTimeNoMillis returns a new OrigTimeField initialized with val without millisecs
func NewOrigTimeNoMillis(val time.Time) OrigTimeField {
	return OrigTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//OrigTradeDateField is a LOCALMKTDATE field
type OrigTradeDateField struct{ quickfix.FIXString }

//Tag returns tag.OrigTradeDate (1125)
func (f OrigTradeDateField) Tag() quickfix.Tag { return tag.OrigTradeDate }

//NewOrigTradeDate returns a new OrigTradeDateField initialized with val
func NewOrigTradeDate(val quickfix.FIXString) OrigTradeDateField {
	return OrigTradeDateField{val}
}

//OrigTradeHandlingInstrField is a CHAR field
type OrigTradeHandlingInstrField struct{ quickfix.FIXString }

//Tag returns tag.OrigTradeHandlingInstr (1124)
func (f OrigTradeHandlingInstrField) Tag() quickfix.Tag { return tag.OrigTradeHandlingInstr }

//NewOrigTradeHandlingInstr returns a new OrigTradeHandlingInstrField initialized with val
func NewOrigTradeHandlingInstr(val quickfix.FIXString) OrigTradeHandlingInstrField {
	return OrigTradeHandlingInstrField{val}
}

//OrigTradeIDField is a STRING field
type OrigTradeIDField struct{ quickfix.FIXString }

//Tag returns tag.OrigTradeID (1126)
func (f OrigTradeIDField) Tag() quickfix.Tag { return tag.OrigTradeID }

//NewOrigTradeID returns a new OrigTradeIDField initialized with val
func NewOrigTradeID(val quickfix.FIXString) OrigTradeIDField {
	return OrigTradeIDField{val}
}

//OriginalNotionalPercentageOutstandingField is a PERCENTAGE field
type OriginalNotionalPercentageOutstandingField struct{ quickfix.FIXFloat }

//Tag returns tag.OriginalNotionalPercentageOutstanding (1452)
func (f OriginalNotionalPercentageOutstandingField) Tag() quickfix.Tag {
	return tag.OriginalNotionalPercentageOutstanding
}

//NewOriginalNotionalPercentageOutstanding returns a new OriginalNotionalPercentageOutstandingField initialized with val
func NewOriginalNotionalPercentageOutstanding(val quickfix.FIXFloat) OriginalNotionalPercentageOutstandingField {
	return OriginalNotionalPercentageOutstandingField{val}
}

//OutMainCntryUIndexField is a AMT field
type OutMainCntryUIndexField struct{ quickfix.FIXFloat }

//Tag returns tag.OutMainCntryUIndex (412)
func (f OutMainCntryUIndexField) Tag() quickfix.Tag { return tag.OutMainCntryUIndex }

//NewOutMainCntryUIndex returns a new OutMainCntryUIndexField initialized with val
func NewOutMainCntryUIndex(val quickfix.FIXFloat) OutMainCntryUIndexField {
	return OutMainCntryUIndexField{val}
}

//OutsideIndexPctField is a PERCENTAGE field
type OutsideIndexPctField struct{ quickfix.FIXFloat }

//Tag returns tag.OutsideIndexPct (407)
func (f OutsideIndexPctField) Tag() quickfix.Tag { return tag.OutsideIndexPct }

//NewOutsideIndexPct returns a new OutsideIndexPctField initialized with val
func NewOutsideIndexPct(val quickfix.FIXFloat) OutsideIndexPctField {
	return OutsideIndexPctField{val}
}

//OwnerTypeField is a INT field
type OwnerTypeField struct{ quickfix.FIXInt }

//Tag returns tag.OwnerType (522)
func (f OwnerTypeField) Tag() quickfix.Tag { return tag.OwnerType }

//NewOwnerType returns a new OwnerTypeField initialized with val
func NewOwnerType(val quickfix.FIXInt) OwnerTypeField {
	return OwnerTypeField{val}
}

//OwnershipTypeField is a CHAR field
type OwnershipTypeField struct{ quickfix.FIXString }

//Tag returns tag.OwnershipType (517)
func (f OwnershipTypeField) Tag() quickfix.Tag { return tag.OwnershipType }

//NewOwnershipType returns a new OwnershipTypeField initialized with val
func NewOwnershipType(val quickfix.FIXString) OwnershipTypeField {
	return OwnershipTypeField{val}
}

//ParentMktSegmIDField is a STRING field
type ParentMktSegmIDField struct{ quickfix.FIXString }

//Tag returns tag.ParentMktSegmID (1325)
func (f ParentMktSegmIDField) Tag() quickfix.Tag { return tag.ParentMktSegmID }

//NewParentMktSegmID returns a new ParentMktSegmIDField initialized with val
func NewParentMktSegmID(val quickfix.FIXString) ParentMktSegmIDField {
	return ParentMktSegmIDField{val}
}

//ParticipationRateField is a PERCENTAGE field
type ParticipationRateField struct{ quickfix.FIXFloat }

//Tag returns tag.ParticipationRate (849)
func (f ParticipationRateField) Tag() quickfix.Tag { return tag.ParticipationRate }

//NewParticipationRate returns a new ParticipationRateField initialized with val
func NewParticipationRate(val quickfix.FIXFloat) ParticipationRateField {
	return ParticipationRateField{val}
}

//PartyAltIDField is a STRING field
type PartyAltIDField struct{ quickfix.FIXString }

//Tag returns tag.PartyAltID (1517)
func (f PartyAltIDField) Tag() quickfix.Tag { return tag.PartyAltID }

//NewPartyAltID returns a new PartyAltIDField initialized with val
func NewPartyAltID(val quickfix.FIXString) PartyAltIDField {
	return PartyAltIDField{val}
}

//PartyAltIDSourceField is a CHAR field
type PartyAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.PartyAltIDSource (1518)
func (f PartyAltIDSourceField) Tag() quickfix.Tag { return tag.PartyAltIDSource }

//NewPartyAltIDSource returns a new PartyAltIDSourceField initialized with val
func NewPartyAltIDSource(val quickfix.FIXString) PartyAltIDSourceField {
	return PartyAltIDSourceField{val}
}

//PartyAltSubIDField is a STRING field
type PartyAltSubIDField struct{ quickfix.FIXString }

//Tag returns tag.PartyAltSubID (1520)
func (f PartyAltSubIDField) Tag() quickfix.Tag { return tag.PartyAltSubID }

//NewPartyAltSubID returns a new PartyAltSubIDField initialized with val
func NewPartyAltSubID(val quickfix.FIXString) PartyAltSubIDField {
	return PartyAltSubIDField{val}
}

//PartyAltSubIDTypeField is a INT field
type PartyAltSubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.PartyAltSubIDType (1521)
func (f PartyAltSubIDTypeField) Tag() quickfix.Tag { return tag.PartyAltSubIDType }

//NewPartyAltSubIDType returns a new PartyAltSubIDTypeField initialized with val
func NewPartyAltSubIDType(val quickfix.FIXInt) PartyAltSubIDTypeField {
	return PartyAltSubIDTypeField{val}
}

//PartyDetailsListReportIDField is a STRING field
type PartyDetailsListReportIDField struct{ quickfix.FIXString }

//Tag returns tag.PartyDetailsListReportID (1510)
func (f PartyDetailsListReportIDField) Tag() quickfix.Tag { return tag.PartyDetailsListReportID }

//NewPartyDetailsListReportID returns a new PartyDetailsListReportIDField initialized with val
func NewPartyDetailsListReportID(val quickfix.FIXString) PartyDetailsListReportIDField {
	return PartyDetailsListReportIDField{val}
}

//PartyDetailsListRequestIDField is a STRING field
type PartyDetailsListRequestIDField struct{ quickfix.FIXString }

//Tag returns tag.PartyDetailsListRequestID (1505)
func (f PartyDetailsListRequestIDField) Tag() quickfix.Tag { return tag.PartyDetailsListRequestID }

//NewPartyDetailsListRequestID returns a new PartyDetailsListRequestIDField initialized with val
func NewPartyDetailsListRequestID(val quickfix.FIXString) PartyDetailsListRequestIDField {
	return PartyDetailsListRequestIDField{val}
}

//PartyDetailsRequestResultField is a INT field
type PartyDetailsRequestResultField struct{ quickfix.FIXInt }

//Tag returns tag.PartyDetailsRequestResult (1511)
func (f PartyDetailsRequestResultField) Tag() quickfix.Tag { return tag.PartyDetailsRequestResult }

//NewPartyDetailsRequestResult returns a new PartyDetailsRequestResultField initialized with val
func NewPartyDetailsRequestResult(val quickfix.FIXInt) PartyDetailsRequestResultField {
	return PartyDetailsRequestResultField{val}
}

//PartyIDField is a STRING field
type PartyIDField struct{ quickfix.FIXString }

//Tag returns tag.PartyID (448)
func (f PartyIDField) Tag() quickfix.Tag { return tag.PartyID }

//NewPartyID returns a new PartyIDField initialized with val
func NewPartyID(val quickfix.FIXString) PartyIDField {
	return PartyIDField{val}
}

//PartyIDSourceField is a CHAR field
type PartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.PartyIDSource (447)
func (f PartyIDSourceField) Tag() quickfix.Tag { return tag.PartyIDSource }

//NewPartyIDSource returns a new PartyIDSourceField initialized with val
func NewPartyIDSource(val quickfix.FIXString) PartyIDSourceField {
	return PartyIDSourceField{val}
}

//PartyListResponseTypeField is a INT field
type PartyListResponseTypeField struct{ quickfix.FIXInt }

//Tag returns tag.PartyListResponseType (1507)
func (f PartyListResponseTypeField) Tag() quickfix.Tag { return tag.PartyListResponseType }

//NewPartyListResponseType returns a new PartyListResponseTypeField initialized with val
func NewPartyListResponseType(val quickfix.FIXInt) PartyListResponseTypeField {
	return PartyListResponseTypeField{val}
}

//PartyRelationshipField is a INT field
type PartyRelationshipField struct{ quickfix.FIXInt }

//Tag returns tag.PartyRelationship (1515)
func (f PartyRelationshipField) Tag() quickfix.Tag { return tag.PartyRelationship }

//NewPartyRelationship returns a new PartyRelationshipField initialized with val
func NewPartyRelationship(val quickfix.FIXInt) PartyRelationshipField {
	return PartyRelationshipField{val}
}

//PartyRoleField is a INT field
type PartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.PartyRole (452)
func (f PartyRoleField) Tag() quickfix.Tag { return tag.PartyRole }

//NewPartyRole returns a new PartyRoleField initialized with val
func NewPartyRole(val quickfix.FIXInt) PartyRoleField {
	return PartyRoleField{val}
}

//PartySubIDField is a STRING field
type PartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.PartySubID (523)
func (f PartySubIDField) Tag() quickfix.Tag { return tag.PartySubID }

//NewPartySubID returns a new PartySubIDField initialized with val
func NewPartySubID(val quickfix.FIXString) PartySubIDField {
	return PartySubIDField{val}
}

//PartySubIDTypeField is a INT field
type PartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.PartySubIDType (803)
func (f PartySubIDTypeField) Tag() quickfix.Tag { return tag.PartySubIDType }

//NewPartySubIDType returns a new PartySubIDTypeField initialized with val
func NewPartySubIDType(val quickfix.FIXInt) PartySubIDTypeField {
	return PartySubIDTypeField{val}
}

//PasswordField is a STRING field
type PasswordField struct{ quickfix.FIXString }

//Tag returns tag.Password (554)
func (f PasswordField) Tag() quickfix.Tag { return tag.Password }

//NewPassword returns a new PasswordField initialized with val
func NewPassword(val quickfix.FIXString) PasswordField {
	return PasswordField{val}
}

//PaymentDateField is a LOCALMKTDATE field
type PaymentDateField struct{ quickfix.FIXString }

//Tag returns tag.PaymentDate (504)
func (f PaymentDateField) Tag() quickfix.Tag { return tag.PaymentDate }

//NewPaymentDate returns a new PaymentDateField initialized with val
func NewPaymentDate(val quickfix.FIXString) PaymentDateField {
	return PaymentDateField{val}
}

//PaymentMethodField is a INT field
type PaymentMethodField struct{ quickfix.FIXInt }

//Tag returns tag.PaymentMethod (492)
func (f PaymentMethodField) Tag() quickfix.Tag { return tag.PaymentMethod }

//NewPaymentMethod returns a new PaymentMethodField initialized with val
func NewPaymentMethod(val quickfix.FIXInt) PaymentMethodField {
	return PaymentMethodField{val}
}

//PaymentRefField is a STRING field
type PaymentRefField struct{ quickfix.FIXString }

//Tag returns tag.PaymentRef (476)
func (f PaymentRefField) Tag() quickfix.Tag { return tag.PaymentRef }

//NewPaymentRef returns a new PaymentRefField initialized with val
func NewPaymentRef(val quickfix.FIXString) PaymentRefField {
	return PaymentRefField{val}
}

//PaymentRemitterIDField is a STRING field
type PaymentRemitterIDField struct{ quickfix.FIXString }

//Tag returns tag.PaymentRemitterID (505)
func (f PaymentRemitterIDField) Tag() quickfix.Tag { return tag.PaymentRemitterID }

//NewPaymentRemitterID returns a new PaymentRemitterIDField initialized with val
func NewPaymentRemitterID(val quickfix.FIXString) PaymentRemitterIDField {
	return PaymentRemitterIDField{val}
}

//PctAtRiskField is a PERCENTAGE field
type PctAtRiskField struct{ quickfix.FIXFloat }

//Tag returns tag.PctAtRisk (869)
func (f PctAtRiskField) Tag() quickfix.Tag { return tag.PctAtRisk }

//NewPctAtRisk returns a new PctAtRiskField initialized with val
func NewPctAtRisk(val quickfix.FIXFloat) PctAtRiskField {
	return PctAtRiskField{val}
}

//PegDifferenceField is a PRICEOFFSET field
type PegDifferenceField struct{ quickfix.FIXFloat }

//Tag returns tag.PegDifference (211)
func (f PegDifferenceField) Tag() quickfix.Tag { return tag.PegDifference }

//NewPegDifference returns a new PegDifferenceField initialized with val
func NewPegDifference(val quickfix.FIXFloat) PegDifferenceField {
	return PegDifferenceField{val}
}

//PegLimitTypeField is a INT field
type PegLimitTypeField struct{ quickfix.FIXInt }

//Tag returns tag.PegLimitType (837)
func (f PegLimitTypeField) Tag() quickfix.Tag { return tag.PegLimitType }

//NewPegLimitType returns a new PegLimitTypeField initialized with val
func NewPegLimitType(val quickfix.FIXInt) PegLimitTypeField {
	return PegLimitTypeField{val}
}

//PegMoveTypeField is a INT field
type PegMoveTypeField struct{ quickfix.FIXInt }

//Tag returns tag.PegMoveType (835)
func (f PegMoveTypeField) Tag() quickfix.Tag { return tag.PegMoveType }

//NewPegMoveType returns a new PegMoveTypeField initialized with val
func NewPegMoveType(val quickfix.FIXInt) PegMoveTypeField {
	return PegMoveTypeField{val}
}

//PegOffsetTypeField is a INT field
type PegOffsetTypeField struct{ quickfix.FIXInt }

//Tag returns tag.PegOffsetType (836)
func (f PegOffsetTypeField) Tag() quickfix.Tag { return tag.PegOffsetType }

//NewPegOffsetType returns a new PegOffsetTypeField initialized with val
func NewPegOffsetType(val quickfix.FIXInt) PegOffsetTypeField {
	return PegOffsetTypeField{val}
}

//PegOffsetValueField is a FLOAT field
type PegOffsetValueField struct{ quickfix.FIXFloat }

//Tag returns tag.PegOffsetValue (211)
func (f PegOffsetValueField) Tag() quickfix.Tag { return tag.PegOffsetValue }

//NewPegOffsetValue returns a new PegOffsetValueField initialized with val
func NewPegOffsetValue(val quickfix.FIXFloat) PegOffsetValueField {
	return PegOffsetValueField{val}
}

//PegPriceTypeField is a INT field
type PegPriceTypeField struct{ quickfix.FIXInt }

//Tag returns tag.PegPriceType (1094)
func (f PegPriceTypeField) Tag() quickfix.Tag { return tag.PegPriceType }

//NewPegPriceType returns a new PegPriceTypeField initialized with val
func NewPegPriceType(val quickfix.FIXInt) PegPriceTypeField {
	return PegPriceTypeField{val}
}

//PegRoundDirectionField is a INT field
type PegRoundDirectionField struct{ quickfix.FIXInt }

//Tag returns tag.PegRoundDirection (838)
func (f PegRoundDirectionField) Tag() quickfix.Tag { return tag.PegRoundDirection }

//NewPegRoundDirection returns a new PegRoundDirectionField initialized with val
func NewPegRoundDirection(val quickfix.FIXInt) PegRoundDirectionField {
	return PegRoundDirectionField{val}
}

//PegScopeField is a INT field
type PegScopeField struct{ quickfix.FIXInt }

//Tag returns tag.PegScope (840)
func (f PegScopeField) Tag() quickfix.Tag { return tag.PegScope }

//NewPegScope returns a new PegScopeField initialized with val
func NewPegScope(val quickfix.FIXInt) PegScopeField {
	return PegScopeField{val}
}

//PegSecurityDescField is a STRING field
type PegSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.PegSecurityDesc (1099)
func (f PegSecurityDescField) Tag() quickfix.Tag { return tag.PegSecurityDesc }

//NewPegSecurityDesc returns a new PegSecurityDescField initialized with val
func NewPegSecurityDesc(val quickfix.FIXString) PegSecurityDescField {
	return PegSecurityDescField{val}
}

//PegSecurityIDField is a STRING field
type PegSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.PegSecurityID (1097)
func (f PegSecurityIDField) Tag() quickfix.Tag { return tag.PegSecurityID }

//NewPegSecurityID returns a new PegSecurityIDField initialized with val
func NewPegSecurityID(val quickfix.FIXString) PegSecurityIDField {
	return PegSecurityIDField{val}
}

//PegSecurityIDSourceField is a STRING field
type PegSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.PegSecurityIDSource (1096)
func (f PegSecurityIDSourceField) Tag() quickfix.Tag { return tag.PegSecurityIDSource }

//NewPegSecurityIDSource returns a new PegSecurityIDSourceField initialized with val
func NewPegSecurityIDSource(val quickfix.FIXString) PegSecurityIDSourceField {
	return PegSecurityIDSourceField{val}
}

//PegSymbolField is a STRING field
type PegSymbolField struct{ quickfix.FIXString }

//Tag returns tag.PegSymbol (1098)
func (f PegSymbolField) Tag() quickfix.Tag { return tag.PegSymbol }

//NewPegSymbol returns a new PegSymbolField initialized with val
func NewPegSymbol(val quickfix.FIXString) PegSymbolField {
	return PegSymbolField{val}
}

//PeggedPriceField is a PRICE field
type PeggedPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.PeggedPrice (839)
func (f PeggedPriceField) Tag() quickfix.Tag { return tag.PeggedPrice }

//NewPeggedPrice returns a new PeggedPriceField initialized with val
func NewPeggedPrice(val quickfix.FIXFloat) PeggedPriceField {
	return PeggedPriceField{val}
}

//PeggedRefPriceField is a PRICE field
type PeggedRefPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.PeggedRefPrice (1095)
func (f PeggedRefPriceField) Tag() quickfix.Tag { return tag.PeggedRefPrice }

//NewPeggedRefPrice returns a new PeggedRefPriceField initialized with val
func NewPeggedRefPrice(val quickfix.FIXFloat) PeggedRefPriceField {
	return PeggedRefPriceField{val}
}

//PoolField is a STRING field
type PoolField struct{ quickfix.FIXString }

//Tag returns tag.Pool (691)
func (f PoolField) Tag() quickfix.Tag { return tag.Pool }

//NewPool returns a new PoolField initialized with val
func NewPool(val quickfix.FIXString) PoolField {
	return PoolField{val}
}

//PosAmtField is a AMT field
type PosAmtField struct{ quickfix.FIXFloat }

//Tag returns tag.PosAmt (708)
func (f PosAmtField) Tag() quickfix.Tag { return tag.PosAmt }

//NewPosAmt returns a new PosAmtField initialized with val
func NewPosAmt(val quickfix.FIXFloat) PosAmtField {
	return PosAmtField{val}
}

//PosAmtTypeField is a STRING field
type PosAmtTypeField struct{ quickfix.FIXString }

//Tag returns tag.PosAmtType (707)
func (f PosAmtTypeField) Tag() quickfix.Tag { return tag.PosAmtType }

//NewPosAmtType returns a new PosAmtTypeField initialized with val
func NewPosAmtType(val quickfix.FIXString) PosAmtTypeField {
	return PosAmtTypeField{val}
}

//PosMaintActionField is a INT field
type PosMaintActionField struct{ quickfix.FIXInt }

//Tag returns tag.PosMaintAction (712)
func (f PosMaintActionField) Tag() quickfix.Tag { return tag.PosMaintAction }

//NewPosMaintAction returns a new PosMaintActionField initialized with val
func NewPosMaintAction(val quickfix.FIXInt) PosMaintActionField {
	return PosMaintActionField{val}
}

//PosMaintResultField is a INT field
type PosMaintResultField struct{ quickfix.FIXInt }

//Tag returns tag.PosMaintResult (723)
func (f PosMaintResultField) Tag() quickfix.Tag { return tag.PosMaintResult }

//NewPosMaintResult returns a new PosMaintResultField initialized with val
func NewPosMaintResult(val quickfix.FIXInt) PosMaintResultField {
	return PosMaintResultField{val}
}

//PosMaintRptIDField is a STRING field
type PosMaintRptIDField struct{ quickfix.FIXString }

//Tag returns tag.PosMaintRptID (721)
func (f PosMaintRptIDField) Tag() quickfix.Tag { return tag.PosMaintRptID }

//NewPosMaintRptID returns a new PosMaintRptIDField initialized with val
func NewPosMaintRptID(val quickfix.FIXString) PosMaintRptIDField {
	return PosMaintRptIDField{val}
}

//PosMaintRptRefIDField is a STRING field
type PosMaintRptRefIDField struct{ quickfix.FIXString }

//Tag returns tag.PosMaintRptRefID (714)
func (f PosMaintRptRefIDField) Tag() quickfix.Tag { return tag.PosMaintRptRefID }

//NewPosMaintRptRefID returns a new PosMaintRptRefIDField initialized with val
func NewPosMaintRptRefID(val quickfix.FIXString) PosMaintRptRefIDField {
	return PosMaintRptRefIDField{val}
}

//PosMaintStatusField is a INT field
type PosMaintStatusField struct{ quickfix.FIXInt }

//Tag returns tag.PosMaintStatus (722)
func (f PosMaintStatusField) Tag() quickfix.Tag { return tag.PosMaintStatus }

//NewPosMaintStatus returns a new PosMaintStatusField initialized with val
func NewPosMaintStatus(val quickfix.FIXInt) PosMaintStatusField {
	return PosMaintStatusField{val}
}

//PosQtyStatusField is a INT field
type PosQtyStatusField struct{ quickfix.FIXInt }

//Tag returns tag.PosQtyStatus (706)
func (f PosQtyStatusField) Tag() quickfix.Tag { return tag.PosQtyStatus }

//NewPosQtyStatus returns a new PosQtyStatusField initialized with val
func NewPosQtyStatus(val quickfix.FIXInt) PosQtyStatusField {
	return PosQtyStatusField{val}
}

//PosReqIDField is a STRING field
type PosReqIDField struct{ quickfix.FIXString }

//Tag returns tag.PosReqID (710)
func (f PosReqIDField) Tag() quickfix.Tag { return tag.PosReqID }

//NewPosReqID returns a new PosReqIDField initialized with val
func NewPosReqID(val quickfix.FIXString) PosReqIDField {
	return PosReqIDField{val}
}

//PosReqResultField is a INT field
type PosReqResultField struct{ quickfix.FIXInt }

//Tag returns tag.PosReqResult (728)
func (f PosReqResultField) Tag() quickfix.Tag { return tag.PosReqResult }

//NewPosReqResult returns a new PosReqResultField initialized with val
func NewPosReqResult(val quickfix.FIXInt) PosReqResultField {
	return PosReqResultField{val}
}

//PosReqStatusField is a INT field
type PosReqStatusField struct{ quickfix.FIXInt }

//Tag returns tag.PosReqStatus (729)
func (f PosReqStatusField) Tag() quickfix.Tag { return tag.PosReqStatus }

//NewPosReqStatus returns a new PosReqStatusField initialized with val
func NewPosReqStatus(val quickfix.FIXInt) PosReqStatusField {
	return PosReqStatusField{val}
}

//PosReqTypeField is a INT field
type PosReqTypeField struct{ quickfix.FIXInt }

//Tag returns tag.PosReqType (724)
func (f PosReqTypeField) Tag() quickfix.Tag { return tag.PosReqType }

//NewPosReqType returns a new PosReqTypeField initialized with val
func NewPosReqType(val quickfix.FIXInt) PosReqTypeField {
	return PosReqTypeField{val}
}

//PosTransTypeField is a INT field
type PosTransTypeField struct{ quickfix.FIXInt }

//Tag returns tag.PosTransType (709)
func (f PosTransTypeField) Tag() quickfix.Tag { return tag.PosTransType }

//NewPosTransType returns a new PosTransTypeField initialized with val
func NewPosTransType(val quickfix.FIXInt) PosTransTypeField {
	return PosTransTypeField{val}
}

//PosTypeField is a STRING field
type PosTypeField struct{ quickfix.FIXString }

//Tag returns tag.PosType (703)
func (f PosTypeField) Tag() quickfix.Tag { return tag.PosType }

//NewPosType returns a new PosTypeField initialized with val
func NewPosType(val quickfix.FIXString) PosTypeField {
	return PosTypeField{val}
}

//PositionCurrencyField is a STRING field
type PositionCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.PositionCurrency (1055)
func (f PositionCurrencyField) Tag() quickfix.Tag { return tag.PositionCurrency }

//NewPositionCurrency returns a new PositionCurrencyField initialized with val
func NewPositionCurrency(val quickfix.FIXString) PositionCurrencyField {
	return PositionCurrencyField{val}
}

//PositionEffectField is a CHAR field
type PositionEffectField struct{ quickfix.FIXString }

//Tag returns tag.PositionEffect (77)
func (f PositionEffectField) Tag() quickfix.Tag { return tag.PositionEffect }

//NewPositionEffect returns a new PositionEffectField initialized with val
func NewPositionEffect(val quickfix.FIXString) PositionEffectField {
	return PositionEffectField{val}
}

//PositionLimitField is a INT field
type PositionLimitField struct{ quickfix.FIXInt }

//Tag returns tag.PositionLimit (970)
func (f PositionLimitField) Tag() quickfix.Tag { return tag.PositionLimit }

//NewPositionLimit returns a new PositionLimitField initialized with val
func NewPositionLimit(val quickfix.FIXInt) PositionLimitField {
	return PositionLimitField{val}
}

//PossDupFlagField is a BOOLEAN field
type PossDupFlagField struct{ quickfix.FIXBoolean }

//Tag returns tag.PossDupFlag (43)
func (f PossDupFlagField) Tag() quickfix.Tag { return tag.PossDupFlag }

//NewPossDupFlag returns a new PossDupFlagField initialized with val
func NewPossDupFlag(val quickfix.FIXBoolean) PossDupFlagField {
	return PossDupFlagField{val}
}

//PossResendField is a BOOLEAN field
type PossResendField struct{ quickfix.FIXBoolean }

//Tag returns tag.PossResend (97)
func (f PossResendField) Tag() quickfix.Tag { return tag.PossResend }

//NewPossResend returns a new PossResendField initialized with val
func NewPossResend(val quickfix.FIXBoolean) PossResendField {
	return PossResendField{val}
}

//PreTradeAnonymityField is a BOOLEAN field
type PreTradeAnonymityField struct{ quickfix.FIXBoolean }

//Tag returns tag.PreTradeAnonymity (1091)
func (f PreTradeAnonymityField) Tag() quickfix.Tag { return tag.PreTradeAnonymity }

//NewPreTradeAnonymity returns a new PreTradeAnonymityField initialized with val
func NewPreTradeAnonymity(val quickfix.FIXBoolean) PreTradeAnonymityField {
	return PreTradeAnonymityField{val}
}

//PreallocMethodField is a CHAR field
type PreallocMethodField struct{ quickfix.FIXString }

//Tag returns tag.PreallocMethod (591)
func (f PreallocMethodField) Tag() quickfix.Tag { return tag.PreallocMethod }

//NewPreallocMethod returns a new PreallocMethodField initialized with val
func NewPreallocMethod(val quickfix.FIXString) PreallocMethodField {
	return PreallocMethodField{val}
}

//PrevClosePxField is a PRICE field
type PrevClosePxField struct{ quickfix.FIXFloat }

//Tag returns tag.PrevClosePx (140)
func (f PrevClosePxField) Tag() quickfix.Tag { return tag.PrevClosePx }

//NewPrevClosePx returns a new PrevClosePxField initialized with val
func NewPrevClosePx(val quickfix.FIXFloat) PrevClosePxField {
	return PrevClosePxField{val}
}

//PreviouslyReportedField is a BOOLEAN field
type PreviouslyReportedField struct{ quickfix.FIXBoolean }

//Tag returns tag.PreviouslyReported (570)
func (f PreviouslyReportedField) Tag() quickfix.Tag { return tag.PreviouslyReported }

//NewPreviouslyReported returns a new PreviouslyReportedField initialized with val
func NewPreviouslyReported(val quickfix.FIXBoolean) PreviouslyReportedField {
	return PreviouslyReportedField{val}
}

//PriceField is a PRICE field
type PriceField struct{ quickfix.FIXFloat }

//Tag returns tag.Price (44)
func (f PriceField) Tag() quickfix.Tag { return tag.Price }

//NewPrice returns a new PriceField initialized with val
func NewPrice(val quickfix.FIXFloat) PriceField {
	return PriceField{val}
}

//Price2Field is a PRICE field
type Price2Field struct{ quickfix.FIXFloat }

//Tag returns tag.Price2 (640)
func (f Price2Field) Tag() quickfix.Tag { return tag.Price2 }

//NewPrice2 returns a new Price2Field initialized with val
func NewPrice2(val quickfix.FIXFloat) Price2Field {
	return Price2Field{val}
}

//PriceDeltaField is a FLOAT field
type PriceDeltaField struct{ quickfix.FIXFloat }

//Tag returns tag.PriceDelta (811)
func (f PriceDeltaField) Tag() quickfix.Tag { return tag.PriceDelta }

//NewPriceDelta returns a new PriceDeltaField initialized with val
func NewPriceDelta(val quickfix.FIXFloat) PriceDeltaField {
	return PriceDeltaField{val}
}

//PriceImprovementField is a PRICEOFFSET field
type PriceImprovementField struct{ quickfix.FIXFloat }

//Tag returns tag.PriceImprovement (639)
func (f PriceImprovementField) Tag() quickfix.Tag { return tag.PriceImprovement }

//NewPriceImprovement returns a new PriceImprovementField initialized with val
func NewPriceImprovement(val quickfix.FIXFloat) PriceImprovementField {
	return PriceImprovementField{val}
}

//PriceLimitTypeField is a INT field
type PriceLimitTypeField struct{ quickfix.FIXInt }

//Tag returns tag.PriceLimitType (1306)
func (f PriceLimitTypeField) Tag() quickfix.Tag { return tag.PriceLimitType }

//NewPriceLimitType returns a new PriceLimitTypeField initialized with val
func NewPriceLimitType(val quickfix.FIXInt) PriceLimitTypeField {
	return PriceLimitTypeField{val}
}

//PriceProtectionScopeField is a CHAR field
type PriceProtectionScopeField struct{ quickfix.FIXString }

//Tag returns tag.PriceProtectionScope (1092)
func (f PriceProtectionScopeField) Tag() quickfix.Tag { return tag.PriceProtectionScope }

//NewPriceProtectionScope returns a new PriceProtectionScopeField initialized with val
func NewPriceProtectionScope(val quickfix.FIXString) PriceProtectionScopeField {
	return PriceProtectionScopeField{val}
}

//PriceQuoteMethodField is a STRING field
type PriceQuoteMethodField struct{ quickfix.FIXString }

//Tag returns tag.PriceQuoteMethod (1196)
func (f PriceQuoteMethodField) Tag() quickfix.Tag { return tag.PriceQuoteMethod }

//NewPriceQuoteMethod returns a new PriceQuoteMethodField initialized with val
func NewPriceQuoteMethod(val quickfix.FIXString) PriceQuoteMethodField {
	return PriceQuoteMethodField{val}
}

//PriceTypeField is a INT field
type PriceTypeField struct{ quickfix.FIXInt }

//Tag returns tag.PriceType (423)
func (f PriceTypeField) Tag() quickfix.Tag { return tag.PriceType }

//NewPriceType returns a new PriceTypeField initialized with val
func NewPriceType(val quickfix.FIXInt) PriceTypeField {
	return PriceTypeField{val}
}

//PriceUnitOfMeasureField is a STRING field
type PriceUnitOfMeasureField struct{ quickfix.FIXString }

//Tag returns tag.PriceUnitOfMeasure (1191)
func (f PriceUnitOfMeasureField) Tag() quickfix.Tag { return tag.PriceUnitOfMeasure }

//NewPriceUnitOfMeasure returns a new PriceUnitOfMeasureField initialized with val
func NewPriceUnitOfMeasure(val quickfix.FIXString) PriceUnitOfMeasureField {
	return PriceUnitOfMeasureField{val}
}

//PriceUnitOfMeasureQtyField is a QTY field
type PriceUnitOfMeasureQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.PriceUnitOfMeasureQty (1192)
func (f PriceUnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.PriceUnitOfMeasureQty }

//NewPriceUnitOfMeasureQty returns a new PriceUnitOfMeasureQtyField initialized with val
func NewPriceUnitOfMeasureQty(val quickfix.FIXFloat) PriceUnitOfMeasureQtyField {
	return PriceUnitOfMeasureQtyField{val}
}

//PriorSettlPriceField is a PRICE field
type PriorSettlPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.PriorSettlPrice (734)
func (f PriorSettlPriceField) Tag() quickfix.Tag { return tag.PriorSettlPrice }

//NewPriorSettlPrice returns a new PriorSettlPriceField initialized with val
func NewPriorSettlPrice(val quickfix.FIXFloat) PriorSettlPriceField {
	return PriorSettlPriceField{val}
}

//PriorSpreadIndicatorField is a BOOLEAN field
type PriorSpreadIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.PriorSpreadIndicator (720)
func (f PriorSpreadIndicatorField) Tag() quickfix.Tag { return tag.PriorSpreadIndicator }

//NewPriorSpreadIndicator returns a new PriorSpreadIndicatorField initialized with val
func NewPriorSpreadIndicator(val quickfix.FIXBoolean) PriorSpreadIndicatorField {
	return PriorSpreadIndicatorField{val}
}

//PriorityIndicatorField is a INT field
type PriorityIndicatorField struct{ quickfix.FIXInt }

//Tag returns tag.PriorityIndicator (638)
func (f PriorityIndicatorField) Tag() quickfix.Tag { return tag.PriorityIndicator }

//NewPriorityIndicator returns a new PriorityIndicatorField initialized with val
func NewPriorityIndicator(val quickfix.FIXInt) PriorityIndicatorField {
	return PriorityIndicatorField{val}
}

//PrivateQuoteField is a BOOLEAN field
type PrivateQuoteField struct{ quickfix.FIXBoolean }

//Tag returns tag.PrivateQuote (1171)
func (f PrivateQuoteField) Tag() quickfix.Tag { return tag.PrivateQuote }

//NewPrivateQuote returns a new PrivateQuoteField initialized with val
func NewPrivateQuote(val quickfix.FIXBoolean) PrivateQuoteField {
	return PrivateQuoteField{val}
}

//ProcessCodeField is a CHAR field
type ProcessCodeField struct{ quickfix.FIXString }

//Tag returns tag.ProcessCode (81)
func (f ProcessCodeField) Tag() quickfix.Tag { return tag.ProcessCode }

//NewProcessCode returns a new ProcessCodeField initialized with val
func NewProcessCode(val quickfix.FIXString) ProcessCodeField {
	return ProcessCodeField{val}
}

//ProductField is a INT field
type ProductField struct{ quickfix.FIXInt }

//Tag returns tag.Product (460)
func (f ProductField) Tag() quickfix.Tag { return tag.Product }

//NewProduct returns a new ProductField initialized with val
func NewProduct(val quickfix.FIXInt) ProductField {
	return ProductField{val}
}

//ProductComplexField is a STRING field
type ProductComplexField struct{ quickfix.FIXString }

//Tag returns tag.ProductComplex (1227)
func (f ProductComplexField) Tag() quickfix.Tag { return tag.ProductComplex }

//NewProductComplex returns a new ProductComplexField initialized with val
func NewProductComplex(val quickfix.FIXString) ProductComplexField {
	return ProductComplexField{val}
}

//ProgPeriodIntervalField is a INT field
type ProgPeriodIntervalField struct{ quickfix.FIXInt }

//Tag returns tag.ProgPeriodInterval (415)
func (f ProgPeriodIntervalField) Tag() quickfix.Tag { return tag.ProgPeriodInterval }

//NewProgPeriodInterval returns a new ProgPeriodIntervalField initialized with val
func NewProgPeriodInterval(val quickfix.FIXInt) ProgPeriodIntervalField {
	return ProgPeriodIntervalField{val}
}

//ProgRptReqsField is a INT field
type ProgRptReqsField struct{ quickfix.FIXInt }

//Tag returns tag.ProgRptReqs (414)
func (f ProgRptReqsField) Tag() quickfix.Tag { return tag.ProgRptReqs }

//NewProgRptReqs returns a new ProgRptReqsField initialized with val
func NewProgRptReqs(val quickfix.FIXInt) ProgRptReqsField {
	return ProgRptReqsField{val}
}

//PublishTrdIndicatorField is a BOOLEAN field
type PublishTrdIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.PublishTrdIndicator (852)
func (f PublishTrdIndicatorField) Tag() quickfix.Tag { return tag.PublishTrdIndicator }

//NewPublishTrdIndicator returns a new PublishTrdIndicatorField initialized with val
func NewPublishTrdIndicator(val quickfix.FIXBoolean) PublishTrdIndicatorField {
	return PublishTrdIndicatorField{val}
}

//PutOrCallField is a INT field
type PutOrCallField struct{ quickfix.FIXInt }

//Tag returns tag.PutOrCall (201)
func (f PutOrCallField) Tag() quickfix.Tag { return tag.PutOrCall }

//NewPutOrCall returns a new PutOrCallField initialized with val
func NewPutOrCall(val quickfix.FIXInt) PutOrCallField {
	return PutOrCallField{val}
}

//QtyTypeField is a INT field
type QtyTypeField struct{ quickfix.FIXInt }

//Tag returns tag.QtyType (854)
func (f QtyTypeField) Tag() quickfix.Tag { return tag.QtyType }

//NewQtyType returns a new QtyTypeField initialized with val
func NewQtyType(val quickfix.FIXInt) QtyTypeField {
	return QtyTypeField{val}
}

//QuantityField is a QTY field
type QuantityField struct{ quickfix.FIXFloat }

//Tag returns tag.Quantity (53)
func (f QuantityField) Tag() quickfix.Tag { return tag.Quantity }

//NewQuantity returns a new QuantityField initialized with val
func NewQuantity(val quickfix.FIXFloat) QuantityField {
	return QuantityField{val}
}

//QuantityDateField is a LOCALMKTDATE field
type QuantityDateField struct{ quickfix.FIXString }

//Tag returns tag.QuantityDate (976)
func (f QuantityDateField) Tag() quickfix.Tag { return tag.QuantityDate }

//NewQuantityDate returns a new QuantityDateField initialized with val
func NewQuantityDate(val quickfix.FIXString) QuantityDateField {
	return QuantityDateField{val}
}

//QuantityTypeField is a INT field
type QuantityTypeField struct{ quickfix.FIXInt }

//Tag returns tag.QuantityType (465)
func (f QuantityTypeField) Tag() quickfix.Tag { return tag.QuantityType }

//NewQuantityType returns a new QuantityTypeField initialized with val
func NewQuantityType(val quickfix.FIXInt) QuantityTypeField {
	return QuantityTypeField{val}
}

//QuoteAckStatusField is a INT field
type QuoteAckStatusField struct{ quickfix.FIXInt }

//Tag returns tag.QuoteAckStatus (297)
func (f QuoteAckStatusField) Tag() quickfix.Tag { return tag.QuoteAckStatus }

//NewQuoteAckStatus returns a new QuoteAckStatusField initialized with val
func NewQuoteAckStatus(val quickfix.FIXInt) QuoteAckStatusField {
	return QuoteAckStatusField{val}
}

//QuoteCancelTypeField is a INT field
type QuoteCancelTypeField struct{ quickfix.FIXInt }

//Tag returns tag.QuoteCancelType (298)
func (f QuoteCancelTypeField) Tag() quickfix.Tag { return tag.QuoteCancelType }

//NewQuoteCancelType returns a new QuoteCancelTypeField initialized with val
func NewQuoteCancelType(val quickfix.FIXInt) QuoteCancelTypeField {
	return QuoteCancelTypeField{val}
}

//QuoteConditionField is a MULTIPLESTRINGVALUE field
type QuoteConditionField struct{ quickfix.FIXString }

//Tag returns tag.QuoteCondition (276)
func (f QuoteConditionField) Tag() quickfix.Tag { return tag.QuoteCondition }

//NewQuoteCondition returns a new QuoteConditionField initialized with val
func NewQuoteCondition(val quickfix.FIXString) QuoteConditionField {
	return QuoteConditionField{val}
}

//QuoteEntryIDField is a STRING field
type QuoteEntryIDField struct{ quickfix.FIXString }

//Tag returns tag.QuoteEntryID (299)
func (f QuoteEntryIDField) Tag() quickfix.Tag { return tag.QuoteEntryID }

//NewQuoteEntryID returns a new QuoteEntryIDField initialized with val
func NewQuoteEntryID(val quickfix.FIXString) QuoteEntryIDField {
	return QuoteEntryIDField{val}
}

//QuoteEntryRejectReasonField is a INT field
type QuoteEntryRejectReasonField struct{ quickfix.FIXInt }

//Tag returns tag.QuoteEntryRejectReason (368)
func (f QuoteEntryRejectReasonField) Tag() quickfix.Tag { return tag.QuoteEntryRejectReason }

//NewQuoteEntryRejectReason returns a new QuoteEntryRejectReasonField initialized with val
func NewQuoteEntryRejectReason(val quickfix.FIXInt) QuoteEntryRejectReasonField {
	return QuoteEntryRejectReasonField{val}
}

//QuoteEntryStatusField is a INT field
type QuoteEntryStatusField struct{ quickfix.FIXInt }

//Tag returns tag.QuoteEntryStatus (1167)
func (f QuoteEntryStatusField) Tag() quickfix.Tag { return tag.QuoteEntryStatus }

//NewQuoteEntryStatus returns a new QuoteEntryStatusField initialized with val
func NewQuoteEntryStatus(val quickfix.FIXInt) QuoteEntryStatusField {
	return QuoteEntryStatusField{val}
}

//QuoteIDField is a STRING field
type QuoteIDField struct{ quickfix.FIXString }

//Tag returns tag.QuoteID (117)
func (f QuoteIDField) Tag() quickfix.Tag { return tag.QuoteID }

//NewQuoteID returns a new QuoteIDField initialized with val
func NewQuoteID(val quickfix.FIXString) QuoteIDField {
	return QuoteIDField{val}
}

//QuoteMsgIDField is a STRING field
type QuoteMsgIDField struct{ quickfix.FIXString }

//Tag returns tag.QuoteMsgID (1166)
func (f QuoteMsgIDField) Tag() quickfix.Tag { return tag.QuoteMsgID }

//NewQuoteMsgID returns a new QuoteMsgIDField initialized with val
func NewQuoteMsgID(val quickfix.FIXString) QuoteMsgIDField {
	return QuoteMsgIDField{val}
}

//QuotePriceTypeField is a INT field
type QuotePriceTypeField struct{ quickfix.FIXInt }

//Tag returns tag.QuotePriceType (692)
func (f QuotePriceTypeField) Tag() quickfix.Tag { return tag.QuotePriceType }

//NewQuotePriceType returns a new QuotePriceTypeField initialized with val
func NewQuotePriceType(val quickfix.FIXInt) QuotePriceTypeField {
	return QuotePriceTypeField{val}
}

//QuoteQualifierField is a CHAR field
type QuoteQualifierField struct{ quickfix.FIXString }

//Tag returns tag.QuoteQualifier (695)
func (f QuoteQualifierField) Tag() quickfix.Tag { return tag.QuoteQualifier }

//NewQuoteQualifier returns a new QuoteQualifierField initialized with val
func NewQuoteQualifier(val quickfix.FIXString) QuoteQualifierField {
	return QuoteQualifierField{val}
}

//QuoteRejectReasonField is a INT field
type QuoteRejectReasonField struct{ quickfix.FIXInt }

//Tag returns tag.QuoteRejectReason (300)
func (f QuoteRejectReasonField) Tag() quickfix.Tag { return tag.QuoteRejectReason }

//NewQuoteRejectReason returns a new QuoteRejectReasonField initialized with val
func NewQuoteRejectReason(val quickfix.FIXInt) QuoteRejectReasonField {
	return QuoteRejectReasonField{val}
}

//QuoteReqIDField is a STRING field
type QuoteReqIDField struct{ quickfix.FIXString }

//Tag returns tag.QuoteReqID (131)
func (f QuoteReqIDField) Tag() quickfix.Tag { return tag.QuoteReqID }

//NewQuoteReqID returns a new QuoteReqIDField initialized with val
func NewQuoteReqID(val quickfix.FIXString) QuoteReqIDField {
	return QuoteReqIDField{val}
}

//QuoteRequestRejectReasonField is a INT field
type QuoteRequestRejectReasonField struct{ quickfix.FIXInt }

//Tag returns tag.QuoteRequestRejectReason (658)
func (f QuoteRequestRejectReasonField) Tag() quickfix.Tag { return tag.QuoteRequestRejectReason }

//NewQuoteRequestRejectReason returns a new QuoteRequestRejectReasonField initialized with val
func NewQuoteRequestRejectReason(val quickfix.FIXInt) QuoteRequestRejectReasonField {
	return QuoteRequestRejectReasonField{val}
}

//QuoteRequestTypeField is a INT field
type QuoteRequestTypeField struct{ quickfix.FIXInt }

//Tag returns tag.QuoteRequestType (303)
func (f QuoteRequestTypeField) Tag() quickfix.Tag { return tag.QuoteRequestType }

//NewQuoteRequestType returns a new QuoteRequestTypeField initialized with val
func NewQuoteRequestType(val quickfix.FIXInt) QuoteRequestTypeField {
	return QuoteRequestTypeField{val}
}

//QuoteRespIDField is a STRING field
type QuoteRespIDField struct{ quickfix.FIXString }

//Tag returns tag.QuoteRespID (693)
func (f QuoteRespIDField) Tag() quickfix.Tag { return tag.QuoteRespID }

//NewQuoteRespID returns a new QuoteRespIDField initialized with val
func NewQuoteRespID(val quickfix.FIXString) QuoteRespIDField {
	return QuoteRespIDField{val}
}

//QuoteRespTypeField is a INT field
type QuoteRespTypeField struct{ quickfix.FIXInt }

//Tag returns tag.QuoteRespType (694)
func (f QuoteRespTypeField) Tag() quickfix.Tag { return tag.QuoteRespType }

//NewQuoteRespType returns a new QuoteRespTypeField initialized with val
func NewQuoteRespType(val quickfix.FIXInt) QuoteRespTypeField {
	return QuoteRespTypeField{val}
}

//QuoteResponseLevelField is a INT field
type QuoteResponseLevelField struct{ quickfix.FIXInt }

//Tag returns tag.QuoteResponseLevel (301)
func (f QuoteResponseLevelField) Tag() quickfix.Tag { return tag.QuoteResponseLevel }

//NewQuoteResponseLevel returns a new QuoteResponseLevelField initialized with val
func NewQuoteResponseLevel(val quickfix.FIXInt) QuoteResponseLevelField {
	return QuoteResponseLevelField{val}
}

//QuoteSetIDField is a STRING field
type QuoteSetIDField struct{ quickfix.FIXString }

//Tag returns tag.QuoteSetID (302)
func (f QuoteSetIDField) Tag() quickfix.Tag { return tag.QuoteSetID }

//NewQuoteSetID returns a new QuoteSetIDField initialized with val
func NewQuoteSetID(val quickfix.FIXString) QuoteSetIDField {
	return QuoteSetIDField{val}
}

//QuoteSetValidUntilTimeField is a UTCTIMESTAMP field
type QuoteSetValidUntilTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.QuoteSetValidUntilTime (367)
func (f QuoteSetValidUntilTimeField) Tag() quickfix.Tag { return tag.QuoteSetValidUntilTime }

//NewQuoteSetValidUntilTime returns a new QuoteSetValidUntilTimeField initialized with val
func NewQuoteSetValidUntilTime(val time.Time) QuoteSetValidUntilTimeField {
	return QuoteSetValidUntilTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewQuoteSetValidUntilTimeNoMillis returns a new QuoteSetValidUntilTimeField initialized with val without millisecs
func NewQuoteSetValidUntilTimeNoMillis(val time.Time) QuoteSetValidUntilTimeField {
	return QuoteSetValidUntilTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//QuoteStatusField is a INT field
type QuoteStatusField struct{ quickfix.FIXInt }

//Tag returns tag.QuoteStatus (297)
func (f QuoteStatusField) Tag() quickfix.Tag { return tag.QuoteStatus }

//NewQuoteStatus returns a new QuoteStatusField initialized with val
func NewQuoteStatus(val quickfix.FIXInt) QuoteStatusField {
	return QuoteStatusField{val}
}

//QuoteStatusReqIDField is a STRING field
type QuoteStatusReqIDField struct{ quickfix.FIXString }

//Tag returns tag.QuoteStatusReqID (649)
func (f QuoteStatusReqIDField) Tag() quickfix.Tag { return tag.QuoteStatusReqID }

//NewQuoteStatusReqID returns a new QuoteStatusReqIDField initialized with val
func NewQuoteStatusReqID(val quickfix.FIXString) QuoteStatusReqIDField {
	return QuoteStatusReqIDField{val}
}

//QuoteTypeField is a INT field
type QuoteTypeField struct{ quickfix.FIXInt }

//Tag returns tag.QuoteType (537)
func (f QuoteTypeField) Tag() quickfix.Tag { return tag.QuoteType }

//NewQuoteType returns a new QuoteTypeField initialized with val
func NewQuoteType(val quickfix.FIXInt) QuoteTypeField {
	return QuoteTypeField{val}
}

//RFQReqIDField is a STRING field
type RFQReqIDField struct{ quickfix.FIXString }

//Tag returns tag.RFQReqID (644)
func (f RFQReqIDField) Tag() quickfix.Tag { return tag.RFQReqID }

//NewRFQReqID returns a new RFQReqIDField initialized with val
func NewRFQReqID(val quickfix.FIXString) RFQReqIDField {
	return RFQReqIDField{val}
}

//RateSourceField is a INT field
type RateSourceField struct{ quickfix.FIXInt }

//Tag returns tag.RateSource (1446)
func (f RateSourceField) Tag() quickfix.Tag { return tag.RateSource }

//NewRateSource returns a new RateSourceField initialized with val
func NewRateSource(val quickfix.FIXInt) RateSourceField {
	return RateSourceField{val}
}

//RateSourceTypeField is a INT field
type RateSourceTypeField struct{ quickfix.FIXInt }

//Tag returns tag.RateSourceType (1447)
func (f RateSourceTypeField) Tag() quickfix.Tag { return tag.RateSourceType }

//NewRateSourceType returns a new RateSourceTypeField initialized with val
func NewRateSourceType(val quickfix.FIXInt) RateSourceTypeField {
	return RateSourceTypeField{val}
}

//RatioQtyField is a QTY field
type RatioQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.RatioQty (319)
func (f RatioQtyField) Tag() quickfix.Tag { return tag.RatioQty }

//NewRatioQty returns a new RatioQtyField initialized with val
func NewRatioQty(val quickfix.FIXFloat) RatioQtyField {
	return RatioQtyField{val}
}

//RawDataField is a DATA field
type RawDataField struct{ quickfix.FIXString }

//Tag returns tag.RawData (96)
func (f RawDataField) Tag() quickfix.Tag { return tag.RawData }

//NewRawData returns a new RawDataField initialized with val
func NewRawData(val quickfix.FIXString) RawDataField {
	return RawDataField{val}
}

//RawDataLengthField is a LENGTH field
type RawDataLengthField struct{ quickfix.FIXInt }

//Tag returns tag.RawDataLength (95)
func (f RawDataLengthField) Tag() quickfix.Tag { return tag.RawDataLength }

//NewRawDataLength returns a new RawDataLengthField initialized with val
func NewRawDataLength(val quickfix.FIXInt) RawDataLengthField {
	return RawDataLengthField{val}
}

//ReceivedDeptIDField is a STRING field
type ReceivedDeptIDField struct{ quickfix.FIXString }

//Tag returns tag.ReceivedDeptID (1030)
func (f ReceivedDeptIDField) Tag() quickfix.Tag { return tag.ReceivedDeptID }

//NewReceivedDeptID returns a new ReceivedDeptIDField initialized with val
func NewReceivedDeptID(val quickfix.FIXString) ReceivedDeptIDField {
	return ReceivedDeptIDField{val}
}

//RedemptionDateField is a LOCALMKTDATE field
type RedemptionDateField struct{ quickfix.FIXString }

//Tag returns tag.RedemptionDate (240)
func (f RedemptionDateField) Tag() quickfix.Tag { return tag.RedemptionDate }

//NewRedemptionDate returns a new RedemptionDateField initialized with val
func NewRedemptionDate(val quickfix.FIXString) RedemptionDateField {
	return RedemptionDateField{val}
}

//RefAllocIDField is a STRING field
type RefAllocIDField struct{ quickfix.FIXString }

//Tag returns tag.RefAllocID (72)
func (f RefAllocIDField) Tag() quickfix.Tag { return tag.RefAllocID }

//NewRefAllocID returns a new RefAllocIDField initialized with val
func NewRefAllocID(val quickfix.FIXString) RefAllocIDField {
	return RefAllocIDField{val}
}

//RefApplExtIDField is a INT field
type RefApplExtIDField struct{ quickfix.FIXInt }

//Tag returns tag.RefApplExtID (1406)
func (f RefApplExtIDField) Tag() quickfix.Tag { return tag.RefApplExtID }

//NewRefApplExtID returns a new RefApplExtIDField initialized with val
func NewRefApplExtID(val quickfix.FIXInt) RefApplExtIDField {
	return RefApplExtIDField{val}
}

//RefApplIDField is a STRING field
type RefApplIDField struct{ quickfix.FIXString }

//Tag returns tag.RefApplID (1355)
func (f RefApplIDField) Tag() quickfix.Tag { return tag.RefApplID }

//NewRefApplID returns a new RefApplIDField initialized with val
func NewRefApplID(val quickfix.FIXString) RefApplIDField {
	return RefApplIDField{val}
}

//RefApplLastSeqNumField is a SEQNUM field
type RefApplLastSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.RefApplLastSeqNum (1357)
func (f RefApplLastSeqNumField) Tag() quickfix.Tag { return tag.RefApplLastSeqNum }

//NewRefApplLastSeqNum returns a new RefApplLastSeqNumField initialized with val
func NewRefApplLastSeqNum(val quickfix.FIXInt) RefApplLastSeqNumField {
	return RefApplLastSeqNumField{val}
}

//RefApplReqIDField is a STRING field
type RefApplReqIDField struct{ quickfix.FIXString }

//Tag returns tag.RefApplReqID (1433)
func (f RefApplReqIDField) Tag() quickfix.Tag { return tag.RefApplReqID }

//NewRefApplReqID returns a new RefApplReqIDField initialized with val
func NewRefApplReqID(val quickfix.FIXString) RefApplReqIDField {
	return RefApplReqIDField{val}
}

//RefApplVerIDField is a STRING field
type RefApplVerIDField struct{ quickfix.FIXString }

//Tag returns tag.RefApplVerID (1130)
func (f RefApplVerIDField) Tag() quickfix.Tag { return tag.RefApplVerID }

//NewRefApplVerID returns a new RefApplVerIDField initialized with val
func NewRefApplVerID(val quickfix.FIXString) RefApplVerIDField {
	return RefApplVerIDField{val}
}

//RefCompIDField is a STRING field
type RefCompIDField struct{ quickfix.FIXString }

//Tag returns tag.RefCompID (930)
func (f RefCompIDField) Tag() quickfix.Tag { return tag.RefCompID }

//NewRefCompID returns a new RefCompIDField initialized with val
func NewRefCompID(val quickfix.FIXString) RefCompIDField {
	return RefCompIDField{val}
}

//RefCstmApplVerIDField is a STRING field
type RefCstmApplVerIDField struct{ quickfix.FIXString }

//Tag returns tag.RefCstmApplVerID (1131)
func (f RefCstmApplVerIDField) Tag() quickfix.Tag { return tag.RefCstmApplVerID }

//NewRefCstmApplVerID returns a new RefCstmApplVerIDField initialized with val
func NewRefCstmApplVerID(val quickfix.FIXString) RefCstmApplVerIDField {
	return RefCstmApplVerIDField{val}
}

//RefMsgTypeField is a STRING field
type RefMsgTypeField struct{ quickfix.FIXString }

//Tag returns tag.RefMsgType (372)
func (f RefMsgTypeField) Tag() quickfix.Tag { return tag.RefMsgType }

//NewRefMsgType returns a new RefMsgTypeField initialized with val
func NewRefMsgType(val quickfix.FIXString) RefMsgTypeField {
	return RefMsgTypeField{val}
}

//RefOrdIDReasonField is a INT field
type RefOrdIDReasonField struct{ quickfix.FIXInt }

//Tag returns tag.RefOrdIDReason (1431)
func (f RefOrdIDReasonField) Tag() quickfix.Tag { return tag.RefOrdIDReason }

//NewRefOrdIDReason returns a new RefOrdIDReasonField initialized with val
func NewRefOrdIDReason(val quickfix.FIXInt) RefOrdIDReasonField {
	return RefOrdIDReasonField{val}
}

//RefOrderIDField is a STRING field
type RefOrderIDField struct{ quickfix.FIXString }

//Tag returns tag.RefOrderID (1080)
func (f RefOrderIDField) Tag() quickfix.Tag { return tag.RefOrderID }

//NewRefOrderID returns a new RefOrderIDField initialized with val
func NewRefOrderID(val quickfix.FIXString) RefOrderIDField {
	return RefOrderIDField{val}
}

//RefOrderIDSourceField is a CHAR field
type RefOrderIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RefOrderIDSource (1081)
func (f RefOrderIDSourceField) Tag() quickfix.Tag { return tag.RefOrderIDSource }

//NewRefOrderIDSource returns a new RefOrderIDSourceField initialized with val
func NewRefOrderIDSource(val quickfix.FIXString) RefOrderIDSourceField {
	return RefOrderIDSourceField{val}
}

//RefSeqNumField is a SEQNUM field
type RefSeqNumField struct{ quickfix.FIXInt }

//Tag returns tag.RefSeqNum (45)
func (f RefSeqNumField) Tag() quickfix.Tag { return tag.RefSeqNum }

//NewRefSeqNum returns a new RefSeqNumField initialized with val
func NewRefSeqNum(val quickfix.FIXInt) RefSeqNumField {
	return RefSeqNumField{val}
}

//RefSubIDField is a STRING field
type RefSubIDField struct{ quickfix.FIXString }

//Tag returns tag.RefSubID (931)
func (f RefSubIDField) Tag() quickfix.Tag { return tag.RefSubID }

//NewRefSubID returns a new RefSubIDField initialized with val
func NewRefSubID(val quickfix.FIXString) RefSubIDField {
	return RefSubIDField{val}
}

//RefTagIDField is a INT field
type RefTagIDField struct{ quickfix.FIXInt }

//Tag returns tag.RefTagID (371)
func (f RefTagIDField) Tag() quickfix.Tag { return tag.RefTagID }

//NewRefTagID returns a new RefTagIDField initialized with val
func NewRefTagID(val quickfix.FIXInt) RefTagIDField {
	return RefTagIDField{val}
}

//ReferencePageField is a STRING field
type ReferencePageField struct{ quickfix.FIXString }

//Tag returns tag.ReferencePage (1448)
func (f ReferencePageField) Tag() quickfix.Tag { return tag.ReferencePage }

//NewReferencePage returns a new ReferencePageField initialized with val
func NewReferencePage(val quickfix.FIXString) ReferencePageField {
	return ReferencePageField{val}
}

//RefreshIndicatorField is a BOOLEAN field
type RefreshIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.RefreshIndicator (1187)
func (f RefreshIndicatorField) Tag() quickfix.Tag { return tag.RefreshIndicator }

//NewRefreshIndicator returns a new RefreshIndicatorField initialized with val
func NewRefreshIndicator(val quickfix.FIXBoolean) RefreshIndicatorField {
	return RefreshIndicatorField{val}
}

//RefreshQtyField is a QTY field
type RefreshQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.RefreshQty (1088)
func (f RefreshQtyField) Tag() quickfix.Tag { return tag.RefreshQty }

//NewRefreshQty returns a new RefreshQtyField initialized with val
func NewRefreshQty(val quickfix.FIXFloat) RefreshQtyField {
	return RefreshQtyField{val}
}

//RegistAcctTypeField is a STRING field
type RegistAcctTypeField struct{ quickfix.FIXString }

//Tag returns tag.RegistAcctType (493)
func (f RegistAcctTypeField) Tag() quickfix.Tag { return tag.RegistAcctType }

//NewRegistAcctType returns a new RegistAcctTypeField initialized with val
func NewRegistAcctType(val quickfix.FIXString) RegistAcctTypeField {
	return RegistAcctTypeField{val}
}

//RegistDetlsField is a STRING field
type RegistDetlsField struct{ quickfix.FIXString }

//Tag returns tag.RegistDetls (509)
func (f RegistDetlsField) Tag() quickfix.Tag { return tag.RegistDetls }

//NewRegistDetls returns a new RegistDetlsField initialized with val
func NewRegistDetls(val quickfix.FIXString) RegistDetlsField {
	return RegistDetlsField{val}
}

//RegistDtlsField is a STRING field
type RegistDtlsField struct{ quickfix.FIXString }

//Tag returns tag.RegistDtls (509)
func (f RegistDtlsField) Tag() quickfix.Tag { return tag.RegistDtls }

//NewRegistDtls returns a new RegistDtlsField initialized with val
func NewRegistDtls(val quickfix.FIXString) RegistDtlsField {
	return RegistDtlsField{val}
}

//RegistEmailField is a STRING field
type RegistEmailField struct{ quickfix.FIXString }

//Tag returns tag.RegistEmail (511)
func (f RegistEmailField) Tag() quickfix.Tag { return tag.RegistEmail }

//NewRegistEmail returns a new RegistEmailField initialized with val
func NewRegistEmail(val quickfix.FIXString) RegistEmailField {
	return RegistEmailField{val}
}

//RegistIDField is a STRING field
type RegistIDField struct{ quickfix.FIXString }

//Tag returns tag.RegistID (513)
func (f RegistIDField) Tag() quickfix.Tag { return tag.RegistID }

//NewRegistID returns a new RegistIDField initialized with val
func NewRegistID(val quickfix.FIXString) RegistIDField {
	return RegistIDField{val}
}

//RegistRefIDField is a STRING field
type RegistRefIDField struct{ quickfix.FIXString }

//Tag returns tag.RegistRefID (508)
func (f RegistRefIDField) Tag() quickfix.Tag { return tag.RegistRefID }

//NewRegistRefID returns a new RegistRefIDField initialized with val
func NewRegistRefID(val quickfix.FIXString) RegistRefIDField {
	return RegistRefIDField{val}
}

//RegistRejReasonCodeField is a INT field
type RegistRejReasonCodeField struct{ quickfix.FIXInt }

//Tag returns tag.RegistRejReasonCode (507)
func (f RegistRejReasonCodeField) Tag() quickfix.Tag { return tag.RegistRejReasonCode }

//NewRegistRejReasonCode returns a new RegistRejReasonCodeField initialized with val
func NewRegistRejReasonCode(val quickfix.FIXInt) RegistRejReasonCodeField {
	return RegistRejReasonCodeField{val}
}

//RegistRejReasonTextField is a STRING field
type RegistRejReasonTextField struct{ quickfix.FIXString }

//Tag returns tag.RegistRejReasonText (496)
func (f RegistRejReasonTextField) Tag() quickfix.Tag { return tag.RegistRejReasonText }

//NewRegistRejReasonText returns a new RegistRejReasonTextField initialized with val
func NewRegistRejReasonText(val quickfix.FIXString) RegistRejReasonTextField {
	return RegistRejReasonTextField{val}
}

//RegistStatusField is a CHAR field
type RegistStatusField struct{ quickfix.FIXString }

//Tag returns tag.RegistStatus (506)
func (f RegistStatusField) Tag() quickfix.Tag { return tag.RegistStatus }

//NewRegistStatus returns a new RegistStatusField initialized with val
func NewRegistStatus(val quickfix.FIXString) RegistStatusField {
	return RegistStatusField{val}
}

//RegistTransTypeField is a CHAR field
type RegistTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.RegistTransType (514)
func (f RegistTransTypeField) Tag() quickfix.Tag { return tag.RegistTransType }

//NewRegistTransType returns a new RegistTransTypeField initialized with val
func NewRegistTransType(val quickfix.FIXString) RegistTransTypeField {
	return RegistTransTypeField{val}
}

//RejectTextField is a STRING field
type RejectTextField struct{ quickfix.FIXString }

//Tag returns tag.RejectText (1328)
func (f RejectTextField) Tag() quickfix.Tag { return tag.RejectText }

//NewRejectText returns a new RejectTextField initialized with val
func NewRejectText(val quickfix.FIXString) RejectTextField {
	return RejectTextField{val}
}

//RelSymTransactTimeField is a UTCTIMESTAMP field
type RelSymTransactTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.RelSymTransactTime (1504)
func (f RelSymTransactTimeField) Tag() quickfix.Tag { return tag.RelSymTransactTime }

//NewRelSymTransactTime returns a new RelSymTransactTimeField initialized with val
func NewRelSymTransactTime(val time.Time) RelSymTransactTimeField {
	return RelSymTransactTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewRelSymTransactTimeNoMillis returns a new RelSymTransactTimeField initialized with val without millisecs
func NewRelSymTransactTimeNoMillis(val time.Time) RelSymTransactTimeField {
	return RelSymTransactTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//RelatdSymField is a STRING field
type RelatdSymField struct{ quickfix.FIXString }

//Tag returns tag.RelatdSym (46)
func (f RelatdSymField) Tag() quickfix.Tag { return tag.RelatdSym }

//NewRelatdSym returns a new RelatdSymField initialized with val
func NewRelatdSym(val quickfix.FIXString) RelatdSymField {
	return RelatdSymField{val}
}

//RelatedContextPartyIDField is a STRING field
type RelatedContextPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.RelatedContextPartyID (1576)
func (f RelatedContextPartyIDField) Tag() quickfix.Tag { return tag.RelatedContextPartyID }

//NewRelatedContextPartyID returns a new RelatedContextPartyIDField initialized with val
func NewRelatedContextPartyID(val quickfix.FIXString) RelatedContextPartyIDField {
	return RelatedContextPartyIDField{val}
}

//RelatedContextPartyIDSourceField is a CHAR field
type RelatedContextPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RelatedContextPartyIDSource (1577)
func (f RelatedContextPartyIDSourceField) Tag() quickfix.Tag { return tag.RelatedContextPartyIDSource }

//NewRelatedContextPartyIDSource returns a new RelatedContextPartyIDSourceField initialized with val
func NewRelatedContextPartyIDSource(val quickfix.FIXString) RelatedContextPartyIDSourceField {
	return RelatedContextPartyIDSourceField{val}
}

//RelatedContextPartyRoleField is a INT field
type RelatedContextPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.RelatedContextPartyRole (1578)
func (f RelatedContextPartyRoleField) Tag() quickfix.Tag { return tag.RelatedContextPartyRole }

//NewRelatedContextPartyRole returns a new RelatedContextPartyRoleField initialized with val
func NewRelatedContextPartyRole(val quickfix.FIXInt) RelatedContextPartyRoleField {
	return RelatedContextPartyRoleField{val}
}

//RelatedContextPartySubIDField is a STRING field
type RelatedContextPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.RelatedContextPartySubID (1580)
func (f RelatedContextPartySubIDField) Tag() quickfix.Tag { return tag.RelatedContextPartySubID }

//NewRelatedContextPartySubID returns a new RelatedContextPartySubIDField initialized with val
func NewRelatedContextPartySubID(val quickfix.FIXString) RelatedContextPartySubIDField {
	return RelatedContextPartySubIDField{val}
}

//RelatedContextPartySubIDTypeField is a INT field
type RelatedContextPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.RelatedContextPartySubIDType (1581)
func (f RelatedContextPartySubIDTypeField) Tag() quickfix.Tag { return tag.RelatedContextPartySubIDType }

//NewRelatedContextPartySubIDType returns a new RelatedContextPartySubIDTypeField initialized with val
func NewRelatedContextPartySubIDType(val quickfix.FIXInt) RelatedContextPartySubIDTypeField {
	return RelatedContextPartySubIDTypeField{val}
}

//RelatedPartyAltIDField is a STRING field
type RelatedPartyAltIDField struct{ quickfix.FIXString }

//Tag returns tag.RelatedPartyAltID (1570)
func (f RelatedPartyAltIDField) Tag() quickfix.Tag { return tag.RelatedPartyAltID }

//NewRelatedPartyAltID returns a new RelatedPartyAltIDField initialized with val
func NewRelatedPartyAltID(val quickfix.FIXString) RelatedPartyAltIDField {
	return RelatedPartyAltIDField{val}
}

//RelatedPartyAltIDSourceField is a CHAR field
type RelatedPartyAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RelatedPartyAltIDSource (1571)
func (f RelatedPartyAltIDSourceField) Tag() quickfix.Tag { return tag.RelatedPartyAltIDSource }

//NewRelatedPartyAltIDSource returns a new RelatedPartyAltIDSourceField initialized with val
func NewRelatedPartyAltIDSource(val quickfix.FIXString) RelatedPartyAltIDSourceField {
	return RelatedPartyAltIDSourceField{val}
}

//RelatedPartyAltSubIDField is a STRING field
type RelatedPartyAltSubIDField struct{ quickfix.FIXString }

//Tag returns tag.RelatedPartyAltSubID (1573)
func (f RelatedPartyAltSubIDField) Tag() quickfix.Tag { return tag.RelatedPartyAltSubID }

//NewRelatedPartyAltSubID returns a new RelatedPartyAltSubIDField initialized with val
func NewRelatedPartyAltSubID(val quickfix.FIXString) RelatedPartyAltSubIDField {
	return RelatedPartyAltSubIDField{val}
}

//RelatedPartyAltSubIDTypeField is a INT field
type RelatedPartyAltSubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.RelatedPartyAltSubIDType (1574)
func (f RelatedPartyAltSubIDTypeField) Tag() quickfix.Tag { return tag.RelatedPartyAltSubIDType }

//NewRelatedPartyAltSubIDType returns a new RelatedPartyAltSubIDTypeField initialized with val
func NewRelatedPartyAltSubIDType(val quickfix.FIXInt) RelatedPartyAltSubIDTypeField {
	return RelatedPartyAltSubIDTypeField{val}
}

//RelatedPartyIDField is a STRING field
type RelatedPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.RelatedPartyID (1563)
func (f RelatedPartyIDField) Tag() quickfix.Tag { return tag.RelatedPartyID }

//NewRelatedPartyID returns a new RelatedPartyIDField initialized with val
func NewRelatedPartyID(val quickfix.FIXString) RelatedPartyIDField {
	return RelatedPartyIDField{val}
}

//RelatedPartyIDSourceField is a CHAR field
type RelatedPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RelatedPartyIDSource (1564)
func (f RelatedPartyIDSourceField) Tag() quickfix.Tag { return tag.RelatedPartyIDSource }

//NewRelatedPartyIDSource returns a new RelatedPartyIDSourceField initialized with val
func NewRelatedPartyIDSource(val quickfix.FIXString) RelatedPartyIDSourceField {
	return RelatedPartyIDSourceField{val}
}

//RelatedPartyRoleField is a INT field
type RelatedPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.RelatedPartyRole (1565)
func (f RelatedPartyRoleField) Tag() quickfix.Tag { return tag.RelatedPartyRole }

//NewRelatedPartyRole returns a new RelatedPartyRoleField initialized with val
func NewRelatedPartyRole(val quickfix.FIXInt) RelatedPartyRoleField {
	return RelatedPartyRoleField{val}
}

//RelatedPartySubIDField is a STRING field
type RelatedPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.RelatedPartySubID (1567)
func (f RelatedPartySubIDField) Tag() quickfix.Tag { return tag.RelatedPartySubID }

//NewRelatedPartySubID returns a new RelatedPartySubIDField initialized with val
func NewRelatedPartySubID(val quickfix.FIXString) RelatedPartySubIDField {
	return RelatedPartySubIDField{val}
}

//RelatedPartySubIDTypeField is a INT field
type RelatedPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.RelatedPartySubIDType (1568)
func (f RelatedPartySubIDTypeField) Tag() quickfix.Tag { return tag.RelatedPartySubIDType }

//NewRelatedPartySubIDType returns a new RelatedPartySubIDTypeField initialized with val
func NewRelatedPartySubIDType(val quickfix.FIXInt) RelatedPartySubIDTypeField {
	return RelatedPartySubIDTypeField{val}
}

//RelationshipRiskCFICodeField is a STRING field
type RelationshipRiskCFICodeField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskCFICode (1599)
func (f RelationshipRiskCFICodeField) Tag() quickfix.Tag { return tag.RelationshipRiskCFICode }

//NewRelationshipRiskCFICode returns a new RelationshipRiskCFICodeField initialized with val
func NewRelationshipRiskCFICode(val quickfix.FIXString) RelationshipRiskCFICodeField {
	return RelationshipRiskCFICodeField{val}
}

//RelationshipRiskCouponRateField is a PERCENTAGE field
type RelationshipRiskCouponRateField struct{ quickfix.FIXFloat }

//Tag returns tag.RelationshipRiskCouponRate (1608)
func (f RelationshipRiskCouponRateField) Tag() quickfix.Tag { return tag.RelationshipRiskCouponRate }

//NewRelationshipRiskCouponRate returns a new RelationshipRiskCouponRateField initialized with val
func NewRelationshipRiskCouponRate(val quickfix.FIXFloat) RelationshipRiskCouponRateField {
	return RelationshipRiskCouponRateField{val}
}

//RelationshipRiskEncodedSecurityDescField is a DATA field
type RelationshipRiskEncodedSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskEncodedSecurityDesc (1619)
func (f RelationshipRiskEncodedSecurityDescField) Tag() quickfix.Tag {
	return tag.RelationshipRiskEncodedSecurityDesc
}

//NewRelationshipRiskEncodedSecurityDesc returns a new RelationshipRiskEncodedSecurityDescField initialized with val
func NewRelationshipRiskEncodedSecurityDesc(val quickfix.FIXString) RelationshipRiskEncodedSecurityDescField {
	return RelationshipRiskEncodedSecurityDescField{val}
}

//RelationshipRiskEncodedSecurityDescLenField is a LENGTH field
type RelationshipRiskEncodedSecurityDescLenField struct{ quickfix.FIXInt }

//Tag returns tag.RelationshipRiskEncodedSecurityDescLen (1618)
func (f RelationshipRiskEncodedSecurityDescLenField) Tag() quickfix.Tag {
	return tag.RelationshipRiskEncodedSecurityDescLen
}

//NewRelationshipRiskEncodedSecurityDescLen returns a new RelationshipRiskEncodedSecurityDescLenField initialized with val
func NewRelationshipRiskEncodedSecurityDescLen(val quickfix.FIXInt) RelationshipRiskEncodedSecurityDescLenField {
	return RelationshipRiskEncodedSecurityDescLenField{val}
}

//RelationshipRiskFlexibleIndicatorField is a BOOLEAN field
type RelationshipRiskFlexibleIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.RelationshipRiskFlexibleIndicator (1607)
func (f RelationshipRiskFlexibleIndicatorField) Tag() quickfix.Tag {
	return tag.RelationshipRiskFlexibleIndicator
}

//NewRelationshipRiskFlexibleIndicator returns a new RelationshipRiskFlexibleIndicatorField initialized with val
func NewRelationshipRiskFlexibleIndicator(val quickfix.FIXBoolean) RelationshipRiskFlexibleIndicatorField {
	return RelationshipRiskFlexibleIndicatorField{val}
}

//RelationshipRiskInstrumentMultiplierField is a FLOAT field
type RelationshipRiskInstrumentMultiplierField struct{ quickfix.FIXFloat }

//Tag returns tag.RelationshipRiskInstrumentMultiplier (1612)
func (f RelationshipRiskInstrumentMultiplierField) Tag() quickfix.Tag {
	return tag.RelationshipRiskInstrumentMultiplier
}

//NewRelationshipRiskInstrumentMultiplier returns a new RelationshipRiskInstrumentMultiplierField initialized with val
func NewRelationshipRiskInstrumentMultiplier(val quickfix.FIXFloat) RelationshipRiskInstrumentMultiplierField {
	return RelationshipRiskInstrumentMultiplierField{val}
}

//RelationshipRiskInstrumentOperatorField is a INT field
type RelationshipRiskInstrumentOperatorField struct{ quickfix.FIXInt }

//Tag returns tag.RelationshipRiskInstrumentOperator (1588)
func (f RelationshipRiskInstrumentOperatorField) Tag() quickfix.Tag {
	return tag.RelationshipRiskInstrumentOperator
}

//NewRelationshipRiskInstrumentOperator returns a new RelationshipRiskInstrumentOperatorField initialized with val
func NewRelationshipRiskInstrumentOperator(val quickfix.FIXInt) RelationshipRiskInstrumentOperatorField {
	return RelationshipRiskInstrumentOperatorField{val}
}

//RelationshipRiskInstrumentSettlTypeField is a STRING field
type RelationshipRiskInstrumentSettlTypeField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskInstrumentSettlType (1611)
func (f RelationshipRiskInstrumentSettlTypeField) Tag() quickfix.Tag {
	return tag.RelationshipRiskInstrumentSettlType
}

//NewRelationshipRiskInstrumentSettlType returns a new RelationshipRiskInstrumentSettlTypeField initialized with val
func NewRelationshipRiskInstrumentSettlType(val quickfix.FIXString) RelationshipRiskInstrumentSettlTypeField {
	return RelationshipRiskInstrumentSettlTypeField{val}
}

//RelationshipRiskLimitAmountField is a AMT field
type RelationshipRiskLimitAmountField struct{ quickfix.FIXFloat }

//Tag returns tag.RelationshipRiskLimitAmount (1584)
func (f RelationshipRiskLimitAmountField) Tag() quickfix.Tag { return tag.RelationshipRiskLimitAmount }

//NewRelationshipRiskLimitAmount returns a new RelationshipRiskLimitAmountField initialized with val
func NewRelationshipRiskLimitAmount(val quickfix.FIXFloat) RelationshipRiskLimitAmountField {
	return RelationshipRiskLimitAmountField{val}
}

//RelationshipRiskLimitCurrencyField is a CURRENCY field
type RelationshipRiskLimitCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskLimitCurrency (1585)
func (f RelationshipRiskLimitCurrencyField) Tag() quickfix.Tag {
	return tag.RelationshipRiskLimitCurrency
}

//NewRelationshipRiskLimitCurrency returns a new RelationshipRiskLimitCurrencyField initialized with val
func NewRelationshipRiskLimitCurrency(val quickfix.FIXString) RelationshipRiskLimitCurrencyField {
	return RelationshipRiskLimitCurrencyField{val}
}

//RelationshipRiskLimitPlatformField is a STRING field
type RelationshipRiskLimitPlatformField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskLimitPlatform (1586)
func (f RelationshipRiskLimitPlatformField) Tag() quickfix.Tag {
	return tag.RelationshipRiskLimitPlatform
}

//NewRelationshipRiskLimitPlatform returns a new RelationshipRiskLimitPlatformField initialized with val
func NewRelationshipRiskLimitPlatform(val quickfix.FIXString) RelationshipRiskLimitPlatformField {
	return RelationshipRiskLimitPlatformField{val}
}

//RelationshipRiskLimitTypeField is a INT field
type RelationshipRiskLimitTypeField struct{ quickfix.FIXInt }

//Tag returns tag.RelationshipRiskLimitType (1583)
func (f RelationshipRiskLimitTypeField) Tag() quickfix.Tag { return tag.RelationshipRiskLimitType }

//NewRelationshipRiskLimitType returns a new RelationshipRiskLimitTypeField initialized with val
func NewRelationshipRiskLimitType(val quickfix.FIXInt) RelationshipRiskLimitTypeField {
	return RelationshipRiskLimitTypeField{val}
}

//RelationshipRiskMaturityMonthYearField is a MONTHYEAR field
type RelationshipRiskMaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskMaturityMonthYear (1602)
func (f RelationshipRiskMaturityMonthYearField) Tag() quickfix.Tag {
	return tag.RelationshipRiskMaturityMonthYear
}

//NewRelationshipRiskMaturityMonthYear returns a new RelationshipRiskMaturityMonthYearField initialized with val
func NewRelationshipRiskMaturityMonthYear(val quickfix.FIXString) RelationshipRiskMaturityMonthYearField {
	return RelationshipRiskMaturityMonthYearField{val}
}

//RelationshipRiskMaturityTimeField is a TZTIMEONLY field
type RelationshipRiskMaturityTimeField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskMaturityTime (1603)
func (f RelationshipRiskMaturityTimeField) Tag() quickfix.Tag { return tag.RelationshipRiskMaturityTime }

//NewRelationshipRiskMaturityTime returns a new RelationshipRiskMaturityTimeField initialized with val
func NewRelationshipRiskMaturityTime(val quickfix.FIXString) RelationshipRiskMaturityTimeField {
	return RelationshipRiskMaturityTimeField{val}
}

//RelationshipRiskProductField is a INT field
type RelationshipRiskProductField struct{ quickfix.FIXInt }

//Tag returns tag.RelationshipRiskProduct (1596)
func (f RelationshipRiskProductField) Tag() quickfix.Tag { return tag.RelationshipRiskProduct }

//NewRelationshipRiskProduct returns a new RelationshipRiskProductField initialized with val
func NewRelationshipRiskProduct(val quickfix.FIXInt) RelationshipRiskProductField {
	return RelationshipRiskProductField{val}
}

//RelationshipRiskProductComplexField is a STRING field
type RelationshipRiskProductComplexField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskProductComplex (1597)
func (f RelationshipRiskProductComplexField) Tag() quickfix.Tag {
	return tag.RelationshipRiskProductComplex
}

//NewRelationshipRiskProductComplex returns a new RelationshipRiskProductComplexField initialized with val
func NewRelationshipRiskProductComplex(val quickfix.FIXString) RelationshipRiskProductComplexField {
	return RelationshipRiskProductComplexField{val}
}

//RelationshipRiskPutOrCallField is a INT field
type RelationshipRiskPutOrCallField struct{ quickfix.FIXInt }

//Tag returns tag.RelationshipRiskPutOrCall (1606)
func (f RelationshipRiskPutOrCallField) Tag() quickfix.Tag { return tag.RelationshipRiskPutOrCall }

//NewRelationshipRiskPutOrCall returns a new RelationshipRiskPutOrCallField initialized with val
func NewRelationshipRiskPutOrCall(val quickfix.FIXInt) RelationshipRiskPutOrCallField {
	return RelationshipRiskPutOrCallField{val}
}

//RelationshipRiskRestructuringTypeField is a STRING field
type RelationshipRiskRestructuringTypeField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskRestructuringType (1604)
func (f RelationshipRiskRestructuringTypeField) Tag() quickfix.Tag {
	return tag.RelationshipRiskRestructuringType
}

//NewRelationshipRiskRestructuringType returns a new RelationshipRiskRestructuringTypeField initialized with val
func NewRelationshipRiskRestructuringType(val quickfix.FIXString) RelationshipRiskRestructuringTypeField {
	return RelationshipRiskRestructuringTypeField{val}
}

//RelationshipRiskSecurityAltIDField is a STRING field
type RelationshipRiskSecurityAltIDField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecurityAltID (1594)
func (f RelationshipRiskSecurityAltIDField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecurityAltID
}

//NewRelationshipRiskSecurityAltID returns a new RelationshipRiskSecurityAltIDField initialized with val
func NewRelationshipRiskSecurityAltID(val quickfix.FIXString) RelationshipRiskSecurityAltIDField {
	return RelationshipRiskSecurityAltIDField{val}
}

//RelationshipRiskSecurityAltIDSourceField is a STRING field
type RelationshipRiskSecurityAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecurityAltIDSource (1595)
func (f RelationshipRiskSecurityAltIDSourceField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecurityAltIDSource
}

//NewRelationshipRiskSecurityAltIDSource returns a new RelationshipRiskSecurityAltIDSourceField initialized with val
func NewRelationshipRiskSecurityAltIDSource(val quickfix.FIXString) RelationshipRiskSecurityAltIDSourceField {
	return RelationshipRiskSecurityAltIDSourceField{val}
}

//RelationshipRiskSecurityDescField is a STRING field
type RelationshipRiskSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecurityDesc (1610)
func (f RelationshipRiskSecurityDescField) Tag() quickfix.Tag { return tag.RelationshipRiskSecurityDesc }

//NewRelationshipRiskSecurityDesc returns a new RelationshipRiskSecurityDescField initialized with val
func NewRelationshipRiskSecurityDesc(val quickfix.FIXString) RelationshipRiskSecurityDescField {
	return RelationshipRiskSecurityDescField{val}
}

//RelationshipRiskSecurityExchangeField is a EXCHANGE field
type RelationshipRiskSecurityExchangeField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecurityExchange (1609)
func (f RelationshipRiskSecurityExchangeField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecurityExchange
}

//NewRelationshipRiskSecurityExchange returns a new RelationshipRiskSecurityExchangeField initialized with val
func NewRelationshipRiskSecurityExchange(val quickfix.FIXString) RelationshipRiskSecurityExchangeField {
	return RelationshipRiskSecurityExchangeField{val}
}

//RelationshipRiskSecurityGroupField is a STRING field
type RelationshipRiskSecurityGroupField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecurityGroup (1598)
func (f RelationshipRiskSecurityGroupField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecurityGroup
}

//NewRelationshipRiskSecurityGroup returns a new RelationshipRiskSecurityGroupField initialized with val
func NewRelationshipRiskSecurityGroup(val quickfix.FIXString) RelationshipRiskSecurityGroupField {
	return RelationshipRiskSecurityGroupField{val}
}

//RelationshipRiskSecurityIDField is a STRING field
type RelationshipRiskSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecurityID (1591)
func (f RelationshipRiskSecurityIDField) Tag() quickfix.Tag { return tag.RelationshipRiskSecurityID }

//NewRelationshipRiskSecurityID returns a new RelationshipRiskSecurityIDField initialized with val
func NewRelationshipRiskSecurityID(val quickfix.FIXString) RelationshipRiskSecurityIDField {
	return RelationshipRiskSecurityIDField{val}
}

//RelationshipRiskSecurityIDSourceField is a STRING field
type RelationshipRiskSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecurityIDSource (1592)
func (f RelationshipRiskSecurityIDSourceField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecurityIDSource
}

//NewRelationshipRiskSecurityIDSource returns a new RelationshipRiskSecurityIDSourceField initialized with val
func NewRelationshipRiskSecurityIDSource(val quickfix.FIXString) RelationshipRiskSecurityIDSourceField {
	return RelationshipRiskSecurityIDSourceField{val}
}

//RelationshipRiskSecuritySubTypeField is a STRING field
type RelationshipRiskSecuritySubTypeField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecuritySubType (1601)
func (f RelationshipRiskSecuritySubTypeField) Tag() quickfix.Tag {
	return tag.RelationshipRiskSecuritySubType
}

//NewRelationshipRiskSecuritySubType returns a new RelationshipRiskSecuritySubTypeField initialized with val
func NewRelationshipRiskSecuritySubType(val quickfix.FIXString) RelationshipRiskSecuritySubTypeField {
	return RelationshipRiskSecuritySubTypeField{val}
}

//RelationshipRiskSecurityTypeField is a STRING field
type RelationshipRiskSecurityTypeField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSecurityType (1600)
func (f RelationshipRiskSecurityTypeField) Tag() quickfix.Tag { return tag.RelationshipRiskSecurityType }

//NewRelationshipRiskSecurityType returns a new RelationshipRiskSecurityTypeField initialized with val
func NewRelationshipRiskSecurityType(val quickfix.FIXString) RelationshipRiskSecurityTypeField {
	return RelationshipRiskSecurityTypeField{val}
}

//RelationshipRiskSeniorityField is a STRING field
type RelationshipRiskSeniorityField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSeniority (1605)
func (f RelationshipRiskSeniorityField) Tag() quickfix.Tag { return tag.RelationshipRiskSeniority }

//NewRelationshipRiskSeniority returns a new RelationshipRiskSeniorityField initialized with val
func NewRelationshipRiskSeniority(val quickfix.FIXString) RelationshipRiskSeniorityField {
	return RelationshipRiskSeniorityField{val}
}

//RelationshipRiskSymbolField is a STRING field
type RelationshipRiskSymbolField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSymbol (1589)
func (f RelationshipRiskSymbolField) Tag() quickfix.Tag { return tag.RelationshipRiskSymbol }

//NewRelationshipRiskSymbol returns a new RelationshipRiskSymbolField initialized with val
func NewRelationshipRiskSymbol(val quickfix.FIXString) RelationshipRiskSymbolField {
	return RelationshipRiskSymbolField{val}
}

//RelationshipRiskSymbolSfxField is a STRING field
type RelationshipRiskSymbolSfxField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskSymbolSfx (1590)
func (f RelationshipRiskSymbolSfxField) Tag() quickfix.Tag { return tag.RelationshipRiskSymbolSfx }

//NewRelationshipRiskSymbolSfx returns a new RelationshipRiskSymbolSfxField initialized with val
func NewRelationshipRiskSymbolSfx(val quickfix.FIXString) RelationshipRiskSymbolSfxField {
	return RelationshipRiskSymbolSfxField{val}
}

//RelationshipRiskWarningLevelNameField is a STRING field
type RelationshipRiskWarningLevelNameField struct{ quickfix.FIXString }

//Tag returns tag.RelationshipRiskWarningLevelName (1615)
func (f RelationshipRiskWarningLevelNameField) Tag() quickfix.Tag {
	return tag.RelationshipRiskWarningLevelName
}

//NewRelationshipRiskWarningLevelName returns a new RelationshipRiskWarningLevelNameField initialized with val
func NewRelationshipRiskWarningLevelName(val quickfix.FIXString) RelationshipRiskWarningLevelNameField {
	return RelationshipRiskWarningLevelNameField{val}
}

//RelationshipRiskWarningLevelPercentField is a PERCENTAGE field
type RelationshipRiskWarningLevelPercentField struct{ quickfix.FIXFloat }

//Tag returns tag.RelationshipRiskWarningLevelPercent (1614)
func (f RelationshipRiskWarningLevelPercentField) Tag() quickfix.Tag {
	return tag.RelationshipRiskWarningLevelPercent
}

//NewRelationshipRiskWarningLevelPercent returns a new RelationshipRiskWarningLevelPercentField initialized with val
func NewRelationshipRiskWarningLevelPercent(val quickfix.FIXFloat) RelationshipRiskWarningLevelPercentField {
	return RelationshipRiskWarningLevelPercentField{val}
}

//RepoCollateralSecurityTypeField is a INT field
type RepoCollateralSecurityTypeField struct{ quickfix.FIXInt }

//Tag returns tag.RepoCollateralSecurityType (239)
func (f RepoCollateralSecurityTypeField) Tag() quickfix.Tag { return tag.RepoCollateralSecurityType }

//NewRepoCollateralSecurityType returns a new RepoCollateralSecurityTypeField initialized with val
func NewRepoCollateralSecurityType(val quickfix.FIXInt) RepoCollateralSecurityTypeField {
	return RepoCollateralSecurityTypeField{val}
}

//ReportToExchField is a BOOLEAN field
type ReportToExchField struct{ quickfix.FIXBoolean }

//Tag returns tag.ReportToExch (113)
func (f ReportToExchField) Tag() quickfix.Tag { return tag.ReportToExch }

//NewReportToExch returns a new ReportToExchField initialized with val
func NewReportToExch(val quickfix.FIXBoolean) ReportToExchField {
	return ReportToExchField{val}
}

//ReportedPxField is a PRICE field
type ReportedPxField struct{ quickfix.FIXFloat }

//Tag returns tag.ReportedPx (861)
func (f ReportedPxField) Tag() quickfix.Tag { return tag.ReportedPx }

//NewReportedPx returns a new ReportedPxField initialized with val
func NewReportedPx(val quickfix.FIXFloat) ReportedPxField {
	return ReportedPxField{val}
}

//ReportedPxDiffField is a BOOLEAN field
type ReportedPxDiffField struct{ quickfix.FIXBoolean }

//Tag returns tag.ReportedPxDiff (1134)
func (f ReportedPxDiffField) Tag() quickfix.Tag { return tag.ReportedPxDiff }

//NewReportedPxDiff returns a new ReportedPxDiffField initialized with val
func NewReportedPxDiff(val quickfix.FIXBoolean) ReportedPxDiffField {
	return ReportedPxDiffField{val}
}

//RepurchaseRateField is a PERCENTAGE field
type RepurchaseRateField struct{ quickfix.FIXFloat }

//Tag returns tag.RepurchaseRate (227)
func (f RepurchaseRateField) Tag() quickfix.Tag { return tag.RepurchaseRate }

//NewRepurchaseRate returns a new RepurchaseRateField initialized with val
func NewRepurchaseRate(val quickfix.FIXFloat) RepurchaseRateField {
	return RepurchaseRateField{val}
}

//RepurchaseTermField is a INT field
type RepurchaseTermField struct{ quickfix.FIXInt }

//Tag returns tag.RepurchaseTerm (226)
func (f RepurchaseTermField) Tag() quickfix.Tag { return tag.RepurchaseTerm }

//NewRepurchaseTerm returns a new RepurchaseTermField initialized with val
func NewRepurchaseTerm(val quickfix.FIXInt) RepurchaseTermField {
	return RepurchaseTermField{val}
}

//RequestedPartyRoleField is a INT field
type RequestedPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.RequestedPartyRole (1509)
func (f RequestedPartyRoleField) Tag() quickfix.Tag { return tag.RequestedPartyRole }

//NewRequestedPartyRole returns a new RequestedPartyRoleField initialized with val
func NewRequestedPartyRole(val quickfix.FIXInt) RequestedPartyRoleField {
	return RequestedPartyRoleField{val}
}

//ResetSeqNumFlagField is a BOOLEAN field
type ResetSeqNumFlagField struct{ quickfix.FIXBoolean }

//Tag returns tag.ResetSeqNumFlag (141)
func (f ResetSeqNumFlagField) Tag() quickfix.Tag { return tag.ResetSeqNumFlag }

//NewResetSeqNumFlag returns a new ResetSeqNumFlagField initialized with val
func NewResetSeqNumFlag(val quickfix.FIXBoolean) ResetSeqNumFlagField {
	return ResetSeqNumFlagField{val}
}

//RespondentTypeField is a INT field
type RespondentTypeField struct{ quickfix.FIXInt }

//Tag returns tag.RespondentType (1172)
func (f RespondentTypeField) Tag() quickfix.Tag { return tag.RespondentType }

//NewRespondentType returns a new RespondentTypeField initialized with val
func NewRespondentType(val quickfix.FIXInt) RespondentTypeField {
	return RespondentTypeField{val}
}

//ResponseDestinationField is a STRING field
type ResponseDestinationField struct{ quickfix.FIXString }

//Tag returns tag.ResponseDestination (726)
func (f ResponseDestinationField) Tag() quickfix.Tag { return tag.ResponseDestination }

//NewResponseDestination returns a new ResponseDestinationField initialized with val
func NewResponseDestination(val quickfix.FIXString) ResponseDestinationField {
	return ResponseDestinationField{val}
}

//ResponseTransportTypeField is a INT field
type ResponseTransportTypeField struct{ quickfix.FIXInt }

//Tag returns tag.ResponseTransportType (725)
func (f ResponseTransportTypeField) Tag() quickfix.Tag { return tag.ResponseTransportType }

//NewResponseTransportType returns a new ResponseTransportTypeField initialized with val
func NewResponseTransportType(val quickfix.FIXInt) ResponseTransportTypeField {
	return ResponseTransportTypeField{val}
}

//RestructuringTypeField is a STRING field
type RestructuringTypeField struct{ quickfix.FIXString }

//Tag returns tag.RestructuringType (1449)
func (f RestructuringTypeField) Tag() quickfix.Tag { return tag.RestructuringType }

//NewRestructuringType returns a new RestructuringTypeField initialized with val
func NewRestructuringType(val quickfix.FIXString) RestructuringTypeField {
	return RestructuringTypeField{val}
}

//ReversalIndicatorField is a BOOLEAN field
type ReversalIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.ReversalIndicator (700)
func (f ReversalIndicatorField) Tag() quickfix.Tag { return tag.ReversalIndicator }

//NewReversalIndicator returns a new ReversalIndicatorField initialized with val
func NewReversalIndicator(val quickfix.FIXBoolean) ReversalIndicatorField {
	return ReversalIndicatorField{val}
}

//RiskCFICodeField is a STRING field
type RiskCFICodeField struct{ quickfix.FIXString }

//Tag returns tag.RiskCFICode (1546)
func (f RiskCFICodeField) Tag() quickfix.Tag { return tag.RiskCFICode }

//NewRiskCFICode returns a new RiskCFICodeField initialized with val
func NewRiskCFICode(val quickfix.FIXString) RiskCFICodeField {
	return RiskCFICodeField{val}
}

//RiskCouponRateField is a PERCENTAGE field
type RiskCouponRateField struct{ quickfix.FIXFloat }

//Tag returns tag.RiskCouponRate (1555)
func (f RiskCouponRateField) Tag() quickfix.Tag { return tag.RiskCouponRate }

//NewRiskCouponRate returns a new RiskCouponRateField initialized with val
func NewRiskCouponRate(val quickfix.FIXFloat) RiskCouponRateField {
	return RiskCouponRateField{val}
}

//RiskEncodedSecurityDescField is a DATA field
type RiskEncodedSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.RiskEncodedSecurityDesc (1621)
func (f RiskEncodedSecurityDescField) Tag() quickfix.Tag { return tag.RiskEncodedSecurityDesc }

//NewRiskEncodedSecurityDesc returns a new RiskEncodedSecurityDescField initialized with val
func NewRiskEncodedSecurityDesc(val quickfix.FIXString) RiskEncodedSecurityDescField {
	return RiskEncodedSecurityDescField{val}
}

//RiskEncodedSecurityDescLenField is a LENGTH field
type RiskEncodedSecurityDescLenField struct{ quickfix.FIXInt }

//Tag returns tag.RiskEncodedSecurityDescLen (1620)
func (f RiskEncodedSecurityDescLenField) Tag() quickfix.Tag { return tag.RiskEncodedSecurityDescLen }

//NewRiskEncodedSecurityDescLen returns a new RiskEncodedSecurityDescLenField initialized with val
func NewRiskEncodedSecurityDescLen(val quickfix.FIXInt) RiskEncodedSecurityDescLenField {
	return RiskEncodedSecurityDescLenField{val}
}

//RiskFlexibleIndicatorField is a BOOLEAN field
type RiskFlexibleIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.RiskFlexibleIndicator (1554)
func (f RiskFlexibleIndicatorField) Tag() quickfix.Tag { return tag.RiskFlexibleIndicator }

//NewRiskFlexibleIndicator returns a new RiskFlexibleIndicatorField initialized with val
func NewRiskFlexibleIndicator(val quickfix.FIXBoolean) RiskFlexibleIndicatorField {
	return RiskFlexibleIndicatorField{val}
}

//RiskFreeRateField is a FLOAT field
type RiskFreeRateField struct{ quickfix.FIXFloat }

//Tag returns tag.RiskFreeRate (1190)
func (f RiskFreeRateField) Tag() quickfix.Tag { return tag.RiskFreeRate }

//NewRiskFreeRate returns a new RiskFreeRateField initialized with val
func NewRiskFreeRate(val quickfix.FIXFloat) RiskFreeRateField {
	return RiskFreeRateField{val}
}

//RiskInstrumentMultiplierField is a FLOAT field
type RiskInstrumentMultiplierField struct{ quickfix.FIXFloat }

//Tag returns tag.RiskInstrumentMultiplier (1558)
func (f RiskInstrumentMultiplierField) Tag() quickfix.Tag { return tag.RiskInstrumentMultiplier }

//NewRiskInstrumentMultiplier returns a new RiskInstrumentMultiplierField initialized with val
func NewRiskInstrumentMultiplier(val quickfix.FIXFloat) RiskInstrumentMultiplierField {
	return RiskInstrumentMultiplierField{val}
}

//RiskInstrumentOperatorField is a INT field
type RiskInstrumentOperatorField struct{ quickfix.FIXInt }

//Tag returns tag.RiskInstrumentOperator (1535)
func (f RiskInstrumentOperatorField) Tag() quickfix.Tag { return tag.RiskInstrumentOperator }

//NewRiskInstrumentOperator returns a new RiskInstrumentOperatorField initialized with val
func NewRiskInstrumentOperator(val quickfix.FIXInt) RiskInstrumentOperatorField {
	return RiskInstrumentOperatorField{val}
}

//RiskInstrumentSettlTypeField is a STRING field
type RiskInstrumentSettlTypeField struct{ quickfix.FIXString }

//Tag returns tag.RiskInstrumentSettlType (1557)
func (f RiskInstrumentSettlTypeField) Tag() quickfix.Tag { return tag.RiskInstrumentSettlType }

//NewRiskInstrumentSettlType returns a new RiskInstrumentSettlTypeField initialized with val
func NewRiskInstrumentSettlType(val quickfix.FIXString) RiskInstrumentSettlTypeField {
	return RiskInstrumentSettlTypeField{val}
}

//RiskLimitAmountField is a AMT field
type RiskLimitAmountField struct{ quickfix.FIXFloat }

//Tag returns tag.RiskLimitAmount (1531)
func (f RiskLimitAmountField) Tag() quickfix.Tag { return tag.RiskLimitAmount }

//NewRiskLimitAmount returns a new RiskLimitAmountField initialized with val
func NewRiskLimitAmount(val quickfix.FIXFloat) RiskLimitAmountField {
	return RiskLimitAmountField{val}
}

//RiskLimitCurrencyField is a CURRENCY field
type RiskLimitCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.RiskLimitCurrency (1532)
func (f RiskLimitCurrencyField) Tag() quickfix.Tag { return tag.RiskLimitCurrency }

//NewRiskLimitCurrency returns a new RiskLimitCurrencyField initialized with val
func NewRiskLimitCurrency(val quickfix.FIXString) RiskLimitCurrencyField {
	return RiskLimitCurrencyField{val}
}

//RiskLimitPlatformField is a STRING field
type RiskLimitPlatformField struct{ quickfix.FIXString }

//Tag returns tag.RiskLimitPlatform (1533)
func (f RiskLimitPlatformField) Tag() quickfix.Tag { return tag.RiskLimitPlatform }

//NewRiskLimitPlatform returns a new RiskLimitPlatformField initialized with val
func NewRiskLimitPlatform(val quickfix.FIXString) RiskLimitPlatformField {
	return RiskLimitPlatformField{val}
}

//RiskLimitTypeField is a INT field
type RiskLimitTypeField struct{ quickfix.FIXInt }

//Tag returns tag.RiskLimitType (1530)
func (f RiskLimitTypeField) Tag() quickfix.Tag { return tag.RiskLimitType }

//NewRiskLimitType returns a new RiskLimitTypeField initialized with val
func NewRiskLimitType(val quickfix.FIXInt) RiskLimitTypeField {
	return RiskLimitTypeField{val}
}

//RiskMaturityMonthYearField is a MONTHYEAR field
type RiskMaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.RiskMaturityMonthYear (1549)
func (f RiskMaturityMonthYearField) Tag() quickfix.Tag { return tag.RiskMaturityMonthYear }

//NewRiskMaturityMonthYear returns a new RiskMaturityMonthYearField initialized with val
func NewRiskMaturityMonthYear(val quickfix.FIXString) RiskMaturityMonthYearField {
	return RiskMaturityMonthYearField{val}
}

//RiskMaturityTimeField is a TZTIMEONLY field
type RiskMaturityTimeField struct{ quickfix.FIXString }

//Tag returns tag.RiskMaturityTime (1550)
func (f RiskMaturityTimeField) Tag() quickfix.Tag { return tag.RiskMaturityTime }

//NewRiskMaturityTime returns a new RiskMaturityTimeField initialized with val
func NewRiskMaturityTime(val quickfix.FIXString) RiskMaturityTimeField {
	return RiskMaturityTimeField{val}
}

//RiskProductField is a INT field
type RiskProductField struct{ quickfix.FIXInt }

//Tag returns tag.RiskProduct (1543)
func (f RiskProductField) Tag() quickfix.Tag { return tag.RiskProduct }

//NewRiskProduct returns a new RiskProductField initialized with val
func NewRiskProduct(val quickfix.FIXInt) RiskProductField {
	return RiskProductField{val}
}

//RiskProductComplexField is a STRING field
type RiskProductComplexField struct{ quickfix.FIXString }

//Tag returns tag.RiskProductComplex (1544)
func (f RiskProductComplexField) Tag() quickfix.Tag { return tag.RiskProductComplex }

//NewRiskProductComplex returns a new RiskProductComplexField initialized with val
func NewRiskProductComplex(val quickfix.FIXString) RiskProductComplexField {
	return RiskProductComplexField{val}
}

//RiskPutOrCallField is a INT field
type RiskPutOrCallField struct{ quickfix.FIXInt }

//Tag returns tag.RiskPutOrCall (1553)
func (f RiskPutOrCallField) Tag() quickfix.Tag { return tag.RiskPutOrCall }

//NewRiskPutOrCall returns a new RiskPutOrCallField initialized with val
func NewRiskPutOrCall(val quickfix.FIXInt) RiskPutOrCallField {
	return RiskPutOrCallField{val}
}

//RiskRestructuringTypeField is a STRING field
type RiskRestructuringTypeField struct{ quickfix.FIXString }

//Tag returns tag.RiskRestructuringType (1551)
func (f RiskRestructuringTypeField) Tag() quickfix.Tag { return tag.RiskRestructuringType }

//NewRiskRestructuringType returns a new RiskRestructuringTypeField initialized with val
func NewRiskRestructuringType(val quickfix.FIXString) RiskRestructuringTypeField {
	return RiskRestructuringTypeField{val}
}

//RiskSecurityAltIDField is a STRING field
type RiskSecurityAltIDField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecurityAltID (1541)
func (f RiskSecurityAltIDField) Tag() quickfix.Tag { return tag.RiskSecurityAltID }

//NewRiskSecurityAltID returns a new RiskSecurityAltIDField initialized with val
func NewRiskSecurityAltID(val quickfix.FIXString) RiskSecurityAltIDField {
	return RiskSecurityAltIDField{val}
}

//RiskSecurityAltIDSourceField is a STRING field
type RiskSecurityAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecurityAltIDSource (1542)
func (f RiskSecurityAltIDSourceField) Tag() quickfix.Tag { return tag.RiskSecurityAltIDSource }

//NewRiskSecurityAltIDSource returns a new RiskSecurityAltIDSourceField initialized with val
func NewRiskSecurityAltIDSource(val quickfix.FIXString) RiskSecurityAltIDSourceField {
	return RiskSecurityAltIDSourceField{val}
}

//RiskSecurityDescField is a STRING field
type RiskSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecurityDesc (1556)
func (f RiskSecurityDescField) Tag() quickfix.Tag { return tag.RiskSecurityDesc }

//NewRiskSecurityDesc returns a new RiskSecurityDescField initialized with val
func NewRiskSecurityDesc(val quickfix.FIXString) RiskSecurityDescField {
	return RiskSecurityDescField{val}
}

//RiskSecurityExchangeField is a EXCHANGE field
type RiskSecurityExchangeField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecurityExchange (1616)
func (f RiskSecurityExchangeField) Tag() quickfix.Tag { return tag.RiskSecurityExchange }

//NewRiskSecurityExchange returns a new RiskSecurityExchangeField initialized with val
func NewRiskSecurityExchange(val quickfix.FIXString) RiskSecurityExchangeField {
	return RiskSecurityExchangeField{val}
}

//RiskSecurityGroupField is a STRING field
type RiskSecurityGroupField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecurityGroup (1545)
func (f RiskSecurityGroupField) Tag() quickfix.Tag { return tag.RiskSecurityGroup }

//NewRiskSecurityGroup returns a new RiskSecurityGroupField initialized with val
func NewRiskSecurityGroup(val quickfix.FIXString) RiskSecurityGroupField {
	return RiskSecurityGroupField{val}
}

//RiskSecurityIDField is a STRING field
type RiskSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecurityID (1538)
func (f RiskSecurityIDField) Tag() quickfix.Tag { return tag.RiskSecurityID }

//NewRiskSecurityID returns a new RiskSecurityIDField initialized with val
func NewRiskSecurityID(val quickfix.FIXString) RiskSecurityIDField {
	return RiskSecurityIDField{val}
}

//RiskSecurityIDSourceField is a STRING field
type RiskSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecurityIDSource (1539)
func (f RiskSecurityIDSourceField) Tag() quickfix.Tag { return tag.RiskSecurityIDSource }

//NewRiskSecurityIDSource returns a new RiskSecurityIDSourceField initialized with val
func NewRiskSecurityIDSource(val quickfix.FIXString) RiskSecurityIDSourceField {
	return RiskSecurityIDSourceField{val}
}

//RiskSecuritySubTypeField is a STRING field
type RiskSecuritySubTypeField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecuritySubType (1548)
func (f RiskSecuritySubTypeField) Tag() quickfix.Tag { return tag.RiskSecuritySubType }

//NewRiskSecuritySubType returns a new RiskSecuritySubTypeField initialized with val
func NewRiskSecuritySubType(val quickfix.FIXString) RiskSecuritySubTypeField {
	return RiskSecuritySubTypeField{val}
}

//RiskSecurityTypeField is a STRING field
type RiskSecurityTypeField struct{ quickfix.FIXString }

//Tag returns tag.RiskSecurityType (1547)
func (f RiskSecurityTypeField) Tag() quickfix.Tag { return tag.RiskSecurityType }

//NewRiskSecurityType returns a new RiskSecurityTypeField initialized with val
func NewRiskSecurityType(val quickfix.FIXString) RiskSecurityTypeField {
	return RiskSecurityTypeField{val}
}

//RiskSeniorityField is a STRING field
type RiskSeniorityField struct{ quickfix.FIXString }

//Tag returns tag.RiskSeniority (1552)
func (f RiskSeniorityField) Tag() quickfix.Tag { return tag.RiskSeniority }

//NewRiskSeniority returns a new RiskSeniorityField initialized with val
func NewRiskSeniority(val quickfix.FIXString) RiskSeniorityField {
	return RiskSeniorityField{val}
}

//RiskSymbolField is a STRING field
type RiskSymbolField struct{ quickfix.FIXString }

//Tag returns tag.RiskSymbol (1536)
func (f RiskSymbolField) Tag() quickfix.Tag { return tag.RiskSymbol }

//NewRiskSymbol returns a new RiskSymbolField initialized with val
func NewRiskSymbol(val quickfix.FIXString) RiskSymbolField {
	return RiskSymbolField{val}
}

//RiskSymbolSfxField is a STRING field
type RiskSymbolSfxField struct{ quickfix.FIXString }

//Tag returns tag.RiskSymbolSfx (1537)
func (f RiskSymbolSfxField) Tag() quickfix.Tag { return tag.RiskSymbolSfx }

//NewRiskSymbolSfx returns a new RiskSymbolSfxField initialized with val
func NewRiskSymbolSfx(val quickfix.FIXString) RiskSymbolSfxField {
	return RiskSymbolSfxField{val}
}

//RiskWarningLevelNameField is a STRING field
type RiskWarningLevelNameField struct{ quickfix.FIXString }

//Tag returns tag.RiskWarningLevelName (1561)
func (f RiskWarningLevelNameField) Tag() quickfix.Tag { return tag.RiskWarningLevelName }

//NewRiskWarningLevelName returns a new RiskWarningLevelNameField initialized with val
func NewRiskWarningLevelName(val quickfix.FIXString) RiskWarningLevelNameField {
	return RiskWarningLevelNameField{val}
}

//RiskWarningLevelPercentField is a PERCENTAGE field
type RiskWarningLevelPercentField struct{ quickfix.FIXFloat }

//Tag returns tag.RiskWarningLevelPercent (1560)
func (f RiskWarningLevelPercentField) Tag() quickfix.Tag { return tag.RiskWarningLevelPercent }

//NewRiskWarningLevelPercent returns a new RiskWarningLevelPercentField initialized with val
func NewRiskWarningLevelPercent(val quickfix.FIXFloat) RiskWarningLevelPercentField {
	return RiskWarningLevelPercentField{val}
}

//RndPxField is a PRICE field
type RndPxField struct{ quickfix.FIXFloat }

//Tag returns tag.RndPx (991)
func (f RndPxField) Tag() quickfix.Tag { return tag.RndPx }

//NewRndPx returns a new RndPxField initialized with val
func NewRndPx(val quickfix.FIXFloat) RndPxField {
	return RndPxField{val}
}

//RootPartyIDField is a STRING field
type RootPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.RootPartyID (1117)
func (f RootPartyIDField) Tag() quickfix.Tag { return tag.RootPartyID }

//NewRootPartyID returns a new RootPartyIDField initialized with val
func NewRootPartyID(val quickfix.FIXString) RootPartyIDField {
	return RootPartyIDField{val}
}

//RootPartyIDSourceField is a CHAR field
type RootPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.RootPartyIDSource (1118)
func (f RootPartyIDSourceField) Tag() quickfix.Tag { return tag.RootPartyIDSource }

//NewRootPartyIDSource returns a new RootPartyIDSourceField initialized with val
func NewRootPartyIDSource(val quickfix.FIXString) RootPartyIDSourceField {
	return RootPartyIDSourceField{val}
}

//RootPartyRoleField is a INT field
type RootPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.RootPartyRole (1119)
func (f RootPartyRoleField) Tag() quickfix.Tag { return tag.RootPartyRole }

//NewRootPartyRole returns a new RootPartyRoleField initialized with val
func NewRootPartyRole(val quickfix.FIXInt) RootPartyRoleField {
	return RootPartyRoleField{val}
}

//RootPartySubIDField is a STRING field
type RootPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.RootPartySubID (1121)
func (f RootPartySubIDField) Tag() quickfix.Tag { return tag.RootPartySubID }

//NewRootPartySubID returns a new RootPartySubIDField initialized with val
func NewRootPartySubID(val quickfix.FIXString) RootPartySubIDField {
	return RootPartySubIDField{val}
}

//RootPartySubIDTypeField is a INT field
type RootPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.RootPartySubIDType (1122)
func (f RootPartySubIDTypeField) Tag() quickfix.Tag { return tag.RootPartySubIDType }

//NewRootPartySubIDType returns a new RootPartySubIDTypeField initialized with val
func NewRootPartySubIDType(val quickfix.FIXInt) RootPartySubIDTypeField {
	return RootPartySubIDTypeField{val}
}

//RoundLotField is a QTY field
type RoundLotField struct{ quickfix.FIXFloat }

//Tag returns tag.RoundLot (561)
func (f RoundLotField) Tag() quickfix.Tag { return tag.RoundLot }

//NewRoundLot returns a new RoundLotField initialized with val
func NewRoundLot(val quickfix.FIXFloat) RoundLotField {
	return RoundLotField{val}
}

//RoundingDirectionField is a CHAR field
type RoundingDirectionField struct{ quickfix.FIXString }

//Tag returns tag.RoundingDirection (468)
func (f RoundingDirectionField) Tag() quickfix.Tag { return tag.RoundingDirection }

//NewRoundingDirection returns a new RoundingDirectionField initialized with val
func NewRoundingDirection(val quickfix.FIXString) RoundingDirectionField {
	return RoundingDirectionField{val}
}

//RoundingModulusField is a FLOAT field
type RoundingModulusField struct{ quickfix.FIXFloat }

//Tag returns tag.RoundingModulus (469)
func (f RoundingModulusField) Tag() quickfix.Tag { return tag.RoundingModulus }

//NewRoundingModulus returns a new RoundingModulusField initialized with val
func NewRoundingModulus(val quickfix.FIXFloat) RoundingModulusField {
	return RoundingModulusField{val}
}

//RoutingIDField is a STRING field
type RoutingIDField struct{ quickfix.FIXString }

//Tag returns tag.RoutingID (217)
func (f RoutingIDField) Tag() quickfix.Tag { return tag.RoutingID }

//NewRoutingID returns a new RoutingIDField initialized with val
func NewRoutingID(val quickfix.FIXString) RoutingIDField {
	return RoutingIDField{val}
}

//RoutingTypeField is a INT field
type RoutingTypeField struct{ quickfix.FIXInt }

//Tag returns tag.RoutingType (216)
func (f RoutingTypeField) Tag() quickfix.Tag { return tag.RoutingType }

//NewRoutingType returns a new RoutingTypeField initialized with val
func NewRoutingType(val quickfix.FIXInt) RoutingTypeField {
	return RoutingTypeField{val}
}

//RptSeqField is a INT field
type RptSeqField struct{ quickfix.FIXInt }

//Tag returns tag.RptSeq (83)
func (f RptSeqField) Tag() quickfix.Tag { return tag.RptSeq }

//NewRptSeq returns a new RptSeqField initialized with val
func NewRptSeq(val quickfix.FIXInt) RptSeqField {
	return RptSeqField{val}
}

//RptSysField is a STRING field
type RptSysField struct{ quickfix.FIXString }

//Tag returns tag.RptSys (1135)
func (f RptSysField) Tag() quickfix.Tag { return tag.RptSys }

//NewRptSys returns a new RptSysField initialized with val
func NewRptSys(val quickfix.FIXString) RptSysField {
	return RptSysField{val}
}

//Rule80AField is a CHAR field
type Rule80AField struct{ quickfix.FIXString }

//Tag returns tag.Rule80A (47)
func (f Rule80AField) Tag() quickfix.Tag { return tag.Rule80A }

//NewRule80A returns a new Rule80AField initialized with val
func NewRule80A(val quickfix.FIXString) Rule80AField {
	return Rule80AField{val}
}

//ScopeField is a MULTIPLECHARVALUE field
type ScopeField struct{ quickfix.FIXString }

//Tag returns tag.Scope (546)
func (f ScopeField) Tag() quickfix.Tag { return tag.Scope }

//NewScope returns a new ScopeField initialized with val
func NewScope(val quickfix.FIXString) ScopeField {
	return ScopeField{val}
}

//SecDefStatusField is a INT field
type SecDefStatusField struct{ quickfix.FIXInt }

//Tag returns tag.SecDefStatus (653)
func (f SecDefStatusField) Tag() quickfix.Tag { return tag.SecDefStatus }

//NewSecDefStatus returns a new SecDefStatusField initialized with val
func NewSecDefStatus(val quickfix.FIXInt) SecDefStatusField {
	return SecDefStatusField{val}
}

//SecondaryAllocIDField is a STRING field
type SecondaryAllocIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryAllocID (793)
func (f SecondaryAllocIDField) Tag() quickfix.Tag { return tag.SecondaryAllocID }

//NewSecondaryAllocID returns a new SecondaryAllocIDField initialized with val
func NewSecondaryAllocID(val quickfix.FIXString) SecondaryAllocIDField {
	return SecondaryAllocIDField{val}
}

//SecondaryClOrdIDField is a STRING field
type SecondaryClOrdIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryClOrdID (526)
func (f SecondaryClOrdIDField) Tag() quickfix.Tag { return tag.SecondaryClOrdID }

//NewSecondaryClOrdID returns a new SecondaryClOrdIDField initialized with val
func NewSecondaryClOrdID(val quickfix.FIXString) SecondaryClOrdIDField {
	return SecondaryClOrdIDField{val}
}

//SecondaryDisplayQtyField is a QTY field
type SecondaryDisplayQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.SecondaryDisplayQty (1082)
func (f SecondaryDisplayQtyField) Tag() quickfix.Tag { return tag.SecondaryDisplayQty }

//NewSecondaryDisplayQty returns a new SecondaryDisplayQtyField initialized with val
func NewSecondaryDisplayQty(val quickfix.FIXFloat) SecondaryDisplayQtyField {
	return SecondaryDisplayQtyField{val}
}

//SecondaryExecIDField is a STRING field
type SecondaryExecIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryExecID (527)
func (f SecondaryExecIDField) Tag() quickfix.Tag { return tag.SecondaryExecID }

//NewSecondaryExecID returns a new SecondaryExecIDField initialized with val
func NewSecondaryExecID(val quickfix.FIXString) SecondaryExecIDField {
	return SecondaryExecIDField{val}
}

//SecondaryFirmTradeIDField is a STRING field
type SecondaryFirmTradeIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryFirmTradeID (1042)
func (f SecondaryFirmTradeIDField) Tag() quickfix.Tag { return tag.SecondaryFirmTradeID }

//NewSecondaryFirmTradeID returns a new SecondaryFirmTradeIDField initialized with val
func NewSecondaryFirmTradeID(val quickfix.FIXString) SecondaryFirmTradeIDField {
	return SecondaryFirmTradeIDField{val}
}

//SecondaryHighLimitPriceField is a PRICE field
type SecondaryHighLimitPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.SecondaryHighLimitPrice (1230)
func (f SecondaryHighLimitPriceField) Tag() quickfix.Tag { return tag.SecondaryHighLimitPrice }

//NewSecondaryHighLimitPrice returns a new SecondaryHighLimitPriceField initialized with val
func NewSecondaryHighLimitPrice(val quickfix.FIXFloat) SecondaryHighLimitPriceField {
	return SecondaryHighLimitPriceField{val}
}

//SecondaryIndividualAllocIDField is a STRING field
type SecondaryIndividualAllocIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryIndividualAllocID (989)
func (f SecondaryIndividualAllocIDField) Tag() quickfix.Tag { return tag.SecondaryIndividualAllocID }

//NewSecondaryIndividualAllocID returns a new SecondaryIndividualAllocIDField initialized with val
func NewSecondaryIndividualAllocID(val quickfix.FIXString) SecondaryIndividualAllocIDField {
	return SecondaryIndividualAllocIDField{val}
}

//SecondaryLowLimitPriceField is a PRICE field
type SecondaryLowLimitPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.SecondaryLowLimitPrice (1221)
func (f SecondaryLowLimitPriceField) Tag() quickfix.Tag { return tag.SecondaryLowLimitPrice }

//NewSecondaryLowLimitPrice returns a new SecondaryLowLimitPriceField initialized with val
func NewSecondaryLowLimitPrice(val quickfix.FIXFloat) SecondaryLowLimitPriceField {
	return SecondaryLowLimitPriceField{val}
}

//SecondaryOrderIDField is a STRING field
type SecondaryOrderIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryOrderID (198)
func (f SecondaryOrderIDField) Tag() quickfix.Tag { return tag.SecondaryOrderID }

//NewSecondaryOrderID returns a new SecondaryOrderIDField initialized with val
func NewSecondaryOrderID(val quickfix.FIXString) SecondaryOrderIDField {
	return SecondaryOrderIDField{val}
}

//SecondaryPriceLimitTypeField is a INT field
type SecondaryPriceLimitTypeField struct{ quickfix.FIXInt }

//Tag returns tag.SecondaryPriceLimitType (1305)
func (f SecondaryPriceLimitTypeField) Tag() quickfix.Tag { return tag.SecondaryPriceLimitType }

//NewSecondaryPriceLimitType returns a new SecondaryPriceLimitTypeField initialized with val
func NewSecondaryPriceLimitType(val quickfix.FIXInt) SecondaryPriceLimitTypeField {
	return SecondaryPriceLimitTypeField{val}
}

//SecondaryTradeIDField is a STRING field
type SecondaryTradeIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryTradeID (1040)
func (f SecondaryTradeIDField) Tag() quickfix.Tag { return tag.SecondaryTradeID }

//NewSecondaryTradeID returns a new SecondaryTradeIDField initialized with val
func NewSecondaryTradeID(val quickfix.FIXString) SecondaryTradeIDField {
	return SecondaryTradeIDField{val}
}

//SecondaryTradeReportIDField is a STRING field
type SecondaryTradeReportIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryTradeReportID (818)
func (f SecondaryTradeReportIDField) Tag() quickfix.Tag { return tag.SecondaryTradeReportID }

//NewSecondaryTradeReportID returns a new SecondaryTradeReportIDField initialized with val
func NewSecondaryTradeReportID(val quickfix.FIXString) SecondaryTradeReportIDField {
	return SecondaryTradeReportIDField{val}
}

//SecondaryTradeReportRefIDField is a STRING field
type SecondaryTradeReportRefIDField struct{ quickfix.FIXString }

//Tag returns tag.SecondaryTradeReportRefID (881)
func (f SecondaryTradeReportRefIDField) Tag() quickfix.Tag { return tag.SecondaryTradeReportRefID }

//NewSecondaryTradeReportRefID returns a new SecondaryTradeReportRefIDField initialized with val
func NewSecondaryTradeReportRefID(val quickfix.FIXString) SecondaryTradeReportRefIDField {
	return SecondaryTradeReportRefIDField{val}
}

//SecondaryTradingReferencePriceField is a PRICE field
type SecondaryTradingReferencePriceField struct{ quickfix.FIXFloat }

//Tag returns tag.SecondaryTradingReferencePrice (1240)
func (f SecondaryTradingReferencePriceField) Tag() quickfix.Tag {
	return tag.SecondaryTradingReferencePrice
}

//NewSecondaryTradingReferencePrice returns a new SecondaryTradingReferencePriceField initialized with val
func NewSecondaryTradingReferencePrice(val quickfix.FIXFloat) SecondaryTradingReferencePriceField {
	return SecondaryTradingReferencePriceField{val}
}

//SecondaryTrdTypeField is a INT field
type SecondaryTrdTypeField struct{ quickfix.FIXInt }

//Tag returns tag.SecondaryTrdType (855)
func (f SecondaryTrdTypeField) Tag() quickfix.Tag { return tag.SecondaryTrdType }

//NewSecondaryTrdType returns a new SecondaryTrdTypeField initialized with val
func NewSecondaryTrdType(val quickfix.FIXInt) SecondaryTrdTypeField {
	return SecondaryTrdTypeField{val}
}

//SecureDataField is a DATA field
type SecureDataField struct{ quickfix.FIXString }

//Tag returns tag.SecureData (91)
func (f SecureDataField) Tag() quickfix.Tag { return tag.SecureData }

//NewSecureData returns a new SecureDataField initialized with val
func NewSecureData(val quickfix.FIXString) SecureDataField {
	return SecureDataField{val}
}

//SecureDataLenField is a LENGTH field
type SecureDataLenField struct{ quickfix.FIXInt }

//Tag returns tag.SecureDataLen (90)
func (f SecureDataLenField) Tag() quickfix.Tag { return tag.SecureDataLen }

//NewSecureDataLen returns a new SecureDataLenField initialized with val
func NewSecureDataLen(val quickfix.FIXInt) SecureDataLenField {
	return SecureDataLenField{val}
}

//SecurityAltIDField is a STRING field
type SecurityAltIDField struct{ quickfix.FIXString }

//Tag returns tag.SecurityAltID (455)
func (f SecurityAltIDField) Tag() quickfix.Tag { return tag.SecurityAltID }

//NewSecurityAltID returns a new SecurityAltIDField initialized with val
func NewSecurityAltID(val quickfix.FIXString) SecurityAltIDField {
	return SecurityAltIDField{val}
}

//SecurityAltIDSourceField is a STRING field
type SecurityAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.SecurityAltIDSource (456)
func (f SecurityAltIDSourceField) Tag() quickfix.Tag { return tag.SecurityAltIDSource }

//NewSecurityAltIDSource returns a new SecurityAltIDSourceField initialized with val
func NewSecurityAltIDSource(val quickfix.FIXString) SecurityAltIDSourceField {
	return SecurityAltIDSourceField{val}
}

//SecurityDescField is a STRING field
type SecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.SecurityDesc (107)
func (f SecurityDescField) Tag() quickfix.Tag { return tag.SecurityDesc }

//NewSecurityDesc returns a new SecurityDescField initialized with val
func NewSecurityDesc(val quickfix.FIXString) SecurityDescField {
	return SecurityDescField{val}
}

//SecurityExchangeField is a EXCHANGE field
type SecurityExchangeField struct{ quickfix.FIXString }

//Tag returns tag.SecurityExchange (207)
func (f SecurityExchangeField) Tag() quickfix.Tag { return tag.SecurityExchange }

//NewSecurityExchange returns a new SecurityExchangeField initialized with val
func NewSecurityExchange(val quickfix.FIXString) SecurityExchangeField {
	return SecurityExchangeField{val}
}

//SecurityGroupField is a STRING field
type SecurityGroupField struct{ quickfix.FIXString }

//Tag returns tag.SecurityGroup (1151)
func (f SecurityGroupField) Tag() quickfix.Tag { return tag.SecurityGroup }

//NewSecurityGroup returns a new SecurityGroupField initialized with val
func NewSecurityGroup(val quickfix.FIXString) SecurityGroupField {
	return SecurityGroupField{val}
}

//SecurityIDField is a STRING field
type SecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.SecurityID (48)
func (f SecurityIDField) Tag() quickfix.Tag { return tag.SecurityID }

//NewSecurityID returns a new SecurityIDField initialized with val
func NewSecurityID(val quickfix.FIXString) SecurityIDField {
	return SecurityIDField{val}
}

//SecurityIDSourceField is a STRING field
type SecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.SecurityIDSource (22)
func (f SecurityIDSourceField) Tag() quickfix.Tag { return tag.SecurityIDSource }

//NewSecurityIDSource returns a new SecurityIDSourceField initialized with val
func NewSecurityIDSource(val quickfix.FIXString) SecurityIDSourceField {
	return SecurityIDSourceField{val}
}

//SecurityListDescField is a STRING field
type SecurityListDescField struct{ quickfix.FIXString }

//Tag returns tag.SecurityListDesc (1467)
func (f SecurityListDescField) Tag() quickfix.Tag { return tag.SecurityListDesc }

//NewSecurityListDesc returns a new SecurityListDescField initialized with val
func NewSecurityListDesc(val quickfix.FIXString) SecurityListDescField {
	return SecurityListDescField{val}
}

//SecurityListIDField is a STRING field
type SecurityListIDField struct{ quickfix.FIXString }

//Tag returns tag.SecurityListID (1465)
func (f SecurityListIDField) Tag() quickfix.Tag { return tag.SecurityListID }

//NewSecurityListID returns a new SecurityListIDField initialized with val
func NewSecurityListID(val quickfix.FIXString) SecurityListIDField {
	return SecurityListIDField{val}
}

//SecurityListRefIDField is a STRING field
type SecurityListRefIDField struct{ quickfix.FIXString }

//Tag returns tag.SecurityListRefID (1466)
func (f SecurityListRefIDField) Tag() quickfix.Tag { return tag.SecurityListRefID }

//NewSecurityListRefID returns a new SecurityListRefIDField initialized with val
func NewSecurityListRefID(val quickfix.FIXString) SecurityListRefIDField {
	return SecurityListRefIDField{val}
}

//SecurityListRequestTypeField is a INT field
type SecurityListRequestTypeField struct{ quickfix.FIXInt }

//Tag returns tag.SecurityListRequestType (559)
func (f SecurityListRequestTypeField) Tag() quickfix.Tag { return tag.SecurityListRequestType }

//NewSecurityListRequestType returns a new SecurityListRequestTypeField initialized with val
func NewSecurityListRequestType(val quickfix.FIXInt) SecurityListRequestTypeField {
	return SecurityListRequestTypeField{val}
}

//SecurityListTypeField is a INT field
type SecurityListTypeField struct{ quickfix.FIXInt }

//Tag returns tag.SecurityListType (1470)
func (f SecurityListTypeField) Tag() quickfix.Tag { return tag.SecurityListType }

//NewSecurityListType returns a new SecurityListTypeField initialized with val
func NewSecurityListType(val quickfix.FIXInt) SecurityListTypeField {
	return SecurityListTypeField{val}
}

//SecurityListTypeSourceField is a INT field
type SecurityListTypeSourceField struct{ quickfix.FIXInt }

//Tag returns tag.SecurityListTypeSource (1471)
func (f SecurityListTypeSourceField) Tag() quickfix.Tag { return tag.SecurityListTypeSource }

//NewSecurityListTypeSource returns a new SecurityListTypeSourceField initialized with val
func NewSecurityListTypeSource(val quickfix.FIXInt) SecurityListTypeSourceField {
	return SecurityListTypeSourceField{val}
}

//SecurityReportIDField is a INT field
type SecurityReportIDField struct{ quickfix.FIXInt }

//Tag returns tag.SecurityReportID (964)
func (f SecurityReportIDField) Tag() quickfix.Tag { return tag.SecurityReportID }

//NewSecurityReportID returns a new SecurityReportIDField initialized with val
func NewSecurityReportID(val quickfix.FIXInt) SecurityReportIDField {
	return SecurityReportIDField{val}
}

//SecurityReqIDField is a STRING field
type SecurityReqIDField struct{ quickfix.FIXString }

//Tag returns tag.SecurityReqID (320)
func (f SecurityReqIDField) Tag() quickfix.Tag { return tag.SecurityReqID }

//NewSecurityReqID returns a new SecurityReqIDField initialized with val
func NewSecurityReqID(val quickfix.FIXString) SecurityReqIDField {
	return SecurityReqIDField{val}
}

//SecurityRequestResultField is a INT field
type SecurityRequestResultField struct{ quickfix.FIXInt }

//Tag returns tag.SecurityRequestResult (560)
func (f SecurityRequestResultField) Tag() quickfix.Tag { return tag.SecurityRequestResult }

//NewSecurityRequestResult returns a new SecurityRequestResultField initialized with val
func NewSecurityRequestResult(val quickfix.FIXInt) SecurityRequestResultField {
	return SecurityRequestResultField{val}
}

//SecurityRequestTypeField is a INT field
type SecurityRequestTypeField struct{ quickfix.FIXInt }

//Tag returns tag.SecurityRequestType (321)
func (f SecurityRequestTypeField) Tag() quickfix.Tag { return tag.SecurityRequestType }

//NewSecurityRequestType returns a new SecurityRequestTypeField initialized with val
func NewSecurityRequestType(val quickfix.FIXInt) SecurityRequestTypeField {
	return SecurityRequestTypeField{val}
}

//SecurityResponseIDField is a STRING field
type SecurityResponseIDField struct{ quickfix.FIXString }

//Tag returns tag.SecurityResponseID (322)
func (f SecurityResponseIDField) Tag() quickfix.Tag { return tag.SecurityResponseID }

//NewSecurityResponseID returns a new SecurityResponseIDField initialized with val
func NewSecurityResponseID(val quickfix.FIXString) SecurityResponseIDField {
	return SecurityResponseIDField{val}
}

//SecurityResponseTypeField is a INT field
type SecurityResponseTypeField struct{ quickfix.FIXInt }

//Tag returns tag.SecurityResponseType (323)
func (f SecurityResponseTypeField) Tag() quickfix.Tag { return tag.SecurityResponseType }

//NewSecurityResponseType returns a new SecurityResponseTypeField initialized with val
func NewSecurityResponseType(val quickfix.FIXInt) SecurityResponseTypeField {
	return SecurityResponseTypeField{val}
}

//SecuritySettlAgentAcctNameField is a STRING field
type SecuritySettlAgentAcctNameField struct{ quickfix.FIXString }

//Tag returns tag.SecuritySettlAgentAcctName (179)
func (f SecuritySettlAgentAcctNameField) Tag() quickfix.Tag { return tag.SecuritySettlAgentAcctName }

//NewSecuritySettlAgentAcctName returns a new SecuritySettlAgentAcctNameField initialized with val
func NewSecuritySettlAgentAcctName(val quickfix.FIXString) SecuritySettlAgentAcctNameField {
	return SecuritySettlAgentAcctNameField{val}
}

//SecuritySettlAgentAcctNumField is a STRING field
type SecuritySettlAgentAcctNumField struct{ quickfix.FIXString }

//Tag returns tag.SecuritySettlAgentAcctNum (178)
func (f SecuritySettlAgentAcctNumField) Tag() quickfix.Tag { return tag.SecuritySettlAgentAcctNum }

//NewSecuritySettlAgentAcctNum returns a new SecuritySettlAgentAcctNumField initialized with val
func NewSecuritySettlAgentAcctNum(val quickfix.FIXString) SecuritySettlAgentAcctNumField {
	return SecuritySettlAgentAcctNumField{val}
}

//SecuritySettlAgentCodeField is a STRING field
type SecuritySettlAgentCodeField struct{ quickfix.FIXString }

//Tag returns tag.SecuritySettlAgentCode (177)
func (f SecuritySettlAgentCodeField) Tag() quickfix.Tag { return tag.SecuritySettlAgentCode }

//NewSecuritySettlAgentCode returns a new SecuritySettlAgentCodeField initialized with val
func NewSecuritySettlAgentCode(val quickfix.FIXString) SecuritySettlAgentCodeField {
	return SecuritySettlAgentCodeField{val}
}

//SecuritySettlAgentContactNameField is a STRING field
type SecuritySettlAgentContactNameField struct{ quickfix.FIXString }

//Tag returns tag.SecuritySettlAgentContactName (180)
func (f SecuritySettlAgentContactNameField) Tag() quickfix.Tag {
	return tag.SecuritySettlAgentContactName
}

//NewSecuritySettlAgentContactName returns a new SecuritySettlAgentContactNameField initialized with val
func NewSecuritySettlAgentContactName(val quickfix.FIXString) SecuritySettlAgentContactNameField {
	return SecuritySettlAgentContactNameField{val}
}

//SecuritySettlAgentContactPhoneField is a STRING field
type SecuritySettlAgentContactPhoneField struct{ quickfix.FIXString }

//Tag returns tag.SecuritySettlAgentContactPhone (181)
func (f SecuritySettlAgentContactPhoneField) Tag() quickfix.Tag {
	return tag.SecuritySettlAgentContactPhone
}

//NewSecuritySettlAgentContactPhone returns a new SecuritySettlAgentContactPhoneField initialized with val
func NewSecuritySettlAgentContactPhone(val quickfix.FIXString) SecuritySettlAgentContactPhoneField {
	return SecuritySettlAgentContactPhoneField{val}
}

//SecuritySettlAgentNameField is a STRING field
type SecuritySettlAgentNameField struct{ quickfix.FIXString }

//Tag returns tag.SecuritySettlAgentName (176)
func (f SecuritySettlAgentNameField) Tag() quickfix.Tag { return tag.SecuritySettlAgentName }

//NewSecuritySettlAgentName returns a new SecuritySettlAgentNameField initialized with val
func NewSecuritySettlAgentName(val quickfix.FIXString) SecuritySettlAgentNameField {
	return SecuritySettlAgentNameField{val}
}

//SecurityStatusField is a STRING field
type SecurityStatusField struct{ quickfix.FIXString }

//Tag returns tag.SecurityStatus (965)
func (f SecurityStatusField) Tag() quickfix.Tag { return tag.SecurityStatus }

//NewSecurityStatus returns a new SecurityStatusField initialized with val
func NewSecurityStatus(val quickfix.FIXString) SecurityStatusField {
	return SecurityStatusField{val}
}

//SecurityStatusReqIDField is a STRING field
type SecurityStatusReqIDField struct{ quickfix.FIXString }

//Tag returns tag.SecurityStatusReqID (324)
func (f SecurityStatusReqIDField) Tag() quickfix.Tag { return tag.SecurityStatusReqID }

//NewSecurityStatusReqID returns a new SecurityStatusReqIDField initialized with val
func NewSecurityStatusReqID(val quickfix.FIXString) SecurityStatusReqIDField {
	return SecurityStatusReqIDField{val}
}

//SecuritySubTypeField is a STRING field
type SecuritySubTypeField struct{ quickfix.FIXString }

//Tag returns tag.SecuritySubType (762)
func (f SecuritySubTypeField) Tag() quickfix.Tag { return tag.SecuritySubType }

//NewSecuritySubType returns a new SecuritySubTypeField initialized with val
func NewSecuritySubType(val quickfix.FIXString) SecuritySubTypeField {
	return SecuritySubTypeField{val}
}

//SecurityTradingEventField is a INT field
type SecurityTradingEventField struct{ quickfix.FIXInt }

//Tag returns tag.SecurityTradingEvent (1174)
func (f SecurityTradingEventField) Tag() quickfix.Tag { return tag.SecurityTradingEvent }

//NewSecurityTradingEvent returns a new SecurityTradingEventField initialized with val
func NewSecurityTradingEvent(val quickfix.FIXInt) SecurityTradingEventField {
	return SecurityTradingEventField{val}
}

//SecurityTradingStatusField is a INT field
type SecurityTradingStatusField struct{ quickfix.FIXInt }

//Tag returns tag.SecurityTradingStatus (326)
func (f SecurityTradingStatusField) Tag() quickfix.Tag { return tag.SecurityTradingStatus }

//NewSecurityTradingStatus returns a new SecurityTradingStatusField initialized with val
func NewSecurityTradingStatus(val quickfix.FIXInt) SecurityTradingStatusField {
	return SecurityTradingStatusField{val}
}

//SecurityTypeField is a STRING field
type SecurityTypeField struct{ quickfix.FIXString }

//Tag returns tag.SecurityType (167)
func (f SecurityTypeField) Tag() quickfix.Tag { return tag.SecurityType }

//NewSecurityType returns a new SecurityTypeField initialized with val
func NewSecurityType(val quickfix.FIXString) SecurityTypeField {
	return SecurityTypeField{val}
}

//SecurityUpdateActionField is a CHAR field
type SecurityUpdateActionField struct{ quickfix.FIXString }

//Tag returns tag.SecurityUpdateAction (980)
func (f SecurityUpdateActionField) Tag() quickfix.Tag { return tag.SecurityUpdateAction }

//NewSecurityUpdateAction returns a new SecurityUpdateActionField initialized with val
func NewSecurityUpdateAction(val quickfix.FIXString) SecurityUpdateActionField {
	return SecurityUpdateActionField{val}
}

//SecurityXMLField is a XMLDATA field
type SecurityXMLField struct{ quickfix.FIXString }

//Tag returns tag.SecurityXML (1185)
func (f SecurityXMLField) Tag() quickfix.Tag { return tag.SecurityXML }

//NewSecurityXML returns a new SecurityXMLField initialized with val
func NewSecurityXML(val quickfix.FIXString) SecurityXMLField {
	return SecurityXMLField{val}
}

//SecurityXMLLenField is a LENGTH field
type SecurityXMLLenField struct{ quickfix.FIXInt }

//Tag returns tag.SecurityXMLLen (1184)
func (f SecurityXMLLenField) Tag() quickfix.Tag { return tag.SecurityXMLLen }

//NewSecurityXMLLen returns a new SecurityXMLLenField initialized with val
func NewSecurityXMLLen(val quickfix.FIXInt) SecurityXMLLenField {
	return SecurityXMLLenField{val}
}

//SecurityXMLSchemaField is a STRING field
type SecurityXMLSchemaField struct{ quickfix.FIXString }

//Tag returns tag.SecurityXMLSchema (1186)
func (f SecurityXMLSchemaField) Tag() quickfix.Tag { return tag.SecurityXMLSchema }

//NewSecurityXMLSchema returns a new SecurityXMLSchemaField initialized with val
func NewSecurityXMLSchema(val quickfix.FIXString) SecurityXMLSchemaField {
	return SecurityXMLSchemaField{val}
}

//SellVolumeField is a QTY field
type SellVolumeField struct{ quickfix.FIXFloat }

//Tag returns tag.SellVolume (331)
func (f SellVolumeField) Tag() quickfix.Tag { return tag.SellVolume }

//NewSellVolume returns a new SellVolumeField initialized with val
func NewSellVolume(val quickfix.FIXFloat) SellVolumeField {
	return SellVolumeField{val}
}

//SellerDaysField is a INT field
type SellerDaysField struct{ quickfix.FIXInt }

//Tag returns tag.SellerDays (287)
func (f SellerDaysField) Tag() quickfix.Tag { return tag.SellerDays }

//NewSellerDays returns a new SellerDaysField initialized with val
func NewSellerDays(val quickfix.FIXInt) SellerDaysField {
	return SellerDaysField{val}
}

//SenderCompIDField is a STRING field
type SenderCompIDField struct{ quickfix.FIXString }

//Tag returns tag.SenderCompID (49)
func (f SenderCompIDField) Tag() quickfix.Tag { return tag.SenderCompID }

//NewSenderCompID returns a new SenderCompIDField initialized with val
func NewSenderCompID(val quickfix.FIXString) SenderCompIDField {
	return SenderCompIDField{val}
}

//SenderLocationIDField is a STRING field
type SenderLocationIDField struct{ quickfix.FIXString }

//Tag returns tag.SenderLocationID (142)
func (f SenderLocationIDField) Tag() quickfix.Tag { return tag.SenderLocationID }

//NewSenderLocationID returns a new SenderLocationIDField initialized with val
func NewSenderLocationID(val quickfix.FIXString) SenderLocationIDField {
	return SenderLocationIDField{val}
}

//SenderSubIDField is a STRING field
type SenderSubIDField struct{ quickfix.FIXString }

//Tag returns tag.SenderSubID (50)
func (f SenderSubIDField) Tag() quickfix.Tag { return tag.SenderSubID }

//NewSenderSubID returns a new SenderSubIDField initialized with val
func NewSenderSubID(val quickfix.FIXString) SenderSubIDField {
	return SenderSubIDField{val}
}

//SendingDateField is a LOCALMKTDATE field
type SendingDateField struct{ quickfix.FIXString }

//Tag returns tag.SendingDate (51)
func (f SendingDateField) Tag() quickfix.Tag { return tag.SendingDate }

//NewSendingDate returns a new SendingDateField initialized with val
func NewSendingDate(val quickfix.FIXString) SendingDateField {
	return SendingDateField{val}
}

//SendingTimeField is a UTCTIMESTAMP field
type SendingTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.SendingTime (52)
func (f SendingTimeField) Tag() quickfix.Tag { return tag.SendingTime }

//NewSendingTime returns a new SendingTimeField initialized with val
func NewSendingTime(val time.Time) SendingTimeField {
	return SendingTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewSendingTimeNoMillis returns a new SendingTimeField initialized with val without millisecs
func NewSendingTimeNoMillis(val time.Time) SendingTimeField {
	return SendingTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//SeniorityField is a STRING field
type SeniorityField struct{ quickfix.FIXString }

//Tag returns tag.Seniority (1450)
func (f SeniorityField) Tag() quickfix.Tag { return tag.Seniority }

//NewSeniority returns a new SeniorityField initialized with val
func NewSeniority(val quickfix.FIXString) SeniorityField {
	return SeniorityField{val}
}

//SessionRejectReasonField is a INT field
type SessionRejectReasonField struct{ quickfix.FIXInt }

//Tag returns tag.SessionRejectReason (373)
func (f SessionRejectReasonField) Tag() quickfix.Tag { return tag.SessionRejectReason }

//NewSessionRejectReason returns a new SessionRejectReasonField initialized with val
func NewSessionRejectReason(val quickfix.FIXInt) SessionRejectReasonField {
	return SessionRejectReasonField{val}
}

//SessionStatusField is a INT field
type SessionStatusField struct{ quickfix.FIXInt }

//Tag returns tag.SessionStatus (1409)
func (f SessionStatusField) Tag() quickfix.Tag { return tag.SessionStatus }

//NewSessionStatus returns a new SessionStatusField initialized with val
func NewSessionStatus(val quickfix.FIXInt) SessionStatusField {
	return SessionStatusField{val}
}

//SettlBrkrCodeField is a STRING field
type SettlBrkrCodeField struct{ quickfix.FIXString }

//Tag returns tag.SettlBrkrCode (174)
func (f SettlBrkrCodeField) Tag() quickfix.Tag { return tag.SettlBrkrCode }

//NewSettlBrkrCode returns a new SettlBrkrCodeField initialized with val
func NewSettlBrkrCode(val quickfix.FIXString) SettlBrkrCodeField {
	return SettlBrkrCodeField{val}
}

//SettlCurrAmtField is a AMT field
type SettlCurrAmtField struct{ quickfix.FIXFloat }

//Tag returns tag.SettlCurrAmt (119)
func (f SettlCurrAmtField) Tag() quickfix.Tag { return tag.SettlCurrAmt }

//NewSettlCurrAmt returns a new SettlCurrAmtField initialized with val
func NewSettlCurrAmt(val quickfix.FIXFloat) SettlCurrAmtField {
	return SettlCurrAmtField{val}
}

//SettlCurrBidFxRateField is a FLOAT field
type SettlCurrBidFxRateField struct{ quickfix.FIXFloat }

//Tag returns tag.SettlCurrBidFxRate (656)
func (f SettlCurrBidFxRateField) Tag() quickfix.Tag { return tag.SettlCurrBidFxRate }

//NewSettlCurrBidFxRate returns a new SettlCurrBidFxRateField initialized with val
func NewSettlCurrBidFxRate(val quickfix.FIXFloat) SettlCurrBidFxRateField {
	return SettlCurrBidFxRateField{val}
}

//SettlCurrFxRateField is a FLOAT field
type SettlCurrFxRateField struct{ quickfix.FIXFloat }

//Tag returns tag.SettlCurrFxRate (155)
func (f SettlCurrFxRateField) Tag() quickfix.Tag { return tag.SettlCurrFxRate }

//NewSettlCurrFxRate returns a new SettlCurrFxRateField initialized with val
func NewSettlCurrFxRate(val quickfix.FIXFloat) SettlCurrFxRateField {
	return SettlCurrFxRateField{val}
}

//SettlCurrFxRateCalcField is a CHAR field
type SettlCurrFxRateCalcField struct{ quickfix.FIXString }

//Tag returns tag.SettlCurrFxRateCalc (156)
func (f SettlCurrFxRateCalcField) Tag() quickfix.Tag { return tag.SettlCurrFxRateCalc }

//NewSettlCurrFxRateCalc returns a new SettlCurrFxRateCalcField initialized with val
func NewSettlCurrFxRateCalc(val quickfix.FIXString) SettlCurrFxRateCalcField {
	return SettlCurrFxRateCalcField{val}
}

//SettlCurrOfferFxRateField is a FLOAT field
type SettlCurrOfferFxRateField struct{ quickfix.FIXFloat }

//Tag returns tag.SettlCurrOfferFxRate (657)
func (f SettlCurrOfferFxRateField) Tag() quickfix.Tag { return tag.SettlCurrOfferFxRate }

//NewSettlCurrOfferFxRate returns a new SettlCurrOfferFxRateField initialized with val
func NewSettlCurrOfferFxRate(val quickfix.FIXFloat) SettlCurrOfferFxRateField {
	return SettlCurrOfferFxRateField{val}
}

//SettlCurrencyField is a CURRENCY field
type SettlCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.SettlCurrency (120)
func (f SettlCurrencyField) Tag() quickfix.Tag { return tag.SettlCurrency }

//NewSettlCurrency returns a new SettlCurrencyField initialized with val
func NewSettlCurrency(val quickfix.FIXString) SettlCurrencyField {
	return SettlCurrencyField{val}
}

//SettlDateField is a LOCALMKTDATE field
type SettlDateField struct{ quickfix.FIXString }

//Tag returns tag.SettlDate (64)
func (f SettlDateField) Tag() quickfix.Tag { return tag.SettlDate }

//NewSettlDate returns a new SettlDateField initialized with val
func NewSettlDate(val quickfix.FIXString) SettlDateField {
	return SettlDateField{val}
}

//SettlDate2Field is a LOCALMKTDATE field
type SettlDate2Field struct{ quickfix.FIXString }

//Tag returns tag.SettlDate2 (193)
func (f SettlDate2Field) Tag() quickfix.Tag { return tag.SettlDate2 }

//NewSettlDate2 returns a new SettlDate2Field initialized with val
func NewSettlDate2(val quickfix.FIXString) SettlDate2Field {
	return SettlDate2Field{val}
}

//SettlDeliveryTypeField is a INT field
type SettlDeliveryTypeField struct{ quickfix.FIXInt }

//Tag returns tag.SettlDeliveryType (172)
func (f SettlDeliveryTypeField) Tag() quickfix.Tag { return tag.SettlDeliveryType }

//NewSettlDeliveryType returns a new SettlDeliveryTypeField initialized with val
func NewSettlDeliveryType(val quickfix.FIXInt) SettlDeliveryTypeField {
	return SettlDeliveryTypeField{val}
}

//SettlDepositoryCodeField is a STRING field
type SettlDepositoryCodeField struct{ quickfix.FIXString }

//Tag returns tag.SettlDepositoryCode (173)
func (f SettlDepositoryCodeField) Tag() quickfix.Tag { return tag.SettlDepositoryCode }

//NewSettlDepositoryCode returns a new SettlDepositoryCodeField initialized with val
func NewSettlDepositoryCode(val quickfix.FIXString) SettlDepositoryCodeField {
	return SettlDepositoryCodeField{val}
}

//SettlInstCodeField is a STRING field
type SettlInstCodeField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstCode (175)
func (f SettlInstCodeField) Tag() quickfix.Tag { return tag.SettlInstCode }

//NewSettlInstCode returns a new SettlInstCodeField initialized with val
func NewSettlInstCode(val quickfix.FIXString) SettlInstCodeField {
	return SettlInstCodeField{val}
}

//SettlInstIDField is a STRING field
type SettlInstIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstID (162)
func (f SettlInstIDField) Tag() quickfix.Tag { return tag.SettlInstID }

//NewSettlInstID returns a new SettlInstIDField initialized with val
func NewSettlInstID(val quickfix.FIXString) SettlInstIDField {
	return SettlInstIDField{val}
}

//SettlInstModeField is a CHAR field
type SettlInstModeField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstMode (160)
func (f SettlInstModeField) Tag() quickfix.Tag { return tag.SettlInstMode }

//NewSettlInstMode returns a new SettlInstModeField initialized with val
func NewSettlInstMode(val quickfix.FIXString) SettlInstModeField {
	return SettlInstModeField{val}
}

//SettlInstMsgIDField is a STRING field
type SettlInstMsgIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstMsgID (777)
func (f SettlInstMsgIDField) Tag() quickfix.Tag { return tag.SettlInstMsgID }

//NewSettlInstMsgID returns a new SettlInstMsgIDField initialized with val
func NewSettlInstMsgID(val quickfix.FIXString) SettlInstMsgIDField {
	return SettlInstMsgIDField{val}
}

//SettlInstRefIDField is a STRING field
type SettlInstRefIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstRefID (214)
func (f SettlInstRefIDField) Tag() quickfix.Tag { return tag.SettlInstRefID }

//NewSettlInstRefID returns a new SettlInstRefIDField initialized with val
func NewSettlInstRefID(val quickfix.FIXString) SettlInstRefIDField {
	return SettlInstRefIDField{val}
}

//SettlInstReqIDField is a STRING field
type SettlInstReqIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstReqID (791)
func (f SettlInstReqIDField) Tag() quickfix.Tag { return tag.SettlInstReqID }

//NewSettlInstReqID returns a new SettlInstReqIDField initialized with val
func NewSettlInstReqID(val quickfix.FIXString) SettlInstReqIDField {
	return SettlInstReqIDField{val}
}

//SettlInstReqRejCodeField is a INT field
type SettlInstReqRejCodeField struct{ quickfix.FIXInt }

//Tag returns tag.SettlInstReqRejCode (792)
func (f SettlInstReqRejCodeField) Tag() quickfix.Tag { return tag.SettlInstReqRejCode }

//NewSettlInstReqRejCode returns a new SettlInstReqRejCodeField initialized with val
func NewSettlInstReqRejCode(val quickfix.FIXInt) SettlInstReqRejCodeField {
	return SettlInstReqRejCodeField{val}
}

//SettlInstSourceField is a CHAR field
type SettlInstSourceField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstSource (165)
func (f SettlInstSourceField) Tag() quickfix.Tag { return tag.SettlInstSource }

//NewSettlInstSource returns a new SettlInstSourceField initialized with val
func NewSettlInstSource(val quickfix.FIXString) SettlInstSourceField {
	return SettlInstSourceField{val}
}

//SettlInstTransTypeField is a CHAR field
type SettlInstTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.SettlInstTransType (163)
func (f SettlInstTransTypeField) Tag() quickfix.Tag { return tag.SettlInstTransType }

//NewSettlInstTransType returns a new SettlInstTransTypeField initialized with val
func NewSettlInstTransType(val quickfix.FIXString) SettlInstTransTypeField {
	return SettlInstTransTypeField{val}
}

//SettlLocationField is a STRING field
type SettlLocationField struct{ quickfix.FIXString }

//Tag returns tag.SettlLocation (166)
func (f SettlLocationField) Tag() quickfix.Tag { return tag.SettlLocation }

//NewSettlLocation returns a new SettlLocationField initialized with val
func NewSettlLocation(val quickfix.FIXString) SettlLocationField {
	return SettlLocationField{val}
}

//SettlMethodField is a CHAR field
type SettlMethodField struct{ quickfix.FIXString }

//Tag returns tag.SettlMethod (1193)
func (f SettlMethodField) Tag() quickfix.Tag { return tag.SettlMethod }

//NewSettlMethod returns a new SettlMethodField initialized with val
func NewSettlMethod(val quickfix.FIXString) SettlMethodField {
	return SettlMethodField{val}
}

//SettlObligIDField is a STRING field
type SettlObligIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlObligID (1161)
func (f SettlObligIDField) Tag() quickfix.Tag { return tag.SettlObligID }

//NewSettlObligID returns a new SettlObligIDField initialized with val
func NewSettlObligID(val quickfix.FIXString) SettlObligIDField {
	return SettlObligIDField{val}
}

//SettlObligModeField is a INT field
type SettlObligModeField struct{ quickfix.FIXInt }

//Tag returns tag.SettlObligMode (1159)
func (f SettlObligModeField) Tag() quickfix.Tag { return tag.SettlObligMode }

//NewSettlObligMode returns a new SettlObligModeField initialized with val
func NewSettlObligMode(val quickfix.FIXInt) SettlObligModeField {
	return SettlObligModeField{val}
}

//SettlObligMsgIDField is a STRING field
type SettlObligMsgIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlObligMsgID (1160)
func (f SettlObligMsgIDField) Tag() quickfix.Tag { return tag.SettlObligMsgID }

//NewSettlObligMsgID returns a new SettlObligMsgIDField initialized with val
func NewSettlObligMsgID(val quickfix.FIXString) SettlObligMsgIDField {
	return SettlObligMsgIDField{val}
}

//SettlObligRefIDField is a STRING field
type SettlObligRefIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlObligRefID (1163)
func (f SettlObligRefIDField) Tag() quickfix.Tag { return tag.SettlObligRefID }

//NewSettlObligRefID returns a new SettlObligRefIDField initialized with val
func NewSettlObligRefID(val quickfix.FIXString) SettlObligRefIDField {
	return SettlObligRefIDField{val}
}

//SettlObligSourceField is a CHAR field
type SettlObligSourceField struct{ quickfix.FIXString }

//Tag returns tag.SettlObligSource (1164)
func (f SettlObligSourceField) Tag() quickfix.Tag { return tag.SettlObligSource }

//NewSettlObligSource returns a new SettlObligSourceField initialized with val
func NewSettlObligSource(val quickfix.FIXString) SettlObligSourceField {
	return SettlObligSourceField{val}
}

//SettlObligTransTypeField is a CHAR field
type SettlObligTransTypeField struct{ quickfix.FIXString }

//Tag returns tag.SettlObligTransType (1162)
func (f SettlObligTransTypeField) Tag() quickfix.Tag { return tag.SettlObligTransType }

//NewSettlObligTransType returns a new SettlObligTransTypeField initialized with val
func NewSettlObligTransType(val quickfix.FIXString) SettlObligTransTypeField {
	return SettlObligTransTypeField{val}
}

//SettlPartyIDField is a STRING field
type SettlPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlPartyID (782)
func (f SettlPartyIDField) Tag() quickfix.Tag { return tag.SettlPartyID }

//NewSettlPartyID returns a new SettlPartyIDField initialized with val
func NewSettlPartyID(val quickfix.FIXString) SettlPartyIDField {
	return SettlPartyIDField{val}
}

//SettlPartyIDSourceField is a CHAR field
type SettlPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.SettlPartyIDSource (783)
func (f SettlPartyIDSourceField) Tag() quickfix.Tag { return tag.SettlPartyIDSource }

//NewSettlPartyIDSource returns a new SettlPartyIDSourceField initialized with val
func NewSettlPartyIDSource(val quickfix.FIXString) SettlPartyIDSourceField {
	return SettlPartyIDSourceField{val}
}

//SettlPartyRoleField is a INT field
type SettlPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.SettlPartyRole (784)
func (f SettlPartyRoleField) Tag() quickfix.Tag { return tag.SettlPartyRole }

//NewSettlPartyRole returns a new SettlPartyRoleField initialized with val
func NewSettlPartyRole(val quickfix.FIXInt) SettlPartyRoleField {
	return SettlPartyRoleField{val}
}

//SettlPartySubIDField is a STRING field
type SettlPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlPartySubID (785)
func (f SettlPartySubIDField) Tag() quickfix.Tag { return tag.SettlPartySubID }

//NewSettlPartySubID returns a new SettlPartySubIDField initialized with val
func NewSettlPartySubID(val quickfix.FIXString) SettlPartySubIDField {
	return SettlPartySubIDField{val}
}

//SettlPartySubIDTypeField is a INT field
type SettlPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.SettlPartySubIDType (786)
func (f SettlPartySubIDTypeField) Tag() quickfix.Tag { return tag.SettlPartySubIDType }

//NewSettlPartySubIDType returns a new SettlPartySubIDTypeField initialized with val
func NewSettlPartySubIDType(val quickfix.FIXInt) SettlPartySubIDTypeField {
	return SettlPartySubIDTypeField{val}
}

//SettlPriceField is a PRICE field
type SettlPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.SettlPrice (730)
func (f SettlPriceField) Tag() quickfix.Tag { return tag.SettlPrice }

//NewSettlPrice returns a new SettlPriceField initialized with val
func NewSettlPrice(val quickfix.FIXFloat) SettlPriceField {
	return SettlPriceField{val}
}

//SettlPriceTypeField is a INT field
type SettlPriceTypeField struct{ quickfix.FIXInt }

//Tag returns tag.SettlPriceType (731)
func (f SettlPriceTypeField) Tag() quickfix.Tag { return tag.SettlPriceType }

//NewSettlPriceType returns a new SettlPriceTypeField initialized with val
func NewSettlPriceType(val quickfix.FIXInt) SettlPriceTypeField {
	return SettlPriceTypeField{val}
}

//SettlSessIDField is a STRING field
type SettlSessIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlSessID (716)
func (f SettlSessIDField) Tag() quickfix.Tag { return tag.SettlSessID }

//NewSettlSessID returns a new SettlSessIDField initialized with val
func NewSettlSessID(val quickfix.FIXString) SettlSessIDField {
	return SettlSessIDField{val}
}

//SettlSessSubIDField is a STRING field
type SettlSessSubIDField struct{ quickfix.FIXString }

//Tag returns tag.SettlSessSubID (717)
func (f SettlSessSubIDField) Tag() quickfix.Tag { return tag.SettlSessSubID }

//NewSettlSessSubID returns a new SettlSessSubIDField initialized with val
func NewSettlSessSubID(val quickfix.FIXString) SettlSessSubIDField {
	return SettlSessSubIDField{val}
}

//SettlTypeField is a STRING field
type SettlTypeField struct{ quickfix.FIXString }

//Tag returns tag.SettlType (63)
func (f SettlTypeField) Tag() quickfix.Tag { return tag.SettlType }

//NewSettlType returns a new SettlTypeField initialized with val
func NewSettlType(val quickfix.FIXString) SettlTypeField {
	return SettlTypeField{val}
}

//SettleOnOpenFlagField is a STRING field
type SettleOnOpenFlagField struct{ quickfix.FIXString }

//Tag returns tag.SettleOnOpenFlag (966)
func (f SettleOnOpenFlagField) Tag() quickfix.Tag { return tag.SettleOnOpenFlag }

//NewSettleOnOpenFlag returns a new SettleOnOpenFlagField initialized with val
func NewSettleOnOpenFlag(val quickfix.FIXString) SettleOnOpenFlagField {
	return SettleOnOpenFlagField{val}
}

//SettlementCycleNoField is a INT field
type SettlementCycleNoField struct{ quickfix.FIXInt }

//Tag returns tag.SettlementCycleNo (1153)
func (f SettlementCycleNoField) Tag() quickfix.Tag { return tag.SettlementCycleNo }

//NewSettlementCycleNo returns a new SettlementCycleNoField initialized with val
func NewSettlementCycleNo(val quickfix.FIXInt) SettlementCycleNoField {
	return SettlementCycleNoField{val}
}

//SettlmntTypField is a CHAR field
type SettlmntTypField struct{ quickfix.FIXString }

//Tag returns tag.SettlmntTyp (63)
func (f SettlmntTypField) Tag() quickfix.Tag { return tag.SettlmntTyp }

//NewSettlmntTyp returns a new SettlmntTypField initialized with val
func NewSettlmntTyp(val quickfix.FIXString) SettlmntTypField {
	return SettlmntTypField{val}
}

//SharedCommissionField is a AMT field
type SharedCommissionField struct{ quickfix.FIXFloat }

//Tag returns tag.SharedCommission (858)
func (f SharedCommissionField) Tag() quickfix.Tag { return tag.SharedCommission }

//NewSharedCommission returns a new SharedCommissionField initialized with val
func NewSharedCommission(val quickfix.FIXFloat) SharedCommissionField {
	return SharedCommissionField{val}
}

//SharesField is a QTY field
type SharesField struct{ quickfix.FIXFloat }

//Tag returns tag.Shares (53)
func (f SharesField) Tag() quickfix.Tag { return tag.Shares }

//NewShares returns a new SharesField initialized with val
func NewShares(val quickfix.FIXFloat) SharesField {
	return SharesField{val}
}

//ShortQtyField is a QTY field
type ShortQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.ShortQty (705)
func (f ShortQtyField) Tag() quickfix.Tag { return tag.ShortQty }

//NewShortQty returns a new ShortQtyField initialized with val
func NewShortQty(val quickfix.FIXFloat) ShortQtyField {
	return ShortQtyField{val}
}

//ShortSaleReasonField is a INT field
type ShortSaleReasonField struct{ quickfix.FIXInt }

//Tag returns tag.ShortSaleReason (853)
func (f ShortSaleReasonField) Tag() quickfix.Tag { return tag.ShortSaleReason }

//NewShortSaleReason returns a new ShortSaleReasonField initialized with val
func NewShortSaleReason(val quickfix.FIXInt) ShortSaleReasonField {
	return ShortSaleReasonField{val}
}

//SideField is a CHAR field
type SideField struct{ quickfix.FIXString }

//Tag returns tag.Side (54)
func (f SideField) Tag() quickfix.Tag { return tag.Side }

//NewSide returns a new SideField initialized with val
func NewSide(val quickfix.FIXString) SideField {
	return SideField{val}
}

//SideComplianceIDField is a STRING field
type SideComplianceIDField struct{ quickfix.FIXString }

//Tag returns tag.SideComplianceID (659)
func (f SideComplianceIDField) Tag() quickfix.Tag { return tag.SideComplianceID }

//NewSideComplianceID returns a new SideComplianceIDField initialized with val
func NewSideComplianceID(val quickfix.FIXString) SideComplianceIDField {
	return SideComplianceIDField{val}
}

//SideCurrencyField is a CURRENCY field
type SideCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.SideCurrency (1154)
func (f SideCurrencyField) Tag() quickfix.Tag { return tag.SideCurrency }

//NewSideCurrency returns a new SideCurrencyField initialized with val
func NewSideCurrency(val quickfix.FIXString) SideCurrencyField {
	return SideCurrencyField{val}
}

//SideExecIDField is a STRING field
type SideExecIDField struct{ quickfix.FIXString }

//Tag returns tag.SideExecID (1427)
func (f SideExecIDField) Tag() quickfix.Tag { return tag.SideExecID }

//NewSideExecID returns a new SideExecIDField initialized with val
func NewSideExecID(val quickfix.FIXString) SideExecIDField {
	return SideExecIDField{val}
}

//SideFillStationCdField is a STRING field
type SideFillStationCdField struct{ quickfix.FIXString }

//Tag returns tag.SideFillStationCd (1006)
func (f SideFillStationCdField) Tag() quickfix.Tag { return tag.SideFillStationCd }

//NewSideFillStationCd returns a new SideFillStationCdField initialized with val
func NewSideFillStationCd(val quickfix.FIXString) SideFillStationCdField {
	return SideFillStationCdField{val}
}

//SideGrossTradeAmtField is a AMT field
type SideGrossTradeAmtField struct{ quickfix.FIXFloat }

//Tag returns tag.SideGrossTradeAmt (1072)
func (f SideGrossTradeAmtField) Tag() quickfix.Tag { return tag.SideGrossTradeAmt }

//NewSideGrossTradeAmt returns a new SideGrossTradeAmtField initialized with val
func NewSideGrossTradeAmt(val quickfix.FIXFloat) SideGrossTradeAmtField {
	return SideGrossTradeAmtField{val}
}

//SideLastQtyField is a INT field
type SideLastQtyField struct{ quickfix.FIXInt }

//Tag returns tag.SideLastQty (1009)
func (f SideLastQtyField) Tag() quickfix.Tag { return tag.SideLastQty }

//NewSideLastQty returns a new SideLastQtyField initialized with val
func NewSideLastQty(val quickfix.FIXInt) SideLastQtyField {
	return SideLastQtyField{val}
}

//SideLiquidityIndField is a INT field
type SideLiquidityIndField struct{ quickfix.FIXInt }

//Tag returns tag.SideLiquidityInd (1444)
func (f SideLiquidityIndField) Tag() quickfix.Tag { return tag.SideLiquidityInd }

//NewSideLiquidityInd returns a new SideLiquidityIndField initialized with val
func NewSideLiquidityInd(val quickfix.FIXInt) SideLiquidityIndField {
	return SideLiquidityIndField{val}
}

//SideMultiLegReportingTypeField is a INT field
type SideMultiLegReportingTypeField struct{ quickfix.FIXInt }

//Tag returns tag.SideMultiLegReportingType (752)
func (f SideMultiLegReportingTypeField) Tag() quickfix.Tag { return tag.SideMultiLegReportingType }

//NewSideMultiLegReportingType returns a new SideMultiLegReportingTypeField initialized with val
func NewSideMultiLegReportingType(val quickfix.FIXInt) SideMultiLegReportingTypeField {
	return SideMultiLegReportingTypeField{val}
}

//SideQtyField is a INT field
type SideQtyField struct{ quickfix.FIXInt }

//Tag returns tag.SideQty (1009)
func (f SideQtyField) Tag() quickfix.Tag { return tag.SideQty }

//NewSideQty returns a new SideQtyField initialized with val
func NewSideQty(val quickfix.FIXInt) SideQtyField {
	return SideQtyField{val}
}

//SideReasonCdField is a STRING field
type SideReasonCdField struct{ quickfix.FIXString }

//Tag returns tag.SideReasonCd (1007)
func (f SideReasonCdField) Tag() quickfix.Tag { return tag.SideReasonCd }

//NewSideReasonCd returns a new SideReasonCdField initialized with val
func NewSideReasonCd(val quickfix.FIXString) SideReasonCdField {
	return SideReasonCdField{val}
}

//SideSettlCurrencyField is a CURRENCY field
type SideSettlCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.SideSettlCurrency (1155)
func (f SideSettlCurrencyField) Tag() quickfix.Tag { return tag.SideSettlCurrency }

//NewSideSettlCurrency returns a new SideSettlCurrencyField initialized with val
func NewSideSettlCurrency(val quickfix.FIXString) SideSettlCurrencyField {
	return SideSettlCurrencyField{val}
}

//SideTimeInForceField is a UTCTIMESTAMP field
type SideTimeInForceField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.SideTimeInForce (962)
func (f SideTimeInForceField) Tag() quickfix.Tag { return tag.SideTimeInForce }

//NewSideTimeInForce returns a new SideTimeInForceField initialized with val
func NewSideTimeInForce(val time.Time) SideTimeInForceField {
	return SideTimeInForceField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewSideTimeInForceNoMillis returns a new SideTimeInForceField initialized with val without millisecs
func NewSideTimeInForceNoMillis(val time.Time) SideTimeInForceField {
	return SideTimeInForceField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//SideTradeReportIDField is a STRING field
type SideTradeReportIDField struct{ quickfix.FIXString }

//Tag returns tag.SideTradeReportID (1005)
func (f SideTradeReportIDField) Tag() quickfix.Tag { return tag.SideTradeReportID }

//NewSideTradeReportID returns a new SideTradeReportIDField initialized with val
func NewSideTradeReportID(val quickfix.FIXString) SideTradeReportIDField {
	return SideTradeReportIDField{val}
}

//SideTrdRegTimestampField is a UTCTIMESTAMP field
type SideTrdRegTimestampField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.SideTrdRegTimestamp (1012)
func (f SideTrdRegTimestampField) Tag() quickfix.Tag { return tag.SideTrdRegTimestamp }

//NewSideTrdRegTimestamp returns a new SideTrdRegTimestampField initialized with val
func NewSideTrdRegTimestamp(val time.Time) SideTrdRegTimestampField {
	return SideTrdRegTimestampField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewSideTrdRegTimestampNoMillis returns a new SideTrdRegTimestampField initialized with val without millisecs
func NewSideTrdRegTimestampNoMillis(val time.Time) SideTrdRegTimestampField {
	return SideTrdRegTimestampField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//SideTrdRegTimestampSrcField is a STRING field
type SideTrdRegTimestampSrcField struct{ quickfix.FIXString }

//Tag returns tag.SideTrdRegTimestampSrc (1014)
func (f SideTrdRegTimestampSrcField) Tag() quickfix.Tag { return tag.SideTrdRegTimestampSrc }

//NewSideTrdRegTimestampSrc returns a new SideTrdRegTimestampSrcField initialized with val
func NewSideTrdRegTimestampSrc(val quickfix.FIXString) SideTrdRegTimestampSrcField {
	return SideTrdRegTimestampSrcField{val}
}

//SideTrdRegTimestampTypeField is a INT field
type SideTrdRegTimestampTypeField struct{ quickfix.FIXInt }

//Tag returns tag.SideTrdRegTimestampType (1013)
func (f SideTrdRegTimestampTypeField) Tag() quickfix.Tag { return tag.SideTrdRegTimestampType }

//NewSideTrdRegTimestampType returns a new SideTrdRegTimestampTypeField initialized with val
func NewSideTrdRegTimestampType(val quickfix.FIXInt) SideTrdRegTimestampTypeField {
	return SideTrdRegTimestampTypeField{val}
}

//SideTrdSubTypField is a INT field
type SideTrdSubTypField struct{ quickfix.FIXInt }

//Tag returns tag.SideTrdSubTyp (1008)
func (f SideTrdSubTypField) Tag() quickfix.Tag { return tag.SideTrdSubTyp }

//NewSideTrdSubTyp returns a new SideTrdSubTypField initialized with val
func NewSideTrdSubTyp(val quickfix.FIXInt) SideTrdSubTypField {
	return SideTrdSubTypField{val}
}

//SideValue1Field is a AMT field
type SideValue1Field struct{ quickfix.FIXFloat }

//Tag returns tag.SideValue1 (396)
func (f SideValue1Field) Tag() quickfix.Tag { return tag.SideValue1 }

//NewSideValue1 returns a new SideValue1Field initialized with val
func NewSideValue1(val quickfix.FIXFloat) SideValue1Field {
	return SideValue1Field{val}
}

//SideValue2Field is a AMT field
type SideValue2Field struct{ quickfix.FIXFloat }

//Tag returns tag.SideValue2 (397)
func (f SideValue2Field) Tag() quickfix.Tag { return tag.SideValue2 }

//NewSideValue2 returns a new SideValue2Field initialized with val
func NewSideValue2(val quickfix.FIXFloat) SideValue2Field {
	return SideValue2Field{val}
}

//SideValueIndField is a INT field
type SideValueIndField struct{ quickfix.FIXInt }

//Tag returns tag.SideValueInd (401)
func (f SideValueIndField) Tag() quickfix.Tag { return tag.SideValueInd }

//NewSideValueInd returns a new SideValueIndField initialized with val
func NewSideValueInd(val quickfix.FIXInt) SideValueIndField {
	return SideValueIndField{val}
}

//SignatureField is a DATA field
type SignatureField struct{ quickfix.FIXString }

//Tag returns tag.Signature (89)
func (f SignatureField) Tag() quickfix.Tag { return tag.Signature }

//NewSignature returns a new SignatureField initialized with val
func NewSignature(val quickfix.FIXString) SignatureField {
	return SignatureField{val}
}

//SignatureLengthField is a LENGTH field
type SignatureLengthField struct{ quickfix.FIXInt }

//Tag returns tag.SignatureLength (93)
func (f SignatureLengthField) Tag() quickfix.Tag { return tag.SignatureLength }

//NewSignatureLength returns a new SignatureLengthField initialized with val
func NewSignatureLength(val quickfix.FIXInt) SignatureLengthField {
	return SignatureLengthField{val}
}

//SolicitedFlagField is a BOOLEAN field
type SolicitedFlagField struct{ quickfix.FIXBoolean }

//Tag returns tag.SolicitedFlag (377)
func (f SolicitedFlagField) Tag() quickfix.Tag { return tag.SolicitedFlag }

//NewSolicitedFlag returns a new SolicitedFlagField initialized with val
func NewSolicitedFlag(val quickfix.FIXBoolean) SolicitedFlagField {
	return SolicitedFlagField{val}
}

//SpreadField is a PRICEOFFSET field
type SpreadField struct{ quickfix.FIXFloat }

//Tag returns tag.Spread (218)
func (f SpreadField) Tag() quickfix.Tag { return tag.Spread }

//NewSpread returns a new SpreadField initialized with val
func NewSpread(val quickfix.FIXFloat) SpreadField {
	return SpreadField{val}
}

//SpreadToBenchmarkField is a PRICEOFFSET field
type SpreadToBenchmarkField struct{ quickfix.FIXFloat }

//Tag returns tag.SpreadToBenchmark (218)
func (f SpreadToBenchmarkField) Tag() quickfix.Tag { return tag.SpreadToBenchmark }

//NewSpreadToBenchmark returns a new SpreadToBenchmarkField initialized with val
func NewSpreadToBenchmark(val quickfix.FIXFloat) SpreadToBenchmarkField {
	return SpreadToBenchmarkField{val}
}

//StandInstDbIDField is a STRING field
type StandInstDbIDField struct{ quickfix.FIXString }

//Tag returns tag.StandInstDbID (171)
func (f StandInstDbIDField) Tag() quickfix.Tag { return tag.StandInstDbID }

//NewStandInstDbID returns a new StandInstDbIDField initialized with val
func NewStandInstDbID(val quickfix.FIXString) StandInstDbIDField {
	return StandInstDbIDField{val}
}

//StandInstDbNameField is a STRING field
type StandInstDbNameField struct{ quickfix.FIXString }

//Tag returns tag.StandInstDbName (170)
func (f StandInstDbNameField) Tag() quickfix.Tag { return tag.StandInstDbName }

//NewStandInstDbName returns a new StandInstDbNameField initialized with val
func NewStandInstDbName(val quickfix.FIXString) StandInstDbNameField {
	return StandInstDbNameField{val}
}

//StandInstDbTypeField is a INT field
type StandInstDbTypeField struct{ quickfix.FIXInt }

//Tag returns tag.StandInstDbType (169)
func (f StandInstDbTypeField) Tag() quickfix.Tag { return tag.StandInstDbType }

//NewStandInstDbType returns a new StandInstDbTypeField initialized with val
func NewStandInstDbType(val quickfix.FIXInt) StandInstDbTypeField {
	return StandInstDbTypeField{val}
}

//StartCashField is a AMT field
type StartCashField struct{ quickfix.FIXFloat }

//Tag returns tag.StartCash (921)
func (f StartCashField) Tag() quickfix.Tag { return tag.StartCash }

//NewStartCash returns a new StartCashField initialized with val
func NewStartCash(val quickfix.FIXFloat) StartCashField {
	return StartCashField{val}
}

//StartDateField is a LOCALMKTDATE field
type StartDateField struct{ quickfix.FIXString }

//Tag returns tag.StartDate (916)
func (f StartDateField) Tag() quickfix.Tag { return tag.StartDate }

//NewStartDate returns a new StartDateField initialized with val
func NewStartDate(val quickfix.FIXString) StartDateField {
	return StartDateField{val}
}

//StartMaturityMonthYearField is a MONTHYEAR field
type StartMaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.StartMaturityMonthYear (1241)
func (f StartMaturityMonthYearField) Tag() quickfix.Tag { return tag.StartMaturityMonthYear }

//NewStartMaturityMonthYear returns a new StartMaturityMonthYearField initialized with val
func NewStartMaturityMonthYear(val quickfix.FIXString) StartMaturityMonthYearField {
	return StartMaturityMonthYearField{val}
}

//StartStrikePxRangeField is a PRICE field
type StartStrikePxRangeField struct{ quickfix.FIXFloat }

//Tag returns tag.StartStrikePxRange (1202)
func (f StartStrikePxRangeField) Tag() quickfix.Tag { return tag.StartStrikePxRange }

//NewStartStrikePxRange returns a new StartStrikePxRangeField initialized with val
func NewStartStrikePxRange(val quickfix.FIXFloat) StartStrikePxRangeField {
	return StartStrikePxRangeField{val}
}

//StartTickPriceRangeField is a PRICE field
type StartTickPriceRangeField struct{ quickfix.FIXFloat }

//Tag returns tag.StartTickPriceRange (1206)
func (f StartTickPriceRangeField) Tag() quickfix.Tag { return tag.StartTickPriceRange }

//NewStartTickPriceRange returns a new StartTickPriceRangeField initialized with val
func NewStartTickPriceRange(val quickfix.FIXFloat) StartTickPriceRangeField {
	return StartTickPriceRangeField{val}
}

//StateOrProvinceOfIssueField is a STRING field
type StateOrProvinceOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.StateOrProvinceOfIssue (471)
func (f StateOrProvinceOfIssueField) Tag() quickfix.Tag { return tag.StateOrProvinceOfIssue }

//NewStateOrProvinceOfIssue returns a new StateOrProvinceOfIssueField initialized with val
func NewStateOrProvinceOfIssue(val quickfix.FIXString) StateOrProvinceOfIssueField {
	return StateOrProvinceOfIssueField{val}
}

//StatsTypeField is a INT field
type StatsTypeField struct{ quickfix.FIXInt }

//Tag returns tag.StatsType (1176)
func (f StatsTypeField) Tag() quickfix.Tag { return tag.StatsType }

//NewStatsType returns a new StatsTypeField initialized with val
func NewStatsType(val quickfix.FIXInt) StatsTypeField {
	return StatsTypeField{val}
}

//StatusTextField is a STRING field
type StatusTextField struct{ quickfix.FIXString }

//Tag returns tag.StatusText (929)
func (f StatusTextField) Tag() quickfix.Tag { return tag.StatusText }

//NewStatusText returns a new StatusTextField initialized with val
func NewStatusText(val quickfix.FIXString) StatusTextField {
	return StatusTextField{val}
}

//StatusValueField is a INT field
type StatusValueField struct{ quickfix.FIXInt }

//Tag returns tag.StatusValue (928)
func (f StatusValueField) Tag() quickfix.Tag { return tag.StatusValue }

//NewStatusValue returns a new StatusValueField initialized with val
func NewStatusValue(val quickfix.FIXInt) StatusValueField {
	return StatusValueField{val}
}

//StipulationTypeField is a STRING field
type StipulationTypeField struct{ quickfix.FIXString }

//Tag returns tag.StipulationType (233)
func (f StipulationTypeField) Tag() quickfix.Tag { return tag.StipulationType }

//NewStipulationType returns a new StipulationTypeField initialized with val
func NewStipulationType(val quickfix.FIXString) StipulationTypeField {
	return StipulationTypeField{val}
}

//StipulationValueField is a STRING field
type StipulationValueField struct{ quickfix.FIXString }

//Tag returns tag.StipulationValue (234)
func (f StipulationValueField) Tag() quickfix.Tag { return tag.StipulationValue }

//NewStipulationValue returns a new StipulationValueField initialized with val
func NewStipulationValue(val quickfix.FIXString) StipulationValueField {
	return StipulationValueField{val}
}

//StopPxField is a PRICE field
type StopPxField struct{ quickfix.FIXFloat }

//Tag returns tag.StopPx (99)
func (f StopPxField) Tag() quickfix.Tag { return tag.StopPx }

//NewStopPx returns a new StopPxField initialized with val
func NewStopPx(val quickfix.FIXFloat) StopPxField {
	return StopPxField{val}
}

//StrategyParameterNameField is a STRING field
type StrategyParameterNameField struct{ quickfix.FIXString }

//Tag returns tag.StrategyParameterName (958)
func (f StrategyParameterNameField) Tag() quickfix.Tag { return tag.StrategyParameterName }

//NewStrategyParameterName returns a new StrategyParameterNameField initialized with val
func NewStrategyParameterName(val quickfix.FIXString) StrategyParameterNameField {
	return StrategyParameterNameField{val}
}

//StrategyParameterTypeField is a INT field
type StrategyParameterTypeField struct{ quickfix.FIXInt }

//Tag returns tag.StrategyParameterType (959)
func (f StrategyParameterTypeField) Tag() quickfix.Tag { return tag.StrategyParameterType }

//NewStrategyParameterType returns a new StrategyParameterTypeField initialized with val
func NewStrategyParameterType(val quickfix.FIXInt) StrategyParameterTypeField {
	return StrategyParameterTypeField{val}
}

//StrategyParameterValueField is a STRING field
type StrategyParameterValueField struct{ quickfix.FIXString }

//Tag returns tag.StrategyParameterValue (960)
func (f StrategyParameterValueField) Tag() quickfix.Tag { return tag.StrategyParameterValue }

//NewStrategyParameterValue returns a new StrategyParameterValueField initialized with val
func NewStrategyParameterValue(val quickfix.FIXString) StrategyParameterValueField {
	return StrategyParameterValueField{val}
}

//StreamAsgnAckTypeField is a INT field
type StreamAsgnAckTypeField struct{ quickfix.FIXInt }

//Tag returns tag.StreamAsgnAckType (1503)
func (f StreamAsgnAckTypeField) Tag() quickfix.Tag { return tag.StreamAsgnAckType }

//NewStreamAsgnAckType returns a new StreamAsgnAckTypeField initialized with val
func NewStreamAsgnAckType(val quickfix.FIXInt) StreamAsgnAckTypeField {
	return StreamAsgnAckTypeField{val}
}

//StreamAsgnRejReasonField is a INT field
type StreamAsgnRejReasonField struct{ quickfix.FIXInt }

//Tag returns tag.StreamAsgnRejReason (1502)
func (f StreamAsgnRejReasonField) Tag() quickfix.Tag { return tag.StreamAsgnRejReason }

//NewStreamAsgnRejReason returns a new StreamAsgnRejReasonField initialized with val
func NewStreamAsgnRejReason(val quickfix.FIXInt) StreamAsgnRejReasonField {
	return StreamAsgnRejReasonField{val}
}

//StreamAsgnReqIDField is a STRING field
type StreamAsgnReqIDField struct{ quickfix.FIXString }

//Tag returns tag.StreamAsgnReqID (1497)
func (f StreamAsgnReqIDField) Tag() quickfix.Tag { return tag.StreamAsgnReqID }

//NewStreamAsgnReqID returns a new StreamAsgnReqIDField initialized with val
func NewStreamAsgnReqID(val quickfix.FIXString) StreamAsgnReqIDField {
	return StreamAsgnReqIDField{val}
}

//StreamAsgnReqTypeField is a INT field
type StreamAsgnReqTypeField struct{ quickfix.FIXInt }

//Tag returns tag.StreamAsgnReqType (1498)
func (f StreamAsgnReqTypeField) Tag() quickfix.Tag { return tag.StreamAsgnReqType }

//NewStreamAsgnReqType returns a new StreamAsgnReqTypeField initialized with val
func NewStreamAsgnReqType(val quickfix.FIXInt) StreamAsgnReqTypeField {
	return StreamAsgnReqTypeField{val}
}

//StreamAsgnRptIDField is a STRING field
type StreamAsgnRptIDField struct{ quickfix.FIXString }

//Tag returns tag.StreamAsgnRptID (1501)
func (f StreamAsgnRptIDField) Tag() quickfix.Tag { return tag.StreamAsgnRptID }

//NewStreamAsgnRptID returns a new StreamAsgnRptIDField initialized with val
func NewStreamAsgnRptID(val quickfix.FIXString) StreamAsgnRptIDField {
	return StreamAsgnRptIDField{val}
}

//StreamAsgnTypeField is a INT field
type StreamAsgnTypeField struct{ quickfix.FIXInt }

//Tag returns tag.StreamAsgnType (1617)
func (f StreamAsgnTypeField) Tag() quickfix.Tag { return tag.StreamAsgnType }

//NewStreamAsgnType returns a new StreamAsgnTypeField initialized with val
func NewStreamAsgnType(val quickfix.FIXInt) StreamAsgnTypeField {
	return StreamAsgnTypeField{val}
}

//StrikeCurrencyField is a CURRENCY field
type StrikeCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.StrikeCurrency (947)
func (f StrikeCurrencyField) Tag() quickfix.Tag { return tag.StrikeCurrency }

//NewStrikeCurrency returns a new StrikeCurrencyField initialized with val
func NewStrikeCurrency(val quickfix.FIXString) StrikeCurrencyField {
	return StrikeCurrencyField{val}
}

//StrikeExerciseStyleField is a INT field
type StrikeExerciseStyleField struct{ quickfix.FIXInt }

//Tag returns tag.StrikeExerciseStyle (1304)
func (f StrikeExerciseStyleField) Tag() quickfix.Tag { return tag.StrikeExerciseStyle }

//NewStrikeExerciseStyle returns a new StrikeExerciseStyleField initialized with val
func NewStrikeExerciseStyle(val quickfix.FIXInt) StrikeExerciseStyleField {
	return StrikeExerciseStyleField{val}
}

//StrikeIncrementField is a FLOAT field
type StrikeIncrementField struct{ quickfix.FIXFloat }

//Tag returns tag.StrikeIncrement (1204)
func (f StrikeIncrementField) Tag() quickfix.Tag { return tag.StrikeIncrement }

//NewStrikeIncrement returns a new StrikeIncrementField initialized with val
func NewStrikeIncrement(val quickfix.FIXFloat) StrikeIncrementField {
	return StrikeIncrementField{val}
}

//StrikeMultiplierField is a FLOAT field
type StrikeMultiplierField struct{ quickfix.FIXFloat }

//Tag returns tag.StrikeMultiplier (967)
func (f StrikeMultiplierField) Tag() quickfix.Tag { return tag.StrikeMultiplier }

//NewStrikeMultiplier returns a new StrikeMultiplierField initialized with val
func NewStrikeMultiplier(val quickfix.FIXFloat) StrikeMultiplierField {
	return StrikeMultiplierField{val}
}

//StrikePriceField is a PRICE field
type StrikePriceField struct{ quickfix.FIXFloat }

//Tag returns tag.StrikePrice (202)
func (f StrikePriceField) Tag() quickfix.Tag { return tag.StrikePrice }

//NewStrikePrice returns a new StrikePriceField initialized with val
func NewStrikePrice(val quickfix.FIXFloat) StrikePriceField {
	return StrikePriceField{val}
}

//StrikePriceBoundaryMethodField is a INT field
type StrikePriceBoundaryMethodField struct{ quickfix.FIXInt }

//Tag returns tag.StrikePriceBoundaryMethod (1479)
func (f StrikePriceBoundaryMethodField) Tag() quickfix.Tag { return tag.StrikePriceBoundaryMethod }

//NewStrikePriceBoundaryMethod returns a new StrikePriceBoundaryMethodField initialized with val
func NewStrikePriceBoundaryMethod(val quickfix.FIXInt) StrikePriceBoundaryMethodField {
	return StrikePriceBoundaryMethodField{val}
}

//StrikePriceBoundaryPrecisionField is a PERCENTAGE field
type StrikePriceBoundaryPrecisionField struct{ quickfix.FIXFloat }

//Tag returns tag.StrikePriceBoundaryPrecision (1480)
func (f StrikePriceBoundaryPrecisionField) Tag() quickfix.Tag { return tag.StrikePriceBoundaryPrecision }

//NewStrikePriceBoundaryPrecision returns a new StrikePriceBoundaryPrecisionField initialized with val
func NewStrikePriceBoundaryPrecision(val quickfix.FIXFloat) StrikePriceBoundaryPrecisionField {
	return StrikePriceBoundaryPrecisionField{val}
}

//StrikePriceDeterminationMethodField is a INT field
type StrikePriceDeterminationMethodField struct{ quickfix.FIXInt }

//Tag returns tag.StrikePriceDeterminationMethod (1478)
func (f StrikePriceDeterminationMethodField) Tag() quickfix.Tag {
	return tag.StrikePriceDeterminationMethod
}

//NewStrikePriceDeterminationMethod returns a new StrikePriceDeterminationMethodField initialized with val
func NewStrikePriceDeterminationMethod(val quickfix.FIXInt) StrikePriceDeterminationMethodField {
	return StrikePriceDeterminationMethodField{val}
}

//StrikeRuleIDField is a STRING field
type StrikeRuleIDField struct{ quickfix.FIXString }

//Tag returns tag.StrikeRuleID (1223)
func (f StrikeRuleIDField) Tag() quickfix.Tag { return tag.StrikeRuleID }

//NewStrikeRuleID returns a new StrikeRuleIDField initialized with val
func NewStrikeRuleID(val quickfix.FIXString) StrikeRuleIDField {
	return StrikeRuleIDField{val}
}

//StrikeTimeField is a UTCTIMESTAMP field
type StrikeTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.StrikeTime (443)
func (f StrikeTimeField) Tag() quickfix.Tag { return tag.StrikeTime }

//NewStrikeTime returns a new StrikeTimeField initialized with val
func NewStrikeTime(val time.Time) StrikeTimeField {
	return StrikeTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewStrikeTimeNoMillis returns a new StrikeTimeField initialized with val without millisecs
func NewStrikeTimeNoMillis(val time.Time) StrikeTimeField {
	return StrikeTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//StrikeValueField is a FLOAT field
type StrikeValueField struct{ quickfix.FIXFloat }

//Tag returns tag.StrikeValue (968)
func (f StrikeValueField) Tag() quickfix.Tag { return tag.StrikeValue }

//NewStrikeValue returns a new StrikeValueField initialized with val
func NewStrikeValue(val quickfix.FIXFloat) StrikeValueField {
	return StrikeValueField{val}
}

//SubjectField is a STRING field
type SubjectField struct{ quickfix.FIXString }

//Tag returns tag.Subject (147)
func (f SubjectField) Tag() quickfix.Tag { return tag.Subject }

//NewSubject returns a new SubjectField initialized with val
func NewSubject(val quickfix.FIXString) SubjectField {
	return SubjectField{val}
}

//SubscriptionRequestTypeField is a CHAR field
type SubscriptionRequestTypeField struct{ quickfix.FIXString }

//Tag returns tag.SubscriptionRequestType (263)
func (f SubscriptionRequestTypeField) Tag() quickfix.Tag { return tag.SubscriptionRequestType }

//NewSubscriptionRequestType returns a new SubscriptionRequestTypeField initialized with val
func NewSubscriptionRequestType(val quickfix.FIXString) SubscriptionRequestTypeField {
	return SubscriptionRequestTypeField{val}
}

//SwapPointsField is a PRICEOFFSET field
type SwapPointsField struct{ quickfix.FIXFloat }

//Tag returns tag.SwapPoints (1069)
func (f SwapPointsField) Tag() quickfix.Tag { return tag.SwapPoints }

//NewSwapPoints returns a new SwapPointsField initialized with val
func NewSwapPoints(val quickfix.FIXFloat) SwapPointsField {
	return SwapPointsField{val}
}

//SymbolField is a STRING field
type SymbolField struct{ quickfix.FIXString }

//Tag returns tag.Symbol (55)
func (f SymbolField) Tag() quickfix.Tag { return tag.Symbol }

//NewSymbol returns a new SymbolField initialized with val
func NewSymbol(val quickfix.FIXString) SymbolField {
	return SymbolField{val}
}

//SymbolSfxField is a STRING field
type SymbolSfxField struct{ quickfix.FIXString }

//Tag returns tag.SymbolSfx (65)
func (f SymbolSfxField) Tag() quickfix.Tag { return tag.SymbolSfx }

//NewSymbolSfx returns a new SymbolSfxField initialized with val
func NewSymbolSfx(val quickfix.FIXString) SymbolSfxField {
	return SymbolSfxField{val}
}

//TZTransactTimeField is a TZTIMESTAMP field
type TZTransactTimeField struct{ quickfix.FIXString }

//Tag returns tag.TZTransactTime (1132)
func (f TZTransactTimeField) Tag() quickfix.Tag { return tag.TZTransactTime }

//NewTZTransactTime returns a new TZTransactTimeField initialized with val
func NewTZTransactTime(val quickfix.FIXString) TZTransactTimeField {
	return TZTransactTimeField{val}
}

//TargetCompIDField is a STRING field
type TargetCompIDField struct{ quickfix.FIXString }

//Tag returns tag.TargetCompID (56)
func (f TargetCompIDField) Tag() quickfix.Tag { return tag.TargetCompID }

//NewTargetCompID returns a new TargetCompIDField initialized with val
func NewTargetCompID(val quickfix.FIXString) TargetCompIDField {
	return TargetCompIDField{val}
}

//TargetLocationIDField is a STRING field
type TargetLocationIDField struct{ quickfix.FIXString }

//Tag returns tag.TargetLocationID (143)
func (f TargetLocationIDField) Tag() quickfix.Tag { return tag.TargetLocationID }

//NewTargetLocationID returns a new TargetLocationIDField initialized with val
func NewTargetLocationID(val quickfix.FIXString) TargetLocationIDField {
	return TargetLocationIDField{val}
}

//TargetPartyIDField is a STRING field
type TargetPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.TargetPartyID (1462)
func (f TargetPartyIDField) Tag() quickfix.Tag { return tag.TargetPartyID }

//NewTargetPartyID returns a new TargetPartyIDField initialized with val
func NewTargetPartyID(val quickfix.FIXString) TargetPartyIDField {
	return TargetPartyIDField{val}
}

//TargetPartyIDSourceField is a CHAR field
type TargetPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.TargetPartyIDSource (1463)
func (f TargetPartyIDSourceField) Tag() quickfix.Tag { return tag.TargetPartyIDSource }

//NewTargetPartyIDSource returns a new TargetPartyIDSourceField initialized with val
func NewTargetPartyIDSource(val quickfix.FIXString) TargetPartyIDSourceField {
	return TargetPartyIDSourceField{val}
}

//TargetPartyRoleField is a INT field
type TargetPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.TargetPartyRole (1464)
func (f TargetPartyRoleField) Tag() quickfix.Tag { return tag.TargetPartyRole }

//NewTargetPartyRole returns a new TargetPartyRoleField initialized with val
func NewTargetPartyRole(val quickfix.FIXInt) TargetPartyRoleField {
	return TargetPartyRoleField{val}
}

//TargetStrategyField is a INT field
type TargetStrategyField struct{ quickfix.FIXInt }

//Tag returns tag.TargetStrategy (847)
func (f TargetStrategyField) Tag() quickfix.Tag { return tag.TargetStrategy }

//NewTargetStrategy returns a new TargetStrategyField initialized with val
func NewTargetStrategy(val quickfix.FIXInt) TargetStrategyField {
	return TargetStrategyField{val}
}

//TargetStrategyParametersField is a STRING field
type TargetStrategyParametersField struct{ quickfix.FIXString }

//Tag returns tag.TargetStrategyParameters (848)
func (f TargetStrategyParametersField) Tag() quickfix.Tag { return tag.TargetStrategyParameters }

//NewTargetStrategyParameters returns a new TargetStrategyParametersField initialized with val
func NewTargetStrategyParameters(val quickfix.FIXString) TargetStrategyParametersField {
	return TargetStrategyParametersField{val}
}

//TargetStrategyPerformanceField is a FLOAT field
type TargetStrategyPerformanceField struct{ quickfix.FIXFloat }

//Tag returns tag.TargetStrategyPerformance (850)
func (f TargetStrategyPerformanceField) Tag() quickfix.Tag { return tag.TargetStrategyPerformance }

//NewTargetStrategyPerformance returns a new TargetStrategyPerformanceField initialized with val
func NewTargetStrategyPerformance(val quickfix.FIXFloat) TargetStrategyPerformanceField {
	return TargetStrategyPerformanceField{val}
}

//TargetSubIDField is a STRING field
type TargetSubIDField struct{ quickfix.FIXString }

//Tag returns tag.TargetSubID (57)
func (f TargetSubIDField) Tag() quickfix.Tag { return tag.TargetSubID }

//NewTargetSubID returns a new TargetSubIDField initialized with val
func NewTargetSubID(val quickfix.FIXString) TargetSubIDField {
	return TargetSubIDField{val}
}

//TaxAdvantageTypeField is a INT field
type TaxAdvantageTypeField struct{ quickfix.FIXInt }

//Tag returns tag.TaxAdvantageType (495)
func (f TaxAdvantageTypeField) Tag() quickfix.Tag { return tag.TaxAdvantageType }

//NewTaxAdvantageType returns a new TaxAdvantageTypeField initialized with val
func NewTaxAdvantageType(val quickfix.FIXInt) TaxAdvantageTypeField {
	return TaxAdvantageTypeField{val}
}

//TerminationTypeField is a INT field
type TerminationTypeField struct{ quickfix.FIXInt }

//Tag returns tag.TerminationType (788)
func (f TerminationTypeField) Tag() quickfix.Tag { return tag.TerminationType }

//NewTerminationType returns a new TerminationTypeField initialized with val
func NewTerminationType(val quickfix.FIXInt) TerminationTypeField {
	return TerminationTypeField{val}
}

//TestMessageIndicatorField is a BOOLEAN field
type TestMessageIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.TestMessageIndicator (464)
func (f TestMessageIndicatorField) Tag() quickfix.Tag { return tag.TestMessageIndicator }

//NewTestMessageIndicator returns a new TestMessageIndicatorField initialized with val
func NewTestMessageIndicator(val quickfix.FIXBoolean) TestMessageIndicatorField {
	return TestMessageIndicatorField{val}
}

//TestReqIDField is a STRING field
type TestReqIDField struct{ quickfix.FIXString }

//Tag returns tag.TestReqID (112)
func (f TestReqIDField) Tag() quickfix.Tag { return tag.TestReqID }

//NewTestReqID returns a new TestReqIDField initialized with val
func NewTestReqID(val quickfix.FIXString) TestReqIDField {
	return TestReqIDField{val}
}

//TextField is a STRING field
type TextField struct{ quickfix.FIXString }

//Tag returns tag.Text (58)
func (f TextField) Tag() quickfix.Tag { return tag.Text }

//NewText returns a new TextField initialized with val
func NewText(val quickfix.FIXString) TextField {
	return TextField{val}
}

//ThresholdAmountField is a PRICEOFFSET field
type ThresholdAmountField struct{ quickfix.FIXFloat }

//Tag returns tag.ThresholdAmount (834)
func (f ThresholdAmountField) Tag() quickfix.Tag { return tag.ThresholdAmount }

//NewThresholdAmount returns a new ThresholdAmountField initialized with val
func NewThresholdAmount(val quickfix.FIXFloat) ThresholdAmountField {
	return ThresholdAmountField{val}
}

//TickDirectionField is a CHAR field
type TickDirectionField struct{ quickfix.FIXString }

//Tag returns tag.TickDirection (274)
func (f TickDirectionField) Tag() quickfix.Tag { return tag.TickDirection }

//NewTickDirection returns a new TickDirectionField initialized with val
func NewTickDirection(val quickfix.FIXString) TickDirectionField {
	return TickDirectionField{val}
}

//TickIncrementField is a PRICE field
type TickIncrementField struct{ quickfix.FIXFloat }

//Tag returns tag.TickIncrement (1208)
func (f TickIncrementField) Tag() quickfix.Tag { return tag.TickIncrement }

//NewTickIncrement returns a new TickIncrementField initialized with val
func NewTickIncrement(val quickfix.FIXFloat) TickIncrementField {
	return TickIncrementField{val}
}

//TickRuleTypeField is a INT field
type TickRuleTypeField struct{ quickfix.FIXInt }

//Tag returns tag.TickRuleType (1209)
func (f TickRuleTypeField) Tag() quickfix.Tag { return tag.TickRuleType }

//NewTickRuleType returns a new TickRuleTypeField initialized with val
func NewTickRuleType(val quickfix.FIXInt) TickRuleTypeField {
	return TickRuleTypeField{val}
}

//TierCodeField is a STRING field
type TierCodeField struct{ quickfix.FIXString }

//Tag returns tag.TierCode (994)
func (f TierCodeField) Tag() quickfix.Tag { return tag.TierCode }

//NewTierCode returns a new TierCodeField initialized with val
func NewTierCode(val quickfix.FIXString) TierCodeField {
	return TierCodeField{val}
}

//TimeBracketField is a STRING field
type TimeBracketField struct{ quickfix.FIXString }

//Tag returns tag.TimeBracket (943)
func (f TimeBracketField) Tag() quickfix.Tag { return tag.TimeBracket }

//NewTimeBracket returns a new TimeBracketField initialized with val
func NewTimeBracket(val quickfix.FIXString) TimeBracketField {
	return TimeBracketField{val}
}

//TimeInForceField is a CHAR field
type TimeInForceField struct{ quickfix.FIXString }

//Tag returns tag.TimeInForce (59)
func (f TimeInForceField) Tag() quickfix.Tag { return tag.TimeInForce }

//NewTimeInForce returns a new TimeInForceField initialized with val
func NewTimeInForce(val quickfix.FIXString) TimeInForceField {
	return TimeInForceField{val}
}

//TimeToExpirationField is a FLOAT field
type TimeToExpirationField struct{ quickfix.FIXFloat }

//Tag returns tag.TimeToExpiration (1189)
func (f TimeToExpirationField) Tag() quickfix.Tag { return tag.TimeToExpiration }

//NewTimeToExpiration returns a new TimeToExpirationField initialized with val
func NewTimeToExpiration(val quickfix.FIXFloat) TimeToExpirationField {
	return TimeToExpirationField{val}
}

//TimeUnitField is a STRING field
type TimeUnitField struct{ quickfix.FIXString }

//Tag returns tag.TimeUnit (997)
func (f TimeUnitField) Tag() quickfix.Tag { return tag.TimeUnit }

//NewTimeUnit returns a new TimeUnitField initialized with val
func NewTimeUnit(val quickfix.FIXString) TimeUnitField {
	return TimeUnitField{val}
}

//TotNoAccQuotesField is a INT field
type TotNoAccQuotesField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoAccQuotes (1169)
func (f TotNoAccQuotesField) Tag() quickfix.Tag { return tag.TotNoAccQuotes }

//NewTotNoAccQuotes returns a new TotNoAccQuotesField initialized with val
func NewTotNoAccQuotes(val quickfix.FIXInt) TotNoAccQuotesField {
	return TotNoAccQuotesField{val}
}

//TotNoAllocsField is a INT field
type TotNoAllocsField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoAllocs (892)
func (f TotNoAllocsField) Tag() quickfix.Tag { return tag.TotNoAllocs }

//NewTotNoAllocs returns a new TotNoAllocsField initialized with val
func NewTotNoAllocs(val quickfix.FIXInt) TotNoAllocsField {
	return TotNoAllocsField{val}
}

//TotNoCxldQuotesField is a INT field
type TotNoCxldQuotesField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoCxldQuotes (1168)
func (f TotNoCxldQuotesField) Tag() quickfix.Tag { return tag.TotNoCxldQuotes }

//NewTotNoCxldQuotes returns a new TotNoCxldQuotesField initialized with val
func NewTotNoCxldQuotes(val quickfix.FIXInt) TotNoCxldQuotesField {
	return TotNoCxldQuotesField{val}
}

//TotNoFillsField is a INT field
type TotNoFillsField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoFills (1361)
func (f TotNoFillsField) Tag() quickfix.Tag { return tag.TotNoFills }

//NewTotNoFills returns a new TotNoFillsField initialized with val
func NewTotNoFills(val quickfix.FIXInt) TotNoFillsField {
	return TotNoFillsField{val}
}

//TotNoOrdersField is a INT field
type TotNoOrdersField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoOrders (68)
func (f TotNoOrdersField) Tag() quickfix.Tag { return tag.TotNoOrders }

//NewTotNoOrders returns a new TotNoOrdersField initialized with val
func NewTotNoOrders(val quickfix.FIXInt) TotNoOrdersField {
	return TotNoOrdersField{val}
}

//TotNoPartyListField is a INT field
type TotNoPartyListField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoPartyList (1512)
func (f TotNoPartyListField) Tag() quickfix.Tag { return tag.TotNoPartyList }

//NewTotNoPartyList returns a new TotNoPartyListField initialized with val
func NewTotNoPartyList(val quickfix.FIXInt) TotNoPartyListField {
	return TotNoPartyListField{val}
}

//TotNoQuoteEntriesField is a INT field
type TotNoQuoteEntriesField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoQuoteEntries (304)
func (f TotNoQuoteEntriesField) Tag() quickfix.Tag { return tag.TotNoQuoteEntries }

//NewTotNoQuoteEntries returns a new TotNoQuoteEntriesField initialized with val
func NewTotNoQuoteEntries(val quickfix.FIXInt) TotNoQuoteEntriesField {
	return TotNoQuoteEntriesField{val}
}

//TotNoRejQuotesField is a INT field
type TotNoRejQuotesField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoRejQuotes (1170)
func (f TotNoRejQuotesField) Tag() quickfix.Tag { return tag.TotNoRejQuotes }

//NewTotNoRejQuotes returns a new TotNoRejQuotesField initialized with val
func NewTotNoRejQuotes(val quickfix.FIXInt) TotNoRejQuotesField {
	return TotNoRejQuotesField{val}
}

//TotNoRelatedSymField is a INT field
type TotNoRelatedSymField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoRelatedSym (393)
func (f TotNoRelatedSymField) Tag() quickfix.Tag { return tag.TotNoRelatedSym }

//NewTotNoRelatedSym returns a new TotNoRelatedSymField initialized with val
func NewTotNoRelatedSym(val quickfix.FIXInt) TotNoRelatedSymField {
	return TotNoRelatedSymField{val}
}

//TotNoSecurityTypesField is a INT field
type TotNoSecurityTypesField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoSecurityTypes (557)
func (f TotNoSecurityTypesField) Tag() quickfix.Tag { return tag.TotNoSecurityTypes }

//NewTotNoSecurityTypes returns a new TotNoSecurityTypesField initialized with val
func NewTotNoSecurityTypes(val quickfix.FIXInt) TotNoSecurityTypesField {
	return TotNoSecurityTypesField{val}
}

//TotNoStrikesField is a INT field
type TotNoStrikesField struct{ quickfix.FIXInt }

//Tag returns tag.TotNoStrikes (422)
func (f TotNoStrikesField) Tag() quickfix.Tag { return tag.TotNoStrikes }

//NewTotNoStrikes returns a new TotNoStrikesField initialized with val
func NewTotNoStrikes(val quickfix.FIXInt) TotNoStrikesField {
	return TotNoStrikesField{val}
}

//TotNumAssignmentReportsField is a INT field
type TotNumAssignmentReportsField struct{ quickfix.FIXInt }

//Tag returns tag.TotNumAssignmentReports (832)
func (f TotNumAssignmentReportsField) Tag() quickfix.Tag { return tag.TotNumAssignmentReports }

//NewTotNumAssignmentReports returns a new TotNumAssignmentReportsField initialized with val
func NewTotNumAssignmentReports(val quickfix.FIXInt) TotNumAssignmentReportsField {
	return TotNumAssignmentReportsField{val}
}

//TotNumReportsField is a INT field
type TotNumReportsField struct{ quickfix.FIXInt }

//Tag returns tag.TotNumReports (911)
func (f TotNumReportsField) Tag() quickfix.Tag { return tag.TotNumReports }

//NewTotNumReports returns a new TotNumReportsField initialized with val
func NewTotNumReports(val quickfix.FIXInt) TotNumReportsField {
	return TotNumReportsField{val}
}

//TotNumTradeReportsField is a INT field
type TotNumTradeReportsField struct{ quickfix.FIXInt }

//Tag returns tag.TotNumTradeReports (748)
func (f TotNumTradeReportsField) Tag() quickfix.Tag { return tag.TotNumTradeReports }

//NewTotNumTradeReports returns a new TotNumTradeReportsField initialized with val
func NewTotNumTradeReports(val quickfix.FIXInt) TotNumTradeReportsField {
	return TotNumTradeReportsField{val}
}

//TotQuoteEntriesField is a INT field
type TotQuoteEntriesField struct{ quickfix.FIXInt }

//Tag returns tag.TotQuoteEntries (304)
func (f TotQuoteEntriesField) Tag() quickfix.Tag { return tag.TotQuoteEntries }

//NewTotQuoteEntries returns a new TotQuoteEntriesField initialized with val
func NewTotQuoteEntries(val quickfix.FIXInt) TotQuoteEntriesField {
	return TotQuoteEntriesField{val}
}

//TotalAccruedInterestAmtField is a AMT field
type TotalAccruedInterestAmtField struct{ quickfix.FIXFloat }

//Tag returns tag.TotalAccruedInterestAmt (540)
func (f TotalAccruedInterestAmtField) Tag() quickfix.Tag { return tag.TotalAccruedInterestAmt }

//NewTotalAccruedInterestAmt returns a new TotalAccruedInterestAmtField initialized with val
func NewTotalAccruedInterestAmt(val quickfix.FIXFloat) TotalAccruedInterestAmtField {
	return TotalAccruedInterestAmtField{val}
}

//TotalAffectedOrdersField is a INT field
type TotalAffectedOrdersField struct{ quickfix.FIXInt }

//Tag returns tag.TotalAffectedOrders (533)
func (f TotalAffectedOrdersField) Tag() quickfix.Tag { return tag.TotalAffectedOrders }

//NewTotalAffectedOrders returns a new TotalAffectedOrdersField initialized with val
func NewTotalAffectedOrders(val quickfix.FIXInt) TotalAffectedOrdersField {
	return TotalAffectedOrdersField{val}
}

//TotalNetValueField is a AMT field
type TotalNetValueField struct{ quickfix.FIXFloat }

//Tag returns tag.TotalNetValue (900)
func (f TotalNetValueField) Tag() quickfix.Tag { return tag.TotalNetValue }

//NewTotalNetValue returns a new TotalNetValueField initialized with val
func NewTotalNetValue(val quickfix.FIXFloat) TotalNetValueField {
	return TotalNetValueField{val}
}

//TotalNumPosReportsField is a INT field
type TotalNumPosReportsField struct{ quickfix.FIXInt }

//Tag returns tag.TotalNumPosReports (727)
func (f TotalNumPosReportsField) Tag() quickfix.Tag { return tag.TotalNumPosReports }

//NewTotalNumPosReports returns a new TotalNumPosReportsField initialized with val
func NewTotalNumPosReports(val quickfix.FIXInt) TotalNumPosReportsField {
	return TotalNumPosReportsField{val}
}

//TotalNumSecuritiesField is a INT field
type TotalNumSecuritiesField struct{ quickfix.FIXInt }

//Tag returns tag.TotalNumSecurities (393)
func (f TotalNumSecuritiesField) Tag() quickfix.Tag { return tag.TotalNumSecurities }

//NewTotalNumSecurities returns a new TotalNumSecuritiesField initialized with val
func NewTotalNumSecurities(val quickfix.FIXInt) TotalNumSecuritiesField {
	return TotalNumSecuritiesField{val}
}

//TotalNumSecurityTypesField is a INT field
type TotalNumSecurityTypesField struct{ quickfix.FIXInt }

//Tag returns tag.TotalNumSecurityTypes (557)
func (f TotalNumSecurityTypesField) Tag() quickfix.Tag { return tag.TotalNumSecurityTypes }

//NewTotalNumSecurityTypes returns a new TotalNumSecurityTypesField initialized with val
func NewTotalNumSecurityTypes(val quickfix.FIXInt) TotalNumSecurityTypesField {
	return TotalNumSecurityTypesField{val}
}

//TotalTakedownField is a AMT field
type TotalTakedownField struct{ quickfix.FIXFloat }

//Tag returns tag.TotalTakedown (237)
func (f TotalTakedownField) Tag() quickfix.Tag { return tag.TotalTakedown }

//NewTotalTakedown returns a new TotalTakedownField initialized with val
func NewTotalTakedown(val quickfix.FIXFloat) TotalTakedownField {
	return TotalTakedownField{val}
}

//TotalVolumeTradedField is a QTY field
type TotalVolumeTradedField struct{ quickfix.FIXFloat }

//Tag returns tag.TotalVolumeTraded (387)
func (f TotalVolumeTradedField) Tag() quickfix.Tag { return tag.TotalVolumeTraded }

//NewTotalVolumeTraded returns a new TotalVolumeTradedField initialized with val
func NewTotalVolumeTraded(val quickfix.FIXFloat) TotalVolumeTradedField {
	return TotalVolumeTradedField{val}
}

//TotalVolumeTradedDateField is a UTCDATEONLY field
type TotalVolumeTradedDateField struct{ quickfix.FIXString }

//Tag returns tag.TotalVolumeTradedDate (449)
func (f TotalVolumeTradedDateField) Tag() quickfix.Tag { return tag.TotalVolumeTradedDate }

//NewTotalVolumeTradedDate returns a new TotalVolumeTradedDateField initialized with val
func NewTotalVolumeTradedDate(val quickfix.FIXString) TotalVolumeTradedDateField {
	return TotalVolumeTradedDateField{val}
}

//TotalVolumeTradedTimeField is a UTCTIMEONLY field
type TotalVolumeTradedTimeField struct{ quickfix.FIXString }

//Tag returns tag.TotalVolumeTradedTime (450)
func (f TotalVolumeTradedTimeField) Tag() quickfix.Tag { return tag.TotalVolumeTradedTime }

//NewTotalVolumeTradedTime returns a new TotalVolumeTradedTimeField initialized with val
func NewTotalVolumeTradedTime(val quickfix.FIXString) TotalVolumeTradedTimeField {
	return TotalVolumeTradedTimeField{val}
}

//TradSesCloseTimeField is a UTCTIMESTAMP field
type TradSesCloseTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.TradSesCloseTime (344)
func (f TradSesCloseTimeField) Tag() quickfix.Tag { return tag.TradSesCloseTime }

//NewTradSesCloseTime returns a new TradSesCloseTimeField initialized with val
func NewTradSesCloseTime(val time.Time) TradSesCloseTimeField {
	return TradSesCloseTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewTradSesCloseTimeNoMillis returns a new TradSesCloseTimeField initialized with val without millisecs
func NewTradSesCloseTimeNoMillis(val time.Time) TradSesCloseTimeField {
	return TradSesCloseTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//TradSesEndTimeField is a UTCTIMESTAMP field
type TradSesEndTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.TradSesEndTime (345)
func (f TradSesEndTimeField) Tag() quickfix.Tag { return tag.TradSesEndTime }

//NewTradSesEndTime returns a new TradSesEndTimeField initialized with val
func NewTradSesEndTime(val time.Time) TradSesEndTimeField {
	return TradSesEndTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewTradSesEndTimeNoMillis returns a new TradSesEndTimeField initialized with val without millisecs
func NewTradSesEndTimeNoMillis(val time.Time) TradSesEndTimeField {
	return TradSesEndTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//TradSesEventField is a INT field
type TradSesEventField struct{ quickfix.FIXInt }

//Tag returns tag.TradSesEvent (1368)
func (f TradSesEventField) Tag() quickfix.Tag { return tag.TradSesEvent }

//NewTradSesEvent returns a new TradSesEventField initialized with val
func NewTradSesEvent(val quickfix.FIXInt) TradSesEventField {
	return TradSesEventField{val}
}

//TradSesMethodField is a INT field
type TradSesMethodField struct{ quickfix.FIXInt }

//Tag returns tag.TradSesMethod (338)
func (f TradSesMethodField) Tag() quickfix.Tag { return tag.TradSesMethod }

//NewTradSesMethod returns a new TradSesMethodField initialized with val
func NewTradSesMethod(val quickfix.FIXInt) TradSesMethodField {
	return TradSesMethodField{val}
}

//TradSesModeField is a INT field
type TradSesModeField struct{ quickfix.FIXInt }

//Tag returns tag.TradSesMode (339)
func (f TradSesModeField) Tag() quickfix.Tag { return tag.TradSesMode }

//NewTradSesMode returns a new TradSesModeField initialized with val
func NewTradSesMode(val quickfix.FIXInt) TradSesModeField {
	return TradSesModeField{val}
}

//TradSesOpenTimeField is a UTCTIMESTAMP field
type TradSesOpenTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.TradSesOpenTime (342)
func (f TradSesOpenTimeField) Tag() quickfix.Tag { return tag.TradSesOpenTime }

//NewTradSesOpenTime returns a new TradSesOpenTimeField initialized with val
func NewTradSesOpenTime(val time.Time) TradSesOpenTimeField {
	return TradSesOpenTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewTradSesOpenTimeNoMillis returns a new TradSesOpenTimeField initialized with val without millisecs
func NewTradSesOpenTimeNoMillis(val time.Time) TradSesOpenTimeField {
	return TradSesOpenTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//TradSesPreCloseTimeField is a UTCTIMESTAMP field
type TradSesPreCloseTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.TradSesPreCloseTime (343)
func (f TradSesPreCloseTimeField) Tag() quickfix.Tag { return tag.TradSesPreCloseTime }

//NewTradSesPreCloseTime returns a new TradSesPreCloseTimeField initialized with val
func NewTradSesPreCloseTime(val time.Time) TradSesPreCloseTimeField {
	return TradSesPreCloseTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewTradSesPreCloseTimeNoMillis returns a new TradSesPreCloseTimeField initialized with val without millisecs
func NewTradSesPreCloseTimeNoMillis(val time.Time) TradSesPreCloseTimeField {
	return TradSesPreCloseTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//TradSesReqIDField is a STRING field
type TradSesReqIDField struct{ quickfix.FIXString }

//Tag returns tag.TradSesReqID (335)
func (f TradSesReqIDField) Tag() quickfix.Tag { return tag.TradSesReqID }

//NewTradSesReqID returns a new TradSesReqIDField initialized with val
func NewTradSesReqID(val quickfix.FIXString) TradSesReqIDField {
	return TradSesReqIDField{val}
}

//TradSesStartTimeField is a UTCTIMESTAMP field
type TradSesStartTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.TradSesStartTime (341)
func (f TradSesStartTimeField) Tag() quickfix.Tag { return tag.TradSesStartTime }

//NewTradSesStartTime returns a new TradSesStartTimeField initialized with val
func NewTradSesStartTime(val time.Time) TradSesStartTimeField {
	return TradSesStartTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewTradSesStartTimeNoMillis returns a new TradSesStartTimeField initialized with val without millisecs
func NewTradSesStartTimeNoMillis(val time.Time) TradSesStartTimeField {
	return TradSesStartTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//TradSesStatusField is a INT field
type TradSesStatusField struct{ quickfix.FIXInt }

//Tag returns tag.TradSesStatus (340)
func (f TradSesStatusField) Tag() quickfix.Tag { return tag.TradSesStatus }

//NewTradSesStatus returns a new TradSesStatusField initialized with val
func NewTradSesStatus(val quickfix.FIXInt) TradSesStatusField {
	return TradSesStatusField{val}
}

//TradSesStatusRejReasonField is a INT field
type TradSesStatusRejReasonField struct{ quickfix.FIXInt }

//Tag returns tag.TradSesStatusRejReason (567)
func (f TradSesStatusRejReasonField) Tag() quickfix.Tag { return tag.TradSesStatusRejReason }

//NewTradSesStatusRejReason returns a new TradSesStatusRejReasonField initialized with val
func NewTradSesStatusRejReason(val quickfix.FIXInt) TradSesStatusRejReasonField {
	return TradSesStatusRejReasonField{val}
}

//TradSesUpdateActionField is a CHAR field
type TradSesUpdateActionField struct{ quickfix.FIXString }

//Tag returns tag.TradSesUpdateAction (1327)
func (f TradSesUpdateActionField) Tag() quickfix.Tag { return tag.TradSesUpdateAction }

//NewTradSesUpdateAction returns a new TradSesUpdateActionField initialized with val
func NewTradSesUpdateAction(val quickfix.FIXString) TradSesUpdateActionField {
	return TradSesUpdateActionField{val}
}

//TradeAllocIndicatorField is a INT field
type TradeAllocIndicatorField struct{ quickfix.FIXInt }

//Tag returns tag.TradeAllocIndicator (826)
func (f TradeAllocIndicatorField) Tag() quickfix.Tag { return tag.TradeAllocIndicator }

//NewTradeAllocIndicator returns a new TradeAllocIndicatorField initialized with val
func NewTradeAllocIndicator(val quickfix.FIXInt) TradeAllocIndicatorField {
	return TradeAllocIndicatorField{val}
}

//TradeConditionField is a MULTIPLESTRINGVALUE field
type TradeConditionField struct{ quickfix.FIXString }

//Tag returns tag.TradeCondition (277)
func (f TradeConditionField) Tag() quickfix.Tag { return tag.TradeCondition }

//NewTradeCondition returns a new TradeConditionField initialized with val
func NewTradeCondition(val quickfix.FIXString) TradeConditionField {
	return TradeConditionField{val}
}

//TradeDateField is a LOCALMKTDATE field
type TradeDateField struct{ quickfix.FIXString }

//Tag returns tag.TradeDate (75)
func (f TradeDateField) Tag() quickfix.Tag { return tag.TradeDate }

//NewTradeDate returns a new TradeDateField initialized with val
func NewTradeDate(val quickfix.FIXString) TradeDateField {
	return TradeDateField{val}
}

//TradeHandlingInstrField is a CHAR field
type TradeHandlingInstrField struct{ quickfix.FIXString }

//Tag returns tag.TradeHandlingInstr (1123)
func (f TradeHandlingInstrField) Tag() quickfix.Tag { return tag.TradeHandlingInstr }

//NewTradeHandlingInstr returns a new TradeHandlingInstrField initialized with val
func NewTradeHandlingInstr(val quickfix.FIXString) TradeHandlingInstrField {
	return TradeHandlingInstrField{val}
}

//TradeIDField is a STRING field
type TradeIDField struct{ quickfix.FIXString }

//Tag returns tag.TradeID (1003)
func (f TradeIDField) Tag() quickfix.Tag { return tag.TradeID }

//NewTradeID returns a new TradeIDField initialized with val
func NewTradeID(val quickfix.FIXString) TradeIDField {
	return TradeIDField{val}
}

//TradeInputDeviceField is a STRING field
type TradeInputDeviceField struct{ quickfix.FIXString }

//Tag returns tag.TradeInputDevice (579)
func (f TradeInputDeviceField) Tag() quickfix.Tag { return tag.TradeInputDevice }

//NewTradeInputDevice returns a new TradeInputDeviceField initialized with val
func NewTradeInputDevice(val quickfix.FIXString) TradeInputDeviceField {
	return TradeInputDeviceField{val}
}

//TradeInputSourceField is a STRING field
type TradeInputSourceField struct{ quickfix.FIXString }

//Tag returns tag.TradeInputSource (578)
func (f TradeInputSourceField) Tag() quickfix.Tag { return tag.TradeInputSource }

//NewTradeInputSource returns a new TradeInputSourceField initialized with val
func NewTradeInputSource(val quickfix.FIXString) TradeInputSourceField {
	return TradeInputSourceField{val}
}

//TradeLegRefIDField is a STRING field
type TradeLegRefIDField struct{ quickfix.FIXString }

//Tag returns tag.TradeLegRefID (824)
func (f TradeLegRefIDField) Tag() quickfix.Tag { return tag.TradeLegRefID }

//NewTradeLegRefID returns a new TradeLegRefIDField initialized with val
func NewTradeLegRefID(val quickfix.FIXString) TradeLegRefIDField {
	return TradeLegRefIDField{val}
}

//TradeLinkIDField is a STRING field
type TradeLinkIDField struct{ quickfix.FIXString }

//Tag returns tag.TradeLinkID (820)
func (f TradeLinkIDField) Tag() quickfix.Tag { return tag.TradeLinkID }

//NewTradeLinkID returns a new TradeLinkIDField initialized with val
func NewTradeLinkID(val quickfix.FIXString) TradeLinkIDField {
	return TradeLinkIDField{val}
}

//TradeOriginationDateField is a LOCALMKTDATE field
type TradeOriginationDateField struct{ quickfix.FIXString }

//Tag returns tag.TradeOriginationDate (229)
func (f TradeOriginationDateField) Tag() quickfix.Tag { return tag.TradeOriginationDate }

//NewTradeOriginationDate returns a new TradeOriginationDateField initialized with val
func NewTradeOriginationDate(val quickfix.FIXString) TradeOriginationDateField {
	return TradeOriginationDateField{val}
}

//TradePublishIndicatorField is a INT field
type TradePublishIndicatorField struct{ quickfix.FIXInt }

//Tag returns tag.TradePublishIndicator (1390)
func (f TradePublishIndicatorField) Tag() quickfix.Tag { return tag.TradePublishIndicator }

//NewTradePublishIndicator returns a new TradePublishIndicatorField initialized with val
func NewTradePublishIndicator(val quickfix.FIXInt) TradePublishIndicatorField {
	return TradePublishIndicatorField{val}
}

//TradeReportIDField is a STRING field
type TradeReportIDField struct{ quickfix.FIXString }

//Tag returns tag.TradeReportID (571)
func (f TradeReportIDField) Tag() quickfix.Tag { return tag.TradeReportID }

//NewTradeReportID returns a new TradeReportIDField initialized with val
func NewTradeReportID(val quickfix.FIXString) TradeReportIDField {
	return TradeReportIDField{val}
}

//TradeReportRefIDField is a STRING field
type TradeReportRefIDField struct{ quickfix.FIXString }

//Tag returns tag.TradeReportRefID (572)
func (f TradeReportRefIDField) Tag() quickfix.Tag { return tag.TradeReportRefID }

//NewTradeReportRefID returns a new TradeReportRefIDField initialized with val
func NewTradeReportRefID(val quickfix.FIXString) TradeReportRefIDField {
	return TradeReportRefIDField{val}
}

//TradeReportRejectReasonField is a INT field
type TradeReportRejectReasonField struct{ quickfix.FIXInt }

//Tag returns tag.TradeReportRejectReason (751)
func (f TradeReportRejectReasonField) Tag() quickfix.Tag { return tag.TradeReportRejectReason }

//NewTradeReportRejectReason returns a new TradeReportRejectReasonField initialized with val
func NewTradeReportRejectReason(val quickfix.FIXInt) TradeReportRejectReasonField {
	return TradeReportRejectReasonField{val}
}

//TradeReportTransTypeField is a INT field
type TradeReportTransTypeField struct{ quickfix.FIXInt }

//Tag returns tag.TradeReportTransType (487)
func (f TradeReportTransTypeField) Tag() quickfix.Tag { return tag.TradeReportTransType }

//NewTradeReportTransType returns a new TradeReportTransTypeField initialized with val
func NewTradeReportTransType(val quickfix.FIXInt) TradeReportTransTypeField {
	return TradeReportTransTypeField{val}
}

//TradeReportTypeField is a INT field
type TradeReportTypeField struct{ quickfix.FIXInt }

//Tag returns tag.TradeReportType (856)
func (f TradeReportTypeField) Tag() quickfix.Tag { return tag.TradeReportType }

//NewTradeReportType returns a new TradeReportTypeField initialized with val
func NewTradeReportType(val quickfix.FIXInt) TradeReportTypeField {
	return TradeReportTypeField{val}
}

//TradeRequestIDField is a STRING field
type TradeRequestIDField struct{ quickfix.FIXString }

//Tag returns tag.TradeRequestID (568)
func (f TradeRequestIDField) Tag() quickfix.Tag { return tag.TradeRequestID }

//NewTradeRequestID returns a new TradeRequestIDField initialized with val
func NewTradeRequestID(val quickfix.FIXString) TradeRequestIDField {
	return TradeRequestIDField{val}
}

//TradeRequestResultField is a INT field
type TradeRequestResultField struct{ quickfix.FIXInt }

//Tag returns tag.TradeRequestResult (749)
func (f TradeRequestResultField) Tag() quickfix.Tag { return tag.TradeRequestResult }

//NewTradeRequestResult returns a new TradeRequestResultField initialized with val
func NewTradeRequestResult(val quickfix.FIXInt) TradeRequestResultField {
	return TradeRequestResultField{val}
}

//TradeRequestStatusField is a INT field
type TradeRequestStatusField struct{ quickfix.FIXInt }

//Tag returns tag.TradeRequestStatus (750)
func (f TradeRequestStatusField) Tag() quickfix.Tag { return tag.TradeRequestStatus }

//NewTradeRequestStatus returns a new TradeRequestStatusField initialized with val
func NewTradeRequestStatus(val quickfix.FIXInt) TradeRequestStatusField {
	return TradeRequestStatusField{val}
}

//TradeRequestTypeField is a INT field
type TradeRequestTypeField struct{ quickfix.FIXInt }

//Tag returns tag.TradeRequestType (569)
func (f TradeRequestTypeField) Tag() quickfix.Tag { return tag.TradeRequestType }

//NewTradeRequestType returns a new TradeRequestTypeField initialized with val
func NewTradeRequestType(val quickfix.FIXInt) TradeRequestTypeField {
	return TradeRequestTypeField{val}
}

//TradeTypeField is a CHAR field
type TradeTypeField struct{ quickfix.FIXString }

//Tag returns tag.TradeType (418)
func (f TradeTypeField) Tag() quickfix.Tag { return tag.TradeType }

//NewTradeType returns a new TradeTypeField initialized with val
func NewTradeType(val quickfix.FIXString) TradeTypeField {
	return TradeTypeField{val}
}

//TradeVolumeField is a QTY field
type TradeVolumeField struct{ quickfix.FIXFloat }

//Tag returns tag.TradeVolume (1020)
func (f TradeVolumeField) Tag() quickfix.Tag { return tag.TradeVolume }

//NewTradeVolume returns a new TradeVolumeField initialized with val
func NewTradeVolume(val quickfix.FIXFloat) TradeVolumeField {
	return TradeVolumeField{val}
}

//TradedFlatSwitchField is a BOOLEAN field
type TradedFlatSwitchField struct{ quickfix.FIXBoolean }

//Tag returns tag.TradedFlatSwitch (258)
func (f TradedFlatSwitchField) Tag() quickfix.Tag { return tag.TradedFlatSwitch }

//NewTradedFlatSwitch returns a new TradedFlatSwitchField initialized with val
func NewTradedFlatSwitch(val quickfix.FIXBoolean) TradedFlatSwitchField {
	return TradedFlatSwitchField{val}
}

//TradingCurrencyField is a CURRENCY field
type TradingCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.TradingCurrency (1245)
func (f TradingCurrencyField) Tag() quickfix.Tag { return tag.TradingCurrency }

//NewTradingCurrency returns a new TradingCurrencyField initialized with val
func NewTradingCurrency(val quickfix.FIXString) TradingCurrencyField {
	return TradingCurrencyField{val}
}

//TradingReferencePriceField is a PRICE field
type TradingReferencePriceField struct{ quickfix.FIXFloat }

//Tag returns tag.TradingReferencePrice (1150)
func (f TradingReferencePriceField) Tag() quickfix.Tag { return tag.TradingReferencePrice }

//NewTradingReferencePrice returns a new TradingReferencePriceField initialized with val
func NewTradingReferencePrice(val quickfix.FIXFloat) TradingReferencePriceField {
	return TradingReferencePriceField{val}
}

//TradingSessionDescField is a STRING field
type TradingSessionDescField struct{ quickfix.FIXString }

//Tag returns tag.TradingSessionDesc (1326)
func (f TradingSessionDescField) Tag() quickfix.Tag { return tag.TradingSessionDesc }

//NewTradingSessionDesc returns a new TradingSessionDescField initialized with val
func NewTradingSessionDesc(val quickfix.FIXString) TradingSessionDescField {
	return TradingSessionDescField{val}
}

//TradingSessionIDField is a STRING field
type TradingSessionIDField struct{ quickfix.FIXString }

//Tag returns tag.TradingSessionID (336)
func (f TradingSessionIDField) Tag() quickfix.Tag { return tag.TradingSessionID }

//NewTradingSessionID returns a new TradingSessionIDField initialized with val
func NewTradingSessionID(val quickfix.FIXString) TradingSessionIDField {
	return TradingSessionIDField{val}
}

//TradingSessionSubIDField is a STRING field
type TradingSessionSubIDField struct{ quickfix.FIXString }

//Tag returns tag.TradingSessionSubID (625)
func (f TradingSessionSubIDField) Tag() quickfix.Tag { return tag.TradingSessionSubID }

//NewTradingSessionSubID returns a new TradingSessionSubIDField initialized with val
func NewTradingSessionSubID(val quickfix.FIXString) TradingSessionSubIDField {
	return TradingSessionSubIDField{val}
}

//TransBkdTimeField is a UTCTIMESTAMP field
type TransBkdTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.TransBkdTime (483)
func (f TransBkdTimeField) Tag() quickfix.Tag { return tag.TransBkdTime }

//NewTransBkdTime returns a new TransBkdTimeField initialized with val
func NewTransBkdTime(val time.Time) TransBkdTimeField {
	return TransBkdTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewTransBkdTimeNoMillis returns a new TransBkdTimeField initialized with val without millisecs
func NewTransBkdTimeNoMillis(val time.Time) TransBkdTimeField {
	return TransBkdTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//TransactTimeField is a UTCTIMESTAMP field
type TransactTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.TransactTime (60)
func (f TransactTimeField) Tag() quickfix.Tag { return tag.TransactTime }

//NewTransactTime returns a new TransactTimeField initialized with val
func NewTransactTime(val time.Time) TransactTimeField {
	return TransactTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewTransactTimeNoMillis returns a new TransactTimeField initialized with val without millisecs
func NewTransactTimeNoMillis(val time.Time) TransactTimeField {
	return TransactTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//TransferReasonField is a STRING field
type TransferReasonField struct{ quickfix.FIXString }

//Tag returns tag.TransferReason (830)
func (f TransferReasonField) Tag() quickfix.Tag { return tag.TransferReason }

//NewTransferReason returns a new TransferReasonField initialized with val
func NewTransferReason(val quickfix.FIXString) TransferReasonField {
	return TransferReasonField{val}
}

//TrdMatchIDField is a STRING field
type TrdMatchIDField struct{ quickfix.FIXString }

//Tag returns tag.TrdMatchID (880)
func (f TrdMatchIDField) Tag() quickfix.Tag { return tag.TrdMatchID }

//NewTrdMatchID returns a new TrdMatchIDField initialized with val
func NewTrdMatchID(val quickfix.FIXString) TrdMatchIDField {
	return TrdMatchIDField{val}
}

//TrdRegTimestampField is a UTCTIMESTAMP field
type TrdRegTimestampField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.TrdRegTimestamp (769)
func (f TrdRegTimestampField) Tag() quickfix.Tag { return tag.TrdRegTimestamp }

//NewTrdRegTimestamp returns a new TrdRegTimestampField initialized with val
func NewTrdRegTimestamp(val time.Time) TrdRegTimestampField {
	return TrdRegTimestampField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewTrdRegTimestampNoMillis returns a new TrdRegTimestampField initialized with val without millisecs
func NewTrdRegTimestampNoMillis(val time.Time) TrdRegTimestampField {
	return TrdRegTimestampField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//TrdRegTimestampOriginField is a STRING field
type TrdRegTimestampOriginField struct{ quickfix.FIXString }

//Tag returns tag.TrdRegTimestampOrigin (771)
func (f TrdRegTimestampOriginField) Tag() quickfix.Tag { return tag.TrdRegTimestampOrigin }

//NewTrdRegTimestampOrigin returns a new TrdRegTimestampOriginField initialized with val
func NewTrdRegTimestampOrigin(val quickfix.FIXString) TrdRegTimestampOriginField {
	return TrdRegTimestampOriginField{val}
}

//TrdRegTimestampTypeField is a INT field
type TrdRegTimestampTypeField struct{ quickfix.FIXInt }

//Tag returns tag.TrdRegTimestampType (770)
func (f TrdRegTimestampTypeField) Tag() quickfix.Tag { return tag.TrdRegTimestampType }

//NewTrdRegTimestampType returns a new TrdRegTimestampTypeField initialized with val
func NewTrdRegTimestampType(val quickfix.FIXInt) TrdRegTimestampTypeField {
	return TrdRegTimestampTypeField{val}
}

//TrdRepIndicatorField is a BOOLEAN field
type TrdRepIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.TrdRepIndicator (1389)
func (f TrdRepIndicatorField) Tag() quickfix.Tag { return tag.TrdRepIndicator }

//NewTrdRepIndicator returns a new TrdRepIndicatorField initialized with val
func NewTrdRepIndicator(val quickfix.FIXBoolean) TrdRepIndicatorField {
	return TrdRepIndicatorField{val}
}

//TrdRepPartyRoleField is a INT field
type TrdRepPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.TrdRepPartyRole (1388)
func (f TrdRepPartyRoleField) Tag() quickfix.Tag { return tag.TrdRepPartyRole }

//NewTrdRepPartyRole returns a new TrdRepPartyRoleField initialized with val
func NewTrdRepPartyRole(val quickfix.FIXInt) TrdRepPartyRoleField {
	return TrdRepPartyRoleField{val}
}

//TrdRptStatusField is a INT field
type TrdRptStatusField struct{ quickfix.FIXInt }

//Tag returns tag.TrdRptStatus (939)
func (f TrdRptStatusField) Tag() quickfix.Tag { return tag.TrdRptStatus }

//NewTrdRptStatus returns a new TrdRptStatusField initialized with val
func NewTrdRptStatus(val quickfix.FIXInt) TrdRptStatusField {
	return TrdRptStatusField{val}
}

//TrdSubTypeField is a INT field
type TrdSubTypeField struct{ quickfix.FIXInt }

//Tag returns tag.TrdSubType (829)
func (f TrdSubTypeField) Tag() quickfix.Tag { return tag.TrdSubType }

//NewTrdSubType returns a new TrdSubTypeField initialized with val
func NewTrdSubType(val quickfix.FIXInt) TrdSubTypeField {
	return TrdSubTypeField{val}
}

//TrdTypeField is a INT field
type TrdTypeField struct{ quickfix.FIXInt }

//Tag returns tag.TrdType (828)
func (f TrdTypeField) Tag() quickfix.Tag { return tag.TrdType }

//NewTrdType returns a new TrdTypeField initialized with val
func NewTrdType(val quickfix.FIXInt) TrdTypeField {
	return TrdTypeField{val}
}

//TriggerActionField is a CHAR field
type TriggerActionField struct{ quickfix.FIXString }

//Tag returns tag.TriggerAction (1101)
func (f TriggerActionField) Tag() quickfix.Tag { return tag.TriggerAction }

//NewTriggerAction returns a new TriggerActionField initialized with val
func NewTriggerAction(val quickfix.FIXString) TriggerActionField {
	return TriggerActionField{val}
}

//TriggerNewPriceField is a PRICE field
type TriggerNewPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.TriggerNewPrice (1110)
func (f TriggerNewPriceField) Tag() quickfix.Tag { return tag.TriggerNewPrice }

//NewTriggerNewPrice returns a new TriggerNewPriceField initialized with val
func NewTriggerNewPrice(val quickfix.FIXFloat) TriggerNewPriceField {
	return TriggerNewPriceField{val}
}

//TriggerNewQtyField is a QTY field
type TriggerNewQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.TriggerNewQty (1112)
func (f TriggerNewQtyField) Tag() quickfix.Tag { return tag.TriggerNewQty }

//NewTriggerNewQty returns a new TriggerNewQtyField initialized with val
func NewTriggerNewQty(val quickfix.FIXFloat) TriggerNewQtyField {
	return TriggerNewQtyField{val}
}

//TriggerOrderTypeField is a CHAR field
type TriggerOrderTypeField struct{ quickfix.FIXString }

//Tag returns tag.TriggerOrderType (1111)
func (f TriggerOrderTypeField) Tag() quickfix.Tag { return tag.TriggerOrderType }

//NewTriggerOrderType returns a new TriggerOrderTypeField initialized with val
func NewTriggerOrderType(val quickfix.FIXString) TriggerOrderTypeField {
	return TriggerOrderTypeField{val}
}

//TriggerPriceField is a PRICE field
type TriggerPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.TriggerPrice (1102)
func (f TriggerPriceField) Tag() quickfix.Tag { return tag.TriggerPrice }

//NewTriggerPrice returns a new TriggerPriceField initialized with val
func NewTriggerPrice(val quickfix.FIXFloat) TriggerPriceField {
	return TriggerPriceField{val}
}

//TriggerPriceDirectionField is a CHAR field
type TriggerPriceDirectionField struct{ quickfix.FIXString }

//Tag returns tag.TriggerPriceDirection (1109)
func (f TriggerPriceDirectionField) Tag() quickfix.Tag { return tag.TriggerPriceDirection }

//NewTriggerPriceDirection returns a new TriggerPriceDirectionField initialized with val
func NewTriggerPriceDirection(val quickfix.FIXString) TriggerPriceDirectionField {
	return TriggerPriceDirectionField{val}
}

//TriggerPriceTypeField is a CHAR field
type TriggerPriceTypeField struct{ quickfix.FIXString }

//Tag returns tag.TriggerPriceType (1107)
func (f TriggerPriceTypeField) Tag() quickfix.Tag { return tag.TriggerPriceType }

//NewTriggerPriceType returns a new TriggerPriceTypeField initialized with val
func NewTriggerPriceType(val quickfix.FIXString) TriggerPriceTypeField {
	return TriggerPriceTypeField{val}
}

//TriggerPriceTypeScopeField is a CHAR field
type TriggerPriceTypeScopeField struct{ quickfix.FIXString }

//Tag returns tag.TriggerPriceTypeScope (1108)
func (f TriggerPriceTypeScopeField) Tag() quickfix.Tag { return tag.TriggerPriceTypeScope }

//NewTriggerPriceTypeScope returns a new TriggerPriceTypeScopeField initialized with val
func NewTriggerPriceTypeScope(val quickfix.FIXString) TriggerPriceTypeScopeField {
	return TriggerPriceTypeScopeField{val}
}

//TriggerSecurityDescField is a STRING field
type TriggerSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.TriggerSecurityDesc (1106)
func (f TriggerSecurityDescField) Tag() quickfix.Tag { return tag.TriggerSecurityDesc }

//NewTriggerSecurityDesc returns a new TriggerSecurityDescField initialized with val
func NewTriggerSecurityDesc(val quickfix.FIXString) TriggerSecurityDescField {
	return TriggerSecurityDescField{val}
}

//TriggerSecurityIDField is a STRING field
type TriggerSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.TriggerSecurityID (1104)
func (f TriggerSecurityIDField) Tag() quickfix.Tag { return tag.TriggerSecurityID }

//NewTriggerSecurityID returns a new TriggerSecurityIDField initialized with val
func NewTriggerSecurityID(val quickfix.FIXString) TriggerSecurityIDField {
	return TriggerSecurityIDField{val}
}

//TriggerSecurityIDSourceField is a STRING field
type TriggerSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.TriggerSecurityIDSource (1105)
func (f TriggerSecurityIDSourceField) Tag() quickfix.Tag { return tag.TriggerSecurityIDSource }

//NewTriggerSecurityIDSource returns a new TriggerSecurityIDSourceField initialized with val
func NewTriggerSecurityIDSource(val quickfix.FIXString) TriggerSecurityIDSourceField {
	return TriggerSecurityIDSourceField{val}
}

//TriggerSymbolField is a STRING field
type TriggerSymbolField struct{ quickfix.FIXString }

//Tag returns tag.TriggerSymbol (1103)
func (f TriggerSymbolField) Tag() quickfix.Tag { return tag.TriggerSymbol }

//NewTriggerSymbol returns a new TriggerSymbolField initialized with val
func NewTriggerSymbol(val quickfix.FIXString) TriggerSymbolField {
	return TriggerSymbolField{val}
}

//TriggerTradingSessionIDField is a STRING field
type TriggerTradingSessionIDField struct{ quickfix.FIXString }

//Tag returns tag.TriggerTradingSessionID (1113)
func (f TriggerTradingSessionIDField) Tag() quickfix.Tag { return tag.TriggerTradingSessionID }

//NewTriggerTradingSessionID returns a new TriggerTradingSessionIDField initialized with val
func NewTriggerTradingSessionID(val quickfix.FIXString) TriggerTradingSessionIDField {
	return TriggerTradingSessionIDField{val}
}

//TriggerTradingSessionSubIDField is a STRING field
type TriggerTradingSessionSubIDField struct{ quickfix.FIXString }

//Tag returns tag.TriggerTradingSessionSubID (1114)
func (f TriggerTradingSessionSubIDField) Tag() quickfix.Tag { return tag.TriggerTradingSessionSubID }

//NewTriggerTradingSessionSubID returns a new TriggerTradingSessionSubIDField initialized with val
func NewTriggerTradingSessionSubID(val quickfix.FIXString) TriggerTradingSessionSubIDField {
	return TriggerTradingSessionSubIDField{val}
}

//TriggerTypeField is a CHAR field
type TriggerTypeField struct{ quickfix.FIXString }

//Tag returns tag.TriggerType (1100)
func (f TriggerTypeField) Tag() quickfix.Tag { return tag.TriggerType }

//NewTriggerType returns a new TriggerTypeField initialized with val
func NewTriggerType(val quickfix.FIXString) TriggerTypeField {
	return TriggerTypeField{val}
}

//URLLinkField is a STRING field
type URLLinkField struct{ quickfix.FIXString }

//Tag returns tag.URLLink (149)
func (f URLLinkField) Tag() quickfix.Tag { return tag.URLLink }

//NewURLLink returns a new URLLinkField initialized with val
func NewURLLink(val quickfix.FIXString) URLLinkField {
	return URLLinkField{val}
}

//UnderlyingAdjustedQuantityField is a QTY field
type UnderlyingAdjustedQuantityField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingAdjustedQuantity (1044)
func (f UnderlyingAdjustedQuantityField) Tag() quickfix.Tag { return tag.UnderlyingAdjustedQuantity }

//NewUnderlyingAdjustedQuantity returns a new UnderlyingAdjustedQuantityField initialized with val
func NewUnderlyingAdjustedQuantity(val quickfix.FIXFloat) UnderlyingAdjustedQuantityField {
	return UnderlyingAdjustedQuantityField{val}
}

//UnderlyingAllocationPercentField is a PERCENTAGE field
type UnderlyingAllocationPercentField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingAllocationPercent (972)
func (f UnderlyingAllocationPercentField) Tag() quickfix.Tag { return tag.UnderlyingAllocationPercent }

//NewUnderlyingAllocationPercent returns a new UnderlyingAllocationPercentField initialized with val
func NewUnderlyingAllocationPercent(val quickfix.FIXFloat) UnderlyingAllocationPercentField {
	return UnderlyingAllocationPercentField{val}
}

//UnderlyingAttachmentPointField is a PERCENTAGE field
type UnderlyingAttachmentPointField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingAttachmentPoint (1459)
func (f UnderlyingAttachmentPointField) Tag() quickfix.Tag { return tag.UnderlyingAttachmentPoint }

//NewUnderlyingAttachmentPoint returns a new UnderlyingAttachmentPointField initialized with val
func NewUnderlyingAttachmentPoint(val quickfix.FIXFloat) UnderlyingAttachmentPointField {
	return UnderlyingAttachmentPointField{val}
}

//UnderlyingCFICodeField is a STRING field
type UnderlyingCFICodeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingCFICode (463)
func (f UnderlyingCFICodeField) Tag() quickfix.Tag { return tag.UnderlyingCFICode }

//NewUnderlyingCFICode returns a new UnderlyingCFICodeField initialized with val
func NewUnderlyingCFICode(val quickfix.FIXString) UnderlyingCFICodeField {
	return UnderlyingCFICodeField{val}
}

//UnderlyingCPProgramField is a STRING field
type UnderlyingCPProgramField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingCPProgram (877)
func (f UnderlyingCPProgramField) Tag() quickfix.Tag { return tag.UnderlyingCPProgram }

//NewUnderlyingCPProgram returns a new UnderlyingCPProgramField initialized with val
func NewUnderlyingCPProgram(val quickfix.FIXString) UnderlyingCPProgramField {
	return UnderlyingCPProgramField{val}
}

//UnderlyingCPRegTypeField is a STRING field
type UnderlyingCPRegTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingCPRegType (878)
func (f UnderlyingCPRegTypeField) Tag() quickfix.Tag { return tag.UnderlyingCPRegType }

//NewUnderlyingCPRegType returns a new UnderlyingCPRegTypeField initialized with val
func NewUnderlyingCPRegType(val quickfix.FIXString) UnderlyingCPRegTypeField {
	return UnderlyingCPRegTypeField{val}
}

//UnderlyingCapValueField is a AMT field
type UnderlyingCapValueField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingCapValue (1038)
func (f UnderlyingCapValueField) Tag() quickfix.Tag { return tag.UnderlyingCapValue }

//NewUnderlyingCapValue returns a new UnderlyingCapValueField initialized with val
func NewUnderlyingCapValue(val quickfix.FIXFloat) UnderlyingCapValueField {
	return UnderlyingCapValueField{val}
}

//UnderlyingCashAmountField is a AMT field
type UnderlyingCashAmountField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingCashAmount (973)
func (f UnderlyingCashAmountField) Tag() quickfix.Tag { return tag.UnderlyingCashAmount }

//NewUnderlyingCashAmount returns a new UnderlyingCashAmountField initialized with val
func NewUnderlyingCashAmount(val quickfix.FIXFloat) UnderlyingCashAmountField {
	return UnderlyingCashAmountField{val}
}

//UnderlyingCashTypeField is a STRING field
type UnderlyingCashTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingCashType (974)
func (f UnderlyingCashTypeField) Tag() quickfix.Tag { return tag.UnderlyingCashType }

//NewUnderlyingCashType returns a new UnderlyingCashTypeField initialized with val
func NewUnderlyingCashType(val quickfix.FIXString) UnderlyingCashTypeField {
	return UnderlyingCashTypeField{val}
}

//UnderlyingCollectAmountField is a AMT field
type UnderlyingCollectAmountField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingCollectAmount (986)
func (f UnderlyingCollectAmountField) Tag() quickfix.Tag { return tag.UnderlyingCollectAmount }

//NewUnderlyingCollectAmount returns a new UnderlyingCollectAmountField initialized with val
func NewUnderlyingCollectAmount(val quickfix.FIXFloat) UnderlyingCollectAmountField {
	return UnderlyingCollectAmountField{val}
}

//UnderlyingContractMultiplierField is a FLOAT field
type UnderlyingContractMultiplierField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingContractMultiplier (436)
func (f UnderlyingContractMultiplierField) Tag() quickfix.Tag { return tag.UnderlyingContractMultiplier }

//NewUnderlyingContractMultiplier returns a new UnderlyingContractMultiplierField initialized with val
func NewUnderlyingContractMultiplier(val quickfix.FIXFloat) UnderlyingContractMultiplierField {
	return UnderlyingContractMultiplierField{val}
}

//UnderlyingContractMultiplierUnitField is a INT field
type UnderlyingContractMultiplierUnitField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingContractMultiplierUnit (1437)
func (f UnderlyingContractMultiplierUnitField) Tag() quickfix.Tag {
	return tag.UnderlyingContractMultiplierUnit
}

//NewUnderlyingContractMultiplierUnit returns a new UnderlyingContractMultiplierUnitField initialized with val
func NewUnderlyingContractMultiplierUnit(val quickfix.FIXInt) UnderlyingContractMultiplierUnitField {
	return UnderlyingContractMultiplierUnitField{val}
}

//UnderlyingCountryOfIssueField is a COUNTRY field
type UnderlyingCountryOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingCountryOfIssue (592)
func (f UnderlyingCountryOfIssueField) Tag() quickfix.Tag { return tag.UnderlyingCountryOfIssue }

//NewUnderlyingCountryOfIssue returns a new UnderlyingCountryOfIssueField initialized with val
func NewUnderlyingCountryOfIssue(val quickfix.FIXString) UnderlyingCountryOfIssueField {
	return UnderlyingCountryOfIssueField{val}
}

//UnderlyingCouponPaymentDateField is a LOCALMKTDATE field
type UnderlyingCouponPaymentDateField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingCouponPaymentDate (241)
func (f UnderlyingCouponPaymentDateField) Tag() quickfix.Tag { return tag.UnderlyingCouponPaymentDate }

//NewUnderlyingCouponPaymentDate returns a new UnderlyingCouponPaymentDateField initialized with val
func NewUnderlyingCouponPaymentDate(val quickfix.FIXString) UnderlyingCouponPaymentDateField {
	return UnderlyingCouponPaymentDateField{val}
}

//UnderlyingCouponRateField is a PERCENTAGE field
type UnderlyingCouponRateField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingCouponRate (435)
func (f UnderlyingCouponRateField) Tag() quickfix.Tag { return tag.UnderlyingCouponRate }

//NewUnderlyingCouponRate returns a new UnderlyingCouponRateField initialized with val
func NewUnderlyingCouponRate(val quickfix.FIXFloat) UnderlyingCouponRateField {
	return UnderlyingCouponRateField{val}
}

//UnderlyingCreditRatingField is a STRING field
type UnderlyingCreditRatingField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingCreditRating (256)
func (f UnderlyingCreditRatingField) Tag() quickfix.Tag { return tag.UnderlyingCreditRating }

//NewUnderlyingCreditRating returns a new UnderlyingCreditRatingField initialized with val
func NewUnderlyingCreditRating(val quickfix.FIXString) UnderlyingCreditRatingField {
	return UnderlyingCreditRatingField{val}
}

//UnderlyingCurrencyField is a CURRENCY field
type UnderlyingCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingCurrency (318)
func (f UnderlyingCurrencyField) Tag() quickfix.Tag { return tag.UnderlyingCurrency }

//NewUnderlyingCurrency returns a new UnderlyingCurrencyField initialized with val
func NewUnderlyingCurrency(val quickfix.FIXString) UnderlyingCurrencyField {
	return UnderlyingCurrencyField{val}
}

//UnderlyingCurrentValueField is a AMT field
type UnderlyingCurrentValueField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingCurrentValue (885)
func (f UnderlyingCurrentValueField) Tag() quickfix.Tag { return tag.UnderlyingCurrentValue }

//NewUnderlyingCurrentValue returns a new UnderlyingCurrentValueField initialized with val
func NewUnderlyingCurrentValue(val quickfix.FIXFloat) UnderlyingCurrentValueField {
	return UnderlyingCurrentValueField{val}
}

//UnderlyingDeliveryAmountField is a AMT field
type UnderlyingDeliveryAmountField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingDeliveryAmount (1037)
func (f UnderlyingDeliveryAmountField) Tag() quickfix.Tag { return tag.UnderlyingDeliveryAmount }

//NewUnderlyingDeliveryAmount returns a new UnderlyingDeliveryAmountField initialized with val
func NewUnderlyingDeliveryAmount(val quickfix.FIXFloat) UnderlyingDeliveryAmountField {
	return UnderlyingDeliveryAmountField{val}
}

//UnderlyingDetachmentPointField is a PERCENTAGE field
type UnderlyingDetachmentPointField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingDetachmentPoint (1460)
func (f UnderlyingDetachmentPointField) Tag() quickfix.Tag { return tag.UnderlyingDetachmentPoint }

//NewUnderlyingDetachmentPoint returns a new UnderlyingDetachmentPointField initialized with val
func NewUnderlyingDetachmentPoint(val quickfix.FIXFloat) UnderlyingDetachmentPointField {
	return UnderlyingDetachmentPointField{val}
}

//UnderlyingDirtyPriceField is a PRICE field
type UnderlyingDirtyPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingDirtyPrice (882)
func (f UnderlyingDirtyPriceField) Tag() quickfix.Tag { return tag.UnderlyingDirtyPrice }

//NewUnderlyingDirtyPrice returns a new UnderlyingDirtyPriceField initialized with val
func NewUnderlyingDirtyPrice(val quickfix.FIXFloat) UnderlyingDirtyPriceField {
	return UnderlyingDirtyPriceField{val}
}

//UnderlyingEndPriceField is a PRICE field
type UnderlyingEndPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingEndPrice (883)
func (f UnderlyingEndPriceField) Tag() quickfix.Tag { return tag.UnderlyingEndPrice }

//NewUnderlyingEndPrice returns a new UnderlyingEndPriceField initialized with val
func NewUnderlyingEndPrice(val quickfix.FIXFloat) UnderlyingEndPriceField {
	return UnderlyingEndPriceField{val}
}

//UnderlyingEndValueField is a AMT field
type UnderlyingEndValueField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingEndValue (886)
func (f UnderlyingEndValueField) Tag() quickfix.Tag { return tag.UnderlyingEndValue }

//NewUnderlyingEndValue returns a new UnderlyingEndValueField initialized with val
func NewUnderlyingEndValue(val quickfix.FIXFloat) UnderlyingEndValueField {
	return UnderlyingEndValueField{val}
}

//UnderlyingExerciseStyleField is a INT field
type UnderlyingExerciseStyleField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingExerciseStyle (1419)
func (f UnderlyingExerciseStyleField) Tag() quickfix.Tag { return tag.UnderlyingExerciseStyle }

//NewUnderlyingExerciseStyle returns a new UnderlyingExerciseStyleField initialized with val
func NewUnderlyingExerciseStyle(val quickfix.FIXInt) UnderlyingExerciseStyleField {
	return UnderlyingExerciseStyleField{val}
}

//UnderlyingFXRateField is a FLOAT field
type UnderlyingFXRateField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingFXRate (1045)
func (f UnderlyingFXRateField) Tag() quickfix.Tag { return tag.UnderlyingFXRate }

//NewUnderlyingFXRate returns a new UnderlyingFXRateField initialized with val
func NewUnderlyingFXRate(val quickfix.FIXFloat) UnderlyingFXRateField {
	return UnderlyingFXRateField{val}
}

//UnderlyingFXRateCalcField is a CHAR field
type UnderlyingFXRateCalcField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingFXRateCalc (1046)
func (f UnderlyingFXRateCalcField) Tag() quickfix.Tag { return tag.UnderlyingFXRateCalc }

//NewUnderlyingFXRateCalc returns a new UnderlyingFXRateCalcField initialized with val
func NewUnderlyingFXRateCalc(val quickfix.FIXString) UnderlyingFXRateCalcField {
	return UnderlyingFXRateCalcField{val}
}

//UnderlyingFactorField is a FLOAT field
type UnderlyingFactorField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingFactor (246)
func (f UnderlyingFactorField) Tag() quickfix.Tag { return tag.UnderlyingFactor }

//NewUnderlyingFactor returns a new UnderlyingFactorField initialized with val
func NewUnderlyingFactor(val quickfix.FIXFloat) UnderlyingFactorField {
	return UnderlyingFactorField{val}
}

//UnderlyingFlowScheduleTypeField is a INT field
type UnderlyingFlowScheduleTypeField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingFlowScheduleType (1441)
func (f UnderlyingFlowScheduleTypeField) Tag() quickfix.Tag { return tag.UnderlyingFlowScheduleType }

//NewUnderlyingFlowScheduleType returns a new UnderlyingFlowScheduleTypeField initialized with val
func NewUnderlyingFlowScheduleType(val quickfix.FIXInt) UnderlyingFlowScheduleTypeField {
	return UnderlyingFlowScheduleTypeField{val}
}

//UnderlyingIDSourceField is a STRING field
type UnderlyingIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingIDSource (305)
func (f UnderlyingIDSourceField) Tag() quickfix.Tag { return tag.UnderlyingIDSource }

//NewUnderlyingIDSource returns a new UnderlyingIDSourceField initialized with val
func NewUnderlyingIDSource(val quickfix.FIXString) UnderlyingIDSourceField {
	return UnderlyingIDSourceField{val}
}

//UnderlyingInstrRegistryField is a STRING field
type UnderlyingInstrRegistryField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingInstrRegistry (595)
func (f UnderlyingInstrRegistryField) Tag() quickfix.Tag { return tag.UnderlyingInstrRegistry }

//NewUnderlyingInstrRegistry returns a new UnderlyingInstrRegistryField initialized with val
func NewUnderlyingInstrRegistry(val quickfix.FIXString) UnderlyingInstrRegistryField {
	return UnderlyingInstrRegistryField{val}
}

//UnderlyingInstrumentPartyIDField is a STRING field
type UnderlyingInstrumentPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingInstrumentPartyID (1059)
func (f UnderlyingInstrumentPartyIDField) Tag() quickfix.Tag { return tag.UnderlyingInstrumentPartyID }

//NewUnderlyingInstrumentPartyID returns a new UnderlyingInstrumentPartyIDField initialized with val
func NewUnderlyingInstrumentPartyID(val quickfix.FIXString) UnderlyingInstrumentPartyIDField {
	return UnderlyingInstrumentPartyIDField{val}
}

//UnderlyingInstrumentPartyIDSourceField is a CHAR field
type UnderlyingInstrumentPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingInstrumentPartyIDSource (1060)
func (f UnderlyingInstrumentPartyIDSourceField) Tag() quickfix.Tag {
	return tag.UnderlyingInstrumentPartyIDSource
}

//NewUnderlyingInstrumentPartyIDSource returns a new UnderlyingInstrumentPartyIDSourceField initialized with val
func NewUnderlyingInstrumentPartyIDSource(val quickfix.FIXString) UnderlyingInstrumentPartyIDSourceField {
	return UnderlyingInstrumentPartyIDSourceField{val}
}

//UnderlyingInstrumentPartyRoleField is a INT field
type UnderlyingInstrumentPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingInstrumentPartyRole (1061)
func (f UnderlyingInstrumentPartyRoleField) Tag() quickfix.Tag {
	return tag.UnderlyingInstrumentPartyRole
}

//NewUnderlyingInstrumentPartyRole returns a new UnderlyingInstrumentPartyRoleField initialized with val
func NewUnderlyingInstrumentPartyRole(val quickfix.FIXInt) UnderlyingInstrumentPartyRoleField {
	return UnderlyingInstrumentPartyRoleField{val}
}

//UnderlyingInstrumentPartySubIDField is a STRING field
type UnderlyingInstrumentPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingInstrumentPartySubID (1063)
func (f UnderlyingInstrumentPartySubIDField) Tag() quickfix.Tag {
	return tag.UnderlyingInstrumentPartySubID
}

//NewUnderlyingInstrumentPartySubID returns a new UnderlyingInstrumentPartySubIDField initialized with val
func NewUnderlyingInstrumentPartySubID(val quickfix.FIXString) UnderlyingInstrumentPartySubIDField {
	return UnderlyingInstrumentPartySubIDField{val}
}

//UnderlyingInstrumentPartySubIDTypeField is a INT field
type UnderlyingInstrumentPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingInstrumentPartySubIDType (1064)
func (f UnderlyingInstrumentPartySubIDTypeField) Tag() quickfix.Tag {
	return tag.UnderlyingInstrumentPartySubIDType
}

//NewUnderlyingInstrumentPartySubIDType returns a new UnderlyingInstrumentPartySubIDTypeField initialized with val
func NewUnderlyingInstrumentPartySubIDType(val quickfix.FIXInt) UnderlyingInstrumentPartySubIDTypeField {
	return UnderlyingInstrumentPartySubIDTypeField{val}
}

//UnderlyingIssueDateField is a LOCALMKTDATE field
type UnderlyingIssueDateField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingIssueDate (242)
func (f UnderlyingIssueDateField) Tag() quickfix.Tag { return tag.UnderlyingIssueDate }

//NewUnderlyingIssueDate returns a new UnderlyingIssueDateField initialized with val
func NewUnderlyingIssueDate(val quickfix.FIXString) UnderlyingIssueDateField {
	return UnderlyingIssueDateField{val}
}

//UnderlyingIssuerField is a STRING field
type UnderlyingIssuerField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingIssuer (306)
func (f UnderlyingIssuerField) Tag() quickfix.Tag { return tag.UnderlyingIssuer }

//NewUnderlyingIssuer returns a new UnderlyingIssuerField initialized with val
func NewUnderlyingIssuer(val quickfix.FIXString) UnderlyingIssuerField {
	return UnderlyingIssuerField{val}
}

//UnderlyingLastPxField is a PRICE field
type UnderlyingLastPxField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingLastPx (651)
func (f UnderlyingLastPxField) Tag() quickfix.Tag { return tag.UnderlyingLastPx }

//NewUnderlyingLastPx returns a new UnderlyingLastPxField initialized with val
func NewUnderlyingLastPx(val quickfix.FIXFloat) UnderlyingLastPxField {
	return UnderlyingLastPxField{val}
}

//UnderlyingLastQtyField is a QTY field
type UnderlyingLastQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingLastQty (652)
func (f UnderlyingLastQtyField) Tag() quickfix.Tag { return tag.UnderlyingLastQty }

//NewUnderlyingLastQty returns a new UnderlyingLastQtyField initialized with val
func NewUnderlyingLastQty(val quickfix.FIXFloat) UnderlyingLastQtyField {
	return UnderlyingLastQtyField{val}
}

//UnderlyingLegCFICodeField is a STRING field
type UnderlyingLegCFICodeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegCFICode (1344)
func (f UnderlyingLegCFICodeField) Tag() quickfix.Tag { return tag.UnderlyingLegCFICode }

//NewUnderlyingLegCFICode returns a new UnderlyingLegCFICodeField initialized with val
func NewUnderlyingLegCFICode(val quickfix.FIXString) UnderlyingLegCFICodeField {
	return UnderlyingLegCFICodeField{val}
}

//UnderlyingLegMaturityDateField is a LOCALMKTDATE field
type UnderlyingLegMaturityDateField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegMaturityDate (1345)
func (f UnderlyingLegMaturityDateField) Tag() quickfix.Tag { return tag.UnderlyingLegMaturityDate }

//NewUnderlyingLegMaturityDate returns a new UnderlyingLegMaturityDateField initialized with val
func NewUnderlyingLegMaturityDate(val quickfix.FIXString) UnderlyingLegMaturityDateField {
	return UnderlyingLegMaturityDateField{val}
}

//UnderlyingLegMaturityMonthYearField is a MONTHYEAR field
type UnderlyingLegMaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegMaturityMonthYear (1339)
func (f UnderlyingLegMaturityMonthYearField) Tag() quickfix.Tag {
	return tag.UnderlyingLegMaturityMonthYear
}

//NewUnderlyingLegMaturityMonthYear returns a new UnderlyingLegMaturityMonthYearField initialized with val
func NewUnderlyingLegMaturityMonthYear(val quickfix.FIXString) UnderlyingLegMaturityMonthYearField {
	return UnderlyingLegMaturityMonthYearField{val}
}

//UnderlyingLegMaturityTimeField is a TZTIMEONLY field
type UnderlyingLegMaturityTimeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegMaturityTime (1405)
func (f UnderlyingLegMaturityTimeField) Tag() quickfix.Tag { return tag.UnderlyingLegMaturityTime }

//NewUnderlyingLegMaturityTime returns a new UnderlyingLegMaturityTimeField initialized with val
func NewUnderlyingLegMaturityTime(val quickfix.FIXString) UnderlyingLegMaturityTimeField {
	return UnderlyingLegMaturityTimeField{val}
}

//UnderlyingLegOptAttributeField is a CHAR field
type UnderlyingLegOptAttributeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegOptAttribute (1391)
func (f UnderlyingLegOptAttributeField) Tag() quickfix.Tag { return tag.UnderlyingLegOptAttribute }

//NewUnderlyingLegOptAttribute returns a new UnderlyingLegOptAttributeField initialized with val
func NewUnderlyingLegOptAttribute(val quickfix.FIXString) UnderlyingLegOptAttributeField {
	return UnderlyingLegOptAttributeField{val}
}

//UnderlyingLegPutOrCallField is a INT field
type UnderlyingLegPutOrCallField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingLegPutOrCall (1343)
func (f UnderlyingLegPutOrCallField) Tag() quickfix.Tag { return tag.UnderlyingLegPutOrCall }

//NewUnderlyingLegPutOrCall returns a new UnderlyingLegPutOrCallField initialized with val
func NewUnderlyingLegPutOrCall(val quickfix.FIXInt) UnderlyingLegPutOrCallField {
	return UnderlyingLegPutOrCallField{val}
}

//UnderlyingLegSecurityAltIDField is a STRING field
type UnderlyingLegSecurityAltIDField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSecurityAltID (1335)
func (f UnderlyingLegSecurityAltIDField) Tag() quickfix.Tag { return tag.UnderlyingLegSecurityAltID }

//NewUnderlyingLegSecurityAltID returns a new UnderlyingLegSecurityAltIDField initialized with val
func NewUnderlyingLegSecurityAltID(val quickfix.FIXString) UnderlyingLegSecurityAltIDField {
	return UnderlyingLegSecurityAltIDField{val}
}

//UnderlyingLegSecurityAltIDSourceField is a STRING field
type UnderlyingLegSecurityAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSecurityAltIDSource (1336)
func (f UnderlyingLegSecurityAltIDSourceField) Tag() quickfix.Tag {
	return tag.UnderlyingLegSecurityAltIDSource
}

//NewUnderlyingLegSecurityAltIDSource returns a new UnderlyingLegSecurityAltIDSourceField initialized with val
func NewUnderlyingLegSecurityAltIDSource(val quickfix.FIXString) UnderlyingLegSecurityAltIDSourceField {
	return UnderlyingLegSecurityAltIDSourceField{val}
}

//UnderlyingLegSecurityDescField is a STRING field
type UnderlyingLegSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSecurityDesc (1392)
func (f UnderlyingLegSecurityDescField) Tag() quickfix.Tag { return tag.UnderlyingLegSecurityDesc }

//NewUnderlyingLegSecurityDesc returns a new UnderlyingLegSecurityDescField initialized with val
func NewUnderlyingLegSecurityDesc(val quickfix.FIXString) UnderlyingLegSecurityDescField {
	return UnderlyingLegSecurityDescField{val}
}

//UnderlyingLegSecurityExchangeField is a STRING field
type UnderlyingLegSecurityExchangeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSecurityExchange (1341)
func (f UnderlyingLegSecurityExchangeField) Tag() quickfix.Tag {
	return tag.UnderlyingLegSecurityExchange
}

//NewUnderlyingLegSecurityExchange returns a new UnderlyingLegSecurityExchangeField initialized with val
func NewUnderlyingLegSecurityExchange(val quickfix.FIXString) UnderlyingLegSecurityExchangeField {
	return UnderlyingLegSecurityExchangeField{val}
}

//UnderlyingLegSecurityIDField is a STRING field
type UnderlyingLegSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSecurityID (1332)
func (f UnderlyingLegSecurityIDField) Tag() quickfix.Tag { return tag.UnderlyingLegSecurityID }

//NewUnderlyingLegSecurityID returns a new UnderlyingLegSecurityIDField initialized with val
func NewUnderlyingLegSecurityID(val quickfix.FIXString) UnderlyingLegSecurityIDField {
	return UnderlyingLegSecurityIDField{val}
}

//UnderlyingLegSecurityIDSourceField is a STRING field
type UnderlyingLegSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSecurityIDSource (1333)
func (f UnderlyingLegSecurityIDSourceField) Tag() quickfix.Tag {
	return tag.UnderlyingLegSecurityIDSource
}

//NewUnderlyingLegSecurityIDSource returns a new UnderlyingLegSecurityIDSourceField initialized with val
func NewUnderlyingLegSecurityIDSource(val quickfix.FIXString) UnderlyingLegSecurityIDSourceField {
	return UnderlyingLegSecurityIDSourceField{val}
}

//UnderlyingLegSecuritySubTypeField is a STRING field
type UnderlyingLegSecuritySubTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSecuritySubType (1338)
func (f UnderlyingLegSecuritySubTypeField) Tag() quickfix.Tag { return tag.UnderlyingLegSecuritySubType }

//NewUnderlyingLegSecuritySubType returns a new UnderlyingLegSecuritySubTypeField initialized with val
func NewUnderlyingLegSecuritySubType(val quickfix.FIXString) UnderlyingLegSecuritySubTypeField {
	return UnderlyingLegSecuritySubTypeField{val}
}

//UnderlyingLegSecurityTypeField is a STRING field
type UnderlyingLegSecurityTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSecurityType (1337)
func (f UnderlyingLegSecurityTypeField) Tag() quickfix.Tag { return tag.UnderlyingLegSecurityType }

//NewUnderlyingLegSecurityType returns a new UnderlyingLegSecurityTypeField initialized with val
func NewUnderlyingLegSecurityType(val quickfix.FIXString) UnderlyingLegSecurityTypeField {
	return UnderlyingLegSecurityTypeField{val}
}

//UnderlyingLegStrikePriceField is a PRICE field
type UnderlyingLegStrikePriceField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingLegStrikePrice (1340)
func (f UnderlyingLegStrikePriceField) Tag() quickfix.Tag { return tag.UnderlyingLegStrikePrice }

//NewUnderlyingLegStrikePrice returns a new UnderlyingLegStrikePriceField initialized with val
func NewUnderlyingLegStrikePrice(val quickfix.FIXFloat) UnderlyingLegStrikePriceField {
	return UnderlyingLegStrikePriceField{val}
}

//UnderlyingLegSymbolField is a STRING field
type UnderlyingLegSymbolField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSymbol (1330)
func (f UnderlyingLegSymbolField) Tag() quickfix.Tag { return tag.UnderlyingLegSymbol }

//NewUnderlyingLegSymbol returns a new UnderlyingLegSymbolField initialized with val
func NewUnderlyingLegSymbol(val quickfix.FIXString) UnderlyingLegSymbolField {
	return UnderlyingLegSymbolField{val}
}

//UnderlyingLegSymbolSfxField is a STRING field
type UnderlyingLegSymbolSfxField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLegSymbolSfx (1331)
func (f UnderlyingLegSymbolSfxField) Tag() quickfix.Tag { return tag.UnderlyingLegSymbolSfx }

//NewUnderlyingLegSymbolSfx returns a new UnderlyingLegSymbolSfxField initialized with val
func NewUnderlyingLegSymbolSfx(val quickfix.FIXString) UnderlyingLegSymbolSfxField {
	return UnderlyingLegSymbolSfxField{val}
}

//UnderlyingLocaleOfIssueField is a STRING field
type UnderlyingLocaleOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingLocaleOfIssue (594)
func (f UnderlyingLocaleOfIssueField) Tag() quickfix.Tag { return tag.UnderlyingLocaleOfIssue }

//NewUnderlyingLocaleOfIssue returns a new UnderlyingLocaleOfIssueField initialized with val
func NewUnderlyingLocaleOfIssue(val quickfix.FIXString) UnderlyingLocaleOfIssueField {
	return UnderlyingLocaleOfIssueField{val}
}

//UnderlyingMaturityDateField is a LOCALMKTDATE field
type UnderlyingMaturityDateField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingMaturityDate (542)
func (f UnderlyingMaturityDateField) Tag() quickfix.Tag { return tag.UnderlyingMaturityDate }

//NewUnderlyingMaturityDate returns a new UnderlyingMaturityDateField initialized with val
func NewUnderlyingMaturityDate(val quickfix.FIXString) UnderlyingMaturityDateField {
	return UnderlyingMaturityDateField{val}
}

//UnderlyingMaturityDayField is a DAYOFMONTH field
type UnderlyingMaturityDayField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingMaturityDay (314)
func (f UnderlyingMaturityDayField) Tag() quickfix.Tag { return tag.UnderlyingMaturityDay }

//NewUnderlyingMaturityDay returns a new UnderlyingMaturityDayField initialized with val
func NewUnderlyingMaturityDay(val quickfix.FIXInt) UnderlyingMaturityDayField {
	return UnderlyingMaturityDayField{val}
}

//UnderlyingMaturityMonthYearField is a MONTHYEAR field
type UnderlyingMaturityMonthYearField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingMaturityMonthYear (313)
func (f UnderlyingMaturityMonthYearField) Tag() quickfix.Tag { return tag.UnderlyingMaturityMonthYear }

//NewUnderlyingMaturityMonthYear returns a new UnderlyingMaturityMonthYearField initialized with val
func NewUnderlyingMaturityMonthYear(val quickfix.FIXString) UnderlyingMaturityMonthYearField {
	return UnderlyingMaturityMonthYearField{val}
}

//UnderlyingMaturityTimeField is a TZTIMEONLY field
type UnderlyingMaturityTimeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingMaturityTime (1213)
func (f UnderlyingMaturityTimeField) Tag() quickfix.Tag { return tag.UnderlyingMaturityTime }

//NewUnderlyingMaturityTime returns a new UnderlyingMaturityTimeField initialized with val
func NewUnderlyingMaturityTime(val quickfix.FIXString) UnderlyingMaturityTimeField {
	return UnderlyingMaturityTimeField{val}
}

//UnderlyingNotionalPercentageOutstandingField is a PERCENTAGE field
type UnderlyingNotionalPercentageOutstandingField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingNotionalPercentageOutstanding (1455)
func (f UnderlyingNotionalPercentageOutstandingField) Tag() quickfix.Tag {
	return tag.UnderlyingNotionalPercentageOutstanding
}

//NewUnderlyingNotionalPercentageOutstanding returns a new UnderlyingNotionalPercentageOutstandingField initialized with val
func NewUnderlyingNotionalPercentageOutstanding(val quickfix.FIXFloat) UnderlyingNotionalPercentageOutstandingField {
	return UnderlyingNotionalPercentageOutstandingField{val}
}

//UnderlyingOptAttributeField is a CHAR field
type UnderlyingOptAttributeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingOptAttribute (317)
func (f UnderlyingOptAttributeField) Tag() quickfix.Tag { return tag.UnderlyingOptAttribute }

//NewUnderlyingOptAttribute returns a new UnderlyingOptAttributeField initialized with val
func NewUnderlyingOptAttribute(val quickfix.FIXString) UnderlyingOptAttributeField {
	return UnderlyingOptAttributeField{val}
}

//UnderlyingOriginalNotionalPercentageOutstandingField is a PERCENTAGE field
type UnderlyingOriginalNotionalPercentageOutstandingField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingOriginalNotionalPercentageOutstanding (1456)
func (f UnderlyingOriginalNotionalPercentageOutstandingField) Tag() quickfix.Tag {
	return tag.UnderlyingOriginalNotionalPercentageOutstanding
}

//NewUnderlyingOriginalNotionalPercentageOutstanding returns a new UnderlyingOriginalNotionalPercentageOutstandingField initialized with val
func NewUnderlyingOriginalNotionalPercentageOutstanding(val quickfix.FIXFloat) UnderlyingOriginalNotionalPercentageOutstandingField {
	return UnderlyingOriginalNotionalPercentageOutstandingField{val}
}

//UnderlyingPayAmountField is a AMT field
type UnderlyingPayAmountField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingPayAmount (985)
func (f UnderlyingPayAmountField) Tag() quickfix.Tag { return tag.UnderlyingPayAmount }

//NewUnderlyingPayAmount returns a new UnderlyingPayAmountField initialized with val
func NewUnderlyingPayAmount(val quickfix.FIXFloat) UnderlyingPayAmountField {
	return UnderlyingPayAmountField{val}
}

//UnderlyingPriceDeterminationMethodField is a INT field
type UnderlyingPriceDeterminationMethodField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingPriceDeterminationMethod (1481)
func (f UnderlyingPriceDeterminationMethodField) Tag() quickfix.Tag {
	return tag.UnderlyingPriceDeterminationMethod
}

//NewUnderlyingPriceDeterminationMethod returns a new UnderlyingPriceDeterminationMethodField initialized with val
func NewUnderlyingPriceDeterminationMethod(val quickfix.FIXInt) UnderlyingPriceDeterminationMethodField {
	return UnderlyingPriceDeterminationMethodField{val}
}

//UnderlyingPriceUnitOfMeasureField is a STRING field
type UnderlyingPriceUnitOfMeasureField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingPriceUnitOfMeasure (1424)
func (f UnderlyingPriceUnitOfMeasureField) Tag() quickfix.Tag { return tag.UnderlyingPriceUnitOfMeasure }

//NewUnderlyingPriceUnitOfMeasure returns a new UnderlyingPriceUnitOfMeasureField initialized with val
func NewUnderlyingPriceUnitOfMeasure(val quickfix.FIXString) UnderlyingPriceUnitOfMeasureField {
	return UnderlyingPriceUnitOfMeasureField{val}
}

//UnderlyingPriceUnitOfMeasureQtyField is a QTY field
type UnderlyingPriceUnitOfMeasureQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingPriceUnitOfMeasureQty (1425)
func (f UnderlyingPriceUnitOfMeasureQtyField) Tag() quickfix.Tag {
	return tag.UnderlyingPriceUnitOfMeasureQty
}

//NewUnderlyingPriceUnitOfMeasureQty returns a new UnderlyingPriceUnitOfMeasureQtyField initialized with val
func NewUnderlyingPriceUnitOfMeasureQty(val quickfix.FIXFloat) UnderlyingPriceUnitOfMeasureQtyField {
	return UnderlyingPriceUnitOfMeasureQtyField{val}
}

//UnderlyingProductField is a INT field
type UnderlyingProductField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingProduct (462)
func (f UnderlyingProductField) Tag() quickfix.Tag { return tag.UnderlyingProduct }

//NewUnderlyingProduct returns a new UnderlyingProductField initialized with val
func NewUnderlyingProduct(val quickfix.FIXInt) UnderlyingProductField {
	return UnderlyingProductField{val}
}

//UnderlyingPutOrCallField is a INT field
type UnderlyingPutOrCallField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingPutOrCall (315)
func (f UnderlyingPutOrCallField) Tag() quickfix.Tag { return tag.UnderlyingPutOrCall }

//NewUnderlyingPutOrCall returns a new UnderlyingPutOrCallField initialized with val
func NewUnderlyingPutOrCall(val quickfix.FIXInt) UnderlyingPutOrCallField {
	return UnderlyingPutOrCallField{val}
}

//UnderlyingPxField is a PRICE field
type UnderlyingPxField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingPx (810)
func (f UnderlyingPxField) Tag() quickfix.Tag { return tag.UnderlyingPx }

//NewUnderlyingPx returns a new UnderlyingPxField initialized with val
func NewUnderlyingPx(val quickfix.FIXFloat) UnderlyingPxField {
	return UnderlyingPxField{val}
}

//UnderlyingQtyField is a QTY field
type UnderlyingQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingQty (879)
func (f UnderlyingQtyField) Tag() quickfix.Tag { return tag.UnderlyingQty }

//NewUnderlyingQty returns a new UnderlyingQtyField initialized with val
func NewUnderlyingQty(val quickfix.FIXFloat) UnderlyingQtyField {
	return UnderlyingQtyField{val}
}

//UnderlyingRedemptionDateField is a LOCALMKTDATE field
type UnderlyingRedemptionDateField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingRedemptionDate (247)
func (f UnderlyingRedemptionDateField) Tag() quickfix.Tag { return tag.UnderlyingRedemptionDate }

//NewUnderlyingRedemptionDate returns a new UnderlyingRedemptionDateField initialized with val
func NewUnderlyingRedemptionDate(val quickfix.FIXString) UnderlyingRedemptionDateField {
	return UnderlyingRedemptionDateField{val}
}

//UnderlyingRepoCollateralSecurityTypeField is a INT field
type UnderlyingRepoCollateralSecurityTypeField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingRepoCollateralSecurityType (243)
func (f UnderlyingRepoCollateralSecurityTypeField) Tag() quickfix.Tag {
	return tag.UnderlyingRepoCollateralSecurityType
}

//NewUnderlyingRepoCollateralSecurityType returns a new UnderlyingRepoCollateralSecurityTypeField initialized with val
func NewUnderlyingRepoCollateralSecurityType(val quickfix.FIXInt) UnderlyingRepoCollateralSecurityTypeField {
	return UnderlyingRepoCollateralSecurityTypeField{val}
}

//UnderlyingRepurchaseRateField is a PERCENTAGE field
type UnderlyingRepurchaseRateField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingRepurchaseRate (245)
func (f UnderlyingRepurchaseRateField) Tag() quickfix.Tag { return tag.UnderlyingRepurchaseRate }

//NewUnderlyingRepurchaseRate returns a new UnderlyingRepurchaseRateField initialized with val
func NewUnderlyingRepurchaseRate(val quickfix.FIXFloat) UnderlyingRepurchaseRateField {
	return UnderlyingRepurchaseRateField{val}
}

//UnderlyingRepurchaseTermField is a INT field
type UnderlyingRepurchaseTermField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingRepurchaseTerm (244)
func (f UnderlyingRepurchaseTermField) Tag() quickfix.Tag { return tag.UnderlyingRepurchaseTerm }

//NewUnderlyingRepurchaseTerm returns a new UnderlyingRepurchaseTermField initialized with val
func NewUnderlyingRepurchaseTerm(val quickfix.FIXInt) UnderlyingRepurchaseTermField {
	return UnderlyingRepurchaseTermField{val}
}

//UnderlyingRestructuringTypeField is a STRING field
type UnderlyingRestructuringTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingRestructuringType (1453)
func (f UnderlyingRestructuringTypeField) Tag() quickfix.Tag { return tag.UnderlyingRestructuringType }

//NewUnderlyingRestructuringType returns a new UnderlyingRestructuringTypeField initialized with val
func NewUnderlyingRestructuringType(val quickfix.FIXString) UnderlyingRestructuringTypeField {
	return UnderlyingRestructuringTypeField{val}
}

//UnderlyingSecurityAltIDField is a STRING field
type UnderlyingSecurityAltIDField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSecurityAltID (458)
func (f UnderlyingSecurityAltIDField) Tag() quickfix.Tag { return tag.UnderlyingSecurityAltID }

//NewUnderlyingSecurityAltID returns a new UnderlyingSecurityAltIDField initialized with val
func NewUnderlyingSecurityAltID(val quickfix.FIXString) UnderlyingSecurityAltIDField {
	return UnderlyingSecurityAltIDField{val}
}

//UnderlyingSecurityAltIDSourceField is a STRING field
type UnderlyingSecurityAltIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSecurityAltIDSource (459)
func (f UnderlyingSecurityAltIDSourceField) Tag() quickfix.Tag {
	return tag.UnderlyingSecurityAltIDSource
}

//NewUnderlyingSecurityAltIDSource returns a new UnderlyingSecurityAltIDSourceField initialized with val
func NewUnderlyingSecurityAltIDSource(val quickfix.FIXString) UnderlyingSecurityAltIDSourceField {
	return UnderlyingSecurityAltIDSourceField{val}
}

//UnderlyingSecurityDescField is a STRING field
type UnderlyingSecurityDescField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSecurityDesc (307)
func (f UnderlyingSecurityDescField) Tag() quickfix.Tag { return tag.UnderlyingSecurityDesc }

//NewUnderlyingSecurityDesc returns a new UnderlyingSecurityDescField initialized with val
func NewUnderlyingSecurityDesc(val quickfix.FIXString) UnderlyingSecurityDescField {
	return UnderlyingSecurityDescField{val}
}

//UnderlyingSecurityExchangeField is a EXCHANGE field
type UnderlyingSecurityExchangeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSecurityExchange (308)
func (f UnderlyingSecurityExchangeField) Tag() quickfix.Tag { return tag.UnderlyingSecurityExchange }

//NewUnderlyingSecurityExchange returns a new UnderlyingSecurityExchangeField initialized with val
func NewUnderlyingSecurityExchange(val quickfix.FIXString) UnderlyingSecurityExchangeField {
	return UnderlyingSecurityExchangeField{val}
}

//UnderlyingSecurityIDField is a STRING field
type UnderlyingSecurityIDField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSecurityID (309)
func (f UnderlyingSecurityIDField) Tag() quickfix.Tag { return tag.UnderlyingSecurityID }

//NewUnderlyingSecurityID returns a new UnderlyingSecurityIDField initialized with val
func NewUnderlyingSecurityID(val quickfix.FIXString) UnderlyingSecurityIDField {
	return UnderlyingSecurityIDField{val}
}

//UnderlyingSecurityIDSourceField is a STRING field
type UnderlyingSecurityIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSecurityIDSource (305)
func (f UnderlyingSecurityIDSourceField) Tag() quickfix.Tag { return tag.UnderlyingSecurityIDSource }

//NewUnderlyingSecurityIDSource returns a new UnderlyingSecurityIDSourceField initialized with val
func NewUnderlyingSecurityIDSource(val quickfix.FIXString) UnderlyingSecurityIDSourceField {
	return UnderlyingSecurityIDSourceField{val}
}

//UnderlyingSecuritySubTypeField is a STRING field
type UnderlyingSecuritySubTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSecuritySubType (763)
func (f UnderlyingSecuritySubTypeField) Tag() quickfix.Tag { return tag.UnderlyingSecuritySubType }

//NewUnderlyingSecuritySubType returns a new UnderlyingSecuritySubTypeField initialized with val
func NewUnderlyingSecuritySubType(val quickfix.FIXString) UnderlyingSecuritySubTypeField {
	return UnderlyingSecuritySubTypeField{val}
}

//UnderlyingSecurityTypeField is a STRING field
type UnderlyingSecurityTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSecurityType (310)
func (f UnderlyingSecurityTypeField) Tag() quickfix.Tag { return tag.UnderlyingSecurityType }

//NewUnderlyingSecurityType returns a new UnderlyingSecurityTypeField initialized with val
func NewUnderlyingSecurityType(val quickfix.FIXString) UnderlyingSecurityTypeField {
	return UnderlyingSecurityTypeField{val}
}

//UnderlyingSeniorityField is a STRING field
type UnderlyingSeniorityField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSeniority (1454)
func (f UnderlyingSeniorityField) Tag() quickfix.Tag { return tag.UnderlyingSeniority }

//NewUnderlyingSeniority returns a new UnderlyingSeniorityField initialized with val
func NewUnderlyingSeniority(val quickfix.FIXString) UnderlyingSeniorityField {
	return UnderlyingSeniorityField{val}
}

//UnderlyingSettlMethodField is a STRING field
type UnderlyingSettlMethodField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSettlMethod (1039)
func (f UnderlyingSettlMethodField) Tag() quickfix.Tag { return tag.UnderlyingSettlMethod }

//NewUnderlyingSettlMethod returns a new UnderlyingSettlMethodField initialized with val
func NewUnderlyingSettlMethod(val quickfix.FIXString) UnderlyingSettlMethodField {
	return UnderlyingSettlMethodField{val}
}

//UnderlyingSettlPriceField is a PRICE field
type UnderlyingSettlPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingSettlPrice (732)
func (f UnderlyingSettlPriceField) Tag() quickfix.Tag { return tag.UnderlyingSettlPrice }

//NewUnderlyingSettlPrice returns a new UnderlyingSettlPriceField initialized with val
func NewUnderlyingSettlPrice(val quickfix.FIXFloat) UnderlyingSettlPriceField {
	return UnderlyingSettlPriceField{val}
}

//UnderlyingSettlPriceTypeField is a INT field
type UnderlyingSettlPriceTypeField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingSettlPriceType (733)
func (f UnderlyingSettlPriceTypeField) Tag() quickfix.Tag { return tag.UnderlyingSettlPriceType }

//NewUnderlyingSettlPriceType returns a new UnderlyingSettlPriceTypeField initialized with val
func NewUnderlyingSettlPriceType(val quickfix.FIXInt) UnderlyingSettlPriceTypeField {
	return UnderlyingSettlPriceTypeField{val}
}

//UnderlyingSettlementDateField is a LOCALMKTDATE field
type UnderlyingSettlementDateField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSettlementDate (987)
func (f UnderlyingSettlementDateField) Tag() quickfix.Tag { return tag.UnderlyingSettlementDate }

//NewUnderlyingSettlementDate returns a new UnderlyingSettlementDateField initialized with val
func NewUnderlyingSettlementDate(val quickfix.FIXString) UnderlyingSettlementDateField {
	return UnderlyingSettlementDateField{val}
}

//UnderlyingSettlementStatusField is a STRING field
type UnderlyingSettlementStatusField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSettlementStatus (988)
func (f UnderlyingSettlementStatusField) Tag() quickfix.Tag { return tag.UnderlyingSettlementStatus }

//NewUnderlyingSettlementStatus returns a new UnderlyingSettlementStatusField initialized with val
func NewUnderlyingSettlementStatus(val quickfix.FIXString) UnderlyingSettlementStatusField {
	return UnderlyingSettlementStatusField{val}
}

//UnderlyingSettlementTypeField is a INT field
type UnderlyingSettlementTypeField struct{ quickfix.FIXInt }

//Tag returns tag.UnderlyingSettlementType (975)
func (f UnderlyingSettlementTypeField) Tag() quickfix.Tag { return tag.UnderlyingSettlementType }

//NewUnderlyingSettlementType returns a new UnderlyingSettlementTypeField initialized with val
func NewUnderlyingSettlementType(val quickfix.FIXInt) UnderlyingSettlementTypeField {
	return UnderlyingSettlementTypeField{val}
}

//UnderlyingStartValueField is a AMT field
type UnderlyingStartValueField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingStartValue (884)
func (f UnderlyingStartValueField) Tag() quickfix.Tag { return tag.UnderlyingStartValue }

//NewUnderlyingStartValue returns a new UnderlyingStartValueField initialized with val
func NewUnderlyingStartValue(val quickfix.FIXFloat) UnderlyingStartValueField {
	return UnderlyingStartValueField{val}
}

//UnderlyingStateOrProvinceOfIssueField is a STRING field
type UnderlyingStateOrProvinceOfIssueField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingStateOrProvinceOfIssue (593)
func (f UnderlyingStateOrProvinceOfIssueField) Tag() quickfix.Tag {
	return tag.UnderlyingStateOrProvinceOfIssue
}

//NewUnderlyingStateOrProvinceOfIssue returns a new UnderlyingStateOrProvinceOfIssueField initialized with val
func NewUnderlyingStateOrProvinceOfIssue(val quickfix.FIXString) UnderlyingStateOrProvinceOfIssueField {
	return UnderlyingStateOrProvinceOfIssueField{val}
}

//UnderlyingStipTypeField is a STRING field
type UnderlyingStipTypeField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingStipType (888)
func (f UnderlyingStipTypeField) Tag() quickfix.Tag { return tag.UnderlyingStipType }

//NewUnderlyingStipType returns a new UnderlyingStipTypeField initialized with val
func NewUnderlyingStipType(val quickfix.FIXString) UnderlyingStipTypeField {
	return UnderlyingStipTypeField{val}
}

//UnderlyingStipValueField is a STRING field
type UnderlyingStipValueField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingStipValue (889)
func (f UnderlyingStipValueField) Tag() quickfix.Tag { return tag.UnderlyingStipValue }

//NewUnderlyingStipValue returns a new UnderlyingStipValueField initialized with val
func NewUnderlyingStipValue(val quickfix.FIXString) UnderlyingStipValueField {
	return UnderlyingStipValueField{val}
}

//UnderlyingStrikeCurrencyField is a CURRENCY field
type UnderlyingStrikeCurrencyField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingStrikeCurrency (941)
func (f UnderlyingStrikeCurrencyField) Tag() quickfix.Tag { return tag.UnderlyingStrikeCurrency }

//NewUnderlyingStrikeCurrency returns a new UnderlyingStrikeCurrencyField initialized with val
func NewUnderlyingStrikeCurrency(val quickfix.FIXString) UnderlyingStrikeCurrencyField {
	return UnderlyingStrikeCurrencyField{val}
}

//UnderlyingStrikePriceField is a PRICE field
type UnderlyingStrikePriceField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingStrikePrice (316)
func (f UnderlyingStrikePriceField) Tag() quickfix.Tag { return tag.UnderlyingStrikePrice }

//NewUnderlyingStrikePrice returns a new UnderlyingStrikePriceField initialized with val
func NewUnderlyingStrikePrice(val quickfix.FIXFloat) UnderlyingStrikePriceField {
	return UnderlyingStrikePriceField{val}
}

//UnderlyingSymbolField is a STRING field
type UnderlyingSymbolField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSymbol (311)
func (f UnderlyingSymbolField) Tag() quickfix.Tag { return tag.UnderlyingSymbol }

//NewUnderlyingSymbol returns a new UnderlyingSymbolField initialized with val
func NewUnderlyingSymbol(val quickfix.FIXString) UnderlyingSymbolField {
	return UnderlyingSymbolField{val}
}

//UnderlyingSymbolSfxField is a STRING field
type UnderlyingSymbolSfxField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingSymbolSfx (312)
func (f UnderlyingSymbolSfxField) Tag() quickfix.Tag { return tag.UnderlyingSymbolSfx }

//NewUnderlyingSymbolSfx returns a new UnderlyingSymbolSfxField initialized with val
func NewUnderlyingSymbolSfx(val quickfix.FIXString) UnderlyingSymbolSfxField {
	return UnderlyingSymbolSfxField{val}
}

//UnderlyingTimeUnitField is a STRING field
type UnderlyingTimeUnitField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingTimeUnit (1000)
func (f UnderlyingTimeUnitField) Tag() quickfix.Tag { return tag.UnderlyingTimeUnit }

//NewUnderlyingTimeUnit returns a new UnderlyingTimeUnitField initialized with val
func NewUnderlyingTimeUnit(val quickfix.FIXString) UnderlyingTimeUnitField {
	return UnderlyingTimeUnitField{val}
}

//UnderlyingTradingSessionIDField is a STRING field
type UnderlyingTradingSessionIDField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingTradingSessionID (822)
func (f UnderlyingTradingSessionIDField) Tag() quickfix.Tag { return tag.UnderlyingTradingSessionID }

//NewUnderlyingTradingSessionID returns a new UnderlyingTradingSessionIDField initialized with val
func NewUnderlyingTradingSessionID(val quickfix.FIXString) UnderlyingTradingSessionIDField {
	return UnderlyingTradingSessionIDField{val}
}

//UnderlyingTradingSessionSubIDField is a STRING field
type UnderlyingTradingSessionSubIDField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingTradingSessionSubID (823)
func (f UnderlyingTradingSessionSubIDField) Tag() quickfix.Tag {
	return tag.UnderlyingTradingSessionSubID
}

//NewUnderlyingTradingSessionSubID returns a new UnderlyingTradingSessionSubIDField initialized with val
func NewUnderlyingTradingSessionSubID(val quickfix.FIXString) UnderlyingTradingSessionSubIDField {
	return UnderlyingTradingSessionSubIDField{val}
}

//UnderlyingUnitOfMeasureField is a STRING field
type UnderlyingUnitOfMeasureField struct{ quickfix.FIXString }

//Tag returns tag.UnderlyingUnitOfMeasure (998)
func (f UnderlyingUnitOfMeasureField) Tag() quickfix.Tag { return tag.UnderlyingUnitOfMeasure }

//NewUnderlyingUnitOfMeasure returns a new UnderlyingUnitOfMeasureField initialized with val
func NewUnderlyingUnitOfMeasure(val quickfix.FIXString) UnderlyingUnitOfMeasureField {
	return UnderlyingUnitOfMeasureField{val}
}

//UnderlyingUnitOfMeasureQtyField is a QTY field
type UnderlyingUnitOfMeasureQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.UnderlyingUnitOfMeasureQty (1423)
func (f UnderlyingUnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.UnderlyingUnitOfMeasureQty }

//NewUnderlyingUnitOfMeasureQty returns a new UnderlyingUnitOfMeasureQtyField initialized with val
func NewUnderlyingUnitOfMeasureQty(val quickfix.FIXFloat) UnderlyingUnitOfMeasureQtyField {
	return UnderlyingUnitOfMeasureQtyField{val}
}

//UndlyInstrumentPartyIDField is a STRING field
type UndlyInstrumentPartyIDField struct{ quickfix.FIXString }

//Tag returns tag.UndlyInstrumentPartyID (1059)
func (f UndlyInstrumentPartyIDField) Tag() quickfix.Tag { return tag.UndlyInstrumentPartyID }

//NewUndlyInstrumentPartyID returns a new UndlyInstrumentPartyIDField initialized with val
func NewUndlyInstrumentPartyID(val quickfix.FIXString) UndlyInstrumentPartyIDField {
	return UndlyInstrumentPartyIDField{val}
}

//UndlyInstrumentPartyIDSourceField is a CHAR field
type UndlyInstrumentPartyIDSourceField struct{ quickfix.FIXString }

//Tag returns tag.UndlyInstrumentPartyIDSource (1060)
func (f UndlyInstrumentPartyIDSourceField) Tag() quickfix.Tag { return tag.UndlyInstrumentPartyIDSource }

//NewUndlyInstrumentPartyIDSource returns a new UndlyInstrumentPartyIDSourceField initialized with val
func NewUndlyInstrumentPartyIDSource(val quickfix.FIXString) UndlyInstrumentPartyIDSourceField {
	return UndlyInstrumentPartyIDSourceField{val}
}

//UndlyInstrumentPartyRoleField is a INT field
type UndlyInstrumentPartyRoleField struct{ quickfix.FIXInt }

//Tag returns tag.UndlyInstrumentPartyRole (1061)
func (f UndlyInstrumentPartyRoleField) Tag() quickfix.Tag { return tag.UndlyInstrumentPartyRole }

//NewUndlyInstrumentPartyRole returns a new UndlyInstrumentPartyRoleField initialized with val
func NewUndlyInstrumentPartyRole(val quickfix.FIXInt) UndlyInstrumentPartyRoleField {
	return UndlyInstrumentPartyRoleField{val}
}

//UndlyInstrumentPartySubIDField is a STRING field
type UndlyInstrumentPartySubIDField struct{ quickfix.FIXString }

//Tag returns tag.UndlyInstrumentPartySubID (1063)
func (f UndlyInstrumentPartySubIDField) Tag() quickfix.Tag { return tag.UndlyInstrumentPartySubID }

//NewUndlyInstrumentPartySubID returns a new UndlyInstrumentPartySubIDField initialized with val
func NewUndlyInstrumentPartySubID(val quickfix.FIXString) UndlyInstrumentPartySubIDField {
	return UndlyInstrumentPartySubIDField{val}
}

//UndlyInstrumentPartySubIDTypeField is a INT field
type UndlyInstrumentPartySubIDTypeField struct{ quickfix.FIXInt }

//Tag returns tag.UndlyInstrumentPartySubIDType (1064)
func (f UndlyInstrumentPartySubIDTypeField) Tag() quickfix.Tag {
	return tag.UndlyInstrumentPartySubIDType
}

//NewUndlyInstrumentPartySubIDType returns a new UndlyInstrumentPartySubIDTypeField initialized with val
func NewUndlyInstrumentPartySubIDType(val quickfix.FIXInt) UndlyInstrumentPartySubIDTypeField {
	return UndlyInstrumentPartySubIDTypeField{val}
}

//UnitOfMeasureField is a STRING field
type UnitOfMeasureField struct{ quickfix.FIXString }

//Tag returns tag.UnitOfMeasure (996)
func (f UnitOfMeasureField) Tag() quickfix.Tag { return tag.UnitOfMeasure }

//NewUnitOfMeasure returns a new UnitOfMeasureField initialized with val
func NewUnitOfMeasure(val quickfix.FIXString) UnitOfMeasureField {
	return UnitOfMeasureField{val}
}

//UnitOfMeasureQtyField is a QTY field
type UnitOfMeasureQtyField struct{ quickfix.FIXFloat }

//Tag returns tag.UnitOfMeasureQty (1147)
func (f UnitOfMeasureQtyField) Tag() quickfix.Tag { return tag.UnitOfMeasureQty }

//NewUnitOfMeasureQty returns a new UnitOfMeasureQtyField initialized with val
func NewUnitOfMeasureQty(val quickfix.FIXFloat) UnitOfMeasureQtyField {
	return UnitOfMeasureQtyField{val}
}

//UnsolicitedIndicatorField is a BOOLEAN field
type UnsolicitedIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.UnsolicitedIndicator (325)
func (f UnsolicitedIndicatorField) Tag() quickfix.Tag { return tag.UnsolicitedIndicator }

//NewUnsolicitedIndicator returns a new UnsolicitedIndicatorField initialized with val
func NewUnsolicitedIndicator(val quickfix.FIXBoolean) UnsolicitedIndicatorField {
	return UnsolicitedIndicatorField{val}
}

//UrgencyField is a CHAR field
type UrgencyField struct{ quickfix.FIXString }

//Tag returns tag.Urgency (61)
func (f UrgencyField) Tag() quickfix.Tag { return tag.Urgency }

//NewUrgency returns a new UrgencyField initialized with val
func NewUrgency(val quickfix.FIXString) UrgencyField {
	return UrgencyField{val}
}

//UserRequestIDField is a STRING field
type UserRequestIDField struct{ quickfix.FIXString }

//Tag returns tag.UserRequestID (923)
func (f UserRequestIDField) Tag() quickfix.Tag { return tag.UserRequestID }

//NewUserRequestID returns a new UserRequestIDField initialized with val
func NewUserRequestID(val quickfix.FIXString) UserRequestIDField {
	return UserRequestIDField{val}
}

//UserRequestTypeField is a INT field
type UserRequestTypeField struct{ quickfix.FIXInt }

//Tag returns tag.UserRequestType (924)
func (f UserRequestTypeField) Tag() quickfix.Tag { return tag.UserRequestType }

//NewUserRequestType returns a new UserRequestTypeField initialized with val
func NewUserRequestType(val quickfix.FIXInt) UserRequestTypeField {
	return UserRequestTypeField{val}
}

//UserStatusField is a INT field
type UserStatusField struct{ quickfix.FIXInt }

//Tag returns tag.UserStatus (926)
func (f UserStatusField) Tag() quickfix.Tag { return tag.UserStatus }

//NewUserStatus returns a new UserStatusField initialized with val
func NewUserStatus(val quickfix.FIXInt) UserStatusField {
	return UserStatusField{val}
}

//UserStatusTextField is a STRING field
type UserStatusTextField struct{ quickfix.FIXString }

//Tag returns tag.UserStatusText (927)
func (f UserStatusTextField) Tag() quickfix.Tag { return tag.UserStatusText }

//NewUserStatusText returns a new UserStatusTextField initialized with val
func NewUserStatusText(val quickfix.FIXString) UserStatusTextField {
	return UserStatusTextField{val}
}

//UsernameField is a STRING field
type UsernameField struct{ quickfix.FIXString }

//Tag returns tag.Username (553)
func (f UsernameField) Tag() quickfix.Tag { return tag.Username }

//NewUsername returns a new UsernameField initialized with val
func NewUsername(val quickfix.FIXString) UsernameField {
	return UsernameField{val}
}

//ValidUntilTimeField is a UTCTIMESTAMP field
type ValidUntilTimeField struct{ quickfix.FIXUTCTimestamp }

//Tag returns tag.ValidUntilTime (62)
func (f ValidUntilTimeField) Tag() quickfix.Tag { return tag.ValidUntilTime }

//NewValidUntilTime returns a new ValidUntilTimeField initialized with val
func NewValidUntilTime(val time.Time) ValidUntilTimeField {
	return ValidUntilTimeField{quickfix.FIXUTCTimestamp{Value: val}}
}

//NewValidUntilTimeNoMillis returns a new ValidUntilTimeField initialized with val without millisecs
func NewValidUntilTimeNoMillis(val time.Time) ValidUntilTimeField {
	return ValidUntilTimeField{quickfix.FIXUTCTimestamp{Value: val, NoMillis: true}}
}

//ValuationMethodField is a STRING field
type ValuationMethodField struct{ quickfix.FIXString }

//Tag returns tag.ValuationMethod (1197)
func (f ValuationMethodField) Tag() quickfix.Tag { return tag.ValuationMethod }

//NewValuationMethod returns a new ValuationMethodField initialized with val
func NewValuationMethod(val quickfix.FIXString) ValuationMethodField {
	return ValuationMethodField{val}
}

//ValueOfFuturesField is a AMT field
type ValueOfFuturesField struct{ quickfix.FIXFloat }

//Tag returns tag.ValueOfFutures (408)
func (f ValueOfFuturesField) Tag() quickfix.Tag { return tag.ValueOfFutures }

//NewValueOfFutures returns a new ValueOfFuturesField initialized with val
func NewValueOfFutures(val quickfix.FIXFloat) ValueOfFuturesField {
	return ValueOfFuturesField{val}
}

//VenueTypeField is a CHAR field
type VenueTypeField struct{ quickfix.FIXString }

//Tag returns tag.VenueType (1430)
func (f VenueTypeField) Tag() quickfix.Tag { return tag.VenueType }

//NewVenueType returns a new VenueTypeField initialized with val
func NewVenueType(val quickfix.FIXString) VenueTypeField {
	return VenueTypeField{val}
}

//VolatilityField is a FLOAT field
type VolatilityField struct{ quickfix.FIXFloat }

//Tag returns tag.Volatility (1188)
func (f VolatilityField) Tag() quickfix.Tag { return tag.Volatility }

//NewVolatility returns a new VolatilityField initialized with val
func NewVolatility(val quickfix.FIXFloat) VolatilityField {
	return VolatilityField{val}
}

//WaveNoField is a STRING field
type WaveNoField struct{ quickfix.FIXString }

//Tag returns tag.WaveNo (105)
func (f WaveNoField) Tag() quickfix.Tag { return tag.WaveNo }

//NewWaveNo returns a new WaveNoField initialized with val
func NewWaveNo(val quickfix.FIXString) WaveNoField {
	return WaveNoField{val}
}

//WorkingIndicatorField is a BOOLEAN field
type WorkingIndicatorField struct{ quickfix.FIXBoolean }

//Tag returns tag.WorkingIndicator (636)
func (f WorkingIndicatorField) Tag() quickfix.Tag { return tag.WorkingIndicator }

//NewWorkingIndicator returns a new WorkingIndicatorField initialized with val
func NewWorkingIndicator(val quickfix.FIXBoolean) WorkingIndicatorField {
	return WorkingIndicatorField{val}
}

//WtAverageLiquidityField is a PERCENTAGE field
type WtAverageLiquidityField struct{ quickfix.FIXFloat }

//Tag returns tag.WtAverageLiquidity (410)
func (f WtAverageLiquidityField) Tag() quickfix.Tag { return tag.WtAverageLiquidity }

//NewWtAverageLiquidity returns a new WtAverageLiquidityField initialized with val
func NewWtAverageLiquidity(val quickfix.FIXFloat) WtAverageLiquidityField {
	return WtAverageLiquidityField{val}
}

//XmlDataField is a DATA field
type XmlDataField struct{ quickfix.FIXString }

//Tag returns tag.XmlData (213)
func (f XmlDataField) Tag() quickfix.Tag { return tag.XmlData }

//NewXmlData returns a new XmlDataField initialized with val
func NewXmlData(val quickfix.FIXString) XmlDataField {
	return XmlDataField{val}
}

//XmlDataLenField is a LENGTH field
type XmlDataLenField struct{ quickfix.FIXInt }

//Tag returns tag.XmlDataLen (212)
func (f XmlDataLenField) Tag() quickfix.Tag { return tag.XmlDataLen }

//NewXmlDataLen returns a new XmlDataLenField initialized with val
func NewXmlDataLen(val quickfix.FIXInt) XmlDataLenField {
	return XmlDataLenField{val}
}

//YieldField is a PERCENTAGE field
type YieldField struct{ quickfix.FIXFloat }

//Tag returns tag.Yield (236)
func (f YieldField) Tag() quickfix.Tag { return tag.Yield }

//NewYield returns a new YieldField initialized with val
func NewYield(val quickfix.FIXFloat) YieldField {
	return YieldField{val}
}

//YieldCalcDateField is a LOCALMKTDATE field
type YieldCalcDateField struct{ quickfix.FIXString }

//Tag returns tag.YieldCalcDate (701)
func (f YieldCalcDateField) Tag() quickfix.Tag { return tag.YieldCalcDate }

//NewYieldCalcDate returns a new YieldCalcDateField initialized with val
func NewYieldCalcDate(val quickfix.FIXString) YieldCalcDateField {
	return YieldCalcDateField{val}
}

//YieldRedemptionDateField is a LOCALMKTDATE field
type YieldRedemptionDateField struct{ quickfix.FIXString }

//Tag returns tag.YieldRedemptionDate (696)
func (f YieldRedemptionDateField) Tag() quickfix.Tag { return tag.YieldRedemptionDate }

//NewYieldRedemptionDate returns a new YieldRedemptionDateField initialized with val
func NewYieldRedemptionDate(val quickfix.FIXString) YieldRedemptionDateField {
	return YieldRedemptionDateField{val}
}

//YieldRedemptionPriceField is a PRICE field
type YieldRedemptionPriceField struct{ quickfix.FIXFloat }

//Tag returns tag.YieldRedemptionPrice (697)
func (f YieldRedemptionPriceField) Tag() quickfix.Tag { return tag.YieldRedemptionPrice }

//NewYieldRedemptionPrice returns a new YieldRedemptionPriceField initialized with val
func NewYieldRedemptionPrice(val quickfix.FIXFloat) YieldRedemptionPriceField {
	return YieldRedemptionPriceField{val}
}

//YieldRedemptionPriceTypeField is a INT field
type YieldRedemptionPriceTypeField struct{ quickfix.FIXInt }

//Tag returns tag.YieldRedemptionPriceType (698)
func (f YieldRedemptionPriceTypeField) Tag() quickfix.Tag { return tag.YieldRedemptionPriceType }

//NewYieldRedemptionPriceType returns a new YieldRedemptionPriceTypeField initialized with val
func NewYieldRedemptionPriceType(val quickfix.FIXInt) YieldRedemptionPriceTypeField {
	return YieldRedemptionPriceTypeField{val}
}

//YieldTypeField is a STRING field
type YieldTypeField struct{ quickfix.FIXString }

//Tag returns tag.YieldType (235)
func (f YieldTypeField) Tag() quickfix.Tag { return tag.YieldType }

//NewYieldType returns a new YieldTypeField initialized with val
func NewYieldType(val quickfix.FIXString) YieldTypeField {
	return YieldTypeField{val}
}
